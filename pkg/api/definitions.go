package api

import (
	"SteamPurchaseService/util"
	"net/url"
	"strconv"
)

type InitTxnRequest struct {
	SteamAccountID string `json:"SteamAccountID"`
	OrderID        string `json:"OrderID"`
	ItemID         int    `json:"ItemID"`
}

type FinalizeTxnRequest struct {
	OrderID string `json:"OrderID"`
}

//{"response":{"result":"OK","params":{"orderid":"321438899","transid":"4844195756231622061"}}}
type FinalizeTxnResponse struct {
	Response FinalizeTxnResponseResponse `json:"response"`
}

type FinalizeTxnResponseResponse struct {
	Result string                    `json:"result"`
	Params FinalizeTxnResponseParams `json:"params"`
}

type FinalizeTxnResponseParams struct {
	OrderId string `json:"orderid"`
	TransId string `json:"transid"`
}

func (t *InitTxnRequest) ToPostBody(config *util.Config, item util.Item) url.Values {
	postBody := url.Values{}
	postBody.Set("key", config.APIKey)
	postBody.Set("orderid", t.OrderID)
	postBody.Set("steamid", t.SteamAccountID)
	postBody.Set("appid", config.SteamAppID)
	postBody.Set("itemcount", "1")
	postBody.Set("currency", "USD")
	postBody.Set("language", "en")
	postBody.Set("usersession", "client")
	postBody.Set("itemid[0]", strconv.Itoa(item.ItemID))
	postBody.Set("qty[0]", "1")
	postBody.Set("amount[0]", item.Price)
	postBody.Set("description[0]", item.Description)
	postBody.Set("category[0]", item.Category)
	return postBody
}

func (t *FinalizeTxnRequest) ToPostBody(config *util.Config) url.Values {
	postBody := url.Values{}
	postBody.Set("key", config.APIKey)
	postBody.Set("orderid", t.OrderID)
	postBody.Set("appid", config.SteamAppID)
	return postBody
}
