// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// GlobalAtmosphericModelHistoryService contains methods and other services that
// help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalAtmosphericModelHistoryService] method instead.
type GlobalAtmosphericModelHistoryService struct {
	Options []option.RequestOption
}

// NewGlobalAtmosphericModelHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGlobalAtmosphericModelHistoryService(opts ...option.RequestOption) (r GlobalAtmosphericModelHistoryService) {
	r = GlobalAtmosphericModelHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *GlobalAtmosphericModelHistoryService) List(ctx context.Context, query GlobalAtmosphericModelHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GlobalAtmosphericModelHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/globalatmosphericmodel/history"
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
func (r *GlobalAtmosphericModelHistoryService) ListAutoPaging(ctx context.Context, query GlobalAtmosphericModelHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GlobalAtmosphericModelHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *GlobalAtmosphericModelHistoryService) Aodr(ctx context.Context, query GlobalAtmosphericModelHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/globalatmosphericmodel/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *GlobalAtmosphericModelHistoryService) Count(ctx context.Context, query GlobalAtmosphericModelHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/globalatmosphericmodel/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The GlobalAtmosphericModel service provides atmospheric model output data for
// use in space situational awareness such as the Global Total Electron Content
// (2D) data, Global Total Electron Density (3D) data, etc.
type GlobalAtmosphericModelHistoryListResponse struct {
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
	DataMode GlobalAtmosphericModelHistoryListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Type of data associated with this record (e.g. Global Total Electron Density,
	// Global Total Electron Content).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Model execution cadence, in minutes.
	Cadence int64 `json:"cadence"`
	// MD5 value of the data file. If not provided, the ingest/create operation will
	// automatically generate the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A unique identification code or label assigned to a particular source from which
	// atmospheric data originates.
	DataSourceIdentifier string `json:"dataSourceIdentifier"`
	// Ending altitude of model outputs, in kilometers.
	EndAlt float64 `json:"endAlt"`
	// WGS-84 ending latitude of model output, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndLat float64 `json:"endLat"`
	// WGS-84 ending longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndLon float64 `json:"endLon"`
	// The file name of the uploaded file.
	Filename string `json:"filename"`
	// The uploaded file size, in bytes. The maximum file size for this service is
	// 104857600 bytes (100MB). Files exceeding the maximum size will be rejected.
	Filesize int64 `json:"filesize"`
	// Number of altitude points.
	NumAlt int64 `json:"numAlt"`
	// Number of latitude points.
	NumLat int64 `json:"numLat"`
	// Number of longitude points.
	NumLon int64 `json:"numLon"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The time that this record was created, in ISO 8601 UTC format with millisecond
	// precision.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Starting altitude of model outputs, in kilometers.
	StartAlt float64 `json:"startAlt"`
	// WGS-84 starting latitude of model output, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	StartLat float64 `json:"startLat"`
	// WGS-84 starting longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	StartLon float64 `json:"startLon"`
	// State value indicating whether the values in this record are PREDICTED or
	// OBSERVED.
	State string `json:"state"`
	// Separation in latitude between subsequent model outputs, in degrees.
	StepLat float64 `json:"stepLat"`
	// Separation in longitude between subsequent model outputs, in degrees.
	StepLon float64 `json:"stepLon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		Cadence               respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataSourceIdentifier  respjson.Field
		EndAlt                respjson.Field
		EndLat                respjson.Field
		EndLon                respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		NumAlt                respjson.Field
		NumLat                respjson.Field
		NumLon                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ReportTime            respjson.Field
		SourceDl              respjson.Field
		StartAlt              respjson.Field
		StartLat              respjson.Field
		StartLon              respjson.Field
		State                 respjson.Field
		StepLat               respjson.Field
		StepLon               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalAtmosphericModelHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalAtmosphericModelHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type GlobalAtmosphericModelHistoryListResponseDataMode string

const (
	GlobalAtmosphericModelHistoryListResponseDataModeReal      GlobalAtmosphericModelHistoryListResponseDataMode = "REAL"
	GlobalAtmosphericModelHistoryListResponseDataModeTest      GlobalAtmosphericModelHistoryListResponseDataMode = "TEST"
	GlobalAtmosphericModelHistoryListResponseDataModeSimulated GlobalAtmosphericModelHistoryListResponseDataMode = "SIMULATED"
	GlobalAtmosphericModelHistoryListResponseDataModeExercise  GlobalAtmosphericModelHistoryListResponseDataMode = "EXERCISE"
)

type GlobalAtmosphericModelHistoryListParams struct {
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts time.Time `query:"ts,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalAtmosphericModelHistoryListParams]'s query parameters
// as `url.Values`.
func (r GlobalAtmosphericModelHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalAtmosphericModelHistoryAodrParams struct {
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts time.Time `query:"ts,required" format:"date-time" json:"-"`
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

// URLQuery serializes [GlobalAtmosphericModelHistoryAodrParams]'s query parameters
// as `url.Values`.
func (r GlobalAtmosphericModelHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalAtmosphericModelHistoryCountParams struct {
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalAtmosphericModelHistoryCountParams]'s query
// parameters as `url.Values`.
func (r GlobalAtmosphericModelHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
