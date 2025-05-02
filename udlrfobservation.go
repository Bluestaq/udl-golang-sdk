// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

// UdlRfobservationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUdlRfobservationService] method instead.
type UdlRfobservationService struct {
	Options []option.RequestOption
	History UdlRfobservationHistoryService
}

// NewUdlRfobservationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUdlRfobservationService(opts ...option.RequestOption) (r UdlRfobservationService) {
	r = UdlRfobservationService{}
	r.Options = opts
	r.History = NewUdlRfobservationHistoryService(opts...)
	return
}
