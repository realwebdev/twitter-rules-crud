package main

// import (
// 	"crypto/rand"
// 	"encoding/base64"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/oauth2"
// )

// var (
// 	clientID            = "azFiYjR0RXI0SDJQLVFxOUc0dS06MTpjaQ"
// 	clientSecret        = "OiLPYzmj6DjXIieQ6NenvdDKKuGmK1I0Ytjx0wSmVoIn3nh2G5"
// 	redirectURL         = "http://127.0.0.1:8080/api/twitterCallback"
// 	codeChallengeMethod = "plain" // or "S256" for PKCE
// )

// var oauthConfig = &oauth2.Config{
// 	ClientID:     clientID,
// 	ClientSecret: clientSecret,
// 	RedirectURL:  redirectURL,
// 	Endpoint: oauth2.Endpoint{
// 		AuthURL:  "https://twitter.com/i/oauth2/authorize",
// 		TokenURL: "https://api.twitter.com/oauth2/token",
// 	},
// 	Scopes: []string{"tweet.read", "users.read", "offline.access"},
// }

// func generateRandomString(length int) (string, error) {
// 	bytes := make([]byte, length)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
// }

// func start(c *gin.Context) {
// 	state, err := generateRandomString(32)
// 	if err != nil {
// 		errorMessage := fmt.Sprintf("Error generating random string: %s", err.Error())
// 		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error_message": errorMessage})
// 		return
// 	}

// 	codeChallenge, err := generateRandomString(43) // 43 bytes base64 URL encoded = 32 bytes random + padding
// 	if err != nil {
// 		errorMessage := fmt.Sprintf("Error generating code challenge: %s", err.Error())
// 		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error_message": errorMessage})
// 		return
// 	}

// 	authURL := oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline,
// 		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
// 		oauth2.SetAuthURLParam("code_challenge_method", codeChallengeMethod),
// 	)

// 	c.Redirect(http.StatusTemporaryRedirect, authURL)
// }

// func callback(c *gin.Context) {
// 	code := c.Query("code")
// 	state := c.Query("state")

// 	// Validate state parameter to prevent CSRF attacks
// 	// (You should store and verify this on your server)
// 	if state != "your_generated_state" {
// 		errorMessage := "Invalid state parameter"
// 		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error_message": errorMessage})
// 		return
// 	}

// 	token, err := oauthConfig.Exchange(oauth2.NoContext, code)
// 	if err != nil {
// 		errorMessage := fmt.Sprintf("Error exchanging authorization code for access token: %s", err.Error())
// 		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error_message": errorMessage})
// 		return
// 	}

// 	// Use the access token to make authenticated requests to the Twitter API
// 	// ...

// 	// Respond with the user's details
// 	c.JSON(http.StatusOK, gin.H{
// 		"access_token": token.AccessToken,
// 		"user_details": "user_details_placeholder", // replace with actual user details
// 	})
// }

// func main() {
// 	router := gin.Default()

// 	router.GET("/start", start)
// 	router.GET("/twitter/callback", callback)

// 	router.Run(":8080")
// }
