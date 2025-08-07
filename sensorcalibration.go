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
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// SensorCalibrationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSensorCalibrationService] method instead.
type SensorCalibrationService struct {
	Options []option.RequestOption
	History SensorCalibrationHistoryService
}

// NewSensorCalibrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSensorCalibrationService(opts ...option.RequestOption) (r SensorCalibrationService) {
	r = SensorCalibrationService{}
	r.Options = opts
	r.History = NewSensorCalibrationHistoryService(opts...)
	return
}

// Service operation to take a single SensorCalibration as a POST body and ingest
// into the database. This operation is not intended to be used for automated feeds
// into UDL. Data providers should contact the UDL team for specific role
// assignments and for instructions on setting up a permanent feed through an
// alternate mechanism.
func (r *SensorCalibrationService) New(ctx context.Context, body SensorCalibrationNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensorcalibration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single SensorCalibration by its unique ID passed as a
// path parameter.
func (r *SensorCalibrationService) Get(ctx context.Context, id string, query SensorCalibrationGetParams, opts ...option.RequestOption) (res *SensorCalibrationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensorcalibration/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SensorCalibrationService) Count(ctx context.Context, query SensorCalibrationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sensorcalibration/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// SensorCalibrations as a POST body and ingest into the database. This operation
// is not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *SensorCalibrationService) NewBulk(ctx context.Context, body SensorCalibrationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensorcalibration/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensorCalibrationService) Query(ctx context.Context, query SensorCalibrationQueryParams, opts ...option.RequestOption) (res *[]SensorCalibrationQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensorcalibration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SensorCalibrationService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *SensorCalibrationQueryHelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensorcalibration/queryhelp"
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
func (r *SensorCalibrationService) Tuple(ctx context.Context, query SensorCalibrationTupleParams, opts ...option.RequestOption) (res *[]SensorCalibrationTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensorcalibration/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple sensorcalibration records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SensorCalibrationService) UnvalidatedPublish(ctx context.Context, body SensorCalibrationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-sensorcalibration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The Sensor Calibration service records data about a sensor's overall accuracy
// and is used to adjust sensor settings to achieve and maintain that accuracy in
// reported sensor observations. Calibration occurs periodically when needed to
// maintain sensor accuracy or on-demand to adjust a sensor for a specific reading.
type SensorCalibrationGetResponse struct {
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
	DataMode SensorCalibrationGetResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the sensor to which this calibration data applies. This ID
	// can be used to obtain additional information on a sensor using the 'get by ID'
	// operation (e.g. /udl/sensor/{id}). For example, the sensor with idSensor = abc
	// would be queried as /udl/sensor/abc.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Sensor azimuth/right-ascension acceleration bias, in degrees per second squared.
	AzRaAccelBias float64 `json:"azRaAccelBias"`
	// The standard deviation of the azimuth/right ascension acceleration residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension acceleration bias.
	AzRaAccelSigma float64 `json:"azRaAccelSigma"`
	// Sensor azimuth/right-ascension bias, in degrees.
	AzRaBias float64 `json:"azRaBias"`
	// Sensor azimuth/right-ascension rate bias, in degrees per second.
	AzRaRateBias float64 `json:"azRaRateBias"`
	// The standard deviation of the azimuth/right ascension rate residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension rate bias.
	AzRaRateSigma float64 `json:"azRaRateSigma"`
	// The root mean square of the azimuth/right-ascension residuals, in degrees, used
	// to determine the sensor azimuth/right-ascension bias.
	AzRaRms float64 `json:"azRaRms"`
	// The standard deviation of the azimuth/right ascension residuals, in degrees,
	// used to determine the sensor azimuth/right-ascension bias.
	AzRaSigma float64 `json:"azRaSigma"`
	// Specifies the calibration reference angle set for this calibration data set.
	// Azimuth and Elevation (AZEL) or Right Ascension and Declination (RADEC).
	CalAngleRef string `json:"calAngleRef"`
	// Specifies that the calibration data are from INTRA_TRACK or INTER_TRACK
	// residuals.
	CalTrackMode string `json:"calTrackMode"`
	// The basis of calibration values contained in this record (COMPUTED,
	// OPERATIONAL).
	CalType string `json:"calType"`
	// The confidence noise bias of the duration span.
	ConfidenceNoiseBias float64 `json:"confidenceNoiseBias"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Duration of the sensor calibration data which produced these values, measured in
	// days.
	Duration float64 `json:"duration"`
	// Three element array, expressing the sensor location in Earth Centered Rotating
	// (ECR) coordinates, in kilometers. The array element order is [x, y, z].
	Ecr []float64 `json:"ecr"`
	// Sensor elevation/declination acceleration bias, in degrees per second squared.
	ElDecAccelBias float64 `json:"elDecAccelBias"`
	// The standard deviation of the elevation/declination acceleration residuals, in
	// degrees, used to determine the sensor elevation/declination acceleration bias.
	ElDecAccelSigma float64 `json:"elDecAccelSigma"`
	// Sensor elevation/declination bias, in degrees.
	ElDecBias float64 `json:"elDecBias"`
	// Sensor elevation/declination rate bias, in degrees per second.
	ElDecRateBias float64 `json:"elDecRateBias"`
	// The standard deviation of the elevation/declination rate residuals, in degrees,
	// used to determine the sensor elevation/declination rate bias.
	ElDecRateSigma float64 `json:"elDecRateSigma"`
	// The root mean square of the elevation/declination residuals, in degrees, used to
	// determine the sensor elevation/declination bias.
	ElDecRms float64 `json:"elDecRms"`
	// The standard deviation of the elevation/declination residuals, in degrees, used
	// to determine the sensor elevation/declination bias.
	ElDecSigma float64 `json:"elDecSigma"`
	// Calibration data span end time in ISO 8601 UTC format with millisecond
	// precision. If provided, the endTime must be greater than or equal to the
	// startTime in the SensorCalibration record.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The number of observables used in determining the azimuth or right-ascension
	// calibration values.
	NumAzRaObs int64 `json:"numAzRaObs"`
	// The number of observables used in determining the elevation or declination
	// calibration values.
	NumElDecObs int64 `json:"numElDecObs"`
	// The total number of observables available over the calibration span.
	NumObs int64 `json:"numObs"`
	// The number of observables used in determining the photometric calibration
	// values.
	NumPhotoObs int64 `json:"numPhotoObs"`
	// The number of observables used in determining the range calibration values.
	NumRangeObs int64 `json:"numRangeObs"`
	// The number of observables used in determining the range rate calibration values.
	NumRangeRateObs int64 `json:"numRangeRateObs"`
	// The number of observables used in determining the radar cross section (RCS)
	// calibration values.
	NumRcsObs int64 `json:"numRcsObs"`
	// The number of observables used in determining the time calibration values.
	NumTimeObs int64 `json:"numTimeObs"`
	// The total number of tracks available over the calibration span.
	NumTracks int64 `json:"numTracks"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The sensor photometric observation magnitude bias, in visual magnitude.
	PhotoBias float64 `json:"photoBias"`
	// The standard deviation of the magnitude residuals, in visual magnitude, used to
	// determine the photometric bias.
	PhotoSigma float64 `json:"photoSigma"`
	// Sensor range rate acceleration bias, in kilometers per second squared.
	RangeAccelBias float64 `json:"rangeAccelBias"`
	// The standard deviation of the range acceleration residuals, in kilometers per
	// second squared, used to determine the sensor range acceleration bias.
	RangeAccelSigma float64 `json:"rangeAccelSigma"`
	// Sensor range bias, in kilometers.
	RangeBias float64 `json:"rangeBias"`
	// Sensor range rate bias, in kilometers per second.
	RangeRateBias float64 `json:"rangeRateBias"`
	// The root mean square of the range rate residuals, in kilometers per second, used
	// to determine the sensor range rate bias.
	RangeRateRms float64 `json:"rangeRateRms"`
	// The standard deviation of the range rate residuals, in kilometers per second,
	// used to determine the sensor range rate bias.
	RangeRateSigma float64 `json:"rangeRateSigma"`
	// The root mean square of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeRms float64 `json:"rangeRms"`
	// The standard deviation of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeSigma float64 `json:"rangeSigma"`
	// The sensor radar cross section (RCS) observation bias, in square meters.
	RcsBias float64 `json:"rcsBias"`
	// The standard deviation of the radar cross section residuals, in square meters,
	// used to determine the radar cross section bias.
	RcsSigma float64 `json:"rcsSigma"`
	// Array of the catalog IDs of the reference targets used in the calibration.
	RefTargets []string `json:"refTargets"`
	// The reference type used in the calibration.
	RefType string `json:"refType"`
	// The sensor type (MECHANICAL, OPTICAL, PHASED ARRAY, RF).
	SenType string `json:"senType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Sensor time bias, in seconds.
	TimeBias float64 `json:"timeBias"`
	// The standard deviation of the time residuals, in seconds, used to determine the
	// sensor time bias.
	TimeBiasSigma float64 `json:"timeBiasSigma"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSensor              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		AzRaAccelBias         respjson.Field
		AzRaAccelSigma        respjson.Field
		AzRaBias              respjson.Field
		AzRaRateBias          respjson.Field
		AzRaRateSigma         respjson.Field
		AzRaRms               respjson.Field
		AzRaSigma             respjson.Field
		CalAngleRef           respjson.Field
		CalTrackMode          respjson.Field
		CalType               respjson.Field
		ConfidenceNoiseBias   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Duration              respjson.Field
		Ecr                   respjson.Field
		ElDecAccelBias        respjson.Field
		ElDecAccelSigma       respjson.Field
		ElDecBias             respjson.Field
		ElDecRateBias         respjson.Field
		ElDecRateSigma        respjson.Field
		ElDecRms              respjson.Field
		ElDecSigma            respjson.Field
		EndTime               respjson.Field
		NumAzRaObs            respjson.Field
		NumElDecObs           respjson.Field
		NumObs                respjson.Field
		NumPhotoObs           respjson.Field
		NumRangeObs           respjson.Field
		NumRangeRateObs       respjson.Field
		NumRcsObs             respjson.Field
		NumTimeObs            respjson.Field
		NumTracks             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PhotoBias             respjson.Field
		PhotoSigma            respjson.Field
		RangeAccelBias        respjson.Field
		RangeAccelSigma       respjson.Field
		RangeBias             respjson.Field
		RangeRateBias         respjson.Field
		RangeRateRms          respjson.Field
		RangeRateSigma        respjson.Field
		RangeRms              respjson.Field
		RangeSigma            respjson.Field
		RcsBias               respjson.Field
		RcsSigma              respjson.Field
		RefTargets            respjson.Field
		RefType               respjson.Field
		SenType               respjson.Field
		SourceDl              respjson.Field
		TimeBias              respjson.Field
		TimeBiasSigma         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorCalibrationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorCalibrationGetResponse) UnmarshalJSON(data []byte) error {
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
type SensorCalibrationGetResponseDataMode string

const (
	SensorCalibrationGetResponseDataModeReal      SensorCalibrationGetResponseDataMode = "REAL"
	SensorCalibrationGetResponseDataModeTest      SensorCalibrationGetResponseDataMode = "TEST"
	SensorCalibrationGetResponseDataModeSimulated SensorCalibrationGetResponseDataMode = "SIMULATED"
	SensorCalibrationGetResponseDataModeExercise  SensorCalibrationGetResponseDataMode = "EXERCISE"
)

// The Sensor Calibration service records data about a sensor's overall accuracy
// and is used to adjust sensor settings to achieve and maintain that accuracy in
// reported sensor observations. Calibration occurs periodically when needed to
// maintain sensor accuracy or on-demand to adjust a sensor for a specific reading.
type SensorCalibrationQueryResponse struct {
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
	DataMode SensorCalibrationQueryResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the sensor to which this calibration data applies. This ID
	// can be used to obtain additional information on a sensor using the 'get by ID'
	// operation (e.g. /udl/sensor/{id}). For example, the sensor with idSensor = abc
	// would be queried as /udl/sensor/abc.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Sensor azimuth/right-ascension acceleration bias, in degrees per second squared.
	AzRaAccelBias float64 `json:"azRaAccelBias"`
	// The standard deviation of the azimuth/right ascension acceleration residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension acceleration bias.
	AzRaAccelSigma float64 `json:"azRaAccelSigma"`
	// Sensor azimuth/right-ascension bias, in degrees.
	AzRaBias float64 `json:"azRaBias"`
	// Sensor azimuth/right-ascension rate bias, in degrees per second.
	AzRaRateBias float64 `json:"azRaRateBias"`
	// The standard deviation of the azimuth/right ascension rate residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension rate bias.
	AzRaRateSigma float64 `json:"azRaRateSigma"`
	// The root mean square of the azimuth/right-ascension residuals, in degrees, used
	// to determine the sensor azimuth/right-ascension bias.
	AzRaRms float64 `json:"azRaRms"`
	// The standard deviation of the azimuth/right ascension residuals, in degrees,
	// used to determine the sensor azimuth/right-ascension bias.
	AzRaSigma float64 `json:"azRaSigma"`
	// Specifies the calibration reference angle set for this calibration data set.
	// Azimuth and Elevation (AZEL) or Right Ascension and Declination (RADEC).
	CalAngleRef string `json:"calAngleRef"`
	// Specifies that the calibration data are from INTRA_TRACK or INTER_TRACK
	// residuals.
	CalTrackMode string `json:"calTrackMode"`
	// The basis of calibration values contained in this record (COMPUTED,
	// OPERATIONAL).
	CalType string `json:"calType"`
	// The confidence noise bias of the duration span.
	ConfidenceNoiseBias float64 `json:"confidenceNoiseBias"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Duration of the sensor calibration data which produced these values, measured in
	// days.
	Duration float64 `json:"duration"`
	// Three element array, expressing the sensor location in Earth Centered Rotating
	// (ECR) coordinates, in kilometers. The array element order is [x, y, z].
	Ecr []float64 `json:"ecr"`
	// Sensor elevation/declination acceleration bias, in degrees per second squared.
	ElDecAccelBias float64 `json:"elDecAccelBias"`
	// The standard deviation of the elevation/declination acceleration residuals, in
	// degrees, used to determine the sensor elevation/declination acceleration bias.
	ElDecAccelSigma float64 `json:"elDecAccelSigma"`
	// Sensor elevation/declination bias, in degrees.
	ElDecBias float64 `json:"elDecBias"`
	// Sensor elevation/declination rate bias, in degrees per second.
	ElDecRateBias float64 `json:"elDecRateBias"`
	// The standard deviation of the elevation/declination rate residuals, in degrees,
	// used to determine the sensor elevation/declination rate bias.
	ElDecRateSigma float64 `json:"elDecRateSigma"`
	// The root mean square of the elevation/declination residuals, in degrees, used to
	// determine the sensor elevation/declination bias.
	ElDecRms float64 `json:"elDecRms"`
	// The standard deviation of the elevation/declination residuals, in degrees, used
	// to determine the sensor elevation/declination bias.
	ElDecSigma float64 `json:"elDecSigma"`
	// Calibration data span end time in ISO 8601 UTC format with millisecond
	// precision. If provided, the endTime must be greater than or equal to the
	// startTime in the SensorCalibration record.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The number of observables used in determining the azimuth or right-ascension
	// calibration values.
	NumAzRaObs int64 `json:"numAzRaObs"`
	// The number of observables used in determining the elevation or declination
	// calibration values.
	NumElDecObs int64 `json:"numElDecObs"`
	// The total number of observables available over the calibration span.
	NumObs int64 `json:"numObs"`
	// The number of observables used in determining the photometric calibration
	// values.
	NumPhotoObs int64 `json:"numPhotoObs"`
	// The number of observables used in determining the range calibration values.
	NumRangeObs int64 `json:"numRangeObs"`
	// The number of observables used in determining the range rate calibration values.
	NumRangeRateObs int64 `json:"numRangeRateObs"`
	// The number of observables used in determining the radar cross section (RCS)
	// calibration values.
	NumRcsObs int64 `json:"numRcsObs"`
	// The number of observables used in determining the time calibration values.
	NumTimeObs int64 `json:"numTimeObs"`
	// The total number of tracks available over the calibration span.
	NumTracks int64 `json:"numTracks"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The sensor photometric observation magnitude bias, in visual magnitude.
	PhotoBias float64 `json:"photoBias"`
	// The standard deviation of the magnitude residuals, in visual magnitude, used to
	// determine the photometric bias.
	PhotoSigma float64 `json:"photoSigma"`
	// Sensor range rate acceleration bias, in kilometers per second squared.
	RangeAccelBias float64 `json:"rangeAccelBias"`
	// The standard deviation of the range acceleration residuals, in kilometers per
	// second squared, used to determine the sensor range acceleration bias.
	RangeAccelSigma float64 `json:"rangeAccelSigma"`
	// Sensor range bias, in kilometers.
	RangeBias float64 `json:"rangeBias"`
	// Sensor range rate bias, in kilometers per second.
	RangeRateBias float64 `json:"rangeRateBias"`
	// The root mean square of the range rate residuals, in kilometers per second, used
	// to determine the sensor range rate bias.
	RangeRateRms float64 `json:"rangeRateRms"`
	// The standard deviation of the range rate residuals, in kilometers per second,
	// used to determine the sensor range rate bias.
	RangeRateSigma float64 `json:"rangeRateSigma"`
	// The root mean square of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeRms float64 `json:"rangeRms"`
	// The standard deviation of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeSigma float64 `json:"rangeSigma"`
	// The sensor radar cross section (RCS) observation bias, in square meters.
	RcsBias float64 `json:"rcsBias"`
	// The standard deviation of the radar cross section residuals, in square meters,
	// used to determine the radar cross section bias.
	RcsSigma float64 `json:"rcsSigma"`
	// Array of the catalog IDs of the reference targets used in the calibration.
	RefTargets []string `json:"refTargets"`
	// The reference type used in the calibration.
	RefType string `json:"refType"`
	// The sensor type (MECHANICAL, OPTICAL, PHASED ARRAY, RF).
	SenType string `json:"senType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Sensor time bias, in seconds.
	TimeBias float64 `json:"timeBias"`
	// The standard deviation of the time residuals, in seconds, used to determine the
	// sensor time bias.
	TimeBiasSigma float64 `json:"timeBiasSigma"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSensor              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		AzRaAccelBias         respjson.Field
		AzRaAccelSigma        respjson.Field
		AzRaBias              respjson.Field
		AzRaRateBias          respjson.Field
		AzRaRateSigma         respjson.Field
		AzRaRms               respjson.Field
		AzRaSigma             respjson.Field
		CalAngleRef           respjson.Field
		CalTrackMode          respjson.Field
		CalType               respjson.Field
		ConfidenceNoiseBias   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Duration              respjson.Field
		Ecr                   respjson.Field
		ElDecAccelBias        respjson.Field
		ElDecAccelSigma       respjson.Field
		ElDecBias             respjson.Field
		ElDecRateBias         respjson.Field
		ElDecRateSigma        respjson.Field
		ElDecRms              respjson.Field
		ElDecSigma            respjson.Field
		EndTime               respjson.Field
		NumAzRaObs            respjson.Field
		NumElDecObs           respjson.Field
		NumObs                respjson.Field
		NumPhotoObs           respjson.Field
		NumRangeObs           respjson.Field
		NumRangeRateObs       respjson.Field
		NumRcsObs             respjson.Field
		NumTimeObs            respjson.Field
		NumTracks             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PhotoBias             respjson.Field
		PhotoSigma            respjson.Field
		RangeAccelBias        respjson.Field
		RangeAccelSigma       respjson.Field
		RangeBias             respjson.Field
		RangeRateBias         respjson.Field
		RangeRateRms          respjson.Field
		RangeRateSigma        respjson.Field
		RangeRms              respjson.Field
		RangeSigma            respjson.Field
		RcsBias               respjson.Field
		RcsSigma              respjson.Field
		RefTargets            respjson.Field
		RefType               respjson.Field
		SenType               respjson.Field
		SourceDl              respjson.Field
		TimeBias              respjson.Field
		TimeBiasSigma         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorCalibrationQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorCalibrationQueryResponse) UnmarshalJSON(data []byte) error {
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
type SensorCalibrationQueryResponseDataMode string

const (
	SensorCalibrationQueryResponseDataModeReal      SensorCalibrationQueryResponseDataMode = "REAL"
	SensorCalibrationQueryResponseDataModeTest      SensorCalibrationQueryResponseDataMode = "TEST"
	SensorCalibrationQueryResponseDataModeSimulated SensorCalibrationQueryResponseDataMode = "SIMULATED"
	SensorCalibrationQueryResponseDataModeExercise  SensorCalibrationQueryResponseDataMode = "EXERCISE"
)

type SensorCalibrationQueryHelpResponse struct {
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
func (r SensorCalibrationQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorCalibrationQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Sensor Calibration service records data about a sensor's overall accuracy
// and is used to adjust sensor settings to achieve and maintain that accuracy in
// reported sensor observations. Calibration occurs periodically when needed to
// maintain sensor accuracy or on-demand to adjust a sensor for a specific reading.
type SensorCalibrationTupleResponse struct {
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
	DataMode SensorCalibrationTupleResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the sensor to which this calibration data applies. This ID
	// can be used to obtain additional information on a sensor using the 'get by ID'
	// operation (e.g. /udl/sensor/{id}). For example, the sensor with idSensor = abc
	// would be queried as /udl/sensor/abc.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Sensor azimuth/right-ascension acceleration bias, in degrees per second squared.
	AzRaAccelBias float64 `json:"azRaAccelBias"`
	// The standard deviation of the azimuth/right ascension acceleration residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension acceleration bias.
	AzRaAccelSigma float64 `json:"azRaAccelSigma"`
	// Sensor azimuth/right-ascension bias, in degrees.
	AzRaBias float64 `json:"azRaBias"`
	// Sensor azimuth/right-ascension rate bias, in degrees per second.
	AzRaRateBias float64 `json:"azRaRateBias"`
	// The standard deviation of the azimuth/right ascension rate residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension rate bias.
	AzRaRateSigma float64 `json:"azRaRateSigma"`
	// The root mean square of the azimuth/right-ascension residuals, in degrees, used
	// to determine the sensor azimuth/right-ascension bias.
	AzRaRms float64 `json:"azRaRms"`
	// The standard deviation of the azimuth/right ascension residuals, in degrees,
	// used to determine the sensor azimuth/right-ascension bias.
	AzRaSigma float64 `json:"azRaSigma"`
	// Specifies the calibration reference angle set for this calibration data set.
	// Azimuth and Elevation (AZEL) or Right Ascension and Declination (RADEC).
	CalAngleRef string `json:"calAngleRef"`
	// Specifies that the calibration data are from INTRA_TRACK or INTER_TRACK
	// residuals.
	CalTrackMode string `json:"calTrackMode"`
	// The basis of calibration values contained in this record (COMPUTED,
	// OPERATIONAL).
	CalType string `json:"calType"`
	// The confidence noise bias of the duration span.
	ConfidenceNoiseBias float64 `json:"confidenceNoiseBias"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Duration of the sensor calibration data which produced these values, measured in
	// days.
	Duration float64 `json:"duration"`
	// Three element array, expressing the sensor location in Earth Centered Rotating
	// (ECR) coordinates, in kilometers. The array element order is [x, y, z].
	Ecr []float64 `json:"ecr"`
	// Sensor elevation/declination acceleration bias, in degrees per second squared.
	ElDecAccelBias float64 `json:"elDecAccelBias"`
	// The standard deviation of the elevation/declination acceleration residuals, in
	// degrees, used to determine the sensor elevation/declination acceleration bias.
	ElDecAccelSigma float64 `json:"elDecAccelSigma"`
	// Sensor elevation/declination bias, in degrees.
	ElDecBias float64 `json:"elDecBias"`
	// Sensor elevation/declination rate bias, in degrees per second.
	ElDecRateBias float64 `json:"elDecRateBias"`
	// The standard deviation of the elevation/declination rate residuals, in degrees,
	// used to determine the sensor elevation/declination rate bias.
	ElDecRateSigma float64 `json:"elDecRateSigma"`
	// The root mean square of the elevation/declination residuals, in degrees, used to
	// determine the sensor elevation/declination bias.
	ElDecRms float64 `json:"elDecRms"`
	// The standard deviation of the elevation/declination residuals, in degrees, used
	// to determine the sensor elevation/declination bias.
	ElDecSigma float64 `json:"elDecSigma"`
	// Calibration data span end time in ISO 8601 UTC format with millisecond
	// precision. If provided, the endTime must be greater than or equal to the
	// startTime in the SensorCalibration record.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The number of observables used in determining the azimuth or right-ascension
	// calibration values.
	NumAzRaObs int64 `json:"numAzRaObs"`
	// The number of observables used in determining the elevation or declination
	// calibration values.
	NumElDecObs int64 `json:"numElDecObs"`
	// The total number of observables available over the calibration span.
	NumObs int64 `json:"numObs"`
	// The number of observables used in determining the photometric calibration
	// values.
	NumPhotoObs int64 `json:"numPhotoObs"`
	// The number of observables used in determining the range calibration values.
	NumRangeObs int64 `json:"numRangeObs"`
	// The number of observables used in determining the range rate calibration values.
	NumRangeRateObs int64 `json:"numRangeRateObs"`
	// The number of observables used in determining the radar cross section (RCS)
	// calibration values.
	NumRcsObs int64 `json:"numRcsObs"`
	// The number of observables used in determining the time calibration values.
	NumTimeObs int64 `json:"numTimeObs"`
	// The total number of tracks available over the calibration span.
	NumTracks int64 `json:"numTracks"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The sensor photometric observation magnitude bias, in visual magnitude.
	PhotoBias float64 `json:"photoBias"`
	// The standard deviation of the magnitude residuals, in visual magnitude, used to
	// determine the photometric bias.
	PhotoSigma float64 `json:"photoSigma"`
	// Sensor range rate acceleration bias, in kilometers per second squared.
	RangeAccelBias float64 `json:"rangeAccelBias"`
	// The standard deviation of the range acceleration residuals, in kilometers per
	// second squared, used to determine the sensor range acceleration bias.
	RangeAccelSigma float64 `json:"rangeAccelSigma"`
	// Sensor range bias, in kilometers.
	RangeBias float64 `json:"rangeBias"`
	// Sensor range rate bias, in kilometers per second.
	RangeRateBias float64 `json:"rangeRateBias"`
	// The root mean square of the range rate residuals, in kilometers per second, used
	// to determine the sensor range rate bias.
	RangeRateRms float64 `json:"rangeRateRms"`
	// The standard deviation of the range rate residuals, in kilometers per second,
	// used to determine the sensor range rate bias.
	RangeRateSigma float64 `json:"rangeRateSigma"`
	// The root mean square of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeRms float64 `json:"rangeRms"`
	// The standard deviation of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeSigma float64 `json:"rangeSigma"`
	// The sensor radar cross section (RCS) observation bias, in square meters.
	RcsBias float64 `json:"rcsBias"`
	// The standard deviation of the radar cross section residuals, in square meters,
	// used to determine the radar cross section bias.
	RcsSigma float64 `json:"rcsSigma"`
	// Array of the catalog IDs of the reference targets used in the calibration.
	RefTargets []string `json:"refTargets"`
	// The reference type used in the calibration.
	RefType string `json:"refType"`
	// The sensor type (MECHANICAL, OPTICAL, PHASED ARRAY, RF).
	SenType string `json:"senType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Sensor time bias, in seconds.
	TimeBias float64 `json:"timeBias"`
	// The standard deviation of the time residuals, in seconds, used to determine the
	// sensor time bias.
	TimeBiasSigma float64 `json:"timeBiasSigma"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSensor              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		AzRaAccelBias         respjson.Field
		AzRaAccelSigma        respjson.Field
		AzRaBias              respjson.Field
		AzRaRateBias          respjson.Field
		AzRaRateSigma         respjson.Field
		AzRaRms               respjson.Field
		AzRaSigma             respjson.Field
		CalAngleRef           respjson.Field
		CalTrackMode          respjson.Field
		CalType               respjson.Field
		ConfidenceNoiseBias   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Duration              respjson.Field
		Ecr                   respjson.Field
		ElDecAccelBias        respjson.Field
		ElDecAccelSigma       respjson.Field
		ElDecBias             respjson.Field
		ElDecRateBias         respjson.Field
		ElDecRateSigma        respjson.Field
		ElDecRms              respjson.Field
		ElDecSigma            respjson.Field
		EndTime               respjson.Field
		NumAzRaObs            respjson.Field
		NumElDecObs           respjson.Field
		NumObs                respjson.Field
		NumPhotoObs           respjson.Field
		NumRangeObs           respjson.Field
		NumRangeRateObs       respjson.Field
		NumRcsObs             respjson.Field
		NumTimeObs            respjson.Field
		NumTracks             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PhotoBias             respjson.Field
		PhotoSigma            respjson.Field
		RangeAccelBias        respjson.Field
		RangeAccelSigma       respjson.Field
		RangeBias             respjson.Field
		RangeRateBias         respjson.Field
		RangeRateRms          respjson.Field
		RangeRateSigma        respjson.Field
		RangeRms              respjson.Field
		RangeSigma            respjson.Field
		RcsBias               respjson.Field
		RcsSigma              respjson.Field
		RefTargets            respjson.Field
		RefType               respjson.Field
		SenType               respjson.Field
		SourceDl              respjson.Field
		TimeBias              respjson.Field
		TimeBiasSigma         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorCalibrationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorCalibrationTupleResponse) UnmarshalJSON(data []byte) error {
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
type SensorCalibrationTupleResponseDataMode string

const (
	SensorCalibrationTupleResponseDataModeReal      SensorCalibrationTupleResponseDataMode = "REAL"
	SensorCalibrationTupleResponseDataModeTest      SensorCalibrationTupleResponseDataMode = "TEST"
	SensorCalibrationTupleResponseDataModeSimulated SensorCalibrationTupleResponseDataMode = "SIMULATED"
	SensorCalibrationTupleResponseDataModeExercise  SensorCalibrationTupleResponseDataMode = "EXERCISE"
)

type SensorCalibrationNewParams struct {
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
	DataMode SensorCalibrationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the sensor to which this calibration data applies. This ID
	// can be used to obtain additional information on a sensor using the 'get by ID'
	// operation (e.g. /udl/sensor/{id}). For example, the sensor with idSensor = abc
	// would be queried as /udl/sensor/abc.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Sensor azimuth/right-ascension acceleration bias, in degrees per second squared.
	AzRaAccelBias param.Opt[float64] `json:"azRaAccelBias,omitzero"`
	// The standard deviation of the azimuth/right ascension acceleration residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension acceleration bias.
	AzRaAccelSigma param.Opt[float64] `json:"azRaAccelSigma,omitzero"`
	// Sensor azimuth/right-ascension bias, in degrees.
	AzRaBias param.Opt[float64] `json:"azRaBias,omitzero"`
	// Sensor azimuth/right-ascension rate bias, in degrees per second.
	AzRaRateBias param.Opt[float64] `json:"azRaRateBias,omitzero"`
	// The standard deviation of the azimuth/right ascension rate residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension rate bias.
	AzRaRateSigma param.Opt[float64] `json:"azRaRateSigma,omitzero"`
	// The root mean square of the azimuth/right-ascension residuals, in degrees, used
	// to determine the sensor azimuth/right-ascension bias.
	AzRaRms param.Opt[float64] `json:"azRaRms,omitzero"`
	// The standard deviation of the azimuth/right ascension residuals, in degrees,
	// used to determine the sensor azimuth/right-ascension bias.
	AzRaSigma param.Opt[float64] `json:"azRaSigma,omitzero"`
	// Specifies the calibration reference angle set for this calibration data set.
	// Azimuth and Elevation (AZEL) or Right Ascension and Declination (RADEC).
	CalAngleRef param.Opt[string] `json:"calAngleRef,omitzero"`
	// Specifies that the calibration data are from INTRA_TRACK or INTER_TRACK
	// residuals.
	CalTrackMode param.Opt[string] `json:"calTrackMode,omitzero"`
	// The basis of calibration values contained in this record (COMPUTED,
	// OPERATIONAL).
	CalType param.Opt[string] `json:"calType,omitzero"`
	// The confidence noise bias of the duration span.
	ConfidenceNoiseBias param.Opt[float64] `json:"confidenceNoiseBias,omitzero"`
	// Duration of the sensor calibration data which produced these values, measured in
	// days.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Sensor elevation/declination acceleration bias, in degrees per second squared.
	ElDecAccelBias param.Opt[float64] `json:"elDecAccelBias,omitzero"`
	// The standard deviation of the elevation/declination acceleration residuals, in
	// degrees, used to determine the sensor elevation/declination acceleration bias.
	ElDecAccelSigma param.Opt[float64] `json:"elDecAccelSigma,omitzero"`
	// Sensor elevation/declination bias, in degrees.
	ElDecBias param.Opt[float64] `json:"elDecBias,omitzero"`
	// Sensor elevation/declination rate bias, in degrees per second.
	ElDecRateBias param.Opt[float64] `json:"elDecRateBias,omitzero"`
	// The standard deviation of the elevation/declination rate residuals, in degrees,
	// used to determine the sensor elevation/declination rate bias.
	ElDecRateSigma param.Opt[float64] `json:"elDecRateSigma,omitzero"`
	// The root mean square of the elevation/declination residuals, in degrees, used to
	// determine the sensor elevation/declination bias.
	ElDecRms param.Opt[float64] `json:"elDecRms,omitzero"`
	// The standard deviation of the elevation/declination residuals, in degrees, used
	// to determine the sensor elevation/declination bias.
	ElDecSigma param.Opt[float64] `json:"elDecSigma,omitzero"`
	// Calibration data span end time in ISO 8601 UTC format with millisecond
	// precision. If provided, the endTime must be greater than or equal to the
	// startTime in the SensorCalibration record.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The number of observables used in determining the azimuth or right-ascension
	// calibration values.
	NumAzRaObs param.Opt[int64] `json:"numAzRaObs,omitzero"`
	// The number of observables used in determining the elevation or declination
	// calibration values.
	NumElDecObs param.Opt[int64] `json:"numElDecObs,omitzero"`
	// The total number of observables available over the calibration span.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// The number of observables used in determining the photometric calibration
	// values.
	NumPhotoObs param.Opt[int64] `json:"numPhotoObs,omitzero"`
	// The number of observables used in determining the range calibration values.
	NumRangeObs param.Opt[int64] `json:"numRangeObs,omitzero"`
	// The number of observables used in determining the range rate calibration values.
	NumRangeRateObs param.Opt[int64] `json:"numRangeRateObs,omitzero"`
	// The number of observables used in determining the radar cross section (RCS)
	// calibration values.
	NumRcsObs param.Opt[int64] `json:"numRcsObs,omitzero"`
	// The number of observables used in determining the time calibration values.
	NumTimeObs param.Opt[int64] `json:"numTimeObs,omitzero"`
	// The total number of tracks available over the calibration span.
	NumTracks param.Opt[int64] `json:"numTracks,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The sensor photometric observation magnitude bias, in visual magnitude.
	PhotoBias param.Opt[float64] `json:"photoBias,omitzero"`
	// The standard deviation of the magnitude residuals, in visual magnitude, used to
	// determine the photometric bias.
	PhotoSigma param.Opt[float64] `json:"photoSigma,omitzero"`
	// Sensor range rate acceleration bias, in kilometers per second squared.
	RangeAccelBias param.Opt[float64] `json:"rangeAccelBias,omitzero"`
	// The standard deviation of the range acceleration residuals, in kilometers per
	// second squared, used to determine the sensor range acceleration bias.
	RangeAccelSigma param.Opt[float64] `json:"rangeAccelSigma,omitzero"`
	// Sensor range bias, in kilometers.
	RangeBias param.Opt[float64] `json:"rangeBias,omitzero"`
	// Sensor range rate bias, in kilometers per second.
	RangeRateBias param.Opt[float64] `json:"rangeRateBias,omitzero"`
	// The root mean square of the range rate residuals, in kilometers per second, used
	// to determine the sensor range rate bias.
	RangeRateRms param.Opt[float64] `json:"rangeRateRms,omitzero"`
	// The standard deviation of the range rate residuals, in kilometers per second,
	// used to determine the sensor range rate bias.
	RangeRateSigma param.Opt[float64] `json:"rangeRateSigma,omitzero"`
	// The root mean square of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeRms param.Opt[float64] `json:"rangeRms,omitzero"`
	// The standard deviation of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeSigma param.Opt[float64] `json:"rangeSigma,omitzero"`
	// The sensor radar cross section (RCS) observation bias, in square meters.
	RcsBias param.Opt[float64] `json:"rcsBias,omitzero"`
	// The standard deviation of the radar cross section residuals, in square meters,
	// used to determine the radar cross section bias.
	RcsSigma param.Opt[float64] `json:"rcsSigma,omitzero"`
	// The reference type used in the calibration.
	RefType param.Opt[string] `json:"refType,omitzero"`
	// The sensor type (MECHANICAL, OPTICAL, PHASED ARRAY, RF).
	SenType param.Opt[string] `json:"senType,omitzero"`
	// Sensor time bias, in seconds.
	TimeBias param.Opt[float64] `json:"timeBias,omitzero"`
	// The standard deviation of the time residuals, in seconds, used to determine the
	// sensor time bias.
	TimeBiasSigma param.Opt[float64] `json:"timeBiasSigma,omitzero"`
	// Three element array, expressing the sensor location in Earth Centered Rotating
	// (ECR) coordinates, in kilometers. The array element order is [x, y, z].
	Ecr []float64 `json:"ecr,omitzero"`
	// Array of the catalog IDs of the reference targets used in the calibration.
	RefTargets []string `json:"refTargets,omitzero"`
	paramObj
}

func (r SensorCalibrationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SensorCalibrationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorCalibrationNewParams) UnmarshalJSON(data []byte) error {
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
type SensorCalibrationNewParamsDataMode string

const (
	SensorCalibrationNewParamsDataModeReal      SensorCalibrationNewParamsDataMode = "REAL"
	SensorCalibrationNewParamsDataModeTest      SensorCalibrationNewParamsDataMode = "TEST"
	SensorCalibrationNewParamsDataModeSimulated SensorCalibrationNewParamsDataMode = "SIMULATED"
	SensorCalibrationNewParamsDataModeExercise  SensorCalibrationNewParamsDataMode = "EXERCISE"
)

type SensorCalibrationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorCalibrationGetParams]'s query parameters as
// `url.Values`.
func (r SensorCalibrationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorCalibrationCountParams struct {
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorCalibrationCountParams]'s query parameters as
// `url.Values`.
func (r SensorCalibrationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorCalibrationNewBulkParams struct {
	Body []SensorCalibrationNewBulkParamsBody
	paramObj
}

func (r SensorCalibrationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SensorCalibrationNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The Sensor Calibration service records data about a sensor's overall accuracy
// and is used to adjust sensor settings to achieve and maintain that accuracy in
// reported sensor observations. Calibration occurs periodically when needed to
// maintain sensor accuracy or on-demand to adjust a sensor for a specific reading.
//
// The properties ClassificationMarking, DataMode, IDSensor, Source, StartTime are
// required.
type SensorCalibrationNewBulkParamsBody struct {
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
	// Unique identifier of the sensor to which this calibration data applies. This ID
	// can be used to obtain additional information on a sensor using the 'get by ID'
	// operation (e.g. /udl/sensor/{id}). For example, the sensor with idSensor = abc
	// would be queried as /udl/sensor/abc.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Sensor azimuth/right-ascension acceleration bias, in degrees per second squared.
	AzRaAccelBias param.Opt[float64] `json:"azRaAccelBias,omitzero"`
	// The standard deviation of the azimuth/right ascension acceleration residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension acceleration bias.
	AzRaAccelSigma param.Opt[float64] `json:"azRaAccelSigma,omitzero"`
	// Sensor azimuth/right-ascension bias, in degrees.
	AzRaBias param.Opt[float64] `json:"azRaBias,omitzero"`
	// Sensor azimuth/right-ascension rate bias, in degrees per second.
	AzRaRateBias param.Opt[float64] `json:"azRaRateBias,omitzero"`
	// The standard deviation of the azimuth/right ascension rate residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension rate bias.
	AzRaRateSigma param.Opt[float64] `json:"azRaRateSigma,omitzero"`
	// The root mean square of the azimuth/right-ascension residuals, in degrees, used
	// to determine the sensor azimuth/right-ascension bias.
	AzRaRms param.Opt[float64] `json:"azRaRms,omitzero"`
	// The standard deviation of the azimuth/right ascension residuals, in degrees,
	// used to determine the sensor azimuth/right-ascension bias.
	AzRaSigma param.Opt[float64] `json:"azRaSigma,omitzero"`
	// Specifies the calibration reference angle set for this calibration data set.
	// Azimuth and Elevation (AZEL) or Right Ascension and Declination (RADEC).
	CalAngleRef param.Opt[string] `json:"calAngleRef,omitzero"`
	// Specifies that the calibration data are from INTRA_TRACK or INTER_TRACK
	// residuals.
	CalTrackMode param.Opt[string] `json:"calTrackMode,omitzero"`
	// The basis of calibration values contained in this record (COMPUTED,
	// OPERATIONAL).
	CalType param.Opt[string] `json:"calType,omitzero"`
	// The confidence noise bias of the duration span.
	ConfidenceNoiseBias param.Opt[float64] `json:"confidenceNoiseBias,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Duration of the sensor calibration data which produced these values, measured in
	// days.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Sensor elevation/declination acceleration bias, in degrees per second squared.
	ElDecAccelBias param.Opt[float64] `json:"elDecAccelBias,omitzero"`
	// The standard deviation of the elevation/declination acceleration residuals, in
	// degrees, used to determine the sensor elevation/declination acceleration bias.
	ElDecAccelSigma param.Opt[float64] `json:"elDecAccelSigma,omitzero"`
	// Sensor elevation/declination bias, in degrees.
	ElDecBias param.Opt[float64] `json:"elDecBias,omitzero"`
	// Sensor elevation/declination rate bias, in degrees per second.
	ElDecRateBias param.Opt[float64] `json:"elDecRateBias,omitzero"`
	// The standard deviation of the elevation/declination rate residuals, in degrees,
	// used to determine the sensor elevation/declination rate bias.
	ElDecRateSigma param.Opt[float64] `json:"elDecRateSigma,omitzero"`
	// The root mean square of the elevation/declination residuals, in degrees, used to
	// determine the sensor elevation/declination bias.
	ElDecRms param.Opt[float64] `json:"elDecRms,omitzero"`
	// The standard deviation of the elevation/declination residuals, in degrees, used
	// to determine the sensor elevation/declination bias.
	ElDecSigma param.Opt[float64] `json:"elDecSigma,omitzero"`
	// Calibration data span end time in ISO 8601 UTC format with millisecond
	// precision. If provided, the endTime must be greater than or equal to the
	// startTime in the SensorCalibration record.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The number of observables used in determining the azimuth or right-ascension
	// calibration values.
	NumAzRaObs param.Opt[int64] `json:"numAzRaObs,omitzero"`
	// The number of observables used in determining the elevation or declination
	// calibration values.
	NumElDecObs param.Opt[int64] `json:"numElDecObs,omitzero"`
	// The total number of observables available over the calibration span.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// The number of observables used in determining the photometric calibration
	// values.
	NumPhotoObs param.Opt[int64] `json:"numPhotoObs,omitzero"`
	// The number of observables used in determining the range calibration values.
	NumRangeObs param.Opt[int64] `json:"numRangeObs,omitzero"`
	// The number of observables used in determining the range rate calibration values.
	NumRangeRateObs param.Opt[int64] `json:"numRangeRateObs,omitzero"`
	// The number of observables used in determining the radar cross section (RCS)
	// calibration values.
	NumRcsObs param.Opt[int64] `json:"numRcsObs,omitzero"`
	// The number of observables used in determining the time calibration values.
	NumTimeObs param.Opt[int64] `json:"numTimeObs,omitzero"`
	// The total number of tracks available over the calibration span.
	NumTracks param.Opt[int64] `json:"numTracks,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The sensor photometric observation magnitude bias, in visual magnitude.
	PhotoBias param.Opt[float64] `json:"photoBias,omitzero"`
	// The standard deviation of the magnitude residuals, in visual magnitude, used to
	// determine the photometric bias.
	PhotoSigma param.Opt[float64] `json:"photoSigma,omitzero"`
	// Sensor range rate acceleration bias, in kilometers per second squared.
	RangeAccelBias param.Opt[float64] `json:"rangeAccelBias,omitzero"`
	// The standard deviation of the range acceleration residuals, in kilometers per
	// second squared, used to determine the sensor range acceleration bias.
	RangeAccelSigma param.Opt[float64] `json:"rangeAccelSigma,omitzero"`
	// Sensor range bias, in kilometers.
	RangeBias param.Opt[float64] `json:"rangeBias,omitzero"`
	// Sensor range rate bias, in kilometers per second.
	RangeRateBias param.Opt[float64] `json:"rangeRateBias,omitzero"`
	// The root mean square of the range rate residuals, in kilometers per second, used
	// to determine the sensor range rate bias.
	RangeRateRms param.Opt[float64] `json:"rangeRateRms,omitzero"`
	// The standard deviation of the range rate residuals, in kilometers per second,
	// used to determine the sensor range rate bias.
	RangeRateSigma param.Opt[float64] `json:"rangeRateSigma,omitzero"`
	// The root mean square of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeRms param.Opt[float64] `json:"rangeRms,omitzero"`
	// The standard deviation of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeSigma param.Opt[float64] `json:"rangeSigma,omitzero"`
	// The sensor radar cross section (RCS) observation bias, in square meters.
	RcsBias param.Opt[float64] `json:"rcsBias,omitzero"`
	// The standard deviation of the radar cross section residuals, in square meters,
	// used to determine the radar cross section bias.
	RcsSigma param.Opt[float64] `json:"rcsSigma,omitzero"`
	// The reference type used in the calibration.
	RefType param.Opt[string] `json:"refType,omitzero"`
	// The sensor type (MECHANICAL, OPTICAL, PHASED ARRAY, RF).
	SenType param.Opt[string] `json:"senType,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Sensor time bias, in seconds.
	TimeBias param.Opt[float64] `json:"timeBias,omitzero"`
	// The standard deviation of the time residuals, in seconds, used to determine the
	// sensor time bias.
	TimeBiasSigma param.Opt[float64] `json:"timeBiasSigma,omitzero"`
	// Three element array, expressing the sensor location in Earth Centered Rotating
	// (ECR) coordinates, in kilometers. The array element order is [x, y, z].
	Ecr []float64 `json:"ecr,omitzero"`
	// Array of the catalog IDs of the reference targets used in the calibration.
	RefTargets []string `json:"refTargets,omitzero"`
	paramObj
}

func (r SensorCalibrationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SensorCalibrationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorCalibrationNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorCalibrationNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SensorCalibrationQueryParams struct {
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorCalibrationQueryParams]'s query parameters as
// `url.Values`.
func (r SensorCalibrationQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorCalibrationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorCalibrationTupleParams]'s query parameters as
// `url.Values`.
func (r SensorCalibrationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorCalibrationUnvalidatedPublishParams struct {
	Body []SensorCalibrationUnvalidatedPublishParamsBody
	paramObj
}

func (r SensorCalibrationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SensorCalibrationUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The Sensor Calibration service records data about a sensor's overall accuracy
// and is used to adjust sensor settings to achieve and maintain that accuracy in
// reported sensor observations. Calibration occurs periodically when needed to
// maintain sensor accuracy or on-demand to adjust a sensor for a specific reading.
//
// The properties ClassificationMarking, DataMode, IDSensor, Source, StartTime are
// required.
type SensorCalibrationUnvalidatedPublishParamsBody struct {
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
	// Unique identifier of the sensor to which this calibration data applies. This ID
	// can be used to obtain additional information on a sensor using the 'get by ID'
	// operation (e.g. /udl/sensor/{id}). For example, the sensor with idSensor = abc
	// would be queried as /udl/sensor/abc.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Sensor azimuth/right-ascension acceleration bias, in degrees per second squared.
	AzRaAccelBias param.Opt[float64] `json:"azRaAccelBias,omitzero"`
	// The standard deviation of the azimuth/right ascension acceleration residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension acceleration bias.
	AzRaAccelSigma param.Opt[float64] `json:"azRaAccelSigma,omitzero"`
	// Sensor azimuth/right-ascension bias, in degrees.
	AzRaBias param.Opt[float64] `json:"azRaBias,omitzero"`
	// Sensor azimuth/right-ascension rate bias, in degrees per second.
	AzRaRateBias param.Opt[float64] `json:"azRaRateBias,omitzero"`
	// The standard deviation of the azimuth/right ascension rate residuals, in
	// degrees, used to determine the sensor azimuth/right-ascension rate bias.
	AzRaRateSigma param.Opt[float64] `json:"azRaRateSigma,omitzero"`
	// The root mean square of the azimuth/right-ascension residuals, in degrees, used
	// to determine the sensor azimuth/right-ascension bias.
	AzRaRms param.Opt[float64] `json:"azRaRms,omitzero"`
	// The standard deviation of the azimuth/right ascension residuals, in degrees,
	// used to determine the sensor azimuth/right-ascension bias.
	AzRaSigma param.Opt[float64] `json:"azRaSigma,omitzero"`
	// Specifies the calibration reference angle set for this calibration data set.
	// Azimuth and Elevation (AZEL) or Right Ascension and Declination (RADEC).
	CalAngleRef param.Opt[string] `json:"calAngleRef,omitzero"`
	// Specifies that the calibration data are from INTRA_TRACK or INTER_TRACK
	// residuals.
	CalTrackMode param.Opt[string] `json:"calTrackMode,omitzero"`
	// The basis of calibration values contained in this record (COMPUTED,
	// OPERATIONAL).
	CalType param.Opt[string] `json:"calType,omitzero"`
	// The confidence noise bias of the duration span.
	ConfidenceNoiseBias param.Opt[float64] `json:"confidenceNoiseBias,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Duration of the sensor calibration data which produced these values, measured in
	// days.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Sensor elevation/declination acceleration bias, in degrees per second squared.
	ElDecAccelBias param.Opt[float64] `json:"elDecAccelBias,omitzero"`
	// The standard deviation of the elevation/declination acceleration residuals, in
	// degrees, used to determine the sensor elevation/declination acceleration bias.
	ElDecAccelSigma param.Opt[float64] `json:"elDecAccelSigma,omitzero"`
	// Sensor elevation/declination bias, in degrees.
	ElDecBias param.Opt[float64] `json:"elDecBias,omitzero"`
	// Sensor elevation/declination rate bias, in degrees per second.
	ElDecRateBias param.Opt[float64] `json:"elDecRateBias,omitzero"`
	// The standard deviation of the elevation/declination rate residuals, in degrees,
	// used to determine the sensor elevation/declination rate bias.
	ElDecRateSigma param.Opt[float64] `json:"elDecRateSigma,omitzero"`
	// The root mean square of the elevation/declination residuals, in degrees, used to
	// determine the sensor elevation/declination bias.
	ElDecRms param.Opt[float64] `json:"elDecRms,omitzero"`
	// The standard deviation of the elevation/declination residuals, in degrees, used
	// to determine the sensor elevation/declination bias.
	ElDecSigma param.Opt[float64] `json:"elDecSigma,omitzero"`
	// Calibration data span end time in ISO 8601 UTC format with millisecond
	// precision. If provided, the endTime must be greater than or equal to the
	// startTime in the SensorCalibration record.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The number of observables used in determining the azimuth or right-ascension
	// calibration values.
	NumAzRaObs param.Opt[int64] `json:"numAzRaObs,omitzero"`
	// The number of observables used in determining the elevation or declination
	// calibration values.
	NumElDecObs param.Opt[int64] `json:"numElDecObs,omitzero"`
	// The total number of observables available over the calibration span.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// The number of observables used in determining the photometric calibration
	// values.
	NumPhotoObs param.Opt[int64] `json:"numPhotoObs,omitzero"`
	// The number of observables used in determining the range calibration values.
	NumRangeObs param.Opt[int64] `json:"numRangeObs,omitzero"`
	// The number of observables used in determining the range rate calibration values.
	NumRangeRateObs param.Opt[int64] `json:"numRangeRateObs,omitzero"`
	// The number of observables used in determining the radar cross section (RCS)
	// calibration values.
	NumRcsObs param.Opt[int64] `json:"numRcsObs,omitzero"`
	// The number of observables used in determining the time calibration values.
	NumTimeObs param.Opt[int64] `json:"numTimeObs,omitzero"`
	// The total number of tracks available over the calibration span.
	NumTracks param.Opt[int64] `json:"numTracks,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The sensor photometric observation magnitude bias, in visual magnitude.
	PhotoBias param.Opt[float64] `json:"photoBias,omitzero"`
	// The standard deviation of the magnitude residuals, in visual magnitude, used to
	// determine the photometric bias.
	PhotoSigma param.Opt[float64] `json:"photoSigma,omitzero"`
	// Sensor range rate acceleration bias, in kilometers per second squared.
	RangeAccelBias param.Opt[float64] `json:"rangeAccelBias,omitzero"`
	// The standard deviation of the range acceleration residuals, in kilometers per
	// second squared, used to determine the sensor range acceleration bias.
	RangeAccelSigma param.Opt[float64] `json:"rangeAccelSigma,omitzero"`
	// Sensor range bias, in kilometers.
	RangeBias param.Opt[float64] `json:"rangeBias,omitzero"`
	// Sensor range rate bias, in kilometers per second.
	RangeRateBias param.Opt[float64] `json:"rangeRateBias,omitzero"`
	// The root mean square of the range rate residuals, in kilometers per second, used
	// to determine the sensor range rate bias.
	RangeRateRms param.Opt[float64] `json:"rangeRateRms,omitzero"`
	// The standard deviation of the range rate residuals, in kilometers per second,
	// used to determine the sensor range rate bias.
	RangeRateSigma param.Opt[float64] `json:"rangeRateSigma,omitzero"`
	// The root mean square of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeRms param.Opt[float64] `json:"rangeRms,omitzero"`
	// The standard deviation of the range residuals, in kilometers, used to determine
	// the sensor range bias.
	RangeSigma param.Opt[float64] `json:"rangeSigma,omitzero"`
	// The sensor radar cross section (RCS) observation bias, in square meters.
	RcsBias param.Opt[float64] `json:"rcsBias,omitzero"`
	// The standard deviation of the radar cross section residuals, in square meters,
	// used to determine the radar cross section bias.
	RcsSigma param.Opt[float64] `json:"rcsSigma,omitzero"`
	// The reference type used in the calibration.
	RefType param.Opt[string] `json:"refType,omitzero"`
	// The sensor type (MECHANICAL, OPTICAL, PHASED ARRAY, RF).
	SenType param.Opt[string] `json:"senType,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Sensor time bias, in seconds.
	TimeBias param.Opt[float64] `json:"timeBias,omitzero"`
	// The standard deviation of the time residuals, in seconds, used to determine the
	// sensor time bias.
	TimeBiasSigma param.Opt[float64] `json:"timeBiasSigma,omitzero"`
	// Three element array, expressing the sensor location in Earth Centered Rotating
	// (ECR) coordinates, in kilometers. The array element order is [x, y, z].
	Ecr []float64 `json:"ecr,omitzero"`
	// Array of the catalog IDs of the reference targets used in the calibration.
	RefTargets []string `json:"refTargets,omitzero"`
	paramObj
}

func (r SensorCalibrationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SensorCalibrationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorCalibrationUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorCalibrationUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
