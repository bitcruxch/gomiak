package newsletter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitcruxch/gomiak"
)

// OperationService handles async operation management.
type OperationService struct {
	s *Service
}

// Cancel cancels an outgoing operation by its ID.
func (os *OperationService) Cancel(ctx context.Context, operationID string) error {
	return gomiak.DoEmpty(ctx, os.s.client, http.MethodDelete, fmt.Sprintf("%s/operations/%s", os.s.basePath, operationID), nil, nil)
}
