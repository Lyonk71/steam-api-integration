// Package api provides support for business logic related to Steam functionality
package api

import (
	"SteamPurchaseService/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SteamService interface {
	InitTxn(initTxn InitTxnRequest) error
	GetUserInfo(initTxn InitTxnRequest) error
	FinalizeTxn(finalizeTxn FinalizeTxnRequest) error
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

func (s *steamService) InitTxn(initTxn InitTxnRequest) error {
	// Start by looking up the item def based on the itemid passed into the request
	//items := s.ItemDef.

	item := s.ItemDef[initTxn.ItemID]

	url := fmt.Sprintf("%s%s/InitTxn/v3", s.Config.SteamAPIUrl, s.Config.SteamInterface)
	resp, err := http.PostForm(
		url,
		initTxn.ToPostBody(s.Config, item))
	// Handle Error
	if err != nil {
		log.Printf("An Error Occurred %v", err)
		return nil
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Print(err)
		}
	}()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil
	}
	sb := string(body)
	fmt.Print(sb)

	return nil
}

func (s *steamService) GetUserInfo(initTxn InitTxnRequest) error {
	return nil
}

func (s *steamService) FinalizeTxn(finalizeTxn FinalizeTxnRequest) error {
	return nil
}
