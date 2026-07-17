// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/Bluestaq/udl-golang-sdk/v2/option"
)

// TdoaFdoaService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTdoaFdoaService] method instead.
type TdoaFdoaService struct {
	Options []option.RequestOption
	// These services provide operations for querying and manipulation of Signal time
	// and frequency difference of arrival (TDOA/FDOA) information obtained by using
	// passive RF based sensor phenomenologies and sensor triangulation. The J2000
	// coordinate frame is the preferred frame for all observations, but in some cases
	// observations may be in another frame depending on the provider.
	Diffofarrival TdoaFdoaDiffofarrivalService
}

// NewTdoaFdoaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTdoaFdoaService(opts ...option.RequestOption) (r TdoaFdoaService) {
	r = TdoaFdoaService{}
	r.Options = opts
	r.Diffofarrival = NewTdoaFdoaDiffofarrivalService(opts...)
	return
}
