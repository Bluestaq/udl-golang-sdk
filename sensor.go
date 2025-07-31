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

// SensorService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSensorService] method instead.
type SensorService struct {
	Options     []option.RequestOption
	Calibration SensorCalibrationService
}

// NewSensorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSensorService(opts ...option.RequestOption) (r SensorService) {
	r = SensorService{}
	r.Options = opts
	r.Calibration = NewSensorCalibrationService(opts...)
	return
}

// Service operation to take a single sensor as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *SensorService) New(ctx context.Context, body SensorNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Sensor. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *SensorService) Update(ctx context.Context, id string, body SensorUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensor/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensorService) List(ctx context.Context, query SensorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SensorListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sensor"
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
func (r *SensorService) ListAutoPaging(ctx context.Context, query SensorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SensorListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Sensor specified by the passed ID path parameter.
// A specific role is required to perform this service operation. Please contact
// the UDL team for assistance.
func (r *SensorService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensor/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SensorService) Count(ctx context.Context, query SensorCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sensor/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Sensor by its unique ID passed as a path
// parameter.
func (r *SensorService) Get(ctx context.Context, id string, query SensorGetParams, opts ...option.RequestOption) (res *SensorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensor/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SensorService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SensorQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensor/queryhelp"
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
func (r *SensorService) Tuple(ctx context.Context, query SensorTupleParams, opts ...option.RequestOption) (res *[]SensorTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensor/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of observation data for electro-optical based sensor
// phenomenologies.
type SensorListResponse struct {
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
	DataMode SensorListResponseDataMode `json:"dataMode,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Optional flag indicating if the sensor is active.
	Active bool `json:"active"`
	// Optional US Air Force identifier for the sensor/ASR site, typically for air
	// surveillance radar (ASR) sensors.
	AfID string `json:"afId"`
	// The sensor type at the site. Optional field, intended primarily for ASRs.
	AsrType string `json:"asrType"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional dissemination control required for accessing data (e.g observations)
	// produced by this sensor. This is typically a proprietary data owner control for
	// commercial sensors.
	DataControl string `json:"dataControl"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity SensorListResponseEntity `json:"entity"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity string `json:"idEntity"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensor string `json:"idSensor"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Collection of Sensorcharacteristics which define characteristics and
	// capabilities of a sensor.
	Sensorcharacteristics []SensorListResponseSensorcharacteristic `json:"sensorcharacteristics"`
	// Sensorlimits define 0 to many limits of a particular sensor in terms of
	// observation coverage of on-orbit objects.
	SensorlimitsCollection []SensorListResponseSensorlimitsCollection `json:"sensorlimitsCollection"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber          int64                                   `json:"sensorNumber"`
	SensorObservationType SensorListResponseSensorObservationType `json:"sensorObservationType"`
	// Collection of SensorStats which contain statistics of a sensor.
	SensorStats []SensorListResponseSensorStat `json:"sensorStats"`
	SensorType  SensorListResponseSensorType   `json:"sensorType"`
	// Optional short name for the sensor.
	ShortName string `json:"shortName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		SensorName             respjson.Field
		Source                 respjson.Field
		Active                 respjson.Field
		AfID                   respjson.Field
		AsrType                respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DataControl            respjson.Field
		Entity                 respjson.Field
		IDEntity               respjson.Field
		IDSensor               respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		Sensorcharacteristics  respjson.Field
		SensorlimitsCollection respjson.Field
		SensorNumber           respjson.Field
		SensorObservationType  respjson.Field
		SensorStats            respjson.Field
		SensorType             respjson.Field
		ShortName              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorListResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorListResponse) UnmarshalJSON(data []byte) error {
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
type SensorListResponseDataMode string

const (
	SensorListResponseDataModeReal      SensorListResponseDataMode = "REAL"
	SensorListResponseDataModeTest      SensorListResponseDataMode = "TEST"
	SensorListResponseDataModeSimulated SensorListResponseDataMode = "SIMULATED"
	SensorListResponseDataModeExercise  SensorListResponseDataMode = "EXERCISE"
)

// An entity is a generic representation of any object within a space/SSA system
// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
// entity can have an operating unit, a location (if terrestrial), and statuses.
type SensorListResponseEntity struct {
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
	// Unique entity name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
	// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
	//
	// Any of "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT",
	// "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL".
	Type string `json:"type,required"`
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
	Location SensorListResponseEntityLocation `json:"location"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit SensorListResponseEntityOnOrbit `json:"onOrbit"`
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
	OwnerType string `json:"ownerType"`
	// Boolean indicating if this entity is taskable.
	Taskable bool `json:"taskable"`
	// Terrestrial identifier of this entity, if applicable.
	TerrestrialID string `json:"terrestrialId"`
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
		TerrestrialID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *SensorListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a location, which is a specific fixed point on the earth
// and is used to denote the locations of fixed sensors, operating units, etc.
type SensorListResponseEntityLocation struct {
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
func (r SensorListResponseEntityLocation) RawJSON() string { return r.JSON.raw }
func (r *SensorListResponseEntityLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model object representing on-orbit objects or satellites in the system.
type SensorListResponseEntityOnOrbit struct {
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
func (r SensorListResponseEntityOnOrbit) RawJSON() string { return r.JSON.raw }
func (r *SensorListResponseEntityOnOrbit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of characteristics and capabilities of a sensor.
type SensorListResponseSensorcharacteristic struct {
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
	// Unique identifier of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of measurement range(s) where radar samples must fall to be acceptable. If
	// this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	AcceptSampleRanges []float64 `json:"acceptSampleRanges"`
	// Number of bits used in the conversion from analog electrons in a pixel well to a
	// digital number. The digital number has a maximum value of 2^N, where N is the
	// number of bits.
	AnalogToDigitalBitSize int64 `json:"analogToDigitalBitSize"`
	// Optical sensor camera aperture.
	Aperture float64 `json:"aperture"`
	// For ASR (Air Surveillance Radar) sensors, the scan (360 deg sweep) rate of the
	// radar, in scans/minute.
	AsrScanRate float64 `json:"asrScanRate"`
	// One-way radar receiver loss factor due to atmospheric effects. This value will
	// often be the same as the corresponding transmission factor but may be different
	// for bistatic systems.
	AtmosReceiverLoss float64 `json:"atmosReceiverLoss"`
	// One-way radar transmission loss factor due to atmospheric effects.
	AtmosTransmissionLoss float64 `json:"atmosTransmissionLoss"`
	// Average atmospheric angular width with no distortion from turbulence at an
	// optical sensor site in arcseconds.
	AvgAtmosSeeingConditions float64 `json:"avgAtmosSeeingConditions"`
	// Array of azimuth angles of a radar beam, in degrees. If this field is populated,
	// the associated beam(s) must be provided in the beamOrder field.
	AzAngs []float64 `json:"azAngs"`
	// Azimuth rate acquisition limit (radians/minute).
	AzimuthRate float64 `json:"azimuthRate"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of watts per square meter per steradian
	// (W/(m^2 str)) consistent with sensor detection bands.
	BackgroundSkyRadiance float64 `json:"backgroundSkyRadiance"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of visual magnitude consistent with sensor
	// detection bands.
	BackgroundSkyVisMag float64 `json:"backgroundSkyVisMag"`
	// Sensor band.
	Band string `json:"band"`
	// The total bandwidth, in megahertz, about the radar center frequency.
	Bandwidth float64 `json:"bandwidth"`
	// Array designating the beam order of provided values (e.g. vb1 for vertical beam
	// 1, ob1 for oblique beam 1, etc.). Required if any of the following array fields
	// are populated: azAngs, elAngs, radarPulseWidths, pulseRepPeriods, delayGates,
	// rangeGates, rangeSpacings, vertGateSpacings, vertGateWidths, specAvgSpectraNums,
	// tempMedFiltCodes, runMeanCodes, totRecNums, reqRecords, acceptSampleRanges.
	BeamOrder []string `json:"beamOrder"`
	// Number of radar beams used by the sensor.
	BeamQty int64 `json:"beamQty"`
	// The angle of the center of a phased array sensor.
	Boresight float64 `json:"boresight"`
	// The number of degrees off of the boresight for the sensor.
	BoresightOffAngle float64 `json:"boresightOffAngle"`
	// Weighted center wavelength for an optical sensor bandpass in micrometers. It is
	// the center wavelength in a weighted integral sense, accounting for the
	// sensitivity vs. wavelength curve for the sensor focal plane array.
	CenterWavelength float64 `json:"centerWavelength"`
	// Collapsing loss in decibels. Collapsing losses occur when two or more sources of
	// noise are added to the target signal. Examples include receiver bandwidth
	// mismatch with filtering bandwidth and elevation or azimuth beam collapse onto
	// position/height indicator displays.
	CollapsingLoss float64 `json:"collapsingLoss"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Threshold shear value beyond which one of the radial velocity values will be
	// rejected, measured in units of inverse second.
	CritShear float64 `json:"critShear"`
	// Current flowing through optical sensor focal plane electronics with a closed
	// shutter in electrons per second.
	DarkCurrent float64 `json:"darkCurrent"`
	// Array of time delay(s) for pulses from a radar beam to get to the first range
	// gate, in nanoseconds. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	DelayGates []float64 `json:"delayGates"`
	// Description of the equipment and data source.
	Description string `json:"description"`
	// Detection signal-to-noise ratio (SNR) threshold in decibels. This value is
	// typically lower than trackSNR.
	DetectSnr float64 `json:"detectSNR"`
	// Sensor duty cycle as a fraction of 1. Duty cycle is the fraction of time a
	// sensor is actively transmitting.
	DutyCycle float64 `json:"dutyCycle"`
	// Sensor Earth limb exclusion height in kilometers and is generally only applied
	// to space-based sensors. Some models used an earth exclusion angle instead, but
	// this assumes the sensor is in a circular orbit with constant altitude relative
	// to the earth. The limb exclusion height can be used for space-based sensors in
	// any orbit (assuming it is constant with sensor altitude). The limb height is
	// defined to be 0 at the surface of the earth.
	EarthLimbExclHgt float64 `json:"earthLimbExclHgt"`
	// Array of elevation angles of a radar beam, in degrees. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	ElAngs []float64 `json:"elAngs"`
	// Elevation rate acquisition limit (radians/minute).
	ElevationRateGeolm float64 `json:"elevationRateGeolm"`
	// Type of equipment used to take measurements.
	EquipmentType string `json:"equipmentType"`
	// The beam width of a Sensor's Fan (range). The values for this range from (0.0 to
	// PI).
	FanBeamWidth float64 `json:"fanBeamWidth"`
	// Number of Fast Fourier Transform (FFT) points used to convert time varying
	// signals into the frequency domain.
	Fft int64 `json:"fft"`
	// Maximum number of times the first guess was propagated in each gate before
	// failing the first guess check.
	FgpCrit int64 `json:"fgpCrit"`
	// Noise term, in decibels, that arises when a radar receiver filter has a
	// non-optimal bandwidth for an incoming signal (i.e., when it does not match the
	// pulse width).
	FilterMismatchFactor float64 `json:"filterMismatchFactor"`
	// F-number for an optical telescope. It is dimensionless and is defined as the
	// ratio of the focal length to the aperture width.
	FNum float64 `json:"fNum"`
	// For radar based sensors, the focal point elevation of the radar at the site, in
	// meters.
	FocalPoint float64 `json:"focalPoint"`
	// Horizontal field of view, in degrees.
	HFov float64 `json:"hFOV"`
	// Horizontal pixel resolution.
	HResPixels int64 `json:"hResPixels"`
	// For radar based sensors, K-factor is a relative indicator of refractivity that
	// infers the amount of radar beam bending due to atmosphere. (1<K<2).
	K float64 `json:"k"`
	// For Orbiting Sensors, First Card Azimuth limit #1 (left, degrees).
	LeftClockAngle float64 `json:"leftClockAngle"`
	// Leftmost GEO belt longitude limit for this sensor (if applicable).
	LeftGeoBeltLimit float64 `json:"leftGeoBeltLimit"`
	// Site where measurement is taken.
	Location string `json:"location"`
	// Aggregated radar range equation gain in decibels for maximum sensitivity. It is
	// a roll-up value for low fidelity modeling and is often the only sensitivity
	// value available for a radar system without access to detailed performance
	// parameters.
	LoopGain float64 `json:"loopGain"`
	// Lowest aspect angle of the full moon in degrees at which the sensor can achieve
	// full performance.
	LunarExclAngle float64 `json:"lunarExclAngle"`
	// Angle between magnetic north and true north at the sensor site, in degrees.
	MagDec float64 `json:"magDec"`
	// Absolute magnitude acquisition limit for optical sensors.
	MagnitudeLimit float64 `json:"magnitudeLimit"`
	// Max deviation angle of the sensor in degrees.
	MaxDeviationAngle float64 `json:"maxDeviationAngle"`
	// Maximum integration time per image frame in seconds for an optical sensor.
	MaxIntegrationTime float64 `json:"maxIntegrationTime"`
	// Maximum observable sensor range, in kilometers.
	MaxObservableRange float64 `json:"maxObservableRange"`
	// Maximum observable range limit in kilometers -- sensor cannot acquire beyond
	// this range.
	MaxRangeLimit float64 `json:"maxRangeLimit"`
	// Maximum wavelength detectable by an optical sensor in micrometers.
	MaxWavelength float64 `json:"maxWavelength"`
	// Minimum integration time per image frame in seconds for an optical sensor.
	MinIntegrationTime float64 `json:"minIntegrationTime"`
	// Minimum range measurement capability of the sensor, in kilometers.
	MinRangeLimit float64 `json:"minRangeLimit"`
	// Signal to Noise Ratio, in decibels. The values for this range from 0.0 - + 99.99
	// dB.
	MinSignalNoiseRatio float64 `json:"minSignalNoiseRatio"`
	// Minimum wavelength detectable by an optical sensor in micrometers.
	MinWavelength float64 `json:"minWavelength"`
	// Negative Range-rate/relative velocity limit (kilometers/second).
	NegativeRangeRateLimit float64 `json:"negativeRangeRateLimit"`
	// Noise figure for a radar system in decibels. This value may be used to compute
	// system noise when the system temperature is unavailable.
	NoiseFigure float64 `json:"noiseFigure"`
	// Number of pulses that are non-coherently integrated during detection processing.
	NonCoherentIntegratedPulses int64 `json:"nonCoherentIntegratedPulses"`
	// For radar based sensors, number of integrated pulses in a transmit cycle.
	NumIntegratedPulses int64 `json:"numIntegratedPulses"`
	// Number of integration frames for an optical sensor.
	NumIntegrationFrames int64 `json:"numIntegrationFrames"`
	// The number of optical integration mode characteristics provided in this record.
	// If provided, the numOpticalIntegrationModes value indicates the number of
	// elements in each of the opticalIntegrationTimes, opticalIntegrationAngularRates,
	// opticalIntegrationFrames, opticalIntegrationPixelBinnings, and
	// opticalIntegrationSNRs arrays.
	NumOpticalIntegrationModes int64 `json:"numOpticalIntegrationModes"`
	// The number of waveforms characteristics provided in this record. If provided,
	// the numWaveforms value indicates the number of elements in each of the
	// waveformPulseWidths, waveformBandWidths, waveformMinRanges, waveformMaxRanges,
	// and waveformLoopGains arrays.
	NumWaveforms int64 `json:"numWaveforms"`
	// Array containing the angular rate, in arcsec/sec, for each provided optical
	// integration mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationAngularRates []float64 `json:"opticalIntegrationAngularRates"`
	// Array containing the number of frames, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationFrames []float64 `json:"opticalIntegrationFrames"`
	// Array containing the pixel binning, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationPixelBinnings []float64 `json:"opticalIntegrationPixelBinnings"`
	// Array of optical integration modes signal to noise ratios. The number of
	// elements must be equal to the value indicated in numOpticalIntegrationModes.
	OpticalIntegrationSnRs []float64 `json:"opticalIntegrationSNRs"`
	// Array containing the time, in seconds, for each provided optical integration
	// mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationTimes []float64 `json:"opticalIntegrationTimes"`
	// Fraction of incident light transmitted to an optical sensor focal plane array.
	// The value is given as a fraction of 1, not as a percent.
	OpticalTransmission float64 `json:"opticalTransmission"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Two-way pattern absorption/propagation loss due to medium in decibels.
	PatternAbsorptionLoss float64 `json:"patternAbsorptionLoss"`
	// Losses from the beam shape, scanning, and pattern factor in decibels. These
	// losses occur when targets are not directly in line with a beam center. For space
	// surveillance, this will occur most often during sensor scanning.
	PatternScanLoss float64 `json:"patternScanLoss"`
	// Maximum instantaneous radar transmit power in watts for use in the radar range
	// equation.
	PeakPower float64 `json:"peakPower"`
	// Angular field-of-view covered by one pixel in a focal plane array in
	// microradians. The pixel is assumed to be a perfect square so that only a single
	// value is required.
	PixelInstantaneousFov float64 `json:"pixelInstantaneousFOV"`
	// Maximum number of electrons that can be collected in a single pixel on an
	// optical sensor focal plane array.
	PixelWellDepth int64 `json:"pixelWellDepth"`
	// Positive Range-rate/relative velocity limit (kilometers/second).
	PositiveRangeRateLimit float64 `json:"positiveRangeRateLimit"`
	// For radar based sensors, pulse repetition frequency (PRF), in hertz. Number of
	// new pulses transmitted per second.
	Prf float64 `json:"prf"`
	// Designated probability of detection at the signal-to-noise detection threshold.
	ProbDetectSnr float64 `json:"probDetectSNR"`
	// For radar based sensors, probability of the indication of the presence of a
	// radar target due to noise or interference.
	ProbFalseAlarm float64 `json:"probFalseAlarm"`
	// Array of interval(s) between the start of one radar pulse and the start of
	// another for a radar beam, in microseconds. If this field is populated, the
	// associated beam(s) must be provided in the beamOrder field.
	PulseRepPeriods []float64 `json:"pulseRepPeriods"`
	// Fraction of incident photons converted to electrons at the focal plane array.
	// This value is a decimal number between 0 and 1, inclusive.
	QuantumEff float64 `json:"quantumEff"`
	// Radar frequency in hertz, of the sensor (if a radar sensor).
	RadarFrequency float64 `json:"radarFrequency"`
	// Message data format transmitted by the sensor digitizer.
	RadarMessageFormat string `json:"radarMessageFormat"`
	// For radar based sensors, radar maximum unambiguous range, in kilometers.
	RadarMur float64 `json:"radarMUR"`
	// Array of transmit time(s) for a radar beam pulse, in microseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RadarPulseWidths []float64 `json:"radarPulseWidths"`
	// Radio frequency (if sensor is RF).
	RadioFrequency float64 `json:"radioFrequency"`
	// Losses due to the presence of a protective radome surrounding a radar sensor, in
	// decibels.
	RadomeLoss float64 `json:"radomeLoss"`
	// Array of the number(s) of discrete altitudes where return signals are sampled by
	// a radar beam. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	RangeGates []int64 `json:"rangeGates"`
	// Array of range gate spacing(s) for a radar beam, in nanoseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RangeSpacings []float64 `json:"rangeSpacings"`
	// Number of false-signal electrons generated by optical sensor focal plane
	// read-out electronics from photon-to-electron conversion during frame
	// integration. The units are in electrons RMS.
	ReadNoise int64 `json:"readNoise"`
	// Radar receive gain in decibels.
	ReceiveGain float64 `json:"receiveGain"`
	// Horizontal/azimuthal receive beamwidth for a radar in degrees.
	ReceiveHorizBeamWidth float64 `json:"receiveHorizBeamWidth"`
	// Aggregate radar receive loss, in decibels.
	ReceiveLoss float64 `json:"receiveLoss"`
	// Vertical/elevation receive beamwidth for a radar in degrees.
	ReceiveVertBeamWidth float64 `json:"receiveVertBeamWidth"`
	// Reference temperature for radar noise in Kelvin. A reference temperature is used
	// when the radar system temperature is unknown and is combined with the system
	// noise figure to estimate signal loss.
	RefTemp float64 `json:"refTemp"`
	// Array of the total number(s) of records required to meet consensus for a radar
	// beam. If this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	ReqRecords []int64 `json:"reqRecords"`
	// For Orbiting Sensors, First Card Azimuth limit #3 (right, degrees).
	RightClockAngle float64 `json:"rightClockAngle"`
	// Rightmost GEO belt longitude limit for this sensor (if applicable).
	RightGeoBeltLimit float64 `json:"rightGeoBeltLimit"`
	// Array of running mean code(s) used by radar data processing. The running mean
	// method involves taking a series of averages of different selections of the full
	// data set to smooth out short-term fluctuations in the data. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	RunMeanCodes []int64 `json:"runMeanCodes"`
	// Radar signal processing losses, in decibels.
	SignalProcessingLoss float64 `json:"signalProcessingLoss"`
	// Site code of the sensor.
	SiteCode string `json:"siteCode"`
	// Sensor and target position vector origins are at the center of the earth. The
	// sun vector origin is at the target position and points toward the sun. Any value
	// between 0 and 180 degrees is acceptable and is assumed to apply in both
	// directions (i.e., a solar exclusion angle of 30 degrees is understood to mean no
	// viewing for any angle between -30 deg and +30 deg).
	SolarExclAngle float64 `json:"solarExclAngle"`
	// Array of the number(s) of Doppler spectra used to process measurements from
	// radar. Spectral averaging involves combining multiple Doppler spectra acquired
	// to obtain a more accurate and representative spectrum. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	SpecAvgSpectraNums []int64 `json:"specAvgSpectraNums"`
	// For radar based sensors, expression of the radar system noise, aggregated as an
	// equivalent thermal noise value, in degrees Kelvin.
	SystemNoiseTemperature float64 `json:"systemNoiseTemperature"`
	// Maximum taskable range of the sensor, in kilometers.
	TaskableRange float64 `json:"taskableRange"`
	// Array of temporal median filter code(s) of a radar beam. Temporal median
	// filtering is a noise-reducing algorithm which involves replacing each data point
	// with the median value of a window of neighboring points over time. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	TempMedFiltCodes []int64 `json:"tempMedFiltCodes"`
	// Test number for the observed measurement.
	TestNumber string `json:"testNumber"`
	// Array of the total number(s) of records for a radar beam. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	TotRecNums []int64 `json:"totRecNums"`
	// For tower sensors, the physical height of the sensor tower, in meters.
	TowerHeight float64 `json:"towerHeight"`
	// Beginning track angle limit, in radians. Track angle is the angle between the
	// camera axis and the gimbal plane. Values range from 0 - PI/2.
	TrackAngle float64 `json:"trackAngle"`
	// Track signal-to-noise ratio (SNR) threshold in decibels. This value is typically
	// higher than detectSNR.
	TrackSnr float64 `json:"trackSNR"`
	// Radar transmit gain in decibels.
	TransmitGain float64 `json:"transmitGain"`
	// Horizontal/azimuthal transmit beamwidth for a radar in degrees.
	TransmitHorizBeamWidth float64 `json:"transmitHorizBeamWidth"`
	// Aggregate radar transmit loss, in decibels.
	TransmitLoss float64 `json:"transmitLoss"`
	// Radar transmit power in Watts.
	TransmitPower float64 `json:"transmitPower"`
	// Vertical/elevation transmit beamwidth for a radar in degrees.
	TransmitVertBeamWidth float64 `json:"transmitVertBeamWidth"`
	// True North correction for the sensor, in ACP (Azimunth Change Pulse) count.
	TrueNorthCorrector int64 `json:"trueNorthCorrector"`
	// Antenna true tilt, in degrees.
	TrueTilt float64 `json:"trueTilt"`
	// Twilight angle for ground-based optical sensors in degrees. A sensor cannot view
	// targets until the sun is below the twilight angle relative to the local horizon.
	// The sign of the angle is positive despite the sun elevation being negative after
	// local sunset. Typical values for the twilight angle are civil twilight (6
	// degrees), nautical twilight (12 degrees), and astronomical twilight (18
	// degrees).
	TwilightAngle float64 `json:"twilightAngle"`
	// Flag indicating if a vertical radar beam was used in the wind calculation.
	VertBeamFlag bool `json:"vertBeamFlag"`
	// Array of vertical distance(s) between points where radar measurements are taken,
	// in meters. If this field is populated, the associated beam(s) must be provided
	// in the beamOrder field.
	VertGateSpacings []float64 `json:"vertGateSpacings"`
	// Array of width(s) of each location where radar measurements are taken, in
	// meters. If this field is populated, the associated beam(s) must be provided in
	// the beamOrder field.
	VertGateWidths []float64 `json:"vertGateWidths"`
	// Vertical field of view, in degrees.
	VFov float64 `json:"vFOV"`
	// Vertical pixel resolution.
	VResPixels int64 `json:"vResPixels"`
	// Array containing the bandwidth, in megahertz, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformBandwidths []float64 `json:"waveformBandwidths"`
	// Array containing the loop gain, in decibels, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field (10 SNR vs. 1 dBsm at 1000 km).
	WaveformLoopGains []float64 `json:"waveformLoopGains"`
	// Array containing the maximum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMaxRanges []float64 `json:"waveformMaxRanges"`
	// Array containing the minimum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMinRanges []float64 `json:"waveformMinRanges"`
	// Array containing the pulse width, in microseconds, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformPulseWidths []float64 `json:"waveformPulseWidths"`
	// Peformance zone-1 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MaxRange float64 `json:"z1MaxRange"`
	// Peformance zone-1 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MinRange float64 `json:"z1MinRange"`
	// Peformance zone-2 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MaxRange float64 `json:"z2MaxRange"`
	// Peformance zone-2 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MinRange float64 `json:"z2MinRange"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking           respjson.Field
		DataMode                        respjson.Field
		IDSensor                        respjson.Field
		Source                          respjson.Field
		ID                              respjson.Field
		AcceptSampleRanges              respjson.Field
		AnalogToDigitalBitSize          respjson.Field
		Aperture                        respjson.Field
		AsrScanRate                     respjson.Field
		AtmosReceiverLoss               respjson.Field
		AtmosTransmissionLoss           respjson.Field
		AvgAtmosSeeingConditions        respjson.Field
		AzAngs                          respjson.Field
		AzimuthRate                     respjson.Field
		BackgroundSkyRadiance           respjson.Field
		BackgroundSkyVisMag             respjson.Field
		Band                            respjson.Field
		Bandwidth                       respjson.Field
		BeamOrder                       respjson.Field
		BeamQty                         respjson.Field
		Boresight                       respjson.Field
		BoresightOffAngle               respjson.Field
		CenterWavelength                respjson.Field
		CollapsingLoss                  respjson.Field
		CreatedAt                       respjson.Field
		CreatedBy                       respjson.Field
		CritShear                       respjson.Field
		DarkCurrent                     respjson.Field
		DelayGates                      respjson.Field
		Description                     respjson.Field
		DetectSnr                       respjson.Field
		DutyCycle                       respjson.Field
		EarthLimbExclHgt                respjson.Field
		ElAngs                          respjson.Field
		ElevationRateGeolm              respjson.Field
		EquipmentType                   respjson.Field
		FanBeamWidth                    respjson.Field
		Fft                             respjson.Field
		FgpCrit                         respjson.Field
		FilterMismatchFactor            respjson.Field
		FNum                            respjson.Field
		FocalPoint                      respjson.Field
		HFov                            respjson.Field
		HResPixels                      respjson.Field
		K                               respjson.Field
		LeftClockAngle                  respjson.Field
		LeftGeoBeltLimit                respjson.Field
		Location                        respjson.Field
		LoopGain                        respjson.Field
		LunarExclAngle                  respjson.Field
		MagDec                          respjson.Field
		MagnitudeLimit                  respjson.Field
		MaxDeviationAngle               respjson.Field
		MaxIntegrationTime              respjson.Field
		MaxObservableRange              respjson.Field
		MaxRangeLimit                   respjson.Field
		MaxWavelength                   respjson.Field
		MinIntegrationTime              respjson.Field
		MinRangeLimit                   respjson.Field
		MinSignalNoiseRatio             respjson.Field
		MinWavelength                   respjson.Field
		NegativeRangeRateLimit          respjson.Field
		NoiseFigure                     respjson.Field
		NonCoherentIntegratedPulses     respjson.Field
		NumIntegratedPulses             respjson.Field
		NumIntegrationFrames            respjson.Field
		NumOpticalIntegrationModes      respjson.Field
		NumWaveforms                    respjson.Field
		OpticalIntegrationAngularRates  respjson.Field
		OpticalIntegrationFrames        respjson.Field
		OpticalIntegrationPixelBinnings respjson.Field
		OpticalIntegrationSnRs          respjson.Field
		OpticalIntegrationTimes         respjson.Field
		OpticalTransmission             respjson.Field
		OrigNetwork                     respjson.Field
		PatternAbsorptionLoss           respjson.Field
		PatternScanLoss                 respjson.Field
		PeakPower                       respjson.Field
		PixelInstantaneousFov           respjson.Field
		PixelWellDepth                  respjson.Field
		PositiveRangeRateLimit          respjson.Field
		Prf                             respjson.Field
		ProbDetectSnr                   respjson.Field
		ProbFalseAlarm                  respjson.Field
		PulseRepPeriods                 respjson.Field
		QuantumEff                      respjson.Field
		RadarFrequency                  respjson.Field
		RadarMessageFormat              respjson.Field
		RadarMur                        respjson.Field
		RadarPulseWidths                respjson.Field
		RadioFrequency                  respjson.Field
		RadomeLoss                      respjson.Field
		RangeGates                      respjson.Field
		RangeSpacings                   respjson.Field
		ReadNoise                       respjson.Field
		ReceiveGain                     respjson.Field
		ReceiveHorizBeamWidth           respjson.Field
		ReceiveLoss                     respjson.Field
		ReceiveVertBeamWidth            respjson.Field
		RefTemp                         respjson.Field
		ReqRecords                      respjson.Field
		RightClockAngle                 respjson.Field
		RightGeoBeltLimit               respjson.Field
		RunMeanCodes                    respjson.Field
		SignalProcessingLoss            respjson.Field
		SiteCode                        respjson.Field
		SolarExclAngle                  respjson.Field
		SpecAvgSpectraNums              respjson.Field
		SystemNoiseTemperature          respjson.Field
		TaskableRange                   respjson.Field
		TempMedFiltCodes                respjson.Field
		TestNumber                      respjson.Field
		TotRecNums                      respjson.Field
		TowerHeight                     respjson.Field
		TrackAngle                      respjson.Field
		TrackSnr                        respjson.Field
		TransmitGain                    respjson.Field
		TransmitHorizBeamWidth          respjson.Field
		TransmitLoss                    respjson.Field
		TransmitPower                   respjson.Field
		TransmitVertBeamWidth           respjson.Field
		TrueNorthCorrector              respjson.Field
		TrueTilt                        respjson.Field
		TwilightAngle                   respjson.Field
		VertBeamFlag                    respjson.Field
		VertGateSpacings                respjson.Field
		VertGateWidths                  respjson.Field
		VFov                            respjson.Field
		VResPixels                      respjson.Field
		WaveformBandwidths              respjson.Field
		WaveformLoopGains               respjson.Field
		WaveformMaxRanges               respjson.Field
		WaveformMinRanges               respjson.Field
		WaveformPulseWidths             respjson.Field
		Z1MaxRange                      respjson.Field
		Z1MinRange                      respjson.Field
		Z2MaxRange                      respjson.Field
		Z2MinRange                      respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorListResponseSensorcharacteristic) RawJSON() string { return r.JSON.raw }
func (r *SensorListResponseSensorcharacteristic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sensorlimits define 0 to many limits of a particular sensor in terms of
// observation coverage of on-orbit objects.
type SensorListResponseSensorlimitsCollection struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the target sensor object.
	IDSensor string `json:"idSensor"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensorLimits string `json:"idSensorLimits"`
	// Leftmost or minimum lower azimuth within this limit. Interpreted according to
	// site types as lower left azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower left azimuth limits would be equal. (degrees).
	LowerLeftAzimuthLimit float64 `json:"lowerLeftAzimuthLimit"`
	// Minimum or lower elevation within this limit. Interpreted according to site
	// types as minimum elevation angle, constant elevation or fan beam centerline.
	// (Degrees).
	LowerLeftElevationLimit float64 `json:"lowerLeftElevationLimit"`
	// Rightmost or maximum lower azimuth within this limit. Interpreted according to
	// site types as 2nd lower azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower right azimuth limits would be equal. (degrees).
	LowerRightAzimuthLimit float64 `json:"lowerRightAzimuthLimit"`
	// Minimum or lower right elevation within this limit. Interpreted according to
	// site types as minimum right elevation angle, constant elevation or fan beam
	// centerline. If the limit rectangle is parallel to the horizon, the left and
	// right lower elevation limits would be equal. (Degrees).
	LowerRightElevationLimit float64 `json:"lowerRightElevationLimit"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Leftmost or minimum upper azimuth within this sensor limit. Interpreted
	// according to site types as beginning upper azimuth limit, left-hand upper
	// boundary limit. If the limit rectangle is parallel to the horizon, the upper and
	// lower left azimuth limits would be equal. (in degrees).
	UpperLeftAzimuthLimit float64 `json:"upperLeftAzimuthLimit"`
	// Maximum or upper elevation within this limit. Interpreted according to site
	// types as maximum elevation angle, half the apex of conical observation pattern
	// or star. (Degrees).
	UpperLeftElevationLimit float64 `json:"upperLeftElevationLimit"`
	// Rightmost or maximum upper azimuth within this limit. Interpreted according to
	// site types as 2nd azimuth limit elevation angle of axis of conical observation
	// pattern. If the limit rectangle is parallel to the horizon, the upper and lower
	// right azimuth limits would be equal. (degrees).
	UpperRightAzimuthLimit float64 `json:"upperRightAzimuthLimit"`
	// Maximum or upper right elevation within this limit. Interpreted according to
	// site types as maximum rightmost elevation angle, half the apex of conical
	// observation pattern or star. If the limit rectangle is parallel to the horizon,
	// the left and right upper elevation limits would be equal. (Degrees).
	UpperRightElevationLimit float64 `json:"upperRightElevationLimit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		Source                   respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		IDSensor                 respjson.Field
		IDSensorLimits           respjson.Field
		LowerLeftAzimuthLimit    respjson.Field
		LowerLeftElevationLimit  respjson.Field
		LowerRightAzimuthLimit   respjson.Field
		LowerRightElevationLimit respjson.Field
		OrigNetwork              respjson.Field
		UpperLeftAzimuthLimit    respjson.Field
		UpperLeftElevationLimit  respjson.Field
		UpperRightAzimuthLimit   respjson.Field
		UpperRightElevationLimit respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorListResponseSensorlimitsCollection) RawJSON() string { return r.JSON.raw }
func (r *SensorListResponseSensorlimitsCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorListResponseSensorObservationType struct {
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The observation measurement type produced by a sensor.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		OrigNetwork respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorListResponseSensorObservationType) RawJSON() string { return r.JSON.raw }
func (r *SensorListResponseSensorObservationType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SensorStats contain statistics on sensors related to observation production such
// as last reported observation time.
type SensorListResponseSensorStat struct {
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
	// Unique ID of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Time of last reported observation in ISO 8601 UTC with microsecond precision.
	LastObTime time.Time `json:"lastObTime" format:"date-time"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSensor              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		LastObTime            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorListResponseSensorStat) RawJSON() string { return r.JSON.raw }
func (r *SensorListResponseSensorStat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorListResponseSensorType struct {
	// Unique identifier of the record, auto-generated by the system.
	ID int64 `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The specific sensor type and/or surveillance capability of this sensor.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		OrigNetwork respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorListResponseSensorType) RawJSON() string { return r.JSON.raw }
func (r *SensorListResponseSensorType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of observation data for electro-optical based sensor
// phenomenologies.
type SensorGetResponse struct {
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
	DataMode SensorGetResponseDataMode `json:"dataMode,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Optional flag indicating if the sensor is active.
	Active bool `json:"active"`
	// Optional US Air Force identifier for the sensor/ASR site, typically for air
	// surveillance radar (ASR) sensors.
	AfID string `json:"afId"`
	// The sensor type at the site. Optional field, intended primarily for ASRs.
	AsrType string `json:"asrType"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional dissemination control required for accessing data (e.g observations)
	// produced by this sensor. This is typically a proprietary data owner control for
	// commercial sensors.
	DataControl string `json:"dataControl"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity SensorGetResponseEntity `json:"entity"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity string `json:"idEntity"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensor string `json:"idSensor"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Collection of Sensorcharacteristics which define characteristics and
	// capabilities of a sensor.
	Sensorcharacteristics []SensorGetResponseSensorcharacteristic `json:"sensorcharacteristics"`
	// Sensorlimits define 0 to many limits of a particular sensor in terms of
	// observation coverage of on-orbit objects.
	SensorlimitsCollection []SensorGetResponseSensorlimitsCollection `json:"sensorlimitsCollection"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber          int64                                  `json:"sensorNumber"`
	SensorObservationType SensorGetResponseSensorObservationType `json:"sensorObservationType"`
	// Collection of SensorStats which contain statistics of a sensor.
	SensorStats []SensorGetResponseSensorStat `json:"sensorStats"`
	SensorType  SensorGetResponseSensorType   `json:"sensorType"`
	// Optional short name for the sensor.
	ShortName string `json:"shortName"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		SensorName             respjson.Field
		Source                 respjson.Field
		Active                 respjson.Field
		AfID                   respjson.Field
		AsrType                respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DataControl            respjson.Field
		Entity                 respjson.Field
		IDEntity               respjson.Field
		IDSensor               respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		Sensorcharacteristics  respjson.Field
		SensorlimitsCollection respjson.Field
		SensorNumber           respjson.Field
		SensorObservationType  respjson.Field
		SensorStats            respjson.Field
		SensorType             respjson.Field
		ShortName              respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponse) UnmarshalJSON(data []byte) error {
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
type SensorGetResponseDataMode string

const (
	SensorGetResponseDataModeReal      SensorGetResponseDataMode = "REAL"
	SensorGetResponseDataModeTest      SensorGetResponseDataMode = "TEST"
	SensorGetResponseDataModeSimulated SensorGetResponseDataMode = "SIMULATED"
	SensorGetResponseDataModeExercise  SensorGetResponseDataMode = "EXERCISE"
)

// An entity is a generic representation of any object within a space/SSA system
// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
// entity can have an operating unit, a location (if terrestrial), and statuses.
type SensorGetResponseEntity struct {
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
	// Unique entity name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
	// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
	//
	// Any of "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT",
	// "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL".
	Type string `json:"type,required"`
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
	Location shared.LocationFull `json:"location"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit SensorGetResponseEntityOnOrbit `json:"onOrbit"`
	// Model representation of a unit or organization which operates or controls a
	// space-related Entity such as an on-orbit payload, a sensor, etc. A contact may
	// belong to an organization.
	OperatingUnit SensorGetResponseEntityOperatingUnit `json:"operatingUnit"`
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
	OwnerType string `json:"ownerType"`
	// Read-only collection of RF bands utilized by this entity for communication
	// and/or operation.
	RfBands []SensorGetResponseEntityRfBand `json:"rfBands"`
	// Read-only collection of statuses which can be collected by multiple sources.
	StatusCollection []SensorGetResponseEntityStatusCollection `json:"statusCollection"`
	// Boolean indicating if this entity is taskable.
	Taskable bool `json:"taskable"`
	// Terrestrial identifier of this entity, if applicable.
	TerrestrialID string `json:"terrestrialId"`
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
		TerrestrialID         respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		URLs                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorGetResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model object representing on-orbit objects or satellites in the system.
type SensorGetResponseEntityOnOrbit struct {
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
	Antennas []SensorGetResponseEntityOnOrbitAntenna `json:"antennas"`
	// Read-only collection of batteries on this on-orbit object.
	Batteries []SensorGetResponseEntityOnOrbitBattery `json:"batteries"`
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
	OnorbitDetails []SensorGetResponseEntityOnOrbitOnorbitDetail `json:"onorbitDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only collection of solar arrays on this on-orbit object.
	SolarArrays []SensorGetResponseEntityOnOrbitSolarArray `json:"solarArrays"`
	// Read-only collection of thrusters (engines) on this on-orbit object.
	Thrusters []SensorGetResponseEntityOnOrbitThruster `json:"thrusters"`
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
func (r SensorGetResponseEntityOnOrbit) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityOnOrbit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorGetResponseEntityOnOrbitAntenna struct {
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
	Antenna shared.AntennaFull `json:"antenna"`
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
func (r SensorGetResponseEntityOnOrbitAntenna) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityOnOrbitAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorGetResponseEntityOnOrbitBattery struct {
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
	Battery shared.BatteryFull `json:"battery"`
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
func (r SensorGetResponseEntityOnOrbitBattery) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityOnOrbitBattery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains details of the OnOrbit object.
type SensorGetResponseEntityOnOrbitOnorbitDetail struct {
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
func (r SensorGetResponseEntityOnOrbitOnorbitDetail) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityOnOrbitOnorbitDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorGetResponseEntityOnOrbitSolarArray struct {
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
	SolarArray SensorGetResponseEntityOnOrbitSolarArraySolarArray `json:"solarArray"`
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
func (r SensorGetResponseEntityOnOrbitSolarArray) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityOnOrbitSolarArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of information on on-orbit/spacecraft solar arrays. A
// spacecraft may have multiple solar arrays and each solar array can have multiple
// 'details' records compiled by different sources.
type SensorGetResponseEntityOnOrbitSolarArraySolarArray struct {
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
	SolarArrayDetails []shared.SolarArrayDetailsFull `json:"solarArrayDetails"`
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
func (r SensorGetResponseEntityOnOrbitSolarArraySolarArray) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityOnOrbitSolarArraySolarArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorGetResponseEntityOnOrbitThruster struct {
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
	Engine shared.Engine `json:"engine"`
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
func (r SensorGetResponseEntityOnOrbitThruster) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityOnOrbitThruster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a unit or organization which operates or controls a
// space-related Entity such as an on-orbit payload, a sensor, etc. A contact may
// belong to an organization.
type SensorGetResponseEntityOperatingUnit struct {
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
	Location shared.LocationFull `json:"location"`
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
	OperatingUnitRemarks []SensorGetResponseEntityOperatingUnitOperatingUnitRemark `json:"operatingUnitRemarks"`
	// The Degree to which an operating unit is ready to perform the overall
	// operational mission(s) for which it was organized and equipped. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	OperStatus string `json:"operStatus"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	Organization shared.OrganizationFull `json:"organization"`
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
func (r SensorGetResponseEntityOperatingUnit) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityOperatingUnit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
type SensorGetResponseEntityOperatingUnitOperatingUnitRemark struct {
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
func (r SensorGetResponseEntityOperatingUnitOperatingUnitRemark) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityOperatingUnitOperatingUnitRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details on a particular Radio Frequency (RF) band, also known as a carrier,
// which may be in use by any type of Entity for communications or operations.
type SensorGetResponseEntityRfBand struct {
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
func (r SensorGetResponseEntityRfBand) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityRfBand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status for a particular Entity. An entity may have multiple status records
// collected by various sources.
type SensorGetResponseEntityStatusCollection struct {
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
	State               string                                                       `json:"state"`
	SubStatusCollection []SensorGetResponseEntityStatusCollectionSubStatusCollection `json:"subStatusCollection"`
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
func (r SensorGetResponseEntityStatusCollection) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseEntityStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional sub-system or capability status for the parent entity.
type SensorGetResponseEntityStatusCollectionSubStatusCollection struct {
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
func (r SensorGetResponseEntityStatusCollectionSubStatusCollection) RawJSON() string {
	return r.JSON.raw
}
func (r *SensorGetResponseEntityStatusCollectionSubStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of characteristics and capabilities of a sensor.
type SensorGetResponseSensorcharacteristic struct {
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
	// Unique identifier of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of measurement range(s) where radar samples must fall to be acceptable. If
	// this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	AcceptSampleRanges []float64 `json:"acceptSampleRanges"`
	// Number of bits used in the conversion from analog electrons in a pixel well to a
	// digital number. The digital number has a maximum value of 2^N, where N is the
	// number of bits.
	AnalogToDigitalBitSize int64 `json:"analogToDigitalBitSize"`
	// Optical sensor camera aperture.
	Aperture float64 `json:"aperture"`
	// For ASR (Air Surveillance Radar) sensors, the scan (360 deg sweep) rate of the
	// radar, in scans/minute.
	AsrScanRate float64 `json:"asrScanRate"`
	// One-way radar receiver loss factor due to atmospheric effects. This value will
	// often be the same as the corresponding transmission factor but may be different
	// for bistatic systems.
	AtmosReceiverLoss float64 `json:"atmosReceiverLoss"`
	// One-way radar transmission loss factor due to atmospheric effects.
	AtmosTransmissionLoss float64 `json:"atmosTransmissionLoss"`
	// Average atmospheric angular width with no distortion from turbulence at an
	// optical sensor site in arcseconds.
	AvgAtmosSeeingConditions float64 `json:"avgAtmosSeeingConditions"`
	// Array of azimuth angles of a radar beam, in degrees. If this field is populated,
	// the associated beam(s) must be provided in the beamOrder field.
	AzAngs []float64 `json:"azAngs"`
	// Azimuth rate acquisition limit (radians/minute).
	AzimuthRate float64 `json:"azimuthRate"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of watts per square meter per steradian
	// (W/(m^2 str)) consistent with sensor detection bands.
	BackgroundSkyRadiance float64 `json:"backgroundSkyRadiance"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of visual magnitude consistent with sensor
	// detection bands.
	BackgroundSkyVisMag float64 `json:"backgroundSkyVisMag"`
	// Sensor band.
	Band string `json:"band"`
	// The total bandwidth, in megahertz, about the radar center frequency.
	Bandwidth float64 `json:"bandwidth"`
	// Array designating the beam order of provided values (e.g. vb1 for vertical beam
	// 1, ob1 for oblique beam 1, etc.). Required if any of the following array fields
	// are populated: azAngs, elAngs, radarPulseWidths, pulseRepPeriods, delayGates,
	// rangeGates, rangeSpacings, vertGateSpacings, vertGateWidths, specAvgSpectraNums,
	// tempMedFiltCodes, runMeanCodes, totRecNums, reqRecords, acceptSampleRanges.
	BeamOrder []string `json:"beamOrder"`
	// Number of radar beams used by the sensor.
	BeamQty int64 `json:"beamQty"`
	// The angle of the center of a phased array sensor.
	Boresight float64 `json:"boresight"`
	// The number of degrees off of the boresight for the sensor.
	BoresightOffAngle float64 `json:"boresightOffAngle"`
	// Weighted center wavelength for an optical sensor bandpass in micrometers. It is
	// the center wavelength in a weighted integral sense, accounting for the
	// sensitivity vs. wavelength curve for the sensor focal plane array.
	CenterWavelength float64 `json:"centerWavelength"`
	// Collapsing loss in decibels. Collapsing losses occur when two or more sources of
	// noise are added to the target signal. Examples include receiver bandwidth
	// mismatch with filtering bandwidth and elevation or azimuth beam collapse onto
	// position/height indicator displays.
	CollapsingLoss float64 `json:"collapsingLoss"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Threshold shear value beyond which one of the radial velocity values will be
	// rejected, measured in units of inverse second.
	CritShear float64 `json:"critShear"`
	// Current flowing through optical sensor focal plane electronics with a closed
	// shutter in electrons per second.
	DarkCurrent float64 `json:"darkCurrent"`
	// Array of time delay(s) for pulses from a radar beam to get to the first range
	// gate, in nanoseconds. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	DelayGates []float64 `json:"delayGates"`
	// Description of the equipment and data source.
	Description string `json:"description"`
	// Detection signal-to-noise ratio (SNR) threshold in decibels. This value is
	// typically lower than trackSNR.
	DetectSnr float64 `json:"detectSNR"`
	// Sensor duty cycle as a fraction of 1. Duty cycle is the fraction of time a
	// sensor is actively transmitting.
	DutyCycle float64 `json:"dutyCycle"`
	// Sensor Earth limb exclusion height in kilometers and is generally only applied
	// to space-based sensors. Some models used an earth exclusion angle instead, but
	// this assumes the sensor is in a circular orbit with constant altitude relative
	// to the earth. The limb exclusion height can be used for space-based sensors in
	// any orbit (assuming it is constant with sensor altitude). The limb height is
	// defined to be 0 at the surface of the earth.
	EarthLimbExclHgt float64 `json:"earthLimbExclHgt"`
	// Array of elevation angles of a radar beam, in degrees. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	ElAngs []float64 `json:"elAngs"`
	// Elevation rate acquisition limit (radians/minute).
	ElevationRateGeolm float64 `json:"elevationRateGeolm"`
	// Type of equipment used to take measurements.
	EquipmentType string `json:"equipmentType"`
	// The beam width of a Sensor's Fan (range). The values for this range from (0.0 to
	// PI).
	FanBeamWidth float64 `json:"fanBeamWidth"`
	// Number of Fast Fourier Transform (FFT) points used to convert time varying
	// signals into the frequency domain.
	Fft int64 `json:"fft"`
	// Maximum number of times the first guess was propagated in each gate before
	// failing the first guess check.
	FgpCrit int64 `json:"fgpCrit"`
	// Noise term, in decibels, that arises when a radar receiver filter has a
	// non-optimal bandwidth for an incoming signal (i.e., when it does not match the
	// pulse width).
	FilterMismatchFactor float64 `json:"filterMismatchFactor"`
	// F-number for an optical telescope. It is dimensionless and is defined as the
	// ratio of the focal length to the aperture width.
	FNum float64 `json:"fNum"`
	// For radar based sensors, the focal point elevation of the radar at the site, in
	// meters.
	FocalPoint float64 `json:"focalPoint"`
	// Horizontal field of view, in degrees.
	HFov float64 `json:"hFOV"`
	// Horizontal pixel resolution.
	HResPixels int64 `json:"hResPixels"`
	// For radar based sensors, K-factor is a relative indicator of refractivity that
	// infers the amount of radar beam bending due to atmosphere. (1<K<2).
	K float64 `json:"k"`
	// For Orbiting Sensors, First Card Azimuth limit #1 (left, degrees).
	LeftClockAngle float64 `json:"leftClockAngle"`
	// Leftmost GEO belt longitude limit for this sensor (if applicable).
	LeftGeoBeltLimit float64 `json:"leftGeoBeltLimit"`
	// Site where measurement is taken.
	Location string `json:"location"`
	// Aggregated radar range equation gain in decibels for maximum sensitivity. It is
	// a roll-up value for low fidelity modeling and is often the only sensitivity
	// value available for a radar system without access to detailed performance
	// parameters.
	LoopGain float64 `json:"loopGain"`
	// Lowest aspect angle of the full moon in degrees at which the sensor can achieve
	// full performance.
	LunarExclAngle float64 `json:"lunarExclAngle"`
	// Angle between magnetic north and true north at the sensor site, in degrees.
	MagDec float64 `json:"magDec"`
	// Absolute magnitude acquisition limit for optical sensors.
	MagnitudeLimit float64 `json:"magnitudeLimit"`
	// Max deviation angle of the sensor in degrees.
	MaxDeviationAngle float64 `json:"maxDeviationAngle"`
	// Maximum integration time per image frame in seconds for an optical sensor.
	MaxIntegrationTime float64 `json:"maxIntegrationTime"`
	// Maximum observable sensor range, in kilometers.
	MaxObservableRange float64 `json:"maxObservableRange"`
	// Maximum observable range limit in kilometers -- sensor cannot acquire beyond
	// this range.
	MaxRangeLimit float64 `json:"maxRangeLimit"`
	// Maximum wavelength detectable by an optical sensor in micrometers.
	MaxWavelength float64 `json:"maxWavelength"`
	// Minimum integration time per image frame in seconds for an optical sensor.
	MinIntegrationTime float64 `json:"minIntegrationTime"`
	// Minimum range measurement capability of the sensor, in kilometers.
	MinRangeLimit float64 `json:"minRangeLimit"`
	// Signal to Noise Ratio, in decibels. The values for this range from 0.0 - + 99.99
	// dB.
	MinSignalNoiseRatio float64 `json:"minSignalNoiseRatio"`
	// Minimum wavelength detectable by an optical sensor in micrometers.
	MinWavelength float64 `json:"minWavelength"`
	// Negative Range-rate/relative velocity limit (kilometers/second).
	NegativeRangeRateLimit float64 `json:"negativeRangeRateLimit"`
	// Noise figure for a radar system in decibels. This value may be used to compute
	// system noise when the system temperature is unavailable.
	NoiseFigure float64 `json:"noiseFigure"`
	// Number of pulses that are non-coherently integrated during detection processing.
	NonCoherentIntegratedPulses int64 `json:"nonCoherentIntegratedPulses"`
	// For radar based sensors, number of integrated pulses in a transmit cycle.
	NumIntegratedPulses int64 `json:"numIntegratedPulses"`
	// Number of integration frames for an optical sensor.
	NumIntegrationFrames int64 `json:"numIntegrationFrames"`
	// The number of optical integration mode characteristics provided in this record.
	// If provided, the numOpticalIntegrationModes value indicates the number of
	// elements in each of the opticalIntegrationTimes, opticalIntegrationAngularRates,
	// opticalIntegrationFrames, opticalIntegrationPixelBinnings, and
	// opticalIntegrationSNRs arrays.
	NumOpticalIntegrationModes int64 `json:"numOpticalIntegrationModes"`
	// The number of waveforms characteristics provided in this record. If provided,
	// the numWaveforms value indicates the number of elements in each of the
	// waveformPulseWidths, waveformBandWidths, waveformMinRanges, waveformMaxRanges,
	// and waveformLoopGains arrays.
	NumWaveforms int64 `json:"numWaveforms"`
	// Array containing the angular rate, in arcsec/sec, for each provided optical
	// integration mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationAngularRates []float64 `json:"opticalIntegrationAngularRates"`
	// Array containing the number of frames, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationFrames []float64 `json:"opticalIntegrationFrames"`
	// Array containing the pixel binning, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationPixelBinnings []float64 `json:"opticalIntegrationPixelBinnings"`
	// Array of optical integration modes signal to noise ratios. The number of
	// elements must be equal to the value indicated in numOpticalIntegrationModes.
	OpticalIntegrationSnRs []float64 `json:"opticalIntegrationSNRs"`
	// Array containing the time, in seconds, for each provided optical integration
	// mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationTimes []float64 `json:"opticalIntegrationTimes"`
	// Fraction of incident light transmitted to an optical sensor focal plane array.
	// The value is given as a fraction of 1, not as a percent.
	OpticalTransmission float64 `json:"opticalTransmission"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Two-way pattern absorption/propagation loss due to medium in decibels.
	PatternAbsorptionLoss float64 `json:"patternAbsorptionLoss"`
	// Losses from the beam shape, scanning, and pattern factor in decibels. These
	// losses occur when targets are not directly in line with a beam center. For space
	// surveillance, this will occur most often during sensor scanning.
	PatternScanLoss float64 `json:"patternScanLoss"`
	// Maximum instantaneous radar transmit power in watts for use in the radar range
	// equation.
	PeakPower float64 `json:"peakPower"`
	// Angular field-of-view covered by one pixel in a focal plane array in
	// microradians. The pixel is assumed to be a perfect square so that only a single
	// value is required.
	PixelInstantaneousFov float64 `json:"pixelInstantaneousFOV"`
	// Maximum number of electrons that can be collected in a single pixel on an
	// optical sensor focal plane array.
	PixelWellDepth int64 `json:"pixelWellDepth"`
	// Positive Range-rate/relative velocity limit (kilometers/second).
	PositiveRangeRateLimit float64 `json:"positiveRangeRateLimit"`
	// For radar based sensors, pulse repetition frequency (PRF), in hertz. Number of
	// new pulses transmitted per second.
	Prf float64 `json:"prf"`
	// Designated probability of detection at the signal-to-noise detection threshold.
	ProbDetectSnr float64 `json:"probDetectSNR"`
	// For radar based sensors, probability of the indication of the presence of a
	// radar target due to noise or interference.
	ProbFalseAlarm float64 `json:"probFalseAlarm"`
	// Array of interval(s) between the start of one radar pulse and the start of
	// another for a radar beam, in microseconds. If this field is populated, the
	// associated beam(s) must be provided in the beamOrder field.
	PulseRepPeriods []float64 `json:"pulseRepPeriods"`
	// Fraction of incident photons converted to electrons at the focal plane array.
	// This value is a decimal number between 0 and 1, inclusive.
	QuantumEff float64 `json:"quantumEff"`
	// Radar frequency in hertz, of the sensor (if a radar sensor).
	RadarFrequency float64 `json:"radarFrequency"`
	// Message data format transmitted by the sensor digitizer.
	RadarMessageFormat string `json:"radarMessageFormat"`
	// For radar based sensors, radar maximum unambiguous range, in kilometers.
	RadarMur float64 `json:"radarMUR"`
	// Array of transmit time(s) for a radar beam pulse, in microseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RadarPulseWidths []float64 `json:"radarPulseWidths"`
	// Radio frequency (if sensor is RF).
	RadioFrequency float64 `json:"radioFrequency"`
	// Losses due to the presence of a protective radome surrounding a radar sensor, in
	// decibels.
	RadomeLoss float64 `json:"radomeLoss"`
	// Array of the number(s) of discrete altitudes where return signals are sampled by
	// a radar beam. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	RangeGates []int64 `json:"rangeGates"`
	// Array of range gate spacing(s) for a radar beam, in nanoseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RangeSpacings []float64 `json:"rangeSpacings"`
	// Number of false-signal electrons generated by optical sensor focal plane
	// read-out electronics from photon-to-electron conversion during frame
	// integration. The units are in electrons RMS.
	ReadNoise int64 `json:"readNoise"`
	// Radar receive gain in decibels.
	ReceiveGain float64 `json:"receiveGain"`
	// Horizontal/azimuthal receive beamwidth for a radar in degrees.
	ReceiveHorizBeamWidth float64 `json:"receiveHorizBeamWidth"`
	// Aggregate radar receive loss, in decibels.
	ReceiveLoss float64 `json:"receiveLoss"`
	// Vertical/elevation receive beamwidth for a radar in degrees.
	ReceiveVertBeamWidth float64 `json:"receiveVertBeamWidth"`
	// Reference temperature for radar noise in Kelvin. A reference temperature is used
	// when the radar system temperature is unknown and is combined with the system
	// noise figure to estimate signal loss.
	RefTemp float64 `json:"refTemp"`
	// Array of the total number(s) of records required to meet consensus for a radar
	// beam. If this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	ReqRecords []int64 `json:"reqRecords"`
	// For Orbiting Sensors, First Card Azimuth limit #3 (right, degrees).
	RightClockAngle float64 `json:"rightClockAngle"`
	// Rightmost GEO belt longitude limit for this sensor (if applicable).
	RightGeoBeltLimit float64 `json:"rightGeoBeltLimit"`
	// Array of running mean code(s) used by radar data processing. The running mean
	// method involves taking a series of averages of different selections of the full
	// data set to smooth out short-term fluctuations in the data. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	RunMeanCodes []int64 `json:"runMeanCodes"`
	// Radar signal processing losses, in decibels.
	SignalProcessingLoss float64 `json:"signalProcessingLoss"`
	// Site code of the sensor.
	SiteCode string `json:"siteCode"`
	// Sensor and target position vector origins are at the center of the earth. The
	// sun vector origin is at the target position and points toward the sun. Any value
	// between 0 and 180 degrees is acceptable and is assumed to apply in both
	// directions (i.e., a solar exclusion angle of 30 degrees is understood to mean no
	// viewing for any angle between -30 deg and +30 deg).
	SolarExclAngle float64 `json:"solarExclAngle"`
	// Array of the number(s) of Doppler spectra used to process measurements from
	// radar. Spectral averaging involves combining multiple Doppler spectra acquired
	// to obtain a more accurate and representative spectrum. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	SpecAvgSpectraNums []int64 `json:"specAvgSpectraNums"`
	// For radar based sensors, expression of the radar system noise, aggregated as an
	// equivalent thermal noise value, in degrees Kelvin.
	SystemNoiseTemperature float64 `json:"systemNoiseTemperature"`
	// Maximum taskable range of the sensor, in kilometers.
	TaskableRange float64 `json:"taskableRange"`
	// Array of temporal median filter code(s) of a radar beam. Temporal median
	// filtering is a noise-reducing algorithm which involves replacing each data point
	// with the median value of a window of neighboring points over time. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	TempMedFiltCodes []int64 `json:"tempMedFiltCodes"`
	// Test number for the observed measurement.
	TestNumber string `json:"testNumber"`
	// Array of the total number(s) of records for a radar beam. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	TotRecNums []int64 `json:"totRecNums"`
	// For tower sensors, the physical height of the sensor tower, in meters.
	TowerHeight float64 `json:"towerHeight"`
	// Beginning track angle limit, in radians. Track angle is the angle between the
	// camera axis and the gimbal plane. Values range from 0 - PI/2.
	TrackAngle float64 `json:"trackAngle"`
	// Track signal-to-noise ratio (SNR) threshold in decibels. This value is typically
	// higher than detectSNR.
	TrackSnr float64 `json:"trackSNR"`
	// Radar transmit gain in decibels.
	TransmitGain float64 `json:"transmitGain"`
	// Horizontal/azimuthal transmit beamwidth for a radar in degrees.
	TransmitHorizBeamWidth float64 `json:"transmitHorizBeamWidth"`
	// Aggregate radar transmit loss, in decibels.
	TransmitLoss float64 `json:"transmitLoss"`
	// Radar transmit power in Watts.
	TransmitPower float64 `json:"transmitPower"`
	// Vertical/elevation transmit beamwidth for a radar in degrees.
	TransmitVertBeamWidth float64 `json:"transmitVertBeamWidth"`
	// True North correction for the sensor, in ACP (Azimunth Change Pulse) count.
	TrueNorthCorrector int64 `json:"trueNorthCorrector"`
	// Antenna true tilt, in degrees.
	TrueTilt float64 `json:"trueTilt"`
	// Twilight angle for ground-based optical sensors in degrees. A sensor cannot view
	// targets until the sun is below the twilight angle relative to the local horizon.
	// The sign of the angle is positive despite the sun elevation being negative after
	// local sunset. Typical values for the twilight angle are civil twilight (6
	// degrees), nautical twilight (12 degrees), and astronomical twilight (18
	// degrees).
	TwilightAngle float64 `json:"twilightAngle"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Flag indicating if a vertical radar beam was used in the wind calculation.
	VertBeamFlag bool `json:"vertBeamFlag"`
	// Array of vertical distance(s) between points where radar measurements are taken,
	// in meters. If this field is populated, the associated beam(s) must be provided
	// in the beamOrder field.
	VertGateSpacings []float64 `json:"vertGateSpacings"`
	// Array of width(s) of each location where radar measurements are taken, in
	// meters. If this field is populated, the associated beam(s) must be provided in
	// the beamOrder field.
	VertGateWidths []float64 `json:"vertGateWidths"`
	// Vertical field of view, in degrees.
	VFov float64 `json:"vFOV"`
	// Vertical pixel resolution.
	VResPixels int64 `json:"vResPixels"`
	// Array containing the bandwidth, in megahertz, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformBandwidths []float64 `json:"waveformBandwidths"`
	// Array containing the loop gain, in decibels, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field (10 SNR vs. 1 dBsm at 1000 km).
	WaveformLoopGains []float64 `json:"waveformLoopGains"`
	// Array containing the maximum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMaxRanges []float64 `json:"waveformMaxRanges"`
	// Array containing the minimum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMinRanges []float64 `json:"waveformMinRanges"`
	// Array containing the pulse width, in microseconds, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformPulseWidths []float64 `json:"waveformPulseWidths"`
	// Peformance zone-1 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MaxRange float64 `json:"z1MaxRange"`
	// Peformance zone-1 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MinRange float64 `json:"z1MinRange"`
	// Peformance zone-2 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MaxRange float64 `json:"z2MaxRange"`
	// Peformance zone-2 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MinRange float64 `json:"z2MinRange"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking           respjson.Field
		DataMode                        respjson.Field
		IDSensor                        respjson.Field
		Source                          respjson.Field
		ID                              respjson.Field
		AcceptSampleRanges              respjson.Field
		AnalogToDigitalBitSize          respjson.Field
		Aperture                        respjson.Field
		AsrScanRate                     respjson.Field
		AtmosReceiverLoss               respjson.Field
		AtmosTransmissionLoss           respjson.Field
		AvgAtmosSeeingConditions        respjson.Field
		AzAngs                          respjson.Field
		AzimuthRate                     respjson.Field
		BackgroundSkyRadiance           respjson.Field
		BackgroundSkyVisMag             respjson.Field
		Band                            respjson.Field
		Bandwidth                       respjson.Field
		BeamOrder                       respjson.Field
		BeamQty                         respjson.Field
		Boresight                       respjson.Field
		BoresightOffAngle               respjson.Field
		CenterWavelength                respjson.Field
		CollapsingLoss                  respjson.Field
		CreatedAt                       respjson.Field
		CreatedBy                       respjson.Field
		CritShear                       respjson.Field
		DarkCurrent                     respjson.Field
		DelayGates                      respjson.Field
		Description                     respjson.Field
		DetectSnr                       respjson.Field
		DutyCycle                       respjson.Field
		EarthLimbExclHgt                respjson.Field
		ElAngs                          respjson.Field
		ElevationRateGeolm              respjson.Field
		EquipmentType                   respjson.Field
		FanBeamWidth                    respjson.Field
		Fft                             respjson.Field
		FgpCrit                         respjson.Field
		FilterMismatchFactor            respjson.Field
		FNum                            respjson.Field
		FocalPoint                      respjson.Field
		HFov                            respjson.Field
		HResPixels                      respjson.Field
		K                               respjson.Field
		LeftClockAngle                  respjson.Field
		LeftGeoBeltLimit                respjson.Field
		Location                        respjson.Field
		LoopGain                        respjson.Field
		LunarExclAngle                  respjson.Field
		MagDec                          respjson.Field
		MagnitudeLimit                  respjson.Field
		MaxDeviationAngle               respjson.Field
		MaxIntegrationTime              respjson.Field
		MaxObservableRange              respjson.Field
		MaxRangeLimit                   respjson.Field
		MaxWavelength                   respjson.Field
		MinIntegrationTime              respjson.Field
		MinRangeLimit                   respjson.Field
		MinSignalNoiseRatio             respjson.Field
		MinWavelength                   respjson.Field
		NegativeRangeRateLimit          respjson.Field
		NoiseFigure                     respjson.Field
		NonCoherentIntegratedPulses     respjson.Field
		NumIntegratedPulses             respjson.Field
		NumIntegrationFrames            respjson.Field
		NumOpticalIntegrationModes      respjson.Field
		NumWaveforms                    respjson.Field
		OpticalIntegrationAngularRates  respjson.Field
		OpticalIntegrationFrames        respjson.Field
		OpticalIntegrationPixelBinnings respjson.Field
		OpticalIntegrationSnRs          respjson.Field
		OpticalIntegrationTimes         respjson.Field
		OpticalTransmission             respjson.Field
		OrigNetwork                     respjson.Field
		PatternAbsorptionLoss           respjson.Field
		PatternScanLoss                 respjson.Field
		PeakPower                       respjson.Field
		PixelInstantaneousFov           respjson.Field
		PixelWellDepth                  respjson.Field
		PositiveRangeRateLimit          respjson.Field
		Prf                             respjson.Field
		ProbDetectSnr                   respjson.Field
		ProbFalseAlarm                  respjson.Field
		PulseRepPeriods                 respjson.Field
		QuantumEff                      respjson.Field
		RadarFrequency                  respjson.Field
		RadarMessageFormat              respjson.Field
		RadarMur                        respjson.Field
		RadarPulseWidths                respjson.Field
		RadioFrequency                  respjson.Field
		RadomeLoss                      respjson.Field
		RangeGates                      respjson.Field
		RangeSpacings                   respjson.Field
		ReadNoise                       respjson.Field
		ReceiveGain                     respjson.Field
		ReceiveHorizBeamWidth           respjson.Field
		ReceiveLoss                     respjson.Field
		ReceiveVertBeamWidth            respjson.Field
		RefTemp                         respjson.Field
		ReqRecords                      respjson.Field
		RightClockAngle                 respjson.Field
		RightGeoBeltLimit               respjson.Field
		RunMeanCodes                    respjson.Field
		SignalProcessingLoss            respjson.Field
		SiteCode                        respjson.Field
		SolarExclAngle                  respjson.Field
		SpecAvgSpectraNums              respjson.Field
		SystemNoiseTemperature          respjson.Field
		TaskableRange                   respjson.Field
		TempMedFiltCodes                respjson.Field
		TestNumber                      respjson.Field
		TotRecNums                      respjson.Field
		TowerHeight                     respjson.Field
		TrackAngle                      respjson.Field
		TrackSnr                        respjson.Field
		TransmitGain                    respjson.Field
		TransmitHorizBeamWidth          respjson.Field
		TransmitLoss                    respjson.Field
		TransmitPower                   respjson.Field
		TransmitVertBeamWidth           respjson.Field
		TrueNorthCorrector              respjson.Field
		TrueTilt                        respjson.Field
		TwilightAngle                   respjson.Field
		UpdatedAt                       respjson.Field
		UpdatedBy                       respjson.Field
		VertBeamFlag                    respjson.Field
		VertGateSpacings                respjson.Field
		VertGateWidths                  respjson.Field
		VFov                            respjson.Field
		VResPixels                      respjson.Field
		WaveformBandwidths              respjson.Field
		WaveformLoopGains               respjson.Field
		WaveformMaxRanges               respjson.Field
		WaveformMinRanges               respjson.Field
		WaveformPulseWidths             respjson.Field
		Z1MaxRange                      respjson.Field
		Z1MinRange                      respjson.Field
		Z2MaxRange                      respjson.Field
		Z2MinRange                      respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorGetResponseSensorcharacteristic) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseSensorcharacteristic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sensorlimits define 0 to many limits of a particular sensor in terms of
// observation coverage of on-orbit objects.
type SensorGetResponseSensorlimitsCollection struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the target sensor object.
	IDSensor string `json:"idSensor"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensorLimits string `json:"idSensorLimits"`
	// Leftmost or minimum lower azimuth within this limit. Interpreted according to
	// site types as lower left azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower left azimuth limits would be equal. (degrees).
	LowerLeftAzimuthLimit float64 `json:"lowerLeftAzimuthLimit"`
	// Minimum or lower elevation within this limit. Interpreted according to site
	// types as minimum elevation angle, constant elevation or fan beam centerline.
	// (Degrees).
	LowerLeftElevationLimit float64 `json:"lowerLeftElevationLimit"`
	// Rightmost or maximum lower azimuth within this limit. Interpreted according to
	// site types as 2nd lower azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower right azimuth limits would be equal. (degrees).
	LowerRightAzimuthLimit float64 `json:"lowerRightAzimuthLimit"`
	// Minimum or lower right elevation within this limit. Interpreted according to
	// site types as minimum right elevation angle, constant elevation or fan beam
	// centerline. If the limit rectangle is parallel to the horizon, the left and
	// right lower elevation limits would be equal. (Degrees).
	LowerRightElevationLimit float64 `json:"lowerRightElevationLimit"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Leftmost or minimum upper azimuth within this sensor limit. Interpreted
	// according to site types as beginning upper azimuth limit, left-hand upper
	// boundary limit. If the limit rectangle is parallel to the horizon, the upper and
	// lower left azimuth limits would be equal. (in degrees).
	UpperLeftAzimuthLimit float64 `json:"upperLeftAzimuthLimit"`
	// Maximum or upper elevation within this limit. Interpreted according to site
	// types as maximum elevation angle, half the apex of conical observation pattern
	// or star. (Degrees).
	UpperLeftElevationLimit float64 `json:"upperLeftElevationLimit"`
	// Rightmost or maximum upper azimuth within this limit. Interpreted according to
	// site types as 2nd azimuth limit elevation angle of axis of conical observation
	// pattern. If the limit rectangle is parallel to the horizon, the upper and lower
	// right azimuth limits would be equal. (degrees).
	UpperRightAzimuthLimit float64 `json:"upperRightAzimuthLimit"`
	// Maximum or upper right elevation within this limit. Interpreted according to
	// site types as maximum rightmost elevation angle, half the apex of conical
	// observation pattern or star. If the limit rectangle is parallel to the horizon,
	// the left and right upper elevation limits would be equal. (Degrees).
	UpperRightElevationLimit float64 `json:"upperRightElevationLimit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		Source                   respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		IDSensor                 respjson.Field
		IDSensorLimits           respjson.Field
		LowerLeftAzimuthLimit    respjson.Field
		LowerLeftElevationLimit  respjson.Field
		LowerRightAzimuthLimit   respjson.Field
		LowerRightElevationLimit respjson.Field
		OrigNetwork              respjson.Field
		UpdatedAt                respjson.Field
		UpdatedBy                respjson.Field
		UpperLeftAzimuthLimit    respjson.Field
		UpperLeftElevationLimit  respjson.Field
		UpperRightAzimuthLimit   respjson.Field
		UpperRightElevationLimit respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorGetResponseSensorlimitsCollection) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseSensorlimitsCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorGetResponseSensorObservationType struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The observation measurement type produced by a sensor.
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataMode    respjson.Field
		Source      respjson.Field
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		OrigNetwork respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorGetResponseSensorObservationType) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseSensorObservationType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SensorStats contain statistics on sensors related to observation production such
// as last reported observation time.
type SensorGetResponseSensorStat struct {
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
	// Unique ID of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Time of last reported observation in ISO 8601 UTC with microsecond precision.
	LastObTime time.Time `json:"lastObTime" format:"date-time"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSensor              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		LastObTime            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorGetResponseSensorStat) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseSensorStat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorGetResponseSensorType struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID int64 `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The specific sensor type and/or surveillance capability of this sensor.
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataMode    respjson.Field
		Source      respjson.Field
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		OrigNetwork respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorGetResponseSensorType) RawJSON() string { return r.JSON.raw }
func (r *SensorGetResponseSensorType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorQueryhelpResponse struct {
	AodrSupported         bool                               `json:"aodrSupported"`
	ClassificationMarking string                             `json:"classificationMarking"`
	Description           string                             `json:"description"`
	HistorySupported      bool                               `json:"historySupported"`
	Name                  string                             `json:"name"`
	Parameters            []SensorQueryhelpResponseParameter `json:"parameters"`
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
func (r SensorQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorQueryhelpResponseParameter struct {
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
func (r SensorQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *SensorQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of observation data for electro-optical based sensor
// phenomenologies.
type SensorTupleResponse struct {
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
	DataMode SensorTupleResponseDataMode `json:"dataMode,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Optional flag indicating if the sensor is active.
	Active bool `json:"active"`
	// Optional US Air Force identifier for the sensor/ASR site, typically for air
	// surveillance radar (ASR) sensors.
	AfID string `json:"afId"`
	// The sensor type at the site. Optional field, intended primarily for ASRs.
	AsrType string `json:"asrType"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional dissemination control required for accessing data (e.g observations)
	// produced by this sensor. This is typically a proprietary data owner control for
	// commercial sensors.
	DataControl string `json:"dataControl"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity SensorTupleResponseEntity `json:"entity"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity string `json:"idEntity"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensor string `json:"idSensor"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Collection of Sensorcharacteristics which define characteristics and
	// capabilities of a sensor.
	Sensorcharacteristics []SensorTupleResponseSensorcharacteristic `json:"sensorcharacteristics"`
	// Sensorlimits define 0 to many limits of a particular sensor in terms of
	// observation coverage of on-orbit objects.
	SensorlimitsCollection []SensorTupleResponseSensorlimitsCollection `json:"sensorlimitsCollection"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber          int64                                    `json:"sensorNumber"`
	SensorObservationType SensorTupleResponseSensorObservationType `json:"sensorObservationType"`
	// Collection of SensorStats which contain statistics of a sensor.
	SensorStats []SensorTupleResponseSensorStat `json:"sensorStats"`
	SensorType  SensorTupleResponseSensorType   `json:"sensorType"`
	// Optional short name for the sensor.
	ShortName string `json:"shortName"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		SensorName             respjson.Field
		Source                 respjson.Field
		Active                 respjson.Field
		AfID                   respjson.Field
		AsrType                respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DataControl            respjson.Field
		Entity                 respjson.Field
		IDEntity               respjson.Field
		IDSensor               respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		Sensorcharacteristics  respjson.Field
		SensorlimitsCollection respjson.Field
		SensorNumber           respjson.Field
		SensorObservationType  respjson.Field
		SensorStats            respjson.Field
		SensorType             respjson.Field
		ShortName              respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponse) UnmarshalJSON(data []byte) error {
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
type SensorTupleResponseDataMode string

const (
	SensorTupleResponseDataModeReal      SensorTupleResponseDataMode = "REAL"
	SensorTupleResponseDataModeTest      SensorTupleResponseDataMode = "TEST"
	SensorTupleResponseDataModeSimulated SensorTupleResponseDataMode = "SIMULATED"
	SensorTupleResponseDataModeExercise  SensorTupleResponseDataMode = "EXERCISE"
)

// An entity is a generic representation of any object within a space/SSA system
// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
// entity can have an operating unit, a location (if terrestrial), and statuses.
type SensorTupleResponseEntity struct {
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
	// Unique entity name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
	// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
	//
	// Any of "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT",
	// "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL".
	Type string `json:"type,required"`
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
	Location shared.LocationFull `json:"location"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit SensorTupleResponseEntityOnOrbit `json:"onOrbit"`
	// Model representation of a unit or organization which operates or controls a
	// space-related Entity such as an on-orbit payload, a sensor, etc. A contact may
	// belong to an organization.
	OperatingUnit SensorTupleResponseEntityOperatingUnit `json:"operatingUnit"`
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
	OwnerType string `json:"ownerType"`
	// Read-only collection of RF bands utilized by this entity for communication
	// and/or operation.
	RfBands []SensorTupleResponseEntityRfBand `json:"rfBands"`
	// Read-only collection of statuses which can be collected by multiple sources.
	StatusCollection []SensorTupleResponseEntityStatusCollection `json:"statusCollection"`
	// Boolean indicating if this entity is taskable.
	Taskable bool `json:"taskable"`
	// Terrestrial identifier of this entity, if applicable.
	TerrestrialID string `json:"terrestrialId"`
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
		TerrestrialID         respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		URLs                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorTupleResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model object representing on-orbit objects or satellites in the system.
type SensorTupleResponseEntityOnOrbit struct {
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
	Antennas []SensorTupleResponseEntityOnOrbitAntenna `json:"antennas"`
	// Read-only collection of batteries on this on-orbit object.
	Batteries []SensorTupleResponseEntityOnOrbitBattery `json:"batteries"`
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
	OnorbitDetails []SensorTupleResponseEntityOnOrbitOnorbitDetail `json:"onorbitDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only collection of solar arrays on this on-orbit object.
	SolarArrays []SensorTupleResponseEntityOnOrbitSolarArray `json:"solarArrays"`
	// Read-only collection of thrusters (engines) on this on-orbit object.
	Thrusters []SensorTupleResponseEntityOnOrbitThruster `json:"thrusters"`
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
func (r SensorTupleResponseEntityOnOrbit) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityOnOrbit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorTupleResponseEntityOnOrbitAntenna struct {
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
	Antenna shared.AntennaFull `json:"antenna"`
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
func (r SensorTupleResponseEntityOnOrbitAntenna) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityOnOrbitAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorTupleResponseEntityOnOrbitBattery struct {
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
	Battery shared.BatteryFull `json:"battery"`
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
func (r SensorTupleResponseEntityOnOrbitBattery) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityOnOrbitBattery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains details of the OnOrbit object.
type SensorTupleResponseEntityOnOrbitOnorbitDetail struct {
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
func (r SensorTupleResponseEntityOnOrbitOnorbitDetail) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityOnOrbitOnorbitDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorTupleResponseEntityOnOrbitSolarArray struct {
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
	SolarArray SensorTupleResponseEntityOnOrbitSolarArraySolarArray `json:"solarArray"`
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
func (r SensorTupleResponseEntityOnOrbitSolarArray) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityOnOrbitSolarArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of information on on-orbit/spacecraft solar arrays. A
// spacecraft may have multiple solar arrays and each solar array can have multiple
// 'details' records compiled by different sources.
type SensorTupleResponseEntityOnOrbitSolarArraySolarArray struct {
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
	SolarArrayDetails []shared.SolarArrayDetailsFull `json:"solarArrayDetails"`
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
func (r SensorTupleResponseEntityOnOrbitSolarArraySolarArray) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityOnOrbitSolarArraySolarArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorTupleResponseEntityOnOrbitThruster struct {
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
	Engine shared.Engine `json:"engine"`
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
func (r SensorTupleResponseEntityOnOrbitThruster) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityOnOrbitThruster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a unit or organization which operates or controls a
// space-related Entity such as an on-orbit payload, a sensor, etc. A contact may
// belong to an organization.
type SensorTupleResponseEntityOperatingUnit struct {
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
	Location shared.LocationFull `json:"location"`
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
	OperatingUnitRemarks []SensorTupleResponseEntityOperatingUnitOperatingUnitRemark `json:"operatingUnitRemarks"`
	// The Degree to which an operating unit is ready to perform the overall
	// operational mission(s) for which it was organized and equipped. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	OperStatus string `json:"operStatus"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	Organization shared.OrganizationFull `json:"organization"`
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
func (r SensorTupleResponseEntityOperatingUnit) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityOperatingUnit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
type SensorTupleResponseEntityOperatingUnitOperatingUnitRemark struct {
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
func (r SensorTupleResponseEntityOperatingUnitOperatingUnitRemark) RawJSON() string {
	return r.JSON.raw
}
func (r *SensorTupleResponseEntityOperatingUnitOperatingUnitRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details on a particular Radio Frequency (RF) band, also known as a carrier,
// which may be in use by any type of Entity for communications or operations.
type SensorTupleResponseEntityRfBand struct {
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
func (r SensorTupleResponseEntityRfBand) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityRfBand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status for a particular Entity. An entity may have multiple status records
// collected by various sources.
type SensorTupleResponseEntityStatusCollection struct {
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
	State               string                                                         `json:"state"`
	SubStatusCollection []SensorTupleResponseEntityStatusCollectionSubStatusCollection `json:"subStatusCollection"`
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
func (r SensorTupleResponseEntityStatusCollection) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseEntityStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional sub-system or capability status for the parent entity.
type SensorTupleResponseEntityStatusCollectionSubStatusCollection struct {
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
func (r SensorTupleResponseEntityStatusCollectionSubStatusCollection) RawJSON() string {
	return r.JSON.raw
}
func (r *SensorTupleResponseEntityStatusCollectionSubStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of characteristics and capabilities of a sensor.
type SensorTupleResponseSensorcharacteristic struct {
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
	// Unique identifier of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of measurement range(s) where radar samples must fall to be acceptable. If
	// this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	AcceptSampleRanges []float64 `json:"acceptSampleRanges"`
	// Number of bits used in the conversion from analog electrons in a pixel well to a
	// digital number. The digital number has a maximum value of 2^N, where N is the
	// number of bits.
	AnalogToDigitalBitSize int64 `json:"analogToDigitalBitSize"`
	// Optical sensor camera aperture.
	Aperture float64 `json:"aperture"`
	// For ASR (Air Surveillance Radar) sensors, the scan (360 deg sweep) rate of the
	// radar, in scans/minute.
	AsrScanRate float64 `json:"asrScanRate"`
	// One-way radar receiver loss factor due to atmospheric effects. This value will
	// often be the same as the corresponding transmission factor but may be different
	// for bistatic systems.
	AtmosReceiverLoss float64 `json:"atmosReceiverLoss"`
	// One-way radar transmission loss factor due to atmospheric effects.
	AtmosTransmissionLoss float64 `json:"atmosTransmissionLoss"`
	// Average atmospheric angular width with no distortion from turbulence at an
	// optical sensor site in arcseconds.
	AvgAtmosSeeingConditions float64 `json:"avgAtmosSeeingConditions"`
	// Array of azimuth angles of a radar beam, in degrees. If this field is populated,
	// the associated beam(s) must be provided in the beamOrder field.
	AzAngs []float64 `json:"azAngs"`
	// Azimuth rate acquisition limit (radians/minute).
	AzimuthRate float64 `json:"azimuthRate"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of watts per square meter per steradian
	// (W/(m^2 str)) consistent with sensor detection bands.
	BackgroundSkyRadiance float64 `json:"backgroundSkyRadiance"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of visual magnitude consistent with sensor
	// detection bands.
	BackgroundSkyVisMag float64 `json:"backgroundSkyVisMag"`
	// Sensor band.
	Band string `json:"band"`
	// The total bandwidth, in megahertz, about the radar center frequency.
	Bandwidth float64 `json:"bandwidth"`
	// Array designating the beam order of provided values (e.g. vb1 for vertical beam
	// 1, ob1 for oblique beam 1, etc.). Required if any of the following array fields
	// are populated: azAngs, elAngs, radarPulseWidths, pulseRepPeriods, delayGates,
	// rangeGates, rangeSpacings, vertGateSpacings, vertGateWidths, specAvgSpectraNums,
	// tempMedFiltCodes, runMeanCodes, totRecNums, reqRecords, acceptSampleRanges.
	BeamOrder []string `json:"beamOrder"`
	// Number of radar beams used by the sensor.
	BeamQty int64 `json:"beamQty"`
	// The angle of the center of a phased array sensor.
	Boresight float64 `json:"boresight"`
	// The number of degrees off of the boresight for the sensor.
	BoresightOffAngle float64 `json:"boresightOffAngle"`
	// Weighted center wavelength for an optical sensor bandpass in micrometers. It is
	// the center wavelength in a weighted integral sense, accounting for the
	// sensitivity vs. wavelength curve for the sensor focal plane array.
	CenterWavelength float64 `json:"centerWavelength"`
	// Collapsing loss in decibels. Collapsing losses occur when two or more sources of
	// noise are added to the target signal. Examples include receiver bandwidth
	// mismatch with filtering bandwidth and elevation or azimuth beam collapse onto
	// position/height indicator displays.
	CollapsingLoss float64 `json:"collapsingLoss"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Threshold shear value beyond which one of the radial velocity values will be
	// rejected, measured in units of inverse second.
	CritShear float64 `json:"critShear"`
	// Current flowing through optical sensor focal plane electronics with a closed
	// shutter in electrons per second.
	DarkCurrent float64 `json:"darkCurrent"`
	// Array of time delay(s) for pulses from a radar beam to get to the first range
	// gate, in nanoseconds. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	DelayGates []float64 `json:"delayGates"`
	// Description of the equipment and data source.
	Description string `json:"description"`
	// Detection signal-to-noise ratio (SNR) threshold in decibels. This value is
	// typically lower than trackSNR.
	DetectSnr float64 `json:"detectSNR"`
	// Sensor duty cycle as a fraction of 1. Duty cycle is the fraction of time a
	// sensor is actively transmitting.
	DutyCycle float64 `json:"dutyCycle"`
	// Sensor Earth limb exclusion height in kilometers and is generally only applied
	// to space-based sensors. Some models used an earth exclusion angle instead, but
	// this assumes the sensor is in a circular orbit with constant altitude relative
	// to the earth. The limb exclusion height can be used for space-based sensors in
	// any orbit (assuming it is constant with sensor altitude). The limb height is
	// defined to be 0 at the surface of the earth.
	EarthLimbExclHgt float64 `json:"earthLimbExclHgt"`
	// Array of elevation angles of a radar beam, in degrees. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	ElAngs []float64 `json:"elAngs"`
	// Elevation rate acquisition limit (radians/minute).
	ElevationRateGeolm float64 `json:"elevationRateGeolm"`
	// Type of equipment used to take measurements.
	EquipmentType string `json:"equipmentType"`
	// The beam width of a Sensor's Fan (range). The values for this range from (0.0 to
	// PI).
	FanBeamWidth float64 `json:"fanBeamWidth"`
	// Number of Fast Fourier Transform (FFT) points used to convert time varying
	// signals into the frequency domain.
	Fft int64 `json:"fft"`
	// Maximum number of times the first guess was propagated in each gate before
	// failing the first guess check.
	FgpCrit int64 `json:"fgpCrit"`
	// Noise term, in decibels, that arises when a radar receiver filter has a
	// non-optimal bandwidth for an incoming signal (i.e., when it does not match the
	// pulse width).
	FilterMismatchFactor float64 `json:"filterMismatchFactor"`
	// F-number for an optical telescope. It is dimensionless and is defined as the
	// ratio of the focal length to the aperture width.
	FNum float64 `json:"fNum"`
	// For radar based sensors, the focal point elevation of the radar at the site, in
	// meters.
	FocalPoint float64 `json:"focalPoint"`
	// Horizontal field of view, in degrees.
	HFov float64 `json:"hFOV"`
	// Horizontal pixel resolution.
	HResPixels int64 `json:"hResPixels"`
	// For radar based sensors, K-factor is a relative indicator of refractivity that
	// infers the amount of radar beam bending due to atmosphere. (1<K<2).
	K float64 `json:"k"`
	// For Orbiting Sensors, First Card Azimuth limit #1 (left, degrees).
	LeftClockAngle float64 `json:"leftClockAngle"`
	// Leftmost GEO belt longitude limit for this sensor (if applicable).
	LeftGeoBeltLimit float64 `json:"leftGeoBeltLimit"`
	// Site where measurement is taken.
	Location string `json:"location"`
	// Aggregated radar range equation gain in decibels for maximum sensitivity. It is
	// a roll-up value for low fidelity modeling and is often the only sensitivity
	// value available for a radar system without access to detailed performance
	// parameters.
	LoopGain float64 `json:"loopGain"`
	// Lowest aspect angle of the full moon in degrees at which the sensor can achieve
	// full performance.
	LunarExclAngle float64 `json:"lunarExclAngle"`
	// Angle between magnetic north and true north at the sensor site, in degrees.
	MagDec float64 `json:"magDec"`
	// Absolute magnitude acquisition limit for optical sensors.
	MagnitudeLimit float64 `json:"magnitudeLimit"`
	// Max deviation angle of the sensor in degrees.
	MaxDeviationAngle float64 `json:"maxDeviationAngle"`
	// Maximum integration time per image frame in seconds for an optical sensor.
	MaxIntegrationTime float64 `json:"maxIntegrationTime"`
	// Maximum observable sensor range, in kilometers.
	MaxObservableRange float64 `json:"maxObservableRange"`
	// Maximum observable range limit in kilometers -- sensor cannot acquire beyond
	// this range.
	MaxRangeLimit float64 `json:"maxRangeLimit"`
	// Maximum wavelength detectable by an optical sensor in micrometers.
	MaxWavelength float64 `json:"maxWavelength"`
	// Minimum integration time per image frame in seconds for an optical sensor.
	MinIntegrationTime float64 `json:"minIntegrationTime"`
	// Minimum range measurement capability of the sensor, in kilometers.
	MinRangeLimit float64 `json:"minRangeLimit"`
	// Signal to Noise Ratio, in decibels. The values for this range from 0.0 - + 99.99
	// dB.
	MinSignalNoiseRatio float64 `json:"minSignalNoiseRatio"`
	// Minimum wavelength detectable by an optical sensor in micrometers.
	MinWavelength float64 `json:"minWavelength"`
	// Negative Range-rate/relative velocity limit (kilometers/second).
	NegativeRangeRateLimit float64 `json:"negativeRangeRateLimit"`
	// Noise figure for a radar system in decibels. This value may be used to compute
	// system noise when the system temperature is unavailable.
	NoiseFigure float64 `json:"noiseFigure"`
	// Number of pulses that are non-coherently integrated during detection processing.
	NonCoherentIntegratedPulses int64 `json:"nonCoherentIntegratedPulses"`
	// For radar based sensors, number of integrated pulses in a transmit cycle.
	NumIntegratedPulses int64 `json:"numIntegratedPulses"`
	// Number of integration frames for an optical sensor.
	NumIntegrationFrames int64 `json:"numIntegrationFrames"`
	// The number of optical integration mode characteristics provided in this record.
	// If provided, the numOpticalIntegrationModes value indicates the number of
	// elements in each of the opticalIntegrationTimes, opticalIntegrationAngularRates,
	// opticalIntegrationFrames, opticalIntegrationPixelBinnings, and
	// opticalIntegrationSNRs arrays.
	NumOpticalIntegrationModes int64 `json:"numOpticalIntegrationModes"`
	// The number of waveforms characteristics provided in this record. If provided,
	// the numWaveforms value indicates the number of elements in each of the
	// waveformPulseWidths, waveformBandWidths, waveformMinRanges, waveformMaxRanges,
	// and waveformLoopGains arrays.
	NumWaveforms int64 `json:"numWaveforms"`
	// Array containing the angular rate, in arcsec/sec, for each provided optical
	// integration mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationAngularRates []float64 `json:"opticalIntegrationAngularRates"`
	// Array containing the number of frames, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationFrames []float64 `json:"opticalIntegrationFrames"`
	// Array containing the pixel binning, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationPixelBinnings []float64 `json:"opticalIntegrationPixelBinnings"`
	// Array of optical integration modes signal to noise ratios. The number of
	// elements must be equal to the value indicated in numOpticalIntegrationModes.
	OpticalIntegrationSnRs []float64 `json:"opticalIntegrationSNRs"`
	// Array containing the time, in seconds, for each provided optical integration
	// mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationTimes []float64 `json:"opticalIntegrationTimes"`
	// Fraction of incident light transmitted to an optical sensor focal plane array.
	// The value is given as a fraction of 1, not as a percent.
	OpticalTransmission float64 `json:"opticalTransmission"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Two-way pattern absorption/propagation loss due to medium in decibels.
	PatternAbsorptionLoss float64 `json:"patternAbsorptionLoss"`
	// Losses from the beam shape, scanning, and pattern factor in decibels. These
	// losses occur when targets are not directly in line with a beam center. For space
	// surveillance, this will occur most often during sensor scanning.
	PatternScanLoss float64 `json:"patternScanLoss"`
	// Maximum instantaneous radar transmit power in watts for use in the radar range
	// equation.
	PeakPower float64 `json:"peakPower"`
	// Angular field-of-view covered by one pixel in a focal plane array in
	// microradians. The pixel is assumed to be a perfect square so that only a single
	// value is required.
	PixelInstantaneousFov float64 `json:"pixelInstantaneousFOV"`
	// Maximum number of electrons that can be collected in a single pixel on an
	// optical sensor focal plane array.
	PixelWellDepth int64 `json:"pixelWellDepth"`
	// Positive Range-rate/relative velocity limit (kilometers/second).
	PositiveRangeRateLimit float64 `json:"positiveRangeRateLimit"`
	// For radar based sensors, pulse repetition frequency (PRF), in hertz. Number of
	// new pulses transmitted per second.
	Prf float64 `json:"prf"`
	// Designated probability of detection at the signal-to-noise detection threshold.
	ProbDetectSnr float64 `json:"probDetectSNR"`
	// For radar based sensors, probability of the indication of the presence of a
	// radar target due to noise or interference.
	ProbFalseAlarm float64 `json:"probFalseAlarm"`
	// Array of interval(s) between the start of one radar pulse and the start of
	// another for a radar beam, in microseconds. If this field is populated, the
	// associated beam(s) must be provided in the beamOrder field.
	PulseRepPeriods []float64 `json:"pulseRepPeriods"`
	// Fraction of incident photons converted to electrons at the focal plane array.
	// This value is a decimal number between 0 and 1, inclusive.
	QuantumEff float64 `json:"quantumEff"`
	// Radar frequency in hertz, of the sensor (if a radar sensor).
	RadarFrequency float64 `json:"radarFrequency"`
	// Message data format transmitted by the sensor digitizer.
	RadarMessageFormat string `json:"radarMessageFormat"`
	// For radar based sensors, radar maximum unambiguous range, in kilometers.
	RadarMur float64 `json:"radarMUR"`
	// Array of transmit time(s) for a radar beam pulse, in microseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RadarPulseWidths []float64 `json:"radarPulseWidths"`
	// Radio frequency (if sensor is RF).
	RadioFrequency float64 `json:"radioFrequency"`
	// Losses due to the presence of a protective radome surrounding a radar sensor, in
	// decibels.
	RadomeLoss float64 `json:"radomeLoss"`
	// Array of the number(s) of discrete altitudes where return signals are sampled by
	// a radar beam. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	RangeGates []int64 `json:"rangeGates"`
	// Array of range gate spacing(s) for a radar beam, in nanoseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RangeSpacings []float64 `json:"rangeSpacings"`
	// Number of false-signal electrons generated by optical sensor focal plane
	// read-out electronics from photon-to-electron conversion during frame
	// integration. The units are in electrons RMS.
	ReadNoise int64 `json:"readNoise"`
	// Radar receive gain in decibels.
	ReceiveGain float64 `json:"receiveGain"`
	// Horizontal/azimuthal receive beamwidth for a radar in degrees.
	ReceiveHorizBeamWidth float64 `json:"receiveHorizBeamWidth"`
	// Aggregate radar receive loss, in decibels.
	ReceiveLoss float64 `json:"receiveLoss"`
	// Vertical/elevation receive beamwidth for a radar in degrees.
	ReceiveVertBeamWidth float64 `json:"receiveVertBeamWidth"`
	// Reference temperature for radar noise in Kelvin. A reference temperature is used
	// when the radar system temperature is unknown and is combined with the system
	// noise figure to estimate signal loss.
	RefTemp float64 `json:"refTemp"`
	// Array of the total number(s) of records required to meet consensus for a radar
	// beam. If this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	ReqRecords []int64 `json:"reqRecords"`
	// For Orbiting Sensors, First Card Azimuth limit #3 (right, degrees).
	RightClockAngle float64 `json:"rightClockAngle"`
	// Rightmost GEO belt longitude limit for this sensor (if applicable).
	RightGeoBeltLimit float64 `json:"rightGeoBeltLimit"`
	// Array of running mean code(s) used by radar data processing. The running mean
	// method involves taking a series of averages of different selections of the full
	// data set to smooth out short-term fluctuations in the data. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	RunMeanCodes []int64 `json:"runMeanCodes"`
	// Radar signal processing losses, in decibels.
	SignalProcessingLoss float64 `json:"signalProcessingLoss"`
	// Site code of the sensor.
	SiteCode string `json:"siteCode"`
	// Sensor and target position vector origins are at the center of the earth. The
	// sun vector origin is at the target position and points toward the sun. Any value
	// between 0 and 180 degrees is acceptable and is assumed to apply in both
	// directions (i.e., a solar exclusion angle of 30 degrees is understood to mean no
	// viewing for any angle between -30 deg and +30 deg).
	SolarExclAngle float64 `json:"solarExclAngle"`
	// Array of the number(s) of Doppler spectra used to process measurements from
	// radar. Spectral averaging involves combining multiple Doppler spectra acquired
	// to obtain a more accurate and representative spectrum. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	SpecAvgSpectraNums []int64 `json:"specAvgSpectraNums"`
	// For radar based sensors, expression of the radar system noise, aggregated as an
	// equivalent thermal noise value, in degrees Kelvin.
	SystemNoiseTemperature float64 `json:"systemNoiseTemperature"`
	// Maximum taskable range of the sensor, in kilometers.
	TaskableRange float64 `json:"taskableRange"`
	// Array of temporal median filter code(s) of a radar beam. Temporal median
	// filtering is a noise-reducing algorithm which involves replacing each data point
	// with the median value of a window of neighboring points over time. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	TempMedFiltCodes []int64 `json:"tempMedFiltCodes"`
	// Test number for the observed measurement.
	TestNumber string `json:"testNumber"`
	// Array of the total number(s) of records for a radar beam. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	TotRecNums []int64 `json:"totRecNums"`
	// For tower sensors, the physical height of the sensor tower, in meters.
	TowerHeight float64 `json:"towerHeight"`
	// Beginning track angle limit, in radians. Track angle is the angle between the
	// camera axis and the gimbal plane. Values range from 0 - PI/2.
	TrackAngle float64 `json:"trackAngle"`
	// Track signal-to-noise ratio (SNR) threshold in decibels. This value is typically
	// higher than detectSNR.
	TrackSnr float64 `json:"trackSNR"`
	// Radar transmit gain in decibels.
	TransmitGain float64 `json:"transmitGain"`
	// Horizontal/azimuthal transmit beamwidth for a radar in degrees.
	TransmitHorizBeamWidth float64 `json:"transmitHorizBeamWidth"`
	// Aggregate radar transmit loss, in decibels.
	TransmitLoss float64 `json:"transmitLoss"`
	// Radar transmit power in Watts.
	TransmitPower float64 `json:"transmitPower"`
	// Vertical/elevation transmit beamwidth for a radar in degrees.
	TransmitVertBeamWidth float64 `json:"transmitVertBeamWidth"`
	// True North correction for the sensor, in ACP (Azimunth Change Pulse) count.
	TrueNorthCorrector int64 `json:"trueNorthCorrector"`
	// Antenna true tilt, in degrees.
	TrueTilt float64 `json:"trueTilt"`
	// Twilight angle for ground-based optical sensors in degrees. A sensor cannot view
	// targets until the sun is below the twilight angle relative to the local horizon.
	// The sign of the angle is positive despite the sun elevation being negative after
	// local sunset. Typical values for the twilight angle are civil twilight (6
	// degrees), nautical twilight (12 degrees), and astronomical twilight (18
	// degrees).
	TwilightAngle float64 `json:"twilightAngle"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Flag indicating if a vertical radar beam was used in the wind calculation.
	VertBeamFlag bool `json:"vertBeamFlag"`
	// Array of vertical distance(s) between points where radar measurements are taken,
	// in meters. If this field is populated, the associated beam(s) must be provided
	// in the beamOrder field.
	VertGateSpacings []float64 `json:"vertGateSpacings"`
	// Array of width(s) of each location where radar measurements are taken, in
	// meters. If this field is populated, the associated beam(s) must be provided in
	// the beamOrder field.
	VertGateWidths []float64 `json:"vertGateWidths"`
	// Vertical field of view, in degrees.
	VFov float64 `json:"vFOV"`
	// Vertical pixel resolution.
	VResPixels int64 `json:"vResPixels"`
	// Array containing the bandwidth, in megahertz, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformBandwidths []float64 `json:"waveformBandwidths"`
	// Array containing the loop gain, in decibels, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field (10 SNR vs. 1 dBsm at 1000 km).
	WaveformLoopGains []float64 `json:"waveformLoopGains"`
	// Array containing the maximum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMaxRanges []float64 `json:"waveformMaxRanges"`
	// Array containing the minimum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMinRanges []float64 `json:"waveformMinRanges"`
	// Array containing the pulse width, in microseconds, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformPulseWidths []float64 `json:"waveformPulseWidths"`
	// Peformance zone-1 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MaxRange float64 `json:"z1MaxRange"`
	// Peformance zone-1 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MinRange float64 `json:"z1MinRange"`
	// Peformance zone-2 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MaxRange float64 `json:"z2MaxRange"`
	// Peformance zone-2 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MinRange float64 `json:"z2MinRange"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking           respjson.Field
		DataMode                        respjson.Field
		IDSensor                        respjson.Field
		Source                          respjson.Field
		ID                              respjson.Field
		AcceptSampleRanges              respjson.Field
		AnalogToDigitalBitSize          respjson.Field
		Aperture                        respjson.Field
		AsrScanRate                     respjson.Field
		AtmosReceiverLoss               respjson.Field
		AtmosTransmissionLoss           respjson.Field
		AvgAtmosSeeingConditions        respjson.Field
		AzAngs                          respjson.Field
		AzimuthRate                     respjson.Field
		BackgroundSkyRadiance           respjson.Field
		BackgroundSkyVisMag             respjson.Field
		Band                            respjson.Field
		Bandwidth                       respjson.Field
		BeamOrder                       respjson.Field
		BeamQty                         respjson.Field
		Boresight                       respjson.Field
		BoresightOffAngle               respjson.Field
		CenterWavelength                respjson.Field
		CollapsingLoss                  respjson.Field
		CreatedAt                       respjson.Field
		CreatedBy                       respjson.Field
		CritShear                       respjson.Field
		DarkCurrent                     respjson.Field
		DelayGates                      respjson.Field
		Description                     respjson.Field
		DetectSnr                       respjson.Field
		DutyCycle                       respjson.Field
		EarthLimbExclHgt                respjson.Field
		ElAngs                          respjson.Field
		ElevationRateGeolm              respjson.Field
		EquipmentType                   respjson.Field
		FanBeamWidth                    respjson.Field
		Fft                             respjson.Field
		FgpCrit                         respjson.Field
		FilterMismatchFactor            respjson.Field
		FNum                            respjson.Field
		FocalPoint                      respjson.Field
		HFov                            respjson.Field
		HResPixels                      respjson.Field
		K                               respjson.Field
		LeftClockAngle                  respjson.Field
		LeftGeoBeltLimit                respjson.Field
		Location                        respjson.Field
		LoopGain                        respjson.Field
		LunarExclAngle                  respjson.Field
		MagDec                          respjson.Field
		MagnitudeLimit                  respjson.Field
		MaxDeviationAngle               respjson.Field
		MaxIntegrationTime              respjson.Field
		MaxObservableRange              respjson.Field
		MaxRangeLimit                   respjson.Field
		MaxWavelength                   respjson.Field
		MinIntegrationTime              respjson.Field
		MinRangeLimit                   respjson.Field
		MinSignalNoiseRatio             respjson.Field
		MinWavelength                   respjson.Field
		NegativeRangeRateLimit          respjson.Field
		NoiseFigure                     respjson.Field
		NonCoherentIntegratedPulses     respjson.Field
		NumIntegratedPulses             respjson.Field
		NumIntegrationFrames            respjson.Field
		NumOpticalIntegrationModes      respjson.Field
		NumWaveforms                    respjson.Field
		OpticalIntegrationAngularRates  respjson.Field
		OpticalIntegrationFrames        respjson.Field
		OpticalIntegrationPixelBinnings respjson.Field
		OpticalIntegrationSnRs          respjson.Field
		OpticalIntegrationTimes         respjson.Field
		OpticalTransmission             respjson.Field
		OrigNetwork                     respjson.Field
		PatternAbsorptionLoss           respjson.Field
		PatternScanLoss                 respjson.Field
		PeakPower                       respjson.Field
		PixelInstantaneousFov           respjson.Field
		PixelWellDepth                  respjson.Field
		PositiveRangeRateLimit          respjson.Field
		Prf                             respjson.Field
		ProbDetectSnr                   respjson.Field
		ProbFalseAlarm                  respjson.Field
		PulseRepPeriods                 respjson.Field
		QuantumEff                      respjson.Field
		RadarFrequency                  respjson.Field
		RadarMessageFormat              respjson.Field
		RadarMur                        respjson.Field
		RadarPulseWidths                respjson.Field
		RadioFrequency                  respjson.Field
		RadomeLoss                      respjson.Field
		RangeGates                      respjson.Field
		RangeSpacings                   respjson.Field
		ReadNoise                       respjson.Field
		ReceiveGain                     respjson.Field
		ReceiveHorizBeamWidth           respjson.Field
		ReceiveLoss                     respjson.Field
		ReceiveVertBeamWidth            respjson.Field
		RefTemp                         respjson.Field
		ReqRecords                      respjson.Field
		RightClockAngle                 respjson.Field
		RightGeoBeltLimit               respjson.Field
		RunMeanCodes                    respjson.Field
		SignalProcessingLoss            respjson.Field
		SiteCode                        respjson.Field
		SolarExclAngle                  respjson.Field
		SpecAvgSpectraNums              respjson.Field
		SystemNoiseTemperature          respjson.Field
		TaskableRange                   respjson.Field
		TempMedFiltCodes                respjson.Field
		TestNumber                      respjson.Field
		TotRecNums                      respjson.Field
		TowerHeight                     respjson.Field
		TrackAngle                      respjson.Field
		TrackSnr                        respjson.Field
		TransmitGain                    respjson.Field
		TransmitHorizBeamWidth          respjson.Field
		TransmitLoss                    respjson.Field
		TransmitPower                   respjson.Field
		TransmitVertBeamWidth           respjson.Field
		TrueNorthCorrector              respjson.Field
		TrueTilt                        respjson.Field
		TwilightAngle                   respjson.Field
		UpdatedAt                       respjson.Field
		UpdatedBy                       respjson.Field
		VertBeamFlag                    respjson.Field
		VertGateSpacings                respjson.Field
		VertGateWidths                  respjson.Field
		VFov                            respjson.Field
		VResPixels                      respjson.Field
		WaveformBandwidths              respjson.Field
		WaveformLoopGains               respjson.Field
		WaveformMaxRanges               respjson.Field
		WaveformMinRanges               respjson.Field
		WaveformPulseWidths             respjson.Field
		Z1MaxRange                      respjson.Field
		Z1MinRange                      respjson.Field
		Z2MaxRange                      respjson.Field
		Z2MinRange                      respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorTupleResponseSensorcharacteristic) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseSensorcharacteristic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sensorlimits define 0 to many limits of a particular sensor in terms of
// observation coverage of on-orbit objects.
type SensorTupleResponseSensorlimitsCollection struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the target sensor object.
	IDSensor string `json:"idSensor"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensorLimits string `json:"idSensorLimits"`
	// Leftmost or minimum lower azimuth within this limit. Interpreted according to
	// site types as lower left azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower left azimuth limits would be equal. (degrees).
	LowerLeftAzimuthLimit float64 `json:"lowerLeftAzimuthLimit"`
	// Minimum or lower elevation within this limit. Interpreted according to site
	// types as minimum elevation angle, constant elevation or fan beam centerline.
	// (Degrees).
	LowerLeftElevationLimit float64 `json:"lowerLeftElevationLimit"`
	// Rightmost or maximum lower azimuth within this limit. Interpreted according to
	// site types as 2nd lower azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower right azimuth limits would be equal. (degrees).
	LowerRightAzimuthLimit float64 `json:"lowerRightAzimuthLimit"`
	// Minimum or lower right elevation within this limit. Interpreted according to
	// site types as minimum right elevation angle, constant elevation or fan beam
	// centerline. If the limit rectangle is parallel to the horizon, the left and
	// right lower elevation limits would be equal. (Degrees).
	LowerRightElevationLimit float64 `json:"lowerRightElevationLimit"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Leftmost or minimum upper azimuth within this sensor limit. Interpreted
	// according to site types as beginning upper azimuth limit, left-hand upper
	// boundary limit. If the limit rectangle is parallel to the horizon, the upper and
	// lower left azimuth limits would be equal. (in degrees).
	UpperLeftAzimuthLimit float64 `json:"upperLeftAzimuthLimit"`
	// Maximum or upper elevation within this limit. Interpreted according to site
	// types as maximum elevation angle, half the apex of conical observation pattern
	// or star. (Degrees).
	UpperLeftElevationLimit float64 `json:"upperLeftElevationLimit"`
	// Rightmost or maximum upper azimuth within this limit. Interpreted according to
	// site types as 2nd azimuth limit elevation angle of axis of conical observation
	// pattern. If the limit rectangle is parallel to the horizon, the upper and lower
	// right azimuth limits would be equal. (degrees).
	UpperRightAzimuthLimit float64 `json:"upperRightAzimuthLimit"`
	// Maximum or upper right elevation within this limit. Interpreted according to
	// site types as maximum rightmost elevation angle, half the apex of conical
	// observation pattern or star. If the limit rectangle is parallel to the horizon,
	// the left and right upper elevation limits would be equal. (Degrees).
	UpperRightElevationLimit float64 `json:"upperRightElevationLimit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		Source                   respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		IDSensor                 respjson.Field
		IDSensorLimits           respjson.Field
		LowerLeftAzimuthLimit    respjson.Field
		LowerLeftElevationLimit  respjson.Field
		LowerRightAzimuthLimit   respjson.Field
		LowerRightElevationLimit respjson.Field
		OrigNetwork              respjson.Field
		UpdatedAt                respjson.Field
		UpdatedBy                respjson.Field
		UpperLeftAzimuthLimit    respjson.Field
		UpperLeftElevationLimit  respjson.Field
		UpperRightAzimuthLimit   respjson.Field
		UpperRightElevationLimit respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorTupleResponseSensorlimitsCollection) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseSensorlimitsCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorTupleResponseSensorObservationType struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The observation measurement type produced by a sensor.
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataMode    respjson.Field
		Source      respjson.Field
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		OrigNetwork respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorTupleResponseSensorObservationType) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseSensorObservationType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SensorStats contain statistics on sensors related to observation production such
// as last reported observation time.
type SensorTupleResponseSensorStat struct {
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
	// Unique ID of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Time of last reported observation in ISO 8601 UTC with microsecond precision.
	LastObTime time.Time `json:"lastObTime" format:"date-time"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSensor              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		LastObTime            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorTupleResponseSensorStat) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseSensorStat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorTupleResponseSensorType struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID int64 `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The specific sensor type and/or surveillance capability of this sensor.
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataMode    respjson.Field
		Source      respjson.Field
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		OrigNetwork respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorTupleResponseSensorType) RawJSON() string { return r.JSON.raw }
func (r *SensorTupleResponseSensorType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorNewParams struct {
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
	DataMode SensorNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Optional flag indicating if the sensor is active.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Optional US Air Force identifier for the sensor/ASR site, typically for air
	// surveillance radar (ASR) sensors.
	AfID param.Opt[string] `json:"afId,omitzero"`
	// The sensor type at the site. Optional field, intended primarily for ASRs.
	AsrType param.Opt[string] `json:"asrType,omitzero"`
	// Optional dissemination control required for accessing data (e.g observations)
	// produced by this sensor. This is typically a proprietary data owner control for
	// commercial sensors.
	DataControl param.Opt[string] `json:"dataControl,omitzero"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber param.Opt[int64] `json:"sensorNumber,omitzero"`
	// Optional short name for the sensor.
	ShortName param.Opt[string] `json:"shortName,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity SensorNewParamsEntity `json:"entity,omitzero"`
	// Collection of Sensorcharacteristics which define characteristics and
	// capabilities of a sensor.
	Sensorcharacteristics []SensorNewParamsSensorcharacteristic `json:"sensorcharacteristics,omitzero"`
	// Sensorlimits define 0 to many limits of a particular sensor in terms of
	// observation coverage of on-orbit objects.
	SensorlimitsCollection []SensorNewParamsSensorlimitsCollection `json:"sensorlimitsCollection,omitzero"`
	SensorObservationType  SensorNewParamsSensorObservationType    `json:"sensorObservationType,omitzero"`
	// Collection of SensorStats which contain statistics of a sensor.
	SensorStats []SensorNewParamsSensorStat `json:"sensorStats,omitzero"`
	SensorType  SensorNewParamsSensorType   `json:"sensorType,omitzero"`
	paramObj
}

func (r SensorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SensorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorNewParams) UnmarshalJSON(data []byte) error {
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
type SensorNewParamsDataMode string

const (
	SensorNewParamsDataModeReal      SensorNewParamsDataMode = "REAL"
	SensorNewParamsDataModeTest      SensorNewParamsDataMode = "TEST"
	SensorNewParamsDataModeSimulated SensorNewParamsDataMode = "SIMULATED"
	SensorNewParamsDataModeExercise  SensorNewParamsDataMode = "EXERCISE"
)

// An entity is a generic representation of any object within a space/SSA system
// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
// entity can have an operating unit, a location (if terrestrial), and statuses.
//
// The properties ClassificationMarking, DataMode, Name, Source, Type are required.
type SensorNewParamsEntity struct {
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
	// Unique entity name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
	// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
	//
	// Any of "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT",
	// "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL".
	Type string `json:"type,omitzero,required"`
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
	// Terrestrial identifier of this entity, if applicable.
	TerrestrialID param.Opt[string] `json:"terrestrialId,omitzero"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LocationIngestParam `json:"location,omitzero"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit SensorNewParamsEntityOnOrbit `json:"onOrbit,omitzero"`
	// Type of organization which owns this entity (e.g. Commercial, Government,
	// Academic, Consortium, etc).
	//
	// Any of "Commercial", "Government", "Academic", "Consortium", "Other".
	OwnerType string `json:"ownerType,omitzero"`
	// List of URLs to additional details/documents for this entity.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

func (r SensorNewParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow SensorNewParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorNewParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorNewParamsEntity](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SensorNewParamsEntity](
		"type", "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT", "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL",
	)
	apijson.RegisterFieldValidator[SensorNewParamsEntity](
		"ownerType", "Commercial", "Government", "Academic", "Consortium", "Other",
	)
}

// Model object representing on-orbit objects or satellites in the system.
//
// The properties ClassificationMarking, DataMode, SatNo, Source are required.
type SensorNewParamsEntityOnOrbit struct {
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

func (r SensorNewParamsEntityOnOrbit) MarshalJSON() (data []byte, err error) {
	type shadow SensorNewParamsEntityOnOrbit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorNewParamsEntityOnOrbit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorNewParamsEntityOnOrbit](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SensorNewParamsEntityOnOrbit](
		"category", "Unknown", "On-Orbit", "Decayed", "Cataloged Without State", "Launch Nominal", "Analyst Satellite", "Cislunar", "Lunar", "Hyperbolic", "Heliocentric", "Interplanetary", "Lagrangian", "Docked",
	)
	apijson.RegisterFieldValidator[SensorNewParamsEntityOnOrbit](
		"objectType", "ROCKET BODY", "DEBRIS", "PAYLOAD", "PLATFORM", "MANNED", "UNKNOWN",
	)
}

// Model representation of characteristics and capabilities of a sensor.
//
// The properties ClassificationMarking, DataMode, IDSensor, Source are required.
type SensorNewParamsSensorcharacteristic struct {
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
	// Unique identifier of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Number of bits used in the conversion from analog electrons in a pixel well to a
	// digital number. The digital number has a maximum value of 2^N, where N is the
	// number of bits.
	AnalogToDigitalBitSize param.Opt[int64] `json:"analogToDigitalBitSize,omitzero"`
	// Optical sensor camera aperture.
	Aperture param.Opt[float64] `json:"aperture,omitzero"`
	// For ASR (Air Surveillance Radar) sensors, the scan (360 deg sweep) rate of the
	// radar, in scans/minute.
	AsrScanRate param.Opt[float64] `json:"asrScanRate,omitzero"`
	// One-way radar receiver loss factor due to atmospheric effects. This value will
	// often be the same as the corresponding transmission factor but may be different
	// for bistatic systems.
	AtmosReceiverLoss param.Opt[float64] `json:"atmosReceiverLoss,omitzero"`
	// One-way radar transmission loss factor due to atmospheric effects.
	AtmosTransmissionLoss param.Opt[float64] `json:"atmosTransmissionLoss,omitzero"`
	// Average atmospheric angular width with no distortion from turbulence at an
	// optical sensor site in arcseconds.
	AvgAtmosSeeingConditions param.Opt[float64] `json:"avgAtmosSeeingConditions,omitzero"`
	// Azimuth rate acquisition limit (radians/minute).
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of watts per square meter per steradian
	// (W/(m^2 str)) consistent with sensor detection bands.
	BackgroundSkyRadiance param.Opt[float64] `json:"backgroundSkyRadiance,omitzero"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of visual magnitude consistent with sensor
	// detection bands.
	BackgroundSkyVisMag param.Opt[float64] `json:"backgroundSkyVisMag,omitzero"`
	// Sensor band.
	Band param.Opt[string] `json:"band,omitzero"`
	// The total bandwidth, in megahertz, about the radar center frequency.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Number of radar beams used by the sensor.
	BeamQty param.Opt[int64] `json:"beamQty,omitzero"`
	// The angle of the center of a phased array sensor.
	Boresight param.Opt[float64] `json:"boresight,omitzero"`
	// The number of degrees off of the boresight for the sensor.
	BoresightOffAngle param.Opt[float64] `json:"boresightOffAngle,omitzero"`
	// Weighted center wavelength for an optical sensor bandpass in micrometers. It is
	// the center wavelength in a weighted integral sense, accounting for the
	// sensitivity vs. wavelength curve for the sensor focal plane array.
	CenterWavelength param.Opt[float64] `json:"centerWavelength,omitzero"`
	// Collapsing loss in decibels. Collapsing losses occur when two or more sources of
	// noise are added to the target signal. Examples include receiver bandwidth
	// mismatch with filtering bandwidth and elevation or azimuth beam collapse onto
	// position/height indicator displays.
	CollapsingLoss param.Opt[float64] `json:"collapsingLoss,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Threshold shear value beyond which one of the radial velocity values will be
	// rejected, measured in units of inverse second.
	CritShear param.Opt[float64] `json:"critShear,omitzero"`
	// Current flowing through optical sensor focal plane electronics with a closed
	// shutter in electrons per second.
	DarkCurrent param.Opt[float64] `json:"darkCurrent,omitzero"`
	// Description of the equipment and data source.
	Description param.Opt[string] `json:"description,omitzero"`
	// Detection signal-to-noise ratio (SNR) threshold in decibels. This value is
	// typically lower than trackSNR.
	DetectSnr param.Opt[float64] `json:"detectSNR,omitzero"`
	// Sensor duty cycle as a fraction of 1. Duty cycle is the fraction of time a
	// sensor is actively transmitting.
	DutyCycle param.Opt[float64] `json:"dutyCycle,omitzero"`
	// Sensor Earth limb exclusion height in kilometers and is generally only applied
	// to space-based sensors. Some models used an earth exclusion angle instead, but
	// this assumes the sensor is in a circular orbit with constant altitude relative
	// to the earth. The limb exclusion height can be used for space-based sensors in
	// any orbit (assuming it is constant with sensor altitude). The limb height is
	// defined to be 0 at the surface of the earth.
	EarthLimbExclHgt param.Opt[float64] `json:"earthLimbExclHgt,omitzero"`
	// Elevation rate acquisition limit (radians/minute).
	ElevationRateGeolm param.Opt[float64] `json:"elevationRateGeolm,omitzero"`
	// Type of equipment used to take measurements.
	EquipmentType param.Opt[string] `json:"equipmentType,omitzero"`
	// The beam width of a Sensor's Fan (range). The values for this range from (0.0 to
	// PI).
	FanBeamWidth param.Opt[float64] `json:"fanBeamWidth,omitzero"`
	// Number of Fast Fourier Transform (FFT) points used to convert time varying
	// signals into the frequency domain.
	Fft param.Opt[int64] `json:"fft,omitzero"`
	// Maximum number of times the first guess was propagated in each gate before
	// failing the first guess check.
	FgpCrit param.Opt[int64] `json:"fgpCrit,omitzero"`
	// Noise term, in decibels, that arises when a radar receiver filter has a
	// non-optimal bandwidth for an incoming signal (i.e., when it does not match the
	// pulse width).
	FilterMismatchFactor param.Opt[float64] `json:"filterMismatchFactor,omitzero"`
	// F-number for an optical telescope. It is dimensionless and is defined as the
	// ratio of the focal length to the aperture width.
	FNum param.Opt[float64] `json:"fNum,omitzero"`
	// For radar based sensors, the focal point elevation of the radar at the site, in
	// meters.
	FocalPoint param.Opt[float64] `json:"focalPoint,omitzero"`
	// Horizontal field of view, in degrees.
	HFov param.Opt[float64] `json:"hFOV,omitzero"`
	// Horizontal pixel resolution.
	HResPixels param.Opt[int64] `json:"hResPixels,omitzero"`
	// For radar based sensors, K-factor is a relative indicator of refractivity that
	// infers the amount of radar beam bending due to atmosphere. (1<K<2).
	K param.Opt[float64] `json:"k,omitzero"`
	// For Orbiting Sensors, First Card Azimuth limit #1 (left, degrees).
	LeftClockAngle param.Opt[float64] `json:"leftClockAngle,omitzero"`
	// Leftmost GEO belt longitude limit for this sensor (if applicable).
	LeftGeoBeltLimit param.Opt[float64] `json:"leftGeoBeltLimit,omitzero"`
	// Site where measurement is taken.
	Location param.Opt[string] `json:"location,omitzero"`
	// Aggregated radar range equation gain in decibels for maximum sensitivity. It is
	// a roll-up value for low fidelity modeling and is often the only sensitivity
	// value available for a radar system without access to detailed performance
	// parameters.
	LoopGain param.Opt[float64] `json:"loopGain,omitzero"`
	// Lowest aspect angle of the full moon in degrees at which the sensor can achieve
	// full performance.
	LunarExclAngle param.Opt[float64] `json:"lunarExclAngle,omitzero"`
	// Angle between magnetic north and true north at the sensor site, in degrees.
	MagDec param.Opt[float64] `json:"magDec,omitzero"`
	// Absolute magnitude acquisition limit for optical sensors.
	MagnitudeLimit param.Opt[float64] `json:"magnitudeLimit,omitzero"`
	// Max deviation angle of the sensor in degrees.
	MaxDeviationAngle param.Opt[float64] `json:"maxDeviationAngle,omitzero"`
	// Maximum integration time per image frame in seconds for an optical sensor.
	MaxIntegrationTime param.Opt[float64] `json:"maxIntegrationTime,omitzero"`
	// Maximum observable sensor range, in kilometers.
	MaxObservableRange param.Opt[float64] `json:"maxObservableRange,omitzero"`
	// Maximum observable range limit in kilometers -- sensor cannot acquire beyond
	// this range.
	MaxRangeLimit param.Opt[float64] `json:"maxRangeLimit,omitzero"`
	// Maximum wavelength detectable by an optical sensor in micrometers.
	MaxWavelength param.Opt[float64] `json:"maxWavelength,omitzero"`
	// Minimum integration time per image frame in seconds for an optical sensor.
	MinIntegrationTime param.Opt[float64] `json:"minIntegrationTime,omitzero"`
	// Minimum range measurement capability of the sensor, in kilometers.
	MinRangeLimit param.Opt[float64] `json:"minRangeLimit,omitzero"`
	// Signal to Noise Ratio, in decibels. The values for this range from 0.0 - + 99.99
	// dB.
	MinSignalNoiseRatio param.Opt[float64] `json:"minSignalNoiseRatio,omitzero"`
	// Minimum wavelength detectable by an optical sensor in micrometers.
	MinWavelength param.Opt[float64] `json:"minWavelength,omitzero"`
	// Negative Range-rate/relative velocity limit (kilometers/second).
	NegativeRangeRateLimit param.Opt[float64] `json:"negativeRangeRateLimit,omitzero"`
	// Noise figure for a radar system in decibels. This value may be used to compute
	// system noise when the system temperature is unavailable.
	NoiseFigure param.Opt[float64] `json:"noiseFigure,omitzero"`
	// Number of pulses that are non-coherently integrated during detection processing.
	NonCoherentIntegratedPulses param.Opt[int64] `json:"nonCoherentIntegratedPulses,omitzero"`
	// For radar based sensors, number of integrated pulses in a transmit cycle.
	NumIntegratedPulses param.Opt[int64] `json:"numIntegratedPulses,omitzero"`
	// Number of integration frames for an optical sensor.
	NumIntegrationFrames param.Opt[int64] `json:"numIntegrationFrames,omitzero"`
	// The number of optical integration mode characteristics provided in this record.
	// If provided, the numOpticalIntegrationModes value indicates the number of
	// elements in each of the opticalIntegrationTimes, opticalIntegrationAngularRates,
	// opticalIntegrationFrames, opticalIntegrationPixelBinnings, and
	// opticalIntegrationSNRs arrays.
	NumOpticalIntegrationModes param.Opt[int64] `json:"numOpticalIntegrationModes,omitzero"`
	// The number of waveforms characteristics provided in this record. If provided,
	// the numWaveforms value indicates the number of elements in each of the
	// waveformPulseWidths, waveformBandWidths, waveformMinRanges, waveformMaxRanges,
	// and waveformLoopGains arrays.
	NumWaveforms param.Opt[int64] `json:"numWaveforms,omitzero"`
	// Fraction of incident light transmitted to an optical sensor focal plane array.
	// The value is given as a fraction of 1, not as a percent.
	OpticalTransmission param.Opt[float64] `json:"opticalTransmission,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Two-way pattern absorption/propagation loss due to medium in decibels.
	PatternAbsorptionLoss param.Opt[float64] `json:"patternAbsorptionLoss,omitzero"`
	// Losses from the beam shape, scanning, and pattern factor in decibels. These
	// losses occur when targets are not directly in line with a beam center. For space
	// surveillance, this will occur most often during sensor scanning.
	PatternScanLoss param.Opt[float64] `json:"patternScanLoss,omitzero"`
	// Maximum instantaneous radar transmit power in watts for use in the radar range
	// equation.
	PeakPower param.Opt[float64] `json:"peakPower,omitzero"`
	// Angular field-of-view covered by one pixel in a focal plane array in
	// microradians. The pixel is assumed to be a perfect square so that only a single
	// value is required.
	PixelInstantaneousFov param.Opt[float64] `json:"pixelInstantaneousFOV,omitzero"`
	// Maximum number of electrons that can be collected in a single pixel on an
	// optical sensor focal plane array.
	PixelWellDepth param.Opt[int64] `json:"pixelWellDepth,omitzero"`
	// Positive Range-rate/relative velocity limit (kilometers/second).
	PositiveRangeRateLimit param.Opt[float64] `json:"positiveRangeRateLimit,omitzero"`
	// For radar based sensors, pulse repetition frequency (PRF), in hertz. Number of
	// new pulses transmitted per second.
	Prf param.Opt[float64] `json:"prf,omitzero"`
	// Designated probability of detection at the signal-to-noise detection threshold.
	ProbDetectSnr param.Opt[float64] `json:"probDetectSNR,omitzero"`
	// For radar based sensors, probability of the indication of the presence of a
	// radar target due to noise or interference.
	ProbFalseAlarm param.Opt[float64] `json:"probFalseAlarm,omitzero"`
	// Fraction of incident photons converted to electrons at the focal plane array.
	// This value is a decimal number between 0 and 1, inclusive.
	QuantumEff param.Opt[float64] `json:"quantumEff,omitzero"`
	// Radar frequency in hertz, of the sensor (if a radar sensor).
	RadarFrequency param.Opt[float64] `json:"radarFrequency,omitzero"`
	// Message data format transmitted by the sensor digitizer.
	RadarMessageFormat param.Opt[string] `json:"radarMessageFormat,omitzero"`
	// For radar based sensors, radar maximum unambiguous range, in kilometers.
	RadarMur param.Opt[float64] `json:"radarMUR,omitzero"`
	// Radio frequency (if sensor is RF).
	RadioFrequency param.Opt[float64] `json:"radioFrequency,omitzero"`
	// Losses due to the presence of a protective radome surrounding a radar sensor, in
	// decibels.
	RadomeLoss param.Opt[float64] `json:"radomeLoss,omitzero"`
	// Number of false-signal electrons generated by optical sensor focal plane
	// read-out electronics from photon-to-electron conversion during frame
	// integration. The units are in electrons RMS.
	ReadNoise param.Opt[int64] `json:"readNoise,omitzero"`
	// Radar receive gain in decibels.
	ReceiveGain param.Opt[float64] `json:"receiveGain,omitzero"`
	// Horizontal/azimuthal receive beamwidth for a radar in degrees.
	ReceiveHorizBeamWidth param.Opt[float64] `json:"receiveHorizBeamWidth,omitzero"`
	// Aggregate radar receive loss, in decibels.
	ReceiveLoss param.Opt[float64] `json:"receiveLoss,omitzero"`
	// Vertical/elevation receive beamwidth for a radar in degrees.
	ReceiveVertBeamWidth param.Opt[float64] `json:"receiveVertBeamWidth,omitzero"`
	// Reference temperature for radar noise in Kelvin. A reference temperature is used
	// when the radar system temperature is unknown and is combined with the system
	// noise figure to estimate signal loss.
	RefTemp param.Opt[float64] `json:"refTemp,omitzero"`
	// For Orbiting Sensors, First Card Azimuth limit #3 (right, degrees).
	RightClockAngle param.Opt[float64] `json:"rightClockAngle,omitzero"`
	// Rightmost GEO belt longitude limit for this sensor (if applicable).
	RightGeoBeltLimit param.Opt[float64] `json:"rightGeoBeltLimit,omitzero"`
	// Radar signal processing losses, in decibels.
	SignalProcessingLoss param.Opt[float64] `json:"signalProcessingLoss,omitzero"`
	// Site code of the sensor.
	SiteCode param.Opt[string] `json:"siteCode,omitzero"`
	// Sensor and target position vector origins are at the center of the earth. The
	// sun vector origin is at the target position and points toward the sun. Any value
	// between 0 and 180 degrees is acceptable and is assumed to apply in both
	// directions (i.e., a solar exclusion angle of 30 degrees is understood to mean no
	// viewing for any angle between -30 deg and +30 deg).
	SolarExclAngle param.Opt[float64] `json:"solarExclAngle,omitzero"`
	// For radar based sensors, expression of the radar system noise, aggregated as an
	// equivalent thermal noise value, in degrees Kelvin.
	SystemNoiseTemperature param.Opt[float64] `json:"systemNoiseTemperature,omitzero"`
	// Maximum taskable range of the sensor, in kilometers.
	TaskableRange param.Opt[float64] `json:"taskableRange,omitzero"`
	// Test number for the observed measurement.
	TestNumber param.Opt[string] `json:"testNumber,omitzero"`
	// For tower sensors, the physical height of the sensor tower, in meters.
	TowerHeight param.Opt[float64] `json:"towerHeight,omitzero"`
	// Beginning track angle limit, in radians. Track angle is the angle between the
	// camera axis and the gimbal plane. Values range from 0 - PI/2.
	TrackAngle param.Opt[float64] `json:"trackAngle,omitzero"`
	// Track signal-to-noise ratio (SNR) threshold in decibels. This value is typically
	// higher than detectSNR.
	TrackSnr param.Opt[float64] `json:"trackSNR,omitzero"`
	// Radar transmit gain in decibels.
	TransmitGain param.Opt[float64] `json:"transmitGain,omitzero"`
	// Horizontal/azimuthal transmit beamwidth for a radar in degrees.
	TransmitHorizBeamWidth param.Opt[float64] `json:"transmitHorizBeamWidth,omitzero"`
	// Aggregate radar transmit loss, in decibels.
	TransmitLoss param.Opt[float64] `json:"transmitLoss,omitzero"`
	// Radar transmit power in Watts.
	TransmitPower param.Opt[float64] `json:"transmitPower,omitzero"`
	// Vertical/elevation transmit beamwidth for a radar in degrees.
	TransmitVertBeamWidth param.Opt[float64] `json:"transmitVertBeamWidth,omitzero"`
	// True North correction for the sensor, in ACP (Azimunth Change Pulse) count.
	TrueNorthCorrector param.Opt[int64] `json:"trueNorthCorrector,omitzero"`
	// Antenna true tilt, in degrees.
	TrueTilt param.Opt[float64] `json:"trueTilt,omitzero"`
	// Twilight angle for ground-based optical sensors in degrees. A sensor cannot view
	// targets until the sun is below the twilight angle relative to the local horizon.
	// The sign of the angle is positive despite the sun elevation being negative after
	// local sunset. Typical values for the twilight angle are civil twilight (6
	// degrees), nautical twilight (12 degrees), and astronomical twilight (18
	// degrees).
	TwilightAngle param.Opt[float64] `json:"twilightAngle,omitzero"`
	// Flag indicating if a vertical radar beam was used in the wind calculation.
	VertBeamFlag param.Opt[bool] `json:"vertBeamFlag,omitzero"`
	// Vertical field of view, in degrees.
	VFov param.Opt[float64] `json:"vFOV,omitzero"`
	// Vertical pixel resolution.
	VResPixels param.Opt[int64] `json:"vResPixels,omitzero"`
	// Peformance zone-1 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MaxRange param.Opt[float64] `json:"z1MaxRange,omitzero"`
	// Peformance zone-1 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MinRange param.Opt[float64] `json:"z1MinRange,omitzero"`
	// Peformance zone-2 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MaxRange param.Opt[float64] `json:"z2MaxRange,omitzero"`
	// Peformance zone-2 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MinRange param.Opt[float64] `json:"z2MinRange,omitzero"`
	// Array of measurement range(s) where radar samples must fall to be acceptable. If
	// this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	AcceptSampleRanges []float64 `json:"acceptSampleRanges,omitzero"`
	// Array of azimuth angles of a radar beam, in degrees. If this field is populated,
	// the associated beam(s) must be provided in the beamOrder field.
	AzAngs []float64 `json:"azAngs,omitzero"`
	// Array designating the beam order of provided values (e.g. vb1 for vertical beam
	// 1, ob1 for oblique beam 1, etc.). Required if any of the following array fields
	// are populated: azAngs, elAngs, radarPulseWidths, pulseRepPeriods, delayGates,
	// rangeGates, rangeSpacings, vertGateSpacings, vertGateWidths, specAvgSpectraNums,
	// tempMedFiltCodes, runMeanCodes, totRecNums, reqRecords, acceptSampleRanges.
	BeamOrder []string `json:"beamOrder,omitzero"`
	// Array of time delay(s) for pulses from a radar beam to get to the first range
	// gate, in nanoseconds. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	DelayGates []float64 `json:"delayGates,omitzero"`
	// Array of elevation angles of a radar beam, in degrees. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	ElAngs []float64 `json:"elAngs,omitzero"`
	// Array containing the angular rate, in arcsec/sec, for each provided optical
	// integration mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationAngularRates []float64 `json:"opticalIntegrationAngularRates,omitzero"`
	// Array containing the number of frames, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationFrames []float64 `json:"opticalIntegrationFrames,omitzero"`
	// Array containing the pixel binning, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationPixelBinnings []float64 `json:"opticalIntegrationPixelBinnings,omitzero"`
	// Array of optical integration modes signal to noise ratios. The number of
	// elements must be equal to the value indicated in numOpticalIntegrationModes.
	OpticalIntegrationSnRs []float64 `json:"opticalIntegrationSNRs,omitzero"`
	// Array containing the time, in seconds, for each provided optical integration
	// mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationTimes []float64 `json:"opticalIntegrationTimes,omitzero"`
	// Array of interval(s) between the start of one radar pulse and the start of
	// another for a radar beam, in microseconds. If this field is populated, the
	// associated beam(s) must be provided in the beamOrder field.
	PulseRepPeriods []float64 `json:"pulseRepPeriods,omitzero"`
	// Array of transmit time(s) for a radar beam pulse, in microseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RadarPulseWidths []float64 `json:"radarPulseWidths,omitzero"`
	// Array of the number(s) of discrete altitudes where return signals are sampled by
	// a radar beam. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	RangeGates []int64 `json:"rangeGates,omitzero"`
	// Array of range gate spacing(s) for a radar beam, in nanoseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RangeSpacings []float64 `json:"rangeSpacings,omitzero"`
	// Array of the total number(s) of records required to meet consensus for a radar
	// beam. If this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	ReqRecords []int64 `json:"reqRecords,omitzero"`
	// Array of running mean code(s) used by radar data processing. The running mean
	// method involves taking a series of averages of different selections of the full
	// data set to smooth out short-term fluctuations in the data. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	RunMeanCodes []int64 `json:"runMeanCodes,omitzero"`
	// Array of the number(s) of Doppler spectra used to process measurements from
	// radar. Spectral averaging involves combining multiple Doppler spectra acquired
	// to obtain a more accurate and representative spectrum. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	SpecAvgSpectraNums []int64 `json:"specAvgSpectraNums,omitzero"`
	// Array of temporal median filter code(s) of a radar beam. Temporal median
	// filtering is a noise-reducing algorithm which involves replacing each data point
	// with the median value of a window of neighboring points over time. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	TempMedFiltCodes []int64 `json:"tempMedFiltCodes,omitzero"`
	// Array of the total number(s) of records for a radar beam. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	TotRecNums []int64 `json:"totRecNums,omitzero"`
	// Array of vertical distance(s) between points where radar measurements are taken,
	// in meters. If this field is populated, the associated beam(s) must be provided
	// in the beamOrder field.
	VertGateSpacings []float64 `json:"vertGateSpacings,omitzero"`
	// Array of width(s) of each location where radar measurements are taken, in
	// meters. If this field is populated, the associated beam(s) must be provided in
	// the beamOrder field.
	VertGateWidths []float64 `json:"vertGateWidths,omitzero"`
	// Array containing the bandwidth, in megahertz, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformBandwidths []float64 `json:"waveformBandwidths,omitzero"`
	// Array containing the loop gain, in decibels, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field (10 SNR vs. 1 dBsm at 1000 km).
	WaveformLoopGains []float64 `json:"waveformLoopGains,omitzero"`
	// Array containing the maximum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMaxRanges []float64 `json:"waveformMaxRanges,omitzero"`
	// Array containing the minimum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMinRanges []float64 `json:"waveformMinRanges,omitzero"`
	// Array containing the pulse width, in microseconds, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformPulseWidths []float64 `json:"waveformPulseWidths,omitzero"`
	paramObj
}

func (r SensorNewParamsSensorcharacteristic) MarshalJSON() (data []byte, err error) {
	type shadow SensorNewParamsSensorcharacteristic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorNewParamsSensorcharacteristic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorNewParamsSensorcharacteristic](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Sensorlimits define 0 to many limits of a particular sensor in terms of
// observation coverage of on-orbit objects.
//
// The properties ClassificationMarking, DataMode, Source are required.
type SensorNewParamsSensorlimitsCollection struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of the target sensor object.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensorLimits param.Opt[string] `json:"idSensorLimits,omitzero"`
	// Leftmost or minimum lower azimuth within this limit. Interpreted according to
	// site types as lower left azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower left azimuth limits would be equal. (degrees).
	LowerLeftAzimuthLimit param.Opt[float64] `json:"lowerLeftAzimuthLimit,omitzero"`
	// Minimum or lower elevation within this limit. Interpreted according to site
	// types as minimum elevation angle, constant elevation or fan beam centerline.
	// (Degrees).
	LowerLeftElevationLimit param.Opt[float64] `json:"lowerLeftElevationLimit,omitzero"`
	// Rightmost or maximum lower azimuth within this limit. Interpreted according to
	// site types as 2nd lower azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower right azimuth limits would be equal. (degrees).
	LowerRightAzimuthLimit param.Opt[float64] `json:"lowerRightAzimuthLimit,omitzero"`
	// Minimum or lower right elevation within this limit. Interpreted according to
	// site types as minimum right elevation angle, constant elevation or fan beam
	// centerline. If the limit rectangle is parallel to the horizon, the left and
	// right lower elevation limits would be equal. (Degrees).
	LowerRightElevationLimit param.Opt[float64] `json:"lowerRightElevationLimit,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Leftmost or minimum upper azimuth within this sensor limit. Interpreted
	// according to site types as beginning upper azimuth limit, left-hand upper
	// boundary limit. If the limit rectangle is parallel to the horizon, the upper and
	// lower left azimuth limits would be equal. (in degrees).
	UpperLeftAzimuthLimit param.Opt[float64] `json:"upperLeftAzimuthLimit,omitzero"`
	// Maximum or upper elevation within this limit. Interpreted according to site
	// types as maximum elevation angle, half the apex of conical observation pattern
	// or star. (Degrees).
	UpperLeftElevationLimit param.Opt[float64] `json:"upperLeftElevationLimit,omitzero"`
	// Rightmost or maximum upper azimuth within this limit. Interpreted according to
	// site types as 2nd azimuth limit elevation angle of axis of conical observation
	// pattern. If the limit rectangle is parallel to the horizon, the upper and lower
	// right azimuth limits would be equal. (degrees).
	UpperRightAzimuthLimit param.Opt[float64] `json:"upperRightAzimuthLimit,omitzero"`
	// Maximum or upper right elevation within this limit. Interpreted according to
	// site types as maximum rightmost elevation angle, half the apex of conical
	// observation pattern or star. If the limit rectangle is parallel to the horizon,
	// the left and right upper elevation limits would be equal. (Degrees).
	UpperRightElevationLimit param.Opt[float64] `json:"upperRightElevationLimit,omitzero"`
	paramObj
}

func (r SensorNewParamsSensorlimitsCollection) MarshalJSON() (data []byte, err error) {
	type shadow SensorNewParamsSensorlimitsCollection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorNewParamsSensorlimitsCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorNewParamsSensorlimitsCollection](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SensorNewParamsSensorObservationType struct {
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The observation measurement type produced by a sensor.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r SensorNewParamsSensorObservationType) MarshalJSON() (data []byte, err error) {
	type shadow SensorNewParamsSensorObservationType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorNewParamsSensorObservationType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SensorStats contain statistics on sensors related to observation production such
// as last reported observation time.
//
// The properties ClassificationMarking, DataMode, IDSensor, Source are required.
type SensorNewParamsSensorStat struct {
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
	// Unique ID of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Time of last reported observation in ISO 8601 UTC with microsecond precision.
	LastObTime param.Opt[time.Time] `json:"lastObTime,omitzero" format:"date-time"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	paramObj
}

func (r SensorNewParamsSensorStat) MarshalJSON() (data []byte, err error) {
	type shadow SensorNewParamsSensorStat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorNewParamsSensorStat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorNewParamsSensorStat](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SensorNewParamsSensorType struct {
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[int64] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The specific sensor type and/or surveillance capability of this sensor.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r SensorNewParamsSensorType) MarshalJSON() (data []byte, err error) {
	type shadow SensorNewParamsSensorType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorNewParamsSensorType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorUpdateParams struct {
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
	DataMode SensorUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Optional flag indicating if the sensor is active.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Optional US Air Force identifier for the sensor/ASR site, typically for air
	// surveillance radar (ASR) sensors.
	AfID param.Opt[string] `json:"afId,omitzero"`
	// The sensor type at the site. Optional field, intended primarily for ASRs.
	AsrType param.Opt[string] `json:"asrType,omitzero"`
	// Optional dissemination control required for accessing data (e.g observations)
	// produced by this sensor. This is typically a proprietary data owner control for
	// commercial sensors.
	DataControl param.Opt[string] `json:"dataControl,omitzero"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber param.Opt[int64] `json:"sensorNumber,omitzero"`
	// Optional short name for the sensor.
	ShortName param.Opt[string] `json:"shortName,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity SensorUpdateParamsEntity `json:"entity,omitzero"`
	// Collection of Sensorcharacteristics which define characteristics and
	// capabilities of a sensor.
	Sensorcharacteristics []SensorUpdateParamsSensorcharacteristic `json:"sensorcharacteristics,omitzero"`
	// Sensorlimits define 0 to many limits of a particular sensor in terms of
	// observation coverage of on-orbit objects.
	SensorlimitsCollection []SensorUpdateParamsSensorlimitsCollection `json:"sensorlimitsCollection,omitzero"`
	SensorObservationType  SensorUpdateParamsSensorObservationType    `json:"sensorObservationType,omitzero"`
	// Collection of SensorStats which contain statistics of a sensor.
	SensorStats []SensorUpdateParamsSensorStat `json:"sensorStats,omitzero"`
	SensorType  SensorUpdateParamsSensorType   `json:"sensorType,omitzero"`
	paramObj
}

func (r SensorUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SensorUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorUpdateParams) UnmarshalJSON(data []byte) error {
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
type SensorUpdateParamsDataMode string

const (
	SensorUpdateParamsDataModeReal      SensorUpdateParamsDataMode = "REAL"
	SensorUpdateParamsDataModeTest      SensorUpdateParamsDataMode = "TEST"
	SensorUpdateParamsDataModeSimulated SensorUpdateParamsDataMode = "SIMULATED"
	SensorUpdateParamsDataModeExercise  SensorUpdateParamsDataMode = "EXERCISE"
)

// An entity is a generic representation of any object within a space/SSA system
// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
// entity can have an operating unit, a location (if terrestrial), and statuses.
//
// The properties ClassificationMarking, DataMode, Name, Source, Type are required.
type SensorUpdateParamsEntity struct {
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
	// Unique entity name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of entity represented by this record (AIRCRAFT, BUS, COMM, IR,
	// LASEREMITTER, NAVIGATION, ONORBIT, RFEMITTER, SCIENTIFIC, SENSOR, SITE, VESSEL).
	//
	// Any of "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT",
	// "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL".
	Type string `json:"type,omitzero,required"`
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
	// Terrestrial identifier of this entity, if applicable.
	TerrestrialID param.Opt[string] `json:"terrestrialId,omitzero"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LocationIngestParam `json:"location,omitzero"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit SensorUpdateParamsEntityOnOrbit `json:"onOrbit,omitzero"`
	// Type of organization which owns this entity (e.g. Commercial, Government,
	// Academic, Consortium, etc).
	//
	// Any of "Commercial", "Government", "Academic", "Consortium", "Other".
	OwnerType string `json:"ownerType,omitzero"`
	// List of URLs to additional details/documents for this entity.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

func (r SensorUpdateParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow SensorUpdateParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorUpdateParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorUpdateParamsEntity](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SensorUpdateParamsEntity](
		"type", "AIRCRAFT", "BUS", "COMM", "IR", "LASEREMITTER", "NAVIGATION", "ONORBIT", "RFEMITTER", "SCIENTIFIC", "SENSOR", "SITE", "VESSEL",
	)
	apijson.RegisterFieldValidator[SensorUpdateParamsEntity](
		"ownerType", "Commercial", "Government", "Academic", "Consortium", "Other",
	)
}

// Model object representing on-orbit objects or satellites in the system.
//
// The properties ClassificationMarking, DataMode, SatNo, Source are required.
type SensorUpdateParamsEntityOnOrbit struct {
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

func (r SensorUpdateParamsEntityOnOrbit) MarshalJSON() (data []byte, err error) {
	type shadow SensorUpdateParamsEntityOnOrbit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorUpdateParamsEntityOnOrbit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorUpdateParamsEntityOnOrbit](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[SensorUpdateParamsEntityOnOrbit](
		"category", "Unknown", "On-Orbit", "Decayed", "Cataloged Without State", "Launch Nominal", "Analyst Satellite", "Cislunar", "Lunar", "Hyperbolic", "Heliocentric", "Interplanetary", "Lagrangian", "Docked",
	)
	apijson.RegisterFieldValidator[SensorUpdateParamsEntityOnOrbit](
		"objectType", "ROCKET BODY", "DEBRIS", "PAYLOAD", "PLATFORM", "MANNED", "UNKNOWN",
	)
}

// Model representation of characteristics and capabilities of a sensor.
//
// The properties ClassificationMarking, DataMode, IDSensor, Source are required.
type SensorUpdateParamsSensorcharacteristic struct {
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
	// Unique identifier of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Number of bits used in the conversion from analog electrons in a pixel well to a
	// digital number. The digital number has a maximum value of 2^N, where N is the
	// number of bits.
	AnalogToDigitalBitSize param.Opt[int64] `json:"analogToDigitalBitSize,omitzero"`
	// Optical sensor camera aperture.
	Aperture param.Opt[float64] `json:"aperture,omitzero"`
	// For ASR (Air Surveillance Radar) sensors, the scan (360 deg sweep) rate of the
	// radar, in scans/minute.
	AsrScanRate param.Opt[float64] `json:"asrScanRate,omitzero"`
	// One-way radar receiver loss factor due to atmospheric effects. This value will
	// often be the same as the corresponding transmission factor but may be different
	// for bistatic systems.
	AtmosReceiverLoss param.Opt[float64] `json:"atmosReceiverLoss,omitzero"`
	// One-way radar transmission loss factor due to atmospheric effects.
	AtmosTransmissionLoss param.Opt[float64] `json:"atmosTransmissionLoss,omitzero"`
	// Average atmospheric angular width with no distortion from turbulence at an
	// optical sensor site in arcseconds.
	AvgAtmosSeeingConditions param.Opt[float64] `json:"avgAtmosSeeingConditions,omitzero"`
	// Azimuth rate acquisition limit (radians/minute).
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of watts per square meter per steradian
	// (W/(m^2 str)) consistent with sensor detection bands.
	BackgroundSkyRadiance param.Opt[float64] `json:"backgroundSkyRadiance,omitzero"`
	// Average background sky brightness at an optical sensor site during new moon
	// conditions. This field uses units of visual magnitude consistent with sensor
	// detection bands.
	BackgroundSkyVisMag param.Opt[float64] `json:"backgroundSkyVisMag,omitzero"`
	// Sensor band.
	Band param.Opt[string] `json:"band,omitzero"`
	// The total bandwidth, in megahertz, about the radar center frequency.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Number of radar beams used by the sensor.
	BeamQty param.Opt[int64] `json:"beamQty,omitzero"`
	// The angle of the center of a phased array sensor.
	Boresight param.Opt[float64] `json:"boresight,omitzero"`
	// The number of degrees off of the boresight for the sensor.
	BoresightOffAngle param.Opt[float64] `json:"boresightOffAngle,omitzero"`
	// Weighted center wavelength for an optical sensor bandpass in micrometers. It is
	// the center wavelength in a weighted integral sense, accounting for the
	// sensitivity vs. wavelength curve for the sensor focal plane array.
	CenterWavelength param.Opt[float64] `json:"centerWavelength,omitzero"`
	// Collapsing loss in decibels. Collapsing losses occur when two or more sources of
	// noise are added to the target signal. Examples include receiver bandwidth
	// mismatch with filtering bandwidth and elevation or azimuth beam collapse onto
	// position/height indicator displays.
	CollapsingLoss param.Opt[float64] `json:"collapsingLoss,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Threshold shear value beyond which one of the radial velocity values will be
	// rejected, measured in units of inverse second.
	CritShear param.Opt[float64] `json:"critShear,omitzero"`
	// Current flowing through optical sensor focal plane electronics with a closed
	// shutter in electrons per second.
	DarkCurrent param.Opt[float64] `json:"darkCurrent,omitzero"`
	// Description of the equipment and data source.
	Description param.Opt[string] `json:"description,omitzero"`
	// Detection signal-to-noise ratio (SNR) threshold in decibels. This value is
	// typically lower than trackSNR.
	DetectSnr param.Opt[float64] `json:"detectSNR,omitzero"`
	// Sensor duty cycle as a fraction of 1. Duty cycle is the fraction of time a
	// sensor is actively transmitting.
	DutyCycle param.Opt[float64] `json:"dutyCycle,omitzero"`
	// Sensor Earth limb exclusion height in kilometers and is generally only applied
	// to space-based sensors. Some models used an earth exclusion angle instead, but
	// this assumes the sensor is in a circular orbit with constant altitude relative
	// to the earth. The limb exclusion height can be used for space-based sensors in
	// any orbit (assuming it is constant with sensor altitude). The limb height is
	// defined to be 0 at the surface of the earth.
	EarthLimbExclHgt param.Opt[float64] `json:"earthLimbExclHgt,omitzero"`
	// Elevation rate acquisition limit (radians/minute).
	ElevationRateGeolm param.Opt[float64] `json:"elevationRateGeolm,omitzero"`
	// Type of equipment used to take measurements.
	EquipmentType param.Opt[string] `json:"equipmentType,omitzero"`
	// The beam width of a Sensor's Fan (range). The values for this range from (0.0 to
	// PI).
	FanBeamWidth param.Opt[float64] `json:"fanBeamWidth,omitzero"`
	// Number of Fast Fourier Transform (FFT) points used to convert time varying
	// signals into the frequency domain.
	Fft param.Opt[int64] `json:"fft,omitzero"`
	// Maximum number of times the first guess was propagated in each gate before
	// failing the first guess check.
	FgpCrit param.Opt[int64] `json:"fgpCrit,omitzero"`
	// Noise term, in decibels, that arises when a radar receiver filter has a
	// non-optimal bandwidth for an incoming signal (i.e., when it does not match the
	// pulse width).
	FilterMismatchFactor param.Opt[float64] `json:"filterMismatchFactor,omitzero"`
	// F-number for an optical telescope. It is dimensionless and is defined as the
	// ratio of the focal length to the aperture width.
	FNum param.Opt[float64] `json:"fNum,omitzero"`
	// For radar based sensors, the focal point elevation of the radar at the site, in
	// meters.
	FocalPoint param.Opt[float64] `json:"focalPoint,omitzero"`
	// Horizontal field of view, in degrees.
	HFov param.Opt[float64] `json:"hFOV,omitzero"`
	// Horizontal pixel resolution.
	HResPixels param.Opt[int64] `json:"hResPixels,omitzero"`
	// For radar based sensors, K-factor is a relative indicator of refractivity that
	// infers the amount of radar beam bending due to atmosphere. (1<K<2).
	K param.Opt[float64] `json:"k,omitzero"`
	// For Orbiting Sensors, First Card Azimuth limit #1 (left, degrees).
	LeftClockAngle param.Opt[float64] `json:"leftClockAngle,omitzero"`
	// Leftmost GEO belt longitude limit for this sensor (if applicable).
	LeftGeoBeltLimit param.Opt[float64] `json:"leftGeoBeltLimit,omitzero"`
	// Site where measurement is taken.
	Location param.Opt[string] `json:"location,omitzero"`
	// Aggregated radar range equation gain in decibels for maximum sensitivity. It is
	// a roll-up value for low fidelity modeling and is often the only sensitivity
	// value available for a radar system without access to detailed performance
	// parameters.
	LoopGain param.Opt[float64] `json:"loopGain,omitzero"`
	// Lowest aspect angle of the full moon in degrees at which the sensor can achieve
	// full performance.
	LunarExclAngle param.Opt[float64] `json:"lunarExclAngle,omitzero"`
	// Angle between magnetic north and true north at the sensor site, in degrees.
	MagDec param.Opt[float64] `json:"magDec,omitzero"`
	// Absolute magnitude acquisition limit for optical sensors.
	MagnitudeLimit param.Opt[float64] `json:"magnitudeLimit,omitzero"`
	// Max deviation angle of the sensor in degrees.
	MaxDeviationAngle param.Opt[float64] `json:"maxDeviationAngle,omitzero"`
	// Maximum integration time per image frame in seconds for an optical sensor.
	MaxIntegrationTime param.Opt[float64] `json:"maxIntegrationTime,omitzero"`
	// Maximum observable sensor range, in kilometers.
	MaxObservableRange param.Opt[float64] `json:"maxObservableRange,omitzero"`
	// Maximum observable range limit in kilometers -- sensor cannot acquire beyond
	// this range.
	MaxRangeLimit param.Opt[float64] `json:"maxRangeLimit,omitzero"`
	// Maximum wavelength detectable by an optical sensor in micrometers.
	MaxWavelength param.Opt[float64] `json:"maxWavelength,omitzero"`
	// Minimum integration time per image frame in seconds for an optical sensor.
	MinIntegrationTime param.Opt[float64] `json:"minIntegrationTime,omitzero"`
	// Minimum range measurement capability of the sensor, in kilometers.
	MinRangeLimit param.Opt[float64] `json:"minRangeLimit,omitzero"`
	// Signal to Noise Ratio, in decibels. The values for this range from 0.0 - + 99.99
	// dB.
	MinSignalNoiseRatio param.Opt[float64] `json:"minSignalNoiseRatio,omitzero"`
	// Minimum wavelength detectable by an optical sensor in micrometers.
	MinWavelength param.Opt[float64] `json:"minWavelength,omitzero"`
	// Negative Range-rate/relative velocity limit (kilometers/second).
	NegativeRangeRateLimit param.Opt[float64] `json:"negativeRangeRateLimit,omitzero"`
	// Noise figure for a radar system in decibels. This value may be used to compute
	// system noise when the system temperature is unavailable.
	NoiseFigure param.Opt[float64] `json:"noiseFigure,omitzero"`
	// Number of pulses that are non-coherently integrated during detection processing.
	NonCoherentIntegratedPulses param.Opt[int64] `json:"nonCoherentIntegratedPulses,omitzero"`
	// For radar based sensors, number of integrated pulses in a transmit cycle.
	NumIntegratedPulses param.Opt[int64] `json:"numIntegratedPulses,omitzero"`
	// Number of integration frames for an optical sensor.
	NumIntegrationFrames param.Opt[int64] `json:"numIntegrationFrames,omitzero"`
	// The number of optical integration mode characteristics provided in this record.
	// If provided, the numOpticalIntegrationModes value indicates the number of
	// elements in each of the opticalIntegrationTimes, opticalIntegrationAngularRates,
	// opticalIntegrationFrames, opticalIntegrationPixelBinnings, and
	// opticalIntegrationSNRs arrays.
	NumOpticalIntegrationModes param.Opt[int64] `json:"numOpticalIntegrationModes,omitzero"`
	// The number of waveforms characteristics provided in this record. If provided,
	// the numWaveforms value indicates the number of elements in each of the
	// waveformPulseWidths, waveformBandWidths, waveformMinRanges, waveformMaxRanges,
	// and waveformLoopGains arrays.
	NumWaveforms param.Opt[int64] `json:"numWaveforms,omitzero"`
	// Fraction of incident light transmitted to an optical sensor focal plane array.
	// The value is given as a fraction of 1, not as a percent.
	OpticalTransmission param.Opt[float64] `json:"opticalTransmission,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Two-way pattern absorption/propagation loss due to medium in decibels.
	PatternAbsorptionLoss param.Opt[float64] `json:"patternAbsorptionLoss,omitzero"`
	// Losses from the beam shape, scanning, and pattern factor in decibels. These
	// losses occur when targets are not directly in line with a beam center. For space
	// surveillance, this will occur most often during sensor scanning.
	PatternScanLoss param.Opt[float64] `json:"patternScanLoss,omitzero"`
	// Maximum instantaneous radar transmit power in watts for use in the radar range
	// equation.
	PeakPower param.Opt[float64] `json:"peakPower,omitzero"`
	// Angular field-of-view covered by one pixel in a focal plane array in
	// microradians. The pixel is assumed to be a perfect square so that only a single
	// value is required.
	PixelInstantaneousFov param.Opt[float64] `json:"pixelInstantaneousFOV,omitzero"`
	// Maximum number of electrons that can be collected in a single pixel on an
	// optical sensor focal plane array.
	PixelWellDepth param.Opt[int64] `json:"pixelWellDepth,omitzero"`
	// Positive Range-rate/relative velocity limit (kilometers/second).
	PositiveRangeRateLimit param.Opt[float64] `json:"positiveRangeRateLimit,omitzero"`
	// For radar based sensors, pulse repetition frequency (PRF), in hertz. Number of
	// new pulses transmitted per second.
	Prf param.Opt[float64] `json:"prf,omitzero"`
	// Designated probability of detection at the signal-to-noise detection threshold.
	ProbDetectSnr param.Opt[float64] `json:"probDetectSNR,omitzero"`
	// For radar based sensors, probability of the indication of the presence of a
	// radar target due to noise or interference.
	ProbFalseAlarm param.Opt[float64] `json:"probFalseAlarm,omitzero"`
	// Fraction of incident photons converted to electrons at the focal plane array.
	// This value is a decimal number between 0 and 1, inclusive.
	QuantumEff param.Opt[float64] `json:"quantumEff,omitzero"`
	// Radar frequency in hertz, of the sensor (if a radar sensor).
	RadarFrequency param.Opt[float64] `json:"radarFrequency,omitzero"`
	// Message data format transmitted by the sensor digitizer.
	RadarMessageFormat param.Opt[string] `json:"radarMessageFormat,omitzero"`
	// For radar based sensors, radar maximum unambiguous range, in kilometers.
	RadarMur param.Opt[float64] `json:"radarMUR,omitzero"`
	// Radio frequency (if sensor is RF).
	RadioFrequency param.Opt[float64] `json:"radioFrequency,omitzero"`
	// Losses due to the presence of a protective radome surrounding a radar sensor, in
	// decibels.
	RadomeLoss param.Opt[float64] `json:"radomeLoss,omitzero"`
	// Number of false-signal electrons generated by optical sensor focal plane
	// read-out electronics from photon-to-electron conversion during frame
	// integration. The units are in electrons RMS.
	ReadNoise param.Opt[int64] `json:"readNoise,omitzero"`
	// Radar receive gain in decibels.
	ReceiveGain param.Opt[float64] `json:"receiveGain,omitzero"`
	// Horizontal/azimuthal receive beamwidth for a radar in degrees.
	ReceiveHorizBeamWidth param.Opt[float64] `json:"receiveHorizBeamWidth,omitzero"`
	// Aggregate radar receive loss, in decibels.
	ReceiveLoss param.Opt[float64] `json:"receiveLoss,omitzero"`
	// Vertical/elevation receive beamwidth for a radar in degrees.
	ReceiveVertBeamWidth param.Opt[float64] `json:"receiveVertBeamWidth,omitzero"`
	// Reference temperature for radar noise in Kelvin. A reference temperature is used
	// when the radar system temperature is unknown and is combined with the system
	// noise figure to estimate signal loss.
	RefTemp param.Opt[float64] `json:"refTemp,omitzero"`
	// For Orbiting Sensors, First Card Azimuth limit #3 (right, degrees).
	RightClockAngle param.Opt[float64] `json:"rightClockAngle,omitzero"`
	// Rightmost GEO belt longitude limit for this sensor (if applicable).
	RightGeoBeltLimit param.Opt[float64] `json:"rightGeoBeltLimit,omitzero"`
	// Radar signal processing losses, in decibels.
	SignalProcessingLoss param.Opt[float64] `json:"signalProcessingLoss,omitzero"`
	// Site code of the sensor.
	SiteCode param.Opt[string] `json:"siteCode,omitzero"`
	// Sensor and target position vector origins are at the center of the earth. The
	// sun vector origin is at the target position and points toward the sun. Any value
	// between 0 and 180 degrees is acceptable and is assumed to apply in both
	// directions (i.e., a solar exclusion angle of 30 degrees is understood to mean no
	// viewing for any angle between -30 deg and +30 deg).
	SolarExclAngle param.Opt[float64] `json:"solarExclAngle,omitzero"`
	// For radar based sensors, expression of the radar system noise, aggregated as an
	// equivalent thermal noise value, in degrees Kelvin.
	SystemNoiseTemperature param.Opt[float64] `json:"systemNoiseTemperature,omitzero"`
	// Maximum taskable range of the sensor, in kilometers.
	TaskableRange param.Opt[float64] `json:"taskableRange,omitzero"`
	// Test number for the observed measurement.
	TestNumber param.Opt[string] `json:"testNumber,omitzero"`
	// For tower sensors, the physical height of the sensor tower, in meters.
	TowerHeight param.Opt[float64] `json:"towerHeight,omitzero"`
	// Beginning track angle limit, in radians. Track angle is the angle between the
	// camera axis and the gimbal plane. Values range from 0 - PI/2.
	TrackAngle param.Opt[float64] `json:"trackAngle,omitzero"`
	// Track signal-to-noise ratio (SNR) threshold in decibels. This value is typically
	// higher than detectSNR.
	TrackSnr param.Opt[float64] `json:"trackSNR,omitzero"`
	// Radar transmit gain in decibels.
	TransmitGain param.Opt[float64] `json:"transmitGain,omitzero"`
	// Horizontal/azimuthal transmit beamwidth for a radar in degrees.
	TransmitHorizBeamWidth param.Opt[float64] `json:"transmitHorizBeamWidth,omitzero"`
	// Aggregate radar transmit loss, in decibels.
	TransmitLoss param.Opt[float64] `json:"transmitLoss,omitzero"`
	// Radar transmit power in Watts.
	TransmitPower param.Opt[float64] `json:"transmitPower,omitzero"`
	// Vertical/elevation transmit beamwidth for a radar in degrees.
	TransmitVertBeamWidth param.Opt[float64] `json:"transmitVertBeamWidth,omitzero"`
	// True North correction for the sensor, in ACP (Azimunth Change Pulse) count.
	TrueNorthCorrector param.Opt[int64] `json:"trueNorthCorrector,omitzero"`
	// Antenna true tilt, in degrees.
	TrueTilt param.Opt[float64] `json:"trueTilt,omitzero"`
	// Twilight angle for ground-based optical sensors in degrees. A sensor cannot view
	// targets until the sun is below the twilight angle relative to the local horizon.
	// The sign of the angle is positive despite the sun elevation being negative after
	// local sunset. Typical values for the twilight angle are civil twilight (6
	// degrees), nautical twilight (12 degrees), and astronomical twilight (18
	// degrees).
	TwilightAngle param.Opt[float64] `json:"twilightAngle,omitzero"`
	// Flag indicating if a vertical radar beam was used in the wind calculation.
	VertBeamFlag param.Opt[bool] `json:"vertBeamFlag,omitzero"`
	// Vertical field of view, in degrees.
	VFov param.Opt[float64] `json:"vFOV,omitzero"`
	// Vertical pixel resolution.
	VResPixels param.Opt[int64] `json:"vResPixels,omitzero"`
	// Peformance zone-1 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MaxRange param.Opt[float64] `json:"z1MaxRange,omitzero"`
	// Peformance zone-1 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z1MinRange param.Opt[float64] `json:"z1MinRange,omitzero"`
	// Peformance zone-2 maximum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MaxRange param.Opt[float64] `json:"z2MaxRange,omitzero"`
	// Peformance zone-2 minimum range, in kilometers. Note that the zones apply only
	// to the PSR/Search radars.
	Z2MinRange param.Opt[float64] `json:"z2MinRange,omitzero"`
	// Array of measurement range(s) where radar samples must fall to be acceptable. If
	// this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	AcceptSampleRanges []float64 `json:"acceptSampleRanges,omitzero"`
	// Array of azimuth angles of a radar beam, in degrees. If this field is populated,
	// the associated beam(s) must be provided in the beamOrder field.
	AzAngs []float64 `json:"azAngs,omitzero"`
	// Array designating the beam order of provided values (e.g. vb1 for vertical beam
	// 1, ob1 for oblique beam 1, etc.). Required if any of the following array fields
	// are populated: azAngs, elAngs, radarPulseWidths, pulseRepPeriods, delayGates,
	// rangeGates, rangeSpacings, vertGateSpacings, vertGateWidths, specAvgSpectraNums,
	// tempMedFiltCodes, runMeanCodes, totRecNums, reqRecords, acceptSampleRanges.
	BeamOrder []string `json:"beamOrder,omitzero"`
	// Array of time delay(s) for pulses from a radar beam to get to the first range
	// gate, in nanoseconds. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	DelayGates []float64 `json:"delayGates,omitzero"`
	// Array of elevation angles of a radar beam, in degrees. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	ElAngs []float64 `json:"elAngs,omitzero"`
	// Array containing the angular rate, in arcsec/sec, for each provided optical
	// integration mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationAngularRates []float64 `json:"opticalIntegrationAngularRates,omitzero"`
	// Array containing the number of frames, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationFrames []float64 `json:"opticalIntegrationFrames,omitzero"`
	// Array containing the pixel binning, for each optical integration mode. The
	// number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationPixelBinnings []float64 `json:"opticalIntegrationPixelBinnings,omitzero"`
	// Array of optical integration modes signal to noise ratios. The number of
	// elements must be equal to the value indicated in numOpticalIntegrationModes.
	OpticalIntegrationSnRs []float64 `json:"opticalIntegrationSNRs,omitzero"`
	// Array containing the time, in seconds, for each provided optical integration
	// mode. The number of elements must be equal to the value indicated in
	// numOpticalIntegrationModes.
	OpticalIntegrationTimes []float64 `json:"opticalIntegrationTimes,omitzero"`
	// Array of interval(s) between the start of one radar pulse and the start of
	// another for a radar beam, in microseconds. If this field is populated, the
	// associated beam(s) must be provided in the beamOrder field.
	PulseRepPeriods []float64 `json:"pulseRepPeriods,omitzero"`
	// Array of transmit time(s) for a radar beam pulse, in microseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RadarPulseWidths []float64 `json:"radarPulseWidths,omitzero"`
	// Array of the number(s) of discrete altitudes where return signals are sampled by
	// a radar beam. If this field is populated, the associated beam(s) must be
	// provided in the beamOrder field.
	RangeGates []int64 `json:"rangeGates,omitzero"`
	// Array of range gate spacing(s) for a radar beam, in nanoseconds. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	RangeSpacings []float64 `json:"rangeSpacings,omitzero"`
	// Array of the total number(s) of records required to meet consensus for a radar
	// beam. If this field is populated, the associated beam(s) must be provided in the
	// beamOrder field.
	ReqRecords []int64 `json:"reqRecords,omitzero"`
	// Array of running mean code(s) used by radar data processing. The running mean
	// method involves taking a series of averages of different selections of the full
	// data set to smooth out short-term fluctuations in the data. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	RunMeanCodes []int64 `json:"runMeanCodes,omitzero"`
	// Array of the number(s) of Doppler spectra used to process measurements from
	// radar. Spectral averaging involves combining multiple Doppler spectra acquired
	// to obtain a more accurate and representative spectrum. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	SpecAvgSpectraNums []int64 `json:"specAvgSpectraNums,omitzero"`
	// Array of temporal median filter code(s) of a radar beam. Temporal median
	// filtering is a noise-reducing algorithm which involves replacing each data point
	// with the median value of a window of neighboring points over time. If this field
	// is populated, the associated beam(s) must be provided in the beamOrder field.
	TempMedFiltCodes []int64 `json:"tempMedFiltCodes,omitzero"`
	// Array of the total number(s) of records for a radar beam. If this field is
	// populated, the associated beam(s) must be provided in the beamOrder field.
	TotRecNums []int64 `json:"totRecNums,omitzero"`
	// Array of vertical distance(s) between points where radar measurements are taken,
	// in meters. If this field is populated, the associated beam(s) must be provided
	// in the beamOrder field.
	VertGateSpacings []float64 `json:"vertGateSpacings,omitzero"`
	// Array of width(s) of each location where radar measurements are taken, in
	// meters. If this field is populated, the associated beam(s) must be provided in
	// the beamOrder field.
	VertGateWidths []float64 `json:"vertGateWidths,omitzero"`
	// Array containing the bandwidth, in megahertz, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformBandwidths []float64 `json:"waveformBandwidths,omitzero"`
	// Array containing the loop gain, in decibels, for each provided waveform. The
	// number of elements in this array must be equal to the value indicated in the
	// numWaveforms field (10 SNR vs. 1 dBsm at 1000 km).
	WaveformLoopGains []float64 `json:"waveformLoopGains,omitzero"`
	// Array containing the maximum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMaxRanges []float64 `json:"waveformMaxRanges,omitzero"`
	// Array containing the minimum range, in kilometers, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformMinRanges []float64 `json:"waveformMinRanges,omitzero"`
	// Array containing the pulse width, in microseconds, for each provided waveform.
	// The number of elements in this array must be equal to the value indicated in the
	// numWaveforms field.
	WaveformPulseWidths []float64 `json:"waveformPulseWidths,omitzero"`
	paramObj
}

func (r SensorUpdateParamsSensorcharacteristic) MarshalJSON() (data []byte, err error) {
	type shadow SensorUpdateParamsSensorcharacteristic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorUpdateParamsSensorcharacteristic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorUpdateParamsSensorcharacteristic](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Sensorlimits define 0 to many limits of a particular sensor in terms of
// observation coverage of on-orbit objects.
//
// The properties ClassificationMarking, DataMode, Source are required.
type SensorUpdateParamsSensorlimitsCollection struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of the target sensor object.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDSensorLimits param.Opt[string] `json:"idSensorLimits,omitzero"`
	// Leftmost or minimum lower azimuth within this limit. Interpreted according to
	// site types as lower left azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower left azimuth limits would be equal. (degrees).
	LowerLeftAzimuthLimit param.Opt[float64] `json:"lowerLeftAzimuthLimit,omitzero"`
	// Minimum or lower elevation within this limit. Interpreted according to site
	// types as minimum elevation angle, constant elevation or fan beam centerline.
	// (Degrees).
	LowerLeftElevationLimit param.Opt[float64] `json:"lowerLeftElevationLimit,omitzero"`
	// Rightmost or maximum lower azimuth within this limit. Interpreted according to
	// site types as 2nd lower azimuth limit elevation angle of axis of conical
	// observation pattern. If the limit rectangle is parallel to the horizon, the
	// upper and lower right azimuth limits would be equal. (degrees).
	LowerRightAzimuthLimit param.Opt[float64] `json:"lowerRightAzimuthLimit,omitzero"`
	// Minimum or lower right elevation within this limit. Interpreted according to
	// site types as minimum right elevation angle, constant elevation or fan beam
	// centerline. If the limit rectangle is parallel to the horizon, the left and
	// right lower elevation limits would be equal. (Degrees).
	LowerRightElevationLimit param.Opt[float64] `json:"lowerRightElevationLimit,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Leftmost or minimum upper azimuth within this sensor limit. Interpreted
	// according to site types as beginning upper azimuth limit, left-hand upper
	// boundary limit. If the limit rectangle is parallel to the horizon, the upper and
	// lower left azimuth limits would be equal. (in degrees).
	UpperLeftAzimuthLimit param.Opt[float64] `json:"upperLeftAzimuthLimit,omitzero"`
	// Maximum or upper elevation within this limit. Interpreted according to site
	// types as maximum elevation angle, half the apex of conical observation pattern
	// or star. (Degrees).
	UpperLeftElevationLimit param.Opt[float64] `json:"upperLeftElevationLimit,omitzero"`
	// Rightmost or maximum upper azimuth within this limit. Interpreted according to
	// site types as 2nd azimuth limit elevation angle of axis of conical observation
	// pattern. If the limit rectangle is parallel to the horizon, the upper and lower
	// right azimuth limits would be equal. (degrees).
	UpperRightAzimuthLimit param.Opt[float64] `json:"upperRightAzimuthLimit,omitzero"`
	// Maximum or upper right elevation within this limit. Interpreted according to
	// site types as maximum rightmost elevation angle, half the apex of conical
	// observation pattern or star. If the limit rectangle is parallel to the horizon,
	// the left and right upper elevation limits would be equal. (Degrees).
	UpperRightElevationLimit param.Opt[float64] `json:"upperRightElevationLimit,omitzero"`
	paramObj
}

func (r SensorUpdateParamsSensorlimitsCollection) MarshalJSON() (data []byte, err error) {
	type shadow SensorUpdateParamsSensorlimitsCollection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorUpdateParamsSensorlimitsCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorUpdateParamsSensorlimitsCollection](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SensorUpdateParamsSensorObservationType struct {
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The observation measurement type produced by a sensor.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r SensorUpdateParamsSensorObservationType) MarshalJSON() (data []byte, err error) {
	type shadow SensorUpdateParamsSensorObservationType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorUpdateParamsSensorObservationType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SensorStats contain statistics on sensors related to observation production such
// as last reported observation time.
//
// The properties ClassificationMarking, DataMode, IDSensor, Source are required.
type SensorUpdateParamsSensorStat struct {
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
	// Unique ID of the parent sensor.
	IDSensor string `json:"idSensor,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Time of last reported observation in ISO 8601 UTC with microsecond precision.
	LastObTime param.Opt[time.Time] `json:"lastObTime,omitzero" format:"date-time"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	paramObj
}

func (r SensorUpdateParamsSensorStat) MarshalJSON() (data []byte, err error) {
	type shadow SensorUpdateParamsSensorStat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorUpdateParamsSensorStat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorUpdateParamsSensorStat](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SensorUpdateParamsSensorType struct {
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[int64] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The specific sensor type and/or surveillance capability of this sensor.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r SensorUpdateParamsSensorType) MarshalJSON() (data []byte, err error) {
	type shadow SensorUpdateParamsSensorType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorUpdateParamsSensorType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorListParams]'s query parameters as `url.Values`.
func (r SensorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorCountParams]'s query parameters as `url.Values`.
func (r SensorCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorGetParams]'s query parameters as `url.Values`.
func (r SensorGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorTupleParams]'s query parameters as `url.Values`.
func (r SensorTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
