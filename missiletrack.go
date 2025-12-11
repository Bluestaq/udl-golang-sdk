// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// MissileTrackService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMissileTrackService] method instead.
type MissileTrackService struct {
	Options []option.RequestOption
	History MissileTrackHistoryService
}

// NewMissileTrackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMissileTrackService(opts ...option.RequestOption) (r MissileTrackService) {
	r = MissileTrackService{}
	r.Options = opts
	r.History = NewMissileTrackHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *MissileTrackService) List(ctx context.Context, query MissileTrackListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[MissileTrackListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/missiletrack"
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
func (r *MissileTrackService) ListAutoPaging(ctx context.Context, query MissileTrackListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[MissileTrackListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *MissileTrackService) Count(ctx context.Context, query MissileTrackCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/missiletrack/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// MissileTrack records as a POST body and ingest into the database. This operation
// is not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *MissileTrackService) NewBulk(ctx context.Context, body MissileTrackNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/missiletrack/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *MissileTrackService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *MissileTrackQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/missiletrack/queryhelp"
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
func (r *MissileTrackService) Tuple(ctx context.Context, query MissileTrackTupleParams, opts ...option.RequestOption) (res *[]MissileTrackTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/missiletrack/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple missile track records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *MissileTrackService) UnvalidatedPublish(ctx context.Context, body MissileTrackUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-missiletrack"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// These services provide operations for querying of all available missile track
// details and amplifying missile data. A missile track is a position and
// optionally a heading/velocity of an object across all environments at a
// particular timestamp. It also includes optional information regarding the
// identity/type of missile, impact location, launch location and other amplifying
// object data, if known.
type MissileTrackListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode MissileTrackListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The receipt time of the data by the processing system, in ISO8601 UTC format
	// with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Subtype is a finer grain categorization of missile types. Examples include but
	// are not limited to SRBM, MRBM, IRBM, LRBM, ICBM, SLBM.
	//
	// &nbsp;SRBM - Short-Range Ballistic Missile
	//
	// &nbsp;MRBM - Medium-Range Ballistic Missile
	//
	// &nbsp;IRBM - Intermediate-Range Ballistic Missile
	//
	// &nbsp;LRBM - Long-Range Ballistic Missile
	//
	// &nbsp;ICBM - Intercontinental Ballistic Missile
	//
	// &nbsp;SLBM - Submarine-Launched Ballistic Missile.
	AcftSubType string `json:"acftSubType"`
	// A track may be designated as a non-alert track or an alert track.
	//
	// Examples include but are not limited to:
	//
	// &nbsp;Non-alert tracks – choose None (Blank).
	//
	// &nbsp;Alert tracks – enter the proper alert classification:
	//
	// &nbsp;HIT - High Interest Track
	//
	// &nbsp;TGT - Target
	//
	// &nbsp;SUS - Suspect Carrier
	//
	// &nbsp;NSP - Cleared Suspect.
	Alert string `json:"alert"`
	// Angle of elevation/depression between observer and missile in degrees.
	AngElev float64 `json:"angElev"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	AouRptData []float64 `json:"aouRptData"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	AouRptType string `json:"aouRptType"`
	// Missile azimuth corridor data.
	AzCorr float64 `json:"azCorr"`
	// Indicates whether or not the missile is currently in a state of boosting.
	Boosting bool `json:"boosting"`
	// Track point burnout altitude relative to WGS-84 ellipsoid, in kilometers.
	BurnoutAlt float64 `json:"burnoutAlt"`
	// The call sign currently assigned to the track object.
	CallSign string `json:"callSign"`
	// The percentage of time that the estimated AoU will "cover" the true position of
	// the track.
	Containment float64 `json:"containment"`
	// An optional string array containing additional data (keys) representing relevant
	// items for context of fields not specifically defined in this schema. This array
	// is paired with the contextValues string array and must contain the same number
	// of items. Please note these fields are intended for contextual use only and do
	// not pertain to core schema information. To ensure proper integration and avoid
	// misuse, coordination of how these fields are populated and consumed is required
	// during onboarding.
	ContextKeys []string `json:"contextKeys"`
	// An optional string array containing the values associated with the contextKeys
	// array. This array is paired with the contextKeys string array and must contain
	// the same number of items. Please note these fields are intended for contextual
	// use only and do not pertain to core schema information. To ensure proper
	// integration and avoid misuse, coordination of how these fields are populated and
	// consumed is required during onboarding.
	ContextValues []string `json:"contextValues"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// The drop-point indicator setting.
	DropPtInd bool `json:"dropPtInd"`
	// Indicates whether or not a track has an emergency.
	EmgInd bool `json:"emgInd"`
	// The track environment type (AIR, LAND, SPACE, SUBSURFACE, SURFACE, UNKNOWN):
	//
	// AIR: Between sea level and the Kármán line, which has an altitude of 100
	// kilometers (62 miles).
	//
	// LAND: On the surface of dry land.
	//
	// SPACE: Above the Kármán line, which has an altitude of 100 kilometers (62
	// miles).
	//
	// SURFACE: On the surface of a body of water.
	//
	// SUBSURFACE: Below the surface of a body of water.
	//
	// UNKNOWN: Environment is not known.
	//
	// Any of "AIR", "LAND", "SPACE", "SURFACE", "SUBSURFACE", "UNKNOWN".
	Env MissileTrackListResponseEnv `json:"env"`
	// Estimated impact point altitude relative to WGS-84 ellipsoid, in kilometers.
	ImpactAlt float64 `json:"impactAlt"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	ImpactAouData []float64 `json:"impactAouData"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	ImpactAouType string `json:"impactAouType"`
	// Confidence level of the impact point estimate. 0 - 100 percent.
	ImpactConf float64 `json:"impactConf"`
	// WGS-84 latitude of the missile object impact point, in degrees. -90 to 90
	// degrees (negative values south of equator).
	ImpactLat float64 `json:"impactLat"`
	// WGS-84 longitude of the missile object impact point, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	ImpactLon float64 `json:"impactLon"`
	// Estimated time of impact timestamp in ISO8601 UTC format with microsecond
	// precision.
	ImpactTime time.Time `json:"impactTime" format:"date-time"`
	// Source code for source of information used to detect track.
	InfoSource string `json:"infoSource"`
	// Estimated launch point altitude relative to WGS-84 ellipsoid, in kilometers.
	LaunchAlt float64 `json:"launchAlt"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	LaunchAouData []float64 `json:"launchAouData"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	LaunchAouType string `json:"launchAouType"`
	// Angle between true north and the object's current position, with respect to the
	// launch point, in degrees. 0 to 360 degrees.
	LaunchAz float64 `json:"launchAz"`
	// Uncertainty of the launch azimuth, in degrees.
	LaunchAzUnc float64 `json:"launchAzUnc"`
	// Confidence level in the accuracy of the launch point estimate. 0 - 100 percent.
	LaunchConf float64 `json:"launchConf"`
	// WGS-84 latitude of the missile launch point, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	LaunchLat float64 `json:"launchLat"`
	// WGS-84 longitude of the missile launch point, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	LaunchLon float64 `json:"launchLon"`
	// Missile launch timestamp in ISO8601 UTC format with microsecond precision.
	LaunchTime time.Time `json:"launchTime" format:"date-time"`
	// Indicates whether or not the missile is considered lost.
	LostTrkInd bool `json:"lostTrkInd"`
	// Maneuver end time, in ISO 8601 UTC format with microsecond precision.
	ManeuverEnd time.Time `json:"maneuverEnd" format:"date-time"`
	// Maneuver start time, in ISO 8601 UTC format with microsecond precision.
	ManeuverStart time.Time `json:"maneuverStart" format:"date-time"`
	// The timestamp of the external message from which this request originated, if
	// applicable, in ISO8601 UTC format with millisecond precision.
	MsgCreateDate time.Time `json:"msgCreateDate" format:"date-time"`
	// The message subtype is a finer grain categorization of message types as many
	// messages can contain a variety of data content within the same structure.
	// Examples include but are not limited to Initial, Final, Launch, Update, etc.
	// Users should consult the appropriate documentation, based on the message type,
	// for the definitions of the subtypes that apply to that message.
	MsgSubType string `json:"msgSubType"`
	// The type of external message from which this request originated.
	MsgType string `json:"msgType"`
	// Missile status enumeration examples include but are not limited to:
	//
	// &nbsp;AT LAUNCH
	//
	// &nbsp;AT OBSERVATION
	//
	// &nbsp;FLYING
	//
	// &nbsp;IMPACTED
	//
	// &nbsp;LOST
	//
	// &nbsp;STALE
	//
	// &nbsp;DEBRIS.
	MslStatus string `json:"mslStatus"`
	// Source of the missile-unique identifier (MUID).
	MuidSrc string `json:"muidSrc"`
	// Track ID for the source of the missile-unique identifier.
	MuidSrcTrk string `json:"muidSrcTrk"`
	// Track name.
	Name string `json:"name"`
	// Space activity (examples: RECONNAISSANCE, ANTISPACE WARFARE, TELEVISION). The
	// activity in which the track object is engaged. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent activity designations. The activity can
	// be reported as either a combination of the code and environment (e.g. 65/AIR) or
	// as the descriptive enumeration (e.g. DIVERTING), which are equivalent.
	ObjAct string `json:"objAct"`
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
	ObjIdent MissileTrackListResponseObjIdent `json:"objIdent"`
	// Space Platform field along with the Space Activity field further defines the
	// identity of a Space track (examples: SATELLITE, WEAPON, PATROL). The object
	// platform type. Intended as, but not constrained to, MIL-STD-6016 environment
	// dependent platform type designations. The platform type can be reported as
	// either a combination of the code and environment (e.g. 14/LAND) or as the
	// descriptive representations (e.g. COMBAT VEHICLE), which are equivalent.
	ObjPlat string `json:"objPlat"`
	// The type of object to which this record refers. The object type may be updated
	// in later records based on assessment of additional data.
	ObjType string `json:"objType"`
	// Confidence of the object type, 0-100.
	ObjTypeConf int64 `json:"objTypeConf"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Track ID of the parent track, within the originating system, from which the
	// track was developed.
	ParentTrackID string `json:"parentTrackId"`
	// Azimuth corridor reference point latitude.
	PolarSingLocLat float64 `json:"polarSingLocLat"`
	// Azimuth corridor reference point longitude.
	PolarSingLocLon float64 `json:"polarSingLocLon"`
	// Last report type received from the sensor (for example, OBSBO = observation
	// burnout).
	SenMode string `json:"senMode"`
	// Space amplification indicates additional information on the space environment
	// being reported (examples: NUCLEAR WARHEAD, FUEL-AIR EXPLOSIVE WARHEAD, DEBRIS).
	SpaceAmp string `json:"spaceAmp"`
	// Confidence level of the amplifying characteristics. Values range from 0 to 6.
	SpaceAmpConf int64 `json:"spaceAmpConf"`
	// Specific type of point or track with an environment of space.
	SpaceSpecType string `json:"spaceSpecType"`
	// Track ID within the originating system.
	TrackID string `json:"trackId"`
	// Overall track confidence estimate (not standardized, but typically a value
	// between 0 and 1, with 0 indicating lowest confidence).
	TrkConf float64 `json:"trkConf"`
	// Track Quality is reported as an integer from 0-15. Track Quality specifies the
	// reliability of the positional information of a reported track, with higher
	// values indicating higher track quality; i.e., lower errors in reported position.
	TrkQual int64 `json:"trkQual"`
	// Array of MissileTrackVector objects. Missile track vectors are cartesian vectors
	// of position, velocity, and acceleration that, together with their time, 'epoch',
	// uniquely determine the trajectory of the missile. ECEF is the preferred
	// coordinate frame but in some cases data may be in another frame as specified by
	// 'referenceFrame', depending on the provider.
	Vectors []MissileTrackListResponseVector `json:"vectors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		AcftSubType           respjson.Field
		Alert                 respjson.Field
		AngElev               respjson.Field
		AouRptData            respjson.Field
		AouRptType            respjson.Field
		AzCorr                respjson.Field
		Boosting              respjson.Field
		BurnoutAlt            respjson.Field
		CallSign              respjson.Field
		Containment           respjson.Field
		ContextKeys           respjson.Field
		ContextValues         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DropPtInd             respjson.Field
		EmgInd                respjson.Field
		Env                   respjson.Field
		ImpactAlt             respjson.Field
		ImpactAouData         respjson.Field
		ImpactAouType         respjson.Field
		ImpactConf            respjson.Field
		ImpactLat             respjson.Field
		ImpactLon             respjson.Field
		ImpactTime            respjson.Field
		InfoSource            respjson.Field
		LaunchAlt             respjson.Field
		LaunchAouData         respjson.Field
		LaunchAouType         respjson.Field
		LaunchAz              respjson.Field
		LaunchAzUnc           respjson.Field
		LaunchConf            respjson.Field
		LaunchLat             respjson.Field
		LaunchLon             respjson.Field
		LaunchTime            respjson.Field
		LostTrkInd            respjson.Field
		ManeuverEnd           respjson.Field
		ManeuverStart         respjson.Field
		MsgCreateDate         respjson.Field
		MsgSubType            respjson.Field
		MsgType               respjson.Field
		MslStatus             respjson.Field
		MuidSrc               respjson.Field
		MuidSrcTrk            respjson.Field
		Name                  respjson.Field
		ObjAct                respjson.Field
		ObjIdent              respjson.Field
		ObjPlat               respjson.Field
		ObjType               respjson.Field
		ObjTypeConf           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ParentTrackID         respjson.Field
		PolarSingLocLat       respjson.Field
		PolarSingLocLon       respjson.Field
		SenMode               respjson.Field
		SpaceAmp              respjson.Field
		SpaceAmpConf          respjson.Field
		SpaceSpecType         respjson.Field
		TrackID               respjson.Field
		TrkConf               respjson.Field
		TrkQual               respjson.Field
		Vectors               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MissileTrackListResponse) RawJSON() string { return r.JSON.raw }
func (r *MissileTrackListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type MissileTrackListResponseDataMode string

const (
	MissileTrackListResponseDataModeReal      MissileTrackListResponseDataMode = "REAL"
	MissileTrackListResponseDataModeTest      MissileTrackListResponseDataMode = "TEST"
	MissileTrackListResponseDataModeSimulated MissileTrackListResponseDataMode = "SIMULATED"
	MissileTrackListResponseDataModeExercise  MissileTrackListResponseDataMode = "EXERCISE"
)

// The track environment type (AIR, LAND, SPACE, SUBSURFACE, SURFACE, UNKNOWN):
//
// AIR: Between sea level and the Kármán line, which has an altitude of 100
// kilometers (62 miles).
//
// LAND: On the surface of dry land.
//
// SPACE: Above the Kármán line, which has an altitude of 100 kilometers (62
// miles).
//
// SURFACE: On the surface of a body of water.
//
// SUBSURFACE: Below the surface of a body of water.
//
// UNKNOWN: Environment is not known.
type MissileTrackListResponseEnv string

const (
	MissileTrackListResponseEnvAir        MissileTrackListResponseEnv = "AIR"
	MissileTrackListResponseEnvLand       MissileTrackListResponseEnv = "LAND"
	MissileTrackListResponseEnvSpace      MissileTrackListResponseEnv = "SPACE"
	MissileTrackListResponseEnvSurface    MissileTrackListResponseEnv = "SURFACE"
	MissileTrackListResponseEnvSubsurface MissileTrackListResponseEnv = "SUBSURFACE"
	MissileTrackListResponseEnvUnknown    MissileTrackListResponseEnv = "UNKNOWN"
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
type MissileTrackListResponseObjIdent string

const (
	MissileTrackListResponseObjIdentAssumedFriend MissileTrackListResponseObjIdent = "ASSUMED FRIEND"
	MissileTrackListResponseObjIdentFriend        MissileTrackListResponseObjIdent = "FRIEND"
	MissileTrackListResponseObjIdentHostile       MissileTrackListResponseObjIdent = "HOSTILE"
	MissileTrackListResponseObjIdentNeutral       MissileTrackListResponseObjIdent = "NEUTRAL"
	MissileTrackListResponseObjIdentPending       MissileTrackListResponseObjIdent = "PENDING"
	MissileTrackListResponseObjIdentSuspect       MissileTrackListResponseObjIdent = "SUSPECT"
	MissileTrackListResponseObjIdentUnknown       MissileTrackListResponseObjIdent = "UNKNOWN"
)

// Schema for Missile Track Vector data.
type MissileTrackListResponseVector struct {
	// Vector timestamp in ISO8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Three element array, expressing the cartesian acceleration vector of the target
	// object, in kilometers/second^2, in the specified referenceFrame. If
	// referenceFrame is null then ECEF should be assumed. The array element order is
	// [x”, y”, z”].
	Accel []float64 `json:"accel"`
	// Confidence of the vector, 0-100.
	Confidence int64 `json:"confidence"`
	// An optional string array containing additional data (keys) representing relevant
	// items for context of fields not specifically defined in this schema. This array
	// is paired with the contextValues string array and must contain the same number
	// of items. Please note these fields are intended for contextual use only and do
	// not pertain to core schema information. To ensure proper integration and avoid
	// misuse, coordination of how these fields are populated and consumed is required
	// during onboarding.
	ContextKeys []string `json:"contextKeys"`
	// An optional string array containing the values associated with the contextKeys
	// array. This array is paired with the contextKeys string array and must contain
	// the same number of items. Please note these fields are intended for contextual
	// use only and do not pertain to core schema information. To ensure proper
	// integration and avoid misuse, coordination of how these fields are populated and
	// consumed is required during onboarding.
	ContextValues []string `json:"contextValues"`
	// Track object course, in degrees clockwise from true north.
	Course float64 `json:"course"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame.
	//
	// If the covReferenceFrame is null it is assumed to be UVW. The array values
	// (1-45) represent the upper triangular half of the position-velocity-acceleration
	// covariance matrix.
	//
	// The covariance elements are position dependent within the array with values
	// ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x"&nbsp;&nbsp;&nbsp;&nbsp;y"&nbsp;&nbsp;&nbsp;&nbsp;z"
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;10&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;&nbsp;12&nbsp;&nbsp;&nbsp;13&nbsp;&nbsp;&nbsp;14&nbsp;&nbsp;&nbsp;15&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;&nbsp;17
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;18&nbsp;&nbsp;&nbsp;19&nbsp;&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;&nbsp;23&nbsp;&nbsp;&nbsp;24
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;25&nbsp;&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;&nbsp;30
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;31&nbsp;&nbsp;&nbsp;32&nbsp;&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;35
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;&nbsp;38&nbsp;&nbsp;&nbsp;39
	//
	// x"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;40&nbsp;&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42
	//
	// y"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;44
	//
	// z"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45
	//
	// The cov array should contain only the upper right triangle values from top left
	// down to bottom right, in order.
	Cov []float64 `json:"cov"`
	// The reference frame of the covariance elements (J2000, UVW, EFG/TDR, ECR/ECEF,
	// TEME, GCRF). If the referenceFrame is null it is assumed to be UVW.
	//
	// Any of "J2000", "UVW", "EFG/TDR", "ECR/ECEF", "TEME", "GCRF".
	CovReferenceFrame string `json:"covReferenceFrame"`
	// The flight azimuth associated with the current state vector (0-360 degrees).
	FlightAz float64 `json:"flightAz"`
	// Unique identifier of the reporting sensor of the object.
	IDSensor string `json:"idSensor"`
	// Object to which this vector applies.
	Object string `json:"object"`
	// Optional identifier provided by the source to indicate the reporting sensor of
	// the object. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Three element array, expressing the cartesian position vector of the target
	// object, in kilometers, in the specified referenceFrame. If referenceFrame is
	// null then ECEF should be assumed. The array element order is [x, y, z].
	Pos []float64 `json:"pos"`
	// Flag indicating whether the vector data was propagated.
	Propagated bool `json:"propagated"`
	// The quaternion describing the attitude of the spacecraft with respect to the
	// reference frame listed in the 'referenceFrame' field. The array element order
	// convention is the three vector components, followed by the scalar component.
	Quat []float64 `json:"quat"`
	// Range from the originating system or sensor to the object, in kilometers.
	Range float64 `json:"range"`
	// The reference frame of the cartesian vector (ECEF, J2000). If the referenceFrame
	// is null it is assumed to be ECEF.
	ReferenceFrame string `json:"referenceFrame"`
	// Track object speed, in kilometers/sec.
	Spd float64 `json:"spd"`
	// Status of the vector (e.g. INITIAL, UPDATE).
	Status string `json:"status"`
	// Source of the epoch time.
	TimeSource string `json:"timeSource"`
	// Type of vector represented (e.g. LOS, PREDICTED, STATE).
	Type string `json:"type"`
	// Object altitude at epoch, expressed in kilometers above WGS-84 ellipsoid.
	VectorAlt float64 `json:"vectorAlt"`
	// WGS-84 object latitude subpoint at epoch, represented as -90 to 90 degrees
	// (negative values south of equator).
	VectorLat float64 `json:"vectorLat"`
	// WGS-84 object longitude subpoint at epoch, represented as -180 to 180 degrees
	// (negative values west of Prime Meridian).
	VectorLon float64 `json:"vectorLon"`
	// Vector track ID within the originating system or sensor.
	VectorTrackID string `json:"vectorTrackId"`
	// Three element array, expressing the cartesian velocity vector of the target
	// object, in kilometers/second, in the specified referenceFrame. If referenceFrame
	// is null then ECEF should be assumed. The array element order is [x', y', z'].
	Vel []float64 `json:"vel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Epoch             respjson.Field
		Accel             respjson.Field
		Confidence        respjson.Field
		ContextKeys       respjson.Field
		ContextValues     respjson.Field
		Course            respjson.Field
		Cov               respjson.Field
		CovReferenceFrame respjson.Field
		FlightAz          respjson.Field
		IDSensor          respjson.Field
		Object            respjson.Field
		OrigSensorID      respjson.Field
		Pos               respjson.Field
		Propagated        respjson.Field
		Quat              respjson.Field
		Range             respjson.Field
		ReferenceFrame    respjson.Field
		Spd               respjson.Field
		Status            respjson.Field
		TimeSource        respjson.Field
		Type              respjson.Field
		VectorAlt         respjson.Field
		VectorLat         respjson.Field
		VectorLon         respjson.Field
		VectorTrackID     respjson.Field
		Vel               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MissileTrackListResponseVector) RawJSON() string { return r.JSON.raw }
func (r *MissileTrackListResponseVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MissileTrackQueryhelpResponse struct {
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
func (r MissileTrackQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *MissileTrackQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These services provide operations for querying of all available missile track
// details and amplifying missile data. A missile track is a position and
// optionally a heading/velocity of an object across all environments at a
// particular timestamp. It also includes optional information regarding the
// identity/type of missile, impact location, launch location and other amplifying
// object data, if known.
type MissileTrackTupleResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode MissileTrackTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The receipt time of the data by the processing system, in ISO8601 UTC format
	// with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Subtype is a finer grain categorization of missile types. Examples include but
	// are not limited to SRBM, MRBM, IRBM, LRBM, ICBM, SLBM.
	//
	// &nbsp;SRBM - Short-Range Ballistic Missile
	//
	// &nbsp;MRBM - Medium-Range Ballistic Missile
	//
	// &nbsp;IRBM - Intermediate-Range Ballistic Missile
	//
	// &nbsp;LRBM - Long-Range Ballistic Missile
	//
	// &nbsp;ICBM - Intercontinental Ballistic Missile
	//
	// &nbsp;SLBM - Submarine-Launched Ballistic Missile.
	AcftSubType string `json:"acftSubType"`
	// A track may be designated as a non-alert track or an alert track.
	//
	// Examples include but are not limited to:
	//
	// &nbsp;Non-alert tracks – choose None (Blank).
	//
	// &nbsp;Alert tracks – enter the proper alert classification:
	//
	// &nbsp;HIT - High Interest Track
	//
	// &nbsp;TGT - Target
	//
	// &nbsp;SUS - Suspect Carrier
	//
	// &nbsp;NSP - Cleared Suspect.
	Alert string `json:"alert"`
	// Angle of elevation/depression between observer and missile in degrees.
	AngElev float64 `json:"angElev"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	AouRptData []float64 `json:"aouRptData"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	AouRptType string `json:"aouRptType"`
	// Missile azimuth corridor data.
	AzCorr float64 `json:"azCorr"`
	// Indicates whether or not the missile is currently in a state of boosting.
	Boosting bool `json:"boosting"`
	// Track point burnout altitude relative to WGS-84 ellipsoid, in kilometers.
	BurnoutAlt float64 `json:"burnoutAlt"`
	// The call sign currently assigned to the track object.
	CallSign string `json:"callSign"`
	// The percentage of time that the estimated AoU will "cover" the true position of
	// the track.
	Containment float64 `json:"containment"`
	// An optional string array containing additional data (keys) representing relevant
	// items for context of fields not specifically defined in this schema. This array
	// is paired with the contextValues string array and must contain the same number
	// of items. Please note these fields are intended for contextual use only and do
	// not pertain to core schema information. To ensure proper integration and avoid
	// misuse, coordination of how these fields are populated and consumed is required
	// during onboarding.
	ContextKeys []string `json:"contextKeys"`
	// An optional string array containing the values associated with the contextKeys
	// array. This array is paired with the contextKeys string array and must contain
	// the same number of items. Please note these fields are intended for contextual
	// use only and do not pertain to core schema information. To ensure proper
	// integration and avoid misuse, coordination of how these fields are populated and
	// consumed is required during onboarding.
	ContextValues []string `json:"contextValues"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// The drop-point indicator setting.
	DropPtInd bool `json:"dropPtInd"`
	// Indicates whether or not a track has an emergency.
	EmgInd bool `json:"emgInd"`
	// The track environment type (AIR, LAND, SPACE, SUBSURFACE, SURFACE, UNKNOWN):
	//
	// AIR: Between sea level and the Kármán line, which has an altitude of 100
	// kilometers (62 miles).
	//
	// LAND: On the surface of dry land.
	//
	// SPACE: Above the Kármán line, which has an altitude of 100 kilometers (62
	// miles).
	//
	// SURFACE: On the surface of a body of water.
	//
	// SUBSURFACE: Below the surface of a body of water.
	//
	// UNKNOWN: Environment is not known.
	//
	// Any of "AIR", "LAND", "SPACE", "SURFACE", "SUBSURFACE", "UNKNOWN".
	Env MissileTrackTupleResponseEnv `json:"env"`
	// Estimated impact point altitude relative to WGS-84 ellipsoid, in kilometers.
	ImpactAlt float64 `json:"impactAlt"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	ImpactAouData []float64 `json:"impactAouData"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	ImpactAouType string `json:"impactAouType"`
	// Confidence level of the impact point estimate. 0 - 100 percent.
	ImpactConf float64 `json:"impactConf"`
	// WGS-84 latitude of the missile object impact point, in degrees. -90 to 90
	// degrees (negative values south of equator).
	ImpactLat float64 `json:"impactLat"`
	// WGS-84 longitude of the missile object impact point, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	ImpactLon float64 `json:"impactLon"`
	// Estimated time of impact timestamp in ISO8601 UTC format with microsecond
	// precision.
	ImpactTime time.Time `json:"impactTime" format:"date-time"`
	// Source code for source of information used to detect track.
	InfoSource string `json:"infoSource"`
	// Estimated launch point altitude relative to WGS-84 ellipsoid, in kilometers.
	LaunchAlt float64 `json:"launchAlt"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	LaunchAouData []float64 `json:"launchAouData"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	LaunchAouType string `json:"launchAouType"`
	// Angle between true north and the object's current position, with respect to the
	// launch point, in degrees. 0 to 360 degrees.
	LaunchAz float64 `json:"launchAz"`
	// Uncertainty of the launch azimuth, in degrees.
	LaunchAzUnc float64 `json:"launchAzUnc"`
	// Confidence level in the accuracy of the launch point estimate. 0 - 100 percent.
	LaunchConf float64 `json:"launchConf"`
	// WGS-84 latitude of the missile launch point, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	LaunchLat float64 `json:"launchLat"`
	// WGS-84 longitude of the missile launch point, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	LaunchLon float64 `json:"launchLon"`
	// Missile launch timestamp in ISO8601 UTC format with microsecond precision.
	LaunchTime time.Time `json:"launchTime" format:"date-time"`
	// Indicates whether or not the missile is considered lost.
	LostTrkInd bool `json:"lostTrkInd"`
	// Maneuver end time, in ISO 8601 UTC format with microsecond precision.
	ManeuverEnd time.Time `json:"maneuverEnd" format:"date-time"`
	// Maneuver start time, in ISO 8601 UTC format with microsecond precision.
	ManeuverStart time.Time `json:"maneuverStart" format:"date-time"`
	// The timestamp of the external message from which this request originated, if
	// applicable, in ISO8601 UTC format with millisecond precision.
	MsgCreateDate time.Time `json:"msgCreateDate" format:"date-time"`
	// The message subtype is a finer grain categorization of message types as many
	// messages can contain a variety of data content within the same structure.
	// Examples include but are not limited to Initial, Final, Launch, Update, etc.
	// Users should consult the appropriate documentation, based on the message type,
	// for the definitions of the subtypes that apply to that message.
	MsgSubType string `json:"msgSubType"`
	// The type of external message from which this request originated.
	MsgType string `json:"msgType"`
	// Missile status enumeration examples include but are not limited to:
	//
	// &nbsp;AT LAUNCH
	//
	// &nbsp;AT OBSERVATION
	//
	// &nbsp;FLYING
	//
	// &nbsp;IMPACTED
	//
	// &nbsp;LOST
	//
	// &nbsp;STALE
	//
	// &nbsp;DEBRIS.
	MslStatus string `json:"mslStatus"`
	// Source of the missile-unique identifier (MUID).
	MuidSrc string `json:"muidSrc"`
	// Track ID for the source of the missile-unique identifier.
	MuidSrcTrk string `json:"muidSrcTrk"`
	// Track name.
	Name string `json:"name"`
	// Space activity (examples: RECONNAISSANCE, ANTISPACE WARFARE, TELEVISION). The
	// activity in which the track object is engaged. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent activity designations. The activity can
	// be reported as either a combination of the code and environment (e.g. 65/AIR) or
	// as the descriptive enumeration (e.g. DIVERTING), which are equivalent.
	ObjAct string `json:"objAct"`
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
	ObjIdent MissileTrackTupleResponseObjIdent `json:"objIdent"`
	// Space Platform field along with the Space Activity field further defines the
	// identity of a Space track (examples: SATELLITE, WEAPON, PATROL). The object
	// platform type. Intended as, but not constrained to, MIL-STD-6016 environment
	// dependent platform type designations. The platform type can be reported as
	// either a combination of the code and environment (e.g. 14/LAND) or as the
	// descriptive representations (e.g. COMBAT VEHICLE), which are equivalent.
	ObjPlat string `json:"objPlat"`
	// The type of object to which this record refers. The object type may be updated
	// in later records based on assessment of additional data.
	ObjType string `json:"objType"`
	// Confidence of the object type, 0-100.
	ObjTypeConf int64 `json:"objTypeConf"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Track ID of the parent track, within the originating system, from which the
	// track was developed.
	ParentTrackID string `json:"parentTrackId"`
	// Azimuth corridor reference point latitude.
	PolarSingLocLat float64 `json:"polarSingLocLat"`
	// Azimuth corridor reference point longitude.
	PolarSingLocLon float64 `json:"polarSingLocLon"`
	// Last report type received from the sensor (for example, OBSBO = observation
	// burnout).
	SenMode string `json:"senMode"`
	// Space amplification indicates additional information on the space environment
	// being reported (examples: NUCLEAR WARHEAD, FUEL-AIR EXPLOSIVE WARHEAD, DEBRIS).
	SpaceAmp string `json:"spaceAmp"`
	// Confidence level of the amplifying characteristics. Values range from 0 to 6.
	SpaceAmpConf int64 `json:"spaceAmpConf"`
	// Specific type of point or track with an environment of space.
	SpaceSpecType string `json:"spaceSpecType"`
	// Track ID within the originating system.
	TrackID string `json:"trackId"`
	// Overall track confidence estimate (not standardized, but typically a value
	// between 0 and 1, with 0 indicating lowest confidence).
	TrkConf float64 `json:"trkConf"`
	// Track Quality is reported as an integer from 0-15. Track Quality specifies the
	// reliability of the positional information of a reported track, with higher
	// values indicating higher track quality; i.e., lower errors in reported position.
	TrkQual int64 `json:"trkQual"`
	// Array of MissileTrackVector objects. Missile track vectors are cartesian vectors
	// of position, velocity, and acceleration that, together with their time, 'epoch',
	// uniquely determine the trajectory of the missile. ECEF is the preferred
	// coordinate frame but in some cases data may be in another frame as specified by
	// 'referenceFrame', depending on the provider.
	Vectors []MissileTrackTupleResponseVector `json:"vectors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		AcftSubType           respjson.Field
		Alert                 respjson.Field
		AngElev               respjson.Field
		AouRptData            respjson.Field
		AouRptType            respjson.Field
		AzCorr                respjson.Field
		Boosting              respjson.Field
		BurnoutAlt            respjson.Field
		CallSign              respjson.Field
		Containment           respjson.Field
		ContextKeys           respjson.Field
		ContextValues         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DropPtInd             respjson.Field
		EmgInd                respjson.Field
		Env                   respjson.Field
		ImpactAlt             respjson.Field
		ImpactAouData         respjson.Field
		ImpactAouType         respjson.Field
		ImpactConf            respjson.Field
		ImpactLat             respjson.Field
		ImpactLon             respjson.Field
		ImpactTime            respjson.Field
		InfoSource            respjson.Field
		LaunchAlt             respjson.Field
		LaunchAouData         respjson.Field
		LaunchAouType         respjson.Field
		LaunchAz              respjson.Field
		LaunchAzUnc           respjson.Field
		LaunchConf            respjson.Field
		LaunchLat             respjson.Field
		LaunchLon             respjson.Field
		LaunchTime            respjson.Field
		LostTrkInd            respjson.Field
		ManeuverEnd           respjson.Field
		ManeuverStart         respjson.Field
		MsgCreateDate         respjson.Field
		MsgSubType            respjson.Field
		MsgType               respjson.Field
		MslStatus             respjson.Field
		MuidSrc               respjson.Field
		MuidSrcTrk            respjson.Field
		Name                  respjson.Field
		ObjAct                respjson.Field
		ObjIdent              respjson.Field
		ObjPlat               respjson.Field
		ObjType               respjson.Field
		ObjTypeConf           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ParentTrackID         respjson.Field
		PolarSingLocLat       respjson.Field
		PolarSingLocLon       respjson.Field
		SenMode               respjson.Field
		SpaceAmp              respjson.Field
		SpaceAmpConf          respjson.Field
		SpaceSpecType         respjson.Field
		TrackID               respjson.Field
		TrkConf               respjson.Field
		TrkQual               respjson.Field
		Vectors               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MissileTrackTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *MissileTrackTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type MissileTrackTupleResponseDataMode string

const (
	MissileTrackTupleResponseDataModeReal      MissileTrackTupleResponseDataMode = "REAL"
	MissileTrackTupleResponseDataModeTest      MissileTrackTupleResponseDataMode = "TEST"
	MissileTrackTupleResponseDataModeSimulated MissileTrackTupleResponseDataMode = "SIMULATED"
	MissileTrackTupleResponseDataModeExercise  MissileTrackTupleResponseDataMode = "EXERCISE"
)

// The track environment type (AIR, LAND, SPACE, SUBSURFACE, SURFACE, UNKNOWN):
//
// AIR: Between sea level and the Kármán line, which has an altitude of 100
// kilometers (62 miles).
//
// LAND: On the surface of dry land.
//
// SPACE: Above the Kármán line, which has an altitude of 100 kilometers (62
// miles).
//
// SURFACE: On the surface of a body of water.
//
// SUBSURFACE: Below the surface of a body of water.
//
// UNKNOWN: Environment is not known.
type MissileTrackTupleResponseEnv string

const (
	MissileTrackTupleResponseEnvAir        MissileTrackTupleResponseEnv = "AIR"
	MissileTrackTupleResponseEnvLand       MissileTrackTupleResponseEnv = "LAND"
	MissileTrackTupleResponseEnvSpace      MissileTrackTupleResponseEnv = "SPACE"
	MissileTrackTupleResponseEnvSurface    MissileTrackTupleResponseEnv = "SURFACE"
	MissileTrackTupleResponseEnvSubsurface MissileTrackTupleResponseEnv = "SUBSURFACE"
	MissileTrackTupleResponseEnvUnknown    MissileTrackTupleResponseEnv = "UNKNOWN"
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
type MissileTrackTupleResponseObjIdent string

const (
	MissileTrackTupleResponseObjIdentAssumedFriend MissileTrackTupleResponseObjIdent = "ASSUMED FRIEND"
	MissileTrackTupleResponseObjIdentFriend        MissileTrackTupleResponseObjIdent = "FRIEND"
	MissileTrackTupleResponseObjIdentHostile       MissileTrackTupleResponseObjIdent = "HOSTILE"
	MissileTrackTupleResponseObjIdentNeutral       MissileTrackTupleResponseObjIdent = "NEUTRAL"
	MissileTrackTupleResponseObjIdentPending       MissileTrackTupleResponseObjIdent = "PENDING"
	MissileTrackTupleResponseObjIdentSuspect       MissileTrackTupleResponseObjIdent = "SUSPECT"
	MissileTrackTupleResponseObjIdentUnknown       MissileTrackTupleResponseObjIdent = "UNKNOWN"
)

// Schema for Missile Track Vector data.
type MissileTrackTupleResponseVector struct {
	// Vector timestamp in ISO8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Three element array, expressing the cartesian acceleration vector of the target
	// object, in kilometers/second^2, in the specified referenceFrame. If
	// referenceFrame is null then ECEF should be assumed. The array element order is
	// [x”, y”, z”].
	Accel []float64 `json:"accel"`
	// Confidence of the vector, 0-100.
	Confidence int64 `json:"confidence"`
	// An optional string array containing additional data (keys) representing relevant
	// items for context of fields not specifically defined in this schema. This array
	// is paired with the contextValues string array and must contain the same number
	// of items. Please note these fields are intended for contextual use only and do
	// not pertain to core schema information. To ensure proper integration and avoid
	// misuse, coordination of how these fields are populated and consumed is required
	// during onboarding.
	ContextKeys []string `json:"contextKeys"`
	// An optional string array containing the values associated with the contextKeys
	// array. This array is paired with the contextKeys string array and must contain
	// the same number of items. Please note these fields are intended for contextual
	// use only and do not pertain to core schema information. To ensure proper
	// integration and avoid misuse, coordination of how these fields are populated and
	// consumed is required during onboarding.
	ContextValues []string `json:"contextValues"`
	// Track object course, in degrees clockwise from true north.
	Course float64 `json:"course"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame.
	//
	// If the covReferenceFrame is null it is assumed to be UVW. The array values
	// (1-45) represent the upper triangular half of the position-velocity-acceleration
	// covariance matrix.
	//
	// The covariance elements are position dependent within the array with values
	// ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x"&nbsp;&nbsp;&nbsp;&nbsp;y"&nbsp;&nbsp;&nbsp;&nbsp;z"
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;10&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;&nbsp;12&nbsp;&nbsp;&nbsp;13&nbsp;&nbsp;&nbsp;14&nbsp;&nbsp;&nbsp;15&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;&nbsp;17
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;18&nbsp;&nbsp;&nbsp;19&nbsp;&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;&nbsp;23&nbsp;&nbsp;&nbsp;24
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;25&nbsp;&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;&nbsp;30
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;31&nbsp;&nbsp;&nbsp;32&nbsp;&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;35
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;&nbsp;38&nbsp;&nbsp;&nbsp;39
	//
	// x"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;40&nbsp;&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42
	//
	// y"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;44
	//
	// z"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45
	//
	// The cov array should contain only the upper right triangle values from top left
	// down to bottom right, in order.
	Cov []float64 `json:"cov"`
	// The reference frame of the covariance elements (J2000, UVW, EFG/TDR, ECR/ECEF,
	// TEME, GCRF). If the referenceFrame is null it is assumed to be UVW.
	//
	// Any of "J2000", "UVW", "EFG/TDR", "ECR/ECEF", "TEME", "GCRF".
	CovReferenceFrame string `json:"covReferenceFrame"`
	// The flight azimuth associated with the current state vector (0-360 degrees).
	FlightAz float64 `json:"flightAz"`
	// Unique identifier of the reporting sensor of the object.
	IDSensor string `json:"idSensor"`
	// Object to which this vector applies.
	Object string `json:"object"`
	// Optional identifier provided by the source to indicate the reporting sensor of
	// the object. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Three element array, expressing the cartesian position vector of the target
	// object, in kilometers, in the specified referenceFrame. If referenceFrame is
	// null then ECEF should be assumed. The array element order is [x, y, z].
	Pos []float64 `json:"pos"`
	// Flag indicating whether the vector data was propagated.
	Propagated bool `json:"propagated"`
	// The quaternion describing the attitude of the spacecraft with respect to the
	// reference frame listed in the 'referenceFrame' field. The array element order
	// convention is the three vector components, followed by the scalar component.
	Quat []float64 `json:"quat"`
	// Range from the originating system or sensor to the object, in kilometers.
	Range float64 `json:"range"`
	// The reference frame of the cartesian vector (ECEF, J2000). If the referenceFrame
	// is null it is assumed to be ECEF.
	ReferenceFrame string `json:"referenceFrame"`
	// Track object speed, in kilometers/sec.
	Spd float64 `json:"spd"`
	// Status of the vector (e.g. INITIAL, UPDATE).
	Status string `json:"status"`
	// Source of the epoch time.
	TimeSource string `json:"timeSource"`
	// Type of vector represented (e.g. LOS, PREDICTED, STATE).
	Type string `json:"type"`
	// Object altitude at epoch, expressed in kilometers above WGS-84 ellipsoid.
	VectorAlt float64 `json:"vectorAlt"`
	// WGS-84 object latitude subpoint at epoch, represented as -90 to 90 degrees
	// (negative values south of equator).
	VectorLat float64 `json:"vectorLat"`
	// WGS-84 object longitude subpoint at epoch, represented as -180 to 180 degrees
	// (negative values west of Prime Meridian).
	VectorLon float64 `json:"vectorLon"`
	// Vector track ID within the originating system or sensor.
	VectorTrackID string `json:"vectorTrackId"`
	// Three element array, expressing the cartesian velocity vector of the target
	// object, in kilometers/second, in the specified referenceFrame. If referenceFrame
	// is null then ECEF should be assumed. The array element order is [x', y', z'].
	Vel []float64 `json:"vel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Epoch             respjson.Field
		Accel             respjson.Field
		Confidence        respjson.Field
		ContextKeys       respjson.Field
		ContextValues     respjson.Field
		Course            respjson.Field
		Cov               respjson.Field
		CovReferenceFrame respjson.Field
		FlightAz          respjson.Field
		IDSensor          respjson.Field
		Object            respjson.Field
		OrigSensorID      respjson.Field
		Pos               respjson.Field
		Propagated        respjson.Field
		Quat              respjson.Field
		Range             respjson.Field
		ReferenceFrame    respjson.Field
		Spd               respjson.Field
		Status            respjson.Field
		TimeSource        respjson.Field
		Type              respjson.Field
		VectorAlt         respjson.Field
		VectorLat         respjson.Field
		VectorLon         respjson.Field
		VectorTrackID     respjson.Field
		Vel               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MissileTrackTupleResponseVector) RawJSON() string { return r.JSON.raw }
func (r *MissileTrackTupleResponseVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MissileTrackListParams struct {
	// The receipt time of the data by the processing system, in ISO8601 UTC format
	// with microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MissileTrackListParams]'s query parameters as `url.Values`.
func (r MissileTrackListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MissileTrackCountParams struct {
	// The receipt time of the data by the processing system, in ISO8601 UTC format
	// with microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MissileTrackCountParams]'s query parameters as
// `url.Values`.
func (r MissileTrackCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MissileTrackNewBulkParams struct {
	Body []MissileTrackNewBulkParamsBody
	paramObj
}

func (r MissileTrackNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *MissileTrackNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide operations for querying of all available missile track
// details and amplifying missile data. A missile track is a position and
// optionally a heading/velocity of an object across all environments at a
// particular timestamp. It also includes optional information regarding the
// identity/type of missile, impact location, launch location and other amplifying
// object data, if known.
//
// The properties ClassificationMarking, DataMode, Source, Ts are required.
type MissileTrackNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The receipt time of the data by the processing system, in ISO8601 UTC format
	// with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Subtype is a finer grain categorization of missile types. Examples include but
	// are not limited to SRBM, MRBM, IRBM, LRBM, ICBM, SLBM.
	//
	// &nbsp;SRBM - Short-Range Ballistic Missile
	//
	// &nbsp;MRBM - Medium-Range Ballistic Missile
	//
	// &nbsp;IRBM - Intermediate-Range Ballistic Missile
	//
	// &nbsp;LRBM - Long-Range Ballistic Missile
	//
	// &nbsp;ICBM - Intercontinental Ballistic Missile
	//
	// &nbsp;SLBM - Submarine-Launched Ballistic Missile.
	AcftSubType param.Opt[string] `json:"acftSubType,omitzero"`
	// A track may be designated as a non-alert track or an alert track.
	//
	// Examples include but are not limited to:
	//
	// &nbsp;Non-alert tracks – choose None (Blank).
	//
	// &nbsp;Alert tracks – enter the proper alert classification:
	//
	// &nbsp;HIT - High Interest Track
	//
	// &nbsp;TGT - Target
	//
	// &nbsp;SUS - Suspect Carrier
	//
	// &nbsp;NSP - Cleared Suspect.
	Alert param.Opt[string] `json:"alert,omitzero"`
	// Angle of elevation/depression between observer and missile in degrees.
	AngElev param.Opt[float64] `json:"angElev,omitzero"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	AouRptType param.Opt[string] `json:"aouRptType,omitzero"`
	// Missile azimuth corridor data.
	AzCorr param.Opt[float64] `json:"azCorr,omitzero"`
	// Indicates whether or not the missile is currently in a state of boosting.
	Boosting param.Opt[bool] `json:"boosting,omitzero"`
	// Track point burnout altitude relative to WGS-84 ellipsoid, in kilometers.
	BurnoutAlt param.Opt[float64] `json:"burnoutAlt,omitzero"`
	// The call sign currently assigned to the track object.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// The percentage of time that the estimated AoU will "cover" the true position of
	// the track.
	Containment param.Opt[float64] `json:"containment,omitzero"`
	// The drop-point indicator setting.
	DropPtInd param.Opt[bool] `json:"dropPtInd,omitzero"`
	// Indicates whether or not a track has an emergency.
	EmgInd param.Opt[bool] `json:"emgInd,omitzero"`
	// Estimated impact point altitude relative to WGS-84 ellipsoid, in kilometers.
	ImpactAlt param.Opt[float64] `json:"impactAlt,omitzero"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	ImpactAouType param.Opt[string] `json:"impactAouType,omitzero"`
	// Confidence level of the impact point estimate. 0 - 100 percent.
	ImpactConf param.Opt[float64] `json:"impactConf,omitzero"`
	// WGS-84 latitude of the missile object impact point, in degrees. -90 to 90
	// degrees (negative values south of equator).
	ImpactLat param.Opt[float64] `json:"impactLat,omitzero"`
	// WGS-84 longitude of the missile object impact point, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	ImpactLon param.Opt[float64] `json:"impactLon,omitzero"`
	// Estimated time of impact timestamp in ISO8601 UTC format with microsecond
	// precision.
	ImpactTime param.Opt[time.Time] `json:"impactTime,omitzero" format:"date-time"`
	// Source code for source of information used to detect track.
	InfoSource param.Opt[string] `json:"infoSource,omitzero"`
	// Estimated launch point altitude relative to WGS-84 ellipsoid, in kilometers.
	LaunchAlt param.Opt[float64] `json:"launchAlt,omitzero"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	LaunchAouType param.Opt[string] `json:"launchAouType,omitzero"`
	// Angle between true north and the object's current position, with respect to the
	// launch point, in degrees. 0 to 360 degrees.
	LaunchAz param.Opt[float64] `json:"launchAz,omitzero"`
	// Uncertainty of the launch azimuth, in degrees.
	LaunchAzUnc param.Opt[float64] `json:"launchAzUnc,omitzero"`
	// Confidence level in the accuracy of the launch point estimate. 0 - 100 percent.
	LaunchConf param.Opt[float64] `json:"launchConf,omitzero"`
	// WGS-84 latitude of the missile launch point, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	LaunchLat param.Opt[float64] `json:"launchLat,omitzero"`
	// WGS-84 longitude of the missile launch point, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	LaunchLon param.Opt[float64] `json:"launchLon,omitzero"`
	// Missile launch timestamp in ISO8601 UTC format with microsecond precision.
	LaunchTime param.Opt[time.Time] `json:"launchTime,omitzero" format:"date-time"`
	// Indicates whether or not the missile is considered lost.
	LostTrkInd param.Opt[bool] `json:"lostTrkInd,omitzero"`
	// Maneuver end time, in ISO 8601 UTC format with microsecond precision.
	ManeuverEnd param.Opt[time.Time] `json:"maneuverEnd,omitzero" format:"date-time"`
	// Maneuver start time, in ISO 8601 UTC format with microsecond precision.
	ManeuverStart param.Opt[time.Time] `json:"maneuverStart,omitzero" format:"date-time"`
	// The timestamp of the external message from which this request originated, if
	// applicable, in ISO8601 UTC format with millisecond precision.
	MsgCreateDate param.Opt[time.Time] `json:"msgCreateDate,omitzero" format:"date-time"`
	// The message subtype is a finer grain categorization of message types as many
	// messages can contain a variety of data content within the same structure.
	// Examples include but are not limited to Initial, Final, Launch, Update, etc.
	// Users should consult the appropriate documentation, based on the message type,
	// for the definitions of the subtypes that apply to that message.
	MsgSubType param.Opt[string] `json:"msgSubType,omitzero"`
	// The type of external message from which this request originated.
	MsgType param.Opt[string] `json:"msgType,omitzero"`
	// Missile status enumeration examples include but are not limited to:
	//
	// &nbsp;AT LAUNCH
	//
	// &nbsp;AT OBSERVATION
	//
	// &nbsp;FLYING
	//
	// &nbsp;IMPACTED
	//
	// &nbsp;LOST
	//
	// &nbsp;STALE
	//
	// &nbsp;DEBRIS.
	MslStatus param.Opt[string] `json:"mslStatus,omitzero"`
	// Source of the missile-unique identifier (MUID).
	MuidSrc param.Opt[string] `json:"muidSrc,omitzero"`
	// Track ID for the source of the missile-unique identifier.
	MuidSrcTrk param.Opt[string] `json:"muidSrcTrk,omitzero"`
	// Track name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Space activity (examples: RECONNAISSANCE, ANTISPACE WARFARE, TELEVISION). The
	// activity in which the track object is engaged. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent activity designations. The activity can
	// be reported as either a combination of the code and environment (e.g. 65/AIR) or
	// as the descriptive enumeration (e.g. DIVERTING), which are equivalent.
	ObjAct param.Opt[string] `json:"objAct,omitzero"`
	// Space Platform field along with the Space Activity field further defines the
	// identity of a Space track (examples: SATELLITE, WEAPON, PATROL). The object
	// platform type. Intended as, but not constrained to, MIL-STD-6016 environment
	// dependent platform type designations. The platform type can be reported as
	// either a combination of the code and environment (e.g. 14/LAND) or as the
	// descriptive representations (e.g. COMBAT VEHICLE), which are equivalent.
	ObjPlat param.Opt[string] `json:"objPlat,omitzero"`
	// The type of object to which this record refers. The object type may be updated
	// in later records based on assessment of additional data.
	ObjType param.Opt[string] `json:"objType,omitzero"`
	// Confidence of the object type, 0-100.
	ObjTypeConf param.Opt[int64] `json:"objTypeConf,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Track ID of the parent track, within the originating system, from which the
	// track was developed.
	ParentTrackID param.Opt[string] `json:"parentTrackId,omitzero"`
	// Azimuth corridor reference point latitude.
	PolarSingLocLat param.Opt[float64] `json:"polarSingLocLat,omitzero"`
	// Azimuth corridor reference point longitude.
	PolarSingLocLon param.Opt[float64] `json:"polarSingLocLon,omitzero"`
	// Last report type received from the sensor (for example, OBSBO = observation
	// burnout).
	SenMode param.Opt[string] `json:"senMode,omitzero"`
	// Space amplification indicates additional information on the space environment
	// being reported (examples: NUCLEAR WARHEAD, FUEL-AIR EXPLOSIVE WARHEAD, DEBRIS).
	SpaceAmp param.Opt[string] `json:"spaceAmp,omitzero"`
	// Confidence level of the amplifying characteristics. Values range from 0 to 6.
	SpaceAmpConf param.Opt[int64] `json:"spaceAmpConf,omitzero"`
	// Specific type of point or track with an environment of space.
	SpaceSpecType param.Opt[string] `json:"spaceSpecType,omitzero"`
	// Track ID within the originating system.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Overall track confidence estimate (not standardized, but typically a value
	// between 0 and 1, with 0 indicating lowest confidence).
	TrkConf param.Opt[float64] `json:"trkConf,omitzero"`
	// Track Quality is reported as an integer from 0-15. Track Quality specifies the
	// reliability of the positional information of a reported track, with higher
	// values indicating higher track quality; i.e., lower errors in reported position.
	TrkQual param.Opt[int64] `json:"trkQual,omitzero"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	AouRptData []float64 `json:"aouRptData,omitzero"`
	// An optional string array containing additional data (keys) representing relevant
	// items for context of fields not specifically defined in this schema. This array
	// is paired with the contextValues string array and must contain the same number
	// of items. Please note these fields are intended for contextual use only and do
	// not pertain to core schema information. To ensure proper integration and avoid
	// misuse, coordination of how these fields are populated and consumed is required
	// during onboarding.
	ContextKeys []string `json:"contextKeys,omitzero"`
	// An optional string array containing the values associated with the contextKeys
	// array. This array is paired with the contextKeys string array and must contain
	// the same number of items. Please note these fields are intended for contextual
	// use only and do not pertain to core schema information. To ensure proper
	// integration and avoid misuse, coordination of how these fields are populated and
	// consumed is required during onboarding.
	ContextValues []string `json:"contextValues,omitzero"`
	// The track environment type (AIR, LAND, SPACE, SUBSURFACE, SURFACE, UNKNOWN):
	//
	// AIR: Between sea level and the Kármán line, which has an altitude of 100
	// kilometers (62 miles).
	//
	// LAND: On the surface of dry land.
	//
	// SPACE: Above the Kármán line, which has an altitude of 100 kilometers (62
	// miles).
	//
	// SURFACE: On the surface of a body of water.
	//
	// SUBSURFACE: Below the surface of a body of water.
	//
	// UNKNOWN: Environment is not known.
	//
	// Any of "AIR", "LAND", "SPACE", "SURFACE", "SUBSURFACE", "UNKNOWN".
	Env string `json:"env,omitzero"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	ImpactAouData []float64 `json:"impactAouData,omitzero"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	LaunchAouData []float64 `json:"launchAouData,omitzero"`
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
	// Array of MissileTrackVector objects. Missile track vectors are cartesian vectors
	// of position, velocity, and acceleration that, together with their time, 'epoch',
	// uniquely determine the trajectory of the missile. ECEF is the preferred
	// coordinate frame but in some cases data may be in another frame as specified by
	// 'referenceFrame', depending on the provider.
	Vectors []MissileTrackNewBulkParamsBodyVector `json:"vectors,omitzero"`
	paramObj
}

func (r MissileTrackNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow MissileTrackNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MissileTrackNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MissileTrackNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[MissileTrackNewBulkParamsBody](
		"env", "AIR", "LAND", "SPACE", "SURFACE", "SUBSURFACE", "UNKNOWN",
	)
	apijson.RegisterFieldValidator[MissileTrackNewBulkParamsBody](
		"objIdent", "ASSUMED FRIEND", "FRIEND", "HOSTILE", "NEUTRAL", "PENDING", "SUSPECT", "UNKNOWN",
	)
}

// Schema for Missile Track Vector data.
//
// The property Epoch is required.
type MissileTrackNewBulkParamsBodyVector struct {
	// Vector timestamp in ISO8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Confidence of the vector, 0-100.
	Confidence param.Opt[int64] `json:"confidence,omitzero"`
	// Track object course, in degrees clockwise from true north.
	Course param.Opt[float64] `json:"course,omitzero"`
	// The flight azimuth associated with the current state vector (0-360 degrees).
	FlightAz param.Opt[float64] `json:"flightAz,omitzero"`
	// Unique identifier of the reporting sensor of the object.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Object to which this vector applies.
	Object param.Opt[string] `json:"object,omitzero"`
	// Optional identifier provided by the source to indicate the reporting sensor of
	// the object. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Flag indicating whether the vector data was propagated.
	Propagated param.Opt[bool] `json:"propagated,omitzero"`
	// Range from the originating system or sensor to the object, in kilometers.
	Range param.Opt[float64] `json:"range,omitzero"`
	// The reference frame of the cartesian vector (ECEF, J2000). If the referenceFrame
	// is null it is assumed to be ECEF.
	ReferenceFrame param.Opt[string] `json:"referenceFrame,omitzero"`
	// Track object speed, in kilometers/sec.
	Spd param.Opt[float64] `json:"spd,omitzero"`
	// Status of the vector (e.g. INITIAL, UPDATE).
	Status param.Opt[string] `json:"status,omitzero"`
	// Source of the epoch time.
	TimeSource param.Opt[string] `json:"timeSource,omitzero"`
	// Type of vector represented (e.g. LOS, PREDICTED, STATE).
	Type param.Opt[string] `json:"type,omitzero"`
	// Object altitude at epoch, expressed in kilometers above WGS-84 ellipsoid.
	VectorAlt param.Opt[float64] `json:"vectorAlt,omitzero"`
	// WGS-84 object latitude subpoint at epoch, represented as -90 to 90 degrees
	// (negative values south of equator).
	VectorLat param.Opt[float64] `json:"vectorLat,omitzero"`
	// WGS-84 object longitude subpoint at epoch, represented as -180 to 180 degrees
	// (negative values west of Prime Meridian).
	VectorLon param.Opt[float64] `json:"vectorLon,omitzero"`
	// Vector track ID within the originating system or sensor.
	VectorTrackID param.Opt[string] `json:"vectorTrackId,omitzero"`
	// Three element array, expressing the cartesian acceleration vector of the target
	// object, in kilometers/second^2, in the specified referenceFrame. If
	// referenceFrame is null then ECEF should be assumed. The array element order is
	// [x”, y”, z”].
	Accel []float64 `json:"accel,omitzero"`
	// An optional string array containing additional data (keys) representing relevant
	// items for context of fields not specifically defined in this schema. This array
	// is paired with the contextValues string array and must contain the same number
	// of items. Please note these fields are intended for contextual use only and do
	// not pertain to core schema information. To ensure proper integration and avoid
	// misuse, coordination of how these fields are populated and consumed is required
	// during onboarding.
	ContextKeys []string `json:"contextKeys,omitzero"`
	// An optional string array containing the values associated with the contextKeys
	// array. This array is paired with the contextKeys string array and must contain
	// the same number of items. Please note these fields are intended for contextual
	// use only and do not pertain to core schema information. To ensure proper
	// integration and avoid misuse, coordination of how these fields are populated and
	// consumed is required during onboarding.
	ContextValues []string `json:"contextValues,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame.
	//
	// If the covReferenceFrame is null it is assumed to be UVW. The array values
	// (1-45) represent the upper triangular half of the position-velocity-acceleration
	// covariance matrix.
	//
	// The covariance elements are position dependent within the array with values
	// ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x"&nbsp;&nbsp;&nbsp;&nbsp;y"&nbsp;&nbsp;&nbsp;&nbsp;z"
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;10&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;&nbsp;12&nbsp;&nbsp;&nbsp;13&nbsp;&nbsp;&nbsp;14&nbsp;&nbsp;&nbsp;15&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;&nbsp;17
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;18&nbsp;&nbsp;&nbsp;19&nbsp;&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;&nbsp;23&nbsp;&nbsp;&nbsp;24
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;25&nbsp;&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;&nbsp;30
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;31&nbsp;&nbsp;&nbsp;32&nbsp;&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;35
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;&nbsp;38&nbsp;&nbsp;&nbsp;39
	//
	// x"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;40&nbsp;&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42
	//
	// y"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;44
	//
	// z"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45
	//
	// The cov array should contain only the upper right triangle values from top left
	// down to bottom right, in order.
	Cov []float64 `json:"cov,omitzero"`
	// The reference frame of the covariance elements (J2000, UVW, EFG/TDR, ECR/ECEF,
	// TEME, GCRF). If the referenceFrame is null it is assumed to be UVW.
	//
	// Any of "J2000", "UVW", "EFG/TDR", "ECR/ECEF", "TEME", "GCRF".
	CovReferenceFrame string `json:"covReferenceFrame,omitzero"`
	// Three element array, expressing the cartesian position vector of the target
	// object, in kilometers, in the specified referenceFrame. If referenceFrame is
	// null then ECEF should be assumed. The array element order is [x, y, z].
	Pos []float64 `json:"pos,omitzero"`
	// The quaternion describing the attitude of the spacecraft with respect to the
	// reference frame listed in the 'referenceFrame' field. The array element order
	// convention is the three vector components, followed by the scalar component.
	Quat []float64 `json:"quat,omitzero"`
	// Three element array, expressing the cartesian velocity vector of the target
	// object, in kilometers/second, in the specified referenceFrame. If referenceFrame
	// is null then ECEF should be assumed. The array element order is [x', y', z'].
	Vel []float64 `json:"vel,omitzero"`
	paramObj
}

func (r MissileTrackNewBulkParamsBodyVector) MarshalJSON() (data []byte, err error) {
	type shadow MissileTrackNewBulkParamsBodyVector
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MissileTrackNewBulkParamsBodyVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MissileTrackNewBulkParamsBodyVector](
		"covReferenceFrame", "J2000", "UVW", "EFG/TDR", "ECR/ECEF", "TEME", "GCRF",
	)
}

type MissileTrackTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The receipt time of the data by the processing system, in ISO8601 UTC format
	// with microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MissileTrackTupleParams]'s query parameters as
// `url.Values`.
func (r MissileTrackTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MissileTrackUnvalidatedPublishParams struct {
	Body []MissileTrackUnvalidatedPublishParamsBody
	paramObj
}

func (r MissileTrackUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *MissileTrackUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide operations for querying of all available missile track
// details and amplifying missile data. A missile track is a position and
// optionally a heading/velocity of an object across all environments at a
// particular timestamp. It also includes optional information regarding the
// identity/type of missile, impact location, launch location and other amplifying
// object data, if known.
//
// The properties ClassificationMarking, DataMode, Source, Ts are required.
type MissileTrackUnvalidatedPublishParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The receipt time of the data by the processing system, in ISO8601 UTC format
	// with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Subtype is a finer grain categorization of missile types. Examples include but
	// are not limited to SRBM, MRBM, IRBM, LRBM, ICBM, SLBM.
	//
	// &nbsp;SRBM - Short-Range Ballistic Missile
	//
	// &nbsp;MRBM - Medium-Range Ballistic Missile
	//
	// &nbsp;IRBM - Intermediate-Range Ballistic Missile
	//
	// &nbsp;LRBM - Long-Range Ballistic Missile
	//
	// &nbsp;ICBM - Intercontinental Ballistic Missile
	//
	// &nbsp;SLBM - Submarine-Launched Ballistic Missile.
	AcftSubType param.Opt[string] `json:"acftSubType,omitzero"`
	// A track may be designated as a non-alert track or an alert track.
	//
	// Examples include but are not limited to:
	//
	// &nbsp;Non-alert tracks – choose None (Blank).
	//
	// &nbsp;Alert tracks – enter the proper alert classification:
	//
	// &nbsp;HIT - High Interest Track
	//
	// &nbsp;TGT - Target
	//
	// &nbsp;SUS - Suspect Carrier
	//
	// &nbsp;NSP - Cleared Suspect.
	Alert param.Opt[string] `json:"alert,omitzero"`
	// Angle of elevation/depression between observer and missile in degrees.
	AngElev param.Opt[float64] `json:"angElev,omitzero"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	AouRptType param.Opt[string] `json:"aouRptType,omitzero"`
	// Missile azimuth corridor data.
	AzCorr param.Opt[float64] `json:"azCorr,omitzero"`
	// Indicates whether or not the missile is currently in a state of boosting.
	Boosting param.Opt[bool] `json:"boosting,omitzero"`
	// Track point burnout altitude relative to WGS-84 ellipsoid, in kilometers.
	BurnoutAlt param.Opt[float64] `json:"burnoutAlt,omitzero"`
	// The call sign currently assigned to the track object.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// The percentage of time that the estimated AoU will "cover" the true position of
	// the track.
	Containment param.Opt[float64] `json:"containment,omitzero"`
	// The drop-point indicator setting.
	DropPtInd param.Opt[bool] `json:"dropPtInd,omitzero"`
	// Indicates whether or not a track has an emergency.
	EmgInd param.Opt[bool] `json:"emgInd,omitzero"`
	// Estimated impact point altitude relative to WGS-84 ellipsoid, in kilometers.
	ImpactAlt param.Opt[float64] `json:"impactAlt,omitzero"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	ImpactAouType param.Opt[string] `json:"impactAouType,omitzero"`
	// Confidence level of the impact point estimate. 0 - 100 percent.
	ImpactConf param.Opt[float64] `json:"impactConf,omitzero"`
	// WGS-84 latitude of the missile object impact point, in degrees. -90 to 90
	// degrees (negative values south of equator).
	ImpactLat param.Opt[float64] `json:"impactLat,omitzero"`
	// WGS-84 longitude of the missile object impact point, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	ImpactLon param.Opt[float64] `json:"impactLon,omitzero"`
	// Estimated time of impact timestamp in ISO8601 UTC format with microsecond
	// precision.
	ImpactTime param.Opt[time.Time] `json:"impactTime,omitzero" format:"date-time"`
	// Source code for source of information used to detect track.
	InfoSource param.Opt[string] `json:"infoSource,omitzero"`
	// Estimated launch point altitude relative to WGS-84 ellipsoid, in kilometers.
	LaunchAlt param.Opt[float64] `json:"launchAlt,omitzero"`
	// The Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition. This
	// type defines the elements of the aouEllp array and is required if aouEllp is not
	// null. See the aouEllp field definition for specific information.
	LaunchAouType param.Opt[string] `json:"launchAouType,omitzero"`
	// Angle between true north and the object's current position, with respect to the
	// launch point, in degrees. 0 to 360 degrees.
	LaunchAz param.Opt[float64] `json:"launchAz,omitzero"`
	// Uncertainty of the launch azimuth, in degrees.
	LaunchAzUnc param.Opt[float64] `json:"launchAzUnc,omitzero"`
	// Confidence level in the accuracy of the launch point estimate. 0 - 100 percent.
	LaunchConf param.Opt[float64] `json:"launchConf,omitzero"`
	// WGS-84 latitude of the missile launch point, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	LaunchLat param.Opt[float64] `json:"launchLat,omitzero"`
	// WGS-84 longitude of the missile launch point, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	LaunchLon param.Opt[float64] `json:"launchLon,omitzero"`
	// Missile launch timestamp in ISO8601 UTC format with microsecond precision.
	LaunchTime param.Opt[time.Time] `json:"launchTime,omitzero" format:"date-time"`
	// Indicates whether or not the missile is considered lost.
	LostTrkInd param.Opt[bool] `json:"lostTrkInd,omitzero"`
	// Maneuver end time, in ISO 8601 UTC format with microsecond precision.
	ManeuverEnd param.Opt[time.Time] `json:"maneuverEnd,omitzero" format:"date-time"`
	// Maneuver start time, in ISO 8601 UTC format with microsecond precision.
	ManeuverStart param.Opt[time.Time] `json:"maneuverStart,omitzero" format:"date-time"`
	// The timestamp of the external message from which this request originated, if
	// applicable, in ISO8601 UTC format with millisecond precision.
	MsgCreateDate param.Opt[time.Time] `json:"msgCreateDate,omitzero" format:"date-time"`
	// The message subtype is a finer grain categorization of message types as many
	// messages can contain a variety of data content within the same structure.
	// Examples include but are not limited to Initial, Final, Launch, Update, etc.
	// Users should consult the appropriate documentation, based on the message type,
	// for the definitions of the subtypes that apply to that message.
	MsgSubType param.Opt[string] `json:"msgSubType,omitzero"`
	// The type of external message from which this request originated.
	MsgType param.Opt[string] `json:"msgType,omitzero"`
	// Missile status enumeration examples include but are not limited to:
	//
	// &nbsp;AT LAUNCH
	//
	// &nbsp;AT OBSERVATION
	//
	// &nbsp;FLYING
	//
	// &nbsp;IMPACTED
	//
	// &nbsp;LOST
	//
	// &nbsp;STALE
	//
	// &nbsp;DEBRIS.
	MslStatus param.Opt[string] `json:"mslStatus,omitzero"`
	// Source of the missile-unique identifier (MUID).
	MuidSrc param.Opt[string] `json:"muidSrc,omitzero"`
	// Track ID for the source of the missile-unique identifier.
	MuidSrcTrk param.Opt[string] `json:"muidSrcTrk,omitzero"`
	// Track name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Space activity (examples: RECONNAISSANCE, ANTISPACE WARFARE, TELEVISION). The
	// activity in which the track object is engaged. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent activity designations. The activity can
	// be reported as either a combination of the code and environment (e.g. 65/AIR) or
	// as the descriptive enumeration (e.g. DIVERTING), which are equivalent.
	ObjAct param.Opt[string] `json:"objAct,omitzero"`
	// Space Platform field along with the Space Activity field further defines the
	// identity of a Space track (examples: SATELLITE, WEAPON, PATROL). The object
	// platform type. Intended as, but not constrained to, MIL-STD-6016 environment
	// dependent platform type designations. The platform type can be reported as
	// either a combination of the code and environment (e.g. 14/LAND) or as the
	// descriptive representations (e.g. COMBAT VEHICLE), which are equivalent.
	ObjPlat param.Opt[string] `json:"objPlat,omitzero"`
	// The type of object to which this record refers. The object type may be updated
	// in later records based on assessment of additional data.
	ObjType param.Opt[string] `json:"objType,omitzero"`
	// Confidence of the object type, 0-100.
	ObjTypeConf param.Opt[int64] `json:"objTypeConf,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Track ID of the parent track, within the originating system, from which the
	// track was developed.
	ParentTrackID param.Opt[string] `json:"parentTrackId,omitzero"`
	// Azimuth corridor reference point latitude.
	PolarSingLocLat param.Opt[float64] `json:"polarSingLocLat,omitzero"`
	// Azimuth corridor reference point longitude.
	PolarSingLocLon param.Opt[float64] `json:"polarSingLocLon,omitzero"`
	// Last report type received from the sensor (for example, OBSBO = observation
	// burnout).
	SenMode param.Opt[string] `json:"senMode,omitzero"`
	// Space amplification indicates additional information on the space environment
	// being reported (examples: NUCLEAR WARHEAD, FUEL-AIR EXPLOSIVE WARHEAD, DEBRIS).
	SpaceAmp param.Opt[string] `json:"spaceAmp,omitzero"`
	// Confidence level of the amplifying characteristics. Values range from 0 to 6.
	SpaceAmpConf param.Opt[int64] `json:"spaceAmpConf,omitzero"`
	// Specific type of point or track with an environment of space.
	SpaceSpecType param.Opt[string] `json:"spaceSpecType,omitzero"`
	// Track ID within the originating system.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Overall track confidence estimate (not standardized, but typically a value
	// between 0 and 1, with 0 indicating lowest confidence).
	TrkConf param.Opt[float64] `json:"trkConf,omitzero"`
	// Track Quality is reported as an integer from 0-15. Track Quality specifies the
	// reliability of the positional information of a reported track, with higher
	// values indicating higher track quality; i.e., lower errors in reported position.
	TrkQual param.Opt[int64] `json:"trkQual,omitzero"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	AouRptData []float64 `json:"aouRptData,omitzero"`
	// An optional string array containing additional data (keys) representing relevant
	// items for context of fields not specifically defined in this schema. This array
	// is paired with the contextValues string array and must contain the same number
	// of items. Please note these fields are intended for contextual use only and do
	// not pertain to core schema information. To ensure proper integration and avoid
	// misuse, coordination of how these fields are populated and consumed is required
	// during onboarding.
	ContextKeys []string `json:"contextKeys,omitzero"`
	// An optional string array containing the values associated with the contextKeys
	// array. This array is paired with the contextKeys string array and must contain
	// the same number of items. Please note these fields are intended for contextual
	// use only and do not pertain to core schema information. To ensure proper
	// integration and avoid misuse, coordination of how these fields are populated and
	// consumed is required during onboarding.
	ContextValues []string `json:"contextValues,omitzero"`
	// The track environment type (AIR, LAND, SPACE, SUBSURFACE, SURFACE, UNKNOWN):
	//
	// AIR: Between sea level and the Kármán line, which has an altitude of 100
	// kilometers (62 miles).
	//
	// LAND: On the surface of dry land.
	//
	// SPACE: Above the Kármán line, which has an altitude of 100 kilometers (62
	// miles).
	//
	// SURFACE: On the surface of a body of water.
	//
	// SUBSURFACE: Below the surface of a body of water.
	//
	// UNKNOWN: Environment is not known.
	//
	// Any of "AIR", "LAND", "SPACE", "SURFACE", "SUBSURFACE", "UNKNOWN".
	Env string `json:"env,omitzero"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	ImpactAouData []float64 `json:"impactAouData,omitzero"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouType specified in
	// this record:
	//
	// &nbsp;ELLIPSE:
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the ellipse
	//
	// &nbsp;&nbsp;a1 - semi-major axis in kilometers
	//
	// &nbsp;&nbsp;a2 - semi-minor axis in kilometers
	//
	// &nbsp;BEARING (BEARING BOX or MTST BEARING BOX):
	//
	// &nbsp;&nbsp;brg - orientation in degrees of the bearing box
	//
	// &nbsp;&nbsp;a1 - length of bearing box in kilometers
	//
	// &nbsp;&nbsp;a2 - half-width of bearing box in kilometers
	//
	// &nbsp;OTHER (All other type values):
	//
	// &nbsp;&nbsp;brg - line of bearing in degrees true
	//
	// &nbsp;&nbsp;a1 - bearing error in degrees
	//
	// &nbsp;&nbsp;a2 - estimated range in kilometers.
	LaunchAouData []float64 `json:"launchAouData,omitzero"`
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
	// Array of MissileTrackVector objects. Missile track vectors are cartesian vectors
	// of position, velocity, and acceleration that, together with their time, 'epoch',
	// uniquely determine the trajectory of the missile. ECEF is the preferred
	// coordinate frame but in some cases data may be in another frame as specified by
	// 'referenceFrame', depending on the provider.
	Vectors []MissileTrackUnvalidatedPublishParamsBodyVector `json:"vectors,omitzero"`
	paramObj
}

func (r MissileTrackUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow MissileTrackUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MissileTrackUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MissileTrackUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[MissileTrackUnvalidatedPublishParamsBody](
		"env", "AIR", "LAND", "SPACE", "SURFACE", "SUBSURFACE", "UNKNOWN",
	)
	apijson.RegisterFieldValidator[MissileTrackUnvalidatedPublishParamsBody](
		"objIdent", "ASSUMED FRIEND", "FRIEND", "HOSTILE", "NEUTRAL", "PENDING", "SUSPECT", "UNKNOWN",
	)
}

// Schema for Missile Track Vector data.
//
// The property Epoch is required.
type MissileTrackUnvalidatedPublishParamsBodyVector struct {
	// Vector timestamp in ISO8601 UTC format, with microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Confidence of the vector, 0-100.
	Confidence param.Opt[int64] `json:"confidence,omitzero"`
	// Track object course, in degrees clockwise from true north.
	Course param.Opt[float64] `json:"course,omitzero"`
	// The flight azimuth associated with the current state vector (0-360 degrees).
	FlightAz param.Opt[float64] `json:"flightAz,omitzero"`
	// Unique identifier of the reporting sensor of the object.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Object to which this vector applies.
	Object param.Opt[string] `json:"object,omitzero"`
	// Optional identifier provided by the source to indicate the reporting sensor of
	// the object. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Flag indicating whether the vector data was propagated.
	Propagated param.Opt[bool] `json:"propagated,omitzero"`
	// Range from the originating system or sensor to the object, in kilometers.
	Range param.Opt[float64] `json:"range,omitzero"`
	// The reference frame of the cartesian vector (ECEF, J2000). If the referenceFrame
	// is null it is assumed to be ECEF.
	ReferenceFrame param.Opt[string] `json:"referenceFrame,omitzero"`
	// Track object speed, in kilometers/sec.
	Spd param.Opt[float64] `json:"spd,omitzero"`
	// Status of the vector (e.g. INITIAL, UPDATE).
	Status param.Opt[string] `json:"status,omitzero"`
	// Source of the epoch time.
	TimeSource param.Opt[string] `json:"timeSource,omitzero"`
	// Type of vector represented (e.g. LOS, PREDICTED, STATE).
	Type param.Opt[string] `json:"type,omitzero"`
	// Object altitude at epoch, expressed in kilometers above WGS-84 ellipsoid.
	VectorAlt param.Opt[float64] `json:"vectorAlt,omitzero"`
	// WGS-84 object latitude subpoint at epoch, represented as -90 to 90 degrees
	// (negative values south of equator).
	VectorLat param.Opt[float64] `json:"vectorLat,omitzero"`
	// WGS-84 object longitude subpoint at epoch, represented as -180 to 180 degrees
	// (negative values west of Prime Meridian).
	VectorLon param.Opt[float64] `json:"vectorLon,omitzero"`
	// Vector track ID within the originating system or sensor.
	VectorTrackID param.Opt[string] `json:"vectorTrackId,omitzero"`
	// Three element array, expressing the cartesian acceleration vector of the target
	// object, in kilometers/second^2, in the specified referenceFrame. If
	// referenceFrame is null then ECEF should be assumed. The array element order is
	// [x”, y”, z”].
	Accel []float64 `json:"accel,omitzero"`
	// An optional string array containing additional data (keys) representing relevant
	// items for context of fields not specifically defined in this schema. This array
	// is paired with the contextValues string array and must contain the same number
	// of items. Please note these fields are intended for contextual use only and do
	// not pertain to core schema information. To ensure proper integration and avoid
	// misuse, coordination of how these fields are populated and consumed is required
	// during onboarding.
	ContextKeys []string `json:"contextKeys,omitzero"`
	// An optional string array containing the values associated with the contextKeys
	// array. This array is paired with the contextKeys string array and must contain
	// the same number of items. Please note these fields are intended for contextual
	// use only and do not pertain to core schema information. To ensure proper
	// integration and avoid misuse, coordination of how these fields are populated and
	// consumed is required during onboarding.
	ContextValues []string `json:"contextValues,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame.
	//
	// If the covReferenceFrame is null it is assumed to be UVW. The array values
	// (1-45) represent the upper triangular half of the position-velocity-acceleration
	// covariance matrix.
	//
	// The covariance elements are position dependent within the array with values
	// ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x"&nbsp;&nbsp;&nbsp;&nbsp;y"&nbsp;&nbsp;&nbsp;&nbsp;z"
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;6&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;9
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;10&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;&nbsp;12&nbsp;&nbsp;&nbsp;13&nbsp;&nbsp;&nbsp;14&nbsp;&nbsp;&nbsp;15&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;&nbsp;17
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;18&nbsp;&nbsp;&nbsp;19&nbsp;&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;&nbsp;23&nbsp;&nbsp;&nbsp;24
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;25&nbsp;&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;&nbsp;30
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;31&nbsp;&nbsp;&nbsp;32&nbsp;&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;35
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;&nbsp;38&nbsp;&nbsp;&nbsp;39
	//
	// x"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;40&nbsp;&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42
	//
	// y"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;44
	//
	// z"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45
	//
	// The cov array should contain only the upper right triangle values from top left
	// down to bottom right, in order.
	Cov []float64 `json:"cov,omitzero"`
	// The reference frame of the covariance elements (J2000, UVW, EFG/TDR, ECR/ECEF,
	// TEME, GCRF). If the referenceFrame is null it is assumed to be UVW.
	//
	// Any of "J2000", "UVW", "EFG/TDR", "ECR/ECEF", "TEME", "GCRF".
	CovReferenceFrame string `json:"covReferenceFrame,omitzero"`
	// Three element array, expressing the cartesian position vector of the target
	// object, in kilometers, in the specified referenceFrame. If referenceFrame is
	// null then ECEF should be assumed. The array element order is [x, y, z].
	Pos []float64 `json:"pos,omitzero"`
	// The quaternion describing the attitude of the spacecraft with respect to the
	// reference frame listed in the 'referenceFrame' field. The array element order
	// convention is the three vector components, followed by the scalar component.
	Quat []float64 `json:"quat,omitzero"`
	// Three element array, expressing the cartesian velocity vector of the target
	// object, in kilometers/second, in the specified referenceFrame. If referenceFrame
	// is null then ECEF should be assumed. The array element order is [x', y', z'].
	Vel []float64 `json:"vel,omitzero"`
	paramObj
}

func (r MissileTrackUnvalidatedPublishParamsBodyVector) MarshalJSON() (data []byte, err error) {
	type shadow MissileTrackUnvalidatedPublishParamsBodyVector
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MissileTrackUnvalidatedPublishParamsBodyVector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MissileTrackUnvalidatedPublishParamsBodyVector](
		"covReferenceFrame", "J2000", "UVW", "EFG/TDR", "ECR/ECEF", "TEME", "GCRF",
	)
}
