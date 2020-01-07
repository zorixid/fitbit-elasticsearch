package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/fitbit-elastic/fitbit"
)

type Config struct {
	Client   string `yml:"client"`
	Callback string `yml:"callback"`
	Code     string `yml:"code"`
	Bearer   string `yml:"bearer"`
	Testing	 string `yml:testing"`
}

var c Config
var configPath = "fitbit-elastic.yml"

func main() {

	NewConfig()

	// WORKAROUND to set client ID as accesstoken because it's needed in auth.go
	err := os.Setenv("CLIENT_ID", c.Client)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Setenv("TESTING", c.Testing)
	if err != nil {
		fmt.Println(err)
	}

	if (c.Testing != "true") {
		// Create a request body with our required auth information.
		// Get auth and refresh token.
		// Save them as env variables.
		rb := fmt.Sprintf("expires_in=3600&grant_type=authorization_code&clientId=%s&code=%s&redirect_uri=%s", c.Client, c.Code, c.Callback)
		body := fitbit.AuthorizeUser(rb, c.Bearer)
		fitbit.SaveTokens(body)
	}

	// Now we're authorized! Let's get some FitBit data
	fitbit.GetActivity()
}

func NewConfig() {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TEST TEST TEST yml: %s\n\n\n\n", c)
}
