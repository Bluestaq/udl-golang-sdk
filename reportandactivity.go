// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/Bluestaq/udl-golang-sdk/v2/option"
)

// ReportAndActivityService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportAndActivityService] method instead.
type ReportAndActivityService struct {
	Options []option.RequestOption
	Poi     ReportAndActivityPoiService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	UdlH3geo ReportAndActivityUdlH3geoService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	UdlSigact ReportAndActivityUdlSigactService
}

// NewReportAndActivityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReportAndActivityService(opts ...option.RequestOption) (r ReportAndActivityService) {
	r = ReportAndActivityService{}
	r.Options = opts
	r.Poi = NewReportAndActivityPoiService(opts...)
	r.UdlH3geo = NewReportAndActivityUdlH3geoService(opts...)
	r.UdlSigact = NewReportAndActivityUdlSigactService(opts...)
	return
}
