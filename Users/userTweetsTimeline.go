package Users

import (
	"fmt"
	"net/http"
	"os"
)

var (
	bearer_token = os.Getenv("BEARER_TOKEN")
	//baseURL      = "https://api.x.com/2/tweets"
	baseURL = "https://api.x.com/2"

// baseULR + "/tweets"+"%v", userid
)

func CreateURL(UserId any) string {
	//url := fmt.Sprintf("https://api.x.com/2/tweets/%v", UserId)
	//url := fmt.Sprintf(baseURL+"/%v", UserId)
	url := fmt.Sprintf(baseURL+"/tweets/%v", UserId)
	return url
}

func GetParams() []string {
	// Tweet fields are adjustable.
	// Options include:
	// attachments, author_id, context_annotations,
	// conversation_id, created_at, entities, geo, id,
	// in_reply_to_user_id, lang, non_public_metrics, organic_metrics,
	// possibly_sensitive, promoted_metrics, public_metrics, referenced_tweets,
	// source, text, and withheld

	fields := []string{"tweet.fields", "created_at"}
	return fields
}

func BearerOAUTH(httpRequest *http.Request) {
	httpRequest.Header.Add("Authorization", "Bearer "+bearer_token)
	httpRequest.Header.Add("User-Agent", "v2UserTweetsGo")

}

func userTweetTimeline(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

	}
	BearerOAUTH(req)
	client := http.Client{}

	res, err := client.Do(req)

	return res
}
