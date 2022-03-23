package owozon

import "github.com/owsup-ru/owozon/api"

func Init(clientid, apikey string) api.OzonAPI {
	return api.OzonAPI{AccountId: clientid, Token: apikey}
}
