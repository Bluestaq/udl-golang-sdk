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

// EvacService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvacService] method instead.
type EvacService struct {
	Options []option.RequestOption
	History EvacHistoryService
	Tuple   EvacTupleService
}

// NewEvacService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvacService(opts ...option.RequestOption) (r EvacService) {
	r = EvacService{}
	r.Options = opts
	r.History = NewEvacHistoryService(opts...)
	r.Tuple = NewEvacTupleService(opts...)
	return
}

// Service operation to take a single evac as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *EvacService) New(ctx context.Context, body EvacNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/evac"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Evac by its unique ID passed as a path
// parameter.
func (r *EvacService) Get(ctx context.Context, id string, query EvacGetParams, opts ...option.RequestOption) (res *shared.EvacFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/evac/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EvacService) List(ctx context.Context, query EvacListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EvacAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/evac"
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
func (r *EvacService) ListAutoPaging(ctx context.Context, query EvacListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EvacAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EvacService) Count(ctx context.Context, query EvacCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/evac/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of Evac
// records as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *EvacService) NewBulk(ctx context.Context, body EvacNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/evac/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EvacService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/evac/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Service operation to take a list of Evac events as a POST body and ingest into
// the database. Requires a specific role, please contact the UDL team to gain
// access. This operation is intended to be used for automated feeds into UDL.
func (r *EvacService) UnvalidatedPublish(ctx context.Context, body EvacUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-evac"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Casualty report and evacuation request. Used to report and request support to
// evacuate friendly and enemy casualties.
type EvacAbridged struct {
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
	DataMode EvacAbridgedDataMode `json:"dataMode,required"`
	// WGS-84 latitude of the pickup location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PickupLat float64 `json:"pickupLat,required"`
	// WGS-84 longitude of the pickup location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PickupLon float64 `json:"pickupLon,required"`
	// The request time, in ISO 8601 UTC format.
	ReqTime time.Time `json:"reqTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of this medevac record (REQUEST, RESPONSE).
	//
	// Any of "REQUEST", "RESPONSE".
	Type EvacAbridgedType `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Identity and medical information on the patient to be evacuated.
	CasualtyInfo []EvacAbridgedCasualtyInfo `json:"casualtyInfo"`
	// Radius of circular area about lat/lon point, in meters (1-sigma, if representing
	// error).
	Ce float64 `json:"ce"`
	// The contact frequency, in Hz, of the agency or zone controller.
	CntctFreq float64 `json:"cntctFreq"`
	// Additional comments for the medevac mission.
	Comments string `json:"comments"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Data defining any enemy intelligence reported by the requestor.
	EnemyData []EvacAbridgedEnemyData `json:"enemyData"`
	// Unique identifier of a weather report associated with this evacuation.
	IDWeatherReport string `json:"idWeatherReport"`
	// Height above lat/lon point, in meters (1-sigma, if representing linear error).
	Le float64 `json:"le"`
	// UUID identifying the medevac mission, which should remain the same on subsequent
	// posts related to the same medevac mission.
	MedevacID string `json:"medevacId"`
	// Flag indicating whether the mission requires medical personnel.
	MedicReq bool `json:"medicReq"`
	// The operation type of the evacuation. (NOT SPECIFIED, AIR, GROUND, SURFACE).
	MissionType string `json:"missionType"`
	// Number of ambulatory personnel requiring evacuation.
	NumAmbulatory int64 `json:"numAmbulatory"`
	// The count of people requiring medevac.
	NumCasualties int64 `json:"numCasualties"`
	// Number of people Killed In Action.
	NumKia int64 `json:"numKIA"`
	// Number of littered personnel requiring evacuation.
	NumLitter int64 `json:"numLitter"`
	// Number of people Wounded In Action.
	NumWia int64 `json:"numWIA"`
	// Amplifying data for the terrain describing important obstacles in or around the
	// zone.
	ObstaclesRemarks string `json:"obstaclesRemarks"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Altitude relative to WGS-84 ellipsoid, in meters. Positive values indicate a
	// point height above ellipsoid, and negative values indicate a point height below
	// ellipsoid.
	PickupAlt float64 `json:"pickupAlt"`
	// The expected pickup time, in ISO 8601 UTC format.
	PickupTime time.Time `json:"pickupTime" format:"date-time"`
	// The call sign of this medevac requestor.
	ReqCallSign string `json:"reqCallSign"`
	// Externally provided Medevac request number (e.g. MED.1.223908).
	ReqNum string `json:"reqNum"`
	// Short description of the terrain features of the pickup location (WOODS, TREES,
	// PLOWED FIELDS, FLAT, STANDING WATER, MARSH, URBAN BUILT-UP AREA, MOUNTAIN, HILL,
	// SAND TD, ROCKY, VALLEY, METAMORPHIC ICE, UNKNOWN TD, SEA, NO STATEMENT).
	Terrain string `json:"terrain"`
	// Amplifying data for the terrain describing any notable additional terrain
	// features.
	TerrainRemarks string `json:"terrainRemarks"`
	// The call sign of the zone controller.
	ZoneContrCallSign string `json:"zoneContrCallSign"`
	// Flag indicating that the pickup site is hot and hostiles are in the area.
	ZoneHot bool `json:"zoneHot"`
	// The expected marker identifying the pickup site (SMOKE ZONE MARKING, FLARES,
	// MIRROR, GLIDE ANGLE INDICATOR LIGHT, LIGHT ZONE MARKING, PANELS, FIRE, LASER
	// DESIGNATOR, STROBE LIGHTS, VEHICLE LIGHTS, COLORED SMOKE, WHITE PHOSPHERUS,
	// INFRARED, ILLUMINATION, FRATRICIDE FENCE).
	ZoneMarking string `json:"zoneMarking"`
	// Color used for the pickup site marking (RED, WHITE, BLUE, YELLOW, GREEN, ORANGE,
	// BLACK, PURPLE, BROWN, TAN, GRAY, SILVER, CAMOUFLAGE, OTHER COLOR).
	ZoneMarkingColor string `json:"zoneMarkingColor"`
	// The name of the zone.
	ZoneName string `json:"zoneName"`
	// The pickup site security (UNKNOWN ZONESECURITY, NO ENEMY, POSSIBLE ENEMY, ENEMY
	// IN AREA USE CAUTION, ENEMY IN AREA ARMED ESCORT REQUIRED).
	ZoneSecurity string `json:"zoneSecurity"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		PickupLat             resp.Field
		PickupLon             resp.Field
		ReqTime               resp.Field
		Source                resp.Field
		Type                  resp.Field
		ID                    resp.Field
		CasualtyInfo          resp.Field
		Ce                    resp.Field
		CntctFreq             resp.Field
		Comments              resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EnemyData             resp.Field
		IDWeatherReport       resp.Field
		Le                    resp.Field
		MedevacID             resp.Field
		MedicReq              resp.Field
		MissionType           resp.Field
		NumAmbulatory         resp.Field
		NumCasualties         resp.Field
		NumKia                resp.Field
		NumLitter             resp.Field
		NumWia                resp.Field
		ObstaclesRemarks      resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		PickupAlt             resp.Field
		PickupTime            resp.Field
		ReqCallSign           resp.Field
		ReqNum                resp.Field
		Terrain               resp.Field
		TerrainRemarks        resp.Field
		ZoneContrCallSign     resp.Field
		ZoneHot               resp.Field
		ZoneMarking           resp.Field
		ZoneMarkingColor      resp.Field
		ZoneName              resp.Field
		ZoneSecurity          resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridged) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridged) UnmarshalJSON(data []byte) error {
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
type EvacAbridgedDataMode string

const (
	EvacAbridgedDataModeReal      EvacAbridgedDataMode = "REAL"
	EvacAbridgedDataModeTest      EvacAbridgedDataMode = "TEST"
	EvacAbridgedDataModeSimulated EvacAbridgedDataMode = "SIMULATED"
	EvacAbridgedDataModeExercise  EvacAbridgedDataMode = "EXERCISE"
)

// The type of this medevac record (REQUEST, RESPONSE).
type EvacAbridgedType string

const (
	EvacAbridgedTypeRequest  EvacAbridgedType = "REQUEST"
	EvacAbridgedTypeResponse EvacAbridgedType = "RESPONSE"
)

type EvacAbridgedCasualtyInfo struct {
	// The patient age, in years.
	Age int64 `json:"age"`
	// Allergy information.
	Allergy []EvacAbridgedCasualtyInfoAllergy `json:"allergy"`
	// The patient blood type (A POS, B POS, AB POS, O POS, A NEG, B NEG, AB NEG, O
	// NEG).
	BloodType string `json:"bloodType"`
	// The body part involved for the patient (HEAD, NECK, ABDOMEN, UPPER EXTREMITIES,
	// BACK, FACE, LOWER EXTREMITIES, FRONT, OBSTETRICAL GYNECOLOGICAL, OTHER BODY
	// PART).
	BodyPart string `json:"bodyPart"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the burial location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	BurialLocation []float64 `json:"burialLocation"`
	// The call sign of this patient.
	CallSign string `json:"callSign"`
	// Unique identifier for the patient care provider.
	CareProviderUrn string `json:"careProviderUrn"`
	// Optional casualty key.
	CasualtyKey string `json:"casualtyKey"`
	// The type of medical issue resulting in the need to evacuate the patient (NON
	// BATTLE, CUT, BURN, SICK, FRACTURE, AMPUTATION, PERFORATION, NUCLEAR, EXHAUSTION,
	// BIOLOGICAL, CHEMICAL, SHOCK, PUNCTURE WOUND, OTHER CUT, WOUNDED IN ACTION,
	// DENIAL, COMBAT STRESS).
	CasualtyType string `json:"casualtyType"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the collection point. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	CollectionPoint []float64 `json:"collectionPoint"`
	// Additional comments on the patient's casualty information.
	Comments string `json:"comments"`
	// Health condition information.
	Condition []EvacAbridgedCasualtyInfoCondition `json:"condition"`
	// The contamination specified for the patient (NONE, RADIATION, BIOLOGICAL,
	// CHEMICAL).
	ContamType string `json:"contamType"`
	// The patient's general medical state (SICK IN QUARTERS, RETURN TO DUTY, EVACUATE
	// WOUNDED, EVACUATE DECEASED, INTERRED).
	Disposition string `json:"disposition"`
	// The expected disposition of this patient (R T D, EVACUATE, EVACUATE TO FORWARD
	// SURGICAL TEAM, EVACUATE TO COMBAT SUPPORT HOSPITAL, EVACUATE TO AERO MEDICAL
	// STAGING FACILITY, EVACUATE TO SUSTAINING BASE MEDICAL TREATMENT FACILITY).
	DispositionType string `json:"dispositionType"`
	// Medical condition causation information.
	Etiology []EvacAbridgedCasualtyInfoEtiology `json:"etiology"`
	// The required evacuation method for this patient (AIR, GROUND, NOT EVACUATED).
	EvacType string `json:"evacType"`
	// The patient sex (MALE, FEMALE).
	Gender string `json:"gender"`
	// Health state information.
	HealthState []EvacAbridgedCasualtyInfoHealthState `json:"healthState"`
	// Injury specifics.
	Injury []EvacAbridgedCasualtyInfoInjury `json:"injury"`
	// Last 4 characters of the patient social security code, or equivalent.
	Last4Ssn string `json:"last4SSN"`
	// Medication specifics.
	Medication []EvacAbridgedCasualtyInfoMedication `json:"medication"`
	// The patient common or legal name.
	Name string `json:"name"`
	// The country code indicating the citizenship of the patient.
	Nationality string `json:"nationality"`
	// The career field of this patient.
	OccSpeciality string `json:"occSpeciality"`
	// The patient service identity (UNKNOWN MILITARY, UNKNOWN CIVILIAN, FRIEND
	// MILITARY, FRIEND CIVILIAN, NEUTRAL MILITARY, NEUTRAL CIVILIAN, HOSTILE MILITARY,
	// HOSTILE CIVILIAN).
	PatientIdentity string `json:"patientIdentity"`
	// The patient service status (US MILITARY, US CIVILIAN, NON US MILITARY, NON US
	// CIVILIAN, ENEMY POW).
	PatientStatus string `json:"patientStatus"`
	// The patient pay grade or rank designation (O-10, O-9, O-8, O-7, O-6, O-5, O-4,
	// O-3, O-2, O-1, CWO-5, CWO-4, CWO-2, CWO-1, E -9, E-8, E-7, E-6, E-5, E-4, E-3,
	// E-2, E-1, NONE, CIVILIAN).
	PayGrade string `json:"payGrade"`
	// The priority of the medevac mission for this patient (URGENT, PRIORITY, ROUTINE,
	// URGENT SURGERY, CONVENIENCE).
	Priority string `json:"priority"`
	// The method used to generate this medevac report (DEVICE, GROUND COMBAT
	// PERSONNEL, EVACUATION PERSONNEL, ECHELON1 PERSONNEL, ECHELON2 PERSONNEL).
	ReportGen string `json:"reportGen"`
	// Datetime of the compiling of the patients casualty report, in ISO 8601 UTC
	// format.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// The patient branch of service (AIR FORCE, ARMY, NAVY, MARINES, CIV, CONTR,
	// UNKNOWN SERVICE).
	Service string `json:"service"`
	// Array specifying if any special equipment is need for each of the evacuation of
	// this patient (EXTRACTION EQUIPMENT, SEMI RIGID LITTER, BACKBOARD, CERVICAL
	// COLLAR ,JUNGLE PENETRATOR, OXYGEN, WHOLE BLOOD, VENTILATOR, HOIST, NONE).
	SpecMedEquip []string `json:"specMedEquip"`
	// Treatment information.
	Treatment []EvacAbridgedCasualtyInfoTreatment `json:"treatment"`
	// Information obtained for vital signs.
	VitalSignData []EvacAbridgedCasualtyInfoVitalSignData `json:"vitalSignData"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Age             resp.Field
		Allergy         resp.Field
		BloodType       resp.Field
		BodyPart        resp.Field
		BurialLocation  resp.Field
		CallSign        resp.Field
		CareProviderUrn resp.Field
		CasualtyKey     resp.Field
		CasualtyType    resp.Field
		CollectionPoint resp.Field
		Comments        resp.Field
		Condition       resp.Field
		ContamType      resp.Field
		Disposition     resp.Field
		DispositionType resp.Field
		Etiology        resp.Field
		EvacType        resp.Field
		Gender          resp.Field
		HealthState     resp.Field
		Injury          resp.Field
		Last4Ssn        resp.Field
		Medication      resp.Field
		Name            resp.Field
		Nationality     resp.Field
		OccSpeciality   resp.Field
		PatientIdentity resp.Field
		PatientStatus   resp.Field
		PayGrade        resp.Field
		Priority        resp.Field
		ReportGen       resp.Field
		ReportTime      resp.Field
		Service         resp.Field
		SpecMedEquip    resp.Field
		Treatment       resp.Field
		VitalSignData   resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedCasualtyInfo) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedCasualtyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacAbridgedCasualtyInfoAllergy struct {
	// Additional comments on the patient's allergy information.
	Comments string `json:"comments"`
	// Type of patient allergy (e.g. PENICILLIN, SULFA, OTHER).
	Type string `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Comments    resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedCasualtyInfoAllergy) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedCasualtyInfoAllergy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacAbridgedCasualtyInfoCondition struct {
	// Body part location or body part referenced in condition. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart string `json:"bodyPart"`
	// Additional comments on the patient's condition.
	Comments string `json:"comments"`
	// Datetime of the condition diagnosis in ISO 8601 UTC datetime format.
	Time time.Time `json:"time" format:"date-time"`
	// Health condition assessment. Intended as, but not constrained to, K07.1
	// Condition Type Enumeration (e.g. ACTIVITY HIGH, ACTIVITY LOW, ACTIVITY MEDIUM,
	// ACTIVITY NONE, AVPU ALERT, AVPU ALTERED MENTAL STATE, AVPU PAIN, AVPU
	// UNRESPONSIVE, etc.).
	Type string `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BodyPart    resp.Field
		Comments    resp.Field
		Time        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedCasualtyInfoCondition) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedCasualtyInfoCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacAbridgedCasualtyInfoEtiology struct {
	// The body part or location affected from the etiology. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart string `json:"bodyPart"`
	// Additional comments on the patient's etiology information.
	Comments string `json:"comments"`
	// Datetime of the discovery of the etiology state in ISO 8601 UTC format.
	Time time.Time `json:"time" format:"date-time"`
	// The cause or manner of causation of the medical condition. Intended as, but not
	// constrained to, K07.1 EtiologyType Enumeration (e.g. ASSAULT, BUILDING COLLAPSE,
	// BURN CHEMICAL, BURN ELECTRICAL, BURN, BURN HOT LIQUID, BURN RADIATION, BURN
	// THERMAL, etc.).
	Type string `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BodyPart    resp.Field
		Comments    resp.Field
		Time        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedCasualtyInfoEtiology) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedCasualtyInfoEtiology) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacAbridgedCasualtyInfoHealthState struct {
	// Medical color code used to quickly identify various medical state (e.g. AMBER,
	// BLACK, BLUE, GRAY, NORMAL, RED).
	HealthStateCode string `json:"healthStateCode"`
	// Medical confidence factor.
	MedConfFactor int64 `json:"medConfFactor"`
	// Datetime of the health state diagnosis in ISO 8601 UTC datetime format.
	Time time.Time `json:"time" format:"date-time"`
	// Generalized state of health type (BIOLOGICAL, CHEMICAL, COGNITIVE, HYDRATION,
	// LIFE SIGN, RADIATION, SHOCK, THERMAL).
	Type string `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		HealthStateCode resp.Field
		MedConfFactor   resp.Field
		Time            resp.Field
		Type            resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedCasualtyInfoHealthState) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedCasualtyInfoHealthState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacAbridgedCasualtyInfoInjury struct {
	// Body part location of the injury. Intended as, but not constrained to, K07.1
	// Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE LEFT FRONT, ANKLE RIGHT
	// BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW BACK, ARM LEFT ELBOW
	// FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart string `json:"bodyPart"`
	// Additional comments on the patient's injury information.
	Comments string `json:"comments"`
	// The time of the injury, in ISO 8601 UTC format.
	Time time.Time `json:"time" format:"date-time"`
	// Classification of the injury type (ABRASION, AMPUTATION IT, AVULATION,
	// BALLISTIC, BLAST WAVE, BURN 1ST DEGREE, BURN 2ND DEGREE, BURN 3RD DEGREE, BURN
	// INHALATION, BURN LOWER AIRWAY, CHEST FLAIL, CHEST OPEN, DEGLOVING, ECCHYMOSIS,
	// FRACTURE CLOSED, FRACTURE CREPITUS, FRACTURE IT, FRACTURE OPEN, HEMATOMA,
	// IRREGULAR CONSISTENCY, IRREGULAR CONSISTENCY RIDGED, IRREGULAR CONSISTENCY
	// SWOLLEN, IRREGULAR CONSISTENCY SWOLLEN DISTENDED, IRREGULAR CONSISTENCY TENDER,
	// IRREGULAR POSITION, IRREGULAR SHAPE, IRREGULAR SHAPE MISSHAPED, IRREGULAR SHAPE
	// NON SYMMETRICAL, LACERATION, NEUROVASCULAR COMPROMISE, NEUROVASCULAR INTACT,
	// PUNCTURE, SEAT BELT SIGN, STAB, TIC TIM).
	Type string `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BodyPart    resp.Field
		Comments    resp.Field
		Time        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedCasualtyInfoInjury) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedCasualtyInfoInjury) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacAbridgedCasualtyInfoMedication struct {
	// Route of medication delivery (e.g. INJECTION, ORAL, etc.).
	AdminRoute string `json:"adminRoute"`
	// Body part location or body part referenced for medication. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart string `json:"bodyPart"`
	// Additional comments on the patient's medication information.
	Comments string `json:"comments"`
	// Quantity of medicine or drug administered or recommended to be taken at a
	// particular time.
	Dose string `json:"dose"`
	// The time that the medication was administered in ISO 8601 UTC format.
	Time time.Time `json:"time" format:"date-time"`
	// The type of medication administered. Intended as, but not constrained to, K07.1
	// Medication Enumeration (CEFOTETAN, ABRASION, ABX, AMOXILOXACIN, ANALGESIC,
	// COLLOID, CRYOPECIPITATES, CRYSTALLOID, EPINEPHRINE, ERTAPENEM, FENTANYL,
	// HEXTEND, LACTATED RINGERS, MOBIC, MORPHINE, NARCOTIC, NS, PENICILLIN, PLASMA,
	// PLATELETS, PRBC, TYLENOL, WHOLE BLOOD MT).
	Type string `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AdminRoute  resp.Field
		BodyPart    resp.Field
		Comments    resp.Field
		Dose        resp.Field
		Time        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedCasualtyInfoMedication) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedCasualtyInfoMedication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacAbridgedCasualtyInfoTreatment struct {
	// Body part location or body part treated or to be treated. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart string `json:"bodyPart"`
	// Additional comments on the patient's treatment information.
	Comments string `json:"comments"`
	// Datetime of the treatment in ISO 8601 UTC format.
	Time time.Time `json:"time" format:"date-time"`
	// Type of treatment administered or to be administered. Intended as, but not
	// constrained to, K07.1 Treatment Type Enumeration (e.g. AIRWAY ADJUNCT, AIRWAY
	// ASSISTED VENTILATION, AIRWAY COMBI TUBE USED, AIRWAY ET NT, AIRWAY INTUBATED,
	// AIRWAY NPA OPA APPLIED, AIRWAY PATIENT, AIRWAY POSITIONAL, AIRWAY SURGICAL CRIC,
	// BREATHING CHEST SEAL, BREATHING CHEST TUBE, etc.).
	Type string `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BodyPart    resp.Field
		Comments    resp.Field
		Time        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedCasualtyInfoTreatment) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedCasualtyInfoTreatment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacAbridgedCasualtyInfoVitalSignData struct {
	// Medical confidence factor.
	MedConfFactor int64 `json:"medConfFactor"`
	// Datetime of the vital sign measurement in ISO 8601 UTC datetime format.
	Time time.Time `json:"time" format:"date-time"`
	// Patient vital sign measured (e.g. HEART RATE, PULSE RATE, RESPIRATION RATE,
	// TEMPERATURE CORE, etc.).
	VitalSign string `json:"vitalSign"`
	// Vital sign value 1. The content of this field is dependent on the type of vital
	// sign being measured (see the vitalSign field).
	VitalSign1 float64 `json:"vitalSign1"`
	// Vital sign value 2. The content of this field is dependent on the type of vital
	// sign being measured (see the vitalSign field).
	VitalSign2 float64 `json:"vitalSign2"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		MedConfFactor resp.Field
		Time          resp.Field
		VitalSign     resp.Field
		VitalSign1    resp.Field
		VitalSign2    resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedCasualtyInfoVitalSignData) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedCasualtyInfoVitalSignData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacAbridgedEnemyData struct {
	// Directions to known enemies in the operation area (NORTH, NORTHEAST, EAST,
	// SOUTHEAST, SOUTH, SOUTHWEST, WEST, NORTHWEST, SURROUNDED).
	DirToEnemy string `json:"dirToEnemy"`
	// Comments provided by friendlies about the evac zone.
	FriendliesRemarks string `json:"friendliesRemarks"`
	// Hot Landing Zone remarks.
	HlzRemarks string `json:"hlzRemarks"`
	// The type of hostile fire received (SMALL ARMS, MORTAR, ARTILLERY, ROCKETS).
	HostileFireType string `json:"hostileFireType"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		DirToEnemy        resp.Field
		FriendliesRemarks resp.Field
		HlzRemarks        resp.Field
		HostileFireType   resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvacAbridgedEnemyData) RawJSON() string { return r.JSON.raw }
func (r *EvacAbridgedEnemyData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvacNewParams struct {
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
	DataMode EvacNewParamsDataMode `json:"dataMode,omitzero,required"`
	// WGS-84 latitude of the pickup location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PickupLat float64 `json:"pickupLat,required"`
	// WGS-84 longitude of the pickup location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PickupLon float64 `json:"pickupLon,required"`
	// The request time, in ISO 8601 UTC format.
	ReqTime time.Time `json:"reqTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of this medevac record (REQUEST, RESPONSE).
	//
	// Any of "REQUEST", "RESPONSE".
	Type EvacNewParamsType `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Radius of circular area about lat/lon point, in meters (1-sigma, if representing
	// error).
	Ce param.Opt[float64] `json:"ce,omitzero"`
	// The contact frequency, in Hz, of the agency or zone controller.
	CntctFreq param.Opt[float64] `json:"cntctFreq,omitzero"`
	// Additional comments for the medevac mission.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Unique identifier of a weather report associated with this evacuation.
	IDWeatherReport param.Opt[string] `json:"idWeatherReport,omitzero"`
	// Height above lat/lon point, in meters (1-sigma, if representing linear error).
	Le param.Opt[float64] `json:"le,omitzero"`
	// UUID identifying the medevac mission, which should remain the same on subsequent
	// posts related to the same medevac mission.
	MedevacID param.Opt[string] `json:"medevacId,omitzero"`
	// Flag indicating whether the mission requires medical personnel.
	MedicReq param.Opt[bool] `json:"medicReq,omitzero"`
	// The operation type of the evacuation. (NOT SPECIFIED, AIR, GROUND, SURFACE).
	MissionType param.Opt[string] `json:"missionType,omitzero"`
	// Number of ambulatory personnel requiring evacuation.
	NumAmbulatory param.Opt[int64] `json:"numAmbulatory,omitzero"`
	// The count of people requiring medevac.
	NumCasualties param.Opt[int64] `json:"numCasualties,omitzero"`
	// Number of people Killed In Action.
	NumKia param.Opt[int64] `json:"numKIA,omitzero"`
	// Number of littered personnel requiring evacuation.
	NumLitter param.Opt[int64] `json:"numLitter,omitzero"`
	// Number of people Wounded In Action.
	NumWia param.Opt[int64] `json:"numWIA,omitzero"`
	// Amplifying data for the terrain describing important obstacles in or around the
	// zone.
	ObstaclesRemarks param.Opt[string] `json:"obstaclesRemarks,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Altitude relative to WGS-84 ellipsoid, in meters. Positive values indicate a
	// point height above ellipsoid, and negative values indicate a point height below
	// ellipsoid.
	PickupAlt param.Opt[float64] `json:"pickupAlt,omitzero"`
	// The expected pickup time, in ISO 8601 UTC format.
	PickupTime param.Opt[time.Time] `json:"pickupTime,omitzero" format:"date-time"`
	// The call sign of this medevac requestor.
	ReqCallSign param.Opt[string] `json:"reqCallSign,omitzero"`
	// Externally provided Medevac request number (e.g. MED.1.223908).
	ReqNum param.Opt[string] `json:"reqNum,omitzero"`
	// Short description of the terrain features of the pickup location (WOODS, TREES,
	// PLOWED FIELDS, FLAT, STANDING WATER, MARSH, URBAN BUILT-UP AREA, MOUNTAIN, HILL,
	// SAND TD, ROCKY, VALLEY, METAMORPHIC ICE, UNKNOWN TD, SEA, NO STATEMENT).
	Terrain param.Opt[string] `json:"terrain,omitzero"`
	// Amplifying data for the terrain describing any notable additional terrain
	// features.
	TerrainRemarks param.Opt[string] `json:"terrainRemarks,omitzero"`
	// The call sign of the zone controller.
	ZoneContrCallSign param.Opt[string] `json:"zoneContrCallSign,omitzero"`
	// Flag indicating that the pickup site is hot and hostiles are in the area.
	ZoneHot param.Opt[bool] `json:"zoneHot,omitzero"`
	// The expected marker identifying the pickup site (SMOKE ZONE MARKING, FLARES,
	// MIRROR, GLIDE ANGLE INDICATOR LIGHT, LIGHT ZONE MARKING, PANELS, FIRE, LASER
	// DESIGNATOR, STROBE LIGHTS, VEHICLE LIGHTS, COLORED SMOKE, WHITE PHOSPHERUS,
	// INFRARED, ILLUMINATION, FRATRICIDE FENCE).
	ZoneMarking param.Opt[string] `json:"zoneMarking,omitzero"`
	// Color used for the pickup site marking (RED, WHITE, BLUE, YELLOW, GREEN, ORANGE,
	// BLACK, PURPLE, BROWN, TAN, GRAY, SILVER, CAMOUFLAGE, OTHER COLOR).
	ZoneMarkingColor param.Opt[string] `json:"zoneMarkingColor,omitzero"`
	// The name of the zone.
	ZoneName param.Opt[string] `json:"zoneName,omitzero"`
	// The pickup site security (UNKNOWN ZONESECURITY, NO ENEMY, POSSIBLE ENEMY, ENEMY
	// IN AREA USE CAUTION, ENEMY IN AREA ARMED ESCORT REQUIRED).
	ZoneSecurity param.Opt[string] `json:"zoneSecurity,omitzero"`
	// Identity and medical information on the patient to be evacuated.
	CasualtyInfo []EvacNewParamsCasualtyInfo `json:"casualtyInfo,omitzero"`
	// Data defining any enemy intelligence reported by the requestor.
	EnemyData []EvacNewParamsEnemyData `json:"enemyData,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r EvacNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParams
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
type EvacNewParamsDataMode string

const (
	EvacNewParamsDataModeReal      EvacNewParamsDataMode = "REAL"
	EvacNewParamsDataModeTest      EvacNewParamsDataMode = "TEST"
	EvacNewParamsDataModeSimulated EvacNewParamsDataMode = "SIMULATED"
	EvacNewParamsDataModeExercise  EvacNewParamsDataMode = "EXERCISE"
)

// The type of this medevac record (REQUEST, RESPONSE).
type EvacNewParamsType string

const (
	EvacNewParamsTypeRequest  EvacNewParamsType = "REQUEST"
	EvacNewParamsTypeResponse EvacNewParamsType = "RESPONSE"
)

type EvacNewParamsCasualtyInfo struct {
	// The patient age, in years.
	Age param.Opt[int64] `json:"age,omitzero"`
	// The patient blood type (A POS, B POS, AB POS, O POS, A NEG, B NEG, AB NEG, O
	// NEG).
	BloodType param.Opt[string] `json:"bloodType,omitzero"`
	// The body part involved for the patient (HEAD, NECK, ABDOMEN, UPPER EXTREMITIES,
	// BACK, FACE, LOWER EXTREMITIES, FRONT, OBSTETRICAL GYNECOLOGICAL, OTHER BODY
	// PART).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// The call sign of this patient.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Unique identifier for the patient care provider.
	CareProviderUrn param.Opt[string] `json:"careProviderUrn,omitzero"`
	// Optional casualty key.
	CasualtyKey param.Opt[string] `json:"casualtyKey,omitzero"`
	// The type of medical issue resulting in the need to evacuate the patient (NON
	// BATTLE, CUT, BURN, SICK, FRACTURE, AMPUTATION, PERFORATION, NUCLEAR, EXHAUSTION,
	// BIOLOGICAL, CHEMICAL, SHOCK, PUNCTURE WOUND, OTHER CUT, WOUNDED IN ACTION,
	// DENIAL, COMBAT STRESS).
	CasualtyType param.Opt[string] `json:"casualtyType,omitzero"`
	// Additional comments on the patient's casualty information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// The contamination specified for the patient (NONE, RADIATION, BIOLOGICAL,
	// CHEMICAL).
	ContamType param.Opt[string] `json:"contamType,omitzero"`
	// The patient's general medical state (SICK IN QUARTERS, RETURN TO DUTY, EVACUATE
	// WOUNDED, EVACUATE DECEASED, INTERRED).
	Disposition param.Opt[string] `json:"disposition,omitzero"`
	// The expected disposition of this patient (R T D, EVACUATE, EVACUATE TO FORWARD
	// SURGICAL TEAM, EVACUATE TO COMBAT SUPPORT HOSPITAL, EVACUATE TO AERO MEDICAL
	// STAGING FACILITY, EVACUATE TO SUSTAINING BASE MEDICAL TREATMENT FACILITY).
	DispositionType param.Opt[string] `json:"dispositionType,omitzero"`
	// The required evacuation method for this patient (AIR, GROUND, NOT EVACUATED).
	EvacType param.Opt[string] `json:"evacType,omitzero"`
	// The patient sex (MALE, FEMALE).
	Gender param.Opt[string] `json:"gender,omitzero"`
	// Last 4 characters of the patient social security code, or equivalent.
	Last4Ssn param.Opt[string] `json:"last4SSN,omitzero"`
	// The patient common or legal name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The country code indicating the citizenship of the patient.
	Nationality param.Opt[string] `json:"nationality,omitzero"`
	// The career field of this patient.
	OccSpeciality param.Opt[string] `json:"occSpeciality,omitzero"`
	// The patient service identity (UNKNOWN MILITARY, UNKNOWN CIVILIAN, FRIEND
	// MILITARY, FRIEND CIVILIAN, NEUTRAL MILITARY, NEUTRAL CIVILIAN, HOSTILE MILITARY,
	// HOSTILE CIVILIAN).
	PatientIdentity param.Opt[string] `json:"patientIdentity,omitzero"`
	// The patient service status (US MILITARY, US CIVILIAN, NON US MILITARY, NON US
	// CIVILIAN, ENEMY POW).
	PatientStatus param.Opt[string] `json:"patientStatus,omitzero"`
	// The patient pay grade or rank designation (O-10, O-9, O-8, O-7, O-6, O-5, O-4,
	// O-3, O-2, O-1, CWO-5, CWO-4, CWO-2, CWO-1, E -9, E-8, E-7, E-6, E-5, E-4, E-3,
	// E-2, E-1, NONE, CIVILIAN).
	PayGrade param.Opt[string] `json:"payGrade,omitzero"`
	// The priority of the medevac mission for this patient (URGENT, PRIORITY, ROUTINE,
	// URGENT SURGERY, CONVENIENCE).
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The method used to generate this medevac report (DEVICE, GROUND COMBAT
	// PERSONNEL, EVACUATION PERSONNEL, ECHELON1 PERSONNEL, ECHELON2 PERSONNEL).
	ReportGen param.Opt[string] `json:"reportGen,omitzero"`
	// Datetime of the compiling of the patients casualty report, in ISO 8601 UTC
	// format.
	ReportTime param.Opt[time.Time] `json:"reportTime,omitzero" format:"date-time"`
	// The patient branch of service (AIR FORCE, ARMY, NAVY, MARINES, CIV, CONTR,
	// UNKNOWN SERVICE).
	Service param.Opt[string] `json:"service,omitzero"`
	// Allergy information.
	Allergy []EvacNewParamsCasualtyInfoAllergy `json:"allergy,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the burial location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	BurialLocation []float64 `json:"burialLocation,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the collection point. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	CollectionPoint []float64 `json:"collectionPoint,omitzero"`
	// Health condition information.
	Condition []EvacNewParamsCasualtyInfoCondition `json:"condition,omitzero"`
	// Medical condition causation information.
	Etiology []EvacNewParamsCasualtyInfoEtiology `json:"etiology,omitzero"`
	// Health state information.
	HealthState []EvacNewParamsCasualtyInfoHealthState `json:"healthState,omitzero"`
	// Injury specifics.
	Injury []EvacNewParamsCasualtyInfoInjury `json:"injury,omitzero"`
	// Medication specifics.
	Medication []EvacNewParamsCasualtyInfoMedication `json:"medication,omitzero"`
	// Array specifying if any special equipment is need for each of the evacuation of
	// this patient (EXTRACTION EQUIPMENT, SEMI RIGID LITTER, BACKBOARD, CERVICAL
	// COLLAR ,JUNGLE PENETRATOR, OXYGEN, WHOLE BLOOD, VENTILATOR, HOIST, NONE).
	SpecMedEquip []string `json:"specMedEquip,omitzero"`
	// Treatment information.
	Treatment []EvacNewParamsCasualtyInfoTreatment `json:"treatment,omitzero"`
	// Information obtained for vital signs.
	VitalSignData []EvacNewParamsCasualtyInfoVitalSignData `json:"vitalSignData,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsCasualtyInfo) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvacNewParamsCasualtyInfo) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsCasualtyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewParamsCasualtyInfoAllergy struct {
	// Additional comments on the patient's allergy information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Type of patient allergy (e.g. PENICILLIN, SULFA, OTHER).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsCasualtyInfoAllergy) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvacNewParamsCasualtyInfoAllergy) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsCasualtyInfoAllergy
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewParamsCasualtyInfoCondition struct {
	// Body part location or body part referenced in condition. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's condition.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Datetime of the condition diagnosis in ISO 8601 UTC datetime format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Health condition assessment. Intended as, but not constrained to, K07.1
	// Condition Type Enumeration (e.g. ACTIVITY HIGH, ACTIVITY LOW, ACTIVITY MEDIUM,
	// ACTIVITY NONE, AVPU ALERT, AVPU ALTERED MENTAL STATE, AVPU PAIN, AVPU
	// UNRESPONSIVE, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsCasualtyInfoCondition) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewParamsCasualtyInfoCondition) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsCasualtyInfoCondition
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewParamsCasualtyInfoEtiology struct {
	// The body part or location affected from the etiology. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's etiology information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Datetime of the discovery of the etiology state in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// The cause or manner of causation of the medical condition. Intended as, but not
	// constrained to, K07.1 EtiologyType Enumeration (e.g. ASSAULT, BUILDING COLLAPSE,
	// BURN CHEMICAL, BURN ELECTRICAL, BURN, BURN HOT LIQUID, BURN RADIATION, BURN
	// THERMAL, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsCasualtyInfoEtiology) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewParamsCasualtyInfoEtiology) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsCasualtyInfoEtiology
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewParamsCasualtyInfoHealthState struct {
	// Medical color code used to quickly identify various medical state (e.g. AMBER,
	// BLACK, BLUE, GRAY, NORMAL, RED).
	HealthStateCode param.Opt[string] `json:"healthStateCode,omitzero"`
	// Medical confidence factor.
	MedConfFactor param.Opt[int64] `json:"medConfFactor,omitzero"`
	// Datetime of the health state diagnosis in ISO 8601 UTC datetime format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Generalized state of health type (BIOLOGICAL, CHEMICAL, COGNITIVE, HYDRATION,
	// LIFE SIGN, RADIATION, SHOCK, THERMAL).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsCasualtyInfoHealthState) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewParamsCasualtyInfoHealthState) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsCasualtyInfoHealthState
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewParamsCasualtyInfoInjury struct {
	// Body part location of the injury. Intended as, but not constrained to, K07.1
	// Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE LEFT FRONT, ANKLE RIGHT
	// BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW BACK, ARM LEFT ELBOW
	// FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's injury information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// The time of the injury, in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Classification of the injury type (ABRASION, AMPUTATION IT, AVULATION,
	// BALLISTIC, BLAST WAVE, BURN 1ST DEGREE, BURN 2ND DEGREE, BURN 3RD DEGREE, BURN
	// INHALATION, BURN LOWER AIRWAY, CHEST FLAIL, CHEST OPEN, DEGLOVING, ECCHYMOSIS,
	// FRACTURE CLOSED, FRACTURE CREPITUS, FRACTURE IT, FRACTURE OPEN, HEMATOMA,
	// IRREGULAR CONSISTENCY, IRREGULAR CONSISTENCY RIDGED, IRREGULAR CONSISTENCY
	// SWOLLEN, IRREGULAR CONSISTENCY SWOLLEN DISTENDED, IRREGULAR CONSISTENCY TENDER,
	// IRREGULAR POSITION, IRREGULAR SHAPE, IRREGULAR SHAPE MISSHAPED, IRREGULAR SHAPE
	// NON SYMMETRICAL, LACERATION, NEUROVASCULAR COMPROMISE, NEUROVASCULAR INTACT,
	// PUNCTURE, SEAT BELT SIGN, STAB, TIC TIM).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsCasualtyInfoInjury) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvacNewParamsCasualtyInfoInjury) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsCasualtyInfoInjury
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewParamsCasualtyInfoMedication struct {
	// Route of medication delivery (e.g. INJECTION, ORAL, etc.).
	AdminRoute param.Opt[string] `json:"adminRoute,omitzero"`
	// Body part location or body part referenced for medication. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's medication information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Quantity of medicine or drug administered or recommended to be taken at a
	// particular time.
	Dose param.Opt[string] `json:"dose,omitzero"`
	// The time that the medication was administered in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// The type of medication administered. Intended as, but not constrained to, K07.1
	// Medication Enumeration (CEFOTETAN, ABRASION, ABX, AMOXILOXACIN, ANALGESIC,
	// COLLOID, CRYOPECIPITATES, CRYSTALLOID, EPINEPHRINE, ERTAPENEM, FENTANYL,
	// HEXTEND, LACTATED RINGERS, MOBIC, MORPHINE, NARCOTIC, NS, PENICILLIN, PLASMA,
	// PLATELETS, PRBC, TYLENOL, WHOLE BLOOD MT).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsCasualtyInfoMedication) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewParamsCasualtyInfoMedication) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsCasualtyInfoMedication
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewParamsCasualtyInfoTreatment struct {
	// Body part location or body part treated or to be treated. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's treatment information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Datetime of the treatment in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Type of treatment administered or to be administered. Intended as, but not
	// constrained to, K07.1 Treatment Type Enumeration (e.g. AIRWAY ADJUNCT, AIRWAY
	// ASSISTED VENTILATION, AIRWAY COMBI TUBE USED, AIRWAY ET NT, AIRWAY INTUBATED,
	// AIRWAY NPA OPA APPLIED, AIRWAY PATIENT, AIRWAY POSITIONAL, AIRWAY SURGICAL CRIC,
	// BREATHING CHEST SEAL, BREATHING CHEST TUBE, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsCasualtyInfoTreatment) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewParamsCasualtyInfoTreatment) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsCasualtyInfoTreatment
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewParamsCasualtyInfoVitalSignData struct {
	// Medical confidence factor.
	MedConfFactor param.Opt[int64] `json:"medConfFactor,omitzero"`
	// Datetime of the vital sign measurement in ISO 8601 UTC datetime format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Patient vital sign measured (e.g. HEART RATE, PULSE RATE, RESPIRATION RATE,
	// TEMPERATURE CORE, etc.).
	VitalSign param.Opt[string] `json:"vitalSign,omitzero"`
	// Vital sign value 1. The content of this field is dependent on the type of vital
	// sign being measured (see the vitalSign field).
	VitalSign1 param.Opt[float64] `json:"vitalSign1,omitzero"`
	// Vital sign value 2. The content of this field is dependent on the type of vital
	// sign being measured (see the vitalSign field).
	VitalSign2 param.Opt[float64] `json:"vitalSign2,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsCasualtyInfoVitalSignData) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewParamsCasualtyInfoVitalSignData) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsCasualtyInfoVitalSignData
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewParamsEnemyData struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewParamsEnemyData) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvacNewParamsEnemyData) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewParamsEnemyData
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [EvacGetParams]'s query parameters as `url.Values`.
func (r EvacGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EvacListParams struct {
	// The request time, in ISO 8601 UTC format. (YYYY-MM-DDTHH:MM:SS.sssZ)
	ReqTime     time.Time        `query:"reqTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [EvacListParams]'s query parameters as `url.Values`.
func (r EvacListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EvacCountParams struct {
	// The request time, in ISO 8601 UTC format. (YYYY-MM-DDTHH:MM:SS.sssZ)
	ReqTime     time.Time        `query:"reqTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [EvacCountParams]'s query parameters as `url.Values`.
func (r EvacCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EvacNewBulkParams struct {
	Body []EvacNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r EvacNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Casualty report and evacuation request. Used to report and request support to
// evacuate friendly and enemy casualties.
//
// The properties ClassificationMarking, DataMode, PickupLat, PickupLon, ReqTime,
// Source, Type are required.
type EvacNewBulkParamsBody struct {
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
	// WGS-84 latitude of the pickup location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PickupLat float64 `json:"pickupLat,required"`
	// WGS-84 longitude of the pickup location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PickupLon float64 `json:"pickupLon,required"`
	// The request time, in ISO 8601 UTC format.
	ReqTime time.Time `json:"reqTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of this medevac record (REQUEST, RESPONSE).
	//
	// Any of "REQUEST", "RESPONSE".
	Type string `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Radius of circular area about lat/lon point, in meters (1-sigma, if representing
	// error).
	Ce param.Opt[float64] `json:"ce,omitzero"`
	// The contact frequency, in Hz, of the agency or zone controller.
	CntctFreq param.Opt[float64] `json:"cntctFreq,omitzero"`
	// Additional comments for the medevac mission.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of a weather report associated with this evacuation.
	IDWeatherReport param.Opt[string] `json:"idWeatherReport,omitzero"`
	// Height above lat/lon point, in meters (1-sigma, if representing linear error).
	Le param.Opt[float64] `json:"le,omitzero"`
	// UUID identifying the medevac mission, which should remain the same on subsequent
	// posts related to the same medevac mission.
	MedevacID param.Opt[string] `json:"medevacId,omitzero"`
	// Flag indicating whether the mission requires medical personnel.
	MedicReq param.Opt[bool] `json:"medicReq,omitzero"`
	// The operation type of the evacuation. (NOT SPECIFIED, AIR, GROUND, SURFACE).
	MissionType param.Opt[string] `json:"missionType,omitzero"`
	// Number of ambulatory personnel requiring evacuation.
	NumAmbulatory param.Opt[int64] `json:"numAmbulatory,omitzero"`
	// The count of people requiring medevac.
	NumCasualties param.Opt[int64] `json:"numCasualties,omitzero"`
	// Number of people Killed In Action.
	NumKia param.Opt[int64] `json:"numKIA,omitzero"`
	// Number of littered personnel requiring evacuation.
	NumLitter param.Opt[int64] `json:"numLitter,omitzero"`
	// Number of people Wounded In Action.
	NumWia param.Opt[int64] `json:"numWIA,omitzero"`
	// Amplifying data for the terrain describing important obstacles in or around the
	// zone.
	ObstaclesRemarks param.Opt[string] `json:"obstaclesRemarks,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Altitude relative to WGS-84 ellipsoid, in meters. Positive values indicate a
	// point height above ellipsoid, and negative values indicate a point height below
	// ellipsoid.
	PickupAlt param.Opt[float64] `json:"pickupAlt,omitzero"`
	// The expected pickup time, in ISO 8601 UTC format.
	PickupTime param.Opt[time.Time] `json:"pickupTime,omitzero" format:"date-time"`
	// The call sign of this medevac requestor.
	ReqCallSign param.Opt[string] `json:"reqCallSign,omitzero"`
	// Externally provided Medevac request number (e.g. MED.1.223908).
	ReqNum param.Opt[string] `json:"reqNum,omitzero"`
	// Short description of the terrain features of the pickup location (WOODS, TREES,
	// PLOWED FIELDS, FLAT, STANDING WATER, MARSH, URBAN BUILT-UP AREA, MOUNTAIN, HILL,
	// SAND TD, ROCKY, VALLEY, METAMORPHIC ICE, UNKNOWN TD, SEA, NO STATEMENT).
	Terrain param.Opt[string] `json:"terrain,omitzero"`
	// Amplifying data for the terrain describing any notable additional terrain
	// features.
	TerrainRemarks param.Opt[string] `json:"terrainRemarks,omitzero"`
	// The call sign of the zone controller.
	ZoneContrCallSign param.Opt[string] `json:"zoneContrCallSign,omitzero"`
	// Flag indicating that the pickup site is hot and hostiles are in the area.
	ZoneHot param.Opt[bool] `json:"zoneHot,omitzero"`
	// The expected marker identifying the pickup site (SMOKE ZONE MARKING, FLARES,
	// MIRROR, GLIDE ANGLE INDICATOR LIGHT, LIGHT ZONE MARKING, PANELS, FIRE, LASER
	// DESIGNATOR, STROBE LIGHTS, VEHICLE LIGHTS, COLORED SMOKE, WHITE PHOSPHERUS,
	// INFRARED, ILLUMINATION, FRATRICIDE FENCE).
	ZoneMarking param.Opt[string] `json:"zoneMarking,omitzero"`
	// Color used for the pickup site marking (RED, WHITE, BLUE, YELLOW, GREEN, ORANGE,
	// BLACK, PURPLE, BROWN, TAN, GRAY, SILVER, CAMOUFLAGE, OTHER COLOR).
	ZoneMarkingColor param.Opt[string] `json:"zoneMarkingColor,omitzero"`
	// The name of the zone.
	ZoneName param.Opt[string] `json:"zoneName,omitzero"`
	// The pickup site security (UNKNOWN ZONESECURITY, NO ENEMY, POSSIBLE ENEMY, ENEMY
	// IN AREA USE CAUTION, ENEMY IN AREA ARMED ESCORT REQUIRED).
	ZoneSecurity param.Opt[string] `json:"zoneSecurity,omitzero"`
	// Identity and medical information on the patient to be evacuated.
	CasualtyInfo []EvacNewBulkParamsBodyCasualtyInfo `json:"casualtyInfo,omitzero"`
	// Data defining any enemy intelligence reported by the requestor.
	EnemyData []EvacNewBulkParamsBodyEnemyData `json:"enemyData,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvacNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EvacNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[EvacNewBulkParamsBody](
		"Type", false, "REQUEST", "RESPONSE",
	)
}

type EvacNewBulkParamsBodyCasualtyInfo struct {
	// The patient age, in years.
	Age param.Opt[int64] `json:"age,omitzero"`
	// The patient blood type (A POS, B POS, AB POS, O POS, A NEG, B NEG, AB NEG, O
	// NEG).
	BloodType param.Opt[string] `json:"bloodType,omitzero"`
	// The body part involved for the patient (HEAD, NECK, ABDOMEN, UPPER EXTREMITIES,
	// BACK, FACE, LOWER EXTREMITIES, FRONT, OBSTETRICAL GYNECOLOGICAL, OTHER BODY
	// PART).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// The call sign of this patient.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Unique identifier for the patient care provider.
	CareProviderUrn param.Opt[string] `json:"careProviderUrn,omitzero"`
	// Optional casualty key.
	CasualtyKey param.Opt[string] `json:"casualtyKey,omitzero"`
	// The type of medical issue resulting in the need to evacuate the patient (NON
	// BATTLE, CUT, BURN, SICK, FRACTURE, AMPUTATION, PERFORATION, NUCLEAR, EXHAUSTION,
	// BIOLOGICAL, CHEMICAL, SHOCK, PUNCTURE WOUND, OTHER CUT, WOUNDED IN ACTION,
	// DENIAL, COMBAT STRESS).
	CasualtyType param.Opt[string] `json:"casualtyType,omitzero"`
	// Additional comments on the patient's casualty information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// The contamination specified for the patient (NONE, RADIATION, BIOLOGICAL,
	// CHEMICAL).
	ContamType param.Opt[string] `json:"contamType,omitzero"`
	// The patient's general medical state (SICK IN QUARTERS, RETURN TO DUTY, EVACUATE
	// WOUNDED, EVACUATE DECEASED, INTERRED).
	Disposition param.Opt[string] `json:"disposition,omitzero"`
	// The expected disposition of this patient (R T D, EVACUATE, EVACUATE TO FORWARD
	// SURGICAL TEAM, EVACUATE TO COMBAT SUPPORT HOSPITAL, EVACUATE TO AERO MEDICAL
	// STAGING FACILITY, EVACUATE TO SUSTAINING BASE MEDICAL TREATMENT FACILITY).
	DispositionType param.Opt[string] `json:"dispositionType,omitzero"`
	// The required evacuation method for this patient (AIR, GROUND, NOT EVACUATED).
	EvacType param.Opt[string] `json:"evacType,omitzero"`
	// The patient sex (MALE, FEMALE).
	Gender param.Opt[string] `json:"gender,omitzero"`
	// Last 4 characters of the patient social security code, or equivalent.
	Last4Ssn param.Opt[string] `json:"last4SSN,omitzero"`
	// The patient common or legal name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The country code indicating the citizenship of the patient.
	Nationality param.Opt[string] `json:"nationality,omitzero"`
	// The career field of this patient.
	OccSpeciality param.Opt[string] `json:"occSpeciality,omitzero"`
	// The patient service identity (UNKNOWN MILITARY, UNKNOWN CIVILIAN, FRIEND
	// MILITARY, FRIEND CIVILIAN, NEUTRAL MILITARY, NEUTRAL CIVILIAN, HOSTILE MILITARY,
	// HOSTILE CIVILIAN).
	PatientIdentity param.Opt[string] `json:"patientIdentity,omitzero"`
	// The patient service status (US MILITARY, US CIVILIAN, NON US MILITARY, NON US
	// CIVILIAN, ENEMY POW).
	PatientStatus param.Opt[string] `json:"patientStatus,omitzero"`
	// The patient pay grade or rank designation (O-10, O-9, O-8, O-7, O-6, O-5, O-4,
	// O-3, O-2, O-1, CWO-5, CWO-4, CWO-2, CWO-1, E -9, E-8, E-7, E-6, E-5, E-4, E-3,
	// E-2, E-1, NONE, CIVILIAN).
	PayGrade param.Opt[string] `json:"payGrade,omitzero"`
	// The priority of the medevac mission for this patient (URGENT, PRIORITY, ROUTINE,
	// URGENT SURGERY, CONVENIENCE).
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The method used to generate this medevac report (DEVICE, GROUND COMBAT
	// PERSONNEL, EVACUATION PERSONNEL, ECHELON1 PERSONNEL, ECHELON2 PERSONNEL).
	ReportGen param.Opt[string] `json:"reportGen,omitzero"`
	// Datetime of the compiling of the patients casualty report, in ISO 8601 UTC
	// format.
	ReportTime param.Opt[time.Time] `json:"reportTime,omitzero" format:"date-time"`
	// The patient branch of service (AIR FORCE, ARMY, NAVY, MARINES, CIV, CONTR,
	// UNKNOWN SERVICE).
	Service param.Opt[string] `json:"service,omitzero"`
	// Allergy information.
	Allergy []EvacNewBulkParamsBodyCasualtyInfoAllergy `json:"allergy,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the burial location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	BurialLocation []float64 `json:"burialLocation,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the collection point. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	CollectionPoint []float64 `json:"collectionPoint,omitzero"`
	// Health condition information.
	Condition []EvacNewBulkParamsBodyCasualtyInfoCondition `json:"condition,omitzero"`
	// Medical condition causation information.
	Etiology []EvacNewBulkParamsBodyCasualtyInfoEtiology `json:"etiology,omitzero"`
	// Health state information.
	HealthState []EvacNewBulkParamsBodyCasualtyInfoHealthState `json:"healthState,omitzero"`
	// Injury specifics.
	Injury []EvacNewBulkParamsBodyCasualtyInfoInjury `json:"injury,omitzero"`
	// Medication specifics.
	Medication []EvacNewBulkParamsBodyCasualtyInfoMedication `json:"medication,omitzero"`
	// Array specifying if any special equipment is need for each of the evacuation of
	// this patient (EXTRACTION EQUIPMENT, SEMI RIGID LITTER, BACKBOARD, CERVICAL
	// COLLAR ,JUNGLE PENETRATOR, OXYGEN, WHOLE BLOOD, VENTILATOR, HOIST, NONE).
	SpecMedEquip []string `json:"specMedEquip,omitzero"`
	// Treatment information.
	Treatment []EvacNewBulkParamsBodyCasualtyInfoTreatment `json:"treatment,omitzero"`
	// Information obtained for vital signs.
	VitalSignData []EvacNewBulkParamsBodyCasualtyInfoVitalSignData `json:"vitalSignData,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyCasualtyInfo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewBulkParamsBodyCasualtyInfo) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyCasualtyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewBulkParamsBodyCasualtyInfoAllergy struct {
	// Additional comments on the patient's allergy information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Type of patient allergy (e.g. PENICILLIN, SULFA, OTHER).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyCasualtyInfoAllergy) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewBulkParamsBodyCasualtyInfoAllergy) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyCasualtyInfoAllergy
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewBulkParamsBodyCasualtyInfoCondition struct {
	// Body part location or body part referenced in condition. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's condition.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Datetime of the condition diagnosis in ISO 8601 UTC datetime format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Health condition assessment. Intended as, but not constrained to, K07.1
	// Condition Type Enumeration (e.g. ACTIVITY HIGH, ACTIVITY LOW, ACTIVITY MEDIUM,
	// ACTIVITY NONE, AVPU ALERT, AVPU ALTERED MENTAL STATE, AVPU PAIN, AVPU
	// UNRESPONSIVE, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyCasualtyInfoCondition) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewBulkParamsBodyCasualtyInfoCondition) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyCasualtyInfoCondition
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewBulkParamsBodyCasualtyInfoEtiology struct {
	// The body part or location affected from the etiology. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's etiology information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Datetime of the discovery of the etiology state in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// The cause or manner of causation of the medical condition. Intended as, but not
	// constrained to, K07.1 EtiologyType Enumeration (e.g. ASSAULT, BUILDING COLLAPSE,
	// BURN CHEMICAL, BURN ELECTRICAL, BURN, BURN HOT LIQUID, BURN RADIATION, BURN
	// THERMAL, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyCasualtyInfoEtiology) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewBulkParamsBodyCasualtyInfoEtiology) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyCasualtyInfoEtiology
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewBulkParamsBodyCasualtyInfoHealthState struct {
	// Medical color code used to quickly identify various medical state (e.g. AMBER,
	// BLACK, BLUE, GRAY, NORMAL, RED).
	HealthStateCode param.Opt[string] `json:"healthStateCode,omitzero"`
	// Medical confidence factor.
	MedConfFactor param.Opt[int64] `json:"medConfFactor,omitzero"`
	// Datetime of the health state diagnosis in ISO 8601 UTC datetime format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Generalized state of health type (BIOLOGICAL, CHEMICAL, COGNITIVE, HYDRATION,
	// LIFE SIGN, RADIATION, SHOCK, THERMAL).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyCasualtyInfoHealthState) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewBulkParamsBodyCasualtyInfoHealthState) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyCasualtyInfoHealthState
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewBulkParamsBodyCasualtyInfoInjury struct {
	// Body part location of the injury. Intended as, but not constrained to, K07.1
	// Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE LEFT FRONT, ANKLE RIGHT
	// BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW BACK, ARM LEFT ELBOW
	// FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's injury information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// The time of the injury, in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Classification of the injury type (ABRASION, AMPUTATION IT, AVULATION,
	// BALLISTIC, BLAST WAVE, BURN 1ST DEGREE, BURN 2ND DEGREE, BURN 3RD DEGREE, BURN
	// INHALATION, BURN LOWER AIRWAY, CHEST FLAIL, CHEST OPEN, DEGLOVING, ECCHYMOSIS,
	// FRACTURE CLOSED, FRACTURE CREPITUS, FRACTURE IT, FRACTURE OPEN, HEMATOMA,
	// IRREGULAR CONSISTENCY, IRREGULAR CONSISTENCY RIDGED, IRREGULAR CONSISTENCY
	// SWOLLEN, IRREGULAR CONSISTENCY SWOLLEN DISTENDED, IRREGULAR CONSISTENCY TENDER,
	// IRREGULAR POSITION, IRREGULAR SHAPE, IRREGULAR SHAPE MISSHAPED, IRREGULAR SHAPE
	// NON SYMMETRICAL, LACERATION, NEUROVASCULAR COMPROMISE, NEUROVASCULAR INTACT,
	// PUNCTURE, SEAT BELT SIGN, STAB, TIC TIM).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyCasualtyInfoInjury) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewBulkParamsBodyCasualtyInfoInjury) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyCasualtyInfoInjury
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewBulkParamsBodyCasualtyInfoMedication struct {
	// Route of medication delivery (e.g. INJECTION, ORAL, etc.).
	AdminRoute param.Opt[string] `json:"adminRoute,omitzero"`
	// Body part location or body part referenced for medication. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's medication information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Quantity of medicine or drug administered or recommended to be taken at a
	// particular time.
	Dose param.Opt[string] `json:"dose,omitzero"`
	// The time that the medication was administered in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// The type of medication administered. Intended as, but not constrained to, K07.1
	// Medication Enumeration (CEFOTETAN, ABRASION, ABX, AMOXILOXACIN, ANALGESIC,
	// COLLOID, CRYOPECIPITATES, CRYSTALLOID, EPINEPHRINE, ERTAPENEM, FENTANYL,
	// HEXTEND, LACTATED RINGERS, MOBIC, MORPHINE, NARCOTIC, NS, PENICILLIN, PLASMA,
	// PLATELETS, PRBC, TYLENOL, WHOLE BLOOD MT).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyCasualtyInfoMedication) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewBulkParamsBodyCasualtyInfoMedication) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyCasualtyInfoMedication
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewBulkParamsBodyCasualtyInfoTreatment struct {
	// Body part location or body part treated or to be treated. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's treatment information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Datetime of the treatment in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Type of treatment administered or to be administered. Intended as, but not
	// constrained to, K07.1 Treatment Type Enumeration (e.g. AIRWAY ADJUNCT, AIRWAY
	// ASSISTED VENTILATION, AIRWAY COMBI TUBE USED, AIRWAY ET NT, AIRWAY INTUBATED,
	// AIRWAY NPA OPA APPLIED, AIRWAY PATIENT, AIRWAY POSITIONAL, AIRWAY SURGICAL CRIC,
	// BREATHING CHEST SEAL, BREATHING CHEST TUBE, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyCasualtyInfoTreatment) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewBulkParamsBodyCasualtyInfoTreatment) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyCasualtyInfoTreatment
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewBulkParamsBodyCasualtyInfoVitalSignData struct {
	// Medical confidence factor.
	MedConfFactor param.Opt[int64] `json:"medConfFactor,omitzero"`
	// Datetime of the vital sign measurement in ISO 8601 UTC datetime format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Patient vital sign measured (e.g. HEART RATE, PULSE RATE, RESPIRATION RATE,
	// TEMPERATURE CORE, etc.).
	VitalSign param.Opt[string] `json:"vitalSign,omitzero"`
	// Vital sign value 1. The content of this field is dependent on the type of vital
	// sign being measured (see the vitalSign field).
	VitalSign1 param.Opt[float64] `json:"vitalSign1,omitzero"`
	// Vital sign value 2. The content of this field is dependent on the type of vital
	// sign being measured (see the vitalSign field).
	VitalSign2 param.Opt[float64] `json:"vitalSign2,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyCasualtyInfoVitalSignData) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacNewBulkParamsBodyCasualtyInfoVitalSignData) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyCasualtyInfoVitalSignData
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacNewBulkParamsBodyEnemyData struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacNewBulkParamsBodyEnemyData) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvacNewBulkParamsBodyEnemyData) MarshalJSON() (data []byte, err error) {
	type shadow EvacNewBulkParamsBodyEnemyData
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParams struct {
	Body []EvacUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r EvacUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Casualty report and evacuation request. Used to report and request support to
// evacuate friendly and enemy casualties.
//
// The properties ClassificationMarking, DataMode, PickupLat, PickupLon, ReqTime,
// Source, Type are required.
type EvacUnvalidatedPublishParamsBody struct {
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
	// WGS-84 latitude of the pickup location, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	PickupLat float64 `json:"pickupLat,required"`
	// WGS-84 longitude of the pickup location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	PickupLon float64 `json:"pickupLon,required"`
	// The request time, in ISO 8601 UTC format.
	ReqTime time.Time `json:"reqTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of this medevac record (REQUEST, RESPONSE).
	//
	// Any of "REQUEST", "RESPONSE".
	Type string `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Radius of circular area about lat/lon point, in meters (1-sigma, if representing
	// error).
	Ce param.Opt[float64] `json:"ce,omitzero"`
	// The contact frequency, in Hz, of the agency or zone controller.
	CntctFreq param.Opt[float64] `json:"cntctFreq,omitzero"`
	// Additional comments for the medevac mission.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of a weather report associated with this evacuation.
	IDWeatherReport param.Opt[string] `json:"idWeatherReport,omitzero"`
	// Height above lat/lon point, in meters (1-sigma, if representing linear error).
	Le param.Opt[float64] `json:"le,omitzero"`
	// UUID identifying the medevac mission, which should remain the same on subsequent
	// posts related to the same medevac mission.
	MedevacID param.Opt[string] `json:"medevacId,omitzero"`
	// Flag indicating whether the mission requires medical personnel.
	MedicReq param.Opt[bool] `json:"medicReq,omitzero"`
	// The operation type of the evacuation. (NOT SPECIFIED, AIR, GROUND, SURFACE).
	MissionType param.Opt[string] `json:"missionType,omitzero"`
	// Number of ambulatory personnel requiring evacuation.
	NumAmbulatory param.Opt[int64] `json:"numAmbulatory,omitzero"`
	// The count of people requiring medevac.
	NumCasualties param.Opt[int64] `json:"numCasualties,omitzero"`
	// Number of people Killed In Action.
	NumKia param.Opt[int64] `json:"numKIA,omitzero"`
	// Number of littered personnel requiring evacuation.
	NumLitter param.Opt[int64] `json:"numLitter,omitzero"`
	// Number of people Wounded In Action.
	NumWia param.Opt[int64] `json:"numWIA,omitzero"`
	// Amplifying data for the terrain describing important obstacles in or around the
	// zone.
	ObstaclesRemarks param.Opt[string] `json:"obstaclesRemarks,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Altitude relative to WGS-84 ellipsoid, in meters. Positive values indicate a
	// point height above ellipsoid, and negative values indicate a point height below
	// ellipsoid.
	PickupAlt param.Opt[float64] `json:"pickupAlt,omitzero"`
	// The expected pickup time, in ISO 8601 UTC format.
	PickupTime param.Opt[time.Time] `json:"pickupTime,omitzero" format:"date-time"`
	// The call sign of this medevac requestor.
	ReqCallSign param.Opt[string] `json:"reqCallSign,omitzero"`
	// Externally provided Medevac request number (e.g. MED.1.223908).
	ReqNum param.Opt[string] `json:"reqNum,omitzero"`
	// Short description of the terrain features of the pickup location (WOODS, TREES,
	// PLOWED FIELDS, FLAT, STANDING WATER, MARSH, URBAN BUILT-UP AREA, MOUNTAIN, HILL,
	// SAND TD, ROCKY, VALLEY, METAMORPHIC ICE, UNKNOWN TD, SEA, NO STATEMENT).
	Terrain param.Opt[string] `json:"terrain,omitzero"`
	// Amplifying data for the terrain describing any notable additional terrain
	// features.
	TerrainRemarks param.Opt[string] `json:"terrainRemarks,omitzero"`
	// The call sign of the zone controller.
	ZoneContrCallSign param.Opt[string] `json:"zoneContrCallSign,omitzero"`
	// Flag indicating that the pickup site is hot and hostiles are in the area.
	ZoneHot param.Opt[bool] `json:"zoneHot,omitzero"`
	// The expected marker identifying the pickup site (SMOKE ZONE MARKING, FLARES,
	// MIRROR, GLIDE ANGLE INDICATOR LIGHT, LIGHT ZONE MARKING, PANELS, FIRE, LASER
	// DESIGNATOR, STROBE LIGHTS, VEHICLE LIGHTS, COLORED SMOKE, WHITE PHOSPHERUS,
	// INFRARED, ILLUMINATION, FRATRICIDE FENCE).
	ZoneMarking param.Opt[string] `json:"zoneMarking,omitzero"`
	// Color used for the pickup site marking (RED, WHITE, BLUE, YELLOW, GREEN, ORANGE,
	// BLACK, PURPLE, BROWN, TAN, GRAY, SILVER, CAMOUFLAGE, OTHER COLOR).
	ZoneMarkingColor param.Opt[string] `json:"zoneMarkingColor,omitzero"`
	// The name of the zone.
	ZoneName param.Opt[string] `json:"zoneName,omitzero"`
	// The pickup site security (UNKNOWN ZONESECURITY, NO ENEMY, POSSIBLE ENEMY, ENEMY
	// IN AREA USE CAUTION, ENEMY IN AREA ARMED ESCORT REQUIRED).
	ZoneSecurity param.Opt[string] `json:"zoneSecurity,omitzero"`
	// Identity and medical information on the patient to be evacuated.
	CasualtyInfo []EvacUnvalidatedPublishParamsBodyCasualtyInfo `json:"casualtyInfo,omitzero"`
	// Data defining any enemy intelligence reported by the requestor.
	EnemyData []EvacUnvalidatedPublishParamsBodyEnemyData `json:"enemyData,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvacUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EvacUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[EvacUnvalidatedPublishParamsBody](
		"Type", false, "REQUEST", "RESPONSE",
	)
}

type EvacUnvalidatedPublishParamsBodyCasualtyInfo struct {
	// The patient age, in years.
	Age param.Opt[int64] `json:"age,omitzero"`
	// The patient blood type (A POS, B POS, AB POS, O POS, A NEG, B NEG, AB NEG, O
	// NEG).
	BloodType param.Opt[string] `json:"bloodType,omitzero"`
	// The body part involved for the patient (HEAD, NECK, ABDOMEN, UPPER EXTREMITIES,
	// BACK, FACE, LOWER EXTREMITIES, FRONT, OBSTETRICAL GYNECOLOGICAL, OTHER BODY
	// PART).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// The call sign of this patient.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Unique identifier for the patient care provider.
	CareProviderUrn param.Opt[string] `json:"careProviderUrn,omitzero"`
	// Optional casualty key.
	CasualtyKey param.Opt[string] `json:"casualtyKey,omitzero"`
	// The type of medical issue resulting in the need to evacuate the patient (NON
	// BATTLE, CUT, BURN, SICK, FRACTURE, AMPUTATION, PERFORATION, NUCLEAR, EXHAUSTION,
	// BIOLOGICAL, CHEMICAL, SHOCK, PUNCTURE WOUND, OTHER CUT, WOUNDED IN ACTION,
	// DENIAL, COMBAT STRESS).
	CasualtyType param.Opt[string] `json:"casualtyType,omitzero"`
	// Additional comments on the patient's casualty information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// The contamination specified for the patient (NONE, RADIATION, BIOLOGICAL,
	// CHEMICAL).
	ContamType param.Opt[string] `json:"contamType,omitzero"`
	// The patient's general medical state (SICK IN QUARTERS, RETURN TO DUTY, EVACUATE
	// WOUNDED, EVACUATE DECEASED, INTERRED).
	Disposition param.Opt[string] `json:"disposition,omitzero"`
	// The expected disposition of this patient (R T D, EVACUATE, EVACUATE TO FORWARD
	// SURGICAL TEAM, EVACUATE TO COMBAT SUPPORT HOSPITAL, EVACUATE TO AERO MEDICAL
	// STAGING FACILITY, EVACUATE TO SUSTAINING BASE MEDICAL TREATMENT FACILITY).
	DispositionType param.Opt[string] `json:"dispositionType,omitzero"`
	// The required evacuation method for this patient (AIR, GROUND, NOT EVACUATED).
	EvacType param.Opt[string] `json:"evacType,omitzero"`
	// The patient sex (MALE, FEMALE).
	Gender param.Opt[string] `json:"gender,omitzero"`
	// Last 4 characters of the patient social security code, or equivalent.
	Last4Ssn param.Opt[string] `json:"last4SSN,omitzero"`
	// The patient common or legal name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The country code indicating the citizenship of the patient.
	Nationality param.Opt[string] `json:"nationality,omitzero"`
	// The career field of this patient.
	OccSpeciality param.Opt[string] `json:"occSpeciality,omitzero"`
	// The patient service identity (UNKNOWN MILITARY, UNKNOWN CIVILIAN, FRIEND
	// MILITARY, FRIEND CIVILIAN, NEUTRAL MILITARY, NEUTRAL CIVILIAN, HOSTILE MILITARY,
	// HOSTILE CIVILIAN).
	PatientIdentity param.Opt[string] `json:"patientIdentity,omitzero"`
	// The patient service status (US MILITARY, US CIVILIAN, NON US MILITARY, NON US
	// CIVILIAN, ENEMY POW).
	PatientStatus param.Opt[string] `json:"patientStatus,omitzero"`
	// The patient pay grade or rank designation (O-10, O-9, O-8, O-7, O-6, O-5, O-4,
	// O-3, O-2, O-1, CWO-5, CWO-4, CWO-2, CWO-1, E -9, E-8, E-7, E-6, E-5, E-4, E-3,
	// E-2, E-1, NONE, CIVILIAN).
	PayGrade param.Opt[string] `json:"payGrade,omitzero"`
	// The priority of the medevac mission for this patient (URGENT, PRIORITY, ROUTINE,
	// URGENT SURGERY, CONVENIENCE).
	Priority param.Opt[string] `json:"priority,omitzero"`
	// The method used to generate this medevac report (DEVICE, GROUND COMBAT
	// PERSONNEL, EVACUATION PERSONNEL, ECHELON1 PERSONNEL, ECHELON2 PERSONNEL).
	ReportGen param.Opt[string] `json:"reportGen,omitzero"`
	// Datetime of the compiling of the patients casualty report, in ISO 8601 UTC
	// format.
	ReportTime param.Opt[time.Time] `json:"reportTime,omitzero" format:"date-time"`
	// The patient branch of service (AIR FORCE, ARMY, NAVY, MARINES, CIV, CONTR,
	// UNKNOWN SERVICE).
	Service param.Opt[string] `json:"service,omitzero"`
	// Allergy information.
	Allergy []EvacUnvalidatedPublishParamsBodyCasualtyInfoAllergy `json:"allergy,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the burial location. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	BurialLocation []float64 `json:"burialLocation,omitzero"`
	// Array of the WGS-84 latitude (-90 to 90, negative values south of the equator)
	// in degrees, longitude (-180 to 180, negative values west of Prime Meridian) in
	// degrees, and altitude, in meters, of the collection point. This array must
	// contain a minimum of 2 elements (latitude and longitude), and may contain an
	// optional 3rd element (altitude).
	CollectionPoint []float64 `json:"collectionPoint,omitzero"`
	// Health condition information.
	Condition []EvacUnvalidatedPublishParamsBodyCasualtyInfoCondition `json:"condition,omitzero"`
	// Medical condition causation information.
	Etiology []EvacUnvalidatedPublishParamsBodyCasualtyInfoEtiology `json:"etiology,omitzero"`
	// Health state information.
	HealthState []EvacUnvalidatedPublishParamsBodyCasualtyInfoHealthState `json:"healthState,omitzero"`
	// Injury specifics.
	Injury []EvacUnvalidatedPublishParamsBodyCasualtyInfoInjury `json:"injury,omitzero"`
	// Medication specifics.
	Medication []EvacUnvalidatedPublishParamsBodyCasualtyInfoMedication `json:"medication,omitzero"`
	// Array specifying if any special equipment is need for each of the evacuation of
	// this patient (EXTRACTION EQUIPMENT, SEMI RIGID LITTER, BACKBOARD, CERVICAL
	// COLLAR ,JUNGLE PENETRATOR, OXYGEN, WHOLE BLOOD, VENTILATOR, HOIST, NONE).
	SpecMedEquip []string `json:"specMedEquip,omitzero"`
	// Treatment information.
	Treatment []EvacUnvalidatedPublishParamsBodyCasualtyInfoTreatment `json:"treatment,omitzero"`
	// Information obtained for vital signs.
	VitalSignData []EvacUnvalidatedPublishParamsBodyCasualtyInfoVitalSignData `json:"vitalSignData,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyCasualtyInfo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyCasualtyInfo) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyCasualtyInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParamsBodyCasualtyInfoAllergy struct {
	// Additional comments on the patient's allergy information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Type of patient allergy (e.g. PENICILLIN, SULFA, OTHER).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyCasualtyInfoAllergy) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyCasualtyInfoAllergy) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyCasualtyInfoAllergy
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParamsBodyCasualtyInfoCondition struct {
	// Body part location or body part referenced in condition. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's condition.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Datetime of the condition diagnosis in ISO 8601 UTC datetime format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Health condition assessment. Intended as, but not constrained to, K07.1
	// Condition Type Enumeration (e.g. ACTIVITY HIGH, ACTIVITY LOW, ACTIVITY MEDIUM,
	// ACTIVITY NONE, AVPU ALERT, AVPU ALTERED MENTAL STATE, AVPU PAIN, AVPU
	// UNRESPONSIVE, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyCasualtyInfoCondition) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyCasualtyInfoCondition) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyCasualtyInfoCondition
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParamsBodyCasualtyInfoEtiology struct {
	// The body part or location affected from the etiology. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's etiology information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Datetime of the discovery of the etiology state in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// The cause or manner of causation of the medical condition. Intended as, but not
	// constrained to, K07.1 EtiologyType Enumeration (e.g. ASSAULT, BUILDING COLLAPSE,
	// BURN CHEMICAL, BURN ELECTRICAL, BURN, BURN HOT LIQUID, BURN RADIATION, BURN
	// THERMAL, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyCasualtyInfoEtiology) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyCasualtyInfoEtiology) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyCasualtyInfoEtiology
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParamsBodyCasualtyInfoHealthState struct {
	// Medical color code used to quickly identify various medical state (e.g. AMBER,
	// BLACK, BLUE, GRAY, NORMAL, RED).
	HealthStateCode param.Opt[string] `json:"healthStateCode,omitzero"`
	// Medical confidence factor.
	MedConfFactor param.Opt[int64] `json:"medConfFactor,omitzero"`
	// Datetime of the health state diagnosis in ISO 8601 UTC datetime format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Generalized state of health type (BIOLOGICAL, CHEMICAL, COGNITIVE, HYDRATION,
	// LIFE SIGN, RADIATION, SHOCK, THERMAL).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyCasualtyInfoHealthState) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyCasualtyInfoHealthState) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyCasualtyInfoHealthState
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParamsBodyCasualtyInfoInjury struct {
	// Body part location of the injury. Intended as, but not constrained to, K07.1
	// Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE LEFT FRONT, ANKLE RIGHT
	// BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW BACK, ARM LEFT ELBOW
	// FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's injury information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// The time of the injury, in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Classification of the injury type (ABRASION, AMPUTATION IT, AVULATION,
	// BALLISTIC, BLAST WAVE, BURN 1ST DEGREE, BURN 2ND DEGREE, BURN 3RD DEGREE, BURN
	// INHALATION, BURN LOWER AIRWAY, CHEST FLAIL, CHEST OPEN, DEGLOVING, ECCHYMOSIS,
	// FRACTURE CLOSED, FRACTURE CREPITUS, FRACTURE IT, FRACTURE OPEN, HEMATOMA,
	// IRREGULAR CONSISTENCY, IRREGULAR CONSISTENCY RIDGED, IRREGULAR CONSISTENCY
	// SWOLLEN, IRREGULAR CONSISTENCY SWOLLEN DISTENDED, IRREGULAR CONSISTENCY TENDER,
	// IRREGULAR POSITION, IRREGULAR SHAPE, IRREGULAR SHAPE MISSHAPED, IRREGULAR SHAPE
	// NON SYMMETRICAL, LACERATION, NEUROVASCULAR COMPROMISE, NEUROVASCULAR INTACT,
	// PUNCTURE, SEAT BELT SIGN, STAB, TIC TIM).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyCasualtyInfoInjury) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyCasualtyInfoInjury) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyCasualtyInfoInjury
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParamsBodyCasualtyInfoMedication struct {
	// Route of medication delivery (e.g. INJECTION, ORAL, etc.).
	AdminRoute param.Opt[string] `json:"adminRoute,omitzero"`
	// Body part location or body part referenced for medication. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's medication information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Quantity of medicine or drug administered or recommended to be taken at a
	// particular time.
	Dose param.Opt[string] `json:"dose,omitzero"`
	// The time that the medication was administered in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// The type of medication administered. Intended as, but not constrained to, K07.1
	// Medication Enumeration (CEFOTETAN, ABRASION, ABX, AMOXILOXACIN, ANALGESIC,
	// COLLOID, CRYOPECIPITATES, CRYSTALLOID, EPINEPHRINE, ERTAPENEM, FENTANYL,
	// HEXTEND, LACTATED RINGERS, MOBIC, MORPHINE, NARCOTIC, NS, PENICILLIN, PLASMA,
	// PLATELETS, PRBC, TYLENOL, WHOLE BLOOD MT).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyCasualtyInfoMedication) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyCasualtyInfoMedication) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyCasualtyInfoMedication
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParamsBodyCasualtyInfoTreatment struct {
	// Body part location or body part treated or to be treated. Intended as, but not
	// constrained to, K07.1 Body Location Enumeration (e.g. ANKLE LEFT BACK, ANKLE
	// LEFT FRONT, ANKLE RIGHT BACK, ANKLE RIGHT FRONT, ARM LEFT BACK, ARM LEFT ELBOW
	// BACK, ARM LEFT ELBOW FRONT, ARM LEFT FRONT, ARM LEFT LOWER BACK, etc.).
	BodyPart param.Opt[string] `json:"bodyPart,omitzero"`
	// Additional comments on the patient's treatment information.
	Comments param.Opt[string] `json:"comments,omitzero"`
	// Datetime of the treatment in ISO 8601 UTC format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Type of treatment administered or to be administered. Intended as, but not
	// constrained to, K07.1 Treatment Type Enumeration (e.g. AIRWAY ADJUNCT, AIRWAY
	// ASSISTED VENTILATION, AIRWAY COMBI TUBE USED, AIRWAY ET NT, AIRWAY INTUBATED,
	// AIRWAY NPA OPA APPLIED, AIRWAY PATIENT, AIRWAY POSITIONAL, AIRWAY SURGICAL CRIC,
	// BREATHING CHEST SEAL, BREATHING CHEST TUBE, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyCasualtyInfoTreatment) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyCasualtyInfoTreatment) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyCasualtyInfoTreatment
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParamsBodyCasualtyInfoVitalSignData struct {
	// Medical confidence factor.
	MedConfFactor param.Opt[int64] `json:"medConfFactor,omitzero"`
	// Datetime of the vital sign measurement in ISO 8601 UTC datetime format.
	Time param.Opt[time.Time] `json:"time,omitzero" format:"date-time"`
	// Patient vital sign measured (e.g. HEART RATE, PULSE RATE, RESPIRATION RATE,
	// TEMPERATURE CORE, etc.).
	VitalSign param.Opt[string] `json:"vitalSign,omitzero"`
	// Vital sign value 1. The content of this field is dependent on the type of vital
	// sign being measured (see the vitalSign field).
	VitalSign1 param.Opt[float64] `json:"vitalSign1,omitzero"`
	// Vital sign value 2. The content of this field is dependent on the type of vital
	// sign being measured (see the vitalSign field).
	VitalSign2 param.Opt[float64] `json:"vitalSign2,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyCasualtyInfoVitalSignData) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyCasualtyInfoVitalSignData) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyCasualtyInfoVitalSignData
	return param.MarshalObject(r, (*shadow)(&r))
}

type EvacUnvalidatedPublishParamsBodyEnemyData struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvacUnvalidatedPublishParamsBodyEnemyData) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvacUnvalidatedPublishParamsBodyEnemyData) MarshalJSON() (data []byte, err error) {
	type shadow EvacUnvalidatedPublishParamsBodyEnemyData
	return param.MarshalObject(r, (*shadow)(&r))
}
