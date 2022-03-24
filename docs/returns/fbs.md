# Получить вазвраты по FBS

POST ```https://api-seller.ozon.ru/v2/returns/company/fbs```

## Пример использование

```go
package main

import (
	"log"

	owozon "github.com/owsup-ru/owozon"
	returns "github.com/owsup-ru/owozon/api/returns"
)

func main() {
	body := returns.ReturnsPayload{
    Limit:  1000,
    Offset: 0,
  }

	ozon := owozon.Init("clientid", "token")
	data, _ := ozon.GetReturnsFbs(body)

	log.Println(data)
}
```

## Ответ 


```go
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
```