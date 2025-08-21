package tweet

import "net/http"

var (
	search_url = "https:api.x.com/2/tweets/counts/recent"
)

func CountTweet() (*http.Response, error) {
	httpRequest, err := http.NewRequest("GET", search_url, nil)
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
