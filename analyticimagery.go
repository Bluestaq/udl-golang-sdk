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
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiform"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// AnalyticImageryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticImageryService] method instead.
type AnalyticImageryService struct {
	Options []option.RequestOption
}

// NewAnalyticImageryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticImageryService(opts ...option.RequestOption) (r AnalyticImageryService) {
	r = AnalyticImageryService{}
	r.Options = opts
	return
}

// Service operation to get a single AnalyticImagery record by its unique ID passed
// as a path parameter. AnalyticImagery represents metadata about an image, as well
// as the actual binary image data.
func (r *AnalyticImageryService) Get(ctx context.Context, id string, query AnalyticImageryGetParams, opts ...option.RequestOption) (res *AnalyticImageryFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/analyticimagery/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AnalyticImageryService) List(ctx context.Context, query AnalyticImageryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AnalyticImageryAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/analyticimagery"
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
func (r *AnalyticImageryService) ListAutoPaging(ctx context.Context, query AnalyticImageryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AnalyticImageryAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AnalyticImageryService) Count(ctx context.Context, query AnalyticImageryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/analyticimagery/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single AnalyticImagery binary image by its unique ID
// passed as a path parameter. The image is returned as an attachment
// Content-Disposition.
func (r *AnalyticImageryService) FileGet(ctx context.Context, id string, query AnalyticImageryFileGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/analyticimagery/getFile/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AnalyticImageryService) History(ctx context.Context, query AnalyticImageryHistoryParams, opts ...option.RequestOption) (res *[]AnalyticImageryFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/analyticimagery/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AnalyticImageryService) HistoryAodr(ctx context.Context, query AnalyticImageryHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/analyticimagery/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AnalyticImageryService) HistoryCount(ctx context.Context, query AnalyticImageryHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/analyticimagery/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AnalyticImageryService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/analyticimagery/queryhelp"
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
func (r *AnalyticImageryService) Tuple(ctx context.Context, query AnalyticImageryTupleParams, opts ...option.RequestOption) (res *[]AnalyticImageryFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/analyticimagery/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a new image with its metadata.
//
// The request body requires a zip file containing exactly two files:\
// 1\) A file with the `.json` file extension whose content conforms to the `AnalyticImagery_Ingest`
// schema.\
// 2\) A binary image file of the allowed types for this service.
//
// The JSON and image files will be associated with each other via the `id` field.
// Query the metadata via `GET /udl/analyticimagery` and use
// `GET /udl/analyticimagery/getFile/{id}` to retrieve the binary image file.
//
// This operation only accepts application/zip media. The application/json request
// body is documented to provide a convenient reference to the ingest schema.
//
// This operation is intended to be used for automated feeds into UDL. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *AnalyticImageryService) UnvalidatedPublish(ctx context.Context, body AnalyticImageryUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-analyticimagery"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The analytic imagery schema supports data plots and graphics of various types.
// The records contain general file information, allows for annotations to
// user-defined areas of interest on the graphics, and supports keyword searching.
type AnalyticImageryAbridged struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// General type of content of this image (CONTOUR, DIAGRAM, HEATMAP, HISTOGRAM,
	// PLOT, SCREENSHOT).
	Content string `json:"content,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode AnalyticImageryAbridgedDataMode `json:"dataMode,required"`
	// Description of the image content and utility.
	Description string `json:"description,required"`
	// The image file name.
	Filename string `json:"filename,required"`
	// The image file size, in bytes. The maximum file size for this service is
	// 40,000,000 bytes (40MB). Files exceeding the maximum size will be rejected.
	Filesize int64 `json:"filesize,required"`
	// The type of image associated with this record (GIF, JPG, PNG, TIF).
	ImageType string `json:"imageType,required"`
	// The message time of this image record, in ISO8601 UTC format with millisecond
	// precision.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Rectangular annotation limits, specified in pixels, as an array of arrays of the
	// coordinates [ [UL1x, UL1y], [UR1x, UR1y], [LR1x, LR1y], [LL1x, LL1y] ],
	// indicating the corners of a rectangle beginning with the Upper Left (UL) and
	// moving clockwise. Allows the image provider to highlight one or more rectangular
	// area(s) of interest. The array must contain Nx4 two-element arrays, where N is
	// the number of rectangles of interest. The associated annotation(s) should be
	// included in the annText array.
	AnnLims [][]int64 `json:"annLims"`
	// Annotation text, a string array of annotation(s) corresponding to the
	// rectangular areas specified in annLims. This array contains the annotation text
	// associated with the areas of interest indicated in annLims, in order. This array
	// should contain one annotation per four values of the area (annLims) array.
	AnnText []string `json:"annText"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground.
	Atype string `json:"atype"`
	// MD5 checksum value of the file. The ingest/create operation will automatically
	// generate the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The start time, in ISO8601 UTC format with millisecond precision, of the data
	// used in the analysis or composition of the image content, when applicable.
	DataStart time.Time `json:"dataStart" format:"date-time"`
	// The stop time, in ISO8601 UTC format with millisecond precision, of the data
	// used in the analysis or composition of the image content, when applicable.
	DataStop time.Time `json:"dataStop" format:"date-time"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// User-defined set ID of a sequence of images. Used to associate related analytic
	// image records.
	ImageSetID string `json:"imageSetId"`
	// The number of images in an image set.
	ImageSetLength int64 `json:"imageSetLength"`
	// The image height (vertical), in pixels.
	ImgHeight int64 `json:"imgHeight"`
	// The image width (horizontal), in pixels.
	ImgWidth int64 `json:"imgWidth"`
	// Array of searchable keywords for this analytic imagery record.
	Keywords []string `json:"keywords"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the source to indicate the sensor for this
	// collection. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Assessed satellite ID (NORAD RSO object number). The 'satId' and 'satIdConf'
	// arrays must match in size.
	SatID []string `json:"satId"`
	// Assessed satellite confidence corresponding to an assessment ID. Values are
	// between 0.0 and 1.0. The 'satId' and 'satIdConf' arrays must match in size.
	SatIDConf []float64 `json:"satIdConf"`
	// The sequence number of an image within an image set. If null, then it is assumed
	// that the order of images in an imageSet is not relevant.
	SequenceID int64 `json:"sequenceId"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Array of UUIDs of the UDL data records that are related to this image. See the
	// associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK) that
	// are related to this image. See the associated 'srcIds' array for the record
	// UUIDs, positionally corresponding to the record types in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Optional field indicating the units that apply to the x-axis of the attached
	// image, when applicable.
	XUnits string `json:"xUnits"`
	// Optional field indicating the units that apply to the y-axis of the attached
	// image, when applicable.
	YUnits string `json:"yUnits"`
	// Optional field indicating the units that apply to the z-axis of the attached
	// image, when applicable.
	ZUnits string `json:"zUnits"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Content               respjson.Field
		DataMode              respjson.Field
		Description           respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		ImageType             respjson.Field
		MsgTime               respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Agjson                respjson.Field
		Andims                respjson.Field
		AnnLims               respjson.Field
		AnnText               respjson.Field
		Asrid                 respjson.Field
		Atext                 respjson.Field
		Atype                 respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataStart             respjson.Field
		DataStop              respjson.Field
		IDSensor              respjson.Field
		ImageSetID            respjson.Field
		ImageSetLength        respjson.Field
		ImgHeight             respjson.Field
		ImgWidth              respjson.Field
		Keywords              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		SatID                 respjson.Field
		SatIDConf             respjson.Field
		SequenceID            respjson.Field
		SourceDl              respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		XUnits                respjson.Field
		YUnits                respjson.Field
		ZUnits                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticImageryAbridged) RawJSON() string { return r.JSON.raw }
func (r *AnalyticImageryAbridged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type AnalyticImageryAbridgedDataMode string

const (
	AnalyticImageryAbridgedDataModeReal      AnalyticImageryAbridgedDataMode = "REAL"
	AnalyticImageryAbridgedDataModeTest      AnalyticImageryAbridgedDataMode = "TEST"
	AnalyticImageryAbridgedDataModeSimulated AnalyticImageryAbridgedDataMode = "SIMULATED"
	AnalyticImageryAbridgedDataModeExercise  AnalyticImageryAbridgedDataMode = "EXERCISE"
)

// The analytic imagery schema supports data plots and graphics of various types.
// The records contain general file information, allows for annotations to
// user-defined areas of interest on the graphics, and supports keyword searching.
type AnalyticImageryFull struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// General type of content of this image (CONTOUR, DIAGRAM, HEATMAP, HISTOGRAM,
	// PLOT, SCREENSHOT).
	Content string `json:"content,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode AnalyticImageryFullDataMode `json:"dataMode,required"`
	// Description of the image content and utility.
	Description string `json:"description,required"`
	// The image file name.
	Filename string `json:"filename,required"`
	// The image file size, in bytes. The maximum file size for this service is
	// 40,000,000 bytes (40MB). Files exceeding the maximum size will be rejected.
	Filesize int64 `json:"filesize,required"`
	// The type of image associated with this record (GIF, JPG, PNG, TIF).
	ImageType string `json:"imageType,required"`
	// The message time of this image record, in ISO8601 UTC format with millisecond
	// precision.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Rectangular annotation limits, specified in pixels, as an array of arrays of the
	// coordinates [ [UL1x, UL1y], [UR1x, UR1y], [LR1x, LR1y], [LL1x, LL1y] ],
	// indicating the corners of a rectangle beginning with the Upper Left (UL) and
	// moving clockwise. Allows the image provider to highlight one or more rectangular
	// area(s) of interest. The array must contain Nx4 two-element arrays, where N is
	// the number of rectangles of interest. The associated annotation(s) should be
	// included in the annText array.
	AnnLims [][]int64 `json:"annLims"`
	// Annotation text, a string array of annotation(s) corresponding to the
	// rectangular areas specified in annLims. This array contains the annotation text
	// associated with the areas of interest indicated in annLims, in order. This array
	// should contain one annotation per four values of the area (annLims) array.
	AnnText []string `json:"annText"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the point of interest as projected on the ground.
	Area string `json:"area"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground.
	Atype string `json:"atype"`
	// MD5 checksum value of the file. The ingest/create operation will automatically
	// generate the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The start time, in ISO8601 UTC format with millisecond precision, of the data
	// used in the analysis or composition of the image content, when applicable.
	DataStart time.Time `json:"dataStart" format:"date-time"`
	// The stop time, in ISO8601 UTC format with millisecond precision, of the data
	// used in the analysis or composition of the image content, when applicable.
	DataStop time.Time `json:"dataStop" format:"date-time"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// User-defined set ID of a sequence of images. Used to associate related analytic
	// image records.
	ImageSetID string `json:"imageSetId"`
	// The number of images in an image set.
	ImageSetLength int64 `json:"imageSetLength"`
	// The image height (vertical), in pixels.
	ImgHeight int64 `json:"imgHeight"`
	// The image width (horizontal), in pixels.
	ImgWidth int64 `json:"imgWidth"`
	// Array of searchable keywords for this analytic imagery record.
	Keywords []string `json:"keywords"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the source to indicate the sensor for this
	// collection. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Assessed satellite ID (NORAD RSO object number). The 'satId' and 'satIdConf'
	// arrays must match in size.
	SatID []string `json:"satId"`
	// Assessed satellite confidence corresponding to an assessment ID. Values are
	// between 0.0 and 1.0. The 'satId' and 'satIdConf' arrays must match in size.
	SatIDConf []float64 `json:"satIdConf"`
	// The sequence number of an image within an image set. If null, then it is assumed
	// that the order of images in an imageSet is not relevant.
	SequenceID int64 `json:"sequenceId"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Array of UUIDs of the UDL data records that are related to this image. See the
	// associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK) that
	// are related to this image. See the associated 'srcIds' array for the record
	// UUIDs, positionally corresponding to the record types in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Optional field indicating the units that apply to the x-axis of the attached
	// image, when applicable.
	XUnits string `json:"xUnits"`
	// Optional field indicating the units that apply to the y-axis of the attached
	// image, when applicable.
	YUnits string `json:"yUnits"`
	// Optional field indicating the units that apply to the z-axis of the attached
	// image, when applicable.
	ZUnits string `json:"zUnits"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Content               respjson.Field
		DataMode              respjson.Field
		Description           respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		ImageType             respjson.Field
		MsgTime               respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Agjson                respjson.Field
		Andims                respjson.Field
		AnnLims               respjson.Field
		AnnText               respjson.Field
		Area                  respjson.Field
		Asrid                 respjson.Field
		Atext                 respjson.Field
		Atype                 respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataStart             respjson.Field
		DataStop              respjson.Field
		IDSensor              respjson.Field
		ImageSetID            respjson.Field
		ImageSetLength        respjson.Field
		ImgHeight             respjson.Field
		ImgWidth              respjson.Field
		Keywords              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		SatID                 respjson.Field
		SatIDConf             respjson.Field
		SequenceID            respjson.Field
		SourceDl              respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		XUnits                respjson.Field
		YUnits                respjson.Field
		ZUnits                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticImageryFull) RawJSON() string { return r.JSON.raw }
func (r *AnalyticImageryFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type AnalyticImageryFullDataMode string

const (
	AnalyticImageryFullDataModeReal      AnalyticImageryFullDataMode = "REAL"
	AnalyticImageryFullDataModeTest      AnalyticImageryFullDataMode = "TEST"
	AnalyticImageryFullDataModeSimulated AnalyticImageryFullDataMode = "SIMULATED"
	AnalyticImageryFullDataModeExercise  AnalyticImageryFullDataMode = "EXERCISE"
)

type AnalyticImageryGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticImageryGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticImageryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticImageryListParams struct {
	// The message time of this image record, in ISO8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticImageryListParams]'s query parameters as
// `url.Values`.
func (r AnalyticImageryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticImageryCountParams struct {
	// The message time of this image record, in ISO8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticImageryCountParams]'s query parameters as
// `url.Values`.
func (r AnalyticImageryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticImageryFileGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticImageryFileGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticImageryFileGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticImageryHistoryParams struct {
	// The message time of this image record, in ISO8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime time.Time `query:"msgTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticImageryHistoryParams]'s query parameters as
// `url.Values`.
func (r AnalyticImageryHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticImageryHistoryAodrParams struct {
	// The message time of this image record, in ISO8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime time.Time `query:"msgTime,required" format:"date-time" json:"-"`
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

// URLQuery serializes [AnalyticImageryHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r AnalyticImageryHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticImageryHistoryCountParams struct {
	// The message time of this image record, in ISO8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticImageryHistoryCountParams]'s query parameters as
// `url.Values`.
func (r AnalyticImageryHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticImageryTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The message time of this image record, in ISO8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticImageryTupleParams]'s query parameters as
// `url.Values`.
func (r AnalyticImageryTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticImageryUnvalidatedPublishParams struct {
	// Zip file containing files described in the specification
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

func (r AnalyticImageryUnvalidatedPublishParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
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
