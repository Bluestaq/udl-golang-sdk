// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// OnorbitassessmentService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitassessmentService] method instead.
type OnorbitassessmentService struct {
	Options []option.RequestOption
	History OnorbitassessmentHistoryService
}

// NewOnorbitassessmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOnorbitassessmentService(opts ...option.RequestOption) (r OnorbitassessmentService) {
	r = OnorbitassessmentService{}
	r.Options = opts
	r.History = NewOnorbitassessmentHistoryService(opts...)
	return
}

// Service operation to take a single OnorbitAssessment as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *OnorbitassessmentService) New(ctx context.Context, body OnorbitassessmentNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitassessment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitassessmentService) List(ctx context.Context, query OnorbitassessmentListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitassessmentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitassessment"
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
func (r *OnorbitassessmentService) ListAutoPaging(ctx context.Context, query OnorbitassessmentListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitassessmentListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnorbitassessmentService) Count(ctx context.Context, query OnorbitassessmentCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onorbitassessment/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// OnorbitAssessment records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *OnorbitassessmentService) NewBulk(ctx context.Context, body OnorbitassessmentNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitassessment/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single OnorbitAssessment record by its unique ID
// passed as a path parameter.
func (r *OnorbitassessmentService) Get(ctx context.Context, id string, query OnorbitassessmentGetParams, opts ...option.RequestOption) (res *OnorbitassessmentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitassessment/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *OnorbitassessmentService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *OnorbitassessmentQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/onorbitassessment/queryhelp"
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
func (r *OnorbitassessmentService) Tuple(ctx context.Context, query OnorbitassessmentTupleParams, opts ...option.RequestOption) (res *[]OnorbitassessmentTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/onorbitassessment/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple OnorbitAssessment records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *OnorbitassessmentService) UnvalidatedPublish(ctx context.Context, body OnorbitassessmentUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-onorbitassessment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Spacecraft characterization results from analysis of MASINT data. Supports
// results of both the overall assessment of a particular spacecraft as well as of
// individual signature packages.
type OnorbitassessmentListResponse struct {
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
	DataMode OnorbitassessmentListResponseDataMode `json:"dataMode,required"`
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
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		SatNo                 respjson.Field
		SigDataType           respjson.Field
		SourceDl              respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitassessmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitassessmentListResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitassessmentListResponseDataMode string

const (
	OnorbitassessmentListResponseDataModeReal      OnorbitassessmentListResponseDataMode = "REAL"
	OnorbitassessmentListResponseDataModeTest      OnorbitassessmentListResponseDataMode = "TEST"
	OnorbitassessmentListResponseDataModeSimulated OnorbitassessmentListResponseDataMode = "SIMULATED"
	OnorbitassessmentListResponseDataModeExercise  OnorbitassessmentListResponseDataMode = "EXERCISE"
)

// Spacecraft characterization results from analysis of MASINT data. Supports
// results of both the overall assessment of a particular spacecraft as well as of
// individual signature packages.
type OnorbitassessmentGetResponse struct {
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
	DataMode OnorbitassessmentGetResponseDataMode `json:"dataMode,required"`
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
func (r OnorbitassessmentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitassessmentGetResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitassessmentGetResponseDataMode string

const (
	OnorbitassessmentGetResponseDataModeReal      OnorbitassessmentGetResponseDataMode = "REAL"
	OnorbitassessmentGetResponseDataModeTest      OnorbitassessmentGetResponseDataMode = "TEST"
	OnorbitassessmentGetResponseDataModeSimulated OnorbitassessmentGetResponseDataMode = "SIMULATED"
	OnorbitassessmentGetResponseDataModeExercise  OnorbitassessmentGetResponseDataMode = "EXERCISE"
)

type OnorbitassessmentQueryhelpResponse struct {
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
func (r OnorbitassessmentQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitassessmentQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spacecraft characterization results from analysis of MASINT data. Supports
// results of both the overall assessment of a particular spacecraft as well as of
// individual signature packages.
type OnorbitassessmentTupleResponse struct {
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
	DataMode OnorbitassessmentTupleResponseDataMode `json:"dataMode,required"`
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
func (r OnorbitassessmentTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitassessmentTupleResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitassessmentTupleResponseDataMode string

const (
	OnorbitassessmentTupleResponseDataModeReal      OnorbitassessmentTupleResponseDataMode = "REAL"
	OnorbitassessmentTupleResponseDataModeTest      OnorbitassessmentTupleResponseDataMode = "TEST"
	OnorbitassessmentTupleResponseDataModeSimulated OnorbitassessmentTupleResponseDataMode = "SIMULATED"
	OnorbitassessmentTupleResponseDataModeExercise  OnorbitassessmentTupleResponseDataMode = "EXERCISE"
)

type OnorbitassessmentNewParams struct {
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
	DataMode OnorbitassessmentNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The overall assessment of the on-orbit object. The assessment will vary
	// depending on the assessment level. If assmtLevel = SATELLITE, then expected
	// values include HEALTHY, NO DATA, UNHEALTHY, and UNKNOWN. If assmtLevel =
	// SIGNATURE, then expected values include ANOMALOUS, BAD, NOMINAL, and UNKNOWN.
	Assessment param.Opt[string] `json:"assessment,omitzero"`
	// Comments and details concerning this assessment.
	AssmtDetails param.Opt[string] `json:"assmtDetails,omitzero"`
	// The level (SATELLITE, SIGNATURE) of this assessment.
	AssmtLevel param.Opt[string] `json:"assmtLevel,omitzero"`
	// The rotational period, in seconds, determined in the assessment of the on-orbit
	// object.
	AssmtRotPeriod param.Opt[float64] `json:"assmtRotPeriod,omitzero"`
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
	AssmtSubType param.Opt[string] `json:"assmtSubType,omitzero"`
	// URL to an external location containing additional assessment information.
	AssmtURL param.Opt[string] `json:"assmtURL,omitzero"`
	// Flag indicating whether this assessment was auto-generated (true) or produced by
	// an analyst.
	AutoAssmt param.Opt[bool] `json:"autoAssmt,omitzero"`
	// URL to an external location containing the data that was used to make this
	// assessment.
	CollectionURL param.Opt[string] `json:"collectionURL,omitzero"`
	// Unique identifier of the target satellite on-orbit object to which this
	// assessment applies. This ID can be used to obtain additional information on an
	// OnOrbit object using the 'get by ID' operation (e.g. /udl/onorbit/{id}). For
	// example, the OnOrbit with idOnOrbit = 25544 would be queried as
	// /udl/onorbit/25544.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the sensor from which the signature data applied in this
	// assessment was collected. This ID can be used to obtain additional information
	// on a sensor using the 'get by ID' operation (e.g. /udl/sensor/{id}). For
	// example, the sensor with idSensor = abc would be queried as /udl/sensor/abc.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The total duration, in hours, of the signature or set of signatures used to
	// create this assessment.
	ObDuration param.Opt[float64] `json:"obDuration,omitzero"`
	// The observation time, in ISO 8601 UTC format with millisecond precision. For
	// non-instantaneous collections, the observation time will correspond to the
	// earliest timestamp in the signature or signature set data.
	ObTime param.Opt[time.Time] `json:"obTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the source to indicate the target on-orbit
	// object to which this assessment applies. This may be an internal identifier and
	// not necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Satellite/Catalog number of the target on-orbit object to which this assessment
	// applies.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The observation data type (LONG DWELL, NARROWBAND, PHOTOMETRIC POL, PHOTOMETRY,
	// RANGE PROFILER, WIDEBAND, etc.) contained in the signature. Applies when
	// assmtLevel = SIGNATURE.
	SigDataType param.Opt[string] `json:"sigDataType,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Array of the affected spacecraft component(s) relevant to this assessment, if
	// non-nominal.
	Components []string `json:"components,omitzero"`
	// Array of UUIDs of the UDL data records that are related to this assessment. See
	// the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (DOA, ELSET, EO, ESID, GROUNDIMAGE, POI, MANEUVER,
	// MTI, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK, etc.) that are related to this
	// assessment. See the associated 'srcIds' array for the record UUIDs, positionally
	// corresponding to the record types in this array. The 'srcTyps' and 'srcIds'
	// arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r OnorbitassessmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitassessmentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitassessmentNewParams) UnmarshalJSON(data []byte) error {
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
type OnorbitassessmentNewParamsDataMode string

const (
	OnorbitassessmentNewParamsDataModeReal      OnorbitassessmentNewParamsDataMode = "REAL"
	OnorbitassessmentNewParamsDataModeTest      OnorbitassessmentNewParamsDataMode = "TEST"
	OnorbitassessmentNewParamsDataModeSimulated OnorbitassessmentNewParamsDataMode = "SIMULATED"
	OnorbitassessmentNewParamsDataModeExercise  OnorbitassessmentNewParamsDataMode = "EXERCISE"
)

type OnorbitassessmentListParams struct {
	// The time of the assessment, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	AssmtTime   time.Time        `query:"assmtTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitassessmentListParams]'s query parameters as
// `url.Values`.
func (r OnorbitassessmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitassessmentCountParams struct {
	// The time of the assessment, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	AssmtTime   time.Time        `query:"assmtTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitassessmentCountParams]'s query parameters as
// `url.Values`.
func (r OnorbitassessmentCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitassessmentNewBulkParams struct {
	Body []OnorbitassessmentNewBulkParamsBody
	paramObj
}

func (r OnorbitassessmentNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *OnorbitassessmentNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Spacecraft characterization results from analysis of MASINT data. Supports
// results of both the overall assessment of a particular spacecraft as well as of
// individual signature packages.
//
// The properties AssmtTime, ClassificationMarking, DataMode, Source are required.
type OnorbitassessmentNewBulkParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The overall assessment of the on-orbit object. The assessment will vary
	// depending on the assessment level. If assmtLevel = SATELLITE, then expected
	// values include HEALTHY, NO DATA, UNHEALTHY, and UNKNOWN. If assmtLevel =
	// SIGNATURE, then expected values include ANOMALOUS, BAD, NOMINAL, and UNKNOWN.
	Assessment param.Opt[string] `json:"assessment,omitzero"`
	// Comments and details concerning this assessment.
	AssmtDetails param.Opt[string] `json:"assmtDetails,omitzero"`
	// The level (SATELLITE, SIGNATURE) of this assessment.
	AssmtLevel param.Opt[string] `json:"assmtLevel,omitzero"`
	// The rotational period, in seconds, determined in the assessment of the on-orbit
	// object.
	AssmtRotPeriod param.Opt[float64] `json:"assmtRotPeriod,omitzero"`
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
	AssmtSubType param.Opt[string] `json:"assmtSubType,omitzero"`
	// URL to an external location containing additional assessment information.
	AssmtURL param.Opt[string] `json:"assmtURL,omitzero"`
	// Flag indicating whether this assessment was auto-generated (true) or produced by
	// an analyst.
	AutoAssmt param.Opt[bool] `json:"autoAssmt,omitzero"`
	// URL to an external location containing the data that was used to make this
	// assessment.
	CollectionURL param.Opt[string] `json:"collectionURL,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of the target satellite on-orbit object to which this
	// assessment applies. This ID can be used to obtain additional information on an
	// OnOrbit object using the 'get by ID' operation (e.g. /udl/onorbit/{id}). For
	// example, the OnOrbit with idOnOrbit = 25544 would be queried as
	// /udl/onorbit/25544.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the sensor from which the signature data applied in this
	// assessment was collected. This ID can be used to obtain additional information
	// on a sensor using the 'get by ID' operation (e.g. /udl/sensor/{id}). For
	// example, the sensor with idSensor = abc would be queried as /udl/sensor/abc.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The total duration, in hours, of the signature or set of signatures used to
	// create this assessment.
	ObDuration param.Opt[float64] `json:"obDuration,omitzero"`
	// The observation time, in ISO 8601 UTC format with millisecond precision. For
	// non-instantaneous collections, the observation time will correspond to the
	// earliest timestamp in the signature or signature set data.
	ObTime param.Opt[time.Time] `json:"obTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the source to indicate the target on-orbit
	// object to which this assessment applies. This may be an internal identifier and
	// not necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Satellite/Catalog number of the target on-orbit object to which this assessment
	// applies.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The observation data type (LONG DWELL, NARROWBAND, PHOTOMETRIC POL, PHOTOMETRY,
	// RANGE PROFILER, WIDEBAND, etc.) contained in the signature. Applies when
	// assmtLevel = SIGNATURE.
	SigDataType param.Opt[string] `json:"sigDataType,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Array of the affected spacecraft component(s) relevant to this assessment, if
	// non-nominal.
	Components []string `json:"components,omitzero"`
	// Array of UUIDs of the UDL data records that are related to this assessment. See
	// the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (DOA, ELSET, EO, ESID, GROUNDIMAGE, POI, MANEUVER,
	// MTI, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK, etc.) that are related to this
	// assessment. See the associated 'srcIds' array for the record UUIDs, positionally
	// corresponding to the record types in this array. The 'srcTyps' and 'srcIds'
	// arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r OnorbitassessmentNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitassessmentNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitassessmentNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OnorbitassessmentNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type OnorbitassessmentGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitassessmentGetParams]'s query parameters as
// `url.Values`.
func (r OnorbitassessmentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitassessmentTupleParams struct {
	// The time of the assessment, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	AssmtTime time.Time `query:"assmtTime,required" format:"date-time" json:"-"`
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitassessmentTupleParams]'s query parameters as
// `url.Values`.
func (r OnorbitassessmentTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitassessmentUnvalidatedPublishParams struct {
	Body []OnorbitassessmentUnvalidatedPublishParamsBody
	paramObj
}

func (r OnorbitassessmentUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *OnorbitassessmentUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Spacecraft characterization results from analysis of MASINT data. Supports
// results of both the overall assessment of a particular spacecraft as well as of
// individual signature packages.
//
// The properties AssmtTime, ClassificationMarking, DataMode, Source are required.
type OnorbitassessmentUnvalidatedPublishParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The overall assessment of the on-orbit object. The assessment will vary
	// depending on the assessment level. If assmtLevel = SATELLITE, then expected
	// values include HEALTHY, NO DATA, UNHEALTHY, and UNKNOWN. If assmtLevel =
	// SIGNATURE, then expected values include ANOMALOUS, BAD, NOMINAL, and UNKNOWN.
	Assessment param.Opt[string] `json:"assessment,omitzero"`
	// Comments and details concerning this assessment.
	AssmtDetails param.Opt[string] `json:"assmtDetails,omitzero"`
	// The level (SATELLITE, SIGNATURE) of this assessment.
	AssmtLevel param.Opt[string] `json:"assmtLevel,omitzero"`
	// The rotational period, in seconds, determined in the assessment of the on-orbit
	// object.
	AssmtRotPeriod param.Opt[float64] `json:"assmtRotPeriod,omitzero"`
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
	AssmtSubType param.Opt[string] `json:"assmtSubType,omitzero"`
	// URL to an external location containing additional assessment information.
	AssmtURL param.Opt[string] `json:"assmtURL,omitzero"`
	// Flag indicating whether this assessment was auto-generated (true) or produced by
	// an analyst.
	AutoAssmt param.Opt[bool] `json:"autoAssmt,omitzero"`
	// URL to an external location containing the data that was used to make this
	// assessment.
	CollectionURL param.Opt[string] `json:"collectionURL,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of the target satellite on-orbit object to which this
	// assessment applies. This ID can be used to obtain additional information on an
	// OnOrbit object using the 'get by ID' operation (e.g. /udl/onorbit/{id}). For
	// example, the OnOrbit with idOnOrbit = 25544 would be queried as
	// /udl/onorbit/25544.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the sensor from which the signature data applied in this
	// assessment was collected. This ID can be used to obtain additional information
	// on a sensor using the 'get by ID' operation (e.g. /udl/sensor/{id}). For
	// example, the sensor with idSensor = abc would be queried as /udl/sensor/abc.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The total duration, in hours, of the signature or set of signatures used to
	// create this assessment.
	ObDuration param.Opt[float64] `json:"obDuration,omitzero"`
	// The observation time, in ISO 8601 UTC format with millisecond precision. For
	// non-instantaneous collections, the observation time will correspond to the
	// earliest timestamp in the signature or signature set data.
	ObTime param.Opt[time.Time] `json:"obTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the source to indicate the target on-orbit
	// object to which this assessment applies. This may be an internal identifier and
	// not necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the observation source to indicate the sensor
	// which produced this observation. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Satellite/Catalog number of the target on-orbit object to which this assessment
	// applies.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The observation data type (LONG DWELL, NARROWBAND, PHOTOMETRIC POL, PHOTOMETRY,
	// RANGE PROFILER, WIDEBAND, etc.) contained in the signature. Applies when
	// assmtLevel = SIGNATURE.
	SigDataType param.Opt[string] `json:"sigDataType,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Array of the affected spacecraft component(s) relevant to this assessment, if
	// non-nominal.
	Components []string `json:"components,omitzero"`
	// Array of UUIDs of the UDL data records that are related to this assessment. See
	// the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (DOA, ELSET, EO, ESID, GROUNDIMAGE, POI, MANEUVER,
	// MTI, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK, etc.) that are related to this
	// assessment. See the associated 'srcIds' array for the record UUIDs, positionally
	// corresponding to the record types in this array. The 'srcTyps' and 'srcIds'
	// arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r OnorbitassessmentUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitassessmentUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitassessmentUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OnorbitassessmentUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
