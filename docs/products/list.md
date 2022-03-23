# Список товаров
Возвращает список товаров

POST: ```v2/product/list```

## Пример использование

```go
body := api.PayloadList{
  Limit: 100,         // Максимальное количество товаров за запрос (Минимум — 1, Максимум — 1000)
  Filter: api.Filter{
    Visibility: "ALL",                              // Фильтр по видимости товара
    ProductId:  []string{"000000000", "000000001"}, // Если нужно взять товары с такими id
		OfferId:    []string{"aa-000", "aa-001"},       // Если нужно взять товары с такими артиклами
  },
  LastId: "WzIxNzYwNjMwOSwyMTc2MDYzMDld"            // Используеться для того что бы получить вторую страницу товаров ( лимит 100, а товаров 200, потребуеться послать запрос два раза )
}

data, err := api.ProductsList(apikey, clientid, body)
```

Фильтр по видимости товара:

- `ALL` — все товары.
- `VISIBLE` — товары, которые видны покупателям.
- `INVISIBLE` — товары, которые по какой-то из причин не видны покупателям.
- `EMPTY_STOCK` — товары, у которых не указано наличие.
- `NOT_MODERATED` — товары, которые не прошли модерацию.
- `MODERATED` — товары, которые прошли модерацию.
- `DISABLED` — товары, которые видны покупателям, но недоступны к покупке.
- `STATE_FAILED` — товары, создание которых завершилось ошибĸой.
- `READY_TO_SUPPLY` — товары, готовые к поставке.
- `VALIDATION_STATE_PENDING` — товары, которые проходят проверку на премодерации (валидатором).
- `VALIDATION_STATE_FAIL` — товары, которые не прошли проверку на премодерации (валидатором).
- `VALIDATION_STATE_SUCCESS` — товары, которые прошли проверку на премодерации (валидатором).
- `TO_SUPPLY` — товары, готовые к продаже.
- `IN_SALE` — товары в продаже.
- `REMOVED_FROM_SALE` — товары, скрытые от покупателей.
- `BANNED` - товары в бане.
- `OVERPRICED`.
- `CRITICALLY_OVERPRICED`.
- `EMPTY_BARCODE`.
- `BARCODE_EXISTS`.
- `QUARANTINE`.
- `ARCHIVED` — товары в архиве.
- `OVERPRICED_WITH_STOCK` — товары в продаже со стоимостью выше, чем у конкурентов.
- `PARTIAL_APPROVED` — товары в продаже с пустым или неполным описанием.
- `IMAGE_ABSENT`.
- `MODERATION_BLOCK`.


## Ответ 

### go:

```go
type IProductsList struct {
	Result struct {
		Items []struct {
			ProductID int    `json:"product_id"`
			OfferId   string `json:"offer_id"`
		} `json:"items"`
		Total  int    `json:"total"`
		LastID string `json:"last_id"`
	} `json:"result"`
}
```

### json:

```json
{
  "result": {
    "items": [{
      "product_id": 223681945,
      "offer_id": "136748"
    }],
    "total": 1,
    "last_id": "bnVсbA=="
  }
}
```