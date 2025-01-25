package interfaces

import (
	"context"
	"database_anonymizer/app/src/structs"
)

type ServicesInterface interface {
	ValidateRules(request structs.AnonymizationRequest) error
	Anonymize(ctx context.Context, request structs.AnonymizationRequest, cacheKey string) error
}
