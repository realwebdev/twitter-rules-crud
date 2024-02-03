package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"os"

// 	"github.com/dghubble/oauth1"
// 	"github.com/gorilla/mux"
// )

// var oauthStore = make(map[string]string)

// var requestTokenURL = "https://api.twitter.com/oauth/request_token"
// var accessTokenURL = "https://api.twitter.com/oauth/access_token"
// var authorizeURL = "https://api.twitter.com/oauth/authorize"
// var showUserURL = "https://api.twitter.com/1.1/users/show.json"

// var (
// 	consumerKey    = os.Getenv("TWAUTH_APP_CONSUMER_KEY")
// 	consumerSecret = os.Getenv("TWAUTH_APP_CONSUMER_SECRET")
// 	callbackURL    = "http://localhost:8080/callback"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err := template.New("index").Parse("<h1>Hello, Twitter OAuth!</h1>")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, nil)
// }

// func startHandler(w http.ResponseWriter, r *http.Request) {
// 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// 	tempCred, requestsecret, err := config.RequestToken(http.DefaultClient, callbackURL, nil)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	oauthStore[tempCred.Token] = tempCred.Secret

// 	redirectURL := config.AuthorizationURL(tempCred, nil)
// 	http.Redirect(w, r, redirectURL.String(), http.StatusFound)
// }

// func callbackHandler(w http.ResponseWriter, r *http.Request) {
// 	values := r.URL.Query()
// 	oauthToken := values.Get("oauth_token")
// 	oauthVerifier := values.Get("oauth_verifier")
// 	oauthDenied := values.Get("denied")

// 	if oauthDenied != "" {
// 		delete(oauthStore, oauthDenied)
// 		http.Error(w, "OAuth request was denied by the user", http.StatusBadRequest)
// 		return
// 	}

// 	if oauthToken == "" || oauthVerifier == "" {
// 		http.Error(w, "Callback param(s) missing", http.StatusBadRequest)
// 		return
// 	}

// 	oauthTokenSecret, ok := oauthStore[oauthToken]
// 	if !ok {
// 		http.Error(w, "oauth_token not found locally", http.StatusBadRequest)
// 		return
// 	}

// 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// 	tokenCred := oauth1.NewToken(oauthToken, oauthTokenSecret)
// 	tokenCred.SetVerifier(oauthVerifier)

// 	httpClient := config.Client(oauth1.NoContext, tokenCred)
// 	accessToken, _, err := config.AccessToken(httpClient, tokenCred, oauthVerifier)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Fetch user details
// 	userInfoURL := fmt.Sprintf("%s?user_id=%s", showUserURL, accessToken.Extra["user_id"])
// 	resp, err := httpClient.Get(userInfoURL)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Handle the response from Twitter API
// 	// (Note: Additional error checking can be added based on the Twitter API response)
// 	var user map[string]interface{}
// 	json.NewDecoder(resp.Body).Decode(&user)

// 	tmpl, err := template.New("callback-success").Parse(`
// 		<h1>Welcome {{.ScreenName}}!</h1>
// 		<p>User ID: {{.UserID}}</p>
// 		<p>Name: {{.Name}}</p>
// 		<p>Friends Count: {{.FriendsCount}}</p>
// 		<p>Statuses Count: {{.StatusesCount}}</p>
// 		<p>Followers Count: {{.FollowersCount}}</p>
// 	`)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl.Execute(w, user)
// }

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", helloHandler)
// 	r.HandleFunc("/start", startHandler)
// 	r.HandleFunc("/callback", callbackHandler)

// 	http.Handle("/", r)

// 	fmt.Println("Server is running on :8080")
// 	http.ListenAndServe(":8080", nil)
// }
