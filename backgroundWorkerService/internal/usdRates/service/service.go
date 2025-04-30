package service

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/internal/usdRates/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	BASE_URL   = "http://data.fixer.io/api/latest"
	PREFIX_KEY = "rates:USD_"
)

type Service struct {
	Redis      *redis.Client
	Ttl        time.Duration
	FixerToken string
}

func NewService(config configs.Config, r *redis.Client) Service {
	return Service{
		Redis:      r,
		Ttl:        time.Duration(config.Redis.TTL) * time.Minute,
		FixerToken: config.FixerAccessToken,
	}
}

func (s Service) GetUSDRates(ctx context.Context) (model.USDRatesResponse, error) {
	address, err := url.Parse(BASE_URL)
	if err != nil {
		return model.USDRatesResponse{}, fmt.Errorf("can`t parse URL %v", err)
	}

	q := address.Query()
	q.Set("access_key", s.FixerToken)
	address.RawQuery = q.Encode()

	resp, err := http.Get(address.String())
	if err != nil {
		return model.USDRatesResponse{}, fmt.Errorf("error while request: %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.USDRatesResponse{}, fmt.Errorf("error while read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return model.USDRatesResponse{}, fmt.Errorf("error, status code %s", resp.Status)
	}

	var response model.USDRatesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return model.USDRatesResponse{}, fmt.Errorf("error while decode response body to JSON: %v", err)
	}

	if response.Success {
		for currency, rate := range response.Rates {
			key := fmt.Sprintf("%s%s", PREFIX_KEY, currency)
			err = s.Redis.Set(ctx, key, rate, s.Ttl).Err()
			if err != nil {
				log.Printf("Can`t add to redis key=%s, value=%s", key, resp)
			}
		}
	} else {
		return model.USDRatesResponse{}, fmt.Errorf("can`t get rates data")
	}

	return response, nil
}

func (s Service) GetUSDRate(ctx context.Context, currency string) (float64, error) {
	usdRate, err := s.getUSDRateCache(ctx, currency)
	if err != nil {
		log.Println(err.Error())
	} else {
		return usdRate, err
	}

	usdRates, err := s.GetUSDRates(ctx)
	if err != nil {
		return 0, err
	}

	rate, ok := usdRates.Rates[currency]
	if !ok {
		return 0, fmt.Errorf("not found range for %s", currency)
	}

	return rate, nil
}

func (s Service) getUSDRateCache(ctx context.Context, currency string) (float64, error) {
	redisKey := fmt.Sprintf("%s%s", PREFIX_KEY, currency)
	rsl := s.Redis.Get(ctx, redisKey)

	if err := rsl.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, fmt.Errorf("value %s not found in Redis", redisKey)
		} else {
			return 0, err
		}
	}

	valueFloat, err := rsl.Float64()
	if err != nil {
		return 0, fmt.Errorf("can`t parse %s to float64", rsl.Val())
	}

	return valueFloat, nil
}
