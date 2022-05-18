// Package api provides support for business logic related to Steam functionality
package api

import (
	"SteamPurchaseService/util"
	"fmt"
)

type SteamService interface {
	InitTxn(initTxn InitTxnRequest) error
	FinalizeTxn(finalizeTxn FinalizeTxnRequest) (*FinalizeTxnResponse, error)
}

type steamService struct {
	Config  *util.Config
	ItemDef util.Items
}

func NewSteamService(config *util.Config, items util.Items) SteamService {
	return &steamService{
		Config:  config,
		ItemDef: items,
	}
}

// InitTxn initializes a transaction with Steam
// Steam API Url: POST https://partner.steam-api.com/ISteamMicroTxn/InitTxn/v3/
// https://partner.steamgames.com/doc/webapi/ISteamMicroTxn#InitTxn
func (s *steamService) InitTxn(initTxn InitTxnRequest) error {
	userInfo, err := SteamGetUserInfo(s.Config, initTxn.SteamAccountID)
	if err != nil {
		fmt.Print(err)
		return err
	}

	// kday todo: validate userInfo
	fmt.Print(userInfo)

	// Start by looking up the item def based on the itemid passed into the request
	item := s.ItemDef[initTxn.ItemID]
	resp, err := SteamInitTxn(s.Config, initTxn.ToPostBody(s.Config, item))
	if err != nil {
		fmt.Print(err)
		return err
	}

	// kday todo: do something with the response
	fmt.Print(resp)
	return nil
}

func (s *steamService) FinalizeTxn(finalizeTxn FinalizeTxnRequest) (*FinalizeTxnResponse, error) {
	resp, err := SteamFinalizeTxn(s.Config, finalizeTxn.ToPostBody(s.Config))
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return resp, nil
}
