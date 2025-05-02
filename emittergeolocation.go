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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// EmittergeolocationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmittergeolocationService] method instead.
type EmittergeolocationService struct {
	Options []option.RequestOption
}

// NewEmittergeolocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmittergeolocationService(opts ...option.RequestOption) (r EmittergeolocationService) {
	r = EmittergeolocationService{}
	r.Options = opts
	return
}

// Service operation to take a single RF geolocation as a POST body and ingest into
// the database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *EmittergeolocationService) New(ctx context.Context, body EmittergeolocationNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/emittergeolocation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single RF geolocation by its unique ID passed as a
// path parameter.
func (r *EmittergeolocationService) Get(ctx context.Context, id string, query EmittergeolocationGetParams, opts ...option.RequestOption) (res *EmittergeolocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/emittergeolocation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to delete a RF geolocation specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance. Note, delete operations do not remove data
// from historical or publish/subscribe stores.
func (r *EmittergeolocationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/emittergeolocation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EmittergeolocationService) Count(ctx context.Context, query EmittergeolocationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/emittergeolocation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of RF
// geolocations as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *EmittergeolocationService) NewBulk(ctx context.Context, body EmittergeolocationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/emittergeolocation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EmittergeolocationService) Query(ctx context.Context, query EmittergeolocationQueryParams, opts ...option.RequestOption) (res *[]EmittergeolocationQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/emittergeolocation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EmittergeolocationService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/emittergeolocation/queryhelp"
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
func (r *EmittergeolocationService) Tuple(ctx context.Context, query EmittergeolocationTupleParams, opts ...option.RequestOption) (res *[]EmittergeolocationTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/emittergeolocation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple emittergeolocation records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *EmittergeolocationService) UnvalidatedPublish(ctx context.Context, body EmittergeolocationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-emittergeolocation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Model representation of Emitter geolocation data for a signal of interest.
type EmittergeolocationGetResponse struct {
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
	DataMode EmittergeolocationGetResponseDataMode `json:"dataMode,required"`
	// Type of the signal of interest of this Emitter Geo Location (e.g. RF).
	SignalOfInterestType string `json:"signalOfInterestType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// The EmitterGeo algorithm type and version used in Emitter geolocation
	// calculations.
	AlgVersion string `json:"algVersion"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the Point of Interest as projected on the ground.
	Area string `json:"area"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground.
	Atype string `json:"atype"`
	// The detected signal frequency in megahertz.
	CenterFreq float64 `json:"centerFreq"`
	// The name(s) of the subset of constellation spacecraft that made this detection.
	Cluster string `json:"cluster"`
	// The area of the confidence ellipse specified in meters squared to contain the
	// emitter with a 95% probability.
	ConfArea float64 `json:"confArea"`
	// The name of the satellite constellation.
	Constellation string `json:"constellation"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Specifies the creation time associated with the order in ISO 8601 UTC with
	// microsecond precision.
	CreatedTs time.Time `json:"createdTs" format:"date-time"`
	// The altitude relative to WGS-84 ellipsoid, in meters.
	DetectAlt float64 `json:"detectAlt"`
	// WGS-84 latitude of the most likely emitter location coordinate point, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	DetectLat float64 `json:"detectLat"`
	// WGS-84 longitude of the most likely emitter location coordinate point, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	DetectLon float64 `json:"detectLon"`
	// The end time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Confidence ellipsoid about the detection location [semi-major axis (m),
	// semi-minor axis (m), orientation (deg)].
	ErrEllp []float64 `json:"errEllp"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// Unique identifier of the satellite used to identify and geolocate Emitter
	// signals of interest. This ID can be used to obtain additional information on an
	// OnOrbit object using the 'get by ID' operation (e.g. /udl/onorbit/{id}). For
	// example, the onorbit object with idOnOrbit = abc would be queried as
	// /udl/onorbit/abc. Used when Emitter geolocation is done by a single satellite.
	IDOnOrbit string `json:"idOnOrbit"`
	// Optional identifier of the geolocated signal of interest RF Emitter for this
	// observation. This ID can be used to obtain additional information on an RF
	// Emitter object using the 'get by ID' operation (e.g. /udl/rfemitter/{id}). For
	// example, the rfemitter object with idRFEmitter = abc would be queried as
	// /udl/rfemitter/abc.
	IDRfEmitter string `json:"idRFEmitter"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc. Used when Emitter geolocation is done by a single sensor.
	IDSensor string `json:"idSensor"`
	// The maximum detected frequency in megahertz.
	MaxFreq float64 `json:"maxFreq"`
	// The minimum detected frequency in megahertz.
	MinFreq float64 `json:"minFreq"`
	// The count of single-burst observations used for this geolocation observation.
	NumBursts int64 `json:"numBursts"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// The order identifier for this Emitter Geo Location data set.
	OrderID string `json:"orderId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier of the satellite used to identify and geolocate Emitter
	// signals of interest of this observation. This may be an internal identifier and
	// not necessarily a valid satellite number. Used when Emitter geolocation is done
	// by a single satellite.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier of the emitter of interest for this observation. This may be
	// an internal identifier and not necessarily a valid emitter Id.
	OrigRfEmitterID string `json:"origRFEmitterId"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this Emitter Geo Location. This may be an internal
	// identifier and not necessarily a valid sensor ID. Used when Emitter geolocation
	// is done by a single sensor.
	OrigSensorID string `json:"origSensorId"`
	// Optional external identifier referencing the entity used in the calculation of
	// the geolocation.
	PassGroupID string `json:"passGroupId"`
	// The time representing the mean of the constituent single-burst observations in
	// ISO 8601 UTC with microsecond precision.
	ReceivedTs time.Time `json:"receivedTs" format:"date-time"`
	// Satellite/catalog number of the on-orbit spacecraft used to identify and
	// geolocate Emitter signals of interest of this detection. Used when Emitter
	// geolocation is done by a single satellite.
	SatNo int64 `json:"satNo"`
	// The name of the signal of interest.
	SignalOfInterest string `json:"signalOfInterest"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		SignalOfInterestType  resp.Field
		Source                resp.Field
		StartTime             resp.Field
		ID                    resp.Field
		Agjson                resp.Field
		AlgVersion            resp.Field
		Andims                resp.Field
		Area                  resp.Field
		Asrid                 resp.Field
		Atext                 resp.Field
		Atype                 resp.Field
		CenterFreq            resp.Field
		Cluster               resp.Field
		ConfArea              resp.Field
		Constellation         resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		CreatedTs             resp.Field
		DetectAlt             resp.Field
		DetectLat             resp.Field
		DetectLon             resp.Field
		EndTime               resp.Field
		ErrEllp               resp.Field
		ExternalID            resp.Field
		IDOnOrbit             resp.Field
		IDRfEmitter           resp.Field
		IDSensor              resp.Field
		MaxFreq               resp.Field
		MinFreq               resp.Field
		NumBursts             resp.Field
		OnOrbit               resp.Field
		OrderID               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		OrigRfEmitterID       resp.Field
		OrigSensorID          resp.Field
		PassGroupID           resp.Field
		ReceivedTs            resp.Field
		SatNo                 resp.Field
		SignalOfInterest      resp.Field
		SourceDl              resp.Field
		Tags                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmittergeolocationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmittergeolocationGetResponse) UnmarshalJSON(data []byte) error {
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
type EmittergeolocationGetResponseDataMode string

const (
	EmittergeolocationGetResponseDataModeReal      EmittergeolocationGetResponseDataMode = "REAL"
	EmittergeolocationGetResponseDataModeTest      EmittergeolocationGetResponseDataMode = "TEST"
	EmittergeolocationGetResponseDataModeSimulated EmittergeolocationGetResponseDataMode = "SIMULATED"
	EmittergeolocationGetResponseDataModeExercise  EmittergeolocationGetResponseDataMode = "EXERCISE"
)

// Model representation of Emitter geolocation data for a signal of interest.
type EmittergeolocationQueryResponse struct {
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
	DataMode EmittergeolocationQueryResponseDataMode `json:"dataMode,required"`
	// Type of the signal of interest of this Emitter Geo Location (e.g. RF).
	SignalOfInterestType string `json:"signalOfInterestType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// The EmitterGeo algorithm type and version used in Emitter geolocation
	// calculations.
	AlgVersion string `json:"algVersion"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground.
	Atype string `json:"atype"`
	// The detected signal frequency in megahertz.
	CenterFreq float64 `json:"centerFreq"`
	// The name(s) of the subset of constellation spacecraft that made this detection.
	Cluster string `json:"cluster"`
	// The area of the confidence ellipse specified in meters squared to contain the
	// emitter with a 95% probability.
	ConfArea float64 `json:"confArea"`
	// The name of the satellite constellation.
	Constellation string `json:"constellation"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Specifies the creation time associated with the order in ISO 8601 UTC with
	// microsecond precision.
	CreatedTs time.Time `json:"createdTs" format:"date-time"`
	// The altitude relative to WGS-84 ellipsoid, in meters.
	DetectAlt float64 `json:"detectAlt"`
	// WGS-84 latitude of the most likely emitter location coordinate point, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	DetectLat float64 `json:"detectLat"`
	// WGS-84 longitude of the most likely emitter location coordinate point, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	DetectLon float64 `json:"detectLon"`
	// The end time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Confidence ellipsoid about the detection location [semi-major axis (m),
	// semi-minor axis (m), orientation (deg)].
	ErrEllp []float64 `json:"errEllp"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// Unique identifier of the satellite used to identify and geolocate Emitter
	// signals of interest. This ID can be used to obtain additional information on an
	// OnOrbit object using the 'get by ID' operation (e.g. /udl/onorbit/{id}). For
	// example, the onorbit object with idOnOrbit = abc would be queried as
	// /udl/onorbit/abc. Used when Emitter geolocation is done by a single satellite.
	IDOnOrbit string `json:"idOnOrbit"`
	// Optional identifier of the geolocated signal of interest RF Emitter for this
	// observation. This ID can be used to obtain additional information on an RF
	// Emitter object using the 'get by ID' operation (e.g. /udl/rfemitter/{id}). For
	// example, the rfemitter object with idRFEmitter = abc would be queried as
	// /udl/rfemitter/abc.
	IDRfEmitter string `json:"idRFEmitter"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc. Used when Emitter geolocation is done by a single sensor.
	IDSensor string `json:"idSensor"`
	// The maximum detected frequency in megahertz.
	MaxFreq float64 `json:"maxFreq"`
	// The minimum detected frequency in megahertz.
	MinFreq float64 `json:"minFreq"`
	// The count of single-burst observations used for this geolocation observation.
	NumBursts int64 `json:"numBursts"`
	// The order identifier for this Emitter Geo Location data set.
	OrderID string `json:"orderId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier of the satellite used to identify and geolocate Emitter
	// signals of interest of this observation. This may be an internal identifier and
	// not necessarily a valid satellite number. Used when Emitter geolocation is done
	// by a single satellite.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier of the emitter of interest for this observation. This may be
	// an internal identifier and not necessarily a valid emitter Id.
	OrigRfEmitterID string `json:"origRFEmitterId"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this Emitter Geo Location. This may be an internal
	// identifier and not necessarily a valid sensor ID. Used when Emitter geolocation
	// is done by a single sensor.
	OrigSensorID string `json:"origSensorId"`
	// Optional external identifier referencing the entity used in the calculation of
	// the geolocation.
	PassGroupID string `json:"passGroupId"`
	// The time representing the mean of the constituent single-burst observations in
	// ISO 8601 UTC with microsecond precision.
	ReceivedTs time.Time `json:"receivedTs" format:"date-time"`
	// Satellite/catalog number of the on-orbit spacecraft used to identify and
	// geolocate Emitter signals of interest of this detection. Used when Emitter
	// geolocation is done by a single satellite.
	SatNo int64 `json:"satNo"`
	// The name of the signal of interest.
	SignalOfInterest string `json:"signalOfInterest"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		SignalOfInterestType  resp.Field
		Source                resp.Field
		StartTime             resp.Field
		ID                    resp.Field
		Agjson                resp.Field
		AlgVersion            resp.Field
		Andims                resp.Field
		Asrid                 resp.Field
		Atext                 resp.Field
		Atype                 resp.Field
		CenterFreq            resp.Field
		Cluster               resp.Field
		ConfArea              resp.Field
		Constellation         resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		CreatedTs             resp.Field
		DetectAlt             resp.Field
		DetectLat             resp.Field
		DetectLon             resp.Field
		EndTime               resp.Field
		ErrEllp               resp.Field
		ExternalID            resp.Field
		IDOnOrbit             resp.Field
		IDRfEmitter           resp.Field
		IDSensor              resp.Field
		MaxFreq               resp.Field
		MinFreq               resp.Field
		NumBursts             resp.Field
		OrderID               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		OrigRfEmitterID       resp.Field
		OrigSensorID          resp.Field
		PassGroupID           resp.Field
		ReceivedTs            resp.Field
		SatNo                 resp.Field
		SignalOfInterest      resp.Field
		SourceDl              resp.Field
		Tags                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmittergeolocationQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *EmittergeolocationQueryResponse) UnmarshalJSON(data []byte) error {
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
type EmittergeolocationQueryResponseDataMode string

const (
	EmittergeolocationQueryResponseDataModeReal      EmittergeolocationQueryResponseDataMode = "REAL"
	EmittergeolocationQueryResponseDataModeTest      EmittergeolocationQueryResponseDataMode = "TEST"
	EmittergeolocationQueryResponseDataModeSimulated EmittergeolocationQueryResponseDataMode = "SIMULATED"
	EmittergeolocationQueryResponseDataModeExercise  EmittergeolocationQueryResponseDataMode = "EXERCISE"
)

// Model representation of Emitter geolocation data for a signal of interest.
type EmittergeolocationTupleResponse struct {
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
	DataMode EmittergeolocationTupleResponseDataMode `json:"dataMode,required"`
	// Type of the signal of interest of this Emitter Geo Location (e.g. RF).
	SignalOfInterestType string `json:"signalOfInterestType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// The EmitterGeo algorithm type and version used in Emitter geolocation
	// calculations.
	AlgVersion string `json:"algVersion"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the Point of Interest as projected on the ground.
	Area string `json:"area"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground.
	Atype string `json:"atype"`
	// The detected signal frequency in megahertz.
	CenterFreq float64 `json:"centerFreq"`
	// The name(s) of the subset of constellation spacecraft that made this detection.
	Cluster string `json:"cluster"`
	// The area of the confidence ellipse specified in meters squared to contain the
	// emitter with a 95% probability.
	ConfArea float64 `json:"confArea"`
	// The name of the satellite constellation.
	Constellation string `json:"constellation"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Specifies the creation time associated with the order in ISO 8601 UTC with
	// microsecond precision.
	CreatedTs time.Time `json:"createdTs" format:"date-time"`
	// The altitude relative to WGS-84 ellipsoid, in meters.
	DetectAlt float64 `json:"detectAlt"`
	// WGS-84 latitude of the most likely emitter location coordinate point, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	DetectLat float64 `json:"detectLat"`
	// WGS-84 longitude of the most likely emitter location coordinate point, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	DetectLon float64 `json:"detectLon"`
	// The end time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Confidence ellipsoid about the detection location [semi-major axis (m),
	// semi-minor axis (m), orientation (deg)].
	ErrEllp []float64 `json:"errEllp"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// Unique identifier of the satellite used to identify and geolocate Emitter
	// signals of interest. This ID can be used to obtain additional information on an
	// OnOrbit object using the 'get by ID' operation (e.g. /udl/onorbit/{id}). For
	// example, the onorbit object with idOnOrbit = abc would be queried as
	// /udl/onorbit/abc. Used when Emitter geolocation is done by a single satellite.
	IDOnOrbit string `json:"idOnOrbit"`
	// Optional identifier of the geolocated signal of interest RF Emitter for this
	// observation. This ID can be used to obtain additional information on an RF
	// Emitter object using the 'get by ID' operation (e.g. /udl/rfemitter/{id}). For
	// example, the rfemitter object with idRFEmitter = abc would be queried as
	// /udl/rfemitter/abc.
	IDRfEmitter string `json:"idRFEmitter"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc. Used when Emitter geolocation is done by a single sensor.
	IDSensor string `json:"idSensor"`
	// The maximum detected frequency in megahertz.
	MaxFreq float64 `json:"maxFreq"`
	// The minimum detected frequency in megahertz.
	MinFreq float64 `json:"minFreq"`
	// The count of single-burst observations used for this geolocation observation.
	NumBursts int64 `json:"numBursts"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// The order identifier for this Emitter Geo Location data set.
	OrderID string `json:"orderId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier of the satellite used to identify and geolocate Emitter
	// signals of interest of this observation. This may be an internal identifier and
	// not necessarily a valid satellite number. Used when Emitter geolocation is done
	// by a single satellite.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier of the emitter of interest for this observation. This may be
	// an internal identifier and not necessarily a valid emitter Id.
	OrigRfEmitterID string `json:"origRFEmitterId"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this Emitter Geo Location. This may be an internal
	// identifier and not necessarily a valid sensor ID. Used when Emitter geolocation
	// is done by a single sensor.
	OrigSensorID string `json:"origSensorId"`
	// Optional external identifier referencing the entity used in the calculation of
	// the geolocation.
	PassGroupID string `json:"passGroupId"`
	// The time representing the mean of the constituent single-burst observations in
	// ISO 8601 UTC with microsecond precision.
	ReceivedTs time.Time `json:"receivedTs" format:"date-time"`
	// Satellite/catalog number of the on-orbit spacecraft used to identify and
	// geolocate Emitter signals of interest of this detection. Used when Emitter
	// geolocation is done by a single satellite.
	SatNo int64 `json:"satNo"`
	// The name of the signal of interest.
	SignalOfInterest string `json:"signalOfInterest"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		SignalOfInterestType  resp.Field
		Source                resp.Field
		StartTime             resp.Field
		ID                    resp.Field
		Agjson                resp.Field
		AlgVersion            resp.Field
		Andims                resp.Field
		Area                  resp.Field
		Asrid                 resp.Field
		Atext                 resp.Field
		Atype                 resp.Field
		CenterFreq            resp.Field
		Cluster               resp.Field
		ConfArea              resp.Field
		Constellation         resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		CreatedTs             resp.Field
		DetectAlt             resp.Field
		DetectLat             resp.Field
		DetectLon             resp.Field
		EndTime               resp.Field
		ErrEllp               resp.Field
		ExternalID            resp.Field
		IDOnOrbit             resp.Field
		IDRfEmitter           resp.Field
		IDSensor              resp.Field
		MaxFreq               resp.Field
		MinFreq               resp.Field
		NumBursts             resp.Field
		OnOrbit               resp.Field
		OrderID               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		OrigRfEmitterID       resp.Field
		OrigSensorID          resp.Field
		PassGroupID           resp.Field
		ReceivedTs            resp.Field
		SatNo                 resp.Field
		SignalOfInterest      resp.Field
		SourceDl              resp.Field
		Tags                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmittergeolocationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *EmittergeolocationTupleResponse) UnmarshalJSON(data []byte) error {
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
type EmittergeolocationTupleResponseDataMode string

const (
	EmittergeolocationTupleResponseDataModeReal      EmittergeolocationTupleResponseDataMode = "REAL"
	EmittergeolocationTupleResponseDataModeTest      EmittergeolocationTupleResponseDataMode = "TEST"
	EmittergeolocationTupleResponseDataModeSimulated EmittergeolocationTupleResponseDataMode = "SIMULATED"
	EmittergeolocationTupleResponseDataModeExercise  EmittergeolocationTupleResponseDataMode = "EXERCISE"
)

type EmittergeolocationNewParams struct {
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
	DataMode EmittergeolocationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Type of the signal of interest of this Emitter Geo Location (e.g. RF).
	SignalOfInterestType string `json:"signalOfInterestType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// The EmitterGeo algorithm type and version used in Emitter geolocation
	// calculations.
	AlgVersion param.Opt[string] `json:"algVersion,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the Point of Interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// The detected signal frequency in megahertz.
	CenterFreq param.Opt[float64] `json:"centerFreq,omitzero"`
	// The name(s) of the subset of constellation spacecraft that made this detection.
	Cluster param.Opt[string] `json:"cluster,omitzero"`
	// The area of the confidence ellipse specified in meters squared to contain the
	// emitter with a 95% probability.
	ConfArea param.Opt[float64] `json:"confArea,omitzero"`
	// The name of the satellite constellation.
	Constellation param.Opt[string] `json:"constellation,omitzero"`
	// Specifies the creation time associated with the order in ISO 8601 UTC with
	// microsecond precision.
	CreatedTs param.Opt[time.Time] `json:"createdTs,omitzero" format:"date-time"`
	// The altitude relative to WGS-84 ellipsoid, in meters.
	DetectAlt param.Opt[float64] `json:"detectAlt,omitzero"`
	// WGS-84 latitude of the most likely emitter location coordinate point, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	DetectLat param.Opt[float64] `json:"detectLat,omitzero"`
	// WGS-84 longitude of the most likely emitter location coordinate point, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	DetectLon param.Opt[float64] `json:"detectLon,omitzero"`
	// The end time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Optional identifier of the geolocated signal of interest RF Emitter for this
	// observation. This ID can be used to obtain additional information on an RF
	// Emitter object using the 'get by ID' operation (e.g. /udl/rfemitter/{id}). For
	// example, the rfemitter object with idRFEmitter = abc would be queried as
	// /udl/rfemitter/abc.
	IDRfEmitter param.Opt[string] `json:"idRFEmitter,omitzero"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc. Used when Emitter geolocation is done by a single sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The maximum detected frequency in megahertz.
	MaxFreq param.Opt[float64] `json:"maxFreq,omitzero"`
	// The minimum detected frequency in megahertz.
	MinFreq param.Opt[float64] `json:"minFreq,omitzero"`
	// The count of single-burst observations used for this geolocation observation.
	NumBursts param.Opt[int64] `json:"numBursts,omitzero"`
	// The order identifier for this Emitter Geo Location data set.
	OrderID param.Opt[string] `json:"orderId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier of the satellite used to identify and geolocate Emitter
	// signals of interest of this observation. This may be an internal identifier and
	// not necessarily a valid satellite number. Used when Emitter geolocation is done
	// by a single satellite.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier of the emitter of interest for this observation. This may be
	// an internal identifier and not necessarily a valid emitter Id.
	OrigRfEmitterID param.Opt[string] `json:"origRFEmitterId,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this Emitter Geo Location. This may be an internal
	// identifier and not necessarily a valid sensor ID. Used when Emitter geolocation
	// is done by a single sensor.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Optional external identifier referencing the entity used in the calculation of
	// the geolocation.
	PassGroupID param.Opt[string] `json:"passGroupId,omitzero"`
	// The time representing the mean of the constituent single-burst observations in
	// ISO 8601 UTC with microsecond precision.
	ReceivedTs param.Opt[time.Time] `json:"receivedTs,omitzero" format:"date-time"`
	// Satellite/catalog number of the on-orbit spacecraft used to identify and
	// geolocate Emitter signals of interest of this detection. Used when Emitter
	// geolocation is done by a single satellite.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The name of the signal of interest.
	SignalOfInterest param.Opt[string] `json:"signalOfInterest,omitzero"`
	// Confidence ellipsoid about the detection location [semi-major axis (m),
	// semi-minor axis (m), orientation (deg)].
	ErrEllp []float64 `json:"errEllp,omitzero"`
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
func (f EmittergeolocationNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r EmittergeolocationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmittergeolocationNewParams
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
type EmittergeolocationNewParamsDataMode string

const (
	EmittergeolocationNewParamsDataModeReal      EmittergeolocationNewParamsDataMode = "REAL"
	EmittergeolocationNewParamsDataModeTest      EmittergeolocationNewParamsDataMode = "TEST"
	EmittergeolocationNewParamsDataModeSimulated EmittergeolocationNewParamsDataMode = "SIMULATED"
	EmittergeolocationNewParamsDataModeExercise  EmittergeolocationNewParamsDataMode = "EXERCISE"
)

type EmittergeolocationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EmittergeolocationGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [EmittergeolocationGetParams]'s query parameters as
// `url.Values`.
func (r EmittergeolocationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmittergeolocationCountParams struct {
	// The start time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EmittergeolocationCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [EmittergeolocationCountParams]'s query parameters as
// `url.Values`.
func (r EmittergeolocationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmittergeolocationNewBulkParams struct {
	Body []EmittergeolocationNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EmittergeolocationNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r EmittergeolocationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Model representation of Emitter geolocation data for a signal of interest.
//
// The properties ClassificationMarking, DataMode, SignalOfInterestType, Source,
// StartTime are required.
type EmittergeolocationNewBulkParamsBody struct {
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
	// Type of the signal of interest of this Emitter Geo Location (e.g. RF).
	SignalOfInterestType string `json:"signalOfInterestType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// The EmitterGeo algorithm type and version used in Emitter geolocation
	// calculations.
	AlgVersion param.Opt[string] `json:"algVersion,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the Point of Interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// The detected signal frequency in megahertz.
	CenterFreq param.Opt[float64] `json:"centerFreq,omitzero"`
	// The name(s) of the subset of constellation spacecraft that made this detection.
	Cluster param.Opt[string] `json:"cluster,omitzero"`
	// The area of the confidence ellipse specified in meters squared to contain the
	// emitter with a 95% probability.
	ConfArea param.Opt[float64] `json:"confArea,omitzero"`
	// The name of the satellite constellation.
	Constellation param.Opt[string] `json:"constellation,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Specifies the creation time associated with the order in ISO 8601 UTC with
	// microsecond precision.
	CreatedTs param.Opt[time.Time] `json:"createdTs,omitzero" format:"date-time"`
	// The altitude relative to WGS-84 ellipsoid, in meters.
	DetectAlt param.Opt[float64] `json:"detectAlt,omitzero"`
	// WGS-84 latitude of the most likely emitter location coordinate point, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	DetectLat param.Opt[float64] `json:"detectLat,omitzero"`
	// WGS-84 longitude of the most likely emitter location coordinate point, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	DetectLon param.Opt[float64] `json:"detectLon,omitzero"`
	// The end time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Unique identifier of the satellite used to identify and geolocate Emitter
	// signals of interest. This ID can be used to obtain additional information on an
	// OnOrbit object using the 'get by ID' operation (e.g. /udl/onorbit/{id}). For
	// example, the onorbit object with idOnOrbit = abc would be queried as
	// /udl/onorbit/abc. Used when Emitter geolocation is done by a single satellite.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Optional identifier of the geolocated signal of interest RF Emitter for this
	// observation. This ID can be used to obtain additional information on an RF
	// Emitter object using the 'get by ID' operation (e.g. /udl/rfemitter/{id}). For
	// example, the rfemitter object with idRFEmitter = abc would be queried as
	// /udl/rfemitter/abc.
	IDRfEmitter param.Opt[string] `json:"idRFEmitter,omitzero"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc. Used when Emitter geolocation is done by a single sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The maximum detected frequency in megahertz.
	MaxFreq param.Opt[float64] `json:"maxFreq,omitzero"`
	// The minimum detected frequency in megahertz.
	MinFreq param.Opt[float64] `json:"minFreq,omitzero"`
	// The count of single-burst observations used for this geolocation observation.
	NumBursts param.Opt[int64] `json:"numBursts,omitzero"`
	// The order identifier for this Emitter Geo Location data set.
	OrderID param.Opt[string] `json:"orderId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier of the satellite used to identify and geolocate Emitter
	// signals of interest of this observation. This may be an internal identifier and
	// not necessarily a valid satellite number. Used when Emitter geolocation is done
	// by a single satellite.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier of the emitter of interest for this observation. This may be
	// an internal identifier and not necessarily a valid emitter Id.
	OrigRfEmitterID param.Opt[string] `json:"origRFEmitterId,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this Emitter Geo Location. This may be an internal
	// identifier and not necessarily a valid sensor ID. Used when Emitter geolocation
	// is done by a single sensor.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Optional external identifier referencing the entity used in the calculation of
	// the geolocation.
	PassGroupID param.Opt[string] `json:"passGroupId,omitzero"`
	// The time representing the mean of the constituent single-burst observations in
	// ISO 8601 UTC with microsecond precision.
	ReceivedTs param.Opt[time.Time] `json:"receivedTs,omitzero" format:"date-time"`
	// Satellite/catalog number of the on-orbit spacecraft used to identify and
	// geolocate Emitter signals of interest of this detection. Used when Emitter
	// geolocation is done by a single satellite.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The name of the signal of interest.
	SignalOfInterest param.Opt[string] `json:"signalOfInterest,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Confidence ellipsoid about the detection location [semi-major axis (m),
	// semi-minor axis (m), orientation (deg)].
	ErrEllp []float64 `json:"errEllp,omitzero"`
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
func (f EmittergeolocationNewBulkParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EmittergeolocationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EmittergeolocationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EmittergeolocationNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type EmittergeolocationQueryParams struct {
	// The start time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EmittergeolocationQueryParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [EmittergeolocationQueryParams]'s query parameters as
// `url.Values`.
func (r EmittergeolocationQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmittergeolocationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The start time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EmittergeolocationTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [EmittergeolocationTupleParams]'s query parameters as
// `url.Values`.
func (r EmittergeolocationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmittergeolocationUnvalidatedPublishParams struct {
	Body []EmittergeolocationUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EmittergeolocationUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r EmittergeolocationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Model representation of Emitter geolocation data for a signal of interest.
//
// The properties ClassificationMarking, DataMode, SignalOfInterestType, Source,
// StartTime are required.
type EmittergeolocationUnvalidatedPublishParamsBody struct {
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
	// Type of the signal of interest of this Emitter Geo Location (e.g. RF).
	SignalOfInterestType string `json:"signalOfInterestType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// The EmitterGeo algorithm type and version used in Emitter geolocation
	// calculations.
	AlgVersion param.Opt[string] `json:"algVersion,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the Point of Interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// The detected signal frequency in megahertz.
	CenterFreq param.Opt[float64] `json:"centerFreq,omitzero"`
	// The name(s) of the subset of constellation spacecraft that made this detection.
	Cluster param.Opt[string] `json:"cluster,omitzero"`
	// The area of the confidence ellipse specified in meters squared to contain the
	// emitter with a 95% probability.
	ConfArea param.Opt[float64] `json:"confArea,omitzero"`
	// The name of the satellite constellation.
	Constellation param.Opt[string] `json:"constellation,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Specifies the creation time associated with the order in ISO 8601 UTC with
	// microsecond precision.
	CreatedTs param.Opt[time.Time] `json:"createdTs,omitzero" format:"date-time"`
	// The altitude relative to WGS-84 ellipsoid, in meters.
	DetectAlt param.Opt[float64] `json:"detectAlt,omitzero"`
	// WGS-84 latitude of the most likely emitter location coordinate point, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	DetectLat param.Opt[float64] `json:"detectLat,omitzero"`
	// WGS-84 longitude of the most likely emitter location coordinate point, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	DetectLon param.Opt[float64] `json:"detectLon,omitzero"`
	// The end time for this Emitter Geo Location data set in ISO 8601 UTC with
	// microsecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Unique identifier of the satellite used to identify and geolocate Emitter
	// signals of interest. This ID can be used to obtain additional information on an
	// OnOrbit object using the 'get by ID' operation (e.g. /udl/onorbit/{id}). For
	// example, the onorbit object with idOnOrbit = abc would be queried as
	// /udl/onorbit/abc. Used when Emitter geolocation is done by a single satellite.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Optional identifier of the geolocated signal of interest RF Emitter for this
	// observation. This ID can be used to obtain additional information on an RF
	// Emitter object using the 'get by ID' operation (e.g. /udl/rfemitter/{id}). For
	// example, the rfemitter object with idRFEmitter = abc would be queried as
	// /udl/rfemitter/abc.
	IDRfEmitter param.Opt[string] `json:"idRFEmitter,omitzero"`
	// Unique identifier of the reporting sensor. This ID can be used to obtain
	// additional information on a sensor using the 'get by ID' operation (e.g.
	// /udl/sensor/{id}). For example, the sensor with idSensor = abc would be queried
	// as /udl/sensor/abc. Used when Emitter geolocation is done by a single sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The maximum detected frequency in megahertz.
	MaxFreq param.Opt[float64] `json:"maxFreq,omitzero"`
	// The minimum detected frequency in megahertz.
	MinFreq param.Opt[float64] `json:"minFreq,omitzero"`
	// The count of single-burst observations used for this geolocation observation.
	NumBursts param.Opt[int64] `json:"numBursts,omitzero"`
	// The order identifier for this Emitter Geo Location data set.
	OrderID param.Opt[string] `json:"orderId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier of the satellite used to identify and geolocate Emitter
	// signals of interest of this observation. This may be an internal identifier and
	// not necessarily a valid satellite number. Used when Emitter geolocation is done
	// by a single satellite.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier of the emitter of interest for this observation. This may be
	// an internal identifier and not necessarily a valid emitter Id.
	OrigRfEmitterID param.Opt[string] `json:"origRFEmitterId,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this Emitter Geo Location. This may be an internal
	// identifier and not necessarily a valid sensor ID. Used when Emitter geolocation
	// is done by a single sensor.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Optional external identifier referencing the entity used in the calculation of
	// the geolocation.
	PassGroupID param.Opt[string] `json:"passGroupId,omitzero"`
	// The time representing the mean of the constituent single-burst observations in
	// ISO 8601 UTC with microsecond precision.
	ReceivedTs param.Opt[time.Time] `json:"receivedTs,omitzero" format:"date-time"`
	// Satellite/catalog number of the on-orbit spacecraft used to identify and
	// geolocate Emitter signals of interest of this detection. Used when Emitter
	// geolocation is done by a single satellite.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The name of the signal of interest.
	SignalOfInterest param.Opt[string] `json:"signalOfInterest,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Confidence ellipsoid about the detection location [semi-major axis (m),
	// semi-minor axis (m), orientation (deg)].
	ErrEllp []float64 `json:"errEllp,omitzero"`
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
func (f EmittergeolocationUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EmittergeolocationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EmittergeolocationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EmittergeolocationUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
