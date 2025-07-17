// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
)

// ElsetCurrentService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewElsetCurrentService] method instead.
type ElsetCurrentService struct {
	Options []option.RequestOption
}

// NewElsetCurrentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewElsetCurrentService(opts ...option.RequestOption) (r ElsetCurrentService) {
	r = ElsetCurrentService{}
	r.Options = opts
	return
}

// Service operation to dynamically query/filter current elsets within the system
// by a variety of query parameters not specified in this API documentation. A
// current elset is the currently active, latest elset for an on-orbit object.
// Current elsets are tracked by source and a source should be provided as a query
// parameter to this service operation to view the 'current' catalog for a
// particular provider. If source is not provided, it will be defaulted to '18th
// SPCS'. See the queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more
// details on additional query parameter information.
func (r *ElsetCurrentService) List(ctx context.Context, query ElsetCurrentListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ElsetAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/elset/current"
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

// Service operation to dynamically query/filter current elsets within the system
// by a variety of query parameters not specified in this API documentation. A
// current elset is the currently active, latest elset for an on-orbit object.
// Current elsets are tracked by source and a source should be provided as a query
// parameter to this service operation to view the 'current' catalog for a
// particular provider. If source is not provided, it will be defaulted to '18th
// SPCS'. See the queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more
// details on additional query parameter information.
func (r *ElsetCurrentService) ListAutoPaging(ctx context.Context, query ElsetCurrentListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ElsetAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query/filter current elsets within the system
// by a variety of query parameters not specified in this API documentation. A
// current elset is the currently active, latest elset for an on-orbit object.
// Current elsets are tracked by source and a source should be provided as a query
// parameter to this service operation to view the 'current' catalog for a
// particular provider. If source is not provided, it will be defaulted to '18th
// SPCS'. See the queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more
// details on additional query parameter information.
func (r *ElsetCurrentService) Tuple(ctx context.Context, query ElsetCurrentTupleParams, opts ...option.RequestOption) (res *[]Elset, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/elset/current/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ElsetCurrentListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ElsetCurrentListParams]'s query parameters as `url.Values`.
func (r ElsetCurrentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElsetCurrentTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ElsetCurrentTupleParams]'s query parameters as
// `url.Values`.
func (r ElsetCurrentTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
