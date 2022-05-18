package api

import (
	"SteamPurchaseService/util"
	"net/url"
	"strconv"
)

// InitTxnRequest is the request data from the game client
type InitTxnRequest struct {
	SteamAccountID string `json:"SteamAccountID"`
	OrderID        string `json:"OrderID"`
	ItemID         int    `json:"ItemID"`
}

// FinalizeTxnRequest is the request data from the game client
type FinalizeTxnRequest struct {
	OrderID string `json:"OrderID"`
}

// FinalizeTxnResponse is the object we build when we receive the result from the Steam API
// Example succesful json response: {"response":{"result":"OK","params":{"orderid":"321438899","transid":"4844195756231622061"}}}
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

// GetUserInfoResponse: Result from steam when calling GetUserInfo
/*
{
    "response": {
        "result": "OK",
        "params": {
            "state": "NY",
            "country": "US",
            "currency": "USD",
            "status": "Trusted"
        }
    }
}
*/
type GetUserInfoResponse struct {
	Response GetUserInfoResponseResponse `json:"response"`
}

type GetUserInfoResponseResponse struct {
	Result string                    `json:"result"`
	Params GetUserInfoResponseParams `json:"params"`
}

type GetUserInfoResponseParams struct {
	State    string `json:"state"`
	Country  string `json:"country"`
	Currency string `json:"currency"`
	Status   string `json:"status"`
}

// ToPostBody for InitTxnRequest is the Steam API request params
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

// ToPostBody for FinalizeTxnRequest is the Steam API request params
func (t *FinalizeTxnRequest) ToPostBody(config *util.Config) url.Values {
	postBody := url.Values{}
	postBody.Set("key", config.APIKey)
	postBody.Set("orderid", t.OrderID)
	postBody.Set("appid", config.SteamAppID)
	return postBody
}
