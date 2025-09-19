// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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
)

// FeatureAssessmentHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeatureAssessmentHistoryService] method instead.
type FeatureAssessmentHistoryService struct {
	Options []option.RequestOption
}

// NewFeatureAssessmentHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFeatureAssessmentHistoryService(opts ...option.RequestOption) (r FeatureAssessmentHistoryService) {
	r = FeatureAssessmentHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *FeatureAssessmentHistoryService) List(ctx context.Context, query FeatureAssessmentHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[FeatureAssessmentHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/featureassessment/history"
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

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *FeatureAssessmentHistoryService) ListAutoPaging(ctx context.Context, query FeatureAssessmentHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[FeatureAssessmentHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *FeatureAssessmentHistoryService) Aodr(ctx context.Context, query FeatureAssessmentHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/featureassessment/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *FeatureAssessmentHistoryService) Count(ctx context.Context, query FeatureAssessmentHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/featureassessment/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Feature assessments obtained from imagery analysis or other data analytics.
// Feature assessments are georeferenced terrestrial features such as marine
// vessels, vehicles, buildings, etc., or contain other types of non terrestrial
// assessments such as spacecraft structures. Geospatial queries are supported
// through either the regionText (WKT) or regionGeoJSON fields.
type FeatureAssessmentHistoryListResponse struct {
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
	DataMode FeatureAssessmentHistoryListResponseDataMode `json:"dataMode,required"`
	// Datetime type value associated with this record, in ISO 8601 UTC format with
	// millisecond precision.
	FeatureTs time.Time `json:"featureTs,required" format:"date-time"`
	// The Unit of Measure associated with this feature. If there are no physical units
	// associated with the feature a value of NONE should be specified.
	FeatureUoM string `json:"featureUoM,required"`
	// Unique identifier of the Analytic Imagery associated with this Feature
	// Assessment record.
	IDAnalyticImagery string `json:"idAnalyticImagery,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the feature assessment as projected
	// on the ground. GeoJSON Reference: https://geojson.org/. Ignored if included with
	// a POST or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Polygonal annotation limits, specified in pixels, as an array of arrays N x M.
	// Allows the image provider to highlight one or more polygonal area(s) of
	// interest. The array must contain NxM two-element arrays, where N is the number
	// of polygons of interest. The associated annotation(s) should be included in the
	// annText array.
	AnnLims [][]int64 `json:"annLims"`
	// Annotation text, a string array of annotation(s) corresponding to the
	// rectangular areas specified in annLims. This array contains the annotation text
	// associated with the areas of interest indicated in annLims, in order. This array
	// should contain one annotation per four values of the area (annLims) array.
	AnnText []string `json:"annText"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the feature assessment as projected on the ground.
	Area string `json:"area"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Descriptive or additional information associated with this feature/assessment.
	Assessment string `json:"assessment"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the feature assessment as
	// projected on the ground. WKT reference:
	// https://www.opengeospatial.org/standards/wkt-crs. Ignored if included with a
	// POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype string `json:"atype"`
	// Analytic confidence of feature accuracy (0 to 1).
	Confidence float64 `json:"confidence"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Feature Assessment ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID string `json:"externalId"`
	// An array of numeric feature/assessment values expressed in the specified unit of
	// measure (obUoM). Because of the variability of the Feature Assessment data
	// types, each record may employ a numeric observation value (featureValue), a
	// string observation value (featureString), a Boolean observation value
	// (featureBool), an array of numeric observation values (featureArray), or any
	// combination of these.
	FeatureArray []float64 `json:"featureArray"`
	// A boolean feature/assessment. Because of the variability of the Feature
	// Assessment data types, each record may employ a numeric observation value
	// (featureValue), a string observation value (featureString), a Boolean
	// observation value (featureBool), an array of numeric observation values
	// (featureArray), or any combination of these.
	FeatureBool bool `json:"featureBool"`
	// A single feature/assessment string expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Feature Assessment data types, each
	// record may employ a numeric observation value (featureValue), a string
	// observation value (featureString), a Boolean observation value (featureBool), an
	// array of numeric observation values (featureArray), or any combination of these.
	FeatureString string `json:"featureString"`
	// An array of string feature/assessment expressions. Because of the variability of
	// the Feature Assessment data types, each record may employ a numeric observation
	// value (featureValue), a string observation value (featureString), a Boolean
	// observation value (featureBool), an array of numeric observation values
	// (featureArray), or any combination of these.
	FeatureStringArray []string `json:"featureStringArray"`
	// A single feature/assessment value expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Feature Assessment data types, each
	// record may employ a numeric observation value (featureValue), a string
	// observation value (featureString), a Boolean observation value (featureBool), an
	// array of numeric observation values (featureArray), or any combination of these.
	FeatureValue float64 `json:"featureValue"`
	// The feature object heading, in degrees clockwise from true North at the object
	// location.
	Heading float64 `json:"heading"`
	// Estimated physical height of the feature, in meters.
	Height float64 `json:"height"`
	// Estimated physical length of the feature, in meters.
	Length float64 `json:"length"`
	// Source defined name of the feature associated with this record. If an annotation
	// for this feature element exists on the parent image it can be referenced here.
	Name string `json:"name"`
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
	// Feature's speed of travel, in meters per second.
	Speed float64 `json:"speed"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this activity or event. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps', 'srcIds', and 'srcTs' arrays must contain the same number of
	// elements. See the corresponding srcTyps array element for the data type of the
	// UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds"`
	// Array of the primary timestamps, in ISO 8601 UTC format, with appropriate
	// precision for the datatype of each correspondng srcTyp/srcId record. See the
	// associated 'srcTyps' and 'srcIds' arrays for the record type and UUID,
	// respectively, positionally corresponding to the record types in this array. The
	// 'srcTyps', 'srcIds', and 'srcTs' arrays must contain the same number of
	// elements. These timestamps are included to support services which do not include
	// a GET by {id} operation. If referencing a datatype which does not include a
	// primary timestamp, the corresponding srcTs array element should be included as
	// null.
	SrcTs []time.Time `json:"srcTs" format:"date-time"`
	// Array of UDL record types (AIS, GROUNDIMAGE, MTI, ONORBIT, POI, SAR, SKYIMAGE,
	// SOI, TRACK) related to this feature assessment. See the associated 'srcIds' and
	// 'srcTs' arrays for the record UUIDs and timetsmps. respectively, positionally
	// corresponding to the record types in this array. The 'srcTyps', 'srcIds', and
	// 'srcTs' arrays must contain the same number of elements.
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
	// The type of feature (e.g. AIRCRAFT, ANTENNA, SOLAR ARRAY, SITE, STRUCTURE,
	// VESSEL, VEHICLE, UNKNOWN, etc.) detailed in this feature assessment record. This
	// type may be a primary feature within an image, for example a VESSEL, or may be a
	// component or characteristic of a primary feature, for example an ANTENNA mounted
	// on a vessel.
	Type string `json:"type"`
	// Estimated physical width of the feature, in meters.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		FeatureTs             respjson.Field
		FeatureUoM            respjson.Field
		IDAnalyticImagery     respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Agjson                respjson.Field
		Andims                respjson.Field
		AnnLims               respjson.Field
		AnnText               respjson.Field
		Area                  respjson.Field
		Asrid                 respjson.Field
		Assessment            respjson.Field
		Atext                 respjson.Field
		Atype                 respjson.Field
		Confidence            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ExternalID            respjson.Field
		FeatureArray          respjson.Field
		FeatureBool           respjson.Field
		FeatureString         respjson.Field
		FeatureStringArray    respjson.Field
		FeatureValue          respjson.Field
		Heading               respjson.Field
		Height                respjson.Field
		Length                respjson.Field
		Name                  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SourceDl              respjson.Field
		Speed                 respjson.Field
		SrcIDs                respjson.Field
		SrcTs                 respjson.Field
		SrcTyps               respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureAssessmentHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureAssessmentHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type FeatureAssessmentHistoryListResponseDataMode string

const (
	FeatureAssessmentHistoryListResponseDataModeReal      FeatureAssessmentHistoryListResponseDataMode = "REAL"
	FeatureAssessmentHistoryListResponseDataModeTest      FeatureAssessmentHistoryListResponseDataMode = "TEST"
	FeatureAssessmentHistoryListResponseDataModeSimulated FeatureAssessmentHistoryListResponseDataMode = "SIMULATED"
	FeatureAssessmentHistoryListResponseDataModeExercise  FeatureAssessmentHistoryListResponseDataMode = "EXERCISE"
)

type FeatureAssessmentHistoryListParams struct {
	// Unique identifier of the Analytic Imagery associated with this Feature
	// Assessment record.
	IDAnalyticImagery string `query:"idAnalyticImagery,required" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FeatureAssessmentHistoryListParams]'s query parameters as
// `url.Values`.
func (r FeatureAssessmentHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FeatureAssessmentHistoryAodrParams struct {
	// Unique identifier of the Analytic Imagery associated with this Feature
	// Assessment record.
	IDAnalyticImagery string `query:"idAnalyticImagery,required" json:"-"`
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

// URLQuery serializes [FeatureAssessmentHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r FeatureAssessmentHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FeatureAssessmentHistoryCountParams struct {
	// Unique identifier of the Analytic Imagery associated with this Feature
	// Assessment record.
	IDAnalyticImagery string           `query:"idAnalyticImagery,required" json:"-"`
	FirstResult       param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults        param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FeatureAssessmentHistoryCountParams]'s query parameters as
// `url.Values`.
func (r FeatureAssessmentHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
