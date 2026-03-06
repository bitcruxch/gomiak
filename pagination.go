package gomiak

import (
	"net/url"
	"strconv"
)

// ListOptions specifies common pagination and sorting parameters.
type ListOptions struct {
	Page    int
	PerPage int
	OrderBy string
	Order   string // "asc" or "desc"
}

// Values encodes ListOptions as URL query parameters.
func (o *ListOptions) Values() url.Values {
	if o == nil {
		return nil
	}
	v := url.Values{}
	if o.Page > 0 {
		v.Set("page", strconv.Itoa(o.Page))
	}
	if o.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(o.PerPage))
	}
	if o.OrderBy != "" {
		v.Set("order_by", o.OrderBy)
	}
	if o.Order != "" {
		v.Set("order", o.Order)
	}
	return v
}
