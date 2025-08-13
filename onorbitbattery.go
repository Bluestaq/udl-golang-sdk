// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

// OnorbitbatteryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitbatteryService] method instead.
type OnorbitbatteryService struct {
	Options []option.RequestOption
}

// NewOnorbitbatteryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOnorbitbatteryService(opts ...option.RequestOption) (r OnorbitbatteryService) {
	r = OnorbitbatteryService{}
	r.Options = opts
	return
}

// Service operation to take a single OnorbitBattery as a POST body and ingest into
// the database. OnorbitBattery is the association between on-orbit spacecraft
// batteries and a particular on-orbit spacecraft. A Battery may be associated with
// many different on-orbit spacecraft. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *OnorbitbatteryService) New(ctx context.Context, body OnorbitbatteryNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitbattery"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single OnorbitBattery. OnorbitBattery is the
// association between on-orbit spacecraft batteries and a particular on-orbit
// spacecraft. A Battery may be associated with many different on-orbit spacecraft.
// A specific role is required to perform this service operation. Please contact
// the UDL team for assistance.
func (r *OnorbitbatteryService) Update(ctx context.Context, id string, body OnorbitbatteryUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitbattery/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitbatteryService) List(ctx context.Context, query OnorbitbatteryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitbatteryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitbattery"
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
func (r *OnorbitbatteryService) ListAutoPaging(ctx context.Context, query OnorbitbatteryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitbatteryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a OnorbitBattery object specified by the passed ID
// path parameter. OnorbitBattery is the association between on-orbit spacecraft
// batteries and a particular on-orbit spacecraft. A Battery may be associated with
// many different on-orbit spacecraft. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *OnorbitbatteryService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitbattery/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to get a single OnorbitBattery record by its unique ID passed
// as a path parameter. OnorbitBattery is the association between on-orbit
// spacecraft batteries and a particular on-orbit spacecraft. A Battery may be
// associated with many different on-orbit spacecraft.
func (r *OnorbitbatteryService) Get(ctx context.Context, id string, query OnorbitbatteryGetParams, opts ...option.RequestOption) (res *shared.OnorbitBatteryFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitbattery/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OnorbitbatteryListResponse struct {
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
	DataMode OnorbitbatteryListResponseDataMode `json:"dataMode,required"`
	// ID of the battery.
	IDBattery string `json:"idBattery,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Model representation of specific spacecraft battery types.
	Battery BatteryAbridged `json:"battery"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The number of batteries on the spacecraft of the type identified by idBattery.
	Quantity int64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDBattery             respjson.Field
		IDOnOrbit             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Battery               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Quantity              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitbatteryListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitbatteryListResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitbatteryListResponseDataMode string

const (
	OnorbitbatteryListResponseDataModeReal      OnorbitbatteryListResponseDataMode = "REAL"
	OnorbitbatteryListResponseDataModeTest      OnorbitbatteryListResponseDataMode = "TEST"
	OnorbitbatteryListResponseDataModeSimulated OnorbitbatteryListResponseDataMode = "SIMULATED"
	OnorbitbatteryListResponseDataModeExercise  OnorbitbatteryListResponseDataMode = "EXERCISE"
)

type OnorbitbatteryNewParams struct {
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
	DataMode OnorbitbatteryNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the battery.
	IDBattery string `json:"idBattery,required"`
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
	// The number of batteries on the spacecraft of the type identified by idBattery.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Model representation of specific spacecraft battery types.
	Battery shared.BatteryIngestParam `json:"battery,omitzero"`
	paramObj
}

func (r OnorbitbatteryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitbatteryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitbatteryNewParams) UnmarshalJSON(data []byte) error {
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
type OnorbitbatteryNewParamsDataMode string

const (
	OnorbitbatteryNewParamsDataModeReal      OnorbitbatteryNewParamsDataMode = "REAL"
	OnorbitbatteryNewParamsDataModeTest      OnorbitbatteryNewParamsDataMode = "TEST"
	OnorbitbatteryNewParamsDataModeSimulated OnorbitbatteryNewParamsDataMode = "SIMULATED"
	OnorbitbatteryNewParamsDataModeExercise  OnorbitbatteryNewParamsDataMode = "EXERCISE"
)

type OnorbitbatteryUpdateParams struct {
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
	DataMode OnorbitbatteryUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the battery.
	IDBattery string `json:"idBattery,required"`
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
	// The number of batteries on the spacecraft of the type identified by idBattery.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Model representation of specific spacecraft battery types.
	Battery shared.BatteryIngestParam `json:"battery,omitzero"`
	paramObj
}

func (r OnorbitbatteryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitbatteryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitbatteryUpdateParams) UnmarshalJSON(data []byte) error {
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
type OnorbitbatteryUpdateParamsDataMode string

const (
	OnorbitbatteryUpdateParamsDataModeReal      OnorbitbatteryUpdateParamsDataMode = "REAL"
	OnorbitbatteryUpdateParamsDataModeTest      OnorbitbatteryUpdateParamsDataMode = "TEST"
	OnorbitbatteryUpdateParamsDataModeSimulated OnorbitbatteryUpdateParamsDataMode = "SIMULATED"
	OnorbitbatteryUpdateParamsDataModeExercise  OnorbitbatteryUpdateParamsDataMode = "EXERCISE"
)

type OnorbitbatteryListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitbatteryListParams]'s query parameters as
// `url.Values`.
func (r OnorbitbatteryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitbatteryGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitbatteryGetParams]'s query parameters as
// `url.Values`.
func (r OnorbitbatteryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
