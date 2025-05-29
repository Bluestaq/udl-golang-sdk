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

// EffectResponseService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEffectResponseService] method instead.
type EffectResponseService struct {
	Options []option.RequestOption
	History EffectResponseHistoryService
}

// NewEffectResponseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEffectResponseService(opts ...option.RequestOption) (r EffectResponseService) {
	r = EffectResponseService{}
	r.Options = opts
	r.History = NewEffectResponseHistoryService(opts...)
	return
}

// Service operation to take a single EffectResponse as a POST body and ingest into
// the database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *EffectResponseService) New(ctx context.Context, body EffectResponseNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/effectresponse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single EffectResponse by its unique ID passed as a
// path parameter.
func (r *EffectResponseService) Get(ctx context.Context, id string, query EffectResponseGetParams, opts ...option.RequestOption) (res *EffectResponseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/effectresponse/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EffectResponseService) List(ctx context.Context, query EffectResponseListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EffectResponseListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/effectresponse"
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
func (r *EffectResponseService) ListAutoPaging(ctx context.Context, query EffectResponseListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EffectResponseListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EffectResponseService) Count(ctx context.Context, query EffectResponseCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/effectresponse/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// EffectResponse records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *EffectResponseService) NewBulk(ctx context.Context, body EffectResponseNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/effectresponse/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EffectResponseService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *EffectResponseQueryHelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/effectresponse/queryhelp"
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
func (r *EffectResponseService) Tuple(ctx context.Context, query EffectResponseTupleParams, opts ...option.RequestOption) (res *[]EffectResponseTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/effectresponse/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple EffectResponses as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *EffectResponseService) UnvalidatedPublish(ctx context.Context, body EffectResponseUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-effectresponse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// A response for various effects on a target.
type EffectResponseGetResponse struct {
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
	DataMode EffectResponseGetResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of response in this record (e.g. COA, SCORECARD, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// List of actions associated with this effect response.
	ActionsList []EffectResponseGetResponseActionsList `json:"actionsList"`
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset.
	ActorSrcID string `json:"actorSrcId"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActorSrcType string `json:"actorSrcType"`
	// List of COA metrics associated with this effect response.
	CoaMetrics []EffectResponseGetResponseCoaMetric `json:"coaMetrics"`
	// The collateral damage estimate (CDE) of the munition being fired.
	CollateralDamageEst float64 `json:"collateralDamageEst"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The deadline time to accept this COA before it's no longer valid, in ISO8601 UTC
	// format.
	DecisionDeadline time.Time `json:"decisionDeadline" format:"date-time"`
	// List of external actions to be executed as part of this task.
	ExternalActions []string `json:"externalActions"`
	// The external system identifier of the associated effect request. A human
	// readable unique id.
	ExternalRequestID string `json:"externalRequestId"`
	// Unique identifier of the EffectRequest associated with this response.
	IDEffectRequest string `json:"idEffectRequest"`
	// Unique identifier of the munition.
	MunitionID string `json:"munitionId"`
	// The type of munition being fired.
	MunitionType string `json:"munitionType"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The probability of kill (0-1) of the target being destroyed.
	ProbabilityOfKill float64 `json:"probabilityOfKill"`
	// The record ID, depending on the type identified in redTargetSrcType, of the red
	// force target. If the redTargetSrcType is POI or TRACK, then this identifier
	// corresponds to either poi.poiid or track.trkId from their respective schemas.
	RedTargetSrcID string `json:"redTargetSrcId"`
	// The source type of the targetId identifier (POI, SITE, TRACK).
	RedTargetSrcType string `json:"redTargetSrcType"`
	// The time to overhead for the red force to be over their target, in ISO8601 UTC
	// format.
	RedTimeToOverhead time.Time `json:"redTimeToOverhead" format:"date-time"`
	// The number of shots required to destroy target.
	ShotsRequired int64 `json:"shotsRequired"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		ActionsList           respjson.Field
		ActorSrcID            respjson.Field
		ActorSrcType          respjson.Field
		CoaMetrics            respjson.Field
		CollateralDamageEst   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecisionDeadline      respjson.Field
		ExternalActions       respjson.Field
		ExternalRequestID     respjson.Field
		IDEffectRequest       respjson.Field
		MunitionID            respjson.Field
		MunitionType          respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ProbabilityOfKill     respjson.Field
		RedTargetSrcID        respjson.Field
		RedTargetSrcType      respjson.Field
		RedTimeToOverhead     respjson.Field
		ShotsRequired         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseGetResponse) UnmarshalJSON(data []byte) error {
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
type EffectResponseGetResponseDataMode string

const (
	EffectResponseGetResponseDataModeReal      EffectResponseGetResponseDataMode = "REAL"
	EffectResponseGetResponseDataModeTest      EffectResponseGetResponseDataMode = "TEST"
	EffectResponseGetResponseDataModeSimulated EffectResponseGetResponseDataMode = "SIMULATED"
	EffectResponseGetResponseDataModeExercise  EffectResponseGetResponseDataMode = "EXERCISE"
)

type EffectResponseGetResponseActionsList struct {
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset/actor.
	ActionActorSrcID string `json:"actionActorSrcId"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActionActorSrcType string `json:"actionActorSrcType"`
	// The desired end time of this task, in ISO8601 UTC format.
	ActionEndTime time.Time `json:"actionEndTime" format:"date-time"`
	// Identifier of this action.
	ActionID string `json:"actionId"`
	// List of metrics associated with this action.
	ActionMetrics []EffectResponseGetResponseActionsListActionMetric `json:"actionMetrics"`
	// The desired start time of this task, in ISO8601 UTC format.
	ActionStartTime time.Time `json:"actionStartTime" format:"date-time"`
	// The WGS-84 altitude of the asset/actor location at weapon launch, in meters.
	ActorInterceptAlt float64 `json:"actorInterceptAlt"`
	// The WGS-84 latitude of the asset/actor location at weapon launch, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	ActorInterceptLat float64 `json:"actorInterceptLat"`
	// The WGS-84 longitude of the asset/actor location at weapon launch, in degrees.
	// -180 to 180 degrees (negative values west of Prime Meridian).
	ActorInterceptLon float64 `json:"actorInterceptLon"`
	// The type of munition or sensor used by this asset/actor.
	Effector string `json:"effector"`
	// A summary string describing different aspects of the action.
	Summary string `json:"summary"`
	// The POI or TRACK ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID string `json:"targetSrcId"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType string `json:"targetSrcType"`
	// The end time of the asset TOT (time over target), in ISO8601 UTC format.
	TotEndTime time.Time `json:"totEndTime" format:"date-time"`
	// The start time of the asset TOT (time over target), in ISO8601 UTC format.
	TotStartTime time.Time `json:"totStartTime" format:"date-time"`
	// The WGS-84 altitude of the weapon destination location, in meters.
	WeaponInterceptAlt float64 `json:"weaponInterceptAlt"`
	// The WGS-84 latitude of the weapon destination location, in degrees. -90 to 90
	// degrees (negative values south of equator).
	WeaponInterceptLat float64 `json:"weaponInterceptLat"`
	// The WGS-84 longitude of the weapon destination location, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	WeaponInterceptLon float64 `json:"weaponInterceptLon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionActorSrcID   respjson.Field
		ActionActorSrcType respjson.Field
		ActionEndTime      respjson.Field
		ActionID           respjson.Field
		ActionMetrics      respjson.Field
		ActionStartTime    respjson.Field
		ActorInterceptAlt  respjson.Field
		ActorInterceptLat  respjson.Field
		ActorInterceptLon  respjson.Field
		Effector           respjson.Field
		Summary            respjson.Field
		TargetSrcID        respjson.Field
		TargetSrcType      respjson.Field
		TotEndTime         respjson.Field
		TotStartTime       respjson.Field
		WeaponInterceptAlt respjson.Field
		WeaponInterceptLat respjson.Field
		WeaponInterceptLon respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseGetResponseActionsList) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseGetResponseActionsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseGetResponseActionsListActionMetric struct {
	// The metric score specific to its domain.
	DomainValue float64 `json:"domainValue"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType string `json:"metricType"`
	// The metric that was used to score this task.
	Provenance string `json:"provenance"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue float64 `json:"relativeValue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DomainValue   respjson.Field
		MetricType    respjson.Field
		Provenance    respjson.Field
		RelativeValue respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseGetResponseActionsListActionMetric) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseGetResponseActionsListActionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseGetResponseCoaMetric struct {
	// The metric score specific to its domain.
	DomainValue float64 `json:"domainValue"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType string `json:"metricType"`
	// The metric that was used to score this task.
	Provenance string `json:"provenance"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue float64 `json:"relativeValue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DomainValue   respjson.Field
		MetricType    respjson.Field
		Provenance    respjson.Field
		RelativeValue respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseGetResponseCoaMetric) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseGetResponseCoaMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response for various effects on a target.
type EffectResponseListResponse struct {
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
	DataMode EffectResponseListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of response in this record (e.g. COA, SCORECARD, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// List of actions associated with this effect response.
	ActionsList []EffectResponseListResponseActionsList `json:"actionsList"`
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset.
	ActorSrcID string `json:"actorSrcId"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActorSrcType string `json:"actorSrcType"`
	// List of COA metrics associated with this effect response.
	CoaMetrics []EffectResponseListResponseCoaMetric `json:"coaMetrics"`
	// The collateral damage estimate (CDE) of the munition being fired.
	CollateralDamageEst float64 `json:"collateralDamageEst"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The deadline time to accept this COA before it's no longer valid, in ISO8601 UTC
	// format.
	DecisionDeadline time.Time `json:"decisionDeadline" format:"date-time"`
	// List of external actions to be executed as part of this task.
	ExternalActions []string `json:"externalActions"`
	// The external system identifier of the associated effect request. A human
	// readable unique id.
	ExternalRequestID string `json:"externalRequestId"`
	// Unique identifier of the EffectRequest associated with this response.
	IDEffectRequest string `json:"idEffectRequest"`
	// Unique identifier of the munition.
	MunitionID string `json:"munitionId"`
	// The type of munition being fired.
	MunitionType string `json:"munitionType"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The probability of kill (0-1) of the target being destroyed.
	ProbabilityOfKill float64 `json:"probabilityOfKill"`
	// The record ID, depending on the type identified in redTargetSrcType, of the red
	// force target. If the redTargetSrcType is POI or TRACK, then this identifier
	// corresponds to either poi.poiid or track.trkId from their respective schemas.
	RedTargetSrcID string `json:"redTargetSrcId"`
	// The source type of the targetId identifier (POI, SITE, TRACK).
	RedTargetSrcType string `json:"redTargetSrcType"`
	// The time to overhead for the red force to be over their target, in ISO8601 UTC
	// format.
	RedTimeToOverhead time.Time `json:"redTimeToOverhead" format:"date-time"`
	// The number of shots required to destroy target.
	ShotsRequired int64 `json:"shotsRequired"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		ActionsList           respjson.Field
		ActorSrcID            respjson.Field
		ActorSrcType          respjson.Field
		CoaMetrics            respjson.Field
		CollateralDamageEst   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecisionDeadline      respjson.Field
		ExternalActions       respjson.Field
		ExternalRequestID     respjson.Field
		IDEffectRequest       respjson.Field
		MunitionID            respjson.Field
		MunitionType          respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ProbabilityOfKill     respjson.Field
		RedTargetSrcID        respjson.Field
		RedTargetSrcType      respjson.Field
		RedTimeToOverhead     respjson.Field
		ShotsRequired         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseListResponse) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseListResponse) UnmarshalJSON(data []byte) error {
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
type EffectResponseListResponseDataMode string

const (
	EffectResponseListResponseDataModeReal      EffectResponseListResponseDataMode = "REAL"
	EffectResponseListResponseDataModeTest      EffectResponseListResponseDataMode = "TEST"
	EffectResponseListResponseDataModeSimulated EffectResponseListResponseDataMode = "SIMULATED"
	EffectResponseListResponseDataModeExercise  EffectResponseListResponseDataMode = "EXERCISE"
)

type EffectResponseListResponseActionsList struct {
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset/actor.
	ActionActorSrcID string `json:"actionActorSrcId"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActionActorSrcType string `json:"actionActorSrcType"`
	// The desired end time of this task, in ISO8601 UTC format.
	ActionEndTime time.Time `json:"actionEndTime" format:"date-time"`
	// Identifier of this action.
	ActionID string `json:"actionId"`
	// List of metrics associated with this action.
	ActionMetrics []EffectResponseListResponseActionsListActionMetric `json:"actionMetrics"`
	// The desired start time of this task, in ISO8601 UTC format.
	ActionStartTime time.Time `json:"actionStartTime" format:"date-time"`
	// The WGS-84 altitude of the asset/actor location at weapon launch, in meters.
	ActorInterceptAlt float64 `json:"actorInterceptAlt"`
	// The WGS-84 latitude of the asset/actor location at weapon launch, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	ActorInterceptLat float64 `json:"actorInterceptLat"`
	// The WGS-84 longitude of the asset/actor location at weapon launch, in degrees.
	// -180 to 180 degrees (negative values west of Prime Meridian).
	ActorInterceptLon float64 `json:"actorInterceptLon"`
	// The type of munition or sensor used by this asset/actor.
	Effector string `json:"effector"`
	// A summary string describing different aspects of the action.
	Summary string `json:"summary"`
	// The POI or TRACK ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID string `json:"targetSrcId"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType string `json:"targetSrcType"`
	// The end time of the asset TOT (time over target), in ISO8601 UTC format.
	TotEndTime time.Time `json:"totEndTime" format:"date-time"`
	// The start time of the asset TOT (time over target), in ISO8601 UTC format.
	TotStartTime time.Time `json:"totStartTime" format:"date-time"`
	// The WGS-84 altitude of the weapon destination location, in meters.
	WeaponInterceptAlt float64 `json:"weaponInterceptAlt"`
	// The WGS-84 latitude of the weapon destination location, in degrees. -90 to 90
	// degrees (negative values south of equator).
	WeaponInterceptLat float64 `json:"weaponInterceptLat"`
	// The WGS-84 longitude of the weapon destination location, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	WeaponInterceptLon float64 `json:"weaponInterceptLon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionActorSrcID   respjson.Field
		ActionActorSrcType respjson.Field
		ActionEndTime      respjson.Field
		ActionID           respjson.Field
		ActionMetrics      respjson.Field
		ActionStartTime    respjson.Field
		ActorInterceptAlt  respjson.Field
		ActorInterceptLat  respjson.Field
		ActorInterceptLon  respjson.Field
		Effector           respjson.Field
		Summary            respjson.Field
		TargetSrcID        respjson.Field
		TargetSrcType      respjson.Field
		TotEndTime         respjson.Field
		TotStartTime       respjson.Field
		WeaponInterceptAlt respjson.Field
		WeaponInterceptLat respjson.Field
		WeaponInterceptLon respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseListResponseActionsList) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseListResponseActionsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseListResponseActionsListActionMetric struct {
	// The metric score specific to its domain.
	DomainValue float64 `json:"domainValue"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType string `json:"metricType"`
	// The metric that was used to score this task.
	Provenance string `json:"provenance"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue float64 `json:"relativeValue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DomainValue   respjson.Field
		MetricType    respjson.Field
		Provenance    respjson.Field
		RelativeValue respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseListResponseActionsListActionMetric) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseListResponseActionsListActionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseListResponseCoaMetric struct {
	// The metric score specific to its domain.
	DomainValue float64 `json:"domainValue"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType string `json:"metricType"`
	// The metric that was used to score this task.
	Provenance string `json:"provenance"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue float64 `json:"relativeValue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DomainValue   respjson.Field
		MetricType    respjson.Field
		Provenance    respjson.Field
		RelativeValue respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseListResponseCoaMetric) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseListResponseCoaMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseQueryHelpResponse struct {
	AodrSupported         bool                                       `json:"aodrSupported"`
	ClassificationMarking string                                     `json:"classificationMarking"`
	Description           string                                     `json:"description"`
	HistorySupported      bool                                       `json:"historySupported"`
	Name                  string                                     `json:"name"`
	Parameters            []EffectResponseQueryHelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                   `json:"requiredRoles"`
	RestSupported         bool                                       `json:"restSupported"`
	SortSupported         bool                                       `json:"sortSupported"`
	TypeName              string                                     `json:"typeName"`
	Uri                   string                                     `json:"uri"`
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
func (r EffectResponseQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseQueryHelpResponseParameter struct {
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
func (r EffectResponseQueryHelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseQueryHelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response for various effects on a target.
type EffectResponseTupleResponse struct {
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
	DataMode EffectResponseTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of response in this record (e.g. COA, SCORECARD, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// List of actions associated with this effect response.
	ActionsList []EffectResponseTupleResponseActionsList `json:"actionsList"`
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset.
	ActorSrcID string `json:"actorSrcId"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActorSrcType string `json:"actorSrcType"`
	// List of COA metrics associated with this effect response.
	CoaMetrics []EffectResponseTupleResponseCoaMetric `json:"coaMetrics"`
	// The collateral damage estimate (CDE) of the munition being fired.
	CollateralDamageEst float64 `json:"collateralDamageEst"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The deadline time to accept this COA before it's no longer valid, in ISO8601 UTC
	// format.
	DecisionDeadline time.Time `json:"decisionDeadline" format:"date-time"`
	// List of external actions to be executed as part of this task.
	ExternalActions []string `json:"externalActions"`
	// The external system identifier of the associated effect request. A human
	// readable unique id.
	ExternalRequestID string `json:"externalRequestId"`
	// Unique identifier of the EffectRequest associated with this response.
	IDEffectRequest string `json:"idEffectRequest"`
	// Unique identifier of the munition.
	MunitionID string `json:"munitionId"`
	// The type of munition being fired.
	MunitionType string `json:"munitionType"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The probability of kill (0-1) of the target being destroyed.
	ProbabilityOfKill float64 `json:"probabilityOfKill"`
	// The record ID, depending on the type identified in redTargetSrcType, of the red
	// force target. If the redTargetSrcType is POI or TRACK, then this identifier
	// corresponds to either poi.poiid or track.trkId from their respective schemas.
	RedTargetSrcID string `json:"redTargetSrcId"`
	// The source type of the targetId identifier (POI, SITE, TRACK).
	RedTargetSrcType string `json:"redTargetSrcType"`
	// The time to overhead for the red force to be over their target, in ISO8601 UTC
	// format.
	RedTimeToOverhead time.Time `json:"redTimeToOverhead" format:"date-time"`
	// The number of shots required to destroy target.
	ShotsRequired int64 `json:"shotsRequired"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		ActionsList           respjson.Field
		ActorSrcID            respjson.Field
		ActorSrcType          respjson.Field
		CoaMetrics            respjson.Field
		CollateralDamageEst   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecisionDeadline      respjson.Field
		ExternalActions       respjson.Field
		ExternalRequestID     respjson.Field
		IDEffectRequest       respjson.Field
		MunitionID            respjson.Field
		MunitionType          respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ProbabilityOfKill     respjson.Field
		RedTargetSrcID        respjson.Field
		RedTargetSrcType      respjson.Field
		RedTimeToOverhead     respjson.Field
		ShotsRequired         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseTupleResponse) UnmarshalJSON(data []byte) error {
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
type EffectResponseTupleResponseDataMode string

const (
	EffectResponseTupleResponseDataModeReal      EffectResponseTupleResponseDataMode = "REAL"
	EffectResponseTupleResponseDataModeTest      EffectResponseTupleResponseDataMode = "TEST"
	EffectResponseTupleResponseDataModeSimulated EffectResponseTupleResponseDataMode = "SIMULATED"
	EffectResponseTupleResponseDataModeExercise  EffectResponseTupleResponseDataMode = "EXERCISE"
)

type EffectResponseTupleResponseActionsList struct {
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset/actor.
	ActionActorSrcID string `json:"actionActorSrcId"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActionActorSrcType string `json:"actionActorSrcType"`
	// The desired end time of this task, in ISO8601 UTC format.
	ActionEndTime time.Time `json:"actionEndTime" format:"date-time"`
	// Identifier of this action.
	ActionID string `json:"actionId"`
	// List of metrics associated with this action.
	ActionMetrics []EffectResponseTupleResponseActionsListActionMetric `json:"actionMetrics"`
	// The desired start time of this task, in ISO8601 UTC format.
	ActionStartTime time.Time `json:"actionStartTime" format:"date-time"`
	// The WGS-84 altitude of the asset/actor location at weapon launch, in meters.
	ActorInterceptAlt float64 `json:"actorInterceptAlt"`
	// The WGS-84 latitude of the asset/actor location at weapon launch, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	ActorInterceptLat float64 `json:"actorInterceptLat"`
	// The WGS-84 longitude of the asset/actor location at weapon launch, in degrees.
	// -180 to 180 degrees (negative values west of Prime Meridian).
	ActorInterceptLon float64 `json:"actorInterceptLon"`
	// The type of munition or sensor used by this asset/actor.
	Effector string `json:"effector"`
	// A summary string describing different aspects of the action.
	Summary string `json:"summary"`
	// The POI or TRACK ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID string `json:"targetSrcId"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType string `json:"targetSrcType"`
	// The end time of the asset TOT (time over target), in ISO8601 UTC format.
	TotEndTime time.Time `json:"totEndTime" format:"date-time"`
	// The start time of the asset TOT (time over target), in ISO8601 UTC format.
	TotStartTime time.Time `json:"totStartTime" format:"date-time"`
	// The WGS-84 altitude of the weapon destination location, in meters.
	WeaponInterceptAlt float64 `json:"weaponInterceptAlt"`
	// The WGS-84 latitude of the weapon destination location, in degrees. -90 to 90
	// degrees (negative values south of equator).
	WeaponInterceptLat float64 `json:"weaponInterceptLat"`
	// The WGS-84 longitude of the weapon destination location, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	WeaponInterceptLon float64 `json:"weaponInterceptLon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionActorSrcID   respjson.Field
		ActionActorSrcType respjson.Field
		ActionEndTime      respjson.Field
		ActionID           respjson.Field
		ActionMetrics      respjson.Field
		ActionStartTime    respjson.Field
		ActorInterceptAlt  respjson.Field
		ActorInterceptLat  respjson.Field
		ActorInterceptLon  respjson.Field
		Effector           respjson.Field
		Summary            respjson.Field
		TargetSrcID        respjson.Field
		TargetSrcType      respjson.Field
		TotEndTime         respjson.Field
		TotStartTime       respjson.Field
		WeaponInterceptAlt respjson.Field
		WeaponInterceptLat respjson.Field
		WeaponInterceptLon respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseTupleResponseActionsList) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseTupleResponseActionsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseTupleResponseActionsListActionMetric struct {
	// The metric score specific to its domain.
	DomainValue float64 `json:"domainValue"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType string `json:"metricType"`
	// The metric that was used to score this task.
	Provenance string `json:"provenance"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue float64 `json:"relativeValue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DomainValue   respjson.Field
		MetricType    respjson.Field
		Provenance    respjson.Field
		RelativeValue respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseTupleResponseActionsListActionMetric) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseTupleResponseActionsListActionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseTupleResponseCoaMetric struct {
	// The metric score specific to its domain.
	DomainValue float64 `json:"domainValue"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType string `json:"metricType"`
	// The metric that was used to score this task.
	Provenance string `json:"provenance"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue float64 `json:"relativeValue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DomainValue   respjson.Field
		MetricType    respjson.Field
		Provenance    respjson.Field
		RelativeValue respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EffectResponseTupleResponseCoaMetric) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseTupleResponseCoaMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseNewParams struct {
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
	DataMode EffectResponseNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of response in this record (e.g. COA, SCORECARD, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset.
	ActorSrcID param.Opt[string] `json:"actorSrcId,omitzero"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActorSrcType param.Opt[string] `json:"actorSrcType,omitzero"`
	// The collateral damage estimate (CDE) of the munition being fired.
	CollateralDamageEst param.Opt[float64] `json:"collateralDamageEst,omitzero"`
	// The deadline time to accept this COA before it's no longer valid, in ISO8601 UTC
	// format.
	DecisionDeadline param.Opt[time.Time] `json:"decisionDeadline,omitzero" format:"date-time"`
	// The external system identifier of the associated effect request. A human
	// readable unique id.
	ExternalRequestID param.Opt[string] `json:"externalRequestId,omitzero"`
	// Unique identifier of the EffectRequest associated with this response.
	IDEffectRequest param.Opt[string] `json:"idEffectRequest,omitzero"`
	// Unique identifier of the munition.
	MunitionID param.Opt[string] `json:"munitionId,omitzero"`
	// The type of munition being fired.
	MunitionType param.Opt[string] `json:"munitionType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The probability of kill (0-1) of the target being destroyed.
	ProbabilityOfKill param.Opt[float64] `json:"probabilityOfKill,omitzero"`
	// The record ID, depending on the type identified in redTargetSrcType, of the red
	// force target. If the redTargetSrcType is POI or TRACK, then this identifier
	// corresponds to either poi.poiid or track.trkId from their respective schemas.
	RedTargetSrcID param.Opt[string] `json:"redTargetSrcId,omitzero"`
	// The source type of the targetId identifier (POI, SITE, TRACK).
	RedTargetSrcType param.Opt[string] `json:"redTargetSrcType,omitzero"`
	// The time to overhead for the red force to be over their target, in ISO8601 UTC
	// format.
	RedTimeToOverhead param.Opt[time.Time] `json:"redTimeToOverhead,omitzero" format:"date-time"`
	// The number of shots required to destroy target.
	ShotsRequired param.Opt[int64] `json:"shotsRequired,omitzero"`
	// List of actions associated with this effect response.
	ActionsList []EffectResponseNewParamsActionsList `json:"actionsList,omitzero"`
	// List of COA metrics associated with this effect response.
	CoaMetrics []EffectResponseNewParamsCoaMetric `json:"coaMetrics,omitzero"`
	// List of external actions to be executed as part of this task.
	ExternalActions []string `json:"externalActions,omitzero"`
	paramObj
}

func (r EffectResponseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseNewParams) UnmarshalJSON(data []byte) error {
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
type EffectResponseNewParamsDataMode string

const (
	EffectResponseNewParamsDataModeReal      EffectResponseNewParamsDataMode = "REAL"
	EffectResponseNewParamsDataModeTest      EffectResponseNewParamsDataMode = "TEST"
	EffectResponseNewParamsDataModeSimulated EffectResponseNewParamsDataMode = "SIMULATED"
	EffectResponseNewParamsDataModeExercise  EffectResponseNewParamsDataMode = "EXERCISE"
)

type EffectResponseNewParamsActionsList struct {
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset/actor.
	ActionActorSrcID param.Opt[string] `json:"actionActorSrcId,omitzero"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActionActorSrcType param.Opt[string] `json:"actionActorSrcType,omitzero"`
	// The desired end time of this task, in ISO8601 UTC format.
	ActionEndTime param.Opt[time.Time] `json:"actionEndTime,omitzero" format:"date-time"`
	// Identifier of this action.
	ActionID param.Opt[string] `json:"actionId,omitzero"`
	// The desired start time of this task, in ISO8601 UTC format.
	ActionStartTime param.Opt[time.Time] `json:"actionStartTime,omitzero" format:"date-time"`
	// The WGS-84 altitude of the asset/actor location at weapon launch, in meters.
	ActorInterceptAlt param.Opt[float64] `json:"actorInterceptAlt,omitzero"`
	// The WGS-84 latitude of the asset/actor location at weapon launch, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	ActorInterceptLat param.Opt[float64] `json:"actorInterceptLat,omitzero"`
	// The WGS-84 longitude of the asset/actor location at weapon launch, in degrees.
	// -180 to 180 degrees (negative values west of Prime Meridian).
	ActorInterceptLon param.Opt[float64] `json:"actorInterceptLon,omitzero"`
	// The type of munition or sensor used by this asset/actor.
	Effector param.Opt[string] `json:"effector,omitzero"`
	// A summary string describing different aspects of the action.
	Summary param.Opt[string] `json:"summary,omitzero"`
	// The POI or TRACK ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID param.Opt[string] `json:"targetSrcId,omitzero"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType param.Opt[string] `json:"targetSrcType,omitzero"`
	// The end time of the asset TOT (time over target), in ISO8601 UTC format.
	TotEndTime param.Opt[time.Time] `json:"totEndTime,omitzero" format:"date-time"`
	// The start time of the asset TOT (time over target), in ISO8601 UTC format.
	TotStartTime param.Opt[time.Time] `json:"totStartTime,omitzero" format:"date-time"`
	// The WGS-84 altitude of the weapon destination location, in meters.
	WeaponInterceptAlt param.Opt[float64] `json:"weaponInterceptAlt,omitzero"`
	// The WGS-84 latitude of the weapon destination location, in degrees. -90 to 90
	// degrees (negative values south of equator).
	WeaponInterceptLat param.Opt[float64] `json:"weaponInterceptLat,omitzero"`
	// The WGS-84 longitude of the weapon destination location, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	WeaponInterceptLon param.Opt[float64] `json:"weaponInterceptLon,omitzero"`
	// List of metrics associated with this action.
	ActionMetrics []EffectResponseNewParamsActionsListActionMetric `json:"actionMetrics,omitzero"`
	paramObj
}

func (r EffectResponseNewParamsActionsList) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseNewParamsActionsList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseNewParamsActionsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseNewParamsActionsListActionMetric struct {
	// The metric score specific to its domain.
	DomainValue param.Opt[float64] `json:"domainValue,omitzero"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType param.Opt[string] `json:"metricType,omitzero"`
	// The metric that was used to score this task.
	Provenance param.Opt[string] `json:"provenance,omitzero"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue param.Opt[float64] `json:"relativeValue,omitzero"`
	paramObj
}

func (r EffectResponseNewParamsActionsListActionMetric) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseNewParamsActionsListActionMetric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseNewParamsActionsListActionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseNewParamsCoaMetric struct {
	// The metric score specific to its domain.
	DomainValue param.Opt[float64] `json:"domainValue,omitzero"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType param.Opt[string] `json:"metricType,omitzero"`
	// The metric that was used to score this task.
	Provenance param.Opt[string] `json:"provenance,omitzero"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue param.Opt[float64] `json:"relativeValue,omitzero"`
	paramObj
}

func (r EffectResponseNewParamsCoaMetric) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseNewParamsCoaMetric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseNewParamsCoaMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EffectResponseGetParams]'s query parameters as
// `url.Values`.
func (r EffectResponseGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectResponseListParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EffectResponseListParams]'s query parameters as
// `url.Values`.
func (r EffectResponseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectResponseCountParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EffectResponseCountParams]'s query parameters as
// `url.Values`.
func (r EffectResponseCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectResponseNewBulkParams struct {
	Body []EffectResponseNewBulkParamsBody
	paramObj
}

func (r EffectResponseNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *EffectResponseNewBulkParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// A response for various effects on a target.
//
// The properties ClassificationMarking, DataMode, Source, Type are required.
type EffectResponseNewBulkParamsBody struct {
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
	// The type of response in this record (e.g. COA, SCORECARD, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset.
	ActorSrcID param.Opt[string] `json:"actorSrcId,omitzero"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActorSrcType param.Opt[string] `json:"actorSrcType,omitzero"`
	// The collateral damage estimate (CDE) of the munition being fired.
	CollateralDamageEst param.Opt[float64] `json:"collateralDamageEst,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The deadline time to accept this COA before it's no longer valid, in ISO8601 UTC
	// format.
	DecisionDeadline param.Opt[time.Time] `json:"decisionDeadline,omitzero" format:"date-time"`
	// The external system identifier of the associated effect request. A human
	// readable unique id.
	ExternalRequestID param.Opt[string] `json:"externalRequestId,omitzero"`
	// Unique identifier of the EffectRequest associated with this response.
	IDEffectRequest param.Opt[string] `json:"idEffectRequest,omitzero"`
	// Unique identifier of the munition.
	MunitionID param.Opt[string] `json:"munitionId,omitzero"`
	// The type of munition being fired.
	MunitionType param.Opt[string] `json:"munitionType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The probability of kill (0-1) of the target being destroyed.
	ProbabilityOfKill param.Opt[float64] `json:"probabilityOfKill,omitzero"`
	// The record ID, depending on the type identified in redTargetSrcType, of the red
	// force target. If the redTargetSrcType is POI or TRACK, then this identifier
	// corresponds to either poi.poiid or track.trkId from their respective schemas.
	RedTargetSrcID param.Opt[string] `json:"redTargetSrcId,omitzero"`
	// The source type of the targetId identifier (POI, SITE, TRACK).
	RedTargetSrcType param.Opt[string] `json:"redTargetSrcType,omitzero"`
	// The time to overhead for the red force to be over their target, in ISO8601 UTC
	// format.
	RedTimeToOverhead param.Opt[time.Time] `json:"redTimeToOverhead,omitzero" format:"date-time"`
	// The number of shots required to destroy target.
	ShotsRequired param.Opt[int64] `json:"shotsRequired,omitzero"`
	// List of actions associated with this effect response.
	ActionsList []EffectResponseNewBulkParamsBodyActionsList `json:"actionsList,omitzero"`
	// List of COA metrics associated with this effect response.
	CoaMetrics []EffectResponseNewBulkParamsBodyCoaMetric `json:"coaMetrics,omitzero"`
	// List of external actions to be executed as part of this task.
	ExternalActions []string `json:"externalActions,omitzero"`
	paramObj
}

func (r EffectResponseNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EffectResponseNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type EffectResponseNewBulkParamsBodyActionsList struct {
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset/actor.
	ActionActorSrcID param.Opt[string] `json:"actionActorSrcId,omitzero"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActionActorSrcType param.Opt[string] `json:"actionActorSrcType,omitzero"`
	// The desired end time of this task, in ISO8601 UTC format.
	ActionEndTime param.Opt[time.Time] `json:"actionEndTime,omitzero" format:"date-time"`
	// Identifier of this action.
	ActionID param.Opt[string] `json:"actionId,omitzero"`
	// The desired start time of this task, in ISO8601 UTC format.
	ActionStartTime param.Opt[time.Time] `json:"actionStartTime,omitzero" format:"date-time"`
	// The WGS-84 altitude of the asset/actor location at weapon launch, in meters.
	ActorInterceptAlt param.Opt[float64] `json:"actorInterceptAlt,omitzero"`
	// The WGS-84 latitude of the asset/actor location at weapon launch, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	ActorInterceptLat param.Opt[float64] `json:"actorInterceptLat,omitzero"`
	// The WGS-84 longitude of the asset/actor location at weapon launch, in degrees.
	// -180 to 180 degrees (negative values west of Prime Meridian).
	ActorInterceptLon param.Opt[float64] `json:"actorInterceptLon,omitzero"`
	// The type of munition or sensor used by this asset/actor.
	Effector param.Opt[string] `json:"effector,omitzero"`
	// A summary string describing different aspects of the action.
	Summary param.Opt[string] `json:"summary,omitzero"`
	// The POI or TRACK ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID param.Opt[string] `json:"targetSrcId,omitzero"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType param.Opt[string] `json:"targetSrcType,omitzero"`
	// The end time of the asset TOT (time over target), in ISO8601 UTC format.
	TotEndTime param.Opt[time.Time] `json:"totEndTime,omitzero" format:"date-time"`
	// The start time of the asset TOT (time over target), in ISO8601 UTC format.
	TotStartTime param.Opt[time.Time] `json:"totStartTime,omitzero" format:"date-time"`
	// The WGS-84 altitude of the weapon destination location, in meters.
	WeaponInterceptAlt param.Opt[float64] `json:"weaponInterceptAlt,omitzero"`
	// The WGS-84 latitude of the weapon destination location, in degrees. -90 to 90
	// degrees (negative values south of equator).
	WeaponInterceptLat param.Opt[float64] `json:"weaponInterceptLat,omitzero"`
	// The WGS-84 longitude of the weapon destination location, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	WeaponInterceptLon param.Opt[float64] `json:"weaponInterceptLon,omitzero"`
	// List of metrics associated with this action.
	ActionMetrics []EffectResponseNewBulkParamsBodyActionsListActionMetric `json:"actionMetrics,omitzero"`
	paramObj
}

func (r EffectResponseNewBulkParamsBodyActionsList) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseNewBulkParamsBodyActionsList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseNewBulkParamsBodyActionsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseNewBulkParamsBodyActionsListActionMetric struct {
	// The metric score specific to its domain.
	DomainValue param.Opt[float64] `json:"domainValue,omitzero"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType param.Opt[string] `json:"metricType,omitzero"`
	// The metric that was used to score this task.
	Provenance param.Opt[string] `json:"provenance,omitzero"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue param.Opt[float64] `json:"relativeValue,omitzero"`
	paramObj
}

func (r EffectResponseNewBulkParamsBodyActionsListActionMetric) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseNewBulkParamsBodyActionsListActionMetric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseNewBulkParamsBodyActionsListActionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseNewBulkParamsBodyCoaMetric struct {
	// The metric score specific to its domain.
	DomainValue param.Opt[float64] `json:"domainValue,omitzero"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType param.Opt[string] `json:"metricType,omitzero"`
	// The metric that was used to score this task.
	Provenance param.Opt[string] `json:"provenance,omitzero"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue param.Opt[float64] `json:"relativeValue,omitzero"`
	paramObj
}

func (r EffectResponseNewBulkParamsBodyCoaMetric) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseNewBulkParamsBodyCoaMetric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseNewBulkParamsBodyCoaMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseTupleParams struct {
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

// URLQuery serializes [EffectResponseTupleParams]'s query parameters as
// `url.Values`.
func (r EffectResponseTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectResponseUnvalidatedPublishParams struct {
	Body []EffectResponseUnvalidatedPublishParamsBody
	paramObj
}

func (r EffectResponseUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *EffectResponseUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// A response for various effects on a target.
//
// The properties ClassificationMarking, DataMode, Source, Type are required.
type EffectResponseUnvalidatedPublishParamsBody struct {
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
	// The type of response in this record (e.g. COA, SCORECARD, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset.
	ActorSrcID param.Opt[string] `json:"actorSrcId,omitzero"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActorSrcType param.Opt[string] `json:"actorSrcType,omitzero"`
	// The collateral damage estimate (CDE) of the munition being fired.
	CollateralDamageEst param.Opt[float64] `json:"collateralDamageEst,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The deadline time to accept this COA before it's no longer valid, in ISO8601 UTC
	// format.
	DecisionDeadline param.Opt[time.Time] `json:"decisionDeadline,omitzero" format:"date-time"`
	// The external system identifier of the associated effect request. A human
	// readable unique id.
	ExternalRequestID param.Opt[string] `json:"externalRequestId,omitzero"`
	// Unique identifier of the EffectRequest associated with this response.
	IDEffectRequest param.Opt[string] `json:"idEffectRequest,omitzero"`
	// Unique identifier of the munition.
	MunitionID param.Opt[string] `json:"munitionId,omitzero"`
	// The type of munition being fired.
	MunitionType param.Opt[string] `json:"munitionType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The probability of kill (0-1) of the target being destroyed.
	ProbabilityOfKill param.Opt[float64] `json:"probabilityOfKill,omitzero"`
	// The record ID, depending on the type identified in redTargetSrcType, of the red
	// force target. If the redTargetSrcType is POI or TRACK, then this identifier
	// corresponds to either poi.poiid or track.trkId from their respective schemas.
	RedTargetSrcID param.Opt[string] `json:"redTargetSrcId,omitzero"`
	// The source type of the targetId identifier (POI, SITE, TRACK).
	RedTargetSrcType param.Opt[string] `json:"redTargetSrcType,omitzero"`
	// The time to overhead for the red force to be over their target, in ISO8601 UTC
	// format.
	RedTimeToOverhead param.Opt[time.Time] `json:"redTimeToOverhead,omitzero" format:"date-time"`
	// The number of shots required to destroy target.
	ShotsRequired param.Opt[int64] `json:"shotsRequired,omitzero"`
	// List of actions associated with this effect response.
	ActionsList []EffectResponseUnvalidatedPublishParamsBodyActionsList `json:"actionsList,omitzero"`
	// List of COA metrics associated with this effect response.
	CoaMetrics []EffectResponseUnvalidatedPublishParamsBodyCoaMetric `json:"coaMetrics,omitzero"`
	// List of external actions to be executed as part of this task.
	ExternalActions []string `json:"externalActions,omitzero"`
	paramObj
}

func (r EffectResponseUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EffectResponseUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type EffectResponseUnvalidatedPublishParamsBodyActionsList struct {
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset/actor.
	ActionActorSrcID param.Opt[string] `json:"actionActorSrcId,omitzero"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActionActorSrcType param.Opt[string] `json:"actionActorSrcType,omitzero"`
	// The desired end time of this task, in ISO8601 UTC format.
	ActionEndTime param.Opt[time.Time] `json:"actionEndTime,omitzero" format:"date-time"`
	// Identifier of this action.
	ActionID param.Opt[string] `json:"actionId,omitzero"`
	// The desired start time of this task, in ISO8601 UTC format.
	ActionStartTime param.Opt[time.Time] `json:"actionStartTime,omitzero" format:"date-time"`
	// The WGS-84 altitude of the asset/actor location at weapon launch, in meters.
	ActorInterceptAlt param.Opt[float64] `json:"actorInterceptAlt,omitzero"`
	// The WGS-84 latitude of the asset/actor location at weapon launch, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	ActorInterceptLat param.Opt[float64] `json:"actorInterceptLat,omitzero"`
	// The WGS-84 longitude of the asset/actor location at weapon launch, in degrees.
	// -180 to 180 degrees (negative values west of Prime Meridian).
	ActorInterceptLon param.Opt[float64] `json:"actorInterceptLon,omitzero"`
	// The type of munition or sensor used by this asset/actor.
	Effector param.Opt[string] `json:"effector,omitzero"`
	// A summary string describing different aspects of the action.
	Summary param.Opt[string] `json:"summary,omitzero"`
	// The POI or TRACK ID, depending on the type identified in targetSrcType, of the
	// requested target. This identifier corresponds to either poi.poiid or track.trkId
	// from their respective schemas.
	TargetSrcID param.Opt[string] `json:"targetSrcId,omitzero"`
	// The source type of the targetId identifier (POI, TRACK).
	TargetSrcType param.Opt[string] `json:"targetSrcType,omitzero"`
	// The end time of the asset TOT (time over target), in ISO8601 UTC format.
	TotEndTime param.Opt[time.Time] `json:"totEndTime,omitzero" format:"date-time"`
	// The start time of the asset TOT (time over target), in ISO8601 UTC format.
	TotStartTime param.Opt[time.Time] `json:"totStartTime,omitzero" format:"date-time"`
	// The WGS-84 altitude of the weapon destination location, in meters.
	WeaponInterceptAlt param.Opt[float64] `json:"weaponInterceptAlt,omitzero"`
	// The WGS-84 latitude of the weapon destination location, in degrees. -90 to 90
	// degrees (negative values south of equator).
	WeaponInterceptLat param.Opt[float64] `json:"weaponInterceptLat,omitzero"`
	// The WGS-84 longitude of the weapon destination location, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	WeaponInterceptLon param.Opt[float64] `json:"weaponInterceptLon,omitzero"`
	// List of metrics associated with this action.
	ActionMetrics []EffectResponseUnvalidatedPublishParamsBodyActionsListActionMetric `json:"actionMetrics,omitzero"`
	paramObj
}

func (r EffectResponseUnvalidatedPublishParamsBodyActionsList) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseUnvalidatedPublishParamsBodyActionsList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseUnvalidatedPublishParamsBodyActionsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseUnvalidatedPublishParamsBodyActionsListActionMetric struct {
	// The metric score specific to its domain.
	DomainValue param.Opt[float64] `json:"domainValue,omitzero"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType param.Opt[string] `json:"metricType,omitzero"`
	// The metric that was used to score this task.
	Provenance param.Opt[string] `json:"provenance,omitzero"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue param.Opt[float64] `json:"relativeValue,omitzero"`
	paramObj
}

func (r EffectResponseUnvalidatedPublishParamsBodyActionsListActionMetric) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseUnvalidatedPublishParamsBodyActionsListActionMetric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseUnvalidatedPublishParamsBodyActionsListActionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EffectResponseUnvalidatedPublishParamsBodyCoaMetric struct {
	// The metric score specific to its domain.
	DomainValue param.Opt[float64] `json:"domainValue,omitzero"`
	// The type of the metric (e.g. CollateralDamage, GoalAchievement, OpportunityCost,
	// Timeliness, Unavailable, etc.).
	MetricType param.Opt[string] `json:"metricType,omitzero"`
	// The metric that was used to score this task.
	Provenance param.Opt[string] `json:"provenance,omitzero"`
	// The metric score adjusted to be relative and comparable to other domains.
	RelativeValue param.Opt[float64] `json:"relativeValue,omitzero"`
	paramObj
}

func (r EffectResponseUnvalidatedPublishParamsBodyCoaMetric) MarshalJSON() (data []byte, err error) {
	type shadow EffectResponseUnvalidatedPublishParamsBodyCoaMetric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EffectResponseUnvalidatedPublishParamsBodyCoaMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
