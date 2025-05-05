// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

// TrackRouteService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrackRouteService] method instead.
type TrackRouteService struct {
	Options []option.RequestOption
	History TrackRouteHistoryService
}

// NewTrackRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTrackRouteService(opts ...option.RequestOption) (r TrackRouteService) {
	r = TrackRouteService{}
	r.Options = opts
	r.History = NewTrackRouteHistoryService(opts...)
	return
}

// Service operation to take a single trackroute record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *TrackRouteService) New(ctx context.Context, body TrackRouteNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/trackroute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single trackroute record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *TrackRouteService) Update(ctx context.Context, id string, body TrackRouteUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/trackroute/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TrackRouteService) List(ctx context.Context, query TrackRouteListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[TrackRouteListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/trackroute"
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

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TrackRouteService) ListAutoPaging(ctx context.Context, query TrackRouteListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[TrackRouteListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a trackroute record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *TrackRouteService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/trackroute/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *TrackRouteService) Count(ctx context.Context, query TrackRouteCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/trackroute/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// trackroute records as a POST body and ingest into the database. This operation
// is not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *TrackRouteService) NewBulk(ctx context.Context, body TrackRouteNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/trackroute/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single trackroute record by its unique ID passed as a
// path parameter.
func (r *TrackRouteService) Get(ctx context.Context, id string, query TrackRouteGetParams, opts ...option.RequestOption) (res *TrackRouteFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/trackroute/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *TrackRouteService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/trackroute/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Service operation to dynamically query data and only return specified
// columns/fields. Requested columns are specified by the 'columns' query parameter
// and should be a comma separated list of valid fields for the specified data
// type. classificationMarking is always returned. See the queryhelp operation
// (/udl/<datatype>/queryhelp) for more details on valid/required query parameter
// information. An example URI: /udl/elset/tuple?columns=satNo,period&epoch=>now-5
// hours would return the satNo and period of elsets with an epoch greater than 5
// hours ago.
func (r *TrackRouteService) Tuple(ctx context.Context, query TrackRouteTupleParams, opts ...option.RequestOption) (res *[]TrackRouteFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/trackroute/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple trackroute records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *TrackRouteService) UnvalidatedPublish(ctx context.Context, body TrackRouteUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-trackroute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// A track route is a prescribed route for performing training events or operations
// such as air refueling.
type TrackRouteListResponse struct {
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
	DataMode TrackRouteListResponseDataMode `json:"dataMode,required"`
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
	AltitudeBlocks []TrackRouteListResponseAltitudeBlock `json:"altitudeBlocks"`
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
	Poc []TrackRouteListResponsePoc `json:"poc"`
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
	RoutePoints []TrackRouteListResponseRoutePoint `json:"routePoints"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		LastUpdateDate        resp.Field
		Source                resp.Field
		Type                  resp.Field
		ID                    resp.Field
		AltitudeBlocks        resp.Field
		ApnSetting            resp.Field
		ApxBeaconCode         resp.Field
		ArtccMessage          resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		CreatingOrg           resp.Field
		Direction             resp.Field
		EffectiveDate         resp.Field
		ExternalID            resp.Field
		LastUsedDate          resp.Field
		LocationTrackID       resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Poc                   resp.Field
		PriFreq               resp.Field
		ReceiverTankerChCode  resp.Field
		RegionCode            resp.Field
		RegionName            resp.Field
		ReviewDate            resp.Field
		RoutePoints           resp.Field
		SchedulerOrgName      resp.Field
		SchedulerOrgUnit      resp.Field
		SecFreq               resp.Field
		ShortName             resp.Field
		Sic                   resp.Field
		SourceDl              resp.Field
		TrackID               resp.Field
		TrackName             resp.Field
		TypeCode              resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackRouteListResponse) RawJSON() string { return r.JSON.raw }
func (r *TrackRouteListResponse) UnmarshalJSON(data []byte) error {
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
type TrackRouteListResponseDataMode string

const (
	TrackRouteListResponseDataModeReal      TrackRouteListResponseDataMode = "REAL"
	TrackRouteListResponseDataModeTest      TrackRouteListResponseDataMode = "TEST"
	TrackRouteListResponseDataModeSimulated TrackRouteListResponseDataMode = "SIMULATED"
	TrackRouteListResponseDataModeExercise  TrackRouteListResponseDataMode = "EXERCISE"
)

// Minimum and maximum altitude bounds for the track.
type TrackRouteListResponseAltitudeBlock struct {
	// Sequencing field for the altitude block.
	AltitudeSequenceID string `json:"altitudeSequenceId"`
	// Lowest altitude of the track route altitude block above mean sea level in feet.
	LowerAltitude float64 `json:"lowerAltitude"`
	// Highest altitude of the track route altitude block above mean sea level in feet.
	UpperAltitude float64 `json:"upperAltitude"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AltitudeSequenceID resp.Field
		LowerAltitude      resp.Field
		UpperAltitude      resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackRouteListResponseAltitudeBlock) RawJSON() string { return r.JSON.raw }
func (r *TrackRouteListResponseAltitudeBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Point of contacts for scheduling or modifying the route.
type TrackRouteListResponsePoc struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Office        resp.Field
		Phone         resp.Field
		PocName       resp.Field
		PocOrg        resp.Field
		PocSequenceID resp.Field
		PocTypeName   resp.Field
		Rank          resp.Field
		Remark        resp.Field
		Username      resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackRouteListResponsePoc) RawJSON() string { return r.JSON.raw }
func (r *TrackRouteListResponsePoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Points identified within the route.
type TrackRouteListResponseRoutePoint struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AltCountryCode resp.Field
		CountryCode    resp.Field
		DafifPt        resp.Field
		MagDec         resp.Field
		Navaid         resp.Field
		NavaidLength   resp.Field
		NavaidType     resp.Field
		PtLat          resp.Field
		PtLon          resp.Field
		PtSequenceID   resp.Field
		PtTypeCode     resp.Field
		PtTypeName     resp.Field
		WaypointName   resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackRouteListResponseRoutePoint) RawJSON() string { return r.JSON.raw }
func (r *TrackRouteListResponseRoutePoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackRouteNewParams struct {
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
	DataMode TrackRouteNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision.
	LastUpdateDate time.Time `json:"lastUpdateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The track route type represented by this record (ex. AIR REFUELING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The APN radar code sent and received by the aircraft for identification.
	ApnSetting param.Opt[string] `json:"apnSetting,omitzero"`
	// The APX radar code sent and received by the aircraft for identification.
	ApxBeaconCode param.Opt[string] `json:"apxBeaconCode,omitzero"`
	// Air Refueling Track Control Center message.
	ArtccMessage param.Opt[string] `json:"artccMessage,omitzero"`
	// The name of the creating organization of the track route.
	CreatingOrg param.Opt[string] `json:"creatingOrg,omitzero"`
	// The principal compass direction (cardinal or ordinal) of the track route.
	Direction param.Opt[string] `json:"direction,omitzero"`
	// The date which the DAFIF track was last updated/validated in ISO 8601 UTC format
	// with millisecond precision.
	EffectiveDate param.Opt[time.Time] `json:"effectiveDate,omitzero" format:"date-time"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Used to show last time the track route was added to an itinerary in ISO 8601 UTC
	// format with millisecond precision.
	LastUsedDate param.Opt[time.Time] `json:"lastUsedDate,omitzero" format:"date-time"`
	// Track location ID.
	LocationTrackID param.Opt[string] `json:"locationTrackId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The primary UHF radio frequency used for the track route in megahertz.
	PriFreq param.Opt[float64] `json:"priFreq,omitzero"`
	// The receiver tanker channel identifer for air refueling tracks.
	ReceiverTankerChCode param.Opt[string] `json:"receiverTankerCHCode,omitzero"`
	// Region code indicating where the track resides as determined by the data source.
	RegionCode param.Opt[string] `json:"regionCode,omitzero"`
	// Region where the track resides.
	RegionName param.Opt[string] `json:"regionName,omitzero"`
	// Date the track needs to be reviewed for accuracy or deletion in ISO 8601 UTC
	// format with millisecond precision.
	ReviewDate param.Opt[time.Time] `json:"reviewDate,omitzero" format:"date-time"`
	// Point of contact for the air refueling track route scheduler.
	SchedulerOrgName param.Opt[string] `json:"schedulerOrgName,omitzero"`
	// The unit responsible for scheduling the track route.
	SchedulerOrgUnit param.Opt[string] `json:"schedulerOrgUnit,omitzero"`
	// The secondary UHF radio frequency used for the track route in megahertz.
	SecFreq param.Opt[float64] `json:"secFreq,omitzero"`
	// Abbreviated name of the track.
	ShortName param.Opt[string] `json:"shortName,omitzero"`
	// Standard Indicator Code of the air refueling track.
	Sic param.Opt[string] `json:"sic,omitzero"`
	// Identifier of the track.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Name of the track.
	TrackName param.Opt[string] `json:"trackName,omitzero"`
	// Type of process used by AMC to schedule an air refueling event. Possible values
	// are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched Theater
	// Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S (Soft Air
	// Refueling), T (Matched Theater Operation Short Notice (Theater Assets)), V
	// (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short Notice (AMC
	// Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z (Other Air
	// Refueling).
	TypeCode param.Opt[string] `json:"typeCode,omitzero"`
	// Minimum and maximum altitude bounds for the track.
	AltitudeBlocks []TrackRouteNewParamsAltitudeBlock `json:"altitudeBlocks,omitzero"`
	// Point of contacts for scheduling or modifying the route.
	Poc []TrackRouteNewParamsPoc `json:"poc,omitzero"`
	// Points identified within the route.
	RoutePoints []TrackRouteNewParamsRoutePoint `json:"routePoints,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r TrackRouteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type TrackRouteNewParamsDataMode string

const (
	TrackRouteNewParamsDataModeReal      TrackRouteNewParamsDataMode = "REAL"
	TrackRouteNewParamsDataModeTest      TrackRouteNewParamsDataMode = "TEST"
	TrackRouteNewParamsDataModeSimulated TrackRouteNewParamsDataMode = "SIMULATED"
	TrackRouteNewParamsDataModeExercise  TrackRouteNewParamsDataMode = "EXERCISE"
)

// Minimum and maximum altitude bounds for the track.
type TrackRouteNewParamsAltitudeBlock struct {
	// Sequencing field for the altitude block.
	AltitudeSequenceID param.Opt[string] `json:"altitudeSequenceId,omitzero"`
	// Lowest altitude of the track route altitude block above mean sea level in feet.
	LowerAltitude param.Opt[float64] `json:"lowerAltitude,omitzero"`
	// Highest altitude of the track route altitude block above mean sea level in feet.
	UpperAltitude param.Opt[float64] `json:"upperAltitude,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteNewParamsAltitudeBlock) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r TrackRouteNewParamsAltitudeBlock) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteNewParamsAltitudeBlock
	return param.MarshalObject(r, (*shadow)(&r))
}

// Point of contacts for scheduling or modifying the route.
type TrackRouteNewParamsPoc struct {
	// Office name for which the contact belongs.
	Office param.Opt[string] `json:"office,omitzero"`
	// Phone number of the contact.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// The name of the contact.
	PocName param.Opt[string] `json:"pocName,omitzero"`
	// Organization name for which the contact belongs.
	PocOrg param.Opt[string] `json:"pocOrg,omitzero"`
	// Sequencing field for point of contact.
	PocSequenceID param.Opt[int64] `json:"pocSequenceId,omitzero"`
	// A code or name that represents the contact's role in association to the track
	// route (ex. Originator, Scheduler, Maintainer, etc.).
	PocTypeName param.Opt[string] `json:"pocTypeName,omitzero"`
	// The rank of contact.
	Rank param.Opt[string] `json:"rank,omitzero"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The username of the contact.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteNewParamsPoc) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r TrackRouteNewParamsPoc) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteNewParamsPoc
	return param.MarshalObject(r, (*shadow)(&r))
}

// Points identified within the route.
type TrackRouteNewParamsRoutePoint struct {
	// Specifies an alternate country code if the data provider code is not part of an
	// official NAVAID Country Code standard such as ISO-3166 or FIPS. This field will
	// be set to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// The DoD Standard Country Code designator for the country where the route point
	// resides. This field should be set to "OTHR" if the source value does not match a
	// UDL country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Flag indicating this is a Digital Aeronautical Flight Information File (DAFIF)
	// point.
	DafifPt param.Opt[bool] `json:"dafifPt,omitzero"`
	// The magnetic declination/variation of the route point location from true north,
	// in degrees. Positive values east of true north and negative values west of true
	// north.
	MagDec param.Opt[float64] `json:"magDec,omitzero"`
	// Navigational Aid (NAVAID) identification code.
	Navaid param.Opt[string] `json:"navaid,omitzero"`
	// The length of the course from the Navigational Aid (NAVAID) in nautical miles.
	NavaidLength param.Opt[float64] `json:"navaidLength,omitzero"`
	// The NAVAID type of this route point (ex. VOR, VORTAC, TACAN, etc.).
	NavaidType param.Opt[string] `json:"navaidType,omitzero"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PtLat param.Opt[float64] `json:"ptLat,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	PtLon param.Opt[float64] `json:"ptLon,omitzero"`
	// Sequencing field for the track route. This is the identifier representing the
	// sequence of waypoints associated to the track route.
	PtSequenceID param.Opt[int64] `json:"ptSequenceId,omitzero"`
	// Code representation of the point within the track route (ex. EP, EX, CP, IP,
	// etc.).
	PtTypeCode param.Opt[string] `json:"ptTypeCode,omitzero"`
	// The name that represents the point within the track route (ex. ENTRY POINT, EXIT
	// POINT, CONTROL POINT, INITIAL POINT, etc.).
	PtTypeName param.Opt[string] `json:"ptTypeName,omitzero"`
	// Name of a waypoint which identifies the location of the point.
	WaypointName param.Opt[string] `json:"waypointName,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteNewParamsRoutePoint) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r TrackRouteNewParamsRoutePoint) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteNewParamsRoutePoint
	return param.MarshalObject(r, (*shadow)(&r))
}

type TrackRouteUpdateParams struct {
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
	DataMode TrackRouteUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision.
	LastUpdateDate time.Time `json:"lastUpdateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The track route type represented by this record (ex. AIR REFUELING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The APN radar code sent and received by the aircraft for identification.
	ApnSetting param.Opt[string] `json:"apnSetting,omitzero"`
	// The APX radar code sent and received by the aircraft for identification.
	ApxBeaconCode param.Opt[string] `json:"apxBeaconCode,omitzero"`
	// Air Refueling Track Control Center message.
	ArtccMessage param.Opt[string] `json:"artccMessage,omitzero"`
	// The name of the creating organization of the track route.
	CreatingOrg param.Opt[string] `json:"creatingOrg,omitzero"`
	// The principal compass direction (cardinal or ordinal) of the track route.
	Direction param.Opt[string] `json:"direction,omitzero"`
	// The date which the DAFIF track was last updated/validated in ISO 8601 UTC format
	// with millisecond precision.
	EffectiveDate param.Opt[time.Time] `json:"effectiveDate,omitzero" format:"date-time"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Used to show last time the track route was added to an itinerary in ISO 8601 UTC
	// format with millisecond precision.
	LastUsedDate param.Opt[time.Time] `json:"lastUsedDate,omitzero" format:"date-time"`
	// Track location ID.
	LocationTrackID param.Opt[string] `json:"locationTrackId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The primary UHF radio frequency used for the track route in megahertz.
	PriFreq param.Opt[float64] `json:"priFreq,omitzero"`
	// The receiver tanker channel identifer for air refueling tracks.
	ReceiverTankerChCode param.Opt[string] `json:"receiverTankerCHCode,omitzero"`
	// Region code indicating where the track resides as determined by the data source.
	RegionCode param.Opt[string] `json:"regionCode,omitzero"`
	// Region where the track resides.
	RegionName param.Opt[string] `json:"regionName,omitzero"`
	// Date the track needs to be reviewed for accuracy or deletion in ISO 8601 UTC
	// format with millisecond precision.
	ReviewDate param.Opt[time.Time] `json:"reviewDate,omitzero" format:"date-time"`
	// Point of contact for the air refueling track route scheduler.
	SchedulerOrgName param.Opt[string] `json:"schedulerOrgName,omitzero"`
	// The unit responsible for scheduling the track route.
	SchedulerOrgUnit param.Opt[string] `json:"schedulerOrgUnit,omitzero"`
	// The secondary UHF radio frequency used for the track route in megahertz.
	SecFreq param.Opt[float64] `json:"secFreq,omitzero"`
	// Abbreviated name of the track.
	ShortName param.Opt[string] `json:"shortName,omitzero"`
	// Standard Indicator Code of the air refueling track.
	Sic param.Opt[string] `json:"sic,omitzero"`
	// Identifier of the track.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Name of the track.
	TrackName param.Opt[string] `json:"trackName,omitzero"`
	// Type of process used by AMC to schedule an air refueling event. Possible values
	// are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched Theater
	// Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S (Soft Air
	// Refueling), T (Matched Theater Operation Short Notice (Theater Assets)), V
	// (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short Notice (AMC
	// Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z (Other Air
	// Refueling).
	TypeCode param.Opt[string] `json:"typeCode,omitzero"`
	// Minimum and maximum altitude bounds for the track.
	AltitudeBlocks []TrackRouteUpdateParamsAltitudeBlock `json:"altitudeBlocks,omitzero"`
	// Point of contacts for scheduling or modifying the route.
	Poc []TrackRouteUpdateParamsPoc `json:"poc,omitzero"`
	// Points identified within the route.
	RoutePoints []TrackRouteUpdateParamsRoutePoint `json:"routePoints,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r TrackRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type TrackRouteUpdateParamsDataMode string

const (
	TrackRouteUpdateParamsDataModeReal      TrackRouteUpdateParamsDataMode = "REAL"
	TrackRouteUpdateParamsDataModeTest      TrackRouteUpdateParamsDataMode = "TEST"
	TrackRouteUpdateParamsDataModeSimulated TrackRouteUpdateParamsDataMode = "SIMULATED"
	TrackRouteUpdateParamsDataModeExercise  TrackRouteUpdateParamsDataMode = "EXERCISE"
)

// Minimum and maximum altitude bounds for the track.
type TrackRouteUpdateParamsAltitudeBlock struct {
	// Sequencing field for the altitude block.
	AltitudeSequenceID param.Opt[string] `json:"altitudeSequenceId,omitzero"`
	// Lowest altitude of the track route altitude block above mean sea level in feet.
	LowerAltitude param.Opt[float64] `json:"lowerAltitude,omitzero"`
	// Highest altitude of the track route altitude block above mean sea level in feet.
	UpperAltitude param.Opt[float64] `json:"upperAltitude,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteUpdateParamsAltitudeBlock) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r TrackRouteUpdateParamsAltitudeBlock) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteUpdateParamsAltitudeBlock
	return param.MarshalObject(r, (*shadow)(&r))
}

// Point of contacts for scheduling or modifying the route.
type TrackRouteUpdateParamsPoc struct {
	// Office name for which the contact belongs.
	Office param.Opt[string] `json:"office,omitzero"`
	// Phone number of the contact.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// The name of the contact.
	PocName param.Opt[string] `json:"pocName,omitzero"`
	// Organization name for which the contact belongs.
	PocOrg param.Opt[string] `json:"pocOrg,omitzero"`
	// Sequencing field for point of contact.
	PocSequenceID param.Opt[int64] `json:"pocSequenceId,omitzero"`
	// A code or name that represents the contact's role in association to the track
	// route (ex. Originator, Scheduler, Maintainer, etc.).
	PocTypeName param.Opt[string] `json:"pocTypeName,omitzero"`
	// The rank of contact.
	Rank param.Opt[string] `json:"rank,omitzero"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The username of the contact.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteUpdateParamsPoc) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r TrackRouteUpdateParamsPoc) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteUpdateParamsPoc
	return param.MarshalObject(r, (*shadow)(&r))
}

// Points identified within the route.
type TrackRouteUpdateParamsRoutePoint struct {
	// Specifies an alternate country code if the data provider code is not part of an
	// official NAVAID Country Code standard such as ISO-3166 or FIPS. This field will
	// be set to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// The DoD Standard Country Code designator for the country where the route point
	// resides. This field should be set to "OTHR" if the source value does not match a
	// UDL country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Flag indicating this is a Digital Aeronautical Flight Information File (DAFIF)
	// point.
	DafifPt param.Opt[bool] `json:"dafifPt,omitzero"`
	// The magnetic declination/variation of the route point location from true north,
	// in degrees. Positive values east of true north and negative values west of true
	// north.
	MagDec param.Opt[float64] `json:"magDec,omitzero"`
	// Navigational Aid (NAVAID) identification code.
	Navaid param.Opt[string] `json:"navaid,omitzero"`
	// The length of the course from the Navigational Aid (NAVAID) in nautical miles.
	NavaidLength param.Opt[float64] `json:"navaidLength,omitzero"`
	// The NAVAID type of this route point (ex. VOR, VORTAC, TACAN, etc.).
	NavaidType param.Opt[string] `json:"navaidType,omitzero"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PtLat param.Opt[float64] `json:"ptLat,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	PtLon param.Opt[float64] `json:"ptLon,omitzero"`
	// Sequencing field for the track route. This is the identifier representing the
	// sequence of waypoints associated to the track route.
	PtSequenceID param.Opt[int64] `json:"ptSequenceId,omitzero"`
	// Code representation of the point within the track route (ex. EP, EX, CP, IP,
	// etc.).
	PtTypeCode param.Opt[string] `json:"ptTypeCode,omitzero"`
	// The name that represents the point within the track route (ex. ENTRY POINT, EXIT
	// POINT, CONTROL POINT, INITIAL POINT, etc.).
	PtTypeName param.Opt[string] `json:"ptTypeName,omitzero"`
	// Name of a waypoint which identifies the location of the point.
	WaypointName param.Opt[string] `json:"waypointName,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteUpdateParamsRoutePoint) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r TrackRouteUpdateParamsRoutePoint) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteUpdateParamsRoutePoint
	return param.MarshalObject(r, (*shadow)(&r))
}

type TrackRouteListParams struct {
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	LastUpdateDate time.Time        `query:"lastUpdateDate,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TrackRouteListParams]'s query parameters as `url.Values`.
func (r TrackRouteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TrackRouteCountParams struct {
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	LastUpdateDate time.Time        `query:"lastUpdateDate,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TrackRouteCountParams]'s query parameters as `url.Values`.
func (r TrackRouteCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TrackRouteNewBulkParams struct {
	Body []TrackRouteNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r TrackRouteNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// A track route is a prescribed route for performing training events or operations
// such as air refueling.
//
// The properties ClassificationMarking, DataMode, LastUpdateDate, Source, Type are
// required.
type TrackRouteNewBulkParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision.
	LastUpdateDate time.Time `json:"lastUpdateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The track route type represented by this record (ex. AIR REFUELING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The APN radar code sent and received by the aircraft for identification.
	ApnSetting param.Opt[string] `json:"apnSetting,omitzero"`
	// The APX radar code sent and received by the aircraft for identification.
	ApxBeaconCode param.Opt[string] `json:"apxBeaconCode,omitzero"`
	// Air Refueling Track Control Center message.
	ArtccMessage param.Opt[string] `json:"artccMessage,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The name of the creating organization of the track route.
	CreatingOrg param.Opt[string] `json:"creatingOrg,omitzero"`
	// The principal compass direction (cardinal or ordinal) of the track route.
	Direction param.Opt[string] `json:"direction,omitzero"`
	// The date which the DAFIF track was last updated/validated in ISO 8601 UTC format
	// with millisecond precision.
	EffectiveDate param.Opt[time.Time] `json:"effectiveDate,omitzero" format:"date-time"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Used to show last time the track route was added to an itinerary in ISO 8601 UTC
	// format with millisecond precision.
	LastUsedDate param.Opt[time.Time] `json:"lastUsedDate,omitzero" format:"date-time"`
	// Track location ID.
	LocationTrackID param.Opt[string] `json:"locationTrackId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The primary UHF radio frequency used for the track route in megahertz.
	PriFreq param.Opt[float64] `json:"priFreq,omitzero"`
	// The receiver tanker channel identifer for air refueling tracks.
	ReceiverTankerChCode param.Opt[string] `json:"receiverTankerCHCode,omitzero"`
	// Region code indicating where the track resides as determined by the data source.
	RegionCode param.Opt[string] `json:"regionCode,omitzero"`
	// Region where the track resides.
	RegionName param.Opt[string] `json:"regionName,omitzero"`
	// Date the track needs to be reviewed for accuracy or deletion in ISO 8601 UTC
	// format with millisecond precision.
	ReviewDate param.Opt[time.Time] `json:"reviewDate,omitzero" format:"date-time"`
	// Point of contact for the air refueling track route scheduler.
	SchedulerOrgName param.Opt[string] `json:"schedulerOrgName,omitzero"`
	// The unit responsible for scheduling the track route.
	SchedulerOrgUnit param.Opt[string] `json:"schedulerOrgUnit,omitzero"`
	// The secondary UHF radio frequency used for the track route in megahertz.
	SecFreq param.Opt[float64] `json:"secFreq,omitzero"`
	// Abbreviated name of the track.
	ShortName param.Opt[string] `json:"shortName,omitzero"`
	// Standard Indicator Code of the air refueling track.
	Sic param.Opt[string] `json:"sic,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Identifier of the track.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Name of the track.
	TrackName param.Opt[string] `json:"trackName,omitzero"`
	// Type of process used by AMC to schedule an air refueling event. Possible values
	// are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched Theater
	// Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S (Soft Air
	// Refueling), T (Matched Theater Operation Short Notice (Theater Assets)), V
	// (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short Notice (AMC
	// Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z (Other Air
	// Refueling).
	TypeCode param.Opt[string] `json:"typeCode,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Minimum and maximum altitude bounds for the track.
	AltitudeBlocks []TrackRouteNewBulkParamsBodyAltitudeBlock `json:"altitudeBlocks,omitzero"`
	// Point of contacts for scheduling or modifying the route.
	Poc []TrackRouteNewBulkParamsBodyPoc `json:"poc,omitzero"`
	// Points identified within the route.
	RoutePoints []TrackRouteNewBulkParamsBodyRoutePoint `json:"routePoints,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteNewBulkParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r TrackRouteNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[TrackRouteNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Minimum and maximum altitude bounds for the track.
type TrackRouteNewBulkParamsBodyAltitudeBlock struct {
	// Sequencing field for the altitude block.
	AltitudeSequenceID param.Opt[string] `json:"altitudeSequenceId,omitzero"`
	// Lowest altitude of the track route altitude block above mean sea level in feet.
	LowerAltitude param.Opt[float64] `json:"lowerAltitude,omitzero"`
	// Highest altitude of the track route altitude block above mean sea level in feet.
	UpperAltitude param.Opt[float64] `json:"upperAltitude,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteNewBulkParamsBodyAltitudeBlock) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r TrackRouteNewBulkParamsBodyAltitudeBlock) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteNewBulkParamsBodyAltitudeBlock
	return param.MarshalObject(r, (*shadow)(&r))
}

// Point of contacts for scheduling or modifying the route.
type TrackRouteNewBulkParamsBodyPoc struct {
	// Office name for which the contact belongs.
	Office param.Opt[string] `json:"office,omitzero"`
	// Phone number of the contact.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// The name of the contact.
	PocName param.Opt[string] `json:"pocName,omitzero"`
	// Organization name for which the contact belongs.
	PocOrg param.Opt[string] `json:"pocOrg,omitzero"`
	// Sequencing field for point of contact.
	PocSequenceID param.Opt[int64] `json:"pocSequenceId,omitzero"`
	// A code or name that represents the contact's role in association to the track
	// route (ex. Originator, Scheduler, Maintainer, etc.).
	PocTypeName param.Opt[string] `json:"pocTypeName,omitzero"`
	// The rank of contact.
	Rank param.Opt[string] `json:"rank,omitzero"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The username of the contact.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteNewBulkParamsBodyPoc) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r TrackRouteNewBulkParamsBodyPoc) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteNewBulkParamsBodyPoc
	return param.MarshalObject(r, (*shadow)(&r))
}

// Points identified within the route.
type TrackRouteNewBulkParamsBodyRoutePoint struct {
	// Specifies an alternate country code if the data provider code is not part of an
	// official NAVAID Country Code standard such as ISO-3166 or FIPS. This field will
	// be set to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// The DoD Standard Country Code designator for the country where the route point
	// resides. This field should be set to "OTHR" if the source value does not match a
	// UDL country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Flag indicating this is a Digital Aeronautical Flight Information File (DAFIF)
	// point.
	DafifPt param.Opt[bool] `json:"dafifPt,omitzero"`
	// The magnetic declination/variation of the route point location from true north,
	// in degrees. Positive values east of true north and negative values west of true
	// north.
	MagDec param.Opt[float64] `json:"magDec,omitzero"`
	// Navigational Aid (NAVAID) identification code.
	Navaid param.Opt[string] `json:"navaid,omitzero"`
	// The length of the course from the Navigational Aid (NAVAID) in nautical miles.
	NavaidLength param.Opt[float64] `json:"navaidLength,omitzero"`
	// The NAVAID type of this route point (ex. VOR, VORTAC, TACAN, etc.).
	NavaidType param.Opt[string] `json:"navaidType,omitzero"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PtLat param.Opt[float64] `json:"ptLat,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	PtLon param.Opt[float64] `json:"ptLon,omitzero"`
	// Sequencing field for the track route. This is the identifier representing the
	// sequence of waypoints associated to the track route.
	PtSequenceID param.Opt[int64] `json:"ptSequenceId,omitzero"`
	// Code representation of the point within the track route (ex. EP, EX, CP, IP,
	// etc.).
	PtTypeCode param.Opt[string] `json:"ptTypeCode,omitzero"`
	// The name that represents the point within the track route (ex. ENTRY POINT, EXIT
	// POINT, CONTROL POINT, INITIAL POINT, etc.).
	PtTypeName param.Opt[string] `json:"ptTypeName,omitzero"`
	// Name of a waypoint which identifies the location of the point.
	WaypointName param.Opt[string] `json:"waypointName,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteNewBulkParamsBodyRoutePoint) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r TrackRouteNewBulkParamsBodyRoutePoint) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteNewBulkParamsBodyRoutePoint
	return param.MarshalObject(r, (*shadow)(&r))
}

type TrackRouteGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TrackRouteGetParams]'s query parameters as `url.Values`.
func (r TrackRouteGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TrackRouteTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	LastUpdateDate time.Time        `query:"lastUpdateDate,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TrackRouteTupleParams]'s query parameters as `url.Values`.
func (r TrackRouteTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TrackRouteUnvalidatedPublishParams struct {
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
	DataMode TrackRouteUnvalidatedPublishParamsDataMode `json:"dataMode,omitzero,required"`
	// The last updated date of the track route in ISO 8601 UTC format with millisecond
	// precision.
	LastUpdateDate time.Time `json:"lastUpdateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The track route type represented by this record (ex. AIR REFUELING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The APN radar code sent and received by the aircraft for identification.
	ApnSetting param.Opt[string] `json:"apnSetting,omitzero"`
	// The APX radar code sent and received by the aircraft for identification.
	ApxBeaconCode param.Opt[string] `json:"apxBeaconCode,omitzero"`
	// Air Refueling Track Control Center message.
	ArtccMessage param.Opt[string] `json:"artccMessage,omitzero"`
	// The name of the creating organization of the track route.
	CreatingOrg param.Opt[string] `json:"creatingOrg,omitzero"`
	// The principal compass direction (cardinal or ordinal) of the track route.
	Direction param.Opt[string] `json:"direction,omitzero"`
	// The date which the DAFIF track was last updated/validated in ISO 8601 UTC format
	// with millisecond precision.
	EffectiveDate param.Opt[time.Time] `json:"effectiveDate,omitzero" format:"date-time"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Used to show last time the track route was added to an itinerary in ISO 8601 UTC
	// format with millisecond precision.
	LastUsedDate param.Opt[time.Time] `json:"lastUsedDate,omitzero" format:"date-time"`
	// Track location ID.
	LocationTrackID param.Opt[string] `json:"locationTrackId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The primary UHF radio frequency used for the track route in megahertz.
	PriFreq param.Opt[float64] `json:"priFreq,omitzero"`
	// The receiver tanker channel identifer for air refueling tracks.
	ReceiverTankerChCode param.Opt[string] `json:"receiverTankerCHCode,omitzero"`
	// Region code indicating where the track resides as determined by the data source.
	RegionCode param.Opt[string] `json:"regionCode,omitzero"`
	// Region where the track resides.
	RegionName param.Opt[string] `json:"regionName,omitzero"`
	// Date the track needs to be reviewed for accuracy or deletion in ISO 8601 UTC
	// format with millisecond precision.
	ReviewDate param.Opt[time.Time] `json:"reviewDate,omitzero" format:"date-time"`
	// Point of contact for the air refueling track route scheduler.
	SchedulerOrgName param.Opt[string] `json:"schedulerOrgName,omitzero"`
	// The unit responsible for scheduling the track route.
	SchedulerOrgUnit param.Opt[string] `json:"schedulerOrgUnit,omitzero"`
	// The secondary UHF radio frequency used for the track route in megahertz.
	SecFreq param.Opt[float64] `json:"secFreq,omitzero"`
	// Abbreviated name of the track.
	ShortName param.Opt[string] `json:"shortName,omitzero"`
	// Standard Indicator Code of the air refueling track.
	Sic param.Opt[string] `json:"sic,omitzero"`
	// Identifier of the track.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Name of the track.
	TrackName param.Opt[string] `json:"trackName,omitzero"`
	// Type of process used by AMC to schedule an air refueling event. Possible values
	// are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched Theater
	// Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S (Soft Air
	// Refueling), T (Matched Theater Operation Short Notice (Theater Assets)), V
	// (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short Notice (AMC
	// Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z (Other Air
	// Refueling).
	TypeCode param.Opt[string] `json:"typeCode,omitzero"`
	// Minimum and maximum altitude bounds for the track.
	AltitudeBlocks []TrackRouteUnvalidatedPublishParamsAltitudeBlock `json:"altitudeBlocks,omitzero"`
	// Point of contacts for scheduling or modifying the route.
	Poc []TrackRouteUnvalidatedPublishParamsPoc `json:"poc,omitzero"`
	// Points identified within the route.
	RoutePoints []TrackRouteUnvalidatedPublishParamsRoutePoint `json:"routePoints,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r TrackRouteUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteUnvalidatedPublishParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type TrackRouteUnvalidatedPublishParamsDataMode string

const (
	TrackRouteUnvalidatedPublishParamsDataModeReal      TrackRouteUnvalidatedPublishParamsDataMode = "REAL"
	TrackRouteUnvalidatedPublishParamsDataModeTest      TrackRouteUnvalidatedPublishParamsDataMode = "TEST"
	TrackRouteUnvalidatedPublishParamsDataModeSimulated TrackRouteUnvalidatedPublishParamsDataMode = "SIMULATED"
	TrackRouteUnvalidatedPublishParamsDataModeExercise  TrackRouteUnvalidatedPublishParamsDataMode = "EXERCISE"
)

// Minimum and maximum altitude bounds for the track.
type TrackRouteUnvalidatedPublishParamsAltitudeBlock struct {
	// Sequencing field for the altitude block.
	AltitudeSequenceID param.Opt[string] `json:"altitudeSequenceId,omitzero"`
	// Lowest altitude of the track route altitude block above mean sea level in feet.
	LowerAltitude param.Opt[float64] `json:"lowerAltitude,omitzero"`
	// Highest altitude of the track route altitude block above mean sea level in feet.
	UpperAltitude param.Opt[float64] `json:"upperAltitude,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteUnvalidatedPublishParamsAltitudeBlock) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r TrackRouteUnvalidatedPublishParamsAltitudeBlock) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteUnvalidatedPublishParamsAltitudeBlock
	return param.MarshalObject(r, (*shadow)(&r))
}

// Point of contacts for scheduling or modifying the route.
type TrackRouteUnvalidatedPublishParamsPoc struct {
	// Office name for which the contact belongs.
	Office param.Opt[string] `json:"office,omitzero"`
	// Phone number of the contact.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// The name of the contact.
	PocName param.Opt[string] `json:"pocName,omitzero"`
	// Organization name for which the contact belongs.
	PocOrg param.Opt[string] `json:"pocOrg,omitzero"`
	// Sequencing field for point of contact.
	PocSequenceID param.Opt[int64] `json:"pocSequenceId,omitzero"`
	// A code or name that represents the contact's role in association to the track
	// route (ex. Originator, Scheduler, Maintainer, etc.).
	PocTypeName param.Opt[string] `json:"pocTypeName,omitzero"`
	// The rank of contact.
	Rank param.Opt[string] `json:"rank,omitzero"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The username of the contact.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteUnvalidatedPublishParamsPoc) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r TrackRouteUnvalidatedPublishParamsPoc) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteUnvalidatedPublishParamsPoc
	return param.MarshalObject(r, (*shadow)(&r))
}

// Points identified within the route.
type TrackRouteUnvalidatedPublishParamsRoutePoint struct {
	// Specifies an alternate country code if the data provider code is not part of an
	// official NAVAID Country Code standard such as ISO-3166 or FIPS. This field will
	// be set to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// The DoD Standard Country Code designator for the country where the route point
	// resides. This field should be set to "OTHR" if the source value does not match a
	// UDL country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Flag indicating this is a Digital Aeronautical Flight Information File (DAFIF)
	// point.
	DafifPt param.Opt[bool] `json:"dafifPt,omitzero"`
	// The magnetic declination/variation of the route point location from true north,
	// in degrees. Positive values east of true north and negative values west of true
	// north.
	MagDec param.Opt[float64] `json:"magDec,omitzero"`
	// Navigational Aid (NAVAID) identification code.
	Navaid param.Opt[string] `json:"navaid,omitzero"`
	// The length of the course from the Navigational Aid (NAVAID) in nautical miles.
	NavaidLength param.Opt[float64] `json:"navaidLength,omitzero"`
	// The NAVAID type of this route point (ex. VOR, VORTAC, TACAN, etc.).
	NavaidType param.Opt[string] `json:"navaidType,omitzero"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PtLat param.Opt[float64] `json:"ptLat,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	PtLon param.Opt[float64] `json:"ptLon,omitzero"`
	// Sequencing field for the track route. This is the identifier representing the
	// sequence of waypoints associated to the track route.
	PtSequenceID param.Opt[int64] `json:"ptSequenceId,omitzero"`
	// Code representation of the point within the track route (ex. EP, EX, CP, IP,
	// etc.).
	PtTypeCode param.Opt[string] `json:"ptTypeCode,omitzero"`
	// The name that represents the point within the track route (ex. ENTRY POINT, EXIT
	// POINT, CONTROL POINT, INITIAL POINT, etc.).
	PtTypeName param.Opt[string] `json:"ptTypeName,omitzero"`
	// Name of a waypoint which identifies the location of the point.
	WaypointName param.Opt[string] `json:"waypointName,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TrackRouteUnvalidatedPublishParamsRoutePoint) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r TrackRouteUnvalidatedPublishParamsRoutePoint) MarshalJSON() (data []byte, err error) {
	type shadow TrackRouteUnvalidatedPublishParamsRoutePoint
	return param.MarshalObject(r, (*shadow)(&r))
}
