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

const getTokenPath = "/v1/api/auth/token"

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

	json.Unmarshal(responseData, &responseObject) //.Decode(&responseObject)
	o.token = responseObject.Token
	return nil
}
