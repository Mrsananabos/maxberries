package model

type USDRatesResponse struct {
	Success bool               `json:"success"`
	Rates   map[string]float64 `json:"rates"`
}
