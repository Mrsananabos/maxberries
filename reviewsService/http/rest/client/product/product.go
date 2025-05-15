package product

import (
	"fmt"
	"net/http"
	"reviewsService/configs"
)

type HttpClient struct {
	cnf configs.Services
}

func NewHttpClient(cnf configs.Services) HttpClient {
	return HttpClient{
		cnf: cnf,
	}
}

func (h HttpClient) GetProductById(id int64) error {
	resp, err := http.Get(fmt.Sprintf("%s/products/%d", h.cnf.CatalogServiceAddress, id))
	if err != nil {
		return fmt.Errorf("error while making request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil

}
