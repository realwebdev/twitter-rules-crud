package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrjones/oauth"
)

// const (
// 	consumerKey    = "azFiYjR0RXI0SDJQLVFxOUc0dS06MTpjaQ"
// 	consumerSecret = "OiLPYzmj6DjXIieQ6NenvdDKKuGmK1I0Ytjx0wSmVoIn3nh2G5"
// 	callbackURL    = "http://127.0.0.1:8080/twitter/callback"
// )

// var (
// 	consumerKey     = "qOrLVqeIaRSNQ4rdP7B42wcyje"
// 	consumerSecret  = "Ui8FyTdhkgPArAlBGqOtsK3LQNL4A3eAYyjBE4l4zTmG296aSRe"
// 	requestTokenURL = "https://api.twitter.com/oauth/request_token"
// 	authorizeURL    = "https://api.twitter.com/oauth/authorize"
)

var oauthStore = make(map[string]string)

func start(c *gin.Context) {
	appCallbackURL := "http://127.0.0.1:8080/twitter/callback" // Update with your actual callback URL

	consumer := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   requestTokenURL,
			AuthorizeTokenUrl: authorizeURL,
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)

	requestToken, url, err := consumer.GetRequestTokenAndUrl(appCallbackURL)
	if err != nil {
		errorMessage := fmt.Sprintf("Error getting request token: %s", err.Error())
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error_message": errorMessage})
		return
	}

	oauthStore[requestToken.Token] = requestToken.Secret

	c.JSON(http.StatusOK, gin.H{
		"authorize_url":     url,
		"oauth_token":       requestToken.Token,
		"request_token_url": requestTokenURL,
	})
}

func callback(c *gin.Context) {
	oauthToken := c.Query("oauth_token")
	oauthVerifier := c.Query("oauth_verifier")

	// Retrieve the corresponding secret from the store
	oauthTokenSecret, ok := oauthStore[oauthToken]
	if !ok {
		errorMessage := "Invalid or expired OAuth token"
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	// Exchange the request token and verifier for an access token
	consumer := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   requestTokenURL,
			AuthorizeTokenUrl: authorizeURL,
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)

	accessToken, err := consumer.AuthorizeToken(&oauth.RequestToken{Token: oauthToken},
		oauthVerifier,
	)
	if err != nil {
		errorMessage := fmt.Sprintf("Error exchanging request token for access token: %s", err.Error())
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error_message": errorMessage})
		return
	}

	// Use the access token to make authenticated requests to the Twitter API
	userURL := "https://api.twitter.com/1.1/account/verify_credentials.json?include_email=true"
	httpClient, err := consumer.MakeHttpClient(accessToken)
	if err != nil {

	}
	resp, err := httpClient.Get(userURL)
	if err != nil {
		errorMessage := fmt.Sprintf("Error retrieving user credentials: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": errorMessage})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errorMessage := fmt.Sprintf("Error reading response body: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": errorMessage})
		return
	}

	// // Respond with the user's details including email
	// c.JSON(http.StatusOK, gin.H{
	// 	"access_token":        accessToken.Token,
	// 	"access_token_secret": accessToken.Secret,
	// 	"oauth_verifier":      oauthVerifier,
	// 	"oauth_token":         oauthToken,
	// 	"oauth_token_secret":  oauthTokenSecret,
	// 	"user_details":        string(responseBody),
	// })

	// Now you have the accessToken and accessTokenSecret
	// You can use them to make authenticated requests to the Twitter API

	c.JSON(http.StatusOK, gin.H{
		"access_token":        accessToken.Token,
		"access_token_secret": accessToken.Secret,
		"oauth_verifier":      oauthVerifier,
		"oauth_token":         oauthToken,
		"oauth_token_secret":  oauthTokenSecret,
		"request_token_url":   requestTokenURL,
		"authorize_token_url": authorizeURL,
		"access_token_url":    "https://api.twitter.com/oauth/access_token",
		"authenticated_user":  accessToken.AdditionalData["screen_name"],
		"user_details":        string(responseBody),
	})
}

// Inside the main function, add the callback route
func main() {
	router := gin.Default()

	// ... (previous code)

	router.GET("/start", start)
	router.GET("/twitter/callback", callback) // Add the callback route

	// ... (other routes)

	// Run the server
	router.Run(":8080")
}

// func main() {
// 	router := gin.Default()
// 	// router.LoadHTMLGlob("templates/*")

// 	router.GET("/start", start)

// 	// Add other routes as needed

// 	// Run the server
// 	router.Run(":8080")
// }
