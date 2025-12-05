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

// AircraftStatusService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAircraftStatusService] method instead.
type AircraftStatusService struct {
	Options []option.RequestOption
	History AircraftStatusHistoryService
}

// NewAircraftStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAircraftStatusService(opts ...option.RequestOption) (r AircraftStatusService) {
	r = AircraftStatusService{}
	r.Options = opts
	r.History = NewAircraftStatusHistoryService(opts...)
	return
}

// Service operation to take a single AircraftStatus as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *AircraftStatusService) New(ctx context.Context, body AircraftStatusNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/aircraftstatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single AircraftStatus record by its unique ID passed
// as a path parameter.
func (r *AircraftStatusService) Get(ctx context.Context, id string, query AircraftStatusGetParams, opts ...option.RequestOption) (res *shared.AircraftstatusFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraftstatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single AircraftStatus. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *AircraftStatusService) Update(ctx context.Context, id string, body AircraftStatusUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraftstatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AircraftStatusService) List(ctx context.Context, query AircraftStatusListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AircraftstatusAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/aircraftstatus"
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
func (r *AircraftStatusService) ListAutoPaging(ctx context.Context, query AircraftStatusListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AircraftstatusAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Status object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *AircraftStatusService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraftstatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AircraftStatusService) Count(ctx context.Context, query AircraftStatusCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/aircraftstatus/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AircraftStatusService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *AircraftStatusQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/aircraftstatus/queryhelp"
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
func (r *AircraftStatusService) Tuple(ctx context.Context, query AircraftStatusTupleParams, opts ...option.RequestOption) (res *[]shared.AircraftstatusFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/aircraftstatus/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Aircraft readiness and status data. Contains the dynamic data associated with
// the specific aircraft status, either in-flight or on-ground, including remaining
// fuel, mission readiness, and inventory, etc.
type AircraftstatusAbridged struct {
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
	DataMode AircraftstatusAbridgedDataMode `json:"dataMode,required"`
	// Unique identifier of the aircraft.
	IDAircraft string `json:"idAircraft,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// List of additional operational systems on this aircraft beyond what is normally
	// available.
	AdditionalSys []string `json:"additionalSys"`
	// The status of the air-to-air weapon release system (OPERATIONAL,
	// NON-OPERATIONAL, OFF).
	//
	// Any of "OPERATIONAL", "NON-OPERATIONAL", "OFF".
	AirToAirStatus AircraftstatusAbridgedAirToAirStatus `json:"airToAirStatus"`
	// The status of the air-to-ground weapon release system (OPERATIONAL,
	// NON-OPERATIONAL, OFF).
	//
	// Any of "OPERATIONAL", "NON-OPERATIONAL", "OFF".
	AirToGroundStatus AircraftstatusAbridgedAirToGroundStatus `json:"airToGroundStatus"`
	// Aircraft alpha status code that indicates the aircraft maintenance status
	// estimated by the pilot.
	AlphaStatusCode string `json:"alphaStatusCode"`
	// Alternate Aircraft Identifier provided by source.
	AltAircraftID string `json:"altAircraftId"`
	// The contamination status of the aircraft (e.g. CLEAR, CONTAMINATED,
	// DECONTAMINATED, UNKNOWN, etc.).
	ContaminationStatus string `json:"contaminationStatus"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The International Civil Aviation Organization (ICAO) code at which this aircraft
	// is currently located or has most recently departed, if airborne.
	CurrentIcao string `json:"currentICAO"`
	// The current readiness state of the aircraft (e.g. AIRBORNE, ALERTCOCKED,
	// AVAILABLE, BATTLESTATION, RUNWAY ALERT, SUITUP).
	CurrentState string `json:"currentState"`
	// The earliest time that turnaround of the aircraft may complete, in ISO 8601 UTC
	// format with millisecond precision.
	EarliestTaEndTime time.Time `json:"earliestTAEndTime" format:"date-time"`
	// The Expected Time in Commission (ETIC) for this aircraft, in ISO 8601 UTC format
	// with millisecond precision. This is the estimated time when the issue will be
	// resolved.
	Etic time.Time `json:"etic" format:"date-time"`
	// Current flight phase (e.g. AIR REFUELING, GROUND, LANDING, etc.) of the
	// aircraft.
	FlightPhase string `json:"flightPhase"`
	// The mass of fuel remaining on the aircraft, in kilograms.
	Fuel int64 `json:"fuel"`
	// Used in conjunction with the fuel field to indicate either burnable or offload
	// fuel.
	FuelFunction string `json:"fuelFunction"`
	// The state of the aircraft fuel status (e.g. DELIVERED, DUMPED, EMPTY, FULL,
	// OTHER, REQUESTED, etc.).
	FuelStatus string `json:"fuelStatus"`
	// US Air Force geographic location code of the airfield where the aircraft is
	// located.
	GeoLoc string `json:"geoLoc"`
	// The ground status of the aircraft (e.g. ALERT, CREW READY, ENGINE START, HANGAR,
	// etc.).
	GroundStatus string `json:"groundStatus"`
	// Flag indicating that the aircraft is capable of making at least one gun pass.
	GunCapable bool `json:"gunCapable"`
	// The upper bound of the estimated number of gun rounds available.
	GunRdsMax int64 `json:"gunRdsMax"`
	// The lower bound of the estimated number of gun rounds available.
	GunRdsMin int64 `json:"gunRdsMin"`
	// The type of gun rounds available (e.g. 7.62 MM, 20 MM, 25 MM, etc.).
	GunRdsType string `json:"gunRdsType"`
	// If not airborne, the unique identifier of the installation currently hosting the
	// aircraft.
	IDAirfield string `json:"idAirfield"`
	// Unique identifier of the Point of Interest (POI) record related to this aircraft
	// status. This will generally represent the location of an aircraft on the ground.
	IDPoi string `json:"idPOI"`
	// Array of inventory item(s) for which estimate(s) are available (e.g. AIM-9
	// SIDEWINDER, AIM-120 AMRAAM, AIM-92 STINGER, CHAFF DECOY, FLARE TP 400, etc.).
	// Intended as, but not constrained to, MIL-STD-6016 environment dependent
	// specific/store type designations. This array must be the same length as
	// inventoryMin and inventoryMax.
	Inventory []string `json:"inventory"`
	// Array of the upper bound quantity for each of the inventory items. The values in
	// this array must correspond to position index in the inventory array. This array
	// must be the same length as inventory and inventoryMin.
	InventoryMax []int64 `json:"inventoryMax"`
	// Array of the lower bound quantity for each of the inventory items. The values in
	// this array must correspond to position index in the inventory array. This array
	// must be the same length as inventory and inventoryMax.
	InventoryMin []int64 `json:"inventoryMin"`
	// Date when the military aircraft inspection was last performed, in ISO 8601 UTC
	// format with millisecond precision.
	LastInspectionDate time.Time `json:"lastInspectionDate" format:"date-time"`
	// The name or ID of the external user that updated this status.
	LastUpdatedBy string `json:"lastUpdatedBy"`
	// Military aircraft maintenance point of contact for this aircraft.
	MaintPoc string `json:"maintPoc"`
	// Indicates the priority of the maintenance effort.
	MaintPriority string `json:"maintPriority"`
	// The maintenance status of the aircraft.
	MaintStatus string `json:"maintStatus"`
	// Indicates the maintenance discrepancy that drives the current maintenance
	// status.
	MaintStatusDriver string `json:"maintStatusDriver"`
	// The time of the last maintenance status update, in ISO 8601 UTC format with
	// millisecond precision.
	MaintStatusUpdate time.Time `json:"maintStatusUpdate" format:"date-time"`
	// The Operational Capability of the reported aircraft (ABLE, LOFUEL, UNABLE).
	MissionReadiness string `json:"missionReadiness"`
	// Maintenance pacing remarks assocociated with this aircraft.
	MxRemark string `json:"mxRemark"`
	// The International Civil Aviation Organization (ICAO) code of the next
	// destination of this aircraft.
	NextIcao string `json:"nextICAO"`
	// Optional notes/comments concerning this aircraft status.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The parking location of this aircraft.
	ParkLocation string `json:"parkLocation"`
	// The system that designated the parking location (e.g. EMOC, GDSS, PEX, etc.).
	ParkLocationSystem string `json:"parkLocationSystem"`
	// The International Civil Aviation Organization (ICAO) code at which this aircraft
	// was previously located.
	PreviousIcao string `json:"previousICAO"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The turnaround start time, in ISO 8601 UTC format with millisecond precision.
	TaStartTime time.Time `json:"taStartTime" format:"date-time"`
	// Estimated Time for Completion (ETIC) of an aircraft issue, in ISO 8601 UTC
	// format with millisecond precision. This is the estimated time when the course of
	// action to resolve the issue will be determined.
	TroubleshootEtic time.Time `json:"troubleshootEtic" format:"date-time"`
	// List of unavailable systems that would normally be on this aircraft.
	UnavailableSys []string `json:"unavailableSys"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDAircraft            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AdditionalSys         respjson.Field
		AirToAirStatus        respjson.Field
		AirToGroundStatus     respjson.Field
		AlphaStatusCode       respjson.Field
		AltAircraftID         respjson.Field
		ContaminationStatus   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CurrentIcao           respjson.Field
		CurrentState          respjson.Field
		EarliestTaEndTime     respjson.Field
		Etic                  respjson.Field
		FlightPhase           respjson.Field
		Fuel                  respjson.Field
		FuelFunction          respjson.Field
		FuelStatus            respjson.Field
		GeoLoc                respjson.Field
		GroundStatus          respjson.Field
		GunCapable            respjson.Field
		GunRdsMax             respjson.Field
		GunRdsMin             respjson.Field
		GunRdsType            respjson.Field
		IDAirfield            respjson.Field
		IDPoi                 respjson.Field
		Inventory             respjson.Field
		InventoryMax          respjson.Field
		InventoryMin          respjson.Field
		LastInspectionDate    respjson.Field
		LastUpdatedBy         respjson.Field
		MaintPoc              respjson.Field
		MaintPriority         respjson.Field
		MaintStatus           respjson.Field
		MaintStatusDriver     respjson.Field
		MaintStatusUpdate     respjson.Field
		MissionReadiness      respjson.Field
		MxRemark              respjson.Field
		NextIcao              respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ParkLocation          respjson.Field
		ParkLocationSystem    respjson.Field
		PreviousIcao          respjson.Field
		SourceDl              respjson.Field
		TaStartTime           respjson.Field
		TroubleshootEtic      respjson.Field
		UnavailableSys        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AircraftstatusAbridged) RawJSON() string { return r.JSON.raw }
func (r *AircraftstatusAbridged) UnmarshalJSON(data []byte) error {
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
type AircraftstatusAbridgedDataMode string

const (
	AircraftstatusAbridgedDataModeReal      AircraftstatusAbridgedDataMode = "REAL"
	AircraftstatusAbridgedDataModeTest      AircraftstatusAbridgedDataMode = "TEST"
	AircraftstatusAbridgedDataModeSimulated AircraftstatusAbridgedDataMode = "SIMULATED"
	AircraftstatusAbridgedDataModeExercise  AircraftstatusAbridgedDataMode = "EXERCISE"
)

// The status of the air-to-air weapon release system (OPERATIONAL,
// NON-OPERATIONAL, OFF).
type AircraftstatusAbridgedAirToAirStatus string

const (
	AircraftstatusAbridgedAirToAirStatusOperational    AircraftstatusAbridgedAirToAirStatus = "OPERATIONAL"
	AircraftstatusAbridgedAirToAirStatusNonOperational AircraftstatusAbridgedAirToAirStatus = "NON-OPERATIONAL"
	AircraftstatusAbridgedAirToAirStatusOff            AircraftstatusAbridgedAirToAirStatus = "OFF"
)

// The status of the air-to-ground weapon release system (OPERATIONAL,
// NON-OPERATIONAL, OFF).
type AircraftstatusAbridgedAirToGroundStatus string

const (
	AircraftstatusAbridgedAirToGroundStatusOperational    AircraftstatusAbridgedAirToGroundStatus = "OPERATIONAL"
	AircraftstatusAbridgedAirToGroundStatusNonOperational AircraftstatusAbridgedAirToGroundStatus = "NON-OPERATIONAL"
	AircraftstatusAbridgedAirToGroundStatusOff            AircraftstatusAbridgedAirToGroundStatus = "OFF"
)

type AircraftStatusQueryhelpResponse struct {
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
func (r AircraftStatusQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *AircraftStatusQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AircraftStatusNewParams struct {
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
	DataMode AircraftStatusNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the aircraft.
	IDAircraft string `json:"idAircraft,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Aircraft alpha status code that indicates the aircraft maintenance status
	// estimated by the pilot.
	AlphaStatusCode param.Opt[string] `json:"alphaStatusCode,omitzero"`
	// Alternate Aircraft Identifier provided by source.
	AltAircraftID param.Opt[string] `json:"altAircraftId,omitzero"`
	// The contamination status of the aircraft (e.g. CLEAR, CONTAMINATED,
	// DECONTAMINATED, UNKNOWN, etc.).
	ContaminationStatus param.Opt[string] `json:"contaminationStatus,omitzero"`
	// The International Civil Aviation Organization (ICAO) code at which this aircraft
	// is currently located or has most recently departed, if airborne.
	CurrentIcao param.Opt[string] `json:"currentICAO,omitzero"`
	// The current readiness state of the aircraft (e.g. AIRBORNE, ALERTCOCKED,
	// AVAILABLE, BATTLESTATION, RUNWAY ALERT, SUITUP).
	CurrentState param.Opt[string] `json:"currentState,omitzero"`
	// The earliest time that turnaround of the aircraft may complete, in ISO 8601 UTC
	// format with millisecond precision.
	EarliestTaEndTime param.Opt[time.Time] `json:"earliestTAEndTime,omitzero" format:"date-time"`
	// The Expected Time in Commission (ETIC) for this aircraft, in ISO 8601 UTC format
	// with millisecond precision. This is the estimated time when the issue will be
	// resolved.
	Etic param.Opt[time.Time] `json:"etic,omitzero" format:"date-time"`
	// Current flight phase (e.g. AIR REFUELING, GROUND, LANDING, etc.) of the
	// aircraft.
	FlightPhase param.Opt[string] `json:"flightPhase,omitzero"`
	// The mass of fuel remaining on the aircraft, in kilograms.
	Fuel param.Opt[int64] `json:"fuel,omitzero"`
	// Used in conjunction with the fuel field to indicate either burnable or offload
	// fuel.
	FuelFunction param.Opt[string] `json:"fuelFunction,omitzero"`
	// The state of the aircraft fuel status (e.g. DELIVERED, DUMPED, EMPTY, FULL,
	// OTHER, REQUESTED, etc.).
	FuelStatus param.Opt[string] `json:"fuelStatus,omitzero"`
	// US Air Force geographic location code of the airfield where the aircraft is
	// located.
	GeoLoc param.Opt[string] `json:"geoLoc,omitzero"`
	// The ground status of the aircraft (e.g. ALERT, CREW READY, ENGINE START, HANGAR,
	// etc.).
	GroundStatus param.Opt[string] `json:"groundStatus,omitzero"`
	// Flag indicating that the aircraft is capable of making at least one gun pass.
	GunCapable param.Opt[bool] `json:"gunCapable,omitzero"`
	// The upper bound of the estimated number of gun rounds available.
	GunRdsMax param.Opt[int64] `json:"gunRdsMax,omitzero"`
	// The lower bound of the estimated number of gun rounds available.
	GunRdsMin param.Opt[int64] `json:"gunRdsMin,omitzero"`
	// The type of gun rounds available (e.g. 7.62 MM, 20 MM, 25 MM, etc.).
	GunRdsType param.Opt[string] `json:"gunRdsType,omitzero"`
	// If not airborne, the unique identifier of the installation currently hosting the
	// aircraft.
	IDAirfield param.Opt[string] `json:"idAirfield,omitzero"`
	// Unique identifier of the Point of Interest (POI) record related to this aircraft
	// status. This will generally represent the location of an aircraft on the ground.
	IDPoi param.Opt[string] `json:"idPOI,omitzero"`
	// Date when the military aircraft inspection was last performed, in ISO 8601 UTC
	// format with millisecond precision.
	LastInspectionDate param.Opt[time.Time] `json:"lastInspectionDate,omitzero" format:"date-time"`
	// The name or ID of the external user that updated this status.
	LastUpdatedBy param.Opt[string] `json:"lastUpdatedBy,omitzero"`
	// Military aircraft maintenance point of contact for this aircraft.
	MaintPoc param.Opt[string] `json:"maintPoc,omitzero"`
	// Indicates the priority of the maintenance effort.
	MaintPriority param.Opt[string] `json:"maintPriority,omitzero"`
	// The maintenance status of the aircraft.
	MaintStatus param.Opt[string] `json:"maintStatus,omitzero"`
	// Indicates the maintenance discrepancy that drives the current maintenance
	// status.
	MaintStatusDriver param.Opt[string] `json:"maintStatusDriver,omitzero"`
	// The time of the last maintenance status update, in ISO 8601 UTC format with
	// millisecond precision.
	MaintStatusUpdate param.Opt[time.Time] `json:"maintStatusUpdate,omitzero" format:"date-time"`
	// The Operational Capability of the reported aircraft (ABLE, LOFUEL, UNABLE).
	MissionReadiness param.Opt[string] `json:"missionReadiness,omitzero"`
	// Maintenance pacing remarks assocociated with this aircraft.
	MxRemark param.Opt[string] `json:"mxRemark,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the next
	// destination of this aircraft.
	NextIcao param.Opt[string] `json:"nextICAO,omitzero"`
	// Optional notes/comments concerning this aircraft status.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The parking location of this aircraft.
	ParkLocation param.Opt[string] `json:"parkLocation,omitzero"`
	// The system that designated the parking location (e.g. EMOC, GDSS, PEX, etc.).
	ParkLocationSystem param.Opt[string] `json:"parkLocationSystem,omitzero"`
	// The International Civil Aviation Organization (ICAO) code at which this aircraft
	// was previously located.
	PreviousIcao param.Opt[string] `json:"previousICAO,omitzero"`
	// The turnaround start time, in ISO 8601 UTC format with millisecond precision.
	TaStartTime param.Opt[time.Time] `json:"taStartTime,omitzero" format:"date-time"`
	// Estimated Time for Completion (ETIC) of an aircraft issue, in ISO 8601 UTC
	// format with millisecond precision. This is the estimated time when the course of
	// action to resolve the issue will be determined.
	TroubleshootEtic param.Opt[time.Time] `json:"troubleshootEtic,omitzero" format:"date-time"`
	// List of additional operational systems on this aircraft beyond what is normally
	// available.
	AdditionalSys []string `json:"additionalSys,omitzero"`
	// The status of the air-to-air weapon release system (OPERATIONAL,
	// NON-OPERATIONAL, OFF).
	//
	// Any of "OPERATIONAL", "NON-OPERATIONAL", "OFF".
	AirToAirStatus AircraftStatusNewParamsAirToAirStatus `json:"airToAirStatus,omitzero"`
	// The status of the air-to-ground weapon release system (OPERATIONAL,
	// NON-OPERATIONAL, OFF).
	//
	// Any of "OPERATIONAL", "NON-OPERATIONAL", "OFF".
	AirToGroundStatus AircraftStatusNewParamsAirToGroundStatus `json:"airToGroundStatus,omitzero"`
	// Array of inventory item(s) for which estimate(s) are available (e.g. AIM-9
	// SIDEWINDER, AIM-120 AMRAAM, AIM-92 STINGER, CHAFF DECOY, FLARE TP 400, etc.).
	// Intended as, but not constrained to, MIL-STD-6016 environment dependent
	// specific/store type designations. This array must be the same length as
	// inventoryMin and inventoryMax.
	Inventory []string `json:"inventory,omitzero"`
	// Array of the upper bound quantity for each of the inventory items. The values in
	// this array must correspond to position index in the inventory array. This array
	// must be the same length as inventory and inventoryMin.
	InventoryMax []int64 `json:"inventoryMax,omitzero"`
	// Array of the lower bound quantity for each of the inventory items. The values in
	// this array must correspond to position index in the inventory array. This array
	// must be the same length as inventory and inventoryMax.
	InventoryMin []int64 `json:"inventoryMin,omitzero"`
	// List of unavailable systems that would normally be on this aircraft.
	UnavailableSys []string `json:"unavailableSys,omitzero"`
	paramObj
}

func (r AircraftStatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AircraftStatusNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AircraftStatusNewParams) UnmarshalJSON(data []byte) error {
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
type AircraftStatusNewParamsDataMode string

const (
	AircraftStatusNewParamsDataModeReal      AircraftStatusNewParamsDataMode = "REAL"
	AircraftStatusNewParamsDataModeTest      AircraftStatusNewParamsDataMode = "TEST"
	AircraftStatusNewParamsDataModeSimulated AircraftStatusNewParamsDataMode = "SIMULATED"
	AircraftStatusNewParamsDataModeExercise  AircraftStatusNewParamsDataMode = "EXERCISE"
)

// The status of the air-to-air weapon release system (OPERATIONAL,
// NON-OPERATIONAL, OFF).
type AircraftStatusNewParamsAirToAirStatus string

const (
	AircraftStatusNewParamsAirToAirStatusOperational    AircraftStatusNewParamsAirToAirStatus = "OPERATIONAL"
	AircraftStatusNewParamsAirToAirStatusNonOperational AircraftStatusNewParamsAirToAirStatus = "NON-OPERATIONAL"
	AircraftStatusNewParamsAirToAirStatusOff            AircraftStatusNewParamsAirToAirStatus = "OFF"
)

// The status of the air-to-ground weapon release system (OPERATIONAL,
// NON-OPERATIONAL, OFF).
type AircraftStatusNewParamsAirToGroundStatus string

const (
	AircraftStatusNewParamsAirToGroundStatusOperational    AircraftStatusNewParamsAirToGroundStatus = "OPERATIONAL"
	AircraftStatusNewParamsAirToGroundStatusNonOperational AircraftStatusNewParamsAirToGroundStatus = "NON-OPERATIONAL"
	AircraftStatusNewParamsAirToGroundStatusOff            AircraftStatusNewParamsAirToGroundStatus = "OFF"
)

type AircraftStatusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftStatusGetParams]'s query parameters as
// `url.Values`.
func (r AircraftStatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftStatusUpdateParams struct {
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
	DataMode AircraftStatusUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the aircraft.
	IDAircraft string `json:"idAircraft,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Aircraft alpha status code that indicates the aircraft maintenance status
	// estimated by the pilot.
	AlphaStatusCode param.Opt[string] `json:"alphaStatusCode,omitzero"`
	// Alternate Aircraft Identifier provided by source.
	AltAircraftID param.Opt[string] `json:"altAircraftId,omitzero"`
	// The contamination status of the aircraft (e.g. CLEAR, CONTAMINATED,
	// DECONTAMINATED, UNKNOWN, etc.).
	ContaminationStatus param.Opt[string] `json:"contaminationStatus,omitzero"`
	// The International Civil Aviation Organization (ICAO) code at which this aircraft
	// is currently located or has most recently departed, if airborne.
	CurrentIcao param.Opt[string] `json:"currentICAO,omitzero"`
	// The current readiness state of the aircraft (e.g. AIRBORNE, ALERTCOCKED,
	// AVAILABLE, BATTLESTATION, RUNWAY ALERT, SUITUP).
	CurrentState param.Opt[string] `json:"currentState,omitzero"`
	// The earliest time that turnaround of the aircraft may complete, in ISO 8601 UTC
	// format with millisecond precision.
	EarliestTaEndTime param.Opt[time.Time] `json:"earliestTAEndTime,omitzero" format:"date-time"`
	// The Expected Time in Commission (ETIC) for this aircraft, in ISO 8601 UTC format
	// with millisecond precision. This is the estimated time when the issue will be
	// resolved.
	Etic param.Opt[time.Time] `json:"etic,omitzero" format:"date-time"`
	// Current flight phase (e.g. AIR REFUELING, GROUND, LANDING, etc.) of the
	// aircraft.
	FlightPhase param.Opt[string] `json:"flightPhase,omitzero"`
	// The mass of fuel remaining on the aircraft, in kilograms.
	Fuel param.Opt[int64] `json:"fuel,omitzero"`
	// Used in conjunction with the fuel field to indicate either burnable or offload
	// fuel.
	FuelFunction param.Opt[string] `json:"fuelFunction,omitzero"`
	// The state of the aircraft fuel status (e.g. DELIVERED, DUMPED, EMPTY, FULL,
	// OTHER, REQUESTED, etc.).
	FuelStatus param.Opt[string] `json:"fuelStatus,omitzero"`
	// US Air Force geographic location code of the airfield where the aircraft is
	// located.
	GeoLoc param.Opt[string] `json:"geoLoc,omitzero"`
	// The ground status of the aircraft (e.g. ALERT, CREW READY, ENGINE START, HANGAR,
	// etc.).
	GroundStatus param.Opt[string] `json:"groundStatus,omitzero"`
	// Flag indicating that the aircraft is capable of making at least one gun pass.
	GunCapable param.Opt[bool] `json:"gunCapable,omitzero"`
	// The upper bound of the estimated number of gun rounds available.
	GunRdsMax param.Opt[int64] `json:"gunRdsMax,omitzero"`
	// The lower bound of the estimated number of gun rounds available.
	GunRdsMin param.Opt[int64] `json:"gunRdsMin,omitzero"`
	// The type of gun rounds available (e.g. 7.62 MM, 20 MM, 25 MM, etc.).
	GunRdsType param.Opt[string] `json:"gunRdsType,omitzero"`
	// If not airborne, the unique identifier of the installation currently hosting the
	// aircraft.
	IDAirfield param.Opt[string] `json:"idAirfield,omitzero"`
	// Unique identifier of the Point of Interest (POI) record related to this aircraft
	// status. This will generally represent the location of an aircraft on the ground.
	IDPoi param.Opt[string] `json:"idPOI,omitzero"`
	// Date when the military aircraft inspection was last performed, in ISO 8601 UTC
	// format with millisecond precision.
	LastInspectionDate param.Opt[time.Time] `json:"lastInspectionDate,omitzero" format:"date-time"`
	// The name or ID of the external user that updated this status.
	LastUpdatedBy param.Opt[string] `json:"lastUpdatedBy,omitzero"`
	// Military aircraft maintenance point of contact for this aircraft.
	MaintPoc param.Opt[string] `json:"maintPoc,omitzero"`
	// Indicates the priority of the maintenance effort.
	MaintPriority param.Opt[string] `json:"maintPriority,omitzero"`
	// The maintenance status of the aircraft.
	MaintStatus param.Opt[string] `json:"maintStatus,omitzero"`
	// Indicates the maintenance discrepancy that drives the current maintenance
	// status.
	MaintStatusDriver param.Opt[string] `json:"maintStatusDriver,omitzero"`
	// The time of the last maintenance status update, in ISO 8601 UTC format with
	// millisecond precision.
	MaintStatusUpdate param.Opt[time.Time] `json:"maintStatusUpdate,omitzero" format:"date-time"`
	// The Operational Capability of the reported aircraft (ABLE, LOFUEL, UNABLE).
	MissionReadiness param.Opt[string] `json:"missionReadiness,omitzero"`
	// Maintenance pacing remarks assocociated with this aircraft.
	MxRemark param.Opt[string] `json:"mxRemark,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the next
	// destination of this aircraft.
	NextIcao param.Opt[string] `json:"nextICAO,omitzero"`
	// Optional notes/comments concerning this aircraft status.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The parking location of this aircraft.
	ParkLocation param.Opt[string] `json:"parkLocation,omitzero"`
	// The system that designated the parking location (e.g. EMOC, GDSS, PEX, etc.).
	ParkLocationSystem param.Opt[string] `json:"parkLocationSystem,omitzero"`
	// The International Civil Aviation Organization (ICAO) code at which this aircraft
	// was previously located.
	PreviousIcao param.Opt[string] `json:"previousICAO,omitzero"`
	// The turnaround start time, in ISO 8601 UTC format with millisecond precision.
	TaStartTime param.Opt[time.Time] `json:"taStartTime,omitzero" format:"date-time"`
	// Estimated Time for Completion (ETIC) of an aircraft issue, in ISO 8601 UTC
	// format with millisecond precision. This is the estimated time when the course of
	// action to resolve the issue will be determined.
	TroubleshootEtic param.Opt[time.Time] `json:"troubleshootEtic,omitzero" format:"date-time"`
	// List of additional operational systems on this aircraft beyond what is normally
	// available.
	AdditionalSys []string `json:"additionalSys,omitzero"`
	// The status of the air-to-air weapon release system (OPERATIONAL,
	// NON-OPERATIONAL, OFF).
	//
	// Any of "OPERATIONAL", "NON-OPERATIONAL", "OFF".
	AirToAirStatus AircraftStatusUpdateParamsAirToAirStatus `json:"airToAirStatus,omitzero"`
	// The status of the air-to-ground weapon release system (OPERATIONAL,
	// NON-OPERATIONAL, OFF).
	//
	// Any of "OPERATIONAL", "NON-OPERATIONAL", "OFF".
	AirToGroundStatus AircraftStatusUpdateParamsAirToGroundStatus `json:"airToGroundStatus,omitzero"`
	// Array of inventory item(s) for which estimate(s) are available (e.g. AIM-9
	// SIDEWINDER, AIM-120 AMRAAM, AIM-92 STINGER, CHAFF DECOY, FLARE TP 400, etc.).
	// Intended as, but not constrained to, MIL-STD-6016 environment dependent
	// specific/store type designations. This array must be the same length as
	// inventoryMin and inventoryMax.
	Inventory []string `json:"inventory,omitzero"`
	// Array of the upper bound quantity for each of the inventory items. The values in
	// this array must correspond to position index in the inventory array. This array
	// must be the same length as inventory and inventoryMin.
	InventoryMax []int64 `json:"inventoryMax,omitzero"`
	// Array of the lower bound quantity for each of the inventory items. The values in
	// this array must correspond to position index in the inventory array. This array
	// must be the same length as inventory and inventoryMax.
	InventoryMin []int64 `json:"inventoryMin,omitzero"`
	// List of unavailable systems that would normally be on this aircraft.
	UnavailableSys []string `json:"unavailableSys,omitzero"`
	paramObj
}

func (r AircraftStatusUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AircraftStatusUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AircraftStatusUpdateParams) UnmarshalJSON(data []byte) error {
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
type AircraftStatusUpdateParamsDataMode string

const (
	AircraftStatusUpdateParamsDataModeReal      AircraftStatusUpdateParamsDataMode = "REAL"
	AircraftStatusUpdateParamsDataModeTest      AircraftStatusUpdateParamsDataMode = "TEST"
	AircraftStatusUpdateParamsDataModeSimulated AircraftStatusUpdateParamsDataMode = "SIMULATED"
	AircraftStatusUpdateParamsDataModeExercise  AircraftStatusUpdateParamsDataMode = "EXERCISE"
)

// The status of the air-to-air weapon release system (OPERATIONAL,
// NON-OPERATIONAL, OFF).
type AircraftStatusUpdateParamsAirToAirStatus string

const (
	AircraftStatusUpdateParamsAirToAirStatusOperational    AircraftStatusUpdateParamsAirToAirStatus = "OPERATIONAL"
	AircraftStatusUpdateParamsAirToAirStatusNonOperational AircraftStatusUpdateParamsAirToAirStatus = "NON-OPERATIONAL"
	AircraftStatusUpdateParamsAirToAirStatusOff            AircraftStatusUpdateParamsAirToAirStatus = "OFF"
)

// The status of the air-to-ground weapon release system (OPERATIONAL,
// NON-OPERATIONAL, OFF).
type AircraftStatusUpdateParamsAirToGroundStatus string

const (
	AircraftStatusUpdateParamsAirToGroundStatusOperational    AircraftStatusUpdateParamsAirToGroundStatus = "OPERATIONAL"
	AircraftStatusUpdateParamsAirToGroundStatusNonOperational AircraftStatusUpdateParamsAirToGroundStatus = "NON-OPERATIONAL"
	AircraftStatusUpdateParamsAirToGroundStatusOff            AircraftStatusUpdateParamsAirToGroundStatus = "OFF"
)

type AircraftStatusListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftStatusListParams]'s query parameters as
// `url.Values`.
func (r AircraftStatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftStatusCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftStatusCountParams]'s query parameters as
// `url.Values`.
func (r AircraftStatusCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftStatusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftStatusTupleParams]'s query parameters as
// `url.Values`.
func (r AircraftStatusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
