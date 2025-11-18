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

// SensorCalibrationHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSensorCalibrationHistoryService] method instead.
type SensorCalibrationHistoryService struct {
	Options []option.RequestOption
}

// NewSensorCalibrationHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSensorCalibrationHistoryService(opts ...option.RequestOption) (r SensorCalibrationHistoryService) {
	r = SensorCalibrationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensorCalibrationHistoryService) List(ctx context.Context, query SensorCalibrationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SensorCalibrationHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sensorcalibration/history"
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
func (r *SensorCalibrationHistoryService) ListAutoPaging(ctx context.Context, query SensorCalibrationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SensorCalibrationHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensorCalibrationHistoryService) Aodr(ctx context.Context, query SensorCalibrationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/sensorcalibration/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SensorCalibrationHistoryService) Count(ctx context.Context, query SensorCalibrationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sensorcalibration/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The Sensor Calibration service records data about a sensor's overall accuracy
// and is used to adjust sensor settings to achieve and maintain that accuracy in
// reported sensor observations. Calibration occurs periodically when needed to
// maintain sensor accuracy or on-demand to adjust a sensor for a specific reading.
type SensorCalibrationHistoryListResponse struct {
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
	DataMode SensorCalibrationHistoryListResponseDataMode `json:"dataMode,required"`
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
func (r SensorCalibrationHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorCalibrationHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type SensorCalibrationHistoryListResponseDataMode string

const (
	SensorCalibrationHistoryListResponseDataModeReal      SensorCalibrationHistoryListResponseDataMode = "REAL"
	SensorCalibrationHistoryListResponseDataModeTest      SensorCalibrationHistoryListResponseDataMode = "TEST"
	SensorCalibrationHistoryListResponseDataModeSimulated SensorCalibrationHistoryListResponseDataMode = "SIMULATED"
	SensorCalibrationHistoryListResponseDataModeExercise  SensorCalibrationHistoryListResponseDataMode = "EXERCISE"
)

type SensorCalibrationHistoryListParams struct {
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime time.Time `query:"startTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorCalibrationHistoryListParams]'s query parameters as
// `url.Values`.
func (r SensorCalibrationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorCalibrationHistoryAodrParams struct {
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
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

// URLQuery serializes [SensorCalibrationHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r SensorCalibrationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorCalibrationHistoryCountParams struct {
	// Calibration data span start time in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorCalibrationHistoryCountParams]'s query parameters as
// `url.Values`.
func (r SensorCalibrationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
