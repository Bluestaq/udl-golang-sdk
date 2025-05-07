// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// AircraftStatusRemarkService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAircraftStatusRemarkService] method instead.
type AircraftStatusRemarkService struct {
	Options []option.RequestOption
}

// NewAircraftStatusRemarkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAircraftStatusRemarkService(opts ...option.RequestOption) (r AircraftStatusRemarkService) {
	r = AircraftStatusRemarkService{}
	r.Options = opts
	return
}

// Service operation to take a single Aircraft Status Remark record as a POST body
// and ingest into the database. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *AircraftStatusRemarkService) New(ctx context.Context, body AircraftStatusRemarkNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aircraftstatusremark"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Aircraft Status Remark record by its unique ID
// passed as a path parameter.
func (r *AircraftStatusRemarkService) Get(ctx context.Context, id string, query AircraftStatusRemarkGetParams, opts ...option.RequestOption) (res *AircraftstatusremarkFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraftstatusremark/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single Aircraft Status Remark record. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *AircraftStatusRemarkService) Update(ctx context.Context, id string, body AircraftStatusRemarkUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraftstatusremark/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AircraftStatusRemarkService) List(ctx context.Context, query AircraftStatusRemarkListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AircraftstatusremarkAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/aircraftstatusremark"
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
func (r *AircraftStatusRemarkService) ListAutoPaging(ctx context.Context, query AircraftStatusRemarkListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AircraftstatusremarkAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a single Aircraft Status Remark record specified by
// the passed ID path parameter. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *AircraftStatusRemarkService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraftstatusremark/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AircraftStatusRemarkService) Count(ctx context.Context, query AircraftStatusRemarkCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/aircraftstatusremark/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AircraftStatusRemarkService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aircraftstatusremark/queryhelp"
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
func (r *AircraftStatusRemarkService) Tuple(ctx context.Context, query AircraftStatusRemarkTupleParams, opts ...option.RequestOption) (res *[]AircraftstatusremarkFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/aircraftstatusremark/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Properties and characteristics of a remark that is associated with an aircraft
// status.
type AircraftstatusremarkAbridged struct {
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
	DataMode AircraftstatusremarkAbridgedDataMode `json:"dataMode,required"`
	// The ID of the Aircraft Status to which this remark applies.
	IDAircraftStatus string `json:"idAircraftStatus,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Unique identifier of the Aircraft Status Remark record from the originating
	// system.
	AltRmkID string `json:"altRmkId"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Time the remark was last updated in the originating system in ISO 8601 UTC
	// format with millisecond precision.
	LastUpdatedAt time.Time `json:"lastUpdatedAt" format:"date-time"`
	// The name or ID of the external user that updated this remark in the originating
	// system.
	LastUpdatedBy string `json:"lastUpdatedBy"`
	// The name of the remark.
	Name string `json:"name"`
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
	// Time the remark was created in the originating system in ISO 8601 UTC format
	// with millisecond precision.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDAircraftStatus      respjson.Field
		Source                respjson.Field
		Text                  respjson.Field
		ID                    respjson.Field
		AltRmkID              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		LastUpdatedAt         respjson.Field
		LastUpdatedBy         respjson.Field
		Name                  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SourceDl              respjson.Field
		Timestamp             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AircraftstatusremarkAbridged) RawJSON() string { return r.JSON.raw }
func (r *AircraftstatusremarkAbridged) UnmarshalJSON(data []byte) error {
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
type AircraftstatusremarkAbridgedDataMode string

const (
	AircraftstatusremarkAbridgedDataModeReal      AircraftstatusremarkAbridgedDataMode = "REAL"
	AircraftstatusremarkAbridgedDataModeTest      AircraftstatusremarkAbridgedDataMode = "TEST"
	AircraftstatusremarkAbridgedDataModeSimulated AircraftstatusremarkAbridgedDataMode = "SIMULATED"
	AircraftstatusremarkAbridgedDataModeExercise  AircraftstatusremarkAbridgedDataMode = "EXERCISE"
)

// Properties and characteristics of a remark that is associated with an aircraft
// status.
type AircraftstatusremarkFull struct {
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
	DataMode AircraftstatusremarkFullDataMode `json:"dataMode,required"`
	// The ID of the Aircraft Status to which this remark applies.
	IDAircraftStatus string `json:"idAircraftStatus,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Unique identifier of the Aircraft Status Remark record from the originating
	// system.
	AltRmkID string `json:"altRmkId"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Time the remark was last updated in the originating system in ISO 8601 UTC
	// format with millisecond precision.
	LastUpdatedAt time.Time `json:"lastUpdatedAt" format:"date-time"`
	// The name or ID of the external user that updated this remark in the originating
	// system.
	LastUpdatedBy string `json:"lastUpdatedBy"`
	// The name of the remark.
	Name string `json:"name"`
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
	// Time the remark was created in the originating system in ISO 8601 UTC format
	// with millisecond precision.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDAircraftStatus      respjson.Field
		Source                respjson.Field
		Text                  respjson.Field
		ID                    respjson.Field
		AltRmkID              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		LastUpdatedAt         respjson.Field
		LastUpdatedBy         respjson.Field
		Name                  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SourceDl              respjson.Field
		Timestamp             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AircraftstatusremarkFull) RawJSON() string { return r.JSON.raw }
func (r *AircraftstatusremarkFull) UnmarshalJSON(data []byte) error {
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
type AircraftstatusremarkFullDataMode string

const (
	AircraftstatusremarkFullDataModeReal      AircraftstatusremarkFullDataMode = "REAL"
	AircraftstatusremarkFullDataModeTest      AircraftstatusremarkFullDataMode = "TEST"
	AircraftstatusremarkFullDataModeSimulated AircraftstatusremarkFullDataMode = "SIMULATED"
	AircraftstatusremarkFullDataModeExercise  AircraftstatusremarkFullDataMode = "EXERCISE"
)

type AircraftStatusRemarkNewParams struct {
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
	DataMode AircraftStatusRemarkNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The ID of the Aircraft Status to which this remark applies.
	IDAircraftStatus string `json:"idAircraftStatus,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Unique identifier of the Aircraft Status Remark record from the originating
	// system.
	AltRmkID param.Opt[string] `json:"altRmkId,omitzero"`
	// Time the remark was last updated in the originating system in ISO 8601 UTC
	// format with millisecond precision.
	LastUpdatedAt param.Opt[time.Time] `json:"lastUpdatedAt,omitzero" format:"date-time"`
	// The name or ID of the external user that updated this remark in the originating
	// system.
	LastUpdatedBy param.Opt[string] `json:"lastUpdatedBy,omitzero"`
	// The name of the remark.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Time the remark was created in the originating system in ISO 8601 UTC format
	// with millisecond precision.
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	paramObj
}

func (r AircraftStatusRemarkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AircraftStatusRemarkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AircraftStatusRemarkNewParams) UnmarshalJSON(data []byte) error {
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
type AircraftStatusRemarkNewParamsDataMode string

const (
	AircraftStatusRemarkNewParamsDataModeReal      AircraftStatusRemarkNewParamsDataMode = "REAL"
	AircraftStatusRemarkNewParamsDataModeTest      AircraftStatusRemarkNewParamsDataMode = "TEST"
	AircraftStatusRemarkNewParamsDataModeSimulated AircraftStatusRemarkNewParamsDataMode = "SIMULATED"
	AircraftStatusRemarkNewParamsDataModeExercise  AircraftStatusRemarkNewParamsDataMode = "EXERCISE"
)

type AircraftStatusRemarkGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftStatusRemarkGetParams]'s query parameters as
// `url.Values`.
func (r AircraftStatusRemarkGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftStatusRemarkUpdateParams struct {
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
	DataMode AircraftStatusRemarkUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The ID of the Aircraft Status to which this remark applies.
	IDAircraftStatus string `json:"idAircraftStatus,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Unique identifier of the Aircraft Status Remark record from the originating
	// system.
	AltRmkID param.Opt[string] `json:"altRmkId,omitzero"`
	// Time the remark was last updated in the originating system in ISO 8601 UTC
	// format with millisecond precision.
	LastUpdatedAt param.Opt[time.Time] `json:"lastUpdatedAt,omitzero" format:"date-time"`
	// The name or ID of the external user that updated this remark in the originating
	// system.
	LastUpdatedBy param.Opt[string] `json:"lastUpdatedBy,omitzero"`
	// The name of the remark.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Time the remark was created in the originating system in ISO 8601 UTC format
	// with millisecond precision.
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	paramObj
}

func (r AircraftStatusRemarkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AircraftStatusRemarkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AircraftStatusRemarkUpdateParams) UnmarshalJSON(data []byte) error {
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
type AircraftStatusRemarkUpdateParamsDataMode string

const (
	AircraftStatusRemarkUpdateParamsDataModeReal      AircraftStatusRemarkUpdateParamsDataMode = "REAL"
	AircraftStatusRemarkUpdateParamsDataModeTest      AircraftStatusRemarkUpdateParamsDataMode = "TEST"
	AircraftStatusRemarkUpdateParamsDataModeSimulated AircraftStatusRemarkUpdateParamsDataMode = "SIMULATED"
	AircraftStatusRemarkUpdateParamsDataModeExercise  AircraftStatusRemarkUpdateParamsDataMode = "EXERCISE"
)

type AircraftStatusRemarkListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftStatusRemarkListParams]'s query parameters as
// `url.Values`.
func (r AircraftStatusRemarkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftStatusRemarkCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftStatusRemarkCountParams]'s query parameters as
// `url.Values`.
func (r AircraftStatusRemarkCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftStatusRemarkTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftStatusRemarkTupleParams]'s query parameters as
// `url.Values`.
func (r AircraftStatusRemarkTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
