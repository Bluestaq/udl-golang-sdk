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

// CloselyspacedobjectHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloselyspacedobjectHistoryService] method instead.
type CloselyspacedobjectHistoryService struct {
	Options []option.RequestOption
}

// NewCloselyspacedobjectHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloselyspacedobjectHistoryService(opts ...option.RequestOption) (r CloselyspacedobjectHistoryService) {
	r = CloselyspacedobjectHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *CloselyspacedobjectHistoryService) List(ctx context.Context, query CloselyspacedobjectHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[CloselyspacedobjectHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/closelyspacedobjects/history"
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
func (r *CloselyspacedobjectHistoryService) ListAutoPaging(ctx context.Context, query CloselyspacedobjectHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[CloselyspacedobjectHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *CloselyspacedobjectHistoryService) Aodr(ctx context.Context, query CloselyspacedobjectHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/closelyspacedobjects/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *CloselyspacedobjectHistoryService) Count(ctx context.Context, query CloselyspacedobjectHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/closelyspacedobjects/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This collection of services provides operations for manipulating and querying of
// closely spaced objects (on orbit) operations including docking, rendezvous,
// proximity and reporting of payload zone engagements observed and characterized
// over a period of time.
type CloselyspacedobjectHistoryListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicates the current state that characterizes Closely Spaced Objects (CSO)
	// analysis steps and conclusions. Values include: ACTIVE, ACTUAL, CANCELED,
	// CLOSED, COMPLETED, DETECTED, INDICATED, PENDING, PLANNED, POSSIBLE, PREDICTED,
	// SEPARATED, UPDATED.
	CsoState string `json:"csoState,required"`
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
	DataMode CloselyspacedobjectHistoryListResponseDataMode `json:"dataMode,required"`
	// Timestamp representing the events start time in ISO 8601 UTC format with
	// millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// The type of event associated with this record. Values include: DOCK, UNDOCK,
	// SEPARATION, RENDEZVOUS, PROXIMITY, PEZ, WEZ.
	EventType string `json:"eventType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// State vector epoch time of the actor satellite in ISO 8601 UTC format with
	// millisecond precision.
	ActorSvEpoch time.Time `json:"actorSVEpoch" format:"date-time"`
	// Timespan of the rendezvous analysis in seconds.
	AnalysisDuration float64 `json:"analysisDuration"`
	// Epoch time of the beginning of the analysis period in ISO 8601 UTC format with
	// millisecond precision.
	AnalysisEpoch time.Time `json:"analysisEpoch" format:"date-time"`
	// Computation type, values (e.g. PLANARALIGNMENT, LONGITUDE).
	CompType string `json:"compType"`
	// An optional string array containing additional data (keys) representing relevant
	// items for context of fields not specifically defined in this schema. This array
	// is paired with the contextValues string array and must contain the same number
	// of items. Please note these fields are intended for contextual use only and do
	// not pertain to core schema information. To ensure proper integration and avoid
	// misuse, coordination of how these fields are populated and consumed is required
	// during onboarding.
	ContextKeys []string `json:"contextKeys"`
	// An optional string array containing the values associated with the contextKeys
	// array. This array is paired with the contextKeys string array and must contain
	// the same number of items. Please note these fields are intended for contextual
	// use only and do not pertain to core schema information. To ensure proper
	// integration and avoid misuse, coordination of how these fields are populated and
	// consumed is required during onboarding.
	ContextValues []string `json:"contextValues"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A collection of orbital metrics for the event at the start and end times, and
	// the mean values of the primary and secondary objects, as well as the deltas
	// between the primary and secondary objects.
	CsoDetails []CloselyspacedobjectHistoryListResponseCsoDetail `json:"csoDetails"`
	// The tolerance value for the DeltaV, in kilometers per second.
	DeltaVTol float64 `json:"deltaVTol"`
	// The threshold of the event duration in seconds.
	DurationThreshold float64 `json:"durationThreshold"`
	// Timestamp representing the events end time in ISO 8601 UTC format with
	// millisecond precision.
	EventEndTime time.Time `json:"eventEndTime" format:"date-time"`
	// Percentage of the event interval that is within the plane tolerance specified as
	// a percent value between 0 and 100.
	EventIntervalCoverage float64 `json:"eventIntervalCoverage"`
	// Unique identifier of the record from the originating system. This field has no
	// meaning within UDL and is provided as a convenience for systems that require
	// tracking of an internal system generated ID.
	ExtID string `json:"extId"`
	// The Hohmann DeltaV (kilometers per second) is the minimum delta velocity for the
	// in-plane orbit change. The in-plane maneuvers change the semi-major axis
	// (perigee and/or apogee). It is the minimum assuming two maneuvers; a lower delta
	// velocity is possible with bi-elliptic transfers involving three maneuvers.
	HohmannDeltaV float64 `json:"hohmannDeltaV"`
	// Optional ID of the UDL State Vector at epoch time of the actor satellite. When
	// performing a create, this id will be ignored in favor of the UDL generated id of
	// the actor state vector.
	IDActorSv string `json:"idActorSV"`
	// Unique identifier of the primary satellite on-orbit object, if correlated. For
	// rendezvous and proximity operations, this is the target on-orbit object. When
	// the secondary object is on the rendezvous capable list, this can be any object.
	IDOnOrbit1 string `json:"idOnOrbit1"`
	// Unique identifier of the secondary satellite on-orbit object, if correlated. For
	// rendezvous and proximity operations, this is the actor. When the primary object
	// is a satellite being protected on the neighborhood watch list (NWL), this can be
	// any object encroaching on the primary.
	IDOnOrbit2 string `json:"idOnOrbit2"`
	// Optional ID of the UDL State Vector at epoch time of the target satellite. When
	// performing a create, this id will be ignored in favor of the UDL generated id of
	// the target state vector.
	IDTargetSv string `json:"idTargetSV"`
	// The Inclination DeltaV is the minimum delta velocity for the out-of-plane
	// change, assuming alignment of the right ascensions measured in kilometers per
	// second.
	InclinationDeltaV float64 `json:"inclinationDeltaV"`
	// Identifies the source of the indication, if the latest event info was manually
	// input, not computed.
	IndicationSource string `json:"indicationSource"`
	// The tolerance value for the longitude in degrees.
	LonTol float64 `json:"lonTol"`
	// Maximum range (apogee and perigee differences) within the event interval
	// measured in kilometers.
	MaxRange float64 `json:"maxRange"`
	// Minimum angle between the target's position and the projection of the actor's
	// position into the target's nominal orbit plane over the event interval measured
	// in degrees.
	MinPlaneSepAngle float64 `json:"minPlaneSepAngle"`
	// Epoch time of the minimum in-plane separation angle occurrence in ISO 8601 UTC
	// format with millisecond precision.
	MinPlaneSepEpoch time.Time `json:"minPlaneSepEpoch" format:"date-time"`
	// Minimum range (apogee and perigee differences) within the event interval
	// measured in kilometers.
	MinRange float64 `json:"minRange"`
	// Timespan of satellites within the range tolerance in seconds.
	MinRangeAnalysisDuration float64 `json:"minRangeAnalysisDuration"`
	// Epoch time of the minimum range occurrence in ISO 8601 UTC format with
	// millisecond precision.
	MinRangeEpoch time.Time `json:"minRangeEpoch" format:"date-time"`
	// Contains other descriptive information associated with an indicated event.
	Notes string `json:"notes"`
	// The number of oscillations within the event interval which are within the plane
	// tolerance.
	NumSubIntervals int64 `json:"numSubIntervals"`
	// The change in angle between the angular momentum vectors between the actor and
	// target relative to plane orientation in degrees.
	OrbitAlignDel float64 `json:"orbitAlignDel"`
	// The tolerance value for the difference in the orbital plane measured in degrees.
	OrbitPlaneTol float64 `json:"orbitPlaneTol"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional place holder for an OnOrbit ID that does not exist in UDL.
	OrigObjectId1 string `json:"origObjectId1"`
	// Optional place holder for an OnOrbit ID that does not exist in UDL.
	OrigObjectId2 string `json:"origObjectId2"`
	// The threshold of the event range in kilometers.
	RangeThreshold float64 `json:"rangeThreshold"`
	// The tolerance value for the range in kilometers.
	RangeTol float64 `json:"rangeTol"`
	// Indicates the relative position vector of the event occurrence measured in
	// kilometers.
	RelPos []float64 `json:"relPos"`
	// Range of closest approach: relative position magnitude, in kilometers, of the
	// difference in the physical position between the actor and target objects.
	RelPosMag float64 `json:"relPosMag"`
	// Indicates the closure rate specified as a relative velocity magnitude in
	// kilometers per second of the difference in the velocities between the actor and
	// target objects.
	RelSpeedMag float64 `json:"relSpeedMag"`
	// Indicates the relative velocity vector of the event occurrence measured in
	// kilometers per second.
	RelVel []float64 `json:"relVel"`
	// Satellite/catalog number of the target on-orbit primary object.
	SatNo1 int64 `json:"satNo1"`
	// Satellite/catalog number of the target on-orbit secondary object.
	SatNo2 int64 `json:"satNo2"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The tolerance value for the optimal longitude for station-keeping in degrees.
	StationLimLonTol float64 `json:"stationLimLonTol"`
	// State vector epoch time of the target satellite in ISO 8601 UTC format with
	// millisecond precision.
	TargetSvEpoch time.Time `json:"targetSVEpoch" format:"date-time"`
	// The Total DeltaV is the sum of the Hohmann and Inclination DeltaVs measured in
	// kilometers per second.
	TotalDeltaV float64 `json:"totalDeltaV"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		CsoState                 respjson.Field
		DataMode                 respjson.Field
		EventStartTime           respjson.Field
		EventType                respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		ActorSvEpoch             respjson.Field
		AnalysisDuration         respjson.Field
		AnalysisEpoch            respjson.Field
		CompType                 respjson.Field
		ContextKeys              respjson.Field
		ContextValues            respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		CsoDetails               respjson.Field
		DeltaVTol                respjson.Field
		DurationThreshold        respjson.Field
		EventEndTime             respjson.Field
		EventIntervalCoverage    respjson.Field
		ExtID                    respjson.Field
		HohmannDeltaV            respjson.Field
		IDActorSv                respjson.Field
		IDOnOrbit1               respjson.Field
		IDOnOrbit2               respjson.Field
		IDTargetSv               respjson.Field
		InclinationDeltaV        respjson.Field
		IndicationSource         respjson.Field
		LonTol                   respjson.Field
		MaxRange                 respjson.Field
		MinPlaneSepAngle         respjson.Field
		MinPlaneSepEpoch         respjson.Field
		MinRange                 respjson.Field
		MinRangeAnalysisDuration respjson.Field
		MinRangeEpoch            respjson.Field
		Notes                    respjson.Field
		NumSubIntervals          respjson.Field
		OrbitAlignDel            respjson.Field
		OrbitPlaneTol            respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		OrigObjectId1            respjson.Field
		OrigObjectId2            respjson.Field
		RangeThreshold           respjson.Field
		RangeTol                 respjson.Field
		RelPos                   respjson.Field
		RelPosMag                respjson.Field
		RelSpeedMag              respjson.Field
		RelVel                   respjson.Field
		SatNo1                   respjson.Field
		SatNo2                   respjson.Field
		SourceDl                 respjson.Field
		StationLimLonTol         respjson.Field
		TargetSvEpoch            respjson.Field
		TotalDeltaV              respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloselyspacedobjectHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *CloselyspacedobjectHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type CloselyspacedobjectHistoryListResponseDataMode string

const (
	CloselyspacedobjectHistoryListResponseDataModeReal      CloselyspacedobjectHistoryListResponseDataMode = "REAL"
	CloselyspacedobjectHistoryListResponseDataModeTest      CloselyspacedobjectHistoryListResponseDataMode = "TEST"
	CloselyspacedobjectHistoryListResponseDataModeSimulated CloselyspacedobjectHistoryListResponseDataMode = "SIMULATED"
	CloselyspacedobjectHistoryListResponseDataModeExercise  CloselyspacedobjectHistoryListResponseDataMode = "EXERCISE"
)

// A collection of orbital metrics for the event at the start and end times, and
// the mean values of the primary and secondary objects, as well as the deltas
// between the primary and secondary objects.
type CloselyspacedobjectHistoryListResponseCsoDetail struct {
	// The type of object event the metrics apply. Values consist of START, END, MEAN.
	ObjectEvent string `json:"objectEvent,required"`
	// The type of the object for which the metrics apply. Values include PRIMARY,
	// SECONDARY, DELTA.
	ObjectType string `json:"objectType,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The orbit point furthest from the center of the earth in kilometers.
	Apogee float64 `json:"apogee"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking"`
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
	DataMode string `json:"dataMode"`
	// Unique identifier of the parent CSO record, auto-populated by the system.
	IDCso string `json:"idCSO"`
	// The angle between the equator and the orbit when looking from the center of the
	// Earth in degrees from 0 to 180.
	Inclination float64 `json:"inclination"`
	// The longitude degree of the object measured in degrees from -180 to 180.
	Longitude float64 `json:"longitude"`
	// The orbit point nearest to the center of the earth in kilometers.
	Perigee float64 `json:"perigee"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectEvent           respjson.Field
		ObjectType            respjson.Field
		ID                    respjson.Field
		Apogee                respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDCso                 respjson.Field
		Inclination           respjson.Field
		Longitude             respjson.Field
		Perigee               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloselyspacedobjectHistoryListResponseCsoDetail) RawJSON() string { return r.JSON.raw }
func (r *CloselyspacedobjectHistoryListResponseCsoDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CloselyspacedobjectHistoryListParams struct {
	// Timestamp representing the events start time in ISO 8601 UTC format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EventStartTime time.Time `query:"eventStartTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CloselyspacedobjectHistoryListParams]'s query parameters as
// `url.Values`.
func (r CloselyspacedobjectHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloselyspacedobjectHistoryAodrParams struct {
	// Timestamp representing the events start time in ISO 8601 UTC format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
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

// URLQuery serializes [CloselyspacedobjectHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r CloselyspacedobjectHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloselyspacedobjectHistoryCountParams struct {
	// Timestamp representing the events start time in ISO 8601 UTC format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EventStartTime time.Time        `query:"eventStartTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CloselyspacedobjectHistoryCountParams]'s query parameters
// as `url.Values`.
func (r CloselyspacedobjectHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
