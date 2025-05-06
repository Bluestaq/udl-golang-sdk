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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// EquipmentRemarkService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEquipmentRemarkService] method instead.
type EquipmentRemarkService struct {
	Options []option.RequestOption
}

// NewEquipmentRemarkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEquipmentRemarkService(opts ...option.RequestOption) (r EquipmentRemarkService) {
	r = EquipmentRemarkService{}
	r.Options = opts
	return
}

// Service operation to take a single equipmentremark record as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *EquipmentRemarkService) New(ctx context.Context, body EquipmentRemarkNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/equipmentremark"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single equipmentremark record by its unique ID passed
// as a path parameter.
func (r *EquipmentRemarkService) Get(ctx context.Context, id string, query EquipmentRemarkGetParams, opts ...option.RequestOption) (res *EquipmentRemarkFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/equipmentremark/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EquipmentRemarkService) List(ctx context.Context, query EquipmentRemarkListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EquipmentRemarkAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/equipmentremark"
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
func (r *EquipmentRemarkService) ListAutoPaging(ctx context.Context, query EquipmentRemarkListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EquipmentRemarkAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EquipmentRemarkService) Count(ctx context.Context, query EquipmentRemarkCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/equipmentremark/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// equipmentremark records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *EquipmentRemarkService) NewBulk(ctx context.Context, body EquipmentRemarkNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/equipmentremark/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EquipmentRemarkService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/equipmentremark/queryhelp"
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
func (r *EquipmentRemarkService) Tuple(ctx context.Context, query EquipmentRemarkTupleParams, opts ...option.RequestOption) (res *[]EquipmentRemarkFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/equipmentremark/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
type EquipmentRemarkAbridged struct {
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
	DataMode EquipmentRemarkAbridgedDataMode `json:"dataMode,required"`
	// The ID of the Equipment to which this remark applies.
	IDEquipment string `json:"idEquipment,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Unique identifier of the Equipment Remark record from the originating system.
	AltRmkID string `json:"altRmkId"`
	// The remark type identifier. For example, the Mobility Air Forces (MAF) remark
	// code, defined in the Airfield Suitability and Restriction Report (ASRR).
	Code string `json:"code"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The name of the remark.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The remark type (e.g. Caution, Information, Misc, Restriction, etc.).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDEquipment           resp.Field
		Source                resp.Field
		Text                  resp.Field
		ID                    resp.Field
		AltRmkID              resp.Field
		Code                  resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Name                  resp.Field
		Origin                resp.Field
		Type                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EquipmentRemarkAbridged) RawJSON() string { return r.JSON.raw }
func (r *EquipmentRemarkAbridged) UnmarshalJSON(data []byte) error {
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
type EquipmentRemarkAbridgedDataMode string

const (
	EquipmentRemarkAbridgedDataModeReal      EquipmentRemarkAbridgedDataMode = "REAL"
	EquipmentRemarkAbridgedDataModeTest      EquipmentRemarkAbridgedDataMode = "TEST"
	EquipmentRemarkAbridgedDataModeSimulated EquipmentRemarkAbridgedDataMode = "SIMULATED"
	EquipmentRemarkAbridgedDataModeExercise  EquipmentRemarkAbridgedDataMode = "EXERCISE"
)

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
type EquipmentRemarkFull struct {
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
	DataMode EquipmentRemarkFullDataMode `json:"dataMode,required"`
	// The ID of the Equipment to which this remark applies.
	IDEquipment string `json:"idEquipment,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Unique identifier of the Equipment Remark record from the originating system.
	AltRmkID string `json:"altRmkId"`
	// The remark type identifier. For example, the Mobility Air Forces (MAF) remark
	// code, defined in the Airfield Suitability and Restriction Report (ASRR).
	Code string `json:"code"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The name of the remark.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The remark type (e.g. Caution, Information, Misc, Restriction, etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDEquipment           resp.Field
		Source                resp.Field
		Text                  resp.Field
		ID                    resp.Field
		AltRmkID              resp.Field
		Code                  resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Name                  resp.Field
		Origin                resp.Field
		Type                  resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EquipmentRemarkFull) RawJSON() string { return r.JSON.raw }
func (r *EquipmentRemarkFull) UnmarshalJSON(data []byte) error {
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
type EquipmentRemarkFullDataMode string

const (
	EquipmentRemarkFullDataModeReal      EquipmentRemarkFullDataMode = "REAL"
	EquipmentRemarkFullDataModeTest      EquipmentRemarkFullDataMode = "TEST"
	EquipmentRemarkFullDataModeSimulated EquipmentRemarkFullDataMode = "SIMULATED"
	EquipmentRemarkFullDataModeExercise  EquipmentRemarkFullDataMode = "EXERCISE"
)

type EquipmentRemarkNewParams struct {
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
	DataMode EquipmentRemarkNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The ID of the Equipment to which this remark applies.
	IDEquipment string `json:"idEquipment,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Unique identifier of the Equipment Remark record from the originating system.
	AltRmkID param.Opt[string] `json:"altRmkId,omitzero"`
	// The remark type identifier. For example, the Mobility Air Forces (MAF) remark
	// code, defined in the Airfield Suitability and Restriction Report (ASRR).
	Code param.Opt[string] `json:"code,omitzero"`
	// The name of the remark.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The remark type (e.g. Caution, Information, Misc, Restriction, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r EquipmentRemarkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EquipmentRemarkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type EquipmentRemarkNewParamsDataMode string

const (
	EquipmentRemarkNewParamsDataModeReal      EquipmentRemarkNewParamsDataMode = "REAL"
	EquipmentRemarkNewParamsDataModeTest      EquipmentRemarkNewParamsDataMode = "TEST"
	EquipmentRemarkNewParamsDataModeSimulated EquipmentRemarkNewParamsDataMode = "SIMULATED"
	EquipmentRemarkNewParamsDataModeExercise  EquipmentRemarkNewParamsDataMode = "EXERCISE"
)

type EquipmentRemarkGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EquipmentRemarkGetParams]'s query parameters as
// `url.Values`.
func (r EquipmentRemarkGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EquipmentRemarkListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EquipmentRemarkListParams]'s query parameters as
// `url.Values`.
func (r EquipmentRemarkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EquipmentRemarkCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EquipmentRemarkCountParams]'s query parameters as
// `url.Values`.
func (r EquipmentRemarkCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EquipmentRemarkNewBulkParams struct {
	Body []EquipmentRemarkNewBulkParamsBody
	paramObj
}

func (r EquipmentRemarkNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
//
// The properties ClassificationMarking, DataMode, IDEquipment, Source, Text are
// required.
type EquipmentRemarkNewBulkParamsBody struct {
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
	// The ID of the Equipment to which this remark applies.
	IDEquipment string `json:"idEquipment,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Unique identifier of the Equipment Remark record from the originating system.
	AltRmkID param.Opt[string] `json:"altRmkId,omitzero"`
	// The remark type identifier. For example, the Mobility Air Forces (MAF) remark
	// code, defined in the Airfield Suitability and Restriction Report (ASRR).
	Code param.Opt[string] `json:"code,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The name of the remark.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The remark type (e.g. Caution, Information, Misc, Restriction, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r EquipmentRemarkNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EquipmentRemarkNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EquipmentRemarkNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type EquipmentRemarkTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EquipmentRemarkTupleParams]'s query parameters as
// `url.Values`.
func (r EquipmentRemarkTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
