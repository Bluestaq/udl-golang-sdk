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

// OnorbitthrusterstatusHistoryService contains methods and other services that
// help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitthrusterstatusHistoryService] method instead.
type OnorbitthrusterstatusHistoryService struct {
	Options []option.RequestOption
}

// NewOnorbitthrusterstatusHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOnorbitthrusterstatusHistoryService(opts ...option.RequestOption) (r OnorbitthrusterstatusHistoryService) {
	r = OnorbitthrusterstatusHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitthrusterstatusHistoryService) List(ctx context.Context, query OnorbitthrusterstatusHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[shared.OnorbitthrusterstatusFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitthrusterstatus/history"
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
func (r *OnorbitthrusterstatusHistoryService) ListAutoPaging(ctx context.Context, query OnorbitthrusterstatusHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[shared.OnorbitthrusterstatusFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnorbitthrusterstatusHistoryService) Count(ctx context.Context, query OnorbitthrusterstatusHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onorbitthrusterstatus/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OnorbitthrusterstatusHistoryListParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) ID of the
	// associated OnorbitThruster record. This ID can be used to obtain additional
	// information on an onorbit thruster object using the 'get by ID' operation (e.g.
	// /udl/onorbitthruster/{id}). For example, the OnorbitThruster object with
	// idOnorbitThruster = abc would be queried as /udl/onorbitthruster/abc.
	IDOnorbitThruster param.Opt[string] `query:"idOnorbitThruster,omitzero" json:"-"`
	MaxResults        param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) Datetime
	// of the thruster status observation in ISO 8601 UTC datetime format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StatusTime param.Opt[time.Time] `query:"statusTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitthrusterstatusHistoryListParams]'s query parameters
// as `url.Values`.
func (r OnorbitthrusterstatusHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitthrusterstatusHistoryCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) ID of the
	// associated OnorbitThruster record. This ID can be used to obtain additional
	// information on an onorbit thruster object using the 'get by ID' operation (e.g.
	// /udl/onorbitthruster/{id}). For example, the OnorbitThruster object with
	// idOnorbitThruster = abc would be queried as /udl/onorbitthruster/abc.
	IDOnorbitThruster param.Opt[string] `query:"idOnorbitThruster,omitzero" json:"-"`
	MaxResults        param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'idOnorbitThruster, statusTime' are required.) Datetime
	// of the thruster status observation in ISO 8601 UTC datetime format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StatusTime param.Opt[time.Time] `query:"statusTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitthrusterstatusHistoryCountParams]'s query parameters
// as `url.Values`.
func (r OnorbitthrusterstatusHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
