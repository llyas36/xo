package tweet

import (
	"fmt"
	"net/http"
	"os"
)

//// find out who retweeted a specific Tweet
//
//

var (
	bearer_token = os.Getenv("BEARER_TOKEN")
)

func CreateURL(id string) string {
	url := fmt.Sprintf("https://api.x.com/2/tweets/%s/retweeted_by", id)
	return url
}

func BearerOAUTH(httpRequest *http.Request) {
	httpRequest.Header.Add("Authorization", "Bearer "+bearer_token)
	httpRequest.Header.Add("User-Agent", "v2RetweetedGo")
}

func RetweetedById(id string) (*http.Response, error) {
	url := CreateURL(id)
	httpRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	BearerOAUTH(httpRequest)
	client := http.Client{}
	res, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}
