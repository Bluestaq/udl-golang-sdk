// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// ObservationPassiveRadarObservationHistoryService contains methods and other
// services that help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationPassiveRadarObservationHistoryService] method instead.
type ObservationPassiveRadarObservationHistoryService struct {
	Options []option.RequestOption
}

// NewObservationPassiveRadarObservationHistoryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewObservationPassiveRadarObservationHistoryService(opts ...option.RequestOption) (r ObservationPassiveRadarObservationHistoryService) {
	r = ObservationPassiveRadarObservationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationPassiveRadarObservationHistoryService) List(ctx context.Context, query ObservationPassiveRadarObservationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ObservationPassiveRadarObservationHistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/passiveradarobservation/history"
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
func (r *ObservationPassiveRadarObservationHistoryService) ListAutoPaging(ctx context.Context, query ObservationPassiveRadarObservationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ObservationPassiveRadarObservationHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationPassiveRadarObservationHistoryService) Aodr(ctx context.Context, query ObservationPassiveRadarObservationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/passiveradarobservation/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationPassiveRadarObservationHistoryService) Count(ctx context.Context, query ObservationPassiveRadarObservationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/passiveradarobservation/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of observation data for passive radar based sensor
// phenomenologies. Passive radar is a form of radar that instead of using a
// dedicated transmitter, as is the case for traditional radar, utilizes radio
// frequency (RF) energy already in the environment. With an abundance of existing
// energy available from geographically diverse sources, passive radar offers wide
// field of view coverage and long observation times. A passive radar system is
// comprised of separately located transmitter (for example, FM radio stations), a
// reference receiver, and a surveillance sensor. The transmitted signal
// illuminates multiple targets over a broad angular extent as well as providing a
// signal at the reference site. The reflected energy is received at the
// surveillance site for processing with the reference signal. The long observation
// durations that are possible with this technology enables an accurate orbit to be
// determined within a single pass.
type ObservationPassiveRadarObservationHistoryListResponse struct {
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
	DataMode ObservationPassiveRadarObservationHistoryListResponseDataMode `json:"dataMode,required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The target Acceleration measurement in kilometers/sec^2 for this observation.
	Accel float64 `json:"accel"`
	// The target Acceleration uncertainty measurement in kilometers/sec^2 for this
	// observation.
	AccelUnc float64 `json:"accelUnc"`
	// The target altitude relative to WGS-84 ellipsoid, in kilometers for this
	// observation.
	Alt float64 `json:"alt"`
	// Line of sight azimuth angle in degrees and topocentric frame.
	Azimuth float64 `json:"azimuth"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias float64 `json:"azimuthBias"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// Target bistatic path distance in kilometers. This is the
	// transmitter-to-target-to-surveillance site distance.
	BistaticRange float64 `json:"bistaticRange"`
	// Bistatic range acceleration in kilometers/sec^2.
	BistaticRangeAccel float64 `json:"bistaticRangeAccel"`
	// One sigma uncertainty in the bistatic range acceleration measurement, in
	// kilometers/sec^2.
	BistaticRangeAccelUnc float64 `json:"bistaticRangeAccelUnc"`
	// Sensor bistatic range bias in kilometers.
	BistaticRangeBias float64 `json:"bistaticRangeBias"`
	// Rate of change of the bistatic path in kilometers/sec.
	BistaticRangeRate float64 `json:"bistaticRangeRate"`
	// One sigma uncertainty in rate of change of the bistatic path in kilometers/sec.
	BistaticRangeRateUnc float64 `json:"bistaticRangeRateUnc"`
	// One sigma uncertainty in bistatic range in kilometers.
	BistaticRangeUnc float64 `json:"bistaticRangeUnc"`
	// Coning angle in degrees.
	Coning float64 `json:"coning"`
	// One sigma uncertainty in the coning angle measurement, in degrees.
	ConingUnc float64 `json:"coningUnc"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination float64 `json:"declination"`
	// The time difference, in seconds, between the signal collected at the
	// surveillance site (after being reflected from the target) and the reference site
	// (direct path line-of-sight signal).
	Delay float64 `json:"delay"`
	// Delay bias in seconds.
	DelayBias float64 `json:"delayBias"`
	// One sigma uncertainty in the delay measurement, in seconds.
	DelayUnc float64 `json:"delayUnc"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Doppler measurement in hertz.
	Doppler float64 `json:"doppler"`
	// One sigma uncertainty in the Doppler measurement in hertz.
	DopplerUnc float64 `json:"dopplerUnc"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation float64 `json:"elevation"`
	// Sensor elevation bias in degrees.
	ElevationBias float64 `json:"elevationBias"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate float64 `json:"elevationRate"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc float64 `json:"elevationUnc"`
	// Optional external observation identifier provided by the source.
	ExtObservationID string `json:"extObservationId"`
	// Unique identifier of the target satellite on-orbit object. This ID can be used
	// to obtain additional information on an OnOrbit object using the 'get by ID'
	// operation (e.g. /udl/onorbit/{id}). For example, the OnOrbit with idOnOrbit =
	// 25544 would be queried as /udl/onorbit/25544.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the transmitter. This ID can be used to obtain additional
	// information on an RFEmitter using the 'get by ID' operation (e.g.
	// /udl/rfemitter/{id}). For example, the RFEmitter with idRFEmitter = abc would be
	// queried as /udl/rfemitter/abc.
	IDRfEmitter string `json:"idRFEmitter"`
	// Unique identifier of the reporting surveillance sensor. This ID can be used to
	// obtain additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensor string `json:"idSensor"`
	// Unique identifier of the reference receiver sensor. This ID can be used to
	// obtain additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensorRefReceiver string `json:"idSensorRefReceiver"`
	// WGS-84 target latitude sub-point at observation time (obTime), represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat float64 `json:"lat"`
	// WGS-84 target longitude sub-point at observation time (obTime), represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The position of this observation within a track (FENCE, FIRST, IN, LAST,
	// SINGLE). This identifier is optional and, if null, no assumption should be made
	// regarding whether other observations may or may not exist to compose a track.
	ObPosition string `json:"obPosition"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Radar cross section in meters squared for orthogonal polarization.
	OrthogonalRcs float64 `json:"orthogonalRcs"`
	// One sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc float64 `json:"orthogonalRcsUnc"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra float64 `json:"ra"`
	// Radar cross section in meters squared for polarization principal.
	Rcs float64 `json:"rcs"`
	// One sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc float64 `json:"rcsUnc"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
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
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID string `json:"taskId"`
	// Sensor timing bias in seconds.
	TimingBias float64 `json:"timingBias"`
	// Time of flight (TOF) in seconds. This is the calculated propagation time from
	// transmitter-to-target-to-surveillance site.
	Tof float64 `json:"tof"`
	// The Time of Flight (TOF) bias in seconds.
	TofBias float64 `json:"tofBias"`
	// One sigma uncertainty in time of flight in seconds.
	TofUnc float64 `json:"tofUnc"`
	// Unique identifier of a track that represents a tracklet for this observation.
	TrackID string `json:"trackId"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Read only enumeration specifying the type of observation (e.g. OPTICAL, RADAR,
	// RF, etc).
	Type string `json:"type"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct bool `json:"uct"`
	// X velocity of target in kilometers/sec in J2000 coordinate frame.
	Xvel float64 `json:"xvel"`
	// Y velocity of target in kilometers/sec in J2000 coordinate frame.
	Yvel float64 `json:"yvel"`
	// Z velocity of target in kilometers/sec in J2000 coordinate frame.
	Zvel float64 `json:"zvel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Accel                 respjson.Field
		AccelUnc              respjson.Field
		Alt                   respjson.Field
		Azimuth               respjson.Field
		AzimuthBias           respjson.Field
		AzimuthRate           respjson.Field
		AzimuthUnc            respjson.Field
		BistaticRange         respjson.Field
		BistaticRangeAccel    respjson.Field
		BistaticRangeAccelUnc respjson.Field
		BistaticRangeBias     respjson.Field
		BistaticRangeRate     respjson.Field
		BistaticRangeRateUnc  respjson.Field
		BistaticRangeUnc      respjson.Field
		Coning                respjson.Field
		ConingUnc             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Declination           respjson.Field
		Delay                 respjson.Field
		DelayBias             respjson.Field
		DelayUnc              respjson.Field
		Descriptor            respjson.Field
		Doppler               respjson.Field
		DopplerUnc            respjson.Field
		Elevation             respjson.Field
		ElevationBias         respjson.Field
		ElevationRate         respjson.Field
		ElevationUnc          respjson.Field
		ExtObservationID      respjson.Field
		IDOnOrbit             respjson.Field
		IDRfEmitter           respjson.Field
		IDSensor              respjson.Field
		IDSensorRefReceiver   respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		ObPosition            respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		OrthogonalRcs         respjson.Field
		OrthogonalRcsUnc      respjson.Field
		Ra                    respjson.Field
		Rcs                   respjson.Field
		RcsUnc                respjson.Field
		SatNo                 respjson.Field
		Snr                   respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		TaskID                respjson.Field
		TimingBias            respjson.Field
		Tof                   respjson.Field
		TofBias               respjson.Field
		TofUnc                respjson.Field
		TrackID               respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		Uct                   respjson.Field
		Xvel                  respjson.Field
		Yvel                  respjson.Field
		Zvel                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObservationPassiveRadarObservationHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationPassiveRadarObservationHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type ObservationPassiveRadarObservationHistoryListResponseDataMode string

const (
	ObservationPassiveRadarObservationHistoryListResponseDataModeReal      ObservationPassiveRadarObservationHistoryListResponseDataMode = "REAL"
	ObservationPassiveRadarObservationHistoryListResponseDataModeTest      ObservationPassiveRadarObservationHistoryListResponseDataMode = "TEST"
	ObservationPassiveRadarObservationHistoryListResponseDataModeSimulated ObservationPassiveRadarObservationHistoryListResponseDataMode = "SIMULATED"
	ObservationPassiveRadarObservationHistoryListResponseDataModeExercise  ObservationPassiveRadarObservationHistoryListResponseDataMode = "EXERCISE"
)

type ObservationPassiveRadarObservationHistoryListParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime time.Time `query:"obTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationPassiveRadarObservationHistoryListParams]'s
// query parameters as `url.Values`.
func (r ObservationPassiveRadarObservationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationPassiveRadarObservationHistoryAodrParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime time.Time `query:"obTime,required" format:"date-time" json:"-"`
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

// URLQuery serializes [ObservationPassiveRadarObservationHistoryAodrParams]'s
// query parameters as `url.Values`.
func (r ObservationPassiveRadarObservationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationPassiveRadarObservationHistoryCountParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationPassiveRadarObservationHistoryCountParams]'s
// query parameters as `url.Values`.
func (r ObservationPassiveRadarObservationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
