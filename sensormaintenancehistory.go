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

// SensormaintenanceHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSensormaintenanceHistoryService] method instead.
type SensormaintenanceHistoryService struct {
	Options []option.RequestOption
}

// NewSensormaintenanceHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSensormaintenanceHistoryService(opts ...option.RequestOption) (r SensormaintenanceHistoryService) {
	r = SensormaintenanceHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensormaintenanceHistoryService) Get(ctx context.Context, query SensormaintenanceHistoryGetParams, opts ...option.RequestOption) (res *[]SensormaintenanceFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensormaintenance/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensormaintenanceHistoryService) Aodr(ctx context.Context, query SensormaintenanceHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensormaintenance/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SensormaintenanceHistoryService) Count(ctx context.Context, query SensormaintenanceHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sensormaintenance/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SensormaintenanceHistoryGetParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns param.Opt[string] `query:"columns,omitzero" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// end time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EndTime     param.Opt[time.Time] `query:"endTime,omitzero" format:"date-time" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceHistoryGetParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [SensormaintenanceHistoryGetParams]'s query parameters as
// `url.Values`.
func (r SensormaintenanceHistoryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensormaintenanceHistoryAodrParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns param.Opt[string] `query:"columns,omitzero" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// end time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EndTime     param.Opt[time.Time] `query:"endTime,omitzero" format:"date-time" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// optional, notification method for the created file link. When omitted, EMAIL is
	// assumed. Current valid values are: EMAIL, SMS.
	Notification param.Opt[string] `query:"notification,omitzero" json:"-"`
	// optional, field delimiter when the created file is not JSON. Must be a single
	// character chosen from this set: (',', ';', ':', '|'). When omitted, "," is used.
	// It is strongly encouraged that your field delimiter be a character unlikely to
	// occur within the data.
	OutputDelimiter param.Opt[string] `query:"outputDelimiter,omitzero" json:"-"`
	// optional, output format for the file. When omitted, JSON is assumed. Current
	// valid values are: JSON and CSV.
	OutputFormat param.Opt[string] `query:"outputFormat,omitzero" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceHistoryAodrParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [SensormaintenanceHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r SensormaintenanceHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensormaintenanceHistoryCountParams struct {
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// end time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EndTime     param.Opt[time.Time] `query:"endTime,omitzero" format:"date-time" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceHistoryCountParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [SensormaintenanceHistoryCountParams]'s query parameters as
// `url.Values`.
func (r SensormaintenanceHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
