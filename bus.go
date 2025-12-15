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

// BusService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBusService] method instead.
type BusService struct {
	Options []option.RequestOption
}

// NewBusService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBusService(opts ...option.RequestOption) (r BusService) {
	r = BusService{}
	r.Options = opts
	return
}

// Service operation to take a single Bus as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *BusService) New(ctx context.Context, body BusNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/bus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Bus record by its unique ID passed as a path
// parameter.
func (r *BusService) Get(ctx context.Context, id string, query BusGetParams, opts ...option.RequestOption) (res *shared.BusFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/bus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single Bus. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *BusService) Update(ctx context.Context, id string, body BusUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/bus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *BusService) List(ctx context.Context, query BusListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[BusAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/bus"
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
func (r *BusService) ListAutoPaging(ctx context.Context, query BusListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[BusAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Bus object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *BusService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/bus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *BusService) Count(ctx context.Context, query BusCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/bus/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *BusService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *BusQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/bus/queryhelp"
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
func (r *BusService) Tuple(ctx context.Context, query BusTupleParams, opts ...option.RequestOption) (res *[]shared.BusFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/bus/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A bus is the physical and software infrastructure backbone to which on-orbit
// satellite payloads are attached for power, control, and other support functions.
type BusAbridged struct {
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
	DataMode BusAbridgedDataMode `json:"dataMode,required"`
	// Name of this bus.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Attitude and Orbital Control Notes/description for the bus.
	AocsNotes string `json:"aocsNotes"`
	// Average mass of this bus without payloads or fuel, in kilograms.
	AvgDryMass float64 `json:"avgDryMass"`
	// Average mass available on this bus for payloads, in kilograms.
	AvgPayloadMass float64 `json:"avgPayloadMass"`
	// Average power available on this bus for payloads, in kilowatts.
	AvgPayloadPower float64 `json:"avgPayloadPower"`
	// Average power available on this bus, in kilowatts.
	AvgSpacecraftPower float64 `json:"avgSpacecraftPower"`
	// Average mass of this bus with fuel, but without payloads, in kilograms.
	AvgWetMass float64 `json:"avgWetMass"`
	// Body dimension in X direction pertaining to length, in meters.
	BodyDimensionX float64 `json:"bodyDimensionX"`
	// Body dimension in Y direction pertaining to height, in meters.
	BodyDimensionY float64 `json:"bodyDimensionY"`
	// Body dimension in Z direction pertaining to width, in meters.
	BodyDimensionZ float64 `json:"bodyDimensionZ"`
	// Unique identifier of the organization which designs the bus kit.
	BusKitDesignerOrgID string `json:"busKitDesignerOrgId"`
	// Country where this bus was manufactured. This value is typically the ISO 3166
	// Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Notes/description of the bus.
	Description string `json:"description"`
	// Boolean indicating if this bus is generic.
	Generic bool `json:"generic"`
	// ID of the parent entity for this bus.
	IDEntity string `json:"idEntity"`
	// Launch envelope dimension in X direction, in meters.
	LaunchEnvelopeDimensionX float64 `json:"launchEnvelopeDimensionX"`
	// Launch envelope dimension in Y direction, in meters.
	LaunchEnvelopeDimensionY float64 `json:"launchEnvelopeDimensionY"`
	// Launch envelope dimension in Z direction, in meters.
	LaunchEnvelopeDimensionZ float64 `json:"launchEnvelopeDimensionZ"`
	// Unique identifier of the organization which manufactures the main onboard
	// computer for this bus.
	MainComputerManufacturerOrgID string `json:"mainComputerManufacturerOrgId"`
	// Unique identifier of the organization which manufactures this bus.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Mass category of this bus (e.g. 1 - 10 kg: Nanosatellite, 10 - 100 kg:
	// Microsatellite, 100 - 500 kg: Minisatellite, 1000 - 2500kg: Medium satellite,
	// etc.).
	MassCategory string `json:"massCategory"`
	// Maximum power at beginning of life, lower bounds, in kilowatts.
	MaxBolPowerLower float64 `json:"maxBOLPowerLower"`
	// Maximum power at beginning of life, upper bounds, in kilowatts.
	MaxBolPowerUpper float64 `json:"maxBOLPowerUpper"`
	// Maximum mass on station at beginning of life, in kilograms.
	MaxBolStationMass float64 `json:"maxBOLStationMass"`
	// Maximum mass of this bus without payloads or fuel, in kilograms.
	MaxDryMass float64 `json:"maxDryMass"`
	// Maximum power at end of life, lower bounds, in kilowatts.
	MaxEolPowerLower float64 `json:"maxEOLPowerLower"`
	// Maximum power at end of life, upper bounds, in kilowatts.
	MaxEolPowerUpper float64 `json:"maxEOLPowerUpper"`
	// Maximum mass at launch, lower bounds, in kilograms.
	MaxLaunchMassLower float64 `json:"maxLaunchMassLower"`
	// Maximum mass at launch, upper bounds, in kilograms.
	MaxLaunchMassUpper float64 `json:"maxLaunchMassUpper"`
	// Maximum payload mass available, in kilograms.
	MaxPayloadMass float64 `json:"maxPayloadMass"`
	// Maximum payload power available, in kilowatts.
	MaxPayloadPower float64 `json:"maxPayloadPower"`
	// Maximum power available on this bus, in kilowatts.
	MaxSpacecraftPower float64 `json:"maxSpacecraftPower"`
	// Maximum mass of this bus with fuel, but without payloads, in kilograms.
	MaxWetMass float64 `json:"maxWetMass"`
	// Median mass of this bus without payloads or fuel, in kilograms.
	MedianDryMass float64 `json:"medianDryMass"`
	// Median mass of this bus with fuel, but without payloads, in kilograms.
	MedianWetMass float64 `json:"medianWetMass"`
	// Minimum mass of this bus without payloads or fuel, in kilograms.
	MinDryMass float64 `json:"minDryMass"`
	// Minimum mass of this bus with fuel, but without payloads, in kilograms.
	MinWetMass float64 `json:"minWetMass"`
	// The number of orbit types this bus can support.
	NumOrbitType int64 `json:"numOrbitType"`
	// Orbit averaged power (the power averaged over one orbit) available on this bus
	// for payloads, in kilowatts.
	OapPayloadPower float64 `json:"oapPayloadPower"`
	// Orbit averaged power (the power averaged over one orbit) available on this bus,
	// in kilowatts.
	OapSpacecraftPower float64 `json:"oapSpacecraftPower"`
	// Array of orbit types this bus can support (e.g. GEO, LEO, etc.). Must contain
	// the same number of elements as the value of numOrbitType.
	OrbitTypes []string `json:"orbitTypes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The radial dimension available on this bus for payloads, in meters.
	PayloadDimensionX float64 `json:"payloadDimensionX"`
	// The in-track dimension available on this bus for payloads, in meters.
	PayloadDimensionY float64 `json:"payloadDimensionY"`
	// The cross-track dimension available on this bus for payloads, in meters.
	PayloadDimensionZ float64 `json:"payloadDimensionZ"`
	// The volume available on this bus for payloads, in cubic meters.
	PayloadVolume float64 `json:"payloadVolume"`
	// Power category of this bus (e.g. 0-1kW low power, etc).
	PowerCategory string `json:"powerCategory"`
	// Unique identifier of the organization which manufactures the telemetry tracking
	// and command subsystem for this bus.
	TelemetryTrackingManufacturerOrgID string `json:"telemetryTrackingManufacturerOrgId"`
	// Type of this bus.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking              respjson.Field
		DataMode                           respjson.Field
		Name                               respjson.Field
		Source                             respjson.Field
		ID                                 respjson.Field
		AocsNotes                          respjson.Field
		AvgDryMass                         respjson.Field
		AvgPayloadMass                     respjson.Field
		AvgPayloadPower                    respjson.Field
		AvgSpacecraftPower                 respjson.Field
		AvgWetMass                         respjson.Field
		BodyDimensionX                     respjson.Field
		BodyDimensionY                     respjson.Field
		BodyDimensionZ                     respjson.Field
		BusKitDesignerOrgID                respjson.Field
		CountryCode                        respjson.Field
		CreatedAt                          respjson.Field
		CreatedBy                          respjson.Field
		Description                        respjson.Field
		Generic                            respjson.Field
		IDEntity                           respjson.Field
		LaunchEnvelopeDimensionX           respjson.Field
		LaunchEnvelopeDimensionY           respjson.Field
		LaunchEnvelopeDimensionZ           respjson.Field
		MainComputerManufacturerOrgID      respjson.Field
		ManufacturerOrgID                  respjson.Field
		MassCategory                       respjson.Field
		MaxBolPowerLower                   respjson.Field
		MaxBolPowerUpper                   respjson.Field
		MaxBolStationMass                  respjson.Field
		MaxDryMass                         respjson.Field
		MaxEolPowerLower                   respjson.Field
		MaxEolPowerUpper                   respjson.Field
		MaxLaunchMassLower                 respjson.Field
		MaxLaunchMassUpper                 respjson.Field
		MaxPayloadMass                     respjson.Field
		MaxPayloadPower                    respjson.Field
		MaxSpacecraftPower                 respjson.Field
		MaxWetMass                         respjson.Field
		MedianDryMass                      respjson.Field
		MedianWetMass                      respjson.Field
		MinDryMass                         respjson.Field
		MinWetMass                         respjson.Field
		NumOrbitType                       respjson.Field
		OapPayloadPower                    respjson.Field
		OapSpacecraftPower                 respjson.Field
		OrbitTypes                         respjson.Field
		Origin                             respjson.Field
		OrigNetwork                        respjson.Field
		PayloadDimensionX                  respjson.Field
		PayloadDimensionY                  respjson.Field
		PayloadDimensionZ                  respjson.Field
		PayloadVolume                      respjson.Field
		PowerCategory                      respjson.Field
		TelemetryTrackingManufacturerOrgID respjson.Field
		Type                               respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BusAbridged) RawJSON() string { return r.JSON.raw }
func (r *BusAbridged) UnmarshalJSON(data []byte) error {
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
type BusAbridgedDataMode string

const (
	BusAbridgedDataModeReal      BusAbridgedDataMode = "REAL"
	BusAbridgedDataModeTest      BusAbridgedDataMode = "TEST"
	BusAbridgedDataModeSimulated BusAbridgedDataMode = "SIMULATED"
	BusAbridgedDataModeExercise  BusAbridgedDataMode = "EXERCISE"
)

type BusQueryHelpResponse struct {
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
func (r BusQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *BusQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BusNewParams struct {
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
	DataMode BusNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Name of this bus.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Attitude and Orbital Control Notes/description for the bus.
	AocsNotes param.Opt[string] `json:"aocsNotes,omitzero"`
	// Average mass of this bus without payloads or fuel, in kilograms.
	AvgDryMass param.Opt[float64] `json:"avgDryMass,omitzero"`
	// Average mass available on this bus for payloads, in kilograms.
	AvgPayloadMass param.Opt[float64] `json:"avgPayloadMass,omitzero"`
	// Average power available on this bus for payloads, in kilowatts.
	AvgPayloadPower param.Opt[float64] `json:"avgPayloadPower,omitzero"`
	// Average power available on this bus, in kilowatts.
	AvgSpacecraftPower param.Opt[float64] `json:"avgSpacecraftPower,omitzero"`
	// Average mass of this bus with fuel, but without payloads, in kilograms.
	AvgWetMass param.Opt[float64] `json:"avgWetMass,omitzero"`
	// Body dimension in X direction pertaining to length, in meters.
	BodyDimensionX param.Opt[float64] `json:"bodyDimensionX,omitzero"`
	// Body dimension in Y direction pertaining to height, in meters.
	BodyDimensionY param.Opt[float64] `json:"bodyDimensionY,omitzero"`
	// Body dimension in Z direction pertaining to width, in meters.
	BodyDimensionZ param.Opt[float64] `json:"bodyDimensionZ,omitzero"`
	// Unique identifier of the organization which designs the bus kit.
	BusKitDesignerOrgID param.Opt[string] `json:"busKitDesignerOrgId,omitzero"`
	// Country where this bus was manufactured. This value is typically the ISO 3166
	// Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Notes/description of the bus.
	Description param.Opt[string] `json:"description,omitzero"`
	// Boolean indicating if this bus is generic.
	Generic param.Opt[bool] `json:"generic,omitzero"`
	// ID of the parent entity for this bus.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Launch envelope dimension in X direction, in meters.
	LaunchEnvelopeDimensionX param.Opt[float64] `json:"launchEnvelopeDimensionX,omitzero"`
	// Launch envelope dimension in Y direction, in meters.
	LaunchEnvelopeDimensionY param.Opt[float64] `json:"launchEnvelopeDimensionY,omitzero"`
	// Launch envelope dimension in Z direction, in meters.
	LaunchEnvelopeDimensionZ param.Opt[float64] `json:"launchEnvelopeDimensionZ,omitzero"`
	// Unique identifier of the organization which manufactures the main onboard
	// computer for this bus.
	MainComputerManufacturerOrgID param.Opt[string] `json:"mainComputerManufacturerOrgId,omitzero"`
	// Unique identifier of the organization which manufactures this bus.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Mass category of this bus (e.g. 1 - 10 kg: Nanosatellite, 10 - 100 kg:
	// Microsatellite, 100 - 500 kg: Minisatellite, 1000 - 2500kg: Medium satellite,
	// etc.).
	MassCategory param.Opt[string] `json:"massCategory,omitzero"`
	// Maximum power at beginning of life, lower bounds, in kilowatts.
	MaxBolPowerLower param.Opt[float64] `json:"maxBOLPowerLower,omitzero"`
	// Maximum power at beginning of life, upper bounds, in kilowatts.
	MaxBolPowerUpper param.Opt[float64] `json:"maxBOLPowerUpper,omitzero"`
	// Maximum mass on station at beginning of life, in kilograms.
	MaxBolStationMass param.Opt[float64] `json:"maxBOLStationMass,omitzero"`
	// Maximum mass of this bus without payloads or fuel, in kilograms.
	MaxDryMass param.Opt[float64] `json:"maxDryMass,omitzero"`
	// Maximum power at end of life, lower bounds, in kilowatts.
	MaxEolPowerLower param.Opt[float64] `json:"maxEOLPowerLower,omitzero"`
	// Maximum power at end of life, upper bounds, in kilowatts.
	MaxEolPowerUpper param.Opt[float64] `json:"maxEOLPowerUpper,omitzero"`
	// Maximum mass at launch, lower bounds, in kilograms.
	MaxLaunchMassLower param.Opt[float64] `json:"maxLaunchMassLower,omitzero"`
	// Maximum mass at launch, upper bounds, in kilograms.
	MaxLaunchMassUpper param.Opt[float64] `json:"maxLaunchMassUpper,omitzero"`
	// Maximum payload mass available, in kilograms.
	MaxPayloadMass param.Opt[float64] `json:"maxPayloadMass,omitzero"`
	// Maximum payload power available, in kilowatts.
	MaxPayloadPower param.Opt[float64] `json:"maxPayloadPower,omitzero"`
	// Maximum power available on this bus, in kilowatts.
	MaxSpacecraftPower param.Opt[float64] `json:"maxSpacecraftPower,omitzero"`
	// Maximum mass of this bus with fuel, but without payloads, in kilograms.
	MaxWetMass param.Opt[float64] `json:"maxWetMass,omitzero"`
	// Median mass of this bus without payloads or fuel, in kilograms.
	MedianDryMass param.Opt[float64] `json:"medianDryMass,omitzero"`
	// Median mass of this bus with fuel, but without payloads, in kilograms.
	MedianWetMass param.Opt[float64] `json:"medianWetMass,omitzero"`
	// Minimum mass of this bus without payloads or fuel, in kilograms.
	MinDryMass param.Opt[float64] `json:"minDryMass,omitzero"`
	// Minimum mass of this bus with fuel, but without payloads, in kilograms.
	MinWetMass param.Opt[float64] `json:"minWetMass,omitzero"`
	// The number of orbit types this bus can support.
	NumOrbitType param.Opt[int64] `json:"numOrbitType,omitzero"`
	// Orbit averaged power (the power averaged over one orbit) available on this bus
	// for payloads, in kilowatts.
	OapPayloadPower param.Opt[float64] `json:"oapPayloadPower,omitzero"`
	// Orbit averaged power (the power averaged over one orbit) available on this bus,
	// in kilowatts.
	OapSpacecraftPower param.Opt[float64] `json:"oapSpacecraftPower,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The radial dimension available on this bus for payloads, in meters.
	PayloadDimensionX param.Opt[float64] `json:"payloadDimensionX,omitzero"`
	// The in-track dimension available on this bus for payloads, in meters.
	PayloadDimensionY param.Opt[float64] `json:"payloadDimensionY,omitzero"`
	// The cross-track dimension available on this bus for payloads, in meters.
	PayloadDimensionZ param.Opt[float64] `json:"payloadDimensionZ,omitzero"`
	// The volume available on this bus for payloads, in cubic meters.
	PayloadVolume param.Opt[float64] `json:"payloadVolume,omitzero"`
	// Power category of this bus (e.g. 0-1kW low power, etc).
	PowerCategory param.Opt[string] `json:"powerCategory,omitzero"`
	// Unique identifier of the organization which manufactures the telemetry tracking
	// and command subsystem for this bus.
	TelemetryTrackingManufacturerOrgID param.Opt[string] `json:"telemetryTrackingManufacturerOrgId,omitzero"`
	// Type of this bus.
	Type param.Opt[string] `json:"type,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	// Array of orbit types this bus can support (e.g. GEO, LEO, etc.). Must contain
	// the same number of elements as the value of numOrbitType.
	OrbitTypes []string `json:"orbitTypes,omitzero"`
	paramObj
}

func (r BusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BusNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BusNewParams) UnmarshalJSON(data []byte) error {
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
type BusNewParamsDataMode string

const (
	BusNewParamsDataModeReal      BusNewParamsDataMode = "REAL"
	BusNewParamsDataModeTest      BusNewParamsDataMode = "TEST"
	BusNewParamsDataModeSimulated BusNewParamsDataMode = "SIMULATED"
	BusNewParamsDataModeExercise  BusNewParamsDataMode = "EXERCISE"
)

type BusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BusGetParams]'s query parameters as `url.Values`.
func (r BusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BusUpdateParams struct {
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
	DataMode BusUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Name of this bus.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Attitude and Orbital Control Notes/description for the bus.
	AocsNotes param.Opt[string] `json:"aocsNotes,omitzero"`
	// Average mass of this bus without payloads or fuel, in kilograms.
	AvgDryMass param.Opt[float64] `json:"avgDryMass,omitzero"`
	// Average mass available on this bus for payloads, in kilograms.
	AvgPayloadMass param.Opt[float64] `json:"avgPayloadMass,omitzero"`
	// Average power available on this bus for payloads, in kilowatts.
	AvgPayloadPower param.Opt[float64] `json:"avgPayloadPower,omitzero"`
	// Average power available on this bus, in kilowatts.
	AvgSpacecraftPower param.Opt[float64] `json:"avgSpacecraftPower,omitzero"`
	// Average mass of this bus with fuel, but without payloads, in kilograms.
	AvgWetMass param.Opt[float64] `json:"avgWetMass,omitzero"`
	// Body dimension in X direction pertaining to length, in meters.
	BodyDimensionX param.Opt[float64] `json:"bodyDimensionX,omitzero"`
	// Body dimension in Y direction pertaining to height, in meters.
	BodyDimensionY param.Opt[float64] `json:"bodyDimensionY,omitzero"`
	// Body dimension in Z direction pertaining to width, in meters.
	BodyDimensionZ param.Opt[float64] `json:"bodyDimensionZ,omitzero"`
	// Unique identifier of the organization which designs the bus kit.
	BusKitDesignerOrgID param.Opt[string] `json:"busKitDesignerOrgId,omitzero"`
	// Country where this bus was manufactured. This value is typically the ISO 3166
	// Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Notes/description of the bus.
	Description param.Opt[string] `json:"description,omitzero"`
	// Boolean indicating if this bus is generic.
	Generic param.Opt[bool] `json:"generic,omitzero"`
	// ID of the parent entity for this bus.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Launch envelope dimension in X direction, in meters.
	LaunchEnvelopeDimensionX param.Opt[float64] `json:"launchEnvelopeDimensionX,omitzero"`
	// Launch envelope dimension in Y direction, in meters.
	LaunchEnvelopeDimensionY param.Opt[float64] `json:"launchEnvelopeDimensionY,omitzero"`
	// Launch envelope dimension in Z direction, in meters.
	LaunchEnvelopeDimensionZ param.Opt[float64] `json:"launchEnvelopeDimensionZ,omitzero"`
	// Unique identifier of the organization which manufactures the main onboard
	// computer for this bus.
	MainComputerManufacturerOrgID param.Opt[string] `json:"mainComputerManufacturerOrgId,omitzero"`
	// Unique identifier of the organization which manufactures this bus.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Mass category of this bus (e.g. 1 - 10 kg: Nanosatellite, 10 - 100 kg:
	// Microsatellite, 100 - 500 kg: Minisatellite, 1000 - 2500kg: Medium satellite,
	// etc.).
	MassCategory param.Opt[string] `json:"massCategory,omitzero"`
	// Maximum power at beginning of life, lower bounds, in kilowatts.
	MaxBolPowerLower param.Opt[float64] `json:"maxBOLPowerLower,omitzero"`
	// Maximum power at beginning of life, upper bounds, in kilowatts.
	MaxBolPowerUpper param.Opt[float64] `json:"maxBOLPowerUpper,omitzero"`
	// Maximum mass on station at beginning of life, in kilograms.
	MaxBolStationMass param.Opt[float64] `json:"maxBOLStationMass,omitzero"`
	// Maximum mass of this bus without payloads or fuel, in kilograms.
	MaxDryMass param.Opt[float64] `json:"maxDryMass,omitzero"`
	// Maximum power at end of life, lower bounds, in kilowatts.
	MaxEolPowerLower param.Opt[float64] `json:"maxEOLPowerLower,omitzero"`
	// Maximum power at end of life, upper bounds, in kilowatts.
	MaxEolPowerUpper param.Opt[float64] `json:"maxEOLPowerUpper,omitzero"`
	// Maximum mass at launch, lower bounds, in kilograms.
	MaxLaunchMassLower param.Opt[float64] `json:"maxLaunchMassLower,omitzero"`
	// Maximum mass at launch, upper bounds, in kilograms.
	MaxLaunchMassUpper param.Opt[float64] `json:"maxLaunchMassUpper,omitzero"`
	// Maximum payload mass available, in kilograms.
	MaxPayloadMass param.Opt[float64] `json:"maxPayloadMass,omitzero"`
	// Maximum payload power available, in kilowatts.
	MaxPayloadPower param.Opt[float64] `json:"maxPayloadPower,omitzero"`
	// Maximum power available on this bus, in kilowatts.
	MaxSpacecraftPower param.Opt[float64] `json:"maxSpacecraftPower,omitzero"`
	// Maximum mass of this bus with fuel, but without payloads, in kilograms.
	MaxWetMass param.Opt[float64] `json:"maxWetMass,omitzero"`
	// Median mass of this bus without payloads or fuel, in kilograms.
	MedianDryMass param.Opt[float64] `json:"medianDryMass,omitzero"`
	// Median mass of this bus with fuel, but without payloads, in kilograms.
	MedianWetMass param.Opt[float64] `json:"medianWetMass,omitzero"`
	// Minimum mass of this bus without payloads or fuel, in kilograms.
	MinDryMass param.Opt[float64] `json:"minDryMass,omitzero"`
	// Minimum mass of this bus with fuel, but without payloads, in kilograms.
	MinWetMass param.Opt[float64] `json:"minWetMass,omitzero"`
	// The number of orbit types this bus can support.
	NumOrbitType param.Opt[int64] `json:"numOrbitType,omitzero"`
	// Orbit averaged power (the power averaged over one orbit) available on this bus
	// for payloads, in kilowatts.
	OapPayloadPower param.Opt[float64] `json:"oapPayloadPower,omitzero"`
	// Orbit averaged power (the power averaged over one orbit) available on this bus,
	// in kilowatts.
	OapSpacecraftPower param.Opt[float64] `json:"oapSpacecraftPower,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The radial dimension available on this bus for payloads, in meters.
	PayloadDimensionX param.Opt[float64] `json:"payloadDimensionX,omitzero"`
	// The in-track dimension available on this bus for payloads, in meters.
	PayloadDimensionY param.Opt[float64] `json:"payloadDimensionY,omitzero"`
	// The cross-track dimension available on this bus for payloads, in meters.
	PayloadDimensionZ param.Opt[float64] `json:"payloadDimensionZ,omitzero"`
	// The volume available on this bus for payloads, in cubic meters.
	PayloadVolume param.Opt[float64] `json:"payloadVolume,omitzero"`
	// Power category of this bus (e.g. 0-1kW low power, etc).
	PowerCategory param.Opt[string] `json:"powerCategory,omitzero"`
	// Unique identifier of the organization which manufactures the telemetry tracking
	// and command subsystem for this bus.
	TelemetryTrackingManufacturerOrgID param.Opt[string] `json:"telemetryTrackingManufacturerOrgId,omitzero"`
	// Type of this bus.
	Type param.Opt[string] `json:"type,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	// Array of orbit types this bus can support (e.g. GEO, LEO, etc.). Must contain
	// the same number of elements as the value of numOrbitType.
	OrbitTypes []string `json:"orbitTypes,omitzero"`
	paramObj
}

func (r BusUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BusUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BusUpdateParams) UnmarshalJSON(data []byte) error {
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
type BusUpdateParamsDataMode string

const (
	BusUpdateParamsDataModeReal      BusUpdateParamsDataMode = "REAL"
	BusUpdateParamsDataModeTest      BusUpdateParamsDataMode = "TEST"
	BusUpdateParamsDataModeSimulated BusUpdateParamsDataMode = "SIMULATED"
	BusUpdateParamsDataModeExercise  BusUpdateParamsDataMode = "EXERCISE"
)

type BusListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BusListParams]'s query parameters as `url.Values`.
func (r BusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BusCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BusCountParams]'s query parameters as `url.Values`.
func (r BusCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BusTupleParams]'s query parameters as `url.Values`.
func (r BusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
