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

// SortiepprService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSortiepprService] method instead.
type SortiepprService struct {
	Options []option.RequestOption
	History SortiepprHistoryService
}

// NewSortiepprService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSortiepprService(opts ...option.RequestOption) (r SortiepprService) {
	r = SortiepprService{}
	r.Options = opts
	r.History = NewSortiepprHistoryService(opts...)
	return
}

// Service operation to take a single sortieppr record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SortiepprService) New(ctx context.Context, body SortiepprNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sortieppr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single sortieppr record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SortiepprService) Update(ctx context.Context, id string, body SortiepprUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
func (r *SortiepprService) List(ctx context.Context, query SortiepprListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SortiepprListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *SortiepprService) ListAutoPaging(ctx context.Context, query SortiepprListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SortiepprListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a sortieppr record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *SortiepprService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
func (r *SortiepprService) Count(ctx context.Context, query SortiepprCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *SortiepprService) NewBulk(ctx context.Context, body SortiepprNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sortieppr/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single sortieppr record by its unique ID passed as a
// path parameter.
func (r *SortiepprService) Get(ctx context.Context, id string, query SortiepprGetParams, opts ...option.RequestOption) (res *SortiePprFull, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *SortiepprService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sortieppr/queryhelp"
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
func (r *SortiepprService) Tuple(ctx context.Context, query SortiepprTupleParams, opts ...option.RequestOption) (res *[]SortiePprFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sortieppr/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take SortiePPR as a POST body and ingest into the database.
// This operation is intended to be used for automated feeds into UDL. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *SortiepprService) UnvalidatedPublish(ctx context.Context, body SortiepprUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-sortieppr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// SortiePPR is a regulatory requirement where operators must obtain permissions to
// full operational access to a runway, taxiway, or airport service.
type SortiepprListResponse struct {
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
	DataMode SortiepprListResponseDataMode `json:"dataMode,required"`
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
	Type SortiepprListResponseType `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDSortie              resp.Field
		Source                resp.Field
		ID                    resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EndTime               resp.Field
		ExternalID            resp.Field
		Grantor               resp.Field
		Number                resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Remarks               resp.Field
		Requestor             resp.Field
		SourceDl              resp.Field
		StartTime             resp.Field
		Type                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SortiepprListResponse) RawJSON() string { return r.JSON.raw }
func (r *SortiepprListResponse) UnmarshalJSON(data []byte) error {
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
type SortiepprListResponseDataMode string

const (
	SortiepprListResponseDataModeReal      SortiepprListResponseDataMode = "REAL"
	SortiepprListResponseDataModeTest      SortiepprListResponseDataMode = "TEST"
	SortiepprListResponseDataModeSimulated SortiepprListResponseDataMode = "SIMULATED"
	SortiepprListResponseDataModeExercise  SortiepprListResponseDataMode = "EXERCISE"
)

// Type of prior permission required (PPR) for a sortie (M - Military or C -
// Civilian). Enum: [M, C].
type SortiepprListResponseType string

const (
	SortiepprListResponseTypeM SortiepprListResponseType = "M"
	SortiepprListResponseTypeC SortiepprListResponseType = "C"
)

type SortiepprNewParams struct {
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
	DataMode SortiepprNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	Type SortiepprNewParamsType `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SortiepprNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SortiepprNewParams
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
type SortiepprNewParamsDataMode string

const (
	SortiepprNewParamsDataModeReal      SortiepprNewParamsDataMode = "REAL"
	SortiepprNewParamsDataModeTest      SortiepprNewParamsDataMode = "TEST"
	SortiepprNewParamsDataModeSimulated SortiepprNewParamsDataMode = "SIMULATED"
	SortiepprNewParamsDataModeExercise  SortiepprNewParamsDataMode = "EXERCISE"
)

// Type of prior permission required (PPR) for a sortie (M - Military or C -
// Civilian). Enum: [M, C].
type SortiepprNewParamsType string

const (
	SortiepprNewParamsTypeM SortiepprNewParamsType = "M"
	SortiepprNewParamsTypeC SortiepprNewParamsType = "C"
)

type SortiepprUpdateParams struct {
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
	DataMode SortiepprUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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
	Type SortiepprUpdateParamsType `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SortiepprUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SortiepprUpdateParams
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
type SortiepprUpdateParamsDataMode string

const (
	SortiepprUpdateParamsDataModeReal      SortiepprUpdateParamsDataMode = "REAL"
	SortiepprUpdateParamsDataModeTest      SortiepprUpdateParamsDataMode = "TEST"
	SortiepprUpdateParamsDataModeSimulated SortiepprUpdateParamsDataMode = "SIMULATED"
	SortiepprUpdateParamsDataModeExercise  SortiepprUpdateParamsDataMode = "EXERCISE"
)

// Type of prior permission required (PPR) for a sortie (M - Military or C -
// Civilian). Enum: [M, C].
type SortiepprUpdateParamsType string

const (
	SortiepprUpdateParamsTypeM SortiepprUpdateParamsType = "M"
	SortiepprUpdateParamsTypeC SortiepprUpdateParamsType = "C"
)

type SortiepprListParams struct {
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie    string           `query:"idSortie,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SortiepprListParams]'s query parameters as `url.Values`.
func (r SortiepprListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SortiepprCountParams struct {
	// Unique identifier of the Aircraft Sortie associated with this prior permission
	// required (PPR) record.
	IDSortie    string           `query:"idSortie,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SortiepprCountParams]'s query parameters as `url.Values`.
func (r SortiepprCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SortiepprNewBulkParams struct {
	Body []SortiepprNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SortiepprNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// SortiePPR is a regulatory requirement where operators must obtain permissions to
// full operational access to a runway, taxiway, or airport service.
//
// The properties ClassificationMarking, DataMode, IDSortie, Source are required.
type SortiepprNewBulkParamsBody struct {
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
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Remarks concerning the prior permission required (PPR) for a sortie.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// The username of the prior permission required (PPR) requestor.
	Requestor param.Opt[string] `json:"requestor,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprNewBulkParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r SortiepprNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SortiepprNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[SortiepprNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SortiepprNewBulkParamsBody](
		"Type", false, "M", "C",
	)
}

type SortiepprGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SortiepprGetParams]'s query parameters as `url.Values`.
func (r SortiepprGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SortiepprTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SortiepprTupleParams]'s query parameters as `url.Values`.
func (r SortiepprTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SortiepprUnvalidatedPublishParams struct {
	Body []SortiepprUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r SortiepprUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// SortiePPR is a regulatory requirement where operators must obtain permissions to
// full operational access to a runway, taxiway, or airport service.
//
// The properties ClassificationMarking, DataMode, IDSortie, Source are required.
type SortiepprUnvalidatedPublishParamsBody struct {
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
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Remarks concerning the prior permission required (PPR) for a sortie.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// The username of the prior permission required (PPR) requestor.
	Requestor param.Opt[string] `json:"requestor,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SortiepprUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r SortiepprUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SortiepprUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[SortiepprUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SortiepprUnvalidatedPublishParamsBody](
		"Type", false, "M", "C",
	)
}
