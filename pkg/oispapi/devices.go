package oispapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const getDevicesPath = "/v1/api/accounts/{accountId}/devices"

//GetDevices is retrieving the list of devices of active account
func (o *Oispapi) GetDevices() (*[]Device, error) {

	client := &http.Client{}

	replacements := map[string]string{"accountId": o.accounts[o.activAccount].ID}
	url := makeURL(o.url, getDevicesPath, replacements)
	req, _ := http.NewRequest("GET", url, nil)
	setAuthHeader(req, o.token)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(string(responseData))
	}
	responseObject := []Device{}
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(string(responseData))
	return &responseObject, nil
}
