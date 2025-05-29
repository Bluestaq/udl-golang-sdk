// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// SpaceEnvObservationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpaceEnvObservationService] method instead.
type SpaceEnvObservationService struct {
	Options []option.RequestOption
	History SpaceEnvObservationHistoryService
}

// NewSpaceEnvObservationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSpaceEnvObservationService(opts ...option.RequestOption) (r SpaceEnvObservationService) {
	r = SpaceEnvObservationService{}
	r.Options = opts
	r.History = NewSpaceEnvObservationHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SpaceEnvObservationService) List(ctx context.Context, query SpaceEnvObservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SpaceEnvObservationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/spaceenvobservation"
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
func (r *SpaceEnvObservationService) ListAutoPaging(ctx context.Context, query SpaceEnvObservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SpaceEnvObservationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SpaceEnvObservationService) Count(ctx context.Context, query SpaceEnvObservationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/spaceenvobservation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// SpaceEnvObservation records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *SpaceEnvObservationService) NewBulk(ctx context.Context, body SpaceEnvObservationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/spaceenvobservation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SpaceEnvObservationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SpaceEnvObservationQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/spaceenvobservation/queryhelp"
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
func (r *SpaceEnvObservationService) Tuple(ctx context.Context, query SpaceEnvObservationTupleParams, opts ...option.RequestOption) (res *[]SpaceEnvObservationFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/spaceenvobservation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to accept one or more SpaceEnvObservation(s) as a POST body
// and ingest into the database. This operation is intended to be used for
// automated feeds into UDL. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SpaceEnvObservationService) UnvalidatedPublish(ctx context.Context, body SpaceEnvObservationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-spaceenvobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// SpaceEnvObservation data.
type SpaceEnvObservationListResponse struct {
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
	DataMode SpaceEnvObservationListResponseDataMode `json:"dataMode,required"`
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Spacecraft/sensor altitude at observation time, expressed in kilometers above
	// WGS-84 ellipsoid.
	Alt float64 `json:"alt"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The data type (e.g. AP, AURORAL FLUX, ECP, KINDEX, PROPAGATED SOLAR WIND, XRAY
	// FLUX, etc.) of observations in this record.
	DataType string `json:"dataType"`
	// Flag indicating that this record contains derived data.
	Derived bool `json:"derived"`
	// Descriptive or additional information associated with this observation record.
	Description string `json:"description"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// Flag indicating that this record contains forecast data.
	Forecast bool `json:"forecast"`
	// The external system which generated the message, if applicable.
	GenSystem string `json:"genSystem"`
	// The time at which the associated data message was generated, in ISO 8601 UTC
	// format with millisecond precision.
	GenTime time.Time `json:"genTime" format:"date-time"`
	// Unique identifier of the on-orbit satellite hosting the sensor which produced
	// this data.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The type of instrument from which this data was collected (e.g. ANTENNA,
	// CHANNELTRON, INTERFEROMETER, MAGNETOMETER, RADIOMETER, etc.).
	InstrumentType string `json:"instrumentType"`
	// WGS-84 spacecraft/sensor latitude sub-point at observation time, represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat float64 `json:"lat"`
	// WGS-84 spacecraft/sensor longitude sub-point at observation time, represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The sensor measurement type of the observation data contained in this record.
	MeasType string `json:"measType"`
	// The type of message associated with this record.
	MsgType string `json:"msgType"`
	// The name of the observatory from which this data was collected.
	ObservatoryName string `json:"observatoryName"`
	// Additional notes concerning the observatory.
	ObservatoryNotes string `json:"observatoryNotes"`
	// The type of observatory from which this data was collected (e.g. FACILITY,
	// ONORBIT, NETWORK, etc.).
	ObservatoryType string `json:"observatoryType"`
	// A user-defined name or ID of a set of observations, if applicable. Used for
	// identifying multiple observation records as part of one observation set.
	ObSetID string `json:"obSetId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor which produced this data. This may be an internal identifier
	// and not necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// The particle type (AEROSOL, ALPHA PARTICLE, ATOM, DUST, ELECTRON, ION, MOLECULE,
	// NEUTRON, POSITRON, PROTON) associated with this measurement.
	ParticleType string `json:"particleType"`
	// The quality of the overall data contained in this record. The quality indicator
	// value may vary among providers and may be a generalized statement (BAD, GOOD,
	// UNCERTAIN, UNKNOWN) or a numeric value. Users should consult the data provider
	// to verify the usage of the quality indicator.
	Quality string `json:"quality"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor which
	// produced this data.
	SatNo int64 `json:"satNo"`
	// The energy level bin of the sensor associated with this measurement.
	SenEnergyLevel string `json:"senEnergyLevel"`
	// Three element array, expressing the observing spacecraft/sensor position vector
	// components at observation time, in kilometers, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xpos, ypos, zpos].
	SenPos []float64 `json:"senPos"`
	// The reference frame of the observing spacecraft/sensor state. If the
	// senReferenceFrame is null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame SpaceEnvObservationListResponseSenReferenceFrame `json:"senReferenceFrame"`
	// Three element array, expressing the observing spacecraft/sensor velocity vector
	// components at observation time, in kilometers/second, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xvel, yvel, zvel].
	SenVel []float64 `json:"senVel"`
	// A collection of individual space environment observations.
	SeoList []SpaceEnvObservationListResponseSeoList `json:"seoList"`
	// Array of UUIDs of the UDL data records that are related to this observation
	// record. See the associated 'srcTyps' array for specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// of the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SGI, SIGACT, SKYIMAGE, SPACEENVOB,
	// SV, TRACK) that are related to this observation record. See the associated
	// 'srcIds' array for the record UUIDs, positionally corresponding to the record
	// types in this array. The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Alt                   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataType              respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		Descriptor            respjson.Field
		ExternalID            respjson.Field
		Forecast              respjson.Field
		GenSystem             respjson.Field
		GenTime               respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		InstrumentType        respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		MeasType              respjson.Field
		MsgType               respjson.Field
		ObservatoryName       respjson.Field
		ObservatoryNotes      respjson.Field
		ObservatoryType       respjson.Field
		ObSetID               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		ParticleType          respjson.Field
		Quality               respjson.Field
		SatNo                 respjson.Field
		SenEnergyLevel        respjson.Field
		SenPos                respjson.Field
		SenReferenceFrame     respjson.Field
		SenVel                respjson.Field
		SeoList               respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpaceEnvObservationListResponse) RawJSON() string { return r.JSON.raw }
func (r *SpaceEnvObservationListResponse) UnmarshalJSON(data []byte) error {
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
type SpaceEnvObservationListResponseDataMode string

const (
	SpaceEnvObservationListResponseDataModeReal      SpaceEnvObservationListResponseDataMode = "REAL"
	SpaceEnvObservationListResponseDataModeTest      SpaceEnvObservationListResponseDataMode = "TEST"
	SpaceEnvObservationListResponseDataModeSimulated SpaceEnvObservationListResponseDataMode = "SIMULATED"
	SpaceEnvObservationListResponseDataModeExercise  SpaceEnvObservationListResponseDataMode = "EXERCISE"
)

// The reference frame of the observing spacecraft/sensor state. If the
// senReferenceFrame is null it is assumed to be J2000.
type SpaceEnvObservationListResponseSenReferenceFrame string

const (
	SpaceEnvObservationListResponseSenReferenceFrameJ2000   SpaceEnvObservationListResponseSenReferenceFrame = "J2000"
	SpaceEnvObservationListResponseSenReferenceFrameEfgTdr  SpaceEnvObservationListResponseSenReferenceFrame = "EFG/TDR"
	SpaceEnvObservationListResponseSenReferenceFrameEcrEcef SpaceEnvObservationListResponseSenReferenceFrame = "ECR/ECEF"
	SpaceEnvObservationListResponseSenReferenceFrameTeme    SpaceEnvObservationListResponseSenReferenceFrame = "TEME"
	SpaceEnvObservationListResponseSenReferenceFrameItrf    SpaceEnvObservationListResponseSenReferenceFrame = "ITRF"
	SpaceEnvObservationListResponseSenReferenceFrameGcrf    SpaceEnvObservationListResponseSenReferenceFrame = "GCRF"
)

// A single space environment observation.
type SpaceEnvObservationListResponseSeoList struct {
	// The type of observation associated with this record.
	ObType string `json:"obType,required"`
	// The Unit of Measure associated with this observation. If there are no physical
	// units associated with the measurement, a value of NONE should be specified.
	ObUoM string `json:"obUoM,required"`
	// An array of observation values expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Space Environment data types, each
	// record may employ a numeric observation value (obValue), a string observation
	// value (obString), a Boolean observation value (obBool), an array of numeric
	// observation values (obArray), or any combination of these.
	ObArray []float64 `json:"obArray"`
	// A Boolean observation. Because of the variability of the Space Environment data
	// types, each record may employ a numeric observation value (obValue), a string
	// observation value (obString), a Boolean observation value (obBool), an array of
	// numeric observation values (obArray), or any combination of these.
	ObBool bool `json:"obBool"`
	// Descriptive or additional information associated with this individual
	// observation.
	ObDescription string `json:"obDescription"`
	// The quality of this individual observation. The observation quality indicator
	// value may vary among providers and may be a generalized statement (BAD, GOOD,
	// UNCERTAIN, UNKNOWN) or a numeric value. Users should consult the data provider
	// to verify the usage of the observation.
	ObQuality string `json:"obQuality"`
	// A single observation string expressed in the specified unit of measure (obUoM).
	// Because of the variability of the Space Environment data types, each record may
	// employ a numeric observation value (obValue), a string observation value
	// (obString), a Boolean observation value (obBool), an array of numeric
	// observation values (obArray), or any combination of these.
	ObString string `json:"obString"`
	// A single observation value expressed in the specified unit of measure (obUoM).
	// Because of the variability of the Space Environment data types, each record may
	// employ a numeric observation value (obValue), a string observation value
	// (obString), a Boolean observation value (obBool), an array of numeric
	// observation values (obArray), or any combination of these.
	ObValue float64 `json:"obValue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObType        respjson.Field
		ObUoM         respjson.Field
		ObArray       respjson.Field
		ObBool        respjson.Field
		ObDescription respjson.Field
		ObQuality     respjson.Field
		ObString      respjson.Field
		ObValue       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpaceEnvObservationListResponseSeoList) RawJSON() string { return r.JSON.raw }
func (r *SpaceEnvObservationListResponseSeoList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpaceEnvObservationQueryhelpResponse struct {
	AodrSupported         bool                                            `json:"aodrSupported"`
	ClassificationMarking string                                          `json:"classificationMarking"`
	Description           string                                          `json:"description"`
	HistorySupported      bool                                            `json:"historySupported"`
	Name                  string                                          `json:"name"`
	Parameters            []SpaceEnvObservationQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                        `json:"requiredRoles"`
	RestSupported         bool                                            `json:"restSupported"`
	SortSupported         bool                                            `json:"sortSupported"`
	TypeName              string                                          `json:"typeName"`
	Uri                   string                                          `json:"uri"`
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
func (r SpaceEnvObservationQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SpaceEnvObservationQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpaceEnvObservationQueryhelpResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpaceEnvObservationQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *SpaceEnvObservationQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpaceEnvObservationListParams struct {
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpaceEnvObservationListParams]'s query parameters as
// `url.Values`.
func (r SpaceEnvObservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SpaceEnvObservationCountParams struct {
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpaceEnvObservationCountParams]'s query parameters as
// `url.Values`.
func (r SpaceEnvObservationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SpaceEnvObservationNewBulkParams struct {
	Body []SpaceEnvObservationNewBulkParamsBody
	paramObj
}

func (r SpaceEnvObservationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *SpaceEnvObservationNewBulkParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// SpaceEnvObservation data.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type SpaceEnvObservationNewBulkParamsBody struct {
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
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Spacecraft/sensor altitude at observation time, expressed in kilometers above
	// WGS-84 ellipsoid.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The data type (e.g. AP, AURORAL FLUX, ECP, KINDEX, PROPAGATED SOLAR WIND, XRAY
	// FLUX, etc.) of observations in this record.
	DataType param.Opt[string] `json:"dataType,omitzero"`
	// Flag indicating that this record contains derived data.
	Derived param.Opt[bool] `json:"derived,omitzero"`
	// Descriptive or additional information associated with this observation record.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Flag indicating that this record contains forecast data.
	Forecast param.Opt[bool] `json:"forecast,omitzero"`
	// The external system which generated the message, if applicable.
	GenSystem param.Opt[string] `json:"genSystem,omitzero"`
	// The time at which the associated data message was generated, in ISO 8601 UTC
	// format with millisecond precision.
	GenTime param.Opt[time.Time] `json:"genTime,omitzero" format:"date-time"`
	// Unique identifier of the on-orbit satellite hosting the sensor which produced
	// this data.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The type of instrument from which this data was collected (e.g. ANTENNA,
	// CHANNELTRON, INTERFEROMETER, MAGNETOMETER, RADIOMETER, etc.).
	InstrumentType param.Opt[string] `json:"instrumentType,omitzero"`
	// WGS-84 spacecraft/sensor latitude sub-point at observation time, represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 spacecraft/sensor longitude sub-point at observation time, represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The sensor measurement type of the observation data contained in this record.
	MeasType param.Opt[string] `json:"measType,omitzero"`
	// The type of message associated with this record.
	MsgType param.Opt[string] `json:"msgType,omitzero"`
	// The name of the observatory from which this data was collected.
	ObservatoryName param.Opt[string] `json:"observatoryName,omitzero"`
	// Additional notes concerning the observatory.
	ObservatoryNotes param.Opt[string] `json:"observatoryNotes,omitzero"`
	// The type of observatory from which this data was collected (e.g. FACILITY,
	// ONORBIT, NETWORK, etc.).
	ObservatoryType param.Opt[string] `json:"observatoryType,omitzero"`
	// A user-defined name or ID of a set of observations, if applicable. Used for
	// identifying multiple observation records as part of one observation set.
	ObSetID param.Opt[string] `json:"obSetId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor which produced this data. This may be an internal identifier
	// and not necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// The particle type (AEROSOL, ALPHA PARTICLE, ATOM, DUST, ELECTRON, ION, MOLECULE,
	// NEUTRON, POSITRON, PROTON) associated with this measurement.
	ParticleType param.Opt[string] `json:"particleType,omitzero"`
	// The quality of the overall data contained in this record. The quality indicator
	// value may vary among providers and may be a generalized statement (BAD, GOOD,
	// UNCERTAIN, UNKNOWN) or a numeric value. Users should consult the data provider
	// to verify the usage of the quality indicator.
	Quality param.Opt[string] `json:"quality,omitzero"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor which
	// produced this data.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The energy level bin of the sensor associated with this measurement.
	SenEnergyLevel param.Opt[string] `json:"senEnergyLevel,omitzero"`
	// Three element array, expressing the observing spacecraft/sensor position vector
	// components at observation time, in kilometers, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xpos, ypos, zpos].
	SenPos []float64 `json:"senPos,omitzero"`
	// The reference frame of the observing spacecraft/sensor state. If the
	// senReferenceFrame is null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame string `json:"senReferenceFrame,omitzero"`
	// Three element array, expressing the observing spacecraft/sensor velocity vector
	// components at observation time, in kilometers/second, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xvel, yvel, zvel].
	SenVel []float64 `json:"senVel,omitzero"`
	// A collection of individual space environment observations.
	SeoList []SpaceEnvObservationNewBulkParamsBodySeoList `json:"seoList,omitzero"`
	// Array of UUIDs of the UDL data records that are related to this observation
	// record. See the associated 'srcTyps' array for specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// of the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SGI, SIGACT, SKYIMAGE, SPACEENVOB,
	// SV, TRACK) that are related to this observation record. See the associated
	// 'srcIds' array for the record UUIDs, positionally corresponding to the record
	// types in this array. The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	paramObj
}

func (r SpaceEnvObservationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SpaceEnvObservationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpaceEnvObservationNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SpaceEnvObservationNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SpaceEnvObservationNewBulkParamsBody](
		"senReferenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// A single space environment observation.
//
// The properties ObType, ObUoM are required.
type SpaceEnvObservationNewBulkParamsBodySeoList struct {
	// The type of observation associated with this record.
	ObType string `json:"obType,required"`
	// The Unit of Measure associated with this observation. If there are no physical
	// units associated with the measurement, a value of NONE should be specified.
	ObUoM string `json:"obUoM,required"`
	// A Boolean observation. Because of the variability of the Space Environment data
	// types, each record may employ a numeric observation value (obValue), a string
	// observation value (obString), a Boolean observation value (obBool), an array of
	// numeric observation values (obArray), or any combination of these.
	ObBool param.Opt[bool] `json:"obBool,omitzero"`
	// Descriptive or additional information associated with this individual
	// observation.
	ObDescription param.Opt[string] `json:"obDescription,omitzero"`
	// The quality of this individual observation. The observation quality indicator
	// value may vary among providers and may be a generalized statement (BAD, GOOD,
	// UNCERTAIN, UNKNOWN) or a numeric value. Users should consult the data provider
	// to verify the usage of the observation.
	ObQuality param.Opt[string] `json:"obQuality,omitzero"`
	// A single observation string expressed in the specified unit of measure (obUoM).
	// Because of the variability of the Space Environment data types, each record may
	// employ a numeric observation value (obValue), a string observation value
	// (obString), a Boolean observation value (obBool), an array of numeric
	// observation values (obArray), or any combination of these.
	ObString param.Opt[string] `json:"obString,omitzero"`
	// A single observation value expressed in the specified unit of measure (obUoM).
	// Because of the variability of the Space Environment data types, each record may
	// employ a numeric observation value (obValue), a string observation value
	// (obString), a Boolean observation value (obBool), an array of numeric
	// observation values (obArray), or any combination of these.
	ObValue param.Opt[float64] `json:"obValue,omitzero"`
	// An array of observation values expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Space Environment data types, each
	// record may employ a numeric observation value (obValue), a string observation
	// value (obString), a Boolean observation value (obBool), an array of numeric
	// observation values (obArray), or any combination of these.
	ObArray []float64 `json:"obArray,omitzero"`
	paramObj
}

func (r SpaceEnvObservationNewBulkParamsBodySeoList) MarshalJSON() (data []byte, err error) {
	type shadow SpaceEnvObservationNewBulkParamsBodySeoList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpaceEnvObservationNewBulkParamsBodySeoList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpaceEnvObservationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpaceEnvObservationTupleParams]'s query parameters as
// `url.Values`.
func (r SpaceEnvObservationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SpaceEnvObservationUnvalidatedPublishParams struct {
	Body []SpaceEnvObservationUnvalidatedPublishParamsBody
	paramObj
}

func (r SpaceEnvObservationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *SpaceEnvObservationUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// SpaceEnvObservation data.
//
// The properties ClassificationMarking, DataMode, ObTime, Source are required.
type SpaceEnvObservationUnvalidatedPublishParamsBody struct {
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
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Spacecraft/sensor altitude at observation time, expressed in kilometers above
	// WGS-84 ellipsoid.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The data type (e.g. AP, AURORAL FLUX, ECP, KINDEX, PROPAGATED SOLAR WIND, XRAY
	// FLUX, etc.) of observations in this record.
	DataType param.Opt[string] `json:"dataType,omitzero"`
	// Flag indicating that this record contains derived data.
	Derived param.Opt[bool] `json:"derived,omitzero"`
	// Descriptive or additional information associated with this observation record.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Flag indicating that this record contains forecast data.
	Forecast param.Opt[bool] `json:"forecast,omitzero"`
	// The external system which generated the message, if applicable.
	GenSystem param.Opt[string] `json:"genSystem,omitzero"`
	// The time at which the associated data message was generated, in ISO 8601 UTC
	// format with millisecond precision.
	GenTime param.Opt[time.Time] `json:"genTime,omitzero" format:"date-time"`
	// Unique identifier of the on-orbit satellite hosting the sensor which produced
	// this data.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The type of instrument from which this data was collected (e.g. ANTENNA,
	// CHANNELTRON, INTERFEROMETER, MAGNETOMETER, RADIOMETER, etc.).
	InstrumentType param.Opt[string] `json:"instrumentType,omitzero"`
	// WGS-84 spacecraft/sensor latitude sub-point at observation time, represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 spacecraft/sensor longitude sub-point at observation time, represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The sensor measurement type of the observation data contained in this record.
	MeasType param.Opt[string] `json:"measType,omitzero"`
	// The type of message associated with this record.
	MsgType param.Opt[string] `json:"msgType,omitzero"`
	// The name of the observatory from which this data was collected.
	ObservatoryName param.Opt[string] `json:"observatoryName,omitzero"`
	// Additional notes concerning the observatory.
	ObservatoryNotes param.Opt[string] `json:"observatoryNotes,omitzero"`
	// The type of observatory from which this data was collected (e.g. FACILITY,
	// ONORBIT, NETWORK, etc.).
	ObservatoryType param.Opt[string] `json:"observatoryType,omitzero"`
	// A user-defined name or ID of a set of observations, if applicable. Used for
	// identifying multiple observation records as part of one observation set.
	ObSetID param.Opt[string] `json:"obSetId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor which produced this data. This may be an internal identifier
	// and not necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// The particle type (AEROSOL, ALPHA PARTICLE, ATOM, DUST, ELECTRON, ION, MOLECULE,
	// NEUTRON, POSITRON, PROTON) associated with this measurement.
	ParticleType param.Opt[string] `json:"particleType,omitzero"`
	// The quality of the overall data contained in this record. The quality indicator
	// value may vary among providers and may be a generalized statement (BAD, GOOD,
	// UNCERTAIN, UNKNOWN) or a numeric value. Users should consult the data provider
	// to verify the usage of the quality indicator.
	Quality param.Opt[string] `json:"quality,omitzero"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor which
	// produced this data.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The energy level bin of the sensor associated with this measurement.
	SenEnergyLevel param.Opt[string] `json:"senEnergyLevel,omitzero"`
	// Three element array, expressing the observing spacecraft/sensor position vector
	// components at observation time, in kilometers, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xpos, ypos, zpos].
	SenPos []float64 `json:"senPos,omitzero"`
	// The reference frame of the observing spacecraft/sensor state. If the
	// senReferenceFrame is null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame string `json:"senReferenceFrame,omitzero"`
	// Three element array, expressing the observing spacecraft/sensor velocity vector
	// components at observation time, in kilometers/second, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xvel, yvel, zvel].
	SenVel []float64 `json:"senVel,omitzero"`
	// A collection of individual space environment observations.
	SeoList []SpaceEnvObservationUnvalidatedPublishParamsBodySeoList `json:"seoList,omitzero"`
	// Array of UUIDs of the UDL data records that are related to this observation
	// record. See the associated 'srcTyps' array for specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// of the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SGI, SIGACT, SKYIMAGE, SPACEENVOB,
	// SV, TRACK) that are related to this observation record. See the associated
	// 'srcIds' array for the record UUIDs, positionally corresponding to the record
	// types in this array. The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	paramObj
}

func (r SpaceEnvObservationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SpaceEnvObservationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpaceEnvObservationUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SpaceEnvObservationUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SpaceEnvObservationUnvalidatedPublishParamsBody](
		"senReferenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// A single space environment observation.
//
// The properties ObType, ObUoM are required.
type SpaceEnvObservationUnvalidatedPublishParamsBodySeoList struct {
	// The type of observation associated with this record.
	ObType string `json:"obType,required"`
	// The Unit of Measure associated with this observation. If there are no physical
	// units associated with the measurement, a value of NONE should be specified.
	ObUoM string `json:"obUoM,required"`
	// A Boolean observation. Because of the variability of the Space Environment data
	// types, each record may employ a numeric observation value (obValue), a string
	// observation value (obString), a Boolean observation value (obBool), an array of
	// numeric observation values (obArray), or any combination of these.
	ObBool param.Opt[bool] `json:"obBool,omitzero"`
	// Descriptive or additional information associated with this individual
	// observation.
	ObDescription param.Opt[string] `json:"obDescription,omitzero"`
	// The quality of this individual observation. The observation quality indicator
	// value may vary among providers and may be a generalized statement (BAD, GOOD,
	// UNCERTAIN, UNKNOWN) or a numeric value. Users should consult the data provider
	// to verify the usage of the observation.
	ObQuality param.Opt[string] `json:"obQuality,omitzero"`
	// A single observation string expressed in the specified unit of measure (obUoM).
	// Because of the variability of the Space Environment data types, each record may
	// employ a numeric observation value (obValue), a string observation value
	// (obString), a Boolean observation value (obBool), an array of numeric
	// observation values (obArray), or any combination of these.
	ObString param.Opt[string] `json:"obString,omitzero"`
	// A single observation value expressed in the specified unit of measure (obUoM).
	// Because of the variability of the Space Environment data types, each record may
	// employ a numeric observation value (obValue), a string observation value
	// (obString), a Boolean observation value (obBool), an array of numeric
	// observation values (obArray), or any combination of these.
	ObValue param.Opt[float64] `json:"obValue,omitzero"`
	// An array of observation values expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Space Environment data types, each
	// record may employ a numeric observation value (obValue), a string observation
	// value (obString), a Boolean observation value (obBool), an array of numeric
	// observation values (obArray), or any combination of these.
	ObArray []float64 `json:"obArray,omitzero"`
	paramObj
}

func (r SpaceEnvObservationUnvalidatedPublishParamsBodySeoList) MarshalJSON() (data []byte, err error) {
	type shadow SpaceEnvObservationUnvalidatedPublishParamsBodySeoList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpaceEnvObservationUnvalidatedPublishParamsBodySeoList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
