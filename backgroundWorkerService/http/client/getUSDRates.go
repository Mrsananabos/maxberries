package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	BASE_URL   = "https://data.fixer.io/api/latest"
	PREFIX_KEY = "rates:USD_"
)

type Response struct {
	Success bool               `json:"success"`
	Rates   map[string]float64 `json:"rates"`
}

func GetUSDRates(ctx context.Context, redis *redis.Client, token string, ttl time.Duration) {
	address, err := url.Parse(BASE_URL)
	if err != nil {
		log.Printf("Can`t parse URL %v", err)
		return
	}

	q := address.Query()
	q.Set("access_key", token)
	address.RawQuery = q.Encode()

	resp, err := http.Get(address.String())
	if err != nil {
		log.Printf("Error while request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error while read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error, status code %s", resp.Status)
	}

	var fixerResponse Response
	err = json.Unmarshal(body, &fixerResponse)
	if err != nil {
		log.Printf("Error while decode response body to JSON: %v", err)
	}

	if fixerResponse.Success {
		for currency, rate := range fixerResponse.Rates {
			key := fmt.Sprintf("%s%s", PREFIX_KEY, currency)
			err = redis.Set(ctx, key, rate, ttl).Err()
			if err != nil {
				log.Printf("Can`t add to redis key=%s, value=%s", key, resp)
			}
		}
	} else {
		log.Println("Can`t get rates data")
	}
}
