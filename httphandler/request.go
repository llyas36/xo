package httphandler

import (
	"ctw/config"
	"log"
	"net/http"
)

func CreateGetRequest(url string) (*http.Request, error) {
	httpRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//log.Fatal(err)
		return httpRequest, err
	}
	setupHeader(httpRequest)
	return httpRequest, nil
}

func setupHeader(httpReqeust *http.Request) {
	httpReqeust.Header.Set("Authorization", "Bearer "+config.BearerToken)

}

func GetResponse(httpRequest *http.Request) (*http.Response, error) {
	client := http.Client{}
	res, err := client.Do(httpRequest)
	if err != nil {
		return res, err
	}
	return res, nil
}

func UsersLookup(usernames string, user_fields string) (*http.Request, error) {
	//usernames := "elonmusk,llyas__"
	//baseURL := "https://api.x.com/2/users/by"
	subURL := "/users/by/username/"
	// user_fields := "user.fields=description,created_at" // optional

	// params := url.Values{}
	// params.Add("usernames", usernames)
	// params.Add("user.Fields", user_fields)

	//fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fullURL := config.CreateURL(usernames, user_fields, subURL)
	// data(username)
	// fields(user_fields)
	// subURL(/users/by/username/{%s}
	// parmas(&parmas)
	httpRequest, err := CreateGetRequest(fullURL)
	if err != nil {
		return httpRequest, err
	}

	return httpRequest, nil

}

func GetTweetsById(ids string, tweet_fields string) *http.Response {
	subURL := "/tweets"
	// data(username)
	// fields(tweet_field)
	// subURL
	//
	url := config.CreateURL("", "", subURL)
	// bearer_token := config.ReturnEnv()
	req, err := CreateGetRequest(url)
	if err != nil {
		log.Fatal(err)
	}

	res, err := GetResponse(req)
	if err != nil {
		log.Fatal(err)
	}

	return res

}
