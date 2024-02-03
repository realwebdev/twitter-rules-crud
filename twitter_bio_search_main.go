package main

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/url"
// 	"os"

// 	"github.com/dghubble/oauth1"
// )

// func main() {
// 	consumerKey := "D22E1Pq81PU648TWTRHSRPDpu"
// 	consumerSecret := "jI0idplXT1i6KhKBl2k1DEum1L06mcnkdi1ZCxm19oL3DlYFDu"
// 	accessToken := "3298684828-RBvjAGn4nGy7zUS3zHcF3HJdcHZvd9ZKaj5xX62"
// 	accessSecret := "oWWWZ0WLpNlhXbOdXHkJZNBmTJTiLtWLElHGOoea6HMHd"

// 	// OAuth1 config
// 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// 	token := oauth1.NewToken(accessToken, accessSecret)
// 	httpClient := config.Client(oauth1.NoContext, token)

// 	// Twitter API endpoint - user search
// 	searchURL := "https://api.twitter.com/2/users/search"

// 	// Parameters for the user search query
// 	queryParams := url.Values{}
// 	queryParams.Set("query", "oncologist") // Replace USERNAME_TO_SEARCH with the username you want to search for
// 	searchURL += "?" + queryParams.Encode()

// 	var nicknames []string

// 	// Make successive GET requests until there is no next page token
// 	// Make successive GET requests until there is no next page token
// 	for {
// 		// Make GET request
// 		resp, err := httpClient.Get(searchURL)
// 		if err != nil {
// 			fmt.Println("Error making GET request:", err)
// 			return
// 		}

// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println("Error reading response body:", err)
// 			return
// 		}

// 		// Unmarshal the response
// 		var response map[string]interface{}
// 		if err := json.Unmarshal(body, &response); err != nil {
// 			fmt.Println("Error unmarshalling response:", err)
// 			return
// 		}

// 		// Check if the "errors" field exists and is not nil
// 		if errors, ok := response["errors"].([]interface{}); ok && len(errors) > 0 {
// 			// Extract error details
// 			if errorDetails, ok := errors[0].(map[string]interface{}); ok {
// 				parameters := errorDetails["parameters"]
// 				message := errorDetails["message"]

// 				fmt.Printf("Error from Twitter API:\nParameters: %v\nMessage: %v\n", parameters, message)
// 			} else {
// 				fmt.Println("Unknown error details from Twitter API")
// 			}
// 			return
// 		}

// 		// Check if the "data" field exists and is not nil
// 		if data, ok := response["data"].([]interface{}); ok && data != nil {
// 			// Extract nicknames from the current page and add to the list
// 			for _, user := range data {
// 				if username, ok := user.(map[string]interface{})["username"].(string); ok {
// 					nicknames = append(nicknames, username)
// 				}
// 			}
// 		}

// 		// Check if the "meta" field exists and is not nil
// 		if meta, ok := response["meta"].(map[string]interface{}); ok && meta != nil {
// 			// Check if there is a next page
// 			nextToken, hasNext := meta["next_token"].(string)
// 			if !hasNext {
// 				break
// 			}

// 			// Update the searchURL with the next page token
// 			queryParams.Set("next_token", nextToken)
// 			searchURL = "https://api.twitter.com/2/users/search?" + queryParams.Encode()
// 		} else {
// 			// Break the loop if "meta" field is nil
// 			break
// 		}

// 		defer resp.Body.Close()
// 	}

// 	// // Write nicknames to CSV file
// 	// if err := writeNicknamesToCSV("nicknames.csv", nicknames); err != nil {
// 	// 	fmt.Println("Error writing nicknames to CSV:", err)
// 	// }
// }

// func writeNicknamesToCSV(filename string, nicknames []string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// Writing header if needed
// 	// writer.Write([]string{"Username"})

// 	for _, nickname := range nicknames {
// 		if err := writer.Write([]string{nickname}); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // package main

// // import (
// // 	"fmt"
// // 	"io/ioutil"
// // 	"net/url"

// // 	"github.com/dghubble/oauth1"
// // )

// // func main() {
// // 	// Replace with your actual Twitter API credentials
// // 	consumerKey := "D22E1Pq81PU648TWTRHSRPDpu"
// // 	consumerSecret := "jI0idplXT1i6KhKBl2k1DEum1L06mcnkdi1ZCxm19oL3DlYFDu"
// // 	accessToken := "3298684828-RBvjAGn4nGy7zUS3zHcF3HJdcHZvd9ZKaj5xX62"
// // 	accessSecret := "oWWWZ0WLpNlhXbOdXHkJZNBmTJTiLtWLElHGOoea6HMHd"

// // 	// OAuth1 config
// // 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// // 	token := oauth1.NewToken(accessToken, accessSecret)
// // 	httpClient := config.Client(oauth1.NoContext, token)

// // 	// Twitter API endpoint - user search
// // 	searchURL := "https://api.twitter.com/2/users/search"

// // 	// Parameters for the user search query
// // 	queryParams := url.Values{}
// // 	queryParams.Set("query", "lahore") // Replace USERNAME_TO_SEARCH with the username you want to search for
// // 	searchURL += "?" + queryParams.Encode()

// // 	// Make GET request
// // 	resp, err := httpClient.Get(searchURL)
// // 	if err != nil {
// // 		fmt.Println("Error making GET request:", err)
// // 		return
// // 	}
// // 	defer resp.Body.Close()

// // 	body, err := ioutil.ReadAll(resp.Body)
// // 	if err != nil {
// // 		fmt.Println("Error reading response body:", err)
// // 		return
// // 	}

// // 	// Print the user search response
// // 	fmt.Println(string(body))
// // }

// // // package main

// // // import (
// // // 	"fmt"
// // // 	"io/ioutil"

// // // 	"github.com/dghubble/oauth1"
// // // )

// // // func main() {
// // // 	// Replace with your actual Twitter API credentials
// // // 	consumerKey := "D22E1Pq81PU648TWTRHSRPDpu"
// // // 	consumerSecret := "jI0idplXT1i6KhKBl2k1DEum1L06mcnkdi1ZCxm19oL3DlYFDu"
// // // 	accessToken := "3298684828-RBvjAGn4nGy7zUS3zHcF3HJdcHZvd9ZKaj5xX62"
// // // 	accessSecret := "oWWWZ0WLpNlhXbOdXHkJZNBmTJTiLtWLElHGOoea6HMHd"

// // // 	// OAuth1 config
// // // 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// // // 	token := oauth1.NewToken(accessToken, accessSecret)
// // // 	httpClient := config.Client(oauth1.NoContext, token)

// // // 	// Twitter API endpoint
// // // 	url := "https://api.twitter.com/1.1/account/verify_credentials.json"

// // // 	// Make GET request
// // // 	resp, err := httpClient.Get(url)
// // // 	if err != nil {
// // // 		fmt.Println("Error making GET request:", err)
// // // 		return
// // // 	}
// // // 	defer resp.Body.Close()

// // // 	body, err := ioutil.ReadAll(resp.Body)
// // // 	if err != nil {
// // // 		fmt.Println("Error reading response body:", err)
// // // 		return
// // // 	}
// // // 	fmt.Println(string(body))
// // // }
