// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// ScNotificationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScNotificationService] method instead.
type ScNotificationService struct {
	Options []option.RequestOption
	Offset  ScNotificationOffsetService
}

// NewScNotificationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScNotificationService(opts ...option.RequestOption) (r ScNotificationService) {
	r = ScNotificationService{}
	r.Options = opts
	r.Offset = NewScNotificationOffsetService(opts...)
	return
}

// Returns a list of notifications for items in a specific folder.
func (r *ScNotificationService) List(ctx context.Context, offset string, query ScNotificationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ScNotificationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if offset == "" {
		err = errors.New("missing required offset parameter")
		return
	}
	path := fmt.Sprintf("scs/notifications/%s", offset)
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

// Returns a list of notifications for items in a specific folder.
func (r *ScNotificationService) ListAutoPaging(ctx context.Context, offset string, query ScNotificationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ScNotificationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, offset, query, opts...))
}

// SCS Event Notification
type ScNotificationListResponse struct {
	// Any of "ROOT_WRITE", "UPLOAD_FILE", "CREATE_FOLDER", "DOWNLOAD_FILE",
	// "DOWNLOAD_FOLDER", "MOVE_RENAME_FILE", "MOVE_RENAME_FOLDER", "COPY_FILE",
	// "COPY_FOLDER", "UPDATE_FILE", "UPDATE_FOLDER", "DELETE_FILE", "DELETE_FOLDER",
	// "DELETE_EMPTY_FOLDER", "CROSS_DOMAIN", "SEND_NOTIFICATION", "DELETE_READ_ACL",
	// "DELETE_WRITE_ACL", "DELETE_FILE_TAGS", "DELETE_FOLDER_TAGS".
	Actions               []string `json:"actions"`
	ClassificationMarking string   `json:"classificationMarking"`
	CrossDomainTo         string   `json:"crossDomainTo"`
	Expires               string   `json:"expires"`
	Overwrite             bool     `json:"overwrite"`
	Path                  string   `json:"path"`
	Timestamp             string   `json:"timestamp"`
	User                  string   `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions               respjson.Field
		ClassificationMarking respjson.Field
		CrossDomainTo         respjson.Field
		Expires               respjson.Field
		Overwrite             respjson.Field
		Path                  respjson.Field
		Timestamp             respjson.Field
		User                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScNotificationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ScNotificationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScNotificationListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	// Path of the folder to retrieve notification for. Must start and end with /. If
	// no path is specified, all notifications will be returned.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ScNotificationListParams]'s query parameters as
// `url.Values`.
func (r ScNotificationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
