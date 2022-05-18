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

// InitTxn initializes a transaction with Steam
// Steam API Url: POST https://partner.steam-api.com/ISteamMicroTxn/InitTxn/v3/
// https://partner.steamgames.com/doc/webapi/ISteamMicroTxn#InitTxn
func SteamInitTxn(c *util.Config, v url.Values) (string, error) {
	url := fmt.Sprintf("%s%s/InitTxn/v3", c.SteamAPIUrl, c.SteamInterface)
	resp, err := http.PostForm(url, v)
	if err != nil {
		return "", err
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
		return "", err
	}

	sb := string(body)
	return sb, nil
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

	sb := string(body)
	fmt.Print(sb)

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
	sb := string(body)
	fmt.Print(sb)

	var finalizeTxnResponse = new(FinalizeTxnResponse)
	err = json.Unmarshal(body, &finalizeTxnResponse)
	if err != nil {
		return nil, err
	}

	return finalizeTxnResponse, nil
}
