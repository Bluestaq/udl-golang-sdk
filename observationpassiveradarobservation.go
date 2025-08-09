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

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// ObservationPassiveRadarObservationService contains methods and other services
// that help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationPassiveRadarObservationService] method instead.
type ObservationPassiveRadarObservationService struct {
	Options []option.RequestOption
	History ObservationPassiveRadarObservationHistoryService
}

// NewObservationPassiveRadarObservationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewObservationPassiveRadarObservationService(opts ...option.RequestOption) (r ObservationPassiveRadarObservationService) {
	r = ObservationPassiveRadarObservationService{}
	r.Options = opts
	r.History = NewObservationPassiveRadarObservationHistoryService(opts...)
	return
}

// Service operation to take a single PassiveRadarObservation as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *ObservationPassiveRadarObservationService) New(ctx context.Context, body ObservationPassiveRadarObservationNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/passiveradarobservation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationPassiveRadarObservationService) List(ctx context.Context, query ObservationPassiveRadarObservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ObservationPassiveRadarObservationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/passiveradarobservation"
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
func (r *ObservationPassiveRadarObservationService) ListAutoPaging(ctx context.Context, query ObservationPassiveRadarObservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ObservationPassiveRadarObservationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationPassiveRadarObservationService) Count(ctx context.Context, query ObservationPassiveRadarObservationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/passiveradarobservation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// PassiveRadarObservation records as a POST body and ingest into the database.
// This operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *ObservationPassiveRadarObservationService) NewBulk(ctx context.Context, body ObservationPassiveRadarObservationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/passiveradarobservation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to take multiple PassiveRadarObservation records as a POST
// body and ingest into the database. This operation is intended to be used for
// automated feeds into UDL. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *ObservationPassiveRadarObservationService) FileNew(ctx context.Context, body ObservationPassiveRadarObservationFileNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-passiveradar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single PassiveRadarObservation record by its unique
// ID passed as a path parameter.
func (r *ObservationPassiveRadarObservationService) Get(ctx context.Context, id string, query ObservationPassiveRadarObservationGetParams, opts ...option.RequestOption) (res *ObservationPassiveRadarObservationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/passiveradarobservation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ObservationPassiveRadarObservationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *ObservationPassiveRadarObservationQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/passiveradarobservation/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *ObservationPassiveRadarObservationService) Tuple(ctx context.Context, query ObservationPassiveRadarObservationTupleParams, opts ...option.RequestOption) (res *[]ObservationPassiveRadarObservationTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/passiveradarobservation/tuple"
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
type ObservationPassiveRadarObservationListResponse struct {
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
	DataMode ObservationPassiveRadarObservationListResponseDataMode `json:"dataMode,required"`
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
func (r ObservationPassiveRadarObservationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationPassiveRadarObservationListResponse) UnmarshalJSON(data []byte) error {
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
type ObservationPassiveRadarObservationListResponseDataMode string

const (
	ObservationPassiveRadarObservationListResponseDataModeReal      ObservationPassiveRadarObservationListResponseDataMode = "REAL"
	ObservationPassiveRadarObservationListResponseDataModeTest      ObservationPassiveRadarObservationListResponseDataMode = "TEST"
	ObservationPassiveRadarObservationListResponseDataModeSimulated ObservationPassiveRadarObservationListResponseDataMode = "SIMULATED"
	ObservationPassiveRadarObservationListResponseDataModeExercise  ObservationPassiveRadarObservationListResponseDataMode = "EXERCISE"
)

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
type ObservationPassiveRadarObservationGetResponse struct {
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
	DataMode ObservationPassiveRadarObservationGetResponseDataMode `json:"dataMode,required"`
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
func (r ObservationPassiveRadarObservationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationPassiveRadarObservationGetResponse) UnmarshalJSON(data []byte) error {
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
type ObservationPassiveRadarObservationGetResponseDataMode string

const (
	ObservationPassiveRadarObservationGetResponseDataModeReal      ObservationPassiveRadarObservationGetResponseDataMode = "REAL"
	ObservationPassiveRadarObservationGetResponseDataModeTest      ObservationPassiveRadarObservationGetResponseDataMode = "TEST"
	ObservationPassiveRadarObservationGetResponseDataModeSimulated ObservationPassiveRadarObservationGetResponseDataMode = "SIMULATED"
	ObservationPassiveRadarObservationGetResponseDataModeExercise  ObservationPassiveRadarObservationGetResponseDataMode = "EXERCISE"
)

type ObservationPassiveRadarObservationQueryhelpResponse struct {
	AodrSupported         bool                         `json:"aodrSupported"`
	ClassificationMarking string                       `json:"classificationMarking"`
	Description           string                       `json:"description"`
	HistorySupported      bool                         `json:"historySupported"`
	Name                  string                       `json:"name"`
	Parameters            []shared.ParamDescriptorResp `json:"parameters"`
	RequiredRoles         []string                     `json:"requiredRoles"`
	RestSupported         bool                         `json:"restSupported"`
	SortSupported         bool                         `json:"sortSupported"`
	TypeName              string                       `json:"typeName"`
	Uri                   string                       `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AodrSupported         respjson.Field
		ClassificationMarking respjson.Field
		Description           respjson.Field
		HistorySupported      respjson.Field
		Name                  respjson.Field
		Parameters            respjson.Field
		RequiredRoles         respjson.Field
		RestSupported         respjson.Field
		SortSupported         respjson.Field
		TypeName              respjson.Field
		Uri                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObservationPassiveRadarObservationQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationPassiveRadarObservationQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
type ObservationPassiveRadarObservationTupleResponse struct {
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
	DataMode ObservationPassiveRadarObservationTupleResponseDataMode `json:"dataMode,required"`
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
func (r ObservationPassiveRadarObservationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationPassiveRadarObservationTupleResponse) UnmarshalJSON(data []byte) error {
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
type ObservationPassiveRadarObservationTupleResponseDataMode string

const (
	ObservationPassiveRadarObservationTupleResponseDataModeReal      ObservationPassiveRadarObservationTupleResponseDataMode = "REAL"
	ObservationPassiveRadarObservationTupleResponseDataModeTest      ObservationPassiveRadarObservationTupleResponseDataMode = "TEST"
	ObservationPassiveRadarObservationTupleResponseDataModeSimulated ObservationPassiveRadarObservationTupleResponseDataMode = "SIMULATED"
	ObservationPassiveRadarObservationTupleResponseDataModeExercise  ObservationPassiveRadarObservationTupleResponseDataMode = "EXERCISE"
)

type ObservationPassiveRadarObservationNewParams struct {
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
	DataMode ObservationPassiveRadarObservationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The target Acceleration measurement in kilometers/sec^2 for this observation.
	Accel param.Opt[float64] `json:"accel,omitzero"`
	// The target Acceleration uncertainty measurement in kilometers/sec^2 for this
	// observation.
	AccelUnc param.Opt[float64] `json:"accelUnc,omitzero"`
	// The target altitude relative to WGS-84 ellipsoid, in kilometers for this
	// observation.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Line of sight azimuth angle in degrees and topocentric frame.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias param.Opt[float64] `json:"azimuthBias,omitzero"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// Target bistatic path distance in kilometers. This is the
	// transmitter-to-target-to-surveillance site distance.
	BistaticRange param.Opt[float64] `json:"bistaticRange,omitzero"`
	// Bistatic range acceleration in kilometers/sec^2.
	BistaticRangeAccel param.Opt[float64] `json:"bistaticRangeAccel,omitzero"`
	// One sigma uncertainty in the bistatic range acceleration measurement, in
	// kilometers/sec^2.
	BistaticRangeAccelUnc param.Opt[float64] `json:"bistaticRangeAccelUnc,omitzero"`
	// Sensor bistatic range bias in kilometers.
	BistaticRangeBias param.Opt[float64] `json:"bistaticRangeBias,omitzero"`
	// Rate of change of the bistatic path in kilometers/sec.
	BistaticRangeRate param.Opt[float64] `json:"bistaticRangeRate,omitzero"`
	// One sigma uncertainty in rate of change of the bistatic path in kilometers/sec.
	BistaticRangeRateUnc param.Opt[float64] `json:"bistaticRangeRateUnc,omitzero"`
	// One sigma uncertainty in bistatic range in kilometers.
	BistaticRangeUnc param.Opt[float64] `json:"bistaticRangeUnc,omitzero"`
	// Coning angle in degrees.
	Coning param.Opt[float64] `json:"coning,omitzero"`
	// One sigma uncertainty in the coning angle measurement, in degrees.
	ConingUnc param.Opt[float64] `json:"coningUnc,omitzero"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// The time difference, in seconds, between the signal collected at the
	// surveillance site (after being reflected from the target) and the reference site
	// (direct path line-of-sight signal).
	Delay param.Opt[float64] `json:"delay,omitzero"`
	// Delay bias in seconds.
	DelayBias param.Opt[float64] `json:"delayBias,omitzero"`
	// One sigma uncertainty in the delay measurement, in seconds.
	DelayUnc param.Opt[float64] `json:"delayUnc,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Doppler measurement in hertz.
	Doppler param.Opt[float64] `json:"doppler,omitzero"`
	// One sigma uncertainty in the Doppler measurement in hertz.
	DopplerUnc param.Opt[float64] `json:"dopplerUnc,omitzero"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Sensor elevation bias in degrees.
	ElevationBias param.Opt[float64] `json:"elevationBias,omitzero"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// Optional external observation identifier provided by the source.
	ExtObservationID param.Opt[string] `json:"extObservationId,omitzero"`
	// Unique identifier of the transmitter. This ID can be used to obtain additional
	// information on an RFEmitter using the 'get by ID' operation (e.g.
	// /udl/rfemitter/{id}). For example, the RFEmitter with idRFEmitter = abc would be
	// queried as /udl/rfemitter/abc.
	IDRfEmitter param.Opt[string] `json:"idRFEmitter,omitzero"`
	// Unique identifier of the reporting surveillance sensor. This ID can be used to
	// obtain additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Unique identifier of the reference receiver sensor. This ID can be used to
	// obtain additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensorRefReceiver param.Opt[string] `json:"idSensorRefReceiver,omitzero"`
	// WGS-84 target latitude sub-point at observation time (obTime), represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 target longitude sub-point at observation time (obTime), represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The position of this observation within a track (FENCE, FIRST, IN, LAST,
	// SINGLE). This identifier is optional and, if null, no assumption should be made
	// regarding whether other observations may or may not exist to compose a track.
	ObPosition param.Opt[string] `json:"obPosition,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Radar cross section in meters squared for orthogonal polarization.
	OrthogonalRcs param.Opt[float64] `json:"orthogonalRcs,omitzero"`
	// One sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc param.Opt[float64] `json:"orthogonalRcsUnc,omitzero"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Radar cross section in meters squared for polarization principal.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// One sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc param.Opt[float64] `json:"rcsUnc,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Sensor timing bias in seconds.
	TimingBias param.Opt[float64] `json:"timingBias,omitzero"`
	// Time of flight (TOF) in seconds. This is the calculated propagation time from
	// transmitter-to-target-to-surveillance site.
	Tof param.Opt[float64] `json:"tof,omitzero"`
	// The Time of Flight (TOF) bias in seconds.
	TofBias param.Opt[float64] `json:"tofBias,omitzero"`
	// One sigma uncertainty in time of flight in seconds.
	TofUnc param.Opt[float64] `json:"tofUnc,omitzero"`
	// Unique identifier of a track that represents a tracklet for this observation.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// X velocity of target in kilometers/sec in J2000 coordinate frame.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Y velocity of target in kilometers/sec in J2000 coordinate frame.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Z velocity of target in kilometers/sec in J2000 coordinate frame.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ObservationPassiveRadarObservationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObservationPassiveRadarObservationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationPassiveRadarObservationNewParams) UnmarshalJSON(data []byte) error {
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
type ObservationPassiveRadarObservationNewParamsDataMode string

const (
	ObservationPassiveRadarObservationNewParamsDataModeReal      ObservationPassiveRadarObservationNewParamsDataMode = "REAL"
	ObservationPassiveRadarObservationNewParamsDataModeTest      ObservationPassiveRadarObservationNewParamsDataMode = "TEST"
	ObservationPassiveRadarObservationNewParamsDataModeSimulated ObservationPassiveRadarObservationNewParamsDataMode = "SIMULATED"
	ObservationPassiveRadarObservationNewParamsDataModeExercise  ObservationPassiveRadarObservationNewParamsDataMode = "EXERCISE"
)

type ObservationPassiveRadarObservationListParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationPassiveRadarObservationListParams]'s query
// parameters as `url.Values`.
func (r ObservationPassiveRadarObservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationPassiveRadarObservationCountParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationPassiveRadarObservationCountParams]'s query
// parameters as `url.Values`.
func (r ObservationPassiveRadarObservationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationPassiveRadarObservationNewBulkParams struct {
	Body []ObservationPassiveRadarObservationNewBulkParamsBody
	paramObj
}

func (r ObservationPassiveRadarObservationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationPassiveRadarObservationNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
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
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type ObservationPassiveRadarObservationNewBulkParamsBody struct {
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
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The target Acceleration measurement in kilometers/sec^2 for this observation.
	Accel param.Opt[float64] `json:"accel,omitzero"`
	// The target Acceleration uncertainty measurement in kilometers/sec^2 for this
	// observation.
	AccelUnc param.Opt[float64] `json:"accelUnc,omitzero"`
	// The target altitude relative to WGS-84 ellipsoid, in kilometers for this
	// observation.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Line of sight azimuth angle in degrees and topocentric frame.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias param.Opt[float64] `json:"azimuthBias,omitzero"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// Target bistatic path distance in kilometers. This is the
	// transmitter-to-target-to-surveillance site distance.
	BistaticRange param.Opt[float64] `json:"bistaticRange,omitzero"`
	// Bistatic range acceleration in kilometers/sec^2.
	BistaticRangeAccel param.Opt[float64] `json:"bistaticRangeAccel,omitzero"`
	// One sigma uncertainty in the bistatic range acceleration measurement, in
	// kilometers/sec^2.
	BistaticRangeAccelUnc param.Opt[float64] `json:"bistaticRangeAccelUnc,omitzero"`
	// Sensor bistatic range bias in kilometers.
	BistaticRangeBias param.Opt[float64] `json:"bistaticRangeBias,omitzero"`
	// Rate of change of the bistatic path in kilometers/sec.
	BistaticRangeRate param.Opt[float64] `json:"bistaticRangeRate,omitzero"`
	// One sigma uncertainty in rate of change of the bistatic path in kilometers/sec.
	BistaticRangeRateUnc param.Opt[float64] `json:"bistaticRangeRateUnc,omitzero"`
	// One sigma uncertainty in bistatic range in kilometers.
	BistaticRangeUnc param.Opt[float64] `json:"bistaticRangeUnc,omitzero"`
	// Coning angle in degrees.
	Coning param.Opt[float64] `json:"coning,omitzero"`
	// One sigma uncertainty in the coning angle measurement, in degrees.
	ConingUnc param.Opt[float64] `json:"coningUnc,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// The time difference, in seconds, between the signal collected at the
	// surveillance site (after being reflected from the target) and the reference site
	// (direct path line-of-sight signal).
	Delay param.Opt[float64] `json:"delay,omitzero"`
	// Delay bias in seconds.
	DelayBias param.Opt[float64] `json:"delayBias,omitzero"`
	// One sigma uncertainty in the delay measurement, in seconds.
	DelayUnc param.Opt[float64] `json:"delayUnc,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Doppler measurement in hertz.
	Doppler param.Opt[float64] `json:"doppler,omitzero"`
	// One sigma uncertainty in the Doppler measurement in hertz.
	DopplerUnc param.Opt[float64] `json:"dopplerUnc,omitzero"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Sensor elevation bias in degrees.
	ElevationBias param.Opt[float64] `json:"elevationBias,omitzero"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// Optional external observation identifier provided by the source.
	ExtObservationID param.Opt[string] `json:"extObservationId,omitzero"`
	// Unique identifier of the target satellite on-orbit object. This ID can be used
	// to obtain additional information on an OnOrbit object using the 'get by ID'
	// operation (e.g. /udl/onorbit/{id}). For example, the OnOrbit with idOnOrbit =
	// 25544 would be queried as /udl/onorbit/25544.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the transmitter. This ID can be used to obtain additional
	// information on an RFEmitter using the 'get by ID' operation (e.g.
	// /udl/rfemitter/{id}). For example, the RFEmitter with idRFEmitter = abc would be
	// queried as /udl/rfemitter/abc.
	IDRfEmitter param.Opt[string] `json:"idRFEmitter,omitzero"`
	// Unique identifier of the reporting surveillance sensor. This ID can be used to
	// obtain additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Unique identifier of the reference receiver sensor. This ID can be used to
	// obtain additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensorRefReceiver param.Opt[string] `json:"idSensorRefReceiver,omitzero"`
	// WGS-84 target latitude sub-point at observation time (obTime), represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 target longitude sub-point at observation time (obTime), represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The position of this observation within a track (FENCE, FIRST, IN, LAST,
	// SINGLE). This identifier is optional and, if null, no assumption should be made
	// regarding whether other observations may or may not exist to compose a track.
	ObPosition param.Opt[string] `json:"obPosition,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Radar cross section in meters squared for orthogonal polarization.
	OrthogonalRcs param.Opt[float64] `json:"orthogonalRcs,omitzero"`
	// One sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc param.Opt[float64] `json:"orthogonalRcsUnc,omitzero"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Radar cross section in meters squared for polarization principal.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// One sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc param.Opt[float64] `json:"rcsUnc,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Sensor timing bias in seconds.
	TimingBias param.Opt[float64] `json:"timingBias,omitzero"`
	// Time of flight (TOF) in seconds. This is the calculated propagation time from
	// transmitter-to-target-to-surveillance site.
	Tof param.Opt[float64] `json:"tof,omitzero"`
	// The Time of Flight (TOF) bias in seconds.
	TofBias param.Opt[float64] `json:"tofBias,omitzero"`
	// One sigma uncertainty in time of flight in seconds.
	TofUnc param.Opt[float64] `json:"tofUnc,omitzero"`
	// Unique identifier of a track that represents a tracklet for this observation.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Read only enumeration specifying the type of observation (e.g. OPTICAL, RADAR,
	// RF, etc).
	Type param.Opt[string] `json:"type,omitzero"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// X velocity of target in kilometers/sec in J2000 coordinate frame.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Y velocity of target in kilometers/sec in J2000 coordinate frame.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Z velocity of target in kilometers/sec in J2000 coordinate frame.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ObservationPassiveRadarObservationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationPassiveRadarObservationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationPassiveRadarObservationNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationPassiveRadarObservationNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type ObservationPassiveRadarObservationFileNewParams struct {
	Body []ObservationPassiveRadarObservationFileNewParamsBody
	paramObj
}

func (r ObservationPassiveRadarObservationFileNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationPassiveRadarObservationFileNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
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
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type ObservationPassiveRadarObservationFileNewParamsBody struct {
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
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The target Acceleration measurement in kilometers/sec^2 for this observation.
	Accel param.Opt[float64] `json:"accel,omitzero"`
	// The target Acceleration uncertainty measurement in kilometers/sec^2 for this
	// observation.
	AccelUnc param.Opt[float64] `json:"accelUnc,omitzero"`
	// The target altitude relative to WGS-84 ellipsoid, in kilometers for this
	// observation.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Line of sight azimuth angle in degrees and topocentric frame.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias param.Opt[float64] `json:"azimuthBias,omitzero"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// Target bistatic path distance in kilometers. This is the
	// transmitter-to-target-to-surveillance site distance.
	BistaticRange param.Opt[float64] `json:"bistaticRange,omitzero"`
	// Bistatic range acceleration in kilometers/sec^2.
	BistaticRangeAccel param.Opt[float64] `json:"bistaticRangeAccel,omitzero"`
	// One sigma uncertainty in the bistatic range acceleration measurement, in
	// kilometers/sec^2.
	BistaticRangeAccelUnc param.Opt[float64] `json:"bistaticRangeAccelUnc,omitzero"`
	// Sensor bistatic range bias in kilometers.
	BistaticRangeBias param.Opt[float64] `json:"bistaticRangeBias,omitzero"`
	// Rate of change of the bistatic path in kilometers/sec.
	BistaticRangeRate param.Opt[float64] `json:"bistaticRangeRate,omitzero"`
	// One sigma uncertainty in rate of change of the bistatic path in kilometers/sec.
	BistaticRangeRateUnc param.Opt[float64] `json:"bistaticRangeRateUnc,omitzero"`
	// One sigma uncertainty in bistatic range in kilometers.
	BistaticRangeUnc param.Opt[float64] `json:"bistaticRangeUnc,omitzero"`
	// Coning angle in degrees.
	Coning param.Opt[float64] `json:"coning,omitzero"`
	// One sigma uncertainty in the coning angle measurement, in degrees.
	ConingUnc param.Opt[float64] `json:"coningUnc,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// The time difference, in seconds, between the signal collected at the
	// surveillance site (after being reflected from the target) and the reference site
	// (direct path line-of-sight signal).
	Delay param.Opt[float64] `json:"delay,omitzero"`
	// Delay bias in seconds.
	DelayBias param.Opt[float64] `json:"delayBias,omitzero"`
	// One sigma uncertainty in the delay measurement, in seconds.
	DelayUnc param.Opt[float64] `json:"delayUnc,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Doppler measurement in hertz.
	Doppler param.Opt[float64] `json:"doppler,omitzero"`
	// One sigma uncertainty in the Doppler measurement in hertz.
	DopplerUnc param.Opt[float64] `json:"dopplerUnc,omitzero"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Sensor elevation bias in degrees.
	ElevationBias param.Opt[float64] `json:"elevationBias,omitzero"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// Optional external observation identifier provided by the source.
	ExtObservationID param.Opt[string] `json:"extObservationId,omitzero"`
	// Unique identifier of the target satellite on-orbit object. This ID can be used
	// to obtain additional information on an OnOrbit object using the 'get by ID'
	// operation (e.g. /udl/onorbit/{id}). For example, the OnOrbit with idOnOrbit =
	// 25544 would be queried as /udl/onorbit/25544.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the transmitter. This ID can be used to obtain additional
	// information on an RFEmitter using the 'get by ID' operation (e.g.
	// /udl/rfemitter/{id}). For example, the RFEmitter with idRFEmitter = abc would be
	// queried as /udl/rfemitter/abc.
	IDRfEmitter param.Opt[string] `json:"idRFEmitter,omitzero"`
	// Unique identifier of the reporting surveillance sensor. This ID can be used to
	// obtain additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Unique identifier of the reference receiver sensor. This ID can be used to
	// obtain additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensorRefReceiver param.Opt[string] `json:"idSensorRefReceiver,omitzero"`
	// WGS-84 target latitude sub-point at observation time (obTime), represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 target longitude sub-point at observation time (obTime), represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The position of this observation within a track (FENCE, FIRST, IN, LAST,
	// SINGLE). This identifier is optional and, if null, no assumption should be made
	// regarding whether other observations may or may not exist to compose a track.
	ObPosition param.Opt[string] `json:"obPosition,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Radar cross section in meters squared for orthogonal polarization.
	OrthogonalRcs param.Opt[float64] `json:"orthogonalRcs,omitzero"`
	// One sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc param.Opt[float64] `json:"orthogonalRcsUnc,omitzero"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Radar cross section in meters squared for polarization principal.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// One sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc param.Opt[float64] `json:"rcsUnc,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Sensor timing bias in seconds.
	TimingBias param.Opt[float64] `json:"timingBias,omitzero"`
	// Time of flight (TOF) in seconds. This is the calculated propagation time from
	// transmitter-to-target-to-surveillance site.
	Tof param.Opt[float64] `json:"tof,omitzero"`
	// The Time of Flight (TOF) bias in seconds.
	TofBias param.Opt[float64] `json:"tofBias,omitzero"`
	// One sigma uncertainty in time of flight in seconds.
	TofUnc param.Opt[float64] `json:"tofUnc,omitzero"`
	// Unique identifier of a track that represents a tracklet for this observation.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Read only enumeration specifying the type of observation (e.g. OPTICAL, RADAR,
	// RF, etc).
	Type param.Opt[string] `json:"type,omitzero"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// X velocity of target in kilometers/sec in J2000 coordinate frame.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Y velocity of target in kilometers/sec in J2000 coordinate frame.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Z velocity of target in kilometers/sec in J2000 coordinate frame.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ObservationPassiveRadarObservationFileNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationPassiveRadarObservationFileNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationPassiveRadarObservationFileNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationPassiveRadarObservationFileNewParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type ObservationPassiveRadarObservationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationPassiveRadarObservationGetParams]'s query
// parameters as `url.Values`.
func (r ObservationPassiveRadarObservationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationPassiveRadarObservationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationPassiveRadarObservationTupleParams]'s query
// parameters as `url.Values`.
func (r ObservationPassiveRadarObservationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
