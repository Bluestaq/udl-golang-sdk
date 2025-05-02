// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// ScFolderService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScFolderService] method instead.
type ScFolderService struct {
	Options []option.RequestOption
}

// NewScFolderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScFolderService(opts ...option.RequestOption) (r ScFolderService) {
	r = ScFolderService{}
	r.Options = opts
	return
}

// Creates a new folder that is passed as part of the path. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *ScFolderService) New(ctx context.Context, body ScFolderNewParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := "scs/folder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a FileData object representing the folder ID that is visible to the
// calling user.
func (r *ScFolderService) Get(ctx context.Context, query ScFolderGetParams, opts ...option.RequestOption) (res *shared.FileData, err error) {
	opts = append(r.Options[:], opts...)
	path := "scs/folder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// operation to update folders metadata. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *ScFolderService) Update(ctx context.Context, body ScFolderUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/folder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

type ScFolderNewParams struct {
	// Path to create folder.
	ID string `query:"id,required" json:"-"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `query:"classificationMarking,required" json:"-"`
	// Optional description to include on folder.
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	// Comma separated list of user ids who can read contents of the folder.
	Read param.Opt[string] `query:"read,omitzero" json:"-"`
	// Comma separated list of tags to add to the folder.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	// Comma separated list of user ids who can write to the folder.
	Write param.Opt[string] `query:"write,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ScFolderNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ScFolderNewParams]'s query parameters as `url.Values`.
func (r ScFolderNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScFolderGetParams struct {
	// The folder ID
	ID          string           `query:"id,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ScFolderGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ScFolderGetParams]'s query parameters as `url.Values`.
func (r ScFolderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScFolderUpdateParams struct {
	FileData shared.FileDataParam
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ScFolderUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ScFolderUpdateParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.FileData)
}
