package Users

import (
	"fmt"
	"net/http"
)

func handleURL(tweet_fields string, ids string) string {
	url := fmt.Sprintf("https://api.x.com/2/tweets?%s&%s", ids, tweet_fields)

	return url

}

func getTweetWithBearer(tweet_fields string, ids string) (*http.Response, error) {
	url := handleURL(tweet_fields, ids)
	req, err := http.NewRequest("GET", url, nil)
	BearerOAUTH(req)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
