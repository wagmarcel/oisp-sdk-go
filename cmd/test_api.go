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
	fmt.Print("Received token ", api)

	api, err = oispapi.NewOispAPIFromToken(api.GetToken(), url)
	if err != nil {
		fmt.Println("Error while trying to get token", err)
	}
	fmt.Print("Received token ", api)
}
