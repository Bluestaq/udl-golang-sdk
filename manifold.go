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

// ManifoldService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManifoldService] method instead.
type ManifoldService struct {
	Options []option.RequestOption
}

// NewManifoldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewManifoldService(opts ...option.RequestOption) (r ManifoldService) {
	r = ManifoldService{}
	r.Options = opts
	return
}

// Service operation to take a single Manifold as a POST body and ingest into the
// database. A manifold represents a set of possible/theoretical orbits for an
// object of interest based on a delta V and delta T. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *ManifoldService) New(ctx context.Context, body ManifoldNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/manifold"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Manifold. A manifold represents a set of
// possible/theoretical orbits for an object of interest based on a delta V and
// delta T. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *ManifoldService) Update(ctx context.Context, id string, body ManifoldUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/manifold/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ManifoldService) List(ctx context.Context, query ManifoldListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ManifoldListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/manifold"
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
func (r *ManifoldService) ListAutoPaging(ctx context.Context, query ManifoldListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ManifoldListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Manifold object specified by the passed ID path
// parameter. A manifold represents a set of possible/theoretical orbits for an
// object of interest based on a delta V and delta T. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *ManifoldService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/manifold/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ManifoldService) Count(ctx context.Context, query ManifoldCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/manifold/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple Manifolds as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *ManifoldService) NewBulk(ctx context.Context, body ManifoldNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/manifold/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Manifold record by its unique ID passed as a
// path parameter. A manifold represents a set of possible/theoretical orbits for
// an object of interest based on a delta V and delta T.
func (r *ManifoldService) Get(ctx context.Context, id string, query ManifoldGetParams, opts ...option.RequestOption) (res *ManifoldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/manifold/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ManifoldService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *ManifoldQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/manifold/queryhelp"
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
func (r *ManifoldService) Tuple(ctx context.Context, query ManifoldTupleParams, opts ...option.RequestOption) (res *[]ManifoldTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/manifold/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A manifold represents a set of possible/theoretical orbits for an object of
// interest based on a delta V and delta T.
type ManifoldListResponse struct {
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
	DataMode ManifoldListResponseDataMode `json:"dataMode,required"`
	// ID of the parent object of interest.
	IDObjectOfInterest string `json:"idObjectOfInterest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Applied delta V duration for this manifold's calculations in seconds.
	DeltaT float64 `json:"deltaT"`
	// Applied delta V for this manifold's calculations, in km/sec.
	DeltaV float64 `json:"deltaV"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Status of the manifold and its associated ManifoldElsets (e.g. PENDING,
	// COMPLETE). PENDING status means element set generation is in progress and
	// COMPLETE indicates all ManifoldElsets have been generated.
	Status string `json:"status"`
	// Weight or probability of this manifold for prioritization purposes, between 0
	// and 1.
	Weight float64 `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDObjectOfInterest    respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeltaT                respjson.Field
		DeltaV                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Status                respjson.Field
		Weight                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManifoldListResponse) RawJSON() string { return r.JSON.raw }
func (r *ManifoldListResponse) UnmarshalJSON(data []byte) error {
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
type ManifoldListResponseDataMode string

const (
	ManifoldListResponseDataModeReal      ManifoldListResponseDataMode = "REAL"
	ManifoldListResponseDataModeTest      ManifoldListResponseDataMode = "TEST"
	ManifoldListResponseDataModeSimulated ManifoldListResponseDataMode = "SIMULATED"
	ManifoldListResponseDataModeExercise  ManifoldListResponseDataMode = "EXERCISE"
)

// A manifold represents a set of possible/theoretical orbits for an object of
// interest based on a delta V and delta T.
type ManifoldGetResponse struct {
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
	DataMode ManifoldGetResponseDataMode `json:"dataMode,required"`
	// ID of the parent object of interest.
	IDObjectOfInterest string `json:"idObjectOfInterest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Applied delta V duration for this manifold's calculations in seconds.
	DeltaT float64 `json:"deltaT"`
	// Applied delta V for this manifold's calculations, in km/sec.
	DeltaV float64 `json:"deltaV"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Status of the manifold and its associated ManifoldElsets (e.g. PENDING,
	// COMPLETE). PENDING status means element set generation is in progress and
	// COMPLETE indicates all ManifoldElsets have been generated.
	Status string `json:"status"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// Weight or probability of this manifold for prioritization purposes, between 0
	// and 1.
	Weight float64 `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDObjectOfInterest    respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeltaT                respjson.Field
		DeltaV                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Status                respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Weight                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManifoldGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ManifoldGetResponse) UnmarshalJSON(data []byte) error {
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
type ManifoldGetResponseDataMode string

const (
	ManifoldGetResponseDataModeReal      ManifoldGetResponseDataMode = "REAL"
	ManifoldGetResponseDataModeTest      ManifoldGetResponseDataMode = "TEST"
	ManifoldGetResponseDataModeSimulated ManifoldGetResponseDataMode = "SIMULATED"
	ManifoldGetResponseDataModeExercise  ManifoldGetResponseDataMode = "EXERCISE"
)

type ManifoldQueryhelpResponse struct {
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
func (r ManifoldQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *ManifoldQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A manifold represents a set of possible/theoretical orbits for an object of
// interest based on a delta V and delta T.
type ManifoldTupleResponse struct {
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
	DataMode ManifoldTupleResponseDataMode `json:"dataMode,required"`
	// ID of the parent object of interest.
	IDObjectOfInterest string `json:"idObjectOfInterest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Applied delta V duration for this manifold's calculations in seconds.
	DeltaT float64 `json:"deltaT"`
	// Applied delta V for this manifold's calculations, in km/sec.
	DeltaV float64 `json:"deltaV"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Status of the manifold and its associated ManifoldElsets (e.g. PENDING,
	// COMPLETE). PENDING status means element set generation is in progress and
	// COMPLETE indicates all ManifoldElsets have been generated.
	Status string `json:"status"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// Weight or probability of this manifold for prioritization purposes, between 0
	// and 1.
	Weight float64 `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDObjectOfInterest    respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeltaT                respjson.Field
		DeltaV                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Status                respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Weight                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManifoldTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ManifoldTupleResponse) UnmarshalJSON(data []byte) error {
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
type ManifoldTupleResponseDataMode string

const (
	ManifoldTupleResponseDataModeReal      ManifoldTupleResponseDataMode = "REAL"
	ManifoldTupleResponseDataModeTest      ManifoldTupleResponseDataMode = "TEST"
	ManifoldTupleResponseDataModeSimulated ManifoldTupleResponseDataMode = "SIMULATED"
	ManifoldTupleResponseDataModeExercise  ManifoldTupleResponseDataMode = "EXERCISE"
)

type ManifoldNewParams struct {
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
	DataMode ManifoldNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the parent object of interest.
	IDObjectOfInterest string `json:"idObjectOfInterest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Applied delta V duration for this manifold's calculations in seconds.
	DeltaT param.Opt[float64] `json:"deltaT,omitzero"`
	// Applied delta V for this manifold's calculations, in km/sec.
	DeltaV param.Opt[float64] `json:"deltaV,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Status of the manifold and its associated ManifoldElsets (e.g. PENDING,
	// COMPLETE). PENDING status means element set generation is in progress and
	// COMPLETE indicates all ManifoldElsets have been generated.
	Status param.Opt[string] `json:"status,omitzero"`
	// Weight or probability of this manifold for prioritization purposes, between 0
	// and 1.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r ManifoldNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ManifoldNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManifoldNewParams) UnmarshalJSON(data []byte) error {
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
type ManifoldNewParamsDataMode string

const (
	ManifoldNewParamsDataModeReal      ManifoldNewParamsDataMode = "REAL"
	ManifoldNewParamsDataModeTest      ManifoldNewParamsDataMode = "TEST"
	ManifoldNewParamsDataModeSimulated ManifoldNewParamsDataMode = "SIMULATED"
	ManifoldNewParamsDataModeExercise  ManifoldNewParamsDataMode = "EXERCISE"
)

type ManifoldUpdateParams struct {
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
	DataMode ManifoldUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the parent object of interest.
	IDObjectOfInterest string `json:"idObjectOfInterest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Applied delta V duration for this manifold's calculations in seconds.
	DeltaT param.Opt[float64] `json:"deltaT,omitzero"`
	// Applied delta V for this manifold's calculations, in km/sec.
	DeltaV param.Opt[float64] `json:"deltaV,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Status of the manifold and its associated ManifoldElsets (e.g. PENDING,
	// COMPLETE). PENDING status means element set generation is in progress and
	// COMPLETE indicates all ManifoldElsets have been generated.
	Status param.Opt[string] `json:"status,omitzero"`
	// Weight or probability of this manifold for prioritization purposes, between 0
	// and 1.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r ManifoldUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ManifoldUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManifoldUpdateParams) UnmarshalJSON(data []byte) error {
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
type ManifoldUpdateParamsDataMode string

const (
	ManifoldUpdateParamsDataModeReal      ManifoldUpdateParamsDataMode = "REAL"
	ManifoldUpdateParamsDataModeTest      ManifoldUpdateParamsDataMode = "TEST"
	ManifoldUpdateParamsDataModeSimulated ManifoldUpdateParamsDataMode = "SIMULATED"
	ManifoldUpdateParamsDataModeExercise  ManifoldUpdateParamsDataMode = "EXERCISE"
)

type ManifoldListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManifoldListParams]'s query parameters as `url.Values`.
func (r ManifoldListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManifoldCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManifoldCountParams]'s query parameters as `url.Values`.
func (r ManifoldCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManifoldNewBulkParams struct {
	Body []ManifoldNewBulkParamsBody
	paramObj
}

func (r ManifoldNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ManifoldNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// A manifold represents a set of possible/theoretical orbits for an object of
// interest based on a delta V and delta T.
//
// The properties ClassificationMarking, DataMode, IDObjectOfInterest, Source are
// required.
type ManifoldNewBulkParamsBody struct {
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
	// ID of the parent object of interest.
	IDObjectOfInterest string `json:"idObjectOfInterest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Applied delta V duration for this manifold's calculations in seconds.
	DeltaT param.Opt[float64] `json:"deltaT,omitzero"`
	// Applied delta V for this manifold's calculations, in km/sec.
	DeltaV param.Opt[float64] `json:"deltaV,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Status of the manifold and its associated ManifoldElsets (e.g. PENDING,
	// COMPLETE). PENDING status means element set generation is in progress and
	// COMPLETE indicates all ManifoldElsets have been generated.
	Status param.Opt[string] `json:"status,omitzero"`
	// Weight or probability of this manifold for prioritization purposes, between 0
	// and 1.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r ManifoldNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ManifoldNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManifoldNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ManifoldNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type ManifoldGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManifoldGetParams]'s query parameters as `url.Values`.
func (r ManifoldGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManifoldTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManifoldTupleParams]'s query parameters as `url.Values`.
func (r ManifoldTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
