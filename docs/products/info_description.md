# Получить описание товара 
Возвращает информация о товаре

POST: ```v1/product/info/description```


## Пример использование

```go
package main

import (
	"log"

	owozon "github.com/owsup-ru/owozon"
	api "github.com/owsup-ru/owozon/api/products"
)

func main() {
	body := api.InfoDescriptionPayload{
    OfferId: "aaa-0001",  // Если нужно получить информацию о товаре по артиклу
    ProductId: 168397488, // Если нужно получить информацию о товаре по id 
  }

	ozon := owozon.Init("clientid", "token")

	data, _ := ozon.ProductInfoList(body)

	log.Println(data)
}
```


## Ответ 

### go:

```go
type IProductInfoDescription struct {
	Result struct {
		ID          int   
		OfferID     string
		Name        string
		Description string
	} 
}
```

### json:

```json
{
  "result": {
    "id": 73453843,
    "offer_id": "артикл",
    "name": "Название товара",
    "description": "Описание товара <html>"
  }
}
```