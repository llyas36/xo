package Space

import "net/http"

var (
	search_url = "https://api.x.com/2/space/search"
)

func SearchSpace(search_term string) (*http.Response, error) {

	// query_parmas := map[string]any{
	// 	"query":        search_term,
	// 	"space.fields": []string{"title, created_at"},
	// 	"expansions":   "creator_id",
	// }
	// optional add params...
	// params := url.Values{}
	// params.Add("user.Fields", query_params[])
	// url := search_url + params.Encode()

	httpReqeust, err := http.NewRequest("GET", search_url, nil)
	if err != nil {
		return nil, err
	}
	HandleHeader(httpReqeust)

	client := http.Client{}
	res, err := client.Do(httpReqeust)
	if err != nil {
		return nil, err
	}

	return res, nil
}
