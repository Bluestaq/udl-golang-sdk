// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// TaiutcHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaiutcHistoryService] method instead.
type TaiutcHistoryService struct {
	Options []option.RequestOption
}

// NewTaiutcHistoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTaiutcHistoryService(opts ...option.RequestOption) (r TaiutcHistoryService) {
	r = TaiutcHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TaiutcHistoryService) List(ctx context.Context, query TaiutcHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[TaiutcFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/taiutc/history"
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
func (r *TaiutcHistoryService) ListAutoPaging(ctx context.Context, query TaiutcHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[TaiutcFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TaiutcHistoryService) Aodr(ctx context.Context, query TaiutcHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/taiutc/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *TaiutcHistoryService) Count(ctx context.Context, query TaiutcHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/taiutc/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// International Atomic Time (TAI) is a statistical atomic time scale based on a
// large number of clocks operating at standards laboratories around the world that
// is maintained by the Bureau International des Poids et Mesures; its unit
// interval is exactly one SI second at sea level. The origin of TAI is such that
// UT1-TAI is approximately 0 (zero) on January 1, 1958. TAI is not adjusted for
// leap seconds. Coordinated Universal Time (UTC) is defined by the CCIR
// Recommendation 460-4 (1986). It differs from TAI by the total number of leap
// seconds, so that UT1-UTC stays smaller than 0.9s in absolute value. The decision
// to introduce a leap second in UTC is the responsibility of the International
// Earth Rotation Service (IERS). According to the CCIR Recommendation, first
// preference is given to the opportunities at the end of December and June, and
// second preference to those at the end of March and September. Since the system
// was introduced in 1972, only dates in June and December have been used. TAI is
// expressed in terms of UTC by the relation TAI = UTC + dAT, where dAT is the
// total algebraic sum of leap seconds. The first leap second was introduced on
// June 30, 1972. The historical list of leap seconds can be found in this table.
type TaiutcFull struct {
	// Effective date/time for the leap second adjustment.
	AdjustmentDate time.Time `json:"adjustmentDate,required" format:"date-time"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode TaiutcFullDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Multiplication factor of the leap second adjustment.
	MultiplicationFactor float64 `json:"multiplicationFactor"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Total/cumulative offset between TAI and UTC time as of adjustmentDate, in
	// seconds.
	TaiUtc float64 `json:"taiUTC"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AdjustmentDate        resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Source                resp.Field
		ID                    resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		MultiplicationFactor  resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		RawFileUri            resp.Field
		TaiUtc                resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaiutcFull) RawJSON() string { return r.JSON.raw }
func (r *TaiutcFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
type TaiutcFullDataMode string

const (
	TaiutcFullDataModeReal      TaiutcFullDataMode = "REAL"
	TaiutcFullDataModeTest      TaiutcFullDataMode = "TEST"
	TaiutcFullDataModeSimulated TaiutcFullDataMode = "SIMULATED"
	TaiutcFullDataModeExercise  TaiutcFullDataMode = "EXERCISE"
)

type TaiutcHistoryListParams struct {
	// Effective date/time for the leap second adjustment. Must be a unique value
	// across all TAIUTC datasets. (YYYY-MM-DDTHH:MM:SS.sssZ)
	AdjustmentDate time.Time `query:"adjustmentDate,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TaiutcHistoryListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TaiutcHistoryListParams]'s query parameters as
// `url.Values`.
func (r TaiutcHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TaiutcHistoryAodrParams struct {
	// Effective date/time for the leap second adjustment. Must be a unique value
	// across all TAIUTC datasets. (YYYY-MM-DDTHH:MM:SS.sssZ)
	AdjustmentDate time.Time `query:"adjustmentDate,required" format:"date-time" json:"-"`
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
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TaiutcHistoryAodrParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TaiutcHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r TaiutcHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TaiutcHistoryCountParams struct {
	// Effective date/time for the leap second adjustment. Must be a unique value
	// across all TAIUTC datasets. (YYYY-MM-DDTHH:MM:SS.sssZ)
	AdjustmentDate time.Time        `query:"adjustmentDate,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TaiutcHistoryCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TaiutcHistoryCountParams]'s query parameters as
// `url.Values`.
func (r TaiutcHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
