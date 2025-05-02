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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// GnssrawifService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGnssrawifService] method instead.
type GnssrawifService struct {
	Options []option.RequestOption
}

// NewGnssrawifService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGnssrawifService(opts ...option.RequestOption) (r GnssrawifService) {
	r = GnssrawifService{}
	r.Options = opts
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *GnssrawifService) List(ctx context.Context, query GnssrawifListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GnssrawifListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/gnssrawif"
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
func (r *GnssrawifService) ListAutoPaging(ctx context.Context, query GnssrawifListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GnssrawifListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *GnssrawifService) Count(ctx context.Context, query GnssrawifCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/gnssrawif/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single GNSSRAWIF hdf5 file by its unique ID passed as
// a path parameter. The file is returned as an attachment Content-Disposition.
func (r *GnssrawifService) FileGet(ctx context.Context, id string, query GnssrawifFileGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/gnssrawif/getFile/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single GNSSRawIF by its unique ID passed as a path
// parameter.
func (r *GnssrawifService) Get(ctx context.Context, id string, query GnssrawifGetParams, opts ...option.RequestOption) (res *GnssRawIfFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/gnssrawif/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *GnssrawifService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/gnssrawif/queryhelp"
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
func (r *GnssrawifService) Tuple(ctx context.Context, query GnssrawifTupleParams, opts ...option.RequestOption) (res *[]GnssRawIfFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/gnssrawif/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload an HDF5 file with its metadata.
//
// The request body requires a zip file containing exactly two files:\
// 1\) A file with the `.json` file extension whose content conforms to the `GNSSRawIF_Ingest`
// schema.\
// 2\) A file with the `.hdf5` file extension.
//
// The JSON and HDF5 files will be associated with each other via the `id` field.
// Query the metadata via `GET /udl/gnssrawif` and use
// `GET /udl/gnssrawif/getFile/{id}` to retrieve the HDF5 file.
//
// This operation only accepts application/zip media. The application/json request
// body is documented to provide a convenient reference to the ingest schema.
//
// This operation is intended to be used for automated feeds into UDL. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *GnssrawifService) UploadZip(ctx context.Context, body GnssrawifUploadZipParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-gnssrawif"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Global Navigation Satellite System (GNSS) Raw Intermediate Frequency (IF) data
// are the recorded streams of raw signal samples after down-conversion of the
// received signal to IF and prior to any processing onboard the receiving
// spacecraft. These data sets are processed in various geophysical applications
// and used to characterize Electromagnetic Interference (EMI) in the operating
// environment.
type GnssrawifListResponse struct {
	// The center frequency, in MHz, of the observation bands. More than one band may
	// be reported in each binary file, so this is an array of the center frequency of
	// each band (including an array of length 1 if only one band is present).
	CenterFreq []float64 `json:"centerFreq,required"`
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
	DataMode GnssrawifListResponseDataMode `json:"dataMode,required"`
	// End time of the data contained in the associated binary file, in ISO 8601 UTC
	// format with microsecond precision.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// The file name of the Raw IF Binary file. The files should be in the Hierarchical
	// Data Format (HDF5).
	FileName string `json:"fileName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Start time of the data contained in the associated binary file, in ISO 8601 UTC
	// format with microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The number of bits in each datum, for example 2 for 2-bit integers or 8 for
	// 8-bit integers.
	BitDepth int64 `json:"bitDepth"`
	// Unit vector of the outward facing direction of the receiver boresight in a
	// body-fixed coordinate system.
	Boresight []float64 `json:"boresight"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The amount of data generated per unit time, expressed in Megabytes/minute.
	DataRate float64 `json:"dataRate"`
	// Differential Code Biases (DCBs) are the systematic errors, or biases, between
	// two GNSS code observations at the same or different frequencies. If applicable,
	// this field should contain an array of DBC with length equal to the number of
	// frequencies in the binary file. The reference frequency should show DCB equal
	// to 0. If null, it is assumed that there is no DCB (all values are 0).
	DiffCodeBias []float64 `json:"diffCodeBias"`
	// Spacecraft altitude at end time (endTime), expressed in kilometers above WGS-84
	// ellipsoid.
	EndAlt float64 `json:"endAlt"`
	// WGS-84 spacecraft latitude sub-point at end time (endTime), represented as -90
	// to 90 degrees (negative values south of equator).
	EndLat float64 `json:"endLat"`
	// WGS-84 spacecraft longitude sub-point at end time (endTime), represented as -180
	// to 180 degrees (negative values west of Prime Meridian).
	EndLon float64 `json:"endLon"`
	// Unique identifier of the parent Ephemeris Set, if this data is correlated with
	// an Ephemeris. If reporting for a spacecraft with multiple onboard GNSS
	// receivers, this ID may be associated with multiple GNSS Raw IF records if each
	// receiver is synced to the ephemeris points.
	EsID string `json:"esId"`
	// Optional source-provided identifier for this collection event. This field can be
	// used to associate records related to the same event.
	EventID string `json:"eventId"`
	// The binary file size, in bytes, auto-populated by the system. The maximum file
	// size for this service is 5,000,000 Bytes (5 MB). Files exceeding the maximum
	// size will be rejected.
	FileSize int64 `json:"fileSize"`
	// Unique identifier of the primary satellite on-orbit object.
	IDOnOrbit string `json:"idOnOrbit"`
	// The center frequency, in MHz, after downconversion to intermediate frequency. If
	// provided, this array should have the same length as centerFreqs.
	IfFreq []float64 `json:"ifFreq"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the data source to indicate the target object of
	// this record. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// After converting the raw data in the file, to generate (frequency-space) spectra
	// some sets require an additional transformation or correction based on details of
	// the hardware that recorded it. This field marks any such transformations.
	// Currently supported options are NONE (no correction) and MIRRORED (frequency
	// axis is flipped around the corresponding value of ifFreq). If null, it is
	// assumed that NONE applies to all frequency bands.
	PostFourier []string `json:"postFourier"`
	// The quaternion describing the rotation of the body-fixed frame used for this
	// system into the local geodetic frame, at the sample start time (startTime). The
	// array element order convention is scalar component first, followed by the three
	// vector components. For a vector u in the body-fixed frame, the corresponding
	// vector u' in the geodetic frame should satisfy u' = quq\*, where q is this
	// quaternion. The quaternion should be normalized to 1.
	Quat []float64 `json:"quat"`
	// The number or ID of the GNSS receiver associated with this data. If reporting
	// for multiple receivers a separate record should be generated for each. If null,
	// it is assumed to indicate that only one receiver is present, or reported.
	Receiver string `json:"receiver"`
	// The number of samples taken per second.
	SampleRate []int64 `json:"sampleRate"`
	// The sample type associated with the IF data. REAL for data with only an
	// I-component or COMPLEX for data with both I & Q components.
	SampleType string `json:"sampleType"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// The sequence number of a raw IF record/file within a set. Sequence number should
	// start at 1. If null, then it is assumed that the order of records within a raw
	// IF set is not relevant.
	SequenceID int64 `json:"sequenceID"`
	// User-defined ID of a set or sequence of records/files. Used to associate a set
	// of related raw IF records.
	SetID string `json:"setId"`
	// The number of raw IF records/files in a set.
	SetLength int64 `json:"setLength"`
	// Array of UUIDs of the UDL data records associated with this GNSSRawIF record.
	// See the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object (e.g.
	// /udl/gnssobservationset/{uuid}).
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (GNSSSET) associated with this GNSSRawIF record. See
	// the associated 'srcIds' array for the record UUIDs, positionally corresponding
	// to the record types in this array. The 'srcTyps' and 'srcIds' arrays must match
	// in size.
	SrcTyps []string `json:"srcTyps"`
	// Spacecraft altitude at start time (startTime), expressed in kilometers above
	// WGS-84 ellipsoid.
	StartAlt float64 `json:"startAlt"`
	// The index of the sample within the associated binary file that corresponds to
	// the startTime indicated in this record. This is especially useful on high
	// sample-rate sensors when some samples are less than one microsecond before the
	// value of startTime. This index is 0-based. If null, the startIndex is assumed to
	// be 0.
	StartIndex int64 `json:"startIndex"`
	// WGS-84 spacecraft latitude sub-point at start time (startTime), represented as
	// -90 to 90 degrees (negative values south of equator).
	StartLat float64 `json:"startLat"`
	// WGS-84 spacecraft longitude sub-point at start time (startTime), represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	StartLon float64 `json:"startLon"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CenterFreq            resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		EndTime               resp.Field
		FileName              resp.Field
		Source                resp.Field
		StartTime             resp.Field
		ID                    resp.Field
		BitDepth              resp.Field
		Boresight             resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DataRate              resp.Field
		DiffCodeBias          resp.Field
		EndAlt                resp.Field
		EndLat                resp.Field
		EndLon                resp.Field
		EsID                  resp.Field
		EventID               resp.Field
		FileSize              resp.Field
		IDOnOrbit             resp.Field
		IfFreq                resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		PostFourier           resp.Field
		Quat                  resp.Field
		Receiver              resp.Field
		SampleRate            resp.Field
		SampleType            resp.Field
		SatNo                 resp.Field
		SequenceID            resp.Field
		SetID                 resp.Field
		SetLength             resp.Field
		SrcIDs                resp.Field
		SrcTyps               resp.Field
		StartAlt              resp.Field
		StartIndex            resp.Field
		StartLat              resp.Field
		StartLon              resp.Field
		Tags                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GnssrawifListResponse) RawJSON() string { return r.JSON.raw }
func (r *GnssrawifListResponse) UnmarshalJSON(data []byte) error {
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
type GnssrawifListResponseDataMode string

const (
	GnssrawifListResponseDataModeReal      GnssrawifListResponseDataMode = "REAL"
	GnssrawifListResponseDataModeTest      GnssrawifListResponseDataMode = "TEST"
	GnssrawifListResponseDataModeSimulated GnssrawifListResponseDataMode = "SIMULATED"
	GnssrawifListResponseDataModeExercise  GnssrawifListResponseDataMode = "EXERCISE"
)

type GnssrawifListParams struct {
	// Start time of the data contained in the associated binary file, in ISO 8601 UTC
	// format with microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GnssrawifListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [GnssrawifListParams]'s query parameters as `url.Values`.
func (r GnssrawifListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GnssrawifCountParams struct {
	// Start time of the data contained in the associated binary file, in ISO 8601 UTC
	// format with microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GnssrawifCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [GnssrawifCountParams]'s query parameters as `url.Values`.
func (r GnssrawifCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GnssrawifFileGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GnssrawifFileGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [GnssrawifFileGetParams]'s query parameters as `url.Values`.
func (r GnssrawifFileGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GnssrawifGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GnssrawifGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [GnssrawifGetParams]'s query parameters as `url.Values`.
func (r GnssrawifGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GnssrawifTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Start time of the data contained in the associated binary file, in ISO 8601 UTC
	// format with microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GnssrawifTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [GnssrawifTupleParams]'s query parameters as `url.Values`.
func (r GnssrawifTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GnssrawifUploadZipParams struct {
	// Zip file containing files described in the specification
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GnssrawifUploadZipParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r GnssrawifUploadZipParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
