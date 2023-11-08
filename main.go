package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Define a struct to store the extracted data
type HeadData struct {
	Likes       string
	Comments    string
	DateCreated string
	Text        string
}

func main() {
	var url1 *url.URL
	fmt.Println(url1)

	resp, err := http.Get("https://www.instagram.com/p/CzPUiEwstRB/?hl=en")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch the page. Status code: %d", resp.StatusCode)
	}

	// Parse the HTML content using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Invalid Page: %s", err.Error())
	}

	// Find the <meta> tag with property="og:description"
	meta := doc.Find("meta[property='og:description']").First()

	// Check if the <meta> tag with the specified property was found
	if meta.Length() > 0 {
		content := meta.AttrOr("content", "")

		// Split the content string using known delimiters
		parts := strings.Split(content, " likes, ")
		if len(parts) >= 2 {
			likesAndComments := parts[0]
			dateAndText := parts[1]

			likesCommentsParts := strings.Split(likesAndComments, " comments - ")
			if len(likesCommentsParts) >= 2 {
				likes := likesCommentsParts[0]
				comments := likesCommentsParts[1]

				dateAndTextParts := strings.Split(dateAndText, ": \"")
				if len(dateAndTextParts) >= 2 {
					dateCreated := dateAndTextParts[0]
					text := dateAndTextParts[1]

					// Create a struct to store the extracted data
					headData := HeadData{
						Likes:       likes,
						Comments:    comments,
						DateCreated: dateCreated,
						Text:        text,
					}

					// Print the extracted data
					fmt.Printf("Likes: %s\n", headData.Likes)
					fmt.Printf("Comments: %s\n", headData.Comments)
					fmt.Printf("Date Created: %s\n", headData.DateCreated)
					fmt.Printf("Text: %s\n", headData.Text)
				} else {
					log.Fatal("Failed to extract date and text.")
				}
			} else {
				log.Fatal("Failed to extract likes and comments.")
			}
		} else {
			log.Fatal("Failed to extract data from the content string.")
		}
	} else {
		log.Fatal("Meta tag not found.")
	}
}

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"regexp"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// func main() {
// 	// var ctx context.Context
// 	var url1 *url.URL
// 	// var language string
// 	fmt.Println(url1)

// 	resp, err := http.Get("https://www.instagram.com/p/CzPUiEwstRB/?hl=en")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		// return nil, fmt.Errorf("failed to fetch the page. Status code: %d", resp.StatusCode)
// 	}

// 	// Parse the HTML content using goquery
// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		// return nil, fmt.Errorf("invalid Page: %s", err.Error())
// 	}

// 	// Find the <script> tag with type="application/ld+json" in the <head>
// 	var jsonScript string
// 	doc.Find("head script[type=\"application/ld+json\"]").Each(func(index int, script *goquery.Selection) {
// 		// Get the content of the script tag
// 		scriptContent := script.Text()
// 		jsonScript = scriptContent
// 	})

// 	// Print the extracted JavaScript code
// 	if jsonScript == "" {
// 		// return nil, fmt.Errorf("no JavaScript code found in the specified script tage")
// 	}

// 	// lnp := ln.LNData{}
// 	// err = json.Unmarshal([]byte(jsonScript), &lnp)
// 	if err != nil {
// 		// return nil, fmt.Errorf("getting error while trying to unmarshal linkedin: %s", err)
// 	}
// 	LinkedINID := ""
// 	// Find the <meta> tag with property="lnkd:url"
// 	meta := doc.Find("meta[property='lnkd:url']")

// 	// Check if the <meta> tag with the specified property was found
// 	if meta.Length() > 0 {
// 		// Extract the 'content' attribute from the <meta> tag
// 		content := meta.First().AttrOr("content", "")
// 		parts := strings.Split(content, ":activity:")
// 		if len(parts) >= 2 {
// 			LinkedINID = parts[1]
// 		}
// 	}
// 	if LinkedINID == "" {
// 		// return nil, fmt.Errorf("invalid Linkedin id:  %s", LinkedINID)
// 	}
// 	// lnp.ClippingID = LinkedINID

// 	pa := doc.Find("p.public-post-author-card__followers")

// 	// Check if the <p> element with the specified classes was found
// 	if pa.Length() > 0 {
// 		followers := pa.Text()
// 		re := regexp.MustCompile(`\d+,\d+`)
// 		matches := re.FindAllString(followers, -1)
// 		if len(matches) > 0 {
// 			numberStr := matches[0]
// 			numberStr = strings.ReplaceAll(numberStr, ",", "")
// 			// number, err := strconv.ParseInt(numberStr, 10, 64)
// 			if err == nil {
// 				// lnp.FollowersCount = number
// 			}
// 		}
// 	}

// 	article := doc.Find("article").First()
// 	if article.Length() > 0 {
// 		reactionCount := article.Find("span[data-test-id='social-actions__reaction-count']")
// 		if reactionCount.Length() > 0 {
// 			likes := reactionCount.Text()
// 			likes = strings.TrimSpace(likes)
// 			likes = strings.ReplaceAll(likes, ",", "")
// 			// number, err := strconv.ParseInt(likes, 10, 64)
// 			if err == nil {
// 				// lnp.LikesCount = int(number)
// 			}
// 		}
// 	}

// }
