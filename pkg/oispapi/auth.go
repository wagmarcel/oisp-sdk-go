package oispapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type getTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type getTokenResponse struct {
	Token string `json:"token"`
}

type getAuthTokenInfoResponse struct {
	Payload struct {
		UserID   string    `json:"sub"`
		Accounts []Account `json:"accounts"`
	} `json:"payload"`
}

const getTokenPath = "/v1/api/auth/token"
const getAuthTokenInfoPath = "/v1/api/auth/tokenInfo"

//GetUserToken is creating the user token from username and password
//Caution: Should not be used for long term services! The token is only valid
//for limited time (e.g. 1h)
func (o *Oispapi) getUserToken(username string, password string) error {

	body := &getTokenRequest{Username: username, Password: password}
	jsonValue, _ := json.Marshal(body)
	response, err := http.Post(o.url+getTokenPath, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(string(responseData))
	}
	responseObject := getTokenResponse{}

	json.Unmarshal(responseData, &responseObject)
	o.token = responseObject.Token
	return nil
}

//getAuthTokenInfo is retrieving the details of a token
func (o *Oispapi) getAuthTokenInfo() error {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", o.url+getAuthTokenInfoPath, nil)
	setAuthHeader(req, o.token)
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(string(responseData))
	}
	responseObject := getAuthTokenInfoResponse{}

	json.Unmarshal(responseData, &responseObject)
	o.accounts = responseObject.Payload.Accounts
	o.userID = responseObject.Payload.UserID
	return nil
}
