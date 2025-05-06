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
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// LaunchEventService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaunchEventService] method instead.
type LaunchEventService struct {
	Options []option.RequestOption
	History LaunchEventHistoryService
}

// NewLaunchEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaunchEventService(opts ...option.RequestOption) (r LaunchEventService) {
	r = LaunchEventService{}
	r.Options = opts
	r.History = NewLaunchEventHistoryService(opts...)
	return
}

// Service operation to take a single LaunchEvent as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LaunchEventService) New(ctx context.Context, body LaunchEventNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchevent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaunchEventService) List(ctx context.Context, query LaunchEventListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaunchEventListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/launchevent"
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
func (r *LaunchEventService) ListAutoPaging(ctx context.Context, query LaunchEventListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaunchEventListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LaunchEventService) Count(ctx context.Context, query LaunchEventCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/launchevent/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// LaunchEvent as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *LaunchEventService) NewBulk(ctx context.Context, body LaunchEventNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchevent/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single LaunchEvent record by its unique ID passed as
// a path parameter.
func (r *LaunchEventService) Get(ctx context.Context, id string, query LaunchEventGetParams, opts ...option.RequestOption) (res *LaunchEventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchevent/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *LaunchEventService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchevent/queryhelp"
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
func (r *LaunchEventService) Tuple(ctx context.Context, query LaunchEventTupleParams, opts ...option.RequestOption) (res *[]LaunchEventTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/launchevent/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take LaunchEvent entries as a POST body and ingest into the
// database. This operation is intended to be used for automated feeds into UDL. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *LaunchEventService) UnvalidatedPublish(ctx context.Context, body LaunchEventUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-launchevent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Information on known launch events.
type LaunchEventListResponse struct {
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
	DataMode LaunchEventListResponseDataMode `json:"dataMode,required"`
	// Timestamp of the originating message in ISO8601 UTC format.
	MsgCreateDate time.Time `json:"msgCreateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The Basic Encyclopedia Number, if applicable.
	BeNumber string `json:"beNumber"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// The launch date, in ISO8601 UTC format.
	LaunchDate time.Time `json:"launchDate" format:"date-time"`
	// The Launch facility name.
	LaunchFacilityName string `json:"launchFacilityName"`
	// The DISOB launch Failure Code, if applicable.
	LaunchFailureCode string `json:"launchFailureCode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional target-id, if missing in UDL.
	OrigObjectID string `json:"origObjectId"`
	// The OSuffix, if applicable.
	OSuffix string `json:"oSuffix"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking  resp.Field
		DataMode               resp.Field
		MsgCreateDate          resp.Field
		Source                 resp.Field
		ID                     resp.Field
		BeNumber               resp.Field
		CreatedAt              resp.Field
		CreatedBy              resp.Field
		DeclassificationDate   resp.Field
		DeclassificationString resp.Field
		DerivedFrom            resp.Field
		IDOnOrbit              resp.Field
		LaunchDate             resp.Field
		LaunchFacilityName     resp.Field
		LaunchFailureCode      resp.Field
		Origin                 resp.Field
		OrigNetwork            resp.Field
		OrigObjectID           resp.Field
		OSuffix                resp.Field
		SatNo                  resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchEventListResponse) UnmarshalJSON(data []byte) error {
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
type LaunchEventListResponseDataMode string

const (
	LaunchEventListResponseDataModeReal      LaunchEventListResponseDataMode = "REAL"
	LaunchEventListResponseDataModeTest      LaunchEventListResponseDataMode = "TEST"
	LaunchEventListResponseDataModeSimulated LaunchEventListResponseDataMode = "SIMULATED"
	LaunchEventListResponseDataModeExercise  LaunchEventListResponseDataMode = "EXERCISE"
)

// Information on known launch events.
type LaunchEventGetResponse struct {
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
	DataMode LaunchEventGetResponseDataMode `json:"dataMode,required"`
	// Timestamp of the originating message in ISO8601 UTC format.
	MsgCreateDate time.Time `json:"msgCreateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The Basic Encyclopedia Number, if applicable.
	BeNumber string `json:"beNumber"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// The launch date, in ISO8601 UTC format.
	LaunchDate time.Time `json:"launchDate" format:"date-time"`
	// The Launch facility name.
	LaunchFacilityName string `json:"launchFacilityName"`
	// The DISOB launch Failure Code, if applicable.
	LaunchFailureCode string `json:"launchFailureCode"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional target-id, if missing in UDL.
	OrigObjectID string `json:"origObjectId"`
	// The OSuffix, if applicable.
	OSuffix string `json:"oSuffix"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking  resp.Field
		DataMode               resp.Field
		MsgCreateDate          resp.Field
		Source                 resp.Field
		ID                     resp.Field
		BeNumber               resp.Field
		CreatedAt              resp.Field
		CreatedBy              resp.Field
		DeclassificationDate   resp.Field
		DeclassificationString resp.Field
		DerivedFrom            resp.Field
		IDOnOrbit              resp.Field
		LaunchDate             resp.Field
		LaunchFacilityName     resp.Field
		LaunchFailureCode      resp.Field
		OnOrbit                resp.Field
		Origin                 resp.Field
		OrigNetwork            resp.Field
		OrigObjectID           resp.Field
		OSuffix                resp.Field
		SatNo                  resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchEventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchEventGetResponse) UnmarshalJSON(data []byte) error {
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
type LaunchEventGetResponseDataMode string

const (
	LaunchEventGetResponseDataModeReal      LaunchEventGetResponseDataMode = "REAL"
	LaunchEventGetResponseDataModeTest      LaunchEventGetResponseDataMode = "TEST"
	LaunchEventGetResponseDataModeSimulated LaunchEventGetResponseDataMode = "SIMULATED"
	LaunchEventGetResponseDataModeExercise  LaunchEventGetResponseDataMode = "EXERCISE"
)

// Information on known launch events.
type LaunchEventTupleResponse struct {
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
	DataMode LaunchEventTupleResponseDataMode `json:"dataMode,required"`
	// Timestamp of the originating message in ISO8601 UTC format.
	MsgCreateDate time.Time `json:"msgCreateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The Basic Encyclopedia Number, if applicable.
	BeNumber string `json:"beNumber"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// The launch date, in ISO8601 UTC format.
	LaunchDate time.Time `json:"launchDate" format:"date-time"`
	// The Launch facility name.
	LaunchFacilityName string `json:"launchFacilityName"`
	// The DISOB launch Failure Code, if applicable.
	LaunchFailureCode string `json:"launchFailureCode"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional target-id, if missing in UDL.
	OrigObjectID string `json:"origObjectId"`
	// The OSuffix, if applicable.
	OSuffix string `json:"oSuffix"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking  resp.Field
		DataMode               resp.Field
		MsgCreateDate          resp.Field
		Source                 resp.Field
		ID                     resp.Field
		BeNumber               resp.Field
		CreatedAt              resp.Field
		CreatedBy              resp.Field
		DeclassificationDate   resp.Field
		DeclassificationString resp.Field
		DerivedFrom            resp.Field
		IDOnOrbit              resp.Field
		LaunchDate             resp.Field
		LaunchFacilityName     resp.Field
		LaunchFailureCode      resp.Field
		OnOrbit                resp.Field
		Origin                 resp.Field
		OrigNetwork            resp.Field
		OrigObjectID           resp.Field
		OSuffix                resp.Field
		SatNo                  resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchEventTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchEventTupleResponse) UnmarshalJSON(data []byte) error {
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
type LaunchEventTupleResponseDataMode string

const (
	LaunchEventTupleResponseDataModeReal      LaunchEventTupleResponseDataMode = "REAL"
	LaunchEventTupleResponseDataModeTest      LaunchEventTupleResponseDataMode = "TEST"
	LaunchEventTupleResponseDataModeSimulated LaunchEventTupleResponseDataMode = "SIMULATED"
	LaunchEventTupleResponseDataModeExercise  LaunchEventTupleResponseDataMode = "EXERCISE"
)

type LaunchEventNewParams struct {
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
	DataMode LaunchEventNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Timestamp of the originating message in ISO8601 UTC format.
	MsgCreateDate time.Time `json:"msgCreateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Basic Encyclopedia Number, if applicable.
	BeNumber param.Opt[string] `json:"beNumber,omitzero"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate param.Opt[time.Time] `json:"declassificationDate,omitzero" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString param.Opt[string] `json:"declassificationString,omitzero"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom param.Opt[string] `json:"derivedFrom,omitzero"`
	// The launch date, in ISO8601 UTC format.
	LaunchDate param.Opt[time.Time] `json:"launchDate,omitzero" format:"date-time"`
	// The Launch facility name.
	LaunchFacilityName param.Opt[string] `json:"launchFacilityName,omitzero"`
	// The DISOB launch Failure Code, if applicable.
	LaunchFailureCode param.Opt[string] `json:"launchFailureCode,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional target-id, if missing in UDL.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// The OSuffix, if applicable.
	OSuffix param.Opt[string] `json:"oSuffix,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	paramObj
}

func (r LaunchEventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchEventNewParams
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
type LaunchEventNewParamsDataMode string

const (
	LaunchEventNewParamsDataModeReal      LaunchEventNewParamsDataMode = "REAL"
	LaunchEventNewParamsDataModeTest      LaunchEventNewParamsDataMode = "TEST"
	LaunchEventNewParamsDataModeSimulated LaunchEventNewParamsDataMode = "SIMULATED"
	LaunchEventNewParamsDataModeExercise  LaunchEventNewParamsDataMode = "EXERCISE"
)

type LaunchEventListParams struct {
	// Timestamp of the originating message in ISO8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgCreateDate time.Time        `query:"msgCreateDate,required" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchEventListParams]'s query parameters as `url.Values`.
func (r LaunchEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchEventCountParams struct {
	// Timestamp of the originating message in ISO8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgCreateDate time.Time        `query:"msgCreateDate,required" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchEventCountParams]'s query parameters as `url.Values`.
func (r LaunchEventCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchEventNewBulkParams struct {
	Body []LaunchEventNewBulkParamsBody
	paramObj
}

func (r LaunchEventNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Information on known launch events.
//
// The properties ClassificationMarking, DataMode, MsgCreateDate, Source are
// required.
type LaunchEventNewBulkParamsBody struct {
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
	// Timestamp of the originating message in ISO8601 UTC format.
	MsgCreateDate time.Time `json:"msgCreateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Basic Encyclopedia Number, if applicable.
	BeNumber param.Opt[string] `json:"beNumber,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate param.Opt[time.Time] `json:"declassificationDate,omitzero" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString param.Opt[string] `json:"declassificationString,omitzero"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom param.Opt[string] `json:"derivedFrom,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// The launch date, in ISO8601 UTC format.
	LaunchDate param.Opt[time.Time] `json:"launchDate,omitzero" format:"date-time"`
	// The Launch facility name.
	LaunchFacilityName param.Opt[string] `json:"launchFacilityName,omitzero"`
	// The DISOB launch Failure Code, if applicable.
	LaunchFailureCode param.Opt[string] `json:"launchFailureCode,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional target-id, if missing in UDL.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// The OSuffix, if applicable.
	OSuffix param.Opt[string] `json:"oSuffix,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	paramObj
}

func (r LaunchEventNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaunchEventNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[LaunchEventNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type LaunchEventGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchEventGetParams]'s query parameters as `url.Values`.
func (r LaunchEventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchEventTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Timestamp of the originating message in ISO8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgCreateDate time.Time        `query:"msgCreateDate,required" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchEventTupleParams]'s query parameters as `url.Values`.
func (r LaunchEventTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchEventUnvalidatedPublishParams struct {
	Body []LaunchEventUnvalidatedPublishParamsBody
	paramObj
}

func (r LaunchEventUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Information on known launch events.
//
// The properties ClassificationMarking, DataMode, MsgCreateDate, Source are
// required.
type LaunchEventUnvalidatedPublishParamsBody struct {
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
	// Timestamp of the originating message in ISO8601 UTC format.
	MsgCreateDate time.Time `json:"msgCreateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Basic Encyclopedia Number, if applicable.
	BeNumber param.Opt[string] `json:"beNumber,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate param.Opt[time.Time] `json:"declassificationDate,omitzero" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString param.Opt[string] `json:"declassificationString,omitzero"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom param.Opt[string] `json:"derivedFrom,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// The launch date, in ISO8601 UTC format.
	LaunchDate param.Opt[time.Time] `json:"launchDate,omitzero" format:"date-time"`
	// The Launch facility name.
	LaunchFacilityName param.Opt[string] `json:"launchFacilityName,omitzero"`
	// The DISOB launch Failure Code, if applicable.
	LaunchFailureCode param.Opt[string] `json:"launchFailureCode,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional target-id, if missing in UDL.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// The OSuffix, if applicable.
	OSuffix param.Opt[string] `json:"oSuffix,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	paramObj
}

func (r LaunchEventUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaunchEventUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[LaunchEventUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
