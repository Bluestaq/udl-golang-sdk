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

// SupportingDataDataTypeService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportingDataDataTypeService] method instead.
type SupportingDataDataTypeService struct {
	Options []option.RequestOption
}

// NewSupportingDataDataTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSupportingDataDataTypeService(opts ...option.RequestOption) (r SupportingDataDataTypeService) {
	r = SupportingDataDataTypeService{}
	r.Options = opts
	return
}

// Retrieves all distinct data owner types.
func (r *SupportingDataDataTypeService) List(ctx context.Context, query SupportingDataDataTypeListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[string], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/dataowner/getDataTypes"
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

// Retrieves all distinct data owner types.
func (r *SupportingDataDataTypeService) ListAutoPaging(ctx context.Context, query SupportingDataDataTypeListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[string] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

type SupportingDataDataTypeListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SupportingDataDataTypeListParams]'s query parameters as
// `url.Values`.
func (r SupportingDataDataTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
