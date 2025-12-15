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

// LogisticsSupportHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogisticsSupportHistoryService] method instead.
type LogisticsSupportHistoryService struct {
	Options []option.RequestOption
}

// NewLogisticsSupportHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogisticsSupportHistoryService(opts ...option.RequestOption) (r LogisticsSupportHistoryService) {
	r = LogisticsSupportHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LogisticsSupportHistoryService) List(ctx context.Context, query LogisticsSupportHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LogisticsSupportHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/logisticssupport/history"
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
func (r *LogisticsSupportHistoryService) ListAutoPaging(ctx context.Context, query LogisticsSupportHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LogisticsSupportHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LogisticsSupportHistoryService) Aodr(ctx context.Context, query LogisticsSupportHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/logisticssupport/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LogisticsSupportHistoryService) Count(ctx context.Context, query LogisticsSupportHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/logisticssupport/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
type LogisticsSupportHistoryListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode LogisticsSupportHistoryListResponseDataMode `json:"dataMode,required"`
	// The time this report was created, in ISO 8601 UTC format with millisecond
	// precision.
	RptCreatedTime time.Time `json:"rptCreatedTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds string `json:"aircraftMDS"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// The current ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	CurrIcao string `json:"currICAO"`
	// The estimated time mission capable for the aircraft, in ISO 8601 UCT format with
	// millisecond precision. This is the estimated time when the aircraft is mission
	// ready.
	Etic time.Time `json:"etic" format:"date-time"`
	// Logistics estimated time mission capable.
	Etmc time.Time `json:"etmc" format:"date-time"`
	// Optional system identifier from external systs. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtSystemID string `json:"extSystemId"`
	// This field identifies the pacing event for bringing the aircraft to Mission
	// Capable status. It is used in calculating the Estimated Time Mission Capable
	// (ETMC) value. Acceptable values are WA (Will Advise), INW (In Work), P+hhh.h
	// (where P=parts and hhh.h is the number of hours up to 999 plus tenths of hours),
	// EQ+hhh.h (EQ=equipment), MRT+hhh.h (MRT=maintenance recovery team).
	LogisticAction string `json:"logisticAction"`
	// Discrepancy information associated with this LogisticsSupport record.
	LogisticsDiscrepancyInfos []LogisticsDiscrepancyInfosFull `json:"logisticsDiscrepancyInfos"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID string `json:"logisticsRecordId"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticsSupportItemsFull `json:"logisticsSupportItems"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticsTransportationPlansFull `json:"logisticsTransportationPlans"`
	// The maintenance status code of the aircraft which may be based on pilot
	// descriptions or evaluation codes. Contact the source provider for details.
	MaintStatusCode string `json:"maintStatusCode"`
	// The time indicating when all mission essential problems with a given aircraft
	// have been fixed and is mission capable. This datetime should be in ISO 8601 UTC
	// format with millisecond precision.
	McTime time.Time `json:"mcTime" format:"date-time"`
	// The time indicating when a given aircraft breaks for a mission essential reason.
	// This datetime should be in ISO 8601 UTC format with millisecond precision.
	MeTime time.Time `json:"meTime" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The organization that owns this logistics record.
	Owner string `json:"owner"`
	// This is used to indicate whether a closed master record has been reopened.
	ReopenFlag bool `json:"reopenFlag"`
	// The time this report was closed, in ISO 8601 UTC format with millisecond
	// precision.
	RptClosedTime time.Time `json:"rptClosedTime" format:"date-time"`
	// The supplying ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	SuppIcao string `json:"suppICAO"`
	// The tail number of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	TailNumber string `json:"tailNumber"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking        respjson.Field
		DataMode                     respjson.Field
		RptCreatedTime               respjson.Field
		Source                       respjson.Field
		ID                           respjson.Field
		AircraftMds                  respjson.Field
		CreatedAt                    respjson.Field
		CreatedBy                    respjson.Field
		CurrIcao                     respjson.Field
		Etic                         respjson.Field
		Etmc                         respjson.Field
		ExtSystemID                  respjson.Field
		LogisticAction               respjson.Field
		LogisticsDiscrepancyInfos    respjson.Field
		LogisticsRecordID            respjson.Field
		LogisticsRemarks             respjson.Field
		LogisticsSupportItems        respjson.Field
		LogisticsTransportationPlans respjson.Field
		MaintStatusCode              respjson.Field
		McTime                       respjson.Field
		MeTime                       respjson.Field
		Origin                       respjson.Field
		OrigNetwork                  respjson.Field
		Owner                        respjson.Field
		ReopenFlag                   respjson.Field
		RptClosedTime                respjson.Field
		SuppIcao                     respjson.Field
		TailNumber                   respjson.Field
		UpdatedAt                    respjson.Field
		UpdatedBy                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportHistoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type LogisticsSupportHistoryListResponseDataMode string

const (
	LogisticsSupportHistoryListResponseDataModeReal      LogisticsSupportHistoryListResponseDataMode = "REAL"
	LogisticsSupportHistoryListResponseDataModeTest      LogisticsSupportHistoryListResponseDataMode = "TEST"
	LogisticsSupportHistoryListResponseDataModeSimulated LogisticsSupportHistoryListResponseDataMode = "SIMULATED"
	LogisticsSupportHistoryListResponseDataModeExercise  LogisticsSupportHistoryListResponseDataMode = "EXERCISE"
)

type LogisticsSupportHistoryListParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogisticsSupportHistoryListParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticsSupportHistoryAodrParams struct {
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

// URLQuery serializes [LogisticsSupportHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticsSupportHistoryCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogisticsSupportHistoryCountParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
