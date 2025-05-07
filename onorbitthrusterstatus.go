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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// OnorbitthrusterstatusService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitthrusterstatusService] method instead.
type OnorbitthrusterstatusService struct {
	Options []option.RequestOption
	History OnorbitthrusterstatusHistoryService
}

// NewOnorbitthrusterstatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOnorbitthrusterstatusService(opts ...option.RequestOption) (r OnorbitthrusterstatusService) {
	r = OnorbitthrusterstatusService{}
	r.Options = opts
	r.History = NewOnorbitthrusterstatusHistoryService(opts...)
	return
}

// Service operation to take a single OnorbitThrusterStatus record as a POST body
// and ingest into the database. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *OnorbitthrusterstatusService) New(ctx context.Context, body OnorbitthrusterstatusNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitthrusterstatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitthrusterstatusService) List(ctx context.Context, query OnorbitthrusterstatusListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitthrusterstatusListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitthrusterstatus"
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
func (r *OnorbitthrusterstatusService) ListAutoPaging(ctx context.Context, query OnorbitthrusterstatusListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitthrusterstatusListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a OnorbitThrusterStatus record specified by the
// passed ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *OnorbitthrusterstatusService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitthrusterstatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnorbitthrusterstatusService) Count(ctx context.Context, query OnorbitthrusterstatusCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onorbitthrusterstatus/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// onorbitthrusterstatus records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *OnorbitthrusterstatusService) NewBulk(ctx context.Context, body OnorbitthrusterstatusNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitthrusterstatus/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single OnorbitThrusterStatus record by its unique ID
// passed as a path parameter. OnorbitThrusterStatus records are information for
// OnorbitThruster objects.
func (r *OnorbitthrusterstatusService) Get(ctx context.Context, id string, query OnorbitthrusterstatusGetParams, opts ...option.RequestOption) (res *OnorbitthrusterstatusFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitthrusterstatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *OnorbitthrusterstatusService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitthrusterstatus/queryhelp"
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
func (r *OnorbitthrusterstatusService) Tuple(ctx context.Context, query OnorbitthrusterstatusTupleParams, opts ...option.RequestOption) (res *[]OnorbitthrusterstatusFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/onorbitthrusterstatus/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Status information for OnorbitThruster objects.
type OnorbitthrusterstatusListResponse struct {
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
	DataMode OnorbitthrusterstatusListResponseDataMode `json:"dataMode,required"`
	// ID of the associated OnorbitThruster record. This ID can be used to obtain
	// additional information on an onorbit thruster object using the 'get by ID'
	// operation (e.g. /udl/onorbitthruster/{id}). For example, the OnorbitThruster
	// object with idOnorbitThruster = abc would be queried as
	// /udl/onorbitthruster/abc.
	IDOnorbitThruster string `json:"idOnorbitThruster,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Datetime of the thruster status observation in ISO 8601 UTC datetime format with
	// millisecond precision.
	StatusTime time.Time `json:"statusTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Estimated available delta-velocity for this thruster, in meters per second.
	EstDeltaV float64 `json:"estDeltaV"`
	// Total fuel mass available for this thruster's type, in kilograms.
	FuelMass float64 `json:"fuelMass"`
	// 1-sigma uncertainty of the total fuel mass available for this thruster type, in
	// kilograms.
	FuelMassUnc float64 `json:"fuelMassUnc"`
	// Specific impulse for this thruster, in seconds.
	Isp float64 `json:"isp"`
	// Maximum available delta-velocity for this thruster, in meters per second.
	MaxDeltaV float64 `json:"maxDeltaV"`
	// Minimum available delta-velocity for this thruster, in meters per second.
	MinDeltaV float64 `json:"minDeltaV"`
	// Identifier of this thruster.
	Name string `json:"name"`
	// Flag indicating if this thruster is operational.
	Operational bool `json:"operational"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Average available propellant mass for this thruster's type, in kilograms.
	PropMassAvg float64 `json:"propMassAvg"`
	// Maximum available propellant mass for this thruster's type, in kilograms.
	PropMassMax float64 `json:"propMassMax"`
	// Median available propellant mass for this thruster's type, in kilograms.
	PropMassMedian float64 `json:"propMassMedian"`
	// Minimum available propellant mass for this thruster's type, in kilograms.
	PropMassMin float64 `json:"propMassMin"`
	// Maximum available thrust for this thruster, in newtons.
	ThrustMax float64 `json:"thrustMax"`
	// Total delta-velocity available for this thruster's type, in meters per second.
	TotalDeltaV float64 `json:"totalDeltaV"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOnorbitThruster     respjson.Field
		Source                respjson.Field
		StatusTime            respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EstDeltaV             respjson.Field
		FuelMass              respjson.Field
		FuelMassUnc           respjson.Field
		Isp                   respjson.Field
		MaxDeltaV             respjson.Field
		MinDeltaV             respjson.Field
		Name                  respjson.Field
		Operational           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PropMassAvg           respjson.Field
		PropMassMax           respjson.Field
		PropMassMedian        respjson.Field
		PropMassMin           respjson.Field
		ThrustMax             respjson.Field
		TotalDeltaV           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitthrusterstatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitthrusterstatusListResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitthrusterstatusListResponseDataMode string

const (
	OnorbitthrusterstatusListResponseDataModeReal      OnorbitthrusterstatusListResponseDataMode = "REAL"
	OnorbitthrusterstatusListResponseDataModeTest      OnorbitthrusterstatusListResponseDataMode = "TEST"
	OnorbitthrusterstatusListResponseDataModeSimulated OnorbitthrusterstatusListResponseDataMode = "SIMULATED"
	OnorbitthrusterstatusListResponseDataModeExercise  OnorbitthrusterstatusListResponseDataMode = "EXERCISE"
)

type OnorbitthrusterstatusNewParams struct {
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
	DataMode OnorbitthrusterstatusNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the associated OnorbitThruster record. This ID can be used to obtain
	// additional information on an onorbit thruster object using the 'get by ID'
	// operation (e.g. /udl/onorbitthruster/{id}). For example, the OnorbitThruster
	// object with idOnorbitThruster = abc would be queried as
	// /udl/onorbitthruster/abc.
	IDOnorbitThruster string `json:"idOnorbitThruster,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Datetime of the thruster status observation in ISO 8601 UTC datetime format with
	// millisecond precision.
	StatusTime time.Time `json:"statusTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Estimated available delta-velocity for this thruster, in meters per second.
	EstDeltaV param.Opt[float64] `json:"estDeltaV,omitzero"`
	// Total fuel mass available for this thruster's type, in kilograms.
	FuelMass param.Opt[float64] `json:"fuelMass,omitzero"`
	// 1-sigma uncertainty of the total fuel mass available for this thruster type, in
	// kilograms.
	FuelMassUnc param.Opt[float64] `json:"fuelMassUnc,omitzero"`
	// Specific impulse for this thruster, in seconds.
	Isp param.Opt[float64] `json:"isp,omitzero"`
	// Maximum available delta-velocity for this thruster, in meters per second.
	MaxDeltaV param.Opt[float64] `json:"maxDeltaV,omitzero"`
	// Minimum available delta-velocity for this thruster, in meters per second.
	MinDeltaV param.Opt[float64] `json:"minDeltaV,omitzero"`
	// Identifier of this thruster.
	Name param.Opt[string] `json:"name,omitzero"`
	// Flag indicating if this thruster is operational.
	Operational param.Opt[bool] `json:"operational,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Average available propellant mass for this thruster's type, in kilograms.
	PropMassAvg param.Opt[float64] `json:"propMassAvg,omitzero"`
	// Maximum available propellant mass for this thruster's type, in kilograms.
	PropMassMax param.Opt[float64] `json:"propMassMax,omitzero"`
	// Median available propellant mass for this thruster's type, in kilograms.
	PropMassMedian param.Opt[float64] `json:"propMassMedian,omitzero"`
	// Minimum available propellant mass for this thruster's type, in kilograms.
	PropMassMin param.Opt[float64] `json:"propMassMin,omitzero"`
	// Maximum available thrust for this thruster, in newtons.
	ThrustMax param.Opt[float64] `json:"thrustMax,omitzero"`
	// Total delta-velocity available for this thruster's type, in meters per second.
	TotalDeltaV param.Opt[float64] `json:"totalDeltaV,omitzero"`
	paramObj
}

func (r OnorbitthrusterstatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitthrusterstatusNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitthrusterstatusNewParams) UnmarshalJSON(data []byte) error {
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
type OnorbitthrusterstatusNewParamsDataMode string

const (
	OnorbitthrusterstatusNewParamsDataModeReal      OnorbitthrusterstatusNewParamsDataMode = "REAL"
	OnorbitthrusterstatusNewParamsDataModeTest      OnorbitthrusterstatusNewParamsDataMode = "TEST"
	OnorbitthrusterstatusNewParamsDataModeSimulated OnorbitthrusterstatusNewParamsDataMode = "SIMULATED"
	OnorbitthrusterstatusNewParamsDataModeExercise  OnorbitthrusterstatusNewParamsDataMode = "EXERCISE"
)

type OnorbitthrusterstatusListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) ID of the
	// associated OnorbitThruster record. This ID can be used to obtain additional
	// information on an onorbit thruster object using the 'get by ID' operation (e.g.
	// /udl/onorbitthruster/{id}). For example, the OnorbitThruster object with
	// idOnorbitThruster = abc would be queried as /udl/onorbitthruster/abc.
	IDOnorbitThruster param.Opt[string] `query:"idOnorbitThruster,omitzero" json:"-"`
	MaxResults        param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) Datetime
	// of the thruster status observation in ISO 8601 UTC datetime format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StatusTime param.Opt[time.Time] `query:"statusTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitthrusterstatusListParams]'s query parameters as
// `url.Values`.
func (r OnorbitthrusterstatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitthrusterstatusCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) ID of the
	// associated OnorbitThruster record. This ID can be used to obtain additional
	// information on an onorbit thruster object using the 'get by ID' operation (e.g.
	// /udl/onorbitthruster/{id}). For example, the OnorbitThruster object with
	// idOnorbitThruster = abc would be queried as /udl/onorbitthruster/abc.
	IDOnorbitThruster param.Opt[string] `query:"idOnorbitThruster,omitzero" json:"-"`
	MaxResults        param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) Datetime
	// of the thruster status observation in ISO 8601 UTC datetime format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StatusTime param.Opt[time.Time] `query:"statusTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitthrusterstatusCountParams]'s query parameters as
// `url.Values`.
func (r OnorbitthrusterstatusCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitthrusterstatusNewBulkParams struct {
	Body []OnorbitthrusterstatusNewBulkParamsBody
	paramObj
}

func (r OnorbitthrusterstatusNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *OnorbitthrusterstatusNewBulkParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// Status information for OnorbitThruster objects.
//
// The properties ClassificationMarking, DataMode, IDOnorbitThruster, Source,
// StatusTime are required.
type OnorbitthrusterstatusNewBulkParamsBody struct {
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
	// ID of the associated OnorbitThruster record. This ID can be used to obtain
	// additional information on an onorbit thruster object using the 'get by ID'
	// operation (e.g. /udl/onorbitthruster/{id}). For example, the OnorbitThruster
	// object with idOnorbitThruster = abc would be queried as
	// /udl/onorbitthruster/abc.
	IDOnorbitThruster string `json:"idOnorbitThruster,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Datetime of the thruster status observation in ISO 8601 UTC datetime format with
	// millisecond precision.
	StatusTime time.Time `json:"statusTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Estimated available delta-velocity for this thruster, in meters per second.
	EstDeltaV param.Opt[float64] `json:"estDeltaV,omitzero"`
	// Total fuel mass available for this thruster's type, in kilograms.
	FuelMass param.Opt[float64] `json:"fuelMass,omitzero"`
	// 1-sigma uncertainty of the total fuel mass available for this thruster type, in
	// kilograms.
	FuelMassUnc param.Opt[float64] `json:"fuelMassUnc,omitzero"`
	// Specific impulse for this thruster, in seconds.
	Isp param.Opt[float64] `json:"isp,omitzero"`
	// Maximum available delta-velocity for this thruster, in meters per second.
	MaxDeltaV param.Opt[float64] `json:"maxDeltaV,omitzero"`
	// Minimum available delta-velocity for this thruster, in meters per second.
	MinDeltaV param.Opt[float64] `json:"minDeltaV,omitzero"`
	// Identifier of this thruster.
	Name param.Opt[string] `json:"name,omitzero"`
	// Flag indicating if this thruster is operational.
	Operational param.Opt[bool] `json:"operational,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Average available propellant mass for this thruster's type, in kilograms.
	PropMassAvg param.Opt[float64] `json:"propMassAvg,omitzero"`
	// Maximum available propellant mass for this thruster's type, in kilograms.
	PropMassMax param.Opt[float64] `json:"propMassMax,omitzero"`
	// Median available propellant mass for this thruster's type, in kilograms.
	PropMassMedian param.Opt[float64] `json:"propMassMedian,omitzero"`
	// Minimum available propellant mass for this thruster's type, in kilograms.
	PropMassMin param.Opt[float64] `json:"propMassMin,omitzero"`
	// Maximum available thrust for this thruster, in newtons.
	ThrustMax param.Opt[float64] `json:"thrustMax,omitzero"`
	// Total delta-velocity available for this thruster's type, in meters per second.
	TotalDeltaV param.Opt[float64] `json:"totalDeltaV,omitzero"`
	paramObj
}

func (r OnorbitthrusterstatusNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitthrusterstatusNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitthrusterstatusNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OnorbitthrusterstatusNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type OnorbitthrusterstatusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitthrusterstatusGetParams]'s query parameters as
// `url.Values`.
func (r OnorbitthrusterstatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitthrusterstatusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) ID of the
	// associated OnorbitThruster record. This ID can be used to obtain additional
	// information on an onorbit thruster object using the 'get by ID' operation (e.g.
	// /udl/onorbitthruster/{id}). For example, the OnorbitThruster object with
	// idOnorbitThruster = abc would be queried as /udl/onorbitthruster/abc.
	IDOnorbitThruster param.Opt[string] `query:"idOnorbitThruster,omitzero" json:"-"`
	MaxResults        param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) Datetime
	// of the thruster status observation in ISO 8601 UTC datetime format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StatusTime param.Opt[time.Time] `query:"statusTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitthrusterstatusTupleParams]'s query parameters as
// `url.Values`.
func (r OnorbitthrusterstatusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
