// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// These services provide operations for manipulating and querying Aircraft Sortie,
// Aircraft Mission, Item Tracking, Flight Plan, Air Event, Sortie Prior Permission
// Required (PPR), Diplomatic Clearance, Diplomatic Clearance Country, Airspace
// Control Order, Air Tasking Order, Navigational Obstruction, Logistics Support,
// Track Route, Air Load Plan, and Aviation Risk Management data. Aircraft Sortie
// information contains static and dynamic aircraft assignments, departure and
// arrival times, and remarks. Aircraft Mission information contains static data
// for mission planning to include assigned aircraft and crews, cargo pickup and
// dropoff locations, unique identifiers, and prioritization. Item Tracking
// information contains data for tracking an item from its origin to destination
// and how it may be configured during transport. Flight Plan information contains
// schedule and route details. Air Event provides information concerning various
// aerial events such as fuel transfer and air drops, as well as the associated
// aircraft involved. Sortie PPR information contains details on operational access
// to a runway, taxiway, or airport service. Diplomatic Clearance information
// contains details on the issuance and coordination of aircraft clearance
// requests. Diplomatic Clearance Country provides information such as entry/exit
// points, requirements, and points of contact for countries diplomatic clearances
// are being created for. Airspace Control Order provides information concerning
// the allocation, restriction, and deconfliction of airspace. Air Tasking Order
// information contains details on the coordination of air missions and their
// tasks, resources, and timelines. Navigational Obstruction provides the
// locations, characteristics, and boundaries of obstacles and structures that can
// restrict or interfere with navigation. Logistics Support contains information
// regarding the transport and maintenance of resources and equipment to sustain
// air operations. Track Route information defines specific flight paths used by
// aircraft during the transport of fuel and other resources. Air Load Plan
// information provides mission actuals concerning the loading and air transport of
// cargo and passengers. Aviation Risk Management information help aid in mission
// planning by accounting for factors such as mission complexity and crew fatigue.
//
// AirTransportMissionHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirTransportMissionHistoryService] method instead.
type AirTransportMissionHistoryService struct {
	Options []option.RequestOption
}

// NewAirTransportMissionHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAirTransportMissionHistoryService(opts ...option.RequestOption) (r AirTransportMissionHistoryService) {
	r = AirTransportMissionHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirTransportMissionHistoryService) List(ctx context.Context, query AirTransportMissionHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[shared.AirTransportMissionFull], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/airtransportmission/history"
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
func (r *AirTransportMissionHistoryService) ListAutoPaging(ctx context.Context, query AirTransportMissionHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[shared.AirTransportMissionFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirTransportMissionHistoryService) Aodr(ctx context.Context, query AirTransportMissionHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/airtransportmission/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirTransportMissionHistoryService) Count(ctx context.Context, query AirTransportMissionHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airtransportmission/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AirTransportMissionHistoryListParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt time.Time `query:"createdAt" api:"required" format:"date" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirTransportMissionHistoryListParams]'s query parameters as
// `url.Values`.
func (r AirTransportMissionHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirTransportMissionHistoryAodrParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt time.Time `query:"createdAt" api:"required" format:"date" json:"-"`
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

// URLQuery serializes [AirTransportMissionHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r AirTransportMissionHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirTransportMissionHistoryCountParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt" api:"required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirTransportMissionHistoryCountParams]'s query parameters
// as `url.Values`.
func (r AirTransportMissionHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
