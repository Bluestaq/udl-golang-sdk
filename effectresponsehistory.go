// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// EffectResponseHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEffectResponseHistoryService] method instead.
type EffectResponseHistoryService struct {
	Options []option.RequestOption
}

// NewEffectResponseHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEffectResponseHistoryService(opts ...option.RequestOption) (r EffectResponseHistoryService) {
	r = EffectResponseHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EffectResponseHistoryService) List(ctx context.Context, query EffectResponseHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EffectResponseHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/effectresponse/history"
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
func (r *EffectResponseHistoryService) ListAutoPaging(ctx context.Context, query EffectResponseHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EffectResponseHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EffectResponseHistoryService) Aodr(ctx context.Context, query EffectResponseHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/effectresponse/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EffectResponseHistoryService) Count(ctx context.Context, query EffectResponseHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/effectresponse/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A response for various effects on a target.
type EffectResponseHistoryListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode EffectResponseHistoryListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of response in this record (e.g. COA, SCORECARD, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// List of actions associated with this effect response.
	ActionsList []EffectResponseActionsListFull `json:"actionsList"`
	// The record ID, depending on the type identified in actorSrcType, of the
	// requested asset.
	ActorSrcID string `json:"actorSrcId"`
	// The source type of the asset/actor identifier (AIRCRAFT, LANDCRAFT, SEACRAFT,
	// TRACK).
	ActorSrcType string `json:"actorSrcType"`
	// List of COA metrics associated with this effect response.
	CoaMetrics []EffectResponseMetricsFull `json:"coaMetrics"`
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
func (r EffectResponseHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *EffectResponseHistoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type EffectResponseHistoryListResponseDataMode string

const (
	EffectResponseHistoryListResponseDataModeReal      EffectResponseHistoryListResponseDataMode = "REAL"
	EffectResponseHistoryListResponseDataModeTest      EffectResponseHistoryListResponseDataMode = "TEST"
	EffectResponseHistoryListResponseDataModeSimulated EffectResponseHistoryListResponseDataMode = "SIMULATED"
	EffectResponseHistoryListResponseDataModeExercise  EffectResponseHistoryListResponseDataMode = "EXERCISE"
)

type EffectResponseHistoryListParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt time.Time `query:"createdAt,required" format:"date" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EffectResponseHistoryListParams]'s query parameters as
// `url.Values`.
func (r EffectResponseHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectResponseHistoryAodrParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt time.Time `query:"createdAt,required" format:"date" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
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
	paramObj
}

// URLQuery serializes [EffectResponseHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r EffectResponseHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EffectResponseHistoryCountParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EffectResponseHistoryCountParams]'s query parameters as
// `url.Values`.
func (r EffectResponseHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
