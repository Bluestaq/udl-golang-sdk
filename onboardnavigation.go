// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// OnboardnavigationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnboardnavigationService] method instead.
type OnboardnavigationService struct {
	Options []option.RequestOption
	History OnboardnavigationHistoryService
}

// NewOnboardnavigationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOnboardnavigationService(opts ...option.RequestOption) (r OnboardnavigationService) {
	r = OnboardnavigationService{}
	r.Options = opts
	r.History = NewOnboardnavigationHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnboardnavigationService) List(ctx context.Context, query OnboardnavigationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnboardnavigationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onboardnavigation"
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
func (r *OnboardnavigationService) ListAutoPaging(ctx context.Context, query OnboardnavigationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnboardnavigationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnboardnavigationService) Count(ctx context.Context, query OnboardnavigationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onboardnavigation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// OnboardNavigation records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *OnboardnavigationService) NewBulk(ctx context.Context, body OnboardnavigationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onboardnavigation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *OnboardnavigationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *OnboardnavigationQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/onboardnavigation/queryhelp"
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
func (r *OnboardnavigationService) Tuple(ctx context.Context, query OnboardnavigationTupleParams, opts ...option.RequestOption) (res *[]OnboardnavigationFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/onboardnavigation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a list of onboard navigation records as a POST body
// and ingest into the database. This operation is intended to be used for
// automated feeds into UDL. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *OnboardnavigationService) UnvalidatedPublish(ctx context.Context, body OnboardnavigationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-onboardnavigation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// These services provide spacecraft positional data derived from on-board
// navigational sensors. Navigational points are provided in kilometers and in a
// user specified reference frame, with ECI J2K being preferred.
type OnboardnavigationListResponse struct {
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
	DataMode OnboardnavigationListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Start time of the sensor data, in ISO 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Arrays of spacecraft delta position (X, Y, Z), in km, in the specified
	// referenceFrame, between the onboard state and the last accepted sensor position.
	DeltaPos [][]float64 `json:"deltaPos"`
	// End time of the sensor data, in ISO 8601 UTC format.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Unique identifier of the parent EphemerisSet, if this data is correlated with an
	// Ephemeris. If multiple nav sensor records are required, this ID may be
	// associated with each of those records if each is synced to the ephemeris points.
	EsID string `json:"esId"`
	// Unique identifier of the primary satellite on-orbit object.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the last onboard state vector.
	IDStateVector string `json:"idStateVector"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard magnetometer, at each epoch.
	Mag [][]float64 `json:"mag"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the data source to indicate the target object of
	// this record. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// The reference frame in which the sensor derived positions are provided. If the
	// referenceFrame is null it is assumed to be J2000. Note that this frame is
	// assumed to apply to all of the sensor data in this record. If onboard sensors
	// process positional estimates in different frames then separate records should be
	// generated.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame OnboardnavigationListResponseReferenceFrame `json:"referenceFrame"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// The last load time of the current star catalog onboard this spacecraft.
	StarCatLoadTime time.Time `json:"starCatLoadTime" format:"date-time"`
	// The name or identifier the star catalog in use.
	StarCatName string `json:"starCatName"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard star tracker, at each epoch.
	StarTracker [][]float64 `json:"starTracker"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard sun sensor, at each epoch.
	SunSensor [][]float64 `json:"sunSensor"`
	// Array of epochs of the observations, in ISO 8601 UTC format. The epochs are
	// assumed to correspond to all sensor data in this record. If sensors do not share
	// a common epoch then separate records should be generated.
	Ts []time.Time `json:"ts" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeltaPos              respjson.Field
		EndTime               respjson.Field
		EsID                  respjson.Field
		IDOnOrbit             respjson.Field
		IDStateVector         respjson.Field
		Mag                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		ReferenceFrame        respjson.Field
		SatNo                 respjson.Field
		StarCatLoadTime       respjson.Field
		StarCatName           respjson.Field
		StarTracker           respjson.Field
		SunSensor             respjson.Field
		Ts                    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnboardnavigationListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnboardnavigationListResponse) UnmarshalJSON(data []byte) error {
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
type OnboardnavigationListResponseDataMode string

const (
	OnboardnavigationListResponseDataModeReal      OnboardnavigationListResponseDataMode = "REAL"
	OnboardnavigationListResponseDataModeTest      OnboardnavigationListResponseDataMode = "TEST"
	OnboardnavigationListResponseDataModeSimulated OnboardnavigationListResponseDataMode = "SIMULATED"
	OnboardnavigationListResponseDataModeExercise  OnboardnavigationListResponseDataMode = "EXERCISE"
)

// The reference frame in which the sensor derived positions are provided. If the
// referenceFrame is null it is assumed to be J2000. Note that this frame is
// assumed to apply to all of the sensor data in this record. If onboard sensors
// process positional estimates in different frames then separate records should be
// generated.
type OnboardnavigationListResponseReferenceFrame string

const (
	OnboardnavigationListResponseReferenceFrameJ2000   OnboardnavigationListResponseReferenceFrame = "J2000"
	OnboardnavigationListResponseReferenceFrameEfgTdr  OnboardnavigationListResponseReferenceFrame = "EFG/TDR"
	OnboardnavigationListResponseReferenceFrameEcrEcef OnboardnavigationListResponseReferenceFrame = "ECR/ECEF"
	OnboardnavigationListResponseReferenceFrameTeme    OnboardnavigationListResponseReferenceFrame = "TEME"
	OnboardnavigationListResponseReferenceFrameItrf    OnboardnavigationListResponseReferenceFrame = "ITRF"
	OnboardnavigationListResponseReferenceFrameGcrf    OnboardnavigationListResponseReferenceFrame = "GCRF"
)

type OnboardnavigationQueryhelpResponse struct {
	AodrSupported         bool                                          `json:"aodrSupported"`
	ClassificationMarking string                                        `json:"classificationMarking"`
	Description           string                                        `json:"description"`
	HistorySupported      bool                                          `json:"historySupported"`
	Name                  string                                        `json:"name"`
	Parameters            []OnboardnavigationQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                      `json:"requiredRoles"`
	RestSupported         bool                                          `json:"restSupported"`
	SortSupported         bool                                          `json:"sortSupported"`
	TypeName              string                                        `json:"typeName"`
	Uri                   string                                        `json:"uri"`
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
func (r OnboardnavigationQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *OnboardnavigationQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnboardnavigationQueryhelpResponseParameter struct {
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
func (r OnboardnavigationQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *OnboardnavigationQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnboardnavigationListParams struct {
	// Start time of the sensor data, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnboardnavigationListParams]'s query parameters as
// `url.Values`.
func (r OnboardnavigationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnboardnavigationCountParams struct {
	// Start time of the sensor data, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnboardnavigationCountParams]'s query parameters as
// `url.Values`.
func (r OnboardnavigationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnboardnavigationNewBulkParams struct {
	Body []OnboardnavigationNewBulkParamsBody
	paramObj
}

func (r OnboardnavigationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *OnboardnavigationNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide spacecraft positional data derived from on-board
// navigational sensors. Navigational points are provided in kilometers and in a
// user specified reference frame, with ECI J2K being preferred.
//
// The properties ClassificationMarking, DataMode, Source, StartTime are required.
type OnboardnavigationNewBulkParamsBody struct {
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
	// Start time of the sensor data, in ISO 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// End time of the sensor data, in ISO 8601 UTC format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Unique identifier of the parent EphemerisSet, if this data is correlated with an
	// Ephemeris. If multiple nav sensor records are required, this ID may be
	// associated with each of those records if each is synced to the ephemeris points.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// Unique identifier of the primary satellite on-orbit object.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the last onboard state vector.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the data source to indicate the target object of
	// this record. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The last load time of the current star catalog onboard this spacecraft.
	StarCatLoadTime param.Opt[time.Time] `json:"starCatLoadTime,omitzero" format:"date-time"`
	// The name or identifier the star catalog in use.
	StarCatName param.Opt[string] `json:"starCatName,omitzero"`
	// Arrays of spacecraft delta position (X, Y, Z), in km, in the specified
	// referenceFrame, between the onboard state and the last accepted sensor position.
	DeltaPos [][]float64 `json:"deltaPos,omitzero"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard magnetometer, at each epoch.
	Mag [][]float64 `json:"mag,omitzero"`
	// The reference frame in which the sensor derived positions are provided. If the
	// referenceFrame is null it is assumed to be J2000. Note that this frame is
	// assumed to apply to all of the sensor data in this record. If onboard sensors
	// process positional estimates in different frames then separate records should be
	// generated.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard star tracker, at each epoch.
	StarTracker [][]float64 `json:"starTracker,omitzero"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard sun sensor, at each epoch.
	SunSensor [][]float64 `json:"sunSensor,omitzero"`
	// Array of epochs of the observations, in ISO 8601 UTC format. The epochs are
	// assumed to correspond to all sensor data in this record. If sensors do not share
	// a common epoch then separate records should be generated.
	Ts []time.Time `json:"ts,omitzero" format:"date-time"`
	paramObj
}

func (r OnboardnavigationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OnboardnavigationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnboardnavigationNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OnboardnavigationNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[OnboardnavigationNewBulkParamsBody](
		"referenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

type OnboardnavigationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Start time of the sensor data, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnboardnavigationTupleParams]'s query parameters as
// `url.Values`.
func (r OnboardnavigationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnboardnavigationUnvalidatedPublishParams struct {
	Body []OnboardnavigationUnvalidatedPublishParamsBody
	paramObj
}

func (r OnboardnavigationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *OnboardnavigationUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide spacecraft positional data derived from on-board
// navigational sensors. Navigational points are provided in kilometers and in a
// user specified reference frame, with ECI J2K being preferred.
//
// The properties ClassificationMarking, DataMode, Source, StartTime are required.
type OnboardnavigationUnvalidatedPublishParamsBody struct {
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
	// Start time of the sensor data, in ISO 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// End time of the sensor data, in ISO 8601 UTC format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Unique identifier of the parent EphemerisSet, if this data is correlated with an
	// Ephemeris. If multiple nav sensor records are required, this ID may be
	// associated with each of those records if each is synced to the ephemeris points.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// Unique identifier of the primary satellite on-orbit object.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the last onboard state vector.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the data source to indicate the target object of
	// this record. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The last load time of the current star catalog onboard this spacecraft.
	StarCatLoadTime param.Opt[time.Time] `json:"starCatLoadTime,omitzero" format:"date-time"`
	// The name or identifier the star catalog in use.
	StarCatName param.Opt[string] `json:"starCatName,omitzero"`
	// Arrays of spacecraft delta position (X, Y, Z), in km, in the specified
	// referenceFrame, between the onboard state and the last accepted sensor position.
	DeltaPos [][]float64 `json:"deltaPos,omitzero"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard magnetometer, at each epoch.
	Mag [][]float64 `json:"mag,omitzero"`
	// The reference frame in which the sensor derived positions are provided. If the
	// referenceFrame is null it is assumed to be J2000. Note that this frame is
	// assumed to apply to all of the sensor data in this record. If onboard sensors
	// process positional estimates in different frames then separate records should be
	// generated.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard star tracker, at each epoch.
	StarTracker [][]float64 `json:"starTracker,omitzero"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard sun sensor, at each epoch.
	SunSensor [][]float64 `json:"sunSensor,omitzero"`
	// Array of epochs of the observations, in ISO 8601 UTC format. The epochs are
	// assumed to correspond to all sensor data in this record. If sensors do not share
	// a common epoch then separate records should be generated.
	Ts []time.Time `json:"ts,omitzero" format:"date-time"`
	paramObj
}

func (r OnboardnavigationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OnboardnavigationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnboardnavigationUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OnboardnavigationUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[OnboardnavigationUnvalidatedPublishParamsBody](
		"referenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}
