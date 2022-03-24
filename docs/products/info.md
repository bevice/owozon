# Информация о товаре
Возвращает информацию о товарах.

POST: ```v2/product/info```

## Пример использование

```go
package main

import (
	"log"

	owozon "github.com/owsup-ru/owozon"
	api "github.com/owsup-ru/owozon/api/products"
)

func main() {
	body := api.InfoPayload{
      ProductId: 168397488, // Если нужно получить информацию о товаре по id 
      Sku: 168397488,       // Если нужно получить информацию о товаре по sku
      OfferId: "aaa-0001",  // Если нужно получить информацию о товаре по артиклу
    }

	ozon := owozon.Init("clientid", "token")
	data, _ := ozon.ProductInfo(body)

	log.Println(data)
}
```

## Ответ 

### go:

```go
type IProductInfo struct {
	Result struct {
		ID               int      
		Name             string   
		OfferID          string   
		Barcode          string   
		BuyboxPrice      string   
		CategoryID       int      
		CreatedAt        time.Time
		Images           []string 
		MarketingPrice   string   
		MinOzonPrice     string   
		OldPrice         string   
		PremiumPrice     string   
		Price            string   
		RecommendedPrice string   
		MinPrice         string   
		Sources          []struct {
			IsEnabled bool   
			Sku       int    
			Source    string 
		}
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
		}
		PriceIndex  string 
		Commissions []struct {
			Percent        int   
			MinValue       int   
			Value          int   
			SaleSchema     string
			DeliveryAmount int   
			ReturnAmount   int   
		}
		VolumeWeight        float64       
		IsPrepayment        bool          
		IsPrepaymentAllowed bool          
		Images360           []interface{} 
		ColorImage          string        
		PrimaryImage        string        
		Status              struct {
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
		State       string 
		ServiceType string 
		FboSku      int    
		FbsSku      int    
	}
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