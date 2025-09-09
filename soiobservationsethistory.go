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
)

// SoiObservationSetHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSoiObservationSetHistoryService] method instead.
type SoiObservationSetHistoryService struct {
	Options []option.RequestOption
}

// NewSoiObservationSetHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSoiObservationSetHistoryService(opts ...option.RequestOption) (r SoiObservationSetHistoryService) {
	r = SoiObservationSetHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SoiObservationSetHistoryService) List(ctx context.Context, query SoiObservationSetHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SoiObservationSetFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/soiobservationset/history"
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
func (r *SoiObservationSetHistoryService) ListAutoPaging(ctx context.Context, query SoiObservationSetHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SoiObservationSetFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SoiObservationSetHistoryService) Aodr(ctx context.Context, query SoiObservationSetHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/soiobservationset/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SoiObservationSetHistoryService) Count(ctx context.Context, query SoiObservationSetHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/soiobservationset/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// These services provide operations for posting space object idenfification
// observation sets.
type SoiObservationSetFull struct {
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
	DataMode SoiObservationSetFullDataMode `json:"dataMode,required"`
	// The number of observation records in the set.
	NumObs int64 `json:"numObs,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Observation type (OPTICAL, RADAR).
	//
	// Any of "OPTICAL", "RADAR".
	Type SoiObservationSetFullType `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The number of pixels binned horizontally.
	BinningHoriz int64 `json:"binningHoriz"`
	// The number of pixels binned vertically.
	BinningVert int64 `json:"binningVert"`
	// Boolean indicating if a brightness variance change event was detected, based on
	// historical collection data for the object.
	BrightnessVarianceChangeDetected bool `json:"brightnessVarianceChangeDetected"`
	// Array of SOI Calibrations associated with this SOIObservationSet.
	Calibrations []SoiObservationSetFullCalibration `json:"calibrations"`
	// Type of calibration used by the Sensor (e.g. ALL SKY, DIFFERENTIAL, DEFAULT,
	// NONE).
	CalibrationType string `json:"calibrationType"`
	// Overall qualitative confidence assessment of change detection results (e.g.
	// HIGH, MEDIUM, LOW).
	ChangeConf string `json:"changeConf"`
	// Boolean indicating if any change event was detected, based on historical
	// collection data for the object.
	ChangeDetected bool `json:"changeDetected"`
	// Qualitative Collection Density assessment, with respect to confidence of
	// detecting a change event (e.g. HIGH, MEDIUM, LOW).
	CollectionDensityConf string `json:"collectionDensityConf"`
	// Universally Unique collection ID. Mechanism to correlate Single Point Photometry
	// (SPP) JSON files to images.
	CollectionID string `json:"collectionId"`
	// Mode indicating telescope movement during collection (AUTOTRACK, MANUAL
	// AUTOTRACK, MANUAL RATE TRACK, MANUAL SIDEREAL, SIDEREAL, RATE TRACK).
	CollectionMode string `json:"collectionMode"`
	// Object Correlation Quality value. Measures how close the observed object's orbit
	// is to matching an object in the catalog. The scale of this field may vary
	// depending on provider. Users should consult the data provider to verify the
	// meaning of the value (e.g. A value of 0.0 indicates a high/strong correlation,
	// while a value closer to 1.0 indicates low/weak correlation).
	CorrQuality float64 `json:"corrQuality"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Observation set detection end time in ISO 8601 UTC with microsecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The gain used during the collection, in units of photoelectrons per
	// analog-to-digital unit (e-/ADU). If no gain is used, the value = 1.
	Gain float64 `json:"gain"`
	// ID of the UDL Elset of the Space Object under observation.
	IDElset string `json:"idElset"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// ID of the observing sensor.
	IDSensor string `json:"idSensor"`
	// Line of sight declination at observation set detection end time. Specified in
	// degrees, in the specified referenceFrame. If referenceFrame is null then J2K
	// should be assumed (e.g -30 to 130.0).
	LosDeclinationEnd float64 `json:"losDeclinationEnd"`
	// Line of sight declination at observation set detection start time. Specified in
	// degrees, in the specified referenceFrame. If referenceFrame is null then J2K
	// should be assumed (e.g -30 to 130.0).
	LosDeclinationStart float64 `json:"losDeclinationStart"`
	// SOI msgCreateDate time in ISO 8601 UTC time, with millisecond precision.
	MsgCreateDate time.Time `json:"msgCreateDate" format:"date-time"`
	// The value is the number of spectral filters used.
	NumSpectralFilters int64 `json:"numSpectralFilters"`
	// OpticalSOIObservations associated with this SOIObservationSet.
	OpticalSoiObservationList []SoiObservationSetFullOpticalSoiObservationList `json:"opticalSOIObservationList"`
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
	// Optional identifier provided by the record source to indicate the sensor
	// identifier to which this attitude set applies if this set is reporting a single
	// sensor orientation. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// A threshold for percent of pixels that make up object signal that are beyond the
	// saturation point for the sensor that are removed in the EOSSA file, in range of
	// 0 to 1.
	PercentSatThreshold float64 `json:"percentSatThreshold"`
	// Boolean indicating if a periodicity change event was detected, based on
	// historical collection data for the object.
	PeriodicityChangeDetected bool `json:"periodicityChangeDetected"`
	// Qualitative assessment of the periodicity detection results from the Attitude
	// and Shape Retrieval (ASR) Periodicity Assessment (PA) Tool (e.g. HIGH, MEDIUM,
	// LOW).
	PeriodicityDetectionConf string `json:"periodicityDetectionConf"`
	// Qualitative Periodicity Sampling assessment, with respect to confidence of
	// detecting a change event (e.g. HIGH, MEDIUM, LOW).
	PeriodicitySamplingConf string `json:"periodicitySamplingConf"`
	// Pixel array size (height) in pixels.
	PixelArrayHeight int64 `json:"pixelArrayHeight"`
	// Pixel array size (width) in pixels.
	PixelArrayWidth int64 `json:"pixelArrayWidth"`
	// The maximum valid pixel value.
	PixelMax int64 `json:"pixelMax"`
	// The minimum valid pixel value.
	PixelMin int64 `json:"pixelMin"`
	// Pointing angle of the Azimuth gimbal/mount at observation set detection end
	// time. Specified in degrees.
	PointingAngleAzEnd float64 `json:"pointingAngleAzEnd"`
	// Pointing angle of the Azimuth gimbal/mount at observation set detection start
	// time. Specified in degrees.
	PointingAngleAzStart float64 `json:"pointingAngleAzStart"`
	// Pointing angle of the Elevation gimbal/mount at observation set detection end
	// time. Specified in degrees.
	PointingAngleElEnd float64 `json:"pointingAngleElEnd"`
	// Pointing angle of the Elevation gimbal/mount at observation set detection start
	// time. Specified in degrees.
	PointingAngleElStart float64 `json:"pointingAngleElStart"`
	// Polar angle of the gimbal/mount at observation set detection end time in
	// degrees.
	PolarAngleEnd float64 `json:"polarAngleEnd"`
	// Polar angle of the gimbal/mount at observation set detection start time in
	// degrees.
	PolarAngleStart float64 `json:"polarAngleStart"`
	// RadarSOIObservations associated with this RadarSOIObservationSet.
	RadarSoiObservationList []SoiObservationSetFullRadarSoiObservationList `json:"radarSOIObservationList"`
	// The reference frame of the observation measurements. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame SoiObservationSetFullReferenceFrame `json:"referenceFrame"`
	// Name of the target satellite.
	SatelliteName string `json:"satelliteName"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sensor altitude at startTime (if mobile/onorbit) in kilometers.
	Senalt float64 `json:"senalt"`
	// Sensor WGS84 latitude at startTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info. -90 to 90 degrees (negative values south of
	// equator).
	Senlat float64 `json:"senlat"`
	// Sensor WGS84 longitude at startTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values south of
	// equator).
	Senlon float64 `json:"senlon"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame SoiObservationSetFullSenReferenceFrame `json:"senReferenceFrame"`
	// ID of the AttitudeSet record for the observing sensor.
	SensorAsID string `json:"sensorAsId"`
	// Cartesian X velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvelx float64 `json:"senvelx"`
	// Cartesian Y velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvely float64 `json:"senvely"`
	// Cartesian Z velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvelz float64 `json:"senvelz"`
	// Cartesian X position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Senx float64 `json:"senx"`
	// Cartesian Y position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Seny float64 `json:"seny"`
	// Cartesian Z position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Senz float64 `json:"senz"`
	// Software Version used to Capture, Process, and Deliver the data.
	SoftwareVersion string `json:"softwareVersion"`
	// The in-band solar magnitude at 1 A.U.
	SolarMag float64 `json:"solarMag"`
	// Boolean indicating if a solar phase angle brightness change event was detected,
	// based on historical collection data for the object.
	SolarPhaseAngleBrightnessChangeDetected bool `json:"solarPhaseAngleBrightnessChangeDetected"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Array of the SpectralFilters keywords, must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	SpectralFilters []string `json:"spectralFilters"`
	// Name of the Star Catalog used for photometry and astrometry.
	StarCatName string `json:"starCatName"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating whether the target object was unable to be correlated to a
	// known object. This flag should only be set to true by data providers after an
	// attempt to correlate to an OnOrbit object was made and failed. If unable to
	// correlate, the 'origObjectId' field may be populated with an internal data
	// provider specific identifier.
	Uct bool `json:"uct"`
	// Key to indicate which, if any of, the pre/post photometer calibrations are valid
	// for use when generating data for the EOSSA file. If the field is not populated,
	// then the provided calibration data will be used when generating the EOSSA file
	// (e.g. PRE, POST, BOTH, NONE).
	ValidCalibrations string `json:"validCalibrations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                   respjson.Field
		DataMode                                respjson.Field
		NumObs                                  respjson.Field
		Source                                  respjson.Field
		StartTime                               respjson.Field
		Type                                    respjson.Field
		ID                                      respjson.Field
		BinningHoriz                            respjson.Field
		BinningVert                             respjson.Field
		BrightnessVarianceChangeDetected        respjson.Field
		Calibrations                            respjson.Field
		CalibrationType                         respjson.Field
		ChangeConf                              respjson.Field
		ChangeDetected                          respjson.Field
		CollectionDensityConf                   respjson.Field
		CollectionID                            respjson.Field
		CollectionMode                          respjson.Field
		CorrQuality                             respjson.Field
		CreatedAt                               respjson.Field
		CreatedBy                               respjson.Field
		EndTime                                 respjson.Field
		Gain                                    respjson.Field
		IDElset                                 respjson.Field
		IDOnOrbit                               respjson.Field
		IDSensor                                respjson.Field
		LosDeclinationEnd                       respjson.Field
		LosDeclinationStart                     respjson.Field
		MsgCreateDate                           respjson.Field
		NumSpectralFilters                      respjson.Field
		OpticalSoiObservationList               respjson.Field
		Origin                                  respjson.Field
		OrigNetwork                             respjson.Field
		OrigObjectID                            respjson.Field
		OrigSensorID                            respjson.Field
		PercentSatThreshold                     respjson.Field
		PeriodicityChangeDetected               respjson.Field
		PeriodicityDetectionConf                respjson.Field
		PeriodicitySamplingConf                 respjson.Field
		PixelArrayHeight                        respjson.Field
		PixelArrayWidth                         respjson.Field
		PixelMax                                respjson.Field
		PixelMin                                respjson.Field
		PointingAngleAzEnd                      respjson.Field
		PointingAngleAzStart                    respjson.Field
		PointingAngleElEnd                      respjson.Field
		PointingAngleElStart                    respjson.Field
		PolarAngleEnd                           respjson.Field
		PolarAngleStart                         respjson.Field
		RadarSoiObservationList                 respjson.Field
		ReferenceFrame                          respjson.Field
		SatelliteName                           respjson.Field
		SatNo                                   respjson.Field
		Senalt                                  respjson.Field
		Senlat                                  respjson.Field
		Senlon                                  respjson.Field
		SenReferenceFrame                       respjson.Field
		SensorAsID                              respjson.Field
		Senvelx                                 respjson.Field
		Senvely                                 respjson.Field
		Senvelz                                 respjson.Field
		Senx                                    respjson.Field
		Seny                                    respjson.Field
		Senz                                    respjson.Field
		SoftwareVersion                         respjson.Field
		SolarMag                                respjson.Field
		SolarPhaseAngleBrightnessChangeDetected respjson.Field
		SourceDl                                respjson.Field
		SpectralFilters                         respjson.Field
		StarCatName                             respjson.Field
		Tags                                    respjson.Field
		TransactionID                           respjson.Field
		Uct                                     respjson.Field
		ValidCalibrations                       respjson.Field
		ExtraFields                             map[string]respjson.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SoiObservationSetFull) RawJSON() string { return r.JSON.raw }
func (r *SoiObservationSetFull) UnmarshalJSON(data []byte) error {
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
type SoiObservationSetFullDataMode string

const (
	SoiObservationSetFullDataModeReal      SoiObservationSetFullDataMode = "REAL"
	SoiObservationSetFullDataModeTest      SoiObservationSetFullDataMode = "TEST"
	SoiObservationSetFullDataModeSimulated SoiObservationSetFullDataMode = "SIMULATED"
	SoiObservationSetFullDataModeExercise  SoiObservationSetFullDataMode = "EXERCISE"
)

// Observation type (OPTICAL, RADAR).
type SoiObservationSetFullType string

const (
	SoiObservationSetFullTypeOptical SoiObservationSetFullType = "OPTICAL"
	SoiObservationSetFullTypeRadar   SoiObservationSetFullType = "RADAR"
)

// Schema for SOI Calibration data.
type SoiObservationSetFullCalibration struct {
	// Background intensity, at calibration, specified in kilowatts per steradian per
	// micrometer.
	CalBgIntensity float64 `json:"calBgIntensity"`
	// Coefficient value for how much signal would be lost to atmospheric attenuation
	// for a star at zenith, in magnitudes per air mass.
	CalExtinctionCoeff float64 `json:"calExtinctionCoeff"`
	// Maximum extinction coefficient uncertainty in magnitudes, at calibration (e.g.
	// -5.0 to 30.0).
	CalExtinctionCoeffMaxUnc float64 `json:"calExtinctionCoeffMaxUnc"`
	// Extinction coefficient uncertainty in magnitudes, at calibration, which
	// represents the difference between the measured brightness and predicted
	// brightness of the star with the extinction removed, making it exo-atmospheric
	// (e.g. -5.0 to 30.0).
	CalExtinctionCoeffUnc float64 `json:"calExtinctionCoeffUnc"`
	// Number of correlated stars in the FOV with the target object, at calibration.
	// Can be 0 for narrow FOV sensors.
	CalNumCorrelatedStars int64 `json:"calNumCorrelatedStars"`
	// Number of detected stars in the FOV with the target object, at calibration.
	// Helps identify frames with clouds.
	CalNumDetectedStars int64 `json:"calNumDetectedStars"`
	// Average Sky Background signals in magnitudes, at calibration. Sky Background
	// refers to the incoming light from an apparently empty part of the night sky.
	CalSkyBg float64 `json:"calSkyBg"`
	// In-band solar magnitudes at 1 A.U, at calibration (e.g. -5.0 to 30.0).
	CalSpectralFilterSolarMag float64 `json:"calSpectralFilterSolarMag"`
	// Start time of calibration in ISO 8601 UTC time, with millisecond precision.
	CalTime time.Time `json:"calTime" format:"date-time"`
	// Type of calibration (e.g. PRE, MID, POST).
	CalType string `json:"calType"`
	// Value representing the difference between the catalog magnitude and instrumental
	// magnitude for a set of standard stars, at calibration (e.g. -5.0 to 30.0).
	CalZeroPoint float64 `json:"calZeroPoint"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CalBgIntensity            respjson.Field
		CalExtinctionCoeff        respjson.Field
		CalExtinctionCoeffMaxUnc  respjson.Field
		CalExtinctionCoeffUnc     respjson.Field
		CalNumCorrelatedStars     respjson.Field
		CalNumDetectedStars       respjson.Field
		CalSkyBg                  respjson.Field
		CalSpectralFilterSolarMag respjson.Field
		CalTime                   respjson.Field
		CalType                   respjson.Field
		CalZeroPoint              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SoiObservationSetFullCalibration) RawJSON() string { return r.JSON.raw }
func (r *SoiObservationSetFullCalibration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Optical SOI record contains observation information taken from a sensor about
// a Space Object.
type SoiObservationSetFullOpticalSoiObservationList struct {
	// Observation detection start time in ISO 8601 UTC with microsecond precision.
	ObStartTime time.Time `json:"obStartTime,required" format:"date-time"`
	// The reference number, x, where x ranges from 1 to n, where n is the number
	// specified in spectralFilters that corresponds to the spectral filter used.
	CurrentSpectralFilterNum int64 `json:"currentSpectralFilterNum"`
	// Array of declination rate values, in degrees per second, measuring the rate
	// speed at which an object's declination changes over time, for each element in
	// the intensities field, at the middle of the frame's exposure time.
	DeclinationRates []float64 `json:"declinationRates"`
	// Array of declination values, in degrees, of the Target object from the frame of
	// reference of the sensor. A value is provided for each element in the intensities
	// field, at the middle of the frame’s exposure time.
	Declinations []float64 `json:"declinations"`
	// Image exposure duration in seconds.
	ExpDuration float64 `json:"expDuration"`
	// Array of coefficients for how much signal would be lost to atmospheric
	// attenuation for a star at zenith for each element in intensities, in magnitudes
	// per air mass.
	ExtinctionCoeffs []float64 `json:"extinctionCoeffs"`
	// Array of extinction coefficient uncertainties for each element in intensities.
	// Each value represents the difference between the measured brightness and
	// predicted brightness of the star with the extinction removed, making it
	// exo-atmospheric (e.g. -5.0 to 30.0).
	ExtinctionCoeffsUnc []float64 `json:"extinctionCoeffsUnc"`
	// Array of intensities of the Space Object for observations, in kilowatts per
	// steradian per micrometer.
	Intensities []float64 `json:"intensities"`
	// Array of start times for each intensity measurement. The 1st value in the array
	// will match obStartTime.
	IntensityTimes []time.Time `json:"intensityTimes" format:"date-time"`
	// Array of local average Sky Background signals, in magnitudes, with a value
	// corresponding to the time of each intensity measurement. Sky Background refers
	// to the incoming light from an apparently empty part of the night sky.
	LocalSkyBgs []float64 `json:"localSkyBgs"`
	// Array of uncertainty of the local average Sky Background signal, in magnitudes,
	// with a value corresponding to the time of each intensity measurement.
	LocalSkyBgsUnc []float64 `json:"localSkyBgsUnc"`
	// Array of the number of correlated stars in the FOV with a value for each element
	// in the intensities field.
	NumCorrelatedStars []int64 `json:"numCorrelatedStars"`
	// Array of the number of detected stars in the FOV with a value for each element
	// in the intensities field.
	NumDetectedStars []int64 `json:"numDetectedStars"`
	// Array of values giving the percent of pixels that make up the object signal that
	// are beyond the saturation point for the sensor, with a value for each element in
	// the intensities field.
	PercentSats []float64 `json:"percentSats"`
	// Array of right ascension rate values, in degrees per second, measuring the rate
	// the telescope is moving to track the Target object from the frame of reference
	// of the sensor, for each element in the intensities field, at the middle of the
	// frame’s exposure time.
	RaRates []float64 `json:"raRates"`
	// Array of right ascension values, in degrees, of the Target object from the frame
	// of reference of the sensor. A value is provided for each element in the
	// intensities field.
	Ras []float64 `json:"ras"`
	// Array of average Sky Background signals, in magnitudes, with a value
	// corresponding to the time of each intensity measurement. Sky Background refers
	// to the incoming light from an apparently empty part of the night sky.
	SkyBgs []float64 `json:"skyBgs"`
	// Array of values for the zero-point in magnitudes, calculated at the time of each
	// intensity measurement. It is the difference between the catalog mag and
	// instrumental mag for a set of standard stars (e.g. -5.0 to 30.0).
	ZeroPoints []float64 `json:"zeroPoints"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObStartTime              respjson.Field
		CurrentSpectralFilterNum respjson.Field
		DeclinationRates         respjson.Field
		Declinations             respjson.Field
		ExpDuration              respjson.Field
		ExtinctionCoeffs         respjson.Field
		ExtinctionCoeffsUnc      respjson.Field
		Intensities              respjson.Field
		IntensityTimes           respjson.Field
		LocalSkyBgs              respjson.Field
		LocalSkyBgsUnc           respjson.Field
		NumCorrelatedStars       respjson.Field
		NumDetectedStars         respjson.Field
		PercentSats              respjson.Field
		RaRates                  respjson.Field
		Ras                      respjson.Field
		SkyBgs                   respjson.Field
		ZeroPoints               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SoiObservationSetFullOpticalSoiObservationList) RawJSON() string { return r.JSON.raw }
func (r *SoiObservationSetFullOpticalSoiObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Radar SOI record contains observation information taken from a sensor about a
// Space Object.
type SoiObservationSetFullRadarSoiObservationList struct {
	// Observation detection start time in ISO 8601 UTC format with microsecond
	// precision.
	ObStartTime time.Time `json:"obStartTime,required" format:"date-time"`
	// Array of the aspect angle at the center of the image in degrees. The 'tovs' and
	// 'aspectAngles' arrays must match in size, if 'aspectAngles' is provided.
	AspectAngles []float64 `json:"aspectAngles"`
	// Array of sensor azimuth angle biases in degrees. The 'tovs' and 'azimuthBiases'
	// arrays must match in size, if 'azimuthBiases' is provided.
	AzimuthBiases []float64 `json:"azimuthBiases"`
	// Array of the azimuth rate of target at image center in degrees per second. The
	// 'tovs' and 'azimuthRates' arrays must match in size, if 'azimuthRates' is
	// provided. If there is an associated image the azimuth rate is assumed to be at
	// image center.
	AzimuthRates []float64 `json:"azimuthRates"`
	// Array of the azimuth angle to target at image center in degrees. The 'tovs' and
	// 'azimuths' arrays must match in size, if 'azimuths' is provided. If there is an
	// associated image the azimuth angle is assumed to be at image center.
	Azimuths []float64 `json:"azimuths"`
	// Beta angle (between target and radar-image frame z axis) in degrees.
	Beta float64 `json:"beta"`
	// Radar center frequency of the radar in hertz.
	CenterFrequency float64 `json:"centerFrequency"`
	// Array of cross-range resolutions (accounting for weighting function) in
	// kilometers. The 'tovs' and 'crossRangeRes' arrays must match in size, if
	// 'crossRangeRes' is provided.
	CrossRangeRes []float64 `json:"crossRangeRes"`
	// Array of average Interpulse spacing in seconds. The 'tovs' and 'deltaTimes'
	// arrays must match in size, if 'deltaTimes' is provided.
	DeltaTimes []float64 `json:"deltaTimes"`
	// Array of conversion factors between Doppler in hertz and cross-range in meters.
	// The 'tovs' and 'doppler2XRs' arrays must match in size, if 'doppler2XRs' is
	// provided.
	Doppler2XRs []float64 `json:"doppler2XRs"`
	// Array of sensor elevation biases in degrees. The 'tovs' and 'elevationBiases'
	// arrays must match in size, if 'elevationBiases' is provided.
	ElevationBiases []float64 `json:"elevationBiases"`
	// Array of the elevation rate of target at image center in degrees per second. The
	// 'tovs' and 'elevationRates' arrays must match in size, if 'elevationRates' is
	// provided. If there is an associated image the elevation rate is assumed to be at
	// image center.
	ElevationRates []float64 `json:"elevationRates"`
	// Array of the elevation angle to target at image center in degrees. The 'tovs'
	// and 'elevations' arrays must match in size, if 'elevations' is provided. If
	// there is an associated image the elevation angle is assumed to be at image
	// center.
	Elevations []float64 `json:"elevations"`
	// Optional id of assumed AttitudeSet of object being observed.
	IDAttitudeSet string `json:"idAttitudeSet"`
	// Optional id of assumed StateVector of object being observed.
	IDStateVector string `json:"idStateVector"`
	// Array of Integration angles in degrees. The 'tovs' and 'integrationAngles'
	// arrays must match in size, if 'integrationAngles' is provided.
	IntegrationAngles []float64 `json:"integrationAngles"`
	// Kappa angle (between radar-line-of-sight and target-frame x axis) in degrees.
	Kappa float64 `json:"kappa"`
	// Array of the peak pixel amplitude for each image in decibels. The 'tovs' and
	// 'peakAmplitudes' arrays must match in size, if 'peakAmplitudes' is provided.
	PeakAmplitudes []float64 `json:"peakAmplitudes"`
	// Array of sensor polarizations when collecting the data. Polarization is a
	// property of electromagnetic waves that describes the orientation of their
	// oscillations. Possible values are H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	Polarizations []string `json:"polarizations"`
	// Array of the component of target angular velocity observable by radar in radians
	// per second. The 'tovs' and 'projAngVels' arrays must match in size, if
	// 'projAngVels' is provided.
	ProjAngVels []float64 `json:"projAngVels"`
	// Bandwidth of radar pulse in hertz.
	PulseBandwidth float64 `json:"pulseBandwidth"`
	// Array of the range accelerations of target in kilometers per second squared. The
	// 'tovs' and 'rangeAccels' arrays must match in size, if 'rangeAccels' is
	// provided. If there is an associated image the range acceleration is assumed to
	// be at image center.
	RangeAccels []float64 `json:"rangeAccels"`
	// Array of sensor range biases in kilometers. The 'tovs' and 'rangeBiases' arrays
	// must match in size, if 'rangeBiases' is provided.
	RangeBiases []float64 `json:"rangeBiases"`
	// Array of the range rate of target at image center in kilometers per second. The
	// 'tovs' and 'rangeRates' arrays must match in size, if 'rangeRates' is provided.
	// If there is an associated image the range rate is assumed to be at image center.
	RangeRates []float64 `json:"rangeRates"`
	// Array of the range to target at image center in kilometers. The 'tovs' and
	// 'ranges' arrays must match in size, if 'ranges' is provided. If there is an
	// associated image the range is assumed to be at image center.
	Ranges []float64 `json:"ranges"`
	// Array of error estimates of RCS values, in square meters.
	RcsErrorEsts []float64 `json:"rcsErrorEsts"`
	// Array of observed radar cross section (RCS) values, in square meters.
	RcsValues []float64 `json:"rcsValues"`
	// Array of range sample spacing in meters. The 'tovs' and 'rspaces' arrays must
	// match in size, if 'rspaces' is provided.
	Rspaces []float64 `json:"rspaces"`
	// Array of spectral widths, in hertz. The spectral width of a satellite can help
	// determine if an object is stable or tumbling which is often useful to
	// distinguish a rocket body from an active stabilized payload or to deduce a
	// rotational period of slowly tumbling objects at GEO.
	SpectralWidths []float64 `json:"spectralWidths"`
	// Array of the times of validity in ISO 8601 UTC format with microsecond
	// precision.
	Tovs []time.Time `json:"tovs" format:"date-time"`
	// Array of the cartesian X accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xaccel' arrays must match in size, if 'xaccel' is provided.
	Xaccel []float64 `json:"xaccel"`
	// Array of the cartesian X positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xpos' arrays must match in size, if 'xpos' is provided.
	Xpos []float64 `json:"xpos"`
	// Array of cross-range sample spacing in meters. The 'tovs' and 'xspaces' arrays
	// must match in size, if 'xspaces' is provided.
	Xspaces []float64 `json:"xspaces"`
	// Array of the cartesian X velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xvel' arrays must match in size, if 'xvel' is provided.
	Xvel []float64 `json:"xvel"`
	// Array of the cartesian Y accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'yaccel' arrays must match in size, if 'yaccel' is provided.
	Yaccel []float64 `json:"yaccel"`
	// Array of the cartesian Y positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'ypos' arrays must match in size, if 'ypos' is provided.
	Ypos []float64 `json:"ypos"`
	// Array of the cartesian Y velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'yvel' arrays must match in size, if 'yvel' is provided.
	Yvel []float64 `json:"yvel"`
	// Array of the cartesian Z accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zaccel' arrays must match in size, if 'zaccel' is provided.
	Zaccel []float64 `json:"zaccel"`
	// Array of the cartesian Z positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zpos' arrays must match in size, if 'zpos' is provided.
	Zpos []float64 `json:"zpos"`
	// Array of the cartesian Z velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zvel' arrays must match in size, if 'zvel' is provided.
	Zvel []float64 `json:"zvel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObStartTime       respjson.Field
		AspectAngles      respjson.Field
		AzimuthBiases     respjson.Field
		AzimuthRates      respjson.Field
		Azimuths          respjson.Field
		Beta              respjson.Field
		CenterFrequency   respjson.Field
		CrossRangeRes     respjson.Field
		DeltaTimes        respjson.Field
		Doppler2XRs       respjson.Field
		ElevationBiases   respjson.Field
		ElevationRates    respjson.Field
		Elevations        respjson.Field
		IDAttitudeSet     respjson.Field
		IDStateVector     respjson.Field
		IntegrationAngles respjson.Field
		Kappa             respjson.Field
		PeakAmplitudes    respjson.Field
		Polarizations     respjson.Field
		ProjAngVels       respjson.Field
		PulseBandwidth    respjson.Field
		RangeAccels       respjson.Field
		RangeBiases       respjson.Field
		RangeRates        respjson.Field
		Ranges            respjson.Field
		RcsErrorEsts      respjson.Field
		RcsValues         respjson.Field
		Rspaces           respjson.Field
		SpectralWidths    respjson.Field
		Tovs              respjson.Field
		Xaccel            respjson.Field
		Xpos              respjson.Field
		Xspaces           respjson.Field
		Xvel              respjson.Field
		Yaccel            respjson.Field
		Ypos              respjson.Field
		Yvel              respjson.Field
		Zaccel            respjson.Field
		Zpos              respjson.Field
		Zvel              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SoiObservationSetFullRadarSoiObservationList) RawJSON() string { return r.JSON.raw }
func (r *SoiObservationSetFullRadarSoiObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reference frame of the observation measurements. If the referenceFrame is
// null it is assumed to be J2000.
type SoiObservationSetFullReferenceFrame string

const (
	SoiObservationSetFullReferenceFrameJ2000   SoiObservationSetFullReferenceFrame = "J2000"
	SoiObservationSetFullReferenceFrameEfgTdr  SoiObservationSetFullReferenceFrame = "EFG/TDR"
	SoiObservationSetFullReferenceFrameEcrEcef SoiObservationSetFullReferenceFrame = "ECR/ECEF"
	SoiObservationSetFullReferenceFrameTeme    SoiObservationSetFullReferenceFrame = "TEME"
	SoiObservationSetFullReferenceFrameItrf    SoiObservationSetFullReferenceFrame = "ITRF"
	SoiObservationSetFullReferenceFrameGcrf    SoiObservationSetFullReferenceFrame = "GCRF"
)

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type SoiObservationSetFullSenReferenceFrame string

const (
	SoiObservationSetFullSenReferenceFrameJ2000   SoiObservationSetFullSenReferenceFrame = "J2000"
	SoiObservationSetFullSenReferenceFrameEfgTdr  SoiObservationSetFullSenReferenceFrame = "EFG/TDR"
	SoiObservationSetFullSenReferenceFrameEcrEcef SoiObservationSetFullSenReferenceFrame = "ECR/ECEF"
	SoiObservationSetFullSenReferenceFrameTeme    SoiObservationSetFullSenReferenceFrame = "TEME"
	SoiObservationSetFullSenReferenceFrameItrf    SoiObservationSetFullSenReferenceFrame = "ITRF"
	SoiObservationSetFullSenReferenceFrameGcrf    SoiObservationSetFullSenReferenceFrame = "GCRF"
)

type SoiObservationSetHistoryListParams struct {
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime time.Time `query:"startTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SoiObservationSetHistoryListParams]'s query parameters as
// `url.Values`.
func (r SoiObservationSetHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SoiObservationSetHistoryAodrParams struct {
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime time.Time `query:"startTime,required" format:"date-time" json:"-"`
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

// URLQuery serializes [SoiObservationSetHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r SoiObservationSetHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SoiObservationSetHistoryCountParams struct {
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SoiObservationSetHistoryCountParams]'s query parameters as
// `url.Values`.
func (r SoiObservationSetHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
