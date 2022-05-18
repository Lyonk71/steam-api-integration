// Package api provides support for business logic related to Steam functionality
package api

import (
	"SteamPurchaseService/util"
	"fmt"
)

type SteamService interface {
	InitTxn(initTxn InitTxnRequest) (*InitTxnResponse, error)
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

// kday todo: return a struct type
func (s *steamService) InitTxn(initTxn InitTxnRequest) (*InitTxnResponse, error) {
	// First let's make check sure that the player owns the game!
	appOwnership, err := SteamCheckAppOwnership(s.Config, initTxn.SteamAccountID)
	if err != nil {
		// kday todo: build proper error response
		fmt.Print(err)
		return nil, err
	}

	if !appOwnership.AppOwnership.OwnsApp || appOwnership.AppOwnership.Result != "OK" {
		// kday todo: build proper error response
		fmt.Print("This player does not own the game")
		return nil, err
	}

	// Then we'll make sure that the player can
	userInfo, err := SteamGetUserInfo(s.Config, initTxn.SteamAccountID)
	if err != nil {
		// kday todo: build proper error response
		fmt.Print(err)
		return nil, err
	}

	// We check if the player's account is ok for making purchases and that their account is not locked.
	if userInfo.Response.Result != "OK" || userInfo.Response.Params.State == "Locked" {
		// kday todo: build proper error response
		fmt.Print("This player cannot make purchases")
		return nil, err
	}

	// Now we can make the initTxn request
	item := s.ItemDef[initTxn.ItemID]
	resp, err := SteamInitTxn(s.Config, initTxn.ToPostBody(s.Config, item))
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return resp, nil
}

func (s *steamService) FinalizeTxn(finalizeTxn FinalizeTxnRequest) (*FinalizeTxnResponse, error) {
	resp, err := SteamFinalizeTxn(s.Config, finalizeTxn.ToPostBody(s.Config))
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return resp, nil
}
