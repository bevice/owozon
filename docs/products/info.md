# Информация о товаре
Возвращает информацию о товарах.

POST: ```v2/product/info```

## Пример использование

```go
body := api.InfoPayload{
  ProductId: 168397488, // Если нужно получить информацию о товаре по id 
  Sku: 168397488,       // Если нужно получить информацию о товаре по sku
  OfferId: "aaa-0001",  // Если нужно получить информацию о товаре по артиклу
}

data, err := api.ProductInfo(apikey, clientid, body)
```

## Ответ 

### go:

```go
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
```

### json:

```json
{
    "result": {
        "id": 0000,
        "name": "Название товара",
        "offer_id": "артикл",
        "barcode": "OZN000000000",
        "buybox_price": "",
        "category_id": 1000001960,
        "created_at": "2022-01-09T11:46:23.043099Z",
        "images": [
            "https://cdn1.ozone.ru/s3/multimedia-m/0000000000.jpg",
            "https://cdn1.ozone.ru/s3/multimedia-6/0000000001.jpg",
        ],
        "marketing_price": "1400.0000",
        "min_ozon_price": "",
        "old_price": "8990.0000",
        "premium_price": "",
        "price": "1400.0000",
        "recommended_price": "",
        "min_price": "",
        "sources": [
            {
                "is_enabled": true,
                "sku": 000000000,
                "source": "fbo"
            },
            {
                "is_enabled": true,
                "sku": 000000001,
                "source": "fbs"
            }
        ],
        "stocks": {
            "coming": 0,
            "present": 8,
            "reserved": 0
        },
        "errors": [],
        "vat": "0.0",
        "visible": true,
        "visibility_details": {
            "has_price": true,
            "has_stock": true,
            "active_product": false
        },
        "price_index": "0.00",
        "commissions": [
            {
                "percent": 5,
                "min_value": 0,
                "value": 70,
                "sale_schema": "fbo",
                "delivery_amount": 0,
                "return_amount": 0
            },
            {
                "percent": 5,
                "min_value": 0,
                "value": 70,
                "sale_schema": "fbs",
                "delivery_amount": 0,
                "return_amount": 0
            },
            {
                "percent": 5,
                "min_value": 0,
                "value": 70,
                "sale_schema": "rfbs",
                "delivery_amount": 0,
                "return_amount": 0
            }
        ],
        "volume_weight": 0.1,
        "is_prepayment": false,
        "is_prepayment_allowed": true,
        "images360": [],
        "color_image": "",
        "primary_image": "https://cdn1.ozone.ru/s3/multimedia-j/0000000000.jpg",
        "status": {
            "state": "price_sent",
            "state_failed": "",
            "moderate_status": "",
            "decline_reasons": [],
            "validation_state": "success",
            "state_name": "Продается",
            "state_description": "",
            "is_failed": false,
            "is_created": true,
            "state_tooltip": "",
            "item_errors": [],
            "state_updated_at": "2022-02-06T12:28:50.987931Z"
        },
        "state": "",
        "service_type": "IS_CODE_SERVICE",
        "fbo_sku": 000000000,
        "fbs_sku": 000000001
    }
}
```