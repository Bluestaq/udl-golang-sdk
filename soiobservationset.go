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

// SoiObservationSetService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSoiObservationSetService] method instead.
type SoiObservationSetService struct {
	Options []option.RequestOption
	History SoiObservationSetHistoryService
}

// NewSoiObservationSetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSoiObservationSetService(opts ...option.RequestOption) (r SoiObservationSetService) {
	r = SoiObservationSetService{}
	r.Options = opts
	r.History = NewSoiObservationSetHistoryService(opts...)
	return
}

// Service operation to take a single SOIObservationSet record as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SoiObservationSetService) New(ctx context.Context, body SoiObservationSetNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/soiobservationset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters.
// The query will return the SOI Observation Sets and not the associated SOI
// Observations. See the queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for
// more details on valid/required query parameter information.
func (r *SoiObservationSetService) List(ctx context.Context, query SoiObservationSetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SoiObservationSetListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/soiobservationset"
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

// Service operation to dynamically query data by a variety of query parameters.
// The query will return the SOI Observation Sets and not the associated SOI
// Observations. See the queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for
// more details on valid/required query parameter information.
func (r *SoiObservationSetService) ListAutoPaging(ctx context.Context, query SoiObservationSetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SoiObservationSetListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SoiObservationSetService) Count(ctx context.Context, query SoiObservationSetCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/soiobservationset/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// SOIObservationSet records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *SoiObservationSetService) NewBulk(ctx context.Context, body SoiObservationSetNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/soiobservationset/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single SOIObservationSet by its unique ID passed as a
// path parameter.
func (r *SoiObservationSetService) Get(ctx context.Context, id string, query SoiObservationSetGetParams, opts ...option.RequestOption) (res *SoiObservationSetFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/soiobservationset/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SoiObservationSetService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SoiObservationSetQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/soiobservationset/queryhelp"
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
func (r *SoiObservationSetService) Tuple(ctx context.Context, query SoiObservationSetTupleParams, opts ...option.RequestOption) (res *[]SoiObservationSetFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/soiobservationset/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple SOIObservationSet records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SoiObservationSetService) UnvalidatedPublish(ctx context.Context, body SoiObservationSetUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-soiobservationset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// These services provide operations for posting space object idenfification
// observation sets.
type SoiObservationSetListResponse struct {
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
	DataMode SoiObservationSetListResponseDataMode `json:"dataMode,required"`
	// The number of observation records in the set.
	NumObs int64 `json:"numObs,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Observation type (OPTICAL, RADAR).
	//
	// Any of "OPTICAL", "RADAR".
	Type SoiObservationSetListResponseType `json:"type,required"`
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
	Calibrations []SoiObservationSetListResponseCalibration `json:"calibrations"`
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
	// The reference frame of the observation measurements. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame SoiObservationSetListResponseReferenceFrame `json:"referenceFrame"`
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
	SenReferenceFrame SoiObservationSetListResponseSenReferenceFrame `json:"senReferenceFrame"`
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
func (r SoiObservationSetListResponse) RawJSON() string { return r.JSON.raw }
func (r *SoiObservationSetListResponse) UnmarshalJSON(data []byte) error {
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
type SoiObservationSetListResponseDataMode string

const (
	SoiObservationSetListResponseDataModeReal      SoiObservationSetListResponseDataMode = "REAL"
	SoiObservationSetListResponseDataModeTest      SoiObservationSetListResponseDataMode = "TEST"
	SoiObservationSetListResponseDataModeSimulated SoiObservationSetListResponseDataMode = "SIMULATED"
	SoiObservationSetListResponseDataModeExercise  SoiObservationSetListResponseDataMode = "EXERCISE"
)

// Observation type (OPTICAL, RADAR).
type SoiObservationSetListResponseType string

const (
	SoiObservationSetListResponseTypeOptical SoiObservationSetListResponseType = "OPTICAL"
	SoiObservationSetListResponseTypeRadar   SoiObservationSetListResponseType = "RADAR"
)

// Schema for SOI Calibration data.
type SoiObservationSetListResponseCalibration struct {
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
func (r SoiObservationSetListResponseCalibration) RawJSON() string { return r.JSON.raw }
func (r *SoiObservationSetListResponseCalibration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reference frame of the observation measurements. If the referenceFrame is
// null it is assumed to be J2000.
type SoiObservationSetListResponseReferenceFrame string

const (
	SoiObservationSetListResponseReferenceFrameJ2000   SoiObservationSetListResponseReferenceFrame = "J2000"
	SoiObservationSetListResponseReferenceFrameEfgTdr  SoiObservationSetListResponseReferenceFrame = "EFG/TDR"
	SoiObservationSetListResponseReferenceFrameEcrEcef SoiObservationSetListResponseReferenceFrame = "ECR/ECEF"
	SoiObservationSetListResponseReferenceFrameTeme    SoiObservationSetListResponseReferenceFrame = "TEME"
	SoiObservationSetListResponseReferenceFrameItrf    SoiObservationSetListResponseReferenceFrame = "ITRF"
	SoiObservationSetListResponseReferenceFrameGcrf    SoiObservationSetListResponseReferenceFrame = "GCRF"
)

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type SoiObservationSetListResponseSenReferenceFrame string

const (
	SoiObservationSetListResponseSenReferenceFrameJ2000   SoiObservationSetListResponseSenReferenceFrame = "J2000"
	SoiObservationSetListResponseSenReferenceFrameEfgTdr  SoiObservationSetListResponseSenReferenceFrame = "EFG/TDR"
	SoiObservationSetListResponseSenReferenceFrameEcrEcef SoiObservationSetListResponseSenReferenceFrame = "ECR/ECEF"
	SoiObservationSetListResponseSenReferenceFrameTeme    SoiObservationSetListResponseSenReferenceFrame = "TEME"
	SoiObservationSetListResponseSenReferenceFrameItrf    SoiObservationSetListResponseSenReferenceFrame = "ITRF"
	SoiObservationSetListResponseSenReferenceFrameGcrf    SoiObservationSetListResponseSenReferenceFrame = "GCRF"
)

type SoiObservationSetQueryhelpResponse struct {
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
func (r SoiObservationSetQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SoiObservationSetQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SoiObservationSetNewParams struct {
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
	DataMode SoiObservationSetNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The number of observation records in the set.
	NumObs int64 `json:"numObs,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Observation type (OPTICAL, RADAR).
	//
	// Any of "OPTICAL", "RADAR".
	Type SoiObservationSetNewParamsType `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The number of pixels binned horizontally.
	BinningHoriz param.Opt[int64] `json:"binningHoriz,omitzero"`
	// The number of pixels binned vertically.
	BinningVert param.Opt[int64] `json:"binningVert,omitzero"`
	// Boolean indicating if a brightness variance change event was detected, based on
	// historical collection data for the object.
	BrightnessVarianceChangeDetected param.Opt[bool] `json:"brightnessVarianceChangeDetected,omitzero"`
	// Type of calibration used by the Sensor (e.g. ALL SKY, DIFFERENTIAL, DEFAULT,
	// NONE).
	CalibrationType param.Opt[string] `json:"calibrationType,omitzero"`
	// Overall qualitative confidence assessment of change detection results (e.g.
	// HIGH, MEDIUM, LOW).
	ChangeConf param.Opt[string] `json:"changeConf,omitzero"`
	// Boolean indicating if any change event was detected, based on historical
	// collection data for the object.
	ChangeDetected param.Opt[bool] `json:"changeDetected,omitzero"`
	// Qualitative Collection Density assessment, with respect to confidence of
	// detecting a change event (e.g. HIGH, MEDIUM, LOW).
	CollectionDensityConf param.Opt[string] `json:"collectionDensityConf,omitzero"`
	// Universally Unique collection ID. Mechanism to correlate Single Point Photometry
	// (SPP) JSON files to images.
	CollectionID param.Opt[string] `json:"collectionId,omitzero"`
	// Mode indicating telescope movement during collection (AUTOTRACK, MANUAL
	// AUTOTRACK, MANUAL RATE TRACK, MANUAL SIDEREAL, SIDEREAL, RATE TRACK).
	CollectionMode param.Opt[string] `json:"collectionMode,omitzero"`
	// Object Correlation Quality value. Measures how close the observed object's orbit
	// is to matching an object in the catalog. The scale of this field may vary
	// depending on provider. Users should consult the data provider to verify the
	// meaning of the value (e.g. A value of 0.0 indicates a high/strong correlation,
	// while a value closer to 1.0 indicates low/weak correlation).
	CorrQuality param.Opt[float64] `json:"corrQuality,omitzero"`
	// Observation set detection end time in ISO 8601 UTC with microsecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The gain used during the collection, in units of photoelectrons per
	// analog-to-digital unit (e-/ADU). If no gain is used, the value = 1.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// ID of the UDL Elset of the Space Object under observation.
	IDElset param.Opt[string] `json:"idElset,omitzero"`
	// ID of the observing sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Line of sight declination at observation set detection end time. Specified in
	// degrees, in the specified referenceFrame. If referenceFrame is null then J2K
	// should be assumed (e.g -30 to 130.0).
	LosDeclinationEnd param.Opt[float64] `json:"losDeclinationEnd,omitzero"`
	// Line of sight declination at observation set detection start time. Specified in
	// degrees, in the specified referenceFrame. If referenceFrame is null then J2K
	// should be assumed (e.g -30 to 130.0).
	LosDeclinationStart param.Opt[float64] `json:"losDeclinationStart,omitzero"`
	// SOI msgCreateDate time in ISO 8601 UTC time, with millisecond precision.
	MsgCreateDate param.Opt[time.Time] `json:"msgCreateDate,omitzero" format:"date-time"`
	// The value is the number of spectral filters used.
	NumSpectralFilters param.Opt[int64] `json:"numSpectralFilters,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the record source to indicate the sensor
	// identifier to which this attitude set applies if this set is reporting a single
	// sensor orientation. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// A threshold for percent of pixels that make up object signal that are beyond the
	// saturation point for the sensor that are removed in the EOSSA file, in range of
	// 0 to 1.
	PercentSatThreshold param.Opt[float64] `json:"percentSatThreshold,omitzero"`
	// Boolean indicating if a periodicity change event was detected, based on
	// historical collection data for the object.
	PeriodicityChangeDetected param.Opt[bool] `json:"periodicityChangeDetected,omitzero"`
	// Qualitative assessment of the periodicity detection results from the Attitude
	// and Shape Retrieval (ASR) Periodicity Assessment (PA) Tool (e.g. HIGH, MEDIUM,
	// LOW).
	PeriodicityDetectionConf param.Opt[string] `json:"periodicityDetectionConf,omitzero"`
	// Qualitative Periodicity Sampling assessment, with respect to confidence of
	// detecting a change event (e.g. HIGH, MEDIUM, LOW).
	PeriodicitySamplingConf param.Opt[string] `json:"periodicitySamplingConf,omitzero"`
	// Pixel array size (height) in pixels.
	PixelArrayHeight param.Opt[int64] `json:"pixelArrayHeight,omitzero"`
	// Pixel array size (width) in pixels.
	PixelArrayWidth param.Opt[int64] `json:"pixelArrayWidth,omitzero"`
	// The maximum valid pixel value.
	PixelMax param.Opt[int64] `json:"pixelMax,omitzero"`
	// The minimum valid pixel value.
	PixelMin param.Opt[int64] `json:"pixelMin,omitzero"`
	// Pointing angle of the Azimuth gimbal/mount at observation set detection end
	// time. Specified in degrees.
	PointingAngleAzEnd param.Opt[float64] `json:"pointingAngleAzEnd,omitzero"`
	// Pointing angle of the Azimuth gimbal/mount at observation set detection start
	// time. Specified in degrees.
	PointingAngleAzStart param.Opt[float64] `json:"pointingAngleAzStart,omitzero"`
	// Pointing angle of the Elevation gimbal/mount at observation set detection end
	// time. Specified in degrees.
	PointingAngleElEnd param.Opt[float64] `json:"pointingAngleElEnd,omitzero"`
	// Pointing angle of the Elevation gimbal/mount at observation set detection start
	// time. Specified in degrees.
	PointingAngleElStart param.Opt[float64] `json:"pointingAngleElStart,omitzero"`
	// Polar angle of the gimbal/mount at observation set detection end time in
	// degrees.
	PolarAngleEnd param.Opt[float64] `json:"polarAngleEnd,omitzero"`
	// Polar angle of the gimbal/mount at observation set detection start time in
	// degrees.
	PolarAngleStart param.Opt[float64] `json:"polarAngleStart,omitzero"`
	// Name of the target satellite.
	SatelliteName param.Opt[string] `json:"satelliteName,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude at startTime (if mobile/onorbit) in kilometers.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at startTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info. -90 to 90 degrees (negative values south of
	// equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at startTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values south of
	// equator).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// ID of the AttitudeSet record for the observing sensor.
	SensorAsID param.Opt[string] `json:"sensorAsId,omitzero"`
	// Cartesian X velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvelx param.Opt[float64] `json:"senvelx,omitzero"`
	// Cartesian Y velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvely param.Opt[float64] `json:"senvely,omitzero"`
	// Cartesian Z velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvelz param.Opt[float64] `json:"senvelz,omitzero"`
	// Cartesian X position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Senx param.Opt[float64] `json:"senx,omitzero"`
	// Cartesian Y position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Seny param.Opt[float64] `json:"seny,omitzero"`
	// Cartesian Z position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Senz param.Opt[float64] `json:"senz,omitzero"`
	// Software Version used to Capture, Process, and Deliver the data.
	SoftwareVersion param.Opt[string] `json:"softwareVersion,omitzero"`
	// The in-band solar magnitude at 1 A.U.
	SolarMag param.Opt[float64] `json:"solarMag,omitzero"`
	// Boolean indicating if a solar phase angle brightness change event was detected,
	// based on historical collection data for the object.
	SolarPhaseAngleBrightnessChangeDetected param.Opt[bool] `json:"solarPhaseAngleBrightnessChangeDetected,omitzero"`
	// Name of the Star Catalog used for photometry and astrometry.
	StarCatName param.Opt[string] `json:"starCatName,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating whether the target object was unable to be correlated to a
	// known object. This flag should only be set to true by data providers after an
	// attempt to correlate to an OnOrbit object was made and failed. If unable to
	// correlate, the 'origObjectId' field may be populated with an internal data
	// provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Key to indicate which, if any of, the pre/post photometer calibrations are valid
	// for use when generating data for the EOSSA file. If the field is not populated,
	// then the provided calibration data will be used when generating the EOSSA file
	// (e.g. PRE, POST, BOTH, NONE).
	ValidCalibrations param.Opt[string] `json:"validCalibrations,omitzero"`
	// Array of SOI Calibrations associated with this SOIObservationSet.
	Calibrations []SoiObservationSetNewParamsCalibration `json:"calibrations,omitzero"`
	// OpticalSOIObservations associated with this SOIObservationSet.
	OpticalSoiObservationList []SoiObservationSetNewParamsOpticalSoiObservationList `json:"opticalSOIObservationList,omitzero"`
	// RadarSOIObservations associated with this RadarSOIObservationSet.
	RadarSoiObservationList []SoiObservationSetNewParamsRadarSoiObservationList `json:"radarSOIObservationList,omitzero"`
	// The reference frame of the observation measurements. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame SoiObservationSetNewParamsReferenceFrame `json:"referenceFrame,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame SoiObservationSetNewParamsSenReferenceFrame `json:"senReferenceFrame,omitzero"`
	// Array of the SpectralFilters keywords, must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	SpectralFilters []string `json:"spectralFilters,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SoiObservationSetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetNewParams) UnmarshalJSON(data []byte) error {
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
type SoiObservationSetNewParamsDataMode string

const (
	SoiObservationSetNewParamsDataModeReal      SoiObservationSetNewParamsDataMode = "REAL"
	SoiObservationSetNewParamsDataModeTest      SoiObservationSetNewParamsDataMode = "TEST"
	SoiObservationSetNewParamsDataModeSimulated SoiObservationSetNewParamsDataMode = "SIMULATED"
	SoiObservationSetNewParamsDataModeExercise  SoiObservationSetNewParamsDataMode = "EXERCISE"
)

// Observation type (OPTICAL, RADAR).
type SoiObservationSetNewParamsType string

const (
	SoiObservationSetNewParamsTypeOptical SoiObservationSetNewParamsType = "OPTICAL"
	SoiObservationSetNewParamsTypeRadar   SoiObservationSetNewParamsType = "RADAR"
)

// Schema for SOI Calibration data.
type SoiObservationSetNewParamsCalibration struct {
	// Background intensity, at calibration, specified in kilowatts per steradian per
	// micrometer.
	CalBgIntensity param.Opt[float64] `json:"calBgIntensity,omitzero"`
	// Coefficient value for how much signal would be lost to atmospheric attenuation
	// for a star at zenith, in magnitudes per air mass.
	CalExtinctionCoeff param.Opt[float64] `json:"calExtinctionCoeff,omitzero"`
	// Maximum extinction coefficient uncertainty in magnitudes, at calibration (e.g.
	// -5.0 to 30.0).
	CalExtinctionCoeffMaxUnc param.Opt[float64] `json:"calExtinctionCoeffMaxUnc,omitzero"`
	// Extinction coefficient uncertainty in magnitudes, at calibration, which
	// represents the difference between the measured brightness and predicted
	// brightness of the star with the extinction removed, making it exo-atmospheric
	// (e.g. -5.0 to 30.0).
	CalExtinctionCoeffUnc param.Opt[float64] `json:"calExtinctionCoeffUnc,omitzero"`
	// Number of correlated stars in the FOV with the target object, at calibration.
	// Can be 0 for narrow FOV sensors.
	CalNumCorrelatedStars param.Opt[int64] `json:"calNumCorrelatedStars,omitzero"`
	// Number of detected stars in the FOV with the target object, at calibration.
	// Helps identify frames with clouds.
	CalNumDetectedStars param.Opt[int64] `json:"calNumDetectedStars,omitzero"`
	// Average Sky Background signals in magnitudes, at calibration. Sky Background
	// refers to the incoming light from an apparently empty part of the night sky.
	CalSkyBg param.Opt[float64] `json:"calSkyBg,omitzero"`
	// In-band solar magnitudes at 1 A.U, at calibration (e.g. -5.0 to 30.0).
	CalSpectralFilterSolarMag param.Opt[float64] `json:"calSpectralFilterSolarMag,omitzero"`
	// Start time of calibration in ISO 8601 UTC time, with millisecond precision.
	CalTime param.Opt[time.Time] `json:"calTime,omitzero" format:"date-time"`
	// Type of calibration (e.g. PRE, MID, POST).
	CalType param.Opt[string] `json:"calType,omitzero"`
	// Value representing the difference between the catalog magnitude and instrumental
	// magnitude for a set of standard stars, at calibration (e.g. -5.0 to 30.0).
	CalZeroPoint param.Opt[float64] `json:"calZeroPoint,omitzero"`
	paramObj
}

func (r SoiObservationSetNewParamsCalibration) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetNewParamsCalibration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetNewParamsCalibration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Optical SOI record contains observation information taken from a sensor about
// a Space Object.
//
// The property ObStartTime is required.
type SoiObservationSetNewParamsOpticalSoiObservationList struct {
	// Observation detection start time in ISO 8601 UTC with microsecond precision.
	ObStartTime time.Time `json:"obStartTime,required" format:"date-time"`
	// The reference number, x, where x ranges from 1 to n, where n is the number
	// specified in spectralFilters that corresponds to the spectral filter used.
	CurrentSpectralFilterNum param.Opt[int64] `json:"currentSpectralFilterNum,omitzero"`
	// Image exposure duration in seconds.
	ExpDuration param.Opt[float64] `json:"expDuration,omitzero"`
	// Array of declination values, in degrees, of the Target object from the frame of
	// reference of the sensor. A value is provided for each element in the intensities
	// field, at the middle of the frame’s exposure time.
	Declinations []float64 `json:"declinations,omitzero"`
	// Array of coefficients for how much signal would be lost to atmospheric
	// attenuation for a star at zenith for each element in intensities, in magnitudes
	// per air mass.
	ExtinctionCoeffs []float64 `json:"extinctionCoeffs,omitzero"`
	// Array of extinction coefficient uncertainties for each element in intensities.
	// Each value represents the difference between the measured brightness and
	// predicted brightness of the star with the extinction removed, making it
	// exo-atmospheric (e.g. -5.0 to 30.0).
	ExtinctionCoeffsUnc []float64 `json:"extinctionCoeffsUnc,omitzero"`
	// Array of intensities of the Space Object for observations, in kilowatts per
	// steradian per micrometer.
	Intensities []float64 `json:"intensities,omitzero"`
	// Array of start times for each intensity measurement. The 1st value in the array
	// will match obStartTime.
	IntensityTimes []time.Time `json:"intensityTimes,omitzero" format:"date-time"`
	// Array of local average Sky Background signals, in magnitudes, with a value
	// corresponding to the time of each intensity measurement. Sky Background refers
	// to the incoming light from an apparently empty part of the night sky.
	LocalSkyBgs []float64 `json:"localSkyBgs,omitzero"`
	// Array of uncertainty of the local average Sky Background signal, in magnitudes,
	// with a value corresponding to the time of each intensity measurement.
	LocalSkyBgsUnc []float64 `json:"localSkyBgsUnc,omitzero"`
	// Array of the number of correlated stars in the FOV with a value for each element
	// in the intensities field.
	NumCorrelatedStars []int64 `json:"numCorrelatedStars,omitzero"`
	// Array of the number of detected stars in the FOV with a value for each element
	// in the intensities field.
	NumDetectedStars []int64 `json:"numDetectedStars,omitzero"`
	// Array of values giving the percent of pixels that make up the object signal that
	// are beyond the saturation point for the sensor, with a value for each element in
	// the intensities field.
	PercentSats []float64 `json:"percentSats,omitzero"`
	// Array of right ascension rate values, in degrees per second, measuring the rate
	// the telescope is moving to track the Target object from the frame of reference
	// of the sensor, for each element in the intensities field, at the middle of the
	// frame’s exposure time.
	RaRates []float64 `json:"raRates,omitzero"`
	// Array of right ascension values, in degrees, of the Target object from the frame
	// of reference of the sensor. A value is provided for each element in the
	// intensities field.
	Ras []float64 `json:"ras,omitzero"`
	// Array of average Sky Background signals, in magnitudes, with a value
	// corresponding to the time of each intensity measurement. Sky Background refers
	// to the incoming light from an apparently empty part of the night sky.
	SkyBgs []float64 `json:"skyBgs,omitzero"`
	// Array of values for the zero-point in magnitudes, calculated at the time of each
	// intensity measurement. It is the difference between the catalog mag and
	// instrumental mag for a set of standard stars (e.g. -5.0 to 30.0).
	ZeroPoints []float64 `json:"zeroPoints,omitzero"`
	paramObj
}

func (r SoiObservationSetNewParamsOpticalSoiObservationList) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetNewParamsOpticalSoiObservationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetNewParamsOpticalSoiObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Radar SOI record contains observation information taken from a sensor about a
// Space Object.
//
// The property ObStartTime is required.
type SoiObservationSetNewParamsRadarSoiObservationList struct {
	// Observation detection start time in ISO 8601 UTC format with microsecond
	// precision.
	ObStartTime time.Time `json:"obStartTime,required" format:"date-time"`
	// Beta angle (between target and radar-image frame z axis) in degrees.
	Beta param.Opt[float64] `json:"beta,omitzero"`
	// Radar center frequency of the radar in hertz.
	CenterFrequency param.Opt[float64] `json:"centerFrequency,omitzero"`
	// Optional id of assumed AttitudeSet of object being observed.
	IDAttitudeSet param.Opt[string] `json:"idAttitudeSet,omitzero"`
	// Optional id of assumed StateVector of object being observed.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Kappa angle (between radar-line-of-sight and target-frame x axis) in degrees.
	Kappa param.Opt[float64] `json:"kappa,omitzero"`
	// Bandwidth of radar pulse in hertz.
	PulseBandwidth param.Opt[float64] `json:"pulseBandwidth,omitzero"`
	// Array of the aspect angle at the center of the image in degrees. The 'tovs' and
	// 'aspectAngles' arrays must match in size, if 'aspectAngles' is provided.
	AspectAngles []float64 `json:"aspectAngles,omitzero"`
	// Array of sensor azimuth angle biases in degrees. The 'tovs' and 'azimuthBiases'
	// arrays must match in size, if 'azimuthBiases' is provided.
	AzimuthBiases []float64 `json:"azimuthBiases,omitzero"`
	// Array of the azimuth rate of target at image center in degrees per second. The
	// 'tovs' and 'azimuthRates' arrays must match in size, if 'azimuthRates' is
	// provided. If there is an associated image the azimuth rate is assumed to be at
	// image center.
	AzimuthRates []float64 `json:"azimuthRates,omitzero"`
	// Array of the azimuth angle to target at image center in degrees. The 'tovs' and
	// 'azimuths' arrays must match in size, if 'azimuths' is provided. If there is an
	// associated image the azimuth angle is assumed to be at image center.
	Azimuths []float64 `json:"azimuths,omitzero"`
	// Array of cross-range resolutions (accounting for weighting function) in
	// kilometers. The 'tovs' and 'crossRangeRes' arrays must match in size, if
	// 'crossRangeRes' is provided.
	CrossRangeRes []float64 `json:"crossRangeRes,omitzero"`
	// Array of average Interpulse spacing in seconds. The 'tovs' and 'deltaTimes'
	// arrays must match in size, if 'deltaTimes' is provided.
	DeltaTimes []float64 `json:"deltaTimes,omitzero"`
	// Array of conversion factors between Doppler in hertz and cross-range in meters.
	// The 'tovs' and 'doppler2XRs' arrays must match in size, if 'doppler2XRs' is
	// provided.
	Doppler2XRs []float64 `json:"doppler2XRs,omitzero"`
	// Array of sensor elevation biases in degrees. The 'tovs' and 'elevationBiases'
	// arrays must match in size, if 'elevationBiases' is provided.
	ElevationBiases []float64 `json:"elevationBiases,omitzero"`
	// Array of the elevation rate of target at image center in degrees per second. The
	// 'tovs' and 'elevationRates' arrays must match in size, if 'elevationRates' is
	// provided. If there is an associated image the elevation rate is assumed to be at
	// image center.
	ElevationRates []float64 `json:"elevationRates,omitzero"`
	// Array of the elevation angle to target at image center in degrees. The 'tovs'
	// and 'elevations' arrays must match in size, if 'elevations' is provided. If
	// there is an associated image the elevation angle is assumed to be at image
	// center.
	Elevations []float64 `json:"elevations,omitzero"`
	// Array of Integration angles in degrees. The 'tovs' and 'integrationAngles'
	// arrays must match in size, if 'integrationAngles' is provided.
	IntegrationAngles []float64 `json:"integrationAngles,omitzero"`
	// Array of the peak pixel amplitude for each image in decibels. The 'tovs' and
	// 'peakAmplitudes' arrays must match in size, if 'peakAmplitudes' is provided.
	PeakAmplitudes []float64 `json:"peakAmplitudes,omitzero"`
	// Array of sensor polarizations when collecting the data. Polarization is a
	// property of electromagnetic waves that describes the orientation of their
	// oscillations. Possible values are H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	Polarizations []string `json:"polarizations,omitzero"`
	// Array of the component of target angular velocity observable by radar in radians
	// per second. The 'tovs' and 'projAngVels' arrays must match in size, if
	// 'projAngVels' is provided.
	ProjAngVels []float64 `json:"projAngVels,omitzero"`
	// Array of the range accelerations of target in kilometers per second squared. The
	// 'tovs' and 'rangeAccels' arrays must match in size, if 'rangeAccels' is
	// provided. If there is an associated image the range acceleration is assumed to
	// be at image center.
	RangeAccels []float64 `json:"rangeAccels,omitzero"`
	// Array of sensor range biases in kilometers. The 'tovs' and 'rangeBiases' arrays
	// must match in size, if 'rangeBiases' is provided.
	RangeBiases []float64 `json:"rangeBiases,omitzero"`
	// Array of the range rate of target at image center in kilometers per second. The
	// 'tovs' and 'rangeRates' arrays must match in size, if 'rangeRates' is provided.
	// If there is an associated image the range rate is assumed to be at image center.
	RangeRates []float64 `json:"rangeRates,omitzero"`
	// Array of the range to target at image center in kilometers. The 'tovs' and
	// 'ranges' arrays must match in size, if 'ranges' is provided. If there is an
	// associated image the range is assumed to be at image center.
	Ranges []float64 `json:"ranges,omitzero"`
	// Array of error estimates of RCS values, in square meters.
	RcsErrorEsts []float64 `json:"rcsErrorEsts,omitzero"`
	// Array of observed radar cross section (RCS) values, in square meters.
	RcsValues []float64 `json:"rcsValues,omitzero"`
	// Array of range sample spacing in meters. The 'tovs' and 'rspaces' arrays must
	// match in size, if 'rspaces' is provided.
	Rspaces []float64 `json:"rspaces,omitzero"`
	// Array of spectral widths, in hertz. The spectral width of a satellite can help
	// determine if an object is stable or tumbling which is often useful to
	// distinguish a rocket body from an active stabilized payload or to deduce a
	// rotational period of slowly tumbling objects at GEO.
	SpectralWidths []float64 `json:"spectralWidths,omitzero"`
	// Array of the times of validity in ISO 8601 UTC format with microsecond
	// precision.
	Tovs []time.Time `json:"tovs,omitzero" format:"date-time"`
	// Array of the cartesian X accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xaccel' arrays must match in size, if 'xaccel' is provided.
	Xaccel []float64 `json:"xaccel,omitzero"`
	// Array of the cartesian X positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xpos' arrays must match in size, if 'xpos' is provided.
	Xpos []float64 `json:"xpos,omitzero"`
	// Array of cross-range sample spacing in meters. The 'tovs' and 'xspaces' arrays
	// must match in size, if 'xspaces' is provided.
	Xspaces []float64 `json:"xspaces,omitzero"`
	// Array of the cartesian X velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xvel' arrays must match in size, if 'xvel' is provided.
	Xvel []float64 `json:"xvel,omitzero"`
	// Array of the cartesian Y accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'yaccel' arrays must match in size, if 'yaccel' is provided.
	Yaccel []float64 `json:"yaccel,omitzero"`
	// Array of the cartesian Y positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'ypos' arrays must match in size, if 'ypos' is provided.
	Ypos []float64 `json:"ypos,omitzero"`
	// Array of the cartesian Y velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'yvel' arrays must match in size, if 'yvel' is provided.
	Yvel []float64 `json:"yvel,omitzero"`
	// Array of the cartesian Z accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zaccel' arrays must match in size, if 'zaccel' is provided.
	Zaccel []float64 `json:"zaccel,omitzero"`
	// Array of the cartesian Z positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zpos' arrays must match in size, if 'zpos' is provided.
	Zpos []float64 `json:"zpos,omitzero"`
	// Array of the cartesian Z velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zvel' arrays must match in size, if 'zvel' is provided.
	Zvel []float64 `json:"zvel,omitzero"`
	paramObj
}

func (r SoiObservationSetNewParamsRadarSoiObservationList) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetNewParamsRadarSoiObservationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetNewParamsRadarSoiObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reference frame of the observation measurements. If the referenceFrame is
// null it is assumed to be J2000.
type SoiObservationSetNewParamsReferenceFrame string

const (
	SoiObservationSetNewParamsReferenceFrameJ2000   SoiObservationSetNewParamsReferenceFrame = "J2000"
	SoiObservationSetNewParamsReferenceFrameEfgTdr  SoiObservationSetNewParamsReferenceFrame = "EFG/TDR"
	SoiObservationSetNewParamsReferenceFrameEcrEcef SoiObservationSetNewParamsReferenceFrame = "ECR/ECEF"
	SoiObservationSetNewParamsReferenceFrameTeme    SoiObservationSetNewParamsReferenceFrame = "TEME"
	SoiObservationSetNewParamsReferenceFrameItrf    SoiObservationSetNewParamsReferenceFrame = "ITRF"
	SoiObservationSetNewParamsReferenceFrameGcrf    SoiObservationSetNewParamsReferenceFrame = "GCRF"
)

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type SoiObservationSetNewParamsSenReferenceFrame string

const (
	SoiObservationSetNewParamsSenReferenceFrameJ2000   SoiObservationSetNewParamsSenReferenceFrame = "J2000"
	SoiObservationSetNewParamsSenReferenceFrameEfgTdr  SoiObservationSetNewParamsSenReferenceFrame = "EFG/TDR"
	SoiObservationSetNewParamsSenReferenceFrameEcrEcef SoiObservationSetNewParamsSenReferenceFrame = "ECR/ECEF"
	SoiObservationSetNewParamsSenReferenceFrameTeme    SoiObservationSetNewParamsSenReferenceFrame = "TEME"
	SoiObservationSetNewParamsSenReferenceFrameItrf    SoiObservationSetNewParamsSenReferenceFrame = "ITRF"
	SoiObservationSetNewParamsSenReferenceFrameGcrf    SoiObservationSetNewParamsSenReferenceFrame = "GCRF"
)

type SoiObservationSetListParams struct {
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SoiObservationSetListParams]'s query parameters as
// `url.Values`.
func (r SoiObservationSetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SoiObservationSetCountParams struct {
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SoiObservationSetCountParams]'s query parameters as
// `url.Values`.
func (r SoiObservationSetCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SoiObservationSetNewBulkParams struct {
	Body []SoiObservationSetNewBulkParamsBody
	paramObj
}

func (r SoiObservationSetNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SoiObservationSetNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide operations for posting space object idenfification
// observation sets.
//
// The properties ClassificationMarking, DataMode, NumObs, Source, StartTime, Type
// are required.
type SoiObservationSetNewBulkParamsBody struct {
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
	// The number of observation records in the set.
	NumObs int64 `json:"numObs,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Observation type (OPTICAL, RADAR).
	//
	// Any of "OPTICAL", "RADAR".
	Type string `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The number of pixels binned horizontally.
	BinningHoriz param.Opt[int64] `json:"binningHoriz,omitzero"`
	// The number of pixels binned vertically.
	BinningVert param.Opt[int64] `json:"binningVert,omitzero"`
	// Boolean indicating if a brightness variance change event was detected, based on
	// historical collection data for the object.
	BrightnessVarianceChangeDetected param.Opt[bool] `json:"brightnessVarianceChangeDetected,omitzero"`
	// Type of calibration used by the Sensor (e.g. ALL SKY, DIFFERENTIAL, DEFAULT,
	// NONE).
	CalibrationType param.Opt[string] `json:"calibrationType,omitzero"`
	// Overall qualitative confidence assessment of change detection results (e.g.
	// HIGH, MEDIUM, LOW).
	ChangeConf param.Opt[string] `json:"changeConf,omitzero"`
	// Boolean indicating if any change event was detected, based on historical
	// collection data for the object.
	ChangeDetected param.Opt[bool] `json:"changeDetected,omitzero"`
	// Qualitative Collection Density assessment, with respect to confidence of
	// detecting a change event (e.g. HIGH, MEDIUM, LOW).
	CollectionDensityConf param.Opt[string] `json:"collectionDensityConf,omitzero"`
	// Universally Unique collection ID. Mechanism to correlate Single Point Photometry
	// (SPP) JSON files to images.
	CollectionID param.Opt[string] `json:"collectionId,omitzero"`
	// Mode indicating telescope movement during collection (AUTOTRACK, MANUAL
	// AUTOTRACK, MANUAL RATE TRACK, MANUAL SIDEREAL, SIDEREAL, RATE TRACK).
	CollectionMode param.Opt[string] `json:"collectionMode,omitzero"`
	// Object Correlation Quality value. Measures how close the observed object's orbit
	// is to matching an object in the catalog. The scale of this field may vary
	// depending on provider. Users should consult the data provider to verify the
	// meaning of the value (e.g. A value of 0.0 indicates a high/strong correlation,
	// while a value closer to 1.0 indicates low/weak correlation).
	CorrQuality param.Opt[float64] `json:"corrQuality,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Observation set detection end time in ISO 8601 UTC with microsecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The gain used during the collection, in units of photoelectrons per
	// analog-to-digital unit (e-/ADU). If no gain is used, the value = 1.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// ID of the UDL Elset of the Space Object under observation.
	IDElset param.Opt[string] `json:"idElset,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// ID of the observing sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Line of sight declination at observation set detection end time. Specified in
	// degrees, in the specified referenceFrame. If referenceFrame is null then J2K
	// should be assumed (e.g -30 to 130.0).
	LosDeclinationEnd param.Opt[float64] `json:"losDeclinationEnd,omitzero"`
	// Line of sight declination at observation set detection start time. Specified in
	// degrees, in the specified referenceFrame. If referenceFrame is null then J2K
	// should be assumed (e.g -30 to 130.0).
	LosDeclinationStart param.Opt[float64] `json:"losDeclinationStart,omitzero"`
	// SOI msgCreateDate time in ISO 8601 UTC time, with millisecond precision.
	MsgCreateDate param.Opt[time.Time] `json:"msgCreateDate,omitzero" format:"date-time"`
	// The value is the number of spectral filters used.
	NumSpectralFilters param.Opt[int64] `json:"numSpectralFilters,omitzero"`
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
	// Optional identifier provided by the record source to indicate the sensor
	// identifier to which this attitude set applies if this set is reporting a single
	// sensor orientation. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// A threshold for percent of pixels that make up object signal that are beyond the
	// saturation point for the sensor that are removed in the EOSSA file, in range of
	// 0 to 1.
	PercentSatThreshold param.Opt[float64] `json:"percentSatThreshold,omitzero"`
	// Boolean indicating if a periodicity change event was detected, based on
	// historical collection data for the object.
	PeriodicityChangeDetected param.Opt[bool] `json:"periodicityChangeDetected,omitzero"`
	// Qualitative assessment of the periodicity detection results from the Attitude
	// and Shape Retrieval (ASR) Periodicity Assessment (PA) Tool (e.g. HIGH, MEDIUM,
	// LOW).
	PeriodicityDetectionConf param.Opt[string] `json:"periodicityDetectionConf,omitzero"`
	// Qualitative Periodicity Sampling assessment, with respect to confidence of
	// detecting a change event (e.g. HIGH, MEDIUM, LOW).
	PeriodicitySamplingConf param.Opt[string] `json:"periodicitySamplingConf,omitzero"`
	// Pixel array size (height) in pixels.
	PixelArrayHeight param.Opt[int64] `json:"pixelArrayHeight,omitzero"`
	// Pixel array size (width) in pixels.
	PixelArrayWidth param.Opt[int64] `json:"pixelArrayWidth,omitzero"`
	// The maximum valid pixel value.
	PixelMax param.Opt[int64] `json:"pixelMax,omitzero"`
	// The minimum valid pixel value.
	PixelMin param.Opt[int64] `json:"pixelMin,omitzero"`
	// Pointing angle of the Azimuth gimbal/mount at observation set detection end
	// time. Specified in degrees.
	PointingAngleAzEnd param.Opt[float64] `json:"pointingAngleAzEnd,omitzero"`
	// Pointing angle of the Azimuth gimbal/mount at observation set detection start
	// time. Specified in degrees.
	PointingAngleAzStart param.Opt[float64] `json:"pointingAngleAzStart,omitzero"`
	// Pointing angle of the Elevation gimbal/mount at observation set detection end
	// time. Specified in degrees.
	PointingAngleElEnd param.Opt[float64] `json:"pointingAngleElEnd,omitzero"`
	// Pointing angle of the Elevation gimbal/mount at observation set detection start
	// time. Specified in degrees.
	PointingAngleElStart param.Opt[float64] `json:"pointingAngleElStart,omitzero"`
	// Polar angle of the gimbal/mount at observation set detection end time in
	// degrees.
	PolarAngleEnd param.Opt[float64] `json:"polarAngleEnd,omitzero"`
	// Polar angle of the gimbal/mount at observation set detection start time in
	// degrees.
	PolarAngleStart param.Opt[float64] `json:"polarAngleStart,omitzero"`
	// Name of the target satellite.
	SatelliteName param.Opt[string] `json:"satelliteName,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude at startTime (if mobile/onorbit) in kilometers.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at startTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info. -90 to 90 degrees (negative values south of
	// equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at startTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values south of
	// equator).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// ID of the AttitudeSet record for the observing sensor.
	SensorAsID param.Opt[string] `json:"sensorAsId,omitzero"`
	// Cartesian X velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvelx param.Opt[float64] `json:"senvelx,omitzero"`
	// Cartesian Y velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvely param.Opt[float64] `json:"senvely,omitzero"`
	// Cartesian Z velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvelz param.Opt[float64] `json:"senvelz,omitzero"`
	// Cartesian X position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Senx param.Opt[float64] `json:"senx,omitzero"`
	// Cartesian Y position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Seny param.Opt[float64] `json:"seny,omitzero"`
	// Cartesian Z position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Senz param.Opt[float64] `json:"senz,omitzero"`
	// Software Version used to Capture, Process, and Deliver the data.
	SoftwareVersion param.Opt[string] `json:"softwareVersion,omitzero"`
	// The in-band solar magnitude at 1 A.U.
	SolarMag param.Opt[float64] `json:"solarMag,omitzero"`
	// Boolean indicating if a solar phase angle brightness change event was detected,
	// based on historical collection data for the object.
	SolarPhaseAngleBrightnessChangeDetected param.Opt[bool] `json:"solarPhaseAngleBrightnessChangeDetected,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Name of the Star Catalog used for photometry and astrometry.
	StarCatName param.Opt[string] `json:"starCatName,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating whether the target object was unable to be correlated to a
	// known object. This flag should only be set to true by data providers after an
	// attempt to correlate to an OnOrbit object was made and failed. If unable to
	// correlate, the 'origObjectId' field may be populated with an internal data
	// provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Key to indicate which, if any of, the pre/post photometer calibrations are valid
	// for use when generating data for the EOSSA file. If the field is not populated,
	// then the provided calibration data will be used when generating the EOSSA file
	// (e.g. PRE, POST, BOTH, NONE).
	ValidCalibrations param.Opt[string] `json:"validCalibrations,omitzero"`
	// Array of SOI Calibrations associated with this SOIObservationSet.
	Calibrations []SoiObservationSetNewBulkParamsBodyCalibration `json:"calibrations,omitzero"`
	// OpticalSOIObservations associated with this SOIObservationSet.
	OpticalSoiObservationList []SoiObservationSetNewBulkParamsBodyOpticalSoiObservationList `json:"opticalSOIObservationList,omitzero"`
	// RadarSOIObservations associated with this RadarSOIObservationSet.
	RadarSoiObservationList []SoiObservationSetNewBulkParamsBodyRadarSoiObservationList `json:"radarSOIObservationList,omitzero"`
	// The reference frame of the observation measurements. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame string `json:"senReferenceFrame,omitzero"`
	// Array of the SpectralFilters keywords, must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	SpectralFilters []string `json:"spectralFilters,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SoiObservationSetNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SoiObservationSetNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SoiObservationSetNewBulkParamsBody](
		"type", "OPTICAL", "RADAR",
	)
	apijson.RegisterFieldValidator[SoiObservationSetNewBulkParamsBody](
		"referenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
	apijson.RegisterFieldValidator[SoiObservationSetNewBulkParamsBody](
		"senReferenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// Schema for SOI Calibration data.
type SoiObservationSetNewBulkParamsBodyCalibration struct {
	// Background intensity, at calibration, specified in kilowatts per steradian per
	// micrometer.
	CalBgIntensity param.Opt[float64] `json:"calBgIntensity,omitzero"`
	// Coefficient value for how much signal would be lost to atmospheric attenuation
	// for a star at zenith, in magnitudes per air mass.
	CalExtinctionCoeff param.Opt[float64] `json:"calExtinctionCoeff,omitzero"`
	// Maximum extinction coefficient uncertainty in magnitudes, at calibration (e.g.
	// -5.0 to 30.0).
	CalExtinctionCoeffMaxUnc param.Opt[float64] `json:"calExtinctionCoeffMaxUnc,omitzero"`
	// Extinction coefficient uncertainty in magnitudes, at calibration, which
	// represents the difference between the measured brightness and predicted
	// brightness of the star with the extinction removed, making it exo-atmospheric
	// (e.g. -5.0 to 30.0).
	CalExtinctionCoeffUnc param.Opt[float64] `json:"calExtinctionCoeffUnc,omitzero"`
	// Number of correlated stars in the FOV with the target object, at calibration.
	// Can be 0 for narrow FOV sensors.
	CalNumCorrelatedStars param.Opt[int64] `json:"calNumCorrelatedStars,omitzero"`
	// Number of detected stars in the FOV with the target object, at calibration.
	// Helps identify frames with clouds.
	CalNumDetectedStars param.Opt[int64] `json:"calNumDetectedStars,omitzero"`
	// Average Sky Background signals in magnitudes, at calibration. Sky Background
	// refers to the incoming light from an apparently empty part of the night sky.
	CalSkyBg param.Opt[float64] `json:"calSkyBg,omitzero"`
	// In-band solar magnitudes at 1 A.U, at calibration (e.g. -5.0 to 30.0).
	CalSpectralFilterSolarMag param.Opt[float64] `json:"calSpectralFilterSolarMag,omitzero"`
	// Start time of calibration in ISO 8601 UTC time, with millisecond precision.
	CalTime param.Opt[time.Time] `json:"calTime,omitzero" format:"date-time"`
	// Type of calibration (e.g. PRE, MID, POST).
	CalType param.Opt[string] `json:"calType,omitzero"`
	// Value representing the difference between the catalog magnitude and instrumental
	// magnitude for a set of standard stars, at calibration (e.g. -5.0 to 30.0).
	CalZeroPoint param.Opt[float64] `json:"calZeroPoint,omitzero"`
	paramObj
}

func (r SoiObservationSetNewBulkParamsBodyCalibration) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetNewBulkParamsBodyCalibration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetNewBulkParamsBodyCalibration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Optical SOI record contains observation information taken from a sensor about
// a Space Object.
//
// The property ObStartTime is required.
type SoiObservationSetNewBulkParamsBodyOpticalSoiObservationList struct {
	// Observation detection start time in ISO 8601 UTC with microsecond precision.
	ObStartTime time.Time `json:"obStartTime,required" format:"date-time"`
	// The reference number, x, where x ranges from 1 to n, where n is the number
	// specified in spectralFilters that corresponds to the spectral filter used.
	CurrentSpectralFilterNum param.Opt[int64] `json:"currentSpectralFilterNum,omitzero"`
	// Image exposure duration in seconds.
	ExpDuration param.Opt[float64] `json:"expDuration,omitzero"`
	// Array of declination values, in degrees, of the Target object from the frame of
	// reference of the sensor. A value is provided for each element in the intensities
	// field, at the middle of the frame’s exposure time.
	Declinations []float64 `json:"declinations,omitzero"`
	// Array of coefficients for how much signal would be lost to atmospheric
	// attenuation for a star at zenith for each element in intensities, in magnitudes
	// per air mass.
	ExtinctionCoeffs []float64 `json:"extinctionCoeffs,omitzero"`
	// Array of extinction coefficient uncertainties for each element in intensities.
	// Each value represents the difference between the measured brightness and
	// predicted brightness of the star with the extinction removed, making it
	// exo-atmospheric (e.g. -5.0 to 30.0).
	ExtinctionCoeffsUnc []float64 `json:"extinctionCoeffsUnc,omitzero"`
	// Array of intensities of the Space Object for observations, in kilowatts per
	// steradian per micrometer.
	Intensities []float64 `json:"intensities,omitzero"`
	// Array of start times for each intensity measurement. The 1st value in the array
	// will match obStartTime.
	IntensityTimes []time.Time `json:"intensityTimes,omitzero" format:"date-time"`
	// Array of local average Sky Background signals, in magnitudes, with a value
	// corresponding to the time of each intensity measurement. Sky Background refers
	// to the incoming light from an apparently empty part of the night sky.
	LocalSkyBgs []float64 `json:"localSkyBgs,omitzero"`
	// Array of uncertainty of the local average Sky Background signal, in magnitudes,
	// with a value corresponding to the time of each intensity measurement.
	LocalSkyBgsUnc []float64 `json:"localSkyBgsUnc,omitzero"`
	// Array of the number of correlated stars in the FOV with a value for each element
	// in the intensities field.
	NumCorrelatedStars []int64 `json:"numCorrelatedStars,omitzero"`
	// Array of the number of detected stars in the FOV with a value for each element
	// in the intensities field.
	NumDetectedStars []int64 `json:"numDetectedStars,omitzero"`
	// Array of values giving the percent of pixels that make up the object signal that
	// are beyond the saturation point for the sensor, with a value for each element in
	// the intensities field.
	PercentSats []float64 `json:"percentSats,omitzero"`
	// Array of right ascension rate values, in degrees per second, measuring the rate
	// the telescope is moving to track the Target object from the frame of reference
	// of the sensor, for each element in the intensities field, at the middle of the
	// frame’s exposure time.
	RaRates []float64 `json:"raRates,omitzero"`
	// Array of right ascension values, in degrees, of the Target object from the frame
	// of reference of the sensor. A value is provided for each element in the
	// intensities field.
	Ras []float64 `json:"ras,omitzero"`
	// Array of average Sky Background signals, in magnitudes, with a value
	// corresponding to the time of each intensity measurement. Sky Background refers
	// to the incoming light from an apparently empty part of the night sky.
	SkyBgs []float64 `json:"skyBgs,omitzero"`
	// Array of values for the zero-point in magnitudes, calculated at the time of each
	// intensity measurement. It is the difference between the catalog mag and
	// instrumental mag for a set of standard stars (e.g. -5.0 to 30.0).
	ZeroPoints []float64 `json:"zeroPoints,omitzero"`
	paramObj
}

func (r SoiObservationSetNewBulkParamsBodyOpticalSoiObservationList) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetNewBulkParamsBodyOpticalSoiObservationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetNewBulkParamsBodyOpticalSoiObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Radar SOI record contains observation information taken from a sensor about a
// Space Object.
//
// The property ObStartTime is required.
type SoiObservationSetNewBulkParamsBodyRadarSoiObservationList struct {
	// Observation detection start time in ISO 8601 UTC format with microsecond
	// precision.
	ObStartTime time.Time `json:"obStartTime,required" format:"date-time"`
	// Beta angle (between target and radar-image frame z axis) in degrees.
	Beta param.Opt[float64] `json:"beta,omitzero"`
	// Radar center frequency of the radar in hertz.
	CenterFrequency param.Opt[float64] `json:"centerFrequency,omitzero"`
	// Optional id of assumed AttitudeSet of object being observed.
	IDAttitudeSet param.Opt[string] `json:"idAttitudeSet,omitzero"`
	// Optional id of assumed StateVector of object being observed.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Kappa angle (between radar-line-of-sight and target-frame x axis) in degrees.
	Kappa param.Opt[float64] `json:"kappa,omitzero"`
	// Bandwidth of radar pulse in hertz.
	PulseBandwidth param.Opt[float64] `json:"pulseBandwidth,omitzero"`
	// Array of the aspect angle at the center of the image in degrees. The 'tovs' and
	// 'aspectAngles' arrays must match in size, if 'aspectAngles' is provided.
	AspectAngles []float64 `json:"aspectAngles,omitzero"`
	// Array of sensor azimuth angle biases in degrees. The 'tovs' and 'azimuthBiases'
	// arrays must match in size, if 'azimuthBiases' is provided.
	AzimuthBiases []float64 `json:"azimuthBiases,omitzero"`
	// Array of the azimuth rate of target at image center in degrees per second. The
	// 'tovs' and 'azimuthRates' arrays must match in size, if 'azimuthRates' is
	// provided. If there is an associated image the azimuth rate is assumed to be at
	// image center.
	AzimuthRates []float64 `json:"azimuthRates,omitzero"`
	// Array of the azimuth angle to target at image center in degrees. The 'tovs' and
	// 'azimuths' arrays must match in size, if 'azimuths' is provided. If there is an
	// associated image the azimuth angle is assumed to be at image center.
	Azimuths []float64 `json:"azimuths,omitzero"`
	// Array of cross-range resolutions (accounting for weighting function) in
	// kilometers. The 'tovs' and 'crossRangeRes' arrays must match in size, if
	// 'crossRangeRes' is provided.
	CrossRangeRes []float64 `json:"crossRangeRes,omitzero"`
	// Array of average Interpulse spacing in seconds. The 'tovs' and 'deltaTimes'
	// arrays must match in size, if 'deltaTimes' is provided.
	DeltaTimes []float64 `json:"deltaTimes,omitzero"`
	// Array of conversion factors between Doppler in hertz and cross-range in meters.
	// The 'tovs' and 'doppler2XRs' arrays must match in size, if 'doppler2XRs' is
	// provided.
	Doppler2XRs []float64 `json:"doppler2XRs,omitzero"`
	// Array of sensor elevation biases in degrees. The 'tovs' and 'elevationBiases'
	// arrays must match in size, if 'elevationBiases' is provided.
	ElevationBiases []float64 `json:"elevationBiases,omitzero"`
	// Array of the elevation rate of target at image center in degrees per second. The
	// 'tovs' and 'elevationRates' arrays must match in size, if 'elevationRates' is
	// provided. If there is an associated image the elevation rate is assumed to be at
	// image center.
	ElevationRates []float64 `json:"elevationRates,omitzero"`
	// Array of the elevation angle to target at image center in degrees. The 'tovs'
	// and 'elevations' arrays must match in size, if 'elevations' is provided. If
	// there is an associated image the elevation angle is assumed to be at image
	// center.
	Elevations []float64 `json:"elevations,omitzero"`
	// Array of Integration angles in degrees. The 'tovs' and 'integrationAngles'
	// arrays must match in size, if 'integrationAngles' is provided.
	IntegrationAngles []float64 `json:"integrationAngles,omitzero"`
	// Array of the peak pixel amplitude for each image in decibels. The 'tovs' and
	// 'peakAmplitudes' arrays must match in size, if 'peakAmplitudes' is provided.
	PeakAmplitudes []float64 `json:"peakAmplitudes,omitzero"`
	// Array of sensor polarizations when collecting the data. Polarization is a
	// property of electromagnetic waves that describes the orientation of their
	// oscillations. Possible values are H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	Polarizations []string `json:"polarizations,omitzero"`
	// Array of the component of target angular velocity observable by radar in radians
	// per second. The 'tovs' and 'projAngVels' arrays must match in size, if
	// 'projAngVels' is provided.
	ProjAngVels []float64 `json:"projAngVels,omitzero"`
	// Array of the range accelerations of target in kilometers per second squared. The
	// 'tovs' and 'rangeAccels' arrays must match in size, if 'rangeAccels' is
	// provided. If there is an associated image the range acceleration is assumed to
	// be at image center.
	RangeAccels []float64 `json:"rangeAccels,omitzero"`
	// Array of sensor range biases in kilometers. The 'tovs' and 'rangeBiases' arrays
	// must match in size, if 'rangeBiases' is provided.
	RangeBiases []float64 `json:"rangeBiases,omitzero"`
	// Array of the range rate of target at image center in kilometers per second. The
	// 'tovs' and 'rangeRates' arrays must match in size, if 'rangeRates' is provided.
	// If there is an associated image the range rate is assumed to be at image center.
	RangeRates []float64 `json:"rangeRates,omitzero"`
	// Array of the range to target at image center in kilometers. The 'tovs' and
	// 'ranges' arrays must match in size, if 'ranges' is provided. If there is an
	// associated image the range is assumed to be at image center.
	Ranges []float64 `json:"ranges,omitzero"`
	// Array of error estimates of RCS values, in square meters.
	RcsErrorEsts []float64 `json:"rcsErrorEsts,omitzero"`
	// Array of observed radar cross section (RCS) values, in square meters.
	RcsValues []float64 `json:"rcsValues,omitzero"`
	// Array of range sample spacing in meters. The 'tovs' and 'rspaces' arrays must
	// match in size, if 'rspaces' is provided.
	Rspaces []float64 `json:"rspaces,omitzero"`
	// Array of spectral widths, in hertz. The spectral width of a satellite can help
	// determine if an object is stable or tumbling which is often useful to
	// distinguish a rocket body from an active stabilized payload or to deduce a
	// rotational period of slowly tumbling objects at GEO.
	SpectralWidths []float64 `json:"spectralWidths,omitzero"`
	// Array of the times of validity in ISO 8601 UTC format with microsecond
	// precision.
	Tovs []time.Time `json:"tovs,omitzero" format:"date-time"`
	// Array of the cartesian X accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xaccel' arrays must match in size, if 'xaccel' is provided.
	Xaccel []float64 `json:"xaccel,omitzero"`
	// Array of the cartesian X positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xpos' arrays must match in size, if 'xpos' is provided.
	Xpos []float64 `json:"xpos,omitzero"`
	// Array of cross-range sample spacing in meters. The 'tovs' and 'xspaces' arrays
	// must match in size, if 'xspaces' is provided.
	Xspaces []float64 `json:"xspaces,omitzero"`
	// Array of the cartesian X velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xvel' arrays must match in size, if 'xvel' is provided.
	Xvel []float64 `json:"xvel,omitzero"`
	// Array of the cartesian Y accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'yaccel' arrays must match in size, if 'yaccel' is provided.
	Yaccel []float64 `json:"yaccel,omitzero"`
	// Array of the cartesian Y positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'ypos' arrays must match in size, if 'ypos' is provided.
	Ypos []float64 `json:"ypos,omitzero"`
	// Array of the cartesian Y velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'yvel' arrays must match in size, if 'yvel' is provided.
	Yvel []float64 `json:"yvel,omitzero"`
	// Array of the cartesian Z accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zaccel' arrays must match in size, if 'zaccel' is provided.
	Zaccel []float64 `json:"zaccel,omitzero"`
	// Array of the cartesian Z positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zpos' arrays must match in size, if 'zpos' is provided.
	Zpos []float64 `json:"zpos,omitzero"`
	// Array of the cartesian Z velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zvel' arrays must match in size, if 'zvel' is provided.
	Zvel []float64 `json:"zvel,omitzero"`
	paramObj
}

func (r SoiObservationSetNewBulkParamsBodyRadarSoiObservationList) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetNewBulkParamsBodyRadarSoiObservationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetNewBulkParamsBodyRadarSoiObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SoiObservationSetGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SoiObservationSetGetParams]'s query parameters as
// `url.Values`.
func (r SoiObservationSetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SoiObservationSetTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SoiObservationSetTupleParams]'s query parameters as
// `url.Values`.
func (r SoiObservationSetTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SoiObservationSetUnvalidatedPublishParams struct {
	Body []SoiObservationSetUnvalidatedPublishParamsBody
	paramObj
}

func (r SoiObservationSetUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SoiObservationSetUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide operations for posting space object idenfification
// observation sets.
//
// The properties ClassificationMarking, DataMode, NumObs, Source, StartTime, Type
// are required.
type SoiObservationSetUnvalidatedPublishParamsBody struct {
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
	// The number of observation records in the set.
	NumObs int64 `json:"numObs,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Observation set detection start time in ISO 8601 UTC with microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Observation type (OPTICAL, RADAR).
	//
	// Any of "OPTICAL", "RADAR".
	Type string `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The number of pixels binned horizontally.
	BinningHoriz param.Opt[int64] `json:"binningHoriz,omitzero"`
	// The number of pixels binned vertically.
	BinningVert param.Opt[int64] `json:"binningVert,omitzero"`
	// Boolean indicating if a brightness variance change event was detected, based on
	// historical collection data for the object.
	BrightnessVarianceChangeDetected param.Opt[bool] `json:"brightnessVarianceChangeDetected,omitzero"`
	// Type of calibration used by the Sensor (e.g. ALL SKY, DIFFERENTIAL, DEFAULT,
	// NONE).
	CalibrationType param.Opt[string] `json:"calibrationType,omitzero"`
	// Overall qualitative confidence assessment of change detection results (e.g.
	// HIGH, MEDIUM, LOW).
	ChangeConf param.Opt[string] `json:"changeConf,omitzero"`
	// Boolean indicating if any change event was detected, based on historical
	// collection data for the object.
	ChangeDetected param.Opt[bool] `json:"changeDetected,omitzero"`
	// Qualitative Collection Density assessment, with respect to confidence of
	// detecting a change event (e.g. HIGH, MEDIUM, LOW).
	CollectionDensityConf param.Opt[string] `json:"collectionDensityConf,omitzero"`
	// Universally Unique collection ID. Mechanism to correlate Single Point Photometry
	// (SPP) JSON files to images.
	CollectionID param.Opt[string] `json:"collectionId,omitzero"`
	// Mode indicating telescope movement during collection (AUTOTRACK, MANUAL
	// AUTOTRACK, MANUAL RATE TRACK, MANUAL SIDEREAL, SIDEREAL, RATE TRACK).
	CollectionMode param.Opt[string] `json:"collectionMode,omitzero"`
	// Object Correlation Quality value. Measures how close the observed object's orbit
	// is to matching an object in the catalog. The scale of this field may vary
	// depending on provider. Users should consult the data provider to verify the
	// meaning of the value (e.g. A value of 0.0 indicates a high/strong correlation,
	// while a value closer to 1.0 indicates low/weak correlation).
	CorrQuality param.Opt[float64] `json:"corrQuality,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Observation set detection end time in ISO 8601 UTC with microsecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The gain used during the collection, in units of photoelectrons per
	// analog-to-digital unit (e-/ADU). If no gain is used, the value = 1.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// ID of the UDL Elset of the Space Object under observation.
	IDElset param.Opt[string] `json:"idElset,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// ID of the observing sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Line of sight declination at observation set detection end time. Specified in
	// degrees, in the specified referenceFrame. If referenceFrame is null then J2K
	// should be assumed (e.g -30 to 130.0).
	LosDeclinationEnd param.Opt[float64] `json:"losDeclinationEnd,omitzero"`
	// Line of sight declination at observation set detection start time. Specified in
	// degrees, in the specified referenceFrame. If referenceFrame is null then J2K
	// should be assumed (e.g -30 to 130.0).
	LosDeclinationStart param.Opt[float64] `json:"losDeclinationStart,omitzero"`
	// SOI msgCreateDate time in ISO 8601 UTC time, with millisecond precision.
	MsgCreateDate param.Opt[time.Time] `json:"msgCreateDate,omitzero" format:"date-time"`
	// The value is the number of spectral filters used.
	NumSpectralFilters param.Opt[int64] `json:"numSpectralFilters,omitzero"`
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
	// Optional identifier provided by the record source to indicate the sensor
	// identifier to which this attitude set applies if this set is reporting a single
	// sensor orientation. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// A threshold for percent of pixels that make up object signal that are beyond the
	// saturation point for the sensor that are removed in the EOSSA file, in range of
	// 0 to 1.
	PercentSatThreshold param.Opt[float64] `json:"percentSatThreshold,omitzero"`
	// Boolean indicating if a periodicity change event was detected, based on
	// historical collection data for the object.
	PeriodicityChangeDetected param.Opt[bool] `json:"periodicityChangeDetected,omitzero"`
	// Qualitative assessment of the periodicity detection results from the Attitude
	// and Shape Retrieval (ASR) Periodicity Assessment (PA) Tool (e.g. HIGH, MEDIUM,
	// LOW).
	PeriodicityDetectionConf param.Opt[string] `json:"periodicityDetectionConf,omitzero"`
	// Qualitative Periodicity Sampling assessment, with respect to confidence of
	// detecting a change event (e.g. HIGH, MEDIUM, LOW).
	PeriodicitySamplingConf param.Opt[string] `json:"periodicitySamplingConf,omitzero"`
	// Pixel array size (height) in pixels.
	PixelArrayHeight param.Opt[int64] `json:"pixelArrayHeight,omitzero"`
	// Pixel array size (width) in pixels.
	PixelArrayWidth param.Opt[int64] `json:"pixelArrayWidth,omitzero"`
	// The maximum valid pixel value.
	PixelMax param.Opt[int64] `json:"pixelMax,omitzero"`
	// The minimum valid pixel value.
	PixelMin param.Opt[int64] `json:"pixelMin,omitzero"`
	// Pointing angle of the Azimuth gimbal/mount at observation set detection end
	// time. Specified in degrees.
	PointingAngleAzEnd param.Opt[float64] `json:"pointingAngleAzEnd,omitzero"`
	// Pointing angle of the Azimuth gimbal/mount at observation set detection start
	// time. Specified in degrees.
	PointingAngleAzStart param.Opt[float64] `json:"pointingAngleAzStart,omitzero"`
	// Pointing angle of the Elevation gimbal/mount at observation set detection end
	// time. Specified in degrees.
	PointingAngleElEnd param.Opt[float64] `json:"pointingAngleElEnd,omitzero"`
	// Pointing angle of the Elevation gimbal/mount at observation set detection start
	// time. Specified in degrees.
	PointingAngleElStart param.Opt[float64] `json:"pointingAngleElStart,omitzero"`
	// Polar angle of the gimbal/mount at observation set detection end time in
	// degrees.
	PolarAngleEnd param.Opt[float64] `json:"polarAngleEnd,omitzero"`
	// Polar angle of the gimbal/mount at observation set detection start time in
	// degrees.
	PolarAngleStart param.Opt[float64] `json:"polarAngleStart,omitzero"`
	// Name of the target satellite.
	SatelliteName param.Opt[string] `json:"satelliteName,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude at startTime (if mobile/onorbit) in kilometers.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at startTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info. -90 to 90 degrees (negative values south of
	// equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at startTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values south of
	// equator).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// ID of the AttitudeSet record for the observing sensor.
	SensorAsID param.Opt[string] `json:"sensorAsId,omitzero"`
	// Cartesian X velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvelx param.Opt[float64] `json:"senvelx,omitzero"`
	// Cartesian Y velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvely param.Opt[float64] `json:"senvely,omitzero"`
	// Cartesian Z velocity of the observing mobile/onorbit sensor at startTime, in
	// kilometers per second, in the specified senReferenceFrame. If senReferenceFrame
	// is null then J2K should be assumed.
	Senvelz param.Opt[float64] `json:"senvelz,omitzero"`
	// Cartesian X position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Senx param.Opt[float64] `json:"senx,omitzero"`
	// Cartesian Y position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Seny param.Opt[float64] `json:"seny,omitzero"`
	// Cartesian Z position of the observing mobile/onorbit sensor at startTime, in
	// kilometers, in the specified senReferenceFrame. If senReferenceFrame is null
	// then J2K should be assumed.
	Senz param.Opt[float64] `json:"senz,omitzero"`
	// Software Version used to Capture, Process, and Deliver the data.
	SoftwareVersion param.Opt[string] `json:"softwareVersion,omitzero"`
	// The in-band solar magnitude at 1 A.U.
	SolarMag param.Opt[float64] `json:"solarMag,omitzero"`
	// Boolean indicating if a solar phase angle brightness change event was detected,
	// based on historical collection data for the object.
	SolarPhaseAngleBrightnessChangeDetected param.Opt[bool] `json:"solarPhaseAngleBrightnessChangeDetected,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Name of the Star Catalog used for photometry and astrometry.
	StarCatName param.Opt[string] `json:"starCatName,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating whether the target object was unable to be correlated to a
	// known object. This flag should only be set to true by data providers after an
	// attempt to correlate to an OnOrbit object was made and failed. If unable to
	// correlate, the 'origObjectId' field may be populated with an internal data
	// provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Key to indicate which, if any of, the pre/post photometer calibrations are valid
	// for use when generating data for the EOSSA file. If the field is not populated,
	// then the provided calibration data will be used when generating the EOSSA file
	// (e.g. PRE, POST, BOTH, NONE).
	ValidCalibrations param.Opt[string] `json:"validCalibrations,omitzero"`
	// Array of SOI Calibrations associated with this SOIObservationSet.
	Calibrations []SoiObservationSetUnvalidatedPublishParamsBodyCalibration `json:"calibrations,omitzero"`
	// OpticalSOIObservations associated with this SOIObservationSet.
	OpticalSoiObservationList []SoiObservationSetUnvalidatedPublishParamsBodyOpticalSoiObservationList `json:"opticalSOIObservationList,omitzero"`
	// RadarSOIObservations associated with this RadarSOIObservationSet.
	RadarSoiObservationList []SoiObservationSetUnvalidatedPublishParamsBodyRadarSoiObservationList `json:"radarSOIObservationList,omitzero"`
	// The reference frame of the observation measurements. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame string `json:"senReferenceFrame,omitzero"`
	// Array of the SpectralFilters keywords, must be present for all values n=1 to
	// numSpectralFilters, in incrementing order of n, and for no other values of n.
	SpectralFilters []string `json:"spectralFilters,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SoiObservationSetUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SoiObservationSetUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SoiObservationSetUnvalidatedPublishParamsBody](
		"type", "OPTICAL", "RADAR",
	)
	apijson.RegisterFieldValidator[SoiObservationSetUnvalidatedPublishParamsBody](
		"referenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
	apijson.RegisterFieldValidator[SoiObservationSetUnvalidatedPublishParamsBody](
		"senReferenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// Schema for SOI Calibration data.
type SoiObservationSetUnvalidatedPublishParamsBodyCalibration struct {
	// Background intensity, at calibration, specified in kilowatts per steradian per
	// micrometer.
	CalBgIntensity param.Opt[float64] `json:"calBgIntensity,omitzero"`
	// Coefficient value for how much signal would be lost to atmospheric attenuation
	// for a star at zenith, in magnitudes per air mass.
	CalExtinctionCoeff param.Opt[float64] `json:"calExtinctionCoeff,omitzero"`
	// Maximum extinction coefficient uncertainty in magnitudes, at calibration (e.g.
	// -5.0 to 30.0).
	CalExtinctionCoeffMaxUnc param.Opt[float64] `json:"calExtinctionCoeffMaxUnc,omitzero"`
	// Extinction coefficient uncertainty in magnitudes, at calibration, which
	// represents the difference between the measured brightness and predicted
	// brightness of the star with the extinction removed, making it exo-atmospheric
	// (e.g. -5.0 to 30.0).
	CalExtinctionCoeffUnc param.Opt[float64] `json:"calExtinctionCoeffUnc,omitzero"`
	// Number of correlated stars in the FOV with the target object, at calibration.
	// Can be 0 for narrow FOV sensors.
	CalNumCorrelatedStars param.Opt[int64] `json:"calNumCorrelatedStars,omitzero"`
	// Number of detected stars in the FOV with the target object, at calibration.
	// Helps identify frames with clouds.
	CalNumDetectedStars param.Opt[int64] `json:"calNumDetectedStars,omitzero"`
	// Average Sky Background signals in magnitudes, at calibration. Sky Background
	// refers to the incoming light from an apparently empty part of the night sky.
	CalSkyBg param.Opt[float64] `json:"calSkyBg,omitzero"`
	// In-band solar magnitudes at 1 A.U, at calibration (e.g. -5.0 to 30.0).
	CalSpectralFilterSolarMag param.Opt[float64] `json:"calSpectralFilterSolarMag,omitzero"`
	// Start time of calibration in ISO 8601 UTC time, with millisecond precision.
	CalTime param.Opt[time.Time] `json:"calTime,omitzero" format:"date-time"`
	// Type of calibration (e.g. PRE, MID, POST).
	CalType param.Opt[string] `json:"calType,omitzero"`
	// Value representing the difference between the catalog magnitude and instrumental
	// magnitude for a set of standard stars, at calibration (e.g. -5.0 to 30.0).
	CalZeroPoint param.Opt[float64] `json:"calZeroPoint,omitzero"`
	paramObj
}

func (r SoiObservationSetUnvalidatedPublishParamsBodyCalibration) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetUnvalidatedPublishParamsBodyCalibration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetUnvalidatedPublishParamsBodyCalibration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Optical SOI record contains observation information taken from a sensor about
// a Space Object.
//
// The property ObStartTime is required.
type SoiObservationSetUnvalidatedPublishParamsBodyOpticalSoiObservationList struct {
	// Observation detection start time in ISO 8601 UTC with microsecond precision.
	ObStartTime time.Time `json:"obStartTime,required" format:"date-time"`
	// The reference number, x, where x ranges from 1 to n, where n is the number
	// specified in spectralFilters that corresponds to the spectral filter used.
	CurrentSpectralFilterNum param.Opt[int64] `json:"currentSpectralFilterNum,omitzero"`
	// Image exposure duration in seconds.
	ExpDuration param.Opt[float64] `json:"expDuration,omitzero"`
	// Array of declination values, in degrees, of the Target object from the frame of
	// reference of the sensor. A value is provided for each element in the intensities
	// field, at the middle of the frame’s exposure time.
	Declinations []float64 `json:"declinations,omitzero"`
	// Array of coefficients for how much signal would be lost to atmospheric
	// attenuation for a star at zenith for each element in intensities, in magnitudes
	// per air mass.
	ExtinctionCoeffs []float64 `json:"extinctionCoeffs,omitzero"`
	// Array of extinction coefficient uncertainties for each element in intensities.
	// Each value represents the difference between the measured brightness and
	// predicted brightness of the star with the extinction removed, making it
	// exo-atmospheric (e.g. -5.0 to 30.0).
	ExtinctionCoeffsUnc []float64 `json:"extinctionCoeffsUnc,omitzero"`
	// Array of intensities of the Space Object for observations, in kilowatts per
	// steradian per micrometer.
	Intensities []float64 `json:"intensities,omitzero"`
	// Array of start times for each intensity measurement. The 1st value in the array
	// will match obStartTime.
	IntensityTimes []time.Time `json:"intensityTimes,omitzero" format:"date-time"`
	// Array of local average Sky Background signals, in magnitudes, with a value
	// corresponding to the time of each intensity measurement. Sky Background refers
	// to the incoming light from an apparently empty part of the night sky.
	LocalSkyBgs []float64 `json:"localSkyBgs,omitzero"`
	// Array of uncertainty of the local average Sky Background signal, in magnitudes,
	// with a value corresponding to the time of each intensity measurement.
	LocalSkyBgsUnc []float64 `json:"localSkyBgsUnc,omitzero"`
	// Array of the number of correlated stars in the FOV with a value for each element
	// in the intensities field.
	NumCorrelatedStars []int64 `json:"numCorrelatedStars,omitzero"`
	// Array of the number of detected stars in the FOV with a value for each element
	// in the intensities field.
	NumDetectedStars []int64 `json:"numDetectedStars,omitzero"`
	// Array of values giving the percent of pixels that make up the object signal that
	// are beyond the saturation point for the sensor, with a value for each element in
	// the intensities field.
	PercentSats []float64 `json:"percentSats,omitzero"`
	// Array of right ascension rate values, in degrees per second, measuring the rate
	// the telescope is moving to track the Target object from the frame of reference
	// of the sensor, for each element in the intensities field, at the middle of the
	// frame’s exposure time.
	RaRates []float64 `json:"raRates,omitzero"`
	// Array of right ascension values, in degrees, of the Target object from the frame
	// of reference of the sensor. A value is provided for each element in the
	// intensities field.
	Ras []float64 `json:"ras,omitzero"`
	// Array of average Sky Background signals, in magnitudes, with a value
	// corresponding to the time of each intensity measurement. Sky Background refers
	// to the incoming light from an apparently empty part of the night sky.
	SkyBgs []float64 `json:"skyBgs,omitzero"`
	// Array of values for the zero-point in magnitudes, calculated at the time of each
	// intensity measurement. It is the difference between the catalog mag and
	// instrumental mag for a set of standard stars (e.g. -5.0 to 30.0).
	ZeroPoints []float64 `json:"zeroPoints,omitzero"`
	paramObj
}

func (r SoiObservationSetUnvalidatedPublishParamsBodyOpticalSoiObservationList) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetUnvalidatedPublishParamsBodyOpticalSoiObservationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetUnvalidatedPublishParamsBodyOpticalSoiObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Radar SOI record contains observation information taken from a sensor about a
// Space Object.
//
// The property ObStartTime is required.
type SoiObservationSetUnvalidatedPublishParamsBodyRadarSoiObservationList struct {
	// Observation detection start time in ISO 8601 UTC format with microsecond
	// precision.
	ObStartTime time.Time `json:"obStartTime,required" format:"date-time"`
	// Beta angle (between target and radar-image frame z axis) in degrees.
	Beta param.Opt[float64] `json:"beta,omitzero"`
	// Radar center frequency of the radar in hertz.
	CenterFrequency param.Opt[float64] `json:"centerFrequency,omitzero"`
	// Optional id of assumed AttitudeSet of object being observed.
	IDAttitudeSet param.Opt[string] `json:"idAttitudeSet,omitzero"`
	// Optional id of assumed StateVector of object being observed.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Kappa angle (between radar-line-of-sight and target-frame x axis) in degrees.
	Kappa param.Opt[float64] `json:"kappa,omitzero"`
	// Bandwidth of radar pulse in hertz.
	PulseBandwidth param.Opt[float64] `json:"pulseBandwidth,omitzero"`
	// Array of the aspect angle at the center of the image in degrees. The 'tovs' and
	// 'aspectAngles' arrays must match in size, if 'aspectAngles' is provided.
	AspectAngles []float64 `json:"aspectAngles,omitzero"`
	// Array of sensor azimuth angle biases in degrees. The 'tovs' and 'azimuthBiases'
	// arrays must match in size, if 'azimuthBiases' is provided.
	AzimuthBiases []float64 `json:"azimuthBiases,omitzero"`
	// Array of the azimuth rate of target at image center in degrees per second. The
	// 'tovs' and 'azimuthRates' arrays must match in size, if 'azimuthRates' is
	// provided. If there is an associated image the azimuth rate is assumed to be at
	// image center.
	AzimuthRates []float64 `json:"azimuthRates,omitzero"`
	// Array of the azimuth angle to target at image center in degrees. The 'tovs' and
	// 'azimuths' arrays must match in size, if 'azimuths' is provided. If there is an
	// associated image the azimuth angle is assumed to be at image center.
	Azimuths []float64 `json:"azimuths,omitzero"`
	// Array of cross-range resolutions (accounting for weighting function) in
	// kilometers. The 'tovs' and 'crossRangeRes' arrays must match in size, if
	// 'crossRangeRes' is provided.
	CrossRangeRes []float64 `json:"crossRangeRes,omitzero"`
	// Array of average Interpulse spacing in seconds. The 'tovs' and 'deltaTimes'
	// arrays must match in size, if 'deltaTimes' is provided.
	DeltaTimes []float64 `json:"deltaTimes,omitzero"`
	// Array of conversion factors between Doppler in hertz and cross-range in meters.
	// The 'tovs' and 'doppler2XRs' arrays must match in size, if 'doppler2XRs' is
	// provided.
	Doppler2XRs []float64 `json:"doppler2XRs,omitzero"`
	// Array of sensor elevation biases in degrees. The 'tovs' and 'elevationBiases'
	// arrays must match in size, if 'elevationBiases' is provided.
	ElevationBiases []float64 `json:"elevationBiases,omitzero"`
	// Array of the elevation rate of target at image center in degrees per second. The
	// 'tovs' and 'elevationRates' arrays must match in size, if 'elevationRates' is
	// provided. If there is an associated image the elevation rate is assumed to be at
	// image center.
	ElevationRates []float64 `json:"elevationRates,omitzero"`
	// Array of the elevation angle to target at image center in degrees. The 'tovs'
	// and 'elevations' arrays must match in size, if 'elevations' is provided. If
	// there is an associated image the elevation angle is assumed to be at image
	// center.
	Elevations []float64 `json:"elevations,omitzero"`
	// Array of Integration angles in degrees. The 'tovs' and 'integrationAngles'
	// arrays must match in size, if 'integrationAngles' is provided.
	IntegrationAngles []float64 `json:"integrationAngles,omitzero"`
	// Array of the peak pixel amplitude for each image in decibels. The 'tovs' and
	// 'peakAmplitudes' arrays must match in size, if 'peakAmplitudes' is provided.
	PeakAmplitudes []float64 `json:"peakAmplitudes,omitzero"`
	// Array of sensor polarizations when collecting the data. Polarization is a
	// property of electromagnetic waves that describes the orientation of their
	// oscillations. Possible values are H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	Polarizations []string `json:"polarizations,omitzero"`
	// Array of the component of target angular velocity observable by radar in radians
	// per second. The 'tovs' and 'projAngVels' arrays must match in size, if
	// 'projAngVels' is provided.
	ProjAngVels []float64 `json:"projAngVels,omitzero"`
	// Array of the range accelerations of target in kilometers per second squared. The
	// 'tovs' and 'rangeAccels' arrays must match in size, if 'rangeAccels' is
	// provided. If there is an associated image the range acceleration is assumed to
	// be at image center.
	RangeAccels []float64 `json:"rangeAccels,omitzero"`
	// Array of sensor range biases in kilometers. The 'tovs' and 'rangeBiases' arrays
	// must match in size, if 'rangeBiases' is provided.
	RangeBiases []float64 `json:"rangeBiases,omitzero"`
	// Array of the range rate of target at image center in kilometers per second. The
	// 'tovs' and 'rangeRates' arrays must match in size, if 'rangeRates' is provided.
	// If there is an associated image the range rate is assumed to be at image center.
	RangeRates []float64 `json:"rangeRates,omitzero"`
	// Array of the range to target at image center in kilometers. The 'tovs' and
	// 'ranges' arrays must match in size, if 'ranges' is provided. If there is an
	// associated image the range is assumed to be at image center.
	Ranges []float64 `json:"ranges,omitzero"`
	// Array of error estimates of RCS values, in square meters.
	RcsErrorEsts []float64 `json:"rcsErrorEsts,omitzero"`
	// Array of observed radar cross section (RCS) values, in square meters.
	RcsValues []float64 `json:"rcsValues,omitzero"`
	// Array of range sample spacing in meters. The 'tovs' and 'rspaces' arrays must
	// match in size, if 'rspaces' is provided.
	Rspaces []float64 `json:"rspaces,omitzero"`
	// Array of spectral widths, in hertz. The spectral width of a satellite can help
	// determine if an object is stable or tumbling which is often useful to
	// distinguish a rocket body from an active stabilized payload or to deduce a
	// rotational period of slowly tumbling objects at GEO.
	SpectralWidths []float64 `json:"spectralWidths,omitzero"`
	// Array of the times of validity in ISO 8601 UTC format with microsecond
	// precision.
	Tovs []time.Time `json:"tovs,omitzero" format:"date-time"`
	// Array of the cartesian X accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xaccel' arrays must match in size, if 'xaccel' is provided.
	Xaccel []float64 `json:"xaccel,omitzero"`
	// Array of the cartesian X positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xpos' arrays must match in size, if 'xpos' is provided.
	Xpos []float64 `json:"xpos,omitzero"`
	// Array of cross-range sample spacing in meters. The 'tovs' and 'xspaces' arrays
	// must match in size, if 'xspaces' is provided.
	Xspaces []float64 `json:"xspaces,omitzero"`
	// Array of the cartesian X velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'xvel' arrays must match in size, if 'xvel' is provided.
	Xvel []float64 `json:"xvel,omitzero"`
	// Array of the cartesian Y accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'yaccel' arrays must match in size, if 'yaccel' is provided.
	Yaccel []float64 `json:"yaccel,omitzero"`
	// Array of the cartesian Y positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'ypos' arrays must match in size, if 'ypos' is provided.
	Ypos []float64 `json:"ypos,omitzero"`
	// Array of the cartesian Y velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'yvel' arrays must match in size, if 'yvel' is provided.
	Yvel []float64 `json:"yvel,omitzero"`
	// Array of the cartesian Z accelerations, in kilometers per second squared, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zaccel' arrays must match in size, if 'zaccel' is provided.
	Zaccel []float64 `json:"zaccel,omitzero"`
	// Array of the cartesian Z positions of the target, in kilometers, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zpos' arrays must match in size, if 'zpos' is provided.
	Zpos []float64 `json:"zpos,omitzero"`
	// Array of the cartesian Z velocities of target, in kilometers per second, in the
	// specified referenceFrame. If referenceFrame is null then J2K should be assumed.
	// The 'tovs' and 'zvel' arrays must match in size, if 'zvel' is provided.
	Zvel []float64 `json:"zvel,omitzero"`
	paramObj
}

func (r SoiObservationSetUnvalidatedPublishParamsBodyRadarSoiObservationList) MarshalJSON() (data []byte, err error) {
	type shadow SoiObservationSetUnvalidatedPublishParamsBodyRadarSoiObservationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoiObservationSetUnvalidatedPublishParamsBodyRadarSoiObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
