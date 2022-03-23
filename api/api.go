package api

import (
	posting "github.com/owsup-ru/owozon/api/posting"
	products "github.com/owsup-ru/owozon/api/products"
	returns "github.com/owsup-ru/owozon/api/returns"
)

type OzonAPI struct {
	AccountId string `json:"id"`
	Token     string `json:"token"`
}

// Получить список заказов по FBO
func (o OzonAPI) FboList(data posting.FboListPayload) (*posting.IFboList, error) {
	return posting.FboList(o.Token, o.AccountId, data)
}

// Получить список заказов по FBS
func (o OzonAPI) FbsList(data posting.FbsListPayload) (*posting.IFbsList, error) {
	return posting.FbsList(o.Token, o.AccountId, data)
}

// Получить список товаров
func (o OzonAPI) ProductsList(data products.PayloadList) (*products.IProductsList, error) {
	return products.ProductsList(o.Token, o.AccountId, data)
}

// Получить информацию о товаре
func (o OzonAPI) ProductInfo(data products.InfoPayload) (*products.IProductInfo, error) {
	return products.ProductInfo(o.Token, o.AccountId, data)
}

// Получить список товаров с описанием
func (o OzonAPI) ProductInfoList(data products.InfoListPayload) (*products.IProductInfoList, error) {
	return products.ProductInfoList(o.Token, o.AccountId, data)
}

// Получить описание товара
func (o OzonAPI) ProductInfoDescription(data products.InfoDescriptionPayload) (*products.IProductInfoDescription, error) {
	return products.ProductInfoDescription(o.Token, o.AccountId, data)
}

// Получить атрибуты товара
func (o OzonAPI) ProductInfoAttributes(data products.InfoAttributesPayload) (*products.IProductInfoAttributes, error) {
	return products.ProductInfoAttributes(o.Token, o.AccountId, data)
}

// Получить возвраты по FBO
func (o OzonAPI) GetReturnsFbo(data returns.ReturnsPayload) (*returns.ReturnResponseFbo, error) {
	return returns.ReturnsFbo(o.Token, o.AccountId, data)
}

// Получить возвраты по FBS
func (o OzonAPI) GetReturnsFbs(data returns.ReturnsPayload) (*returns.ReturnResponseFbs, error) {
	return returns.ReturnsFbs(o.Token, o.AccountId, data)
}
