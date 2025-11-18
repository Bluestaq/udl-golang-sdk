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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/sensor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Sensor. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *SensorService) Update(ctx context.Context, id string, body SensorUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sensor/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Sensor by its unique ID passed as a path
// parameter.
func (r *SensorService) Get(ctx context.Context, id string, query SensorGetParams, opts ...option.RequestOption) (res *SensorGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	Location shared.LocationAbridged `json:"location"`
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
	OperatingUnit shared.OperatingunitFull `json:"operatingUnit"`
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
	RfBands []shared.RfBandFull `json:"rfBands"`
	// Read-only collection of statuses which can be collected by multiple sources.
	StatusCollection []shared.StatusFull `json:"statusCollection"`
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
	Antennas []shared.OnorbitAntennaFull `json:"antennas"`
	// Read-only collection of batteries on this on-orbit object.
	Batteries []shared.OnorbitBatteryFull `json:"batteries"`
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
	OnorbitDetails []shared.OnorbitDetailsFull `json:"onorbitDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only collection of solar arrays on this on-orbit object.
	SolarArrays []shared.OnorbitSolarArrayFull `json:"solarArrays"`
	// Read-only collection of thrusters (engines) on this on-orbit object.
	Thrusters []shared.OnorbitThrusterFull `json:"thrusters"`
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
func (r SensorQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorQueryhelpResponse) UnmarshalJSON(data []byte) error {
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
	OperatingUnit shared.OperatingunitFull `json:"operatingUnit"`
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
	RfBands []shared.RfBandFull `json:"rfBands"`
	// Read-only collection of statuses which can be collected by multiple sources.
	StatusCollection []shared.StatusFull `json:"statusCollection"`
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
	Antennas []shared.OnorbitAntennaFull `json:"antennas"`
	// Read-only collection of batteries on this on-orbit object.
	Batteries []shared.OnorbitBatteryFull `json:"batteries"`
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
	OnorbitDetails []shared.OnorbitDetailsFull `json:"onorbitDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only collection of solar arrays on this on-orbit object.
	SolarArrays []shared.OnorbitSolarArrayFull `json:"solarArrays"`
	// Read-only collection of thrusters (engines) on this on-orbit object.
	Thrusters []shared.OnorbitThrusterFull `json:"thrusters"`
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
	// Time of last reported observation in ISO 8601 UTC with microsecond precision.
	LastObTime param.Opt[time.Time] `json:"lastObTime,omitzero" format:"date-time"`
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
	// Time of last reported observation in ISO 8601 UTC with microsecond precision.
	LastObTime param.Opt[time.Time] `json:"lastObTime,omitzero" format:"date-time"`
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
