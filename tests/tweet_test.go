package tests

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	config := oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "oob",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
			AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
			AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
		},
	}

	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		log.Fatalf("Failed to get request token: %v", err)
	}
	fmt.Printf("Got request token: %s\n", requestToken)

	authorizationURL, _ := config.AuthorizationURL(requestToken)
	fmt.Printf("Please go here and authorize: %s\n", authorizationURL.String())

	fmt.Print("Paste the PIN here: ")
	reader := bufio.NewReader(os.Stdin)
	verifier, _ := reader.ReadString('\n')

	accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, verifier)
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}

	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	payload := map[string]string{"text": "Hello world!"}
	payloadBytes, _ := json.Marshal(payload)
	fmt.Printf(">>:%v\n", payloadBytes)
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bufio.NewReader(os.Stdin))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = http.NoBody


	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Request returned an error: %d %s", resp.StatusCode, resp.Status)
	}

	var jsonResponse map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&jsonResponse)
	responseFormatted, _ := json.MarshalIndent(jsonResponse, "", "  ")
	fmt.Println(string(responseFormatted))
}
