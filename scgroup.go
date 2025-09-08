// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/Bluestaq/udl-golang-sdk/option"
)

// ScGroupService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScGroupService] method instead.
type ScGroupService struct {
	Options []option.RequestOption
}

// NewScGroupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScGroupService(opts ...option.RequestOption) (r ScGroupService) {
	r = ScGroupService{}
	r.Options = opts
	return
}
