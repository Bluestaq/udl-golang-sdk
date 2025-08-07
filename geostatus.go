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

// GeoStatusService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeoStatusService] method instead.
type GeoStatusService struct {
	Options []option.RequestOption
	History GeoStatusHistoryService
}

// NewGeoStatusService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGeoStatusService(opts ...option.RequestOption) (r GeoStatusService) {
	r = GeoStatusService{}
	r.Options = opts
	r.History = NewGeoStatusHistoryService(opts...)
	return
}

// Service operation to take a single GEOStatus record as a POST body and ingest
// into the database. This operation is not intended to be used for automated feeds
// into UDL. Data providers should contact the UDL team for specific role
// assignments and for instructions on setting up a permanent feed through an
// alternate mechanism.
func (r *GeoStatusService) New(ctx context.Context, body GeoStatusNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/geostatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *GeoStatusService) List(ctx context.Context, query GeoStatusListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GeoStatusListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/geostatus"
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
func (r *GeoStatusService) ListAutoPaging(ctx context.Context, query GeoStatusListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GeoStatusListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *GeoStatusService) Count(ctx context.Context, query GeoStatusCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/geostatus/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// GEOStatus records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *GeoStatusService) NewBulk(ctx context.Context, body GeoStatusNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/geostatus/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single GEOStatus record by its unique ID passed as a
// path parameter.
func (r *GeoStatusService) Get(ctx context.Context, id string, query GeoStatusGetParams, opts ...option.RequestOption) (res *GeoStatusFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/geostatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *GeoStatusService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *GeoStatusQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/geostatus/queryhelp"
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
func (r *GeoStatusService) Tuple(ctx context.Context, query GeoStatusTupleParams, opts ...option.RequestOption) (res *[]GeoStatusFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/geostatus/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information for the specified on-orbit GEO spacecraft, including status,
// expected longitude limits, and drift rates.
type GeoStatusListResponse struct {
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
	DataMode GeoStatusListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Indicates the confidence level in the entry. (Low, Medium, High).
	ConfidenceLevel string `json:"confidenceLevel"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the object on-orbit object.
	IDOnOrbit string `json:"idOnOrbit"`
	// Maximum longitude for this object. WGS-84 longitude of the spacecraft position,
	// in degrees. 0 to 360 degrees.
	LongitudeMax float64 `json:"longitudeMax"`
	// Minimum longitude for this object. WGS-84 longitude of the spacecraft position,
	// in degrees. 0 to 360 degrees.
	LongitudeMin float64 `json:"longitudeMin"`
	// Corrective or overriding long term trend for longitudinal change in degrees/day.
	LongitudeRate float64 `json:"longitudeRate"`
	// Lost space object indicator. (True or False).
	LostFlag bool `json:"lostFlag"`
	// Space object status. (Active, Dead, Unknown).
	ObjectStatus string `json:"objectStatus"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided to indicate the target onorbit. This may be an
	// internal identifier and not necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Geosynchronous plane changing status. (Current, Never, Former, Future).
	PlaneChangeStatus string `json:"planeChangeStatus"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Objects displacement from geostationary orbit in deg^2/day^2.
	RelativeEnergy float64 `json:"relativeEnergy"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sine of inclination times the cosine of right ascension.
	Sc float64 `json:"sc"`
	// Semi-annual correction. (True or False).
	SemiAnnualCorrFlag bool `json:"semiAnnualCorrFlag"`
	// Sine of inclination times the sine of right ascension.
	SS float64 `json:"ss"`
	// Indicates the trough (gravity well) or drift direction of a space object:
	//
	// 255 - Influenced by 255° longitude trough.
	//
	// 75 - Influenced by 75° longitude trough.
	//
	// Both - Oscillating between both 255 and 75 troughs.
	//
	// East - Drifting eastward; large relative energy and a period less than 1436.1
	// minutes.
	//
	// West - Drifting westward; large relative energy and a period greater than 1436.2
	// minutes.
	TroughType string `json:"troughType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ConfidenceLevel       respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDOnOrbit             respjson.Field
		LongitudeMax          respjson.Field
		LongitudeMin          respjson.Field
		LongitudeRate         respjson.Field
		LostFlag              respjson.Field
		ObjectStatus          respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		PlaneChangeStatus     respjson.Field
		RawFileUri            respjson.Field
		RelativeEnergy        respjson.Field
		SatNo                 respjson.Field
		Sc                    respjson.Field
		SemiAnnualCorrFlag    respjson.Field
		SS                    respjson.Field
		TroughType            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeoStatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *GeoStatusListResponse) UnmarshalJSON(data []byte) error {
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
type GeoStatusListResponseDataMode string

const (
	GeoStatusListResponseDataModeReal      GeoStatusListResponseDataMode = "REAL"
	GeoStatusListResponseDataModeTest      GeoStatusListResponseDataMode = "TEST"
	GeoStatusListResponseDataModeSimulated GeoStatusListResponseDataMode = "SIMULATED"
	GeoStatusListResponseDataModeExercise  GeoStatusListResponseDataMode = "EXERCISE"
)

type GeoStatusQueryhelpResponse struct {
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
func (r GeoStatusQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *GeoStatusQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeoStatusNewParams struct {
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
	DataMode GeoStatusNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Indicates the confidence level in the entry. (Low, Medium, High).
	ConfidenceLevel param.Opt[string] `json:"confidenceLevel,omitzero"`
	// Maximum longitude for this object. WGS-84 longitude of the spacecraft position,
	// in degrees. 0 to 360 degrees.
	LongitudeMax param.Opt[float64] `json:"longitudeMax,omitzero"`
	// Minimum longitude for this object. WGS-84 longitude of the spacecraft position,
	// in degrees. 0 to 360 degrees.
	LongitudeMin param.Opt[float64] `json:"longitudeMin,omitzero"`
	// Corrective or overriding long term trend for longitudinal change in degrees/day.
	LongitudeRate param.Opt[float64] `json:"longitudeRate,omitzero"`
	// Lost space object indicator. (True or False).
	LostFlag param.Opt[bool] `json:"lostFlag,omitzero"`
	// Space object status. (Active, Dead, Unknown).
	ObjectStatus param.Opt[string] `json:"objectStatus,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided to indicate the target onorbit. This may be an
	// internal identifier and not necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Geosynchronous plane changing status. (Current, Never, Former, Future).
	PlaneChangeStatus param.Opt[string] `json:"planeChangeStatus,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Objects displacement from geostationary orbit in deg^2/day^2.
	RelativeEnergy param.Opt[float64] `json:"relativeEnergy,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sine of inclination times the cosine of right ascension.
	Sc param.Opt[float64] `json:"sc,omitzero"`
	// Semi-annual correction. (True or False).
	SemiAnnualCorrFlag param.Opt[bool] `json:"semiAnnualCorrFlag,omitzero"`
	// Sine of inclination times the sine of right ascension.
	SS param.Opt[float64] `json:"ss,omitzero"`
	// Indicates the trough (gravity well) or drift direction of a space object:
	//
	// 255 - Influenced by 255° longitude trough.
	//
	// 75 - Influenced by 75° longitude trough.
	//
	// Both - Oscillating between both 255 and 75 troughs.
	//
	// East - Drifting eastward; large relative energy and a period less than 1436.1
	// minutes.
	//
	// West - Drifting westward; large relative energy and a period greater than 1436.2
	// minutes.
	TroughType param.Opt[string] `json:"troughType,omitzero"`
	paramObj
}

func (r GeoStatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GeoStatusNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeoStatusNewParams) UnmarshalJSON(data []byte) error {
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
type GeoStatusNewParamsDataMode string

const (
	GeoStatusNewParamsDataModeReal      GeoStatusNewParamsDataMode = "REAL"
	GeoStatusNewParamsDataModeTest      GeoStatusNewParamsDataMode = "TEST"
	GeoStatusNewParamsDataModeSimulated GeoStatusNewParamsDataMode = "SIMULATED"
	GeoStatusNewParamsDataModeExercise  GeoStatusNewParamsDataMode = "EXERCISE"
)

type GeoStatusListParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GeoStatusListParams]'s query parameters as `url.Values`.
func (r GeoStatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeoStatusCountParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GeoStatusCountParams]'s query parameters as `url.Values`.
func (r GeoStatusCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeoStatusNewBulkParams struct {
	Body []GeoStatusNewBulkParamsBody
	paramObj
}

func (r GeoStatusNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *GeoStatusNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Information for the specified on-orbit GEO spacecraft, including status,
// expected longitude limits, and drift rates.
//
// The properties ClassificationMarking, DataMode, Source are required.
type GeoStatusNewBulkParamsBody struct {
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
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Indicates the confidence level in the entry. (Low, Medium, High).
	ConfidenceLevel param.Opt[string] `json:"confidenceLevel,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of the object on-orbit object.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Maximum longitude for this object. WGS-84 longitude of the spacecraft position,
	// in degrees. 0 to 360 degrees.
	LongitudeMax param.Opt[float64] `json:"longitudeMax,omitzero"`
	// Minimum longitude for this object. WGS-84 longitude of the spacecraft position,
	// in degrees. 0 to 360 degrees.
	LongitudeMin param.Opt[float64] `json:"longitudeMin,omitzero"`
	// Corrective or overriding long term trend for longitudinal change in degrees/day.
	LongitudeRate param.Opt[float64] `json:"longitudeRate,omitzero"`
	// Lost space object indicator. (True or False).
	LostFlag param.Opt[bool] `json:"lostFlag,omitzero"`
	// Space object status. (Active, Dead, Unknown).
	ObjectStatus param.Opt[string] `json:"objectStatus,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided to indicate the target onorbit. This may be an
	// internal identifier and not necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Geosynchronous plane changing status. (Current, Never, Former, Future).
	PlaneChangeStatus param.Opt[string] `json:"planeChangeStatus,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Objects displacement from geostationary orbit in deg^2/day^2.
	RelativeEnergy param.Opt[float64] `json:"relativeEnergy,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sine of inclination times the cosine of right ascension.
	Sc param.Opt[float64] `json:"sc,omitzero"`
	// Semi-annual correction. (True or False).
	SemiAnnualCorrFlag param.Opt[bool] `json:"semiAnnualCorrFlag,omitzero"`
	// Sine of inclination times the sine of right ascension.
	SS param.Opt[float64] `json:"ss,omitzero"`
	// Indicates the trough (gravity well) or drift direction of a space object:
	//
	// 255 - Influenced by 255° longitude trough.
	//
	// 75 - Influenced by 75° longitude trough.
	//
	// Both - Oscillating between both 255 and 75 troughs.
	//
	// East - Drifting eastward; large relative energy and a period less than 1436.1
	// minutes.
	//
	// West - Drifting westward; large relative energy and a period greater than 1436.2
	// minutes.
	TroughType param.Opt[string] `json:"troughType,omitzero"`
	paramObj
}

func (r GeoStatusNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow GeoStatusNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeoStatusNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GeoStatusNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type GeoStatusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GeoStatusGetParams]'s query parameters as `url.Values`.
func (r GeoStatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeoStatusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GeoStatusTupleParams]'s query parameters as `url.Values`.
func (r GeoStatusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
