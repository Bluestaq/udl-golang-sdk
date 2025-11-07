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

// ManifoldelsetService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManifoldelsetService] method instead.
type ManifoldelsetService struct {
	Options []option.RequestOption
}

// NewManifoldelsetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewManifoldelsetService(opts ...option.RequestOption) (r ManifoldelsetService) {
	r = ManifoldelsetService{}
	r.Options = opts
	return
}

// Service operation to take a single ManifoldElset as a POST body and ingest into
// the database. A ManifoldElset represents theoretical Keplarian orbital elements
// belonging to an object of interest's manifold describing a possible/theoretical
// orbit for an object of interest for tasking purposes. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *ManifoldelsetService) New(ctx context.Context, body ManifoldelsetNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/manifoldelset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single ManifoldElset. A ManifoldElset represents
// theoretical Keplarian orbital elements belonging to an object of interest's
// manifold describing a possible/theoretical orbit for an object of interest for
// tasking purposes. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *ManifoldelsetService) Update(ctx context.Context, id string, body ManifoldelsetUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/manifoldelset/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ManifoldelsetService) List(ctx context.Context, query ManifoldelsetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ManifoldelsetListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/manifoldelset"
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
func (r *ManifoldelsetService) ListAutoPaging(ctx context.Context, query ManifoldelsetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ManifoldelsetListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a ManifoldElset object specified by the passed ID
// path parameter. A ManifoldElset represents theoretical Keplarian orbital
// elements belonging to an object of interest's manifold describing a
// possible/theoretical orbit for an object of interest for tasking purposes. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *ManifoldelsetService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/manifoldelset/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ManifoldelsetService) Count(ctx context.Context, query ManifoldelsetCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/manifoldelset/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple ManifoldElsets as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *ManifoldelsetService) NewBulk(ctx context.Context, body ManifoldelsetNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/manifoldelset/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single ManifoldElset record by its unique ID passed
// as a path parameter. A ManifoldElset represents theoretical Keplarian orbital
// elements belonging to an object of interest's manifold describing a
// possible/theoretical orbit for an object of interest for tasking purposes.
func (r *ManifoldelsetService) Get(ctx context.Context, id string, query ManifoldelsetGetParams, opts ...option.RequestOption) (res *ManifoldelsetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/manifoldelset/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ManifoldelsetService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *ManifoldelsetQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/manifoldelset/queryhelp"
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
func (r *ManifoldelsetService) Tuple(ctx context.Context, query ManifoldelsetTupleParams, opts ...option.RequestOption) (res *[]ManifoldelsetTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/manifoldelset/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Theoretical Keplarian orbital elements belonging to an object of interest's
// manifold describing a possible/theoretical orbit for an object of interest for
// tasking purposes.
type ManifoldelsetListResponse struct {
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
	DataMode ManifoldelsetListResponseDataMode `json:"dataMode,required"`
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Identifier of the parent Manifold record.
	IDManifold string `json:"idManifold,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// A placeholder satellite number and not a true NORAD catalog number.
	TmpSatNo int64 `json:"tmpSatNo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The Orbit point furthest from the center of the earth in kilometers.
	Apogee float64 `json:"apogee"`
	// The argument of perigee is the angle in degrees formed between the perigee and
	// the ascending node. If the perigee would occur at the ascending node, the
	// argument of perigee would be 0.
	ArgOfPerigee float64 `json:"argOfPerigee"`
	// The drag term for SGP4 orbital model, used for calculating decay constants for
	// altitude, eccentricity etc, measured in inverse earth radii.
	BStar float64 `json:"bStar"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle. A value of 0 is a circular orbit, values between 0 and 1 form an
	// elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a
	// hyperbolic escape orbit.
	Eccentricity float64 `json:"eccentricity"`
	// The angle between the equator and the orbit when looking from the center of the
	// Earth. If the orbit went exactly around the equator from left to right, then the
	// inclination would be 0. The inclination ranges from 0 to 180 degrees.
	Inclination float64 `json:"inclination"`
	// Read only derived/generated line1 of a legacy TLE (two line element set) format,
	// ignored on create/edit operations.
	Line1 string `json:"line1"`
	// Read only derived/generated line2 of a legacy TLE (two line element set) format,
	// ignored on create/edit operations.
	Line2 string `json:"line2"`
	// Where the satellite is in its orbital path. The mean anomaly ranges from 0 to
	// 360 degrees. The mean anomaly is referenced to the perigee. If the satellite
	// were at the perigee, the mean anomaly would be 0.
	MeanAnomaly float64 `json:"meanAnomaly"`
	// The constant angular speed required for the body to complete one circular orbit
	// in the same amount of time as the actual elliptical orbit with variable speed.
	// Measured in revolutions per day.
	MeanMotion float64 `json:"meanMotion"`
	// 2nd derivative of the mean motion with respect to time. Units are revolutions
	// per day cubed.
	MeanMotionDDot float64 `json:"meanMotionDDot"`
	// 1st derivative of the mean motion with respect to time. Units are revolutions
	// per day squared.
	MeanMotionDot float64 `json:"meanMotionDot"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The orbit point nearest to the center of the earth in kilometers.
	Perigee float64 `json:"perigee"`
	// Period of the orbit equal to inverse of mean motion.
	Period float64 `json:"period"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node, which is where the orbit crosses the
	// equator when traveling north.
	Raan float64 `json:"raan"`
	// The current revolution number. The value is incremented when a satellite crosses
	// the equator on an ascending pass.
	RevNo int64 `json:"revNo"`
	// The sum of the periapsis and apoapsis distances divided by two. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies,
	// not the distance of the bodies from the center of mass.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Epoch                 respjson.Field
		IDManifold            respjson.Field
		Source                respjson.Field
		TmpSatNo              respjson.Field
		ID                    respjson.Field
		Apogee                respjson.Field
		ArgOfPerigee          respjson.Field
		BStar                 respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Eccentricity          respjson.Field
		Inclination           respjson.Field
		Line1                 respjson.Field
		Line2                 respjson.Field
		MeanAnomaly           respjson.Field
		MeanMotion            respjson.Field
		MeanMotionDDot        respjson.Field
		MeanMotionDot         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Perigee               respjson.Field
		Period                respjson.Field
		Raan                  respjson.Field
		RevNo                 respjson.Field
		SemiMajorAxis         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManifoldelsetListResponse) RawJSON() string { return r.JSON.raw }
func (r *ManifoldelsetListResponse) UnmarshalJSON(data []byte) error {
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
type ManifoldelsetListResponseDataMode string

const (
	ManifoldelsetListResponseDataModeReal      ManifoldelsetListResponseDataMode = "REAL"
	ManifoldelsetListResponseDataModeTest      ManifoldelsetListResponseDataMode = "TEST"
	ManifoldelsetListResponseDataModeSimulated ManifoldelsetListResponseDataMode = "SIMULATED"
	ManifoldelsetListResponseDataModeExercise  ManifoldelsetListResponseDataMode = "EXERCISE"
)

// Theoretical Keplarian orbital elements belonging to an object of interest's
// manifold describing a possible/theoretical orbit for an object of interest for
// tasking purposes.
type ManifoldelsetGetResponse struct {
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
	DataMode ManifoldelsetGetResponseDataMode `json:"dataMode,required"`
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Identifier of the parent Manifold record.
	IDManifold string `json:"idManifold,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// A placeholder satellite number and not a true NORAD catalog number.
	TmpSatNo int64 `json:"tmpSatNo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The Orbit point furthest from the center of the earth in kilometers.
	Apogee float64 `json:"apogee"`
	// The argument of perigee is the angle in degrees formed between the perigee and
	// the ascending node. If the perigee would occur at the ascending node, the
	// argument of perigee would be 0.
	ArgOfPerigee float64 `json:"argOfPerigee"`
	// The drag term for SGP4 orbital model, used for calculating decay constants for
	// altitude, eccentricity etc, measured in inverse earth radii.
	BStar float64 `json:"bStar"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle. A value of 0 is a circular orbit, values between 0 and 1 form an
	// elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a
	// hyperbolic escape orbit.
	Eccentricity float64 `json:"eccentricity"`
	// The angle between the equator and the orbit when looking from the center of the
	// Earth. If the orbit went exactly around the equator from left to right, then the
	// inclination would be 0. The inclination ranges from 0 to 180 degrees.
	Inclination float64 `json:"inclination"`
	// Read only derived/generated line1 of a legacy TLE (two line element set) format,
	// ignored on create/edit operations.
	Line1 string `json:"line1"`
	// Read only derived/generated line2 of a legacy TLE (two line element set) format,
	// ignored on create/edit operations.
	Line2 string `json:"line2"`
	// Where the satellite is in its orbital path. The mean anomaly ranges from 0 to
	// 360 degrees. The mean anomaly is referenced to the perigee. If the satellite
	// were at the perigee, the mean anomaly would be 0.
	MeanAnomaly float64 `json:"meanAnomaly"`
	// The constant angular speed required for the body to complete one circular orbit
	// in the same amount of time as the actual elliptical orbit with variable speed.
	// Measured in revolutions per day.
	MeanMotion float64 `json:"meanMotion"`
	// 2nd derivative of the mean motion with respect to time. Units are revolutions
	// per day cubed.
	MeanMotionDDot float64 `json:"meanMotionDDot"`
	// 1st derivative of the mean motion with respect to time. Units are revolutions
	// per day squared.
	MeanMotionDot float64 `json:"meanMotionDot"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The orbit point nearest to the center of the earth in kilometers.
	Perigee float64 `json:"perigee"`
	// Period of the orbit equal to inverse of mean motion.
	Period float64 `json:"period"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node, which is where the orbit crosses the
	// equator when traveling north.
	Raan float64 `json:"raan"`
	// The current revolution number. The value is incremented when a satellite crosses
	// the equator on an ascending pass.
	RevNo int64 `json:"revNo"`
	// The sum of the periapsis and apoapsis distances divided by two. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies,
	// not the distance of the bodies from the center of mass.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Epoch                 respjson.Field
		IDManifold            respjson.Field
		Source                respjson.Field
		TmpSatNo              respjson.Field
		ID                    respjson.Field
		Apogee                respjson.Field
		ArgOfPerigee          respjson.Field
		BStar                 respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Eccentricity          respjson.Field
		Inclination           respjson.Field
		Line1                 respjson.Field
		Line2                 respjson.Field
		MeanAnomaly           respjson.Field
		MeanMotion            respjson.Field
		MeanMotionDDot        respjson.Field
		MeanMotionDot         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Perigee               respjson.Field
		Period                respjson.Field
		Raan                  respjson.Field
		RevNo                 respjson.Field
		SemiMajorAxis         respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManifoldelsetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ManifoldelsetGetResponse) UnmarshalJSON(data []byte) error {
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
type ManifoldelsetGetResponseDataMode string

const (
	ManifoldelsetGetResponseDataModeReal      ManifoldelsetGetResponseDataMode = "REAL"
	ManifoldelsetGetResponseDataModeTest      ManifoldelsetGetResponseDataMode = "TEST"
	ManifoldelsetGetResponseDataModeSimulated ManifoldelsetGetResponseDataMode = "SIMULATED"
	ManifoldelsetGetResponseDataModeExercise  ManifoldelsetGetResponseDataMode = "EXERCISE"
)

type ManifoldelsetQueryhelpResponse struct {
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
func (r ManifoldelsetQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *ManifoldelsetQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Theoretical Keplarian orbital elements belonging to an object of interest's
// manifold describing a possible/theoretical orbit for an object of interest for
// tasking purposes.
type ManifoldelsetTupleResponse struct {
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
	DataMode ManifoldelsetTupleResponseDataMode `json:"dataMode,required"`
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Identifier of the parent Manifold record.
	IDManifold string `json:"idManifold,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// A placeholder satellite number and not a true NORAD catalog number.
	TmpSatNo int64 `json:"tmpSatNo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The Orbit point furthest from the center of the earth in kilometers.
	Apogee float64 `json:"apogee"`
	// The argument of perigee is the angle in degrees formed between the perigee and
	// the ascending node. If the perigee would occur at the ascending node, the
	// argument of perigee would be 0.
	ArgOfPerigee float64 `json:"argOfPerigee"`
	// The drag term for SGP4 orbital model, used for calculating decay constants for
	// altitude, eccentricity etc, measured in inverse earth radii.
	BStar float64 `json:"bStar"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle. A value of 0 is a circular orbit, values between 0 and 1 form an
	// elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a
	// hyperbolic escape orbit.
	Eccentricity float64 `json:"eccentricity"`
	// The angle between the equator and the orbit when looking from the center of the
	// Earth. If the orbit went exactly around the equator from left to right, then the
	// inclination would be 0. The inclination ranges from 0 to 180 degrees.
	Inclination float64 `json:"inclination"`
	// Read only derived/generated line1 of a legacy TLE (two line element set) format,
	// ignored on create/edit operations.
	Line1 string `json:"line1"`
	// Read only derived/generated line2 of a legacy TLE (two line element set) format,
	// ignored on create/edit operations.
	Line2 string `json:"line2"`
	// Where the satellite is in its orbital path. The mean anomaly ranges from 0 to
	// 360 degrees. The mean anomaly is referenced to the perigee. If the satellite
	// were at the perigee, the mean anomaly would be 0.
	MeanAnomaly float64 `json:"meanAnomaly"`
	// The constant angular speed required for the body to complete one circular orbit
	// in the same amount of time as the actual elliptical orbit with variable speed.
	// Measured in revolutions per day.
	MeanMotion float64 `json:"meanMotion"`
	// 2nd derivative of the mean motion with respect to time. Units are revolutions
	// per day cubed.
	MeanMotionDDot float64 `json:"meanMotionDDot"`
	// 1st derivative of the mean motion with respect to time. Units are revolutions
	// per day squared.
	MeanMotionDot float64 `json:"meanMotionDot"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The orbit point nearest to the center of the earth in kilometers.
	Perigee float64 `json:"perigee"`
	// Period of the orbit equal to inverse of mean motion.
	Period float64 `json:"period"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node, which is where the orbit crosses the
	// equator when traveling north.
	Raan float64 `json:"raan"`
	// The current revolution number. The value is incremented when a satellite crosses
	// the equator on an ascending pass.
	RevNo int64 `json:"revNo"`
	// The sum of the periapsis and apoapsis distances divided by two. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies,
	// not the distance of the bodies from the center of mass.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Epoch                 respjson.Field
		IDManifold            respjson.Field
		Source                respjson.Field
		TmpSatNo              respjson.Field
		ID                    respjson.Field
		Apogee                respjson.Field
		ArgOfPerigee          respjson.Field
		BStar                 respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Eccentricity          respjson.Field
		Inclination           respjson.Field
		Line1                 respjson.Field
		Line2                 respjson.Field
		MeanAnomaly           respjson.Field
		MeanMotion            respjson.Field
		MeanMotionDDot        respjson.Field
		MeanMotionDot         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Perigee               respjson.Field
		Period                respjson.Field
		Raan                  respjson.Field
		RevNo                 respjson.Field
		SemiMajorAxis         respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManifoldelsetTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ManifoldelsetTupleResponse) UnmarshalJSON(data []byte) error {
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
type ManifoldelsetTupleResponseDataMode string

const (
	ManifoldelsetTupleResponseDataModeReal      ManifoldelsetTupleResponseDataMode = "REAL"
	ManifoldelsetTupleResponseDataModeTest      ManifoldelsetTupleResponseDataMode = "TEST"
	ManifoldelsetTupleResponseDataModeSimulated ManifoldelsetTupleResponseDataMode = "SIMULATED"
	ManifoldelsetTupleResponseDataModeExercise  ManifoldelsetTupleResponseDataMode = "EXERCISE"
)

type ManifoldelsetNewParams struct {
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
	DataMode ManifoldelsetNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Identifier of the parent Manifold record.
	IDManifold string `json:"idManifold,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// A placeholder satellite number and not a true NORAD catalog number.
	TmpSatNo int64 `json:"tmpSatNo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Orbit point furthest from the center of the earth in kilometers.
	Apogee param.Opt[float64] `json:"apogee,omitzero"`
	// The argument of perigee is the angle in degrees formed between the perigee and
	// the ascending node. If the perigee would occur at the ascending node, the
	// argument of perigee would be 0.
	ArgOfPerigee param.Opt[float64] `json:"argOfPerigee,omitzero"`
	// The drag term for SGP4 orbital model, used for calculating decay constants for
	// altitude, eccentricity etc, measured in inverse earth radii.
	BStar param.Opt[float64] `json:"bStar,omitzero"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle. A value of 0 is a circular orbit, values between 0 and 1 form an
	// elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a
	// hyperbolic escape orbit.
	Eccentricity param.Opt[float64] `json:"eccentricity,omitzero"`
	// The angle between the equator and the orbit when looking from the center of the
	// Earth. If the orbit went exactly around the equator from left to right, then the
	// inclination would be 0. The inclination ranges from 0 to 180 degrees.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Where the satellite is in its orbital path. The mean anomaly ranges from 0 to
	// 360 degrees. The mean anomaly is referenced to the perigee. If the satellite
	// were at the perigee, the mean anomaly would be 0.
	MeanAnomaly param.Opt[float64] `json:"meanAnomaly,omitzero"`
	// The constant angular speed required for the body to complete one circular orbit
	// in the same amount of time as the actual elliptical orbit with variable speed.
	// Measured in revolutions per day.
	MeanMotion param.Opt[float64] `json:"meanMotion,omitzero"`
	// 2nd derivative of the mean motion with respect to time. Units are revolutions
	// per day cubed.
	MeanMotionDDot param.Opt[float64] `json:"meanMotionDDot,omitzero"`
	// 1st derivative of the mean motion with respect to time. Units are revolutions
	// per day squared.
	MeanMotionDot param.Opt[float64] `json:"meanMotionDot,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The orbit point nearest to the center of the earth in kilometers.
	Perigee param.Opt[float64] `json:"perigee,omitzero"`
	// Period of the orbit equal to inverse of mean motion.
	Period param.Opt[float64] `json:"period,omitzero"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node, which is where the orbit crosses the
	// equator when traveling north.
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// The current revolution number. The value is incremented when a satellite crosses
	// the equator on an ascending pass.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The sum of the periapsis and apoapsis distances divided by two. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies,
	// not the distance of the bodies from the center of mass.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	paramObj
}

func (r ManifoldelsetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ManifoldelsetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManifoldelsetNewParams) UnmarshalJSON(data []byte) error {
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
type ManifoldelsetNewParamsDataMode string

const (
	ManifoldelsetNewParamsDataModeReal      ManifoldelsetNewParamsDataMode = "REAL"
	ManifoldelsetNewParamsDataModeTest      ManifoldelsetNewParamsDataMode = "TEST"
	ManifoldelsetNewParamsDataModeSimulated ManifoldelsetNewParamsDataMode = "SIMULATED"
	ManifoldelsetNewParamsDataModeExercise  ManifoldelsetNewParamsDataMode = "EXERCISE"
)

type ManifoldelsetUpdateParams struct {
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
	DataMode ManifoldelsetUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Identifier of the parent Manifold record.
	IDManifold string `json:"idManifold,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// A placeholder satellite number and not a true NORAD catalog number.
	TmpSatNo int64 `json:"tmpSatNo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Orbit point furthest from the center of the earth in kilometers.
	Apogee param.Opt[float64] `json:"apogee,omitzero"`
	// The argument of perigee is the angle in degrees formed between the perigee and
	// the ascending node. If the perigee would occur at the ascending node, the
	// argument of perigee would be 0.
	ArgOfPerigee param.Opt[float64] `json:"argOfPerigee,omitzero"`
	// The drag term for SGP4 orbital model, used for calculating decay constants for
	// altitude, eccentricity etc, measured in inverse earth radii.
	BStar param.Opt[float64] `json:"bStar,omitzero"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle. A value of 0 is a circular orbit, values between 0 and 1 form an
	// elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a
	// hyperbolic escape orbit.
	Eccentricity param.Opt[float64] `json:"eccentricity,omitzero"`
	// The angle between the equator and the orbit when looking from the center of the
	// Earth. If the orbit went exactly around the equator from left to right, then the
	// inclination would be 0. The inclination ranges from 0 to 180 degrees.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Where the satellite is in its orbital path. The mean anomaly ranges from 0 to
	// 360 degrees. The mean anomaly is referenced to the perigee. If the satellite
	// were at the perigee, the mean anomaly would be 0.
	MeanAnomaly param.Opt[float64] `json:"meanAnomaly,omitzero"`
	// The constant angular speed required for the body to complete one circular orbit
	// in the same amount of time as the actual elliptical orbit with variable speed.
	// Measured in revolutions per day.
	MeanMotion param.Opt[float64] `json:"meanMotion,omitzero"`
	// 2nd derivative of the mean motion with respect to time. Units are revolutions
	// per day cubed.
	MeanMotionDDot param.Opt[float64] `json:"meanMotionDDot,omitzero"`
	// 1st derivative of the mean motion with respect to time. Units are revolutions
	// per day squared.
	MeanMotionDot param.Opt[float64] `json:"meanMotionDot,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The orbit point nearest to the center of the earth in kilometers.
	Perigee param.Opt[float64] `json:"perigee,omitzero"`
	// Period of the orbit equal to inverse of mean motion.
	Period param.Opt[float64] `json:"period,omitzero"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node, which is where the orbit crosses the
	// equator when traveling north.
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// The current revolution number. The value is incremented when a satellite crosses
	// the equator on an ascending pass.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The sum of the periapsis and apoapsis distances divided by two. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies,
	// not the distance of the bodies from the center of mass.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	paramObj
}

func (r ManifoldelsetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ManifoldelsetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManifoldelsetUpdateParams) UnmarshalJSON(data []byte) error {
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
type ManifoldelsetUpdateParamsDataMode string

const (
	ManifoldelsetUpdateParamsDataModeReal      ManifoldelsetUpdateParamsDataMode = "REAL"
	ManifoldelsetUpdateParamsDataModeTest      ManifoldelsetUpdateParamsDataMode = "TEST"
	ManifoldelsetUpdateParamsDataModeSimulated ManifoldelsetUpdateParamsDataMode = "SIMULATED"
	ManifoldelsetUpdateParamsDataModeExercise  ManifoldelsetUpdateParamsDataMode = "EXERCISE"
)

type ManifoldelsetListParams struct {
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Epoch       time.Time        `query:"epoch,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManifoldelsetListParams]'s query parameters as
// `url.Values`.
func (r ManifoldelsetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManifoldelsetCountParams struct {
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Epoch       time.Time        `query:"epoch,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManifoldelsetCountParams]'s query parameters as
// `url.Values`.
func (r ManifoldelsetCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManifoldelsetNewBulkParams struct {
	Body []ManifoldelsetNewBulkParamsBody
	paramObj
}

func (r ManifoldelsetNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ManifoldelsetNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Theoretical Keplarian orbital elements belonging to an object of interest's
// manifold describing a possible/theoretical orbit for an object of interest for
// tasking purposes.
//
// The properties ClassificationMarking, DataMode, Epoch, IDManifold, Source,
// TmpSatNo are required.
type ManifoldelsetNewBulkParamsBody struct {
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
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Identifier of the parent Manifold record.
	IDManifold string `json:"idManifold,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// A placeholder satellite number and not a true NORAD catalog number.
	TmpSatNo int64 `json:"tmpSatNo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Orbit point furthest from the center of the earth in kilometers.
	Apogee param.Opt[float64] `json:"apogee,omitzero"`
	// The argument of perigee is the angle in degrees formed between the perigee and
	// the ascending node. If the perigee would occur at the ascending node, the
	// argument of perigee would be 0.
	ArgOfPerigee param.Opt[float64] `json:"argOfPerigee,omitzero"`
	// The drag term for SGP4 orbital model, used for calculating decay constants for
	// altitude, eccentricity etc, measured in inverse earth radii.
	BStar param.Opt[float64] `json:"bStar,omitzero"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle. A value of 0 is a circular orbit, values between 0 and 1 form an
	// elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a
	// hyperbolic escape orbit.
	Eccentricity param.Opt[float64] `json:"eccentricity,omitzero"`
	// The angle between the equator and the orbit when looking from the center of the
	// Earth. If the orbit went exactly around the equator from left to right, then the
	// inclination would be 0. The inclination ranges from 0 to 180 degrees.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Where the satellite is in its orbital path. The mean anomaly ranges from 0 to
	// 360 degrees. The mean anomaly is referenced to the perigee. If the satellite
	// were at the perigee, the mean anomaly would be 0.
	MeanAnomaly param.Opt[float64] `json:"meanAnomaly,omitzero"`
	// The constant angular speed required for the body to complete one circular orbit
	// in the same amount of time as the actual elliptical orbit with variable speed.
	// Measured in revolutions per day.
	MeanMotion param.Opt[float64] `json:"meanMotion,omitzero"`
	// 2nd derivative of the mean motion with respect to time. Units are revolutions
	// per day cubed.
	MeanMotionDDot param.Opt[float64] `json:"meanMotionDDot,omitzero"`
	// 1st derivative of the mean motion with respect to time. Units are revolutions
	// per day squared.
	MeanMotionDot param.Opt[float64] `json:"meanMotionDot,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The orbit point nearest to the center of the earth in kilometers.
	Perigee param.Opt[float64] `json:"perigee,omitzero"`
	// Period of the orbit equal to inverse of mean motion.
	Period param.Opt[float64] `json:"period,omitzero"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node, which is where the orbit crosses the
	// equator when traveling north.
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// The current revolution number. The value is incremented when a satellite crosses
	// the equator on an ascending pass.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The sum of the periapsis and apoapsis distances divided by two. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies,
	// not the distance of the bodies from the center of mass.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	paramObj
}

func (r ManifoldelsetNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ManifoldelsetNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManifoldelsetNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ManifoldelsetNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type ManifoldelsetGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManifoldelsetGetParams]'s query parameters as `url.Values`.
func (r ManifoldelsetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManifoldelsetTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Epoch       time.Time        `query:"epoch,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManifoldelsetTupleParams]'s query parameters as
// `url.Values`.
func (r ManifoldelsetTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
