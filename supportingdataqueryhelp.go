// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

// SupportingDataQueryHelpService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportingDataQueryHelpService] method instead.
type SupportingDataQueryHelpService struct {
	Options []option.RequestOption
}

// NewSupportingDataQueryHelpService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSupportingDataQueryHelpService(opts ...option.RequestOption) (r SupportingDataQueryHelpService) {
	r = SupportingDataQueryHelpService{}
	r.Options = opts
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SupportingDataQueryHelpService) Get(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/dataowner/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}
