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

// SurfaceObstructionService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSurfaceObstructionService] method instead.
type SurfaceObstructionService struct {
	Options []option.RequestOption
}

// NewSurfaceObstructionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSurfaceObstructionService(opts ...option.RequestOption) (r SurfaceObstructionService) {
	r = SurfaceObstructionService{}
	r.Options = opts
	return
}

// Service operation to take a single surfaceobstruction record as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SurfaceObstructionService) New(ctx context.Context, body SurfaceObstructionNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/surfaceobstruction"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single surfaceobstruction record. A specific role
// is required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SurfaceObstructionService) Update(ctx context.Context, id string, body SurfaceObstructionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/surfaceobstruction/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SurfaceObstructionService) List(ctx context.Context, query SurfaceObstructionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SurfaceObstructionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/surfaceobstruction"
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
func (r *SurfaceObstructionService) ListAutoPaging(ctx context.Context, query SurfaceObstructionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SurfaceObstructionListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a surfaceobstruction record specified by the passed
// ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SurfaceObstructionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/surfaceobstruction/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SurfaceObstructionService) Count(ctx context.Context, query SurfaceObstructionCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/surfaceobstruction/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single surfaceobstruction record by its unique ID
// passed as a path parameter.
func (r *SurfaceObstructionService) Get(ctx context.Context, id string, query SurfaceObstructionGetParams, opts ...option.RequestOption) (res *SurfaceObstructionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/surfaceobstruction/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SurfaceObstructionService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SurfaceObstructionQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/surfaceobstruction/queryhelp"
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
func (r *SurfaceObstructionService) Tuple(ctx context.Context, query SurfaceObstructionTupleParams, opts ...option.RequestOption) (res *[]SurfaceObstructionTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/surfaceobstruction/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple surfaceobstruction records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SurfaceObstructionService) UnvalidatedPublish(ctx context.Context, body SurfaceObstructionUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-surfaceobstruction"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type SurfaceObstructionListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceObstructionListResponseDataMode `json:"dataMode,required"`
	// The unique identifier of the associated surface record. This field is required
	// when posting, updating, or deleting a SurfaceObstruction record.
	IDSurface string `json:"idSurface,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an advisory for usage.
	AdvisoryRequired []string `json:"advisoryRequired"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an approval for usage.
	ApprovalRequired []string `json:"approvalRequired"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The distance from the surface center line to this obstruction, in feet.
	DistanceFromCenterLine float64 `json:"distanceFromCenterLine"`
	// The distance from the surface edge to this obstruction, in feet.
	DistanceFromEdge float64 `json:"distanceFromEdge"`
	// The distance from the surface threshold to this obstruction, in feet.
	DistanceFromThreshold float64 `json:"distanceFromThreshold"`
	// The unique identifier of the associated NavigationalObstruction record.
	IDNavigationalObstruction string `json:"idNavigationalObstruction"`
	// Description of this surface obstruction.
	ObstructionDesc string `json:"obstructionDesc"`
	// The height above ground level of the surface obstruction, in feet.
	ObstructionHeight float64 `json:"obstructionHeight"`
	// A code that indicates which side of the surface end is affected by this
	// obstruction.
	ObstructionSideCode string `json:"obstructionSideCode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking     respjson.Field
		DataMode                  respjson.Field
		IDSurface                 respjson.Field
		Source                    respjson.Field
		ID                        respjson.Field
		AdvisoryRequired          respjson.Field
		ApprovalRequired          respjson.Field
		CreatedAt                 respjson.Field
		CreatedBy                 respjson.Field
		DistanceFromCenterLine    respjson.Field
		DistanceFromEdge          respjson.Field
		DistanceFromThreshold     respjson.Field
		IDNavigationalObstruction respjson.Field
		ObstructionDesc           respjson.Field
		ObstructionHeight         respjson.Field
		ObstructionSideCode       respjson.Field
		Origin                    respjson.Field
		OrigNetwork               respjson.Field
		SourceDl                  respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SurfaceObstructionListResponse) RawJSON() string { return r.JSON.raw }
func (r *SurfaceObstructionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceObstructionListResponseDataMode string

const (
	SurfaceObstructionListResponseDataModeReal      SurfaceObstructionListResponseDataMode = "REAL"
	SurfaceObstructionListResponseDataModeTest      SurfaceObstructionListResponseDataMode = "TEST"
	SurfaceObstructionListResponseDataModeSimulated SurfaceObstructionListResponseDataMode = "SIMULATED"
	SurfaceObstructionListResponseDataModeExercise  SurfaceObstructionListResponseDataMode = "EXERCISE"
)

type SurfaceObstructionGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceObstructionGetResponseDataMode `json:"dataMode,required"`
	// The unique identifier of the associated surface record. This field is required
	// when posting, updating, or deleting a SurfaceObstruction record.
	IDSurface string `json:"idSurface,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an advisory for usage.
	AdvisoryRequired []string `json:"advisoryRequired"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an approval for usage.
	ApprovalRequired []string `json:"approvalRequired"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The distance from the surface center line to this obstruction, in feet.
	DistanceFromCenterLine float64 `json:"distanceFromCenterLine"`
	// The distance from the surface edge to this obstruction, in feet.
	DistanceFromEdge float64 `json:"distanceFromEdge"`
	// The distance from the surface threshold to this obstruction, in feet.
	DistanceFromThreshold float64 `json:"distanceFromThreshold"`
	// The unique identifier of the associated NavigationalObstruction record.
	IDNavigationalObstruction string `json:"idNavigationalObstruction"`
	// Description of this surface obstruction.
	ObstructionDesc string `json:"obstructionDesc"`
	// The height above ground level of the surface obstruction, in feet.
	ObstructionHeight float64 `json:"obstructionHeight"`
	// A code that indicates which side of the surface end is affected by this
	// obstruction.
	ObstructionSideCode string `json:"obstructionSideCode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking     respjson.Field
		DataMode                  respjson.Field
		IDSurface                 respjson.Field
		Source                    respjson.Field
		ID                        respjson.Field
		AdvisoryRequired          respjson.Field
		ApprovalRequired          respjson.Field
		CreatedAt                 respjson.Field
		CreatedBy                 respjson.Field
		DistanceFromCenterLine    respjson.Field
		DistanceFromEdge          respjson.Field
		DistanceFromThreshold     respjson.Field
		IDNavigationalObstruction respjson.Field
		ObstructionDesc           respjson.Field
		ObstructionHeight         respjson.Field
		ObstructionSideCode       respjson.Field
		Origin                    respjson.Field
		OrigNetwork               respjson.Field
		SourceDl                  respjson.Field
		UpdatedAt                 respjson.Field
		UpdatedBy                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SurfaceObstructionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SurfaceObstructionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceObstructionGetResponseDataMode string

const (
	SurfaceObstructionGetResponseDataModeReal      SurfaceObstructionGetResponseDataMode = "REAL"
	SurfaceObstructionGetResponseDataModeTest      SurfaceObstructionGetResponseDataMode = "TEST"
	SurfaceObstructionGetResponseDataModeSimulated SurfaceObstructionGetResponseDataMode = "SIMULATED"
	SurfaceObstructionGetResponseDataModeExercise  SurfaceObstructionGetResponseDataMode = "EXERCISE"
)

type SurfaceObstructionQueryhelpResponse struct {
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
func (r SurfaceObstructionQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SurfaceObstructionQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SurfaceObstructionTupleResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceObstructionTupleResponseDataMode `json:"dataMode,required"`
	// The unique identifier of the associated surface record. This field is required
	// when posting, updating, or deleting a SurfaceObstruction record.
	IDSurface string `json:"idSurface,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an advisory for usage.
	AdvisoryRequired []string `json:"advisoryRequired"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an approval for usage.
	ApprovalRequired []string `json:"approvalRequired"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The distance from the surface center line to this obstruction, in feet.
	DistanceFromCenterLine float64 `json:"distanceFromCenterLine"`
	// The distance from the surface edge to this obstruction, in feet.
	DistanceFromEdge float64 `json:"distanceFromEdge"`
	// The distance from the surface threshold to this obstruction, in feet.
	DistanceFromThreshold float64 `json:"distanceFromThreshold"`
	// The unique identifier of the associated NavigationalObstruction record.
	IDNavigationalObstruction string `json:"idNavigationalObstruction"`
	// Description of this surface obstruction.
	ObstructionDesc string `json:"obstructionDesc"`
	// The height above ground level of the surface obstruction, in feet.
	ObstructionHeight float64 `json:"obstructionHeight"`
	// A code that indicates which side of the surface end is affected by this
	// obstruction.
	ObstructionSideCode string `json:"obstructionSideCode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking     respjson.Field
		DataMode                  respjson.Field
		IDSurface                 respjson.Field
		Source                    respjson.Field
		ID                        respjson.Field
		AdvisoryRequired          respjson.Field
		ApprovalRequired          respjson.Field
		CreatedAt                 respjson.Field
		CreatedBy                 respjson.Field
		DistanceFromCenterLine    respjson.Field
		DistanceFromEdge          respjson.Field
		DistanceFromThreshold     respjson.Field
		IDNavigationalObstruction respjson.Field
		ObstructionDesc           respjson.Field
		ObstructionHeight         respjson.Field
		ObstructionSideCode       respjson.Field
		Origin                    respjson.Field
		OrigNetwork               respjson.Field
		SourceDl                  respjson.Field
		UpdatedAt                 respjson.Field
		UpdatedBy                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SurfaceObstructionTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SurfaceObstructionTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceObstructionTupleResponseDataMode string

const (
	SurfaceObstructionTupleResponseDataModeReal      SurfaceObstructionTupleResponseDataMode = "REAL"
	SurfaceObstructionTupleResponseDataModeTest      SurfaceObstructionTupleResponseDataMode = "TEST"
	SurfaceObstructionTupleResponseDataModeSimulated SurfaceObstructionTupleResponseDataMode = "SIMULATED"
	SurfaceObstructionTupleResponseDataModeExercise  SurfaceObstructionTupleResponseDataMode = "EXERCISE"
)

type SurfaceObstructionNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceObstructionNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The unique identifier of the associated surface record. This field is required
	// when posting, updating, or deleting a SurfaceObstruction record.
	IDSurface string `json:"idSurface,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The distance from the surface center line to this obstruction, in feet.
	DistanceFromCenterLine param.Opt[float64] `json:"distanceFromCenterLine,omitzero"`
	// The distance from the surface edge to this obstruction, in feet.
	DistanceFromEdge param.Opt[float64] `json:"distanceFromEdge,omitzero"`
	// The distance from the surface threshold to this obstruction, in feet.
	DistanceFromThreshold param.Opt[float64] `json:"distanceFromThreshold,omitzero"`
	// The unique identifier of the associated NavigationalObstruction record.
	IDNavigationalObstruction param.Opt[string] `json:"idNavigationalObstruction,omitzero"`
	// Description of this surface obstruction.
	ObstructionDesc param.Opt[string] `json:"obstructionDesc,omitzero"`
	// The height above ground level of the surface obstruction, in feet.
	ObstructionHeight param.Opt[float64] `json:"obstructionHeight,omitzero"`
	// A code that indicates which side of the surface end is affected by this
	// obstruction.
	ObstructionSideCode param.Opt[string] `json:"obstructionSideCode,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an advisory for usage.
	AdvisoryRequired []string `json:"advisoryRequired,omitzero"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an approval for usage.
	ApprovalRequired []string `json:"approvalRequired,omitzero"`
	paramObj
}

func (r SurfaceObstructionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SurfaceObstructionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SurfaceObstructionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceObstructionNewParamsDataMode string

const (
	SurfaceObstructionNewParamsDataModeReal      SurfaceObstructionNewParamsDataMode = "REAL"
	SurfaceObstructionNewParamsDataModeTest      SurfaceObstructionNewParamsDataMode = "TEST"
	SurfaceObstructionNewParamsDataModeSimulated SurfaceObstructionNewParamsDataMode = "SIMULATED"
	SurfaceObstructionNewParamsDataModeExercise  SurfaceObstructionNewParamsDataMode = "EXERCISE"
)

type SurfaceObstructionUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceObstructionUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The unique identifier of the associated surface record. This field is required
	// when posting, updating, or deleting a SurfaceObstruction record.
	IDSurface string `json:"idSurface,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The distance from the surface center line to this obstruction, in feet.
	DistanceFromCenterLine param.Opt[float64] `json:"distanceFromCenterLine,omitzero"`
	// The distance from the surface edge to this obstruction, in feet.
	DistanceFromEdge param.Opt[float64] `json:"distanceFromEdge,omitzero"`
	// The distance from the surface threshold to this obstruction, in feet.
	DistanceFromThreshold param.Opt[float64] `json:"distanceFromThreshold,omitzero"`
	// The unique identifier of the associated NavigationalObstruction record.
	IDNavigationalObstruction param.Opt[string] `json:"idNavigationalObstruction,omitzero"`
	// Description of this surface obstruction.
	ObstructionDesc param.Opt[string] `json:"obstructionDesc,omitzero"`
	// The height above ground level of the surface obstruction, in feet.
	ObstructionHeight param.Opt[float64] `json:"obstructionHeight,omitzero"`
	// A code that indicates which side of the surface end is affected by this
	// obstruction.
	ObstructionSideCode param.Opt[string] `json:"obstructionSideCode,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an advisory for usage.
	AdvisoryRequired []string `json:"advisoryRequired,omitzero"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an approval for usage.
	ApprovalRequired []string `json:"approvalRequired,omitzero"`
	paramObj
}

func (r SurfaceObstructionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SurfaceObstructionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SurfaceObstructionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceObstructionUpdateParamsDataMode string

const (
	SurfaceObstructionUpdateParamsDataModeReal      SurfaceObstructionUpdateParamsDataMode = "REAL"
	SurfaceObstructionUpdateParamsDataModeTest      SurfaceObstructionUpdateParamsDataMode = "TEST"
	SurfaceObstructionUpdateParamsDataModeSimulated SurfaceObstructionUpdateParamsDataMode = "SIMULATED"
	SurfaceObstructionUpdateParamsDataModeExercise  SurfaceObstructionUpdateParamsDataMode = "EXERCISE"
)

type SurfaceObstructionListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SurfaceObstructionListParams]'s query parameters as
// `url.Values`.
func (r SurfaceObstructionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SurfaceObstructionCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SurfaceObstructionCountParams]'s query parameters as
// `url.Values`.
func (r SurfaceObstructionCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SurfaceObstructionGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SurfaceObstructionGetParams]'s query parameters as
// `url.Values`.
func (r SurfaceObstructionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SurfaceObstructionTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SurfaceObstructionTupleParams]'s query parameters as
// `url.Values`.
func (r SurfaceObstructionTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SurfaceObstructionUnvalidatedPublishParams struct {
	Body []SurfaceObstructionUnvalidatedPublishParamsBody
	paramObj
}

func (r SurfaceObstructionUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SurfaceObstructionUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The properties ClassificationMarking, DataMode, IDSurface, Source are required.
type SurfaceObstructionUnvalidatedPublishParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,omitzero,required"`
	// The unique identifier of the associated surface record. This field is required
	// when posting, updating, or deleting a SurfaceObstruction record.
	IDSurface string `json:"idSurface,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The distance from the surface center line to this obstruction, in feet.
	DistanceFromCenterLine param.Opt[float64] `json:"distanceFromCenterLine,omitzero"`
	// The distance from the surface edge to this obstruction, in feet.
	DistanceFromEdge param.Opt[float64] `json:"distanceFromEdge,omitzero"`
	// The distance from the surface threshold to this obstruction, in feet.
	DistanceFromThreshold param.Opt[float64] `json:"distanceFromThreshold,omitzero"`
	// The unique identifier of the associated NavigationalObstruction record.
	IDNavigationalObstruction param.Opt[string] `json:"idNavigationalObstruction,omitzero"`
	// Description of this surface obstruction.
	ObstructionDesc param.Opt[string] `json:"obstructionDesc,omitzero"`
	// The height above ground level of the surface obstruction, in feet.
	ObstructionHeight param.Opt[float64] `json:"obstructionHeight,omitzero"`
	// A code that indicates which side of the surface end is affected by this
	// obstruction.
	ObstructionSideCode param.Opt[string] `json:"obstructionSideCode,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an advisory for usage.
	AdvisoryRequired []string `json:"advisoryRequired,omitzero"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an approval for usage.
	ApprovalRequired []string `json:"approvalRequired,omitzero"`
	paramObj
}

func (r SurfaceObstructionUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SurfaceObstructionUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SurfaceObstructionUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SurfaceObstructionUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
