// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

// UdlManeuverService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUdlManeuverService] method instead.
type UdlManeuverService struct {
	Options []option.RequestOption
	History UdlManeuverHistoryService
}

// NewUdlManeuverService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUdlManeuverService(opts ...option.RequestOption) (r UdlManeuverService) {
	r = UdlManeuverService{}
	r.Options = opts
	r.History = NewUdlManeuverHistoryService(opts...)
	return
}
