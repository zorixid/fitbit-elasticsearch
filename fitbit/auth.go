// This file holds all of the code we need to authorize our API usage
// AuthorizeUser() --> Get an access token and refresh token
// saveTokens() --> Save tokens as env variables
// isTokenActive() --> Determines if a response is valid
// refreshAuthToken() --> Refreshes and saves new access and refresh tokens

package fitbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Idk how else to pass this variable into this file yet. Dumb workaround.
var bearer string

// Use the FitBit API to get an access token and request token
// Returns the tokens in json form
func AuthorizeUser(b string, bear string) []byte {
	// sets bear to the global bearer. Dumb workaround.
	bearer = bear
	url := "https://api.fitbit.com/oauth2/token"
	requestBody := strings.NewReader(b)
	client := http.Client{}

	// Create a new request and add the appropriate headers
	request, err := http.NewRequest("POST", url, requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "Basic "+bearer)

	// Pass the request to our client and execute it
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	// Close the request when the func ends
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// for now, we'll print the response. Then, return it.
	fmt.Printf("Oauth response: %s\n", body)
	return body
}

// parses the access and refresh tokens into a struct
// saves the tokens as env variables
func SaveTokens(a []byte) {
	// create a token struct
	type Token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		UserID       string `json:"user_id"` // can probs remove this later
	}
	var t Token

	// unmarshal the JSON into the token struct
	err := json.Unmarshal([]byte(a), &t)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// save data as ENV variables
	err = os.Setenv("ACCESS_TOKEN", t.AccessToken)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Setenv("REFRESH_TOKEN", t.RefreshToken)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Setenv("USER_ID", t.UserID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Env variables saved successfully")
}

// Unmarshals the FitBit API response body to determine if the auth token is still valid
func isTokenActive(str string) bool {
	type FitBit struct {
		Success bool `json:"success"`
	}

	// Set a default value of true. If the value is missing, this wont overwrite it
	fb := FitBit{Success: true}
	err := json.Unmarshal([]byte(str), &fb)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// right now, we just print the error if something went wrong
	// in the future, I'll want to fix this
	if fb.Success == false {
		fmt.Printf("Request failed: %s\n", str)
	}

	return fb.Success
}

// Refreshes the auth token
func refreshAuthToken() {
	fmt.Println("Refreshing Auth Token")

	// Get the current variables, and create a request body from them
	access_token := os.Getenv("ACCESS_TOKEN")
	refresh_token := os.Getenv("REFRESH_TOKEN")
	// clientID := os.Getenv("CLIENT_ID")
	rb := fmt.Sprintf("grant_type=refresh_token&refresh_token=%s&clientId=%s", refresh_token, bearer)

	// Using the AuthrizeUser() func, let's refresh our auth token
	body := AuthorizeUser(rb, access_token)
	// fmt.Printf("Print body: %s\n", string(body))

	// Let's see if it worked
	if isTokenActive(string(body)) {
		// If it did, save the new tokens
		SaveTokens(body)
	} else {
		// If it didn't, panic
		// In the future, maybe change this
		panic("Failed to get new Auth Token")
	}
}
