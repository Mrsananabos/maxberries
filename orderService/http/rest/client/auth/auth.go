package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orderService/configs"
)

type Claims struct {
	Sub         string   `json:"sub"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

type HttpClient struct {
	cnf    configs.Services
	client *http.Client
}

func NewHttpClient(cnf configs.Services, client *http.Client) HttpClient {
	return HttpClient{
		cnf:    cnf,
		client: client,
	}
}

func (h HttpClient) GetAuthClaims(token string) (Claims, error) {
	var response Claims

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/me", h.cnf.AuthServiceAddress), nil)
	if err != nil {
		return response, fmt.Errorf("error while creating request: %w", err)
	}

	req.Header.Set("Authorization", token)

	resp, err := h.client.Do(req)
	if err != nil {
		return response, fmt.Errorf("error while making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response, fmt.Errorf("error while decoding response: %w", err)
	}

	return response, nil
}
