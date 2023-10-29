package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	bearerToken := os.Getenv("TWITTER_BEARER_TOKEN")
	if bearerToken == "" {
		fmt.Println("Twitter Bearer Token not found in environment variables.")
		os.Exit(1)
	}
	// Step 1: Fetch the existing rules and extract their IDs
	existingRuleIDs, err := fetchExistingRuleIDs(bearerToken)
	if err != nil {
		fmt.Println("Error fetching existing rule IDs:", err)
		os.Exit(1)
		return
	}

	// Step 2: Delete the rules using the extracted IDs
	if err := deleteRules(existingRuleIDs, bearerToken); err != nil {
		fmt.Println("Error deleting rules:", err)
		os.Exit(1)
		return
	}

	fmt.Println("Rules deleted successfully.")
}

// Step 1: Fetch the existing rules and extract their IDs
func fetchExistingRuleIDs(bearerToken string) ([]string, error) {
	url := "https://api.twitter.com/2/tweets/search/stream/rules"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+bearerToken)
	time.Sleep(3 * time.Second)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response to extract rule IDs
	var response struct {
		Rules []struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	ruleIDs := make([]string, len(response.Rules))
	for i, rule := range response.Rules {
		ruleIDs[i] = rule.ID
	}

	return ruleIDs, nil
}

// Step 2: Delete rules using the extracted IDs
func deleteRules(ruleIDs []string, bearerToken string) error {
	url := "https://api.twitter.com/2/tweets/search/stream/rules"
	method := "POST"

	payload := fmt.Sprintf(`{
		"delete": {
			"ids": %s
		}
	}`, toJSON(ruleIDs))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerToken)

	time.Sleep(3 * time.Second)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response:", string(body))
	fmt.Println("Response status:", res.Status)
	return nil
}

// Helper function to convert a slice to JSON array
func toJSON(data interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}

func fetchRulesToSaveInFile(bearerToken string) {
	url := "https://api.twitter.com/2/tweets/search/stream/rules"
	method := "GET"
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
	fmt.Println(string(body))

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
	if err := os.WriteFile("rules.json", prettyJSON, 0644); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response saved to response.json")

}
