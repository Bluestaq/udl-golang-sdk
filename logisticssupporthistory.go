// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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

// LogisticsSupportHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogisticsSupportHistoryService] method instead.
type LogisticsSupportHistoryService struct {
	Options []option.RequestOption
}

// NewLogisticsSupportHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogisticsSupportHistoryService(opts ...option.RequestOption) (r LogisticsSupportHistoryService) {
	r = LogisticsSupportHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LogisticsSupportHistoryService) List(ctx context.Context, query LogisticsSupportHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LogisticsSupportHistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/logisticssupport/history"
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

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LogisticsSupportHistoryService) ListAutoPaging(ctx context.Context, query LogisticsSupportHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LogisticsSupportHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LogisticsSupportHistoryService) Aodr(ctx context.Context, query LogisticsSupportHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/logisticssupport/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LogisticsSupportHistoryService) Count(ctx context.Context, query LogisticsSupportHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/logisticssupport/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
type LogisticsSupportHistoryListResponse struct {
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
	DataMode LogisticsSupportHistoryListResponseDataMode `json:"dataMode,required"`
	// The time this report was created, in ISO 8601 UTC format with millisecond
	// precision.
	RptCreatedTime time.Time `json:"rptCreatedTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds string `json:"aircraftMDS"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// The current ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	CurrIcao string `json:"currICAO"`
	// The estimated time mission capable for the aircraft, in ISO 8601 UCT format with
	// millisecond precision. This is the estimated time when the aircraft is mission
	// ready.
	Etic time.Time `json:"etic" format:"date-time"`
	// Logistics estimated time mission capable.
	Etmc time.Time `json:"etmc" format:"date-time"`
	// Optional system identifier from external systs. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtSystemID string `json:"extSystemId"`
	// This field identifies the pacing event for bringing the aircraft to Mission
	// Capable status. It is used in calculating the Estimated Time Mission Capable
	// (ETMC) value. Acceptable values are WA (Will Advise), INW (In Work), P+hhh.h
	// (where P=parts and hhh.h is the number of hours up to 999 plus tenths of hours),
	// EQ+hhh.h (EQ=equipment), MRT+hhh.h (MRT=maintenance recovery team).
	LogisticAction string `json:"logisticAction"`
	// Discrepancy information associated with this LogisticsSupport record.
	LogisticsDiscrepancyInfos []LogisticsSupportHistoryListResponseLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID string `json:"logisticsRecordId"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticsSupportHistoryListResponseLogisticsSupportItem `json:"logisticsSupportItems"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticsSupportHistoryListResponseLogisticsTransportationPlan `json:"logisticsTransportationPlans"`
	// The maintenance status code of the aircraft which may be based on pilot
	// descriptions or evaluation codes. Contact the source provider for details.
	MaintStatusCode string `json:"maintStatusCode"`
	// The time indicating when all mission essential problems with a given aircraft
	// have been fixed and is mission capable. This datetime should be in ISO 8601 UTC
	// format with millisecond precision.
	McTime time.Time `json:"mcTime" format:"date-time"`
	// The time indicating when a given aircraft breaks for a mission essential reason.
	// This datetime should be in ISO 8601 UTC format with millisecond precision.
	MeTime time.Time `json:"meTime" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The organization that owns this logistics record.
	Owner string `json:"owner"`
	// This is used to indicate whether a closed master record has been reopened.
	ReopenFlag bool `json:"reopenFlag"`
	// The time this report was closed, in ISO 8601 UTC format with millisecond
	// precision.
	RptClosedTime time.Time `json:"rptClosedTime" format:"date-time"`
	// The supplying ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	SuppIcao string `json:"suppICAO"`
	// The tail number of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	TailNumber string `json:"tailNumber"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking        respjson.Field
		DataMode                     respjson.Field
		RptCreatedTime               respjson.Field
		Source                       respjson.Field
		ID                           respjson.Field
		AircraftMds                  respjson.Field
		CreatedAt                    respjson.Field
		CreatedBy                    respjson.Field
		CurrIcao                     respjson.Field
		Etic                         respjson.Field
		Etmc                         respjson.Field
		ExtSystemID                  respjson.Field
		LogisticAction               respjson.Field
		LogisticsDiscrepancyInfos    respjson.Field
		LogisticsRecordID            respjson.Field
		LogisticsRemarks             respjson.Field
		LogisticsSupportItems        respjson.Field
		LogisticsTransportationPlans respjson.Field
		MaintStatusCode              respjson.Field
		McTime                       respjson.Field
		MeTime                       respjson.Field
		Origin                       respjson.Field
		OrigNetwork                  respjson.Field
		Owner                        respjson.Field
		ReopenFlag                   respjson.Field
		RptClosedTime                respjson.Field
		SuppIcao                     respjson.Field
		TailNumber                   respjson.Field
		UpdatedAt                    respjson.Field
		UpdatedBy                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type LogisticsSupportHistoryListResponseDataMode string

const (
	LogisticsSupportHistoryListResponseDataModeReal      LogisticsSupportHistoryListResponseDataMode = "REAL"
	LogisticsSupportHistoryListResponseDataModeTest      LogisticsSupportHistoryListResponseDataMode = "TEST"
	LogisticsSupportHistoryListResponseDataModeSimulated LogisticsSupportHistoryListResponseDataMode = "SIMULATED"
	LogisticsSupportHistoryListResponseDataModeExercise  LogisticsSupportHistoryListResponseDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticsSupportHistoryListResponseLogisticsDiscrepancyInfo struct {
	// The discrepancy closure time, in ISO 8601 UTC format with millisecond precision.
	ClosureTime time.Time `json:"closureTime" format:"date-time"`
	// The aircraft discrepancy description.
	DiscrepancyInfo string `json:"discrepancyInfo"`
	// Job Control Number of the discrepancy.
	Jcn string `json:"jcn"`
	// The job start time, in ISO 8601 UTC format with millisecond precision.
	JobStTime time.Time `json:"jobStTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClosureTime     respjson.Field
		DiscrepancyInfo respjson.Field
		Jcn             respjson.Field
		JobStTime       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportHistoryListResponseLogisticsDiscrepancyInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportHistoryListResponseLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticsSupportHistoryListResponseLogisticsSupportItem struct {
	// This element indicates whether or not the supplied item is contained within
	// another item.
	Cannibalized bool `json:"cannibalized"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	DeployPlanNumber string `json:"deployPlanNumber"`
	// The technical order name of the part ordered.
	Description string `json:"description"`
	// The last time this supported item was updated, in ISO 8601 UTC format with
	// millisecond precision.
	ItemLastChangedDate time.Time `json:"itemLastChangedDate" format:"date-time"`
	// A number assigned by Job Control to monitor and record maintenance actions
	// required to correct the associated aircraft maintenance discrepancy. It is
	// seven, nine or twelve characters, depending on the base-specific numbering
	// scheme. If seven characters: characters 1-3 are Julian date, 4-7 are sequence
	// numbers. If nine characters: characters 1-2 are last two digits of the year,
	// characters 3-5 are Julian date, 6-9 are sequence numbers. If twelve characters:
	// characters 1-2 are last two digits of the year, 3-5 are Julian date, 6-9 are
	// sequence numbers, and 10-12 are a three-digit supplemental number.
	JobControlNumber string `json:"jobControlNumber"`
	// The parts associated with this support item.
	LogisticsParts []LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsPart `json:"logisticsParts"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties"`
	// Military aircraft discrepancy logistics requisition ordered quantity. The
	// quantity of equipment ordered that is required to fix the aircraft.
	Quantity int64 `json:"quantity"`
	// The time the item is ready, in ISO 8601 UTC format with millisecond precision.
	ReadyTime time.Time `json:"readyTime" format:"date-time"`
	// The time the item is received, in ISO 8601 UTC format with millisecond
	// precision.
	ReceivedTime time.Time `json:"receivedTime" format:"date-time"`
	// The type of recovery request needed. Contact the source provider for details.
	RecoveryRequestTypeCode string `json:"recoveryRequestTypeCode"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	RedeployPlanNumber string `json:"redeployPlanNumber"`
	// This is the Redeploy (return) Transportation Control Number/Tracking Reference
	// Number for the selected item.
	RedeployShipmentUnitID string `json:"redeployShipmentUnitId"`
	// The request or record number for this item type (Equipent, Part, or MRT).
	RequestNumber string `json:"requestNumber"`
	// This element indicates if the supplied item is characterized as additional
	// support.
	ResupportFlag bool `json:"resupportFlag"`
	// Shipment Unit Identifier is the Transportation Control Number (TCN) for shipping
	// that piece of equipment being requested.
	ShipmentUnitID string `json:"shipmentUnitId"`
	// The point of contact is a free text field to add information about the
	// individual(s) with knowledge of the referenced requested or supplied item(s).
	// The default value for this field is the last name, first name, and middle
	// initial of the operator who created the records and/or generated the
	// transaction.
	SiPoc string `json:"siPOC"`
	// The code that represents the International Civil Aviation Organization (ICAO)
	// designations of an airport.
	SourceIcao string `json:"sourceICAO"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cannibalized            respjson.Field
		DeployPlanNumber        respjson.Field
		Description             respjson.Field
		ItemLastChangedDate     respjson.Field
		JobControlNumber        respjson.Field
		LogisticsParts          respjson.Field
		LogisticsRemarks        respjson.Field
		LogisticsSpecialties    respjson.Field
		Quantity                respjson.Field
		ReadyTime               respjson.Field
		ReceivedTime            respjson.Field
		RecoveryRequestTypeCode respjson.Field
		RedeployPlanNumber      respjson.Field
		RedeployShipmentUnitID  respjson.Field
		RequestNumber           respjson.Field
		ResupportFlag           respjson.Field
		ShipmentUnitID          respjson.Field
		SiPoc                   respjson.Field
		SourceIcao              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportHistoryListResponseLogisticsSupportItem) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportHistoryListResponseLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber string `json:"figureNumber"`
	// Technical order manual index number for the requested part.
	IndexNumber string `json:"indexNumber"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier string `json:"locationVerifier"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks"`
	// Code for a unit of measurement.
	MeasurementUnitCode string `json:"measurementUnitCode"`
	// The National Stock Number of the part being requested or supplied.
	NationalStockNumber string `json:"nationalStockNumber"`
	// Requested or supplied part number.
	PartNumber string `json:"partNumber"`
	// The person who validated the request for parts.
	RequestVerifier string `json:"requestVerifier"`
	// The supply document number.
	SupplyDocumentNumber string `json:"supplyDocumentNumber"`
	// Indicates the specified Technical Order manual holding the aircraft information
	// for use in diagnosing a problem or condition.
	TechnicalOrderText string `json:"technicalOrderText"`
	// Work Unit Code (WUC), or for some aircraft types, the Reference Designator.
	WorkUnitCode string `json:"workUnitCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FigureNumber         respjson.Field
		IndexNumber          respjson.Field
		LocationVerifier     respjson.Field
		LogisticsStocks      respjson.Field
		MeasurementUnitCode  respjson.Field
		NationalStockNumber  respjson.Field
		PartNumber           respjson.Field
		RequestVerifier      respjson.Field
		SupplyDocumentNumber respjson.Field
		TechnicalOrderText   respjson.Field
		WorkUnitCode         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsPart) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsPartLogisticsStock struct {
	// The quantity of available parts needed from sourceICAO.
	Quantity int64 `json:"quantity"`
	// The ICAO code for the primary location with available parts.
	SourceIcao string `json:"sourceICAO"`
	// The datetime when the parts were sourced, in ISO 8601 UTC format with
	// millisecond precision.
	StockCheckTime time.Time `json:"stockCheckTime" format:"date-time"`
	// The point of contact at the sourced location.
	StockPoc string `json:"stockPOC"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity       respjson.Field
		SourceIcao     respjson.Field
		StockCheckTime respjson.Field
		StockPoc       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsPartLogisticsStock) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsSpecialty struct {
	// The first name of the specialist.
	FirstName string `json:"firstName"`
	// The last four digits of the specialist's social security number.
	Last4Ssn string `json:"last4Ssn"`
	// The last name of the specialist.
	LastName string `json:"lastName"`
	// Military service rank designation.
	RankCode string `json:"rankCode"`
	// Type code that determines role of the mission response team member. TC - Team
	// Chief, TM - Team Member.
	RoleTypeCode string `json:"roleTypeCode"`
	// Skill level of the mission response team member.
	SkillLevel int64 `json:"skillLevel"`
	// Indicates where the repairs will be performed, or which shop specialty has been
	// assigned responsibility for correcting the discrepancy. Shop specialties are
	// normally listed in abbreviated format.
	Specialty string `json:"specialty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName    respjson.Field
		Last4Ssn     respjson.Field
		LastName     respjson.Field
		RankCode     respjson.Field
		RoleTypeCode respjson.Field
		SkillLevel   respjson.Field
		Specialty    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsSpecialty) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportHistoryListResponseLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticsSupportHistoryListResponseLogisticsTransportationPlan struct {
	// Actual time of departure of first segment, in ISO 8601 UTC format with
	// millisecond precision.
	ActDepTime time.Time `json:"actDepTime" format:"date-time"`
	// These are the initial maintenance values entered based on the pilot descriptions
	// or the official maintenance evaluation code.
	AircraftStatus string `json:"aircraftStatus"`
	// Approximate time of arrival of final segement, in ISO 8601 UTC format with
	// millisecond precision.
	ApproxArrTime time.Time `json:"approxArrTime" format:"date-time"`
	// GC. LGTP_CANX_DT. GD2: Date when the transportation plan was cancelled, in ISO
	// 8601 UTC format with millisecond precision.
	CancelledDate time.Time `json:"cancelledDate" format:"date-time"`
	// GC. LGTP_CLSD_DT. GD2: Date when the transportation plan was closed, in ISO 8601
	// UTC format with millisecond precision.
	ClosedDate time.Time `json:"closedDate" format:"date-time"`
	// The AMS username of the operator who alters the coordination status.
	// Automatically captured by the system.
	Coordinator string `json:"coordinator"`
	// The AMS user unit_id of the operator who alters the coordination status.
	// Automatically captured by the system from table AMS_USER.
	CoordinatorUnit string `json:"coordinatorUnit"`
	// Destination location ICAO.
	DestinationIcao string `json:"destinationICAO"`
	// Transportation plan duration, expressed in the format MMM:SS.
	Duration string `json:"duration"`
	// ETA of the final segment, in ISO 8601 UTC format with millisecond precision.
	EstArrTime time.Time `json:"estArrTime" format:"date-time"`
	// ETD of the first segment, in ISO 8601 UTC format with millisecond precision.
	EstDepTime time.Time `json:"estDepTime" format:"date-time"`
	// Last time transportation plan was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastChangedDate time.Time `json:"lastChangedDate" format:"date-time"`
	// The identifier that represents a Logistics Master Record.
	LogisticMasterRecordID string `json:"logisticMasterRecordId"`
	// The transportation segments associated with this transportation plan.
	LogisticsSegments []LogisticsSupportHistoryListResponseLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticsRemarksFull `json:"logisticsTransportationPlansRemarks"`
	// The major command for the current unit.
	Majcom string `json:"majcom"`
	// Indicates whether there have been changes to changes to ICAOs, estArrTime, or
	// estDepTime since this Transportation Plan was last edited.
	MissionChange bool `json:"missionChange"`
	// Transportation plan enroute stops.
	NumEnrouteStops int64 `json:"numEnrouteStops"`
	// The number of transloads for this Transportation Plan.
	NumTransLoads int64 `json:"numTransLoads"`
	// The origin location.
	OriginIcao string `json:"originICAO"`
	// Defines the transporation plan as either a deployment or redeployment.
	PlanDefinition string `json:"planDefinition"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	PlansNumber string `json:"plansNumber"`
	// GDSS2 uses an 8 character serial number to uniquely identify the aircraft and
	// MDS combination. This is a portion of the full manufacturer serial number.
	SerialNumber string `json:"serialNumber"`
	// Transporation Coordination status code. Cancel, Send to APCC, working, agree,
	// disapprove or blank.
	StatusCode string `json:"statusCode"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	TpAircraftMds string `json:"tpAircraftMDS"`
	// Contains the tail number displayed by GDSS2.
	TpTailNumber string `json:"tpTailNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActDepTime                          respjson.Field
		AircraftStatus                      respjson.Field
		ApproxArrTime                       respjson.Field
		CancelledDate                       respjson.Field
		ClosedDate                          respjson.Field
		Coordinator                         respjson.Field
		CoordinatorUnit                     respjson.Field
		DestinationIcao                     respjson.Field
		Duration                            respjson.Field
		EstArrTime                          respjson.Field
		EstDepTime                          respjson.Field
		LastChangedDate                     respjson.Field
		LogisticMasterRecordID              respjson.Field
		LogisticsSegments                   respjson.Field
		LogisticsTransportationPlansRemarks respjson.Field
		Majcom                              respjson.Field
		MissionChange                       respjson.Field
		NumEnrouteStops                     respjson.Field
		NumTransLoads                       respjson.Field
		OriginIcao                          respjson.Field
		PlanDefinition                      respjson.Field
		PlansNumber                         respjson.Field
		SerialNumber                        respjson.Field
		StatusCode                          respjson.Field
		TpAircraftMds                       respjson.Field
		TpTailNumber                        respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportHistoryListResponseLogisticsTransportationPlan) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportHistoryListResponseLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportHistoryListResponseLogisticsTransportationPlanLogisticsSegment struct {
	// Airport ICAO arrival code.
	ArrivalIcao string `json:"arrivalICAO"`
	// Airport ICAO departure code.
	DepartureIcao string `json:"departureICAO"`
	// The GDSS mission ID for this segment.
	ExtMissionID string `json:"extMissionId"`
	// The unique identifier of the mission to which this logistics record is assigned.
	IDMission string `json:"idMission"`
	// Start air mission itinerary point identifier.
	Itin int64 `json:"itin"`
	// The user generated identifier for an air mission subgroup.
	MissionNumber string `json:"missionNumber"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	MissionType string `json:"missionType"`
	// Transportation mode. AMC airlift, Commercial airlift, Other, or surface
	// transportation.
	ModeCode string `json:"modeCode"`
	// Actual arrival time to segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActArrTime time.Time `json:"segActArrTime" format:"date-time"`
	// Actual departure time to the segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActDepTime time.Time `json:"segActDepTime" format:"date-time"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	SegAircraftMds string `json:"segAircraftMDS"`
	// GC. LGTPS_C_DT_EST_ARR. GD2: Estimated arrival time to the segment destination.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETA
	// is derived from the Mission segment destination point. This datetime should be
	// in ISO 8601 UTC format with millisecond precision.
	SegEstArrTime time.Time `json:"segEstArrTime" format:"date-time"`
	// GC. LGTPS_C_DT_EST_DEP. GD2: Estimated departure time from the segment origin.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETD
	// is derived from the Mission segment origin point. This datetime should be in ISO
	// 8601 UTC format with millisecond precision.
	SegEstDepTime time.Time `json:"segEstDepTime" format:"date-time"`
	// Used to sequence the segments in the transportation plan.
	SegmentNumber int64 `json:"segmentNumber"`
	// The identifier that represents a specific aircraft within an aircraft type.
	SegTailNumber string `json:"segTailNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArrivalIcao    respjson.Field
		DepartureIcao  respjson.Field
		ExtMissionID   respjson.Field
		IDMission      respjson.Field
		Itin           respjson.Field
		MissionNumber  respjson.Field
		MissionType    respjson.Field
		ModeCode       respjson.Field
		SegActArrTime  respjson.Field
		SegActDepTime  respjson.Field
		SegAircraftMds respjson.Field
		SegEstArrTime  respjson.Field
		SegEstDepTime  respjson.Field
		SegmentNumber  respjson.Field
		SegTailNumber  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportHistoryListResponseLogisticsTransportationPlanLogisticsSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportHistoryListResponseLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogisticsSupportHistoryListParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogisticsSupportHistoryListParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticsSupportHistoryAodrParams struct {
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

// URLQuery serializes [LogisticsSupportHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticsSupportHistoryCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogisticsSupportHistoryCountParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
