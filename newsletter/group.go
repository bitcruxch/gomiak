package newsletter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitcrux/gomiak"
)

// Group represents a subscriber group.
type Group struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	UpdatedAt       int64       `json:"updated_at"`
	SubscribedCount string      `json:"subscribed_count,omitempty"`
	Statistics      *GroupStats `json:"statistics,omitempty"`
	Domain          string      `json:"domain,omitempty"`
}

// GroupStats holds engagement statistics for a group.
type GroupStats struct {
	OpenRate  float64 `json:"open_rate"`
	ClickRate float64 `json:"click_rate"`
}

// CreateGroupRequest is the request body for creating a group.
type CreateGroupRequest struct {
	Name string `json:"name"`
}

// UpdateGroupRequest is the request body for updating a group.
type UpdateGroupRequest struct {
	Name string `json:"name"`
}

// GroupAssignRequest is the request body for assigning subscribers to a group.
type GroupAssignRequest struct {
	SubscriberIDs []int `json:"subscriber_ids"`
}

// GroupService handles group-related API calls.
type GroupService struct {
	s *Service
}

// List returns all groups.
func (gs *GroupService) List(ctx context.Context, opts *gomiak.ListOptions) ([]Group, *gomiak.Pagination, error) {
	return gomiak.Do[[]Group](ctx, gs.s.client, http.MethodGet, gs.s.basePath+"/groups", opts.Values(), nil)
}

// Get returns a single group by ID.
func (gs *GroupService) Get(ctx context.Context, id int) (*Group, error) {
	g, _, err := gomiak.Do[*Group](ctx, gs.s.client, http.MethodGet, fmt.Sprintf("%s/groups/%d", gs.s.basePath, id), nil, nil)
	return g, err
}

// Create creates a new group.
func (gs *GroupService) Create(ctx context.Context, req *CreateGroupRequest) (*Group, error) {
	g, _, err := gomiak.Do[*Group](ctx, gs.s.client, http.MethodPost, gs.s.basePath+"/groups", nil, req)
	return g, err
}

// Update updates an existing group.
func (gs *GroupService) Update(ctx context.Context, id int, req *UpdateGroupRequest) (*Group, error) {
	g, _, err := gomiak.Do[*Group](ctx, gs.s.client, http.MethodPut, fmt.Sprintf("%s/groups/%d", gs.s.basePath, id), nil, req)
	return g, err
}

// Delete deletes a single group.
func (gs *GroupService) Delete(ctx context.Context, id int) error {
	return gomiak.DoEmpty(ctx, gs.s.client, http.MethodDelete, fmt.Sprintf("%s/groups/%d", gs.s.basePath, id), nil, nil)
}

// DeleteBulk deletes multiple groups. Returns an operation ID.
func (gs *GroupService) DeleteBulk(ctx context.Context, sel *Selection) (string, error) {
	body := struct {
		Select *Selection `json:"select"`
	}{Select: sel}
	opID, _, err := gomiak.Do[string](ctx, gs.s.client, http.MethodDelete, gs.s.basePath+"/groups", nil, &body)
	return opID, err
}

// ListSubscribers returns subscribers belonging to a group.
func (gs *GroupService) ListSubscribers(ctx context.Context, id int, opts *gomiak.ListOptions) ([]Subscriber, *gomiak.Pagination, error) {
	return gomiak.Do[[]Subscriber](ctx, gs.s.client, http.MethodGet, fmt.Sprintf("%s/groups/%d/subscribers", gs.s.basePath, id), opts.Values(), nil)
}

// AssignSubscribers assigns subscribers to a group by their IDs.
func (gs *GroupService) AssignSubscribers(ctx context.Context, groupID int, req *GroupAssignRequest) ([]Subscriber, error) {
	data, _, err := gomiak.Do[[]Subscriber](ctx, gs.s.client, http.MethodPost, fmt.Sprintf("%s/groups/%d/subscribers/assign", gs.s.basePath, groupID), nil, req)
	return data, err
}

// UnassignSubscribers removes subscribers from a group by their IDs.
func (gs *GroupService) UnassignSubscribers(ctx context.Context, groupID int, req *GroupAssignRequest) ([]Subscriber, error) {
	data, _, err := gomiak.Do[[]Subscriber](ctx, gs.s.client, http.MethodPost, fmt.Sprintf("%s/groups/%d/subscribers/unassign", gs.s.basePath, groupID), nil, req)
	return data, err
}
