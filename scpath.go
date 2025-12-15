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
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
)

// ScPathService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScPathService] method instead.
type ScPathService struct {
	Options []option.RequestOption
}

// NewScPathService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScPathService(opts ...option.RequestOption) (r ScPathService) {
	r = ScPathService{}
	r.Options = opts
	return
}

// Creates the path and uploads file that is passed. If folder exist it will only
// create folders that are missing. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
//
// Deprecated: deprecated
func (r *ScPathService) NewWithFile(ctx context.Context, fileContent io.Reader, params ScPathNewWithFileParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", fileContent)}, opts...)
	path := "scs/path"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ScPathNewWithFileParams struct {
	// The full path to create, including path and file name
	ID string `query:"id,required" json:"-"`
	// Classification marking of the file being uploaded.
	ClassificationMarking string `query:"classificationMarking,required" json:"-"`
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

func (r ScPathNewWithFileParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [ScPathNewWithFileParams]'s query parameters as
// `url.Values`.
func (r ScPathNewWithFileParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
