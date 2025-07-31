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

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// CrewService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCrewService] method instead.
type CrewService struct {
	Options []option.RequestOption
}

// NewCrewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCrewService(opts ...option.RequestOption) (r CrewService) {
	r = CrewService{}
	r.Options = opts
	return
}

// Service operation to take a single Crew object as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *CrewService) New(ctx context.Context, body CrewNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/crew"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Crew record by its unique ID passed as a path
// parameter.
func (r *CrewService) Get(ctx context.Context, id string, query CrewGetParams, opts ...option.RequestOption) (res *shared.CrewFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/crew/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single Crew record. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *CrewService) Update(ctx context.Context, id string, body CrewUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/crew/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *CrewService) List(ctx context.Context, query CrewListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[CrewAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/crew"
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
func (r *CrewService) ListAutoPaging(ctx context.Context, query CrewListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[CrewAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *CrewService) Count(ctx context.Context, query CrewCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/crew/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *CrewService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *CrewQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/crew/queryhelp"
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
func (r *CrewService) Tuple(ctx context.Context, query CrewTupleParams, opts ...option.RequestOption) (res *[]shared.CrewFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/crew/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple Crew objects as a POST body and ingest into
// the database. This operation is intended to be used for automated feeds into
// UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *CrewService) UnvalidatedPublish(ctx context.Context, body CrewUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-crew"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Crew Services.
type CrewAbridged struct {
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
	DataMode CrewAbridgedDataMode `json:"dataMode,required"`
	// Unique identifier of the formed crew provided by the originating source.
	// Provided for systems that require tracking of an internal system generated ID.
	OrigCrewID string `json:"origCrewId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Adjusted return time, in ISO 8601 UTC format with millisecond precision.
	AdjReturnTime time.Time `json:"adjReturnTime" format:"date-time"`
	// Last name of the adjusted return time approver.
	AdjReturnTimeApprover string `json:"adjReturnTimeApprover"`
	// The aircraft Model Design Series designation assigned for this crew.
	AircraftMds string `json:"aircraftMDS"`
	// Time the crew was alerted, in ISO 8601 UTC format with millisecond precision.
	AlertedTime time.Time `json:"alertedTime" format:"date-time"`
	// Type of alert for the crew (e.g., ALPHA for maximum readiness, BRAVO for
	// standby, etc.).
	AlertType string `json:"alertType"`
	// The crew's Aviation Resource Management System (ARMS) unit. If multiple units
	// exist, use the Aircraft Commander's Unit.
	ArmsCrewUnit string `json:"armsCrewUnit"`
	// Array of qualification codes assigned to this crew (e.g., AL for Aircraft
	// Leader, CS for Combat Systems Operator, etc.).
	AssignedQualCode []string `json:"assignedQualCode"`
	// Unique identifier of the crew commander assigned by the originating source.
	CommanderID string `json:"commanderId"`
	// Last four digits of the crew commander's social security number.
	CommanderLast4Ssn string `json:"commanderLast4SSN"`
	// The name of the crew commander.
	CommanderName string `json:"commanderName"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Flag indicating whether this crew task takes the crew home and out of the stage.
	CrewHome bool `json:"crewHome"`
	// CrewMembers Collection.
	CrewMembers []CrewAbridgedCrewMember `json:"crewMembers"`
	// Name of the formed crew.
	CrewName string `json:"crewName"`
	// The resource management system managing and reporting data on this crew.
	CrewRms string `json:"crewRMS"`
	// The crew's role on the mission (e.g., DEADHEAD, MEDICAL, PRIMARY).
	CrewRole string `json:"crewRole"`
	// The military component that comprises the crew (e.g., ACTIVE, RESERVE, GUARD,
	// MIXED, UNKNOWN, etc.).
	CrewSource string `json:"crewSource"`
	// The squadron the crew serves.
	CrewSquadron string `json:"crewSquadron"`
	// The type of crew required to meet mission objectives (e.g., AIRDROP, AIRLAND,
	// AIR REFUELING, etc.).
	CrewType string `json:"crewType"`
	// The crew's squadron as identified in its resource management system. If the crew
	// is composed of members from multiple units, then the Crew Commander's unit
	// should be indicated as the crew unit.
	CrewUnit string `json:"crewUnit"`
	// The wing the crew serves.
	CrewWing string `json:"crewWing"`
	// The International Civil Aviation Organization (ICAO) code of the airfield at
	// which the crew is currently located.
	CurrentIcao string `json:"currentICAO"`
	// Crew Flight Duty Period (FDP) eligibility type.
	FdpEligType string `json:"fdpEligType"`
	// Flight Duty Period (FDP) type.
	FdpType string `json:"fdpType"`
	// The number of female enlisted crew members.
	FemaleEnlistedQty int64 `json:"femaleEnlistedQty"`
	// The number of female officer crew members.
	FemaleOfficerQty int64 `json:"femaleOfficerQty"`
	// Authorization number used on the flight order.
	FltAuthNum string `json:"fltAuthNum"`
	// Unique identifier of the Site at which the crew is currently located. This ID
	// can be used to obtain additional information on a Site using the 'get by ID'
	// operation (e.g. /udl/site/{id}). For example, the Site object with idSite = abc
	// would be queried as /udl/site/abc.
	IDSiteCurrent string `json:"idSiteCurrent"`
	// Unique identifier of the Aircraft Sortie associated with this crew record.
	IDSortie string `json:"idSortie"`
	// Initial start time of the crew's linked task that was delinked due to mission
	// closure, in ISO 8601 UTC format with millisecond precision.
	InitStartTime time.Time `json:"initStartTime" format:"date-time"`
	// The last time the crew can be alerted, in ISO 8601 UTC format with millisecond
	// precision.
	LastAlertTime time.Time `json:"lastAlertTime" format:"date-time"`
	// Time the crew is legal for alert, in ISO 8601 UTC format with millisecond
	// precision.
	LegalAlertTime time.Time `json:"legalAlertTime" format:"date-time"`
	// Time the crew is legally authorized or scheduled to remain on standby for duty,
	// in ISO 8601 UTC format with millisecond precision.
	LegalBravoTime time.Time `json:"legalBravoTime" format:"date-time"`
	// Flag indicating whether this crew is part of a linked flying task.
	LinkedTask bool `json:"linkedTask"`
	// The number of male enlisted crew members.
	MaleEnlistedQty int64 `json:"maleEnlistedQty"`
	// The number of male officer crew members.
	MaleOfficerQty int64 `json:"maleOfficerQty"`
	// User-defined alias designation for the mission.
	MissionAlias string `json:"missionAlias"`
	// The mission ID the crew is supporting according to the source system.
	MissionID string `json:"missionId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The type of personnel that comprises the crew (e.g., AIRCREW, MEDCREW, etc.).
	PersonnelType string `json:"personnelType"`
	// Time the crew will be picked up from lodging, in ISO 8601 UTC format with
	// millisecond precision.
	PickupTime time.Time `json:"pickupTime" format:"date-time"`
	// Flag indicating whether post-mission crew rest is applied to the last sortie of
	// a crew's task.
	PostRestApplied bool `json:"postRestApplied"`
	// End time of the crew rest period after the mission, in ISO 8601 UTC format with
	// millisecond precision.
	PostRestEnd time.Time `json:"postRestEnd" format:"date-time"`
	// The scheduled delay or adjustment in the start time of a crew's rest period
	// after a mission, expressed as +/-HH:MM.
	PostRestOffset string `json:"postRestOffset"`
	// Flag indicating whether pre-mission crew rest is applied to the first sortie of
	// a crew's task.
	PreRestApplied bool `json:"preRestApplied"`
	// Start time of the crew rest period before the mission, in ISO 8601 UTC format
	// with millisecond precision.
	PreRestStart time.Time `json:"preRestStart" format:"date-time"`
	// Array of qualification codes required for this crew (e.g., AL for Aircraft
	// Leader, CS for Combat Systems Operator, etc.).
	ReqQualCode []string `json:"reqQualCode"`
	// Scheduled return time, in ISO 8601 UTC format with millisecond precision.
	ReturnTime time.Time `json:"returnTime" format:"date-time"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The stage 1 qualifications the crew must have for a mission, such as having
	// basic knowledge of crew operations and aircraft systems.
	Stage1Qual string `json:"stage1Qual"`
	// The stage 2 qualifications the crew must have for a mission, such as completion
	// of advanced mission-specific training.
	Stage2Qual string `json:"stage2Qual"`
	// The stage 3 qualifications the crew must have for a mission, such as full
	// mission-ready certification and the capability of leading complex operations.
	Stage3Qual string `json:"stage3Qual"`
	// Stage name for the crew. A stage is a pool of crews supporting a given operation
	// plan.
	StageName string `json:"stageName"`
	// Time the crew entered the stage, in ISO 8601 UTC format with millisecond
	// precision.
	StageTime time.Time `json:"stageTime" format:"date-time"`
	// Crew status (e.g. NEEDCREW, ASSIGNED, APPROVED, NOTIFIED, PARTIAL, UNKNOWN,
	// etc.).
	Status string `json:"status"`
	// Flag indicating that one or more crew members requires transportation to the
	// departure location.
	TransportReq bool `json:"transportReq"`
	// Identifies the trip kit needed by the crew. A trip kit contains charts,
	// regulations, maps, etc. carried by the crew during missions.
	TripKit string `json:"tripKit"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		OrigCrewID            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AdjReturnTime         respjson.Field
		AdjReturnTimeApprover respjson.Field
		AircraftMds           respjson.Field
		AlertedTime           respjson.Field
		AlertType             respjson.Field
		ArmsCrewUnit          respjson.Field
		AssignedQualCode      respjson.Field
		CommanderID           respjson.Field
		CommanderLast4Ssn     respjson.Field
		CommanderName         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CrewHome              respjson.Field
		CrewMembers           respjson.Field
		CrewName              respjson.Field
		CrewRms               respjson.Field
		CrewRole              respjson.Field
		CrewSource            respjson.Field
		CrewSquadron          respjson.Field
		CrewType              respjson.Field
		CrewUnit              respjson.Field
		CrewWing              respjson.Field
		CurrentIcao           respjson.Field
		FdpEligType           respjson.Field
		FdpType               respjson.Field
		FemaleEnlistedQty     respjson.Field
		FemaleOfficerQty      respjson.Field
		FltAuthNum            respjson.Field
		IDSiteCurrent         respjson.Field
		IDSortie              respjson.Field
		InitStartTime         respjson.Field
		LastAlertTime         respjson.Field
		LegalAlertTime        respjson.Field
		LegalBravoTime        respjson.Field
		LinkedTask            respjson.Field
		MaleEnlistedQty       respjson.Field
		MaleOfficerQty        respjson.Field
		MissionAlias          respjson.Field
		MissionID             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PersonnelType         respjson.Field
		PickupTime            respjson.Field
		PostRestApplied       respjson.Field
		PostRestEnd           respjson.Field
		PostRestOffset        respjson.Field
		PreRestApplied        respjson.Field
		PreRestStart          respjson.Field
		ReqQualCode           respjson.Field
		ReturnTime            respjson.Field
		SourceDl              respjson.Field
		Stage1Qual            respjson.Field
		Stage2Qual            respjson.Field
		Stage3Qual            respjson.Field
		StageName             respjson.Field
		StageTime             respjson.Field
		Status                respjson.Field
		TransportReq          respjson.Field
		TripKit               respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrewAbridged) RawJSON() string { return r.JSON.raw }
func (r *CrewAbridged) UnmarshalJSON(data []byte) error {
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
type CrewAbridgedDataMode string

const (
	CrewAbridgedDataModeReal      CrewAbridgedDataMode = "REAL"
	CrewAbridgedDataModeTest      CrewAbridgedDataMode = "TEST"
	CrewAbridgedDataModeSimulated CrewAbridgedDataMode = "SIMULATED"
	CrewAbridgedDataModeExercise  CrewAbridgedDataMode = "EXERCISE"
)

// Schema for Crew Member data.
type CrewAbridgedCrewMember struct {
	// Flag indicating whether this crew member has been alerted of the associated
	// task.
	Alerted bool `json:"alerted"`
	// Flag indicating this crew member is assigned to all sorties of the crew
	// itinerary.
	AllSortie bool `json:"allSortie"`
	// Flag indicating whether this crew member has been approved for the associated
	// task.
	Approved bool `json:"approved"`
	// Flag indicating whether this crew member is attached to his/her squadron. Crew
	// members that are not attached are considered assigned.
	Attached bool `json:"attached"`
	// The military branch assignment of the crew member.
	Branch string `json:"branch"`
	// Flag indicating this crew member is a civilian or non-military person.
	Civilian bool `json:"civilian"`
	// Flag indicating this person is the aircraft commander.
	Commander bool `json:"commander"`
	// The crew position of the crew member.
	CrewPosition string `json:"crewPosition"`
	// The crew member's 10-digit DoD ID number.
	DodID string `json:"dodID"`
	// The duty position of the crew member.
	DutyPosition string `json:"dutyPosition"`
	// The current duty status code of this crew member (e.g., AGR for Active Guard and
	// Reserve, IDT for Inactive Duty Training, etc.).
	DutyStatus string `json:"dutyStatus"`
	// Flag indicating whether this crew member has been notified of an event by email.
	Emailed bool `json:"emailed"`
	// Flag indicating whether this crew member requires an additional amount of time
	// to report for duty.
	ExtraTime bool `json:"extraTime"`
	// The first name of the crew member.
	FirstName string `json:"firstName"`
	// The earliest flying currency expiration date for this crew member, in ISO 8601
	// UTC format with millisecond precision.
	FltCurrencyExp time.Time `json:"fltCurrencyExp" format:"date-time"`
	// The training task identifier associated with the flying currency expiration date
	// for this crew member.
	FltCurrencyExpID string `json:"fltCurrencyExpId"`
	// The date this crew member's records review was completed, in ISO 8601 UTC format
	// with millisecond precision.
	FltRecDate time.Time `json:"fltRecDate" format:"date-time"`
	// The date this crew member's records review is due, in ISO 8601 UTC format with
	// millisecond precision.
	FltRecDue time.Time `json:"fltRecDue" format:"date-time"`
	// The flying squadron assignment of the crew member.
	FlySquadron string `json:"flySquadron"`
	// Flag indicating whether this crew member is funded.
	Funded bool `json:"funded"`
	// Gender of the crew member.
	Gender string `json:"gender"`
	// The earliest ground currency expiration date for this crew member, in ISO 8601
	// UTC format with millisecond precision.
	GndCurrencyExp time.Time `json:"gndCurrencyExp" format:"date-time"`
	// The training task identifier associated with the ground currency expiration date
	// for this crew member.
	GndCurrencyExpID string `json:"gndCurrencyExpId"`
	// Flag indicating whether this crew member is grounded (i.e., his/her duties do
	// not include flying).
	Grounded bool `json:"grounded"`
	// Date when this crew member starts acting as guest help for the squadron, in ISO
	// 8601 UTC format with millisecond precision.
	GuestStart time.Time `json:"guestStart" format:"date-time"`
	// Date when this crew member stops acting as guest help for the squadron, in ISO
	// 8601 UTC format with millisecond precision.
	GuestStop time.Time `json:"guestStop" format:"date-time"`
	// Last four digits of the crew member's social security number.
	Last4Ssn string `json:"last4SSN"`
	// Date of the last flight for this crew member, in ISO 8601 UTC format with
	// millisecond precision.
	LastFltDate time.Time `json:"lastFltDate" format:"date-time"`
	// The last name of the crew member.
	LastName string `json:"lastName"`
	// The squadron the crew member has been temporarily loaned to.
	LoanedTo string `json:"loanedTo"`
	// Crew member lodging location.
	Lodging string `json:"lodging"`
	// Time this crew member was actually alerted for the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberActualAlertTime time.Time `json:"memberActualAlertTime" format:"date-time"`
	// Adjusted return time for the crew member, in ISO 8601 UTC format with
	// millisecond precision.
	MemberAdjReturnTime time.Time `json:"memberAdjReturnTime" format:"date-time"`
	// Last name of the crew member's adjusted return time approver.
	MemberAdjReturnTimeApprover string `json:"memberAdjReturnTimeApprover"`
	// Unique identifier of the crew member assigned by the originating source.
	MemberID string `json:"memberId"`
	// Initial start time of the crew member's linked task that was delinked due to
	// mission closure, in ISO 8601 UTC format with millisecond precision.
	MemberInitStartTime time.Time `json:"memberInitStartTime" format:"date-time"`
	// The latest possible time the crew member can legally be alerted for a task, in
	// ISO 8601 UTC format with millisecond precision.
	MemberLastAlertTime time.Time `json:"memberLastAlertTime" format:"date-time"`
	// Time this crew member becomes eligible to be alerted for the mission, in ISO
	// 8601 UTC format with millisecond precision.
	MemberLegalAlertTime time.Time `json:"memberLegalAlertTime" format:"date-time"`
	// Time this crew member will be picked up from lodging, in ISO 8601 UTC format
	// with millisecond precision.
	MemberPickupTime time.Time `json:"memberPickupTime" format:"date-time"`
	// The scheduled delay or adjustment in the start time of a crew member's rest
	// period after a mission, expressed as +/-HH:MM.
	MemberPostRestOffset string `json:"memberPostRestOffset"`
	// End time of this crew member's rest period after the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberPostRestTime time.Time `json:"memberPostRestTime" format:"date-time"`
	// Start time of this crew member's rest period before the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberPreRestTime time.Time `json:"memberPreRestTime" format:"date-time"`
	// Remarks concerning the crew member.
	MemberRemarks string `json:"memberRemarks"`
	// Scheduled return time for this crew member, in ISO 8601 UTC format with
	// millisecond precision.
	MemberReturnTime time.Time `json:"memberReturnTime" format:"date-time"`
	// Time this crew member is scheduled to be alerted for the mission, in ISO 8601
	// UTC format with millisecond precision.
	MemberSchedAlertTime time.Time `json:"memberSchedAlertTime" format:"date-time"`
	// The military component for the crew member (e.g., ACTIVE, RESERVE, GUARD,
	// UNKNOWN, etc.).
	MemberSource string `json:"memberSource"`
	// Stage name for the crew member. A stage is a pool of crews supporting a given
	// operation plan.
	MemberStageName string `json:"memberStageName"`
	// Flag indicating whether this crew member needs transportation to the departure
	// location.
	MemberTransportReq bool `json:"memberTransportReq"`
	// Amplifying details about the crew member type (e.g. RAVEN, FCC, COMCAM, AIRCREW,
	// MEP, OTHER, etc.).
	MemberType string `json:"memberType"`
	// The middle initial of the crew member.
	MiddleInitial string `json:"middleInitial"`
	// Flag indicating whether this crew member has been notified of an event.
	Notified bool `json:"notified"`
	// Crew member lodging phone number.
	PhoneNumber string `json:"phoneNumber"`
	// Code indicating a crew member's current physical fitness status and whether they
	// are medically cleared to fly (e.g., D for Duties Not Including Flying, E for
	// Physical Overdue, etc.).
	PhysAvCode string `json:"physAvCode"`
	// Code indicating a crew member's physical availabiility status (e.g.,
	// DISQUALIFIED, OVERDUE, etc.).
	PhysAvStatus string `json:"physAvStatus"`
	// Due date for the crew member's physical, in ISO 8601 UTC format with millisecond
	// precision.
	PhysDue time.Time `json:"physDue" format:"date-time"`
	// The rank of the crew member.
	Rank string `json:"rank"`
	// Remark code used to designate attributes of this crew member. For more
	// information, contact the provider source.
	RemarkCode string `json:"remarkCode"`
	// The primary aircraft type for the crew member according to the personnel
	// resource management system indicated in the crewRMS field.
	RmsMds string `json:"rmsMDS"`
	// Time this crew member is required to report for duty before this flight/mission,
	// in ISO 8601 UTC format with millisecond precision.
	ShowTime time.Time `json:"showTime" format:"date-time"`
	// The squadron the crew member serves.
	Squadron string `json:"squadron"`
	// The date this crew member accomplished physiological or altitude chamber
	// training, in ISO 8601 UTC format with millisecond precision.
	TrainingDate time.Time `json:"trainingDate" format:"date-time"`
	// The Mattermost username of this crew member.
	Username string `json:"username"`
	// The wing the crew member serves.
	Wing string `json:"wing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alerted                     respjson.Field
		AllSortie                   respjson.Field
		Approved                    respjson.Field
		Attached                    respjson.Field
		Branch                      respjson.Field
		Civilian                    respjson.Field
		Commander                   respjson.Field
		CrewPosition                respjson.Field
		DodID                       respjson.Field
		DutyPosition                respjson.Field
		DutyStatus                  respjson.Field
		Emailed                     respjson.Field
		ExtraTime                   respjson.Field
		FirstName                   respjson.Field
		FltCurrencyExp              respjson.Field
		FltCurrencyExpID            respjson.Field
		FltRecDate                  respjson.Field
		FltRecDue                   respjson.Field
		FlySquadron                 respjson.Field
		Funded                      respjson.Field
		Gender                      respjson.Field
		GndCurrencyExp              respjson.Field
		GndCurrencyExpID            respjson.Field
		Grounded                    respjson.Field
		GuestStart                  respjson.Field
		GuestStop                   respjson.Field
		Last4Ssn                    respjson.Field
		LastFltDate                 respjson.Field
		LastName                    respjson.Field
		LoanedTo                    respjson.Field
		Lodging                     respjson.Field
		MemberActualAlertTime       respjson.Field
		MemberAdjReturnTime         respjson.Field
		MemberAdjReturnTimeApprover respjson.Field
		MemberID                    respjson.Field
		MemberInitStartTime         respjson.Field
		MemberLastAlertTime         respjson.Field
		MemberLegalAlertTime        respjson.Field
		MemberPickupTime            respjson.Field
		MemberPostRestOffset        respjson.Field
		MemberPostRestTime          respjson.Field
		MemberPreRestTime           respjson.Field
		MemberRemarks               respjson.Field
		MemberReturnTime            respjson.Field
		MemberSchedAlertTime        respjson.Field
		MemberSource                respjson.Field
		MemberStageName             respjson.Field
		MemberTransportReq          respjson.Field
		MemberType                  respjson.Field
		MiddleInitial               respjson.Field
		Notified                    respjson.Field
		PhoneNumber                 respjson.Field
		PhysAvCode                  respjson.Field
		PhysAvStatus                respjson.Field
		PhysDue                     respjson.Field
		Rank                        respjson.Field
		RemarkCode                  respjson.Field
		RmsMds                      respjson.Field
		ShowTime                    respjson.Field
		Squadron                    respjson.Field
		TrainingDate                respjson.Field
		Username                    respjson.Field
		Wing                        respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrewAbridgedCrewMember) RawJSON() string { return r.JSON.raw }
func (r *CrewAbridgedCrewMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrewQueryhelpResponse struct {
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
func (r CrewQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *CrewQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrewNewParams struct {
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
	DataMode CrewNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the formed crew provided by the originating source.
	// Provided for systems that require tracking of an internal system generated ID.
	OrigCrewID string `json:"origCrewId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Adjusted return time, in ISO 8601 UTC format with millisecond precision.
	AdjReturnTime param.Opt[time.Time] `json:"adjReturnTime,omitzero" format:"date-time"`
	// Last name of the adjusted return time approver.
	AdjReturnTimeApprover param.Opt[string] `json:"adjReturnTimeApprover,omitzero"`
	// The aircraft Model Design Series designation assigned for this crew.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Time the crew was alerted, in ISO 8601 UTC format with millisecond precision.
	AlertedTime param.Opt[time.Time] `json:"alertedTime,omitzero" format:"date-time"`
	// Type of alert for the crew (e.g., ALPHA for maximum readiness, BRAVO for
	// standby, etc.).
	AlertType param.Opt[string] `json:"alertType,omitzero"`
	// The crew's Aviation Resource Management System (ARMS) unit. If multiple units
	// exist, use the Aircraft Commander's Unit.
	ArmsCrewUnit param.Opt[string] `json:"armsCrewUnit,omitzero"`
	// Unique identifier of the crew commander assigned by the originating source.
	CommanderID param.Opt[string] `json:"commanderId,omitzero"`
	// Last four digits of the crew commander's social security number.
	CommanderLast4Ssn param.Opt[string] `json:"commanderLast4SSN,omitzero"`
	// The name of the crew commander.
	CommanderName param.Opt[string] `json:"commanderName,omitzero"`
	// Flag indicating whether this crew task takes the crew home and out of the stage.
	CrewHome param.Opt[bool] `json:"crewHome,omitzero"`
	// Name of the formed crew.
	CrewName param.Opt[string] `json:"crewName,omitzero"`
	// The resource management system managing and reporting data on this crew.
	CrewRms param.Opt[string] `json:"crewRMS,omitzero"`
	// The crew's role on the mission (e.g., DEADHEAD, MEDICAL, PRIMARY).
	CrewRole param.Opt[string] `json:"crewRole,omitzero"`
	// The military component that comprises the crew (e.g., ACTIVE, RESERVE, GUARD,
	// MIXED, UNKNOWN, etc.).
	CrewSource param.Opt[string] `json:"crewSource,omitzero"`
	// The squadron the crew serves.
	CrewSquadron param.Opt[string] `json:"crewSquadron,omitzero"`
	// The type of crew required to meet mission objectives (e.g., AIRDROP, AIRLAND,
	// AIR REFUELING, etc.).
	CrewType param.Opt[string] `json:"crewType,omitzero"`
	// The crew's squadron as identified in its resource management system. If the crew
	// is composed of members from multiple units, then the Crew Commander's unit
	// should be indicated as the crew unit.
	CrewUnit param.Opt[string] `json:"crewUnit,omitzero"`
	// The wing the crew serves.
	CrewWing param.Opt[string] `json:"crewWing,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the airfield at
	// which the crew is currently located.
	CurrentIcao param.Opt[string] `json:"currentICAO,omitzero"`
	// Crew Flight Duty Period (FDP) eligibility type.
	FdpEligType param.Opt[string] `json:"fdpEligType,omitzero"`
	// Flight Duty Period (FDP) type.
	FdpType param.Opt[string] `json:"fdpType,omitzero"`
	// The number of female enlisted crew members.
	FemaleEnlistedQty param.Opt[int64] `json:"femaleEnlistedQty,omitzero"`
	// The number of female officer crew members.
	FemaleOfficerQty param.Opt[int64] `json:"femaleOfficerQty,omitzero"`
	// Authorization number used on the flight order.
	FltAuthNum param.Opt[string] `json:"fltAuthNum,omitzero"`
	// Unique identifier of the Site at which the crew is currently located. This ID
	// can be used to obtain additional information on a Site using the 'get by ID'
	// operation (e.g. /udl/site/{id}). For example, the Site object with idSite = abc
	// would be queried as /udl/site/abc.
	IDSiteCurrent param.Opt[string] `json:"idSiteCurrent,omitzero"`
	// Unique identifier of the Aircraft Sortie associated with this crew record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Initial start time of the crew's linked task that was delinked due to mission
	// closure, in ISO 8601 UTC format with millisecond precision.
	InitStartTime param.Opt[time.Time] `json:"initStartTime,omitzero" format:"date-time"`
	// The last time the crew can be alerted, in ISO 8601 UTC format with millisecond
	// precision.
	LastAlertTime param.Opt[time.Time] `json:"lastAlertTime,omitzero" format:"date-time"`
	// Time the crew is legal for alert, in ISO 8601 UTC format with millisecond
	// precision.
	LegalAlertTime param.Opt[time.Time] `json:"legalAlertTime,omitzero" format:"date-time"`
	// Time the crew is legally authorized or scheduled to remain on standby for duty,
	// in ISO 8601 UTC format with millisecond precision.
	LegalBravoTime param.Opt[time.Time] `json:"legalBravoTime,omitzero" format:"date-time"`
	// Flag indicating whether this crew is part of a linked flying task.
	LinkedTask param.Opt[bool] `json:"linkedTask,omitzero"`
	// The number of male enlisted crew members.
	MaleEnlistedQty param.Opt[int64] `json:"maleEnlistedQty,omitzero"`
	// The number of male officer crew members.
	MaleOfficerQty param.Opt[int64] `json:"maleOfficerQty,omitzero"`
	// User-defined alias designation for the mission.
	MissionAlias param.Opt[string] `json:"missionAlias,omitzero"`
	// The mission ID the crew is supporting according to the source system.
	MissionID param.Opt[string] `json:"missionId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The type of personnel that comprises the crew (e.g., AIRCREW, MEDCREW, etc.).
	PersonnelType param.Opt[string] `json:"personnelType,omitzero"`
	// Time the crew will be picked up from lodging, in ISO 8601 UTC format with
	// millisecond precision.
	PickupTime param.Opt[time.Time] `json:"pickupTime,omitzero" format:"date-time"`
	// Flag indicating whether post-mission crew rest is applied to the last sortie of
	// a crew's task.
	PostRestApplied param.Opt[bool] `json:"postRestApplied,omitzero"`
	// End time of the crew rest period after the mission, in ISO 8601 UTC format with
	// millisecond precision.
	PostRestEnd param.Opt[time.Time] `json:"postRestEnd,omitzero" format:"date-time"`
	// The scheduled delay or adjustment in the start time of a crew's rest period
	// after a mission, expressed as +/-HH:MM.
	PostRestOffset param.Opt[string] `json:"postRestOffset,omitzero"`
	// Flag indicating whether pre-mission crew rest is applied to the first sortie of
	// a crew's task.
	PreRestApplied param.Opt[bool] `json:"preRestApplied,omitzero"`
	// Start time of the crew rest period before the mission, in ISO 8601 UTC format
	// with millisecond precision.
	PreRestStart param.Opt[time.Time] `json:"preRestStart,omitzero" format:"date-time"`
	// Scheduled return time, in ISO 8601 UTC format with millisecond precision.
	ReturnTime param.Opt[time.Time] `json:"returnTime,omitzero" format:"date-time"`
	// The stage 1 qualifications the crew must have for a mission, such as having
	// basic knowledge of crew operations and aircraft systems.
	Stage1Qual param.Opt[string] `json:"stage1Qual,omitzero"`
	// The stage 2 qualifications the crew must have for a mission, such as completion
	// of advanced mission-specific training.
	Stage2Qual param.Opt[string] `json:"stage2Qual,omitzero"`
	// The stage 3 qualifications the crew must have for a mission, such as full
	// mission-ready certification and the capability of leading complex operations.
	Stage3Qual param.Opt[string] `json:"stage3Qual,omitzero"`
	// Stage name for the crew. A stage is a pool of crews supporting a given operation
	// plan.
	StageName param.Opt[string] `json:"stageName,omitzero"`
	// Time the crew entered the stage, in ISO 8601 UTC format with millisecond
	// precision.
	StageTime param.Opt[time.Time] `json:"stageTime,omitzero" format:"date-time"`
	// Crew status (e.g. NEEDCREW, ASSIGNED, APPROVED, NOTIFIED, PARTIAL, UNKNOWN,
	// etc.).
	Status param.Opt[string] `json:"status,omitzero"`
	// Flag indicating that one or more crew members requires transportation to the
	// departure location.
	TransportReq param.Opt[bool] `json:"transportReq,omitzero"`
	// Identifies the trip kit needed by the crew. A trip kit contains charts,
	// regulations, maps, etc. carried by the crew during missions.
	TripKit param.Opt[string] `json:"tripKit,omitzero"`
	// Array of qualification codes assigned to this crew (e.g., AL for Aircraft
	// Leader, CS for Combat Systems Operator, etc.).
	AssignedQualCode []string `json:"assignedQualCode,omitzero"`
	// CrewMembers Collection.
	CrewMembers []CrewNewParamsCrewMember `json:"crewMembers,omitzero"`
	// Array of qualification codes required for this crew (e.g., AL for Aircraft
	// Leader, CS for Combat Systems Operator, etc.).
	ReqQualCode []string `json:"reqQualCode,omitzero"`
	paramObj
}

func (r CrewNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CrewNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrewNewParams) UnmarshalJSON(data []byte) error {
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
type CrewNewParamsDataMode string

const (
	CrewNewParamsDataModeReal      CrewNewParamsDataMode = "REAL"
	CrewNewParamsDataModeTest      CrewNewParamsDataMode = "TEST"
	CrewNewParamsDataModeSimulated CrewNewParamsDataMode = "SIMULATED"
	CrewNewParamsDataModeExercise  CrewNewParamsDataMode = "EXERCISE"
)

// Schema for Crew Member data.
type CrewNewParamsCrewMember struct {
	// Flag indicating whether this crew member has been alerted of the associated
	// task.
	Alerted param.Opt[bool] `json:"alerted,omitzero"`
	// Flag indicating this crew member is assigned to all sorties of the crew
	// itinerary.
	AllSortie param.Opt[bool] `json:"allSortie,omitzero"`
	// Flag indicating whether this crew member has been approved for the associated
	// task.
	Approved param.Opt[bool] `json:"approved,omitzero"`
	// Flag indicating whether this crew member is attached to his/her squadron. Crew
	// members that are not attached are considered assigned.
	Attached param.Opt[bool] `json:"attached,omitzero"`
	// The military branch assignment of the crew member.
	Branch param.Opt[string] `json:"branch,omitzero"`
	// Flag indicating this crew member is a civilian or non-military person.
	Civilian param.Opt[bool] `json:"civilian,omitzero"`
	// Flag indicating this person is the aircraft commander.
	Commander param.Opt[bool] `json:"commander,omitzero"`
	// The crew position of the crew member.
	CrewPosition param.Opt[string] `json:"crewPosition,omitzero"`
	// The crew member's 10-digit DoD ID number.
	DodID param.Opt[string] `json:"dodID,omitzero"`
	// The duty position of the crew member.
	DutyPosition param.Opt[string] `json:"dutyPosition,omitzero"`
	// The current duty status code of this crew member (e.g., AGR for Active Guard and
	// Reserve, IDT for Inactive Duty Training, etc.).
	DutyStatus param.Opt[string] `json:"dutyStatus,omitzero"`
	// Flag indicating whether this crew member has been notified of an event by email.
	Emailed param.Opt[bool] `json:"emailed,omitzero"`
	// Flag indicating whether this crew member requires an additional amount of time
	// to report for duty.
	ExtraTime param.Opt[bool] `json:"extraTime,omitzero"`
	// The first name of the crew member.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The earliest flying currency expiration date for this crew member, in ISO 8601
	// UTC format with millisecond precision.
	FltCurrencyExp param.Opt[time.Time] `json:"fltCurrencyExp,omitzero" format:"date-time"`
	// The training task identifier associated with the flying currency expiration date
	// for this crew member.
	FltCurrencyExpID param.Opt[string] `json:"fltCurrencyExpId,omitzero"`
	// The date this crew member's records review was completed, in ISO 8601 UTC format
	// with millisecond precision.
	FltRecDate param.Opt[time.Time] `json:"fltRecDate,omitzero" format:"date-time"`
	// The date this crew member's records review is due, in ISO 8601 UTC format with
	// millisecond precision.
	FltRecDue param.Opt[time.Time] `json:"fltRecDue,omitzero" format:"date-time"`
	// The flying squadron assignment of the crew member.
	FlySquadron param.Opt[string] `json:"flySquadron,omitzero"`
	// Flag indicating whether this crew member is funded.
	Funded param.Opt[bool] `json:"funded,omitzero"`
	// Gender of the crew member.
	Gender param.Opt[string] `json:"gender,omitzero"`
	// The earliest ground currency expiration date for this crew member, in ISO 8601
	// UTC format with millisecond precision.
	GndCurrencyExp param.Opt[time.Time] `json:"gndCurrencyExp,omitzero" format:"date-time"`
	// The training task identifier associated with the ground currency expiration date
	// for this crew member.
	GndCurrencyExpID param.Opt[string] `json:"gndCurrencyExpId,omitzero"`
	// Flag indicating whether this crew member is grounded (i.e., his/her duties do
	// not include flying).
	Grounded param.Opt[bool] `json:"grounded,omitzero"`
	// Date when this crew member starts acting as guest help for the squadron, in ISO
	// 8601 UTC format with millisecond precision.
	GuestStart param.Opt[time.Time] `json:"guestStart,omitzero" format:"date-time"`
	// Date when this crew member stops acting as guest help for the squadron, in ISO
	// 8601 UTC format with millisecond precision.
	GuestStop param.Opt[time.Time] `json:"guestStop,omitzero" format:"date-time"`
	// Last four digits of the crew member's social security number.
	Last4Ssn param.Opt[string] `json:"last4SSN,omitzero"`
	// Date of the last flight for this crew member, in ISO 8601 UTC format with
	// millisecond precision.
	LastFltDate param.Opt[time.Time] `json:"lastFltDate,omitzero" format:"date-time"`
	// The last name of the crew member.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// The squadron the crew member has been temporarily loaned to.
	LoanedTo param.Opt[string] `json:"loanedTo,omitzero"`
	// Crew member lodging location.
	Lodging param.Opt[string] `json:"lodging,omitzero"`
	// Time this crew member was actually alerted for the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberActualAlertTime param.Opt[time.Time] `json:"memberActualAlertTime,omitzero" format:"date-time"`
	// Adjusted return time for the crew member, in ISO 8601 UTC format with
	// millisecond precision.
	MemberAdjReturnTime param.Opt[time.Time] `json:"memberAdjReturnTime,omitzero" format:"date-time"`
	// Last name of the crew member's adjusted return time approver.
	MemberAdjReturnTimeApprover param.Opt[string] `json:"memberAdjReturnTimeApprover,omitzero"`
	// Unique identifier of the crew member assigned by the originating source.
	MemberID param.Opt[string] `json:"memberId,omitzero"`
	// Initial start time of the crew member's linked task that was delinked due to
	// mission closure, in ISO 8601 UTC format with millisecond precision.
	MemberInitStartTime param.Opt[time.Time] `json:"memberInitStartTime,omitzero" format:"date-time"`
	// The latest possible time the crew member can legally be alerted for a task, in
	// ISO 8601 UTC format with millisecond precision.
	MemberLastAlertTime param.Opt[time.Time] `json:"memberLastAlertTime,omitzero" format:"date-time"`
	// Time this crew member becomes eligible to be alerted for the mission, in ISO
	// 8601 UTC format with millisecond precision.
	MemberLegalAlertTime param.Opt[time.Time] `json:"memberLegalAlertTime,omitzero" format:"date-time"`
	// Time this crew member will be picked up from lodging, in ISO 8601 UTC format
	// with millisecond precision.
	MemberPickupTime param.Opt[time.Time] `json:"memberPickupTime,omitzero" format:"date-time"`
	// The scheduled delay or adjustment in the start time of a crew member's rest
	// period after a mission, expressed as +/-HH:MM.
	MemberPostRestOffset param.Opt[string] `json:"memberPostRestOffset,omitzero"`
	// End time of this crew member's rest period after the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberPostRestTime param.Opt[time.Time] `json:"memberPostRestTime,omitzero" format:"date-time"`
	// Start time of this crew member's rest period before the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberPreRestTime param.Opt[time.Time] `json:"memberPreRestTime,omitzero" format:"date-time"`
	// Remarks concerning the crew member.
	MemberRemarks param.Opt[string] `json:"memberRemarks,omitzero"`
	// Scheduled return time for this crew member, in ISO 8601 UTC format with
	// millisecond precision.
	MemberReturnTime param.Opt[time.Time] `json:"memberReturnTime,omitzero" format:"date-time"`
	// Time this crew member is scheduled to be alerted for the mission, in ISO 8601
	// UTC format with millisecond precision.
	MemberSchedAlertTime param.Opt[time.Time] `json:"memberSchedAlertTime,omitzero" format:"date-time"`
	// The military component for the crew member (e.g., ACTIVE, RESERVE, GUARD,
	// UNKNOWN, etc.).
	MemberSource param.Opt[string] `json:"memberSource,omitzero"`
	// Stage name for the crew member. A stage is a pool of crews supporting a given
	// operation plan.
	MemberStageName param.Opt[string] `json:"memberStageName,omitzero"`
	// Flag indicating whether this crew member needs transportation to the departure
	// location.
	MemberTransportReq param.Opt[bool] `json:"memberTransportReq,omitzero"`
	// Amplifying details about the crew member type (e.g. RAVEN, FCC, COMCAM, AIRCREW,
	// MEP, OTHER, etc.).
	MemberType param.Opt[string] `json:"memberType,omitzero"`
	// The middle initial of the crew member.
	MiddleInitial param.Opt[string] `json:"middleInitial,omitzero"`
	// Flag indicating whether this crew member has been notified of an event.
	Notified param.Opt[bool] `json:"notified,omitzero"`
	// Crew member lodging phone number.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// Code indicating a crew member's current physical fitness status and whether they
	// are medically cleared to fly (e.g., D for Duties Not Including Flying, E for
	// Physical Overdue, etc.).
	PhysAvCode param.Opt[string] `json:"physAvCode,omitzero"`
	// Code indicating a crew member's physical availabiility status (e.g.,
	// DISQUALIFIED, OVERDUE, etc.).
	PhysAvStatus param.Opt[string] `json:"physAvStatus,omitzero"`
	// Due date for the crew member's physical, in ISO 8601 UTC format with millisecond
	// precision.
	PhysDue param.Opt[time.Time] `json:"physDue,omitzero" format:"date-time"`
	// The rank of the crew member.
	Rank param.Opt[string] `json:"rank,omitzero"`
	// Remark code used to designate attributes of this crew member. For more
	// information, contact the provider source.
	RemarkCode param.Opt[string] `json:"remarkCode,omitzero"`
	// The primary aircraft type for the crew member according to the personnel
	// resource management system indicated in the crewRMS field.
	RmsMds param.Opt[string] `json:"rmsMDS,omitzero"`
	// Time this crew member is required to report for duty before this flight/mission,
	// in ISO 8601 UTC format with millisecond precision.
	ShowTime param.Opt[time.Time] `json:"showTime,omitzero" format:"date-time"`
	// The squadron the crew member serves.
	Squadron param.Opt[string] `json:"squadron,omitzero"`
	// The date this crew member accomplished physiological or altitude chamber
	// training, in ISO 8601 UTC format with millisecond precision.
	TrainingDate param.Opt[time.Time] `json:"trainingDate,omitzero" format:"date-time"`
	// The Mattermost username of this crew member.
	Username param.Opt[string] `json:"username,omitzero"`
	// The wing the crew member serves.
	Wing param.Opt[string] `json:"wing,omitzero"`
	paramObj
}

func (r CrewNewParamsCrewMember) MarshalJSON() (data []byte, err error) {
	type shadow CrewNewParamsCrewMember
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrewNewParamsCrewMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrewGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CrewGetParams]'s query parameters as `url.Values`.
func (r CrewGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CrewUpdateParams struct {
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
	DataMode CrewUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the formed crew provided by the originating source.
	// Provided for systems that require tracking of an internal system generated ID.
	OrigCrewID string `json:"origCrewId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Adjusted return time, in ISO 8601 UTC format with millisecond precision.
	AdjReturnTime param.Opt[time.Time] `json:"adjReturnTime,omitzero" format:"date-time"`
	// Last name of the adjusted return time approver.
	AdjReturnTimeApprover param.Opt[string] `json:"adjReturnTimeApprover,omitzero"`
	// The aircraft Model Design Series designation assigned for this crew.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Time the crew was alerted, in ISO 8601 UTC format with millisecond precision.
	AlertedTime param.Opt[time.Time] `json:"alertedTime,omitzero" format:"date-time"`
	// Type of alert for the crew (e.g., ALPHA for maximum readiness, BRAVO for
	// standby, etc.).
	AlertType param.Opt[string] `json:"alertType,omitzero"`
	// The crew's Aviation Resource Management System (ARMS) unit. If multiple units
	// exist, use the Aircraft Commander's Unit.
	ArmsCrewUnit param.Opt[string] `json:"armsCrewUnit,omitzero"`
	// Unique identifier of the crew commander assigned by the originating source.
	CommanderID param.Opt[string] `json:"commanderId,omitzero"`
	// Last four digits of the crew commander's social security number.
	CommanderLast4Ssn param.Opt[string] `json:"commanderLast4SSN,omitzero"`
	// The name of the crew commander.
	CommanderName param.Opt[string] `json:"commanderName,omitzero"`
	// Flag indicating whether this crew task takes the crew home and out of the stage.
	CrewHome param.Opt[bool] `json:"crewHome,omitzero"`
	// Name of the formed crew.
	CrewName param.Opt[string] `json:"crewName,omitzero"`
	// The resource management system managing and reporting data on this crew.
	CrewRms param.Opt[string] `json:"crewRMS,omitzero"`
	// The crew's role on the mission (e.g., DEADHEAD, MEDICAL, PRIMARY).
	CrewRole param.Opt[string] `json:"crewRole,omitzero"`
	// The military component that comprises the crew (e.g., ACTIVE, RESERVE, GUARD,
	// MIXED, UNKNOWN, etc.).
	CrewSource param.Opt[string] `json:"crewSource,omitzero"`
	// The squadron the crew serves.
	CrewSquadron param.Opt[string] `json:"crewSquadron,omitzero"`
	// The type of crew required to meet mission objectives (e.g., AIRDROP, AIRLAND,
	// AIR REFUELING, etc.).
	CrewType param.Opt[string] `json:"crewType,omitzero"`
	// The crew's squadron as identified in its resource management system. If the crew
	// is composed of members from multiple units, then the Crew Commander's unit
	// should be indicated as the crew unit.
	CrewUnit param.Opt[string] `json:"crewUnit,omitzero"`
	// The wing the crew serves.
	CrewWing param.Opt[string] `json:"crewWing,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the airfield at
	// which the crew is currently located.
	CurrentIcao param.Opt[string] `json:"currentICAO,omitzero"`
	// Crew Flight Duty Period (FDP) eligibility type.
	FdpEligType param.Opt[string] `json:"fdpEligType,omitzero"`
	// Flight Duty Period (FDP) type.
	FdpType param.Opt[string] `json:"fdpType,omitzero"`
	// The number of female enlisted crew members.
	FemaleEnlistedQty param.Opt[int64] `json:"femaleEnlistedQty,omitzero"`
	// The number of female officer crew members.
	FemaleOfficerQty param.Opt[int64] `json:"femaleOfficerQty,omitzero"`
	// Authorization number used on the flight order.
	FltAuthNum param.Opt[string] `json:"fltAuthNum,omitzero"`
	// Unique identifier of the Site at which the crew is currently located. This ID
	// can be used to obtain additional information on a Site using the 'get by ID'
	// operation (e.g. /udl/site/{id}). For example, the Site object with idSite = abc
	// would be queried as /udl/site/abc.
	IDSiteCurrent param.Opt[string] `json:"idSiteCurrent,omitzero"`
	// Unique identifier of the Aircraft Sortie associated with this crew record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Initial start time of the crew's linked task that was delinked due to mission
	// closure, in ISO 8601 UTC format with millisecond precision.
	InitStartTime param.Opt[time.Time] `json:"initStartTime,omitzero" format:"date-time"`
	// The last time the crew can be alerted, in ISO 8601 UTC format with millisecond
	// precision.
	LastAlertTime param.Opt[time.Time] `json:"lastAlertTime,omitzero" format:"date-time"`
	// Time the crew is legal for alert, in ISO 8601 UTC format with millisecond
	// precision.
	LegalAlertTime param.Opt[time.Time] `json:"legalAlertTime,omitzero" format:"date-time"`
	// Time the crew is legally authorized or scheduled to remain on standby for duty,
	// in ISO 8601 UTC format with millisecond precision.
	LegalBravoTime param.Opt[time.Time] `json:"legalBravoTime,omitzero" format:"date-time"`
	// Flag indicating whether this crew is part of a linked flying task.
	LinkedTask param.Opt[bool] `json:"linkedTask,omitzero"`
	// The number of male enlisted crew members.
	MaleEnlistedQty param.Opt[int64] `json:"maleEnlistedQty,omitzero"`
	// The number of male officer crew members.
	MaleOfficerQty param.Opt[int64] `json:"maleOfficerQty,omitzero"`
	// User-defined alias designation for the mission.
	MissionAlias param.Opt[string] `json:"missionAlias,omitzero"`
	// The mission ID the crew is supporting according to the source system.
	MissionID param.Opt[string] `json:"missionId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The type of personnel that comprises the crew (e.g., AIRCREW, MEDCREW, etc.).
	PersonnelType param.Opt[string] `json:"personnelType,omitzero"`
	// Time the crew will be picked up from lodging, in ISO 8601 UTC format with
	// millisecond precision.
	PickupTime param.Opt[time.Time] `json:"pickupTime,omitzero" format:"date-time"`
	// Flag indicating whether post-mission crew rest is applied to the last sortie of
	// a crew's task.
	PostRestApplied param.Opt[bool] `json:"postRestApplied,omitzero"`
	// End time of the crew rest period after the mission, in ISO 8601 UTC format with
	// millisecond precision.
	PostRestEnd param.Opt[time.Time] `json:"postRestEnd,omitzero" format:"date-time"`
	// The scheduled delay or adjustment in the start time of a crew's rest period
	// after a mission, expressed as +/-HH:MM.
	PostRestOffset param.Opt[string] `json:"postRestOffset,omitzero"`
	// Flag indicating whether pre-mission crew rest is applied to the first sortie of
	// a crew's task.
	PreRestApplied param.Opt[bool] `json:"preRestApplied,omitzero"`
	// Start time of the crew rest period before the mission, in ISO 8601 UTC format
	// with millisecond precision.
	PreRestStart param.Opt[time.Time] `json:"preRestStart,omitzero" format:"date-time"`
	// Scheduled return time, in ISO 8601 UTC format with millisecond precision.
	ReturnTime param.Opt[time.Time] `json:"returnTime,omitzero" format:"date-time"`
	// The stage 1 qualifications the crew must have for a mission, such as having
	// basic knowledge of crew operations and aircraft systems.
	Stage1Qual param.Opt[string] `json:"stage1Qual,omitzero"`
	// The stage 2 qualifications the crew must have for a mission, such as completion
	// of advanced mission-specific training.
	Stage2Qual param.Opt[string] `json:"stage2Qual,omitzero"`
	// The stage 3 qualifications the crew must have for a mission, such as full
	// mission-ready certification and the capability of leading complex operations.
	Stage3Qual param.Opt[string] `json:"stage3Qual,omitzero"`
	// Stage name for the crew. A stage is a pool of crews supporting a given operation
	// plan.
	StageName param.Opt[string] `json:"stageName,omitzero"`
	// Time the crew entered the stage, in ISO 8601 UTC format with millisecond
	// precision.
	StageTime param.Opt[time.Time] `json:"stageTime,omitzero" format:"date-time"`
	// Crew status (e.g. NEEDCREW, ASSIGNED, APPROVED, NOTIFIED, PARTIAL, UNKNOWN,
	// etc.).
	Status param.Opt[string] `json:"status,omitzero"`
	// Flag indicating that one or more crew members requires transportation to the
	// departure location.
	TransportReq param.Opt[bool] `json:"transportReq,omitzero"`
	// Identifies the trip kit needed by the crew. A trip kit contains charts,
	// regulations, maps, etc. carried by the crew during missions.
	TripKit param.Opt[string] `json:"tripKit,omitzero"`
	// Array of qualification codes assigned to this crew (e.g., AL for Aircraft
	// Leader, CS for Combat Systems Operator, etc.).
	AssignedQualCode []string `json:"assignedQualCode,omitzero"`
	// CrewMembers Collection.
	CrewMembers []CrewUpdateParamsCrewMember `json:"crewMembers,omitzero"`
	// Array of qualification codes required for this crew (e.g., AL for Aircraft
	// Leader, CS for Combat Systems Operator, etc.).
	ReqQualCode []string `json:"reqQualCode,omitzero"`
	paramObj
}

func (r CrewUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CrewUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrewUpdateParams) UnmarshalJSON(data []byte) error {
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
type CrewUpdateParamsDataMode string

const (
	CrewUpdateParamsDataModeReal      CrewUpdateParamsDataMode = "REAL"
	CrewUpdateParamsDataModeTest      CrewUpdateParamsDataMode = "TEST"
	CrewUpdateParamsDataModeSimulated CrewUpdateParamsDataMode = "SIMULATED"
	CrewUpdateParamsDataModeExercise  CrewUpdateParamsDataMode = "EXERCISE"
)

// Schema for Crew Member data.
type CrewUpdateParamsCrewMember struct {
	// Flag indicating whether this crew member has been alerted of the associated
	// task.
	Alerted param.Opt[bool] `json:"alerted,omitzero"`
	// Flag indicating this crew member is assigned to all sorties of the crew
	// itinerary.
	AllSortie param.Opt[bool] `json:"allSortie,omitzero"`
	// Flag indicating whether this crew member has been approved for the associated
	// task.
	Approved param.Opt[bool] `json:"approved,omitzero"`
	// Flag indicating whether this crew member is attached to his/her squadron. Crew
	// members that are not attached are considered assigned.
	Attached param.Opt[bool] `json:"attached,omitzero"`
	// The military branch assignment of the crew member.
	Branch param.Opt[string] `json:"branch,omitzero"`
	// Flag indicating this crew member is a civilian or non-military person.
	Civilian param.Opt[bool] `json:"civilian,omitzero"`
	// Flag indicating this person is the aircraft commander.
	Commander param.Opt[bool] `json:"commander,omitzero"`
	// The crew position of the crew member.
	CrewPosition param.Opt[string] `json:"crewPosition,omitzero"`
	// The crew member's 10-digit DoD ID number.
	DodID param.Opt[string] `json:"dodID,omitzero"`
	// The duty position of the crew member.
	DutyPosition param.Opt[string] `json:"dutyPosition,omitzero"`
	// The current duty status code of this crew member (e.g., AGR for Active Guard and
	// Reserve, IDT for Inactive Duty Training, etc.).
	DutyStatus param.Opt[string] `json:"dutyStatus,omitzero"`
	// Flag indicating whether this crew member has been notified of an event by email.
	Emailed param.Opt[bool] `json:"emailed,omitzero"`
	// Flag indicating whether this crew member requires an additional amount of time
	// to report for duty.
	ExtraTime param.Opt[bool] `json:"extraTime,omitzero"`
	// The first name of the crew member.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The earliest flying currency expiration date for this crew member, in ISO 8601
	// UTC format with millisecond precision.
	FltCurrencyExp param.Opt[time.Time] `json:"fltCurrencyExp,omitzero" format:"date-time"`
	// The training task identifier associated with the flying currency expiration date
	// for this crew member.
	FltCurrencyExpID param.Opt[string] `json:"fltCurrencyExpId,omitzero"`
	// The date this crew member's records review was completed, in ISO 8601 UTC format
	// with millisecond precision.
	FltRecDate param.Opt[time.Time] `json:"fltRecDate,omitzero" format:"date-time"`
	// The date this crew member's records review is due, in ISO 8601 UTC format with
	// millisecond precision.
	FltRecDue param.Opt[time.Time] `json:"fltRecDue,omitzero" format:"date-time"`
	// The flying squadron assignment of the crew member.
	FlySquadron param.Opt[string] `json:"flySquadron,omitzero"`
	// Flag indicating whether this crew member is funded.
	Funded param.Opt[bool] `json:"funded,omitzero"`
	// Gender of the crew member.
	Gender param.Opt[string] `json:"gender,omitzero"`
	// The earliest ground currency expiration date for this crew member, in ISO 8601
	// UTC format with millisecond precision.
	GndCurrencyExp param.Opt[time.Time] `json:"gndCurrencyExp,omitzero" format:"date-time"`
	// The training task identifier associated with the ground currency expiration date
	// for this crew member.
	GndCurrencyExpID param.Opt[string] `json:"gndCurrencyExpId,omitzero"`
	// Flag indicating whether this crew member is grounded (i.e., his/her duties do
	// not include flying).
	Grounded param.Opt[bool] `json:"grounded,omitzero"`
	// Date when this crew member starts acting as guest help for the squadron, in ISO
	// 8601 UTC format with millisecond precision.
	GuestStart param.Opt[time.Time] `json:"guestStart,omitzero" format:"date-time"`
	// Date when this crew member stops acting as guest help for the squadron, in ISO
	// 8601 UTC format with millisecond precision.
	GuestStop param.Opt[time.Time] `json:"guestStop,omitzero" format:"date-time"`
	// Last four digits of the crew member's social security number.
	Last4Ssn param.Opt[string] `json:"last4SSN,omitzero"`
	// Date of the last flight for this crew member, in ISO 8601 UTC format with
	// millisecond precision.
	LastFltDate param.Opt[time.Time] `json:"lastFltDate,omitzero" format:"date-time"`
	// The last name of the crew member.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// The squadron the crew member has been temporarily loaned to.
	LoanedTo param.Opt[string] `json:"loanedTo,omitzero"`
	// Crew member lodging location.
	Lodging param.Opt[string] `json:"lodging,omitzero"`
	// Time this crew member was actually alerted for the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberActualAlertTime param.Opt[time.Time] `json:"memberActualAlertTime,omitzero" format:"date-time"`
	// Adjusted return time for the crew member, in ISO 8601 UTC format with
	// millisecond precision.
	MemberAdjReturnTime param.Opt[time.Time] `json:"memberAdjReturnTime,omitzero" format:"date-time"`
	// Last name of the crew member's adjusted return time approver.
	MemberAdjReturnTimeApprover param.Opt[string] `json:"memberAdjReturnTimeApprover,omitzero"`
	// Unique identifier of the crew member assigned by the originating source.
	MemberID param.Opt[string] `json:"memberId,omitzero"`
	// Initial start time of the crew member's linked task that was delinked due to
	// mission closure, in ISO 8601 UTC format with millisecond precision.
	MemberInitStartTime param.Opt[time.Time] `json:"memberInitStartTime,omitzero" format:"date-time"`
	// The latest possible time the crew member can legally be alerted for a task, in
	// ISO 8601 UTC format with millisecond precision.
	MemberLastAlertTime param.Opt[time.Time] `json:"memberLastAlertTime,omitzero" format:"date-time"`
	// Time this crew member becomes eligible to be alerted for the mission, in ISO
	// 8601 UTC format with millisecond precision.
	MemberLegalAlertTime param.Opt[time.Time] `json:"memberLegalAlertTime,omitzero" format:"date-time"`
	// Time this crew member will be picked up from lodging, in ISO 8601 UTC format
	// with millisecond precision.
	MemberPickupTime param.Opt[time.Time] `json:"memberPickupTime,omitzero" format:"date-time"`
	// The scheduled delay or adjustment in the start time of a crew member's rest
	// period after a mission, expressed as +/-HH:MM.
	MemberPostRestOffset param.Opt[string] `json:"memberPostRestOffset,omitzero"`
	// End time of this crew member's rest period after the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberPostRestTime param.Opt[time.Time] `json:"memberPostRestTime,omitzero" format:"date-time"`
	// Start time of this crew member's rest period before the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberPreRestTime param.Opt[time.Time] `json:"memberPreRestTime,omitzero" format:"date-time"`
	// Remarks concerning the crew member.
	MemberRemarks param.Opt[string] `json:"memberRemarks,omitzero"`
	// Scheduled return time for this crew member, in ISO 8601 UTC format with
	// millisecond precision.
	MemberReturnTime param.Opt[time.Time] `json:"memberReturnTime,omitzero" format:"date-time"`
	// Time this crew member is scheduled to be alerted for the mission, in ISO 8601
	// UTC format with millisecond precision.
	MemberSchedAlertTime param.Opt[time.Time] `json:"memberSchedAlertTime,omitzero" format:"date-time"`
	// The military component for the crew member (e.g., ACTIVE, RESERVE, GUARD,
	// UNKNOWN, etc.).
	MemberSource param.Opt[string] `json:"memberSource,omitzero"`
	// Stage name for the crew member. A stage is a pool of crews supporting a given
	// operation plan.
	MemberStageName param.Opt[string] `json:"memberStageName,omitzero"`
	// Flag indicating whether this crew member needs transportation to the departure
	// location.
	MemberTransportReq param.Opt[bool] `json:"memberTransportReq,omitzero"`
	// Amplifying details about the crew member type (e.g. RAVEN, FCC, COMCAM, AIRCREW,
	// MEP, OTHER, etc.).
	MemberType param.Opt[string] `json:"memberType,omitzero"`
	// The middle initial of the crew member.
	MiddleInitial param.Opt[string] `json:"middleInitial,omitzero"`
	// Flag indicating whether this crew member has been notified of an event.
	Notified param.Opt[bool] `json:"notified,omitzero"`
	// Crew member lodging phone number.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// Code indicating a crew member's current physical fitness status and whether they
	// are medically cleared to fly (e.g., D for Duties Not Including Flying, E for
	// Physical Overdue, etc.).
	PhysAvCode param.Opt[string] `json:"physAvCode,omitzero"`
	// Code indicating a crew member's physical availabiility status (e.g.,
	// DISQUALIFIED, OVERDUE, etc.).
	PhysAvStatus param.Opt[string] `json:"physAvStatus,omitzero"`
	// Due date for the crew member's physical, in ISO 8601 UTC format with millisecond
	// precision.
	PhysDue param.Opt[time.Time] `json:"physDue,omitzero" format:"date-time"`
	// The rank of the crew member.
	Rank param.Opt[string] `json:"rank,omitzero"`
	// Remark code used to designate attributes of this crew member. For more
	// information, contact the provider source.
	RemarkCode param.Opt[string] `json:"remarkCode,omitzero"`
	// The primary aircraft type for the crew member according to the personnel
	// resource management system indicated in the crewRMS field.
	RmsMds param.Opt[string] `json:"rmsMDS,omitzero"`
	// Time this crew member is required to report for duty before this flight/mission,
	// in ISO 8601 UTC format with millisecond precision.
	ShowTime param.Opt[time.Time] `json:"showTime,omitzero" format:"date-time"`
	// The squadron the crew member serves.
	Squadron param.Opt[string] `json:"squadron,omitzero"`
	// The date this crew member accomplished physiological or altitude chamber
	// training, in ISO 8601 UTC format with millisecond precision.
	TrainingDate param.Opt[time.Time] `json:"trainingDate,omitzero" format:"date-time"`
	// The Mattermost username of this crew member.
	Username param.Opt[string] `json:"username,omitzero"`
	// The wing the crew member serves.
	Wing param.Opt[string] `json:"wing,omitzero"`
	paramObj
}

func (r CrewUpdateParamsCrewMember) MarshalJSON() (data []byte, err error) {
	type shadow CrewUpdateParamsCrewMember
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrewUpdateParamsCrewMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrewListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CrewListParams]'s query parameters as `url.Values`.
func (r CrewListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CrewCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CrewCountParams]'s query parameters as `url.Values`.
func (r CrewCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CrewTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CrewTupleParams]'s query parameters as `url.Values`.
func (r CrewTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CrewUnvalidatedPublishParams struct {
	Body []CrewUnvalidatedPublishParamsBody
	paramObj
}

func (r CrewUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *CrewUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Crew Services.
//
// The properties ClassificationMarking, DataMode, OrigCrewID, Source are required.
type CrewUnvalidatedPublishParamsBody struct {
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
	// Unique identifier of the formed crew provided by the originating source.
	// Provided for systems that require tracking of an internal system generated ID.
	OrigCrewID string `json:"origCrewId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Adjusted return time, in ISO 8601 UTC format with millisecond precision.
	AdjReturnTime param.Opt[time.Time] `json:"adjReturnTime,omitzero" format:"date-time"`
	// Last name of the adjusted return time approver.
	AdjReturnTimeApprover param.Opt[string] `json:"adjReturnTimeApprover,omitzero"`
	// The aircraft Model Design Series designation assigned for this crew.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Time the crew was alerted, in ISO 8601 UTC format with millisecond precision.
	AlertedTime param.Opt[time.Time] `json:"alertedTime,omitzero" format:"date-time"`
	// Type of alert for the crew (e.g., ALPHA for maximum readiness, BRAVO for
	// standby, etc.).
	AlertType param.Opt[string] `json:"alertType,omitzero"`
	// The crew's Aviation Resource Management System (ARMS) unit. If multiple units
	// exist, use the Aircraft Commander's Unit.
	ArmsCrewUnit param.Opt[string] `json:"armsCrewUnit,omitzero"`
	// Unique identifier of the crew commander assigned by the originating source.
	CommanderID param.Opt[string] `json:"commanderId,omitzero"`
	// Last four digits of the crew commander's social security number.
	CommanderLast4Ssn param.Opt[string] `json:"commanderLast4SSN,omitzero"`
	// The name of the crew commander.
	CommanderName param.Opt[string] `json:"commanderName,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Flag indicating whether this crew task takes the crew home and out of the stage.
	CrewHome param.Opt[bool] `json:"crewHome,omitzero"`
	// Name of the formed crew.
	CrewName param.Opt[string] `json:"crewName,omitzero"`
	// The resource management system managing and reporting data on this crew.
	CrewRms param.Opt[string] `json:"crewRMS,omitzero"`
	// The crew's role on the mission (e.g., DEADHEAD, MEDICAL, PRIMARY).
	CrewRole param.Opt[string] `json:"crewRole,omitzero"`
	// The military component that comprises the crew (e.g., ACTIVE, RESERVE, GUARD,
	// MIXED, UNKNOWN, etc.).
	CrewSource param.Opt[string] `json:"crewSource,omitzero"`
	// The squadron the crew serves.
	CrewSquadron param.Opt[string] `json:"crewSquadron,omitzero"`
	// The type of crew required to meet mission objectives (e.g., AIRDROP, AIRLAND,
	// AIR REFUELING, etc.).
	CrewType param.Opt[string] `json:"crewType,omitzero"`
	// The crew's squadron as identified in its resource management system. If the crew
	// is composed of members from multiple units, then the Crew Commander's unit
	// should be indicated as the crew unit.
	CrewUnit param.Opt[string] `json:"crewUnit,omitzero"`
	// The wing the crew serves.
	CrewWing param.Opt[string] `json:"crewWing,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the airfield at
	// which the crew is currently located.
	CurrentIcao param.Opt[string] `json:"currentICAO,omitzero"`
	// Crew Flight Duty Period (FDP) eligibility type.
	FdpEligType param.Opt[string] `json:"fdpEligType,omitzero"`
	// Flight Duty Period (FDP) type.
	FdpType param.Opt[string] `json:"fdpType,omitzero"`
	// The number of female enlisted crew members.
	FemaleEnlistedQty param.Opt[int64] `json:"femaleEnlistedQty,omitzero"`
	// The number of female officer crew members.
	FemaleOfficerQty param.Opt[int64] `json:"femaleOfficerQty,omitzero"`
	// Authorization number used on the flight order.
	FltAuthNum param.Opt[string] `json:"fltAuthNum,omitzero"`
	// Unique identifier of the Site at which the crew is currently located. This ID
	// can be used to obtain additional information on a Site using the 'get by ID'
	// operation (e.g. /udl/site/{id}). For example, the Site object with idSite = abc
	// would be queried as /udl/site/abc.
	IDSiteCurrent param.Opt[string] `json:"idSiteCurrent,omitzero"`
	// Unique identifier of the Aircraft Sortie associated with this crew record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Initial start time of the crew's linked task that was delinked due to mission
	// closure, in ISO 8601 UTC format with millisecond precision.
	InitStartTime param.Opt[time.Time] `json:"initStartTime,omitzero" format:"date-time"`
	// The last time the crew can be alerted, in ISO 8601 UTC format with millisecond
	// precision.
	LastAlertTime param.Opt[time.Time] `json:"lastAlertTime,omitzero" format:"date-time"`
	// Time the crew is legal for alert, in ISO 8601 UTC format with millisecond
	// precision.
	LegalAlertTime param.Opt[time.Time] `json:"legalAlertTime,omitzero" format:"date-time"`
	// Time the crew is legally authorized or scheduled to remain on standby for duty,
	// in ISO 8601 UTC format with millisecond precision.
	LegalBravoTime param.Opt[time.Time] `json:"legalBravoTime,omitzero" format:"date-time"`
	// Flag indicating whether this crew is part of a linked flying task.
	LinkedTask param.Opt[bool] `json:"linkedTask,omitzero"`
	// The number of male enlisted crew members.
	MaleEnlistedQty param.Opt[int64] `json:"maleEnlistedQty,omitzero"`
	// The number of male officer crew members.
	MaleOfficerQty param.Opt[int64] `json:"maleOfficerQty,omitzero"`
	// User-defined alias designation for the mission.
	MissionAlias param.Opt[string] `json:"missionAlias,omitzero"`
	// The mission ID the crew is supporting according to the source system.
	MissionID param.Opt[string] `json:"missionId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The type of personnel that comprises the crew (e.g., AIRCREW, MEDCREW, etc.).
	PersonnelType param.Opt[string] `json:"personnelType,omitzero"`
	// Time the crew will be picked up from lodging, in ISO 8601 UTC format with
	// millisecond precision.
	PickupTime param.Opt[time.Time] `json:"pickupTime,omitzero" format:"date-time"`
	// Flag indicating whether post-mission crew rest is applied to the last sortie of
	// a crew's task.
	PostRestApplied param.Opt[bool] `json:"postRestApplied,omitzero"`
	// End time of the crew rest period after the mission, in ISO 8601 UTC format with
	// millisecond precision.
	PostRestEnd param.Opt[time.Time] `json:"postRestEnd,omitzero" format:"date-time"`
	// The scheduled delay or adjustment in the start time of a crew's rest period
	// after a mission, expressed as +/-HH:MM.
	PostRestOffset param.Opt[string] `json:"postRestOffset,omitzero"`
	// Flag indicating whether pre-mission crew rest is applied to the first sortie of
	// a crew's task.
	PreRestApplied param.Opt[bool] `json:"preRestApplied,omitzero"`
	// Start time of the crew rest period before the mission, in ISO 8601 UTC format
	// with millisecond precision.
	PreRestStart param.Opt[time.Time] `json:"preRestStart,omitzero" format:"date-time"`
	// Scheduled return time, in ISO 8601 UTC format with millisecond precision.
	ReturnTime param.Opt[time.Time] `json:"returnTime,omitzero" format:"date-time"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The stage 1 qualifications the crew must have for a mission, such as having
	// basic knowledge of crew operations and aircraft systems.
	Stage1Qual param.Opt[string] `json:"stage1Qual,omitzero"`
	// The stage 2 qualifications the crew must have for a mission, such as completion
	// of advanced mission-specific training.
	Stage2Qual param.Opt[string] `json:"stage2Qual,omitzero"`
	// The stage 3 qualifications the crew must have for a mission, such as full
	// mission-ready certification and the capability of leading complex operations.
	Stage3Qual param.Opt[string] `json:"stage3Qual,omitzero"`
	// Stage name for the crew. A stage is a pool of crews supporting a given operation
	// plan.
	StageName param.Opt[string] `json:"stageName,omitzero"`
	// Time the crew entered the stage, in ISO 8601 UTC format with millisecond
	// precision.
	StageTime param.Opt[time.Time] `json:"stageTime,omitzero" format:"date-time"`
	// Crew status (e.g. NEEDCREW, ASSIGNED, APPROVED, NOTIFIED, PARTIAL, UNKNOWN,
	// etc.).
	Status param.Opt[string] `json:"status,omitzero"`
	// Flag indicating that one or more crew members requires transportation to the
	// departure location.
	TransportReq param.Opt[bool] `json:"transportReq,omitzero"`
	// Identifies the trip kit needed by the crew. A trip kit contains charts,
	// regulations, maps, etc. carried by the crew during missions.
	TripKit param.Opt[string] `json:"tripKit,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Array of qualification codes assigned to this crew (e.g., AL for Aircraft
	// Leader, CS for Combat Systems Operator, etc.).
	AssignedQualCode []string `json:"assignedQualCode,omitzero"`
	// CrewMembers Collection.
	CrewMembers []CrewUnvalidatedPublishParamsBodyCrewMember `json:"crewMembers,omitzero"`
	// Array of qualification codes required for this crew (e.g., AL for Aircraft
	// Leader, CS for Combat Systems Operator, etc.).
	ReqQualCode []string `json:"reqQualCode,omitzero"`
	paramObj
}

func (r CrewUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow CrewUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrewUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrewUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Schema for Crew Member data.
type CrewUnvalidatedPublishParamsBodyCrewMember struct {
	// Flag indicating whether this crew member has been alerted of the associated
	// task.
	Alerted param.Opt[bool] `json:"alerted,omitzero"`
	// Flag indicating this crew member is assigned to all sorties of the crew
	// itinerary.
	AllSortie param.Opt[bool] `json:"allSortie,omitzero"`
	// Flag indicating whether this crew member has been approved for the associated
	// task.
	Approved param.Opt[bool] `json:"approved,omitzero"`
	// Flag indicating whether this crew member is attached to his/her squadron. Crew
	// members that are not attached are considered assigned.
	Attached param.Opt[bool] `json:"attached,omitzero"`
	// The military branch assignment of the crew member.
	Branch param.Opt[string] `json:"branch,omitzero"`
	// Flag indicating this crew member is a civilian or non-military person.
	Civilian param.Opt[bool] `json:"civilian,omitzero"`
	// Flag indicating this person is the aircraft commander.
	Commander param.Opt[bool] `json:"commander,omitzero"`
	// The crew position of the crew member.
	CrewPosition param.Opt[string] `json:"crewPosition,omitzero"`
	// The crew member's 10-digit DoD ID number.
	DodID param.Opt[string] `json:"dodID,omitzero"`
	// The duty position of the crew member.
	DutyPosition param.Opt[string] `json:"dutyPosition,omitzero"`
	// The current duty status code of this crew member (e.g., AGR for Active Guard and
	// Reserve, IDT for Inactive Duty Training, etc.).
	DutyStatus param.Opt[string] `json:"dutyStatus,omitzero"`
	// Flag indicating whether this crew member has been notified of an event by email.
	Emailed param.Opt[bool] `json:"emailed,omitzero"`
	// Flag indicating whether this crew member requires an additional amount of time
	// to report for duty.
	ExtraTime param.Opt[bool] `json:"extraTime,omitzero"`
	// The first name of the crew member.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The earliest flying currency expiration date for this crew member, in ISO 8601
	// UTC format with millisecond precision.
	FltCurrencyExp param.Opt[time.Time] `json:"fltCurrencyExp,omitzero" format:"date-time"`
	// The training task identifier associated with the flying currency expiration date
	// for this crew member.
	FltCurrencyExpID param.Opt[string] `json:"fltCurrencyExpId,omitzero"`
	// The date this crew member's records review was completed, in ISO 8601 UTC format
	// with millisecond precision.
	FltRecDate param.Opt[time.Time] `json:"fltRecDate,omitzero" format:"date-time"`
	// The date this crew member's records review is due, in ISO 8601 UTC format with
	// millisecond precision.
	FltRecDue param.Opt[time.Time] `json:"fltRecDue,omitzero" format:"date-time"`
	// The flying squadron assignment of the crew member.
	FlySquadron param.Opt[string] `json:"flySquadron,omitzero"`
	// Flag indicating whether this crew member is funded.
	Funded param.Opt[bool] `json:"funded,omitzero"`
	// Gender of the crew member.
	Gender param.Opt[string] `json:"gender,omitzero"`
	// The earliest ground currency expiration date for this crew member, in ISO 8601
	// UTC format with millisecond precision.
	GndCurrencyExp param.Opt[time.Time] `json:"gndCurrencyExp,omitzero" format:"date-time"`
	// The training task identifier associated with the ground currency expiration date
	// for this crew member.
	GndCurrencyExpID param.Opt[string] `json:"gndCurrencyExpId,omitzero"`
	// Flag indicating whether this crew member is grounded (i.e., his/her duties do
	// not include flying).
	Grounded param.Opt[bool] `json:"grounded,omitzero"`
	// Date when this crew member starts acting as guest help for the squadron, in ISO
	// 8601 UTC format with millisecond precision.
	GuestStart param.Opt[time.Time] `json:"guestStart,omitzero" format:"date-time"`
	// Date when this crew member stops acting as guest help for the squadron, in ISO
	// 8601 UTC format with millisecond precision.
	GuestStop param.Opt[time.Time] `json:"guestStop,omitzero" format:"date-time"`
	// Last four digits of the crew member's social security number.
	Last4Ssn param.Opt[string] `json:"last4SSN,omitzero"`
	// Date of the last flight for this crew member, in ISO 8601 UTC format with
	// millisecond precision.
	LastFltDate param.Opt[time.Time] `json:"lastFltDate,omitzero" format:"date-time"`
	// The last name of the crew member.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// The squadron the crew member has been temporarily loaned to.
	LoanedTo param.Opt[string] `json:"loanedTo,omitzero"`
	// Crew member lodging location.
	Lodging param.Opt[string] `json:"lodging,omitzero"`
	// Time this crew member was actually alerted for the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberActualAlertTime param.Opt[time.Time] `json:"memberActualAlertTime,omitzero" format:"date-time"`
	// Adjusted return time for the crew member, in ISO 8601 UTC format with
	// millisecond precision.
	MemberAdjReturnTime param.Opt[time.Time] `json:"memberAdjReturnTime,omitzero" format:"date-time"`
	// Last name of the crew member's adjusted return time approver.
	MemberAdjReturnTimeApprover param.Opt[string] `json:"memberAdjReturnTimeApprover,omitzero"`
	// Unique identifier of the crew member assigned by the originating source.
	MemberID param.Opt[string] `json:"memberId,omitzero"`
	// Initial start time of the crew member's linked task that was delinked due to
	// mission closure, in ISO 8601 UTC format with millisecond precision.
	MemberInitStartTime param.Opt[time.Time] `json:"memberInitStartTime,omitzero" format:"date-time"`
	// The latest possible time the crew member can legally be alerted for a task, in
	// ISO 8601 UTC format with millisecond precision.
	MemberLastAlertTime param.Opt[time.Time] `json:"memberLastAlertTime,omitzero" format:"date-time"`
	// Time this crew member becomes eligible to be alerted for the mission, in ISO
	// 8601 UTC format with millisecond precision.
	MemberLegalAlertTime param.Opt[time.Time] `json:"memberLegalAlertTime,omitzero" format:"date-time"`
	// Time this crew member will be picked up from lodging, in ISO 8601 UTC format
	// with millisecond precision.
	MemberPickupTime param.Opt[time.Time] `json:"memberPickupTime,omitzero" format:"date-time"`
	// The scheduled delay or adjustment in the start time of a crew member's rest
	// period after a mission, expressed as +/-HH:MM.
	MemberPostRestOffset param.Opt[string] `json:"memberPostRestOffset,omitzero"`
	// End time of this crew member's rest period after the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberPostRestTime param.Opt[time.Time] `json:"memberPostRestTime,omitzero" format:"date-time"`
	// Start time of this crew member's rest period before the mission, in ISO 8601 UTC
	// format with millisecond precision.
	MemberPreRestTime param.Opt[time.Time] `json:"memberPreRestTime,omitzero" format:"date-time"`
	// Remarks concerning the crew member.
	MemberRemarks param.Opt[string] `json:"memberRemarks,omitzero"`
	// Scheduled return time for this crew member, in ISO 8601 UTC format with
	// millisecond precision.
	MemberReturnTime param.Opt[time.Time] `json:"memberReturnTime,omitzero" format:"date-time"`
	// Time this crew member is scheduled to be alerted for the mission, in ISO 8601
	// UTC format with millisecond precision.
	MemberSchedAlertTime param.Opt[time.Time] `json:"memberSchedAlertTime,omitzero" format:"date-time"`
	// The military component for the crew member (e.g., ACTIVE, RESERVE, GUARD,
	// UNKNOWN, etc.).
	MemberSource param.Opt[string] `json:"memberSource,omitzero"`
	// Stage name for the crew member. A stage is a pool of crews supporting a given
	// operation plan.
	MemberStageName param.Opt[string] `json:"memberStageName,omitzero"`
	// Flag indicating whether this crew member needs transportation to the departure
	// location.
	MemberTransportReq param.Opt[bool] `json:"memberTransportReq,omitzero"`
	// Amplifying details about the crew member type (e.g. RAVEN, FCC, COMCAM, AIRCREW,
	// MEP, OTHER, etc.).
	MemberType param.Opt[string] `json:"memberType,omitzero"`
	// The middle initial of the crew member.
	MiddleInitial param.Opt[string] `json:"middleInitial,omitzero"`
	// Flag indicating whether this crew member has been notified of an event.
	Notified param.Opt[bool] `json:"notified,omitzero"`
	// Crew member lodging phone number.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// Code indicating a crew member's current physical fitness status and whether they
	// are medically cleared to fly (e.g., D for Duties Not Including Flying, E for
	// Physical Overdue, etc.).
	PhysAvCode param.Opt[string] `json:"physAvCode,omitzero"`
	// Code indicating a crew member's physical availabiility status (e.g.,
	// DISQUALIFIED, OVERDUE, etc.).
	PhysAvStatus param.Opt[string] `json:"physAvStatus,omitzero"`
	// Due date for the crew member's physical, in ISO 8601 UTC format with millisecond
	// precision.
	PhysDue param.Opt[time.Time] `json:"physDue,omitzero" format:"date-time"`
	// The rank of the crew member.
	Rank param.Opt[string] `json:"rank,omitzero"`
	// Remark code used to designate attributes of this crew member. For more
	// information, contact the provider source.
	RemarkCode param.Opt[string] `json:"remarkCode,omitzero"`
	// The primary aircraft type for the crew member according to the personnel
	// resource management system indicated in the crewRMS field.
	RmsMds param.Opt[string] `json:"rmsMDS,omitzero"`
	// Time this crew member is required to report for duty before this flight/mission,
	// in ISO 8601 UTC format with millisecond precision.
	ShowTime param.Opt[time.Time] `json:"showTime,omitzero" format:"date-time"`
	// The squadron the crew member serves.
	Squadron param.Opt[string] `json:"squadron,omitzero"`
	// The date this crew member accomplished physiological or altitude chamber
	// training, in ISO 8601 UTC format with millisecond precision.
	TrainingDate param.Opt[time.Time] `json:"trainingDate,omitzero" format:"date-time"`
	// The Mattermost username of this crew member.
	Username param.Opt[string] `json:"username,omitzero"`
	// The wing the crew member serves.
	Wing param.Opt[string] `json:"wing,omitzero"`
	paramObj
}

func (r CrewUnvalidatedPublishParamsBodyCrewMember) MarshalJSON() (data []byte, err error) {
	type shadow CrewUnvalidatedPublishParamsBodyCrewMember
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrewUnvalidatedPublishParamsBodyCrewMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
