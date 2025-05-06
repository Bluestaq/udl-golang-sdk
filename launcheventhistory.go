// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// LaunchEventHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaunchEventHistoryService] method instead.
type LaunchEventHistoryService struct {
	Options []option.RequestOption
}

// NewLaunchEventHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLaunchEventHistoryService(opts ...option.RequestOption) (r LaunchEventHistoryService) {
	r = LaunchEventHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaunchEventHistoryService) List(ctx context.Context, query LaunchEventHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaunchEventHistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/launchevent/history"
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
func (r *LaunchEventHistoryService) ListAutoPaging(ctx context.Context, query LaunchEventHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaunchEventHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaunchEventHistoryService) Aodr(ctx context.Context, query LaunchEventHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchevent/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LaunchEventHistoryService) Count(ctx context.Context, query LaunchEventHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/launchevent/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information on known launch events.
type LaunchEventHistoryListResponse struct {
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
	DataMode LaunchEventHistoryListResponseDataMode `json:"dataMode,required"`
	// Timestamp of the originating message in ISO8601 UTC format.
	MsgCreateDate time.Time `json:"msgCreateDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The Basic Encyclopedia Number, if applicable.
	BeNumber string `json:"beNumber"`
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
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// The launch date, in ISO8601 UTC format.
	LaunchDate time.Time `json:"launchDate" format:"date-time"`
	// The Launch facility name.
	LaunchFacilityName string `json:"launchFacilityName"`
	// The DISOB launch Failure Code, if applicable.
	LaunchFailureCode string `json:"launchFailureCode"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional target-id, if missing in UDL.
	OrigObjectID string `json:"origObjectId"`
	// The OSuffix, if applicable.
	OSuffix string `json:"oSuffix"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking  resp.Field
		DataMode               resp.Field
		MsgCreateDate          resp.Field
		Source                 resp.Field
		ID                     resp.Field
		BeNumber               resp.Field
		CreatedAt              resp.Field
		CreatedBy              resp.Field
		DeclassificationDate   resp.Field
		DeclassificationString resp.Field
		DerivedFrom            resp.Field
		IDOnOrbit              resp.Field
		LaunchDate             resp.Field
		LaunchFacilityName     resp.Field
		LaunchFailureCode      resp.Field
		OnOrbit                resp.Field
		Origin                 resp.Field
		OrigNetwork            resp.Field
		OrigObjectID           resp.Field
		OSuffix                resp.Field
		SatNo                  resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchEventHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchEventHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type LaunchEventHistoryListResponseDataMode string

const (
	LaunchEventHistoryListResponseDataModeReal      LaunchEventHistoryListResponseDataMode = "REAL"
	LaunchEventHistoryListResponseDataModeTest      LaunchEventHistoryListResponseDataMode = "TEST"
	LaunchEventHistoryListResponseDataModeSimulated LaunchEventHistoryListResponseDataMode = "SIMULATED"
	LaunchEventHistoryListResponseDataModeExercise  LaunchEventHistoryListResponseDataMode = "EXERCISE"
)

type LaunchEventHistoryListParams struct {
	// Timestamp of the originating message in ISO8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgCreateDate time.Time `query:"msgCreateDate,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchEventHistoryListParams]'s query parameters as
// `url.Values`.
func (r LaunchEventHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchEventHistoryAodrParams struct {
	// Timestamp of the originating message in ISO8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgCreateDate time.Time `query:"msgCreateDate,required" format:"date-time" json:"-"`
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

// URLQuery serializes [LaunchEventHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r LaunchEventHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchEventHistoryCountParams struct {
	// Timestamp of the originating message in ISO8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgCreateDate time.Time        `query:"msgCreateDate,required" format:"date-time" json:"-"`
	FirstResult   param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults    param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchEventHistoryCountParams]'s query parameters as
// `url.Values`.
func (r LaunchEventHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
