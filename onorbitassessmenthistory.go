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
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// OnorbitassessmentHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitassessmentHistoryService] method instead.
type OnorbitassessmentHistoryService struct {
	Options []option.RequestOption
}

// NewOnorbitassessmentHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOnorbitassessmentHistoryService(opts ...option.RequestOption) (r OnorbitassessmentHistoryService) {
	r = OnorbitassessmentHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitassessmentHistoryService) List(ctx context.Context, query OnorbitassessmentHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitassessmentHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitassessment/history"
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
func (r *OnorbitassessmentHistoryService) ListAutoPaging(ctx context.Context, query OnorbitassessmentHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitassessmentHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitassessmentHistoryService) Aodr(ctx context.Context, query OnorbitassessmentHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitassessment/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnorbitassessmentHistoryService) Count(ctx context.Context, query OnorbitassessmentHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onorbitassessment/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Spacecraft characterization results from analysis of MASINT data. Supports
// results of both the overall assessment of a particular spacecraft as well as of
// individual signature packages.
type OnorbitassessmentHistoryListResponse struct {
	// The time of the assessment, in ISO 8601 UTC format with millisecond precision.
	AssmtTime time.Time `json:"assmtTime,required" format:"date-time"`
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
	DataMode OnorbitassessmentHistoryListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The overall assessment of the on-orbit object. The assessment will vary
	// depending on the assessment level. If assmtLevel = SATELLITE, then expected
	// values include HEALTHY, NO DATA, UNHEALTHY, and UNKNOWN. If assmtLevel =
	// SIGNATURE, then expected values include ANOMALOUS, BAD, NOMINAL, and UNKNOWN.
	Assessment string `json:"assessment"`
	// Comments and details concerning this assessment.
	AssmtDetails string `json:"assmtDetails"`
	// The level (SATELLITE, SIGNATURE) of this assessment.
	AssmtLevel string `json:"assmtLevel"`
	// The rotational period, in seconds, determined in the assessment of the on-orbit
	// object.
	AssmtRotPeriod float64 `json:"assmtRotPeriod"`
	// The sub-type for this assessment. The sub-type provides an additional level of
	// specification when the assessment level = SIGNATURE, and will vary based on the
	// overall assessment. If assessment = ANOMALOUS, then expected values include
	// ACTIVITY OBSERVED, BAD CONFIGURATION, MANEUVERING, OTHER, POSSIBLE SAFE MODE,
	// UNSTABLE, and WRONG ATTITUDE. If assessment = BAD, then expected values include
	// BAD-MISSING STATE VECTOR, CORRUPT-NOISY, CROSS-TAG, DROPOUTS, INSUFFICIENT DATA,
	// INSUFFICIENT LOOK ANGLE, NO CROSSOVER, and SHORT. If assessment = NOMINAL, then
	// expected values include ACTIVITY OBSERVED, GRAVITY GRADIENT, HORIZON STABLE,
	// INERTIAL, MANEUVERING, SPIN STABLE, and STABLE. If assessment = UNKNOWN, then
	// expected values include NO COHORT, and OTHER.
	AssmtSubType string `json:"assmtSubType"`
	// URL to an external location containing additional assessment information.
	AssmtURL string `json:"assmtURL"`
	// Flag indicating whether this assessment was auto-generated (true) or produced by
	// an analyst.
	AutoAssmt bool `json:"autoAssmt"`
	// URL to an external location containing the data that was used to make this
	// assessment.
	CollectionURL string `json:"collectionURL"`
	// Array of the affected spacecraft component(s) relevant to this assessment, if
	// non-nominal.
	Components []string `json:"components"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the target satellite on-orbit object to which this
	// assessment applies. This ID can be used to obtain additional information on an
	// OnOrbit object using the 'get by ID' operation (e.g. /udl/onorbit/{id}). For
	// example, the OnOrbit with idOnOrbit = 25544 would be queried as
	// /udl/onorbit/25544.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the sensor from which the signature data applied in this
	// assessment was collected. This ID can be used to obtain additional information
	// on a sensor using the 'get by ID' operation (e.g. /udl/sensor/{id}). For
	// example, the sensor with idSensor = abc would be queried as /udl/sensor/abc.
	IDSensor string `json:"idSensor"`
	// The total duration, in hours, of the signature or set of signatures used to
	// create this assessment.
	ObDuration float64 `json:"obDuration"`
	// The observation time, in ISO 8601 UTC format with millisecond precision. For
	// non-instantaneous collections, the observation time will correspond to the
	// earliest timestamp in the signature or signature set data.
	ObTime time.Time `json:"obTime" format:"date-time"`
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
	// Optional identifier provided by the source to indicate the target on-orbit
	// object to which this assessment applies. This may be an internal identifier and
	// not necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Satellite/Catalog number of the target on-orbit object to which this assessment
	// applies.
	SatNo int64 `json:"satNo"`
	// The observation data type (LONG DWELL, NARROWBAND, PHOTOMETRIC POL, PHOTOMETRY,
	// RANGE PROFILER, WIDEBAND, etc.) contained in the signature. Applies when
	// assmtLevel = SIGNATURE.
	SigDataType string `json:"sigDataType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Array of UUIDs of the UDL data records that are related to this assessment. See
	// the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (DOA, ELSET, EO, ESID, GROUNDIMAGE, POI, MANEUVER,
	// MTI, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK, etc.) that are related to this
	// assessment. See the associated 'srcIds' array for the record UUIDs, positionally
	// corresponding to the record types in this array. The 'srcTyps' and 'srcIds'
	// arrays must match in size.
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssmtTime             respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Assessment            respjson.Field
		AssmtDetails          respjson.Field
		AssmtLevel            respjson.Field
		AssmtRotPeriod        respjson.Field
		AssmtSubType          respjson.Field
		AssmtURL              respjson.Field
		AutoAssmt             respjson.Field
		CollectionURL         respjson.Field
		Components            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		ObDuration            respjson.Field
		ObTime                respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		SatNo                 respjson.Field
		SigDataType           respjson.Field
		SourceDl              respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitassessmentHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitassessmentHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitassessmentHistoryListResponseDataMode string

const (
	OnorbitassessmentHistoryListResponseDataModeReal      OnorbitassessmentHistoryListResponseDataMode = "REAL"
	OnorbitassessmentHistoryListResponseDataModeTest      OnorbitassessmentHistoryListResponseDataMode = "TEST"
	OnorbitassessmentHistoryListResponseDataModeSimulated OnorbitassessmentHistoryListResponseDataMode = "SIMULATED"
	OnorbitassessmentHistoryListResponseDataModeExercise  OnorbitassessmentHistoryListResponseDataMode = "EXERCISE"
)

type OnorbitassessmentHistoryListParams struct {
	// The time of the assessment, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	AssmtTime time.Time `query:"assmtTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitassessmentHistoryListParams]'s query parameters as
// `url.Values`.
func (r OnorbitassessmentHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitassessmentHistoryAodrParams struct {
	// The time of the assessment, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	AssmtTime time.Time `query:"assmtTime,required" format:"date-time" json:"-"`
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

// URLQuery serializes [OnorbitassessmentHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r OnorbitassessmentHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitassessmentHistoryCountParams struct {
	// The time of the assessment, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	AssmtTime   time.Time        `query:"assmtTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitassessmentHistoryCountParams]'s query parameters as
// `url.Values`.
func (r OnorbitassessmentHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
