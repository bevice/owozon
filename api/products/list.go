package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Request body
type Filter struct {
	Visibility string   `json:"visibility"`
	ProductId  []string `json:"product_id"`
	OfferId    []string `json:"offer_id"`
}
type PayloadList struct {
	Filter Filter `json:"filter"`
	Limit  int    `json:"limit"`
	LastId string `json:"last_id "`
}

// Response
type IProductsList struct {
	Result struct {
		Items []struct {
			ProductID int    `json:"product_id"`
			OfferId   string `json:"offer_id"`
		} `json:"items"`
		Total  int    `json:"total"`
		LastID string `json:"last_id"`
	} `json:"result"`
}

func ProductsList(ApiKey, ClientId string, data PayloadList) (*IProductsList, error) {
	payloadBytes, err := json.Marshal(data)

	if err != nil {
		return nil, errors.New("json parse error")
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api-seller.ozon.ru/v2/product/list", body)

	if err != nil {
		return nil, errors.New("make request error")
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", ApiKey)
	req.Header.Set("Client-Id", ClientId)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.New("make request error")
	}

	defer resp.Body.Close()

	result := IProductsList{}

	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
