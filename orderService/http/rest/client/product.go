package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orderService/configs"
)

type HttpClient struct {
	cnf configs.Services
}

func NewHttpClient(cnf configs.Services) HttpClient {
	return HttpClient{
		cnf: cnf,
	}
}

type ProductInfoResponse struct {
	Price float64
}

func (h HttpClient) GetProductPrice(id int64) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("%s/products/%d", h.cnf.CatalogServiceAddress, id))
	if err != nil {
		return 0, fmt.Errorf("error while making request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response ProductInfoResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("error while decoding response body: %w", err)
	}

	return response.Price, nil
}
