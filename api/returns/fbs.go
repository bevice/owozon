package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type ReturnResponseFbs struct {
	Result struct {
		Returns []struct {
			ID                         int64       `json:"id"`
			ClearingID                 int64       `json:"clearing_id"`
			PostingNumber              string      `json:"posting_number"`
			ProductID                  int         `json:"product_id"`
			Sku                        int         `json:"sku"`
			Status                     string      `json:"status"`
			ReturnsKeepingCost         int         `json:"returns_keeping_cost"`
			ReturnReasonName           string      `json:"return_reason_name"`
			ReturnDate                 time.Time   `json:"return_date"`
			Quantity                   int         `json:"quantity"`
			ProductName                string      `json:"product_name"`
			Price                      int         `json:"price"`
			WaitingForSellerDateTime   time.Time   `json:"waiting_for_seller_date_time"`
			ReturnedToSellerDateTime   time.Time   `json:"returned_to_seller_date_time"`
			LastFreeWaitingDay         time.Time   `json:"last_free_waiting_day"`
			IsOpened                   bool        `json:"is_opened"`
			PlaceID                    int         `json:"place_id"`
			CommissionPercent          int         `json:"commission_percent"`
			Commission                 int         `json:"commission"`
			PriceWithoutCommission     int         `json:"price_without_commission"`
			IsMoving                   bool        `json:"is_moving"`
			MovingToPlaceName          string      `json:"moving_to_place_name"`
			WaitingForSellerDays       int         `json:"waiting_for_seller_days"`
			PickingAmount              interface{} `json:"picking_amount"`
			AcceptedFromCustomerMoment interface{} `json:"accepted_from_customer_moment"`
			PickingTag                 interface{} `json:"picking_tag"`
		} `json:"returns"`
		Count int `json:"count"`
	} `json:"result"`
}

func ReturnsFbs(ApiKey, ClientId string, data ReturnsPayload) (*ReturnResponseFbs, error) {
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("json parse error")
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://api-seller.ozon.ru/v2/returns/company/fbs", body)
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

	result := ReturnResponseFbs{}

	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
