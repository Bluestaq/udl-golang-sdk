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
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// NavigationalObstructionService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNavigationalObstructionService] method instead.
type NavigationalObstructionService struct {
	Options []option.RequestOption
}

// NewNavigationalObstructionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNavigationalObstructionService(opts ...option.RequestOption) (r NavigationalObstructionService) {
	r = NavigationalObstructionService{}
	r.Options = opts
	return
}

// Service operation to take a single navigational obstruction record as a POST
// body and ingest into the database. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *NavigationalObstructionService) New(ctx context.Context, body NavigationalObstructionNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/navigationalobstruction"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single navigational obstruction record. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *NavigationalObstructionService) Update(ctx context.Context, id string, body NavigationalObstructionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/navigationalobstruction/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *NavigationalObstructionService) List(ctx context.Context, query NavigationalObstructionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[NavigationalObstructionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/navigationalobstruction"
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
func (r *NavigationalObstructionService) ListAutoPaging(ctx context.Context, query NavigationalObstructionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[NavigationalObstructionListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *NavigationalObstructionService) Count(ctx context.Context, query NavigationalObstructionCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/navigationalobstruction/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// navigational obstruction records as a POST body and ingest into the database.
// This operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *NavigationalObstructionService) NewBulk(ctx context.Context, body NavigationalObstructionNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/navigationalobstruction/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single navigational obstruction record by its unique
// ID passed as a path parameter.
func (r *NavigationalObstructionService) Get(ctx context.Context, id string, query NavigationalObstructionGetParams, opts ...option.RequestOption) (res *NavigationalObstructionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/navigationalobstruction/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *NavigationalObstructionService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *NavigationalObstructionQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/navigationalobstruction/queryhelp"
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
func (r *NavigationalObstructionService) Tuple(ctx context.Context, query NavigationalObstructionTupleParams, opts ...option.RequestOption) (res *[]NavigationalObstructionTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/navigationalobstruction/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Beta Version Navigational Obstruction: Information describing navigational
// obstructions, such as applicable boundaries, locations, heights, data ownership,
// and currency.
type NavigationalObstructionListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Start date of this obstruction data set's currency, in ISO 8601 date-only
	// format.
	CycleDate time.Time `json:"cycleDate,required" format:"date"`
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
	DataMode NavigationalObstructionListResponseDataMode `json:"dataMode,required"`
	// The ID of this obstacle.
	ObstacleID string `json:"obstacleId,required"`
	// Type of obstacle (e.g. P for point, V for vector, L for line).
	ObstacleType string `json:"obstacleType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Indicates if this obstacle record is Active (A) or Deleted (D).
	ActDelCode string `json:"actDelCode"`
	// The Aeronautical Information Regulation and Control (AIRAC) cycle of this
	// obstruction data set. The format is YYNN where YY is the last two digits of the
	// year and NN is the cycle number.
	AiracCycle int64 `json:"airacCycle"`
	// The baseline Aeronautical Information Regulation and Control (AIRAC) cycle for
	// change sets only. The format is YYNN where YY is the last two digits of the year
	// and NN is the cycle number.
	BaseAiracCycle int64 `json:"baseAiracCycle"`
	// Earliest record date possible in this obstruction data set (not the earliest
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD). If null, this data
	// set is assumed to be a full data pull of holdings until the cutoffDate. If this
	// field is populated, this data set only contains updates since the last baseline
	// data set.
	BaselineCutoffDate time.Time `json:"baselineCutoffDate" format:"date"`
	// WGS-84 latitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundNeLat float64 `json:"boundNELat"`
	// WGS-84 longitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundNeLon float64 `json:"boundNELon"`
	// WGS-84 latitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundSwLat float64 `json:"boundSWLat"`
	// WGS-84 longitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundSwLon float64 `json:"boundSWLon"`
	// The DoD Standard Country Code designator for the country issuing the diplomatic
	// clearance. This field will be set to "OTHR" if the source value does not match a
	// UDL Country code value (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Latest record date possible in this obstruction data set (not the most recent
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	CutoffDate time.Time `json:"cutoffDate" format:"date"`
	// Remarks concerning this obstruction's data set.
	DataSetRemarks string `json:"dataSetRemarks"`
	// The organization that deleted this obstacle record.
	DeletingOrg string `json:"deletingOrg"`
	// The organization that entered obstacle data other than the producer.
	DerivingOrg string `json:"derivingOrg"`
	// The side or sides of this obstruction feature which produces the greatest
	// reflectivity potential.
	DirectivityCode int64 `json:"directivityCode"`
	// The elevation at the point obstacle's location in feet.
	Elevation float64 `json:"elevation"`
	// The difference between the assigned elevation of this point and its true
	// elevation, in feet.
	ElevationAcc float64 `json:"elevationAcc"`
	// Optional obstacle ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID string `json:"externalId"`
	// FACC (Feature and Attribute Coding Catalog) is a five-character code for
	// encoding real-world entities and objects. The first letter of the code is an
	// alphabetic value from "A" to "Z" which will map to a feature category. The
	// second character will map to a subcategory. Characters three to five are numeric
	// and range from 000 to 999. This value will provide a unit feature type
	// identification within the categories.
	Facc string `json:"facc"`
	// Identifying code for the type of this point obstacle.
	FeatureCode string `json:"featureCode"`
	// Description of this obstacle, corresponding to the FACC (Feature and Attribute
	// Coding Catalog) value.
	FeatureDescription string `json:"featureDescription"`
	// Type name of point obstacle.
	FeatureName string `json:"featureName"`
	// Identifying code for the type of this point obstacle.
	FeatureType string `json:"featureType"`
	// The height Above Ground Level (AGL) of the point obstacle in feet.
	HeightAgl float64 `json:"heightAGL"`
	// The accuracy of the height Above Ground Level (AGL) value for this point
	// obstacle, in feet.
	HeightAglAcc float64 `json:"heightAGLAcc"`
	// The height Above Mean Sea Level (AMSL) of the point obstacle in feet.
	HeightMsl float64 `json:"heightMSL"`
	// The accuracy of the height Above Mean Sea Level (AMSL) value for this point
	// obstacle in feet.
	HeightMslAcc float64 `json:"heightMSLAcc"`
	// The difference between the recorded horizontal coordinates of this point
	// obstacle and its true position, in feet.
	HorizAcc float64 `json:"horizAcc"`
	// Code representing the mathematical model of Earth used to calculate coordinates
	// for this obstacle (e.g. WGS-84, U for undetermined, etc.). US Forces use the
	// World Geodetic System 1984 (WGS-84), but also use maps by allied countries with
	// local datums.
	HorizDatumCode string `json:"horizDatumCode"`
	// Date this obstacle was initially added to the data set, in ISO 8601 date-only
	// format (ex. YYYY-MM-DD).
	InitRecordDate time.Time `json:"initRecordDate" format:"date"`
	// This field provides an array of keys that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the values array. This array
	// must be the same length as values.
	Keys []string `json:"keys"`
	// Code specifying if this obstacle is lit (e.g. Y = Yes, N = No, U = Unknown).
	LightingCode string `json:"lightingCode"`
	// WGS-84 latitude of the northeastern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineNeLat float64 `json:"lineNELat"`
	// WGS-84 longitude of the northeastern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineNeLon float64 `json:"lineNELon"`
	// The name of the line file associated with this obstruction data set.
	LinesFilename string `json:"linesFilename"`
	// WGS-84 latitude of the southwestern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineSwLat float64 `json:"lineSWLat"`
	// WGS-84 longitude of the southwestern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineSwLon float64 `json:"lineSWLon"`
	// The minimum height Above Ground Level (AGL) of the shortest obstruction
	// contained in this data set, in feet.
	MinHeightAgl float64 `json:"minHeightAGL"`
	// Indicates if the feature has multiple obstructions (e.g. S = Single, M =
	// Multiple, U = Undetermined).
	MultObs string `json:"multObs"`
	// The date after which this obstruction data set’s currency is stale and should be
	// refreshed, in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	NextCycleDate time.Time `json:"nextCycleDate" format:"date"`
	// The number of line features associated with this obstruction data set.
	NumLines int64 `json:"numLines"`
	// Indicates the number of obstructions associated with a feature.
	NumObs int64 `json:"numObs"`
	// The number of point features associated with this obstruction data set.
	NumPoints int64 `json:"numPoints"`
	// Remarks regarding this obstacle.
	ObstacleRemarks string `json:"obstacleRemarks"`
	// The original ID for this obstacle.
	OrigID string `json:"origId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The DoD Standard Country Code designator for the country or political entity
	// that owns the data set associated with this obstruction. This field will be set
	// to "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	OwnerCountryCode string `json:"ownerCountryCode"`
	// WGS-84 latitude of this point obstacle, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PointLat float64 `json:"pointLat"`
	// WGS-84 longitude of this point obstacle, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PointLon float64 `json:"pointLon"`
	// The name of the point file associated with this obstruction data set.
	PointsFilename string `json:"pointsFilename"`
	// Code denoting the action, review, or process that updated this obstacle.
	ProcessCode string `json:"processCode"`
	// Name of the agency that produced this obstruction data set.
	Producer string `json:"producer"`
	// The Federal Information Processing Standards (FIPS) state/province numeric code
	// of this obstacle's location.
	ProvinceCode string `json:"provinceCode"`
	// When horizontal and/or vertical accuracy requirements cannot be met because of
	// inadequate source material, this code indicates the quality of the data.
	Quality string `json:"quality"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Date this obstacle data was revised, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	RevDate time.Time `json:"revDate" format:"date"`
	// ID of the end point of a line segment.
	SegEndPoint int64 `json:"segEndPoint"`
	// Identifies the sequence number of a line segment.
	SegNum int64 `json:"segNum"`
	// ID of the starting point of a line segment.
	SegStartPoint int64 `json:"segStartPoint"`
	// Source date of this obstacle data, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	SourceDate time.Time `json:"sourceDate" format:"date"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The surface material composition code of this point obstacle.
	SurfaceMatCode string `json:"surfaceMatCode"`
	// The transaction type/code for this obstacle (e.g. "D", "N", "R", "S", "V", "X").
	TransactionCode string `json:"transactionCode"`
	// Method used to confirm the existence of this obstacle.
	ValidationCode int64 `json:"validationCode"`
	// This field provides an array of values that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the keys array. This array
	// must be the same length as keys.
	Values []string `json:"values"`
	// The name of the vector file associated with this obstruction data set.
	VectorsFilename string `json:"vectorsFilename"`
	// The World Aeronautical Chart (WAC) identifier for the area in which this
	// obstacle is located.
	Wac string `json:"wac"`
	// This obstacle's World Area Code installation number (WAC-INNR).
	WacInnr string `json:"wacINNR"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		CycleDate             respjson.Field
		DataMode              respjson.Field
		ObstacleID            respjson.Field
		ObstacleType          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ActDelCode            respjson.Field
		AiracCycle            respjson.Field
		BaseAiracCycle        respjson.Field
		BaselineCutoffDate    respjson.Field
		BoundNeLat            respjson.Field
		BoundNeLon            respjson.Field
		BoundSwLat            respjson.Field
		BoundSwLon            respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CutoffDate            respjson.Field
		DataSetRemarks        respjson.Field
		DeletingOrg           respjson.Field
		DerivingOrg           respjson.Field
		DirectivityCode       respjson.Field
		Elevation             respjson.Field
		ElevationAcc          respjson.Field
		ExternalID            respjson.Field
		Facc                  respjson.Field
		FeatureCode           respjson.Field
		FeatureDescription    respjson.Field
		FeatureName           respjson.Field
		FeatureType           respjson.Field
		HeightAgl             respjson.Field
		HeightAglAcc          respjson.Field
		HeightMsl             respjson.Field
		HeightMslAcc          respjson.Field
		HorizAcc              respjson.Field
		HorizDatumCode        respjson.Field
		InitRecordDate        respjson.Field
		Keys                  respjson.Field
		LightingCode          respjson.Field
		LineNeLat             respjson.Field
		LineNeLon             respjson.Field
		LinesFilename         respjson.Field
		LineSwLat             respjson.Field
		LineSwLon             respjson.Field
		MinHeightAgl          respjson.Field
		MultObs               respjson.Field
		NextCycleDate         respjson.Field
		NumLines              respjson.Field
		NumObs                respjson.Field
		NumPoints             respjson.Field
		ObstacleRemarks       respjson.Field
		OrigID                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OwnerCountryCode      respjson.Field
		PointLat              respjson.Field
		PointLon              respjson.Field
		PointsFilename        respjson.Field
		ProcessCode           respjson.Field
		Producer              respjson.Field
		ProvinceCode          respjson.Field
		Quality               respjson.Field
		RawFileUri            respjson.Field
		RevDate               respjson.Field
		SegEndPoint           respjson.Field
		SegNum                respjson.Field
		SegStartPoint         respjson.Field
		SourceDate            respjson.Field
		SourceDl              respjson.Field
		SurfaceMatCode        respjson.Field
		TransactionCode       respjson.Field
		ValidationCode        respjson.Field
		Values                respjson.Field
		VectorsFilename       respjson.Field
		Wac                   respjson.Field
		WacInnr               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationalObstructionListResponse) RawJSON() string { return r.JSON.raw }
func (r *NavigationalObstructionListResponse) UnmarshalJSON(data []byte) error {
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
type NavigationalObstructionListResponseDataMode string

const (
	NavigationalObstructionListResponseDataModeReal      NavigationalObstructionListResponseDataMode = "REAL"
	NavigationalObstructionListResponseDataModeTest      NavigationalObstructionListResponseDataMode = "TEST"
	NavigationalObstructionListResponseDataModeSimulated NavigationalObstructionListResponseDataMode = "SIMULATED"
	NavigationalObstructionListResponseDataModeExercise  NavigationalObstructionListResponseDataMode = "EXERCISE"
)

// Beta Version Navigational Obstruction: Information describing navigational
// obstructions, such as applicable boundaries, locations, heights, data ownership,
// and currency.
type NavigationalObstructionGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Start date of this obstruction data set's currency, in ISO 8601 date-only
	// format.
	CycleDate time.Time `json:"cycleDate,required" format:"date"`
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
	DataMode NavigationalObstructionGetResponseDataMode `json:"dataMode,required"`
	// The ID of this obstacle.
	ObstacleID string `json:"obstacleId,required"`
	// Type of obstacle (e.g. P for point, V for vector, L for line).
	ObstacleType string `json:"obstacleType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Indicates if this obstacle record is Active (A) or Deleted (D).
	ActDelCode string `json:"actDelCode"`
	// The Aeronautical Information Regulation and Control (AIRAC) cycle of this
	// obstruction data set. The format is YYNN where YY is the last two digits of the
	// year and NN is the cycle number.
	AiracCycle int64 `json:"airacCycle"`
	// The baseline Aeronautical Information Regulation and Control (AIRAC) cycle for
	// change sets only. The format is YYNN where YY is the last two digits of the year
	// and NN is the cycle number.
	BaseAiracCycle int64 `json:"baseAiracCycle"`
	// Earliest record date possible in this obstruction data set (not the earliest
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD). If null, this data
	// set is assumed to be a full data pull of holdings until the cutoffDate. If this
	// field is populated, this data set only contains updates since the last baseline
	// data set.
	BaselineCutoffDate time.Time `json:"baselineCutoffDate" format:"date"`
	// WGS-84 latitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundNeLat float64 `json:"boundNELat"`
	// WGS-84 longitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundNeLon float64 `json:"boundNELon"`
	// WGS-84 latitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundSwLat float64 `json:"boundSWLat"`
	// WGS-84 longitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundSwLon float64 `json:"boundSWLon"`
	// The DoD Standard Country Code designator for the country issuing the diplomatic
	// clearance. This field will be set to "OTHR" if the source value does not match a
	// UDL Country code value (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Latest record date possible in this obstruction data set (not the most recent
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	CutoffDate time.Time `json:"cutoffDate" format:"date"`
	// Remarks concerning this obstruction's data set.
	DataSetRemarks string `json:"dataSetRemarks"`
	// The organization that deleted this obstacle record.
	DeletingOrg string `json:"deletingOrg"`
	// The organization that entered obstacle data other than the producer.
	DerivingOrg string `json:"derivingOrg"`
	// The side or sides of this obstruction feature which produces the greatest
	// reflectivity potential.
	DirectivityCode int64 `json:"directivityCode"`
	// The elevation at the point obstacle's location in feet.
	Elevation float64 `json:"elevation"`
	// The difference between the assigned elevation of this point and its true
	// elevation, in feet.
	ElevationAcc float64 `json:"elevationAcc"`
	// Optional obstacle ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID string `json:"externalId"`
	// FACC (Feature and Attribute Coding Catalog) is a five-character code for
	// encoding real-world entities and objects. The first letter of the code is an
	// alphabetic value from "A" to "Z" which will map to a feature category. The
	// second character will map to a subcategory. Characters three to five are numeric
	// and range from 000 to 999. This value will provide a unit feature type
	// identification within the categories.
	Facc string `json:"facc"`
	// Identifying code for the type of this point obstacle.
	FeatureCode string `json:"featureCode"`
	// Description of this obstacle, corresponding to the FACC (Feature and Attribute
	// Coding Catalog) value.
	FeatureDescription string `json:"featureDescription"`
	// Type name of point obstacle.
	FeatureName string `json:"featureName"`
	// Identifying code for the type of this point obstacle.
	FeatureType string `json:"featureType"`
	// The height Above Ground Level (AGL) of the point obstacle in feet.
	HeightAgl float64 `json:"heightAGL"`
	// The accuracy of the height Above Ground Level (AGL) value for this point
	// obstacle, in feet.
	HeightAglAcc float64 `json:"heightAGLAcc"`
	// The height Above Mean Sea Level (AMSL) of the point obstacle in feet.
	HeightMsl float64 `json:"heightMSL"`
	// The accuracy of the height Above Mean Sea Level (AMSL) value for this point
	// obstacle in feet.
	HeightMslAcc float64 `json:"heightMSLAcc"`
	// The difference between the recorded horizontal coordinates of this point
	// obstacle and its true position, in feet.
	HorizAcc float64 `json:"horizAcc"`
	// Code representing the mathematical model of Earth used to calculate coordinates
	// for this obstacle (e.g. WGS-84, U for undetermined, etc.). US Forces use the
	// World Geodetic System 1984 (WGS-84), but also use maps by allied countries with
	// local datums.
	HorizDatumCode string `json:"horizDatumCode"`
	// Date this obstacle was initially added to the data set, in ISO 8601 date-only
	// format (ex. YYYY-MM-DD).
	InitRecordDate time.Time `json:"initRecordDate" format:"date"`
	// This field provides an array of keys that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the values array. This array
	// must be the same length as values.
	Keys []string `json:"keys"`
	// Code specifying if this obstacle is lit (e.g. Y = Yes, N = No, U = Unknown).
	LightingCode string `json:"lightingCode"`
	// WGS-84 latitude of the northeastern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineNeLat float64 `json:"lineNELat"`
	// WGS-84 longitude of the northeastern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineNeLon float64 `json:"lineNELon"`
	// The name of the line file associated with this obstruction data set.
	LinesFilename string `json:"linesFilename"`
	// WGS-84 latitude of the southwestern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineSwLat float64 `json:"lineSWLat"`
	// WGS-84 longitude of the southwestern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineSwLon float64 `json:"lineSWLon"`
	// The minimum height Above Ground Level (AGL) of the shortest obstruction
	// contained in this data set, in feet.
	MinHeightAgl float64 `json:"minHeightAGL"`
	// Indicates if the feature has multiple obstructions (e.g. S = Single, M =
	// Multiple, U = Undetermined).
	MultObs string `json:"multObs"`
	// The date after which this obstruction data set’s currency is stale and should be
	// refreshed, in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	NextCycleDate time.Time `json:"nextCycleDate" format:"date"`
	// The number of line features associated with this obstruction data set.
	NumLines int64 `json:"numLines"`
	// Indicates the number of obstructions associated with a feature.
	NumObs int64 `json:"numObs"`
	// The number of point features associated with this obstruction data set.
	NumPoints int64 `json:"numPoints"`
	// Remarks regarding this obstacle.
	ObstacleRemarks string `json:"obstacleRemarks"`
	// The original ID for this obstacle.
	OrigID string `json:"origId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The DoD Standard Country Code designator for the country or political entity
	// that owns the data set associated with this obstruction. This field will be set
	// to "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	OwnerCountryCode string `json:"ownerCountryCode"`
	// WGS-84 latitude of this point obstacle, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PointLat float64 `json:"pointLat"`
	// WGS-84 longitude of this point obstacle, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PointLon float64 `json:"pointLon"`
	// The name of the point file associated with this obstruction data set.
	PointsFilename string `json:"pointsFilename"`
	// Code denoting the action, review, or process that updated this obstacle.
	ProcessCode string `json:"processCode"`
	// Name of the agency that produced this obstruction data set.
	Producer string `json:"producer"`
	// The Federal Information Processing Standards (FIPS) state/province numeric code
	// of this obstacle's location.
	ProvinceCode string `json:"provinceCode"`
	// When horizontal and/or vertical accuracy requirements cannot be met because of
	// inadequate source material, this code indicates the quality of the data.
	Quality string `json:"quality"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Date this obstacle data was revised, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	RevDate time.Time `json:"revDate" format:"date"`
	// ID of the end point of a line segment.
	SegEndPoint int64 `json:"segEndPoint"`
	// Identifies the sequence number of a line segment.
	SegNum int64 `json:"segNum"`
	// ID of the starting point of a line segment.
	SegStartPoint int64 `json:"segStartPoint"`
	// Source date of this obstacle data, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	SourceDate time.Time `json:"sourceDate" format:"date"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The surface material composition code of this point obstacle.
	SurfaceMatCode string `json:"surfaceMatCode"`
	// The transaction type/code for this obstacle (e.g. "D", "N", "R", "S", "V", "X").
	TransactionCode string `json:"transactionCode"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Method used to confirm the existence of this obstacle.
	ValidationCode int64 `json:"validationCode"`
	// This field provides an array of values that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the keys array. This array
	// must be the same length as keys.
	Values []string `json:"values"`
	// The name of the vector file associated with this obstruction data set.
	VectorsFilename string `json:"vectorsFilename"`
	// The World Aeronautical Chart (WAC) identifier for the area in which this
	// obstacle is located.
	Wac string `json:"wac"`
	// This obstacle's World Area Code installation number (WAC-INNR).
	WacInnr string `json:"wacINNR"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		CycleDate             respjson.Field
		DataMode              respjson.Field
		ObstacleID            respjson.Field
		ObstacleType          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ActDelCode            respjson.Field
		AiracCycle            respjson.Field
		BaseAiracCycle        respjson.Field
		BaselineCutoffDate    respjson.Field
		BoundNeLat            respjson.Field
		BoundNeLon            respjson.Field
		BoundSwLat            respjson.Field
		BoundSwLon            respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CutoffDate            respjson.Field
		DataSetRemarks        respjson.Field
		DeletingOrg           respjson.Field
		DerivingOrg           respjson.Field
		DirectivityCode       respjson.Field
		Elevation             respjson.Field
		ElevationAcc          respjson.Field
		ExternalID            respjson.Field
		Facc                  respjson.Field
		FeatureCode           respjson.Field
		FeatureDescription    respjson.Field
		FeatureName           respjson.Field
		FeatureType           respjson.Field
		HeightAgl             respjson.Field
		HeightAglAcc          respjson.Field
		HeightMsl             respjson.Field
		HeightMslAcc          respjson.Field
		HorizAcc              respjson.Field
		HorizDatumCode        respjson.Field
		InitRecordDate        respjson.Field
		Keys                  respjson.Field
		LightingCode          respjson.Field
		LineNeLat             respjson.Field
		LineNeLon             respjson.Field
		LinesFilename         respjson.Field
		LineSwLat             respjson.Field
		LineSwLon             respjson.Field
		MinHeightAgl          respjson.Field
		MultObs               respjson.Field
		NextCycleDate         respjson.Field
		NumLines              respjson.Field
		NumObs                respjson.Field
		NumPoints             respjson.Field
		ObstacleRemarks       respjson.Field
		OrigID                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OwnerCountryCode      respjson.Field
		PointLat              respjson.Field
		PointLon              respjson.Field
		PointsFilename        respjson.Field
		ProcessCode           respjson.Field
		Producer              respjson.Field
		ProvinceCode          respjson.Field
		Quality               respjson.Field
		RawFileUri            respjson.Field
		RevDate               respjson.Field
		SegEndPoint           respjson.Field
		SegNum                respjson.Field
		SegStartPoint         respjson.Field
		SourceDate            respjson.Field
		SourceDl              respjson.Field
		SurfaceMatCode        respjson.Field
		TransactionCode       respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ValidationCode        respjson.Field
		Values                respjson.Field
		VectorsFilename       respjson.Field
		Wac                   respjson.Field
		WacInnr               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationalObstructionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NavigationalObstructionGetResponse) UnmarshalJSON(data []byte) error {
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
type NavigationalObstructionGetResponseDataMode string

const (
	NavigationalObstructionGetResponseDataModeReal      NavigationalObstructionGetResponseDataMode = "REAL"
	NavigationalObstructionGetResponseDataModeTest      NavigationalObstructionGetResponseDataMode = "TEST"
	NavigationalObstructionGetResponseDataModeSimulated NavigationalObstructionGetResponseDataMode = "SIMULATED"
	NavigationalObstructionGetResponseDataModeExercise  NavigationalObstructionGetResponseDataMode = "EXERCISE"
)

type NavigationalObstructionQueryhelpResponse struct {
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
func (r NavigationalObstructionQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *NavigationalObstructionQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Beta Version Navigational Obstruction: Information describing navigational
// obstructions, such as applicable boundaries, locations, heights, data ownership,
// and currency.
type NavigationalObstructionTupleResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Start date of this obstruction data set's currency, in ISO 8601 date-only
	// format.
	CycleDate time.Time `json:"cycleDate,required" format:"date"`
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
	DataMode NavigationalObstructionTupleResponseDataMode `json:"dataMode,required"`
	// The ID of this obstacle.
	ObstacleID string `json:"obstacleId,required"`
	// Type of obstacle (e.g. P for point, V for vector, L for line).
	ObstacleType string `json:"obstacleType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Indicates if this obstacle record is Active (A) or Deleted (D).
	ActDelCode string `json:"actDelCode"`
	// The Aeronautical Information Regulation and Control (AIRAC) cycle of this
	// obstruction data set. The format is YYNN where YY is the last two digits of the
	// year and NN is the cycle number.
	AiracCycle int64 `json:"airacCycle"`
	// The baseline Aeronautical Information Regulation and Control (AIRAC) cycle for
	// change sets only. The format is YYNN where YY is the last two digits of the year
	// and NN is the cycle number.
	BaseAiracCycle int64 `json:"baseAiracCycle"`
	// Earliest record date possible in this obstruction data set (not the earliest
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD). If null, this data
	// set is assumed to be a full data pull of holdings until the cutoffDate. If this
	// field is populated, this data set only contains updates since the last baseline
	// data set.
	BaselineCutoffDate time.Time `json:"baselineCutoffDate" format:"date"`
	// WGS-84 latitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundNeLat float64 `json:"boundNELat"`
	// WGS-84 longitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundNeLon float64 `json:"boundNELon"`
	// WGS-84 latitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundSwLat float64 `json:"boundSWLat"`
	// WGS-84 longitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundSwLon float64 `json:"boundSWLon"`
	// The DoD Standard Country Code designator for the country issuing the diplomatic
	// clearance. This field will be set to "OTHR" if the source value does not match a
	// UDL Country code value (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Latest record date possible in this obstruction data set (not the most recent
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	CutoffDate time.Time `json:"cutoffDate" format:"date"`
	// Remarks concerning this obstruction's data set.
	DataSetRemarks string `json:"dataSetRemarks"`
	// The organization that deleted this obstacle record.
	DeletingOrg string `json:"deletingOrg"`
	// The organization that entered obstacle data other than the producer.
	DerivingOrg string `json:"derivingOrg"`
	// The side or sides of this obstruction feature which produces the greatest
	// reflectivity potential.
	DirectivityCode int64 `json:"directivityCode"`
	// The elevation at the point obstacle's location in feet.
	Elevation float64 `json:"elevation"`
	// The difference between the assigned elevation of this point and its true
	// elevation, in feet.
	ElevationAcc float64 `json:"elevationAcc"`
	// Optional obstacle ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID string `json:"externalId"`
	// FACC (Feature and Attribute Coding Catalog) is a five-character code for
	// encoding real-world entities and objects. The first letter of the code is an
	// alphabetic value from "A" to "Z" which will map to a feature category. The
	// second character will map to a subcategory. Characters three to five are numeric
	// and range from 000 to 999. This value will provide a unit feature type
	// identification within the categories.
	Facc string `json:"facc"`
	// Identifying code for the type of this point obstacle.
	FeatureCode string `json:"featureCode"`
	// Description of this obstacle, corresponding to the FACC (Feature and Attribute
	// Coding Catalog) value.
	FeatureDescription string `json:"featureDescription"`
	// Type name of point obstacle.
	FeatureName string `json:"featureName"`
	// Identifying code for the type of this point obstacle.
	FeatureType string `json:"featureType"`
	// The height Above Ground Level (AGL) of the point obstacle in feet.
	HeightAgl float64 `json:"heightAGL"`
	// The accuracy of the height Above Ground Level (AGL) value for this point
	// obstacle, in feet.
	HeightAglAcc float64 `json:"heightAGLAcc"`
	// The height Above Mean Sea Level (AMSL) of the point obstacle in feet.
	HeightMsl float64 `json:"heightMSL"`
	// The accuracy of the height Above Mean Sea Level (AMSL) value for this point
	// obstacle in feet.
	HeightMslAcc float64 `json:"heightMSLAcc"`
	// The difference between the recorded horizontal coordinates of this point
	// obstacle and its true position, in feet.
	HorizAcc float64 `json:"horizAcc"`
	// Code representing the mathematical model of Earth used to calculate coordinates
	// for this obstacle (e.g. WGS-84, U for undetermined, etc.). US Forces use the
	// World Geodetic System 1984 (WGS-84), but also use maps by allied countries with
	// local datums.
	HorizDatumCode string `json:"horizDatumCode"`
	// Date this obstacle was initially added to the data set, in ISO 8601 date-only
	// format (ex. YYYY-MM-DD).
	InitRecordDate time.Time `json:"initRecordDate" format:"date"`
	// This field provides an array of keys that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the values array. This array
	// must be the same length as values.
	Keys []string `json:"keys"`
	// Code specifying if this obstacle is lit (e.g. Y = Yes, N = No, U = Unknown).
	LightingCode string `json:"lightingCode"`
	// WGS-84 latitude of the northeastern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineNeLat float64 `json:"lineNELat"`
	// WGS-84 longitude of the northeastern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineNeLon float64 `json:"lineNELon"`
	// The name of the line file associated with this obstruction data set.
	LinesFilename string `json:"linesFilename"`
	// WGS-84 latitude of the southwestern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineSwLat float64 `json:"lineSWLat"`
	// WGS-84 longitude of the southwestern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineSwLon float64 `json:"lineSWLon"`
	// The minimum height Above Ground Level (AGL) of the shortest obstruction
	// contained in this data set, in feet.
	MinHeightAgl float64 `json:"minHeightAGL"`
	// Indicates if the feature has multiple obstructions (e.g. S = Single, M =
	// Multiple, U = Undetermined).
	MultObs string `json:"multObs"`
	// The date after which this obstruction data set’s currency is stale and should be
	// refreshed, in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	NextCycleDate time.Time `json:"nextCycleDate" format:"date"`
	// The number of line features associated with this obstruction data set.
	NumLines int64 `json:"numLines"`
	// Indicates the number of obstructions associated with a feature.
	NumObs int64 `json:"numObs"`
	// The number of point features associated with this obstruction data set.
	NumPoints int64 `json:"numPoints"`
	// Remarks regarding this obstacle.
	ObstacleRemarks string `json:"obstacleRemarks"`
	// The original ID for this obstacle.
	OrigID string `json:"origId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The DoD Standard Country Code designator for the country or political entity
	// that owns the data set associated with this obstruction. This field will be set
	// to "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	OwnerCountryCode string `json:"ownerCountryCode"`
	// WGS-84 latitude of this point obstacle, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PointLat float64 `json:"pointLat"`
	// WGS-84 longitude of this point obstacle, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PointLon float64 `json:"pointLon"`
	// The name of the point file associated with this obstruction data set.
	PointsFilename string `json:"pointsFilename"`
	// Code denoting the action, review, or process that updated this obstacle.
	ProcessCode string `json:"processCode"`
	// Name of the agency that produced this obstruction data set.
	Producer string `json:"producer"`
	// The Federal Information Processing Standards (FIPS) state/province numeric code
	// of this obstacle's location.
	ProvinceCode string `json:"provinceCode"`
	// When horizontal and/or vertical accuracy requirements cannot be met because of
	// inadequate source material, this code indicates the quality of the data.
	Quality string `json:"quality"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Date this obstacle data was revised, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	RevDate time.Time `json:"revDate" format:"date"`
	// ID of the end point of a line segment.
	SegEndPoint int64 `json:"segEndPoint"`
	// Identifies the sequence number of a line segment.
	SegNum int64 `json:"segNum"`
	// ID of the starting point of a line segment.
	SegStartPoint int64 `json:"segStartPoint"`
	// Source date of this obstacle data, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	SourceDate time.Time `json:"sourceDate" format:"date"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The surface material composition code of this point obstacle.
	SurfaceMatCode string `json:"surfaceMatCode"`
	// The transaction type/code for this obstacle (e.g. "D", "N", "R", "S", "V", "X").
	TransactionCode string `json:"transactionCode"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Method used to confirm the existence of this obstacle.
	ValidationCode int64 `json:"validationCode"`
	// This field provides an array of values that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the keys array. This array
	// must be the same length as keys.
	Values []string `json:"values"`
	// The name of the vector file associated with this obstruction data set.
	VectorsFilename string `json:"vectorsFilename"`
	// The World Aeronautical Chart (WAC) identifier for the area in which this
	// obstacle is located.
	Wac string `json:"wac"`
	// This obstacle's World Area Code installation number (WAC-INNR).
	WacInnr string `json:"wacINNR"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		CycleDate             respjson.Field
		DataMode              respjson.Field
		ObstacleID            respjson.Field
		ObstacleType          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ActDelCode            respjson.Field
		AiracCycle            respjson.Field
		BaseAiracCycle        respjson.Field
		BaselineCutoffDate    respjson.Field
		BoundNeLat            respjson.Field
		BoundNeLon            respjson.Field
		BoundSwLat            respjson.Field
		BoundSwLon            respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CutoffDate            respjson.Field
		DataSetRemarks        respjson.Field
		DeletingOrg           respjson.Field
		DerivingOrg           respjson.Field
		DirectivityCode       respjson.Field
		Elevation             respjson.Field
		ElevationAcc          respjson.Field
		ExternalID            respjson.Field
		Facc                  respjson.Field
		FeatureCode           respjson.Field
		FeatureDescription    respjson.Field
		FeatureName           respjson.Field
		FeatureType           respjson.Field
		HeightAgl             respjson.Field
		HeightAglAcc          respjson.Field
		HeightMsl             respjson.Field
		HeightMslAcc          respjson.Field
		HorizAcc              respjson.Field
		HorizDatumCode        respjson.Field
		InitRecordDate        respjson.Field
		Keys                  respjson.Field
		LightingCode          respjson.Field
		LineNeLat             respjson.Field
		LineNeLon             respjson.Field
		LinesFilename         respjson.Field
		LineSwLat             respjson.Field
		LineSwLon             respjson.Field
		MinHeightAgl          respjson.Field
		MultObs               respjson.Field
		NextCycleDate         respjson.Field
		NumLines              respjson.Field
		NumObs                respjson.Field
		NumPoints             respjson.Field
		ObstacleRemarks       respjson.Field
		OrigID                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OwnerCountryCode      respjson.Field
		PointLat              respjson.Field
		PointLon              respjson.Field
		PointsFilename        respjson.Field
		ProcessCode           respjson.Field
		Producer              respjson.Field
		ProvinceCode          respjson.Field
		Quality               respjson.Field
		RawFileUri            respjson.Field
		RevDate               respjson.Field
		SegEndPoint           respjson.Field
		SegNum                respjson.Field
		SegStartPoint         respjson.Field
		SourceDate            respjson.Field
		SourceDl              respjson.Field
		SurfaceMatCode        respjson.Field
		TransactionCode       respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ValidationCode        respjson.Field
		Values                respjson.Field
		VectorsFilename       respjson.Field
		Wac                   respjson.Field
		WacInnr               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationalObstructionTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *NavigationalObstructionTupleResponse) UnmarshalJSON(data []byte) error {
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
type NavigationalObstructionTupleResponseDataMode string

const (
	NavigationalObstructionTupleResponseDataModeReal      NavigationalObstructionTupleResponseDataMode = "REAL"
	NavigationalObstructionTupleResponseDataModeTest      NavigationalObstructionTupleResponseDataMode = "TEST"
	NavigationalObstructionTupleResponseDataModeSimulated NavigationalObstructionTupleResponseDataMode = "SIMULATED"
	NavigationalObstructionTupleResponseDataModeExercise  NavigationalObstructionTupleResponseDataMode = "EXERCISE"
)

type NavigationalObstructionNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Start date of this obstruction data set's currency, in ISO 8601 date-only
	// format.
	CycleDate time.Time `json:"cycleDate,required" format:"date"`
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
	DataMode NavigationalObstructionNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The ID of this obstacle.
	ObstacleID string `json:"obstacleId,required"`
	// Type of obstacle (e.g. P for point, V for vector, L for line).
	ObstacleType string `json:"obstacleType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Indicates if this obstacle record is Active (A) or Deleted (D).
	ActDelCode param.Opt[string] `json:"actDelCode,omitzero"`
	// The Aeronautical Information Regulation and Control (AIRAC) cycle of this
	// obstruction data set. The format is YYNN where YY is the last two digits of the
	// year and NN is the cycle number.
	AiracCycle param.Opt[int64] `json:"airacCycle,omitzero"`
	// The baseline Aeronautical Information Regulation and Control (AIRAC) cycle for
	// change sets only. The format is YYNN where YY is the last two digits of the year
	// and NN is the cycle number.
	BaseAiracCycle param.Opt[int64] `json:"baseAiracCycle,omitzero"`
	// Earliest record date possible in this obstruction data set (not the earliest
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD). If null, this data
	// set is assumed to be a full data pull of holdings until the cutoffDate. If this
	// field is populated, this data set only contains updates since the last baseline
	// data set.
	BaselineCutoffDate param.Opt[time.Time] `json:"baselineCutoffDate,omitzero" format:"date"`
	// WGS-84 latitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundNeLat param.Opt[float64] `json:"boundNELat,omitzero"`
	// WGS-84 longitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundNeLon param.Opt[float64] `json:"boundNELon,omitzero"`
	// WGS-84 latitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundSwLat param.Opt[float64] `json:"boundSWLat,omitzero"`
	// WGS-84 longitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundSwLon param.Opt[float64] `json:"boundSWLon,omitzero"`
	// The DoD Standard Country Code designator for the country issuing the diplomatic
	// clearance. This field will be set to "OTHR" if the source value does not match a
	// UDL Country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Latest record date possible in this obstruction data set (not the most recent
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	CutoffDate param.Opt[time.Time] `json:"cutoffDate,omitzero" format:"date"`
	// Remarks concerning this obstruction's data set.
	DataSetRemarks param.Opt[string] `json:"dataSetRemarks,omitzero"`
	// The organization that deleted this obstacle record.
	DeletingOrg param.Opt[string] `json:"deletingOrg,omitzero"`
	// The organization that entered obstacle data other than the producer.
	DerivingOrg param.Opt[string] `json:"derivingOrg,omitzero"`
	// The side or sides of this obstruction feature which produces the greatest
	// reflectivity potential.
	DirectivityCode param.Opt[int64] `json:"directivityCode,omitzero"`
	// The elevation at the point obstacle's location in feet.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// The difference between the assigned elevation of this point and its true
	// elevation, in feet.
	ElevationAcc param.Opt[float64] `json:"elevationAcc,omitzero"`
	// Optional obstacle ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// FACC (Feature and Attribute Coding Catalog) is a five-character code for
	// encoding real-world entities and objects. The first letter of the code is an
	// alphabetic value from "A" to "Z" which will map to a feature category. The
	// second character will map to a subcategory. Characters three to five are numeric
	// and range from 000 to 999. This value will provide a unit feature type
	// identification within the categories.
	Facc param.Opt[string] `json:"facc,omitzero"`
	// Identifying code for the type of this point obstacle.
	FeatureCode param.Opt[string] `json:"featureCode,omitzero"`
	// Description of this obstacle, corresponding to the FACC (Feature and Attribute
	// Coding Catalog) value.
	FeatureDescription param.Opt[string] `json:"featureDescription,omitzero"`
	// Type name of point obstacle.
	FeatureName param.Opt[string] `json:"featureName,omitzero"`
	// Identifying code for the type of this point obstacle.
	FeatureType param.Opt[string] `json:"featureType,omitzero"`
	// The height Above Ground Level (AGL) of the point obstacle in feet.
	HeightAgl param.Opt[float64] `json:"heightAGL,omitzero"`
	// The accuracy of the height Above Ground Level (AGL) value for this point
	// obstacle, in feet.
	HeightAglAcc param.Opt[float64] `json:"heightAGLAcc,omitzero"`
	// The height Above Mean Sea Level (AMSL) of the point obstacle in feet.
	HeightMsl param.Opt[float64] `json:"heightMSL,omitzero"`
	// The accuracy of the height Above Mean Sea Level (AMSL) value for this point
	// obstacle in feet.
	HeightMslAcc param.Opt[float64] `json:"heightMSLAcc,omitzero"`
	// The difference between the recorded horizontal coordinates of this point
	// obstacle and its true position, in feet.
	HorizAcc param.Opt[float64] `json:"horizAcc,omitzero"`
	// Code representing the mathematical model of Earth used to calculate coordinates
	// for this obstacle (e.g. WGS-84, U for undetermined, etc.). US Forces use the
	// World Geodetic System 1984 (WGS-84), but also use maps by allied countries with
	// local datums.
	HorizDatumCode param.Opt[string] `json:"horizDatumCode,omitzero"`
	// Date this obstacle was initially added to the data set, in ISO 8601 date-only
	// format (ex. YYYY-MM-DD).
	InitRecordDate param.Opt[time.Time] `json:"initRecordDate,omitzero" format:"date"`
	// Code specifying if this obstacle is lit (e.g. Y = Yes, N = No, U = Unknown).
	LightingCode param.Opt[string] `json:"lightingCode,omitzero"`
	// WGS-84 latitude of the northeastern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineNeLat param.Opt[float64] `json:"lineNELat,omitzero"`
	// WGS-84 longitude of the northeastern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineNeLon param.Opt[float64] `json:"lineNELon,omitzero"`
	// The name of the line file associated with this obstruction data set.
	LinesFilename param.Opt[string] `json:"linesFilename,omitzero"`
	// WGS-84 latitude of the southwestern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineSwLat param.Opt[float64] `json:"lineSWLat,omitzero"`
	// WGS-84 longitude of the southwestern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineSwLon param.Opt[float64] `json:"lineSWLon,omitzero"`
	// The minimum height Above Ground Level (AGL) of the shortest obstruction
	// contained in this data set, in feet.
	MinHeightAgl param.Opt[float64] `json:"minHeightAGL,omitzero"`
	// Indicates if the feature has multiple obstructions (e.g. S = Single, M =
	// Multiple, U = Undetermined).
	MultObs param.Opt[string] `json:"multObs,omitzero"`
	// The date after which this obstruction data set’s currency is stale and should be
	// refreshed, in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	NextCycleDate param.Opt[time.Time] `json:"nextCycleDate,omitzero" format:"date"`
	// The number of line features associated with this obstruction data set.
	NumLines param.Opt[int64] `json:"numLines,omitzero"`
	// Indicates the number of obstructions associated with a feature.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// The number of point features associated with this obstruction data set.
	NumPoints param.Opt[int64] `json:"numPoints,omitzero"`
	// Remarks regarding this obstacle.
	ObstacleRemarks param.Opt[string] `json:"obstacleRemarks,omitzero"`
	// The original ID for this obstacle.
	OrigID param.Opt[string] `json:"origId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity
	// that owns the data set associated with this obstruction. This field will be set
	// to "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	OwnerCountryCode param.Opt[string] `json:"ownerCountryCode,omitzero"`
	// WGS-84 latitude of this point obstacle, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PointLat param.Opt[float64] `json:"pointLat,omitzero"`
	// WGS-84 longitude of this point obstacle, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PointLon param.Opt[float64] `json:"pointLon,omitzero"`
	// The name of the point file associated with this obstruction data set.
	PointsFilename param.Opt[string] `json:"pointsFilename,omitzero"`
	// Code denoting the action, review, or process that updated this obstacle.
	ProcessCode param.Opt[string] `json:"processCode,omitzero"`
	// Name of the agency that produced this obstruction data set.
	Producer param.Opt[string] `json:"producer,omitzero"`
	// The Federal Information Processing Standards (FIPS) state/province numeric code
	// of this obstacle's location.
	ProvinceCode param.Opt[string] `json:"provinceCode,omitzero"`
	// When horizontal and/or vertical accuracy requirements cannot be met because of
	// inadequate source material, this code indicates the quality of the data.
	Quality param.Opt[string] `json:"quality,omitzero"`
	// Date this obstacle data was revised, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	RevDate param.Opt[time.Time] `json:"revDate,omitzero" format:"date"`
	// ID of the end point of a line segment.
	SegEndPoint param.Opt[int64] `json:"segEndPoint,omitzero"`
	// Identifies the sequence number of a line segment.
	SegNum param.Opt[int64] `json:"segNum,omitzero"`
	// ID of the starting point of a line segment.
	SegStartPoint param.Opt[int64] `json:"segStartPoint,omitzero"`
	// Source date of this obstacle data, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	SourceDate param.Opt[time.Time] `json:"sourceDate,omitzero" format:"date"`
	// The surface material composition code of this point obstacle.
	SurfaceMatCode param.Opt[string] `json:"surfaceMatCode,omitzero"`
	// The transaction type/code for this obstacle (e.g. "D", "N", "R", "S", "V", "X").
	TransactionCode param.Opt[string] `json:"transactionCode,omitzero"`
	// Method used to confirm the existence of this obstacle.
	ValidationCode param.Opt[int64] `json:"validationCode,omitzero"`
	// The name of the vector file associated with this obstruction data set.
	VectorsFilename param.Opt[string] `json:"vectorsFilename,omitzero"`
	// The World Aeronautical Chart (WAC) identifier for the area in which this
	// obstacle is located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	// This obstacle's World Area Code installation number (WAC-INNR).
	WacInnr param.Opt[string] `json:"wacINNR,omitzero"`
	// This field provides an array of keys that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the values array. This array
	// must be the same length as values.
	Keys []string `json:"keys,omitzero"`
	// This field provides an array of values that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the keys array. This array
	// must be the same length as keys.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r NavigationalObstructionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NavigationalObstructionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NavigationalObstructionNewParams) UnmarshalJSON(data []byte) error {
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
type NavigationalObstructionNewParamsDataMode string

const (
	NavigationalObstructionNewParamsDataModeReal      NavigationalObstructionNewParamsDataMode = "REAL"
	NavigationalObstructionNewParamsDataModeTest      NavigationalObstructionNewParamsDataMode = "TEST"
	NavigationalObstructionNewParamsDataModeSimulated NavigationalObstructionNewParamsDataMode = "SIMULATED"
	NavigationalObstructionNewParamsDataModeExercise  NavigationalObstructionNewParamsDataMode = "EXERCISE"
)

type NavigationalObstructionUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Start date of this obstruction data set's currency, in ISO 8601 date-only
	// format.
	CycleDate time.Time `json:"cycleDate,required" format:"date"`
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
	DataMode NavigationalObstructionUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The ID of this obstacle.
	ObstacleID string `json:"obstacleId,required"`
	// Type of obstacle (e.g. P for point, V for vector, L for line).
	ObstacleType string `json:"obstacleType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Indicates if this obstacle record is Active (A) or Deleted (D).
	ActDelCode param.Opt[string] `json:"actDelCode,omitzero"`
	// The Aeronautical Information Regulation and Control (AIRAC) cycle of this
	// obstruction data set. The format is YYNN where YY is the last two digits of the
	// year and NN is the cycle number.
	AiracCycle param.Opt[int64] `json:"airacCycle,omitzero"`
	// The baseline Aeronautical Information Regulation and Control (AIRAC) cycle for
	// change sets only. The format is YYNN where YY is the last two digits of the year
	// and NN is the cycle number.
	BaseAiracCycle param.Opt[int64] `json:"baseAiracCycle,omitzero"`
	// Earliest record date possible in this obstruction data set (not the earliest
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD). If null, this data
	// set is assumed to be a full data pull of holdings until the cutoffDate. If this
	// field is populated, this data set only contains updates since the last baseline
	// data set.
	BaselineCutoffDate param.Opt[time.Time] `json:"baselineCutoffDate,omitzero" format:"date"`
	// WGS-84 latitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundNeLat param.Opt[float64] `json:"boundNELat,omitzero"`
	// WGS-84 longitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundNeLon param.Opt[float64] `json:"boundNELon,omitzero"`
	// WGS-84 latitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundSwLat param.Opt[float64] `json:"boundSWLat,omitzero"`
	// WGS-84 longitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundSwLon param.Opt[float64] `json:"boundSWLon,omitzero"`
	// The DoD Standard Country Code designator for the country issuing the diplomatic
	// clearance. This field will be set to "OTHR" if the source value does not match a
	// UDL Country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Latest record date possible in this obstruction data set (not the most recent
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	CutoffDate param.Opt[time.Time] `json:"cutoffDate,omitzero" format:"date"`
	// Remarks concerning this obstruction's data set.
	DataSetRemarks param.Opt[string] `json:"dataSetRemarks,omitzero"`
	// The organization that deleted this obstacle record.
	DeletingOrg param.Opt[string] `json:"deletingOrg,omitzero"`
	// The organization that entered obstacle data other than the producer.
	DerivingOrg param.Opt[string] `json:"derivingOrg,omitzero"`
	// The side or sides of this obstruction feature which produces the greatest
	// reflectivity potential.
	DirectivityCode param.Opt[int64] `json:"directivityCode,omitzero"`
	// The elevation at the point obstacle's location in feet.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// The difference between the assigned elevation of this point and its true
	// elevation, in feet.
	ElevationAcc param.Opt[float64] `json:"elevationAcc,omitzero"`
	// Optional obstacle ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// FACC (Feature and Attribute Coding Catalog) is a five-character code for
	// encoding real-world entities and objects. The first letter of the code is an
	// alphabetic value from "A" to "Z" which will map to a feature category. The
	// second character will map to a subcategory. Characters three to five are numeric
	// and range from 000 to 999. This value will provide a unit feature type
	// identification within the categories.
	Facc param.Opt[string] `json:"facc,omitzero"`
	// Identifying code for the type of this point obstacle.
	FeatureCode param.Opt[string] `json:"featureCode,omitzero"`
	// Description of this obstacle, corresponding to the FACC (Feature and Attribute
	// Coding Catalog) value.
	FeatureDescription param.Opt[string] `json:"featureDescription,omitzero"`
	// Type name of point obstacle.
	FeatureName param.Opt[string] `json:"featureName,omitzero"`
	// Identifying code for the type of this point obstacle.
	FeatureType param.Opt[string] `json:"featureType,omitzero"`
	// The height Above Ground Level (AGL) of the point obstacle in feet.
	HeightAgl param.Opt[float64] `json:"heightAGL,omitzero"`
	// The accuracy of the height Above Ground Level (AGL) value for this point
	// obstacle, in feet.
	HeightAglAcc param.Opt[float64] `json:"heightAGLAcc,omitzero"`
	// The height Above Mean Sea Level (AMSL) of the point obstacle in feet.
	HeightMsl param.Opt[float64] `json:"heightMSL,omitzero"`
	// The accuracy of the height Above Mean Sea Level (AMSL) value for this point
	// obstacle in feet.
	HeightMslAcc param.Opt[float64] `json:"heightMSLAcc,omitzero"`
	// The difference between the recorded horizontal coordinates of this point
	// obstacle and its true position, in feet.
	HorizAcc param.Opt[float64] `json:"horizAcc,omitzero"`
	// Code representing the mathematical model of Earth used to calculate coordinates
	// for this obstacle (e.g. WGS-84, U for undetermined, etc.). US Forces use the
	// World Geodetic System 1984 (WGS-84), but also use maps by allied countries with
	// local datums.
	HorizDatumCode param.Opt[string] `json:"horizDatumCode,omitzero"`
	// Date this obstacle was initially added to the data set, in ISO 8601 date-only
	// format (ex. YYYY-MM-DD).
	InitRecordDate param.Opt[time.Time] `json:"initRecordDate,omitzero" format:"date"`
	// Code specifying if this obstacle is lit (e.g. Y = Yes, N = No, U = Unknown).
	LightingCode param.Opt[string] `json:"lightingCode,omitzero"`
	// WGS-84 latitude of the northeastern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineNeLat param.Opt[float64] `json:"lineNELat,omitzero"`
	// WGS-84 longitude of the northeastern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineNeLon param.Opt[float64] `json:"lineNELon,omitzero"`
	// The name of the line file associated with this obstruction data set.
	LinesFilename param.Opt[string] `json:"linesFilename,omitzero"`
	// WGS-84 latitude of the southwestern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineSwLat param.Opt[float64] `json:"lineSWLat,omitzero"`
	// WGS-84 longitude of the southwestern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineSwLon param.Opt[float64] `json:"lineSWLon,omitzero"`
	// The minimum height Above Ground Level (AGL) of the shortest obstruction
	// contained in this data set, in feet.
	MinHeightAgl param.Opt[float64] `json:"minHeightAGL,omitzero"`
	// Indicates if the feature has multiple obstructions (e.g. S = Single, M =
	// Multiple, U = Undetermined).
	MultObs param.Opt[string] `json:"multObs,omitzero"`
	// The date after which this obstruction data set’s currency is stale and should be
	// refreshed, in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	NextCycleDate param.Opt[time.Time] `json:"nextCycleDate,omitzero" format:"date"`
	// The number of line features associated with this obstruction data set.
	NumLines param.Opt[int64] `json:"numLines,omitzero"`
	// Indicates the number of obstructions associated with a feature.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// The number of point features associated with this obstruction data set.
	NumPoints param.Opt[int64] `json:"numPoints,omitzero"`
	// Remarks regarding this obstacle.
	ObstacleRemarks param.Opt[string] `json:"obstacleRemarks,omitzero"`
	// The original ID for this obstacle.
	OrigID param.Opt[string] `json:"origId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity
	// that owns the data set associated with this obstruction. This field will be set
	// to "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	OwnerCountryCode param.Opt[string] `json:"ownerCountryCode,omitzero"`
	// WGS-84 latitude of this point obstacle, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PointLat param.Opt[float64] `json:"pointLat,omitzero"`
	// WGS-84 longitude of this point obstacle, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PointLon param.Opt[float64] `json:"pointLon,omitzero"`
	// The name of the point file associated with this obstruction data set.
	PointsFilename param.Opt[string] `json:"pointsFilename,omitzero"`
	// Code denoting the action, review, or process that updated this obstacle.
	ProcessCode param.Opt[string] `json:"processCode,omitzero"`
	// Name of the agency that produced this obstruction data set.
	Producer param.Opt[string] `json:"producer,omitzero"`
	// The Federal Information Processing Standards (FIPS) state/province numeric code
	// of this obstacle's location.
	ProvinceCode param.Opt[string] `json:"provinceCode,omitzero"`
	// When horizontal and/or vertical accuracy requirements cannot be met because of
	// inadequate source material, this code indicates the quality of the data.
	Quality param.Opt[string] `json:"quality,omitzero"`
	// Date this obstacle data was revised, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	RevDate param.Opt[time.Time] `json:"revDate,omitzero" format:"date"`
	// ID of the end point of a line segment.
	SegEndPoint param.Opt[int64] `json:"segEndPoint,omitzero"`
	// Identifies the sequence number of a line segment.
	SegNum param.Opt[int64] `json:"segNum,omitzero"`
	// ID of the starting point of a line segment.
	SegStartPoint param.Opt[int64] `json:"segStartPoint,omitzero"`
	// Source date of this obstacle data, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	SourceDate param.Opt[time.Time] `json:"sourceDate,omitzero" format:"date"`
	// The surface material composition code of this point obstacle.
	SurfaceMatCode param.Opt[string] `json:"surfaceMatCode,omitzero"`
	// The transaction type/code for this obstacle (e.g. "D", "N", "R", "S", "V", "X").
	TransactionCode param.Opt[string] `json:"transactionCode,omitzero"`
	// Method used to confirm the existence of this obstacle.
	ValidationCode param.Opt[int64] `json:"validationCode,omitzero"`
	// The name of the vector file associated with this obstruction data set.
	VectorsFilename param.Opt[string] `json:"vectorsFilename,omitzero"`
	// The World Aeronautical Chart (WAC) identifier for the area in which this
	// obstacle is located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	// This obstacle's World Area Code installation number (WAC-INNR).
	WacInnr param.Opt[string] `json:"wacINNR,omitzero"`
	// This field provides an array of keys that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the values array. This array
	// must be the same length as values.
	Keys []string `json:"keys,omitzero"`
	// This field provides an array of values that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the keys array. This array
	// must be the same length as keys.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r NavigationalObstructionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NavigationalObstructionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NavigationalObstructionUpdateParams) UnmarshalJSON(data []byte) error {
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
type NavigationalObstructionUpdateParamsDataMode string

const (
	NavigationalObstructionUpdateParamsDataModeReal      NavigationalObstructionUpdateParamsDataMode = "REAL"
	NavigationalObstructionUpdateParamsDataModeTest      NavigationalObstructionUpdateParamsDataMode = "TEST"
	NavigationalObstructionUpdateParamsDataModeSimulated NavigationalObstructionUpdateParamsDataMode = "SIMULATED"
	NavigationalObstructionUpdateParamsDataModeExercise  NavigationalObstructionUpdateParamsDataMode = "EXERCISE"
)

type NavigationalObstructionListParams struct {
	// (One or more of fields 'cycleDate, obstacleId' are required.) Start date of this
	// obstruction data set's currency, in ISO 8601 date-only format. (YYYY-MM-DD)
	CycleDate   param.Opt[time.Time] `query:"cycleDate,omitzero" format:"date" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'cycleDate, obstacleId' are required.) The ID of this
	// obstacle.
	ObstacleID param.Opt[string] `query:"obstacleId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NavigationalObstructionListParams]'s query parameters as
// `url.Values`.
func (r NavigationalObstructionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NavigationalObstructionCountParams struct {
	// (One or more of fields 'cycleDate, obstacleId' are required.) Start date of this
	// obstruction data set's currency, in ISO 8601 date-only format. (YYYY-MM-DD)
	CycleDate   param.Opt[time.Time] `query:"cycleDate,omitzero" format:"date" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'cycleDate, obstacleId' are required.) The ID of this
	// obstacle.
	ObstacleID param.Opt[string] `query:"obstacleId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NavigationalObstructionCountParams]'s query parameters as
// `url.Values`.
func (r NavigationalObstructionCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NavigationalObstructionNewBulkParams struct {
	Body []NavigationalObstructionNewBulkParamsBody
	paramObj
}

func (r NavigationalObstructionNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *NavigationalObstructionNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Beta Version Navigational Obstruction: Information describing navigational
// obstructions, such as applicable boundaries, locations, heights, data ownership,
// and currency.
//
// The properties ClassificationMarking, CycleDate, DataMode, ObstacleID,
// ObstacleType, Source are required.
type NavigationalObstructionNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Start date of this obstruction data set's currency, in ISO 8601 date-only
	// format.
	CycleDate time.Time `json:"cycleDate,required" format:"date"`
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
	// The ID of this obstacle.
	ObstacleID string `json:"obstacleId,required"`
	// Type of obstacle (e.g. P for point, V for vector, L for line).
	ObstacleType string `json:"obstacleType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Indicates if this obstacle record is Active (A) or Deleted (D).
	ActDelCode param.Opt[string] `json:"actDelCode,omitzero"`
	// The Aeronautical Information Regulation and Control (AIRAC) cycle of this
	// obstruction data set. The format is YYNN where YY is the last two digits of the
	// year and NN is the cycle number.
	AiracCycle param.Opt[int64] `json:"airacCycle,omitzero"`
	// The baseline Aeronautical Information Regulation and Control (AIRAC) cycle for
	// change sets only. The format is YYNN where YY is the last two digits of the year
	// and NN is the cycle number.
	BaseAiracCycle param.Opt[int64] `json:"baseAiracCycle,omitzero"`
	// Earliest record date possible in this obstruction data set (not the earliest
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD). If null, this data
	// set is assumed to be a full data pull of holdings until the cutoffDate. If this
	// field is populated, this data set only contains updates since the last baseline
	// data set.
	BaselineCutoffDate param.Opt[time.Time] `json:"baselineCutoffDate,omitzero" format:"date"`
	// WGS-84 latitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundNeLat param.Opt[float64] `json:"boundNELat,omitzero"`
	// WGS-84 longitude of the northeastern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundNeLon param.Opt[float64] `json:"boundNELon,omitzero"`
	// WGS-84 latitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -90 to 90 degrees (negative values south of equator).
	BoundSwLat param.Opt[float64] `json:"boundSWLat,omitzero"`
	// WGS-84 longitude of the southwestern boundary for obstructions contained in this
	// data set, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	BoundSwLon param.Opt[float64] `json:"boundSWLon,omitzero"`
	// The DoD Standard Country Code designator for the country issuing the diplomatic
	// clearance. This field will be set to "OTHR" if the source value does not match a
	// UDL Country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Latest record date possible in this obstruction data set (not the most recent
	// data item), in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	CutoffDate param.Opt[time.Time] `json:"cutoffDate,omitzero" format:"date"`
	// Remarks concerning this obstruction's data set.
	DataSetRemarks param.Opt[string] `json:"dataSetRemarks,omitzero"`
	// The organization that deleted this obstacle record.
	DeletingOrg param.Opt[string] `json:"deletingOrg,omitzero"`
	// The organization that entered obstacle data other than the producer.
	DerivingOrg param.Opt[string] `json:"derivingOrg,omitzero"`
	// The side or sides of this obstruction feature which produces the greatest
	// reflectivity potential.
	DirectivityCode param.Opt[int64] `json:"directivityCode,omitzero"`
	// The elevation at the point obstacle's location in feet.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// The difference between the assigned elevation of this point and its true
	// elevation, in feet.
	ElevationAcc param.Opt[float64] `json:"elevationAcc,omitzero"`
	// Optional obstacle ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// FACC (Feature and Attribute Coding Catalog) is a five-character code for
	// encoding real-world entities and objects. The first letter of the code is an
	// alphabetic value from "A" to "Z" which will map to a feature category. The
	// second character will map to a subcategory. Characters three to five are numeric
	// and range from 000 to 999. This value will provide a unit feature type
	// identification within the categories.
	Facc param.Opt[string] `json:"facc,omitzero"`
	// Identifying code for the type of this point obstacle.
	FeatureCode param.Opt[string] `json:"featureCode,omitzero"`
	// Description of this obstacle, corresponding to the FACC (Feature and Attribute
	// Coding Catalog) value.
	FeatureDescription param.Opt[string] `json:"featureDescription,omitzero"`
	// Type name of point obstacle.
	FeatureName param.Opt[string] `json:"featureName,omitzero"`
	// Identifying code for the type of this point obstacle.
	FeatureType param.Opt[string] `json:"featureType,omitzero"`
	// The height Above Ground Level (AGL) of the point obstacle in feet.
	HeightAgl param.Opt[float64] `json:"heightAGL,omitzero"`
	// The accuracy of the height Above Ground Level (AGL) value for this point
	// obstacle, in feet.
	HeightAglAcc param.Opt[float64] `json:"heightAGLAcc,omitzero"`
	// The height Above Mean Sea Level (AMSL) of the point obstacle in feet.
	HeightMsl param.Opt[float64] `json:"heightMSL,omitzero"`
	// The accuracy of the height Above Mean Sea Level (AMSL) value for this point
	// obstacle in feet.
	HeightMslAcc param.Opt[float64] `json:"heightMSLAcc,omitzero"`
	// The difference between the recorded horizontal coordinates of this point
	// obstacle and its true position, in feet.
	HorizAcc param.Opt[float64] `json:"horizAcc,omitzero"`
	// Code representing the mathematical model of Earth used to calculate coordinates
	// for this obstacle (e.g. WGS-84, U for undetermined, etc.). US Forces use the
	// World Geodetic System 1984 (WGS-84), but also use maps by allied countries with
	// local datums.
	HorizDatumCode param.Opt[string] `json:"horizDatumCode,omitzero"`
	// Date this obstacle was initially added to the data set, in ISO 8601 date-only
	// format (ex. YYYY-MM-DD).
	InitRecordDate param.Opt[time.Time] `json:"initRecordDate,omitzero" format:"date"`
	// Code specifying if this obstacle is lit (e.g. Y = Yes, N = No, U = Unknown).
	LightingCode param.Opt[string] `json:"lightingCode,omitzero"`
	// WGS-84 latitude of the northeastern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineNeLat param.Opt[float64] `json:"lineNELat,omitzero"`
	// WGS-84 longitude of the northeastern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineNeLon param.Opt[float64] `json:"lineNELon,omitzero"`
	// The name of the line file associated with this obstruction data set.
	LinesFilename param.Opt[string] `json:"linesFilename,omitzero"`
	// WGS-84 latitude of the southwestern point of the line, in degrees. -90 to 90
	// degrees (negative values south of equator).
	LineSwLat param.Opt[float64] `json:"lineSWLat,omitzero"`
	// WGS-84 longitude of the southwestern point of the line, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	LineSwLon param.Opt[float64] `json:"lineSWLon,omitzero"`
	// The minimum height Above Ground Level (AGL) of the shortest obstruction
	// contained in this data set, in feet.
	MinHeightAgl param.Opt[float64] `json:"minHeightAGL,omitzero"`
	// Indicates if the feature has multiple obstructions (e.g. S = Single, M =
	// Multiple, U = Undetermined).
	MultObs param.Opt[string] `json:"multObs,omitzero"`
	// The date after which this obstruction data set’s currency is stale and should be
	// refreshed, in ISO 8601 date-only format (e.g. YYYY-MM-DD).
	NextCycleDate param.Opt[time.Time] `json:"nextCycleDate,omitzero" format:"date"`
	// The number of line features associated with this obstruction data set.
	NumLines param.Opt[int64] `json:"numLines,omitzero"`
	// Indicates the number of obstructions associated with a feature.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// The number of point features associated with this obstruction data set.
	NumPoints param.Opt[int64] `json:"numPoints,omitzero"`
	// Remarks regarding this obstacle.
	ObstacleRemarks param.Opt[string] `json:"obstacleRemarks,omitzero"`
	// The original ID for this obstacle.
	OrigID param.Opt[string] `json:"origId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity
	// that owns the data set associated with this obstruction. This field will be set
	// to "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	OwnerCountryCode param.Opt[string] `json:"ownerCountryCode,omitzero"`
	// WGS-84 latitude of this point obstacle, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PointLat param.Opt[float64] `json:"pointLat,omitzero"`
	// WGS-84 longitude of this point obstacle, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PointLon param.Opt[float64] `json:"pointLon,omitzero"`
	// The name of the point file associated with this obstruction data set.
	PointsFilename param.Opt[string] `json:"pointsFilename,omitzero"`
	// Code denoting the action, review, or process that updated this obstacle.
	ProcessCode param.Opt[string] `json:"processCode,omitzero"`
	// Name of the agency that produced this obstruction data set.
	Producer param.Opt[string] `json:"producer,omitzero"`
	// The Federal Information Processing Standards (FIPS) state/province numeric code
	// of this obstacle's location.
	ProvinceCode param.Opt[string] `json:"provinceCode,omitzero"`
	// When horizontal and/or vertical accuracy requirements cannot be met because of
	// inadequate source material, this code indicates the quality of the data.
	Quality param.Opt[string] `json:"quality,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Date this obstacle data was revised, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	RevDate param.Opt[time.Time] `json:"revDate,omitzero" format:"date"`
	// ID of the end point of a line segment.
	SegEndPoint param.Opt[int64] `json:"segEndPoint,omitzero"`
	// Identifies the sequence number of a line segment.
	SegNum param.Opt[int64] `json:"segNum,omitzero"`
	// ID of the starting point of a line segment.
	SegStartPoint param.Opt[int64] `json:"segStartPoint,omitzero"`
	// Source date of this obstacle data, in ISO 8601 date-only format (ex.
	// YYYY-MM-DD).
	SourceDate param.Opt[time.Time] `json:"sourceDate,omitzero" format:"date"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The surface material composition code of this point obstacle.
	SurfaceMatCode param.Opt[string] `json:"surfaceMatCode,omitzero"`
	// The transaction type/code for this obstacle (e.g. "D", "N", "R", "S", "V", "X").
	TransactionCode param.Opt[string] `json:"transactionCode,omitzero"`
	// Method used to confirm the existence of this obstacle.
	ValidationCode param.Opt[int64] `json:"validationCode,omitzero"`
	// The name of the vector file associated with this obstruction data set.
	VectorsFilename param.Opt[string] `json:"vectorsFilename,omitzero"`
	// The World Aeronautical Chart (WAC) identifier for the area in which this
	// obstacle is located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	// This obstacle's World Area Code installation number (WAC-INNR).
	WacInnr param.Opt[string] `json:"wacINNR,omitzero"`
	// This field provides an array of keys that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the values array. This array
	// must be the same length as values.
	Keys []string `json:"keys,omitzero"`
	// This field provides an array of values that can be added to any obstruction
	// feature to provide information that is not already supported. The entries in
	// this array must correspond to the position index in the keys array. This array
	// must be the same length as keys.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r NavigationalObstructionNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow NavigationalObstructionNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NavigationalObstructionNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NavigationalObstructionNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type NavigationalObstructionGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NavigationalObstructionGetParams]'s query parameters as
// `url.Values`.
func (r NavigationalObstructionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NavigationalObstructionTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// (One or more of fields 'cycleDate, obstacleId' are required.) Start date of this
	// obstruction data set's currency, in ISO 8601 date-only format. (YYYY-MM-DD)
	CycleDate   param.Opt[time.Time] `query:"cycleDate,omitzero" format:"date" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'cycleDate, obstacleId' are required.) The ID of this
	// obstacle.
	ObstacleID param.Opt[string] `query:"obstacleId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NavigationalObstructionTupleParams]'s query parameters as
// `url.Values`.
func (r NavigationalObstructionTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
