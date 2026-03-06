package newsletter

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bitcrux/gomiak"
)

// Subscriber represents a newsletter subscriber.
type Subscriber struct {
	ID           int              `json:"id"`
	Email        string           `json:"email"`
	Status       *string          `json:"status"`
	CreatedAt    int64            `json:"created_at"`
	AddedAt      int64            `json:"added_at,omitempty"`
	Source       *string          `json:"source,omitempty"`
	LastLocation *string          `json:"last_location,omitempty"`
	Statistics   *SubscriberStats `json:"statistics,omitempty"`
	Groups       []Group          `json:"groups,omitempty"`
	Sources      []Source         `json:"sources,omitempty"`
	Fields       []Field          `json:"fields,omitempty"`
}

// SubscriberStats holds engagement statistics for a subscriber.
type SubscriberStats struct {
	SentCount  int     `json:"sent_count"`
	OpenCount  int     `json:"open_count"`
	ClickCount int     `json:"click_count"`
	OpenRate   float64 `json:"open_rate"`
	ClickRate  float64 `json:"click_rate"`
}

// Source represents the origin of a subscriber.
type Source struct {
	ID             int    `json:"id"`
	SourceableType string `json:"sourceable_type"`
	SourceableID   int    `json:"sourceable_id"`
	Type           int    `json:"type"`
	IP             string `json:"ip"`
	CreatedAt      int64  `json:"created_at"`
}

// SubscriberFilter is used in bulk subscriber operations.
type SubscriberFilter struct {
	Groups  []int  `json:"groups,omitempty"`
	Search  string `json:"search,omitempty"`
	Status  string `json:"status,omitempty"`
}

// CreateSubscriberRequest is the request body for creating a subscriber.
type CreateSubscriberRequest struct {
	Email  string            `json:"email"`
	Status string            `json:"status,omitempty"`
	Groups []string          `json:"groups,omitempty"`
	Fields map[string]string `json:"fields,omitempty"`
}

// UpdateSubscriberRequest is the request body for updating a subscriber.
type UpdateSubscriberRequest struct {
	Status string            `json:"status,omitempty"`
	Groups []string          `json:"groups,omitempty"`
	Fields map[string]string `json:"fields,omitempty"`
}

// BulkDeleteSubscribersRequest is the request body for bulk deleting subscribers.
type BulkDeleteSubscribersRequest struct {
	Select  *Selection        `json:"select"`
	Filter  *SubscriberFilter `json:"filter,omitempty"`
	Segment string            `json:"segment,omitempty"`
}

// AssignSubscribersRequest is the request body for assigning subscribers to groups.
type AssignSubscribersRequest struct {
	Select  *Selection        `json:"select"`
	Groups  []string          `json:"groups,omitempty"`
	Filter  *SubscriberFilter `json:"filter,omitempty"`
	Segment string            `json:"segment,omitempty"`
}

// UnassignSubscribersRequest is the request body for unassigning subscribers from a group.
type UnassignSubscribersRequest struct {
	Select  *Selection        `json:"select"`
	GroupID int               `json:"group_id"`
	Filter  *SubscriberFilter `json:"filter,omitempty"`
	Segment string            `json:"segment,omitempty"`
}

// UnsubscribeSubscribersRequest is the request body for bulk unsubscribing.
type UnsubscribeSubscribersRequest struct {
	Select  *Selection        `json:"select"`
	Filter  *SubscriberFilter `json:"filter,omitempty"`
	Segment string            `json:"segment,omitempty"`
}

// ExportSubscribersRequest is the request body for exporting subscribers.
type ExportSubscribersRequest struct {
	Select  *Selection        `json:"select"`
	Filter  *SubscriberFilter `json:"filter,omitempty"`
	Segment string            `json:"segment,omitempty"`
}

// FilterSubscribersRequest is the request body for filtering subscribers.
type FilterSubscribersRequest struct {
	Filter  *SubscriberFilter `json:"filter,omitempty"`
	Segment string            `json:"segment,omitempty"`
}

// ImportSubscribersRequest is the request body for importing subscribers.
type ImportSubscribersRequest struct {
	Fields        []string `json:"fields"`
	Groups        []string `json:"groups,omitempty"`
	UploadID      int      `json:"upload_id,omitempty"`
	IPDUUID       string   `json:"ipd_uuid,omitempty"`
	ReplaceFields *bool    `json:"replace_fields,omitempty"`
	CSVSeparator  string   `json:"csv_separator,omitempty"`
	CSVEnclosure  string   `json:"csv_enclosure,omitempty"`
}

// UploadCSVRequest is the request body for uploading a CSV file.
type UploadCSVRequest struct {
	File    string `json:"file,omitempty"`
	Content string `json:"content,omitempty"`
}

// SubscriberService handles subscriber-related API calls.
type SubscriberService struct {
	s *Service
}

// List returns all subscribers.
func (ss *SubscriberService) List(ctx context.Context, opts *gomiak.ListOptions) ([]Subscriber, *gomiak.Pagination, error) {
	return gomiak.Do[[]Subscriber](ctx, ss.s.client, http.MethodGet, ss.s.basePath+"/subscribers", opts.Values(), nil)
}

// Get returns a single subscriber by ID.
func (ss *SubscriberService) Get(ctx context.Context, id int) (*Subscriber, error) {
	s, _, err := gomiak.Do[*Subscriber](ctx, ss.s.client, http.MethodGet, fmt.Sprintf("%s/subscribers/%d", ss.s.basePath, id), nil, nil)
	return s, err
}

// Create creates a new subscriber.
func (ss *SubscriberService) Create(ctx context.Context, req *CreateSubscriberRequest) (*Subscriber, error) {
	s, _, err := gomiak.Do[*Subscriber](ctx, ss.s.client, http.MethodPost, ss.s.basePath+"/subscribers", nil, req)
	return s, err
}

// Update updates an existing subscriber.
func (ss *SubscriberService) Update(ctx context.Context, id int, req *UpdateSubscriberRequest) (*Subscriber, error) {
	s, _, err := gomiak.Do[*Subscriber](ctx, ss.s.client, http.MethodPut, fmt.Sprintf("%s/subscribers/%d", ss.s.basePath, id), nil, req)
	return s, err
}

// Delete deletes a single subscriber.
func (ss *SubscriberService) Delete(ctx context.Context, id int) error {
	return gomiak.DoEmpty(ctx, ss.s.client, http.MethodDelete, fmt.Sprintf("%s/subscribers/%d", ss.s.basePath, id), nil, nil)
}

// Forget permanently deletes a subscriber (GDPR-compliant).
func (ss *SubscriberService) Forget(ctx context.Context, id int) error {
	return gomiak.DoEmpty(ctx, ss.s.client, http.MethodDelete, fmt.Sprintf("%s/subscribers/%d/forget", ss.s.basePath, id), nil, nil)
}

// DeleteBulk deletes multiple subscribers. Returns an operation ID.
func (ss *SubscriberService) DeleteBulk(ctx context.Context, req *BulkDeleteSubscribersRequest) (string, error) {
	opID, _, err := gomiak.Do[string](ctx, ss.s.client, http.MethodDelete, ss.s.basePath+"/subscribers", nil, req)
	return opID, err
}

// Assign assigns subscribers to groups. Returns an operation ID.
func (ss *SubscriberService) Assign(ctx context.Context, req *AssignSubscribersRequest) (string, error) {
	opID, _, err := gomiak.Do[string](ctx, ss.s.client, http.MethodPut, ss.s.basePath+"/subscribers/assign", nil, req)
	return opID, err
}

// Unassign unassigns subscribers from a group. Returns an operation ID.
func (ss *SubscriberService) Unassign(ctx context.Context, req *UnassignSubscribersRequest) (string, error) {
	opID, _, err := gomiak.Do[string](ctx, ss.s.client, http.MethodPut, ss.s.basePath+"/subscribers/unassign", nil, req)
	return opID, err
}

// Unsubscribe bulk unsubscribes subscribers. Returns an operation ID.
func (ss *SubscriberService) Unsubscribe(ctx context.Context, req *UnsubscribeSubscribersRequest) (string, error) {
	opID, _, err := gomiak.Do[string](ctx, ss.s.client, http.MethodPut, ss.s.basePath+"/subscribers/unsubscribe", nil, req)
	return opID, err
}

// Export exports subscribers. Returns an operation ID.
func (ss *SubscriberService) Export(ctx context.Context, req *ExportSubscribersRequest) (string, error) {
	opID, _, err := gomiak.Do[string](ctx, ss.s.client, http.MethodPost, ss.s.basePath+"/subscribers/export", nil, req)
	return opID, err
}

// Filter returns subscribers matching advanced filter criteria.
func (ss *SubscriberService) Filter(ctx context.Context, opts *gomiak.ListOptions, req *FilterSubscribersRequest) ([]Subscriber, *gomiak.Pagination, error) {
	return gomiak.Do[[]Subscriber](ctx, ss.s.client, http.MethodPost, ss.s.basePath+"/subscribers/filter", opts.Values(), req)
}

// CountStatus returns subscriber counts by status.
func (ss *SubscriberService) CountStatus(ctx context.Context) (any, error) {
	data, _, err := gomiak.Do[any](ctx, ss.s.client, http.MethodGet, ss.s.basePath+"/subscribers/count_status", nil, nil)
	return data, err
}

// Import imports subscribers from an uploaded CSV.
func (ss *SubscriberService) Import(ctx context.Context, req *ImportSubscribersRequest) error {
	return gomiak.DoEmpty(ctx, ss.s.client, http.MethodPost, ss.s.basePath+"/subscribers/import", nil, req)
}

// UploadCSV uploads a CSV file for import. Returns preview data.
func (ss *SubscriberService) UploadCSV(ctx context.Context, req *UploadCSVRequest) ([]string, error) {
	data, _, err := gomiak.Do[[]string](ctx, ss.s.client, http.MethodPost, ss.s.basePath+"/subscribers/import/upload", nil, req)
	return data, err
}

// addWith adds a "with" query parameter if non-empty.
func addWith(v url.Values, with string) url.Values {
	if with == "" {
		return v
	}
	if v == nil {
		v = url.Values{}
	}
	v.Set("with", with)
	return v
}
