package gomiak

import "net/http"

// Option configures a Client.
type Option func(*Client)

// WithBearerToken sets the API bearer token for authentication.
func WithBearerToken(token string) Option {
	return func(c *Client) {
		c.HTTPClient = &http.Client{
			Transport: &authTransport{
				token: token,
				base:  c.HTTPClient.Transport,
			},
		}
	}
}

// WithHTTPClient sets the underlying HTTP client (e.g., an OAuth2 client).
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = hc
	}
}

// WithBaseURL overrides the default API base URL.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.BaseURL = url
	}
}

// WithUserAgent sets a custom User-Agent header.
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.UserAgent = ua
	}
}
