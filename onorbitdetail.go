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

// OnorbitdetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitdetailService] method instead.
type OnorbitdetailService struct {
	Options []option.RequestOption
}

// NewOnorbitdetailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOnorbitdetailService(opts ...option.RequestOption) (r OnorbitdetailService) {
	r = OnorbitdetailService{}
	r.Options = opts
	return
}

// Service operation to take a single OnorbitDetails as a POST body and ingest into
// the database. An OnorbitDetails is a collection of additional characteristics on
// an on-orbit object. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *OnorbitdetailService) New(ctx context.Context, body OnorbitdetailNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/onorbitdetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single OnorbitDetails. An OnorbitDetails is a
// collection of additional characteristics on an on-orbit object. A specific role
// is required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *OnorbitdetailService) Update(ctx context.Context, id string, body OnorbitdetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitdetailService) List(ctx context.Context, query OnorbitdetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitdetailListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitdetails"
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
func (r *OnorbitdetailService) ListAutoPaging(ctx context.Context, query OnorbitdetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitdetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a OnorbitDetails object specified by the passed ID
// path parameter. An OnorbitDetails is a collection of additional characteristics
// on an on-orbit object. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *OnorbitdetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to get a single OnorbitDetails record by its unique ID passed
// as a path parameter. An OnorbitDetails is a collection of additional
// characteristics on an on-orbit object.
func (r *OnorbitdetailService) Get(ctx context.Context, id string, query OnorbitdetailGetParams, opts ...option.RequestOption) (res *shared.OnorbitDetailsFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Contains details of the OnOrbit object.
type OnorbitdetailListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode OnorbitdetailListResponseDataMode `json:"dataMode,required"`
	// UUID of the parent Onorbit record.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Mass of fuel and disposables at launch time in kilograms.
	AdditionalMass float64 `json:"additionalMass"`
	// The radius used for long-term debris environment projection analyses that is not
	// as conservative as COLA Radius, in meters.
	AdeptRadius float64 `json:"adeptRadius"`
	// The total beginning of life delta V of the spacecraft, in meters per second.
	BolDeltaV float64 `json:"bolDeltaV"`
	// Spacecraft beginning of life fuel mass, in orbit, in kilograms.
	BolFuelMass float64 `json:"bolFuelMass"`
	// Average cross sectional area of the bus in meters squared.
	BusCrossSection float64 `json:"busCrossSection"`
	// Type of the bus on the spacecraft.
	BusType string `json:"busType"`
	// Maximum dimension of the box circumscribing the spacecraft (d = sqrt(a*a + b*b +
	// c\*c) where a is the tip-to-tip dimension, b and c are perpendicular to that.)
	// in meters.
	ColaRadius float64 `json:"colaRadius"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Average cross sectional area in meters squared.
	CrossSection float64 `json:"crossSection"`
	// The estimated total current mass of the spacecraft, in kilograms.
	CurrentMass float64 `json:"currentMass"`
	// The 1-sigma uncertainty of the total spacecraft delta V, in meters per second.
	DeltaVUnc float64 `json:"deltaVUnc"`
	// Array of the estimated mass of each deployable object, in kilograms. Must
	// contain the same number of elements as the value of numDeployable.
	DepEstMasses []float64 `json:"depEstMasses"`
	// Array of the 1-sigma uncertainty of the mass for each deployable object, in
	// kilograms. Must contain the same number of elements as the value of
	// numDeployable.
	DepMassUncs []float64 `json:"depMassUncs"`
	// Array of satellite deployable objects. Must contain the same number of elements
	// as the value of numDeployable.
	DepNames []string `json:"depNames"`
	// GEO drift rate, if applicable in degrees per day.
	DriftRate float64 `json:"driftRate"`
	// Spacecraft dry mass (without fuel or disposables) in kilograms.
	DryMass float64 `json:"dryMass"`
	// Estimated maximum burn duration for the object, in seconds.
	EstDeltaVDuration float64 `json:"estDeltaVDuration"`
	// Estimated remaining fuel for the object in kilograms.
	FuelRemaining float64 `json:"fuelRemaining"`
	// GEO slot if applicable, in degrees. -180 (West of Prime Meridian) to 180 degrees
	// (East of Prime Meridian). Prime Meridian is 0.
	GeoSlot float64 `json:"geoSlot"`
	// The name of the source who last provided an observation for this idOnOrbit.
	LastObSource string `json:"lastObSource"`
	// Time of last reported observation for this object in ISO 8601 UTC with
	// microsecond precision.
	LastObTime time.Time `json:"lastObTime" format:"date-time"`
	// Nominal mass of spacecraft and fuel at launch time, in kilograms.
	LaunchMass float64 `json:"launchMass"`
	// Maximum (estimated) mass of spacecraft and fuel at launch time, in kilograms.
	LaunchMassMax float64 `json:"launchMassMax"`
	// Minimum (estimated) mass of spacecraft and fuel at launch time, in kilograms.
	LaunchMassMin float64 `json:"launchMassMin"`
	// Boolean indicating whether a spacecraft is maneuverable. Note that a spacecraft
	// may have propulsion capability but may not be maneuverable due to lack of fuel,
	// anomalous condition, or other operational constraints.
	Maneuverable bool `json:"maneuverable"`
	// Maximum delta V available for this on-orbit spacecraft, in meters per second.
	MaxDeltaV float64 `json:"maxDeltaV"`
	// Maximum dimension across the spacecraft (e.g., tip-to-tip across the solar panel
	// arrays) in meters.
	MaxRadius float64 `json:"maxRadius"`
	// Array of the type of missions the spacecraft performs. Must contain the same
	// number of elements as the value of numMission.
	MissionTypes []string `json:"missionTypes"`
	// The number of sub-satellites or deployable objects on the spacecraft.
	NumDeployable int64 `json:"numDeployable"`
	// The number of distinct missions the spacecraft performs.
	NumMission int64 `json:"numMission"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Current/latest radar cross section in meters squared.
	Rcs float64 `json:"rcs"`
	// Maximum radar cross section in meters squared.
	RcsMax float64 `json:"rcsMax"`
	// Mean radar cross section in meters squared.
	RcsMean float64 `json:"rcsMean"`
	// Minimum radar cross section in meters squared.
	RcsMin float64 `json:"rcsMin"`
	// The reference source, sources, or URL from which the data in this record was
	// obtained.
	RefSource string `json:"refSource"`
	// Spacecraft deployed area of solar array in meters squared.
	SolarArrayArea float64 `json:"solarArrayArea"`
	// The 1-sigma uncertainty of the total spacecraft mass, in kilograms.
	TotalMassUnc float64 `json:"totalMassUnc"`
	// Current/latest visual magnitude in M.
	Vismag float64 `json:"vismag"`
	// Maximum visual magnitude in M.
	VismagMax float64 `json:"vismagMax"`
	// Mean visual magnitude in M.
	VismagMean float64 `json:"vismagMean"`
	// Minimum visual magnitude in M.
	VismagMin float64 `json:"vismagMin"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOnOrbit             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AdditionalMass        respjson.Field
		AdeptRadius           respjson.Field
		BolDeltaV             respjson.Field
		BolFuelMass           respjson.Field
		BusCrossSection       respjson.Field
		BusType               respjson.Field
		ColaRadius            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CrossSection          respjson.Field
		CurrentMass           respjson.Field
		DeltaVUnc             respjson.Field
		DepEstMasses          respjson.Field
		DepMassUncs           respjson.Field
		DepNames              respjson.Field
		DriftRate             respjson.Field
		DryMass               respjson.Field
		EstDeltaVDuration     respjson.Field
		FuelRemaining         respjson.Field
		GeoSlot               respjson.Field
		LastObSource          respjson.Field
		LastObTime            respjson.Field
		LaunchMass            respjson.Field
		LaunchMassMax         respjson.Field
		LaunchMassMin         respjson.Field
		Maneuverable          respjson.Field
		MaxDeltaV             respjson.Field
		MaxRadius             respjson.Field
		MissionTypes          respjson.Field
		NumDeployable         respjson.Field
		NumMission            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Rcs                   respjson.Field
		RcsMax                respjson.Field
		RcsMean               respjson.Field
		RcsMin                respjson.Field
		RefSource             respjson.Field
		SolarArrayArea        respjson.Field
		TotalMassUnc          respjson.Field
		Vismag                respjson.Field
		VismagMax             respjson.Field
		VismagMean            respjson.Field
		VismagMin             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitdetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitdetailListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type OnorbitdetailListResponseDataMode string

const (
	OnorbitdetailListResponseDataModeReal      OnorbitdetailListResponseDataMode = "REAL"
	OnorbitdetailListResponseDataModeTest      OnorbitdetailListResponseDataMode = "TEST"
	OnorbitdetailListResponseDataModeSimulated OnorbitdetailListResponseDataMode = "SIMULATED"
	OnorbitdetailListResponseDataModeExercise  OnorbitdetailListResponseDataMode = "EXERCISE"
)

type OnorbitdetailNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode OnorbitdetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// UUID of the parent Onorbit record.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Mass of fuel and disposables at launch time in kilograms.
	AdditionalMass param.Opt[float64] `json:"additionalMass,omitzero"`
	// The radius used for long-term debris environment projection analyses that is not
	// as conservative as COLA Radius, in meters.
	AdeptRadius param.Opt[float64] `json:"adeptRadius,omitzero"`
	// The total beginning of life delta V of the spacecraft, in meters per second.
	BolDeltaV param.Opt[float64] `json:"bolDeltaV,omitzero"`
	// Spacecraft beginning of life fuel mass, in orbit, in kilograms.
	BolFuelMass param.Opt[float64] `json:"bolFuelMass,omitzero"`
	// Average cross sectional area of the bus in meters squared.
	BusCrossSection param.Opt[float64] `json:"busCrossSection,omitzero"`
	// Type of the bus on the spacecraft.
	BusType param.Opt[string] `json:"busType,omitzero"`
	// Maximum dimension of the box circumscribing the spacecraft (d = sqrt(a*a + b*b +
	// c\*c) where a is the tip-to-tip dimension, b and c are perpendicular to that.)
	// in meters.
	ColaRadius param.Opt[float64] `json:"colaRadius,omitzero"`
	// Average cross sectional area in meters squared.
	CrossSection param.Opt[float64] `json:"crossSection,omitzero"`
	// The estimated total current mass of the spacecraft, in kilograms.
	CurrentMass param.Opt[float64] `json:"currentMass,omitzero"`
	// The 1-sigma uncertainty of the total spacecraft delta V, in meters per second.
	DeltaVUnc param.Opt[float64] `json:"deltaVUnc,omitzero"`
	// GEO drift rate, if applicable in degrees per day.
	DriftRate param.Opt[float64] `json:"driftRate,omitzero"`
	// Spacecraft dry mass (without fuel or disposables) in kilograms.
	DryMass param.Opt[float64] `json:"dryMass,omitzero"`
	// Estimated maximum burn duration for the object, in seconds.
	EstDeltaVDuration param.Opt[float64] `json:"estDeltaVDuration,omitzero"`
	// Estimated remaining fuel for the object in kilograms.
	FuelRemaining param.Opt[float64] `json:"fuelRemaining,omitzero"`
	// GEO slot if applicable, in degrees. -180 (West of Prime Meridian) to 180 degrees
	// (East of Prime Meridian). Prime Meridian is 0.
	GeoSlot param.Opt[float64] `json:"geoSlot,omitzero"`
	// The name of the source who last provided an observation for this idOnOrbit.
	LastObSource param.Opt[string] `json:"lastObSource,omitzero"`
	// Time of last reported observation for this object in ISO 8601 UTC with
	// microsecond precision.
	LastObTime param.Opt[time.Time] `json:"lastObTime,omitzero" format:"date-time"`
	// Nominal mass of spacecraft and fuel at launch time, in kilograms.
	LaunchMass param.Opt[float64] `json:"launchMass,omitzero"`
	// Maximum (estimated) mass of spacecraft and fuel at launch time, in kilograms.
	LaunchMassMax param.Opt[float64] `json:"launchMassMax,omitzero"`
	// Minimum (estimated) mass of spacecraft and fuel at launch time, in kilograms.
	LaunchMassMin param.Opt[float64] `json:"launchMassMin,omitzero"`
	// Boolean indicating whether a spacecraft is maneuverable. Note that a spacecraft
	// may have propulsion capability but may not be maneuverable due to lack of fuel,
	// anomalous condition, or other operational constraints.
	Maneuverable param.Opt[bool] `json:"maneuverable,omitzero"`
	// Maximum delta V available for this on-orbit spacecraft, in meters per second.
	MaxDeltaV param.Opt[float64] `json:"maxDeltaV,omitzero"`
	// Maximum dimension across the spacecraft (e.g., tip-to-tip across the solar panel
	// arrays) in meters.
	MaxRadius param.Opt[float64] `json:"maxRadius,omitzero"`
	// The number of sub-satellites or deployable objects on the spacecraft.
	NumDeployable param.Opt[int64] `json:"numDeployable,omitzero"`
	// The number of distinct missions the spacecraft performs.
	NumMission param.Opt[int64] `json:"numMission,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Current/latest radar cross section in meters squared.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// Maximum radar cross section in meters squared.
	RcsMax param.Opt[float64] `json:"rcsMax,omitzero"`
	// Mean radar cross section in meters squared.
	RcsMean param.Opt[float64] `json:"rcsMean,omitzero"`
	// Minimum radar cross section in meters squared.
	RcsMin param.Opt[float64] `json:"rcsMin,omitzero"`
	// The reference source, sources, or URL from which the data in this record was
	// obtained.
	RefSource param.Opt[string] `json:"refSource,omitzero"`
	// Spacecraft deployed area of solar array in meters squared.
	SolarArrayArea param.Opt[float64] `json:"solarArrayArea,omitzero"`
	// The 1-sigma uncertainty of the total spacecraft mass, in kilograms.
	TotalMassUnc param.Opt[float64] `json:"totalMassUnc,omitzero"`
	// Current/latest visual magnitude in M.
	Vismag param.Opt[float64] `json:"vismag,omitzero"`
	// Maximum visual magnitude in M.
	VismagMax param.Opt[float64] `json:"vismagMax,omitzero"`
	// Mean visual magnitude in M.
	VismagMean param.Opt[float64] `json:"vismagMean,omitzero"`
	// Minimum visual magnitude in M.
	VismagMin param.Opt[float64] `json:"vismagMin,omitzero"`
	// Array of the estimated mass of each deployable object, in kilograms. Must
	// contain the same number of elements as the value of numDeployable.
	DepEstMasses []float64 `json:"depEstMasses,omitzero"`
	// Array of the 1-sigma uncertainty of the mass for each deployable object, in
	// kilograms. Must contain the same number of elements as the value of
	// numDeployable.
	DepMassUncs []float64 `json:"depMassUncs,omitzero"`
	// Array of satellite deployable objects. Must contain the same number of elements
	// as the value of numDeployable.
	DepNames []string `json:"depNames,omitzero"`
	// Array of the type of missions the spacecraft performs. Must contain the same
	// number of elements as the value of numMission.
	MissionTypes []string `json:"missionTypes,omitzero"`
	paramObj
}

func (r OnorbitdetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitdetailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitdetailNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type OnorbitdetailNewParamsDataMode string

const (
	OnorbitdetailNewParamsDataModeReal      OnorbitdetailNewParamsDataMode = "REAL"
	OnorbitdetailNewParamsDataModeTest      OnorbitdetailNewParamsDataMode = "TEST"
	OnorbitdetailNewParamsDataModeSimulated OnorbitdetailNewParamsDataMode = "SIMULATED"
	OnorbitdetailNewParamsDataModeExercise  OnorbitdetailNewParamsDataMode = "EXERCISE"
)

type OnorbitdetailUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode OnorbitdetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// UUID of the parent Onorbit record.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Mass of fuel and disposables at launch time in kilograms.
	AdditionalMass param.Opt[float64] `json:"additionalMass,omitzero"`
	// The radius used for long-term debris environment projection analyses that is not
	// as conservative as COLA Radius, in meters.
	AdeptRadius param.Opt[float64] `json:"adeptRadius,omitzero"`
	// The total beginning of life delta V of the spacecraft, in meters per second.
	BolDeltaV param.Opt[float64] `json:"bolDeltaV,omitzero"`
	// Spacecraft beginning of life fuel mass, in orbit, in kilograms.
	BolFuelMass param.Opt[float64] `json:"bolFuelMass,omitzero"`
	// Average cross sectional area of the bus in meters squared.
	BusCrossSection param.Opt[float64] `json:"busCrossSection,omitzero"`
	// Type of the bus on the spacecraft.
	BusType param.Opt[string] `json:"busType,omitzero"`
	// Maximum dimension of the box circumscribing the spacecraft (d = sqrt(a*a + b*b +
	// c\*c) where a is the tip-to-tip dimension, b and c are perpendicular to that.)
	// in meters.
	ColaRadius param.Opt[float64] `json:"colaRadius,omitzero"`
	// Average cross sectional area in meters squared.
	CrossSection param.Opt[float64] `json:"crossSection,omitzero"`
	// The estimated total current mass of the spacecraft, in kilograms.
	CurrentMass param.Opt[float64] `json:"currentMass,omitzero"`
	// The 1-sigma uncertainty of the total spacecraft delta V, in meters per second.
	DeltaVUnc param.Opt[float64] `json:"deltaVUnc,omitzero"`
	// GEO drift rate, if applicable in degrees per day.
	DriftRate param.Opt[float64] `json:"driftRate,omitzero"`
	// Spacecraft dry mass (without fuel or disposables) in kilograms.
	DryMass param.Opt[float64] `json:"dryMass,omitzero"`
	// Estimated maximum burn duration for the object, in seconds.
	EstDeltaVDuration param.Opt[float64] `json:"estDeltaVDuration,omitzero"`
	// Estimated remaining fuel for the object in kilograms.
	FuelRemaining param.Opt[float64] `json:"fuelRemaining,omitzero"`
	// GEO slot if applicable, in degrees. -180 (West of Prime Meridian) to 180 degrees
	// (East of Prime Meridian). Prime Meridian is 0.
	GeoSlot param.Opt[float64] `json:"geoSlot,omitzero"`
	// The name of the source who last provided an observation for this idOnOrbit.
	LastObSource param.Opt[string] `json:"lastObSource,omitzero"`
	// Time of last reported observation for this object in ISO 8601 UTC with
	// microsecond precision.
	LastObTime param.Opt[time.Time] `json:"lastObTime,omitzero" format:"date-time"`
	// Nominal mass of spacecraft and fuel at launch time, in kilograms.
	LaunchMass param.Opt[float64] `json:"launchMass,omitzero"`
	// Maximum (estimated) mass of spacecraft and fuel at launch time, in kilograms.
	LaunchMassMax param.Opt[float64] `json:"launchMassMax,omitzero"`
	// Minimum (estimated) mass of spacecraft and fuel at launch time, in kilograms.
	LaunchMassMin param.Opt[float64] `json:"launchMassMin,omitzero"`
	// Boolean indicating whether a spacecraft is maneuverable. Note that a spacecraft
	// may have propulsion capability but may not be maneuverable due to lack of fuel,
	// anomalous condition, or other operational constraints.
	Maneuverable param.Opt[bool] `json:"maneuverable,omitzero"`
	// Maximum delta V available for this on-orbit spacecraft, in meters per second.
	MaxDeltaV param.Opt[float64] `json:"maxDeltaV,omitzero"`
	// Maximum dimension across the spacecraft (e.g., tip-to-tip across the solar panel
	// arrays) in meters.
	MaxRadius param.Opt[float64] `json:"maxRadius,omitzero"`
	// The number of sub-satellites or deployable objects on the spacecraft.
	NumDeployable param.Opt[int64] `json:"numDeployable,omitzero"`
	// The number of distinct missions the spacecraft performs.
	NumMission param.Opt[int64] `json:"numMission,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Current/latest radar cross section in meters squared.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// Maximum radar cross section in meters squared.
	RcsMax param.Opt[float64] `json:"rcsMax,omitzero"`
	// Mean radar cross section in meters squared.
	RcsMean param.Opt[float64] `json:"rcsMean,omitzero"`
	// Minimum radar cross section in meters squared.
	RcsMin param.Opt[float64] `json:"rcsMin,omitzero"`
	// The reference source, sources, or URL from which the data in this record was
	// obtained.
	RefSource param.Opt[string] `json:"refSource,omitzero"`
	// Spacecraft deployed area of solar array in meters squared.
	SolarArrayArea param.Opt[float64] `json:"solarArrayArea,omitzero"`
	// The 1-sigma uncertainty of the total spacecraft mass, in kilograms.
	TotalMassUnc param.Opt[float64] `json:"totalMassUnc,omitzero"`
	// Current/latest visual magnitude in M.
	Vismag param.Opt[float64] `json:"vismag,omitzero"`
	// Maximum visual magnitude in M.
	VismagMax param.Opt[float64] `json:"vismagMax,omitzero"`
	// Mean visual magnitude in M.
	VismagMean param.Opt[float64] `json:"vismagMean,omitzero"`
	// Minimum visual magnitude in M.
	VismagMin param.Opt[float64] `json:"vismagMin,omitzero"`
	// Array of the estimated mass of each deployable object, in kilograms. Must
	// contain the same number of elements as the value of numDeployable.
	DepEstMasses []float64 `json:"depEstMasses,omitzero"`
	// Array of the 1-sigma uncertainty of the mass for each deployable object, in
	// kilograms. Must contain the same number of elements as the value of
	// numDeployable.
	DepMassUncs []float64 `json:"depMassUncs,omitzero"`
	// Array of satellite deployable objects. Must contain the same number of elements
	// as the value of numDeployable.
	DepNames []string `json:"depNames,omitzero"`
	// Array of the type of missions the spacecraft performs. Must contain the same
	// number of elements as the value of numMission.
	MissionTypes []string `json:"missionTypes,omitzero"`
	paramObj
}

func (r OnorbitdetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitdetailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitdetailUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type OnorbitdetailUpdateParamsDataMode string

const (
	OnorbitdetailUpdateParamsDataModeReal      OnorbitdetailUpdateParamsDataMode = "REAL"
	OnorbitdetailUpdateParamsDataModeTest      OnorbitdetailUpdateParamsDataMode = "TEST"
	OnorbitdetailUpdateParamsDataModeSimulated OnorbitdetailUpdateParamsDataMode = "SIMULATED"
	OnorbitdetailUpdateParamsDataModeExercise  OnorbitdetailUpdateParamsDataMode = "EXERCISE"
)

type OnorbitdetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitdetailListParams]'s query parameters as
// `url.Values`.
func (r OnorbitdetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitdetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitdetailGetParams]'s query parameters as `url.Values`.
func (r OnorbitdetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
