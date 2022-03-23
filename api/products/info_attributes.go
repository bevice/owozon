package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Request body

type FilterInfoAttributes struct {
	OfferId    []string `json:"offer_id"`
	ProductId  []int64  `json:"product_id"`
	Visibility string   `json:"visibility"`
}

type InfoAttributesPayload struct {
	Filter  `json:"filter"`
	LastId  string `json:"last_id"`
	Limit   int64  `json:"limit"`
	SortBy  string `json:"sort_by"`
	Sortdir string `json:"sort_dir"`
}

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

func ProductInfoAttributes(ApiKey, ClientId string, data InfoAttributesPayload) (*IProductInfoAttributes, error) {
	payloadBytes, err := json.Marshal(data)

	if err != nil {
		return nil, errors.New("json parse error")
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api-seller.ozon.ru/v3/products/info/attributes", body)

	if err != nil {
		return nil, errors.New("make request error")
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", ApiKey)
	req.Header.Set("Client-Id", ClientId)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.New("make request error")
	}

	defer resp.Body.Close()

	result := IProductInfoAttributes{}

	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
