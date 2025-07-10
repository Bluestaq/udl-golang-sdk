// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// FeatureAssessmentService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeatureAssessmentService] method instead.
type FeatureAssessmentService struct {
	Options []option.RequestOption
	History FeatureAssessmentHistoryService
}

// NewFeatureAssessmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFeatureAssessmentService(opts ...option.RequestOption) (r FeatureAssessmentService) {
	r = FeatureAssessmentService{}
	r.Options = opts
	r.History = NewFeatureAssessmentHistoryService(opts...)
	return
}

// Service operation to take a single FeatureAssessment record as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *FeatureAssessmentService) New(ctx context.Context, body FeatureAssessmentNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/featureassessment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single FeatureAssessment record by its unique ID
// passed as a path parameter.
func (r *FeatureAssessmentService) Get(ctx context.Context, id string, query FeatureAssessmentGetParams, opts ...option.RequestOption) (res *FeatureAssessmentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/featureassessment/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *FeatureAssessmentService) Count(ctx context.Context, query FeatureAssessmentCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/featureassessment/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// FeatureAssessment records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *FeatureAssessmentService) NewBulk(ctx context.Context, body FeatureAssessmentNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/featureassessment/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *FeatureAssessmentService) Query(ctx context.Context, query FeatureAssessmentQueryParams, opts ...option.RequestOption) (res *[]FeatureAssessmentQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/featureassessment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *FeatureAssessmentService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *FeatureAssessmentQueryHelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/featureassessment/queryhelp"
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
func (r *FeatureAssessmentService) Tuple(ctx context.Context, query FeatureAssessmentTupleParams, opts ...option.RequestOption) (res *[]FeatureAssessmentTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/featureassessment/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple FeatureAssessment records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *FeatureAssessmentService) UnvalidatedPublish(ctx context.Context, body FeatureAssessmentUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-featureassessment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Feature assessments obtained from imagery analysis or other data analytics.
// Feature assessments are georeferenced terrestrial features such as marine
// vessels, vehicles, buildings, etc., or contain other types of non terrestrial
// assessments such as spacecraft structures. Geospatial queries are supported
// through either the regionText (WKT) or regionGeoJSON fields.
type FeatureAssessmentGetResponse struct {
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
	DataMode FeatureAssessmentGetResponseDataMode `json:"dataMode,required"`
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
func (r FeatureAssessmentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureAssessmentGetResponse) UnmarshalJSON(data []byte) error {
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
type FeatureAssessmentGetResponseDataMode string

const (
	FeatureAssessmentGetResponseDataModeReal      FeatureAssessmentGetResponseDataMode = "REAL"
	FeatureAssessmentGetResponseDataModeTest      FeatureAssessmentGetResponseDataMode = "TEST"
	FeatureAssessmentGetResponseDataModeSimulated FeatureAssessmentGetResponseDataMode = "SIMULATED"
	FeatureAssessmentGetResponseDataModeExercise  FeatureAssessmentGetResponseDataMode = "EXERCISE"
)

// Feature assessments obtained from imagery analysis or other data analytics.
// Feature assessments are georeferenced terrestrial features such as marine
// vessels, vehicles, buildings, etc., or contain other types of non terrestrial
// assessments such as spacecraft structures. Geospatial queries are supported
// through either the regionText (WKT) or regionGeoJSON fields.
type FeatureAssessmentQueryResponse struct {
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
	DataMode FeatureAssessmentQueryResponseDataMode `json:"dataMode,required"`
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
		TransactionID         respjson.Field
		Type                  respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureAssessmentQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureAssessmentQueryResponse) UnmarshalJSON(data []byte) error {
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
type FeatureAssessmentQueryResponseDataMode string

const (
	FeatureAssessmentQueryResponseDataModeReal      FeatureAssessmentQueryResponseDataMode = "REAL"
	FeatureAssessmentQueryResponseDataModeTest      FeatureAssessmentQueryResponseDataMode = "TEST"
	FeatureAssessmentQueryResponseDataModeSimulated FeatureAssessmentQueryResponseDataMode = "SIMULATED"
	FeatureAssessmentQueryResponseDataModeExercise  FeatureAssessmentQueryResponseDataMode = "EXERCISE"
)

type FeatureAssessmentQueryHelpResponse struct {
	AodrSupported         bool                                          `json:"aodrSupported"`
	ClassificationMarking string                                        `json:"classificationMarking"`
	Description           string                                        `json:"description"`
	HistorySupported      bool                                          `json:"historySupported"`
	Name                  string                                        `json:"name"`
	Parameters            []FeatureAssessmentQueryHelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                      `json:"requiredRoles"`
	RestSupported         bool                                          `json:"restSupported"`
	SortSupported         bool                                          `json:"sortSupported"`
	TypeName              string                                        `json:"typeName"`
	Uri                   string                                        `json:"uri"`
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
func (r FeatureAssessmentQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureAssessmentQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureAssessmentQueryHelpResponseParameter struct {
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
func (r FeatureAssessmentQueryHelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *FeatureAssessmentQueryHelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature assessments obtained from imagery analysis or other data analytics.
// Feature assessments are georeferenced terrestrial features such as marine
// vessels, vehicles, buildings, etc., or contain other types of non terrestrial
// assessments such as spacecraft structures. Geospatial queries are supported
// through either the regionText (WKT) or regionGeoJSON fields.
type FeatureAssessmentTupleResponse struct {
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
	DataMode FeatureAssessmentTupleResponseDataMode `json:"dataMode,required"`
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
func (r FeatureAssessmentTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureAssessmentTupleResponse) UnmarshalJSON(data []byte) error {
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
type FeatureAssessmentTupleResponseDataMode string

const (
	FeatureAssessmentTupleResponseDataModeReal      FeatureAssessmentTupleResponseDataMode = "REAL"
	FeatureAssessmentTupleResponseDataModeTest      FeatureAssessmentTupleResponseDataMode = "TEST"
	FeatureAssessmentTupleResponseDataModeSimulated FeatureAssessmentTupleResponseDataMode = "SIMULATED"
	FeatureAssessmentTupleResponseDataModeExercise  FeatureAssessmentTupleResponseDataMode = "EXERCISE"
)

type FeatureAssessmentNewParams struct {
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
	DataMode FeatureAssessmentNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the feature assessment as projected
	// on the ground. GeoJSON Reference: https://geojson.org/. Ignored if included with
	// a POST or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the feature assessment as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Descriptive or additional information associated with this feature/assessment.
	Assessment param.Opt[string] `json:"assessment,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the feature assessment as
	// projected on the ground. WKT reference:
	// https://www.opengeospatial.org/standards/wkt-crs. Ignored if included with a
	// POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype param.Opt[string] `json:"atype,omitzero"`
	// Analytic confidence of feature accuracy (0 to 1).
	Confidence param.Opt[float64] `json:"confidence,omitzero"`
	// Feature Assessment ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// A boolean feature/assessment. Because of the variability of the Feature
	// Assessment data types, each record may employ a numeric observation value
	// (featureValue), a string observation value (featureString), a Boolean
	// observation value (featureBool), an array of numeric observation values
	// (featureArray), or any combination of these.
	FeatureBool param.Opt[bool] `json:"featureBool,omitzero"`
	// A single feature/assessment string expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Feature Assessment data types, each
	// record may employ a numeric observation value (featureValue), a string
	// observation value (featureString), a Boolean observation value (featureBool), an
	// array of numeric observation values (featureArray), or any combination of these.
	FeatureString param.Opt[string] `json:"featureString,omitzero"`
	// A single feature/assessment value expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Feature Assessment data types, each
	// record may employ a numeric observation value (featureValue), a string
	// observation value (featureString), a Boolean observation value (featureBool), an
	// array of numeric observation values (featureArray), or any combination of these.
	FeatureValue param.Opt[float64] `json:"featureValue,omitzero"`
	// The feature object heading, in degrees clockwise from true North at the object
	// location.
	Heading param.Opt[float64] `json:"heading,omitzero"`
	// Estimated physical height of the feature, in meters.
	Height param.Opt[float64] `json:"height,omitzero"`
	// Estimated physical length of the feature, in meters.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Source defined name of the feature associated with this record. If an annotation
	// for this feature element exists on the parent image it can be referenced here.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Feature's speed of travel, in meters per second.
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The type of feature (e.g. AIRCRAFT, ANTENNA, SOLAR ARRAY, SITE, STRUCTURE,
	// VESSEL, VEHICLE, UNKNOWN, etc.) detailed in this feature assessment record. This
	// type may be a primary feature within an image, for example a VESSEL, or may be a
	// component or characteristic of a primary feature, for example an ANTENNA mounted
	// on a vessel.
	Type param.Opt[string] `json:"type,omitzero"`
	// Estimated physical width of the feature, in meters.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Polygonal annotation limits, specified in pixels, as an array of arrays N x M.
	// Allows the image provider to highlight one or more polygonal area(s) of
	// interest. The array must contain NxM two-element arrays, where N is the number
	// of polygons of interest. The associated annotation(s) should be included in the
	// annText array.
	AnnLims [][]int64 `json:"annLims,omitzero"`
	// Annotation text, a string array of annotation(s) corresponding to the
	// rectangular areas specified in annLims. This array contains the annotation text
	// associated with the areas of interest indicated in annLims, in order. This array
	// should contain one annotation per four values of the area (annLims) array.
	AnnText []string `json:"annText,omitzero"`
	// An array of numeric feature/assessment values expressed in the specified unit of
	// measure (obUoM). Because of the variability of the Feature Assessment data
	// types, each record may employ a numeric observation value (featureValue), a
	// string observation value (featureString), a Boolean observation value
	// (featureBool), an array of numeric observation values (featureArray), or any
	// combination of these.
	FeatureArray []float64 `json:"featureArray,omitzero"`
	// An array of string feature/assessment expressions. Because of the variability of
	// the Feature Assessment data types, each record may employ a numeric observation
	// value (featureValue), a string observation value (featureString), a Boolean
	// observation value (featureBool), an array of numeric observation values
	// (featureArray), or any combination of these.
	FeatureStringArray []string `json:"featureStringArray,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this activity or event. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps', 'srcIds', and 'srcTs' arrays must contain the same number of
	// elements. See the corresponding srcTyps array element for the data type of the
	// UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of the primary timestamps, in ISO 8601 UTC format, with appropriate
	// precision for the datatype of each correspondng srcTyp/srcId record. See the
	// associated 'srcTyps' and 'srcIds' arrays for the record type and UUID,
	// respectively, positionally corresponding to the record types in this array. The
	// 'srcTyps', 'srcIds', and 'srcTs' arrays must contain the same number of
	// elements. These timestamps are included to support services which do not include
	// a GET by {id} operation. If referencing a datatype which does not include a
	// primary timestamp, the corresponding srcTs array element should be included as
	// null.
	SrcTs []time.Time `json:"srcTs,omitzero" format:"date-time"`
	// Array of UDL record types (AIS, GROUNDIMAGE, MTI, ONORBIT, POI, SAR, SKYIMAGE,
	// SOI, TRACK) related to this feature assessment. See the associated 'srcIds' and
	// 'srcTs' arrays for the record UUIDs and timetsmps. respectively, positionally
	// corresponding to the record types in this array. The 'srcTyps', 'srcIds', and
	// 'srcTs' arrays must contain the same number of elements.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FeatureAssessmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FeatureAssessmentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeatureAssessmentNewParams) UnmarshalJSON(data []byte) error {
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
type FeatureAssessmentNewParamsDataMode string

const (
	FeatureAssessmentNewParamsDataModeReal      FeatureAssessmentNewParamsDataMode = "REAL"
	FeatureAssessmentNewParamsDataModeTest      FeatureAssessmentNewParamsDataMode = "TEST"
	FeatureAssessmentNewParamsDataModeSimulated FeatureAssessmentNewParamsDataMode = "SIMULATED"
	FeatureAssessmentNewParamsDataModeExercise  FeatureAssessmentNewParamsDataMode = "EXERCISE"
)

type FeatureAssessmentGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FeatureAssessmentGetParams]'s query parameters as
// `url.Values`.
func (r FeatureAssessmentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FeatureAssessmentCountParams struct {
	// Unique identifier of the Analytic Imagery associated with this Feature
	// Assessment record.
	IDAnalyticImagery string           `query:"idAnalyticImagery,required" json:"-"`
	FirstResult       param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults        param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FeatureAssessmentCountParams]'s query parameters as
// `url.Values`.
func (r FeatureAssessmentCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FeatureAssessmentNewBulkParams struct {
	Body []FeatureAssessmentNewBulkParamsBody
	paramObj
}

func (r FeatureAssessmentNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *FeatureAssessmentNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Feature assessments obtained from imagery analysis or other data analytics.
// Feature assessments are georeferenced terrestrial features such as marine
// vessels, vehicles, buildings, etc., or contain other types of non terrestrial
// assessments such as spacecraft structures. Geospatial queries are supported
// through either the regionText (WKT) or regionGeoJSON fields.
//
// The properties ClassificationMarking, DataMode, FeatureTs, FeatureUoM,
// IDAnalyticImagery, Source are required.
type FeatureAssessmentNewBulkParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
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
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the feature assessment as projected
	// on the ground. GeoJSON Reference: https://geojson.org/. Ignored if included with
	// a POST or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the feature assessment as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Descriptive or additional information associated with this feature/assessment.
	Assessment param.Opt[string] `json:"assessment,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the feature assessment as
	// projected on the ground. WKT reference:
	// https://www.opengeospatial.org/standards/wkt-crs. Ignored if included with a
	// POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype param.Opt[string] `json:"atype,omitzero"`
	// Analytic confidence of feature accuracy (0 to 1).
	Confidence param.Opt[float64] `json:"confidence,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Feature Assessment ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// A boolean feature/assessment. Because of the variability of the Feature
	// Assessment data types, each record may employ a numeric observation value
	// (featureValue), a string observation value (featureString), a Boolean
	// observation value (featureBool), an array of numeric observation values
	// (featureArray), or any combination of these.
	FeatureBool param.Opt[bool] `json:"featureBool,omitzero"`
	// A single feature/assessment string expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Feature Assessment data types, each
	// record may employ a numeric observation value (featureValue), a string
	// observation value (featureString), a Boolean observation value (featureBool), an
	// array of numeric observation values (featureArray), or any combination of these.
	FeatureString param.Opt[string] `json:"featureString,omitzero"`
	// A single feature/assessment value expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Feature Assessment data types, each
	// record may employ a numeric observation value (featureValue), a string
	// observation value (featureString), a Boolean observation value (featureBool), an
	// array of numeric observation values (featureArray), or any combination of these.
	FeatureValue param.Opt[float64] `json:"featureValue,omitzero"`
	// The feature object heading, in degrees clockwise from true North at the object
	// location.
	Heading param.Opt[float64] `json:"heading,omitzero"`
	// Estimated physical height of the feature, in meters.
	Height param.Opt[float64] `json:"height,omitzero"`
	// Estimated physical length of the feature, in meters.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Source defined name of the feature associated with this record. If an annotation
	// for this feature element exists on the parent image it can be referenced here.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Feature's speed of travel, in meters per second.
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The type of feature (e.g. AIRCRAFT, ANTENNA, SOLAR ARRAY, SITE, STRUCTURE,
	// VESSEL, VEHICLE, UNKNOWN, etc.) detailed in this feature assessment record. This
	// type may be a primary feature within an image, for example a VESSEL, or may be a
	// component or characteristic of a primary feature, for example an ANTENNA mounted
	// on a vessel.
	Type param.Opt[string] `json:"type,omitzero"`
	// Estimated physical width of the feature, in meters.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Polygonal annotation limits, specified in pixels, as an array of arrays N x M.
	// Allows the image provider to highlight one or more polygonal area(s) of
	// interest. The array must contain NxM two-element arrays, where N is the number
	// of polygons of interest. The associated annotation(s) should be included in the
	// annText array.
	AnnLims [][]int64 `json:"annLims,omitzero"`
	// Annotation text, a string array of annotation(s) corresponding to the
	// rectangular areas specified in annLims. This array contains the annotation text
	// associated with the areas of interest indicated in annLims, in order. This array
	// should contain one annotation per four values of the area (annLims) array.
	AnnText []string `json:"annText,omitzero"`
	// An array of numeric feature/assessment values expressed in the specified unit of
	// measure (obUoM). Because of the variability of the Feature Assessment data
	// types, each record may employ a numeric observation value (featureValue), a
	// string observation value (featureString), a Boolean observation value
	// (featureBool), an array of numeric observation values (featureArray), or any
	// combination of these.
	FeatureArray []float64 `json:"featureArray,omitzero"`
	// An array of string feature/assessment expressions. Because of the variability of
	// the Feature Assessment data types, each record may employ a numeric observation
	// value (featureValue), a string observation value (featureString), a Boolean
	// observation value (featureBool), an array of numeric observation values
	// (featureArray), or any combination of these.
	FeatureStringArray []string `json:"featureStringArray,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this activity or event. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps', 'srcIds', and 'srcTs' arrays must contain the same number of
	// elements. See the corresponding srcTyps array element for the data type of the
	// UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of the primary timestamps, in ISO 8601 UTC format, with appropriate
	// precision for the datatype of each correspondng srcTyp/srcId record. See the
	// associated 'srcTyps' and 'srcIds' arrays for the record type and UUID,
	// respectively, positionally corresponding to the record types in this array. The
	// 'srcTyps', 'srcIds', and 'srcTs' arrays must contain the same number of
	// elements. These timestamps are included to support services which do not include
	// a GET by {id} operation. If referencing a datatype which does not include a
	// primary timestamp, the corresponding srcTs array element should be included as
	// null.
	SrcTs []time.Time `json:"srcTs,omitzero" format:"date-time"`
	// Array of UDL record types (AIS, GROUNDIMAGE, MTI, ONORBIT, POI, SAR, SKYIMAGE,
	// SOI, TRACK) related to this feature assessment. See the associated 'srcIds' and
	// 'srcTs' arrays for the record UUIDs and timetsmps. respectively, positionally
	// corresponding to the record types in this array. The 'srcTyps', 'srcIds', and
	// 'srcTs' arrays must contain the same number of elements.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FeatureAssessmentNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow FeatureAssessmentNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeatureAssessmentNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FeatureAssessmentNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type FeatureAssessmentQueryParams struct {
	// Unique identifier of the Analytic Imagery associated with this Feature
	// Assessment record.
	IDAnalyticImagery string           `query:"idAnalyticImagery,required" json:"-"`
	FirstResult       param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults        param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FeatureAssessmentQueryParams]'s query parameters as
// `url.Values`.
func (r FeatureAssessmentQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FeatureAssessmentTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Unique identifier of the Analytic Imagery associated with this Feature
	// Assessment record.
	IDAnalyticImagery string           `query:"idAnalyticImagery,required" json:"-"`
	FirstResult       param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults        param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FeatureAssessmentTupleParams]'s query parameters as
// `url.Values`.
func (r FeatureAssessmentTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FeatureAssessmentUnvalidatedPublishParams struct {
	Body []FeatureAssessmentUnvalidatedPublishParamsBody
	paramObj
}

func (r FeatureAssessmentUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *FeatureAssessmentUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Feature assessments obtained from imagery analysis or other data analytics.
// Feature assessments are georeferenced terrestrial features such as marine
// vessels, vehicles, buildings, etc., or contain other types of non terrestrial
// assessments such as spacecraft structures. Geospatial queries are supported
// through either the regionText (WKT) or regionGeoJSON fields.
//
// The properties ClassificationMarking, DataMode, FeatureTs, FeatureUoM,
// IDAnalyticImagery, Source are required.
type FeatureAssessmentUnvalidatedPublishParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
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
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the feature assessment as projected
	// on the ground. GeoJSON Reference: https://geojson.org/. Ignored if included with
	// a POST or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the feature assessment as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Descriptive or additional information associated with this feature/assessment.
	Assessment param.Opt[string] `json:"assessment,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the feature assessment as
	// projected on the ground. WKT reference:
	// https://www.opengeospatial.org/standards/wkt-crs. Ignored if included with a
	// POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype param.Opt[string] `json:"atype,omitzero"`
	// Analytic confidence of feature accuracy (0 to 1).
	Confidence param.Opt[float64] `json:"confidence,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Feature Assessment ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// A boolean feature/assessment. Because of the variability of the Feature
	// Assessment data types, each record may employ a numeric observation value
	// (featureValue), a string observation value (featureString), a Boolean
	// observation value (featureBool), an array of numeric observation values
	// (featureArray), or any combination of these.
	FeatureBool param.Opt[bool] `json:"featureBool,omitzero"`
	// A single feature/assessment string expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Feature Assessment data types, each
	// record may employ a numeric observation value (featureValue), a string
	// observation value (featureString), a Boolean observation value (featureBool), an
	// array of numeric observation values (featureArray), or any combination of these.
	FeatureString param.Opt[string] `json:"featureString,omitzero"`
	// A single feature/assessment value expressed in the specified unit of measure
	// (obUoM). Because of the variability of the Feature Assessment data types, each
	// record may employ a numeric observation value (featureValue), a string
	// observation value (featureString), a Boolean observation value (featureBool), an
	// array of numeric observation values (featureArray), or any combination of these.
	FeatureValue param.Opt[float64] `json:"featureValue,omitzero"`
	// The feature object heading, in degrees clockwise from true North at the object
	// location.
	Heading param.Opt[float64] `json:"heading,omitzero"`
	// Estimated physical height of the feature, in meters.
	Height param.Opt[float64] `json:"height,omitzero"`
	// Estimated physical length of the feature, in meters.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Source defined name of the feature associated with this record. If an annotation
	// for this feature element exists on the parent image it can be referenced here.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Feature's speed of travel, in meters per second.
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The type of feature (e.g. AIRCRAFT, ANTENNA, SOLAR ARRAY, SITE, STRUCTURE,
	// VESSEL, VEHICLE, UNKNOWN, etc.) detailed in this feature assessment record. This
	// type may be a primary feature within an image, for example a VESSEL, or may be a
	// component or characteristic of a primary feature, for example an ANTENNA mounted
	// on a vessel.
	Type param.Opt[string] `json:"type,omitzero"`
	// Estimated physical width of the feature, in meters.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Polygonal annotation limits, specified in pixels, as an array of arrays N x M.
	// Allows the image provider to highlight one or more polygonal area(s) of
	// interest. The array must contain NxM two-element arrays, where N is the number
	// of polygons of interest. The associated annotation(s) should be included in the
	// annText array.
	AnnLims [][]int64 `json:"annLims,omitzero"`
	// Annotation text, a string array of annotation(s) corresponding to the
	// rectangular areas specified in annLims. This array contains the annotation text
	// associated with the areas of interest indicated in annLims, in order. This array
	// should contain one annotation per four values of the area (annLims) array.
	AnnText []string `json:"annText,omitzero"`
	// An array of numeric feature/assessment values expressed in the specified unit of
	// measure (obUoM). Because of the variability of the Feature Assessment data
	// types, each record may employ a numeric observation value (featureValue), a
	// string observation value (featureString), a Boolean observation value
	// (featureBool), an array of numeric observation values (featureArray), or any
	// combination of these.
	FeatureArray []float64 `json:"featureArray,omitzero"`
	// An array of string feature/assessment expressions. Because of the variability of
	// the Feature Assessment data types, each record may employ a numeric observation
	// value (featureValue), a string observation value (featureString), a Boolean
	// observation value (featureBool), an array of numeric observation values
	// (featureArray), or any combination of these.
	FeatureStringArray []string `json:"featureStringArray,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this activity or event. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps', 'srcIds', and 'srcTs' arrays must contain the same number of
	// elements. See the corresponding srcTyps array element for the data type of the
	// UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of the primary timestamps, in ISO 8601 UTC format, with appropriate
	// precision for the datatype of each correspondng srcTyp/srcId record. See the
	// associated 'srcTyps' and 'srcIds' arrays for the record type and UUID,
	// respectively, positionally corresponding to the record types in this array. The
	// 'srcTyps', 'srcIds', and 'srcTs' arrays must contain the same number of
	// elements. These timestamps are included to support services which do not include
	// a GET by {id} operation. If referencing a datatype which does not include a
	// primary timestamp, the corresponding srcTs array element should be included as
	// null.
	SrcTs []time.Time `json:"srcTs,omitzero" format:"date-time"`
	// Array of UDL record types (AIS, GROUNDIMAGE, MTI, ONORBIT, POI, SAR, SKYIMAGE,
	// SOI, TRACK) related to this feature assessment. See the associated 'srcIds' and
	// 'srcTs' arrays for the record UUIDs and timetsmps. respectively, positionally
	// corresponding to the record types in this array. The 'srcTyps', 'srcIds', and
	// 'srcTs' arrays must contain the same number of elements.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FeatureAssessmentUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow FeatureAssessmentUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeatureAssessmentUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FeatureAssessmentUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
