// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiform"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
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
func (r *ScPathService) New(ctx context.Context, params ScPathNewParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := "scs/path"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ScPathNewParams struct {
	// The full path to create, including path and file name
	ID string `query:"id,required" json:"-"`
	// Classification (ex. U//FOUO)
	ClassificationMarking string `query:"classificationMarking,required" json:"-"`
	Body                  io.Reader
	// Description
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	// Whether or not to overwrite a file with the same name and path, if one exists.
	Overwrite param.Opt[bool] `query:"overwrite,omitzero" json:"-"`
	// Tags
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	paramObj
}

func (r ScPathNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
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

// URLQuery serializes [ScPathNewParams]'s query parameters as `url.Values`.
func (r ScPathNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
