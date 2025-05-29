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
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// CollectResponseService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectResponseService] method instead.
type CollectResponseService struct {
	Options []option.RequestOption
	History CollectResponseHistoryService
	Tuple   CollectResponseTupleService
}

// NewCollectResponseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectResponseService(opts ...option.RequestOption) (r CollectResponseService) {
	r = CollectResponseService{}
	r.Options = opts
	r.History = NewCollectResponseHistoryService(opts...)
	r.Tuple = NewCollectResponseTupleService(opts...)
	return
}

// Service operation to take a single Collect Response object as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *CollectResponseService) New(ctx context.Context, body CollectResponseNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/collectresponse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Collect Response record by its unique ID
// passed as a path parameter.
func (r *CollectResponseService) Get(ctx context.Context, id string, query CollectResponseGetParams, opts ...option.RequestOption) (res *shared.CollectResponseFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/collectresponse/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *CollectResponseService) List(ctx context.Context, query CollectResponseListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[CollectResponseAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/collectresponse"
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
func (r *CollectResponseService) ListAutoPaging(ctx context.Context, query CollectResponseListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[CollectResponseAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *CollectResponseService) Count(ctx context.Context, query CollectResponseCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/collectresponse/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// Collect Response objects as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *CollectResponseService) NewBulk(ctx context.Context, body CollectResponseNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/collectresponse/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *CollectResponseService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *CollectResponseQueryHelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/collectresponse/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Service operation to take a list of CollectResponse as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *CollectResponseService) UnvalidatedPublish(ctx context.Context, body CollectResponseUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-collectresponse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Collect response supports the response and status of individual collect
// requests. Each response is referenced by the UUID of the request, and contains
// information including the status of the request, collection times and types, and
// reference(s) to the observations collected. There may be multiple responses
// associated with a request, either from multiple collectors or to relay status
// changes prior to completion and delivery.
type CollectResponseAbridged struct {
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
	DataMode CollectResponseAbridgedDataMode `json:"dataMode,required"`
	// Unique identifier of the request associated with this response.
	IDRequest string `json:"idRequest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The actual end time of the collect or contact, in ISO 8601 UTC format.
	ActualEndTime time.Time `json:"actualEndTime" format:"date-time"`
	// The actual start time of the collect or contact, in ISO 8601 UTC format.
	ActualStartTime time.Time `json:"actualStartTime" format:"date-time"`
	// Proposed alternative end time, in ISO 8601 UTC format.
	AltEndTime time.Time `json:"altEndTime" format:"date-time"`
	// Proposed alternative start time, in ISO 8601 UTC format.
	AltStartTime time.Time `json:"altStartTime" format:"date-time"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Error code associated with this request/response.
	ErrCode string `json:"errCode"`
	// UUID from external systems. This field has no meaning within UDL and is provided
	// as a convenience for systems that require tracking of internal system generated
	// ID.
	ExternalID string `json:"externalId"`
	// Unique identifier of the target on-orbit object associated with this response.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the parent plan or schedule associated with the
	// request/response.
	IDPlan string `json:"idPlan"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// Notes or comments associated with this response.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by data source to indicate the target object of
	// this response. This may be an internal identifier and not necessarily a valid
	// satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the collection source to indicate the sensor
	// identifier responding to this collect or contact. This may be an internal
	// identifier and not necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Array of UUIDs of the UDL data record(s) collected in response to the associated
	// request. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. The appropriate API operation can be used to
	// retrieve the specified object(s) (e.g. /udl/rfobservation/{uuid}).
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record type(s) (DOA, ELSET, EO, RADAR, RF, SV) collected or
	// produced in response to the associated request. See the associated 'srcIds'
	// array for the record UUIDs, positionally corresponding to the record types in
	// this array. The 'srcTyps' and 'srcIds' arrays must match in size. The
	// appropriate API operation can be used to retrieve the specified object(s) (e.g.
	// /udl/rfobservation/{uuid}).
	SrcTyps []string `json:"srcTyps"`
	// The status of the request (ACCEPTED, CANCELLED, COLLECTED, COMPLETED, DELIVERED,
	// FAILED, PARTIAL, PROPOSED, REJECTED, SCHEDULED):
	//
	// ACCEPTED: The collect or contact request has been received and accepted.
	//
	// CANCELLED: A previously scheduled collect or contact whose execution was
	// cancelled.
	//
	// COLLECTED: The collect has been accomplished. A collected state implies that
	// additional activity is required for delivery/completion.
	//
	// COMPLETED: The collect or contact has been completed. For many systems completed
	// and delivered constitute an equivalent successful end state.
	//
	// DELIVERED: The collected observation(s) have been delivered to the requestor.
	// For many systems completed and delivered constitute an equivalent successful end
	// state. A DELIVERED state is typically used for systems that exhibit a delay
	// between collect and delivery, such as with space-based systems which require
	// ground contact to deliver observations.
	//
	// FAILED: The collect or contact was attempted and failed, or the delivery of the
	// collected observation(s) failed. A FAILED status may be accompanied by an error
	// code (errCode), if available.
	//
	// PARTIAL: A PARTIAL state indicates that a part of a multi-track request has been
	// accomplished, but the full request is incomplete. A PARTIAL status should
	// ultimately be resolved to an end state.
	//
	// PROPOSED: Indicates that the request was received and alternate collect or
	// contact time(s) (altStartTime, altEndTime) have been proposed. If an alternate
	// is accepted by the requestor the current request should be cancelled and a new
	// request created.
	//
	// REJECTED: The request has been received and rejected by the provider. A REJECTED
	// status may be accompanied by an explanation (notes) of the reason that the
	// request was rejected.
	//
	// SCHEDULED: The request was received and has been scheduled for execution.
	Status string `json:"status"`
	// Optional task ID associated with the request/response.
	TaskID string `json:"taskId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDRequest             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ActualEndTime         respjson.Field
		ActualStartTime       respjson.Field
		AltEndTime            respjson.Field
		AltStartTime          respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ErrCode               respjson.Field
		ExternalID            respjson.Field
		IDOnOrbit             respjson.Field
		IDPlan                respjson.Field
		IDSensor              respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		SatNo                 respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		Status                respjson.Field
		TaskID                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectResponseAbridged) RawJSON() string { return r.JSON.raw }
func (r *CollectResponseAbridged) UnmarshalJSON(data []byte) error {
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
type CollectResponseAbridgedDataMode string

const (
	CollectResponseAbridgedDataModeReal      CollectResponseAbridgedDataMode = "REAL"
	CollectResponseAbridgedDataModeTest      CollectResponseAbridgedDataMode = "TEST"
	CollectResponseAbridgedDataModeSimulated CollectResponseAbridgedDataMode = "SIMULATED"
	CollectResponseAbridgedDataModeExercise  CollectResponseAbridgedDataMode = "EXERCISE"
)

type CollectResponseQueryHelpResponse struct {
	AodrSupported         bool                                        `json:"aodrSupported"`
	ClassificationMarking string                                      `json:"classificationMarking"`
	Description           string                                      `json:"description"`
	HistorySupported      bool                                        `json:"historySupported"`
	Name                  string                                      `json:"name"`
	Parameters            []CollectResponseQueryHelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                    `json:"requiredRoles"`
	RestSupported         bool                                        `json:"restSupported"`
	SortSupported         bool                                        `json:"sortSupported"`
	TypeName              string                                      `json:"typeName"`
	Uri                   string                                      `json:"uri"`
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
func (r CollectResponseQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectResponseQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectResponseQueryHelpResponseParameter struct {
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
func (r CollectResponseQueryHelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *CollectResponseQueryHelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectResponseNewParams struct {
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
	DataMode CollectResponseNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the request associated with this response.
	IDRequest string `json:"idRequest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The actual end time of the collect or contact, in ISO 8601 UTC format.
	ActualEndTime param.Opt[time.Time] `json:"actualEndTime,omitzero" format:"date-time"`
	// The actual start time of the collect or contact, in ISO 8601 UTC format.
	ActualStartTime param.Opt[time.Time] `json:"actualStartTime,omitzero" format:"date-time"`
	// Proposed alternative end time, in ISO 8601 UTC format.
	AltEndTime param.Opt[time.Time] `json:"altEndTime,omitzero" format:"date-time"`
	// Proposed alternative start time, in ISO 8601 UTC format.
	AltStartTime param.Opt[time.Time] `json:"altStartTime,omitzero" format:"date-time"`
	// Error code associated with this request/response.
	ErrCode param.Opt[string] `json:"errCode,omitzero"`
	// UUID from external systems. This field has no meaning within UDL and is provided
	// as a convenience for systems that require tracking of internal system generated
	// ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Unique identifier of the parent plan or schedule associated with the
	// request/response.
	IDPlan param.Opt[string] `json:"idPlan,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Notes or comments associated with this response.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by data source to indicate the target object of
	// this response. This may be an internal identifier and not necessarily a valid
	// satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the collection source to indicate the sensor
	// identifier responding to this collect or contact. This may be an internal
	// identifier and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The status of the request (ACCEPTED, CANCELLED, COLLECTED, COMPLETED, DELIVERED,
	// FAILED, PARTIAL, PROPOSED, REJECTED, SCHEDULED):
	//
	// ACCEPTED: The collect or contact request has been received and accepted.
	//
	// CANCELLED: A previously scheduled collect or contact whose execution was
	// cancelled.
	//
	// COLLECTED: The collect has been accomplished. A collected state implies that
	// additional activity is required for delivery/completion.
	//
	// COMPLETED: The collect or contact has been completed. For many systems completed
	// and delivered constitute an equivalent successful end state.
	//
	// DELIVERED: The collected observation(s) have been delivered to the requestor.
	// For many systems completed and delivered constitute an equivalent successful end
	// state. A DELIVERED state is typically used for systems that exhibit a delay
	// between collect and delivery, such as with space-based systems which require
	// ground contact to deliver observations.
	//
	// FAILED: The collect or contact was attempted and failed, or the delivery of the
	// collected observation(s) failed. A FAILED status may be accompanied by an error
	// code (errCode), if available.
	//
	// PARTIAL: A PARTIAL state indicates that a part of a multi-track request has been
	// accomplished, but the full request is incomplete. A PARTIAL status should
	// ultimately be resolved to an end state.
	//
	// PROPOSED: Indicates that the request was received and alternate collect or
	// contact time(s) (altStartTime, altEndTime) have been proposed. If an alternate
	// is accepted by the requestor the current request should be cancelled and a new
	// request created.
	//
	// REJECTED: The request has been received and rejected by the provider. A REJECTED
	// status may be accompanied by an explanation (notes) of the reason that the
	// request was rejected.
	//
	// SCHEDULED: The request was received and has been scheduled for execution.
	Status param.Opt[string] `json:"status,omitzero"`
	// Optional task ID associated with the request/response.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Array of UUIDs of the UDL data record(s) collected in response to the associated
	// request. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. The appropriate API operation can be used to
	// retrieve the specified object(s) (e.g. /udl/rfobservation/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record type(s) (DOA, ELSET, EO, RADAR, RF, SV) collected or
	// produced in response to the associated request. See the associated 'srcIds'
	// array for the record UUIDs, positionally corresponding to the record types in
	// this array. The 'srcTyps' and 'srcIds' arrays must match in size. The
	// appropriate API operation can be used to retrieve the specified object(s) (e.g.
	// /udl/rfobservation/{uuid}).
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r CollectResponseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectResponseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectResponseNewParams) UnmarshalJSON(data []byte) error {
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
type CollectResponseNewParamsDataMode string

const (
	CollectResponseNewParamsDataModeReal      CollectResponseNewParamsDataMode = "REAL"
	CollectResponseNewParamsDataModeTest      CollectResponseNewParamsDataMode = "TEST"
	CollectResponseNewParamsDataModeSimulated CollectResponseNewParamsDataMode = "SIMULATED"
	CollectResponseNewParamsDataModeExercise  CollectResponseNewParamsDataMode = "EXERCISE"
)

type CollectResponseGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectResponseGetParams]'s query parameters as
// `url.Values`.
func (r CollectResponseGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectResponseListParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectResponseListParams]'s query parameters as
// `url.Values`.
func (r CollectResponseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectResponseCountParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectResponseCountParams]'s query parameters as
// `url.Values`.
func (r CollectResponseCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectResponseNewBulkParams struct {
	Body []CollectResponseNewBulkParamsBody
	paramObj
}

func (r CollectResponseNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *CollectResponseNewBulkParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// Collect response supports the response and status of individual collect
// requests. Each response is referenced by the UUID of the request, and contains
// information including the status of the request, collection times and types, and
// reference(s) to the observations collected. There may be multiple responses
// associated with a request, either from multiple collectors or to relay status
// changes prior to completion and delivery.
//
// The properties ClassificationMarking, DataMode, IDRequest, Source are required.
type CollectResponseNewBulkParamsBody struct {
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
	// Unique identifier of the request associated with this response.
	IDRequest string `json:"idRequest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The actual end time of the collect or contact, in ISO 8601 UTC format.
	ActualEndTime param.Opt[time.Time] `json:"actualEndTime,omitzero" format:"date-time"`
	// The actual start time of the collect or contact, in ISO 8601 UTC format.
	ActualStartTime param.Opt[time.Time] `json:"actualStartTime,omitzero" format:"date-time"`
	// Proposed alternative end time, in ISO 8601 UTC format.
	AltEndTime param.Opt[time.Time] `json:"altEndTime,omitzero" format:"date-time"`
	// Proposed alternative start time, in ISO 8601 UTC format.
	AltStartTime param.Opt[time.Time] `json:"altStartTime,omitzero" format:"date-time"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Error code associated with this request/response.
	ErrCode param.Opt[string] `json:"errCode,omitzero"`
	// UUID from external systems. This field has no meaning within UDL and is provided
	// as a convenience for systems that require tracking of internal system generated
	// ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Unique identifier of the target on-orbit object associated with this response.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the parent plan or schedule associated with the
	// request/response.
	IDPlan param.Opt[string] `json:"idPlan,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Notes or comments associated with this response.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by data source to indicate the target object of
	// this response. This may be an internal identifier and not necessarily a valid
	// satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the collection source to indicate the sensor
	// identifier responding to this collect or contact. This may be an internal
	// identifier and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The status of the request (ACCEPTED, CANCELLED, COLLECTED, COMPLETED, DELIVERED,
	// FAILED, PARTIAL, PROPOSED, REJECTED, SCHEDULED):
	//
	// ACCEPTED: The collect or contact request has been received and accepted.
	//
	// CANCELLED: A previously scheduled collect or contact whose execution was
	// cancelled.
	//
	// COLLECTED: The collect has been accomplished. A collected state implies that
	// additional activity is required for delivery/completion.
	//
	// COMPLETED: The collect or contact has been completed. For many systems completed
	// and delivered constitute an equivalent successful end state.
	//
	// DELIVERED: The collected observation(s) have been delivered to the requestor.
	// For many systems completed and delivered constitute an equivalent successful end
	// state. A DELIVERED state is typically used for systems that exhibit a delay
	// between collect and delivery, such as with space-based systems which require
	// ground contact to deliver observations.
	//
	// FAILED: The collect or contact was attempted and failed, or the delivery of the
	// collected observation(s) failed. A FAILED status may be accompanied by an error
	// code (errCode), if available.
	//
	// PARTIAL: A PARTIAL state indicates that a part of a multi-track request has been
	// accomplished, but the full request is incomplete. A PARTIAL status should
	// ultimately be resolved to an end state.
	//
	// PROPOSED: Indicates that the request was received and alternate collect or
	// contact time(s) (altStartTime, altEndTime) have been proposed. If an alternate
	// is accepted by the requestor the current request should be cancelled and a new
	// request created.
	//
	// REJECTED: The request has been received and rejected by the provider. A REJECTED
	// status may be accompanied by an explanation (notes) of the reason that the
	// request was rejected.
	//
	// SCHEDULED: The request was received and has been scheduled for execution.
	Status param.Opt[string] `json:"status,omitzero"`
	// Optional task ID associated with the request/response.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Array of UUIDs of the UDL data record(s) collected in response to the associated
	// request. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. The appropriate API operation can be used to
	// retrieve the specified object(s) (e.g. /udl/rfobservation/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record type(s) (DOA, ELSET, EO, RADAR, RF, SV) collected or
	// produced in response to the associated request. See the associated 'srcIds'
	// array for the record UUIDs, positionally corresponding to the record types in
	// this array. The 'srcTyps' and 'srcIds' arrays must match in size. The
	// appropriate API operation can be used to retrieve the specified object(s) (e.g.
	// /udl/rfobservation/{uuid}).
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r CollectResponseNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow CollectResponseNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectResponseNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectResponseNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type CollectResponseUnvalidatedPublishParams struct {
	Body []CollectResponseUnvalidatedPublishParamsBody
	paramObj
}

func (r CollectResponseUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *CollectResponseUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// Collect response supports the response and status of individual collect
// requests. Each response is referenced by the UUID of the request, and contains
// information including the status of the request, collection times and types, and
// reference(s) to the observations collected. There may be multiple responses
// associated with a request, either from multiple collectors or to relay status
// changes prior to completion and delivery.
//
// The properties ClassificationMarking, DataMode, IDRequest, Source are required.
type CollectResponseUnvalidatedPublishParamsBody struct {
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
	// Unique identifier of the request associated with this response.
	IDRequest string `json:"idRequest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The actual end time of the collect or contact, in ISO 8601 UTC format.
	ActualEndTime param.Opt[time.Time] `json:"actualEndTime,omitzero" format:"date-time"`
	// The actual start time of the collect or contact, in ISO 8601 UTC format.
	ActualStartTime param.Opt[time.Time] `json:"actualStartTime,omitzero" format:"date-time"`
	// Proposed alternative end time, in ISO 8601 UTC format.
	AltEndTime param.Opt[time.Time] `json:"altEndTime,omitzero" format:"date-time"`
	// Proposed alternative start time, in ISO 8601 UTC format.
	AltStartTime param.Opt[time.Time] `json:"altStartTime,omitzero" format:"date-time"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Error code associated with this request/response.
	ErrCode param.Opt[string] `json:"errCode,omitzero"`
	// UUID from external systems. This field has no meaning within UDL and is provided
	// as a convenience for systems that require tracking of internal system generated
	// ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Unique identifier of the target on-orbit object associated with this response.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the parent plan or schedule associated with the
	// request/response.
	IDPlan param.Opt[string] `json:"idPlan,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Notes or comments associated with this response.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by data source to indicate the target object of
	// this response. This may be an internal identifier and not necessarily a valid
	// satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the collection source to indicate the sensor
	// identifier responding to this collect or contact. This may be an internal
	// identifier and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The status of the request (ACCEPTED, CANCELLED, COLLECTED, COMPLETED, DELIVERED,
	// FAILED, PARTIAL, PROPOSED, REJECTED, SCHEDULED):
	//
	// ACCEPTED: The collect or contact request has been received and accepted.
	//
	// CANCELLED: A previously scheduled collect or contact whose execution was
	// cancelled.
	//
	// COLLECTED: The collect has been accomplished. A collected state implies that
	// additional activity is required for delivery/completion.
	//
	// COMPLETED: The collect or contact has been completed. For many systems completed
	// and delivered constitute an equivalent successful end state.
	//
	// DELIVERED: The collected observation(s) have been delivered to the requestor.
	// For many systems completed and delivered constitute an equivalent successful end
	// state. A DELIVERED state is typically used for systems that exhibit a delay
	// between collect and delivery, such as with space-based systems which require
	// ground contact to deliver observations.
	//
	// FAILED: The collect or contact was attempted and failed, or the delivery of the
	// collected observation(s) failed. A FAILED status may be accompanied by an error
	// code (errCode), if available.
	//
	// PARTIAL: A PARTIAL state indicates that a part of a multi-track request has been
	// accomplished, but the full request is incomplete. A PARTIAL status should
	// ultimately be resolved to an end state.
	//
	// PROPOSED: Indicates that the request was received and alternate collect or
	// contact time(s) (altStartTime, altEndTime) have been proposed. If an alternate
	// is accepted by the requestor the current request should be cancelled and a new
	// request created.
	//
	// REJECTED: The request has been received and rejected by the provider. A REJECTED
	// status may be accompanied by an explanation (notes) of the reason that the
	// request was rejected.
	//
	// SCHEDULED: The request was received and has been scheduled for execution.
	Status param.Opt[string] `json:"status,omitzero"`
	// Optional task ID associated with the request/response.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Array of UUIDs of the UDL data record(s) collected in response to the associated
	// request. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. The appropriate API operation can be used to
	// retrieve the specified object(s) (e.g. /udl/rfobservation/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record type(s) (DOA, ELSET, EO, RADAR, RF, SV) collected or
	// produced in response to the associated request. See the associated 'srcIds'
	// array for the record UUIDs, positionally corresponding to the record types in
	// this array. The 'srcTyps' and 'srcIds' arrays must match in size. The
	// appropriate API operation can be used to retrieve the specified object(s) (e.g.
	// /udl/rfobservation/{uuid}).
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r CollectResponseUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow CollectResponseUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectResponseUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectResponseUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
