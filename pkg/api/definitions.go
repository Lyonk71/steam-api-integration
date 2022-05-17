package api

type InitTxnRequest struct {
	SteamAccountID string `json:"steamAccountId"`
	OrderID        string `json:"orderId"`
	ItemID         string `json:"itemId"`
}

type FinalizeTxnRequest struct {
	OrderID string `json:"orderId"`
}
