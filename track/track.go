package track

import "net/http"

const (
	root = "suivi"
	version = "v1"
)

type Client struct {
	APIKey string
	httpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

func New(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		httpClient: http.DefaultClient,
	}
}

func (cl *Client) WithClient(httpClient *http.Client) {
	cl.httpClient = httpClient
}

func (cl *Client) auth(req *http.Request) {
	req.Header.Set("X-Okapi-Key", cl.APIKey)
}
