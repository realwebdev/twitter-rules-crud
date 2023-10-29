package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "https://api.twitter.com/2/tweets/search/stream/rules"
	method := "GET"
	bearerToken := "AAAAAAAAAAAAAAAAAAAAAHz%2BpAEAAAAATBeQ85JC3V9jjImUubyVSRPHWvY%3DppIMyd8Yo4315kOdkfaca8YMxOGGTuksxES0sQCENTnDjMECFb"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+bearerToken)
	req.Header.Set("User-Agent", "v2FilteredStreamRules")

	time.Sleep(3 * time.Second)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response struct {
		Rules []struct {
			ID     string `json:"id"`
			Value  string `json:"value"`
			Tag    string `json:"tag"`
			Create string `json:"created_at"`
		} `json:"data"`
	}

	// Unmarshal the JSON response into the struct
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err)
		return
	}

	// Serialize the struct back to JSON
	prettyJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {	
		fmt.Println(err)
		return
	}

	// Write the json to a file
	if err := os.WriteFile("rules.json", prettyJSON)

}
