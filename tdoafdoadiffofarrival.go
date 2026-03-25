// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// These services provide operations for querying and manipulation of Signal time
// and frequency difference of arrival (TDOA/FDOA) information obtained by using
// passive RF based sensor phenomenologies and sensor triangulation. The J2000
// coordinate frame is the preferred frame for all observations, but in some cases
// observations may be in another frame depending on the provider.
//
// TdoaFdoaDiffofarrivalService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTdoaFdoaDiffofarrivalService] method instead.
type TdoaFdoaDiffofarrivalService struct {
	Options []option.RequestOption
	// These services provide operations for querying and manipulation of Signal time
	// and frequency difference of arrival (TDOA/FDOA) information obtained by using
	// passive RF based sensor phenomenologies and sensor triangulation. The J2000
	// coordinate frame is the preferred frame for all observations, but in some cases
	// observations may be in another frame depending on the provider.
	History TdoaFdoaDiffofarrivalHistoryService
}

// NewTdoaFdoaDiffofarrivalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTdoaFdoaDiffofarrivalService(opts ...option.RequestOption) (r TdoaFdoaDiffofarrivalService) {
	r = TdoaFdoaDiffofarrivalService{}
	r.Options = opts
	r.History = NewTdoaFdoaDiffofarrivalHistoryService(opts...)
	return
}

// Service operation to take a single TDOA/FDOA record as a POST body and ingest
// into the database. This operation is not intended to be used for automated feeds
// into UDL. Data providers should contact the UDL team for specific role
// assignments and for instructions on setting up a permanent feed through an
// alternate mechanism.
func (r *TdoaFdoaDiffofarrivalService) New(ctx context.Context, body TdoaFdoaDiffofarrivalNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/diffofarrival"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TdoaFdoaDiffofarrivalService) List(ctx context.Context, query TdoaFdoaDiffofarrivalListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[DiffofarrivalAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/diffofarrival"
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
func (r *TdoaFdoaDiffofarrivalService) ListAutoPaging(ctx context.Context, query TdoaFdoaDiffofarrivalListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[DiffofarrivalAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *TdoaFdoaDiffofarrivalService) Count(ctx context.Context, query TdoaFdoaDiffofarrivalCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/diffofarrival/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// TDOA/FDOA records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *TdoaFdoaDiffofarrivalService) NewBulk(ctx context.Context, body TdoaFdoaDiffofarrivalNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/diffofarrival/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Model representation of Signal time and frequency difference of arrival
// (TDOA/FDOA) information obtained by using passive RF based sensor
// phenomenologies and sensor triangulation.
type DiffofarrivalAbridged struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
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
	DataMode DiffofarrivalAbridgedDataMode `json:"dataMode" api:"required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime" api:"required" format:"date-time"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Line of sight azimuth angle for sensor 1, in degrees and topocentric frame.
	// Azimuth ranges from 0 to 360 degrees.
	Azimuth1 float64 `json:"azimuth1"`
	// Line of sight azimuth angle for sensor 2, in degrees and topocentric frame.
	// Azimuth ranges from 0 to 360 degrees.
	Azimuth2 float64 `json:"azimuth2"`
	// Bandwidth of the signal, in hertz.
	Bandwidth float64 `json:"bandwidth"`
	// Collection mode (e.g., DIRECTED_SEARCH, MANUAL, NEIGHBORHOOD_WATCH, SPOT_SEARCH,
	// SURVEY, etc.).
	CollectionMode string `json:"collectionMode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Delta range, in kilometers. Delta range calculation convention is (sensor2 -
	// sensor1).
	DeltaRange float64 `json:"deltaRange"`
	// Delta range rate, in kilometers per second. Delta range rate calculation
	// convention is (sensor2 - sensor1).
	DeltaRangeRate float64 `json:"deltaRangeRate"`
	// One sigma uncertainty in the delta range rate, in kilometers per second.
	DeltaRangeRateUnc float64 `json:"deltaRangeRateUnc"`
	// One sigma uncertainty in delta range, in kilometers.
	DeltaRangeUnc float64 `json:"deltaRangeUnc"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Line of sight elevation angle for sensor 1, in degrees and topocentric frame.
	// Elevation ranges from -90 to 90 degrees.
	Elevation1 float64 `json:"elevation1"`
	// Line of sight elevation angle for sensor 2, in degrees and topocentric frame.
	// Elevation ranges from -90 to 90 degrees.
	Elevation2 float64 `json:"elevation2"`
	// Frequency difference of arrival of the center frequency signal, in hertz. FDOA
	// calculation convention is (sensor2 - sensor1).
	Fdoa float64 `json:"fdoa"`
	// One sigma uncertainty in frequency difference of arrival of the center frequency
	// signal, in hertz.
	FdoaUnc float64 `json:"fdoaUnc"`
	// Center frequency of the collect, in hertz.
	Frequency float64 `json:"frequency"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Sensor ID of the primary/1st sensor used for this measurement.
	IDSensor1 string `json:"idSensor1"`
	// Sensor ID of the secondary/2nd sensor used for this measurement.
	IDSensor2 string `json:"idSensor2"`
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
	// Optional identifier provided by DOA source to indicate the primary/1st sensor
	// identifier used for this measurement. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorId1 string `json:"origSensorId1"`
	// Optional identifier provided by DOA source to indicate the secondary/2nd sensor
	// identifier used for this this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorId2 string `json:"origSensorId2"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	PolarityType string `json:"polarityType"`
	// Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sensor 2 altitude at obTime (if mobile/onorbit), in kilometers. If null, can be
	// obtained from sensor info.
	Sen2alt float64 `json:"sen2alt"`
	// Sensor 2 WGS84 latitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -90 to 90 degrees (negative values south of
	// equator).
	Sen2lat float64 `json:"sen2lat"`
	// Sensor 2 WGS84 longitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Sen2lon float64 `json:"sen2lon"`
	// Sensor altitude at obTime (if mobile/onorbit), in kilometers. If null, can be
	// obtained from sensor info.
	Senalt float64 `json:"senalt"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit), in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon float64 `json:"senlon"`
	// The signal arrival delay relative to sensor 1, in seconds.
	Sensor1Delay float64 `json:"sensor1Delay"`
	// The signal arrival delay relative to sensor 2, in seconds.
	Sensor2Delay float64 `json:"sensor2Delay"`
	// Signal to noise ratio, in decibels.
	Snr float64 `json:"snr"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID string `json:"taskId"`
	// Time difference of arrival of the center frequency signal, in seconds. TDOA
	// calculation convention is (sensor2 - sensor1).
	Tdoa float64 `json:"tdoa"`
	// The number of full signal cycles the time difference of arrival measurement
	// could be shifted due to repeating patterns in the signal (0 = no ambiguity, +1 =
	// one positive cycle ambiguity, -1 = one negative cycle ambiguity).
	TdoaAmb int64 `json:"tdoaAmb"`
	// One sigma uncertainty in time difference of arrival of the center frequency
	// signal, in seconds.
	TdoaUnc float64 `json:"tdoaUnc"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct bool `json:"uct"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Azimuth1              respjson.Field
		Azimuth2              respjson.Field
		Bandwidth             respjson.Field
		CollectionMode        respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeltaRange            respjson.Field
		DeltaRangeRate        respjson.Field
		DeltaRangeRateUnc     respjson.Field
		DeltaRangeUnc         respjson.Field
		Descriptor            respjson.Field
		Elevation1            respjson.Field
		Elevation2            respjson.Field
		Fdoa                  respjson.Field
		FdoaUnc               respjson.Field
		Frequency             respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor1             respjson.Field
		IDSensor2             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorId1         respjson.Field
		OrigSensorId2         respjson.Field
		PolarityType          respjson.Field
		SatNo                 respjson.Field
		Sen2alt               respjson.Field
		Sen2lat               respjson.Field
		Sen2lon               respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		Sensor1Delay          respjson.Field
		Sensor2Delay          respjson.Field
		Snr                   respjson.Field
		SourceDl              respjson.Field
		TaskID                respjson.Field
		Tdoa                  respjson.Field
		TdoaAmb               respjson.Field
		TdoaUnc               respjson.Field
		TransactionID         respjson.Field
		Uct                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiffofarrivalAbridged) RawJSON() string { return r.JSON.raw }
func (r *DiffofarrivalAbridged) UnmarshalJSON(data []byte) error {
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
type DiffofarrivalAbridgedDataMode string

const (
	DiffofarrivalAbridgedDataModeReal      DiffofarrivalAbridgedDataMode = "REAL"
	DiffofarrivalAbridgedDataModeTest      DiffofarrivalAbridgedDataMode = "TEST"
	DiffofarrivalAbridgedDataModeSimulated DiffofarrivalAbridgedDataMode = "SIMULATED"
	DiffofarrivalAbridgedDataModeExercise  DiffofarrivalAbridgedDataMode = "EXERCISE"
)

// Model representation of Signal time and frequency difference of arrival
// (TDOA/FDOA) information obtained by using passive RF based sensor
// phenomenologies and sensor triangulation.
type DiffofarrivalFull struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
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
	DataMode DiffofarrivalFullDataMode `json:"dataMode" api:"required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime" api:"required" format:"date-time"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Line of sight azimuth angle for sensor 1, in degrees and topocentric frame.
	// Azimuth ranges from 0 to 360 degrees.
	Azimuth1 float64 `json:"azimuth1"`
	// Line of sight azimuth angle for sensor 2, in degrees and topocentric frame.
	// Azimuth ranges from 0 to 360 degrees.
	Azimuth2 float64 `json:"azimuth2"`
	// Bandwidth of the signal, in hertz.
	Bandwidth float64 `json:"bandwidth"`
	// Collection mode (e.g., DIRECTED_SEARCH, MANUAL, NEIGHBORHOOD_WATCH, SPOT_SEARCH,
	// SURVEY, etc.).
	CollectionMode string `json:"collectionMode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Delta range, in kilometers. Delta range calculation convention is (sensor2 -
	// sensor1).
	DeltaRange float64 `json:"deltaRange"`
	// Delta range rate, in kilometers per second. Delta range rate calculation
	// convention is (sensor2 - sensor1).
	DeltaRangeRate float64 `json:"deltaRangeRate"`
	// One sigma uncertainty in the delta range rate, in kilometers per second.
	DeltaRangeRateUnc float64 `json:"deltaRangeRateUnc"`
	// One sigma uncertainty in delta range, in kilometers.
	DeltaRangeUnc float64 `json:"deltaRangeUnc"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Line of sight elevation angle for sensor 1, in degrees and topocentric frame.
	// Elevation ranges from -90 to 90 degrees.
	Elevation1 float64 `json:"elevation1"`
	// Line of sight elevation angle for sensor 2, in degrees and topocentric frame.
	// Elevation ranges from -90 to 90 degrees.
	Elevation2 float64 `json:"elevation2"`
	// Frequency difference of arrival of the center frequency signal, in hertz. FDOA
	// calculation convention is (sensor2 - sensor1).
	Fdoa float64 `json:"fdoa"`
	// One sigma uncertainty in frequency difference of arrival of the center frequency
	// signal, in hertz.
	FdoaUnc float64 `json:"fdoaUnc"`
	// Center frequency of the collect, in hertz.
	Frequency float64 `json:"frequency"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Sensor ID of the primary/1st sensor used for this measurement.
	IDSensor1 string `json:"idSensor1"`
	// Sensor ID of the secondary/2nd sensor used for this measurement.
	IDSensor2 string `json:"idSensor2"`
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
	// Optional identifier provided by DOA source to indicate the primary/1st sensor
	// identifier used for this measurement. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorId1 string `json:"origSensorId1"`
	// Optional identifier provided by DOA source to indicate the secondary/2nd sensor
	// identifier used for this this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorId2 string `json:"origSensorId2"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	PolarityType string `json:"polarityType"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value. This field should not be
	// used unless coordinated with the UDL onboarding team.
	RawFileUri string `json:"rawFileURI"`
	// Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sensor 2 altitude at obTime (if mobile/onorbit), in kilometers. If null, can be
	// obtained from sensor info.
	Sen2alt float64 `json:"sen2alt"`
	// Sensor 2 WGS84 latitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -90 to 90 degrees (negative values south of
	// equator).
	Sen2lat float64 `json:"sen2lat"`
	// Sensor 2 WGS84 longitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Sen2lon float64 `json:"sen2lon"`
	// Sensor altitude at obTime (if mobile/onorbit), in kilometers. If null, can be
	// obtained from sensor info.
	Senalt float64 `json:"senalt"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit), in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon float64 `json:"senlon"`
	// The signal arrival delay relative to sensor 1, in seconds.
	Sensor1Delay float64 `json:"sensor1Delay"`
	// The signal arrival delay relative to sensor 2, in seconds.
	Sensor2Delay float64 `json:"sensor2Delay"`
	// Signal to noise ratio, in decibels.
	Snr float64 `json:"snr"`
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
	// Time difference of arrival of the center frequency signal, in seconds. TDOA
	// calculation convention is (sensor2 - sensor1).
	Tdoa float64 `json:"tdoa"`
	// The number of full signal cycles the time difference of arrival measurement
	// could be shifted due to repeating patterns in the signal (0 = no ambiguity, +1 =
	// one positive cycle ambiguity, -1 = one negative cycle ambiguity).
	TdoaAmb int64 `json:"tdoaAmb"`
	// One sigma uncertainty in time difference of arrival of the center frequency
	// signal, in seconds.
	TdoaUnc float64 `json:"tdoaUnc"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct bool `json:"uct"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Azimuth1              respjson.Field
		Azimuth2              respjson.Field
		Bandwidth             respjson.Field
		CollectionMode        respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeltaRange            respjson.Field
		DeltaRangeRate        respjson.Field
		DeltaRangeRateUnc     respjson.Field
		DeltaRangeUnc         respjson.Field
		Descriptor            respjson.Field
		Elevation1            respjson.Field
		Elevation2            respjson.Field
		Fdoa                  respjson.Field
		FdoaUnc               respjson.Field
		Frequency             respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor1             respjson.Field
		IDSensor2             respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorId1         respjson.Field
		OrigSensorId2         respjson.Field
		PolarityType          respjson.Field
		RawFileUri            respjson.Field
		SatNo                 respjson.Field
		Sen2alt               respjson.Field
		Sen2lat               respjson.Field
		Sen2lon               respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		Sensor1Delay          respjson.Field
		Sensor2Delay          respjson.Field
		Snr                   respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		TaskID                respjson.Field
		Tdoa                  respjson.Field
		TdoaAmb               respjson.Field
		TdoaUnc               respjson.Field
		TransactionID         respjson.Field
		Uct                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiffofarrivalFull) RawJSON() string { return r.JSON.raw }
func (r *DiffofarrivalFull) UnmarshalJSON(data []byte) error {
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
type DiffofarrivalFullDataMode string

const (
	DiffofarrivalFullDataModeReal      DiffofarrivalFullDataMode = "REAL"
	DiffofarrivalFullDataModeTest      DiffofarrivalFullDataMode = "TEST"
	DiffofarrivalFullDataModeSimulated DiffofarrivalFullDataMode = "SIMULATED"
	DiffofarrivalFullDataModeExercise  DiffofarrivalFullDataMode = "EXERCISE"
)

type TdoaFdoaDiffofarrivalNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
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
	DataMode TdoaFdoaDiffofarrivalNewParamsDataMode `json:"dataMode,omitzero" api:"required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime" api:"required" format:"date-time"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Line of sight azimuth angle for sensor 1, in degrees and topocentric frame.
	// Azimuth ranges from 0 to 360 degrees.
	Azimuth1 param.Opt[float64] `json:"azimuth1,omitzero"`
	// Line of sight azimuth angle for sensor 2, in degrees and topocentric frame.
	// Azimuth ranges from 0 to 360 degrees.
	Azimuth2 param.Opt[float64] `json:"azimuth2,omitzero"`
	// Bandwidth of the signal, in hertz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Collection mode (e.g., DIRECTED_SEARCH, MANUAL, NEIGHBORHOOD_WATCH, SPOT_SEARCH,
	// SURVEY, etc.).
	CollectionMode param.Opt[string] `json:"collectionMode,omitzero"`
	// Delta range, in kilometers. Delta range calculation convention is (sensor2 -
	// sensor1).
	DeltaRange param.Opt[float64] `json:"deltaRange,omitzero"`
	// Delta range rate, in kilometers per second. Delta range rate calculation
	// convention is (sensor2 - sensor1).
	DeltaRangeRate param.Opt[float64] `json:"deltaRangeRate,omitzero"`
	// One sigma uncertainty in the delta range rate, in kilometers per second.
	DeltaRangeRateUnc param.Opt[float64] `json:"deltaRangeRateUnc,omitzero"`
	// One sigma uncertainty in delta range, in kilometers.
	DeltaRangeUnc param.Opt[float64] `json:"deltaRangeUnc,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Line of sight elevation angle for sensor 1, in degrees and topocentric frame.
	// Elevation ranges from -90 to 90 degrees.
	Elevation1 param.Opt[float64] `json:"elevation1,omitzero"`
	// Line of sight elevation angle for sensor 2, in degrees and topocentric frame.
	// Elevation ranges from -90 to 90 degrees.
	Elevation2 param.Opt[float64] `json:"elevation2,omitzero"`
	// Frequency difference of arrival of the center frequency signal, in hertz. FDOA
	// calculation convention is (sensor2 - sensor1).
	Fdoa param.Opt[float64] `json:"fdoa,omitzero"`
	// One sigma uncertainty in frequency difference of arrival of the center frequency
	// signal, in hertz.
	FdoaUnc param.Opt[float64] `json:"fdoaUnc,omitzero"`
	// Center frequency of the collect, in hertz.
	Frequency param.Opt[float64] `json:"frequency,omitzero"`
	// Sensor ID of the primary/1st sensor used for this measurement.
	IDSensor1 param.Opt[string] `json:"idSensor1,omitzero"`
	// Sensor ID of the secondary/2nd sensor used for this measurement.
	IDSensor2 param.Opt[string] `json:"idSensor2,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
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
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	PolarityType param.Opt[string] `json:"polarityType,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value. This field should not be
	// used unless coordinated with the UDL onboarding team.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor 2 altitude at obTime (if mobile/onorbit), in kilometers. If null, can be
	// obtained from sensor info.
	Sen2alt param.Opt[float64] `json:"sen2alt,omitzero"`
	// Sensor 2 WGS84 latitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -90 to 90 degrees (negative values south of
	// equator).
	Sen2lat param.Opt[float64] `json:"sen2lat,omitzero"`
	// Sensor 2 WGS84 longitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Sen2lon param.Opt[float64] `json:"sen2lon,omitzero"`
	// Sensor altitude at obTime (if mobile/onorbit), in kilometers. If null, can be
	// obtained from sensor info.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit), in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// The signal arrival delay relative to sensor 1, in seconds.
	Sensor1Delay param.Opt[float64] `json:"sensor1Delay,omitzero"`
	// The signal arrival delay relative to sensor 2, in seconds.
	Sensor2Delay param.Opt[float64] `json:"sensor2Delay,omitzero"`
	// Signal to noise ratio, in decibels.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Time difference of arrival of the center frequency signal, in seconds. TDOA
	// calculation convention is (sensor2 - sensor1).
	Tdoa param.Opt[float64] `json:"tdoa,omitzero"`
	// The number of full signal cycles the time difference of arrival measurement
	// could be shifted due to repeating patterns in the signal (0 = no ambiguity, +1 =
	// one positive cycle ambiguity, -1 = one negative cycle ambiguity).
	TdoaAmb param.Opt[int64] `json:"tdoaAmb,omitzero"`
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

func (r TdoaFdoaDiffofarrivalNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TdoaFdoaDiffofarrivalNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TdoaFdoaDiffofarrivalNewParams) UnmarshalJSON(data []byte) error {
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
type TdoaFdoaDiffofarrivalNewParamsDataMode string

const (
	TdoaFdoaDiffofarrivalNewParamsDataModeReal      TdoaFdoaDiffofarrivalNewParamsDataMode = "REAL"
	TdoaFdoaDiffofarrivalNewParamsDataModeTest      TdoaFdoaDiffofarrivalNewParamsDataMode = "TEST"
	TdoaFdoaDiffofarrivalNewParamsDataModeSimulated TdoaFdoaDiffofarrivalNewParamsDataMode = "SIMULATED"
	TdoaFdoaDiffofarrivalNewParamsDataModeExercise  TdoaFdoaDiffofarrivalNewParamsDataMode = "EXERCISE"
)

type TdoaFdoaDiffofarrivalListParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime" api:"required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TdoaFdoaDiffofarrivalListParams]'s query parameters as
// `url.Values`.
func (r TdoaFdoaDiffofarrivalListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TdoaFdoaDiffofarrivalCountParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime" api:"required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TdoaFdoaDiffofarrivalCountParams]'s query parameters as
// `url.Values`.
func (r TdoaFdoaDiffofarrivalCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TdoaFdoaDiffofarrivalNewBulkParams struct {
	Body []TdoaFdoaDiffofarrivalNewBulkParamsBody
	paramObj
}

func (r TdoaFdoaDiffofarrivalNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *TdoaFdoaDiffofarrivalNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Model representation of Signal time and frequency difference of arrival
// (TDOA/FDOA) information obtained by using passive RF based sensor
// phenomenologies and sensor triangulation.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type TdoaFdoaDiffofarrivalNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
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
	DataMode string `json:"dataMode,omitzero" api:"required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime" api:"required" format:"date-time"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Line of sight azimuth angle for sensor 1, in degrees and topocentric frame.
	// Azimuth ranges from 0 to 360 degrees.
	Azimuth1 param.Opt[float64] `json:"azimuth1,omitzero"`
	// Line of sight azimuth angle for sensor 2, in degrees and topocentric frame.
	// Azimuth ranges from 0 to 360 degrees.
	Azimuth2 param.Opt[float64] `json:"azimuth2,omitzero"`
	// Bandwidth of the signal, in hertz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Collection mode (e.g., DIRECTED_SEARCH, MANUAL, NEIGHBORHOOD_WATCH, SPOT_SEARCH,
	// SURVEY, etc.).
	CollectionMode param.Opt[string] `json:"collectionMode,omitzero"`
	// Delta range, in kilometers. Delta range calculation convention is (sensor2 -
	// sensor1).
	DeltaRange param.Opt[float64] `json:"deltaRange,omitzero"`
	// Delta range rate, in kilometers per second. Delta range rate calculation
	// convention is (sensor2 - sensor1).
	DeltaRangeRate param.Opt[float64] `json:"deltaRangeRate,omitzero"`
	// One sigma uncertainty in the delta range rate, in kilometers per second.
	DeltaRangeRateUnc param.Opt[float64] `json:"deltaRangeRateUnc,omitzero"`
	// One sigma uncertainty in delta range, in kilometers.
	DeltaRangeUnc param.Opt[float64] `json:"deltaRangeUnc,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Line of sight elevation angle for sensor 1, in degrees and topocentric frame.
	// Elevation ranges from -90 to 90 degrees.
	Elevation1 param.Opt[float64] `json:"elevation1,omitzero"`
	// Line of sight elevation angle for sensor 2, in degrees and topocentric frame.
	// Elevation ranges from -90 to 90 degrees.
	Elevation2 param.Opt[float64] `json:"elevation2,omitzero"`
	// Frequency difference of arrival of the center frequency signal, in hertz. FDOA
	// calculation convention is (sensor2 - sensor1).
	Fdoa param.Opt[float64] `json:"fdoa,omitzero"`
	// One sigma uncertainty in frequency difference of arrival of the center frequency
	// signal, in hertz.
	FdoaUnc param.Opt[float64] `json:"fdoaUnc,omitzero"`
	// Center frequency of the collect, in hertz.
	Frequency param.Opt[float64] `json:"frequency,omitzero"`
	// Sensor ID of the primary/1st sensor used for this measurement.
	IDSensor1 param.Opt[string] `json:"idSensor1,omitzero"`
	// Sensor ID of the secondary/2nd sensor used for this measurement.
	IDSensor2 param.Opt[string] `json:"idSensor2,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
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
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	PolarityType param.Opt[string] `json:"polarityType,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value. This field should not be
	// used unless coordinated with the UDL onboarding team.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor 2 altitude at obTime (if mobile/onorbit), in kilometers. If null, can be
	// obtained from sensor info.
	Sen2alt param.Opt[float64] `json:"sen2alt,omitzero"`
	// Sensor 2 WGS84 latitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -90 to 90 degrees (negative values south of
	// equator).
	Sen2lat param.Opt[float64] `json:"sen2lat,omitzero"`
	// Sensor 2 WGS84 longitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Sen2lon param.Opt[float64] `json:"sen2lon,omitzero"`
	// Sensor altitude at obTime (if mobile/onorbit), in kilometers. If null, can be
	// obtained from sensor info.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit), in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit), in degrees. If null, can
	// be obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// The signal arrival delay relative to sensor 1, in seconds.
	Sensor1Delay param.Opt[float64] `json:"sensor1Delay,omitzero"`
	// The signal arrival delay relative to sensor 2, in seconds.
	Sensor2Delay param.Opt[float64] `json:"sensor2Delay,omitzero"`
	// Signal to noise ratio, in decibels.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Time difference of arrival of the center frequency signal, in seconds. TDOA
	// calculation convention is (sensor2 - sensor1).
	Tdoa param.Opt[float64] `json:"tdoa,omitzero"`
	// The number of full signal cycles the time difference of arrival measurement
	// could be shifted due to repeating patterns in the signal (0 = no ambiguity, +1 =
	// one positive cycle ambiguity, -1 = one negative cycle ambiguity).
	TdoaAmb param.Opt[int64] `json:"tdoaAmb,omitzero"`
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

func (r TdoaFdoaDiffofarrivalNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow TdoaFdoaDiffofarrivalNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TdoaFdoaDiffofarrivalNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TdoaFdoaDiffofarrivalNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
