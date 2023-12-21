package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/joho/godotenv"
// )

// type RuleData struct {
// 	Value string `json:"value"`
// 	Tag   string `json:"tag"`
// }

// func main() {
// 	// Load environment variables from .env
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Error loading .env file")
// 		return
// 	}
// 	bearerTokentoFetchRules := os.Getenv("TWITTER_BEARER_TOKEN_TO_FETCH_RULES")
// 	if bearerTokentoFetchRules == "" {
// 		fmt.Println("Twitter Bearer Token not found in environment variables.")
// 		os.Exit(1)
// 	}

// 	bearerTokentoCreateRules := os.Getenv("TWITTER_BEARER_TOKEN_TO_CREATE_RULES")
// 	if bearerTokentoCreateRules == "" {
// 		fmt.Println("Twitter Bearer Token not found in environment variables.")
// 		os.Exit(1)
// 	}
// 	// Step 1: Fetch the existing rules and extract their IDs
// 	existingRuleIDs, err := fetchExistingRule(bearerTokentoFetchRules)
// 	if err != nil {
// 		fmt.Println("Error fetching existing rule IDs:", err)
// 		os.Exit(1)
// 		return
// 	}

// 	// Step 2: Delete the rules using the extracted IDs
// 	if err := createRules(existingRuleIDs, bearerTokentoCreateRules); err != nil {
// 		fmt.Println("Error deleting rules:", err)
// 		os.Exit(1)
// 		return
// 	}

// 	fmt.Println("Rules deleted successfully.")
// }

// func fetchExistingRule(bearerToken string) ([]RuleData, error) {
// 	url := "https://api.twitter.com/2/tweets/search/stream/rules"
// 	method := "GET"

// 	client := &http.Client{}

// 	req, err := http.NewRequest(method, url, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", "Bearer "+bearerToken)
// 	time.Sleep(3 * time.Second)

// 	res, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Parse the JSON response to extract rule IDs
// 	var response struct {
// 		Rules []RuleData `json:"data"`
// 	}

// 	if err := json.Unmarshal(body, &response); err != nil {
// 		return nil, err
// 	}

// 	// Unmarshal the JSON response into the struct
// 	if err := json.Unmarshal(body, &response); err != nil {
// 		fmt.Println(err)
// 	}

// 	// Serialize the struct back to JSON
// 	prettyJSON, err := json.MarshalIndent(response, "", "  ")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Write the json to a file
// 	if err := os.WriteFile("rules2.json", prettyJSON, 0644); err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Response saved to response.json")

// 	return response.Rules, nil
// }

// func createRules(rules []RuleData, bearerToken string) error {
// 	url := "https://api.twitter.com/2/tweets/search/stream/rules"
// 	method := "POST"

// 	// Create a slice of rules to be added
// 	var rulesToAdd []map[string]interface{}
// 	for _, rule := range rules {
// 		rulesToAdd = append(rulesToAdd, map[string]interface{}{
// 			"value": rule.Value,
// 			"tag":   rule.Tag,
// 		})
// 	}

// 	// Create the payload with the rules to be added
// 	payload := map[string][]map[string]interface{}{"add": rulesToAdd}
// 	payloadJSON, err := json.Marshal(payload)
// 	if err != nil {
// 		return err
// 	}

// 	client := &http.Client{}

// 	req, err := http.NewRequest(method, url, bytes.NewReader(payloadJSON))
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", "Bearer "+bearerToken)

// 	time.Sleep(3 * time.Second)

// 	res, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Response:", string(body))
// 	fmt.Println("Response status:", res.Status)

// 	return nil
// }

type StreamFs struct {
	Data struct {
		Attachments struct {
			MediaKeys []string `json:"media_keys"`
		} `json:"attachments"`
		AuthorID       string `json:"author_id"`
		ConversationID string `json:"conversation_id"`
		CreatedAt      string `json:"created_at"`
		EditControls   struct {
			EditsRemaining int    `json:"edits_remaining"`
			IsEditEligible bool   `json:"is_edit_eligible"`
			EditableUntil  string `json:"editable_until"`
		} `json:"edit_controls"`
		EditHistoryTweetIDs []string `json:"edit_history_tweet_ids"`
		Entities            struct {
			Hashtags []struct {
				Start int    `json:"start"`
				End   int    `json:"end"`
				Tag   string `json:"tag"`
			} `json:"hashtags"`
			Mentions []struct {
				Start    int    `json:"start"`
				End      int    `json:"end"`
				Username string `json:"username"`
				ID       string `json:"id"`
			} `json:"mentions"`
			URLs []struct {
				Start       int     `json:"start"`
				End         int     `json:"end"`
				URL         string  `json:"url"`
				ExpandedURL string  `json:"expanded_url"`
				DisplayURL  string  `json:"display_url"`
				Images      []Image `json:"images"`
				Status      int     `json:"status"`
				Title       string  `json:"title"`
				Description string  `json:"description"`
				UnwoundURL  string  `json:"unwound_url"`
				MediaKey    string  `json:"media_key"`
			} `json:"urls"`
		} `json:"entities"`
		Geo               struct{} `json:"geo"`
		ID                string   `json:"id"`
		InReplyToUserID   string   `json:"in_reply_to_user_id"`
		Lang              string   `json:"lang"`
		PossiblySensitive bool     `json:"possibly_sensitive"`
		PublicMetrics     struct {
			RetweetCount    int `json:"retweet_count"`
			ReplyCount      int `json:"reply_count"`
			LikeCount       int `json:"like_count"`
			QuoteCount      int `json:"quote_count"`
			BookmarkCount   int `json:"bookmark_count"`
			ImpressionCount int `json:"impression_count"`
		} `json:"public_metrics"`
		ReferencedTweets []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"referenced_tweets"`
		ReplySettings string `json:"reply_settings"`
		Text          string `json:"text"`
	} `json:"data"`
	Includes struct {
		Places []struct {
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			FullName    string `json:"full_name"`
			Name        string `json:"name"`
			PlaceType   string `json:"place_type"`
			ID          string `json:"id"`
			Geo         struct {
				Type       string    `json:"type"`
				Bbox       []float64 `json:"bbox"`
				Properties struct {
				} `json:"properties"`
			} `json:"geo"`
		} `json:"places"`
		Media []struct {
			Height        int      `json:"height"`
			MediaKey      string   `json:"media_key"`
			PublicMetrics struct{} `json:"public_metrics"`
			Type          string   `json:"type"`
			URL           string   `json:"url"`
			Width         int      `json:"width"`
		} `json:"media"`
		Users []struct {
			CreatedAt       string `json:"created_at"`
			Description     string `json:"description"`
			ID              string `json:"id"`
			Location        string `json:"location"`
			Name            string `json:"name"`
			ProfileImageURL string `json:"profile_image_url"`
			Protected       bool   `json:"protected"`
			PublicMetrics   struct {
				FollowersCount int `json:"followers_count"`
				FollowingCount int `json:"following_count"`
				TweetCount     int `json:"tweet_count"`
				ListedCount    int `json:"listed_count"`
			} `json:"public_metrics"`
			Username     string `json:"username"`
			Verified     bool   `json:"verified"`
			VerifiedType string `json:"verified_type"`
			Entities     struct {
				URL struct {
					Urls []struct {
						Start       int    `json:"start"`
						End         int    `json:"end"`
						URL         string `json:"url"`
						ExpandedURL string `json:"expanded_url"`
						DisplayURL  string `json:"display_url"`
					} `json:"urls"`
				} `json:"url"`
				Description struct {
					Urls []struct {
						Start       int    `json:"start"`
						End         int    `json:"end"`
						URL         string `json:"url"`
						ExpandedURL string `json:"expanded_url"`
						DisplayURL  string `json:"display_url"`
					} `json:"urls"`
					Mentions []struct {
						Start    int    `json:"start"`
						End      int    `json:"end"`
						Username string `json:"username"`
					} `json:"mentions"`
				} `json:"description"`
			} `json:"entities,omitempty"`
		} `json:"users"`
		Tweets []struct {
			Attachments struct {
				MediaKeys []string `json:"media_keys"`
			} `json:"attachments"`
			AuthorID       string `json:"author_id"`
			ConversationID string `json:"conversation_id"`
			CreatedAt      string `json:"created_at"`
			EditControls   struct {
				EditsRemaining int    `json:"edits_remaining"`
				IsEditEligible bool   `json:"is_edit_eligible"`
				EditableUntil  string `json:"editable_until"`
			} `json:"edit_controls"`
			EditHistoryTweetIDs []string `json:"edit_history_tweet_ids"`
			Entities            struct {
				Annotations []struct {
					Start          int     `json:"start"`
					End            int     `json:"end"`
					Probability    float64 `json:"probability"`
					Type           string  `json:"type"`
					NormalizedText string  `json:"normalized_text"`
				} `json:"annotations"`
				Hashtags []struct {
					Start int    `json:"start"`
					End   int    `json:"end"`
					Tag   string `json:"tag"`
				} `json:"hashtags"`
				Mentions []struct {
					Start    int    `json:"start"`
					End      int    `json:"end"`
					Username string `json:"username"`
					ID       string `json:"id"`
				} `json:"mentions"`
				Urls []struct {
					Start       int    `json:"start"`
					End         int    `json:"end"`
					URL         string `json:"url"`
					ExpandedURL string `json:"expanded_url"`
					DisplayURL  string `json:"display_url"`
					MediaKey    string `json:"media_key"`
				} `json:"urls"`
			} `json:"entities"`
			Geo               struct{} `json:"geo"`
			ID                string   `json:"id"`
			InReplyToUserID   string   `json:"in_reply_to_user_id"`
			Lang              string   `json:"lang"`
			PossiblySensitive bool     `json:"possibly_sensitive"`
			PublicMetrics     struct {
				RetweetCount    int `json:"retweet_count"`
				ReplyCount      int `json:"reply_count"`
				LikeCount       int `json:"like_count"`
				QuoteCount      int `json:"quote_count"`
				BookmarkCount   int `json:"bookmark_count"`
				ImpressionCount int `json:"impression_count"`
			} `json:"public_metrics"`
			ReferencedTweets []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"referenced_tweets"`
			ReplySettings string `json:"reply_settings"`
			Text          string `json:"text"`
			Attachments0  struct {
				PollIds []string `json:"poll_ids"`
			} `json:"attachments0,omitempty"`
			ContextAnnotations []struct {
				Domain struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
				} `json:"domain"`
				Entity struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
				} `json:"entity,omitempty"`
				Entity0 struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"entity0,omitempty"`
			} `json:"context_annotations,omitempty"`
			Entities0 struct {
				Annotations []struct {
					Start          int     `json:"start"`
					End            int     `json:"end"`
					Probability    float64 `json:"probability"`
					Type           string  `json:"type"`
					NormalizedText string  `json:"normalized_text"`
				} `json:"annotations"`
				Hashtags []struct {
					Start int    `json:"start"`
					End   int    `json:"end"`
					Tag   string `json:"tag"`
				} `json:"hashtags"`
				Mentions []struct {
					Start    int    `json:"start"`
					End      int    `json:"end"`
					Username string `json:"username"`
					ID       string `json:"id"`
				} `json:"mentions"`
			} `json:"entities0,omitempty"`
		} `json:"tweets"`
	} `json:"includes"`
	Errors []struct {
		Value        string `json:"value"`
		Detail       string `json:"detail"`
		Title        string `json:"title"`
		ResourceType string `json:"resource_type"`
		Parameter    string `json:"parameter"`
		ResourceID   string `json:"resource_id"`
	} `json:"errors"`
	MatchingRules []struct {
		ID  string `json:"id"`
		Tag string `json:"tag"`
	} `json:"matching_rules"`
}

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
