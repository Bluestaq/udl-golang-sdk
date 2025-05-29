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
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// OnorbiteventService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbiteventService] method instead.
type OnorbiteventService struct {
	Options []option.RequestOption
}

// NewOnorbiteventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOnorbiteventService(opts ...option.RequestOption) (r OnorbiteventService) {
	r = OnorbiteventService{}
	r.Options = opts
	return
}

// Service operation to take a single OnorbitEvent as a POST body and ingest into
// the database. An OnorbitEvent is an event associated with a particular on-orbit
// spacecraft including insurance related losses, anomalies and incidents. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *OnorbiteventService) New(ctx context.Context, body OnorbiteventNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitevent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single OnorbitEvent. An OnorbitEvent is an event
// associated with a particular on-orbit spacecraft including insurance related
// losses, anomalies and incidents. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *OnorbiteventService) Update(ctx context.Context, id string, body OnorbiteventUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitevent/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbiteventService) List(ctx context.Context, query OnorbiteventListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbiteventListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitevent"
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
func (r *OnorbiteventService) ListAutoPaging(ctx context.Context, query OnorbiteventListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbiteventListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a OnorbitEvent object specified by the passed ID
// path parameter. An OnorbitEvent is an event associated with a particular
// on-orbit spacecraft including insurance related losses, anomalies and incidents.
// A specific role is required to perform this service operation. Please contact
// the UDL team for assistance.
func (r *OnorbiteventService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitevent/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnorbiteventService) Count(ctx context.Context, query OnorbiteventCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onorbitevent/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single OnorbitEvent record by its unique ID passed as
// a path parameter. An OnorbitEvent is an event associated with a particular
// on-orbit spacecraft including insurance related losses, anomalies and incidents.
func (r *OnorbiteventService) Get(ctx context.Context, id string, query OnorbiteventGetParams, opts ...option.RequestOption) (res *OnorbiteventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitevent/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *OnorbiteventService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *OnorbiteventQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/onorbitevent/queryhelp"
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
func (r *OnorbiteventService) Tuple(ctx context.Context, query OnorbiteventTupleParams, opts ...option.RequestOption) (res *[]OnorbiteventTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/onorbitevent/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OnorbiteventListResponse struct {
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
	DataMode OnorbiteventListResponseDataMode `json:"dataMode,required"`
	// Date/Time of the event. See eventTimeNotes for remarks on the accuracy of the
	// date time.
	EventTime time.Time `json:"eventTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Achieved phase of flight prior to the event.
	AchievedFlightPhase string `json:"achievedFlightPhase"`
	// Spacecraft age at the event in years.
	AgeAtEvent float64 `json:"ageAtEvent"`
	// Spacecraft capability loss incurred, as a fraction of 1.
	CapabilityLoss float64 `json:"capabilityLoss"`
	// Notes on capability loss at the time of event.
	CapabilityLossNotes string `json:"capabilityLossNotes"`
	// Spacecraft capacity loss incurred, as a fraction of 1.
	CapacityLoss float64 `json:"capacityLoss"`
	// Additional equipment which failed as a result of faulty equipment on the
	// spacecraft during the event.
	ConsequentialEquipmentFailure string `json:"consequentialEquipmentFailure"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Notes/description of the event.
	Description string `json:"description"`
	// Equipment on the spacecraft which caused the event.
	EquipmentAtFault string `json:"equipmentAtFault"`
	// Additional notes on the equipment causing the event/loss.
	EquipmentCausingLossNotes string `json:"equipmentCausingLossNotes"`
	// Specific part of the equipment on the spacecraft which caused the event.
	EquipmentPartAtFault string `json:"equipmentPartAtFault"`
	// Type of the equipment on the spacecraft which caused the event.
	EquipmentTypeAtFault string `json:"equipmentTypeAtFault"`
	// The result of the reported event.
	EventResult string `json:"eventResult"`
	// Notes/remarks on the validity/accuracy of the eventTime.
	EventTimeNotes string `json:"eventTimeNotes"`
	// The type of on-orbit event being reported.
	EventType string `json:"eventType"`
	// GEO position longitude at event time if applicable. Negative values are west.
	GeoPosition float64 `json:"geoPosition"`
	// Unique identifier of the on-orbit object for this event.
	IDOnOrbit string `json:"idOnOrbit"`
	// Boolean indicating if the spacecraft is inclined.
	Inclined bool `json:"inclined"`
	// Number of humans injured in the event.
	Injured int64 `json:"injured"`
	// Additional insurance notes on coverages at the time of event.
	InsuranceCarriedNotes string `json:"insuranceCarriedNotes"`
	// Insurance loss incurred, as a fraction of 1.
	InsuranceLoss float64 `json:"insuranceLoss"`
	// Additional insurance notes if the event is an official loss.
	InsuranceLossNotes string `json:"insuranceLossNotes"`
	// Number of humans killed in the event.
	Killed int64 `json:"killed"`
	// Unique identifier of the organization which leases this on-orbit spacecraft.
	LesseeOrgID string `json:"lesseeOrgId"`
	// Spacecraft life lost due to the event as a percent/fraction of 1.
	LifeLost float64 `json:"lifeLost"`
	// Net amount of the insurance claim for the event, in USD.
	NetAmount float64 `json:"netAmount"`
	// The status of the on-orbit object.
	ObjectStatus string `json:"objectStatus"`
	// Phase of flight during which the event occurred.
	OccurrenceFlightPhase string `json:"occurrenceFlightPhase"`
	// Date time of official loss of the spacecraft.
	OfficialLossDate time.Time `json:"officialLossDate" format:"date-time"`
	// Unique identifier of the organization on whose behalf the on-orbit spacecraft is
	// operated.
	OperatedOnBehalfOfOrgID string `json:"operatedOnBehalfOfOrgId"`
	// Organization ID of the operator of the on-orbit spacecraft at the time of the
	// event.
	OperatorOrgID string `json:"operatorOrgId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Original object ID or Catalog Number provided by source (may not map to an
	// existing idOnOrbit in UDL).
	OrigObjectID string `json:"origObjectId"`
	// Organization ID of the owner of the on-orbit spacecraft at the time of the
	// event.
	OwnerOrgID string `json:"ownerOrgId"`
	// GEO slot plane number/designator of the spacecraft at event time.
	PlaneNumber string `json:"planeNumber"`
	// GEO plane slot of the spacecraft at event time.
	PlaneSlot string `json:"planeSlot"`
	// Position status of the spacecraft at event time (e.g. Stable, Drifting/Tumbling,
	// etc).
	PositionStatus string `json:"positionStatus"`
	// Additional remarks on the event description.
	Remarks string `json:"remarks"`
	// Description of the satellite orbital position or regime.
	SatellitePosition string `json:"satellitePosition"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Faulty stage of flight for the event.
	StageAtFault string `json:"stageAtFault"`
	// Insurance loss incurred by 3rd party insurance, in USD.
	ThirdPartyInsuranceLoss float64 `json:"thirdPartyInsuranceLoss"`
	// Underlying cause of the event.
	UnderlyingCause string `json:"underlyingCause"`
	// Maximum validity time of the event.
	UntilTime time.Time `json:"untilTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking         respjson.Field
		DataMode                      respjson.Field
		EventTime                     respjson.Field
		Source                        respjson.Field
		ID                            respjson.Field
		AchievedFlightPhase           respjson.Field
		AgeAtEvent                    respjson.Field
		CapabilityLoss                respjson.Field
		CapabilityLossNotes           respjson.Field
		CapacityLoss                  respjson.Field
		ConsequentialEquipmentFailure respjson.Field
		CreatedAt                     respjson.Field
		CreatedBy                     respjson.Field
		DeclassificationDate          respjson.Field
		DeclassificationString        respjson.Field
		DerivedFrom                   respjson.Field
		Description                   respjson.Field
		EquipmentAtFault              respjson.Field
		EquipmentCausingLossNotes     respjson.Field
		EquipmentPartAtFault          respjson.Field
		EquipmentTypeAtFault          respjson.Field
		EventResult                   respjson.Field
		EventTimeNotes                respjson.Field
		EventType                     respjson.Field
		GeoPosition                   respjson.Field
		IDOnOrbit                     respjson.Field
		Inclined                      respjson.Field
		Injured                       respjson.Field
		InsuranceCarriedNotes         respjson.Field
		InsuranceLoss                 respjson.Field
		InsuranceLossNotes            respjson.Field
		Killed                        respjson.Field
		LesseeOrgID                   respjson.Field
		LifeLost                      respjson.Field
		NetAmount                     respjson.Field
		ObjectStatus                  respjson.Field
		OccurrenceFlightPhase         respjson.Field
		OfficialLossDate              respjson.Field
		OperatedOnBehalfOfOrgID       respjson.Field
		OperatorOrgID                 respjson.Field
		Origin                        respjson.Field
		OrigNetwork                   respjson.Field
		OrigObjectID                  respjson.Field
		OwnerOrgID                    respjson.Field
		PlaneNumber                   respjson.Field
		PlaneSlot                     respjson.Field
		PositionStatus                respjson.Field
		Remarks                       respjson.Field
		SatellitePosition             respjson.Field
		SatNo                         respjson.Field
		StageAtFault                  respjson.Field
		ThirdPartyInsuranceLoss       respjson.Field
		UnderlyingCause               respjson.Field
		UntilTime                     respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbiteventListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbiteventListResponse) UnmarshalJSON(data []byte) error {
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
type OnorbiteventListResponseDataMode string

const (
	OnorbiteventListResponseDataModeReal      OnorbiteventListResponseDataMode = "REAL"
	OnorbiteventListResponseDataModeTest      OnorbiteventListResponseDataMode = "TEST"
	OnorbiteventListResponseDataModeSimulated OnorbiteventListResponseDataMode = "SIMULATED"
	OnorbiteventListResponseDataModeExercise  OnorbiteventListResponseDataMode = "EXERCISE"
)

type OnorbiteventGetResponse struct {
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
	DataMode OnorbiteventGetResponseDataMode `json:"dataMode,required"`
	// Date/Time of the event. See eventTimeNotes for remarks on the accuracy of the
	// date time.
	EventTime time.Time `json:"eventTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Achieved phase of flight prior to the event.
	AchievedFlightPhase string `json:"achievedFlightPhase"`
	// Spacecraft age at the event in years.
	AgeAtEvent float64 `json:"ageAtEvent"`
	// Spacecraft capability loss incurred, as a fraction of 1.
	CapabilityLoss float64 `json:"capabilityLoss"`
	// Notes on capability loss at the time of event.
	CapabilityLossNotes string `json:"capabilityLossNotes"`
	// Spacecraft capacity loss incurred, as a fraction of 1.
	CapacityLoss float64 `json:"capacityLoss"`
	// Additional equipment which failed as a result of faulty equipment on the
	// spacecraft during the event.
	ConsequentialEquipmentFailure string `json:"consequentialEquipmentFailure"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Notes/description of the event.
	Description string `json:"description"`
	// Equipment on the spacecraft which caused the event.
	EquipmentAtFault string `json:"equipmentAtFault"`
	// Additional notes on the equipment causing the event/loss.
	EquipmentCausingLossNotes string `json:"equipmentCausingLossNotes"`
	// Specific part of the equipment on the spacecraft which caused the event.
	EquipmentPartAtFault string `json:"equipmentPartAtFault"`
	// Type of the equipment on the spacecraft which caused the event.
	EquipmentTypeAtFault string `json:"equipmentTypeAtFault"`
	// The result of the reported event.
	EventResult string `json:"eventResult"`
	// Notes/remarks on the validity/accuracy of the eventTime.
	EventTimeNotes string `json:"eventTimeNotes"`
	// The type of on-orbit event being reported.
	EventType string `json:"eventType"`
	// GEO position longitude at event time if applicable. Negative values are west.
	GeoPosition float64 `json:"geoPosition"`
	// Unique identifier of the on-orbit object for this event.
	IDOnOrbit string `json:"idOnOrbit"`
	// Boolean indicating if the spacecraft is inclined.
	Inclined bool `json:"inclined"`
	// Number of humans injured in the event.
	Injured int64 `json:"injured"`
	// Additional insurance notes on coverages at the time of event.
	InsuranceCarriedNotes string `json:"insuranceCarriedNotes"`
	// Insurance loss incurred, as a fraction of 1.
	InsuranceLoss float64 `json:"insuranceLoss"`
	// Additional insurance notes if the event is an official loss.
	InsuranceLossNotes string `json:"insuranceLossNotes"`
	// Number of humans killed in the event.
	Killed int64 `json:"killed"`
	// Unique identifier of the organization which leases this on-orbit spacecraft.
	LesseeOrgID string `json:"lesseeOrgId"`
	// Spacecraft life lost due to the event as a percent/fraction of 1.
	LifeLost float64 `json:"lifeLost"`
	// Net amount of the insurance claim for the event, in USD.
	NetAmount float64 `json:"netAmount"`
	// The status of the on-orbit object.
	ObjectStatus string `json:"objectStatus"`
	// Phase of flight during which the event occurred.
	OccurrenceFlightPhase string `json:"occurrenceFlightPhase"`
	// Date time of official loss of the spacecraft.
	OfficialLossDate time.Time `json:"officialLossDate" format:"date-time"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Unique identifier of the organization on whose behalf the on-orbit spacecraft is
	// operated.
	OperatedOnBehalfOfOrgID string `json:"operatedOnBehalfOfOrgId"`
	// Organization ID of the operator of the on-orbit spacecraft at the time of the
	// event.
	OperatorOrgID string `json:"operatorOrgId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Original object ID or Catalog Number provided by source (may not map to an
	// existing idOnOrbit in UDL).
	OrigObjectID string `json:"origObjectId"`
	// Organization ID of the owner of the on-orbit spacecraft at the time of the
	// event.
	OwnerOrgID string `json:"ownerOrgId"`
	// GEO slot plane number/designator of the spacecraft at event time.
	PlaneNumber string `json:"planeNumber"`
	// GEO plane slot of the spacecraft at event time.
	PlaneSlot string `json:"planeSlot"`
	// Position status of the spacecraft at event time (e.g. Stable, Drifting/Tumbling,
	// etc).
	PositionStatus string `json:"positionStatus"`
	// Additional remarks on the event description.
	Remarks string `json:"remarks"`
	// Description of the satellite orbital position or regime.
	SatellitePosition string `json:"satellitePosition"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Faulty stage of flight for the event.
	StageAtFault string `json:"stageAtFault"`
	// Insurance loss incurred by 3rd party insurance, in USD.
	ThirdPartyInsuranceLoss float64 `json:"thirdPartyInsuranceLoss"`
	// Underlying cause of the event.
	UnderlyingCause string `json:"underlyingCause"`
	// Maximum validity time of the event.
	UntilTime time.Time `json:"untilTime" format:"date-time"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking         respjson.Field
		DataMode                      respjson.Field
		EventTime                     respjson.Field
		Source                        respjson.Field
		ID                            respjson.Field
		AchievedFlightPhase           respjson.Field
		AgeAtEvent                    respjson.Field
		CapabilityLoss                respjson.Field
		CapabilityLossNotes           respjson.Field
		CapacityLoss                  respjson.Field
		ConsequentialEquipmentFailure respjson.Field
		CreatedAt                     respjson.Field
		CreatedBy                     respjson.Field
		DeclassificationDate          respjson.Field
		DeclassificationString        respjson.Field
		DerivedFrom                   respjson.Field
		Description                   respjson.Field
		EquipmentAtFault              respjson.Field
		EquipmentCausingLossNotes     respjson.Field
		EquipmentPartAtFault          respjson.Field
		EquipmentTypeAtFault          respjson.Field
		EventResult                   respjson.Field
		EventTimeNotes                respjson.Field
		EventType                     respjson.Field
		GeoPosition                   respjson.Field
		IDOnOrbit                     respjson.Field
		Inclined                      respjson.Field
		Injured                       respjson.Field
		InsuranceCarriedNotes         respjson.Field
		InsuranceLoss                 respjson.Field
		InsuranceLossNotes            respjson.Field
		Killed                        respjson.Field
		LesseeOrgID                   respjson.Field
		LifeLost                      respjson.Field
		NetAmount                     respjson.Field
		ObjectStatus                  respjson.Field
		OccurrenceFlightPhase         respjson.Field
		OfficialLossDate              respjson.Field
		OnOrbit                       respjson.Field
		OperatedOnBehalfOfOrgID       respjson.Field
		OperatorOrgID                 respjson.Field
		Origin                        respjson.Field
		OrigNetwork                   respjson.Field
		OrigObjectID                  respjson.Field
		OwnerOrgID                    respjson.Field
		PlaneNumber                   respjson.Field
		PlaneSlot                     respjson.Field
		PositionStatus                respjson.Field
		Remarks                       respjson.Field
		SatellitePosition             respjson.Field
		SatNo                         respjson.Field
		StageAtFault                  respjson.Field
		ThirdPartyInsuranceLoss       respjson.Field
		UnderlyingCause               respjson.Field
		UntilTime                     respjson.Field
		UpdatedAt                     respjson.Field
		UpdatedBy                     respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbiteventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbiteventGetResponse) UnmarshalJSON(data []byte) error {
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
type OnorbiteventGetResponseDataMode string

const (
	OnorbiteventGetResponseDataModeReal      OnorbiteventGetResponseDataMode = "REAL"
	OnorbiteventGetResponseDataModeTest      OnorbiteventGetResponseDataMode = "TEST"
	OnorbiteventGetResponseDataModeSimulated OnorbiteventGetResponseDataMode = "SIMULATED"
	OnorbiteventGetResponseDataModeExercise  OnorbiteventGetResponseDataMode = "EXERCISE"
)

type OnorbiteventQueryhelpResponse struct {
	AodrSupported         bool                                     `json:"aodrSupported"`
	ClassificationMarking string                                   `json:"classificationMarking"`
	Description           string                                   `json:"description"`
	HistorySupported      bool                                     `json:"historySupported"`
	Name                  string                                   `json:"name"`
	Parameters            []OnorbiteventQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                 `json:"requiredRoles"`
	RestSupported         bool                                     `json:"restSupported"`
	SortSupported         bool                                     `json:"sortSupported"`
	TypeName              string                                   `json:"typeName"`
	Uri                   string                                   `json:"uri"`
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
func (r OnorbiteventQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbiteventQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbiteventQueryhelpResponseParameter struct {
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
func (r OnorbiteventQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *OnorbiteventQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbiteventTupleResponse struct {
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
	DataMode OnorbiteventTupleResponseDataMode `json:"dataMode,required"`
	// Date/Time of the event. See eventTimeNotes for remarks on the accuracy of the
	// date time.
	EventTime time.Time `json:"eventTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Achieved phase of flight prior to the event.
	AchievedFlightPhase string `json:"achievedFlightPhase"`
	// Spacecraft age at the event in years.
	AgeAtEvent float64 `json:"ageAtEvent"`
	// Spacecraft capability loss incurred, as a fraction of 1.
	CapabilityLoss float64 `json:"capabilityLoss"`
	// Notes on capability loss at the time of event.
	CapabilityLossNotes string `json:"capabilityLossNotes"`
	// Spacecraft capacity loss incurred, as a fraction of 1.
	CapacityLoss float64 `json:"capacityLoss"`
	// Additional equipment which failed as a result of faulty equipment on the
	// spacecraft during the event.
	ConsequentialEquipmentFailure string `json:"consequentialEquipmentFailure"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Notes/description of the event.
	Description string `json:"description"`
	// Equipment on the spacecraft which caused the event.
	EquipmentAtFault string `json:"equipmentAtFault"`
	// Additional notes on the equipment causing the event/loss.
	EquipmentCausingLossNotes string `json:"equipmentCausingLossNotes"`
	// Specific part of the equipment on the spacecraft which caused the event.
	EquipmentPartAtFault string `json:"equipmentPartAtFault"`
	// Type of the equipment on the spacecraft which caused the event.
	EquipmentTypeAtFault string `json:"equipmentTypeAtFault"`
	// The result of the reported event.
	EventResult string `json:"eventResult"`
	// Notes/remarks on the validity/accuracy of the eventTime.
	EventTimeNotes string `json:"eventTimeNotes"`
	// The type of on-orbit event being reported.
	EventType string `json:"eventType"`
	// GEO position longitude at event time if applicable. Negative values are west.
	GeoPosition float64 `json:"geoPosition"`
	// Unique identifier of the on-orbit object for this event.
	IDOnOrbit string `json:"idOnOrbit"`
	// Boolean indicating if the spacecraft is inclined.
	Inclined bool `json:"inclined"`
	// Number of humans injured in the event.
	Injured int64 `json:"injured"`
	// Additional insurance notes on coverages at the time of event.
	InsuranceCarriedNotes string `json:"insuranceCarriedNotes"`
	// Insurance loss incurred, as a fraction of 1.
	InsuranceLoss float64 `json:"insuranceLoss"`
	// Additional insurance notes if the event is an official loss.
	InsuranceLossNotes string `json:"insuranceLossNotes"`
	// Number of humans killed in the event.
	Killed int64 `json:"killed"`
	// Unique identifier of the organization which leases this on-orbit spacecraft.
	LesseeOrgID string `json:"lesseeOrgId"`
	// Spacecraft life lost due to the event as a percent/fraction of 1.
	LifeLost float64 `json:"lifeLost"`
	// Net amount of the insurance claim for the event, in USD.
	NetAmount float64 `json:"netAmount"`
	// The status of the on-orbit object.
	ObjectStatus string `json:"objectStatus"`
	// Phase of flight during which the event occurred.
	OccurrenceFlightPhase string `json:"occurrenceFlightPhase"`
	// Date time of official loss of the spacecraft.
	OfficialLossDate time.Time `json:"officialLossDate" format:"date-time"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Unique identifier of the organization on whose behalf the on-orbit spacecraft is
	// operated.
	OperatedOnBehalfOfOrgID string `json:"operatedOnBehalfOfOrgId"`
	// Organization ID of the operator of the on-orbit spacecraft at the time of the
	// event.
	OperatorOrgID string `json:"operatorOrgId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Original object ID or Catalog Number provided by source (may not map to an
	// existing idOnOrbit in UDL).
	OrigObjectID string `json:"origObjectId"`
	// Organization ID of the owner of the on-orbit spacecraft at the time of the
	// event.
	OwnerOrgID string `json:"ownerOrgId"`
	// GEO slot plane number/designator of the spacecraft at event time.
	PlaneNumber string `json:"planeNumber"`
	// GEO plane slot of the spacecraft at event time.
	PlaneSlot string `json:"planeSlot"`
	// Position status of the spacecraft at event time (e.g. Stable, Drifting/Tumbling,
	// etc).
	PositionStatus string `json:"positionStatus"`
	// Additional remarks on the event description.
	Remarks string `json:"remarks"`
	// Description of the satellite orbital position or regime.
	SatellitePosition string `json:"satellitePosition"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Faulty stage of flight for the event.
	StageAtFault string `json:"stageAtFault"`
	// Insurance loss incurred by 3rd party insurance, in USD.
	ThirdPartyInsuranceLoss float64 `json:"thirdPartyInsuranceLoss"`
	// Underlying cause of the event.
	UnderlyingCause string `json:"underlyingCause"`
	// Maximum validity time of the event.
	UntilTime time.Time `json:"untilTime" format:"date-time"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking         respjson.Field
		DataMode                      respjson.Field
		EventTime                     respjson.Field
		Source                        respjson.Field
		ID                            respjson.Field
		AchievedFlightPhase           respjson.Field
		AgeAtEvent                    respjson.Field
		CapabilityLoss                respjson.Field
		CapabilityLossNotes           respjson.Field
		CapacityLoss                  respjson.Field
		ConsequentialEquipmentFailure respjson.Field
		CreatedAt                     respjson.Field
		CreatedBy                     respjson.Field
		DeclassificationDate          respjson.Field
		DeclassificationString        respjson.Field
		DerivedFrom                   respjson.Field
		Description                   respjson.Field
		EquipmentAtFault              respjson.Field
		EquipmentCausingLossNotes     respjson.Field
		EquipmentPartAtFault          respjson.Field
		EquipmentTypeAtFault          respjson.Field
		EventResult                   respjson.Field
		EventTimeNotes                respjson.Field
		EventType                     respjson.Field
		GeoPosition                   respjson.Field
		IDOnOrbit                     respjson.Field
		Inclined                      respjson.Field
		Injured                       respjson.Field
		InsuranceCarriedNotes         respjson.Field
		InsuranceLoss                 respjson.Field
		InsuranceLossNotes            respjson.Field
		Killed                        respjson.Field
		LesseeOrgID                   respjson.Field
		LifeLost                      respjson.Field
		NetAmount                     respjson.Field
		ObjectStatus                  respjson.Field
		OccurrenceFlightPhase         respjson.Field
		OfficialLossDate              respjson.Field
		OnOrbit                       respjson.Field
		OperatedOnBehalfOfOrgID       respjson.Field
		OperatorOrgID                 respjson.Field
		Origin                        respjson.Field
		OrigNetwork                   respjson.Field
		OrigObjectID                  respjson.Field
		OwnerOrgID                    respjson.Field
		PlaneNumber                   respjson.Field
		PlaneSlot                     respjson.Field
		PositionStatus                respjson.Field
		Remarks                       respjson.Field
		SatellitePosition             respjson.Field
		SatNo                         respjson.Field
		StageAtFault                  respjson.Field
		ThirdPartyInsuranceLoss       respjson.Field
		UnderlyingCause               respjson.Field
		UntilTime                     respjson.Field
		UpdatedAt                     respjson.Field
		UpdatedBy                     respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbiteventTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbiteventTupleResponse) UnmarshalJSON(data []byte) error {
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
type OnorbiteventTupleResponseDataMode string

const (
	OnorbiteventTupleResponseDataModeReal      OnorbiteventTupleResponseDataMode = "REAL"
	OnorbiteventTupleResponseDataModeTest      OnorbiteventTupleResponseDataMode = "TEST"
	OnorbiteventTupleResponseDataModeSimulated OnorbiteventTupleResponseDataMode = "SIMULATED"
	OnorbiteventTupleResponseDataModeExercise  OnorbiteventTupleResponseDataMode = "EXERCISE"
)

type OnorbiteventNewParams struct {
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
	DataMode OnorbiteventNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Date/Time of the event. See eventTimeNotes for remarks on the accuracy of the
	// date time.
	EventTime time.Time `json:"eventTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Achieved phase of flight prior to the event.
	AchievedFlightPhase param.Opt[string] `json:"achievedFlightPhase,omitzero"`
	// Spacecraft age at the event in years.
	AgeAtEvent param.Opt[float64] `json:"ageAtEvent,omitzero"`
	// Spacecraft capability loss incurred, as a fraction of 1.
	CapabilityLoss param.Opt[float64] `json:"capabilityLoss,omitzero"`
	// Notes on capability loss at the time of event.
	CapabilityLossNotes param.Opt[string] `json:"capabilityLossNotes,omitzero"`
	// Spacecraft capacity loss incurred, as a fraction of 1.
	CapacityLoss param.Opt[float64] `json:"capacityLoss,omitzero"`
	// Additional equipment which failed as a result of faulty equipment on the
	// spacecraft during the event.
	ConsequentialEquipmentFailure param.Opt[string] `json:"consequentialEquipmentFailure,omitzero"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate param.Opt[time.Time] `json:"declassificationDate,omitzero" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString param.Opt[string] `json:"declassificationString,omitzero"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom param.Opt[string] `json:"derivedFrom,omitzero"`
	// Notes/description of the event.
	Description param.Opt[string] `json:"description,omitzero"`
	// Equipment on the spacecraft which caused the event.
	EquipmentAtFault param.Opt[string] `json:"equipmentAtFault,omitzero"`
	// Additional notes on the equipment causing the event/loss.
	EquipmentCausingLossNotes param.Opt[string] `json:"equipmentCausingLossNotes,omitzero"`
	// Specific part of the equipment on the spacecraft which caused the event.
	EquipmentPartAtFault param.Opt[string] `json:"equipmentPartAtFault,omitzero"`
	// Type of the equipment on the spacecraft which caused the event.
	EquipmentTypeAtFault param.Opt[string] `json:"equipmentTypeAtFault,omitzero"`
	// The result of the reported event.
	EventResult param.Opt[string] `json:"eventResult,omitzero"`
	// Notes/remarks on the validity/accuracy of the eventTime.
	EventTimeNotes param.Opt[string] `json:"eventTimeNotes,omitzero"`
	// The type of on-orbit event being reported.
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// GEO position longitude at event time if applicable. Negative values are west.
	GeoPosition param.Opt[float64] `json:"geoPosition,omitzero"`
	// Unique identifier of the on-orbit object for this event.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Boolean indicating if the spacecraft is inclined.
	Inclined param.Opt[bool] `json:"inclined,omitzero"`
	// Number of humans injured in the event.
	Injured param.Opt[int64] `json:"injured,omitzero"`
	// Additional insurance notes on coverages at the time of event.
	InsuranceCarriedNotes param.Opt[string] `json:"insuranceCarriedNotes,omitzero"`
	// Insurance loss incurred, as a fraction of 1.
	InsuranceLoss param.Opt[float64] `json:"insuranceLoss,omitzero"`
	// Additional insurance notes if the event is an official loss.
	InsuranceLossNotes param.Opt[string] `json:"insuranceLossNotes,omitzero"`
	// Number of humans killed in the event.
	Killed param.Opt[int64] `json:"killed,omitzero"`
	// Unique identifier of the organization which leases this on-orbit spacecraft.
	LesseeOrgID param.Opt[string] `json:"lesseeOrgId,omitzero"`
	// Spacecraft life lost due to the event as a percent/fraction of 1.
	LifeLost param.Opt[float64] `json:"lifeLost,omitzero"`
	// Net amount of the insurance claim for the event, in USD.
	NetAmount param.Opt[float64] `json:"netAmount,omitzero"`
	// The status of the on-orbit object.
	ObjectStatus param.Opt[string] `json:"objectStatus,omitzero"`
	// Phase of flight during which the event occurred.
	OccurrenceFlightPhase param.Opt[string] `json:"occurrenceFlightPhase,omitzero"`
	// Date time of official loss of the spacecraft.
	OfficialLossDate param.Opt[time.Time] `json:"officialLossDate,omitzero" format:"date-time"`
	// Unique identifier of the organization on whose behalf the on-orbit spacecraft is
	// operated.
	OperatedOnBehalfOfOrgID param.Opt[string] `json:"operatedOnBehalfOfOrgId,omitzero"`
	// Organization ID of the operator of the on-orbit spacecraft at the time of the
	// event.
	OperatorOrgID param.Opt[string] `json:"operatorOrgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Original object ID or Catalog Number provided by source (may not map to an
	// existing idOnOrbit in UDL).
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Organization ID of the owner of the on-orbit spacecraft at the time of the
	// event.
	OwnerOrgID param.Opt[string] `json:"ownerOrgId,omitzero"`
	// GEO slot plane number/designator of the spacecraft at event time.
	PlaneNumber param.Opt[string] `json:"planeNumber,omitzero"`
	// GEO plane slot of the spacecraft at event time.
	PlaneSlot param.Opt[string] `json:"planeSlot,omitzero"`
	// Position status of the spacecraft at event time (e.g. Stable, Drifting/Tumbling,
	// etc).
	PositionStatus param.Opt[string] `json:"positionStatus,omitzero"`
	// Additional remarks on the event description.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// Description of the satellite orbital position or regime.
	SatellitePosition param.Opt[string] `json:"satellitePosition,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Faulty stage of flight for the event.
	StageAtFault param.Opt[string] `json:"stageAtFault,omitzero"`
	// Insurance loss incurred by 3rd party insurance, in USD.
	ThirdPartyInsuranceLoss param.Opt[float64] `json:"thirdPartyInsuranceLoss,omitzero"`
	// Underlying cause of the event.
	UnderlyingCause param.Opt[string] `json:"underlyingCause,omitzero"`
	// Maximum validity time of the event.
	UntilTime param.Opt[time.Time] `json:"untilTime,omitzero" format:"date-time"`
	paramObj
}

func (r OnorbiteventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbiteventNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbiteventNewParams) UnmarshalJSON(data []byte) error {
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
type OnorbiteventNewParamsDataMode string

const (
	OnorbiteventNewParamsDataModeReal      OnorbiteventNewParamsDataMode = "REAL"
	OnorbiteventNewParamsDataModeTest      OnorbiteventNewParamsDataMode = "TEST"
	OnorbiteventNewParamsDataModeSimulated OnorbiteventNewParamsDataMode = "SIMULATED"
	OnorbiteventNewParamsDataModeExercise  OnorbiteventNewParamsDataMode = "EXERCISE"
)

type OnorbiteventUpdateParams struct {
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
	DataMode OnorbiteventUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Date/Time of the event. See eventTimeNotes for remarks on the accuracy of the
	// date time.
	EventTime time.Time `json:"eventTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Achieved phase of flight prior to the event.
	AchievedFlightPhase param.Opt[string] `json:"achievedFlightPhase,omitzero"`
	// Spacecraft age at the event in years.
	AgeAtEvent param.Opt[float64] `json:"ageAtEvent,omitzero"`
	// Spacecraft capability loss incurred, as a fraction of 1.
	CapabilityLoss param.Opt[float64] `json:"capabilityLoss,omitzero"`
	// Notes on capability loss at the time of event.
	CapabilityLossNotes param.Opt[string] `json:"capabilityLossNotes,omitzero"`
	// Spacecraft capacity loss incurred, as a fraction of 1.
	CapacityLoss param.Opt[float64] `json:"capacityLoss,omitzero"`
	// Additional equipment which failed as a result of faulty equipment on the
	// spacecraft during the event.
	ConsequentialEquipmentFailure param.Opt[string] `json:"consequentialEquipmentFailure,omitzero"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate param.Opt[time.Time] `json:"declassificationDate,omitzero" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString param.Opt[string] `json:"declassificationString,omitzero"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom param.Opt[string] `json:"derivedFrom,omitzero"`
	// Notes/description of the event.
	Description param.Opt[string] `json:"description,omitzero"`
	// Equipment on the spacecraft which caused the event.
	EquipmentAtFault param.Opt[string] `json:"equipmentAtFault,omitzero"`
	// Additional notes on the equipment causing the event/loss.
	EquipmentCausingLossNotes param.Opt[string] `json:"equipmentCausingLossNotes,omitzero"`
	// Specific part of the equipment on the spacecraft which caused the event.
	EquipmentPartAtFault param.Opt[string] `json:"equipmentPartAtFault,omitzero"`
	// Type of the equipment on the spacecraft which caused the event.
	EquipmentTypeAtFault param.Opt[string] `json:"equipmentTypeAtFault,omitzero"`
	// The result of the reported event.
	EventResult param.Opt[string] `json:"eventResult,omitzero"`
	// Notes/remarks on the validity/accuracy of the eventTime.
	EventTimeNotes param.Opt[string] `json:"eventTimeNotes,omitzero"`
	// The type of on-orbit event being reported.
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// GEO position longitude at event time if applicable. Negative values are west.
	GeoPosition param.Opt[float64] `json:"geoPosition,omitzero"`
	// Unique identifier of the on-orbit object for this event.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Boolean indicating if the spacecraft is inclined.
	Inclined param.Opt[bool] `json:"inclined,omitzero"`
	// Number of humans injured in the event.
	Injured param.Opt[int64] `json:"injured,omitzero"`
	// Additional insurance notes on coverages at the time of event.
	InsuranceCarriedNotes param.Opt[string] `json:"insuranceCarriedNotes,omitzero"`
	// Insurance loss incurred, as a fraction of 1.
	InsuranceLoss param.Opt[float64] `json:"insuranceLoss,omitzero"`
	// Additional insurance notes if the event is an official loss.
	InsuranceLossNotes param.Opt[string] `json:"insuranceLossNotes,omitzero"`
	// Number of humans killed in the event.
	Killed param.Opt[int64] `json:"killed,omitzero"`
	// Unique identifier of the organization which leases this on-orbit spacecraft.
	LesseeOrgID param.Opt[string] `json:"lesseeOrgId,omitzero"`
	// Spacecraft life lost due to the event as a percent/fraction of 1.
	LifeLost param.Opt[float64] `json:"lifeLost,omitzero"`
	// Net amount of the insurance claim for the event, in USD.
	NetAmount param.Opt[float64] `json:"netAmount,omitzero"`
	// The status of the on-orbit object.
	ObjectStatus param.Opt[string] `json:"objectStatus,omitzero"`
	// Phase of flight during which the event occurred.
	OccurrenceFlightPhase param.Opt[string] `json:"occurrenceFlightPhase,omitzero"`
	// Date time of official loss of the spacecraft.
	OfficialLossDate param.Opt[time.Time] `json:"officialLossDate,omitzero" format:"date-time"`
	// Unique identifier of the organization on whose behalf the on-orbit spacecraft is
	// operated.
	OperatedOnBehalfOfOrgID param.Opt[string] `json:"operatedOnBehalfOfOrgId,omitzero"`
	// Organization ID of the operator of the on-orbit spacecraft at the time of the
	// event.
	OperatorOrgID param.Opt[string] `json:"operatorOrgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Original object ID or Catalog Number provided by source (may not map to an
	// existing idOnOrbit in UDL).
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Organization ID of the owner of the on-orbit spacecraft at the time of the
	// event.
	OwnerOrgID param.Opt[string] `json:"ownerOrgId,omitzero"`
	// GEO slot plane number/designator of the spacecraft at event time.
	PlaneNumber param.Opt[string] `json:"planeNumber,omitzero"`
	// GEO plane slot of the spacecraft at event time.
	PlaneSlot param.Opt[string] `json:"planeSlot,omitzero"`
	// Position status of the spacecraft at event time (e.g. Stable, Drifting/Tumbling,
	// etc).
	PositionStatus param.Opt[string] `json:"positionStatus,omitzero"`
	// Additional remarks on the event description.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// Description of the satellite orbital position or regime.
	SatellitePosition param.Opt[string] `json:"satellitePosition,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Faulty stage of flight for the event.
	StageAtFault param.Opt[string] `json:"stageAtFault,omitzero"`
	// Insurance loss incurred by 3rd party insurance, in USD.
	ThirdPartyInsuranceLoss param.Opt[float64] `json:"thirdPartyInsuranceLoss,omitzero"`
	// Underlying cause of the event.
	UnderlyingCause param.Opt[string] `json:"underlyingCause,omitzero"`
	// Maximum validity time of the event.
	UntilTime param.Opt[time.Time] `json:"untilTime,omitzero" format:"date-time"`
	paramObj
}

func (r OnorbiteventUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbiteventUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbiteventUpdateParams) UnmarshalJSON(data []byte) error {
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
type OnorbiteventUpdateParamsDataMode string

const (
	OnorbiteventUpdateParamsDataModeReal      OnorbiteventUpdateParamsDataMode = "REAL"
	OnorbiteventUpdateParamsDataModeTest      OnorbiteventUpdateParamsDataMode = "TEST"
	OnorbiteventUpdateParamsDataModeSimulated OnorbiteventUpdateParamsDataMode = "SIMULATED"
	OnorbiteventUpdateParamsDataModeExercise  OnorbiteventUpdateParamsDataMode = "EXERCISE"
)

type OnorbiteventListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbiteventListParams]'s query parameters as `url.Values`.
func (r OnorbiteventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbiteventCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbiteventCountParams]'s query parameters as
// `url.Values`.
func (r OnorbiteventCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbiteventGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbiteventGetParams]'s query parameters as `url.Values`.
func (r OnorbiteventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbiteventTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbiteventTupleParams]'s query parameters as
// `url.Values`.
func (r OnorbiteventTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
