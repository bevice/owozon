package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type FboListWith struct {
	AnalyticsData bool `json:"analytics_data"`
	FinancialData bool `json:"financial_data"`
}

type FboListFilter struct {
	Since string `json:"since"`
	To    string `json:"to"`
}

type FboListPayload struct {
	Filter FboListFilter `json:"filter"`
	Limit  int           `json:"limit"`
	Offset int           `json:"offset"`
	With   FboListWith   `json:"with"`
}

type IFboList struct {
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

func FboList(ApiKey, ClientId string, data FboListPayload) (*IFboList, error) {

	time_since, _ := time.Parse("2006-01-02", data.Filter.Since)
	time_to, _ := time.Parse("2006-01-02", data.Filter.To)

	data.Filter.Since = time_since.Format("2006-01-02T00:00:00.000Z")
	data.Filter.To = time_to.Format("2006-01-02T00:00:00.000Z")

	payloadBytes, err := json.Marshal(data)

	if err != nil {
		return nil, errors.New("json parse error")
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api-seller.ozon.ru/v2/posting/fbo/list", body)

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

	result := IFboList{}

	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
