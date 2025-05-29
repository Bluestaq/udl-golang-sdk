// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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

// AirfieldSlotConsumptionService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirfieldSlotConsumptionService] method instead.
type AirfieldSlotConsumptionService struct {
	Options []option.RequestOption
}

// NewAirfieldSlotConsumptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAirfieldSlotConsumptionService(opts ...option.RequestOption) (r AirfieldSlotConsumptionService) {
	r = AirfieldSlotConsumptionService{}
	r.Options = opts
	return
}

// Service operation to take a single airfieldslotconsumption record as a POST body
// and ingest into the database. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *AirfieldSlotConsumptionService) New(ctx context.Context, body AirfieldSlotConsumptionNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airfieldslotconsumption"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single airfieldslotconsumption record by its unique
// ID passed as a path parameter.
func (r *AirfieldSlotConsumptionService) Get(ctx context.Context, id string, query AirfieldSlotConsumptionGetParams, opts ...option.RequestOption) (res *AirfieldslotconsumptionFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfieldslotconsumption/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single AirfieldSlotConsumption. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *AirfieldSlotConsumptionService) Update(ctx context.Context, id string, body AirfieldSlotConsumptionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfieldslotconsumption/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirfieldSlotConsumptionService) List(ctx context.Context, query AirfieldSlotConsumptionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AirfieldslotconsumptionAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/airfieldslotconsumption"
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
func (r *AirfieldSlotConsumptionService) ListAutoPaging(ctx context.Context, query AirfieldSlotConsumptionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AirfieldslotconsumptionAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an airfieldslotconsumption record specified by the
// passed ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AirfieldSlotConsumptionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfieldslotconsumption/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirfieldSlotConsumptionService) Count(ctx context.Context, query AirfieldSlotConsumptionCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airfieldslotconsumption/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AirfieldSlotConsumptionService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *AirfieldSlotConsumptionQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airfieldslotconsumption/queryhelp"
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
func (r *AirfieldSlotConsumptionService) Tuple(ctx context.Context, query AirfieldSlotConsumptionTupleParams, opts ...option.RequestOption) (res *[]AirfieldslotconsumptionFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airfieldslotconsumption/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Airfield slot use data. Contains the dynamic data associated with the status and
// use of specific airfield slots.
type AirfieldslotconsumptionAbridged struct {
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
	DataMode AirfieldslotconsumptionAbridgedDataMode `json:"dataMode,required"`
	// Unique identifier of the airfield slot for which this slot consumption record is
	// referencing.
	IDAirfieldSlot string `json:"idAirfieldSlot,required"`
	// Number of aircraft using this slot for this time.
	NumAircraft int64 `json:"numAircraft,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the slot window, in ISO 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate identifier of the sortie arriving at the slot start time provided by
	// the source.
	AltArrSortieID string `json:"altArrSortieId"`
	// Alternate identifier of the sortie departing at the slot end time provided by
	// the source.
	AltDepSortieID string `json:"altDepSortieId"`
	// Comments from the approver.
	AppComment string `json:"appComment"`
	// Initials of the person approving the use of this slot. Use SYSTEM if
	// auto-approved without human involvement.
	AppInitials string `json:"appInitials"`
	// Short name of the organization approving the use of this slot.
	AppOrg string `json:"appOrg"`
	// Array of call signs of the aircraft using this slot.
	CallSigns []string `json:"callSigns"`
	// Identifying name of the aircraft using this slot. Names are often Prior
	// Permission Required (PPR) numbers or other similar human-readable identifiers.
	Consumer string `json:"consumer"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The end of the slot window, in ISO 8601 UTC format.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Unique identifier of the sortie arriving at the slot start time.
	IDArrSortie string `json:"idArrSortie"`
	// Unique identifier of the sortie departing at the slot end time.
	IDDepSortie string `json:"idDepSortie"`
	// Mission identifier using this slot according to Mobility Air Forces (MAF)
	// Encode/Decode procedures.
	MissionID string `json:"missionId"`
	// The aircraft Model Design Series designation of the aircraft occupying this
	// slot.
	OccAircraftMds string `json:"occAircraftMDS"`
	// Time the aircraft began occupying this slot, in ISO 8601 UTC format with
	// millisecond precision.
	OccStartTime time.Time `json:"occStartTime" format:"date-time"`
	// The tail number of the aircraft occupying this slot.
	OccTailNumber string `json:"occTailNumber"`
	// Flag indicating if the slot is occupied.
	Occupied bool `json:"occupied"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Comments from the requester.
	ReqComment string `json:"reqComment"`
	// Initials of the person requesting the use of this slot. Use SYSTEM if this
	// request is auto-generated by an auto-planning system.
	ReqInitials string `json:"reqInitials"`
	// Short name of the organization requesting use of this slot.
	ReqOrg string `json:"reqOrg"`
	// The aircraft Model Design Series designation of the aircraft this slot is
	// reserved for.
	ResAircraftMds string `json:"resAircraftMDS"`
	// Mission identifier reserving this slot according to Mobility Air Forces (MAF)
	// Encode/Decode procedures.
	ResMissionID string `json:"resMissionId"`
	// The reason the slot reservation was made.
	ResReason string `json:"resReason"`
	// The tail number of the aircraft this slot is reserved for.
	ResTailNumber string `json:"resTailNumber"`
	// Indicates the type of reservation (e.g. M for Mission, A for Aircraft, O for
	// Other).
	ResType string `json:"resType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Current status of this slot (REQUESTED / APPROVED / DENIED / BLOCKED / OTHER).
	//
	// Any of "REQUESTED", "APPROVED", "DENIED", "BLOCKED", "OTHER".
	Status AirfieldslotconsumptionAbridgedStatus `json:"status"`
	// The desired time for aircraft action such as landing, take off, parking, etc.,
	// in ISO 8601 UTC format.
	TargetTime time.Time `json:"targetTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDAirfieldSlot        respjson.Field
		NumAircraft           respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		AltArrSortieID        respjson.Field
		AltDepSortieID        respjson.Field
		AppComment            respjson.Field
		AppInitials           respjson.Field
		AppOrg                respjson.Field
		CallSigns             respjson.Field
		Consumer              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndTime               respjson.Field
		IDArrSortie           respjson.Field
		IDDepSortie           respjson.Field
		MissionID             respjson.Field
		OccAircraftMds        respjson.Field
		OccStartTime          respjson.Field
		OccTailNumber         respjson.Field
		Occupied              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ReqComment            respjson.Field
		ReqInitials           respjson.Field
		ReqOrg                respjson.Field
		ResAircraftMds        respjson.Field
		ResMissionID          respjson.Field
		ResReason             respjson.Field
		ResTailNumber         respjson.Field
		ResType               respjson.Field
		SourceDl              respjson.Field
		Status                respjson.Field
		TargetTime            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirfieldslotconsumptionAbridged) RawJSON() string { return r.JSON.raw }
func (r *AirfieldslotconsumptionAbridged) UnmarshalJSON(data []byte) error {
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
type AirfieldslotconsumptionAbridgedDataMode string

const (
	AirfieldslotconsumptionAbridgedDataModeReal      AirfieldslotconsumptionAbridgedDataMode = "REAL"
	AirfieldslotconsumptionAbridgedDataModeTest      AirfieldslotconsumptionAbridgedDataMode = "TEST"
	AirfieldslotconsumptionAbridgedDataModeSimulated AirfieldslotconsumptionAbridgedDataMode = "SIMULATED"
	AirfieldslotconsumptionAbridgedDataModeExercise  AirfieldslotconsumptionAbridgedDataMode = "EXERCISE"
)

// Current status of this slot (REQUESTED / APPROVED / DENIED / BLOCKED / OTHER).
type AirfieldslotconsumptionAbridgedStatus string

const (
	AirfieldslotconsumptionAbridgedStatusRequested AirfieldslotconsumptionAbridgedStatus = "REQUESTED"
	AirfieldslotconsumptionAbridgedStatusApproved  AirfieldslotconsumptionAbridgedStatus = "APPROVED"
	AirfieldslotconsumptionAbridgedStatusDenied    AirfieldslotconsumptionAbridgedStatus = "DENIED"
	AirfieldslotconsumptionAbridgedStatusBlocked   AirfieldslotconsumptionAbridgedStatus = "BLOCKED"
	AirfieldslotconsumptionAbridgedStatusOther     AirfieldslotconsumptionAbridgedStatus = "OTHER"
)

// Airfield slot use data. Contains the dynamic data associated with the status and
// use of specific airfield slots.
type AirfieldslotconsumptionFull struct {
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
	DataMode AirfieldslotconsumptionFullDataMode `json:"dataMode,required"`
	// Unique identifier of the airfield slot for which this slot consumption record is
	// referencing.
	IDAirfieldSlot string `json:"idAirfieldSlot,required"`
	// Number of aircraft using this slot for this time.
	NumAircraft int64 `json:"numAircraft,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the slot window, in ISO 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate identifier of the sortie arriving at the slot start time provided by
	// the source.
	AltArrSortieID string `json:"altArrSortieId"`
	// Alternate identifier of the sortie departing at the slot end time provided by
	// the source.
	AltDepSortieID string `json:"altDepSortieId"`
	// Comments from the approver.
	AppComment string `json:"appComment"`
	// Initials of the person approving the use of this slot. Use SYSTEM if
	// auto-approved without human involvement.
	AppInitials string `json:"appInitials"`
	// Short name of the organization approving the use of this slot.
	AppOrg string `json:"appOrg"`
	// Array of call signs of the aircraft using this slot.
	CallSigns []string `json:"callSigns"`
	// Identifying name of the aircraft using this slot. Names are often Prior
	// Permission Required (PPR) numbers or other similar human-readable identifiers.
	Consumer string `json:"consumer"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The end of the slot window, in ISO 8601 UTC format.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Unique identifier of the sortie arriving at the slot start time.
	IDArrSortie string `json:"idArrSortie"`
	// Unique identifier of the sortie departing at the slot end time.
	IDDepSortie string `json:"idDepSortie"`
	// Mission identifier using this slot according to Mobility Air Forces (MAF)
	// Encode/Decode procedures.
	MissionID string `json:"missionId"`
	// The aircraft Model Design Series designation of the aircraft occupying this
	// slot.
	OccAircraftMds string `json:"occAircraftMDS"`
	// Time the aircraft began occupying this slot, in ISO 8601 UTC format with
	// millisecond precision.
	OccStartTime time.Time `json:"occStartTime" format:"date-time"`
	// The tail number of the aircraft occupying this slot.
	OccTailNumber string `json:"occTailNumber"`
	// Flag indicating if the slot is occupied.
	Occupied bool `json:"occupied"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Comments from the requester.
	ReqComment string `json:"reqComment"`
	// Initials of the person requesting the use of this slot. Use SYSTEM if this
	// request is auto-generated by an auto-planning system.
	ReqInitials string `json:"reqInitials"`
	// Short name of the organization requesting use of this slot.
	ReqOrg string `json:"reqOrg"`
	// The aircraft Model Design Series designation of the aircraft this slot is
	// reserved for.
	ResAircraftMds string `json:"resAircraftMDS"`
	// Mission identifier reserving this slot according to Mobility Air Forces (MAF)
	// Encode/Decode procedures.
	ResMissionID string `json:"resMissionId"`
	// The reason the slot reservation was made.
	ResReason string `json:"resReason"`
	// The tail number of the aircraft this slot is reserved for.
	ResTailNumber string `json:"resTailNumber"`
	// Indicates the type of reservation (e.g. M for Mission, A for Aircraft, O for
	// Other).
	ResType string `json:"resType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Current status of this slot (REQUESTED / APPROVED / DENIED / BLOCKED / OTHER).
	//
	// Any of "REQUESTED", "APPROVED", "DENIED", "BLOCKED", "OTHER".
	Status AirfieldslotconsumptionFullStatus `json:"status"`
	// The desired time for aircraft action such as landing, take off, parking, etc.,
	// in ISO 8601 UTC format.
	TargetTime time.Time `json:"targetTime" format:"date-time"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDAirfieldSlot        respjson.Field
		NumAircraft           respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		AltArrSortieID        respjson.Field
		AltDepSortieID        respjson.Field
		AppComment            respjson.Field
		AppInitials           respjson.Field
		AppOrg                respjson.Field
		CallSigns             respjson.Field
		Consumer              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndTime               respjson.Field
		IDArrSortie           respjson.Field
		IDDepSortie           respjson.Field
		MissionID             respjson.Field
		OccAircraftMds        respjson.Field
		OccStartTime          respjson.Field
		OccTailNumber         respjson.Field
		Occupied              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ReqComment            respjson.Field
		ReqInitials           respjson.Field
		ReqOrg                respjson.Field
		ResAircraftMds        respjson.Field
		ResMissionID          respjson.Field
		ResReason             respjson.Field
		ResTailNumber         respjson.Field
		ResType               respjson.Field
		SourceDl              respjson.Field
		Status                respjson.Field
		TargetTime            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirfieldslotconsumptionFull) RawJSON() string { return r.JSON.raw }
func (r *AirfieldslotconsumptionFull) UnmarshalJSON(data []byte) error {
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
type AirfieldslotconsumptionFullDataMode string

const (
	AirfieldslotconsumptionFullDataModeReal      AirfieldslotconsumptionFullDataMode = "REAL"
	AirfieldslotconsumptionFullDataModeTest      AirfieldslotconsumptionFullDataMode = "TEST"
	AirfieldslotconsumptionFullDataModeSimulated AirfieldslotconsumptionFullDataMode = "SIMULATED"
	AirfieldslotconsumptionFullDataModeExercise  AirfieldslotconsumptionFullDataMode = "EXERCISE"
)

// Current status of this slot (REQUESTED / APPROVED / DENIED / BLOCKED / OTHER).
type AirfieldslotconsumptionFullStatus string

const (
	AirfieldslotconsumptionFullStatusRequested AirfieldslotconsumptionFullStatus = "REQUESTED"
	AirfieldslotconsumptionFullStatusApproved  AirfieldslotconsumptionFullStatus = "APPROVED"
	AirfieldslotconsumptionFullStatusDenied    AirfieldslotconsumptionFullStatus = "DENIED"
	AirfieldslotconsumptionFullStatusBlocked   AirfieldslotconsumptionFullStatus = "BLOCKED"
	AirfieldslotconsumptionFullStatusOther     AirfieldslotconsumptionFullStatus = "OTHER"
)

type AirfieldSlotConsumptionQueryhelpResponse struct {
	AodrSupported         bool                                                `json:"aodrSupported"`
	ClassificationMarking string                                              `json:"classificationMarking"`
	Description           string                                              `json:"description"`
	HistorySupported      bool                                                `json:"historySupported"`
	Name                  string                                              `json:"name"`
	Parameters            []AirfieldSlotConsumptionQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                            `json:"requiredRoles"`
	RestSupported         bool                                                `json:"restSupported"`
	SortSupported         bool                                                `json:"sortSupported"`
	TypeName              string                                              `json:"typeName"`
	Uri                   string                                              `json:"uri"`
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
func (r AirfieldSlotConsumptionQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *AirfieldSlotConsumptionQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirfieldSlotConsumptionQueryhelpResponseParameter struct {
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
func (r AirfieldSlotConsumptionQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *AirfieldSlotConsumptionQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirfieldSlotConsumptionNewParams struct {
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
	DataMode AirfieldSlotConsumptionNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the airfield slot for which this slot consumption record is
	// referencing.
	IDAirfieldSlot string `json:"idAirfieldSlot,required"`
	// Number of aircraft using this slot for this time.
	NumAircraft int64 `json:"numAircraft,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the slot window, in ISO 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate identifier of the sortie arriving at the slot start time provided by
	// the source.
	AltArrSortieID param.Opt[string] `json:"altArrSortieId,omitzero"`
	// Alternate identifier of the sortie departing at the slot end time provided by
	// the source.
	AltDepSortieID param.Opt[string] `json:"altDepSortieId,omitzero"`
	// Comments from the approver.
	AppComment param.Opt[string] `json:"appComment,omitzero"`
	// Initials of the person approving the use of this slot. Use SYSTEM if
	// auto-approved without human involvement.
	AppInitials param.Opt[string] `json:"appInitials,omitzero"`
	// Short name of the organization approving the use of this slot.
	AppOrg param.Opt[string] `json:"appOrg,omitzero"`
	// Identifying name of the aircraft using this slot. Names are often Prior
	// Permission Required (PPR) numbers or other similar human-readable identifiers.
	Consumer param.Opt[string] `json:"consumer,omitzero"`
	// The end of the slot window, in ISO 8601 UTC format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Unique identifier of the sortie arriving at the slot start time.
	IDArrSortie param.Opt[string] `json:"idArrSortie,omitzero"`
	// Unique identifier of the sortie departing at the slot end time.
	IDDepSortie param.Opt[string] `json:"idDepSortie,omitzero"`
	// Mission identifier using this slot according to Mobility Air Forces (MAF)
	// Encode/Decode procedures.
	MissionID param.Opt[string] `json:"missionId,omitzero"`
	// The aircraft Model Design Series designation of the aircraft occupying this
	// slot.
	OccAircraftMds param.Opt[string] `json:"occAircraftMDS,omitzero"`
	// Time the aircraft began occupying this slot, in ISO 8601 UTC format with
	// millisecond precision.
	OccStartTime param.Opt[time.Time] `json:"occStartTime,omitzero" format:"date-time"`
	// The tail number of the aircraft occupying this slot.
	OccTailNumber param.Opt[string] `json:"occTailNumber,omitzero"`
	// Flag indicating if the slot is occupied.
	Occupied param.Opt[bool] `json:"occupied,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Comments from the requester.
	ReqComment param.Opt[string] `json:"reqComment,omitzero"`
	// Initials of the person requesting the use of this slot. Use SYSTEM if this
	// request is auto-generated by an auto-planning system.
	ReqInitials param.Opt[string] `json:"reqInitials,omitzero"`
	// Short name of the organization requesting use of this slot.
	ReqOrg param.Opt[string] `json:"reqOrg,omitzero"`
	// The aircraft Model Design Series designation of the aircraft this slot is
	// reserved for.
	ResAircraftMds param.Opt[string] `json:"resAircraftMDS,omitzero"`
	// Mission identifier reserving this slot according to Mobility Air Forces (MAF)
	// Encode/Decode procedures.
	ResMissionID param.Opt[string] `json:"resMissionId,omitzero"`
	// The reason the slot reservation was made.
	ResReason param.Opt[string] `json:"resReason,omitzero"`
	// The tail number of the aircraft this slot is reserved for.
	ResTailNumber param.Opt[string] `json:"resTailNumber,omitzero"`
	// Indicates the type of reservation (e.g. M for Mission, A for Aircraft, O for
	// Other).
	ResType param.Opt[string] `json:"resType,omitzero"`
	// The desired time for aircraft action such as landing, take off, parking, etc.,
	// in ISO 8601 UTC format.
	TargetTime param.Opt[time.Time] `json:"targetTime,omitzero" format:"date-time"`
	// Array of call signs of the aircraft using this slot.
	CallSigns []string `json:"callSigns,omitzero"`
	// Current status of this slot (REQUESTED / APPROVED / DENIED / BLOCKED / OTHER).
	//
	// Any of "REQUESTED", "APPROVED", "DENIED", "BLOCKED", "OTHER".
	Status AirfieldSlotConsumptionNewParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r AirfieldSlotConsumptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AirfieldSlotConsumptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirfieldSlotConsumptionNewParams) UnmarshalJSON(data []byte) error {
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
type AirfieldSlotConsumptionNewParamsDataMode string

const (
	AirfieldSlotConsumptionNewParamsDataModeReal      AirfieldSlotConsumptionNewParamsDataMode = "REAL"
	AirfieldSlotConsumptionNewParamsDataModeTest      AirfieldSlotConsumptionNewParamsDataMode = "TEST"
	AirfieldSlotConsumptionNewParamsDataModeSimulated AirfieldSlotConsumptionNewParamsDataMode = "SIMULATED"
	AirfieldSlotConsumptionNewParamsDataModeExercise  AirfieldSlotConsumptionNewParamsDataMode = "EXERCISE"
)

// Current status of this slot (REQUESTED / APPROVED / DENIED / BLOCKED / OTHER).
type AirfieldSlotConsumptionNewParamsStatus string

const (
	AirfieldSlotConsumptionNewParamsStatusRequested AirfieldSlotConsumptionNewParamsStatus = "REQUESTED"
	AirfieldSlotConsumptionNewParamsStatusApproved  AirfieldSlotConsumptionNewParamsStatus = "APPROVED"
	AirfieldSlotConsumptionNewParamsStatusDenied    AirfieldSlotConsumptionNewParamsStatus = "DENIED"
	AirfieldSlotConsumptionNewParamsStatusBlocked   AirfieldSlotConsumptionNewParamsStatus = "BLOCKED"
	AirfieldSlotConsumptionNewParamsStatusOther     AirfieldSlotConsumptionNewParamsStatus = "OTHER"
)

type AirfieldSlotConsumptionGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirfieldSlotConsumptionGetParams]'s query parameters as
// `url.Values`.
func (r AirfieldSlotConsumptionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldSlotConsumptionUpdateParams struct {
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
	DataMode AirfieldSlotConsumptionUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the airfield slot for which this slot consumption record is
	// referencing.
	IDAirfieldSlot string `json:"idAirfieldSlot,required"`
	// Number of aircraft using this slot for this time.
	NumAircraft int64 `json:"numAircraft,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the slot window, in ISO 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate identifier of the sortie arriving at the slot start time provided by
	// the source.
	AltArrSortieID param.Opt[string] `json:"altArrSortieId,omitzero"`
	// Alternate identifier of the sortie departing at the slot end time provided by
	// the source.
	AltDepSortieID param.Opt[string] `json:"altDepSortieId,omitzero"`
	// Comments from the approver.
	AppComment param.Opt[string] `json:"appComment,omitzero"`
	// Initials of the person approving the use of this slot. Use SYSTEM if
	// auto-approved without human involvement.
	AppInitials param.Opt[string] `json:"appInitials,omitzero"`
	// Short name of the organization approving the use of this slot.
	AppOrg param.Opt[string] `json:"appOrg,omitzero"`
	// Identifying name of the aircraft using this slot. Names are often Prior
	// Permission Required (PPR) numbers or other similar human-readable identifiers.
	Consumer param.Opt[string] `json:"consumer,omitzero"`
	// The end of the slot window, in ISO 8601 UTC format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Unique identifier of the sortie arriving at the slot start time.
	IDArrSortie param.Opt[string] `json:"idArrSortie,omitzero"`
	// Unique identifier of the sortie departing at the slot end time.
	IDDepSortie param.Opt[string] `json:"idDepSortie,omitzero"`
	// Mission identifier using this slot according to Mobility Air Forces (MAF)
	// Encode/Decode procedures.
	MissionID param.Opt[string] `json:"missionId,omitzero"`
	// The aircraft Model Design Series designation of the aircraft occupying this
	// slot.
	OccAircraftMds param.Opt[string] `json:"occAircraftMDS,omitzero"`
	// Time the aircraft began occupying this slot, in ISO 8601 UTC format with
	// millisecond precision.
	OccStartTime param.Opt[time.Time] `json:"occStartTime,omitzero" format:"date-time"`
	// The tail number of the aircraft occupying this slot.
	OccTailNumber param.Opt[string] `json:"occTailNumber,omitzero"`
	// Flag indicating if the slot is occupied.
	Occupied param.Opt[bool] `json:"occupied,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Comments from the requester.
	ReqComment param.Opt[string] `json:"reqComment,omitzero"`
	// Initials of the person requesting the use of this slot. Use SYSTEM if this
	// request is auto-generated by an auto-planning system.
	ReqInitials param.Opt[string] `json:"reqInitials,omitzero"`
	// Short name of the organization requesting use of this slot.
	ReqOrg param.Opt[string] `json:"reqOrg,omitzero"`
	// The aircraft Model Design Series designation of the aircraft this slot is
	// reserved for.
	ResAircraftMds param.Opt[string] `json:"resAircraftMDS,omitzero"`
	// Mission identifier reserving this slot according to Mobility Air Forces (MAF)
	// Encode/Decode procedures.
	ResMissionID param.Opt[string] `json:"resMissionId,omitzero"`
	// The reason the slot reservation was made.
	ResReason param.Opt[string] `json:"resReason,omitzero"`
	// The tail number of the aircraft this slot is reserved for.
	ResTailNumber param.Opt[string] `json:"resTailNumber,omitzero"`
	// Indicates the type of reservation (e.g. M for Mission, A for Aircraft, O for
	// Other).
	ResType param.Opt[string] `json:"resType,omitzero"`
	// The desired time for aircraft action such as landing, take off, parking, etc.,
	// in ISO 8601 UTC format.
	TargetTime param.Opt[time.Time] `json:"targetTime,omitzero" format:"date-time"`
	// Array of call signs of the aircraft using this slot.
	CallSigns []string `json:"callSigns,omitzero"`
	// Current status of this slot (REQUESTED / APPROVED / DENIED / BLOCKED / OTHER).
	//
	// Any of "REQUESTED", "APPROVED", "DENIED", "BLOCKED", "OTHER".
	Status AirfieldSlotConsumptionUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r AirfieldSlotConsumptionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AirfieldSlotConsumptionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirfieldSlotConsumptionUpdateParams) UnmarshalJSON(data []byte) error {
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
type AirfieldSlotConsumptionUpdateParamsDataMode string

const (
	AirfieldSlotConsumptionUpdateParamsDataModeReal      AirfieldSlotConsumptionUpdateParamsDataMode = "REAL"
	AirfieldSlotConsumptionUpdateParamsDataModeTest      AirfieldSlotConsumptionUpdateParamsDataMode = "TEST"
	AirfieldSlotConsumptionUpdateParamsDataModeSimulated AirfieldSlotConsumptionUpdateParamsDataMode = "SIMULATED"
	AirfieldSlotConsumptionUpdateParamsDataModeExercise  AirfieldSlotConsumptionUpdateParamsDataMode = "EXERCISE"
)

// Current status of this slot (REQUESTED / APPROVED / DENIED / BLOCKED / OTHER).
type AirfieldSlotConsumptionUpdateParamsStatus string

const (
	AirfieldSlotConsumptionUpdateParamsStatusRequested AirfieldSlotConsumptionUpdateParamsStatus = "REQUESTED"
	AirfieldSlotConsumptionUpdateParamsStatusApproved  AirfieldSlotConsumptionUpdateParamsStatus = "APPROVED"
	AirfieldSlotConsumptionUpdateParamsStatusDenied    AirfieldSlotConsumptionUpdateParamsStatus = "DENIED"
	AirfieldSlotConsumptionUpdateParamsStatusBlocked   AirfieldSlotConsumptionUpdateParamsStatus = "BLOCKED"
	AirfieldSlotConsumptionUpdateParamsStatusOther     AirfieldSlotConsumptionUpdateParamsStatus = "OTHER"
)

type AirfieldSlotConsumptionListParams struct {
	// The start of the slot window, in ISO 8601 UTC format. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirfieldSlotConsumptionListParams]'s query parameters as
// `url.Values`.
func (r AirfieldSlotConsumptionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldSlotConsumptionCountParams struct {
	// The start of the slot window, in ISO 8601 UTC format. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirfieldSlotConsumptionCountParams]'s query parameters as
// `url.Values`.
func (r AirfieldSlotConsumptionCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldSlotConsumptionTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The start of the slot window, in ISO 8601 UTC format. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirfieldSlotConsumptionTupleParams]'s query parameters as
// `url.Values`.
func (r AirfieldSlotConsumptionTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
