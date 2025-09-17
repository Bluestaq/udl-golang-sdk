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

// SensorMaintenanceHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSensorMaintenanceHistoryService] method instead.
type SensorMaintenanceHistoryService struct {
	Options []option.RequestOption
}

// NewSensorMaintenanceHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSensorMaintenanceHistoryService(opts ...option.RequestOption) (r SensorMaintenanceHistoryService) {
	r = SensorMaintenanceHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensorMaintenanceHistoryService) List(ctx context.Context, query SensorMaintenanceHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SensorMaintenanceHistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sensormaintenance/history"
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
func (r *SensorMaintenanceHistoryService) ListAutoPaging(ctx context.Context, query SensorMaintenanceHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SensorMaintenanceHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensorMaintenanceHistoryService) Aodr(ctx context.Context, query SensorMaintenanceHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sensormaintenance/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SensorMaintenanceHistoryService) Count(ctx context.Context, query SensorMaintenanceHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sensormaintenance/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Maintenance schedule and operational status of Sensor.
type SensorMaintenanceHistoryListResponse struct {
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
	DataMode SensorMaintenanceHistoryListResponseDataMode `json:"dataMode,required"`
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
func (r SensorMaintenanceHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorMaintenanceHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type SensorMaintenanceHistoryListResponseDataMode string

const (
	SensorMaintenanceHistoryListResponseDataModeReal      SensorMaintenanceHistoryListResponseDataMode = "REAL"
	SensorMaintenanceHistoryListResponseDataModeTest      SensorMaintenanceHistoryListResponseDataMode = "TEST"
	SensorMaintenanceHistoryListResponseDataModeSimulated SensorMaintenanceHistoryListResponseDataMode = "SIMULATED"
	SensorMaintenanceHistoryListResponseDataModeExercise  SensorMaintenanceHistoryListResponseDataMode = "EXERCISE"
)

type SensorMaintenanceHistoryListParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns param.Opt[string] `query:"columns,omitzero" json:"-"`
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

// URLQuery serializes [SensorMaintenanceHistoryListParams]'s query parameters as
// `url.Values`.
func (r SensorMaintenanceHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorMaintenanceHistoryAodrParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns param.Opt[string] `query:"columns,omitzero" json:"-"`
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// end time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EndTime     param.Opt[time.Time] `query:"endTime,omitzero" format:"date-time" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
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
	// (One or more of fields 'endTime, startTime' are required.) The planned outage
	// start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SensorMaintenanceHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r SensorMaintenanceHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorMaintenanceHistoryCountParams struct {
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

// URLQuery serializes [SensorMaintenanceHistoryCountParams]'s query parameters as
// `url.Values`.
func (r SensorMaintenanceHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
