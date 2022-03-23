# Подробный список товаров
Возвращает массив товаров по их идентификаторам

POST: ```v2/product/info/list```


## Пример использование

```go
body := api.InfoListPayload{
  ProductId: []int64{168397488}, // Если нужно получить информацию о товарах по id 
  Sku: []int64{168397488},       // Если нужно получить информацию о товарах по sku
  OfferId: []string{"aaa-0001"},  // Если нужно получить информацию о товарах по артиклам
}

data, err := api.ProductInfoList(apikey, clientid, body)
```

## Ответ 

### go:

```go
type IProductInfoList struct {
	Result struct {
		Items []struct {
			ID               int           `json:"id"`
			Name             string        `json:"name"`
			OfferID          string        `json:"offer_id"`
			Barcode          string        `json:"barcode"`
			BuyboxPrice      string        `json:"buybox_price"`
			CategoryID       int           `json:"category_id"`
			CreatedAt        time.Time     `json:"created_at"`
			Images           []interface{} `json:"images"`
			MarketingPrice   string        `json:"marketing_price"`
			MinPrice         string        `json:"min_price"`
			OldPrice         string        `json:"old_price"`
			PremiumPrice     string        `json:"premium_price"`
			Price            string        `json:"price"`
			RecommendedPrice string        `json:"recommended_price"`
			Sources          []struct {
				IsEnabled bool   `json:"is_enabled"`
				Sku       int    `json:"sku"`
				Source    string `json:"source"`
			} `json:"sources"`
			State  string `json:"state"`
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
				Reasons       struct {
				} `json:"reasons"`
			} `json:"visibility_details"`
			PriceIndex   string        `json:"price_index"`
			Images360    []interface{} `json:"images360"`
			ColorImage   string        `json:"color_image"`
			PrimaryImage string        `json:"primary_image"`
			Status       struct {
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
		} `json:"items"`
	} `json:"result"`
}
```

### json:

```json
{
  "result": {
    "items": [{
        "id": 00000000,
        "name": "Название товара",
        "offer_id": "артикул",
        "barcode": "",
        "buybox_price": "",
        "category_id": 93726157,
        "created_at": "2021-06-03T03:40:05.871465Z",
        "images": [],
        "marketing_price": "",
        "min_price": "",
        "old_price": "1000.0000",
        "premium_price": "590.0000",
        "price": "690.0000",
        "recommended_price": "",
        "sources": [
          { "is_enabled": true, "sku": 000000000, "source": "fbo" },
          { "is_enabled": true, "sku": 000000000, "source": "fbs" }
        ],
        "state": "",
        "stocks": { "coming": 0, "present": 13, "reserved": 0 },
        "errors": [],
        "vat": "0.0",
        "visible": true,
        "visibility_details": {
          "has_price": false,
          "has_stock": true,
          "active_product": false,
          "reasons": {}
        },
        "price_index": "0.00",
        "images360": [],
        "color_image": "",
        "primary_image": "https://cdn1.ozone.ru/s3/multimedia-y/0000000000.jpg",
        "status": {
          "state": "price_sent",
          "state_failed": "",
          "moderate_status": "approved",
          "decline_reasons": [],
          "validation_state": "success",
          "state_name": "Продается",
          "state_description": "",
          "is_failed": false,
          "is_created": true,
          "state_tooltip": "",
          "item_errors": [],
          "state_updated_at": "2021-07-26T04:50:08.486697Z"
        }
      }]
  }
}
```