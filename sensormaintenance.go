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

// SensormaintenanceService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSensormaintenanceService] method instead.
type SensormaintenanceService struct {
	Options []option.RequestOption
	History SensormaintenanceHistoryService
}

// NewSensormaintenanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSensormaintenanceService(opts ...option.RequestOption) (r SensormaintenanceService) {
	r = SensormaintenanceService{}
	r.Options = opts
	r.History = NewSensormaintenanceHistoryService(opts...)
	return
}

// Service operation to take a single SensorMaintenance as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SensormaintenanceService) New(ctx context.Context, body SensormaintenanceNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensormaintenance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single SensorMaintenance. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SensormaintenanceService) Update(ctx context.Context, id string, body SensormaintenanceUpdateParams, opts ...option.RequestOption) (err error) {
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
func (r *SensormaintenanceService) List(ctx context.Context, query SensormaintenanceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SensormaintenanceListResponse], err error) {
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
func (r *SensormaintenanceService) ListAutoPaging(ctx context.Context, query SensormaintenanceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SensormaintenanceListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a SensorMaintenance object specified by the passed
// ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SensormaintenanceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
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
func (r *SensormaintenanceService) Count(ctx context.Context, query SensormaintenanceCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sensormaintenance/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple SensorMaintenance as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SensormaintenanceService) NewBulk(ctx context.Context, params SensormaintenanceNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensormaintenance/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Service operation to get current Sensor Maintenance records using any number of
// additional parameters.
func (r *SensormaintenanceService) Current(ctx context.Context, query SensormaintenanceCurrentParams, opts ...option.RequestOption) (res *[]SensormaintenanceFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensormaintenance/current"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SensorMaintenance record by its unique ID
// passed as a path parameter.
func (r *SensormaintenanceService) Get(ctx context.Context, id string, query SensormaintenanceGetParams, opts ...option.RequestOption) (res *SensormaintenanceFull, err error) {
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
func (r *SensormaintenanceService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensormaintenance/queryhelp"
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
func (r *SensormaintenanceService) Tuple(ctx context.Context, query SensormaintenanceTupleParams, opts ...option.RequestOption) (res *[]SensormaintenanceFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sensormaintenance/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Maintenance schedule and operational status of Sensor.
type SensormaintenanceListResponse struct {
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
	DataMode SensormaintenanceListResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		EndTime               resp.Field
		SiteCode              resp.Field
		Source                resp.Field
		StartTime             resp.Field
		ID                    resp.Field
		Activity              resp.Field
		Approver              resp.Field
		Changer               resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Duration              resp.Field
		EowID                 resp.Field
		EquipStatus           resp.Field
		IDSensor              resp.Field
		ImpactedFaces         resp.Field
		InactiveDate          resp.Field
		LineNumber            resp.Field
		MdOpsCap              resp.Field
		MwOpsCap              resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Priority              resp.Field
		Recall                resp.Field
		Rel                   resp.Field
		Remark                resp.Field
		Requestor             resp.Field
		Resource              resp.Field
		Rev                   resp.Field
		SSOpsCap              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensormaintenanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *SensormaintenanceListResponse) UnmarshalJSON(data []byte) error {
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
type SensormaintenanceListResponseDataMode string

const (
	SensormaintenanceListResponseDataModeReal      SensormaintenanceListResponseDataMode = "REAL"
	SensormaintenanceListResponseDataModeTest      SensormaintenanceListResponseDataMode = "TEST"
	SensormaintenanceListResponseDataModeSimulated SensormaintenanceListResponseDataMode = "SIMULATED"
	SensormaintenanceListResponseDataModeExercise  SensormaintenanceListResponseDataMode = "EXERCISE"
)

type SensormaintenanceNewParams struct {
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
	DataMode SensormaintenanceNewParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SensormaintenanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SensormaintenanceNewParams
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
type SensormaintenanceNewParamsDataMode string

const (
	SensormaintenanceNewParamsDataModeReal      SensormaintenanceNewParamsDataMode = "REAL"
	SensormaintenanceNewParamsDataModeTest      SensormaintenanceNewParamsDataMode = "TEST"
	SensormaintenanceNewParamsDataModeSimulated SensormaintenanceNewParamsDataMode = "SIMULATED"
	SensormaintenanceNewParamsDataModeExercise  SensormaintenanceNewParamsDataMode = "EXERCISE"
)

type SensormaintenanceUpdateParams struct {
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
	DataMode SensormaintenanceUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SensormaintenanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SensormaintenanceUpdateParams
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
type SensormaintenanceUpdateParamsDataMode string

const (
	SensormaintenanceUpdateParamsDataModeReal      SensormaintenanceUpdateParamsDataMode = "REAL"
	SensormaintenanceUpdateParamsDataModeTest      SensormaintenanceUpdateParamsDataMode = "TEST"
	SensormaintenanceUpdateParamsDataModeSimulated SensormaintenanceUpdateParamsDataMode = "SIMULATED"
	SensormaintenanceUpdateParamsDataModeExercise  SensormaintenanceUpdateParamsDataMode = "EXERCISE"
)

type SensormaintenanceListParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SensormaintenanceListParams]'s query parameters as
// `url.Values`.
func (r SensormaintenanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensormaintenanceCountParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SensormaintenanceCountParams]'s query parameters as
// `url.Values`.
func (r SensormaintenanceCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensormaintenanceNewBulkParams struct {
	Body []SensormaintenanceNewBulkParamsBody
	// Origin of the SensorMaintenance data.
	Origin param.Opt[string] `query:"origin,omitzero" json:"-"`
	// Source of the SensorMaintenance data.
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SensormaintenanceNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// URLQuery serializes [SensormaintenanceNewBulkParams]'s query parameters as
// `url.Values`.
func (r SensormaintenanceNewBulkParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Maintenance schedule and operational status of Sensor.
//
// The properties ClassificationMarking, DataMode, EndTime, SiteCode, Source,
// StartTime are required.
type SensormaintenanceNewBulkParamsBody struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceNewBulkParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r SensormaintenanceNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SensormaintenanceNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[SensormaintenanceNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SensormaintenanceCurrentParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceCurrentParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SensormaintenanceCurrentParams]'s query parameters as
// `url.Values`.
func (r SensormaintenanceCurrentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensormaintenanceGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SensormaintenanceGetParams]'s query parameters as
// `url.Values`.
func (r SensormaintenanceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensormaintenanceTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SensormaintenanceTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SensormaintenanceTupleParams]'s query parameters as
// `url.Values`.
func (r SensormaintenanceTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
