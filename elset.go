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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// ElsetService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewElsetService] method instead.
type ElsetService struct {
	Options []option.RequestOption
	Current ElsetCurrentService
	History ElsetHistoryService
}

// NewElsetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewElsetService(opts ...option.RequestOption) (r ElsetService) {
	r = ElsetService{}
	r.Options = opts
	r.Current = NewElsetCurrentService(opts...)
	r.History = NewElsetHistoryService(opts...)
	return
}

// Service operation to take a single elset as a POST body and ingest into the
// database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *ElsetService) New(ctx context.Context, body ElsetNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/elset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single elset by its unique ID passed as a path
// parameter.
func (r *ElsetService) Get(ctx context.Context, id string, query ElsetGetParams, opts ...option.RequestOption) (res *Elset, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/elset/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ElsetService) List(ctx context.Context, query ElsetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ElsetAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/elset"
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
func (r *ElsetService) ListAutoPaging(ctx context.Context, query ElsetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ElsetAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ElsetService) Count(ctx context.Context, query ElsetCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/elset/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// elsets as a POST body and ingest into the database with or without dupe
// detection. Default is no dupe checking. This operation is not intended to be
// used for automated feeds into UDL. Data providers should contact the UDL team
// for specific role assignments and for instructions on setting up a permanent
// feed through an alternate mechanism.
func (r *ElsetService) NewBulk(ctx context.Context, params ElsetNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/elset/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Service operation to take a multiple TLEs as a POST body and ingest into the
// database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *ElsetService) NewBulkFromTle(ctx context.Context, params ElsetNewBulkFromTleParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/elset/createBulkFromTLE"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ElsetService) QueryCurrentElsetHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/currentelset/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ElsetService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/elset/queryhelp"
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
func (r *ElsetService) Tuple(ctx context.Context, query ElsetTupleParams, opts ...option.RequestOption) (res *[]Elset, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/elset/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take elsets as a POST body and ingest into the database
// with or without dupe detection. Default is no dupe checking. This operation is
// intended to be used for automated feeds into UDL.
func (r *ElsetService) UnvalidatedPublish(ctx context.Context, body ElsetUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-elset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
type Elset struct {
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
	DataMode ElsetDataMode `json:"dataMode,required"`
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// AGOM, expressed in m^2/kg, is the value of the (averaged) object Area times the
	// solar radiation pressure coefficient(Gamma) over the object Mass. Applicable
	// only with ephemType4.
	Agom float64 `json:"agom"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The orbit point furthest from the center of the earth in kilometers. If not
	// provided, apogee will be computed from the TLE according to the following. Using
	// mu, the standard gravitational parameter for the earth (398600.4418), semi-major
	// axis A = (mu/(n _ 2 _ pi/(24*3600))^2)(1/3). Using semi-major axis A,
	// eccentricity E, apogee = (A * (1 + E)) in km. Note that the calculations are for
	// computing the apogee radius from the center of the earth, to compute apogee
	// altitude the radius of the earth should be subtracted (6378.135 km).
	Apogee float64 `json:"apogee"`
	// The argument of perigee is the angle in degrees formed between the perigee and
	// the ascending node. If the perigee would occur at the ascending node, the
	// argument of perigee would be 0.
	ArgOfPerigee float64 `json:"argOfPerigee"`
	// Ballistic coefficient, in m^2/kg. Applicable only with ephemType4.
	BallisticCoeff float64 `json:"ballisticCoeff"`
	// The drag term for SGP4 orbital model, used for calculating decay constants for
	// altitude, eccentricity etc, measured in inverse earth radii.
	BStar float64 `json:"bStar"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle. A value of 0 is a circular orbit, values between 0 and 1 form an
	// elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a
	// hyperbolic escape orbit.
	Eccentricity float64 `json:"eccentricity"`
	// Read-only start time at which this elset was the 'current' elset for its
	// satellite. This field and is set by the system automatically and ignored on
	// create/edit operations.
	EffectiveFrom time.Time `json:"effectiveFrom" format:"date-time"`
	// Read-only end time at which this elset was no longer the 'current' elset for its
	// satellite. This field and is set by the system automatically and ignored on
	// create/edit operations.
	EffectiveUntil time.Time `json:"effectiveUntil" format:"date-time"`
	// The ephemeris type associated with this TLE:
	//
	// 0:&nbsp;SGP (or SGP4 with Kozai mean motion)
	//
	// 1:&nbsp;SGP (Kozai mean motion)
	//
	// 2:&nbsp;SGP4 (Brouver mean motion)
	//
	// 3:&nbsp;SDP4
	//
	// 4:&nbsp;SGP4-XP
	//
	// 5:&nbsp;SDP8
	//
	// 6:&nbsp;SP (osculating mean motion)
	EphemType int64 `json:"ephemType"`
	// Unique identifier of the record, auto-generated by the system.
	IDElset string `json:"idElset"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this elset. This ID
	// can be used to obtain additional information on an OrbitDetermination object
	// using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queried
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
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
	// Mean motion is the angular speed required for a body to complete one orbit,
	// assuming constant speed in a circular orbit which completes in the same time as
	// the variable speed, elliptical orbit of the actual body. Measured in revolutions
	// per day.
	MeanMotion float64 `json:"meanMotion"`
	// 2nd derivative of the mean motion with respect to time. Units are revolutions
	// per day cubed.
	MeanMotionDDot float64 `json:"meanMotionDDot"`
	// 1st derivative of the mean motion with respect to time. Units are revolutions
	// per day squared.
	MeanMotionDot float64 `json:"meanMotionDot"`
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
	// Optional identifier provided by elset source to indicate the target onorbit
	// object of this elset. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// The orbit point nearest to the center of the earth in kilometers. If not
	// provided, perigee will be computed from the TLE according to the following.
	// Using mu, the standard gravitational parameter for the earth (398600.4418),
	// semi-major axis A = (mu/(n _ 2 _ pi/(24*3600))^2)(1/3). Using semi-major axis A,
	// eccentricity E, perigee = (A * (1 - E)) in km. Note that the calculations are
	// for computing the perigee radius from the center of the earth, to compute
	// perigee altitude the radius of the earth should be subtracted (6378.135 km).
	Perigee float64 `json:"perigee"`
	// Period of the orbit equal to inverse of mean motion, in minutes.
	Period float64 `json:"period"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node, which is where the orbit crosses the
	// equator when traveling north.
	Raan float64 `json:"raan"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The current revolution number. The value is incremented when a satellite crosses
	// the equator on an ascending pass.
	RevNo int64 `json:"revNo"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// The sum of the periapsis and apoapsis distances divided by two. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies,
	// not the distance of the bodies from the center of mass. Units are kilometers.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// Optional array of UDL data (observation) UUIDs used to build this element set.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData"`
	// Optional array of UDL observation data types used to build this element set
	// (e.g. EO, RADAR, RF, DOA). See the associated sourcedData array for the specific
	// UUIDs of observations for the positionally corresponding data types in this
	// array (the two arrays must match in size).
	//
	// Any of "EO", "RADAR", "RF", "DOA", "ELSET", "SV".
	SourcedDataTypes []string `json:"sourcedDataTypes"`
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
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this Elset was unable to be correlated to a known object.
	// This flag should only be set to true by data providers after an attempt to
	// correlate to an on-orbit object was made and failed. If unable to correlate, the
	// 'origObjectId' field may be populated with an internal data provider specific
	// identifier.
	Uct bool `json:"uct"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Epoch                 resp.Field
		Source                resp.Field
		Agom                  resp.Field
		Algorithm             resp.Field
		Apogee                resp.Field
		ArgOfPerigee          resp.Field
		BallisticCoeff        resp.Field
		BStar                 resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		Eccentricity          resp.Field
		EffectiveFrom         resp.Field
		EffectiveUntil        resp.Field
		EphemType             resp.Field
		IDElset               resp.Field
		IDOnOrbit             resp.Field
		IDOrbitDetermination  resp.Field
		Inclination           resp.Field
		Line1                 resp.Field
		Line2                 resp.Field
		MeanAnomaly           resp.Field
		MeanMotion            resp.Field
		MeanMotionDDot        resp.Field
		MeanMotionDot         resp.Field
		OnOrbit               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		Perigee               resp.Field
		Period                resp.Field
		Raan                  resp.Field
		RawFileUri            resp.Field
		RevNo                 resp.Field
		SatNo                 resp.Field
		SemiMajorAxis         resp.Field
		SourcedData           resp.Field
		SourcedDataTypes      resp.Field
		SourceDl              resp.Field
		Tags                  resp.Field
		TransactionID         resp.Field
		Uct                   resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Elset) RawJSON() string { return r.JSON.raw }
func (r *Elset) UnmarshalJSON(data []byte) error {
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
type ElsetDataMode string

const (
	ElsetDataModeReal      ElsetDataMode = "REAL"
	ElsetDataModeTest      ElsetDataMode = "TEST"
	ElsetDataModeSimulated ElsetDataMode = "SIMULATED"
	ElsetDataModeExercise  ElsetDataMode = "EXERCISE"
)

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ElsetIngestParam struct {
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
	DataMode ElsetIngestDataMode `json:"dataMode,omitzero,required"`
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// AGOM, expressed in m^2/kg, is the value of the (averaged) object Area times the
	// solar radiation pressure coefficient(Gamma) over the object Mass. Applicable
	// only with ephemType4.
	Agom param.Opt[float64] `json:"agom,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// The orbit point furthest from the center of the earth in kilometers. If not
	// provided, apogee will be computed from the TLE according to the following. Using
	// mu, the standard gravitational parameter for the earth (398600.4418), semi-major
	// axis A = (mu/(n _ 2 _ pi/(24*3600))^2)(1/3). Using semi-major axis A,
	// eccentricity E, apogee = (A * (1 + E)) in km. Note that the calculations are for
	// computing the apogee radius from the center of the earth, to compute apogee
	// altitude the radius of the earth should be subtracted (6378.135 km).
	Apogee param.Opt[float64] `json:"apogee,omitzero"`
	// The argument of perigee is the angle in degrees formed between the perigee and
	// the ascending node. If the perigee would occur at the ascending node, the
	// argument of perigee would be 0.
	ArgOfPerigee param.Opt[float64] `json:"argOfPerigee,omitzero"`
	// Ballistic coefficient, in m^2/kg. Applicable only with ephemType4.
	BallisticCoeff param.Opt[float64] `json:"ballisticCoeff,omitzero"`
	// The drag term for SGP4 orbital model, used for calculating decay constants for
	// altitude, eccentricity etc, measured in inverse earth radii.
	BStar param.Opt[float64] `json:"bStar,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle. A value of 0 is a circular orbit, values between 0 and 1 form an
	// elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a
	// hyperbolic escape orbit.
	Eccentricity param.Opt[float64] `json:"eccentricity,omitzero"`
	// The ephemeris type associated with this TLE:
	//
	// 0:&nbsp;SGP (or SGP4 with Kozai mean motion)
	//
	// 1:&nbsp;SGP (Kozai mean motion)
	//
	// 2:&nbsp;SGP4 (Brouver mean motion)
	//
	// 3:&nbsp;SDP4
	//
	// 4:&nbsp;SGP4-XP
	//
	// 5:&nbsp;SDP8
	//
	// 6:&nbsp;SP (osculating mean motion)
	EphemType param.Opt[int64] `json:"ephemType,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDElset param.Opt[string] `json:"idElset,omitzero"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the OD solution record that produced this elset. This ID
	// can be used to obtain additional information on an OrbitDetermination object
	// using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queried
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination param.Opt[string] `json:"idOrbitDetermination,omitzero"`
	// The angle between the equator and the orbit when looking from the center of the
	// Earth. If the orbit went exactly around the equator from left to right, then the
	// inclination would be 0. The inclination ranges from 0 to 180 degrees.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Read only derived/generated line1 of a legacy TLE (two line element set) format,
	// ignored on create/edit operations.
	Line1 param.Opt[string] `json:"line1,omitzero"`
	// Read only derived/generated line2 of a legacy TLE (two line element set) format,
	// ignored on create/edit operations.
	Line2 param.Opt[string] `json:"line2,omitzero"`
	// Where the satellite is in its orbital path. The mean anomaly ranges from 0 to
	// 360 degrees. The mean anomaly is referenced to the perigee. If the satellite
	// were at the perigee, the mean anomaly would be 0.
	MeanAnomaly param.Opt[float64] `json:"meanAnomaly,omitzero"`
	// Mean motion is the angular speed required for a body to complete one orbit,
	// assuming constant speed in a circular orbit which completes in the same time as
	// the variable speed, elliptical orbit of the actual body. Measured in revolutions
	// per day.
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by elset source to indicate the target onorbit
	// object of this elset. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// The orbit point nearest to the center of the earth in kilometers. If not
	// provided, perigee will be computed from the TLE according to the following.
	// Using mu, the standard gravitational parameter for the earth (398600.4418),
	// semi-major axis A = (mu/(n _ 2 _ pi/(24*3600))^2)(1/3). Using semi-major axis A,
	// eccentricity E, perigee = (A * (1 - E)) in km. Note that the calculations are
	// for computing the perigee radius from the center of the earth, to compute
	// perigee altitude the radius of the earth should be subtracted (6378.135 km).
	Perigee param.Opt[float64] `json:"perigee,omitzero"`
	// Period of the orbit equal to inverse of mean motion, in minutes.
	Period param.Opt[float64] `json:"period,omitzero"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node, which is where the orbit crosses the
	// equator when traveling north.
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The current revolution number. The value is incremented when a satellite crosses
	// the equator on an ascending pass.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The sum of the periapsis and apoapsis distances divided by two. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies,
	// not the distance of the bodies from the center of mass. Units are kilometers.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this Elset was unable to be correlated to a known object.
	// This flag should only be set to true by data providers after an attempt to
	// correlate to an on-orbit object was made and failed. If unable to correlate, the
	// 'origObjectId' field may be populated with an internal data provider specific
	// identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Optional array of UDL data (observation) UUIDs used to build this element set.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL observation data types used to build this element set
	// (e.g. EO, RADAR, RF, DOA). See the associated sourcedData array for the specific
	// UUIDs of observations for the positionally corresponding data types in this
	// array (the two arrays must match in size).
	//
	// Any of "EO", "RADAR", "RF", "DOA", "ELSET", "SV".
	SourcedDataTypes []string `json:"sourcedDataTypes,omitzero"`
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
func (f ElsetIngestParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ElsetIngestParam) MarshalJSON() (data []byte, err error) {
	type shadow ElsetIngestParam
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
type ElsetIngestDataMode string

const (
	ElsetIngestDataModeReal      ElsetIngestDataMode = "REAL"
	ElsetIngestDataModeTest      ElsetIngestDataMode = "TEST"
	ElsetIngestDataModeSimulated ElsetIngestDataMode = "SIMULATED"
	ElsetIngestDataModeExercise  ElsetIngestDataMode = "EXERCISE"
)

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
type ElsetAbridged struct {
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
	DataMode ElsetAbridgedDataMode `json:"dataMode,required"`
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// AGOM, expressed in m^2/kg, is the value of the (averaged) object Area times the
	// solar radiation pressure coefficient(Gamma) over the object Mass. Applicable
	// only with ephemType4.
	Agom float64 `json:"agom"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The orbit point furthest from the center of the earth in kilometers. If not
	// provided, apogee will be computed from the TLE according to the following. Using
	// mu, the standard gravitational parameter for the earth (398600.4418), semi-major
	// axis A = (mu/(n _ 2 _ pi/(24*3600))^2)(1/3). Using semi-major axis A,
	// eccentricity E, apogee = (A * (1 + E)) in km. Note that the calculations are for
	// computing the apogee radius from the center of the earth, to compute apogee
	// altitude the radius of the earth should be subtracted (6378.135 km).
	Apogee float64 `json:"apogee"`
	// The argument of perigee is the angle in degrees formed between the perigee and
	// the ascending node. If the perigee would occur at the ascending node, the
	// argument of perigee would be 0.
	ArgOfPerigee float64 `json:"argOfPerigee"`
	// Ballistic coefficient, in m^2/kg. Applicable only with ephemType4.
	BallisticCoeff float64 `json:"ballisticCoeff"`
	// The drag term for SGP4 orbital model, used for calculating decay constants for
	// altitude, eccentricity etc, measured in inverse earth radii.
	BStar float64 `json:"bStar"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle. A value of 0 is a circular orbit, values between 0 and 1 form an
	// elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a
	// hyperbolic escape orbit.
	Eccentricity float64 `json:"eccentricity"`
	// The ephemeris type associated with this TLE:
	//
	// 0:&nbsp;SGP (or SGP4 with Kozai mean motion)
	//
	// 1:&nbsp;SGP (Kozai mean motion)
	//
	// 2:&nbsp;SGP4 (Brouver mean motion)
	//
	// 3:&nbsp;SDP4
	//
	// 4:&nbsp;SGP4-XP
	//
	// 5:&nbsp;SDP8
	//
	// 6:&nbsp;SP (osculating mean motion)
	EphemType int64 `json:"ephemType"`
	// Unique identifier of the record, auto-generated by the system.
	IDElset string `json:"idElset"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this elset. This ID
	// can be used to obtain additional information on an OrbitDetermination object
	// using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queried
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
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
	// Mean motion is the angular speed required for a body to complete one orbit,
	// assuming constant speed in a circular orbit which completes in the same time as
	// the variable speed, elliptical orbit of the actual body. Measured in revolutions
	// per day.
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
	// Optional identifier provided by elset source to indicate the target onorbit
	// object of this elset. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// The orbit point nearest to the center of the earth in kilometers. If not
	// provided, perigee will be computed from the TLE according to the following.
	// Using mu, the standard gravitational parameter for the earth (398600.4418),
	// semi-major axis A = (mu/(n _ 2 _ pi/(24*3600))^2)(1/3). Using semi-major axis A,
	// eccentricity E, perigee = (A * (1 - E)) in km. Note that the calculations are
	// for computing the perigee radius from the center of the earth, to compute
	// perigee altitude the radius of the earth should be subtracted (6378.135 km).
	Perigee float64 `json:"perigee"`
	// Period of the orbit equal to inverse of mean motion, in minutes.
	Period float64 `json:"period"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node, which is where the orbit crosses the
	// equator when traveling north.
	Raan float64 `json:"raan"`
	// The current revolution number. The value is incremented when a satellite crosses
	// the equator on an ascending pass.
	RevNo int64 `json:"revNo"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// The sum of the periapsis and apoapsis distances divided by two. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies,
	// not the distance of the bodies from the center of mass. Units are kilometers.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this Elset was unable to be correlated to a known object.
	// This flag should only be set to true by data providers after an attempt to
	// correlate to an on-orbit object was made and failed. If unable to correlate, the
	// 'origObjectId' field may be populated with an internal data provider specific
	// identifier.
	Uct bool `json:"uct"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Epoch                 resp.Field
		Source                resp.Field
		Agom                  resp.Field
		Algorithm             resp.Field
		Apogee                resp.Field
		ArgOfPerigee          resp.Field
		BallisticCoeff        resp.Field
		BStar                 resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		Eccentricity          resp.Field
		EphemType             resp.Field
		IDElset               resp.Field
		IDOnOrbit             resp.Field
		IDOrbitDetermination  resp.Field
		Inclination           resp.Field
		Line1                 resp.Field
		Line2                 resp.Field
		MeanAnomaly           resp.Field
		MeanMotion            resp.Field
		MeanMotionDDot        resp.Field
		MeanMotionDot         resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		Perigee               resp.Field
		Period                resp.Field
		Raan                  resp.Field
		RevNo                 resp.Field
		SatNo                 resp.Field
		SemiMajorAxis         resp.Field
		SourceDl              resp.Field
		TransactionID         resp.Field
		Uct                   resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElsetAbridged) RawJSON() string { return r.JSON.raw }
func (r *ElsetAbridged) UnmarshalJSON(data []byte) error {
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
type ElsetAbridgedDataMode string

const (
	ElsetAbridgedDataModeReal      ElsetAbridgedDataMode = "REAL"
	ElsetAbridgedDataModeTest      ElsetAbridgedDataMode = "TEST"
	ElsetAbridgedDataModeSimulated ElsetAbridgedDataMode = "SIMULATED"
	ElsetAbridgedDataModeExercise  ElsetAbridgedDataMode = "EXERCISE"
)

type ElsetNewParams struct {
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	ElsetIngest ElsetIngestParam
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ElsetNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ElsetNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.ElsetIngest)
}

type ElsetGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ElsetGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ElsetGetParams]'s query parameters as `url.Values`.
func (r ElsetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElsetListParams struct {
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Epoch       time.Time        `query:"epoch,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ElsetListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ElsetListParams]'s query parameters as `url.Values`.
func (r ElsetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElsetCountParams struct {
	// Elset epoch time in ISO 8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Epoch       time.Time        `query:"epoch,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ElsetCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ElsetCountParams]'s query parameters as `url.Values`.
func (r ElsetCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElsetNewBulkParams struct {
	Body []ElsetIngestParam
	// Boolean indicating if these elsets should be checked for duplicates, default is
	// not to.
	DupeCheck param.Opt[bool] `query:"dupeCheck,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ElsetNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ElsetNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// URLQuery serializes [ElsetNewBulkParams]'s query parameters as `url.Values`.
func (r ElsetNewBulkParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElsetNewBulkFromTleParams struct {
	// Data mode of the passed elsets (REAL, TEST, etc).
	DataMode string `query:"dataMode,required" json:"-"`
	// Boolean indicating if these elsets should be set as the 'current' for their
	// corresponding on-orbit/satellite numbers.
	MakeCurrent bool `query:"makeCurrent,required" json:"-"`
	// Source of the elset data.
	Source string `query:"source,required" json:"-"`
	Body   string
	// Boolean indicating if a shell Onorbit/satellite should be created if the passed
	// satellite number doesn't exist.
	AutoCreateSats param.Opt[bool] `query:"autoCreateSats,omitzero" json:"-"`
	// Dissemination control of the passed elsets (e.g. to support tagging with
	// proprietary markings).
	Control param.Opt[string] `query:"control,omitzero" json:"-"`
	// Origin of the elset data.
	Origin param.Opt[string] `query:"origin,omitzero" json:"-"`
	// Optional comma-delineated list of provider/source specific tags for this data,
	// where each element is no longer than 32 characters, used for implementing data
	// owner conditional access controls to restrict access to the data. Should be left
	// null by data providers unless conditional access controls are coordinated with
	// the UDL team.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ElsetNewBulkFromTleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ElsetNewBulkFromTleParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// URLQuery serializes [ElsetNewBulkFromTleParams]'s query parameters as
// `url.Values`.
func (r ElsetNewBulkFromTleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElsetTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ElsetTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ElsetTupleParams]'s query parameters as `url.Values`.
func (r ElsetTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElsetUnvalidatedPublishParams struct {
	Body []ElsetIngestParam
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ElsetUnvalidatedPublishParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ElsetUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
