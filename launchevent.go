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

// LauncheventService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLauncheventService] method instead.
type LauncheventService struct {
	Options []option.RequestOption
}

// NewLauncheventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLauncheventService(opts ...option.RequestOption) (r LauncheventService) {
	r = LauncheventService{}
	r.Options = opts
	return
}

// Service operation to take a single LaunchEvent as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LauncheventService) New(ctx context.Context, body LauncheventNewParams, opts ...option.RequestOption) (err error) {
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
func (r *LauncheventService) List(ctx context.Context, query LauncheventListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LauncheventListResponse], err error) {
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
func (r *LauncheventService) ListAutoPaging(ctx context.Context, query LauncheventListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LauncheventListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LauncheventService) Count(ctx context.Context, query LauncheventCountParams, opts ...option.RequestOption) (res *string, err error) {
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
func (r *LauncheventService) NewBulk(ctx context.Context, body LauncheventNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchevent/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single LaunchEvent record by its unique ID passed as
// a path parameter.
func (r *LauncheventService) Get(ctx context.Context, id string, query LauncheventGetParams, opts ...option.RequestOption) (res *LauncheventGetResponse, err error) {
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
func (r *LauncheventService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
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
func (r *LauncheventService) Tuple(ctx context.Context, query LauncheventTupleParams, opts ...option.RequestOption) (res *[]LauncheventTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/launchevent/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take LaunchEvent entries as a POST body and ingest into the
// database. This operation is intended to be used for automated feeds into UDL. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *LauncheventService) UnvalidatedPublish(ctx context.Context, body LauncheventUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-launchevent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Information on known launch events.
type LauncheventListResponse struct {
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
	DataMode LauncheventListResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
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
func (r LauncheventListResponse) RawJSON() string { return r.JSON.raw }
func (r *LauncheventListResponse) UnmarshalJSON(data []byte) error {
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
type LauncheventListResponseDataMode string

const (
	LauncheventListResponseDataModeReal      LauncheventListResponseDataMode = "REAL"
	LauncheventListResponseDataModeTest      LauncheventListResponseDataMode = "TEST"
	LauncheventListResponseDataModeSimulated LauncheventListResponseDataMode = "SIMULATED"
	LauncheventListResponseDataModeExercise  LauncheventListResponseDataMode = "EXERCISE"
)

// Information on known launch events.
type LauncheventGetResponse struct {
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
	DataMode LauncheventGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
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
func (r LauncheventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LauncheventGetResponse) UnmarshalJSON(data []byte) error {
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
type LauncheventGetResponseDataMode string

const (
	LauncheventGetResponseDataModeReal      LauncheventGetResponseDataMode = "REAL"
	LauncheventGetResponseDataModeTest      LauncheventGetResponseDataMode = "TEST"
	LauncheventGetResponseDataModeSimulated LauncheventGetResponseDataMode = "SIMULATED"
	LauncheventGetResponseDataModeExercise  LauncheventGetResponseDataMode = "EXERCISE"
)

// Information on known launch events.
type LauncheventTupleResponse struct {
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
	DataMode LauncheventTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
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
func (r LauncheventTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LauncheventTupleResponse) UnmarshalJSON(data []byte) error {
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
type LauncheventTupleResponseDataMode string

const (
	LauncheventTupleResponseDataModeReal      LauncheventTupleResponseDataMode = "REAL"
	LauncheventTupleResponseDataModeTest      LauncheventTupleResponseDataMode = "TEST"
	LauncheventTupleResponseDataModeSimulated LauncheventTupleResponseDataMode = "SIMULATED"
	LauncheventTupleResponseDataModeExercise  LauncheventTupleResponseDataMode = "EXERCISE"
)

type LauncheventNewParams struct {
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
	DataMode LauncheventNewParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LauncheventNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LauncheventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LauncheventNewParams
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
type LauncheventNewParamsDataMode string

const (
	LauncheventNewParamsDataModeReal      LauncheventNewParamsDataMode = "REAL"
	LauncheventNewParamsDataModeTest      LauncheventNewParamsDataMode = "TEST"
	LauncheventNewParamsDataModeSimulated LauncheventNewParamsDataMode = "SIMULATED"
	LauncheventNewParamsDataModeExercise  LauncheventNewParamsDataMode = "EXERCISE"
)

type LauncheventListParams struct {
	// Timestamp of the originating message in ISO8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgCreateDate time.Time        `query:"msgCreateDate,required" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LauncheventListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LauncheventListParams]'s query parameters as `url.Values`.
func (r LauncheventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LauncheventCountParams struct {
	// Timestamp of the originating message in ISO8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgCreateDate time.Time        `query:"msgCreateDate,required" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LauncheventCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LauncheventCountParams]'s query parameters as `url.Values`.
func (r LauncheventCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LauncheventNewBulkParams struct {
	Body []LauncheventNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LauncheventNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LauncheventNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Information on known launch events.
//
// The properties ClassificationMarking, DataMode, MsgCreateDate, Source are
// required.
type LauncheventNewBulkParamsBody struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LauncheventNewBulkParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r LauncheventNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LauncheventNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[LauncheventNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type LauncheventGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LauncheventGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LauncheventGetParams]'s query parameters as `url.Values`.
func (r LauncheventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LauncheventTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LauncheventTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LauncheventTupleParams]'s query parameters as `url.Values`.
func (r LauncheventTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LauncheventUnvalidatedPublishParams struct {
	Body []LauncheventUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LauncheventUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r LauncheventUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Information on known launch events.
//
// The properties ClassificationMarking, DataMode, MsgCreateDate, Source are
// required.
type LauncheventUnvalidatedPublishParamsBody struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LauncheventUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LauncheventUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LauncheventUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[LauncheventUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
