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

// SwirService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSwirService] method instead.
type SwirService struct {
	Options []option.RequestOption
	History SwirHistoryService
}

// NewSwirService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSwirService(opts ...option.RequestOption) (r SwirService) {
	r = SwirService{}
	r.Options = opts
	r.History = NewSwirHistoryService(opts...)
	return
}

// Service operation to take a single SWIR record as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SwirService) New(ctx context.Context, body SwirNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/swir"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SwirService) List(ctx context.Context, query SwirListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SwirListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/swir"
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
func (r *SwirService) ListAutoPaging(ctx context.Context, query SwirListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SwirListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SwirService) Count(ctx context.Context, query SwirCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/swir/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of SWIR
// records as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *SwirService) NewBulk(ctx context.Context, body SwirNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/swir/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single SWIR record by its unique ID passed as a path
// parameter.
func (r *SwirService) Get(ctx context.Context, id string, query SwirGetParams, opts ...option.RequestOption) (res *SwirFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/swir/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SwirService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SwirQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/swir/queryhelp"
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
func (r *SwirService) Tuple(ctx context.Context, query SwirTupleParams, opts ...option.RequestOption) (res *[]SwirFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/swir/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Data representing observed short wave infrared (SWIR) measurements.
type SwirListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
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
	DataMode SwirListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Data timestamp in ISO8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// User comments concerning sensor or data limitations.
	BadWave string `json:"badWave"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the target on-orbit object.
	IDOnOrbit string `json:"idOnOrbit"`
	// Spacecraft WGS84 latitude, in degrees at obTime. -90 to 90 degrees (negative
	// values south of equator).
	Lat float64 `json:"lat"`
	// Location/name of the observing sensor.
	LocationName string `json:"locationName"`
	// Spacecraft WGS84 longitude at ob time, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Original object ID or Catalog Number provided by source.
	OrigObjectID string `json:"origObjectId"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// The angle, in degrees, between the target-to-observer vector and the
	// target-to-sun vector.
	SolarPhaseAngle float64 `json:"solarPhaseAngle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		BadWave               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDOnOrbit             respjson.Field
		Lat                   respjson.Field
		LocationName          respjson.Field
		Lon                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		SatNo                 respjson.Field
		SolarPhaseAngle       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SwirListResponse) RawJSON() string { return r.JSON.raw }
func (r *SwirListResponse) UnmarshalJSON(data []byte) error {
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
type SwirListResponseDataMode string

const (
	SwirListResponseDataModeReal      SwirListResponseDataMode = "REAL"
	SwirListResponseDataModeTest      SwirListResponseDataMode = "TEST"
	SwirListResponseDataModeSimulated SwirListResponseDataMode = "SIMULATED"
	SwirListResponseDataModeExercise  SwirListResponseDataMode = "EXERCISE"
)

type SwirQueryhelpResponse struct {
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
func (r SwirQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SwirQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwirNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
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
	DataMode SwirNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Data timestamp in ISO8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// User comments concerning sensor or data limitations.
	BadWave param.Opt[string] `json:"badWave,omitzero"`
	// Spacecraft WGS84 latitude, in degrees at obTime. -90 to 90 degrees (negative
	// values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Location/name of the observing sensor.
	LocationName param.Opt[string] `json:"locationName,omitzero"`
	// Spacecraft WGS84 longitude at ob time, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Original object ID or Catalog Number provided by source.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The angle, in degrees, between the target-to-observer vector and the
	// target-to-sun vector.
	SolarPhaseAngle param.Opt[float64] `json:"solarPhaseAngle,omitzero"`
	// Array of absolute flux measurement data, in Watts per square centimeter per
	// micron. This array should correspond with the same-sized array of wavelengths.
	AbsFluxes []float64 `json:"absFluxes,omitzero"`
	// Array of flux ratio data. This array should correspond with the same-sized array
	// of ratioWavelengths.
	FluxRatios []float64 `json:"fluxRatios,omitzero"`
	// Array of ratio wavelength data. This array should correspond with the same-sized
	// array of fluxRatios.
	RatioWavelengths []float64 `json:"ratioWavelengths,omitzero"`
	// Array of wavelengths, in microns. This array should correspond with the
	// same-sized array of absFluxes.
	Wavelengths []float64 `json:"wavelengths,omitzero"`
	paramObj
}

func (r SwirNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SwirNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SwirNewParams) UnmarshalJSON(data []byte) error {
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
type SwirNewParamsDataMode string

const (
	SwirNewParamsDataModeReal      SwirNewParamsDataMode = "REAL"
	SwirNewParamsDataModeTest      SwirNewParamsDataMode = "TEST"
	SwirNewParamsDataModeSimulated SwirNewParamsDataMode = "SIMULATED"
	SwirNewParamsDataModeExercise  SwirNewParamsDataMode = "EXERCISE"
)

type SwirListParams struct {
	// Data timestamp in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SwirListParams]'s query parameters as `url.Values`.
func (r SwirListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SwirCountParams struct {
	// Data timestamp in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SwirCountParams]'s query parameters as `url.Values`.
func (r SwirCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SwirNewBulkParams struct {
	Body []SwirNewBulkParamsBody
	paramObj
}

func (r SwirNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SwirNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Data representing observed short wave infrared (SWIR) measurements.
//
// The properties ClassificationMarking, DataMode, Source, Ts are required.
type SwirNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Data timestamp in ISO8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// User comments concerning sensor or data limitations.
	BadWave param.Opt[string] `json:"badWave,omitzero"`
	// Spacecraft WGS84 latitude, in degrees at obTime. -90 to 90 degrees (negative
	// values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Location/name of the observing sensor.
	LocationName param.Opt[string] `json:"locationName,omitzero"`
	// Spacecraft WGS84 longitude at ob time, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Original object ID or Catalog Number provided by source.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The angle, in degrees, between the target-to-observer vector and the
	// target-to-sun vector.
	SolarPhaseAngle param.Opt[float64] `json:"solarPhaseAngle,omitzero"`
	// Array of absolute flux measurement data, in Watts per square centimeter per
	// micron. This array should correspond with the same-sized array of wavelengths.
	AbsFluxes []float64 `json:"absFluxes,omitzero"`
	// Array of flux ratio data. This array should correspond with the same-sized array
	// of ratioWavelengths.
	FluxRatios []float64 `json:"fluxRatios,omitzero"`
	// Array of ratio wavelength data. This array should correspond with the same-sized
	// array of fluxRatios.
	RatioWavelengths []float64 `json:"ratioWavelengths,omitzero"`
	// Array of wavelengths, in microns. This array should correspond with the
	// same-sized array of absFluxes.
	Wavelengths []float64 `json:"wavelengths,omitzero"`
	paramObj
}

func (r SwirNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SwirNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SwirNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SwirNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SwirGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SwirGetParams]'s query parameters as `url.Values`.
func (r SwirGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SwirTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Data timestamp in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SwirTupleParams]'s query parameters as `url.Values`.
func (r SwirTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
