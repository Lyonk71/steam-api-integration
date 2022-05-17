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
	OrderID string `json:"orderId"`
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
