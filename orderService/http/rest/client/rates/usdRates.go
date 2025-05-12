package rates

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

type USDRatesResponse struct {
	Rate float64
}

func (h HttpClient) GetUsdRate(currency string) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("%s/rates/%s", h.cnf.BackgroundServiceAddress, currency))
	if err != nil {
		return 0, fmt.Errorf("error while making request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("get rate for '%s' unexpected status code: %d", currency, resp.StatusCode)
	}

	var response USDRatesResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("error while decoding response body: %w", err)
	}

	return response.Rate, nil
}
