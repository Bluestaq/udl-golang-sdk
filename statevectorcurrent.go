// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
)

// StateVectorCurrentService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStateVectorCurrentService] method instead.
type StateVectorCurrentService struct {
	Options []option.RequestOption
}

// NewStateVectorCurrentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStateVectorCurrentService(opts ...option.RequestOption) (r StateVectorCurrentService) {
	r = StateVectorCurrentService{}
	r.Options = opts
	return
}

// Service operation to dynamically query/filter current StateVector data within
// the system by a variety of query parameters not specified in this API
// documentation. A current StateVector is the currently active, latest StateVector
// for an on-orbit object. Current state vectors are tracked by source and a source
// should be provided as a query parameter to this service operation to view the
// 'current' catalog for a particular provider. Default current state vector
// sources may vary by UDL environment. Please contact the UDL help desk for more
// information, or explicitly specify the desired source. See the queryhelp
// operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required
// query parameter information.
func (r *StateVectorCurrentService) List(ctx context.Context, query StateVectorCurrentListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[StateVectorAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/statevector/current"
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

// Service operation to dynamically query/filter current StateVector data within
// the system by a variety of query parameters not specified in this API
// documentation. A current StateVector is the currently active, latest StateVector
// for an on-orbit object. Current state vectors are tracked by source and a source
// should be provided as a query parameter to this service operation to view the
// 'current' catalog for a particular provider. Default current state vector
// sources may vary by UDL environment. Please contact the UDL help desk for more
// information, or explicitly specify the desired source. See the queryhelp
// operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required
// query parameter information.
func (r *StateVectorCurrentService) ListAutoPaging(ctx context.Context, query StateVectorCurrentListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[StateVectorAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query/filter current StateVector data within
// the system by a variety of query parameters not specified in this API
// documentation. A current StateVector is the currently active, latest StateVector
// for an on-orbit object. Current state vectors are tracked by source and a source
// should be provided as a query parameter to this service operation to view the
// 'current' catalog for a particular provider. Default current state vector
// sources may vary by UDL environment. Please contact the UDL help desk for more
// information, or explicitly specify the desired source. See the queryhelp
// operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required
// query parameter information.
func (r *StateVectorCurrentService) Tuple(ctx context.Context, query StateVectorCurrentTupleParams, opts ...option.RequestOption) (res *[]StateVectorFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/statevector/current/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StateVectorCurrentListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StateVectorCurrentListParams]'s query parameters as
// `url.Values`.
func (r StateVectorCurrentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StateVectorCurrentTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StateVectorCurrentTupleParams]'s query parameters as
// `url.Values`.
func (r StateVectorCurrentTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
