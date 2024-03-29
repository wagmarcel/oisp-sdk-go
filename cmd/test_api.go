package main

import (
	"fmt"
	"os"

	"github.com/oisp-sdk-go/pkg/oispapi"
)

func main() {

	username := os.Getenv("OISP_USERNAME")
	password := os.Getenv("OISP_PASSWORD")
	url := os.Getenv("OISP_URL")

	api, err := oispapi.NewOispAPIFromUser(username, password, url)
	if err != nil {
		fmt.Println("Error while trying to get token", err)
	}
	fmt.Println("Received token ", api)

	api, err = oispapi.NewOispAPIFromToken(api.GetToken(), url)
	if err != nil {
		fmt.Println("Error while trying to get token", err)
	}
	fmt.Println("Received token ", api)

	devices, err := api.GetDevices()
	if err != nil {
		fmt.Println("Error while trying to get devices", err)
	}
	fmt.Println("Retrieved devices: ", devices)

	device := oispapi.Device{
		DeviceID:  "11-22-33-44-55-66",
		Name:      "MyGoDevice",
		Tags:      []string{"hello", "world"},
		GatewayID: "mygogateway",
	}

	err = api.CreateDevice(&device)
	if err != nil {
		fmt.Println("Error while creating device:", err)
	} else {
		fmt.Println("Device created successfully")
	}

	newdevice, err := api.GetDevice(device.DeviceID)
	if err != nil {
		fmt.Println("Error while retrieving device: ", err)
	}
	fmt.Println("Device retrieved: ", newdevice)

	updateDevice := oispapi.Device{
		DeviceID:  device.DeviceID,
		Name:      "MyGoDevice2",
		Tags:      []string{"hello2", "world2"},
		GatewayID: "mygogateway2",
	}
	err = api.UpdateDevice(&updateDevice)
	if err != nil {
		fmt.Println("Error while updating device:", err)
	} else {
		fmt.Println("Device updated successfully")
	}

	err = api.DeleteDevice(device.DeviceID)
	if err != nil {
		fmt.Println("Error deleting device:", err)
	} else {
		fmt.Println("Device deleted successfully")
	}
}
