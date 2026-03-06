package gomiak_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bitcrux/gomiak"
	"github.com/bitcrux/gomiak/newsletter"
)

func TestClientConstruction(t *testing.T) {
	c := gomiak.New(
		gomiak.WithBearerToken("test-token"),
		gomiak.WithUserAgent("test/1.0"),
	)
	if c.BaseURL != "https://api.infomaniak.com" {
		t.Errorf("unexpected base URL: %s", c.BaseURL)
	}
	if c.UserAgent != "test/1.0" {
		t.Errorf("unexpected user agent: %s", c.UserAgent)
	}
}

func TestNewsletterServiceConstruction(t *testing.T) {
	c := gomiak.New(gomiak.WithBearerToken("test-token"))
	nl := newsletter.New(c, 12345)

	if nl.Campaigns == nil {
		t.Error("Campaigns service is nil")
	}
	if nl.Subscribers == nil {
		t.Error("Subscribers service is nil")
	}
	if nl.Groups == nil {
		t.Error("Groups service is nil")
	}
	if nl.Segments == nil {
		t.Error("Segments service is nil")
	}
	if nl.Fields == nil {
		t.Error("Fields service is nil")
	}
	if nl.Templates == nil {
		t.Error("Templates service is nil")
	}
	if nl.Webforms == nil {
		t.Error("Webforms service is nil")
	}
	if nl.Credits == nil {
		t.Error("Credits service is nil")
	}
	if nl.Domains == nil {
		t.Error("Domains service is nil")
	}
	if nl.Operations == nil {
		t.Error("Operations service is nil")
	}
}

func TestListCampaigns(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/1/newsletters/123/campaigns" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("missing or wrong auth header: %s", r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gomiak.Response[[]newsletter.Campaign]{
			Result: "success",
			Data: []newsletter.Campaign{
				{ID: 1, Subject: "Test Campaign", Status: "draft"},
			},
			Pagination: &gomiak.Pagination{Total: 1, Page: 1, Pages: 1, ItemsPerPage: 25},
		})
	}))
	defer srv.Close()

	c := gomiak.New(
		gomiak.WithBearerToken("test-token"),
		gomiak.WithBaseURL(srv.URL),
	)
	nl := newsletter.New(c, 123)

	campaigns, pg, err := nl.Campaigns.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(campaigns) != 1 {
		t.Fatalf("expected 1 campaign, got %d", len(campaigns))
	}
	if campaigns[0].Subject != "Test Campaign" {
		t.Errorf("unexpected subject: %s", campaigns[0].Subject)
	}
	if pg.Total != 1 {
		t.Errorf("unexpected total: %d", pg.Total)
	}
}

func TestAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"result": "error",
			"error": map[string]any{
				"code":        "not_found",
				"description": "Campaign not found",
			},
		})
	}))
	defer srv.Close()

	c := gomiak.New(
		gomiak.WithBearerToken("test-token"),
		gomiak.WithBaseURL(srv.URL),
	)
	nl := newsletter.New(c, 123)

	_, err := nl.Campaigns.Get(context.Background(), 999)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	apiErr, ok := err.(*gomiak.APIError)
	if !ok {
		t.Fatalf("expected *gomiak.APIError, got %T", err)
	}
	if apiErr.HTTPStatus != 404 {
		t.Errorf("expected status 404, got %d", apiErr.HTTPStatus)
	}
	if apiErr.Code != "not_found" {
		t.Errorf("expected code 'not_found', got %s", apiErr.Code)
	}
}

func TestListOptions(t *testing.T) {
	opts := &gomiak.ListOptions{Page: 2, PerPage: 50, OrderBy: "name", Order: "asc"}
	v := opts.Values()
	if v.Get("page") != "2" {
		t.Errorf("unexpected page: %s", v.Get("page"))
	}
	if v.Get("per_page") != "50" {
		t.Errorf("unexpected per_page: %s", v.Get("per_page"))
	}
	if v.Get("order_by") != "name" {
		t.Errorf("unexpected order_by: %s", v.Get("order_by"))
	}
}
