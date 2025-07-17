// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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
)

// AirOperationAircraftSortieService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirOperationAircraftSortieService] method instead.
type AirOperationAircraftSortieService struct {
	Options []option.RequestOption
}

// NewAirOperationAircraftSortieService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAirOperationAircraftSortieService(opts ...option.RequestOption) (r AirOperationAircraftSortieService) {
	r = AirOperationAircraftSortieService{}
	r.Options = opts
	return
}

// Service operation to take a single AircraftSortie as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *AirOperationAircraftSortieService) New(ctx context.Context, body AirOperationAircraftSortieNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aircraftsortie"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirOperationAircraftSortieService) List(ctx context.Context, query AirOperationAircraftSortieListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AircraftsortieAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/aircraftsortie"
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
func (r *AirOperationAircraftSortieService) ListAutoPaging(ctx context.Context, query AirOperationAircraftSortieListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AircraftsortieAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirOperationAircraftSortieService) Count(ctx context.Context, query AirOperationAircraftSortieCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/aircraftsortie/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// AircraftSorties as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *AirOperationAircraftSortieService) NewBulk(ctx context.Context, body AirOperationAircraftSortieNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aircraftsortie/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirOperationAircraftSortieService) HistoryAodr(ctx context.Context, query AirOperationAircraftSortieHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aircraftsortie/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirOperationAircraftSortieService) HistoryCount(ctx context.Context, query AirOperationAircraftSortieHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/aircraftsortie/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirOperationAircraftSortieService) HistoryQuery(ctx context.Context, query AirOperationAircraftSortieHistoryQueryParams, opts ...option.RequestOption) (res *[]AircraftsortieFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/aircraftsortie/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information related to the planning, load, status, and deployment or dispatch of
// one aircraft to carry out a mission.
type AircraftsortieAbridged struct {
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
	DataMode AircraftsortieAbridgedDataMode `json:"dataMode,required"`
	// The scheduled time that the Aircraft sortie is planned to depart, in ISO 8601
	// UTC format with millisecond precision.
	PlannedDepTime time.Time `json:"plannedDepTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The actual arrival time, in ISO 8601 UTC format with millisecond precision.
	ActualArrTime time.Time `json:"actualArrTime" format:"date-time"`
	// The actual time the Aircraft comes to a complete stop in its parking position,
	// in ISO 8601 UTC format with millisecond precision.
	ActualBlockInTime time.Time `json:"actualBlockInTime" format:"date-time"`
	// The actual time the Aircraft begins to taxi from its parking position, in ISO
	// 8601 UTC format with millisecond precision.
	ActualBlockOutTime time.Time `json:"actualBlockOutTime" format:"date-time"`
	// The actual departure time, in ISO 8601 UTC format.
	ActualDepTime time.Time `json:"actualDepTime" format:"date-time"`
	// The Automatic Dependent Surveillance-Broadcast (ADS-B) device identifier.
	AircraftAdsb string `json:"aircraftADSB"`
	// Alternate Aircraft Identifier provided by source.
	AircraftAltID string `json:"aircraftAltId"`
	// Aircraft event text.
	AircraftEvent string `json:"aircraftEvent"`
	// The aircraft Model Design Series designation assigned to this sortie.
	AircraftMds string `json:"aircraftMDS"`
	// Remarks concerning the aircraft.
	AircraftRemarks string `json:"aircraftRemarks"`
	// The amount of time allowed between launch order and takeoff, in seconds.
	AlertStatus int64 `json:"alertStatus"`
	// The Alert Status code.
	AlertStatusCode string `json:"alertStatusCode"`
	// The Air Mobility Command (AMC) mission number of the sortie.
	AmcMsnNum string `json:"amcMsnNum"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	AmcMsnType string `json:"amcMsnType"`
	// The arrival Federal Aviation Administration (FAA) code of this sortie.
	ArrFaa string `json:"arrFAA"`
	// The arrival International Aviation Transport Association (IATA) code of this
	// sortie.
	ArrIata string `json:"arrIATA"`
	// The arrival International Civil Aviation Organization (ICAO) of this sortie.
	ArrIcao string `json:"arrICAO"`
	// The itinerary identifier of the arrival location.
	ArrItinerary int64 `json:"arrItinerary"`
	// Purpose code at the arrival location of this sortie.
	ArrPurposeCode string `json:"arrPurposeCode"`
	// The call sign assigned to the aircraft on this sortie.
	CallSign string `json:"callSign"`
	// Description of the cargo configuration (e.g. C-1, C-2, C-3, DV-1, DV-2, AE-1,
	// etc.) currently on board the aircraft. Configuration meanings are determined by
	// the data source.
	CargoConfig string `json:"cargoConfig"`
	// The last name of the aircraft commander.
	CommanderName string `json:"commanderName"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The current state of this sortie.
	CurrentState string `json:"currentState"`
	// The primary delay code.
	DelayCode string `json:"delayCode"`
	// The departure Federal Aviation Administration (FAA) code of this sortie.
	DepFaa string `json:"depFAA"`
	// The departure International Aviation Transport Association (IATA) code of this
	// sortie.
	DepIata string `json:"depIATA"`
	// The departure International Civil Aviation Organization (ICAO) of this sortie.
	DepIcao string `json:"depICAO"`
	// The itinerary identifier of the departure location.
	DepItinerary int64 `json:"depItinerary"`
	// Purpose code at the departure location of this sortie.
	DepPurposeCode string `json:"depPurposeCode"`
	// Due home date by which the aircraft must return to its home station, in ISO 8601
	// UTC format with millisecond precision.
	Dhd time.Time `json:"dhd" format:"date-time"`
	// Reason the aircraft must return to home station by its due home date.
	DhdReason string `json:"dhdReason"`
	// The current estimated time that the Aircraft is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime time.Time `json:"estArrTime" format:"date-time"`
	// The estimated time the Aircraft will come to a complete stop in its parking
	// position, in ISO 8601 UTC format with millisecond precision.
	EstBlockInTime time.Time `json:"estBlockInTime" format:"date-time"`
	// The estimated time the Aircraft will begin to taxi from its parking position, in
	// ISO 8601 UTC format with millisecond precision.
	EstBlockOutTime time.Time `json:"estBlockOutTime" format:"date-time"`
	// The current estimated time that the Aircraft is planned to depart, in ISO 8601
	// UTC format with millisecond precision.
	EstDepTime time.Time `json:"estDepTime" format:"date-time"`
	// Name of the uploaded PDF.
	Filename string `json:"filename"`
	// Size of the supporting PDF, in bytes.
	Filesize int64 `json:"filesize"`
	// The planned flight time for this sortie, in minutes.
	FlightTime float64 `json:"flightTime"`
	// Desk phone number of the flight manager assigned to the sortie. Null when no
	// flight manager is assigned.
	FmDeskNum string `json:"fmDeskNum"`
	// Last name of the flight manager assigned to the sortie. Null when no flight
	// manager is assigned.
	FmName string `json:"fmName"`
	// Mass of fuel required for this leg of the sortie, in kilograms.
	FuelReq float64 `json:"fuelReq"`
	// Scheduled ground time, in minutes.
	GndTime float64 `json:"gndTime"`
	// Unique identifier of the aircraft.
	IDAircraft string `json:"idAircraft"`
	// The unique identifier of the mission to which this sortie is assigned.
	IDMission string `json:"idMission"`
	// Joint Chiefs of Staff priority of this sortie.
	JcsPriority string `json:"jcsPriority"`
	// The leg number of this sortie.
	LegNum int64 `json:"legNum"`
	// The external system line number of this sortie.
	LineNumber int64 `json:"lineNumber"`
	// The mission ID according to the source system.
	MissionID string `json:"missionId"`
	// Time the associated mission data was last updated in relation to the aircraft
	// assignment, in ISO 8601 UTC format with millisecond precision. If this time is
	// coming from an external system, it may not sync with the latest mission time
	// associated to this record.
	MissionUpdate time.Time `json:"missionUpdate" format:"date-time"`
	// Remarks concerning the sortie objective.
	ObjectiveRemarks string `json:"objectiveRemarks"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The sortie identifier provided by the originating source.
	OrigSortieID string `json:"origSortieId"`
	// Liquid oxygen onboard the aircraft for the crew compartment, in liters.
	OxyOnCrew float64 `json:"oxyOnCrew"`
	// Liquid oxygen onboard the aircraft for the troop compartment, in liters.
	OxyOnPax float64 `json:"oxyOnPax"`
	// Liquid oxygen required on the aircraft for the crew compartment, in liters.
	OxyReqCrew float64 `json:"oxyReqCrew"`
	// Liquid oxygen required on the aircraft for the troop compartment, in liters.
	OxyReqPax float64 `json:"oxyReqPax"`
	// The status of the supporting document.
	//
	// Any of "PUBLISHED", "DELETED", "UPDATED", "READ".
	PaperStatus AircraftsortieAbridgedPaperStatus `json:"paperStatus"`
	// The version number of the crew paper.
	PapersVersion string `json:"papersVersion"`
	// The POI parking location.
	ParkingLoc string `json:"parkingLoc"`
	// The number of passengers tasked for this sortie.
	Passengers int64 `json:"passengers"`
	// The scheduled time that the Aircraft sortie is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	PlannedArrTime time.Time `json:"plannedArrTime" format:"date-time"`
	// The prior permission required (PPR) status.
	//
	// Any of "NOT REQUIRED", "REQUIRED NOT REQUESTED", "GRANTED", "PENDING".
	PprStatus AircraftsortieAbridgedPprStatus `json:"pprStatus"`
	// The planned primary Standard Conventional Load of the aircraft for this sortie.
	PrimaryScl string `json:"primarySCL"`
	// When crew papers are associated to this sortie, the system updates this value.
	// This field is the URI location in the document repository of that raw file. To
	// download the raw file, prepend https://udl-hostname/scs/download?id= to this
	// field's value.
	RawFileUri string `json:"rawFileURI"`
	// Aircraft configuration required for the mission.
	ReqConfig string `json:"reqConfig"`
	// Remarks concerning the results of this sortie.
	ResultRemarks string `json:"resultRemarks"`
	// Type of Ravens required for this sortie (N - None, R - Raven (Security Team)
	// required, C6 - Consider ravens (Ground time over 6 hours), R6 - Ravens required
	// (Ground time over 6 hours)).
	//
	// Any of "N", "R", "C6", "R6".
	RvnReq AircraftsortieAbridgedRvnReq `json:"rvnReq"`
	// Remarks concerning the schedule.
	ScheduleRemarks string `json:"scheduleRemarks"`
	// The planned secondary Standard Conventional Load of the aircraft for this
	// sortie.
	SecondaryScl string `json:"secondarySCL"`
	// Indicates the group responsible for recording the completion time of the next
	// event in the sequence of events assigned to this sortie (e.g. OPS - Operations,
	// MX - Maintenance, TR - Transportation, etc.).
	Soe string `json:"soe"`
	// The scheduled UTC date for this sortie, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	SortieDate time.Time `json:"sortieDate" format:"date"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The tail number of the aircraft assigned to this sortie.
	TailNumber string `json:"tailNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		PlannedDepTime        respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ActualArrTime         respjson.Field
		ActualBlockInTime     respjson.Field
		ActualBlockOutTime    respjson.Field
		ActualDepTime         respjson.Field
		AircraftAdsb          respjson.Field
		AircraftAltID         respjson.Field
		AircraftEvent         respjson.Field
		AircraftMds           respjson.Field
		AircraftRemarks       respjson.Field
		AlertStatus           respjson.Field
		AlertStatusCode       respjson.Field
		AmcMsnNum             respjson.Field
		AmcMsnType            respjson.Field
		ArrFaa                respjson.Field
		ArrIata               respjson.Field
		ArrIcao               respjson.Field
		ArrItinerary          respjson.Field
		ArrPurposeCode        respjson.Field
		CallSign              respjson.Field
		CargoConfig           respjson.Field
		CommanderName         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CurrentState          respjson.Field
		DelayCode             respjson.Field
		DepFaa                respjson.Field
		DepIata               respjson.Field
		DepIcao               respjson.Field
		DepItinerary          respjson.Field
		DepPurposeCode        respjson.Field
		Dhd                   respjson.Field
		DhdReason             respjson.Field
		EstArrTime            respjson.Field
		EstBlockInTime        respjson.Field
		EstBlockOutTime       respjson.Field
		EstDepTime            respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		FlightTime            respjson.Field
		FmDeskNum             respjson.Field
		FmName                respjson.Field
		FuelReq               respjson.Field
		GndTime               respjson.Field
		IDAircraft            respjson.Field
		IDMission             respjson.Field
		JcsPriority           respjson.Field
		LegNum                respjson.Field
		LineNumber            respjson.Field
		MissionID             respjson.Field
		MissionUpdate         respjson.Field
		ObjectiveRemarks      respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSortieID          respjson.Field
		OxyOnCrew             respjson.Field
		OxyOnPax              respjson.Field
		OxyReqCrew            respjson.Field
		OxyReqPax             respjson.Field
		PaperStatus           respjson.Field
		PapersVersion         respjson.Field
		ParkingLoc            respjson.Field
		Passengers            respjson.Field
		PlannedArrTime        respjson.Field
		PprStatus             respjson.Field
		PrimaryScl            respjson.Field
		RawFileUri            respjson.Field
		ReqConfig             respjson.Field
		ResultRemarks         respjson.Field
		RvnReq                respjson.Field
		ScheduleRemarks       respjson.Field
		SecondaryScl          respjson.Field
		Soe                   respjson.Field
		SortieDate            respjson.Field
		SourceDl              respjson.Field
		TailNumber            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AircraftsortieAbridged) RawJSON() string { return r.JSON.raw }
func (r *AircraftsortieAbridged) UnmarshalJSON(data []byte) error {
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
type AircraftsortieAbridgedDataMode string

const (
	AircraftsortieAbridgedDataModeReal      AircraftsortieAbridgedDataMode = "REAL"
	AircraftsortieAbridgedDataModeTest      AircraftsortieAbridgedDataMode = "TEST"
	AircraftsortieAbridgedDataModeSimulated AircraftsortieAbridgedDataMode = "SIMULATED"
	AircraftsortieAbridgedDataModeExercise  AircraftsortieAbridgedDataMode = "EXERCISE"
)

// The status of the supporting document.
type AircraftsortieAbridgedPaperStatus string

const (
	AircraftsortieAbridgedPaperStatusPublished AircraftsortieAbridgedPaperStatus = "PUBLISHED"
	AircraftsortieAbridgedPaperStatusDeleted   AircraftsortieAbridgedPaperStatus = "DELETED"
	AircraftsortieAbridgedPaperStatusUpdated   AircraftsortieAbridgedPaperStatus = "UPDATED"
	AircraftsortieAbridgedPaperStatusRead      AircraftsortieAbridgedPaperStatus = "READ"
)

// The prior permission required (PPR) status.
type AircraftsortieAbridgedPprStatus string

const (
	AircraftsortieAbridgedPprStatusNotRequired          AircraftsortieAbridgedPprStatus = "NOT REQUIRED"
	AircraftsortieAbridgedPprStatusRequiredNotRequested AircraftsortieAbridgedPprStatus = "REQUIRED NOT REQUESTED"
	AircraftsortieAbridgedPprStatusGranted              AircraftsortieAbridgedPprStatus = "GRANTED"
	AircraftsortieAbridgedPprStatusPending              AircraftsortieAbridgedPprStatus = "PENDING"
)

// Type of Ravens required for this sortie (N - None, R - Raven (Security Team)
// required, C6 - Consider ravens (Ground time over 6 hours), R6 - Ravens required
// (Ground time over 6 hours)).
type AircraftsortieAbridgedRvnReq string

const (
	AircraftsortieAbridgedRvnReqN  AircraftsortieAbridgedRvnReq = "N"
	AircraftsortieAbridgedRvnReqR  AircraftsortieAbridgedRvnReq = "R"
	AircraftsortieAbridgedRvnReqC6 AircraftsortieAbridgedRvnReq = "C6"
	AircraftsortieAbridgedRvnReqR6 AircraftsortieAbridgedRvnReq = "R6"
)

// Information related to the planning, load, status, and deployment or dispatch of
// one aircraft to carry out a mission.
type AircraftsortieFull struct {
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
	DataMode AircraftsortieFullDataMode `json:"dataMode,required"`
	// The scheduled time that the Aircraft sortie is planned to depart, in ISO 8601
	// UTC format with millisecond precision.
	PlannedDepTime time.Time `json:"plannedDepTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The actual arrival time, in ISO 8601 UTC format with millisecond precision.
	ActualArrTime time.Time `json:"actualArrTime" format:"date-time"`
	// The actual time the Aircraft comes to a complete stop in its parking position,
	// in ISO 8601 UTC format with millisecond precision.
	ActualBlockInTime time.Time `json:"actualBlockInTime" format:"date-time"`
	// The actual time the Aircraft begins to taxi from its parking position, in ISO
	// 8601 UTC format with millisecond precision.
	ActualBlockOutTime time.Time `json:"actualBlockOutTime" format:"date-time"`
	// The actual departure time, in ISO 8601 UTC format.
	ActualDepTime time.Time `json:"actualDepTime" format:"date-time"`
	// The Automatic Dependent Surveillance-Broadcast (ADS-B) device identifier.
	AircraftAdsb string `json:"aircraftADSB"`
	// Alternate Aircraft Identifier provided by source.
	AircraftAltID string `json:"aircraftAltId"`
	// Aircraft event text.
	AircraftEvent string `json:"aircraftEvent"`
	// The aircraft Model Design Series designation assigned to this sortie.
	AircraftMds string `json:"aircraftMDS"`
	// Remarks concerning the aircraft.
	AircraftRemarks string `json:"aircraftRemarks"`
	// The amount of time allowed between launch order and takeoff, in seconds.
	AlertStatus int64 `json:"alertStatus"`
	// The Alert Status code.
	AlertStatusCode string `json:"alertStatusCode"`
	// The Air Mobility Command (AMC) mission number of the sortie.
	AmcMsnNum string `json:"amcMsnNum"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	AmcMsnType string `json:"amcMsnType"`
	// The arrival Federal Aviation Administration (FAA) code of this sortie.
	ArrFaa string `json:"arrFAA"`
	// The arrival International Aviation Transport Association (IATA) code of this
	// sortie.
	ArrIata string `json:"arrIATA"`
	// The arrival International Civil Aviation Organization (ICAO) of this sortie.
	ArrIcao string `json:"arrICAO"`
	// The itinerary identifier of the arrival location.
	ArrItinerary int64 `json:"arrItinerary"`
	// Purpose code at the arrival location of this sortie.
	ArrPurposeCode string `json:"arrPurposeCode"`
	// The call sign assigned to the aircraft on this sortie.
	CallSign string `json:"callSign"`
	// Description of the cargo configuration (e.g. C-1, C-2, C-3, DV-1, DV-2, AE-1,
	// etc.) currently on board the aircraft. Configuration meanings are determined by
	// the data source.
	CargoConfig string `json:"cargoConfig"`
	// The last name of the aircraft commander.
	CommanderName string `json:"commanderName"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Crew Services.
	Crew CrewFull `json:"crew"`
	// The current state of this sortie.
	CurrentState string `json:"currentState"`
	// The primary delay code.
	DelayCode string `json:"delayCode"`
	// The departure Federal Aviation Administration (FAA) code of this sortie.
	DepFaa string `json:"depFAA"`
	// The departure International Aviation Transport Association (IATA) code of this
	// sortie.
	DepIata string `json:"depIATA"`
	// The departure International Civil Aviation Organization (ICAO) of this sortie.
	DepIcao string `json:"depICAO"`
	// The itinerary identifier of the departure location.
	DepItinerary int64 `json:"depItinerary"`
	// Purpose code at the departure location of this sortie.
	DepPurposeCode string `json:"depPurposeCode"`
	// Due home date by which the aircraft must return to its home station, in ISO 8601
	// UTC format with millisecond precision.
	Dhd time.Time `json:"dhd" format:"date-time"`
	// Reason the aircraft must return to home station by its due home date.
	DhdReason string `json:"dhdReason"`
	// The current estimated time that the Aircraft is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime time.Time `json:"estArrTime" format:"date-time"`
	// The estimated time the Aircraft will come to a complete stop in its parking
	// position, in ISO 8601 UTC format with millisecond precision.
	EstBlockInTime time.Time `json:"estBlockInTime" format:"date-time"`
	// The estimated time the Aircraft will begin to taxi from its parking position, in
	// ISO 8601 UTC format with millisecond precision.
	EstBlockOutTime time.Time `json:"estBlockOutTime" format:"date-time"`
	// The current estimated time that the Aircraft is planned to depart, in ISO 8601
	// UTC format with millisecond precision.
	EstDepTime time.Time `json:"estDepTime" format:"date-time"`
	// Name of the uploaded PDF.
	Filename string `json:"filename"`
	// Size of the supporting PDF, in bytes.
	Filesize int64 `json:"filesize"`
	// The planned flight time for this sortie, in minutes.
	FlightTime float64 `json:"flightTime"`
	// Desk phone number of the flight manager assigned to the sortie. Null when no
	// flight manager is assigned.
	FmDeskNum string `json:"fmDeskNum"`
	// Last name of the flight manager assigned to the sortie. Null when no flight
	// manager is assigned.
	FmName string `json:"fmName"`
	// Mass of fuel required for this leg of the sortie, in kilograms.
	FuelReq float64 `json:"fuelReq"`
	// Scheduled ground time, in minutes.
	GndTime float64 `json:"gndTime"`
	// Unique identifier of the aircraft.
	IDAircraft string `json:"idAircraft"`
	// The unique identifier of the mission to which this sortie is assigned.
	IDMission string `json:"idMission"`
	// Joint Chiefs of Staff priority of this sortie.
	JcsPriority string `json:"jcsPriority"`
	// The leg number of this sortie.
	LegNum int64 `json:"legNum"`
	// The external system line number of this sortie.
	LineNumber int64 `json:"lineNumber"`
	// The mission ID according to the source system.
	MissionID string `json:"missionId"`
	// Time the associated mission data was last updated in relation to the aircraft
	// assignment, in ISO 8601 UTC format with millisecond precision. If this time is
	// coming from an external system, it may not sync with the latest mission time
	// associated to this record.
	MissionUpdate time.Time `json:"missionUpdate" format:"date-time"`
	// Remarks concerning the sortie objective.
	ObjectiveRemarks string `json:"objectiveRemarks"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The sortie identifier provided by the originating source.
	OrigSortieID string `json:"origSortieId"`
	// Liquid oxygen onboard the aircraft for the crew compartment, in liters.
	OxyOnCrew float64 `json:"oxyOnCrew"`
	// Liquid oxygen onboard the aircraft for the troop compartment, in liters.
	OxyOnPax float64 `json:"oxyOnPax"`
	// Liquid oxygen required on the aircraft for the crew compartment, in liters.
	OxyReqCrew float64 `json:"oxyReqCrew"`
	// Liquid oxygen required on the aircraft for the troop compartment, in liters.
	OxyReqPax float64 `json:"oxyReqPax"`
	// The status of the supporting document.
	//
	// Any of "PUBLISHED", "DELETED", "UPDATED", "READ".
	PaperStatus AircraftsortieFullPaperStatus `json:"paperStatus"`
	// The version number of the crew paper.
	PapersVersion string `json:"papersVersion"`
	// The POI parking location.
	ParkingLoc string `json:"parkingLoc"`
	// The number of passengers tasked for this sortie.
	Passengers int64 `json:"passengers"`
	// The scheduled time that the Aircraft sortie is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	PlannedArrTime time.Time `json:"plannedArrTime" format:"date-time"`
	// The prior permission required (PPR) status.
	//
	// Any of "NOT REQUIRED", "REQUIRED NOT REQUESTED", "GRANTED", "PENDING".
	PprStatus AircraftsortieFullPprStatus `json:"pprStatus"`
	// The planned primary Standard Conventional Load of the aircraft for this sortie.
	PrimaryScl string `json:"primarySCL"`
	// When crew papers are associated to this sortie, the system updates this value.
	// This field is the URI location in the document repository of that raw file. To
	// download the raw file, prepend https://udl-hostname/scs/download?id= to this
	// field's value.
	RawFileUri string `json:"rawFileURI"`
	// Aircraft configuration required for the mission.
	ReqConfig string `json:"reqConfig"`
	// Remarks concerning the results of this sortie.
	ResultRemarks string `json:"resultRemarks"`
	// Type of Ravens required for this sortie (N - None, R - Raven (Security Team)
	// required, C6 - Consider ravens (Ground time over 6 hours), R6 - Ravens required
	// (Ground time over 6 hours)).
	//
	// Any of "N", "R", "C6", "R6".
	RvnReq AircraftsortieFullRvnReq `json:"rvnReq"`
	// Remarks concerning the schedule.
	ScheduleRemarks string `json:"scheduleRemarks"`
	// The planned secondary Standard Conventional Load of the aircraft for this
	// sortie.
	SecondaryScl string `json:"secondarySCL"`
	// Indicates the group responsible for recording the completion time of the next
	// event in the sequence of events assigned to this sortie (e.g. OPS - Operations,
	// MX - Maintenance, TR - Transportation, etc.).
	Soe string `json:"soe"`
	// The scheduled UTC date for this sortie, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	SortieDate time.Time       `json:"sortieDate" format:"date"`
	SortiePpr  []SortiePprFull `json:"sortiePPR"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The tail number of the aircraft assigned to this sortie.
	TailNumber string `json:"tailNumber"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		PlannedDepTime        respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ActualArrTime         respjson.Field
		ActualBlockInTime     respjson.Field
		ActualBlockOutTime    respjson.Field
		ActualDepTime         respjson.Field
		AircraftAdsb          respjson.Field
		AircraftAltID         respjson.Field
		AircraftEvent         respjson.Field
		AircraftMds           respjson.Field
		AircraftRemarks       respjson.Field
		AlertStatus           respjson.Field
		AlertStatusCode       respjson.Field
		AmcMsnNum             respjson.Field
		AmcMsnType            respjson.Field
		ArrFaa                respjson.Field
		ArrIata               respjson.Field
		ArrIcao               respjson.Field
		ArrItinerary          respjson.Field
		ArrPurposeCode        respjson.Field
		CallSign              respjson.Field
		CargoConfig           respjson.Field
		CommanderName         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Crew                  respjson.Field
		CurrentState          respjson.Field
		DelayCode             respjson.Field
		DepFaa                respjson.Field
		DepIata               respjson.Field
		DepIcao               respjson.Field
		DepItinerary          respjson.Field
		DepPurposeCode        respjson.Field
		Dhd                   respjson.Field
		DhdReason             respjson.Field
		EstArrTime            respjson.Field
		EstBlockInTime        respjson.Field
		EstBlockOutTime       respjson.Field
		EstDepTime            respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		FlightTime            respjson.Field
		FmDeskNum             respjson.Field
		FmName                respjson.Field
		FuelReq               respjson.Field
		GndTime               respjson.Field
		IDAircraft            respjson.Field
		IDMission             respjson.Field
		JcsPriority           respjson.Field
		LegNum                respjson.Field
		LineNumber            respjson.Field
		MissionID             respjson.Field
		MissionUpdate         respjson.Field
		ObjectiveRemarks      respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSortieID          respjson.Field
		OxyOnCrew             respjson.Field
		OxyOnPax              respjson.Field
		OxyReqCrew            respjson.Field
		OxyReqPax             respjson.Field
		PaperStatus           respjson.Field
		PapersVersion         respjson.Field
		ParkingLoc            respjson.Field
		Passengers            respjson.Field
		PlannedArrTime        respjson.Field
		PprStatus             respjson.Field
		PrimaryScl            respjson.Field
		RawFileUri            respjson.Field
		ReqConfig             respjson.Field
		ResultRemarks         respjson.Field
		RvnReq                respjson.Field
		ScheduleRemarks       respjson.Field
		SecondaryScl          respjson.Field
		Soe                   respjson.Field
		SortieDate            respjson.Field
		SortiePpr             respjson.Field
		SourceDl              respjson.Field
		TailNumber            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AircraftsortieFull) RawJSON() string { return r.JSON.raw }
func (r *AircraftsortieFull) UnmarshalJSON(data []byte) error {
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
type AircraftsortieFullDataMode string

const (
	AircraftsortieFullDataModeReal      AircraftsortieFullDataMode = "REAL"
	AircraftsortieFullDataModeTest      AircraftsortieFullDataMode = "TEST"
	AircraftsortieFullDataModeSimulated AircraftsortieFullDataMode = "SIMULATED"
	AircraftsortieFullDataModeExercise  AircraftsortieFullDataMode = "EXERCISE"
)

// The status of the supporting document.
type AircraftsortieFullPaperStatus string

const (
	AircraftsortieFullPaperStatusPublished AircraftsortieFullPaperStatus = "PUBLISHED"
	AircraftsortieFullPaperStatusDeleted   AircraftsortieFullPaperStatus = "DELETED"
	AircraftsortieFullPaperStatusUpdated   AircraftsortieFullPaperStatus = "UPDATED"
	AircraftsortieFullPaperStatusRead      AircraftsortieFullPaperStatus = "READ"
)

// The prior permission required (PPR) status.
type AircraftsortieFullPprStatus string

const (
	AircraftsortieFullPprStatusNotRequired          AircraftsortieFullPprStatus = "NOT REQUIRED"
	AircraftsortieFullPprStatusRequiredNotRequested AircraftsortieFullPprStatus = "REQUIRED NOT REQUESTED"
	AircraftsortieFullPprStatusGranted              AircraftsortieFullPprStatus = "GRANTED"
	AircraftsortieFullPprStatusPending              AircraftsortieFullPprStatus = "PENDING"
)

// Type of Ravens required for this sortie (N - None, R - Raven (Security Team)
// required, C6 - Consider ravens (Ground time over 6 hours), R6 - Ravens required
// (Ground time over 6 hours)).
type AircraftsortieFullRvnReq string

const (
	AircraftsortieFullRvnReqN  AircraftsortieFullRvnReq = "N"
	AircraftsortieFullRvnReqR  AircraftsortieFullRvnReq = "R"
	AircraftsortieFullRvnReqC6 AircraftsortieFullRvnReq = "C6"
	AircraftsortieFullRvnReqR6 AircraftsortieFullRvnReq = "R6"
)

type AirOperationAircraftSortieNewParams struct {
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
	DataMode AirOperationAircraftSortieNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	PprStatus AirOperationAircraftSortieNewParamsPprStatus `json:"pprStatus,omitzero"`
	// Type of Ravens required for this sortie (N - None, R - Raven (Security Team)
	// required, C6 - Consider ravens (Ground time over 6 hours), R6 - Ravens required
	// (Ground time over 6 hours)).
	//
	// Any of "N", "R", "C6", "R6".
	RvnReq AirOperationAircraftSortieNewParamsRvnReq `json:"rvnReq,omitzero"`
	paramObj
}

func (r AirOperationAircraftSortieNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAircraftSortieNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAircraftSortieNewParams) UnmarshalJSON(data []byte) error {
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
type AirOperationAircraftSortieNewParamsDataMode string

const (
	AirOperationAircraftSortieNewParamsDataModeReal      AirOperationAircraftSortieNewParamsDataMode = "REAL"
	AirOperationAircraftSortieNewParamsDataModeTest      AirOperationAircraftSortieNewParamsDataMode = "TEST"
	AirOperationAircraftSortieNewParamsDataModeSimulated AirOperationAircraftSortieNewParamsDataMode = "SIMULATED"
	AirOperationAircraftSortieNewParamsDataModeExercise  AirOperationAircraftSortieNewParamsDataMode = "EXERCISE"
)

// The prior permission required (PPR) status.
type AirOperationAircraftSortieNewParamsPprStatus string

const (
	AirOperationAircraftSortieNewParamsPprStatusNotRequired          AirOperationAircraftSortieNewParamsPprStatus = "NOT REQUIRED"
	AirOperationAircraftSortieNewParamsPprStatusRequiredNotRequested AirOperationAircraftSortieNewParamsPprStatus = "REQUIRED NOT REQUESTED"
	AirOperationAircraftSortieNewParamsPprStatusGranted              AirOperationAircraftSortieNewParamsPprStatus = "GRANTED"
	AirOperationAircraftSortieNewParamsPprStatusPending              AirOperationAircraftSortieNewParamsPprStatus = "PENDING"
)

// Type of Ravens required for this sortie (N - None, R - Raven (Security Team)
// required, C6 - Consider ravens (Ground time over 6 hours), R6 - Ravens required
// (Ground time over 6 hours)).
type AirOperationAircraftSortieNewParamsRvnReq string

const (
	AirOperationAircraftSortieNewParamsRvnReqN  AirOperationAircraftSortieNewParamsRvnReq = "N"
	AirOperationAircraftSortieNewParamsRvnReqR  AirOperationAircraftSortieNewParamsRvnReq = "R"
	AirOperationAircraftSortieNewParamsRvnReqC6 AirOperationAircraftSortieNewParamsRvnReq = "C6"
	AirOperationAircraftSortieNewParamsRvnReqR6 AirOperationAircraftSortieNewParamsRvnReq = "R6"
)

type AirOperationAircraftSortieListParams struct {
	// The scheduled time that the Aircraft sortie is planned to depart, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	PlannedDepTime time.Time        `query:"plannedDepTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationAircraftSortieListParams]'s query parameters as
// `url.Values`.
func (r AirOperationAircraftSortieListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirOperationAircraftSortieCountParams struct {
	// The scheduled time that the Aircraft sortie is planned to depart, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	PlannedDepTime time.Time        `query:"plannedDepTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationAircraftSortieCountParams]'s query parameters
// as `url.Values`.
func (r AirOperationAircraftSortieCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirOperationAircraftSortieNewBulkParams struct {
	Body []AirOperationAircraftSortieNewBulkParamsBody
	paramObj
}

func (r AirOperationAircraftSortieNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *AirOperationAircraftSortieNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Information related to the planning, load, status, and deployment or dispatch of
// one aircraft to carry out a mission.
//
// The properties ClassificationMarking, DataMode, PlannedDepTime, Source are
// required.
type AirOperationAircraftSortieNewBulkParamsBody struct {
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
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
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
	// Name of the uploaded PDF.
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Size of the supporting PDF, in bytes.
	Filesize param.Opt[int64] `json:"filesize,omitzero"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
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
	// The version number of the crew paper.
	PapersVersion param.Opt[string] `json:"papersVersion,omitzero"`
	// The POI parking location.
	ParkingLoc param.Opt[string] `json:"parkingLoc,omitzero"`
	// The number of passengers tasked for this sortie.
	Passengers param.Opt[int64] `json:"passengers,omitzero"`
	// The scheduled time that the Aircraft sortie is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	PlannedArrTime param.Opt[time.Time] `json:"plannedArrTime,omitzero" format:"date-time"`
	// The planned primary Standard Conventional Load of the aircraft for this sortie.
	PrimaryScl param.Opt[string] `json:"primarySCL,omitzero"`
	// When crew papers are associated to this sortie, the system updates this value.
	// This field is the URI location in the document repository of that raw file. To
	// download the raw file, prepend https://udl-hostname/scs/download?id= to this
	// field's value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
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
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The tail number of the aircraft assigned to this sortie.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// The status of the supporting document.
	//
	// Any of "PUBLISHED", "DELETED", "UPDATED", "READ".
	PaperStatus string `json:"paperStatus,omitzero"`
	// The prior permission required (PPR) status.
	//
	// Any of "NOT REQUIRED", "REQUIRED NOT REQUESTED", "GRANTED", "PENDING".
	PprStatus string `json:"pprStatus,omitzero"`
	// Type of Ravens required for this sortie (N - None, R - Raven (Security Team)
	// required, C6 - Consider ravens (Ground time over 6 hours), R6 - Ravens required
	// (Ground time over 6 hours)).
	//
	// Any of "N", "R", "C6", "R6".
	RvnReq string `json:"rvnReq,omitzero"`
	paramObj
}

func (r AirOperationAircraftSortieNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAircraftSortieNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAircraftSortieNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AirOperationAircraftSortieNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[AirOperationAircraftSortieNewBulkParamsBody](
		"paperStatus", "PUBLISHED", "DELETED", "UPDATED", "READ",
	)
	apijson.RegisterFieldValidator[AirOperationAircraftSortieNewBulkParamsBody](
		"pprStatus", "NOT REQUIRED", "REQUIRED NOT REQUESTED", "GRANTED", "PENDING",
	)
	apijson.RegisterFieldValidator[AirOperationAircraftSortieNewBulkParamsBody](
		"rvnReq", "N", "R", "C6", "R6",
	)
}

type AirOperationAircraftSortieHistoryAodrParams struct {
	// The scheduled time that the Aircraft sortie is planned to depart, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	PlannedDepTime time.Time `query:"plannedDepTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// optional, notification method for the created file link. When omitted, EMAIL is
	// assumed. Current valid values are: EMAIL, SMS.
	Notification param.Opt[string] `query:"notification,omitzero" json:"-"`
	// optional, field delimiter when the created file is not JSON. Must be a single
	// character chosen from this set: (',', ';', ':', '|'). When omitted, "," is used.
	// It is strongly encouraged that your field delimiter be a character unlikely to
	// occur within the data.
	OutputDelimiter param.Opt[string] `query:"outputDelimiter,omitzero" json:"-"`
	// optional, output format for the file. When omitted, JSON is assumed. Current
	// valid values are: JSON and CSV.
	OutputFormat param.Opt[string] `query:"outputFormat,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationAircraftSortieHistoryAodrParams]'s query
// parameters as `url.Values`.
func (r AirOperationAircraftSortieHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirOperationAircraftSortieHistoryCountParams struct {
	// The scheduled time that the Aircraft sortie is planned to depart, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	PlannedDepTime time.Time        `query:"plannedDepTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationAircraftSortieHistoryCountParams]'s query
// parameters as `url.Values`.
func (r AirOperationAircraftSortieHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirOperationAircraftSortieHistoryQueryParams struct {
	// The scheduled time that the Aircraft sortie is planned to depart, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	PlannedDepTime time.Time `query:"plannedDepTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationAircraftSortieHistoryQueryParams]'s query
// parameters as `url.Values`.
func (r AirOperationAircraftSortieHistoryQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
