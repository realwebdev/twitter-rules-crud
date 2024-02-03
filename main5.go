package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// func main() {
// 	url := "https://api.twitter.com/2/users/search"
// 	method := "GET"

// 	// Set the query parameters
// 	queryParams := map[string]string{
// 		"query":        "oncologist",
// 		"user.fields":  "created_at,description,entities,id,location,most_recent_tweet_id,name,pinned_tweet_id,profile_image_url,protected,public_metrics,receives_your_dm,subscription_type,url,username,verified,verified_type,withheld",
// 		"tweet.fields": "attachments,author_id,card_uri,context_annotations,conversation_id,created_at,edit_controls,edit_history_tweet_ids,entities,geo,id,in_reply_to_user_id,lang,non_public_metrics,note_tweet,organic_metrics,possibly_sensitive,promoted_metrics,public_metrics,referenced_tweets,reply_settings,source,text,withheld",
// 		"expansions":   "pinned_tweet_id,most_recent_tweet_id",
// 	}

// 	// Construct the URL with query parameters
// 	query := make([]string, 0, len(queryParams))
// 	for key, value := range queryParams {
// 		query = append(query, fmt.Sprintf("%s=%s", key, value))
// 	}
// 	url += "?" + strings.Join(query, "&")

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, nil)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Replace the following OAuth header with your actual OAuth header
// 	req.Header.Add("Authorization", "OAuth oauth_consumer_key=\"D22E1Pq81PU648TWTRHSRPDpu\",oauth_token=\"3298684828-RBvjAGn4nGy7zUS3zHcF3HJdcHZvd9ZKaj5xX62\",oauth_signature_method=\"HMAC-SHA1\",oauth_timestamp=\"1706267507\",oauth_nonce=\"iHk6bCyEyT3\",oauth_signature=\"8SpU0KwOEZHYziSEkAjO8ftSd84%3D\"")
// 	req.Header.Add("Cookie", "guest_id=v1%3A168996171338363383")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(body))
// }
