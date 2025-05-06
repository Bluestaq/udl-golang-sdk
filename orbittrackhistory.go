// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// OrbittrackHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrbittrackHistoryService] method instead.
type OrbittrackHistoryService struct {
	Options []option.RequestOption
}

// NewOrbittrackHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrbittrackHistoryService(opts ...option.RequestOption) (r OrbittrackHistoryService) {
	r = OrbittrackHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OrbittrackHistoryService) List(ctx context.Context, query OrbittrackHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OrbittrackHistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/orbittrack/history"
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

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OrbittrackHistoryService) ListAutoPaging(ctx context.Context, query OrbittrackHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OrbittrackHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OrbittrackHistoryService) Aodr(ctx context.Context, query OrbittrackHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/orbittrack/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OrbittrackHistoryService) Count(ctx context.Context, query OrbittrackHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/orbittrack/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Keplerian orbital elements describing an orbit for a particular on-orbit
// satellite and applicable sensor data aiding in the orbit prediction.
type OrbittrackHistoryListResponse struct {
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
	DataMode OrbittrackHistoryListResponseDataMode `json:"dataMode,required"`
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
	ObjectType OrbittrackHistoryListResponseObjectType `json:"objectType"`
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
	ObjIdent OrbittrackHistoryListResponseObjIdent `json:"objIdent"`
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
	TrackSensors []OrbittrackHistoryListResponseTrackSensor `json:"trackSensors"`
	// UUID identifying the track, which should remain the same on subsequent tracks of
	// the same object.
	TrkID string `json:"trkId"`
	// The type of vehicle with which the device is associated. Based on MIL-STD-2525
	// symbology definitions.
	VehType string `json:"vehType"`
	// Source cross-reference code for the command that originated the track report.
	Xref string `json:"xref"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		Alt                   respjson.Field
		Amplification         respjson.Field
		AngElev               respjson.Field
		AouData               respjson.Field
		AouType               respjson.Field
		CallSign              respjson.Field
		CharlieLine           respjson.Field
		ChXRef                respjson.Field
		Cntnmnt               respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Decay                 respjson.Field
		Dummy                 respjson.Field
		Feint                 respjson.Field
		Hq                    respjson.Field
		IDElset               respjson.Field
		IdentAmp              respjson.Field
		IDOnOrbit             respjson.Field
		Iff                   respjson.Field
		Installation          respjson.Field
		ObjectType            respjson.Field
		ObjIdent              respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		RdfRf                 respjson.Field
		Reduced               respjson.Field
		Reinforced            respjson.Field
		RptNum                respjson.Field
		SatNo                 respjson.Field
		SatStatus             respjson.Field
		Spd                   respjson.Field
		TaskForce             respjson.Field
		TrackSensors          respjson.Field
		TrkID                 respjson.Field
		VehType               respjson.Field
		Xref                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrbittrackHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrbittrackHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type OrbittrackHistoryListResponseDataMode string

const (
	OrbittrackHistoryListResponseDataModeReal      OrbittrackHistoryListResponseDataMode = "REAL"
	OrbittrackHistoryListResponseDataModeTest      OrbittrackHistoryListResponseDataMode = "TEST"
	OrbittrackHistoryListResponseDataModeSimulated OrbittrackHistoryListResponseDataMode = "SIMULATED"
	OrbittrackHistoryListResponseDataModeExercise  OrbittrackHistoryListResponseDataMode = "EXERCISE"
)

// The on-orbit category assigned to this track object (DEBRIS, MANNED, PAYLOAD,
// PLATFORM, ROCKET BODY, UNKNOWN).
type OrbittrackHistoryListResponseObjectType string

const (
	OrbittrackHistoryListResponseObjectTypeDebris     OrbittrackHistoryListResponseObjectType = "DEBRIS"
	OrbittrackHistoryListResponseObjectTypeManned     OrbittrackHistoryListResponseObjectType = "MANNED"
	OrbittrackHistoryListResponseObjectTypePayload    OrbittrackHistoryListResponseObjectType = "PAYLOAD"
	OrbittrackHistoryListResponseObjectTypePlatform   OrbittrackHistoryListResponseObjectType = "PLATFORM"
	OrbittrackHistoryListResponseObjectTypeRocketBody OrbittrackHistoryListResponseObjectType = "ROCKET BODY"
	OrbittrackHistoryListResponseObjectTypeUnknown    OrbittrackHistoryListResponseObjectType = "UNKNOWN"
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
type OrbittrackHistoryListResponseObjIdent string

const (
	OrbittrackHistoryListResponseObjIdentAssumedFriend OrbittrackHistoryListResponseObjIdent = "ASSUMED FRIEND"
	OrbittrackHistoryListResponseObjIdentFriend        OrbittrackHistoryListResponseObjIdent = "FRIEND"
	OrbittrackHistoryListResponseObjIdentHostile       OrbittrackHistoryListResponseObjIdent = "HOSTILE"
	OrbittrackHistoryListResponseObjIdentNeutral       OrbittrackHistoryListResponseObjIdent = "NEUTRAL"
	OrbittrackHistoryListResponseObjIdentPending       OrbittrackHistoryListResponseObjIdent = "PENDING"
	OrbittrackHistoryListResponseObjIdentSuspect       OrbittrackHistoryListResponseObjIdent = "SUSPECT"
	OrbittrackHistoryListResponseObjIdentUnknown       OrbittrackHistoryListResponseObjIdent = "UNKNOWN"
)

// Schema for Track Sensor data.
type OrbittrackHistoryListResponseTrackSensor struct {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Az            respjson.Field
		Range         respjson.Field
		MinRangeLimit respjson.Field
		MissionNumber respjson.Field
		SensorFovType respjson.Field
		SensorName    respjson.Field
		SensorNumber  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrbittrackHistoryListResponseTrackSensor) RawJSON() string { return r.JSON.raw }
func (r *OrbittrackHistoryListResponseTrackSensor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrbittrackHistoryListParams struct {
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts time.Time `query:"ts,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrbittrackHistoryListParams]'s query parameters as
// `url.Values`.
func (r OrbittrackHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrbittrackHistoryAodrParams struct {
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts time.Time `query:"ts,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// optional, notification method for the created file link. When omitted, EMAIL is
	// assumed. Current valid values are: EMAIL, SMS.
	Notification param.Opt[string] `query:"notification,omitzero" json:"-"`
	// optional, field delimiter when the created file is not JSON. Must be a single
	// character chosen from this set: (',', ';', ':', '|'). When omitted, "," is used.
	// It is strongly encouraged that your field delimiter be a character unlikely to
	// occur within the data.
	OutputDelimiter param.Opt[string] `query:"outputDelimiter,omitzero" json:"-"`
	// optional, output format for the file. When omitted, JSON is assumed. Current
	// valid values are: JSON and CSV.
	OutputFormat param.Opt[string] `query:"outputFormat,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrbittrackHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r OrbittrackHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrbittrackHistoryCountParams struct {
	// Track timestamp in ISO8601 UTC format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrbittrackHistoryCountParams]'s query parameters as
// `url.Values`.
func (r OrbittrackHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
