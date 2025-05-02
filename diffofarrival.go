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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
)

// DiffofarrivalService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiffofarrivalService] method instead.
type DiffofarrivalService struct {
	Options []option.RequestOption
	History DiffofarrivalHistoryService
}

// NewDiffofarrivalService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDiffofarrivalService(opts ...option.RequestOption) (r DiffofarrivalService) {
	r = DiffofarrivalService{}
	r.Options = opts
	r.History = NewDiffofarrivalHistoryService(opts...)
	return
}

// Service operation to get a single TDOA/FDOA record by its unique ID passed as a
// path parameter.
func (r *DiffofarrivalService) Get(ctx context.Context, id string, query DiffofarrivalGetParams, opts ...option.RequestOption) (res *DiffofarrivalFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/diffofarrival/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *DiffofarrivalService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/diffofarrival/queryhelp"
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
func (r *DiffofarrivalService) Tuple(ctx context.Context, query DiffofarrivalTupleParams, opts ...option.RequestOption) (res *[]DiffofarrivalFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/diffofarrival/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple TDOA/FDOA records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *DiffofarrivalService) UnvalidatedPublish(ctx context.Context, body DiffofarrivalUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-diffofarrival"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type DiffofarrivalGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiffofarrivalGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [DiffofarrivalGetParams]'s query parameters as `url.Values`.
func (r DiffofarrivalGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DiffofarrivalTupleParams struct {
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
func (f DiffofarrivalTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [DiffofarrivalTupleParams]'s query parameters as
// `url.Values`.
func (r DiffofarrivalTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DiffofarrivalUnvalidatedPublishParams struct {
	Body []DiffofarrivalUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiffofarrivalUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r DiffofarrivalUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Model representation of Signal time and frequency difference of arrival
// (TDOA/FDOA) information obtained by using passive RF based sensor
// phenomenologies and sensor triangulation.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type DiffofarrivalUnvalidatedPublishParamsBody struct {
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
	// Bandwidth of the signal in Hz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Collection mode (e.g. SURVEY, SPOT_SEARCH, NEIGHBORHOOD_WATCH, DIRECTED_SEARCH,
	// MANUAL, etc).
	CollectionMode param.Opt[string] `json:"collectionMode,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Delta range, in km. Delta range calculation convention is (sensor2 - sensor1).
	DeltaRange param.Opt[float64] `json:"deltaRange,omitzero"`
	// Delta range rate, in km/sec. Delta range rate calculation convention is
	// (sensor2 - sensor1).
	DeltaRangeRate param.Opt[float64] `json:"deltaRangeRate,omitzero"`
	// One sigma uncertainty in the delta range rate, in km/sec.
	DeltaRangeRateUnc param.Opt[float64] `json:"deltaRangeRateUnc,omitzero"`
	// One sigma uncertainty in delta range, in km.
	DeltaRangeUnc param.Opt[float64] `json:"deltaRangeUnc,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Frequency difference of arrival of the center frequency signal, in Hz. FDOA
	// calculation convention is (sensor2 - sensor1).
	Fdoa param.Opt[float64] `json:"fdoa,omitzero"`
	// One sigma uncertainty in frequency difference of arrival of the center frequency
	// signal, in Hz.
	FdoaUnc param.Opt[float64] `json:"fdoaUnc,omitzero"`
	// Center frequency of the collect in Hz.
	Frequency param.Opt[float64] `json:"frequency,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Sensor ID of the primary/1st sensor used for this measurement.
	IDSensor1 param.Opt[string] `json:"idSensor1,omitzero"`
	// Sensor ID of the secondary/2nd sensor used for this measurement.
	IDSensor2 param.Opt[string] `json:"idSensor2,omitzero"`
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
	// Optional identifier provided by DOA source to indicate the primary/1st sensor
	// identifier used for this measurement. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorId1 param.Opt[string] `json:"origSensorId1,omitzero"`
	// Optional identifier provided by DOA source to indicate the secondary/2nd sensor
	// identifier used for this this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorId2 param.Opt[string] `json:"origSensorId2,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor 2 altitude at obTime (if mobile/onorbit) in km. If null, can be obtained
	// from sensor info.
	Sen2alt param.Opt[float64] `json:"sen2alt,omitzero"`
	// Sensor 2 WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info.
	Sen2lat param.Opt[float64] `json:"sen2lat,omitzero"`
	// Sensor 2 WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can
	// be obtained from sensor info.
	Sen2lon param.Opt[float64] `json:"sen2lon,omitzero"`
	// Sensor altitude at obTime (if mobile/onorbit) in km. If null, can be obtained
	// from sensor info.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// The signal arrival delay relative to sensor 1 in seconds.
	Sensor1Delay param.Opt[float64] `json:"sensor1Delay,omitzero"`
	// The signal arrival delay relative to sensor 2 in seconds.
	Sensor2Delay param.Opt[float64] `json:"sensor2Delay,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Time difference of arrival of the center frequency signal, in seconds. TDOA
	// calculation convention is (sensor2 - sensor1).
	Tdoa param.Opt[float64] `json:"tdoa,omitzero"`
	// One sigma uncertainty in time difference of arrival of the center frequency
	// signal, in seconds.
	TdoaUnc param.Opt[float64] `json:"tdoaUnc,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
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
func (f DiffofarrivalUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r DiffofarrivalUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow DiffofarrivalUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[DiffofarrivalUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
