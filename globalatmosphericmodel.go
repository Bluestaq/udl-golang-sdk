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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// GlobalatmosphericmodelService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalatmosphericmodelService] method instead.
type GlobalatmosphericmodelService struct {
	Options []option.RequestOption
	History GlobalatmosphericmodelHistoryService
}

// NewGlobalatmosphericmodelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalatmosphericmodelService(opts ...option.RequestOption) (r GlobalatmosphericmodelService) {
	r = GlobalatmosphericmodelService{}
	r.Options = opts
	r.History = NewGlobalatmosphericmodelHistoryService(opts...)
	return
}

// Service operation to get a single GlobalAtmosphericModel record by its unique ID
// passed as a path parameter.
func (r *GlobalatmosphericmodelService) Get(ctx context.Context, id string, query GlobalatmosphericmodelGetParams, opts ...option.RequestOption) (res *GlobalatmosphericmodelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/globalatmosphericmodel/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *GlobalatmosphericmodelService) Count(ctx context.Context, query GlobalatmosphericmodelCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/globalatmosphericmodel/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single GlobalAtmosphericModel compressed data file by
// its unique ID passed as a path parameter. The compressed data file is returned
// as an attachment Content-Disposition.
func (r *GlobalatmosphericmodelService) GetFile(ctx context.Context, id string, query GlobalatmosphericmodelGetFileParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/globalatmosphericmodel/getFile/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *GlobalatmosphericmodelService) Query(ctx context.Context, query GlobalatmosphericmodelQueryParams, opts ...option.RequestOption) (res *[]GlobalatmosphericmodelQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/globalatmosphericmodel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *GlobalatmosphericmodelService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/globalatmosphericmodel/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
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
func (r *GlobalatmosphericmodelService) Tuple(ctx context.Context, query GlobalatmosphericmodelTupleParams, opts ...option.RequestOption) (res *[]GlobalatmosphericmodelTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *GlobalatmosphericmodelService) UnvalidatedPublish(ctx context.Context, body GlobalatmosphericmodelUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-globalatmosphericmodel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The GlobalAtmosphericModel service provides atmospheric model output data for
// use in space situational awareness such as the Global Total Electron Content
// (2D) data, Global Total Electron Density (3D) data, etc.
type GlobalatmosphericmodelGetResponse struct {
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
	DataMode GlobalatmosphericmodelGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Source                resp.Field
		Ts                    resp.Field
		Type                  resp.Field
		ID                    resp.Field
		Cadence               resp.Field
		ChecksumValue         resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EndAlt                resp.Field
		EndLat                resp.Field
		EndLon                resp.Field
		Filename              resp.Field
		Filesize              resp.Field
		NumAlt                resp.Field
		NumLat                resp.Field
		NumLon                resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		ReportTime            resp.Field
		SourceDl              resp.Field
		StartAlt              resp.Field
		StartLat              resp.Field
		StartLon              resp.Field
		State                 resp.Field
		StepLat               resp.Field
		StepLon               resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalatmosphericmodelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalatmosphericmodelGetResponse) UnmarshalJSON(data []byte) error {
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
type GlobalatmosphericmodelGetResponseDataMode string

const (
	GlobalatmosphericmodelGetResponseDataModeReal      GlobalatmosphericmodelGetResponseDataMode = "REAL"
	GlobalatmosphericmodelGetResponseDataModeTest      GlobalatmosphericmodelGetResponseDataMode = "TEST"
	GlobalatmosphericmodelGetResponseDataModeSimulated GlobalatmosphericmodelGetResponseDataMode = "SIMULATED"
	GlobalatmosphericmodelGetResponseDataModeExercise  GlobalatmosphericmodelGetResponseDataMode = "EXERCISE"
)

// The GlobalAtmosphericModel service provides atmospheric model output data for
// use in space situational awareness such as the Global Total Electron Content
// (2D) data, Global Total Electron Density (3D) data, etc.
type GlobalatmosphericmodelQueryResponse struct {
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
	DataMode GlobalatmosphericmodelQueryResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Source                resp.Field
		Ts                    resp.Field
		Type                  resp.Field
		ID                    resp.Field
		Cadence               resp.Field
		ChecksumValue         resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EndAlt                resp.Field
		EndLat                resp.Field
		EndLon                resp.Field
		Filename              resp.Field
		Filesize              resp.Field
		NumAlt                resp.Field
		NumLat                resp.Field
		NumLon                resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		ReportTime            resp.Field
		SourceDl              resp.Field
		StartAlt              resp.Field
		StartLat              resp.Field
		StartLon              resp.Field
		State                 resp.Field
		StepLat               resp.Field
		StepLon               resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalatmosphericmodelQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalatmosphericmodelQueryResponse) UnmarshalJSON(data []byte) error {
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
type GlobalatmosphericmodelQueryResponseDataMode string

const (
	GlobalatmosphericmodelQueryResponseDataModeReal      GlobalatmosphericmodelQueryResponseDataMode = "REAL"
	GlobalatmosphericmodelQueryResponseDataModeTest      GlobalatmosphericmodelQueryResponseDataMode = "TEST"
	GlobalatmosphericmodelQueryResponseDataModeSimulated GlobalatmosphericmodelQueryResponseDataMode = "SIMULATED"
	GlobalatmosphericmodelQueryResponseDataModeExercise  GlobalatmosphericmodelQueryResponseDataMode = "EXERCISE"
)

// The GlobalAtmosphericModel service provides atmospheric model output data for
// use in space situational awareness such as the Global Total Electron Content
// (2D) data, Global Total Electron Density (3D) data, etc.
type GlobalatmosphericmodelTupleResponse struct {
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
	DataMode GlobalatmosphericmodelTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Source                resp.Field
		Ts                    resp.Field
		Type                  resp.Field
		ID                    resp.Field
		Cadence               resp.Field
		ChecksumValue         resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EndAlt                resp.Field
		EndLat                resp.Field
		EndLon                resp.Field
		Filename              resp.Field
		Filesize              resp.Field
		NumAlt                resp.Field
		NumLat                resp.Field
		NumLon                resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		ReportTime            resp.Field
		SourceDl              resp.Field
		StartAlt              resp.Field
		StartLat              resp.Field
		StartLon              resp.Field
		State                 resp.Field
		StepLat               resp.Field
		StepLon               resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalatmosphericmodelTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalatmosphericmodelTupleResponse) UnmarshalJSON(data []byte) error {
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
type GlobalatmosphericmodelTupleResponseDataMode string

const (
	GlobalatmosphericmodelTupleResponseDataModeReal      GlobalatmosphericmodelTupleResponseDataMode = "REAL"
	GlobalatmosphericmodelTupleResponseDataModeTest      GlobalatmosphericmodelTupleResponseDataMode = "TEST"
	GlobalatmosphericmodelTupleResponseDataModeSimulated GlobalatmosphericmodelTupleResponseDataMode = "SIMULATED"
	GlobalatmosphericmodelTupleResponseDataModeExercise  GlobalatmosphericmodelTupleResponseDataMode = "EXERCISE"
)

type GlobalatmosphericmodelGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GlobalatmosphericmodelGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [GlobalatmosphericmodelGetParams]'s query parameters as
// `url.Values`.
func (r GlobalatmosphericmodelGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalatmosphericmodelCountParams struct {
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GlobalatmosphericmodelCountParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [GlobalatmosphericmodelCountParams]'s query parameters as
// `url.Values`.
func (r GlobalatmosphericmodelCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalatmosphericmodelGetFileParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GlobalatmosphericmodelGetFileParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [GlobalatmosphericmodelGetFileParams]'s query parameters as
// `url.Values`.
func (r GlobalatmosphericmodelGetFileParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalatmosphericmodelQueryParams struct {
	// Target time of the model in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GlobalatmosphericmodelQueryParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [GlobalatmosphericmodelQueryParams]'s query parameters as
// `url.Values`.
func (r GlobalatmosphericmodelQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalatmosphericmodelTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GlobalatmosphericmodelTupleParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [GlobalatmosphericmodelTupleParams]'s query parameters as
// `url.Values`.
func (r GlobalatmosphericmodelTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalatmosphericmodelUnvalidatedPublishParams struct {
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
	DataMode GlobalatmosphericmodelUnvalidatedPublishParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GlobalatmosphericmodelUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r GlobalatmosphericmodelUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	type shadow GlobalatmosphericmodelUnvalidatedPublishParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type GlobalatmosphericmodelUnvalidatedPublishParamsDataMode string

const (
	GlobalatmosphericmodelUnvalidatedPublishParamsDataModeReal      GlobalatmosphericmodelUnvalidatedPublishParamsDataMode = "REAL"
	GlobalatmosphericmodelUnvalidatedPublishParamsDataModeTest      GlobalatmosphericmodelUnvalidatedPublishParamsDataMode = "TEST"
	GlobalatmosphericmodelUnvalidatedPublishParamsDataModeSimulated GlobalatmosphericmodelUnvalidatedPublishParamsDataMode = "SIMULATED"
	GlobalatmosphericmodelUnvalidatedPublishParamsDataModeExercise  GlobalatmosphericmodelUnvalidatedPublishParamsDataMode = "EXERCISE"
)
