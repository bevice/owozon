# Получить список заказов (fbo)

POST ```https://api-seller.ozon.ru/v2/posting/fbo/list```

## Пример использование

```go
package main

import (
	"log"

	owozon "github.com/owsup-ru/owozon"
	posting "github.com/owsup-ru/owozon/api/posting"
)

func main() {
	body := posting.FboListPayload{
		Limit:  1000,
		Offset: 0,
		With: posting.FboListWith{
			AnalyticsData: true,
			FinancialData: true,
		},
		Filter: posting.FboListFilter{
			Since: "2022-01-01",
			To:    "2022-02-19",
		},
	}


	ozon := owozon.Init("clientid", "token")

	data, _ := ozon.FboList(body)

	log.Println(data)
}
```

## Ответ 

### go:

```go
// Response

type IItemService struct {
	MarketplaceServiceItemFulfillment                float64
	MarketplaceServiceItemPickup                     float64
	MarketplaceServiceItemDropoffPvz                 float64
	MarketplaceServiceItemDropoffSc                  float64
	MarketplaceServiceItemDropoffFf                  float64
	MarketplaceServiceItemDirectFlowTrans            float64
	MarketplaceServiceItemReturnFlowTrans            float64
	MarketplaceServiceItemDelivToCustomer            float64
	MarketplaceServiceItemReturnNotDelivToCustomer   float64
	MarketplaceServiceItemReturnPartGoodsCustomer    float64
	MarketplaceServiceItemReturnAfterDelivToCustomer float64
}

type IFboList struct {
	Result []struct {
		OrderID        int64     
		OrderNumber    string    
		PostingNumber  string    
		Status         string    
		CancelReasonID int64     
		CreatedAt      time.Time 
		InProcessAt    time.Time 
		Products       []struct {
			Sku          int64        
			Name         string       
			Quantity     int          
			OfferID      string       
			Price        string       
			DigitalCodes []interface{}
		}
		AnalyticsData struct {
			Region               string
			City                 string
			DeliveryType         string
			IsPremium            bool  
			PaymentTypeGroupName string
			WarehouseID          int64 
			WarehouseName        string
			IsLegal              bool  
		}
		FinancialData struct {
			Products []struct {
				CommissionAmount     float64     
				CommissionPercent    float64     
				Payout               float64     
				ProductID            float64     
				OldPrice             float64     
				Price                float64     
				TotalDiscountValue   float64     
				TotalDiscountPercent float64     
				Actions              []string    
				Picking              interface{} 
				Quantity             int         
				ClientPrice          string      
				ItemServices         IItemService
			}
			PostingServices IItemService
		}
		AdditionalData []interface{}
	}
}

```