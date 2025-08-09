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

// CollectRequestService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectRequestService] method instead.
type CollectRequestService struct {
	Options []option.RequestOption
	History CollectRequestHistoryService
}

// NewCollectRequestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectRequestService(opts ...option.RequestOption) (r CollectRequestService) {
	r = CollectRequestService{}
	r.Options = opts
	r.History = NewCollectRequestHistoryService(opts...)
	return
}

// Service operation to take a single CollectRequest as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *CollectRequestService) New(ctx context.Context, body CollectRequestNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/collectrequest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single CollectRequest by its unique ID passed as a
// path parameter.
func (r *CollectRequestService) Get(ctx context.Context, id string, query CollectRequestGetParams, opts ...option.RequestOption) (res *shared.CollectRequestFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/collectrequest/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *CollectRequestService) List(ctx context.Context, query CollectRequestListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[CollectRequestAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/collectrequest"
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
func (r *CollectRequestService) ListAutoPaging(ctx context.Context, query CollectRequestListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[CollectRequestAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *CollectRequestService) Count(ctx context.Context, query CollectRequestCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/collectrequest/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// CollectRequest as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *CollectRequestService) NewBulk(ctx context.Context, body CollectRequestNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/collectrequest/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *CollectRequestService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *CollectRequestQueryHelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/collectrequest/queryhelp"
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
func (r *CollectRequestService) Tuple(ctx context.Context, query CollectRequestTupleParams, opts ...option.RequestOption) (res *[]shared.CollectRequestFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/collectrequest/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a list of CollectRequest as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *CollectRequestService) UnvalidatedPublish(ctx context.Context, body CollectRequestUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-collectrequest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Collect Requests support several types of individual requests, or
// planned/scheduled tasks on sensors and/or orbital objects. Options are provided
// to accomodate most common sensor contact and collection applications, including
// single sensor-object tasking, search operations, and TT&C support. Multiple
// requests originating from a plan or schedule may be associated to a sensor plan
// if desired.
type CollectRequestAbridged struct {
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
	DataMode CollectRequestAbridgedDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start time or earliest time of the collect or contact request window, in ISO
	// 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The type of this collect or contact request (DIRECTED SEARCH, DWELL, OBJECT,
	// POL, RATE TRACK, SEARCH, SOI, STARE, TTC, VOLUME SEARCH, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Height above WGS-84 ellipsoid (HAE), in kilometers. If an accompanying stopAlt
	// is provided, then alt value can be assumed to be the starting altitude of a
	// volume definition.
	Alt float64 `json:"alt"`
	// The argument of perigee is the angle, in degrees, formed between the perigee and
	// the ascending node.
	ArgOfPerigee float64 `json:"argOfPerigee"`
	// The expected or directed azimuth angle, in degrees, for search or target
	// acquisition.
	Az float64 `json:"az"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The customer for this request.
	Customer string `json:"customer"`
	// The expected or directed declination angle, in degrees, for search or target
	// acquisition.
	Dec float64 `json:"dec"`
	// The duration of the collect request, in seconds. If both duration and endTime
	// are provided, the endTime is assumed to take precedence.
	Duration int64 `json:"duration"`
	// The dwell ID associated with this request. A dwell ID is dwell point specific
	// and a DWELL request consist of many dwell point requests.
	DwellID string `json:"dwellId"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle.
	Eccentricity float64 `json:"eccentricity"`
	// The expected or directed elevation angle, in degrees, for search or target
	// acquisition.
	El float64 `json:"el"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	Elset CollectRequestAbridgedElset `json:"elset"`
	// The end time of the collect or contact request window, in ISO 8601 UTC format.
	// If no endTime or duration is provided it is assumed the request is either
	// ongoing or that the request is for a specified number of tracks (numTracks). If
	// both duration and endTime are provided, the endTime is assumed to take
	// precedence.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Epoch time, in ISO 8601 UTC format, of the orbital elements.
	Epoch time.Time `json:"epoch" format:"date-time"`
	// ID of the UDL Ephemeris Set of the object associated with this request.
	EsID string `json:"esId"`
	// The extent of the azimuth angle, in degrees, from center azimuth to define a
	// spatial volume.
	ExtentAz float64 `json:"extentAz"`
	// The extent of the elevation angle, in degrees, from center elevation to define a
	// spatial volume.
	ExtentEl float64 `json:"extentEl"`
	// The extent of the range, in km, from center range to define a spatial volume.
	ExtentRange float64 `json:"extentRange"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// For optical sensors, the frame rate of the camera, in Hz.
	FrameRate float64 `json:"frameRate"`
	// The estimated or expected emission frequency of the target, in MHz.
	Freq float64 `json:"freq"`
	// The maximum frequency of interest, in MHz.
	FreqMax float64 `json:"freqMax"`
	// The minimum frequency of interest, in MHz. If only minimum frequency is provided
	// it is assumed to be minimum reportable frequency.
	FreqMin float64 `json:"freqMin"`
	// ID of the UDL Elset of the object associated with this request.
	IDElset string `json:"idElset"`
	// ID of the UDL Manifold Elset of the object associated with this request. A
	// Manifold Elset provides theoretical Keplerian orbital elements belonging to an
	// object of interest's manifold describing a possible/theoretical orbit for an
	// object of interest for tasking purposes.
	IDManifold string `json:"idManifold"`
	// Unique identifier of the target on-orbit object for this request.
	IDOnOrbit string `json:"idOnOrbit"`
	// The unique ID of the collect request record from which this request originated.
	// This may be used for cases of sensor-to-sensor tasking, such as tip/cue
	// operations.
	IDParentReq string `json:"idParentReq"`
	// Unique identifier of the parent plan or schedule associated with this request.
	// If null, this request is assumed not associated with a plan or schedule.
	IDPlan string `json:"idPlan"`
	// Unique identifier of the requested/scheduled/planned sensor associated with this
	// request. If both idSensor and origSensorId are null then the request is assumed
	// to be a general request for observations or contact on an object, if specified,
	// or an area/volume. In this case, the requester may specify a desired obType.
	IDSensor string `json:"idSensor"`
	// ID of the UDL State Vector of the object or central vector associated with this
	// request.
	IDStateVector string `json:"idStateVector"`
	// The angle, in degrees, between the equator and the orbit plane when looking from
	// the center of the Earth. Inclination ranges from 0-180 degrees, with 0-90
	// representing posigrade orbits and 90-180 representing retrograde orbits.
	Inclination float64 `json:"inclination"`
	// For optical sensors, the integration time per camera frame, in milliseconds.
	IntegrationTime float64 `json:"integrationTime"`
	// Inter-Range Operations Number. Four-digit identifier used to schedule and
	// identify AFSCN contact support for booster, launch, and on-orbit operations.
	Iron int64 `json:"iron"`
	// The target object irradiance value.
	Irradiance float64 `json:"irradiance"`
	// WGS-84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator). If an accompanying stopLat is provided, then the lat value can be
	// assumed to be the starting latitude of a volume definition.
	Lat float64 `json:"lat"`
	// WGS-84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian). If an accompanying stopLon is provided, then lon value can be assumed
	// to be the starting longitude of a volume definition.
	Lon float64 `json:"lon"`
	// The timestamp of the external message from which this request originated, if
	// applicable, in ISO8601 UTC format with millisecond precision.
	MsgCreateDate time.Time `json:"msgCreateDate" format:"date-time"`
	// The type of external message from which this request originated.
	MsgType string `json:"msgType"`
	// Notes or comments associated with this request.
	Notes string `json:"notes"`
	// For optical sensors, the requested number of frames to capture at each sensor
	// step.
	NumFrames int64 `json:"numFrames"`
	// The number of requested observations on the target.
	NumObs int64 `json:"numObs"`
	// The number of requested tracks on the target. If numTracks is not provided it is
	// assumed to indicate all possible observations every pass over the request
	// duration or within the request start/end window.
	NumTracks int64 `json:"numTracks"`
	// Optional type of observation (EO, IR, RADAR, RF-ACTIVE, RF-PASSIVE, OTHER)
	// requested. This field may correspond to a request of a specific sensor, or to a
	// general non sensor specific request.
	ObType string `json:"obType"`
	// The orbit regime of the target (GEO, HEO, LAUNCH, LEO, MEO, OTHER).
	OrbitRegime string `json:"orbitRegime"`
	// The magnitude of rotation, in degrees, between the xAngle direction and locally
	// defined equinoctial plane. A positive value indicates clockwise rotation about
	// the sensor boresight vector.
	OrientAngle float64 `json:"orientAngle"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the data source to indicate the target object of
	// this request. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the source to indicate the sensor identifier
	// requested/scheduled/planned for this request. This may be an internal identifier
	// and not necessarily a valid sensor ID. If both idSensor and origSensorId are
	// null then the request is assumed to be a general request for observations or
	// contact on an object, if specified, or an area/volume. In this case, the
	// requester may specify a desired obType.
	OrigSensorID string `json:"origSensorId"`
	// Index number (integer) for records within a collection plan or schedule.
	PlanIndex int64 `json:"planIndex"`
	// The RF polarization (H, LHC, RHC, V).
	Polarization string `json:"polarization"`
	// The priority of the collect request (EMERGENCY, FLASH, IMMEDIATE, PRIORITY,
	// ROUTINE).
	Priority string `json:"priority"`
	// The expected or directed right ascension angle, in degrees, for search or target
	// acquisition.
	Ra float64 `json:"ra"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node.
	Raan float64 `json:"raan"`
	// The expected acquisition range or defined center range, in km.
	Range float64 `json:"range"`
	// The Radar Cross-Section of the target, in m^2.
	Rcs float64 `json:"rcs"`
	// The maximum Radar Cross-Section of the target, in m^2.
	RcsMax float64 `json:"rcsMax"`
	// The minimum Radar Cross-Section of the target, in m^2. If only minimum RCS is
	// provided it is assumed to be minimum reportable RCS.
	RcsMin float64 `json:"rcsMin"`
	// The fraction of solar energy reflected from target.
	Reflectance float64 `json:"reflectance"`
	// Satellite/catalog number of the target on-orbit object for this request.
	SatNo int64 `json:"satNo"`
	// Pre-coordinated code, direction, or configuration to be executed by the sensor
	// or site for this collect or contact.
	Scenario string `json:"scenario"`
	// The average of the periapsis and apoapsis distances, in kilometers. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies.
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// The spectral model used for the irradiance calculation.
	SpectralModel string `json:"spectralModel"`
	// The maximum inclination, in degrees, to be used in search operations.
	SrchInc float64 `json:"srchInc"`
	// The search pattern to be executed for this request (e.g. PICKET-FENCE, SCAN,
	// etc.).
	SrchPattern string `json:"srchPattern"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	StateVector CollectRequestAbridgedStateVector `json:"stateVector"`
	// The stopping HAE WGS-84 height above ellipsoid (HAE), of a volume definition, in
	// kilometers. The stopAlt value is only meaningful if a (starting) alt value is
	// provided.
	StopAlt float64 `json:"stopAlt"`
	// The stopping WGS-84 latitude of a volume definition, in degrees. -90 to 90
	// degrees (negative values south of equator). The stopLat value is only meaningful
	// if a (starting) lat value is provided.
	StopLat float64 `json:"stopLat"`
	// The stopping WGS-84 longitude of a volume definition, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian). The stopLon value is only
	// meaningful if a (starting) lon value is provided.
	StopLon float64 `json:"stopLon"`
	// The (SSN) tasking suffix (A-Z) associated with this request. The suffix defines
	// the amount of observational data and the frequency of collection. Note that
	// suffix definitions are sensor type specific.
	Suffix string `json:"suffix"`
	// The minimum object (diameter) size, in meters, to be reported.
	TargetSize float64 `json:"targetSize"`
	// The (SSN) tasking category (1-5) associated with this request. The tasking
	// category defines the priority of gathering and transmitting the requested
	// observational data. Note that category definitions are sensor type specific.
	TaskCategory int64 `json:"taskCategory"`
	// The tasking group to which the target object is assigned.
	TaskGroup string `json:"taskGroup"`
	// Task ID associated with this request. A task ID may be associated with a single
	// collect request or may be used to tie together the sub-requests of a full
	// collect, for example a DWELL consisting of many dwell points.
	TaskID string `json:"taskId"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// The true anomaly defines the angular position, in degrees, of the object on it's
	// orbital path as measured from the orbit focal point at epoch. The true anomaly
	// is referenced from perigee.
	TrueAnomoly float64 `json:"trueAnomoly"`
	// Boolean indicating that this collect request is UCT follow-up.
	UctFollowUp bool `json:"uctFollowUp"`
	// The estimated or expected visual magnitude of the target, in Magnitudes (M).
	VisMag float64 `json:"visMag"`
	// The maximum estimated or expected visual magnitude of the target, in Magnitudes
	// (M).
	VisMagMax float64 `json:"visMagMax"`
	// The minimum estimated or expected visual magnitude of the target, in Magnitudes
	// (M). If only minimum vismag is provided it is assumed to be minimum reportable
	// vismag.
	VisMagMin float64 `json:"visMagMin"`
	// The angular distance, in degrees, in the sensor-x direction from scan center
	// defined by the central vector. The specification of xAngle and yAngle defines a
	// rectangle of width 2*xAngle and height 2*yAngle centered about the central
	// vector.
	XAngle float64 `json:"xAngle"`
	// The angular distance, in degrees, in the sensor-y direction from scan center
	// defined by the central vector. The specification of xAngle and yAngle defines a
	// rectangle of width 2*xAngle and height 2*yAngle centered about the central
	// vector.
	YAngle float64 `json:"yAngle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		Alt                   respjson.Field
		ArgOfPerigee          respjson.Field
		Az                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Customer              respjson.Field
		Dec                   respjson.Field
		Duration              respjson.Field
		DwellID               respjson.Field
		Eccentricity          respjson.Field
		El                    respjson.Field
		Elset                 respjson.Field
		EndTime               respjson.Field
		Epoch                 respjson.Field
		EsID                  respjson.Field
		ExtentAz              respjson.Field
		ExtentEl              respjson.Field
		ExtentRange           respjson.Field
		ExternalID            respjson.Field
		FrameRate             respjson.Field
		Freq                  respjson.Field
		FreqMax               respjson.Field
		FreqMin               respjson.Field
		IDElset               respjson.Field
		IDManifold            respjson.Field
		IDOnOrbit             respjson.Field
		IDParentReq           respjson.Field
		IDPlan                respjson.Field
		IDSensor              respjson.Field
		IDStateVector         respjson.Field
		Inclination           respjson.Field
		IntegrationTime       respjson.Field
		Iron                  respjson.Field
		Irradiance            respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		MsgCreateDate         respjson.Field
		MsgType               respjson.Field
		Notes                 respjson.Field
		NumFrames             respjson.Field
		NumObs                respjson.Field
		NumTracks             respjson.Field
		ObType                respjson.Field
		OrbitRegime           respjson.Field
		OrientAngle           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		PlanIndex             respjson.Field
		Polarization          respjson.Field
		Priority              respjson.Field
		Ra                    respjson.Field
		Raan                  respjson.Field
		Range                 respjson.Field
		Rcs                   respjson.Field
		RcsMax                respjson.Field
		RcsMin                respjson.Field
		Reflectance           respjson.Field
		SatNo                 respjson.Field
		Scenario              respjson.Field
		SemiMajorAxis         respjson.Field
		SpectralModel         respjson.Field
		SrchInc               respjson.Field
		SrchPattern           respjson.Field
		StateVector           respjson.Field
		StopAlt               respjson.Field
		StopLat               respjson.Field
		StopLon               respjson.Field
		Suffix                respjson.Field
		TargetSize            respjson.Field
		TaskCategory          respjson.Field
		TaskGroup             respjson.Field
		TaskID                respjson.Field
		TransactionID         respjson.Field
		TrueAnomoly           respjson.Field
		UctFollowUp           respjson.Field
		VisMag                respjson.Field
		VisMagMax             respjson.Field
		VisMagMin             respjson.Field
		XAngle                respjson.Field
		YAngle                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectRequestAbridged) RawJSON() string { return r.JSON.raw }
func (r *CollectRequestAbridged) UnmarshalJSON(data []byte) error {
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
type CollectRequestAbridgedDataMode string

const (
	CollectRequestAbridgedDataModeReal      CollectRequestAbridgedDataMode = "REAL"
	CollectRequestAbridgedDataModeTest      CollectRequestAbridgedDataMode = "TEST"
	CollectRequestAbridgedDataModeSimulated CollectRequestAbridgedDataMode = "SIMULATED"
	CollectRequestAbridgedDataModeExercise  CollectRequestAbridgedDataMode = "EXERCISE"
)

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
type CollectRequestAbridgedElset struct {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Epoch                 respjson.Field
		Source                respjson.Field
		Agom                  respjson.Field
		Algorithm             respjson.Field
		Apogee                respjson.Field
		ArgOfPerigee          respjson.Field
		BallisticCoeff        respjson.Field
		BStar                 respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		Eccentricity          respjson.Field
		EphemType             respjson.Field
		IDElset               respjson.Field
		IDOnOrbit             respjson.Field
		IDOrbitDetermination  respjson.Field
		Inclination           respjson.Field
		Line1                 respjson.Field
		Line2                 respjson.Field
		MeanAnomaly           respjson.Field
		MeanMotion            respjson.Field
		MeanMotionDDot        respjson.Field
		MeanMotionDot         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		Perigee               respjson.Field
		Period                respjson.Field
		Raan                  respjson.Field
		RevNo                 respjson.Field
		SatNo                 respjson.Field
		SemiMajorAxis         respjson.Field
		SourceDl              respjson.Field
		TransactionID         respjson.Field
		Uct                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectRequestAbridgedElset) RawJSON() string { return r.JSON.raw }
func (r *CollectRequestAbridgedElset) UnmarshalJSON(data []byte) error {
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
type CollectRequestAbridgedStateVector struct {
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
	// Any of "J2000", "UVW", "EFG/TDR", "TEME", "GCRF".
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Epoch                 respjson.Field
		Source                respjson.Field
		ActualOdSpan          respjson.Field
		Algorithm             respjson.Field
		Alt1ReferenceFrame    respjson.Field
		Alt2ReferenceFrame    respjson.Field
		Area                  respjson.Field
		BDot                  respjson.Field
		CmOffset              respjson.Field
		Cov                   respjson.Field
		CovMethod             respjson.Field
		CovReferenceFrame     respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		DragArea              respjson.Field
		DragCoeff             respjson.Field
		DragModel             respjson.Field
		Edr                   respjson.Field
		EqCov                 respjson.Field
		ErrorControl          respjson.Field
		FixedStep             respjson.Field
		GeopotentialModel     respjson.Field
		Iau1980Terms          respjson.Field
		IDOnOrbit             respjson.Field
		IDOrbitDetermination  respjson.Field
		IDStateVector         respjson.Field
		IntegratorMode        respjson.Field
		InTrackThrust         respjson.Field
		LastObEnd             respjson.Field
		LastObStart           respjson.Field
		LeapSecondTime        respjson.Field
		LunarSolar            respjson.Field
		Mass                  respjson.Field
		ObsAvailable          respjson.Field
		ObsUsed               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		Partials              respjson.Field
		Pedigree              respjson.Field
		PolarMotionX          respjson.Field
		PolarMotionY          respjson.Field
		PosUnc                respjson.Field
		RecOdSpan             respjson.Field
		ReferenceFrame        respjson.Field
		ResidualsAcc          respjson.Field
		RevNo                 respjson.Field
		Rms                   respjson.Field
		SatNo                 respjson.Field
		SigmaPosUvw           respjson.Field
		SigmaVelUvw           respjson.Field
		SolarFluxApAvg        respjson.Field
		SolarFluxF10          respjson.Field
		SolarFluxF10Avg       respjson.Field
		SolarRadPress         respjson.Field
		SolarRadPressCoeff    respjson.Field
		SolidEarthTides       respjson.Field
		SourceDl              respjson.Field
		SrpArea               respjson.Field
		StepMode              respjson.Field
		StepSize              respjson.Field
		StepSizeSelection     respjson.Field
		TaiUtc                respjson.Field
		ThrustAccel           respjson.Field
		TracksAvail           respjson.Field
		TracksUsed            respjson.Field
		TransactionID         respjson.Field
		Uct                   respjson.Field
		Ut1Rate               respjson.Field
		Ut1Utc                respjson.Field
		VelUnc                respjson.Field
		Xaccel                respjson.Field
		Xpos                  respjson.Field
		XposAlt1              respjson.Field
		XposAlt2              respjson.Field
		Xvel                  respjson.Field
		XvelAlt1              respjson.Field
		XvelAlt2              respjson.Field
		Yaccel                respjson.Field
		Ypos                  respjson.Field
		YposAlt1              respjson.Field
		YposAlt2              respjson.Field
		Yvel                  respjson.Field
		YvelAlt1              respjson.Field
		YvelAlt2              respjson.Field
		Zaccel                respjson.Field
		Zpos                  respjson.Field
		ZposAlt1              respjson.Field
		ZposAlt2              respjson.Field
		Zvel                  respjson.Field
		ZvelAlt1              respjson.Field
		ZvelAlt2              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectRequestAbridgedStateVector) RawJSON() string { return r.JSON.raw }
func (r *CollectRequestAbridgedStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectRequestQueryHelpResponse struct {
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
func (r CollectRequestQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectRequestQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectRequestNewParams struct {
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
	DataMode CollectRequestNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start time or earliest time of the collect or contact request window, in ISO
	// 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The type of this collect or contact request (DIRECTED SEARCH, DWELL, OBJECT,
	// POL, RATE TRACK, SEARCH, SOI, STARE, TTC, VOLUME SEARCH, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Height above WGS-84 ellipsoid (HAE), in kilometers. If an accompanying stopAlt
	// is provided, then alt value can be assumed to be the starting altitude of a
	// volume definition.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// The argument of perigee is the angle, in degrees, formed between the perigee and
	// the ascending node.
	ArgOfPerigee param.Opt[float64] `json:"argOfPerigee,omitzero"`
	// The expected or directed azimuth angle, in degrees, for search or target
	// acquisition.
	Az param.Opt[float64] `json:"az,omitzero"`
	// The customer for this request.
	Customer param.Opt[string] `json:"customer,omitzero"`
	// The expected or directed declination angle, in degrees, for search or target
	// acquisition.
	Dec param.Opt[float64] `json:"dec,omitzero"`
	// The duration of the collect request, in seconds. If both duration and endTime
	// are provided, the endTime is assumed to take precedence.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// The dwell ID associated with this request. A dwell ID is dwell point specific
	// and a DWELL request consist of many dwell point requests.
	DwellID param.Opt[string] `json:"dwellId,omitzero"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle.
	Eccentricity param.Opt[float64] `json:"eccentricity,omitzero"`
	// The expected or directed elevation angle, in degrees, for search or target
	// acquisition.
	El param.Opt[float64] `json:"el,omitzero"`
	// The end time of the collect or contact request window, in ISO 8601 UTC format.
	// If no endTime or duration is provided it is assumed the request is either
	// ongoing or that the request is for a specified number of tracks (numTracks). If
	// both duration and endTime are provided, the endTime is assumed to take
	// precedence.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Epoch time, in ISO 8601 UTC format, of the orbital elements.
	Epoch param.Opt[time.Time] `json:"epoch,omitzero" format:"date-time"`
	// ID of the UDL Ephemeris Set of the object associated with this request.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// The extent of the azimuth angle, in degrees, from center azimuth to define a
	// spatial volume.
	ExtentAz param.Opt[float64] `json:"extentAz,omitzero"`
	// The extent of the elevation angle, in degrees, from center elevation to define a
	// spatial volume.
	ExtentEl param.Opt[float64] `json:"extentEl,omitzero"`
	// The extent of the range, in km, from center range to define a spatial volume.
	ExtentRange param.Opt[float64] `json:"extentRange,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// For optical sensors, the frame rate of the camera, in Hz.
	FrameRate param.Opt[float64] `json:"frameRate,omitzero"`
	// The estimated or expected emission frequency of the target, in MHz.
	Freq param.Opt[float64] `json:"freq,omitzero"`
	// The maximum frequency of interest, in MHz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// The minimum frequency of interest, in MHz. If only minimum frequency is provided
	// it is assumed to be minimum reportable frequency.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// ID of the UDL Elset of the object associated with this request.
	IDElset param.Opt[string] `json:"idElset,omitzero"`
	// ID of the UDL Manifold Elset of the object associated with this request. A
	// Manifold Elset provides theoretical Keplerian orbital elements belonging to an
	// object of interest's manifold describing a possible/theoretical orbit for an
	// object of interest for tasking purposes.
	IDManifold param.Opt[string] `json:"idManifold,omitzero"`
	// The unique ID of the collect request record from which this request originated.
	// This may be used for cases of sensor-to-sensor tasking, such as tip/cue
	// operations.
	IDParentReq param.Opt[string] `json:"idParentReq,omitzero"`
	// Unique identifier of the parent plan or schedule associated with this request.
	// If null, this request is assumed not associated with a plan or schedule.
	IDPlan param.Opt[string] `json:"idPlan,omitzero"`
	// Unique identifier of the requested/scheduled/planned sensor associated with this
	// request. If both idSensor and origSensorId are null then the request is assumed
	// to be a general request for observations or contact on an object, if specified,
	// or an area/volume. In this case, the requester may specify a desired obType.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// ID of the UDL State Vector of the object or central vector associated with this
	// request.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// The angle, in degrees, between the equator and the orbit plane when looking from
	// the center of the Earth. Inclination ranges from 0-180 degrees, with 0-90
	// representing posigrade orbits and 90-180 representing retrograde orbits.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// For optical sensors, the integration time per camera frame, in milliseconds.
	IntegrationTime param.Opt[float64] `json:"integrationTime,omitzero"`
	// Inter-Range Operations Number. Four-digit identifier used to schedule and
	// identify AFSCN contact support for booster, launch, and on-orbit operations.
	Iron param.Opt[int64] `json:"iron,omitzero"`
	// The target object irradiance value.
	Irradiance param.Opt[float64] `json:"irradiance,omitzero"`
	// WGS-84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator). If an accompanying stopLat is provided, then the lat value can be
	// assumed to be the starting latitude of a volume definition.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian). If an accompanying stopLon is provided, then lon value can be assumed
	// to be the starting longitude of a volume definition.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The timestamp of the external message from which this request originated, if
	// applicable, in ISO8601 UTC format with millisecond precision.
	MsgCreateDate param.Opt[time.Time] `json:"msgCreateDate,omitzero" format:"date-time"`
	// The type of external message from which this request originated.
	MsgType param.Opt[string] `json:"msgType,omitzero"`
	// Notes or comments associated with this request.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// For optical sensors, the requested number of frames to capture at each sensor
	// step.
	NumFrames param.Opt[int64] `json:"numFrames,omitzero"`
	// The number of requested observations on the target.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// The number of requested tracks on the target. If numTracks is not provided it is
	// assumed to indicate all possible observations every pass over the request
	// duration or within the request start/end window.
	NumTracks param.Opt[int64] `json:"numTracks,omitzero"`
	// Optional type of observation (EO, IR, RADAR, RF-ACTIVE, RF-PASSIVE, OTHER)
	// requested. This field may correspond to a request of a specific sensor, or to a
	// general non sensor specific request.
	ObType param.Opt[string] `json:"obType,omitzero"`
	// The orbit regime of the target (GEO, HEO, LAUNCH, LEO, MEO, OTHER).
	OrbitRegime param.Opt[string] `json:"orbitRegime,omitzero"`
	// The magnitude of rotation, in degrees, between the xAngle direction and locally
	// defined equinoctial plane. A positive value indicates clockwise rotation about
	// the sensor boresight vector.
	OrientAngle param.Opt[float64] `json:"orientAngle,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the data source to indicate the target object of
	// this request. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the source to indicate the sensor identifier
	// requested/scheduled/planned for this request. This may be an internal identifier
	// and not necessarily a valid sensor ID. If both idSensor and origSensorId are
	// null then the request is assumed to be a general request for observations or
	// contact on an object, if specified, or an area/volume. In this case, the
	// requester may specify a desired obType.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Index number (integer) for records within a collection plan or schedule.
	PlanIndex param.Opt[int64] `json:"planIndex,omitzero"`
	// The RF polarization (H, LHC, RHC, V).
	Polarization param.Opt[string] `json:"polarization,omitzero"`
	// The priority of the collect request (EMERGENCY, FLASH, IMMEDIATE, PRIORITY,
	// ROUTINE).
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The expected or directed right ascension angle, in degrees, for search or target
	// acquisition.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node.
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// The expected acquisition range or defined center range, in km.
	Range param.Opt[float64] `json:"range,omitzero"`
	// The Radar Cross-Section of the target, in m^2.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// The maximum Radar Cross-Section of the target, in m^2.
	RcsMax param.Opt[float64] `json:"rcsMax,omitzero"`
	// The minimum Radar Cross-Section of the target, in m^2. If only minimum RCS is
	// provided it is assumed to be minimum reportable RCS.
	RcsMin param.Opt[float64] `json:"rcsMin,omitzero"`
	// The fraction of solar energy reflected from target.
	Reflectance param.Opt[float64] `json:"reflectance,omitzero"`
	// Satellite/catalog number of the target on-orbit object for this request.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Pre-coordinated code, direction, or configuration to be executed by the sensor
	// or site for this collect or contact.
	Scenario param.Opt[string] `json:"scenario,omitzero"`
	// The average of the periapsis and apoapsis distances, in kilometers. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	// The spectral model used for the irradiance calculation.
	SpectralModel param.Opt[string] `json:"spectralModel,omitzero"`
	// The maximum inclination, in degrees, to be used in search operations.
	SrchInc param.Opt[float64] `json:"srchInc,omitzero"`
	// The search pattern to be executed for this request (e.g. PICKET-FENCE, SCAN,
	// etc.).
	SrchPattern param.Opt[string] `json:"srchPattern,omitzero"`
	// The stopping HAE WGS-84 height above ellipsoid (HAE), of a volume definition, in
	// kilometers. The stopAlt value is only meaningful if a (starting) alt value is
	// provided.
	StopAlt param.Opt[float64] `json:"stopAlt,omitzero"`
	// The stopping WGS-84 latitude of a volume definition, in degrees. -90 to 90
	// degrees (negative values south of equator). The stopLat value is only meaningful
	// if a (starting) lat value is provided.
	StopLat param.Opt[float64] `json:"stopLat,omitzero"`
	// The stopping WGS-84 longitude of a volume definition, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian). The stopLon value is only
	// meaningful if a (starting) lon value is provided.
	StopLon param.Opt[float64] `json:"stopLon,omitzero"`
	// The (SSN) tasking suffix (A-Z) associated with this request. The suffix defines
	// the amount of observational data and the frequency of collection. Note that
	// suffix definitions are sensor type specific.
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// The minimum object (diameter) size, in meters, to be reported.
	TargetSize param.Opt[float64] `json:"targetSize,omitzero"`
	// The (SSN) tasking category (1-5) associated with this request. The tasking
	// category defines the priority of gathering and transmitting the requested
	// observational data. Note that category definitions are sensor type specific.
	TaskCategory param.Opt[int64] `json:"taskCategory,omitzero"`
	// The tasking group to which the target object is assigned.
	TaskGroup param.Opt[string] `json:"taskGroup,omitzero"`
	// Task ID associated with this request. A task ID may be associated with a single
	// collect request or may be used to tie together the sub-requests of a full
	// collect, for example a DWELL consisting of many dwell points.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The true anomaly defines the angular position, in degrees, of the object on it's
	// orbital path as measured from the orbit focal point at epoch. The true anomaly
	// is referenced from perigee.
	TrueAnomoly param.Opt[float64] `json:"trueAnomoly,omitzero"`
	// Boolean indicating that this collect request is UCT follow-up.
	UctFollowUp param.Opt[bool] `json:"uctFollowUp,omitzero"`
	// The estimated or expected visual magnitude of the target, in Magnitudes (M).
	VisMag param.Opt[float64] `json:"visMag,omitzero"`
	// The maximum estimated or expected visual magnitude of the target, in Magnitudes
	// (M).
	VisMagMax param.Opt[float64] `json:"visMagMax,omitzero"`
	// The minimum estimated or expected visual magnitude of the target, in Magnitudes
	// (M). If only minimum vismag is provided it is assumed to be minimum reportable
	// vismag.
	VisMagMin param.Opt[float64] `json:"visMagMin,omitzero"`
	// The angular distance, in degrees, in the sensor-x direction from scan center
	// defined by the central vector. The specification of xAngle and yAngle defines a
	// rectangle of width 2*xAngle and height 2*yAngle centered about the central
	// vector.
	XAngle param.Opt[float64] `json:"xAngle,omitzero"`
	// The angular distance, in degrees, in the sensor-y direction from scan center
	// defined by the central vector. The specification of xAngle and yAngle defines a
	// rectangle of width 2*xAngle and height 2*yAngle centered about the central
	// vector.
	YAngle param.Opt[float64] `json:"yAngle,omitzero"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	Elset CollectRequestNewParamsElset `json:"elset,omitzero"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	StateVector CollectRequestNewParamsStateVector `json:"stateVector,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r CollectRequestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectRequestNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectRequestNewParams) UnmarshalJSON(data []byte) error {
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
type CollectRequestNewParamsDataMode string

const (
	CollectRequestNewParamsDataModeReal      CollectRequestNewParamsDataMode = "REAL"
	CollectRequestNewParamsDataModeTest      CollectRequestNewParamsDataMode = "TEST"
	CollectRequestNewParamsDataModeSimulated CollectRequestNewParamsDataMode = "SIMULATED"
	CollectRequestNewParamsDataModeExercise  CollectRequestNewParamsDataMode = "EXERCISE"
)

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type CollectRequestNewParamsElset struct {
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

func (r CollectRequestNewParamsElset) MarshalJSON() (data []byte, err error) {
	type shadow CollectRequestNewParamsElset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectRequestNewParamsElset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectRequestNewParamsElset](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
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
type CollectRequestNewParamsStateVector struct {
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
	// Any of "J2000", "UVW", "EFG/TDR", "TEME", "GCRF".
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

func (r CollectRequestNewParamsStateVector) MarshalJSON() (data []byte, err error) {
	type shadow CollectRequestNewParamsStateVector
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectRequestNewParamsStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectRequestNewParamsStateVector](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[CollectRequestNewParamsStateVector](
		"covReferenceFrame", "J2000", "UVW", "EFG/TDR", "TEME", "GCRF",
	)
	apijson.RegisterFieldValidator[CollectRequestNewParamsStateVector](
		"referenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

type CollectRequestGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectRequestGetParams]'s query parameters as
// `url.Values`.
func (r CollectRequestGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectRequestListParams struct {
	// The start time or earliest time of the collect or contact request window, in ISO
	// 8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectRequestListParams]'s query parameters as
// `url.Values`.
func (r CollectRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectRequestCountParams struct {
	// The start time or earliest time of the collect or contact request window, in ISO
	// 8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectRequestCountParams]'s query parameters as
// `url.Values`.
func (r CollectRequestCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectRequestNewBulkParams struct {
	Body []CollectRequestNewBulkParamsBody
	paramObj
}

func (r CollectRequestNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *CollectRequestNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Collect Requests support several types of individual requests, or
// planned/scheduled tasks on sensors and/or orbital objects. Options are provided
// to accomodate most common sensor contact and collection applications, including
// single sensor-object tasking, search operations, and TT&C support. Multiple
// requests originating from a plan or schedule may be associated to a sensor plan
// if desired.
//
// The properties ClassificationMarking, DataMode, Source, StartTime, Type are
// required.
type CollectRequestNewBulkParamsBody struct {
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
	// The start time or earliest time of the collect or contact request window, in ISO
	// 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The type of this collect or contact request (DIRECTED SEARCH, DWELL, OBJECT,
	// POL, RATE TRACK, SEARCH, SOI, STARE, TTC, VOLUME SEARCH, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Height above WGS-84 ellipsoid (HAE), in kilometers. If an accompanying stopAlt
	// is provided, then alt value can be assumed to be the starting altitude of a
	// volume definition.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// The argument of perigee is the angle, in degrees, formed between the perigee and
	// the ascending node.
	ArgOfPerigee param.Opt[float64] `json:"argOfPerigee,omitzero"`
	// The expected or directed azimuth angle, in degrees, for search or target
	// acquisition.
	Az param.Opt[float64] `json:"az,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The customer for this request.
	Customer param.Opt[string] `json:"customer,omitzero"`
	// The expected or directed declination angle, in degrees, for search or target
	// acquisition.
	Dec param.Opt[float64] `json:"dec,omitzero"`
	// The duration of the collect request, in seconds. If both duration and endTime
	// are provided, the endTime is assumed to take precedence.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// The dwell ID associated with this request. A dwell ID is dwell point specific
	// and a DWELL request consist of many dwell point requests.
	DwellID param.Opt[string] `json:"dwellId,omitzero"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle.
	Eccentricity param.Opt[float64] `json:"eccentricity,omitzero"`
	// The expected or directed elevation angle, in degrees, for search or target
	// acquisition.
	El param.Opt[float64] `json:"el,omitzero"`
	// The end time of the collect or contact request window, in ISO 8601 UTC format.
	// If no endTime or duration is provided it is assumed the request is either
	// ongoing or that the request is for a specified number of tracks (numTracks). If
	// both duration and endTime are provided, the endTime is assumed to take
	// precedence.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Epoch time, in ISO 8601 UTC format, of the orbital elements.
	Epoch param.Opt[time.Time] `json:"epoch,omitzero" format:"date-time"`
	// ID of the UDL Ephemeris Set of the object associated with this request.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// The extent of the azimuth angle, in degrees, from center azimuth to define a
	// spatial volume.
	ExtentAz param.Opt[float64] `json:"extentAz,omitzero"`
	// The extent of the elevation angle, in degrees, from center elevation to define a
	// spatial volume.
	ExtentEl param.Opt[float64] `json:"extentEl,omitzero"`
	// The extent of the range, in km, from center range to define a spatial volume.
	ExtentRange param.Opt[float64] `json:"extentRange,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// For optical sensors, the frame rate of the camera, in Hz.
	FrameRate param.Opt[float64] `json:"frameRate,omitzero"`
	// The estimated or expected emission frequency of the target, in MHz.
	Freq param.Opt[float64] `json:"freq,omitzero"`
	// The maximum frequency of interest, in MHz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// The minimum frequency of interest, in MHz. If only minimum frequency is provided
	// it is assumed to be minimum reportable frequency.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// ID of the UDL Elset of the object associated with this request.
	IDElset param.Opt[string] `json:"idElset,omitzero"`
	// ID of the UDL Manifold Elset of the object associated with this request. A
	// Manifold Elset provides theoretical Keplerian orbital elements belonging to an
	// object of interest's manifold describing a possible/theoretical orbit for an
	// object of interest for tasking purposes.
	IDManifold param.Opt[string] `json:"idManifold,omitzero"`
	// Unique identifier of the target on-orbit object for this request.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// The unique ID of the collect request record from which this request originated.
	// This may be used for cases of sensor-to-sensor tasking, such as tip/cue
	// operations.
	IDParentReq param.Opt[string] `json:"idParentReq,omitzero"`
	// Unique identifier of the parent plan or schedule associated with this request.
	// If null, this request is assumed not associated with a plan or schedule.
	IDPlan param.Opt[string] `json:"idPlan,omitzero"`
	// Unique identifier of the requested/scheduled/planned sensor associated with this
	// request. If both idSensor and origSensorId are null then the request is assumed
	// to be a general request for observations or contact on an object, if specified,
	// or an area/volume. In this case, the requester may specify a desired obType.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// ID of the UDL State Vector of the object or central vector associated with this
	// request.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// The angle, in degrees, between the equator and the orbit plane when looking from
	// the center of the Earth. Inclination ranges from 0-180 degrees, with 0-90
	// representing posigrade orbits and 90-180 representing retrograde orbits.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// For optical sensors, the integration time per camera frame, in milliseconds.
	IntegrationTime param.Opt[float64] `json:"integrationTime,omitzero"`
	// Inter-Range Operations Number. Four-digit identifier used to schedule and
	// identify AFSCN contact support for booster, launch, and on-orbit operations.
	Iron param.Opt[int64] `json:"iron,omitzero"`
	// The target object irradiance value.
	Irradiance param.Opt[float64] `json:"irradiance,omitzero"`
	// WGS-84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator). If an accompanying stopLat is provided, then the lat value can be
	// assumed to be the starting latitude of a volume definition.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian). If an accompanying stopLon is provided, then lon value can be assumed
	// to be the starting longitude of a volume definition.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The timestamp of the external message from which this request originated, if
	// applicable, in ISO8601 UTC format with millisecond precision.
	MsgCreateDate param.Opt[time.Time] `json:"msgCreateDate,omitzero" format:"date-time"`
	// The type of external message from which this request originated.
	MsgType param.Opt[string] `json:"msgType,omitzero"`
	// Notes or comments associated with this request.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// For optical sensors, the requested number of frames to capture at each sensor
	// step.
	NumFrames param.Opt[int64] `json:"numFrames,omitzero"`
	// The number of requested observations on the target.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// The number of requested tracks on the target. If numTracks is not provided it is
	// assumed to indicate all possible observations every pass over the request
	// duration or within the request start/end window.
	NumTracks param.Opt[int64] `json:"numTracks,omitzero"`
	// Optional type of observation (EO, IR, RADAR, RF-ACTIVE, RF-PASSIVE, OTHER)
	// requested. This field may correspond to a request of a specific sensor, or to a
	// general non sensor specific request.
	ObType param.Opt[string] `json:"obType,omitzero"`
	// The orbit regime of the target (GEO, HEO, LAUNCH, LEO, MEO, OTHER).
	OrbitRegime param.Opt[string] `json:"orbitRegime,omitzero"`
	// The magnitude of rotation, in degrees, between the xAngle direction and locally
	// defined equinoctial plane. A positive value indicates clockwise rotation about
	// the sensor boresight vector.
	OrientAngle param.Opt[float64] `json:"orientAngle,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the data source to indicate the target object of
	// this request. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the source to indicate the sensor identifier
	// requested/scheduled/planned for this request. This may be an internal identifier
	// and not necessarily a valid sensor ID. If both idSensor and origSensorId are
	// null then the request is assumed to be a general request for observations or
	// contact on an object, if specified, or an area/volume. In this case, the
	// requester may specify a desired obType.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Index number (integer) for records within a collection plan or schedule.
	PlanIndex param.Opt[int64] `json:"planIndex,omitzero"`
	// The RF polarization (H, LHC, RHC, V).
	Polarization param.Opt[string] `json:"polarization,omitzero"`
	// The priority of the collect request (EMERGENCY, FLASH, IMMEDIATE, PRIORITY,
	// ROUTINE).
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The expected or directed right ascension angle, in degrees, for search or target
	// acquisition.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node.
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// The expected acquisition range or defined center range, in km.
	Range param.Opt[float64] `json:"range,omitzero"`
	// The Radar Cross-Section of the target, in m^2.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// The maximum Radar Cross-Section of the target, in m^2.
	RcsMax param.Opt[float64] `json:"rcsMax,omitzero"`
	// The minimum Radar Cross-Section of the target, in m^2. If only minimum RCS is
	// provided it is assumed to be minimum reportable RCS.
	RcsMin param.Opt[float64] `json:"rcsMin,omitzero"`
	// The fraction of solar energy reflected from target.
	Reflectance param.Opt[float64] `json:"reflectance,omitzero"`
	// Satellite/catalog number of the target on-orbit object for this request.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Pre-coordinated code, direction, or configuration to be executed by the sensor
	// or site for this collect or contact.
	Scenario param.Opt[string] `json:"scenario,omitzero"`
	// The average of the periapsis and apoapsis distances, in kilometers. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	// The spectral model used for the irradiance calculation.
	SpectralModel param.Opt[string] `json:"spectralModel,omitzero"`
	// The maximum inclination, in degrees, to be used in search operations.
	SrchInc param.Opt[float64] `json:"srchInc,omitzero"`
	// The search pattern to be executed for this request (e.g. PICKET-FENCE, SCAN,
	// etc.).
	SrchPattern param.Opt[string] `json:"srchPattern,omitzero"`
	// The stopping HAE WGS-84 height above ellipsoid (HAE), of a volume definition, in
	// kilometers. The stopAlt value is only meaningful if a (starting) alt value is
	// provided.
	StopAlt param.Opt[float64] `json:"stopAlt,omitzero"`
	// The stopping WGS-84 latitude of a volume definition, in degrees. -90 to 90
	// degrees (negative values south of equator). The stopLat value is only meaningful
	// if a (starting) lat value is provided.
	StopLat param.Opt[float64] `json:"stopLat,omitzero"`
	// The stopping WGS-84 longitude of a volume definition, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian). The stopLon value is only
	// meaningful if a (starting) lon value is provided.
	StopLon param.Opt[float64] `json:"stopLon,omitzero"`
	// The (SSN) tasking suffix (A-Z) associated with this request. The suffix defines
	// the amount of observational data and the frequency of collection. Note that
	// suffix definitions are sensor type specific.
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// The minimum object (diameter) size, in meters, to be reported.
	TargetSize param.Opt[float64] `json:"targetSize,omitzero"`
	// The (SSN) tasking category (1-5) associated with this request. The tasking
	// category defines the priority of gathering and transmitting the requested
	// observational data. Note that category definitions are sensor type specific.
	TaskCategory param.Opt[int64] `json:"taskCategory,omitzero"`
	// The tasking group to which the target object is assigned.
	TaskGroup param.Opt[string] `json:"taskGroup,omitzero"`
	// Task ID associated with this request. A task ID may be associated with a single
	// collect request or may be used to tie together the sub-requests of a full
	// collect, for example a DWELL consisting of many dwell points.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The true anomaly defines the angular position, in degrees, of the object on it's
	// orbital path as measured from the orbit focal point at epoch. The true anomaly
	// is referenced from perigee.
	TrueAnomoly param.Opt[float64] `json:"trueAnomoly,omitzero"`
	// Boolean indicating that this collect request is UCT follow-up.
	UctFollowUp param.Opt[bool] `json:"uctFollowUp,omitzero"`
	// The estimated or expected visual magnitude of the target, in Magnitudes (M).
	VisMag param.Opt[float64] `json:"visMag,omitzero"`
	// The maximum estimated or expected visual magnitude of the target, in Magnitudes
	// (M).
	VisMagMax param.Opt[float64] `json:"visMagMax,omitzero"`
	// The minimum estimated or expected visual magnitude of the target, in Magnitudes
	// (M). If only minimum vismag is provided it is assumed to be minimum reportable
	// vismag.
	VisMagMin param.Opt[float64] `json:"visMagMin,omitzero"`
	// The angular distance, in degrees, in the sensor-x direction from scan center
	// defined by the central vector. The specification of xAngle and yAngle defines a
	// rectangle of width 2*xAngle and height 2*yAngle centered about the central
	// vector.
	XAngle param.Opt[float64] `json:"xAngle,omitzero"`
	// The angular distance, in degrees, in the sensor-y direction from scan center
	// defined by the central vector. The specification of xAngle and yAngle defines a
	// rectangle of width 2*xAngle and height 2*yAngle centered about the central
	// vector.
	YAngle param.Opt[float64] `json:"yAngle,omitzero"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	Elset CollectRequestNewBulkParamsBodyElset `json:"elset,omitzero"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	StateVector CollectRequestNewBulkParamsBodyStateVector `json:"stateVector,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r CollectRequestNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow CollectRequestNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectRequestNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectRequestNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type CollectRequestNewBulkParamsBodyElset struct {
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

func (r CollectRequestNewBulkParamsBodyElset) MarshalJSON() (data []byte, err error) {
	type shadow CollectRequestNewBulkParamsBodyElset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectRequestNewBulkParamsBodyElset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectRequestNewBulkParamsBodyElset](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
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
type CollectRequestNewBulkParamsBodyStateVector struct {
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
	// Any of "J2000", "UVW", "EFG/TDR", "TEME", "GCRF".
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

func (r CollectRequestNewBulkParamsBodyStateVector) MarshalJSON() (data []byte, err error) {
	type shadow CollectRequestNewBulkParamsBodyStateVector
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectRequestNewBulkParamsBodyStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectRequestNewBulkParamsBodyStateVector](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[CollectRequestNewBulkParamsBodyStateVector](
		"covReferenceFrame", "J2000", "UVW", "EFG/TDR", "TEME", "GCRF",
	)
	apijson.RegisterFieldValidator[CollectRequestNewBulkParamsBodyStateVector](
		"referenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}

type CollectRequestTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The start time or earliest time of the collect or contact request window, in ISO
	// 8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectRequestTupleParams]'s query parameters as
// `url.Values`.
func (r CollectRequestTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectRequestUnvalidatedPublishParams struct {
	Body []CollectRequestUnvalidatedPublishParamsBody
	paramObj
}

func (r CollectRequestUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *CollectRequestUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Collect Requests support several types of individual requests, or
// planned/scheduled tasks on sensors and/or orbital objects. Options are provided
// to accomodate most common sensor contact and collection applications, including
// single sensor-object tasking, search operations, and TT&C support. Multiple
// requests originating from a plan or schedule may be associated to a sensor plan
// if desired.
//
// The properties ClassificationMarking, DataMode, Source, StartTime, Type are
// required.
type CollectRequestUnvalidatedPublishParamsBody struct {
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
	// The start time or earliest time of the collect or contact request window, in ISO
	// 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The type of this collect or contact request (DIRECTED SEARCH, DWELL, OBJECT,
	// POL, RATE TRACK, SEARCH, SOI, STARE, TTC, VOLUME SEARCH, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Height above WGS-84 ellipsoid (HAE), in kilometers. If an accompanying stopAlt
	// is provided, then alt value can be assumed to be the starting altitude of a
	// volume definition.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// The argument of perigee is the angle, in degrees, formed between the perigee and
	// the ascending node.
	ArgOfPerigee param.Opt[float64] `json:"argOfPerigee,omitzero"`
	// The expected or directed azimuth angle, in degrees, for search or target
	// acquisition.
	Az param.Opt[float64] `json:"az,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The customer for this request.
	Customer param.Opt[string] `json:"customer,omitzero"`
	// The expected or directed declination angle, in degrees, for search or target
	// acquisition.
	Dec param.Opt[float64] `json:"dec,omitzero"`
	// The duration of the collect request, in seconds. If both duration and endTime
	// are provided, the endTime is assumed to take precedence.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// The dwell ID associated with this request. A dwell ID is dwell point specific
	// and a DWELL request consist of many dwell point requests.
	DwellID param.Opt[string] `json:"dwellId,omitzero"`
	// The orbital eccentricity of an astronomical object is a parameter that
	// determines the amount by which its orbit around another body deviates from a
	// perfect circle.
	Eccentricity param.Opt[float64] `json:"eccentricity,omitzero"`
	// The expected or directed elevation angle, in degrees, for search or target
	// acquisition.
	El param.Opt[float64] `json:"el,omitzero"`
	// The end time of the collect or contact request window, in ISO 8601 UTC format.
	// If no endTime or duration is provided it is assumed the request is either
	// ongoing or that the request is for a specified number of tracks (numTracks). If
	// both duration and endTime are provided, the endTime is assumed to take
	// precedence.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Epoch time, in ISO 8601 UTC format, of the orbital elements.
	Epoch param.Opt[time.Time] `json:"epoch,omitzero" format:"date-time"`
	// ID of the UDL Ephemeris Set of the object associated with this request.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// The extent of the azimuth angle, in degrees, from center azimuth to define a
	// spatial volume.
	ExtentAz param.Opt[float64] `json:"extentAz,omitzero"`
	// The extent of the elevation angle, in degrees, from center elevation to define a
	// spatial volume.
	ExtentEl param.Opt[float64] `json:"extentEl,omitzero"`
	// The extent of the range, in km, from center range to define a spatial volume.
	ExtentRange param.Opt[float64] `json:"extentRange,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// For optical sensors, the frame rate of the camera, in Hz.
	FrameRate param.Opt[float64] `json:"frameRate,omitzero"`
	// The estimated or expected emission frequency of the target, in MHz.
	Freq param.Opt[float64] `json:"freq,omitzero"`
	// The maximum frequency of interest, in MHz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// The minimum frequency of interest, in MHz. If only minimum frequency is provided
	// it is assumed to be minimum reportable frequency.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// ID of the UDL Elset of the object associated with this request.
	IDElset param.Opt[string] `json:"idElset,omitzero"`
	// ID of the UDL Manifold Elset of the object associated with this request. A
	// Manifold Elset provides theoretical Keplerian orbital elements belonging to an
	// object of interest's manifold describing a possible/theoretical orbit for an
	// object of interest for tasking purposes.
	IDManifold param.Opt[string] `json:"idManifold,omitzero"`
	// Unique identifier of the target on-orbit object for this request.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// The unique ID of the collect request record from which this request originated.
	// This may be used for cases of sensor-to-sensor tasking, such as tip/cue
	// operations.
	IDParentReq param.Opt[string] `json:"idParentReq,omitzero"`
	// Unique identifier of the parent plan or schedule associated with this request.
	// If null, this request is assumed not associated with a plan or schedule.
	IDPlan param.Opt[string] `json:"idPlan,omitzero"`
	// Unique identifier of the requested/scheduled/planned sensor associated with this
	// request. If both idSensor and origSensorId are null then the request is assumed
	// to be a general request for observations or contact on an object, if specified,
	// or an area/volume. In this case, the requester may specify a desired obType.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// ID of the UDL State Vector of the object or central vector associated with this
	// request.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// The angle, in degrees, between the equator and the orbit plane when looking from
	// the center of the Earth. Inclination ranges from 0-180 degrees, with 0-90
	// representing posigrade orbits and 90-180 representing retrograde orbits.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// For optical sensors, the integration time per camera frame, in milliseconds.
	IntegrationTime param.Opt[float64] `json:"integrationTime,omitzero"`
	// Inter-Range Operations Number. Four-digit identifier used to schedule and
	// identify AFSCN contact support for booster, launch, and on-orbit operations.
	Iron param.Opt[int64] `json:"iron,omitzero"`
	// The target object irradiance value.
	Irradiance param.Opt[float64] `json:"irradiance,omitzero"`
	// WGS-84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator). If an accompanying stopLat is provided, then the lat value can be
	// assumed to be the starting latitude of a volume definition.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian). If an accompanying stopLon is provided, then lon value can be assumed
	// to be the starting longitude of a volume definition.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The timestamp of the external message from which this request originated, if
	// applicable, in ISO8601 UTC format with millisecond precision.
	MsgCreateDate param.Opt[time.Time] `json:"msgCreateDate,omitzero" format:"date-time"`
	// The type of external message from which this request originated.
	MsgType param.Opt[string] `json:"msgType,omitzero"`
	// Notes or comments associated with this request.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// For optical sensors, the requested number of frames to capture at each sensor
	// step.
	NumFrames param.Opt[int64] `json:"numFrames,omitzero"`
	// The number of requested observations on the target.
	NumObs param.Opt[int64] `json:"numObs,omitzero"`
	// The number of requested tracks on the target. If numTracks is not provided it is
	// assumed to indicate all possible observations every pass over the request
	// duration or within the request start/end window.
	NumTracks param.Opt[int64] `json:"numTracks,omitzero"`
	// Optional type of observation (EO, IR, RADAR, RF-ACTIVE, RF-PASSIVE, OTHER)
	// requested. This field may correspond to a request of a specific sensor, or to a
	// general non sensor specific request.
	ObType param.Opt[string] `json:"obType,omitzero"`
	// The orbit regime of the target (GEO, HEO, LAUNCH, LEO, MEO, OTHER).
	OrbitRegime param.Opt[string] `json:"orbitRegime,omitzero"`
	// The magnitude of rotation, in degrees, between the xAngle direction and locally
	// defined equinoctial plane. A positive value indicates clockwise rotation about
	// the sensor boresight vector.
	OrientAngle param.Opt[float64] `json:"orientAngle,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the data source to indicate the target object of
	// this request. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the source to indicate the sensor identifier
	// requested/scheduled/planned for this request. This may be an internal identifier
	// and not necessarily a valid sensor ID. If both idSensor and origSensorId are
	// null then the request is assumed to be a general request for observations or
	// contact on an object, if specified, or an area/volume. In this case, the
	// requester may specify a desired obType.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Index number (integer) for records within a collection plan or schedule.
	PlanIndex param.Opt[int64] `json:"planIndex,omitzero"`
	// The RF polarization (H, LHC, RHC, V).
	Polarization param.Opt[string] `json:"polarization,omitzero"`
	// The priority of the collect request (EMERGENCY, FLASH, IMMEDIATE, PRIORITY,
	// ROUTINE).
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The expected or directed right ascension angle, in degrees, for search or target
	// acquisition.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Right ascension of the ascending node, or RAAN is the angle as measured in
	// degrees eastwards (or, as seen from the north, counterclockwise) from the First
	// Point of Aries to the ascending node.
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// The expected acquisition range or defined center range, in km.
	Range param.Opt[float64] `json:"range,omitzero"`
	// The Radar Cross-Section of the target, in m^2.
	Rcs param.Opt[float64] `json:"rcs,omitzero"`
	// The maximum Radar Cross-Section of the target, in m^2.
	RcsMax param.Opt[float64] `json:"rcsMax,omitzero"`
	// The minimum Radar Cross-Section of the target, in m^2. If only minimum RCS is
	// provided it is assumed to be minimum reportable RCS.
	RcsMin param.Opt[float64] `json:"rcsMin,omitzero"`
	// The fraction of solar energy reflected from target.
	Reflectance param.Opt[float64] `json:"reflectance,omitzero"`
	// Satellite/catalog number of the target on-orbit object for this request.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Pre-coordinated code, direction, or configuration to be executed by the sensor
	// or site for this collect or contact.
	Scenario param.Opt[string] `json:"scenario,omitzero"`
	// The average of the periapsis and apoapsis distances, in kilometers. For circular
	// orbits, the semimajor axis is the distance between the centers of the bodies.
	SemiMajorAxis param.Opt[float64] `json:"semiMajorAxis,omitzero"`
	// The spectral model used for the irradiance calculation.
	SpectralModel param.Opt[string] `json:"spectralModel,omitzero"`
	// The maximum inclination, in degrees, to be used in search operations.
	SrchInc param.Opt[float64] `json:"srchInc,omitzero"`
	// The search pattern to be executed for this request (e.g. PICKET-FENCE, SCAN,
	// etc.).
	SrchPattern param.Opt[string] `json:"srchPattern,omitzero"`
	// The stopping HAE WGS-84 height above ellipsoid (HAE), of a volume definition, in
	// kilometers. The stopAlt value is only meaningful if a (starting) alt value is
	// provided.
	StopAlt param.Opt[float64] `json:"stopAlt,omitzero"`
	// The stopping WGS-84 latitude of a volume definition, in degrees. -90 to 90
	// degrees (negative values south of equator). The stopLat value is only meaningful
	// if a (starting) lat value is provided.
	StopLat param.Opt[float64] `json:"stopLat,omitzero"`
	// The stopping WGS-84 longitude of a volume definition, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian). The stopLon value is only
	// meaningful if a (starting) lon value is provided.
	StopLon param.Opt[float64] `json:"stopLon,omitzero"`
	// The (SSN) tasking suffix (A-Z) associated with this request. The suffix defines
	// the amount of observational data and the frequency of collection. Note that
	// suffix definitions are sensor type specific.
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// The minimum object (diameter) size, in meters, to be reported.
	TargetSize param.Opt[float64] `json:"targetSize,omitzero"`
	// The (SSN) tasking category (1-5) associated with this request. The tasking
	// category defines the priority of gathering and transmitting the requested
	// observational data. Note that category definitions are sensor type specific.
	TaskCategory param.Opt[int64] `json:"taskCategory,omitzero"`
	// The tasking group to which the target object is assigned.
	TaskGroup param.Opt[string] `json:"taskGroup,omitzero"`
	// Task ID associated with this request. A task ID may be associated with a single
	// collect request or may be used to tie together the sub-requests of a full
	// collect, for example a DWELL consisting of many dwell points.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// The true anomaly defines the angular position, in degrees, of the object on it's
	// orbital path as measured from the orbit focal point at epoch. The true anomaly
	// is referenced from perigee.
	TrueAnomoly param.Opt[float64] `json:"trueAnomoly,omitzero"`
	// Boolean indicating that this collect request is UCT follow-up.
	UctFollowUp param.Opt[bool] `json:"uctFollowUp,omitzero"`
	// The estimated or expected visual magnitude of the target, in Magnitudes (M).
	VisMag param.Opt[float64] `json:"visMag,omitzero"`
	// The maximum estimated or expected visual magnitude of the target, in Magnitudes
	// (M).
	VisMagMax param.Opt[float64] `json:"visMagMax,omitzero"`
	// The minimum estimated or expected visual magnitude of the target, in Magnitudes
	// (M). If only minimum vismag is provided it is assumed to be minimum reportable
	// vismag.
	VisMagMin param.Opt[float64] `json:"visMagMin,omitzero"`
	// The angular distance, in degrees, in the sensor-x direction from scan center
	// defined by the central vector. The specification of xAngle and yAngle defines a
	// rectangle of width 2*xAngle and height 2*yAngle centered about the central
	// vector.
	XAngle param.Opt[float64] `json:"xAngle,omitzero"`
	// The angular distance, in degrees, in the sensor-y direction from scan center
	// defined by the central vector. The specification of xAngle and yAngle defines a
	// rectangle of width 2*xAngle and height 2*yAngle centered about the central
	// vector.
	YAngle param.Opt[float64] `json:"yAngle,omitzero"`
	// An element set is a collection of Keplerian orbital elements describing an orbit
	// of a particular satellite. The data is used along with an orbit propagator in
	// order to predict the motion of a satellite. The element set, or elset for short,
	// consists of identification data, the classical elements and drag parameters.
	Elset CollectRequestUnvalidatedPublishParamsBodyElset `json:"elset,omitzero"`
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	StateVector CollectRequestUnvalidatedPublishParamsBodyStateVector `json:"stateVector,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r CollectRequestUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow CollectRequestUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectRequestUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectRequestUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// An element set is a collection of Keplerian orbital elements describing an orbit
// of a particular satellite. The data is used along with an orbit propagator in
// order to predict the motion of a satellite. The element set, or elset for short,
// consists of identification data, the classical elements and drag parameters.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type CollectRequestUnvalidatedPublishParamsBodyElset struct {
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

func (r CollectRequestUnvalidatedPublishParamsBodyElset) MarshalJSON() (data []byte, err error) {
	type shadow CollectRequestUnvalidatedPublishParamsBodyElset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectRequestUnvalidatedPublishParamsBodyElset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectRequestUnvalidatedPublishParamsBodyElset](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
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
type CollectRequestUnvalidatedPublishParamsBodyStateVector struct {
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
	// Any of "J2000", "UVW", "EFG/TDR", "TEME", "GCRF".
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

func (r CollectRequestUnvalidatedPublishParamsBodyStateVector) MarshalJSON() (data []byte, err error) {
	type shadow CollectRequestUnvalidatedPublishParamsBodyStateVector
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectRequestUnvalidatedPublishParamsBodyStateVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectRequestUnvalidatedPublishParamsBodyStateVector](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[CollectRequestUnvalidatedPublishParamsBodyStateVector](
		"covReferenceFrame", "J2000", "UVW", "EFG/TDR", "TEME", "GCRF",
	)
	apijson.RegisterFieldValidator[CollectRequestUnvalidatedPublishParamsBodyStateVector](
		"referenceFrame", "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF",
	)
}
