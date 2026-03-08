package newsletter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitcruxch/gomiak"
)

// Segment represents a subscriber segment with dynamic conditions.
type Segment struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Data            any     `json:"data,omitempty"`
	SubscribedCount int     `json:"subscribed_count"`
	UpdatedAt       int64   `json:"updated_at"`
	Statistics      any     `json:"statistics,omitempty"`
}

// SegmentConditions defines the top-level condition structure.
type SegmentConditions struct {
	Closure string           `json:"closure,omitempty"` // "OR"
	Conds   []ConditionGroup `json:"conds,omitempty"`
}

// ConditionGroup is a group of conditions joined by AND.
type ConditionGroup struct {
	Closure string      `json:"closure"` // "AND"
	Conds   []Condition `json:"conds"`
}

// Condition is a single filter condition.
type Condition struct {
	Type          string `json:"type"`                     // "campaign", "created_at", "mailinglist", "meta"
	Operator      string `json:"operator"`                 // "<", "<>", "=", ">", "between", "click", etc.
	ActionID      string `json:"actionId,omitempty"`       // "all", "any", "custom", "last5"
	CampaignID    int    `json:"campaign_id,omitempty"`
	Date1         string `json:"date1,omitempty"`          // Y-m-d
	Date2         string `json:"date2,omitempty"`
	MailinglistID int    `json:"mailinglist_id,omitempty"`
	MetaID        int    `json:"meta_id,omitempty"`
	Value         string `json:"value,omitempty"`
}

// CreateSegmentRequest is the request body for creating a segment.
type CreateSegmentRequest struct {
	Name string             `json:"name"`
	Data *SegmentConditions `json:"data"`
}

// UpdateSegmentRequest is the request body for updating a segment.
type UpdateSegmentRequest struct {
	Name string             `json:"name,omitempty"`
	Data *SegmentConditions `json:"data,omitempty"`
}

// SegmentService handles segment-related API calls.
type SegmentService struct {
	s *Service
}

// List returns all segments.
func (ss *SegmentService) List(ctx context.Context, opts *gomiak.ListOptions) ([]Segment, *gomiak.Pagination, error) {
	return gomiak.Do[[]Segment](ctx, ss.s.client, http.MethodGet, ss.s.basePath+"/segments", opts.Values(), nil)
}

// Get returns a single segment by ID.
func (ss *SegmentService) Get(ctx context.Context, id int) (*Segment, error) {
	s, _, err := gomiak.Do[*Segment](ctx, ss.s.client, http.MethodGet, fmt.Sprintf("%s/segments/%d", ss.s.basePath, id), nil, nil)
	return s, err
}

// Create creates a new segment.
func (ss *SegmentService) Create(ctx context.Context, req *CreateSegmentRequest) (*Segment, error) {
	s, _, err := gomiak.Do[*Segment](ctx, ss.s.client, http.MethodPost, ss.s.basePath+"/segments", nil, req)
	return s, err
}

// Update updates an existing segment.
func (ss *SegmentService) Update(ctx context.Context, id int, req *UpdateSegmentRequest) (*Segment, error) {
	s, _, err := gomiak.Do[*Segment](ctx, ss.s.client, http.MethodPut, fmt.Sprintf("%s/segments/%d", ss.s.basePath, id), nil, req)
	return s, err
}

// Delete deletes a single segment.
func (ss *SegmentService) Delete(ctx context.Context, id int) error {
	return gomiak.DoEmpty(ctx, ss.s.client, http.MethodDelete, fmt.Sprintf("%s/segments/%d", ss.s.basePath, id), nil, nil)
}

// DeleteBulk deletes multiple segments.
func (ss *SegmentService) DeleteBulk(ctx context.Context, sel *Selection) error {
	body := struct {
		Select *Selection `json:"select"`
	}{Select: sel}
	return gomiak.DoEmpty(ctx, ss.s.client, http.MethodDelete, ss.s.basePath+"/segments", nil, &body)
}

// ListSubscribers returns subscribers matching a segment.
func (ss *SegmentService) ListSubscribers(ctx context.Context, id int, opts *gomiak.ListOptions) ([]Subscriber, *gomiak.Pagination, error) {
	return gomiak.Do[[]Subscriber](ctx, ss.s.client, http.MethodGet, fmt.Sprintf("%s/segments/%d/subscribers", ss.s.basePath, id), opts.Values(), nil)
}
