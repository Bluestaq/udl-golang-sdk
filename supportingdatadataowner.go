// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// SupportingDataDataownerService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportingDataDataownerService] method instead.
type SupportingDataDataownerService struct {
	Options []option.RequestOption
}

// NewSupportingDataDataownerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSupportingDataDataownerService(opts ...option.RequestOption) (r SupportingDataDataownerService) {
	r = SupportingDataDataownerService{}
	r.Options = opts
	return
}

func (r *SupportingDataDataownerService) Get(ctx context.Context, query SupportingDataDataownerGetParams, opts ...option.RequestOption) (res *[]DataownerAbridged, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/dataowner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SupportingDataDataownerService) Count(ctx context.Context, query SupportingDataDataownerCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/dataowner/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SupportingDataDataownerService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *SupportingDataDataownerQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/dataowner/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all distinct data owner types.
func (r *SupportingDataDataownerService) GetDataOwnerTypes(ctx context.Context, query SupportingDataDataownerGetDataOwnerTypesParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/dataowner/getDataOwnerTypes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *SupportingDataDataownerService) GetProviderMetadata(ctx context.Context, query SupportingDataDataownerGetProviderMetadataParams, opts ...option.RequestOption) (res *[]DataownerAbridged, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/dataowner/providerMetadata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information pertaining to UDL data owners.
type DataownerAbridged struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Description of this data owner.
	Description string `json:"description,required"`
	// The name of the data owner.
	DoName string `json:"doName,required"`
	// Unique identifier of the contact for this data owner.
	IDContact string `json:"idContact,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Boolean indicating if the data owner is coming soon or not yet available.
	ComingSoon bool `json:"comingSoon"`
	// Optional control required to access this data type from this owner.
	Control string `json:"control"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Type of data this data owner owns (e.g. EPHEMERIS, IMAGERY, MANEUVER, etc.).
	DataType string `json:"dataType"`
	// Boolean indicating if the data owner is enabled (if not enabled, they should not
	// appear on the data products screen on the storefront).
	Enabled bool `json:"enabled"`
	// Type of organization which this data owner belongs to (e.g. Commercial,
	// Government, Academic, Consortium, etc.).
	OwnerType string `json:"ownerType"`
	// Organization name for the data provider.
	Provider string `json:"provider"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Description           respjson.Field
		DoName                respjson.Field
		IDContact             respjson.Field
		Source                respjson.Field
		ComingSoon            respjson.Field
		Control               respjson.Field
		CountryCode           respjson.Field
		DataType              respjson.Field
		Enabled               respjson.Field
		OwnerType             respjson.Field
		Provider              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataownerAbridged) RawJSON() string { return r.JSON.raw }
func (r *DataownerAbridged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SupportingDataDataownerQueryHelpResponse struct {
	AodrSupported         bool                         `json:"aodrSupported"`
	ClassificationMarking string                       `json:"classificationMarking"`
	Description           string                       `json:"description"`
	HistorySupported      bool                         `json:"historySupported"`
	Name                  string                       `json:"name"`
	Parameters            []shared.ParamDescriptorResp `json:"parameters"`
	RequiredRoles         []string                     `json:"requiredRoles"`
	RestSupported         bool                         `json:"restSupported"`
	SortSupported         bool                         `json:"sortSupported"`
	TypeName              string                       `json:"typeName"`
	Uri                   string                       `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AodrSupported         respjson.Field
		ClassificationMarking respjson.Field
		Description           respjson.Field
		HistorySupported      respjson.Field
		Name                  respjson.Field
		Parameters            respjson.Field
		RequiredRoles         respjson.Field
		RestSupported         respjson.Field
		SortSupported         respjson.Field
		TypeName              respjson.Field
		Uri                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportingDataDataownerQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportingDataDataownerQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SupportingDataDataownerGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SupportingDataDataownerGetParams]'s query parameters as
// `url.Values`.
func (r SupportingDataDataownerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SupportingDataDataownerCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SupportingDataDataownerCountParams]'s query parameters as
// `url.Values`.
func (r SupportingDataDataownerCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SupportingDataDataownerGetDataOwnerTypesParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SupportingDataDataownerGetDataOwnerTypesParams]'s query
// parameters as `url.Values`.
func (r SupportingDataDataownerGetDataOwnerTypesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SupportingDataDataownerGetProviderMetadataParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SupportingDataDataownerGetProviderMetadataParams]'s query
// parameters as `url.Values`.
func (r SupportingDataDataownerGetProviderMetadataParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
