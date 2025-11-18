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

// SortiePprService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSortiePprService] method instead.
type SortiePprService struct {
	Options []option.RequestOption
	History SortiePprHistoryService
}

// NewSortiePprService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSortiePprService(opts ...option.RequestOption) (r SortiePprService) {
	r = SortiePprService{}
	r.Options = opts
	r.History = NewSortiePprHistoryService(opts...)
	return
}

// Service operation to take a single sortieppr record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SortiePprService) New(ctx context.Context, body SortiePprNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/sortieppr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single sortieppr record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SortiePprService) Update(ctx context.Context, id string, body SortiePprUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sortieppr/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SortiePprService) List(ctx context.Context, query SortiePprListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SortiePprListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sortieppr"
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
func (r *SortiePprService) ListAutoPaging(ctx context.Context, query SortiePprListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SortiePprListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a sortieppr record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *SortiePprService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sortieppr/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SortiePprService) Count(ctx context.Context, query SortiePprCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sortieppr/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// SortiePPR records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *SortiePprService) NewBulk(ctx context.Context, body SortiePprNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/sortieppr/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single sortieppr record by its unique ID passed as a
// path parameter.
func (r *SortiePprService) Get(ctx context.Context, id string, query SortiePprGetParams, opts ...option.RequestOption) (res *shared.SortiePprFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sortieppr/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SortiePprService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SortiePprQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/sortieppr/queryhelp"
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
func (r *SortiePprService) Tuple(ctx context.Context, query SortiePprTupleParams, opts ...option.RequestOption) (res *[]shared.SortiePprFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/sortieppr/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take SortiePPR as a POST body and ingest into the database.
// This operation is intended to be used for automated feeds into UDL. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *SortiePprService) UnvalidatedPublish(ctx context.Context, body SortiePprUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-sortieppr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// SortiePPR is a regulatory requirement where operators must obtain permissions to
// full operational access to a runway, taxiway, or airport service.
type SortiePprListResponse struct {
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
	DataMode SortiePprListResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie string `json:"idSortie,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Time the prior permission required (PPR) valid window ends, in ISO 8601 UTC
	// format with millisecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// Identifier of the prior permission required (PPR) grantor.
	Grantor string `json:"grantor"`
	// The prior permission required (PPR) number issued by the airfield for a sortie.
	Number string `json:"number"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Remarks concerning the prior permission required (PPR) for a sortie.
	Remarks string `json:"remarks"`
	// The username of the prior permission required (PPR) requestor.
	Requestor string `json:"requestor"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the prior permission required (PPR) valid window begins, in ISO 8601 UTC
	// format with millisecond precision.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// Type of prior permission required (PPR) for a sortie (M - Military or C -
	// Civilian). Enum: [M, C].
	//
	// Any of "M", "C".
	Type SortiePprListResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSortie              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndTime               respjson.Field
		ExternalID            respjson.Field
		Grantor               respjson.Field
		Number                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Remarks               respjson.Field
		Requestor             respjson.Field
		SourceDl              respjson.Field
		StartTime             respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SortiePprListResponse) RawJSON() string { return r.JSON.raw }
func (r *SortiePprListResponse) UnmarshalJSON(data []byte) error {
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
type SortiePprListResponseDataMode string

const (
	SortiePprListResponseDataModeReal      SortiePprListResponseDataMode = "REAL"
	SortiePprListResponseDataModeTest      SortiePprListResponseDataMode = "TEST"
	SortiePprListResponseDataModeSimulated SortiePprListResponseDataMode = "SIMULATED"
	SortiePprListResponseDataModeExercise  SortiePprListResponseDataMode = "EXERCISE"
)

// Type of prior permission required (PPR) for a sortie (M - Military or C -
// Civilian). Enum: [M, C].
type SortiePprListResponseType string

const (
	SortiePprListResponseTypeM SortiePprListResponseType = "M"
	SortiePprListResponseTypeC SortiePprListResponseType = "C"
)

type SortiePprQueryhelpResponse struct {
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
func (r SortiePprQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SortiePprQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SortiePprNewParams struct {
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
	DataMode SortiePprNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie string `json:"idSortie,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the prior permission required (PPR) valid window ends, in ISO 8601 UTC
	// format with millisecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Identifier of the prior permission required (PPR) grantor.
	Grantor param.Opt[string] `json:"grantor,omitzero"`
	// The prior permission required (PPR) number issued by the airfield for a sortie.
	Number param.Opt[string] `json:"number,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Remarks concerning the prior permission required (PPR) for a sortie.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// The username of the prior permission required (PPR) requestor.
	Requestor param.Opt[string] `json:"requestor,omitzero"`
	// Time the prior permission required (PPR) valid window begins, in ISO 8601 UTC
	// format with millisecond precision.
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// Type of prior permission required (PPR) for a sortie (M - Military or C -
	// Civilian). Enum: [M, C].
	//
	// Any of "M", "C".
	Type SortiePprNewParamsType `json:"type,omitzero"`
	paramObj
}

func (r SortiePprNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SortiePprNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SortiePprNewParams) UnmarshalJSON(data []byte) error {
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
type SortiePprNewParamsDataMode string

const (
	SortiePprNewParamsDataModeReal      SortiePprNewParamsDataMode = "REAL"
	SortiePprNewParamsDataModeTest      SortiePprNewParamsDataMode = "TEST"
	SortiePprNewParamsDataModeSimulated SortiePprNewParamsDataMode = "SIMULATED"
	SortiePprNewParamsDataModeExercise  SortiePprNewParamsDataMode = "EXERCISE"
)

// Type of prior permission required (PPR) for a sortie (M - Military or C -
// Civilian). Enum: [M, C].
type SortiePprNewParamsType string

const (
	SortiePprNewParamsTypeM SortiePprNewParamsType = "M"
	SortiePprNewParamsTypeC SortiePprNewParamsType = "C"
)

type SortiePprUpdateParams struct {
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
	DataMode SortiePprUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie string `json:"idSortie,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the prior permission required (PPR) valid window ends, in ISO 8601 UTC
	// format with millisecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Identifier of the prior permission required (PPR) grantor.
	Grantor param.Opt[string] `json:"grantor,omitzero"`
	// The prior permission required (PPR) number issued by the airfield for a sortie.
	Number param.Opt[string] `json:"number,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Remarks concerning the prior permission required (PPR) for a sortie.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// The username of the prior permission required (PPR) requestor.
	Requestor param.Opt[string] `json:"requestor,omitzero"`
	// Time the prior permission required (PPR) valid window begins, in ISO 8601 UTC
	// format with millisecond precision.
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// Type of prior permission required (PPR) for a sortie (M - Military or C -
	// Civilian). Enum: [M, C].
	//
	// Any of "M", "C".
	Type SortiePprUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r SortiePprUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SortiePprUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SortiePprUpdateParams) UnmarshalJSON(data []byte) error {
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
type SortiePprUpdateParamsDataMode string

const (
	SortiePprUpdateParamsDataModeReal      SortiePprUpdateParamsDataMode = "REAL"
	SortiePprUpdateParamsDataModeTest      SortiePprUpdateParamsDataMode = "TEST"
	SortiePprUpdateParamsDataModeSimulated SortiePprUpdateParamsDataMode = "SIMULATED"
	SortiePprUpdateParamsDataModeExercise  SortiePprUpdateParamsDataMode = "EXERCISE"
)

// Type of prior permission required (PPR) for a sortie (M - Military or C -
// Civilian). Enum: [M, C].
type SortiePprUpdateParamsType string

const (
	SortiePprUpdateParamsTypeM SortiePprUpdateParamsType = "M"
	SortiePprUpdateParamsTypeC SortiePprUpdateParamsType = "C"
)

type SortiePprListParams struct {
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie    string           `query:"idSortie,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SortiePprListParams]'s query parameters as `url.Values`.
func (r SortiePprListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SortiePprCountParams struct {
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie    string           `query:"idSortie,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SortiePprCountParams]'s query parameters as `url.Values`.
func (r SortiePprCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SortiePprNewBulkParams struct {
	Body []SortiePprNewBulkParamsBody
	paramObj
}

func (r SortiePprNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SortiePprNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// SortiePPR is a regulatory requirement where operators must obtain permissions to
// full operational access to a runway, taxiway, or airport service.
//
// The properties ClassificationMarking, DataMode, IDSortie, Source are required.
type SortiePprNewBulkParamsBody struct {
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
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie string `json:"idSortie,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the prior permission required (PPR) valid window ends, in ISO 8601 UTC
	// format with millisecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Identifier of the prior permission required (PPR) grantor.
	Grantor param.Opt[string] `json:"grantor,omitzero"`
	// The prior permission required (PPR) number issued by the airfield for a sortie.
	Number param.Opt[string] `json:"number,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Remarks concerning the prior permission required (PPR) for a sortie.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// The username of the prior permission required (PPR) requestor.
	Requestor param.Opt[string] `json:"requestor,omitzero"`
	// Time the prior permission required (PPR) valid window begins, in ISO 8601 UTC
	// format with millisecond precision.
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// Type of prior permission required (PPR) for a sortie (M - Military or C -
	// Civilian). Enum: [M, C].
	//
	// Any of "M", "C".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SortiePprNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SortiePprNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SortiePprNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SortiePprNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SortiePprNewBulkParamsBody](
		"type", "M", "C",
	)
}

type SortiePprGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SortiePprGetParams]'s query parameters as `url.Values`.
func (r SortiePprGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SortiePprTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie    string           `query:"idSortie,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SortiePprTupleParams]'s query parameters as `url.Values`.
func (r SortiePprTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SortiePprUnvalidatedPublishParams struct {
	Body []SortiePprUnvalidatedPublishParamsBody
	paramObj
}

func (r SortiePprUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SortiePprUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// SortiePPR is a regulatory requirement where operators must obtain permissions to
// full operational access to a runway, taxiway, or airport service.
//
// The properties ClassificationMarking, DataMode, IDSortie, Source are required.
type SortiePprUnvalidatedPublishParamsBody struct {
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
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie string `json:"idSortie,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the prior permission required (PPR) valid window ends, in ISO 8601 UTC
	// format with millisecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Identifier of the prior permission required (PPR) grantor.
	Grantor param.Opt[string] `json:"grantor,omitzero"`
	// The prior permission required (PPR) number issued by the airfield for a sortie.
	Number param.Opt[string] `json:"number,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Remarks concerning the prior permission required (PPR) for a sortie.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// The username of the prior permission required (PPR) requestor.
	Requestor param.Opt[string] `json:"requestor,omitzero"`
	// Time the prior permission required (PPR) valid window begins, in ISO 8601 UTC
	// format with millisecond precision.
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// Type of prior permission required (PPR) for a sortie (M - Military or C -
	// Civilian). Enum: [M, C].
	//
	// Any of "M", "C".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SortiePprUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SortiePprUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SortiePprUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SortiePprUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SortiePprUnvalidatedPublishParamsBody](
		"type", "M", "C",
	)
}
