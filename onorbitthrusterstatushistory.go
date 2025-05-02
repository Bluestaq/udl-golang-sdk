// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
)

// OnorbitthrusterstatusHistoryService contains methods and other services that
// help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitthrusterstatusHistoryService] method instead.
type OnorbitthrusterstatusHistoryService struct {
	Options []option.RequestOption
}

// NewOnorbitthrusterstatusHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOnorbitthrusterstatusHistoryService(opts ...option.RequestOption) (r OnorbitthrusterstatusHistoryService) {
	r = OnorbitthrusterstatusHistoryService{}
	r.Options = opts
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnorbitthrusterstatusHistoryService) Count(ctx context.Context, query OnorbitthrusterstatusHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onorbitthrusterstatus/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OnorbitthrusterstatusHistoryCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) ID of the
	// associated OnorbitThruster record. This ID can be used to obtain additional
	// information on an onorbit thruster object using the 'get by ID' operation (e.g.
	// /udl/onorbitthruster/{id}). For example, the OnorbitThruster object with
	// idOnorbitThruster = abc would be queried as /udl/onorbitthruster/abc.
	IDOnorbitThruster param.Opt[string] `query:"idOnorbitThruster,omitzero" json:"-"`
	MaxResults        param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) Datetime
	// of the thruster status observation in ISO 8601 UTC datetime format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StatusTime param.Opt[time.Time] `query:"statusTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnorbitthrusterstatusHistoryCountParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [OnorbitthrusterstatusHistoryCountParams]'s query parameters
// as `url.Values`.
func (r OnorbitthrusterstatusHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
