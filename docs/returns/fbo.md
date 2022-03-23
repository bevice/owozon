# Получить вазвраты по FBO

POST ```https://api-seller.ozon.ru/v2/returns/company/fbo```

## Пример использование

```go
ozon := api.OzonAPI{AccountId: clientid, Token: apikey}

body := returns.ReturnsPayload{
  Limit:  1000,
  Offset: 0,
}

data, err := ozon.GetReturnsFbo(body)
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