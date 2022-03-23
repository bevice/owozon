# Получить описание характеристик товара 

POST ```https://api-seller.ozon.ru/v3/products/info/attributes```

## Пример использование

```go
body := api.InfoAttributesPayload{}

data, err := api.ProductInfoAttributes(apikey, clientid, body)
```

## Ответ 

### go:

```go
// Response
type IProductInfoAttributes struct {
	Result []struct {
		ID            int    `json:"id"`
		Barcode       string `json:"barcode"`
		CategoryID    int    `json:"category_id"`
		Name          string `json:"name"`
		OfferID       string `json:"offer_id"`
		Height        int    `json:"height"`
		Depth         int    `json:"depth"`
		Width         int    `json:"width"`
		DimensionUnit string `json:"dimension_unit"`
		Weight        int    `json:"weight"`
		WeightUnit    string `json:"weight_unit"`
		Images        []struct {
			FileName string `json:"file_name"`
			Default  bool   `json:"default"`
			Index    int    `json:"index"`
		} `json:"images"`
		ImageGroupID string        `json:"image_group_id"`
		Images360    []interface{} `json:"images360"`
		PdfList      []interface{} `json:"pdf_list"`
		Attributes   []struct {
			AttributeID int `json:"attribute_id"`
			ComplexID   int `json:"complex_id"`
			Values      []struct {
				DictionaryValueID int    `json:"dictionary_value_id"`
				Value             string `json:"value"`
			} `json:"values"`
		} `json:"attributes"`
		ComplexAttributes []interface{} `json:"complex_attributes"`
		ColorImage        string        `json:"color_image"`
		LastID            string        `json:"last_id"`
	} `json:"result"`
	Total  int    `json:"total"`
	LastID string `json:"last_id"`
}
```