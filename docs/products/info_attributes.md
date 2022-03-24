# Получить описание характеристик товара 

POST ```https://api-seller.ozon.ru/v3/products/info/attributes```

## Пример использование

```go
package main

import (
	"log"

	owozon "github.com/owsup-ru/owozon"
	api "github.com/owsup-ru/owozon/api/products"
)

func main() {
	body := api.InfoAttributesPayload{}

	ozon := owozon.Init("clientid", "token")

	data, _ := ozon.ProductInfoAttributes(body)

	log.Println(data)
}
```


## Ответ 

### go:

```go
// Response
type IProductInfoAttributes struct {
	Result []struct {
		ID            int   
		Barcode       string
		CategoryID    int   
		Name          string
		OfferID       string
		Height        int   
		Depth         int   
		Width         int   
		DimensionUnit string
		Weight        int   
		WeightUnit    string
		Images        []struct {
			FileName string
			Default  bool  
			Index    int   
		}
		ImageGroupID string       
		Images360    []interface{}
		PdfList      []interface{}
		Attributes   []struct {
			AttributeID int
			ComplexID   int
			Values      []struct {
				DictionaryValueID int   
				Value             string
			}
		}
		ComplexAttributes []interface{}
		ColorImage        string       
		LastID            string       
	}
	Total  int   
	LastID string
}
```