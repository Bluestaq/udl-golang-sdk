// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// VideoService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVideoService] method instead.
type VideoService struct {
	Options []option.RequestOption
	History VideoHistoryService
}

// NewVideoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVideoService(opts ...option.RequestOption) (r VideoService) {
	r = VideoService{}
	r.Options = opts
	r.History = NewVideoHistoryService(opts...)
	return
}

// Service operation to take a single Video Stream record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *VideoService) New(ctx context.Context, body VideoNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/video"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *VideoService) List(ctx context.Context, query VideoListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[VideoListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/video"
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
func (r *VideoService) ListAutoPaging(ctx context.Context, query VideoListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[VideoListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *VideoService) Count(ctx context.Context, query VideoCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/video/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Video Stream record by its unique ID passed as
// a path parameter.
func (r *VideoService) Get(ctx context.Context, id string, query VideoGetParams, opts ...option.RequestOption) (res *VideoStreamsFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/video/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to retrieve player URL and token for given stream name and
// source.
func (r *VideoService) GetPlayerStreamingInfo(ctx context.Context, query VideoGetPlayerStreamingInfoParams, opts ...option.RequestOption) (res *VideoGetPlayerStreamingInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/video/getPlayerStreamingInfo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to retrieve publisher URL and token for given stream name and
// source.
func (r *VideoService) GetPublisherStreamingInfo(ctx context.Context, query VideoGetPublisherStreamingInfoParams, opts ...option.RequestOption) (res *VideoGetPublisherStreamingInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/video/getPublisherStreamingInfo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to retrieve a static pre-configured SRT or UDP streamfile URL
// for given stream name and source.
func (r *VideoService) GetStreamFile(ctx context.Context, query VideoGetStreamFileParams, opts ...option.RequestOption) (res *VideoGetStreamFileResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/video/getStreamFile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *VideoService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *VideoQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/video/queryhelp"
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
func (r *VideoService) Tuple(ctx context.Context, query VideoTupleParams, opts ...option.RequestOption) (res *[]VideoStreamsFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/video/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The table captures metadata associated with the published videos in UDL.
type VideoListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode VideoListResponseDataMode `json:"dataMode,required"`
	// Description/notes associated with the video stream.
	Description string `json:"description,required"`
	// Name of the video stream.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Video Publisher Token.
	TokenValue string `json:"tokenValue,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	Origin    string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The timestamp when the stream is available from. The unit is ISO 8601 format.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// The timestamp when the stream is available until. The unit is ISO 8601 format.
	StopTime time.Time `json:"stopTime" format:"date-time"`
	// Video Streaming Support URLs.
	VideoURLs []string `json:"videoUrls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		TokenValue            respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		StartTime             respjson.Field
		StopTime              respjson.Field
		VideoURLs             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoListResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
type VideoListResponseDataMode string

const (
	VideoListResponseDataModeReal      VideoListResponseDataMode = "REAL"
	VideoListResponseDataModeTest      VideoListResponseDataMode = "TEST"
	VideoListResponseDataModeSimulated VideoListResponseDataMode = "SIMULATED"
	VideoListResponseDataModeExercise  VideoListResponseDataMode = "EXERCISE"
)

// The table captures metadata associated with the published videos in UDL.
type VideoGetPlayerStreamingInfoResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode VideoGetPlayerStreamingInfoResponseDataMode `json:"dataMode,required"`
	// Description/notes associated with the video stream.
	Description string `json:"description,required"`
	// Name of the video stream.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Video Publisher Token.
	TokenValue string `json:"tokenValue,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	Origin    string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The timestamp when the stream is available from. The unit is ISO 8601 format.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// The timestamp when the stream is available until. The unit is ISO 8601 format.
	StopTime time.Time `json:"stopTime" format:"date-time"`
	// Video Streaming Support URLs.
	VideoURLs []string `json:"videoUrls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		TokenValue            respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		StartTime             respjson.Field
		StopTime              respjson.Field
		VideoURLs             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoGetPlayerStreamingInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoGetPlayerStreamingInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
type VideoGetPlayerStreamingInfoResponseDataMode string

const (
	VideoGetPlayerStreamingInfoResponseDataModeReal      VideoGetPlayerStreamingInfoResponseDataMode = "REAL"
	VideoGetPlayerStreamingInfoResponseDataModeTest      VideoGetPlayerStreamingInfoResponseDataMode = "TEST"
	VideoGetPlayerStreamingInfoResponseDataModeSimulated VideoGetPlayerStreamingInfoResponseDataMode = "SIMULATED"
	VideoGetPlayerStreamingInfoResponseDataModeExercise  VideoGetPlayerStreamingInfoResponseDataMode = "EXERCISE"
)

// The table captures metadata associated with the published videos in UDL.
type VideoGetPublisherStreamingInfoResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode VideoGetPublisherStreamingInfoResponseDataMode `json:"dataMode,required"`
	// Description/notes associated with the video stream.
	Description string `json:"description,required"`
	// Name of the video stream.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Video Publisher Token.
	TokenValue string `json:"tokenValue,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	Origin    string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The timestamp when the stream is available from. The unit is ISO 8601 format.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// The timestamp when the stream is available until. The unit is ISO 8601 format.
	StopTime time.Time `json:"stopTime" format:"date-time"`
	// Video Streaming Support URLs.
	VideoURLs []string `json:"videoUrls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		TokenValue            respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		StartTime             respjson.Field
		StopTime              respjson.Field
		VideoURLs             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoGetPublisherStreamingInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoGetPublisherStreamingInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
type VideoGetPublisherStreamingInfoResponseDataMode string

const (
	VideoGetPublisherStreamingInfoResponseDataModeReal      VideoGetPublisherStreamingInfoResponseDataMode = "REAL"
	VideoGetPublisherStreamingInfoResponseDataModeTest      VideoGetPublisherStreamingInfoResponseDataMode = "TEST"
	VideoGetPublisherStreamingInfoResponseDataModeSimulated VideoGetPublisherStreamingInfoResponseDataMode = "SIMULATED"
	VideoGetPublisherStreamingInfoResponseDataModeExercise  VideoGetPublisherStreamingInfoResponseDataMode = "EXERCISE"
)

// The table captures metadata associated with the published videos in UDL.
type VideoGetStreamFileResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode VideoGetStreamFileResponseDataMode `json:"dataMode,required"`
	// Description/notes associated with the video stream.
	Description string `json:"description,required"`
	// Name of the video stream.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Video Publisher Token.
	TokenValue string `json:"tokenValue,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	Origin    string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The timestamp when the stream is available from. The unit is ISO 8601 format.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// The timestamp when the stream is available until. The unit is ISO 8601 format.
	StopTime time.Time `json:"stopTime" format:"date-time"`
	// Video Streaming Support URLs.
	VideoURLs []string `json:"videoUrls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		TokenValue            respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		StartTime             respjson.Field
		StopTime              respjson.Field
		VideoURLs             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoGetStreamFileResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoGetStreamFileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
type VideoGetStreamFileResponseDataMode string

const (
	VideoGetStreamFileResponseDataModeReal      VideoGetStreamFileResponseDataMode = "REAL"
	VideoGetStreamFileResponseDataModeTest      VideoGetStreamFileResponseDataMode = "TEST"
	VideoGetStreamFileResponseDataModeSimulated VideoGetStreamFileResponseDataMode = "SIMULATED"
	VideoGetStreamFileResponseDataModeExercise  VideoGetStreamFileResponseDataMode = "EXERCISE"
)

type VideoQueryhelpResponse struct {
	AodrSupported         bool                              `json:"aodrSupported"`
	ClassificationMarking string                            `json:"classificationMarking"`
	Description           string                            `json:"description"`
	HistorySupported      bool                              `json:"historySupported"`
	Name                  string                            `json:"name"`
	Parameters            []VideoQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                          `json:"requiredRoles"`
	RestSupported         bool                              `json:"restSupported"`
	SortSupported         bool                              `json:"sortSupported"`
	TypeName              string                            `json:"typeName"`
	Uri                   string                            `json:"uri"`
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
func (r VideoQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoQueryhelpResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *VideoQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode VideoNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Description/notes associated with the video stream.
	Description string `json:"description,required"`
	// Name of the video stream.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID     param.Opt[string] `json:"id,omitzero"`
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The timestamp when the stream is available from. The unit is ISO 8601 format.
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// The timestamp when the stream is available until. The unit is ISO 8601 format.
	StopTime param.Opt[time.Time] `json:"stopTime,omitzero" format:"date-time"`
	paramObj
}

func (r VideoNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
type VideoNewParamsDataMode string

const (
	VideoNewParamsDataModeReal      VideoNewParamsDataMode = "REAL"
	VideoNewParamsDataModeTest      VideoNewParamsDataMode = "TEST"
	VideoNewParamsDataModeSimulated VideoNewParamsDataMode = "SIMULATED"
	VideoNewParamsDataModeExercise  VideoNewParamsDataMode = "EXERCISE"
)

type VideoListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoListParams]'s query parameters as `url.Values`.
func (r VideoListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoCountParams]'s query parameters as `url.Values`.
func (r VideoCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoGetParams]'s query parameters as `url.Values`.
func (r VideoGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoGetPlayerStreamingInfoParams struct {
	// The video source name.
	SourceName string `query:"sourceName,required" json:"-"`
	// The video stream name.
	StreamName  string           `query:"streamName,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoGetPlayerStreamingInfoParams]'s query parameters as
// `url.Values`.
func (r VideoGetPlayerStreamingInfoParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoGetPublisherStreamingInfoParams struct {
	// The video source name.
	SourceName string `query:"sourceName,required" json:"-"`
	// The video stream name.
	StreamName  string           `query:"streamName,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoGetPublisherStreamingInfoParams]'s query parameters as
// `url.Values`.
func (r VideoGetPublisherStreamingInfoParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoGetStreamFileParams struct {
	// The video source name.
	SourceName string `query:"sourceName,required" json:"-"`
	// The video stream name.
	StreamName  string           `query:"streamName,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoGetStreamFileParams]'s query parameters as
// `url.Values`.
func (r VideoGetStreamFileParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoTupleParams]'s query parameters as `url.Values`.
func (r VideoTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
