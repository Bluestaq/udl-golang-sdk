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

// SupportingDataDataownerTypeService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportingDataDataownerTypeService] method instead.
type SupportingDataDataownerTypeService struct {
	Options []option.RequestOption
}

// NewSupportingDataDataownerTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSupportingDataDataownerTypeService(opts ...option.RequestOption) (r SupportingDataDataownerTypeService) {
	r = SupportingDataDataownerTypeService{}
	r.Options = opts
	return
}

// Retrieves all distinct data owner types.
func (r *SupportingDataDataownerTypeService) List(ctx context.Context, query SupportingDataDataownerTypeListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[string], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/dataowner/getDataOwnerTypes"
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
func (r *SupportingDataDataownerTypeService) ListAutoPaging(ctx context.Context, query SupportingDataDataownerTypeListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[string] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

type SupportingDataDataownerTypeListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SupportingDataDataownerTypeListParams]'s query parameters
// as `url.Values`.
func (r SupportingDataDataownerTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
