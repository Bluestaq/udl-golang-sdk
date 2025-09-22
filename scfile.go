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
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// ScFileService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScFileService] method instead.
type ScFileService struct {
	Options []option.RequestOption
}

// NewScFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScFileService(opts ...option.RequestOption) (r ScFileService) {
	r = ScFileService{}
	r.Options = opts
	return
}

// Returns a FileData object representing the file with the given ID that is
// visible to the calling user.
//
// Deprecated: deprecated
func (r *ScFileService) Get(ctx context.Context, query ScFileGetParams, opts ...option.RequestOption) (res *shared.FileData, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scs/file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// operation to update files metadata. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
//
// Deprecated: deprecated
func (r *ScFileService) Update(ctx context.Context, body ScFileUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Returns a non-recursive list of FileData objects representing the files and
// subdirectories in the passed-in path directory that are visible to the calling
// user.
//
// Deprecated: deprecated
func (r *ScFileService) List(ctx context.Context, query ScFileListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[shared.FileData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "scs/list"
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

// Returns a non-recursive list of FileData objects representing the files and
// subdirectories in the passed-in path directory that are visible to the calling
// user.
//
// Deprecated: deprecated
func (r *ScFileService) ListAutoPaging(ctx context.Context, query ScFileListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[shared.FileData] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

type ScFileGetParams struct {
	// The file ID to view
	ID          string           `query:"id,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ScFileGetParams]'s query parameters as `url.Values`.
func (r ScFileGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScFileUpdateParams struct {
	FileDataList []shared.FileDataParam `json:"fileDataList,omitzero"`
	paramObj
}

func (r ScFileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ScFileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScFileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScFileListParams struct {
	// The base path to list
	Path string `query:"path,required" json:"-"`
	// Number of items per page
	Count       param.Opt[int64] `query:"count,omitzero" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	// First result to return
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ScFileListParams]'s query parameters as `url.Values`.
func (r ScFileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
