package newsletter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitcruxch/gomiak"
)

// Campaign represents a newsletter campaign.
type Campaign struct {
	ID                int                `json:"id"`
	Domain            string             `json:"domain,omitempty"`
	Status            string             `json:"status"`
	SendingBatch      *string            `json:"sending_batch"`
	Lang              string             `json:"lang"`
	Subject           string             `json:"subject"`
	Preheader         string             `json:"preheader"`
	EmailFromName     string             `json:"email_from_name"`
	EmailFromAddr     string             `json:"email_from_addr"`
	TrackingLink      bool               `json:"tracking_link"`
	TrackingOpening   bool               `json:"tracking_opening"`
	TrackingUTM       any                `json:"tracking_utm,omitempty"`
	ThumbnailURL      *string            `json:"thumbnail_url,omitempty"`
	DesignerURL       *string            `json:"designer_url,omitempty"`
	PreviewURL        string             `json:"preview_url,omitempty"`
	Content           string             `json:"content,omitempty"`
	Recipients        any                `json:"recipients,omitempty"`
	SubscribersCount  int                `json:"subscribers_count,omitempty"`
	LimitByHour       int                `json:"limit_by_hour,omitempty"`
	LimitByDay        int                `json:"limit_by_day,omitempty"`
	NextStep          string             `json:"next_step"`
	UnsubLink         *string            `json:"unsub_link,omitempty"`
	ForceSended       *string            `json:"force_sended,omitempty"`
	StartedAt         int64              `json:"started_at"`
	CreatedAt         int64              `json:"created_at"`
	Statistics        *CampaignStats     `json:"statistics,omitempty"`
	AdvancedStatistics any               `json:"advanced_statistics,omitempty"`
}

// CampaignStats holds delivery and engagement statistics for a campaign.
type CampaignStats struct {
	SentCount        int     `json:"sent_count"`
	OpenCount        int     `json:"open_count"`
	ClickCount       int     `json:"click_count"`
	DeliveryCount    int     `json:"delivery_count"`
	ComplaintCount   int     `json:"complaint_count"`
	BounceCount      int     `json:"bounce_count"`
	SoftBounceCount  int     `json:"soft_bounce_count"`
	UnsubscribeCount int     `json:"unsubscribe_count"`
	OpenRate         float64 `json:"open_rate"`
	ClickRate        float64 `json:"click_rate"`
	DeliveryRate     float64 `json:"delivery_rate"`
	ComplaintRate    float64 `json:"complaint_rate"`
	BounceRate       float64 `json:"bounce_rate"`
	SoftBounceRate   float64 `json:"soft_bounce_rate"`
	UnsubscribeRate  float64 `json:"unsubscribe_rate"`
}

// CampaignRecipients specifies target audience for a campaign.
type CampaignRecipients struct {
	AllSubscribers bool                    `json:"all_subscribers,omitempty"`
	Emails         []string                `json:"emails,omitempty"`
	Groups         *CampaignRecipientGroup `json:"groups,omitempty"`
	Segments       *CampaignRecipientGroup `json:"segments,omitempty"`
}

// CampaignRecipientGroup specifies include/exclude IDs for recipient targeting.
type CampaignRecipientGroup struct {
	Include []int `json:"include,omitempty"`
	Exclude []int `json:"exclude,omitempty"`
}

// TrackingUTM holds UTM tracking parameters.
type TrackingUTM struct {
	UTMSource   string `json:"utm_source,omitempty"`
	UTMMedium   string `json:"utm_medium,omitempty"`
	UTMCampaign string `json:"utm_campaign,omitempty"`
}

// CreateCampaignRequest is the request body for creating a campaign.
type CreateCampaignRequest struct {
	Subject         string              `json:"subject"`
	EmailFromName   string              `json:"email_from_name"`
	EmailFromAddr   string              `json:"email_from_addr"`
	Lang            string              `json:"lang"`
	Preheader       string              `json:"preheader,omitempty"`
	ContentHTML     string              `json:"content_html,omitempty"`
	TemplateID      int                 `json:"template_id,omitempty"`
	ForceSended     bool                `json:"force_sended,omitempty"`
	TrackingLink    *bool               `json:"tracking_link,omitempty"`
	TrackingOpening *bool               `json:"tracking_opening,omitempty"`
	TrackingUTM     *TrackingUTM        `json:"tracking_utm,omitempty"`
	UnsubLink       *bool               `json:"unsub_link,omitempty"`
	Recipients      *CampaignRecipients `json:"recipients,omitempty"`
}

// UpdateCampaignRequest is the request body for updating a campaign.
type UpdateCampaignRequest struct {
	Subject         string              `json:"subject,omitempty"`
	EmailFromName   string              `json:"email_from_name,omitempty"`
	EmailFromAddr   string              `json:"email_from_addr,omitempty"`
	Lang            string              `json:"lang,omitempty"`
	Preheader       string              `json:"preheader,omitempty"`
	ContentHTML     string              `json:"content_html,omitempty"`
	TemplateID      int                 `json:"template_id,omitempty"`
	ForceSended     bool                `json:"force_sended,omitempty"`
	TrackingLink    *bool               `json:"tracking_link,omitempty"`
	TrackingOpening *bool               `json:"tracking_opening,omitempty"`
	TrackingUTM     *TrackingUTM        `json:"tracking_utm,omitempty"`
	UnsubLink       *bool               `json:"unsub_link,omitempty"`
	Recipients      *CampaignRecipients `json:"recipients,omitempty"`
}

// ScheduleCampaignRequest is the request body for scheduling a campaign.
type ScheduleCampaignRequest struct {
	StartedAt int64 `json:"started_at,omitempty"`
}

// TestCampaignRequest is the request body for testing a campaign by ID.
type TestCampaignRequest struct {
	Email string `json:"email"`
}

// TestCampaignByTemplateRequest is the request body for testing via template UUID.
type TestCampaignByTemplateRequest struct {
	Emails []string `json:"emails"`
}

// CampaignService handles campaign-related API calls.
type CampaignService struct {
	s *Service
}

// List returns all campaigns.
func (cs *CampaignService) List(ctx context.Context, opts *gomiak.ListOptions) ([]Campaign, *gomiak.Pagination, error) {
	return gomiak.Do[[]Campaign](ctx, cs.s.client, http.MethodGet, cs.s.basePath+"/campaigns", opts.Values(), nil)
}

// Get returns a single campaign by ID.
func (cs *CampaignService) Get(ctx context.Context, id int) (*Campaign, error) {
	c, _, err := gomiak.Do[*Campaign](ctx, cs.s.client, http.MethodGet, fmt.Sprintf("%s/campaigns/%d", cs.s.basePath, id), nil, nil)
	return c, err
}

// Create creates a new campaign.
func (cs *CampaignService) Create(ctx context.Context, req *CreateCampaignRequest) (*Campaign, error) {
	c, _, err := gomiak.Do[*Campaign](ctx, cs.s.client, http.MethodPost, cs.s.basePath+"/campaigns", nil, req)
	return c, err
}

// Update updates an existing campaign.
func (cs *CampaignService) Update(ctx context.Context, id int, req *UpdateCampaignRequest) (*Campaign, error) {
	c, _, err := gomiak.Do[*Campaign](ctx, cs.s.client, http.MethodPut, fmt.Sprintf("%s/campaigns/%d", cs.s.basePath, id), nil, req)
	return c, err
}

// Delete deletes a single campaign.
func (cs *CampaignService) Delete(ctx context.Context, id int) error {
	return gomiak.DoEmpty(ctx, cs.s.client, http.MethodDelete, fmt.Sprintf("%s/campaigns/%d", cs.s.basePath, id), nil, nil)
}

// DeleteBulk deletes multiple campaigns.
func (cs *CampaignService) DeleteBulk(ctx context.Context, sel *Selection) (string, error) {
	body := struct {
		Select *Selection `json:"select"`
	}{Select: sel}
	opID, _, err := gomiak.Do[string](ctx, cs.s.client, http.MethodDelete, cs.s.basePath+"/campaigns", nil, &body)
	return opID, err
}

// Schedule schedules a campaign for sending.
func (cs *CampaignService) Schedule(ctx context.Context, id int, req *ScheduleCampaignRequest) error {
	return gomiak.DoEmpty(ctx, cs.s.client, http.MethodPut, fmt.Sprintf("%s/campaigns/%d/schedule", cs.s.basePath, id), nil, req)
}

// Cancel cancels a scheduled campaign.
func (cs *CampaignService) Cancel(ctx context.Context, id int) error {
	return gomiak.DoEmpty(ctx, cs.s.client, http.MethodPut, fmt.Sprintf("%s/campaigns/%d/cancel", cs.s.basePath, id), nil, nil)
}

// Duplicate duplicates a campaign.
func (cs *CampaignService) Duplicate(ctx context.Context, id int) (*Campaign, error) {
	c, _, err := gomiak.Do[*Campaign](ctx, cs.s.client, http.MethodPost, fmt.Sprintf("%s/campaigns/%d/duplicate", cs.s.basePath, id), nil, nil)
	return c, err
}

// Test sends a test email for a campaign.
func (cs *CampaignService) Test(ctx context.Context, id int, req *TestCampaignRequest) error {
	return gomiak.DoEmpty(ctx, cs.s.client, http.MethodPost, fmt.Sprintf("%s/campaigns/%d/test", cs.s.basePath, id), nil, req)
}

// TestByTemplate sends test emails for a campaign identified by template UUID.
func (cs *CampaignService) TestByTemplate(ctx context.Context, templateUUID string, req *TestCampaignByTemplateRequest) error {
	return gomiak.DoEmpty(ctx, cs.s.client, http.MethodPost, fmt.Sprintf("%s/campaigns/template/%s/test", cs.s.basePath, templateUUID), nil, req)
}

// CreateTemplateFromCampaign creates a template from an existing campaign.
func (cs *CampaignService) CreateTemplateFromCampaign(ctx context.Context, campaignID int, templateName string) (string, error) {
	id, _, err := gomiak.Do[string](ctx, cs.s.client, http.MethodPost, fmt.Sprintf("%s/campaigns/%d/template/%s", cs.s.basePath, campaignID, templateName), nil, nil)
	return id, err
}

// GetActivity returns subscriber activity statistics for a campaign.
func (cs *CampaignService) GetActivity(ctx context.Context, id int, opts *gomiak.ListOptions) (any, *gomiak.Pagination, error) {
	return gomiak.Do[any](ctx, cs.s.client, http.MethodGet, fmt.Sprintf("%s/campaigns/%d/report/activity", cs.s.basePath, id), opts.Values(), nil)
}

// GetLinks returns link activity statistics for a campaign.
func (cs *CampaignService) GetLinks(ctx context.Context, id int) (any, error) {
	data, _, err := gomiak.Do[any](ctx, cs.s.client, http.MethodGet, fmt.Sprintf("%s/campaigns/%d/report/links", cs.s.basePath, id), nil, nil)
	return data, err
}

// GetTracking returns tracking records for a campaign.
func (cs *CampaignService) GetTracking(ctx context.Context, id int, opts *gomiak.ListOptions) (any, *gomiak.Pagination, error) {
	return gomiak.Do[any](ctx, cs.s.client, http.MethodGet, fmt.Sprintf("%s/campaigns/%d/tracking", cs.s.basePath, id), opts.Values(), nil)
}
