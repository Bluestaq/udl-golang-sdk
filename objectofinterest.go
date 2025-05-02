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
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// ObjectofinterestService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectofinterestService] method instead.
type ObjectofinterestService struct {
	Options []option.RequestOption
}

// NewObjectofinterestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectofinterestService(opts ...option.RequestOption) (r ObjectofinterestService) {
	r = ObjectofinterestService{}
	r.Options = opts
	return
}

// Service operation to take a single ObjectOfInterest as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *ObjectofinterestService) New(ctx context.Context, body ObjectofinterestNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/objectofinterest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single ObjectOfInterest. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *ObjectofinterestService) Update(ctx context.Context, id string, body ObjectofinterestUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/objectofinterest/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObjectofinterestService) List(ctx context.Context, query ObjectofinterestListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ObjectofinterestListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/objectofinterest"
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
func (r *ObjectofinterestService) ListAutoPaging(ctx context.Context, query ObjectofinterestListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ObjectofinterestListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a ObjectOfInterest object specified by the passed ID
// path parameter. A ObjectOfInterest is an on-orbit ObjectOfInterestunications
// payload, including supporting data such as transponders and channels, etc. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *ObjectofinterestService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/objectofinterest/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObjectofinterestService) Count(ctx context.Context, query ObjectofinterestCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/objectofinterest/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single ObjectOfInterest record by its unique ID
// passed as a path parameter.
func (r *ObjectofinterestService) Get(ctx context.Context, id string, query ObjectofinterestGetParams, opts ...option.RequestOption) (res *ObjectofinterestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/objectofinterest/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ObjectofinterestService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/objectofinterest/queryhelp"
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
func (r *ObjectofinterestService) Tuple(ctx context.Context, query ObjectofinterestTupleParams, opts ...option.RequestOption) (res *[]ObjectofinterestTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/objectofinterest/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// OnOrbit objects of interest, which include information about the last known
// state of the object.
type ObjectofinterestListResponse struct {
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
	DataMode ObjectofinterestListResponseDataMode `json:"dataMode,required"`
	// UUID of the parent Onorbit record.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Sensor tasking start time for object of interest.
	SensorTaskingStartTime time.Time `json:"sensorTaskingStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Time of last status change of the object of interest event.
	StatusDate time.Time `json:"statusDate,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Last reported apogee. The Orbit point furthest from the center of the earth in
	// kilometers.
	Apogee float64 `json:"apogee"`
	// Last reported argument of perigee. The argument of perigee is the angle in
	// degrees formed between the perigee and the ascending node. If the perigee would
	// occur at the ascending node, the argument of perigee would be 0.
	ArgOfPerigee float64 `json:"argOfPerigee"`
	// Last reported drag term for SGP4 orbital model, used for calculating decay
	// constants for altitude, eccentricity etc, measured in inverse earth radii.
	BStar float64 `json:"bStar"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Possible delta time applications for the object of interest, in seconds.
	DeltaTs []float64 `json:"deltaTs"`
	// Possible delta V applications for the object of interest, in km/sec.
	DeltaVs []float64 `json:"deltaVs"`
	// Description of the object of interest event.
	Description string `json:"description"`
	// Last reported eccentricity of the object. The orbital eccentricity of an
	// astronomical object is a parameter that determines the amount by which its orbit
	// around another body deviates from a perfect circle. A value of 0 is a circular
	// orbit, values between 0 and 1 form an elliptic orbit, 1 is a parabolic escape
	// orbit, and greater than 1 is a hyperbolic escape orbit.
	Eccentricity float64 `json:"eccentricity"`
	// Last reported elset epoch time in ISO 8601 UTC time, with microsecond precision.
	ElsetEpoch time.Time `json:"elsetEpoch" format:"date-time"`
	// Last reported inclination of the object. Inclination is the angle between the
	// equator and the orbit when looking from the center of the Earth. If the orbit
	// went exactly around the equator from left to right, then the inclination would
	// be 0. The inclination ranges from 0 to 180 degrees.
	Inclination float64 `json:"inclination"`
	// Last reported observation time in ISO 8601 UTC time, with microsecond precision.
	LastObTime time.Time `json:"lastObTime" format:"date-time"`
	// Last reported meanAnomaly. Mean anomoly is where the satellite is in its orbital
	// path. The mean anomaly ranges from 0 to 360 degrees. The mean anomaly is
	// referenced to the perigee. If the satellite were at the perigee, the mean
	// anomaly would be 0.
	MeanAnomaly float64 `json:"meanAnomaly"`
	// Last reported mean motion of the object. Mean motion is the angular speed
	// required for a body to complete one orbit, assuming constant speed in a circular
	// orbit which completes in the same time as the variable speed, elliptical orbit
	// of the actual body. Measured in revolutions per day.
	MeanMotion float64 `json:"meanMotion"`
	// Last reported 2nd derivative of the mean motion with respect to time. Units are
	// revolutions per day cubed.
	MeanMotionDDot float64 `json:"meanMotionDDot"`
	// Last reported 1st derivative of the mean motion with respect to time. Units are
	// revolutions per day squared.
	MeanMotionDot float64 `json:"meanMotionDot"`
	// The time at which an attempted observation of the object of interest noticed it
	// was missing, in ISO 8601 UTC time, with microsecond precision.
	MissedObTime time.Time `json:"missedObTime" format:"date-time"`
	// Unique name of the object of interest event.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Last reported perigee. The orbit point nearest to the center of the earth in
	// kilometers.
	Perigee float64 `json:"perigee"`
	// Last reported orbit period. Period of the orbit is equal to inverse of mean
	// motion.
	Period float64 `json:"period"`
	// Priority of the object of interest as an integer (1=highest priority).
	Priority int64 `json:"priority"`
	// Last reported raan. Right ascension of the ascending node, or RAAN is the angle
	// as measured in degrees eastwards (or, as seen from the north, counterclockwise)
	// from the First Point of Aries to the ascending node, which is where the orbit
	// crosses the equator when traveling north.
	Raan float64 `json:"raan"`
	// The last reported revolution number. The value is incremented when a satellite
	// crosses the equator on an ascending pass.
	RevNo int64 `json:"revNo"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Last reported semi major axis, which is the sum of the periapsis and apoapsis
	// distances divided by two. For circular orbits, the semimajor axis is the
	// distance between the centers of the bodies, not the distance of the bodies from
	// the center of mass.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// Sensor tasking stop time for object of interest.
	SensorTaskingStopTime time.Time `json:"sensorTaskingStopTime" format:"date-time"`
	// Status of the object of interest event (e.g. OPEN, CLOSED, CANCELLED).
	Status string `json:"status"`
	// Last reported state vector epoch time in ISO 8601 UTC time, with microsecond
	// precision.
	SvEpoch time.Time `json:"svEpoch" format:"date-time"`
	// Last reported x position of the object in km, in J2000 coordinates.
	X float64 `json:"x"`
	// Last reported x velocity of the object in km/sec, in J2000 coordinates.
	Xvel float64 `json:"xvel"`
	// Last reported y position of the object in km, in J2000 coordinates.
	Y float64 `json:"y"`
	// Last reported y velocity of the object in km/sec, in J2000 coordinates.
	Yvel float64 `json:"yvel"`
	// Last reported z position of the object in km, in J2000 coordinates.
	Z float64 `json:"z"`
	// Last reported z velocity of the object in km/sec, in J2000 coordinates.
	Zvel float64 `json:"zvel"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking  resp.Field
		DataMode               resp.Field
		IDOnOrbit              resp.Field
		SensorTaskingStartTime resp.Field
		Source                 resp.Field
		StatusDate             resp.Field
		ID                     resp.Field
		Apogee                 resp.Field
		ArgOfPerigee           resp.Field
		BStar                  resp.Field
		CreatedAt              resp.Field
		CreatedBy              resp.Field
		DeltaTs                resp.Field
		DeltaVs                resp.Field
		Description            resp.Field
		Eccentricity           resp.Field
		ElsetEpoch             resp.Field
		Inclination            resp.Field
		LastObTime             resp.Field
		MeanAnomaly            resp.Field
		MeanMotion             resp.Field
		MeanMotionDDot         resp.Field
		MeanMotionDot          resp.Field
		MissedObTime           resp.Field
		Name                   resp.Field
		Origin                 resp.Field
		OrigNetwork            resp.Field
		Perigee                resp.Field
		Period                 resp.Field
		Priority               resp.Field
		Raan                   resp.Field
		RevNo                  resp.Field
		SatNo                  resp.Field
		SemiMajorAxis          resp.Field
		SensorTaskingStopTime  resp.Field
		Status                 resp.Field
		SvEpoch                resp.Field
		X                      resp.Field
		Xvel                   resp.Field
		Y                      resp.Field
		Yvel                   resp.Field
		Z                      resp.Field
		Zvel                   resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectofinterestListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectofinterestListResponse) UnmarshalJSON(data []byte) error {
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
type ObjectofinterestListResponseDataMode string

const (
	ObjectofinterestListResponseDataModeReal      ObjectofinterestListResponseDataMode = "REAL"
	ObjectofinterestListResponseDataModeTest      ObjectofinterestListResponseDataMode = "TEST"
	ObjectofinterestListResponseDataModeSimulated ObjectofinterestListResponseDataMode = "SIMULATED"
	ObjectofinterestListResponseDataModeExercise  ObjectofinterestListResponseDataMode = "EXERCISE"
)

// OnOrbit objects of interest, which include information about the last known
// state of the object.
type ObjectofinterestGetResponse struct {
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
	DataMode ObjectofinterestGetResponseDataMode `json:"dataMode,required"`
	// UUID of the parent Onorbit record.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Sensor tasking start time for object of interest.
	SensorTaskingStartTime time.Time `json:"sensorTaskingStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Time of last status change of the object of interest event.
	StatusDate time.Time `json:"statusDate,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Optional array of Onorbit IDs (idOnOrbit) representing satellites potentially
	// affected by this object of interest.
	AffectedObjects []string `json:"affectedObjects"`
	// Last reported apogee. The Orbit point furthest from the center of the earth in
	// kilometers.
	Apogee float64 `json:"apogee"`
	// Last reported argument of perigee. The argument of perigee is the angle in
	// degrees formed between the perigee and the ascending node. If the perigee would
	// occur at the ascending node, the argument of perigee would be 0.
	ArgOfPerigee float64 `json:"argOfPerigee"`
	// Last reported drag term for SGP4 orbital model, used for calculating decay
	// constants for altitude, eccentricity etc, measured in inverse earth radii.
	BStar float64 `json:"bStar"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Possible delta time applications for the object of interest, in seconds.
	DeltaTs []float64 `json:"deltaTs"`
	// Possible delta V applications for the object of interest, in km/sec.
	DeltaVs []float64 `json:"deltaVs"`
	// Description of the object of interest event.
	Description string `json:"description"`
	// Last reported eccentricity of the object. The orbital eccentricity of an
	// astronomical object is a parameter that determines the amount by which its orbit
	// around another body deviates from a perfect circle. A value of 0 is a circular
	// orbit, values between 0 and 1 form an elliptic orbit, 1 is a parabolic escape
	// orbit, and greater than 1 is a hyperbolic escape orbit.
	Eccentricity float64 `json:"eccentricity"`
	// Last reported elset epoch time in ISO 8601 UTC time, with microsecond precision.
	ElsetEpoch time.Time `json:"elsetEpoch" format:"date-time"`
	// Last reported inclination of the object. Inclination is the angle between the
	// equator and the orbit when looking from the center of the Earth. If the orbit
	// went exactly around the equator from left to right, then the inclination would
	// be 0. The inclination ranges from 0 to 180 degrees.
	Inclination float64 `json:"inclination"`
	// Last reported observation time in ISO 8601 UTC time, with microsecond precision.
	LastObTime time.Time `json:"lastObTime" format:"date-time"`
	// Manifolds associated with this object of interest.
	Manifolds []ObjectofinterestGetResponseManifold `json:"manifolds"`
	// Last reported meanAnomaly. Mean anomoly is where the satellite is in its orbital
	// path. The mean anomaly ranges from 0 to 360 degrees. The mean anomaly is
	// referenced to the perigee. If the satellite were at the perigee, the mean
	// anomaly would be 0.
	MeanAnomaly float64 `json:"meanAnomaly"`
	// Last reported mean motion of the object. Mean motion is the angular speed
	// required for a body to complete one orbit, assuming constant speed in a circular
	// orbit which completes in the same time as the variable speed, elliptical orbit
	// of the actual body. Measured in revolutions per day.
	MeanMotion float64 `json:"meanMotion"`
	// Last reported 2nd derivative of the mean motion with respect to time. Units are
	// revolutions per day cubed.
	MeanMotionDDot float64 `json:"meanMotionDDot"`
	// Last reported 1st derivative of the mean motion with respect to time. Units are
	// revolutions per day squared.
	MeanMotionDot float64 `json:"meanMotionDot"`
	// The time at which an attempted observation of the object of interest noticed it
	// was missing, in ISO 8601 UTC time, with microsecond precision.
	MissedObTime time.Time `json:"missedObTime" format:"date-time"`
	// Unique name of the object of interest event.
	Name string `json:"name"`
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
	// Last reported perigee. The orbit point nearest to the center of the earth in
	// kilometers.
	Perigee float64 `json:"perigee"`
	// Last reported orbit period. Period of the orbit is equal to inverse of mean
	// motion.
	Period float64 `json:"period"`
	// Priority of the object of interest as an integer (1=highest priority).
	Priority int64 `json:"priority"`
	// Last reported raan. Right ascension of the ascending node, or RAAN is the angle
	// as measured in degrees eastwards (or, as seen from the north, counterclockwise)
	// from the First Point of Aries to the ascending node, which is where the orbit
	// crosses the equator when traveling north.
	Raan float64 `json:"raan"`
	// The last reported revolution number. The value is incremented when a satellite
	// crosses the equator on an ascending pass.
	RevNo int64 `json:"revNo"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Last reported semi major axis, which is the sum of the periapsis and apoapsis
	// distances divided by two. For circular orbits, the semimajor axis is the
	// distance between the centers of the bodies, not the distance of the bodies from
	// the center of mass.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// Sensor tasking stop time for object of interest.
	SensorTaskingStopTime time.Time `json:"sensorTaskingStopTime" format:"date-time"`
	// Status of the object of interest event (e.g. OPEN, CLOSED, CANCELLED).
	Status string `json:"status"`
	// Last reported state vector epoch time in ISO 8601 UTC time, with microsecond
	// precision.
	SvEpoch time.Time `json:"svEpoch" format:"date-time"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Last reported x position of the object in km, in J2000 coordinates.
	X float64 `json:"x"`
	// Last reported x velocity of the object in km/sec, in J2000 coordinates.
	Xvel float64 `json:"xvel"`
	// Last reported y position of the object in km, in J2000 coordinates.
	Y float64 `json:"y"`
	// Last reported y velocity of the object in km/sec, in J2000 coordinates.
	Yvel float64 `json:"yvel"`
	// Last reported z position of the object in km, in J2000 coordinates.
	Z float64 `json:"z"`
	// Last reported z velocity of the object in km/sec, in J2000 coordinates.
	Zvel float64 `json:"zvel"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking  resp.Field
		DataMode               resp.Field
		IDOnOrbit              resp.Field
		SensorTaskingStartTime resp.Field
		Source                 resp.Field
		StatusDate             resp.Field
		ID                     resp.Field
		AffectedObjects        resp.Field
		Apogee                 resp.Field
		ArgOfPerigee           resp.Field
		BStar                  resp.Field
		CreatedAt              resp.Field
		CreatedBy              resp.Field
		DeltaTs                resp.Field
		DeltaVs                resp.Field
		Description            resp.Field
		Eccentricity           resp.Field
		ElsetEpoch             resp.Field
		Inclination            resp.Field
		LastObTime             resp.Field
		Manifolds              resp.Field
		MeanAnomaly            resp.Field
		MeanMotion             resp.Field
		MeanMotionDDot         resp.Field
		MeanMotionDot          resp.Field
		MissedObTime           resp.Field
		Name                   resp.Field
		OnOrbit                resp.Field
		Origin                 resp.Field
		OrigNetwork            resp.Field
		Perigee                resp.Field
		Period                 resp.Field
		Priority               resp.Field
		Raan                   resp.Field
		RevNo                  resp.Field
		SatNo                  resp.Field
		SemiMajorAxis          resp.Field
		SensorTaskingStopTime  resp.Field
		Status                 resp.Field
		SvEpoch                resp.Field
		UpdatedAt              resp.Field
		UpdatedBy              resp.Field
		X                      resp.Field
		Xvel                   resp.Field
		Y                      resp.Field
		Yvel                   resp.Field
		Z                      resp.Field
		Zvel                   resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectofinterestGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectofinterestGetResponse) UnmarshalJSON(data []byte) error {
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
type ObjectofinterestGetResponseDataMode string

const (
	ObjectofinterestGetResponseDataModeReal      ObjectofinterestGetResponseDataMode = "REAL"
	ObjectofinterestGetResponseDataModeTest      ObjectofinterestGetResponseDataMode = "TEST"
	ObjectofinterestGetResponseDataModeSimulated ObjectofinterestGetResponseDataMode = "SIMULATED"
	ObjectofinterestGetResponseDataModeExercise  ObjectofinterestGetResponseDataMode = "EXERCISE"
)

// A manifold represents a set of possible/theoretical orbits for an object of
// interest based on a delta V and delta T.
type ObjectofinterestGetResponseManifold struct {
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
	// ID of the parent object of interest.
	IDObjectOfInterest string `json:"idObjectOfInterest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Applied delta V duration for this manifold's calculations in seconds.
	DeltaT float64 `json:"deltaT"`
	// Applied delta V for this manifold's calculations, in km/sec.
	DeltaV float64 `json:"deltaV"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Status of the manifold and its associated ManifoldElsets (e.g. PENDING,
	// COMPLETE). PENDING status means element set generation is in progress and
	// COMPLETE indicates all ManifoldElsets have been generated.
	Status string `json:"status"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// Weight or probability of this manifold for prioritization purposes, between 0
	// and 1.
	Weight float64 `json:"weight"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDObjectOfInterest    resp.Field
		Source                resp.Field
		ID                    resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DeltaT                resp.Field
		DeltaV                resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Status                resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		Weight                resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectofinterestGetResponseManifold) RawJSON() string { return r.JSON.raw }
func (r *ObjectofinterestGetResponseManifold) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OnOrbit objects of interest, which include information about the last known
// state of the object.
type ObjectofinterestTupleResponse struct {
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
	DataMode ObjectofinterestTupleResponseDataMode `json:"dataMode,required"`
	// UUID of the parent Onorbit record.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Sensor tasking start time for object of interest.
	SensorTaskingStartTime time.Time `json:"sensorTaskingStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Time of last status change of the object of interest event.
	StatusDate time.Time `json:"statusDate,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Optional array of Onorbit IDs (idOnOrbit) representing satellites potentially
	// affected by this object of interest.
	AffectedObjects []string `json:"affectedObjects"`
	// Last reported apogee. The Orbit point furthest from the center of the earth in
	// kilometers.
	Apogee float64 `json:"apogee"`
	// Last reported argument of perigee. The argument of perigee is the angle in
	// degrees formed between the perigee and the ascending node. If the perigee would
	// occur at the ascending node, the argument of perigee would be 0.
	ArgOfPerigee float64 `json:"argOfPerigee"`
	// Last reported drag term for SGP4 orbital model, used for calculating decay
	// constants for altitude, eccentricity etc, measured in inverse earth radii.
	BStar float64 `json:"bStar"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Possible delta time applications for the object of interest, in seconds.
	DeltaTs []float64 `json:"deltaTs"`
	// Possible delta V applications for the object of interest, in km/sec.
	DeltaVs []float64 `json:"deltaVs"`
	// Description of the object of interest event.
	Description string `json:"description"`
	// Last reported eccentricity of the object. The orbital eccentricity of an
	// astronomical object is a parameter that determines the amount by which its orbit
	// around another body deviates from a perfect circle. A value of 0 is a circular
	// orbit, values between 0 and 1 form an elliptic orbit, 1 is a parabolic escape
	// orbit, and greater than 1 is a hyperbolic escape orbit.
	Eccentricity float64 `json:"eccentricity"`
	// Last reported elset epoch time in ISO 8601 UTC time, with microsecond precision.
	ElsetEpoch time.Time `json:"elsetEpoch" format:"date-time"`
	// Last reported inclination of the object. Inclination is the angle between the
	// equator and the orbit when looking from the center of the Earth. If the orbit
	// went exactly around the equator from left to right, then the inclination would
	// be 0. The inclination ranges from 0 to 180 degrees.
	Inclination float64 `json:"inclination"`
	// Last reported observation time in ISO 8601 UTC time, with microsecond precision.
	LastObTime time.Time `json:"lastObTime" format:"date-time"`
	// Manifolds associated with this object of interest.
	Manifolds []ObjectofinterestTupleResponseManifold `json:"manifolds"`
	// Last reported meanAnomaly. Mean anomoly is where the satellite is in its orbital
	// path. The mean anomaly ranges from 0 to 360 degrees. The mean anomaly is
	// referenced to the perigee. If the satellite were at the perigee, the mean
	// anomaly would be 0.
	MeanAnomaly float64 `json:"meanAnomaly"`
	// Last reported mean motion of the object. Mean motion is the angular speed
	// required for a body to complete one orbit, assuming constant speed in a circular
	// orbit which completes in the same time as the variable speed, elliptical orbit
	// of the actual body. Measured in revolutions per day.
	MeanMotion float64 `json:"meanMotion"`
	// Last reported 2nd derivative of the mean motion with respect to time. Units are
	// revolutions per day cubed.
	MeanMotionDDot float64 `json:"meanMotionDDot"`
	// Last reported 1st derivative of the mean motion with respect to time. Units are
	// revolutions per day squared.
	MeanMotionDot float64 `json:"meanMotionDot"`
	// The time at which an attempted observation of the object of interest noticed it
	// was missing, in ISO 8601 UTC time, with microsecond precision.
	MissedObTime time.Time `json:"missedObTime" format:"date-time"`
	// Unique name of the object of interest event.
	Name string `json:"name"`
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
	// Last reported perigee. The orbit point nearest to the center of the earth in
	// kilometers.
	Perigee float64 `json:"perigee"`
	// Last reported orbit period. Period of the orbit is equal to inverse of mean
	// motion.
	Period float64 `json:"period"`
	// Priority of the object of interest as an integer (1=highest priority).
	Priority int64 `json:"priority"`
	// Last reported raan. Right ascension of the ascending node, or RAAN is the angle
	// as measured in degrees eastwards (or, as seen from the north, counterclockwise)
	// from the First Point of Aries to the ascending node, which is where the orbit
	// crosses the equator when traveling north.
	Raan float64 `json:"raan"`
	// The last reported revolution number. The value is incremented when a satellite
	// crosses the equator on an ascending pass.
	RevNo int64 `json:"revNo"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Last reported semi major axis, which is the sum of the periapsis and apoapsis
	// distances divided by two. For circular orbits, the semimajor axis is the
	// distance between the centers of the bodies, not the distance of the bodies from
	// the center of mass.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// Sensor tasking stop time for object of interest.
	SensorTaskingStopTime time.Time `json:"sensorTaskingStopTime" format:"date-time"`
	// Status of the object of interest event (e.g. OPEN, CLOSED, CANCELLED).
	Status string `json:"status"`
	// Last reported state vector epoch time in ISO 8601 UTC time, with microsecond
	// precision.
	SvEpoch time.Time `json:"svEpoch" format:"date-time"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Last reported x position of the object in km, in J2000 coordinates.
	X float64 `json:"x"`
	// Last reported x velocity of the object in km/sec, in J2000 coordinates.
	Xvel float64 `json:"xvel"`
	// Last reported y position of the object in km, in J2000 coordinates.
	Y float64 `json:"y"`
	// Last reported y velocity of the object in km/sec, in J2000 coordinates.
	Yvel float64 `json:"yvel"`
	// Last reported z position of the object in km, in J2000 coordinates.
	Z float64 `json:"z"`
	// Last reported z velocity of the object in km/sec, in J2000 coordinates.
	Zvel float64 `json:"zvel"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking  resp.Field
		DataMode               resp.Field
		IDOnOrbit              resp.Field
		SensorTaskingStartTime resp.Field
		Source                 resp.Field
		StatusDate             resp.Field
		ID                     resp.Field
		AffectedObjects        resp.Field
		Apogee                 resp.Field
		ArgOfPerigee           resp.Field
		BStar                  resp.Field
		CreatedAt              resp.Field
		CreatedBy              resp.Field
		DeltaTs                resp.Field
		DeltaVs                resp.Field
		Description            resp.Field
		Eccentricity           resp.Field
		ElsetEpoch             resp.Field
		Inclination            resp.Field
		LastObTime             resp.Field
		Manifolds              resp.Field
		MeanAnomaly            resp.Field
		MeanMotion             resp.Field
		MeanMotionDDot         resp.Field
		MeanMotionDot          resp.Field
		MissedObTime           resp.Field
		Name                   resp.Field
		OnOrbit                resp.Field
		Origin                 resp.Field
		OrigNetwork            resp.Field
		Perigee                resp.Field
		Period                 resp.Field
		Priority               resp.Field
		Raan                   resp.Field
		RevNo                  resp.Field
		SatNo                  resp.Field
		SemiMajorAxis          resp.Field
		SensorTaskingStopTime  resp.Field
		Status                 resp.Field
		SvEpoch                resp.Field
		UpdatedAt              resp.Field
		UpdatedBy              resp.Field
		X                      resp.Field
		Xvel                   resp.Field
		Y                      resp.Field
		Yvel                   resp.Field
		Z                      resp.Field
		Zvel                   resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectofinterestTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectofinterestTupleResponse) UnmarshalJSON(data []byte) error {
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
type ObjectofinterestTupleResponseDataMode string

const (
	ObjectofinterestTupleResponseDataModeReal      ObjectofinterestTupleResponseDataMode = "REAL"
	ObjectofinterestTupleResponseDataModeTest      ObjectofinterestTupleResponseDataMode = "TEST"
	ObjectofinterestTupleResponseDataModeSimulated ObjectofinterestTupleResponseDataMode = "SIMULATED"
	ObjectofinterestTupleResponseDataModeExercise  ObjectofinterestTupleResponseDataMode = "EXERCISE"
)

// A manifold represents a set of possible/theoretical orbits for an object of
// interest based on a delta V and delta T.
type ObjectofinterestTupleResponseManifold struct {
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
	// ID of the parent object of interest.
	IDObjectOfInterest string `json:"idObjectOfInterest,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Applied delta V duration for this manifold's calculations in seconds.
	DeltaT float64 `json:"deltaT"`
	// Applied delta V for this manifold's calculations, in km/sec.
	DeltaV float64 `json:"deltaV"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Status of the manifold and its associated ManifoldElsets (e.g. PENDING,
	// COMPLETE). PENDING status means element set generation is in progress and
	// COMPLETE indicates all ManifoldElsets have been generated.
	Status string `json:"status"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// Weight or probability of this manifold for prioritization purposes, between 0
	// and 1.
	Weight float64 `json:"weight"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDObjectOfInterest    resp.Field
		Source                resp.Field
		ID                    resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DeltaT                resp.Field
		DeltaV                resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Status                resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		Weight                resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectofinterestTupleResponseManifold) RawJSON() string { return r.JSON.raw }
func (r *ObjectofinterestTupleResponseManifold) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectofinterestNewParams struct {
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
	DataMode ObjectofinterestNewParamsDataMode `json:"dataMode,omitzero,required"`
	// UUID of the parent Onorbit record.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Sensor tasking start time for object of interest.
	SensorTaskingStartTime time.Time `json:"sensorTaskingStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Time of last status change of the object of interest event.
	StatusDate time.Time `json:"statusDate,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Last reported apogee. The Orbit point furthest from the center of the earth in
	// kilometers.
	Apogee param.Opt[float64] `json:"apogee,omitzero"`
	// Last reported argument of perigee. The argument of perigee is the angle in
	// degrees formed between the perigee and the ascending node. If the perigee would
	// occur at the ascending node, the argument of perigee would be 0.
	ArgOfPerigee param.Opt[float64] `json:"argOfPerigee,omitzero"`
	// Last reported drag term for SGP4 orbital model, used for calculating decay
	// constants for altitude, eccentricity etc, measured in inverse earth radii.
	BStar param.Opt[float64] `json:"bStar,omitzero"`
	// Description of the object of interest event.
	Description param.Opt[string] `json:"description,omitzero"`
	// Last reported eccentricity of the object. The orbital eccentricity of an
	// astronomical object is a parameter that determines the amount by which its orbit
	// around another body deviates from a perfect circle. A value of 0 is a circular
	// orbit, values between 0 and 1 form an elliptic orbit, 1 is a parabolic escape
	// orbit, and greater than 1 is a hyperbolic escape orbit.
	Eccentricity param.Opt[float64] `json:"eccentricity,omitzero"`
	// Last reported elset epoch time in ISO 8601 UTC time, with microsecond precision.
	ElsetEpoch param.Opt[time.Time] `json:"elsetEpoch,omitzero" format:"date-time"`
	// Last reported inclination of the object. Inclination is the angle between the
	// equator and the orbit when looking from the center of the Earth. If the orbit
	// went exactly around the equator from left to right, then the inclination would
	// be 0. The inclination ranges from 0 to 180 degrees.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Last reported observation time in ISO 8601 UTC time, with microsecond precision.
	LastObTime param.Opt[time.Time] `json:"lastObTime,omitzero" format:"date-time"`
	// Last reported meanAnomaly. Mean anomoly is where the satellite is in its orbital
	// path. The mean anomaly ranges from 0 to 360 degrees. The mean anomaly is
	// referenced to the perigee. If the satellite were at the perigee, the mean
	// anomaly would be 0.
	MeanAnomaly param.Opt[float64] `json:"meanAnomaly,omitzero"`
	// Last reported mean motion of the object. Mean motion is the angular speed
	// required for a body to complete one orbit, assuming constant speed in a circular
	// orbit which completes in the same time as the variable speed, elliptical orbit
	// of the actual body. Measured in revolutions per day.
	MeanMotion param.Opt[float64] `json:"meanMotion,omitzero"`
	// Last reported 2nd derivative of the mean motion with respect to time. Units are
	// revolutions per day cubed.
	MeanMotionDDot param.Opt[float64] `json:"meanMotionDDot,omitzero"`
	// Last reported 1st derivative of the mean motion with respect to time. Units are
	// revolutions per day squared.
	MeanMotionDot param.Opt[float64] `json:"meanMotionDot,omitzero"`
	// The time at which an attempted observation of the object of interest noticed it
	// was missing, in ISO 8601 UTC time, with microsecond precision.
	MissedObTime param.Opt[time.Time] `json:"missedObTime,omitzero" format:"date-time"`
	// Unique name of the object of interest event.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Last reported perigee. The orbit point nearest to the center of the earth in
	// kilometers.
	Perigee param.Opt[float64] `json:"perigee,omitzero"`
	// Last reported orbit period. Period of the orbit is equal to inverse of mean
	// motion.
	Period param.Opt[float64] `json:"period,omitzero"`
	// Priority of the object of interest as an integer (1=highest priority).
	Priority param.Opt[int64] `json:"priority,omitzero"`
	// Last reported raan. Right ascension of the ascending node, or RAAN is the angle
	// as measured in degrees eastwards (or, as seen from the north, counterclockwise)
	// from the First Point of Aries to the ascending node, which is where the orbit
	// crosses the equator when traveling north.
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// The last reported revolution number. The value is incremented when a satellite
	// crosses the equator on an ascending pass.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Last reported semi major axis, which is the sum of the periapsis and apoapsis
	// distances divided by two. For circular orbits, the semimajor axis is the
	// distance between the centers of the bodies, not the distance of the bodies from
	// the center of mass.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	// Sensor tasking stop time for object of interest.
	SensorTaskingStopTime param.Opt[time.Time] `json:"sensorTaskingStopTime,omitzero" format:"date-time"`
	// Status of the object of interest event (e.g. OPEN, CLOSED, CANCELLED).
	Status param.Opt[string] `json:"status,omitzero"`
	// Last reported state vector epoch time in ISO 8601 UTC time, with microsecond
	// precision.
	SvEpoch param.Opt[time.Time] `json:"svEpoch,omitzero" format:"date-time"`
	// Last reported x position of the object in km, in J2000 coordinates.
	X param.Opt[float64] `json:"x,omitzero"`
	// Last reported x velocity of the object in km/sec, in J2000 coordinates.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Last reported y position of the object in km, in J2000 coordinates.
	Y param.Opt[float64] `json:"y,omitzero"`
	// Last reported y velocity of the object in km/sec, in J2000 coordinates.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Last reported z position of the object in km, in J2000 coordinates.
	Z param.Opt[float64] `json:"z,omitzero"`
	// Last reported z velocity of the object in km/sec, in J2000 coordinates.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Optional array of Onorbit IDs (idOnOrbit) representing satellites potentially
	// affected by this object of interest.
	AffectedObjects []string `json:"affectedObjects,omitzero"`
	// Possible delta time applications for the object of interest, in seconds.
	DeltaTs []float64 `json:"deltaTs,omitzero"`
	// Possible delta V applications for the object of interest, in km/sec.
	DeltaVs []float64 `json:"deltaVs,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ObjectofinterestNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ObjectofinterestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectofinterestNewParams
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
type ObjectofinterestNewParamsDataMode string

const (
	ObjectofinterestNewParamsDataModeReal      ObjectofinterestNewParamsDataMode = "REAL"
	ObjectofinterestNewParamsDataModeTest      ObjectofinterestNewParamsDataMode = "TEST"
	ObjectofinterestNewParamsDataModeSimulated ObjectofinterestNewParamsDataMode = "SIMULATED"
	ObjectofinterestNewParamsDataModeExercise  ObjectofinterestNewParamsDataMode = "EXERCISE"
)

type ObjectofinterestUpdateParams struct {
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
	DataMode ObjectofinterestUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// UUID of the parent Onorbit record.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Sensor tasking start time for object of interest.
	SensorTaskingStartTime time.Time `json:"sensorTaskingStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Time of last status change of the object of interest event.
	StatusDate time.Time `json:"statusDate,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Last reported apogee. The Orbit point furthest from the center of the earth in
	// kilometers.
	Apogee param.Opt[float64] `json:"apogee,omitzero"`
	// Last reported argument of perigee. The argument of perigee is the angle in
	// degrees formed between the perigee and the ascending node. If the perigee would
	// occur at the ascending node, the argument of perigee would be 0.
	ArgOfPerigee param.Opt[float64] `json:"argOfPerigee,omitzero"`
	// Last reported drag term for SGP4 orbital model, used for calculating decay
	// constants for altitude, eccentricity etc, measured in inverse earth radii.
	BStar param.Opt[float64] `json:"bStar,omitzero"`
	// Description of the object of interest event.
	Description param.Opt[string] `json:"description,omitzero"`
	// Last reported eccentricity of the object. The orbital eccentricity of an
	// astronomical object is a parameter that determines the amount by which its orbit
	// around another body deviates from a perfect circle. A value of 0 is a circular
	// orbit, values between 0 and 1 form an elliptic orbit, 1 is a parabolic escape
	// orbit, and greater than 1 is a hyperbolic escape orbit.
	Eccentricity param.Opt[float64] `json:"eccentricity,omitzero"`
	// Last reported elset epoch time in ISO 8601 UTC time, with microsecond precision.
	ElsetEpoch param.Opt[time.Time] `json:"elsetEpoch,omitzero" format:"date-time"`
	// Last reported inclination of the object. Inclination is the angle between the
	// equator and the orbit when looking from the center of the Earth. If the orbit
	// went exactly around the equator from left to right, then the inclination would
	// be 0. The inclination ranges from 0 to 180 degrees.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Last reported observation time in ISO 8601 UTC time, with microsecond precision.
	LastObTime param.Opt[time.Time] `json:"lastObTime,omitzero" format:"date-time"`
	// Last reported meanAnomaly. Mean anomoly is where the satellite is in its orbital
	// path. The mean anomaly ranges from 0 to 360 degrees. The mean anomaly is
	// referenced to the perigee. If the satellite were at the perigee, the mean
	// anomaly would be 0.
	MeanAnomaly param.Opt[float64] `json:"meanAnomaly,omitzero"`
	// Last reported mean motion of the object. Mean motion is the angular speed
	// required for a body to complete one orbit, assuming constant speed in a circular
	// orbit which completes in the same time as the variable speed, elliptical orbit
	// of the actual body. Measured in revolutions per day.
	MeanMotion param.Opt[float64] `json:"meanMotion,omitzero"`
	// Last reported 2nd derivative of the mean motion with respect to time. Units are
	// revolutions per day cubed.
	MeanMotionDDot param.Opt[float64] `json:"meanMotionDDot,omitzero"`
	// Last reported 1st derivative of the mean motion with respect to time. Units are
	// revolutions per day squared.
	MeanMotionDot param.Opt[float64] `json:"meanMotionDot,omitzero"`
	// The time at which an attempted observation of the object of interest noticed it
	// was missing, in ISO 8601 UTC time, with microsecond precision.
	MissedObTime param.Opt[time.Time] `json:"missedObTime,omitzero" format:"date-time"`
	// Unique name of the object of interest event.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Last reported perigee. The orbit point nearest to the center of the earth in
	// kilometers.
	Perigee param.Opt[float64] `json:"perigee,omitzero"`
	// Last reported orbit period. Period of the orbit is equal to inverse of mean
	// motion.
	Period param.Opt[float64] `json:"period,omitzero"`
	// Priority of the object of interest as an integer (1=highest priority).
	Priority param.Opt[int64] `json:"priority,omitzero"`
	// Last reported raan. Right ascension of the ascending node, or RAAN is the angle
	// as measured in degrees eastwards (or, as seen from the north, counterclockwise)
	// from the First Point of Aries to the ascending node, which is where the orbit
	// crosses the equator when traveling north.
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// The last reported revolution number. The value is incremented when a satellite
	// crosses the equator on an ascending pass.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Last reported semi major axis, which is the sum of the periapsis and apoapsis
	// distances divided by two. For circular orbits, the semimajor axis is the
	// distance between the centers of the bodies, not the distance of the bodies from
	// the center of mass.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	// Sensor tasking stop time for object of interest.
	SensorTaskingStopTime param.Opt[time.Time] `json:"sensorTaskingStopTime,omitzero" format:"date-time"`
	// Status of the object of interest event (e.g. OPEN, CLOSED, CANCELLED).
	Status param.Opt[string] `json:"status,omitzero"`
	// Last reported state vector epoch time in ISO 8601 UTC time, with microsecond
	// precision.
	SvEpoch param.Opt[time.Time] `json:"svEpoch,omitzero" format:"date-time"`
	// Last reported x position of the object in km, in J2000 coordinates.
	X param.Opt[float64] `json:"x,omitzero"`
	// Last reported x velocity of the object in km/sec, in J2000 coordinates.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Last reported y position of the object in km, in J2000 coordinates.
	Y param.Opt[float64] `json:"y,omitzero"`
	// Last reported y velocity of the object in km/sec, in J2000 coordinates.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Last reported z position of the object in km, in J2000 coordinates.
	Z param.Opt[float64] `json:"z,omitzero"`
	// Last reported z velocity of the object in km/sec, in J2000 coordinates.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Optional array of Onorbit IDs (idOnOrbit) representing satellites potentially
	// affected by this object of interest.
	AffectedObjects []string `json:"affectedObjects,omitzero"`
	// Possible delta time applications for the object of interest, in seconds.
	DeltaTs []float64 `json:"deltaTs,omitzero"`
	// Possible delta V applications for the object of interest, in km/sec.
	DeltaVs []float64 `json:"deltaVs,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ObjectofinterestUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ObjectofinterestUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectofinterestUpdateParams
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
type ObjectofinterestUpdateParamsDataMode string

const (
	ObjectofinterestUpdateParamsDataModeReal      ObjectofinterestUpdateParamsDataMode = "REAL"
	ObjectofinterestUpdateParamsDataModeTest      ObjectofinterestUpdateParamsDataMode = "TEST"
	ObjectofinterestUpdateParamsDataModeSimulated ObjectofinterestUpdateParamsDataMode = "SIMULATED"
	ObjectofinterestUpdateParamsDataModeExercise  ObjectofinterestUpdateParamsDataMode = "EXERCISE"
)

type ObjectofinterestListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ObjectofinterestListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ObjectofinterestListParams]'s query parameters as
// `url.Values`.
func (r ObjectofinterestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectofinterestCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ObjectofinterestCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ObjectofinterestCountParams]'s query parameters as
// `url.Values`.
func (r ObjectofinterestCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectofinterestGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ObjectofinterestGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ObjectofinterestGetParams]'s query parameters as
// `url.Values`.
func (r ObjectofinterestGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectofinterestTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ObjectofinterestTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ObjectofinterestTupleParams]'s query parameters as
// `url.Values`.
func (r ObjectofinterestTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
