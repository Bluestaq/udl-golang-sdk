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

// StageService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStageService] method instead.
type StageService struct {
	Options []option.RequestOption
}

// NewStageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStageService(opts ...option.RequestOption) (r StageService) {
	r = StageService{}
	r.Options = opts
	return
}

// Service operation to take a single Stage as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance. A Stage represents various stages of a
// particular launch vehicle, compiled by a particular source. A vehicle may have
// multiple stage records.
func (r *StageService) New(ctx context.Context, body StageNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/stage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Stage. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance. A
// Stage represents various stages of a particular launch vehicle, compiled by a
// particular source. A vehicle may have multiple stage records.
func (r *StageService) Update(ctx context.Context, id string, body StageUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/stage/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *StageService) List(ctx context.Context, query StageListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[StageListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/stage"
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
func (r *StageService) ListAutoPaging(ctx context.Context, query StageListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[StageListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Stage object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance. A Stage represents various stages of a
// particular launch vehicle, compiled by a particular source. A vehicle may have
// multiple stage records.
func (r *StageService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/stage/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *StageService) Count(ctx context.Context, query StageCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/stage/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Stage record by its unique ID passed as a path
// parameter. A Stage represents various stages of a particular launch vehicle,
// compiled by a particular source. A vehicle may have multiple stage records.
func (r *StageService) Get(ctx context.Context, id string, query StageGetParams, opts ...option.RequestOption) (res *StageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/stage/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *StageService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *StageQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/stage/queryhelp"
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
func (r *StageService) Tuple(ctx context.Context, query StageTupleParams, opts ...option.RequestOption) (res *[]StageTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/stage/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Launch stage information for a particular launch vehicle. A launch vehicle can
// have several stages, each with 1 to many engines.
type StageListResponse struct {
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
	DataMode StageListResponseDataMode `json:"dataMode,required"`
	// Identifier of the Engine record for this stage.
	IDEngine string `json:"idEngine,required"`
	// Identifier of the launch vehicle record for this stage.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Description/notes of the stage avionics.
	AvionicsNotes string `json:"avionicsNotes"`
	// Total burn time of the stage engines in seconds.
	BurnTime float64 `json:"burnTime"`
	// Control thruster 1 type.
	ControlThruster1 string `json:"controlThruster1"`
	// Control thruster 2 type.
	ControlThruster2 string `json:"controlThruster2"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Stage maximum external diameter in meters.
	Diameter float64 `json:"diameter"`
	// Stage length in meters.
	Length float64 `json:"length"`
	// Thrust of the stage main engine at sea level in kN.
	MainEngineThrustSeaLevel float64 `json:"mainEngineThrustSeaLevel"`
	// Thrust of the stage main engine in a vacuum in kN.
	MainEngineThrustVacuum float64 `json:"mainEngineThrustVacuum"`
	// ID of the organization that manufactures this launch stage.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Stage gross mass in kg.
	Mass float64 `json:"mass"`
	// Description/notes of the stage.
	Notes string `json:"notes"`
	// Number of burns for the stage engines.
	NumBurns int64 `json:"numBurns"`
	// Number of type control thruster 1.
	NumControlThruster1 int64 `json:"numControlThruster1"`
	// Number of type control thruster 2.
	NumControlThruster2 int64 `json:"numControlThruster2"`
	// The number of the specified engines on this launch stage.
	NumEngines int64 `json:"numEngines"`
	// Number of launch stage elements used in this stage.
	NumStageElements int64 `json:"numStageElements"`
	// Number of vernier or additional engines.
	NumVernier int64 `json:"numVernier"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Array of URLs of photos of the stage.
	PhotoURLs []string `json:"photoURLs"`
	// Boolean indicating if this launch stage can be restarted.
	Restartable bool `json:"restartable"`
	// Boolean indicating if this launch stage is reusable.
	Reusable bool `json:"reusable"`
	// The stage number of this launch stage.
	StageNumber int64 `json:"stageNumber"`
	// Total thrust of the stage at sea level in kN.
	ThrustSeaLevel float64 `json:"thrustSeaLevel"`
	// Total thrust of the stage in a vacuum in kN.
	ThrustVacuum float64 `json:"thrustVacuum"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	Type string `json:"type"`
	// Engine vernier or additional engine type.
	Vernier string `json:"vernier"`
	// Total burn time of the vernier or additional stage engines in seconds.
	VernierBurnTime float64 `json:"vernierBurnTime"`
	// Total number of burns of the vernier or additional stage engines.
	VernierNumBurns int64 `json:"vernierNumBurns"`
	// Total thrust of one of the vernier or additional engines at sea level in kN.
	VernierThrustSeaLevel float64 `json:"vernierThrustSeaLevel"`
	// Total thrust of one of the vernier or additional engines in a vacuum in kN.
	VernierThrustVacuum float64 `json:"vernierThrustVacuum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		IDEngine                 respjson.Field
		IDLaunchVehicle          respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		AvionicsNotes            respjson.Field
		BurnTime                 respjson.Field
		ControlThruster1         respjson.Field
		ControlThruster2         respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		Diameter                 respjson.Field
		Length                   respjson.Field
		MainEngineThrustSeaLevel respjson.Field
		MainEngineThrustVacuum   respjson.Field
		ManufacturerOrgID        respjson.Field
		Mass                     respjson.Field
		Notes                    respjson.Field
		NumBurns                 respjson.Field
		NumControlThruster1      respjson.Field
		NumControlThruster2      respjson.Field
		NumEngines               respjson.Field
		NumStageElements         respjson.Field
		NumVernier               respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		PhotoURLs                respjson.Field
		Restartable              respjson.Field
		Reusable                 respjson.Field
		StageNumber              respjson.Field
		ThrustSeaLevel           respjson.Field
		ThrustVacuum             respjson.Field
		Type                     respjson.Field
		Vernier                  respjson.Field
		VernierBurnTime          respjson.Field
		VernierNumBurns          respjson.Field
		VernierThrustSeaLevel    respjson.Field
		VernierThrustVacuum      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StageListResponse) RawJSON() string { return r.JSON.raw }
func (r *StageListResponse) UnmarshalJSON(data []byte) error {
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
type StageListResponseDataMode string

const (
	StageListResponseDataModeReal      StageListResponseDataMode = "REAL"
	StageListResponseDataModeTest      StageListResponseDataMode = "TEST"
	StageListResponseDataModeSimulated StageListResponseDataMode = "SIMULATED"
	StageListResponseDataModeExercise  StageListResponseDataMode = "EXERCISE"
)

// Launch stage information for a particular launch vehicle. A launch vehicle can
// have several stages, each with 1 to many engines.
type StageGetResponse struct {
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
	DataMode StageGetResponseDataMode `json:"dataMode,required"`
	// Identifier of the Engine record for this stage.
	IDEngine string `json:"idEngine,required"`
	// Identifier of the launch vehicle record for this stage.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Description/notes of the stage avionics.
	AvionicsNotes string `json:"avionicsNotes"`
	// Total burn time of the stage engines in seconds.
	BurnTime float64 `json:"burnTime"`
	// Control thruster 1 type.
	ControlThruster1 string `json:"controlThruster1"`
	// Control thruster 2 type.
	ControlThruster2 string `json:"controlThruster2"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Stage maximum external diameter in meters.
	Diameter float64 `json:"diameter"`
	// Known launch vehicle engines and their performance characteristics and limits. A
	// launch vehicle has 1 to many engines per stage.
	Engine shared.Engine `json:"engine"`
	// Stage length in meters.
	Length float64 `json:"length"`
	// Thrust of the stage main engine at sea level in kN.
	MainEngineThrustSeaLevel float64 `json:"mainEngineThrustSeaLevel"`
	// Thrust of the stage main engine in a vacuum in kN.
	MainEngineThrustVacuum float64 `json:"mainEngineThrustVacuum"`
	// ID of the organization that manufactures this launch stage.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Stage gross mass in kg.
	Mass float64 `json:"mass"`
	// Description/notes of the stage.
	Notes string `json:"notes"`
	// Number of burns for the stage engines.
	NumBurns int64 `json:"numBurns"`
	// Number of type control thruster 1.
	NumControlThruster1 int64 `json:"numControlThruster1"`
	// Number of type control thruster 2.
	NumControlThruster2 int64 `json:"numControlThruster2"`
	// The number of the specified engines on this launch stage.
	NumEngines int64 `json:"numEngines"`
	// Number of launch stage elements used in this stage.
	NumStageElements int64 `json:"numStageElements"`
	// Number of vernier or additional engines.
	NumVernier int64 `json:"numVernier"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Array of URLs of photos of the stage.
	PhotoURLs []string `json:"photoURLs"`
	// Boolean indicating if this launch stage can be restarted.
	Restartable bool `json:"restartable"`
	// Boolean indicating if this launch stage is reusable.
	Reusable bool `json:"reusable"`
	// The stage number of this launch stage.
	StageNumber int64 `json:"stageNumber"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Total thrust of the stage at sea level in kN.
	ThrustSeaLevel float64 `json:"thrustSeaLevel"`
	// Total thrust of the stage in a vacuum in kN.
	ThrustVacuum float64 `json:"thrustVacuum"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Engine vernier or additional engine type.
	Vernier string `json:"vernier"`
	// Total burn time of the vernier or additional stage engines in seconds.
	VernierBurnTime float64 `json:"vernierBurnTime"`
	// Total number of burns of the vernier or additional stage engines.
	VernierNumBurns int64 `json:"vernierNumBurns"`
	// Total thrust of one of the vernier or additional engines at sea level in kN.
	VernierThrustSeaLevel float64 `json:"vernierThrustSeaLevel"`
	// Total thrust of one of the vernier or additional engines in a vacuum in kN.
	VernierThrustVacuum float64 `json:"vernierThrustVacuum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		IDEngine                 respjson.Field
		IDLaunchVehicle          respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		AvionicsNotes            respjson.Field
		BurnTime                 respjson.Field
		ControlThruster1         respjson.Field
		ControlThruster2         respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		Diameter                 respjson.Field
		Engine                   respjson.Field
		Length                   respjson.Field
		MainEngineThrustSeaLevel respjson.Field
		MainEngineThrustVacuum   respjson.Field
		ManufacturerOrgID        respjson.Field
		Mass                     respjson.Field
		Notes                    respjson.Field
		NumBurns                 respjson.Field
		NumControlThruster1      respjson.Field
		NumControlThruster2      respjson.Field
		NumEngines               respjson.Field
		NumStageElements         respjson.Field
		NumVernier               respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		PhotoURLs                respjson.Field
		Restartable              respjson.Field
		Reusable                 respjson.Field
		StageNumber              respjson.Field
		Tags                     respjson.Field
		ThrustSeaLevel           respjson.Field
		ThrustVacuum             respjson.Field
		Type                     respjson.Field
		UpdatedAt                respjson.Field
		UpdatedBy                respjson.Field
		Vernier                  respjson.Field
		VernierBurnTime          respjson.Field
		VernierNumBurns          respjson.Field
		VernierThrustSeaLevel    respjson.Field
		VernierThrustVacuum      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StageGetResponse) UnmarshalJSON(data []byte) error {
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
type StageGetResponseDataMode string

const (
	StageGetResponseDataModeReal      StageGetResponseDataMode = "REAL"
	StageGetResponseDataModeTest      StageGetResponseDataMode = "TEST"
	StageGetResponseDataModeSimulated StageGetResponseDataMode = "SIMULATED"
	StageGetResponseDataModeExercise  StageGetResponseDataMode = "EXERCISE"
)

type StageQueryhelpResponse struct {
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
func (r StageQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *StageQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Launch stage information for a particular launch vehicle. A launch vehicle can
// have several stages, each with 1 to many engines.
type StageTupleResponse struct {
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
	DataMode StageTupleResponseDataMode `json:"dataMode,required"`
	// Identifier of the Engine record for this stage.
	IDEngine string `json:"idEngine,required"`
	// Identifier of the launch vehicle record for this stage.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Description/notes of the stage avionics.
	AvionicsNotes string `json:"avionicsNotes"`
	// Total burn time of the stage engines in seconds.
	BurnTime float64 `json:"burnTime"`
	// Control thruster 1 type.
	ControlThruster1 string `json:"controlThruster1"`
	// Control thruster 2 type.
	ControlThruster2 string `json:"controlThruster2"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Stage maximum external diameter in meters.
	Diameter float64 `json:"diameter"`
	// Known launch vehicle engines and their performance characteristics and limits. A
	// launch vehicle has 1 to many engines per stage.
	Engine shared.Engine `json:"engine"`
	// Stage length in meters.
	Length float64 `json:"length"`
	// Thrust of the stage main engine at sea level in kN.
	MainEngineThrustSeaLevel float64 `json:"mainEngineThrustSeaLevel"`
	// Thrust of the stage main engine in a vacuum in kN.
	MainEngineThrustVacuum float64 `json:"mainEngineThrustVacuum"`
	// ID of the organization that manufactures this launch stage.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Stage gross mass in kg.
	Mass float64 `json:"mass"`
	// Description/notes of the stage.
	Notes string `json:"notes"`
	// Number of burns for the stage engines.
	NumBurns int64 `json:"numBurns"`
	// Number of type control thruster 1.
	NumControlThruster1 int64 `json:"numControlThruster1"`
	// Number of type control thruster 2.
	NumControlThruster2 int64 `json:"numControlThruster2"`
	// The number of the specified engines on this launch stage.
	NumEngines int64 `json:"numEngines"`
	// Number of launch stage elements used in this stage.
	NumStageElements int64 `json:"numStageElements"`
	// Number of vernier or additional engines.
	NumVernier int64 `json:"numVernier"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Array of URLs of photos of the stage.
	PhotoURLs []string `json:"photoURLs"`
	// Boolean indicating if this launch stage can be restarted.
	Restartable bool `json:"restartable"`
	// Boolean indicating if this launch stage is reusable.
	Reusable bool `json:"reusable"`
	// The stage number of this launch stage.
	StageNumber int64 `json:"stageNumber"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Total thrust of the stage at sea level in kN.
	ThrustSeaLevel float64 `json:"thrustSeaLevel"`
	// Total thrust of the stage in a vacuum in kN.
	ThrustVacuum float64 `json:"thrustVacuum"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Engine vernier or additional engine type.
	Vernier string `json:"vernier"`
	// Total burn time of the vernier or additional stage engines in seconds.
	VernierBurnTime float64 `json:"vernierBurnTime"`
	// Total number of burns of the vernier or additional stage engines.
	VernierNumBurns int64 `json:"vernierNumBurns"`
	// Total thrust of one of the vernier or additional engines at sea level in kN.
	VernierThrustSeaLevel float64 `json:"vernierThrustSeaLevel"`
	// Total thrust of one of the vernier or additional engines in a vacuum in kN.
	VernierThrustVacuum float64 `json:"vernierThrustVacuum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		IDEngine                 respjson.Field
		IDLaunchVehicle          respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		AvionicsNotes            respjson.Field
		BurnTime                 respjson.Field
		ControlThruster1         respjson.Field
		ControlThruster2         respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		Diameter                 respjson.Field
		Engine                   respjson.Field
		Length                   respjson.Field
		MainEngineThrustSeaLevel respjson.Field
		MainEngineThrustVacuum   respjson.Field
		ManufacturerOrgID        respjson.Field
		Mass                     respjson.Field
		Notes                    respjson.Field
		NumBurns                 respjson.Field
		NumControlThruster1      respjson.Field
		NumControlThruster2      respjson.Field
		NumEngines               respjson.Field
		NumStageElements         respjson.Field
		NumVernier               respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		PhotoURLs                respjson.Field
		Restartable              respjson.Field
		Reusable                 respjson.Field
		StageNumber              respjson.Field
		Tags                     respjson.Field
		ThrustSeaLevel           respjson.Field
		ThrustVacuum             respjson.Field
		Type                     respjson.Field
		UpdatedAt                respjson.Field
		UpdatedBy                respjson.Field
		Vernier                  respjson.Field
		VernierBurnTime          respjson.Field
		VernierNumBurns          respjson.Field
		VernierThrustSeaLevel    respjson.Field
		VernierThrustVacuum      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StageTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *StageTupleResponse) UnmarshalJSON(data []byte) error {
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
type StageTupleResponseDataMode string

const (
	StageTupleResponseDataModeReal      StageTupleResponseDataMode = "REAL"
	StageTupleResponseDataModeTest      StageTupleResponseDataMode = "TEST"
	StageTupleResponseDataModeSimulated StageTupleResponseDataMode = "SIMULATED"
	StageTupleResponseDataModeExercise  StageTupleResponseDataMode = "EXERCISE"
)

type StageNewParams struct {
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
	DataMode StageNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the Engine record for this stage.
	IDEngine string `json:"idEngine,required"`
	// Identifier of the launch vehicle record for this stage.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Description/notes of the stage avionics.
	AvionicsNotes param.Opt[string] `json:"avionicsNotes,omitzero"`
	// Total burn time of the stage engines in seconds.
	BurnTime param.Opt[float64] `json:"burnTime,omitzero"`
	// Control thruster 1 type.
	ControlThruster1 param.Opt[string] `json:"controlThruster1,omitzero"`
	// Control thruster 2 type.
	ControlThruster2 param.Opt[string] `json:"controlThruster2,omitzero"`
	// Stage maximum external diameter in meters.
	Diameter param.Opt[float64] `json:"diameter,omitzero"`
	// Stage length in meters.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Thrust of the stage main engine at sea level in kN.
	MainEngineThrustSeaLevel param.Opt[float64] `json:"mainEngineThrustSeaLevel,omitzero"`
	// Thrust of the stage main engine in a vacuum in kN.
	MainEngineThrustVacuum param.Opt[float64] `json:"mainEngineThrustVacuum,omitzero"`
	// ID of the organization that manufactures this launch stage.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Stage gross mass in kg.
	Mass param.Opt[float64] `json:"mass,omitzero"`
	// Description/notes of the stage.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of burns for the stage engines.
	NumBurns param.Opt[int64] `json:"numBurns,omitzero"`
	// Number of type control thruster 1.
	NumControlThruster1 param.Opt[int64] `json:"numControlThruster1,omitzero"`
	// Number of type control thruster 2.
	NumControlThruster2 param.Opt[int64] `json:"numControlThruster2,omitzero"`
	// The number of the specified engines on this launch stage.
	NumEngines param.Opt[int64] `json:"numEngines,omitzero"`
	// Number of launch stage elements used in this stage.
	NumStageElements param.Opt[int64] `json:"numStageElements,omitzero"`
	// Number of vernier or additional engines.
	NumVernier param.Opt[int64] `json:"numVernier,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Boolean indicating if this launch stage can be restarted.
	Restartable param.Opt[bool] `json:"restartable,omitzero"`
	// Boolean indicating if this launch stage is reusable.
	Reusable param.Opt[bool] `json:"reusable,omitzero"`
	// The stage number of this launch stage.
	StageNumber param.Opt[int64] `json:"stageNumber,omitzero"`
	// Total thrust of the stage at sea level in kN.
	ThrustSeaLevel param.Opt[float64] `json:"thrustSeaLevel,omitzero"`
	// Total thrust of the stage in a vacuum in kN.
	ThrustVacuum param.Opt[float64] `json:"thrustVacuum,omitzero"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// Engine vernier or additional engine type.
	Vernier param.Opt[string] `json:"vernier,omitzero"`
	// Total burn time of the vernier or additional stage engines in seconds.
	VernierBurnTime param.Opt[float64] `json:"vernierBurnTime,omitzero"`
	// Total number of burns of the vernier or additional stage engines.
	VernierNumBurns param.Opt[int64] `json:"vernierNumBurns,omitzero"`
	// Total thrust of one of the vernier or additional engines at sea level in kN.
	VernierThrustSeaLevel param.Opt[float64] `json:"vernierThrustSeaLevel,omitzero"`
	// Total thrust of one of the vernier or additional engines in a vacuum in kN.
	VernierThrustVacuum param.Opt[float64] `json:"vernierThrustVacuum,omitzero"`
	// Array of URLs of photos of the stage.
	PhotoURLs []string `json:"photoURLs,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r StageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StageNewParams) UnmarshalJSON(data []byte) error {
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
type StageNewParamsDataMode string

const (
	StageNewParamsDataModeReal      StageNewParamsDataMode = "REAL"
	StageNewParamsDataModeTest      StageNewParamsDataMode = "TEST"
	StageNewParamsDataModeSimulated StageNewParamsDataMode = "SIMULATED"
	StageNewParamsDataModeExercise  StageNewParamsDataMode = "EXERCISE"
)

type StageUpdateParams struct {
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
	DataMode StageUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the Engine record for this stage.
	IDEngine string `json:"idEngine,required"`
	// Identifier of the launch vehicle record for this stage.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Description/notes of the stage avionics.
	AvionicsNotes param.Opt[string] `json:"avionicsNotes,omitzero"`
	// Total burn time of the stage engines in seconds.
	BurnTime param.Opt[float64] `json:"burnTime,omitzero"`
	// Control thruster 1 type.
	ControlThruster1 param.Opt[string] `json:"controlThruster1,omitzero"`
	// Control thruster 2 type.
	ControlThruster2 param.Opt[string] `json:"controlThruster2,omitzero"`
	// Stage maximum external diameter in meters.
	Diameter param.Opt[float64] `json:"diameter,omitzero"`
	// Stage length in meters.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Thrust of the stage main engine at sea level in kN.
	MainEngineThrustSeaLevel param.Opt[float64] `json:"mainEngineThrustSeaLevel,omitzero"`
	// Thrust of the stage main engine in a vacuum in kN.
	MainEngineThrustVacuum param.Opt[float64] `json:"mainEngineThrustVacuum,omitzero"`
	// ID of the organization that manufactures this launch stage.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Stage gross mass in kg.
	Mass param.Opt[float64] `json:"mass,omitzero"`
	// Description/notes of the stage.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of burns for the stage engines.
	NumBurns param.Opt[int64] `json:"numBurns,omitzero"`
	// Number of type control thruster 1.
	NumControlThruster1 param.Opt[int64] `json:"numControlThruster1,omitzero"`
	// Number of type control thruster 2.
	NumControlThruster2 param.Opt[int64] `json:"numControlThruster2,omitzero"`
	// The number of the specified engines on this launch stage.
	NumEngines param.Opt[int64] `json:"numEngines,omitzero"`
	// Number of launch stage elements used in this stage.
	NumStageElements param.Opt[int64] `json:"numStageElements,omitzero"`
	// Number of vernier or additional engines.
	NumVernier param.Opt[int64] `json:"numVernier,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Boolean indicating if this launch stage can be restarted.
	Restartable param.Opt[bool] `json:"restartable,omitzero"`
	// Boolean indicating if this launch stage is reusable.
	Reusable param.Opt[bool] `json:"reusable,omitzero"`
	// The stage number of this launch stage.
	StageNumber param.Opt[int64] `json:"stageNumber,omitzero"`
	// Total thrust of the stage at sea level in kN.
	ThrustSeaLevel param.Opt[float64] `json:"thrustSeaLevel,omitzero"`
	// Total thrust of the stage in a vacuum in kN.
	ThrustVacuum param.Opt[float64] `json:"thrustVacuum,omitzero"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// Engine vernier or additional engine type.
	Vernier param.Opt[string] `json:"vernier,omitzero"`
	// Total burn time of the vernier or additional stage engines in seconds.
	VernierBurnTime param.Opt[float64] `json:"vernierBurnTime,omitzero"`
	// Total number of burns of the vernier or additional stage engines.
	VernierNumBurns param.Opt[int64] `json:"vernierNumBurns,omitzero"`
	// Total thrust of one of the vernier or additional engines at sea level in kN.
	VernierThrustSeaLevel param.Opt[float64] `json:"vernierThrustSeaLevel,omitzero"`
	// Total thrust of one of the vernier or additional engines in a vacuum in kN.
	VernierThrustVacuum param.Opt[float64] `json:"vernierThrustVacuum,omitzero"`
	// Array of URLs of photos of the stage.
	PhotoURLs []string `json:"photoURLs,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r StageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StageUpdateParams) UnmarshalJSON(data []byte) error {
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
type StageUpdateParamsDataMode string

const (
	StageUpdateParamsDataModeReal      StageUpdateParamsDataMode = "REAL"
	StageUpdateParamsDataModeTest      StageUpdateParamsDataMode = "TEST"
	StageUpdateParamsDataModeSimulated StageUpdateParamsDataMode = "SIMULATED"
	StageUpdateParamsDataModeExercise  StageUpdateParamsDataMode = "EXERCISE"
)

type StageListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StageListParams]'s query parameters as `url.Values`.
func (r StageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StageCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StageCountParams]'s query parameters as `url.Values`.
func (r StageCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StageGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StageGetParams]'s query parameters as `url.Values`.
func (r StageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StageTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StageTupleParams]'s query parameters as `url.Values`.
func (r StageTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
