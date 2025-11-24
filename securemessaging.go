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

// SecureMessagingService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecureMessagingService] method instead.
type SecureMessagingService struct {
	Options []option.RequestOption
}

// NewSecureMessagingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecureMessagingService(opts ...option.RequestOption) (r SecureMessagingService) {
	r = SecureMessagingService{}
	r.Options = opts
	return
}

// Retrieve the details of the specified topic or data type.
func (r *SecureMessagingService) DescribeTopic(ctx context.Context, topic string, query SecureMessagingDescribeTopicParams, opts ...option.RequestOption) (res *TopicDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if topic == "" {
		err = errors.New("missing required topic parameter")
		return
	}
	path := fmt.Sprintf("sm/describeTopic/%s", topic)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns the current/latest offset for the passed topic name.
func (r *SecureMessagingService) GetLatestOffset(ctx context.Context, topic string, query SecureMessagingGetLatestOffsetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if topic == "" {
		err = errors.New("missing required topic parameter")
		return
	}
	path := fmt.Sprintf("sm/getLatestOffset/%s", topic)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Retrieve a set of messages from the given topic at the given offset. See Help >
// Secure Messaging API on Storefront for more details on how to use getMessages.
func (r *SecureMessagingService) GetMessages(ctx context.Context, offset int64, params SecureMessagingGetMessagesParams, opts ...option.RequestOption) (res *pagination.KafkaOffsetPage[interface{}], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*"), option.WithResponseInto(&raw)}, opts...)
	if params.Topic == "" {
		err = errors.New("missing required topic parameter")
		return
	}
	path := fmt.Sprintf("sm/getMessages/%s/%v", params.Topic, offset)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw, func(nextOffset int64) string {
		return fmt.Sprintf("sm/getMessages/%s/%v", params.Topic, nextOffset)
	})
	return res, nil
}

// Retrieve a set of messages from the given topic at the given offset. See Help >
// Secure Messaging API on Storefront for more details on how to use getMessages.
func (r *SecureMessagingService) GetMessagesAutoPaging(ctx context.Context, offset int64, params SecureMessagingGetMessagesParams, opts ...option.RequestOption) *pagination.KafkaOffsetPageAutoPager[interface{}] {
	return pagination.NewKafkaOffsetPageAutoPager(r.GetMessages(ctx, offset, params, opts...))
}

// Retrieve the list of available secure messaging topics or data types available.
func (r *SecureMessagingService) ListTopics(ctx context.Context, opts ...option.RequestOption) (res *[]TopicDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sm/listTopics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TopicDetails struct {
	// A description of the data content of this topic.
	Description string `json:"description"`
	// The maximum (latest) kafka offset for this topic.
	MaxPos int64 `json:"maxPos"`
	// The minimum (oldest) kafka offset for this topic.
	MinPos int64 `json:"minPos"`
	// The name of the topic in kafka.
	Topic string `json:"topic"`
	// The UDL schema that the objects in this topic apply to.
	UdlOpenAPISchema string `json:"udlOpenAPISchema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description      respjson.Field
		MaxPos           respjson.Field
		MinPos           respjson.Field
		Topic            respjson.Field
		UdlOpenAPISchema respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopicDetails) RawJSON() string { return r.JSON.raw }
func (r *TopicDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecureMessagingDescribeTopicParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecureMessagingDescribeTopicParams]'s query parameters as
// `url.Values`.
func (r SecureMessagingDescribeTopicParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SecureMessagingGetLatestOffsetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecureMessagingGetLatestOffsetParams]'s query parameters as
// `url.Values`.
func (r SecureMessagingGetLatestOffsetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SecureMessagingGetMessagesParams struct {
	Topic       string           `path:"topic,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecureMessagingGetMessagesParams]'s query parameters as
// `url.Values`.
func (r SecureMessagingGetMessagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
