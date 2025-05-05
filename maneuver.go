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

// ManeuverService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManeuverService] method instead.
type ManeuverService struct {
	Options []option.RequestOption
	History ManeuverHistoryService
}

// NewManeuverService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewManeuverService(opts ...option.RequestOption) (r ManeuverService) {
	r = ManeuverService{}
	r.Options = opts
	r.History = NewManeuverHistoryService(opts...)
	return
}

// Service operation to take a single maneuver as a POST body and ingest into the
// database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *ManeuverService) New(ctx context.Context, body ManeuverNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/maneuver"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ManeuverService) List(ctx context.Context, query ManeuverListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ManeuverListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/maneuver"
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
func (r *ManeuverService) ListAutoPaging(ctx context.Context, query ManeuverListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ManeuverListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ManeuverService) Count(ctx context.Context, query ManeuverCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/maneuver/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// maneuvers as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *ManeuverService) NewBulk(ctx context.Context, body ManeuverNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/maneuver/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single maneuver by its unique ID passed as a path
// parameter.
func (r *ManeuverService) Get(ctx context.Context, id string, query ManeuverGetParams, opts ...option.RequestOption) (res *ManeuverGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/maneuver/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ManeuverService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/maneuver/queryhelp"
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
func (r *ManeuverService) Tuple(ctx context.Context, query ManeuverTupleParams, opts ...option.RequestOption) (res *[]ManeuverTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/maneuver/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple maneuvers as a POST body and ingest into the
// database. This operation is intended to be used for automated feeds into UDL. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *ManeuverService) UnvalidatedPublish(ctx context.Context, body ManeuverUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-maneuver"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Model representation of on-orbit object maneuver information for detected,
// possible, and confirmed maneuvers.
type ManeuverListResponse struct {
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
	DataMode ManeuverListResponseDataMode `json:"dataMode,required"`
	// Maneuver event start time in ISO 8601 UTC with microsecond precision. For
	// maneuvers without start and end times, the start time is considered to be the
	// maneuver event time.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// Optional purpose of the maneuver if known (e.g. North-South Station Keeping,
	// East-West Station Keeping, Longitude Shift, Unknown).
	Characterization string `json:"characterization"`
	// Uncertainty in the characterization or purpose assessment of this maneuver (0 -
	// 1).
	CharacterizationUnc float64 `json:"characterizationUnc"`
	// Optional maneuver cross-track/radial/in-track covariance array, in meter and
	// second based units, in the following order: CR_R, CI_R, CI_I, CC_R, CC_I, CC_C,
	// CT_R, CT_I, CT_C, CT_T.
	Cov []float64 `json:"cov"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Difference in mass before and after the maneuver, in kg.
	DeltaMass float64 `json:"deltaMass"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors at the maneuver event time.
	DeltaPos float64 `json:"deltaPos"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaPosU float64 `json:"deltaPosU"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaPosV float64 `json:"deltaPosV"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaPosW float64 `json:"deltaPosW"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors at the maneuver event time.
	DeltaVel float64 `json:"deltaVel"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaVelU float64 `json:"deltaVelU"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaVelV float64 `json:"deltaVelV"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaVelW float64 `json:"deltaVelW"`
	// Description and notes of the maneuver.
	Description string `json:"description"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Maneuver event end time in ISO 8601 UTC with microsecond precision.
	EventEndTime time.Time `json:"eventEndTime" format:"date-time"`
	// Optional source-provided identifier for this maneuver event. In the case where
	// multiple maneuver records are submitted for the same event, this field can be
	// used to tie them together to the same event.
	EventID string `json:"eventId"`
	// Target maneuvering on-orbit object. For the public catalog, the idOnOrbit is
	// typically the satellite number as a string, but may be a UUID for analyst or
	// other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Optional ID of the sensor that detected this maneuver (for example, if detected
	// by passive RF anomalies).
	IDSensor string `json:"idSensor"`
	// Uncertainty in the occurrence of this maneuver (0 - 1).
	ManeuverUnc float64 `json:"maneuverUnc"`
	// Array of estimated acceleration values, in meters per second squared. Number of
	// elements must match the numAccelPoints.
	MnvrAccels []float64 `json:"mnvrAccels"`
	// Array of elapsed times, in seconds from maneuver start time, at which each
	// acceleration point is estimated. Number of elements must match the
	// numAccelPoints.
	MnvrAccelTimes []float64 `json:"mnvrAccelTimes"`
	// Array of the 1-sigma uncertainties in estimated accelerations, in meters per
	// second squared. Number of elements must match the numAccelPoints.
	MnvrAccelUncs []float64 `json:"mnvrAccelUncs"`
	// The total number of estimated acceleration points during the maneuver.
	NumAccelPoints int64 `json:"numAccelPoints"`
	// Number of observations used to generate the maneuver data.
	NumObs int64 `json:"numObs"`
	// Maneuver orbit determination fit data end time in ISO 8601 UTC with microsecond
	// precision.
	OdFitEndTime time.Time `json:"odFitEndTime" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Identifier provided by source to indicate the target on-orbit object performing
	// this maneuver. This may be an internal identifier and not necessarily a valid
	// satellite number/ID.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Post-event spacecraft apogee (measured from Earth center), in kilometers.
	PostApogee float64 `json:"postApogee"`
	// Estimated area of the object following the maneuver, in meters squared.
	PostArea float64 `json:"postArea"`
	// Post-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PostBallisticCoeff float64 `json:"postBallisticCoeff"`
	// Post-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PostDriftRate float64 `json:"postDriftRate"`
	// Post-event spacecraft eccentricity.
	PostEccentricity float64 `json:"postEccentricity"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PostEventElset ManeuverListResponsePostEventElset `json:"postEventElset"`
	// Optional identifier of the element set for the post-maneuver orbit.
	PostEventIDElset string `json:"postEventIdElset"`
	// Optional identifier of the state vector for the post-maneuver trajectory of the
	// spacecraft.
	PostEventIDStateVector string `json:"postEventIdStateVector"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PostEventStateVector ManeuverListResponsePostEventStateVector `json:"postEventStateVector"`
	// Post-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PostGeoLongitude float64 `json:"postGeoLongitude"`
	// Post-event spacecraft orbital inclination, in degrees. 0-180.
	PostInclination float64 `json:"postInclination"`
	// Estimated mass of the object following the maneuver, in kg.
	PostMass float64 `json:"postMass"`
	// Post-event spacecraft perigee (measured from Earth center), in kilometers.
	PostPerigee float64 `json:"postPerigee"`
	// Post-event spacecraft orbital period, in minutes.
	PostPeriod float64 `json:"postPeriod"`
	// Post-event X component of position in ECI space, in km.
	PostPosX float64 `json:"postPosX"`
	// Post-event Y component of position in ECI space, in km.
	PostPosY float64 `json:"postPosY"`
	// Post-event Z component of position in ECI space, in km.
	PostPosZ float64 `json:"postPosZ"`
	// Post-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PostRaan float64 `json:"postRAAN"`
	// Post-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PostRadiationPressCoeff float64 `json:"postRadiationPressCoeff"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PostSigmaU float64 `json:"postSigmaU"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PostSigmaV float64 `json:"postSigmaV"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PostSigmaW float64 `json:"postSigmaW"`
	// Post-event spacecraft Semi-Major Axis (SMA), in kilometers.
	PostSma float64 `json:"postSMA"`
	// Post-event X component of velocity in ECI space, in km/sec.
	PostVelX float64 `json:"postVelX"`
	// Post-event Y component of velocity in ECI space, in km/sec.
	PostVelY float64 `json:"postVelY"`
	// Post-event Z component of velocity in ECI space, in km/sec.
	PostVelZ float64 `json:"postVelZ"`
	// Pre-event spacecraft apogee (measured from Earth center), in kilometers.
	PreApogee float64 `json:"preApogee"`
	// Pre-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PreBallisticCoeff float64 `json:"preBallisticCoeff"`
	// Pre-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PreDriftRate float64 `json:"preDriftRate"`
	// Pre-event spacecraft eccentricity.
	PreEccentricity float64 `json:"preEccentricity"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PreEventElset ManeuverListResponsePreEventElset `json:"preEventElset"`
	// Optional identifier of the element set for the pre-maneuver orbit.
	PreEventIDElset string `json:"preEventIdElset"`
	// Optional identifier of the state vector for the pre-maneuver trajectory of the
	// spacecraft.
	PreEventIDStateVector string `json:"preEventIdStateVector"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PreEventStateVector ManeuverListResponsePreEventStateVector `json:"preEventStateVector"`
	// Pre-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PreGeoLongitude float64 `json:"preGeoLongitude"`
	// Pre-event spacecraft orbital inclination, in degrees. 0-180.
	PreInclination float64 `json:"preInclination"`
	// Pre-event spacecraft perigee (measured from Earth center), in kilometers.
	PrePerigee float64 `json:"prePerigee"`
	// Pre-event spacecraft orbital period, in minutes.
	PrePeriod float64 `json:"prePeriod"`
	// Pre-event X component of position in ECI space, in km.
	PrePosX float64 `json:"prePosX"`
	// Pre-event Y component of position in ECI space, in km.
	PrePosY float64 `json:"prePosY"`
	// Pre-event Z component of position in ECI space, in km.
	PrePosZ float64 `json:"prePosZ"`
	// Pre-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PreRaan float64 `json:"preRAAN"`
	// Pre-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PreRadiationPressCoeff float64 `json:"preRadiationPressCoeff"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PreSigmaU float64 `json:"preSigmaU"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PreSigmaV float64 `json:"preSigmaV"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PreSigmaW float64 `json:"preSigmaW"`
	// Pre-event spacecraft orbital Semi-Major Axis (SMA), in kilometers.
	PreSma float64 `json:"preSMA"`
	// Pre-event X component of velocity in ECI space, in km/sec.
	PreVelX float64 `json:"preVelX"`
	// Pre-event Y component of velocity in ECI space, in km/sec.
	PreVelY float64 `json:"preVelY"`
	// Pre-event Z component of velocity in ECI space, in km/sec.
	PreVelZ float64 `json:"preVelZ"`
	// The time that the report or alert of this maneuver was generated, in ISO 8601
	// UTC format.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Name of the state model used to generate the maneuver data.
	StateModel string `json:"stateModel"`
	// Version of the state model used to generate the maneuver data.
	StateModelVersion float64 `json:"stateModelVersion"`
	// Status of this maneuver (CANCELLED, PLANNED, POSSIBLE, REDACTED, VERIFIED).
	//
	// CANCELLED: A previously planned maneuver whose execution was cancelled.
	//
	// PLANNED: A maneuver planned to take place at the eventStartTime.
	//
	// POSSIBLE: A possible maneuver detected by observation of the spacecraft or by
	// evaluation of the spacecraft orbit.
	//
	// REDACTED: A redaction of a reported possible maneuver that has been determined
	// to have not taken place after further observation/evaluation.
	//
	// VERIFIED: A maneuver whose execution has been verified, either by the
	// owner/operator or observation/evaluation.
	Status string `json:"status"`
	// The estimated total active burn time of a maneuver, in seconds. This includes
	// the sum of all burns in numAccelPoints. Not to be confused with the total
	// duration of the maneuver.
	TotalBurnTime float64 `json:"totalBurnTime"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this maneuver was unable to be correlated to a known object.
	// This flag should only be set to true by data providers after an attempt to
	// correlate to an on-orbit object was made and failed. If unable to correlate, the
	// 'origObjectId' field may be populated with an internal data provider specific
	// identifier.
	Uct bool `json:"uct"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking   resp.Field
		DataMode                resp.Field
		EventStartTime          resp.Field
		Source                  resp.Field
		ID                      resp.Field
		Algorithm               resp.Field
		Characterization        resp.Field
		CharacterizationUnc     resp.Field
		Cov                     resp.Field
		CreatedAt               resp.Field
		CreatedBy               resp.Field
		DeltaMass               resp.Field
		DeltaPos                resp.Field
		DeltaPosU               resp.Field
		DeltaPosV               resp.Field
		DeltaPosW               resp.Field
		DeltaVel                resp.Field
		DeltaVelU               resp.Field
		DeltaVelV               resp.Field
		DeltaVelW               resp.Field
		Description             resp.Field
		Descriptor              resp.Field
		EventEndTime            resp.Field
		EventID                 resp.Field
		IDOnOrbit               resp.Field
		IDSensor                resp.Field
		ManeuverUnc             resp.Field
		MnvrAccels              resp.Field
		MnvrAccelTimes          resp.Field
		MnvrAccelUncs           resp.Field
		NumAccelPoints          resp.Field
		NumObs                  resp.Field
		OdFitEndTime            resp.Field
		Origin                  resp.Field
		OrigNetwork             resp.Field
		OrigObjectID            resp.Field
		OrigSensorID            resp.Field
		PostApogee              resp.Field
		PostArea                resp.Field
		PostBallisticCoeff      resp.Field
		PostDriftRate           resp.Field
		PostEccentricity        resp.Field
		PostEventElset          resp.Field
		PostEventIDElset        resp.Field
		PostEventIDStateVector  resp.Field
		PostEventStateVector    resp.Field
		PostGeoLongitude        resp.Field
		PostInclination         resp.Field
		PostMass                resp.Field
		PostPerigee             resp.Field
		PostPeriod              resp.Field
		PostPosX                resp.Field
		PostPosY                resp.Field
		PostPosZ                resp.Field
		PostRaan                resp.Field
		PostRadiationPressCoeff resp.Field
		PostSigmaU              resp.Field
		PostSigmaV              resp.Field
		PostSigmaW              resp.Field
		PostSma                 resp.Field
		PostVelX                resp.Field
		PostVelY                resp.Field
		PostVelZ                resp.Field
		PreApogee               resp.Field
		PreBallisticCoeff       resp.Field
		PreDriftRate            resp.Field
		PreEccentricity         resp.Field
		PreEventElset           resp.Field
		PreEventIDElset         resp.Field
		PreEventIDStateVector   resp.Field
		PreEventStateVector     resp.Field
		PreGeoLongitude         resp.Field
		PreInclination          resp.Field
		PrePerigee              resp.Field
		PrePeriod               resp.Field
		PrePosX                 resp.Field
		PrePosY                 resp.Field
		PrePosZ                 resp.Field
		PreRaan                 resp.Field
		PreRadiationPressCoeff  resp.Field
		PreSigmaU               resp.Field
		PreSigmaV               resp.Field
		PreSigmaW               resp.Field
		PreSma                  resp.Field
		PreVelX                 resp.Field
		PreVelY                 resp.Field
		PreVelZ                 resp.Field
		ReportTime              resp.Field
		SatNo                   resp.Field
		StateModel              resp.Field
		StateModelVersion       resp.Field
		Status                  resp.Field
		TotalBurnTime           resp.Field
		TransactionID           resp.Field
		Uct                     resp.Field
		ExtraFields             map[string]resp.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManeuverListResponse) RawJSON() string { return r.JSON.raw }
func (r *ManeuverListResponse) UnmarshalJSON(data []byte) error {
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
type ManeuverListResponseDataMode string

const (
	ManeuverListResponseDataModeReal      ManeuverListResponseDataMode = "REAL"
	ManeuverListResponseDataModeTest      ManeuverListResponseDataMode = "TEST"
	ManeuverListResponseDataModeSimulated ManeuverListResponseDataMode = "SIMULATED"
	ManeuverListResponseDataModeExercise  ManeuverListResponseDataMode = "EXERCISE"
)

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
type ManeuverListResponsePostEventElset struct {
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
func (r ManeuverListResponsePostEventElset) RawJSON() string { return r.JSON.raw }
func (r *ManeuverListResponsePostEventElset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
type ManeuverListResponsePostEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan float64 `json:"actualODSpan"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame string `json:"alt1ReferenceFrame"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame string `json:"alt2ReferenceFrame"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area float64 `json:"area"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot float64 `json:"bDot"`
	// Model parameter value for center of mass offset (m).
	CmOffset float64 `json:"cmOffset"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod string `json:"covMethod"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea float64 `json:"dragArea"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff float64 `json:"dragCoeff"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel string `json:"dragModel"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr float64 `json:"edr"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov"`
	// Integrator error control.
	ErrorControl float64 `json:"errorControl"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep bool `json:"fixedStep"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel string `json:"geopotentialModel"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms int64 `json:"iau1980Terms"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector string `json:"idStateVector"`
	// Integrator Mode.
	IntegratorMode string `json:"integratorMode"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust bool `json:"inTrackThrust"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd time.Time `json:"lastObEnd" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart time.Time `json:"lastObStart" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime time.Time `json:"leapSecondTime" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar bool `json:"lunarSolar"`
	// The mass of the object, in kilograms.
	Mass float64 `json:"mass"`
	// The number of observations available for the OD of the object.
	ObsAvailable int64 `json:"obsAvailable"`
	// The number of observations accepted for the OD of the object.
	ObsUsed int64 `json:"obsUsed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials string `json:"partials"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree string `json:"pedigree"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX float64 `json:"polarMotionX"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY float64 `json:"polarMotionY"`
	// One sigma position uncertainty, in kilometers.
	PosUnc float64 `json:"posUnc"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan float64 `json:"recODSpan"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc float64 `json:"residualsAcc"`
	// Epoch revolution number.
	RevNo int64 `json:"revNo"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms float64 `json:"rms"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo int64 `json:"satNo"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg float64 `json:"solarFluxAPAvg"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 float64 `json:"solarFluxF10"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg float64 `json:"solarFluxF10Avg"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress bool `json:"solarRadPress"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff float64 `json:"solarRadPressCoeff"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides bool `json:"solidEarthTides"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea float64 `json:"srpArea"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode string `json:"stepMode"`
	// Initial integration step size (seconds).
	StepSize float64 `json:"stepSize"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection string `json:"stepSizeSelection"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc float64 `json:"taiUtc"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel float64 `json:"thrustAccel"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail int64 `json:"tracksAvail"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed int64 `json:"tracksUsed"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct bool `json:"uct"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate float64 `json:"ut1Rate"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc float64 `json:"ut1Utc"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc float64 `json:"velUnc"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel float64 `json:"xaccel"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 float64 `json:"xposAlt1"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 float64 `json:"xposAlt2"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 float64 `json:"xvelAlt1"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 float64 `json:"xvelAlt2"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel float64 `json:"yaccel"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 float64 `json:"yposAlt1"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 float64 `json:"yposAlt2"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 float64 `json:"yvelAlt1"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 float64 `json:"yvelAlt2"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel float64 `json:"zaccel"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 float64 `json:"zposAlt1"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 float64 `json:"zposAlt2"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 float64 `json:"zvelAlt1"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 float64 `json:"zvelAlt2"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Epoch                 resp.Field
		Source                resp.Field
		ActualOdSpan          resp.Field
		Algorithm             resp.Field
		Alt1ReferenceFrame    resp.Field
		Alt2ReferenceFrame    resp.Field
		Area                  resp.Field
		BDot                  resp.Field
		CmOffset              resp.Field
		Cov                   resp.Field
		CovMethod             resp.Field
		CovReferenceFrame     resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		DragArea              resp.Field
		DragCoeff             resp.Field
		DragModel             resp.Field
		Edr                   resp.Field
		EqCov                 resp.Field
		ErrorControl          resp.Field
		FixedStep             resp.Field
		GeopotentialModel     resp.Field
		Iau1980Terms          resp.Field
		IDOnOrbit             resp.Field
		IDOrbitDetermination  resp.Field
		IDStateVector         resp.Field
		IntegratorMode        resp.Field
		InTrackThrust         resp.Field
		LastObEnd             resp.Field
		LastObStart           resp.Field
		LeapSecondTime        resp.Field
		LunarSolar            resp.Field
		Mass                  resp.Field
		ObsAvailable          resp.Field
		ObsUsed               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		Partials              resp.Field
		Pedigree              resp.Field
		PolarMotionX          resp.Field
		PolarMotionY          resp.Field
		PosUnc                resp.Field
		RecOdSpan             resp.Field
		ReferenceFrame        resp.Field
		ResidualsAcc          resp.Field
		RevNo                 resp.Field
		Rms                   resp.Field
		SatNo                 resp.Field
		SigmaPosUvw           resp.Field
		SigmaVelUvw           resp.Field
		SolarFluxApAvg        resp.Field
		SolarFluxF10          resp.Field
		SolarFluxF10Avg       resp.Field
		SolarRadPress         resp.Field
		SolarRadPressCoeff    resp.Field
		SolidEarthTides       resp.Field
		SourceDl              resp.Field
		SrpArea               resp.Field
		StepMode              resp.Field
		StepSize              resp.Field
		StepSizeSelection     resp.Field
		TaiUtc                resp.Field
		ThrustAccel           resp.Field
		TracksAvail           resp.Field
		TracksUsed            resp.Field
		TransactionID         resp.Field
		Uct                   resp.Field
		Ut1Rate               resp.Field
		Ut1Utc                resp.Field
		VelUnc                resp.Field
		Xaccel                resp.Field
		Xpos                  resp.Field
		XposAlt1              resp.Field
		XposAlt2              resp.Field
		Xvel                  resp.Field
		XvelAlt1              resp.Field
		XvelAlt2              resp.Field
		Yaccel                resp.Field
		Ypos                  resp.Field
		YposAlt1              resp.Field
		YposAlt2              resp.Field
		Yvel                  resp.Field
		YvelAlt1              resp.Field
		YvelAlt2              resp.Field
		Zaccel                resp.Field
		Zpos                  resp.Field
		ZposAlt1              resp.Field
		ZposAlt2              resp.Field
		Zvel                  resp.Field
		ZvelAlt1              resp.Field
		ZvelAlt2              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManeuverListResponsePostEventStateVector) RawJSON() string { return r.JSON.raw }
func (r *ManeuverListResponsePostEventStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
type ManeuverListResponsePreEventElset struct {
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
func (r ManeuverListResponsePreEventElset) RawJSON() string { return r.JSON.raw }
func (r *ManeuverListResponsePreEventElset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
type ManeuverListResponsePreEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan float64 `json:"actualODSpan"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame string `json:"alt1ReferenceFrame"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame string `json:"alt2ReferenceFrame"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area float64 `json:"area"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot float64 `json:"bDot"`
	// Model parameter value for center of mass offset (m).
	CmOffset float64 `json:"cmOffset"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod string `json:"covMethod"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea float64 `json:"dragArea"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff float64 `json:"dragCoeff"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel string `json:"dragModel"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr float64 `json:"edr"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov"`
	// Integrator error control.
	ErrorControl float64 `json:"errorControl"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep bool `json:"fixedStep"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel string `json:"geopotentialModel"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms int64 `json:"iau1980Terms"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector string `json:"idStateVector"`
	// Integrator Mode.
	IntegratorMode string `json:"integratorMode"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust bool `json:"inTrackThrust"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd time.Time `json:"lastObEnd" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart time.Time `json:"lastObStart" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime time.Time `json:"leapSecondTime" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar bool `json:"lunarSolar"`
	// The mass of the object, in kilograms.
	Mass float64 `json:"mass"`
	// The number of observations available for the OD of the object.
	ObsAvailable int64 `json:"obsAvailable"`
	// The number of observations accepted for the OD of the object.
	ObsUsed int64 `json:"obsUsed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials string `json:"partials"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree string `json:"pedigree"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX float64 `json:"polarMotionX"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY float64 `json:"polarMotionY"`
	// One sigma position uncertainty, in kilometers.
	PosUnc float64 `json:"posUnc"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan float64 `json:"recODSpan"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc float64 `json:"residualsAcc"`
	// Epoch revolution number.
	RevNo int64 `json:"revNo"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms float64 `json:"rms"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo int64 `json:"satNo"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg float64 `json:"solarFluxAPAvg"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 float64 `json:"solarFluxF10"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg float64 `json:"solarFluxF10Avg"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress bool `json:"solarRadPress"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff float64 `json:"solarRadPressCoeff"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides bool `json:"solidEarthTides"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea float64 `json:"srpArea"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode string `json:"stepMode"`
	// Initial integration step size (seconds).
	StepSize float64 `json:"stepSize"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection string `json:"stepSizeSelection"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc float64 `json:"taiUtc"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel float64 `json:"thrustAccel"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail int64 `json:"tracksAvail"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed int64 `json:"tracksUsed"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct bool `json:"uct"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate float64 `json:"ut1Rate"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc float64 `json:"ut1Utc"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc float64 `json:"velUnc"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel float64 `json:"xaccel"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 float64 `json:"xposAlt1"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 float64 `json:"xposAlt2"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 float64 `json:"xvelAlt1"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 float64 `json:"xvelAlt2"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel float64 `json:"yaccel"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 float64 `json:"yposAlt1"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 float64 `json:"yposAlt2"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 float64 `json:"yvelAlt1"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 float64 `json:"yvelAlt2"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel float64 `json:"zaccel"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 float64 `json:"zposAlt1"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 float64 `json:"zposAlt2"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 float64 `json:"zvelAlt1"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 float64 `json:"zvelAlt2"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Epoch                 resp.Field
		Source                resp.Field
		ActualOdSpan          resp.Field
		Algorithm             resp.Field
		Alt1ReferenceFrame    resp.Field
		Alt2ReferenceFrame    resp.Field
		Area                  resp.Field
		BDot                  resp.Field
		CmOffset              resp.Field
		Cov                   resp.Field
		CovMethod             resp.Field
		CovReferenceFrame     resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		DragArea              resp.Field
		DragCoeff             resp.Field
		DragModel             resp.Field
		Edr                   resp.Field
		EqCov                 resp.Field
		ErrorControl          resp.Field
		FixedStep             resp.Field
		GeopotentialModel     resp.Field
		Iau1980Terms          resp.Field
		IDOnOrbit             resp.Field
		IDOrbitDetermination  resp.Field
		IDStateVector         resp.Field
		IntegratorMode        resp.Field
		InTrackThrust         resp.Field
		LastObEnd             resp.Field
		LastObStart           resp.Field
		LeapSecondTime        resp.Field
		LunarSolar            resp.Field
		Mass                  resp.Field
		ObsAvailable          resp.Field
		ObsUsed               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		Partials              resp.Field
		Pedigree              resp.Field
		PolarMotionX          resp.Field
		PolarMotionY          resp.Field
		PosUnc                resp.Field
		RecOdSpan             resp.Field
		ReferenceFrame        resp.Field
		ResidualsAcc          resp.Field
		RevNo                 resp.Field
		Rms                   resp.Field
		SatNo                 resp.Field
		SigmaPosUvw           resp.Field
		SigmaVelUvw           resp.Field
		SolarFluxApAvg        resp.Field
		SolarFluxF10          resp.Field
		SolarFluxF10Avg       resp.Field
		SolarRadPress         resp.Field
		SolarRadPressCoeff    resp.Field
		SolidEarthTides       resp.Field
		SourceDl              resp.Field
		SrpArea               resp.Field
		StepMode              resp.Field
		StepSize              resp.Field
		StepSizeSelection     resp.Field
		TaiUtc                resp.Field
		ThrustAccel           resp.Field
		TracksAvail           resp.Field
		TracksUsed            resp.Field
		TransactionID         resp.Field
		Uct                   resp.Field
		Ut1Rate               resp.Field
		Ut1Utc                resp.Field
		VelUnc                resp.Field
		Xaccel                resp.Field
		Xpos                  resp.Field
		XposAlt1              resp.Field
		XposAlt2              resp.Field
		Xvel                  resp.Field
		XvelAlt1              resp.Field
		XvelAlt2              resp.Field
		Yaccel                resp.Field
		Ypos                  resp.Field
		YposAlt1              resp.Field
		YposAlt2              resp.Field
		Yvel                  resp.Field
		YvelAlt1              resp.Field
		YvelAlt2              resp.Field
		Zaccel                resp.Field
		Zpos                  resp.Field
		ZposAlt1              resp.Field
		ZposAlt2              resp.Field
		Zvel                  resp.Field
		ZvelAlt1              resp.Field
		ZvelAlt2              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManeuverListResponsePreEventStateVector) RawJSON() string { return r.JSON.raw }
func (r *ManeuverListResponsePreEventStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of on-orbit object maneuver information for detected,
// possible, and confirmed maneuvers.
type ManeuverGetResponse struct {
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
	DataMode ManeuverGetResponseDataMode `json:"dataMode,required"`
	// Maneuver event start time in ISO 8601 UTC with microsecond precision. For
	// maneuvers without start and end times, the start time is considered to be the
	// maneuver event time.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// Optional purpose of the maneuver if known (e.g. North-South Station Keeping,
	// East-West Station Keeping, Longitude Shift, Unknown).
	Characterization string `json:"characterization"`
	// Uncertainty in the characterization or purpose assessment of this maneuver (0 -
	// 1).
	CharacterizationUnc float64 `json:"characterizationUnc"`
	// Optional maneuver cross-track/radial/in-track covariance array, in meter and
	// second based units, in the following order: CR_R, CI_R, CI_I, CC_R, CC_I, CC_C,
	// CT_R, CT_I, CT_C, CT_T.
	Cov []float64 `json:"cov"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Difference in mass before and after the maneuver, in kg.
	DeltaMass float64 `json:"deltaMass"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors at the maneuver event time.
	DeltaPos float64 `json:"deltaPos"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaPosU float64 `json:"deltaPosU"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaPosV float64 `json:"deltaPosV"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaPosW float64 `json:"deltaPosW"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors at the maneuver event time.
	DeltaVel float64 `json:"deltaVel"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaVelU float64 `json:"deltaVelU"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaVelV float64 `json:"deltaVelV"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaVelW float64 `json:"deltaVelW"`
	// Description and notes of the maneuver.
	Description string `json:"description"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Maneuver event end time in ISO 8601 UTC with microsecond precision.
	EventEndTime time.Time `json:"eventEndTime" format:"date-time"`
	// Optional source-provided identifier for this maneuver event. In the case where
	// multiple maneuver records are submitted for the same event, this field can be
	// used to tie them together to the same event.
	EventID string `json:"eventId"`
	// Target maneuvering on-orbit object. For the public catalog, the idOnOrbit is
	// typically the satellite number as a string, but may be a UUID for analyst or
	// other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Optional ID of the sensor that detected this maneuver (for example, if detected
	// by passive RF anomalies).
	IDSensor string `json:"idSensor"`
	// Uncertainty in the occurrence of this maneuver (0 - 1).
	ManeuverUnc float64 `json:"maneuverUnc"`
	// Array of estimated acceleration values, in meters per second squared. Number of
	// elements must match the numAccelPoints.
	MnvrAccels []float64 `json:"mnvrAccels"`
	// Array of elapsed times, in seconds from maneuver start time, at which each
	// acceleration point is estimated. Number of elements must match the
	// numAccelPoints.
	MnvrAccelTimes []float64 `json:"mnvrAccelTimes"`
	// Array of the 1-sigma uncertainties in estimated accelerations, in meters per
	// second squared. Number of elements must match the numAccelPoints.
	MnvrAccelUncs []float64 `json:"mnvrAccelUncs"`
	// The total number of estimated acceleration points during the maneuver.
	NumAccelPoints int64 `json:"numAccelPoints"`
	// Number of observations used to generate the maneuver data.
	NumObs int64 `json:"numObs"`
	// Maneuver orbit determination fit data end time in ISO 8601 UTC with microsecond
	// precision.
	OdFitEndTime time.Time `json:"odFitEndTime" format:"date-time"`
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
	// Identifier provided by source to indicate the target on-orbit object performing
	// this maneuver. This may be an internal identifier and not necessarily a valid
	// satellite number/ID.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Post-event spacecraft apogee (measured from Earth center), in kilometers.
	PostApogee float64 `json:"postApogee"`
	// Estimated area of the object following the maneuver, in meters squared.
	PostArea float64 `json:"postArea"`
	// Post-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PostBallisticCoeff float64 `json:"postBallisticCoeff"`
	// Post-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PostDriftRate float64 `json:"postDriftRate"`
	// Post-event spacecraft eccentricity.
	PostEccentricity float64 `json:"postEccentricity"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PostEventElset ManeuverGetResponsePostEventElset `json:"postEventElset"`
	// Optional identifier of the element set for the post-maneuver orbit.
	PostEventIDElset string `json:"postEventIdElset"`
	// Optional identifier of the state vector for the post-maneuver trajectory of the
	// spacecraft.
	PostEventIDStateVector string `json:"postEventIdStateVector"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PostEventStateVector ManeuverGetResponsePostEventStateVector `json:"postEventStateVector"`
	// Post-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PostGeoLongitude float64 `json:"postGeoLongitude"`
	// Post-event spacecraft orbital inclination, in degrees. 0-180.
	PostInclination float64 `json:"postInclination"`
	// Estimated mass of the object following the maneuver, in kg.
	PostMass float64 `json:"postMass"`
	// Post-event spacecraft perigee (measured from Earth center), in kilometers.
	PostPerigee float64 `json:"postPerigee"`
	// Post-event spacecraft orbital period, in minutes.
	PostPeriod float64 `json:"postPeriod"`
	// Post-event X component of position in ECI space, in km.
	PostPosX float64 `json:"postPosX"`
	// Post-event Y component of position in ECI space, in km.
	PostPosY float64 `json:"postPosY"`
	// Post-event Z component of position in ECI space, in km.
	PostPosZ float64 `json:"postPosZ"`
	// Post-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PostRaan float64 `json:"postRAAN"`
	// Post-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PostRadiationPressCoeff float64 `json:"postRadiationPressCoeff"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PostSigmaU float64 `json:"postSigmaU"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PostSigmaV float64 `json:"postSigmaV"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PostSigmaW float64 `json:"postSigmaW"`
	// Post-event spacecraft Semi-Major Axis (SMA), in kilometers.
	PostSma float64 `json:"postSMA"`
	// Post-event X component of velocity in ECI space, in km/sec.
	PostVelX float64 `json:"postVelX"`
	// Post-event Y component of velocity in ECI space, in km/sec.
	PostVelY float64 `json:"postVelY"`
	// Post-event Z component of velocity in ECI space, in km/sec.
	PostVelZ float64 `json:"postVelZ"`
	// Pre-event spacecraft apogee (measured from Earth center), in kilometers.
	PreApogee float64 `json:"preApogee"`
	// Pre-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PreBallisticCoeff float64 `json:"preBallisticCoeff"`
	// Pre-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PreDriftRate float64 `json:"preDriftRate"`
	// Pre-event spacecraft eccentricity.
	PreEccentricity float64 `json:"preEccentricity"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PreEventElset ManeuverGetResponsePreEventElset `json:"preEventElset"`
	// Optional identifier of the element set for the pre-maneuver orbit.
	PreEventIDElset string `json:"preEventIdElset"`
	// Optional identifier of the state vector for the pre-maneuver trajectory of the
	// spacecraft.
	PreEventIDStateVector string `json:"preEventIdStateVector"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PreEventStateVector ManeuverGetResponsePreEventStateVector `json:"preEventStateVector"`
	// Pre-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PreGeoLongitude float64 `json:"preGeoLongitude"`
	// Pre-event spacecraft orbital inclination, in degrees. 0-180.
	PreInclination float64 `json:"preInclination"`
	// Pre-event spacecraft perigee (measured from Earth center), in kilometers.
	PrePerigee float64 `json:"prePerigee"`
	// Pre-event spacecraft orbital period, in minutes.
	PrePeriod float64 `json:"prePeriod"`
	// Pre-event X component of position in ECI space, in km.
	PrePosX float64 `json:"prePosX"`
	// Pre-event Y component of position in ECI space, in km.
	PrePosY float64 `json:"prePosY"`
	// Pre-event Z component of position in ECI space, in km.
	PrePosZ float64 `json:"prePosZ"`
	// Pre-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PreRaan float64 `json:"preRAAN"`
	// Pre-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PreRadiationPressCoeff float64 `json:"preRadiationPressCoeff"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PreSigmaU float64 `json:"preSigmaU"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PreSigmaV float64 `json:"preSigmaV"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PreSigmaW float64 `json:"preSigmaW"`
	// Pre-event spacecraft orbital Semi-Major Axis (SMA), in kilometers.
	PreSma float64 `json:"preSMA"`
	// Pre-event X component of velocity in ECI space, in km/sec.
	PreVelX float64 `json:"preVelX"`
	// Pre-event Y component of velocity in ECI space, in km/sec.
	PreVelY float64 `json:"preVelY"`
	// Pre-event Z component of velocity in ECI space, in km/sec.
	PreVelZ float64 `json:"preVelZ"`
	// The time that the report or alert of this maneuver was generated, in ISO 8601
	// UTC format.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Optional array of UDL data (elsets, state vectors, etc) UUIDs used to build this
	// maneuver. See the associated sourcedDataTypes array for the specific types of
	// data for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData"`
	// Optional array of UDL data types used to build this maneuver (e.g. EO, RADAR,
	// RF, DOA, ELSET, SV). See the associated sourcedData array for the specific UUIDs
	// of data for the positionally corresponding data types in this array (the two
	// arrays must match in size).
	//
	// Any of "EO", "RADAR", "RF", "DOA", "ELSET", "SV".
	SourcedDataTypes []string `json:"sourcedDataTypes"`
	// Name of the state model used to generate the maneuver data.
	StateModel string `json:"stateModel"`
	// Version of the state model used to generate the maneuver data.
	StateModelVersion float64 `json:"stateModelVersion"`
	// Status of this maneuver (CANCELLED, PLANNED, POSSIBLE, REDACTED, VERIFIED).
	//
	// CANCELLED: A previously planned maneuver whose execution was cancelled.
	//
	// PLANNED: A maneuver planned to take place at the eventStartTime.
	//
	// POSSIBLE: A possible maneuver detected by observation of the spacecraft or by
	// evaluation of the spacecraft orbit.
	//
	// REDACTED: A redaction of a reported possible maneuver that has been determined
	// to have not taken place after further observation/evaluation.
	//
	// VERIFIED: A maneuver whose execution has been verified, either by the
	// owner/operator or observation/evaluation.
	Status string `json:"status"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// The estimated total active burn time of a maneuver, in seconds. This includes
	// the sum of all burns in numAccelPoints. Not to be confused with the total
	// duration of the maneuver.
	TotalBurnTime float64 `json:"totalBurnTime"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this maneuver was unable to be correlated to a known object.
	// This flag should only be set to true by data providers after an attempt to
	// correlate to an on-orbit object was made and failed. If unable to correlate, the
	// 'origObjectId' field may be populated with an internal data provider specific
	// identifier.
	Uct bool `json:"uct"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking   resp.Field
		DataMode                resp.Field
		EventStartTime          resp.Field
		Source                  resp.Field
		ID                      resp.Field
		Algorithm               resp.Field
		Characterization        resp.Field
		CharacterizationUnc     resp.Field
		Cov                     resp.Field
		CreatedAt               resp.Field
		CreatedBy               resp.Field
		DeltaMass               resp.Field
		DeltaPos                resp.Field
		DeltaPosU               resp.Field
		DeltaPosV               resp.Field
		DeltaPosW               resp.Field
		DeltaVel                resp.Field
		DeltaVelU               resp.Field
		DeltaVelV               resp.Field
		DeltaVelW               resp.Field
		Description             resp.Field
		Descriptor              resp.Field
		EventEndTime            resp.Field
		EventID                 resp.Field
		IDOnOrbit               resp.Field
		IDSensor                resp.Field
		ManeuverUnc             resp.Field
		MnvrAccels              resp.Field
		MnvrAccelTimes          resp.Field
		MnvrAccelUncs           resp.Field
		NumAccelPoints          resp.Field
		NumObs                  resp.Field
		OdFitEndTime            resp.Field
		OnOrbit                 resp.Field
		Origin                  resp.Field
		OrigNetwork             resp.Field
		OrigObjectID            resp.Field
		OrigSensorID            resp.Field
		PostApogee              resp.Field
		PostArea                resp.Field
		PostBallisticCoeff      resp.Field
		PostDriftRate           resp.Field
		PostEccentricity        resp.Field
		PostEventElset          resp.Field
		PostEventIDElset        resp.Field
		PostEventIDStateVector  resp.Field
		PostEventStateVector    resp.Field
		PostGeoLongitude        resp.Field
		PostInclination         resp.Field
		PostMass                resp.Field
		PostPerigee             resp.Field
		PostPeriod              resp.Field
		PostPosX                resp.Field
		PostPosY                resp.Field
		PostPosZ                resp.Field
		PostRaan                resp.Field
		PostRadiationPressCoeff resp.Field
		PostSigmaU              resp.Field
		PostSigmaV              resp.Field
		PostSigmaW              resp.Field
		PostSma                 resp.Field
		PostVelX                resp.Field
		PostVelY                resp.Field
		PostVelZ                resp.Field
		PreApogee               resp.Field
		PreBallisticCoeff       resp.Field
		PreDriftRate            resp.Field
		PreEccentricity         resp.Field
		PreEventElset           resp.Field
		PreEventIDElset         resp.Field
		PreEventIDStateVector   resp.Field
		PreEventStateVector     resp.Field
		PreGeoLongitude         resp.Field
		PreInclination          resp.Field
		PrePerigee              resp.Field
		PrePeriod               resp.Field
		PrePosX                 resp.Field
		PrePosY                 resp.Field
		PrePosZ                 resp.Field
		PreRaan                 resp.Field
		PreRadiationPressCoeff  resp.Field
		PreSigmaU               resp.Field
		PreSigmaV               resp.Field
		PreSigmaW               resp.Field
		PreSma                  resp.Field
		PreVelX                 resp.Field
		PreVelY                 resp.Field
		PreVelZ                 resp.Field
		ReportTime              resp.Field
		SatNo                   resp.Field
		SourcedData             resp.Field
		SourcedDataTypes        resp.Field
		StateModel              resp.Field
		StateModelVersion       resp.Field
		Status                  resp.Field
		Tags                    resp.Field
		TotalBurnTime           resp.Field
		TransactionID           resp.Field
		Uct                     resp.Field
		ExtraFields             map[string]resp.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManeuverGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ManeuverGetResponse) UnmarshalJSON(data []byte) error {
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
type ManeuverGetResponseDataMode string

const (
	ManeuverGetResponseDataModeReal      ManeuverGetResponseDataMode = "REAL"
	ManeuverGetResponseDataModeTest      ManeuverGetResponseDataMode = "TEST"
	ManeuverGetResponseDataModeSimulated ManeuverGetResponseDataMode = "SIMULATED"
	ManeuverGetResponseDataModeExercise  ManeuverGetResponseDataMode = "EXERCISE"
)

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
type ManeuverGetResponsePostEventElset struct {
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
func (r ManeuverGetResponsePostEventElset) RawJSON() string { return r.JSON.raw }
func (r *ManeuverGetResponsePostEventElset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
type ManeuverGetResponsePostEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan float64 `json:"actualODSpan"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame string `json:"alt1ReferenceFrame"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame string `json:"alt2ReferenceFrame"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area float64 `json:"area"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot float64 `json:"bDot"`
	// Model parameter value for center of mass offset (m).
	CmOffset float64 `json:"cmOffset"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod string `json:"covMethod"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea float64 `json:"dragArea"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff float64 `json:"dragCoeff"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel string `json:"dragModel"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr float64 `json:"edr"`
	// Start time at which this state vector was the 'current' state vector for its
	// satellite.
	EffectiveFrom time.Time `json:"effectiveFrom" format:"date-time"`
	// End time at which this state vector was no longer the 'current' state vector for
	// its satellite.
	EffectiveUntil time.Time `json:"effectiveUntil" format:"date-time"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov"`
	// Integrator error control.
	ErrorControl float64 `json:"errorControl"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep bool `json:"fixedStep"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel string `json:"geopotentialModel"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms int64 `json:"iau1980Terms"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector string `json:"idStateVector"`
	// Integrator Mode.
	IntegratorMode string `json:"integratorMode"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust bool `json:"inTrackThrust"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd time.Time `json:"lastObEnd" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart time.Time `json:"lastObStart" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime time.Time `json:"leapSecondTime" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar bool `json:"lunarSolar"`
	// The mass of the object, in kilograms.
	Mass float64 `json:"mass"`
	// The number of observations available for the OD of the object.
	ObsAvailable int64 `json:"obsAvailable"`
	// The number of observations accepted for the OD of the object.
	ObsUsed int64 `json:"obsUsed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials string `json:"partials"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree string `json:"pedigree"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX float64 `json:"polarMotionX"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY float64 `json:"polarMotionY"`
	// One sigma position uncertainty, in kilometers.
	PosUnc float64 `json:"posUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan float64 `json:"recODSpan"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc float64 `json:"residualsAcc"`
	// Epoch revolution number.
	RevNo int64 `json:"revNo"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms float64 `json:"rms"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo int64 `json:"satNo"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg float64 `json:"solarFluxAPAvg"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 float64 `json:"solarFluxF10"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg float64 `json:"solarFluxF10Avg"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress bool `json:"solarRadPress"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff float64 `json:"solarRadPressCoeff"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides bool `json:"solidEarthTides"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData"`
	// Optional array of UDL observation data types used to build this state vector
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
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea float64 `json:"srpArea"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode string `json:"stepMode"`
	// Initial integration step size (seconds).
	StepSize float64 `json:"stepSize"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection string `json:"stepSizeSelection"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc float64 `json:"taiUtc"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel float64 `json:"thrustAccel"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail int64 `json:"tracksAvail"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed int64 `json:"tracksUsed"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct bool `json:"uct"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate float64 `json:"ut1Rate"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc float64 `json:"ut1Utc"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc float64 `json:"velUnc"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel float64 `json:"xaccel"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 float64 `json:"xposAlt1"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 float64 `json:"xposAlt2"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 float64 `json:"xvelAlt1"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 float64 `json:"xvelAlt2"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel float64 `json:"yaccel"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 float64 `json:"yposAlt1"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 float64 `json:"yposAlt2"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 float64 `json:"yvelAlt1"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 float64 `json:"yvelAlt2"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel float64 `json:"zaccel"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 float64 `json:"zposAlt1"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 float64 `json:"zposAlt2"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 float64 `json:"zvelAlt1"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 float64 `json:"zvelAlt2"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Epoch                 resp.Field
		Source                resp.Field
		ActualOdSpan          resp.Field
		Algorithm             resp.Field
		Alt1ReferenceFrame    resp.Field
		Alt2ReferenceFrame    resp.Field
		Area                  resp.Field
		BDot                  resp.Field
		CmOffset              resp.Field
		Cov                   resp.Field
		CovMethod             resp.Field
		CovReferenceFrame     resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		DragArea              resp.Field
		DragCoeff             resp.Field
		DragModel             resp.Field
		Edr                   resp.Field
		EffectiveFrom         resp.Field
		EffectiveUntil        resp.Field
		EqCov                 resp.Field
		ErrorControl          resp.Field
		FixedStep             resp.Field
		GeopotentialModel     resp.Field
		Iau1980Terms          resp.Field
		IDOnOrbit             resp.Field
		IDOrbitDetermination  resp.Field
		IDStateVector         resp.Field
		IntegratorMode        resp.Field
		InTrackThrust         resp.Field
		LastObEnd             resp.Field
		LastObStart           resp.Field
		LeapSecondTime        resp.Field
		LunarSolar            resp.Field
		Mass                  resp.Field
		ObsAvailable          resp.Field
		ObsUsed               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		Partials              resp.Field
		Pedigree              resp.Field
		PolarMotionX          resp.Field
		PolarMotionY          resp.Field
		PosUnc                resp.Field
		RawFileUri            resp.Field
		RecOdSpan             resp.Field
		ReferenceFrame        resp.Field
		ResidualsAcc          resp.Field
		RevNo                 resp.Field
		Rms                   resp.Field
		SatNo                 resp.Field
		SigmaPosUvw           resp.Field
		SigmaVelUvw           resp.Field
		SolarFluxApAvg        resp.Field
		SolarFluxF10          resp.Field
		SolarFluxF10Avg       resp.Field
		SolarRadPress         resp.Field
		SolarRadPressCoeff    resp.Field
		SolidEarthTides       resp.Field
		SourcedData           resp.Field
		SourcedDataTypes      resp.Field
		SourceDl              resp.Field
		SrpArea               resp.Field
		StepMode              resp.Field
		StepSize              resp.Field
		StepSizeSelection     resp.Field
		Tags                  resp.Field
		TaiUtc                resp.Field
		ThrustAccel           resp.Field
		TracksAvail           resp.Field
		TracksUsed            resp.Field
		TransactionID         resp.Field
		Uct                   resp.Field
		Ut1Rate               resp.Field
		Ut1Utc                resp.Field
		VelUnc                resp.Field
		Xaccel                resp.Field
		Xpos                  resp.Field
		XposAlt1              resp.Field
		XposAlt2              resp.Field
		Xvel                  resp.Field
		XvelAlt1              resp.Field
		XvelAlt2              resp.Field
		Yaccel                resp.Field
		Ypos                  resp.Field
		YposAlt1              resp.Field
		YposAlt2              resp.Field
		Yvel                  resp.Field
		YvelAlt1              resp.Field
		YvelAlt2              resp.Field
		Zaccel                resp.Field
		Zpos                  resp.Field
		ZposAlt1              resp.Field
		ZposAlt2              resp.Field
		Zvel                  resp.Field
		ZvelAlt1              resp.Field
		ZvelAlt2              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManeuverGetResponsePostEventStateVector) RawJSON() string { return r.JSON.raw }
func (r *ManeuverGetResponsePostEventStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
type ManeuverGetResponsePreEventElset struct {
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
func (r ManeuverGetResponsePreEventElset) RawJSON() string { return r.JSON.raw }
func (r *ManeuverGetResponsePreEventElset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
type ManeuverGetResponsePreEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan float64 `json:"actualODSpan"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame string `json:"alt1ReferenceFrame"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame string `json:"alt2ReferenceFrame"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area float64 `json:"area"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot float64 `json:"bDot"`
	// Model parameter value for center of mass offset (m).
	CmOffset float64 `json:"cmOffset"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod string `json:"covMethod"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea float64 `json:"dragArea"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff float64 `json:"dragCoeff"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel string `json:"dragModel"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr float64 `json:"edr"`
	// Start time at which this state vector was the 'current' state vector for its
	// satellite.
	EffectiveFrom time.Time `json:"effectiveFrom" format:"date-time"`
	// End time at which this state vector was no longer the 'current' state vector for
	// its satellite.
	EffectiveUntil time.Time `json:"effectiveUntil" format:"date-time"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov"`
	// Integrator error control.
	ErrorControl float64 `json:"errorControl"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep bool `json:"fixedStep"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel string `json:"geopotentialModel"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms int64 `json:"iau1980Terms"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector string `json:"idStateVector"`
	// Integrator Mode.
	IntegratorMode string `json:"integratorMode"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust bool `json:"inTrackThrust"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd time.Time `json:"lastObEnd" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart time.Time `json:"lastObStart" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime time.Time `json:"leapSecondTime" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar bool `json:"lunarSolar"`
	// The mass of the object, in kilograms.
	Mass float64 `json:"mass"`
	// The number of observations available for the OD of the object.
	ObsAvailable int64 `json:"obsAvailable"`
	// The number of observations accepted for the OD of the object.
	ObsUsed int64 `json:"obsUsed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials string `json:"partials"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree string `json:"pedigree"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX float64 `json:"polarMotionX"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY float64 `json:"polarMotionY"`
	// One sigma position uncertainty, in kilometers.
	PosUnc float64 `json:"posUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan float64 `json:"recODSpan"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc float64 `json:"residualsAcc"`
	// Epoch revolution number.
	RevNo int64 `json:"revNo"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms float64 `json:"rms"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo int64 `json:"satNo"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg float64 `json:"solarFluxAPAvg"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 float64 `json:"solarFluxF10"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg float64 `json:"solarFluxF10Avg"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress bool `json:"solarRadPress"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff float64 `json:"solarRadPressCoeff"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides bool `json:"solidEarthTides"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData"`
	// Optional array of UDL observation data types used to build this state vector
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
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea float64 `json:"srpArea"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode string `json:"stepMode"`
	// Initial integration step size (seconds).
	StepSize float64 `json:"stepSize"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection string `json:"stepSizeSelection"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc float64 `json:"taiUtc"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel float64 `json:"thrustAccel"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail int64 `json:"tracksAvail"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed int64 `json:"tracksUsed"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct bool `json:"uct"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate float64 `json:"ut1Rate"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc float64 `json:"ut1Utc"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc float64 `json:"velUnc"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel float64 `json:"xaccel"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 float64 `json:"xposAlt1"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 float64 `json:"xposAlt2"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 float64 `json:"xvelAlt1"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 float64 `json:"xvelAlt2"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel float64 `json:"yaccel"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 float64 `json:"yposAlt1"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 float64 `json:"yposAlt2"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 float64 `json:"yvelAlt1"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 float64 `json:"yvelAlt2"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel float64 `json:"zaccel"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 float64 `json:"zposAlt1"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 float64 `json:"zposAlt2"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 float64 `json:"zvelAlt1"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 float64 `json:"zvelAlt2"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Epoch                 resp.Field
		Source                resp.Field
		ActualOdSpan          resp.Field
		Algorithm             resp.Field
		Alt1ReferenceFrame    resp.Field
		Alt2ReferenceFrame    resp.Field
		Area                  resp.Field
		BDot                  resp.Field
		CmOffset              resp.Field
		Cov                   resp.Field
		CovMethod             resp.Field
		CovReferenceFrame     resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		DragArea              resp.Field
		DragCoeff             resp.Field
		DragModel             resp.Field
		Edr                   resp.Field
		EffectiveFrom         resp.Field
		EffectiveUntil        resp.Field
		EqCov                 resp.Field
		ErrorControl          resp.Field
		FixedStep             resp.Field
		GeopotentialModel     resp.Field
		Iau1980Terms          resp.Field
		IDOnOrbit             resp.Field
		IDOrbitDetermination  resp.Field
		IDStateVector         resp.Field
		IntegratorMode        resp.Field
		InTrackThrust         resp.Field
		LastObEnd             resp.Field
		LastObStart           resp.Field
		LeapSecondTime        resp.Field
		LunarSolar            resp.Field
		Mass                  resp.Field
		ObsAvailable          resp.Field
		ObsUsed               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		Partials              resp.Field
		Pedigree              resp.Field
		PolarMotionX          resp.Field
		PolarMotionY          resp.Field
		PosUnc                resp.Field
		RawFileUri            resp.Field
		RecOdSpan             resp.Field
		ReferenceFrame        resp.Field
		ResidualsAcc          resp.Field
		RevNo                 resp.Field
		Rms                   resp.Field
		SatNo                 resp.Field
		SigmaPosUvw           resp.Field
		SigmaVelUvw           resp.Field
		SolarFluxApAvg        resp.Field
		SolarFluxF10          resp.Field
		SolarFluxF10Avg       resp.Field
		SolarRadPress         resp.Field
		SolarRadPressCoeff    resp.Field
		SolidEarthTides       resp.Field
		SourcedData           resp.Field
		SourcedDataTypes      resp.Field
		SourceDl              resp.Field
		SrpArea               resp.Field
		StepMode              resp.Field
		StepSize              resp.Field
		StepSizeSelection     resp.Field
		Tags                  resp.Field
		TaiUtc                resp.Field
		ThrustAccel           resp.Field
		TracksAvail           resp.Field
		TracksUsed            resp.Field
		TransactionID         resp.Field
		Uct                   resp.Field
		Ut1Rate               resp.Field
		Ut1Utc                resp.Field
		VelUnc                resp.Field
		Xaccel                resp.Field
		Xpos                  resp.Field
		XposAlt1              resp.Field
		XposAlt2              resp.Field
		Xvel                  resp.Field
		XvelAlt1              resp.Field
		XvelAlt2              resp.Field
		Yaccel                resp.Field
		Ypos                  resp.Field
		YposAlt1              resp.Field
		YposAlt2              resp.Field
		Yvel                  resp.Field
		YvelAlt1              resp.Field
		YvelAlt2              resp.Field
		Zaccel                resp.Field
		Zpos                  resp.Field
		ZposAlt1              resp.Field
		ZposAlt2              resp.Field
		Zvel                  resp.Field
		ZvelAlt1              resp.Field
		ZvelAlt2              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManeuverGetResponsePreEventStateVector) RawJSON() string { return r.JSON.raw }
func (r *ManeuverGetResponsePreEventStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of on-orbit object maneuver information for detected,
// possible, and confirmed maneuvers.
type ManeuverTupleResponse struct {
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
	DataMode ManeuverTupleResponseDataMode `json:"dataMode,required"`
	// Maneuver event start time in ISO 8601 UTC with microsecond precision. For
	// maneuvers without start and end times, the start time is considered to be the
	// maneuver event time.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// Optional purpose of the maneuver if known (e.g. North-South Station Keeping,
	// East-West Station Keeping, Longitude Shift, Unknown).
	Characterization string `json:"characterization"`
	// Uncertainty in the characterization or purpose assessment of this maneuver (0 -
	// 1).
	CharacterizationUnc float64 `json:"characterizationUnc"`
	// Optional maneuver cross-track/radial/in-track covariance array, in meter and
	// second based units, in the following order: CR_R, CI_R, CI_I, CC_R, CC_I, CC_C,
	// CT_R, CT_I, CT_C, CT_T.
	Cov []float64 `json:"cov"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Difference in mass before and after the maneuver, in kg.
	DeltaMass float64 `json:"deltaMass"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors at the maneuver event time.
	DeltaPos float64 `json:"deltaPos"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaPosU float64 `json:"deltaPosU"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaPosV float64 `json:"deltaPosV"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaPosW float64 `json:"deltaPosW"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors at the maneuver event time.
	DeltaVel float64 `json:"deltaVel"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaVelU float64 `json:"deltaVelU"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaVelV float64 `json:"deltaVelV"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaVelW float64 `json:"deltaVelW"`
	// Description and notes of the maneuver.
	Description string `json:"description"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Maneuver event end time in ISO 8601 UTC with microsecond precision.
	EventEndTime time.Time `json:"eventEndTime" format:"date-time"`
	// Optional source-provided identifier for this maneuver event. In the case where
	// multiple maneuver records are submitted for the same event, this field can be
	// used to tie them together to the same event.
	EventID string `json:"eventId"`
	// Target maneuvering on-orbit object. For the public catalog, the idOnOrbit is
	// typically the satellite number as a string, but may be a UUID for analyst or
	// other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Optional ID of the sensor that detected this maneuver (for example, if detected
	// by passive RF anomalies).
	IDSensor string `json:"idSensor"`
	// Uncertainty in the occurrence of this maneuver (0 - 1).
	ManeuverUnc float64 `json:"maneuverUnc"`
	// Array of estimated acceleration values, in meters per second squared. Number of
	// elements must match the numAccelPoints.
	MnvrAccels []float64 `json:"mnvrAccels"`
	// Array of elapsed times, in seconds from maneuver start time, at which each
	// acceleration point is estimated. Number of elements must match the
	// numAccelPoints.
	MnvrAccelTimes []float64 `json:"mnvrAccelTimes"`
	// Array of the 1-sigma uncertainties in estimated accelerations, in meters per
	// second squared. Number of elements must match the numAccelPoints.
	MnvrAccelUncs []float64 `json:"mnvrAccelUncs"`
	// The total number of estimated acceleration points during the maneuver.
	NumAccelPoints int64 `json:"numAccelPoints"`
	// Number of observations used to generate the maneuver data.
	NumObs int64 `json:"numObs"`
	// Maneuver orbit determination fit data end time in ISO 8601 UTC with microsecond
	// precision.
	OdFitEndTime time.Time `json:"odFitEndTime" format:"date-time"`
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
	// Identifier provided by source to indicate the target on-orbit object performing
	// this maneuver. This may be an internal identifier and not necessarily a valid
	// satellite number/ID.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Post-event spacecraft apogee (measured from Earth center), in kilometers.
	PostApogee float64 `json:"postApogee"`
	// Estimated area of the object following the maneuver, in meters squared.
	PostArea float64 `json:"postArea"`
	// Post-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PostBallisticCoeff float64 `json:"postBallisticCoeff"`
	// Post-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PostDriftRate float64 `json:"postDriftRate"`
	// Post-event spacecraft eccentricity.
	PostEccentricity float64 `json:"postEccentricity"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PostEventElset ManeuverTupleResponsePostEventElset `json:"postEventElset"`
	// Optional identifier of the element set for the post-maneuver orbit.
	PostEventIDElset string `json:"postEventIdElset"`
	// Optional identifier of the state vector for the post-maneuver trajectory of the
	// spacecraft.
	PostEventIDStateVector string `json:"postEventIdStateVector"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PostEventStateVector ManeuverTupleResponsePostEventStateVector `json:"postEventStateVector"`
	// Post-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PostGeoLongitude float64 `json:"postGeoLongitude"`
	// Post-event spacecraft orbital inclination, in degrees. 0-180.
	PostInclination float64 `json:"postInclination"`
	// Estimated mass of the object following the maneuver, in kg.
	PostMass float64 `json:"postMass"`
	// Post-event spacecraft perigee (measured from Earth center), in kilometers.
	PostPerigee float64 `json:"postPerigee"`
	// Post-event spacecraft orbital period, in minutes.
	PostPeriod float64 `json:"postPeriod"`
	// Post-event X component of position in ECI space, in km.
	PostPosX float64 `json:"postPosX"`
	// Post-event Y component of position in ECI space, in km.
	PostPosY float64 `json:"postPosY"`
	// Post-event Z component of position in ECI space, in km.
	PostPosZ float64 `json:"postPosZ"`
	// Post-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PostRaan float64 `json:"postRAAN"`
	// Post-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PostRadiationPressCoeff float64 `json:"postRadiationPressCoeff"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PostSigmaU float64 `json:"postSigmaU"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PostSigmaV float64 `json:"postSigmaV"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PostSigmaW float64 `json:"postSigmaW"`
	// Post-event spacecraft Semi-Major Axis (SMA), in kilometers.
	PostSma float64 `json:"postSMA"`
	// Post-event X component of velocity in ECI space, in km/sec.
	PostVelX float64 `json:"postVelX"`
	// Post-event Y component of velocity in ECI space, in km/sec.
	PostVelY float64 `json:"postVelY"`
	// Post-event Z component of velocity in ECI space, in km/sec.
	PostVelZ float64 `json:"postVelZ"`
	// Pre-event spacecraft apogee (measured from Earth center), in kilometers.
	PreApogee float64 `json:"preApogee"`
	// Pre-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PreBallisticCoeff float64 `json:"preBallisticCoeff"`
	// Pre-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PreDriftRate float64 `json:"preDriftRate"`
	// Pre-event spacecraft eccentricity.
	PreEccentricity float64 `json:"preEccentricity"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PreEventElset ManeuverTupleResponsePreEventElset `json:"preEventElset"`
	// Optional identifier of the element set for the pre-maneuver orbit.
	PreEventIDElset string `json:"preEventIdElset"`
	// Optional identifier of the state vector for the pre-maneuver trajectory of the
	// spacecraft.
	PreEventIDStateVector string `json:"preEventIdStateVector"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PreEventStateVector ManeuverTupleResponsePreEventStateVector `json:"preEventStateVector"`
	// Pre-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PreGeoLongitude float64 `json:"preGeoLongitude"`
	// Pre-event spacecraft orbital inclination, in degrees. 0-180.
	PreInclination float64 `json:"preInclination"`
	// Pre-event spacecraft perigee (measured from Earth center), in kilometers.
	PrePerigee float64 `json:"prePerigee"`
	// Pre-event spacecraft orbital period, in minutes.
	PrePeriod float64 `json:"prePeriod"`
	// Pre-event X component of position in ECI space, in km.
	PrePosX float64 `json:"prePosX"`
	// Pre-event Y component of position in ECI space, in km.
	PrePosY float64 `json:"prePosY"`
	// Pre-event Z component of position in ECI space, in km.
	PrePosZ float64 `json:"prePosZ"`
	// Pre-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PreRaan float64 `json:"preRAAN"`
	// Pre-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PreRadiationPressCoeff float64 `json:"preRadiationPressCoeff"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PreSigmaU float64 `json:"preSigmaU"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PreSigmaV float64 `json:"preSigmaV"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PreSigmaW float64 `json:"preSigmaW"`
	// Pre-event spacecraft orbital Semi-Major Axis (SMA), in kilometers.
	PreSma float64 `json:"preSMA"`
	// Pre-event X component of velocity in ECI space, in km/sec.
	PreVelX float64 `json:"preVelX"`
	// Pre-event Y component of velocity in ECI space, in km/sec.
	PreVelY float64 `json:"preVelY"`
	// Pre-event Z component of velocity in ECI space, in km/sec.
	PreVelZ float64 `json:"preVelZ"`
	// The time that the report or alert of this maneuver was generated, in ISO 8601
	// UTC format.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Optional array of UDL data (elsets, state vectors, etc) UUIDs used to build this
	// maneuver. See the associated sourcedDataTypes array for the specific types of
	// data for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData"`
	// Optional array of UDL data types used to build this maneuver (e.g. EO, RADAR,
	// RF, DOA, ELSET, SV). See the associated sourcedData array for the specific UUIDs
	// of data for the positionally corresponding data types in this array (the two
	// arrays must match in size).
	//
	// Any of "EO", "RADAR", "RF", "DOA", "ELSET", "SV".
	SourcedDataTypes []string `json:"sourcedDataTypes"`
	// Name of the state model used to generate the maneuver data.
	StateModel string `json:"stateModel"`
	// Version of the state model used to generate the maneuver data.
	StateModelVersion float64 `json:"stateModelVersion"`
	// Status of this maneuver (CANCELLED, PLANNED, POSSIBLE, REDACTED, VERIFIED).
	//
	// CANCELLED: A previously planned maneuver whose execution was cancelled.
	//
	// PLANNED: A maneuver planned to take place at the eventStartTime.
	//
	// POSSIBLE: A possible maneuver detected by observation of the spacecraft or by
	// evaluation of the spacecraft orbit.
	//
	// REDACTED: A redaction of a reported possible maneuver that has been determined
	// to have not taken place after further observation/evaluation.
	//
	// VERIFIED: A maneuver whose execution has been verified, either by the
	// owner/operator or observation/evaluation.
	Status string `json:"status"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// The estimated total active burn time of a maneuver, in seconds. This includes
	// the sum of all burns in numAccelPoints. Not to be confused with the total
	// duration of the maneuver.
	TotalBurnTime float64 `json:"totalBurnTime"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this maneuver was unable to be correlated to a known object.
	// This flag should only be set to true by data providers after an attempt to
	// correlate to an on-orbit object was made and failed. If unable to correlate, the
	// 'origObjectId' field may be populated with an internal data provider specific
	// identifier.
	Uct bool `json:"uct"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking   resp.Field
		DataMode                resp.Field
		EventStartTime          resp.Field
		Source                  resp.Field
		ID                      resp.Field
		Algorithm               resp.Field
		Characterization        resp.Field
		CharacterizationUnc     resp.Field
		Cov                     resp.Field
		CreatedAt               resp.Field
		CreatedBy               resp.Field
		DeltaMass               resp.Field
		DeltaPos                resp.Field
		DeltaPosU               resp.Field
		DeltaPosV               resp.Field
		DeltaPosW               resp.Field
		DeltaVel                resp.Field
		DeltaVelU               resp.Field
		DeltaVelV               resp.Field
		DeltaVelW               resp.Field
		Description             resp.Field
		Descriptor              resp.Field
		EventEndTime            resp.Field
		EventID                 resp.Field
		IDOnOrbit               resp.Field
		IDSensor                resp.Field
		ManeuverUnc             resp.Field
		MnvrAccels              resp.Field
		MnvrAccelTimes          resp.Field
		MnvrAccelUncs           resp.Field
		NumAccelPoints          resp.Field
		NumObs                  resp.Field
		OdFitEndTime            resp.Field
		OnOrbit                 resp.Field
		Origin                  resp.Field
		OrigNetwork             resp.Field
		OrigObjectID            resp.Field
		OrigSensorID            resp.Field
		PostApogee              resp.Field
		PostArea                resp.Field
		PostBallisticCoeff      resp.Field
		PostDriftRate           resp.Field
		PostEccentricity        resp.Field
		PostEventElset          resp.Field
		PostEventIDElset        resp.Field
		PostEventIDStateVector  resp.Field
		PostEventStateVector    resp.Field
		PostGeoLongitude        resp.Field
		PostInclination         resp.Field
		PostMass                resp.Field
		PostPerigee             resp.Field
		PostPeriod              resp.Field
		PostPosX                resp.Field
		PostPosY                resp.Field
		PostPosZ                resp.Field
		PostRaan                resp.Field
		PostRadiationPressCoeff resp.Field
		PostSigmaU              resp.Field
		PostSigmaV              resp.Field
		PostSigmaW              resp.Field
		PostSma                 resp.Field
		PostVelX                resp.Field
		PostVelY                resp.Field
		PostVelZ                resp.Field
		PreApogee               resp.Field
		PreBallisticCoeff       resp.Field
		PreDriftRate            resp.Field
		PreEccentricity         resp.Field
		PreEventElset           resp.Field
		PreEventIDElset         resp.Field
		PreEventIDStateVector   resp.Field
		PreEventStateVector     resp.Field
		PreGeoLongitude         resp.Field
		PreInclination          resp.Field
		PrePerigee              resp.Field
		PrePeriod               resp.Field
		PrePosX                 resp.Field
		PrePosY                 resp.Field
		PrePosZ                 resp.Field
		PreRaan                 resp.Field
		PreRadiationPressCoeff  resp.Field
		PreSigmaU               resp.Field
		PreSigmaV               resp.Field
		PreSigmaW               resp.Field
		PreSma                  resp.Field
		PreVelX                 resp.Field
		PreVelY                 resp.Field
		PreVelZ                 resp.Field
		ReportTime              resp.Field
		SatNo                   resp.Field
		SourcedData             resp.Field
		SourcedDataTypes        resp.Field
		StateModel              resp.Field
		StateModelVersion       resp.Field
		Status                  resp.Field
		Tags                    resp.Field
		TotalBurnTime           resp.Field
		TransactionID           resp.Field
		Uct                     resp.Field
		ExtraFields             map[string]resp.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManeuverTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ManeuverTupleResponse) UnmarshalJSON(data []byte) error {
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
type ManeuverTupleResponseDataMode string

const (
	ManeuverTupleResponseDataModeReal      ManeuverTupleResponseDataMode = "REAL"
	ManeuverTupleResponseDataModeTest      ManeuverTupleResponseDataMode = "TEST"
	ManeuverTupleResponseDataModeSimulated ManeuverTupleResponseDataMode = "SIMULATED"
	ManeuverTupleResponseDataModeExercise  ManeuverTupleResponseDataMode = "EXERCISE"
)

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
type ManeuverTupleResponsePostEventElset struct {
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
func (r ManeuverTupleResponsePostEventElset) RawJSON() string { return r.JSON.raw }
func (r *ManeuverTupleResponsePostEventElset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
type ManeuverTupleResponsePostEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan float64 `json:"actualODSpan"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame string `json:"alt1ReferenceFrame"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame string `json:"alt2ReferenceFrame"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area float64 `json:"area"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot float64 `json:"bDot"`
	// Model parameter value for center of mass offset (m).
	CmOffset float64 `json:"cmOffset"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod string `json:"covMethod"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea float64 `json:"dragArea"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff float64 `json:"dragCoeff"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel string `json:"dragModel"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr float64 `json:"edr"`
	// Start time at which this state vector was the 'current' state vector for its
	// satellite.
	EffectiveFrom time.Time `json:"effectiveFrom" format:"date-time"`
	// End time at which this state vector was no longer the 'current' state vector for
	// its satellite.
	EffectiveUntil time.Time `json:"effectiveUntil" format:"date-time"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov"`
	// Integrator error control.
	ErrorControl float64 `json:"errorControl"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep bool `json:"fixedStep"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel string `json:"geopotentialModel"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms int64 `json:"iau1980Terms"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector string `json:"idStateVector"`
	// Integrator Mode.
	IntegratorMode string `json:"integratorMode"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust bool `json:"inTrackThrust"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd time.Time `json:"lastObEnd" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart time.Time `json:"lastObStart" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime time.Time `json:"leapSecondTime" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar bool `json:"lunarSolar"`
	// The mass of the object, in kilograms.
	Mass float64 `json:"mass"`
	// The number of observations available for the OD of the object.
	ObsAvailable int64 `json:"obsAvailable"`
	// The number of observations accepted for the OD of the object.
	ObsUsed int64 `json:"obsUsed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials string `json:"partials"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree string `json:"pedigree"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX float64 `json:"polarMotionX"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY float64 `json:"polarMotionY"`
	// One sigma position uncertainty, in kilometers.
	PosUnc float64 `json:"posUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan float64 `json:"recODSpan"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc float64 `json:"residualsAcc"`
	// Epoch revolution number.
	RevNo int64 `json:"revNo"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms float64 `json:"rms"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo int64 `json:"satNo"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg float64 `json:"solarFluxAPAvg"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 float64 `json:"solarFluxF10"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg float64 `json:"solarFluxF10Avg"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress bool `json:"solarRadPress"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff float64 `json:"solarRadPressCoeff"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides bool `json:"solidEarthTides"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData"`
	// Optional array of UDL observation data types used to build this state vector
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
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea float64 `json:"srpArea"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode string `json:"stepMode"`
	// Initial integration step size (seconds).
	StepSize float64 `json:"stepSize"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection string `json:"stepSizeSelection"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc float64 `json:"taiUtc"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel float64 `json:"thrustAccel"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail int64 `json:"tracksAvail"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed int64 `json:"tracksUsed"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct bool `json:"uct"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate float64 `json:"ut1Rate"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc float64 `json:"ut1Utc"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc float64 `json:"velUnc"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel float64 `json:"xaccel"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 float64 `json:"xposAlt1"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 float64 `json:"xposAlt2"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 float64 `json:"xvelAlt1"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 float64 `json:"xvelAlt2"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel float64 `json:"yaccel"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 float64 `json:"yposAlt1"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 float64 `json:"yposAlt2"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 float64 `json:"yvelAlt1"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 float64 `json:"yvelAlt2"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel float64 `json:"zaccel"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 float64 `json:"zposAlt1"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 float64 `json:"zposAlt2"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 float64 `json:"zvelAlt1"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 float64 `json:"zvelAlt2"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Epoch                 resp.Field
		Source                resp.Field
		ActualOdSpan          resp.Field
		Algorithm             resp.Field
		Alt1ReferenceFrame    resp.Field
		Alt2ReferenceFrame    resp.Field
		Area                  resp.Field
		BDot                  resp.Field
		CmOffset              resp.Field
		Cov                   resp.Field
		CovMethod             resp.Field
		CovReferenceFrame     resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		DragArea              resp.Field
		DragCoeff             resp.Field
		DragModel             resp.Field
		Edr                   resp.Field
		EffectiveFrom         resp.Field
		EffectiveUntil        resp.Field
		EqCov                 resp.Field
		ErrorControl          resp.Field
		FixedStep             resp.Field
		GeopotentialModel     resp.Field
		Iau1980Terms          resp.Field
		IDOnOrbit             resp.Field
		IDOrbitDetermination  resp.Field
		IDStateVector         resp.Field
		IntegratorMode        resp.Field
		InTrackThrust         resp.Field
		LastObEnd             resp.Field
		LastObStart           resp.Field
		LeapSecondTime        resp.Field
		LunarSolar            resp.Field
		Mass                  resp.Field
		ObsAvailable          resp.Field
		ObsUsed               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		Partials              resp.Field
		Pedigree              resp.Field
		PolarMotionX          resp.Field
		PolarMotionY          resp.Field
		PosUnc                resp.Field
		RawFileUri            resp.Field
		RecOdSpan             resp.Field
		ReferenceFrame        resp.Field
		ResidualsAcc          resp.Field
		RevNo                 resp.Field
		Rms                   resp.Field
		SatNo                 resp.Field
		SigmaPosUvw           resp.Field
		SigmaVelUvw           resp.Field
		SolarFluxApAvg        resp.Field
		SolarFluxF10          resp.Field
		SolarFluxF10Avg       resp.Field
		SolarRadPress         resp.Field
		SolarRadPressCoeff    resp.Field
		SolidEarthTides       resp.Field
		SourcedData           resp.Field
		SourcedDataTypes      resp.Field
		SourceDl              resp.Field
		SrpArea               resp.Field
		StepMode              resp.Field
		StepSize              resp.Field
		StepSizeSelection     resp.Field
		Tags                  resp.Field
		TaiUtc                resp.Field
		ThrustAccel           resp.Field
		TracksAvail           resp.Field
		TracksUsed            resp.Field
		TransactionID         resp.Field
		Uct                   resp.Field
		Ut1Rate               resp.Field
		Ut1Utc                resp.Field
		VelUnc                resp.Field
		Xaccel                resp.Field
		Xpos                  resp.Field
		XposAlt1              resp.Field
		XposAlt2              resp.Field
		Xvel                  resp.Field
		XvelAlt1              resp.Field
		XvelAlt2              resp.Field
		Yaccel                resp.Field
		Ypos                  resp.Field
		YposAlt1              resp.Field
		YposAlt2              resp.Field
		Yvel                  resp.Field
		YvelAlt1              resp.Field
		YvelAlt2              resp.Field
		Zaccel                resp.Field
		Zpos                  resp.Field
		ZposAlt1              resp.Field
		ZposAlt2              resp.Field
		Zvel                  resp.Field
		ZvelAlt1              resp.Field
		ZvelAlt2              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManeuverTupleResponsePostEventStateVector) RawJSON() string { return r.JSON.raw }
func (r *ManeuverTupleResponsePostEventStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
type ManeuverTupleResponsePreEventElset struct {
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
func (r ManeuverTupleResponsePreEventElset) RawJSON() string { return r.JSON.raw }
func (r *ManeuverTupleResponsePreEventElset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
type ManeuverTupleResponsePreEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan float64 `json:"actualODSpan"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame string `json:"alt1ReferenceFrame"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame string `json:"alt2ReferenceFrame"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area float64 `json:"area"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot float64 `json:"bDot"`
	// Model parameter value for center of mass offset (m).
	CmOffset float64 `json:"cmOffset"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod string `json:"covMethod"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea float64 `json:"dragArea"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff float64 `json:"dragCoeff"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel string `json:"dragModel"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr float64 `json:"edr"`
	// Start time at which this state vector was the 'current' state vector for its
	// satellite.
	EffectiveFrom time.Time `json:"effectiveFrom" format:"date-time"`
	// End time at which this state vector was no longer the 'current' state vector for
	// its satellite.
	EffectiveUntil time.Time `json:"effectiveUntil" format:"date-time"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov"`
	// Integrator error control.
	ErrorControl float64 `json:"errorControl"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep bool `json:"fixedStep"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel string `json:"geopotentialModel"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms int64 `json:"iau1980Terms"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector string `json:"idStateVector"`
	// Integrator Mode.
	IntegratorMode string `json:"integratorMode"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust bool `json:"inTrackThrust"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd time.Time `json:"lastObEnd" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart time.Time `json:"lastObStart" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime time.Time `json:"leapSecondTime" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar bool `json:"lunarSolar"`
	// The mass of the object, in kilograms.
	Mass float64 `json:"mass"`
	// The number of observations available for the OD of the object.
	ObsAvailable int64 `json:"obsAvailable"`
	// The number of observations accepted for the OD of the object.
	ObsUsed int64 `json:"obsUsed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials string `json:"partials"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree string `json:"pedigree"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX float64 `json:"polarMotionX"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY float64 `json:"polarMotionY"`
	// One sigma position uncertainty, in kilometers.
	PosUnc float64 `json:"posUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan float64 `json:"recODSpan"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc float64 `json:"residualsAcc"`
	// Epoch revolution number.
	RevNo int64 `json:"revNo"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms float64 `json:"rms"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo int64 `json:"satNo"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg float64 `json:"solarFluxAPAvg"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 float64 `json:"solarFluxF10"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg float64 `json:"solarFluxF10Avg"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress bool `json:"solarRadPress"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff float64 `json:"solarRadPressCoeff"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides bool `json:"solidEarthTides"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData"`
	// Optional array of UDL observation data types used to build this state vector
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
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea float64 `json:"srpArea"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode string `json:"stepMode"`
	// Initial integration step size (seconds).
	StepSize float64 `json:"stepSize"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection string `json:"stepSizeSelection"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc float64 `json:"taiUtc"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel float64 `json:"thrustAccel"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail int64 `json:"tracksAvail"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed int64 `json:"tracksUsed"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct bool `json:"uct"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate float64 `json:"ut1Rate"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc float64 `json:"ut1Utc"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc float64 `json:"velUnc"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel float64 `json:"xaccel"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 float64 `json:"xposAlt1"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 float64 `json:"xposAlt2"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 float64 `json:"xvelAlt1"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 float64 `json:"xvelAlt2"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel float64 `json:"yaccel"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 float64 `json:"yposAlt1"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 float64 `json:"yposAlt2"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 float64 `json:"yvelAlt1"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 float64 `json:"yvelAlt2"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel float64 `json:"zaccel"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 float64 `json:"zposAlt1"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 float64 `json:"zposAlt2"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 float64 `json:"zvelAlt1"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 float64 `json:"zvelAlt2"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Epoch                 resp.Field
		Source                resp.Field
		ActualOdSpan          resp.Field
		Algorithm             resp.Field
		Alt1ReferenceFrame    resp.Field
		Alt2ReferenceFrame    resp.Field
		Area                  resp.Field
		BDot                  resp.Field
		CmOffset              resp.Field
		Cov                   resp.Field
		CovMethod             resp.Field
		CovReferenceFrame     resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		DragArea              resp.Field
		DragCoeff             resp.Field
		DragModel             resp.Field
		Edr                   resp.Field
		EffectiveFrom         resp.Field
		EffectiveUntil        resp.Field
		EqCov                 resp.Field
		ErrorControl          resp.Field
		FixedStep             resp.Field
		GeopotentialModel     resp.Field
		Iau1980Terms          resp.Field
		IDOnOrbit             resp.Field
		IDOrbitDetermination  resp.Field
		IDStateVector         resp.Field
		IntegratorMode        resp.Field
		InTrackThrust         resp.Field
		LastObEnd             resp.Field
		LastObStart           resp.Field
		LeapSecondTime        resp.Field
		LunarSolar            resp.Field
		Mass                  resp.Field
		ObsAvailable          resp.Field
		ObsUsed               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		Partials              resp.Field
		Pedigree              resp.Field
		PolarMotionX          resp.Field
		PolarMotionY          resp.Field
		PosUnc                resp.Field
		RawFileUri            resp.Field
		RecOdSpan             resp.Field
		ReferenceFrame        resp.Field
		ResidualsAcc          resp.Field
		RevNo                 resp.Field
		Rms                   resp.Field
		SatNo                 resp.Field
		SigmaPosUvw           resp.Field
		SigmaVelUvw           resp.Field
		SolarFluxApAvg        resp.Field
		SolarFluxF10          resp.Field
		SolarFluxF10Avg       resp.Field
		SolarRadPress         resp.Field
		SolarRadPressCoeff    resp.Field
		SolidEarthTides       resp.Field
		SourcedData           resp.Field
		SourcedDataTypes      resp.Field
		SourceDl              resp.Field
		SrpArea               resp.Field
		StepMode              resp.Field
		StepSize              resp.Field
		StepSizeSelection     resp.Field
		Tags                  resp.Field
		TaiUtc                resp.Field
		ThrustAccel           resp.Field
		TracksAvail           resp.Field
		TracksUsed            resp.Field
		TransactionID         resp.Field
		Uct                   resp.Field
		Ut1Rate               resp.Field
		Ut1Utc                resp.Field
		VelUnc                resp.Field
		Xaccel                resp.Field
		Xpos                  resp.Field
		XposAlt1              resp.Field
		XposAlt2              resp.Field
		Xvel                  resp.Field
		XvelAlt1              resp.Field
		XvelAlt2              resp.Field
		Yaccel                resp.Field
		Ypos                  resp.Field
		YposAlt1              resp.Field
		YposAlt2              resp.Field
		Yvel                  resp.Field
		YvelAlt1              resp.Field
		YvelAlt2              resp.Field
		Zaccel                resp.Field
		Zpos                  resp.Field
		ZposAlt1              resp.Field
		ZposAlt2              resp.Field
		Zvel                  resp.Field
		ZvelAlt1              resp.Field
		ZvelAlt2              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManeuverTupleResponsePreEventStateVector) RawJSON() string { return r.JSON.raw }
func (r *ManeuverTupleResponsePreEventStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManeuverNewParams struct {
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
	DataMode ManeuverNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Maneuver event start time in ISO 8601 UTC with microsecond precision. For
	// maneuvers without start and end times, the start time is considered to be the
	// maneuver event time.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// Optional purpose of the maneuver if known (e.g. North-South Station Keeping,
	// East-West Station Keeping, Longitude Shift, Unknown).
	Characterization param.Opt[string] `json:"characterization,omitzero"`
	// Uncertainty in the characterization or purpose assessment of this maneuver (0 -
	// 1).
	CharacterizationUnc param.Opt[float64] `json:"characterizationUnc,omitzero"`
	// Difference in mass before and after the maneuver, in kg.
	DeltaMass param.Opt[float64] `json:"deltaMass,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors at the maneuver event time.
	DeltaPos param.Opt[float64] `json:"deltaPos,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaPosU param.Opt[float64] `json:"deltaPosU,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaPosV param.Opt[float64] `json:"deltaPosV,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaPosW param.Opt[float64] `json:"deltaPosW,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors at the maneuver event time.
	DeltaVel param.Opt[float64] `json:"deltaVel,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaVelU param.Opt[float64] `json:"deltaVelU,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaVelV param.Opt[float64] `json:"deltaVelV,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaVelW param.Opt[float64] `json:"deltaVelW,omitzero"`
	// Description and notes of the maneuver.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Maneuver event end time in ISO 8601 UTC with microsecond precision.
	EventEndTime param.Opt[time.Time] `json:"eventEndTime,omitzero" format:"date-time"`
	// Optional source-provided identifier for this maneuver event. In the case where
	// multiple maneuver records are submitted for the same event, this field can be
	// used to tie them together to the same event.
	EventID param.Opt[string] `json:"eventId,omitzero"`
	// Optional ID of the sensor that detected this maneuver (for example, if detected
	// by passive RF anomalies).
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Uncertainty in the occurrence of this maneuver (0 - 1).
	ManeuverUnc param.Opt[float64] `json:"maneuverUnc,omitzero"`
	// The total number of estimated acceleration points during the maneuver.
	NumAccelPoints param.Opt[int64] `json:"numAccelPoints,omitzero"`
	// Number of observations used to generate the maneuver data.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// Maneuver orbit determination fit data end time in ISO 8601 UTC with microsecond
	// precision.
	OdFitEndTime param.Opt[time.Time] `json:"odFitEndTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Identifier provided by source to indicate the target on-orbit object performing
	// this maneuver. This may be an internal identifier and not necessarily a valid
	// satellite number/ID.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Post-event spacecraft apogee (measured from Earth center), in kilometers.
	PostApogee param.Opt[float64] `json:"postApogee,omitzero"`
	// Estimated area of the object following the maneuver, in meters squared.
	PostArea param.Opt[float64] `json:"postArea,omitzero"`
	// Post-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PostBallisticCoeff param.Opt[float64] `json:"postBallisticCoeff,omitzero"`
	// Post-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PostDriftRate param.Opt[float64] `json:"postDriftRate,omitzero"`
	// Post-event spacecraft eccentricity.
	PostEccentricity param.Opt[float64] `json:"postEccentricity,omitzero"`
	// Optional identifier of the element set for the post-maneuver orbit.
	PostEventIDElset param.Opt[string] `json:"postEventIdElset,omitzero"`
	// Optional identifier of the state vector for the post-maneuver trajectory of the
	// spacecraft.
	PostEventIDStateVector param.Opt[string] `json:"postEventIdStateVector,omitzero"`
	// Post-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PostGeoLongitude param.Opt[float64] `json:"postGeoLongitude,omitzero"`
	// Post-event spacecraft orbital inclination, in degrees. 0-180.
	PostInclination param.Opt[float64] `json:"postInclination,omitzero"`
	// Estimated mass of the object following the maneuver, in kg.
	PostMass param.Opt[float64] `json:"postMass,omitzero"`
	// Post-event spacecraft perigee (measured from Earth center), in kilometers.
	PostPerigee param.Opt[float64] `json:"postPerigee,omitzero"`
	// Post-event spacecraft orbital period, in minutes.
	PostPeriod param.Opt[float64] `json:"postPeriod,omitzero"`
	// Post-event X component of position in ECI space, in km.
	PostPosX param.Opt[float64] `json:"postPosX,omitzero"`
	// Post-event Y component of position in ECI space, in km.
	PostPosY param.Opt[float64] `json:"postPosY,omitzero"`
	// Post-event Z component of position in ECI space, in km.
	PostPosZ param.Opt[float64] `json:"postPosZ,omitzero"`
	// Post-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PostRaan param.Opt[float64] `json:"postRAAN,omitzero"`
	// Post-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PostRadiationPressCoeff param.Opt[float64] `json:"postRadiationPressCoeff,omitzero"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PostSigmaU param.Opt[float64] `json:"postSigmaU,omitzero"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PostSigmaV param.Opt[float64] `json:"postSigmaV,omitzero"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PostSigmaW param.Opt[float64] `json:"postSigmaW,omitzero"`
	// Post-event spacecraft Semi-Major Axis (SMA), in kilometers.
	PostSma param.Opt[float64] `json:"postSMA,omitzero"`
	// Post-event X component of velocity in ECI space, in km/sec.
	PostVelX param.Opt[float64] `json:"postVelX,omitzero"`
	// Post-event Y component of velocity in ECI space, in km/sec.
	PostVelY param.Opt[float64] `json:"postVelY,omitzero"`
	// Post-event Z component of velocity in ECI space, in km/sec.
	PostVelZ param.Opt[float64] `json:"postVelZ,omitzero"`
	// Pre-event spacecraft apogee (measured from Earth center), in kilometers.
	PreApogee param.Opt[float64] `json:"preApogee,omitzero"`
	// Pre-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PreBallisticCoeff param.Opt[float64] `json:"preBallisticCoeff,omitzero"`
	// Pre-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PreDriftRate param.Opt[float64] `json:"preDriftRate,omitzero"`
	// Pre-event spacecraft eccentricity.
	PreEccentricity param.Opt[float64] `json:"preEccentricity,omitzero"`
	// Optional identifier of the element set for the pre-maneuver orbit.
	PreEventIDElset param.Opt[string] `json:"preEventIdElset,omitzero"`
	// Optional identifier of the state vector for the pre-maneuver trajectory of the
	// spacecraft.
	PreEventIDStateVector param.Opt[string] `json:"preEventIdStateVector,omitzero"`
	// Pre-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PreGeoLongitude param.Opt[float64] `json:"preGeoLongitude,omitzero"`
	// Pre-event spacecraft orbital inclination, in degrees. 0-180.
	PreInclination param.Opt[float64] `json:"preInclination,omitzero"`
	// Pre-event spacecraft perigee (measured from Earth center), in kilometers.
	PrePerigee param.Opt[float64] `json:"prePerigee,omitzero"`
	// Pre-event spacecraft orbital period, in minutes.
	PrePeriod param.Opt[float64] `json:"prePeriod,omitzero"`
	// Pre-event X component of position in ECI space, in km.
	PrePosX param.Opt[float64] `json:"prePosX,omitzero"`
	// Pre-event Y component of position in ECI space, in km.
	PrePosY param.Opt[float64] `json:"prePosY,omitzero"`
	// Pre-event Z component of position in ECI space, in km.
	PrePosZ param.Opt[float64] `json:"prePosZ,omitzero"`
	// Pre-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PreRaan param.Opt[float64] `json:"preRAAN,omitzero"`
	// Pre-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PreRadiationPressCoeff param.Opt[float64] `json:"preRadiationPressCoeff,omitzero"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PreSigmaU param.Opt[float64] `json:"preSigmaU,omitzero"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PreSigmaV param.Opt[float64] `json:"preSigmaV,omitzero"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PreSigmaW param.Opt[float64] `json:"preSigmaW,omitzero"`
	// Pre-event spacecraft orbital Semi-Major Axis (SMA), in kilometers.
	PreSma param.Opt[float64] `json:"preSMA,omitzero"`
	// Pre-event X component of velocity in ECI space, in km/sec.
	PreVelX param.Opt[float64] `json:"preVelX,omitzero"`
	// Pre-event Y component of velocity in ECI space, in km/sec.
	PreVelY param.Opt[float64] `json:"preVelY,omitzero"`
	// Pre-event Z component of velocity in ECI space, in km/sec.
	PreVelZ param.Opt[float64] `json:"preVelZ,omitzero"`
	// The time that the report or alert of this maneuver was generated, in ISO 8601
	// UTC format.
	ReportTime param.Opt[time.Time] `json:"reportTime,omitzero" format:"date-time"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Name of the state model used to generate the maneuver data.
	StateModel param.Opt[string] `json:"stateModel,omitzero"`
	// Version of the state model used to generate the maneuver data.
	StateModelVersion param.Opt[float64] `json:"stateModelVersion,omitzero"`
	// Status of this maneuver (CANCELLED, PLANNED, POSSIBLE, REDACTED, VERIFIED).
	//
	// CANCELLED: A previously planned maneuver whose execution was cancelled.
	//
	// PLANNED: A maneuver planned to take place at the eventStartTime.
	//
	// POSSIBLE: A possible maneuver detected by observation of the spacecraft or by
	// evaluation of the spacecraft orbit.
	//
	// REDACTED: A redaction of a reported possible maneuver that has been determined
	// to have not taken place after further observation/evaluation.
	//
	// VERIFIED: A maneuver whose execution has been verified, either by the
	// owner/operator or observation/evaluation.
	Status param.Opt[string] `json:"status,omitzero"`
	// The estimated total active burn time of a maneuver, in seconds. This includes
	// the sum of all burns in numAccelPoints. Not to be confused with the total
	// duration of the maneuver.
	TotalBurnTime param.Opt[float64] `json:"totalBurnTime,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this maneuver was unable to be correlated to a known object.
	// This flag should only be set to true by data providers after an attempt to
	// correlate to an on-orbit object was made and failed. If unable to correlate, the
	// 'origObjectId' field may be populated with an internal data provider specific
	// identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Optional maneuver cross-track/radial/in-track covariance array, in meter and
	// second based units, in the following order: CR_R, CI_R, CI_I, CC_R, CC_I, CC_C,
	// CT_R, CT_I, CT_C, CT_T.
	Cov []float64 `json:"cov,omitzero"`
	// Array of estimated acceleration values, in meters per second squared. Number of
	// elements must match the numAccelPoints.
	MnvrAccels []float64 `json:"mnvrAccels,omitzero"`
	// Array of elapsed times, in seconds from maneuver start time, at which each
	// acceleration point is estimated. Number of elements must match the
	// numAccelPoints.
	MnvrAccelTimes []float64 `json:"mnvrAccelTimes,omitzero"`
	// Array of the 1-sigma uncertainties in estimated accelerations, in meters per
	// second squared. Number of elements must match the numAccelPoints.
	MnvrAccelUncs []float64 `json:"mnvrAccelUncs,omitzero"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PostEventElset ManeuverNewParamsPostEventElset `json:"postEventElset,omitzero"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PostEventStateVector ManeuverNewParamsPostEventStateVector `json:"postEventStateVector,omitzero"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PreEventElset ManeuverNewParamsPreEventElset `json:"preEventElset,omitzero"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PreEventStateVector ManeuverNewParamsPreEventStateVector `json:"preEventStateVector,omitzero"`
	// Optional array of UDL data (elsets, state vectors, etc) UUIDs used to build this
	// maneuver. See the associated sourcedDataTypes array for the specific types of
	// data for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL data types used to build this maneuver (e.g. EO, RADAR,
	// RF, DOA, ELSET, SV). See the associated sourcedData array for the specific UUIDs
	// of data for the positionally corresponding data types in this array (the two
	// arrays must match in size).
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
func (f ManeuverNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ManeuverNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewParams
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
type ManeuverNewParamsDataMode string

const (
	ManeuverNewParamsDataModeReal      ManeuverNewParamsDataMode = "REAL"
	ManeuverNewParamsDataModeTest      ManeuverNewParamsDataMode = "TEST"
	ManeuverNewParamsDataModeSimulated ManeuverNewParamsDataMode = "SIMULATED"
	ManeuverNewParamsDataModeExercise  ManeuverNewParamsDataMode = "EXERCISE"
)

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverNewParamsPostEventElset struct {
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
func (f ManeuverNewParamsPostEventElset) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ManeuverNewParamsPostEventElset) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewParamsPostEventElset
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverNewParamsPostEventElset](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverNewParamsPostEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan param.Opt[float64] `json:"actualODSpan,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame param.Opt[string] `json:"alt1ReferenceFrame,omitzero"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame param.Opt[string] `json:"alt2ReferenceFrame,omitzero"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area param.Opt[float64] `json:"area,omitzero"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot param.Opt[float64] `json:"bDot,omitzero"`
	// Model parameter value for center of mass offset (m).
	CmOffset param.Opt[float64] `json:"cmOffset,omitzero"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod param.Opt[string] `json:"covMethod,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea param.Opt[float64] `json:"dragArea,omitzero"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff param.Opt[float64] `json:"dragCoeff,omitzero"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel param.Opt[string] `json:"dragModel,omitzero"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr param.Opt[float64] `json:"edr,omitzero"`
	// Integrator error control.
	ErrorControl param.Opt[float64] `json:"errorControl,omitzero"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep param.Opt[bool] `json:"fixedStep,omitzero"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel param.Opt[string] `json:"geopotentialModel,omitzero"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms param.Opt[int64] `json:"iau1980Terms,omitzero"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination param.Opt[string] `json:"idOrbitDetermination,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Integrator Mode.
	IntegratorMode param.Opt[string] `json:"integratorMode,omitzero"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust param.Opt[bool] `json:"inTrackThrust,omitzero"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd param.Opt[time.Time] `json:"lastObEnd,omitzero" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart param.Opt[time.Time] `json:"lastObStart,omitzero" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime param.Opt[time.Time] `json:"leapSecondTime,omitzero" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar param.Opt[bool] `json:"lunarSolar,omitzero"`
	// The mass of the object, in kilograms.
	Mass param.Opt[float64] `json:"mass,omitzero"`
	// The number of observations available for the OD of the object.
	ObsAvailable param.Opt[int64] `json:"obsAvailable,omitzero"`
	// The number of observations accepted for the OD of the object.
	ObsUsed param.Opt[int64] `json:"obsUsed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials param.Opt[string] `json:"partials,omitzero"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree param.Opt[string] `json:"pedigree,omitzero"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX param.Opt[float64] `json:"polarMotionX,omitzero"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY param.Opt[float64] `json:"polarMotionY,omitzero"`
	// One sigma position uncertainty, in kilometers.
	PosUnc param.Opt[float64] `json:"posUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan param.Opt[float64] `json:"recODSpan,omitzero"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc param.Opt[float64] `json:"residualsAcc,omitzero"`
	// Epoch revolution number.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms param.Opt[float64] `json:"rms,omitzero"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg param.Opt[float64] `json:"solarFluxAPAvg,omitzero"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 param.Opt[float64] `json:"solarFluxF10,omitzero"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg param.Opt[float64] `json:"solarFluxF10Avg,omitzero"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress param.Opt[bool] `json:"solarRadPress,omitzero"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff param.Opt[float64] `json:"solarRadPressCoeff,omitzero"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides param.Opt[bool] `json:"solidEarthTides,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea param.Opt[float64] `json:"srpArea,omitzero"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode param.Opt[string] `json:"stepMode,omitzero"`
	// Initial integration step size (seconds).
	StepSize param.Opt[float64] `json:"stepSize,omitzero"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection param.Opt[string] `json:"stepSizeSelection,omitzero"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc param.Opt[float64] `json:"taiUtc,omitzero"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel param.Opt[float64] `json:"thrustAccel,omitzero"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail param.Opt[int64] `json:"tracksAvail,omitzero"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed param.Opt[int64] `json:"tracksUsed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate param.Opt[float64] `json:"ut1Rate,omitzero"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc param.Opt[float64] `json:"ut1Utc,omitzero"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc param.Opt[float64] `json:"velUnc,omitzero"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel param.Opt[float64] `json:"xaccel,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos param.Opt[float64] `json:"xpos,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 param.Opt[float64] `json:"xposAlt1,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 param.Opt[float64] `json:"xposAlt2,omitzero"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 param.Opt[float64] `json:"xvelAlt1,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 param.Opt[float64] `json:"xvelAlt2,omitzero"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel param.Opt[float64] `json:"yaccel,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos param.Opt[float64] `json:"ypos,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 param.Opt[float64] `json:"yposAlt1,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 param.Opt[float64] `json:"yposAlt2,omitzero"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 param.Opt[float64] `json:"yvelAlt1,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 param.Opt[float64] `json:"yvelAlt2,omitzero"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel param.Opt[float64] `json:"zaccel,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos param.Opt[float64] `json:"zpos,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 param.Opt[float64] `json:"zposAlt1,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 param.Opt[float64] `json:"zposAlt2,omitzero"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 param.Opt[float64] `json:"zvelAlt1,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 param.Opt[float64] `json:"zvelAlt2,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov,omitzero"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame,omitzero"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW,omitzero"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW,omitzero"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL observation data types used to build this state vector
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
func (f ManeuverNewParamsPostEventStateVector) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverNewParamsPostEventStateVector) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewParamsPostEventStateVector
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverNewParamsPostEventStateVector](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ManeuverNewParamsPostEventStateVector](
		"CovReferenceFrame", false, "J2000", "UVW",
	)
	apijson.RegisterFieldValidator[ManeuverNewParamsPostEventStateVector](
		"ReferenceFrame", false, "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverNewParamsPreEventElset struct {
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
func (f ManeuverNewParamsPreEventElset) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ManeuverNewParamsPreEventElset) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewParamsPreEventElset
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverNewParamsPreEventElset](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverNewParamsPreEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan param.Opt[float64] `json:"actualODSpan,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame param.Opt[string] `json:"alt1ReferenceFrame,omitzero"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame param.Opt[string] `json:"alt2ReferenceFrame,omitzero"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area param.Opt[float64] `json:"area,omitzero"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot param.Opt[float64] `json:"bDot,omitzero"`
	// Model parameter value for center of mass offset (m).
	CmOffset param.Opt[float64] `json:"cmOffset,omitzero"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod param.Opt[string] `json:"covMethod,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea param.Opt[float64] `json:"dragArea,omitzero"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff param.Opt[float64] `json:"dragCoeff,omitzero"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel param.Opt[string] `json:"dragModel,omitzero"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr param.Opt[float64] `json:"edr,omitzero"`
	// Integrator error control.
	ErrorControl param.Opt[float64] `json:"errorControl,omitzero"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep param.Opt[bool] `json:"fixedStep,omitzero"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel param.Opt[string] `json:"geopotentialModel,omitzero"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms param.Opt[int64] `json:"iau1980Terms,omitzero"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination param.Opt[string] `json:"idOrbitDetermination,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Integrator Mode.
	IntegratorMode param.Opt[string] `json:"integratorMode,omitzero"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust param.Opt[bool] `json:"inTrackThrust,omitzero"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd param.Opt[time.Time] `json:"lastObEnd,omitzero" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart param.Opt[time.Time] `json:"lastObStart,omitzero" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime param.Opt[time.Time] `json:"leapSecondTime,omitzero" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar param.Opt[bool] `json:"lunarSolar,omitzero"`
	// The mass of the object, in kilograms.
	Mass param.Opt[float64] `json:"mass,omitzero"`
	// The number of observations available for the OD of the object.
	ObsAvailable param.Opt[int64] `json:"obsAvailable,omitzero"`
	// The number of observations accepted for the OD of the object.
	ObsUsed param.Opt[int64] `json:"obsUsed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials param.Opt[string] `json:"partials,omitzero"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree param.Opt[string] `json:"pedigree,omitzero"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX param.Opt[float64] `json:"polarMotionX,omitzero"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY param.Opt[float64] `json:"polarMotionY,omitzero"`
	// One sigma position uncertainty, in kilometers.
	PosUnc param.Opt[float64] `json:"posUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan param.Opt[float64] `json:"recODSpan,omitzero"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc param.Opt[float64] `json:"residualsAcc,omitzero"`
	// Epoch revolution number.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms param.Opt[float64] `json:"rms,omitzero"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg param.Opt[float64] `json:"solarFluxAPAvg,omitzero"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 param.Opt[float64] `json:"solarFluxF10,omitzero"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg param.Opt[float64] `json:"solarFluxF10Avg,omitzero"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress param.Opt[bool] `json:"solarRadPress,omitzero"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff param.Opt[float64] `json:"solarRadPressCoeff,omitzero"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides param.Opt[bool] `json:"solidEarthTides,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea param.Opt[float64] `json:"srpArea,omitzero"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode param.Opt[string] `json:"stepMode,omitzero"`
	// Initial integration step size (seconds).
	StepSize param.Opt[float64] `json:"stepSize,omitzero"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection param.Opt[string] `json:"stepSizeSelection,omitzero"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc param.Opt[float64] `json:"taiUtc,omitzero"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel param.Opt[float64] `json:"thrustAccel,omitzero"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail param.Opt[int64] `json:"tracksAvail,omitzero"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed param.Opt[int64] `json:"tracksUsed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate param.Opt[float64] `json:"ut1Rate,omitzero"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc param.Opt[float64] `json:"ut1Utc,omitzero"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc param.Opt[float64] `json:"velUnc,omitzero"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel param.Opt[float64] `json:"xaccel,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos param.Opt[float64] `json:"xpos,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 param.Opt[float64] `json:"xposAlt1,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 param.Opt[float64] `json:"xposAlt2,omitzero"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 param.Opt[float64] `json:"xvelAlt1,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 param.Opt[float64] `json:"xvelAlt2,omitzero"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel param.Opt[float64] `json:"yaccel,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos param.Opt[float64] `json:"ypos,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 param.Opt[float64] `json:"yposAlt1,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 param.Opt[float64] `json:"yposAlt2,omitzero"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 param.Opt[float64] `json:"yvelAlt1,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 param.Opt[float64] `json:"yvelAlt2,omitzero"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel param.Opt[float64] `json:"zaccel,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos param.Opt[float64] `json:"zpos,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 param.Opt[float64] `json:"zposAlt1,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 param.Opt[float64] `json:"zposAlt2,omitzero"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 param.Opt[float64] `json:"zvelAlt1,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 param.Opt[float64] `json:"zvelAlt2,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov,omitzero"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame,omitzero"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW,omitzero"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW,omitzero"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL observation data types used to build this state vector
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
func (f ManeuverNewParamsPreEventStateVector) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverNewParamsPreEventStateVector) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewParamsPreEventStateVector
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverNewParamsPreEventStateVector](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ManeuverNewParamsPreEventStateVector](
		"CovReferenceFrame", false, "J2000", "UVW",
	)
	apijson.RegisterFieldValidator[ManeuverNewParamsPreEventStateVector](
		"ReferenceFrame", false, "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

type ManeuverListParams struct {
	// Maneuver event start time in ISO 8601 UTC with microsecond precision. For
	// maneuvers without start and end times, the start time is considered to be the
	// maneuver event time. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EventStartTime time.Time        `query:"eventStartTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ManeuverListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ManeuverListParams]'s query parameters as `url.Values`.
func (r ManeuverListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManeuverCountParams struct {
	// Maneuver event start time in ISO 8601 UTC with microsecond precision. For
	// maneuvers without start and end times, the start time is considered to be the
	// maneuver event time. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EventStartTime time.Time        `query:"eventStartTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ManeuverCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ManeuverCountParams]'s query parameters as `url.Values`.
func (r ManeuverCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManeuverNewBulkParams struct {
	Body []ManeuverNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ManeuverNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ManeuverNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Model representation of on-orbit object maneuver information for detected,
// possible, and confirmed maneuvers.
//
// The properties ClassificationMarking, DataMode, EventStartTime, Source are
// required.
type ManeuverNewBulkParamsBody struct {
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
	// Maneuver event start time in ISO 8601 UTC with microsecond precision. For
	// maneuvers without start and end times, the start time is considered to be the
	// maneuver event time.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// Optional purpose of the maneuver if known (e.g. North-South Station Keeping,
	// East-West Station Keeping, Longitude Shift, Unknown).
	Characterization param.Opt[string] `json:"characterization,omitzero"`
	// Uncertainty in the characterization or purpose assessment of this maneuver (0 -
	// 1).
	CharacterizationUnc param.Opt[float64] `json:"characterizationUnc,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Difference in mass before and after the maneuver, in kg.
	DeltaMass param.Opt[float64] `json:"deltaMass,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors at the maneuver event time.
	DeltaPos param.Opt[float64] `json:"deltaPos,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaPosU param.Opt[float64] `json:"deltaPosU,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaPosV param.Opt[float64] `json:"deltaPosV,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaPosW param.Opt[float64] `json:"deltaPosW,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors at the maneuver event time.
	DeltaVel param.Opt[float64] `json:"deltaVel,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaVelU param.Opt[float64] `json:"deltaVelU,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaVelV param.Opt[float64] `json:"deltaVelV,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaVelW param.Opt[float64] `json:"deltaVelW,omitzero"`
	// Description and notes of the maneuver.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Maneuver event end time in ISO 8601 UTC with microsecond precision.
	EventEndTime param.Opt[time.Time] `json:"eventEndTime,omitzero" format:"date-time"`
	// Optional source-provided identifier for this maneuver event. In the case where
	// multiple maneuver records are submitted for the same event, this field can be
	// used to tie them together to the same event.
	EventID param.Opt[string] `json:"eventId,omitzero"`
	// Target maneuvering on-orbit object. For the public catalog, the idOnOrbit is
	// typically the satellite number as a string, but may be a UUID for analyst or
	// other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Optional ID of the sensor that detected this maneuver (for example, if detected
	// by passive RF anomalies).
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Uncertainty in the occurrence of this maneuver (0 - 1).
	ManeuverUnc param.Opt[float64] `json:"maneuverUnc,omitzero"`
	// The total number of estimated acceleration points during the maneuver.
	NumAccelPoints param.Opt[int64] `json:"numAccelPoints,omitzero"`
	// Number of observations used to generate the maneuver data.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// Maneuver orbit determination fit data end time in ISO 8601 UTC with microsecond
	// precision.
	OdFitEndTime param.Opt[time.Time] `json:"odFitEndTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Identifier provided by source to indicate the target on-orbit object performing
	// this maneuver. This may be an internal identifier and not necessarily a valid
	// satellite number/ID.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Post-event spacecraft apogee (measured from Earth center), in kilometers.
	PostApogee param.Opt[float64] `json:"postApogee,omitzero"`
	// Estimated area of the object following the maneuver, in meters squared.
	PostArea param.Opt[float64] `json:"postArea,omitzero"`
	// Post-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PostBallisticCoeff param.Opt[float64] `json:"postBallisticCoeff,omitzero"`
	// Post-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PostDriftRate param.Opt[float64] `json:"postDriftRate,omitzero"`
	// Post-event spacecraft eccentricity.
	PostEccentricity param.Opt[float64] `json:"postEccentricity,omitzero"`
	// Optional identifier of the element set for the post-maneuver orbit.
	PostEventIDElset param.Opt[string] `json:"postEventIdElset,omitzero"`
	// Optional identifier of the state vector for the post-maneuver trajectory of the
	// spacecraft.
	PostEventIDStateVector param.Opt[string] `json:"postEventIdStateVector,omitzero"`
	// Post-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PostGeoLongitude param.Opt[float64] `json:"postGeoLongitude,omitzero"`
	// Post-event spacecraft orbital inclination, in degrees. 0-180.
	PostInclination param.Opt[float64] `json:"postInclination,omitzero"`
	// Estimated mass of the object following the maneuver, in kg.
	PostMass param.Opt[float64] `json:"postMass,omitzero"`
	// Post-event spacecraft perigee (measured from Earth center), in kilometers.
	PostPerigee param.Opt[float64] `json:"postPerigee,omitzero"`
	// Post-event spacecraft orbital period, in minutes.
	PostPeriod param.Opt[float64] `json:"postPeriod,omitzero"`
	// Post-event X component of position in ECI space, in km.
	PostPosX param.Opt[float64] `json:"postPosX,omitzero"`
	// Post-event Y component of position in ECI space, in km.
	PostPosY param.Opt[float64] `json:"postPosY,omitzero"`
	// Post-event Z component of position in ECI space, in km.
	PostPosZ param.Opt[float64] `json:"postPosZ,omitzero"`
	// Post-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PostRaan param.Opt[float64] `json:"postRAAN,omitzero"`
	// Post-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PostRadiationPressCoeff param.Opt[float64] `json:"postRadiationPressCoeff,omitzero"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PostSigmaU param.Opt[float64] `json:"postSigmaU,omitzero"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PostSigmaV param.Opt[float64] `json:"postSigmaV,omitzero"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PostSigmaW param.Opt[float64] `json:"postSigmaW,omitzero"`
	// Post-event spacecraft Semi-Major Axis (SMA), in kilometers.
	PostSma param.Opt[float64] `json:"postSMA,omitzero"`
	// Post-event X component of velocity in ECI space, in km/sec.
	PostVelX param.Opt[float64] `json:"postVelX,omitzero"`
	// Post-event Y component of velocity in ECI space, in km/sec.
	PostVelY param.Opt[float64] `json:"postVelY,omitzero"`
	// Post-event Z component of velocity in ECI space, in km/sec.
	PostVelZ param.Opt[float64] `json:"postVelZ,omitzero"`
	// Pre-event spacecraft apogee (measured from Earth center), in kilometers.
	PreApogee param.Opt[float64] `json:"preApogee,omitzero"`
	// Pre-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PreBallisticCoeff param.Opt[float64] `json:"preBallisticCoeff,omitzero"`
	// Pre-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PreDriftRate param.Opt[float64] `json:"preDriftRate,omitzero"`
	// Pre-event spacecraft eccentricity.
	PreEccentricity param.Opt[float64] `json:"preEccentricity,omitzero"`
	// Optional identifier of the element set for the pre-maneuver orbit.
	PreEventIDElset param.Opt[string] `json:"preEventIdElset,omitzero"`
	// Optional identifier of the state vector for the pre-maneuver trajectory of the
	// spacecraft.
	PreEventIDStateVector param.Opt[string] `json:"preEventIdStateVector,omitzero"`
	// Pre-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PreGeoLongitude param.Opt[float64] `json:"preGeoLongitude,omitzero"`
	// Pre-event spacecraft orbital inclination, in degrees. 0-180.
	PreInclination param.Opt[float64] `json:"preInclination,omitzero"`
	// Pre-event spacecraft perigee (measured from Earth center), in kilometers.
	PrePerigee param.Opt[float64] `json:"prePerigee,omitzero"`
	// Pre-event spacecraft orbital period, in minutes.
	PrePeriod param.Opt[float64] `json:"prePeriod,omitzero"`
	// Pre-event X component of position in ECI space, in km.
	PrePosX param.Opt[float64] `json:"prePosX,omitzero"`
	// Pre-event Y component of position in ECI space, in km.
	PrePosY param.Opt[float64] `json:"prePosY,omitzero"`
	// Pre-event Z component of position in ECI space, in km.
	PrePosZ param.Opt[float64] `json:"prePosZ,omitzero"`
	// Pre-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PreRaan param.Opt[float64] `json:"preRAAN,omitzero"`
	// Pre-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PreRadiationPressCoeff param.Opt[float64] `json:"preRadiationPressCoeff,omitzero"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PreSigmaU param.Opt[float64] `json:"preSigmaU,omitzero"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PreSigmaV param.Opt[float64] `json:"preSigmaV,omitzero"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PreSigmaW param.Opt[float64] `json:"preSigmaW,omitzero"`
	// Pre-event spacecraft orbital Semi-Major Axis (SMA), in kilometers.
	PreSma param.Opt[float64] `json:"preSMA,omitzero"`
	// Pre-event X component of velocity in ECI space, in km/sec.
	PreVelX param.Opt[float64] `json:"preVelX,omitzero"`
	// Pre-event Y component of velocity in ECI space, in km/sec.
	PreVelY param.Opt[float64] `json:"preVelY,omitzero"`
	// Pre-event Z component of velocity in ECI space, in km/sec.
	PreVelZ param.Opt[float64] `json:"preVelZ,omitzero"`
	// The time that the report or alert of this maneuver was generated, in ISO 8601
	// UTC format.
	ReportTime param.Opt[time.Time] `json:"reportTime,omitzero" format:"date-time"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Name of the state model used to generate the maneuver data.
	StateModel param.Opt[string] `json:"stateModel,omitzero"`
	// Version of the state model used to generate the maneuver data.
	StateModelVersion param.Opt[float64] `json:"stateModelVersion,omitzero"`
	// Status of this maneuver (CANCELLED, PLANNED, POSSIBLE, REDACTED, VERIFIED).
	//
	// CANCELLED: A previously planned maneuver whose execution was cancelled.
	//
	// PLANNED: A maneuver planned to take place at the eventStartTime.
	//
	// POSSIBLE: A possible maneuver detected by observation of the spacecraft or by
	// evaluation of the spacecraft orbit.
	//
	// REDACTED: A redaction of a reported possible maneuver that has been determined
	// to have not taken place after further observation/evaluation.
	//
	// VERIFIED: A maneuver whose execution has been verified, either by the
	// owner/operator or observation/evaluation.
	Status param.Opt[string] `json:"status,omitzero"`
	// The estimated total active burn time of a maneuver, in seconds. This includes
	// the sum of all burns in numAccelPoints. Not to be confused with the total
	// duration of the maneuver.
	TotalBurnTime param.Opt[float64] `json:"totalBurnTime,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this maneuver was unable to be correlated to a known object.
	// This flag should only be set to true by data providers after an attempt to
	// correlate to an on-orbit object was made and failed. If unable to correlate, the
	// 'origObjectId' field may be populated with an internal data provider specific
	// identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Optional maneuver cross-track/radial/in-track covariance array, in meter and
	// second based units, in the following order: CR_R, CI_R, CI_I, CC_R, CC_I, CC_C,
	// CT_R, CT_I, CT_C, CT_T.
	Cov []float64 `json:"cov,omitzero"`
	// Array of estimated acceleration values, in meters per second squared. Number of
	// elements must match the numAccelPoints.
	MnvrAccels []float64 `json:"mnvrAccels,omitzero"`
	// Array of elapsed times, in seconds from maneuver start time, at which each
	// acceleration point is estimated. Number of elements must match the
	// numAccelPoints.
	MnvrAccelTimes []float64 `json:"mnvrAccelTimes,omitzero"`
	// Array of the 1-sigma uncertainties in estimated accelerations, in meters per
	// second squared. Number of elements must match the numAccelPoints.
	MnvrAccelUncs []float64 `json:"mnvrAccelUncs,omitzero"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PostEventElset ManeuverNewBulkParamsBodyPostEventElset `json:"postEventElset,omitzero"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PostEventStateVector ManeuverNewBulkParamsBodyPostEventStateVector `json:"postEventStateVector,omitzero"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PreEventElset ManeuverNewBulkParamsBodyPreEventElset `json:"preEventElset,omitzero"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PreEventStateVector ManeuverNewBulkParamsBodyPreEventStateVector `json:"preEventStateVector,omitzero"`
	// Optional array of UDL data (elsets, state vectors, etc) UUIDs used to build this
	// maneuver. See the associated sourcedDataTypes array for the specific types of
	// data for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL data types used to build this maneuver (e.g. EO, RADAR,
	// RF, DOA, ELSET, SV). See the associated sourcedData array for the specific UUIDs
	// of data for the positionally corresponding data types in this array (the two
	// arrays must match in size).
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
func (f ManeuverNewBulkParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ManeuverNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverNewBulkParamsBodyPostEventElset struct {
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
func (f ManeuverNewBulkParamsBodyPostEventElset) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverNewBulkParamsBodyPostEventElset) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewBulkParamsBodyPostEventElset
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverNewBulkParamsBodyPostEventElset](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverNewBulkParamsBodyPostEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan param.Opt[float64] `json:"actualODSpan,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame param.Opt[string] `json:"alt1ReferenceFrame,omitzero"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame param.Opt[string] `json:"alt2ReferenceFrame,omitzero"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area param.Opt[float64] `json:"area,omitzero"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot param.Opt[float64] `json:"bDot,omitzero"`
	// Model parameter value for center of mass offset (m).
	CmOffset param.Opt[float64] `json:"cmOffset,omitzero"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod param.Opt[string] `json:"covMethod,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea param.Opt[float64] `json:"dragArea,omitzero"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff param.Opt[float64] `json:"dragCoeff,omitzero"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel param.Opt[string] `json:"dragModel,omitzero"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr param.Opt[float64] `json:"edr,omitzero"`
	// Integrator error control.
	ErrorControl param.Opt[float64] `json:"errorControl,omitzero"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep param.Opt[bool] `json:"fixedStep,omitzero"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel param.Opt[string] `json:"geopotentialModel,omitzero"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms param.Opt[int64] `json:"iau1980Terms,omitzero"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination param.Opt[string] `json:"idOrbitDetermination,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Integrator Mode.
	IntegratorMode param.Opt[string] `json:"integratorMode,omitzero"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust param.Opt[bool] `json:"inTrackThrust,omitzero"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd param.Opt[time.Time] `json:"lastObEnd,omitzero" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart param.Opt[time.Time] `json:"lastObStart,omitzero" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime param.Opt[time.Time] `json:"leapSecondTime,omitzero" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar param.Opt[bool] `json:"lunarSolar,omitzero"`
	// The mass of the object, in kilograms.
	Mass param.Opt[float64] `json:"mass,omitzero"`
	// The number of observations available for the OD of the object.
	ObsAvailable param.Opt[int64] `json:"obsAvailable,omitzero"`
	// The number of observations accepted for the OD of the object.
	ObsUsed param.Opt[int64] `json:"obsUsed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials param.Opt[string] `json:"partials,omitzero"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree param.Opt[string] `json:"pedigree,omitzero"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX param.Opt[float64] `json:"polarMotionX,omitzero"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY param.Opt[float64] `json:"polarMotionY,omitzero"`
	// One sigma position uncertainty, in kilometers.
	PosUnc param.Opt[float64] `json:"posUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan param.Opt[float64] `json:"recODSpan,omitzero"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc param.Opt[float64] `json:"residualsAcc,omitzero"`
	// Epoch revolution number.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms param.Opt[float64] `json:"rms,omitzero"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg param.Opt[float64] `json:"solarFluxAPAvg,omitzero"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 param.Opt[float64] `json:"solarFluxF10,omitzero"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg param.Opt[float64] `json:"solarFluxF10Avg,omitzero"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress param.Opt[bool] `json:"solarRadPress,omitzero"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff param.Opt[float64] `json:"solarRadPressCoeff,omitzero"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides param.Opt[bool] `json:"solidEarthTides,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea param.Opt[float64] `json:"srpArea,omitzero"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode param.Opt[string] `json:"stepMode,omitzero"`
	// Initial integration step size (seconds).
	StepSize param.Opt[float64] `json:"stepSize,omitzero"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection param.Opt[string] `json:"stepSizeSelection,omitzero"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc param.Opt[float64] `json:"taiUtc,omitzero"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel param.Opt[float64] `json:"thrustAccel,omitzero"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail param.Opt[int64] `json:"tracksAvail,omitzero"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed param.Opt[int64] `json:"tracksUsed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate param.Opt[float64] `json:"ut1Rate,omitzero"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc param.Opt[float64] `json:"ut1Utc,omitzero"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc param.Opt[float64] `json:"velUnc,omitzero"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel param.Opt[float64] `json:"xaccel,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos param.Opt[float64] `json:"xpos,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 param.Opt[float64] `json:"xposAlt1,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 param.Opt[float64] `json:"xposAlt2,omitzero"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 param.Opt[float64] `json:"xvelAlt1,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 param.Opt[float64] `json:"xvelAlt2,omitzero"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel param.Opt[float64] `json:"yaccel,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos param.Opt[float64] `json:"ypos,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 param.Opt[float64] `json:"yposAlt1,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 param.Opt[float64] `json:"yposAlt2,omitzero"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 param.Opt[float64] `json:"yvelAlt1,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 param.Opt[float64] `json:"yvelAlt2,omitzero"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel param.Opt[float64] `json:"zaccel,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos param.Opt[float64] `json:"zpos,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 param.Opt[float64] `json:"zposAlt1,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 param.Opt[float64] `json:"zposAlt2,omitzero"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 param.Opt[float64] `json:"zvelAlt1,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 param.Opt[float64] `json:"zvelAlt2,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov,omitzero"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame,omitzero"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW,omitzero"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW,omitzero"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL observation data types used to build this state vector
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
func (f ManeuverNewBulkParamsBodyPostEventStateVector) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverNewBulkParamsBodyPostEventStateVector) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewBulkParamsBodyPostEventStateVector
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverNewBulkParamsBodyPostEventStateVector](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ManeuverNewBulkParamsBodyPostEventStateVector](
		"CovReferenceFrame", false, "J2000", "UVW",
	)
	apijson.RegisterFieldValidator[ManeuverNewBulkParamsBodyPostEventStateVector](
		"ReferenceFrame", false, "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverNewBulkParamsBodyPreEventElset struct {
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
func (f ManeuverNewBulkParamsBodyPreEventElset) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverNewBulkParamsBodyPreEventElset) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewBulkParamsBodyPreEventElset
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverNewBulkParamsBodyPreEventElset](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverNewBulkParamsBodyPreEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan param.Opt[float64] `json:"actualODSpan,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame param.Opt[string] `json:"alt1ReferenceFrame,omitzero"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame param.Opt[string] `json:"alt2ReferenceFrame,omitzero"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area param.Opt[float64] `json:"area,omitzero"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot param.Opt[float64] `json:"bDot,omitzero"`
	// Model parameter value for center of mass offset (m).
	CmOffset param.Opt[float64] `json:"cmOffset,omitzero"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod param.Opt[string] `json:"covMethod,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea param.Opt[float64] `json:"dragArea,omitzero"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff param.Opt[float64] `json:"dragCoeff,omitzero"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel param.Opt[string] `json:"dragModel,omitzero"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr param.Opt[float64] `json:"edr,omitzero"`
	// Integrator error control.
	ErrorControl param.Opt[float64] `json:"errorControl,omitzero"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep param.Opt[bool] `json:"fixedStep,omitzero"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel param.Opt[string] `json:"geopotentialModel,omitzero"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms param.Opt[int64] `json:"iau1980Terms,omitzero"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination param.Opt[string] `json:"idOrbitDetermination,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Integrator Mode.
	IntegratorMode param.Opt[string] `json:"integratorMode,omitzero"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust param.Opt[bool] `json:"inTrackThrust,omitzero"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd param.Opt[time.Time] `json:"lastObEnd,omitzero" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart param.Opt[time.Time] `json:"lastObStart,omitzero" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime param.Opt[time.Time] `json:"leapSecondTime,omitzero" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar param.Opt[bool] `json:"lunarSolar,omitzero"`
	// The mass of the object, in kilograms.
	Mass param.Opt[float64] `json:"mass,omitzero"`
	// The number of observations available for the OD of the object.
	ObsAvailable param.Opt[int64] `json:"obsAvailable,omitzero"`
	// The number of observations accepted for the OD of the object.
	ObsUsed param.Opt[int64] `json:"obsUsed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials param.Opt[string] `json:"partials,omitzero"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree param.Opt[string] `json:"pedigree,omitzero"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX param.Opt[float64] `json:"polarMotionX,omitzero"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY param.Opt[float64] `json:"polarMotionY,omitzero"`
	// One sigma position uncertainty, in kilometers.
	PosUnc param.Opt[float64] `json:"posUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan param.Opt[float64] `json:"recODSpan,omitzero"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc param.Opt[float64] `json:"residualsAcc,omitzero"`
	// Epoch revolution number.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms param.Opt[float64] `json:"rms,omitzero"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg param.Opt[float64] `json:"solarFluxAPAvg,omitzero"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 param.Opt[float64] `json:"solarFluxF10,omitzero"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg param.Opt[float64] `json:"solarFluxF10Avg,omitzero"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress param.Opt[bool] `json:"solarRadPress,omitzero"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff param.Opt[float64] `json:"solarRadPressCoeff,omitzero"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides param.Opt[bool] `json:"solidEarthTides,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea param.Opt[float64] `json:"srpArea,omitzero"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode param.Opt[string] `json:"stepMode,omitzero"`
	// Initial integration step size (seconds).
	StepSize param.Opt[float64] `json:"stepSize,omitzero"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection param.Opt[string] `json:"stepSizeSelection,omitzero"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc param.Opt[float64] `json:"taiUtc,omitzero"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel param.Opt[float64] `json:"thrustAccel,omitzero"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail param.Opt[int64] `json:"tracksAvail,omitzero"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed param.Opt[int64] `json:"tracksUsed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate param.Opt[float64] `json:"ut1Rate,omitzero"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc param.Opt[float64] `json:"ut1Utc,omitzero"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc param.Opt[float64] `json:"velUnc,omitzero"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel param.Opt[float64] `json:"xaccel,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos param.Opt[float64] `json:"xpos,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 param.Opt[float64] `json:"xposAlt1,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 param.Opt[float64] `json:"xposAlt2,omitzero"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 param.Opt[float64] `json:"xvelAlt1,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 param.Opt[float64] `json:"xvelAlt2,omitzero"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel param.Opt[float64] `json:"yaccel,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos param.Opt[float64] `json:"ypos,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 param.Opt[float64] `json:"yposAlt1,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 param.Opt[float64] `json:"yposAlt2,omitzero"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 param.Opt[float64] `json:"yvelAlt1,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 param.Opt[float64] `json:"yvelAlt2,omitzero"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel param.Opt[float64] `json:"zaccel,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos param.Opt[float64] `json:"zpos,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 param.Opt[float64] `json:"zposAlt1,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 param.Opt[float64] `json:"zposAlt2,omitzero"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 param.Opt[float64] `json:"zvelAlt1,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 param.Opt[float64] `json:"zvelAlt2,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov,omitzero"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame,omitzero"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW,omitzero"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW,omitzero"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL observation data types used to build this state vector
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
func (f ManeuverNewBulkParamsBodyPreEventStateVector) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverNewBulkParamsBodyPreEventStateVector) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverNewBulkParamsBodyPreEventStateVector
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverNewBulkParamsBodyPreEventStateVector](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ManeuverNewBulkParamsBodyPreEventStateVector](
		"CovReferenceFrame", false, "J2000", "UVW",
	)
	apijson.RegisterFieldValidator[ManeuverNewBulkParamsBodyPreEventStateVector](
		"ReferenceFrame", false, "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

type ManeuverGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ManeuverGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ManeuverGetParams]'s query parameters as `url.Values`.
func (r ManeuverGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManeuverTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Maneuver event start time in ISO 8601 UTC with microsecond precision. For
	// maneuvers without start and end times, the start time is considered to be the
	// maneuver event time. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	EventStartTime time.Time        `query:"eventStartTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ManeuverTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ManeuverTupleParams]'s query parameters as `url.Values`.
func (r ManeuverTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManeuverUnvalidatedPublishParams struct {
	Body []ManeuverUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ManeuverUnvalidatedPublishParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ManeuverUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Model representation of on-orbit object maneuver information for detected,
// possible, and confirmed maneuvers.
//
// The properties ClassificationMarking, DataMode, EventStartTime, Source are
// required.
type ManeuverUnvalidatedPublishParamsBody struct {
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
	// Maneuver event start time in ISO 8601 UTC with microsecond precision. For
	// maneuvers without start and end times, the start time is considered to be the
	// maneuver event time.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// Optional purpose of the maneuver if known (e.g. North-South Station Keeping,
	// East-West Station Keeping, Longitude Shift, Unknown).
	Characterization param.Opt[string] `json:"characterization,omitzero"`
	// Uncertainty in the characterization or purpose assessment of this maneuver (0 -
	// 1).
	CharacterizationUnc param.Opt[float64] `json:"characterizationUnc,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Difference in mass before and after the maneuver, in kg.
	DeltaMass param.Opt[float64] `json:"deltaMass,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors at the maneuver event time.
	DeltaPos param.Opt[float64] `json:"deltaPos,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaPosU param.Opt[float64] `json:"deltaPosU,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaPosV param.Opt[float64] `json:"deltaPosV,omitzero"`
	// Magnitude, in km, of the difference in the pre- and post-maneuver position
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaPosW param.Opt[float64] `json:"deltaPosW,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors at the maneuver event time.
	DeltaVel param.Opt[float64] `json:"deltaVel,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'U' unit vector at the maneuver
	// event time.
	DeltaVelU param.Opt[float64] `json:"deltaVelU,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'V' unit vector at the maneuver
	// event time.
	DeltaVelV param.Opt[float64] `json:"deltaVelV,omitzero"`
	// Magnitude, in km/sec, of the difference in the pre- and post-maneuver velocity
	// vectors in the direction of the pre-maneuver 'W' unit vector at the maneuver
	// event time.
	DeltaVelW param.Opt[float64] `json:"deltaVelW,omitzero"`
	// Description and notes of the maneuver.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Maneuver event end time in ISO 8601 UTC with microsecond precision.
	EventEndTime param.Opt[time.Time] `json:"eventEndTime,omitzero" format:"date-time"`
	// Optional source-provided identifier for this maneuver event. In the case where
	// multiple maneuver records are submitted for the same event, this field can be
	// used to tie them together to the same event.
	EventID param.Opt[string] `json:"eventId,omitzero"`
	// Target maneuvering on-orbit object. For the public catalog, the idOnOrbit is
	// typically the satellite number as a string, but may be a UUID for analyst or
	// other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Optional ID of the sensor that detected this maneuver (for example, if detected
	// by passive RF anomalies).
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Uncertainty in the occurrence of this maneuver (0 - 1).
	ManeuverUnc param.Opt[float64] `json:"maneuverUnc,omitzero"`
	// The total number of estimated acceleration points during the maneuver.
	NumAccelPoints param.Opt[int64] `json:"numAccelPoints,omitzero"`
	// Number of observations used to generate the maneuver data.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// Maneuver orbit determination fit data end time in ISO 8601 UTC with microsecond
	// precision.
	OdFitEndTime param.Opt[time.Time] `json:"odFitEndTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Identifier provided by source to indicate the target on-orbit object performing
	// this maneuver. This may be an internal identifier and not necessarily a valid
	// satellite number/ID.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by source to indicate the sensor identifier used to
	// detect this event. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Post-event spacecraft apogee (measured from Earth center), in kilometers.
	PostApogee param.Opt[float64] `json:"postApogee,omitzero"`
	// Estimated area of the object following the maneuver, in meters squared.
	PostArea param.Opt[float64] `json:"postArea,omitzero"`
	// Post-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PostBallisticCoeff param.Opt[float64] `json:"postBallisticCoeff,omitzero"`
	// Post-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PostDriftRate param.Opt[float64] `json:"postDriftRate,omitzero"`
	// Post-event spacecraft eccentricity.
	PostEccentricity param.Opt[float64] `json:"postEccentricity,omitzero"`
	// Optional identifier of the element set for the post-maneuver orbit.
	PostEventIDElset param.Opt[string] `json:"postEventIdElset,omitzero"`
	// Optional identifier of the state vector for the post-maneuver trajectory of the
	// spacecraft.
	PostEventIDStateVector param.Opt[string] `json:"postEventIdStateVector,omitzero"`
	// Post-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PostGeoLongitude param.Opt[float64] `json:"postGeoLongitude,omitzero"`
	// Post-event spacecraft orbital inclination, in degrees. 0-180.
	PostInclination param.Opt[float64] `json:"postInclination,omitzero"`
	// Estimated mass of the object following the maneuver, in kg.
	PostMass param.Opt[float64] `json:"postMass,omitzero"`
	// Post-event spacecraft perigee (measured from Earth center), in kilometers.
	PostPerigee param.Opt[float64] `json:"postPerigee,omitzero"`
	// Post-event spacecraft orbital period, in minutes.
	PostPeriod param.Opt[float64] `json:"postPeriod,omitzero"`
	// Post-event X component of position in ECI space, in km.
	PostPosX param.Opt[float64] `json:"postPosX,omitzero"`
	// Post-event Y component of position in ECI space, in km.
	PostPosY param.Opt[float64] `json:"postPosY,omitzero"`
	// Post-event Z component of position in ECI space, in km.
	PostPosZ param.Opt[float64] `json:"postPosZ,omitzero"`
	// Post-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PostRaan param.Opt[float64] `json:"postRAAN,omitzero"`
	// Post-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PostRadiationPressCoeff param.Opt[float64] `json:"postRadiationPressCoeff,omitzero"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PostSigmaU param.Opt[float64] `json:"postSigmaU,omitzero"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PostSigmaV param.Opt[float64] `json:"postSigmaV,omitzero"`
	// Post-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PostSigmaW param.Opt[float64] `json:"postSigmaW,omitzero"`
	// Post-event spacecraft Semi-Major Axis (SMA), in kilometers.
	PostSma param.Opt[float64] `json:"postSMA,omitzero"`
	// Post-event X component of velocity in ECI space, in km/sec.
	PostVelX param.Opt[float64] `json:"postVelX,omitzero"`
	// Post-event Y component of velocity in ECI space, in km/sec.
	PostVelY param.Opt[float64] `json:"postVelY,omitzero"`
	// Post-event Z component of velocity in ECI space, in km/sec.
	PostVelZ param.Opt[float64] `json:"postVelZ,omitzero"`
	// Pre-event spacecraft apogee (measured from Earth center), in kilometers.
	PreApogee param.Opt[float64] `json:"preApogee,omitzero"`
	// Pre-event ballistic coefficient. The units of the ballistic coefficient vary
	// depending on provider. Users should consult the data provider to verify the
	// units of the ballistic coefficient.
	PreBallisticCoeff param.Opt[float64] `json:"preBallisticCoeff,omitzero"`
	// Pre-event GEO drift rate of the spacecraft, in degrees per day. Negative values
	// indicate westward drift.
	PreDriftRate param.Opt[float64] `json:"preDriftRate,omitzero"`
	// Pre-event spacecraft eccentricity.
	PreEccentricity param.Opt[float64] `json:"preEccentricity,omitzero"`
	// Optional identifier of the element set for the pre-maneuver orbit.
	PreEventIDElset param.Opt[string] `json:"preEventIdElset,omitzero"`
	// Optional identifier of the state vector for the pre-maneuver trajectory of the
	// spacecraft.
	PreEventIDStateVector param.Opt[string] `json:"preEventIdStateVector,omitzero"`
	// Pre-event spacecraft WGS-84 GEO belt longitude, represented as -180 to 180
	// degrees (negative values west of Prime Meridian).
	PreGeoLongitude param.Opt[float64] `json:"preGeoLongitude,omitzero"`
	// Pre-event spacecraft orbital inclination, in degrees. 0-180.
	PreInclination param.Opt[float64] `json:"preInclination,omitzero"`
	// Pre-event spacecraft perigee (measured from Earth center), in kilometers.
	PrePerigee param.Opt[float64] `json:"prePerigee,omitzero"`
	// Pre-event spacecraft orbital period, in minutes.
	PrePeriod param.Opt[float64] `json:"prePeriod,omitzero"`
	// Pre-event X component of position in ECI space, in km.
	PrePosX param.Opt[float64] `json:"prePosX,omitzero"`
	// Pre-event Y component of position in ECI space, in km.
	PrePosY param.Opt[float64] `json:"prePosY,omitzero"`
	// Pre-event Z component of position in ECI space, in km.
	PrePosZ param.Opt[float64] `json:"prePosZ,omitzero"`
	// Pre-event spacecraft Right Ascension of the Ascending Node (RAAN), in degrees.
	PreRaan param.Opt[float64] `json:"preRAAN,omitzero"`
	// Pre-event radiation pressure coefficient. The units of the radiation pressure
	// coefficient vary depending on provider. Users should consult the data provider
	// to verify the units of the radiation pressure coefficient.
	PreRadiationPressCoeff param.Opt[float64] `json:"preRadiationPressCoeff,omitzero"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'U'
	// unit vector direction.
	PreSigmaU param.Opt[float64] `json:"preSigmaU,omitzero"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'V'
	// unit vector direction.
	PreSigmaV param.Opt[float64] `json:"preSigmaV,omitzero"`
	// Pre-event standard deviation, in kilometers, of spacecraft position in the 'W'
	// unit vector direction.
	PreSigmaW param.Opt[float64] `json:"preSigmaW,omitzero"`
	// Pre-event spacecraft orbital Semi-Major Axis (SMA), in kilometers.
	PreSma param.Opt[float64] `json:"preSMA,omitzero"`
	// Pre-event X component of velocity in ECI space, in km/sec.
	PreVelX param.Opt[float64] `json:"preVelX,omitzero"`
	// Pre-event Y component of velocity in ECI space, in km/sec.
	PreVelY param.Opt[float64] `json:"preVelY,omitzero"`
	// Pre-event Z component of velocity in ECI space, in km/sec.
	PreVelZ param.Opt[float64] `json:"preVelZ,omitzero"`
	// The time that the report or alert of this maneuver was generated, in ISO 8601
	// UTC format.
	ReportTime param.Opt[time.Time] `json:"reportTime,omitzero" format:"date-time"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Name of the state model used to generate the maneuver data.
	StateModel param.Opt[string] `json:"stateModel,omitzero"`
	// Version of the state model used to generate the maneuver data.
	StateModelVersion param.Opt[float64] `json:"stateModelVersion,omitzero"`
	// Status of this maneuver (CANCELLED, PLANNED, POSSIBLE, REDACTED, VERIFIED).
	//
	// CANCELLED: A previously planned maneuver whose execution was cancelled.
	//
	// PLANNED: A maneuver planned to take place at the eventStartTime.
	//
	// POSSIBLE: A possible maneuver detected by observation of the spacecraft or by
	// evaluation of the spacecraft orbit.
	//
	// REDACTED: A redaction of a reported possible maneuver that has been determined
	// to have not taken place after further observation/evaluation.
	//
	// VERIFIED: A maneuver whose execution has been verified, either by the
	// owner/operator or observation/evaluation.
	Status param.Opt[string] `json:"status,omitzero"`
	// The estimated total active burn time of a maneuver, in seconds. This includes
	// the sum of all burns in numAccelPoints. Not to be confused with the total
	// duration of the maneuver.
	TotalBurnTime param.Opt[float64] `json:"totalBurnTime,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this maneuver was unable to be correlated to a known object.
	// This flag should only be set to true by data providers after an attempt to
	// correlate to an on-orbit object was made and failed. If unable to correlate, the
	// 'origObjectId' field may be populated with an internal data provider specific
	// identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Optional maneuver cross-track/radial/in-track covariance array, in meter and
	// second based units, in the following order: CR_R, CI_R, CI_I, CC_R, CC_I, CC_C,
	// CT_R, CT_I, CT_C, CT_T.
	Cov []float64 `json:"cov,omitzero"`
	// Array of estimated acceleration values, in meters per second squared. Number of
	// elements must match the numAccelPoints.
	MnvrAccels []float64 `json:"mnvrAccels,omitzero"`
	// Array of elapsed times, in seconds from maneuver start time, at which each
	// acceleration point is estimated. Number of elements must match the
	// numAccelPoints.
	MnvrAccelTimes []float64 `json:"mnvrAccelTimes,omitzero"`
	// Array of the 1-sigma uncertainties in estimated accelerations, in meters per
	// second squared. Number of elements must match the numAccelPoints.
	MnvrAccelUncs []float64 `json:"mnvrAccelUncs,omitzero"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PostEventElset ManeuverUnvalidatedPublishParamsBodyPostEventElset `json:"postEventElset,omitzero"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PostEventStateVector ManeuverUnvalidatedPublishParamsBodyPostEventStateVector `json:"postEventStateVector,omitzero"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	PreEventElset ManeuverUnvalidatedPublishParamsBodyPreEventElset `json:"preEventElset,omitzero"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	PreEventStateVector ManeuverUnvalidatedPublishParamsBodyPreEventStateVector `json:"preEventStateVector,omitzero"`
	// Optional array of UDL data (elsets, state vectors, etc) UUIDs used to build this
	// maneuver. See the associated sourcedDataTypes array for the specific types of
	// data for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL data types used to build this maneuver (e.g. EO, RADAR,
	// RF, DOA, ELSET, SV). See the associated sourcedData array for the specific UUIDs
	// of data for the positionally corresponding data types in this array (the two
	// arrays must match in size).
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
func (f ManeuverUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverUnvalidatedPublishParamsBodyPostEventElset struct {
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
func (f ManeuverUnvalidatedPublishParamsBodyPostEventElset) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverUnvalidatedPublishParamsBodyPostEventElset) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverUnvalidatedPublishParamsBodyPostEventElset
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverUnvalidatedPublishParamsBodyPostEventElset](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverUnvalidatedPublishParamsBodyPostEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan param.Opt[float64] `json:"actualODSpan,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame param.Opt[string] `json:"alt1ReferenceFrame,omitzero"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame param.Opt[string] `json:"alt2ReferenceFrame,omitzero"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area param.Opt[float64] `json:"area,omitzero"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot param.Opt[float64] `json:"bDot,omitzero"`
	// Model parameter value for center of mass offset (m).
	CmOffset param.Opt[float64] `json:"cmOffset,omitzero"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod param.Opt[string] `json:"covMethod,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea param.Opt[float64] `json:"dragArea,omitzero"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff param.Opt[float64] `json:"dragCoeff,omitzero"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel param.Opt[string] `json:"dragModel,omitzero"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr param.Opt[float64] `json:"edr,omitzero"`
	// Integrator error control.
	ErrorControl param.Opt[float64] `json:"errorControl,omitzero"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep param.Opt[bool] `json:"fixedStep,omitzero"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel param.Opt[string] `json:"geopotentialModel,omitzero"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms param.Opt[int64] `json:"iau1980Terms,omitzero"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination param.Opt[string] `json:"idOrbitDetermination,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Integrator Mode.
	IntegratorMode param.Opt[string] `json:"integratorMode,omitzero"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust param.Opt[bool] `json:"inTrackThrust,omitzero"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd param.Opt[time.Time] `json:"lastObEnd,omitzero" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart param.Opt[time.Time] `json:"lastObStart,omitzero" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime param.Opt[time.Time] `json:"leapSecondTime,omitzero" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar param.Opt[bool] `json:"lunarSolar,omitzero"`
	// The mass of the object, in kilograms.
	Mass param.Opt[float64] `json:"mass,omitzero"`
	// The number of observations available for the OD of the object.
	ObsAvailable param.Opt[int64] `json:"obsAvailable,omitzero"`
	// The number of observations accepted for the OD of the object.
	ObsUsed param.Opt[int64] `json:"obsUsed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials param.Opt[string] `json:"partials,omitzero"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree param.Opt[string] `json:"pedigree,omitzero"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX param.Opt[float64] `json:"polarMotionX,omitzero"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY param.Opt[float64] `json:"polarMotionY,omitzero"`
	// One sigma position uncertainty, in kilometers.
	PosUnc param.Opt[float64] `json:"posUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan param.Opt[float64] `json:"recODSpan,omitzero"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc param.Opt[float64] `json:"residualsAcc,omitzero"`
	// Epoch revolution number.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms param.Opt[float64] `json:"rms,omitzero"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg param.Opt[float64] `json:"solarFluxAPAvg,omitzero"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 param.Opt[float64] `json:"solarFluxF10,omitzero"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg param.Opt[float64] `json:"solarFluxF10Avg,omitzero"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress param.Opt[bool] `json:"solarRadPress,omitzero"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff param.Opt[float64] `json:"solarRadPressCoeff,omitzero"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides param.Opt[bool] `json:"solidEarthTides,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea param.Opt[float64] `json:"srpArea,omitzero"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode param.Opt[string] `json:"stepMode,omitzero"`
	// Initial integration step size (seconds).
	StepSize param.Opt[float64] `json:"stepSize,omitzero"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection param.Opt[string] `json:"stepSizeSelection,omitzero"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc param.Opt[float64] `json:"taiUtc,omitzero"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel param.Opt[float64] `json:"thrustAccel,omitzero"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail param.Opt[int64] `json:"tracksAvail,omitzero"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed param.Opt[int64] `json:"tracksUsed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate param.Opt[float64] `json:"ut1Rate,omitzero"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc param.Opt[float64] `json:"ut1Utc,omitzero"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc param.Opt[float64] `json:"velUnc,omitzero"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel param.Opt[float64] `json:"xaccel,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos param.Opt[float64] `json:"xpos,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 param.Opt[float64] `json:"xposAlt1,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 param.Opt[float64] `json:"xposAlt2,omitzero"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 param.Opt[float64] `json:"xvelAlt1,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 param.Opt[float64] `json:"xvelAlt2,omitzero"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel param.Opt[float64] `json:"yaccel,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos param.Opt[float64] `json:"ypos,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 param.Opt[float64] `json:"yposAlt1,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 param.Opt[float64] `json:"yposAlt2,omitzero"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 param.Opt[float64] `json:"yvelAlt1,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 param.Opt[float64] `json:"yvelAlt2,omitzero"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel param.Opt[float64] `json:"zaccel,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos param.Opt[float64] `json:"zpos,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 param.Opt[float64] `json:"zposAlt1,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 param.Opt[float64] `json:"zposAlt2,omitzero"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 param.Opt[float64] `json:"zvelAlt1,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 param.Opt[float64] `json:"zvelAlt2,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov,omitzero"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame,omitzero"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW,omitzero"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW,omitzero"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL observation data types used to build this state vector
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
func (f ManeuverUnvalidatedPublishParamsBodyPostEventStateVector) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverUnvalidatedPublishParamsBodyPostEventStateVector) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverUnvalidatedPublishParamsBodyPostEventStateVector
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverUnvalidatedPublishParamsBodyPostEventStateVector](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ManeuverUnvalidatedPublishParamsBodyPostEventStateVector](
		"CovReferenceFrame", false, "J2000", "UVW",
	)
	apijson.RegisterFieldValidator[ManeuverUnvalidatedPublishParamsBodyPostEventStateVector](
		"ReferenceFrame", false, "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverUnvalidatedPublishParamsBodyPreEventElset struct {
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
func (f ManeuverUnvalidatedPublishParamsBodyPreEventElset) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverUnvalidatedPublishParamsBodyPreEventElset) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverUnvalidatedPublishParamsBodyPreEventElset
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverUnvalidatedPublishParamsBodyPreEventElset](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type ManeuverUnvalidatedPublishParamsBodyPreEventStateVector struct {
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
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan param.Opt[float64] `json:"actualODSpan,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame param.Opt[string] `json:"alt1ReferenceFrame,omitzero"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame param.Opt[string] `json:"alt2ReferenceFrame,omitzero"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area param.Opt[float64] `json:"area,omitzero"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot param.Opt[float64] `json:"bDot,omitzero"`
	// Model parameter value for center of mass offset (m).
	CmOffset param.Opt[float64] `json:"cmOffset,omitzero"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod param.Opt[string] `json:"covMethod,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea param.Opt[float64] `json:"dragArea,omitzero"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff param.Opt[float64] `json:"dragCoeff,omitzero"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel param.Opt[string] `json:"dragModel,omitzero"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr param.Opt[float64] `json:"edr,omitzero"`
	// Integrator error control.
	ErrorControl param.Opt[float64] `json:"errorControl,omitzero"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep param.Opt[bool] `json:"fixedStep,omitzero"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel param.Opt[string] `json:"geopotentialModel,omitzero"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms param.Opt[int64] `json:"iau1980Terms,omitzero"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination param.Opt[string] `json:"idOrbitDetermination,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Integrator Mode.
	IntegratorMode param.Opt[string] `json:"integratorMode,omitzero"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust param.Opt[bool] `json:"inTrackThrust,omitzero"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd param.Opt[time.Time] `json:"lastObEnd,omitzero" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart param.Opt[time.Time] `json:"lastObStart,omitzero" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime param.Opt[time.Time] `json:"leapSecondTime,omitzero" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar param.Opt[bool] `json:"lunarSolar,omitzero"`
	// The mass of the object, in kilograms.
	Mass param.Opt[float64] `json:"mass,omitzero"`
	// The number of observations available for the OD of the object.
	ObsAvailable param.Opt[int64] `json:"obsAvailable,omitzero"`
	// The number of observations accepted for the OD of the object.
	ObsUsed param.Opt[int64] `json:"obsUsed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials param.Opt[string] `json:"partials,omitzero"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree param.Opt[string] `json:"pedigree,omitzero"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX param.Opt[float64] `json:"polarMotionX,omitzero"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY param.Opt[float64] `json:"polarMotionY,omitzero"`
	// One sigma position uncertainty, in kilometers.
	PosUnc param.Opt[float64] `json:"posUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan param.Opt[float64] `json:"recODSpan,omitzero"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc param.Opt[float64] `json:"residualsAcc,omitzero"`
	// Epoch revolution number.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms param.Opt[float64] `json:"rms,omitzero"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg param.Opt[float64] `json:"solarFluxAPAvg,omitzero"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 param.Opt[float64] `json:"solarFluxF10,omitzero"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg param.Opt[float64] `json:"solarFluxF10Avg,omitzero"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress param.Opt[bool] `json:"solarRadPress,omitzero"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff param.Opt[float64] `json:"solarRadPressCoeff,omitzero"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides param.Opt[bool] `json:"solidEarthTides,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea param.Opt[float64] `json:"srpArea,omitzero"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode param.Opt[string] `json:"stepMode,omitzero"`
	// Initial integration step size (seconds).
	StepSize param.Opt[float64] `json:"stepSize,omitzero"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection param.Opt[string] `json:"stepSizeSelection,omitzero"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc param.Opt[float64] `json:"taiUtc,omitzero"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel param.Opt[float64] `json:"thrustAccel,omitzero"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail param.Opt[int64] `json:"tracksAvail,omitzero"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed param.Opt[int64] `json:"tracksUsed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate param.Opt[float64] `json:"ut1Rate,omitzero"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc param.Opt[float64] `json:"ut1Utc,omitzero"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc param.Opt[float64] `json:"velUnc,omitzero"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel param.Opt[float64] `json:"xaccel,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos param.Opt[float64] `json:"xpos,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 param.Opt[float64] `json:"xposAlt1,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 param.Opt[float64] `json:"xposAlt2,omitzero"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 param.Opt[float64] `json:"xvelAlt1,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 param.Opt[float64] `json:"xvelAlt2,omitzero"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel param.Opt[float64] `json:"yaccel,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos param.Opt[float64] `json:"ypos,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 param.Opt[float64] `json:"yposAlt1,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 param.Opt[float64] `json:"yposAlt2,omitzero"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 param.Opt[float64] `json:"yvelAlt1,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 param.Opt[float64] `json:"yvelAlt2,omitzero"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel param.Opt[float64] `json:"zaccel,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos param.Opt[float64] `json:"zpos,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 param.Opt[float64] `json:"zposAlt1,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 param.Opt[float64] `json:"zposAlt2,omitzero"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 param.Opt[float64] `json:"zvelAlt1,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 param.Opt[float64] `json:"zvelAlt2,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov,omitzero"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW".
	CovReferenceFrame string `json:"covReferenceFrame,omitzero"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame string `json:"referenceFrame,omitzero"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW,omitzero"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW,omitzero"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL observation data types used to build this state vector
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
func (f ManeuverUnvalidatedPublishParamsBodyPreEventStateVector) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ManeuverUnvalidatedPublishParamsBodyPreEventStateVector) MarshalJSON() (data []byte, err error) {
	type shadow ManeuverUnvalidatedPublishParamsBodyPreEventStateVector
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ManeuverUnvalidatedPublishParamsBodyPreEventStateVector](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ManeuverUnvalidatedPublishParamsBodyPreEventStateVector](
		"CovReferenceFrame", false, "J2000", "UVW",
	)
	apijson.RegisterFieldValidator[ManeuverUnvalidatedPublishParamsBodyPreEventStateVector](
		"ReferenceFrame", false, "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}
