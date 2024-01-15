package main

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/dghubble/oauth1"
)

func main() {
	// Replace with your actual Twitter API credentials
	consumerKey := "D22E1Pq81PU648TWTRHSRPDpu"
	consumerSecret := "jI0idplXT1i6KhKBl2k1DEum1L06mcnkdi1ZCxm19oL3DlYFDu"
	accessToken := "3298684828-RBvjAGn4nGy7zUS3zHcF3HJdcHZvd9ZKaj5xX62"
	accessSecret := "oWWWZ0WLpNlhXbOdXHkJZNBmTJTiLtWLElHGOoea6HMHd"

	// OAuth1 config
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter API endpoint - user search
	searchURL := "https://api.twitter.com/2/users/search"

	// Parameters for the user search query
	queryParams := url.Values{}
	queryParams.Set("query", "oncologist") // Replace USERNAME_TO_SEARCH with the username you want to search for
	searchURL += "?" + queryParams.Encode()

	// Make GET request
	resp, err := httpClient.Get(searchURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the user search response
	fmt.Println(string(body))
}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"

// 	"github.com/dghubble/oauth1"
// )

// func main() {
// 	// Replace with your actual Twitter API credentials
// 	consumerKey := "D22E1Pq81PU648TWTRHSRPDpu"
// 	consumerSecret := "jI0idplXT1i6KhKBl2k1DEum1L06mcnkdi1ZCxm19oL3DlYFDu"
// 	accessToken := "3298684828-RBvjAGn4nGy7zUS3zHcF3HJdcHZvd9ZKaj5xX62"
// 	accessSecret := "oWWWZ0WLpNlhXbOdXHkJZNBmTJTiLtWLElHGOoea6HMHd"

// 	// OAuth1 config
// 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// 	token := oauth1.NewToken(accessToken, accessSecret)
// 	httpClient := config.Client(oauth1.NoContext, token)

// 	// Twitter API endpoint
// 	url := "https://api.twitter.com/1.1/account/verify_credentials.json"

// 	// Make GET request
// 	resp, err := httpClient.Get(url)
// 	if err != nil {
// 		fmt.Println("Error making GET request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}
// 	fmt.Println(string(body))
// }
