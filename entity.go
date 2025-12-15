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

// EntityService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityService] method instead.
type EntityService struct {
	Options []option.RequestOption
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEntityService(opts ...option.RequestOption) (r EntityService) {
	r = EntityService{}
	r.Options = opts
	return
}

// Service operation to take a single Entity as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *EntityService) New(ctx context.Context, body EntityNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/entity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Entity record by its unique ID passed as a
// path parameter.
func (r *EntityService) Get(ctx context.Context, id string, query EntityGetParams, opts ...option.RequestOption) (res *shared.EntityFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/entity/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single Entity. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *EntityService) Update(ctx context.Context, id string, body EntityUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/entity/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EntityService) List(ctx context.Context, query EntityListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EntityAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/entity"
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
func (r *EntityService) ListAutoPaging(ctx context.Context, query EntityListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EntityAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an Entity object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *EntityService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/entity/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EntityService) Count(ctx context.Context, query EntityCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/entity/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves all distinct entity types.
func (r *EntityService) GetAllTypes(ctx context.Context, query EntityGetAllTypesParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/entity/getAllTypes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EntityService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *EntityQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/entity/queryhelp"
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
func (r *EntityService) Tuple(ctx context.Context, query EntityTupleParams, opts ...option.RequestOption) (res *[]shared.EntityFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/entity/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An entity is a generic representation of any object within a space/SSA system
// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
// entity can have an operating unit, a location (if terrestrial), and statuses.
type EntityAbridged struct {
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
	DataMode EntityAbridgedDataMode `json:"dataMode,required"`
	// Unique entity name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
	// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
	//
	// Any of "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT",
	// "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL".
	Type EntityAbridgedType `json:"type,required"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the record.
	IDEntity string `json:"idEntity"`
	// Unique identifier of the entity location, if terrestrial/fixed.
	IDLocation string `json:"idLocation"`
	// Onorbit identifier if this entity is part of an on-orbit object. For the public
	// catalog, the idOnOrbit is typically the satellite number as a string, but may be
	// a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the associated operating unit object.
	IDOperatingUnit string `json:"idOperatingUnit"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location shared.LocationAbridged `json:"location"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit EntityAbridgedOnOrbit `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Type of organization which owns this entity (e.g. Commercial, Government,
	// Academic, Consortium, etc).
	//
	// Any of "Commercial", "Government", "Academic", "Consortium", "Other".
	OwnerType EntityAbridgedOwnerType `json:"ownerType"`
	// Boolean indicating if this entity is taskable.
	Taskable bool `json:"taskable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDEntity              respjson.Field
		IDLocation            respjson.Field
		IDOnOrbit             respjson.Field
		IDOperatingUnit       respjson.Field
		Location              respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OwnerType             respjson.Field
		Taskable              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityAbridged) RawJSON() string { return r.JSON.raw }
func (r *EntityAbridged) UnmarshalJSON(data []byte) error {
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
type EntityAbridgedDataMode string

const (
	EntityAbridgedDataModeReal      EntityAbridgedDataMode = "REAL"
	EntityAbridgedDataModeTest      EntityAbridgedDataMode = "TEST"
	EntityAbridgedDataModeSimulated EntityAbridgedDataMode = "SIMULATED"
	EntityAbridgedDataModeExercise  EntityAbridgedDataMode = "EXERCISE"
)

// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
type EntityAbridgedType string

const (
	EntityAbridgedTypeAircraft     EntityAbridgedType = "AIRCRAFT"
	EntityAbridgedTypeBus          EntityAbridgedType = "BUS"
	EntityAbridgedTypeComm         EntityAbridgedType = "COMM"
	EntityAbridgedTypeIr           EntityAbridgedType = "IR"
	EntityAbridgedTypeLaseremitter EntityAbridgedType = "LASEREMITTER"
	EntityAbridgedTypeNavigation   EntityAbridgedType = "NAVIGATION"
	EntityAbridgedTypeOnorbit      EntityAbridgedType = "ONORBIT"
	EntityAbridgedTypeRfemitter    EntityAbridgedType = "RFEMITTER"
	EntityAbridgedTypeScientific   EntityAbridgedType = "SCIENTIFIC"
	EntityAbridgedTypeSensor       EntityAbridgedType = "SENSOR"
	EntityAbridgedTypeSite         EntityAbridgedType = "SITE"
	EntityAbridgedTypeVessel       EntityAbridgedType = "VESSEL"
)

// Model object representing on-orbit objects or satellites in the system.
type EntityAbridgedOnOrbit struct {
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
	DataMode string `json:"dataMode,required"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Alternate name of the on-orbit object.
	AltName string `json:"altName"`
	// Category of the on-orbit object. (Unknown, On-Orbit, Decayed, Cataloged Without
	// State, Launch Nominal, Analyst Satellite, Cislunar, Lunar, Hyperbolic,
	// Heliocentric, Interplanetary, Lagrangian, Docked).
	//
	// Any of "Unknown", "On-Orbit", "Decayed", "Cataloged Without State", "Launch
	// Nominal", "Analyst Satellite", "Cislunar", "Lunar", "Hyperbolic",
	// "Heliocentric", "Interplanetary", "Lagrangian", "Docked".
	Category string `json:"category"`
	// Common name of the on-orbit object.
	CommonName string `json:"commonName"`
	// Constellation to which this satellite belongs.
	Constellation string `json:"constellation"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Date of decay.
	DecayDate time.Time `json:"decayDate" format:"date-time"`
	// For the public catalog, the idOnOrbit is typically the satellite number as a
	// string, but may be a UUID for analyst or other unknown or untracked satellites,
	// auto-generated by the system.
	IDOnOrbit string `json:"idOnOrbit"`
	// International Designator, typically of the format YYYYLLLAAA, where YYYY is the
	// launch year, LLL is the sequential launch number of that year, and AAA is an
	// optional launch piece designator for the launch.
	IntlDes string `json:"intlDes"`
	// Date of launch.
	LaunchDate time.Time `json:"launchDate" format:"date"`
	// Id of the associated launchSite entity.
	LaunchSiteID string `json:"launchSiteId"`
	// Estimated lifetime of the on-orbit payload, if known.
	LifetimeYears int64 `json:"lifetimeYears"`
	// Mission number of the on-orbit object.
	MissionNumber string `json:"missionNumber"`
	// Type of on-orbit object: ROCKET BODY, DEBRIS, PAYLOAD, PLATFORM, MANNED,
	// UNKNOWN.
	//
	// Any of "ROCKET BODY", "DEBRIS", "PAYLOAD", "PLATFORM", "MANNED", "UNKNOWN".
	ObjectType string `json:"objectType"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		SatNo                 respjson.Field
		Source                respjson.Field
		AltName               respjson.Field
		Category              respjson.Field
		CommonName            respjson.Field
		Constellation         respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecayDate             respjson.Field
		IDOnOrbit             respjson.Field
		IntlDes               respjson.Field
		LaunchDate            respjson.Field
		LaunchSiteID          respjson.Field
		LifetimeYears         respjson.Field
		MissionNumber         respjson.Field
		ObjectType            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityAbridgedOnOrbit) RawJSON() string { return r.JSON.raw }
func (r *EntityAbridgedOnOrbit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of organization which owns this entity (e.g. Commercial, Government,
// Academic, Consortium, etc).
type EntityAbridgedOwnerType string

const (
	EntityAbridgedOwnerTypeCommercial EntityAbridgedOwnerType = "Commercial"
	EntityAbridgedOwnerTypeGovernment EntityAbridgedOwnerType = "Government"
	EntityAbridgedOwnerTypeAcademic   EntityAbridgedOwnerType = "Academic"
	EntityAbridgedOwnerTypeConsortium EntityAbridgedOwnerType = "Consortium"
	EntityAbridgedOwnerTypeOther      EntityAbridgedOwnerType = "Other"
)

// An entity is a generic representation of any object within a space/SSA system
// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
// entity can have an operating unit, a location (if terrestrial), and statuses.
//
// The properties ClassificationMarking, DataMode, Name, Source, Type are required.
type EntityIngestParam struct {
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
	DataMode EntityIngestDataMode `json:"dataMode,omitzero,required"`
	// Unique entity name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
	// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
	//
	// Any of "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT",
	// "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL".
	Type EntityIngestType `json:"type,omitzero,required"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Unique identifier of the record.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Unique identifier of the entity location, if terrestrial/fixed.
	IDLocation param.Opt[string] `json:"idLocation,omitzero"`
	// Onorbit identifier if this entity is part of an on-orbit object. For the public
	// catalog, the idOnOrbit is typically the satellite number as a string, but may be
	// a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the associated operating unit object.
	IDOperatingUnit param.Opt[string] `json:"idOperatingUnit,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Boolean indicating if this entity is taskable.
	Taskable param.Opt[bool] `json:"taskable,omitzero"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LocationIngestParam `json:"location,omitzero"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit EntityIngestOnOrbitParam `json:"onOrbit,omitzero"`
	// Type of organization which owns this entity (e.g. Commercial, Government,
	// Academic, Consortium, etc).
	//
	// Any of "Commercial", "Government", "Academic", "Consortium", "Other".
	OwnerType EntityIngestOwnerType `json:"ownerType,omitzero"`
	// List of URLs to additional details/documents for this entity.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

func (r EntityIngestParam) MarshalJSON() (data []byte, err error) {
	type shadow EntityIngestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityIngestParam) UnmarshalJSON(data []byte) error {
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
type EntityIngestDataMode string

const (
	EntityIngestDataModeReal      EntityIngestDataMode = "REAL"
	EntityIngestDataModeTest      EntityIngestDataMode = "TEST"
	EntityIngestDataModeSimulated EntityIngestDataMode = "SIMULATED"
	EntityIngestDataModeExercise  EntityIngestDataMode = "EXERCISE"
)

// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
type EntityIngestType string

const (
	EntityIngestTypeAircraft     EntityIngestType = "AIRCRAFT"
	EntityIngestTypeBus          EntityIngestType = "BUS"
	EntityIngestTypeComm         EntityIngestType = "COMM"
	EntityIngestTypeIr           EntityIngestType = "IR"
	EntityIngestTypeLaseremitter EntityIngestType = "LASEREMITTER"
	EntityIngestTypeNavigation   EntityIngestType = "NAVIGATION"
	EntityIngestTypeOnorbit      EntityIngestType = "ONORBIT"
	EntityIngestTypeRfemitter    EntityIngestType = "RFEMITTER"
	EntityIngestTypeScientific   EntityIngestType = "SCIENTIFIC"
	EntityIngestTypeSensor       EntityIngestType = "SENSOR"
	EntityIngestTypeSite         EntityIngestType = "SITE"
	EntityIngestTypeVessel       EntityIngestType = "VESSEL"
)

// Model object representing on-orbit objects or satellites in the system.
//
// The properties ClassificationMarking, DataMode, SatNo, Source are required.
type EntityIngestOnOrbitParam struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Alternate name of the on-orbit object.
	AltName param.Opt[string] `json:"altName,omitzero"`
	// Common name of the on-orbit object.
	CommonName param.Opt[string] `json:"commonName,omitzero"`
	// Constellation to which this satellite belongs.
	Constellation param.Opt[string] `json:"constellation,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Date of decay.
	DecayDate param.Opt[time.Time] `json:"decayDate,omitzero" format:"date-time"`
	// For the public catalog, the idOnOrbit is typically the satellite number as a
	// string, but may be a UUID for analyst or other unknown or untracked satellites,
	// auto-generated by the system.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// International Designator, typically of the format YYYYLLLAAA, where YYYY is the
	// launch year, LLL is the sequential launch number of that year, and AAA is an
	// optional launch piece designator for the launch.
	IntlDes param.Opt[string] `json:"intlDes,omitzero"`
	// Date of launch.
	LaunchDate param.Opt[time.Time] `json:"launchDate,omitzero" format:"date"`
	// Id of the associated launchSite entity.
	LaunchSiteID param.Opt[string] `json:"launchSiteId,omitzero"`
	// Estimated lifetime of the on-orbit payload, if known.
	LifetimeYears param.Opt[int64] `json:"lifetimeYears,omitzero"`
	// Mission number of the on-orbit object.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Category of the on-orbit object. (Unknown, On-Orbit, Decayed, Cataloged Without
	// State, Launch Nominal, Analyst Satellite, Cislunar, Lunar, Hyperbolic,
	// Heliocentric, Interplanetary, Lagrangian, Docked).
	//
	// Any of "Unknown", "On-Orbit", "Decayed", "Cataloged Without State", "Launch
	// Nominal", "Analyst Satellite", "Cislunar", "Lunar", "Hyperbolic",
	// "Heliocentric", "Interplanetary", "Lagrangian", "Docked".
	Category string `json:"category,omitzero"`
	// Type of on-orbit object: ROCKET BODY, DEBRIS, PAYLOAD, PLATFORM, MANNED,
	// UNKNOWN.
	//
	// Any of "ROCKET BODY", "DEBRIS", "PAYLOAD", "PLATFORM", "MANNED", "UNKNOWN".
	ObjectType string `json:"objectType,omitzero"`
	paramObj
}

func (r EntityIngestOnOrbitParam) MarshalJSON() (data []byte, err error) {
	type shadow EntityIngestOnOrbitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityIngestOnOrbitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EntityIngestOnOrbitParam](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[EntityIngestOnOrbitParam](
		"category", "Unknown", "On-Orbit", "Decayed", "Cataloged Without State", "Launch Nominal", "Analyst Satellite", "Cislunar", "Lunar", "Hyperbolic", "Heliocentric", "Interplanetary", "Lagrangian", "Docked",
	)
	apijson.RegisterFieldValidator[EntityIngestOnOrbitParam](
		"objectType", "ROCKET BODY", "DEBRIS", "PAYLOAD", "PLATFORM", "MANNED", "UNKNOWN",
	)
}

// Type of organization which owns this entity (e.g. Commercial, Government,
// Academic, Consortium, etc).
type EntityIngestOwnerType string

const (
	EntityIngestOwnerTypeCommercial EntityIngestOwnerType = "Commercial"
	EntityIngestOwnerTypeGovernment EntityIngestOwnerType = "Government"
	EntityIngestOwnerTypeAcademic   EntityIngestOwnerType = "Academic"
	EntityIngestOwnerTypeConsortium EntityIngestOwnerType = "Consortium"
	EntityIngestOwnerTypeOther      EntityIngestOwnerType = "Other"
)

type EntityQueryHelpResponse struct {
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
func (r EntityQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityNewParams struct {
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	EntityIngest EntityIngestParam
	paramObj
}

func (r EntityNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EntityIngest)
}
func (r *EntityNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.EntityIngest)
}

type EntityGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntityGetParams]'s query parameters as `url.Values`.
func (r EntityGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EntityUpdateParams struct {
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	EntityIngest EntityIngestParam
	paramObj
}

func (r EntityUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EntityIngest)
}
func (r *EntityUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.EntityIngest)
}

type EntityListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntityListParams]'s query parameters as `url.Values`.
func (r EntityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EntityCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntityCountParams]'s query parameters as `url.Values`.
func (r EntityCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EntityGetAllTypesParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntityGetAllTypesParams]'s query parameters as
// `url.Values`.
func (r EntityGetAllTypesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EntityTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntityTupleParams]'s query parameters as `url.Values`.
func (r EntityTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
