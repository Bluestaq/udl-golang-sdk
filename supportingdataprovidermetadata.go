// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
)

// SupportingDataProviderMetadataService contains methods and other services that
// help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportingDataProviderMetadataService] method instead.
type SupportingDataProviderMetadataService struct {
	Options []option.RequestOption
}

// NewSupportingDataProviderMetadataService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSupportingDataProviderMetadataService(opts ...option.RequestOption) (r SupportingDataProviderMetadataService) {
	r = SupportingDataProviderMetadataService{}
	r.Options = opts
	return
}

func (r *SupportingDataProviderMetadataService) Get(ctx context.Context, query SupportingDataProviderMetadataGetParams, opts ...option.RequestOption) (res *[]DataownerAbridged, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/dataowner/providerMetadata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SupportingDataProviderMetadataGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SupportingDataProviderMetadataGetParams]'s query parameters
// as `url.Values`.
func (r SupportingDataProviderMetadataGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
