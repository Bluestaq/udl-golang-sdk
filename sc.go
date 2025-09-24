// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/Bluestaq/udl-golang-sdk/internal/apiform"
	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// ScService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScService] method instead.
type ScService struct {
	Options       []option.RequestOption
	Notifications ScNotificationService
	File          ScFileService
	Folders       ScFolderService
	Paths         ScPathService
	View          ScViewService
	V2            ScV2Service
}

// NewScService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScService(opts ...option.RequestOption) (r ScService) {
	r = ScService{}
	r.Options = opts
	r.Notifications = NewScNotificationService(opts...)
	r.File = NewScFileService(opts...)
	r.Folders = NewScFolderService(opts...)
	r.Paths = NewScPathService(opts...)
	r.View = NewScViewService(opts...)
	r.V2 = NewScV2Service(opts...)
	return
}

// Deletes the requested file or folder in the passed path directory that is
// visible to the calling user. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
//
// Deprecated: deprecated
func (r *ScService) Delete(ctx context.Context, body ScDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Returns a list of the allowed filename extensions.
func (r *ScService) AllowableFileExtensions(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scs/allowableFileExtensions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of the allowed file upload mime types.
func (r *ScService) AllowableFileMimes(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scs/allowableFileMimes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// operation to copy folders or files. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
//
// Deprecated: deprecated
func (r *ScService) Copy(ctx context.Context, body ScCopyParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scs/copy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Downloads a zip of one or more files and/or folders.
func (r *ScService) Download(ctx context.Context, body ScDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "scs/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Download a single file from SCS.
func (r *ScService) FileDownload(ctx context.Context, query ScFileDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "scs/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Operation to upload a file. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
//
// Deprecated: deprecated
func (r *ScService) FileUpload(ctx context.Context, fileContent io.Reader, body ScFileUploadParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", fileContent)}, opts...)
	path := "scs/file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// operation to move folders or files. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
//
// Deprecated: deprecated
func (r *ScService) Move(ctx context.Context, body ScMoveParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scs/move"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Operation to rename folders or files. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
//
// Deprecated: deprecated
func (r *ScService) Rename(ctx context.Context, body ScRenameParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/rename"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Search for files by metadata and/or text in file content.
//
// Deprecated: deprecated
func (r *ScService) Search(ctx context.Context, params ScSearchParams, opts ...option.RequestOption) (res *[]shared.FileData, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scs/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SearchCriterionUnionParam struct {
	OfSearchCriterionScsSearchFieldCriterion *SearchCriterionScsSearchFieldCriterionParam `json:",omitzero,inline"`
	OfSearchLogicalCriterion                 *SearchLogicalCriterionParam                 `json:",omitzero,inline"`
	paramUnion
}

func (u SearchCriterionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSearchCriterionScsSearchFieldCriterion, u.OfSearchLogicalCriterion)
}
func (u *SearchCriterionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SearchCriterionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSearchCriterionScsSearchFieldCriterion) {
		return u.OfSearchCriterionScsSearchFieldCriterion
	} else if !param.IsOmitted(u.OfSearchLogicalCriterion) {
		return u.OfSearchLogicalCriterion
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SearchCriterionUnionParam) GetField() *string {
	if vt := u.OfSearchCriterionScsSearchFieldCriterion; vt != nil && vt.Field.Valid() {
		return &vt.Field.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SearchCriterionUnionParam) GetValue() *string {
	if vt := u.OfSearchCriterionScsSearchFieldCriterion; vt != nil && vt.Value.Valid() {
		return &vt.Value.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SearchCriterionUnionParam) GetCriteria() []SearchCriterionUnionParam {
	if vt := u.OfSearchLogicalCriterion; vt != nil {
		return vt.Criteria
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SearchCriterionUnionParam) GetOperator() *string {
	if vt := u.OfSearchCriterionScsSearchFieldCriterion; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSearchLogicalCriterion; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// A search on a specific field with a given value and operator.
type SearchCriterionScsSearchFieldCriterionParam struct {
	// The field to search on (e.g., attachment.content, createdBy).
	Field param.Opt[string] `json:"field,omitzero"`
	// The value to compare against (e.g., The Great Gatsby)
	Value param.Opt[string] `json:"value,omitzero"`
	// Supported search operators
	//
	// Any of "EXACT_MATCH", "WILDCARD", "FUZZY", "GTE", "LTE", "GT", "LT".
	Operator string `json:"operator,omitzero"`
	paramObj
}

func (r SearchCriterionScsSearchFieldCriterionParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchCriterionScsSearchFieldCriterionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchCriterionScsSearchFieldCriterionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SearchCriterionScsSearchFieldCriterionParam](
		"operator", "EXACT_MATCH", "WILDCARD", "FUZZY", "GTE", "LTE", "GT", "LT",
	)
}

// Combines multiple search criteria with a logical operator (AND, OR, NOT).
type SearchLogicalCriterionParam struct {
	// List of search criteria to combine
	Criteria []SearchCriterionUnionParam `json:"criteria,omitzero"`
	// Supported search operators
	//
	// Any of "AND", "OR", "NOT".
	Operator SearchLogicalCriterionOperator `json:"operator,omitzero"`
	paramObj
}

func (r SearchLogicalCriterionParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchLogicalCriterionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchLogicalCriterionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Supported search operators
type SearchLogicalCriterionOperator string

const (
	SearchLogicalCriterionOperatorAnd SearchLogicalCriterionOperator = "AND"
	SearchLogicalCriterionOperatorOr  SearchLogicalCriterionOperator = "OR"
	SearchLogicalCriterionOperatorNot SearchLogicalCriterionOperator = "NOT"
)

type ScDeleteParams struct {
	// The id of the item to delete
	ID string `query:"id,required" json:"-"`
	paramObj
}

// URLQuery serializes [ScDeleteParams]'s query parameters as `url.Values`.
func (r ScDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScCopyParams struct {
	// The path of the item to copy
	ID string `query:"id,required" json:"-"`
	// The path to copy to
	TargetPath string `query:"targetPath,required" json:"-"`
	paramObj
}

// URLQuery serializes [ScCopyParams]'s query parameters as `url.Values`.
func (r ScCopyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScDownloadParams struct {
	Body []string
	paramObj
}

func (r ScDownloadParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ScDownloadParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type ScFileDownloadParams struct {
	// The complete path and filename of the file to download.
	ID          string           `query:"id,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ScFileDownloadParams]'s query parameters as `url.Values`.
func (r ScFileDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScFileUploadParams struct {
	// Classification marking of the file being uploaded.
	ClassificationMarking string `query:"classificationMarking,required" json:"-"`
	// Name of the file to upload.
	FileName string `query:"fileName,required" json:"-"`
	// The base path to upload file
	Path string `query:"path,required" json:"-"`
	// Length of time after which to automatically delete the file.
	DeleteAfter param.Opt[string] `query:"deleteAfter,omitzero" json:"-"`
	// Description
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	// Whether or not to overwrite a file with the same name and path, if one exists.
	Overwrite param.Opt[bool] `query:"overwrite,omitzero" json:"-"`
	// Whether or not to send a notification that this file was uploaded.
	SendNotification param.Opt[bool] `query:"sendNotification,omitzero" json:"-"`
	// Tags
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	paramObj
}

func (r ScFileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [ScFileUploadParams]'s query parameters as `url.Values`.
func (r ScFileUploadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScMoveParams struct {
	// The path of the item to move
	ID string `query:"id,required" json:"-"`
	// The path to copy to
	TargetPath string `query:"targetPath,required" json:"-"`
	paramObj
}

// URLQuery serializes [ScMoveParams]'s query parameters as `url.Values`.
func (r ScMoveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScRenameParams struct {
	// The path of the item to rename.
	ID string `query:"id,required" json:"-"`
	// The new name for the file or folder. Do not include the path.
	NewName string `query:"newName,required" json:"-"`
	paramObj
}

// URLQuery serializes [ScRenameParams]'s query parameters as `url.Values`.
func (r ScRenameParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScSearchParams struct {
	// The path to search from
	Path string `query:"path,required" json:"-"`
	// Number of items per page
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// First result to return
	Offset           param.Opt[int64]    `query:"offset,omitzero" json:"-"`
	ContentCriteria  param.Opt[string]   `json:"contentCriteria,omitzero"`
	SearchAfter      param.Opt[string]   `json:"searchAfter,omitzero"`
	MetaDataCriteria map[string][]string `json:"metaDataCriteria,omitzero"`
	NonRangeCriteria map[string][]string `json:"nonRangeCriteria,omitzero"`
	RangeCriteria    map[string][]string `json:"rangeCriteria,omitzero"`
	paramObj
}

func (r ScSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow ScSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ScSearchParams]'s query parameters as `url.Values`.
func (r ScSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
