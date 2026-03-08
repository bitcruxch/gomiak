package newsletter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitcruxch/gomiak"
)

// Webform represents a newsletter subscription webform.
type Webform struct {
	ID                       int    `json:"id"`
	Name                     string `json:"name"`
	Title                    string `json:"title"`
	Subtitle                 string `json:"subtitle"`
	Button                   string `json:"button"`
	MsgOK                    string `json:"msg_ok"`
	MsgOKRedir               string `json:"msg_ok_redir"`
	Placeholder              string `json:"placeholder"`
	Design                   string `json:"design"`
	RGPD                     string `json:"rgpd"`
	RGPDMsg                  string `json:"rgpd_msg"`
	Notify                   string `json:"notify"`
	NotifyAddress            string `json:"notify_address"`
	NotifyLang               string `json:"notify_lang"`
	ConfirmationURL          *string `json:"confirmation_url"`
	ValidationURL            *string `json:"validation_url"`
	TemplateEmailID          int    `json:"template_email_id"`
	TemplateConfirmationID   int    `json:"template_confirmation_id"`
	TemplateValidationID     int    `json:"template_validation_id"`
	EmailTitle               string `json:"email_title"`
	EmailFromAddr            string `json:"email_from_addr"`
	EmailFromName            string `json:"email_from_name"`
	Codes                    any    `json:"codes,omitempty"`
	Fields                   any    `json:"fields,omitempty"`
	Groups                   any    `json:"groups,omitempty"`
	Statistics               any    `json:"statistics,omitempty"`
	CreatedAt                int64  `json:"created_at"`
	UpdatedAt                int64  `json:"updated_at"`
}

// WebformField represents a field configuration in a webform.
type WebformField struct {
	ID       int    `json:"id"`
	Required bool   `json:"required,omitempty"`
	Selected bool   `json:"selected,omitempty"`
	Deleted  bool   `json:"deleted,omitempty"`
	Error    string `json:"error,omitempty"`
}

// CreateWebformRequest is the request body for creating a webform.
type CreateWebformRequest struct {
	Name            string         `json:"name,omitempty"`
	Title           string         `json:"title,omitempty"`
	Subtitle        string         `json:"subtitle,omitempty"`
	Button          string         `json:"button,omitempty"`
	MsgOK           string         `json:"msg_ok,omitempty"`
	Placeholder     bool           `json:"placeholder,omitempty"`
	Design          string         `json:"design,omitempty"`
	RGPD            bool           `json:"rgpd,omitempty"`
	RGPDMsg         string         `json:"rgpd_msg,omitempty"`
	Notify          bool           `json:"notify,omitempty"`
	NotifyAddress   string         `json:"notify_address,omitempty"`
	NotifyLang      string         `json:"notify_lang,omitempty"`
	ConfirmationURL string         `json:"confirmation_url,omitempty"`
	ValidationURL   string         `json:"validation_url,omitempty"`
	EmailTitle      string         `json:"email_title,omitempty"`
	EmailFromAddr   string         `json:"email_from_addr,omitempty"`
	EmailFromName   string         `json:"email_from_name,omitempty"`
	Fields          []WebformField `json:"fields,omitempty"`
	Groups          []int          `json:"groups,omitempty"`
}

// UpdateWebformRequest is the request body for updating a webform.
type UpdateWebformRequest = CreateWebformRequest

// WebformService handles webform-related API calls.
type WebformService struct {
	s *Service
}

// List returns all webforms.
func (ws *WebformService) List(ctx context.Context, opts *gomiak.ListOptions) ([]Webform, *gomiak.Pagination, error) {
	return gomiak.Do[[]Webform](ctx, ws.s.client, http.MethodGet, ws.s.basePath+"/webforms", opts.Values(), nil)
}

// Get returns a single webform by ID.
func (ws *WebformService) Get(ctx context.Context, id int) (*Webform, error) {
	w, _, err := gomiak.Do[*Webform](ctx, ws.s.client, http.MethodGet, fmt.Sprintf("%s/webforms/%d", ws.s.basePath, id), nil, nil)
	return w, err
}

// Create creates a new webform.
func (ws *WebformService) Create(ctx context.Context, req *CreateWebformRequest) (*Webform, error) {
	w, _, err := gomiak.Do[*Webform](ctx, ws.s.client, http.MethodPost, ws.s.basePath+"/webforms", nil, req)
	return w, err
}

// Update updates an existing webform.
func (ws *WebformService) Update(ctx context.Context, id int, req *UpdateWebformRequest) (*Webform, error) {
	w, _, err := gomiak.Do[*Webform](ctx, ws.s.client, http.MethodPut, fmt.Sprintf("%s/webforms/%d", ws.s.basePath, id), nil, req)
	return w, err
}

// Delete deletes a single webform.
func (ws *WebformService) Delete(ctx context.Context, id int) (*Webform, error) {
	w, _, err := gomiak.Do[*Webform](ctx, ws.s.client, http.MethodDelete, fmt.Sprintf("%s/webforms/%d", ws.s.basePath, id), nil, nil)
	return w, err
}

// DeleteBulk deletes multiple webforms.
func (ws *WebformService) DeleteBulk(ctx context.Context, sel *Selection) error {
	body := struct {
		Select *Selection `json:"select"`
	}{Select: sel}
	return gomiak.DoEmpty(ctx, ws.s.client, http.MethodDelete, ws.s.basePath+"/webforms", nil, &body)
}

// ListThemes returns available webform themes.
func (ws *WebformService) ListThemes(ctx context.Context) (any, error) {
	data, _, err := gomiak.Do[any](ctx, ws.s.client, http.MethodGet, ws.s.basePath+"/webforms/themes", nil, nil)
	return data, err
}

// ListFields returns the fields configured for a webform.
func (ws *WebformService) ListFields(ctx context.Context, id int) (any, error) {
	data, _, err := gomiak.Do[any](ctx, ws.s.client, http.MethodGet, fmt.Sprintf("%s/webforms/%d/fields", ws.s.basePath, id), nil, nil)
	return data, err
}
