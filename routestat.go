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

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// RouteStatService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRouteStatService] method instead.
type RouteStatService struct {
	Options []option.RequestOption
}

// NewRouteStatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRouteStatService(opts ...option.RequestOption) (r RouteStatService) {
	r = RouteStatService{}
	r.Options = opts
	return
}

// Service operation to take a single routeStats record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *RouteStatService) New(ctx context.Context, body RouteStatNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/routestats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single routeStats record by its unique ID passed as a
// path parameter.
func (r *RouteStatService) Get(ctx context.Context, id string, query RouteStatGetParams, opts ...option.RequestOption) (res *RouteStatGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/routestats/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single RouteStats. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *RouteStatService) Update(ctx context.Context, id string, body RouteStatUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/routestats/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to delete a routeStats record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RouteStatService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/routestats/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *RouteStatService) Count(ctx context.Context, query RouteStatCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/routestats/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// RouteStats as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *RouteStatService) NewBulk(ctx context.Context, body RouteStatNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/routestats/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *RouteStatService) Query(ctx context.Context, query RouteStatQueryParams, opts ...option.RequestOption) (res *[]RouteStatQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/routestats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *RouteStatService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/routestats/queryhelp"
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
func (r *RouteStatService) Tuple(ctx context.Context, query RouteStatTupleParams, opts ...option.RequestOption) (res *[]RouteStatTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/routestats/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple routestats records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RouteStatService) UnvalidatedPublish(ctx context.Context, body RouteStatUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-routestats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// General statistics applying to navigation routes utilized by vessels, aircraft,
// ground vehicles, etc.
type RouteStatGetResponse struct {
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
	DataMode RouteStatGetResponseDataMode `json:"dataMode,required"`
	// End location of the vehicle.
	LocationEnd string `json:"locationEnd,required"`
	// Starting location of the vehicle.
	LocationStart string `json:"locationStart,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Average travel duration for the indicated distance and type of vehicle in hours.
	AvgDuration float64 `json:"avgDuration"`
	// Average speed during travel in the indicated unit of measurement, speedUnit.
	AvgSpeed float64 `json:"avgSpeed"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The number of data points used in this travel duration calculation.
	DataPtsUsed int64 `json:"dataPtsUsed"`
	// Distance between the departure and arrival locations in the indicated unit of
	// measurement, distUnit.
	Distance float64 `json:"distance"`
	// The unit of measurement used for distance in this calculation.
	DistUnit string `json:"distUnit"`
	// Date of the first data point used in this calculation, in ISO8601 UTC format
	// with millisecond precision.
	FirstPt time.Time `json:"firstPt" format:"date-time"`
	// Description of the portion of travel used to estimate the value of the
	// idealDuration field.
	IdealDesc string `json:"idealDesc"`
	// Estimated ideal travel duration in hours for the full distance using the
	// indicated vehicle type. The field "idealDesc" should be used to describe the
	// ideal travel route.
	IdealDuration float64 `json:"idealDuration"`
	// Unique identifier of the Site at the route's end location. This ID can be used
	// to obtain additional information on a Site using the 'get by ID' operation (e.g.
	// /udl/site/{id}). For example, the Site object with idSite = abc would be queried
	// as /udl/site/abc.
	IDSiteEnd string `json:"idSiteEnd"`
	// Unique identifier of the Site at the route's starting location. This ID can be
	// used to obtain additional information on a Site using the 'get by ID' operation
	// (e.g. /udl/site/{id}). For example, the Site object with idSite = abc would be
	// queried as /udl/site/abc.
	IDSiteStart string `json:"idSiteStart"`
	// Date of the last data point used in this calculation, in ISO8601 UTC format with
	// millisecond precision.
	LastPt time.Time `json:"lastPt" format:"date-time"`
	// Type of location used for route start and end points (e.g., ICAO, PORT, etc.).
	LocationType string `json:"locationType"`
	// Maximum travel duration for the indicated distance and type of vehicle in hours.
	MaxDuration float64 `json:"maxDuration"`
	// Maximum speed during travel in the indicated unit of measurement, speedUnit.
	MaxSpeed float64 `json:"maxSpeed"`
	// Minimum travel duration for the indicated distance and type of vehicle in hours.
	MinDuration float64 `json:"minDuration"`
	// Minimum speed during travel in the indicated unit of measurement, speedUnit.
	MinSpeed float64 `json:"minSpeed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Description of the portion of travel used to estimate the value of the
	// partialDuration field.
	PartialDesc string `json:"partialDesc"`
	// Estimated ideal travel duration in hours for a partial distance using the
	// indicated vehicle type. The field "partialDesc" should be used to specify the
	// intended portion of travel.
	PartialDuration float64 `json:"partialDuration"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The unit of measurement used for speed in this calculation.
	SpeedUnit string `json:"speedUnit"`
	// The time period this data was collected.
	TimePeriod string `json:"timePeriod"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The vehicle category that is the subject of this calculation (e.g., AIRCRAFT,
	// CAR, BOAT, etc.).
	VehicleCategory string `json:"vehicleCategory"`
	// The vehicle type that is the subject of this calculation (e.g., C-17, F-15,
	// etc.).
	VehicleType string `json:"vehicleType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		LocationEnd           respjson.Field
		LocationStart         respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AvgDuration           respjson.Field
		AvgSpeed              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataPtsUsed           respjson.Field
		Distance              respjson.Field
		DistUnit              respjson.Field
		FirstPt               respjson.Field
		IdealDesc             respjson.Field
		IdealDuration         respjson.Field
		IDSiteEnd             respjson.Field
		IDSiteStart           respjson.Field
		LastPt                respjson.Field
		LocationType          respjson.Field
		MaxDuration           respjson.Field
		MaxSpeed              respjson.Field
		MinDuration           respjson.Field
		MinSpeed              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PartialDesc           respjson.Field
		PartialDuration       respjson.Field
		SourceDl              respjson.Field
		SpeedUnit             respjson.Field
		TimePeriod            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VehicleCategory       respjson.Field
		VehicleType           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteStatGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RouteStatGetResponse) UnmarshalJSON(data []byte) error {
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
type RouteStatGetResponseDataMode string

const (
	RouteStatGetResponseDataModeReal      RouteStatGetResponseDataMode = "REAL"
	RouteStatGetResponseDataModeTest      RouteStatGetResponseDataMode = "TEST"
	RouteStatGetResponseDataModeSimulated RouteStatGetResponseDataMode = "SIMULATED"
	RouteStatGetResponseDataModeExercise  RouteStatGetResponseDataMode = "EXERCISE"
)

// General statistics applying to navigation routes utilized by vessels, aircraft,
// ground vehicles, etc.
type RouteStatQueryResponse struct {
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
	DataMode RouteStatQueryResponseDataMode `json:"dataMode,required"`
	// End location of the vehicle.
	LocationEnd string `json:"locationEnd,required"`
	// Starting location of the vehicle.
	LocationStart string `json:"locationStart,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Average travel duration for the indicated distance and type of vehicle in hours.
	AvgDuration float64 `json:"avgDuration"`
	// Average speed during travel in the indicated unit of measurement, speedUnit.
	AvgSpeed float64 `json:"avgSpeed"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The number of data points used in this travel duration calculation.
	DataPtsUsed int64 `json:"dataPtsUsed"`
	// Distance between the departure and arrival locations in the indicated unit of
	// measurement, distUnit.
	Distance float64 `json:"distance"`
	// The unit of measurement used for distance in this calculation.
	DistUnit string `json:"distUnit"`
	// Date of the first data point used in this calculation, in ISO8601 UTC format
	// with millisecond precision.
	FirstPt time.Time `json:"firstPt" format:"date-time"`
	// Description of the portion of travel used to estimate the value of the
	// idealDuration field.
	IdealDesc string `json:"idealDesc"`
	// Estimated ideal travel duration in hours for the full distance using the
	// indicated vehicle type. The field "idealDesc" should be used to describe the
	// ideal travel route.
	IdealDuration float64 `json:"idealDuration"`
	// Unique identifier of the Site at the route's end location. This ID can be used
	// to obtain additional information on a Site using the 'get by ID' operation (e.g.
	// /udl/site/{id}). For example, the Site object with idSite = abc would be queried
	// as /udl/site/abc.
	IDSiteEnd string `json:"idSiteEnd"`
	// Unique identifier of the Site at the route's starting location. This ID can be
	// used to obtain additional information on a Site using the 'get by ID' operation
	// (e.g. /udl/site/{id}). For example, the Site object with idSite = abc would be
	// queried as /udl/site/abc.
	IDSiteStart string `json:"idSiteStart"`
	// Date of the last data point used in this calculation, in ISO8601 UTC format with
	// millisecond precision.
	LastPt time.Time `json:"lastPt" format:"date-time"`
	// Type of location used for route start and end points (e.g., ICAO, PORT, etc.).
	LocationType string `json:"locationType"`
	// Maximum travel duration for the indicated distance and type of vehicle in hours.
	MaxDuration float64 `json:"maxDuration"`
	// Maximum speed during travel in the indicated unit of measurement, speedUnit.
	MaxSpeed float64 `json:"maxSpeed"`
	// Minimum travel duration for the indicated distance and type of vehicle in hours.
	MinDuration float64 `json:"minDuration"`
	// Minimum speed during travel in the indicated unit of measurement, speedUnit.
	MinSpeed float64 `json:"minSpeed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Description of the portion of travel used to estimate the value of the
	// partialDuration field.
	PartialDesc string `json:"partialDesc"`
	// Estimated ideal travel duration in hours for a partial distance using the
	// indicated vehicle type. The field "partialDesc" should be used to specify the
	// intended portion of travel.
	PartialDuration float64 `json:"partialDuration"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The unit of measurement used for speed in this calculation.
	SpeedUnit string `json:"speedUnit"`
	// The time period this data was collected.
	TimePeriod string `json:"timePeriod"`
	// The vehicle category that is the subject of this calculation (e.g., AIRCRAFT,
	// CAR, BOAT, etc.).
	VehicleCategory string `json:"vehicleCategory"`
	// The vehicle type that is the subject of this calculation (e.g., C-17, F-15,
	// etc.).
	VehicleType string `json:"vehicleType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		LocationEnd           respjson.Field
		LocationStart         respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AvgDuration           respjson.Field
		AvgSpeed              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataPtsUsed           respjson.Field
		Distance              respjson.Field
		DistUnit              respjson.Field
		FirstPt               respjson.Field
		IdealDesc             respjson.Field
		IdealDuration         respjson.Field
		IDSiteEnd             respjson.Field
		IDSiteStart           respjson.Field
		LastPt                respjson.Field
		LocationType          respjson.Field
		MaxDuration           respjson.Field
		MaxSpeed              respjson.Field
		MinDuration           respjson.Field
		MinSpeed              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PartialDesc           respjson.Field
		PartialDuration       respjson.Field
		SourceDl              respjson.Field
		SpeedUnit             respjson.Field
		TimePeriod            respjson.Field
		VehicleCategory       respjson.Field
		VehicleType           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteStatQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *RouteStatQueryResponse) UnmarshalJSON(data []byte) error {
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
type RouteStatQueryResponseDataMode string

const (
	RouteStatQueryResponseDataModeReal      RouteStatQueryResponseDataMode = "REAL"
	RouteStatQueryResponseDataModeTest      RouteStatQueryResponseDataMode = "TEST"
	RouteStatQueryResponseDataModeSimulated RouteStatQueryResponseDataMode = "SIMULATED"
	RouteStatQueryResponseDataModeExercise  RouteStatQueryResponseDataMode = "EXERCISE"
)

// General statistics applying to navigation routes utilized by vessels, aircraft,
// ground vehicles, etc.
type RouteStatTupleResponse struct {
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
	DataMode RouteStatTupleResponseDataMode `json:"dataMode,required"`
	// End location of the vehicle.
	LocationEnd string `json:"locationEnd,required"`
	// Starting location of the vehicle.
	LocationStart string `json:"locationStart,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Average travel duration for the indicated distance and type of vehicle in hours.
	AvgDuration float64 `json:"avgDuration"`
	// Average speed during travel in the indicated unit of measurement, speedUnit.
	AvgSpeed float64 `json:"avgSpeed"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The number of data points used in this travel duration calculation.
	DataPtsUsed int64 `json:"dataPtsUsed"`
	// Distance between the departure and arrival locations in the indicated unit of
	// measurement, distUnit.
	Distance float64 `json:"distance"`
	// The unit of measurement used for distance in this calculation.
	DistUnit string `json:"distUnit"`
	// Date of the first data point used in this calculation, in ISO8601 UTC format
	// with millisecond precision.
	FirstPt time.Time `json:"firstPt" format:"date-time"`
	// Description of the portion of travel used to estimate the value of the
	// idealDuration field.
	IdealDesc string `json:"idealDesc"`
	// Estimated ideal travel duration in hours for the full distance using the
	// indicated vehicle type. The field "idealDesc" should be used to describe the
	// ideal travel route.
	IdealDuration float64 `json:"idealDuration"`
	// Unique identifier of the Site at the route's end location. This ID can be used
	// to obtain additional information on a Site using the 'get by ID' operation (e.g.
	// /udl/site/{id}). For example, the Site object with idSite = abc would be queried
	// as /udl/site/abc.
	IDSiteEnd string `json:"idSiteEnd"`
	// Unique identifier of the Site at the route's starting location. This ID can be
	// used to obtain additional information on a Site using the 'get by ID' operation
	// (e.g. /udl/site/{id}). For example, the Site object with idSite = abc would be
	// queried as /udl/site/abc.
	IDSiteStart string `json:"idSiteStart"`
	// Date of the last data point used in this calculation, in ISO8601 UTC format with
	// millisecond precision.
	LastPt time.Time `json:"lastPt" format:"date-time"`
	// Type of location used for route start and end points (e.g., ICAO, PORT, etc.).
	LocationType string `json:"locationType"`
	// Maximum travel duration for the indicated distance and type of vehicle in hours.
	MaxDuration float64 `json:"maxDuration"`
	// Maximum speed during travel in the indicated unit of measurement, speedUnit.
	MaxSpeed float64 `json:"maxSpeed"`
	// Minimum travel duration for the indicated distance and type of vehicle in hours.
	MinDuration float64 `json:"minDuration"`
	// Minimum speed during travel in the indicated unit of measurement, speedUnit.
	MinSpeed float64 `json:"minSpeed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Description of the portion of travel used to estimate the value of the
	// partialDuration field.
	PartialDesc string `json:"partialDesc"`
	// Estimated ideal travel duration in hours for a partial distance using the
	// indicated vehicle type. The field "partialDesc" should be used to specify the
	// intended portion of travel.
	PartialDuration float64 `json:"partialDuration"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The unit of measurement used for speed in this calculation.
	SpeedUnit string `json:"speedUnit"`
	// The time period this data was collected.
	TimePeriod string `json:"timePeriod"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The vehicle category that is the subject of this calculation (e.g., AIRCRAFT,
	// CAR, BOAT, etc.).
	VehicleCategory string `json:"vehicleCategory"`
	// The vehicle type that is the subject of this calculation (e.g., C-17, F-15,
	// etc.).
	VehicleType string `json:"vehicleType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		LocationEnd           respjson.Field
		LocationStart         respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AvgDuration           respjson.Field
		AvgSpeed              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataPtsUsed           respjson.Field
		Distance              respjson.Field
		DistUnit              respjson.Field
		FirstPt               respjson.Field
		IdealDesc             respjson.Field
		IdealDuration         respjson.Field
		IDSiteEnd             respjson.Field
		IDSiteStart           respjson.Field
		LastPt                respjson.Field
		LocationType          respjson.Field
		MaxDuration           respjson.Field
		MaxSpeed              respjson.Field
		MinDuration           respjson.Field
		MinSpeed              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PartialDesc           respjson.Field
		PartialDuration       respjson.Field
		SourceDl              respjson.Field
		SpeedUnit             respjson.Field
		TimePeriod            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VehicleCategory       respjson.Field
		VehicleType           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteStatTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *RouteStatTupleResponse) UnmarshalJSON(data []byte) error {
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
type RouteStatTupleResponseDataMode string

const (
	RouteStatTupleResponseDataModeReal      RouteStatTupleResponseDataMode = "REAL"
	RouteStatTupleResponseDataModeTest      RouteStatTupleResponseDataMode = "TEST"
	RouteStatTupleResponseDataModeSimulated RouteStatTupleResponseDataMode = "SIMULATED"
	RouteStatTupleResponseDataModeExercise  RouteStatTupleResponseDataMode = "EXERCISE"
)

type RouteStatNewParams struct {
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
	DataMode RouteStatNewParamsDataMode `json:"dataMode,omitzero,required"`
	// End location of the vehicle.
	LocationEnd string `json:"locationEnd,required"`
	// Starting location of the vehicle.
	LocationStart string `json:"locationStart,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Average travel duration for the indicated distance and type of vehicle in hours.
	AvgDuration param.Opt[float64] `json:"avgDuration,omitzero"`
	// Average speed during travel in the indicated unit of measurement, speedUnit.
	AvgSpeed param.Opt[float64] `json:"avgSpeed,omitzero"`
	// The number of data points used in this travel duration calculation.
	DataPtsUsed param.Opt[int64] `json:"dataPtsUsed,omitzero"`
	// Distance between the departure and arrival locations in the indicated unit of
	// measurement, distUnit.
	Distance param.Opt[float64] `json:"distance,omitzero"`
	// The unit of measurement used for distance in this calculation.
	DistUnit param.Opt[string] `json:"distUnit,omitzero"`
	// Date of the first data point used in this calculation, in ISO8601 UTC format
	// with millisecond precision.
	FirstPt param.Opt[time.Time] `json:"firstPt,omitzero" format:"date-time"`
	// Description of the portion of travel used to estimate the value of the
	// idealDuration field.
	IdealDesc param.Opt[string] `json:"idealDesc,omitzero"`
	// Estimated ideal travel duration in hours for the full distance using the
	// indicated vehicle type. The field "idealDesc" should be used to describe the
	// ideal travel route.
	IdealDuration param.Opt[float64] `json:"idealDuration,omitzero"`
	// Unique identifier of the Site at the route's end location. This ID can be used
	// to obtain additional information on a Site using the 'get by ID' operation (e.g.
	// /udl/site/{id}). For example, the Site object with idSite = abc would be queried
	// as /udl/site/abc.
	IDSiteEnd param.Opt[string] `json:"idSiteEnd,omitzero"`
	// Unique identifier of the Site at the route's starting location. This ID can be
	// used to obtain additional information on a Site using the 'get by ID' operation
	// (e.g. /udl/site/{id}). For example, the Site object with idSite = abc would be
	// queried as /udl/site/abc.
	IDSiteStart param.Opt[string] `json:"idSiteStart,omitzero"`
	// Date of the last data point used in this calculation, in ISO8601 UTC format with
	// millisecond precision.
	LastPt param.Opt[time.Time] `json:"lastPt,omitzero" format:"date-time"`
	// Type of location used for route start and end points (e.g., ICAO, PORT, etc.).
	LocationType param.Opt[string] `json:"locationType,omitzero"`
	// Maximum travel duration for the indicated distance and type of vehicle in hours.
	MaxDuration param.Opt[float64] `json:"maxDuration,omitzero"`
	// Maximum speed during travel in the indicated unit of measurement, speedUnit.
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// Minimum travel duration for the indicated distance and type of vehicle in hours.
	MinDuration param.Opt[float64] `json:"minDuration,omitzero"`
	// Minimum speed during travel in the indicated unit of measurement, speedUnit.
	MinSpeed param.Opt[float64] `json:"minSpeed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Description of the portion of travel used to estimate the value of the
	// partialDuration field.
	PartialDesc param.Opt[string] `json:"partialDesc,omitzero"`
	// Estimated ideal travel duration in hours for a partial distance using the
	// indicated vehicle type. The field "partialDesc" should be used to specify the
	// intended portion of travel.
	PartialDuration param.Opt[float64] `json:"partialDuration,omitzero"`
	// The unit of measurement used for speed in this calculation.
	SpeedUnit param.Opt[string] `json:"speedUnit,omitzero"`
	// The time period this data was collected.
	TimePeriod param.Opt[string] `json:"timePeriod,omitzero"`
	// The vehicle category that is the subject of this calculation (e.g., AIRCRAFT,
	// CAR, BOAT, etc.).
	VehicleCategory param.Opt[string] `json:"vehicleCategory,omitzero"`
	// The vehicle type that is the subject of this calculation (e.g., C-17, F-15,
	// etc.).
	VehicleType param.Opt[string] `json:"vehicleType,omitzero"`
	paramObj
}

func (r RouteStatNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RouteStatNewParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type RouteStatNewParamsDataMode string

const (
	RouteStatNewParamsDataModeReal      RouteStatNewParamsDataMode = "REAL"
	RouteStatNewParamsDataModeTest      RouteStatNewParamsDataMode = "TEST"
	RouteStatNewParamsDataModeSimulated RouteStatNewParamsDataMode = "SIMULATED"
	RouteStatNewParamsDataModeExercise  RouteStatNewParamsDataMode = "EXERCISE"
)

type RouteStatGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RouteStatGetParams]'s query parameters as `url.Values`.
func (r RouteStatGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RouteStatUpdateParams struct {
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
	DataMode RouteStatUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// End location of the vehicle.
	LocationEnd string `json:"locationEnd,required"`
	// Starting location of the vehicle.
	LocationStart string `json:"locationStart,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Average travel duration for the indicated distance and type of vehicle in hours.
	AvgDuration param.Opt[float64] `json:"avgDuration,omitzero"`
	// Average speed during travel in the indicated unit of measurement, speedUnit.
	AvgSpeed param.Opt[float64] `json:"avgSpeed,omitzero"`
	// The number of data points used in this travel duration calculation.
	DataPtsUsed param.Opt[int64] `json:"dataPtsUsed,omitzero"`
	// Distance between the departure and arrival locations in the indicated unit of
	// measurement, distUnit.
	Distance param.Opt[float64] `json:"distance,omitzero"`
	// The unit of measurement used for distance in this calculation.
	DistUnit param.Opt[string] `json:"distUnit,omitzero"`
	// Date of the first data point used in this calculation, in ISO8601 UTC format
	// with millisecond precision.
	FirstPt param.Opt[time.Time] `json:"firstPt,omitzero" format:"date-time"`
	// Description of the portion of travel used to estimate the value of the
	// idealDuration field.
	IdealDesc param.Opt[string] `json:"idealDesc,omitzero"`
	// Estimated ideal travel duration in hours for the full distance using the
	// indicated vehicle type. The field "idealDesc" should be used to describe the
	// ideal travel route.
	IdealDuration param.Opt[float64] `json:"idealDuration,omitzero"`
	// Unique identifier of the Site at the route's end location. This ID can be used
	// to obtain additional information on a Site using the 'get by ID' operation (e.g.
	// /udl/site/{id}). For example, the Site object with idSite = abc would be queried
	// as /udl/site/abc.
	IDSiteEnd param.Opt[string] `json:"idSiteEnd,omitzero"`
	// Unique identifier of the Site at the route's starting location. This ID can be
	// used to obtain additional information on a Site using the 'get by ID' operation
	// (e.g. /udl/site/{id}). For example, the Site object with idSite = abc would be
	// queried as /udl/site/abc.
	IDSiteStart param.Opt[string] `json:"idSiteStart,omitzero"`
	// Date of the last data point used in this calculation, in ISO8601 UTC format with
	// millisecond precision.
	LastPt param.Opt[time.Time] `json:"lastPt,omitzero" format:"date-time"`
	// Type of location used for route start and end points (e.g., ICAO, PORT, etc.).
	LocationType param.Opt[string] `json:"locationType,omitzero"`
	// Maximum travel duration for the indicated distance and type of vehicle in hours.
	MaxDuration param.Opt[float64] `json:"maxDuration,omitzero"`
	// Maximum speed during travel in the indicated unit of measurement, speedUnit.
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// Minimum travel duration for the indicated distance and type of vehicle in hours.
	MinDuration param.Opt[float64] `json:"minDuration,omitzero"`
	// Minimum speed during travel in the indicated unit of measurement, speedUnit.
	MinSpeed param.Opt[float64] `json:"minSpeed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Description of the portion of travel used to estimate the value of the
	// partialDuration field.
	PartialDesc param.Opt[string] `json:"partialDesc,omitzero"`
	// Estimated ideal travel duration in hours for a partial distance using the
	// indicated vehicle type. The field "partialDesc" should be used to specify the
	// intended portion of travel.
	PartialDuration param.Opt[float64] `json:"partialDuration,omitzero"`
	// The unit of measurement used for speed in this calculation.
	SpeedUnit param.Opt[string] `json:"speedUnit,omitzero"`
	// The time period this data was collected.
	TimePeriod param.Opt[string] `json:"timePeriod,omitzero"`
	// The vehicle category that is the subject of this calculation (e.g., AIRCRAFT,
	// CAR, BOAT, etc.).
	VehicleCategory param.Opt[string] `json:"vehicleCategory,omitzero"`
	// The vehicle type that is the subject of this calculation (e.g., C-17, F-15,
	// etc.).
	VehicleType param.Opt[string] `json:"vehicleType,omitzero"`
	paramObj
}

func (r RouteStatUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RouteStatUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type RouteStatUpdateParamsDataMode string

const (
	RouteStatUpdateParamsDataModeReal      RouteStatUpdateParamsDataMode = "REAL"
	RouteStatUpdateParamsDataModeTest      RouteStatUpdateParamsDataMode = "TEST"
	RouteStatUpdateParamsDataModeSimulated RouteStatUpdateParamsDataMode = "SIMULATED"
	RouteStatUpdateParamsDataModeExercise  RouteStatUpdateParamsDataMode = "EXERCISE"
)

type RouteStatCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RouteStatCountParams]'s query parameters as `url.Values`.
func (r RouteStatCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RouteStatNewBulkParams struct {
	Body []RouteStatNewBulkParamsBody
	paramObj
}

func (r RouteStatNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// General statistics applying to navigation routes utilized by vessels, aircraft,
// ground vehicles, etc.
//
// The properties ClassificationMarking, DataMode, LocationEnd, LocationStart,
// Source are required.
type RouteStatNewBulkParamsBody struct {
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
	// End location of the vehicle.
	LocationEnd string `json:"locationEnd,required"`
	// Starting location of the vehicle.
	LocationStart string `json:"locationStart,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Average travel duration for the indicated distance and type of vehicle in hours.
	AvgDuration param.Opt[float64] `json:"avgDuration,omitzero"`
	// Average speed during travel in the indicated unit of measurement, speedUnit.
	AvgSpeed param.Opt[float64] `json:"avgSpeed,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The number of data points used in this travel duration calculation.
	DataPtsUsed param.Opt[int64] `json:"dataPtsUsed,omitzero"`
	// Distance between the departure and arrival locations in the indicated unit of
	// measurement, distUnit.
	Distance param.Opt[float64] `json:"distance,omitzero"`
	// The unit of measurement used for distance in this calculation.
	DistUnit param.Opt[string] `json:"distUnit,omitzero"`
	// Date of the first data point used in this calculation, in ISO8601 UTC format
	// with millisecond precision.
	FirstPt param.Opt[time.Time] `json:"firstPt,omitzero" format:"date-time"`
	// Description of the portion of travel used to estimate the value of the
	// idealDuration field.
	IdealDesc param.Opt[string] `json:"idealDesc,omitzero"`
	// Estimated ideal travel duration in hours for the full distance using the
	// indicated vehicle type. The field "idealDesc" should be used to describe the
	// ideal travel route.
	IdealDuration param.Opt[float64] `json:"idealDuration,omitzero"`
	// Unique identifier of the Site at the route's end location. This ID can be used
	// to obtain additional information on a Site using the 'get by ID' operation (e.g.
	// /udl/site/{id}). For example, the Site object with idSite = abc would be queried
	// as /udl/site/abc.
	IDSiteEnd param.Opt[string] `json:"idSiteEnd,omitzero"`
	// Unique identifier of the Site at the route's starting location. This ID can be
	// used to obtain additional information on a Site using the 'get by ID' operation
	// (e.g. /udl/site/{id}). For example, the Site object with idSite = abc would be
	// queried as /udl/site/abc.
	IDSiteStart param.Opt[string] `json:"idSiteStart,omitzero"`
	// Date of the last data point used in this calculation, in ISO8601 UTC format with
	// millisecond precision.
	LastPt param.Opt[time.Time] `json:"lastPt,omitzero" format:"date-time"`
	// Type of location used for route start and end points (e.g., ICAO, PORT, etc.).
	LocationType param.Opt[string] `json:"locationType,omitzero"`
	// Maximum travel duration for the indicated distance and type of vehicle in hours.
	MaxDuration param.Opt[float64] `json:"maxDuration,omitzero"`
	// Maximum speed during travel in the indicated unit of measurement, speedUnit.
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// Minimum travel duration for the indicated distance and type of vehicle in hours.
	MinDuration param.Opt[float64] `json:"minDuration,omitzero"`
	// Minimum speed during travel in the indicated unit of measurement, speedUnit.
	MinSpeed param.Opt[float64] `json:"minSpeed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Description of the portion of travel used to estimate the value of the
	// partialDuration field.
	PartialDesc param.Opt[string] `json:"partialDesc,omitzero"`
	// Estimated ideal travel duration in hours for a partial distance using the
	// indicated vehicle type. The field "partialDesc" should be used to specify the
	// intended portion of travel.
	PartialDuration param.Opt[float64] `json:"partialDuration,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The unit of measurement used for speed in this calculation.
	SpeedUnit param.Opt[string] `json:"speedUnit,omitzero"`
	// The time period this data was collected.
	TimePeriod param.Opt[string] `json:"timePeriod,omitzero"`
	// The vehicle category that is the subject of this calculation (e.g., AIRCRAFT,
	// CAR, BOAT, etc.).
	VehicleCategory param.Opt[string] `json:"vehicleCategory,omitzero"`
	// The vehicle type that is the subject of this calculation (e.g., C-17, F-15,
	// etc.).
	VehicleType param.Opt[string] `json:"vehicleType,omitzero"`
	paramObj
}

func (r RouteStatNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow RouteStatNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[RouteStatNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type RouteStatQueryParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RouteStatQueryParams]'s query parameters as `url.Values`.
func (r RouteStatQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RouteStatTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RouteStatTupleParams]'s query parameters as `url.Values`.
func (r RouteStatTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RouteStatUnvalidatedPublishParams struct {
	Body []RouteStatUnvalidatedPublishParamsBody
	paramObj
}

func (r RouteStatUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// General statistics applying to navigation routes utilized by vessels, aircraft,
// ground vehicles, etc.
//
// The properties ClassificationMarking, DataMode, LocationEnd, LocationStart,
// Source are required.
type RouteStatUnvalidatedPublishParamsBody struct {
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
	// End location of the vehicle.
	LocationEnd string `json:"locationEnd,required"`
	// Starting location of the vehicle.
	LocationStart string `json:"locationStart,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Average travel duration for the indicated distance and type of vehicle in hours.
	AvgDuration param.Opt[float64] `json:"avgDuration,omitzero"`
	// Average speed during travel in the indicated unit of measurement, speedUnit.
	AvgSpeed param.Opt[float64] `json:"avgSpeed,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The number of data points used in this travel duration calculation.
	DataPtsUsed param.Opt[int64] `json:"dataPtsUsed,omitzero"`
	// Distance between the departure and arrival locations in the indicated unit of
	// measurement, distUnit.
	Distance param.Opt[float64] `json:"distance,omitzero"`
	// The unit of measurement used for distance in this calculation.
	DistUnit param.Opt[string] `json:"distUnit,omitzero"`
	// Date of the first data point used in this calculation, in ISO8601 UTC format
	// with millisecond precision.
	FirstPt param.Opt[time.Time] `json:"firstPt,omitzero" format:"date-time"`
	// Description of the portion of travel used to estimate the value of the
	// idealDuration field.
	IdealDesc param.Opt[string] `json:"idealDesc,omitzero"`
	// Estimated ideal travel duration in hours for the full distance using the
	// indicated vehicle type. The field "idealDesc" should be used to describe the
	// ideal travel route.
	IdealDuration param.Opt[float64] `json:"idealDuration,omitzero"`
	// Unique identifier of the Site at the route's end location. This ID can be used
	// to obtain additional information on a Site using the 'get by ID' operation (e.g.
	// /udl/site/{id}). For example, the Site object with idSite = abc would be queried
	// as /udl/site/abc.
	IDSiteEnd param.Opt[string] `json:"idSiteEnd,omitzero"`
	// Unique identifier of the Site at the route's starting location. This ID can be
	// used to obtain additional information on a Site using the 'get by ID' operation
	// (e.g. /udl/site/{id}). For example, the Site object with idSite = abc would be
	// queried as /udl/site/abc.
	IDSiteStart param.Opt[string] `json:"idSiteStart,omitzero"`
	// Date of the last data point used in this calculation, in ISO8601 UTC format with
	// millisecond precision.
	LastPt param.Opt[time.Time] `json:"lastPt,omitzero" format:"date-time"`
	// Type of location used for route start and end points (e.g., ICAO, PORT, etc.).
	LocationType param.Opt[string] `json:"locationType,omitzero"`
	// Maximum travel duration for the indicated distance and type of vehicle in hours.
	MaxDuration param.Opt[float64] `json:"maxDuration,omitzero"`
	// Maximum speed during travel in the indicated unit of measurement, speedUnit.
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// Minimum travel duration for the indicated distance and type of vehicle in hours.
	MinDuration param.Opt[float64] `json:"minDuration,omitzero"`
	// Minimum speed during travel in the indicated unit of measurement, speedUnit.
	MinSpeed param.Opt[float64] `json:"minSpeed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Description of the portion of travel used to estimate the value of the
	// partialDuration field.
	PartialDesc param.Opt[string] `json:"partialDesc,omitzero"`
	// Estimated ideal travel duration in hours for a partial distance using the
	// indicated vehicle type. The field "partialDesc" should be used to specify the
	// intended portion of travel.
	PartialDuration param.Opt[float64] `json:"partialDuration,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The unit of measurement used for speed in this calculation.
	SpeedUnit param.Opt[string] `json:"speedUnit,omitzero"`
	// The time period this data was collected.
	TimePeriod param.Opt[string] `json:"timePeriod,omitzero"`
	// The vehicle category that is the subject of this calculation (e.g., AIRCRAFT,
	// CAR, BOAT, etc.).
	VehicleCategory param.Opt[string] `json:"vehicleCategory,omitzero"`
	// The vehicle type that is the subject of this calculation (e.g., C-17, F-15,
	// etc.).
	VehicleType param.Opt[string] `json:"vehicleType,omitzero"`
	paramObj
}

func (r RouteStatUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow RouteStatUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[RouteStatUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
