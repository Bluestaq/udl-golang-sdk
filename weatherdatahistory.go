// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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

// WeatherDataHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeatherDataHistoryService] method instead.
type WeatherDataHistoryService struct {
	Options []option.RequestOption
}

// NewWeatherDataHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWeatherDataHistoryService(opts ...option.RequestOption) (r WeatherDataHistoryService) {
	r = WeatherDataHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *WeatherDataHistoryService) List(ctx context.Context, query WeatherDataHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WeatherDataFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/weatherdata/history"
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
func (r *WeatherDataHistoryService) ListAutoPaging(ctx context.Context, query WeatherDataHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WeatherDataFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *WeatherDataHistoryService) Aodr(ctx context.Context, query WeatherDataHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/weatherdata/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *WeatherDataHistoryService) Count(ctx context.Context, query WeatherDataHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/weatherdata/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// These services provide for posting and querying Weather Data. Weather Data
// integrates dynamic data measured by Doppler/CG such as signal power and noise
// levels, to produce useful weather information.
type WeatherDataFull struct {
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
	DataMode WeatherDataFullDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		ObTime                resp.Field
		Source                resp.Field
		ID                    resp.Field
		AngleOrientation      resp.Field
		AvgRefPwr             resp.Field
		AvgTxPwr              resp.Field
		Checksum              resp.Field
		CoIntegs              resp.Field
		ConsRecs              resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DoppVels              resp.Field
		FileCreation          resp.Field
		FirstGuessAvgs        resp.Field
		IDSensor              resp.Field
		InterpulsePeriods     resp.Field
		LightDetSensors       resp.Field
		LightEventNum         resp.Field
		NoiseLvls             resp.Field
		NumElements           resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigSensorID          resp.Field
		PosConfidence         resp.Field
		QcValue               resp.Field
		SectorNum             resp.Field
		SemiMajorAxis         resp.Field
		SemiMinorAxis         resp.Field
		SigPwrs               resp.Field
		SigStrength           resp.Field
		Snrs                  resp.Field
		SpecAvgs              resp.Field
		SpecWidths            resp.Field
		SrcIDs                resp.Field
		SrcTyps               resp.Field
		TdAvgSampleNums       resp.Field
		TermAlt               resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WeatherDataFull) RawJSON() string { return r.JSON.raw }
func (r *WeatherDataFull) UnmarshalJSON(data []byte) error {
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
type WeatherDataFullDataMode string

const (
	WeatherDataFullDataModeReal      WeatherDataFullDataMode = "REAL"
	WeatherDataFullDataModeTest      WeatherDataFullDataMode = "TEST"
	WeatherDataFullDataModeSimulated WeatherDataFullDataMode = "SIMULATED"
	WeatherDataFullDataModeExercise  WeatherDataFullDataMode = "EXERCISE"
)

type WeatherDataHistoryListParams struct {
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime time.Time `query:"obTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WeatherDataHistoryListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [WeatherDataHistoryListParams]'s query parameters as
// `url.Values`.
func (r WeatherDataHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WeatherDataHistoryAodrParams struct {
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WeatherDataHistoryAodrParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [WeatherDataHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r WeatherDataHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WeatherDataHistoryCountParams struct {
	// Datetime of the weather observation in ISO 8601 UTC datetime format with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WeatherDataHistoryCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [WeatherDataHistoryCountParams]'s query parameters as
// `url.Values`.
func (r WeatherDataHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
