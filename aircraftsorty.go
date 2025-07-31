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
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// AircraftSortyService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAircraftSortyService] method instead.
type AircraftSortyService struct {
	Options []option.RequestOption
}

// NewAircraftSortyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAircraftSortyService(opts ...option.RequestOption) (r AircraftSortyService) {
	r = AircraftSortyService{}
	r.Options = opts
	return
}

// Service operation to get a single AircraftSortie record by its unique ID passed
// as a path parameter.
func (r *AircraftSortyService) Get(ctx context.Context, id string, query AircraftSortyGetParams, opts ...option.RequestOption) (res *shared.AircraftsortieFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraftsortie/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single AircraftSortie. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *AircraftSortyService) Update(ctx context.Context, id string, body AircraftSortyUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraftsortie/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AircraftSortyService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *AircraftSortyQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/aircraftsortie/queryhelp"
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
func (r *AircraftSortyService) Tuple(ctx context.Context, query AircraftSortyTupleParams, opts ...option.RequestOption) (res *[]shared.AircraftsortieFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/aircraftsortie/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AircraftSortyQueryhelpResponse struct {
	AodrSupported         bool                                      `json:"aodrSupported"`
	ClassificationMarking string                                    `json:"classificationMarking"`
	Description           string                                    `json:"description"`
	HistorySupported      bool                                      `json:"historySupported"`
	Name                  string                                    `json:"name"`
	Parameters            []AircraftSortyQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                  `json:"requiredRoles"`
	RestSupported         bool                                      `json:"restSupported"`
	SortSupported         bool                                      `json:"sortSupported"`
	TypeName              string                                    `json:"typeName"`
	Uri                   string                                    `json:"uri"`
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
func (r AircraftSortyQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *AircraftSortyQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AircraftSortyQueryhelpResponseParameter struct {
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
func (r AircraftSortyQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *AircraftSortyQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AircraftSortyGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftSortyGetParams]'s query parameters as `url.Values`.
func (r AircraftSortyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftSortyUpdateParams struct {
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
	DataMode AircraftSortyUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The scheduled time that the Aircraft sortie is planned to depart, in ISO 8601
	// UTC format with millisecond precision.
	PlannedDepTime time.Time `json:"plannedDepTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The actual arrival time, in ISO 8601 UTC format with millisecond precision.
	ActualArrTime param.Opt[time.Time] `json:"actualArrTime,omitzero" format:"date-time"`
	// The actual time the Aircraft comes to a complete stop in its parking position,
	// in ISO 8601 UTC format with millisecond precision.
	ActualBlockInTime param.Opt[time.Time] `json:"actualBlockInTime,omitzero" format:"date-time"`
	// The actual time the Aircraft begins to taxi from its parking position, in ISO
	// 8601 UTC format with millisecond precision.
	ActualBlockOutTime param.Opt[time.Time] `json:"actualBlockOutTime,omitzero" format:"date-time"`
	// The actual departure time, in ISO 8601 UTC format.
	ActualDepTime param.Opt[time.Time] `json:"actualDepTime,omitzero" format:"date-time"`
	// The Automatic Dependent Surveillance-Broadcast (ADS-B) device identifier.
	AircraftAdsb param.Opt[string] `json:"aircraftADSB,omitzero"`
	// Alternate Aircraft Identifier provided by source.
	AircraftAltID param.Opt[string] `json:"aircraftAltId,omitzero"`
	// Aircraft event text.
	AircraftEvent param.Opt[string] `json:"aircraftEvent,omitzero"`
	// The aircraft Model Design Series designation assigned to this sortie.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Remarks concerning the aircraft.
	AircraftRemarks param.Opt[string] `json:"aircraftRemarks,omitzero"`
	// The amount of time allowed between launch order and takeoff, in seconds.
	AlertStatus param.Opt[int64] `json:"alertStatus,omitzero"`
	// The Alert Status code.
	AlertStatusCode param.Opt[string] `json:"alertStatusCode,omitzero"`
	// The Air Mobility Command (AMC) mission number of the sortie.
	AmcMsnNum param.Opt[string] `json:"amcMsnNum,omitzero"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	AmcMsnType param.Opt[string] `json:"amcMsnType,omitzero"`
	// The arrival Federal Aviation Administration (FAA) code of this sortie.
	ArrFaa param.Opt[string] `json:"arrFAA,omitzero"`
	// The arrival International Aviation Transport Association (IATA) code of this
	// sortie.
	ArrIata param.Opt[string] `json:"arrIATA,omitzero"`
	// The arrival International Civil Aviation Organization (ICAO) of this sortie.
	ArrIcao param.Opt[string] `json:"arrICAO,omitzero"`
	// The itinerary identifier of the arrival location.
	ArrItinerary param.Opt[int64] `json:"arrItinerary,omitzero"`
	// Purpose code at the arrival location of this sortie.
	ArrPurposeCode param.Opt[string] `json:"arrPurposeCode,omitzero"`
	// The call sign assigned to the aircraft on this sortie.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Description of the cargo configuration (e.g. C-1, C-2, C-3, DV-1, DV-2, AE-1,
	// etc.) currently on board the aircraft. Configuration meanings are determined by
	// the data source.
	CargoConfig param.Opt[string] `json:"cargoConfig,omitzero"`
	// The last name of the aircraft commander.
	CommanderName param.Opt[string] `json:"commanderName,omitzero"`
	// The current state of this sortie.
	CurrentState param.Opt[string] `json:"currentState,omitzero"`
	// The primary delay code.
	DelayCode param.Opt[string] `json:"delayCode,omitzero"`
	// The departure Federal Aviation Administration (FAA) code of this sortie.
	DepFaa param.Opt[string] `json:"depFAA,omitzero"`
	// The departure International Aviation Transport Association (IATA) code of this
	// sortie.
	DepIata param.Opt[string] `json:"depIATA,omitzero"`
	// The departure International Civil Aviation Organization (ICAO) of this sortie.
	DepIcao param.Opt[string] `json:"depICAO,omitzero"`
	// The itinerary identifier of the departure location.
	DepItinerary param.Opt[int64] `json:"depItinerary,omitzero"`
	// Purpose code at the departure location of this sortie.
	DepPurposeCode param.Opt[string] `json:"depPurposeCode,omitzero"`
	// Due home date by which the aircraft must return to its home station, in ISO 8601
	// UTC format with millisecond precision.
	Dhd param.Opt[time.Time] `json:"dhd,omitzero" format:"date-time"`
	// Reason the aircraft must return to home station by its due home date.
	DhdReason param.Opt[string] `json:"dhdReason,omitzero"`
	// The current estimated time that the Aircraft is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// The estimated time the Aircraft will come to a complete stop in its parking
	// position, in ISO 8601 UTC format with millisecond precision.
	EstBlockInTime param.Opt[time.Time] `json:"estBlockInTime,omitzero" format:"date-time"`
	// The estimated time the Aircraft will begin to taxi from its parking position, in
	// ISO 8601 UTC format with millisecond precision.
	EstBlockOutTime param.Opt[time.Time] `json:"estBlockOutTime,omitzero" format:"date-time"`
	// The current estimated time that the Aircraft is planned to depart, in ISO 8601
	// UTC format with millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// The planned flight time for this sortie, in minutes.
	FlightTime param.Opt[float64] `json:"flightTime,omitzero"`
	// Desk phone number of the flight manager assigned to the sortie. Null when no
	// flight manager is assigned.
	FmDeskNum param.Opt[string] `json:"fmDeskNum,omitzero"`
	// Last name of the flight manager assigned to the sortie. Null when no flight
	// manager is assigned.
	FmName param.Opt[string] `json:"fmName,omitzero"`
	// Mass of fuel required for this leg of the sortie, in kilograms.
	FuelReq param.Opt[float64] `json:"fuelReq,omitzero"`
	// Scheduled ground time, in minutes.
	GndTime param.Opt[float64] `json:"gndTime,omitzero"`
	// Unique identifier of the aircraft.
	IDAircraft param.Opt[string] `json:"idAircraft,omitzero"`
	// The unique identifier of the mission to which this sortie is assigned.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// Joint Chiefs of Staff priority of this sortie.
	JcsPriority param.Opt[string] `json:"jcsPriority,omitzero"`
	// The leg number of this sortie.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The external system line number of this sortie.
	LineNumber param.Opt[int64] `json:"lineNumber,omitzero"`
	// The mission ID according to the source system.
	MissionID param.Opt[string] `json:"missionId,omitzero"`
	// Time the associated mission data was last updated in relation to the aircraft
	// assignment, in ISO 8601 UTC format with millisecond precision. If this time is
	// coming from an external system, it may not sync with the latest mission time
	// associated to this record.
	MissionUpdate param.Opt[time.Time] `json:"missionUpdate,omitzero" format:"date-time"`
	// Remarks concerning the sortie objective.
	ObjectiveRemarks param.Opt[string] `json:"objectiveRemarks,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The sortie identifier provided by the originating source.
	OrigSortieID param.Opt[string] `json:"origSortieId,omitzero"`
	// Liquid oxygen onboard the aircraft for the crew compartment, in liters.
	OxyOnCrew param.Opt[float64] `json:"oxyOnCrew,omitzero"`
	// Liquid oxygen onboard the aircraft for the troop compartment, in liters.
	OxyOnPax param.Opt[float64] `json:"oxyOnPax,omitzero"`
	// Liquid oxygen required on the aircraft for the crew compartment, in liters.
	OxyReqCrew param.Opt[float64] `json:"oxyReqCrew,omitzero"`
	// Liquid oxygen required on the aircraft for the troop compartment, in liters.
	OxyReqPax param.Opt[float64] `json:"oxyReqPax,omitzero"`
	// The POI parking location.
	ParkingLoc param.Opt[string] `json:"parkingLoc,omitzero"`
	// The number of passengers tasked for this sortie.
	Passengers param.Opt[int64] `json:"passengers,omitzero"`
	// The scheduled time that the Aircraft sortie is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	PlannedArrTime param.Opt[time.Time] `json:"plannedArrTime,omitzero" format:"date-time"`
	// The planned primary Standard Conventional Load of the aircraft for this sortie.
	PrimaryScl param.Opt[string] `json:"primarySCL,omitzero"`
	// Aircraft configuration required for the mission.
	ReqConfig param.Opt[string] `json:"reqConfig,omitzero"`
	// Remarks concerning the results of this sortie.
	ResultRemarks param.Opt[string] `json:"resultRemarks,omitzero"`
	// Remarks concerning the schedule.
	ScheduleRemarks param.Opt[string] `json:"scheduleRemarks,omitzero"`
	// The planned secondary Standard Conventional Load of the aircraft for this
	// sortie.
	SecondaryScl param.Opt[string] `json:"secondarySCL,omitzero"`
	// Indicates the group responsible for recording the completion time of the next
	// event in the sequence of events assigned to this sortie (e.g. OPS - Operations,
	// MX - Maintenance, TR - Transportation, etc.).
	Soe param.Opt[string] `json:"soe,omitzero"`
	// The scheduled UTC date for this sortie, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	SortieDate param.Opt[time.Time] `json:"sortieDate,omitzero" format:"date"`
	// The tail number of the aircraft assigned to this sortie.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// The prior permission required (PPR) status.
	//
	// Any of "NOT REQUIRED", "REQUIRED NOT REQUESTED", "GRANTED", "PENDING".
	PprStatus AircraftSortyUpdateParamsPprStatus `json:"pprStatus,omitzero"`
	// Type of Ravens required for this sortie (N - None, R - Raven (Security Team)
	// required, C6 - Consider ravens (Ground time over 6 hours), R6 - Ravens required
	// (Ground time over 6 hours)).
	//
	// Any of "N", "R", "C6", "R6".
	RvnReq AircraftSortyUpdateParamsRvnReq `json:"rvnReq,omitzero"`
	paramObj
}

func (r AircraftSortyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AircraftSortyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AircraftSortyUpdateParams) UnmarshalJSON(data []byte) error {
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
type AircraftSortyUpdateParamsDataMode string

const (
	AircraftSortyUpdateParamsDataModeReal      AircraftSortyUpdateParamsDataMode = "REAL"
	AircraftSortyUpdateParamsDataModeTest      AircraftSortyUpdateParamsDataMode = "TEST"
	AircraftSortyUpdateParamsDataModeSimulated AircraftSortyUpdateParamsDataMode = "SIMULATED"
	AircraftSortyUpdateParamsDataModeExercise  AircraftSortyUpdateParamsDataMode = "EXERCISE"
)

// The prior permission required (PPR) status.
type AircraftSortyUpdateParamsPprStatus string

const (
	AircraftSortyUpdateParamsPprStatusNotRequired          AircraftSortyUpdateParamsPprStatus = "NOT REQUIRED"
	AircraftSortyUpdateParamsPprStatusRequiredNotRequested AircraftSortyUpdateParamsPprStatus = "REQUIRED NOT REQUESTED"
	AircraftSortyUpdateParamsPprStatusGranted              AircraftSortyUpdateParamsPprStatus = "GRANTED"
	AircraftSortyUpdateParamsPprStatusPending              AircraftSortyUpdateParamsPprStatus = "PENDING"
)

// Type of Ravens required for this sortie (N - None, R - Raven (Security Team)
// required, C6 - Consider ravens (Ground time over 6 hours), R6 - Ravens required
// (Ground time over 6 hours)).
type AircraftSortyUpdateParamsRvnReq string

const (
	AircraftSortyUpdateParamsRvnReqN  AircraftSortyUpdateParamsRvnReq = "N"
	AircraftSortyUpdateParamsRvnReqR  AircraftSortyUpdateParamsRvnReq = "R"
	AircraftSortyUpdateParamsRvnReqC6 AircraftSortyUpdateParamsRvnReq = "C6"
	AircraftSortyUpdateParamsRvnReqR6 AircraftSortyUpdateParamsRvnReq = "R6"
)

type AircraftSortyTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The scheduled time that the Aircraft sortie is planned to depart, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	PlannedDepTime time.Time        `query:"plannedDepTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftSortyTupleParams]'s query parameters as
// `url.Values`.
func (r AircraftSortyTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
