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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// LogisticssupportService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogisticssupportService] method instead.
type LogisticssupportService struct {
	Options []option.RequestOption
	History LogisticssupportHistoryService
}

// NewLogisticssupportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogisticssupportService(opts ...option.RequestOption) (r LogisticssupportService) {
	r = LogisticssupportService{}
	r.Options = opts
	r.History = NewLogisticssupportHistoryService(opts...)
	return
}

// Service operation to take a single LogisticsSupport record as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *LogisticssupportService) New(ctx context.Context, body LogisticssupportNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/logisticssupport"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single LogisticsSupport record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *LogisticssupportService) Update(ctx context.Context, id string, body LogisticssupportUpdateParams, opts ...option.RequestOption) (err error) {
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
func (r *LogisticssupportService) List(ctx context.Context, query LogisticssupportListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LogisticssupportListResponse], err error) {
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
func (r *LogisticssupportService) ListAutoPaging(ctx context.Context, query LogisticssupportListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LogisticssupportListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LogisticssupportService) Count(ctx context.Context, query LogisticssupportCountParams, opts ...option.RequestOption) (res *string, err error) {
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
func (r *LogisticssupportService) NewBulk(ctx context.Context, body LogisticssupportNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/logisticssupport/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single LogisticsSupport record by its unique ID
// passed as a path parameter.
func (r *LogisticssupportService) Get(ctx context.Context, id string, query LogisticssupportGetParams, opts ...option.RequestOption) (res *LogisticssupportGetResponse, err error) {
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
func (r *LogisticssupportService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/logisticssupport/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
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
func (r *LogisticssupportService) Tuple(ctx context.Context, query LogisticssupportTupleParams, opts ...option.RequestOption) (res *[]LogisticssupportTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/logisticssupport/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple logisticssupport records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LogisticssupportService) UnvalidatedPublish(ctx context.Context, body LogisticssupportUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		LastChanged resp.Field
		Remark      resp.Field
		Username    resp.Field
		ExtraFields map[string]resp.Field
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
type LogisticssupportListResponse struct {
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
	DataMode LogisticssupportListResponseDataMode `json:"dataMode,required"`
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
	LogisticsDiscrepancyInfos []LogisticssupportListResponseLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID string `json:"logisticsRecordId"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticssupportListResponseLogisticsRemark `json:"logisticsRemarks"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticssupportListResponseLogisticsSupportItem `json:"logisticsSupportItems"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticssupportListResponseLogisticsTransportationPlan `json:"logisticsTransportationPlans"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking        resp.Field
		DataMode                     resp.Field
		RptCreatedTime               resp.Field
		Source                       resp.Field
		ID                           resp.Field
		AircraftMds                  resp.Field
		CreatedAt                    resp.Field
		CreatedBy                    resp.Field
		CurrIcao                     resp.Field
		Etic                         resp.Field
		Etmc                         resp.Field
		ExtSystemID                  resp.Field
		LogisticAction               resp.Field
		LogisticsDiscrepancyInfos    resp.Field
		LogisticsRecordID            resp.Field
		LogisticsRemarks             resp.Field
		LogisticsSupportItems        resp.Field
		LogisticsTransportationPlans resp.Field
		MaintStatusCode              resp.Field
		McTime                       resp.Field
		MeTime                       resp.Field
		Origin                       resp.Field
		OrigNetwork                  resp.Field
		Owner                        resp.Field
		ReopenFlag                   resp.Field
		RptClosedTime                resp.Field
		SuppIcao                     resp.Field
		TailNumber                   resp.Field
		UpdatedAt                    resp.Field
		UpdatedBy                    resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponse) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportListResponse) UnmarshalJSON(data []byte) error {
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
type LogisticssupportListResponseDataMode string

const (
	LogisticssupportListResponseDataModeReal      LogisticssupportListResponseDataMode = "REAL"
	LogisticssupportListResponseDataModeTest      LogisticssupportListResponseDataMode = "TEST"
	LogisticssupportListResponseDataModeSimulated LogisticssupportListResponseDataMode = "SIMULATED"
	LogisticssupportListResponseDataModeExercise  LogisticssupportListResponseDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticssupportListResponseLogisticsDiscrepancyInfo struct {
	// The discrepancy closure time, in ISO 8601 UTC format with millisecond precision.
	ClosureTime time.Time `json:"closureTime" format:"date-time"`
	// The aircraft discrepancy description.
	DiscrepancyInfo string `json:"discrepancyInfo"`
	// Job Control Number of the discrepancy.
	Jcn string `json:"jcn"`
	// The job start time, in ISO 8601 UTC format with millisecond precision.
	JobStTime time.Time `json:"jobStTime" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClosureTime     resp.Field
		DiscrepancyInfo resp.Field
		Jcn             resp.Field
		JobStTime       resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsDiscrepancyInfo) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportListResponseLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportListResponseLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged time.Time `json:"lastChanged" format:"date-time"`
	// Text of the remark.
	Remark string `json:"remark"`
	// User who published the remark.
	Username string `json:"username"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		LastChanged resp.Field
		Remark      resp.Field
		Username    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsRemark) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportListResponseLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticssupportListResponseLogisticsSupportItem struct {
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
	LogisticsParts []LogisticssupportListResponseLogisticsSupportItemLogisticsPart `json:"logisticsParts"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticssupportListResponseLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticssupportListResponseLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Cannibalized            resp.Field
		DeployPlanNumber        resp.Field
		Description             resp.Field
		ItemLastChangedDate     resp.Field
		JobControlNumber        resp.Field
		LogisticsParts          resp.Field
		LogisticsRemarks        resp.Field
		LogisticsSpecialties    resp.Field
		Quantity                resp.Field
		ReadyTime               resp.Field
		ReceivedTime            resp.Field
		RecoveryRequestTypeCode resp.Field
		RedeployPlanNumber      resp.Field
		RedeployShipmentUnitID  resp.Field
		RequestNumber           resp.Field
		ResupportFlag           resp.Field
		ShipmentUnitID          resp.Field
		SiPoc                   resp.Field
		SourceIcao              resp.Field
		ExtraFields             map[string]resp.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsSupportItem) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportListResponseLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticssupportListResponseLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber string `json:"figureNumber"`
	// Technical order manual index number for the requested part.
	IndexNumber string `json:"indexNumber"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier string `json:"locationVerifier"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticssupportListResponseLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FigureNumber         resp.Field
		IndexNumber          resp.Field
		LocationVerifier     resp.Field
		LogisticsStocks      resp.Field
		MeasurementUnitCode  resp.Field
		NationalStockNumber  resp.Field
		PartNumber           resp.Field
		RequestVerifier      resp.Field
		SupplyDocumentNumber resp.Field
		TechnicalOrderText   resp.Field
		WorkUnitCode         resp.Field
		ExtraFields          map[string]resp.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsSupportItemLogisticsPart) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportListResponseLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticssupportListResponseLogisticsSupportItemLogisticsPartLogisticsStock struct {
	// The quantity of available parts needed from sourceICAO.
	Quantity int64 `json:"quantity"`
	// The ICAO code for the primary location with available parts.
	SourceIcao string `json:"sourceICAO"`
	// The datetime when the parts were sourced, in ISO 8601 UTC format with
	// millisecond precision.
	StockCheckTime time.Time `json:"stockCheckTime" format:"date-time"`
	// The point of contact at the sourced location.
	StockPoc string `json:"stockPOC"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Quantity       resp.Field
		SourceIcao     resp.Field
		StockCheckTime resp.Field
		StockPoc       resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsSupportItemLogisticsPartLogisticsStock) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportListResponseLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportListResponseLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged time.Time `json:"lastChanged" format:"date-time"`
	// Text of the remark.
	Remark string `json:"remark"`
	// User who published the remark.
	Username string `json:"username"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		LastChanged resp.Field
		Remark      resp.Field
		Username    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsSupportItemLogisticsRemark) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportListResponseLogisticsSupportItemLogisticsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticssupportListResponseLogisticsSupportItemLogisticsSpecialty struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FirstName    resp.Field
		Last4Ssn     resp.Field
		LastName     resp.Field
		RankCode     resp.Field
		RoleTypeCode resp.Field
		SkillLevel   resp.Field
		Specialty    resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsSupportItemLogisticsSpecialty) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportListResponseLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticssupportListResponseLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticssupportListResponseLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticssupportListResponseLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ActDepTime                          resp.Field
		AircraftStatus                      resp.Field
		ApproxArrTime                       resp.Field
		CancelledDate                       resp.Field
		ClosedDate                          resp.Field
		Coordinator                         resp.Field
		CoordinatorUnit                     resp.Field
		DestinationIcao                     resp.Field
		Duration                            resp.Field
		EstArrTime                          resp.Field
		EstDepTime                          resp.Field
		LastChangedDate                     resp.Field
		LogisticMasterRecordID              resp.Field
		LogisticsSegments                   resp.Field
		LogisticsTransportationPlansRemarks resp.Field
		Majcom                              resp.Field
		MissionChange                       resp.Field
		NumEnrouteStops                     resp.Field
		NumTransLoads                       resp.Field
		OriginIcao                          resp.Field
		PlanDefinition                      resp.Field
		PlansNumber                         resp.Field
		SerialNumber                        resp.Field
		StatusCode                          resp.Field
		TpAircraftMds                       resp.Field
		TpTailNumber                        resp.Field
		ExtraFields                         map[string]resp.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsTransportationPlan) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportListResponseLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportListResponseLogisticsTransportationPlanLogisticsSegment struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ArrivalIcao    resp.Field
		DepartureIcao  resp.Field
		ExtMissionID   resp.Field
		IDMission      resp.Field
		Itin           resp.Field
		MissionNumber  resp.Field
		MissionType    resp.Field
		ModeCode       resp.Field
		SegActArrTime  resp.Field
		SegActDepTime  resp.Field
		SegAircraftMds resp.Field
		SegEstArrTime  resp.Field
		SegEstDepTime  resp.Field
		SegmentNumber  resp.Field
		SegTailNumber  resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsTransportationPlanLogisticsSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportListResponseLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportListResponseLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged time.Time `json:"lastChanged" format:"date-time"`
	// Text of the remark.
	Remark string `json:"remark"`
	// User who published the remark.
	Username string `json:"username"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		LastChanged resp.Field
		Remark      resp.Field
		Username    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportListResponseLogisticsTransportationPlanLogisticsTransportationPlansRemark) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportListResponseLogisticsTransportationPlanLogisticsTransportationPlansRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
type LogisticssupportGetResponse struct {
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
	DataMode LogisticssupportGetResponseDataMode `json:"dataMode,required"`
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
	LogisticsDiscrepancyInfos []LogisticssupportGetResponseLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID string `json:"logisticsRecordId"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticssupportGetResponseLogisticsSupportItem `json:"logisticsSupportItems"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticssupportGetResponseLogisticsTransportationPlan `json:"logisticsTransportationPlans"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking        resp.Field
		DataMode                     resp.Field
		RptCreatedTime               resp.Field
		Source                       resp.Field
		ID                           resp.Field
		AircraftMds                  resp.Field
		CreatedAt                    resp.Field
		CreatedBy                    resp.Field
		CurrIcao                     resp.Field
		Etic                         resp.Field
		Etmc                         resp.Field
		ExtSystemID                  resp.Field
		LogisticAction               resp.Field
		LogisticsDiscrepancyInfos    resp.Field
		LogisticsRecordID            resp.Field
		LogisticsRemarks             resp.Field
		LogisticsSupportItems        resp.Field
		LogisticsTransportationPlans resp.Field
		MaintStatusCode              resp.Field
		McTime                       resp.Field
		MeTime                       resp.Field
		Origin                       resp.Field
		OrigNetwork                  resp.Field
		Owner                        resp.Field
		ReopenFlag                   resp.Field
		RptClosedTime                resp.Field
		SuppIcao                     resp.Field
		TailNumber                   resp.Field
		UpdatedAt                    resp.Field
		UpdatedBy                    resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportGetResponse) UnmarshalJSON(data []byte) error {
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
type LogisticssupportGetResponseDataMode string

const (
	LogisticssupportGetResponseDataModeReal      LogisticssupportGetResponseDataMode = "REAL"
	LogisticssupportGetResponseDataModeTest      LogisticssupportGetResponseDataMode = "TEST"
	LogisticssupportGetResponseDataModeSimulated LogisticssupportGetResponseDataMode = "SIMULATED"
	LogisticssupportGetResponseDataModeExercise  LogisticssupportGetResponseDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticssupportGetResponseLogisticsDiscrepancyInfo struct {
	// The discrepancy closure time, in ISO 8601 UTC format with millisecond precision.
	ClosureTime time.Time `json:"closureTime" format:"date-time"`
	// The aircraft discrepancy description.
	DiscrepancyInfo string `json:"discrepancyInfo"`
	// Job Control Number of the discrepancy.
	Jcn string `json:"jcn"`
	// The job start time, in ISO 8601 UTC format with millisecond precision.
	JobStTime time.Time `json:"jobStTime" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClosureTime     resp.Field
		DiscrepancyInfo resp.Field
		Jcn             resp.Field
		JobStTime       resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportGetResponseLogisticsDiscrepancyInfo) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportGetResponseLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticssupportGetResponseLogisticsSupportItem struct {
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
	LogisticsParts []LogisticssupportGetResponseLogisticsSupportItemLogisticsPart `json:"logisticsParts"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticssupportGetResponseLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Cannibalized            resp.Field
		DeployPlanNumber        resp.Field
		Description             resp.Field
		ItemLastChangedDate     resp.Field
		JobControlNumber        resp.Field
		LogisticsParts          resp.Field
		LogisticsRemarks        resp.Field
		LogisticsSpecialties    resp.Field
		Quantity                resp.Field
		ReadyTime               resp.Field
		ReceivedTime            resp.Field
		RecoveryRequestTypeCode resp.Field
		RedeployPlanNumber      resp.Field
		RedeployShipmentUnitID  resp.Field
		RequestNumber           resp.Field
		ResupportFlag           resp.Field
		ShipmentUnitID          resp.Field
		SiPoc                   resp.Field
		SourceIcao              resp.Field
		ExtraFields             map[string]resp.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportGetResponseLogisticsSupportItem) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportGetResponseLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticssupportGetResponseLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber string `json:"figureNumber"`
	// Technical order manual index number for the requested part.
	IndexNumber string `json:"indexNumber"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier string `json:"locationVerifier"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticssupportGetResponseLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FigureNumber         resp.Field
		IndexNumber          resp.Field
		LocationVerifier     resp.Field
		LogisticsStocks      resp.Field
		MeasurementUnitCode  resp.Field
		NationalStockNumber  resp.Field
		PartNumber           resp.Field
		RequestVerifier      resp.Field
		SupplyDocumentNumber resp.Field
		TechnicalOrderText   resp.Field
		WorkUnitCode         resp.Field
		ExtraFields          map[string]resp.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportGetResponseLogisticsSupportItemLogisticsPart) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportGetResponseLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticssupportGetResponseLogisticsSupportItemLogisticsPartLogisticsStock struct {
	// The quantity of available parts needed from sourceICAO.
	Quantity int64 `json:"quantity"`
	// The ICAO code for the primary location with available parts.
	SourceIcao string `json:"sourceICAO"`
	// The datetime when the parts were sourced, in ISO 8601 UTC format with
	// millisecond precision.
	StockCheckTime time.Time `json:"stockCheckTime" format:"date-time"`
	// The point of contact at the sourced location.
	StockPoc string `json:"stockPOC"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Quantity       resp.Field
		SourceIcao     resp.Field
		StockCheckTime resp.Field
		StockPoc       resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportGetResponseLogisticsSupportItemLogisticsPartLogisticsStock) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportGetResponseLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticssupportGetResponseLogisticsSupportItemLogisticsSpecialty struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FirstName    resp.Field
		Last4Ssn     resp.Field
		LastName     resp.Field
		RankCode     resp.Field
		RoleTypeCode resp.Field
		SkillLevel   resp.Field
		Specialty    resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportGetResponseLogisticsSupportItemLogisticsSpecialty) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportGetResponseLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticssupportGetResponseLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticssupportGetResponseLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ActDepTime                          resp.Field
		AircraftStatus                      resp.Field
		ApproxArrTime                       resp.Field
		CancelledDate                       resp.Field
		ClosedDate                          resp.Field
		Coordinator                         resp.Field
		CoordinatorUnit                     resp.Field
		DestinationIcao                     resp.Field
		Duration                            resp.Field
		EstArrTime                          resp.Field
		EstDepTime                          resp.Field
		LastChangedDate                     resp.Field
		LogisticMasterRecordID              resp.Field
		LogisticsSegments                   resp.Field
		LogisticsTransportationPlansRemarks resp.Field
		Majcom                              resp.Field
		MissionChange                       resp.Field
		NumEnrouteStops                     resp.Field
		NumTransLoads                       resp.Field
		OriginIcao                          resp.Field
		PlanDefinition                      resp.Field
		PlansNumber                         resp.Field
		SerialNumber                        resp.Field
		StatusCode                          resp.Field
		TpAircraftMds                       resp.Field
		TpTailNumber                        resp.Field
		ExtraFields                         map[string]resp.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportGetResponseLogisticsTransportationPlan) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportGetResponseLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportGetResponseLogisticsTransportationPlanLogisticsSegment struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ArrivalIcao    resp.Field
		DepartureIcao  resp.Field
		ExtMissionID   resp.Field
		IDMission      resp.Field
		Itin           resp.Field
		MissionNumber  resp.Field
		MissionType    resp.Field
		ModeCode       resp.Field
		SegActArrTime  resp.Field
		SegActDepTime  resp.Field
		SegAircraftMds resp.Field
		SegEstArrTime  resp.Field
		SegEstDepTime  resp.Field
		SegmentNumber  resp.Field
		SegTailNumber  resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportGetResponseLogisticsTransportationPlanLogisticsSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportGetResponseLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
type LogisticssupportTupleResponse struct {
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
	DataMode LogisticssupportTupleResponseDataMode `json:"dataMode,required"`
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
	LogisticsDiscrepancyInfos []LogisticssupportTupleResponseLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos"`
	// The identifier that represents a Logistics Master Record.
	LogisticsRecordID string `json:"logisticsRecordId"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticssupportTupleResponseLogisticsSupportItem `json:"logisticsSupportItems"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticssupportTupleResponseLogisticsTransportationPlan `json:"logisticsTransportationPlans"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking        resp.Field
		DataMode                     resp.Field
		RptCreatedTime               resp.Field
		Source                       resp.Field
		ID                           resp.Field
		AircraftMds                  resp.Field
		CreatedAt                    resp.Field
		CreatedBy                    resp.Field
		CurrIcao                     resp.Field
		Etic                         resp.Field
		Etmc                         resp.Field
		ExtSystemID                  resp.Field
		LogisticAction               resp.Field
		LogisticsDiscrepancyInfos    resp.Field
		LogisticsRecordID            resp.Field
		LogisticsRemarks             resp.Field
		LogisticsSupportItems        resp.Field
		LogisticsTransportationPlans resp.Field
		MaintStatusCode              resp.Field
		McTime                       resp.Field
		MeTime                       resp.Field
		Origin                       resp.Field
		OrigNetwork                  resp.Field
		Owner                        resp.Field
		ReopenFlag                   resp.Field
		RptClosedTime                resp.Field
		SuppIcao                     resp.Field
		TailNumber                   resp.Field
		UpdatedAt                    resp.Field
		UpdatedBy                    resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportTupleResponse) UnmarshalJSON(data []byte) error {
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
type LogisticssupportTupleResponseDataMode string

const (
	LogisticssupportTupleResponseDataModeReal      LogisticssupportTupleResponseDataMode = "REAL"
	LogisticssupportTupleResponseDataModeTest      LogisticssupportTupleResponseDataMode = "TEST"
	LogisticssupportTupleResponseDataModeSimulated LogisticssupportTupleResponseDataMode = "SIMULATED"
	LogisticssupportTupleResponseDataModeExercise  LogisticssupportTupleResponseDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticssupportTupleResponseLogisticsDiscrepancyInfo struct {
	// The discrepancy closure time, in ISO 8601 UTC format with millisecond precision.
	ClosureTime time.Time `json:"closureTime" format:"date-time"`
	// The aircraft discrepancy description.
	DiscrepancyInfo string `json:"discrepancyInfo"`
	// Job Control Number of the discrepancy.
	Jcn string `json:"jcn"`
	// The job start time, in ISO 8601 UTC format with millisecond precision.
	JobStTime time.Time `json:"jobStTime" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClosureTime     resp.Field
		DiscrepancyInfo resp.Field
		Jcn             resp.Field
		JobStTime       resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportTupleResponseLogisticsDiscrepancyInfo) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportTupleResponseLogisticsDiscrepancyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Support items associated with this LogisticsSupport record.
type LogisticssupportTupleResponseLogisticsSupportItem struct {
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
	LogisticsParts []LogisticssupportTupleResponseLogisticsSupportItemLogisticsPart `json:"logisticsParts"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticsRemarksFull `json:"logisticsRemarks"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticssupportTupleResponseLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Cannibalized            resp.Field
		DeployPlanNumber        resp.Field
		Description             resp.Field
		ItemLastChangedDate     resp.Field
		JobControlNumber        resp.Field
		LogisticsParts          resp.Field
		LogisticsRemarks        resp.Field
		LogisticsSpecialties    resp.Field
		Quantity                resp.Field
		ReadyTime               resp.Field
		ReceivedTime            resp.Field
		RecoveryRequestTypeCode resp.Field
		RedeployPlanNumber      resp.Field
		RedeployShipmentUnitID  resp.Field
		RequestNumber           resp.Field
		ResupportFlag           resp.Field
		ShipmentUnitID          resp.Field
		SiPoc                   resp.Field
		SourceIcao              resp.Field
		ExtraFields             map[string]resp.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportTupleResponseLogisticsSupportItem) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportTupleResponseLogisticsSupportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parts associated with this support item.
type LogisticssupportTupleResponseLogisticsSupportItemLogisticsPart struct {
	// Technical order manual figure number for the requested / supplied part.
	FigureNumber string `json:"figureNumber"`
	// Technical order manual index number for the requested part.
	IndexNumber string `json:"indexNumber"`
	// The person who validated that the sourced location has, and can supply, the
	// requested parts.
	LocationVerifier string `json:"locationVerifier"`
	// The supply stocks for this support item.
	LogisticsStocks []LogisticssupportTupleResponseLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FigureNumber         resp.Field
		IndexNumber          resp.Field
		LocationVerifier     resp.Field
		LogisticsStocks      resp.Field
		MeasurementUnitCode  resp.Field
		NationalStockNumber  resp.Field
		PartNumber           resp.Field
		RequestVerifier      resp.Field
		SupplyDocumentNumber resp.Field
		TechnicalOrderText   resp.Field
		WorkUnitCode         resp.Field
		ExtraFields          map[string]resp.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportTupleResponseLogisticsSupportItemLogisticsPart) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportTupleResponseLogisticsSupportItemLogisticsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supply stocks for this support item.
type LogisticssupportTupleResponseLogisticsSupportItemLogisticsPartLogisticsStock struct {
	// The quantity of available parts needed from sourceICAO.
	Quantity int64 `json:"quantity"`
	// The ICAO code for the primary location with available parts.
	SourceIcao string `json:"sourceICAO"`
	// The datetime when the parts were sourced, in ISO 8601 UTC format with
	// millisecond precision.
	StockCheckTime time.Time `json:"stockCheckTime" format:"date-time"`
	// The point of contact at the sourced location.
	StockPoc string `json:"stockPOC"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Quantity       resp.Field
		SourceIcao     resp.Field
		StockCheckTime resp.Field
		StockPoc       resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportTupleResponseLogisticsSupportItemLogisticsPartLogisticsStock) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportTupleResponseLogisticsSupportItemLogisticsPartLogisticsStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The specialties required to implement this support item.
type LogisticssupportTupleResponseLogisticsSupportItemLogisticsSpecialty struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FirstName    resp.Field
		Last4Ssn     resp.Field
		LastName     resp.Field
		RankCode     resp.Field
		RoleTypeCode resp.Field
		SkillLevel   resp.Field
		Specialty    resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportTupleResponseLogisticsSupportItemLogisticsSpecialty) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportTupleResponseLogisticsSupportItemLogisticsSpecialty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticssupportTupleResponseLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticssupportTupleResponseLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ActDepTime                          resp.Field
		AircraftStatus                      resp.Field
		ApproxArrTime                       resp.Field
		CancelledDate                       resp.Field
		ClosedDate                          resp.Field
		Coordinator                         resp.Field
		CoordinatorUnit                     resp.Field
		DestinationIcao                     resp.Field
		Duration                            resp.Field
		EstArrTime                          resp.Field
		EstDepTime                          resp.Field
		LastChangedDate                     resp.Field
		LogisticMasterRecordID              resp.Field
		LogisticsSegments                   resp.Field
		LogisticsTransportationPlansRemarks resp.Field
		Majcom                              resp.Field
		MissionChange                       resp.Field
		NumEnrouteStops                     resp.Field
		NumTransLoads                       resp.Field
		OriginIcao                          resp.Field
		PlanDefinition                      resp.Field
		PlansNumber                         resp.Field
		SerialNumber                        resp.Field
		StatusCode                          resp.Field
		TpAircraftMds                       resp.Field
		TpTailNumber                        resp.Field
		ExtraFields                         map[string]resp.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportTupleResponseLogisticsTransportationPlan) RawJSON() string { return r.JSON.raw }
func (r *LogisticssupportTupleResponseLogisticsTransportationPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportTupleResponseLogisticsTransportationPlanLogisticsSegment struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ArrivalIcao    resp.Field
		DepartureIcao  resp.Field
		ExtMissionID   resp.Field
		IDMission      resp.Field
		Itin           resp.Field
		MissionNumber  resp.Field
		MissionType    resp.Field
		ModeCode       resp.Field
		SegActArrTime  resp.Field
		SegActDepTime  resp.Field
		SegAircraftMds resp.Field
		SegEstArrTime  resp.Field
		SegEstDepTime  resp.Field
		SegmentNumber  resp.Field
		SegTailNumber  resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogisticssupportTupleResponseLogisticsTransportationPlanLogisticsSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *LogisticssupportTupleResponseLogisticsTransportationPlanLogisticsSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogisticssupportNewParams struct {
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
	DataMode LogisticssupportNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	LogisticsDiscrepancyInfos []LogisticssupportNewParamsLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos,omitzero"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticssupportNewParamsLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticssupportNewParamsLogisticsSupportItem `json:"logisticsSupportItems,omitzero"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticssupportNewParamsLogisticsTransportationPlan `json:"logisticsTransportationPlans,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LogisticssupportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type LogisticssupportNewParamsDataMode string

const (
	LogisticssupportNewParamsDataModeReal      LogisticssupportNewParamsDataMode = "REAL"
	LogisticssupportNewParamsDataModeTest      LogisticssupportNewParamsDataMode = "TEST"
	LogisticssupportNewParamsDataModeSimulated LogisticssupportNewParamsDataMode = "SIMULATED"
	LogisticssupportNewParamsDataModeExercise  LogisticssupportNewParamsDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticssupportNewParamsLogisticsDiscrepancyInfo struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsDiscrepancyInfo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsDiscrepancyInfo) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsDiscrepancyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportNewParamsLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

// Support items associated with this LogisticsSupport record.
type LogisticssupportNewParamsLogisticsSupportItem struct {
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
	LogisticsParts []LogisticssupportNewParamsLogisticsSupportItemLogisticsPart `json:"logisticsParts,omitzero"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticssupportNewParamsLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticssupportNewParamsLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsSupportItem) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsSupportItem) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsSupportItem
	return param.MarshalObject(r, (*shadow)(&r))
}

// The parts associated with this support item.
type LogisticssupportNewParamsLogisticsSupportItemLogisticsPart struct {
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
	LogisticsStocks []LogisticssupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsSupportItemLogisticsPart) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsSupportItemLogisticsPart) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsSupportItemLogisticsPart
	return param.MarshalObject(r, (*shadow)(&r))
}

// The supply stocks for this support item.
type LogisticssupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportNewParamsLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsSupportItemLogisticsRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsSupportItemLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsSupportItemLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

// The specialties required to implement this support item.
type LogisticssupportNewParamsLogisticsSupportItemLogisticsSpecialty struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsSupportItemLogisticsSpecialty) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsSupportItemLogisticsSpecialty) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsSupportItemLogisticsSpecialty
	return param.MarshalObject(r, (*shadow)(&r))
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticssupportNewParamsLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticssupportNewParamsLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments,omitzero"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticssupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsTransportationPlan) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsTransportationPlan) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsTransportationPlan
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportNewParamsLogisticsTransportationPlanLogisticsSegment struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsTransportationPlanLogisticsSegment) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsTransportationPlanLogisticsSegment) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsTransportationPlanLogisticsSegment
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

type LogisticssupportUpdateParams struct {
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
	DataMode LogisticssupportUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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
	LogisticsDiscrepancyInfos []LogisticssupportUpdateParamsLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos,omitzero"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticssupportUpdateParamsLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticssupportUpdateParamsLogisticsSupportItem `json:"logisticsSupportItems,omitzero"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticssupportUpdateParamsLogisticsTransportationPlan `json:"logisticsTransportationPlans,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LogisticssupportUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type LogisticssupportUpdateParamsDataMode string

const (
	LogisticssupportUpdateParamsDataModeReal      LogisticssupportUpdateParamsDataMode = "REAL"
	LogisticssupportUpdateParamsDataModeTest      LogisticssupportUpdateParamsDataMode = "TEST"
	LogisticssupportUpdateParamsDataModeSimulated LogisticssupportUpdateParamsDataMode = "SIMULATED"
	LogisticssupportUpdateParamsDataModeExercise  LogisticssupportUpdateParamsDataMode = "EXERCISE"
)

// Discrepancy information associated with this LogisticsSupport record.
type LogisticssupportUpdateParamsLogisticsDiscrepancyInfo struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsDiscrepancyInfo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsDiscrepancyInfo) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsDiscrepancyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportUpdateParamsLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

// Support items associated with this LogisticsSupport record.
type LogisticssupportUpdateParamsLogisticsSupportItem struct {
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
	LogisticsParts []LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPart `json:"logisticsParts,omitzero"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticssupportUpdateParamsLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticssupportUpdateParamsLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsSupportItem) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsSupportItem) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsSupportItem
	return param.MarshalObject(r, (*shadow)(&r))
}

// The parts associated with this support item.
type LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPart struct {
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
	LogisticsStocks []LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPart) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPart) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPart
	return param.MarshalObject(r, (*shadow)(&r))
}

// The supply stocks for this support item.
type LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportUpdateParamsLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsSupportItemLogisticsRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsSupportItemLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsSupportItemLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

// The specialties required to implement this support item.
type LogisticssupportUpdateParamsLogisticsSupportItemLogisticsSpecialty struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsSupportItemLogisticsSpecialty) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsSupportItemLogisticsSpecialty) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsSupportItemLogisticsSpecialty
	return param.MarshalObject(r, (*shadow)(&r))
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticssupportUpdateParamsLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments,omitzero"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsTransportationPlan) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsTransportationPlan) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsTransportationPlan
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsSegment struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsSegment) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsSegment) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsSegment
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

type LogisticssupportListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LogisticssupportListParams]'s query parameters as
// `url.Values`.
func (r LogisticssupportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticssupportCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LogisticssupportCountParams]'s query parameters as
// `url.Values`.
func (r LogisticssupportCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticssupportNewBulkParams struct {
	Body []LogisticssupportNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LogisticssupportNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
//
// The properties ClassificationMarking, DataMode, RptCreatedTime, Source are
// required.
type LogisticssupportNewBulkParamsBody struct {
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
	LogisticsDiscrepancyInfos []LogisticssupportNewBulkParamsBodyLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos,omitzero"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticssupportNewBulkParamsBodyLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticssupportNewBulkParamsBodyLogisticsSupportItem `json:"logisticsSupportItems,omitzero"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticssupportNewBulkParamsBodyLogisticsTransportationPlan `json:"logisticsTransportationPlans,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[LogisticssupportNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Discrepancy information associated with this LogisticsSupport record.
type LogisticssupportNewBulkParamsBodyLogisticsDiscrepancyInfo struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsDiscrepancyInfo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsDiscrepancyInfo) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsDiscrepancyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportNewBulkParamsBodyLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

// Support items associated with this LogisticsSupport record.
type LogisticssupportNewBulkParamsBodyLogisticsSupportItem struct {
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
	LogisticsParts []LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart `json:"logisticsParts,omitzero"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsSupportItem) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsSupportItem) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsSupportItem
	return param.MarshalObject(r, (*shadow)(&r))
}

// The parts associated with this support item.
type LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart struct {
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
	LogisticsStocks []LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart
	return param.MarshalObject(r, (*shadow)(&r))
}

// The supply stocks for this support item.
type LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

// The specialties required to implement this support item.
type LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty
	return param.MarshalObject(r, (*shadow)(&r))
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticssupportNewBulkParamsBodyLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments,omitzero"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsTransportationPlan) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsTransportationPlan) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsTransportationPlan
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

type LogisticssupportGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LogisticssupportGetParams]'s query parameters as
// `url.Values`.
func (r LogisticssupportGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticssupportTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LogisticssupportTupleParams]'s query parameters as
// `url.Values`.
func (r LogisticssupportTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogisticssupportUnvalidatedPublishParams struct {
	Body []LogisticssupportUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r LogisticssupportUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Comprehensive logistical details concerning the planned support of maintenance
// operations required by an aircraft, including transportation information,
// supplies coordination, and service personnel.
//
// The properties ClassificationMarking, DataMode, RptCreatedTime, Source are
// required.
type LogisticssupportUnvalidatedPublishParamsBody struct {
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
	LogisticsDiscrepancyInfos []LogisticssupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo `json:"logisticsDiscrepancyInfos,omitzero"`
	// Remarks associated with this LogisticsSupport record.
	LogisticsRemarks []LogisticssupportUnvalidatedPublishParamsBodyLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// Support items associated with this LogisticsSupport record.
	LogisticsSupportItems []LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItem `json:"logisticsSupportItems,omitzero"`
	// Transportation plans associated with this LogisticsSupport record, used to
	// coordinate maintenance efforts.
	LogisticsTransportationPlans []LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan `json:"logisticsTransportationPlans,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[LogisticssupportUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Discrepancy information associated with this LogisticsSupport record.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

// Support items associated with this LogisticsSupport record.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItem struct {
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
	LogisticsParts []LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart `json:"logisticsParts,omitzero"`
	// Remarks associated with this support item.
	LogisticsRemarks []LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark `json:"logisticsRemarks,omitzero"`
	// The specialties required to implement this support item.
	LogisticsSpecialties []LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty `json:"logisticsSpecialties,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItem) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItem) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItem
	return param.MarshalObject(r, (*shadow)(&r))
}

// The parts associated with this support item.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart struct {
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
	LogisticsStocks []LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock `json:"logisticsStocks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart
	return param.MarshalObject(r, (*shadow)(&r))
}

// The supply stocks for this support item.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

// The specialties required to implement this support item.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty
	return param.MarshalObject(r, (*shadow)(&r))
}

// Transportation plans associated with this LogisticsSupport record, used to
// coordinate maintenance efforts.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan struct {
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
	LogisticsSegments []LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment `json:"logisticsSegments,omitzero"`
	// Remarks associated with this transportation plan.
	LogisticsTransportationPlansRemarks []LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark `json:"logisticsTransportationPlansRemarks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment
	return param.MarshalObject(r, (*shadow)(&r))
}

// Remarks associated with this LogisticsSupport record.
type LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark struct {
	// Date the remark was published or updated, in ISO 8601 UTC format, with
	// millisecond precision.
	LastChanged param.Opt[time.Time] `json:"lastChanged,omitzero" format:"date-time"`
	// Text of the remark.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// User who published the remark.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark) MarshalJSON() (data []byte, err error) {
	type shadow LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
