// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
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

// ObservationEoObservationService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationEoObservationService] method instead.
type ObservationEoObservationService struct {
	Options []option.RequestOption
	History ObservationEoObservationHistoryService
}

// NewObservationEoObservationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservationEoObservationService(opts ...option.RequestOption) (r ObservationEoObservationService) {
	r = ObservationEoObservationService{}
	r.Options = opts
	r.History = NewObservationEoObservationHistoryService(opts...)
	return
}

// Service operation to take a single EO observation as a POST body and ingest into
// the database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *ObservationEoObservationService) New(ctx context.Context, body ObservationEoObservationNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/eoobservation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single EO observation by its unique ID passed as a
// path parameter.
func (r *ObservationEoObservationService) Get(ctx context.Context, id string, query ObservationEoObservationGetParams, opts ...option.RequestOption) (res *shared.EoObservationFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/eoobservation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationEoObservationService) List(ctx context.Context, query ObservationEoObservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EoObservationAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/eoobservation"
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
func (r *ObservationEoObservationService) ListAutoPaging(ctx context.Context, query ObservationEoObservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EoObservationAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationEoObservationService) Count(ctx context.Context, query ObservationEoObservationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/eoobservation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of EO
// observations as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *ObservationEoObservationService) NewBulk(ctx context.Context, params ObservationEoObservationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/eoobservation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ObservationEoObservationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *ObservationEoObservationQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/eoobservation/queryhelp"
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
func (r *ObservationEoObservationService) Tuple(ctx context.Context, query ObservationEoObservationTupleParams, opts ...option.RequestOption) (res *[]shared.EoObservationFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/eoobservation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple EO observations as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *ObservationEoObservationService) UnvalidatedPublish(ctx context.Context, body ObservationEoObservationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-eo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Model representation of observation data for electro-optical based sensor
// phenomenologies. ECI J2K is the preferred reference frame for EOObservations,
// however, several user-specified reference frames are accommodated. Users should
// check the EOObservation record as well as the 'Discover' tab in the storefront
// to confirm the coordinate frames used by the data provider.
type EoObservationAbridged struct {
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
	DataMode EoObservationAbridgedDataMode `json:"dataMode,required"`
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Line of sight azimuth angle in degrees and topocentric frame. Reported value
	// should include all applicable corrections as specified on the source provider
	// data card. If uncertain, consumers should contact the provider for details on
	// the applied corrections.
	Azimuth float64 `json:"azimuth"`
	// Sensor line of sight azimuth angle bias in degrees.
	AzimuthBias float64 `json:"azimuthBias"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured bool `json:"azimuthMeasured"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the line of sight azimuth angle, in degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// Background intensity for IR observations, in kw/sr/um.
	BgIntensity float64 `json:"bgIntensity"`
	// Method indicating telescope movement during collection (AUTOTRACK, MANUAL
	// AUTOTRACK, MANUAL RATE TRACK, MANUAL SIDEREAL, SIDEREAL, RATE TRACK).
	CollectMethod string `json:"collectMethod"`
	// Object Correlation Quality score of the observation when compared to a known
	// orbit state, (non-standardized). Users should consult data providers regarding
	// the expected range of values.
	CorrQuality float64 `json:"corrQuality"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Line of sight declination, in degrees, in the specified referenceFrame. If
	// referenceFrame is null then J2K should be assumed. Reported value should include
	// all applicable corrections as specified on the source provider data card. If
	// uncertain, consumers should contact the provider for details on the applied
	// corrections.
	Declination float64 `json:"declination"`
	// Sensor line of sight declination angle bias in degrees.
	DeclinationBias float64 `json:"declinationBias"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured bool `json:"declinationMeasured"`
	// Line of sight declination rate of change, in degrees/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	DeclinationRate float64 `json:"declinationRate"`
	// One sigma uncertainty in the line of sight declination angle, in degrees.
	DeclinationUnc float64 `json:"declinationUnc"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Line of sight elevation in degrees and topocentric frame. Reported value should
	// include all applicable corrections as specified on the source provider data
	// card. If uncertain, consumers should contact the provider for details on the
	// applied corrections.
	Elevation float64 `json:"elevation"`
	// Sensor line of sight elevation bias in degrees.
	ElevationBias float64 `json:"elevationBias"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured bool `json:"elevationMeasured"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate float64 `json:"elevationRate"`
	// One sigma uncertainty in the line of sight elevation angle, in degrees.
	ElevationUnc float64 `json:"elevationUnc"`
	// Image exposure duration in seconds. For observations performed using frame
	// stacking or synthetic tracking methods, the exposure duration should be the
	// total integration time. This field is highly recommended / required if the
	// observations are going to be used for photometric processing.
	ExpDuration float64 `json:"expDuration"`
	// The number of RSOs detected in the sensor field of view.
	FovCount int64 `json:"fovCount"`
	// The number of uncorrelated tracks in the field of view.
	FovCountUct int64 `json:"fovCountUCT"`
	// For GEO detections, the altitude in km.
	Geoalt float64 `json:"geoalt"`
	// For GEO detections, the latitude in degrees north.
	Geolat float64 `json:"geolat"`
	// For GEO detections, the longitude in degrees east.
	Geolon float64 `json:"geolon"`
	// For GEO detections, the range in km.
	Georange float64 `json:"georange"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// Unique identifier of the Sky Imagery.
	IDSkyImagery string `json:"idSkyImagery"`
	// Intensity of the target for IR observations, in kw/sr/um.
	Intensity float64 `json:"intensity"`
	// One sigma uncertainty in the line of sight pointing in micro-radians.
	LosUnc float64 `json:"losUnc"`
	// Line-of-sight cartesian X position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losx float64 `json:"losx"`
	// Line-of-sight cartesian X velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losxvel float64 `json:"losxvel"`
	// Line-of-sight cartesian Y position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losy float64 `json:"losy"`
	// Line-of-sight cartesian Y velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losyvel float64 `json:"losyvel"`
	// Line-of-sight cartesian Z position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losz float64 `json:"losz"`
	// Line-of-sight cartesian Z velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Loszvel float64 `json:"loszvel"`
	// Measure of observed brightness calibrated against the Gaia G-band in units of
	// magnitudes.
	Mag float64 `json:"mag"`
	// Formula: mag - 5.0 \* log_10(geo_range / 1000000.0).
	MagNormRange float64 `json:"magNormRange"`
	// Uncertainty of the observed brightness in units of magnitudes.
	MagUnc float64 `json:"magUnc"`
	// Net object signature = counts / expDuration.
	NetObjSig float64 `json:"netObjSig"`
	// Net object signature uncertainty = counts uncertainty / expDuration.
	NetObjSigUnc float64 `json:"netObjSigUnc"`
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
	// Boolean indicating that the target object was in a penumbral eclipse at the time
	// of this observation.
	Penumbra bool `json:"penumbra"`
	// Primary Extinction Coefficient, in Magnitudes. Primary Extinction is the
	// coefficient applied to the airmass to determine how much the observed visual
	// magnitude has been attenuated by the atmosphere. Extinction, in general,
	// describes the absorption and scattering of electromagnetic radiation by dust and
	// gas between an emitting astronomical object and the observer. See the
	// EOObservationDetails API for specification of extinction coefficients for
	// multiple spectral filters.
	PrimaryExtinction float64 `json:"primaryExtinction"`
	// Primary Extinction Coefficient Uncertainty, in Magnitudes.
	PrimaryExtinctionUnc float64 `json:"primaryExtinctionUnc"`
	// Line of sight right ascension, in degrees, in the specified referenceFrame. If
	// referenceFrame is null then J2K should be assumed. Reported value should include
	// all applicable corrections as specified on the source provider data card. If
	// uncertain, consumers should contact the provider for details on the applied
	// corrections.
	Ra float64 `json:"ra"`
	// Sensor line of sight right ascension bias in degrees.
	RaBias float64 `json:"raBias"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured bool `json:"raMeasured"`
	// Line of sight range in km. If referenceFrame is null then J2K should be assumed.
	// Reported value should include all applicable corrections as specified on the
	// source provider data card. If uncertain, consumers should contact the provider
	// for details on the applied corrections.
	Range float64 `json:"range"`
	// Sensor line of sight range bias in km.
	RangeBias float64 `json:"rangeBias"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured bool `json:"rangeMeasured"`
	// Range rate in km/s. If referenceFrame is null then J2K should be assumed.
	// Reported value should include all applicable corrections as specified on the
	// source provider data card. If uncertain, consumers should contact the provider
	// for details on the applied corrections.
	RangeRate float64 `json:"rangeRate"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured bool `json:"rangeRateMeasured"`
	// One sigma uncertainty in the line of sight range rate, in kilometers/second.
	RangeRateUnc float64 `json:"rangeRateUnc"`
	// One sigma uncertainty in the line of sight range, in kilometers.
	RangeUnc float64 `json:"rangeUnc"`
	// Line of sight right ascension rate of change, in degrees/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	RaRate float64 `json:"raRate"`
	// One sigma uncertainty in the line of sight right ascension angle, in degrees.
	RaUnc float64 `json:"raUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The reference frame of the EOObservation measurements. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "GCRF", "ITRF", "TEME".
	ReferenceFrame EoObservationAbridgedReferenceFrame `json:"referenceFrame"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sensor altitude at obTime (if mobile/onorbit) in km.
	Senalt float64 `json:"senalt"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon float64 `json:"senlon"`
	// The quaternion describing the rotation of the sensor in relation to the
	// body-fixed frame used for this system into the local geodetic frame, at
	// observation time (obTime). The array element order convention is scalar
	// component first, followed by the three vector components (qc, q1, q2, q3).
	SenQuat []float64 `json:"senQuat"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame EoObservationAbridgedSenReferenceFrame `json:"senReferenceFrame"`
	// Cartesian X velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvelx float64 `json:"senvelx"`
	// Cartesian Y velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvely float64 `json:"senvely"`
	// Cartesian Z velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvelz float64 `json:"senvelz"`
	// Cartesian X position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Senx float64 `json:"senx"`
	// Cartesian Y position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Seny float64 `json:"seny"`
	// Cartesian Z position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Senz float64 `json:"senz"`
	// Shutter delay in seconds.
	ShutterDelay float64 `json:"shutterDelay"`
	// Average Sky Background signal, in Magnitudes. Sky Background refers to the
	// incoming light from an apparently empty part of the night sky.
	SkyBkgrnd float64 `json:"skyBkgrnd"`
	// Angle from the sun to the equatorial plane.
	SolarDecAngle float64 `json:"solarDecAngle"`
	// The angle, in degrees, between the projections of the target-to-observer vector
	// and the target-to-sun vector onto the equatorial plane. The angle is represented
	// as negative when closing (i.e. before the opposition) and positive when opening
	// (after the opposition).
	SolarEqPhaseAngle float64 `json:"solarEqPhaseAngle"`
	// The angle, in degrees, between the target-to-observer vector and the
	// target-to-sun vector.
	SolarPhaseAngle float64 `json:"solarPhaseAngle"`
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
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Read only field specifying the type of observation (e.g. OPTICAL, OPTICAL_IR,
	// LASER_RANGING, etc).
	Type string `json:"type"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct bool `json:"uct"`
	// Boolean indicating that the target object was in umbral eclipse at the time of
	// this observation.
	Umbra bool `json:"umbra"`
	// Formula: 2.5 \* log_10 (zero_mag_counts / expDuration).
	Zeroptd float64 `json:"zeroptd"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Azimuth               respjson.Field
		AzimuthBias           respjson.Field
		AzimuthMeasured       respjson.Field
		AzimuthRate           respjson.Field
		AzimuthUnc            respjson.Field
		BgIntensity           respjson.Field
		CollectMethod         respjson.Field
		CorrQuality           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Declination           respjson.Field
		DeclinationBias       respjson.Field
		DeclinationMeasured   respjson.Field
		DeclinationRate       respjson.Field
		DeclinationUnc        respjson.Field
		Descriptor            respjson.Field
		Elevation             respjson.Field
		ElevationBias         respjson.Field
		ElevationMeasured     respjson.Field
		ElevationRate         respjson.Field
		ElevationUnc          respjson.Field
		ExpDuration           respjson.Field
		FovCount              respjson.Field
		FovCountUct           respjson.Field
		Geoalt                respjson.Field
		Geolat                respjson.Field
		Geolon                respjson.Field
		Georange              respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		IDSkyImagery          respjson.Field
		Intensity             respjson.Field
		LosUnc                respjson.Field
		Losx                  respjson.Field
		Losxvel               respjson.Field
		Losy                  respjson.Field
		Losyvel               respjson.Field
		Losz                  respjson.Field
		Loszvel               respjson.Field
		Mag                   respjson.Field
		MagNormRange          respjson.Field
		MagUnc                respjson.Field
		NetObjSig             respjson.Field
		NetObjSigUnc          respjson.Field
		ObPosition            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		Penumbra              respjson.Field
		PrimaryExtinction     respjson.Field
		PrimaryExtinctionUnc  respjson.Field
		Ra                    respjson.Field
		RaBias                respjson.Field
		RaMeasured            respjson.Field
		Range                 respjson.Field
		RangeBias             respjson.Field
		RangeMeasured         respjson.Field
		RangeRate             respjson.Field
		RangeRateMeasured     respjson.Field
		RangeRateUnc          respjson.Field
		RangeUnc              respjson.Field
		RaRate                respjson.Field
		RaUnc                 respjson.Field
		RawFileUri            respjson.Field
		ReferenceFrame        respjson.Field
		SatNo                 respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		SenQuat               respjson.Field
		SenReferenceFrame     respjson.Field
		Senvelx               respjson.Field
		Senvely               respjson.Field
		Senvelz               respjson.Field
		Senx                  respjson.Field
		Seny                  respjson.Field
		Senz                  respjson.Field
		ShutterDelay          respjson.Field
		SkyBkgrnd             respjson.Field
		SolarDecAngle         respjson.Field
		SolarEqPhaseAngle     respjson.Field
		SolarPhaseAngle       respjson.Field
		SourceDl              respjson.Field
		TaskID                respjson.Field
		TimingBias            respjson.Field
		TrackID               respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		Uct                   respjson.Field
		Umbra                 respjson.Field
		Zeroptd               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EoObservationAbridged) RawJSON() string { return r.JSON.raw }
func (r *EoObservationAbridged) UnmarshalJSON(data []byte) error {
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
type EoObservationAbridgedDataMode string

const (
	EoObservationAbridgedDataModeReal      EoObservationAbridgedDataMode = "REAL"
	EoObservationAbridgedDataModeTest      EoObservationAbridgedDataMode = "TEST"
	EoObservationAbridgedDataModeSimulated EoObservationAbridgedDataMode = "SIMULATED"
	EoObservationAbridgedDataModeExercise  EoObservationAbridgedDataMode = "EXERCISE"
)

// The reference frame of the EOObservation measurements. If the referenceFrame is
// null it is assumed to be J2000.
type EoObservationAbridgedReferenceFrame string

const (
	EoObservationAbridgedReferenceFrameJ2000 EoObservationAbridgedReferenceFrame = "J2000"
	EoObservationAbridgedReferenceFrameGcrf  EoObservationAbridgedReferenceFrame = "GCRF"
	EoObservationAbridgedReferenceFrameItrf  EoObservationAbridgedReferenceFrame = "ITRF"
	EoObservationAbridgedReferenceFrameTeme  EoObservationAbridgedReferenceFrame = "TEME"
)

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type EoObservationAbridgedSenReferenceFrame string

const (
	EoObservationAbridgedSenReferenceFrameJ2000   EoObservationAbridgedSenReferenceFrame = "J2000"
	EoObservationAbridgedSenReferenceFrameEfgTdr  EoObservationAbridgedSenReferenceFrame = "EFG/TDR"
	EoObservationAbridgedSenReferenceFrameEcrEcef EoObservationAbridgedSenReferenceFrame = "ECR/ECEF"
	EoObservationAbridgedSenReferenceFrameTeme    EoObservationAbridgedSenReferenceFrame = "TEME"
	EoObservationAbridgedSenReferenceFrameItrf    EoObservationAbridgedSenReferenceFrame = "ITRF"
	EoObservationAbridgedSenReferenceFrameGcrf    EoObservationAbridgedSenReferenceFrame = "GCRF"
)

type ObservationEoObservationQueryhelpResponse struct {
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
func (r ObservationEoObservationQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationEoObservationQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObservationEoObservationNewParams struct {
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
	DataMode ObservationEoObservationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Line of sight azimuth angle in degrees and topocentric frame. Reported value
	// should include all applicable corrections as specified on the source provider
	// data card. If uncertain, consumers should contact the provider for details on
	// the applied corrections.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Sensor line of sight azimuth angle bias in degrees.
	AzimuthBias param.Opt[float64] `json:"azimuthBias,omitzero"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured param.Opt[bool] `json:"azimuthMeasured,omitzero"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the line of sight azimuth angle, in degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// Background intensity for IR observations, in kw/sr/um.
	BgIntensity param.Opt[float64] `json:"bgIntensity,omitzero"`
	// Method indicating telescope movement during collection (AUTOTRACK, MANUAL
	// AUTOTRACK, MANUAL RATE TRACK, MANUAL SIDEREAL, SIDEREAL, RATE TRACK).
	CollectMethod param.Opt[string] `json:"collectMethod,omitzero"`
	// Object Correlation Quality score of the observation when compared to a known
	// orbit state, (non-standardized). Users should consult data providers regarding
	// the expected range of values.
	CorrQuality param.Opt[float64] `json:"corrQuality,omitzero"`
	// Line of sight declination, in degrees, in the specified referenceFrame. If
	// referenceFrame is null then J2K should be assumed. Reported value should include
	// all applicable corrections as specified on the source provider data card. If
	// uncertain, consumers should contact the provider for details on the applied
	// corrections.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// Sensor line of sight declination angle bias in degrees.
	DeclinationBias param.Opt[float64] `json:"declinationBias,omitzero"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured param.Opt[bool] `json:"declinationMeasured,omitzero"`
	// Line of sight declination rate of change, in degrees/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	DeclinationRate param.Opt[float64] `json:"declinationRate,omitzero"`
	// One sigma uncertainty in the line of sight declination angle, in degrees.
	DeclinationUnc param.Opt[float64] `json:"declinationUnc,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Line of sight elevation in degrees and topocentric frame. Reported value should
	// include all applicable corrections as specified on the source provider data
	// card. If uncertain, consumers should contact the provider for details on the
	// applied corrections.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Sensor line of sight elevation bias in degrees.
	ElevationBias param.Opt[float64] `json:"elevationBias,omitzero"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured param.Opt[bool] `json:"elevationMeasured,omitzero"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the line of sight elevation angle, in degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// Image exposure duration in seconds. For observations performed using frame
	// stacking or synthetic tracking methods, the exposure duration should be the
	// total integration time. This field is highly recommended / required if the
	// observations are going to be used for photometric processing.
	ExpDuration param.Opt[float64] `json:"expDuration,omitzero"`
	// The number of RSOs detected in the sensor field of view.
	FovCount param.Opt[int64] `json:"fovCount,omitzero"`
	// The number of uncorrelated tracks in the field of view.
	FovCountUct param.Opt[int64] `json:"fovCountUCT,omitzero"`
	// For GEO detections, the altitude in km.
	Geoalt param.Opt[float64] `json:"geoalt,omitzero"`
	// For GEO detections, the latitude in degrees north.
	Geolat param.Opt[float64] `json:"geolat,omitzero"`
	// For GEO detections, the longitude in degrees east.
	Geolon param.Opt[float64] `json:"geolon,omitzero"`
	// For GEO detections, the range in km.
	Georange param.Opt[float64] `json:"georange,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Unique identifier of the Sky Imagery.
	IDSkyImagery param.Opt[string] `json:"idSkyImagery,omitzero"`
	// Intensity of the target for IR observations, in kw/sr/um.
	Intensity param.Opt[float64] `json:"intensity,omitzero"`
	// One sigma uncertainty in the line of sight pointing in micro-radians.
	LosUnc param.Opt[float64] `json:"losUnc,omitzero"`
	// Line-of-sight cartesian X position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losx param.Opt[float64] `json:"losx,omitzero"`
	// Line-of-sight cartesian X velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losxvel param.Opt[float64] `json:"losxvel,omitzero"`
	// Line-of-sight cartesian Y position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losy param.Opt[float64] `json:"losy,omitzero"`
	// Line-of-sight cartesian Y velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losyvel param.Opt[float64] `json:"losyvel,omitzero"`
	// Line-of-sight cartesian Z position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losz param.Opt[float64] `json:"losz,omitzero"`
	// Line-of-sight cartesian Z velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Loszvel param.Opt[float64] `json:"loszvel,omitzero"`
	// Measure of observed brightness calibrated against the Gaia G-band in units of
	// magnitudes.
	Mag param.Opt[float64] `json:"mag,omitzero"`
	// Formula: mag - 5.0 \* log_10(geo_range / 1000000.0).
	MagNormRange param.Opt[float64] `json:"magNormRange,omitzero"`
	// Uncertainty of the observed brightness in units of magnitudes.
	MagUnc param.Opt[float64] `json:"magUnc,omitzero"`
	// Net object signature = counts / expDuration.
	NetObjSig param.Opt[float64] `json:"netObjSig,omitzero"`
	// Net object signature uncertainty = counts uncertainty / expDuration.
	NetObjSigUnc param.Opt[float64] `json:"netObjSigUnc,omitzero"`
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
	// Boolean indicating that the target object was in a penumbral eclipse at the time
	// of this observation.
	Penumbra param.Opt[bool] `json:"penumbra,omitzero"`
	// Primary Extinction Coefficient, in Magnitudes. Primary Extinction is the
	// coefficient applied to the airmass to determine how much the observed visual
	// magnitude has been attenuated by the atmosphere. Extinction, in general,
	// describes the absorption and scattering of electromagnetic radiation by dust and
	// gas between an emitting astronomical object and the observer. See the
	// EOObservationDetails API for specification of extinction coefficients for
	// multiple spectral filters.
	PrimaryExtinction param.Opt[float64] `json:"primaryExtinction,omitzero"`
	// Primary Extinction Coefficient Uncertainty, in Magnitudes.
	PrimaryExtinctionUnc param.Opt[float64] `json:"primaryExtinctionUnc,omitzero"`
	// Line of sight right ascension, in degrees, in the specified referenceFrame. If
	// referenceFrame is null then J2K should be assumed. Reported value should include
	// all applicable corrections as specified on the source provider data card. If
	// uncertain, consumers should contact the provider for details on the applied
	// corrections.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Sensor line of sight right ascension bias in degrees.
	RaBias param.Opt[float64] `json:"raBias,omitzero"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured param.Opt[bool] `json:"raMeasured,omitzero"`
	// Line of sight range in km. If referenceFrame is null then J2K should be assumed.
	// Reported value should include all applicable corrections as specified on the
	// source provider data card. If uncertain, consumers should contact the provider
	// for details on the applied corrections.
	Range param.Opt[float64] `json:"range,omitzero"`
	// Sensor line of sight range bias in km.
	RangeBias param.Opt[float64] `json:"rangeBias,omitzero"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured param.Opt[bool] `json:"rangeMeasured,omitzero"`
	// Range rate in km/s. If referenceFrame is null then J2K should be assumed.
	// Reported value should include all applicable corrections as specified on the
	// source provider data card. If uncertain, consumers should contact the provider
	// for details on the applied corrections.
	RangeRate param.Opt[float64] `json:"rangeRate,omitzero"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured param.Opt[bool] `json:"rangeRateMeasured,omitzero"`
	// One sigma uncertainty in the line of sight range rate, in kilometers/second.
	RangeRateUnc param.Opt[float64] `json:"rangeRateUnc,omitzero"`
	// One sigma uncertainty in the line of sight range, in kilometers.
	RangeUnc param.Opt[float64] `json:"rangeUnc,omitzero"`
	// Line of sight right ascension rate of change, in degrees/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	RaRate param.Opt[float64] `json:"raRate,omitzero"`
	// One sigma uncertainty in the line of sight right ascension angle, in degrees.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude at obTime (if mobile/onorbit) in km.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// Cartesian X velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvelx param.Opt[float64] `json:"senvelx,omitzero"`
	// Cartesian Y velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvely param.Opt[float64] `json:"senvely,omitzero"`
	// Cartesian Z velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvelz param.Opt[float64] `json:"senvelz,omitzero"`
	// Cartesian X position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Senx param.Opt[float64] `json:"senx,omitzero"`
	// Cartesian Y position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Seny param.Opt[float64] `json:"seny,omitzero"`
	// Cartesian Z position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Senz param.Opt[float64] `json:"senz,omitzero"`
	// Shutter delay in seconds.
	ShutterDelay param.Opt[float64] `json:"shutterDelay,omitzero"`
	// Average Sky Background signal, in Magnitudes. Sky Background refers to the
	// incoming light from an apparently empty part of the night sky.
	SkyBkgrnd param.Opt[float64] `json:"skyBkgrnd,omitzero"`
	// Angle from the sun to the equatorial plane.
	SolarDecAngle param.Opt[float64] `json:"solarDecAngle,omitzero"`
	// The angle, in degrees, between the projections of the target-to-observer vector
	// and the target-to-sun vector onto the equatorial plane. The angle is represented
	// as negative when closing (i.e. before the opposition) and positive when opening
	// (after the opposition).
	SolarEqPhaseAngle param.Opt[float64] `json:"solarEqPhaseAngle,omitzero"`
	// The angle, in degrees, between the target-to-observer vector and the
	// target-to-sun vector.
	SolarPhaseAngle param.Opt[float64] `json:"solarPhaseAngle,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Sensor timing bias in seconds.
	TimingBias param.Opt[float64] `json:"timingBias,omitzero"`
	// Optional identifier of the track to which this observation belongs.
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
	// Boolean indicating that the target object was in umbral eclipse at the time of
	// this observation.
	Umbra param.Opt[bool] `json:"umbra,omitzero"`
	// Formula: 2.5 \* log_10 (zero_mag_counts / expDuration).
	Zeroptd param.Opt[float64] `json:"zeroptd,omitzero"`
	// This is the uncertainty in the zero point for the filter used for this
	// observation/row in units of mag. For use with differential photometry.
	ZeroPtdUnc param.Opt[float64] `json:"zeroPtdUnc,omitzero"`
	// Model representation of additional detailed observation data for electro-optical
	// based sensor phenomenologies.
	EoobservationDetails ObservationEoObservationNewParamsEoobservationDetails `json:"eoobservationDetails,omitzero"`
	// The reference frame of the EOObservation measurements. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "GCRF", "ITRF", "TEME".
	ReferenceFrame ObservationEoObservationNewParamsReferenceFrame `json:"referenceFrame,omitzero"`
	// The quaternion describing the rotation of the sensor in relation to the
	// body-fixed frame used for this system into the local geodetic frame, at
	// observation time (obTime). The array element order convention is scalar
	// component first, followed by the three vector components (qc, q1, q2, q3).
	SenQuat []float64 `json:"senQuat,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame ObservationEoObservationNewParamsSenReferenceFrame `json:"senReferenceFrame,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ObservationEoObservationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObservationEoObservationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationEoObservationNewParams) UnmarshalJSON(data []byte) error {
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
type ObservationEoObservationNewParamsDataMode string

const (
	ObservationEoObservationNewParamsDataModeReal      ObservationEoObservationNewParamsDataMode = "REAL"
	ObservationEoObservationNewParamsDataModeTest      ObservationEoObservationNewParamsDataMode = "TEST"
	ObservationEoObservationNewParamsDataModeSimulated ObservationEoObservationNewParamsDataMode = "SIMULATED"
	ObservationEoObservationNewParamsDataModeExercise  ObservationEoObservationNewParamsDataMode = "EXERCISE"
)

// Model representation of additional detailed observation data for electro-optical
// based sensor phenomenologies.
type ObservationEoObservationNewParamsEoobservationDetails struct {
	// World Coordinate System (WCS) X pixel origin in astrometric fit.
	AcalCrPixX param.Opt[float64] `json:"acalCrPixX,omitzero"`
	// World Coordinate System (WCS) Y pixel origin in astrometric fit.
	AcalCrPixY param.Opt[float64] `json:"acalCrPixY,omitzero"`
	// World Coordinate System (WCS) equatorial coordinate X origin corresponding to
	// CRPIX in astrometric fit in degrees.
	AcalCrValX param.Opt[float64] `json:"acalCrValX,omitzero"`
	// World Coordinate System (WCS) equatorial coordinate Y origin corresponding to
	// CRPIX in astrometric fit in degrees.
	AcalCrValY param.Opt[float64] `json:"acalCrValY,omitzero"`
	// Number of stars used in astrometric fit.
	AcalNumStars param.Opt[int64] `json:"acalNumStars,omitzero"`
	// This is the background signal at or in the vicinity of the radiometric source
	// position. Specifically, this is the average background count level (DN/pixel)
	// divided by the exposure time in seconds of the background pixels used in the
	// photometric extraction. DN/pixel/sec.
	BackgroundSignal param.Opt[float64] `json:"backgroundSignal,omitzero"`
	// Estimated 1-sigma uncertainty in the background signal at or in the vicinity of
	// the radiometric source position. DN/pixel/sec.
	BackgroundSignalUnc param.Opt[float64] `json:"backgroundSignalUnc,omitzero"`
	// The number of pixels binned horizontally.
	BinningHoriz param.Opt[int64] `json:"binningHoriz,omitzero"`
	// The number of pixels binned vertically.
	BinningVert param.Opt[int64] `json:"binningVert,omitzero"`
	// The x centroid position on the CCD of the target object in pixels.
	CcdObjPosX param.Opt[float64] `json:"ccdObjPosX,omitzero"`
	// The y centroid position on the CCD of the target object in pixels.
	CcdObjPosY param.Opt[float64] `json:"ccdObjPosY,omitzero"`
	// This is the pixel width of the target. This is either a frame-by-frame
	// measurement or a constant point spread function or synthetic aperture used in
	// the extraction.
	CcdObjWidth param.Opt[float64] `json:"ccdObjWidth,omitzero"`
	// Operating temperature of CCD recorded during exposure or measured during
	// calibrations in K.
	CcdTemp param.Opt[float64] `json:"ccdTemp,omitzero"`
	// Observed centroid column number on the focal plane in pixels (0 is left edge,
	// 0.5 is center of pixels along left of image).
	CentroidColumn param.Opt[float64] `json:"centroidColumn,omitzero"`
	// Observed centroid row number on the focal plane in pixels (0 is top edge, 0.5 is
	// center of pixels along top of image).
	CentroidRow param.Opt[float64] `json:"centroidRow,omitzero"`
	// Classification marking of the data in IC/CAPCO Portion-marked format, will be
	// set to EOObservation classificationMarking if blank.
	ClassificationMarking param.Opt[string] `json:"classificationMarking,omitzero"`
	// Spatial variance of image distribution in horizontal direction measured in
	// pixels squared.
	ColumnVariance param.Opt[float64] `json:"columnVariance,omitzero"`
	// The reference number n, in neutralDensityFilters for the currently used neutral
	// density filter.
	CurrentNeutralDensityFilterNum param.Opt[int64] `json:"currentNeutralDensityFilterNum,omitzero"`
	// The reference number, x, where x ranges from 1 to n, where n is the number
	// specified in spectralFilters that corresponds to the spectral filter given in
	// the corresponding spectralFilterNames.
	CurrentSpectralFilterNum param.Opt[int64] `json:"currentSpectralFilterNum,omitzero"`
	// Covariance (Y^2) in measured declination (Y) in deg^2.
	DeclinationCov param.Opt[float64] `json:"declinationCov,omitzero"`
	// Angle off element set reported in degrees.
	Does param.Opt[float64] `json:"does,omitzero"`
	// Some sensors have gain settings. This value is the gain used during the
	// observation in units e-/ADU. If no gain is used, the value = 1.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// Unique identifier of the parent EOObservation.
	IDEoObservation param.Opt[string] `json:"idEOObservation,omitzero"`
	// Sensor instantaneous field of view (ratio of pixel pitch to focal length).
	Ifov param.Opt[float64] `json:"ifov,omitzero"`
	// Instrumental magnitude of a sensor before corrections are applied for atmosphere
	// or to transform to standard magnitude scale.
	MagInstrumental param.Opt[float64] `json:"magInstrumental,omitzero"`
	// Uncertainty in the instrumental magnitude.
	MagInstrumentalUnc param.Opt[float64] `json:"magInstrumentalUnc,omitzero"`
	// Number of catalog stars in the detector field of view (FOV) with the target
	// object. Can be 0 for narrow FOV sensors.
	NumCatalogStars param.Opt[int64] `json:"numCatalogStars,omitzero"`
	// Number of correlated stars in the FOV with the target object. Can be 0 for
	// narrow FOV sensors.
	NumCorrelatedStars param.Opt[int64] `json:"numCorrelatedStars,omitzero"`
	// Number of detected stars in the FOV with the target object. Helps identify
	// frames with clouds.
	NumDetectedStars param.Opt[int64] `json:"numDetectedStars,omitzero"`
	// The value is the number of neutral density filters used.
	NumNeutralDensityFilters param.Opt[int64] `json:"numNeutralDensityFilters,omitzero"`
	// The value is the number of spectral filters used.
	NumSpectralFilters param.Opt[int64] `json:"numSpectralFilters,omitzero"`
	// Distance from the target object to the sun during the observation in meters.
	ObjSunRange param.Opt[float64] `json:"objSunRange,omitzero"`
	// Ob detection time in ISO 8601 UTC with microsecond precision, will be set to
	// EOObservation obTime if blank.
	ObTime param.Opt[time.Time] `json:"obTime,omitzero" format:"date-time"`
	// Optical Cross Section computed in units m(2)/ster.
	OpticalCrossSection param.Opt[float64] `json:"opticalCrossSection,omitzero"`
	// Uncertainty in Optical Cross Section computed in units m(2)/ster.
	OpticalCrossSectionUnc param.Opt[float64] `json:"opticalCrossSectionUnc,omitzero"`
	// Number of stars used in photometric fit count.
	PcalNumStars param.Opt[int64] `json:"pcalNumStars,omitzero"`
	// Peak Aperture Raw Counts is the value of the peak pixel in the real or synthetic
	// aperture containing the target signal.
	PeakApertureCount param.Opt[float64] `json:"peakApertureCount,omitzero"`
	// Peak Background Raw Counts is the largest pixel value used in background signal.
	PeakBackgroundCount param.Opt[int64] `json:"peakBackgroundCount,omitzero"`
	// Solar phase angle bisector vector. The vector that bisects the solar phase
	// angle. The phase angle bisector is the angle that is << of the value in #48.
	// Then calculate the point on the RA/DEC (ECI J2000.0) sphere that a vector at
	// this angle would intersect.
	PhaseAngBisect param.Opt[float64] `json:"phaseAngBisect,omitzero"`
	// Pixel array size (height) in pixels.
	PixelArrayHeight param.Opt[int64] `json:"pixelArrayHeight,omitzero"`
	// Pixel array size (width) in pixels.
	PixelArrayWidth param.Opt[int64] `json:"pixelArrayWidth,omitzero"`
	// Maximum valid pixel value, this is defined as 2^(number of bits per pixel). For
	// example, a CCD with 8-bitpixels, would have a maximum valid pixel value of 2^8
	// = 256. This can represent the saturation value of the detector, but some sensors
	// will saturate at a value significantly lower than full well depth. This is the
	// analog-to-digital conversion (ADC) saturation value.
	PixelMax param.Opt[int64] `json:"pixelMax,omitzero"`
	// Minimum valid pixel value, this is typically 0.
	PixelMin param.Opt[int64] `json:"pixelMin,omitzero"`
	// Predicted Azimuth angle of the target object from a ground -based sensor (no
	// atmospheric refraction correction required) in degrees. AZ_EL implies apparent
	// topocentric place in true of date reference frame as seen from the observer with
	// aberration due to the observer velocity and light travel time applied.
	PredictedAzimuth param.Opt[float64] `json:"predictedAzimuth,omitzero"`
	// Predicted Declination of the Target object from the frame of reference of the
	// sensor (J2000, geocentric velocity aberration). SGP4 and VCMs produce geocentric
	// origin and velocity aberration and subtracting the sensor geocentric position of
	// the sensor places in its reference frame.
	PredictedDeclination param.Opt[float64] `json:"predictedDeclination,omitzero"`
	// Uncertainty of Predicted Declination of the Target object from the frame of
	// reference of the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs
	// produce geocentric origin and velocity aberration and subtracting the sensor
	// geocentric position of the sensor places in its reference frame.
	PredictedDeclinationUnc param.Opt[float64] `json:"predictedDeclinationUnc,omitzero"`
	// Predicted elevation angle of the target object from a ground -based sensor (no
	// atmospheric refraction correction required) in degrees. AZ_EL implies apparent
	// topocentric place in true of date reference frame as seen from the observer with
	// aberration due to the observer velocity and light travel time applied.
	PredictedElevation param.Opt[float64] `json:"predictedElevation,omitzero"`
	// Predicted Right Ascension of the Target object from the frame of reference of
	// the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs produce
	// geocentric origin and velocity aberration and subtracting the sensor geocentric
	// position of the sensor places in its reference frame.
	PredictedRa param.Opt[float64] `json:"predictedRa,omitzero"`
	// Uncertainty of predicted Right Ascension of the Target object from the frame of
	// reference of the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs
	// produce geocentric origin and velocity aberration and subtracting the sensor
	// geocentric position of the sensor places in its reference frame.
	PredictedRaUnc param.Opt[float64] `json:"predictedRaUnc,omitzero"`
	// Covariance (x^2) in measured Right Ascension (X) in deg^2.
	RaCov param.Opt[float64] `json:"raCov,omitzero"`
	// Covariance (XY) in measured ra/declination (XY) in deg^2.
	RaDeclinationCov param.Opt[float64] `json:"raDeclinationCov,omitzero"`
	// Spatial covariance of image distribution across horizontal and vertical
	// directions measured in pixels squared.
	RowColCov param.Opt[float64] `json:"rowColCov,omitzero"`
	// Spatial variance of image distribution in vertical direction measured in pixels
	// squared.
	RowVariance param.Opt[float64] `json:"rowVariance,omitzero"`
	// Estimated signal-to-noise ratio (SNR) for the total radiometric signal. Under
	// some algorithms, this can be a constant per target (not per observation). Note:
	// this SNR applies to the total signal of the radiometric source (i.e.,
	// Net_Obj_Sig with units DN/sec), not to be confused with the SNR of the signal in
	// the peak pixel (i.e., DN/pixel/sec).
	SnrEst param.Opt[float64] `json:"snrEst,omitzero"`
	// Fraction of the sun that is illuminating the target object. This indicates if
	// the target is in the Earth’s penumbra or umbra. (It is 0 when object is in umbra
	// and 1 when object is fully illuminated.).
	SolarDiskFrac param.Opt[float64] `json:"solarDiskFrac,omitzero"`
	// Source of the data, will be set to EOObservation source if blank.
	Source param.Opt[string] `json:"source,omitzero"`
	// Azimuth angle of the sun from a ground-based telescope (no atmospheric
	// refraction correction required) the observer with aberration due to the observer
	// velocity and light travel time applied in degrees.
	SunAzimuth param.Opt[float64] `json:"sunAzimuth,omitzero"`
	// Elevation angle of the sun from a ground-based telescope (no atmospheric
	// refraction correction required) in degrees.
	SunElevation param.Opt[float64] `json:"sunElevation,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosX param.Opt[float64] `json:"sunStatePosX,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosY param.Opt[float64] `json:"sunStatePosY,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosZ param.Opt[float64] `json:"sunStatePosZ,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelX param.Opt[float64] `json:"sunStateVelX,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelY param.Opt[float64] `json:"sunStateVelY,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelZ param.Opt[float64] `json:"sunStateVelZ,omitzero"`
	// Uncertainty in the times reported in UTC in seconds.
	TimesUnc param.Opt[float64] `json:"timesUnc,omitzero"`
	// Time off element set reported in seconds.
	Toes param.Opt[float64] `json:"toes,omitzero"`
	// Color coefficient for filter n for a space-based sensor where there is no
	// atmospheric extinction. Must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	ColorCoeffs []float64 `json:"colorCoeffs,omitzero"`
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
	// , will be set to EOObservation dataMode if blank.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,omitzero"`
	// An array of measurements that correspond to the distance from the streak center
	// measured from the optical image in pixels that show change over an interval of
	// time. The array length is dependent on the length of the streak. The
	// distFromStreakCenter, surfBrightness, and surfBrightnessUnc arrays will match in
	// size.
	DistFromStreakCenter []float64 `json:"distFromStreakCenter,omitzero"`
	// The extinction coefficient computed for the nth filter. Must be present for all
	// values n=1 to numSpectralFilters, in incrementing order of n, and for no other
	// values of n. Units = mag/airmass.
	ExtinctionCoeffs []float64 `json:"extinctionCoeffs,omitzero"`
	// This is the uncertainty in the extinction coefficient for the nth filter. Must
	// be present for all values n=1 to numSpectralFilters, in incrementing order of n,
	// and for no other values of n. -9999 for space-based sensors. Units =
	// mag/airmass.
	ExtinctionCoeffsUnc []float64 `json:"extinctionCoeffsUnc,omitzero"`
	// Must be present for all values n=1 to numNeutralDensityFilters, in incrementing
	// order of n, and for no other values of n.
	NeutralDensityFilterNames []string `json:"neutralDensityFilterNames,omitzero"`
	// The transmission of the nth neutral density filter. Must be present for all
	// values n=1 to numNeutralDensityFilters, in incrementing order of n, and for no
	// other values of n.
	NeutralDensityFilterTransmissions []float64 `json:"neutralDensityFilterTransmissions,omitzero"`
	// This is the uncertainty in the transmission for the nth filter. Must be present
	// for all values n=1 to numNeutralDensityFilters, in incrementing order of n, and
	// for no other values of n.
	NeutralDensityFilterTransmissionsUnc []float64 `json:"neutralDensityFilterTransmissionsUnc,omitzero"`
	// Array of the SpectralFilters keywords, must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	SpectralFilters []string `json:"spectralFilters,omitzero"`
	// This is the in-band solar magnitude at 1 A.U. Must be present for all values n=1
	// to numSpectralFilters, in incrementing order of n, and for no other values of n.
	// Units = mag.
	SpectralFilterSolarMag []float64 `json:"spectralFilterSolarMag,omitzero"`
	// This is the in-band average irradiance of a 0th mag source. Must be present for
	// all values n=1 to numSpectralFilters, in incrementing order of n, and for no
	// other values of n. Units = W/m2/nm.
	SpectralZmfl []float64 `json:"spectralZMFL,omitzero"`
	// An array of surface brightness measurements in magnitudes per square arcsecond
	// from the optical image that show change over an interval of time. The array
	// length is dependent on the length of the streak. The distFromStreakCenter,
	// surfBrightness, and surfBrightnessUnc arrays will match in size.
	SurfBrightness []float64 `json:"surfBrightness,omitzero"`
	// An array of surface brightness uncertainty measurements in magnitudes per square
	// arcsecond from the optical image that show change over an interval of time. The
	// array length is dependent on the length of the streak. The distFromStreakCenter,
	// surfBrightness, and surfBrightnessUnc arrays will match in size.
	SurfBrightnessUnc []float64 `json:"surfBrightnessUnc,omitzero"`
	// This is the value for the zero-point calculated for each filter denoted in
	// spectralFilters. It is the difference between the catalog mag and instrumental
	// mag for a set of standard stars. For use with All Sky photometry. Must be
	// present for all values n=1 to numSpectralFilters, in incrementing order of n,
	// and for no other values of n.
	ZeroPoints []float64 `json:"zeroPoints,omitzero"`
	// This is the uncertainty in the zero point for the filter denoted in
	// spectralFilters. For use with All Sky photometry. Must be present for all values
	// n=1 to numSpectralFilters, in incrementing order of n, and for no other values
	// of n.
	ZeroPointsUnc []float64 `json:"zeroPointsUnc,omitzero"`
	paramObj
}

func (r ObservationEoObservationNewParamsEoobservationDetails) MarshalJSON() (data []byte, err error) {
	type shadow ObservationEoObservationNewParamsEoobservationDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationEoObservationNewParamsEoobservationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationEoObservationNewParamsEoobservationDetails](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// The reference frame of the EOObservation measurements. If the referenceFrame is
// null it is assumed to be J2000.
type ObservationEoObservationNewParamsReferenceFrame string

const (
	ObservationEoObservationNewParamsReferenceFrameJ2000 ObservationEoObservationNewParamsReferenceFrame = "J2000"
	ObservationEoObservationNewParamsReferenceFrameGcrf  ObservationEoObservationNewParamsReferenceFrame = "GCRF"
	ObservationEoObservationNewParamsReferenceFrameItrf  ObservationEoObservationNewParamsReferenceFrame = "ITRF"
	ObservationEoObservationNewParamsReferenceFrameTeme  ObservationEoObservationNewParamsReferenceFrame = "TEME"
)

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type ObservationEoObservationNewParamsSenReferenceFrame string

const (
	ObservationEoObservationNewParamsSenReferenceFrameJ2000   ObservationEoObservationNewParamsSenReferenceFrame = "J2000"
	ObservationEoObservationNewParamsSenReferenceFrameEfgTdr  ObservationEoObservationNewParamsSenReferenceFrame = "EFG/TDR"
	ObservationEoObservationNewParamsSenReferenceFrameEcrEcef ObservationEoObservationNewParamsSenReferenceFrame = "ECR/ECEF"
	ObservationEoObservationNewParamsSenReferenceFrameTeme    ObservationEoObservationNewParamsSenReferenceFrame = "TEME"
	ObservationEoObservationNewParamsSenReferenceFrameItrf    ObservationEoObservationNewParamsSenReferenceFrame = "ITRF"
	ObservationEoObservationNewParamsSenReferenceFrameGcrf    ObservationEoObservationNewParamsSenReferenceFrame = "GCRF"
)

type ObservationEoObservationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationEoObservationGetParams]'s query parameters as
// `url.Values`.
func (r ObservationEoObservationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEoObservationListParams struct {
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationEoObservationListParams]'s query parameters as
// `url.Values`.
func (r ObservationEoObservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEoObservationCountParams struct {
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationEoObservationCountParams]'s query parameters as
// `url.Values`.
func (r ObservationEoObservationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEoObservationNewBulkParams struct {
	Body []ObservationEoObservationNewBulkParamsBody
	// Flag to convert observation reference frame into J2000.
	ConvertToJ2K param.Opt[bool] `query:"convertToJ2K,omitzero" json:"-"`
	paramObj
}

func (r ObservationEoObservationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationEoObservationNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [ObservationEoObservationNewBulkParams]'s query parameters
// as `url.Values`.
func (r ObservationEoObservationNewBulkParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Model representation of observation data for electro-optical based sensor
// phenomenologies. ECI J2K is the preferred reference frame for EOObservations,
// however, several user-specified reference frames are accommodated. Users should
// check the EOObservation record as well as the 'Discover' tab in the storefront
// to confirm the coordinate frames used by the data provider.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type ObservationEoObservationNewBulkParamsBody struct {
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
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Line of sight azimuth angle in degrees and topocentric frame. Reported value
	// should include all applicable corrections as specified on the source provider
	// data card. If uncertain, consumers should contact the provider for details on
	// the applied corrections.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Sensor line of sight azimuth angle bias in degrees.
	AzimuthBias param.Opt[float64] `json:"azimuthBias,omitzero"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured param.Opt[bool] `json:"azimuthMeasured,omitzero"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the line of sight azimuth angle, in degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// Background intensity for IR observations, in kw/sr/um.
	BgIntensity param.Opt[float64] `json:"bgIntensity,omitzero"`
	// Method indicating telescope movement during collection (AUTOTRACK, MANUAL
	// AUTOTRACK, MANUAL RATE TRACK, MANUAL SIDEREAL, SIDEREAL, RATE TRACK).
	CollectMethod param.Opt[string] `json:"collectMethod,omitzero"`
	// Object Correlation Quality score of the observation when compared to a known
	// orbit state, (non-standardized). Users should consult data providers regarding
	// the expected range of values.
	CorrQuality param.Opt[float64] `json:"corrQuality,omitzero"`
	// Line of sight declination, in degrees, in the specified referenceFrame. If
	// referenceFrame is null then J2K should be assumed. Reported value should include
	// all applicable corrections as specified on the source provider data card. If
	// uncertain, consumers should contact the provider for details on the applied
	// corrections.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// Sensor line of sight declination angle bias in degrees.
	DeclinationBias param.Opt[float64] `json:"declinationBias,omitzero"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured param.Opt[bool] `json:"declinationMeasured,omitzero"`
	// Line of sight declination rate of change, in degrees/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	DeclinationRate param.Opt[float64] `json:"declinationRate,omitzero"`
	// One sigma uncertainty in the line of sight declination angle, in degrees.
	DeclinationUnc param.Opt[float64] `json:"declinationUnc,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Line of sight elevation in degrees and topocentric frame. Reported value should
	// include all applicable corrections as specified on the source provider data
	// card. If uncertain, consumers should contact the provider for details on the
	// applied corrections.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Sensor line of sight elevation bias in degrees.
	ElevationBias param.Opt[float64] `json:"elevationBias,omitzero"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured param.Opt[bool] `json:"elevationMeasured,omitzero"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the line of sight elevation angle, in degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// Image exposure duration in seconds. For observations performed using frame
	// stacking or synthetic tracking methods, the exposure duration should be the
	// total integration time. This field is highly recommended / required if the
	// observations are going to be used for photometric processing.
	ExpDuration param.Opt[float64] `json:"expDuration,omitzero"`
	// The number of RSOs detected in the sensor field of view.
	FovCount param.Opt[int64] `json:"fovCount,omitzero"`
	// The number of uncorrelated tracks in the field of view.
	FovCountUct param.Opt[int64] `json:"fovCountUCT,omitzero"`
	// For GEO detections, the altitude in km.
	Geoalt param.Opt[float64] `json:"geoalt,omitzero"`
	// For GEO detections, the latitude in degrees north.
	Geolat param.Opt[float64] `json:"geolat,omitzero"`
	// For GEO detections, the longitude in degrees east.
	Geolon param.Opt[float64] `json:"geolon,omitzero"`
	// For GEO detections, the range in km.
	Georange param.Opt[float64] `json:"georange,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Unique identifier of the Sky Imagery.
	IDSkyImagery param.Opt[string] `json:"idSkyImagery,omitzero"`
	// Intensity of the target for IR observations, in kw/sr/um.
	Intensity param.Opt[float64] `json:"intensity,omitzero"`
	// One sigma uncertainty in the line of sight pointing in micro-radians.
	LosUnc param.Opt[float64] `json:"losUnc,omitzero"`
	// Line-of-sight cartesian X position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losx param.Opt[float64] `json:"losx,omitzero"`
	// Line-of-sight cartesian X velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losxvel param.Opt[float64] `json:"losxvel,omitzero"`
	// Line-of-sight cartesian Y position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losy param.Opt[float64] `json:"losy,omitzero"`
	// Line-of-sight cartesian Y velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losyvel param.Opt[float64] `json:"losyvel,omitzero"`
	// Line-of-sight cartesian Z position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losz param.Opt[float64] `json:"losz,omitzero"`
	// Line-of-sight cartesian Z velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Loszvel param.Opt[float64] `json:"loszvel,omitzero"`
	// Measure of observed brightness calibrated against the Gaia G-band in units of
	// magnitudes.
	Mag param.Opt[float64] `json:"mag,omitzero"`
	// Formula: mag - 5.0 \* log_10(geo_range / 1000000.0).
	MagNormRange param.Opt[float64] `json:"magNormRange,omitzero"`
	// Uncertainty of the observed brightness in units of magnitudes.
	MagUnc param.Opt[float64] `json:"magUnc,omitzero"`
	// Net object signature = counts / expDuration.
	NetObjSig param.Opt[float64] `json:"netObjSig,omitzero"`
	// Net object signature uncertainty = counts uncertainty / expDuration.
	NetObjSigUnc param.Opt[float64] `json:"netObjSigUnc,omitzero"`
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
	// Boolean indicating that the target object was in a penumbral eclipse at the time
	// of this observation.
	Penumbra param.Opt[bool] `json:"penumbra,omitzero"`
	// Primary Extinction Coefficient, in Magnitudes. Primary Extinction is the
	// coefficient applied to the airmass to determine how much the observed visual
	// magnitude has been attenuated by the atmosphere. Extinction, in general,
	// describes the absorption and scattering of electromagnetic radiation by dust and
	// gas between an emitting astronomical object and the observer. See the
	// EOObservationDetails API for specification of extinction coefficients for
	// multiple spectral filters.
	PrimaryExtinction param.Opt[float64] `json:"primaryExtinction,omitzero"`
	// Primary Extinction Coefficient Uncertainty, in Magnitudes.
	PrimaryExtinctionUnc param.Opt[float64] `json:"primaryExtinctionUnc,omitzero"`
	// Line of sight right ascension, in degrees, in the specified referenceFrame. If
	// referenceFrame is null then J2K should be assumed. Reported value should include
	// all applicable corrections as specified on the source provider data card. If
	// uncertain, consumers should contact the provider for details on the applied
	// corrections.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Sensor line of sight right ascension bias in degrees.
	RaBias param.Opt[float64] `json:"raBias,omitzero"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured param.Opt[bool] `json:"raMeasured,omitzero"`
	// Line of sight range in km. If referenceFrame is null then J2K should be assumed.
	// Reported value should include all applicable corrections as specified on the
	// source provider data card. If uncertain, consumers should contact the provider
	// for details on the applied corrections.
	Range param.Opt[float64] `json:"range,omitzero"`
	// Sensor line of sight range bias in km.
	RangeBias param.Opt[float64] `json:"rangeBias,omitzero"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured param.Opt[bool] `json:"rangeMeasured,omitzero"`
	// Range rate in km/s. If referenceFrame is null then J2K should be assumed.
	// Reported value should include all applicable corrections as specified on the
	// source provider data card. If uncertain, consumers should contact the provider
	// for details on the applied corrections.
	RangeRate param.Opt[float64] `json:"rangeRate,omitzero"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured param.Opt[bool] `json:"rangeRateMeasured,omitzero"`
	// One sigma uncertainty in the line of sight range rate, in kilometers/second.
	RangeRateUnc param.Opt[float64] `json:"rangeRateUnc,omitzero"`
	// One sigma uncertainty in the line of sight range, in kilometers.
	RangeUnc param.Opt[float64] `json:"rangeUnc,omitzero"`
	// Line of sight right ascension rate of change, in degrees/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	RaRate param.Opt[float64] `json:"raRate,omitzero"`
	// One sigma uncertainty in the line of sight right ascension angle, in degrees.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude at obTime (if mobile/onorbit) in km.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// Cartesian X velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvelx param.Opt[float64] `json:"senvelx,omitzero"`
	// Cartesian Y velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvely param.Opt[float64] `json:"senvely,omitzero"`
	// Cartesian Z velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvelz param.Opt[float64] `json:"senvelz,omitzero"`
	// Cartesian X position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Senx param.Opt[float64] `json:"senx,omitzero"`
	// Cartesian Y position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Seny param.Opt[float64] `json:"seny,omitzero"`
	// Cartesian Z position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Senz param.Opt[float64] `json:"senz,omitzero"`
	// Shutter delay in seconds.
	ShutterDelay param.Opt[float64] `json:"shutterDelay,omitzero"`
	// Average Sky Background signal, in Magnitudes. Sky Background refers to the
	// incoming light from an apparently empty part of the night sky.
	SkyBkgrnd param.Opt[float64] `json:"skyBkgrnd,omitzero"`
	// Angle from the sun to the equatorial plane.
	SolarDecAngle param.Opt[float64] `json:"solarDecAngle,omitzero"`
	// The angle, in degrees, between the projections of the target-to-observer vector
	// and the target-to-sun vector onto the equatorial plane. The angle is represented
	// as negative when closing (i.e. before the opposition) and positive when opening
	// (after the opposition).
	SolarEqPhaseAngle param.Opt[float64] `json:"solarEqPhaseAngle,omitzero"`
	// The angle, in degrees, between the target-to-observer vector and the
	// target-to-sun vector.
	SolarPhaseAngle param.Opt[float64] `json:"solarPhaseAngle,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Sensor timing bias in seconds.
	TimingBias param.Opt[float64] `json:"timingBias,omitzero"`
	// Optional identifier of the track to which this observation belongs.
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
	// Boolean indicating that the target object was in umbral eclipse at the time of
	// this observation.
	Umbra param.Opt[bool] `json:"umbra,omitzero"`
	// Formula: 2.5 \* log_10 (zero_mag_counts / expDuration).
	Zeroptd param.Opt[float64] `json:"zeroptd,omitzero"`
	// This is the uncertainty in the zero point for the filter used for this
	// observation/row in units of mag. For use with differential photometry.
	ZeroPtdUnc param.Opt[float64] `json:"zeroPtdUnc,omitzero"`
	// Model representation of additional detailed observation data for electro-optical
	// based sensor phenomenologies.
	EoobservationDetails ObservationEoObservationNewBulkParamsBodyEoobservationDetails `json:"eoobservationDetails,omitzero"`
	// The reference frame of the EOObservation measurements. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "GCRF", "ITRF", "TEME".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// The quaternion describing the rotation of the sensor in relation to the
	// body-fixed frame used for this system into the local geodetic frame, at
	// observation time (obTime). The array element order convention is scalar
	// component first, followed by the three vector components (qc, q1, q2, q3).
	SenQuat []float64 `json:"senQuat,omitzero"`
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

func (r ObservationEoObservationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationEoObservationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationEoObservationNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationEoObservationNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ObservationEoObservationNewBulkParamsBody](
		"referenceFrame", "J2000", "GCRF", "ITRF", "TEME",
	)
	apijson.RegisterFieldValidator[ObservationEoObservationNewBulkParamsBody](
		"senReferenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// Model representation of additional detailed observation data for electro-optical
// based sensor phenomenologies.
type ObservationEoObservationNewBulkParamsBodyEoobservationDetails struct {
	// World Coordinate System (WCS) X pixel origin in astrometric fit.
	AcalCrPixX param.Opt[float64] `json:"acalCrPixX,omitzero"`
	// World Coordinate System (WCS) Y pixel origin in astrometric fit.
	AcalCrPixY param.Opt[float64] `json:"acalCrPixY,omitzero"`
	// World Coordinate System (WCS) equatorial coordinate X origin corresponding to
	// CRPIX in astrometric fit in degrees.
	AcalCrValX param.Opt[float64] `json:"acalCrValX,omitzero"`
	// World Coordinate System (WCS) equatorial coordinate Y origin corresponding to
	// CRPIX in astrometric fit in degrees.
	AcalCrValY param.Opt[float64] `json:"acalCrValY,omitzero"`
	// Number of stars used in astrometric fit.
	AcalNumStars param.Opt[int64] `json:"acalNumStars,omitzero"`
	// This is the background signal at or in the vicinity of the radiometric source
	// position. Specifically, this is the average background count level (DN/pixel)
	// divided by the exposure time in seconds of the background pixels used in the
	// photometric extraction. DN/pixel/sec.
	BackgroundSignal param.Opt[float64] `json:"backgroundSignal,omitzero"`
	// Estimated 1-sigma uncertainty in the background signal at or in the vicinity of
	// the radiometric source position. DN/pixel/sec.
	BackgroundSignalUnc param.Opt[float64] `json:"backgroundSignalUnc,omitzero"`
	// The number of pixels binned horizontally.
	BinningHoriz param.Opt[int64] `json:"binningHoriz,omitzero"`
	// The number of pixels binned vertically.
	BinningVert param.Opt[int64] `json:"binningVert,omitzero"`
	// The x centroid position on the CCD of the target object in pixels.
	CcdObjPosX param.Opt[float64] `json:"ccdObjPosX,omitzero"`
	// The y centroid position on the CCD of the target object in pixels.
	CcdObjPosY param.Opt[float64] `json:"ccdObjPosY,omitzero"`
	// This is the pixel width of the target. This is either a frame-by-frame
	// measurement or a constant point spread function or synthetic aperture used in
	// the extraction.
	CcdObjWidth param.Opt[float64] `json:"ccdObjWidth,omitzero"`
	// Operating temperature of CCD recorded during exposure or measured during
	// calibrations in K.
	CcdTemp param.Opt[float64] `json:"ccdTemp,omitzero"`
	// Observed centroid column number on the focal plane in pixels (0 is left edge,
	// 0.5 is center of pixels along left of image).
	CentroidColumn param.Opt[float64] `json:"centroidColumn,omitzero"`
	// Observed centroid row number on the focal plane in pixels (0 is top edge, 0.5 is
	// center of pixels along top of image).
	CentroidRow param.Opt[float64] `json:"centroidRow,omitzero"`
	// Classification marking of the data in IC/CAPCO Portion-marked format, will be
	// set to EOObservation classificationMarking if blank.
	ClassificationMarking param.Opt[string] `json:"classificationMarking,omitzero"`
	// Spatial variance of image distribution in horizontal direction measured in
	// pixels squared.
	ColumnVariance param.Opt[float64] `json:"columnVariance,omitzero"`
	// The reference number n, in neutralDensityFilters for the currently used neutral
	// density filter.
	CurrentNeutralDensityFilterNum param.Opt[int64] `json:"currentNeutralDensityFilterNum,omitzero"`
	// The reference number, x, where x ranges from 1 to n, where n is the number
	// specified in spectralFilters that corresponds to the spectral filter given in
	// the corresponding spectralFilterNames.
	CurrentSpectralFilterNum param.Opt[int64] `json:"currentSpectralFilterNum,omitzero"`
	// Covariance (Y^2) in measured declination (Y) in deg^2.
	DeclinationCov param.Opt[float64] `json:"declinationCov,omitzero"`
	// Angle off element set reported in degrees.
	Does param.Opt[float64] `json:"does,omitzero"`
	// Some sensors have gain settings. This value is the gain used during the
	// observation in units e-/ADU. If no gain is used, the value = 1.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// Unique identifier of the parent EOObservation.
	IDEoObservation param.Opt[string] `json:"idEOObservation,omitzero"`
	// Sensor instantaneous field of view (ratio of pixel pitch to focal length).
	Ifov param.Opt[float64] `json:"ifov,omitzero"`
	// Instrumental magnitude of a sensor before corrections are applied for atmosphere
	// or to transform to standard magnitude scale.
	MagInstrumental param.Opt[float64] `json:"magInstrumental,omitzero"`
	// Uncertainty in the instrumental magnitude.
	MagInstrumentalUnc param.Opt[float64] `json:"magInstrumentalUnc,omitzero"`
	// Number of catalog stars in the detector field of view (FOV) with the target
	// object. Can be 0 for narrow FOV sensors.
	NumCatalogStars param.Opt[int64] `json:"numCatalogStars,omitzero"`
	// Number of correlated stars in the FOV with the target object. Can be 0 for
	// narrow FOV sensors.
	NumCorrelatedStars param.Opt[int64] `json:"numCorrelatedStars,omitzero"`
	// Number of detected stars in the FOV with the target object. Helps identify
	// frames with clouds.
	NumDetectedStars param.Opt[int64] `json:"numDetectedStars,omitzero"`
	// The value is the number of neutral density filters used.
	NumNeutralDensityFilters param.Opt[int64] `json:"numNeutralDensityFilters,omitzero"`
	// The value is the number of spectral filters used.
	NumSpectralFilters param.Opt[int64] `json:"numSpectralFilters,omitzero"`
	// Distance from the target object to the sun during the observation in meters.
	ObjSunRange param.Opt[float64] `json:"objSunRange,omitzero"`
	// Ob detection time in ISO 8601 UTC with microsecond precision, will be set to
	// EOObservation obTime if blank.
	ObTime param.Opt[time.Time] `json:"obTime,omitzero" format:"date-time"`
	// Optical Cross Section computed in units m(2)/ster.
	OpticalCrossSection param.Opt[float64] `json:"opticalCrossSection,omitzero"`
	// Uncertainty in Optical Cross Section computed in units m(2)/ster.
	OpticalCrossSectionUnc param.Opt[float64] `json:"opticalCrossSectionUnc,omitzero"`
	// Number of stars used in photometric fit count.
	PcalNumStars param.Opt[int64] `json:"pcalNumStars,omitzero"`
	// Peak Aperture Raw Counts is the value of the peak pixel in the real or synthetic
	// aperture containing the target signal.
	PeakApertureCount param.Opt[float64] `json:"peakApertureCount,omitzero"`
	// Peak Background Raw Counts is the largest pixel value used in background signal.
	PeakBackgroundCount param.Opt[int64] `json:"peakBackgroundCount,omitzero"`
	// Solar phase angle bisector vector. The vector that bisects the solar phase
	// angle. The phase angle bisector is the angle that is << of the value in #48.
	// Then calculate the point on the RA/DEC (ECI J2000.0) sphere that a vector at
	// this angle would intersect.
	PhaseAngBisect param.Opt[float64] `json:"phaseAngBisect,omitzero"`
	// Pixel array size (height) in pixels.
	PixelArrayHeight param.Opt[int64] `json:"pixelArrayHeight,omitzero"`
	// Pixel array size (width) in pixels.
	PixelArrayWidth param.Opt[int64] `json:"pixelArrayWidth,omitzero"`
	// Maximum valid pixel value, this is defined as 2^(number of bits per pixel). For
	// example, a CCD with 8-bitpixels, would have a maximum valid pixel value of 2^8
	// = 256. This can represent the saturation value of the detector, but some sensors
	// will saturate at a value significantly lower than full well depth. This is the
	// analog-to-digital conversion (ADC) saturation value.
	PixelMax param.Opt[int64] `json:"pixelMax,omitzero"`
	// Minimum valid pixel value, this is typically 0.
	PixelMin param.Opt[int64] `json:"pixelMin,omitzero"`
	// Predicted Azimuth angle of the target object from a ground -based sensor (no
	// atmospheric refraction correction required) in degrees. AZ_EL implies apparent
	// topocentric place in true of date reference frame as seen from the observer with
	// aberration due to the observer velocity and light travel time applied.
	PredictedAzimuth param.Opt[float64] `json:"predictedAzimuth,omitzero"`
	// Predicted Declination of the Target object from the frame of reference of the
	// sensor (J2000, geocentric velocity aberration). SGP4 and VCMs produce geocentric
	// origin and velocity aberration and subtracting the sensor geocentric position of
	// the sensor places in its reference frame.
	PredictedDeclination param.Opt[float64] `json:"predictedDeclination,omitzero"`
	// Uncertainty of Predicted Declination of the Target object from the frame of
	// reference of the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs
	// produce geocentric origin and velocity aberration and subtracting the sensor
	// geocentric position of the sensor places in its reference frame.
	PredictedDeclinationUnc param.Opt[float64] `json:"predictedDeclinationUnc,omitzero"`
	// Predicted elevation angle of the target object from a ground -based sensor (no
	// atmospheric refraction correction required) in degrees. AZ_EL implies apparent
	// topocentric place in true of date reference frame as seen from the observer with
	// aberration due to the observer velocity and light travel time applied.
	PredictedElevation param.Opt[float64] `json:"predictedElevation,omitzero"`
	// Predicted Right Ascension of the Target object from the frame of reference of
	// the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs produce
	// geocentric origin and velocity aberration and subtracting the sensor geocentric
	// position of the sensor places in its reference frame.
	PredictedRa param.Opt[float64] `json:"predictedRa,omitzero"`
	// Uncertainty of predicted Right Ascension of the Target object from the frame of
	// reference of the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs
	// produce geocentric origin and velocity aberration and subtracting the sensor
	// geocentric position of the sensor places in its reference frame.
	PredictedRaUnc param.Opt[float64] `json:"predictedRaUnc,omitzero"`
	// Covariance (x^2) in measured Right Ascension (X) in deg^2.
	RaCov param.Opt[float64] `json:"raCov,omitzero"`
	// Covariance (XY) in measured ra/declination (XY) in deg^2.
	RaDeclinationCov param.Opt[float64] `json:"raDeclinationCov,omitzero"`
	// Spatial covariance of image distribution across horizontal and vertical
	// directions measured in pixels squared.
	RowColCov param.Opt[float64] `json:"rowColCov,omitzero"`
	// Spatial variance of image distribution in vertical direction measured in pixels
	// squared.
	RowVariance param.Opt[float64] `json:"rowVariance,omitzero"`
	// Estimated signal-to-noise ratio (SNR) for the total radiometric signal. Under
	// some algorithms, this can be a constant per target (not per observation). Note:
	// this SNR applies to the total signal of the radiometric source (i.e.,
	// Net_Obj_Sig with units DN/sec), not to be confused with the SNR of the signal in
	// the peak pixel (i.e., DN/pixel/sec).
	SnrEst param.Opt[float64] `json:"snrEst,omitzero"`
	// Fraction of the sun that is illuminating the target object. This indicates if
	// the target is in the Earth’s penumbra or umbra. (It is 0 when object is in umbra
	// and 1 when object is fully illuminated.).
	SolarDiskFrac param.Opt[float64] `json:"solarDiskFrac,omitzero"`
	// Source of the data, will be set to EOObservation source if blank.
	Source param.Opt[string] `json:"source,omitzero"`
	// Azimuth angle of the sun from a ground-based telescope (no atmospheric
	// refraction correction required) the observer with aberration due to the observer
	// velocity and light travel time applied in degrees.
	SunAzimuth param.Opt[float64] `json:"sunAzimuth,omitzero"`
	// Elevation angle of the sun from a ground-based telescope (no atmospheric
	// refraction correction required) in degrees.
	SunElevation param.Opt[float64] `json:"sunElevation,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosX param.Opt[float64] `json:"sunStatePosX,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosY param.Opt[float64] `json:"sunStatePosY,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosZ param.Opt[float64] `json:"sunStatePosZ,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelX param.Opt[float64] `json:"sunStateVelX,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelY param.Opt[float64] `json:"sunStateVelY,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelZ param.Opt[float64] `json:"sunStateVelZ,omitzero"`
	// Uncertainty in the times reported in UTC in seconds.
	TimesUnc param.Opt[float64] `json:"timesUnc,omitzero"`
	// Time off element set reported in seconds.
	Toes param.Opt[float64] `json:"toes,omitzero"`
	// Color coefficient for filter n for a space-based sensor where there is no
	// atmospheric extinction. Must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	ColorCoeffs []float64 `json:"colorCoeffs,omitzero"`
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
	// , will be set to EOObservation dataMode if blank.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,omitzero"`
	// An array of measurements that correspond to the distance from the streak center
	// measured from the optical image in pixels that show change over an interval of
	// time. The array length is dependent on the length of the streak. The
	// distFromStreakCenter, surfBrightness, and surfBrightnessUnc arrays will match in
	// size.
	DistFromStreakCenter []float64 `json:"distFromStreakCenter,omitzero"`
	// The extinction coefficient computed for the nth filter. Must be present for all
	// values n=1 to numSpectralFilters, in incrementing order of n, and for no other
	// values of n. Units = mag/airmass.
	ExtinctionCoeffs []float64 `json:"extinctionCoeffs,omitzero"`
	// This is the uncertainty in the extinction coefficient for the nth filter. Must
	// be present for all values n=1 to numSpectralFilters, in incrementing order of n,
	// and for no other values of n. -9999 for space-based sensors. Units =
	// mag/airmass.
	ExtinctionCoeffsUnc []float64 `json:"extinctionCoeffsUnc,omitzero"`
	// Must be present for all values n=1 to numNeutralDensityFilters, in incrementing
	// order of n, and for no other values of n.
	NeutralDensityFilterNames []string `json:"neutralDensityFilterNames,omitzero"`
	// The transmission of the nth neutral density filter. Must be present for all
	// values n=1 to numNeutralDensityFilters, in incrementing order of n, and for no
	// other values of n.
	NeutralDensityFilterTransmissions []float64 `json:"neutralDensityFilterTransmissions,omitzero"`
	// This is the uncertainty in the transmission for the nth filter. Must be present
	// for all values n=1 to numNeutralDensityFilters, in incrementing order of n, and
	// for no other values of n.
	NeutralDensityFilterTransmissionsUnc []float64 `json:"neutralDensityFilterTransmissionsUnc,omitzero"`
	// Array of the SpectralFilters keywords, must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	SpectralFilters []string `json:"spectralFilters,omitzero"`
	// This is the in-band solar magnitude at 1 A.U. Must be present for all values n=1
	// to numSpectralFilters, in incrementing order of n, and for no other values of n.
	// Units = mag.
	SpectralFilterSolarMag []float64 `json:"spectralFilterSolarMag,omitzero"`
	// This is the in-band average irradiance of a 0th mag source. Must be present for
	// all values n=1 to numSpectralFilters, in incrementing order of n, and for no
	// other values of n. Units = W/m2/nm.
	SpectralZmfl []float64 `json:"spectralZMFL,omitzero"`
	// An array of surface brightness measurements in magnitudes per square arcsecond
	// from the optical image that show change over an interval of time. The array
	// length is dependent on the length of the streak. The distFromStreakCenter,
	// surfBrightness, and surfBrightnessUnc arrays will match in size.
	SurfBrightness []float64 `json:"surfBrightness,omitzero"`
	// An array of surface brightness uncertainty measurements in magnitudes per square
	// arcsecond from the optical image that show change over an interval of time. The
	// array length is dependent on the length of the streak. The distFromStreakCenter,
	// surfBrightness, and surfBrightnessUnc arrays will match in size.
	SurfBrightnessUnc []float64 `json:"surfBrightnessUnc,omitzero"`
	// This is the value for the zero-point calculated for each filter denoted in
	// spectralFilters. It is the difference between the catalog mag and instrumental
	// mag for a set of standard stars. For use with All Sky photometry. Must be
	// present for all values n=1 to numSpectralFilters, in incrementing order of n,
	// and for no other values of n.
	ZeroPoints []float64 `json:"zeroPoints,omitzero"`
	// This is the uncertainty in the zero point for the filter denoted in
	// spectralFilters. For use with All Sky photometry. Must be present for all values
	// n=1 to numSpectralFilters, in incrementing order of n, and for no other values
	// of n.
	ZeroPointsUnc []float64 `json:"zeroPointsUnc,omitzero"`
	paramObj
}

func (r ObservationEoObservationNewBulkParamsBodyEoobservationDetails) MarshalJSON() (data []byte, err error) {
	type shadow ObservationEoObservationNewBulkParamsBodyEoobservationDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationEoObservationNewBulkParamsBodyEoobservationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationEoObservationNewBulkParamsBodyEoobservationDetails](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type ObservationEoObservationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationEoObservationTupleParams]'s query parameters as
// `url.Values`.
func (r ObservationEoObservationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEoObservationUnvalidatedPublishParams struct {
	Body []ObservationEoObservationUnvalidatedPublishParamsBody
	paramObj
}

func (r ObservationEoObservationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationEoObservationUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Model representation of observation data for electro-optical based sensor
// phenomenologies. ECI J2K is the preferred reference frame for EOObservations,
// however, several user-specified reference frames are accommodated. Users should
// check the EOObservation record as well as the 'Discover' tab in the storefront
// to confirm the coordinate frames used by the data provider.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type ObservationEoObservationUnvalidatedPublishParamsBody struct {
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
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Line of sight azimuth angle in degrees and topocentric frame. Reported value
	// should include all applicable corrections as specified on the source provider
	// data card. If uncertain, consumers should contact the provider for details on
	// the applied corrections.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Sensor line of sight azimuth angle bias in degrees.
	AzimuthBias param.Opt[float64] `json:"azimuthBias,omitzero"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured param.Opt[bool] `json:"azimuthMeasured,omitzero"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the line of sight azimuth angle, in degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// Background intensity for IR observations, in kw/sr/um.
	BgIntensity param.Opt[float64] `json:"bgIntensity,omitzero"`
	// Method indicating telescope movement during collection (AUTOTRACK, MANUAL
	// AUTOTRACK, MANUAL RATE TRACK, MANUAL SIDEREAL, SIDEREAL, RATE TRACK).
	CollectMethod param.Opt[string] `json:"collectMethod,omitzero"`
	// Object Correlation Quality score of the observation when compared to a known
	// orbit state, (non-standardized). Users should consult data providers regarding
	// the expected range of values.
	CorrQuality param.Opt[float64] `json:"corrQuality,omitzero"`
	// Line of sight declination, in degrees, in the specified referenceFrame. If
	// referenceFrame is null then J2K should be assumed. Reported value should include
	// all applicable corrections as specified on the source provider data card. If
	// uncertain, consumers should contact the provider for details on the applied
	// corrections.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// Sensor line of sight declination angle bias in degrees.
	DeclinationBias param.Opt[float64] `json:"declinationBias,omitzero"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured param.Opt[bool] `json:"declinationMeasured,omitzero"`
	// Line of sight declination rate of change, in degrees/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	DeclinationRate param.Opt[float64] `json:"declinationRate,omitzero"`
	// One sigma uncertainty in the line of sight declination angle, in degrees.
	DeclinationUnc param.Opt[float64] `json:"declinationUnc,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Line of sight elevation in degrees and topocentric frame. Reported value should
	// include all applicable corrections as specified on the source provider data
	// card. If uncertain, consumers should contact the provider for details on the
	// applied corrections.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Sensor line of sight elevation bias in degrees.
	ElevationBias param.Opt[float64] `json:"elevationBias,omitzero"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured param.Opt[bool] `json:"elevationMeasured,omitzero"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the line of sight elevation angle, in degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// Image exposure duration in seconds. For observations performed using frame
	// stacking or synthetic tracking methods, the exposure duration should be the
	// total integration time. This field is highly recommended / required if the
	// observations are going to be used for photometric processing.
	ExpDuration param.Opt[float64] `json:"expDuration,omitzero"`
	// The number of RSOs detected in the sensor field of view.
	FovCount param.Opt[int64] `json:"fovCount,omitzero"`
	// The number of uncorrelated tracks in the field of view.
	FovCountUct param.Opt[int64] `json:"fovCountUCT,omitzero"`
	// For GEO detections, the altitude in km.
	Geoalt param.Opt[float64] `json:"geoalt,omitzero"`
	// For GEO detections, the latitude in degrees north.
	Geolat param.Opt[float64] `json:"geolat,omitzero"`
	// For GEO detections, the longitude in degrees east.
	Geolon param.Opt[float64] `json:"geolon,omitzero"`
	// For GEO detections, the range in km.
	Georange param.Opt[float64] `json:"georange,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Unique identifier of the Sky Imagery.
	IDSkyImagery param.Opt[string] `json:"idSkyImagery,omitzero"`
	// Intensity of the target for IR observations, in kw/sr/um.
	Intensity param.Opt[float64] `json:"intensity,omitzero"`
	// One sigma uncertainty in the line of sight pointing in micro-radians.
	LosUnc param.Opt[float64] `json:"losUnc,omitzero"`
	// Line-of-sight cartesian X position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losx param.Opt[float64] `json:"losx,omitzero"`
	// Line-of-sight cartesian X velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losxvel param.Opt[float64] `json:"losxvel,omitzero"`
	// Line-of-sight cartesian Y position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losy param.Opt[float64] `json:"losy,omitzero"`
	// Line-of-sight cartesian Y velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losyvel param.Opt[float64] `json:"losyvel,omitzero"`
	// Line-of-sight cartesian Z position of the target, in km, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Losz param.Opt[float64] `json:"losz,omitzero"`
	// Line-of-sight cartesian Z velocity of target, in km/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Loszvel param.Opt[float64] `json:"loszvel,omitzero"`
	// Measure of observed brightness calibrated against the Gaia G-band in units of
	// magnitudes.
	Mag param.Opt[float64] `json:"mag,omitzero"`
	// Formula: mag - 5.0 \* log_10(geo_range / 1000000.0).
	MagNormRange param.Opt[float64] `json:"magNormRange,omitzero"`
	// Uncertainty of the observed brightness in units of magnitudes.
	MagUnc param.Opt[float64] `json:"magUnc,omitzero"`
	// Net object signature = counts / expDuration.
	NetObjSig param.Opt[float64] `json:"netObjSig,omitzero"`
	// Net object signature uncertainty = counts uncertainty / expDuration.
	NetObjSigUnc param.Opt[float64] `json:"netObjSigUnc,omitzero"`
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
	// Boolean indicating that the target object was in a penumbral eclipse at the time
	// of this observation.
	Penumbra param.Opt[bool] `json:"penumbra,omitzero"`
	// Primary Extinction Coefficient, in Magnitudes. Primary Extinction is the
	// coefficient applied to the airmass to determine how much the observed visual
	// magnitude has been attenuated by the atmosphere. Extinction, in general,
	// describes the absorption and scattering of electromagnetic radiation by dust and
	// gas between an emitting astronomical object and the observer. See the
	// EOObservationDetails API for specification of extinction coefficients for
	// multiple spectral filters.
	PrimaryExtinction param.Opt[float64] `json:"primaryExtinction,omitzero"`
	// Primary Extinction Coefficient Uncertainty, in Magnitudes.
	PrimaryExtinctionUnc param.Opt[float64] `json:"primaryExtinctionUnc,omitzero"`
	// Line of sight right ascension, in degrees, in the specified referenceFrame. If
	// referenceFrame is null then J2K should be assumed. Reported value should include
	// all applicable corrections as specified on the source provider data card. If
	// uncertain, consumers should contact the provider for details on the applied
	// corrections.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Sensor line of sight right ascension bias in degrees.
	RaBias param.Opt[float64] `json:"raBias,omitzero"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured param.Opt[bool] `json:"raMeasured,omitzero"`
	// Line of sight range in km. If referenceFrame is null then J2K should be assumed.
	// Reported value should include all applicable corrections as specified on the
	// source provider data card. If uncertain, consumers should contact the provider
	// for details on the applied corrections.
	Range param.Opt[float64] `json:"range,omitzero"`
	// Sensor line of sight range bias in km.
	RangeBias param.Opt[float64] `json:"rangeBias,omitzero"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured param.Opt[bool] `json:"rangeMeasured,omitzero"`
	// Range rate in km/s. If referenceFrame is null then J2K should be assumed.
	// Reported value should include all applicable corrections as specified on the
	// source provider data card. If uncertain, consumers should contact the provider
	// for details on the applied corrections.
	RangeRate param.Opt[float64] `json:"rangeRate,omitzero"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured param.Opt[bool] `json:"rangeRateMeasured,omitzero"`
	// One sigma uncertainty in the line of sight range rate, in kilometers/second.
	RangeRateUnc param.Opt[float64] `json:"rangeRateUnc,omitzero"`
	// One sigma uncertainty in the line of sight range, in kilometers.
	RangeUnc param.Opt[float64] `json:"rangeUnc,omitzero"`
	// Line of sight right ascension rate of change, in degrees/sec, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	RaRate param.Opt[float64] `json:"raRate,omitzero"`
	// One sigma uncertainty in the line of sight right ascension angle, in degrees.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude at obTime (if mobile/onorbit) in km.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// Cartesian X velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvelx param.Opt[float64] `json:"senvelx,omitzero"`
	// Cartesian Y velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvely param.Opt[float64] `json:"senvely,omitzero"`
	// Cartesian Z velocity of the observing mobile/onorbit sensor at obTime, in
	// km/sec, in the specified senReferenceFrame. If senReferenceFrame is null then
	// J2K should be assumed.
	Senvelz param.Opt[float64] `json:"senvelz,omitzero"`
	// Cartesian X position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Senx param.Opt[float64] `json:"senx,omitzero"`
	// Cartesian Y position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Seny param.Opt[float64] `json:"seny,omitzero"`
	// Cartesian Z position of the observing mobile/onorbit sensor at obTime, in km, in
	// the specified senReferenceFrame. If senReferenceFrame is null then J2K should be
	// assumed.
	Senz param.Opt[float64] `json:"senz,omitzero"`
	// Shutter delay in seconds.
	ShutterDelay param.Opt[float64] `json:"shutterDelay,omitzero"`
	// Average Sky Background signal, in Magnitudes. Sky Background refers to the
	// incoming light from an apparently empty part of the night sky.
	SkyBkgrnd param.Opt[float64] `json:"skyBkgrnd,omitzero"`
	// Angle from the sun to the equatorial plane.
	SolarDecAngle param.Opt[float64] `json:"solarDecAngle,omitzero"`
	// The angle, in degrees, between the projections of the target-to-observer vector
	// and the target-to-sun vector onto the equatorial plane. The angle is represented
	// as negative when closing (i.e. before the opposition) and positive when opening
	// (after the opposition).
	SolarEqPhaseAngle param.Opt[float64] `json:"solarEqPhaseAngle,omitzero"`
	// The angle, in degrees, between the target-to-observer vector and the
	// target-to-sun vector.
	SolarPhaseAngle param.Opt[float64] `json:"solarPhaseAngle,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Sensor timing bias in seconds.
	TimingBias param.Opt[float64] `json:"timingBias,omitzero"`
	// Optional identifier of the track to which this observation belongs.
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
	// Boolean indicating that the target object was in umbral eclipse at the time of
	// this observation.
	Umbra param.Opt[bool] `json:"umbra,omitzero"`
	// Formula: 2.5 \* log_10 (zero_mag_counts / expDuration).
	Zeroptd param.Opt[float64] `json:"zeroptd,omitzero"`
	// This is the uncertainty in the zero point for the filter used for this
	// observation/row in units of mag. For use with differential photometry.
	ZeroPtdUnc param.Opt[float64] `json:"zeroPtdUnc,omitzero"`
	// Model representation of additional detailed observation data for electro-optical
	// based sensor phenomenologies.
	EoobservationDetails ObservationEoObservationUnvalidatedPublishParamsBodyEoobservationDetails `json:"eoobservationDetails,omitzero"`
	// The reference frame of the EOObservation measurements. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "GCRF", "ITRF", "TEME".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// The quaternion describing the rotation of the sensor in relation to the
	// body-fixed frame used for this system into the local geodetic frame, at
	// observation time (obTime). The array element order convention is scalar
	// component first, followed by the three vector components (qc, q1, q2, q3).
	SenQuat []float64 `json:"senQuat,omitzero"`
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

func (r ObservationEoObservationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationEoObservationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationEoObservationUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationEoObservationUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ObservationEoObservationUnvalidatedPublishParamsBody](
		"referenceFrame", "J2000", "GCRF", "ITRF", "TEME",
	)
	apijson.RegisterFieldValidator[ObservationEoObservationUnvalidatedPublishParamsBody](
		"senReferenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// Model representation of additional detailed observation data for electro-optical
// based sensor phenomenologies.
type ObservationEoObservationUnvalidatedPublishParamsBodyEoobservationDetails struct {
	// World Coordinate System (WCS) X pixel origin in astrometric fit.
	AcalCrPixX param.Opt[float64] `json:"acalCrPixX,omitzero"`
	// World Coordinate System (WCS) Y pixel origin in astrometric fit.
	AcalCrPixY param.Opt[float64] `json:"acalCrPixY,omitzero"`
	// World Coordinate System (WCS) equatorial coordinate X origin corresponding to
	// CRPIX in astrometric fit in degrees.
	AcalCrValX param.Opt[float64] `json:"acalCrValX,omitzero"`
	// World Coordinate System (WCS) equatorial coordinate Y origin corresponding to
	// CRPIX in astrometric fit in degrees.
	AcalCrValY param.Opt[float64] `json:"acalCrValY,omitzero"`
	// Number of stars used in astrometric fit.
	AcalNumStars param.Opt[int64] `json:"acalNumStars,omitzero"`
	// This is the background signal at or in the vicinity of the radiometric source
	// position. Specifically, this is the average background count level (DN/pixel)
	// divided by the exposure time in seconds of the background pixels used in the
	// photometric extraction. DN/pixel/sec.
	BackgroundSignal param.Opt[float64] `json:"backgroundSignal,omitzero"`
	// Estimated 1-sigma uncertainty in the background signal at or in the vicinity of
	// the radiometric source position. DN/pixel/sec.
	BackgroundSignalUnc param.Opt[float64] `json:"backgroundSignalUnc,omitzero"`
	// The number of pixels binned horizontally.
	BinningHoriz param.Opt[int64] `json:"binningHoriz,omitzero"`
	// The number of pixels binned vertically.
	BinningVert param.Opt[int64] `json:"binningVert,omitzero"`
	// The x centroid position on the CCD of the target object in pixels.
	CcdObjPosX param.Opt[float64] `json:"ccdObjPosX,omitzero"`
	// The y centroid position on the CCD of the target object in pixels.
	CcdObjPosY param.Opt[float64] `json:"ccdObjPosY,omitzero"`
	// This is the pixel width of the target. This is either a frame-by-frame
	// measurement or a constant point spread function or synthetic aperture used in
	// the extraction.
	CcdObjWidth param.Opt[float64] `json:"ccdObjWidth,omitzero"`
	// Operating temperature of CCD recorded during exposure or measured during
	// calibrations in K.
	CcdTemp param.Opt[float64] `json:"ccdTemp,omitzero"`
	// Observed centroid column number on the focal plane in pixels (0 is left edge,
	// 0.5 is center of pixels along left of image).
	CentroidColumn param.Opt[float64] `json:"centroidColumn,omitzero"`
	// Observed centroid row number on the focal plane in pixels (0 is top edge, 0.5 is
	// center of pixels along top of image).
	CentroidRow param.Opt[float64] `json:"centroidRow,omitzero"`
	// Classification marking of the data in IC/CAPCO Portion-marked format, will be
	// set to EOObservation classificationMarking if blank.
	ClassificationMarking param.Opt[string] `json:"classificationMarking,omitzero"`
	// Spatial variance of image distribution in horizontal direction measured in
	// pixels squared.
	ColumnVariance param.Opt[float64] `json:"columnVariance,omitzero"`
	// The reference number n, in neutralDensityFilters for the currently used neutral
	// density filter.
	CurrentNeutralDensityFilterNum param.Opt[int64] `json:"currentNeutralDensityFilterNum,omitzero"`
	// The reference number, x, where x ranges from 1 to n, where n is the number
	// specified in spectralFilters that corresponds to the spectral filter given in
	// the corresponding spectralFilterNames.
	CurrentSpectralFilterNum param.Opt[int64] `json:"currentSpectralFilterNum,omitzero"`
	// Covariance (Y^2) in measured declination (Y) in deg^2.
	DeclinationCov param.Opt[float64] `json:"declinationCov,omitzero"`
	// Angle off element set reported in degrees.
	Does param.Opt[float64] `json:"does,omitzero"`
	// Some sensors have gain settings. This value is the gain used during the
	// observation in units e-/ADU. If no gain is used, the value = 1.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// Unique identifier of the parent EOObservation.
	IDEoObservation param.Opt[string] `json:"idEOObservation,omitzero"`
	// Sensor instantaneous field of view (ratio of pixel pitch to focal length).
	Ifov param.Opt[float64] `json:"ifov,omitzero"`
	// Instrumental magnitude of a sensor before corrections are applied for atmosphere
	// or to transform to standard magnitude scale.
	MagInstrumental param.Opt[float64] `json:"magInstrumental,omitzero"`
	// Uncertainty in the instrumental magnitude.
	MagInstrumentalUnc param.Opt[float64] `json:"magInstrumentalUnc,omitzero"`
	// Number of catalog stars in the detector field of view (FOV) with the target
	// object. Can be 0 for narrow FOV sensors.
	NumCatalogStars param.Opt[int64] `json:"numCatalogStars,omitzero"`
	// Number of correlated stars in the FOV with the target object. Can be 0 for
	// narrow FOV sensors.
	NumCorrelatedStars param.Opt[int64] `json:"numCorrelatedStars,omitzero"`
	// Number of detected stars in the FOV with the target object. Helps identify
	// frames with clouds.
	NumDetectedStars param.Opt[int64] `json:"numDetectedStars,omitzero"`
	// The value is the number of neutral density filters used.
	NumNeutralDensityFilters param.Opt[int64] `json:"numNeutralDensityFilters,omitzero"`
	// The value is the number of spectral filters used.
	NumSpectralFilters param.Opt[int64] `json:"numSpectralFilters,omitzero"`
	// Distance from the target object to the sun during the observation in meters.
	ObjSunRange param.Opt[float64] `json:"objSunRange,omitzero"`
	// Ob detection time in ISO 8601 UTC with microsecond precision, will be set to
	// EOObservation obTime if blank.
	ObTime param.Opt[time.Time] `json:"obTime,omitzero" format:"date-time"`
	// Optical Cross Section computed in units m(2)/ster.
	OpticalCrossSection param.Opt[float64] `json:"opticalCrossSection,omitzero"`
	// Uncertainty in Optical Cross Section computed in units m(2)/ster.
	OpticalCrossSectionUnc param.Opt[float64] `json:"opticalCrossSectionUnc,omitzero"`
	// Number of stars used in photometric fit count.
	PcalNumStars param.Opt[int64] `json:"pcalNumStars,omitzero"`
	// Peak Aperture Raw Counts is the value of the peak pixel in the real or synthetic
	// aperture containing the target signal.
	PeakApertureCount param.Opt[float64] `json:"peakApertureCount,omitzero"`
	// Peak Background Raw Counts is the largest pixel value used in background signal.
	PeakBackgroundCount param.Opt[int64] `json:"peakBackgroundCount,omitzero"`
	// Solar phase angle bisector vector. The vector that bisects the solar phase
	// angle. The phase angle bisector is the angle that is << of the value in #48.
	// Then calculate the point on the RA/DEC (ECI J2000.0) sphere that a vector at
	// this angle would intersect.
	PhaseAngBisect param.Opt[float64] `json:"phaseAngBisect,omitzero"`
	// Pixel array size (height) in pixels.
	PixelArrayHeight param.Opt[int64] `json:"pixelArrayHeight,omitzero"`
	// Pixel array size (width) in pixels.
	PixelArrayWidth param.Opt[int64] `json:"pixelArrayWidth,omitzero"`
	// Maximum valid pixel value, this is defined as 2^(number of bits per pixel). For
	// example, a CCD with 8-bitpixels, would have a maximum valid pixel value of 2^8
	// = 256. This can represent the saturation value of the detector, but some sensors
	// will saturate at a value significantly lower than full well depth. This is the
	// analog-to-digital conversion (ADC) saturation value.
	PixelMax param.Opt[int64] `json:"pixelMax,omitzero"`
	// Minimum valid pixel value, this is typically 0.
	PixelMin param.Opt[int64] `json:"pixelMin,omitzero"`
	// Predicted Azimuth angle of the target object from a ground -based sensor (no
	// atmospheric refraction correction required) in degrees. AZ_EL implies apparent
	// topocentric place in true of date reference frame as seen from the observer with
	// aberration due to the observer velocity and light travel time applied.
	PredictedAzimuth param.Opt[float64] `json:"predictedAzimuth,omitzero"`
	// Predicted Declination of the Target object from the frame of reference of the
	// sensor (J2000, geocentric velocity aberration). SGP4 and VCMs produce geocentric
	// origin and velocity aberration and subtracting the sensor geocentric position of
	// the sensor places in its reference frame.
	PredictedDeclination param.Opt[float64] `json:"predictedDeclination,omitzero"`
	// Uncertainty of Predicted Declination of the Target object from the frame of
	// reference of the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs
	// produce geocentric origin and velocity aberration and subtracting the sensor
	// geocentric position of the sensor places in its reference frame.
	PredictedDeclinationUnc param.Opt[float64] `json:"predictedDeclinationUnc,omitzero"`
	// Predicted elevation angle of the target object from a ground -based sensor (no
	// atmospheric refraction correction required) in degrees. AZ_EL implies apparent
	// topocentric place in true of date reference frame as seen from the observer with
	// aberration due to the observer velocity and light travel time applied.
	PredictedElevation param.Opt[float64] `json:"predictedElevation,omitzero"`
	// Predicted Right Ascension of the Target object from the frame of reference of
	// the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs produce
	// geocentric origin and velocity aberration and subtracting the sensor geocentric
	// position of the sensor places in its reference frame.
	PredictedRa param.Opt[float64] `json:"predictedRa,omitzero"`
	// Uncertainty of predicted Right Ascension of the Target object from the frame of
	// reference of the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs
	// produce geocentric origin and velocity aberration and subtracting the sensor
	// geocentric position of the sensor places in its reference frame.
	PredictedRaUnc param.Opt[float64] `json:"predictedRaUnc,omitzero"`
	// Covariance (x^2) in measured Right Ascension (X) in deg^2.
	RaCov param.Opt[float64] `json:"raCov,omitzero"`
	// Covariance (XY) in measured ra/declination (XY) in deg^2.
	RaDeclinationCov param.Opt[float64] `json:"raDeclinationCov,omitzero"`
	// Spatial covariance of image distribution across horizontal and vertical
	// directions measured in pixels squared.
	RowColCov param.Opt[float64] `json:"rowColCov,omitzero"`
	// Spatial variance of image distribution in vertical direction measured in pixels
	// squared.
	RowVariance param.Opt[float64] `json:"rowVariance,omitzero"`
	// Estimated signal-to-noise ratio (SNR) for the total radiometric signal. Under
	// some algorithms, this can be a constant per target (not per observation). Note:
	// this SNR applies to the total signal of the radiometric source (i.e.,
	// Net_Obj_Sig with units DN/sec), not to be confused with the SNR of the signal in
	// the peak pixel (i.e., DN/pixel/sec).
	SnrEst param.Opt[float64] `json:"snrEst,omitzero"`
	// Fraction of the sun that is illuminating the target object. This indicates if
	// the target is in the Earth’s penumbra or umbra. (It is 0 when object is in umbra
	// and 1 when object is fully illuminated.).
	SolarDiskFrac param.Opt[float64] `json:"solarDiskFrac,omitzero"`
	// Source of the data, will be set to EOObservation source if blank.
	Source param.Opt[string] `json:"source,omitzero"`
	// Azimuth angle of the sun from a ground-based telescope (no atmospheric
	// refraction correction required) the observer with aberration due to the observer
	// velocity and light travel time applied in degrees.
	SunAzimuth param.Opt[float64] `json:"sunAzimuth,omitzero"`
	// Elevation angle of the sun from a ground-based telescope (no atmospheric
	// refraction correction required) in degrees.
	SunElevation param.Opt[float64] `json:"sunElevation,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosX param.Opt[float64] `json:"sunStatePosX,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosY param.Opt[float64] `json:"sunStatePosY,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosZ param.Opt[float64] `json:"sunStatePosZ,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelX param.Opt[float64] `json:"sunStateVelX,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelY param.Opt[float64] `json:"sunStateVelY,omitzero"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelZ param.Opt[float64] `json:"sunStateVelZ,omitzero"`
	// Uncertainty in the times reported in UTC in seconds.
	TimesUnc param.Opt[float64] `json:"timesUnc,omitzero"`
	// Time off element set reported in seconds.
	Toes param.Opt[float64] `json:"toes,omitzero"`
	// Color coefficient for filter n for a space-based sensor where there is no
	// atmospheric extinction. Must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	ColorCoeffs []float64 `json:"colorCoeffs,omitzero"`
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
	// , will be set to EOObservation dataMode if blank.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,omitzero"`
	// An array of measurements that correspond to the distance from the streak center
	// measured from the optical image in pixels that show change over an interval of
	// time. The array length is dependent on the length of the streak. The
	// distFromStreakCenter, surfBrightness, and surfBrightnessUnc arrays will match in
	// size.
	DistFromStreakCenter []float64 `json:"distFromStreakCenter,omitzero"`
	// The extinction coefficient computed for the nth filter. Must be present for all
	// values n=1 to numSpectralFilters, in incrementing order of n, and for no other
	// values of n. Units = mag/airmass.
	ExtinctionCoeffs []float64 `json:"extinctionCoeffs,omitzero"`
	// This is the uncertainty in the extinction coefficient for the nth filter. Must
	// be present for all values n=1 to numSpectralFilters, in incrementing order of n,
	// and for no other values of n. -9999 for space-based sensors. Units =
	// mag/airmass.
	ExtinctionCoeffsUnc []float64 `json:"extinctionCoeffsUnc,omitzero"`
	// Must be present for all values n=1 to numNeutralDensityFilters, in incrementing
	// order of n, and for no other values of n.
	NeutralDensityFilterNames []string `json:"neutralDensityFilterNames,omitzero"`
	// The transmission of the nth neutral density filter. Must be present for all
	// values n=1 to numNeutralDensityFilters, in incrementing order of n, and for no
	// other values of n.
	NeutralDensityFilterTransmissions []float64 `json:"neutralDensityFilterTransmissions,omitzero"`
	// This is the uncertainty in the transmission for the nth filter. Must be present
	// for all values n=1 to numNeutralDensityFilters, in incrementing order of n, and
	// for no other values of n.
	NeutralDensityFilterTransmissionsUnc []float64 `json:"neutralDensityFilterTransmissionsUnc,omitzero"`
	// Array of the SpectralFilters keywords, must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	SpectralFilters []string `json:"spectralFilters,omitzero"`
	// This is the in-band solar magnitude at 1 A.U. Must be present for all values n=1
	// to numSpectralFilters, in incrementing order of n, and for no other values of n.
	// Units = mag.
	SpectralFilterSolarMag []float64 `json:"spectralFilterSolarMag,omitzero"`
	// This is the in-band average irradiance of a 0th mag source. Must be present for
	// all values n=1 to numSpectralFilters, in incrementing order of n, and for no
	// other values of n. Units = W/m2/nm.
	SpectralZmfl []float64 `json:"spectralZMFL,omitzero"`
	// An array of surface brightness measurements in magnitudes per square arcsecond
	// from the optical image that show change over an interval of time. The array
	// length is dependent on the length of the streak. The distFromStreakCenter,
	// surfBrightness, and surfBrightnessUnc arrays will match in size.
	SurfBrightness []float64 `json:"surfBrightness,omitzero"`
	// An array of surface brightness uncertainty measurements in magnitudes per square
	// arcsecond from the optical image that show change over an interval of time. The
	// array length is dependent on the length of the streak. The distFromStreakCenter,
	// surfBrightness, and surfBrightnessUnc arrays will match in size.
	SurfBrightnessUnc []float64 `json:"surfBrightnessUnc,omitzero"`
	// This is the value for the zero-point calculated for each filter denoted in
	// spectralFilters. It is the difference between the catalog mag and instrumental
	// mag for a set of standard stars. For use with All Sky photometry. Must be
	// present for all values n=1 to numSpectralFilters, in incrementing order of n,
	// and for no other values of n.
	ZeroPoints []float64 `json:"zeroPoints,omitzero"`
	// This is the uncertainty in the zero point for the filter denoted in
	// spectralFilters. For use with All Sky photometry. Must be present for all values
	// n=1 to numSpectralFilters, in incrementing order of n, and for no other values
	// of n.
	ZeroPointsUnc []float64 `json:"zeroPointsUnc,omitzero"`
	paramObj
}

func (r ObservationEoObservationUnvalidatedPublishParamsBodyEoobservationDetails) MarshalJSON() (data []byte, err error) {
	type shadow ObservationEoObservationUnvalidatedPublishParamsBodyEoobservationDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationEoObservationUnvalidatedPublishParamsBodyEoobservationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationEoObservationUnvalidatedPublishParamsBodyEoobservationDetails](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
