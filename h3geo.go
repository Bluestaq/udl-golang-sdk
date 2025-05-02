// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// H3geoService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewH3geoService] method instead.
type H3geoService struct {
	Options []option.RequestOption
	History H3geoHistoryService
}

// NewH3geoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewH3geoService(opts ...option.RequestOption) (r H3geoService) {
	r = H3geoService{}
	r.Options = opts
	r.History = NewH3geoHistoryService(opts...)
	return
}

// Service operation to take a single H3Geo record as a POST body and ingest into
// the database. This operation does not persist any H3GeoHexCell points that may
// be present in the body of the request. This operation is not intended to be used
// for automated feeds into UDL. Data providers should contact the UDL team for
// specific role assignments and for instructions on setting up a permanent feed
// through an alternate mechanism.
func (r *H3geoService) New(ctx context.Context, body H3geoNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/h3geo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *H3geoService) List(ctx context.Context, query H3geoListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[H3geoListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *H3geoService) ListAutoPaging(ctx context.Context, query H3geoListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[H3geoListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *H3geoService) Count(ctx context.Context, query H3geoCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/h3geo/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single RF geolocation by its unique ID passed as a
// path parameter.
func (r *H3geoService) Get(ctx context.Context, id string, query H3geoGetParams, opts ...option.RequestOption) (res *H3geoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *H3geoService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/h3geo/queryhelp"
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
func (r *H3geoService) Tuple(ctx context.Context, query H3geoTupleParams, opts ...option.RequestOption) (res *[]H3geoTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/h3geo/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// H3 Geospatial Binning is a discrete global grid system for indexing geographies
// into a hexagonal grid. The H3 schema is a collection of hexagonal cells that
// contain data for producing geospatial maps over a specified time span.
type H3geoListResponse struct {
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
	DataMode H3geoListResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		NumCells              resp.Field
		Source                resp.Field
		StartTime             resp.Field
		ID                    resp.Field
		CenterFreq            resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EndTime               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Resolution            resp.Field
		SourceDl              resp.Field
		Type                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3geoListResponse) RawJSON() string { return r.JSON.raw }
func (r *H3geoListResponse) UnmarshalJSON(data []byte) error {
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
type H3geoListResponseDataMode string

const (
	H3geoListResponseDataModeReal      H3geoListResponseDataMode = "REAL"
	H3geoListResponseDataModeTest      H3geoListResponseDataMode = "TEST"
	H3geoListResponseDataModeSimulated H3geoListResponseDataMode = "SIMULATED"
	H3geoListResponseDataModeExercise  H3geoListResponseDataMode = "EXERCISE"
)

// H3 Geospatial Binning is a discrete global grid system for indexing geographies
// into a hexagonal grid. The H3 schema is a collection of hexagonal cells that
// contain data for producing geospatial maps over a specified time span.
type H3geoGetResponse struct {
	// The collection of hex cells contained in this H3 data set. The number of cells
	// is a function of the specified resolution.
	Cells []H3geoGetResponseCell `json:"cells,required"`
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
	DataMode H3geoGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Cells                 resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		NumCells              resp.Field
		Source                resp.Field
		StartTime             resp.Field
		ID                    resp.Field
		CenterFreq            resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EndTime               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Resolution            resp.Field
		SourceDl              resp.Field
		Tags                  resp.Field
		Type                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3geoGetResponse) RawJSON() string { return r.JSON.raw }
func (r *H3geoGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a hex cell array containing data for a set of
// observations.
type H3geoGetResponseCell struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CellID                resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Source                resp.Field
		ID                    resp.Field
		AltMean               resp.Field
		AltSigma              resp.Field
		AnomScoreInterference resp.Field
		AnomScoreSpoofing     resp.Field
		ChangeScore           resp.Field
		Coverage              resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		IDH3Geo               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		RpmMax                resp.Field
		RpmMean               resp.Field
		RpmMedian             resp.Field
		RpmMin                resp.Field
		RpmSigma              resp.Field
		SourceDl              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3geoGetResponseCell) RawJSON() string { return r.JSON.raw }
func (r *H3geoGetResponseCell) UnmarshalJSON(data []byte) error {
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
type H3geoGetResponseDataMode string

const (
	H3geoGetResponseDataModeReal      H3geoGetResponseDataMode = "REAL"
	H3geoGetResponseDataModeTest      H3geoGetResponseDataMode = "TEST"
	H3geoGetResponseDataModeSimulated H3geoGetResponseDataMode = "SIMULATED"
	H3geoGetResponseDataModeExercise  H3geoGetResponseDataMode = "EXERCISE"
)

// H3 Geospatial Binning is a discrete global grid system for indexing geographies
// into a hexagonal grid. The H3 schema is a collection of hexagonal cells that
// contain data for producing geospatial maps over a specified time span.
type H3geoTupleResponse struct {
	// The collection of hex cells contained in this H3 data set. The number of cells
	// is a function of the specified resolution.
	Cells []H3geoTupleResponseCell `json:"cells,required"`
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
	DataMode H3geoTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Cells                 resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		NumCells              resp.Field
		Source                resp.Field
		StartTime             resp.Field
		ID                    resp.Field
		CenterFreq            resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EndTime               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Resolution            resp.Field
		SourceDl              resp.Field
		Tags                  resp.Field
		Type                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3geoTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *H3geoTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a hex cell array containing data for a set of
// observations.
type H3geoTupleResponseCell struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CellID                resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Source                resp.Field
		ID                    resp.Field
		AltMean               resp.Field
		AltSigma              resp.Field
		AnomScoreInterference resp.Field
		AnomScoreSpoofing     resp.Field
		ChangeScore           resp.Field
		Coverage              resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		IDH3Geo               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		RpmMax                resp.Field
		RpmMean               resp.Field
		RpmMedian             resp.Field
		RpmMin                resp.Field
		RpmSigma              resp.Field
		SourceDl              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3geoTupleResponseCell) RawJSON() string { return r.JSON.raw }
func (r *H3geoTupleResponseCell) UnmarshalJSON(data []byte) error {
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
type H3geoTupleResponseDataMode string

const (
	H3geoTupleResponseDataModeReal      H3geoTupleResponseDataMode = "REAL"
	H3geoTupleResponseDataModeTest      H3geoTupleResponseDataMode = "TEST"
	H3geoTupleResponseDataModeSimulated H3geoTupleResponseDataMode = "SIMULATED"
	H3geoTupleResponseDataModeExercise  H3geoTupleResponseDataMode = "EXERCISE"
)

type H3geoNewParams struct {
	// The collection of hex cells contained in this H3 data set. The number of cells
	// is a function of the specified resolution.
	Cells []H3geoNewParamsCell `json:"cells,omitzero,required"`
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
	DataMode H3geoNewParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f H3geoNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r H3geoNewParams) MarshalJSON() (data []byte, err error) {
	type shadow H3geoNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Model representation of a hex cell array containing data for a set of
// observations.
//
// The properties CellID, ClassificationMarking, DataMode, Source are required.
type H3geoNewParamsCell struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f H3geoNewParamsCell) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r H3geoNewParamsCell) MarshalJSON() (data []byte, err error) {
	type shadow H3geoNewParamsCell
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[H3geoNewParamsCell](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
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
type H3geoNewParamsDataMode string

const (
	H3geoNewParamsDataModeReal      H3geoNewParamsDataMode = "REAL"
	H3geoNewParamsDataModeTest      H3geoNewParamsDataMode = "TEST"
	H3geoNewParamsDataModeSimulated H3geoNewParamsDataMode = "SIMULATED"
	H3geoNewParamsDataModeExercise  H3geoNewParamsDataMode = "EXERCISE"
)

type H3geoListParams struct {
	// Start time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f H3geoListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [H3geoListParams]'s query parameters as `url.Values`.
func (r H3geoListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type H3geoCountParams struct {
	// Start time for this H3 Geo data set in ISO 8601 UTC with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f H3geoCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [H3geoCountParams]'s query parameters as `url.Values`.
func (r H3geoCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type H3geoGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f H3geoGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [H3geoGetParams]'s query parameters as `url.Values`.
func (r H3geoGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type H3geoTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f H3geoTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [H3geoTupleParams]'s query parameters as `url.Values`.
func (r H3geoTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
