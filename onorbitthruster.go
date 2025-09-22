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

// OnorbitthrusterService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitthrusterService] method instead.
type OnorbitthrusterService struct {
	Options []option.RequestOption
}

// NewOnorbitthrusterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOnorbitthrusterService(opts ...option.RequestOption) (r OnorbitthrusterService) {
	r = OnorbitthrusterService{}
	r.Options = opts
	return
}

// Service operation to take a single OnorbitThruster as a POST body and ingest
// into the database. An OnorbitThruster is the association between an on-orbit
// spacecraft's engine and a particular on-orbit spacecraft. An Engine type may be
// associated with many different on-orbit spacecraft. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *OnorbitthrusterService) New(ctx context.Context, body OnorbitthrusterNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitthruster"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single OnorbitThruster. An OnorbitThruster is the
// association between an on-orbit spacecraft's engine and a particular on-orbit
// spacecraft. An Engine type may be associated with many different on-orbit
// spacecraft. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *OnorbitthrusterService) Update(ctx context.Context, id string, body OnorbitthrusterUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitthruster/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitthrusterService) List(ctx context.Context, query OnorbitthrusterListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitthrusterListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitthruster"
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
func (r *OnorbitthrusterService) ListAutoPaging(ctx context.Context, query OnorbitthrusterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitthrusterListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a OnorbitThruster object specified by the passed ID
// path parameter. An OnorbitThruster is the association between an on-orbit
// spacecraft's engine and a particular on-orbit spacecraft. An Engine type may be
// associated with many different on-orbit spacecraft. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *OnorbitthrusterService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitthruster/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to get a single OnorbitThruster record by its unique ID passed
// as a path parameter. An OnorbitThruster is the association between an on-orbit
// spacecraft's engine and a particular on-orbit spacecraft. An Engine type may be
// associated with many different on-orbit spacecraft.
func (r *OnorbitthrusterService) Get(ctx context.Context, id string, query OnorbitthrusterGetParams, opts ...option.RequestOption) (res *shared.OnorbitThrusterFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitthruster/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OnorbitthrusterListResponse struct {
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
	DataMode OnorbitthrusterListResponseDataMode `json:"dataMode,required"`
	// ID of the Engine.
	IDEngine string `json:"idEngine,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Known launch vehicle engines and their performance characteristics and limits. A
	// launch vehicle has 1 to many engines per stage.
	Engine EngineAbridged `json:"engine"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The number of engines/thrusters on the spacecraft of the type identified by
	// idEngine.
	Quantity int64 `json:"quantity"`
	// The type of thruster associated with this record (e.g. LAE, Hydrazine REA,
	// etc.).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDEngine              respjson.Field
		IDOnOrbit             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Engine                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Quantity              respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitthrusterListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitthrusterListResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitthrusterListResponseDataMode string

const (
	OnorbitthrusterListResponseDataModeReal      OnorbitthrusterListResponseDataMode = "REAL"
	OnorbitthrusterListResponseDataModeTest      OnorbitthrusterListResponseDataMode = "TEST"
	OnorbitthrusterListResponseDataModeSimulated OnorbitthrusterListResponseDataMode = "SIMULATED"
	OnorbitthrusterListResponseDataModeExercise  OnorbitthrusterListResponseDataMode = "EXERCISE"
)

type OnorbitthrusterNewParams struct {
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
	DataMode OnorbitthrusterNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the Engine.
	IDEngine string `json:"idEngine,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The number of engines/thrusters on the spacecraft of the type identified by
	// idEngine.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The type of thruster associated with this record (e.g. LAE, Hydrazine REA,
	// etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// Known launch vehicle engines and their performance characteristics and limits. A
	// launch vehicle has 1 to many engines per stage.
	Engine shared.EngineIngestParam `json:"engine,omitzero"`
	paramObj
}

func (r OnorbitthrusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitthrusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitthrusterNewParams) UnmarshalJSON(data []byte) error {
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
type OnorbitthrusterNewParamsDataMode string

const (
	OnorbitthrusterNewParamsDataModeReal      OnorbitthrusterNewParamsDataMode = "REAL"
	OnorbitthrusterNewParamsDataModeTest      OnorbitthrusterNewParamsDataMode = "TEST"
	OnorbitthrusterNewParamsDataModeSimulated OnorbitthrusterNewParamsDataMode = "SIMULATED"
	OnorbitthrusterNewParamsDataModeExercise  OnorbitthrusterNewParamsDataMode = "EXERCISE"
)

type OnorbitthrusterUpdateParams struct {
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
	DataMode OnorbitthrusterUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the Engine.
	IDEngine string `json:"idEngine,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The number of engines/thrusters on the spacecraft of the type identified by
	// idEngine.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The type of thruster associated with this record (e.g. LAE, Hydrazine REA,
	// etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// Known launch vehicle engines and their performance characteristics and limits. A
	// launch vehicle has 1 to many engines per stage.
	Engine shared.EngineIngestParam `json:"engine,omitzero"`
	paramObj
}

func (r OnorbitthrusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitthrusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitthrusterUpdateParams) UnmarshalJSON(data []byte) error {
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
type OnorbitthrusterUpdateParamsDataMode string

const (
	OnorbitthrusterUpdateParamsDataModeReal      OnorbitthrusterUpdateParamsDataMode = "REAL"
	OnorbitthrusterUpdateParamsDataModeTest      OnorbitthrusterUpdateParamsDataMode = "TEST"
	OnorbitthrusterUpdateParamsDataModeSimulated OnorbitthrusterUpdateParamsDataMode = "SIMULATED"
	OnorbitthrusterUpdateParamsDataModeExercise  OnorbitthrusterUpdateParamsDataMode = "EXERCISE"
)

type OnorbitthrusterListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitthrusterListParams]'s query parameters as
// `url.Values`.
func (r OnorbitthrusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitthrusterGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitthrusterGetParams]'s query parameters as
// `url.Values`.
func (r OnorbitthrusterGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
