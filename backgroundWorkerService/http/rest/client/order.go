package client

import (
	"backgroundWorkerService/configs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct {
	Cnf        configs.Services
	HttpClient *http.Client
}

func NewHttpClient(cnf configs.Services) HttpClient {
	return HttpClient{
		Cnf:        cnf,
		HttpClient: http.DefaultClient,
	}
}

type OrderPriceInfo struct {
	Status        string  `json:"status,omitempty"`
	TotalPrice    float64 `json:"total_price,omitempty"`
	DeliveryPrice float64 `json:"delivery_price,omitempty"`
}

func (h HttpClient) UpdateOrderPriceInfo(id int64, orderInfo OrderPriceInfo) error {
	body, err := json.Marshal(orderInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal orderInfo to JSON: %w", err)
	}

	url := fmt.Sprintf("%s/orders/%d", h.Cnf.OrderServiceAddress, id)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := h.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute patch request: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}
	return nil
}
