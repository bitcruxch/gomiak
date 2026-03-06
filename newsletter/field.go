package newsletter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitcrux/gomiak"
)

// Field represents a custom subscriber field.
type Field struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

// CreateFieldRequest is the request body for creating a field.
type CreateFieldRequest struct {
	Name string `json:"name"`
	Type string `json:"type"` // "boolean", "date_en", "date_fr", "email", "float", "number", "text"
	Slug string `json:"slug,omitempty"`
}

// UpdateFieldRequest is the request body for updating a field.
type UpdateFieldRequest struct {
	Name string `json:"name"`
}

// FieldService handles field-related API calls.
type FieldService struct {
	s *Service
}

// List returns all custom fields.
func (fs *FieldService) List(ctx context.Context, opts *gomiak.ListOptions) ([]Field, *gomiak.Pagination, error) {
	return gomiak.Do[[]Field](ctx, fs.s.client, http.MethodGet, fs.s.basePath+"/fields", opts.Values(), nil)
}

// Create creates a new custom field.
func (fs *FieldService) Create(ctx context.Context, req *CreateFieldRequest) (*Field, error) {
	f, _, err := gomiak.Do[*Field](ctx, fs.s.client, http.MethodPost, fs.s.basePath+"/fields", nil, req)
	return f, err
}

// Update updates an existing field.
func (fs *FieldService) Update(ctx context.Context, id int, req *UpdateFieldRequest) (*Field, error) {
	f, _, err := gomiak.Do[*Field](ctx, fs.s.client, http.MethodPut, fmt.Sprintf("%s/fields/%d", fs.s.basePath, id), nil, req)
	return f, err
}

// Delete deletes a single field.
func (fs *FieldService) Delete(ctx context.Context, id int) (*Field, error) {
	f, _, err := gomiak.Do[*Field](ctx, fs.s.client, http.MethodDelete, fmt.Sprintf("%s/fields/%d", fs.s.basePath, id), nil, nil)
	return f, err
}

// DeleteBulk deletes multiple fields.
func (fs *FieldService) DeleteBulk(ctx context.Context, sel *Selection) error {
	body := struct {
		Select *Selection `json:"select"`
	}{Select: sel}
	return gomiak.DoEmpty(ctx, fs.s.client, http.MethodDelete, fs.s.basePath+"/fields", nil, &body)
}
