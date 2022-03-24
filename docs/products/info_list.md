# Подробный список товаров
Возвращает массив товаров по их идентификаторам

POST: ```v2/product/info/list```


## Пример использование

```go
package main

import (
	"log"

	owozon "github.com/owsup-ru/owozon"
	api "github.com/owsup-ru/owozon/api/products"
)

func main() {
	body := api.InfoListPayload{
		ProductId: []int64{168397488}, // Если нужно получить информацию о товарах по id 
		Sku: []int64{168397488},       // Если нужно получить информацию о товарах по sku
		OfferId: []string{"aaa-0001"},  // Если нужно получить информацию о товарах по артиклам
	}

	ozon := owozon.Init("clientid", "token")

	data, _ := ozon.ProductInfoList(body)

	log.Println(data)
}
```

## Ответ 

### go:

```go
type IProductInfoList struct {
	Result struct {
		Items []struct {
			ID               int          
			Name             string       
			OfferID          string       
			Barcode          string       
			BuyboxPrice      string       
			CategoryID       int          
			CreatedAt        time.Time    
			Images           []interface{}
			MarketingPrice   string       
			MinPrice         string       
			OldPrice         string       
			PremiumPrice     string       
			Price            string       
			RecommendedPrice string       
			Sources          []struct {
				IsEnabled bool  
				Sku       int   
				Source    string
			}
			State  string 
			Stocks struct {
				Coming   int
				Present  int
				Reserved int
			} 
			Errors            []interface{}
			Vat               string       
			Visible           bool         
			VisibilityDetails struct {
				HasPrice      bool
				HasStock      bool
				ActiveProduct bool
				Reasons       struct {
				}
			}
			PriceIndex   string       
			Images360    []interface{}
			ColorImage   string       
			PrimaryImage string       
			Status       struct {
				State            string        
				StateFailed      string        
				ModerateStatus   string        
				DeclineReasons   []interface{} 
				ValidationState  string        
				StateName        string        
				StateDescription string        
				IsFailed         bool          
				IsCreated        bool          
				StateTooltip     string        
				ItemErrors       []interface{} 
				StateUpdatedAt   time.Time     
			}
		} 
	} 
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