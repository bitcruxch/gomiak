package gomiak

// Response is the standard Infomaniak API response envelope.
type Response[T any] struct {
	Result     string      `json:"result"`
	Data       T           `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Pagination holds paging metadata returned by list endpoints.
type Pagination struct {
	Total       int `json:"total"`
	Page        int `json:"page"`
	Pages       int `json:"pages"`
	ItemsPerPage int `json:"items_per_page"`
}
