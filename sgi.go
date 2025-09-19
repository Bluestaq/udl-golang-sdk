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

// SgiService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSgiService] method instead.
type SgiService struct {
	Options []option.RequestOption
	History SgiHistoryService
}

// NewSgiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSgiService(opts ...option.RequestOption) (r SgiService) {
	r = SgiService{}
	r.Options = opts
	r.History = NewSgiHistoryService(opts...)
	return
}

// Service operation to take a single SGI record as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *SgiService) New(ctx context.Context, body SgiNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sgi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single SGI record. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *SgiService) Update(ctx context.Context, id string, body SgiUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sgi/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SgiService) List(ctx context.Context, query SgiListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SgiListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sgi"
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
func (r *SgiService) ListAutoPaging(ctx context.Context, query SgiListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SgiListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a SGI record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *SgiService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sgi/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SgiService) Count(ctx context.Context, query SgiCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sgi/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of SGI
// as a POST body and ingest into the database. This operation is not intended to
// be used for automated feeds into UDL. Data providers should contact the UDL team
// for specific role assignments and for instructions on setting up a permanent
// feed through an alternate mechanism.
func (r *SgiService) NewBulk(ctx context.Context, body SgiNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sgi/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single SGI record by its unique ID passed as a path
// parameter.
func (r *SgiService) Get(ctx context.Context, id string, query SgiGetParams, opts ...option.RequestOption) (res *SgiGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sgi/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service to return matching SGI records as of the effective date.
func (r *SgiService) GetDataByEffectiveAsOfDate(ctx context.Context, query SgiGetDataByEffectiveAsOfDateParams, opts ...option.RequestOption) (res *SgiGetDataByEffectiveAsOfDateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/sgi/getSGIDataByEffectiveAsOfDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SgiService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SgiQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/sgi/queryhelp"
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
func (r *SgiService) Tuple(ctx context.Context, query SgiTupleParams, opts ...option.RequestOption) (res *[]SgiTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/sgi/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple SGI as a POST body and ingest into the
// database. This operation is intended to be used for automated feeds into UDL. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *SgiService) UnvalidatedPublish(ctx context.Context, body SgiUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-sgi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Model representation of space weather/solar, geomagnetic, and radiation belt
// indices.
type SgiListResponse struct {
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
	DataMode SgiListResponseDataMode `json:"dataMode,required"`
	// ISO8601 UTC Time the data was received and processed from the source. Typically
	// a source provides data for a date window with each transmission including past,
	// present, and future predicted values.
	EffectiveDate time.Time `json:"effectiveDate,required" format:"date-time"`
	// ISO8601 UTC Time of the index value. This could be a past, current, or future
	// predicted value. Note: sgiDate defines the start time of the time window for
	// this data record.
	SgiDate time.Time `json:"sgiDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Signal analyzer's input attenuation level, in decibels. Attenuation is a setting
	// on the hardware that measures the power of a signal.
	AnalyzerAttenuation float64 `json:"analyzerAttenuation"`
	// Ap is the planetary geomagnetic 2 nT index (00-21 UT) for the timespan specified
	// in apDuration. If apDuration is null, a 3 hour duration should be assumed.
	Ap float64 `json:"ap"`
	// The time, in hours, for which the Ap index value is valid. If null, a span of 3
	// hours is assumed.
	ApDuration int64 `json:"apDuration"`
	// Array containing the degree of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffDegree []int64 `json:"coeffDegree"`
	// Array containing the order of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffOrder []int64 `json:"coeffOrder"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array containing the cosine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctce []float64 `json:"ctce"`
	// Array containing the cosine spherical-harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctci []float64 `json:"ctci"`
	// Disturbance Storm Time geomagnetic index in nT.
	Dst float64 `json:"dst"`
	// delta exospheric temperature correction in units of K.
	Dtc float64 `json:"dtc"`
	// Extreme Ultraviolet (EUV) proxy, E10.7, in x10-22 Watts per meter squared per
	// Hertz, is the integrated solar EUV energy flux at the top of atmosphere and
	// normalized to solar flux units.
	E10 float64 `json:"e10"`
	// E54 (E10-Bar), in x10-22 Watts per meter squared per Hertz, uses the past
	// 54-days E10 values to determine the E10 average.
	E54 float64 `json:"e54"`
	// Daily solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F10 float64 `json:"f10"`
	// Daily F10.7 index - high range, in x10-22 Watts per meter squared per Hertz.
	// This field usually applies to forecast values, based on the consensus of the
	// Solar Cycle 24 Prediction Panel.
	F10High float64 `json:"f10High"`
	// Daily F10.7 index - low range, in x10-22 Watts per meter squared per Hertz. This
	// field usually applies to forecast values, based on the consensus of the Solar
	// Cycle 24 Prediction Panel.
	F10Low float64 `json:"f10Low"`
	// 54 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F54 float64 `json:"f54"`
	// 81 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F81 float64 `json:"f81"`
	// The maximum measured gamma deflection during the kpDuration timespan. If
	// kpDuration is null, a 3 hour duration should be assumed.
	Gamma int64 `json:"gamma"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The maximum measured K-Index at the associated station during the kpDuration
	// timespan. The K-Index is a unitless measure (0 - 9) of the deviation in the
	// Earth's magnetic field from normal at the station geographic location, with 0
	// indicating the absence of geomagnetic disturbance, and 9 indicating the most
	// significant disturbance. If kpDuration is null, a 3 hour duration should be
	// assumed.
	KIndex int64 `json:"kIndex"`
	// The Planetary K-index (Kp) over the kpDuration timespan. The Kp-Index is the
	// average K-Index for the entire Earth, utilizing a unitless scale (0-9, in
	// incremenets of 1/3), with 0 indicating the absence of geomagnetic disturbance,
	// and 9 indicating the most significant disturbance. If kpDuration is null, a 3
	// hour duration should be assumed.
	Kp float64 `json:"kp"`
	// The time, in hours, over which the K, Kp, and/or gamma index values are
	// measured. If null, a span of 3 hours is assumed.
	KpDuration int64 `json:"kpDuration"`
	// Daily M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M10 float64 `json:"m10"`
	// 54 day M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M54 float64 `json:"m54"`
	// The transmitted DCA mode of the record (1-3).
	Mode int64 `json:"mode"`
	// The normalization factor that has already been applied to the index value prior
	// to record ingest. Typically used to normalize the index value to a particular
	// interval. Units of the normalization factor may vary depending on the provider
	// of this data (REACH, POES, CEASE3, etc.).
	NormFactor float64 `json:"normFactor"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the reporting source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
	// Y = Emergency, Z = Flash).
	//
	// Any of "O", "P", "R", "Y", "Z".
	Precedence SgiListResponsePrecedence `json:"precedence"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The timespan over which the associated radiation belt index is factored. If
	// rbDuration is null, a 24 hour duration should be assumed. Note: rbDuration
	// defines the length of the time window for this data record. The time window
	// start time is defined by sgiDate, and the time window end time is defined by
	// sgiDate plus rbDuration.
	RbDuration int64 `json:"rbDuration"`
	// The value of the radiation belt index. This is the ratio of current intensity of
	// a radiation belt to long-term average value. It's long-term average should be
	// close to 1. Depending on the type of belt sensor, this ratio may measure Flux
	// (number of particles / (cm^2 sec energy solid-angle)), dose rate (rad per
	// second), or relative counts of particles per time (counts per second). The index
	// value may also be normalized, the normalization value typically represents an
	// average of the sensor measurements taken within a region over a given time
	// interval. See the normFactor field for the specific normalization factor, if
	// provided.
	RbIndex float64 `json:"rbIndex"`
	// Region code for the associated radiation belt index. This is the code associated
	// with the corresponding radiation belt location. See the provider card for
	// reference to specific region code definitions.
	RbRegionCode int64 `json:"rbRegionCode"`
	// Daily S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S10 float64 `json:"s10"`
	// 54 day S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S54 float64 `json:"s54"`
	// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
	// record.
	//
	// Any of "I", "N", "P".
	State SgiListResponseState `json:"state"`
	// The name/location of the station that collected the geomagnetic data for this
	// record.
	StationName string `json:"stationName"`
	// Array containing the sine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stce []float64 `json:"stce"`
	// Array containing the sine spherical harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stci []float64 `json:"stci"`
	// Daily sunspot number.
	SunspotNum float64 `json:"sunspotNum"`
	// Daily sunspot number - high range. This field usually applies to forecast
	// values, based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumHigh float64 `json:"sunspotNumHigh"`
	// Daily sunspot number - low range. This field usually applies to forecast values,
	// based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumLow float64 `json:"sunspotNumLow"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// The type of data contained in this record (e.g. HASDM, JBH09, K-Index, PSD-dB,
	// RBI, RFI-SFU, etc).
	Type string `json:"type"`
	// Daily Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y10 float64 `json:"y10"`
	// 54 day Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y54 float64 `json:"y54"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EffectiveDate         respjson.Field
		SgiDate               respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AnalyzerAttenuation   respjson.Field
		Ap                    respjson.Field
		ApDuration            respjson.Field
		CoeffDegree           respjson.Field
		CoeffOrder            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Ctce                  respjson.Field
		Ctci                  respjson.Field
		Dst                   respjson.Field
		Dtc                   respjson.Field
		E10                   respjson.Field
		E54                   respjson.Field
		F10                   respjson.Field
		F10High               respjson.Field
		F10Low                respjson.Field
		F54                   respjson.Field
		F81                   respjson.Field
		Gamma                 respjson.Field
		IDSensor              respjson.Field
		KIndex                respjson.Field
		Kp                    respjson.Field
		KpDuration            respjson.Field
		M10                   respjson.Field
		M54                   respjson.Field
		Mode                  respjson.Field
		NormFactor            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		Precedence            respjson.Field
		RawFileUri            respjson.Field
		RbDuration            respjson.Field
		RbIndex               respjson.Field
		RbRegionCode          respjson.Field
		S10                   respjson.Field
		S54                   respjson.Field
		State                 respjson.Field
		StationName           respjson.Field
		Stce                  respjson.Field
		Stci                  respjson.Field
		SunspotNum            respjson.Field
		SunspotNumHigh        respjson.Field
		SunspotNumLow         respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		Y10                   respjson.Field
		Y54                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SgiListResponse) RawJSON() string { return r.JSON.raw }
func (r *SgiListResponse) UnmarshalJSON(data []byte) error {
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
type SgiListResponseDataMode string

const (
	SgiListResponseDataModeReal      SgiListResponseDataMode = "REAL"
	SgiListResponseDataModeTest      SgiListResponseDataMode = "TEST"
	SgiListResponseDataModeSimulated SgiListResponseDataMode = "SIMULATED"
	SgiListResponseDataModeExercise  SgiListResponseDataMode = "EXERCISE"
)

// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
// Y = Emergency, Z = Flash).
type SgiListResponsePrecedence string

const (
	SgiListResponsePrecedenceO SgiListResponsePrecedence = "O"
	SgiListResponsePrecedenceP SgiListResponsePrecedence = "P"
	SgiListResponsePrecedenceR SgiListResponsePrecedence = "R"
	SgiListResponsePrecedenceY SgiListResponsePrecedence = "Y"
	SgiListResponsePrecedenceZ SgiListResponsePrecedence = "Z"
)

// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
// record.
type SgiListResponseState string

const (
	SgiListResponseStateI SgiListResponseState = "I"
	SgiListResponseStateN SgiListResponseState = "N"
	SgiListResponseStateP SgiListResponseState = "P"
)

// Model representation of space weather/solar, geomagnetic, and radiation belt
// indices.
type SgiGetResponse struct {
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
	DataMode SgiGetResponseDataMode `json:"dataMode,required"`
	// ISO8601 UTC Time the data was received and processed from the source. Typically
	// a source provides data for a date window with each transmission including past,
	// present, and future predicted values.
	EffectiveDate time.Time `json:"effectiveDate,required" format:"date-time"`
	// ISO8601 UTC Time of the index value. This could be a past, current, or future
	// predicted value. Note: sgiDate defines the start time of the time window for
	// this data record.
	SgiDate time.Time `json:"sgiDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Signal analyzer's input attenuation level, in decibels. Attenuation is a setting
	// on the hardware that measures the power of a signal.
	AnalyzerAttenuation float64 `json:"analyzerAttenuation"`
	// Ap is the planetary geomagnetic 2 nT index (00-21 UT) for the timespan specified
	// in apDuration. If apDuration is null, a 3 hour duration should be assumed.
	Ap float64 `json:"ap"`
	// The time, in hours, for which the Ap index value is valid. If null, a span of 3
	// hours is assumed.
	ApDuration int64 `json:"apDuration"`
	// Array containing the degree of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffDegree []int64 `json:"coeffDegree"`
	// Array containing the order of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffOrder []int64 `json:"coeffOrder"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array containing the cosine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctce []float64 `json:"ctce"`
	// Array containing the cosine spherical-harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctci []float64 `json:"ctci"`
	// Disturbance Storm Time geomagnetic index in nT.
	Dst float64 `json:"dst"`
	// delta exospheric temperature correction in units of K.
	Dtc float64 `json:"dtc"`
	// Extreme Ultraviolet (EUV) proxy, E10.7, in x10-22 Watts per meter squared per
	// Hertz, is the integrated solar EUV energy flux at the top of atmosphere and
	// normalized to solar flux units.
	E10 float64 `json:"e10"`
	// E54 (E10-Bar), in x10-22 Watts per meter squared per Hertz, uses the past
	// 54-days E10 values to determine the E10 average.
	E54 float64 `json:"e54"`
	// Daily solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F10 float64 `json:"f10"`
	// Daily F10.7 index - high range, in x10-22 Watts per meter squared per Hertz.
	// This field usually applies to forecast values, based on the consensus of the
	// Solar Cycle 24 Prediction Panel.
	F10High float64 `json:"f10High"`
	// Daily F10.7 index - low range, in x10-22 Watts per meter squared per Hertz. This
	// field usually applies to forecast values, based on the consensus of the Solar
	// Cycle 24 Prediction Panel.
	F10Low float64 `json:"f10Low"`
	// 54 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F54 float64 `json:"f54"`
	// 81 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F81 float64 `json:"f81"`
	// Array of individual power spectral density (PSD) frequencies of the signal, in
	// megahertz. This array should correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies"`
	// The maximum measured gamma deflection during the kpDuration timespan. If
	// kpDuration is null, a 3 hour duration should be assumed.
	Gamma int64 `json:"gamma"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The maximum measured K-Index at the associated station during the kpDuration
	// timespan. The K-Index is a unitless measure (0 - 9) of the deviation in the
	// Earth's magnetic field from normal at the station geographic location, with 0
	// indicating the absence of geomagnetic disturbance, and 9 indicating the most
	// significant disturbance. If kpDuration is null, a 3 hour duration should be
	// assumed.
	KIndex int64 `json:"kIndex"`
	// The Planetary K-index (Kp) over the kpDuration timespan. The Kp-Index is the
	// average K-Index for the entire Earth, utilizing a unitless scale (0-9, in
	// incremenets of 1/3), with 0 indicating the absence of geomagnetic disturbance,
	// and 9 indicating the most significant disturbance. If kpDuration is null, a 3
	// hour duration should be assumed.
	Kp float64 `json:"kp"`
	// The time, in hours, over which the K, Kp, and/or gamma index values are
	// measured. If null, a span of 3 hours is assumed.
	KpDuration int64 `json:"kpDuration"`
	// Daily M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M10 float64 `json:"m10"`
	// 54 day M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M54 float64 `json:"m54"`
	// The transmitted DCA mode of the record (1-3).
	Mode int64 `json:"mode"`
	// The normalization factor that has already been applied to the index value prior
	// to record ingest. Typically used to normalize the index value to a particular
	// interval. Units of the normalization factor may vary depending on the provider
	// of this data (REACH, POES, CEASE3, etc.).
	NormFactor float64 `json:"normFactor"`
	// Observed baseline values of the frequencies specified in the frequencies field,
	// in solar flux units. The baseline values will be used to help detect abnormal
	// readings from the sun that might indicate a flare or other solar activity.
	ObservedBaseline []int64 `json:"observedBaseline"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the reporting source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Array of individual power spectral density (PSD) powers of the signal, in watts.
	// This array should correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers"`
	// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
	// Y = Emergency, Z = Flash).
	//
	// Any of "O", "P", "R", "Y", "Z".
	Precedence SgiGetResponsePrecedence `json:"precedence"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The timespan over which the associated radiation belt index is factored. If
	// rbDuration is null, a 24 hour duration should be assumed. Note: rbDuration
	// defines the length of the time window for this data record. The time window
	// start time is defined by sgiDate, and the time window end time is defined by
	// sgiDate plus rbDuration.
	RbDuration int64 `json:"rbDuration"`
	// The value of the radiation belt index. This is the ratio of current intensity of
	// a radiation belt to long-term average value. It's long-term average should be
	// close to 1. Depending on the type of belt sensor, this ratio may measure Flux
	// (number of particles / (cm^2 sec energy solid-angle)), dose rate (rad per
	// second), or relative counts of particles per time (counts per second). The index
	// value may also be normalized, the normalization value typically represents an
	// average of the sensor measurements taken within a region over a given time
	// interval. See the normFactor field for the specific normalization factor, if
	// provided.
	RbIndex float64 `json:"rbIndex"`
	// Region code for the associated radiation belt index. This is the code associated
	// with the corresponding radiation belt location. See the provider card for
	// reference to specific region code definitions.
	RbRegionCode int64 `json:"rbRegionCode"`
	// Daily S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S10 float64 `json:"s10"`
	// 54 day S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S54 float64 `json:"s54"`
	// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
	// record.
	//
	// Any of "I", "N", "P".
	State SgiGetResponseState `json:"state"`
	// The name/location of the station that collected the geomagnetic data for this
	// record.
	StationName string `json:"stationName"`
	// Array containing the sine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stce []float64 `json:"stce"`
	// Array containing the sine spherical harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stci []float64 `json:"stci"`
	// Daily sunspot number.
	SunspotNum float64 `json:"sunspotNum"`
	// Daily sunspot number - high range. This field usually applies to forecast
	// values, based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumHigh float64 `json:"sunspotNumHigh"`
	// Daily sunspot number - low range. This field usually applies to forecast values,
	// based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumLow float64 `json:"sunspotNumLow"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// The type of data contained in this record (e.g. HASDM, JBH09, K-Index, PSD-dB,
	// RBI, RFI-SFU, etc).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Daily Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y10 float64 `json:"y10"`
	// 54 day Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y54 float64 `json:"y54"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EffectiveDate         respjson.Field
		SgiDate               respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AnalyzerAttenuation   respjson.Field
		Ap                    respjson.Field
		ApDuration            respjson.Field
		CoeffDegree           respjson.Field
		CoeffOrder            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Ctce                  respjson.Field
		Ctci                  respjson.Field
		Dst                   respjson.Field
		Dtc                   respjson.Field
		E10                   respjson.Field
		E54                   respjson.Field
		F10                   respjson.Field
		F10High               respjson.Field
		F10Low                respjson.Field
		F54                   respjson.Field
		F81                   respjson.Field
		Frequencies           respjson.Field
		Gamma                 respjson.Field
		IDSensor              respjson.Field
		KIndex                respjson.Field
		Kp                    respjson.Field
		KpDuration            respjson.Field
		M10                   respjson.Field
		M54                   respjson.Field
		Mode                  respjson.Field
		NormFactor            respjson.Field
		ObservedBaseline      respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		Powers                respjson.Field
		Precedence            respjson.Field
		RawFileUri            respjson.Field
		RbDuration            respjson.Field
		RbIndex               respjson.Field
		RbRegionCode          respjson.Field
		S10                   respjson.Field
		S54                   respjson.Field
		State                 respjson.Field
		StationName           respjson.Field
		Stce                  respjson.Field
		Stci                  respjson.Field
		SunspotNum            respjson.Field
		SunspotNumHigh        respjson.Field
		SunspotNumLow         respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Y10                   respjson.Field
		Y54                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SgiGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SgiGetResponse) UnmarshalJSON(data []byte) error {
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
type SgiGetResponseDataMode string

const (
	SgiGetResponseDataModeReal      SgiGetResponseDataMode = "REAL"
	SgiGetResponseDataModeTest      SgiGetResponseDataMode = "TEST"
	SgiGetResponseDataModeSimulated SgiGetResponseDataMode = "SIMULATED"
	SgiGetResponseDataModeExercise  SgiGetResponseDataMode = "EXERCISE"
)

// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
// Y = Emergency, Z = Flash).
type SgiGetResponsePrecedence string

const (
	SgiGetResponsePrecedenceO SgiGetResponsePrecedence = "O"
	SgiGetResponsePrecedenceP SgiGetResponsePrecedence = "P"
	SgiGetResponsePrecedenceR SgiGetResponsePrecedence = "R"
	SgiGetResponsePrecedenceY SgiGetResponsePrecedence = "Y"
	SgiGetResponsePrecedenceZ SgiGetResponsePrecedence = "Z"
)

// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
// record.
type SgiGetResponseState string

const (
	SgiGetResponseStateI SgiGetResponseState = "I"
	SgiGetResponseStateN SgiGetResponseState = "N"
	SgiGetResponseStateP SgiGetResponseState = "P"
)

// Model representation of space weather/solar, geomagnetic, and radiation belt
// indices.
type SgiGetDataByEffectiveAsOfDateResponse struct {
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
	DataMode SgiGetDataByEffectiveAsOfDateResponseDataMode `json:"dataMode,required"`
	// ISO8601 UTC Time the data was received and processed from the source. Typically
	// a source provides data for a date window with each transmission including past,
	// present, and future predicted values.
	EffectiveDate time.Time `json:"effectiveDate,required" format:"date-time"`
	// ISO8601 UTC Time of the index value. This could be a past, current, or future
	// predicted value. Note: sgiDate defines the start time of the time window for
	// this data record.
	SgiDate time.Time `json:"sgiDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Signal analyzer's input attenuation level, in decibels. Attenuation is a setting
	// on the hardware that measures the power of a signal.
	AnalyzerAttenuation float64 `json:"analyzerAttenuation"`
	// Ap is the planetary geomagnetic 2 nT index (00-21 UT) for the timespan specified
	// in apDuration. If apDuration is null, a 3 hour duration should be assumed.
	Ap float64 `json:"ap"`
	// The time, in hours, for which the Ap index value is valid. If null, a span of 3
	// hours is assumed.
	ApDuration int64 `json:"apDuration"`
	// Array containing the degree of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffDegree []int64 `json:"coeffDegree"`
	// Array containing the order of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffOrder []int64 `json:"coeffOrder"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array containing the cosine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctce []float64 `json:"ctce"`
	// Array containing the cosine spherical-harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctci []float64 `json:"ctci"`
	// Disturbance Storm Time geomagnetic index in nT.
	Dst float64 `json:"dst"`
	// delta exospheric temperature correction in units of K.
	Dtc float64 `json:"dtc"`
	// Extreme Ultraviolet (EUV) proxy, E10.7, in x10-22 Watts per meter squared per
	// Hertz, is the integrated solar EUV energy flux at the top of atmosphere and
	// normalized to solar flux units.
	E10 float64 `json:"e10"`
	// E54 (E10-Bar), in x10-22 Watts per meter squared per Hertz, uses the past
	// 54-days E10 values to determine the E10 average.
	E54 float64 `json:"e54"`
	// Daily solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F10 float64 `json:"f10"`
	// Daily F10.7 index - high range, in x10-22 Watts per meter squared per Hertz.
	// This field usually applies to forecast values, based on the consensus of the
	// Solar Cycle 24 Prediction Panel.
	F10High float64 `json:"f10High"`
	// Daily F10.7 index - low range, in x10-22 Watts per meter squared per Hertz. This
	// field usually applies to forecast values, based on the consensus of the Solar
	// Cycle 24 Prediction Panel.
	F10Low float64 `json:"f10Low"`
	// 54 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F54 float64 `json:"f54"`
	// 81 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F81 float64 `json:"f81"`
	// Array of individual power spectral density (PSD) frequencies of the signal, in
	// megahertz. This array should correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies"`
	// The maximum measured gamma deflection during the kpDuration timespan. If
	// kpDuration is null, a 3 hour duration should be assumed.
	Gamma int64 `json:"gamma"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The maximum measured K-Index at the associated station during the kpDuration
	// timespan. The K-Index is a unitless measure (0 - 9) of the deviation in the
	// Earth's magnetic field from normal at the station geographic location, with 0
	// indicating the absence of geomagnetic disturbance, and 9 indicating the most
	// significant disturbance. If kpDuration is null, a 3 hour duration should be
	// assumed.
	KIndex int64 `json:"kIndex"`
	// The Planetary K-index (Kp) over the kpDuration timespan. The Kp-Index is the
	// average K-Index for the entire Earth, utilizing a unitless scale (0-9, in
	// incremenets of 1/3), with 0 indicating the absence of geomagnetic disturbance,
	// and 9 indicating the most significant disturbance. If kpDuration is null, a 3
	// hour duration should be assumed.
	Kp float64 `json:"kp"`
	// The time, in hours, over which the K, Kp, and/or gamma index values are
	// measured. If null, a span of 3 hours is assumed.
	KpDuration int64 `json:"kpDuration"`
	// Daily M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M10 float64 `json:"m10"`
	// 54 day M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M54 float64 `json:"m54"`
	// The transmitted DCA mode of the record (1-3).
	Mode int64 `json:"mode"`
	// The normalization factor that has already been applied to the index value prior
	// to record ingest. Typically used to normalize the index value to a particular
	// interval. Units of the normalization factor may vary depending on the provider
	// of this data (REACH, POES, CEASE3, etc.).
	NormFactor float64 `json:"normFactor"`
	// Observed baseline values of the frequencies specified in the frequencies field,
	// in solar flux units. The baseline values will be used to help detect abnormal
	// readings from the sun that might indicate a flare or other solar activity.
	ObservedBaseline []int64 `json:"observedBaseline"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the reporting source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Array of individual power spectral density (PSD) powers of the signal, in watts.
	// This array should correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers"`
	// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
	// Y = Emergency, Z = Flash).
	//
	// Any of "O", "P", "R", "Y", "Z".
	Precedence SgiGetDataByEffectiveAsOfDateResponsePrecedence `json:"precedence"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The timespan over which the associated radiation belt index is factored. If
	// rbDuration is null, a 24 hour duration should be assumed. Note: rbDuration
	// defines the length of the time window for this data record. The time window
	// start time is defined by sgiDate, and the time window end time is defined by
	// sgiDate plus rbDuration.
	RbDuration int64 `json:"rbDuration"`
	// The value of the radiation belt index. This is the ratio of current intensity of
	// a radiation belt to long-term average value. It's long-term average should be
	// close to 1. Depending on the type of belt sensor, this ratio may measure Flux
	// (number of particles / (cm^2 sec energy solid-angle)), dose rate (rad per
	// second), or relative counts of particles per time (counts per second). The index
	// value may also be normalized, the normalization value typically represents an
	// average of the sensor measurements taken within a region over a given time
	// interval. See the normFactor field for the specific normalization factor, if
	// provided.
	RbIndex float64 `json:"rbIndex"`
	// Region code for the associated radiation belt index. This is the code associated
	// with the corresponding radiation belt location. See the provider card for
	// reference to specific region code definitions.
	RbRegionCode int64 `json:"rbRegionCode"`
	// Daily S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S10 float64 `json:"s10"`
	// 54 day S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S54 float64 `json:"s54"`
	// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
	// record.
	//
	// Any of "I", "N", "P".
	State SgiGetDataByEffectiveAsOfDateResponseState `json:"state"`
	// The name/location of the station that collected the geomagnetic data for this
	// record.
	StationName string `json:"stationName"`
	// Array containing the sine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stce []float64 `json:"stce"`
	// Array containing the sine spherical harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stci []float64 `json:"stci"`
	// Daily sunspot number.
	SunspotNum float64 `json:"sunspotNum"`
	// Daily sunspot number - high range. This field usually applies to forecast
	// values, based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumHigh float64 `json:"sunspotNumHigh"`
	// Daily sunspot number - low range. This field usually applies to forecast values,
	// based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumLow float64 `json:"sunspotNumLow"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// The type of data contained in this record (e.g. HASDM, JBH09, K-Index, PSD-dB,
	// RBI, RFI-SFU, etc).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Daily Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y10 float64 `json:"y10"`
	// 54 day Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y54 float64 `json:"y54"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EffectiveDate         respjson.Field
		SgiDate               respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AnalyzerAttenuation   respjson.Field
		Ap                    respjson.Field
		ApDuration            respjson.Field
		CoeffDegree           respjson.Field
		CoeffOrder            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Ctce                  respjson.Field
		Ctci                  respjson.Field
		Dst                   respjson.Field
		Dtc                   respjson.Field
		E10                   respjson.Field
		E54                   respjson.Field
		F10                   respjson.Field
		F10High               respjson.Field
		F10Low                respjson.Field
		F54                   respjson.Field
		F81                   respjson.Field
		Frequencies           respjson.Field
		Gamma                 respjson.Field
		IDSensor              respjson.Field
		KIndex                respjson.Field
		Kp                    respjson.Field
		KpDuration            respjson.Field
		M10                   respjson.Field
		M54                   respjson.Field
		Mode                  respjson.Field
		NormFactor            respjson.Field
		ObservedBaseline      respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		Powers                respjson.Field
		Precedence            respjson.Field
		RawFileUri            respjson.Field
		RbDuration            respjson.Field
		RbIndex               respjson.Field
		RbRegionCode          respjson.Field
		S10                   respjson.Field
		S54                   respjson.Field
		State                 respjson.Field
		StationName           respjson.Field
		Stce                  respjson.Field
		Stci                  respjson.Field
		SunspotNum            respjson.Field
		SunspotNumHigh        respjson.Field
		SunspotNumLow         respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Y10                   respjson.Field
		Y54                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SgiGetDataByEffectiveAsOfDateResponse) RawJSON() string { return r.JSON.raw }
func (r *SgiGetDataByEffectiveAsOfDateResponse) UnmarshalJSON(data []byte) error {
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
type SgiGetDataByEffectiveAsOfDateResponseDataMode string

const (
	SgiGetDataByEffectiveAsOfDateResponseDataModeReal      SgiGetDataByEffectiveAsOfDateResponseDataMode = "REAL"
	SgiGetDataByEffectiveAsOfDateResponseDataModeTest      SgiGetDataByEffectiveAsOfDateResponseDataMode = "TEST"
	SgiGetDataByEffectiveAsOfDateResponseDataModeSimulated SgiGetDataByEffectiveAsOfDateResponseDataMode = "SIMULATED"
	SgiGetDataByEffectiveAsOfDateResponseDataModeExercise  SgiGetDataByEffectiveAsOfDateResponseDataMode = "EXERCISE"
)

// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
// Y = Emergency, Z = Flash).
type SgiGetDataByEffectiveAsOfDateResponsePrecedence string

const (
	SgiGetDataByEffectiveAsOfDateResponsePrecedenceO SgiGetDataByEffectiveAsOfDateResponsePrecedence = "O"
	SgiGetDataByEffectiveAsOfDateResponsePrecedenceP SgiGetDataByEffectiveAsOfDateResponsePrecedence = "P"
	SgiGetDataByEffectiveAsOfDateResponsePrecedenceR SgiGetDataByEffectiveAsOfDateResponsePrecedence = "R"
	SgiGetDataByEffectiveAsOfDateResponsePrecedenceY SgiGetDataByEffectiveAsOfDateResponsePrecedence = "Y"
	SgiGetDataByEffectiveAsOfDateResponsePrecedenceZ SgiGetDataByEffectiveAsOfDateResponsePrecedence = "Z"
)

// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
// record.
type SgiGetDataByEffectiveAsOfDateResponseState string

const (
	SgiGetDataByEffectiveAsOfDateResponseStateI SgiGetDataByEffectiveAsOfDateResponseState = "I"
	SgiGetDataByEffectiveAsOfDateResponseStateN SgiGetDataByEffectiveAsOfDateResponseState = "N"
	SgiGetDataByEffectiveAsOfDateResponseStateP SgiGetDataByEffectiveAsOfDateResponseState = "P"
)

type SgiQueryhelpResponse struct {
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
func (r SgiQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SgiQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of space weather/solar, geomagnetic, and radiation belt
// indices.
type SgiTupleResponse struct {
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
	DataMode SgiTupleResponseDataMode `json:"dataMode,required"`
	// ISO8601 UTC Time the data was received and processed from the source. Typically
	// a source provides data for a date window with each transmission including past,
	// present, and future predicted values.
	EffectiveDate time.Time `json:"effectiveDate,required" format:"date-time"`
	// ISO8601 UTC Time of the index value. This could be a past, current, or future
	// predicted value. Note: sgiDate defines the start time of the time window for
	// this data record.
	SgiDate time.Time `json:"sgiDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Signal analyzer's input attenuation level, in decibels. Attenuation is a setting
	// on the hardware that measures the power of a signal.
	AnalyzerAttenuation float64 `json:"analyzerAttenuation"`
	// Ap is the planetary geomagnetic 2 nT index (00-21 UT) for the timespan specified
	// in apDuration. If apDuration is null, a 3 hour duration should be assumed.
	Ap float64 `json:"ap"`
	// The time, in hours, for which the Ap index value is valid. If null, a span of 3
	// hours is assumed.
	ApDuration int64 `json:"apDuration"`
	// Array containing the degree of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffDegree []int64 `json:"coeffDegree"`
	// Array containing the order of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffOrder []int64 `json:"coeffOrder"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array containing the cosine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctce []float64 `json:"ctce"`
	// Array containing the cosine spherical-harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctci []float64 `json:"ctci"`
	// Disturbance Storm Time geomagnetic index in nT.
	Dst float64 `json:"dst"`
	// delta exospheric temperature correction in units of K.
	Dtc float64 `json:"dtc"`
	// Extreme Ultraviolet (EUV) proxy, E10.7, in x10-22 Watts per meter squared per
	// Hertz, is the integrated solar EUV energy flux at the top of atmosphere and
	// normalized to solar flux units.
	E10 float64 `json:"e10"`
	// E54 (E10-Bar), in x10-22 Watts per meter squared per Hertz, uses the past
	// 54-days E10 values to determine the E10 average.
	E54 float64 `json:"e54"`
	// Daily solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F10 float64 `json:"f10"`
	// Daily F10.7 index - high range, in x10-22 Watts per meter squared per Hertz.
	// This field usually applies to forecast values, based on the consensus of the
	// Solar Cycle 24 Prediction Panel.
	F10High float64 `json:"f10High"`
	// Daily F10.7 index - low range, in x10-22 Watts per meter squared per Hertz. This
	// field usually applies to forecast values, based on the consensus of the Solar
	// Cycle 24 Prediction Panel.
	F10Low float64 `json:"f10Low"`
	// 54 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F54 float64 `json:"f54"`
	// 81 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F81 float64 `json:"f81"`
	// Array of individual power spectral density (PSD) frequencies of the signal, in
	// megahertz. This array should correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies"`
	// The maximum measured gamma deflection during the kpDuration timespan. If
	// kpDuration is null, a 3 hour duration should be assumed.
	Gamma int64 `json:"gamma"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The maximum measured K-Index at the associated station during the kpDuration
	// timespan. The K-Index is a unitless measure (0 - 9) of the deviation in the
	// Earth's magnetic field from normal at the station geographic location, with 0
	// indicating the absence of geomagnetic disturbance, and 9 indicating the most
	// significant disturbance. If kpDuration is null, a 3 hour duration should be
	// assumed.
	KIndex int64 `json:"kIndex"`
	// The Planetary K-index (Kp) over the kpDuration timespan. The Kp-Index is the
	// average K-Index for the entire Earth, utilizing a unitless scale (0-9, in
	// incremenets of 1/3), with 0 indicating the absence of geomagnetic disturbance,
	// and 9 indicating the most significant disturbance. If kpDuration is null, a 3
	// hour duration should be assumed.
	Kp float64 `json:"kp"`
	// The time, in hours, over which the K, Kp, and/or gamma index values are
	// measured. If null, a span of 3 hours is assumed.
	KpDuration int64 `json:"kpDuration"`
	// Daily M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M10 float64 `json:"m10"`
	// 54 day M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M54 float64 `json:"m54"`
	// The transmitted DCA mode of the record (1-3).
	Mode int64 `json:"mode"`
	// The normalization factor that has already been applied to the index value prior
	// to record ingest. Typically used to normalize the index value to a particular
	// interval. Units of the normalization factor may vary depending on the provider
	// of this data (REACH, POES, CEASE3, etc.).
	NormFactor float64 `json:"normFactor"`
	// Observed baseline values of the frequencies specified in the frequencies field,
	// in solar flux units. The baseline values will be used to help detect abnormal
	// readings from the sun that might indicate a flare or other solar activity.
	ObservedBaseline []int64 `json:"observedBaseline"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the reporting source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Array of individual power spectral density (PSD) powers of the signal, in watts.
	// This array should correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers"`
	// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
	// Y = Emergency, Z = Flash).
	//
	// Any of "O", "P", "R", "Y", "Z".
	Precedence SgiTupleResponsePrecedence `json:"precedence"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The timespan over which the associated radiation belt index is factored. If
	// rbDuration is null, a 24 hour duration should be assumed. Note: rbDuration
	// defines the length of the time window for this data record. The time window
	// start time is defined by sgiDate, and the time window end time is defined by
	// sgiDate plus rbDuration.
	RbDuration int64 `json:"rbDuration"`
	// The value of the radiation belt index. This is the ratio of current intensity of
	// a radiation belt to long-term average value. It's long-term average should be
	// close to 1. Depending on the type of belt sensor, this ratio may measure Flux
	// (number of particles / (cm^2 sec energy solid-angle)), dose rate (rad per
	// second), or relative counts of particles per time (counts per second). The index
	// value may also be normalized, the normalization value typically represents an
	// average of the sensor measurements taken within a region over a given time
	// interval. See the normFactor field for the specific normalization factor, if
	// provided.
	RbIndex float64 `json:"rbIndex"`
	// Region code for the associated radiation belt index. This is the code associated
	// with the corresponding radiation belt location. See the provider card for
	// reference to specific region code definitions.
	RbRegionCode int64 `json:"rbRegionCode"`
	// Daily S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S10 float64 `json:"s10"`
	// 54 day S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S54 float64 `json:"s54"`
	// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
	// record.
	//
	// Any of "I", "N", "P".
	State SgiTupleResponseState `json:"state"`
	// The name/location of the station that collected the geomagnetic data for this
	// record.
	StationName string `json:"stationName"`
	// Array containing the sine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stce []float64 `json:"stce"`
	// Array containing the sine spherical harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stci []float64 `json:"stci"`
	// Daily sunspot number.
	SunspotNum float64 `json:"sunspotNum"`
	// Daily sunspot number - high range. This field usually applies to forecast
	// values, based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumHigh float64 `json:"sunspotNumHigh"`
	// Daily sunspot number - low range. This field usually applies to forecast values,
	// based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumLow float64 `json:"sunspotNumLow"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// The type of data contained in this record (e.g. HASDM, JBH09, K-Index, PSD-dB,
	// RBI, RFI-SFU, etc).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Daily Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y10 float64 `json:"y10"`
	// 54 day Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y54 float64 `json:"y54"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EffectiveDate         respjson.Field
		SgiDate               respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AnalyzerAttenuation   respjson.Field
		Ap                    respjson.Field
		ApDuration            respjson.Field
		CoeffDegree           respjson.Field
		CoeffOrder            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Ctce                  respjson.Field
		Ctci                  respjson.Field
		Dst                   respjson.Field
		Dtc                   respjson.Field
		E10                   respjson.Field
		E54                   respjson.Field
		F10                   respjson.Field
		F10High               respjson.Field
		F10Low                respjson.Field
		F54                   respjson.Field
		F81                   respjson.Field
		Frequencies           respjson.Field
		Gamma                 respjson.Field
		IDSensor              respjson.Field
		KIndex                respjson.Field
		Kp                    respjson.Field
		KpDuration            respjson.Field
		M10                   respjson.Field
		M54                   respjson.Field
		Mode                  respjson.Field
		NormFactor            respjson.Field
		ObservedBaseline      respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		Powers                respjson.Field
		Precedence            respjson.Field
		RawFileUri            respjson.Field
		RbDuration            respjson.Field
		RbIndex               respjson.Field
		RbRegionCode          respjson.Field
		S10                   respjson.Field
		S54                   respjson.Field
		State                 respjson.Field
		StationName           respjson.Field
		Stce                  respjson.Field
		Stci                  respjson.Field
		SunspotNum            respjson.Field
		SunspotNumHigh        respjson.Field
		SunspotNumLow         respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Y10                   respjson.Field
		Y54                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SgiTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SgiTupleResponse) UnmarshalJSON(data []byte) error {
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
type SgiTupleResponseDataMode string

const (
	SgiTupleResponseDataModeReal      SgiTupleResponseDataMode = "REAL"
	SgiTupleResponseDataModeTest      SgiTupleResponseDataMode = "TEST"
	SgiTupleResponseDataModeSimulated SgiTupleResponseDataMode = "SIMULATED"
	SgiTupleResponseDataModeExercise  SgiTupleResponseDataMode = "EXERCISE"
)

// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
// Y = Emergency, Z = Flash).
type SgiTupleResponsePrecedence string

const (
	SgiTupleResponsePrecedenceO SgiTupleResponsePrecedence = "O"
	SgiTupleResponsePrecedenceP SgiTupleResponsePrecedence = "P"
	SgiTupleResponsePrecedenceR SgiTupleResponsePrecedence = "R"
	SgiTupleResponsePrecedenceY SgiTupleResponsePrecedence = "Y"
	SgiTupleResponsePrecedenceZ SgiTupleResponsePrecedence = "Z"
)

// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
// record.
type SgiTupleResponseState string

const (
	SgiTupleResponseStateI SgiTupleResponseState = "I"
	SgiTupleResponseStateN SgiTupleResponseState = "N"
	SgiTupleResponseStateP SgiTupleResponseState = "P"
)

type SgiNewParams struct {
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
	DataMode SgiNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ISO8601 UTC Time the data was received and processed from the source. Typically
	// a source provides data for a date window with each transmission including past,
	// present, and future predicted values.
	EffectiveDate time.Time `json:"effectiveDate,required" format:"date-time"`
	// ISO8601 UTC Time of the index value. This could be a past, current, or future
	// predicted value. Note: sgiDate defines the start time of the time window for
	// this data record.
	SgiDate time.Time `json:"sgiDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Signal analyzer's input attenuation level, in decibels. Attenuation is a setting
	// on the hardware that measures the power of a signal.
	AnalyzerAttenuation param.Opt[float64] `json:"analyzerAttenuation,omitzero"`
	// Ap is the planetary geomagnetic 2 nT index (00-21 UT) for the timespan specified
	// in apDuration. If apDuration is null, a 3 hour duration should be assumed.
	Ap param.Opt[float64] `json:"ap,omitzero"`
	// The time, in hours, for which the Ap index value is valid. If null, a span of 3
	// hours is assumed.
	ApDuration param.Opt[int64] `json:"apDuration,omitzero"`
	// Disturbance Storm Time geomagnetic index in nT.
	Dst param.Opt[float64] `json:"dst,omitzero"`
	// delta exospheric temperature correction in units of K.
	Dtc param.Opt[float64] `json:"dtc,omitzero"`
	// Extreme Ultraviolet (EUV) proxy, E10.7, in x10-22 Watts per meter squared per
	// Hertz, is the integrated solar EUV energy flux at the top of atmosphere and
	// normalized to solar flux units.
	E10 param.Opt[float64] `json:"e10,omitzero"`
	// E54 (E10-Bar), in x10-22 Watts per meter squared per Hertz, uses the past
	// 54-days E10 values to determine the E10 average.
	E54 param.Opt[float64] `json:"e54,omitzero"`
	// Daily solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F10 param.Opt[float64] `json:"f10,omitzero"`
	// Daily F10.7 index - high range, in x10-22 Watts per meter squared per Hertz.
	// This field usually applies to forecast values, based on the consensus of the
	// Solar Cycle 24 Prediction Panel.
	F10High param.Opt[float64] `json:"f10High,omitzero"`
	// Daily F10.7 index - low range, in x10-22 Watts per meter squared per Hertz. This
	// field usually applies to forecast values, based on the consensus of the Solar
	// Cycle 24 Prediction Panel.
	F10Low param.Opt[float64] `json:"f10Low,omitzero"`
	// 54 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F54 param.Opt[float64] `json:"f54,omitzero"`
	// 81 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F81 param.Opt[float64] `json:"f81,omitzero"`
	// The maximum measured gamma deflection during the kpDuration timespan. If
	// kpDuration is null, a 3 hour duration should be assumed.
	Gamma param.Opt[int64] `json:"gamma,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The maximum measured K-Index at the associated station during the kpDuration
	// timespan. The K-Index is a unitless measure (0 - 9) of the deviation in the
	// Earth's magnetic field from normal at the station geographic location, with 0
	// indicating the absence of geomagnetic disturbance, and 9 indicating the most
	// significant disturbance. If kpDuration is null, a 3 hour duration should be
	// assumed.
	KIndex param.Opt[int64] `json:"kIndex,omitzero"`
	// The Planetary K-index (Kp) over the kpDuration timespan. The Kp-Index is the
	// average K-Index for the entire Earth, utilizing a unitless scale (0-9, in
	// incremenets of 1/3), with 0 indicating the absence of geomagnetic disturbance,
	// and 9 indicating the most significant disturbance. If kpDuration is null, a 3
	// hour duration should be assumed.
	Kp param.Opt[float64] `json:"kp,omitzero"`
	// The time, in hours, over which the K, Kp, and/or gamma index values are
	// measured. If null, a span of 3 hours is assumed.
	KpDuration param.Opt[int64] `json:"kpDuration,omitzero"`
	// Daily M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M10 param.Opt[float64] `json:"m10,omitzero"`
	// 54 day M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M54 param.Opt[float64] `json:"m54,omitzero"`
	// The transmitted DCA mode of the record (1-3).
	Mode param.Opt[int64] `json:"mode,omitzero"`
	// The normalization factor that has already been applied to the index value prior
	// to record ingest. Typically used to normalize the index value to a particular
	// interval. Units of the normalization factor may vary depending on the provider
	// of this data (REACH, POES, CEASE3, etc.).
	NormFactor param.Opt[float64] `json:"normFactor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the reporting source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The timespan over which the associated radiation belt index is factored. If
	// rbDuration is null, a 24 hour duration should be assumed. Note: rbDuration
	// defines the length of the time window for this data record. The time window
	// start time is defined by sgiDate, and the time window end time is defined by
	// sgiDate plus rbDuration.
	RbDuration param.Opt[int64] `json:"rbDuration,omitzero"`
	// The value of the radiation belt index. This is the ratio of current intensity of
	// a radiation belt to long-term average value. It's long-term average should be
	// close to 1. Depending on the type of belt sensor, this ratio may measure Flux
	// (number of particles / (cm^2 sec energy solid-angle)), dose rate (rad per
	// second), or relative counts of particles per time (counts per second). The index
	// value may also be normalized, the normalization value typically represents an
	// average of the sensor measurements taken within a region over a given time
	// interval. See the normFactor field for the specific normalization factor, if
	// provided.
	RbIndex param.Opt[float64] `json:"rbIndex,omitzero"`
	// Region code for the associated radiation belt index. This is the code associated
	// with the corresponding radiation belt location. See the provider card for
	// reference to specific region code definitions.
	RbRegionCode param.Opt[int64] `json:"rbRegionCode,omitzero"`
	// Daily S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S10 param.Opt[float64] `json:"s10,omitzero"`
	// 54 day S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S54 param.Opt[float64] `json:"s54,omitzero"`
	// The name/location of the station that collected the geomagnetic data for this
	// record.
	StationName param.Opt[string] `json:"stationName,omitzero"`
	// Daily sunspot number.
	SunspotNum param.Opt[float64] `json:"sunspotNum,omitzero"`
	// Daily sunspot number - high range. This field usually applies to forecast
	// values, based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumHigh param.Opt[float64] `json:"sunspotNumHigh,omitzero"`
	// Daily sunspot number - low range. This field usually applies to forecast values,
	// based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumLow param.Opt[float64] `json:"sunspotNumLow,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The type of data contained in this record (e.g. HASDM, JBH09, K-Index, PSD-dB,
	// RBI, RFI-SFU, etc).
	Type param.Opt[string] `json:"type,omitzero"`
	// Daily Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y10 param.Opt[float64] `json:"y10,omitzero"`
	// 54 day Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y54 param.Opt[float64] `json:"y54,omitzero"`
	// Array containing the degree of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffDegree []int64 `json:"coeffDegree,omitzero"`
	// Array containing the order of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffOrder []int64 `json:"coeffOrder,omitzero"`
	// Array containing the cosine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctce []float64 `json:"ctce,omitzero"`
	// Array containing the cosine spherical-harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctci []float64 `json:"ctci,omitzero"`
	// Array of individual power spectral density (PSD) frequencies of the signal, in
	// megahertz. This array should correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies,omitzero"`
	// Observed baseline values of the frequencies specified in the frequencies field,
	// in solar flux units. The baseline values will be used to help detect abnormal
	// readings from the sun that might indicate a flare or other solar activity.
	ObservedBaseline []int64 `json:"observedBaseline,omitzero"`
	// Array of individual power spectral density (PSD) powers of the signal, in watts.
	// This array should correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers,omitzero"`
	// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
	// Y = Emergency, Z = Flash).
	//
	// Any of "O", "P", "R", "Y", "Z".
	Precedence SgiNewParamsPrecedence `json:"precedence,omitzero"`
	// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
	// record.
	//
	// Any of "I", "N", "P".
	State SgiNewParamsState `json:"state,omitzero"`
	// Array containing the sine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stce []float64 `json:"stce,omitzero"`
	// Array containing the sine spherical harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stci []float64 `json:"stci,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SgiNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SgiNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SgiNewParams) UnmarshalJSON(data []byte) error {
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
type SgiNewParamsDataMode string

const (
	SgiNewParamsDataModeReal      SgiNewParamsDataMode = "REAL"
	SgiNewParamsDataModeTest      SgiNewParamsDataMode = "TEST"
	SgiNewParamsDataModeSimulated SgiNewParamsDataMode = "SIMULATED"
	SgiNewParamsDataModeExercise  SgiNewParamsDataMode = "EXERCISE"
)

// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
// Y = Emergency, Z = Flash).
type SgiNewParamsPrecedence string

const (
	SgiNewParamsPrecedenceO SgiNewParamsPrecedence = "O"
	SgiNewParamsPrecedenceP SgiNewParamsPrecedence = "P"
	SgiNewParamsPrecedenceR SgiNewParamsPrecedence = "R"
	SgiNewParamsPrecedenceY SgiNewParamsPrecedence = "Y"
	SgiNewParamsPrecedenceZ SgiNewParamsPrecedence = "Z"
)

// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
// record.
type SgiNewParamsState string

const (
	SgiNewParamsStateI SgiNewParamsState = "I"
	SgiNewParamsStateN SgiNewParamsState = "N"
	SgiNewParamsStateP SgiNewParamsState = "P"
)

type SgiUpdateParams struct {
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
	DataMode SgiUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// ISO8601 UTC Time the data was received and processed from the source. Typically
	// a source provides data for a date window with each transmission including past,
	// present, and future predicted values.
	EffectiveDate time.Time `json:"effectiveDate,required" format:"date-time"`
	// ISO8601 UTC Time of the index value. This could be a past, current, or future
	// predicted value. Note: sgiDate defines the start time of the time window for
	// this data record.
	SgiDate time.Time `json:"sgiDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Signal analyzer's input attenuation level, in decibels. Attenuation is a setting
	// on the hardware that measures the power of a signal.
	AnalyzerAttenuation param.Opt[float64] `json:"analyzerAttenuation,omitzero"`
	// Ap is the planetary geomagnetic 2 nT index (00-21 UT) for the timespan specified
	// in apDuration. If apDuration is null, a 3 hour duration should be assumed.
	Ap param.Opt[float64] `json:"ap,omitzero"`
	// The time, in hours, for which the Ap index value is valid. If null, a span of 3
	// hours is assumed.
	ApDuration param.Opt[int64] `json:"apDuration,omitzero"`
	// Disturbance Storm Time geomagnetic index in nT.
	Dst param.Opt[float64] `json:"dst,omitzero"`
	// delta exospheric temperature correction in units of K.
	Dtc param.Opt[float64] `json:"dtc,omitzero"`
	// Extreme Ultraviolet (EUV) proxy, E10.7, in x10-22 Watts per meter squared per
	// Hertz, is the integrated solar EUV energy flux at the top of atmosphere and
	// normalized to solar flux units.
	E10 param.Opt[float64] `json:"e10,omitzero"`
	// E54 (E10-Bar), in x10-22 Watts per meter squared per Hertz, uses the past
	// 54-days E10 values to determine the E10 average.
	E54 param.Opt[float64] `json:"e54,omitzero"`
	// Daily solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F10 param.Opt[float64] `json:"f10,omitzero"`
	// Daily F10.7 index - high range, in x10-22 Watts per meter squared per Hertz.
	// This field usually applies to forecast values, based on the consensus of the
	// Solar Cycle 24 Prediction Panel.
	F10High param.Opt[float64] `json:"f10High,omitzero"`
	// Daily F10.7 index - low range, in x10-22 Watts per meter squared per Hertz. This
	// field usually applies to forecast values, based on the consensus of the Solar
	// Cycle 24 Prediction Panel.
	F10Low param.Opt[float64] `json:"f10Low,omitzero"`
	// 54 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F54 param.Opt[float64] `json:"f54,omitzero"`
	// 81 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F81 param.Opt[float64] `json:"f81,omitzero"`
	// The maximum measured gamma deflection during the kpDuration timespan. If
	// kpDuration is null, a 3 hour duration should be assumed.
	Gamma param.Opt[int64] `json:"gamma,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The maximum measured K-Index at the associated station during the kpDuration
	// timespan. The K-Index is a unitless measure (0 - 9) of the deviation in the
	// Earth's magnetic field from normal at the station geographic location, with 0
	// indicating the absence of geomagnetic disturbance, and 9 indicating the most
	// significant disturbance. If kpDuration is null, a 3 hour duration should be
	// assumed.
	KIndex param.Opt[int64] `json:"kIndex,omitzero"`
	// The Planetary K-index (Kp) over the kpDuration timespan. The Kp-Index is the
	// average K-Index for the entire Earth, utilizing a unitless scale (0-9, in
	// incremenets of 1/3), with 0 indicating the absence of geomagnetic disturbance,
	// and 9 indicating the most significant disturbance. If kpDuration is null, a 3
	// hour duration should be assumed.
	Kp param.Opt[float64] `json:"kp,omitzero"`
	// The time, in hours, over which the K, Kp, and/or gamma index values are
	// measured. If null, a span of 3 hours is assumed.
	KpDuration param.Opt[int64] `json:"kpDuration,omitzero"`
	// Daily M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M10 param.Opt[float64] `json:"m10,omitzero"`
	// 54 day M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M54 param.Opt[float64] `json:"m54,omitzero"`
	// The transmitted DCA mode of the record (1-3).
	Mode param.Opt[int64] `json:"mode,omitzero"`
	// The normalization factor that has already been applied to the index value prior
	// to record ingest. Typically used to normalize the index value to a particular
	// interval. Units of the normalization factor may vary depending on the provider
	// of this data (REACH, POES, CEASE3, etc.).
	NormFactor param.Opt[float64] `json:"normFactor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the reporting source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The timespan over which the associated radiation belt index is factored. If
	// rbDuration is null, a 24 hour duration should be assumed. Note: rbDuration
	// defines the length of the time window for this data record. The time window
	// start time is defined by sgiDate, and the time window end time is defined by
	// sgiDate plus rbDuration.
	RbDuration param.Opt[int64] `json:"rbDuration,omitzero"`
	// The value of the radiation belt index. This is the ratio of current intensity of
	// a radiation belt to long-term average value. It's long-term average should be
	// close to 1. Depending on the type of belt sensor, this ratio may measure Flux
	// (number of particles / (cm^2 sec energy solid-angle)), dose rate (rad per
	// second), or relative counts of particles per time (counts per second). The index
	// value may also be normalized, the normalization value typically represents an
	// average of the sensor measurements taken within a region over a given time
	// interval. See the normFactor field for the specific normalization factor, if
	// provided.
	RbIndex param.Opt[float64] `json:"rbIndex,omitzero"`
	// Region code for the associated radiation belt index. This is the code associated
	// with the corresponding radiation belt location. See the provider card for
	// reference to specific region code definitions.
	RbRegionCode param.Opt[int64] `json:"rbRegionCode,omitzero"`
	// Daily S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S10 param.Opt[float64] `json:"s10,omitzero"`
	// 54 day S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S54 param.Opt[float64] `json:"s54,omitzero"`
	// The name/location of the station that collected the geomagnetic data for this
	// record.
	StationName param.Opt[string] `json:"stationName,omitzero"`
	// Daily sunspot number.
	SunspotNum param.Opt[float64] `json:"sunspotNum,omitzero"`
	// Daily sunspot number - high range. This field usually applies to forecast
	// values, based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumHigh param.Opt[float64] `json:"sunspotNumHigh,omitzero"`
	// Daily sunspot number - low range. This field usually applies to forecast values,
	// based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumLow param.Opt[float64] `json:"sunspotNumLow,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The type of data contained in this record (e.g. HASDM, JBH09, K-Index, PSD-dB,
	// RBI, RFI-SFU, etc).
	Type param.Opt[string] `json:"type,omitzero"`
	// Daily Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y10 param.Opt[float64] `json:"y10,omitzero"`
	// 54 day Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y54 param.Opt[float64] `json:"y54,omitzero"`
	// Array containing the degree of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffDegree []int64 `json:"coeffDegree,omitzero"`
	// Array containing the order of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffOrder []int64 `json:"coeffOrder,omitzero"`
	// Array containing the cosine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctce []float64 `json:"ctce,omitzero"`
	// Array containing the cosine spherical-harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctci []float64 `json:"ctci,omitzero"`
	// Array of individual power spectral density (PSD) frequencies of the signal, in
	// megahertz. This array should correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies,omitzero"`
	// Observed baseline values of the frequencies specified in the frequencies field,
	// in solar flux units. The baseline values will be used to help detect abnormal
	// readings from the sun that might indicate a flare or other solar activity.
	ObservedBaseline []int64 `json:"observedBaseline,omitzero"`
	// Array of individual power spectral density (PSD) powers of the signal, in watts.
	// This array should correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers,omitzero"`
	// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
	// Y = Emergency, Z = Flash).
	//
	// Any of "O", "P", "R", "Y", "Z".
	Precedence SgiUpdateParamsPrecedence `json:"precedence,omitzero"`
	// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
	// record.
	//
	// Any of "I", "N", "P".
	State SgiUpdateParamsState `json:"state,omitzero"`
	// Array containing the sine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stce []float64 `json:"stce,omitzero"`
	// Array containing the sine spherical harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stci []float64 `json:"stci,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SgiUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SgiUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SgiUpdateParams) UnmarshalJSON(data []byte) error {
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
type SgiUpdateParamsDataMode string

const (
	SgiUpdateParamsDataModeReal      SgiUpdateParamsDataMode = "REAL"
	SgiUpdateParamsDataModeTest      SgiUpdateParamsDataMode = "TEST"
	SgiUpdateParamsDataModeSimulated SgiUpdateParamsDataMode = "SIMULATED"
	SgiUpdateParamsDataModeExercise  SgiUpdateParamsDataMode = "EXERCISE"
)

// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
// Y = Emergency, Z = Flash).
type SgiUpdateParamsPrecedence string

const (
	SgiUpdateParamsPrecedenceO SgiUpdateParamsPrecedence = "O"
	SgiUpdateParamsPrecedenceP SgiUpdateParamsPrecedence = "P"
	SgiUpdateParamsPrecedenceR SgiUpdateParamsPrecedence = "R"
	SgiUpdateParamsPrecedenceY SgiUpdateParamsPrecedence = "Y"
	SgiUpdateParamsPrecedenceZ SgiUpdateParamsPrecedence = "Z"
)

// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
// record.
type SgiUpdateParamsState string

const (
	SgiUpdateParamsStateI SgiUpdateParamsState = "I"
	SgiUpdateParamsStateN SgiUpdateParamsState = "N"
	SgiUpdateParamsStateP SgiUpdateParamsState = "P"
)

type SgiListParams struct {
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// the data was received and processed from the source. Typically a source provides
	// solar data for a date window with each transmission including past, present, and
	// future predicted values. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EffectiveDate param.Opt[time.Time] `query:"effectiveDate,omitzero" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// of the index value. This could be a past, current, or future predicted value.
	// Note: sgiDate defines the start time of the time window for this data record.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	SgiDate param.Opt[time.Time] `query:"sgiDate,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SgiListParams]'s query parameters as `url.Values`.
func (r SgiListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SgiCountParams struct {
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// the data was received and processed from the source. Typically a source provides
	// solar data for a date window with each transmission including past, present, and
	// future predicted values. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EffectiveDate param.Opt[time.Time] `query:"effectiveDate,omitzero" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// of the index value. This could be a past, current, or future predicted value.
	// Note: sgiDate defines the start time of the time window for this data record.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	SgiDate param.Opt[time.Time] `query:"sgiDate,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SgiCountParams]'s query parameters as `url.Values`.
func (r SgiCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SgiNewBulkParams struct {
	Body []SgiNewBulkParamsBody
	paramObj
}

func (r SgiNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SgiNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Model representation of space weather/solar, geomagnetic, and radiation belt
// indices.
//
// The properties ClassificationMarking, DataMode, EffectiveDate, SgiDate, Source
// are required.
type SgiNewBulkParamsBody struct {
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
	// ISO8601 UTC Time the data was received and processed from the source. Typically
	// a source provides data for a date window with each transmission including past,
	// present, and future predicted values.
	EffectiveDate time.Time `json:"effectiveDate,required" format:"date-time"`
	// ISO8601 UTC Time of the index value. This could be a past, current, or future
	// predicted value. Note: sgiDate defines the start time of the time window for
	// this data record.
	SgiDate time.Time `json:"sgiDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Signal analyzer's input attenuation level, in decibels. Attenuation is a setting
	// on the hardware that measures the power of a signal.
	AnalyzerAttenuation param.Opt[float64] `json:"analyzerAttenuation,omitzero"`
	// Ap is the planetary geomagnetic 2 nT index (00-21 UT) for the timespan specified
	// in apDuration. If apDuration is null, a 3 hour duration should be assumed.
	Ap param.Opt[float64] `json:"ap,omitzero"`
	// The time, in hours, for which the Ap index value is valid. If null, a span of 3
	// hours is assumed.
	ApDuration param.Opt[int64] `json:"apDuration,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Disturbance Storm Time geomagnetic index in nT.
	Dst param.Opt[float64] `json:"dst,omitzero"`
	// delta exospheric temperature correction in units of K.
	Dtc param.Opt[float64] `json:"dtc,omitzero"`
	// Extreme Ultraviolet (EUV) proxy, E10.7, in x10-22 Watts per meter squared per
	// Hertz, is the integrated solar EUV energy flux at the top of atmosphere and
	// normalized to solar flux units.
	E10 param.Opt[float64] `json:"e10,omitzero"`
	// E54 (E10-Bar), in x10-22 Watts per meter squared per Hertz, uses the past
	// 54-days E10 values to determine the E10 average.
	E54 param.Opt[float64] `json:"e54,omitzero"`
	// Daily solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F10 param.Opt[float64] `json:"f10,omitzero"`
	// Daily F10.7 index - high range, in x10-22 Watts per meter squared per Hertz.
	// This field usually applies to forecast values, based on the consensus of the
	// Solar Cycle 24 Prediction Panel.
	F10High param.Opt[float64] `json:"f10High,omitzero"`
	// Daily F10.7 index - low range, in x10-22 Watts per meter squared per Hertz. This
	// field usually applies to forecast values, based on the consensus of the Solar
	// Cycle 24 Prediction Panel.
	F10Low param.Opt[float64] `json:"f10Low,omitzero"`
	// 54 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F54 param.Opt[float64] `json:"f54,omitzero"`
	// 81 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F81 param.Opt[float64] `json:"f81,omitzero"`
	// The maximum measured gamma deflection during the kpDuration timespan. If
	// kpDuration is null, a 3 hour duration should be assumed.
	Gamma param.Opt[int64] `json:"gamma,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The maximum measured K-Index at the associated station during the kpDuration
	// timespan. The K-Index is a unitless measure (0 - 9) of the deviation in the
	// Earth's magnetic field from normal at the station geographic location, with 0
	// indicating the absence of geomagnetic disturbance, and 9 indicating the most
	// significant disturbance. If kpDuration is null, a 3 hour duration should be
	// assumed.
	KIndex param.Opt[int64] `json:"kIndex,omitzero"`
	// The Planetary K-index (Kp) over the kpDuration timespan. The Kp-Index is the
	// average K-Index for the entire Earth, utilizing a unitless scale (0-9, in
	// incremenets of 1/3), with 0 indicating the absence of geomagnetic disturbance,
	// and 9 indicating the most significant disturbance. If kpDuration is null, a 3
	// hour duration should be assumed.
	Kp param.Opt[float64] `json:"kp,omitzero"`
	// The time, in hours, over which the K, Kp, and/or gamma index values are
	// measured. If null, a span of 3 hours is assumed.
	KpDuration param.Opt[int64] `json:"kpDuration,omitzero"`
	// Daily M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M10 param.Opt[float64] `json:"m10,omitzero"`
	// 54 day M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M54 param.Opt[float64] `json:"m54,omitzero"`
	// The transmitted DCA mode of the record (1-3).
	Mode param.Opt[int64] `json:"mode,omitzero"`
	// The normalization factor that has already been applied to the index value prior
	// to record ingest. Typically used to normalize the index value to a particular
	// interval. Units of the normalization factor may vary depending on the provider
	// of this data (REACH, POES, CEASE3, etc.).
	NormFactor param.Opt[float64] `json:"normFactor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the reporting source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The timespan over which the associated radiation belt index is factored. If
	// rbDuration is null, a 24 hour duration should be assumed. Note: rbDuration
	// defines the length of the time window for this data record. The time window
	// start time is defined by sgiDate, and the time window end time is defined by
	// sgiDate plus rbDuration.
	RbDuration param.Opt[int64] `json:"rbDuration,omitzero"`
	// The value of the radiation belt index. This is the ratio of current intensity of
	// a radiation belt to long-term average value. It's long-term average should be
	// close to 1. Depending on the type of belt sensor, this ratio may measure Flux
	// (number of particles / (cm^2 sec energy solid-angle)), dose rate (rad per
	// second), or relative counts of particles per time (counts per second). The index
	// value may also be normalized, the normalization value typically represents an
	// average of the sensor measurements taken within a region over a given time
	// interval. See the normFactor field for the specific normalization factor, if
	// provided.
	RbIndex param.Opt[float64] `json:"rbIndex,omitzero"`
	// Region code for the associated radiation belt index. This is the code associated
	// with the corresponding radiation belt location. See the provider card for
	// reference to specific region code definitions.
	RbRegionCode param.Opt[int64] `json:"rbRegionCode,omitzero"`
	// Daily S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S10 param.Opt[float64] `json:"s10,omitzero"`
	// 54 day S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S54 param.Opt[float64] `json:"s54,omitzero"`
	// The name/location of the station that collected the geomagnetic data for this
	// record.
	StationName param.Opt[string] `json:"stationName,omitzero"`
	// Daily sunspot number.
	SunspotNum param.Opt[float64] `json:"sunspotNum,omitzero"`
	// Daily sunspot number - high range. This field usually applies to forecast
	// values, based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumHigh param.Opt[float64] `json:"sunspotNumHigh,omitzero"`
	// Daily sunspot number - low range. This field usually applies to forecast values,
	// based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumLow param.Opt[float64] `json:"sunspotNumLow,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The type of data contained in this record (e.g. HASDM, JBH09, K-Index, PSD-dB,
	// RBI, RFI-SFU, etc).
	Type param.Opt[string] `json:"type,omitzero"`
	// Daily Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y10 param.Opt[float64] `json:"y10,omitzero"`
	// 54 day Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y54 param.Opt[float64] `json:"y54,omitzero"`
	// Array containing the degree of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffDegree []int64 `json:"coeffDegree,omitzero"`
	// Array containing the order of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffOrder []int64 `json:"coeffOrder,omitzero"`
	// Array containing the cosine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctce []float64 `json:"ctce,omitzero"`
	// Array containing the cosine spherical-harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctci []float64 `json:"ctci,omitzero"`
	// Array of individual power spectral density (PSD) frequencies of the signal, in
	// megahertz. This array should correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies,omitzero"`
	// Observed baseline values of the frequencies specified in the frequencies field,
	// in solar flux units. The baseline values will be used to help detect abnormal
	// readings from the sun that might indicate a flare or other solar activity.
	ObservedBaseline []int64 `json:"observedBaseline,omitzero"`
	// Array of individual power spectral density (PSD) powers of the signal, in watts.
	// This array should correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers,omitzero"`
	// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
	// Y = Emergency, Z = Flash).
	//
	// Any of "O", "P", "R", "Y", "Z".
	Precedence string `json:"precedence,omitzero"`
	// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
	// record.
	//
	// Any of "I", "N", "P".
	State string `json:"state,omitzero"`
	// Array containing the sine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stce []float64 `json:"stce,omitzero"`
	// Array containing the sine spherical harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stci []float64 `json:"stci,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SgiNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SgiNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SgiNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SgiNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SgiNewBulkParamsBody](
		"precedence", "O", "P", "R", "Y", "Z",
	)
	apijson.RegisterFieldValidator[SgiNewBulkParamsBody](
		"state", "I", "N", "P",
	)
}

type SgiGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SgiGetParams]'s query parameters as `url.Values`.
func (r SgiGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SgiGetDataByEffectiveAsOfDateParams struct {
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// the data was received and processed from the source. Typically a source provides
	// solar data for a date window with each transmission including past, present, and
	// future predicted values. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EffectiveDate param.Opt[time.Time] `query:"effectiveDate,omitzero" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// of the index value. This could be a past, current, or future predicted value.
	// Note: sgiDate defines the start time of the time window for this data record.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	SgiDate param.Opt[time.Time] `query:"sgiDate,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SgiGetDataByEffectiveAsOfDateParams]'s query parameters as
// `url.Values`.
func (r SgiGetDataByEffectiveAsOfDateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SgiTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// the data was received and processed from the source. Typically a source provides
	// solar data for a date window with each transmission including past, present, and
	// future predicted values. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EffectiveDate param.Opt[time.Time] `query:"effectiveDate,omitzero" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'effectiveDate, sgiDate' are required.) ISO8601 UTC Time
	// of the index value. This could be a past, current, or future predicted value.
	// Note: sgiDate defines the start time of the time window for this data record.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	SgiDate param.Opt[time.Time] `query:"sgiDate,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SgiTupleParams]'s query parameters as `url.Values`.
func (r SgiTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SgiUnvalidatedPublishParams struct {
	Body []SgiUnvalidatedPublishParamsBody
	paramObj
}

func (r SgiUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SgiUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Model representation of space weather/solar, geomagnetic, and radiation belt
// indices.
//
// The properties ClassificationMarking, DataMode, EffectiveDate, SgiDate, Source
// are required.
type SgiUnvalidatedPublishParamsBody struct {
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
	// ISO8601 UTC Time the data was received and processed from the source. Typically
	// a source provides data for a date window with each transmission including past,
	// present, and future predicted values.
	EffectiveDate time.Time `json:"effectiveDate,required" format:"date-time"`
	// ISO8601 UTC Time of the index value. This could be a past, current, or future
	// predicted value. Note: sgiDate defines the start time of the time window for
	// this data record.
	SgiDate time.Time `json:"sgiDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Signal analyzer's input attenuation level, in decibels. Attenuation is a setting
	// on the hardware that measures the power of a signal.
	AnalyzerAttenuation param.Opt[float64] `json:"analyzerAttenuation,omitzero"`
	// Ap is the planetary geomagnetic 2 nT index (00-21 UT) for the timespan specified
	// in apDuration. If apDuration is null, a 3 hour duration should be assumed.
	Ap param.Opt[float64] `json:"ap,omitzero"`
	// The time, in hours, for which the Ap index value is valid. If null, a span of 3
	// hours is assumed.
	ApDuration param.Opt[int64] `json:"apDuration,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Disturbance Storm Time geomagnetic index in nT.
	Dst param.Opt[float64] `json:"dst,omitzero"`
	// delta exospheric temperature correction in units of K.
	Dtc param.Opt[float64] `json:"dtc,omitzero"`
	// Extreme Ultraviolet (EUV) proxy, E10.7, in x10-22 Watts per meter squared per
	// Hertz, is the integrated solar EUV energy flux at the top of atmosphere and
	// normalized to solar flux units.
	E10 param.Opt[float64] `json:"e10,omitzero"`
	// E54 (E10-Bar), in x10-22 Watts per meter squared per Hertz, uses the past
	// 54-days E10 values to determine the E10 average.
	E54 param.Opt[float64] `json:"e54,omitzero"`
	// Daily solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F10 param.Opt[float64] `json:"f10,omitzero"`
	// Daily F10.7 index - high range, in x10-22 Watts per meter squared per Hertz.
	// This field usually applies to forecast values, based on the consensus of the
	// Solar Cycle 24 Prediction Panel.
	F10High param.Opt[float64] `json:"f10High,omitzero"`
	// Daily F10.7 index - low range, in x10-22 Watts per meter squared per Hertz. This
	// field usually applies to forecast values, based on the consensus of the Solar
	// Cycle 24 Prediction Panel.
	F10Low param.Opt[float64] `json:"f10Low,omitzero"`
	// 54 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F54 param.Opt[float64] `json:"f54,omitzero"`
	// 81 day solar 10.7 cm radio flux in x10-22 Watts per meter squared per Hertz.
	F81 param.Opt[float64] `json:"f81,omitzero"`
	// The maximum measured gamma deflection during the kpDuration timespan. If
	// kpDuration is null, a 3 hour duration should be assumed.
	Gamma param.Opt[int64] `json:"gamma,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The maximum measured K-Index at the associated station during the kpDuration
	// timespan. The K-Index is a unitless measure (0 - 9) of the deviation in the
	// Earth's magnetic field from normal at the station geographic location, with 0
	// indicating the absence of geomagnetic disturbance, and 9 indicating the most
	// significant disturbance. If kpDuration is null, a 3 hour duration should be
	// assumed.
	KIndex param.Opt[int64] `json:"kIndex,omitzero"`
	// The Planetary K-index (Kp) over the kpDuration timespan. The Kp-Index is the
	// average K-Index for the entire Earth, utilizing a unitless scale (0-9, in
	// incremenets of 1/3), with 0 indicating the absence of geomagnetic disturbance,
	// and 9 indicating the most significant disturbance. If kpDuration is null, a 3
	// hour duration should be assumed.
	Kp param.Opt[float64] `json:"kp,omitzero"`
	// The time, in hours, over which the K, Kp, and/or gamma index values are
	// measured. If null, a span of 3 hours is assumed.
	KpDuration param.Opt[int64] `json:"kpDuration,omitzero"`
	// Daily M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M10 param.Opt[float64] `json:"m10,omitzero"`
	// 54 day M10.7 index for 100-110 km heating of O2 by solar photosphere. 160 nm SRC
	// emissions in x10-22 Watts per meter squared per Hertz.
	M54 param.Opt[float64] `json:"m54,omitzero"`
	// The transmitted DCA mode of the record (1-3).
	Mode param.Opt[int64] `json:"mode,omitzero"`
	// The normalization factor that has already been applied to the index value prior
	// to record ingest. Typically used to normalize the index value to a particular
	// interval. Units of the normalization factor may vary depending on the provider
	// of this data (REACH, POES, CEASE3, etc.).
	NormFactor param.Opt[float64] `json:"normFactor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the reporting source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The timespan over which the associated radiation belt index is factored. If
	// rbDuration is null, a 24 hour duration should be assumed. Note: rbDuration
	// defines the length of the time window for this data record. The time window
	// start time is defined by sgiDate, and the time window end time is defined by
	// sgiDate plus rbDuration.
	RbDuration param.Opt[int64] `json:"rbDuration,omitzero"`
	// The value of the radiation belt index. This is the ratio of current intensity of
	// a radiation belt to long-term average value. It's long-term average should be
	// close to 1. Depending on the type of belt sensor, this ratio may measure Flux
	// (number of particles / (cm^2 sec energy solid-angle)), dose rate (rad per
	// second), or relative counts of particles per time (counts per second). The index
	// value may also be normalized, the normalization value typically represents an
	// average of the sensor measurements taken within a region over a given time
	// interval. See the normFactor field for the specific normalization factor, if
	// provided.
	RbIndex param.Opt[float64] `json:"rbIndex,omitzero"`
	// Region code for the associated radiation belt index. This is the code associated
	// with the corresponding radiation belt location. See the provider card for
	// reference to specific region code definitions.
	RbRegionCode param.Opt[int64] `json:"rbRegionCode,omitzero"`
	// Daily S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S10 param.Opt[float64] `json:"s10,omitzero"`
	// 54 day S10.7 index for >200 km heating of O by solar chromosphere. 28.4-30.4 nm
	// emissions in x10-22 Watts per meter squared per Hertz.
	S54 param.Opt[float64] `json:"s54,omitzero"`
	// The name/location of the station that collected the geomagnetic data for this
	// record.
	StationName param.Opt[string] `json:"stationName,omitzero"`
	// Daily sunspot number.
	SunspotNum param.Opt[float64] `json:"sunspotNum,omitzero"`
	// Daily sunspot number - high range. This field usually applies to forecast
	// values, based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumHigh param.Opt[float64] `json:"sunspotNumHigh,omitzero"`
	// Daily sunspot number - low range. This field usually applies to forecast values,
	// based on the consensus of the Solar Cycle 24 Prediction Panel.
	SunspotNumLow param.Opt[float64] `json:"sunspotNumLow,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The type of data contained in this record (e.g. HASDM, JBH09, K-Index, PSD-dB,
	// RBI, RFI-SFU, etc).
	Type param.Opt[string] `json:"type,omitzero"`
	// Daily Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y10 param.Opt[float64] `json:"y10,omitzero"`
	// 54 day Y10.7 index for 85-90 km heating of N2, O2, H2O, NO by solar coronal.
	// 0.1-0.8 nm and Lya 121 nm emissions in x10-22 Watts per meter squared per Hertz.
	Y54 param.Opt[float64] `json:"y54,omitzero"`
	// Array containing the degree of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffDegree []int64 `json:"coeffDegree,omitzero"`
	// Array containing the order of the temperature coefficients. The coeffDegree and
	// coeffOrder arrays must be the same length.
	CoeffOrder []int64 `json:"coeffOrder,omitzero"`
	// Array containing the cosine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctce []float64 `json:"ctce,omitzero"`
	// Array containing the cosine spherical-harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Ctci []float64 `json:"ctci,omitzero"`
	// Array of individual power spectral density (PSD) frequencies of the signal, in
	// megahertz. This array should correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies,omitzero"`
	// Observed baseline values of the frequencies specified in the frequencies field,
	// in solar flux units. The baseline values will be used to help detect abnormal
	// readings from the sun that might indicate a flare or other solar activity.
	ObservedBaseline []int64 `json:"observedBaseline,omitzero"`
	// Array of individual power spectral density (PSD) powers of the signal, in watts.
	// This array should correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers,omitzero"`
	// The precedence of data in this record (O = Immediate, P = Priority, R = Routine,
	// Y = Emergency, Z = Flash).
	//
	// Any of "O", "P", "R", "Y", "Z".
	Precedence string `json:"precedence,omitzero"`
	// State indicating Issued (I), Nowcast (N), or Predicted (P) values for this
	// record.
	//
	// Any of "I", "N", "P".
	State string `json:"state,omitzero"`
	// Array containing the sine spherical-harmonic coefficients for Exospheric
	// temperature (DTC) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stce []float64 `json:"stce,omitzero"`
	// Array containing the sine spherical harmonic coefficients for Inflection
	// temperature (DTX) difference. Each array element corresponds to the positional
	// index of the coeffDegree and coeffOrder arrays.
	Stci []float64 `json:"stci,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SgiUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SgiUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SgiUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SgiUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SgiUnvalidatedPublishParamsBody](
		"precedence", "O", "P", "R", "Y", "Z",
	)
	apijson.RegisterFieldValidator[SgiUnvalidatedPublishParamsBody](
		"state", "I", "N", "P",
	)
}
