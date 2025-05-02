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

// SpaceenvobservationHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpaceenvobservationHistoryService] method instead.
type SpaceenvobservationHistoryService struct {
	Options []option.RequestOption
}

// NewSpaceenvobservationHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSpaceenvobservationHistoryService(opts ...option.RequestOption) (r SpaceenvobservationHistoryService) {
	r = SpaceenvobservationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SpaceenvobservationHistoryService) List(ctx context.Context, query SpaceenvobservationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SpaceEnvObservationFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/spaceenvobservation/history"
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
func (r *SpaceenvobservationHistoryService) ListAutoPaging(ctx context.Context, query SpaceenvobservationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SpaceEnvObservationFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SpaceenvobservationHistoryService) Aodr(ctx context.Context, query SpaceenvobservationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/spaceenvobservation/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SpaceenvobservationHistoryService) Count(ctx context.Context, query SpaceenvobservationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/spaceenvobservation/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// SpaceEnvObservation data.
type SpaceEnvObservationFull struct {
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
	DataMode SpaceEnvObservationFullDataMode `json:"dataMode,required"`
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
	SenReferenceFrame SpaceEnvObservationFullSenReferenceFrame `json:"senReferenceFrame"`
	// Three element array, expressing the observing spacecraft/sensor velocity vector
	// components at observation time, in kilometers/second, in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	// The array element order is [xvel, yvel, zvel].
	SenVel []float64 `json:"senVel"`
	// A collection of individual space environment observations.
	SeoList []SpaceEnvObservationFullSeoList `json:"seoList"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		ObTime                resp.Field
		Source                resp.Field
		ID                    resp.Field
		Alt                   resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DataType              resp.Field
		Derived               resp.Field
		Description           resp.Field
		Descriptor            resp.Field
		ExternalID            resp.Field
		Forecast              resp.Field
		GenSystem             resp.Field
		GenTime               resp.Field
		IDOnOrbit             resp.Field
		IDSensor              resp.Field
		InstrumentType        resp.Field
		Lat                   resp.Field
		Lon                   resp.Field
		MeasType              resp.Field
		MsgType               resp.Field
		ObservatoryName       resp.Field
		ObservatoryNotes      resp.Field
		ObservatoryType       resp.Field
		ObSetID               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		OrigSensorID          resp.Field
		ParticleType          resp.Field
		Quality               resp.Field
		SatNo                 resp.Field
		SenEnergyLevel        resp.Field
		SenPos                resp.Field
		SenReferenceFrame     resp.Field
		SenVel                resp.Field
		SeoList               resp.Field
		SrcIDs                resp.Field
		SrcTyps               resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpaceEnvObservationFull) RawJSON() string { return r.JSON.raw }
func (r *SpaceEnvObservationFull) UnmarshalJSON(data []byte) error {
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
type SpaceEnvObservationFullDataMode string

const (
	SpaceEnvObservationFullDataModeReal      SpaceEnvObservationFullDataMode = "REAL"
	SpaceEnvObservationFullDataModeTest      SpaceEnvObservationFullDataMode = "TEST"
	SpaceEnvObservationFullDataModeSimulated SpaceEnvObservationFullDataMode = "SIMULATED"
	SpaceEnvObservationFullDataModeExercise  SpaceEnvObservationFullDataMode = "EXERCISE"
)

// The reference frame of the observing spacecraft/sensor state. If the
// senReferenceFrame is null it is assumed to be J2000.
type SpaceEnvObservationFullSenReferenceFrame string

const (
	SpaceEnvObservationFullSenReferenceFrameJ2000   SpaceEnvObservationFullSenReferenceFrame = "J2000"
	SpaceEnvObservationFullSenReferenceFrameEfgTdr  SpaceEnvObservationFullSenReferenceFrame = "EFG/TDR"
	SpaceEnvObservationFullSenReferenceFrameEcrEcef SpaceEnvObservationFullSenReferenceFrame = "ECR/ECEF"
	SpaceEnvObservationFullSenReferenceFrameTeme    SpaceEnvObservationFullSenReferenceFrame = "TEME"
	SpaceEnvObservationFullSenReferenceFrameItrf    SpaceEnvObservationFullSenReferenceFrame = "ITRF"
	SpaceEnvObservationFullSenReferenceFrameGcrf    SpaceEnvObservationFullSenReferenceFrame = "GCRF"
)

// A single space environment observation.
type SpaceEnvObservationFullSeoList struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ObType        resp.Field
		ObUoM         resp.Field
		ObArray       resp.Field
		ObBool        resp.Field
		ObDescription resp.Field
		ObQuality     resp.Field
		ObString      resp.Field
		ObValue       resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpaceEnvObservationFullSeoList) RawJSON() string { return r.JSON.raw }
func (r *SpaceEnvObservationFullSeoList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpaceenvobservationHistoryListParams struct {
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
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
func (f SpaceenvobservationHistoryListParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [SpaceenvobservationHistoryListParams]'s query parameters as
// `url.Values`.
func (r SpaceenvobservationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SpaceenvobservationHistoryAodrParams struct {
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
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
func (f SpaceenvobservationHistoryAodrParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [SpaceenvobservationHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r SpaceenvobservationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SpaceenvobservationHistoryCountParams struct {
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SpaceenvobservationHistoryCountParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [SpaceenvobservationHistoryCountParams]'s query parameters
// as `url.Values`.
func (r SpaceenvobservationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
