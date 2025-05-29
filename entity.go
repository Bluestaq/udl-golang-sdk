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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/entity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Entity record by its unique ID passed as a
// path parameter.
func (r *EntityService) Get(ctx context.Context, id string, query EntityGetParams, opts ...option.RequestOption) (res *EntityFull, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/entity/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves all distinct entity types.
func (r *EntityService) GetAllTypes(ctx context.Context, query EntityGetAllTypesParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/entity/getAllTypes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EntityService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *EntityQueryHelpResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *EntityService) Tuple(ctx context.Context, query EntityTupleParams, opts ...option.RequestOption) (res *[]EntityFull, err error) {
	opts = append(r.Options[:], opts...)
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
	Location EntityAbridgedLocation `json:"location"`
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

// Model representation of a location, which is a specific fixed point on the earth
// and is used to denote the locations of fixed sensors, operating units, etc.
type EntityAbridgedLocation struct {
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
	DataMode string `json:"dataMode,required"`
	// Location name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Altitude of the location, in kilometers.
	Altitude float64 `json:"altitude"`
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
	// Unique identifier of the location, auto-generated by the system.
	IDLocation string `json:"idLocation"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
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
		Name                  respjson.Field
		Source                respjson.Field
		Altitude              respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDLocation            respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityAbridgedLocation) RawJSON() string { return r.JSON.raw }
func (r *EntityAbridgedLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model object representing on-orbit objects or satellites in the system.
type EntityAbridgedOnOrbit struct {
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
type EntityFull struct {
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
	DataMode EntityFullDataMode `json:"dataMode,required"`
	// Unique entity name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
	// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
	//
	// Any of "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT",
	// "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL".
	Type EntityFullType `json:"type,required"`
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
	Location LocationFull `json:"location"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit EntityFullOnOrbit `json:"onOrbit"`
	// Model representation of a unit or organization which operates or controls a
	// space-related Entity such as an on-orbit payload, a sensor, etc. A contact may
	// belong to an organization.
	OperatingUnit EntityFullOperatingUnit `json:"operatingUnit"`
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
	OwnerType EntityFullOwnerType `json:"ownerType"`
	// Read-only collection of RF bands utilized by this entity for communication
	// and/or operation.
	RfBands []EntityFullRfBand `json:"rfBands"`
	// Read-only collection of statuses which can be collected by multiple sources.
	StatusCollection []EntityFullStatusCollection `json:"statusCollection"`
	// Boolean indicating if this entity is taskable.
	Taskable bool `json:"taskable"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// List of URLs to additional details/documents for this entity.
	URLs []string `json:"urls"`
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
		OperatingUnit         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OwnerType             respjson.Field
		RfBands               respjson.Field
		StatusCollection      respjson.Field
		Taskable              respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		URLs                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFull) RawJSON() string { return r.JSON.raw }
func (r *EntityFull) UnmarshalJSON(data []byte) error {
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
type EntityFullDataMode string

const (
	EntityFullDataModeReal      EntityFullDataMode = "REAL"
	EntityFullDataModeTest      EntityFullDataMode = "TEST"
	EntityFullDataModeSimulated EntityFullDataMode = "SIMULATED"
	EntityFullDataModeExercise  EntityFullDataMode = "EXERCISE"
)

// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
type EntityFullType string

const (
	EntityFullTypeAircraft     EntityFullType = "AIRCRAFT"
	EntityFullTypeBus          EntityFullType = "BUS"
	EntityFullTypeComm         EntityFullType = "COMM"
	EntityFullTypeIr           EntityFullType = "IR"
	EntityFullTypeLaseremitter EntityFullType = "LASEREMITTER"
	EntityFullTypeNavigation   EntityFullType = "NAVIGATION"
	EntityFullTypeOnorbit      EntityFullType = "ONORBIT"
	EntityFullTypeRfemitter    EntityFullType = "RFEMITTER"
	EntityFullTypeScientific   EntityFullType = "SCIENTIFIC"
	EntityFullTypeSensor       EntityFullType = "SENSOR"
	EntityFullTypeSite         EntityFullType = "SITE"
	EntityFullTypeVessel       EntityFullType = "VESSEL"
)

// Model object representing on-orbit objects or satellites in the system.
type EntityFullOnOrbit struct {
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
	DataMode string `json:"dataMode,required"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Alternate name of the on-orbit object.
	AltName string `json:"altName"`
	// Read-only collection of antennas on this on-orbit object.
	Antennas []EntityFullOnOrbitAntenna `json:"antennas"`
	// Read-only collection of batteries on this on-orbit object.
	Batteries []EntityFullOnOrbitBattery `json:"batteries"`
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
	// Read-only collection of details for this on-orbit object.
	OnorbitDetails []EntityFullOnOrbitOnorbitDetail `json:"onorbitDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only collection of solar arrays on this on-orbit object.
	SolarArrays []EntityFullOnOrbitSolarArray `json:"solarArrays"`
	// Read-only collection of thrusters (engines) on this on-orbit object.
	Thrusters []EntityFullOnOrbitThruster `json:"thrusters"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		SatNo                 respjson.Field
		Source                respjson.Field
		AltName               respjson.Field
		Antennas              respjson.Field
		Batteries             respjson.Field
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
		OnorbitDetails        respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SolarArrays           respjson.Field
		Thrusters             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullOnOrbit) RawJSON() string { return r.JSON.raw }
func (r *EntityFullOnOrbit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityFullOnOrbitAntenna struct {
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
	DataMode string `json:"dataMode,required"`
	// ID of the antenna.
	IDAntenna string `json:"idAntenna,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Model representation of information on on-orbit/spacecraft communication
	// antennas. A spacecraft may have multiple antennas and each antenna can have
	// multiple 'details' records compiled by different sources.
	Antenna AntennaFull `json:"antenna"`
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
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDAntenna             respjson.Field
		IDOnOrbit             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Antenna               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullOnOrbitAntenna) RawJSON() string { return r.JSON.raw }
func (r *EntityFullOnOrbitAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityFullOnOrbitBattery struct {
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
	DataMode string `json:"dataMode,required"`
	// ID of the battery.
	IDBattery string `json:"idBattery,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Model representation of specific spacecraft battery types.
	Battery BatteryFull `json:"battery"`
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
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
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
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullOnOrbitBattery) RawJSON() string { return r.JSON.raw }
func (r *EntityFullOnOrbitBattery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains details of the OnOrbit object.
type EntityFullOnOrbitOnorbitDetail struct {
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
	DataMode string `json:"dataMode,required"`
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
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
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
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Vismag                respjson.Field
		VismagMax             respjson.Field
		VismagMean            respjson.Field
		VismagMin             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullOnOrbitOnorbitDetail) RawJSON() string { return r.JSON.raw }
func (r *EntityFullOnOrbitOnorbitDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityFullOnOrbitSolarArray struct {
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
	DataMode string `json:"dataMode,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// ID of the SolarArray.
	IDSolarArray string `json:"idSolarArray,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
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
	// The number of solar arrays on the spacecraft of the type identified by
	// idSolarArray.
	Quantity int64 `json:"quantity"`
	// Model representation of information on on-orbit/spacecraft solar arrays. A
	// spacecraft may have multiple solar arrays and each solar array can have multiple
	// 'details' records compiled by different sources.
	SolarArray EntityFullOnOrbitSolarArraySolarArray `json:"solarArray"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOnOrbit             respjson.Field
		IDSolarArray          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Quantity              respjson.Field
		SolarArray            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullOnOrbitSolarArray) RawJSON() string { return r.JSON.raw }
func (r *EntityFullOnOrbitSolarArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of information on on-orbit/spacecraft solar arrays. A
// spacecraft may have multiple solar arrays and each solar array can have multiple
// 'details' records compiled by different sources.
type EntityFullOnOrbitSolarArraySolarArray struct {
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
	DataMode string `json:"dataMode,required"`
	// Solar Array name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
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
	// Read-only collection of additional SolarArrayDetails by various sources for this
	// organization, ignored on create/update. These details must be created separately
	// via the /udl/solararraydetails operations.
	SolarArrayDetails []SolarArrayDetailsFull `json:"solarArrayDetails"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataMode          respjson.Field
		Name              respjson.Field
		Source            respjson.Field
		ID                respjson.Field
		CreatedAt         respjson.Field
		CreatedBy         respjson.Field
		Origin            respjson.Field
		OrigNetwork       respjson.Field
		SolarArrayDetails respjson.Field
		UpdatedAt         respjson.Field
		UpdatedBy         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullOnOrbitSolarArraySolarArray) RawJSON() string { return r.JSON.raw }
func (r *EntityFullOnOrbitSolarArraySolarArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityFullOnOrbitThruster struct {
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
	DataMode string `json:"dataMode,required"`
	// ID of the Engine.
	IDEngine string `json:"idEngine,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Known launch vehicle engines and their performance characteristics and limits. A
	// launch vehicle has 1 to many engines per stage.
	Engine Engine `json:"engine"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The number of engines/thrusters on the spacecraft of the type identified by
	// idEngine.
	Quantity int64 `json:"quantity"`
	// The type of thruster associated with this record (e.g. LAE, Hydrazine REA,
	// etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDEngine              respjson.Field
		IDOnOrbit             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Engine                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Quantity              respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullOnOrbitThruster) RawJSON() string { return r.JSON.raw }
func (r *EntityFullOnOrbitThruster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a unit or organization which operates or controls a
// space-related Entity such as an on-orbit payload, a sensor, etc. A contact may
// belong to an organization.
type EntityFullOperatingUnit struct {
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
	DataMode string `json:"dataMode,required"`
	// Name of the operating unit.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea string `json:"airDefArea"`
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit owes its allegiance. This field will be set to "OTHR"
	// if the source value does not match a UDL country code value (ISO-3166-ALPHA-2).
	Allegiance string `json:"allegiance"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance string `json:"altAllegiance"`
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS. This field will be set
	// to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode string `json:"altCountryCode"`
	// Unique identifier of the operating unit record from the originating system.
	AltOperatingUnitID string `json:"altOperatingUnitId"`
	// Indicates the importance of the operating unit to the OES or MIR system. This
	// data element is restricted to update by DIA (DB-4). Valid values are: 0 - Does
	// not meet criteria above 1 - Primary importance to system 2 - Secondary
	// importance to system 3 - Tertiary importance to system O - Other. Explain in
	// Remarks.
	ClassRating string `json:"classRating"`
	// The physical manner of being or state of existence of the operating unit. A
	// physical condition that must be considered in the determining of a course of
	// action. The specific usage and enumerations contained in this field may be found
	// in the documentation provided in the referenceDoc field. If referenceDoc not
	// provided, users may consult the data provider.
	Condition string `json:"condition"`
	// Availability of the operating unit relative to its condition. Indicates the
	// reason the operating unit is not fully operational. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	ConditionAvail string `json:"conditionAvail"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc.
	//
	// Pos. 1-2. Latitude Degrees [00-90]
	//
	// Pos. 3-4. Latitude Minutes [00-59]
	//
	// Pos. 5-6. Latitude Seconds [00-59]
	//
	// Pos. 7-9. Latitude Thousandths Of Seconds [000-999]
	//
	// Pos. 10. Latitude Hemisphere [NS]
	//
	// Pos. 11-13. Longitude Degrees [00-180]
	//
	// Pos. 14-15. Longitude Minutes [00-59]
	//
	// Pos. 16-17. Longitude Seconds [00-59]
	//
	// Pos. 18-20. Longitude Thousandths Of Seconds [000-999]
	//
	// Pos. 21. Longitude Hemisphere [EW]
	//
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]
	Coord string `json:"coord"`
	// A mathematical model of the earth used to calculate coordinates on a map. US
	// Forces use the World Geodetic System 1984 (WGS 84), but also use maps by allied
	// countries with local datums. The datum must be specified to ensure accuracy of
	// coordinates. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	CoordDatum string `json:"coordDatum"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// coordinate.
	CoordDerivAcc float64 `json:"coordDerivAcc"`
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit geographic coordinates reside . This field will be set
	// to "OTHR" if the source value does not match a UDL country code value
	// (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A code describing the amount of operating unit participation in a deployment.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	DeployStatus string `json:"deployStatus"`
	// Description of the operating unit.
	Description string `json:"description"`
	// Combat status of a divisional or equivalent operating unit. Currently, this data
	// element applies only to operating units of the Former Soviet Union. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	DivCat string `json:"divCat"`
	// Organizational level of the operating unit. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Echelon string `json:"echelon"`
	// Indicates the major group or level to which an echelon belongs. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	EchelonTier string `json:"echelonTier"`
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum.
	ElevMsl float64 `json:"elevMsl"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl int64 `json:"elevMslConfLvl"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc float64 `json:"elevMslDerivAcc"`
	// The Intelligence Confidence Level or the Reliability/degree of confidence that
	// the analyst has assigned to the data within this record. The numerical range is
	// from 1 to 9 with 1 representing the highest confidence level.
	Eval int64 `json:"eval"`
	// The country code of the observed flag flown.
	FlagFlown string `json:"flagFlown"`
	// Naval fleet to which an operating unit is assigned. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FleetID string `json:"fleetId"`
	// An aggregation of military units within a single service (i.e., ARMY, AIR FORCE,
	// etc.) which operates under a single authority to accomplish a common mission.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Force string `json:"force"`
	// The specific name for a given force. For example, Force = ADF (Air Defense
	// Force) and Force Name = Army Air Defense Force.
	ForceName string `json:"forceName"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa string `json:"fpa"`
	// Principal combat-related role that an operating unit is organized, structured
	// and equipped to perform. Or, the specialized military or paramilitary branch in
	// which an individual serves, their specialization. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctRole string `json:"functRole"`
	// The distance between Mean Sea Level and a referenced ellipsoid.
	GeoidalMslSep float64 `json:"geoidalMslSep"`
	// Unique identifier of the contact for this operating unit.
	IDContact string `json:"idContact"`
	// Estimated identity of the Site (ASSUMED FRIEND, FRIEND, HOSTILE, FAKER, JOKER,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	Ident string `json:"ident"`
	// Unique identifier of the location record for this operating unit.
	IDLocation string `json:"idLocation"`
	// Unique identifier of the record, auto-generated by the system.
	IDOperatingUnit string `json:"idOperatingUnit"`
	// Unique identifier of the organization record for this operating unit.
	IDOrganization string `json:"idOrganization"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LocationFull `json:"location"`
	// Location name for the coordinates.
	LocName string `json:"locName"`
	// Indicates the reason that the operating unit is at that location. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason string `json:"locReason"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// This field contains a value indicating whether the record is a master unit
	// record (True) or a detail record (False). Master records contain basic
	// information that does not change over time for each unit that has been selected
	// to be projected.
	MasterUnit bool `json:"masterUnit"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts: 4Q (grid zone
	// designator, GZD) FJ (the 100,000-meter square identifier) 12345678 (numerical
	// location; easting is 1234 and northing is 5678, in this case specifying a
	// location with 10 m resolution).
	MilGrid string `json:"milGrid"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator
	MilGridSys string `json:"milGridSys"`
	// Indicates the principal type of mission that an operating unit is organized and
	// equipped to perform. The specific usage and enumerations contained in this field
	// may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimary string `json:"msnPrimary"`
	// Indicates the principal specialty type of mission that an operating unit is
	// organized and equipped to perform. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpecialty string `json:"msnPrimarySpecialty"`
	// Remarks contain amplifying information for a specific service. The information
	// may contain context and interpretations for consumer use.
	OperatingUnitRemarks []EntityFullOperatingUnitOperatingUnitRemark `json:"operatingUnitRemarks"`
	// The Degree to which an operating unit is ready to perform the overall
	// operational mission(s) for which it was organized and equipped. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	OperStatus string `json:"operStatus"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	Organization OrganizationFull `json:"organization"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv string `json:"polSubdiv"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Values are: A - Active I -
	// Inactive K - Acknowledged L - Local Q - A nominated (NOM) or Data Change Request
	// (DCR) record R - Production reduced by CMD decision W - Working Record.
	RecStatus string `json:"recStatus"`
	// The reference documentiation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc string `json:"referenceDoc"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd string `json:"resProd"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency. This date cannot be greater than the current
	// date.
	ReviewDate time.Time `json:"reviewDate" format:"date"`
	// This field contains a value indicating whether the record is a stylized
	// operating unit record (True) or a regular operating unit record (False). A
	// stylized operating unit is a type of operating unit with one set of equipment
	// that can be assigned to one or more superiors. A stylized operating unit is
	// generally useful for lower echelon operating units where the number of operating
	// units and types of equipment are equal for multiple organizations. In lieu of
	// creating unique operating unit records for each operating unit, a template is
	// created for the operating unit and its equipment. This template enables the user
	// to assign the operating unit to multiple organizations.
	StylizedUnit bool `json:"stylizedUnit"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode string `json:"symCode"`
	// An optional identifier for this operating unit that may be composed from items
	// such as the originating organization, allegiance, one-up number, etc.
	UnitIdentifier string `json:"unitIdentifier"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Universal Transverse Mercator (UTM) grid coordinates. Pos. 1-2, UTM Zone Column
	// [01-60 Pos. 3, UTM Zone Row [C-HJ-NP-X] Pos. 4, UTM False Easting [0-9] Pos.
	// 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9] Pos. 10-11, UTM False Northing
	// [0-9][0-9] Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm string `json:"utm"`
	// World Aeronautical Chart identifier for the area in which a designated operating
	// unit is located.
	Wac string `json:"wac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		AirDefArea            respjson.Field
		Allegiance            respjson.Field
		AltAllegiance         respjson.Field
		AltCountryCode        respjson.Field
		AltOperatingUnitID    respjson.Field
		ClassRating           respjson.Field
		Condition             respjson.Field
		ConditionAvail        respjson.Field
		Coord                 respjson.Field
		CoordDatum            respjson.Field
		CoordDerivAcc         respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeployStatus          respjson.Field
		Description           respjson.Field
		DivCat                respjson.Field
		Echelon               respjson.Field
		EchelonTier           respjson.Field
		ElevMsl               respjson.Field
		ElevMslConfLvl        respjson.Field
		ElevMslDerivAcc       respjson.Field
		Eval                  respjson.Field
		FlagFlown             respjson.Field
		FleetID               respjson.Field
		Force                 respjson.Field
		ForceName             respjson.Field
		Fpa                   respjson.Field
		FunctRole             respjson.Field
		GeoidalMslSep         respjson.Field
		IDContact             respjson.Field
		Ident                 respjson.Field
		IDLocation            respjson.Field
		IDOperatingUnit       respjson.Field
		IDOrganization        respjson.Field
		Lat                   respjson.Field
		Location              respjson.Field
		LocName               respjson.Field
		LocReason             respjson.Field
		Lon                   respjson.Field
		MasterUnit            respjson.Field
		MilGrid               respjson.Field
		MilGridSys            respjson.Field
		MsnPrimary            respjson.Field
		MsnPrimarySpecialty   respjson.Field
		OperatingUnitRemarks  respjson.Field
		OperStatus            respjson.Field
		Organization          respjson.Field
		Origin                respjson.Field
		PolSubdiv             respjson.Field
		RecStatus             respjson.Field
		ReferenceDoc          respjson.Field
		ResProd               respjson.Field
		ReviewDate            respjson.Field
		StylizedUnit          respjson.Field
		SymCode               respjson.Field
		UnitIdentifier        respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Utm                   respjson.Field
		Wac                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullOperatingUnit) RawJSON() string { return r.JSON.raw }
func (r *EntityFullOperatingUnit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
type EntityFullOperatingUnitOperatingUnitRemark struct {
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
	DataMode string `json:"dataMode,required"`
	// The ID of the operating unit to which this remark applies.
	IDOperatingUnit string `json:"idOperatingUnit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Unique identifier of the unit remark record from the originating system.
	AltRmkID string `json:"altRmkId"`
	// The remark type identifier. For example, the Mobility Air Forces (MAF) remark
	// code, defined in the Airfield Suitability and Restriction Report (ASRR).
	Code string `json:"code"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The name of the remark.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The remark type (e.g. Caution, Information, Misc, Restriction, etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOperatingUnit       respjson.Field
		Source                respjson.Field
		Text                  respjson.Field
		ID                    respjson.Field
		AltRmkID              respjson.Field
		Code                  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Name                  respjson.Field
		Origin                respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullOperatingUnitOperatingUnitRemark) RawJSON() string { return r.JSON.raw }
func (r *EntityFullOperatingUnitOperatingUnitRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of organization which owns this entity (e.g. Commercial, Government,
// Academic, Consortium, etc).
type EntityFullOwnerType string

const (
	EntityFullOwnerTypeCommercial EntityFullOwnerType = "Commercial"
	EntityFullOwnerTypeGovernment EntityFullOwnerType = "Government"
	EntityFullOwnerTypeAcademic   EntityFullOwnerType = "Academic"
	EntityFullOwnerTypeConsortium EntityFullOwnerType = "Consortium"
	EntityFullOwnerTypeOther      EntityFullOwnerType = "Other"
)

// Details on a particular Radio Frequency (RF) band, also known as a carrier,
// which may be in use by any type of Entity for communications or operations.
type EntityFullRfBand struct {
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
	DataMode string `json:"dataMode,required"`
	// Unique identifier of the parent Entity which uses this band.
	IDEntity string `json:"idEntity,required"`
	// RF Band name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band string `json:"band"`
	// RF Band frequency range bandwidth in Mhz.
	Bandwidth float64 `json:"bandwidth"`
	// Angle between the half-power (-3 dB) points of the main lobe of the antenna, in
	// degrees.
	Beamwidth float64 `json:"beamwidth"`
	// Center frequency of RF frequency range, if applicable, in Mhz.
	CenterFreq float64 `json:"centerFreq"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// RF Range edge gain, in dBi.
	EdgeGain float64 `json:"edgeGain"`
	// EIRP is defined as the RMS power input in decibel watts required to a lossless
	// half-wave dipole antenna to give the same maximum power density far from the
	// antenna as the actual transmitter. It is equal to the power input to the
	// transmitter's antenna multiplied by the antenna gain relative to a half-wave
	// dipole. Effective radiated power and effective isotropic radiated power both
	// measure the amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Eirp float64 `json:"eirp"`
	// Effective Radiated Power (ERP) is the total power in decibel watts radiated by
	// an actual antenna relative to a half-wave dipole rather than a theoretical
	// isotropic antenna. A half-wave dipole has a gain of 2.15 dB compared to an
	// isotropic antenna. EIRP(dB) = ERP (dB)+2.15 dB or EIRP (W) = 1.64\*ERP(W).
	// Effective radiated power and effective isotropic radiated power both measure the
	// amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Erp float64 `json:"erp"`
	// End/maximum of transmit RF frequency range, if applicable, in Mhz.
	FreqMax float64 `json:"freqMax"`
	// Start/minimum of transmit RF frequency range, if applicable, in Mhz.
	FreqMin float64 `json:"freqMin"`
	// RF Band mode (e.g. TX, RX).
	//
	// Any of "TX", "RX".
	Mode string `json:"mode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// RF Range maximum gain, in dBi.
	PeakGain float64 `json:"peakGain"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	Polarization string `json:"polarization"`
	// Purpose or use of the RF Band -- COMM = communications, TTC =
	// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
	//
	// Any of "COMM", "TTC", "OPS", "OTHER".
	Purpose string `json:"purpose"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDEntity              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Band                  respjson.Field
		Bandwidth             respjson.Field
		Beamwidth             respjson.Field
		CenterFreq            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EdgeGain              respjson.Field
		Eirp                  respjson.Field
		Erp                   respjson.Field
		FreqMax               respjson.Field
		FreqMin               respjson.Field
		Mode                  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PeakGain              respjson.Field
		Polarization          respjson.Field
		Purpose               respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullRfBand) RawJSON() string { return r.JSON.raw }
func (r *EntityFullRfBand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status for a particular Entity. An entity may have multiple status records
// collected by various sources.
type EntityFullStatusCollection struct {
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
	DataMode string `json:"dataMode,required"`
	// Unique identifier of the parent entity.
	IDEntity string `json:"idEntity,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
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
	// Comments describing the status creation and or updates to an entity.
	Notes string `json:"notes"`
	// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	OpsCap string `json:"opsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
	// ACTIVE, STANDBY).
	//
	// Any of "UNKNOWN", "DEAD", "ACTIVE", "RF ACTIVE", "STANDBY".
	State               string                                          `json:"state"`
	SubStatusCollection []EntityFullStatusCollectionSubStatusCollection `json:"subStatusCollection"`
	// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	SysCap string `json:"sysCap"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		IDEntity               respjson.Field
		Source                 respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DeclassificationDate   respjson.Field
		DeclassificationString respjson.Field
		DerivedFrom            respjson.Field
		Notes                  respjson.Field
		OpsCap                 respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		State                  respjson.Field
		SubStatusCollection    respjson.Field
		SysCap                 respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullStatusCollection) RawJSON() string { return r.JSON.raw }
func (r *EntityFullStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional sub-system or capability status for the parent entity.
type EntityFullStatusCollectionSubStatusCollection struct {
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
	DataMode string `json:"dataMode,required"`
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status string `json:"status,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
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
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Notes                 respjson.Field
		Source                respjson.Field
		Status                respjson.Field
		StatusID              respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityFullStatusCollectionSubStatusCollection) RawJSON() string { return r.JSON.raw }
func (r *EntityFullStatusCollectionSubStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An entity is a generic representation of any object within a space/SSA system
// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
// entity can have an operating unit, a location (if terrestrial), and statuses.
//
// The properties ClassificationMarking, DataMode, Name, Source, Type are required.
type EntityIngestParam struct {
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
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
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
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
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
	AodrSupported         bool                               `json:"aodrSupported"`
	ClassificationMarking string                             `json:"classificationMarking"`
	Description           string                             `json:"description"`
	HistorySupported      bool                               `json:"historySupported"`
	Name                  string                             `json:"name"`
	Parameters            []EntityQueryHelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                           `json:"requiredRoles"`
	RestSupported         bool                               `json:"restSupported"`
	SortSupported         bool                               `json:"sortSupported"`
	TypeName              string                             `json:"typeName"`
	Uri                   string                             `json:"uri"`
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

type EntityQueryHelpResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityQueryHelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *EntityQueryHelpResponseParameter) UnmarshalJSON(data []byte) error {
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
	return json.Marshal(r.EntityIngest)
}
func (r *EntityNewParams) UnmarshalJSON(data []byte) error {
	return r.EntityIngest.UnmarshalJSON(data)
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
	return json.Marshal(r.EntityIngest)
}
func (r *EntityUpdateParams) UnmarshalJSON(data []byte) error {
	return r.EntityIngest.UnmarshalJSON(data)
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
