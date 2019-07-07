package oispapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const getDevicesPath = "/v1/api/accounts/{accountId}/devices"
const getOneDevicePath = "/v1/api/accounts/{accountId}/devices/{deviceId}"

//GetDevices is retrieving the list of devices of active account
func (o *Oispapi) GetDevices() (*[]Device, error) {
	client := &http.Client{}
	replacements := map[string]string{"accountId": o.accounts[o.activAccount].ID}
	url := makeURL(o.url, getDevicesPath, replacements)
	req, _ := http.NewRequest("GET", url, nil)
	setHeaders(req, o.token)
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
	return &responseObject, nil
}

//CreateDevice is creating a new device
func (o *Oispapi) CreateDevice(device *Device) error {
	client := &http.Client{}

	replacements := map[string]string{"accountId": o.accounts[o.activAccount].ID}
	url := makeURL(o.url, getDevicesPath, replacements)
	jsonBody, _ := json.Marshal(device)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	setHeaders(req, o.token)
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 201 {
		return errors.New(string(responseData))
	}
	responseObject := []Device{}
	json.Unmarshal(responseData, &responseObject)
	return nil
}

//GetDevice is retrieving the list of devices of active account
func (o *Oispapi) GetDevice(deviceID string) (*Device, error) {
	client := &http.Client{}
	replacements := map[string]string{
		"accountId": o.accounts[o.activAccount].ID,
		"deviceId":  deviceID,
	}
	url := makeURL(o.url, getOneDevicePath, replacements)
	req, _ := http.NewRequest("GET", url, nil)
	setHeaders(req, o.token)
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
	responseObject := Device{}
	json.Unmarshal(responseData, &responseObject)
	return &responseObject, nil
}

//UpdateDevice is updating a new device
func (o *Oispapi) UpdateDevice(device *Device) error {
	client := &http.Client{}

	replacements := map[string]string{"accountId": o.accounts[o.activAccount].ID,
		"deviceId": device.DeviceID,
	}
	url := makeURL(o.url, getOneDevicePath, replacements)
	updatedDevice := Device(*device)
	updatedDevice.DeviceID = ""
	jsonBody, _ := json.Marshal(updatedDevice)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	setHeaders(req, o.token)
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
	responseObject := []Device{}
	json.Unmarshal(responseData, &responseObject)
	return nil
}

//DeleteDevice is deleting a device
func (o *Oispapi) DeleteDevice(deviceID string) error {
	client := &http.Client{}
	replacements := map[string]string{
		"accountId": o.accounts[o.activAccount].ID,
		"deviceId":  deviceID,
	}
	url := makeURL(o.url, getOneDevicePath, replacements)
	req, _ := http.NewRequest("DELETE", url, nil)
	setHeaders(req, o.token)
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return errors.New(string(responseData))
	}

	return nil
}
