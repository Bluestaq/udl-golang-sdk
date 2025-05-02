// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

// ScClassificationMarkingService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScClassificationMarkingService] method instead.
type ScClassificationMarkingService struct {
	Options []option.RequestOption
}

// NewScClassificationMarkingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScClassificationMarkingService(opts ...option.RequestOption) (r ScClassificationMarkingService) {
	r = ScClassificationMarkingService{}
	r.Options = opts
	return
}

// Returns a list of all classification markings appropriate to the current user.
func (r *ScClassificationMarkingService) List(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	path := "scs/getClassificationMarkings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
