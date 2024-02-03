package main

// import (
// 	"crypto/hmac"
// 	"crypto/rand"
// 	"crypto/sha1"
// 	"encoding/base64"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"sort"
// 	"strings"
// 	"time"
// )

// const (
// 	consumerKey    = "azFiYjR0RXI0SDJQLVFxOUc0dS06MTpjaQ"
// 	consumerSecret = "OiLPYzmj6DjXIieQ6NenvdDKKuGmK1I0Ytjx0wSmVoIn3nh2G5"
// 	callbackURL    = "http://127.0.0.1:8080/twitter/callback"
// )

// func main() {
// 	// Step 1: Obtain a request token
// 	oauthToken, oauthTokenSecret, err := getRequestToken()
// 	if err != nil {
// 		fmt.Println("Error obtaining request token:", err)
// 		return
// 	}

// 	// Step 2: Redirect the user to Twitter for authorization
// 	authorizeURL := "https://api.twitter.com/oauth/authorize?oauth_token=" + oauthToken
// 	fmt.Println("Please authorize the application by visiting:", authorizeURL)

// 	// Step 3: Get user verifier
// 	fmt.Print("Enter the verifier: ")
// 	var verifier string
// 	fmt.Scanln(&verifier)

// 	// Step 4: Obtain an access token
// 	accessToken, accessTokenSecret, err := getAccessToken(oauthToken, oauthTokenSecret, verifier)
// 	if err != nil {
// 		fmt.Println("Error obtaining access token:", err)
// 		return
// 	}

// 	fmt.Println("Access Token:", accessToken)
// 	fmt.Println("Access Token Secret:", accessTokenSecret)
// }

// func getRequestToken() (string, string, error) {
// 	oauthNonce := generateNonce()
// 	oauthTimestamp := fmt.Sprint(time.Now().Unix())

// 	params := map[string]string{
// 		"oauth_callback":         callbackURL,
// 		"oauth_consumer_key":     consumerKey,
// 		"oauth_nonce":            oauthNonce,
// 		"oauth_signature_method": "HMAC-SHA1",
// 		"oauth_timestamp":        oauthTimestamp,
// 		"oauth_version":          "1.0",
// 	}

// 	baseURL := "https://api.twitter.com/oauth/request_token"
// 	oauthSignature := generateSignature("POST", baseURL, params)

// 	// Add the signature to the params
// 	params["oauth_signature"] = oauthSignature

// 	// Prepare the authorization header
// 	authHeader := generateAuthHeader(params)

// 	// Make the request to obtain the request token
// 	req, err := http.NewRequest("POST", baseURL, nil)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	// Set the authorization header
// 	req.Header.Set("Authorization", authHeader)

// 	// Print the request details
// 	fmt.Println("Request URL:", baseURL)
// 	fmt.Println("Request Headers:", req.Header)
// 	fmt.Println("Request Parameters:", params)
// 	requestBody := paramsToURLValues(params).Encode()
// 	fmt.Println("Request Body:", requestBody)

// 	// Make the request to obtain the request token
// 	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(requestBody))
// 	if err != nil {
// 		return "", "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", "", fmt.Errorf("failed to obtain request token, status code: %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	// Parse the response
// 	tokenValues, err := url.ParseQuery(string(body))
// 	if err != nil {
// 		return "", "", err
// 	}

// 	oauthToken := tokenValues.Get("oauth_token")
// 	oauthTokenSecret := tokenValues.Get("oauth_token_secret")

// 	return oauthToken, oauthTokenSecret, nil
// }

// func getAccessToken(requestToken, requestTokenSecret, verifier string) (string, string, error) {
// 	oauthNonce := generateNonce()
// 	oauthTimestamp := fmt.Sprint(time.Now().Unix())

// 	params := map[string]string{
// 		"oauth_consumer_key":     consumerKey,
// 		"oauth_nonce":            oauthNonce,
// 		"oauth_signature_method": "HMAC-SHA1",
// 		"oauth_timestamp":        oauthTimestamp,
// 		"oauth_version":          "1.0",
// 		"oauth_token":            requestToken,
// 		"oauth_verifier":         verifier,
// 	}

// 	baseURL := "https://api.twitter.com/oauth/access_token"
// 	oauthSignature := generateSignature("POST", baseURL, params)

// 	// Add the signature to the params
// 	params["oauth_signature"] = oauthSignature

// 	// Prepare the authorization header
// 	authHeader := generateAuthHeader(params)

// 	// Make the request to obtain the request token
// 	req, err := http.NewRequest("POST", baseURL, nil)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	// Set the authorization header
// 	req.Header.Set("Authorization", authHeader)

// 	// Print the request details
// 	fmt.Println("Request URL:", baseURL)
// 	fmt.Println("Request Headers:", req.Header)
// 	fmt.Println("Request Parameters:", params)
// 	requestBody := paramsToURLValues(params).Encode()
// 	fmt.Println("Request Body:", requestBody)

// 	// Make the request to obtain the access token
// 	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(requestBody))
// 	if err != nil {
// 		return "", "", err
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", "", fmt.Errorf("failed to obtain access token, status code: %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	// Parse the response
// 	tokenValues, err := url.ParseQuery(string(body))
// 	if err != nil {
// 		return "", "", err
// 	}

// 	accessToken := tokenValues.Get("oauth_token")
// 	accessTokenSecret := tokenValues.Get("oauth_token_secret")

// 	return accessToken, accessTokenSecret, nil
// }

// func generateNonce() string {
// 	// Generate a random string for the nonce
// 	b := make([]byte, 32)
// 	_, _ = rand.Read(b)
// 	return base64.StdEncoding.EncodeToString(b)
// }

// func generateSignature(method, baseURL string, params map[string]string) string {
// 	// Sort the parameters
// 	var keys []string
// 	for k := range params {
// 		keys = append(keys, k)
// 	}
// 	sort.Strings(keys)

// 	// Construct the parameter string
// 	var paramStrings []string
// 	for _, k := range keys {
// 		paramStrings = append(paramStrings, url.QueryEscape(k)+"="+url.QueryEscape(params[k]))
// 	}
// 	paramString := strings.Join(paramStrings, "&")

// 	// Construct the base string
// 	baseString := strings.Join([]string{
// 		method,
// 		url.QueryEscape(baseURL),
// 		url.QueryEscape(paramString),
// 	}, "&")

// 	// Construct the signing key
// 	signingKey := url.QueryEscape(consumerSecret) + "&"

// 	// Calculate the HMAC-SHA1 signature
// 	h := hmac.New(sha1.New, []byte(signingKey))
// 	h.Write([]byte(baseString))
// 	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

// 	return signature
// }

// func generateAuthHeader(params map[string]string) string {
// 	// Construct the OAuth header
// 	var authHeaderStrings []string
// 	for k, v := range params {
// 		authHeaderStrings = append(authHeaderStrings, k+"=\""+url.QueryEscape(v)+"\"")
// 	}
// 	authHeader := "OAuth " + strings.Join(authHeaderStrings, ", ")

// 	return authHeader
// }

// func paramsToURLValues(params map[string]string) url.Values {
// 	values := url.Values{}
// 	for key, value := range params {
// 		values.Add(key, value)

// 	}
// 	return values
// }
