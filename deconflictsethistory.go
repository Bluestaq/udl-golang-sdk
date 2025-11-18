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

// DeconflictsetHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeconflictsetHistoryService] method instead.
type DeconflictsetHistoryService struct {
	Options []option.RequestOption
}

// NewDeconflictsetHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeconflictsetHistoryService(opts ...option.RequestOption) (r DeconflictsetHistoryService) {
	r = DeconflictsetHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *DeconflictsetHistoryService) List(ctx context.Context, query DeconflictsetHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[DeconflictsetHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/deconflictset/history"
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
func (r *DeconflictsetHistoryService) ListAutoPaging(ctx context.Context, query DeconflictsetHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[DeconflictsetHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *DeconflictsetHistoryService) Aodr(ctx context.Context, query DeconflictsetHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/deconflictset/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *DeconflictsetHistoryService) Count(ctx context.Context, query DeconflictsetHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/deconflictset/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The DeconflictSet service provides access to a set of DeconflictWindows and
// metadata about those data. A DeconflictWindow describes a time window during
// which an action, such as target engagement, may either occur or is prohibited
// from occurring. The DeconflictWindow model includes information about the
// spatial details for specific target types. A flag is provided to specify whether
// the window should be associated with taking action (OPEN), or if no action
// should occur (CLOSED).
type DeconflictsetHistoryListResponse struct {
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
	DataMode DeconflictsetHistoryListResponseDataMode `json:"dataMode,required"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// The number of windows provided by this DeconflictSet record.
	NumWindows int64 `json:"numWindows,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The time at which the window calculations completed, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationEndTime time.Time `json:"calculationEndTime" format:"date-time"`
	// The algorithm execution id associated with the generation of this DeconflictSet.
	CalculationID string `json:"calculationId"`
	// The time at which the window calculations started, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationStartTime time.Time `json:"calculationStartTime" format:"date-time"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array of DeconflictWindow records associated with this DeconflictSet.
	DeconflictWindows []DeconflictsetHistoryListResponseDeconflictWindow `json:"deconflictWindows"`
	// Array of error messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Errors []string `json:"errors"`
	// The end time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventEndTime time.Time `json:"eventEndTime" format:"date-time"`
	// The type of event associated with this DeconflictSet record.
	EventType string `json:"eventType"`
	// The id of the LaserDeconflictRequest record used as input in the generation of
	// this DeconflictSet, if applicable.
	IDLaserDeconflictRequest string `json:"idLaserDeconflictRequest"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	ReferenceFrame string `json:"referenceFrame"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Array of warning messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		EventStartTime           respjson.Field
		NumWindows               respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		CalculationEndTime       respjson.Field
		CalculationID            respjson.Field
		CalculationStartTime     respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		DeconflictWindows        respjson.Field
		Errors                   respjson.Field
		EventEndTime             respjson.Field
		EventType                respjson.Field
		IDLaserDeconflictRequest respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		ReferenceFrame           respjson.Field
		SourceDl                 respjson.Field
		Tags                     respjson.Field
		TransactionID            respjson.Field
		Warnings                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeconflictsetHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeconflictsetHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type DeconflictsetHistoryListResponseDataMode string

const (
	DeconflictsetHistoryListResponseDataModeReal      DeconflictsetHistoryListResponseDataMode = "REAL"
	DeconflictsetHistoryListResponseDataModeTest      DeconflictsetHistoryListResponseDataMode = "TEST"
	DeconflictsetHistoryListResponseDataModeSimulated DeconflictsetHistoryListResponseDataMode = "SIMULATED"
	DeconflictsetHistoryListResponseDataModeExercise  DeconflictsetHistoryListResponseDataMode = "EXERCISE"
)

// A DeconflictWindow describes a time window during which an action, such as
// target engagement, may either occur or is prohibited from occurring. The
// DeconflictWindow model includes information about the spatial details for
// specific target types. A flag is provided to specify whether the window should
// be associated with taking action (OPEN), or if no action should occur (CLOSED).
type DeconflictsetHistoryListResponseDeconflictWindow struct {
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
	DataMode string `json:"dataMode,required"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The window start time, in ISO 8601 UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The window stop time, in ISO 8601 UTC format with millisecond precision.
	StopTime time.Time `json:"stopTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The angle at which the victim enters the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfEntry float64 `json:"angleOfEntry"`
	// The angle at which the victim exits the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfExit float64 `json:"angleOfExit"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The X, Y, Z coordinates of entry, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	EntryCoords []float64 `json:"entryCoords"`
	// The type of event associated with the window status.
	EventType string `json:"eventType"`
	// The X, Y, Z coordinates of exit, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	ExitCoords []float64 `json:"exitCoords"`
	// Unique identifier of the parent DeconflictSet, auto-generated by the system. The
	// idDeconflictSet is used to identify all individual windows associated with a
	// parent DeconflictSet record.
	IDDeconflictSet string `json:"idDeconflictSet"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The target identifier. If the target is a satellite, the target is the
	// satellite/catalog number of the target on-orbit object.
	Target string `json:"target"`
	// The target type associated with this window (e.g. VICTIM, EARTH, etc.).
	TargetType string `json:"targetType"`
	// The victim identifier associated with this window. If the victim is a satellite,
	// the victim is the satellite/catalog number of the target on-orbit object.
	Victim string `json:"victim"`
	// The window status indicating whether possibility of action may occur. In other
	// words, OPEN is akin to a "green light," during which taking action is warranted
	// or authorized (though not necessarily required) over this timeframe, while
	// CLOSED represents a "red light," meaning that absolutely no action is warranted
	// or authorized to take place during this timeframe.
	WindowType string `json:"windowType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EventStartTime        respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		StopTime              respjson.Field
		ID                    respjson.Field
		AngleOfEntry          respjson.Field
		AngleOfExit           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EntryCoords           respjson.Field
		EventType             respjson.Field
		ExitCoords            respjson.Field
		IDDeconflictSet       respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SourceDl              respjson.Field
		Target                respjson.Field
		TargetType            respjson.Field
		Victim                respjson.Field
		WindowType            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeconflictsetHistoryListResponseDeconflictWindow) RawJSON() string { return r.JSON.raw }
func (r *DeconflictsetHistoryListResponseDeconflictWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeconflictsetHistoryListParams struct {
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EventStartTime time.Time `query:"eventStartTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeconflictsetHistoryListParams]'s query parameters as
// `url.Values`.
func (r DeconflictsetHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeconflictsetHistoryAodrParams struct {
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EventStartTime time.Time `query:"eventStartTime,required" format:"date-time" json:"-"`
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

// URLQuery serializes [DeconflictsetHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r DeconflictsetHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeconflictsetHistoryCountParams struct {
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EventStartTime time.Time        `query:"eventStartTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeconflictsetHistoryCountParams]'s query parameters as
// `url.Values`.
func (r DeconflictsetHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
