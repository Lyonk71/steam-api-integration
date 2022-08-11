package api

import (
	"SteamPurchaseService/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// SteamInitTxn initializes a transaction with Steam
// Steam API Url: POST https://partner.steam-api.com/ISteamMicroTxn/InitTxn/v3/
// https://partner.steamgames.com/doc/webapi/ISteamMicroTxn#InitTxn
func SteamInitTxn(c *util.Config, v url.Values) (*InitTxnResponse, error) {
	url := fmt.Sprintf("%s%s/InitTxn/v3", c.SteamAPIUrl, c.SteamInterface)
	resp, err := http.PostForm(url, v)
	if err != nil {
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
		return nil, err
	}

	var initTxnResponse = new(InitTxnResponse)
	err = json.Unmarshal(body, &initTxnResponse)
	if err != nil {
		return nil, err
	}

	return initTxnResponse, nil
}

// SteamCheckAppOwnership makes sure that the player owns the game
// Steam API Url: GET https://partner.steam-api.com/ISteamUser/CheckAppOwnership/v2/
// https://partner.steamgames.com/doc/webapi/ISteamUser#CheckAppOwnership
func SteamCheckAppOwnership(c *util.Config, steamID string) (*CheckAppOwnershipResponse, error) {

	payload := url.Values{}
	payload.Add("key", c.APIKey)
	payload.Add("appid", c.SteamAppID)
	payload.Add("steamid", steamID)

	url := fmt.Sprintf("%s%s/CheckAppOwnership/v2", c.SteamAPIUrl, "ISteamUser")
	resp, err := http.Get(url + "?" + payload.Encode())
	if err != nil {
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
		return nil, err
	}

	var getCheckAppOwnershipResponse = new(CheckAppOwnershipResponse)
	err = json.Unmarshal(body, &getCheckAppOwnershipResponse)
	if err != nil {
		log.Printf("You are not listed as the publisher of this app.")
		return nil, err
	}

	return getCheckAppOwnershipResponse, nil
}

// GetUserInfo fetches the player's info based on their user wallet.
// This is useful for getting the player's local currency and purchasing status.
// Steam API Url: GET https://partner.steam-api.com/ISteamMicroTxn/GetUserInfo/v2/
// https://partner.steamgames.com/doc/webapi/ISteamMicroTxn#GetUserInfo
func SteamGetUserInfo(c *util.Config, steamID string) (*GetUserInfoResponse, error) {
	payload := url.Values{}
	payload.Add("key", c.APIKey)
	payload.Add("appid", c.SteamAppID)
	payload.Add("steamid", steamID)

	url := fmt.Sprintf("%s%s/GetUserInfo/v2", c.SteamAPIUrl, c.SteamInterface)
	resp, err := http.Get(url + "?" + payload.Encode())
	if err != nil {
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
		return nil, err
	}

	var getUserInfoResponse = new(GetUserInfoResponse)
	err = json.Unmarshal(body, &getUserInfoResponse)
	if err != nil {
		return nil, err
	}

	return getUserInfoResponse, nil
}

// FinalizeTxn will finalize the transaction by charging the player's wallet after they authorized the purchase
// Steam API Url: POST https://partner.steam-api.com/ISteamMicroTxn/FinalizeTxn/v2/
// https://partnaer.steamgames.com/doc/webapi/ISteamMicroTxn#FinalizeTxn
func SteamFinalizeTxn(c *util.Config, v url.Values) (*FinalizeTxnResponse, error) {
	url := fmt.Sprintf("%s%s/FinalizeTxn/v2", c.SteamAPIUrl, c.SteamInterface)
	resp, err := http.PostForm(url, v)
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

	var finalizeTxnResponse = new(FinalizeTxnResponse)
	err = json.Unmarshal(body, &finalizeTxnResponse)
	if err != nil {
		return nil, err
	}

	return finalizeTxnResponse, nil
}
