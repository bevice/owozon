# Получить описание товара 
Возвращает информация о товаре

POST: ```v1/product/info/description```


## Пример использование

```go
body := api.InfoDescriptionPayload{
  OfferId: "aaa-0001",  // Если нужно получить информацию о товаре по артиклу
  ProductId: 168397488, // Если нужно получить информацию о товаре по id 
}

data, err := api.ProductInfoList(apikey, clientid, body)
```

## Ответ 

### go:

```go
type IProductInfoDescription struct {
	Result struct {
		ID          int    `json:"id"`
		OfferID     string `json:"offer_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"result"`
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