// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// TrackDetailHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrackDetailHistoryService] method instead.
type TrackDetailHistoryService struct {
	Options []option.RequestOption
}

// NewTrackDetailHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTrackDetailHistoryService(opts ...option.RequestOption) (r TrackDetailHistoryService) {
	r = TrackDetailHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TrackDetailHistoryService) List(ctx context.Context, query TrackDetailHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[TrackDetailsFull], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/trackdetails/history"
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
func (r *TrackDetailHistoryService) ListAutoPaging(ctx context.Context, query TrackDetailHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[TrackDetailsFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TrackDetailHistoryService) Aodr(ctx context.Context, query TrackDetailHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/trackdetails/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *TrackDetailHistoryService) Count(ctx context.Context, query TrackDetailHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/trackdetails/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// These services provide operations for querying of all available track details
// and amplifying track data. A track is a position and optionally a
// heading/velocity of an object such as an aircraft, marine vessel, etc at a
// particular timestamp. It also includes optional information regarding the
// identity/type of the target object and other amplifying object data, if known.
type TrackDetailsFull struct {
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
	DataMode TrackDetailsFullDataMode `json:"dataMode,required"`
	// WGS-84 latitude of the track object, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	Lat float64 `json:"lat,required"`
	// WGS-84 longitude of the track object, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Track timestamp in ISO8601 UTC format with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Subtype is a finer grain categorization of missile types. Examples include but
	// are not limited to SRBM, MRBM, IRBM, LRBM, ICBM, SLBM:
	//
	// # SRBM - Short-Range Ballistic Missile
	//
	// # MRBM - Medium-Range Ballistic Missile
	//
	// # IRBM - Intermediate-Range Ballistic Missile
	//
	// # LRBM - Long-Range Ballistic Missile
	//
	// # ICBM - Intercontinental Ballistic Missile
	//
	// SLBM - Submarine-Launched Ballistic Missile.
	AcftSubType string `json:"acftSubType"`
	// A text amplifier for units, equipment and installations; content is
	// implementation specific.
	AddInfo string `json:"addInfo"`
	// A track may be designated as an alert track with the following designations:
	//
	// # HIT - High Interest Track
	//
	// # TGT - Target
	//
	// # SUS - Suspect Carrier
	//
	// # NSP - Cleared Suspect
	//
	// If alert is null, the track is assumed to be of non-alert status.
	Alert string `json:"alert"`
	// Track point altitude relative to WGS-84 ellipsoid, in meters. Positive values
	// indicate a track object height above ellipsoid, and negative values indicate a
	// track object below ellipsoid, applicable to the depth estimate for a subsurface
	// track.
	Alt float64 `json:"alt"`
	// The angle formed between the line of sight of the observer and the horizon, in
	// degrees. The angular range is -90 to 90, with negative values representing angle
	// of depression.
	AngElev float64 `json:"angElev"`
	// The reference dimensions of the vessel, reported as [A, B, C, D], in meters.
	// Where the array values represent the distance fore (A), aft (B), to port (C),
	// and to starboard (D) of the navigation antenna. Array with values A = C = 0 and
	// B, D > 0 indicate the length (B) and width (D) of the vessel without antenna
	// position reference.
	AntennaRefDimensions []float64 `json:"antennaRefDimensions"`
	// Three element array representing an Area of Uncertainty (AoU). The array element
	// definitions and units are type specific depending on the aouRptType specified in
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
	AouRptData []float64 `json:"aouRptData"`
	// The track Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER) definition.
	// This type defines the elements of the aouRptData array and is required if
	// aouRptData is not null. See the aouRptData field definition for specific
	// information.
	AouRptType string `json:"aouRptType"`
	// Appearance group code.
	//
	// GP1 - Superstructure exceeds 1/3 of ship length.
	//
	// GP2 - Superstructure less than 1/3 of ship length.
	//
	// GP3 - Stack aft.
	AppGrp string `json:"appGrp"`
	// The reported arrival cargo type. Intended as, but not constrained to, the USCG
	// NAVCEN AIS cargo definitions. Users should refer to USCG Navigation Center
	// documentation for specific definitions associated with ship and cargo types.
	// USCG NAVCEN documentation may be found at https://www.navcen.uscg.gov.
	ArrCargo string `json:"arrCargo"`
	// The flag of the arrival port.
	ArrFlag string `json:"arrFlag"`
	// The Arrival Port of the vessel according to the AIS transmission.
	ArrPort string `json:"arrPort"`
	// The Arrival Time of the vessel at the destination, in ISO 8601 UTC format with
	// microsecond precision.
	ArrTime time.Time `json:"arrTime" format:"date-time"`
	// Type of Aid to Navigation. Intended as, but not constrained to, the USCG NAVCEN
	// aids to navigation. Users should refer to USCG Navigation Center documentation
	// for specific device type information. USCG NAVCEN documentation may be found at
	// https://www.navcen.uscg.gov.
	Aton string `json:"aton"`
	// The average speed, in kilometers/hour, calculated for the subject during the
	// latest voyage/excursion.
	AvgSpd float64 `json:"avgSpd"`
	// Azimuth corridor arc distance measured in meters from reference point of azimuth
	// corridor to far edge of bounded azimuth corridor wedge, measured along azimuth
	// corridor center line.
	AzCorrArcWidth float64 `json:"azCorrArcWidth"`
	// The azimuth corridor centerline angle measured in degrees clockwise from true
	// north, of the center line of an azimuth corridor. The center line extends from
	// the referenced corridor origin location.
	AzCorrCenterLine float64 `json:"azCorrCenterLine"`
	// The Basic Encyclopedia (BE) number associated with this installation or area.
	BeNumber string `json:"beNumber"`
	// Flag indicating that the missile is currently in a state of boosting, if
	// reporting a missile track.
	Boosting bool `json:"boosting"`
	// Track point burnout altitude relative to WGS-84 ellipsoid, in meters.
	BurnoutAlt float64 `json:"burnoutAlt"`
	// The call sign currently assigned to this track object.
	CallSign string `json:"callSign"`
	// The reported cargo type. Intended as, but not constrained to, the USCG NAVCEN
	// AIS cargo definitions. Users should refer to USCG Navigation Center
	// documentation for specific definitions associated with ship and cargo types.
	// USCG NAVCEN documentation may be found at https://www.navcen.uscg.gov.
	CargoType string `json:"cargoType"`
	// Correlation Index; reference code for the site that originally reported the
	// track.
	CI string `json:"cI"`
	// The Area Of Uncertainty (AOU) percentage (0 - 100) containment value. The
	// percentage of time (90%) that the estimated area of uncertainty will cover the
	// true position of the track object.
	Containment float64 `json:"containment"`
	// The Cooperative Location Indicator specifies whether the reported entity
	// location was derived using reported locations from sensors on more than one
	// platform.
	//
	// 0 - COOPERATIVE_LOCATOR_NONE
	//
	// 1 - SINGLE_PLATFORM
	//
	// 2 - FRAGMENT
	//
	// 3 - COOPERATIVE.
	CoopLocInd string `json:"coopLocInd"`
	// The track object course-over-ground, in degrees clockwise from true North at the
	// object location (0-360 degrees).
	Course float64 `json:"course"`
	// The distance, in meters, of the closest point of approach between this track to
	// the master reference track.
	Cpa float64 `json:"cpa"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// The reported departure cargo type. Intended as, but not constrained to, the USCG
	// NAVCEN AIS cargo definitions. Users should refer to USCG Navigation Center
	// documentation for specific definitions associated with ship and cargo types.
	// USCG NAVCEN documentation may be found at https://www.navcen.uscg.gov.
	DepCargo string `json:"depCargo"`
	// The flag of the departure port.
	DepFlag string `json:"depFlag"`
	// The Departure Port of the vessel according to the AIS transmission.
	DepPort string `json:"depPort"`
	// The reported destination cargo type. Intended as, but not constrained to, the
	// USCG NAVCEN AIS cargo definitions. Users should refer to USCG Navigation Center
	// documentation for specific definitions associated with ship and cargo types.
	// USCG NAVCEN documentation may be found at https://www.navcen.uscg.gov.
	DesCargo string `json:"desCargo"`
	// The flag of the destination port.
	DesFlag string `json:"desFlag"`
	// The destination of the vessel according to the AIS transmission.
	Destination string `json:"destination"`
	// The Intel Discrete Identifier (IDI) code assigned to this track. The IDI is a
	// four-digit code representing hostile or unknown tracks.
	DisID string `json:"disId"`
	// The maximum static draught, in meters, of the vessel according to the AIS
	// transmission.
	Draught float64 `json:"draught"`
	// The drop-point indicator setting.
	DropPtInd bool `json:"dropPtInd"`
	// Flag indicating that this track represents a dummy object or group. Identifies
	// offensive or defensive units, equipment, and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Dummy bool `json:"dummy"`
	// Track object location in ECEF [x, y, z], meters. When provided, array must
	// always contain 3 values.
	EcefPos []float64 `json:"ecefPos"`
	// Track object velocity in ECEF [x', y', z'], meters/sec. When provided, array
	// must always contain 3 values.
	EcefVel []float64 `json:"ecefVel"`
	// Primary ELINT Notification (ELNOT), a five character identifier assigned to each
	// non-communication emission for collection and reporting purposes. This
	// five-digit field begins with an alpha character, followed by three numbers,
	// ending with another alpha character.
	Elnot1 string `json:"elnot1"`
	// Secondary ELINT Notification (ELNOT), a five character identifier assigned to
	// each non-communication emission for collection and reporting purposes. This
	// five-digit field begins with an alpha character, followed by three numbers,
	// ending with another alpha character.
	Elnot2 string `json:"elnot2"`
	// Flag indicating that the track object has an emergency.
	EmgInd bool `json:"emgInd"`
	// Radar name of the sensor tracking this object (e.g., RAY1500, SPN-43, HEADNET).
	EmitterName string `json:"emitterName"`
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
	Env TrackDetailsFullEnv `json:"env"`
	// The error in the Area Orientation. Area Orientation is the angle or roll in
	// degrees, between area semi-minor axis and plane defined by local vertical and
	// area semi-major axis. When reported without major and minor axis, it is angle
	// between an axis perpendicular to a 2D true north axis and plane defined by local
	// vertical and a 2D true north axis.
	ErrAreaOrient float64 `json:"errAreaOrient"`
	// Geometric area switch identifies the 3D shape of the error volume by providing
	// the 2D shape for the 3D plane. The 3D plane is the plane orthogonal to the area
	// semi-major axis and area semi-minor axis. Depending on whether the 2D plane is
	// reported as an error ellipse 2D or as an error rectangle 2D, this switch reports
	// a complete error volume that is cubical, 3D rectangular, cylindrical, 3D
	// ellipsoidal, or spherical.
	//
	// 0 - SWITCH_TYPE_UNKNOWN
	//
	// 1 - SQUARE_RECTANGLE
	//
	// 2 - CIRCLE_ELLIPSE.
	ErrGeoAreaSwitch string `json:"errGeoAreaSwitch"`
	// The error in the semi-intermediate axis. The semi-intermediate axis is
	// intermediate in length between semi-major and semi-minor axes. This field is
	// doubled and centered on intersection of area semi-major axis and area semi-minor
	// axis at 90 degrees to the plane defined by those axes. For these shapes, the
	// volume is defined as having a 50-percent probability of containing the true
	// location of the referenced entity.
	ErrSemiIntAxis float64 `json:"errSemiIntAxis"`
	// The error in the Semi-major elevation axis. Semi-major elevation axis is the
	// elevation of the cubical, 3D rectangular, cylindrical, 3D ellipsoidal, or
	// spherical semi-major axis, in degrees, measured from local horizontal.
	ErrSemiMajElev float64 `json:"errSemiMajElev"`
	// The Estimated Time of Arrival of the vessel at the destination port, in ISO 8601
	// UTC format with microsecond precision.
	Eta time.Time `json:"eta" format:"date-time"`
	// The Estimated Time of Departure of the vessel from the departure port (depPort),
	// according to Marine Traffic calculations, in ISO 8601 UTC format with
	// microsecond precision.
	Etd time.Time `json:"etd" format:"date-time"`
	// A text amplifier code for units, equipment, and installations that consists of a
	// one-letter reliability rating and a one-number credibility rating based on the
	// following definitions of each:
	//
	// Reliability Ratings:
	//
	// # A-completely reliable
	//
	// # B-usually reliable
	//
	// # C-fairly reliable
	//
	// # D-not usually reliable
	//
	// # E-unreliable
	//
	// # F-reliability cannot be judged
	//
	// Credibility Ratings:
	//
	// 1-confirmed by other sources
	//
	// 2-probably true
	//
	// 3-possibly true
	//
	// 4-doubtfully true
	//
	// 5-improbable
	//
	// 6-truth cannot be judged.
	EvalRating string `json:"evalRating"`
	// Flag indicating that this track represents a feint object or group. Identifies
	// offensive or defensive units, equipment, and/or installations intended to draw
	// the enemy's attention away from the area of the main attack. Based on
	// MIL-STD-2525 symbology definitions.
	Feint bool `json:"feint"`
	// Frequency, in hertz, for the signature report.
	Freq float64 `json:"freq"`
	// An ftn used to associate information and directives with the track.
	Ftn string `json:"ftn"`
	// The name of the Command reporting the Force Over-The-Horizon Track Coordinator
	// (FOTC) track number.
	FtnCmd string `json:"ftnCmd"`
	// The message timestamp that the ftn track position was recorded, in ISO 8601 UTC
	// format with microsecond precision.
	FtnMsgTs time.Time `json:"ftnMsgTs" format:"date-time"`
	// List of harmonics of the signature report in descending order of predominance
	// using 1-2 digit combinations separated by commas, e.g., 8,12,4. (1-22NS).
	Harmonics string `json:"harmonics"`
	// Track object heading, in degrees clockwise from true north.
	Hdng float64 `json:"hdng"`
	// Flag indicating that this track represents a headquarters object. Based on
	// MIL-STD-2525 symbology definitions.
	Hq bool `json:"hq"`
	// The vessel hull number designation of this maritime vessel. The hull number is a
	// 1-6 character alphanumeric entry assigned to a ship and painted on the hull.
	HullNum string `json:"hullNum"`
	// Hull profile code. Based on GCCS-J hull profiles.
	//
	// FLUSH No breaks in Hull Profile.
	//
	// RAISED 1 Hull Profile shows distinct raised area at bow. Remainder of deck is
	// flush.
	//
	// RAISED 2 Hull Profile shows distinct raised area amidships. Bow and stern are
	// flush.
	//
	// RAISED 3 Hull Profile shows distinct raised area at stern. Remainder of deck is
	// flush.
	//
	// RAISED 1-2-3 Distinct raised areas at bow, midships, and stern with breaks
	// between each raise.
	//
	// RAISED 1-2 Raised area at bow and midships with break between.
	//
	// RAISED 1-3 Raised area at bow and stern with break between.
	//
	// RAISED 12 Continuous raised area encompassing both bow and midships.
	//
	// RAISED 23 Continuous raised area encompassing midships and stern.
	//
	// RAISED 12-3 Raised areas at bow, midships and stern. Bow and midship raises are
	// continuous. Break between midship and stern raises.
	//
	// RAISED 1-23 Raised areas at bow, midships, and stern. Midship and stern raises
	// are continuous with break between bow and midship raises.
	//
	// RAISED 1-L2-3 Raised areas at bow, midships, and stern with break between each
	// raise. Midships raise is longer than that associated with raised 1-2-3.
	HullProf string `json:"hullProf"`
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
	// A text amplifier displaying IFF/SIF/AIS Identification modes and codes.
	Iff string `json:"iff"`
	// The International Maritime Organization Number of the vessel. IMON is a
	// seven-digit number that uniquely identifies the vessel.
	Imon int64 `json:"imon"`
	// Three element array representing the impact point Area of Uncertainty (AoU). The
	// array element definitions and units are type specific depending on the
	// impactAouType specified in this record:
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
	ImpactAouData []float64 `json:"impactAouData"`
	// The impact point Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER)
	// definition. This type defines the elements of the impactAouData array and is
	// required if impactAouData is not null. See the impactAouData field definition
	// for specific information.
	ImpactAouType string `json:"impactAouType"`
	// WGS-84 latitude of the missile impact point, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	ImpactLat float64 `json:"impactLat"`
	// WGS-84 longitude of the missile impact point, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	ImpactLon float64 `json:"impactLon"`
	// Missile impact timestamp in ISO8601 UTC format with microsecond precision.
	ImpactTime time.Time `json:"impactTime" format:"date-time"`
	// Source code for source of information used to detect track.
	InfoSource string `json:"infoSource"`
	// Flag indicating that this track represents an installation. Based on
	// MIL-STD-2525 symbology definitions.
	Installation bool `json:"installation"`
	// Three element array representing the launch location Area of Uncertainty (AoU).
	// The array element definitions and units are type specific depending on the
	// launchAouType specified in this record:
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
	LaunchAouData []float64 `json:"launchAouData"`
	// The launch location Area of Uncertainty (AoU) type (BEARING, ELLIPSE, OTHER)
	// definition. This type defines the elements of the launchAouData array and is
	// required if launchAouData is not null. See the launchAouData field definition
	// for specific information.
	LaunchAouType string `json:"launchAouType"`
	// WGS-84 latitude of the missile launch point, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	LaunchLat float64 `json:"launchLat"`
	// WGS-84 longitude of the missile launch point, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	LaunchLon float64 `json:"launchLon"`
	// Missile launch timestamp in ISO8601 UTC format with microsecond precision.
	LaunchTime time.Time `json:"launchTime" format:"date-time"`
	// The overall length of the vessel, in meters. A value of 511 indicates a vessel
	// length of 511 meters or greater.
	Length float64 `json:"length"`
	// Flag indicating whether the missile is considered lost, if reporting a missile
	// track.
	LostTrkInd bool `json:"lostTrkInd"`
	// The manuevering indicator specifying the missile boost phase.
	//
	// 0 - POST_BOOST_NONE
	//
	// 1 - POST_BOOST_MANUEVER
	//
	// 2 - POST_BOOST_COMPLETE_MANUEVER.
	ManeuverInd string `json:"maneuverInd"`
	// Maximum frequency, in hertz, reported for this acoustic track.
	MaxFreq float64 `json:"maxFreq"`
	// The category code that represents the associated facility purpose within the
	// target system. This value is the category code in the MIDB (Modernized
	// Intelligence Database).
	MidbCat string `json:"midbCat"`
	// The MIL-STD-2525B symbology code that applies to the subject of this track.
	Mil2525Bstr string `json:"mil2525Bstr"`
	// The Maritime Mobile Service Identity of the vessel. MMSI is a nine-digit number
	// that identifies the transmitter station of the vessel.
	Mmsi int64 `json:"mmsi"`
	// Optional message type designation.
	MsgType string `json:"msgType"`
	// The status of the missile track in this record, if reporting a missile track
	// (e.g. AT LAUNCH, AT OBSERVATION, FLYING, IMPACTED, LOST, STALE, DEBRIS).
	MslStatus string `json:"mslStatus"`
	// Source of the Missile-Unique Identifier (MUID), if reporting a missile track.
	MuidSrc string `json:"muidSrc"`
	// Track ID for the source of the Missile-Unique Identifier (MUID), if reporting a
	// missile track.
	MuidSrcTrk string `json:"muidSrcTrk"`
	// Track name.
	Name string `json:"name"`
	// The AIS Navigational Status of the vessel (e.g. Underway Using Engine, Moored,
	// Aground, etc.). Intended as, but not constrained to, the USCG NAVCEN navigation
	// status definitions. Users should refer to USCG Navigation Center documentation
	// for specific definitions associated with navigation status. USCG NAVCEN
	// documentation may be found at https://www.navcen.uscg.gov.
	NavStatus string `json:"navStatus"`
	// The Naval Tactical Data System (NTDS) track number assigned to this track.
	Ntds string `json:"ntds"`
	// The number of blades per shaft of the track object. Applicable for maritime
	// vessels.
	NumBlades int64 `json:"numBlades"`
	// The number of shafts on the track object. Applicable for maritime vessels.
	NumShafts int64 `json:"numShafts"`
	// The activity in which the track object is engaged. Intended as, but not
	// constrained to, MIL-STD-6016 environment dependent activity designations. The
	// activity can be reported as either a combination of the code and environment
	// (e.g. 65/AIR) or as the descriptive enumeration (e.g. DIVERTING), which are
	// equivalent. For cases in which no MIl-STD-6016 designation exists, a general
	// description can be used (e.g. ANTISPACE WARFARE).
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
	ObjIdent TrackDetailsFullObjIdent `json:"objIdent"`
	// Nationality of the tracked object.
	ObjNat string `json:"objNat"`
	// The object platform type is intended as, but not constrained to, MIL-STD-6016
	// environment dependent platform type designations. The platform type can be
	// reported as either a combination of the code and environment (e.g. 14/LAND) or
	// as the descriptive representations (e.g. COMBAT VEHICLE), which are equivalent.
	// For cases in which no MIl-STD-6016 designation exists, a general description can
	// be used (e.g. SATELLITE).
	ObjPlat string `json:"objPlat"`
	// The generic classification of the track object/group (e.g., BALLISTIC,
	// HELICOPTER, TRACKED, WATERCRAFT, WHEELED, etc.). Referenced, but not constrained
	// to, NATO STANAG 4676 object type classifications.
	ObjType string `json:"objType"`
	// Indicator position (OFF, ON, UNK) for optional floating navigational aids only.
	OffPosInd string `json:"offPosInd"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Original source cross-reference code for the Command that originated the report.
	OrigXref string `json:"origXref"`
	// The O-suffix associated with this facility. The O-suffix is a five-character
	// alpha/numeric system used to identify a facility, or demographic area, within an
	// installation. The Installation Basic Encyclopedia (beNumber), in conjunction
	// with the O-suffix, uniquely identifies the facility within the Modernized
	// Integrated Database (MIDB). The Installation beNumber and oSuffix are also used
	// in conjunction with the midbCat code to classify the function or purpose of the
	// facility.
	OSuffix string `json:"oSuffix"`
	// The Pseudo Identification Feature (PIF) number is a four digit code that
	// provides an exact ID for the ship or aircraft. Friendly military only.
	Pif string `json:"pif"`
	// This value represents the site number of a specific electronic site or its
	// associated equipment.
	Pin string `json:"pin"`
	// WGS-84 azimuth corridor reference point latitude, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	PolarSingLocLat float64 `json:"polarSingLocLat"`
	// WGS-84 azimuth corridor reference point longitude, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	PolarSingLocLon float64 `json:"polarSingLocLon"`
	// The type of electronic position fixing device (e.g. GPS, GLONASS, etc.).
	// Intended as, but not constrained to, the USCG NAVCEN electronic position fixing
	// device definitions. Users should refer to USCG Navigation Center documentation
	// for specific device type information. USCG NAVCEN documentation may be found at
	// https://www.navcen.uscg.gov.
	PosDeviceType string `json:"posDeviceType"`
	// Pulse repetition frequency of the emitter, measured in pulses per second. PRF is
	// the number of pulses transmitted per second. This is the reciprocal of the pri
	// value.
	Prf float64 `json:"prf"`
	// Pulse repetition interval of the emitter, measured in microseconds. The interval
	// between the start of one pulse and the start of another.
	Pri float64 `json:"pri"`
	// The number of propeller revolutions per minute for a submarine or ship.
	PropRpm float64 `json:"propRPM"`
	// The type of propulsion employed by the track object (Diesel, Nuclear).
	PropType string `json:"propType"`
	// Pulse width of the emitter, measured in microseconds. This is the duration of
	// the pulse.
	Pw float64 `json:"pw"`
	// Flag indicating for the track represents a reduced force object or group. Based
	// on MIL-STD-2525 symbology definitions.
	Reduced bool `json:"reduced"`
	// Flag indicating that this track represents a reinforced object or group. Based
	// on MIL-STD-2525 symbology definitions.
	Reinforced bool `json:"reinforced"`
	// Flag indicating whether this track is archived.
	RptArchived bool `json:"rptArchived"`
	// Source cross-reference code for the Command that originated the track report.
	RptChxref string `json:"rptChxref"`
	// A Reference Track Number used to associate information and directives with the
	// track. Referenced, but not constrained to, MIL-STD-6016F Reference Track Number.
	// The 'rtnMsgTs' and 'rtn' arrays must match in size.
	Rtn []string `json:"rtn"`
	// The name of the Command reporting the Received Track Number (RTN).
	RtnCmd string `json:"rtnCmd"`
	// The message timestamp that the reference track position was recorded, in ISO
	// 8601 UTC format with microsecond precision. The 'rtnMsgTs' and 'rtn' arrays must
	// match in size.
	RtnMsgTs []time.Time `json:"rtnMsgTs" format:"date-time"`
	// Value representing the state of the Received Track.
	RtnTrkState string `json:"rtnTrkState"`
	// Scan rate of the emitter, measured in seconds per rotation (SPR).
	ScanRate float64 `json:"scanRate"`
	// Type of radar scan.
	ScanType string `json:"scanType"`
	// The Sequential Contact Number (SCN) for this track.
	Scn int64 `json:"scn"`
	// The Ship Control Number (SCONUM) is a naval vessel identification number
	// (alphanumeric code) assigned by the Office of Naval Intelligence. SCONUM is
	// sometimes referred to as NOIC ID. SCONUMs are typically of the form A#####,
	// where A is an alpha character and # is numerical.
	Sconum string `json:"sconum"`
	// Flag indicating that this track is self reported.
	SelfReport bool `json:"selfReport"`
	// Id/name of sensor providing the track data.
	Sen string `json:"sen"`
	// The common name for a group of ships with similar design, usually named for the
	// first vessel of the class.
	ShipClass string `json:"shipClass"`
	// Abbreviated track name.
	ShortName string `json:"shortName"`
	// The unique identifier of the source node.
	SourceUid string `json:"sourceUid"`
	// Space amplification indicates additional information on the space environment
	// object being reported (e.g. DEBRIS, FUEL-AIR EXPLOSIVE, NUCLEAR WARHEAD).
	SpaceAmp string `json:"spaceAmp"`
	// Confidence level of the amplifying characteristics. Values range from 0 to 6,
	// with 0 indicating the lowest confidence and 6 indicating the highest.
	SpaceAmpConf int64 `json:"spaceAmpConf"`
	// Specific type of point or track with an environment of space.
	SpaceSpecType string `json:"spaceSpecType"`
	// Track object speed, in meters/sec.
	Spd float64 `json:"spd"`
	// A text amplifier for units, equipment and installations; content is
	// implementation specific.
	StaffCmts string `json:"staffCmts"`
	// The stern type code (Counter, Cruiser) associated with the track object.
	SternType string `json:"sternType"`
	// Flag indicating that this track represents a task force. Based on MIL-STD-2525
	// symbology definitions.
	TaskForce bool `json:"taskForce"`
	// The time, in ISO 8601 UTC format with millisecond precision, of the closest
	// point of approach between this track and the master reference track.
	Tcpa time.Time `json:"tcpa" format:"date-time"`
	// Threat Event System Track ID.
	TesEventID string `json:"tesEventId"`
	// Motion model Time On Leg in hours.
	Tol float64 `json:"tol"`
	// The number of turns of the vessel propellers per knot of forward motion.
	Tpk float64 `json:"tpk"`
	// Overall track confidence estimate (not standardized, but typically a value
	// between 0 and 1, with 0 indicating lowest confidence).
	TrkConf float64 `json:"trkConf"`
	// UUID identifying the track, which should remain the same on subsequent tracks of
	// the same object.
	TrkID string `json:"trkId"`
	// The track number (TN) of a surveillance entity. Intended as, but not constrained
	// to, the J-series track number encoded as five character alpha-numeric
	// characters. Users should refer to J-series documentation for specific TN
	// definitions.
	TrkNum string `json:"trkNum"`
	// Track Quality is reported as an integer from 0-15. Track Quality specifies the
	// reliability of the positional information of a reported track; Higher values
	// indicate higher track quality, i.e., lower errors in reported position.
	TrkQual int64 `json:"trkQual"`
	// Value Indicating the scope of this track: 1 - TERMINAL (Terminal) - available
	// only on the workstation where they were created. 2 - LOCAL (Local) - available
	// only on workstations in the local area network 3 - OTH (Over the Horizon) -
	// available to everyone.
	TrkScope string `json:"trkScope"`
	// Transponder ID for the track. This does not correspond to the UDL transponder
	// schema.
	TrnspdrID string `json:"trnspdrId"`
	// Transponder type for the track.
	TrnspdrType string `json:"trnspdrType"`
	// The weight, in tons, of the vessel associated with this track.
	VslWt float64 `json:"vslWt"`
	// The breadth of the vessel, in meters. A value of 63 indicates a vessel breadth
	// of 63 meters or greater.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		AcftSubType           respjson.Field
		AddInfo               respjson.Field
		Alert                 respjson.Field
		Alt                   respjson.Field
		AngElev               respjson.Field
		AntennaRefDimensions  respjson.Field
		AouRptData            respjson.Field
		AouRptType            respjson.Field
		AppGrp                respjson.Field
		ArrCargo              respjson.Field
		ArrFlag               respjson.Field
		ArrPort               respjson.Field
		ArrTime               respjson.Field
		Aton                  respjson.Field
		AvgSpd                respjson.Field
		AzCorrArcWidth        respjson.Field
		AzCorrCenterLine      respjson.Field
		BeNumber              respjson.Field
		Boosting              respjson.Field
		BurnoutAlt            respjson.Field
		CallSign              respjson.Field
		CargoType             respjson.Field
		CI                    respjson.Field
		Containment           respjson.Field
		CoopLocInd            respjson.Field
		Course                respjson.Field
		Cpa                   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DepCargo              respjson.Field
		DepFlag               respjson.Field
		DepPort               respjson.Field
		DesCargo              respjson.Field
		DesFlag               respjson.Field
		Destination           respjson.Field
		DisID                 respjson.Field
		Draught               respjson.Field
		DropPtInd             respjson.Field
		Dummy                 respjson.Field
		EcefPos               respjson.Field
		EcefVel               respjson.Field
		Elnot1                respjson.Field
		Elnot2                respjson.Field
		EmgInd                respjson.Field
		EmitterName           respjson.Field
		Env                   respjson.Field
		ErrAreaOrient         respjson.Field
		ErrGeoAreaSwitch      respjson.Field
		ErrSemiIntAxis        respjson.Field
		ErrSemiMajElev        respjson.Field
		Eta                   respjson.Field
		Etd                   respjson.Field
		EvalRating            respjson.Field
		Feint                 respjson.Field
		Freq                  respjson.Field
		Ftn                   respjson.Field
		FtnCmd                respjson.Field
		FtnMsgTs              respjson.Field
		Harmonics             respjson.Field
		Hdng                  respjson.Field
		Hq                    respjson.Field
		HullNum               respjson.Field
		HullProf              respjson.Field
		IdentAmp              respjson.Field
		Iff                   respjson.Field
		Imon                  respjson.Field
		ImpactAouData         respjson.Field
		ImpactAouType         respjson.Field
		ImpactLat             respjson.Field
		ImpactLon             respjson.Field
		ImpactTime            respjson.Field
		InfoSource            respjson.Field
		Installation          respjson.Field
		LaunchAouData         respjson.Field
		LaunchAouType         respjson.Field
		LaunchLat             respjson.Field
		LaunchLon             respjson.Field
		LaunchTime            respjson.Field
		Length                respjson.Field
		LostTrkInd            respjson.Field
		ManeuverInd           respjson.Field
		MaxFreq               respjson.Field
		MidbCat               respjson.Field
		Mil2525Bstr           respjson.Field
		Mmsi                  respjson.Field
		MsgType               respjson.Field
		MslStatus             respjson.Field
		MuidSrc               respjson.Field
		MuidSrcTrk            respjson.Field
		Name                  respjson.Field
		NavStatus             respjson.Field
		Ntds                  respjson.Field
		NumBlades             respjson.Field
		NumShafts             respjson.Field
		ObjAct                respjson.Field
		ObjIdent              respjson.Field
		ObjNat                respjson.Field
		ObjPlat               respjson.Field
		ObjType               respjson.Field
		OffPosInd             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigXref              respjson.Field
		OSuffix               respjson.Field
		Pif                   respjson.Field
		Pin                   respjson.Field
		PolarSingLocLat       respjson.Field
		PolarSingLocLon       respjson.Field
		PosDeviceType         respjson.Field
		Prf                   respjson.Field
		Pri                   respjson.Field
		PropRpm               respjson.Field
		PropType              respjson.Field
		Pw                    respjson.Field
		Reduced               respjson.Field
		Reinforced            respjson.Field
		RptArchived           respjson.Field
		RptChxref             respjson.Field
		Rtn                   respjson.Field
		RtnCmd                respjson.Field
		RtnMsgTs              respjson.Field
		RtnTrkState           respjson.Field
		ScanRate              respjson.Field
		ScanType              respjson.Field
		Scn                   respjson.Field
		Sconum                respjson.Field
		SelfReport            respjson.Field
		Sen                   respjson.Field
		ShipClass             respjson.Field
		ShortName             respjson.Field
		SourceUid             respjson.Field
		SpaceAmp              respjson.Field
		SpaceAmpConf          respjson.Field
		SpaceSpecType         respjson.Field
		Spd                   respjson.Field
		StaffCmts             respjson.Field
		SternType             respjson.Field
		TaskForce             respjson.Field
		Tcpa                  respjson.Field
		TesEventID            respjson.Field
		Tol                   respjson.Field
		Tpk                   respjson.Field
		TrkConf               respjson.Field
		TrkID                 respjson.Field
		TrkNum                respjson.Field
		TrkQual               respjson.Field
		TrkScope              respjson.Field
		TrnspdrID             respjson.Field
		TrnspdrType           respjson.Field
		VslWt                 respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackDetailsFull) RawJSON() string { return r.JSON.raw }
func (r *TrackDetailsFull) UnmarshalJSON(data []byte) error {
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
type TrackDetailsFullDataMode string

const (
	TrackDetailsFullDataModeReal      TrackDetailsFullDataMode = "REAL"
	TrackDetailsFullDataModeTest      TrackDetailsFullDataMode = "TEST"
	TrackDetailsFullDataModeSimulated TrackDetailsFullDataMode = "SIMULATED"
	TrackDetailsFullDataModeExercise  TrackDetailsFullDataMode = "EXERCISE"
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
type TrackDetailsFullEnv string

const (
	TrackDetailsFullEnvAir        TrackDetailsFullEnv = "AIR"
	TrackDetailsFullEnvLand       TrackDetailsFullEnv = "LAND"
	TrackDetailsFullEnvSpace      TrackDetailsFullEnv = "SPACE"
	TrackDetailsFullEnvSurface    TrackDetailsFullEnv = "SURFACE"
	TrackDetailsFullEnvSubsurface TrackDetailsFullEnv = "SUBSURFACE"
	TrackDetailsFullEnvUnknown    TrackDetailsFullEnv = "UNKNOWN"
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
type TrackDetailsFullObjIdent string

const (
	TrackDetailsFullObjIdentAssumedFriend TrackDetailsFullObjIdent = "ASSUMED FRIEND"
	TrackDetailsFullObjIdentFriend        TrackDetailsFullObjIdent = "FRIEND"
	TrackDetailsFullObjIdentHostile       TrackDetailsFullObjIdent = "HOSTILE"
	TrackDetailsFullObjIdentNeutral       TrackDetailsFullObjIdent = "NEUTRAL"
	TrackDetailsFullObjIdentPending       TrackDetailsFullObjIdent = "PENDING"
	TrackDetailsFullObjIdentSuspect       TrackDetailsFullObjIdent = "SUSPECT"
	TrackDetailsFullObjIdentUnknown       TrackDetailsFullObjIdent = "UNKNOWN"
)

type TrackDetailHistoryListParams struct {
	// Track timestamp in ISO8601 UTC format with microsecond precision.
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

// URLQuery serializes [TrackDetailHistoryListParams]'s query parameters as
// `url.Values`.
func (r TrackDetailHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TrackDetailHistoryAodrParams struct {
	// Track timestamp in ISO8601 UTC format with microsecond precision.
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

// URLQuery serializes [TrackDetailHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r TrackDetailHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TrackDetailHistoryCountParams struct {
	// Track timestamp in ISO8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TrackDetailHistoryCountParams]'s query parameters as
// `url.Values`.
func (r TrackDetailHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
