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

// ObservationEoObservationHistoryService contains methods and other services that
// help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationEoObservationHistoryService] method instead.
type ObservationEoObservationHistoryService struct {
	Options []option.RequestOption
}

// NewObservationEoObservationHistoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservationEoObservationHistoryService(opts ...option.RequestOption) (r ObservationEoObservationHistoryService) {
	r = ObservationEoObservationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationEoObservationHistoryService) List(ctx context.Context, query ObservationEoObservationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EoObservationFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/eoobservation/history"
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
func (r *ObservationEoObservationHistoryService) ListAutoPaging(ctx context.Context, query ObservationEoObservationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EoObservationFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationEoObservationHistoryService) Aodr(ctx context.Context, query ObservationEoObservationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/eoobservation/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationEoObservationHistoryService) Count(ctx context.Context, query ObservationEoObservationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/eoobservation/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of observation data for electro-optical based sensor
// phenomenologies. ECI J2K is the preferred reference frame for EOObservations,
// however, several user-specified reference frames are accommodated. Users should
// check the EOObservation record as well as the 'Discover' tab in the storefront
// to confirm the coordinate frames used by the data provider.
type EoObservationFull struct {
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
	DataMode EoObservationFullDataMode `json:"dataMode,required"`
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
	// Model representation of additional detailed observation data for electro-optical
	// based sensor phenomenologies.
	EoobservationDetails EoObservationFullEoobservationDetails `json:"eoobservationDetails"`
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
	ReferenceFrame EoObservationFullReferenceFrame `json:"referenceFrame"`
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
	SenReferenceFrame EoObservationFullSenReferenceFrame `json:"senReferenceFrame"`
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
	// This is the uncertainty in the zero point for the filter used for this
	// observation/row in units of mag. For use with differential photometry.
	ZeroPtdUnc float64 `json:"zeroPtdUnc"`
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
		EoobservationDetails  respjson.Field
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
		OnOrbit               respjson.Field
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
		Tags                  respjson.Field
		TaskID                respjson.Field
		TimingBias            respjson.Field
		TrackID               respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		Uct                   respjson.Field
		Umbra                 respjson.Field
		Zeroptd               respjson.Field
		ZeroPtdUnc            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EoObservationFull) RawJSON() string { return r.JSON.raw }
func (r *EoObservationFull) UnmarshalJSON(data []byte) error {
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
type EoObservationFullDataMode string

const (
	EoObservationFullDataModeReal      EoObservationFullDataMode = "REAL"
	EoObservationFullDataModeTest      EoObservationFullDataMode = "TEST"
	EoObservationFullDataModeSimulated EoObservationFullDataMode = "SIMULATED"
	EoObservationFullDataModeExercise  EoObservationFullDataMode = "EXERCISE"
)

// Model representation of additional detailed observation data for electro-optical
// based sensor phenomenologies.
type EoObservationFullEoobservationDetails struct {
	// World Coordinate System (WCS) X pixel origin in astrometric fit.
	AcalCrPixX float64 `json:"acalCrPixX"`
	// World Coordinate System (WCS) Y pixel origin in astrometric fit.
	AcalCrPixY float64 `json:"acalCrPixY"`
	// World Coordinate System (WCS) equatorial coordinate X origin corresponding to
	// CRPIX in astrometric fit in degrees.
	AcalCrValX float64 `json:"acalCrValX"`
	// World Coordinate System (WCS) equatorial coordinate Y origin corresponding to
	// CRPIX in astrometric fit in degrees.
	AcalCrValY float64 `json:"acalCrValY"`
	// Number of stars used in astrometric fit.
	AcalNumStars int64 `json:"acalNumStars"`
	// This is the background signal at or in the vicinity of the radiometric source
	// position. Specifically, this is the average background count level (DN/pixel)
	// divided by the exposure time in seconds of the background pixels used in the
	// photometric extraction. DN/pixel/sec.
	BackgroundSignal float64 `json:"backgroundSignal"`
	// Estimated 1-sigma uncertainty in the background signal at or in the vicinity of
	// the radiometric source position. DN/pixel/sec.
	BackgroundSignalUnc float64 `json:"backgroundSignalUnc"`
	// The number of pixels binned horizontally.
	BinningHoriz int64 `json:"binningHoriz"`
	// The number of pixels binned vertically.
	BinningVert int64 `json:"binningVert"`
	// The x centroid position on the CCD of the target object in pixels.
	CcdObjPosX float64 `json:"ccdObjPosX"`
	// The y centroid position on the CCD of the target object in pixels.
	CcdObjPosY float64 `json:"ccdObjPosY"`
	// This is the pixel width of the target. This is either a frame-by-frame
	// measurement or a constant point spread function or synthetic aperture used in
	// the extraction.
	CcdObjWidth float64 `json:"ccdObjWidth"`
	// Operating temperature of CCD recorded during exposure or measured during
	// calibrations in K.
	CcdTemp float64 `json:"ccdTemp"`
	// Observed centroid column number on the focal plane in pixels (0 is left edge,
	// 0.5 is center of pixels along left of image).
	CentroidColumn float64 `json:"centroidColumn"`
	// Observed centroid row number on the focal plane in pixels (0 is top edge, 0.5 is
	// center of pixels along top of image).
	CentroidRow float64 `json:"centroidRow"`
	// Classification marking of the data in IC/CAPCO Portion-marked format, will be
	// set to EOObservation classificationMarking if blank.
	ClassificationMarking string `json:"classificationMarking"`
	// Color coefficient for filter n for a space-based sensor where there is no
	// atmospheric extinction. Must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	ColorCoeffs []float64 `json:"colorCoeffs"`
	// Spatial variance of image distribution in horizontal direction measured in
	// pixels squared.
	ColumnVariance float64 `json:"columnVariance"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The reference number n, in neutralDensityFilters for the currently used neutral
	// density filter.
	CurrentNeutralDensityFilterNum int64 `json:"currentNeutralDensityFilterNum"`
	// The reference number, x, where x ranges from 1 to n, where n is the number
	// specified in spectralFilters that corresponds to the spectral filter given in
	// the corresponding spectralFilterNames.
	CurrentSpectralFilterNum int64 `json:"currentSpectralFilterNum"`
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
	DataMode string `json:"dataMode"`
	// Covariance (Y^2) in measured declination (Y) in deg^2.
	DeclinationCov float64 `json:"declinationCov"`
	// An array of measurements that correspond to the distance from the streak center
	// measured from the optical image in pixels that show change over an interval of
	// time. The array length is dependent on the length of the streak. The
	// distFromStreakCenter, surfBrightness, and surfBrightnessUnc arrays will match in
	// size.
	DistFromStreakCenter []float64 `json:"distFromStreakCenter"`
	// Angle off element set reported in degrees.
	Does float64 `json:"does"`
	// The extinction coefficient computed for the nth filter. Must be present for all
	// values n=1 to numSpectralFilters, in incrementing order of n, and for no other
	// values of n. Units = mag/airmass.
	ExtinctionCoeffs []float64 `json:"extinctionCoeffs"`
	// This is the uncertainty in the extinction coefficient for the nth filter. Must
	// be present for all values n=1 to numSpectralFilters, in incrementing order of n,
	// and for no other values of n. -9999 for space-based sensors. Units =
	// mag/airmass.
	ExtinctionCoeffsUnc []float64 `json:"extinctionCoeffsUnc"`
	// Some sensors have gain settings. This value is the gain used during the
	// observation in units e-/ADU. If no gain is used, the value = 1.
	Gain float64 `json:"gain"`
	// Unique identifier of the parent EOObservation.
	IDEoObservation string `json:"idEOObservation"`
	// Sensor instantaneous field of view (ratio of pixel pitch to focal length).
	Ifov float64 `json:"ifov"`
	// Instrumental magnitude of a sensor before corrections are applied for atmosphere
	// or to transform to standard magnitude scale.
	MagInstrumental float64 `json:"magInstrumental"`
	// Uncertainty in the instrumental magnitude.
	MagInstrumentalUnc float64 `json:"magInstrumentalUnc"`
	// Must be present for all values n=1 to numNeutralDensityFilters, in incrementing
	// order of n, and for no other values of n.
	NeutralDensityFilterNames []string `json:"neutralDensityFilterNames"`
	// The transmission of the nth neutral density filter. Must be present for all
	// values n=1 to numNeutralDensityFilters, in incrementing order of n, and for no
	// other values of n.
	NeutralDensityFilterTransmissions []float64 `json:"neutralDensityFilterTransmissions"`
	// This is the uncertainty in the transmission for the nth filter. Must be present
	// for all values n=1 to numNeutralDensityFilters, in incrementing order of n, and
	// for no other values of n.
	NeutralDensityFilterTransmissionsUnc []float64 `json:"neutralDensityFilterTransmissionsUnc"`
	// Number of catalog stars in the detector field of view (FOV) with the target
	// object. Can be 0 for narrow FOV sensors.
	NumCatalogStars int64 `json:"numCatalogStars"`
	// Number of correlated stars in the FOV with the target object. Can be 0 for
	// narrow FOV sensors.
	NumCorrelatedStars int64 `json:"numCorrelatedStars"`
	// Number of detected stars in the FOV with the target object. Helps identify
	// frames with clouds.
	NumDetectedStars int64 `json:"numDetectedStars"`
	// The value is the number of neutral density filters used.
	NumNeutralDensityFilters int64 `json:"numNeutralDensityFilters"`
	// The value is the number of spectral filters used.
	NumSpectralFilters int64 `json:"numSpectralFilters"`
	// Distance from the target object to the sun during the observation in meters.
	ObjSunRange float64 `json:"objSunRange"`
	// Ob detection time in ISO 8601 UTC with microsecond precision, will be set to
	// EOObservation obTime if blank.
	ObTime time.Time `json:"obTime" format:"date-time"`
	// Optical Cross Section computed in units m(2)/ster.
	OpticalCrossSection float64 `json:"opticalCrossSection"`
	// Uncertainty in Optical Cross Section computed in units m(2)/ster.
	OpticalCrossSectionUnc float64 `json:"opticalCrossSectionUnc"`
	// Number of stars used in photometric fit count.
	PcalNumStars int64 `json:"pcalNumStars"`
	// Peak Aperture Raw Counts is the value of the peak pixel in the real or synthetic
	// aperture containing the target signal.
	PeakApertureCount float64 `json:"peakApertureCount"`
	// Peak Background Raw Counts is the largest pixel value used in background signal.
	PeakBackgroundCount int64 `json:"peakBackgroundCount"`
	// Solar phase angle bisector vector. The vector that bisects the solar phase
	// angle. The phase angle bisector is the angle that is << of the value in #48.
	// Then calculate the point on the RA/DEC (ECI J2000.0) sphere that a vector at
	// this angle would intersect.
	PhaseAngBisect float64 `json:"phaseAngBisect"`
	// Pixel array size (height) in pixels.
	PixelArrayHeight int64 `json:"pixelArrayHeight"`
	// Pixel array size (width) in pixels.
	PixelArrayWidth int64 `json:"pixelArrayWidth"`
	// Maximum valid pixel value, this is defined as 2^(number of bits per pixel). For
	// example, a CCD with 8-bitpixels, would have a maximum valid pixel value of 2^8
	// = 256. This can represent the saturation value of the detector, but some sensors
	// will saturate at a value significantly lower than full well depth. This is the
	// analog-to-digital conversion (ADC) saturation value.
	PixelMax int64 `json:"pixelMax"`
	// Minimum valid pixel value, this is typically 0.
	PixelMin int64 `json:"pixelMin"`
	// Predicted Azimuth angle of the target object from a ground -based sensor (no
	// atmospheric refraction correction required) in degrees. AZ_EL implies apparent
	// topocentric place in true of date reference frame as seen from the observer with
	// aberration due to the observer velocity and light travel time applied.
	PredictedAzimuth float64 `json:"predictedAzimuth"`
	// Predicted Declination of the Target object from the frame of reference of the
	// sensor (J2000, geocentric velocity aberration). SGP4 and VCMs produce geocentric
	// origin and velocity aberration and subtracting the sensor geocentric position of
	// the sensor places in its reference frame.
	PredictedDeclination float64 `json:"predictedDeclination"`
	// Uncertainty of Predicted Declination of the Target object from the frame of
	// reference of the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs
	// produce geocentric origin and velocity aberration and subtracting the sensor
	// geocentric position of the sensor places in its reference frame.
	PredictedDeclinationUnc float64 `json:"predictedDeclinationUnc"`
	// Predicted elevation angle of the target object from a ground -based sensor (no
	// atmospheric refraction correction required) in degrees. AZ_EL implies apparent
	// topocentric place in true of date reference frame as seen from the observer with
	// aberration due to the observer velocity and light travel time applied.
	PredictedElevation float64 `json:"predictedElevation"`
	// Predicted Right Ascension of the Target object from the frame of reference of
	// the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs produce
	// geocentric origin and velocity aberration and subtracting the sensor geocentric
	// position of the sensor places in its reference frame.
	PredictedRa float64 `json:"predictedRa"`
	// Uncertainty of predicted Right Ascension of the Target object from the frame of
	// reference of the sensor (J2000, geocentric velocity aberration). SGP4 and VCMs
	// produce geocentric origin and velocity aberration and subtracting the sensor
	// geocentric position of the sensor places in its reference frame.
	PredictedRaUnc float64 `json:"predictedRaUnc"`
	// Covariance (x^2) in measured Right Ascension (X) in deg^2.
	RaCov float64 `json:"raCov"`
	// Covariance (XY) in measured ra/declination (XY) in deg^2.
	RaDeclinationCov float64 `json:"raDeclinationCov"`
	// Spatial covariance of image distribution across horizontal and vertical
	// directions measured in pixels squared.
	RowColCov float64 `json:"rowColCov"`
	// Spatial variance of image distribution in vertical direction measured in pixels
	// squared.
	RowVariance float64 `json:"rowVariance"`
	// Estimated signal-to-noise ratio (SNR) for the total radiometric signal. Under
	// some algorithms, this can be a constant per target (not per observation). Note:
	// this SNR applies to the total signal of the radiometric source (i.e.,
	// Net_Obj_Sig with units DN/sec), not to be confused with the SNR of the signal in
	// the peak pixel (i.e., DN/pixel/sec).
	SnrEst float64 `json:"snrEst"`
	// Fraction of the sun that is illuminating the target object. This indicates if
	// the target is in the Earth’s penumbra or umbra. (It is 0 when object is in umbra
	// and 1 when object is fully illuminated.).
	SolarDiskFrac float64 `json:"solarDiskFrac"`
	// Source of the data, will be set to EOObservation source if blank.
	Source string `json:"source"`
	// Array of the SpectralFilters keywords, must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	SpectralFilters []string `json:"spectralFilters"`
	// This is the in-band solar magnitude at 1 A.U. Must be present for all values n=1
	// to numSpectralFilters, in incrementing order of n, and for no other values of n.
	// Units = mag.
	SpectralFilterSolarMag []float64 `json:"spectralFilterSolarMag"`
	// This is the in-band average irradiance of a 0th mag source. Must be present for
	// all values n=1 to numSpectralFilters, in incrementing order of n, and for no
	// other values of n. Units = W/m2/nm.
	SpectralZmfl []float64 `json:"spectralZMFL"`
	// Azimuth angle of the sun from a ground-based telescope (no atmospheric
	// refraction correction required) the observer with aberration due to the observer
	// velocity and light travel time applied in degrees.
	SunAzimuth float64 `json:"sunAzimuth"`
	// Elevation angle of the sun from a ground-based telescope (no atmospheric
	// refraction correction required) in degrees.
	SunElevation float64 `json:"sunElevation"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosX float64 `json:"sunStatePosX"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosY float64 `json:"sunStatePosY"`
	// Sun state vector in ECI J2000 coordinate frame in km.
	SunStatePosZ float64 `json:"sunStatePosZ"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelX float64 `json:"sunStateVelX"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelY float64 `json:"sunStateVelY"`
	// Sun state vector in ECI J2000 coordinate frame in km/sec.
	SunStateVelZ float64 `json:"sunStateVelZ"`
	// An array of surface brightness measurements in magnitudes per square arcsecond
	// from the optical image that show change over an interval of time. The array
	// length is dependent on the length of the streak. The distFromStreakCenter,
	// surfBrightness, and surfBrightnessUnc arrays will match in size.
	SurfBrightness []float64 `json:"surfBrightness"`
	// An array of surface brightness uncertainty measurements in magnitudes per square
	// arcsecond from the optical image that show change over an interval of time. The
	// array length is dependent on the length of the streak. The distFromStreakCenter,
	// surfBrightness, and surfBrightnessUnc arrays will match in size.
	SurfBrightnessUnc []float64 `json:"surfBrightnessUnc"`
	// Uncertainty in the times reported in UTC in seconds.
	TimesUnc float64 `json:"timesUnc"`
	// Time off element set reported in seconds.
	Toes float64 `json:"toes"`
	// This is the value for the zero-point calculated for each filter denoted in
	// spectralFilters. It is the difference between the catalog mag and instrumental
	// mag for a set of standard stars. For use with All Sky photometry. Must be
	// present for all values n=1 to numSpectralFilters, in incrementing order of n,
	// and for no other values of n.
	ZeroPoints []float64 `json:"zeroPoints"`
	// This is the uncertainty in the zero point for the filter denoted in
	// spectralFilters. For use with All Sky photometry. Must be present for all values
	// n=1 to numSpectralFilters, in incrementing order of n, and for no other values
	// of n.
	ZeroPointsUnc []float64 `json:"zeroPointsUnc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcalCrPixX                           respjson.Field
		AcalCrPixY                           respjson.Field
		AcalCrValX                           respjson.Field
		AcalCrValY                           respjson.Field
		AcalNumStars                         respjson.Field
		BackgroundSignal                     respjson.Field
		BackgroundSignalUnc                  respjson.Field
		BinningHoriz                         respjson.Field
		BinningVert                          respjson.Field
		CcdObjPosX                           respjson.Field
		CcdObjPosY                           respjson.Field
		CcdObjWidth                          respjson.Field
		CcdTemp                              respjson.Field
		CentroidColumn                       respjson.Field
		CentroidRow                          respjson.Field
		ClassificationMarking                respjson.Field
		ColorCoeffs                          respjson.Field
		ColumnVariance                       respjson.Field
		CreatedAt                            respjson.Field
		CreatedBy                            respjson.Field
		CurrentNeutralDensityFilterNum       respjson.Field
		CurrentSpectralFilterNum             respjson.Field
		DataMode                             respjson.Field
		DeclinationCov                       respjson.Field
		DistFromStreakCenter                 respjson.Field
		Does                                 respjson.Field
		ExtinctionCoeffs                     respjson.Field
		ExtinctionCoeffsUnc                  respjson.Field
		Gain                                 respjson.Field
		IDEoObservation                      respjson.Field
		Ifov                                 respjson.Field
		MagInstrumental                      respjson.Field
		MagInstrumentalUnc                   respjson.Field
		NeutralDensityFilterNames            respjson.Field
		NeutralDensityFilterTransmissions    respjson.Field
		NeutralDensityFilterTransmissionsUnc respjson.Field
		NumCatalogStars                      respjson.Field
		NumCorrelatedStars                   respjson.Field
		NumDetectedStars                     respjson.Field
		NumNeutralDensityFilters             respjson.Field
		NumSpectralFilters                   respjson.Field
		ObjSunRange                          respjson.Field
		ObTime                               respjson.Field
		OpticalCrossSection                  respjson.Field
		OpticalCrossSectionUnc               respjson.Field
		PcalNumStars                         respjson.Field
		PeakApertureCount                    respjson.Field
		PeakBackgroundCount                  respjson.Field
		PhaseAngBisect                       respjson.Field
		PixelArrayHeight                     respjson.Field
		PixelArrayWidth                      respjson.Field
		PixelMax                             respjson.Field
		PixelMin                             respjson.Field
		PredictedAzimuth                     respjson.Field
		PredictedDeclination                 respjson.Field
		PredictedDeclinationUnc              respjson.Field
		PredictedElevation                   respjson.Field
		PredictedRa                          respjson.Field
		PredictedRaUnc                       respjson.Field
		RaCov                                respjson.Field
		RaDeclinationCov                     respjson.Field
		RowColCov                            respjson.Field
		RowVariance                          respjson.Field
		SnrEst                               respjson.Field
		SolarDiskFrac                        respjson.Field
		Source                               respjson.Field
		SpectralFilters                      respjson.Field
		SpectralFilterSolarMag               respjson.Field
		SpectralZmfl                         respjson.Field
		SunAzimuth                           respjson.Field
		SunElevation                         respjson.Field
		SunStatePosX                         respjson.Field
		SunStatePosY                         respjson.Field
		SunStatePosZ                         respjson.Field
		SunStateVelX                         respjson.Field
		SunStateVelY                         respjson.Field
		SunStateVelZ                         respjson.Field
		SurfBrightness                       respjson.Field
		SurfBrightnessUnc                    respjson.Field
		TimesUnc                             respjson.Field
		Toes                                 respjson.Field
		ZeroPoints                           respjson.Field
		ZeroPointsUnc                        respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EoObservationFullEoobservationDetails) RawJSON() string { return r.JSON.raw }
func (r *EoObservationFullEoobservationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reference frame of the EOObservation measurements. If the referenceFrame is
// null it is assumed to be J2000.
type EoObservationFullReferenceFrame string

const (
	EoObservationFullReferenceFrameJ2000 EoObservationFullReferenceFrame = "J2000"
	EoObservationFullReferenceFrameGcrf  EoObservationFullReferenceFrame = "GCRF"
	EoObservationFullReferenceFrameItrf  EoObservationFullReferenceFrame = "ITRF"
	EoObservationFullReferenceFrameTeme  EoObservationFullReferenceFrame = "TEME"
)

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type EoObservationFullSenReferenceFrame string

const (
	EoObservationFullSenReferenceFrameJ2000   EoObservationFullSenReferenceFrame = "J2000"
	EoObservationFullSenReferenceFrameEfgTdr  EoObservationFullSenReferenceFrame = "EFG/TDR"
	EoObservationFullSenReferenceFrameEcrEcef EoObservationFullSenReferenceFrame = "ECR/ECEF"
	EoObservationFullSenReferenceFrameTeme    EoObservationFullSenReferenceFrame = "TEME"
	EoObservationFullSenReferenceFrameItrf    EoObservationFullSenReferenceFrame = "ITRF"
	EoObservationFullSenReferenceFrameGcrf    EoObservationFullSenReferenceFrame = "GCRF"
)

type ObservationEoObservationHistoryListParams struct {
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
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

// URLQuery serializes [ObservationEoObservationHistoryListParams]'s query
// parameters as `url.Values`.
func (r ObservationEoObservationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEoObservationHistoryAodrParams struct {
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
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

// URLQuery serializes [ObservationEoObservationHistoryAodrParams]'s query
// parameters as `url.Values`.
func (r ObservationEoObservationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEoObservationHistoryCountParams struct {
	// Ob detection time in ISO 8601 UTC, up to microsecond precision. Consumers should
	// contact the provider for details on their obTime specifications.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationEoObservationHistoryCountParams]'s query
// parameters as `url.Values`.
func (r ObservationEoObservationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
