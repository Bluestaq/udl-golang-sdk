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

// SolarArrayDetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolarArrayDetailService] method instead.
type SolarArrayDetailService struct {
	Options []option.RequestOption
}

// NewSolarArrayDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSolarArrayDetailService(opts ...option.RequestOption) (r SolarArrayDetailService) {
	r = SolarArrayDetailService{}
	r.Options = opts
	return
}

// Service operation to take a single SolarArrayDetails as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance. A SolarArray may have
// multiple details records compiled by various sources.
func (r *SolarArrayDetailService) New(ctx context.Context, body SolarArrayDetailNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/solararraydetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single SolarArrayDetails. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance. A SolarArray may have multiple details records compiled by various
// sources.
func (r *SolarArrayDetailService) Update(ctx context.Context, id string, body SolarArrayDetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/solararraydetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SolarArrayDetailService) List(ctx context.Context, query SolarArrayDetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SolarArrayDetailListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/solararraydetails"
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
func (r *SolarArrayDetailService) ListAutoPaging(ctx context.Context, query SolarArrayDetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SolarArrayDetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a SolarArrayDetails object specified by the passed
// ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance. A SolarArray may have
// multiple details records compiled by various sources.
func (r *SolarArrayDetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/solararraydetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to get a single SolarArrayDetails record by its unique ID
// passed as a path parameter. A SolarArray may have multiple details records
// compiled by various sources.
func (r *SolarArrayDetailService) Get(ctx context.Context, id string, query SolarArrayDetailGetParams, opts ...option.RequestOption) (res *shared.SolarArrayDetailsFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/solararraydetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of Information on spacecraft SolarArrayDetails. A
// SolarArray may have multiple details records compiled by various sources.
type SolarArrayDetailListResponse struct {
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
	DataMode SolarArrayDetailListResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the parent SolarArray.
	IDSolarArray string `json:"idSolarArray,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Solar Array area in square meters.
	Area float64 `json:"area"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Solar array description/notes.
	Description string `json:"description"`
	// Solar array junction technology (e.g. Triple).
	JunctionTechnology string `json:"junctionTechnology"`
	// Unique identifier of the organization that manufactures the solar array.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Solar Array span in meters.
	Span float64 `json:"span"`
	// Solar array technology (e.g. Ga-As).
	Technology string `json:"technology"`
	// Type of solar array (e.g. U Shaped).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSolarArray          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Area                  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Description           respjson.Field
		JunctionTechnology    respjson.Field
		ManufacturerOrgID     respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Span                  respjson.Field
		Technology            respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolarArrayDetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *SolarArrayDetailListResponse) UnmarshalJSON(data []byte) error {
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
type SolarArrayDetailListResponseDataMode string

const (
	SolarArrayDetailListResponseDataModeReal      SolarArrayDetailListResponseDataMode = "REAL"
	SolarArrayDetailListResponseDataModeTest      SolarArrayDetailListResponseDataMode = "TEST"
	SolarArrayDetailListResponseDataModeSimulated SolarArrayDetailListResponseDataMode = "SIMULATED"
	SolarArrayDetailListResponseDataModeExercise  SolarArrayDetailListResponseDataMode = "EXERCISE"
)

type SolarArrayDetailNewParams struct {
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
	DataMode SolarArrayDetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent SolarArray.
	IDSolarArray string `json:"idSolarArray,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Solar Array area in square meters.
	Area param.Opt[float64] `json:"area,omitzero"`
	// Solar array description/notes.
	Description param.Opt[string] `json:"description,omitzero"`
	// Solar array junction technology (e.g. Triple).
	JunctionTechnology param.Opt[string] `json:"junctionTechnology,omitzero"`
	// Unique identifier of the organization that manufactures the solar array.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Solar Array span in meters.
	Span param.Opt[float64] `json:"span,omitzero"`
	// Solar array technology (e.g. Ga-As).
	Technology param.Opt[string] `json:"technology,omitzero"`
	// Type of solar array (e.g. U Shaped).
	Type param.Opt[string] `json:"type,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SolarArrayDetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SolarArrayDetailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolarArrayDetailNewParams) UnmarshalJSON(data []byte) error {
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
type SolarArrayDetailNewParamsDataMode string

const (
	SolarArrayDetailNewParamsDataModeReal      SolarArrayDetailNewParamsDataMode = "REAL"
	SolarArrayDetailNewParamsDataModeTest      SolarArrayDetailNewParamsDataMode = "TEST"
	SolarArrayDetailNewParamsDataModeSimulated SolarArrayDetailNewParamsDataMode = "SIMULATED"
	SolarArrayDetailNewParamsDataModeExercise  SolarArrayDetailNewParamsDataMode = "EXERCISE"
)

type SolarArrayDetailUpdateParams struct {
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
	DataMode SolarArrayDetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent SolarArray.
	IDSolarArray string `json:"idSolarArray,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Solar Array area in square meters.
	Area param.Opt[float64] `json:"area,omitzero"`
	// Solar array description/notes.
	Description param.Opt[string] `json:"description,omitzero"`
	// Solar array junction technology (e.g. Triple).
	JunctionTechnology param.Opt[string] `json:"junctionTechnology,omitzero"`
	// Unique identifier of the organization that manufactures the solar array.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Solar Array span in meters.
	Span param.Opt[float64] `json:"span,omitzero"`
	// Solar array technology (e.g. Ga-As).
	Technology param.Opt[string] `json:"technology,omitzero"`
	// Type of solar array (e.g. U Shaped).
	Type param.Opt[string] `json:"type,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SolarArrayDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SolarArrayDetailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolarArrayDetailUpdateParams) UnmarshalJSON(data []byte) error {
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
type SolarArrayDetailUpdateParamsDataMode string

const (
	SolarArrayDetailUpdateParamsDataModeReal      SolarArrayDetailUpdateParamsDataMode = "REAL"
	SolarArrayDetailUpdateParamsDataModeTest      SolarArrayDetailUpdateParamsDataMode = "TEST"
	SolarArrayDetailUpdateParamsDataModeSimulated SolarArrayDetailUpdateParamsDataMode = "SIMULATED"
	SolarArrayDetailUpdateParamsDataModeExercise  SolarArrayDetailUpdateParamsDataMode = "EXERCISE"
)

type SolarArrayDetailListParams struct {
	// (One or more of fields 'classificationMarking, dataMode, source' are required.)
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking param.Opt[string] `query:"classificationMarking,omitzero" json:"-"`
	// (One or more of fields 'classificationMarking, dataMode, source' are required.)
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data. (REAL,
	// TEST, SIMULATED, or EXERCISE)
	DataMode    param.Opt[string] `query:"dataMode,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'classificationMarking, dataMode, source' are required.)
	// Source of the data.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SolarArrayDetailListParams]'s query parameters as
// `url.Values`.
func (r SolarArrayDetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SolarArrayDetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SolarArrayDetailGetParams]'s query parameters as
// `url.Values`.
func (r SolarArrayDetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
