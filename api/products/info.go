package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// Request body
type InfoPayload struct {
	OfferId   string `json:"offer_id"`
	ProductId int64  `json:"product_id"`
	Sku       int64  `json:"sku"`
}

// Response
type IProductInfo struct {
	Result struct {
		ID               int       `json:"id"`
		Name             string    `json:"name"`
		OfferID          string    `json:"offer_id"`
		Barcode          string    `json:"barcode"`
		BuyboxPrice      string    `json:"buybox_price"`
		CategoryID       int       `json:"category_id"`
		CreatedAt        time.Time `json:"created_at"`
		Images           []string  `json:"images"`
		MarketingPrice   string    `json:"marketing_price"`
		MinOzonPrice     string    `json:"min_ozon_price"`
		OldPrice         string    `json:"old_price"`
		PremiumPrice     string    `json:"premium_price"`
		Price            string    `json:"price"`
		RecommendedPrice string    `json:"recommended_price"`
		MinPrice         string    `json:"min_price"`
		Sources          []struct {
			IsEnabled bool   `json:"is_enabled"`
			Sku       int    `json:"sku"`
			Source    string `json:"source"`
		} `json:"sources"`
		Stocks struct {
			Coming   int `json:"coming"`
			Present  int `json:"present"`
			Reserved int `json:"reserved"`
		} `json:"stocks"`
		Errors            []interface{} `json:"errors"`
		Vat               string        `json:"vat"`
		Visible           bool          `json:"visible"`
		VisibilityDetails struct {
			HasPrice      bool `json:"has_price"`
			HasStock      bool `json:"has_stock"`
			ActiveProduct bool `json:"active_product"`
		} `json:"visibility_details"`
		PriceIndex  string `json:"price_index"`
		Commissions []struct {
			Percent        int    `json:"percent"`
			MinValue       int    `json:"min_value"`
			Value          int    `json:"value"`
			SaleSchema     string `json:"sale_schema"`
			DeliveryAmount int    `json:"delivery_amount"`
			ReturnAmount   int    `json:"return_amount"`
		} `json:"commissions"`
		VolumeWeight        float64       `json:"volume_weight"`
		IsPrepayment        bool          `json:"is_prepayment"`
		IsPrepaymentAllowed bool          `json:"is_prepayment_allowed"`
		Images360           []interface{} `json:"images360"`
		ColorImage          string        `json:"color_image"`
		PrimaryImage        string        `json:"primary_image"`
		Status              struct {
			State            string        `json:"state"`
			StateFailed      string        `json:"state_failed"`
			ModerateStatus   string        `json:"moderate_status"`
			DeclineReasons   []interface{} `json:"decline_reasons"`
			ValidationState  string        `json:"validation_state"`
			StateName        string        `json:"state_name"`
			StateDescription string        `json:"state_description"`
			IsFailed         bool          `json:"is_failed"`
			IsCreated        bool          `json:"is_created"`
			StateTooltip     string        `json:"state_tooltip"`
			ItemErrors       []interface{} `json:"item_errors"`
			StateUpdatedAt   time.Time     `json:"state_updated_at"`
		} `json:"status"`
		State       string `json:"state"`
		ServiceType string `json:"service_type"`
		FboSku      int    `json:"fbo_sku"`
		FbsSku      int    `json:"fbs_sku"`
	} `json:"result"`
}

func ProductInfo(ApiKey, ClientId string, data InfoPayload) (*IProductInfo, error) {
	payloadBytes, err := json.Marshal(data)

	if err != nil {
		return nil, errors.New("json parse error")
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api-seller.ozon.ru/v2/product/info", body)

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

	result := IProductInfo{}

	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
