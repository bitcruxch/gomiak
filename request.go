package gomiak

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Do executes an API request and decodes the typed response.
// It is exported so sub-packages (e.g., newsletter) can use it.
func Do[T any](ctx context.Context, c *Client, method, path string, query url.Values, body any) (T, *Pagination, error) {
	var zero T

	u := c.BaseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return zero, nil, fmt.Errorf("gomiak: marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, bodyReader)
	if err != nil {
		return zero, nil, fmt.Errorf("gomiak: create request: %w", err)
	}
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return zero, nil, fmt.Errorf("gomiak: execute request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, nil, fmt.Errorf("gomiak: read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		var errResp errorResponse
		if json.Unmarshal(data, &errResp) == nil && errResp.Error != nil {
			errResp.Error.HTTPStatus = resp.StatusCode
			return zero, nil, errResp.Error
		}
		return zero, nil, &APIError{HTTPStatus: resp.StatusCode}
	}

	var envelope Response[T]
	if err := json.Unmarshal(data, &envelope); err != nil {
		return zero, nil, fmt.Errorf("gomiak: decode response: %w", err)
	}

	return envelope.Data, envelope.Pagination, nil
}

// DoEmpty executes an API request that returns no meaningful data (e.g., boolean true).
func DoEmpty(ctx context.Context, c *Client, method, path string, query url.Values, body any) error {
	_, _, err := Do[json.RawMessage](ctx, c, method, path, query, body)
	return err
}
