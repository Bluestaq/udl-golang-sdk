// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
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

// EffectRequestService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEffectRequestService] method instead.
type EffectRequestService struct {
	Options []option.RequestOption
	History EffectRequestHistoryService
}

// NewEffectRequestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEffectRequestService(opts ...option.RequestOption) (r EffectRequestService) {
	r = EffectRequestService{}
	r.Options = opts
	r.History = NewEffectRequestHistoryService(opts...)
	return
}

// Service operation to take a single EffectRequest as a POST body and ingest into
// the database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *EffectRequestService) New(ctx context.Context, body EffectRequestNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/effectrequest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single EffectRequest by its unique ID passed as a
// path parameter.
func (r *EffectRequestService) Get(ctx context.Context, id string, query EffectRequestGetParams, opts ...option.RequestOption) (res *EffectRequestGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/effectrequest/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EffectRequestService) List(ctx context.Context, query EffectRequestListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EffectRequestListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/effectrequest"
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
func (r *EffectRequestService) ListAutoPaging(ctx context.Context, query EffectRequestListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EffectRequestListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EffectRequestService) Count(ctx context.Context, query EffectRequestCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/effectrequest/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// EffectRequest records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *EffectRequestService) NewBulk(ctx context.Context, body EffectRequestNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/effectrequest/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EffectRequestService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *EffectRequestQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/effectrequest/queryhelp"
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
func (r *EffectRequestService) Tuple(ctx context.Context, query EffectRequestTupleParams, opts ...option.RequestOption) (res *[]EffectRequestTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/effectrequest/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple EffectRequests as a POST body and ingest into
// the database. This operation is intended to be used for automated feeds into
// UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *EffectRequestService) UnvalidatedPublish(ctx context.Context, body EffectRequestUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-effectrequest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// A request for various effects on a target.
type EffectRequestGetResponse struct {
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
	DataMode EffectRequestGetResponseDataMode `json:"dataMode,required"`
	// List of effects to be achieved on the target (e.g. COVER, DECEIVE, DEGRADE,
	// DENY, DESTROY, DISRUPT, DIVERSION, DIVERT, FIX, INSPECT, INTERCEPT, ISOLATE,
	// MANIPULATE, NEUTRALIZE, SHADOW, SUPPRESS, etc.). The effects included in this
	// list are connected by implied AND.
	EffectList []string `json:"effectList,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Specific descriptive instantiation of the effect, e.g., playbook to be used.
	Context string `json:"context"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The indicator of deadline of the bid request (e.g. BETWEEN, IMMEDIATE,
	// NOEARLIERTHAN, NOLATERTHAN, etc.): BETWEEN:&nbsp;Produce effect any time between
	// the given start and end times, equal penalty for being early or late
	// IMMEDIATE:&nbsp;Start as soon as possible, earlier is always better
	// NOEARLIERTHAN:&nbsp;Produce effect at this time or later. Large penalty for
	// being earlier, no reward for being later NOLATERTHAN:&nbsp;Produce effect no
	// later than the given startTime. Large penalty for being later, no reward for
	// being even earlier as long as the effect starts by the given time.
	DeadlineType string `json:"deadlineType"`
	// The time the effect should end, in ISO8601 UTC format.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The extenal system identifier of this request. A human readable unique id.
	ExternalRequestID string `json:"externalRequestId"`
	// Array of the the metric classes to be evaluated (e.g. Cost, GoalAchievement,
	// OpportunityCost, Risk, Timeliness, Unavailable, etc.). See the associated
	// 'metricWeights' array for the weighting values, positionally corresponding to
	// these types. The 'metricTypes' and 'metricWeights' arrays must match in size.
	MetricTypes []string `json:"metricTypes"`
	// Array of the weights for the metric in the final evaluation score. Normalized (0
	// to 1). See the associated 'metricTypes' array for the metric classes,
	// positionally corresponding to these values. The 'metricTypes' and
	// 'metricWeights' arrays must match in size.
	MetricWeights []float64 `json:"metricWeights"`
	// The type or class of the preference model used to evaluate this offer.
	ModelClass string `json:"modelClass"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The priority (LOW, MEDIUM, HIGH) of this request.
	Priority string `json:"priority"`
	// The time the effect should start, in ISO8601 UTC format.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// State of this effect request (e.g. CREATED, UPDATED, DELETED, etc.).
	State string `json:"state"`
	// The record ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID string `json:"targetSrcId"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType string `json:"targetSrcType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EffectList            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Context               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeadlineType          respjson.Field
		EndTime               respjson.Field
		ExternalRequestID     respjson.Field
		MetricTypes           respjson.Field
		MetricWeights         respjson.Field
		ModelClass            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Priority              respjson.Field
		StartTime             respjson.Field
		State                 respjson.Field
		TargetSrcID           respjson.Field
		TargetSrcType         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectRequestGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EffectRequestGetResponse) UnmarshalJSON(data []byte) error {
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
type EffectRequestGetResponseDataMode string

const (
	EffectRequestGetResponseDataModeReal      EffectRequestGetResponseDataMode = "REAL"
	EffectRequestGetResponseDataModeTest      EffectRequestGetResponseDataMode = "TEST"
	EffectRequestGetResponseDataModeSimulated EffectRequestGetResponseDataMode = "SIMULATED"
	EffectRequestGetResponseDataModeExercise  EffectRequestGetResponseDataMode = "EXERCISE"
)

// A request for various effects on a target.
type EffectRequestListResponse struct {
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
	DataMode EffectRequestListResponseDataMode `json:"dataMode,required"`
	// List of effects to be achieved on the target (e.g. COVER, DECEIVE, DEGRADE,
	// DENY, DESTROY, DISRUPT, DIVERSION, DIVERT, FIX, INSPECT, INTERCEPT, ISOLATE,
	// MANIPULATE, NEUTRALIZE, SHADOW, SUPPRESS, etc.). The effects included in this
	// list are connected by implied AND.
	EffectList []string `json:"effectList,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Specific descriptive instantiation of the effect, e.g., playbook to be used.
	Context string `json:"context"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The indicator of deadline of the bid request (e.g. BETWEEN, IMMEDIATE,
	// NOEARLIERTHAN, NOLATERTHAN, etc.): BETWEEN:&nbsp;Produce effect any time between
	// the given start and end times, equal penalty for being early or late
	// IMMEDIATE:&nbsp;Start as soon as possible, earlier is always better
	// NOEARLIERTHAN:&nbsp;Produce effect at this time or later. Large penalty for
	// being earlier, no reward for being later NOLATERTHAN:&nbsp;Produce effect no
	// later than the given startTime. Large penalty for being later, no reward for
	// being even earlier as long as the effect starts by the given time.
	DeadlineType string `json:"deadlineType"`
	// The time the effect should end, in ISO8601 UTC format.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The extenal system identifier of this request. A human readable unique id.
	ExternalRequestID string `json:"externalRequestId"`
	// Array of the the metric classes to be evaluated (e.g. Cost, GoalAchievement,
	// OpportunityCost, Risk, Timeliness, Unavailable, etc.). See the associated
	// 'metricWeights' array for the weighting values, positionally corresponding to
	// these types. The 'metricTypes' and 'metricWeights' arrays must match in size.
	MetricTypes []string `json:"metricTypes"`
	// Array of the weights for the metric in the final evaluation score. Normalized (0
	// to 1). See the associated 'metricTypes' array for the metric classes,
	// positionally corresponding to these values. The 'metricTypes' and
	// 'metricWeights' arrays must match in size.
	MetricWeights []float64 `json:"metricWeights"`
	// The type or class of the preference model used to evaluate this offer.
	ModelClass string `json:"modelClass"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The priority (LOW, MEDIUM, HIGH) of this request.
	Priority string `json:"priority"`
	// The time the effect should start, in ISO8601 UTC format.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// State of this effect request (e.g. CREATED, UPDATED, DELETED, etc.).
	State string `json:"state"`
	// The record ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID string `json:"targetSrcId"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType string `json:"targetSrcType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EffectList            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Context               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeadlineType          respjson.Field
		EndTime               respjson.Field
		ExternalRequestID     respjson.Field
		MetricTypes           respjson.Field
		MetricWeights         respjson.Field
		ModelClass            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Priority              respjson.Field
		StartTime             respjson.Field
		State                 respjson.Field
		TargetSrcID           respjson.Field
		TargetSrcType         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectRequestListResponse) RawJSON() string { return r.JSON.raw }
func (r *EffectRequestListResponse) UnmarshalJSON(data []byte) error {
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
type EffectRequestListResponseDataMode string

const (
	EffectRequestListResponseDataModeReal      EffectRequestListResponseDataMode = "REAL"
	EffectRequestListResponseDataModeTest      EffectRequestListResponseDataMode = "TEST"
	EffectRequestListResponseDataModeSimulated EffectRequestListResponseDataMode = "SIMULATED"
	EffectRequestListResponseDataModeExercise  EffectRequestListResponseDataMode = "EXERCISE"
)

type EffectRequestQueryHelpResponse struct {
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
func (r EffectRequestQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *EffectRequestQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for various effects on a target.
type EffectRequestTupleResponse struct {
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
	DataMode EffectRequestTupleResponseDataMode `json:"dataMode,required"`
	// List of effects to be achieved on the target (e.g. COVER, DECEIVE, DEGRADE,
	// DENY, DESTROY, DISRUPT, DIVERSION, DIVERT, FIX, INSPECT, INTERCEPT, ISOLATE,
	// MANIPULATE, NEUTRALIZE, SHADOW, SUPPRESS, etc.). The effects included in this
	// list are connected by implied AND.
	EffectList []string `json:"effectList,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Specific descriptive instantiation of the effect, e.g., playbook to be used.
	Context string `json:"context"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The indicator of deadline of the bid request (e.g. BETWEEN, IMMEDIATE,
	// NOEARLIERTHAN, NOLATERTHAN, etc.): BETWEEN:&nbsp;Produce effect any time between
	// the given start and end times, equal penalty for being early or late
	// IMMEDIATE:&nbsp;Start as soon as possible, earlier is always better
	// NOEARLIERTHAN:&nbsp;Produce effect at this time or later. Large penalty for
	// being earlier, no reward for being later NOLATERTHAN:&nbsp;Produce effect no
	// later than the given startTime. Large penalty for being later, no reward for
	// being even earlier as long as the effect starts by the given time.
	DeadlineType string `json:"deadlineType"`
	// The time the effect should end, in ISO8601 UTC format.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The extenal system identifier of this request. A human readable unique id.
	ExternalRequestID string `json:"externalRequestId"`
	// Array of the the metric classes to be evaluated (e.g. Cost, GoalAchievement,
	// OpportunityCost, Risk, Timeliness, Unavailable, etc.). See the associated
	// 'metricWeights' array for the weighting values, positionally corresponding to
	// these types. The 'metricTypes' and 'metricWeights' arrays must match in size.
	MetricTypes []string `json:"metricTypes"`
	// Array of the weights for the metric in the final evaluation score. Normalized (0
	// to 1). See the associated 'metricTypes' array for the metric classes,
	// positionally corresponding to these values. The 'metricTypes' and
	// 'metricWeights' arrays must match in size.
	MetricWeights []float64 `json:"metricWeights"`
	// The type or class of the preference model used to evaluate this offer.
	ModelClass string `json:"modelClass"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The priority (LOW, MEDIUM, HIGH) of this request.
	Priority string `json:"priority"`
	// The time the effect should start, in ISO8601 UTC format.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// State of this effect request (e.g. CREATED, UPDATED, DELETED, etc.).
	State string `json:"state"`
	// The record ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID string `json:"targetSrcId"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType string `json:"targetSrcType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EffectList            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Context               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeadlineType          respjson.Field
		EndTime               respjson.Field
		ExternalRequestID     respjson.Field
		MetricTypes           respjson.Field
		MetricWeights         respjson.Field
		ModelClass            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Priority              respjson.Field
		StartTime             respjson.Field
		State                 respjson.Field
		TargetSrcID           respjson.Field
		TargetSrcType         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectRequestTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *EffectRequestTupleResponse) UnmarshalJSON(data []byte) error {
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
type EffectRequestTupleResponseDataMode string

const (
	EffectRequestTupleResponseDataModeReal      EffectRequestTupleResponseDataMode = "REAL"
	EffectRequestTupleResponseDataModeTest      EffectRequestTupleResponseDataMode = "TEST"
	EffectRequestTupleResponseDataModeSimulated EffectRequestTupleResponseDataMode = "SIMULATED"
	EffectRequestTupleResponseDataModeExercise  EffectRequestTupleResponseDataMode = "EXERCISE"
)

type EffectRequestNewParams struct {
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
	DataMode EffectRequestNewParamsDataMode `json:"dataMode,omitzero,required"`
	// List of effects to be achieved on the target (e.g. COVER, DECEIVE, DEGRADE,
	// DENY, DESTROY, DISRUPT, DIVERSION, DIVERT, FIX, INSPECT, INTERCEPT, ISOLATE,
	// MANIPULATE, NEUTRALIZE, SHADOW, SUPPRESS, etc.). The effects included in this
	// list are connected by implied AND.
	EffectList []string `json:"effectList,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Specific descriptive instantiation of the effect, e.g., playbook to be used.
	Context param.Opt[string] `json:"context,omitzero"`
	// The indicator of deadline of the bid request (e.g. BETWEEN, IMMEDIATE,
	// NOEARLIERTHAN, NOLATERTHAN, etc.): BETWEEN:&nbsp;Produce effect any time between
	// the given start and end times, equal penalty for being early or late
	// IMMEDIATE:&nbsp;Start as soon as possible, earlier is always better
	// NOEARLIERTHAN:&nbsp;Produce effect at this time or later. Large penalty for
	// being earlier, no reward for being later NOLATERTHAN:&nbsp;Produce effect no
	// later than the given startTime. Large penalty for being later, no reward for
	// being even earlier as long as the effect starts by the given time.
	DeadlineType param.Opt[string] `json:"deadlineType,omitzero"`
	// The time the effect should end, in ISO8601 UTC format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The extenal system identifier of this request. A human readable unique id.
	ExternalRequestID param.Opt[string] `json:"externalRequestId,omitzero"`
	// The type or class of the preference model used to evaluate this offer.
	ModelClass param.Opt[string] `json:"modelClass,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The priority (LOW, MEDIUM, HIGH) of this request.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The time the effect should start, in ISO8601 UTC format.
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// State of this effect request (e.g. CREATED, UPDATED, DELETED, etc.).
	State param.Opt[string] `json:"state,omitzero"`
	// The record ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID param.Opt[string] `json:"targetSrcId,omitzero"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType param.Opt[string] `json:"targetSrcType,omitzero"`
	// Array of the the metric classes to be evaluated (e.g. Cost, GoalAchievement,
	// OpportunityCost, Risk, Timeliness, Unavailable, etc.). See the associated
	// 'metricWeights' array for the weighting values, positionally corresponding to
	// these types. The 'metricTypes' and 'metricWeights' arrays must match in size.
	MetricTypes []string `json:"metricTypes,omitzero"`
	// Array of the weights for the metric in the final evaluation score. Normalized (0
	// to 1). See the associated 'metricTypes' array for the metric classes,
	// positionally corresponding to these values. The 'metricTypes' and
	// 'metricWeights' arrays must match in size.
	MetricWeights []float64 `json:"metricWeights,omitzero"`
	paramObj
}

func (r EffectRequestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EffectRequestNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectRequestNewParams) UnmarshalJSON(data []byte) error {
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
type EffectRequestNewParamsDataMode string

const (
	EffectRequestNewParamsDataModeReal      EffectRequestNewParamsDataMode = "REAL"
	EffectRequestNewParamsDataModeTest      EffectRequestNewParamsDataMode = "TEST"
	EffectRequestNewParamsDataModeSimulated EffectRequestNewParamsDataMode = "SIMULATED"
	EffectRequestNewParamsDataModeExercise  EffectRequestNewParamsDataMode = "EXERCISE"
)

type EffectRequestGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EffectRequestGetParams]'s query parameters as `url.Values`.
func (r EffectRequestGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectRequestListParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EffectRequestListParams]'s query parameters as
// `url.Values`.
func (r EffectRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectRequestCountParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EffectRequestCountParams]'s query parameters as
// `url.Values`.
func (r EffectRequestCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectRequestNewBulkParams struct {
	Body []EffectRequestNewBulkParamsBody
	paramObj
}

func (r EffectRequestNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *EffectRequestNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// A request for various effects on a target.
//
// The properties ClassificationMarking, DataMode, EffectList, Source are required.
type EffectRequestNewBulkParamsBody struct {
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
	// List of effects to be achieved on the target (e.g. COVER, DECEIVE, DEGRADE,
	// DENY, DESTROY, DISRUPT, DIVERSION, DIVERT, FIX, INSPECT, INTERCEPT, ISOLATE,
	// MANIPULATE, NEUTRALIZE, SHADOW, SUPPRESS, etc.). The effects included in this
	// list are connected by implied AND.
	EffectList []string `json:"effectList,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Specific descriptive instantiation of the effect, e.g., playbook to be used.
	Context param.Opt[string] `json:"context,omitzero"`
	// The indicator of deadline of the bid request (e.g. BETWEEN, IMMEDIATE,
	// NOEARLIERTHAN, NOLATERTHAN, etc.): BETWEEN:&nbsp;Produce effect any time between
	// the given start and end times, equal penalty for being early or late
	// IMMEDIATE:&nbsp;Start as soon as possible, earlier is always better
	// NOEARLIERTHAN:&nbsp;Produce effect at this time or later. Large penalty for
	// being earlier, no reward for being later NOLATERTHAN:&nbsp;Produce effect no
	// later than the given startTime. Large penalty for being later, no reward for
	// being even earlier as long as the effect starts by the given time.
	DeadlineType param.Opt[string] `json:"deadlineType,omitzero"`
	// The time the effect should end, in ISO8601 UTC format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The extenal system identifier of this request. A human readable unique id.
	ExternalRequestID param.Opt[string] `json:"externalRequestId,omitzero"`
	// The type or class of the preference model used to evaluate this offer.
	ModelClass param.Opt[string] `json:"modelClass,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The priority (LOW, MEDIUM, HIGH) of this request.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The time the effect should start, in ISO8601 UTC format.
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// State of this effect request (e.g. CREATED, UPDATED, DELETED, etc.).
	State param.Opt[string] `json:"state,omitzero"`
	// The record ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID param.Opt[string] `json:"targetSrcId,omitzero"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType param.Opt[string] `json:"targetSrcType,omitzero"`
	// Array of the the metric classes to be evaluated (e.g. Cost, GoalAchievement,
	// OpportunityCost, Risk, Timeliness, Unavailable, etc.). See the associated
	// 'metricWeights' array for the weighting values, positionally corresponding to
	// these types. The 'metricTypes' and 'metricWeights' arrays must match in size.
	MetricTypes []string `json:"metricTypes,omitzero"`
	// Array of the weights for the metric in the final evaluation score. Normalized (0
	// to 1). See the associated 'metricTypes' array for the metric classes,
	// positionally corresponding to these values. The 'metricTypes' and
	// 'metricWeights' arrays must match in size.
	MetricWeights []float64 `json:"metricWeights,omitzero"`
	paramObj
}

func (r EffectRequestNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EffectRequestNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectRequestNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EffectRequestNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type EffectRequestTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EffectRequestTupleParams]'s query parameters as
// `url.Values`.
func (r EffectRequestTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectRequestUnvalidatedPublishParams struct {
	Body []EffectRequestUnvalidatedPublishParamsBody
	paramObj
}

func (r EffectRequestUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *EffectRequestUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// A request for various effects on a target.
//
// The properties ClassificationMarking, DataMode, EffectList, Source are required.
type EffectRequestUnvalidatedPublishParamsBody struct {
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
	// List of effects to be achieved on the target (e.g. COVER, DECEIVE, DEGRADE,
	// DENY, DESTROY, DISRUPT, DIVERSION, DIVERT, FIX, INSPECT, INTERCEPT, ISOLATE,
	// MANIPULATE, NEUTRALIZE, SHADOW, SUPPRESS, etc.). The effects included in this
	// list are connected by implied AND.
	EffectList []string `json:"effectList,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Specific descriptive instantiation of the effect, e.g., playbook to be used.
	Context param.Opt[string] `json:"context,omitzero"`
	// The indicator of deadline of the bid request (e.g. BETWEEN, IMMEDIATE,
	// NOEARLIERTHAN, NOLATERTHAN, etc.): BETWEEN:&nbsp;Produce effect any time between
	// the given start and end times, equal penalty for being early or late
	// IMMEDIATE:&nbsp;Start as soon as possible, earlier is always better
	// NOEARLIERTHAN:&nbsp;Produce effect at this time or later. Large penalty for
	// being earlier, no reward for being later NOLATERTHAN:&nbsp;Produce effect no
	// later than the given startTime. Large penalty for being later, no reward for
	// being even earlier as long as the effect starts by the given time.
	DeadlineType param.Opt[string] `json:"deadlineType,omitzero"`
	// The time the effect should end, in ISO8601 UTC format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The extenal system identifier of this request. A human readable unique id.
	ExternalRequestID param.Opt[string] `json:"externalRequestId,omitzero"`
	// The type or class of the preference model used to evaluate this offer.
	ModelClass param.Opt[string] `json:"modelClass,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The priority (LOW, MEDIUM, HIGH) of this request.
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The time the effect should start, in ISO8601 UTC format.
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// State of this effect request (e.g. CREATED, UPDATED, DELETED, etc.).
	State param.Opt[string] `json:"state,omitzero"`
	// The record ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID param.Opt[string] `json:"targetSrcId,omitzero"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType param.Opt[string] `json:"targetSrcType,omitzero"`
	// Array of the the metric classes to be evaluated (e.g. Cost, GoalAchievement,
	// OpportunityCost, Risk, Timeliness, Unavailable, etc.). See the associated
	// 'metricWeights' array for the weighting values, positionally corresponding to
	// these types. The 'metricTypes' and 'metricWeights' arrays must match in size.
	MetricTypes []string `json:"metricTypes,omitzero"`
	// Array of the weights for the metric in the final evaluation score. Normalized (0
	// to 1). See the associated 'metricTypes' array for the metric classes,
	// positionally corresponding to these values. The 'metricTypes' and
	// 'metricWeights' arrays must match in size.
	MetricWeights []float64 `json:"metricWeights,omitzero"`
	paramObj
}

func (r EffectRequestUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EffectRequestUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectRequestUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EffectRequestUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
