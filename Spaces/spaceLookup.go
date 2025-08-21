package Space

import (
	"net/http"
	"os"
)

var (
	bearer_token = os.Getenv("BEARER_TOKEN")
	space_url    = "https://api.x.com/2/spaces"
	//optional params...
	query_parmas = []string{
		"ids",
		"SPACE_ID",
		"space.fields",
		"title",
		"created_at",
		"expansions",
		"creator_id",
	}
)

func HandleHeader(httpRequest *http.Request) {
	httpRequest.Header.Add("Authrization", "Bearer "+bearer_token)
	httpRequest.Header.Add("User-Agent", "v2SpaceLookupGo")

}

func SpaceLookup(user_fields []string) (*http.Response, error) {
	// optional add params...
	// params := url.Values{}
	// params.Add("user.Fields", user_fields[0])
	// url := search_url + params.Encode()
	httpRequest, err := http.NewRequest("GET", space_url, nil)

	if err != nil {
		return nil, err
	}
	HandleHeader(httpRequest)

	client := http.Client{}
	res, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return res, nil
}
