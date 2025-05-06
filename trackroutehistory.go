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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// TrackRouteHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrackRouteHistoryService] method instead.
type TrackRouteHistoryService struct {
	Options []option.RequestOption
}

// NewTrackRouteHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTrackRouteHistoryService(opts ...option.RequestOption) (r TrackRouteHistoryService) {
	r = TrackRouteHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TrackRouteHistoryService) List(ctx context.Context, query TrackRouteHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[TrackRouteFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/trackroute/history"
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
func (r *TrackRouteHistoryService) ListAutoPaging(ctx context.Context, query TrackRouteHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[TrackRouteFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TrackRouteHistoryService) Aodr(ctx context.Context, query TrackRouteHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/trackroute/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *TrackRouteHistoryService) Count(ctx context.Context, query TrackRouteHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/trackroute/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A track route is a prescribed route for performing training events or operations
// such as air refueling.
type TrackRouteFull struct {
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
	DataMode TrackRouteFullDataMode `json:"dataMode,required"`
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision.
	LastUpdateDate time.Time `json:"lastUpdateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The track route type represented by this record (ex. AIR REFUELING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Minimum and maximum altitude bounds for the track.
	AltitudeBlocks []TrackRouteFullAltitudeBlock `json:"altitudeBlocks"`
	// The APN radar code sent and received by the aircraft for identification.
	ApnSetting string `json:"apnSetting"`
	// The APX radar code sent and received by the aircraft for identification.
	ApxBeaconCode string `json:"apxBeaconCode"`
	// Air Refueling Track Control Center message.
	ArtccMessage string `json:"artccMessage"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The name of the creating organization of the track route.
	CreatingOrg string `json:"creatingOrg"`
	// The principal compass direction (cardinal or ordinal) of the track route.
	Direction string `json:"direction"`
	// The date which the DAFIF track was last updated/validated in ISO 8601 UTC format
	// with millisecond precision.
	EffectiveDate time.Time `json:"effectiveDate" format:"date-time"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalID string `json:"externalId"`
	// Used to show last time the track route was added to an itinerary in ISO 8601 UTC
	// format with millisecond precision.
	LastUsedDate time.Time `json:"lastUsedDate" format:"date-time"`
	// Track location ID.
	LocationTrackID string `json:"locationTrackId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Point of contacts for scheduling or modifying the route.
	Poc []TrackRouteFullPoc `json:"poc"`
	// The primary UHF radio frequency used for the track route in megahertz.
	PriFreq float64 `json:"priFreq"`
	// The receiver tanker channel identifer for air refueling tracks.
	ReceiverTankerChCode string `json:"receiverTankerCHCode"`
	// Region code indicating where the track resides as determined by the data source.
	RegionCode string `json:"regionCode"`
	// Region where the track resides.
	RegionName string `json:"regionName"`
	// Date the track needs to be reviewed for accuracy or deletion in ISO 8601 UTC
	// format with millisecond precision.
	ReviewDate time.Time `json:"reviewDate" format:"date-time"`
	// Points identified within the route.
	RoutePoints []TrackRouteFullRoutePoint `json:"routePoints"`
	// Point of contact for the air refueling track route scheduler.
	SchedulerOrgName string `json:"schedulerOrgName"`
	// The unit responsible for scheduling the track route.
	SchedulerOrgUnit string `json:"schedulerOrgUnit"`
	// The secondary UHF radio frequency used for the track route in megahertz.
	SecFreq float64 `json:"secFreq"`
	// Abbreviated name of the track.
	ShortName string `json:"shortName"`
	// Standard Indicator Code of the air refueling track.
	Sic string `json:"sic"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Identifier of the track.
	TrackID string `json:"trackId"`
	// Name of the track.
	TrackName string `json:"trackName"`
	// Type of process used by AMC to schedule an air refueling event. Possible values
	// are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched Theater
	// Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S (Soft Air
	// Refueling), T (Matched Theater Operation Short Notice (Theater Assets)), V
	// (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short Notice (AMC
	// Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z (Other Air
	// Refueling).
	TypeCode string `json:"typeCode"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		LastUpdateDate        respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AltitudeBlocks        respjson.Field
		ApnSetting            respjson.Field
		ApxBeaconCode         respjson.Field
		ArtccMessage          respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CreatingOrg           respjson.Field
		Direction             respjson.Field
		EffectiveDate         respjson.Field
		ExternalID            respjson.Field
		LastUsedDate          respjson.Field
		LocationTrackID       respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Poc                   respjson.Field
		PriFreq               respjson.Field
		ReceiverTankerChCode  respjson.Field
		RegionCode            respjson.Field
		RegionName            respjson.Field
		ReviewDate            respjson.Field
		RoutePoints           respjson.Field
		SchedulerOrgName      respjson.Field
		SchedulerOrgUnit      respjson.Field
		SecFreq               respjson.Field
		ShortName             respjson.Field
		Sic                   respjson.Field
		SourceDl              respjson.Field
		TrackID               respjson.Field
		TrackName             respjson.Field
		TypeCode              respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackRouteFull) RawJSON() string { return r.JSON.raw }
func (r *TrackRouteFull) UnmarshalJSON(data []byte) error {
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
type TrackRouteFullDataMode string

const (
	TrackRouteFullDataModeReal      TrackRouteFullDataMode = "REAL"
	TrackRouteFullDataModeTest      TrackRouteFullDataMode = "TEST"
	TrackRouteFullDataModeSimulated TrackRouteFullDataMode = "SIMULATED"
	TrackRouteFullDataModeExercise  TrackRouteFullDataMode = "EXERCISE"
)

// Minimum and maximum altitude bounds for the track.
type TrackRouteFullAltitudeBlock struct {
	// Sequencing field for the altitude block.
	AltitudeSequenceID string `json:"altitudeSequenceId"`
	// Lowest altitude of the track route altitude block above mean sea level in feet.
	LowerAltitude float64 `json:"lowerAltitude"`
	// Highest altitude of the track route altitude block above mean sea level in feet.
	UpperAltitude float64 `json:"upperAltitude"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltitudeSequenceID respjson.Field
		LowerAltitude      respjson.Field
		UpperAltitude      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackRouteFullAltitudeBlock) RawJSON() string { return r.JSON.raw }
func (r *TrackRouteFullAltitudeBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Point of contacts for scheduling or modifying the route.
type TrackRouteFullPoc struct {
	// Office name for which the contact belongs.
	Office string `json:"office"`
	// Phone number of the contact.
	Phone string `json:"phone"`
	// The name of the contact.
	PocName string `json:"pocName"`
	// Organization name for which the contact belongs.
	PocOrg string `json:"pocOrg"`
	// Sequencing field for point of contact.
	PocSequenceID int64 `json:"pocSequenceId"`
	// A code or name that represents the contact's role in association to the track
	// route (ex. Originator, Scheduler, Maintainer, etc.).
	PocTypeName string `json:"pocTypeName"`
	// The rank of contact.
	Rank string `json:"rank"`
	// Text of the remark.
	Remark string `json:"remark"`
	// The username of the contact.
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Office        respjson.Field
		Phone         respjson.Field
		PocName       respjson.Field
		PocOrg        respjson.Field
		PocSequenceID respjson.Field
		PocTypeName   respjson.Field
		Rank          respjson.Field
		Remark        respjson.Field
		Username      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackRouteFullPoc) RawJSON() string { return r.JSON.raw }
func (r *TrackRouteFullPoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Points identified within the route.
type TrackRouteFullRoutePoint struct {
	// Specifies an alternate country code if the data provider code is not part of an
	// official NAVAID Country Code standard such as ISO-3166 or FIPS. This field will
	// be set to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode string `json:"altCountryCode"`
	// The DoD Standard Country Code designator for the country where the route point
	// resides. This field should be set to "OTHR" if the source value does not match a
	// UDL country code value (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode"`
	// Flag indicating this is a Digital Aeronautical Flight Information File (DAFIF)
	// point.
	DafifPt bool `json:"dafifPt"`
	// The magnetic declination/variation of the route point location from true north,
	// in degrees. Positive values east of true north and negative values west of true
	// north.
	MagDec float64 `json:"magDec"`
	// Navigational Aid (NAVAID) identification code.
	Navaid string `json:"navaid"`
	// The length of the course from the Navigational Aid (NAVAID) in nautical miles.
	NavaidLength float64 `json:"navaidLength"`
	// The NAVAID type of this route point (ex. VOR, VORTAC, TACAN, etc.).
	NavaidType string `json:"navaidType"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PtLat float64 `json:"ptLat"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	PtLon float64 `json:"ptLon"`
	// Sequencing field for the track route. This is the identifier representing the
	// sequence of waypoints associated to the track route.
	PtSequenceID int64 `json:"ptSequenceId"`
	// Code representation of the point within the track route (ex. EP, EX, CP, IP,
	// etc.).
	PtTypeCode string `json:"ptTypeCode"`
	// The name that represents the point within the track route (ex. ENTRY POINT, EXIT
	// POINT, CONTROL POINT, INITIAL POINT, etc.).
	PtTypeName string `json:"ptTypeName"`
	// Name of a waypoint which identifies the location of the point.
	WaypointName string `json:"waypointName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltCountryCode respjson.Field
		CountryCode    respjson.Field
		DafifPt        respjson.Field
		MagDec         respjson.Field
		Navaid         respjson.Field
		NavaidLength   respjson.Field
		NavaidType     respjson.Field
		PtLat          respjson.Field
		PtLon          respjson.Field
		PtSequenceID   respjson.Field
		PtTypeCode     respjson.Field
		PtTypeName     respjson.Field
		WaypointName   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackRouteFullRoutePoint) RawJSON() string { return r.JSON.raw }
func (r *TrackRouteFullRoutePoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackRouteHistoryListParams struct {
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	LastUpdateDate time.Time `query:"lastUpdateDate,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TrackRouteHistoryListParams]'s query parameters as
// `url.Values`.
func (r TrackRouteHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TrackRouteHistoryAodrParams struct {
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	LastUpdateDate time.Time `query:"lastUpdateDate,required" format:"date-time" json:"-"`
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

// URLQuery serializes [TrackRouteHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r TrackRouteHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TrackRouteHistoryCountParams struct {
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	LastUpdateDate time.Time        `query:"lastUpdateDate,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TrackRouteHistoryCountParams]'s query parameters as
// `url.Values`.
func (r TrackRouteHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
