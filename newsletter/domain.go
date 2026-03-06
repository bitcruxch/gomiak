package newsletter

import (
	"context"
	"net/http"

	"github.com/bitcrux/gomiak"
)

// Domain represents a newsletter domain.
type Domain struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	CreditMonth         int     `json:"credit_month"`
	LimitByHour         int     `json:"limit_by_hour,omitempty"`
	SenderType          string  `json:"sender_type"`
	Status              string  `json:"status"` // "blocked", "disabled", "enabled", "error", "suspended", "unknown", "waiting"
	DNS                 any     `json:"dns,omitempty"`
	AWSIdentity         string  `json:"aws_identity,omitempty"`
	CanUseOtherDomains  bool    `json:"can_use_other_domains"`
	SubDomains          any     `json:"sub_domains,omitempty"`
	Statistics          any     `json:"statistics,omitempty"`
	CanUseNewMailBuilder bool   `json:"can_use_new_mail_builder,omitempty"`
	CurrentCredit       string  `json:"current_credit,omitempty"`
}

// DomainService handles domain-related API calls.
type DomainService struct {
	s *Service
}

// Get returns domain information.
func (ds *DomainService) Get(ctx context.Context) (*Domain, error) {
	d, _, err := gomiak.Do[*Domain](ctx, ds.s.client, http.MethodGet, ds.s.basePath, nil, nil)
	return d, err
}

// Delete permanently deletes the domain.
func (ds *DomainService) Delete(ctx context.Context) error {
	return gomiak.DoEmpty(ctx, ds.s.client, http.MethodDelete, ds.s.basePath, nil, nil)
}

// GetAPIKey returns API authentication information.
func (ds *DomainService) GetAPIKey(ctx context.Context) ([]string, error) {
	data, _, err := gomiak.Do[[]string](ctx, ds.s.client, http.MethodGet, ds.s.basePath+"/api-key", nil, nil)
	return data, err
}
