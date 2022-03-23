package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type ReturnsPayload struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ReturnResponseFbo struct {
	Returns []struct {
		ID                         int64     `json:"id"`
		Sku                        int       `json:"sku"`
		CompanyID                  int       `json:"company_id"`
		PostingNumber              string    `json:"posting_number"`
		AcceptedFromCustomerMoment time.Time `json:"accepted_from_customer_moment"`
		ReturnReasonName           string    `json:"return_reason_name"`
		IsOpened                   bool      `json:"is_opened"`
		StatusName                 string    `json:"status_name"`
		ReturnedToOzonMoment       time.Time `json:"returned_to_ozon_moment"`
		CurrentPlaceName           string    `json:"current_place_name"`
		DstPlaceName               string    `json:"dst_place_name"`
	} `json:"returns"`
	Count int `json:"count"`
}

func ReturnsFbo(ApiKey, ClientId string, data ReturnsPayload) (*ReturnResponseFbo, error) {
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("json parse error")
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://api-seller.ozon.ru/v2/returns/company/fbo", body)
	if err != nil {
		return nil, errors.New("make request error")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Client-Id", ClientId)
	req.Header.Set("Api-Key", ApiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("make request error")
	}

	defer resp.Body.Close()

	result := ReturnResponseFbo{}

	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
