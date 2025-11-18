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

// LaserdeconflictrequestHistoryService contains methods and other services that
// help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaserdeconflictrequestHistoryService] method instead.
type LaserdeconflictrequestHistoryService struct {
	Options []option.RequestOption
}

// NewLaserdeconflictrequestHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLaserdeconflictrequestHistoryService(opts ...option.RequestOption) (r LaserdeconflictrequestHistoryService) {
	r = LaserdeconflictrequestHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaserdeconflictrequestHistoryService) List(ctx context.Context, query LaserdeconflictrequestHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaserdeconflictrequestHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/laserdeconflictrequest/history"
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
func (r *LaserdeconflictrequestHistoryService) ListAutoPaging(ctx context.Context, query LaserdeconflictrequestHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaserdeconflictrequestHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaserdeconflictrequestHistoryService) Aodr(ctx context.Context, query LaserdeconflictrequestHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/laserdeconflictrequest/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LaserdeconflictrequestHistoryService) Count(ctx context.Context, query LaserdeconflictrequestHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/laserdeconflictrequest/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The LaserDeconflictionRequest service is designed to process and manage requests
// related to the safe operation of high-powered lasers, ensuring that laser
// systems do not interfere with other critical operations, such as satellite based
// activities. The service facilitates real-time coordination between different
// entities to prevent accidental exposure to laser hazards, ensuring compliance
// with safety protocols and regulations.
type LaserdeconflictrequestHistoryListResponse struct {
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
	DataMode LaserdeconflictrequestHistoryListResponseDataMode `json:"dataMode,required"`
	// End date of the time windows associated with this LaserDeconflictRequest, in ISO
	// 8601 UTC format with millisecond precision.
	EndDate time.Time `json:"endDate,required" format:"date-time"`
	// A list containing the id strings of the LaserEmitter records in UDL detailing
	// the physical parameters of each laser/emitter operationally involved with this
	// request. All laser/emitter components must be accurately described using the
	// LaserEmitter schema and ingested into the UDL LaserEmitter service before
	// creating a LaserDeconflictRequest. Users should create new LaserEmitter records
	// for non-existent emitters and update existing records for any modifications.
	IDLaserEmitters []string `json:"idLaserEmitters,required"`
	// The number of targets included in this request.
	NumTargets int64 `json:"numTargets,required"`
	// External identifier for this LaserDeconflictRequest record.
	RequestID string `json:"requestId,required"`
	// The datetime that this LaserDeconflictRequest record was created, in ISO 8601
	// UTC format with millisecond precision.
	RequestTs time.Time `json:"requestTs,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Start date of the time windows associated with this LaserDeconflictRequest, in
	// ISO 8601 UTC format with millisecond precision.
	StartDate time.Time `json:"startDate,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The azimuth angle of the centerline of the geospatial box that confines
	// positions of the laser platform, in degrees.
	CenterlineAzimuth float64 `json:"centerlineAzimuth"`
	// The elevation angle of the centerline of the geospatial box that confines the
	// positions of the laser platform, in degrees.
	CenterlineElevation float64 `json:"centerlineElevation"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The half-angle of the safety cone of the laser beam, in degrees.
	DefaultCha float64 `json:"defaultCHA"`
	// Boolean enabling Dynamic Satellite Susceptibility (DSS) algorithms.
	EnableDss bool `json:"enableDSS"`
	// A collection of latitude, longitude, and altitude fields which can be used to
	// specify the geometry of the coordinate space in which the laser platform(s) will
	// be operational for this request. For example, a BOX_2_WAYPOINTS would include
	// two data points, while a BOX_4_SURFACE_POINTS would include four data points.
	FixedPoints []FixedPointFull `json:"fixedPoints"`
	// Indicates the geopotential model used in the propagation calculation for this
	// request (e.g. EGM-96, WGS-84, WGS-72, WGS66, WGS60, JGM-2, or GEM-T3).
	GeopotentialModel string `json:"geopotentialModel"`
	// Unique identifier of the on-orbit laser platform.
	IDOnOrbit string `json:"idOnOrbit"`
	// A list containing all laser illumination target object specifications for which
	// deconflictions must be calculated, as planned for this request.
	LaserDeconflictTargets []LaserdeconflictrequestHistoryListResponseLaserDeconflictTarget `json:"laserDeconflictTargets"`
	// The name of the laser/beam director system. The Laser Clearinghouse will append
	// identifiers to the name using standard conventions.
	LaserSystemName string `json:"laserSystemName"`
	// The length of the centerline that passes through the center point of the
	// geospatial box that confines the positions of the laser platform, in kilometers.
	LengthCenterline float64 `json:"lengthCenterline"`
	// Specifies the length of the horizontal dimension of the geospatial box that
	// confines the positions of the laser platform, in kilometers.
	LengthLeftRight float64 `json:"lengthLeftRight"`
	// Specifies the length of the vertical dimension of the geospatial box that
	// confines the positions of the laser platform, in kilometers.
	LengthUpDown float64 `json:"lengthUpDown"`
	// The maximum laser operating altitude specified as the height above/below the
	// WGS84 ellipsoid where the laser is omitted from, in kilometers.
	MaximumHeight float64 `json:"maximumHeight"`
	// The minimum laser operating altitude specified as the height above/below the
	// WGS84 ellipsoid where the laser is omitted from, in kilometers.
	MinimumHeight float64 `json:"minimumHeight"`
	// The name of the mission with which this LaserDeconflictRequest is associated.
	MissionName string `json:"missionName"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the source provider to indicate the on-orbit
	// laser platform. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// The name of the laser platform.
	PlatformLocationName string `json:"platformLocationName"`
	// Indicates the type of location(s) the laser platform will be operational for
	// this request (BOX_2_WAYPOINTS, BOX_4_SURFACE_POINTS, BOX_CENTER_POINT_LINE,
	// EXTERNAL_EPHEMERIS, FIXED_POINT, SATELLITE).
	PlatformLocationType string `json:"platformLocationType"`
	// External identifier for the program that is responsible for this
	// LaserDeconflictRequest.
	ProgramID string `json:"programId"`
	// The type of propagator utilized in the deconfliction/predictive avoidance
	// calculation.
	Propagator string `json:"propagator"`
	// A list of satellite/catalog numbers that should be protected from any and all
	// incidence of laser illumination for the duration of this request.
	ProtectList []int64 `json:"protectList"`
	// The satellite/catalog number of the on-orbit laser platform.
	SatNo int64 `json:"satNo"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Boolean indicating whether error growth of the laser beam is enabled for this
	// request.
	SourceEnabled bool `json:"sourceEnabled"`
	// Status of this request (APPROVED, COMPLETE_WITH_ERRORS, COMPLETE_WITH_WARNINGS,
	// FAILURE, IN_PROGRESS, REQUESTED, SUCCESS).
	Status string `json:"status"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Boolean indicating whether target error growth is enabled for this request.
	TargetEnabled bool `json:"targetEnabled"`
	// The target type that concerns this request (BOX_2_WAYPOINTS,
	// BOX_4_SURFACE_POINTS, BOX_CENTER_POINT_LINE, EXTERNAL_EPHEMERIS, FIXED_POINT,
	// SATELLITE).
	TargetType string `json:"targetType"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Indicates the treatment of earth (INVISIBLE, VICTIM, SHIELD) for this
	// LaserDeconflictRequest record.
	TreatEarthAs string `json:"treatEarthAs"`
	// Boolean indicating that, for deconfliction events in which the potential target
	// is an optical imaging satellite, line of sight computation between target and
	// source is ensured when the source emitter is contained within the field of
	// regard (field of view) of the satellite's optical telescope.
	UseFieldOfRegard bool `json:"useFieldOfRegard"`
	// Boolean indicating whether victim error growth is enabled as input to the
	// deconfliction calculations for this request.
	VictimEnabled bool `json:"victimEnabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		EndDate                respjson.Field
		IDLaserEmitters        respjson.Field
		NumTargets             respjson.Field
		RequestID              respjson.Field
		RequestTs              respjson.Field
		Source                 respjson.Field
		StartDate              respjson.Field
		ID                     respjson.Field
		CenterlineAzimuth      respjson.Field
		CenterlineElevation    respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DefaultCha             respjson.Field
		EnableDss              respjson.Field
		FixedPoints            respjson.Field
		GeopotentialModel      respjson.Field
		IDOnOrbit              respjson.Field
		LaserDeconflictTargets respjson.Field
		LaserSystemName        respjson.Field
		LengthCenterline       respjson.Field
		LengthLeftRight        respjson.Field
		LengthUpDown           respjson.Field
		MaximumHeight          respjson.Field
		MinimumHeight          respjson.Field
		MissionName            respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		OrigObjectID           respjson.Field
		PlatformLocationName   respjson.Field
		PlatformLocationType   respjson.Field
		ProgramID              respjson.Field
		Propagator             respjson.Field
		ProtectList            respjson.Field
		SatNo                  respjson.Field
		SourceDl               respjson.Field
		SourceEnabled          respjson.Field
		Status                 respjson.Field
		Tags                   respjson.Field
		TargetEnabled          respjson.Field
		TargetType             respjson.Field
		TransactionID          respjson.Field
		TreatEarthAs           respjson.Field
		UseFieldOfRegard       respjson.Field
		VictimEnabled          respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaserdeconflictrequestHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaserdeconflictrequestHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type LaserdeconflictrequestHistoryListResponseDataMode string

const (
	LaserdeconflictrequestHistoryListResponseDataModeReal      LaserdeconflictrequestHistoryListResponseDataMode = "REAL"
	LaserdeconflictrequestHistoryListResponseDataModeTest      LaserdeconflictrequestHistoryListResponseDataMode = "TEST"
	LaserdeconflictrequestHistoryListResponseDataModeSimulated LaserdeconflictrequestHistoryListResponseDataMode = "SIMULATED"
	LaserdeconflictrequestHistoryListResponseDataModeExercise  LaserdeconflictrequestHistoryListResponseDataMode = "EXERCISE"
)

// Model representation which can be used for the specification of target object
// for a laser operation. The specification of various target types common across
// space-to-space, space-to-ground, and ground-to-space laser operations are
// supported by this model.
type LaserdeconflictrequestHistoryListResponseLaserDeconflictTarget struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The expected or directed azimuth angle of this target, in degrees.
	Azimuth float64 `json:"azimuth"`
	// The secondary azimuth angle specifying the end of the azimuthal range that
	// defines this target area, in degrees.
	AzimuthEnd float64 `json:"azimuthEnd"`
	// The incremental change in angle over the given azimuthal range at which the
	// target area will be engaged, in degrees.
	AzimuthIncrement float64 `json:"azimuthIncrement"`
	// The initial azimuth angle specifying the start of the azimuthal range that
	// defines this target area, in degrees.
	AzimuthStart float64 `json:"azimuthStart"`
	// The azimuth angle of the centerline of the geospatial box that confines
	// positions of the target, in degrees.
	CenterlineAzimuth float64 `json:"centerlineAzimuth"`
	// The elevation angle of the centerline of the geospatial box that confines the
	// positions of the target, in degrees.
	CenterlineElevation float64 `json:"centerlineElevation"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The expected or directed declination angle of this target, in degrees.
	Declination float64 `json:"declination"`
	// The expected or directed elevation angle of this target, in degrees.
	Elevation float64 `json:"elevation"`
	// The secondary elevation angle specifying the end of the elevation range that
	// defines this target area, in degrees.
	ElevationEnd float64 `json:"elevationEnd"`
	// The incremental change in angle across the given elevation range at which the
	// target area will be engaged, in degrees.
	ElevationIncrement float64 `json:"elevationIncrement"`
	// The initial elevation angle specifying the start of the elevation range that
	// defines this target area, in degrees.
	ElevationStart float64 `json:"elevationStart"`
	// A collection of latitude, longitude, and altitude fields which can be used to
	// specify the geometry of the coordinate space of this target. For example, a
	// BOX_2_WAYPOINTS would include two data points, while a BOX_4_SURFACE_POINTS
	// would include four data points.
	FixedPoints []FixedPointFull `json:"fixedPoints"`
	// Unique identifier of the parent LaserDeconflictRequest record containing this
	// target.
	IDLaserDeconflictRequest string `json:"idLaserDeconflictRequest"`
	// The length of the centerline that passes through the center point of the
	// geospatial box that confines the positions of the target, in kilometers.
	LengthCenterline float64 `json:"lengthCenterline"`
	// Specifies the length of the horizontal dimension of the geospatial box that
	// confines the positions of the target, in kilometers.
	LengthLeftRight float64 `json:"lengthLeftRight"`
	// Specifies the length of the vertical dimension of the geospatial box that
	// confines the positions of the target, in kilometers.
	LengthUpDown float64 `json:"lengthUpDown"`
	// The maximum target altitude specified as the height above/below the WGS84
	// ellipsoid, in kilometers.
	MaximumHeight float64 `json:"maximumHeight"`
	// The minimum target altitude specified as the height above/below the WGS84
	// ellipsoid, in kilometers.
	MinimumHeight float64 `json:"minimumHeight"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The expected or directed right ascension angle of this target, in degrees.
	Ra float64 `json:"ra"`
	// The name of the target celestial body in Earth's solar system (JUPITER, MARS,
	// MERCURY, MOON, NEPTUNE, PLUTO, SATURN, SUN, URANUS, VENUS).
	SolarSystemBody string `json:"solarSystemBody"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The reference number of the target star.
	StarNumber int64 `json:"starNumber"`
	// Start date of the time windows associated with this LaserDeconflictRequest, in
	// ISO 8601 UTC format with millisecond precision.
	StartDate time.Time `json:"startDate" format:"date-time"`
	// The number assigned to this target instance for a request.
	TargetNumber int64 `json:"targetNumber"`
	// Optional target identifier provided by the source provider. This may be an
	// internal identifier and not necessarily map to UDL entities.
	TargetObjectID string `json:"targetObjectId"`
	// If this is an on-orbit target, this is the satellite/catalog number.
	TargetObjectNo int64 `json:"targetObjectNo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		Azimuth                  respjson.Field
		AzimuthEnd               respjson.Field
		AzimuthIncrement         respjson.Field
		AzimuthStart             respjson.Field
		CenterlineAzimuth        respjson.Field
		CenterlineElevation      respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		Declination              respjson.Field
		Elevation                respjson.Field
		ElevationEnd             respjson.Field
		ElevationIncrement       respjson.Field
		ElevationStart           respjson.Field
		FixedPoints              respjson.Field
		IDLaserDeconflictRequest respjson.Field
		LengthCenterline         respjson.Field
		LengthLeftRight          respjson.Field
		LengthUpDown             respjson.Field
		MaximumHeight            respjson.Field
		MinimumHeight            respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		Ra                       respjson.Field
		SolarSystemBody          respjson.Field
		SourceDl                 respjson.Field
		StarNumber               respjson.Field
		StartDate                respjson.Field
		TargetNumber             respjson.Field
		TargetObjectID           respjson.Field
		TargetObjectNo           respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaserdeconflictrequestHistoryListResponseLaserDeconflictTarget) RawJSON() string {
	return r.JSON.raw
}
func (r *LaserdeconflictrequestHistoryListResponseLaserDeconflictTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaserdeconflictrequestHistoryListParams struct {
	// Start date of the time windows associated with this LaserDeconflictRequest, in
	// ISO 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartDate time.Time `query:"startDate,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaserdeconflictrequestHistoryListParams]'s query parameters
// as `url.Values`.
func (r LaserdeconflictrequestHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaserdeconflictrequestHistoryAodrParams struct {
	// Start date of the time windows associated with this LaserDeconflictRequest, in
	// ISO 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartDate time.Time `query:"startDate,required" format:"date-time" json:"-"`
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

// URLQuery serializes [LaserdeconflictrequestHistoryAodrParams]'s query parameters
// as `url.Values`.
func (r LaserdeconflictrequestHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaserdeconflictrequestHistoryCountParams struct {
	// Start date of the time windows associated with this LaserDeconflictRequest, in
	// ISO 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartDate   time.Time        `query:"startDate,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaserdeconflictrequestHistoryCountParams]'s query
// parameters as `url.Values`.
func (r LaserdeconflictrequestHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
