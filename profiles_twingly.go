package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Data struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
}

type JSONData struct {
	Data []Data `json:"data"`
}

func main() {
	// Read the JSON file
	file, _ := ioutil.ReadFile("profiles_twingly.json")

	// Unmarshal the JSON file into a variable
	var jsonData JSONData
	_ = json.Unmarshal([]byte(file), &jsonData)

	// Create a map to store unique tags
	uniqueTags := make(map[string]bool)

	// Iterate over the data
	for _, d := range jsonData.Data {
		// Split the tag and take the number part
		tagParts := strings.Split(d.Tag, "#")
		if len(tagParts) > 0 {
			// Check if the tag is already in the map
			if _, exists := uniqueTags[tagParts[0]]; !exists {
				// If not, add it to the map
				uniqueTags[tagParts[0]] = true
			}
		}
	}

	// Create a slice to store the unique tags
	var tags []string
	for tag := range uniqueTags {
		tags = append(tags, tag)
	}

	// Marshal the tags into a JSON string
	newJSON, _ := json.Marshal(tags)

	// Write the new JSON string to a new file
	_ = ioutil.WriteFile("profiles_twingly_output.json", newJSON, 0644)
}
