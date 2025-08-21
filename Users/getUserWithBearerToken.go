package Users

import (
	"fmt"
	"net/http"
)

func getUserWithBearerToken(userName string) (*http.Response, error) {
	url := fmt.Sprintf(baseURL+"/users/by/%v", userName)
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
