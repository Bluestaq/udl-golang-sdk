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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
)

// SgiHistoryService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSgiHistoryService] method instead.
type SgiHistoryService struct {
	Options []option.RequestOption
}

// NewSgiHistoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSgiHistoryService(opts ...option.RequestOption) (r SgiHistoryService) {
	r = SgiHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SgiHistoryService) List(ctx context.Context, query SgiHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SgiFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sgi/history"
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
func (r *SgiHistoryService) ListAutoPaging(ctx context.Context, query SgiHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SgiFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SgiHistoryService) Aodr(ctx context.Context, query SgiHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sgi/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SgiHistoryService) Count(ctx context.Context, query SgiHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sgi/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SgiHistoryListParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns param.Opt[string] `query:"columns,omitzero" json:"-"`
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// the data was received and processed from the source. Typically a source provides
	// solar data for a date window with each transmission including past, present, and
	// future predicted values. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EffectiveDate param.Opt[time.Time] `query:"effectiveDate,omitzero" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// of the index value. This could be a past, current, or future predicted value.
	// Note: sgiDate defines the start time of the time window for this data record.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	SgiDate param.Opt[time.Time] `query:"sgiDate,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SgiHistoryListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SgiHistoryListParams]'s query parameters as `url.Values`.
func (r SgiHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SgiHistoryAodrParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns param.Opt[string] `query:"columns,omitzero" json:"-"`
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// the data was received and processed from the source. Typically a source provides
	// solar data for a date window with each transmission including past, present, and
	// future predicted values. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EffectiveDate param.Opt[time.Time] `query:"effectiveDate,omitzero" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
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
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// of the index value. This could be a past, current, or future predicted value.
	// Note: sgiDate defines the start time of the time window for this data record.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	SgiDate param.Opt[time.Time] `query:"sgiDate,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SgiHistoryAodrParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SgiHistoryAodrParams]'s query parameters as `url.Values`.
func (r SgiHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SgiHistoryCountParams struct {
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// the data was received and processed from the source. Typically a source provides
	// solar data for a date window with each transmission including past, present, and
	// future predicted values. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EffectiveDate param.Opt[time.Time] `query:"effectiveDate,omitzero" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// of the index value. This could be a past, current, or future predicted value.
	// Note: sgiDate defines the start time of the time window for this data record.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	SgiDate param.Opt[time.Time] `query:"sgiDate,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SgiHistoryCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SgiHistoryCountParams]'s query parameters as `url.Values`.
func (r SgiHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
