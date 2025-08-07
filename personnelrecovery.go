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

// PersonnelrecoveryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonnelrecoveryService] method instead.
type PersonnelrecoveryService struct {
	Options []option.RequestOption
	History PersonnelrecoveryHistoryService
}

// NewPersonnelrecoveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPersonnelrecoveryService(opts ...option.RequestOption) (r PersonnelrecoveryService) {
	r = PersonnelrecoveryService{}
	r.Options = opts
	r.History = NewPersonnelrecoveryHistoryService(opts...)
	return
}

// Service operation to take a single Personnel Recovery object as a POST body and
// ingest into the database. Requires a specific role, please contact the UDL team
// to gain access.
func (r *PersonnelrecoveryService) New(ctx context.Context, body PersonnelrecoveryNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/personnelrecovery"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *PersonnelrecoveryService) List(ctx context.Context, query PersonnelrecoveryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[PersonnelrecoveryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/personnelrecovery"
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
func (r *PersonnelrecoveryService) ListAutoPaging(ctx context.Context, query PersonnelrecoveryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[PersonnelrecoveryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *PersonnelrecoveryService) Count(ctx context.Context, query PersonnelrecoveryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/personnelrecovery/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// Personnel Recovery records as a POST body and ingest into the database. Requires
// specific roles, please contact the UDL team to gain access. This operation is
// not intended to be used for automated feeds into UDL...data providers should
// contact the UDL team for instructions on setting up a permanent feed through an
// alternate mechanism.
func (r *PersonnelrecoveryService) NewBulk(ctx context.Context, body PersonnelrecoveryNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/personnelrecovery/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to take a list of Personnel Recovery records as a POST body
// and ingest into the database. Requires a specific role, please contact the UDL
// team to gain access. This operation is intended to be used for automated feeds
// into UDL.
func (r *PersonnelrecoveryService) FileNew(ctx context.Context, body PersonnelrecoveryFileNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-personnelrecovery"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single PersonnelRecovery by its unique ID passed as a
// path parameter.
func (r *PersonnelrecoveryService) Get(ctx context.Context, id string, query PersonnelrecoveryGetParams, opts ...option.RequestOption) (res *PersonnelRecoveryFullL, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/personnelrecovery/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *PersonnelrecoveryService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *PersonnelrecoveryQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/personnelrecovery/queryhelp"
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
func (r *PersonnelrecoveryService) Tuple(ctx context.Context, query PersonnelrecoveryTupleParams, opts ...option.RequestOption) (res *[]PersonnelRecoveryFullL, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/personnelrecovery/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Provides information concerning search and rescue operations and other
// situations involving personnel recovery.
type PersonnelRecoveryFullL struct {
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
	DataMode PersonnelRecoveryFullLDataMode `json:"dataMode,required"`
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// WGS-84 latitude of the pickup location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PickupLat float64 `json:"pickupLat,required"`
	// WGS-84 longitude of the pickup location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PickupLon float64 `json:"pickupLon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Specifies the type of incident resulting in a recovery or evacuation mission.
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Emergency Type (e.g. NO
	// STATEMENT, DOWN AIRCRAFT, MAN IN WATER, DITCHING, BAILOUT, DISTRESSED VEHICLE,
	// GROUND INCIDENT, MEDICAL, ISOLATED PERSONS, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Mechanism used to verify the survivors identity.
	AuthMethod string `json:"authMethod"`
	// The confirmation status of the isolated personnel identity. Intended as, but not
	// constrained to, MIL-STD-6016 J6.1 Authentication Status, Isolated Personnel (NO
	// STATEMENT, AUTHENTICATED, NOT AUTHENTICATED, AUTHENTICATED UNDER DURESS, NOT
	// APPLICABLE):
	//
	// AUTHENTICATED: Confirmed Friend
	//
	// NOT AUTHENTICATED: Unconfirmed status
	//
	// AUTHENTICATED UNDER DURESS: Authentication comprised by hostiles.
	//
	// NOT APPLICABLE: Authentication not required.
	AuthStatus string `json:"authStatus"`
	// Flag indicating whether a radio identifier is reported.
	BeaconInd bool `json:"beaconInd"`
	// The call sign of the personnel to be recovered.
	CallSign string `json:"callSign"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq1 string `json:"commEq1"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq2 string `json:"commEq2"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq3 string `json:"commEq3"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy     string                              `json:"createdBy"`
	ExecutionInfo PersonnelRecoveryFullLExecutionInfo `json:"executionInfo"`
	// The survivor service identity (UNKNOWN MILITARY, UNKNOWN CIVILIAN, FRIEND
	// MILITARY, FRIEND CIVIILIAN, NEUTRAL MILITARY, NEUTRAL CIVILIAN, HOSTILE
	// MILITARY, HOSTILE CIVILIAN).
	Identity string `json:"identity"`
	// Unique identifier of a weather report associated with this recovery.
	IDWeatherReport string `json:"idWeatherReport"`
	// The military classification of the personnel to be recovered. Intended as, but
	// not constrained to, MIL-STD-6016 J6.1 Isolated Personnel Classification (NO
	// STATEMENT, MILITARY, GOVERNMENT CIVILIAN, GOVERNMENT CONTRACTOR, CIVILIAN,
	// MULTIPLE CLASSIFICATIONS).
	MilClass string `json:"milClass"`
	// The country of origin or political entity of an isolated person subject to
	// rescue or evacuation. If natAlliance is set to 126, then natAlliance1 must be
	// non 0. If natAlliance is any number other than 126, then natAlliance1 will be
	// set to 0 regardless. Defined in MIL-STD-6016 J6.1 Nationality/Alliance isolated
	// person(s).
	NatAlliance int64 `json:"natAlliance"`
	// Extended country of origin or political entity of an isolated person subject to
	// rescue or evacuation. Specify an entry here only if natAlliance is 126. Defined
	// in MIL-STD-6016 J6.1 Nationality/Alliance isolated person(s), 1.
	NatAlliance1 int64 `json:"natAlliance1"`
	// Number of ambulatory personnel requiring recovery.
	NumAmbulatory int64 `json:"numAmbulatory"`
	// Number of injured, but ambulatory, personnel requiring recovery.
	NumAmbulatoryInjured int64 `json:"numAmbulatoryInjured"`
	// Number of littered personnel requiring recovery.
	NumNonAmbulatory int64 `json:"numNonAmbulatory"`
	// The count of persons requiring recovery.
	NumPersons        int64                                   `json:"numPersons"`
	ObjectiveAreaInfo PersonnelRecoveryFullLObjectiveAreaInfo `json:"objectiveAreaInfo"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Altitude relative to WGS-84 ellipsoid, in meters. Positive values indicate a
	// point height above ellipsoid, and negative values indicate a point eight below
	// ellipsoid.
	PickupAlt float64 `json:"pickupAlt"`
	// UUID identifying the Personnel Recovery mission, which should remain the same on
	// subsequent posts related to the same recovery mission.
	RecovID string `json:"recovId"`
	// Receive voice frequency in 5Hz increments. This field will auto populate with
	// the txFreq value if the post element is null.
	RxFreq float64 `json:"rxFreq"`
	// Preloaded message conveying the situation confronting the isolated person(s).
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Survivor Radio Messages
	// (e.g. INJURED CANT MOVE NO KNOWN HOSTILES, INJURED CANT MOVE HOSTILES NEARBY,
	// UNINJURED CANT MOVE HOSTILES NEARBY, UNINJURED NO KNOWN HOSTILES, INJURED
	// LIMITED MOBILITY).
	SurvivorMessages string `json:"survivorMessages"`
	// Survivor radio equipment. Intended as, but not constrained to, MIL-STD-6016 J6.1
	// Survivor Radio Type (NO STATEMENT, PRQ7SEL, PRC90, PRC112, PRC112B B1, PRC112C,
	// PRC112D, PRC148 MBITR, PRC148 JEM, PRC149, PRC152, ACRPLB, OTHER).
	SurvivorRadio string `json:"survivorRadio"`
	// Flag indicating the cancellation of this recovery.
	TermInd bool `json:"termInd"`
	// Additional specific messages received from survivor.
	TextMsg string `json:"textMsg"`
	// Transmit voice frequency in 5Hz increments.
	TxFreq float64 `json:"txFreq"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		MsgTime               respjson.Field
		PickupLat             respjson.Field
		PickupLon             respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AuthMethod            respjson.Field
		AuthStatus            respjson.Field
		BeaconInd             respjson.Field
		CallSign              respjson.Field
		CommEq1               respjson.Field
		CommEq2               respjson.Field
		CommEq3               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ExecutionInfo         respjson.Field
		Identity              respjson.Field
		IDWeatherReport       respjson.Field
		MilClass              respjson.Field
		NatAlliance           respjson.Field
		NatAlliance1          respjson.Field
		NumAmbulatory         respjson.Field
		NumAmbulatoryInjured  respjson.Field
		NumNonAmbulatory      respjson.Field
		NumPersons            respjson.Field
		ObjectiveAreaInfo     respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PickupAlt             respjson.Field
		RecovID               respjson.Field
		RxFreq                respjson.Field
		SurvivorMessages      respjson.Field
		SurvivorRadio         respjson.Field
		TermInd               respjson.Field
		TextMsg               respjson.Field
		TxFreq                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelRecoveryFullL) RawJSON() string { return r.JSON.raw }
func (r *PersonnelRecoveryFullL) UnmarshalJSON(data []byte) error {
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
type PersonnelRecoveryFullLDataMode string

const (
	PersonnelRecoveryFullLDataModeReal      PersonnelRecoveryFullLDataMode = "REAL"
	PersonnelRecoveryFullLDataModeTest      PersonnelRecoveryFullLDataMode = "TEST"
	PersonnelRecoveryFullLDataModeSimulated PersonnelRecoveryFullLDataMode = "SIMULATED"
	PersonnelRecoveryFullLDataModeExercise  PersonnelRecoveryFullLDataMode = "EXERCISE"
)

type PersonnelRecoveryFullLExecutionInfo struct {
	// The heading, in degrees, of leaving the recovery zone.
	Egress float64 `json:"egress"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the egress location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	EgressPoint   []float64                                        `json:"egressPoint"`
	EscortVehicle PersonnelRecoveryFullLExecutionInfoEscortVehicle `json:"escortVehicle"`
	// The heading, in degrees clockwise from North, of entering the recovery zone.
	Ingress float64 `json:"ingress"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the initial location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	InitialPoint []float64 `json:"initialPoint"`
	// Description of the objective strategy plan.
	ObjStrategy     string                                             `json:"objStrategy"`
	RecoveryVehicle PersonnelRecoveryFullLExecutionInfoRecoveryVehicle `json:"recoveryVehicle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Egress          respjson.Field
		EgressPoint     respjson.Field
		EscortVehicle   respjson.Field
		Ingress         respjson.Field
		InitialPoint    respjson.Field
		ObjStrategy     respjson.Field
		RecoveryVehicle respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelRecoveryFullLExecutionInfo) RawJSON() string { return r.JSON.raw }
func (r *PersonnelRecoveryFullLExecutionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelRecoveryFullLExecutionInfoEscortVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign string `json:"callSign"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq float64 `json:"primaryFreq"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength int64 `json:"strength"`
	// The particular type of recovery vehicle to be used.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallSign    respjson.Field
		PrimaryFreq respjson.Field
		Strength    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelRecoveryFullLExecutionInfoEscortVehicle) RawJSON() string { return r.JSON.raw }
func (r *PersonnelRecoveryFullLExecutionInfoEscortVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelRecoveryFullLExecutionInfoRecoveryVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign string `json:"callSign"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq float64 `json:"primaryFreq"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength int64 `json:"strength"`
	// The particular type of recovery vehicle to be used.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallSign    respjson.Field
		PrimaryFreq respjson.Field
		Strength    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelRecoveryFullLExecutionInfoRecoveryVehicle) RawJSON() string { return r.JSON.raw }
func (r *PersonnelRecoveryFullLExecutionInfoRecoveryVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelRecoveryFullLObjectiveAreaInfo struct {
	// Information detailing knowledge of enemies in the area.
	EnemyData []PersonnelRecoveryFullLObjectiveAreaInfoEnemyData `json:"enemyData"`
	// The call sign of the on-scene commander.
	OscCallSign string `json:"oscCallSign"`
	// The radio frequency of the on-scene commander.
	OscFreq float64 `json:"oscFreq"`
	// Description of the pickup zone location.
	PzDesc string `json:"pzDesc"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the pz location. This array must contain a
	// minimum of 2 elements (latitude and longitude), and may contain an optional 3rd
	// element (altitude).
	PzLocation []float64 `json:"pzLocation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnemyData   respjson.Field
		OscCallSign respjson.Field
		OscFreq     respjson.Field
		PzDesc      respjson.Field
		PzLocation  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelRecoveryFullLObjectiveAreaInfo) RawJSON() string { return r.JSON.raw }
func (r *PersonnelRecoveryFullLObjectiveAreaInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelRecoveryFullLObjectiveAreaInfoEnemyData struct {
	// Directions to known enemies in the operation area (NORTH, NORTHEAST, EAST,
	// SOUTHEAST, SOUTH, SOUTHWEST, WEST, NORTHWEST, SURROUNDED).
	DirToEnemy string `json:"dirToEnemy"`
	// Comments provided by friendlies about the evac zone.
	FriendliesRemarks string `json:"friendliesRemarks"`
	// Hot Landing Zone remarks.
	HlzRemarks string `json:"hlzRemarks"`
	// The type of hostile fire received (SMALL ARMS, MORTAR, ARTILLERY, ROCKETS).
	HostileFireType string `json:"hostileFireType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DirToEnemy        respjson.Field
		FriendliesRemarks respjson.Field
		HlzRemarks        respjson.Field
		HostileFireType   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelRecoveryFullLObjectiveAreaInfoEnemyData) RawJSON() string { return r.JSON.raw }
func (r *PersonnelRecoveryFullLObjectiveAreaInfoEnemyData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides information concerning search and rescue operations and other
// situations involving personnel recovery.
type PersonnelrecoveryListResponse struct {
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
	DataMode PersonnelrecoveryListResponseDataMode `json:"dataMode,required"`
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// WGS-84 latitude of the pickup location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PickupLat float64 `json:"pickupLat,required"`
	// WGS-84 longitude of the pickup location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PickupLon float64 `json:"pickupLon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Specifies the type of incident resulting in a recovery or evacuation mission.
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Emergency Type (e.g. NO
	// STATEMENT, DOWN AIRCRAFT, MAN IN WATER, DITCHING, BAILOUT, DISTRESSED VEHICLE,
	// GROUND INCIDENT, MEDICAL, ISOLATED PERSONS, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Mechanism used to verify the survivors identity.
	AuthMethod string `json:"authMethod"`
	// The confirmation status of the isolated personnel identity. Intended as, but not
	// constrained to, MIL-STD-6016 J6.1 Authentication Status, Isolated Personnel (NO
	// STATEMENT, AUTHENTICATED, NOT AUTHENTICATED, AUTHENTICATED UNDER DURESS, NOT
	// APPLICABLE):
	//
	// AUTHENTICATED: Confirmed Friend
	//
	// NOT AUTHENTICATED: Unconfirmed status
	//
	// AUTHENTICATED UNDER DURESS: Authentication comprised by hostiles.
	//
	// NOT APPLICABLE: Authentication not required.
	AuthStatus string `json:"authStatus"`
	// Flag indicating whether a radio identifier is reported.
	BeaconInd bool `json:"beaconInd"`
	// The call sign of the personnel to be recovered.
	CallSign string `json:"callSign"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq1 string `json:"commEq1"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq2 string `json:"commEq2"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq3 string `json:"commEq3"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy     string                                     `json:"createdBy"`
	ExecutionInfo PersonnelrecoveryListResponseExecutionInfo `json:"executionInfo"`
	// The survivor service identity (UNKNOWN MILITARY, UNKNOWN CIVILIAN, FRIEND
	// MILITARY, FRIEND CIVIILIAN, NEUTRAL MILITARY, NEUTRAL CIVILIAN, HOSTILE
	// MILITARY, HOSTILE CIVILIAN).
	Identity string `json:"identity"`
	// Unique identifier of a weather report associated with this recovery.
	IDWeatherReport string `json:"idWeatherReport"`
	// The military classification of the personnel to be recovered. Intended as, but
	// not constrained to, MIL-STD-6016 J6.1 Isolated Personnel Classification (NO
	// STATEMENT, MILITARY, GOVERNMENT CIVILIAN, GOVERNMENT CONTRACTOR, CIVILIAN,
	// MULTIPLE CLASSIFICATIONS).
	MilClass string `json:"milClass"`
	// The country of origin or political entity of an isolated person subject to
	// rescue or evacuation. If natAlliance is set to 126, then natAlliance1 must be
	// non 0. If natAlliance is any number other than 126, then natAlliance1 will be
	// set to 0 regardless. Defined in MIL-STD-6016 J6.1 Nationality/Alliance isolated
	// person(s).
	NatAlliance int64 `json:"natAlliance"`
	// Extended country of origin or political entity of an isolated person subject to
	// rescue or evacuation. Specify an entry here only if natAlliance is 126. Defined
	// in MIL-STD-6016 J6.1 Nationality/Alliance isolated person(s), 1.
	NatAlliance1 int64 `json:"natAlliance1"`
	// Number of ambulatory personnel requiring recovery.
	NumAmbulatory int64 `json:"numAmbulatory"`
	// Number of injured, but ambulatory, personnel requiring recovery.
	NumAmbulatoryInjured int64 `json:"numAmbulatoryInjured"`
	// Number of littered personnel requiring recovery.
	NumNonAmbulatory int64 `json:"numNonAmbulatory"`
	// The count of persons requiring recovery.
	NumPersons        int64                                          `json:"numPersons"`
	ObjectiveAreaInfo PersonnelrecoveryListResponseObjectiveAreaInfo `json:"objectiveAreaInfo"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Altitude relative to WGS-84 ellipsoid, in meters. Positive values indicate a
	// point height above ellipsoid, and negative values indicate a point eight below
	// ellipsoid.
	PickupAlt float64 `json:"pickupAlt"`
	// UUID identifying the Personnel Recovery mission, which should remain the same on
	// subsequent posts related to the same recovery mission.
	RecovID string `json:"recovId"`
	// Receive voice frequency in 5Hz increments. This field will auto populate with
	// the txFreq value if the post element is null.
	RxFreq float64 `json:"rxFreq"`
	// Preloaded message conveying the situation confronting the isolated person(s).
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Survivor Radio Messages
	// (e.g. INJURED CANT MOVE NO KNOWN HOSTILES, INJURED CANT MOVE HOSTILES NEARBY,
	// UNINJURED CANT MOVE HOSTILES NEARBY, UNINJURED NO KNOWN HOSTILES, INJURED
	// LIMITED MOBILITY).
	SurvivorMessages string `json:"survivorMessages"`
	// Survivor radio equipment. Intended as, but not constrained to, MIL-STD-6016 J6.1
	// Survivor Radio Type (NO STATEMENT, PRQ7SEL, PRC90, PRC112, PRC112B B1, PRC112C,
	// PRC112D, PRC148 MBITR, PRC148 JEM, PRC149, PRC152, ACRPLB, OTHER).
	SurvivorRadio string `json:"survivorRadio"`
	// Flag indicating the cancellation of this recovery.
	TermInd bool `json:"termInd"`
	// Additional specific messages received from survivor.
	TextMsg string `json:"textMsg"`
	// Transmit voice frequency in 5Hz increments.
	TxFreq float64 `json:"txFreq"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		MsgTime               respjson.Field
		PickupLat             respjson.Field
		PickupLon             respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AuthMethod            respjson.Field
		AuthStatus            respjson.Field
		BeaconInd             respjson.Field
		CallSign              respjson.Field
		CommEq1               respjson.Field
		CommEq2               respjson.Field
		CommEq3               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ExecutionInfo         respjson.Field
		Identity              respjson.Field
		IDWeatherReport       respjson.Field
		MilClass              respjson.Field
		NatAlliance           respjson.Field
		NatAlliance1          respjson.Field
		NumAmbulatory         respjson.Field
		NumAmbulatoryInjured  respjson.Field
		NumNonAmbulatory      respjson.Field
		NumPersons            respjson.Field
		ObjectiveAreaInfo     respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PickupAlt             respjson.Field
		RecovID               respjson.Field
		RxFreq                respjson.Field
		SurvivorMessages      respjson.Field
		SurvivorRadio         respjson.Field
		TermInd               respjson.Field
		TextMsg               respjson.Field
		TxFreq                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelrecoveryListResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonnelrecoveryListResponse) UnmarshalJSON(data []byte) error {
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
type PersonnelrecoveryListResponseDataMode string

const (
	PersonnelrecoveryListResponseDataModeReal      PersonnelrecoveryListResponseDataMode = "REAL"
	PersonnelrecoveryListResponseDataModeTest      PersonnelrecoveryListResponseDataMode = "TEST"
	PersonnelrecoveryListResponseDataModeSimulated PersonnelrecoveryListResponseDataMode = "SIMULATED"
	PersonnelrecoveryListResponseDataModeExercise  PersonnelrecoveryListResponseDataMode = "EXERCISE"
)

type PersonnelrecoveryListResponseExecutionInfo struct {
	// The heading, in degrees, of leaving the recovery zone.
	Egress float64 `json:"egress"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the egress location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	EgressPoint   []float64                                               `json:"egressPoint"`
	EscortVehicle PersonnelrecoveryListResponseExecutionInfoEscortVehicle `json:"escortVehicle"`
	// The heading, in degrees clockwise from North, of entering the recovery zone.
	Ingress float64 `json:"ingress"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the initial location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	InitialPoint []float64 `json:"initialPoint"`
	// Description of the objective strategy plan.
	ObjStrategy     string                                                    `json:"objStrategy"`
	RecoveryVehicle PersonnelrecoveryListResponseExecutionInfoRecoveryVehicle `json:"recoveryVehicle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Egress          respjson.Field
		EgressPoint     respjson.Field
		EscortVehicle   respjson.Field
		Ingress         respjson.Field
		InitialPoint    respjson.Field
		ObjStrategy     respjson.Field
		RecoveryVehicle respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelrecoveryListResponseExecutionInfo) RawJSON() string { return r.JSON.raw }
func (r *PersonnelrecoveryListResponseExecutionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryListResponseExecutionInfoEscortVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign string `json:"callSign"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq float64 `json:"primaryFreq"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength int64 `json:"strength"`
	// The particular type of recovery vehicle to be used.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallSign    respjson.Field
		PrimaryFreq respjson.Field
		Strength    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelrecoveryListResponseExecutionInfoEscortVehicle) RawJSON() string { return r.JSON.raw }
func (r *PersonnelrecoveryListResponseExecutionInfoEscortVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryListResponseExecutionInfoRecoveryVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign string `json:"callSign"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq float64 `json:"primaryFreq"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength int64 `json:"strength"`
	// The particular type of recovery vehicle to be used.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallSign    respjson.Field
		PrimaryFreq respjson.Field
		Strength    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelrecoveryListResponseExecutionInfoRecoveryVehicle) RawJSON() string {
	return r.JSON.raw
}
func (r *PersonnelrecoveryListResponseExecutionInfoRecoveryVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryListResponseObjectiveAreaInfo struct {
	// Information detailing knowledge of enemies in the area.
	EnemyData []PersonnelrecoveryListResponseObjectiveAreaInfoEnemyData `json:"enemyData"`
	// The call sign of the on-scene commander.
	OscCallSign string `json:"oscCallSign"`
	// The radio frequency of the on-scene commander.
	OscFreq float64 `json:"oscFreq"`
	// Description of the pickup zone location.
	PzDesc string `json:"pzDesc"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the pz location. This array must contain a
	// minimum of 2 elements (latitude and longitude), and may contain an optional 3rd
	// element (altitude).
	PzLocation []float64 `json:"pzLocation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnemyData   respjson.Field
		OscCallSign respjson.Field
		OscFreq     respjson.Field
		PzDesc      respjson.Field
		PzLocation  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelrecoveryListResponseObjectiveAreaInfo) RawJSON() string { return r.JSON.raw }
func (r *PersonnelrecoveryListResponseObjectiveAreaInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryListResponseObjectiveAreaInfoEnemyData struct {
	// Directions to known enemies in the operation area (NORTH, NORTHEAST, EAST,
	// SOUTHEAST, SOUTH, SOUTHWEST, WEST, NORTHWEST, SURROUNDED).
	DirToEnemy string `json:"dirToEnemy"`
	// Comments provided by friendlies about the evac zone.
	FriendliesRemarks string `json:"friendliesRemarks"`
	// Hot Landing Zone remarks.
	HlzRemarks string `json:"hlzRemarks"`
	// The type of hostile fire received (SMALL ARMS, MORTAR, ARTILLERY, ROCKETS).
	HostileFireType string `json:"hostileFireType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DirToEnemy        respjson.Field
		FriendliesRemarks respjson.Field
		HlzRemarks        respjson.Field
		HostileFireType   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonnelrecoveryListResponseObjectiveAreaInfoEnemyData) RawJSON() string { return r.JSON.raw }
func (r *PersonnelrecoveryListResponseObjectiveAreaInfoEnemyData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryQueryhelpResponse struct {
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
func (r PersonnelrecoveryQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonnelrecoveryQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryNewParams struct {
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
	DataMode PersonnelrecoveryNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// WGS-84 latitude of the pickup location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PickupLat float64 `json:"pickupLat,required"`
	// WGS-84 longitude of the pickup location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PickupLon float64 `json:"pickupLon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Specifies the type of incident resulting in a recovery or evacuation mission.
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Emergency Type (e.g. NO
	// STATEMENT, DOWN AIRCRAFT, MAN IN WATER, DITCHING, BAILOUT, DISTRESSED VEHICLE,
	// GROUND INCIDENT, MEDICAL, ISOLATED PERSONS, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Mechanism used to verify the survivors identity.
	AuthMethod param.Opt[string] `json:"authMethod,omitzero"`
	// The confirmation status of the isolated personnel identity. Intended as, but not
	// constrained to, MIL-STD-6016 J6.1 Authentication Status, Isolated Personnel (NO
	// STATEMENT, AUTHENTICATED, NOT AUTHENTICATED, AUTHENTICATED UNDER DURESS, NOT
	// APPLICABLE):
	//
	// AUTHENTICATED: Confirmed Friend
	//
	// NOT AUTHENTICATED: Unconfirmed status
	//
	// AUTHENTICATED UNDER DURESS: Authentication comprised by hostiles.
	//
	// NOT APPLICABLE: Authentication not required.
	AuthStatus param.Opt[string] `json:"authStatus,omitzero"`
	// Flag indicating whether a radio identifier is reported.
	BeaconInd param.Opt[bool] `json:"beaconInd,omitzero"`
	// The call sign of the personnel to be recovered.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq1 param.Opt[string] `json:"commEq1,omitzero"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq2 param.Opt[string] `json:"commEq2,omitzero"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq3 param.Opt[string] `json:"commEq3,omitzero"`
	// The survivor service identity (UNKNOWN MILITARY, UNKNOWN CIVILIAN, FRIEND
	// MILITARY, FRIEND CIVIILIAN, NEUTRAL MILITARY, NEUTRAL CIVILIAN, HOSTILE
	// MILITARY, HOSTILE CIVILIAN).
	Identity param.Opt[string] `json:"identity,omitzero"`
	// Unique identifier of a weather report associated with this recovery.
	IDWeatherReport param.Opt[string] `json:"idWeatherReport,omitzero"`
	// The military classification of the personnel to be recovered. Intended as, but
	// not constrained to, MIL-STD-6016 J6.1 Isolated Personnel Classification (NO
	// STATEMENT, MILITARY, GOVERNMENT CIVILIAN, GOVERNMENT CONTRACTOR, CIVILIAN,
	// MULTIPLE CLASSIFICATIONS).
	MilClass param.Opt[string] `json:"milClass,omitzero"`
	// The country of origin or political entity of an isolated person subject to
	// rescue or evacuation. If natAlliance is set to 126, then natAlliance1 must be
	// non 0. If natAlliance is any number other than 126, then natAlliance1 will be
	// set to 0 regardless. Defined in MIL-STD-6016 J6.1 Nationality/Alliance isolated
	// person(s).
	NatAlliance param.Opt[int64] `json:"natAlliance,omitzero"`
	// Extended country of origin or political entity of an isolated person subject to
	// rescue or evacuation. Specify an entry here only if natAlliance is 126. Defined
	// in MIL-STD-6016 J6.1 Nationality/Alliance isolated person(s), 1.
	NatAlliance1 param.Opt[int64] `json:"natAlliance1,omitzero"`
	// Number of ambulatory personnel requiring recovery.
	NumAmbulatory param.Opt[int64] `json:"numAmbulatory,omitzero"`
	// Number of injured, but ambulatory, personnel requiring recovery.
	NumAmbulatoryInjured param.Opt[int64] `json:"numAmbulatoryInjured,omitzero"`
	// Number of littered personnel requiring recovery.
	NumNonAmbulatory param.Opt[int64] `json:"numNonAmbulatory,omitzero"`
	// The count of persons requiring recovery.
	NumPersons param.Opt[int64] `json:"numPersons,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Altitude relative to WGS-84 ellipsoid, in meters. Positive values indicate a
	// point height above ellipsoid, and negative values indicate a point eight below
	// ellipsoid.
	PickupAlt param.Opt[float64] `json:"pickupAlt,omitzero"`
	// UUID identifying the Personnel Recovery mission, which should remain the same on
	// subsequent posts related to the same recovery mission.
	RecovID param.Opt[string] `json:"recovId,omitzero"`
	// Receive voice frequency in 5Hz increments. This field will auto populate with
	// the txFreq value if the post element is null.
	RxFreq param.Opt[float64] `json:"rxFreq,omitzero"`
	// Preloaded message conveying the situation confronting the isolated person(s).
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Survivor Radio Messages
	// (e.g. INJURED CANT MOVE NO KNOWN HOSTILES, INJURED CANT MOVE HOSTILES NEARBY,
	// UNINJURED CANT MOVE HOSTILES NEARBY, UNINJURED NO KNOWN HOSTILES, INJURED
	// LIMITED MOBILITY).
	SurvivorMessages param.Opt[string] `json:"survivorMessages,omitzero"`
	// Survivor radio equipment. Intended as, but not constrained to, MIL-STD-6016 J6.1
	// Survivor Radio Type (NO STATEMENT, PRQ7SEL, PRC90, PRC112, PRC112B B1, PRC112C,
	// PRC112D, PRC148 MBITR, PRC148 JEM, PRC149, PRC152, ACRPLB, OTHER).
	SurvivorRadio param.Opt[string] `json:"survivorRadio,omitzero"`
	// Flag indicating the cancellation of this recovery.
	TermInd param.Opt[bool] `json:"termInd,omitzero"`
	// Additional specific messages received from survivor.
	TextMsg param.Opt[string] `json:"textMsg,omitzero"`
	// Transmit voice frequency in 5Hz increments.
	TxFreq            param.Opt[float64]                          `json:"txFreq,omitzero"`
	ExecutionInfo     PersonnelrecoveryNewParamsExecutionInfo     `json:"executionInfo,omitzero"`
	ObjectiveAreaInfo PersonnelrecoveryNewParamsObjectiveAreaInfo `json:"objectiveAreaInfo,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewParams) UnmarshalJSON(data []byte) error {
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
type PersonnelrecoveryNewParamsDataMode string

const (
	PersonnelrecoveryNewParamsDataModeReal      PersonnelrecoveryNewParamsDataMode = "REAL"
	PersonnelrecoveryNewParamsDataModeTest      PersonnelrecoveryNewParamsDataMode = "TEST"
	PersonnelrecoveryNewParamsDataModeSimulated PersonnelrecoveryNewParamsDataMode = "SIMULATED"
	PersonnelrecoveryNewParamsDataModeExercise  PersonnelrecoveryNewParamsDataMode = "EXERCISE"
)

type PersonnelrecoveryNewParamsExecutionInfo struct {
	// The heading, in degrees, of leaving the recovery zone.
	Egress param.Opt[float64] `json:"egress,omitzero"`
	// The heading, in degrees clockwise from North, of entering the recovery zone.
	Ingress param.Opt[float64] `json:"ingress,omitzero"`
	// Description of the objective strategy plan.
	ObjStrategy param.Opt[string] `json:"objStrategy,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the egress location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	EgressPoint   []float64                                            `json:"egressPoint,omitzero"`
	EscortVehicle PersonnelrecoveryNewParamsExecutionInfoEscortVehicle `json:"escortVehicle,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the initial location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	InitialPoint    []float64                                              `json:"initialPoint,omitzero"`
	RecoveryVehicle PersonnelrecoveryNewParamsExecutionInfoRecoveryVehicle `json:"recoveryVehicle,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewParamsExecutionInfo) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewParamsExecutionInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewParamsExecutionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryNewParamsExecutionInfoEscortVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq param.Opt[float64] `json:"primaryFreq,omitzero"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength param.Opt[int64] `json:"strength,omitzero"`
	// The particular type of recovery vehicle to be used.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewParamsExecutionInfoEscortVehicle) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewParamsExecutionInfoEscortVehicle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewParamsExecutionInfoEscortVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryNewParamsExecutionInfoRecoveryVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq param.Opt[float64] `json:"primaryFreq,omitzero"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength param.Opt[int64] `json:"strength,omitzero"`
	// The particular type of recovery vehicle to be used.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewParamsExecutionInfoRecoveryVehicle) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewParamsExecutionInfoRecoveryVehicle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewParamsExecutionInfoRecoveryVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryNewParamsObjectiveAreaInfo struct {
	// The call sign of the on-scene commander.
	OscCallSign param.Opt[string] `json:"oscCallSign,omitzero"`
	// The radio frequency of the on-scene commander.
	OscFreq param.Opt[float64] `json:"oscFreq,omitzero"`
	// Description of the pickup zone location.
	PzDesc param.Opt[string] `json:"pzDesc,omitzero"`
	// Information detailing knowledge of enemies in the area.
	EnemyData []PersonnelrecoveryNewParamsObjectiveAreaInfoEnemyData `json:"enemyData,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the pz location. This array must contain a
	// minimum of 2 elements (latitude and longitude), and may contain an optional 3rd
	// element (altitude).
	PzLocation []float64 `json:"pzLocation,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewParamsObjectiveAreaInfo) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewParamsObjectiveAreaInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewParamsObjectiveAreaInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryNewParamsObjectiveAreaInfoEnemyData struct {
	// Directions to known enemies in the operation area (NORTH, NORTHEAST, EAST,
	// SOUTHEAST, SOUTH, SOUTHWEST, WEST, NORTHWEST, SURROUNDED).
	DirToEnemy param.Opt[string] `json:"dirToEnemy,omitzero"`
	// Comments provided by friendlies about the evac zone.
	FriendliesRemarks param.Opt[string] `json:"friendliesRemarks,omitzero"`
	// Hot Landing Zone remarks.
	HlzRemarks param.Opt[string] `json:"hlzRemarks,omitzero"`
	// The type of hostile fire received (SMALL ARMS, MORTAR, ARTILLERY, ROCKETS).
	HostileFireType param.Opt[string] `json:"hostileFireType,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewParamsObjectiveAreaInfoEnemyData) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewParamsObjectiveAreaInfoEnemyData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewParamsObjectiveAreaInfoEnemyData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryListParams struct {
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonnelrecoveryListParams]'s query parameters as
// `url.Values`.
func (r PersonnelrecoveryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PersonnelrecoveryCountParams struct {
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonnelrecoveryCountParams]'s query parameters as
// `url.Values`.
func (r PersonnelrecoveryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PersonnelrecoveryNewBulkParams struct {
	Body []PersonnelrecoveryNewBulkParamsBody
	paramObj
}

func (r PersonnelrecoveryNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PersonnelrecoveryNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Provides information concerning search and rescue operations and other
// situations involving personnel recovery.
//
// The properties ClassificationMarking, DataMode, MsgTime, PickupLat, PickupLon,
// Source, Type are required.
type PersonnelrecoveryNewBulkParamsBody struct {
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
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// WGS-84 latitude of the pickup location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PickupLat float64 `json:"pickupLat,required"`
	// WGS-84 longitude of the pickup location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PickupLon float64 `json:"pickupLon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Specifies the type of incident resulting in a recovery or evacuation mission.
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Emergency Type (e.g. NO
	// STATEMENT, DOWN AIRCRAFT, MAN IN WATER, DITCHING, BAILOUT, DISTRESSED VEHICLE,
	// GROUND INCIDENT, MEDICAL, ISOLATED PERSONS, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Mechanism used to verify the survivors identity.
	AuthMethod param.Opt[string] `json:"authMethod,omitzero"`
	// The confirmation status of the isolated personnel identity. Intended as, but not
	// constrained to, MIL-STD-6016 J6.1 Authentication Status, Isolated Personnel (NO
	// STATEMENT, AUTHENTICATED, NOT AUTHENTICATED, AUTHENTICATED UNDER DURESS, NOT
	// APPLICABLE):
	//
	// AUTHENTICATED: Confirmed Friend
	//
	// NOT AUTHENTICATED: Unconfirmed status
	//
	// AUTHENTICATED UNDER DURESS: Authentication comprised by hostiles.
	//
	// NOT APPLICABLE: Authentication not required.
	AuthStatus param.Opt[string] `json:"authStatus,omitzero"`
	// Flag indicating whether a radio identifier is reported.
	BeaconInd param.Opt[bool] `json:"beaconInd,omitzero"`
	// The call sign of the personnel to be recovered.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq1 param.Opt[string] `json:"commEq1,omitzero"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq2 param.Opt[string] `json:"commEq2,omitzero"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq3 param.Opt[string] `json:"commEq3,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The survivor service identity (UNKNOWN MILITARY, UNKNOWN CIVILIAN, FRIEND
	// MILITARY, FRIEND CIVIILIAN, NEUTRAL MILITARY, NEUTRAL CIVILIAN, HOSTILE
	// MILITARY, HOSTILE CIVILIAN).
	Identity param.Opt[string] `json:"identity,omitzero"`
	// Unique identifier of a weather report associated with this recovery.
	IDWeatherReport param.Opt[string] `json:"idWeatherReport,omitzero"`
	// The military classification of the personnel to be recovered. Intended as, but
	// not constrained to, MIL-STD-6016 J6.1 Isolated Personnel Classification (NO
	// STATEMENT, MILITARY, GOVERNMENT CIVILIAN, GOVERNMENT CONTRACTOR, CIVILIAN,
	// MULTIPLE CLASSIFICATIONS).
	MilClass param.Opt[string] `json:"milClass,omitzero"`
	// The country of origin or political entity of an isolated person subject to
	// rescue or evacuation. If natAlliance is set to 126, then natAlliance1 must be
	// non 0. If natAlliance is any number other than 126, then natAlliance1 will be
	// set to 0 regardless. Defined in MIL-STD-6016 J6.1 Nationality/Alliance isolated
	// person(s).
	NatAlliance param.Opt[int64] `json:"natAlliance,omitzero"`
	// Extended country of origin or political entity of an isolated person subject to
	// rescue or evacuation. Specify an entry here only if natAlliance is 126. Defined
	// in MIL-STD-6016 J6.1 Nationality/Alliance isolated person(s), 1.
	NatAlliance1 param.Opt[int64] `json:"natAlliance1,omitzero"`
	// Number of ambulatory personnel requiring recovery.
	NumAmbulatory param.Opt[int64] `json:"numAmbulatory,omitzero"`
	// Number of injured, but ambulatory, personnel requiring recovery.
	NumAmbulatoryInjured param.Opt[int64] `json:"numAmbulatoryInjured,omitzero"`
	// Number of littered personnel requiring recovery.
	NumNonAmbulatory param.Opt[int64] `json:"numNonAmbulatory,omitzero"`
	// The count of persons requiring recovery.
	NumPersons param.Opt[int64] `json:"numPersons,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Altitude relative to WGS-84 ellipsoid, in meters. Positive values indicate a
	// point height above ellipsoid, and negative values indicate a point eight below
	// ellipsoid.
	PickupAlt param.Opt[float64] `json:"pickupAlt,omitzero"`
	// UUID identifying the Personnel Recovery mission, which should remain the same on
	// subsequent posts related to the same recovery mission.
	RecovID param.Opt[string] `json:"recovId,omitzero"`
	// Receive voice frequency in 5Hz increments. This field will auto populate with
	// the txFreq value if the post element is null.
	RxFreq param.Opt[float64] `json:"rxFreq,omitzero"`
	// Preloaded message conveying the situation confronting the isolated person(s).
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Survivor Radio Messages
	// (e.g. INJURED CANT MOVE NO KNOWN HOSTILES, INJURED CANT MOVE HOSTILES NEARBY,
	// UNINJURED CANT MOVE HOSTILES NEARBY, UNINJURED NO KNOWN HOSTILES, INJURED
	// LIMITED MOBILITY).
	SurvivorMessages param.Opt[string] `json:"survivorMessages,omitzero"`
	// Survivor radio equipment. Intended as, but not constrained to, MIL-STD-6016 J6.1
	// Survivor Radio Type (NO STATEMENT, PRQ7SEL, PRC90, PRC112, PRC112B B1, PRC112C,
	// PRC112D, PRC148 MBITR, PRC148 JEM, PRC149, PRC152, ACRPLB, OTHER).
	SurvivorRadio param.Opt[string] `json:"survivorRadio,omitzero"`
	// Flag indicating the cancellation of this recovery.
	TermInd param.Opt[bool] `json:"termInd,omitzero"`
	// Additional specific messages received from survivor.
	TextMsg param.Opt[string] `json:"textMsg,omitzero"`
	// Transmit voice frequency in 5Hz increments.
	TxFreq            param.Opt[float64]                                  `json:"txFreq,omitzero"`
	ExecutionInfo     PersonnelrecoveryNewBulkParamsBodyExecutionInfo     `json:"executionInfo,omitzero"`
	ObjectiveAreaInfo PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfo `json:"objectiveAreaInfo,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PersonnelrecoveryNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type PersonnelrecoveryNewBulkParamsBodyExecutionInfo struct {
	// The heading, in degrees, of leaving the recovery zone.
	Egress param.Opt[float64] `json:"egress,omitzero"`
	// The heading, in degrees clockwise from North, of entering the recovery zone.
	Ingress param.Opt[float64] `json:"ingress,omitzero"`
	// Description of the objective strategy plan.
	ObjStrategy param.Opt[string] `json:"objStrategy,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the egress location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	EgressPoint   []float64                                                    `json:"egressPoint,omitzero"`
	EscortVehicle PersonnelrecoveryNewBulkParamsBodyExecutionInfoEscortVehicle `json:"escortVehicle,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the initial location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	InitialPoint    []float64                                                      `json:"initialPoint,omitzero"`
	RecoveryVehicle PersonnelrecoveryNewBulkParamsBodyExecutionInfoRecoveryVehicle `json:"recoveryVehicle,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewBulkParamsBodyExecutionInfo) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewBulkParamsBodyExecutionInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewBulkParamsBodyExecutionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryNewBulkParamsBodyExecutionInfoEscortVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq param.Opt[float64] `json:"primaryFreq,omitzero"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength param.Opt[int64] `json:"strength,omitzero"`
	// The particular type of recovery vehicle to be used.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewBulkParamsBodyExecutionInfoEscortVehicle) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewBulkParamsBodyExecutionInfoEscortVehicle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewBulkParamsBodyExecutionInfoEscortVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryNewBulkParamsBodyExecutionInfoRecoveryVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq param.Opt[float64] `json:"primaryFreq,omitzero"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength param.Opt[int64] `json:"strength,omitzero"`
	// The particular type of recovery vehicle to be used.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewBulkParamsBodyExecutionInfoRecoveryVehicle) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewBulkParamsBodyExecutionInfoRecoveryVehicle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewBulkParamsBodyExecutionInfoRecoveryVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfo struct {
	// The call sign of the on-scene commander.
	OscCallSign param.Opt[string] `json:"oscCallSign,omitzero"`
	// The radio frequency of the on-scene commander.
	OscFreq param.Opt[float64] `json:"oscFreq,omitzero"`
	// Description of the pickup zone location.
	PzDesc param.Opt[string] `json:"pzDesc,omitzero"`
	// Information detailing knowledge of enemies in the area.
	EnemyData []PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfoEnemyData `json:"enemyData,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the pz location. This array must contain a
	// minimum of 2 elements (latitude and longitude), and may contain an optional 3rd
	// element (altitude).
	PzLocation []float64 `json:"pzLocation,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfo) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfoEnemyData struct {
	// Directions to known enemies in the operation area (NORTH, NORTHEAST, EAST,
	// SOUTHEAST, SOUTH, SOUTHWEST, WEST, NORTHWEST, SURROUNDED).
	DirToEnemy param.Opt[string] `json:"dirToEnemy,omitzero"`
	// Comments provided by friendlies about the evac zone.
	FriendliesRemarks param.Opt[string] `json:"friendliesRemarks,omitzero"`
	// Hot Landing Zone remarks.
	HlzRemarks param.Opt[string] `json:"hlzRemarks,omitzero"`
	// The type of hostile fire received (SMALL ARMS, MORTAR, ARTILLERY, ROCKETS).
	HostileFireType param.Opt[string] `json:"hostileFireType,omitzero"`
	paramObj
}

func (r PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfoEnemyData) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfoEnemyData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfoEnemyData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryFileNewParams struct {
	Body []PersonnelrecoveryFileNewParamsBody
	paramObj
}

func (r PersonnelrecoveryFileNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PersonnelrecoveryFileNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Provides information concerning search and rescue operations and other
// situations involving personnel recovery.
//
// The properties ClassificationMarking, DataMode, MsgTime, PickupLat, PickupLon,
// Source, Type are required.
type PersonnelrecoveryFileNewParamsBody struct {
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
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// WGS-84 latitude of the pickup location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PickupLat float64 `json:"pickupLat,required"`
	// WGS-84 longitude of the pickup location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PickupLon float64 `json:"pickupLon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Specifies the type of incident resulting in a recovery or evacuation mission.
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Emergency Type (e.g. NO
	// STATEMENT, DOWN AIRCRAFT, MAN IN WATER, DITCHING, BAILOUT, DISTRESSED VEHICLE,
	// GROUND INCIDENT, MEDICAL, ISOLATED PERSONS, etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Mechanism used to verify the survivors identity.
	AuthMethod param.Opt[string] `json:"authMethod,omitzero"`
	// The confirmation status of the isolated personnel identity. Intended as, but not
	// constrained to, MIL-STD-6016 J6.1 Authentication Status, Isolated Personnel (NO
	// STATEMENT, AUTHENTICATED, NOT AUTHENTICATED, AUTHENTICATED UNDER DURESS, NOT
	// APPLICABLE):
	//
	// AUTHENTICATED: Confirmed Friend
	//
	// NOT AUTHENTICATED: Unconfirmed status
	//
	// AUTHENTICATED UNDER DURESS: Authentication comprised by hostiles.
	//
	// NOT APPLICABLE: Authentication not required.
	AuthStatus param.Opt[string] `json:"authStatus,omitzero"`
	// Flag indicating whether a radio identifier is reported.
	BeaconInd param.Opt[bool] `json:"beaconInd,omitzero"`
	// The call sign of the personnel to be recovered.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq1 param.Opt[string] `json:"commEq1,omitzero"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq2 param.Opt[string] `json:"commEq2,omitzero"`
	// Survivor communications equipment. Intended as, but not constrained to,
	// MIL-STD-6016 J6.1 Communications Equipment, Isolated Personnel (NO STATEMENT,
	// SURVIVAL RADIO, RADIO BEACON, EPLRS, SIGNAL MIRROR, SMOKE FLARE, IR SIGNALLING
	// DEVICE, SIGNALLING PANEL, FRIENDLY FORCE TRACKER, GPS BEACON, LL PHONE, TACTICAL
	// RADIO LOS, TACTICAL RADIO BLOS).
	CommEq3 param.Opt[string] `json:"commEq3,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The survivor service identity (UNKNOWN MILITARY, UNKNOWN CIVILIAN, FRIEND
	// MILITARY, FRIEND CIVIILIAN, NEUTRAL MILITARY, NEUTRAL CIVILIAN, HOSTILE
	// MILITARY, HOSTILE CIVILIAN).
	Identity param.Opt[string] `json:"identity,omitzero"`
	// Unique identifier of a weather report associated with this recovery.
	IDWeatherReport param.Opt[string] `json:"idWeatherReport,omitzero"`
	// The military classification of the personnel to be recovered. Intended as, but
	// not constrained to, MIL-STD-6016 J6.1 Isolated Personnel Classification (NO
	// STATEMENT, MILITARY, GOVERNMENT CIVILIAN, GOVERNMENT CONTRACTOR, CIVILIAN,
	// MULTIPLE CLASSIFICATIONS).
	MilClass param.Opt[string] `json:"milClass,omitzero"`
	// The country of origin or political entity of an isolated person subject to
	// rescue or evacuation. If natAlliance is set to 126, then natAlliance1 must be
	// non 0. If natAlliance is any number other than 126, then natAlliance1 will be
	// set to 0 regardless. Defined in MIL-STD-6016 J6.1 Nationality/Alliance isolated
	// person(s).
	NatAlliance param.Opt[int64] `json:"natAlliance,omitzero"`
	// Extended country of origin or political entity of an isolated person subject to
	// rescue or evacuation. Specify an entry here only if natAlliance is 126. Defined
	// in MIL-STD-6016 J6.1 Nationality/Alliance isolated person(s), 1.
	NatAlliance1 param.Opt[int64] `json:"natAlliance1,omitzero"`
	// Number of ambulatory personnel requiring recovery.
	NumAmbulatory param.Opt[int64] `json:"numAmbulatory,omitzero"`
	// Number of injured, but ambulatory, personnel requiring recovery.
	NumAmbulatoryInjured param.Opt[int64] `json:"numAmbulatoryInjured,omitzero"`
	// Number of littered personnel requiring recovery.
	NumNonAmbulatory param.Opt[int64] `json:"numNonAmbulatory,omitzero"`
	// The count of persons requiring recovery.
	NumPersons param.Opt[int64] `json:"numPersons,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Altitude relative to WGS-84 ellipsoid, in meters. Positive values indicate a
	// point height above ellipsoid, and negative values indicate a point eight below
	// ellipsoid.
	PickupAlt param.Opt[float64] `json:"pickupAlt,omitzero"`
	// UUID identifying the Personnel Recovery mission, which should remain the same on
	// subsequent posts related to the same recovery mission.
	RecovID param.Opt[string] `json:"recovId,omitzero"`
	// Receive voice frequency in 5Hz increments. This field will auto populate with
	// the txFreq value if the post element is null.
	RxFreq param.Opt[float64] `json:"rxFreq,omitzero"`
	// Preloaded message conveying the situation confronting the isolated person(s).
	// Intended as, but not constrained to, MIL-STD-6016 J6.1 Survivor Radio Messages
	// (e.g. INJURED CANT MOVE NO KNOWN HOSTILES, INJURED CANT MOVE HOSTILES NEARBY,
	// UNINJURED CANT MOVE HOSTILES NEARBY, UNINJURED NO KNOWN HOSTILES, INJURED
	// LIMITED MOBILITY).
	SurvivorMessages param.Opt[string] `json:"survivorMessages,omitzero"`
	// Survivor radio equipment. Intended as, but not constrained to, MIL-STD-6016 J6.1
	// Survivor Radio Type (NO STATEMENT, PRQ7SEL, PRC90, PRC112, PRC112B B1, PRC112C,
	// PRC112D, PRC148 MBITR, PRC148 JEM, PRC149, PRC152, ACRPLB, OTHER).
	SurvivorRadio param.Opt[string] `json:"survivorRadio,omitzero"`
	// Flag indicating the cancellation of this recovery.
	TermInd param.Opt[bool] `json:"termInd,omitzero"`
	// Additional specific messages received from survivor.
	TextMsg param.Opt[string] `json:"textMsg,omitzero"`
	// Transmit voice frequency in 5Hz increments.
	TxFreq            param.Opt[float64]                                  `json:"txFreq,omitzero"`
	ExecutionInfo     PersonnelrecoveryFileNewParamsBodyExecutionInfo     `json:"executionInfo,omitzero"`
	ObjectiveAreaInfo PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfo `json:"objectiveAreaInfo,omitzero"`
	paramObj
}

func (r PersonnelrecoveryFileNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryFileNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryFileNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PersonnelrecoveryFileNewParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type PersonnelrecoveryFileNewParamsBodyExecutionInfo struct {
	// The heading, in degrees, of leaving the recovery zone.
	Egress param.Opt[float64] `json:"egress,omitzero"`
	// The heading, in degrees clockwise from North, of entering the recovery zone.
	Ingress param.Opt[float64] `json:"ingress,omitzero"`
	// Description of the objective strategy plan.
	ObjStrategy param.Opt[string] `json:"objStrategy,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the egress location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	EgressPoint   []float64                                                    `json:"egressPoint,omitzero"`
	EscortVehicle PersonnelrecoveryFileNewParamsBodyExecutionInfoEscortVehicle `json:"escortVehicle,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the initial location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	InitialPoint    []float64                                                      `json:"initialPoint,omitzero"`
	RecoveryVehicle PersonnelrecoveryFileNewParamsBodyExecutionInfoRecoveryVehicle `json:"recoveryVehicle,omitzero"`
	paramObj
}

func (r PersonnelrecoveryFileNewParamsBodyExecutionInfo) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryFileNewParamsBodyExecutionInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryFileNewParamsBodyExecutionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryFileNewParamsBodyExecutionInfoEscortVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq param.Opt[float64] `json:"primaryFreq,omitzero"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength param.Opt[int64] `json:"strength,omitzero"`
	// The particular type of recovery vehicle to be used.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r PersonnelrecoveryFileNewParamsBodyExecutionInfoEscortVehicle) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryFileNewParamsBodyExecutionInfoEscortVehicle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryFileNewParamsBodyExecutionInfoEscortVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryFileNewParamsBodyExecutionInfoRecoveryVehicle struct {
	// The call sign of the recovery vehicle.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Primary contact frequency of the recovery vehicle.
	PrimaryFreq param.Opt[float64] `json:"primaryFreq,omitzero"`
	// The number of objects or units moving as a group and represented as a single
	// entity in this recovery vehicle message. If null, the strength is assumed to
	// represent a single object. Note that if this recovery derives from a J-series
	// message then special definitions apply for the following values: 13 indicates an
	// estimated 2-7 units, 14 indicates an estimated more than 7 units, and 15
	// indicates an estimated more than 12 units.
	Strength param.Opt[int64] `json:"strength,omitzero"`
	// The particular type of recovery vehicle to be used.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r PersonnelrecoveryFileNewParamsBodyExecutionInfoRecoveryVehicle) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryFileNewParamsBodyExecutionInfoRecoveryVehicle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryFileNewParamsBodyExecutionInfoRecoveryVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfo struct {
	// The call sign of the on-scene commander.
	OscCallSign param.Opt[string] `json:"oscCallSign,omitzero"`
	// The radio frequency of the on-scene commander.
	OscFreq param.Opt[float64] `json:"oscFreq,omitzero"`
	// Description of the pickup zone location.
	PzDesc param.Opt[string] `json:"pzDesc,omitzero"`
	// Information detailing knowledge of enemies in the area.
	EnemyData []PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfoEnemyData `json:"enemyData,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the pz location. This array must contain a
	// minimum of 2 elements (latitude and longitude), and may contain an optional 3rd
	// element (altitude).
	PzLocation []float64 `json:"pzLocation,omitzero"`
	paramObj
}

func (r PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfo) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfoEnemyData struct {
	// Directions to known enemies in the operation area (NORTH, NORTHEAST, EAST,
	// SOUTHEAST, SOUTH, SOUTHWEST, WEST, NORTHWEST, SURROUNDED).
	DirToEnemy param.Opt[string] `json:"dirToEnemy,omitzero"`
	// Comments provided by friendlies about the evac zone.
	FriendliesRemarks param.Opt[string] `json:"friendliesRemarks,omitzero"`
	// Hot Landing Zone remarks.
	HlzRemarks param.Opt[string] `json:"hlzRemarks,omitzero"`
	// The type of hostile fire received (SMALL ARMS, MORTAR, ARTILLERY, ROCKETS).
	HostileFireType param.Opt[string] `json:"hostileFireType,omitzero"`
	paramObj
}

func (r PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfoEnemyData) MarshalJSON() (data []byte, err error) {
	type shadow PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfoEnemyData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfoEnemyData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonnelrecoveryGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonnelrecoveryGetParams]'s query parameters as
// `url.Values`.
func (r PersonnelrecoveryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PersonnelrecoveryTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time stamp of the original personnel recovery message, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonnelrecoveryTupleParams]'s query parameters as
// `url.Values`.
func (r PersonnelrecoveryTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
