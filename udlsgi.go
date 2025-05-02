// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

// UdlSgiService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUdlSgiService] method instead.
type UdlSgiService struct {
	Options []option.RequestOption
	History UdlSgiHistoryService
}

// NewUdlSgiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUdlSgiService(opts ...option.RequestOption) (r UdlSgiService) {
	r = UdlSgiService{}
	r.Options = opts
	r.History = NewUdlSgiHistoryService(opts...)
	return
}
