// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// EventEvolutionHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventEvolutionHistoryService] method instead.
type EventEvolutionHistoryService struct {
	Options []option.RequestOption
}

// NewEventEvolutionHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEventEvolutionHistoryService(opts ...option.RequestOption) (r EventEvolutionHistoryService) {
	r = EventEvolutionHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EventEvolutionHistoryService) List(ctx context.Context, query EventEvolutionHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[shared.EventEvolutionFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/eventevolution/history"
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
func (r *EventEvolutionHistoryService) ListAutoPaging(ctx context.Context, query EventEvolutionHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[shared.EventEvolutionFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EventEvolutionHistoryService) Aodr(ctx context.Context, query EventEvolutionHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/eventevolution/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EventEvolutionHistoryService) Count(ctx context.Context, query EventEvolutionHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/eventevolution/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type EventEvolutionHistoryListParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns param.Opt[string] `query:"columns,omitzero" json:"-"`
	// (One or more of fields 'eventId, startTime' are required.) User-provided unique
	// identifier of this activity or event. This ID should remain the same on
	// subsequent updates in order to associate all records pertaining to the activity
	// or event.
	EventID     param.Opt[string] `query:"eventId,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'eventId, startTime' are required.) The actual or
	// estimated start time of the activity or event, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EventEvolutionHistoryListParams]'s query parameters as
// `url.Values`.
func (r EventEvolutionHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventEvolutionHistoryAodrParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns param.Opt[string] `query:"columns,omitzero" json:"-"`
	// (One or more of fields 'eventId, startTime' are required.) User-provided unique
	// identifier of this activity or event. This ID should remain the same on
	// subsequent updates in order to associate all records pertaining to the activity
	// or event.
	EventID     param.Opt[string] `query:"eventId,omitzero" json:"-"`
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
	// (One or more of fields 'eventId, startTime' are required.) The actual or
	// estimated start time of the activity or event, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EventEvolutionHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r EventEvolutionHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventEvolutionHistoryCountParams struct {
	// (One or more of fields 'eventId, startTime' are required.) User-provided unique
	// identifier of this activity or event. This ID should remain the same on
	// subsequent updates in order to associate all records pertaining to the activity
	// or event.
	EventID     param.Opt[string] `query:"eventId,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'eventId, startTime' are required.) The actual or
	// estimated start time of the activity or event, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EventEvolutionHistoryCountParams]'s query parameters as
// `url.Values`.
func (r EventEvolutionHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
