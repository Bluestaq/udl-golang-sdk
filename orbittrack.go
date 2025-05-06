// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// OrbittrackService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrbittrackService] method instead.
type OrbittrackService struct {
	Options []option.RequestOption
	History OrbittrackHistoryService
}

// NewOrbittrackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrbittrackService(opts ...option.RequestOption) (r OrbittrackService) {
	r = OrbittrackService{}
	r.Options = opts
	r.History = NewOrbittrackHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OrbittrackService) List(ctx context.Context, query OrbittrackListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OrbittrackListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/orbittrack"
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
func (r *OrbittrackService) ListAutoPaging(ctx context.Context, query OrbittrackListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OrbittrackListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OrbittrackService) Count(ctx context.Context, query OrbittrackCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/orbittrack/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of orbit
// track records as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *OrbittrackService) NewBulk(ctx context.Context, body OrbittrackNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/orbittrack/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *OrbittrackService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/orbittrack/queryhelp"
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
func (r *OrbittrackService) Tuple(ctx context.Context, query OrbittrackTupleParams, opts ...option.RequestOption) (res *[]OrbittrackTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/orbittrack/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple orbit track records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *OrbittrackService) UnvalidatedPublish(ctx context.Context, body OrbittrackUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-orbittrack"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Keplerian orbital elements describing an orbit for a particular on-orbit
// satellite and applicable sensor data aiding in the orbit prediction.
type OrbittrackListResponse struct {
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
	DataMode OrbittrackListResponseDataMode `json:"dataMode,required"`
	// WGS-84 latitude of the track object subpoint, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	Lat float64 `json:"lat,required"`
	// WGS-84 longitude of the track object subpoint, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Track point altitude relative to WGS-84 ellipsoid, in meters.
	Alt float64 `json:"alt"`
	// Free-form remarks entered for the satellite.
	Amplification string `json:"amplification"`
	// The angle formed between the line of sight of the observer and the horizon at
	// track timestamp, in degrees. The angular range is -90 to 90, with negative
	// values representing angle of depression.
	AngElev float64 `json:"angElev"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// ELLIPSE:
	//
	// brg - orientation in degrees of the ellipse
	//
	// a1 - semi-major axis in meters
	//
	// a2 - semi-minor axis in meters
	//
	// BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// brg - orientation in degrees of the bearing box
	//
	// a1 - length of bearing box in meters
	//
	// a2 - half-width of bearing box in meters
	//
	// OTHER (All other type values):
	//
	// brg - line of bearing in degrees true
	//
	// a1 - bearing error in degrees
	//
	// a2 - estimated range in meters.
	AouData []float64 `json:"aouData"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouData array and is required if aouData is not
	// null. See the aouData field definition for specific information.
	AouType string `json:"aouType"`
	// International radio call sign assigned to the track. This is an 8-character
	// alphanumeric code.
	CallSign string `json:"callSign"`
	// One-line Charlie elements set.
	CharlieLine string `json:"charlieLine"`
	// The cross-reference code of the channel on which this track report was received,
	// if the report came over a comms channel.
	ChXRef string `json:"chXRef"`
	// The Area Of Uncertainty (AOU) percentage (0 - 100) containment value. The
	// percentage of time (90%) that the estimated area of uncertainty will cover the
	// true position of the track object.
	Cntnmnt float64 `json:"cntnmnt"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Predicted change in Mean Motion (velocity) in radians/herg^2. herg is a unit of
	// time measure equal to 806.8120769 seconds, and is the orbital period of an
	// imaginary satellite rotating about the Earth at zero altitude.
	Decay float64 `json:"decay"`
	// Flag indicating that this track represents a dummy object or group. Identifies
	// offensive or defensive units, equipment and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Dummy bool `json:"dummy"`
	// Flag indicating that this track represents a feint object or group. Identifies
	// offensive or defensive units, equipment and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Feint bool `json:"feint"`
	// Flag indicating that this track represents a headquarters object. Based on
	// MIL-STD-2525 symbology definitions.
	Hq bool `json:"hq"`
	// Unique identifier of the Elset associated with this object.
	IDElset string `json:"idElset"`
	// Additional track object identity/status information, typically used for EXERCISE
	// identity amplification (FAKER, JOKER, KILO, TRAVELLER, ZOMBIE):
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// KILO: Friendly high-value object.
	//
	// TRAVELLER: Suspect land or surface track following a recognized traffic route.
	//
	// ZOMBIE: Suspect track, object, or entity of special interest.
	IdentAmp string `json:"identAmp"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// A text amplifier displaying IFF/SIF/AIS Identification modes and codes.
	Iff string `json:"iff"`
	// Flag indicating that this track represents an installation. Based on
	// MIL-STD-2525 symbology definitions.
	Installation bool `json:"installation"`
	// The on-orbit category assigned to this track object (DEBRIS, MANNED, PAYLOAD,
	// PLATFORM, ROCKET BODY, UNKNOWN).
	//
	// Any of "DEBRIS", "MANNED", "PAYLOAD", "PLATFORM", "ROCKET BODY", "UNKNOWN".
	ObjectType OrbittrackListResponseObjectType `json:"objectType"`
	// The estimated identity of the track object (ASSUMED FRIEND, FRIEND, HOSTILE,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	//
	// Any of "ASSUMED FRIEND", "FRIEND", "HOSTILE", "NEUTRAL", "PENDING", "SUSPECT",
	// "UNKNOWN".
	ObjIdent OrbittrackListResponseObjIdent `json:"objIdent"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by observation source to indicate the target
	// on-orbit object of this track. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Radio frequency of the track, measured in megahertz (MHz).
	RdfRf float64 `json:"rdfRF"`
	// Flag indicating that this track represents a reduced object or group. Based on
	// MIL-STD-2525 symbology definitions.
	Reduced bool `json:"reduced"`
	// Flag indicating that this track represents a reinforced object or group. Based
	// on MIL-STD-2525 symbology definitions.
	Reinforced bool `json:"reinforced"`
	// Report number received from the reporting source for this track.
	RptNum string `json:"rptNum"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Status of the satellite.
	SatStatus string `json:"satStatus"`
	// Track object speed, in km/sec.
	Spd float64 `json:"spd"`
	// Flag indicating that this track represents a task force. Based on MIL-STD-2525
	// symbology definitions.
	TaskForce bool `json:"taskForce"`
	// TrackSensor Collection.
	TrackSensors []OrbittrackListResponseTrackSensor `json:"trackSensors"`
	// UUID identifying the track, which should remain the same on subsequent tracks of
	// the same object.
	TrkID string `json:"trkId"`
	// The type of vehicle with which the device is associated. Based on MIL-STD-2525
	// symbology definitions.
	VehType string `json:"vehType"`
	// Source cross-reference code for the command that originated the track report.
	Xref string `json:"xref"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Lat                   resp.Field
		Lon                   resp.Field
		Source                resp.Field
		Ts                    resp.Field
		ID                    resp.Field
		Alt                   resp.Field
		Amplification         resp.Field
		AngElev               resp.Field
		AouData               resp.Field
		AouType               resp.Field
		CallSign              resp.Field
		CharlieLine           resp.Field
		ChXRef                resp.Field
		Cntnmnt               resp.Field
		CountryCode           resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Decay                 resp.Field
		Dummy                 resp.Field
		Feint                 resp.Field
		Hq                    resp.Field
		IDElset               resp.Field
		IdentAmp              resp.Field
		IDOnOrbit             resp.Field
		Iff                   resp.Field
		Installation          resp.Field
		ObjectType            resp.Field
		ObjIdent              resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		RdfRf                 resp.Field
		Reduced               resp.Field
		Reinforced            resp.Field
		RptNum                resp.Field
		SatNo                 resp.Field
		SatStatus             resp.Field
		Spd                   resp.Field
		TaskForce             resp.Field
		TrackSensors          resp.Field
		TrkID                 resp.Field
		VehType               resp.Field
		Xref                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrbittrackListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrbittrackListResponse) UnmarshalJSON(data []byte) error {
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
type OrbittrackListResponseDataMode string

const (
	OrbittrackListResponseDataModeReal      OrbittrackListResponseDataMode = "REAL"
	OrbittrackListResponseDataModeTest      OrbittrackListResponseDataMode = "TEST"
	OrbittrackListResponseDataModeSimulated OrbittrackListResponseDataMode = "SIMULATED"
	OrbittrackListResponseDataModeExercise  OrbittrackListResponseDataMode = "EXERCISE"
)

// The on-orbit category assigned to this track object (DEBRIS, MANNED, PAYLOAD,
// PLATFORM, ROCKET BODY, UNKNOWN).
type OrbittrackListResponseObjectType string

const (
	OrbittrackListResponseObjectTypeDebris     OrbittrackListResponseObjectType = "DEBRIS"
	OrbittrackListResponseObjectTypeManned     OrbittrackListResponseObjectType = "MANNED"
	OrbittrackListResponseObjectTypePayload    OrbittrackListResponseObjectType = "PAYLOAD"
	OrbittrackListResponseObjectTypePlatform   OrbittrackListResponseObjectType = "PLATFORM"
	OrbittrackListResponseObjectTypeRocketBody OrbittrackListResponseObjectType = "ROCKET BODY"
	OrbittrackListResponseObjectTypeUnknown    OrbittrackListResponseObjectType = "UNKNOWN"
)

// The estimated identity of the track object (ASSUMED FRIEND, FRIEND, HOSTILE,
// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
//
// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
// behavior, and/or origin.
//
// FRIEND: Track object supporting friendly forces and belonging to a declared
// friendly nation or entity.
//
// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
// deemed to contribute to a threat to friendly forces or their mission due to its
// behavior, characteristics, nationality, or origin.
//
// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
// origin indicate that it is neither supporting nor opposing friendly forces or
// their mission.
//
// PENDING: Track object which has not been evaluated.
//
// SUSPECT: Track object deemed potentially hostile due to the object
// characteristics, behavior, nationality, and/or origin.
//
// UNKNOWN: Track object which has been evaluated and does not meet criteria for
// any standard identity.
type OrbittrackListResponseObjIdent string

const (
	OrbittrackListResponseObjIdentAssumedFriend OrbittrackListResponseObjIdent = "ASSUMED FRIEND"
	OrbittrackListResponseObjIdentFriend        OrbittrackListResponseObjIdent = "FRIEND"
	OrbittrackListResponseObjIdentHostile       OrbittrackListResponseObjIdent = "HOSTILE"
	OrbittrackListResponseObjIdentNeutral       OrbittrackListResponseObjIdent = "NEUTRAL"
	OrbittrackListResponseObjIdentPending       OrbittrackListResponseObjIdent = "PENDING"
	OrbittrackListResponseObjIdentSuspect       OrbittrackListResponseObjIdent = "SUSPECT"
	OrbittrackListResponseObjIdentUnknown       OrbittrackListResponseObjIdent = "UNKNOWN"
)

// Schema for Track Sensor data.
type OrbittrackListResponseTrackSensor struct {
	// The observing sensor azimuth angle, in degrees and topocentric frame.
	Az float64 `json:"az,required"`
	// The track object range from the observing sensor, in kilometers.
	Range float64 `json:"range,required"`
	// Minimum range measurement capability of the sensor, in kilometers.
	MinRangeLimit float64 `json:"minRangeLimit"`
	// The mission number which produced this track observation.
	MissionNumber string `json:"missionNumber"`
	// The field of view (FOV) type (Butterfly, Cone Angular, Cone Distance, Horizon to
	// Horizon, Unknown) employed by the sensor observing this object.
	//
	// Any of "BUTTERFLY", "CONE ANGULAR", "CONE DISTANCE", "HORIZON TO HORIZON",
	// "UNKNOWN".
	SensorFovType string `json:"sensorFOVType"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber int64 `json:"sensorNumber"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Az            resp.Field
		Range         resp.Field
		MinRangeLimit resp.Field
		MissionNumber resp.Field
		SensorFovType resp.Field
		SensorName    resp.Field
		SensorNumber  resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrbittrackListResponseTrackSensor) RawJSON() string { return r.JSON.raw }
func (r *OrbittrackListResponseTrackSensor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Keplerian orbital elements describing an orbit for a particular on-orbit
// satellite and applicable sensor data aiding in the orbit prediction.
type OrbittrackTupleResponse struct {
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
	DataMode OrbittrackTupleResponseDataMode `json:"dataMode,required"`
	// WGS-84 latitude of the track object subpoint, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	Lat float64 `json:"lat,required"`
	// WGS-84 longitude of the track object subpoint, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Track point altitude relative to WGS-84 ellipsoid, in meters.
	Alt float64 `json:"alt"`
	// Free-form remarks entered for the satellite.
	Amplification string `json:"amplification"`
	// The angle formed between the line of sight of the observer and the horizon at
	// track timestamp, in degrees. The angular range is -90 to 90, with negative
	// values representing angle of depression.
	AngElev float64 `json:"angElev"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// ELLIPSE:
	//
	// brg - orientation in degrees of the ellipse
	//
	// a1 - semi-major axis in meters
	//
	// a2 - semi-minor axis in meters
	//
	// BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// brg - orientation in degrees of the bearing box
	//
	// a1 - length of bearing box in meters
	//
	// a2 - half-width of bearing box in meters
	//
	// OTHER (All other type values):
	//
	// brg - line of bearing in degrees true
	//
	// a1 - bearing error in degrees
	//
	// a2 - estimated range in meters.
	AouData []float64 `json:"aouData"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouData array and is required if aouData is not
	// null. See the aouData field definition for specific information.
	AouType string `json:"aouType"`
	// International radio call sign assigned to the track. This is an 8-character
	// alphanumeric code.
	CallSign string `json:"callSign"`
	// One-line Charlie elements set.
	CharlieLine string `json:"charlieLine"`
	// The cross-reference code of the channel on which this track report was received,
	// if the report came over a comms channel.
	ChXRef string `json:"chXRef"`
	// The Area Of Uncertainty (AOU) percentage (0 - 100) containment value. The
	// percentage of time (90%) that the estimated area of uncertainty will cover the
	// true position of the track object.
	Cntnmnt float64 `json:"cntnmnt"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Predicted change in Mean Motion (velocity) in radians/herg^2. herg is a unit of
	// time measure equal to 806.8120769 seconds, and is the orbital period of an
	// imaginary satellite rotating about the Earth at zero altitude.
	Decay float64 `json:"decay"`
	// Flag indicating that this track represents a dummy object or group. Identifies
	// offensive or defensive units, equipment and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Dummy bool `json:"dummy"`
	// Flag indicating that this track represents a feint object or group. Identifies
	// offensive or defensive units, equipment and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Feint bool `json:"feint"`
	// Flag indicating that this track represents a headquarters object. Based on
	// MIL-STD-2525 symbology definitions.
	Hq bool `json:"hq"`
	// Unique identifier of the Elset associated with this object.
	IDElset string `json:"idElset"`
	// Additional track object identity/status information, typically used for EXERCISE
	// identity amplification (FAKER, JOKER, KILO, TRAVELLER, ZOMBIE):
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// KILO: Friendly high-value object.
	//
	// TRAVELLER: Suspect land or surface track following a recognized traffic route.
	//
	// ZOMBIE: Suspect track, object, or entity of special interest.
	IdentAmp string `json:"identAmp"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// A text amplifier displaying IFF/SIF/AIS Identification modes and codes.
	Iff string `json:"iff"`
	// Flag indicating that this track represents an installation. Based on
	// MIL-STD-2525 symbology definitions.
	Installation bool `json:"installation"`
	// The on-orbit category assigned to this track object (DEBRIS, MANNED, PAYLOAD,
	// PLATFORM, ROCKET BODY, UNKNOWN).
	//
	// Any of "DEBRIS", "MANNED", "PAYLOAD", "PLATFORM", "ROCKET BODY", "UNKNOWN".
	ObjectType OrbittrackTupleResponseObjectType `json:"objectType"`
	// The estimated identity of the track object (ASSUMED FRIEND, FRIEND, HOSTILE,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	//
	// Any of "ASSUMED FRIEND", "FRIEND", "HOSTILE", "NEUTRAL", "PENDING", "SUSPECT",
	// "UNKNOWN".
	ObjIdent OrbittrackTupleResponseObjIdent `json:"objIdent"`
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
	// Optional identifier provided by observation source to indicate the target
	// on-orbit object of this track. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Radio frequency of the track, measured in megahertz (MHz).
	RdfRf float64 `json:"rdfRF"`
	// Flag indicating that this track represents a reduced object or group. Based on
	// MIL-STD-2525 symbology definitions.
	Reduced bool `json:"reduced"`
	// Flag indicating that this track represents a reinforced object or group. Based
	// on MIL-STD-2525 symbology definitions.
	Reinforced bool `json:"reinforced"`
	// Report number received from the reporting source for this track.
	RptNum string `json:"rptNum"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Status of the satellite.
	SatStatus string `json:"satStatus"`
	// Track object speed, in km/sec.
	Spd float64 `json:"spd"`
	// Flag indicating that this track represents a task force. Based on MIL-STD-2525
	// symbology definitions.
	TaskForce bool `json:"taskForce"`
	// TrackSensor Collection.
	TrackSensors []OrbittrackTupleResponseTrackSensor `json:"trackSensors"`
	// UUID identifying the track, which should remain the same on subsequent tracks of
	// the same object.
	TrkID string `json:"trkId"`
	// The type of vehicle with which the device is associated. Based on MIL-STD-2525
	// symbology definitions.
	VehType string `json:"vehType"`
	// Source cross-reference code for the command that originated the track report.
	Xref string `json:"xref"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Lat                   resp.Field
		Lon                   resp.Field
		Source                resp.Field
		Ts                    resp.Field
		ID                    resp.Field
		Alt                   resp.Field
		Amplification         resp.Field
		AngElev               resp.Field
		AouData               resp.Field
		AouType               resp.Field
		CallSign              resp.Field
		CharlieLine           resp.Field
		ChXRef                resp.Field
		Cntnmnt               resp.Field
		CountryCode           resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Decay                 resp.Field
		Dummy                 resp.Field
		Feint                 resp.Field
		Hq                    resp.Field
		IDElset               resp.Field
		IdentAmp              resp.Field
		IDOnOrbit             resp.Field
		Iff                   resp.Field
		Installation          resp.Field
		ObjectType            resp.Field
		ObjIdent              resp.Field
		OnOrbit               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		RdfRf                 resp.Field
		Reduced               resp.Field
		Reinforced            resp.Field
		RptNum                resp.Field
		SatNo                 resp.Field
		SatStatus             resp.Field
		Spd                   resp.Field
		TaskForce             resp.Field
		TrackSensors          resp.Field
		TrkID                 resp.Field
		VehType               resp.Field
		Xref                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrbittrackTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *OrbittrackTupleResponse) UnmarshalJSON(data []byte) error {
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
type OrbittrackTupleResponseDataMode string

const (
	OrbittrackTupleResponseDataModeReal      OrbittrackTupleResponseDataMode = "REAL"
	OrbittrackTupleResponseDataModeTest      OrbittrackTupleResponseDataMode = "TEST"
	OrbittrackTupleResponseDataModeSimulated OrbittrackTupleResponseDataMode = "SIMULATED"
	OrbittrackTupleResponseDataModeExercise  OrbittrackTupleResponseDataMode = "EXERCISE"
)

// The on-orbit category assigned to this track object (DEBRIS, MANNED, PAYLOAD,
// PLATFORM, ROCKET BODY, UNKNOWN).
type OrbittrackTupleResponseObjectType string

const (
	OrbittrackTupleResponseObjectTypeDebris     OrbittrackTupleResponseObjectType = "DEBRIS"
	OrbittrackTupleResponseObjectTypeManned     OrbittrackTupleResponseObjectType = "MANNED"
	OrbittrackTupleResponseObjectTypePayload    OrbittrackTupleResponseObjectType = "PAYLOAD"
	OrbittrackTupleResponseObjectTypePlatform   OrbittrackTupleResponseObjectType = "PLATFORM"
	OrbittrackTupleResponseObjectTypeRocketBody OrbittrackTupleResponseObjectType = "ROCKET BODY"
	OrbittrackTupleResponseObjectTypeUnknown    OrbittrackTupleResponseObjectType = "UNKNOWN"
)

// The estimated identity of the track object (ASSUMED FRIEND, FRIEND, HOSTILE,
// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
//
// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
// behavior, and/or origin.
//
// FRIEND: Track object supporting friendly forces and belonging to a declared
// friendly nation or entity.
//
// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
// deemed to contribute to a threat to friendly forces or their mission due to its
// behavior, characteristics, nationality, or origin.
//
// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
// origin indicate that it is neither supporting nor opposing friendly forces or
// their mission.
//
// PENDING: Track object which has not been evaluated.
//
// SUSPECT: Track object deemed potentially hostile due to the object
// characteristics, behavior, nationality, and/or origin.
//
// UNKNOWN: Track object which has been evaluated and does not meet criteria for
// any standard identity.
type OrbittrackTupleResponseObjIdent string

const (
	OrbittrackTupleResponseObjIdentAssumedFriend OrbittrackTupleResponseObjIdent = "ASSUMED FRIEND"
	OrbittrackTupleResponseObjIdentFriend        OrbittrackTupleResponseObjIdent = "FRIEND"
	OrbittrackTupleResponseObjIdentHostile       OrbittrackTupleResponseObjIdent = "HOSTILE"
	OrbittrackTupleResponseObjIdentNeutral       OrbittrackTupleResponseObjIdent = "NEUTRAL"
	OrbittrackTupleResponseObjIdentPending       OrbittrackTupleResponseObjIdent = "PENDING"
	OrbittrackTupleResponseObjIdentSuspect       OrbittrackTupleResponseObjIdent = "SUSPECT"
	OrbittrackTupleResponseObjIdentUnknown       OrbittrackTupleResponseObjIdent = "UNKNOWN"
)

// Schema for Track Sensor data.
type OrbittrackTupleResponseTrackSensor struct {
	// The observing sensor azimuth angle, in degrees and topocentric frame.
	Az float64 `json:"az,required"`
	// The track object range from the observing sensor, in kilometers.
	Range float64 `json:"range,required"`
	// Minimum range measurement capability of the sensor, in kilometers.
	MinRangeLimit float64 `json:"minRangeLimit"`
	// The mission number which produced this track observation.
	MissionNumber string `json:"missionNumber"`
	// The field of view (FOV) type (Butterfly, Cone Angular, Cone Distance, Horizon to
	// Horizon, Unknown) employed by the sensor observing this object.
	//
	// Any of "BUTTERFLY", "CONE ANGULAR", "CONE DISTANCE", "HORIZON TO HORIZON",
	// "UNKNOWN".
	SensorFovType string `json:"sensorFOVType"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber int64 `json:"sensorNumber"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Az            resp.Field
		Range         resp.Field
		MinRangeLimit resp.Field
		MissionNumber resp.Field
		SensorFovType resp.Field
		SensorName    resp.Field
		SensorNumber  resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrbittrackTupleResponseTrackSensor) RawJSON() string { return r.JSON.raw }
func (r *OrbittrackTupleResponseTrackSensor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrbittrackListParams struct {
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrbittrackListParams]'s query parameters as `url.Values`.
func (r OrbittrackListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrbittrackCountParams struct {
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrbittrackCountParams]'s query parameters as `url.Values`.
func (r OrbittrackCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrbittrackNewBulkParams struct {
	Body []OrbittrackNewBulkParamsBody
	paramObj
}

func (r OrbittrackNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Keplerian orbital elements describing an orbit for a particular on-orbit
// satellite and applicable sensor data aiding in the orbit prediction.
//
// The properties ClassificationMarking, DataMode, Lat, Lon, Source, Ts are
// required.
type OrbittrackNewBulkParamsBody struct {
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
	// WGS-84 latitude of the track object subpoint, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	Lat float64 `json:"lat,required"`
	// WGS-84 longitude of the track object subpoint, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Track point altitude relative to WGS-84 ellipsoid, in meters.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Free-form remarks entered for the satellite.
	Amplification param.Opt[string] `json:"amplification,omitzero"`
	// The angle formed between the line of sight of the observer and the horizon at
	// track timestamp, in degrees. The angular range is -90 to 90, with negative
	// values representing angle of depression.
	AngElev param.Opt[float64] `json:"angElev,omitzero"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouData array and is required if aouData is not
	// null. See the aouData field definition for specific information.
	AouType param.Opt[string] `json:"aouType,omitzero"`
	// International radio call sign assigned to the track. This is an 8-character
	// alphanumeric code.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// One-line Charlie elements set.
	CharlieLine param.Opt[string] `json:"charlieLine,omitzero"`
	// The cross-reference code of the channel on which this track report was received,
	// if the report came over a comms channel.
	ChXRef param.Opt[string] `json:"chXRef,omitzero"`
	// The Area Of Uncertainty (AOU) percentage (0 - 100) containment value. The
	// percentage of time (90%) that the estimated area of uncertainty will cover the
	// true position of the track object.
	Cntnmnt param.Opt[float64] `json:"cntnmnt,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Time the row was created in the database.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Predicted change in Mean Motion (velocity) in radians/herg^2. herg is a unit of
	// time measure equal to 806.8120769 seconds, and is the orbital period of an
	// imaginary satellite rotating about the Earth at zero altitude.
	Decay param.Opt[float64] `json:"decay,omitzero"`
	// Flag indicating that this track represents a dummy object or group. Identifies
	// offensive or defensive units, equipment and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Dummy param.Opt[bool] `json:"dummy,omitzero"`
	// Flag indicating that this track represents a feint object or group. Identifies
	// offensive or defensive units, equipment and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Feint param.Opt[bool] `json:"feint,omitzero"`
	// Flag indicating that this track represents a headquarters object. Based on
	// MIL-STD-2525 symbology definitions.
	Hq param.Opt[bool] `json:"hq,omitzero"`
	// Unique identifier of the Elset associated with this object.
	IDElset param.Opt[string] `json:"idElset,omitzero"`
	// Additional track object identity/status information, typically used for EXERCISE
	// identity amplification (FAKER, JOKER, KILO, TRAVELLER, ZOMBIE):
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// KILO: Friendly high-value object.
	//
	// TRAVELLER: Suspect land or surface track following a recognized traffic route.
	//
	// ZOMBIE: Suspect track, object, or entity of special interest.
	IdentAmp param.Opt[string] `json:"identAmp,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// A text amplifier displaying IFF/SIF/AIS Identification modes and codes.
	Iff param.Opt[string] `json:"iff,omitzero"`
	// Flag indicating that this track represents an installation. Based on
	// MIL-STD-2525 symbology definitions.
	Installation param.Opt[bool] `json:"installation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by observation source to indicate the target
	// on-orbit object of this track. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Radio frequency of the track, measured in megahertz (MHz).
	RdfRf param.Opt[float64] `json:"rdfRF,omitzero"`
	// Flag indicating that this track represents a reduced object or group. Based on
	// MIL-STD-2525 symbology definitions.
	Reduced param.Opt[bool] `json:"reduced,omitzero"`
	// Flag indicating that this track represents a reinforced object or group. Based
	// on MIL-STD-2525 symbology definitions.
	Reinforced param.Opt[bool] `json:"reinforced,omitzero"`
	// Report number received from the reporting source for this track.
	RptNum param.Opt[string] `json:"rptNum,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Status of the satellite.
	SatStatus param.Opt[string] `json:"satStatus,omitzero"`
	// Track object speed, in km/sec.
	Spd param.Opt[float64] `json:"spd,omitzero"`
	// Flag indicating that this track represents a task force. Based on MIL-STD-2525
	// symbology definitions.
	TaskForce param.Opt[bool] `json:"taskForce,omitzero"`
	// UUID identifying the track, which should remain the same on subsequent tracks of
	// the same object.
	TrkID param.Opt[string] `json:"trkId,omitzero"`
	// The type of vehicle with which the device is associated. Based on MIL-STD-2525
	// symbology definitions.
	VehType param.Opt[string] `json:"vehType,omitzero"`
	// Source cross-reference code for the command that originated the track report.
	Xref param.Opt[string] `json:"xref,omitzero"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// ELLIPSE:
	//
	// brg - orientation in degrees of the ellipse
	//
	// a1 - semi-major axis in meters
	//
	// a2 - semi-minor axis in meters
	//
	// BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// brg - orientation in degrees of the bearing box
	//
	// a1 - length of bearing box in meters
	//
	// a2 - half-width of bearing box in meters
	//
	// OTHER (All other type values):
	//
	// brg - line of bearing in degrees true
	//
	// a1 - bearing error in degrees
	//
	// a2 - estimated range in meters.
	AouData []float64 `json:"aouData,omitzero"`
	// The on-orbit category assigned to this track object (DEBRIS, MANNED, PAYLOAD,
	// PLATFORM, ROCKET BODY, UNKNOWN).
	//
	// Any of "DEBRIS", "MANNED", "PAYLOAD", "PLATFORM", "ROCKET BODY", "UNKNOWN".
	ObjectType string `json:"objectType,omitzero"`
	// The estimated identity of the track object (ASSUMED FRIEND, FRIEND, HOSTILE,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	//
	// Any of "ASSUMED FRIEND", "FRIEND", "HOSTILE", "NEUTRAL", "PENDING", "SUSPECT",
	// "UNKNOWN".
	ObjIdent string `json:"objIdent,omitzero"`
	// TrackSensor Collection.
	TrackSensors []OrbittrackNewBulkParamsBodyTrackSensor `json:"trackSensors,omitzero"`
	paramObj
}

func (r OrbittrackNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OrbittrackNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OrbittrackNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[OrbittrackNewBulkParamsBody](
		"ObjectType", false, "DEBRIS", "MANNED", "PAYLOAD", "PLATFORM", "ROCKET BODY", "UNKNOWN",
	)
	apijson.RegisterFieldValidator[OrbittrackNewBulkParamsBody](
		"ObjIdent", false, "ASSUMED FRIEND", "FRIEND", "HOSTILE", "NEUTRAL", "PENDING", "SUSPECT", "UNKNOWN",
	)
}

// Schema for Track Sensor data.
//
// The properties Az, Range are required.
type OrbittrackNewBulkParamsBodyTrackSensor struct {
	// The observing sensor azimuth angle, in degrees and topocentric frame.
	Az float64 `json:"az,required"`
	// The track object range from the observing sensor, in kilometers.
	Range float64 `json:"range,required"`
	// Minimum range measurement capability of the sensor, in kilometers.
	MinRangeLimit param.Opt[float64] `json:"minRangeLimit,omitzero"`
	// The mission number which produced this track observation.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// Unique name of this sensor.
	SensorName param.Opt[string] `json:"sensorName,omitzero"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber param.Opt[int64] `json:"sensorNumber,omitzero"`
	// The field of view (FOV) type (Butterfly, Cone Angular, Cone Distance, Horizon to
	// Horizon, Unknown) employed by the sensor observing this object.
	//
	// Any of "BUTTERFLY", "CONE ANGULAR", "CONE DISTANCE", "HORIZON TO HORIZON",
	// "UNKNOWN".
	SensorFovType string `json:"sensorFOVType,omitzero"`
	paramObj
}

func (r OrbittrackNewBulkParamsBodyTrackSensor) MarshalJSON() (data []byte, err error) {
	type shadow OrbittrackNewBulkParamsBodyTrackSensor
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OrbittrackNewBulkParamsBodyTrackSensor](
		"SensorFovType", false, "BUTTERFLY", "CONE ANGULAR", "CONE DISTANCE", "HORIZON TO HORIZON", "UNKNOWN",
	)
}

type OrbittrackTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrbittrackTupleParams]'s query parameters as `url.Values`.
func (r OrbittrackTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrbittrackUnvalidatedPublishParams struct {
	Body []OrbittrackUnvalidatedPublishParamsBody
	paramObj
}

func (r OrbittrackUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Keplerian orbital elements describing an orbit for a particular on-orbit
// satellite and applicable sensor data aiding in the orbit prediction.
//
// The properties ClassificationMarking, DataMode, Lat, Lon, Source, Ts are
// required.
type OrbittrackUnvalidatedPublishParamsBody struct {
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
	// WGS-84 latitude of the track object subpoint, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	Lat float64 `json:"lat,required"`
	// WGS-84 longitude of the track object subpoint, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Track point altitude relative to WGS-84 ellipsoid, in meters.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Free-form remarks entered for the satellite.
	Amplification param.Opt[string] `json:"amplification,omitzero"`
	// The angle formed between the line of sight of the observer and the horizon at
	// track timestamp, in degrees. The angular range is -90 to 90, with negative
	// values representing angle of depression.
	AngElev param.Opt[float64] `json:"angElev,omitzero"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouData array and is required if aouData is not
	// null. See the aouData field definition for specific information.
	AouType param.Opt[string] `json:"aouType,omitzero"`
	// International radio call sign assigned to the track. This is an 8-character
	// alphanumeric code.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// One-line Charlie elements set.
	CharlieLine param.Opt[string] `json:"charlieLine,omitzero"`
	// The cross-reference code of the channel on which this track report was received,
	// if the report came over a comms channel.
	ChXRef param.Opt[string] `json:"chXRef,omitzero"`
	// The Area Of Uncertainty (AOU) percentage (0 - 100) containment value. The
	// percentage of time (90%) that the estimated area of uncertainty will cover the
	// true position of the track object.
	Cntnmnt param.Opt[float64] `json:"cntnmnt,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Time the row was created in the database.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Predicted change in Mean Motion (velocity) in radians/herg^2. herg is a unit of
	// time measure equal to 806.8120769 seconds, and is the orbital period of an
	// imaginary satellite rotating about the Earth at zero altitude.
	Decay param.Opt[float64] `json:"decay,omitzero"`
	// Flag indicating that this track represents a dummy object or group. Identifies
	// offensive or defensive units, equipment and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Dummy param.Opt[bool] `json:"dummy,omitzero"`
	// Flag indicating that this track represents a feint object or group. Identifies
	// offensive or defensive units, equipment and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Feint param.Opt[bool] `json:"feint,omitzero"`
	// Flag indicating that this track represents a headquarters object. Based on
	// MIL-STD-2525 symbology definitions.
	Hq param.Opt[bool] `json:"hq,omitzero"`
	// Unique identifier of the Elset associated with this object.
	IDElset param.Opt[string] `json:"idElset,omitzero"`
	// Additional track object identity/status information, typically used for EXERCISE
	// identity amplification (FAKER, JOKER, KILO, TRAVELLER, ZOMBIE):
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// KILO: Friendly high-value object.
	//
	// TRAVELLER: Suspect land or surface track following a recognized traffic route.
	//
	// ZOMBIE: Suspect track, object, or entity of special interest.
	IdentAmp param.Opt[string] `json:"identAmp,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// A text amplifier displaying IFF/SIF/AIS Identification modes and codes.
	Iff param.Opt[string] `json:"iff,omitzero"`
	// Flag indicating that this track represents an installation. Based on
	// MIL-STD-2525 symbology definitions.
	Installation param.Opt[bool] `json:"installation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by observation source to indicate the target
	// on-orbit object of this track. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Radio frequency of the track, measured in megahertz (MHz).
	RdfRf param.Opt[float64] `json:"rdfRF,omitzero"`
	// Flag indicating that this track represents a reduced object or group. Based on
	// MIL-STD-2525 symbology definitions.
	Reduced param.Opt[bool] `json:"reduced,omitzero"`
	// Flag indicating that this track represents a reinforced object or group. Based
	// on MIL-STD-2525 symbology definitions.
	Reinforced param.Opt[bool] `json:"reinforced,omitzero"`
	// Report number received from the reporting source for this track.
	RptNum param.Opt[string] `json:"rptNum,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Status of the satellite.
	SatStatus param.Opt[string] `json:"satStatus,omitzero"`
	// Track object speed, in km/sec.
	Spd param.Opt[float64] `json:"spd,omitzero"`
	// Flag indicating that this track represents a task force. Based on MIL-STD-2525
	// symbology definitions.
	TaskForce param.Opt[bool] `json:"taskForce,omitzero"`
	// UUID identifying the track, which should remain the same on subsequent tracks of
	// the same object.
	TrkID param.Opt[string] `json:"trkId,omitzero"`
	// The type of vehicle with which the device is associated. Based on MIL-STD-2525
	// symbology definitions.
	VehType param.Opt[string] `json:"vehType,omitzero"`
	// Source cross-reference code for the command that originated the track report.
	Xref param.Opt[string] `json:"xref,omitzero"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// ELLIPSE:
	//
	// brg - orientation in degrees of the ellipse
	//
	// a1 - semi-major axis in meters
	//
	// a2 - semi-minor axis in meters
	//
	// BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// brg - orientation in degrees of the bearing box
	//
	// a1 - length of bearing box in meters
	//
	// a2 - half-width of bearing box in meters
	//
	// OTHER (All other type values):
	//
	// brg - line of bearing in degrees true
	//
	// a1 - bearing error in degrees
	//
	// a2 - estimated range in meters.
	AouData []float64 `json:"aouData,omitzero"`
	// The on-orbit category assigned to this track object (DEBRIS, MANNED, PAYLOAD,
	// PLATFORM, ROCKET BODY, UNKNOWN).
	//
	// Any of "DEBRIS", "MANNED", "PAYLOAD", "PLATFORM", "ROCKET BODY", "UNKNOWN".
	ObjectType string `json:"objectType,omitzero"`
	// The estimated identity of the track object (ASSUMED FRIEND, FRIEND, HOSTILE,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	//
	// Any of "ASSUMED FRIEND", "FRIEND", "HOSTILE", "NEUTRAL", "PENDING", "SUSPECT",
	// "UNKNOWN".
	ObjIdent string `json:"objIdent,omitzero"`
	// TrackSensor Collection.
	TrackSensors []OrbittrackUnvalidatedPublishParamsBodyTrackSensor `json:"trackSensors,omitzero"`
	paramObj
}

func (r OrbittrackUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OrbittrackUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OrbittrackUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[OrbittrackUnvalidatedPublishParamsBody](
		"ObjectType", false, "DEBRIS", "MANNED", "PAYLOAD", "PLATFORM", "ROCKET BODY", "UNKNOWN",
	)
	apijson.RegisterFieldValidator[OrbittrackUnvalidatedPublishParamsBody](
		"ObjIdent", false, "ASSUMED FRIEND", "FRIEND", "HOSTILE", "NEUTRAL", "PENDING", "SUSPECT", "UNKNOWN",
	)
}

// Schema for Track Sensor data.
//
// The properties Az, Range are required.
type OrbittrackUnvalidatedPublishParamsBodyTrackSensor struct {
	// The observing sensor azimuth angle, in degrees and topocentric frame.
	Az float64 `json:"az,required"`
	// The track object range from the observing sensor, in kilometers.
	Range float64 `json:"range,required"`
	// Minimum range measurement capability of the sensor, in kilometers.
	MinRangeLimit param.Opt[float64] `json:"minRangeLimit,omitzero"`
	// The mission number which produced this track observation.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// Unique name of this sensor.
	SensorName param.Opt[string] `json:"sensorName,omitzero"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber param.Opt[int64] `json:"sensorNumber,omitzero"`
	// The field of view (FOV) type (Butterfly, Cone Angular, Cone Distance, Horizon to
	// Horizon, Unknown) employed by the sensor observing this object.
	//
	// Any of "BUTTERFLY", "CONE ANGULAR", "CONE DISTANCE", "HORIZON TO HORIZON",
	// "UNKNOWN".
	SensorFovType string `json:"sensorFOVType,omitzero"`
	paramObj
}

func (r OrbittrackUnvalidatedPublishParamsBodyTrackSensor) MarshalJSON() (data []byte, err error) {
	type shadow OrbittrackUnvalidatedPublishParamsBodyTrackSensor
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OrbittrackUnvalidatedPublishParamsBodyTrackSensor](
		"SensorFovType", false, "BUTTERFLY", "CONE ANGULAR", "CONE DISTANCE", "HORIZON TO HORIZON", "UNKNOWN",
	)
}
