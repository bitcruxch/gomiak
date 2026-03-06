package newsletter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitcrux/gomiak"
)

// Credit represents a newsletter credit record.
type Credit struct {
	ID              int    `json:"id"`
	DomainName      string `json:"domain_name,omitempty"`
	CampaignID      int    `json:"campaign_id,omitempty"`
	CampaignSubject string `json:"campaign_subject,omitempty"`
	Quantity        int    `json:"quantity"`
	Type            string `json:"type"`
	Genre           string `json:"genre"`
	CreatedAt       int64  `json:"created_at"`
}

// CreditService handles credit-related API calls.
type CreditService struct {
	s *Service
}

// List returns all credits.
func (cs *CreditService) List(ctx context.Context, opts *gomiak.ListOptions) ([]Credit, *gomiak.Pagination, error) {
	return gomiak.Do[[]Credit](ctx, cs.s.client, http.MethodGet, cs.s.basePath+"/credits", opts.Values(), nil)
}

// GetAccount returns account-level credit information.
func (cs *CreditService) GetAccount(ctx context.Context, opts *gomiak.ListOptions) (*Credit, *gomiak.Pagination, error) {
	return gomiak.Do[*Credit](ctx, cs.s.client, http.MethodGet, cs.s.basePath+"/credits/accounts", opts.Values(), nil)
}

// GetDetails returns credit details.
func (cs *CreditService) GetDetails(ctx context.Context) (any, error) {
	data, _, err := gomiak.Do[any](ctx, cs.s.client, http.MethodGet, cs.s.basePath+"/credits/details", nil, nil)
	return data, err
}

// ListPacks returns available credit packs for purchase.
func (cs *CreditService) ListPacks(ctx context.Context) (any, error) {
	data, _, err := gomiak.Do[any](ctx, cs.s.client, http.MethodGet, cs.s.basePath+"/credits/packs", nil, nil)
	return data, err
}

// Checkout redirects to checkout for a credit pack.
func (cs *CreditService) Checkout(ctx context.Context, packID int) error {
	return gomiak.DoEmpty(ctx, cs.s.client, http.MethodGet, fmt.Sprintf("%s/credits/checkout/%d", cs.s.basePath, packID), nil, nil)
}
