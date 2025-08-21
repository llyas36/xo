package stream

import (
	"net/http"
	"os"
)

var (
	bearer_token = os.Getenv("BEARER_TOKEN")
	url          = "https://api.x.com/2/tweets/sample/stream"
)

func BearerOAUTH(httpRequest *http.Request) {
	httpRequest.Header.Add("Authorization", "Bearer "+bearer_token)
	httpRequest.Header.Add("User-Agent", "v2SampledStreamGo")

}

func SampleStream() (*http.Response, error) {
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
