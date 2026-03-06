package gomiak

import "fmt"

// APIError represents an error response from the Infomaniak API.
type APIError struct {
	HTTPStatus  int    `json:"-"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (e *APIError) Error() string {
	if e.Description != "" {
		return fmt.Sprintf("infomaniak: %d %s: %s", e.HTTPStatus, e.Code, e.Description)
	}
	if e.Code != "" {
		return fmt.Sprintf("infomaniak: %d %s", e.HTTPStatus, e.Code)
	}
	return fmt.Sprintf("infomaniak: %d", e.HTTPStatus)
}

// errorResponse is used to parse error payloads from the API.
type errorResponse struct {
	Result string    `json:"result"`
	Error  *APIError `json:"error"`
}
