package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	// "github.com/websays-intelligence/wclients/fbgo"
	// "github.com/websays-intelligence/wclients/fbgo"
)

// Define a struct to store the extracted data

type PostIndexer struct {
	// post           fbgo.MediaResponseData
	inputChannelID string
	//
	Likes                    string
	Comments                 string
	DateCreated              time.Time
	Text                     string
	ContentID                string
	IGImageURL               string
	PostPageURL              string
	Author                   string
	ProfileID                int64
	Followers                int64
	OwnerProfileThumbnailURL string
}

// var _ winput.PostIndexer = (*PostIndexer)(nil)

// func NewPostIndexer(inputChannelID string, post fbgo.MediaResponseData) *PostIndexer {
// 	return &PostIndexer{
// 		// post:           post,
// 		inputChannelID: inputChannelID,
// 	}
// }

func convertKMtoDigits(s string) int64 {
	if strings.Contains(s, "K") {
		numberPart, _ := strconv.ParseFloat(strings.TrimSuffix(s, "K"), 64)
		return int64(numberPart * 1000)
	} else if strings.Contains(s, "M") {
		numberPart, _ := strconv.ParseFloat(strings.TrimSuffix(s, "M"), 64)
		return int64(numberPart * 1000000)
	} else {
		numberPart, _ := strconv.ParseInt(s, 10, 64)
		return numberPart
	}
}

func extractLikesAndComments(content string) (int64, int64, error) {
	parts := strings.Split(content, " - ")
	if len(parts) < 2 {
		return 0, 0, fmt.Errorf("failed to extract data from the content string")
	}

	likesComments := parts[0]
	dateAndTextStr := parts[1]

	fmt.Println(dateAndTextStr)

	likesCommentsParts := strings.Split(likesComments, ", ")
	if len(likesCommentsParts) < 2 {
		return 0, 0, fmt.Errorf("failed to extract likes and comments")
	}

	likes := convertKMtoDigits(strings.TrimSpace(strings.TrimSuffix(likesCommentsParts[0], " likes")))
	comments := convertKMtoDigits(strings.TrimSpace(strings.TrimSuffix(likesCommentsParts[1], " comments")))

	return likes, comments, nil
}

func extractDateAndText(dateAndText string) (time.Time, string, error) {
	dateAndTextParts := strings.SplitN(dateAndText, ": \"", 2)
	if len(dateAndTextParts) < 2 {
		return time.Time{}, "", fmt.Errorf("failed to extract date and text")
	}

	dateCreated := dateAndTextParts[0]
	text := strings.TrimSuffix(dateAndTextParts[1], "\" ")

	// Extract the "Author" and "Date Created" from the "Date Created" string
	authorAndDateCreated := strings.Split(dateCreated, " on ")
	if len(authorAndDateCreated) == 2 {
		// author := authorAndDateCreated[0]
		dateCreatedStr := authorAndDateCreated[1]

		// Parse the "Date Created" to a time.Time value
		dateCreatedTime, err := time.Parse("January 2, 2006", dateCreatedStr)
		if err != nil {
			log.Fatal("Failed to parse date: ", err)
		}

		// Convert the parsed time to UTC
		dateCreatedUTC := dateCreatedTime.UTC()
		return dateCreatedUTC, text, nil
	}

	return time.Time{}, text, nil
}

func extractContentID(doc *goquery.Document) (string, error) {
	metaContentID := doc.Find("meta[property='al:ios:url']").First()
	if metaContentID.Length() == 0 {
		return "", fmt.Errorf("no <meta> tag with property='al:ios:url' found")
	}

	contentID := metaContentID.AttrOr("content", "")
	contentIDParts := strings.Split(contentID, "id=")
	if len(contentIDParts) != 2 {
		return "", fmt.Errorf("failed to extract content ID")
	}

	contentIDValue := contentIDParts[1]

	// Retrieve the current date and time
	currentTime := time.Now()
	// Format it as a string without separators
	dateTimeStr := currentTime.Format("200601021504")

	// Combine the ContentID with the formatted date and time
	contentIDWithTime := contentIDValue + dateTimeStr

	return contentIDWithTime, nil
}

func extractIGImageURL(doc *goquery.Document) (string, error) {
	metaIGImageURL := doc.Find("meta[name='twitter:image']").First()
	if metaIGImageURL.Length() == 0 {
		return "", fmt.Errorf("no <meta> tag with name='twitter:image' found")
	}

	return metaIGImageURL.AttrOr("content", ""), nil
}

func extractAuthor(dateCreated string) (string, error) {
	author := strings.Split(dateCreated, " on ")
	if len(author) != 2 {
		return "", fmt.Errorf("failed to extract author and date")
	}

	return author[0], nil
}

func extractPostImageURL(doc *goquery.Document) (string, error) {
	metaOGImage := doc.Find("meta[property='og:image']").First()
	if metaOGImage.Length() == 0 {
		return "", fmt.Errorf("no <meta> tag with property='og:image' found")
	}

	return metaOGImage.AttrOr("content", ""), nil
}

func extractHeadData(resp *http.Response, postPageURL string) (*PostIndexer, error) {
	// var headData PostIndexer

	// Parse the HTML content using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Invalid Page: %s", err.Error())
	}

	meta := doc.Find("meta[property='og:description']").First()
	if meta.Length() == 0 {
		return nil, fmt.Errorf("no <meta> tag with property='og:description' found")
	}

	content := meta.AttrOr("content", "")

	likes, comments, err := extractLikesAndComments(content)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(content, " - ")
	dateCreated, text, dateAndText := time.Time{}, "", ""

	if len(parts) >= 2 {
		dateAndText = parts[1]
		dateCreated, text, err = extractDateAndText(dateAndText)
		if err != nil {
			return nil, err
		}
	}

	// ok above

	contentID, err := extractContentID(doc)
	if err != nil {
		return nil, err
	}

	postImageURL, err := extractPostImageURL(doc) // the image thumbnail for post
	if err != nil {
		return nil, err
	}

	author, err := extractAuthor(dateAndText)
	if err != nil {
		return nil, err
	}

	// Get the owner's image URL from owner page using api call
	// ownerImageURL, err := extractOwnerImageURL(doc)
	// if err != nil {
	// 	return nil, err
	// }

	authorPageURL := "https://www.instagram.com/" + author
	followers, ownerProfileImageURL, err := extractFollowers(authorPageURL)
	if err != nil {
		log.Fatal(err)
	}

	// PostIndexer.Followers = followers

	// Print the extracted owner image URL
	fmt.Printf("Owner Image URL: %s\n", ownerProfileImageURL)

	return &PostIndexer{
		Likes:       strconv.FormatInt(likes, 10),
		Comments:    strconv.FormatInt(comments, 10),
		DateCreated: dateCreated,
		Text:        text,
		ContentID:   contentID,
		IGImageURL:  postImageURL,
		PostPageURL: postPageURL,
		Author:      author,
		// ProfileID:                profileID,
		Followers:                followers, // Placeholder followers count
		OwnerProfileThumbnailURL: ownerProfileImageURL,
	}, nil
}

// return nil, nil

// }

func extractFollowers(authorPageURL string) (int64, string, error) {
	resp, err := http.Get(authorPageURL)
	if err != nil {
		return 0, "", fmt.Errorf("failed to make HTTP request to the author's page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, "", fmt.Errorf("failed to fetch the author's page. Status code: %d", resp.StatusCode)
	}

	authorDoc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, "", fmt.Errorf("invalid author page: %w", err)
	}

	metaAuthorDescription := authorDoc.Find("meta[property='og:description']").First()
	if metaAuthorDescription.Length() == 0 {
		return 0, "", fmt.Errorf("no <meta> tag with property='og:description' found in the author page")
	}

	authorDescription := metaAuthorDescription.AttrOr("content", "")

	// Extract the followers count from the content string
	// followersPart := strings.Split(authorDescription, " ")[0]
	// followers, err := strconv.ParseInt(followersPart, 10, 64)
	// followers2 := convertKMtoDigits(strings.TrimSuffix(followersPart[0], "M"))

	followersPart := strings.Split(authorDescription, " ")[0]
	var multiplier int64

	switch {
	case strings.Contains(followersPart, "K"):
		multiplier = 1000
	case strings.Contains(followersPart, "M"):
		multiplier = 1000000
	default:
		multiplier = 1
	}

	// Remove 'K' or 'M' from the followers part
	followersPart = strings.TrimSuffix(strings.TrimSuffix(followersPart, "K"), "M")

	// Parse the followers count and multiply by the appropriate factor
	followers, err := strconv.ParseInt(followersPart, 10, 64)
	if err != nil {
		return 0, "", err
	}

	if err != nil {
		log.Fatal("Failed to parse followers count: ", err)
	}

	f1 := followers * multiplier

	// followersIndex := strings.Index(authorDescription, " Followers")
	// if followersIndex == -1 {
	// 	return 0, fmt.Errorf("failed to find ' Followers' in the author description")
	// }

	// followersStr := authorDescription[:followersIndex]
	// followers, err := strconv.ParseInt(followersStr, 10, 64)
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to parse followers count: %w", err)
	// }

	metaOGImage := authorDoc.Find("meta[property='og:image']").First()

	// Check if the <meta> tag with the specified property was found
	ownerImageURL := ""
	if metaOGImage.Length() > 0 {
		ownerImageURL = metaOGImage.AttrOr("content", "")
	}

	// Print the extracted followers count
	fmt.Printf("Followers: %d\n", followers)

	return f1, ownerImageURL, nil
}

func main() {
	resp, err := http.Get("https://www.instagram.com/p/CzPUiEwstRB/?hl=en")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	PostPageURL := "https://www.instagram.com/p/CzPUiEwstRB/?hl=en"
	headData, err := extractHeadData(resp, PostPageURL)
	if err != nil {
		log.Fatal(err)
	}

	// authorPageURL := "https://www.instagram.com/" + headData.Author
	// followers, err := extractFollowers(authorPageURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// headData.Followers = followers

	// Print the final extracted data
	fmt.Printf("Final Extracted Data:\n%+v\n", headData)
}

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// // Define a struct to store the extracted data
// type HeadData struct {
// 	Likes       string
// 	Comments    string
// 	DateCreated string
// 	Text        string
// }

// func main() {
// 	var url1 *url.URL
// 	fmt.Println(url1)

// 	resp, err := http.Get("https://www.instagram.com/p/CzPUiEwstRB/?hl=en")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalf("Failed to fetch the page. Status code: %d", resp.StatusCode)
// 	}

// 	// Parse the HTML content using goquery
// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Invalid Page: %s", err.Error())
// 	}

// 	// Find the <meta> tag with property="og:description"
// 	meta := doc.Find("meta[property='og:description']").First()

// 	// Check if the <meta> tag with the specified property was found
// 	if meta.Length() > 0 {
// 		content := meta.AttrOr("content", "")

// 		// Split the content string using known delimiters
// 		parts := strings.Split(content, " likes, ")
// 		if len(parts) >= 2 {
// 			likesAndComments := parts[0]
// 			dateAndText := parts[1]

// 			likesCommentsParts := strings.Split(likesAndComments, " comments - ")
// 			if len(likesCommentsParts) >= 2 {
// 				likes := likesCommentsParts[0]
// 				comments := likesCommentsParts[1]

// 				dateAndTextParts := strings.Split(dateAndText, ": \"")
// 				if len(dateAndTextParts) >= 2 {
// 					dateCreated := dateAndTextParts[0]
// 					text := dateAndTextParts[1]

// 					// Create a struct to store the extracted data
// 					headData := HeadData{
// 						Likes:       likes,
// 						Comments:    comments,
// 						DateCreated: dateCreated,
// 						Text:        text,
// 					}

// 					// Print the extracted data
// 					fmt.Printf("Likes: %s\n", headData.Likes)
// 					fmt.Printf("Comments: %s\n", headData.Comments)
// 					fmt.Printf("Date Created: %s\n", headData.DateCreated)
// 					fmt.Printf("Text: %s\n", headData.Text)
// 				} else {
// 					log.Fatal("Failed to extract date and text.")
// 				}
// 			} else {
// 				log.Fatal("Failed to extract likes and comments.")
// 			}
// 		} else {
// 			log.Fatal("Failed to extract data from the content string.")
// 		}
// 	} else {
// 		log.Fatal("Meta tag not found.")
// 	}
// }

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
