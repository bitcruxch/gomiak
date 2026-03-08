package newsletter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitcruxch/gomiak"
)

// Template represents a newsletter template.
type Template struct {
	ID        int    `json:"id"`
	DomainID  int    `json:"domain_id"`
	Name      string `json:"name"`
	Content   any    `json:"content"`
	Thumbnail string `json:"thumbnail"`
}

// TemplateService handles template-related API calls.
type TemplateService struct {
	s *Service
}

// List returns all templates.
func (ts *TemplateService) List(ctx context.Context, opts *gomiak.ListOptions) ([]Template, *gomiak.Pagination, error) {
	return gomiak.Do[[]Template](ctx, ts.s.client, http.MethodGet, ts.s.basePath+"/templates", opts.Values(), nil)
}

// GetHTML returns the HTML content of a template.
func (ts *TemplateService) GetHTML(ctx context.Context, id int) (string, error) {
	// This endpoint returns HTML directly in the data field.
	data, _, err := gomiak.Do[string](ctx, ts.s.client, http.MethodGet, fmt.Sprintf("%s/templates/%d/html", ts.s.basePath, id), nil, nil)
	return data, err
}

// UpdateThumbnail regenerates the thumbnail for a template.
func (ts *TemplateService) UpdateThumbnail(ctx context.Context, id int) error {
	return gomiak.DoEmpty(ctx, ts.s.client, http.MethodPut, fmt.Sprintf("%s/templates/%d/update-thumbnails", ts.s.basePath, id), nil, nil)
}
