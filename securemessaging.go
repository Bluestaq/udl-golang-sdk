// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
func (r *SecureMessagingService) GetMessages(ctx context.Context, offset int64, params SecureMessagingGetMessagesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.Topic == "" {
		err = errors.New("missing required topic parameter")
		return
	}
	path := fmt.Sprintf("sm/getMessages/%s/%v", params.Topic, offset)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

// Retrieve the list of available secure messaging topics or data types available.
func (r *SecureMessagingService) ListTopics(ctx context.Context, opts ...option.RequestOption) (res *[]TopicDetails, err error) {
	opts = append(r.Options[:], opts...)
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Description      resp.Field
		MaxPos           resp.Field
		MinPos           resp.Field
		Topic            resp.Field
		UdlOpenAPISchema resp.Field
		ExtraFields      map[string]resp.Field
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecureMessagingDescribeTopicParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecureMessagingGetLatestOffsetParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecureMessagingGetMessagesParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SecureMessagingGetMessagesParams]'s query parameters as
// `url.Values`.
func (r SecureMessagingGetMessagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
