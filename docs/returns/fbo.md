# Получить вазвраты по FBO

POST ```https://api-seller.ozon.ru/v2/returns/company/fbo```

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
	data, _ := ozon.GetReturnsFbo(body)

	log.Println(data)
}
```


## Ответ 


```go
type ReturnResponseFbo struct {
	Returns []struct {
		ID                         int64     `json:"id"`
		Sku                        int       `json:"sku"`
		CompanyID                  int       `json:"company_id"`
		PostingNumber              string    `json:"posting_number"`
		AcceptedFromCustomerMoment time.Time `json:"accepted_from_customer_moment"`
		ReturnReasonName           string    `json:"return_reason_name"`
		IsOpened                   bool      `json:"is_opened"`
		StatusName                 string    `json:"status_name"`
		ReturnedToOzonMoment       time.Time `json:"returned_to_ozon_moment"`
		CurrentPlaceName           string    `json:"current_place_name"`
		DstPlaceName               string    `json:"dst_place_name"`
	} `json:"returns"`
	Count int `json:"count"`
}
```