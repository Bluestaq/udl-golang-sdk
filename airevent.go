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

// AirEventService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirEventService] method instead.
type AirEventService struct {
	Options []option.RequestOption
}

// NewAirEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAirEventService(opts ...option.RequestOption) (r AirEventService) {
	r = AirEventService{}
	r.Options = opts
	return
}

// Service operation to take a single airevent object as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AirEventService) New(ctx context.Context, body AirEventNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airevent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single airevent record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *AirEventService) Update(ctx context.Context, id string, body AirEventUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airevent/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirEventService) List(ctx context.Context, query AirEventListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AirEventListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/airevent"
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
func (r *AirEventService) ListAutoPaging(ctx context.Context, query AirEventListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AirEventListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an airevent record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *AirEventService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airevent/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirEventService) Count(ctx context.Context, query AirEventCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airevent/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list
// ofService operation intended for initial integration only, to take a list of
// airevent records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *AirEventService) NewBulk(ctx context.Context, body AirEventNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airevent/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single airevent record by its unique ID passed as a
// path parameter.
func (r *AirEventService) Get(ctx context.Context, id string, query AirEventGetParams, opts ...option.RequestOption) (res *AirEventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airevent/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AirEventService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *AirEventQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airevent/queryhelp"
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
func (r *AirEventService) Tuple(ctx context.Context, query AirEventTupleParams, opts ...option.RequestOption) (res *[]AirEventTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airevent/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple airevent records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *AirEventService) UnvalidatedPublish(ctx context.Context, body AirEventUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-airevent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Information related to an air event (e.g. FUEL TRANSFER, AIR DROP) and the
// associated aircraft.
type AirEventListResponse struct {
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
	DataMode AirEventListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of air event (e.g. FUEL TRANSFER, AIR DROP, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The actual arrival time of the aircraft at the air event, in ISO 8601 UTC format
	// with millisecond precision.
	ActualArrTime time.Time `json:"actualArrTime" format:"date-time"`
	// The actual departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	ActualDepTime time.Time `json:"actualDepTime" format:"date-time"`
	// The Air Refueling Control Time is the planned time the tanker aircraft will
	// transfer fuel to the receiver aircraft, in ISO 8601 UTC format, with millisecond
	// precision.
	Arct time.Time `json:"arct" format:"date-time"`
	// Type of process used by AMC to schedule this air refueling event. Possible
	// values are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched
	// Theater Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S
	// (Soft Air Refueling), T (Matched Theater Operation Short Notice (Theater
	// Assets)), V (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short
	// Notice (AMC Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z
	// (Other Air Refueling).
	ArEventType string `json:"arEventType"`
	// The purpose of the air event at the arrival location. Can be either descriptive
	// text such as 'fuel onload' or a purpose code specified by the provider, such as
	// 'A'.
	ArrPurpose string `json:"arrPurpose"`
	// Identifier of the air refueling track, if applicable.
	ArTrackID string `json:"arTrackId"`
	// Name of the air refueling track, if applicable.
	ArTrackName string `json:"arTrackName"`
	// Altitude of this air event, in feet.
	BaseAlt float64 `json:"baseAlt"`
	// Flag indicating that this air refueling event has been cancelled.
	Cancelled bool `json:"cancelled"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The purpose of the air event at the departure location. Can be either
	// descriptive text such as 'fuel onload' or a purpose code specified by the
	// provider, such as 'A'.
	DepPurpose string `json:"depPurpose"`
	// The current estimated arrival time of the aircraft at the air event, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime time.Time `json:"estArrTime" format:"date-time"`
	// The current estimated departure time of the aircraft from the air event, in ISO
	// 8601 UTC format with millisecond precision.
	EstDepTime time.Time `json:"estDepTime" format:"date-time"`
	// Optional air event ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalAirEventID string `json:"externalAirEventId"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalArTrackID string `json:"externalARTrackId"`
	// The UDL unique identifier of the mission associated with this air event.
	IDMission string `json:"idMission"`
	// The UDL unique identifier of the sortie associated with this air event.
	IDSortie string `json:"idSortie"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum int64 `json:"legNum"`
	// The location representing this air event specified as a feature Id. Locations
	// specified include air refueling track Ids and air drop event locations.
	Location string `json:"location"`
	// The number of tankers requested for an air refueling event.
	NumTankers int64 `json:"numTankers"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The scheduled arrival time of the aircraft at the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedArrTime time.Time `json:"plannedArrTime" format:"date-time"`
	// The scheduled departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedDepTime time.Time `json:"plannedDepTime" format:"date-time"`
	// Priority of this air event.
	Priority string `json:"priority"`
	// Collection of receiver aircraft associated with this Air Event.
	Receivers []AirEventListResponseReceiver `json:"receivers"`
	// Collection of remarks associated with this Air Event.
	Remarks []AirEventListResponseRemark `json:"remarks"`
	// Flag indicating if the receiver unit has requested flying an air refueling track
	// in both directions.
	RevTrack bool `json:"revTrack"`
	// The Rendezvous Control Time is the planned time the tanker and receiver aircraft
	// will rendezvous for an en route type air refueling event, in ISO 8601 UTC
	// format, with millisecond precision.
	Rzct time.Time `json:"rzct" format:"date-time"`
	// Rendezvous point for the tanker and receiver during this air refueling event.
	// Possible values are AN (Anchor Nav Point), AP (Anchor Pattern), CP (Control
	// Point), ET (Entry Point), EX (Exit Point), IP (Initial Point), NC (Nav Check
	// Point).
	RzPoint string `json:"rzPoint"`
	// Type of rendezvous used for this air refueling event. Possible values are BUD
	// (Buddy), EN (Enroute), GCI (Ground Control), PP (Point Parallel).
	RzType string `json:"rzType"`
	// Flag indicating that the receiver unit has requested flying a short portion of
	// an air refueling track.
	ShortTrack bool `json:"shortTrack"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Status of this air refueling event track reservation. Receivers are responsible
	// for scheduling or reserving air refueling tracks. Possible values are A
	// (Altitude Reservation), R (Reserved), or Q (Questionable).
	StatusCode string `json:"statusCode"`
	// Collection of tanker aircraft associated with this Air Event.
	Tankers []AirEventListResponseTanker `json:"tankers"`
	// Length of time the receiver unit has requested for an air event, in hours.
	TrackTime float64 `json:"trackTime"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		ActualArrTime         respjson.Field
		ActualDepTime         respjson.Field
		Arct                  respjson.Field
		ArEventType           respjson.Field
		ArrPurpose            respjson.Field
		ArTrackID             respjson.Field
		ArTrackName           respjson.Field
		BaseAlt               respjson.Field
		Cancelled             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DepPurpose            respjson.Field
		EstArrTime            respjson.Field
		EstDepTime            respjson.Field
		ExternalAirEventID    respjson.Field
		ExternalArTrackID     respjson.Field
		IDMission             respjson.Field
		IDSortie              respjson.Field
		LegNum                respjson.Field
		Location              respjson.Field
		NumTankers            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PlannedArrTime        respjson.Field
		PlannedDepTime        respjson.Field
		Priority              respjson.Field
		Receivers             respjson.Field
		Remarks               respjson.Field
		RevTrack              respjson.Field
		Rzct                  respjson.Field
		RzPoint               respjson.Field
		RzType                respjson.Field
		ShortTrack            respjson.Field
		SourceDl              respjson.Field
		StatusCode            respjson.Field
		Tankers               respjson.Field
		TrackTime             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *AirEventListResponse) UnmarshalJSON(data []byte) error {
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
type AirEventListResponseDataMode string

const (
	AirEventListResponseDataModeReal      AirEventListResponseDataMode = "REAL"
	AirEventListResponseDataModeTest      AirEventListResponseDataMode = "TEST"
	AirEventListResponseDataModeSimulated AirEventListResponseDataMode = "SIMULATED"
	AirEventListResponseDataModeExercise  AirEventListResponseDataMode = "EXERCISE"
)

// Collection of receiver aircraft associated with this Air Event.
type AirEventListResponseReceiver struct {
	// Alternate mission identifier of this receiver provided by source.
	AltReceiverMissionID string `json:"altReceiverMissionId"`
	// The Air Mobility Command (AMC) mission identifier of this receiver.
	AmcReceiverMissionID string `json:"amcReceiverMissionId"`
	// Optional receiver identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalReceiverID string `json:"externalReceiverId"`
	// Total weight of the fuel transferred to this receiver during an air refueling
	// event, in pounds.
	FuelOn float64 `json:"fuelOn"`
	// The UDL ID of the airfield this receiver is associated with for this event.
	IDReceiverAirfield string `json:"idReceiverAirfield"`
	// The UDL ID of the mission this receiver is associated with for this event.
	IDReceiverMission string `json:"idReceiverMission"`
	// The UDL ID of the aircraft sortie this receiver is associated with for this
	// event.
	IDReceiverSortie string `json:"idReceiverSortie"`
	// Number of aircraft contained within one receiver coordination record for an air
	// refueling event.
	NumRecAircraft int64 `json:"numRecAircraft"`
	// The package identifier for the receiver in an air refueling event.
	PackageID string `json:"packageId"`
	// The call sign assigned to this receiver.
	ReceiverCallSign string `json:"receiverCallSign"`
	// Position of this receiver within a group of receivers in an air refueling event.
	ReceiverCellPosition int64 `json:"receiverCellPosition"`
	// Coordination record identifier of this receiver.
	ReceiverCoord string `json:"receiverCoord"`
	// Type of fuel delivery method used by the receiver during an air refueling event
	// (BOOM, DROGUE, BOTH).
	ReceiverDeliveryMethod string `json:"receiverDeliveryMethod"`
	// Location the receiver is deployed to for an air refueling event.
	ReceiverDeployedIcao string `json:"receiverDeployedICAO"`
	// Name of the receiver exercise associated with an air refueling event.
	ReceiverExercise string `json:"receiverExercise"`
	// Type of fuel being transferred to the receiver in an air refueling event.
	ReceiverFuelType string `json:"receiverFuelType"`
	// Identifies the itinerary point of a mission that this receiver is linked to.
	ReceiverLegNum int64 `json:"receiverLegNum"`
	// The Model Design Series designation of this receiver.
	ReceiverMds string `json:"receiverMDS"`
	// The wing or unit that owns this receiver.
	ReceiverOwner string `json:"receiverOwner"`
	// The name and/or number of the point of contact for this receiver.
	ReceiverPoc string `json:"receiverPOC"`
	// The major command level (MAJCOM) or foreign military sales (FMS) name of the
	// receiver's organization. The tanker flying hours used for an air refueling event
	// are logged against the receiver MAJCOM or foreign government being supported.
	RecOrg string `json:"recOrg"`
	// Indicates the unique number by Unit ID, which identifies an air refueling event.
	SequenceNum string `json:"sequenceNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltReceiverMissionID   respjson.Field
		AmcReceiverMissionID   respjson.Field
		ExternalReceiverID     respjson.Field
		FuelOn                 respjson.Field
		IDReceiverAirfield     respjson.Field
		IDReceiverMission      respjson.Field
		IDReceiverSortie       respjson.Field
		NumRecAircraft         respjson.Field
		PackageID              respjson.Field
		ReceiverCallSign       respjson.Field
		ReceiverCellPosition   respjson.Field
		ReceiverCoord          respjson.Field
		ReceiverDeliveryMethod respjson.Field
		ReceiverDeployedIcao   respjson.Field
		ReceiverExercise       respjson.Field
		ReceiverFuelType       respjson.Field
		ReceiverLegNum         respjson.Field
		ReceiverMds            respjson.Field
		ReceiverOwner          respjson.Field
		ReceiverPoc            respjson.Field
		RecOrg                 respjson.Field
		SequenceNum            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventListResponseReceiver) RawJSON() string { return r.JSON.raw }
func (r *AirEventListResponseReceiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this Air Event.
type AirEventListResponseRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date time.Time `json:"date" format:"date-time"`
	// Optional remark ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalRemarkID string `json:"externalRemarkId"`
	// Text of the remark.
	Text string `json:"text"`
	// User who published the remark.
	User string `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date             respjson.Field
		ExternalRemarkID respjson.Field
		Text             respjson.Field
		User             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventListResponseRemark) RawJSON() string { return r.JSON.raw }
func (r *AirEventListResponseRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of tanker aircraft associated with this Air Event.
type AirEventListResponseTanker struct {
	// Alternate mission identifier of this tanker provided by source.
	AltTankerMissionID string `json:"altTankerMissionId"`
	// The Air Mobility Command (AMC) mission identifier of this tanker.
	AmcTankerMissionID string `json:"amcTankerMissionId"`
	// Flag indicating that this tanker is flying a dual role mission in an air
	// refueling event.
	DualRole bool `json:"dualRole"`
	// Optional tanker identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalTankerID string `json:"externalTankerId"`
	// Total weight of the fuel transferred from this tanker during an air refueling
	// event, in pounds.
	FuelOff float64 `json:"fuelOff"`
	// The UDL ID of the airfield this tanker is associated with for this event.
	IDTankerAirfield string `json:"idTankerAirfield"`
	// The UDL ID of the mission this tanker is associated with for this event.
	IDTankerMission string `json:"idTankerMission"`
	// The UDL ID of the aircraft sortie this tanker is associated with for this event.
	IDTankerSortie string `json:"idTankerSortie"`
	// The call sign assigned to this tanker.
	TankerCallSign string `json:"tankerCallSign"`
	// Position of this tanker within a group of tankers in an air refueling event.
	TankerCellPosition int64 `json:"tankerCellPosition"`
	// Coordination record identifier of this tanker.
	TankerCoord string `json:"tankerCoord"`
	// Type of fuel delivery method used by the tanker during an air refueling event
	// (BOOM, DROGUE, BOTH).
	TankerDeliveryMethod string `json:"tankerDeliveryMethod"`
	// Location the tanker has been deployed to in preparation for an air refueling
	// event.
	TankerDeployedIcao string `json:"tankerDeployedICAO"`
	// Type of fuel being transferred from the tanker in an air refueling event.
	TankerFuelType string `json:"tankerFuelType"`
	// Identifies the itinerary point of a mission that this tanker is linked to.
	TankerLegNum int64 `json:"tankerLegNum"`
	// The Model Design Series designation of this tanker.
	TankerMds string `json:"tankerMDS"`
	// The wing or unit that owns this tanker.
	TankerOwner string `json:"tankerOwner"`
	// The name and/or number of the point of contact for this tanker.
	TankerPoc string `json:"tankerPOC"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltTankerMissionID   respjson.Field
		AmcTankerMissionID   respjson.Field
		DualRole             respjson.Field
		ExternalTankerID     respjson.Field
		FuelOff              respjson.Field
		IDTankerAirfield     respjson.Field
		IDTankerMission      respjson.Field
		IDTankerSortie       respjson.Field
		TankerCallSign       respjson.Field
		TankerCellPosition   respjson.Field
		TankerCoord          respjson.Field
		TankerDeliveryMethod respjson.Field
		TankerDeployedIcao   respjson.Field
		TankerFuelType       respjson.Field
		TankerLegNum         respjson.Field
		TankerMds            respjson.Field
		TankerOwner          respjson.Field
		TankerPoc            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventListResponseTanker) RawJSON() string { return r.JSON.raw }
func (r *AirEventListResponseTanker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information related to an air event (e.g. FUEL TRANSFER, AIR DROP) and the
// associated aircraft.
type AirEventGetResponse struct {
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
	DataMode AirEventGetResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of air event (e.g. FUEL TRANSFER, AIR DROP, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The actual arrival time of the aircraft at the air event, in ISO 8601 UTC format
	// with millisecond precision.
	ActualArrTime time.Time `json:"actualArrTime" format:"date-time"`
	// The actual departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	ActualDepTime time.Time `json:"actualDepTime" format:"date-time"`
	// The Air Refueling Control Time is the planned time the tanker aircraft will
	// transfer fuel to the receiver aircraft, in ISO 8601 UTC format, with millisecond
	// precision.
	Arct time.Time `json:"arct" format:"date-time"`
	// Type of process used by AMC to schedule this air refueling event. Possible
	// values are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched
	// Theater Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S
	// (Soft Air Refueling), T (Matched Theater Operation Short Notice (Theater
	// Assets)), V (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short
	// Notice (AMC Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z
	// (Other Air Refueling).
	ArEventType string `json:"arEventType"`
	// The purpose of the air event at the arrival location. Can be either descriptive
	// text such as 'fuel onload' or a purpose code specified by the provider, such as
	// 'A'.
	ArrPurpose string `json:"arrPurpose"`
	// Identifier of the air refueling track, if applicable.
	ArTrackID string `json:"arTrackId"`
	// Name of the air refueling track, if applicable.
	ArTrackName string `json:"arTrackName"`
	// Altitude of this air event, in feet.
	BaseAlt float64 `json:"baseAlt"`
	// Flag indicating that this air refueling event has been cancelled.
	Cancelled bool `json:"cancelled"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The purpose of the air event at the departure location. Can be either
	// descriptive text such as 'fuel onload' or a purpose code specified by the
	// provider, such as 'A'.
	DepPurpose string `json:"depPurpose"`
	// The current estimated arrival time of the aircraft at the air event, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime time.Time `json:"estArrTime" format:"date-time"`
	// The current estimated departure time of the aircraft from the air event, in ISO
	// 8601 UTC format with millisecond precision.
	EstDepTime time.Time `json:"estDepTime" format:"date-time"`
	// Optional air event ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalAirEventID string `json:"externalAirEventId"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalArTrackID string `json:"externalARTrackId"`
	// The UDL unique identifier of the mission associated with this air event.
	IDMission string `json:"idMission"`
	// The UDL unique identifier of the sortie associated with this air event.
	IDSortie string `json:"idSortie"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum int64 `json:"legNum"`
	// The location representing this air event specified as a feature Id. Locations
	// specified include air refueling track Ids and air drop event locations.
	Location string `json:"location"`
	// The number of tankers requested for an air refueling event.
	NumTankers int64 `json:"numTankers"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The scheduled arrival time of the aircraft at the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedArrTime time.Time `json:"plannedArrTime" format:"date-time"`
	// The scheduled departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedDepTime time.Time `json:"plannedDepTime" format:"date-time"`
	// Priority of this air event.
	Priority string `json:"priority"`
	// Collection of receiver aircraft associated with this Air Event.
	Receivers []AirEventGetResponseReceiver `json:"receivers"`
	// Collection of remarks associated with this Air Event.
	Remarks []AirEventGetResponseRemark `json:"remarks"`
	// Flag indicating if the receiver unit has requested flying an air refueling track
	// in both directions.
	RevTrack bool `json:"revTrack"`
	// The Rendezvous Control Time is the planned time the tanker and receiver aircraft
	// will rendezvous for an en route type air refueling event, in ISO 8601 UTC
	// format, with millisecond precision.
	Rzct time.Time `json:"rzct" format:"date-time"`
	// Rendezvous point for the tanker and receiver during this air refueling event.
	// Possible values are AN (Anchor Nav Point), AP (Anchor Pattern), CP (Control
	// Point), ET (Entry Point), EX (Exit Point), IP (Initial Point), NC (Nav Check
	// Point).
	RzPoint string `json:"rzPoint"`
	// Type of rendezvous used for this air refueling event. Possible values are BUD
	// (Buddy), EN (Enroute), GCI (Ground Control), PP (Point Parallel).
	RzType string `json:"rzType"`
	// Flag indicating that the receiver unit has requested flying a short portion of
	// an air refueling track.
	ShortTrack bool `json:"shortTrack"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Status of this air refueling event track reservation. Receivers are responsible
	// for scheduling or reserving air refueling tracks. Possible values are A
	// (Altitude Reservation), R (Reserved), or Q (Questionable).
	StatusCode string `json:"statusCode"`
	// Collection of tanker aircraft associated with this Air Event.
	Tankers []AirEventGetResponseTanker `json:"tankers"`
	// Length of time the receiver unit has requested for an air event, in hours.
	TrackTime float64 `json:"trackTime"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		ActualArrTime         respjson.Field
		ActualDepTime         respjson.Field
		Arct                  respjson.Field
		ArEventType           respjson.Field
		ArrPurpose            respjson.Field
		ArTrackID             respjson.Field
		ArTrackName           respjson.Field
		BaseAlt               respjson.Field
		Cancelled             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DepPurpose            respjson.Field
		EstArrTime            respjson.Field
		EstDepTime            respjson.Field
		ExternalAirEventID    respjson.Field
		ExternalArTrackID     respjson.Field
		IDMission             respjson.Field
		IDSortie              respjson.Field
		LegNum                respjson.Field
		Location              respjson.Field
		NumTankers            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PlannedArrTime        respjson.Field
		PlannedDepTime        respjson.Field
		Priority              respjson.Field
		Receivers             respjson.Field
		Remarks               respjson.Field
		RevTrack              respjson.Field
		Rzct                  respjson.Field
		RzPoint               respjson.Field
		RzType                respjson.Field
		ShortTrack            respjson.Field
		SourceDl              respjson.Field
		StatusCode            respjson.Field
		Tankers               respjson.Field
		TrackTime             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AirEventGetResponse) UnmarshalJSON(data []byte) error {
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
type AirEventGetResponseDataMode string

const (
	AirEventGetResponseDataModeReal      AirEventGetResponseDataMode = "REAL"
	AirEventGetResponseDataModeTest      AirEventGetResponseDataMode = "TEST"
	AirEventGetResponseDataModeSimulated AirEventGetResponseDataMode = "SIMULATED"
	AirEventGetResponseDataModeExercise  AirEventGetResponseDataMode = "EXERCISE"
)

// Collection of receiver aircraft associated with this Air Event.
type AirEventGetResponseReceiver struct {
	// Alternate mission identifier of this receiver provided by source.
	AltReceiverMissionID string `json:"altReceiverMissionId"`
	// The Air Mobility Command (AMC) mission identifier of this receiver.
	AmcReceiverMissionID string `json:"amcReceiverMissionId"`
	// Optional receiver identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalReceiverID string `json:"externalReceiverId"`
	// Total weight of the fuel transferred to this receiver during an air refueling
	// event, in pounds.
	FuelOn float64 `json:"fuelOn"`
	// The UDL ID of the airfield this receiver is associated with for this event.
	IDReceiverAirfield string `json:"idReceiverAirfield"`
	// The UDL ID of the mission this receiver is associated with for this event.
	IDReceiverMission string `json:"idReceiverMission"`
	// The UDL ID of the aircraft sortie this receiver is associated with for this
	// event.
	IDReceiverSortie string `json:"idReceiverSortie"`
	// Number of aircraft contained within one receiver coordination record for an air
	// refueling event.
	NumRecAircraft int64 `json:"numRecAircraft"`
	// The package identifier for the receiver in an air refueling event.
	PackageID string `json:"packageId"`
	// The call sign assigned to this receiver.
	ReceiverCallSign string `json:"receiverCallSign"`
	// Position of this receiver within a group of receivers in an air refueling event.
	ReceiverCellPosition int64 `json:"receiverCellPosition"`
	// Coordination record identifier of this receiver.
	ReceiverCoord string `json:"receiverCoord"`
	// Type of fuel delivery method used by the receiver during an air refueling event
	// (BOOM, DROGUE, BOTH).
	ReceiverDeliveryMethod string `json:"receiverDeliveryMethod"`
	// Location the receiver is deployed to for an air refueling event.
	ReceiverDeployedIcao string `json:"receiverDeployedICAO"`
	// Name of the receiver exercise associated with an air refueling event.
	ReceiverExercise string `json:"receiverExercise"`
	// Type of fuel being transferred to the receiver in an air refueling event.
	ReceiverFuelType string `json:"receiverFuelType"`
	// Identifies the itinerary point of a mission that this receiver is linked to.
	ReceiverLegNum int64 `json:"receiverLegNum"`
	// The Model Design Series designation of this receiver.
	ReceiverMds string `json:"receiverMDS"`
	// The wing or unit that owns this receiver.
	ReceiverOwner string `json:"receiverOwner"`
	// The name and/or number of the point of contact for this receiver.
	ReceiverPoc string `json:"receiverPOC"`
	// The major command level (MAJCOM) or foreign military sales (FMS) name of the
	// receiver's organization. The tanker flying hours used for an air refueling event
	// are logged against the receiver MAJCOM or foreign government being supported.
	RecOrg string `json:"recOrg"`
	// Indicates the unique number by Unit ID, which identifies an air refueling event.
	SequenceNum string `json:"sequenceNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltReceiverMissionID   respjson.Field
		AmcReceiverMissionID   respjson.Field
		ExternalReceiverID     respjson.Field
		FuelOn                 respjson.Field
		IDReceiverAirfield     respjson.Field
		IDReceiverMission      respjson.Field
		IDReceiverSortie       respjson.Field
		NumRecAircraft         respjson.Field
		PackageID              respjson.Field
		ReceiverCallSign       respjson.Field
		ReceiverCellPosition   respjson.Field
		ReceiverCoord          respjson.Field
		ReceiverDeliveryMethod respjson.Field
		ReceiverDeployedIcao   respjson.Field
		ReceiverExercise       respjson.Field
		ReceiverFuelType       respjson.Field
		ReceiverLegNum         respjson.Field
		ReceiverMds            respjson.Field
		ReceiverOwner          respjson.Field
		ReceiverPoc            respjson.Field
		RecOrg                 respjson.Field
		SequenceNum            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventGetResponseReceiver) RawJSON() string { return r.JSON.raw }
func (r *AirEventGetResponseReceiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this Air Event.
type AirEventGetResponseRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date time.Time `json:"date" format:"date-time"`
	// Optional remark ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalRemarkID string `json:"externalRemarkId"`
	// Text of the remark.
	Text string `json:"text"`
	// User who published the remark.
	User string `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date             respjson.Field
		ExternalRemarkID respjson.Field
		Text             respjson.Field
		User             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventGetResponseRemark) RawJSON() string { return r.JSON.raw }
func (r *AirEventGetResponseRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of tanker aircraft associated with this Air Event.
type AirEventGetResponseTanker struct {
	// Alternate mission identifier of this tanker provided by source.
	AltTankerMissionID string `json:"altTankerMissionId"`
	// The Air Mobility Command (AMC) mission identifier of this tanker.
	AmcTankerMissionID string `json:"amcTankerMissionId"`
	// Flag indicating that this tanker is flying a dual role mission in an air
	// refueling event.
	DualRole bool `json:"dualRole"`
	// Optional tanker identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalTankerID string `json:"externalTankerId"`
	// Total weight of the fuel transferred from this tanker during an air refueling
	// event, in pounds.
	FuelOff float64 `json:"fuelOff"`
	// The UDL ID of the airfield this tanker is associated with for this event.
	IDTankerAirfield string `json:"idTankerAirfield"`
	// The UDL ID of the mission this tanker is associated with for this event.
	IDTankerMission string `json:"idTankerMission"`
	// The UDL ID of the aircraft sortie this tanker is associated with for this event.
	IDTankerSortie string `json:"idTankerSortie"`
	// The call sign assigned to this tanker.
	TankerCallSign string `json:"tankerCallSign"`
	// Position of this tanker within a group of tankers in an air refueling event.
	TankerCellPosition int64 `json:"tankerCellPosition"`
	// Coordination record identifier of this tanker.
	TankerCoord string `json:"tankerCoord"`
	// Type of fuel delivery method used by the tanker during an air refueling event
	// (BOOM, DROGUE, BOTH).
	TankerDeliveryMethod string `json:"tankerDeliveryMethod"`
	// Location the tanker has been deployed to in preparation for an air refueling
	// event.
	TankerDeployedIcao string `json:"tankerDeployedICAO"`
	// Type of fuel being transferred from the tanker in an air refueling event.
	TankerFuelType string `json:"tankerFuelType"`
	// Identifies the itinerary point of a mission that this tanker is linked to.
	TankerLegNum int64 `json:"tankerLegNum"`
	// The Model Design Series designation of this tanker.
	TankerMds string `json:"tankerMDS"`
	// The wing or unit that owns this tanker.
	TankerOwner string `json:"tankerOwner"`
	// The name and/or number of the point of contact for this tanker.
	TankerPoc string `json:"tankerPOC"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltTankerMissionID   respjson.Field
		AmcTankerMissionID   respjson.Field
		DualRole             respjson.Field
		ExternalTankerID     respjson.Field
		FuelOff              respjson.Field
		IDTankerAirfield     respjson.Field
		IDTankerMission      respjson.Field
		IDTankerSortie       respjson.Field
		TankerCallSign       respjson.Field
		TankerCellPosition   respjson.Field
		TankerCoord          respjson.Field
		TankerDeliveryMethod respjson.Field
		TankerDeployedIcao   respjson.Field
		TankerFuelType       respjson.Field
		TankerLegNum         respjson.Field
		TankerMds            respjson.Field
		TankerOwner          respjson.Field
		TankerPoc            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventGetResponseTanker) RawJSON() string { return r.JSON.raw }
func (r *AirEventGetResponseTanker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirEventQueryhelpResponse struct {
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
func (r AirEventQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *AirEventQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information related to an air event (e.g. FUEL TRANSFER, AIR DROP) and the
// associated aircraft.
type AirEventTupleResponse struct {
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
	DataMode AirEventTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of air event (e.g. FUEL TRANSFER, AIR DROP, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The actual arrival time of the aircraft at the air event, in ISO 8601 UTC format
	// with millisecond precision.
	ActualArrTime time.Time `json:"actualArrTime" format:"date-time"`
	// The actual departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	ActualDepTime time.Time `json:"actualDepTime" format:"date-time"`
	// The Air Refueling Control Time is the planned time the tanker aircraft will
	// transfer fuel to the receiver aircraft, in ISO 8601 UTC format, with millisecond
	// precision.
	Arct time.Time `json:"arct" format:"date-time"`
	// Type of process used by AMC to schedule this air refueling event. Possible
	// values are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched
	// Theater Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S
	// (Soft Air Refueling), T (Matched Theater Operation Short Notice (Theater
	// Assets)), V (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short
	// Notice (AMC Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z
	// (Other Air Refueling).
	ArEventType string `json:"arEventType"`
	// The purpose of the air event at the arrival location. Can be either descriptive
	// text such as 'fuel onload' or a purpose code specified by the provider, such as
	// 'A'.
	ArrPurpose string `json:"arrPurpose"`
	// Identifier of the air refueling track, if applicable.
	ArTrackID string `json:"arTrackId"`
	// Name of the air refueling track, if applicable.
	ArTrackName string `json:"arTrackName"`
	// Altitude of this air event, in feet.
	BaseAlt float64 `json:"baseAlt"`
	// Flag indicating that this air refueling event has been cancelled.
	Cancelled bool `json:"cancelled"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The purpose of the air event at the departure location. Can be either
	// descriptive text such as 'fuel onload' or a purpose code specified by the
	// provider, such as 'A'.
	DepPurpose string `json:"depPurpose"`
	// The current estimated arrival time of the aircraft at the air event, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime time.Time `json:"estArrTime" format:"date-time"`
	// The current estimated departure time of the aircraft from the air event, in ISO
	// 8601 UTC format with millisecond precision.
	EstDepTime time.Time `json:"estDepTime" format:"date-time"`
	// Optional air event ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalAirEventID string `json:"externalAirEventId"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalArTrackID string `json:"externalARTrackId"`
	// The UDL unique identifier of the mission associated with this air event.
	IDMission string `json:"idMission"`
	// The UDL unique identifier of the sortie associated with this air event.
	IDSortie string `json:"idSortie"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum int64 `json:"legNum"`
	// The location representing this air event specified as a feature Id. Locations
	// specified include air refueling track Ids and air drop event locations.
	Location string `json:"location"`
	// The number of tankers requested for an air refueling event.
	NumTankers int64 `json:"numTankers"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The scheduled arrival time of the aircraft at the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedArrTime time.Time `json:"plannedArrTime" format:"date-time"`
	// The scheduled departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedDepTime time.Time `json:"plannedDepTime" format:"date-time"`
	// Priority of this air event.
	Priority string `json:"priority"`
	// Collection of receiver aircraft associated with this Air Event.
	Receivers []AirEventTupleResponseReceiver `json:"receivers"`
	// Collection of remarks associated with this Air Event.
	Remarks []AirEventTupleResponseRemark `json:"remarks"`
	// Flag indicating if the receiver unit has requested flying an air refueling track
	// in both directions.
	RevTrack bool `json:"revTrack"`
	// The Rendezvous Control Time is the planned time the tanker and receiver aircraft
	// will rendezvous for an en route type air refueling event, in ISO 8601 UTC
	// format, with millisecond precision.
	Rzct time.Time `json:"rzct" format:"date-time"`
	// Rendezvous point for the tanker and receiver during this air refueling event.
	// Possible values are AN (Anchor Nav Point), AP (Anchor Pattern), CP (Control
	// Point), ET (Entry Point), EX (Exit Point), IP (Initial Point), NC (Nav Check
	// Point).
	RzPoint string `json:"rzPoint"`
	// Type of rendezvous used for this air refueling event. Possible values are BUD
	// (Buddy), EN (Enroute), GCI (Ground Control), PP (Point Parallel).
	RzType string `json:"rzType"`
	// Flag indicating that the receiver unit has requested flying a short portion of
	// an air refueling track.
	ShortTrack bool `json:"shortTrack"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Status of this air refueling event track reservation. Receivers are responsible
	// for scheduling or reserving air refueling tracks. Possible values are A
	// (Altitude Reservation), R (Reserved), or Q (Questionable).
	StatusCode string `json:"statusCode"`
	// Collection of tanker aircraft associated with this Air Event.
	Tankers []AirEventTupleResponseTanker `json:"tankers"`
	// Length of time the receiver unit has requested for an air event, in hours.
	TrackTime float64 `json:"trackTime"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		ActualArrTime         respjson.Field
		ActualDepTime         respjson.Field
		Arct                  respjson.Field
		ArEventType           respjson.Field
		ArrPurpose            respjson.Field
		ArTrackID             respjson.Field
		ArTrackName           respjson.Field
		BaseAlt               respjson.Field
		Cancelled             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DepPurpose            respjson.Field
		EstArrTime            respjson.Field
		EstDepTime            respjson.Field
		ExternalAirEventID    respjson.Field
		ExternalArTrackID     respjson.Field
		IDMission             respjson.Field
		IDSortie              respjson.Field
		LegNum                respjson.Field
		Location              respjson.Field
		NumTankers            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PlannedArrTime        respjson.Field
		PlannedDepTime        respjson.Field
		Priority              respjson.Field
		Receivers             respjson.Field
		Remarks               respjson.Field
		RevTrack              respjson.Field
		Rzct                  respjson.Field
		RzPoint               respjson.Field
		RzType                respjson.Field
		ShortTrack            respjson.Field
		SourceDl              respjson.Field
		StatusCode            respjson.Field
		Tankers               respjson.Field
		TrackTime             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *AirEventTupleResponse) UnmarshalJSON(data []byte) error {
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
type AirEventTupleResponseDataMode string

const (
	AirEventTupleResponseDataModeReal      AirEventTupleResponseDataMode = "REAL"
	AirEventTupleResponseDataModeTest      AirEventTupleResponseDataMode = "TEST"
	AirEventTupleResponseDataModeSimulated AirEventTupleResponseDataMode = "SIMULATED"
	AirEventTupleResponseDataModeExercise  AirEventTupleResponseDataMode = "EXERCISE"
)

// Collection of receiver aircraft associated with this Air Event.
type AirEventTupleResponseReceiver struct {
	// Alternate mission identifier of this receiver provided by source.
	AltReceiverMissionID string `json:"altReceiverMissionId"`
	// The Air Mobility Command (AMC) mission identifier of this receiver.
	AmcReceiverMissionID string `json:"amcReceiverMissionId"`
	// Optional receiver identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalReceiverID string `json:"externalReceiverId"`
	// Total weight of the fuel transferred to this receiver during an air refueling
	// event, in pounds.
	FuelOn float64 `json:"fuelOn"`
	// The UDL ID of the airfield this receiver is associated with for this event.
	IDReceiverAirfield string `json:"idReceiverAirfield"`
	// The UDL ID of the mission this receiver is associated with for this event.
	IDReceiverMission string `json:"idReceiverMission"`
	// The UDL ID of the aircraft sortie this receiver is associated with for this
	// event.
	IDReceiverSortie string `json:"idReceiverSortie"`
	// Number of aircraft contained within one receiver coordination record for an air
	// refueling event.
	NumRecAircraft int64 `json:"numRecAircraft"`
	// The package identifier for the receiver in an air refueling event.
	PackageID string `json:"packageId"`
	// The call sign assigned to this receiver.
	ReceiverCallSign string `json:"receiverCallSign"`
	// Position of this receiver within a group of receivers in an air refueling event.
	ReceiverCellPosition int64 `json:"receiverCellPosition"`
	// Coordination record identifier of this receiver.
	ReceiverCoord string `json:"receiverCoord"`
	// Type of fuel delivery method used by the receiver during an air refueling event
	// (BOOM, DROGUE, BOTH).
	ReceiverDeliveryMethod string `json:"receiverDeliveryMethod"`
	// Location the receiver is deployed to for an air refueling event.
	ReceiverDeployedIcao string `json:"receiverDeployedICAO"`
	// Name of the receiver exercise associated with an air refueling event.
	ReceiverExercise string `json:"receiverExercise"`
	// Type of fuel being transferred to the receiver in an air refueling event.
	ReceiverFuelType string `json:"receiverFuelType"`
	// Identifies the itinerary point of a mission that this receiver is linked to.
	ReceiverLegNum int64 `json:"receiverLegNum"`
	// The Model Design Series designation of this receiver.
	ReceiverMds string `json:"receiverMDS"`
	// The wing or unit that owns this receiver.
	ReceiverOwner string `json:"receiverOwner"`
	// The name and/or number of the point of contact for this receiver.
	ReceiverPoc string `json:"receiverPOC"`
	// The major command level (MAJCOM) or foreign military sales (FMS) name of the
	// receiver's organization. The tanker flying hours used for an air refueling event
	// are logged against the receiver MAJCOM or foreign government being supported.
	RecOrg string `json:"recOrg"`
	// Indicates the unique number by Unit ID, which identifies an air refueling event.
	SequenceNum string `json:"sequenceNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltReceiverMissionID   respjson.Field
		AmcReceiverMissionID   respjson.Field
		ExternalReceiverID     respjson.Field
		FuelOn                 respjson.Field
		IDReceiverAirfield     respjson.Field
		IDReceiverMission      respjson.Field
		IDReceiverSortie       respjson.Field
		NumRecAircraft         respjson.Field
		PackageID              respjson.Field
		ReceiverCallSign       respjson.Field
		ReceiverCellPosition   respjson.Field
		ReceiverCoord          respjson.Field
		ReceiverDeliveryMethod respjson.Field
		ReceiverDeployedIcao   respjson.Field
		ReceiverExercise       respjson.Field
		ReceiverFuelType       respjson.Field
		ReceiverLegNum         respjson.Field
		ReceiverMds            respjson.Field
		ReceiverOwner          respjson.Field
		ReceiverPoc            respjson.Field
		RecOrg                 respjson.Field
		SequenceNum            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventTupleResponseReceiver) RawJSON() string { return r.JSON.raw }
func (r *AirEventTupleResponseReceiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this Air Event.
type AirEventTupleResponseRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date time.Time `json:"date" format:"date-time"`
	// Optional remark ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalRemarkID string `json:"externalRemarkId"`
	// Text of the remark.
	Text string `json:"text"`
	// User who published the remark.
	User string `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date             respjson.Field
		ExternalRemarkID respjson.Field
		Text             respjson.Field
		User             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventTupleResponseRemark) RawJSON() string { return r.JSON.raw }
func (r *AirEventTupleResponseRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of tanker aircraft associated with this Air Event.
type AirEventTupleResponseTanker struct {
	// Alternate mission identifier of this tanker provided by source.
	AltTankerMissionID string `json:"altTankerMissionId"`
	// The Air Mobility Command (AMC) mission identifier of this tanker.
	AmcTankerMissionID string `json:"amcTankerMissionId"`
	// Flag indicating that this tanker is flying a dual role mission in an air
	// refueling event.
	DualRole bool `json:"dualRole"`
	// Optional tanker identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalTankerID string `json:"externalTankerId"`
	// Total weight of the fuel transferred from this tanker during an air refueling
	// event, in pounds.
	FuelOff float64 `json:"fuelOff"`
	// The UDL ID of the airfield this tanker is associated with for this event.
	IDTankerAirfield string `json:"idTankerAirfield"`
	// The UDL ID of the mission this tanker is associated with for this event.
	IDTankerMission string `json:"idTankerMission"`
	// The UDL ID of the aircraft sortie this tanker is associated with for this event.
	IDTankerSortie string `json:"idTankerSortie"`
	// The call sign assigned to this tanker.
	TankerCallSign string `json:"tankerCallSign"`
	// Position of this tanker within a group of tankers in an air refueling event.
	TankerCellPosition int64 `json:"tankerCellPosition"`
	// Coordination record identifier of this tanker.
	TankerCoord string `json:"tankerCoord"`
	// Type of fuel delivery method used by the tanker during an air refueling event
	// (BOOM, DROGUE, BOTH).
	TankerDeliveryMethod string `json:"tankerDeliveryMethod"`
	// Location the tanker has been deployed to in preparation for an air refueling
	// event.
	TankerDeployedIcao string `json:"tankerDeployedICAO"`
	// Type of fuel being transferred from the tanker in an air refueling event.
	TankerFuelType string `json:"tankerFuelType"`
	// Identifies the itinerary point of a mission that this tanker is linked to.
	TankerLegNum int64 `json:"tankerLegNum"`
	// The Model Design Series designation of this tanker.
	TankerMds string `json:"tankerMDS"`
	// The wing or unit that owns this tanker.
	TankerOwner string `json:"tankerOwner"`
	// The name and/or number of the point of contact for this tanker.
	TankerPoc string `json:"tankerPOC"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltTankerMissionID   respjson.Field
		AmcTankerMissionID   respjson.Field
		DualRole             respjson.Field
		ExternalTankerID     respjson.Field
		FuelOff              respjson.Field
		IDTankerAirfield     respjson.Field
		IDTankerMission      respjson.Field
		IDTankerSortie       respjson.Field
		TankerCallSign       respjson.Field
		TankerCellPosition   respjson.Field
		TankerCoord          respjson.Field
		TankerDeliveryMethod respjson.Field
		TankerDeployedIcao   respjson.Field
		TankerFuelType       respjson.Field
		TankerLegNum         respjson.Field
		TankerMds            respjson.Field
		TankerOwner          respjson.Field
		TankerPoc            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirEventTupleResponseTanker) RawJSON() string { return r.JSON.raw }
func (r *AirEventTupleResponseTanker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirEventNewParams struct {
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
	DataMode AirEventNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of air event (e.g. FUEL TRANSFER, AIR DROP, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The actual arrival time of the aircraft at the air event, in ISO 8601 UTC format
	// with millisecond precision.
	ActualArrTime param.Opt[time.Time] `json:"actualArrTime,omitzero" format:"date-time"`
	// The actual departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	ActualDepTime param.Opt[time.Time] `json:"actualDepTime,omitzero" format:"date-time"`
	// The Air Refueling Control Time is the planned time the tanker aircraft will
	// transfer fuel to the receiver aircraft, in ISO 8601 UTC format, with millisecond
	// precision.
	Arct param.Opt[time.Time] `json:"arct,omitzero" format:"date-time"`
	// Type of process used by AMC to schedule this air refueling event. Possible
	// values are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched
	// Theater Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S
	// (Soft Air Refueling), T (Matched Theater Operation Short Notice (Theater
	// Assets)), V (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short
	// Notice (AMC Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z
	// (Other Air Refueling).
	ArEventType param.Opt[string] `json:"arEventType,omitzero"`
	// The purpose of the air event at the arrival location. Can be either descriptive
	// text such as 'fuel onload' or a purpose code specified by the provider, such as
	// 'A'.
	ArrPurpose param.Opt[string] `json:"arrPurpose,omitzero"`
	// Identifier of the air refueling track, if applicable.
	ArTrackID param.Opt[string] `json:"arTrackId,omitzero"`
	// Name of the air refueling track, if applicable.
	ArTrackName param.Opt[string] `json:"arTrackName,omitzero"`
	// Altitude of this air event, in feet.
	BaseAlt param.Opt[float64] `json:"baseAlt,omitzero"`
	// Flag indicating that this air refueling event has been cancelled.
	Cancelled param.Opt[bool] `json:"cancelled,omitzero"`
	// The purpose of the air event at the departure location. Can be either
	// descriptive text such as 'fuel onload' or a purpose code specified by the
	// provider, such as 'A'.
	DepPurpose param.Opt[string] `json:"depPurpose,omitzero"`
	// The current estimated arrival time of the aircraft at the air event, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// The current estimated departure time of the aircraft from the air event, in ISO
	// 8601 UTC format with millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// Optional air event ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalAirEventID param.Opt[string] `json:"externalAirEventId,omitzero"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalArTrackID param.Opt[string] `json:"externalARTrackId,omitzero"`
	// The UDL unique identifier of the mission associated with this air event.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// The UDL unique identifier of the sortie associated with this air event.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The location representing this air event specified as a feature Id. Locations
	// specified include air refueling track Ids and air drop event locations.
	Location param.Opt[string] `json:"location,omitzero"`
	// The number of tankers requested for an air refueling event.
	NumTankers param.Opt[int64] `json:"numTankers,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The scheduled arrival time of the aircraft at the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedArrTime param.Opt[time.Time] `json:"plannedArrTime,omitzero" format:"date-time"`
	// The scheduled departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedDepTime param.Opt[time.Time] `json:"plannedDepTime,omitzero" format:"date-time"`
	// Priority of this air event.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// Flag indicating if the receiver unit has requested flying an air refueling track
	// in both directions.
	RevTrack param.Opt[bool] `json:"revTrack,omitzero"`
	// The Rendezvous Control Time is the planned time the tanker and receiver aircraft
	// will rendezvous for an en route type air refueling event, in ISO 8601 UTC
	// format, with millisecond precision.
	Rzct param.Opt[time.Time] `json:"rzct,omitzero" format:"date-time"`
	// Rendezvous point for the tanker and receiver during this air refueling event.
	// Possible values are AN (Anchor Nav Point), AP (Anchor Pattern), CP (Control
	// Point), ET (Entry Point), EX (Exit Point), IP (Initial Point), NC (Nav Check
	// Point).
	RzPoint param.Opt[string] `json:"rzPoint,omitzero"`
	// Type of rendezvous used for this air refueling event. Possible values are BUD
	// (Buddy), EN (Enroute), GCI (Ground Control), PP (Point Parallel).
	RzType param.Opt[string] `json:"rzType,omitzero"`
	// Flag indicating that the receiver unit has requested flying a short portion of
	// an air refueling track.
	ShortTrack param.Opt[bool] `json:"shortTrack,omitzero"`
	// Status of this air refueling event track reservation. Receivers are responsible
	// for scheduling or reserving air refueling tracks. Possible values are A
	// (Altitude Reservation), R (Reserved), or Q (Questionable).
	StatusCode param.Opt[string] `json:"statusCode,omitzero"`
	// Length of time the receiver unit has requested for an air event, in hours.
	TrackTime param.Opt[float64] `json:"trackTime,omitzero"`
	// Collection of receiver aircraft associated with this Air Event.
	Receivers []AirEventNewParamsReceiver `json:"receivers,omitzero"`
	// Collection of remarks associated with this Air Event.
	Remarks []AirEventNewParamsRemark `json:"remarks,omitzero"`
	// Collection of tanker aircraft associated with this Air Event.
	Tankers []AirEventNewParamsTanker `json:"tankers,omitzero"`
	paramObj
}

func (r AirEventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AirEventNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventNewParams) UnmarshalJSON(data []byte) error {
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
type AirEventNewParamsDataMode string

const (
	AirEventNewParamsDataModeReal      AirEventNewParamsDataMode = "REAL"
	AirEventNewParamsDataModeTest      AirEventNewParamsDataMode = "TEST"
	AirEventNewParamsDataModeSimulated AirEventNewParamsDataMode = "SIMULATED"
	AirEventNewParamsDataModeExercise  AirEventNewParamsDataMode = "EXERCISE"
)

// Collection of receiver aircraft associated with this Air Event.
type AirEventNewParamsReceiver struct {
	// Alternate mission identifier of this receiver provided by source.
	AltReceiverMissionID param.Opt[string] `json:"altReceiverMissionId,omitzero"`
	// The Air Mobility Command (AMC) mission identifier of this receiver.
	AmcReceiverMissionID param.Opt[string] `json:"amcReceiverMissionId,omitzero"`
	// Optional receiver identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalReceiverID param.Opt[string] `json:"externalReceiverId,omitzero"`
	// Total weight of the fuel transferred to this receiver during an air refueling
	// event, in pounds.
	FuelOn param.Opt[float64] `json:"fuelOn,omitzero"`
	// The UDL ID of the airfield this receiver is associated with for this event.
	IDReceiverAirfield param.Opt[string] `json:"idReceiverAirfield,omitzero"`
	// The UDL ID of the mission this receiver is associated with for this event.
	IDReceiverMission param.Opt[string] `json:"idReceiverMission,omitzero"`
	// The UDL ID of the aircraft sortie this receiver is associated with for this
	// event.
	IDReceiverSortie param.Opt[string] `json:"idReceiverSortie,omitzero"`
	// Number of aircraft contained within one receiver coordination record for an air
	// refueling event.
	NumRecAircraft param.Opt[int64] `json:"numRecAircraft,omitzero"`
	// The package identifier for the receiver in an air refueling event.
	PackageID param.Opt[string] `json:"packageId,omitzero"`
	// The call sign assigned to this receiver.
	ReceiverCallSign param.Opt[string] `json:"receiverCallSign,omitzero"`
	// Position of this receiver within a group of receivers in an air refueling event.
	ReceiverCellPosition param.Opt[int64] `json:"receiverCellPosition,omitzero"`
	// Coordination record identifier of this receiver.
	ReceiverCoord param.Opt[string] `json:"receiverCoord,omitzero"`
	// Type of fuel delivery method used by the receiver during an air refueling event
	// (BOOM, DROGUE, BOTH).
	ReceiverDeliveryMethod param.Opt[string] `json:"receiverDeliveryMethod,omitzero"`
	// Location the receiver is deployed to for an air refueling event.
	ReceiverDeployedIcao param.Opt[string] `json:"receiverDeployedICAO,omitzero"`
	// Name of the receiver exercise associated with an air refueling event.
	ReceiverExercise param.Opt[string] `json:"receiverExercise,omitzero"`
	// Type of fuel being transferred to the receiver in an air refueling event.
	ReceiverFuelType param.Opt[string] `json:"receiverFuelType,omitzero"`
	// Identifies the itinerary point of a mission that this receiver is linked to.
	ReceiverLegNum param.Opt[int64] `json:"receiverLegNum,omitzero"`
	// The Model Design Series designation of this receiver.
	ReceiverMds param.Opt[string] `json:"receiverMDS,omitzero"`
	// The wing or unit that owns this receiver.
	ReceiverOwner param.Opt[string] `json:"receiverOwner,omitzero"`
	// The name and/or number of the point of contact for this receiver.
	ReceiverPoc param.Opt[string] `json:"receiverPOC,omitzero"`
	// The major command level (MAJCOM) or foreign military sales (FMS) name of the
	// receiver's organization. The tanker flying hours used for an air refueling event
	// are logged against the receiver MAJCOM or foreign government being supported.
	RecOrg param.Opt[string] `json:"recOrg,omitzero"`
	// Indicates the unique number by Unit ID, which identifies an air refueling event.
	SequenceNum param.Opt[string] `json:"sequenceNum,omitzero"`
	paramObj
}

func (r AirEventNewParamsReceiver) MarshalJSON() (data []byte, err error) {
	type shadow AirEventNewParamsReceiver
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventNewParamsReceiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this Air Event.
type AirEventNewParamsRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Optional remark ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalRemarkID param.Opt[string] `json:"externalRemarkId,omitzero"`
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// User who published the remark.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r AirEventNewParamsRemark) MarshalJSON() (data []byte, err error) {
	type shadow AirEventNewParamsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventNewParamsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of tanker aircraft associated with this Air Event.
type AirEventNewParamsTanker struct {
	// Alternate mission identifier of this tanker provided by source.
	AltTankerMissionID param.Opt[string] `json:"altTankerMissionId,omitzero"`
	// The Air Mobility Command (AMC) mission identifier of this tanker.
	AmcTankerMissionID param.Opt[string] `json:"amcTankerMissionId,omitzero"`
	// Flag indicating that this tanker is flying a dual role mission in an air
	// refueling event.
	DualRole param.Opt[bool] `json:"dualRole,omitzero"`
	// Optional tanker identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalTankerID param.Opt[string] `json:"externalTankerId,omitzero"`
	// Total weight of the fuel transferred from this tanker during an air refueling
	// event, in pounds.
	FuelOff param.Opt[float64] `json:"fuelOff,omitzero"`
	// The UDL ID of the airfield this tanker is associated with for this event.
	IDTankerAirfield param.Opt[string] `json:"idTankerAirfield,omitzero"`
	// The UDL ID of the mission this tanker is associated with for this event.
	IDTankerMission param.Opt[string] `json:"idTankerMission,omitzero"`
	// The UDL ID of the aircraft sortie this tanker is associated with for this event.
	IDTankerSortie param.Opt[string] `json:"idTankerSortie,omitzero"`
	// The call sign assigned to this tanker.
	TankerCallSign param.Opt[string] `json:"tankerCallSign,omitzero"`
	// Position of this tanker within a group of tankers in an air refueling event.
	TankerCellPosition param.Opt[int64] `json:"tankerCellPosition,omitzero"`
	// Coordination record identifier of this tanker.
	TankerCoord param.Opt[string] `json:"tankerCoord,omitzero"`
	// Type of fuel delivery method used by the tanker during an air refueling event
	// (BOOM, DROGUE, BOTH).
	TankerDeliveryMethod param.Opt[string] `json:"tankerDeliveryMethod,omitzero"`
	// Location the tanker has been deployed to in preparation for an air refueling
	// event.
	TankerDeployedIcao param.Opt[string] `json:"tankerDeployedICAO,omitzero"`
	// Type of fuel being transferred from the tanker in an air refueling event.
	TankerFuelType param.Opt[string] `json:"tankerFuelType,omitzero"`
	// Identifies the itinerary point of a mission that this tanker is linked to.
	TankerLegNum param.Opt[int64] `json:"tankerLegNum,omitzero"`
	// The Model Design Series designation of this tanker.
	TankerMds param.Opt[string] `json:"tankerMDS,omitzero"`
	// The wing or unit that owns this tanker.
	TankerOwner param.Opt[string] `json:"tankerOwner,omitzero"`
	// The name and/or number of the point of contact for this tanker.
	TankerPoc param.Opt[string] `json:"tankerPOC,omitzero"`
	paramObj
}

func (r AirEventNewParamsTanker) MarshalJSON() (data []byte, err error) {
	type shadow AirEventNewParamsTanker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventNewParamsTanker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirEventUpdateParams struct {
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
	DataMode AirEventUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of air event (e.g. FUEL TRANSFER, AIR DROP, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The actual arrival time of the aircraft at the air event, in ISO 8601 UTC format
	// with millisecond precision.
	ActualArrTime param.Opt[time.Time] `json:"actualArrTime,omitzero" format:"date-time"`
	// The actual departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	ActualDepTime param.Opt[time.Time] `json:"actualDepTime,omitzero" format:"date-time"`
	// The Air Refueling Control Time is the planned time the tanker aircraft will
	// transfer fuel to the receiver aircraft, in ISO 8601 UTC format, with millisecond
	// precision.
	Arct param.Opt[time.Time] `json:"arct,omitzero" format:"date-time"`
	// Type of process used by AMC to schedule this air refueling event. Possible
	// values are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched
	// Theater Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S
	// (Soft Air Refueling), T (Matched Theater Operation Short Notice (Theater
	// Assets)), V (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short
	// Notice (AMC Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z
	// (Other Air Refueling).
	ArEventType param.Opt[string] `json:"arEventType,omitzero"`
	// The purpose of the air event at the arrival location. Can be either descriptive
	// text such as 'fuel onload' or a purpose code specified by the provider, such as
	// 'A'.
	ArrPurpose param.Opt[string] `json:"arrPurpose,omitzero"`
	// Identifier of the air refueling track, if applicable.
	ArTrackID param.Opt[string] `json:"arTrackId,omitzero"`
	// Name of the air refueling track, if applicable.
	ArTrackName param.Opt[string] `json:"arTrackName,omitzero"`
	// Altitude of this air event, in feet.
	BaseAlt param.Opt[float64] `json:"baseAlt,omitzero"`
	// Flag indicating that this air refueling event has been cancelled.
	Cancelled param.Opt[bool] `json:"cancelled,omitzero"`
	// The purpose of the air event at the departure location. Can be either
	// descriptive text such as 'fuel onload' or a purpose code specified by the
	// provider, such as 'A'.
	DepPurpose param.Opt[string] `json:"depPurpose,omitzero"`
	// The current estimated arrival time of the aircraft at the air event, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// The current estimated departure time of the aircraft from the air event, in ISO
	// 8601 UTC format with millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// Optional air event ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalAirEventID param.Opt[string] `json:"externalAirEventId,omitzero"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalArTrackID param.Opt[string] `json:"externalARTrackId,omitzero"`
	// The UDL unique identifier of the mission associated with this air event.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// The UDL unique identifier of the sortie associated with this air event.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The location representing this air event specified as a feature Id. Locations
	// specified include air refueling track Ids and air drop event locations.
	Location param.Opt[string] `json:"location,omitzero"`
	// The number of tankers requested for an air refueling event.
	NumTankers param.Opt[int64] `json:"numTankers,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The scheduled arrival time of the aircraft at the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedArrTime param.Opt[time.Time] `json:"plannedArrTime,omitzero" format:"date-time"`
	// The scheduled departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedDepTime param.Opt[time.Time] `json:"plannedDepTime,omitzero" format:"date-time"`
	// Priority of this air event.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// Flag indicating if the receiver unit has requested flying an air refueling track
	// in both directions.
	RevTrack param.Opt[bool] `json:"revTrack,omitzero"`
	// The Rendezvous Control Time is the planned time the tanker and receiver aircraft
	// will rendezvous for an en route type air refueling event, in ISO 8601 UTC
	// format, with millisecond precision.
	Rzct param.Opt[time.Time] `json:"rzct,omitzero" format:"date-time"`
	// Rendezvous point for the tanker and receiver during this air refueling event.
	// Possible values are AN (Anchor Nav Point), AP (Anchor Pattern), CP (Control
	// Point), ET (Entry Point), EX (Exit Point), IP (Initial Point), NC (Nav Check
	// Point).
	RzPoint param.Opt[string] `json:"rzPoint,omitzero"`
	// Type of rendezvous used for this air refueling event. Possible values are BUD
	// (Buddy), EN (Enroute), GCI (Ground Control), PP (Point Parallel).
	RzType param.Opt[string] `json:"rzType,omitzero"`
	// Flag indicating that the receiver unit has requested flying a short portion of
	// an air refueling track.
	ShortTrack param.Opt[bool] `json:"shortTrack,omitzero"`
	// Status of this air refueling event track reservation. Receivers are responsible
	// for scheduling or reserving air refueling tracks. Possible values are A
	// (Altitude Reservation), R (Reserved), or Q (Questionable).
	StatusCode param.Opt[string] `json:"statusCode,omitzero"`
	// Length of time the receiver unit has requested for an air event, in hours.
	TrackTime param.Opt[float64] `json:"trackTime,omitzero"`
	// Collection of receiver aircraft associated with this Air Event.
	Receivers []AirEventUpdateParamsReceiver `json:"receivers,omitzero"`
	// Collection of remarks associated with this Air Event.
	Remarks []AirEventUpdateParamsRemark `json:"remarks,omitzero"`
	// Collection of tanker aircraft associated with this Air Event.
	Tankers []AirEventUpdateParamsTanker `json:"tankers,omitzero"`
	paramObj
}

func (r AirEventUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AirEventUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventUpdateParams) UnmarshalJSON(data []byte) error {
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
type AirEventUpdateParamsDataMode string

const (
	AirEventUpdateParamsDataModeReal      AirEventUpdateParamsDataMode = "REAL"
	AirEventUpdateParamsDataModeTest      AirEventUpdateParamsDataMode = "TEST"
	AirEventUpdateParamsDataModeSimulated AirEventUpdateParamsDataMode = "SIMULATED"
	AirEventUpdateParamsDataModeExercise  AirEventUpdateParamsDataMode = "EXERCISE"
)

// Collection of receiver aircraft associated with this Air Event.
type AirEventUpdateParamsReceiver struct {
	// Alternate mission identifier of this receiver provided by source.
	AltReceiverMissionID param.Opt[string] `json:"altReceiverMissionId,omitzero"`
	// The Air Mobility Command (AMC) mission identifier of this receiver.
	AmcReceiverMissionID param.Opt[string] `json:"amcReceiverMissionId,omitzero"`
	// Optional receiver identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalReceiverID param.Opt[string] `json:"externalReceiverId,omitzero"`
	// Total weight of the fuel transferred to this receiver during an air refueling
	// event, in pounds.
	FuelOn param.Opt[float64] `json:"fuelOn,omitzero"`
	// The UDL ID of the airfield this receiver is associated with for this event.
	IDReceiverAirfield param.Opt[string] `json:"idReceiverAirfield,omitzero"`
	// The UDL ID of the mission this receiver is associated with for this event.
	IDReceiverMission param.Opt[string] `json:"idReceiverMission,omitzero"`
	// The UDL ID of the aircraft sortie this receiver is associated with for this
	// event.
	IDReceiverSortie param.Opt[string] `json:"idReceiverSortie,omitzero"`
	// Number of aircraft contained within one receiver coordination record for an air
	// refueling event.
	NumRecAircraft param.Opt[int64] `json:"numRecAircraft,omitzero"`
	// The package identifier for the receiver in an air refueling event.
	PackageID param.Opt[string] `json:"packageId,omitzero"`
	// The call sign assigned to this receiver.
	ReceiverCallSign param.Opt[string] `json:"receiverCallSign,omitzero"`
	// Position of this receiver within a group of receivers in an air refueling event.
	ReceiverCellPosition param.Opt[int64] `json:"receiverCellPosition,omitzero"`
	// Coordination record identifier of this receiver.
	ReceiverCoord param.Opt[string] `json:"receiverCoord,omitzero"`
	// Type of fuel delivery method used by the receiver during an air refueling event
	// (BOOM, DROGUE, BOTH).
	ReceiverDeliveryMethod param.Opt[string] `json:"receiverDeliveryMethod,omitzero"`
	// Location the receiver is deployed to for an air refueling event.
	ReceiverDeployedIcao param.Opt[string] `json:"receiverDeployedICAO,omitzero"`
	// Name of the receiver exercise associated with an air refueling event.
	ReceiverExercise param.Opt[string] `json:"receiverExercise,omitzero"`
	// Type of fuel being transferred to the receiver in an air refueling event.
	ReceiverFuelType param.Opt[string] `json:"receiverFuelType,omitzero"`
	// Identifies the itinerary point of a mission that this receiver is linked to.
	ReceiverLegNum param.Opt[int64] `json:"receiverLegNum,omitzero"`
	// The Model Design Series designation of this receiver.
	ReceiverMds param.Opt[string] `json:"receiverMDS,omitzero"`
	// The wing or unit that owns this receiver.
	ReceiverOwner param.Opt[string] `json:"receiverOwner,omitzero"`
	// The name and/or number of the point of contact for this receiver.
	ReceiverPoc param.Opt[string] `json:"receiverPOC,omitzero"`
	// The major command level (MAJCOM) or foreign military sales (FMS) name of the
	// receiver's organization. The tanker flying hours used for an air refueling event
	// are logged against the receiver MAJCOM or foreign government being supported.
	RecOrg param.Opt[string] `json:"recOrg,omitzero"`
	// Indicates the unique number by Unit ID, which identifies an air refueling event.
	SequenceNum param.Opt[string] `json:"sequenceNum,omitzero"`
	paramObj
}

func (r AirEventUpdateParamsReceiver) MarshalJSON() (data []byte, err error) {
	type shadow AirEventUpdateParamsReceiver
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventUpdateParamsReceiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this Air Event.
type AirEventUpdateParamsRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Optional remark ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalRemarkID param.Opt[string] `json:"externalRemarkId,omitzero"`
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// User who published the remark.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r AirEventUpdateParamsRemark) MarshalJSON() (data []byte, err error) {
	type shadow AirEventUpdateParamsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventUpdateParamsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of tanker aircraft associated with this Air Event.
type AirEventUpdateParamsTanker struct {
	// Alternate mission identifier of this tanker provided by source.
	AltTankerMissionID param.Opt[string] `json:"altTankerMissionId,omitzero"`
	// The Air Mobility Command (AMC) mission identifier of this tanker.
	AmcTankerMissionID param.Opt[string] `json:"amcTankerMissionId,omitzero"`
	// Flag indicating that this tanker is flying a dual role mission in an air
	// refueling event.
	DualRole param.Opt[bool] `json:"dualRole,omitzero"`
	// Optional tanker identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalTankerID param.Opt[string] `json:"externalTankerId,omitzero"`
	// Total weight of the fuel transferred from this tanker during an air refueling
	// event, in pounds.
	FuelOff param.Opt[float64] `json:"fuelOff,omitzero"`
	// The UDL ID of the airfield this tanker is associated with for this event.
	IDTankerAirfield param.Opt[string] `json:"idTankerAirfield,omitzero"`
	// The UDL ID of the mission this tanker is associated with for this event.
	IDTankerMission param.Opt[string] `json:"idTankerMission,omitzero"`
	// The UDL ID of the aircraft sortie this tanker is associated with for this event.
	IDTankerSortie param.Opt[string] `json:"idTankerSortie,omitzero"`
	// The call sign assigned to this tanker.
	TankerCallSign param.Opt[string] `json:"tankerCallSign,omitzero"`
	// Position of this tanker within a group of tankers in an air refueling event.
	TankerCellPosition param.Opt[int64] `json:"tankerCellPosition,omitzero"`
	// Coordination record identifier of this tanker.
	TankerCoord param.Opt[string] `json:"tankerCoord,omitzero"`
	// Type of fuel delivery method used by the tanker during an air refueling event
	// (BOOM, DROGUE, BOTH).
	TankerDeliveryMethod param.Opt[string] `json:"tankerDeliveryMethod,omitzero"`
	// Location the tanker has been deployed to in preparation for an air refueling
	// event.
	TankerDeployedIcao param.Opt[string] `json:"tankerDeployedICAO,omitzero"`
	// Type of fuel being transferred from the tanker in an air refueling event.
	TankerFuelType param.Opt[string] `json:"tankerFuelType,omitzero"`
	// Identifies the itinerary point of a mission that this tanker is linked to.
	TankerLegNum param.Opt[int64] `json:"tankerLegNum,omitzero"`
	// The Model Design Series designation of this tanker.
	TankerMds param.Opt[string] `json:"tankerMDS,omitzero"`
	// The wing or unit that owns this tanker.
	TankerOwner param.Opt[string] `json:"tankerOwner,omitzero"`
	// The name and/or number of the point of contact for this tanker.
	TankerPoc param.Opt[string] `json:"tankerPOC,omitzero"`
	paramObj
}

func (r AirEventUpdateParamsTanker) MarshalJSON() (data []byte, err error) {
	type shadow AirEventUpdateParamsTanker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventUpdateParamsTanker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirEventListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirEventListParams]'s query parameters as `url.Values`.
func (r AirEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirEventCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirEventCountParams]'s query parameters as `url.Values`.
func (r AirEventCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirEventNewBulkParams struct {
	Body []AirEventNewBulkParamsBody
	paramObj
}

func (r AirEventNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *AirEventNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Information related to an air event (e.g. FUEL TRANSFER, AIR DROP) and the
// associated aircraft.
//
// The properties ClassificationMarking, DataMode, Source, Type are required.
type AirEventNewBulkParamsBody struct {
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
	// Type of air event (e.g. FUEL TRANSFER, AIR DROP, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The actual arrival time of the aircraft at the air event, in ISO 8601 UTC format
	// with millisecond precision.
	ActualArrTime param.Opt[time.Time] `json:"actualArrTime,omitzero" format:"date-time"`
	// The actual departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	ActualDepTime param.Opt[time.Time] `json:"actualDepTime,omitzero" format:"date-time"`
	// The Air Refueling Control Time is the planned time the tanker aircraft will
	// transfer fuel to the receiver aircraft, in ISO 8601 UTC format, with millisecond
	// precision.
	Arct param.Opt[time.Time] `json:"arct,omitzero" format:"date-time"`
	// Type of process used by AMC to schedule this air refueling event. Possible
	// values are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched
	// Theater Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S
	// (Soft Air Refueling), T (Matched Theater Operation Short Notice (Theater
	// Assets)), V (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short
	// Notice (AMC Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z
	// (Other Air Refueling).
	ArEventType param.Opt[string] `json:"arEventType,omitzero"`
	// The purpose of the air event at the arrival location. Can be either descriptive
	// text such as 'fuel onload' or a purpose code specified by the provider, such as
	// 'A'.
	ArrPurpose param.Opt[string] `json:"arrPurpose,omitzero"`
	// Identifier of the air refueling track, if applicable.
	ArTrackID param.Opt[string] `json:"arTrackId,omitzero"`
	// Name of the air refueling track, if applicable.
	ArTrackName param.Opt[string] `json:"arTrackName,omitzero"`
	// Altitude of this air event, in feet.
	BaseAlt param.Opt[float64] `json:"baseAlt,omitzero"`
	// Flag indicating that this air refueling event has been cancelled.
	Cancelled param.Opt[bool] `json:"cancelled,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The purpose of the air event at the departure location. Can be either
	// descriptive text such as 'fuel onload' or a purpose code specified by the
	// provider, such as 'A'.
	DepPurpose param.Opt[string] `json:"depPurpose,omitzero"`
	// The current estimated arrival time of the aircraft at the air event, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// The current estimated departure time of the aircraft from the air event, in ISO
	// 8601 UTC format with millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// Optional air event ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalAirEventID param.Opt[string] `json:"externalAirEventId,omitzero"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalArTrackID param.Opt[string] `json:"externalARTrackId,omitzero"`
	// The UDL unique identifier of the mission associated with this air event.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// The UDL unique identifier of the sortie associated with this air event.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The location representing this air event specified as a feature Id. Locations
	// specified include air refueling track Ids and air drop event locations.
	Location param.Opt[string] `json:"location,omitzero"`
	// The number of tankers requested for an air refueling event.
	NumTankers param.Opt[int64] `json:"numTankers,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The scheduled arrival time of the aircraft at the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedArrTime param.Opt[time.Time] `json:"plannedArrTime,omitzero" format:"date-time"`
	// The scheduled departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedDepTime param.Opt[time.Time] `json:"plannedDepTime,omitzero" format:"date-time"`
	// Priority of this air event.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// Flag indicating if the receiver unit has requested flying an air refueling track
	// in both directions.
	RevTrack param.Opt[bool] `json:"revTrack,omitzero"`
	// The Rendezvous Control Time is the planned time the tanker and receiver aircraft
	// will rendezvous for an en route type air refueling event, in ISO 8601 UTC
	// format, with millisecond precision.
	Rzct param.Opt[time.Time] `json:"rzct,omitzero" format:"date-time"`
	// Rendezvous point for the tanker and receiver during this air refueling event.
	// Possible values are AN (Anchor Nav Point), AP (Anchor Pattern), CP (Control
	// Point), ET (Entry Point), EX (Exit Point), IP (Initial Point), NC (Nav Check
	// Point).
	RzPoint param.Opt[string] `json:"rzPoint,omitzero"`
	// Type of rendezvous used for this air refueling event. Possible values are BUD
	// (Buddy), EN (Enroute), GCI (Ground Control), PP (Point Parallel).
	RzType param.Opt[string] `json:"rzType,omitzero"`
	// Flag indicating that the receiver unit has requested flying a short portion of
	// an air refueling track.
	ShortTrack param.Opt[bool] `json:"shortTrack,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Status of this air refueling event track reservation. Receivers are responsible
	// for scheduling or reserving air refueling tracks. Possible values are A
	// (Altitude Reservation), R (Reserved), or Q (Questionable).
	StatusCode param.Opt[string] `json:"statusCode,omitzero"`
	// Length of time the receiver unit has requested for an air event, in hours.
	TrackTime param.Opt[float64] `json:"trackTime,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Collection of receiver aircraft associated with this Air Event.
	Receivers []AirEventNewBulkParamsBodyReceiver `json:"receivers,omitzero"`
	// Collection of remarks associated with this Air Event.
	Remarks []AirEventNewBulkParamsBodyRemark `json:"remarks,omitzero"`
	// Collection of tanker aircraft associated with this Air Event.
	Tankers []AirEventNewBulkParamsBodyTanker `json:"tankers,omitzero"`
	paramObj
}

func (r AirEventNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AirEventNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AirEventNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of receiver aircraft associated with this Air Event.
type AirEventNewBulkParamsBodyReceiver struct {
	// Alternate mission identifier of this receiver provided by source.
	AltReceiverMissionID param.Opt[string] `json:"altReceiverMissionId,omitzero"`
	// The Air Mobility Command (AMC) mission identifier of this receiver.
	AmcReceiverMissionID param.Opt[string] `json:"amcReceiverMissionId,omitzero"`
	// Optional receiver identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalReceiverID param.Opt[string] `json:"externalReceiverId,omitzero"`
	// Total weight of the fuel transferred to this receiver during an air refueling
	// event, in pounds.
	FuelOn param.Opt[float64] `json:"fuelOn,omitzero"`
	// The UDL ID of the airfield this receiver is associated with for this event.
	IDReceiverAirfield param.Opt[string] `json:"idReceiverAirfield,omitzero"`
	// The UDL ID of the mission this receiver is associated with for this event.
	IDReceiverMission param.Opt[string] `json:"idReceiverMission,omitzero"`
	// The UDL ID of the aircraft sortie this receiver is associated with for this
	// event.
	IDReceiverSortie param.Opt[string] `json:"idReceiverSortie,omitzero"`
	// Number of aircraft contained within one receiver coordination record for an air
	// refueling event.
	NumRecAircraft param.Opt[int64] `json:"numRecAircraft,omitzero"`
	// The package identifier for the receiver in an air refueling event.
	PackageID param.Opt[string] `json:"packageId,omitzero"`
	// The call sign assigned to this receiver.
	ReceiverCallSign param.Opt[string] `json:"receiverCallSign,omitzero"`
	// Position of this receiver within a group of receivers in an air refueling event.
	ReceiverCellPosition param.Opt[int64] `json:"receiverCellPosition,omitzero"`
	// Coordination record identifier of this receiver.
	ReceiverCoord param.Opt[string] `json:"receiverCoord,omitzero"`
	// Type of fuel delivery method used by the receiver during an air refueling event
	// (BOOM, DROGUE, BOTH).
	ReceiverDeliveryMethod param.Opt[string] `json:"receiverDeliveryMethod,omitzero"`
	// Location the receiver is deployed to for an air refueling event.
	ReceiverDeployedIcao param.Opt[string] `json:"receiverDeployedICAO,omitzero"`
	// Name of the receiver exercise associated with an air refueling event.
	ReceiverExercise param.Opt[string] `json:"receiverExercise,omitzero"`
	// Type of fuel being transferred to the receiver in an air refueling event.
	ReceiverFuelType param.Opt[string] `json:"receiverFuelType,omitzero"`
	// Identifies the itinerary point of a mission that this receiver is linked to.
	ReceiverLegNum param.Opt[int64] `json:"receiverLegNum,omitzero"`
	// The Model Design Series designation of this receiver.
	ReceiverMds param.Opt[string] `json:"receiverMDS,omitzero"`
	// The wing or unit that owns this receiver.
	ReceiverOwner param.Opt[string] `json:"receiverOwner,omitzero"`
	// The name and/or number of the point of contact for this receiver.
	ReceiverPoc param.Opt[string] `json:"receiverPOC,omitzero"`
	// The major command level (MAJCOM) or foreign military sales (FMS) name of the
	// receiver's organization. The tanker flying hours used for an air refueling event
	// are logged against the receiver MAJCOM or foreign government being supported.
	RecOrg param.Opt[string] `json:"recOrg,omitzero"`
	// Indicates the unique number by Unit ID, which identifies an air refueling event.
	SequenceNum param.Opt[string] `json:"sequenceNum,omitzero"`
	paramObj
}

func (r AirEventNewBulkParamsBodyReceiver) MarshalJSON() (data []byte, err error) {
	type shadow AirEventNewBulkParamsBodyReceiver
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventNewBulkParamsBodyReceiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this Air Event.
type AirEventNewBulkParamsBodyRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Optional remark ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalRemarkID param.Opt[string] `json:"externalRemarkId,omitzero"`
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// User who published the remark.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r AirEventNewBulkParamsBodyRemark) MarshalJSON() (data []byte, err error) {
	type shadow AirEventNewBulkParamsBodyRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventNewBulkParamsBodyRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of tanker aircraft associated with this Air Event.
type AirEventNewBulkParamsBodyTanker struct {
	// Alternate mission identifier of this tanker provided by source.
	AltTankerMissionID param.Opt[string] `json:"altTankerMissionId,omitzero"`
	// The Air Mobility Command (AMC) mission identifier of this tanker.
	AmcTankerMissionID param.Opt[string] `json:"amcTankerMissionId,omitzero"`
	// Flag indicating that this tanker is flying a dual role mission in an air
	// refueling event.
	DualRole param.Opt[bool] `json:"dualRole,omitzero"`
	// Optional tanker identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalTankerID param.Opt[string] `json:"externalTankerId,omitzero"`
	// Total weight of the fuel transferred from this tanker during an air refueling
	// event, in pounds.
	FuelOff param.Opt[float64] `json:"fuelOff,omitzero"`
	// The UDL ID of the airfield this tanker is associated with for this event.
	IDTankerAirfield param.Opt[string] `json:"idTankerAirfield,omitzero"`
	// The UDL ID of the mission this tanker is associated with for this event.
	IDTankerMission param.Opt[string] `json:"idTankerMission,omitzero"`
	// The UDL ID of the aircraft sortie this tanker is associated with for this event.
	IDTankerSortie param.Opt[string] `json:"idTankerSortie,omitzero"`
	// The call sign assigned to this tanker.
	TankerCallSign param.Opt[string] `json:"tankerCallSign,omitzero"`
	// Position of this tanker within a group of tankers in an air refueling event.
	TankerCellPosition param.Opt[int64] `json:"tankerCellPosition,omitzero"`
	// Coordination record identifier of this tanker.
	TankerCoord param.Opt[string] `json:"tankerCoord,omitzero"`
	// Type of fuel delivery method used by the tanker during an air refueling event
	// (BOOM, DROGUE, BOTH).
	TankerDeliveryMethod param.Opt[string] `json:"tankerDeliveryMethod,omitzero"`
	// Location the tanker has been deployed to in preparation for an air refueling
	// event.
	TankerDeployedIcao param.Opt[string] `json:"tankerDeployedICAO,omitzero"`
	// Type of fuel being transferred from the tanker in an air refueling event.
	TankerFuelType param.Opt[string] `json:"tankerFuelType,omitzero"`
	// Identifies the itinerary point of a mission that this tanker is linked to.
	TankerLegNum param.Opt[int64] `json:"tankerLegNum,omitzero"`
	// The Model Design Series designation of this tanker.
	TankerMds param.Opt[string] `json:"tankerMDS,omitzero"`
	// The wing or unit that owns this tanker.
	TankerOwner param.Opt[string] `json:"tankerOwner,omitzero"`
	// The name and/or number of the point of contact for this tanker.
	TankerPoc param.Opt[string] `json:"tankerPOC,omitzero"`
	paramObj
}

func (r AirEventNewBulkParamsBodyTanker) MarshalJSON() (data []byte, err error) {
	type shadow AirEventNewBulkParamsBodyTanker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventNewBulkParamsBodyTanker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirEventGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirEventGetParams]'s query parameters as `url.Values`.
func (r AirEventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirEventTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirEventTupleParams]'s query parameters as `url.Values`.
func (r AirEventTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirEventUnvalidatedPublishParams struct {
	Body []AirEventUnvalidatedPublishParamsBody
	paramObj
}

func (r AirEventUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *AirEventUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Information related to an air event (e.g. FUEL TRANSFER, AIR DROP) and the
// associated aircraft.
//
// The properties ClassificationMarking, DataMode, Source, Type are required.
type AirEventUnvalidatedPublishParamsBody struct {
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
	// Type of air event (e.g. FUEL TRANSFER, AIR DROP, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The actual arrival time of the aircraft at the air event, in ISO 8601 UTC format
	// with millisecond precision.
	ActualArrTime param.Opt[time.Time] `json:"actualArrTime,omitzero" format:"date-time"`
	// The actual departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	ActualDepTime param.Opt[time.Time] `json:"actualDepTime,omitzero" format:"date-time"`
	// The Air Refueling Control Time is the planned time the tanker aircraft will
	// transfer fuel to the receiver aircraft, in ISO 8601 UTC format, with millisecond
	// precision.
	Arct param.Opt[time.Time] `json:"arct,omitzero" format:"date-time"`
	// Type of process used by AMC to schedule this air refueling event. Possible
	// values are A (Matched Long Range), F (Matched AMC Short Notice), N (Unmatched
	// Theater Operation Short Notice (Theater Assets)), R, Unmatched Long Range, S
	// (Soft Air Refueling), T (Matched Theater Operation Short Notice (Theater
	// Assets)), V (Unmatched AMC Short Notice), X (Unmatched Theater Operation Short
	// Notice (AMC Assets)), Y (Matched Theater Operation Short Notice (AMC Assets)), Z
	// (Other Air Refueling).
	ArEventType param.Opt[string] `json:"arEventType,omitzero"`
	// The purpose of the air event at the arrival location. Can be either descriptive
	// text such as 'fuel onload' or a purpose code specified by the provider, such as
	// 'A'.
	ArrPurpose param.Opt[string] `json:"arrPurpose,omitzero"`
	// Identifier of the air refueling track, if applicable.
	ArTrackID param.Opt[string] `json:"arTrackId,omitzero"`
	// Name of the air refueling track, if applicable.
	ArTrackName param.Opt[string] `json:"arTrackName,omitzero"`
	// Altitude of this air event, in feet.
	BaseAlt param.Opt[float64] `json:"baseAlt,omitzero"`
	// Flag indicating that this air refueling event has been cancelled.
	Cancelled param.Opt[bool] `json:"cancelled,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The purpose of the air event at the departure location. Can be either
	// descriptive text such as 'fuel onload' or a purpose code specified by the
	// provider, such as 'A'.
	DepPurpose param.Opt[string] `json:"depPurpose,omitzero"`
	// The current estimated arrival time of the aircraft at the air event, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// The current estimated departure time of the aircraft from the air event, in ISO
	// 8601 UTC format with millisecond precision.
	EstDepTime param.Opt[time.Time] `json:"estDepTime,omitzero" format:"date-time"`
	// Optional air event ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalAirEventID param.Opt[string] `json:"externalAirEventId,omitzero"`
	// Optional air refueling track ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalArTrackID param.Opt[string] `json:"externalARTrackId,omitzero"`
	// The UDL unique identifier of the mission associated with this air event.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// The UDL unique identifier of the sortie associated with this air event.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The location representing this air event specified as a feature Id. Locations
	// specified include air refueling track Ids and air drop event locations.
	Location param.Opt[string] `json:"location,omitzero"`
	// The number of tankers requested for an air refueling event.
	NumTankers param.Opt[int64] `json:"numTankers,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The scheduled arrival time of the aircraft at the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedArrTime param.Opt[time.Time] `json:"plannedArrTime,omitzero" format:"date-time"`
	// The scheduled departure time of the aircraft from the air event, in ISO 8601 UTC
	// format with millisecond precision.
	PlannedDepTime param.Opt[time.Time] `json:"plannedDepTime,omitzero" format:"date-time"`
	// Priority of this air event.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// Flag indicating if the receiver unit has requested flying an air refueling track
	// in both directions.
	RevTrack param.Opt[bool] `json:"revTrack,omitzero"`
	// The Rendezvous Control Time is the planned time the tanker and receiver aircraft
	// will rendezvous for an en route type air refueling event, in ISO 8601 UTC
	// format, with millisecond precision.
	Rzct param.Opt[time.Time] `json:"rzct,omitzero" format:"date-time"`
	// Rendezvous point for the tanker and receiver during this air refueling event.
	// Possible values are AN (Anchor Nav Point), AP (Anchor Pattern), CP (Control
	// Point), ET (Entry Point), EX (Exit Point), IP (Initial Point), NC (Nav Check
	// Point).
	RzPoint param.Opt[string] `json:"rzPoint,omitzero"`
	// Type of rendezvous used for this air refueling event. Possible values are BUD
	// (Buddy), EN (Enroute), GCI (Ground Control), PP (Point Parallel).
	RzType param.Opt[string] `json:"rzType,omitzero"`
	// Flag indicating that the receiver unit has requested flying a short portion of
	// an air refueling track.
	ShortTrack param.Opt[bool] `json:"shortTrack,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Status of this air refueling event track reservation. Receivers are responsible
	// for scheduling or reserving air refueling tracks. Possible values are A
	// (Altitude Reservation), R (Reserved), or Q (Questionable).
	StatusCode param.Opt[string] `json:"statusCode,omitzero"`
	// Length of time the receiver unit has requested for an air event, in hours.
	TrackTime param.Opt[float64] `json:"trackTime,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Collection of receiver aircraft associated with this Air Event.
	Receivers []AirEventUnvalidatedPublishParamsBodyReceiver `json:"receivers,omitzero"`
	// Collection of remarks associated with this Air Event.
	Remarks []AirEventUnvalidatedPublishParamsBodyRemark `json:"remarks,omitzero"`
	// Collection of tanker aircraft associated with this Air Event.
	Tankers []AirEventUnvalidatedPublishParamsBodyTanker `json:"tankers,omitzero"`
	paramObj
}

func (r AirEventUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AirEventUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AirEventUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of receiver aircraft associated with this Air Event.
type AirEventUnvalidatedPublishParamsBodyReceiver struct {
	// Alternate mission identifier of this receiver provided by source.
	AltReceiverMissionID param.Opt[string] `json:"altReceiverMissionId,omitzero"`
	// The Air Mobility Command (AMC) mission identifier of this receiver.
	AmcReceiverMissionID param.Opt[string] `json:"amcReceiverMissionId,omitzero"`
	// Optional receiver identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalReceiverID param.Opt[string] `json:"externalReceiverId,omitzero"`
	// Total weight of the fuel transferred to this receiver during an air refueling
	// event, in pounds.
	FuelOn param.Opt[float64] `json:"fuelOn,omitzero"`
	// The UDL ID of the airfield this receiver is associated with for this event.
	IDReceiverAirfield param.Opt[string] `json:"idReceiverAirfield,omitzero"`
	// The UDL ID of the mission this receiver is associated with for this event.
	IDReceiverMission param.Opt[string] `json:"idReceiverMission,omitzero"`
	// The UDL ID of the aircraft sortie this receiver is associated with for this
	// event.
	IDReceiverSortie param.Opt[string] `json:"idReceiverSortie,omitzero"`
	// Number of aircraft contained within one receiver coordination record for an air
	// refueling event.
	NumRecAircraft param.Opt[int64] `json:"numRecAircraft,omitzero"`
	// The package identifier for the receiver in an air refueling event.
	PackageID param.Opt[string] `json:"packageId,omitzero"`
	// The call sign assigned to this receiver.
	ReceiverCallSign param.Opt[string] `json:"receiverCallSign,omitzero"`
	// Position of this receiver within a group of receivers in an air refueling event.
	ReceiverCellPosition param.Opt[int64] `json:"receiverCellPosition,omitzero"`
	// Coordination record identifier of this receiver.
	ReceiverCoord param.Opt[string] `json:"receiverCoord,omitzero"`
	// Type of fuel delivery method used by the receiver during an air refueling event
	// (BOOM, DROGUE, BOTH).
	ReceiverDeliveryMethod param.Opt[string] `json:"receiverDeliveryMethod,omitzero"`
	// Location the receiver is deployed to for an air refueling event.
	ReceiverDeployedIcao param.Opt[string] `json:"receiverDeployedICAO,omitzero"`
	// Name of the receiver exercise associated with an air refueling event.
	ReceiverExercise param.Opt[string] `json:"receiverExercise,omitzero"`
	// Type of fuel being transferred to the receiver in an air refueling event.
	ReceiverFuelType param.Opt[string] `json:"receiverFuelType,omitzero"`
	// Identifies the itinerary point of a mission that this receiver is linked to.
	ReceiverLegNum param.Opt[int64] `json:"receiverLegNum,omitzero"`
	// The Model Design Series designation of this receiver.
	ReceiverMds param.Opt[string] `json:"receiverMDS,omitzero"`
	// The wing or unit that owns this receiver.
	ReceiverOwner param.Opt[string] `json:"receiverOwner,omitzero"`
	// The name and/or number of the point of contact for this receiver.
	ReceiverPoc param.Opt[string] `json:"receiverPOC,omitzero"`
	// The major command level (MAJCOM) or foreign military sales (FMS) name of the
	// receiver's organization. The tanker flying hours used for an air refueling event
	// are logged against the receiver MAJCOM or foreign government being supported.
	RecOrg param.Opt[string] `json:"recOrg,omitzero"`
	// Indicates the unique number by Unit ID, which identifies an air refueling event.
	SequenceNum param.Opt[string] `json:"sequenceNum,omitzero"`
	paramObj
}

func (r AirEventUnvalidatedPublishParamsBodyReceiver) MarshalJSON() (data []byte, err error) {
	type shadow AirEventUnvalidatedPublishParamsBodyReceiver
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventUnvalidatedPublishParamsBodyReceiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this Air Event.
type AirEventUnvalidatedPublishParamsBodyRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Optional remark ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalRemarkID param.Opt[string] `json:"externalRemarkId,omitzero"`
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// User who published the remark.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r AirEventUnvalidatedPublishParamsBodyRemark) MarshalJSON() (data []byte, err error) {
	type shadow AirEventUnvalidatedPublishParamsBodyRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventUnvalidatedPublishParamsBodyRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of tanker aircraft associated with this Air Event.
type AirEventUnvalidatedPublishParamsBodyTanker struct {
	// Alternate mission identifier of this tanker provided by source.
	AltTankerMissionID param.Opt[string] `json:"altTankerMissionId,omitzero"`
	// The Air Mobility Command (AMC) mission identifier of this tanker.
	AmcTankerMissionID param.Opt[string] `json:"amcTankerMissionId,omitzero"`
	// Flag indicating that this tanker is flying a dual role mission in an air
	// refueling event.
	DualRole param.Opt[bool] `json:"dualRole,omitzero"`
	// Optional tanker identifier from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExternalTankerID param.Opt[string] `json:"externalTankerId,omitzero"`
	// Total weight of the fuel transferred from this tanker during an air refueling
	// event, in pounds.
	FuelOff param.Opt[float64] `json:"fuelOff,omitzero"`
	// The UDL ID of the airfield this tanker is associated with for this event.
	IDTankerAirfield param.Opt[string] `json:"idTankerAirfield,omitzero"`
	// The UDL ID of the mission this tanker is associated with for this event.
	IDTankerMission param.Opt[string] `json:"idTankerMission,omitzero"`
	// The UDL ID of the aircraft sortie this tanker is associated with for this event.
	IDTankerSortie param.Opt[string] `json:"idTankerSortie,omitzero"`
	// The call sign assigned to this tanker.
	TankerCallSign param.Opt[string] `json:"tankerCallSign,omitzero"`
	// Position of this tanker within a group of tankers in an air refueling event.
	TankerCellPosition param.Opt[int64] `json:"tankerCellPosition,omitzero"`
	// Coordination record identifier of this tanker.
	TankerCoord param.Opt[string] `json:"tankerCoord,omitzero"`
	// Type of fuel delivery method used by the tanker during an air refueling event
	// (BOOM, DROGUE, BOTH).
	TankerDeliveryMethod param.Opt[string] `json:"tankerDeliveryMethod,omitzero"`
	// Location the tanker has been deployed to in preparation for an air refueling
	// event.
	TankerDeployedIcao param.Opt[string] `json:"tankerDeployedICAO,omitzero"`
	// Type of fuel being transferred from the tanker in an air refueling event.
	TankerFuelType param.Opt[string] `json:"tankerFuelType,omitzero"`
	// Identifies the itinerary point of a mission that this tanker is linked to.
	TankerLegNum param.Opt[int64] `json:"tankerLegNum,omitzero"`
	// The Model Design Series designation of this tanker.
	TankerMds param.Opt[string] `json:"tankerMDS,omitzero"`
	// The wing or unit that owns this tanker.
	TankerOwner param.Opt[string] `json:"tankerOwner,omitzero"`
	// The name and/or number of the point of contact for this tanker.
	TankerPoc param.Opt[string] `json:"tankerPOC,omitzero"`
	paramObj
}

func (r AirEventUnvalidatedPublishParamsBodyTanker) MarshalJSON() (data []byte, err error) {
	type shadow AirEventUnvalidatedPublishParamsBodyTanker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirEventUnvalidatedPublishParamsBodyTanker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
