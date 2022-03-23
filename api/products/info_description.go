package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Request body
type InfoDescriptionPayload struct {
	OfferId   string `json:"offer_id"`
	ProductId int64  `json:"product_id"`
}

// Response
type IProductInfoDescription struct {
	Result struct {
		ID          int    `json:"id"`
		OfferID     string `json:"offer_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"result"`
}

func ProductInfoDescription(ApiKey, ClientId string, data InfoDescriptionPayload) (*IProductInfoDescription, error) {
	payloadBytes, err := json.Marshal(data)

	if err != nil {
		return nil, errors.New("json parse error")
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api-seller.ozon.ru/v1/product/info/description", body)

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

	result := IProductInfoDescription{}

	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
