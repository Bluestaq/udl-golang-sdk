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

// LogisticsSupportService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogisticsSupportService] method instead.
type LogisticsSupportService struct {
	Options []option.RequestOption
	History LogisticsSupportHistoryService
}

// NewLogisticsSupportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogisticsSupportService(opts ...option.RequestOption) (r LogisticsSupportService) {
	r = LogisticsSupportService{}
	r.Options = opts
	r.History = NewLogisticsSupportHistoryService(opts...)
	return
}

// Service operation to take a single LogisticsSupport record as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *LogisticsSupportService) New(ctx context.Context, body LogisticsSupportNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/logisticssupport"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single LogisticsSupport record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *LogisticsSupportService) Update(ctx context.Context, id string, body LogisticsSupportUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/logisticssupport/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LogisticsSupportService) List(ctx context.Context, query LogisticsSupportListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LogisticsSupportListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/logisticssupport"
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
func (r *LogisticsSupportService) ListAutoPaging(ctx context.Context, query LogisticsSupportListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LogisticsSupportListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LogisticsSupportService) Count(ctx context.Context, query LogisticsSupportCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/logisticssupport/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// LogisticsSupport records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *LogisticsSupportService) NewBulk(ctx context.Context, body LogisticsSupportNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/logisticssupport/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single LogisticsSupport record by its unique ID
// passed as a path parameter.
func (r *LogisticsSupportService) Get(ctx context.Context, id string, query LogisticsSupportGetParams, opts ...option.RequestOption) (res *LogisticsSupportGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/logisticssupport/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *LogisticsSupportService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *LogisticsSupportQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/logisticssupport/queryhelp"
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
func (r *LogisticsSupportService) Tuple(ctx context.Context, query LogisticsSupportTupleParams, opts ...option.RequestOption) (res *[]LogisticsSupportTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/logisticssupport/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple logisticssupport records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LogisticsSupportService) UnvalidatedPublish(ctx context.Context, body LogisticsSupportUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-logisticssupport"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Remarks associated with this LogisticsSupport record.
type LogisticsRemarksFull struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged time.Time `json:"lastChanged" format:"date-time"`
	// Text of the remark.
	Remark string `json:"remark"`
	// User who published the remark.
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastChanged respjson.Field
		Remark      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsRemarksFull) RawJSON() string { return r.JSON.raw }
func (r *LogisticsRemarksFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
type LogisticsSupportListResponse struct {
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
	DataMode LogisticsSupportListResponseDataMode `json:"dataMode,required"`
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
	LogisticsDiscrepancyInfos []LogisticsSupportListResponseLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID string `json:"logisticsRecordId"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsSupportListResponseLogisticsRemark `json:"logisticsRemarks"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticsSupportListResponseLogisticsSupportItem `json:"logisticsSupportItems"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticsSupportListResponseLogisticsTransportationPlan `json:"logisticsTransportationPlans"`
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
func (r LogisticsSupportListResponse) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportListResponse) UnmarshalJSON(data []byte) error {
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
type LogisticsSupportListResponseDataMode string

const (
	LogisticsSupportListResponseDataModeReal      LogisticsSupportListResponseDataMode = "REAL"
	LogisticsSupportListResponseDataModeTest      LogisticsSupportListResponseDataMode = "TEST"
	LogisticsSupportListResponseDataModeSimulated LogisticsSupportListResponseDataMode = "SIMULATED"
	LogisticsSupportListResponseDataModeExercise  LogisticsSupportListResponseDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticsSupportListResponseLogisticsDiscrepancyInfo struct {
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
func (r LogisticsSupportListResponseLogisticsDiscrepancyInfo) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportListResponseLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportListResponseLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged time.Time `json:"lastChanged" format:"date-time"`
	// Text of the remark.
	Remark string `json:"remark"`
	// User who published the remark.
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastChanged respjson.Field
		Remark      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportListResponseLogisticsRemark) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportListResponseLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticsSupportListResponseLogisticsSupportItem struct {
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
	LogisticsParts []LogisticsSupportListResponseLogisticsSupportItemLogisticsPart `json:"logisticsParts"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsSupportListResponseLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticsSupportListResponseLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties"`
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
func (r LogisticsSupportListResponseLogisticsSupportItem) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportListResponseLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticsSupportListResponseLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber string `json:"figureNumber"`
	// Technical order manual index number for the requested part.
	IndexNumber string `json:"indexNumber"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier string `json:"locationVerifier"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticsSupportListResponseLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks"`
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
func (r LogisticsSupportListResponseLogisticsSupportItemLogisticsPart) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportListResponseLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticsSupportListResponseLogisticsSupportItemLogisticsPartLogisticsStock struct {
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
func (r LogisticsSupportListResponseLogisticsSupportItemLogisticsPartLogisticsStock) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportListResponseLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportListResponseLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged time.Time `json:"lastChanged" format:"date-time"`
	// Text of the remark.
	Remark string `json:"remark"`
	// User who published the remark.
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastChanged respjson.Field
		Remark      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportListResponseLogisticsSupportItemLogisticsRemark) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportListResponseLogisticsSupportItemLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticsSupportListResponseLogisticsSupportItemLogisticsSpecialty struct {
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
func (r LogisticsSupportListResponseLogisticsSupportItemLogisticsSpecialty) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportListResponseLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticsSupportListResponseLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticsSupportListResponseLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticsSupportListResponseLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks"`
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
func (r LogisticsSupportListResponseLogisticsTransportationPlan) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportListResponseLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportListResponseLogisticsTransportationPlanLogisticsSegment struct {
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
func (r LogisticsSupportListResponseLogisticsTransportationPlanLogisticsSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportListResponseLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportListResponseLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged time.Time `json:"lastChanged" format:"date-time"`
	// Text of the remark.
	Remark string `json:"remark"`
	// User who published the remark.
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastChanged respjson.Field
		Remark      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticsSupportListResponseLogisticsTransportationPlanLogisticsTransportationPlansRemark) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportListResponseLogisticsTransportationPlanLogisticsTransportationPlansRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
type LogisticsSupportGetResponse struct {
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
	DataMode LogisticsSupportGetResponseDataMode `json:"dataMode,required"`
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
	LogisticsDiscrepancyInfos []LogisticsSupportGetResponseLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID string `json:"logisticsRecordId"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticsSupportGetResponseLogisticsSupportItem `json:"logisticsSupportItems"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticsSupportGetResponseLogisticsTransportationPlan `json:"logisticsTransportationPlans"`
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
func (r LogisticsSupportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportGetResponse) UnmarshalJSON(data []byte) error {
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
type LogisticsSupportGetResponseDataMode string

const (
	LogisticsSupportGetResponseDataModeReal      LogisticsSupportGetResponseDataMode = "REAL"
	LogisticsSupportGetResponseDataModeTest      LogisticsSupportGetResponseDataMode = "TEST"
	LogisticsSupportGetResponseDataModeSimulated LogisticsSupportGetResponseDataMode = "SIMULATED"
	LogisticsSupportGetResponseDataModeExercise  LogisticsSupportGetResponseDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticsSupportGetResponseLogisticsDiscrepancyInfo struct {
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
func (r LogisticsSupportGetResponseLogisticsDiscrepancyInfo) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportGetResponseLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticsSupportGetResponseLogisticsSupportItem struct {
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
	LogisticsParts []LogisticsSupportGetResponseLogisticsSupportItemLogisticsPart `json:"logisticsParts"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticsSupportGetResponseLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties"`
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
func (r LogisticsSupportGetResponseLogisticsSupportItem) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportGetResponseLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticsSupportGetResponseLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber string `json:"figureNumber"`
	// Technical order manual index number for the requested part.
	IndexNumber string `json:"indexNumber"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier string `json:"locationVerifier"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticsSupportGetResponseLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks"`
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
func (r LogisticsSupportGetResponseLogisticsSupportItemLogisticsPart) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportGetResponseLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticsSupportGetResponseLogisticsSupportItemLogisticsPartLogisticsStock struct {
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
func (r LogisticsSupportGetResponseLogisticsSupportItemLogisticsPartLogisticsStock) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportGetResponseLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticsSupportGetResponseLogisticsSupportItemLogisticsSpecialty struct {
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
func (r LogisticsSupportGetResponseLogisticsSupportItemLogisticsSpecialty) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportGetResponseLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticsSupportGetResponseLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticsSupportGetResponseLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments"`
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
func (r LogisticsSupportGetResponseLogisticsTransportationPlan) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportGetResponseLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportGetResponseLogisticsTransportationPlanLogisticsSegment struct {
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
func (r LogisticsSupportGetResponseLogisticsTransportationPlanLogisticsSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportGetResponseLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogisticsSupportQueryhelpResponse struct {
	AodrSupported         bool                                         `json:"aodrSupported"`
	ClassificationMarking string                                       `json:"classificationMarking"`
	Description           string                                       `json:"description"`
	HistorySupported      bool                                         `json:"historySupported"`
	Name                  string                                       `json:"name"`
	Parameters            []LogisticsSupportQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                     `json:"requiredRoles"`
	RestSupported         bool                                         `json:"restSupported"`
	SortSupported         bool                                         `json:"sortSupported"`
	TypeName              string                                       `json:"typeName"`
	Uri                   string                                       `json:"uri"`
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
func (r LogisticsSupportQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogisticsSupportQueryhelpResponseParameter struct {
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
func (r LogisticsSupportQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
type LogisticsSupportTupleResponse struct {
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
	DataMode LogisticsSupportTupleResponseDataMode `json:"dataMode,required"`
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
	LogisticsDiscrepancyInfos []LogisticsSupportTupleResponseLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID string `json:"logisticsRecordId"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticsSupportTupleResponseLogisticsSupportItem `json:"logisticsSupportItems"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticsSupportTupleResponseLogisticsTransportationPlan `json:"logisticsTransportationPlans"`
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
func (r LogisticsSupportTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportTupleResponse) UnmarshalJSON(data []byte) error {
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
type LogisticsSupportTupleResponseDataMode string

const (
	LogisticsSupportTupleResponseDataModeReal      LogisticsSupportTupleResponseDataMode = "REAL"
	LogisticsSupportTupleResponseDataModeTest      LogisticsSupportTupleResponseDataMode = "TEST"
	LogisticsSupportTupleResponseDataModeSimulated LogisticsSupportTupleResponseDataMode = "SIMULATED"
	LogisticsSupportTupleResponseDataModeExercise  LogisticsSupportTupleResponseDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticsSupportTupleResponseLogisticsDiscrepancyInfo struct {
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
func (r LogisticsSupportTupleResponseLogisticsDiscrepancyInfo) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportTupleResponseLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticsSupportTupleResponseLogisticsSupportItem struct {
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
	LogisticsParts []LogisticsSupportTupleResponseLogisticsSupportItemLogisticsPart `json:"logisticsParts"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticsSupportTupleResponseLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties"`
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
func (r LogisticsSupportTupleResponseLogisticsSupportItem) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportTupleResponseLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticsSupportTupleResponseLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber string `json:"figureNumber"`
	// Technical order manual index number for the requested part.
	IndexNumber string `json:"indexNumber"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier string `json:"locationVerifier"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticsSupportTupleResponseLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks"`
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
func (r LogisticsSupportTupleResponseLogisticsSupportItemLogisticsPart) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportTupleResponseLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticsSupportTupleResponseLogisticsSupportItemLogisticsPartLogisticsStock struct {
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
func (r LogisticsSupportTupleResponseLogisticsSupportItemLogisticsPartLogisticsStock) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportTupleResponseLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticsSupportTupleResponseLogisticsSupportItemLogisticsSpecialty struct {
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
func (r LogisticsSupportTupleResponseLogisticsSupportItemLogisticsSpecialty) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportTupleResponseLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticsSupportTupleResponseLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticsSupportTupleResponseLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments"`
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
func (r LogisticsSupportTupleResponseLogisticsTransportationPlan) RawJSON() string { return r.JSON.raw }
func (r *LogisticsSupportTupleResponseLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportTupleResponseLogisticsTransportationPlanLogisticsSegment struct {
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
func (r LogisticsSupportTupleResponseLogisticsTransportationPlanLogisticsSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticsSupportTupleResponseLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogisticsSupportNewParams struct {
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
	DataMode LogisticsSupportNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The time this report was created, in ISO 8601 UTC format with millisecond
	// precision.
	RptCreatedTime time.Time `json:"rptCreatedTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// The current ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	CurrIcao param.Opt[string] `json:"currICAO,omitzero"`
	// The estimated time mission capable for the aircraft, in ISO 8601 UCT format with
	// millisecond precision. This is the estimated time when the aircraft is mission
	// ready.
	Etic param.Opt[time.Time] `json:"etic,omitzero" format:"date-time"`
	// Logistics estimated time mission capable.
	Etmc param.Opt[time.Time] `json:"etmc,omitzero" format:"date-time"`
	// Optional system identifier from external systs. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtSystemID param.Opt[string] `json:"extSystemId,omitzero"`
	// This field identifies the pacing event for bringing the aircraft to Mission
	// Capable status. It is used in calculating the Estimated Time Mission Capable
	// (ETMC) value. Acceptable values are WA (Will Advise), INW (In Work), P+hhh.h
	// (where P=parts and hhh.h is the number of hours up to 999 plus tenths of hours),
	// EQ+hhh.h (EQ=equipment), MRT+hhh.h (MRT=maintenance recovery team).
	LogisticAction param.Opt[string] `json:"logisticAction,omitzero"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID param.Opt[string] `json:"logisticsRecordId,omitzero"`
	// The maintenance status code of the aircraft which may be based on pilot
	// descriptions or evaluation codes. Contact the source provider for details.
	MaintStatusCode param.Opt[string] `json:"maintStatusCode,omitzero"`
	// The time indicating when all mission essential problems with a given aircraft
	// have been fixed and is mission capable. This datetime should be in ISO 8601 UTC
	// format with millisecond precision.
	McTime param.Opt[time.Time] `json:"mcTime,omitzero" format:"date-time"`
	// The time indicating when a given aircraft breaks for a mission essential reason.
	// This datetime should be in ISO 8601 UTC format with millisecond precision.
	MeTime param.Opt[time.Time] `json:"meTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The organization that owns this logistics record.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// This is used to indicate whether a closed master record has been reopened.
	ReopenFlag param.Opt[bool] `json:"reopenFlag,omitzero"`
	// The time this report was closed, in ISO 8601 UTC format with millisecond
	// precision.
	RptClosedTime param.Opt[time.Time] `json:"rptClosedTime,omitzero" format:"date-time"`
	// The supplying ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	SuppIcao param.Opt[string] `json:"suppICAO,omitzero"`
	// The tail number of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// Discrepancy information associated with this LogisticsSupport record.
	LogisticsDiscrepancyInfos []LogisticsSupportNewParamsLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos,omitzero"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsSupportNewParamsLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticsSupportNewParamsLogisticsSupportItem `json:"logisticsSupportItems,omitzero"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticsSupportNewParamsLogisticsTransportationPlan `json:"logisticsTransportationPlans,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParams) UnmarshalJSON(data []byte) error {
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
type LogisticsSupportNewParamsDataMode string

const (
	LogisticsSupportNewParamsDataModeReal      LogisticsSupportNewParamsDataMode = "REAL"
	LogisticsSupportNewParamsDataModeTest      LogisticsSupportNewParamsDataMode = "TEST"
	LogisticsSupportNewParamsDataModeSimulated LogisticsSupportNewParamsDataMode = "SIMULATED"
	LogisticsSupportNewParamsDataModeExercise  LogisticsSupportNewParamsDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticsSupportNewParamsLogisticsDiscrepancyInfo struct {
	// The discrepancy closure time, in ISO 8601 UTC format with millisecond precision.
	ClosureTime param.Opt[time.Time] `json:"closureTime,omitzero" format:"date-time"`
	// The aircraft discrepancy description.
	DiscrepancyInfo param.Opt[string] `json:"discrepancyInfo,omitzero"`
	// Job Control Number of the discrepancy.
	Jcn param.Opt[string] `json:"jcn,omitzero"`
	// The job start time, in ISO 8601 UTC format with millisecond precision.
	JobStTime param.Opt[time.Time] `json:"jobStTime,omitzero" format:"date-time"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsDiscrepancyInfo) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsDiscrepancyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportNewParamsLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticsSupportNewParamsLogisticsSupportItem struct {
	// This element indicates whether or not the supplied item is contained within
	// another item.
	Cannibalized param.Opt[bool] `json:"cannibalized,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	DeployPlanNumber param.Opt[string] `json:"deployPlanNumber,omitzero"`
	// The technical order name of the part ordered.
	Description param.Opt[string] `json:"description,omitzero"`
	// The last time this supported item was updated, in ISO 8601 UTC format with
	// millisecond precision.
	ItemLastChangedDate param.Opt[time.Time] `json:"itemLastChangedDate,omitzero" format:"date-time"`
	// A number assigned by Job Control to monitor and record maintenance actions
	// required to correct the associated aircraft maintenance discrepancy. It is
	// seven, nine or twelve characters, depending on the base-specific numbering
	// scheme. If seven characters: characters 1-3 are Julian date, 4-7 are sequence
	// numbers. If nine characters: characters 1-2 are last two digits of the year,
	// characters 3-5 are Julian date, 6-9 are sequence numbers. If twelve characters:
	// characters 1-2 are last two digits of the year, 3-5 are Julian date, 6-9 are
	// sequence numbers, and 10-12 are a three-digit supplemental number.
	JobControlNumber param.Opt[string] `json:"jobControlNumber,omitzero"`
	// Military aircraft discrepancy logistics requisition ordered quantity. The
	// quantity of equipment ordered that is required to fix the aircraft.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The time the item is ready, in ISO 8601 UTC format with millisecond precision.
	ReadyTime param.Opt[time.Time] `json:"readyTime,omitzero" format:"date-time"`
	// The time the item is received, in ISO 8601 UTC format with millisecond
	// precision.
	ReceivedTime param.Opt[time.Time] `json:"receivedTime,omitzero" format:"date-time"`
	// The type of recovery request needed. Contact the source provider for details.
	RecoveryRequestTypeCode param.Opt[string] `json:"recoveryRequestTypeCode,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	RedeployPlanNumber param.Opt[string] `json:"redeployPlanNumber,omitzero"`
	// This is the Redeploy (return) Transportation Control Number/Tracking Reference
	// Number for the selected item.
	RedeployShipmentUnitID param.Opt[string] `json:"redeployShipmentUnitId,omitzero"`
	// The request or record number for this item type (Equipent, Part, or MRT).
	RequestNumber param.Opt[string] `json:"requestNumber,omitzero"`
	// This element indicates if the supplied item is characterized as additional
	// support.
	ResupportFlag param.Opt[bool] `json:"resupportFlag,omitzero"`
	// Shipment Unit Identifier is the Transportation Control Number (TCN) for shipping
	// that piece of equipment being requested.
	ShipmentUnitID param.Opt[string] `json:"shipmentUnitId,omitzero"`
	// The point of contact is a free text field to add information about the
	// individual(s) with knowledge of the referenced requested or supplied item(s).
	// The default value for this field is the last name, first name, and middle
	// initial of the operator who created the records and/or generated the
	// transaction.
	SiPoc param.Opt[string] `json:"siPOC,omitzero"`
	// The code that represents the International Civil Aviation Organization (ICAO)
	// designations of an airport.
	SourceIcao param.Opt[string] `json:"sourceICAO,omitzero"`
	// The parts associated with this support item.
	LogisticsParts []LogisticsSupportNewParamsLogisticsSupportItemLogisticsPart `json:"logisticsParts,omitzero"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsSupportNewParamsLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticsSupportNewParamsLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsSupportItem) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsSupportItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticsSupportNewParamsLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber param.Opt[string] `json:"figureNumber,omitzero"`
	// Technical order manual index number for the requested part.
	IndexNumber param.Opt[string] `json:"indexNumber,omitzero"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier param.Opt[string] `json:"locationVerifier,omitzero"`
	// Code for a unit of measurement.
	MeasurementUnitCode param.Opt[string] `json:"measurementUnitCode,omitzero"`
	// The National Stock Number of the part being requested or supplied.
	NationalStockNumber param.Opt[string] `json:"nationalStockNumber,omitzero"`
	// Requested or supplied part number.
	PartNumber param.Opt[string] `json:"partNumber,omitzero"`
	// The person who validated the request for parts.
	RequestVerifier param.Opt[string] `json:"requestVerifier,omitzero"`
	// The supply document number.
	SupplyDocumentNumber param.Opt[string] `json:"supplyDocumentNumber,omitzero"`
	// Indicates the specified Technical Order manual holding the aircraft information
	// for use in diagnosing a problem or condition.
	TechnicalOrderText param.Opt[string] `json:"technicalOrderText,omitzero"`
	// Work Unit Code (WUC), or for some aircraft types, the Reference Designator.
	WorkUnitCode param.Opt[string] `json:"workUnitCode,omitzero"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticsSupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsSupportItemLogisticsPart) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsSupportItemLogisticsPart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticsSupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock struct {
	// The quantity of available parts needed from sourceICAO.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The ICAO code for the primary location with available parts.
	SourceIcao param.Opt[string] `json:"sourceICAO,omitzero"`
	// The datetime when the parts were sourced, in ISO 8601 UTC format with
	// millisecond precision.
	StockCheckTime param.Opt[time.Time] `json:"stockCheckTime,omitzero" format:"date-time"`
	// The point of contact at the sourced location.
	StockPoc param.Opt[string] `json:"stockPOC,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportNewParamsLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsSupportItemLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsSupportItemLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsSupportItemLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticsSupportNewParamsLogisticsSupportItemLogisticsSpecialty struct {
	// The first name of the specialist.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The last four digits of the specialist's social security number.
	Last4Ssn param.Opt[string] `json:"last4Ssn,omitzero"`
	// The last name of the specialist.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// Military service rank designation.
	RankCode param.Opt[string] `json:"rankCode,omitzero"`
	// Type code that determines role of the mission response team member. TC - Team
	// Chief, TM - Team Member.
	RoleTypeCode param.Opt[string] `json:"roleTypeCode,omitzero"`
	// Skill level of the mission response team member.
	SkillLevel param.Opt[int64] `json:"skillLevel,omitzero"`
	// Indicates where the repairs will be performed, or which shop specialty has been
	// assigned responsibility for correcting the discrepancy. Shop specialties are
	// normally listed in abbreviated format.
	Specialty param.Opt[string] `json:"specialty,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsSupportItemLogisticsSpecialty) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsSupportItemLogisticsSpecialty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticsSupportNewParamsLogisticsTransportationPlan struct {
	// Actual time of departure of first segment, in ISO 8601 UTC format with
	// millisecond precision.
	ActDepTime param.Opt[time.Time] `json:"actDepTime,omitzero" format:"date-time"`
	// These are the initial maintenance values entered based on the pilot descriptions
	// or the official maintenance evaluation code.
	AircraftStatus param.Opt[string] `json:"aircraftStatus,omitzero"`
	// Approximate time of arrival of final segement, in ISO 8601 UTC format with
	// millisecond precision.
	ApproxArrTime param.Opt[time.Time] `json:"approxArrTime,omitzero" format:"date-time"`
	// GC. LGTP_CANX_DT. GD2: Date when the transportation plan was cancelled, in ISO
	// 8601 UTC format with millisecond precision.
	CancelledDate param.Opt[time.Time] `json:"cancelledDate,omitzero" format:"date-time"`
	// GC. LGTP_CLSD_DT. GD2: Date when the transportation plan was closed, in ISO 8601
	// UTC format with millisecond precision.
	ClosedDate param.Opt[time.Time] `json:"closedDate,omitzero" format:"date-time"`
	// The AMS username of the operator who alters the coordination status.
	// Automatically captured by the system.
	Coordinator param.Opt[string] `json:"coordinator,omitzero"`
	// The AMS user unit_id of the operator who alters the coordination status.
	// Automatically captured by the system from table AMS_USER.
	CoordinatorUnit param.Opt[string] `json:"coordinatorUnit,omitzero"`
	// Destination location ICAO.
	DestinationIcao param.Opt[string] `json:"destinationICAO,omitzero"`
	// Transportation plan duration, expressed in the format MMM:SS.
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ETA of the final segment, in ISO 8601 UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// ETD of the first segment, in ISO 8601 UTC format with millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// Last time transportation plan was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastChangedDate param.Opt[time.Time] `json:"lastChangedDate,omitzero" format:"date-time"`
	// The identifier that represents a Logistics Master Record.
	LogisticMasterRecordID param.Opt[string] `json:"logisticMasterRecordId,omitzero"`
	// The major command for the current unit.
	Majcom param.Opt[string] `json:"majcom,omitzero"`
	// Indicates whether there have been changes to changes to ICAOs, estArrTime, or
	// estDepTime since this Transportation Plan was last edited.
	MissionChange param.Opt[bool] `json:"missionChange,omitzero"`
	// Transportation plan enroute stops.
	NumEnrouteStops param.Opt[int64] `json:"numEnrouteStops,omitzero"`
	// The number of transloads for this Transportation Plan.
	NumTransLoads param.Opt[int64] `json:"numTransLoads,omitzero"`
	// The origin location.
	OriginIcao param.Opt[string] `json:"originICAO,omitzero"`
	// Defines the transporation plan as either a deployment or redeployment.
	PlanDefinition param.Opt[string] `json:"planDefinition,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	PlansNumber param.Opt[string] `json:"plansNumber,omitzero"`
	// GDSS2 uses an 8 character serial number to uniquely identify the aircraft and
	// MDS combination. This is a portion of the full manufacturer serial number.
	SerialNumber param.Opt[string] `json:"serialNumber,omitzero"`
	// Transporation Coordination status code. Cancel, Send to APCC, working, agree,
	// disapprove or blank.
	StatusCode param.Opt[string] `json:"statusCode,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	TpAircraftMds param.Opt[string] `json:"tpAircraftMDS,omitzero"`
	// Contains the tail number displayed by GDSS2.
	TpTailNumber param.Opt[string] `json:"tpTailNumber,omitzero"`
	// The transportation segments associated with this transportation plan.
	LogisticsSegments []LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments,omitzero"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsTransportationPlan) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsTransportationPlan
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsSegment struct {
	// Airport ICAO arrival code.
	ArrivalIcao param.Opt[string] `json:"arrivalICAO,omitzero"`
	// Airport ICAO departure code.
	DepartureIcao param.Opt[string] `json:"departureICAO,omitzero"`
	// The GDSS mission ID for this segment.
	ExtMissionID param.Opt[string] `json:"extMissionId,omitzero"`
	// The unique identifier of the mission to which this logistics record is assigned.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// Start air mission itinerary point identifier.
	Itin param.Opt[int64] `json:"itin,omitzero"`
	// The user generated identifier for an air mission subgroup.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	MissionType param.Opt[string] `json:"missionType,omitzero"`
	// Transportation mode. AMC airlift, Commercial airlift, Other, or surface
	// transportation.
	ModeCode param.Opt[string] `json:"modeCode,omitzero"`
	// Actual arrival time to segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActArrTime param.Opt[time.Time] `json:"segActArrTime,omitzero" format:"date-time"`
	// Actual departure time to the segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActDepTime param.Opt[time.Time] `json:"segActDepTime,omitzero" format:"date-time"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	SegAircraftMds param.Opt[string] `json:"segAircraftMDS,omitzero"`
	// GC. LGTPS_C_DT_EST_ARR. GD2: Estimated arrival time to the segment destination.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETA
	// is derived from the Mission segment destination point. This datetime should be
	// in ISO 8601 UTC format with millisecond precision.
	SegEstArrTime param.Opt[time.Time] `json:"segEstArrTime,omitzero" format:"date-time"`
	// GC. LGTPS_C_DT_EST_DEP. GD2: Estimated departure time from the segment origin.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETD
	// is derived from the Mission segment origin point. This datetime should be in ISO
	// 8601 UTC format with millisecond precision.
	SegEstDepTime param.Opt[time.Time] `json:"segEstDepTime,omitzero" format:"date-time"`
	// Used to sequence the segments in the transportation plan.
	SegmentNumber param.Opt[int64] `json:"segmentNumber,omitzero"`
	// The identifier that represents a specific aircraft within an aircraft type.
	SegTailNumber param.Opt[string] `json:"segTailNumber,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsSegment) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsSegment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogisticsSupportUpdateParams struct {
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
	DataMode LogisticsSupportUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The time this report was created, in ISO 8601 UTC format with millisecond
	// precision.
	RptCreatedTime time.Time `json:"rptCreatedTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// The current ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	CurrIcao param.Opt[string] `json:"currICAO,omitzero"`
	// The estimated time mission capable for the aircraft, in ISO 8601 UCT format with
	// millisecond precision. This is the estimated time when the aircraft is mission
	// ready.
	Etic param.Opt[time.Time] `json:"etic,omitzero" format:"date-time"`
	// Logistics estimated time mission capable.
	Etmc param.Opt[time.Time] `json:"etmc,omitzero" format:"date-time"`
	// Optional system identifier from external systs. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtSystemID param.Opt[string] `json:"extSystemId,omitzero"`
	// This field identifies the pacing event for bringing the aircraft to Mission
	// Capable status. It is used in calculating the Estimated Time Mission Capable
	// (ETMC) value. Acceptable values are WA (Will Advise), INW (In Work), P+hhh.h
	// (where P=parts and hhh.h is the number of hours up to 999 plus tenths of hours),
	// EQ+hhh.h (EQ=equipment), MRT+hhh.h (MRT=maintenance recovery team).
	LogisticAction param.Opt[string] `json:"logisticAction,omitzero"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID param.Opt[string] `json:"logisticsRecordId,omitzero"`
	// The maintenance status code of the aircraft which may be based on pilot
	// descriptions or evaluation codes. Contact the source provider for details.
	MaintStatusCode param.Opt[string] `json:"maintStatusCode,omitzero"`
	// The time indicating when all mission essential problems with a given aircraft
	// have been fixed and is mission capable. This datetime should be in ISO 8601 UTC
	// format with millisecond precision.
	McTime param.Opt[time.Time] `json:"mcTime,omitzero" format:"date-time"`
	// The time indicating when a given aircraft breaks for a mission essential reason.
	// This datetime should be in ISO 8601 UTC format with millisecond precision.
	MeTime param.Opt[time.Time] `json:"meTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The organization that owns this logistics record.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// This is used to indicate whether a closed master record has been reopened.
	ReopenFlag param.Opt[bool] `json:"reopenFlag,omitzero"`
	// The time this report was closed, in ISO 8601 UTC format with millisecond
	// precision.
	RptClosedTime param.Opt[time.Time] `json:"rptClosedTime,omitzero" format:"date-time"`
	// The supplying ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	SuppIcao param.Opt[string] `json:"suppICAO,omitzero"`
	// The tail number of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// Discrepancy information associated with this LogisticsSupport record.
	LogisticsDiscrepancyInfos []LogisticsSupportUpdateParamsLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos,omitzero"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsSupportUpdateParamsLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticsSupportUpdateParamsLogisticsSupportItem `json:"logisticsSupportItems,omitzero"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticsSupportUpdateParamsLogisticsTransportationPlan `json:"logisticsTransportationPlans,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParams) UnmarshalJSON(data []byte) error {
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
type LogisticsSupportUpdateParamsDataMode string

const (
	LogisticsSupportUpdateParamsDataModeReal      LogisticsSupportUpdateParamsDataMode = "REAL"
	LogisticsSupportUpdateParamsDataModeTest      LogisticsSupportUpdateParamsDataMode = "TEST"
	LogisticsSupportUpdateParamsDataModeSimulated LogisticsSupportUpdateParamsDataMode = "SIMULATED"
	LogisticsSupportUpdateParamsDataModeExercise  LogisticsSupportUpdateParamsDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticsSupportUpdateParamsLogisticsDiscrepancyInfo struct {
	// The discrepancy closure time, in ISO 8601 UTC format with millisecond precision.
	ClosureTime param.Opt[time.Time] `json:"closureTime,omitzero" format:"date-time"`
	// The aircraft discrepancy description.
	DiscrepancyInfo param.Opt[string] `json:"discrepancyInfo,omitzero"`
	// Job Control Number of the discrepancy.
	Jcn param.Opt[string] `json:"jcn,omitzero"`
	// The job start time, in ISO 8601 UTC format with millisecond precision.
	JobStTime param.Opt[time.Time] `json:"jobStTime,omitzero" format:"date-time"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsDiscrepancyInfo) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsDiscrepancyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportUpdateParamsLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticsSupportUpdateParamsLogisticsSupportItem struct {
	// This element indicates whether or not the supplied item is contained within
	// another item.
	Cannibalized param.Opt[bool] `json:"cannibalized,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	DeployPlanNumber param.Opt[string] `json:"deployPlanNumber,omitzero"`
	// The technical order name of the part ordered.
	Description param.Opt[string] `json:"description,omitzero"`
	// The last time this supported item was updated, in ISO 8601 UTC format with
	// millisecond precision.
	ItemLastChangedDate param.Opt[time.Time] `json:"itemLastChangedDate,omitzero" format:"date-time"`
	// A number assigned by Job Control to monitor and record maintenance actions
	// required to correct the associated aircraft maintenance discrepancy. It is
	// seven, nine or twelve characters, depending on the base-specific numbering
	// scheme. If seven characters: characters 1-3 are Julian date, 4-7 are sequence
	// numbers. If nine characters: characters 1-2 are last two digits of the year,
	// characters 3-5 are Julian date, 6-9 are sequence numbers. If twelve characters:
	// characters 1-2 are last two digits of the year, 3-5 are Julian date, 6-9 are
	// sequence numbers, and 10-12 are a three-digit supplemental number.
	JobControlNumber param.Opt[string] `json:"jobControlNumber,omitzero"`
	// Military aircraft discrepancy logistics requisition ordered quantity. The
	// quantity of equipment ordered that is required to fix the aircraft.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The time the item is ready, in ISO 8601 UTC format with millisecond precision.
	ReadyTime param.Opt[time.Time] `json:"readyTime,omitzero" format:"date-time"`
	// The time the item is received, in ISO 8601 UTC format with millisecond
	// precision.
	ReceivedTime param.Opt[time.Time] `json:"receivedTime,omitzero" format:"date-time"`
	// The type of recovery request needed. Contact the source provider for details.
	RecoveryRequestTypeCode param.Opt[string] `json:"recoveryRequestTypeCode,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	RedeployPlanNumber param.Opt[string] `json:"redeployPlanNumber,omitzero"`
	// This is the Redeploy (return) Transportation Control Number/Tracking Reference
	// Number for the selected item.
	RedeployShipmentUnitID param.Opt[string] `json:"redeployShipmentUnitId,omitzero"`
	// The request or record number for this item type (Equipent, Part, or MRT).
	RequestNumber param.Opt[string] `json:"requestNumber,omitzero"`
	// This element indicates if the supplied item is characterized as additional
	// support.
	ResupportFlag param.Opt[bool] `json:"resupportFlag,omitzero"`
	// Shipment Unit Identifier is the Transportation Control Number (TCN) for shipping
	// that piece of equipment being requested.
	ShipmentUnitID param.Opt[string] `json:"shipmentUnitId,omitzero"`
	// The point of contact is a free text field to add information about the
	// individual(s) with knowledge of the referenced requested or supplied item(s).
	// The default value for this field is the last name, first name, and middle
	// initial of the operator who created the records and/or generated the
	// transaction.
	SiPoc param.Opt[string] `json:"siPOC,omitzero"`
	// The code that represents the International Civil Aviation Organization (ICAO)
	// designations of an airport.
	SourceIcao param.Opt[string] `json:"sourceICAO,omitzero"`
	// The parts associated with this support item.
	LogisticsParts []LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPart `json:"logisticsParts,omitzero"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsSupportItem) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsSupportItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber param.Opt[string] `json:"figureNumber,omitzero"`
	// Technical order manual index number for the requested part.
	IndexNumber param.Opt[string] `json:"indexNumber,omitzero"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier param.Opt[string] `json:"locationVerifier,omitzero"`
	// Code for a unit of measurement.
	MeasurementUnitCode param.Opt[string] `json:"measurementUnitCode,omitzero"`
	// The National Stock Number of the part being requested or supplied.
	NationalStockNumber param.Opt[string] `json:"nationalStockNumber,omitzero"`
	// Requested or supplied part number.
	PartNumber param.Opt[string] `json:"partNumber,omitzero"`
	// The person who validated the request for parts.
	RequestVerifier param.Opt[string] `json:"requestVerifier,omitzero"`
	// The supply document number.
	SupplyDocumentNumber param.Opt[string] `json:"supplyDocumentNumber,omitzero"`
	// Indicates the specified Technical Order manual holding the aircraft information
	// for use in diagnosing a problem or condition.
	TechnicalOrderText param.Opt[string] `json:"technicalOrderText,omitzero"`
	// Work Unit Code (WUC), or for some aircraft types, the Reference Designator.
	WorkUnitCode param.Opt[string] `json:"workUnitCode,omitzero"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPart) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock struct {
	// The quantity of available parts needed from sourceICAO.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The ICAO code for the primary location with available parts.
	SourceIcao param.Opt[string] `json:"sourceICAO,omitzero"`
	// The datetime when the parts were sourced, in ISO 8601 UTC format with
	// millisecond precision.
	StockCheckTime param.Opt[time.Time] `json:"stockCheckTime,omitzero" format:"date-time"`
	// The point of contact at the sourced location.
	StockPoc param.Opt[string] `json:"stockPOC,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsSpecialty struct {
	// The first name of the specialist.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The last four digits of the specialist's social security number.
	Last4Ssn param.Opt[string] `json:"last4Ssn,omitzero"`
	// The last name of the specialist.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// Military service rank designation.
	RankCode param.Opt[string] `json:"rankCode,omitzero"`
	// Type code that determines role of the mission response team member. TC - Team
	// Chief, TM - Team Member.
	RoleTypeCode param.Opt[string] `json:"roleTypeCode,omitzero"`
	// Skill level of the mission response team member.
	SkillLevel param.Opt[int64] `json:"skillLevel,omitzero"`
	// Indicates where the repairs will be performed, or which shop specialty has been
	// assigned responsibility for correcting the discrepancy. Shop specialties are
	// normally listed in abbreviated format.
	Specialty param.Opt[string] `json:"specialty,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsSpecialty) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsSpecialty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticsSupportUpdateParamsLogisticsTransportationPlan struct {
	// Actual time of departure of first segment, in ISO 8601 UTC format with
	// millisecond precision.
	ActDepTime param.Opt[time.Time] `json:"actDepTime,omitzero" format:"date-time"`
	// These are the initial maintenance values entered based on the pilot descriptions
	// or the official maintenance evaluation code.
	AircraftStatus param.Opt[string] `json:"aircraftStatus,omitzero"`
	// Approximate time of arrival of final segement, in ISO 8601 UTC format with
	// millisecond precision.
	ApproxArrTime param.Opt[time.Time] `json:"approxArrTime,omitzero" format:"date-time"`
	// GC. LGTP_CANX_DT. GD2: Date when the transportation plan was cancelled, in ISO
	// 8601 UTC format with millisecond precision.
	CancelledDate param.Opt[time.Time] `json:"cancelledDate,omitzero" format:"date-time"`
	// GC. LGTP_CLSD_DT. GD2: Date when the transportation plan was closed, in ISO 8601
	// UTC format with millisecond precision.
	ClosedDate param.Opt[time.Time] `json:"closedDate,omitzero" format:"date-time"`
	// The AMS username of the operator who alters the coordination status.
	// Automatically captured by the system.
	Coordinator param.Opt[string] `json:"coordinator,omitzero"`
	// The AMS user unit_id of the operator who alters the coordination status.
	// Automatically captured by the system from table AMS_USER.
	CoordinatorUnit param.Opt[string] `json:"coordinatorUnit,omitzero"`
	// Destination location ICAO.
	DestinationIcao param.Opt[string] `json:"destinationICAO,omitzero"`
	// Transportation plan duration, expressed in the format MMM:SS.
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ETA of the final segment, in ISO 8601 UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// ETD of the first segment, in ISO 8601 UTC format with millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// Last time transportation plan was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastChangedDate param.Opt[time.Time] `json:"lastChangedDate,omitzero" format:"date-time"`
	// The identifier that represents a Logistics Master Record.
	LogisticMasterRecordID param.Opt[string] `json:"logisticMasterRecordId,omitzero"`
	// The major command for the current unit.
	Majcom param.Opt[string] `json:"majcom,omitzero"`
	// Indicates whether there have been changes to changes to ICAOs, estArrTime, or
	// estDepTime since this Transportation Plan was last edited.
	MissionChange param.Opt[bool] `json:"missionChange,omitzero"`
	// Transportation plan enroute stops.
	NumEnrouteStops param.Opt[int64] `json:"numEnrouteStops,omitzero"`
	// The number of transloads for this Transportation Plan.
	NumTransLoads param.Opt[int64] `json:"numTransLoads,omitzero"`
	// The origin location.
	OriginIcao param.Opt[string] `json:"originICAO,omitzero"`
	// Defines the transporation plan as either a deployment or redeployment.
	PlanDefinition param.Opt[string] `json:"planDefinition,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	PlansNumber param.Opt[string] `json:"plansNumber,omitzero"`
	// GDSS2 uses an 8 character serial number to uniquely identify the aircraft and
	// MDS combination. This is a portion of the full manufacturer serial number.
	SerialNumber param.Opt[string] `json:"serialNumber,omitzero"`
	// Transporation Coordination status code. Cancel, Send to APCC, working, agree,
	// disapprove or blank.
	StatusCode param.Opt[string] `json:"statusCode,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	TpAircraftMds param.Opt[string] `json:"tpAircraftMDS,omitzero"`
	// Contains the tail number displayed by GDSS2.
	TpTailNumber param.Opt[string] `json:"tpTailNumber,omitzero"`
	// The transportation segments associated with this transportation plan.
	LogisticsSegments []LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments,omitzero"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsTransportationPlan) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsTransportationPlan
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsSegment struct {
	// Airport ICAO arrival code.
	ArrivalIcao param.Opt[string] `json:"arrivalICAO,omitzero"`
	// Airport ICAO departure code.
	DepartureIcao param.Opt[string] `json:"departureICAO,omitzero"`
	// The GDSS mission ID for this segment.
	ExtMissionID param.Opt[string] `json:"extMissionId,omitzero"`
	// The unique identifier of the mission to which this logistics record is assigned.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// Start air mission itinerary point identifier.
	Itin param.Opt[int64] `json:"itin,omitzero"`
	// The user generated identifier for an air mission subgroup.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	MissionType param.Opt[string] `json:"missionType,omitzero"`
	// Transportation mode. AMC airlift, Commercial airlift, Other, or surface
	// transportation.
	ModeCode param.Opt[string] `json:"modeCode,omitzero"`
	// Actual arrival time to segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActArrTime param.Opt[time.Time] `json:"segActArrTime,omitzero" format:"date-time"`
	// Actual departure time to the segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActDepTime param.Opt[time.Time] `json:"segActDepTime,omitzero" format:"date-time"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	SegAircraftMds param.Opt[string] `json:"segAircraftMDS,omitzero"`
	// GC. LGTPS_C_DT_EST_ARR. GD2: Estimated arrival time to the segment destination.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETA
	// is derived from the Mission segment destination point. This datetime should be
	// in ISO 8601 UTC format with millisecond precision.
	SegEstArrTime param.Opt[time.Time] `json:"segEstArrTime,omitzero" format:"date-time"`
	// GC. LGTPS_C_DT_EST_DEP. GD2: Estimated departure time from the segment origin.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETD
	// is derived from the Mission segment origin point. This datetime should be in ISO
	// 8601 UTC format with millisecond precision.
	SegEstDepTime param.Opt[time.Time] `json:"segEstDepTime,omitzero" format:"date-time"`
	// Used to sequence the segments in the transportation plan.
	SegmentNumber param.Opt[int64] `json:"segmentNumber,omitzero"`
	// The identifier that represents a specific aircraft within an aircraft type.
	SegTailNumber param.Opt[string] `json:"segTailNumber,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsSegment) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsSegment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogisticsSupportListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogisticsSupportListParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticsSupportCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogisticsSupportCountParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticsSupportNewBulkParams struct {
	Body []LogisticsSupportNewBulkParamsBody
	paramObj
}

func (r LogisticsSupportNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *LogisticsSupportNewBulkParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
//
// The properties ClassificationMarking, DataMode, RptCreatedTime, Source are
// required.
type LogisticsSupportNewBulkParamsBody struct {
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
	// The time this report was created, in ISO 8601 UTC format with millisecond
	// precision.
	RptCreatedTime time.Time `json:"rptCreatedTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Time the row was created in the database.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The current ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	CurrIcao param.Opt[string] `json:"currICAO,omitzero"`
	// The estimated time mission capable for the aircraft, in ISO 8601 UCT format with
	// millisecond precision. This is the estimated time when the aircraft is mission
	// ready.
	Etic param.Opt[time.Time] `json:"etic,omitzero" format:"date-time"`
	// Logistics estimated time mission capable.
	Etmc param.Opt[time.Time] `json:"etmc,omitzero" format:"date-time"`
	// Optional system identifier from external systs. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtSystemID param.Opt[string] `json:"extSystemId,omitzero"`
	// This field identifies the pacing event for bringing the aircraft to Mission
	// Capable status. It is used in calculating the Estimated Time Mission Capable
	// (ETMC) value. Acceptable values are WA (Will Advise), INW (In Work), P+hhh.h
	// (where P=parts and hhh.h is the number of hours up to 999 plus tenths of hours),
	// EQ+hhh.h (EQ=equipment), MRT+hhh.h (MRT=maintenance recovery team).
	LogisticAction param.Opt[string] `json:"logisticAction,omitzero"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID param.Opt[string] `json:"logisticsRecordId,omitzero"`
	// The maintenance status code of the aircraft which may be based on pilot
	// descriptions or evaluation codes. Contact the source provider for details.
	MaintStatusCode param.Opt[string] `json:"maintStatusCode,omitzero"`
	// The time indicating when all mission essential problems with a given aircraft
	// have been fixed and is mission capable. This datetime should be in ISO 8601 UTC
	// format with millisecond precision.
	McTime param.Opt[time.Time] `json:"mcTime,omitzero" format:"date-time"`
	// The time indicating when a given aircraft breaks for a mission essential reason.
	// This datetime should be in ISO 8601 UTC format with millisecond precision.
	MeTime param.Opt[time.Time] `json:"meTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The organization that owns this logistics record.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// This is used to indicate whether a closed master record has been reopened.
	ReopenFlag param.Opt[bool] `json:"reopenFlag,omitzero"`
	// The time this report was closed, in ISO 8601 UTC format with millisecond
	// precision.
	RptClosedTime param.Opt[time.Time] `json:"rptClosedTime,omitzero" format:"date-time"`
	// The supplying ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	SuppIcao param.Opt[string] `json:"suppICAO,omitzero"`
	// The tail number of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Discrepancy information associated with this LogisticsSupport record.
	LogisticsDiscrepancyInfos []LogisticsSupportNewBulkParamsBodyLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos,omitzero"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsSupportNewBulkParamsBodyLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticsSupportNewBulkParamsBodyLogisticsSupportItem `json:"logisticsSupportItems,omitzero"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlan `json:"logisticsTransportationPlans,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogisticsSupportNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Discrepancy information associated with this LogisticsSupport record.
type LogisticsSupportNewBulkParamsBodyLogisticsDiscrepancyInfo struct {
	// The discrepancy closure time, in ISO 8601 UTC format with millisecond precision.
	ClosureTime param.Opt[time.Time] `json:"closureTime,omitzero" format:"date-time"`
	// The aircraft discrepancy description.
	DiscrepancyInfo param.Opt[string] `json:"discrepancyInfo,omitzero"`
	// Job Control Number of the discrepancy.
	Jcn param.Opt[string] `json:"jcn,omitzero"`
	// The job start time, in ISO 8601 UTC format with millisecond precision.
	JobStTime param.Opt[time.Time] `json:"jobStTime,omitzero" format:"date-time"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsDiscrepancyInfo) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsDiscrepancyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportNewBulkParamsBodyLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticsSupportNewBulkParamsBodyLogisticsSupportItem struct {
	// This element indicates whether or not the supplied item is contained within
	// another item.
	Cannibalized param.Opt[bool] `json:"cannibalized,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	DeployPlanNumber param.Opt[string] `json:"deployPlanNumber,omitzero"`
	// The technical order name of the part ordered.
	Description param.Opt[string] `json:"description,omitzero"`
	// The last time this supported item was updated, in ISO 8601 UTC format with
	// millisecond precision.
	ItemLastChangedDate param.Opt[time.Time] `json:"itemLastChangedDate,omitzero" format:"date-time"`
	// A number assigned by Job Control to monitor and record maintenance actions
	// required to correct the associated aircraft maintenance discrepancy. It is
	// seven, nine or twelve characters, depending on the base-specific numbering
	// scheme. If seven characters: characters 1-3 are Julian date, 4-7 are sequence
	// numbers. If nine characters: characters 1-2 are last two digits of the year,
	// characters 3-5 are Julian date, 6-9 are sequence numbers. If twelve characters:
	// characters 1-2 are last two digits of the year, 3-5 are Julian date, 6-9 are
	// sequence numbers, and 10-12 are a three-digit supplemental number.
	JobControlNumber param.Opt[string] `json:"jobControlNumber,omitzero"`
	// Military aircraft discrepancy logistics requisition ordered quantity. The
	// quantity of equipment ordered that is required to fix the aircraft.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The time the item is ready, in ISO 8601 UTC format with millisecond precision.
	ReadyTime param.Opt[time.Time] `json:"readyTime,omitzero" format:"date-time"`
	// The time the item is received, in ISO 8601 UTC format with millisecond
	// precision.
	ReceivedTime param.Opt[time.Time] `json:"receivedTime,omitzero" format:"date-time"`
	// The type of recovery request needed. Contact the source provider for details.
	RecoveryRequestTypeCode param.Opt[string] `json:"recoveryRequestTypeCode,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	RedeployPlanNumber param.Opt[string] `json:"redeployPlanNumber,omitzero"`
	// This is the Redeploy (return) Transportation Control Number/Tracking Reference
	// Number for the selected item.
	RedeployShipmentUnitID param.Opt[string] `json:"redeployShipmentUnitId,omitzero"`
	// The request or record number for this item type (Equipent, Part, or MRT).
	RequestNumber param.Opt[string] `json:"requestNumber,omitzero"`
	// This element indicates if the supplied item is characterized as additional
	// support.
	ResupportFlag param.Opt[bool] `json:"resupportFlag,omitzero"`
	// Shipment Unit Identifier is the Transportation Control Number (TCN) for shipping
	// that piece of equipment being requested.
	ShipmentUnitID param.Opt[string] `json:"shipmentUnitId,omitzero"`
	// The point of contact is a free text field to add information about the
	// individual(s) with knowledge of the referenced requested or supplied item(s).
	// The default value for this field is the last name, first name, and middle
	// initial of the operator who created the records and/or generated the
	// transaction.
	SiPoc param.Opt[string] `json:"siPOC,omitzero"`
	// The code that represents the International Civil Aviation Organization (ICAO)
	// designations of an airport.
	SourceIcao param.Opt[string] `json:"sourceICAO,omitzero"`
	// The parts associated with this support item.
	LogisticsParts []LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart `json:"logisticsParts,omitzero"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsSupportItem) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsSupportItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber param.Opt[string] `json:"figureNumber,omitzero"`
	// Technical order manual index number for the requested part.
	IndexNumber param.Opt[string] `json:"indexNumber,omitzero"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier param.Opt[string] `json:"locationVerifier,omitzero"`
	// Code for a unit of measurement.
	MeasurementUnitCode param.Opt[string] `json:"measurementUnitCode,omitzero"`
	// The National Stock Number of the part being requested or supplied.
	NationalStockNumber param.Opt[string] `json:"nationalStockNumber,omitzero"`
	// Requested or supplied part number.
	PartNumber param.Opt[string] `json:"partNumber,omitzero"`
	// The person who validated the request for parts.
	RequestVerifier param.Opt[string] `json:"requestVerifier,omitzero"`
	// The supply document number.
	SupplyDocumentNumber param.Opt[string] `json:"supplyDocumentNumber,omitzero"`
	// Indicates the specified Technical Order manual holding the aircraft information
	// for use in diagnosing a problem or condition.
	TechnicalOrderText param.Opt[string] `json:"technicalOrderText,omitzero"`
	// Work Unit Code (WUC), or for some aircraft types, the Reference Designator.
	WorkUnitCode param.Opt[string] `json:"workUnitCode,omitzero"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock struct {
	// The quantity of available parts needed from sourceICAO.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The ICAO code for the primary location with available parts.
	SourceIcao param.Opt[string] `json:"sourceICAO,omitzero"`
	// The datetime when the parts were sourced, in ISO 8601 UTC format with
	// millisecond precision.
	StockCheckTime param.Opt[time.Time] `json:"stockCheckTime,omitzero" format:"date-time"`
	// The point of contact at the sourced location.
	StockPoc param.Opt[string] `json:"stockPOC,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty struct {
	// The first name of the specialist.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The last four digits of the specialist's social security number.
	Last4Ssn param.Opt[string] `json:"last4Ssn,omitzero"`
	// The last name of the specialist.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// Military service rank designation.
	RankCode param.Opt[string] `json:"rankCode,omitzero"`
	// Type code that determines role of the mission response team member. TC - Team
	// Chief, TM - Team Member.
	RoleTypeCode param.Opt[string] `json:"roleTypeCode,omitzero"`
	// Skill level of the mission response team member.
	SkillLevel param.Opt[int64] `json:"skillLevel,omitzero"`
	// Indicates where the repairs will be performed, or which shop specialty has been
	// assigned responsibility for correcting the discrepancy. Shop specialties are
	// normally listed in abbreviated format.
	Specialty param.Opt[string] `json:"specialty,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlan struct {
	// Actual time of departure of first segment, in ISO 8601 UTC format with
	// millisecond precision.
	ActDepTime param.Opt[time.Time] `json:"actDepTime,omitzero" format:"date-time"`
	// These are the initial maintenance values entered based on the pilot descriptions
	// or the official maintenance evaluation code.
	AircraftStatus param.Opt[string] `json:"aircraftStatus,omitzero"`
	// Approximate time of arrival of final segement, in ISO 8601 UTC format with
	// millisecond precision.
	ApproxArrTime param.Opt[time.Time] `json:"approxArrTime,omitzero" format:"date-time"`
	// GC. LGTP_CANX_DT. GD2: Date when the transportation plan was cancelled, in ISO
	// 8601 UTC format with millisecond precision.
	CancelledDate param.Opt[time.Time] `json:"cancelledDate,omitzero" format:"date-time"`
	// GC. LGTP_CLSD_DT. GD2: Date when the transportation plan was closed, in ISO 8601
	// UTC format with millisecond precision.
	ClosedDate param.Opt[time.Time] `json:"closedDate,omitzero" format:"date-time"`
	// The AMS username of the operator who alters the coordination status.
	// Automatically captured by the system.
	Coordinator param.Opt[string] `json:"coordinator,omitzero"`
	// The AMS user unit_id of the operator who alters the coordination status.
	// Automatically captured by the system from table AMS_USER.
	CoordinatorUnit param.Opt[string] `json:"coordinatorUnit,omitzero"`
	// Destination location ICAO.
	DestinationIcao param.Opt[string] `json:"destinationICAO,omitzero"`
	// Transportation plan duration, expressed in the format MMM:SS.
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ETA of the final segment, in ISO 8601 UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// ETD of the first segment, in ISO 8601 UTC format with millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// Last time transportation plan was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastChangedDate param.Opt[time.Time] `json:"lastChangedDate,omitzero" format:"date-time"`
	// The identifier that represents a Logistics Master Record.
	LogisticMasterRecordID param.Opt[string] `json:"logisticMasterRecordId,omitzero"`
	// The major command for the current unit.
	Majcom param.Opt[string] `json:"majcom,omitzero"`
	// Indicates whether there have been changes to changes to ICAOs, estArrTime, or
	// estDepTime since this Transportation Plan was last edited.
	MissionChange param.Opt[bool] `json:"missionChange,omitzero"`
	// Transportation plan enroute stops.
	NumEnrouteStops param.Opt[int64] `json:"numEnrouteStops,omitzero"`
	// The number of transloads for this Transportation Plan.
	NumTransLoads param.Opt[int64] `json:"numTransLoads,omitzero"`
	// The origin location.
	OriginIcao param.Opt[string] `json:"originICAO,omitzero"`
	// Defines the transporation plan as either a deployment or redeployment.
	PlanDefinition param.Opt[string] `json:"planDefinition,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	PlansNumber param.Opt[string] `json:"plansNumber,omitzero"`
	// GDSS2 uses an 8 character serial number to uniquely identify the aircraft and
	// MDS combination. This is a portion of the full manufacturer serial number.
	SerialNumber param.Opt[string] `json:"serialNumber,omitzero"`
	// Transporation Coordination status code. Cancel, Send to APCC, working, agree,
	// disapprove or blank.
	StatusCode param.Opt[string] `json:"statusCode,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	TpAircraftMds param.Opt[string] `json:"tpAircraftMDS,omitzero"`
	// Contains the tail number displayed by GDSS2.
	TpTailNumber param.Opt[string] `json:"tpTailNumber,omitzero"`
	// The transportation segments associated with this transportation plan.
	LogisticsSegments []LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments,omitzero"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlan) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlan
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment struct {
	// Airport ICAO arrival code.
	ArrivalIcao param.Opt[string] `json:"arrivalICAO,omitzero"`
	// Airport ICAO departure code.
	DepartureIcao param.Opt[string] `json:"departureICAO,omitzero"`
	// The GDSS mission ID for this segment.
	ExtMissionID param.Opt[string] `json:"extMissionId,omitzero"`
	// The unique identifier of the mission to which this logistics record is assigned.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// Start air mission itinerary point identifier.
	Itin param.Opt[int64] `json:"itin,omitzero"`
	// The user generated identifier for an air mission subgroup.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	MissionType param.Opt[string] `json:"missionType,omitzero"`
	// Transportation mode. AMC airlift, Commercial airlift, Other, or surface
	// transportation.
	ModeCode param.Opt[string] `json:"modeCode,omitzero"`
	// Actual arrival time to segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActArrTime param.Opt[time.Time] `json:"segActArrTime,omitzero" format:"date-time"`
	// Actual departure time to the segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActDepTime param.Opt[time.Time] `json:"segActDepTime,omitzero" format:"date-time"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	SegAircraftMds param.Opt[string] `json:"segAircraftMDS,omitzero"`
	// GC. LGTPS_C_DT_EST_ARR. GD2: Estimated arrival time to the segment destination.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETA
	// is derived from the Mission segment destination point. This datetime should be
	// in ISO 8601 UTC format with millisecond precision.
	SegEstArrTime param.Opt[time.Time] `json:"segEstArrTime,omitzero" format:"date-time"`
	// GC. LGTPS_C_DT_EST_DEP. GD2: Estimated departure time from the segment origin.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETD
	// is derived from the Mission segment origin point. This datetime should be in ISO
	// 8601 UTC format with millisecond precision.
	SegEstDepTime param.Opt[time.Time] `json:"segEstDepTime,omitzero" format:"date-time"`
	// Used to sequence the segments in the transportation plan.
	SegmentNumber param.Opt[int64] `json:"segmentNumber,omitzero"`
	// The identifier that represents a specific aircraft within an aircraft type.
	SegTailNumber param.Opt[string] `json:"segTailNumber,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogisticsSupportGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogisticsSupportGetParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticsSupportTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogisticsSupportTupleParams]'s query parameters as
// `url.Values`.
func (r LogisticsSupportTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticsSupportUnvalidatedPublishParams struct {
	Body []LogisticsSupportUnvalidatedPublishParamsBody
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *LogisticsSupportUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
//
// The properties ClassificationMarking, DataMode, RptCreatedTime, Source are
// required.
type LogisticsSupportUnvalidatedPublishParamsBody struct {
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
	// The time this report was created, in ISO 8601 UTC format with millisecond
	// precision.
	RptCreatedTime time.Time `json:"rptCreatedTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Time the row was created in the database.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The current ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	CurrIcao param.Opt[string] `json:"currICAO,omitzero"`
	// The estimated time mission capable for the aircraft, in ISO 8601 UCT format with
	// millisecond precision. This is the estimated time when the aircraft is mission
	// ready.
	Etic param.Opt[time.Time] `json:"etic,omitzero" format:"date-time"`
	// Logistics estimated time mission capable.
	Etmc param.Opt[time.Time] `json:"etmc,omitzero" format:"date-time"`
	// Optional system identifier from external systs. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtSystemID param.Opt[string] `json:"extSystemId,omitzero"`
	// This field identifies the pacing event for bringing the aircraft to Mission
	// Capable status. It is used in calculating the Estimated Time Mission Capable
	// (ETMC) value. Acceptable values are WA (Will Advise), INW (In Work), P+hhh.h
	// (where P=parts and hhh.h is the number of hours up to 999 plus tenths of hours),
	// EQ+hhh.h (EQ=equipment), MRT+hhh.h (MRT=maintenance recovery team).
	LogisticAction param.Opt[string] `json:"logisticAction,omitzero"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID param.Opt[string] `json:"logisticsRecordId,omitzero"`
	// The maintenance status code of the aircraft which may be based on pilot
	// descriptions or evaluation codes. Contact the source provider for details.
	MaintStatusCode param.Opt[string] `json:"maintStatusCode,omitzero"`
	// The time indicating when all mission essential problems with a given aircraft
	// have been fixed and is mission capable. This datetime should be in ISO 8601 UTC
	// format with millisecond precision.
	McTime param.Opt[time.Time] `json:"mcTime,omitzero" format:"date-time"`
	// The time indicating when a given aircraft breaks for a mission essential reason.
	// This datetime should be in ISO 8601 UTC format with millisecond precision.
	MeTime param.Opt[time.Time] `json:"meTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The organization that owns this logistics record.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// This is used to indicate whether a closed master record has been reopened.
	ReopenFlag param.Opt[bool] `json:"reopenFlag,omitzero"`
	// The time this report was closed, in ISO 8601 UTC format with millisecond
	// precision.
	RptClosedTime param.Opt[time.Time] `json:"rptClosedTime,omitzero" format:"date-time"`
	// The supplying ICAO of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	SuppIcao param.Opt[string] `json:"suppICAO,omitzero"`
	// The tail number of the aircraft that is the subject of this
	// LogisticsSupportDetails record.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Discrepancy information associated with this LogisticsSupport record.
	LogisticsDiscrepancyInfos []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos,omitzero"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItem `json:"logisticsSupportItems,omitzero"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan `json:"logisticsTransportationPlans,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogisticsSupportUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Discrepancy information associated with this LogisticsSupport record.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo struct {
	// The discrepancy closure time, in ISO 8601 UTC format with millisecond precision.
	ClosureTime param.Opt[time.Time] `json:"closureTime,omitzero" format:"date-time"`
	// The aircraft discrepancy description.
	DiscrepancyInfo param.Opt[string] `json:"discrepancyInfo,omitzero"`
	// Job Control Number of the discrepancy.
	Jcn param.Opt[string] `json:"jcn,omitzero"`
	// The job start time, in ISO 8601 UTC format with millisecond precision.
	JobStTime param.Opt[time.Time] `json:"jobStTime,omitzero" format:"date-time"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItem struct {
	// This element indicates whether or not the supplied item is contained within
	// another item.
	Cannibalized param.Opt[bool] `json:"cannibalized,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	DeployPlanNumber param.Opt[string] `json:"deployPlanNumber,omitzero"`
	// The technical order name of the part ordered.
	Description param.Opt[string] `json:"description,omitzero"`
	// The last time this supported item was updated, in ISO 8601 UTC format with
	// millisecond precision.
	ItemLastChangedDate param.Opt[time.Time] `json:"itemLastChangedDate,omitzero" format:"date-time"`
	// A number assigned by Job Control to monitor and record maintenance actions
	// required to correct the associated aircraft maintenance discrepancy. It is
	// seven, nine or twelve characters, depending on the base-specific numbering
	// scheme. If seven characters: characters 1-3 are Julian date, 4-7 are sequence
	// numbers. If nine characters: characters 1-2 are last two digits of the year,
	// characters 3-5 are Julian date, 6-9 are sequence numbers. If twelve characters:
	// characters 1-2 are last two digits of the year, 3-5 are Julian date, 6-9 are
	// sequence numbers, and 10-12 are a three-digit supplemental number.
	JobControlNumber param.Opt[string] `json:"jobControlNumber,omitzero"`
	// Military aircraft discrepancy logistics requisition ordered quantity. The
	// quantity of equipment ordered that is required to fix the aircraft.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The time the item is ready, in ISO 8601 UTC format with millisecond precision.
	ReadyTime param.Opt[time.Time] `json:"readyTime,omitzero" format:"date-time"`
	// The time the item is received, in ISO 8601 UTC format with millisecond
	// precision.
	ReceivedTime param.Opt[time.Time] `json:"receivedTime,omitzero" format:"date-time"`
	// The type of recovery request needed. Contact the source provider for details.
	RecoveryRequestTypeCode param.Opt[string] `json:"recoveryRequestTypeCode,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	RedeployPlanNumber param.Opt[string] `json:"redeployPlanNumber,omitzero"`
	// This is the Redeploy (return) Transportation Control Number/Tracking Reference
	// Number for the selected item.
	RedeployShipmentUnitID param.Opt[string] `json:"redeployShipmentUnitId,omitzero"`
	// The request or record number for this item type (Equipent, Part, or MRT).
	RequestNumber param.Opt[string] `json:"requestNumber,omitzero"`
	// This element indicates if the supplied item is characterized as additional
	// support.
	ResupportFlag param.Opt[bool] `json:"resupportFlag,omitzero"`
	// Shipment Unit Identifier is the Transportation Control Number (TCN) for shipping
	// that piece of equipment being requested.
	ShipmentUnitID param.Opt[string] `json:"shipmentUnitId,omitzero"`
	// The point of contact is a free text field to add information about the
	// individual(s) with knowledge of the referenced requested or supplied item(s).
	// The default value for this field is the last name, first name, and middle
	// initial of the operator who created the records and/or generated the
	// transaction.
	SiPoc param.Opt[string] `json:"siPOC,omitzero"`
	// The code that represents the International Civil Aviation Organization (ICAO)
	// designations of an airport.
	SourceIcao param.Opt[string] `json:"sourceICAO,omitzero"`
	// The parts associated with this support item.
	LogisticsParts []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart `json:"logisticsParts,omitzero"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItem) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber param.Opt[string] `json:"figureNumber,omitzero"`
	// Technical order manual index number for the requested part.
	IndexNumber param.Opt[string] `json:"indexNumber,omitzero"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier param.Opt[string] `json:"locationVerifier,omitzero"`
	// Code for a unit of measurement.
	MeasurementUnitCode param.Opt[string] `json:"measurementUnitCode,omitzero"`
	// The National Stock Number of the part being requested or supplied.
	NationalStockNumber param.Opt[string] `json:"nationalStockNumber,omitzero"`
	// Requested or supplied part number.
	PartNumber param.Opt[string] `json:"partNumber,omitzero"`
	// The person who validated the request for parts.
	RequestVerifier param.Opt[string] `json:"requestVerifier,omitzero"`
	// The supply document number.
	SupplyDocumentNumber param.Opt[string] `json:"supplyDocumentNumber,omitzero"`
	// Indicates the specified Technical Order manual holding the aircraft information
	// for use in diagnosing a problem or condition.
	TechnicalOrderText param.Opt[string] `json:"technicalOrderText,omitzero"`
	// Work Unit Code (WUC), or for some aircraft types, the Reference Designator.
	WorkUnitCode param.Opt[string] `json:"workUnitCode,omitzero"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock struct {
	// The quantity of available parts needed from sourceICAO.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The ICAO code for the primary location with available parts.
	SourceIcao param.Opt[string] `json:"sourceICAO,omitzero"`
	// The datetime when the parts were sourced, in ISO 8601 UTC format with
	// millisecond precision.
	StockCheckTime param.Opt[time.Time] `json:"stockCheckTime,omitzero" format:"date-time"`
	// The point of contact at the sourced location.
	StockPoc param.Opt[string] `json:"stockPOC,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty struct {
	// The first name of the specialist.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The last four digits of the specialist's social security number.
	Last4Ssn param.Opt[string] `json:"last4Ssn,omitzero"`
	// The last name of the specialist.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// Military service rank designation.
	RankCode param.Opt[string] `json:"rankCode,omitzero"`
	// Type code that determines role of the mission response team member. TC - Team
	// Chief, TM - Team Member.
	RoleTypeCode param.Opt[string] `json:"roleTypeCode,omitzero"`
	// Skill level of the mission response team member.
	SkillLevel param.Opt[int64] `json:"skillLevel,omitzero"`
	// Indicates where the repairs will be performed, or which shop specialty has been
	// assigned responsibility for correcting the discrepancy. Shop specialties are
	// normally listed in abbreviated format.
	Specialty param.Opt[string] `json:"specialty,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan struct {
	// Actual time of departure of first segment, in ISO 8601 UTC format with
	// millisecond precision.
	ActDepTime param.Opt[time.Time] `json:"actDepTime,omitzero" format:"date-time"`
	// These are the initial maintenance values entered based on the pilot descriptions
	// or the official maintenance evaluation code.
	AircraftStatus param.Opt[string] `json:"aircraftStatus,omitzero"`
	// Approximate time of arrival of final segement, in ISO 8601 UTC format with
	// millisecond precision.
	ApproxArrTime param.Opt[time.Time] `json:"approxArrTime,omitzero" format:"date-time"`
	// GC. LGTP_CANX_DT. GD2: Date when the transportation plan was cancelled, in ISO
	// 8601 UTC format with millisecond precision.
	CancelledDate param.Opt[time.Time] `json:"cancelledDate,omitzero" format:"date-time"`
	// GC. LGTP_CLSD_DT. GD2: Date when the transportation plan was closed, in ISO 8601
	// UTC format with millisecond precision.
	ClosedDate param.Opt[time.Time] `json:"closedDate,omitzero" format:"date-time"`
	// The AMS username of the operator who alters the coordination status.
	// Automatically captured by the system.
	Coordinator param.Opt[string] `json:"coordinator,omitzero"`
	// The AMS user unit_id of the operator who alters the coordination status.
	// Automatically captured by the system from table AMS_USER.
	CoordinatorUnit param.Opt[string] `json:"coordinatorUnit,omitzero"`
	// Destination location ICAO.
	DestinationIcao param.Opt[string] `json:"destinationICAO,omitzero"`
	// Transportation plan duration, expressed in the format MMM:SS.
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ETA of the final segment, in ISO 8601 UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// ETD of the first segment, in ISO 8601 UTC format with millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// Last time transportation plan was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastChangedDate param.Opt[time.Time] `json:"lastChangedDate,omitzero" format:"date-time"`
	// The identifier that represents a Logistics Master Record.
	LogisticMasterRecordID param.Opt[string] `json:"logisticMasterRecordId,omitzero"`
	// The major command for the current unit.
	Majcom param.Opt[string] `json:"majcom,omitzero"`
	// Indicates whether there have been changes to changes to ICAOs, estArrTime, or
	// estDepTime since this Transportation Plan was last edited.
	MissionChange param.Opt[bool] `json:"missionChange,omitzero"`
	// Transportation plan enroute stops.
	NumEnrouteStops param.Opt[int64] `json:"numEnrouteStops,omitzero"`
	// The number of transloads for this Transportation Plan.
	NumTransLoads param.Opt[int64] `json:"numTransLoads,omitzero"`
	// The origin location.
	OriginIcao param.Opt[string] `json:"originICAO,omitzero"`
	// Defines the transporation plan as either a deployment or redeployment.
	PlanDefinition param.Opt[string] `json:"planDefinition,omitzero"`
	// System generated reference id for the transportation plan. Format: TXXXXXNNNN
	// T - Transportation, Sequence Number, Node Id.
	PlansNumber param.Opt[string] `json:"plansNumber,omitzero"`
	// GDSS2 uses an 8 character serial number to uniquely identify the aircraft and
	// MDS combination. This is a portion of the full manufacturer serial number.
	SerialNumber param.Opt[string] `json:"serialNumber,omitzero"`
	// Transporation Coordination status code. Cancel, Send to APCC, working, agree,
	// disapprove or blank.
	StatusCode param.Opt[string] `json:"statusCode,omitzero"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	TpAircraftMds param.Opt[string] `json:"tpAircraftMDS,omitzero"`
	// Contains the tail number displayed by GDSS2.
	TpTailNumber param.Opt[string] `json:"tpTailNumber,omitzero"`
	// The transportation segments associated with this transportation plan.
	LogisticsSegments []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments,omitzero"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment struct {
	// Airport ICAO arrival code.
	ArrivalIcao param.Opt[string] `json:"arrivalICAO,omitzero"`
	// Airport ICAO departure code.
	DepartureIcao param.Opt[string] `json:"departureICAO,omitzero"`
	// The GDSS mission ID for this segment.
	ExtMissionID param.Opt[string] `json:"extMissionId,omitzero"`
	// The unique identifier of the mission to which this logistics record is assigned.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// Start air mission itinerary point identifier.
	Itin param.Opt[int64] `json:"itin,omitzero"`
	// The user generated identifier for an air mission subgroup.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	MissionType param.Opt[string] `json:"missionType,omitzero"`
	// Transportation mode. AMC airlift, Commercial airlift, Other, or surface
	// transportation.
	ModeCode param.Opt[string] `json:"modeCode,omitzero"`
	// Actual arrival time to segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActArrTime param.Opt[time.Time] `json:"segActArrTime,omitzero" format:"date-time"`
	// Actual departure time to the segment destination, in ISO 8601 UTC format with
	// millisecond precision.
	SegActDepTime param.Opt[time.Time] `json:"segActDepTime,omitzero" format:"date-time"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	SegAircraftMds param.Opt[string] `json:"segAircraftMDS,omitzero"`
	// GC. LGTPS_C_DT_EST_ARR. GD2: Estimated arrival time to the segment destination.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETA
	// is derived from the Mission segment destination point. This datetime should be
	// in ISO 8601 UTC format with millisecond precision.
	SegEstArrTime param.Opt[time.Time] `json:"segEstArrTime,omitzero" format:"date-time"`
	// GC. LGTPS_C_DT_EST_DEP. GD2: Estimated departure time from the segment origin.
	// Only supplied when the segment is not attached to a Mission, otherwise the ETD
	// is derived from the Mission segment origin point. This datetime should be in ISO
	// 8601 UTC format with millisecond precision.
	SegEstDepTime param.Opt[time.Time] `json:"segEstDepTime,omitzero" format:"date-time"`
	// Used to sequence the segments in the transportation plan.
	SegmentNumber param.Opt[int64] `json:"segmentNumber,omitzero"`
	// The identifier that represents a specific aircraft within an aircraft type.
	SegTailNumber param.Opt[string] `json:"segTailNumber,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
