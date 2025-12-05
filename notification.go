// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// NotificationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationService] method instead.
type NotificationService struct {
	Options []option.RequestOption
	History NotificationHistoryService
}

// NewNotificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationService(opts ...option.RequestOption) (r NotificationService) {
	r = NotificationService{}
	r.Options = opts
	r.History = NewNotificationHistoryService(opts...)
	return
}

// Service operation to push a generic Notification/Alert JSON message into the
// UDL. This operation accepts a UDL-formatted Notification JSON notification/alert
// message. See the Notification schema for required fields such as
// classificationMarking, msgType, etc. Messages pushed through this service may be
// pulled via Secure Messaging and historical REST services. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *NotificationService) New(ctx context.Context, body NotificationNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/notification"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *NotificationService) List(ctx context.Context, query NotificationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[NotificationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/notification"
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

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *NotificationService) ListAutoPaging(ctx context.Context, query NotificationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[NotificationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *NotificationService) Count(ctx context.Context, query NotificationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/notification/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to push a Notification/Alert message into the UDL. This
// operation uses query parameters to pass UDL-required fields such as
// classificationMarking, msgType, etc and takes a raw string payload which can be
// XML, JSON, or plain text. The preferred mechanism for posting notifications is
// to use the standard POST which takes the proper UDL JSON Notification schema as
// this service may convert the message in an undesirable manner. The service will
// wrap the passed payload with the appropriate UDL JSON schema, using escape
// characters as necessary for the payload to produce a valid JSON document. XML
// payloads (not recommended) are automatically converted to JSON. Messages pushed
// through this service may be pulled via Secure Messaging and historical REST
// services. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *NotificationService) NewRaw(ctx context.Context, params NotificationNewRawParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/notification/createRaw"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Service operation to get a single notification by its unique ID passed as a path
// parameter.
func (r *NotificationService) Get(ctx context.Context, id string, query NotificationGetParams, opts ...option.RequestOption) (res *shared.NotificationFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/notification/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *NotificationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *NotificationQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/notification/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Service operation to dynamically query data and only return specified
// columns/fields. Requested columns are specified by the 'columns' query parameter
// and should be a comma separated list of valid fields for the specified data
// type. classificationMarking is always returned. See the queryhelp operation
// (/udl/<datatype>/queryhelp) for more details on valid/required query parameter
// information. An example URI: /udl/elset/tuple?columns=satNo,period&epoch=>now-5
// hours would return the satNo and period of elsets with an epoch greater than 5
// hours ago.
func (r *NotificationService) Tuple(ctx context.Context, query NotificationTupleParams, opts ...option.RequestOption) (res *[]shared.NotificationFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/notification/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of client generated notification data. Contains a message
// type and message body field to store notification information.
type NotificationListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode NotificationListResponseDataMode `json:"dataMode,required"`
	// The message body content as a text string, XML, JSON, etc. If JSON is used for
	// the msgBody, it should be 'inline' with the notification message JSON (without
	// quotation marks or escape characters). Size of the msg body cannot be over 1MB.
	MsgBody string `json:"msgBody,required"`
	// Source provided message type.
	MsgType string `json:"msgType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The date and time the notification was created, auto-generated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The user that created the notification.
	CreatedBy string `json:"createdBy"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional array of provider/source specific tags for this data, used for
	// implementing data owner conditional access controls to restrict access to the
	// data.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		MsgBody               respjson.Field
		MsgType               respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type NotificationListResponseDataMode string

const (
	NotificationListResponseDataModeReal      NotificationListResponseDataMode = "REAL"
	NotificationListResponseDataModeTest      NotificationListResponseDataMode = "TEST"
	NotificationListResponseDataModeSimulated NotificationListResponseDataMode = "SIMULATED"
	NotificationListResponseDataModeExercise  NotificationListResponseDataMode = "EXERCISE"
)

type NotificationQueryhelpResponse struct {
	AodrSupported         bool                         `json:"aodrSupported"`
	ClassificationMarking string                       `json:"classificationMarking"`
	Description           string                       `json:"description"`
	HistorySupported      bool                         `json:"historySupported"`
	Name                  string                       `json:"name"`
	Parameters            []shared.ParamDescriptorResp `json:"parameters"`
	RequiredRoles         []string                     `json:"requiredRoles"`
	RestSupported         bool                         `json:"restSupported"`
	SortSupported         bool                         `json:"sortSupported"`
	TypeName              string                       `json:"typeName"`
	Uri                   string                       `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AodrSupported         respjson.Field
		ClassificationMarking respjson.Field
		Description           respjson.Field
		HistorySupported      respjson.Field
		Name                  respjson.Field
		Parameters            respjson.Field
		RequiredRoles         respjson.Field
		RestSupported         respjson.Field
		SortSupported         respjson.Field
		TypeName              respjson.Field
		Uri                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode NotificationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The message body content as a text string, XML, JSON, etc. If JSON is used for
	// the msgBody, it should be 'inline' with the notification message JSON (without
	// quotation marks or escape characters). Size of the msg body cannot be over 1MB.
	MsgBody string `json:"msgBody,required"`
	// Source provided message type.
	MsgType string `json:"msgType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional array of provider/source specific tags for this data, used for
	// implementing data owner conditional access controls to restrict access to the
	// data.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r NotificationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type NotificationNewParamsDataMode string

const (
	NotificationNewParamsDataModeReal      NotificationNewParamsDataMode = "REAL"
	NotificationNewParamsDataModeTest      NotificationNewParamsDataMode = "TEST"
	NotificationNewParamsDataModeSimulated NotificationNewParamsDataMode = "SIMULATED"
	NotificationNewParamsDataModeExercise  NotificationNewParamsDataMode = "EXERCISE"
)

type NotificationListParams struct {
	// Time the row was created in the database. (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationListParams]'s query parameters as `url.Values`.
func (r NotificationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationCountParams struct {
	// Time the row was created in the database. (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationCountParams]'s query parameters as
// `url.Values`.
func (r NotificationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationNewRawParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `query:"classificationMarking,required" json:"-"`
	// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
	//
	// EXERCISE: Data pertaining to a government or military exercise. The data may
	// include both real and simulated data.
	//
	// REAL: Data collected or produced that pertains to real-world objects, events,
	// and analysis.
	//
	// SIMULATED: Synthetic data generated by a model to mimic real-world datasets.
	//
	// TEST: Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	DataMode string `query:"dataMode,required" json:"-"`
	// Source provided message type.
	MsgType string `query:"msgType,required" json:"-"`
	// Origin of the data.
	Origin string `query:"origin,required" json:"-"`
	// Source of the data.
	Source string `query:"source,required" json:"-"`
	Body   string
	// Optional message identifier...if not provided an id will be automatically
	// created.
	MsgID param.Opt[string] `query:"msgId,omitzero" json:"-"`
	// Optional list of provider/source specific tags for this data.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

func (r NotificationNewRawParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *NotificationNewRawParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [NotificationNewRawParams]'s query parameters as
// `url.Values`.
func (r NotificationNewRawParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationGetParams]'s query parameters as `url.Values`.
func (r NotificationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time the row was created in the database. (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationTupleParams]'s query parameters as
// `url.Values`.
func (r NotificationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
