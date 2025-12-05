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

// SkyImageryService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkyImageryService] method instead.
type SkyImageryService struct {
	Options []option.RequestOption
	History SkyImageryHistoryService
}

// NewSkyImageryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSkyImageryService(opts ...option.RequestOption) (r SkyImageryService) {
	r = SkyImageryService{}
	r.Options = opts
	r.History = NewSkyImageryHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SkyImageryService) List(ctx context.Context, query SkyImageryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SkyImageryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/skyimagery"
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
func (r *SkyImageryService) ListAutoPaging(ctx context.Context, query SkyImageryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SkyImageryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SkyImageryService) Count(ctx context.Context, query SkyImageryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/skyimagery/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SkyImagery binary image by its unique ID
// passed as a path parameter. The image is returned as an attachment
// Content-Disposition.
func (r *SkyImageryService) FileGet(ctx context.Context, id string, query SkyImageryFileGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/skyimagery/getFile/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SkyImagery record by its unique ID passed as a
// path parameter. SkyImagery represents metadata about a sky image, as well as the
// actual binary image data.
func (r *SkyImageryService) Get(ctx context.Context, id string, query SkyImageryGetParams, opts ...option.RequestOption) (res *SkyImageryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/skyimagery/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SkyImageryService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SkyImageryQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/skyimagery/queryhelp"
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
func (r *SkyImageryService) Tuple(ctx context.Context, query SkyImageryTupleParams, opts ...option.RequestOption) (res *[]SkyImageryTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/skyimagery/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a new image with its metadata.
//
// The request body requires a zip file containing exactly two files:\
// 1\) A file with the `.json` file extension whose content conforms to the `SkyImagery_Ingest`
// schema.\
// 2\) A binary image file of the allowed types for this service.
//
// The JSON and image files will be associated with each other via the `id` field.
// Query the metadata via `GET /udl/skyimagery` and use
// `GET /udl/skyimagery/getFile/{id}` to retrieve the binary image file.
//
// This operation only accepts application/zip media. The application/json request
// body is documented to provide a convenient reference to the ingest schema.
//
// This operation is intended to be used for automated feeds into UDL. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *SkyImageryService) UploadZip(ctx context.Context, body SkyImageryUploadZipParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://imagery.unifieddatalibrary.com/")}, opts...)
	path := "filedrop/udl-skyimagery"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Model representation of sky imagery data. Sky imagery is ground or space based
// telescope imagery of RSO's and includes metadata on the image (time, source,
// etc) as well as binary image content (e.g. FITS, EOSSA, EOCHIP, MP4). Binary
// content must be downloaded individually by ID using the 'getFile' operation.
type SkyImageryListResponse struct {
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
	DataMode SkyImageryListResponseDataMode `json:"dataMode,required"`
	// Start time of the exposure, in ISO 8601 UTC format with microsecond precision.
	ExpStartTime time.Time `json:"expStartTime,required" format:"date-time"`
	// The type of image associated with this record (e.g. FITS, EOSSA, EOCHIP, MP4).
	ImageType string `json:"imageType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Reference to an annotation document associated with this image.
	AnnotationKey string `json:"annotationKey"`
	// Reference to a calibration document associated with this image.
	CalibrationKey string `json:"calibrationKey"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional name/description associated with this image.
	Description string `json:"description"`
	// End time of the exposure, in ISO 8601 UTC format with microsecond precision.
	ExpEndTime time.Time `json:"expEndTime" format:"date-time"`
	// Name of the uploaded image file.
	Filename string `json:"filename"`
	// Size of the image file, in bytes.
	Filesize int64 `json:"filesize"`
	// Field Of View frame height, in degrees.
	FrameFovHeight float64 `json:"frameFOVHeight"`
	// Field Of View frame width, in degrees.
	FrameFovWidth float64 `json:"frameFOVWidth"`
	// Frame height of the image, in number of pixels.
	FrameHeightPixels int64 `json:"frameHeightPixels"`
	// Frame width of the image, in number of pixels.
	FrameWidthPixels int64 `json:"frameWidthPixels"`
	// Optional identifier of the AttitudeSet data record describing the orientation of
	// an object body.
	IDAttitudeSet string `json:"idAttitudeSet"`
	// Optional identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// Optional unique identifier of the SOI Observation Set associated with this
	// image.
	IDSoiSet string `json:"idSOISet"`
	// The user-defined set ID of a sequence of images.
	ImageSetID string `json:"imageSetId"`
	// The number of images in an image set.
	ImageSetLength int64 `json:"imageSetLength"`
	// String that uniquely identifies the data source.
	ImageSourceInfo string `json:"imageSourceInfo"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the imaging source to indicate the target
	// onorbit object of this image. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the imaging source to indicate the sensor
	// identifier which produced this image. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Bit depth of the image, in number of pixels.
	PixelBitDepth int64 `json:"pixelBitDepth"`
	// Field Of View pixel height, in degrees.
	PixelFovHeight float64 `json:"pixelFOVHeight"`
	// Field Of View pixel width, in degrees.
	PixelFovWidth float64 `json:"pixelFOVWidth"`
	// Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sensor altitude at exposure start epoch, in km. If null, can be obtained from
	// sensor info.
	Senalt float64 `json:"senalt"`
	// Sensor latitude at exposure start epoch, in degrees. If null, can be obtained
	// from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor longitude at exposure start epoch, in degrees. If null, can be obtained
	// from sensor info. -180 to 180 degrees (negative values west of Prime Meridian).
	Senlon float64 `json:"senlon"`
	// The quaternion describing the rotation of the body-fixed frame used for this
	// system into the local geodetic frame, at exposure start epoch (expStartTime).
	// The array element order convention is scalar component first, followed by the
	// three vector components. For a vector u in the body-fixed frame, the
	// corresponding vector u' in the geodetic frame should satisfy u' = quq\*, where q
	// is this quaternion.
	SenQuat []float64 `json:"senQuat"`
	// The derivative of the quaternion describing the rotation of the body-fixed frame
	// used for this system into the local geodetic frame, exposure start epoch
	// (expStartTime). The array element order convention is scalar component first,
	// followed by the three vector components. For a vector u in the body-fixed frame,
	// the corresponding vector u' in the geodetic frame should satisfy u' = quq\*,
	// where q is this quaternion.
	SenQuatDot []float64 `json:"senQuatDot"`
	// Sensor x position at exposure start epoch, in km (if mobile/onorbit) in J2000
	// coordinate frame.
	Senx float64 `json:"senx"`
	// Sensor y position at exposure start epoch, in km (if mobile/onorbit) in J2000
	// coordinate frame.
	Seny float64 `json:"seny"`
	// Sensor z position at exposure start epoch, in km (if mobile/onorbit) in J2000
	// coordinate frame.
	Senz float64 `json:"senz"`
	// The sequence ID of an image within an image set.
	SequenceID int64 `json:"sequenceId"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// The telescope pointing azimuth, in degrees, at the exposure start epoch.
	TopLeftStartAz float64 `json:"topLeftStartAz"`
	// The telescope pointing elevation, in degrees, at the exposure start epoch.
	TopLeftStartEl float64 `json:"topLeftStartEl"`
	// The telescope pointing azimuth, in degrees, at the exposure stop epoch.
	TopLeftStopAz float64 `json:"topLeftStopAz"`
	// The telescope pointing elevation, in degrees, at the exposure stop epoch.
	TopLeftStopEl float64 `json:"topLeftStopEl"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ExpStartTime          respjson.Field
		ImageType             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AnnotationKey         respjson.Field
		CalibrationKey        respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Description           respjson.Field
		ExpEndTime            respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		FrameFovHeight        respjson.Field
		FrameFovWidth         respjson.Field
		FrameHeightPixels     respjson.Field
		FrameWidthPixels      respjson.Field
		IDAttitudeSet         respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		IDSoiSet              respjson.Field
		ImageSetID            respjson.Field
		ImageSetLength        respjson.Field
		ImageSourceInfo       respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		PixelBitDepth         respjson.Field
		PixelFovHeight        respjson.Field
		PixelFovWidth         respjson.Field
		SatNo                 respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		SenQuat               respjson.Field
		SenQuatDot            respjson.Field
		Senx                  respjson.Field
		Seny                  respjson.Field
		Senz                  respjson.Field
		SequenceID            respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		TopLeftStartAz        respjson.Field
		TopLeftStartEl        respjson.Field
		TopLeftStopAz         respjson.Field
		TopLeftStopEl         respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkyImageryListResponse) RawJSON() string { return r.JSON.raw }
func (r *SkyImageryListResponse) UnmarshalJSON(data []byte) error {
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
type SkyImageryListResponseDataMode string

const (
	SkyImageryListResponseDataModeReal      SkyImageryListResponseDataMode = "REAL"
	SkyImageryListResponseDataModeTest      SkyImageryListResponseDataMode = "TEST"
	SkyImageryListResponseDataModeSimulated SkyImageryListResponseDataMode = "SIMULATED"
	SkyImageryListResponseDataModeExercise  SkyImageryListResponseDataMode = "EXERCISE"
)

// Model representation of sky imagery data. Sky imagery is ground or space based
// telescope imagery of RSO's and includes metadata on the image (time, source,
// etc) as well as binary image content (e.g. FITS, EOSSA, EOCHIP, MP4). Binary
// content must be downloaded individually by ID using the 'getFile' operation.
type SkyImageryGetResponse struct {
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
	DataMode SkyImageryGetResponseDataMode `json:"dataMode,required"`
	// Start time of the exposure, in ISO 8601 UTC format with microsecond precision.
	ExpStartTime time.Time `json:"expStartTime,required" format:"date-time"`
	// The type of image associated with this record (e.g. FITS, EOSSA, EOCHIP, MP4).
	ImageType string `json:"imageType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Reference to an annotation document associated with this image.
	AnnotationKey string `json:"annotationKey"`
	// Reference to a calibration document associated with this image.
	CalibrationKey string `json:"calibrationKey"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional name/description associated with this image.
	Description string `json:"description"`
	// Collection of linked EOObservations.
	EoObservations []shared.EoObservationFull `json:"eoObservations"`
	// End time of the exposure, in ISO 8601 UTC format with microsecond precision.
	ExpEndTime time.Time `json:"expEndTime" format:"date-time"`
	// Name of the uploaded image file.
	Filename string `json:"filename"`
	// Size of the image file, in bytes.
	Filesize int64 `json:"filesize"`
	// Field Of View frame height, in degrees.
	FrameFovHeight float64 `json:"frameFOVHeight"`
	// Field Of View frame width, in degrees.
	FrameFovWidth float64 `json:"frameFOVWidth"`
	// Frame height of the image, in number of pixels.
	FrameHeightPixels int64 `json:"frameHeightPixels"`
	// Frame width of the image, in number of pixels.
	FrameWidthPixels int64 `json:"frameWidthPixels"`
	// Optional identifier of the AttitudeSet data record describing the orientation of
	// an object body.
	IDAttitudeSet string `json:"idAttitudeSet"`
	// Optional identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// Optional unique identifier of the SOI Observation Set associated with this
	// image.
	IDSoiSet string `json:"idSOISet"`
	// The user-defined set ID of a sequence of images.
	ImageSetID string `json:"imageSetId"`
	// The number of images in an image set.
	ImageSetLength int64 `json:"imageSetLength"`
	// String that uniquely identifies the data source.
	ImageSourceInfo string `json:"imageSourceInfo"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the imaging source to indicate the target
	// onorbit object of this image. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the imaging source to indicate the sensor
	// identifier which produced this image. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Bit depth of the image, in number of pixels.
	PixelBitDepth int64 `json:"pixelBitDepth"`
	// Field Of View pixel height, in degrees.
	PixelFovHeight float64 `json:"pixelFOVHeight"`
	// Field Of View pixel width, in degrees.
	PixelFovWidth float64 `json:"pixelFOVWidth"`
	// Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sensor altitude at exposure start epoch, in km. If null, can be obtained from
	// sensor info.
	Senalt float64 `json:"senalt"`
	// Sensor latitude at exposure start epoch, in degrees. If null, can be obtained
	// from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor longitude at exposure start epoch, in degrees. If null, can be obtained
	// from sensor info. -180 to 180 degrees (negative values west of Prime Meridian).
	Senlon float64 `json:"senlon"`
	// The quaternion describing the rotation of the body-fixed frame used for this
	// system into the local geodetic frame, at exposure start epoch (expStartTime).
	// The array element order convention is scalar component first, followed by the
	// three vector components. For a vector u in the body-fixed frame, the
	// corresponding vector u' in the geodetic frame should satisfy u' = quq\*, where q
	// is this quaternion.
	SenQuat []float64 `json:"senQuat"`
	// The derivative of the quaternion describing the rotation of the body-fixed frame
	// used for this system into the local geodetic frame, exposure start epoch
	// (expStartTime). The array element order convention is scalar component first,
	// followed by the three vector components. For a vector u in the body-fixed frame,
	// the corresponding vector u' in the geodetic frame should satisfy u' = quq\*,
	// where q is this quaternion.
	SenQuatDot []float64 `json:"senQuatDot"`
	// Sensor x position at exposure start epoch, in km (if mobile/onorbit) in J2000
	// coordinate frame.
	Senx float64 `json:"senx"`
	// Sensor y position at exposure start epoch, in km (if mobile/onorbit) in J2000
	// coordinate frame.
	Seny float64 `json:"seny"`
	// Sensor z position at exposure start epoch, in km (if mobile/onorbit) in J2000
	// coordinate frame.
	Senz float64 `json:"senz"`
	// The sequence ID of an image within an image set.
	SequenceID int64 `json:"sequenceId"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// The telescope pointing azimuth, in degrees, at the exposure start epoch.
	TopLeftStartAz float64 `json:"topLeftStartAz"`
	// The telescope pointing elevation, in degrees, at the exposure start epoch.
	TopLeftStartEl float64 `json:"topLeftStartEl"`
	// The telescope pointing azimuth, in degrees, at the exposure stop epoch.
	TopLeftStopAz float64 `json:"topLeftStopAz"`
	// The telescope pointing elevation, in degrees, at the exposure stop epoch.
	TopLeftStopEl float64 `json:"topLeftStopEl"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ExpStartTime          respjson.Field
		ImageType             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AnnotationKey         respjson.Field
		CalibrationKey        respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Description           respjson.Field
		EoObservations        respjson.Field
		ExpEndTime            respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		FrameFovHeight        respjson.Field
		FrameFovWidth         respjson.Field
		FrameHeightPixels     respjson.Field
		FrameWidthPixels      respjson.Field
		IDAttitudeSet         respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		IDSoiSet              respjson.Field
		ImageSetID            respjson.Field
		ImageSetLength        respjson.Field
		ImageSourceInfo       respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		PixelBitDepth         respjson.Field
		PixelFovHeight        respjson.Field
		PixelFovWidth         respjson.Field
		SatNo                 respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		SenQuat               respjson.Field
		SenQuatDot            respjson.Field
		Senx                  respjson.Field
		Seny                  respjson.Field
		Senz                  respjson.Field
		SequenceID            respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		TopLeftStartAz        respjson.Field
		TopLeftStartEl        respjson.Field
		TopLeftStopAz         respjson.Field
		TopLeftStopEl         respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkyImageryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SkyImageryGetResponse) UnmarshalJSON(data []byte) error {
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
type SkyImageryGetResponseDataMode string

const (
	SkyImageryGetResponseDataModeReal      SkyImageryGetResponseDataMode = "REAL"
	SkyImageryGetResponseDataModeTest      SkyImageryGetResponseDataMode = "TEST"
	SkyImageryGetResponseDataModeSimulated SkyImageryGetResponseDataMode = "SIMULATED"
	SkyImageryGetResponseDataModeExercise  SkyImageryGetResponseDataMode = "EXERCISE"
)

type SkyImageryQueryhelpResponse struct {
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
func (r SkyImageryQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SkyImageryQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of sky imagery data. Sky imagery is ground or space based
// telescope imagery of RSO's and includes metadata on the image (time, source,
// etc) as well as binary image content (e.g. FITS, EOSSA, EOCHIP, MP4). Binary
// content must be downloaded individually by ID using the 'getFile' operation.
type SkyImageryTupleResponse struct {
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
	DataMode SkyImageryTupleResponseDataMode `json:"dataMode,required"`
	// Start time of the exposure, in ISO 8601 UTC format with microsecond precision.
	ExpStartTime time.Time `json:"expStartTime,required" format:"date-time"`
	// The type of image associated with this record (e.g. FITS, EOSSA, EOCHIP, MP4).
	ImageType string `json:"imageType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Reference to an annotation document associated with this image.
	AnnotationKey string `json:"annotationKey"`
	// Reference to a calibration document associated with this image.
	CalibrationKey string `json:"calibrationKey"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue string `json:"checksumValue"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional name/description associated with this image.
	Description string `json:"description"`
	// Collection of linked EOObservations.
	EoObservations []shared.EoObservationFull `json:"eoObservations"`
	// End time of the exposure, in ISO 8601 UTC format with microsecond precision.
	ExpEndTime time.Time `json:"expEndTime" format:"date-time"`
	// Name of the uploaded image file.
	Filename string `json:"filename"`
	// Size of the image file, in bytes.
	Filesize int64 `json:"filesize"`
	// Field Of View frame height, in degrees.
	FrameFovHeight float64 `json:"frameFOVHeight"`
	// Field Of View frame width, in degrees.
	FrameFovWidth float64 `json:"frameFOVWidth"`
	// Frame height of the image, in number of pixels.
	FrameHeightPixels int64 `json:"frameHeightPixels"`
	// Frame width of the image, in number of pixels.
	FrameWidthPixels int64 `json:"frameWidthPixels"`
	// Optional identifier of the AttitudeSet data record describing the orientation of
	// an object body.
	IDAttitudeSet string `json:"idAttitudeSet"`
	// Optional identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// Optional unique identifier of the SOI Observation Set associated with this
	// image.
	IDSoiSet string `json:"idSOISet"`
	// The user-defined set ID of a sequence of images.
	ImageSetID string `json:"imageSetId"`
	// The number of images in an image set.
	ImageSetLength int64 `json:"imageSetLength"`
	// String that uniquely identifies the data source.
	ImageSourceInfo string `json:"imageSourceInfo"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the imaging source to indicate the target
	// onorbit object of this image. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the imaging source to indicate the sensor
	// identifier which produced this image. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Bit depth of the image, in number of pixels.
	PixelBitDepth int64 `json:"pixelBitDepth"`
	// Field Of View pixel height, in degrees.
	PixelFovHeight float64 `json:"pixelFOVHeight"`
	// Field Of View pixel width, in degrees.
	PixelFovWidth float64 `json:"pixelFOVWidth"`
	// Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sensor altitude at exposure start epoch, in km. If null, can be obtained from
	// sensor info.
	Senalt float64 `json:"senalt"`
	// Sensor latitude at exposure start epoch, in degrees. If null, can be obtained
	// from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor longitude at exposure start epoch, in degrees. If null, can be obtained
	// from sensor info. -180 to 180 degrees (negative values west of Prime Meridian).
	Senlon float64 `json:"senlon"`
	// The quaternion describing the rotation of the body-fixed frame used for this
	// system into the local geodetic frame, at exposure start epoch (expStartTime).
	// The array element order convention is scalar component first, followed by the
	// three vector components. For a vector u in the body-fixed frame, the
	// corresponding vector u' in the geodetic frame should satisfy u' = quq\*, where q
	// is this quaternion.
	SenQuat []float64 `json:"senQuat"`
	// The derivative of the quaternion describing the rotation of the body-fixed frame
	// used for this system into the local geodetic frame, exposure start epoch
	// (expStartTime). The array element order convention is scalar component first,
	// followed by the three vector components. For a vector u in the body-fixed frame,
	// the corresponding vector u' in the geodetic frame should satisfy u' = quq\*,
	// where q is this quaternion.
	SenQuatDot []float64 `json:"senQuatDot"`
	// Sensor x position at exposure start epoch, in km (if mobile/onorbit) in J2000
	// coordinate frame.
	Senx float64 `json:"senx"`
	// Sensor y position at exposure start epoch, in km (if mobile/onorbit) in J2000
	// coordinate frame.
	Seny float64 `json:"seny"`
	// Sensor z position at exposure start epoch, in km (if mobile/onorbit) in J2000
	// coordinate frame.
	Senz float64 `json:"senz"`
	// The sequence ID of an image within an image set.
	SequenceID int64 `json:"sequenceId"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// The telescope pointing azimuth, in degrees, at the exposure start epoch.
	TopLeftStartAz float64 `json:"topLeftStartAz"`
	// The telescope pointing elevation, in degrees, at the exposure start epoch.
	TopLeftStartEl float64 `json:"topLeftStartEl"`
	// The telescope pointing azimuth, in degrees, at the exposure stop epoch.
	TopLeftStopAz float64 `json:"topLeftStopAz"`
	// The telescope pointing elevation, in degrees, at the exposure stop epoch.
	TopLeftStopEl float64 `json:"topLeftStopEl"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ExpStartTime          respjson.Field
		ImageType             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AnnotationKey         respjson.Field
		CalibrationKey        respjson.Field
		ChecksumValue         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Description           respjson.Field
		EoObservations        respjson.Field
		ExpEndTime            respjson.Field
		Filename              respjson.Field
		Filesize              respjson.Field
		FrameFovHeight        respjson.Field
		FrameFovWidth         respjson.Field
		FrameHeightPixels     respjson.Field
		FrameWidthPixels      respjson.Field
		IDAttitudeSet         respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		IDSoiSet              respjson.Field
		ImageSetID            respjson.Field
		ImageSetLength        respjson.Field
		ImageSourceInfo       respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		PixelBitDepth         respjson.Field
		PixelFovHeight        respjson.Field
		PixelFovWidth         respjson.Field
		SatNo                 respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		SenQuat               respjson.Field
		SenQuatDot            respjson.Field
		Senx                  respjson.Field
		Seny                  respjson.Field
		Senz                  respjson.Field
		SequenceID            respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		TopLeftStartAz        respjson.Field
		TopLeftStartEl        respjson.Field
		TopLeftStopAz         respjson.Field
		TopLeftStopEl         respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkyImageryTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SkyImageryTupleResponse) UnmarshalJSON(data []byte) error {
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
type SkyImageryTupleResponseDataMode string

const (
	SkyImageryTupleResponseDataModeReal      SkyImageryTupleResponseDataMode = "REAL"
	SkyImageryTupleResponseDataModeTest      SkyImageryTupleResponseDataMode = "TEST"
	SkyImageryTupleResponseDataModeSimulated SkyImageryTupleResponseDataMode = "SIMULATED"
	SkyImageryTupleResponseDataModeExercise  SkyImageryTupleResponseDataMode = "EXERCISE"
)

type SkyImageryListParams struct {
	// Start time of the exposure, in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ExpStartTime time.Time        `query:"expStartTime,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkyImageryListParams]'s query parameters as `url.Values`.
func (r SkyImageryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SkyImageryCountParams struct {
	// Start time of the exposure, in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ExpStartTime time.Time        `query:"expStartTime,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkyImageryCountParams]'s query parameters as `url.Values`.
func (r SkyImageryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SkyImageryFileGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkyImageryFileGetParams]'s query parameters as
// `url.Values`.
func (r SkyImageryFileGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SkyImageryGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkyImageryGetParams]'s query parameters as `url.Values`.
func (r SkyImageryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SkyImageryTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Start time of the exposure, in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ExpStartTime time.Time        `query:"expStartTime,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkyImageryTupleParams]'s query parameters as `url.Values`.
func (r SkyImageryTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SkyImageryUploadZipParams struct {
	// Zip file containing files described in the specification
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

func (r SkyImageryUploadZipParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
