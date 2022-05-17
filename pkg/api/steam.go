// Package api provides support for business logic related to Steam functionality
package api

import (
	"SteamPurchaseService/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type SteamService interface {
	InitTxn(initTxn InitTxnRequest) error
	FinalizeTxn(finalizeTxn FinalizeTxnRequest) error
}

type steamService struct {
	Config *util.Config
}

func NewSteamService(config *util.Config) SteamService {
	return &steamService{
		Config: config,
	}
}

func (s *steamService) InitTxn(initTxn InitTxnRequest) error {
	// kday TODO: redo all this to be flexible and dynamic
	postBody := url.Values{}
	postBody.Set("key", s.Config.APIKey)
	postBody.Set("orderid", initTxn.OrderID)
	postBody.Set("steamid", initTxn.SteamAccountID)
	postBody.Set("appid", s.Config.SteamAppID)
	postBody.Set("itemcount", "1")
	postBody.Set("currency", "USD")
	postBody.Set("language", "en")
	postBody.Set("usersession", "client")
	postBody.Set("itemid[0]", initTxn.ItemID)
	postBody.Set("qty[0]", "1")
	postBody.Set("amount[0]", "100USD")
	postBody.Set("description[0]", "Some Spellcraft Points")
	postBody.Set("category[0]", "Points")

	resp, err := http.PostForm("https://partner.steam-api.com/ISteamMicroTxnSandbox/InitTxn/v3", postBody)
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

func (s *steamService) FinalizeTxn(finalizeTxn FinalizeTxnRequest) error {
	return nil
}
