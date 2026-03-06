// Package gomiak provides a Go client for the Infomaniak API.
package gomiak

import "net/http"

const (
	defaultBaseURL   = "https://api.infomaniak.com"
	defaultUserAgent = "gomiak/0.1"
)

// Client manages communication with the Infomaniak API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	UserAgent  string
}

// New creates a new Infomaniak API client with the given options.
func New(opts ...Option) *Client {
	c := &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    defaultBaseURL,
		UserAgent:  defaultUserAgent,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
