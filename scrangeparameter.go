// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"

	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

// ScRangeParameterService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScRangeParameterService] method instead.
type ScRangeParameterService struct {
	Options []option.RequestOption
}

// NewScRangeParameterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScRangeParameterService(opts ...option.RequestOption) (r ScRangeParameterService) {
	r = ScRangeParameterService{}
	r.Options = opts
	return
}

// Returns a set of File Metadata that can be used for search endpoint.
func (r *ScRangeParameterService) List(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	path := "scs/listRangeParameters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
