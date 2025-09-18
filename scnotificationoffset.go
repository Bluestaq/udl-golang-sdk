// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// ScNotificationOffsetService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScNotificationOffsetService] method instead.
type ScNotificationOffsetService struct {
	Options []option.RequestOption
}

// NewScNotificationOffsetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScNotificationOffsetService(opts ...option.RequestOption) (r ScNotificationOffsetService) {
	r = ScNotificationOffsetService{}
	r.Options = opts
	return
}

// Retrieve the min and max offsets of the SCS Event Notification Kafka topic.
func (r *ScNotificationOffsetService) Get(ctx context.Context, opts ...option.RequestOption) (res *ScNotificationOffsetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "scs/notifications/offsets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the current/latest offset for the SCS Event Notification Kafka topic.
func (r *ScNotificationOffsetService) GetLatest(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "scs/notifications/getLatestOffset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type ScNotificationOffsetGetResponse struct {
	// The maximum (latest) Kafka offset for this topic.
	MaxOffset int64 `json:"maxOffset"`
	// The minimum (oldest) Kafka offset for this topic.
	MinOffset int64 `json:"minOffset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxOffset   respjson.Field
		MinOffset   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScNotificationOffsetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ScNotificationOffsetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
