// Package newsletter provides a client for the Infomaniak Newsletter API.
package newsletter

import (
	"fmt"

	"github.com/bitcruxch/gomiak"
)

// Service provides access to the Infomaniak Newsletter API for a specific domain.
type Service struct {
	client   *gomiak.Client
	basePath string

	Campaigns  *CampaignService
	Subscribers *SubscriberService
	Groups     *GroupService
	Segments   *SegmentService
	Fields     *FieldService
	Templates  *TemplateService
	Webforms   *WebformService
	Credits    *CreditService
	Domains    *DomainService
	Operations *OperationService
}

// New creates a Newsletter service bound to a specific domain ID.
func New(client *gomiak.Client, domainID int) *Service {
	s := &Service{
		client:   client,
		basePath: fmt.Sprintf("/1/newsletters/%d", domainID),
	}
	s.Campaigns = &CampaignService{s: s}
	s.Subscribers = &SubscriberService{s: s}
	s.Groups = &GroupService{s: s}
	s.Segments = &SegmentService{s: s}
	s.Fields = &FieldService{s: s}
	s.Templates = &TemplateService{s: s}
	s.Webforms = &WebformService{s: s}
	s.Credits = &CreditService{s: s}
	s.Domains = &DomainService{s: s}
	s.Operations = &OperationService{s: s}
	return s
}

// Selection is used for bulk operations to select items by inclusion/exclusion.
type Selection struct {
	All     bool  `json:"all"`
	Include []int `json:"include,omitempty"`
	Exclude []int `json:"exclude,omitempty"`
}
