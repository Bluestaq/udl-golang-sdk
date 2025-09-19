// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/Bluestaq/udl-golang-sdk/internal/apiform"
	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// ScV2Service contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScV2Service] method instead.
type ScV2Service struct {
	Options []option.RequestOption
}

// NewScV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScV2Service(opts ...option.RequestOption) (r ScV2Service) {
	r = ScV2Service{}
	r.Options = opts
	return
}

// Update folders and files. For a folder, you may update description, writeAcl,
// readAcl, classificationMarking, and tags. For a file, you may update
// description, classificationMarking, and tags. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *ScV2Service) Update(ctx context.Context, params ScV2UpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/v2/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

// Returns a list of ScsEntity objects, each directly nested under the provided
// path.
func (r *ScV2Service) List(ctx context.Context, query ScV2ListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ScsEntity], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "scs/v2/list"
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

// Returns a list of ScsEntity objects, each directly nested under the provided
// path.
func (r *ScV2Service) ListAutoPaging(ctx context.Context, query ScV2ListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ScsEntity] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Operation to delete a folder or file. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *ScV2Service) Delete(ctx context.Context, body ScV2DeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/v2/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Operation to copy a folder or file. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *ScV2Service) Copy(ctx context.Context, body ScV2CopyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/v2/copy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Operation to upload a file. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *ScV2Service) FileUpload(ctx context.Context, fileContent io.Reader, body ScV2FileUploadParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", ""), option.WithRequestBody("application/octet-stream", fileContent)}, opts...)
	path := "scs/v2/file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Creates all folders in provided path that don't exist. Can be used to create a
// single folder or a new folder structure. Provided classificationMarking,
// description, writeAcl, readAcl, and tags are applied to the deepest folder in
// the provided path. If parent folders are created by this request, each parent
// folder will be created with the same classificationMarking and tags. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *ScV2Service) FolderNew(ctx context.Context, params ScV2FolderNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/v2/folder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Operation to move or rename a folder or file. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *ScV2Service) Move(ctx context.Context, body ScV2MoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/v2/move"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Operation to search for files in the Secure Content Store.
func (r *ScV2Service) Search(ctx context.Context, params ScV2SearchParams, opts ...option.RequestOption) (res *[]ScsEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scs/v2/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type Attachment struct {
	Author        string `json:"author"`
	Content       string `json:"content"`
	ContentLength int64  `json:"content_length"`
	ContentType   string `json:"content_type"`
	Date          string `json:"date"`
	Keywords      string `json:"keywords"`
	Language      string `json:"language"`
	Title         string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Author        respjson.Field
		Content       respjson.Field
		ContentLength respjson.Field
		ContentType   respjson.Field
		Date          respjson.Field
		Keywords      respjson.Field
		Language      respjson.Field
		Title         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Attachment) RawJSON() string { return r.JSON.raw }
func (r *Attachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An SCS file or folder.
type ScsEntity struct {
	ID         string     `json:"id"`
	Attachment Attachment `json:"attachment"`
	// Classification marking of the folder or file in IC/CAPCO portion-marked format.
	ClassificationMarking string `json:"classificationMarking"`
	CreatedAt             string `json:"createdAt"`
	CreatedBy             string `json:"createdBy"`
	Data                  string `json:"data"`
	DeleteOn              int64  `json:"deleteOn"`
	// Optional description for the file or folder.
	Description string `json:"description"`
	Filename    string `json:"filename"`
	FilePath    string `json:"filePath"`
	Keywords    string `json:"keywords"`
	ParentPath  string `json:"parentPath"`
	PathType    string `json:"pathType"`
	// For folders only. Comma separated list of user and group ids that should have
	// read access on this folder and the items nested in it.
	ReadACL string `json:"readAcl"`
	Size    int64  `json:"size"`
	// Array of provider/source specific tags for this data, used for implementing data
	// owner conditional access controls to restrict access to the data.
	Tags      []string `json:"tags"`
	UpdatedAt string   `json:"updatedAt"`
	UpdatedBy string   `json:"updatedBy"`
	// For folders only. Comma separated list of user and group ids that should have
	// write access on this folder and the items nested in it.
	WriteACL string `json:"writeAcl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Attachment            respjson.Field
		ClassificationMarking respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Data                  respjson.Field
		DeleteOn              respjson.Field
		Description           respjson.Field
		Filename              respjson.Field
		FilePath              respjson.Field
		Keywords              respjson.Field
		ParentPath            respjson.Field
		PathType              respjson.Field
		ReadACL               respjson.Field
		Size                  respjson.Field
		Tags                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		WriteACL              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScsEntity) RawJSON() string { return r.JSON.raw }
func (r *ScsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScV2UpdateParams struct {
	// The complete path for the object to be updated.
	Path string `query:"path,required" json:"-"`
	// Whether or not to send a notification that the target file/folder was updated.
	SendNotification param.Opt[bool] `query:"sendNotification,omitzero" json:"-"`
	// Classification marking of the folder or file in IC/CAPCO portion-marked format.
	ClassificationMarking param.Opt[string] `json:"classificationMarking,omitzero"`
	DeleteOn              param.Opt[int64]  `json:"deleteOn,omitzero"`
	// Optional description for the file or folder.
	Description param.Opt[string] `json:"description,omitzero"`
	// For folders only. Comma separated list of user and group ids that should have
	// read access on this folder and the items nested in it.
	ReadACL param.Opt[string] `json:"readAcl,omitzero"`
	// For folders only. Comma separated list of user and group ids that should have
	// write access on this folder and the items nested in it.
	WriteACL param.Opt[string] `json:"writeAcl,omitzero"`
	// Array of provider/source specific tags for this data, used for implementing data
	// owner conditional access controls to restrict access to the data.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ScV2UpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ScV2UpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScV2UpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ScV2UpdateParams]'s query parameters as `url.Values`.
func (r ScV2UpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScV2ListParams struct {
	// The base path to list
	Path        string           `query:"path,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	// The order in which entries should be sorted
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
	// The starting point for pagination results, usually set to the value of the
	// SEARCH_AFTER header returned in the previous request.
	SearchAfter param.Opt[string] `query:"searchAfter,omitzero" json:"-"`
	// The number of results to retrieve.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// The field on which to sort entries
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ScV2ListParams]'s query parameters as `url.Values`.
func (r ScV2ListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScV2DeleteParams struct {
	// The complete path for the object to be deleted. Must start with '/'.
	Path string `query:"path,required" json:"-"`
	paramObj
}

// URLQuery serializes [ScV2DeleteParams]'s query parameters as `url.Values`.
func (r ScV2DeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScV2CopyParams struct {
	// The path of the file or folder to copy. Must start with '/'.
	FromPath string `query:"fromPath,required" json:"-"`
	// The destination path to copy to. Must start with '/'.
	ToPath string `query:"toPath,required" json:"-"`
	paramObj
}

// URLQuery serializes [ScV2CopyParams]'s query parameters as `url.Values`.
func (r ScV2CopyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScV2FileUploadParams struct {
	// Classification marking of uploaded document. If folders are created, they will
	// also have this classification marking.
	ClassificationMarking string `query:"classificationMarking,required" json:"-"`
	// The complete path for the upload including filename. Will attempt to create
	// folders in path if necessary. Must start with '/'.
	Path string `query:"path,required" json:"-"`
	// Length of time after which to automatically delete the file.
	DeleteAfter param.Opt[string] `query:"deleteAfter,omitzero" json:"-"`
	// Optional description of uploaded document.
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	// Whether or not to overwrite a file with the same name and path, if one exists.
	Overwrite param.Opt[bool] `query:"overwrite,omitzero" json:"-"`
	// Whether or not to send a notification that this file was uploaded.
	SendNotification param.Opt[bool] `query:"sendNotification,omitzero" json:"-"`
	// Optional array of provider/source specific tags for this data, used for
	// implementing data owner conditional access controls to restrict access to the
	// data.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	paramObj
}

func (r ScV2FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [ScV2FileUploadParams]'s query parameters as `url.Values`.
func (r ScV2FileUploadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScV2FolderNewParams struct {
	// Path to create. Will attempt to create all folders in the path that do not
	// exist. Must start and end with '/'.
	Path string `query:"path,required" json:"-"`
	// Whether or not to send a notification that this folder was created.
	SendNotification param.Opt[bool] `query:"sendNotification,omitzero" json:"-"`
	// Classification marking of the folder or file in IC/CAPCO portion-marked format.
	ClassificationMarking param.Opt[string] `json:"classificationMarking,omitzero"`
	DeleteOn              param.Opt[int64]  `json:"deleteOn,omitzero"`
	// Optional description for the file or folder.
	Description param.Opt[string] `json:"description,omitzero"`
	// For folders only. Comma separated list of user and group ids that should have
	// read access on this folder and the items nested in it.
	ReadACL param.Opt[string] `json:"readAcl,omitzero"`
	// For folders only. Comma separated list of user and group ids that should have
	// write access on this folder and the items nested in it.
	WriteACL param.Opt[string] `json:"writeAcl,omitzero"`
	// Array of provider/source specific tags for this data, used for implementing data
	// owner conditional access controls to restrict access to the data.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ScV2FolderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ScV2FolderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScV2FolderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ScV2FolderNewParams]'s query parameters as `url.Values`.
func (r ScV2FolderNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScV2MoveParams struct {
	// The path of the file or folder to move or rename. Must start with '/'.
	FromPath string `query:"fromPath,required" json:"-"`
	// The destination path of the file or folder after moving or renaming. Must start
	// with '/'.
	ToPath string `query:"toPath,required" json:"-"`
	paramObj
}

// URLQuery serializes [ScV2MoveParams]'s query parameters as `url.Values`.
func (r ScV2MoveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScV2SearchParams struct {
	// The order in which entries should be sorted
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
	// The starting point for pagination results, usually set to the value of the
	// SEARCH_AFTER header returned in the previous request.
	SearchAfter param.Opt[string] `query:"searchAfter,omitzero" json:"-"`
	// The number of results to retrieve.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// The field on which to sort entries
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// A search criterion, which can be a simple field comparison or a logical
	// combination of other criteria.
	Query SearchCriterionUnionParam `json:"query,omitzero"`
	paramObj
}

func (r ScV2SearchParams) MarshalJSON() (data []byte, err error) {
	type shadow ScV2SearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScV2SearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ScV2SearchParams]'s query parameters as `url.Values`.
func (r ScV2SearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
