// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
)

// PersonnelrecoveryHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonnelrecoveryHistoryService] method instead.
type PersonnelrecoveryHistoryService struct {
	Options []option.RequestOption
}

// NewPersonnelrecoveryHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPersonnelrecoveryHistoryService(opts ...option.RequestOption) (r PersonnelrecoveryHistoryService) {
	r = PersonnelrecoveryHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *PersonnelrecoveryHistoryService) List(ctx context.Context, query PersonnelrecoveryHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[PersonnelRecoveryFullL], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/personnelrecovery/history"
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
func (r *PersonnelrecoveryHistoryService) ListAutoPaging(ctx context.Context, query PersonnelrecoveryHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[PersonnelRecoveryFullL] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *PersonnelrecoveryHistoryService) Count(ctx context.Context, query PersonnelrecoveryHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/personnelrecovery/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PersonnelrecoveryHistoryListParams struct {
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime time.Time `query:"msgTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PersonnelrecoveryHistoryListParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [PersonnelrecoveryHistoryListParams]'s query parameters as
// `url.Values`.
func (r PersonnelrecoveryHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PersonnelrecoveryHistoryCountParams struct {
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PersonnelrecoveryHistoryCountParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [PersonnelrecoveryHistoryCountParams]'s query parameters as
// `url.Values`.
func (r PersonnelrecoveryHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
