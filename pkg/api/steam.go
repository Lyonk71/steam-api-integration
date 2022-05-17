// Package api provides support for business logic related to Steam functionality
package api

import (
	"SteamPurchaseService/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SteamService interface {
	InitTxn(initTxn InitTxnRequest) error
	GetUserInfo(initTxn InitTxnRequest) error
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

func (s *steamService) InitTxn(initTxn InitTxnRequest) error {
	// Start by looking up the item def based on the itemid passed into the request
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

func (s *steamService) FinalizeTxn(finalizeTxn FinalizeTxnRequest) (*FinalizeTxnResponse, error) {
	url := fmt.Sprintf("%s%s/FinalizeTxn/v2", s.Config.SteamAPIUrl, s.Config.SteamInterface)
	resp, err := http.PostForm(url, finalizeTxn.ToPostBody(s.Config))
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	// Handle Error
	if err != nil {
		log.Printf("An Error Occurred %v", err)
		return nil, err
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
		return nil, err
	}
	sb := string(body)
	fmt.Print(sb)

	var finalizeTxnResponse = new(FinalizeTxnResponse)
	err = json.Unmarshal(body, &finalizeTxnResponse)
	if err != nil {
		return nil, err
	}

	return finalizeTxnResponse, nil
}
