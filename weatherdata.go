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

// WeatherDataService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeatherDataService] method instead.
type WeatherDataService struct {
	Options []option.RequestOption
	History WeatherDataHistoryService
}

// NewWeatherDataService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWeatherDataService(opts ...option.RequestOption) (r WeatherDataService) {
	r = WeatherDataService{}
	r.Options = opts
	r.History = NewWeatherDataHistoryService(opts...)
	return
}

// Service operation to take a single WeatherData as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *WeatherDataService) New(ctx context.Context, body WeatherDataNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/weatherdata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *WeatherDataService) List(ctx context.Context, query WeatherDataListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WeatherDataListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/weatherdata"
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
func (r *WeatherDataService) ListAutoPaging(ctx context.Context, query WeatherDataListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WeatherDataListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *WeatherDataService) Count(ctx context.Context, query WeatherDataCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/weatherdata/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple WeatherData as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *WeatherDataService) NewBulk(ctx context.Context, body WeatherDataNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/weatherdata/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single WeatherData by its unique ID passed as a path
// parameter.
func (r *WeatherDataService) Get(ctx context.Context, id string, query WeatherDataGetParams, opts ...option.RequestOption) (res *WeatherDataFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/weatherdata/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *WeatherDataService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *WeatherDataQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/weatherdata/queryhelp"
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
func (r *WeatherDataService) Tuple(ctx context.Context, query WeatherDataTupleParams, opts ...option.RequestOption) (res *[]WeatherDataFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/weatherdata/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a list of WeatherData as a POST body and ingest into
// the database. This operation is intended to be used for automated feeds into
// UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *WeatherDataService) UnvalidatedPublish(ctx context.Context, body WeatherDataUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-weatherdata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// These services provide for posting and querying Weather Data. Weather Data
// integrates dynamic data measured by Doppler/CG such as signal power and noise
// levels, to produce useful weather information.
type WeatherDataListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode WeatherDataListResponseDataMode `json:"dataMode,required"`
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Angle of orientation of the 50% positional confidence ellipse, in degrees
	// clockwise from true north.
	AngleOrientation float64 `json:"angleOrientation"`
	// Average power of the reflected signal received by the radar, in Watts.
	AvgRefPwr float64 `json:"avgRefPwr"`
	// Average transmitted power of the radar, in kilowatts.
	AvgTxPwr float64 `json:"avgTxPwr"`
	// Checksum value for the data.
	Checksum int64 `json:"checksum"`
	// Array of the number(s) of measurements used in coherent integrations used for
	// radar data processing. Users should consult the data provider for information on
	// the coherent integrations array structure.
	CoIntegs []int64 `json:"coIntegs"`
	// Array of the number(s) of records in consensus for a radar beam. Users should
	// consult the data provider for information on the consensus records array
	// structure.
	ConsRecs []int64 `json:"consRecs"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Array of full scale Nyquist Doppler velocities measured by radar, in meters per
	// second. Nyquist velocity refers to the maximum velocity magnitude that the radar
	// system can unambiguously detect. Doppler velocities with absolute values
	// exceeding the Nyquist threshold suffer from aliasing at the time of collection.
	// Users should consult the data provider for information on the doppler velocities
	// array structure.
	DoppVels []float64 `json:"doppVels"`
	// Datetime the system files were created.
	FileCreation time.Time `json:"fileCreation" format:"date-time"`
	// Array of average maximum number(s) of consecutive instances in which the same
	// first guess velocity is used in radar data processing to estimate wind speed.
	// Users should consult the data provider for information on the first guess
	// averages array structure.
	FirstGuessAvgs []int64 `json:"firstGuessAvgs"`
	// Unique identifier of the sensor making the weather measurement.
	IDSensor string `json:"idSensor"`
	// Array of the elapsed time(s) from the beginning of one pulse to the beginning of
	// the next pulse for a radar beam, in microseconds. Users should consult the data
	// provider for information on the interpulse periods array structure.
	InterpulsePeriods []float64 `json:"interpulsePeriods"`
	// Array of sensor(s) that participated in the lightning event location
	// determination.
	LightDetSensors []int64 `json:"lightDetSensors"`
	// Number of sensors used in the lightning event location solution.
	LightEventNum int64 `json:"lightEventNum"`
	// Array of noise level(s) measured by radar, in decibels. Users should consult the
	// data provider for information on the noise levels array structure.
	NoiseLvls []float64 `json:"noiseLvls"`
	// Number of antennas across all sectors within the radar coverage area.
	NumElements int64 `json:"numElements"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the record source. This may be an internal
	// identifier and not necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// The positional confidence of the calculated lightning event location using the
	// chi-square statistical method.
	PosConfidence float64 `json:"posConfidence"`
	// Quality control flag value, as defined by the data provider.
	QcValue int64 `json:"qcValue"`
	// Number of sectors within the radar coverage area, each containing a number of
	// antennas.
	SectorNum int64 `json:"sectorNum"`
	// Semi-major axis of the 50% positional confidence ellipse, in kilometers.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// Semi-minor axis of the 50% positional confidence ellipse, in kilometers.
	SemiMinorAxis float64 `json:"semiMinorAxis"`
	// Array of signal power(s) measured by the sensor, in decibels. Users should
	// consult the data provider for information on the signal powers array structure.
	SigPwrs []float64 `json:"sigPwrs"`
	// Signal strength of the electromagnetic energy received due to a lightning event,
	// in kiloamps.
	SigStrength float64 `json:"sigStrength"`
	// Array of signal to noise ratio(s) for a radar beam, in decibels. Users should
	// consult the data provider for information on the signal to noise ratios array
	// structure.
	Snrs []float64 `json:"snrs"`
	// Array of the number(s) of spectral averages used in radar data processing. Users
	// should consult the data provider for information on the spectral averages array
	// structure.
	SpecAvgs []int64 `json:"specAvgs"`
	// Array of width(s) of the distribution in Doppler velocity measured by radar, in
	// meters/second. Spectral width depends on the particle size distribution, the
	// wind shear across the radar beam, and turbulence. Users should consult the data
	// provider for information on the spectral widths array structure.
	SpecWidths []float64 `json:"specWidths"`
	// Array of UUID(s) of the UDL data record(s) that are related to this WeatherData
	// record. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (SENSOR, WEATHERREPORT) that are related to this
	// WeatherData record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// Array of the number(s) of radar samples used in time domain averaging for radar
	// data processing. Time domain averaging improves the quality of the measured
	// signal by reducing random noise and enhancing the signal-to-noise ratio. Users
	// should consult the data provider for information on the time domain sample
	// numbers array structure.
	TdAvgSampleNums []int64 `json:"tdAvgSampleNums"`
	// Last altitude with recorded measurements in this record, in meters.
	TermAlt float64 `json:"termAlt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AngleOrientation      respjson.Field
		AvgRefPwr             respjson.Field
		AvgTxPwr              respjson.Field
		Checksum              respjson.Field
		CoIntegs              respjson.Field
		ConsRecs              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DoppVels              respjson.Field
		FileCreation          respjson.Field
		FirstGuessAvgs        respjson.Field
		IDSensor              respjson.Field
		InterpulsePeriods     respjson.Field
		LightDetSensors       respjson.Field
		LightEventNum         respjson.Field
		NoiseLvls             respjson.Field
		NumElements           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		PosConfidence         respjson.Field
		QcValue               respjson.Field
		SectorNum             respjson.Field
		SemiMajorAxis         respjson.Field
		SemiMinorAxis         respjson.Field
		SigPwrs               respjson.Field
		SigStrength           respjson.Field
		Snrs                  respjson.Field
		SpecAvgs              respjson.Field
		SpecWidths            respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		TdAvgSampleNums       respjson.Field
		TermAlt               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WeatherDataListResponse) RawJSON() string { return r.JSON.raw }
func (r *WeatherDataListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type WeatherDataListResponseDataMode string

const (
	WeatherDataListResponseDataModeReal      WeatherDataListResponseDataMode = "REAL"
	WeatherDataListResponseDataModeTest      WeatherDataListResponseDataMode = "TEST"
	WeatherDataListResponseDataModeSimulated WeatherDataListResponseDataMode = "SIMULATED"
	WeatherDataListResponseDataModeExercise  WeatherDataListResponseDataMode = "EXERCISE"
)

type WeatherDataQueryhelpResponse struct {
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
func (r WeatherDataQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *WeatherDataQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WeatherDataNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode WeatherDataNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Angle of orientation of the 50% positional confidence ellipse, in degrees
	// clockwise from true north.
	AngleOrientation param.Opt[float64] `json:"angleOrientation,omitzero"`
	// Average power of the reflected signal received by the radar, in Watts.
	AvgRefPwr param.Opt[float64] `json:"avgRefPwr,omitzero"`
	// Average transmitted power of the radar, in kilowatts.
	AvgTxPwr param.Opt[float64] `json:"avgTxPwr,omitzero"`
	// Checksum value for the data.
	Checksum param.Opt[int64] `json:"checksum,omitzero"`
	// Datetime the system files were created.
	FileCreation param.Opt[time.Time] `json:"fileCreation,omitzero" format:"date-time"`
	// Unique identifier of the sensor making the weather measurement.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Number of sensors used in the lightning event location solution.
	LightEventNum param.Opt[int64] `json:"lightEventNum,omitzero"`
	// Number of antennas across all sectors within the radar coverage area.
	NumElements param.Opt[int64] `json:"numElements,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source. This may be an internal
	// identifier and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// The positional confidence of the calculated lightning event location using the
	// chi-square statistical method.
	PosConfidence param.Opt[float64] `json:"posConfidence,omitzero"`
	// Quality control flag value, as defined by the data provider.
	QcValue param.Opt[int64] `json:"qcValue,omitzero"`
	// Number of sectors within the radar coverage area, each containing a number of
	// antennas.
	SectorNum param.Opt[int64] `json:"sectorNum,omitzero"`
	// Semi-major axis of the 50% positional confidence ellipse, in kilometers.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	// Semi-minor axis of the 50% positional confidence ellipse, in kilometers.
	SemiMinorAxis param.Opt[float64] `json:"semiMinorAxis,omitzero"`
	// Signal strength of the electromagnetic energy received due to a lightning event,
	// in kiloamps.
	SigStrength param.Opt[float64] `json:"sigStrength,omitzero"`
	// Last altitude with recorded measurements in this record, in meters.
	TermAlt param.Opt[float64] `json:"termAlt,omitzero"`
	// Array of the number(s) of measurements used in coherent integrations used for
	// radar data processing. Users should consult the data provider for information on
	// the coherent integrations array structure.
	CoIntegs []int64 `json:"coIntegs,omitzero"`
	// Array of the number(s) of records in consensus for a radar beam. Users should
	// consult the data provider for information on the consensus records array
	// structure.
	ConsRecs []int64 `json:"consRecs,omitzero"`
	// Array of full scale Nyquist Doppler velocities measured by radar, in meters per
	// second. Nyquist velocity refers to the maximum velocity magnitude that the radar
	// system can unambiguously detect. Doppler velocities with absolute values
	// exceeding the Nyquist threshold suffer from aliasing at the time of collection.
	// Users should consult the data provider for information on the doppler velocities
	// array structure.
	DoppVels []float64 `json:"doppVels,omitzero"`
	// Array of average maximum number(s) of consecutive instances in which the same
	// first guess velocity is used in radar data processing to estimate wind speed.
	// Users should consult the data provider for information on the first guess
	// averages array structure.
	FirstGuessAvgs []int64 `json:"firstGuessAvgs,omitzero"`
	// Array of the elapsed time(s) from the beginning of one pulse to the beginning of
	// the next pulse for a radar beam, in microseconds. Users should consult the data
	// provider for information on the interpulse periods array structure.
	InterpulsePeriods []float64 `json:"interpulsePeriods,omitzero"`
	// Array of sensor(s) that participated in the lightning event location
	// determination.
	LightDetSensors []int64 `json:"lightDetSensors,omitzero"`
	// Array of noise level(s) measured by radar, in decibels. Users should consult the
	// data provider for information on the noise levels array structure.
	NoiseLvls []float64 `json:"noiseLvls,omitzero"`
	// Array of signal power(s) measured by the sensor, in decibels. Users should
	// consult the data provider for information on the signal powers array structure.
	SigPwrs []float64 `json:"sigPwrs,omitzero"`
	// Array of signal to noise ratio(s) for a radar beam, in decibels. Users should
	// consult the data provider for information on the signal to noise ratios array
	// structure.
	Snrs []float64 `json:"snrs,omitzero"`
	// Array of the number(s) of spectral averages used in radar data processing. Users
	// should consult the data provider for information on the spectral averages array
	// structure.
	SpecAvgs []int64 `json:"specAvgs,omitzero"`
	// Array of width(s) of the distribution in Doppler velocity measured by radar, in
	// meters/second. Spectral width depends on the particle size distribution, the
	// wind shear across the radar beam, and turbulence. Users should consult the data
	// provider for information on the spectral widths array structure.
	SpecWidths []float64 `json:"specWidths,omitzero"`
	// Array of UUID(s) of the UDL data record(s) that are related to this WeatherData
	// record. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (SENSOR, WEATHERREPORT) that are related to this
	// WeatherData record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Array of the number(s) of radar samples used in time domain averaging for radar
	// data processing. Time domain averaging improves the quality of the measured
	// signal by reducing random noise and enhancing the signal-to-noise ratio. Users
	// should consult the data provider for information on the time domain sample
	// numbers array structure.
	TdAvgSampleNums []int64 `json:"tdAvgSampleNums,omitzero"`
	paramObj
}

func (r WeatherDataNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WeatherDataNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WeatherDataNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type WeatherDataNewParamsDataMode string

const (
	WeatherDataNewParamsDataModeReal      WeatherDataNewParamsDataMode = "REAL"
	WeatherDataNewParamsDataModeTest      WeatherDataNewParamsDataMode = "TEST"
	WeatherDataNewParamsDataModeSimulated WeatherDataNewParamsDataMode = "SIMULATED"
	WeatherDataNewParamsDataModeExercise  WeatherDataNewParamsDataMode = "EXERCISE"
)

type WeatherDataListParams struct {
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WeatherDataListParams]'s query parameters as `url.Values`.
func (r WeatherDataListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WeatherDataCountParams struct {
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WeatherDataCountParams]'s query parameters as `url.Values`.
func (r WeatherDataCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WeatherDataNewBulkParams struct {
	Body []WeatherDataNewBulkParamsBody
	paramObj
}

func (r WeatherDataNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *WeatherDataNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide for posting and querying Weather Data. Weather Data
// integrates dynamic data measured by Doppler/CG such as signal power and noise
// levels, to produce useful weather information.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type WeatherDataNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,omitzero,required"`
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Angle of orientation of the 50% positional confidence ellipse, in degrees
	// clockwise from true north.
	AngleOrientation param.Opt[float64] `json:"angleOrientation,omitzero"`
	// Average power of the reflected signal received by the radar, in Watts.
	AvgRefPwr param.Opt[float64] `json:"avgRefPwr,omitzero"`
	// Average transmitted power of the radar, in kilowatts.
	AvgTxPwr param.Opt[float64] `json:"avgTxPwr,omitzero"`
	// Checksum value for the data.
	Checksum param.Opt[int64] `json:"checksum,omitzero"`
	// Datetime the system files were created.
	FileCreation param.Opt[time.Time] `json:"fileCreation,omitzero" format:"date-time"`
	// Unique identifier of the sensor making the weather measurement.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Number of sensors used in the lightning event location solution.
	LightEventNum param.Opt[int64] `json:"lightEventNum,omitzero"`
	// Number of antennas across all sectors within the radar coverage area.
	NumElements param.Opt[int64] `json:"numElements,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source. This may be an internal
	// identifier and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// The positional confidence of the calculated lightning event location using the
	// chi-square statistical method.
	PosConfidence param.Opt[float64] `json:"posConfidence,omitzero"`
	// Quality control flag value, as defined by the data provider.
	QcValue param.Opt[int64] `json:"qcValue,omitzero"`
	// Number of sectors within the radar coverage area, each containing a number of
	// antennas.
	SectorNum param.Opt[int64] `json:"sectorNum,omitzero"`
	// Semi-major axis of the 50% positional confidence ellipse, in kilometers.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	// Semi-minor axis of the 50% positional confidence ellipse, in kilometers.
	SemiMinorAxis param.Opt[float64] `json:"semiMinorAxis,omitzero"`
	// Signal strength of the electromagnetic energy received due to a lightning event,
	// in kiloamps.
	SigStrength param.Opt[float64] `json:"sigStrength,omitzero"`
	// Last altitude with recorded measurements in this record, in meters.
	TermAlt param.Opt[float64] `json:"termAlt,omitzero"`
	// Array of the number(s) of measurements used in coherent integrations used for
	// radar data processing. Users should consult the data provider for information on
	// the coherent integrations array structure.
	CoIntegs []int64 `json:"coIntegs,omitzero"`
	// Array of the number(s) of records in consensus for a radar beam. Users should
	// consult the data provider for information on the consensus records array
	// structure.
	ConsRecs []int64 `json:"consRecs,omitzero"`
	// Array of full scale Nyquist Doppler velocities measured by radar, in meters per
	// second. Nyquist velocity refers to the maximum velocity magnitude that the radar
	// system can unambiguously detect. Doppler velocities with absolute values
	// exceeding the Nyquist threshold suffer from aliasing at the time of collection.
	// Users should consult the data provider for information on the doppler velocities
	// array structure.
	DoppVels []float64 `json:"doppVels,omitzero"`
	// Array of average maximum number(s) of consecutive instances in which the same
	// first guess velocity is used in radar data processing to estimate wind speed.
	// Users should consult the data provider for information on the first guess
	// averages array structure.
	FirstGuessAvgs []int64 `json:"firstGuessAvgs,omitzero"`
	// Array of the elapsed time(s) from the beginning of one pulse to the beginning of
	// the next pulse for a radar beam, in microseconds. Users should consult the data
	// provider for information on the interpulse periods array structure.
	InterpulsePeriods []float64 `json:"interpulsePeriods,omitzero"`
	// Array of sensor(s) that participated in the lightning event location
	// determination.
	LightDetSensors []int64 `json:"lightDetSensors,omitzero"`
	// Array of noise level(s) measured by radar, in decibels. Users should consult the
	// data provider for information on the noise levels array structure.
	NoiseLvls []float64 `json:"noiseLvls,omitzero"`
	// Array of signal power(s) measured by the sensor, in decibels. Users should
	// consult the data provider for information on the signal powers array structure.
	SigPwrs []float64 `json:"sigPwrs,omitzero"`
	// Array of signal to noise ratio(s) for a radar beam, in decibels. Users should
	// consult the data provider for information on the signal to noise ratios array
	// structure.
	Snrs []float64 `json:"snrs,omitzero"`
	// Array of the number(s) of spectral averages used in radar data processing. Users
	// should consult the data provider for information on the spectral averages array
	// structure.
	SpecAvgs []int64 `json:"specAvgs,omitzero"`
	// Array of width(s) of the distribution in Doppler velocity measured by radar, in
	// meters/second. Spectral width depends on the particle size distribution, the
	// wind shear across the radar beam, and turbulence. Users should consult the data
	// provider for information on the spectral widths array structure.
	SpecWidths []float64 `json:"specWidths,omitzero"`
	// Array of UUID(s) of the UDL data record(s) that are related to this WeatherData
	// record. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (SENSOR, WEATHERREPORT) that are related to this
	// WeatherData record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Array of the number(s) of radar samples used in time domain averaging for radar
	// data processing. Time domain averaging improves the quality of the measured
	// signal by reducing random noise and enhancing the signal-to-noise ratio. Users
	// should consult the data provider for information on the time domain sample
	// numbers array structure.
	TdAvgSampleNums []int64 `json:"tdAvgSampleNums,omitzero"`
	paramObj
}

func (r WeatherDataNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow WeatherDataNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WeatherDataNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WeatherDataNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type WeatherDataGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WeatherDataGetParams]'s query parameters as `url.Values`.
func (r WeatherDataGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WeatherDataTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WeatherDataTupleParams]'s query parameters as `url.Values`.
func (r WeatherDataTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WeatherDataUnvalidatedPublishParams struct {
	Body []WeatherDataUnvalidatedPublishParamsBody
	paramObj
}

func (r WeatherDataUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *WeatherDataUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide for posting and querying Weather Data. Weather Data
// integrates dynamic data measured by Doppler/CG such as signal power and noise
// levels, to produce useful weather information.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type WeatherDataUnvalidatedPublishParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,omitzero,required"`
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Angle of orientation of the 50% positional confidence ellipse, in degrees
	// clockwise from true north.
	AngleOrientation param.Opt[float64] `json:"angleOrientation,omitzero"`
	// Average power of the reflected signal received by the radar, in Watts.
	AvgRefPwr param.Opt[float64] `json:"avgRefPwr,omitzero"`
	// Average transmitted power of the radar, in kilowatts.
	AvgTxPwr param.Opt[float64] `json:"avgTxPwr,omitzero"`
	// Checksum value for the data.
	Checksum param.Opt[int64] `json:"checksum,omitzero"`
	// Datetime the system files were created.
	FileCreation param.Opt[time.Time] `json:"fileCreation,omitzero" format:"date-time"`
	// Unique identifier of the sensor making the weather measurement.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Number of sensors used in the lightning event location solution.
	LightEventNum param.Opt[int64] `json:"lightEventNum,omitzero"`
	// Number of antennas across all sectors within the radar coverage area.
	NumElements param.Opt[int64] `json:"numElements,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source. This may be an internal
	// identifier and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// The positional confidence of the calculated lightning event location using the
	// chi-square statistical method.
	PosConfidence param.Opt[float64] `json:"posConfidence,omitzero"`
	// Quality control flag value, as defined by the data provider.
	QcValue param.Opt[int64] `json:"qcValue,omitzero"`
	// Number of sectors within the radar coverage area, each containing a number of
	// antennas.
	SectorNum param.Opt[int64] `json:"sectorNum,omitzero"`
	// Semi-major axis of the 50% positional confidence ellipse, in kilometers.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	// Semi-minor axis of the 50% positional confidence ellipse, in kilometers.
	SemiMinorAxis param.Opt[float64] `json:"semiMinorAxis,omitzero"`
	// Signal strength of the electromagnetic energy received due to a lightning event,
	// in kiloamps.
	SigStrength param.Opt[float64] `json:"sigStrength,omitzero"`
	// Last altitude with recorded measurements in this record, in meters.
	TermAlt param.Opt[float64] `json:"termAlt,omitzero"`
	// Array of the number(s) of measurements used in coherent integrations used for
	// radar data processing. Users should consult the data provider for information on
	// the coherent integrations array structure.
	CoIntegs []int64 `json:"coIntegs,omitzero"`
	// Array of the number(s) of records in consensus for a radar beam. Users should
	// consult the data provider for information on the consensus records array
	// structure.
	ConsRecs []int64 `json:"consRecs,omitzero"`
	// Array of full scale Nyquist Doppler velocities measured by radar, in meters per
	// second. Nyquist velocity refers to the maximum velocity magnitude that the radar
	// system can unambiguously detect. Doppler velocities with absolute values
	// exceeding the Nyquist threshold suffer from aliasing at the time of collection.
	// Users should consult the data provider for information on the doppler velocities
	// array structure.
	DoppVels []float64 `json:"doppVels,omitzero"`
	// Array of average maximum number(s) of consecutive instances in which the same
	// first guess velocity is used in radar data processing to estimate wind speed.
	// Users should consult the data provider for information on the first guess
	// averages array structure.
	FirstGuessAvgs []int64 `json:"firstGuessAvgs,omitzero"`
	// Array of the elapsed time(s) from the beginning of one pulse to the beginning of
	// the next pulse for a radar beam, in microseconds. Users should consult the data
	// provider for information on the interpulse periods array structure.
	InterpulsePeriods []float64 `json:"interpulsePeriods,omitzero"`
	// Array of sensor(s) that participated in the lightning event location
	// determination.
	LightDetSensors []int64 `json:"lightDetSensors,omitzero"`
	// Array of noise level(s) measured by radar, in decibels. Users should consult the
	// data provider for information on the noise levels array structure.
	NoiseLvls []float64 `json:"noiseLvls,omitzero"`
	// Array of signal power(s) measured by the sensor, in decibels. Users should
	// consult the data provider for information on the signal powers array structure.
	SigPwrs []float64 `json:"sigPwrs,omitzero"`
	// Array of signal to noise ratio(s) for a radar beam, in decibels. Users should
	// consult the data provider for information on the signal to noise ratios array
	// structure.
	Snrs []float64 `json:"snrs,omitzero"`
	// Array of the number(s) of spectral averages used in radar data processing. Users
	// should consult the data provider for information on the spectral averages array
	// structure.
	SpecAvgs []int64 `json:"specAvgs,omitzero"`
	// Array of width(s) of the distribution in Doppler velocity measured by radar, in
	// meters/second. Spectral width depends on the particle size distribution, the
	// wind shear across the radar beam, and turbulence. Users should consult the data
	// provider for information on the spectral widths array structure.
	SpecWidths []float64 `json:"specWidths,omitzero"`
	// Array of UUID(s) of the UDL data record(s) that are related to this WeatherData
	// record. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (SENSOR, WEATHERREPORT) that are related to this
	// WeatherData record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Array of the number(s) of radar samples used in time domain averaging for radar
	// data processing. Time domain averaging improves the quality of the measured
	// signal by reducing random noise and enhancing the signal-to-noise ratio. Users
	// should consult the data provider for information on the time domain sample
	// numbers array structure.
	TdAvgSampleNums []int64 `json:"tdAvgSampleNums,omitzero"`
	paramObj
}

func (r WeatherDataUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow WeatherDataUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WeatherDataUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WeatherDataUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
