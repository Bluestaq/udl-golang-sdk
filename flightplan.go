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

// FlightplanService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFlightplanService] method instead.
type FlightplanService struct {
	Options []option.RequestOption
}

// NewFlightplanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFlightplanService(opts ...option.RequestOption) (r FlightplanService) {
	r = FlightplanService{}
	r.Options = opts
	return
}

// Service operation to take a single FlightPlan object as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *FlightplanService) New(ctx context.Context, body FlightplanNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/flightplan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single FlightPlan record by its unique ID passed as a
// path parameter.
func (r *FlightplanService) Get(ctx context.Context, id string, query FlightplanGetParams, opts ...option.RequestOption) (res *shared.FlightPlanFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/flightplan/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single flightplan record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *FlightplanService) Update(ctx context.Context, id string, body FlightplanUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/flightplan/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *FlightplanService) List(ctx context.Context, query FlightplanListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[FlightPlanAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/flightplan"
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
func (r *FlightplanService) ListAutoPaging(ctx context.Context, query FlightplanListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[FlightPlanAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a flight plan record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *FlightplanService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/flightplan/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *FlightplanService) Count(ctx context.Context, query FlightplanCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/flightplan/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *FlightplanService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *FlightplanQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/flightplan/queryhelp"
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
func (r *FlightplanService) Tuple(ctx context.Context, query FlightplanTupleParams, opts ...option.RequestOption) (res *[]shared.FlightPlanFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/flightplan/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take one or many flight plan records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *FlightplanService) UnvalidatedPublish(ctx context.Context, body FlightplanUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-flightplan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Flight Plan contains data specifying the details of an intended flight including
// schedule and expected route.
type FlightPlanAbridged struct {
	// The airfield identifier of the arrival location, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAirfield string `json:"arrAirfield,required"`
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
	DataMode FlightPlanAbridgedDataMode `json:"dataMode,required"`
	// The airfield identifier of the departure location, International Civil Aviation
	// Organization (ICAO) code preferred.
	DepAirfield string `json:"depAirfield,required"`
	// The generation time of this flight plan in ISO 8601 UTC format, with millisecond
	// precision.
	GenTs time.Time `json:"genTS,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this flight plan.
	// Intended as, but not constrained to, MIL-STD-6016 environment dependent specific
	// type designations.
	AircraftMds string `json:"aircraftMDS"`
	// Collection of air refueling events occurring on this flight.
	AirRefuelEvents []FlightPlanAbridgedAirRefuelEvent `json:"airRefuelEvents"`
	// Air Mobility Command (AMC) mission identifier according to Mobility Air Forces
	// (MAF) encode/decode procedures.
	AmcMissionID string `json:"amcMissionId"`
	// Fuel burned from the initial approach point to landing in pounds.
	AppLandingFuel float64 `json:"appLandingFuel"`
	// The first designated alternate arrival airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAlternate1 string `json:"arrAlternate1"`
	// Fuel required to fly to alternate landing site 1 and land in pounds.
	ArrAlternate1Fuel float64 `json:"arrAlternate1Fuel"`
	// The second designated alternate arrival airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAlternate2 string `json:"arrAlternate2"`
	// Fuel required to fly to alternate landing site 2 and land in pounds.
	ArrAlternate2Fuel float64 `json:"arrAlternate2Fuel"`
	// Additional fuel burned at landing/missed approach for icing during arrival in
	// pounds.
	ArrIceFuel float64 `json:"arrIceFuel"`
	// The arrival runway for this flight.
	ArrRunway string `json:"arrRunway"`
	// Array of Air Traffic Control (ATC) addresses.
	AtcAddresses []string `json:"atcAddresses"`
	// Average temperature deviation of the primary, divert, and alternate path for the
	// route between first Top of Climb and last Top of Descent in degrees Celsius.
	AvgTempDev float64 `json:"avgTempDev"`
	// Fuel planned to be burned during the flight in pounds.
	BurnedFuel float64 `json:"burnedFuel"`
	// The call sign assigned to the aircraft for this flight plan.
	CallSign string `json:"callSign"`
	// Remarks about the planned cargo associated with this flight plan.
	CargoRemark string `json:"cargoRemark"`
	// Fuel required from brake release to Top of Climb in pounds.
	ClimbFuel float64 `json:"climbFuel"`
	// Time required from brake release to Top of Climb expressed as HH:MM.
	ClimbTime string `json:"climbTime"`
	// The amount of contingency fuel in pounds.
	ContingencyFuel float64 `json:"contingencyFuel"`
	// Array of country codes for the countries overflown during this flight in ISO
	// 3166-1 Alpha-2 format.
	CountryCodes []string `json:"countryCodes"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The designated alternate departure airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	DepAlternate string `json:"depAlternate"`
	// The depressurization fuel required to fly from the Equal Time Point to the Last
	// Suitable/First Suitable airfield at depressurization altitude in pounds.
	DepressFuel float64 `json:"depressFuel"`
	// The departure runway for this flight.
	DepRunway string `json:"depRunway"`
	// The percent degrade due to drag for this aircraft.
	DragIndex float64 `json:"dragIndex"`
	// Additional fuel burned at landing/missed approach for an early descent in
	// pounds.
	EarlyDescentFuel float64 `json:"earlyDescentFuel"`
	// Total endurance time based on the fuel on board expressed as HH:MM.
	EnduranceTime string `json:"enduranceTime"`
	// Fuel required to fly from Top of Climb to Top of Descent in pounds.
	EnrouteFuel float64 `json:"enrouteFuel"`
	// Time required to fly from Top of Climb to Top of Descent expressed as HH:MM.
	EnrouteTime string `json:"enrouteTime"`
	// The list of equipment on the aircraft as defined in the Flight Information
	// Publications (FLIP) General Planning (GP) manual.
	Equipment string `json:"equipment"`
	// The estimated time of departure for the aircraft, in ISO 8601 UTC format, with
	// millisecond precision.
	EstDepTime time.Time `json:"estDepTime" format:"date-time"`
	// Array of Extended Operations (ETOPS) adequate landing airfields that are within
	// the mission region.
	EtopsAirfields []string `json:"etopsAirfields"`
	// Array of Extended Operations (ETOPS) alternate suitable landing airfields that
	// are within the mission region.
	EtopsAltAirfields []string `json:"etopsAltAirfields"`
	// The Extended Operations (ETOPS) rating used to calculate this flight plan.
	EtopsRating string `json:"etopsRating"`
	// The Extended Operations (ETOPS) validity window for the alternate airfield.
	EtopsValWindow string `json:"etopsValWindow"`
	// The source ID of the flight plan from the generating system.
	ExternalID string `json:"externalId"`
	// Collection of messages associated with this flight plan indicating the severity,
	// the point where the message was generated, the path (Primary, Alternate, etc.),
	// and the text of the message.
	FlightPlanMessages []FlightPlanAbridgedFlightPlanMessage `json:"flightPlanMessages"`
	// Collection of point groups generated for this flight plan. Groups include point
	// sets for Extended Operations (ETOPS), Critical Fuel Point, and Equal Time Point
	// (ETP).
	FlightPlanPointGroups []FlightPlanAbridgedFlightPlanPointGroup `json:"flightPlanPointGroups"`
	// Collection of waypoints associated with this flight plan.
	FlightPlanWaypoints []FlightPlanAbridgedFlightPlanWaypoint `json:"flightPlanWaypoints"`
	// The flight rules this flight plan is being filed under.
	FlightRules string `json:"flightRules"`
	// The type of flight (MILITARY, CIVILIAN, etc).
	FlightType string `json:"flightType"`
	// The fuel degrade percentage used for this mission.
	FuelDegrade float64 `json:"fuelDegrade"`
	// The GPS Receiver Autonomous Integrity Monitoring (RAIM) message. A RAIM system
	// assesses the integrity of the GPS signals. This system predicts outages for a
	// specified geographical area. These predictions are based on the location, path,
	// and scheduled GPS satellite outages.
	GpsRaim string `json:"gpsRAIM"`
	// Additional fuel burned at Top of Climb in pounds.
	HoldDownFuel float64 `json:"holdDownFuel"`
	// Additional fuel burned at the destination for holding in pounds.
	HoldFuel float64 `json:"holdFuel"`
	// Additional time for holding at the destination expressed as HH:MM.
	HoldTime string `json:"holdTime"`
	// The UDL unique identifier of the aircraft associated with this flight plan.
	IDAircraft string `json:"idAircraft"`
	// The UDL unique identifier of the arrival airfield associated with this flight
	// plan.
	IDArrAirfield string `json:"idArrAirfield"`
	// The UDL unique identifier of the departure airfield associated with this flight
	// plan.
	IDDepAirfield string `json:"idDepAirfield"`
	// The amount of identified extra fuel carried and not available in the burn plan
	// in pounds.
	IdentExtraFuel float64 `json:"identExtraFuel"`
	// The UDL unique identifier of the aircraft sortie associated with this flight
	// plan.
	IDSortie string `json:"idSortie"`
	// A character string representation of the initial filed cruise speed for this
	// flight (prepended values of K, N, and M represent kilometers per hour, knots,
	// and Mach, respectively).
	InitialCruiseSpeed string `json:"initialCruiseSpeed"`
	// A character string representation of the initial filed altitude level for this
	// flight (prepended values of F, S, A, and M represent flight level in hundreds of
	// feet, standard metric level in tens of meters, altitude in hundreds of feet, and
	// altitude in tens of meters, respectively).
	InitialFlightLevel string `json:"initialFlightLevel"`
	// Fuel planned to be remaining on the airplane at landing in pounds.
	LandingFuel float64 `json:"landingFuel"`
	// The leg number of this flight plan.
	LegNum int64 `json:"legNum"`
	// The minimum fuel on board required to divert in pounds.
	MinDivertFuel float64 `json:"minDivertFuel"`
	// The mission index value for this mission. The mission index is the ratio of
	// time-related cost of aircraft operation to the cost of fuel.
	MsnIndex float64 `json:"msnIndex"`
	// Additional remarks for air traffic control for this flight.
	Notes string `json:"notes"`
	// The number of aircraft flying this flight plan.
	NumAircraft int64 `json:"numAircraft"`
	// Additional fuel burned at Top of Descent for the operational condition in
	// pounds.
	OpConditionFuel float64 `json:"opConditionFuel"`
	// Operating weight of the aircraft in pounds.
	OpWeight float64 `json:"opWeight"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// Air Traffic Control address filing the flight plan.
	Originator string `json:"originator"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Remarks from the planners concerning this flight plan.
	PlannerRemark string `json:"plannerRemark"`
	// Total of all fuel required to complete the flight in pounds, including fuel to
	// be dispensed on a refueling mission.
	RampFuel float64 `json:"rampFuel"`
	// Total fuel remaining at alternate landing site 1 in pounds.
	RemAlternate1Fuel float64 `json:"remAlternate1Fuel"`
	// Total fuel remaining at alternate landing site 2 in pounds.
	RemAlternate2Fuel float64 `json:"remAlternate2Fuel"`
	// The amount of reserve fuel in pounds.
	ReserveFuel float64 `json:"reserveFuel"`
	// The 1801 fileable route of flight string for this flight. The route of flight
	// string contains route designators, significant points, change of speed/altitude,
	// change of flight rules, and cruise climbs.
	RouteString string `json:"routeString"`
	// Name of the planned Standard Instrument Departure (SID) procedure.
	Sid string `json:"sid"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Name of the planned Standard Terminal Arrival (STAR) procedure.
	Star string `json:"star"`
	// Status of this flight plan (e.g., ACTIVE, APPROVED, PLANNED, etc.).
	Status string `json:"status"`
	// The tail number of the aircraft associated with this flight plan.
	TailNumber string `json:"tailNumber"`
	// Fuel at takeoff, which is calculated as the ramp fuel minus the taxi fuel in
	// pounds.
	TakeoffFuel float64 `json:"takeoffFuel"`
	// Fuel required to start engines and taxi to the end of the runway in pounds.
	TaxiFuel float64 `json:"taxiFuel"`
	// Additional fuel burned at Top of Descent for thunderstorm avoidance in pounds.
	ThunderAvoidFuel float64 `json:"thunderAvoidFuel"`
	// Fuel remaining at Top of Climb in pounds.
	TocFuel float64 `json:"tocFuel"`
	// Additional fuel burned at Top of Climb for icing in pounds.
	TocIceFuel float64 `json:"tocIceFuel"`
	// Fuel remaining at Top of Descent in pounds.
	TodFuel float64 `json:"todFuel"`
	// Additional fuel burned at Top of Descent for icing in pounds.
	TodIceFuel float64 `json:"todIceFuel"`
	// The amount of unidentified extra fuel required to get to min landing in pounds.
	UnidentExtraFuel float64 `json:"unidentExtraFuel"`
	// The amount of unusable fuel in pounds.
	UnusableFuel float64 `json:"unusableFuel"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The wake turbulence category for this flight. The categories are assigned by the
	// International Civil Aviation Organization (ICAO) and are based on maximum
	// certified takeoff mass for the purpose of separating aircraft in flight due to
	// wake turbulence. Valid values include LIGHT, MEDIUM, LARGE, HEAVY, and SUPER.
	WakeTurbCat string `json:"wakeTurbCat"`
	// Wind factor for the first half of the route. This is the average wind factor
	// from first Top of Climb to the mid-time of the entire route in knots. A positive
	// value indicates a headwind, while a negative value indicates a tailwind.
	WindFac1 float64 `json:"windFac1"`
	// Wind factor for the second half of the route. This is the average wind factor
	// from the mid-time of the entire route to last Top of Descent in knots. A
	// positive value indicates a headwind, while a negative value indicates a
	// tailwind.
	WindFac2 float64 `json:"windFac2"`
	// Average wind factor from Top of Climb to Top of Descent in knots. A positive
	// value indicates a headwind, while a negative value indicates a tailwind.
	WindFacAvg float64 `json:"windFacAvg"`
	// The date and time the weather valid period ends in ISO 8601 UTC format, with
	// millisecond precision.
	WxValidEnd time.Time `json:"wxValidEnd" format:"date-time"`
	// The date and time the weather valid period begins in ISO 8601 UTC format, with
	// millisecond precision.
	WxValidStart time.Time `json:"wxValidStart" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArrAirfield           respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		DepAirfield           respjson.Field
		GenTs                 respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AircraftMds           respjson.Field
		AirRefuelEvents       respjson.Field
		AmcMissionID          respjson.Field
		AppLandingFuel        respjson.Field
		ArrAlternate1         respjson.Field
		ArrAlternate1Fuel     respjson.Field
		ArrAlternate2         respjson.Field
		ArrAlternate2Fuel     respjson.Field
		ArrIceFuel            respjson.Field
		ArrRunway             respjson.Field
		AtcAddresses          respjson.Field
		AvgTempDev            respjson.Field
		BurnedFuel            respjson.Field
		CallSign              respjson.Field
		CargoRemark           respjson.Field
		ClimbFuel             respjson.Field
		ClimbTime             respjson.Field
		ContingencyFuel       respjson.Field
		CountryCodes          respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DepAlternate          respjson.Field
		DepressFuel           respjson.Field
		DepRunway             respjson.Field
		DragIndex             respjson.Field
		EarlyDescentFuel      respjson.Field
		EnduranceTime         respjson.Field
		EnrouteFuel           respjson.Field
		EnrouteTime           respjson.Field
		Equipment             respjson.Field
		EstDepTime            respjson.Field
		EtopsAirfields        respjson.Field
		EtopsAltAirfields     respjson.Field
		EtopsRating           respjson.Field
		EtopsValWindow        respjson.Field
		ExternalID            respjson.Field
		FlightPlanMessages    respjson.Field
		FlightPlanPointGroups respjson.Field
		FlightPlanWaypoints   respjson.Field
		FlightRules           respjson.Field
		FlightType            respjson.Field
		FuelDegrade           respjson.Field
		GpsRaim               respjson.Field
		HoldDownFuel          respjson.Field
		HoldFuel              respjson.Field
		HoldTime              respjson.Field
		IDAircraft            respjson.Field
		IDArrAirfield         respjson.Field
		IDDepAirfield         respjson.Field
		IdentExtraFuel        respjson.Field
		IDSortie              respjson.Field
		InitialCruiseSpeed    respjson.Field
		InitialFlightLevel    respjson.Field
		LandingFuel           respjson.Field
		LegNum                respjson.Field
		MinDivertFuel         respjson.Field
		MsnIndex              respjson.Field
		Notes                 respjson.Field
		NumAircraft           respjson.Field
		OpConditionFuel       respjson.Field
		OpWeight              respjson.Field
		Origin                respjson.Field
		Originator            respjson.Field
		OrigNetwork           respjson.Field
		PlannerRemark         respjson.Field
		RampFuel              respjson.Field
		RemAlternate1Fuel     respjson.Field
		RemAlternate2Fuel     respjson.Field
		ReserveFuel           respjson.Field
		RouteString           respjson.Field
		Sid                   respjson.Field
		SourceDl              respjson.Field
		Star                  respjson.Field
		Status                respjson.Field
		TailNumber            respjson.Field
		TakeoffFuel           respjson.Field
		TaxiFuel              respjson.Field
		ThunderAvoidFuel      respjson.Field
		TocFuel               respjson.Field
		TocIceFuel            respjson.Field
		TodFuel               respjson.Field
		TodIceFuel            respjson.Field
		UnidentExtraFuel      respjson.Field
		UnusableFuel          respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		WakeTurbCat           respjson.Field
		WindFac1              respjson.Field
		WindFac2              respjson.Field
		WindFacAvg            respjson.Field
		WxValidEnd            respjson.Field
		WxValidStart          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlightPlanAbridged) RawJSON() string { return r.JSON.raw }
func (r *FlightPlanAbridged) UnmarshalJSON(data []byte) error {
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
type FlightPlanAbridgedDataMode string

const (
	FlightPlanAbridgedDataModeReal      FlightPlanAbridgedDataMode = "REAL"
	FlightPlanAbridgedDataModeTest      FlightPlanAbridgedDataMode = "TEST"
	FlightPlanAbridgedDataModeSimulated FlightPlanAbridgedDataMode = "SIMULATED"
	FlightPlanAbridgedDataModeExercise  FlightPlanAbridgedDataMode = "EXERCISE"
)

// Collection of air refueling events occurring on this flight.
type FlightPlanAbridgedAirRefuelEvent struct {
	// Additional degrade for air refueling, cumulative with fuelDegrade field percent.
	ArDegrade float64 `json:"arDegrade"`
	// Fuel onloaded (use positive numbers) or fuel offloaded (use negative numbers) in
	// pounds.
	ArExchangedFuel float64 `json:"arExchangedFuel"`
	// The number of this air refueling event within the flight plan.
	ArNum int64 `json:"arNum"`
	// Fuel required to fly from air refueling exit point to air refueling divert
	// alternate airfield in pounds.
	DivertFuel float64 `json:"divertFuel"`
	// Fuel remaining at the air refueling exit in pounds.
	ExitFuel float64 `json:"exitFuel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArDegrade       respjson.Field
		ArExchangedFuel respjson.Field
		ArNum           respjson.Field
		DivertFuel      respjson.Field
		ExitFuel        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlightPlanAbridgedAirRefuelEvent) RawJSON() string { return r.JSON.raw }
func (r *FlightPlanAbridgedAirRefuelEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of messages associated with this flight plan indicating the severity,
// the point where the message was generated, the path (Primary, Alternate, etc.),
// and the text of the message.
type FlightPlanAbridgedFlightPlanMessage struct {
	// The text of the message.
	MsgText string `json:"msgText"`
	// The flight path that generated the message (PRIMARY, ALTERNATE, etc.).
	RoutePath string `json:"routePath"`
	// The severity of the message.
	Severity string `json:"severity"`
	// The waypoint number for which the message was generated, or enter "PLAN" for a
	// message impacting the entire route.
	WpNum string `json:"wpNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MsgText     respjson.Field
		RoutePath   respjson.Field
		Severity    respjson.Field
		WpNum       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlightPlanAbridgedFlightPlanMessage) RawJSON() string { return r.JSON.raw }
func (r *FlightPlanAbridgedFlightPlanMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of point groups generated for this flight plan. Groups include point
// sets for Extended Operations (ETOPS), Critical Fuel Point, and Equal Time Point
// (ETP).
type FlightPlanAbridgedFlightPlanPointGroup struct {
	// Average fuel flow at which the fuel was calculated in pounds per hour.
	AvgFuelFlow float64 `json:"avgFuelFlow"`
	// Average wind factor from the Extended Operations (ETOPS) point to the recovery
	// airfield in knots.
	EtopsAvgWindFactor float64 `json:"etopsAvgWindFactor"`
	// Distance from the Extended Operations (ETOPS) point to the recovery airfield in
	// nautical miles.
	EtopsDistance float64 `json:"etopsDistance"`
	// Fuel required to fly from the Extended Operations (ETOPS) point to the recovery
	// airfield in pounds.
	EtopsReqFuel float64 `json:"etopsReqFuel"`
	// Temperature deviation from the Extended Operations (ETOPS) point to the recovery
	// airfield in degrees Celsius.
	EtopsTempDev float64 `json:"etopsTempDev"`
	// Time to fly from the Extended Operations (ETOPS) point to the recovery airfield
	// expressed in HH:MM format.
	EtopsTime string `json:"etopsTime"`
	// Array of point data for this Point Group.
	FlightPlanPoints []FlightPlanAbridgedFlightPlanPointGroupFlightPlanPoint `json:"flightPlanPoints"`
	// Total time from takeoff when the point is reached expressed in HH:MM format.
	FromTakeoffTime string `json:"fromTakeoffTime"`
	// Average wind factor from the Equal Time Point (ETP) to the first suitable
	// airfield in knots.
	FsafAvgWindFactor float64 `json:"fsafAvgWindFactor"`
	// Distance from the Equal Time Point (ETP) to the first suitable airfield in
	// nautical miles.
	FsafDistance float64 `json:"fsafDistance"`
	// Fuel required to fly from the Equal Time Point (ETP) to the first suitable
	// airfield in pounds.
	FsafReqFuel float64 `json:"fsafReqFuel"`
	// Temperature deviation from the Equal Time Point (ETP) to the first suitable
	// airfield in degrees Celsius.
	FsafTempDev float64 `json:"fsafTempDev"`
	// Time to fly from the Equal Time Point (ETP) to the first suitable airfield
	// expressed in HH:MM format.
	FsafTime string `json:"fsafTime"`
	// Flight level of the point at which the fuel was calculated in feet.
	FuelCalcAlt float64 `json:"fuelCalcAlt"`
	// True airspeed at which the fuel was calculated in knots.
	FuelCalcSpd float64 `json:"fuelCalcSpd"`
	// Average wind factor from the Equal Time Point (ETP) to the last suitable
	// airfield in knots.
	LsafAvgWindFactor float64 `json:"lsafAvgWindFactor"`
	// Distance from the Equal Time Point (ETP) to the last suitable airfield in
	// nautical miles.
	LsafDistance float64 `json:"lsafDistance"`
	// Name of the last suitable airfield, International Civil Aviation Organization
	// (ICAO) code preferred.
	LsafName string `json:"lsafName"`
	// Fuel required to fly from the Equal Time Point (ETP) to the last suitable
	// airfield in pounds.
	LsafReqFuel float64 `json:"lsafReqFuel"`
	// Temperature deviation from the Equal Time Point (ETP) to the last suitable
	// airfield in degrees Celsius.
	LsafTempDev float64 `json:"lsafTempDev"`
	// Time to fly from the Equal Time Point (ETP) to the last suitable airfield
	// expressed in HH:MM format.
	LsafTime string `json:"lsafTime"`
	// Amount of planned fuel on board when the point is reached in pounds.
	PlannedFuel float64 `json:"plannedFuel"`
	// Name of the point group, usually Extended Operations (ETOPS), Critical Fuel
	// Point, and Equal Time Point (ETP) sections.
	PointGroupName string `json:"pointGroupName"`
	// Specifies which Point Group case requires the most fuel.
	WorstFuelCase string `json:"worstFuelCase"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgFuelFlow        respjson.Field
		EtopsAvgWindFactor respjson.Field
		EtopsDistance      respjson.Field
		EtopsReqFuel       respjson.Field
		EtopsTempDev       respjson.Field
		EtopsTime          respjson.Field
		FlightPlanPoints   respjson.Field
		FromTakeoffTime    respjson.Field
		FsafAvgWindFactor  respjson.Field
		FsafDistance       respjson.Field
		FsafReqFuel        respjson.Field
		FsafTempDev        respjson.Field
		FsafTime           respjson.Field
		FuelCalcAlt        respjson.Field
		FuelCalcSpd        respjson.Field
		LsafAvgWindFactor  respjson.Field
		LsafDistance       respjson.Field
		LsafName           respjson.Field
		LsafReqFuel        respjson.Field
		LsafTempDev        respjson.Field
		LsafTime           respjson.Field
		PlannedFuel        respjson.Field
		PointGroupName     respjson.Field
		WorstFuelCase      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlightPlanAbridgedFlightPlanPointGroup) RawJSON() string { return r.JSON.raw }
func (r *FlightPlanAbridgedFlightPlanPointGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Array of point data for this Point Group.
type FlightPlanAbridgedFlightPlanPointGroupFlightPlanPoint struct {
	// Estimated Time of Arrival (ETA) at this point in ISO 8601 UTC format, with
	// millisecond precision.
	FppEta time.Time `json:"fppEta" format:"date-time"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	FppLat float64 `json:"fppLat"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	FppLon float64 `json:"fppLon"`
	// Fuel required at this point to execute an Equal Time Point (ETP) or Extended
	// Operations (ETOPS) plan in pounds.
	FppReqFuel float64 `json:"fppReqFuel"`
	// Name of this point.
	PointName string `json:"pointName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FppEta      respjson.Field
		FppLat      respjson.Field
		FppLon      respjson.Field
		FppReqFuel  respjson.Field
		PointName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlightPlanAbridgedFlightPlanPointGroupFlightPlanPoint) RawJSON() string { return r.JSON.raw }
func (r *FlightPlanAbridgedFlightPlanPointGroupFlightPlanPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of waypoints associated with this flight plan.
type FlightPlanAbridgedFlightPlanWaypoint struct {
	// Points are designated by type as either a comment point or a waypoint. A comment
	// point conveys important information about the point for pilots but is not
	// entered into a flight management system. A waypoint is a point that is entered
	// into a flight management system and/or filed with Air Traffic Control.
	Type string `json:"type,required"`
	// Name of the point. The name of a comment point identifies important information
	// about that point, e.g. Top of Climb. The name of a waypoint identifies the
	// location of that point.
	WaypointName string `json:"waypointName,required"`
	// The air-to-air Tactical Air Navigation (TACAN) channels used by the
	// receiver/tanker during air refueling.
	AaTacanChannel string `json:"aaTacanChannel"`
	// The air distance of this leg in nautical miles.
	AirDistance float64 `json:"airDistance"`
	// The flight path flown for this leg.
	Airway string `json:"airway"`
	// Altitude of a level, point, or object measured in feet above mean sea level.
	Alt float64 `json:"alt"`
	// The ID of the air refueling track/anchor or fixed track.
	ArID string `json:"arId"`
	// Point identifying an air refueling track/anchor or fixed track.
	Arpt string `json:"arpt"`
	// Actual Time of Arrival (ATA) at this waypoint in ISO 8601 UTC format, with
	// millisecond precision.
	Ata time.Time `json:"ata" format:"date-time"`
	// The average calibrated airspeed (CAS) for this leg in knots.
	AvgCalAirspeed float64 `json:"avgCalAirspeed"`
	// The average drift angle for this leg in degrees from true north.
	AvgDriftAng float64 `json:"avgDriftAng"`
	// The average ground speed for this leg in knots.
	AvgGroundSpeed float64 `json:"avgGroundSpeed"`
	// The average true airspeed (TAS) for this leg in knots.
	AvgTrueAirspeed float64 `json:"avgTrueAirspeed"`
	// The average wind direction for this leg in degrees from true north.
	AvgWindDir float64 `json:"avgWindDir"`
	// The average wind speed for this leg in knots.
	AvgWindSpeed float64 `json:"avgWindSpeed"`
	// The day low level altitude in feet above sea level for the leg ending at this
	// waypoint.
	DayLowAlt float64 `json:"dayLowAlt"`
	// Estimated Time of Arrival (ETA) at this waypoint in ISO 8601 UTC format, with
	// millisecond precision.
	Eta time.Time `json:"eta" format:"date-time"`
	// The amount of fuel onloaded or offloaded at this waypoint in pounds (negative
	// value for offload).
	ExchangedFuel float64 `json:"exchangedFuel"`
	// The leg fuel flow in pounds per hour.
	FuelFlow float64 `json:"fuelFlow"`
	// The icing intensity classification for this flight (LIGHT, MODERATE, etc).
	IceCat string `json:"iceCat"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	Lat float64 `json:"lat"`
	// The planned alternate leg based on user-defined constraints, International Civil
	// Aviation Organization (ICAO) code preferred.
	LegAlternate string `json:"legAlternate"`
	// The percent degrade due to drag for this aircraft for this leg.
	LegDragIndex float64 `json:"legDragIndex"`
	// The fuel degrade percentage used for this leg.
	LegFuelDegrade float64 `json:"legFuelDegrade"`
	// The average Mach speed for this leg.
	LegMach float64 `json:"legMach"`
	// The mission index value for this leg. The mission index is the ratio of
	// time-related cost of aircraft operation to the cost of fuel.
	LegMsnIndex float64 `json:"legMsnIndex"`
	// The wind factor for this leg in knots. A positive value indicates a headwind,
	// while a negative value indicates a tailwind.
	LegWindFac float64 `json:"legWindFac"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The magnetic course at leg midpoint in degrees from true north.
	MagCourse float64 `json:"magCourse"`
	// The magnetic heading at leg midpoint in degrees from true north.
	MagHeading float64 `json:"magHeading"`
	// The magnetic variation for this leg in degrees.
	MagVar float64 `json:"magVar"`
	// Navigational Aid (NAVAID) identification code.
	Navaid string `json:"navaid"`
	// The night low level altitude in feet above sea level for the leg ending at this
	// waypoint.
	NightLowAlt float64 `json:"nightLowAlt"`
	// The night vision goggle low level altitude in feet above sea level for the leg
	// ending at this waypoint.
	NvgLowAlt float64 `json:"nvgLowAlt"`
	// The wind direction at this specific point in degrees from true north.
	PointWindDir float64 `json:"pointWindDir"`
	// The wind velocity at this specific point in knots.
	PointWindSpeed float64 `json:"pointWindSpeed"`
	// The primary UHF radio frequency used for the air refueling track or anchor in
	// megahertz.
	PriFreq float64 `json:"priFreq"`
	// The secondary UHF radio frequency used for the air refueling track or anchor in
	// megahertz.
	SecFreq float64 `json:"secFreq"`
	// Tactical Air Navigation (TACAN) channel for the Navigational Aid (NAVAID).
	TacanChannel string `json:"tacanChannel"`
	// Average temperature deviation from standard day profile for this leg in degrees
	// Celsius.
	TempDev float64 `json:"tempDev"`
	// The thunderstorm intensity classification for this flight (LIGHT, MODERATE,
	// etc).
	ThunderCat string `json:"thunderCat"`
	// The total air distance to this waypoint in nautical miles.
	TotalAirDistance float64 `json:"totalAirDistance"`
	// The total distance flown to this waypoint calculated from point of departure in
	// nautical miles.
	TotalFlownDistance float64 `json:"totalFlownDistance"`
	// The total distance remaining from this waypoint to the point of arrival in
	// nautical miles.
	TotalRemDistance float64 `json:"totalRemDistance"`
	// The total fuel remaining at this waypoint in pounds.
	TotalRemFuel float64 `json:"totalRemFuel"`
	// The total time accumulated from takeoff to this waypoint expressed as HH:MM.
	TotalTime string `json:"totalTime"`
	// The total time remaining from this waypoint to the point of arrival expressed as
	// HH:MM.
	TotalTimeRem string `json:"totalTimeRem"`
	// The total fuel used to this waypoint from point of departure in pounds.
	TotalUsedFuel float64 `json:"totalUsedFuel"`
	// The total weight of the aircraft at this waypoint in pounds.
	TotalWeight float64 `json:"totalWeight"`
	// The true course at leg midpoint in degrees from true north.
	TrueCourse float64 `json:"trueCourse"`
	// The turbulence intensity classification for this flight (LIGHT, MODERATE, etc).
	TurbCat string `json:"turbCat"`
	// VHF Omni-directional Range (VOR) frequency for the Navigational Aid (NAVAID) in
	// megahertz.
	VorFreq float64 `json:"vorFreq"`
	// The waypoint number on the route. Comment points do not get a waypoint number.
	WaypointNum int64 `json:"waypointNum"`
	// The zone/leg distance flown in nautical miles.
	ZoneDistance float64 `json:"zoneDistance"`
	// The amount of fuel used on this zone/leg in pounds.
	ZoneFuel float64 `json:"zoneFuel"`
	// The time to fly this zone/leg in minutes.
	ZoneTime float64 `json:"zoneTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type               respjson.Field
		WaypointName       respjson.Field
		AaTacanChannel     respjson.Field
		AirDistance        respjson.Field
		Airway             respjson.Field
		Alt                respjson.Field
		ArID               respjson.Field
		Arpt               respjson.Field
		Ata                respjson.Field
		AvgCalAirspeed     respjson.Field
		AvgDriftAng        respjson.Field
		AvgGroundSpeed     respjson.Field
		AvgTrueAirspeed    respjson.Field
		AvgWindDir         respjson.Field
		AvgWindSpeed       respjson.Field
		DayLowAlt          respjson.Field
		Eta                respjson.Field
		ExchangedFuel      respjson.Field
		FuelFlow           respjson.Field
		IceCat             respjson.Field
		Lat                respjson.Field
		LegAlternate       respjson.Field
		LegDragIndex       respjson.Field
		LegFuelDegrade     respjson.Field
		LegMach            respjson.Field
		LegMsnIndex        respjson.Field
		LegWindFac         respjson.Field
		Lon                respjson.Field
		MagCourse          respjson.Field
		MagHeading         respjson.Field
		MagVar             respjson.Field
		Navaid             respjson.Field
		NightLowAlt        respjson.Field
		NvgLowAlt          respjson.Field
		PointWindDir       respjson.Field
		PointWindSpeed     respjson.Field
		PriFreq            respjson.Field
		SecFreq            respjson.Field
		TacanChannel       respjson.Field
		TempDev            respjson.Field
		ThunderCat         respjson.Field
		TotalAirDistance   respjson.Field
		TotalFlownDistance respjson.Field
		TotalRemDistance   respjson.Field
		TotalRemFuel       respjson.Field
		TotalTime          respjson.Field
		TotalTimeRem       respjson.Field
		TotalUsedFuel      respjson.Field
		TotalWeight        respjson.Field
		TrueCourse         respjson.Field
		TurbCat            respjson.Field
		VorFreq            respjson.Field
		WaypointNum        respjson.Field
		ZoneDistance       respjson.Field
		ZoneFuel           respjson.Field
		ZoneTime           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlightPlanAbridgedFlightPlanWaypoint) RawJSON() string { return r.JSON.raw }
func (r *FlightPlanAbridgedFlightPlanWaypoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlightplanQueryhelpResponse struct {
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
func (r FlightplanQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *FlightplanQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlightplanNewParams struct {
	// The airfield identifier of the arrival location, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAirfield string `json:"arrAirfield,required"`
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
	DataMode FlightplanNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The airfield identifier of the departure location, International Civil Aviation
	// Organization (ICAO) code preferred.
	DepAirfield string `json:"depAirfield,required"`
	// The generation time of this flight plan in ISO 8601 UTC format, with millisecond
	// precision.
	GenTs time.Time `json:"genTS,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this flight plan.
	// Intended as, but not constrained to, MIL-STD-6016 environment dependent specific
	// type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Air Mobility Command (AMC) mission identifier according to Mobility Air Forces
	// (MAF) encode/decode procedures.
	AmcMissionID param.Opt[string] `json:"amcMissionId,omitzero"`
	// Fuel burned from the initial approach point to landing in pounds.
	AppLandingFuel param.Opt[float64] `json:"appLandingFuel,omitzero"`
	// The first designated alternate arrival airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAlternate1 param.Opt[string] `json:"arrAlternate1,omitzero"`
	// Fuel required to fly to alternate landing site 1 and land in pounds.
	ArrAlternate1Fuel param.Opt[float64] `json:"arrAlternate1Fuel,omitzero"`
	// The second designated alternate arrival airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAlternate2 param.Opt[string] `json:"arrAlternate2,omitzero"`
	// Fuel required to fly to alternate landing site 2 and land in pounds.
	ArrAlternate2Fuel param.Opt[float64] `json:"arrAlternate2Fuel,omitzero"`
	// Additional fuel burned at landing/missed approach for icing during arrival in
	// pounds.
	ArrIceFuel param.Opt[float64] `json:"arrIceFuel,omitzero"`
	// The arrival runway for this flight.
	ArrRunway param.Opt[string] `json:"arrRunway,omitzero"`
	// Average temperature deviation of the primary, divert, and alternate path for the
	// route between first Top of Climb and last Top of Descent in degrees Celsius.
	AvgTempDev param.Opt[float64] `json:"avgTempDev,omitzero"`
	// Fuel planned to be burned during the flight in pounds.
	BurnedFuel param.Opt[float64] `json:"burnedFuel,omitzero"`
	// The call sign assigned to the aircraft for this flight plan.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Remarks about the planned cargo associated with this flight plan.
	CargoRemark param.Opt[string] `json:"cargoRemark,omitzero"`
	// Fuel required from brake release to Top of Climb in pounds.
	ClimbFuel param.Opt[float64] `json:"climbFuel,omitzero"`
	// Time required from brake release to Top of Climb expressed as HH:MM.
	ClimbTime param.Opt[string] `json:"climbTime,omitzero"`
	// The amount of contingency fuel in pounds.
	ContingencyFuel param.Opt[float64] `json:"contingencyFuel,omitzero"`
	// The designated alternate departure airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	DepAlternate param.Opt[string] `json:"depAlternate,omitzero"`
	// The depressurization fuel required to fly from the Equal Time Point to the Last
	// Suitable/First Suitable airfield at depressurization altitude in pounds.
	DepressFuel param.Opt[float64] `json:"depressFuel,omitzero"`
	// The departure runway for this flight.
	DepRunway param.Opt[string] `json:"depRunway,omitzero"`
	// The percent degrade due to drag for this aircraft.
	DragIndex param.Opt[float64] `json:"dragIndex,omitzero"`
	// Additional fuel burned at landing/missed approach for an early descent in
	// pounds.
	EarlyDescentFuel param.Opt[float64] `json:"earlyDescentFuel,omitzero"`
	// Total endurance time based on the fuel on board expressed as HH:MM.
	EnduranceTime param.Opt[string] `json:"enduranceTime,omitzero"`
	// Fuel required to fly from Top of Climb to Top of Descent in pounds.
	EnrouteFuel param.Opt[float64] `json:"enrouteFuel,omitzero"`
	// Time required to fly from Top of Climb to Top of Descent expressed as HH:MM.
	EnrouteTime param.Opt[string] `json:"enrouteTime,omitzero"`
	// The list of equipment on the aircraft as defined in the Flight Information
	// Publications (FLIP) General Planning (GP) manual.
	Equipment param.Opt[string] `json:"equipment,omitzero"`
	// The estimated time of departure for the aircraft, in ISO 8601 UTC format, with
	// millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// The Extended Operations (ETOPS) rating used to calculate this flight plan.
	EtopsRating param.Opt[string] `json:"etopsRating,omitzero"`
	// The Extended Operations (ETOPS) validity window for the alternate airfield.
	EtopsValWindow param.Opt[string] `json:"etopsValWindow,omitzero"`
	// The source ID of the flight plan from the generating system.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The flight rules this flight plan is being filed under.
	FlightRules param.Opt[string] `json:"flightRules,omitzero"`
	// The type of flight (MILITARY, CIVILIAN, etc).
	FlightType param.Opt[string] `json:"flightType,omitzero"`
	// The fuel degrade percentage used for this mission.
	FuelDegrade param.Opt[float64] `json:"fuelDegrade,omitzero"`
	// The GPS Receiver Autonomous Integrity Monitoring (RAIM) message. A RAIM system
	// assesses the integrity of the GPS signals. This system predicts outages for a
	// specified geographical area. These predictions are based on the location, path,
	// and scheduled GPS satellite outages.
	GpsRaim param.Opt[string] `json:"gpsRAIM,omitzero"`
	// Additional fuel burned at Top of Climb in pounds.
	HoldDownFuel param.Opt[float64] `json:"holdDownFuel,omitzero"`
	// Additional fuel burned at the destination for holding in pounds.
	HoldFuel param.Opt[float64] `json:"holdFuel,omitzero"`
	// Additional time for holding at the destination expressed as HH:MM.
	HoldTime param.Opt[string] `json:"holdTime,omitzero"`
	// The UDL unique identifier of the aircraft associated with this flight plan.
	IDAircraft param.Opt[string] `json:"idAircraft,omitzero"`
	// The UDL unique identifier of the arrival airfield associated with this flight
	// plan.
	IDArrAirfield param.Opt[string] `json:"idArrAirfield,omitzero"`
	// The UDL unique identifier of the departure airfield associated with this flight
	// plan.
	IDDepAirfield param.Opt[string] `json:"idDepAirfield,omitzero"`
	// The amount of identified extra fuel carried and not available in the burn plan
	// in pounds.
	IdentExtraFuel param.Opt[float64] `json:"identExtraFuel,omitzero"`
	// The UDL unique identifier of the aircraft sortie associated with this flight
	// plan.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// A character string representation of the initial filed cruise speed for this
	// flight (prepended values of K, N, and M represent kilometers per hour, knots,
	// and Mach, respectively).
	InitialCruiseSpeed param.Opt[string] `json:"initialCruiseSpeed,omitzero"`
	// A character string representation of the initial filed altitude level for this
	// flight (prepended values of F, S, A, and M represent flight level in hundreds of
	// feet, standard metric level in tens of meters, altitude in hundreds of feet, and
	// altitude in tens of meters, respectively).
	InitialFlightLevel param.Opt[string] `json:"initialFlightLevel,omitzero"`
	// Fuel planned to be remaining on the airplane at landing in pounds.
	LandingFuel param.Opt[float64] `json:"landingFuel,omitzero"`
	// The leg number of this flight plan.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The minimum fuel on board required to divert in pounds.
	MinDivertFuel param.Opt[float64] `json:"minDivertFuel,omitzero"`
	// The mission index value for this mission. The mission index is the ratio of
	// time-related cost of aircraft operation to the cost of fuel.
	MsnIndex param.Opt[float64] `json:"msnIndex,omitzero"`
	// Additional remarks for air traffic control for this flight.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// The number of aircraft flying this flight plan.
	NumAircraft param.Opt[int64] `json:"numAircraft,omitzero"`
	// Additional fuel burned at Top of Descent for the operational condition in
	// pounds.
	OpConditionFuel param.Opt[float64] `json:"opConditionFuel,omitzero"`
	// Operating weight of the aircraft in pounds.
	OpWeight param.Opt[float64] `json:"opWeight,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Air Traffic Control address filing the flight plan.
	Originator param.Opt[string] `json:"originator,omitzero"`
	// Remarks from the planners concerning this flight plan.
	PlannerRemark param.Opt[string] `json:"plannerRemark,omitzero"`
	// Total of all fuel required to complete the flight in pounds, including fuel to
	// be dispensed on a refueling mission.
	RampFuel param.Opt[float64] `json:"rampFuel,omitzero"`
	// Total fuel remaining at alternate landing site 1 in pounds.
	RemAlternate1Fuel param.Opt[float64] `json:"remAlternate1Fuel,omitzero"`
	// Total fuel remaining at alternate landing site 2 in pounds.
	RemAlternate2Fuel param.Opt[float64] `json:"remAlternate2Fuel,omitzero"`
	// The amount of reserve fuel in pounds.
	ReserveFuel param.Opt[float64] `json:"reserveFuel,omitzero"`
	// The 1801 fileable route of flight string for this flight. The route of flight
	// string contains route designators, significant points, change of speed/altitude,
	// change of flight rules, and cruise climbs.
	RouteString param.Opt[string] `json:"routeString,omitzero"`
	// Name of the planned Standard Instrument Departure (SID) procedure.
	Sid param.Opt[string] `json:"sid,omitzero"`
	// Name of the planned Standard Terminal Arrival (STAR) procedure.
	Star param.Opt[string] `json:"star,omitzero"`
	// Status of this flight plan (e.g., ACTIVE, APPROVED, PLANNED, etc.).
	Status param.Opt[string] `json:"status,omitzero"`
	// The tail number of the aircraft associated with this flight plan.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// Fuel at takeoff, which is calculated as the ramp fuel minus the taxi fuel in
	// pounds.
	TakeoffFuel param.Opt[float64] `json:"takeoffFuel,omitzero"`
	// Fuel required to start engines and taxi to the end of the runway in pounds.
	TaxiFuel param.Opt[float64] `json:"taxiFuel,omitzero"`
	// Additional fuel burned at Top of Descent for thunderstorm avoidance in pounds.
	ThunderAvoidFuel param.Opt[float64] `json:"thunderAvoidFuel,omitzero"`
	// Fuel remaining at Top of Climb in pounds.
	TocFuel param.Opt[float64] `json:"tocFuel,omitzero"`
	// Additional fuel burned at Top of Climb for icing in pounds.
	TocIceFuel param.Opt[float64] `json:"tocIceFuel,omitzero"`
	// Fuel remaining at Top of Descent in pounds.
	TodFuel param.Opt[float64] `json:"todFuel,omitzero"`
	// Additional fuel burned at Top of Descent for icing in pounds.
	TodIceFuel param.Opt[float64] `json:"todIceFuel,omitzero"`
	// The amount of unidentified extra fuel required to get to min landing in pounds.
	UnidentExtraFuel param.Opt[float64] `json:"unidentExtraFuel,omitzero"`
	// The amount of unusable fuel in pounds.
	UnusableFuel param.Opt[float64] `json:"unusableFuel,omitzero"`
	// The wake turbulence category for this flight. The categories are assigned by the
	// International Civil Aviation Organization (ICAO) and are based on maximum
	// certified takeoff mass for the purpose of separating aircraft in flight due to
	// wake turbulence. Valid values include LIGHT, MEDIUM, LARGE, HEAVY, and SUPER.
	WakeTurbCat param.Opt[string] `json:"wakeTurbCat,omitzero"`
	// Wind factor for the first half of the route. This is the average wind factor
	// from first Top of Climb to the mid-time of the entire route in knots. A positive
	// value indicates a headwind, while a negative value indicates a tailwind.
	WindFac1 param.Opt[float64] `json:"windFac1,omitzero"`
	// Wind factor for the second half of the route. This is the average wind factor
	// from the mid-time of the entire route to last Top of Descent in knots. A
	// positive value indicates a headwind, while a negative value indicates a
	// tailwind.
	WindFac2 param.Opt[float64] `json:"windFac2,omitzero"`
	// Average wind factor from Top of Climb to Top of Descent in knots. A positive
	// value indicates a headwind, while a negative value indicates a tailwind.
	WindFacAvg param.Opt[float64] `json:"windFacAvg,omitzero"`
	// The date and time the weather valid period ends in ISO 8601 UTC format, with
	// millisecond precision.
	WxValidEnd param.Opt[time.Time] `json:"wxValidEnd,omitzero" format:"date-time"`
	// The date and time the weather valid period begins in ISO 8601 UTC format, with
	// millisecond precision.
	WxValidStart param.Opt[time.Time] `json:"wxValidStart,omitzero" format:"date-time"`
	// Collection of air refueling events occurring on this flight.
	AirRefuelEvents []FlightplanNewParamsAirRefuelEvent `json:"airRefuelEvents,omitzero"`
	// Array of Air Traffic Control (ATC) addresses.
	AtcAddresses []string `json:"atcAddresses,omitzero"`
	// Array of country codes for the countries overflown during this flight in ISO
	// 3166-1 Alpha-2 format.
	CountryCodes []string `json:"countryCodes,omitzero"`
	// Array of Extended Operations (ETOPS) adequate landing airfields that are within
	// the mission region.
	EtopsAirfields []string `json:"etopsAirfields,omitzero"`
	// Array of Extended Operations (ETOPS) alternate suitable landing airfields that
	// are within the mission region.
	EtopsAltAirfields []string `json:"etopsAltAirfields,omitzero"`
	// Collection of messages associated with this flight plan indicating the severity,
	// the point where the message was generated, the path (Primary, Alternate, etc.),
	// and the text of the message.
	FlightPlanMessages []FlightplanNewParamsFlightPlanMessage `json:"flightPlanMessages,omitzero"`
	// Collection of point groups generated for this flight plan. Groups include point
	// sets for Extended Operations (ETOPS), Critical Fuel Point, and Equal Time Point
	// (ETP).
	FlightPlanPointGroups []FlightplanNewParamsFlightPlanPointGroup `json:"flightPlanPointGroups,omitzero"`
	// Collection of waypoints associated with this flight plan.
	FlightPlanWaypoints []FlightplanNewParamsFlightPlanWaypoint `json:"flightPlanWaypoints,omitzero"`
	paramObj
}

func (r FlightplanNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanNewParams) UnmarshalJSON(data []byte) error {
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
type FlightplanNewParamsDataMode string

const (
	FlightplanNewParamsDataModeReal      FlightplanNewParamsDataMode = "REAL"
	FlightplanNewParamsDataModeTest      FlightplanNewParamsDataMode = "TEST"
	FlightplanNewParamsDataModeSimulated FlightplanNewParamsDataMode = "SIMULATED"
	FlightplanNewParamsDataModeExercise  FlightplanNewParamsDataMode = "EXERCISE"
)

// Collection of air refueling events occurring on this flight.
type FlightplanNewParamsAirRefuelEvent struct {
	// Additional degrade for air refueling, cumulative with fuelDegrade field percent.
	ArDegrade param.Opt[float64] `json:"arDegrade,omitzero"`
	// Fuel onloaded (use positive numbers) or fuel offloaded (use negative numbers) in
	// pounds.
	ArExchangedFuel param.Opt[float64] `json:"arExchangedFuel,omitzero"`
	// The number of this air refueling event within the flight plan.
	ArNum param.Opt[int64] `json:"arNum,omitzero"`
	// Fuel required to fly from air refueling exit point to air refueling divert
	// alternate airfield in pounds.
	DivertFuel param.Opt[float64] `json:"divertFuel,omitzero"`
	// Fuel remaining at the air refueling exit in pounds.
	ExitFuel param.Opt[float64] `json:"exitFuel,omitzero"`
	paramObj
}

func (r FlightplanNewParamsAirRefuelEvent) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanNewParamsAirRefuelEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanNewParamsAirRefuelEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of messages associated with this flight plan indicating the severity,
// the point where the message was generated, the path (Primary, Alternate, etc.),
// and the text of the message.
type FlightplanNewParamsFlightPlanMessage struct {
	// The text of the message.
	MsgText param.Opt[string] `json:"msgText,omitzero"`
	// The flight path that generated the message (PRIMARY, ALTERNATE, etc.).
	RoutePath param.Opt[string] `json:"routePath,omitzero"`
	// The severity of the message.
	Severity param.Opt[string] `json:"severity,omitzero"`
	// The waypoint number for which the message was generated, or enter "PLAN" for a
	// message impacting the entire route.
	WpNum param.Opt[string] `json:"wpNum,omitzero"`
	paramObj
}

func (r FlightplanNewParamsFlightPlanMessage) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanNewParamsFlightPlanMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanNewParamsFlightPlanMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of point groups generated for this flight plan. Groups include point
// sets for Extended Operations (ETOPS), Critical Fuel Point, and Equal Time Point
// (ETP).
type FlightplanNewParamsFlightPlanPointGroup struct {
	// Average fuel flow at which the fuel was calculated in pounds per hour.
	AvgFuelFlow param.Opt[float64] `json:"avgFuelFlow,omitzero"`
	// Average wind factor from the Extended Operations (ETOPS) point to the recovery
	// airfield in knots.
	EtopsAvgWindFactor param.Opt[float64] `json:"etopsAvgWindFactor,omitzero"`
	// Distance from the Extended Operations (ETOPS) point to the recovery airfield in
	// nautical miles.
	EtopsDistance param.Opt[float64] `json:"etopsDistance,omitzero"`
	// Fuel required to fly from the Extended Operations (ETOPS) point to the recovery
	// airfield in pounds.
	EtopsReqFuel param.Opt[float64] `json:"etopsReqFuel,omitzero"`
	// Temperature deviation from the Extended Operations (ETOPS) point to the recovery
	// airfield in degrees Celsius.
	EtopsTempDev param.Opt[float64] `json:"etopsTempDev,omitzero"`
	// Time to fly from the Extended Operations (ETOPS) point to the recovery airfield
	// expressed in HH:MM format.
	EtopsTime param.Opt[string] `json:"etopsTime,omitzero"`
	// Total time from takeoff when the point is reached expressed in HH:MM format.
	FromTakeoffTime param.Opt[string] `json:"fromTakeoffTime,omitzero"`
	// Average wind factor from the Equal Time Point (ETP) to the first suitable
	// airfield in knots.
	FsafAvgWindFactor param.Opt[float64] `json:"fsafAvgWindFactor,omitzero"`
	// Distance from the Equal Time Point (ETP) to the first suitable airfield in
	// nautical miles.
	FsafDistance param.Opt[float64] `json:"fsafDistance,omitzero"`
	// Fuel required to fly from the Equal Time Point (ETP) to the first suitable
	// airfield in pounds.
	FsafReqFuel param.Opt[float64] `json:"fsafReqFuel,omitzero"`
	// Temperature deviation from the Equal Time Point (ETP) to the first suitable
	// airfield in degrees Celsius.
	FsafTempDev param.Opt[float64] `json:"fsafTempDev,omitzero"`
	// Time to fly from the Equal Time Point (ETP) to the first suitable airfield
	// expressed in HH:MM format.
	FsafTime param.Opt[string] `json:"fsafTime,omitzero"`
	// Flight level of the point at which the fuel was calculated in feet.
	FuelCalcAlt param.Opt[float64] `json:"fuelCalcAlt,omitzero"`
	// True airspeed at which the fuel was calculated in knots.
	FuelCalcSpd param.Opt[float64] `json:"fuelCalcSpd,omitzero"`
	// Average wind factor from the Equal Time Point (ETP) to the last suitable
	// airfield in knots.
	LsafAvgWindFactor param.Opt[float64] `json:"lsafAvgWindFactor,omitzero"`
	// Distance from the Equal Time Point (ETP) to the last suitable airfield in
	// nautical miles.
	LsafDistance param.Opt[float64] `json:"lsafDistance,omitzero"`
	// Name of the last suitable airfield, International Civil Aviation Organization
	// (ICAO) code preferred.
	LsafName param.Opt[string] `json:"lsafName,omitzero"`
	// Fuel required to fly from the Equal Time Point (ETP) to the last suitable
	// airfield in pounds.
	LsafReqFuel param.Opt[float64] `json:"lsafReqFuel,omitzero"`
	// Temperature deviation from the Equal Time Point (ETP) to the last suitable
	// airfield in degrees Celsius.
	LsafTempDev param.Opt[float64] `json:"lsafTempDev,omitzero"`
	// Time to fly from the Equal Time Point (ETP) to the last suitable airfield
	// expressed in HH:MM format.
	LsafTime param.Opt[string] `json:"lsafTime,omitzero"`
	// Amount of planned fuel on board when the point is reached in pounds.
	PlannedFuel param.Opt[float64] `json:"plannedFuel,omitzero"`
	// Name of the point group, usually Extended Operations (ETOPS), Critical Fuel
	// Point, and Equal Time Point (ETP) sections.
	PointGroupName param.Opt[string] `json:"pointGroupName,omitzero"`
	// Specifies which Point Group case requires the most fuel.
	WorstFuelCase param.Opt[string] `json:"worstFuelCase,omitzero"`
	// Array of point data for this Point Group.
	FlightPlanPoints []FlightplanNewParamsFlightPlanPointGroupFlightPlanPoint `json:"flightPlanPoints,omitzero"`
	paramObj
}

func (r FlightplanNewParamsFlightPlanPointGroup) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanNewParamsFlightPlanPointGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanNewParamsFlightPlanPointGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Array of point data for this Point Group.
type FlightplanNewParamsFlightPlanPointGroupFlightPlanPoint struct {
	// Estimated Time of Arrival (ETA) at this point in ISO 8601 UTC format, with
	// millisecond precision.
	FppEta param.Opt[time.Time] `json:"fppEta,omitzero" format:"date-time"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	FppLat param.Opt[float64] `json:"fppLat,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	FppLon param.Opt[float64] `json:"fppLon,omitzero"`
	// Fuel required at this point to execute an Equal Time Point (ETP) or Extended
	// Operations (ETOPS) plan in pounds.
	FppReqFuel param.Opt[float64] `json:"fppReqFuel,omitzero"`
	// Name of this point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	paramObj
}

func (r FlightplanNewParamsFlightPlanPointGroupFlightPlanPoint) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanNewParamsFlightPlanPointGroupFlightPlanPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanNewParamsFlightPlanPointGroupFlightPlanPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of waypoints associated with this flight plan.
//
// The properties Type, WaypointName are required.
type FlightplanNewParamsFlightPlanWaypoint struct {
	// Points are designated by type as either a comment point or a waypoint. A comment
	// point conveys important information about the point for pilots but is not
	// entered into a flight management system. A waypoint is a point that is entered
	// into a flight management system and/or filed with Air Traffic Control.
	Type string `json:"type,required"`
	// Name of the point. The name of a comment point identifies important information
	// about that point, e.g. Top of Climb. The name of a waypoint identifies the
	// location of that point.
	WaypointName string `json:"waypointName,required"`
	// The air-to-air Tactical Air Navigation (TACAN) channels used by the
	// receiver/tanker during air refueling.
	AaTacanChannel param.Opt[string] `json:"aaTacanChannel,omitzero"`
	// The air distance of this leg in nautical miles.
	AirDistance param.Opt[float64] `json:"airDistance,omitzero"`
	// The flight path flown for this leg.
	Airway param.Opt[string] `json:"airway,omitzero"`
	// Altitude of a level, point, or object measured in feet above mean sea level.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// The ID of the air refueling track/anchor or fixed track.
	ArID param.Opt[string] `json:"arId,omitzero"`
	// Point identifying an air refueling track/anchor or fixed track.
	Arpt param.Opt[string] `json:"arpt,omitzero"`
	// Actual Time of Arrival (ATA) at this waypoint in ISO 8601 UTC format, with
	// millisecond precision.
	Ata param.Opt[time.Time] `json:"ata,omitzero" format:"date-time"`
	// The average calibrated airspeed (CAS) for this leg in knots.
	AvgCalAirspeed param.Opt[float64] `json:"avgCalAirspeed,omitzero"`
	// The average drift angle for this leg in degrees from true north.
	AvgDriftAng param.Opt[float64] `json:"avgDriftAng,omitzero"`
	// The average ground speed for this leg in knots.
	AvgGroundSpeed param.Opt[float64] `json:"avgGroundSpeed,omitzero"`
	// The average true airspeed (TAS) for this leg in knots.
	AvgTrueAirspeed param.Opt[float64] `json:"avgTrueAirspeed,omitzero"`
	// The average wind direction for this leg in degrees from true north.
	AvgWindDir param.Opt[float64] `json:"avgWindDir,omitzero"`
	// The average wind speed for this leg in knots.
	AvgWindSpeed param.Opt[float64] `json:"avgWindSpeed,omitzero"`
	// The day low level altitude in feet above sea level for the leg ending at this
	// waypoint.
	DayLowAlt param.Opt[float64] `json:"dayLowAlt,omitzero"`
	// Estimated Time of Arrival (ETA) at this waypoint in ISO 8601 UTC format, with
	// millisecond precision.
	Eta param.Opt[time.Time] `json:"eta,omitzero" format:"date-time"`
	// The amount of fuel onloaded or offloaded at this waypoint in pounds (negative
	// value for offload).
	ExchangedFuel param.Opt[float64] `json:"exchangedFuel,omitzero"`
	// The leg fuel flow in pounds per hour.
	FuelFlow param.Opt[float64] `json:"fuelFlow,omitzero"`
	// The icing intensity classification for this flight (LIGHT, MODERATE, etc).
	IceCat param.Opt[string] `json:"iceCat,omitzero"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The planned alternate leg based on user-defined constraints, International Civil
	// Aviation Organization (ICAO) code preferred.
	LegAlternate param.Opt[string] `json:"legAlternate,omitzero"`
	// The percent degrade due to drag for this aircraft for this leg.
	LegDragIndex param.Opt[float64] `json:"legDragIndex,omitzero"`
	// The fuel degrade percentage used for this leg.
	LegFuelDegrade param.Opt[float64] `json:"legFuelDegrade,omitzero"`
	// The average Mach speed for this leg.
	LegMach param.Opt[float64] `json:"legMach,omitzero"`
	// The mission index value for this leg. The mission index is the ratio of
	// time-related cost of aircraft operation to the cost of fuel.
	LegMsnIndex param.Opt[float64] `json:"legMsnIndex,omitzero"`
	// The wind factor for this leg in knots. A positive value indicates a headwind,
	// while a negative value indicates a tailwind.
	LegWindFac param.Opt[float64] `json:"legWindFac,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The magnetic course at leg midpoint in degrees from true north.
	MagCourse param.Opt[float64] `json:"magCourse,omitzero"`
	// The magnetic heading at leg midpoint in degrees from true north.
	MagHeading param.Opt[float64] `json:"magHeading,omitzero"`
	// The magnetic variation for this leg in degrees.
	MagVar param.Opt[float64] `json:"magVar,omitzero"`
	// Navigational Aid (NAVAID) identification code.
	Navaid param.Opt[string] `json:"navaid,omitzero"`
	// The night low level altitude in feet above sea level for the leg ending at this
	// waypoint.
	NightLowAlt param.Opt[float64] `json:"nightLowAlt,omitzero"`
	// The night vision goggle low level altitude in feet above sea level for the leg
	// ending at this waypoint.
	NvgLowAlt param.Opt[float64] `json:"nvgLowAlt,omitzero"`
	// The wind direction at this specific point in degrees from true north.
	PointWindDir param.Opt[float64] `json:"pointWindDir,omitzero"`
	// The wind velocity at this specific point in knots.
	PointWindSpeed param.Opt[float64] `json:"pointWindSpeed,omitzero"`
	// The primary UHF radio frequency used for the air refueling track or anchor in
	// megahertz.
	PriFreq param.Opt[float64] `json:"priFreq,omitzero"`
	// The secondary UHF radio frequency used for the air refueling track or anchor in
	// megahertz.
	SecFreq param.Opt[float64] `json:"secFreq,omitzero"`
	// Tactical Air Navigation (TACAN) channel for the Navigational Aid (NAVAID).
	TacanChannel param.Opt[string] `json:"tacanChannel,omitzero"`
	// Average temperature deviation from standard day profile for this leg in degrees
	// Celsius.
	TempDev param.Opt[float64] `json:"tempDev,omitzero"`
	// The thunderstorm intensity classification for this flight (LIGHT, MODERATE,
	// etc).
	ThunderCat param.Opt[string] `json:"thunderCat,omitzero"`
	// The total air distance to this waypoint in nautical miles.
	TotalAirDistance param.Opt[float64] `json:"totalAirDistance,omitzero"`
	// The total distance flown to this waypoint calculated from point of departure in
	// nautical miles.
	TotalFlownDistance param.Opt[float64] `json:"totalFlownDistance,omitzero"`
	// The total distance remaining from this waypoint to the point of arrival in
	// nautical miles.
	TotalRemDistance param.Opt[float64] `json:"totalRemDistance,omitzero"`
	// The total fuel remaining at this waypoint in pounds.
	TotalRemFuel param.Opt[float64] `json:"totalRemFuel,omitzero"`
	// The total time accumulated from takeoff to this waypoint expressed as HH:MM.
	TotalTime param.Opt[string] `json:"totalTime,omitzero"`
	// The total time remaining from this waypoint to the point of arrival expressed as
	// HH:MM.
	TotalTimeRem param.Opt[string] `json:"totalTimeRem,omitzero"`
	// The total fuel used to this waypoint from point of departure in pounds.
	TotalUsedFuel param.Opt[float64] `json:"totalUsedFuel,omitzero"`
	// The total weight of the aircraft at this waypoint in pounds.
	TotalWeight param.Opt[float64] `json:"totalWeight,omitzero"`
	// The true course at leg midpoint in degrees from true north.
	TrueCourse param.Opt[float64] `json:"trueCourse,omitzero"`
	// The turbulence intensity classification for this flight (LIGHT, MODERATE, etc).
	TurbCat param.Opt[string] `json:"turbCat,omitzero"`
	// VHF Omni-directional Range (VOR) frequency for the Navigational Aid (NAVAID) in
	// megahertz.
	VorFreq param.Opt[float64] `json:"vorFreq,omitzero"`
	// The waypoint number on the route. Comment points do not get a waypoint number.
	WaypointNum param.Opt[int64] `json:"waypointNum,omitzero"`
	// The zone/leg distance flown in nautical miles.
	ZoneDistance param.Opt[float64] `json:"zoneDistance,omitzero"`
	// The amount of fuel used on this zone/leg in pounds.
	ZoneFuel param.Opt[float64] `json:"zoneFuel,omitzero"`
	// The time to fly this zone/leg in minutes.
	ZoneTime param.Opt[float64] `json:"zoneTime,omitzero"`
	paramObj
}

func (r FlightplanNewParamsFlightPlanWaypoint) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanNewParamsFlightPlanWaypoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanNewParamsFlightPlanWaypoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlightplanGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FlightplanGetParams]'s query parameters as `url.Values`.
func (r FlightplanGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FlightplanUpdateParams struct {
	// The airfield identifier of the arrival location, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAirfield string `json:"arrAirfield,required"`
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
	DataMode FlightplanUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The airfield identifier of the departure location, International Civil Aviation
	// Organization (ICAO) code preferred.
	DepAirfield string `json:"depAirfield,required"`
	// The generation time of this flight plan in ISO 8601 UTC format, with millisecond
	// precision.
	GenTs time.Time `json:"genTS,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this flight plan.
	// Intended as, but not constrained to, MIL-STD-6016 environment dependent specific
	// type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Air Mobility Command (AMC) mission identifier according to Mobility Air Forces
	// (MAF) encode/decode procedures.
	AmcMissionID param.Opt[string] `json:"amcMissionId,omitzero"`
	// Fuel burned from the initial approach point to landing in pounds.
	AppLandingFuel param.Opt[float64] `json:"appLandingFuel,omitzero"`
	// The first designated alternate arrival airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAlternate1 param.Opt[string] `json:"arrAlternate1,omitzero"`
	// Fuel required to fly to alternate landing site 1 and land in pounds.
	ArrAlternate1Fuel param.Opt[float64] `json:"arrAlternate1Fuel,omitzero"`
	// The second designated alternate arrival airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAlternate2 param.Opt[string] `json:"arrAlternate2,omitzero"`
	// Fuel required to fly to alternate landing site 2 and land in pounds.
	ArrAlternate2Fuel param.Opt[float64] `json:"arrAlternate2Fuel,omitzero"`
	// Additional fuel burned at landing/missed approach for icing during arrival in
	// pounds.
	ArrIceFuel param.Opt[float64] `json:"arrIceFuel,omitzero"`
	// The arrival runway for this flight.
	ArrRunway param.Opt[string] `json:"arrRunway,omitzero"`
	// Average temperature deviation of the primary, divert, and alternate path for the
	// route between first Top of Climb and last Top of Descent in degrees Celsius.
	AvgTempDev param.Opt[float64] `json:"avgTempDev,omitzero"`
	// Fuel planned to be burned during the flight in pounds.
	BurnedFuel param.Opt[float64] `json:"burnedFuel,omitzero"`
	// The call sign assigned to the aircraft for this flight plan.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Remarks about the planned cargo associated with this flight plan.
	CargoRemark param.Opt[string] `json:"cargoRemark,omitzero"`
	// Fuel required from brake release to Top of Climb in pounds.
	ClimbFuel param.Opt[float64] `json:"climbFuel,omitzero"`
	// Time required from brake release to Top of Climb expressed as HH:MM.
	ClimbTime param.Opt[string] `json:"climbTime,omitzero"`
	// The amount of contingency fuel in pounds.
	ContingencyFuel param.Opt[float64] `json:"contingencyFuel,omitzero"`
	// The designated alternate departure airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	DepAlternate param.Opt[string] `json:"depAlternate,omitzero"`
	// The depressurization fuel required to fly from the Equal Time Point to the Last
	// Suitable/First Suitable airfield at depressurization altitude in pounds.
	DepressFuel param.Opt[float64] `json:"depressFuel,omitzero"`
	// The departure runway for this flight.
	DepRunway param.Opt[string] `json:"depRunway,omitzero"`
	// The percent degrade due to drag for this aircraft.
	DragIndex param.Opt[float64] `json:"dragIndex,omitzero"`
	// Additional fuel burned at landing/missed approach for an early descent in
	// pounds.
	EarlyDescentFuel param.Opt[float64] `json:"earlyDescentFuel,omitzero"`
	// Total endurance time based on the fuel on board expressed as HH:MM.
	EnduranceTime param.Opt[string] `json:"enduranceTime,omitzero"`
	// Fuel required to fly from Top of Climb to Top of Descent in pounds.
	EnrouteFuel param.Opt[float64] `json:"enrouteFuel,omitzero"`
	// Time required to fly from Top of Climb to Top of Descent expressed as HH:MM.
	EnrouteTime param.Opt[string] `json:"enrouteTime,omitzero"`
	// The list of equipment on the aircraft as defined in the Flight Information
	// Publications (FLIP) General Planning (GP) manual.
	Equipment param.Opt[string] `json:"equipment,omitzero"`
	// The estimated time of departure for the aircraft, in ISO 8601 UTC format, with
	// millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// The Extended Operations (ETOPS) rating used to calculate this flight plan.
	EtopsRating param.Opt[string] `json:"etopsRating,omitzero"`
	// The Extended Operations (ETOPS) validity window for the alternate airfield.
	EtopsValWindow param.Opt[string] `json:"etopsValWindow,omitzero"`
	// The source ID of the flight plan from the generating system.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The flight rules this flight plan is being filed under.
	FlightRules param.Opt[string] `json:"flightRules,omitzero"`
	// The type of flight (MILITARY, CIVILIAN, etc).
	FlightType param.Opt[string] `json:"flightType,omitzero"`
	// The fuel degrade percentage used for this mission.
	FuelDegrade param.Opt[float64] `json:"fuelDegrade,omitzero"`
	// The GPS Receiver Autonomous Integrity Monitoring (RAIM) message. A RAIM system
	// assesses the integrity of the GPS signals. This system predicts outages for a
	// specified geographical area. These predictions are based on the location, path,
	// and scheduled GPS satellite outages.
	GpsRaim param.Opt[string] `json:"gpsRAIM,omitzero"`
	// Additional fuel burned at Top of Climb in pounds.
	HoldDownFuel param.Opt[float64] `json:"holdDownFuel,omitzero"`
	// Additional fuel burned at the destination for holding in pounds.
	HoldFuel param.Opt[float64] `json:"holdFuel,omitzero"`
	// Additional time for holding at the destination expressed as HH:MM.
	HoldTime param.Opt[string] `json:"holdTime,omitzero"`
	// The UDL unique identifier of the aircraft associated with this flight plan.
	IDAircraft param.Opt[string] `json:"idAircraft,omitzero"`
	// The UDL unique identifier of the arrival airfield associated with this flight
	// plan.
	IDArrAirfield param.Opt[string] `json:"idArrAirfield,omitzero"`
	// The UDL unique identifier of the departure airfield associated with this flight
	// plan.
	IDDepAirfield param.Opt[string] `json:"idDepAirfield,omitzero"`
	// The amount of identified extra fuel carried and not available in the burn plan
	// in pounds.
	IdentExtraFuel param.Opt[float64] `json:"identExtraFuel,omitzero"`
	// The UDL unique identifier of the aircraft sortie associated with this flight
	// plan.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// A character string representation of the initial filed cruise speed for this
	// flight (prepended values of K, N, and M represent kilometers per hour, knots,
	// and Mach, respectively).
	InitialCruiseSpeed param.Opt[string] `json:"initialCruiseSpeed,omitzero"`
	// A character string representation of the initial filed altitude level for this
	// flight (prepended values of F, S, A, and M represent flight level in hundreds of
	// feet, standard metric level in tens of meters, altitude in hundreds of feet, and
	// altitude in tens of meters, respectively).
	InitialFlightLevel param.Opt[string] `json:"initialFlightLevel,omitzero"`
	// Fuel planned to be remaining on the airplane at landing in pounds.
	LandingFuel param.Opt[float64] `json:"landingFuel,omitzero"`
	// The leg number of this flight plan.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The minimum fuel on board required to divert in pounds.
	MinDivertFuel param.Opt[float64] `json:"minDivertFuel,omitzero"`
	// The mission index value for this mission. The mission index is the ratio of
	// time-related cost of aircraft operation to the cost of fuel.
	MsnIndex param.Opt[float64] `json:"msnIndex,omitzero"`
	// Additional remarks for air traffic control for this flight.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// The number of aircraft flying this flight plan.
	NumAircraft param.Opt[int64] `json:"numAircraft,omitzero"`
	// Additional fuel burned at Top of Descent for the operational condition in
	// pounds.
	OpConditionFuel param.Opt[float64] `json:"opConditionFuel,omitzero"`
	// Operating weight of the aircraft in pounds.
	OpWeight param.Opt[float64] `json:"opWeight,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Air Traffic Control address filing the flight plan.
	Originator param.Opt[string] `json:"originator,omitzero"`
	// Remarks from the planners concerning this flight plan.
	PlannerRemark param.Opt[string] `json:"plannerRemark,omitzero"`
	// Total of all fuel required to complete the flight in pounds, including fuel to
	// be dispensed on a refueling mission.
	RampFuel param.Opt[float64] `json:"rampFuel,omitzero"`
	// Total fuel remaining at alternate landing site 1 in pounds.
	RemAlternate1Fuel param.Opt[float64] `json:"remAlternate1Fuel,omitzero"`
	// Total fuel remaining at alternate landing site 2 in pounds.
	RemAlternate2Fuel param.Opt[float64] `json:"remAlternate2Fuel,omitzero"`
	// The amount of reserve fuel in pounds.
	ReserveFuel param.Opt[float64] `json:"reserveFuel,omitzero"`
	// The 1801 fileable route of flight string for this flight. The route of flight
	// string contains route designators, significant points, change of speed/altitude,
	// change of flight rules, and cruise climbs.
	RouteString param.Opt[string] `json:"routeString,omitzero"`
	// Name of the planned Standard Instrument Departure (SID) procedure.
	Sid param.Opt[string] `json:"sid,omitzero"`
	// Name of the planned Standard Terminal Arrival (STAR) procedure.
	Star param.Opt[string] `json:"star,omitzero"`
	// Status of this flight plan (e.g., ACTIVE, APPROVED, PLANNED, etc.).
	Status param.Opt[string] `json:"status,omitzero"`
	// The tail number of the aircraft associated with this flight plan.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// Fuel at takeoff, which is calculated as the ramp fuel minus the taxi fuel in
	// pounds.
	TakeoffFuel param.Opt[float64] `json:"takeoffFuel,omitzero"`
	// Fuel required to start engines and taxi to the end of the runway in pounds.
	TaxiFuel param.Opt[float64] `json:"taxiFuel,omitzero"`
	// Additional fuel burned at Top of Descent for thunderstorm avoidance in pounds.
	ThunderAvoidFuel param.Opt[float64] `json:"thunderAvoidFuel,omitzero"`
	// Fuel remaining at Top of Climb in pounds.
	TocFuel param.Opt[float64] `json:"tocFuel,omitzero"`
	// Additional fuel burned at Top of Climb for icing in pounds.
	TocIceFuel param.Opt[float64] `json:"tocIceFuel,omitzero"`
	// Fuel remaining at Top of Descent in pounds.
	TodFuel param.Opt[float64] `json:"todFuel,omitzero"`
	// Additional fuel burned at Top of Descent for icing in pounds.
	TodIceFuel param.Opt[float64] `json:"todIceFuel,omitzero"`
	// The amount of unidentified extra fuel required to get to min landing in pounds.
	UnidentExtraFuel param.Opt[float64] `json:"unidentExtraFuel,omitzero"`
	// The amount of unusable fuel in pounds.
	UnusableFuel param.Opt[float64] `json:"unusableFuel,omitzero"`
	// The wake turbulence category for this flight. The categories are assigned by the
	// International Civil Aviation Organization (ICAO) and are based on maximum
	// certified takeoff mass for the purpose of separating aircraft in flight due to
	// wake turbulence. Valid values include LIGHT, MEDIUM, LARGE, HEAVY, and SUPER.
	WakeTurbCat param.Opt[string] `json:"wakeTurbCat,omitzero"`
	// Wind factor for the first half of the route. This is the average wind factor
	// from first Top of Climb to the mid-time of the entire route in knots. A positive
	// value indicates a headwind, while a negative value indicates a tailwind.
	WindFac1 param.Opt[float64] `json:"windFac1,omitzero"`
	// Wind factor for the second half of the route. This is the average wind factor
	// from the mid-time of the entire route to last Top of Descent in knots. A
	// positive value indicates a headwind, while a negative value indicates a
	// tailwind.
	WindFac2 param.Opt[float64] `json:"windFac2,omitzero"`
	// Average wind factor from Top of Climb to Top of Descent in knots. A positive
	// value indicates a headwind, while a negative value indicates a tailwind.
	WindFacAvg param.Opt[float64] `json:"windFacAvg,omitzero"`
	// The date and time the weather valid period ends in ISO 8601 UTC format, with
	// millisecond precision.
	WxValidEnd param.Opt[time.Time] `json:"wxValidEnd,omitzero" format:"date-time"`
	// The date and time the weather valid period begins in ISO 8601 UTC format, with
	// millisecond precision.
	WxValidStart param.Opt[time.Time] `json:"wxValidStart,omitzero" format:"date-time"`
	// Collection of air refueling events occurring on this flight.
	AirRefuelEvents []FlightplanUpdateParamsAirRefuelEvent `json:"airRefuelEvents,omitzero"`
	// Array of Air Traffic Control (ATC) addresses.
	AtcAddresses []string `json:"atcAddresses,omitzero"`
	// Array of country codes for the countries overflown during this flight in ISO
	// 3166-1 Alpha-2 format.
	CountryCodes []string `json:"countryCodes,omitzero"`
	// Array of Extended Operations (ETOPS) adequate landing airfields that are within
	// the mission region.
	EtopsAirfields []string `json:"etopsAirfields,omitzero"`
	// Array of Extended Operations (ETOPS) alternate suitable landing airfields that
	// are within the mission region.
	EtopsAltAirfields []string `json:"etopsAltAirfields,omitzero"`
	// Collection of messages associated with this flight plan indicating the severity,
	// the point where the message was generated, the path (Primary, Alternate, etc.),
	// and the text of the message.
	FlightPlanMessages []FlightplanUpdateParamsFlightPlanMessage `json:"flightPlanMessages,omitzero"`
	// Collection of point groups generated for this flight plan. Groups include point
	// sets for Extended Operations (ETOPS), Critical Fuel Point, and Equal Time Point
	// (ETP).
	FlightPlanPointGroups []FlightplanUpdateParamsFlightPlanPointGroup `json:"flightPlanPointGroups,omitzero"`
	// Collection of waypoints associated with this flight plan.
	FlightPlanWaypoints []FlightplanUpdateParamsFlightPlanWaypoint `json:"flightPlanWaypoints,omitzero"`
	paramObj
}

func (r FlightplanUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUpdateParams) UnmarshalJSON(data []byte) error {
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
type FlightplanUpdateParamsDataMode string

const (
	FlightplanUpdateParamsDataModeReal      FlightplanUpdateParamsDataMode = "REAL"
	FlightplanUpdateParamsDataModeTest      FlightplanUpdateParamsDataMode = "TEST"
	FlightplanUpdateParamsDataModeSimulated FlightplanUpdateParamsDataMode = "SIMULATED"
	FlightplanUpdateParamsDataModeExercise  FlightplanUpdateParamsDataMode = "EXERCISE"
)

// Collection of air refueling events occurring on this flight.
type FlightplanUpdateParamsAirRefuelEvent struct {
	// Additional degrade for air refueling, cumulative with fuelDegrade field percent.
	ArDegrade param.Opt[float64] `json:"arDegrade,omitzero"`
	// Fuel onloaded (use positive numbers) or fuel offloaded (use negative numbers) in
	// pounds.
	ArExchangedFuel param.Opt[float64] `json:"arExchangedFuel,omitzero"`
	// The number of this air refueling event within the flight plan.
	ArNum param.Opt[int64] `json:"arNum,omitzero"`
	// Fuel required to fly from air refueling exit point to air refueling divert
	// alternate airfield in pounds.
	DivertFuel param.Opt[float64] `json:"divertFuel,omitzero"`
	// Fuel remaining at the air refueling exit in pounds.
	ExitFuel param.Opt[float64] `json:"exitFuel,omitzero"`
	paramObj
}

func (r FlightplanUpdateParamsAirRefuelEvent) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUpdateParamsAirRefuelEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUpdateParamsAirRefuelEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of messages associated with this flight plan indicating the severity,
// the point where the message was generated, the path (Primary, Alternate, etc.),
// and the text of the message.
type FlightplanUpdateParamsFlightPlanMessage struct {
	// The text of the message.
	MsgText param.Opt[string] `json:"msgText,omitzero"`
	// The flight path that generated the message (PRIMARY, ALTERNATE, etc.).
	RoutePath param.Opt[string] `json:"routePath,omitzero"`
	// The severity of the message.
	Severity param.Opt[string] `json:"severity,omitzero"`
	// The waypoint number for which the message was generated, or enter "PLAN" for a
	// message impacting the entire route.
	WpNum param.Opt[string] `json:"wpNum,omitzero"`
	paramObj
}

func (r FlightplanUpdateParamsFlightPlanMessage) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUpdateParamsFlightPlanMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUpdateParamsFlightPlanMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of point groups generated for this flight plan. Groups include point
// sets for Extended Operations (ETOPS), Critical Fuel Point, and Equal Time Point
// (ETP).
type FlightplanUpdateParamsFlightPlanPointGroup struct {
	// Average fuel flow at which the fuel was calculated in pounds per hour.
	AvgFuelFlow param.Opt[float64] `json:"avgFuelFlow,omitzero"`
	// Average wind factor from the Extended Operations (ETOPS) point to the recovery
	// airfield in knots.
	EtopsAvgWindFactor param.Opt[float64] `json:"etopsAvgWindFactor,omitzero"`
	// Distance from the Extended Operations (ETOPS) point to the recovery airfield in
	// nautical miles.
	EtopsDistance param.Opt[float64] `json:"etopsDistance,omitzero"`
	// Fuel required to fly from the Extended Operations (ETOPS) point to the recovery
	// airfield in pounds.
	EtopsReqFuel param.Opt[float64] `json:"etopsReqFuel,omitzero"`
	// Temperature deviation from the Extended Operations (ETOPS) point to the recovery
	// airfield in degrees Celsius.
	EtopsTempDev param.Opt[float64] `json:"etopsTempDev,omitzero"`
	// Time to fly from the Extended Operations (ETOPS) point to the recovery airfield
	// expressed in HH:MM format.
	EtopsTime param.Opt[string] `json:"etopsTime,omitzero"`
	// Total time from takeoff when the point is reached expressed in HH:MM format.
	FromTakeoffTime param.Opt[string] `json:"fromTakeoffTime,omitzero"`
	// Average wind factor from the Equal Time Point (ETP) to the first suitable
	// airfield in knots.
	FsafAvgWindFactor param.Opt[float64] `json:"fsafAvgWindFactor,omitzero"`
	// Distance from the Equal Time Point (ETP) to the first suitable airfield in
	// nautical miles.
	FsafDistance param.Opt[float64] `json:"fsafDistance,omitzero"`
	// Fuel required to fly from the Equal Time Point (ETP) to the first suitable
	// airfield in pounds.
	FsafReqFuel param.Opt[float64] `json:"fsafReqFuel,omitzero"`
	// Temperature deviation from the Equal Time Point (ETP) to the first suitable
	// airfield in degrees Celsius.
	FsafTempDev param.Opt[float64] `json:"fsafTempDev,omitzero"`
	// Time to fly from the Equal Time Point (ETP) to the first suitable airfield
	// expressed in HH:MM format.
	FsafTime param.Opt[string] `json:"fsafTime,omitzero"`
	// Flight level of the point at which the fuel was calculated in feet.
	FuelCalcAlt param.Opt[float64] `json:"fuelCalcAlt,omitzero"`
	// True airspeed at which the fuel was calculated in knots.
	FuelCalcSpd param.Opt[float64] `json:"fuelCalcSpd,omitzero"`
	// Average wind factor from the Equal Time Point (ETP) to the last suitable
	// airfield in knots.
	LsafAvgWindFactor param.Opt[float64] `json:"lsafAvgWindFactor,omitzero"`
	// Distance from the Equal Time Point (ETP) to the last suitable airfield in
	// nautical miles.
	LsafDistance param.Opt[float64] `json:"lsafDistance,omitzero"`
	// Name of the last suitable airfield, International Civil Aviation Organization
	// (ICAO) code preferred.
	LsafName param.Opt[string] `json:"lsafName,omitzero"`
	// Fuel required to fly from the Equal Time Point (ETP) to the last suitable
	// airfield in pounds.
	LsafReqFuel param.Opt[float64] `json:"lsafReqFuel,omitzero"`
	// Temperature deviation from the Equal Time Point (ETP) to the last suitable
	// airfield in degrees Celsius.
	LsafTempDev param.Opt[float64] `json:"lsafTempDev,omitzero"`
	// Time to fly from the Equal Time Point (ETP) to the last suitable airfield
	// expressed in HH:MM format.
	LsafTime param.Opt[string] `json:"lsafTime,omitzero"`
	// Amount of planned fuel on board when the point is reached in pounds.
	PlannedFuel param.Opt[float64] `json:"plannedFuel,omitzero"`
	// Name of the point group, usually Extended Operations (ETOPS), Critical Fuel
	// Point, and Equal Time Point (ETP) sections.
	PointGroupName param.Opt[string] `json:"pointGroupName,omitzero"`
	// Specifies which Point Group case requires the most fuel.
	WorstFuelCase param.Opt[string] `json:"worstFuelCase,omitzero"`
	// Array of point data for this Point Group.
	FlightPlanPoints []FlightplanUpdateParamsFlightPlanPointGroupFlightPlanPoint `json:"flightPlanPoints,omitzero"`
	paramObj
}

func (r FlightplanUpdateParamsFlightPlanPointGroup) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUpdateParamsFlightPlanPointGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUpdateParamsFlightPlanPointGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Array of point data for this Point Group.
type FlightplanUpdateParamsFlightPlanPointGroupFlightPlanPoint struct {
	// Estimated Time of Arrival (ETA) at this point in ISO 8601 UTC format, with
	// millisecond precision.
	FppEta param.Opt[time.Time] `json:"fppEta,omitzero" format:"date-time"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	FppLat param.Opt[float64] `json:"fppLat,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	FppLon param.Opt[float64] `json:"fppLon,omitzero"`
	// Fuel required at this point to execute an Equal Time Point (ETP) or Extended
	// Operations (ETOPS) plan in pounds.
	FppReqFuel param.Opt[float64] `json:"fppReqFuel,omitzero"`
	// Name of this point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	paramObj
}

func (r FlightplanUpdateParamsFlightPlanPointGroupFlightPlanPoint) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUpdateParamsFlightPlanPointGroupFlightPlanPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUpdateParamsFlightPlanPointGroupFlightPlanPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of waypoints associated with this flight plan.
//
// The properties Type, WaypointName are required.
type FlightplanUpdateParamsFlightPlanWaypoint struct {
	// Points are designated by type as either a comment point or a waypoint. A comment
	// point conveys important information about the point for pilots but is not
	// entered into a flight management system. A waypoint is a point that is entered
	// into a flight management system and/or filed with Air Traffic Control.
	Type string `json:"type,required"`
	// Name of the point. The name of a comment point identifies important information
	// about that point, e.g. Top of Climb. The name of a waypoint identifies the
	// location of that point.
	WaypointName string `json:"waypointName,required"`
	// The air-to-air Tactical Air Navigation (TACAN) channels used by the
	// receiver/tanker during air refueling.
	AaTacanChannel param.Opt[string] `json:"aaTacanChannel,omitzero"`
	// The air distance of this leg in nautical miles.
	AirDistance param.Opt[float64] `json:"airDistance,omitzero"`
	// The flight path flown for this leg.
	Airway param.Opt[string] `json:"airway,omitzero"`
	// Altitude of a level, point, or object measured in feet above mean sea level.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// The ID of the air refueling track/anchor or fixed track.
	ArID param.Opt[string] `json:"arId,omitzero"`
	// Point identifying an air refueling track/anchor or fixed track.
	Arpt param.Opt[string] `json:"arpt,omitzero"`
	// Actual Time of Arrival (ATA) at this waypoint in ISO 8601 UTC format, with
	// millisecond precision.
	Ata param.Opt[time.Time] `json:"ata,omitzero" format:"date-time"`
	// The average calibrated airspeed (CAS) for this leg in knots.
	AvgCalAirspeed param.Opt[float64] `json:"avgCalAirspeed,omitzero"`
	// The average drift angle for this leg in degrees from true north.
	AvgDriftAng param.Opt[float64] `json:"avgDriftAng,omitzero"`
	// The average ground speed for this leg in knots.
	AvgGroundSpeed param.Opt[float64] `json:"avgGroundSpeed,omitzero"`
	// The average true airspeed (TAS) for this leg in knots.
	AvgTrueAirspeed param.Opt[float64] `json:"avgTrueAirspeed,omitzero"`
	// The average wind direction for this leg in degrees from true north.
	AvgWindDir param.Opt[float64] `json:"avgWindDir,omitzero"`
	// The average wind speed for this leg in knots.
	AvgWindSpeed param.Opt[float64] `json:"avgWindSpeed,omitzero"`
	// The day low level altitude in feet above sea level for the leg ending at this
	// waypoint.
	DayLowAlt param.Opt[float64] `json:"dayLowAlt,omitzero"`
	// Estimated Time of Arrival (ETA) at this waypoint in ISO 8601 UTC format, with
	// millisecond precision.
	Eta param.Opt[time.Time] `json:"eta,omitzero" format:"date-time"`
	// The amount of fuel onloaded or offloaded at this waypoint in pounds (negative
	// value for offload).
	ExchangedFuel param.Opt[float64] `json:"exchangedFuel,omitzero"`
	// The leg fuel flow in pounds per hour.
	FuelFlow param.Opt[float64] `json:"fuelFlow,omitzero"`
	// The icing intensity classification for this flight (LIGHT, MODERATE, etc).
	IceCat param.Opt[string] `json:"iceCat,omitzero"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The planned alternate leg based on user-defined constraints, International Civil
	// Aviation Organization (ICAO) code preferred.
	LegAlternate param.Opt[string] `json:"legAlternate,omitzero"`
	// The percent degrade due to drag for this aircraft for this leg.
	LegDragIndex param.Opt[float64] `json:"legDragIndex,omitzero"`
	// The fuel degrade percentage used for this leg.
	LegFuelDegrade param.Opt[float64] `json:"legFuelDegrade,omitzero"`
	// The average Mach speed for this leg.
	LegMach param.Opt[float64] `json:"legMach,omitzero"`
	// The mission index value for this leg. The mission index is the ratio of
	// time-related cost of aircraft operation to the cost of fuel.
	LegMsnIndex param.Opt[float64] `json:"legMsnIndex,omitzero"`
	// The wind factor for this leg in knots. A positive value indicates a headwind,
	// while a negative value indicates a tailwind.
	LegWindFac param.Opt[float64] `json:"legWindFac,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The magnetic course at leg midpoint in degrees from true north.
	MagCourse param.Opt[float64] `json:"magCourse,omitzero"`
	// The magnetic heading at leg midpoint in degrees from true north.
	MagHeading param.Opt[float64] `json:"magHeading,omitzero"`
	// The magnetic variation for this leg in degrees.
	MagVar param.Opt[float64] `json:"magVar,omitzero"`
	// Navigational Aid (NAVAID) identification code.
	Navaid param.Opt[string] `json:"navaid,omitzero"`
	// The night low level altitude in feet above sea level for the leg ending at this
	// waypoint.
	NightLowAlt param.Opt[float64] `json:"nightLowAlt,omitzero"`
	// The night vision goggle low level altitude in feet above sea level for the leg
	// ending at this waypoint.
	NvgLowAlt param.Opt[float64] `json:"nvgLowAlt,omitzero"`
	// The wind direction at this specific point in degrees from true north.
	PointWindDir param.Opt[float64] `json:"pointWindDir,omitzero"`
	// The wind velocity at this specific point in knots.
	PointWindSpeed param.Opt[float64] `json:"pointWindSpeed,omitzero"`
	// The primary UHF radio frequency used for the air refueling track or anchor in
	// megahertz.
	PriFreq param.Opt[float64] `json:"priFreq,omitzero"`
	// The secondary UHF radio frequency used for the air refueling track or anchor in
	// megahertz.
	SecFreq param.Opt[float64] `json:"secFreq,omitzero"`
	// Tactical Air Navigation (TACAN) channel for the Navigational Aid (NAVAID).
	TacanChannel param.Opt[string] `json:"tacanChannel,omitzero"`
	// Average temperature deviation from standard day profile for this leg in degrees
	// Celsius.
	TempDev param.Opt[float64] `json:"tempDev,omitzero"`
	// The thunderstorm intensity classification for this flight (LIGHT, MODERATE,
	// etc).
	ThunderCat param.Opt[string] `json:"thunderCat,omitzero"`
	// The total air distance to this waypoint in nautical miles.
	TotalAirDistance param.Opt[float64] `json:"totalAirDistance,omitzero"`
	// The total distance flown to this waypoint calculated from point of departure in
	// nautical miles.
	TotalFlownDistance param.Opt[float64] `json:"totalFlownDistance,omitzero"`
	// The total distance remaining from this waypoint to the point of arrival in
	// nautical miles.
	TotalRemDistance param.Opt[float64] `json:"totalRemDistance,omitzero"`
	// The total fuel remaining at this waypoint in pounds.
	TotalRemFuel param.Opt[float64] `json:"totalRemFuel,omitzero"`
	// The total time accumulated from takeoff to this waypoint expressed as HH:MM.
	TotalTime param.Opt[string] `json:"totalTime,omitzero"`
	// The total time remaining from this waypoint to the point of arrival expressed as
	// HH:MM.
	TotalTimeRem param.Opt[string] `json:"totalTimeRem,omitzero"`
	// The total fuel used to this waypoint from point of departure in pounds.
	TotalUsedFuel param.Opt[float64] `json:"totalUsedFuel,omitzero"`
	// The total weight of the aircraft at this waypoint in pounds.
	TotalWeight param.Opt[float64] `json:"totalWeight,omitzero"`
	// The true course at leg midpoint in degrees from true north.
	TrueCourse param.Opt[float64] `json:"trueCourse,omitzero"`
	// The turbulence intensity classification for this flight (LIGHT, MODERATE, etc).
	TurbCat param.Opt[string] `json:"turbCat,omitzero"`
	// VHF Omni-directional Range (VOR) frequency for the Navigational Aid (NAVAID) in
	// megahertz.
	VorFreq param.Opt[float64] `json:"vorFreq,omitzero"`
	// The waypoint number on the route. Comment points do not get a waypoint number.
	WaypointNum param.Opt[int64] `json:"waypointNum,omitzero"`
	// The zone/leg distance flown in nautical miles.
	ZoneDistance param.Opt[float64] `json:"zoneDistance,omitzero"`
	// The amount of fuel used on this zone/leg in pounds.
	ZoneFuel param.Opt[float64] `json:"zoneFuel,omitzero"`
	// The time to fly this zone/leg in minutes.
	ZoneTime param.Opt[float64] `json:"zoneTime,omitzero"`
	paramObj
}

func (r FlightplanUpdateParamsFlightPlanWaypoint) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUpdateParamsFlightPlanWaypoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUpdateParamsFlightPlanWaypoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlightplanListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FlightplanListParams]'s query parameters as `url.Values`.
func (r FlightplanListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FlightplanCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FlightplanCountParams]'s query parameters as `url.Values`.
func (r FlightplanCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FlightplanTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FlightplanTupleParams]'s query parameters as `url.Values`.
func (r FlightplanTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FlightplanUnvalidatedPublishParams struct {
	Body []FlightplanUnvalidatedPublishParamsBody
	paramObj
}

func (r FlightplanUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *FlightplanUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Flight Plan contains data specifying the details of an intended flight including
// schedule and expected route.
//
// The properties ArrAirfield, ClassificationMarking, DataMode, DepAirfield, GenTs,
// Source are required.
type FlightplanUnvalidatedPublishParamsBody struct {
	// The airfield identifier of the arrival location, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAirfield string `json:"arrAirfield,required"`
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
	// The airfield identifier of the departure location, International Civil Aviation
	// Organization (ICAO) code preferred.
	DepAirfield string `json:"depAirfield,required"`
	// The generation time of this flight plan in ISO 8601 UTC format, with millisecond
	// precision.
	GenTs time.Time `json:"genTS,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this flight plan.
	// Intended as, but not constrained to, MIL-STD-6016 environment dependent specific
	// type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Air Mobility Command (AMC) mission identifier according to Mobility Air Forces
	// (MAF) encode/decode procedures.
	AmcMissionID param.Opt[string] `json:"amcMissionId,omitzero"`
	// Fuel burned from the initial approach point to landing in pounds.
	AppLandingFuel param.Opt[float64] `json:"appLandingFuel,omitzero"`
	// The first designated alternate arrival airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAlternate1 param.Opt[string] `json:"arrAlternate1,omitzero"`
	// Fuel required to fly to alternate landing site 1 and land in pounds.
	ArrAlternate1Fuel param.Opt[float64] `json:"arrAlternate1Fuel,omitzero"`
	// The second designated alternate arrival airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	ArrAlternate2 param.Opt[string] `json:"arrAlternate2,omitzero"`
	// Fuel required to fly to alternate landing site 2 and land in pounds.
	ArrAlternate2Fuel param.Opt[float64] `json:"arrAlternate2Fuel,omitzero"`
	// Additional fuel burned at landing/missed approach for icing during arrival in
	// pounds.
	ArrIceFuel param.Opt[float64] `json:"arrIceFuel,omitzero"`
	// The arrival runway for this flight.
	ArrRunway param.Opt[string] `json:"arrRunway,omitzero"`
	// Average temperature deviation of the primary, divert, and alternate path for the
	// route between first Top of Climb and last Top of Descent in degrees Celsius.
	AvgTempDev param.Opt[float64] `json:"avgTempDev,omitzero"`
	// Fuel planned to be burned during the flight in pounds.
	BurnedFuel param.Opt[float64] `json:"burnedFuel,omitzero"`
	// The call sign assigned to the aircraft for this flight plan.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Remarks about the planned cargo associated with this flight plan.
	CargoRemark param.Opt[string] `json:"cargoRemark,omitzero"`
	// Fuel required from brake release to Top of Climb in pounds.
	ClimbFuel param.Opt[float64] `json:"climbFuel,omitzero"`
	// Time required from brake release to Top of Climb expressed as HH:MM.
	ClimbTime param.Opt[string] `json:"climbTime,omitzero"`
	// The amount of contingency fuel in pounds.
	ContingencyFuel param.Opt[float64] `json:"contingencyFuel,omitzero"`
	// The designated alternate departure airfield, International Civil Aviation
	// Organization (ICAO) code preferred.
	DepAlternate param.Opt[string] `json:"depAlternate,omitzero"`
	// The depressurization fuel required to fly from the Equal Time Point to the Last
	// Suitable/First Suitable airfield at depressurization altitude in pounds.
	DepressFuel param.Opt[float64] `json:"depressFuel,omitzero"`
	// The departure runway for this flight.
	DepRunway param.Opt[string] `json:"depRunway,omitzero"`
	// The percent degrade due to drag for this aircraft.
	DragIndex param.Opt[float64] `json:"dragIndex,omitzero"`
	// Additional fuel burned at landing/missed approach for an early descent in
	// pounds.
	EarlyDescentFuel param.Opt[float64] `json:"earlyDescentFuel,omitzero"`
	// Total endurance time based on the fuel on board expressed as HH:MM.
	EnduranceTime param.Opt[string] `json:"enduranceTime,omitzero"`
	// Fuel required to fly from Top of Climb to Top of Descent in pounds.
	EnrouteFuel param.Opt[float64] `json:"enrouteFuel,omitzero"`
	// Time required to fly from Top of Climb to Top of Descent expressed as HH:MM.
	EnrouteTime param.Opt[string] `json:"enrouteTime,omitzero"`
	// The list of equipment on the aircraft as defined in the Flight Information
	// Publications (FLIP) General Planning (GP) manual.
	Equipment param.Opt[string] `json:"equipment,omitzero"`
	// The estimated time of departure for the aircraft, in ISO 8601 UTC format, with
	// millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// The Extended Operations (ETOPS) rating used to calculate this flight plan.
	EtopsRating param.Opt[string] `json:"etopsRating,omitzero"`
	// The Extended Operations (ETOPS) validity window for the alternate airfield.
	EtopsValWindow param.Opt[string] `json:"etopsValWindow,omitzero"`
	// The source ID of the flight plan from the generating system.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The flight rules this flight plan is being filed under.
	FlightRules param.Opt[string] `json:"flightRules,omitzero"`
	// The type of flight (MILITARY, CIVILIAN, etc).
	FlightType param.Opt[string] `json:"flightType,omitzero"`
	// The fuel degrade percentage used for this mission.
	FuelDegrade param.Opt[float64] `json:"fuelDegrade,omitzero"`
	// The GPS Receiver Autonomous Integrity Monitoring (RAIM) message. A RAIM system
	// assesses the integrity of the GPS signals. This system predicts outages for a
	// specified geographical area. These predictions are based on the location, path,
	// and scheduled GPS satellite outages.
	GpsRaim param.Opt[string] `json:"gpsRAIM,omitzero"`
	// Additional fuel burned at Top of Climb in pounds.
	HoldDownFuel param.Opt[float64] `json:"holdDownFuel,omitzero"`
	// Additional fuel burned at the destination for holding in pounds.
	HoldFuel param.Opt[float64] `json:"holdFuel,omitzero"`
	// Additional time for holding at the destination expressed as HH:MM.
	HoldTime param.Opt[string] `json:"holdTime,omitzero"`
	// The UDL unique identifier of the aircraft associated with this flight plan.
	IDAircraft param.Opt[string] `json:"idAircraft,omitzero"`
	// The UDL unique identifier of the arrival airfield associated with this flight
	// plan.
	IDArrAirfield param.Opt[string] `json:"idArrAirfield,omitzero"`
	// The UDL unique identifier of the departure airfield associated with this flight
	// plan.
	IDDepAirfield param.Opt[string] `json:"idDepAirfield,omitzero"`
	// The amount of identified extra fuel carried and not available in the burn plan
	// in pounds.
	IdentExtraFuel param.Opt[float64] `json:"identExtraFuel,omitzero"`
	// The UDL unique identifier of the aircraft sortie associated with this flight
	// plan.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// A character string representation of the initial filed cruise speed for this
	// flight (prepended values of K, N, and M represent kilometers per hour, knots,
	// and Mach, respectively).
	InitialCruiseSpeed param.Opt[string] `json:"initialCruiseSpeed,omitzero"`
	// A character string representation of the initial filed altitude level for this
	// flight (prepended values of F, S, A, and M represent flight level in hundreds of
	// feet, standard metric level in tens of meters, altitude in hundreds of feet, and
	// altitude in tens of meters, respectively).
	InitialFlightLevel param.Opt[string] `json:"initialFlightLevel,omitzero"`
	// Fuel planned to be remaining on the airplane at landing in pounds.
	LandingFuel param.Opt[float64] `json:"landingFuel,omitzero"`
	// The leg number of this flight plan.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The minimum fuel on board required to divert in pounds.
	MinDivertFuel param.Opt[float64] `json:"minDivertFuel,omitzero"`
	// The mission index value for this mission. The mission index is the ratio of
	// time-related cost of aircraft operation to the cost of fuel.
	MsnIndex param.Opt[float64] `json:"msnIndex,omitzero"`
	// Additional remarks for air traffic control for this flight.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// The number of aircraft flying this flight plan.
	NumAircraft param.Opt[int64] `json:"numAircraft,omitzero"`
	// Additional fuel burned at Top of Descent for the operational condition in
	// pounds.
	OpConditionFuel param.Opt[float64] `json:"opConditionFuel,omitzero"`
	// Operating weight of the aircraft in pounds.
	OpWeight param.Opt[float64] `json:"opWeight,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Air Traffic Control address filing the flight plan.
	Originator param.Opt[string] `json:"originator,omitzero"`
	// Remarks from the planners concerning this flight plan.
	PlannerRemark param.Opt[string] `json:"plannerRemark,omitzero"`
	// Total of all fuel required to complete the flight in pounds, including fuel to
	// be dispensed on a refueling mission.
	RampFuel param.Opt[float64] `json:"rampFuel,omitzero"`
	// Total fuel remaining at alternate landing site 1 in pounds.
	RemAlternate1Fuel param.Opt[float64] `json:"remAlternate1Fuel,omitzero"`
	// Total fuel remaining at alternate landing site 2 in pounds.
	RemAlternate2Fuel param.Opt[float64] `json:"remAlternate2Fuel,omitzero"`
	// The amount of reserve fuel in pounds.
	ReserveFuel param.Opt[float64] `json:"reserveFuel,omitzero"`
	// The 1801 fileable route of flight string for this flight. The route of flight
	// string contains route designators, significant points, change of speed/altitude,
	// change of flight rules, and cruise climbs.
	RouteString param.Opt[string] `json:"routeString,omitzero"`
	// Name of the planned Standard Instrument Departure (SID) procedure.
	Sid param.Opt[string] `json:"sid,omitzero"`
	// Name of the planned Standard Terminal Arrival (STAR) procedure.
	Star param.Opt[string] `json:"star,omitzero"`
	// Status of this flight plan (e.g., ACTIVE, APPROVED, PLANNED, etc.).
	Status param.Opt[string] `json:"status,omitzero"`
	// The tail number of the aircraft associated with this flight plan.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// Fuel at takeoff, which is calculated as the ramp fuel minus the taxi fuel in
	// pounds.
	TakeoffFuel param.Opt[float64] `json:"takeoffFuel,omitzero"`
	// Fuel required to start engines and taxi to the end of the runway in pounds.
	TaxiFuel param.Opt[float64] `json:"taxiFuel,omitzero"`
	// Additional fuel burned at Top of Descent for thunderstorm avoidance in pounds.
	ThunderAvoidFuel param.Opt[float64] `json:"thunderAvoidFuel,omitzero"`
	// Fuel remaining at Top of Climb in pounds.
	TocFuel param.Opt[float64] `json:"tocFuel,omitzero"`
	// Additional fuel burned at Top of Climb for icing in pounds.
	TocIceFuel param.Opt[float64] `json:"tocIceFuel,omitzero"`
	// Fuel remaining at Top of Descent in pounds.
	TodFuel param.Opt[float64] `json:"todFuel,omitzero"`
	// Additional fuel burned at Top of Descent for icing in pounds.
	TodIceFuel param.Opt[float64] `json:"todIceFuel,omitzero"`
	// The amount of unidentified extra fuel required to get to min landing in pounds.
	UnidentExtraFuel param.Opt[float64] `json:"unidentExtraFuel,omitzero"`
	// The amount of unusable fuel in pounds.
	UnusableFuel param.Opt[float64] `json:"unusableFuel,omitzero"`
	// The wake turbulence category for this flight. The categories are assigned by the
	// International Civil Aviation Organization (ICAO) and are based on maximum
	// certified takeoff mass for the purpose of separating aircraft in flight due to
	// wake turbulence. Valid values include LIGHT, MEDIUM, LARGE, HEAVY, and SUPER.
	WakeTurbCat param.Opt[string] `json:"wakeTurbCat,omitzero"`
	// Wind factor for the first half of the route. This is the average wind factor
	// from first Top of Climb to the mid-time of the entire route in knots. A positive
	// value indicates a headwind, while a negative value indicates a tailwind.
	WindFac1 param.Opt[float64] `json:"windFac1,omitzero"`
	// Wind factor for the second half of the route. This is the average wind factor
	// from the mid-time of the entire route to last Top of Descent in knots. A
	// positive value indicates a headwind, while a negative value indicates a
	// tailwind.
	WindFac2 param.Opt[float64] `json:"windFac2,omitzero"`
	// Average wind factor from Top of Climb to Top of Descent in knots. A positive
	// value indicates a headwind, while a negative value indicates a tailwind.
	WindFacAvg param.Opt[float64] `json:"windFacAvg,omitzero"`
	// The date and time the weather valid period ends in ISO 8601 UTC format, with
	// millisecond precision.
	WxValidEnd param.Opt[time.Time] `json:"wxValidEnd,omitzero" format:"date-time"`
	// The date and time the weather valid period begins in ISO 8601 UTC format, with
	// millisecond precision.
	WxValidStart param.Opt[time.Time] `json:"wxValidStart,omitzero" format:"date-time"`
	// Collection of air refueling events occurring on this flight.
	AirRefuelEvents []FlightplanUnvalidatedPublishParamsBodyAirRefuelEvent `json:"airRefuelEvents,omitzero"`
	// Array of Air Traffic Control (ATC) addresses.
	AtcAddresses []string `json:"atcAddresses,omitzero"`
	// Array of country codes for the countries overflown during this flight in ISO
	// 3166-1 Alpha-2 format.
	CountryCodes []string `json:"countryCodes,omitzero"`
	// Array of Extended Operations (ETOPS) adequate landing airfields that are within
	// the mission region.
	EtopsAirfields []string `json:"etopsAirfields,omitzero"`
	// Array of Extended Operations (ETOPS) alternate suitable landing airfields that
	// are within the mission region.
	EtopsAltAirfields []string `json:"etopsAltAirfields,omitzero"`
	// Collection of messages associated with this flight plan indicating the severity,
	// the point where the message was generated, the path (Primary, Alternate, etc.),
	// and the text of the message.
	FlightPlanMessages []FlightplanUnvalidatedPublishParamsBodyFlightPlanMessage `json:"flightPlanMessages,omitzero"`
	// Collection of point groups generated for this flight plan. Groups include point
	// sets for Extended Operations (ETOPS), Critical Fuel Point, and Equal Time Point
	// (ETP).
	FlightPlanPointGroups []FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroup `json:"flightPlanPointGroups,omitzero"`
	// Collection of waypoints associated with this flight plan.
	FlightPlanWaypoints []FlightplanUnvalidatedPublishParamsBodyFlightPlanWaypoint `json:"flightPlanWaypoints,omitzero"`
	paramObj
}

func (r FlightplanUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FlightplanUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of air refueling events occurring on this flight.
type FlightplanUnvalidatedPublishParamsBodyAirRefuelEvent struct {
	// Additional degrade for air refueling, cumulative with fuelDegrade field percent.
	ArDegrade param.Opt[float64] `json:"arDegrade,omitzero"`
	// Fuel onloaded (use positive numbers) or fuel offloaded (use negative numbers) in
	// pounds.
	ArExchangedFuel param.Opt[float64] `json:"arExchangedFuel,omitzero"`
	// The number of this air refueling event within the flight plan.
	ArNum param.Opt[int64] `json:"arNum,omitzero"`
	// Fuel required to fly from air refueling exit point to air refueling divert
	// alternate airfield in pounds.
	DivertFuel param.Opt[float64] `json:"divertFuel,omitzero"`
	// Fuel remaining at the air refueling exit in pounds.
	ExitFuel param.Opt[float64] `json:"exitFuel,omitzero"`
	paramObj
}

func (r FlightplanUnvalidatedPublishParamsBodyAirRefuelEvent) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUnvalidatedPublishParamsBodyAirRefuelEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUnvalidatedPublishParamsBodyAirRefuelEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of messages associated with this flight plan indicating the severity,
// the point where the message was generated, the path (Primary, Alternate, etc.),
// and the text of the message.
type FlightplanUnvalidatedPublishParamsBodyFlightPlanMessage struct {
	// The text of the message.
	MsgText param.Opt[string] `json:"msgText,omitzero"`
	// The flight path that generated the message (PRIMARY, ALTERNATE, etc.).
	RoutePath param.Opt[string] `json:"routePath,omitzero"`
	// The severity of the message.
	Severity param.Opt[string] `json:"severity,omitzero"`
	// The waypoint number for which the message was generated, or enter "PLAN" for a
	// message impacting the entire route.
	WpNum param.Opt[string] `json:"wpNum,omitzero"`
	paramObj
}

func (r FlightplanUnvalidatedPublishParamsBodyFlightPlanMessage) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUnvalidatedPublishParamsBodyFlightPlanMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUnvalidatedPublishParamsBodyFlightPlanMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of point groups generated for this flight plan. Groups include point
// sets for Extended Operations (ETOPS), Critical Fuel Point, and Equal Time Point
// (ETP).
type FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroup struct {
	// Average fuel flow at which the fuel was calculated in pounds per hour.
	AvgFuelFlow param.Opt[float64] `json:"avgFuelFlow,omitzero"`
	// Average wind factor from the Extended Operations (ETOPS) point to the recovery
	// airfield in knots.
	EtopsAvgWindFactor param.Opt[float64] `json:"etopsAvgWindFactor,omitzero"`
	// Distance from the Extended Operations (ETOPS) point to the recovery airfield in
	// nautical miles.
	EtopsDistance param.Opt[float64] `json:"etopsDistance,omitzero"`
	// Fuel required to fly from the Extended Operations (ETOPS) point to the recovery
	// airfield in pounds.
	EtopsReqFuel param.Opt[float64] `json:"etopsReqFuel,omitzero"`
	// Temperature deviation from the Extended Operations (ETOPS) point to the recovery
	// airfield in degrees Celsius.
	EtopsTempDev param.Opt[float64] `json:"etopsTempDev,omitzero"`
	// Time to fly from the Extended Operations (ETOPS) point to the recovery airfield
	// expressed in HH:MM format.
	EtopsTime param.Opt[string] `json:"etopsTime,omitzero"`
	// Total time from takeoff when the point is reached expressed in HH:MM format.
	FromTakeoffTime param.Opt[string] `json:"fromTakeoffTime,omitzero"`
	// Average wind factor from the Equal Time Point (ETP) to the first suitable
	// airfield in knots.
	FsafAvgWindFactor param.Opt[float64] `json:"fsafAvgWindFactor,omitzero"`
	// Distance from the Equal Time Point (ETP) to the first suitable airfield in
	// nautical miles.
	FsafDistance param.Opt[float64] `json:"fsafDistance,omitzero"`
	// Fuel required to fly from the Equal Time Point (ETP) to the first suitable
	// airfield in pounds.
	FsafReqFuel param.Opt[float64] `json:"fsafReqFuel,omitzero"`
	// Temperature deviation from the Equal Time Point (ETP) to the first suitable
	// airfield in degrees Celsius.
	FsafTempDev param.Opt[float64] `json:"fsafTempDev,omitzero"`
	// Time to fly from the Equal Time Point (ETP) to the first suitable airfield
	// expressed in HH:MM format.
	FsafTime param.Opt[string] `json:"fsafTime,omitzero"`
	// Flight level of the point at which the fuel was calculated in feet.
	FuelCalcAlt param.Opt[float64] `json:"fuelCalcAlt,omitzero"`
	// True airspeed at which the fuel was calculated in knots.
	FuelCalcSpd param.Opt[float64] `json:"fuelCalcSpd,omitzero"`
	// Average wind factor from the Equal Time Point (ETP) to the last suitable
	// airfield in knots.
	LsafAvgWindFactor param.Opt[float64] `json:"lsafAvgWindFactor,omitzero"`
	// Distance from the Equal Time Point (ETP) to the last suitable airfield in
	// nautical miles.
	LsafDistance param.Opt[float64] `json:"lsafDistance,omitzero"`
	// Name of the last suitable airfield, International Civil Aviation Organization
	// (ICAO) code preferred.
	LsafName param.Opt[string] `json:"lsafName,omitzero"`
	// Fuel required to fly from the Equal Time Point (ETP) to the last suitable
	// airfield in pounds.
	LsafReqFuel param.Opt[float64] `json:"lsafReqFuel,omitzero"`
	// Temperature deviation from the Equal Time Point (ETP) to the last suitable
	// airfield in degrees Celsius.
	LsafTempDev param.Opt[float64] `json:"lsafTempDev,omitzero"`
	// Time to fly from the Equal Time Point (ETP) to the last suitable airfield
	// expressed in HH:MM format.
	LsafTime param.Opt[string] `json:"lsafTime,omitzero"`
	// Amount of planned fuel on board when the point is reached in pounds.
	PlannedFuel param.Opt[float64] `json:"plannedFuel,omitzero"`
	// Name of the point group, usually Extended Operations (ETOPS), Critical Fuel
	// Point, and Equal Time Point (ETP) sections.
	PointGroupName param.Opt[string] `json:"pointGroupName,omitzero"`
	// Specifies which Point Group case requires the most fuel.
	WorstFuelCase param.Opt[string] `json:"worstFuelCase,omitzero"`
	// Array of point data for this Point Group.
	FlightPlanPoints []FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroupFlightPlanPoint `json:"flightPlanPoints,omitzero"`
	paramObj
}

func (r FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroup) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Array of point data for this Point Group.
type FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroupFlightPlanPoint struct {
	// Estimated Time of Arrival (ETA) at this point in ISO 8601 UTC format, with
	// millisecond precision.
	FppEta param.Opt[time.Time] `json:"fppEta,omitzero" format:"date-time"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	FppLat param.Opt[float64] `json:"fppLat,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	FppLon param.Opt[float64] `json:"fppLon,omitzero"`
	// Fuel required at this point to execute an Equal Time Point (ETP) or Extended
	// Operations (ETOPS) plan in pounds.
	FppReqFuel param.Opt[float64] `json:"fppReqFuel,omitzero"`
	// Name of this point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	paramObj
}

func (r FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroupFlightPlanPoint) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroupFlightPlanPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroupFlightPlanPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of waypoints associated with this flight plan.
//
// The properties Type, WaypointName are required.
type FlightplanUnvalidatedPublishParamsBodyFlightPlanWaypoint struct {
	// Points are designated by type as either a comment point or a waypoint. A comment
	// point conveys important information about the point for pilots but is not
	// entered into a flight management system. A waypoint is a point that is entered
	// into a flight management system and/or filed with Air Traffic Control.
	Type string `json:"type,required"`
	// Name of the point. The name of a comment point identifies important information
	// about that point, e.g. Top of Climb. The name of a waypoint identifies the
	// location of that point.
	WaypointName string `json:"waypointName,required"`
	// The air-to-air Tactical Air Navigation (TACAN) channels used by the
	// receiver/tanker during air refueling.
	AaTacanChannel param.Opt[string] `json:"aaTacanChannel,omitzero"`
	// The air distance of this leg in nautical miles.
	AirDistance param.Opt[float64] `json:"airDistance,omitzero"`
	// The flight path flown for this leg.
	Airway param.Opt[string] `json:"airway,omitzero"`
	// Altitude of a level, point, or object measured in feet above mean sea level.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// The ID of the air refueling track/anchor or fixed track.
	ArID param.Opt[string] `json:"arId,omitzero"`
	// Point identifying an air refueling track/anchor or fixed track.
	Arpt param.Opt[string] `json:"arpt,omitzero"`
	// Actual Time of Arrival (ATA) at this waypoint in ISO 8601 UTC format, with
	// millisecond precision.
	Ata param.Opt[time.Time] `json:"ata,omitzero" format:"date-time"`
	// The average calibrated airspeed (CAS) for this leg in knots.
	AvgCalAirspeed param.Opt[float64] `json:"avgCalAirspeed,omitzero"`
	// The average drift angle for this leg in degrees from true north.
	AvgDriftAng param.Opt[float64] `json:"avgDriftAng,omitzero"`
	// The average ground speed for this leg in knots.
	AvgGroundSpeed param.Opt[float64] `json:"avgGroundSpeed,omitzero"`
	// The average true airspeed (TAS) for this leg in knots.
	AvgTrueAirspeed param.Opt[float64] `json:"avgTrueAirspeed,omitzero"`
	// The average wind direction for this leg in degrees from true north.
	AvgWindDir param.Opt[float64] `json:"avgWindDir,omitzero"`
	// The average wind speed for this leg in knots.
	AvgWindSpeed param.Opt[float64] `json:"avgWindSpeed,omitzero"`
	// The day low level altitude in feet above sea level for the leg ending at this
	// waypoint.
	DayLowAlt param.Opt[float64] `json:"dayLowAlt,omitzero"`
	// Estimated Time of Arrival (ETA) at this waypoint in ISO 8601 UTC format, with
	// millisecond precision.
	Eta param.Opt[time.Time] `json:"eta,omitzero" format:"date-time"`
	// The amount of fuel onloaded or offloaded at this waypoint in pounds (negative
	// value for offload).
	ExchangedFuel param.Opt[float64] `json:"exchangedFuel,omitzero"`
	// The leg fuel flow in pounds per hour.
	FuelFlow param.Opt[float64] `json:"fuelFlow,omitzero"`
	// The icing intensity classification for this flight (LIGHT, MODERATE, etc).
	IceCat param.Opt[string] `json:"iceCat,omitzero"`
	// WGS84 latitude of the point location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The planned alternate leg based on user-defined constraints, International Civil
	// Aviation Organization (ICAO) code preferred.
	LegAlternate param.Opt[string] `json:"legAlternate,omitzero"`
	// The percent degrade due to drag for this aircraft for this leg.
	LegDragIndex param.Opt[float64] `json:"legDragIndex,omitzero"`
	// The fuel degrade percentage used for this leg.
	LegFuelDegrade param.Opt[float64] `json:"legFuelDegrade,omitzero"`
	// The average Mach speed for this leg.
	LegMach param.Opt[float64] `json:"legMach,omitzero"`
	// The mission index value for this leg. The mission index is the ratio of
	// time-related cost of aircraft operation to the cost of fuel.
	LegMsnIndex param.Opt[float64] `json:"legMsnIndex,omitzero"`
	// The wind factor for this leg in knots. A positive value indicates a headwind,
	// while a negative value indicates a tailwind.
	LegWindFac param.Opt[float64] `json:"legWindFac,omitzero"`
	// WGS84 longitude of the point location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The magnetic course at leg midpoint in degrees from true north.
	MagCourse param.Opt[float64] `json:"magCourse,omitzero"`
	// The magnetic heading at leg midpoint in degrees from true north.
	MagHeading param.Opt[float64] `json:"magHeading,omitzero"`
	// The magnetic variation for this leg in degrees.
	MagVar param.Opt[float64] `json:"magVar,omitzero"`
	// Navigational Aid (NAVAID) identification code.
	Navaid param.Opt[string] `json:"navaid,omitzero"`
	// The night low level altitude in feet above sea level for the leg ending at this
	// waypoint.
	NightLowAlt param.Opt[float64] `json:"nightLowAlt,omitzero"`
	// The night vision goggle low level altitude in feet above sea level for the leg
	// ending at this waypoint.
	NvgLowAlt param.Opt[float64] `json:"nvgLowAlt,omitzero"`
	// The wind direction at this specific point in degrees from true north.
	PointWindDir param.Opt[float64] `json:"pointWindDir,omitzero"`
	// The wind velocity at this specific point in knots.
	PointWindSpeed param.Opt[float64] `json:"pointWindSpeed,omitzero"`
	// The primary UHF radio frequency used for the air refueling track or anchor in
	// megahertz.
	PriFreq param.Opt[float64] `json:"priFreq,omitzero"`
	// The secondary UHF radio frequency used for the air refueling track or anchor in
	// megahertz.
	SecFreq param.Opt[float64] `json:"secFreq,omitzero"`
	// Tactical Air Navigation (TACAN) channel for the Navigational Aid (NAVAID).
	TacanChannel param.Opt[string] `json:"tacanChannel,omitzero"`
	// Average temperature deviation from standard day profile for this leg in degrees
	// Celsius.
	TempDev param.Opt[float64] `json:"tempDev,omitzero"`
	// The thunderstorm intensity classification for this flight (LIGHT, MODERATE,
	// etc).
	ThunderCat param.Opt[string] `json:"thunderCat,omitzero"`
	// The total air distance to this waypoint in nautical miles.
	TotalAirDistance param.Opt[float64] `json:"totalAirDistance,omitzero"`
	// The total distance flown to this waypoint calculated from point of departure in
	// nautical miles.
	TotalFlownDistance param.Opt[float64] `json:"totalFlownDistance,omitzero"`
	// The total distance remaining from this waypoint to the point of arrival in
	// nautical miles.
	TotalRemDistance param.Opt[float64] `json:"totalRemDistance,omitzero"`
	// The total fuel remaining at this waypoint in pounds.
	TotalRemFuel param.Opt[float64] `json:"totalRemFuel,omitzero"`
	// The total time accumulated from takeoff to this waypoint expressed as HH:MM.
	TotalTime param.Opt[string] `json:"totalTime,omitzero"`
	// The total time remaining from this waypoint to the point of arrival expressed as
	// HH:MM.
	TotalTimeRem param.Opt[string] `json:"totalTimeRem,omitzero"`
	// The total fuel used to this waypoint from point of departure in pounds.
	TotalUsedFuel param.Opt[float64] `json:"totalUsedFuel,omitzero"`
	// The total weight of the aircraft at this waypoint in pounds.
	TotalWeight param.Opt[float64] `json:"totalWeight,omitzero"`
	// The true course at leg midpoint in degrees from true north.
	TrueCourse param.Opt[float64] `json:"trueCourse,omitzero"`
	// The turbulence intensity classification for this flight (LIGHT, MODERATE, etc).
	TurbCat param.Opt[string] `json:"turbCat,omitzero"`
	// VHF Omni-directional Range (VOR) frequency for the Navigational Aid (NAVAID) in
	// megahertz.
	VorFreq param.Opt[float64] `json:"vorFreq,omitzero"`
	// The waypoint number on the route. Comment points do not get a waypoint number.
	WaypointNum param.Opt[int64] `json:"waypointNum,omitzero"`
	// The zone/leg distance flown in nautical miles.
	ZoneDistance param.Opt[float64] `json:"zoneDistance,omitzero"`
	// The amount of fuel used on this zone/leg in pounds.
	ZoneFuel param.Opt[float64] `json:"zoneFuel,omitzero"`
	// The time to fly this zone/leg in minutes.
	ZoneTime param.Opt[float64] `json:"zoneTime,omitzero"`
	paramObj
}

func (r FlightplanUnvalidatedPublishParamsBodyFlightPlanWaypoint) MarshalJSON() (data []byte, err error) {
	type shadow FlightplanUnvalidatedPublishParamsBodyFlightPlanWaypoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlightplanUnvalidatedPublishParamsBodyFlightPlanWaypoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
