// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apiform"
	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// GroundImageryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroundImageryService] method instead.
type GroundImageryService struct {
	Options []option.RequestOption
	History GroundImageryHistoryService
}

// NewGroundImageryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGroundImageryService(opts ...option.RequestOption) (r GroundImageryService) {
	r = GroundImageryService{}
	r.Options = opts
	r.History = NewGroundImageryHistoryService(opts...)
	return
}

// Service operation to take a single GroundImagery object as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *GroundImageryService) New(ctx context.Context, body GroundImageryNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/groundimagery"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *GroundImageryService) List(ctx context.Context, query GroundImageryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GroundImageryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/groundimagery"
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
func (r *GroundImageryService) ListAutoPaging(ctx context.Context, query GroundImageryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GroundImageryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *GroundImageryService) Aodr(ctx context.Context, query GroundImageryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/groundimagery/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *GroundImageryService) Count(ctx context.Context, query GroundImageryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/groundimagery/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single GroundImagery record by its unique ID passed
// as a path parameter. GroundImagery represents metadata about a ground image, as
// well as the actual binary image data.
func (r *GroundImageryService) Get(ctx context.Context, id string, query GroundImageryGetParams, opts ...option.RequestOption) (res *GroundImageryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/groundimagery/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single GroundImagery binary image by its unique ID
// passed as a path parameter. The image is returned as an attachment
// Content-Disposition.
func (r *GroundImageryService) GetFile(ctx context.Context, id string, query GroundImageryGetFileParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/groundimagery/getFile/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *GroundImageryService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *GroundImageryQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/groundimagery/queryhelp"
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
func (r *GroundImageryService) Tuple(ctx context.Context, query GroundImageryTupleParams, opts ...option.RequestOption) (res *[]GroundImageryTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/groundimagery/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a new image with its metadata.
//
// The request body requires a zip file containing exactly two files:\
// 1\) A file with the `.json` file extension whose content conforms to the `GroundImagery_Ingest`
// schema. 2\) A binary image file of the allowed types for this service.
//
// The JSON and image files will be associated with each other via the `id` field.
// Query the metadata via `GET /udl/groundimagery` and use
// `GET /udl/groundimagery/getFile/{id}` to retrieve the binary image file.
//
// This operation only accepts application/zip media. The application/json request
// body is documented to provide a convenient reference to the ingest schema.
//
// This operation is intended to be used for automated feeds into UDL. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *GroundImageryService) UploadZip(ctx context.Context, body GroundImageryUploadZipParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://imagery.unifieddatalibrary.com/")}, opts...)
	path := "filedrop/udl-groundimagery"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Imagery of terrestrial regions from on-orbit, air, and other sensors.
type GroundImageryListResponse struct {
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
	DataMode GroundImageryListResponseDataMode `json:"dataMode,required"`
	// Name of the image file.
	Filename string `json:"filename,required"`
	// Timestamp the image was captured/produced.
	ImageTime time.Time `json:"imageTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Size of the image file. Units in bytes. If filesize is provided without an
	// associated file, it defaults to 0.
	Filesize int64 `json:"filesize"`
	// Optional, field indicating type of image, NITF, PNG, etc.
	Format string `json:"format"`
	// Optional ID of the sensor that produced this ground image.
	IDSensor string `json:"idSensor"`
	// Optional name/description associated with this image.
	Name string `json:"name"`
	// Description and notes of the image.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. Reference: https://geojson.org/. Ignored if included with a create
	// operation that also specifies a valid region or regionText.
	RegionGeoJson string `json:"regionGeoJSON"`
	// Number of dimensions of the geometry depicted by region.
	RegionNDims int64 `json:"regionNDims"`
	// Geographical spatial_ref_sys for region.
	RegionSRid int64 `json:"regionSRid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a create operation that also specifies a valid region.
	RegionText string `json:"regionText"`
	// Type of region as projected on the ground.
	RegionType string `json:"regionType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional identifier of the subject/target of the image, useful for correlating
	// multiple images of the same subject.
	SubjectID string `json:"subjectId"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Filename              respjson.Field
		ImageTime             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Filesize              respjson.Field
		Format                respjson.Field
		IDSensor              respjson.Field
		Name                  respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		RegionGeoJson         respjson.Field
		RegionNDims           respjson.Field
		RegionSRid            respjson.Field
		RegionText            respjson.Field
		RegionType            respjson.Field
		SourceDl              respjson.Field
		SubjectID             respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroundImageryListResponse) RawJSON() string { return r.JSON.raw }
func (r *GroundImageryListResponse) UnmarshalJSON(data []byte) error {
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
type GroundImageryListResponseDataMode string

const (
	GroundImageryListResponseDataModeReal      GroundImageryListResponseDataMode = "REAL"
	GroundImageryListResponseDataModeTest      GroundImageryListResponseDataMode = "TEST"
	GroundImageryListResponseDataModeSimulated GroundImageryListResponseDataMode = "SIMULATED"
	GroundImageryListResponseDataModeExercise  GroundImageryListResponseDataMode = "EXERCISE"
)

// Imagery of terrestrial regions from on-orbit, air, and other sensors.
type GroundImageryGetResponse struct {
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
	DataMode GroundImageryGetResponseDataMode `json:"dataMode,required"`
	// Name of the image file.
	Filename string `json:"filename,required"`
	// Timestamp the image was captured/produced.
	ImageTime time.Time `json:"imageTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Size of the image file. Units in bytes. If filesize is provided without an
	// associated file, it defaults to 0.
	Filesize int64 `json:"filesize"`
	// Optional, field indicating type of image, NITF, PNG, etc.
	Format string `json:"format"`
	// Optional ID of the sensor that produced this ground image.
	IDSensor string `json:"idSensor"`
	// Optional array of keywords for this image.
	Keywords []string `json:"keywords"`
	// Optional name/description associated with this image.
	Name string `json:"name"`
	// Description and notes of the image.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Geographical region or polygon (lon/lat pairs) of the image as projected on the
	// ground in geoJSON or geoText format. This is an optional convenience field only
	// used for create operations. The system will auto-detect the format (Well Known
	// Text or GeoJSON) and populate both regionText and regionGeoJSON fields
	// appropriately. When omitted, regionText or regionGeoJSON is expected.
	Region string `json:"region"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. Reference: https://geojson.org/. Ignored if included with a create
	// operation that also specifies a valid region or regionText.
	RegionGeoJson string `json:"regionGeoJSON"`
	// Number of dimensions of the geometry depicted by region.
	RegionNDims int64 `json:"regionNDims"`
	// Geographical spatial_ref_sys for region.
	RegionSRid int64 `json:"regionSRid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a create operation that also specifies a valid region.
	RegionText string `json:"regionText"`
	// Type of region as projected on the ground.
	RegionType string `json:"regionType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional identifier of the subject/target of the image, useful for correlating
	// multiple images of the same subject.
	SubjectID string `json:"subjectId"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Filename              respjson.Field
		ImageTime             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Filesize              respjson.Field
		Format                respjson.Field
		IDSensor              respjson.Field
		Keywords              respjson.Field
		Name                  respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		Region                respjson.Field
		RegionGeoJson         respjson.Field
		RegionNDims           respjson.Field
		RegionSRid            respjson.Field
		RegionText            respjson.Field
		RegionType            respjson.Field
		SourceDl              respjson.Field
		SubjectID             respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroundImageryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GroundImageryGetResponse) UnmarshalJSON(data []byte) error {
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
type GroundImageryGetResponseDataMode string

const (
	GroundImageryGetResponseDataModeReal      GroundImageryGetResponseDataMode = "REAL"
	GroundImageryGetResponseDataModeTest      GroundImageryGetResponseDataMode = "TEST"
	GroundImageryGetResponseDataModeSimulated GroundImageryGetResponseDataMode = "SIMULATED"
	GroundImageryGetResponseDataModeExercise  GroundImageryGetResponseDataMode = "EXERCISE"
)

type GroundImageryQueryhelpResponse struct {
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
func (r GroundImageryQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *GroundImageryQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Imagery of terrestrial regions from on-orbit, air, and other sensors.
type GroundImageryTupleResponse struct {
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
	DataMode GroundImageryTupleResponseDataMode `json:"dataMode,required"`
	// Name of the image file.
	Filename string `json:"filename,required"`
	// Timestamp the image was captured/produced.
	ImageTime time.Time `json:"imageTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Size of the image file. Units in bytes. If filesize is provided without an
	// associated file, it defaults to 0.
	Filesize int64 `json:"filesize"`
	// Optional, field indicating type of image, NITF, PNG, etc.
	Format string `json:"format"`
	// Optional ID of the sensor that produced this ground image.
	IDSensor string `json:"idSensor"`
	// Optional array of keywords for this image.
	Keywords []string `json:"keywords"`
	// Optional name/description associated with this image.
	Name string `json:"name"`
	// Description and notes of the image.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Geographical region or polygon (lon/lat pairs) of the image as projected on the
	// ground in geoJSON or geoText format. This is an optional convenience field only
	// used for create operations. The system will auto-detect the format (Well Known
	// Text or GeoJSON) and populate both regionText and regionGeoJSON fields
	// appropriately. When omitted, regionText or regionGeoJSON is expected.
	Region string `json:"region"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. Reference: https://geojson.org/. Ignored if included with a create
	// operation that also specifies a valid region or regionText.
	RegionGeoJson string `json:"regionGeoJSON"`
	// Number of dimensions of the geometry depicted by region.
	RegionNDims int64 `json:"regionNDims"`
	// Geographical spatial_ref_sys for region.
	RegionSRid int64 `json:"regionSRid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a create operation that also specifies a valid region.
	RegionText string `json:"regionText"`
	// Type of region as projected on the ground.
	RegionType string `json:"regionType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional identifier of the subject/target of the image, useful for correlating
	// multiple images of the same subject.
	SubjectID string `json:"subjectId"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Filename              respjson.Field
		ImageTime             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Filesize              respjson.Field
		Format                respjson.Field
		IDSensor              respjson.Field
		Keywords              respjson.Field
		Name                  respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		Region                respjson.Field
		RegionGeoJson         respjson.Field
		RegionNDims           respjson.Field
		RegionSRid            respjson.Field
		RegionText            respjson.Field
		RegionType            respjson.Field
		SourceDl              respjson.Field
		SubjectID             respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroundImageryTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *GroundImageryTupleResponse) UnmarshalJSON(data []byte) error {
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
type GroundImageryTupleResponseDataMode string

const (
	GroundImageryTupleResponseDataModeReal      GroundImageryTupleResponseDataMode = "REAL"
	GroundImageryTupleResponseDataModeTest      GroundImageryTupleResponseDataMode = "TEST"
	GroundImageryTupleResponseDataModeSimulated GroundImageryTupleResponseDataMode = "SIMULATED"
	GroundImageryTupleResponseDataModeExercise  GroundImageryTupleResponseDataMode = "EXERCISE"
)

type GroundImageryNewParams struct {
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
	DataMode GroundImageryNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Name of the image file.
	Filename string `json:"filename,required"`
	// Timestamp the image was captured/produced.
	ImageTime time.Time `json:"imageTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue param.Opt[string] `json:"checksumValue,omitzero"`
	// Size of the image file. Units in bytes. If filesize is provided without an
	// associated file, it defaults to 0.
	Filesize param.Opt[int64] `json:"filesize,omitzero"`
	// Optional, field indicating type of image, NITF, PNG, etc.
	Format param.Opt[string] `json:"format,omitzero"`
	// Optional ID of the sensor that produced this ground image.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Optional name/description associated with this image.
	Name param.Opt[string] `json:"name,omitzero"`
	// Description and notes of the image.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Geographical region or polygon (lon/lat pairs) of the image as projected on the
	// ground in geoJSON or geoText format. This is an optional convenience field only
	// used for create operations. The system will auto-detect the format (Well Known
	// Text or GeoJSON) and populate both regionText and regionGeoJSON fields
	// appropriately. When omitted, regionText or regionGeoJSON is expected.
	Region param.Opt[string] `json:"region,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. Reference: https://geojson.org/. Ignored if included with a create
	// operation that also specifies a valid region or regionText.
	RegionGeoJson param.Opt[string] `json:"regionGeoJSON,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	RegionNDims param.Opt[int64] `json:"regionNDims,omitzero"`
	// Geographical spatial_ref_sys for region.
	RegionSRid param.Opt[int64] `json:"regionSRid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a create operation that also specifies a valid region.
	RegionText param.Opt[string] `json:"regionText,omitzero"`
	// Type of region as projected on the ground.
	RegionType param.Opt[string] `json:"regionType,omitzero"`
	// Optional identifier of the subject/target of the image, useful for correlating
	// multiple images of the same subject.
	SubjectID param.Opt[string] `json:"subjectId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Optional array of keywords for this image.
	Keywords []string `json:"keywords,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r GroundImageryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GroundImageryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroundImageryNewParams) UnmarshalJSON(data []byte) error {
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
type GroundImageryNewParamsDataMode string

const (
	GroundImageryNewParamsDataModeReal      GroundImageryNewParamsDataMode = "REAL"
	GroundImageryNewParamsDataModeTest      GroundImageryNewParamsDataMode = "TEST"
	GroundImageryNewParamsDataModeSimulated GroundImageryNewParamsDataMode = "SIMULATED"
	GroundImageryNewParamsDataModeExercise  GroundImageryNewParamsDataMode = "EXERCISE"
)

type GroundImageryListParams struct {
	// Timestamp the image was captured/produced. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ImageTime   time.Time        `query:"imageTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroundImageryListParams]'s query parameters as
// `url.Values`.
func (r GroundImageryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroundImageryAodrParams struct {
	// Timestamp the image was captured/produced. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ImageTime time.Time `query:"imageTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// optional, notification method for the created file link. When omitted, EMAIL is
	// assumed. Current valid values are: EMAIL, SMS.
	Notification param.Opt[string] `query:"notification,omitzero" json:"-"`
	// optional, field delimiter when the created file is not JSON. Must be a single
	// character chosen from this set: (',', ';', ':', '|'). When omitted, "," is used.
	// It is strongly encouraged that your field delimiter be a character unlikely to
	// occur within the data.
	OutputDelimiter param.Opt[string] `query:"outputDelimiter,omitzero" json:"-"`
	// optional, output format for the file. When omitted, JSON is assumed. Current
	// valid values are: JSON and CSV.
	OutputFormat param.Opt[string] `query:"outputFormat,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroundImageryAodrParams]'s query parameters as
// `url.Values`.
func (r GroundImageryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroundImageryCountParams struct {
	// Timestamp the image was captured/produced. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ImageTime   time.Time        `query:"imageTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroundImageryCountParams]'s query parameters as
// `url.Values`.
func (r GroundImageryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroundImageryGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroundImageryGetParams]'s query parameters as `url.Values`.
func (r GroundImageryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroundImageryGetFileParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroundImageryGetFileParams]'s query parameters as
// `url.Values`.
func (r GroundImageryGetFileParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroundImageryTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Timestamp the image was captured/produced. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ImageTime   time.Time        `query:"imageTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroundImageryTupleParams]'s query parameters as
// `url.Values`.
func (r GroundImageryTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroundImageryUploadZipParams struct {
	// Zip file containing files described in the specification
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

func (r GroundImageryUploadZipParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
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
