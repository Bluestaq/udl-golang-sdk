// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
)

// EphemerisSetHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEphemerisSetHistoryService] method instead.
type EphemerisSetHistoryService struct {
	Options []option.RequestOption
}

// NewEphemerisSetHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEphemerisSetHistoryService(opts ...option.RequestOption) (r EphemerisSetHistoryService) {
	r = EphemerisSetHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EphemerisSetHistoryService) List(ctx context.Context, query EphemerisSetHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EphemerisSet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/ephemerisset/history"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EphemerisSetHistoryService) ListAutoPaging(ctx context.Context, query EphemerisSetHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EphemerisSet] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EphemerisSetHistoryService) Aodr(ctx context.Context, query EphemerisSetHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/ephemerisset/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EphemerisSetHistoryService) Count(ctx context.Context, query EphemerisSetHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/ephemerisset/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type EphemerisSetHistoryListParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'pointEndTime, pointStartTime' are required.) End
	// time/last time point of the ephemeris, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	PointEndTime param.Opt[time.Time] `query:"pointEndTime,omitzero" format:"date-time" json:"-"`
	// (One or more of fields 'pointEndTime, pointStartTime' are required.) Start
	// time/first time point of the ephemeris, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	PointStartTime param.Opt[time.Time] `query:"pointStartTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EphemerisSetHistoryListParams]'s query parameters as
// `url.Values`.
func (r EphemerisSetHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EphemerisSetHistoryAodrParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
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
	// (One or more of fields 'pointEndTime, pointStartTime' are required.) End
	// time/last time point of the ephemeris, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	PointEndTime param.Opt[time.Time] `query:"pointEndTime,omitzero" format:"date-time" json:"-"`
	// (One or more of fields 'pointEndTime, pointStartTime' are required.) Start
	// time/first time point of the ephemeris, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	PointStartTime param.Opt[time.Time] `query:"pointStartTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EphemerisSetHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r EphemerisSetHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EphemerisSetHistoryCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'pointEndTime, pointStartTime' are required.) End
	// time/last time point of the ephemeris, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	PointEndTime param.Opt[time.Time] `query:"pointEndTime,omitzero" format:"date-time" json:"-"`
	// (One or more of fields 'pointEndTime, pointStartTime' are required.) Start
	// time/first time point of the ephemeris, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	PointStartTime param.Opt[time.Time] `query:"pointStartTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EphemerisSetHistoryCountParams]'s query parameters as
// `url.Values`.
func (r EphemerisSetHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
