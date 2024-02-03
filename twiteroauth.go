package main

// import (
// 	"fmt"
// 	"net/http"
// 	"net/url"

// 	"github.com/dghubble/oauth1"
// )

// func main3() {
// 	consumerKey := "azFiYjR0RXI0SDJQLVFxOUc0dS06MTpjaQ"
// 	consumerSecret := "OiLPYzmj6DjXIieQ6NenvdDKKuGmK1I0Ytjx0wSmVoIn3nh2G5"
// 	accessToken := "1531051229094240259-b3B00ZgzsanMfartm6sAM8hF3AVhye"
// 	accessSecret := "GH4o7d1p8UgSRMGVUnja318IRUdXBqRJKB7AijzeXiIsV"

// 	// OAuth1 config
// 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// 	token := oauth1.NewToken(accessToken, accessSecret)

// 	// Create an http.Client with the OAuth1 credentials
// 	httpClient := config.Client(oauth1.NoContext, token)

// 	// Set the callback URL and any additional parameters
// 	callbackURL := "http://localhost:8080/twitter/callback" // replace with your callback URL
// 	params := url.Values{}
// 	params.Set("oauth_callback", callbackURL)

// 	// Make the request to obtain the OAuth Request Token
// 	resp, err := httpClient.PostForm("https://api.twitter.com/oauth/request_token", params)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Check the response status code
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Println("Error:", resp.Status)

// 		// Print the response body for more details
// 		body := make([]byte, 1024)
// 		n, _ := resp.Body.Read(body)
// 		fmt.Println("Response Body:", string(body[:n]))
// 		return
// 	}
// 	// Parse the response body
// 	// Note: You might want to use a library like github.com/mrjones/oauth to handle OAuth responses
// 	// For simplicity, we are printing the response body here
// 	body := make([]byte, 1024)
// 	n, _ := resp.Body.Read(body)
// 	fmt.Println("Response:", string(body[:n]))
// }

// // package main

// // import (
// // 	"fmt"
// // 	"net/http"
// // 	"net/url"

// // 	"github.com/dghubble/oauth1"
// // )

// // func main() {
// // 	consumerKey := "azFiYjR0RXI0SDJQLVFxOUc0dS06MTpjaQ"
// // 	consumerSecret := "OiLPYzmj6DjXIieQ6NenvdDKKuGmK1I0Ytjx0wSmVoIn3nh2G5"

// // 	// OAuth1 config
// // 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// // 	requestToken, requestSecret, err := config.RequestToken()
// // 	if err != nil {
// // 		fmt.Println("Error obtaining request token:", err)
// // 		return
// // 	}

// // 	token := oauth1.NewToken(requestToken, requestSecret)

// // 	// Create an http.Client with the OAuth1 credentials
// // 	httpClient := config.Client(oauth1.NoContext, token)

// // 	// Set the callback URL and any additional parameters
// // 	callbackURL := "http://localhost:8080/twitter/callback" // replace with your callback URL
// // 	params := url.Values{}
// // 	params.Set("oauth_callback", callbackURL)

// // 	// Make the request to obtain the OAuth Request Token
// // 	resp, err := httpClient.PostForm("https://api.twitter.com/oauth/request_token", params)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		return
// // 	}
// // 	defer resp.Body.Close()

// // 	// Check the response status code
// // 	if resp.StatusCode != http.StatusOK {
// // 		fmt.Println("Error:", resp.Status)

// // 		// Print the response body for more details
// // 		body := make([]byte, 1024)
// // 		n, _ := resp.Body.Read(body)
// // 		fmt.Println("Response Body:", string(body[:n]))
// // 		return
// // 	}

// // 	// Parse the response body
// // 	body := make([]byte, 1024)
// // 	n, _ := resp.Body.Read(body)
// // 	fmt.Println("Response:", string(body[:n]))
// // }
