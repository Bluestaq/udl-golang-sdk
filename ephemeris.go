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

// EphemerisService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEphemerisService] method instead.
type EphemerisService struct {
	Options      []option.RequestOption
	AttitudeData EphemerisAttitudeDataService
	History      EphemerisHistoryService
}

// NewEphemerisService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEphemerisService(opts ...option.RequestOption) (r EphemerisService) {
	r = EphemerisService{}
	r.Options = opts
	r.AttitudeData = NewEphemerisAttitudeDataService(opts...)
	r.History = NewEphemerisHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EphemerisService) List(ctx context.Context, query EphemerisListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EphemerisAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/ephemeris"
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
func (r *EphemerisService) ListAutoPaging(ctx context.Context, query EphemerisListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EphemerisAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EphemerisService) Count(ctx context.Context, query EphemerisCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/ephemeris/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to post/store Ephemeris data. This operation is intended to be
// used for automated feeds into UDL. The payload is in Ephemeris format as
// described by the "Flight Safety Handbook" published by 18th Space Command. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
//
// **Example:**
// /filedrop/ephem?classification=U&dataMode=TEST&source=Bluestaq&satNo=25544&ephemFormatType=NASA&hasMnvr=false&type=ROUTINE&category=EXTERNAL&origin=NASA&tags=tag1,tag2
func (r *EphemerisService) FileUpload(ctx context.Context, params EphemerisFileUploadParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/ephem"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EphemerisService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *EphemerisQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/ephemeris/queryhelp"
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
func (r *EphemerisService) Tuple(ctx context.Context, query EphemerisTupleParams, opts ...option.RequestOption) (res *[]shared.EphemerisFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/ephemeris/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a single EphemerisSet and many associated Ephemeris
// records as a POST body for ingest into the database. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
//
// The following rules apply to this operation:
//
// <h3>
//   - Ephemeris Set numPoints value must correspond exactly to the number of Ephemeris states associated with that Ephemeris Set. The numPoints value is checked against the actual posted number of states, and a mismatch will result in the post being rejected.
//   - Ephemeris Set pointStartTime and pointEndTime must correspond to the earliest and latest state times, respectively, of the associated Ephemeris states.
//   - Either satNo, idOnOrbit, or origObjectId must be provided. The preferred option is to post with satNo for a cataloged object, and with (only) origObjectId for an unknown, uncataloged, or internal/test object. There are several cases for the logic associated with these fields:
//   - If satNo is provided and correlates to a known UDL sat number then the idOnOrbit will be populated appropriately in addition to the satNo.
//   - If satNo is provided and does not correlate to a known UDL sat number then the provided satNo value will be moved to the origObjectId field and satNo left null.
//   - If origObjectId and a valid satNo or idOnOrbit are provided then both the satNo/idOnOrbit and origObjectId will maintain the provided values.
//   - If only origObjectId is provided then origObjectId will be populated with the posted value. In this case, no checks are made against existing UDL sat numbers.
//
// </h3>
func (r *EphemerisService) UnvalidatedPublish(ctx context.Context, body EphemerisUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-ephset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// An ephemeris record is a position and velocity vector identifying the location
// and trajectory of an on-orbit object at a specified time. Ephemeris points,
// including covariance, are in kilometer and second based units in a user
// specified reference frame, with ECI J2K being preferred. The EphemerisSet ID
// (esId) links all points associated with an ephemeris set. The 'EphemerisSet'
// record contains details of the underlying data and propagation models used in
// the generation of the ephemeris. Ephemeris points must be retrieved by
// specifying the parent EphemerisSet ID (esId).
type EphemerisAbridged struct {
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
	DataMode EphemerisAbridgedDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Time associated with the Ephemeris Point, in ISO8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Cartesian X position of target, in km, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos,required"`
	// Cartesian X velocity of target, in km/sec, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel,required"`
	// Cartesian Y position of target, in km, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos,required"`
	// Cartesian Y velocity of target, in km/sec, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel,required"`
	// Cartesian Z position of target, in km, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos,required"`
	// Cartesian Z velocity of target, in km/sec, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// Ephemeris Set covReferenceFrame. If the covReferenceFrame from the EphemerisSet
	// table is null it is assumed to be J2000. The array values represent the lower
	// triangular half of the position-velocity covariance matrix. The size of the
	// covariance matrix is dynamic, depending on whether the covariance for position
	// only or position & velocity. The covariance elements are position dependent
	// within the array with values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;y&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;y'&nbsp;z'&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;
	// 21
	//
	// The array containing the covariance matrix elements will be of length 6 for
	// position only covariance, or length 21 for position-velocity covariance. The cov
	// array should contain only the lower left triangle values from top left down to
	// bottom right, in order.
	Cov []float64 `json:"cov"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent EphemerisSet, auto-generated by the system. The
	// esId (ephemerisSet id) is used to identify all individual ephemeris states
	// associated with a parent ephemerisSet.
	EsID string `json:"esId"`
	// Unique identifier of the on-orbit satellite object.
	IDOnOrbit string `json:"idOnOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// Optional identifier provided by ephemeris source to indicate the target object
	// of this ephemeris. This may be an internal identifier and not necessarily map to
	// a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Cartesian X acceleration of target, in km/sec^2, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel float64 `json:"xaccel"`
	// Cartesian Y acceleration of target, in km/sec^2, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel float64 `json:"yaccel"`
	// Cartesian Z acceleration of target, in km/sec^2, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel float64 `json:"zaccel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		Xpos                  respjson.Field
		Xvel                  respjson.Field
		Ypos                  respjson.Field
		Yvel                  respjson.Field
		Zpos                  respjson.Field
		Zvel                  respjson.Field
		ID                    respjson.Field
		Cov                   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EsID                  respjson.Field
		IDOnOrbit             respjson.Field
		Origin                respjson.Field
		OrigObjectID          respjson.Field
		Xaccel                respjson.Field
		Yaccel                respjson.Field
		Zaccel                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EphemerisAbridged) RawJSON() string { return r.JSON.raw }
func (r *EphemerisAbridged) UnmarshalJSON(data []byte) error {
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
type EphemerisAbridgedDataMode string

const (
	EphemerisAbridgedDataModeReal      EphemerisAbridgedDataMode = "REAL"
	EphemerisAbridgedDataModeTest      EphemerisAbridgedDataMode = "TEST"
	EphemerisAbridgedDataModeSimulated EphemerisAbridgedDataMode = "SIMULATED"
	EphemerisAbridgedDataModeExercise  EphemerisAbridgedDataMode = "EXERCISE"
)

type EphemerisQueryhelpResponse struct {
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
func (r EphemerisQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *EphemerisQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EphemerisListParams struct {
	// Unique identifier of the parent EphemerisSet, auto-generated by the system. The
	// esId (ephemerisSet id) is used to identify all individual ephemeris states
	// associated with a parent ephemerisSet. (uuid)
	EsID        string           `query:"esId,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EphemerisListParams]'s query parameters as `url.Values`.
func (r EphemerisListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EphemerisCountParams struct {
	// Unique identifier of the parent EphemerisSet, auto-generated by the system. The
	// esId (ephemerisSet id) is used to identify all individual ephemeris states
	// associated with a parent ephemerisSet. (uuid)
	EsID        string           `query:"esId,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EphemerisCountParams]'s query parameters as `url.Values`.
func (r EphemerisCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EphemerisFileUploadParams struct {
	// Ephemeris category.
	Category string `query:"category,required" json:"-"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	Classification string `query:"classification,required" json:"-"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode EphemerisFileUploadParamsDataMode `query:"dataMode,omitzero,required" json:"-"`
	// Ephemeris format as documented in Flight Safety Handbook.
	//
	// Any of "ModITC", "GOO", "NASA", "OEM", "OASYS".
	EphemFormatType EphemerisFileUploadParamsEphemFormatType `query:"ephemFormatType,omitzero,required" json:"-"`
	// Boolean indicating whether maneuver(s) are incorporated into the ephemeris.
	HasMnvr bool `query:"hasMnvr,required" json:"-"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `query:"satNo,required" json:"-"`
	// Source of the Ephemeris data.
	Source string `query:"source,required" json:"-"`
	// Ephemeris type.
	Type string `query:"type,required" json:"-"`
	Body string
	// Optional origin of the Ephemeris.
	Origin param.Opt[string] `query:"origin,omitzero" json:"-"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	paramObj
}

func (r EphemerisFileUploadParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *EphemerisFileUploadParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [EphemerisFileUploadParams]'s query parameters as
// `url.Values`.
func (r EphemerisFileUploadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type EphemerisFileUploadParamsDataMode string

const (
	EphemerisFileUploadParamsDataModeReal      EphemerisFileUploadParamsDataMode = "REAL"
	EphemerisFileUploadParamsDataModeTest      EphemerisFileUploadParamsDataMode = "TEST"
	EphemerisFileUploadParamsDataModeSimulated EphemerisFileUploadParamsDataMode = "SIMULATED"
	EphemerisFileUploadParamsDataModeExercise  EphemerisFileUploadParamsDataMode = "EXERCISE"
)

// Ephemeris format as documented in Flight Safety Handbook.
type EphemerisFileUploadParamsEphemFormatType string

const (
	EphemerisFileUploadParamsEphemFormatTypeModItc EphemerisFileUploadParamsEphemFormatType = "ModITC"
	EphemerisFileUploadParamsEphemFormatTypeGoo    EphemerisFileUploadParamsEphemFormatType = "GOO"
	EphemerisFileUploadParamsEphemFormatTypeNasa   EphemerisFileUploadParamsEphemFormatType = "NASA"
	EphemerisFileUploadParamsEphemFormatTypeOem    EphemerisFileUploadParamsEphemFormatType = "OEM"
	EphemerisFileUploadParamsEphemFormatTypeOasys  EphemerisFileUploadParamsEphemFormatType = "OASYS"
)

type EphemerisTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Unique identifier of the parent EphemerisSet, auto-generated by the system. The
	// esId (ephemerisSet id) is used to identify all individual ephemeris states
	// associated with a parent ephemerisSet. (uuid)
	EsID        string           `query:"esId,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EphemerisTupleParams]'s query parameters as `url.Values`.
func (r EphemerisTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EphemerisUnvalidatedPublishParams struct {
	// The source category of the ephemeris (e.g. OWNER_OPERATOR, ANALYST, EXTERNAL).
	Category string `json:"category,required"`
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
	DataMode EphemerisUnvalidatedPublishParamsDataMode `json:"dataMode,omitzero,required"`
	// Number of points contained in the ephemeris.
	NumPoints int64 `json:"numPoints,required"`
	// End time/last time point of the ephemeris, in ISO 8601 UTC format.
	PointEndTime time.Time `json:"pointEndTime,required" format:"date-time"`
	// Start time/first time point of the ephemeris, in ISO 8601 UTC format.
	PointStartTime time.Time `json:"pointStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type/purpose of the ephemeris (e.g., CALIBRATION, LAUNCH, MNVR_PLAN,
	// ROUTINE, SCREENING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// First derivative of ballistic coefficient (m^2/kg-s).
	BDot param.Opt[float64] `json:"bDot,omitzero"`
	// The Central Body of the ephemeris. Assumed to be Earth, unless otherwise
	// indicated.
	CentBody param.Opt[string] `json:"centBody,omitzero"`
	// Additional source provided comments associated with the ephemeris.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Notes/description of the provided ephemeris. A value of DSTOP signifies the
	// ephemeris were generated using the last observation available.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Drag model used in ephemeris generation (e.g. JAC70, MSIS90, NONE, etc.).
	DragModel param.Opt[string] `json:"dragModel,omitzero"`
	// Model parameter value for energy dissipation rate (EDR), expressed in w/kg.
	Edr param.Opt[float64] `json:"edr,omitzero"`
	// Filename of the raw file used to provide the ephemeris data including filetype
	// extension, if applicable. This file may be retrieved using the 'getFile'
	// operation as specified in the 'EphemerisSet' OpenAPI docs.
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Geopotential model used in ephemeris generation (e.g. EGM-96, WGS-84, WGS-72,
	// JGM-2, GEM-T3), including mm degree zonals, nn degree/order tesserals (e.g.
	// EGM-96 24Z,24T).
	GeopotentialModel param.Opt[string] `json:"geopotentialModel,omitzero"`
	// Boolean indicating whether acceleration data is provided with the ephemeris.
	HasAccel param.Opt[bool] `json:"hasAccel,omitzero"`
	// Boolean indicating whether covariance data is provided with the ephemeris.
	HasCov param.Opt[bool] `json:"hasCov,omitzero"`
	// Boolean indicating whether maneuver(s) are incorporated into the ephemeris.
	HasMnvr param.Opt[bool] `json:"hasMnvr,omitzero"`
	// Unique identifier of the primary satellite on-orbit object.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// ID of the State Vector used to generate the ephemeris.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Integrator used in ephemeris generation (e.g. RK7(8), RK8(9), COWELL, TWO-BODY).
	Integrator param.Opt[string] `json:"integrator,omitzero"`
	// The recommended interpolation method for the ephemeris data.
	Interpolation param.Opt[string] `json:"interpolation,omitzero"`
	// The recommended interpolation degree for the ephemeris data.
	InterpolationDegree param.Opt[int64] `json:"interpolationDegree,omitzero"`
	// Boolean indicating use of lunar/solar data in ephemeris generation.
	LunarSolar param.Opt[bool] `json:"lunarSolar,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by ephemeris source to indicate the target object
	// of this ephemeris. This may be an internal identifier and not necessarily map to
	// a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// The pedigree of the ephemeris or source data used for ephemeris generation (e.g.
	// DOPPLER, GPS, HYBRID, PROPAGATED, RANGING, SLR).
	Pedigree param.Opt[string] `json:"pedigree,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Boolean indicating use of solid earth tide data in ephemeris generation.
	SolidEarthTides param.Opt[bool] `json:"solidEarthTides,omitzero"`
	// Ephemeris step size, in seconds.
	StepSize param.Opt[int64] `json:"stepSize,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Optional end time of the usable time span for the ephemeris data, in ISO 8601
	// UTC format with microsecond precision.
	UsableEndTime param.Opt[time.Time] `json:"usableEndTime,omitzero" format:"date-time"`
	// Optional start time of the usable time span for the ephemeris data, in ISO 8601
	// UTC format with microsecond precision.
	UsableStartTime param.Opt[time.Time] `json:"usableStartTime,omitzero" format:"date-time"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW", "EFG/TDR", "ECR/ECEF", "TEME", "GCRF".
	CovReferenceFrame EphemerisUnvalidatedPublishParamsCovReferenceFrame `json:"covReferenceFrame,omitzero"`
	// The list of ephemeris states belonging to the EphemerisSet. Each ephemeris point
	// is associated with a parent Ephemeris Set via the EphemerisSet ID (esId).
	EphemerisList []EphemerisUnvalidatedPublishParamsEphemerisList `json:"ephemerisList,omitzero"`
	// Array of the maneuver IDs of all maneuvers incorporated in the ephemeris.
	IDManeuvers []string `json:"idManeuvers,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame EphemerisUnvalidatedPublishParamsReferenceFrame `json:"referenceFrame,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r EphemerisUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	type shadow EphemerisUnvalidatedPublishParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EphemerisUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
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
type EphemerisUnvalidatedPublishParamsDataMode string

const (
	EphemerisUnvalidatedPublishParamsDataModeReal      EphemerisUnvalidatedPublishParamsDataMode = "REAL"
	EphemerisUnvalidatedPublishParamsDataModeTest      EphemerisUnvalidatedPublishParamsDataMode = "TEST"
	EphemerisUnvalidatedPublishParamsDataModeSimulated EphemerisUnvalidatedPublishParamsDataMode = "SIMULATED"
	EphemerisUnvalidatedPublishParamsDataModeExercise  EphemerisUnvalidatedPublishParamsDataMode = "EXERCISE"
)

// The reference frame of the covariance matrix elements. If the covReferenceFrame
// is null it is assumed to be J2000.
type EphemerisUnvalidatedPublishParamsCovReferenceFrame string

const (
	EphemerisUnvalidatedPublishParamsCovReferenceFrameJ2000   EphemerisUnvalidatedPublishParamsCovReferenceFrame = "J2000"
	EphemerisUnvalidatedPublishParamsCovReferenceFrameUvw     EphemerisUnvalidatedPublishParamsCovReferenceFrame = "UVW"
	EphemerisUnvalidatedPublishParamsCovReferenceFrameEfgTdr  EphemerisUnvalidatedPublishParamsCovReferenceFrame = "EFG/TDR"
	EphemerisUnvalidatedPublishParamsCovReferenceFrameEcrEcef EphemerisUnvalidatedPublishParamsCovReferenceFrame = "ECR/ECEF"
	EphemerisUnvalidatedPublishParamsCovReferenceFrameTeme    EphemerisUnvalidatedPublishParamsCovReferenceFrame = "TEME"
	EphemerisUnvalidatedPublishParamsCovReferenceFrameGcrf    EphemerisUnvalidatedPublishParamsCovReferenceFrame = "GCRF"
)

// An ephemeris record is a position and velocity vector identifying the location
// and trajectory of an on-orbit object at a specified time. Ephemeris points,
// including covariance, are in kilometer and second based units in a user
// specified reference frame, with ECI J2K being preferred. The EphemerisSet ID
// (esId) links all points associated with an ephemeris set. The 'EphemerisSet'
// record contains details of the underlying data and propagation models used in
// the generation of the ephemeris. Ephemeris points must be retrieved by
// specifying the parent EphemerisSet ID (esId).
//
// The properties ClassificationMarking, DataMode, Source, Ts, Xpos, Xvel, Ypos,
// Yvel, Zpos, Zvel are required.
type EphemerisUnvalidatedPublishParamsEphemerisList struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Time associated with the Ephemeris Point, in ISO8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Cartesian X position of target, in km, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos,required"`
	// Cartesian X velocity of target, in km/sec, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel,required"`
	// Cartesian Y position of target, in km, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos,required"`
	// Cartesian Y velocity of target, in km/sec, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel,required"`
	// Cartesian Z position of target, in km, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos,required"`
	// Cartesian Z velocity of target, in km/sec, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Unique identifier of the parent EphemerisSet, auto-generated by the system. The
	// esId (ephemerisSet id) is used to identify all individual ephemeris states
	// associated with a parent ephemerisSet.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// Unique identifier of the on-orbit satellite object.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by ephemeris source to indicate the target object
	// of this ephemeris. This may be an internal identifier and not necessarily map to
	// a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Cartesian X acceleration of target, in km/sec^2, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel param.Opt[float64] `json:"xaccel,omitzero"`
	// Cartesian Y acceleration of target, in km/sec^2, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel param.Opt[float64] `json:"yaccel,omitzero"`
	// Cartesian Z acceleration of target, in km/sec^2, in the specified EphemerisSet
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel param.Opt[float64] `json:"zaccel,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// Ephemeris Set covReferenceFrame. If the covReferenceFrame from the EphemerisSet
	// table is null it is assumed to be J2000. The array values represent the lower
	// triangular half of the position-velocity covariance matrix. The size of the
	// covariance matrix is dynamic, depending on whether the covariance for position
	// only or position & velocity. The covariance elements are position dependent
	// within the array with values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;y&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;y'&nbsp;z'&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;
	// 21
	//
	// The array containing the covariance matrix elements will be of length 6 for
	// position only covariance, or length 21 for position-velocity covariance. The cov
	// array should contain only the lower left triangle values from top left down to
	// bottom right, in order.
	Cov []float64 `json:"cov,omitzero"`
	paramObj
}

func (r EphemerisUnvalidatedPublishParamsEphemerisList) MarshalJSON() (data []byte, err error) {
	type shadow EphemerisUnvalidatedPublishParamsEphemerisList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EphemerisUnvalidatedPublishParamsEphemerisList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EphemerisUnvalidatedPublishParamsEphemerisList](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// The reference frame of the cartesian orbital states. If the referenceFrame is
// null it is assumed to be J2000.
type EphemerisUnvalidatedPublishParamsReferenceFrame string

const (
	EphemerisUnvalidatedPublishParamsReferenceFrameJ2000   EphemerisUnvalidatedPublishParamsReferenceFrame = "J2000"
	EphemerisUnvalidatedPublishParamsReferenceFrameEfgTdr  EphemerisUnvalidatedPublishParamsReferenceFrame = "EFG/TDR"
	EphemerisUnvalidatedPublishParamsReferenceFrameEcrEcef EphemerisUnvalidatedPublishParamsReferenceFrame = "ECR/ECEF"
	EphemerisUnvalidatedPublishParamsReferenceFrameTeme    EphemerisUnvalidatedPublishParamsReferenceFrame = "TEME"
	EphemerisUnvalidatedPublishParamsReferenceFrameItrf    EphemerisUnvalidatedPublishParamsReferenceFrame = "ITRF"
	EphemerisUnvalidatedPublishParamsReferenceFrameGcrf    EphemerisUnvalidatedPublishParamsReferenceFrame = "GCRF"
)
