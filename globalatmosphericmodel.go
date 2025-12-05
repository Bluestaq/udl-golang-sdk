// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// GlobalAtmosphericModelService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalAtmosphericModelService] method instead.
type GlobalAtmosphericModelService struct {
	Options []option.RequestOption
	History GlobalAtmosphericModelHistoryService
}

// NewGlobalAtmosphericModelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalAtmosphericModelService(opts ...option.RequestOption) (r GlobalAtmosphericModelService) {
	r = GlobalAtmosphericModelService{}
	r.Options = opts
	r.History = NewGlobalAtmosphericModelHistoryService(opts...)
	return
}

// Service operation to get a single GlobalAtmosphericModel record by its unique ID
// passed as a path parameter.
func (r *GlobalAtmosphericModelService) Get(ctx context.Context, id string, query GlobalAtmosphericModelGetParams, opts ...option.RequestOption) (res *GlobalAtmosphericModelGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/globalatmosphericmodel/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *GlobalAtmosphericModelService) List(ctx context.Context, query GlobalAtmosphericModelListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GlobalAtmosphericModelListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/globalatmosphericmodel"
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
func (r *GlobalAtmosphericModelService) ListAutoPaging(ctx context.Context, query GlobalAtmosphericModelListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GlobalAtmosphericModelListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *GlobalAtmosphericModelService) Count(ctx context.Context, query GlobalAtmosphericModelCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/globalatmosphericmodel/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single GlobalAtmosphericModel compressed data file by
// its unique ID passed as a path parameter. The compressed data file is returned
// as an attachment Content-Disposition.
func (r *GlobalAtmosphericModelService) GetFile(ctx context.Context, id string, query GlobalAtmosphericModelGetFileParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/globalatmosphericmodel/getFile/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *GlobalAtmosphericModelService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *GlobalAtmosphericModelQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/globalatmosphericmodel/queryhelp"
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
func (r *GlobalAtmosphericModelService) Tuple(ctx context.Context, query GlobalAtmosphericModelTupleParams, opts ...option.RequestOption) (res *[]GlobalAtmosphericModelTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/globalatmosphericmodel/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a file with its metadata.
//
// The request body requires a zip file containing exactly two files:\
// 1\) A file with the `.json` file extension whose content conforms to the `GlobalAtmosphericModel_Ingest`
// schema.\
// 2\) A file with the `.geojson` file extension.
//
// The JSON and GEOJSON files will be associated with each other other via the `id`
// field. Query the metadata via `GET /udl/globalatmosphericmodel` and use
// `GET /udl/globalatmosphericmodel/getFile/{id}` to retrieve the compressed
// GEOJSON file as `.gz` extension.
//
// This operation only accepts application/zip media. The application/json request
// body is documented to provide a convenient reference to the ingest schema.
//
// This operation is intended to be used for automated feeds into UDL. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *GlobalAtmosphericModelService) UnvalidatedPublish(ctx context.Context, body GlobalAtmosphericModelUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-globalatmosphericmodel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The GlobalAtmosphericModel service provides atmospheric model output data for
// use in space situational awareness such as the Global Total Electron Content
// (2D) data, Global Total Electron Density (3D) data, etc.
type GlobalAtmosphericModelGetResponse struct {
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
	DataMode GlobalAtmosphericModelGetResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Type of data associated with this record (e.g. Global Total Electron Density,
	// Global Total Electron Content).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Model execution cadence, in minutes.
	Cadence int64 `json:"cadence"`
	// MD5 value of the data file. If not provided, the ingest/create operation will
	// automatically generate the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A unique identification code or label assigned to a particular source from which
	// atmospheric data originates.
	DataSourceIdentifier string `json:"dataSourceIdentifier"`
	// Ending altitude of model outputs, in kilometers.
	EndAlt float64 `json:"endAlt"`
	// WGS-84 ending latitude of model output, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndLat float64 `json:"endLat"`
	// WGS-84 ending longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndLon float64 `json:"endLon"`
	// The file name of the uploaded file.
	Filename string `json:"filename"`
	// The uploaded file size, in bytes. The maximum file size for this service is
	// 104857600 bytes (100MB). Files exceeding the maximum size will be rejected.
	Filesize int64 `json:"filesize"`
	// Number of altitude points.
	NumAlt int64 `json:"numAlt"`
	// Number of latitude points.
	NumLat int64 `json:"numLat"`
	// Number of longitude points.
	NumLon int64 `json:"numLon"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The time that this record was created, in ISO 8601 UTC format with millisecond
	// precision.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Starting altitude of model outputs, in kilometers.
	StartAlt float64 `json:"startAlt"`
	// WGS-84 starting latitude of model output, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	StartLat float64 `json:"startLat"`
	// WGS-84 starting longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	StartLon float64 `json:"startLon"`
	// State value indicating whether the values in this record are PREDICTED or
	// OBSERVED.
	State string `json:"state"`
	// Separation in latitude between subsequent model outputs, in degrees.
	StepLat float64 `json:"stepLat"`
	// Separation in longitude between subsequent model outputs, in degrees.
	StepLon float64 `json:"stepLon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		Cadence               respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataSourceIdentifier  respjson.Field
		EndAlt                respjson.Field
		EndLat                respjson.Field
		EndLon                respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		NumAlt                respjson.Field
		NumLat                respjson.Field
		NumLon                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ReportTime            respjson.Field
		SourceDl              respjson.Field
		StartAlt              respjson.Field
		StartLat              respjson.Field
		StartLon              respjson.Field
		State                 respjson.Field
		StepLat               respjson.Field
		StepLon               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalAtmosphericModelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalAtmosphericModelGetResponse) UnmarshalJSON(data []byte) error {
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
type GlobalAtmosphericModelGetResponseDataMode string

const (
	GlobalAtmosphericModelGetResponseDataModeReal      GlobalAtmosphericModelGetResponseDataMode = "REAL"
	GlobalAtmosphericModelGetResponseDataModeTest      GlobalAtmosphericModelGetResponseDataMode = "TEST"
	GlobalAtmosphericModelGetResponseDataModeSimulated GlobalAtmosphericModelGetResponseDataMode = "SIMULATED"
	GlobalAtmosphericModelGetResponseDataModeExercise  GlobalAtmosphericModelGetResponseDataMode = "EXERCISE"
)

// The GlobalAtmosphericModel service provides atmospheric model output data for
// use in space situational awareness such as the Global Total Electron Content
// (2D) data, Global Total Electron Density (3D) data, etc.
type GlobalAtmosphericModelListResponse struct {
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
	DataMode GlobalAtmosphericModelListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Type of data associated with this record (e.g. Global Total Electron Density,
	// Global Total Electron Content).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Model execution cadence, in minutes.
	Cadence int64 `json:"cadence"`
	// MD5 value of the data file. If not provided, the ingest/create operation will
	// automatically generate the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A unique identification code or label assigned to a particular source from which
	// atmospheric data originates.
	DataSourceIdentifier string `json:"dataSourceIdentifier"`
	// Ending altitude of model outputs, in kilometers.
	EndAlt float64 `json:"endAlt"`
	// WGS-84 ending latitude of model output, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndLat float64 `json:"endLat"`
	// WGS-84 ending longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndLon float64 `json:"endLon"`
	// The file name of the uploaded file.
	Filename string `json:"filename"`
	// The uploaded file size, in bytes. The maximum file size for this service is
	// 104857600 bytes (100MB). Files exceeding the maximum size will be rejected.
	Filesize int64 `json:"filesize"`
	// Number of altitude points.
	NumAlt int64 `json:"numAlt"`
	// Number of latitude points.
	NumLat int64 `json:"numLat"`
	// Number of longitude points.
	NumLon int64 `json:"numLon"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The time that this record was created, in ISO 8601 UTC format with millisecond
	// precision.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Starting altitude of model outputs, in kilometers.
	StartAlt float64 `json:"startAlt"`
	// WGS-84 starting latitude of model output, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	StartLat float64 `json:"startLat"`
	// WGS-84 starting longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	StartLon float64 `json:"startLon"`
	// State value indicating whether the values in this record are PREDICTED or
	// OBSERVED.
	State string `json:"state"`
	// Separation in latitude between subsequent model outputs, in degrees.
	StepLat float64 `json:"stepLat"`
	// Separation in longitude between subsequent model outputs, in degrees.
	StepLon float64 `json:"stepLon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		Cadence               respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataSourceIdentifier  respjson.Field
		EndAlt                respjson.Field
		EndLat                respjson.Field
		EndLon                respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		NumAlt                respjson.Field
		NumLat                respjson.Field
		NumLon                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ReportTime            respjson.Field
		SourceDl              respjson.Field
		StartAlt              respjson.Field
		StartLat              respjson.Field
		StartLon              respjson.Field
		State                 respjson.Field
		StepLat               respjson.Field
		StepLon               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalAtmosphericModelListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalAtmosphericModelListResponse) UnmarshalJSON(data []byte) error {
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
type GlobalAtmosphericModelListResponseDataMode string

const (
	GlobalAtmosphericModelListResponseDataModeReal      GlobalAtmosphericModelListResponseDataMode = "REAL"
	GlobalAtmosphericModelListResponseDataModeTest      GlobalAtmosphericModelListResponseDataMode = "TEST"
	GlobalAtmosphericModelListResponseDataModeSimulated GlobalAtmosphericModelListResponseDataMode = "SIMULATED"
	GlobalAtmosphericModelListResponseDataModeExercise  GlobalAtmosphericModelListResponseDataMode = "EXERCISE"
)

type GlobalAtmosphericModelQueryHelpResponse struct {
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
func (r GlobalAtmosphericModelQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalAtmosphericModelQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The GlobalAtmosphericModel service provides atmospheric model output data for
// use in space situational awareness such as the Global Total Electron Content
// (2D) data, Global Total Electron Density (3D) data, etc.
type GlobalAtmosphericModelTupleResponse struct {
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
	DataMode GlobalAtmosphericModelTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Type of data associated with this record (e.g. Global Total Electron Density,
	// Global Total Electron Content).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Model execution cadence, in minutes.
	Cadence int64 `json:"cadence"`
	// MD5 value of the data file. If not provided, the ingest/create operation will
	// automatically generate the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A unique identification code or label assigned to a particular source from which
	// atmospheric data originates.
	DataSourceIdentifier string `json:"dataSourceIdentifier"`
	// Ending altitude of model outputs, in kilometers.
	EndAlt float64 `json:"endAlt"`
	// WGS-84 ending latitude of model output, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndLat float64 `json:"endLat"`
	// WGS-84 ending longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndLon float64 `json:"endLon"`
	// The file name of the uploaded file.
	Filename string `json:"filename"`
	// The uploaded file size, in bytes. The maximum file size for this service is
	// 104857600 bytes (100MB). Files exceeding the maximum size will be rejected.
	Filesize int64 `json:"filesize"`
	// Number of altitude points.
	NumAlt int64 `json:"numAlt"`
	// Number of latitude points.
	NumLat int64 `json:"numLat"`
	// Number of longitude points.
	NumLon int64 `json:"numLon"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The time that this record was created, in ISO 8601 UTC format with millisecond
	// precision.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Starting altitude of model outputs, in kilometers.
	StartAlt float64 `json:"startAlt"`
	// WGS-84 starting latitude of model output, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	StartLat float64 `json:"startLat"`
	// WGS-84 starting longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	StartLon float64 `json:"startLon"`
	// State value indicating whether the values in this record are PREDICTED or
	// OBSERVED.
	State string `json:"state"`
	// Separation in latitude between subsequent model outputs, in degrees.
	StepLat float64 `json:"stepLat"`
	// Separation in longitude between subsequent model outputs, in degrees.
	StepLon float64 `json:"stepLon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		Cadence               respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataSourceIdentifier  respjson.Field
		EndAlt                respjson.Field
		EndLat                respjson.Field
		EndLon                respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		NumAlt                respjson.Field
		NumLat                respjson.Field
		NumLon                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ReportTime            respjson.Field
		SourceDl              respjson.Field
		StartAlt              respjson.Field
		StartLat              respjson.Field
		StartLon              respjson.Field
		State                 respjson.Field
		StepLat               respjson.Field
		StepLon               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalAtmosphericModelTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalAtmosphericModelTupleResponse) UnmarshalJSON(data []byte) error {
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
type GlobalAtmosphericModelTupleResponseDataMode string

const (
	GlobalAtmosphericModelTupleResponseDataModeReal      GlobalAtmosphericModelTupleResponseDataMode = "REAL"
	GlobalAtmosphericModelTupleResponseDataModeTest      GlobalAtmosphericModelTupleResponseDataMode = "TEST"
	GlobalAtmosphericModelTupleResponseDataModeSimulated GlobalAtmosphericModelTupleResponseDataMode = "SIMULATED"
	GlobalAtmosphericModelTupleResponseDataModeExercise  GlobalAtmosphericModelTupleResponseDataMode = "EXERCISE"
)

type GlobalAtmosphericModelGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalAtmosphericModelGetParams]'s query parameters as
// `url.Values`.
func (r GlobalAtmosphericModelGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalAtmosphericModelListParams struct {
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalAtmosphericModelListParams]'s query parameters as
// `url.Values`.
func (r GlobalAtmosphericModelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalAtmosphericModelCountParams struct {
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalAtmosphericModelCountParams]'s query parameters as
// `url.Values`.
func (r GlobalAtmosphericModelCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalAtmosphericModelGetFileParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalAtmosphericModelGetFileParams]'s query parameters as
// `url.Values`.
func (r GlobalAtmosphericModelGetFileParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalAtmosphericModelTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalAtmosphericModelTupleParams]'s query parameters as
// `url.Values`.
func (r GlobalAtmosphericModelTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalAtmosphericModelUnvalidatedPublishParams struct {
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
	DataMode GlobalAtmosphericModelUnvalidatedPublishParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Type of data associated with this record (e.g. Global Total Electron Density,
	// Global Total Electron Content).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Model execution cadence, in minutes.
	Cadence param.Opt[int64] `json:"cadence,omitzero"`
	// A unique identification code or label assigned to a particular source from which
	// atmospheric data originates.
	DataSourceIdentifier param.Opt[string] `json:"dataSourceIdentifier,omitzero"`
	// Ending altitude of model outputs, in kilometers.
	EndAlt param.Opt[float64] `json:"endAlt,omitzero"`
	// WGS-84 ending latitude of model output, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndLat param.Opt[float64] `json:"endLat,omitzero"`
	// WGS-84 ending longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndLon param.Opt[float64] `json:"endLon,omitzero"`
	// The file name of the uploaded file.
	Filename param.Opt[string] `json:"filename,omitzero"`
	// The uploaded file size, in bytes. The maximum file size for this service is
	// 104857600 bytes (100MB). Files exceeding the maximum size will be rejected.
	Filesize param.Opt[int64] `json:"filesize,omitzero"`
	// Number of altitude points.
	NumAlt param.Opt[int64] `json:"numAlt,omitzero"`
	// Number of latitude points.
	NumLat param.Opt[int64] `json:"numLat,omitzero"`
	// Number of longitude points.
	NumLon param.Opt[int64] `json:"numLon,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The time that this record was created, in ISO 8601 UTC format with millisecond
	// precision.
	ReportTime param.Opt[time.Time] `json:"reportTime,omitzero" format:"date-time"`
	// Starting altitude of model outputs, in kilometers.
	StartAlt param.Opt[float64] `json:"startAlt,omitzero"`
	// WGS-84 starting latitude of model output, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	StartLat param.Opt[float64] `json:"startLat,omitzero"`
	// WGS-84 starting longitude of model output, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	StartLon param.Opt[float64] `json:"startLon,omitzero"`
	// State value indicating whether the values in this record are PREDICTED or
	// OBSERVED.
	State param.Opt[string] `json:"state,omitzero"`
	// Separation in latitude between subsequent model outputs, in degrees.
	StepLat param.Opt[float64] `json:"stepLat,omitzero"`
	// Separation in longitude between subsequent model outputs, in degrees.
	StepLon param.Opt[float64] `json:"stepLon,omitzero"`
	paramObj
}

func (r GlobalAtmosphericModelUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	type shadow GlobalAtmosphericModelUnvalidatedPublishParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalAtmosphericModelUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
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
type GlobalAtmosphericModelUnvalidatedPublishParamsDataMode string

const (
	GlobalAtmosphericModelUnvalidatedPublishParamsDataModeReal      GlobalAtmosphericModelUnvalidatedPublishParamsDataMode = "REAL"
	GlobalAtmosphericModelUnvalidatedPublishParamsDataModeTest      GlobalAtmosphericModelUnvalidatedPublishParamsDataMode = "TEST"
	GlobalAtmosphericModelUnvalidatedPublishParamsDataModeSimulated GlobalAtmosphericModelUnvalidatedPublishParamsDataMode = "SIMULATED"
	GlobalAtmosphericModelUnvalidatedPublishParamsDataModeExercise  GlobalAtmosphericModelUnvalidatedPublishParamsDataMode = "EXERCISE"
)
