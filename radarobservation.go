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

// RadarobservationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarobservationService] method instead.
type RadarobservationService struct {
	Options []option.RequestOption
}

// NewRadarobservationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarobservationService(opts ...option.RequestOption) (r RadarobservationService) {
	r = RadarobservationService{}
	r.Options = opts
	return
}

// Service operation to take a single radar observation as a POST body and ingest
// into the database. This operation is not intended to be used for automated feeds
// into UDL. Data providers should contact the UDL team for specific role
// assignments and for instructions on setting up a permanent feed through an
// alternate mechanism.
func (r *RadarobservationService) New(ctx context.Context, body RadarobservationNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/radarobservation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *RadarobservationService) List(ctx context.Context, query RadarobservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[RadarobservationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/radarobservation"
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
func (r *RadarobservationService) ListAutoPaging(ctx context.Context, query RadarobservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[RadarobservationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *RadarobservationService) Count(ctx context.Context, query RadarobservationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/radarobservation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of radar
// observations as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *RadarobservationService) NewBulk(ctx context.Context, body RadarobservationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/radarobservation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single radar observations by its unique ID passed as
// a path parameter.
func (r *RadarobservationService) Get(ctx context.Context, id string, query RadarobservationGetParams, opts ...option.RequestOption) (res *RadarobservationFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/radarobservation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *RadarobservationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/radarobservation/queryhelp"
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
func (r *RadarobservationService) Tuple(ctx context.Context, query RadarobservationTupleParams, opts ...option.RequestOption) (res *[]RadarobservationFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/radarobservation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple radar observations as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RadarobservationService) UnvalidatedPublish(ctx context.Context, body RadarobservationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-radar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Model representation of observation data for radar based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
type RadarobservationListResponse struct {
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
	DataMode RadarobservationListResponseDataMode `json:"dataMode,required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// azimuth angle in degrees and topocentric frame.
	Azimuth float64 `json:"azimuth"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias float64 `json:"azimuthBias"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured bool `json:"azimuthMeasured"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// ID of the beam that produced this observation.
	Beam float64 `json:"beam"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination float64 `json:"declination"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured bool `json:"declinationMeasured"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Corrected doppler measurement in meters per second.
	Doppler float64 `json:"doppler"`
	// One sigma uncertainty in the corrected doppler measurement, in meters/second.
	DopplerUnc float64 `json:"dopplerUnc"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation float64 `json:"elevation"`
	// Sensor elevation bias in degrees.
	ElevationBias float64 `json:"elevationBias"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured bool `json:"elevationMeasured"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate float64 `json:"elevationRate"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc float64 `json:"elevationUnc"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
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
	// one sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc float64 `json:"orthogonalRcsUnc"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra float64 `json:"ra"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured bool `json:"raMeasured"`
	// Target range in km.
	Range float64 `json:"range"`
	// Range accelaration in km/s2.
	RangeAccel float64 `json:"rangeAccel"`
	// One sigma uncertainty in the range acceleration measurement, in
	// kilometers/(second^2).
	RangeAccelUnc float64 `json:"rangeAccelUnc"`
	// Sensor range bias in km.
	RangeBias float64 `json:"rangeBias"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured bool `json:"rangeMeasured"`
	// Rate of change of the line of sight range in km/sec.
	RangeRate float64 `json:"rangeRate"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured bool `json:"rangeRateMeasured"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc float64 `json:"rangeRateUnc"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc float64 `json:"rangeUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Radar cross section in meters squared for polarization principal.
	Rcs float64 `json:"rcs"`
	// one sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc float64 `json:"rcsUnc"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame RadarobservationListResponseSenReferenceFrame `json:"senReferenceFrame"`
	// Sensor x position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senx float64 `json:"senx"`
	// Sensor y position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Seny float64 `json:"seny"`
	// Sensor z position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senz float64 `json:"senz"`
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
	// Optional identifier of the track to which this observation belongs.
	TrackID string `json:"trackId"`
	// The beam type (or tracking state) in use at the time of collection of this
	// observation. Values include:
	//
	// INIT ACQ WITH INIT VALUES: Initial acquisition based on predefined initial
	// values such as position, velocity, or other specific parameters.
	//
	// INIT ACQ: Initial acquisition when no prior information or initial values such
	// as position or velocity are available.
	//
	// TRACKING SINGLE BEAM: Continuously tracks and monitors a single target using one
	// specific radar beam.
	//
	// TRACKING SEQUENTIAL ROVING: Sequentially tracks different targets or areas by
	// "roving" from one sector to the next in a systematic order.
	//
	// SELF ACQ WITH INIT VALUES: Autonomously acquires targets using predefined
	// starting parameters or initial values.
	//
	// SELF ACQ: Automatically detects and locks onto targets without the need for
	// predefined initial settings.
	//
	// NON-TRACKING: Non-tracking.
	TrackingState string `json:"trackingState"`
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
	// X position of target in km in J2000 coordinate frame.
	X float64 `json:"x"`
	// X velocity of target in km/sec in J2000 coordinate frame.
	Xvel float64 `json:"xvel"`
	// Y position of target in km in J2000 coordinate frame.
	Y float64 `json:"y"`
	// Y velocity of target in km/sec in J2000 coordinate frame.
	Yvel float64 `json:"yvel"`
	// Z position of target in km in J2000 coordinate frame.
	Z float64 `json:"z"`
	// Z velocity of target in km/sec in J2000 coordinate frame.
	Zvel float64 `json:"zvel"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		ObTime                resp.Field
		Source                resp.Field
		ID                    resp.Field
		Azimuth               resp.Field
		AzimuthBias           resp.Field
		AzimuthMeasured       resp.Field
		AzimuthRate           resp.Field
		AzimuthUnc            resp.Field
		Beam                  resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Declination           resp.Field
		DeclinationMeasured   resp.Field
		Descriptor            resp.Field
		Doppler               resp.Field
		DopplerUnc            resp.Field
		Elevation             resp.Field
		ElevationBias         resp.Field
		ElevationMeasured     resp.Field
		ElevationRate         resp.Field
		ElevationUnc          resp.Field
		IDOnOrbit             resp.Field
		IDSensor              resp.Field
		ObPosition            resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		OrigSensorID          resp.Field
		OrthogonalRcs         resp.Field
		OrthogonalRcsUnc      resp.Field
		Ra                    resp.Field
		RaMeasured            resp.Field
		Range                 resp.Field
		RangeAccel            resp.Field
		RangeAccelUnc         resp.Field
		RangeBias             resp.Field
		RangeMeasured         resp.Field
		RangeRate             resp.Field
		RangeRateMeasured     resp.Field
		RangeRateUnc          resp.Field
		RangeUnc              resp.Field
		RawFileUri            resp.Field
		Rcs                   resp.Field
		RcsUnc                resp.Field
		SatNo                 resp.Field
		SenReferenceFrame     resp.Field
		Senx                  resp.Field
		Seny                  resp.Field
		Senz                  resp.Field
		Snr                   resp.Field
		SourceDl              resp.Field
		TaskID                resp.Field
		TimingBias            resp.Field
		TrackID               resp.Field
		TrackingState         resp.Field
		TransactionID         resp.Field
		Type                  resp.Field
		Uct                   resp.Field
		X                     resp.Field
		Xvel                  resp.Field
		Y                     resp.Field
		Yvel                  resp.Field
		Z                     resp.Field
		Zvel                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RadarobservationListResponse) RawJSON() string { return r.JSON.raw }
func (r *RadarobservationListResponse) UnmarshalJSON(data []byte) error {
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
type RadarobservationListResponseDataMode string

const (
	RadarobservationListResponseDataModeReal      RadarobservationListResponseDataMode = "REAL"
	RadarobservationListResponseDataModeTest      RadarobservationListResponseDataMode = "TEST"
	RadarobservationListResponseDataModeSimulated RadarobservationListResponseDataMode = "SIMULATED"
	RadarobservationListResponseDataModeExercise  RadarobservationListResponseDataMode = "EXERCISE"
)

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type RadarobservationListResponseSenReferenceFrame string

const (
	RadarobservationListResponseSenReferenceFrameJ2000   RadarobservationListResponseSenReferenceFrame = "J2000"
	RadarobservationListResponseSenReferenceFrameEfgTdr  RadarobservationListResponseSenReferenceFrame = "EFG/TDR"
	RadarobservationListResponseSenReferenceFrameEcrEcef RadarobservationListResponseSenReferenceFrame = "ECR/ECEF"
	RadarobservationListResponseSenReferenceFrameTeme    RadarobservationListResponseSenReferenceFrame = "TEME"
	RadarobservationListResponseSenReferenceFrameItrf    RadarobservationListResponseSenReferenceFrame = "ITRF"
	RadarobservationListResponseSenReferenceFrameGcrf    RadarobservationListResponseSenReferenceFrame = "GCRF"
)

type RadarobservationNewParams struct {
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
	DataMode RadarobservationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// azimuth angle in degrees and topocentric frame.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias param.Opt[float64] `json:"azimuthBias,omitzero"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured param.Opt[bool] `json:"azimuthMeasured,omitzero"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// ID of the beam that produced this observation.
	Beam param.Opt[float64] `json:"beam,omitzero"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured param.Opt[bool] `json:"declinationMeasured,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Corrected doppler measurement in meters per second.
	Doppler param.Opt[float64] `json:"doppler,omitzero"`
	// One sigma uncertainty in the corrected doppler measurement, in meters/second.
	DopplerUnc param.Opt[float64] `json:"dopplerUnc,omitzero"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Sensor elevation bias in degrees.
	ElevationBias param.Opt[float64] `json:"elevationBias,omitzero"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured param.Opt[bool] `json:"elevationMeasured,omitzero"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
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
	// one sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc param.Opt[float64] `json:"orthogonalRcsUnc,omitzero"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured param.Opt[bool] `json:"raMeasured,omitzero"`
	// Target range in km.
	Range param.Opt[float64] `json:"range,omitzero"`
	// Range accelaration in km/s2.
	RangeAccel param.Opt[float64] `json:"rangeAccel,omitzero"`
	// One sigma uncertainty in the range acceleration measurement, in
	// kilometers/(second^2).
	RangeAccelUnc param.Opt[float64] `json:"rangeAccelUnc,omitzero"`
	// Sensor range bias in km.
	RangeBias param.Opt[float64] `json:"rangeBias,omitzero"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured param.Opt[bool] `json:"rangeMeasured,omitzero"`
	// Rate of change of the line of sight range in km/sec.
	RangeRate param.Opt[float64] `json:"rangeRate,omitzero"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured param.Opt[bool] `json:"rangeRateMeasured,omitzero"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc param.Opt[float64] `json:"rangeRateUnc,omitzero"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc param.Opt[float64] `json:"rangeUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Radar cross section in meters squared for polarization principal.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// one sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc param.Opt[float64] `json:"rcsUnc,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor x position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senx param.Opt[float64] `json:"senx,omitzero"`
	// Sensor y position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Seny param.Opt[float64] `json:"seny,omitzero"`
	// Sensor z position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senz param.Opt[float64] `json:"senz,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Sensor timing bias in seconds.
	TimingBias param.Opt[float64] `json:"timingBias,omitzero"`
	// Optional identifier of the track to which this observation belongs.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// The beam type (or tracking state) in use at the time of collection of this
	// observation. Values include:
	//
	// INIT ACQ WITH INIT VALUES: Initial acquisition based on predefined initial
	// values such as position, velocity, or other specific parameters.
	//
	// INIT ACQ: Initial acquisition when no prior information or initial values such
	// as position or velocity are available.
	//
	// TRACKING SINGLE BEAM: Continuously tracks and monitors a single target using one
	// specific radar beam.
	//
	// TRACKING SEQUENTIAL ROVING: Sequentially tracks different targets or areas by
	// "roving" from one sector to the next in a systematic order.
	//
	// SELF ACQ WITH INIT VALUES: Autonomously acquires targets using predefined
	// starting parameters or initial values.
	//
	// SELF ACQ: Automatically detects and locks onto targets without the need for
	// predefined initial settings.
	//
	// NON-TRACKING: Non-tracking.
	TrackingState param.Opt[string] `json:"trackingState,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// X position of target in km in J2000 coordinate frame.
	X param.Opt[float64] `json:"x,omitzero"`
	// X velocity of target in km/sec in J2000 coordinate frame.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Y position of target in km in J2000 coordinate frame.
	Y param.Opt[float64] `json:"y,omitzero"`
	// Y velocity of target in km/sec in J2000 coordinate frame.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Z position of target in km in J2000 coordinate frame.
	Z param.Opt[float64] `json:"z,omitzero"`
	// Z velocity of target in km/sec in J2000 coordinate frame.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame RadarobservationNewParamsSenReferenceFrame `json:"senReferenceFrame,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RadarobservationNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RadarobservationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RadarobservationNewParams
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
type RadarobservationNewParamsDataMode string

const (
	RadarobservationNewParamsDataModeReal      RadarobservationNewParamsDataMode = "REAL"
	RadarobservationNewParamsDataModeTest      RadarobservationNewParamsDataMode = "TEST"
	RadarobservationNewParamsDataModeSimulated RadarobservationNewParamsDataMode = "SIMULATED"
	RadarobservationNewParamsDataModeExercise  RadarobservationNewParamsDataMode = "EXERCISE"
)

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type RadarobservationNewParamsSenReferenceFrame string

const (
	RadarobservationNewParamsSenReferenceFrameJ2000   RadarobservationNewParamsSenReferenceFrame = "J2000"
	RadarobservationNewParamsSenReferenceFrameEfgTdr  RadarobservationNewParamsSenReferenceFrame = "EFG/TDR"
	RadarobservationNewParamsSenReferenceFrameEcrEcef RadarobservationNewParamsSenReferenceFrame = "ECR/ECEF"
	RadarobservationNewParamsSenReferenceFrameTeme    RadarobservationNewParamsSenReferenceFrame = "TEME"
	RadarobservationNewParamsSenReferenceFrameItrf    RadarobservationNewParamsSenReferenceFrame = "ITRF"
	RadarobservationNewParamsSenReferenceFrameGcrf    RadarobservationNewParamsSenReferenceFrame = "GCRF"
)

type RadarobservationListParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RadarobservationListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RadarobservationListParams]'s query parameters as
// `url.Values`.
func (r RadarobservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarobservationCountParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RadarobservationCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RadarobservationCountParams]'s query parameters as
// `url.Values`.
func (r RadarobservationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarobservationNewBulkParams struct {
	Body []RadarobservationNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RadarobservationNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RadarobservationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Model representation of observation data for radar based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type RadarobservationNewBulkParamsBody struct {
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
	// azimuth angle in degrees and topocentric frame.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias param.Opt[float64] `json:"azimuthBias,omitzero"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured param.Opt[bool] `json:"azimuthMeasured,omitzero"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// ID of the beam that produced this observation.
	Beam param.Opt[float64] `json:"beam,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured param.Opt[bool] `json:"declinationMeasured,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Corrected doppler measurement in meters per second.
	Doppler param.Opt[float64] `json:"doppler,omitzero"`
	// One sigma uncertainty in the corrected doppler measurement, in meters/second.
	DopplerUnc param.Opt[float64] `json:"dopplerUnc,omitzero"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Sensor elevation bias in degrees.
	ElevationBias param.Opt[float64] `json:"elevationBias,omitzero"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured param.Opt[bool] `json:"elevationMeasured,omitzero"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
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
	// one sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc param.Opt[float64] `json:"orthogonalRcsUnc,omitzero"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured param.Opt[bool] `json:"raMeasured,omitzero"`
	// Target range in km.
	Range param.Opt[float64] `json:"range,omitzero"`
	// Range accelaration in km/s2.
	RangeAccel param.Opt[float64] `json:"rangeAccel,omitzero"`
	// One sigma uncertainty in the range acceleration measurement, in
	// kilometers/(second^2).
	RangeAccelUnc param.Opt[float64] `json:"rangeAccelUnc,omitzero"`
	// Sensor range bias in km.
	RangeBias param.Opt[float64] `json:"rangeBias,omitzero"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured param.Opt[bool] `json:"rangeMeasured,omitzero"`
	// Rate of change of the line of sight range in km/sec.
	RangeRate param.Opt[float64] `json:"rangeRate,omitzero"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured param.Opt[bool] `json:"rangeRateMeasured,omitzero"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc param.Opt[float64] `json:"rangeRateUnc,omitzero"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc param.Opt[float64] `json:"rangeUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Radar cross section in meters squared for polarization principal.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// one sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc param.Opt[float64] `json:"rcsUnc,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor x position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senx param.Opt[float64] `json:"senx,omitzero"`
	// Sensor y position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Seny param.Opt[float64] `json:"seny,omitzero"`
	// Sensor z position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senz param.Opt[float64] `json:"senz,omitzero"`
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
	// Optional identifier of the track to which this observation belongs.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// The beam type (or tracking state) in use at the time of collection of this
	// observation. Values include:
	//
	// INIT ACQ WITH INIT VALUES: Initial acquisition based on predefined initial
	// values such as position, velocity, or other specific parameters.
	//
	// INIT ACQ: Initial acquisition when no prior information or initial values such
	// as position or velocity are available.
	//
	// TRACKING SINGLE BEAM: Continuously tracks and monitors a single target using one
	// specific radar beam.
	//
	// TRACKING SEQUENTIAL ROVING: Sequentially tracks different targets or areas by
	// "roving" from one sector to the next in a systematic order.
	//
	// SELF ACQ WITH INIT VALUES: Autonomously acquires targets using predefined
	// starting parameters or initial values.
	//
	// SELF ACQ: Automatically detects and locks onto targets without the need for
	// predefined initial settings.
	//
	// NON-TRACKING: Non-tracking.
	TrackingState param.Opt[string] `json:"trackingState,omitzero"`
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
	// X position of target in km in J2000 coordinate frame.
	X param.Opt[float64] `json:"x,omitzero"`
	// X velocity of target in km/sec in J2000 coordinate frame.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Y position of target in km in J2000 coordinate frame.
	Y param.Opt[float64] `json:"y,omitzero"`
	// Y velocity of target in km/sec in J2000 coordinate frame.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Z position of target in km in J2000 coordinate frame.
	Z param.Opt[float64] `json:"z,omitzero"`
	// Z velocity of target in km/sec in J2000 coordinate frame.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame string `json:"senReferenceFrame,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RadarobservationNewBulkParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r RadarobservationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow RadarobservationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[RadarobservationNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[RadarobservationNewBulkParamsBody](
		"SenReferenceFrame", false, "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

type RadarobservationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RadarobservationGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RadarobservationGetParams]'s query parameters as
// `url.Values`.
func (r RadarobservationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarobservationTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RadarobservationTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RadarobservationTupleParams]'s query parameters as
// `url.Values`.
func (r RadarobservationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarobservationUnvalidatedPublishParams struct {
	Body []RadarobservationUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RadarobservationUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r RadarobservationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Model representation of observation data for radar based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type RadarobservationUnvalidatedPublishParamsBody struct {
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
	// azimuth angle in degrees and topocentric frame.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias param.Opt[float64] `json:"azimuthBias,omitzero"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured param.Opt[bool] `json:"azimuthMeasured,omitzero"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// ID of the beam that produced this observation.
	Beam param.Opt[float64] `json:"beam,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured param.Opt[bool] `json:"declinationMeasured,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Corrected doppler measurement in meters per second.
	Doppler param.Opt[float64] `json:"doppler,omitzero"`
	// One sigma uncertainty in the corrected doppler measurement, in meters/second.
	DopplerUnc param.Opt[float64] `json:"dopplerUnc,omitzero"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Sensor elevation bias in degrees.
	ElevationBias param.Opt[float64] `json:"elevationBias,omitzero"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured param.Opt[bool] `json:"elevationMeasured,omitzero"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
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
	// one sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc param.Opt[float64] `json:"orthogonalRcsUnc,omitzero"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured param.Opt[bool] `json:"raMeasured,omitzero"`
	// Target range in km.
	Range param.Opt[float64] `json:"range,omitzero"`
	// Range accelaration in km/s2.
	RangeAccel param.Opt[float64] `json:"rangeAccel,omitzero"`
	// One sigma uncertainty in the range acceleration measurement, in
	// kilometers/(second^2).
	RangeAccelUnc param.Opt[float64] `json:"rangeAccelUnc,omitzero"`
	// Sensor range bias in km.
	RangeBias param.Opt[float64] `json:"rangeBias,omitzero"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured param.Opt[bool] `json:"rangeMeasured,omitzero"`
	// Rate of change of the line of sight range in km/sec.
	RangeRate param.Opt[float64] `json:"rangeRate,omitzero"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured param.Opt[bool] `json:"rangeRateMeasured,omitzero"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc param.Opt[float64] `json:"rangeRateUnc,omitzero"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc param.Opt[float64] `json:"rangeUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Radar cross section in meters squared for polarization principal.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// one sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc param.Opt[float64] `json:"rcsUnc,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor x position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senx param.Opt[float64] `json:"senx,omitzero"`
	// Sensor y position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Seny param.Opt[float64] `json:"seny,omitzero"`
	// Sensor z position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senz param.Opt[float64] `json:"senz,omitzero"`
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
	// Optional identifier of the track to which this observation belongs.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// The beam type (or tracking state) in use at the time of collection of this
	// observation. Values include:
	//
	// INIT ACQ WITH INIT VALUES: Initial acquisition based on predefined initial
	// values such as position, velocity, or other specific parameters.
	//
	// INIT ACQ: Initial acquisition when no prior information or initial values such
	// as position or velocity are available.
	//
	// TRACKING SINGLE BEAM: Continuously tracks and monitors a single target using one
	// specific radar beam.
	//
	// TRACKING SEQUENTIAL ROVING: Sequentially tracks different targets or areas by
	// "roving" from one sector to the next in a systematic order.
	//
	// SELF ACQ WITH INIT VALUES: Autonomously acquires targets using predefined
	// starting parameters or initial values.
	//
	// SELF ACQ: Automatically detects and locks onto targets without the need for
	// predefined initial settings.
	//
	// NON-TRACKING: Non-tracking.
	TrackingState param.Opt[string] `json:"trackingState,omitzero"`
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
	// X position of target in km in J2000 coordinate frame.
	X param.Opt[float64] `json:"x,omitzero"`
	// X velocity of target in km/sec in J2000 coordinate frame.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Y position of target in km in J2000 coordinate frame.
	Y param.Opt[float64] `json:"y,omitzero"`
	// Y velocity of target in km/sec in J2000 coordinate frame.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Z position of target in km in J2000 coordinate frame.
	Z param.Opt[float64] `json:"z,omitzero"`
	// Z velocity of target in km/sec in J2000 coordinate frame.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame string `json:"senReferenceFrame,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RadarobservationUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r RadarobservationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow RadarobservationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[RadarobservationUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[RadarobservationUnvalidatedPublishParamsBody](
		"SenReferenceFrame", false, "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}
