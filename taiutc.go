// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// TaiUtcService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaiUtcService] method instead.
type TaiUtcService struct {
	Options []option.RequestOption
	History TaiUtcHistoryService
}

// NewTaiUtcService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaiUtcService(opts ...option.RequestOption) (r TaiUtcService) {
	r = TaiUtcService{}
	r.Options = opts
	r.History = NewTaiUtcHistoryService(opts...)
	return
}

// Service operation to take a single TAIUTC object as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *TaiUtcService) New(ctx context.Context, body TaiUtcNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/taiutc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single TAIUTC object. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *TaiUtcService) Update(ctx context.Context, id string, body TaiUtcUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/taiutc/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TaiUtcService) List(ctx context.Context, query TaiUtcListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[TaiUtcListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/taiutc"
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
func (r *TaiUtcService) ListAutoPaging(ctx context.Context, query TaiUtcListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[TaiUtcListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an TAIUTC object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance. Note, delete operations do not remove data
// from historical or publish/subscribe stores.
func (r *TaiUtcService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/taiutc/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *TaiUtcService) Count(ctx context.Context, query TaiUtcCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/taiutc/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single TAIUTC record by its unique ID passed as a
// path parameter.
func (r *TaiUtcService) Get(ctx context.Context, id string, query TaiUtcGetParams, opts ...option.RequestOption) (res *TaiutcFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/taiutc/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *TaiUtcService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *TaiUtcQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/taiutc/queryhelp"
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
func (r *TaiUtcService) Tuple(ctx context.Context, query TaiUtcTupleParams, opts ...option.RequestOption) (res *[]TaiutcFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/taiutc/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// International Atomic Time (TAI) is a statistical atomic time scale based on a
// large number of clocks operating at standards laboratories around the world that
// is maintained by the Bureau International des Poids et Mesures; its unit
// interval is exactly one SI second at sea level. The origin of TAI is such that
// UT1-TAI is approximately 0 (zero) on January 1, 1958. TAI is not adjusted for
// leap seconds. Coordinated Universal Time (UTC) is defined by the CCIR
// Recommendation 460-4 (1986). It differs from TAI by the total number of leap
// seconds, so that UT1-UTC stays smaller than 0.9s in absolute value. The decision
// to introduce a leap second in UTC is the responsibility of the International
// Earth Rotation Service (IERS). According to the CCIR Recommendation, first
// preference is given to the opportunities at the end of December and June, and
// second preference to those at the end of March and September. Since the system
// was introduced in 1972, only dates in June and December have been used. TAI is
// expressed in terms of UTC by the relation TAI = UTC + dAT, where dAT is the
// total algebraic sum of leap seconds. The first leap second was introduced on
// June 30, 1972. The historical list of leap seconds can be found in this table.
type TaiUtcListResponse struct {
	// Effective date/time for the leap second adjustment.
	AdjustmentDate time.Time `json:"adjustmentDate,required" format:"date-time"`
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
	DataMode TaiUtcListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Multiplication factor of the leap second adjustment.
	MultiplicationFactor float64 `json:"multiplicationFactor"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Total/cumulative offset between TAI and UTC time as of adjustmentDate, in
	// seconds.
	TaiUtc float64 `json:"taiUTC"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdjustmentDate        respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		MultiplicationFactor  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RawFileUri            respjson.Field
		TaiUtc                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaiUtcListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaiUtcListResponse) UnmarshalJSON(data []byte) error {
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
type TaiUtcListResponseDataMode string

const (
	TaiUtcListResponseDataModeReal      TaiUtcListResponseDataMode = "REAL"
	TaiUtcListResponseDataModeTest      TaiUtcListResponseDataMode = "TEST"
	TaiUtcListResponseDataModeSimulated TaiUtcListResponseDataMode = "SIMULATED"
	TaiUtcListResponseDataModeExercise  TaiUtcListResponseDataMode = "EXERCISE"
)

type TaiUtcQueryhelpResponse struct {
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
func (r TaiUtcQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *TaiUtcQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaiUtcNewParams struct {
	// Effective date/time for the leap second adjustment.
	AdjustmentDate time.Time `json:"adjustmentDate,required" format:"date-time"`
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
	DataMode TaiUtcNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Multiplication factor of the leap second adjustment.
	MultiplicationFactor param.Opt[float64] `json:"multiplicationFactor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Total/cumulative offset between TAI and UTC time as of adjustmentDate, in
	// seconds.
	TaiUtc param.Opt[float64] `json:"taiUTC,omitzero"`
	paramObj
}

func (r TaiUtcNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TaiUtcNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaiUtcNewParams) UnmarshalJSON(data []byte) error {
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
type TaiUtcNewParamsDataMode string

const (
	TaiUtcNewParamsDataModeReal      TaiUtcNewParamsDataMode = "REAL"
	TaiUtcNewParamsDataModeTest      TaiUtcNewParamsDataMode = "TEST"
	TaiUtcNewParamsDataModeSimulated TaiUtcNewParamsDataMode = "SIMULATED"
	TaiUtcNewParamsDataModeExercise  TaiUtcNewParamsDataMode = "EXERCISE"
)

type TaiUtcUpdateParams struct {
	// Effective date/time for the leap second adjustment.
	AdjustmentDate time.Time `json:"adjustmentDate,required" format:"date-time"`
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
	DataMode TaiUtcUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Multiplication factor of the leap second adjustment.
	MultiplicationFactor param.Opt[float64] `json:"multiplicationFactor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Total/cumulative offset between TAI and UTC time as of adjustmentDate, in
	// seconds.
	TaiUtc param.Opt[float64] `json:"taiUTC,omitzero"`
	paramObj
}

func (r TaiUtcUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TaiUtcUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaiUtcUpdateParams) UnmarshalJSON(data []byte) error {
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
type TaiUtcUpdateParamsDataMode string

const (
	TaiUtcUpdateParamsDataModeReal      TaiUtcUpdateParamsDataMode = "REAL"
	TaiUtcUpdateParamsDataModeTest      TaiUtcUpdateParamsDataMode = "TEST"
	TaiUtcUpdateParamsDataModeSimulated TaiUtcUpdateParamsDataMode = "SIMULATED"
	TaiUtcUpdateParamsDataModeExercise  TaiUtcUpdateParamsDataMode = "EXERCISE"
)

type TaiUtcListParams struct {
	// Effective date/time for the leap second adjustment. Must be a unique value
	// across all TAIUTC datasets. (YYYY-MM-DDTHH:MM:SS.sssZ)
	AdjustmentDate time.Time        `query:"adjustmentDate,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaiUtcListParams]'s query parameters as `url.Values`.
func (r TaiUtcListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TaiUtcCountParams struct {
	// Effective date/time for the leap second adjustment. Must be a unique value
	// across all TAIUTC datasets. (YYYY-MM-DDTHH:MM:SS.sssZ)
	AdjustmentDate time.Time        `query:"adjustmentDate,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaiUtcCountParams]'s query parameters as `url.Values`.
func (r TaiUtcCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TaiUtcGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaiUtcGetParams]'s query parameters as `url.Values`.
func (r TaiUtcGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TaiUtcTupleParams struct {
	// Effective date/time for the leap second adjustment. Must be a unique value
	// across all TAIUTC datasets. (YYYY-MM-DDTHH:MM:SS.sssZ)
	AdjustmentDate time.Time `query:"adjustmentDate,required" format:"date-time" json:"-"`
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaiUtcTupleParams]'s query parameters as `url.Values`.
func (r TaiUtcTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
