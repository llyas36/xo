package config

import (
	"fmt"
	"net/url"
	"os"
)

var (
	// the tokens for the authorization
	BearerToken string = ReturnEnv()
	UserAgent   string = "CERN-LineMode/2.15 libwww/2.17b3"
)

func Init(bearerToken string) {
	BearerToken = bearerToken
}

func SetUserAgent(userAgent string) {
	UserAgent = userAgent
}

func ReturnEnv() string {
	os.Setenv("BEARER_TOKEN", "AAAAAAAAAAAAAAAAAAAAAIPe2wEAAAAAxT%2F09a1KkzOnbfmHxkx7Fpl%2F5Q0%3Dx8Mg64KS9FUebdTJM4VZOkwTis4SlKBkTZDOcPveEA6h3bWpmJ")
	return os.Getenv("BEARER_TOKEN")
}

func CreateURL(data string, fields string, subURL string) string {
	baseURL := "https://api.x.com/2/"
	params := url.Values{}
	params.Add("data", data)
	params.Add("User.fields", fields)

	fullURL := fmt.Sprintf("%s%s?%s", baseURL, subURL, params.Encode())
	// https://api.x.com/2/users/by

	return fullURL
}

// data(username)
// fields(user_fields)
// subURL(/users/by/username/{%s}
// parmas(&parmas)
/// CreateURL(data, fields, params)
