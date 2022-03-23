package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FbsListInProcessAt struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type FbsListOrderCreatedAt struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type FbsListUpdatedAt struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type FbsListFilter struct {
	InProcessAt    FbsListInProcessAt    `json:"in_process_at"`
	OrderCreatedAt FbsListOrderCreatedAt `json:"order_created_at"`
	Since          string                `json:"since"`
	To             string                `json:"to"`
	UpdatedAt      FbsListUpdatedAt      `json:"updated_at"`
}
type FbsListWith struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes      bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
}

type FbsListPayload struct {
	Filter FbsListFilter `json:"filter"`
	Limit  int           `json:"limit"`
	Offset int           `json:"offset"`
	With   FbsListWith   `json:"with"`
}

type IFbsList struct {
	Result []struct {
		OrderID        int64     `json:"order_id"`
		OrderNumber    string    `json:"order_number"`
		PostingNumber  string    `json:"posting_number"`
		Status         string    `json:"status"`
		CancelReasonID int64     `json:"cancel_reason_id"`
		CreatedAt      time.Time `json:"created_at"`
		InProcessAt    time.Time `json:"in_process_at"`
		Products       []struct {
			Sku          int64         `json:"sku"`
			Name         string        `json:"name"`
			Quantity     int           `json:"quantity"`
			OfferID      string        `json:"offer_id"`
			Price        string        `json:"price"`
			DigitalCodes []interface{} `json:"digital_codes"`
		} `json:"products"`
		AnalyticsData struct {
			Region               string `json:"region"`
			City                 string `json:"city"`
			DeliveryType         string `json:"delivery_type"`
			IsPremium            bool   `json:"is_premium"`
			PaymentTypeGroupName string `json:"payment_type_group_name"`
			WarehouseID          int64  `json:"warehouse_id"`
			WarehouseName        string `json:"warehouse_name"`
			IsLegal              bool   `json:"is_legal"`
		} `json:"analytics_data"`
		FinancialData struct {
			Products []struct {
				CommissionAmount     float64     `json:"commission_amount"`
				CommissionPercent    float64     `json:"commission_percent"`
				Payout               float64     `json:"payout"`
				ProductID            float64     `json:"product_id"`
				OldPrice             float64     `json:"old_price"`
				Price                float64     `json:"price"`
				TotalDiscountValue   float64     `json:"total_discount_value"`
				TotalDiscountPercent float64     `json:"total_discount_percent"`
				Actions              []string    `json:"actions"`
				Picking              interface{} `json:"picking"`
				Quantity             int         `json:"quantity"`
				ClientPrice          string      `json:"client_price"`
				ItemServices         struct {
					MarketplaceServiceItemFulfillment                float64 `json:"marketplace_service_item_fulfillment"`
					MarketplaceServiceItemPickup                     float64 `json:"marketplace_service_item_pickup"`
					MarketplaceServiceItemDropoffPvz                 float64 `json:"marketplace_service_item_dropoff_pvz"`
					MarketplaceServiceItemDropoffSc                  float64 `json:"marketplace_service_item_dropoff_sc"`
					MarketplaceServiceItemDropoffFf                  float64 `json:"marketplace_service_item_dropoff_ff"`
					MarketplaceServiceItemDirectFlowTrans            float64 `json:"marketplace_service_item_direct_flow_trans"`
					MarketplaceServiceItemReturnFlowTrans            float64 `json:"marketplace_service_item_return_flow_trans"`
					MarketplaceServiceItemDelivToCustomer            float64 `json:"marketplace_service_item_deliv_to_customer"`
					MarketplaceServiceItemReturnNotDelivToCustomer   float64 `json:"marketplace_service_item_return_not_deliv_to_customer"`
					MarketplaceServiceItemReturnPartGoodsCustomer    float64 `json:"marketplace_service_item_return_part_goods_customer"`
					MarketplaceServiceItemReturnAfterDelivToCustomer float64 `json:"marketplace_service_item_return_after_deliv_to_customer"`
				} `json:"item_services"`
			} `json:"products"`
			PostingServices struct {
				MarketplaceServiceItemFulfillment                float64 `json:"marketplace_service_item_fulfillment"`
				MarketplaceServiceItemPickup                     float64 `json:"marketplace_service_item_pickup"`
				MarketplaceServiceItemDropoffPvz                 float64 `json:"marketplace_service_item_dropoff_pvz"`
				MarketplaceServiceItemDropoffSc                  float64 `json:"marketplace_service_item_dropoff_sc"`
				MarketplaceServiceItemDropoffFf                  float64 `json:"marketplace_service_item_dropoff_ff"`
				MarketplaceServiceItemDirectFlowTrans            float64 `json:"marketplace_service_item_direct_flow_trans"`
				MarketplaceServiceItemReturnFlowTrans            float64 `json:"marketplace_service_item_return_flow_trans"`
				MarketplaceServiceItemDelivToCustomer            float64 `json:"marketplace_service_item_deliv_to_customer"`
				MarketplaceServiceItemReturnNotDelivToCustomer   float64 `json:"marketplace_service_item_return_not_deliv_to_customer"`
				MarketplaceServiceItemReturnPartGoodsCustomer    float64 `json:"marketplace_service_item_return_part_goods_customer"`
				MarketplaceServiceItemReturnAfterDelivToCustomer float64 `json:"marketplace_service_item_return_after_deliv_to_customer"`
			} `json:"posting_services"`
		} `json:"financial_data"`
		AdditionalData []interface{} `json:"additional_data"`
	} `json:"result"`
}

func FbsList(ApiKey, ClientId string, data FbsListPayload) (*IFbsList, error) {
	time_since, _ := time.Parse("2006-01-02", data.Filter.Since)
	time_to, _ := time.Parse("2006-01-02", data.Filter.To)

	to := time_to.Format("2006-01-02T00:00:00.000Z")
	from := time_since.Format("2006-01-02T00:00:00.000Z")

	data.Filter.Since = from
	data.Filter.To = to

	data.Filter.InProcessAt.From = from
	data.Filter.InProcessAt.To = to

	data.Filter.OrderCreatedAt.From = from
	data.Filter.OrderCreatedAt.To = to

	data.Filter.UpdatedAt.From = from
	data.Filter.UpdatedAt.To = to

	payloadBytes, err := json.Marshal(data)

	if err != nil {
		return nil, errors.New("json parse error")
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api-seller.ozon.ru/v2/posting/fbs/list", body)

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

	result := IFbsList{}

	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)

	if err != nil {
		return nil, err
	}

	fmt.Printf("result: %v\n", result)

	return &result, err
}
