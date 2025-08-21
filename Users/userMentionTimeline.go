package Users

import (
	"fmt"
	"net/http"
)

func userMentionsTimeline(tweetID any) (*http.Response, error) {
	//url := fmt.Sprintf("https://api.x.com/2/users/%v/mentions", tweetID)
	url := fmt.Sprintf(baseURL+"/users/%v/mentions", tweetID)
	req, err := http.NewRequest("GET", url, nil)
	BearerOAUTH(req)
	if err != nil {
		return nil, err

	}
	BearerOAUTH(req)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil

}
