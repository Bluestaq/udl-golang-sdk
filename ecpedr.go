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

// EcpedrService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEcpedrService] method instead.
type EcpedrService struct {
	Options []option.RequestOption
	History EcpedrHistoryService
}

// NewEcpedrService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEcpedrService(opts ...option.RequestOption) (r EcpedrService) {
	r = EcpedrService{}
	r.Options = opts
	r.History = NewEcpedrHistoryService(opts...)
	return
}

// Service operation to take a single ECPEDR record as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *EcpedrService) New(ctx context.Context, body EcpedrNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/ecpedr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EcpedrService) List(ctx context.Context, query EcpedrListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EcpedrListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/ecpedr"
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
func (r *EcpedrService) ListAutoPaging(ctx context.Context, query EcpedrListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EcpedrListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EcpedrService) Count(ctx context.Context, query EcpedrCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/ecpedr/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// ECPEDR records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *EcpedrService) NewBulk(ctx context.Context, body EcpedrNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/ecpedr/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EcpedrService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *EcpedrQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/ecpedr/queryhelp"
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
func (r *EcpedrService) Tuple(ctx context.Context, query EcpedrTupleParams, opts ...option.RequestOption) (res *[]EcpedrTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/ecpedr/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple ECPEDR records as a POST body and ingest into
// the database. This operation is intended to be used for automated feeds into
// UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *EcpedrService) UnvalidatedPublish(ctx context.Context, body EcpedrUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-ecpedr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Energetic Charged Particles (ECP) Environmental Data Records (EDRs).
type EcpedrListResponse struct {
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
	DataMode EcpedrListResponseDataMode `json:"dataMode,required"`
	// Collection of measurements associated with this ECP EDR record.
	EcpedrMeasurements []EcpedrListResponseEcpedrMeasurement `json:"ecpedrMeasurements,required"`
	// Time of the observation, in ISO 8601 UTC format with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// System which generated the message.
	GenSystem string `json:"genSystem"`
	// Time when message was generated in ISO 8601 UTC format with millisecond
	// precision.
	GenTime time.Time `json:"genTime" format:"date-time"`
	// Unique identifier of the on-orbit satellite hosting the sensor.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensor string `json:"idSensor"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo int64 `json:"satNo"`
	// Three element array, expressing the observing spacecraft/sensor position vector
	// components at observation time, in kilometers, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xpos, ypos, zpos].
	SenPos []float64 `json:"senPos"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame EcpedrListResponseSenReferenceFrame `json:"senReferenceFrame"`
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
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EcpedrMeasurements    respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		GenSystem             respjson.Field
		GenTime               respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		SatNo                 respjson.Field
		SenPos                respjson.Field
		SenReferenceFrame     respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EcpedrListResponse) RawJSON() string { return r.JSON.raw }
func (r *EcpedrListResponse) UnmarshalJSON(data []byte) error {
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
type EcpedrListResponseDataMode string

const (
	EcpedrListResponseDataModeReal      EcpedrListResponseDataMode = "REAL"
	EcpedrListResponseDataModeTest      EcpedrListResponseDataMode = "TEST"
	EcpedrListResponseDataModeSimulated EcpedrListResponseDataMode = "SIMULATED"
	EcpedrListResponseDataModeExercise  EcpedrListResponseDataMode = "EXERCISE"
)

// Collection of measurements associated with this ECP EDR record.
type EcpedrListResponseEcpedrMeasurement struct {
	// The type of observation associated with this record. (e.g., FLUX, CHARGE, etc.).
	ObType string `json:"obType,required"`
	// The Unit of Measure associated with this observation. If there are no physical
	// units associated with the measurement, a value of NONE should be specified.
	ObUoM string `json:"obUoM,required"`
	// Higher energy threshold of the channel for event detection and data collection.
	ChanEnergyHigh float64 `json:"chanEnergyHigh"`
	// Lower energy threshold of the channel for event detection and data collection.
	ChanEnergyLow float64 `json:"chanEnergyLow"`
	// Identifier of the channel based on energy levels and particle species.
	ChanID string `json:"chanId"`
	// Type of channel based on the measurement method (e.g., INTEGRAL, DIFFERENTIAL,
	// etc.).
	ChanType string `json:"chanType"`
	// Units used for defining channel energy boundaries (e.g., eV, keV, MeV, etc.).
	ChanUnit string `json:"chanUnit"`
	// Designates a specific group of measurements made.
	MsgNumber int64 `json:"msgNumber"`
	// A single observation value expressed in the specified unit of measure (obUoM).
	ObValue float64 `json:"obValue"`
	// Type of particle species being measured by a channel (e.g., ELECTRON, PROTON,
	// etc.).
	Species string `json:"species"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObType         respjson.Field
		ObUoM          respjson.Field
		ChanEnergyHigh respjson.Field
		ChanEnergyLow  respjson.Field
		ChanID         respjson.Field
		ChanType       respjson.Field
		ChanUnit       respjson.Field
		MsgNumber      respjson.Field
		ObValue        respjson.Field
		Species        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EcpedrListResponseEcpedrMeasurement) RawJSON() string { return r.JSON.raw }
func (r *EcpedrListResponseEcpedrMeasurement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type EcpedrListResponseSenReferenceFrame string

const (
	EcpedrListResponseSenReferenceFrameJ2000   EcpedrListResponseSenReferenceFrame = "J2000"
	EcpedrListResponseSenReferenceFrameEfgTdr  EcpedrListResponseSenReferenceFrame = "EFG/TDR"
	EcpedrListResponseSenReferenceFrameEcrEcef EcpedrListResponseSenReferenceFrame = "ECR/ECEF"
	EcpedrListResponseSenReferenceFrameTeme    EcpedrListResponseSenReferenceFrame = "TEME"
	EcpedrListResponseSenReferenceFrameItrf    EcpedrListResponseSenReferenceFrame = "ITRF"
	EcpedrListResponseSenReferenceFrameGcrf    EcpedrListResponseSenReferenceFrame = "GCRF"
)

type EcpedrQueryhelpResponse struct {
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
func (r EcpedrQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *EcpedrQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Energetic Charged Particles (ECP) Environmental Data Records (EDRs).
type EcpedrTupleResponse struct {
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
	DataMode EcpedrTupleResponseDataMode `json:"dataMode,required"`
	// Collection of measurements associated with this ECP EDR record.
	EcpedrMeasurements []EcpedrTupleResponseEcpedrMeasurement `json:"ecpedrMeasurements,required"`
	// Time of the observation, in ISO 8601 UTC format with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// System which generated the message.
	GenSystem string `json:"genSystem"`
	// Time when message was generated in ISO 8601 UTC format with millisecond
	// precision.
	GenTime time.Time `json:"genTime" format:"date-time"`
	// Unique identifier of the on-orbit satellite hosting the sensor.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensor string `json:"idSensor"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo int64 `json:"satNo"`
	// Three element array, expressing the observing spacecraft/sensor position vector
	// components at observation time, in kilometers, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xpos, ypos, zpos].
	SenPos []float64 `json:"senPos"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame EcpedrTupleResponseSenReferenceFrame `json:"senReferenceFrame"`
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
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EcpedrMeasurements    respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		GenSystem             respjson.Field
		GenTime               respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		SatNo                 respjson.Field
		SenPos                respjson.Field
		SenReferenceFrame     respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EcpedrTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *EcpedrTupleResponse) UnmarshalJSON(data []byte) error {
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
type EcpedrTupleResponseDataMode string

const (
	EcpedrTupleResponseDataModeReal      EcpedrTupleResponseDataMode = "REAL"
	EcpedrTupleResponseDataModeTest      EcpedrTupleResponseDataMode = "TEST"
	EcpedrTupleResponseDataModeSimulated EcpedrTupleResponseDataMode = "SIMULATED"
	EcpedrTupleResponseDataModeExercise  EcpedrTupleResponseDataMode = "EXERCISE"
)

// Collection of measurements associated with this ECP EDR record.
type EcpedrTupleResponseEcpedrMeasurement struct {
	// The type of observation associated with this record. (e.g., FLUX, CHARGE, etc.).
	ObType string `json:"obType,required"`
	// The Unit of Measure associated with this observation. If there are no physical
	// units associated with the measurement, a value of NONE should be specified.
	ObUoM string `json:"obUoM,required"`
	// Higher energy threshold of the channel for event detection and data collection.
	ChanEnergyHigh float64 `json:"chanEnergyHigh"`
	// Lower energy threshold of the channel for event detection and data collection.
	ChanEnergyLow float64 `json:"chanEnergyLow"`
	// Identifier of the channel based on energy levels and particle species.
	ChanID string `json:"chanId"`
	// Type of channel based on the measurement method (e.g., INTEGRAL, DIFFERENTIAL,
	// etc.).
	ChanType string `json:"chanType"`
	// Units used for defining channel energy boundaries (e.g., eV, keV, MeV, etc.).
	ChanUnit string `json:"chanUnit"`
	// Designates a specific group of measurements made.
	MsgNumber int64 `json:"msgNumber"`
	// A single observation value expressed in the specified unit of measure (obUoM).
	ObValue float64 `json:"obValue"`
	// Type of particle species being measured by a channel (e.g., ELECTRON, PROTON,
	// etc.).
	Species string `json:"species"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObType         respjson.Field
		ObUoM          respjson.Field
		ChanEnergyHigh respjson.Field
		ChanEnergyLow  respjson.Field
		ChanID         respjson.Field
		ChanType       respjson.Field
		ChanUnit       respjson.Field
		MsgNumber      respjson.Field
		ObValue        respjson.Field
		Species        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EcpedrTupleResponseEcpedrMeasurement) RawJSON() string { return r.JSON.raw }
func (r *EcpedrTupleResponseEcpedrMeasurement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type EcpedrTupleResponseSenReferenceFrame string

const (
	EcpedrTupleResponseSenReferenceFrameJ2000   EcpedrTupleResponseSenReferenceFrame = "J2000"
	EcpedrTupleResponseSenReferenceFrameEfgTdr  EcpedrTupleResponseSenReferenceFrame = "EFG/TDR"
	EcpedrTupleResponseSenReferenceFrameEcrEcef EcpedrTupleResponseSenReferenceFrame = "ECR/ECEF"
	EcpedrTupleResponseSenReferenceFrameTeme    EcpedrTupleResponseSenReferenceFrame = "TEME"
	EcpedrTupleResponseSenReferenceFrameItrf    EcpedrTupleResponseSenReferenceFrame = "ITRF"
	EcpedrTupleResponseSenReferenceFrameGcrf    EcpedrTupleResponseSenReferenceFrame = "GCRF"
)

type EcpedrNewParams struct {
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
	DataMode EcpedrNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Collection of measurements associated with this ECP EDR record.
	EcpedrMeasurements []EcpedrNewParamsEcpedrMeasurement `json:"ecpedrMeasurements,omitzero,required"`
	// Time of the observation, in ISO 8601 UTC format with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// System which generated the message.
	GenSystem param.Opt[string] `json:"genSystem,omitzero"`
	// Time when message was generated in ISO 8601 UTC format with millisecond
	// precision.
	GenTime param.Opt[time.Time] `json:"genTime,omitzero" format:"date-time"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Three element array, expressing the observing spacecraft/sensor position vector
	// components at observation time, in kilometers, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xpos, ypos, zpos].
	SenPos []float64 `json:"senPos,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame EcpedrNewParamsSenReferenceFrame `json:"senReferenceFrame,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r EcpedrNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EcpedrNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EcpedrNewParams) UnmarshalJSON(data []byte) error {
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
type EcpedrNewParamsDataMode string

const (
	EcpedrNewParamsDataModeReal      EcpedrNewParamsDataMode = "REAL"
	EcpedrNewParamsDataModeTest      EcpedrNewParamsDataMode = "TEST"
	EcpedrNewParamsDataModeSimulated EcpedrNewParamsDataMode = "SIMULATED"
	EcpedrNewParamsDataModeExercise  EcpedrNewParamsDataMode = "EXERCISE"
)

// Collection of measurements associated with this ECP EDR record.
//
// The properties ObType, ObUoM are required.
type EcpedrNewParamsEcpedrMeasurement struct {
	// The type of observation associated with this record. (e.g., FLUX, CHARGE, etc.).
	ObType string `json:"obType,required"`
	// The Unit of Measure associated with this observation. If there are no physical
	// units associated with the measurement, a value of NONE should be specified.
	ObUoM string `json:"obUoM,required"`
	// Higher energy threshold of the channel for event detection and data collection.
	ChanEnergyHigh param.Opt[float64] `json:"chanEnergyHigh,omitzero"`
	// Lower energy threshold of the channel for event detection and data collection.
	ChanEnergyLow param.Opt[float64] `json:"chanEnergyLow,omitzero"`
	// Identifier of the channel based on energy levels and particle species.
	ChanID param.Opt[string] `json:"chanId,omitzero"`
	// Type of channel based on the measurement method (e.g., INTEGRAL, DIFFERENTIAL,
	// etc.).
	ChanType param.Opt[string] `json:"chanType,omitzero"`
	// Units used for defining channel energy boundaries (e.g., eV, keV, MeV, etc.).
	ChanUnit param.Opt[string] `json:"chanUnit,omitzero"`
	// Designates a specific group of measurements made.
	MsgNumber param.Opt[int64] `json:"msgNumber,omitzero"`
	// A single observation value expressed in the specified unit of measure (obUoM).
	ObValue param.Opt[float64] `json:"obValue,omitzero"`
	// Type of particle species being measured by a channel (e.g., ELECTRON, PROTON,
	// etc.).
	Species param.Opt[string] `json:"species,omitzero"`
	paramObj
}

func (r EcpedrNewParamsEcpedrMeasurement) MarshalJSON() (data []byte, err error) {
	type shadow EcpedrNewParamsEcpedrMeasurement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EcpedrNewParamsEcpedrMeasurement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type EcpedrNewParamsSenReferenceFrame string

const (
	EcpedrNewParamsSenReferenceFrameJ2000   EcpedrNewParamsSenReferenceFrame = "J2000"
	EcpedrNewParamsSenReferenceFrameEfgTdr  EcpedrNewParamsSenReferenceFrame = "EFG/TDR"
	EcpedrNewParamsSenReferenceFrameEcrEcef EcpedrNewParamsSenReferenceFrame = "ECR/ECEF"
	EcpedrNewParamsSenReferenceFrameTeme    EcpedrNewParamsSenReferenceFrame = "TEME"
	EcpedrNewParamsSenReferenceFrameItrf    EcpedrNewParamsSenReferenceFrame = "ITRF"
	EcpedrNewParamsSenReferenceFrameGcrf    EcpedrNewParamsSenReferenceFrame = "GCRF"
)

type EcpedrListParams struct {
	// Time of the observation, in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EcpedrListParams]'s query parameters as `url.Values`.
func (r EcpedrListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EcpedrCountParams struct {
	// Time of the observation, in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EcpedrCountParams]'s query parameters as `url.Values`.
func (r EcpedrCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EcpedrNewBulkParams struct {
	Body []EcpedrNewBulkParamsBody
	paramObj
}

func (r EcpedrNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *EcpedrNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Energetic Charged Particles (ECP) Environmental Data Records (EDRs).
//
// The properties ClassificationMarking, DataMode, EcpedrMeasurements, ObTime,
// Source are required.
type EcpedrNewBulkParamsBody struct {
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
	// Collection of measurements associated with this ECP EDR record.
	EcpedrMeasurements []EcpedrNewBulkParamsBodyEcpedrMeasurement `json:"ecpedrMeasurements,omitzero,required"`
	// Time of the observation, in ISO 8601 UTC format with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// System which generated the message.
	GenSystem param.Opt[string] `json:"genSystem,omitzero"`
	// Time when message was generated in ISO 8601 UTC format with millisecond
	// precision.
	GenTime param.Opt[time.Time] `json:"genTime,omitzero" format:"date-time"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Three element array, expressing the observing spacecraft/sensor position vector
	// components at observation time, in kilometers, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xpos, ypos, zpos].
	SenPos []float64 `json:"senPos,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame string `json:"senReferenceFrame,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r EcpedrNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EcpedrNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EcpedrNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EcpedrNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[EcpedrNewBulkParamsBody](
		"senReferenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// Collection of measurements associated with this ECP EDR record.
//
// The properties ObType, ObUoM are required.
type EcpedrNewBulkParamsBodyEcpedrMeasurement struct {
	// The type of observation associated with this record. (e.g., FLUX, CHARGE, etc.).
	ObType string `json:"obType,required"`
	// The Unit of Measure associated with this observation. If there are no physical
	// units associated with the measurement, a value of NONE should be specified.
	ObUoM string `json:"obUoM,required"`
	// Higher energy threshold of the channel for event detection and data collection.
	ChanEnergyHigh param.Opt[float64] `json:"chanEnergyHigh,omitzero"`
	// Lower energy threshold of the channel for event detection and data collection.
	ChanEnergyLow param.Opt[float64] `json:"chanEnergyLow,omitzero"`
	// Identifier of the channel based on energy levels and particle species.
	ChanID param.Opt[string] `json:"chanId,omitzero"`
	// Type of channel based on the measurement method (e.g., INTEGRAL, DIFFERENTIAL,
	// etc.).
	ChanType param.Opt[string] `json:"chanType,omitzero"`
	// Units used for defining channel energy boundaries (e.g., eV, keV, MeV, etc.).
	ChanUnit param.Opt[string] `json:"chanUnit,omitzero"`
	// Designates a specific group of measurements made.
	MsgNumber param.Opt[int64] `json:"msgNumber,omitzero"`
	// A single observation value expressed in the specified unit of measure (obUoM).
	ObValue param.Opt[float64] `json:"obValue,omitzero"`
	// Type of particle species being measured by a channel (e.g., ELECTRON, PROTON,
	// etc.).
	Species param.Opt[string] `json:"species,omitzero"`
	paramObj
}

func (r EcpedrNewBulkParamsBodyEcpedrMeasurement) MarshalJSON() (data []byte, err error) {
	type shadow EcpedrNewBulkParamsBodyEcpedrMeasurement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EcpedrNewBulkParamsBodyEcpedrMeasurement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EcpedrTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time of the observation, in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EcpedrTupleParams]'s query parameters as `url.Values`.
func (r EcpedrTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EcpedrUnvalidatedPublishParams struct {
	Body []EcpedrUnvalidatedPublishParamsBody
	paramObj
}

func (r EcpedrUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *EcpedrUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Energetic Charged Particles (ECP) Environmental Data Records (EDRs).
//
// The properties ClassificationMarking, DataMode, EcpedrMeasurements, ObTime,
// Source are required.
type EcpedrUnvalidatedPublishParamsBody struct {
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
	// Collection of measurements associated with this ECP EDR record.
	EcpedrMeasurements []EcpedrUnvalidatedPublishParamsBodyEcpedrMeasurement `json:"ecpedrMeasurements,omitzero,required"`
	// Time of the observation, in ISO 8601 UTC format with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// System which generated the message.
	GenSystem param.Opt[string] `json:"genSystem,omitzero"`
	// Time when message was generated in ISO 8601 UTC format with millisecond
	// precision.
	GenTime param.Opt[time.Time] `json:"genTime,omitzero" format:"date-time"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Three element array, expressing the observing spacecraft/sensor position vector
	// components at observation time, in kilometers, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xpos, ypos, zpos].
	SenPos []float64 `json:"senPos,omitzero"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame string `json:"senReferenceFrame,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r EcpedrUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EcpedrUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EcpedrUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EcpedrUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[EcpedrUnvalidatedPublishParamsBody](
		"senReferenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// Collection of measurements associated with this ECP EDR record.
//
// The properties ObType, ObUoM are required.
type EcpedrUnvalidatedPublishParamsBodyEcpedrMeasurement struct {
	// The type of observation associated with this record. (e.g., FLUX, CHARGE, etc.).
	ObType string `json:"obType,required"`
	// The Unit of Measure associated with this observation. If there are no physical
	// units associated with the measurement, a value of NONE should be specified.
	ObUoM string `json:"obUoM,required"`
	// Higher energy threshold of the channel for event detection and data collection.
	ChanEnergyHigh param.Opt[float64] `json:"chanEnergyHigh,omitzero"`
	// Lower energy threshold of the channel for event detection and data collection.
	ChanEnergyLow param.Opt[float64] `json:"chanEnergyLow,omitzero"`
	// Identifier of the channel based on energy levels and particle species.
	ChanID param.Opt[string] `json:"chanId,omitzero"`
	// Type of channel based on the measurement method (e.g., INTEGRAL, DIFFERENTIAL,
	// etc.).
	ChanType param.Opt[string] `json:"chanType,omitzero"`
	// Units used for defining channel energy boundaries (e.g., eV, keV, MeV, etc.).
	ChanUnit param.Opt[string] `json:"chanUnit,omitzero"`
	// Designates a specific group of measurements made.
	MsgNumber param.Opt[int64] `json:"msgNumber,omitzero"`
	// A single observation value expressed in the specified unit of measure (obUoM).
	ObValue param.Opt[float64] `json:"obValue,omitzero"`
	// Type of particle species being measured by a channel (e.g., ELECTRON, PROTON,
	// etc.).
	Species param.Opt[string] `json:"species,omitzero"`
	paramObj
}

func (r EcpedrUnvalidatedPublishParamsBodyEcpedrMeasurement) MarshalJSON() (data []byte, err error) {
	type shadow EcpedrUnvalidatedPublishParamsBodyEcpedrMeasurement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EcpedrUnvalidatedPublishParamsBodyEcpedrMeasurement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
