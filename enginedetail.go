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

// EngineDetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEngineDetailService] method instead.
type EngineDetailService struct {
	Options []option.RequestOption
}

// NewEngineDetailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEngineDetailService(opts ...option.RequestOption) (r EngineDetailService) {
	r = EngineDetailService{}
	r.Options = opts
	return
}

// Service operation to take a single EngineDetails as a POST body and ingest into
// the database. EngineDetails are launch vehicle engine details and performance
// characteristics/limits compiled by a particular source. A launch vehicle engine
// may have several details records from multiple sources. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *EngineDetailService) New(ctx context.Context, body EngineDetailNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/enginedetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single EngineDetails record by its unique ID passed
// as a path parameter. EngineDetails are launch vehicle engine details and
// performance characteristics/limits compiled by a particular source. A launch
// vehicle engine may have several details records from multiple sources.
func (r *EngineDetailService) Get(ctx context.Context, id string, query EngineDetailGetParams, opts ...option.RequestOption) (res *shared.EngineDetailsFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/enginedetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single EngineDetails. EngineDetails are launch
// vehicle engine details and performance characteristics/limits compiled by a
// particular source. A launch vehicle engine may have several details records from
// multiple sources. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *EngineDetailService) Update(ctx context.Context, id string, body EngineDetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/enginedetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EngineDetailService) List(ctx context.Context, query EngineDetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EngineDetailsAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/enginedetails"
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
func (r *EngineDetailService) ListAutoPaging(ctx context.Context, query EngineDetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EngineDetailsAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a EngineDetails object specified by the passed ID
// path parameter. EngineDetails are launch vehicle engine details and performance
// characteristics/limits compiled by a particular source. A launch vehicle engine
// may have several details records from multiple sources. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *EngineDetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/enginedetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Known launch vehicle engine details and performance characteristics and limits
// compiled by a particular source. A launch vehicle engine may have several
// details records from multiple sources.
type EngineDetailsAbridged struct {
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
	DataMode EngineDetailsAbridgedDataMode `json:"dataMode,required"`
	// Identifier of the parent engine record.
	IDEngine string `json:"idEngine,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Launch engine maximum burn time in seconds.
	BurnTime float64 `json:"burnTime"`
	// Engine chamber pressure in bars.
	ChamberPressure float64 `json:"chamberPressure"`
	// Engine characteristic type (e.g. Electric, Mono-propellant, Bi-propellant,
	// etc.).
	CharacteristicType string `json:"characteristicType"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	CycleType string `json:"cycleType"`
	// Engine type or family.
	Family string `json:"family"`
	// Organization ID of the engine manufacturer.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Engine maximum number of firings.
	MaxFirings int64 `json:"maxFirings"`
	// Notes/Description of the engine.
	Notes string `json:"notes"`
	// Engine nozzle expansion ratio.
	NozzleExpansionRatio float64 `json:"nozzleExpansionRatio"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Oxidizer type (e.g. Nitrogen Tetroxide, Liquid Oxygen, etc).
	Oxidizer string `json:"oxidizer"`
	// Propellant/fuel type of the engine (e.g. Liquid Hydrogen, Kerosene, Aerozine,
	// etc).
	Propellant string `json:"propellant"`
	// Engine maximum thrust at sea level in Kilo-Newtons.
	SeaLevelThrust float64 `json:"seaLevelThrust"`
	// Launch engine specific impulse in seconds.
	SpecificImpulse float64 `json:"specificImpulse"`
	// Engine maximum thrust in a vacuum in Kilo-Newtons.
	VacuumThrust float64 `json:"vacuumThrust"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDEngine              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		BurnTime              respjson.Field
		ChamberPressure       respjson.Field
		CharacteristicType    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CycleType             respjson.Field
		Family                respjson.Field
		ManufacturerOrgID     respjson.Field
		MaxFirings            respjson.Field
		Notes                 respjson.Field
		NozzleExpansionRatio  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Oxidizer              respjson.Field
		Propellant            respjson.Field
		SeaLevelThrust        respjson.Field
		SpecificImpulse       respjson.Field
		VacuumThrust          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EngineDetailsAbridged) RawJSON() string { return r.JSON.raw }
func (r *EngineDetailsAbridged) UnmarshalJSON(data []byte) error {
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
type EngineDetailsAbridgedDataMode string

const (
	EngineDetailsAbridgedDataModeReal      EngineDetailsAbridgedDataMode = "REAL"
	EngineDetailsAbridgedDataModeTest      EngineDetailsAbridgedDataMode = "TEST"
	EngineDetailsAbridgedDataModeSimulated EngineDetailsAbridgedDataMode = "SIMULATED"
	EngineDetailsAbridgedDataModeExercise  EngineDetailsAbridgedDataMode = "EXERCISE"
)

type EngineDetailNewParams struct {
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
	DataMode EngineDetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the parent engine record.
	IDEngine string `json:"idEngine,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Launch engine maximum burn time in seconds.
	BurnTime param.Opt[float64] `json:"burnTime,omitzero"`
	// Engine chamber pressure in bars.
	ChamberPressure param.Opt[float64] `json:"chamberPressure,omitzero"`
	// Engine characteristic type (e.g. Electric, Mono-propellant, Bi-propellant,
	// etc.).
	CharacteristicType param.Opt[string] `json:"characteristicType,omitzero"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	CycleType param.Opt[string] `json:"cycleType,omitzero"`
	// Engine type or family.
	Family param.Opt[string] `json:"family,omitzero"`
	// Organization ID of the engine manufacturer.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Engine maximum number of firings.
	MaxFirings param.Opt[int64] `json:"maxFirings,omitzero"`
	// Notes/Description of the engine.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Engine nozzle expansion ratio.
	NozzleExpansionRatio param.Opt[float64] `json:"nozzleExpansionRatio,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Oxidizer type (e.g. Nitrogen Tetroxide, Liquid Oxygen, etc).
	Oxidizer param.Opt[string] `json:"oxidizer,omitzero"`
	// Propellant/fuel type of the engine (e.g. Liquid Hydrogen, Kerosene, Aerozine,
	// etc).
	Propellant param.Opt[string] `json:"propellant,omitzero"`
	// Engine maximum thrust at sea level in Kilo-Newtons.
	SeaLevelThrust param.Opt[float64] `json:"seaLevelThrust,omitzero"`
	// Launch engine specific impulse in seconds.
	SpecificImpulse param.Opt[float64] `json:"specificImpulse,omitzero"`
	// Engine maximum thrust in a vacuum in Kilo-Newtons.
	VacuumThrust param.Opt[float64] `json:"vacuumThrust,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r EngineDetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EngineDetailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EngineDetailNewParams) UnmarshalJSON(data []byte) error {
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
type EngineDetailNewParamsDataMode string

const (
	EngineDetailNewParamsDataModeReal      EngineDetailNewParamsDataMode = "REAL"
	EngineDetailNewParamsDataModeTest      EngineDetailNewParamsDataMode = "TEST"
	EngineDetailNewParamsDataModeSimulated EngineDetailNewParamsDataMode = "SIMULATED"
	EngineDetailNewParamsDataModeExercise  EngineDetailNewParamsDataMode = "EXERCISE"
)

type EngineDetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EngineDetailGetParams]'s query parameters as `url.Values`.
func (r EngineDetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EngineDetailUpdateParams struct {
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
	DataMode EngineDetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the parent engine record.
	IDEngine string `json:"idEngine,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Launch engine maximum burn time in seconds.
	BurnTime param.Opt[float64] `json:"burnTime,omitzero"`
	// Engine chamber pressure in bars.
	ChamberPressure param.Opt[float64] `json:"chamberPressure,omitzero"`
	// Engine characteristic type (e.g. Electric, Mono-propellant, Bi-propellant,
	// etc.).
	CharacteristicType param.Opt[string] `json:"characteristicType,omitzero"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	CycleType param.Opt[string] `json:"cycleType,omitzero"`
	// Engine type or family.
	Family param.Opt[string] `json:"family,omitzero"`
	// Organization ID of the engine manufacturer.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Engine maximum number of firings.
	MaxFirings param.Opt[int64] `json:"maxFirings,omitzero"`
	// Notes/Description of the engine.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Engine nozzle expansion ratio.
	NozzleExpansionRatio param.Opt[float64] `json:"nozzleExpansionRatio,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Oxidizer type (e.g. Nitrogen Tetroxide, Liquid Oxygen, etc).
	Oxidizer param.Opt[string] `json:"oxidizer,omitzero"`
	// Propellant/fuel type of the engine (e.g. Liquid Hydrogen, Kerosene, Aerozine,
	// etc).
	Propellant param.Opt[string] `json:"propellant,omitzero"`
	// Engine maximum thrust at sea level in Kilo-Newtons.
	SeaLevelThrust param.Opt[float64] `json:"seaLevelThrust,omitzero"`
	// Launch engine specific impulse in seconds.
	SpecificImpulse param.Opt[float64] `json:"specificImpulse,omitzero"`
	// Engine maximum thrust in a vacuum in Kilo-Newtons.
	VacuumThrust param.Opt[float64] `json:"vacuumThrust,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r EngineDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EngineDetailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EngineDetailUpdateParams) UnmarshalJSON(data []byte) error {
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
type EngineDetailUpdateParamsDataMode string

const (
	EngineDetailUpdateParamsDataModeReal      EngineDetailUpdateParamsDataMode = "REAL"
	EngineDetailUpdateParamsDataModeTest      EngineDetailUpdateParamsDataMode = "TEST"
	EngineDetailUpdateParamsDataModeSimulated EngineDetailUpdateParamsDataMode = "SIMULATED"
	EngineDetailUpdateParamsDataModeExercise  EngineDetailUpdateParamsDataMode = "EXERCISE"
)

type EngineDetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EngineDetailListParams]'s query parameters as `url.Values`.
func (r EngineDetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
