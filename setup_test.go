package accountviewnet_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	accountviewnet "github.com/omniboost/go-accountview.net"
	"golang.org/x/oauth2"
)

var (
	client *accountviewnet.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	tokenURL := os.Getenv("TOKEN_URL")
	companyID := os.Getenv("COMPANY_ID")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	oauthConfig := accountviewnet.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	if baseURL != nil {
		oauthConfig.SetBaseURL(baseURL)
	}

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	// b, _ := json.MarshalIndent(oauthConfig, "", "  ")
	// log.Fatal(string(b))

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(oauth2.NoContext, token)

	client = accountviewnet.NewClient(httpClient, companyID)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
