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
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// SensorMaintenanceService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSensorMaintenanceService] method instead.
type SensorMaintenanceService struct {
	Options []option.RequestOption
	History SensorMaintenanceHistoryService
}

// NewSensorMaintenanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSensorMaintenanceService(opts ...option.RequestOption) (r SensorMaintenanceService) {
	r = SensorMaintenanceService{}
	r.Options = opts
	r.History = NewSensorMaintenanceHistoryService(opts...)
	return
}

// Service operation to take a single SensorMaintenance as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SensorMaintenanceService) New(ctx context.Context, body SensorMaintenanceNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensormaintenance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single SensorMaintenance. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SensorMaintenanceService) Update(ctx context.Context, id string, body SensorMaintenanceUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensormaintenance/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensorMaintenanceService) List(ctx context.Context, query SensorMaintenanceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SensorMaintenanceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sensormaintenance"
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
func (r *SensorMaintenanceService) ListAutoPaging(ctx context.Context, query SensorMaintenanceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SensorMaintenanceListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a SensorMaintenance object specified by the passed
// ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SensorMaintenanceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensormaintenance/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SensorMaintenanceService) Count(ctx context.Context, query SensorMaintenanceCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sensormaintenance/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple SensorMaintenance as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SensorMaintenanceService) NewBulk(ctx context.Context, params SensorMaintenanceNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensormaintenance/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Service operation to get current Sensor Maintenance records using any number of
// additional parameters.
func (r *SensorMaintenanceService) Current(ctx context.Context, query SensorMaintenanceCurrentParams, opts ...option.RequestOption) (res *[]SensorMaintenanceCurrentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensormaintenance/current"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SensorMaintenance record by its unique ID
// passed as a path parameter.
func (r *SensorMaintenanceService) Get(ctx context.Context, id string, query SensorMaintenanceGetParams, opts ...option.RequestOption) (res *SensorMaintenanceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensormaintenance/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SensorMaintenanceService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SensorMaintenanceQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensormaintenance/queryhelp"
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
func (r *SensorMaintenanceService) Tuple(ctx context.Context, query SensorMaintenanceTupleParams, opts ...option.RequestOption) (res *[]SensorMaintenanceTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensormaintenance/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Maintenance schedule and operational status of Sensor.
type SensorMaintenanceListResponse struct {
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
	DataMode SensorMaintenanceListResponseDataMode `json:"dataMode,required"`
	// The planned outage end time in ISO8601 UTC format.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// The site to which this item applies. NOTE - this site code is COLT specific and
	// may not identically correspond to other UDL site IDs.
	SiteCode string `json:"siteCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The planned outage start time in ISO8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Description of the activity taking place during this outage.
	Activity string `json:"activity"`
	// The name of the approver.
	Approver string `json:"approver"`
	// The name of the changer, if applicable.
	Changer string `json:"changer"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The duration of the planned outage, expressed as ddd:hh:mm.
	Duration string `json:"duration"`
	// COLT EOWID.
	EowID string `json:"eowId"`
	// The mission capability status of the equipment (e.g. FMC, NMC, PMC, UNK, etc.).
	EquipStatus string `json:"equipStatus"`
	// UUID of the sensor.
	IDSensor string `json:"idSensor"`
	// The sensor face(s) to which this COLT maintenance item applies, if applicable.
	ImpactedFaces string `json:"impactedFaces"`
	// The date that this item became inactive in ISO8601 UTC format.
	InactiveDate time.Time `json:"inactiveDate" format:"date-time"`
	// The internal COLT line number assigned to this item.
	LineNumber string `json:"lineNumber"`
	// The Missile Defense operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MdOpsCap string `json:"mdOpsCap"`
	// The Missile Warning operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MwOpsCap string `json:"mwOpsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The priority of this maintenance item.
	Priority string `json:"priority"`
	// The minimum time required to recall this activity, expressed as ddd:hh:mm.
	Recall string `json:"recall"`
	// Release.
	Rel string `json:"rel"`
	// Remarks concerning this outage.
	Remark string `json:"remark"`
	// The name of the requestor.
	Requestor string `json:"requestor"`
	// The name of the resource(s) affected by this maintenance item.
	Resource string `json:"resource"`
	// The revision number for this maintenance item.
	Rev string `json:"rev"`
	// The Space Surveillance operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	SSOpsCap string `json:"ssOpsCap"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EndTime               respjson.Field
		SiteCode              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		Activity              respjson.Field
		Approver              respjson.Field
		Changer               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Duration              respjson.Field
		EowID                 respjson.Field
		EquipStatus           respjson.Field
		IDSensor              respjson.Field
		ImpactedFaces         respjson.Field
		InactiveDate          respjson.Field
		LineNumber            respjson.Field
		MdOpsCap              respjson.Field
		MwOpsCap              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Priority              respjson.Field
		Recall                respjson.Field
		Rel                   respjson.Field
		Remark                respjson.Field
		Requestor             respjson.Field
		Resource              respjson.Field
		Rev                   respjson.Field
		SSOpsCap              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorMaintenanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorMaintenanceListResponse) UnmarshalJSON(data []byte) error {
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
type SensorMaintenanceListResponseDataMode string

const (
	SensorMaintenanceListResponseDataModeReal      SensorMaintenanceListResponseDataMode = "REAL"
	SensorMaintenanceListResponseDataModeTest      SensorMaintenanceListResponseDataMode = "TEST"
	SensorMaintenanceListResponseDataModeSimulated SensorMaintenanceListResponseDataMode = "SIMULATED"
	SensorMaintenanceListResponseDataModeExercise  SensorMaintenanceListResponseDataMode = "EXERCISE"
)

// Maintenance schedule and operational status of Sensor.
type SensorMaintenanceCurrentResponse struct {
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
	DataMode SensorMaintenanceCurrentResponseDataMode `json:"dataMode,required"`
	// The planned outage end time in ISO8601 UTC format.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// The site to which this item applies. NOTE - this site code is COLT specific and
	// may not identically correspond to other UDL site IDs.
	SiteCode string `json:"siteCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The planned outage start time in ISO8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Description of the activity taking place during this outage.
	Activity string `json:"activity"`
	// The name of the approver.
	Approver string `json:"approver"`
	// The name of the changer, if applicable.
	Changer string `json:"changer"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The duration of the planned outage, expressed as ddd:hh:mm.
	Duration string `json:"duration"`
	// COLT EOWID.
	EowID string `json:"eowId"`
	// The mission capability status of the equipment (e.g. FMC, NMC, PMC, UNK, etc.).
	EquipStatus string `json:"equipStatus"`
	// UUID of the sensor.
	IDSensor string `json:"idSensor"`
	// The sensor face(s) to which this COLT maintenance item applies, if applicable.
	ImpactedFaces string `json:"impactedFaces"`
	// The date that this item became inactive in ISO8601 UTC format.
	InactiveDate time.Time `json:"inactiveDate" format:"date-time"`
	// The internal COLT line number assigned to this item.
	LineNumber string `json:"lineNumber"`
	// The Missile Defense operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MdOpsCap string `json:"mdOpsCap"`
	// The Missile Warning operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MwOpsCap string `json:"mwOpsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The priority of this maintenance item.
	Priority string `json:"priority"`
	// The minimum time required to recall this activity, expressed as ddd:hh:mm.
	Recall string `json:"recall"`
	// Release.
	Rel string `json:"rel"`
	// Remarks concerning this outage.
	Remark string `json:"remark"`
	// The name of the requestor.
	Requestor string `json:"requestor"`
	// The name of the resource(s) affected by this maintenance item.
	Resource string `json:"resource"`
	// The revision number for this maintenance item.
	Rev string `json:"rev"`
	// The Space Surveillance operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	SSOpsCap string `json:"ssOpsCap"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EndTime               respjson.Field
		SiteCode              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		Activity              respjson.Field
		Approver              respjson.Field
		Changer               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Duration              respjson.Field
		EowID                 respjson.Field
		EquipStatus           respjson.Field
		IDSensor              respjson.Field
		ImpactedFaces         respjson.Field
		InactiveDate          respjson.Field
		LineNumber            respjson.Field
		MdOpsCap              respjson.Field
		MwOpsCap              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Priority              respjson.Field
		Recall                respjson.Field
		Rel                   respjson.Field
		Remark                respjson.Field
		Requestor             respjson.Field
		Resource              respjson.Field
		Rev                   respjson.Field
		SSOpsCap              respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorMaintenanceCurrentResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorMaintenanceCurrentResponse) UnmarshalJSON(data []byte) error {
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
type SensorMaintenanceCurrentResponseDataMode string

const (
	SensorMaintenanceCurrentResponseDataModeReal      SensorMaintenanceCurrentResponseDataMode = "REAL"
	SensorMaintenanceCurrentResponseDataModeTest      SensorMaintenanceCurrentResponseDataMode = "TEST"
	SensorMaintenanceCurrentResponseDataModeSimulated SensorMaintenanceCurrentResponseDataMode = "SIMULATED"
	SensorMaintenanceCurrentResponseDataModeExercise  SensorMaintenanceCurrentResponseDataMode = "EXERCISE"
)

// Maintenance schedule and operational status of Sensor.
type SensorMaintenanceGetResponse struct {
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
	DataMode SensorMaintenanceGetResponseDataMode `json:"dataMode,required"`
	// The planned outage end time in ISO8601 UTC format.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// The site to which this item applies. NOTE - this site code is COLT specific and
	// may not identically correspond to other UDL site IDs.
	SiteCode string `json:"siteCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The planned outage start time in ISO8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Description of the activity taking place during this outage.
	Activity string `json:"activity"`
	// The name of the approver.
	Approver string `json:"approver"`
	// The name of the changer, if applicable.
	Changer string `json:"changer"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The duration of the planned outage, expressed as ddd:hh:mm.
	Duration string `json:"duration"`
	// COLT EOWID.
	EowID string `json:"eowId"`
	// The mission capability status of the equipment (e.g. FMC, NMC, PMC, UNK, etc.).
	EquipStatus string `json:"equipStatus"`
	// UUID of the sensor.
	IDSensor string `json:"idSensor"`
	// The sensor face(s) to which this COLT maintenance item applies, if applicable.
	ImpactedFaces string `json:"impactedFaces"`
	// The date that this item became inactive in ISO8601 UTC format.
	InactiveDate time.Time `json:"inactiveDate" format:"date-time"`
	// The internal COLT line number assigned to this item.
	LineNumber string `json:"lineNumber"`
	// The Missile Defense operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MdOpsCap string `json:"mdOpsCap"`
	// The Missile Warning operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MwOpsCap string `json:"mwOpsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The priority of this maintenance item.
	Priority string `json:"priority"`
	// The minimum time required to recall this activity, expressed as ddd:hh:mm.
	Recall string `json:"recall"`
	// Release.
	Rel string `json:"rel"`
	// Remarks concerning this outage.
	Remark string `json:"remark"`
	// The name of the requestor.
	Requestor string `json:"requestor"`
	// The name of the resource(s) affected by this maintenance item.
	Resource string `json:"resource"`
	// The revision number for this maintenance item.
	Rev string `json:"rev"`
	// The Space Surveillance operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	SSOpsCap string `json:"ssOpsCap"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EndTime               respjson.Field
		SiteCode              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		Activity              respjson.Field
		Approver              respjson.Field
		Changer               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Duration              respjson.Field
		EowID                 respjson.Field
		EquipStatus           respjson.Field
		IDSensor              respjson.Field
		ImpactedFaces         respjson.Field
		InactiveDate          respjson.Field
		LineNumber            respjson.Field
		MdOpsCap              respjson.Field
		MwOpsCap              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Priority              respjson.Field
		Recall                respjson.Field
		Rel                   respjson.Field
		Remark                respjson.Field
		Requestor             respjson.Field
		Resource              respjson.Field
		Rev                   respjson.Field
		SSOpsCap              respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorMaintenanceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorMaintenanceGetResponse) UnmarshalJSON(data []byte) error {
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
type SensorMaintenanceGetResponseDataMode string

const (
	SensorMaintenanceGetResponseDataModeReal      SensorMaintenanceGetResponseDataMode = "REAL"
	SensorMaintenanceGetResponseDataModeTest      SensorMaintenanceGetResponseDataMode = "TEST"
	SensorMaintenanceGetResponseDataModeSimulated SensorMaintenanceGetResponseDataMode = "SIMULATED"
	SensorMaintenanceGetResponseDataModeExercise  SensorMaintenanceGetResponseDataMode = "EXERCISE"
)

type SensorMaintenanceQueryhelpResponse struct {
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
func (r SensorMaintenanceQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorMaintenanceQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maintenance schedule and operational status of Sensor.
type SensorMaintenanceTupleResponse struct {
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
	DataMode SensorMaintenanceTupleResponseDataMode `json:"dataMode,required"`
	// The planned outage end time in ISO8601 UTC format.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// The site to which this item applies. NOTE - this site code is COLT specific and
	// may not identically correspond to other UDL site IDs.
	SiteCode string `json:"siteCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The planned outage start time in ISO8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Description of the activity taking place during this outage.
	Activity string `json:"activity"`
	// The name of the approver.
	Approver string `json:"approver"`
	// The name of the changer, if applicable.
	Changer string `json:"changer"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The duration of the planned outage, expressed as ddd:hh:mm.
	Duration string `json:"duration"`
	// COLT EOWID.
	EowID string `json:"eowId"`
	// The mission capability status of the equipment (e.g. FMC, NMC, PMC, UNK, etc.).
	EquipStatus string `json:"equipStatus"`
	// UUID of the sensor.
	IDSensor string `json:"idSensor"`
	// The sensor face(s) to which this COLT maintenance item applies, if applicable.
	ImpactedFaces string `json:"impactedFaces"`
	// The date that this item became inactive in ISO8601 UTC format.
	InactiveDate time.Time `json:"inactiveDate" format:"date-time"`
	// The internal COLT line number assigned to this item.
	LineNumber string `json:"lineNumber"`
	// The Missile Defense operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MdOpsCap string `json:"mdOpsCap"`
	// The Missile Warning operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MwOpsCap string `json:"mwOpsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The priority of this maintenance item.
	Priority string `json:"priority"`
	// The minimum time required to recall this activity, expressed as ddd:hh:mm.
	Recall string `json:"recall"`
	// Release.
	Rel string `json:"rel"`
	// Remarks concerning this outage.
	Remark string `json:"remark"`
	// The name of the requestor.
	Requestor string `json:"requestor"`
	// The name of the resource(s) affected by this maintenance item.
	Resource string `json:"resource"`
	// The revision number for this maintenance item.
	Rev string `json:"rev"`
	// The Space Surveillance operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	SSOpsCap string `json:"ssOpsCap"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EndTime               respjson.Field
		SiteCode              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		Activity              respjson.Field
		Approver              respjson.Field
		Changer               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Duration              respjson.Field
		EowID                 respjson.Field
		EquipStatus           respjson.Field
		IDSensor              respjson.Field
		ImpactedFaces         respjson.Field
		InactiveDate          respjson.Field
		LineNumber            respjson.Field
		MdOpsCap              respjson.Field
		MwOpsCap              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Priority              respjson.Field
		Recall                respjson.Field
		Rel                   respjson.Field
		Remark                respjson.Field
		Requestor             respjson.Field
		Resource              respjson.Field
		Rev                   respjson.Field
		SSOpsCap              respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorMaintenanceTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorMaintenanceTupleResponse) UnmarshalJSON(data []byte) error {
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
type SensorMaintenanceTupleResponseDataMode string

const (
	SensorMaintenanceTupleResponseDataModeReal      SensorMaintenanceTupleResponseDataMode = "REAL"
	SensorMaintenanceTupleResponseDataModeTest      SensorMaintenanceTupleResponseDataMode = "TEST"
	SensorMaintenanceTupleResponseDataModeSimulated SensorMaintenanceTupleResponseDataMode = "SIMULATED"
	SensorMaintenanceTupleResponseDataModeExercise  SensorMaintenanceTupleResponseDataMode = "EXERCISE"
)

type SensorMaintenanceNewParams struct {
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
	DataMode SensorMaintenanceNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The planned outage end time in ISO8601 UTC format.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// The site to which this item applies. NOTE - this site code is COLT specific and
	// may not identically correspond to other UDL site IDs.
	SiteCode string `json:"siteCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The planned outage start time in ISO8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Description of the activity taking place during this outage.
	Activity param.Opt[string] `json:"activity,omitzero"`
	// The name of the approver.
	Approver param.Opt[string] `json:"approver,omitzero"`
	// The name of the changer, if applicable.
	Changer param.Opt[string] `json:"changer,omitzero"`
	// The duration of the planned outage, expressed as ddd:hh:mm.
	Duration param.Opt[string] `json:"duration,omitzero"`
	// COLT EOWID.
	EowID param.Opt[string] `json:"eowId,omitzero"`
	// The mission capability status of the equipment (e.g. FMC, NMC, PMC, UNK, etc.).
	EquipStatus param.Opt[string] `json:"equipStatus,omitzero"`
	// UUID of the sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The sensor face(s) to which this COLT maintenance item applies, if applicable.
	ImpactedFaces param.Opt[string] `json:"impactedFaces,omitzero"`
	// The internal COLT line number assigned to this item.
	LineNumber param.Opt[string] `json:"lineNumber,omitzero"`
	// The Missile Defense operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MdOpsCap param.Opt[string] `json:"mdOpsCap,omitzero"`
	// The Missile Warning operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MwOpsCap param.Opt[string] `json:"mwOpsCap,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The priority of this maintenance item.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The minimum time required to recall this activity, expressed as ddd:hh:mm.
	Recall param.Opt[string] `json:"recall,omitzero"`
	// Release.
	Rel param.Opt[string] `json:"rel,omitzero"`
	// Remarks concerning this outage.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The name of the requestor.
	Requestor param.Opt[string] `json:"requestor,omitzero"`
	// The name of the resource(s) affected by this maintenance item.
	Resource param.Opt[string] `json:"resource,omitzero"`
	// The revision number for this maintenance item.
	Rev param.Opt[string] `json:"rev,omitzero"`
	// The Space Surveillance operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	SSOpsCap param.Opt[string] `json:"ssOpsCap,omitzero"`
	paramObj
}

func (r SensorMaintenanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SensorMaintenanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorMaintenanceNewParams) UnmarshalJSON(data []byte) error {
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
type SensorMaintenanceNewParamsDataMode string

const (
	SensorMaintenanceNewParamsDataModeReal      SensorMaintenanceNewParamsDataMode = "REAL"
	SensorMaintenanceNewParamsDataModeTest      SensorMaintenanceNewParamsDataMode = "TEST"
	SensorMaintenanceNewParamsDataModeSimulated SensorMaintenanceNewParamsDataMode = "SIMULATED"
	SensorMaintenanceNewParamsDataModeExercise  SensorMaintenanceNewParamsDataMode = "EXERCISE"
)

type SensorMaintenanceUpdateParams struct {
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
	DataMode SensorMaintenanceUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The planned outage end time in ISO8601 UTC format.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// The site to which this item applies. NOTE - this site code is COLT specific and
	// may not identically correspond to other UDL site IDs.
	SiteCode string `json:"siteCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The planned outage start time in ISO8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Description of the activity taking place during this outage.
	Activity param.Opt[string] `json:"activity,omitzero"`
	// The name of the approver.
	Approver param.Opt[string] `json:"approver,omitzero"`
	// The name of the changer, if applicable.
	Changer param.Opt[string] `json:"changer,omitzero"`
	// The duration of the planned outage, expressed as ddd:hh:mm.
	Duration param.Opt[string] `json:"duration,omitzero"`
	// COLT EOWID.
	EowID param.Opt[string] `json:"eowId,omitzero"`
	// The mission capability status of the equipment (e.g. FMC, NMC, PMC, UNK, etc.).
	EquipStatus param.Opt[string] `json:"equipStatus,omitzero"`
	// UUID of the sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The sensor face(s) to which this COLT maintenance item applies, if applicable.
	ImpactedFaces param.Opt[string] `json:"impactedFaces,omitzero"`
	// The internal COLT line number assigned to this item.
	LineNumber param.Opt[string] `json:"lineNumber,omitzero"`
	// The Missile Defense operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MdOpsCap param.Opt[string] `json:"mdOpsCap,omitzero"`
	// The Missile Warning operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MwOpsCap param.Opt[string] `json:"mwOpsCap,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The priority of this maintenance item.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The minimum time required to recall this activity, expressed as ddd:hh:mm.
	Recall param.Opt[string] `json:"recall,omitzero"`
	// Release.
	Rel param.Opt[string] `json:"rel,omitzero"`
	// Remarks concerning this outage.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The name of the requestor.
	Requestor param.Opt[string] `json:"requestor,omitzero"`
	// The name of the resource(s) affected by this maintenance item.
	Resource param.Opt[string] `json:"resource,omitzero"`
	// The revision number for this maintenance item.
	Rev param.Opt[string] `json:"rev,omitzero"`
	// The Space Surveillance operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	SSOpsCap param.Opt[string] `json:"ssOpsCap,omitzero"`
	paramObj
}

func (r SensorMaintenanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SensorMaintenanceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorMaintenanceUpdateParams) UnmarshalJSON(data []byte) error {
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
type SensorMaintenanceUpdateParamsDataMode string

const (
	SensorMaintenanceUpdateParamsDataModeReal      SensorMaintenanceUpdateParamsDataMode = "REAL"
	SensorMaintenanceUpdateParamsDataModeTest      SensorMaintenanceUpdateParamsDataMode = "TEST"
	SensorMaintenanceUpdateParamsDataModeSimulated SensorMaintenanceUpdateParamsDataMode = "SIMULATED"
	SensorMaintenanceUpdateParamsDataModeExercise  SensorMaintenanceUpdateParamsDataMode = "EXERCISE"
)

type SensorMaintenanceListParams struct {
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// end time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EndTime     param.Opt[time.Time] `query:"endTime,omitzero" format:"date-time" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SensorMaintenanceListParams]'s query parameters as
// `url.Values`.
func (r SensorMaintenanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorMaintenanceCountParams struct {
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// end time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EndTime     param.Opt[time.Time] `query:"endTime,omitzero" format:"date-time" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SensorMaintenanceCountParams]'s query parameters as
// `url.Values`.
func (r SensorMaintenanceCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorMaintenanceNewBulkParams struct {
	Body []SensorMaintenanceNewBulkParamsBody
	// Origin of the SensorMaintenance data.
	Origin param.Opt[string] `query:"origin,omitzero" json:"-"`
	// Source of the SensorMaintenance data.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	paramObj
}

func (r SensorMaintenanceNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SensorMaintenanceNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [SensorMaintenanceNewBulkParams]'s query parameters as
// `url.Values`.
func (r SensorMaintenanceNewBulkParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Maintenance schedule and operational status of Sensor.
//
// The properties ClassificationMarking, DataMode, EndTime, SiteCode, Source,
// StartTime are required.
type SensorMaintenanceNewBulkParamsBody struct {
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
	// The planned outage end time in ISO8601 UTC format.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// The site to which this item applies. NOTE - this site code is COLT specific and
	// may not identically correspond to other UDL site IDs.
	SiteCode string `json:"siteCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The planned outage start time in ISO8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Description of the activity taking place during this outage.
	Activity param.Opt[string] `json:"activity,omitzero"`
	// The name of the approver.
	Approver param.Opt[string] `json:"approver,omitzero"`
	// The name of the changer, if applicable.
	Changer param.Opt[string] `json:"changer,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The duration of the planned outage, expressed as ddd:hh:mm.
	Duration param.Opt[string] `json:"duration,omitzero"`
	// COLT EOWID.
	EowID param.Opt[string] `json:"eowId,omitzero"`
	// The mission capability status of the equipment (e.g. FMC, NMC, PMC, UNK, etc.).
	EquipStatus param.Opt[string] `json:"equipStatus,omitzero"`
	// UUID of the sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The sensor face(s) to which this COLT maintenance item applies, if applicable.
	ImpactedFaces param.Opt[string] `json:"impactedFaces,omitzero"`
	// The date that this item became inactive in ISO8601 UTC format.
	InactiveDate param.Opt[time.Time] `json:"inactiveDate,omitzero" format:"date-time"`
	// The internal COLT line number assigned to this item.
	LineNumber param.Opt[string] `json:"lineNumber,omitzero"`
	// The Missile Defense operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MdOpsCap param.Opt[string] `json:"mdOpsCap,omitzero"`
	// The Missile Warning operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	MwOpsCap param.Opt[string] `json:"mwOpsCap,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The priority of this maintenance item.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The minimum time required to recall this activity, expressed as ddd:hh:mm.
	Recall param.Opt[string] `json:"recall,omitzero"`
	// Release.
	Rel param.Opt[string] `json:"rel,omitzero"`
	// Remarks concerning this outage.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The name of the requestor.
	Requestor param.Opt[string] `json:"requestor,omitzero"`
	// The name of the resource(s) affected by this maintenance item.
	Resource param.Opt[string] `json:"resource,omitzero"`
	// The revision number for this maintenance item.
	Rev param.Opt[string] `json:"rev,omitzero"`
	// The Space Surveillance operational capability of this maintenance item. Typical
	// values are G, Y, R, and - for non-applicable sites.
	SSOpsCap param.Opt[string] `json:"ssOpsCap,omitzero"`
	paramObj
}

func (r SensorMaintenanceNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SensorMaintenanceNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorMaintenanceNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SensorMaintenanceNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SensorMaintenanceCurrentParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorMaintenanceCurrentParams]'s query parameters as
// `url.Values`.
func (r SensorMaintenanceCurrentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorMaintenanceGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorMaintenanceGetParams]'s query parameters as
// `url.Values`.
func (r SensorMaintenanceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorMaintenanceTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// end time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EndTime     param.Opt[time.Time] `query:"endTime,omitzero" format:"date-time" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SensorMaintenanceTupleParams]'s query parameters as
// `url.Values`.
func (r SensorMaintenanceTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
