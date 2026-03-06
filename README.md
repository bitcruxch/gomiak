# gomiak

A Go client library for the [Infomaniak API](https://developer.infomaniak.com/).

> **Status:** Early development. The Newsletter API client is implemented. Additional API clients (core platform, kSuite, etc.) will follow.

## Installation

```bash
go get github.com/bitcrux/gomiak
```

To use only the Newsletter client:

```go
import "github.com/bitcrux/gomiak/newsletter"
```

Go's module system will only pull the dependencies needed by the packages you import — no unrelated API clients are included.

**Requirements:** Go 1.22+, zero external dependencies (stdlib only).

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bitcrux/gomiak"
	"github.com/bitcrux/gomiak/newsletter"
)

func main() {
	client := gomiak.New(gomiak.WithBearerToken("your-api-token"))
	nl := newsletter.New(client, 12345) // your newsletter domain ID

	ctx := context.Background()

	// List campaigns
	campaigns, pg, err := nl.Campaigns.List(ctx, &gomiak.ListOptions{Page: 1, PerPage: 25})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d campaigns (page %d/%d)\n", pg.Total, pg.Page, pg.Pages)
	for _, c := range campaigns {
		fmt.Printf("  [%d] %s (%s)\n", c.ID, c.Subject, c.Status)
	}

	// Create a subscriber
	sub, err := nl.Subscribers.Create(ctx, &newsletter.CreateSubscriberRequest{
		Email:  "user@example.com",
		Groups: []string{"newsletter"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created subscriber %d: %s\n", sub.ID, sub.Email)
}
```

## Authentication

### Bearer Token

The simplest method. Create an API token in the [Infomaniak Manager](https://manager.infomaniak.com/v3/ng/accounts/token/list).

```go
client := gomiak.New(gomiak.WithBearerToken("your-api-token"))
```

### OAuth2

Bring your own `*http.Client` from [`golang.org/x/oauth2`](https://pkg.go.dev/golang.org/x/oauth2):

```go
client := gomiak.New(gomiak.WithHTTPClient(oauth2Client))
```

See the [Infomaniak OAuth2 documentation](https://developer.infomaniak.com/docs/api/authentication/oauth2) for details.

## Client Options

| Option | Description |
|--------|-------------|
| `WithBearerToken(token)` | Set a static bearer token for authentication |
| `WithHTTPClient(client)` | Use a custom `*http.Client` (e.g., for OAuth2) |
| `WithBaseURL(url)` | Override the API base URL (default: `https://api.infomaniak.com`) |
| `WithUserAgent(ua)` | Set a custom User-Agent header |

## Newsletter API

The Newsletter client covers all 70 operations from the [Infomaniak Newsletter API](https://developer.infomaniak.com/docs/api/post/1/newsletters/-domain-/campaigns).

Create a service bound to a newsletter domain ID:

```go
nl := newsletter.New(client, domainID)
```

### Sub-services

| Service | Methods | Description |
|---------|---------|-------------|
| `nl.Campaigns` | List, Get, Create, Update, Delete, DeleteBulk, Schedule, Cancel, Duplicate, Test, TestByTemplate, CreateTemplateFromCampaign, GetActivity, GetLinks, GetTracking | Campaign management and reporting |
| `nl.Subscribers` | List, Get, Create, Update, Delete, Forget, DeleteBulk, Assign, Unassign, Unsubscribe, Export, Filter, CountStatus, Import, UploadCSV | Subscriber lifecycle management |
| `nl.Groups` | List, Get, Create, Update, Delete, DeleteBulk, ListSubscribers, AssignSubscribers, UnassignSubscribers | Subscriber grouping |
| `nl.Segments` | List, Get, Create, Update, Delete, DeleteBulk, ListSubscribers | Dynamic segmentation |
| `nl.Fields` | List, Create, Update, Delete, DeleteBulk | Custom subscriber fields |
| `nl.Templates` | List, GetHTML, UpdateThumbnail | Email templates |
| `nl.Webforms` | List, Get, Create, Update, Delete, DeleteBulk, ListThemes, ListFields | Subscription webforms |
| `nl.Credits` | List, GetAccount, GetDetails, ListPacks, Checkout | Credit management |
| `nl.Domains` | Get, Delete, GetAPIKey | Newsletter domain info |
| `nl.Operations` | Cancel | Cancel async operations |

### Return Patterns

```go
// List endpoints return a slice, pagination metadata, and an error
campaigns, pagination, err := nl.Campaigns.List(ctx, opts)

// Single-resource endpoints return a pointer and an error
campaign, err := nl.Campaigns.Get(ctx, id)

// Action endpoints return just an error
err := nl.Campaigns.Cancel(ctx, id)

// Bulk async operations return an operation ID and an error
operationID, err := nl.Subscribers.DeleteBulk(ctx, req)
```

### Error Handling

API errors are returned as `*gomiak.APIError`:

```go
campaign, err := nl.Campaigns.Get(ctx, 999)
if err != nil {
	var apiErr *gomiak.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("HTTP %d: %s — %s\n", apiErr.HTTPStatus, apiErr.Code, apiErr.Description)
	}
}
```

### Pagination

```go
opts := &gomiak.ListOptions{
	Page:    1,
	PerPage: 50,
	OrderBy: "created_at",
	Order:   "desc",
}
campaigns, pg, err := nl.Campaigns.List(ctx, opts)
// pg.Total, pg.Page, pg.Pages, pg.ItemsPerPage
```

## API Reference

- [Infomaniak API Documentation](https://developer.infomaniak.com/)
- [Infomaniak API Authentication](https://developer.infomaniak.com/docs/api/authentication)
- [Newsletter API Endpoints](https://developer.infomaniak.com/docs/api/post/1/newsletters/-domain-/campaigns)
- [API Token Management](https://manager.infomaniak.com/v3/ng/accounts/token/list)

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

## Disclaimer

This is an unofficial, community-maintained client library. It is not affiliated with, endorsed by, or supported by [Infomaniak Network SA](https://www.infomaniak.com/). "Infomaniak" is a trademark of Infomaniak Network SA.
