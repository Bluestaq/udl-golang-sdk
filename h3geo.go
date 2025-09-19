// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
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

// H3GeoService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewH3GeoService] method instead.
type H3GeoService struct {
	Options []option.RequestOption
	History H3GeoHistoryService
}

// NewH3GeoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewH3GeoService(opts ...option.RequestOption) (r H3GeoService) {
	r = H3GeoService{}
	r.Options = opts
	r.History = NewH3GeoHistoryService(opts...)
	return
}

// Service operation to take a single H3Geo record as a POST body and ingest into
// the database. This operation does not persist any H3GeoHexCell points that may
// be present in the body of the request. This operation is not intended to be used
// for automated feeds into UDL. Data providers should contact the UDL team for
// specific role assignments and for instructions on setting up a permanent feed
// through an alternate mechanism.
func (r *H3GeoService) New(ctx context.Context, body H3GeoNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/h3geo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *H3GeoService) List(ctx context.Context, query H3GeoListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[H3GeoListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/h3geo"
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
func (r *H3GeoService) ListAutoPaging(ctx context.Context, query H3GeoListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[H3GeoListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *H3GeoService) Count(ctx context.Context, query H3GeoCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/h3geo/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single RF geolocation by its unique ID passed as a
// path parameter.
func (r *H3GeoService) Get(ctx context.Context, id string, query H3GeoGetParams, opts ...option.RequestOption) (res *H3GeoGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/h3geo/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *H3GeoService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *H3GeoQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/h3geo/queryhelp"
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
func (r *H3GeoService) Tuple(ctx context.Context, query H3GeoTupleParams, opts ...option.RequestOption) (res *[]H3GeoTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/h3geo/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// H3 Geospatial Binning is a discrete global grid system for indexing geographies
// into a hexagonal grid. The H3 schema is a collection of hexagonal cells that
// contain data for producing geospatial maps over a specified time span.
type H3GeoListResponse struct {
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
	DataMode H3GeoListResponseDataMode `json:"dataMode,required"`
	// The number of cells associated with this H3 Geo data set. At this time, UDL
	// supports up to 50,000 cells.
	NumCells int64 `json:"numCells,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Start time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The center frequency of this H3 Geo data set measured in megahertz.
	CenterFreq float64 `json:"centerFreq"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// End time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// H3 resolution (0 – 15) for the data set. At this time, UDL supports a resolution
	// of 3 or less.
	Resolution int64 `json:"resolution"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// An optional field containing the type of data that is represented by this H3 Geo
	// data set.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		NumCells              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		CenterFreq            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndTime               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Resolution            respjson.Field
		SourceDl              respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3GeoListResponse) RawJSON() string { return r.JSON.raw }
func (r *H3GeoListResponse) UnmarshalJSON(data []byte) error {
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
type H3GeoListResponseDataMode string

const (
	H3GeoListResponseDataModeReal      H3GeoListResponseDataMode = "REAL"
	H3GeoListResponseDataModeTest      H3GeoListResponseDataMode = "TEST"
	H3GeoListResponseDataModeSimulated H3GeoListResponseDataMode = "SIMULATED"
	H3GeoListResponseDataModeExercise  H3GeoListResponseDataMode = "EXERCISE"
)

// H3 Geospatial Binning is a discrete global grid system for indexing geographies
// into a hexagonal grid. The H3 schema is a collection of hexagonal cells that
// contain data for producing geospatial maps over a specified time span.
type H3GeoGetResponse struct {
	// The collection of hex cells contained in this H3 data set. The number of cells
	// is a function of the specified resolution.
	Cells []H3GeoGetResponseCell `json:"cells,required"`
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
	DataMode H3GeoGetResponseDataMode `json:"dataMode,required"`
	// The number of cells associated with this H3 Geo data set. At this time, UDL
	// supports up to 50,000 cells.
	NumCells int64 `json:"numCells,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Start time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The center frequency of this H3 Geo data set measured in megahertz.
	CenterFreq float64 `json:"centerFreq"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// End time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// H3 resolution (0 – 15) for the data set. At this time, UDL supports a resolution
	// of 3 or less.
	Resolution int64 `json:"resolution"`
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
	// An optional field containing the type of data that is represented by this H3 Geo
	// data set.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cells                 respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		NumCells              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		CenterFreq            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndTime               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Resolution            respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3GeoGetResponse) RawJSON() string { return r.JSON.raw }
func (r *H3GeoGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a hex cell array containing data for a set of
// observations.
type H3GeoGetResponseCell struct {
	// The H3 index represented as a 16 character hexadecimal string.
	CellID string `json:"cellId,required"`
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The mean altitude of the set of observations within this cell, measured in
	// kilometers.
	AltMean float64 `json:"altMean"`
	// The standard deviation of alttitude in the set of observations within this cell,
	// measured in kilometers.
	AltSigma float64 `json:"altSigma"`
	// The anomaly score for probable manufactured interference or RF interference;
	// calculated as a ratio of #anomalous obs / #total obs or coverage.
	AnomScoreInterference float64 `json:"anomScoreInterference"`
	// The anomaly score for probable spoofing; calculated as a ratio of #anomalous obs
	// / #total obs or coverage.
	AnomScoreSpoofing float64 `json:"anomScoreSpoofing"`
	// The percentage degree of change in the aggregated observables for a particular
	// H3 bin.
	ChangeScore float64 `json:"changeScore"`
	// The total number of available observations in the H3 cell during the start/end
	// times.
	Coverage int64 `json:"coverage"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent H3 Geo record containing this hex cell.
	IDH3Geo string `json:"idH3Geo"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The max received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMax float64 `json:"rpmMax"`
	// The mean received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMean float64 `json:"rpmMean"`
	// The median received power monitor (RPM) output value for the set of data
	// contained within this cell.
	RpmMedian float64 `json:"rpmMedian"`
	// The min received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMin float64 `json:"rpmMin"`
	// The standard deviation of the received power monitor (RPM) output value for the
	// set of data contained within this cell.
	RpmSigma float64 `json:"rpmSigma"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CellID                respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AltMean               respjson.Field
		AltSigma              respjson.Field
		AnomScoreInterference respjson.Field
		AnomScoreSpoofing     respjson.Field
		ChangeScore           respjson.Field
		Coverage              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDH3Geo               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RpmMax                respjson.Field
		RpmMean               respjson.Field
		RpmMedian             respjson.Field
		RpmMin                respjson.Field
		RpmSigma              respjson.Field
		SourceDl              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3GeoGetResponseCell) RawJSON() string { return r.JSON.raw }
func (r *H3GeoGetResponseCell) UnmarshalJSON(data []byte) error {
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
type H3GeoGetResponseDataMode string

const (
	H3GeoGetResponseDataModeReal      H3GeoGetResponseDataMode = "REAL"
	H3GeoGetResponseDataModeTest      H3GeoGetResponseDataMode = "TEST"
	H3GeoGetResponseDataModeSimulated H3GeoGetResponseDataMode = "SIMULATED"
	H3GeoGetResponseDataModeExercise  H3GeoGetResponseDataMode = "EXERCISE"
)

type H3GeoQueryhelpResponse struct {
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
func (r H3GeoQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *H3GeoQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// H3 Geospatial Binning is a discrete global grid system for indexing geographies
// into a hexagonal grid. The H3 schema is a collection of hexagonal cells that
// contain data for producing geospatial maps over a specified time span.
type H3GeoTupleResponse struct {
	// The collection of hex cells contained in this H3 data set. The number of cells
	// is a function of the specified resolution.
	Cells []H3GeoTupleResponseCell `json:"cells,required"`
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
	DataMode H3GeoTupleResponseDataMode `json:"dataMode,required"`
	// The number of cells associated with this H3 Geo data set. At this time, UDL
	// supports up to 50,000 cells.
	NumCells int64 `json:"numCells,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Start time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The center frequency of this H3 Geo data set measured in megahertz.
	CenterFreq float64 `json:"centerFreq"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// End time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// H3 resolution (0 – 15) for the data set. At this time, UDL supports a resolution
	// of 3 or less.
	Resolution int64 `json:"resolution"`
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
	// An optional field containing the type of data that is represented by this H3 Geo
	// data set.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cells                 respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		NumCells              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		CenterFreq            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndTime               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Resolution            respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3GeoTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *H3GeoTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a hex cell array containing data for a set of
// observations.
type H3GeoTupleResponseCell struct {
	// The H3 index represented as a 16 character hexadecimal string.
	CellID string `json:"cellId,required"`
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The mean altitude of the set of observations within this cell, measured in
	// kilometers.
	AltMean float64 `json:"altMean"`
	// The standard deviation of alttitude in the set of observations within this cell,
	// measured in kilometers.
	AltSigma float64 `json:"altSigma"`
	// The anomaly score for probable manufactured interference or RF interference;
	// calculated as a ratio of #anomalous obs / #total obs or coverage.
	AnomScoreInterference float64 `json:"anomScoreInterference"`
	// The anomaly score for probable spoofing; calculated as a ratio of #anomalous obs
	// / #total obs or coverage.
	AnomScoreSpoofing float64 `json:"anomScoreSpoofing"`
	// The percentage degree of change in the aggregated observables for a particular
	// H3 bin.
	ChangeScore float64 `json:"changeScore"`
	// The total number of available observations in the H3 cell during the start/end
	// times.
	Coverage int64 `json:"coverage"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent H3 Geo record containing this hex cell.
	IDH3Geo string `json:"idH3Geo"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The max received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMax float64 `json:"rpmMax"`
	// The mean received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMean float64 `json:"rpmMean"`
	// The median received power monitor (RPM) output value for the set of data
	// contained within this cell.
	RpmMedian float64 `json:"rpmMedian"`
	// The min received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMin float64 `json:"rpmMin"`
	// The standard deviation of the received power monitor (RPM) output value for the
	// set of data contained within this cell.
	RpmSigma float64 `json:"rpmSigma"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CellID                respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AltMean               respjson.Field
		AltSigma              respjson.Field
		AnomScoreInterference respjson.Field
		AnomScoreSpoofing     respjson.Field
		ChangeScore           respjson.Field
		Coverage              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDH3Geo               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RpmMax                respjson.Field
		RpmMean               respjson.Field
		RpmMedian             respjson.Field
		RpmMin                respjson.Field
		RpmSigma              respjson.Field
		SourceDl              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3GeoTupleResponseCell) RawJSON() string { return r.JSON.raw }
func (r *H3GeoTupleResponseCell) UnmarshalJSON(data []byte) error {
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
type H3GeoTupleResponseDataMode string

const (
	H3GeoTupleResponseDataModeReal      H3GeoTupleResponseDataMode = "REAL"
	H3GeoTupleResponseDataModeTest      H3GeoTupleResponseDataMode = "TEST"
	H3GeoTupleResponseDataModeSimulated H3GeoTupleResponseDataMode = "SIMULATED"
	H3GeoTupleResponseDataModeExercise  H3GeoTupleResponseDataMode = "EXERCISE"
)

type H3GeoNewParams struct {
	// The collection of hex cells contained in this H3 data set. The number of cells
	// is a function of the specified resolution.
	Cells []H3GeoNewParamsCell `json:"cells,omitzero,required"`
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
	DataMode H3GeoNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The number of cells associated with this H3 Geo data set. At this time, UDL
	// supports up to 50,000 cells.
	NumCells int64 `json:"numCells,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Start time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The center frequency of this H3 Geo data set measured in megahertz.
	CenterFreq param.Opt[float64] `json:"centerFreq,omitzero"`
	// End time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// H3 resolution (0 – 15) for the data set. At this time, UDL supports a resolution
	// of 3 or less.
	Resolution param.Opt[int64] `json:"resolution,omitzero"`
	// An optional field containing the type of data that is represented by this H3 Geo
	// data set.
	Type param.Opt[string] `json:"type,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r H3GeoNewParams) MarshalJSON() (data []byte, err error) {
	type shadow H3GeoNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *H3GeoNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a hex cell array containing data for a set of
// observations.
//
// The properties CellID, ClassificationMarking, DataMode, Source are required.
type H3GeoNewParamsCell struct {
	// The H3 index represented as a 16 character hexadecimal string.
	CellID string `json:"cellId,required"`
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
	// The mean altitude of the set of observations within this cell, measured in
	// kilometers.
	AltMean param.Opt[float64] `json:"altMean,omitzero"`
	// The standard deviation of alttitude in the set of observations within this cell,
	// measured in kilometers.
	AltSigma param.Opt[float64] `json:"altSigma,omitzero"`
	// The anomaly score for probable manufactured interference or RF interference;
	// calculated as a ratio of #anomalous obs / #total obs or coverage.
	AnomScoreInterference param.Opt[float64] `json:"anomScoreInterference,omitzero"`
	// The anomaly score for probable spoofing; calculated as a ratio of #anomalous obs
	// / #total obs or coverage.
	AnomScoreSpoofing param.Opt[float64] `json:"anomScoreSpoofing,omitzero"`
	// The percentage degree of change in the aggregated observables for a particular
	// H3 bin.
	ChangeScore param.Opt[float64] `json:"changeScore,omitzero"`
	// The total number of available observations in the H3 cell during the start/end
	// times.
	Coverage param.Opt[int64] `json:"coverage,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of the parent H3 Geo record containing this hex cell.
	IDH3Geo param.Opt[string] `json:"idH3Geo,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The max received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMax param.Opt[float64] `json:"rpmMax,omitzero"`
	// The mean received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMean param.Opt[float64] `json:"rpmMean,omitzero"`
	// The median received power monitor (RPM) output value for the set of data
	// contained within this cell.
	RpmMedian param.Opt[float64] `json:"rpmMedian,omitzero"`
	// The min received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMin param.Opt[float64] `json:"rpmMin,omitzero"`
	// The standard deviation of the received power monitor (RPM) output value for the
	// set of data contained within this cell.
	RpmSigma param.Opt[float64] `json:"rpmSigma,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	paramObj
}

func (r H3GeoNewParamsCell) MarshalJSON() (data []byte, err error) {
	type shadow H3GeoNewParamsCell
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *H3GeoNewParamsCell) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[H3GeoNewParamsCell](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
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
type H3GeoNewParamsDataMode string

const (
	H3GeoNewParamsDataModeReal      H3GeoNewParamsDataMode = "REAL"
	H3GeoNewParamsDataModeTest      H3GeoNewParamsDataMode = "TEST"
	H3GeoNewParamsDataModeSimulated H3GeoNewParamsDataMode = "SIMULATED"
	H3GeoNewParamsDataModeExercise  H3GeoNewParamsDataMode = "EXERCISE"
)

type H3GeoListParams struct {
	// Start time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [H3GeoListParams]'s query parameters as `url.Values`.
func (r H3GeoListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type H3GeoCountParams struct {
	// Start time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [H3GeoCountParams]'s query parameters as `url.Values`.
func (r H3GeoCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type H3GeoGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [H3GeoGetParams]'s query parameters as `url.Values`.
func (r H3GeoGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type H3GeoTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Start time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [H3GeoTupleParams]'s query parameters as `url.Values`.
func (r H3GeoTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
