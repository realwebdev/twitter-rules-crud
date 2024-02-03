package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/dghubble/go-twitter/twitter"
// 	"github.com/dghubble/oauth1"
// 	"github.com/gorilla/mux"
// )

// var (
// 	consumerKey    = "azFiYjR0RXI0SDJQLVFxOUc0dS06MTpjaQ"
// 	consumerSecret = "OiLPYzmj6DjXIieQ6NenvdDKKuGmK1I0Ytjx0wSmVoIn3nh2G5"
// 	callbackURL    = "http://localhost:8080/twitter/callback"
// )

// func main() {
// 	router := mux.NewRouter()

// 	// Home page
// 	router.HandleFunc("/", HomeHandler).Methods("GET")

// 	// Twitter authentication
// 	router.HandleFunc("/twitter/login", TwitterLoginHandler).Methods("GET")
// 	router.HandleFunc("twitter/callback", CallbackHandler).Methods("GET")

// 	fmt.Println("Server is running on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Home page!")
// }

// func TwitterLoginHandler(w http.ResponseWriter, r *http.Request) {
// 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// 	// token := oauth1.NoContext

// 	requestToken, tt, err := config.RequestToken()
// 	log.Println(tt)
// 	if err != nil {
// 		http.Error(w, "Error getting request token", http.StatusInternalServerError)
// 		return
// 	}

// 	authorizationURL, err := config.AuthorizationURL(requestToken)
// 	if err != nil {
// 		http.Error(w, "Error getting authorization URL", http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, authorizationURL.String(), http.StatusSeeOther)
// }

// func CallbackHandler(w http.ResponseWriter, r *http.Request) {
// 	config := oauth1.NewConfig(consumerKey, consumerSecret)

// 	requestToken, verifier, err := oauth1.ParseAuthorizationCallback(r)
// 	if err != nil {
// 		http.Error(w, "Error parsing authorization callback", http.StatusInternalServerError)
// 		return
// 	}

// 	accessToken, accessSecret, err := config.AccessToken(requestToken, "", verifier)
// 	if err != nil {
// 		http.Error(w, "Error getting access token", http.StatusInternalServerError)
// 		return
// 	}

// 	client := config.Client(oauth1.NoContext, oauth1.NewToken(accessToken, accessSecret))
// 	twitterClient := twitter.NewClient(client)

// 	user, _, err := twitterClient.Accounts.VerifyCredentials(nil)
// 	if err != nil {
// 		http.Error(w, "Error verifying Twitter credentials", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "Logged in as %s", user.ScreenName)
// }
