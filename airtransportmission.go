// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// AirTransportMissionService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirTransportMissionService] method instead.
type AirTransportMissionService struct {
	Options []option.RequestOption
	History AirTransportMissionHistoryService
}

// NewAirTransportMissionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAirTransportMissionService(opts ...option.RequestOption) (r AirTransportMissionService) {
	r = AirTransportMissionService{}
	r.Options = opts
	r.History = NewAirTransportMissionHistoryService(opts...)
	return
}

// Service operation to take a single AirTransportMission object as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AirTransportMissionService) New(ctx context.Context, body AirTransportMissionNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airtransportmission"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Air Transport Mission record by its unique ID
// passed as a path parameter.
func (r *AirTransportMissionService) Get(ctx context.Context, id string, query AirTransportMissionGetParams, opts ...option.RequestOption) (res *shared.AirTransportMissionFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airtransportmission/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single AirTransportMission record. A specific role
// is required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *AirTransportMissionService) Update(ctx context.Context, id string, body AirTransportMissionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airtransportmission/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirTransportMissionService) List(ctx context.Context, query AirTransportMissionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AirTransportMissionAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/airtransportmission"
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
func (r *AirTransportMissionService) ListAutoPaging(ctx context.Context, query AirTransportMissionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AirTransportMissionAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirTransportMissionService) Count(ctx context.Context, query AirTransportMissionCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airtransportmission/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AirTransportMissionService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *AirTransportMissionQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airtransportmission/queryhelp"
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
func (r *AirTransportMissionService) Tuple(ctx context.Context, query AirTransportMissionTupleParams, opts ...option.RequestOption) (res *[]shared.AirTransportMissionFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airtransportmission/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The information in an Air Transport Mission contains unique identification,
// description of the mission objective, aircraft and crew assignments, mission
// alias, embarkation/debarkation cargo locations, priority, and other mission
// characteristics.
type AirTransportMissionAbridged struct {
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
	DataMode AirTransportMissionAbridgedDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The Air Battle Plan used to coordinate and integrate air assets for this
	// mission.
	Abp string `json:"abp"`
	// Mission alias.
	Alias string `json:"alias"`
	// The unit the mission is allocated to.
	AllocatedUnit string `json:"allocatedUnit"`
	// Air Mobility Command (AMC) mission identifier according to Mobility Air Forces
	// (MAF) Encode/Decode procedures.
	AmcMissionID string `json:"amcMissionId"`
	// The Aircraft and Personnel Automated Clearance System (APACS) system identifier
	// used to process and approve clearance requests.
	ApacsID string `json:"apacsId"`
	// The call sign assigned to this mission according to the Air Tasking Order (ATO).
	AtoCallSign string `json:"atoCallSign"`
	// The mission number according to the Air Tasking Order (ATO).
	AtoMissionID string `json:"atoMissionId"`
	// The call sign for this mission.
	CallSign string `json:"callSign"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Flag indicating this is a close watch mission.
	Cw bool `json:"cw"`
	// Identifier of the Diplomatic Clearance Worksheet used to coordinate aircraft
	// clearance requests.
	DipWorksheetName string `json:"dipWorksheetName"`
	// The International Civil Aviation Organization (ICAO) site code of first cargo
	// pick up.
	FirstPickUp string `json:"firstPickUp"`
	// Global Decision Support System (GDSS) mission unique identifier.
	GdssMissionID string `json:"gdssMissionId"`
	// Collection of Hazardous Material information planned to be associated with this
	// Air Transport Mission.
	HazMat []AirTransportMissionAbridgedHazMat `json:"hazMat"`
	// Highest Joint Chiefs of Staff priority of this mission.
	JcsPriority string `json:"jcsPriority"`
	// The International Civil Aviation Organization (ICAO) site code of last cargo
	// drop off.
	LastDropOff string `json:"lastDropOff"`
	// Load type of this mission (e.g. CARGO, MIXED, PASSENGER).
	LoadCategoryType string `json:"loadCategoryType"`
	// Mode-1 interrogation response (mission code), indicating mission or aircraft
	// type.
	M1 string `json:"m1"`
	// Mode-2 interrogation response (military identification code).
	M2 string `json:"m2"`
	// Mode-3/A interrogation response (aircraft identification), provides a 4-digit
	// octal identification code for the aircraft, assigned by the air traffic
	// controller. Mode-3/A is shared military/civilian use.
	M3a string `json:"m3a"`
	// Numbered Air Force (NAF) organization that owns the mission.
	Naf string `json:"naf"`
	// Air Mobility Command (AMC) mission identifier of the next air transport mission.
	// Provides a method for AMC to link air transport missions together
	// chronologically for tasking and planning purposes.
	NextAmcMissionID string `json:"nextAMCMissionId"`
	// Unique identifier of the next mission provided by the originating source.
	// Provides a method for the data provider to link air transport missions together
	// chronologically for tasking and planning purposes.
	NextMissionID string `json:"nextMissionId"`
	// Designates the location responsible for mission transportation, logistics, or
	// distribution activities for an Area of Responsibility (AOR) within USTRANSCOM.
	Node string `json:"node"`
	// A description of this mission's objective.
	Objective string `json:"objective"`
	// The name of the operation that this mission supports.
	Operation string `json:"operation"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The mission identifier provided by the originating source.
	OrigMissionID string `json:"origMissionId"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Air Mobility Command (AMC) mission identifier of the previous air transport
	// mission. Provides a method for AMC to link air transport missions together
	// chronologically for tasking and planning purposes.
	PrevAmcMissionID string `json:"prevAMCMissionId"`
	// Unique identifier of the previous air transport mission provided by the
	// originating source. Provides a method for the data provider to link air
	// transport missions together chronologically for tasking and planning purposes.
	PrevMissionID string `json:"prevMissionId"`
	// A description of this mission's purpose (e.g. why this mission needs to happen,
	// what is the mission supporting, etc.).
	Purpose string `json:"purpose"`
	// Information related to the planning, load, status, and deployment or dispatch of
	// one aircraft to carry out a mission.
	Remarks []AirTransportMissionAbridgedRemark `json:"remarks"`
	// Collection of Requirements planned to be associated with this Air Transport
	// Mission.
	Requirements []AirTransportMissionAbridgedRequirement `json:"requirements"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The number of minutes a mission is off schedule based on the source system's
	// business rules. Positive numbers are early, negative numbers are late.
	SourceSysDeviation float64 `json:"sourceSysDeviation"`
	// Current state of the mission.
	State string `json:"state"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	Type string `json:"type"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Abp                   respjson.Field
		Alias                 respjson.Field
		AllocatedUnit         respjson.Field
		AmcMissionID          respjson.Field
		ApacsID               respjson.Field
		AtoCallSign           respjson.Field
		AtoMissionID          respjson.Field
		CallSign              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Cw                    respjson.Field
		DipWorksheetName      respjson.Field
		FirstPickUp           respjson.Field
		GdssMissionID         respjson.Field
		HazMat                respjson.Field
		JcsPriority           respjson.Field
		LastDropOff           respjson.Field
		LoadCategoryType      respjson.Field
		M1                    respjson.Field
		M2                    respjson.Field
		M3a                   respjson.Field
		Naf                   respjson.Field
		NextAmcMissionID      respjson.Field
		NextMissionID         respjson.Field
		Node                  respjson.Field
		Objective             respjson.Field
		Operation             respjson.Field
		Origin                respjson.Field
		OrigMissionID         respjson.Field
		OrigNetwork           respjson.Field
		PrevAmcMissionID      respjson.Field
		PrevMissionID         respjson.Field
		Purpose               respjson.Field
		Remarks               respjson.Field
		Requirements          respjson.Field
		SourceDl              respjson.Field
		SourceSysDeviation    respjson.Field
		State                 respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTransportMissionAbridged) RawJSON() string { return r.JSON.raw }
func (r *AirTransportMissionAbridged) UnmarshalJSON(data []byte) error {
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
type AirTransportMissionAbridgedDataMode string

const (
	AirTransportMissionAbridgedDataModeReal      AirTransportMissionAbridgedDataMode = "REAL"
	AirTransportMissionAbridgedDataModeTest      AirTransportMissionAbridgedDataMode = "TEST"
	AirTransportMissionAbridgedDataModeSimulated AirTransportMissionAbridgedDataMode = "SIMULATED"
	AirTransportMissionAbridgedDataModeExercise  AirTransportMissionAbridgedDataMode = "EXERCISE"
)

// Collection of Hazardous Material information planned to be associated with this
// Air Transport Mission.
type AirTransportMissionAbridgedHazMat struct {
	// Comma delimited list of Note IDs for Item Class Segregation groups, specific to
	// GDSS systems.
	ApplicableNotes string `json:"applicableNotes"`
	// Compatibility group code used to specify the controls for the transportation and
	// storage of hazardous materials according to the Hazardous Materials Regulations
	// issued by the U.S. Department of Transportation.
	Cgc string `json:"cgc"`
	// Comma delimited list of Note IDs for compatibility groups, specific to GDSS
	// systems.
	Cgn string `json:"cgn"`
	// Class and division of the hazardous material according to the Hazardous
	// Materials Regulations issued by the U.S. Department of Transportation.
	ClassDiv float64 `json:"classDiv"`
	// The hazMat identifier provided by the originating source.
	ExtHazMatID string `json:"extHazMatId"`
	// United Nations proper shipping name of the hazardous material according to the
	// Hazardous Materials Regulations issued by the U.S. Department of Transportation.
	ItemName string `json:"itemName"`
	// Net explosive weight of the hazardous material, in kilograms.
	NetExpWt float64 `json:"netExpWt"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is unloaded.
	OffIcao string `json:"offICAO"`
	// Itinerary number that identifies where the hazardous material is unloaded.
	OffItin int64 `json:"offItin"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is loaded.
	OnIcao string `json:"onICAO"`
	// Itinerary number that identifies where the hazardous material is loaded.
	OnItin int64 `json:"onItin"`
	// Number of pieces of hazardous cargo.
	Pieces int64 `json:"pieces"`
	// Flag indicating if hazardous material is associated with this air transport
	// mission. Possible values are P (planned to be associated with the mission) or A
	// (actually associated with the mission). Enum: [P, A].
	Planned string `json:"planned"`
	// United Nations number or North America number that identifies hazardous
	// materials according to the Hazardous Materials Regulations issued by the U.S.
	// Department of Transportation.
	UnNum string `json:"unNum"`
	// Total weight of hazardous cargo, including non-explosive parts, in kilograms.
	Weight float64 `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicableNotes respjson.Field
		Cgc             respjson.Field
		Cgn             respjson.Field
		ClassDiv        respjson.Field
		ExtHazMatID     respjson.Field
		ItemName        respjson.Field
		NetExpWt        respjson.Field
		OffIcao         respjson.Field
		OffItin         respjson.Field
		OnIcao          respjson.Field
		OnItin          respjson.Field
		Pieces          respjson.Field
		Planned         respjson.Field
		UnNum           respjson.Field
		Weight          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTransportMissionAbridgedHazMat) RawJSON() string { return r.JSON.raw }
func (r *AirTransportMissionAbridgedHazMat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Remarks associated with this Air Transport Mission.
type AirTransportMissionAbridgedRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date time.Time `json:"date" format:"date-time"`
	// Global Decision Support System (GDSS) remark identifier.
	GdssRemarkID string `json:"gdssRemarkId"`
	// If the remark is sortie specific, this is the number of the sortie it applies
	// to.
	ItineraryNum int64 `json:"itineraryNum"`
	// Text of the remark.
	Text string `json:"text"`
	// Remark type.
	Type string `json:"type"`
	// User who published the remark.
	User string `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date         respjson.Field
		GdssRemarkID respjson.Field
		ItineraryNum respjson.Field
		Text         respjson.Field
		Type         respjson.Field
		User         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTransportMissionAbridgedRemark) RawJSON() string { return r.JSON.raw }
func (r *AirTransportMissionAbridgedRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Requirements planned to be associated with this Air Transport
// Mission.
type AirTransportMissionAbridgedRequirement struct {
	// Total weight of the bulk cargo, in kilograms.
	BulkWeight float64 `json:"bulkWeight"`
	// Earliest available date the cargo can be picked up, in ISO 8601 UTC format with
	// millisecond precision.
	Ead time.Time `json:"ead" format:"date-time"`
	// Global Decision Support System (GDSS) mission requirement identifier.
	GdssReqID string `json:"gdssReqId"`
	// Latest available date the cargo may be delivered, in ISO 8601 UTC format with
	// millisecond precision.
	Lad time.Time `json:"lad" format:"date-time"`
	// Number of ambulatory patients tasked for the mission.
	NumAmbulatory int64 `json:"numAmbulatory"`
	// Number of attendants tasked for the mission.
	NumAttendant int64 `json:"numAttendant"`
	// Number of litter patients tasked for the mission.
	NumLitter int64 `json:"numLitter"`
	// Number of passengers associated with the mission.
	NumPax int64 `json:"numPax"`
	// Identifier of the offload itinerary location.
	OffloadID int64 `json:"offloadId"`
	// Offload location code.
	OffloadLoCode string `json:"offloadLOCode"`
	// Identifier of the onload itinerary location.
	OnloadID int64 `json:"onloadId"`
	// Onload location code.
	OnloadLoCode string `json:"onloadLOCode"`
	// Identification number of the Operation Plan (OPLAN) associated with this
	// mission.
	Oplan string `json:"oplan"`
	// Total weight of the outsize cargo, in kilograms.
	OutsizeWeight float64 `json:"outsizeWeight"`
	// Total weight of the oversized cargo, in kilograms.
	OversizeWeight float64 `json:"oversizeWeight"`
	// Project name.
	ProjName string `json:"projName"`
	// Transportation requirement number.
	TransReqNum string `json:"transReqNum"`
	// Unit line number.
	Uln string `json:"uln"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BulkWeight     respjson.Field
		Ead            respjson.Field
		GdssReqID      respjson.Field
		Lad            respjson.Field
		NumAmbulatory  respjson.Field
		NumAttendant   respjson.Field
		NumLitter      respjson.Field
		NumPax         respjson.Field
		OffloadID      respjson.Field
		OffloadLoCode  respjson.Field
		OnloadID       respjson.Field
		OnloadLoCode   respjson.Field
		Oplan          respjson.Field
		OutsizeWeight  respjson.Field
		OversizeWeight respjson.Field
		ProjName       respjson.Field
		TransReqNum    respjson.Field
		Uln            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTransportMissionAbridgedRequirement) RawJSON() string { return r.JSON.raw }
func (r *AirTransportMissionAbridgedRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirTransportMissionQueryhelpResponse struct {
	AodrSupported         bool                                            `json:"aodrSupported"`
	ClassificationMarking string                                          `json:"classificationMarking"`
	Description           string                                          `json:"description"`
	HistorySupported      bool                                            `json:"historySupported"`
	Name                  string                                          `json:"name"`
	Parameters            []AirTransportMissionQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                        `json:"requiredRoles"`
	RestSupported         bool                                            `json:"restSupported"`
	SortSupported         bool                                            `json:"sortSupported"`
	TypeName              string                                          `json:"typeName"`
	Uri                   string                                          `json:"uri"`
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
func (r AirTransportMissionQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *AirTransportMissionQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirTransportMissionQueryhelpResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTransportMissionQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *AirTransportMissionQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirTransportMissionNewParams struct {
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
	DataMode AirTransportMissionNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Air Battle Plan used to coordinate and integrate air assets for this
	// mission.
	Abp param.Opt[string] `json:"abp,omitzero"`
	// Mission alias.
	Alias param.Opt[string] `json:"alias,omitzero"`
	// The unit the mission is allocated to.
	AllocatedUnit param.Opt[string] `json:"allocatedUnit,omitzero"`
	// Air Mobility Command (AMC) mission identifier according to Mobility Air Forces
	// (MAF) Encode/Decode procedures.
	AmcMissionID param.Opt[string] `json:"amcMissionId,omitzero"`
	// The Aircraft and Personnel Automated Clearance System (APACS) system identifier
	// used to process and approve clearance requests.
	ApacsID param.Opt[string] `json:"apacsId,omitzero"`
	// The call sign assigned to this mission according to the Air Tasking Order (ATO).
	AtoCallSign param.Opt[string] `json:"atoCallSign,omitzero"`
	// The mission number according to the Air Tasking Order (ATO).
	AtoMissionID param.Opt[string] `json:"atoMissionId,omitzero"`
	// The call sign for this mission.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Flag indicating this is a close watch mission.
	Cw param.Opt[bool] `json:"cw,omitzero"`
	// Identifier of the Diplomatic Clearance Worksheet used to coordinate aircraft
	// clearance requests.
	DipWorksheetName param.Opt[string] `json:"dipWorksheetName,omitzero"`
	// The International Civil Aviation Organization (ICAO) site code of first cargo
	// pick up.
	FirstPickUp param.Opt[string] `json:"firstPickUp,omitzero"`
	// Global Decision Support System (GDSS) mission unique identifier.
	GdssMissionID param.Opt[string] `json:"gdssMissionId,omitzero"`
	// Highest Joint Chiefs of Staff priority of this mission.
	JcsPriority param.Opt[string] `json:"jcsPriority,omitzero"`
	// The International Civil Aviation Organization (ICAO) site code of last cargo
	// drop off.
	LastDropOff param.Opt[string] `json:"lastDropOff,omitzero"`
	// Load type of this mission (e.g. CARGO, MIXED, PASSENGER).
	LoadCategoryType param.Opt[string] `json:"loadCategoryType,omitzero"`
	// Mode-1 interrogation response (mission code), indicating mission or aircraft
	// type.
	M1 param.Opt[string] `json:"m1,omitzero"`
	// Mode-2 interrogation response (military identification code).
	M2 param.Opt[string] `json:"m2,omitzero"`
	// Mode-3/A interrogation response (aircraft identification), provides a 4-digit
	// octal identification code for the aircraft, assigned by the air traffic
	// controller. Mode-3/A is shared military/civilian use.
	M3a param.Opt[string] `json:"m3a,omitzero"`
	// Numbered Air Force (NAF) organization that owns the mission.
	Naf param.Opt[string] `json:"naf,omitzero"`
	// Air Mobility Command (AMC) mission identifier of the next air transport mission.
	// Provides a method for AMC to link air transport missions together
	// chronologically for tasking and planning purposes.
	NextAmcMissionID param.Opt[string] `json:"nextAMCMissionId,omitzero"`
	// Unique identifier of the next mission provided by the originating source.
	// Provides a method for the data provider to link air transport missions together
	// chronologically for tasking and planning purposes.
	NextMissionID param.Opt[string] `json:"nextMissionId,omitzero"`
	// Designates the location responsible for mission transportation, logistics, or
	// distribution activities for an Area of Responsibility (AOR) within USTRANSCOM.
	Node param.Opt[string] `json:"node,omitzero"`
	// A description of this mission's objective.
	Objective param.Opt[string] `json:"objective,omitzero"`
	// The name of the operation that this mission supports.
	Operation param.Opt[string] `json:"operation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The mission identifier provided by the originating source.
	OrigMissionID param.Opt[string] `json:"origMissionId,omitzero"`
	// Air Mobility Command (AMC) mission identifier of the previous air transport
	// mission. Provides a method for AMC to link air transport missions together
	// chronologically for tasking and planning purposes.
	PrevAmcMissionID param.Opt[string] `json:"prevAMCMissionId,omitzero"`
	// Unique identifier of the previous air transport mission provided by the
	// originating source. Provides a method for the data provider to link air
	// transport missions together chronologically for tasking and planning purposes.
	PrevMissionID param.Opt[string] `json:"prevMissionId,omitzero"`
	// A description of this mission's purpose (e.g. why this mission needs to happen,
	// what is the mission supporting, etc.).
	Purpose param.Opt[string] `json:"purpose,omitzero"`
	// The number of minutes a mission is off schedule based on the source system's
	// business rules. Positive numbers are early, negative numbers are late.
	SourceSysDeviation param.Opt[float64] `json:"sourceSysDeviation,omitzero"`
	// Current state of the mission.
	State param.Opt[string] `json:"state,omitzero"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// Collection of Hazardous Material information planned to be associated with this
	// Air Transport Mission.
	HazMat []AirTransportMissionNewParamsHazMat `json:"hazMat,omitzero"`
	// Information related to the planning, load, status, and deployment or dispatch of
	// one aircraft to carry out a mission.
	Remarks []AirTransportMissionNewParamsRemark `json:"remarks,omitzero"`
	// Collection of Requirements planned to be associated with this Air Transport
	// Mission.
	Requirements []AirTransportMissionNewParamsRequirement `json:"requirements,omitzero"`
	paramObj
}

func (r AirTransportMissionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AirTransportMissionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirTransportMissionNewParams) UnmarshalJSON(data []byte) error {
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
type AirTransportMissionNewParamsDataMode string

const (
	AirTransportMissionNewParamsDataModeReal      AirTransportMissionNewParamsDataMode = "REAL"
	AirTransportMissionNewParamsDataModeTest      AirTransportMissionNewParamsDataMode = "TEST"
	AirTransportMissionNewParamsDataModeSimulated AirTransportMissionNewParamsDataMode = "SIMULATED"
	AirTransportMissionNewParamsDataModeExercise  AirTransportMissionNewParamsDataMode = "EXERCISE"
)

// Collection of Hazardous Material information planned to be associated with this
// Air Transport Mission.
type AirTransportMissionNewParamsHazMat struct {
	// Comma delimited list of Note IDs for Item Class Segregation groups, specific to
	// GDSS systems.
	ApplicableNotes param.Opt[string] `json:"applicableNotes,omitzero"`
	// Compatibility group code used to specify the controls for the transportation and
	// storage of hazardous materials according to the Hazardous Materials Regulations
	// issued by the U.S. Department of Transportation.
	Cgc param.Opt[string] `json:"cgc,omitzero"`
	// Comma delimited list of Note IDs for compatibility groups, specific to GDSS
	// systems.
	Cgn param.Opt[string] `json:"cgn,omitzero"`
	// Class and division of the hazardous material according to the Hazardous
	// Materials Regulations issued by the U.S. Department of Transportation.
	ClassDiv param.Opt[float64] `json:"classDiv,omitzero"`
	// The hazMat identifier provided by the originating source.
	ExtHazMatID param.Opt[string] `json:"extHazMatId,omitzero"`
	// United Nations proper shipping name of the hazardous material according to the
	// Hazardous Materials Regulations issued by the U.S. Department of Transportation.
	ItemName param.Opt[string] `json:"itemName,omitzero"`
	// Net explosive weight of the hazardous material, in kilograms.
	NetExpWt param.Opt[float64] `json:"netExpWt,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is unloaded.
	OffIcao param.Opt[string] `json:"offICAO,omitzero"`
	// Itinerary number that identifies where the hazardous material is unloaded.
	OffItin param.Opt[int64] `json:"offItin,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is loaded.
	OnIcao param.Opt[string] `json:"onICAO,omitzero"`
	// Itinerary number that identifies where the hazardous material is loaded.
	OnItin param.Opt[int64] `json:"onItin,omitzero"`
	// Number of pieces of hazardous cargo.
	Pieces param.Opt[int64] `json:"pieces,omitzero"`
	// Flag indicating if hazardous material is associated with this air transport
	// mission. Possible values are P (planned to be associated with the mission) or A
	// (actually associated with the mission). Enum: [P, A].
	Planned param.Opt[string] `json:"planned,omitzero"`
	// United Nations number or North America number that identifies hazardous
	// materials according to the Hazardous Materials Regulations issued by the U.S.
	// Department of Transportation.
	UnNum param.Opt[string] `json:"unNum,omitzero"`
	// Total weight of hazardous cargo, including non-explosive parts, in kilograms.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r AirTransportMissionNewParamsHazMat) MarshalJSON() (data []byte, err error) {
	type shadow AirTransportMissionNewParamsHazMat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirTransportMissionNewParamsHazMat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Remarks associated with this Air Transport Mission.
type AirTransportMissionNewParamsRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Global Decision Support System (GDSS) remark identifier.
	GdssRemarkID param.Opt[string] `json:"gdssRemarkId,omitzero"`
	// If the remark is sortie specific, this is the number of the sortie it applies
	// to.
	ItineraryNum param.Opt[int64] `json:"itineraryNum,omitzero"`
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// Remark type.
	Type param.Opt[string] `json:"type,omitzero"`
	// User who published the remark.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r AirTransportMissionNewParamsRemark) MarshalJSON() (data []byte, err error) {
	type shadow AirTransportMissionNewParamsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirTransportMissionNewParamsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Requirements planned to be associated with this Air Transport
// Mission.
type AirTransportMissionNewParamsRequirement struct {
	// Total weight of the bulk cargo, in kilograms.
	BulkWeight param.Opt[float64] `json:"bulkWeight,omitzero"`
	// Earliest available date the cargo can be picked up, in ISO 8601 UTC format with
	// millisecond precision.
	Ead param.Opt[time.Time] `json:"ead,omitzero" format:"date-time"`
	// Global Decision Support System (GDSS) mission requirement identifier.
	GdssReqID param.Opt[string] `json:"gdssReqId,omitzero"`
	// Latest available date the cargo may be delivered, in ISO 8601 UTC format with
	// millisecond precision.
	Lad param.Opt[time.Time] `json:"lad,omitzero" format:"date-time"`
	// Number of ambulatory patients tasked for the mission.
	NumAmbulatory param.Opt[int64] `json:"numAmbulatory,omitzero"`
	// Number of attendants tasked for the mission.
	NumAttendant param.Opt[int64] `json:"numAttendant,omitzero"`
	// Number of litter patients tasked for the mission.
	NumLitter param.Opt[int64] `json:"numLitter,omitzero"`
	// Number of passengers associated with the mission.
	NumPax param.Opt[int64] `json:"numPax,omitzero"`
	// Identifier of the offload itinerary location.
	OffloadID param.Opt[int64] `json:"offloadId,omitzero"`
	// Offload location code.
	OffloadLoCode param.Opt[string] `json:"offloadLOCode,omitzero"`
	// Identifier of the onload itinerary location.
	OnloadID param.Opt[int64] `json:"onloadId,omitzero"`
	// Onload location code.
	OnloadLoCode param.Opt[string] `json:"onloadLOCode,omitzero"`
	// Identification number of the Operation Plan (OPLAN) associated with this
	// mission.
	Oplan param.Opt[string] `json:"oplan,omitzero"`
	// Total weight of the outsize cargo, in kilograms.
	OutsizeWeight param.Opt[float64] `json:"outsizeWeight,omitzero"`
	// Total weight of the oversized cargo, in kilograms.
	OversizeWeight param.Opt[float64] `json:"oversizeWeight,omitzero"`
	// Project name.
	ProjName param.Opt[string] `json:"projName,omitzero"`
	// Transportation requirement number.
	TransReqNum param.Opt[string] `json:"transReqNum,omitzero"`
	// Unit line number.
	Uln param.Opt[string] `json:"uln,omitzero"`
	paramObj
}

func (r AirTransportMissionNewParamsRequirement) MarshalJSON() (data []byte, err error) {
	type shadow AirTransportMissionNewParamsRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirTransportMissionNewParamsRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirTransportMissionGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirTransportMissionGetParams]'s query parameters as
// `url.Values`.
func (r AirTransportMissionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirTransportMissionUpdateParams struct {
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
	DataMode AirTransportMissionUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Air Battle Plan used to coordinate and integrate air assets for this
	// mission.
	Abp param.Opt[string] `json:"abp,omitzero"`
	// Mission alias.
	Alias param.Opt[string] `json:"alias,omitzero"`
	// The unit the mission is allocated to.
	AllocatedUnit param.Opt[string] `json:"allocatedUnit,omitzero"`
	// Air Mobility Command (AMC) mission identifier according to Mobility Air Forces
	// (MAF) Encode/Decode procedures.
	AmcMissionID param.Opt[string] `json:"amcMissionId,omitzero"`
	// The Aircraft and Personnel Automated Clearance System (APACS) system identifier
	// used to process and approve clearance requests.
	ApacsID param.Opt[string] `json:"apacsId,omitzero"`
	// The call sign assigned to this mission according to the Air Tasking Order (ATO).
	AtoCallSign param.Opt[string] `json:"atoCallSign,omitzero"`
	// The mission number according to the Air Tasking Order (ATO).
	AtoMissionID param.Opt[string] `json:"atoMissionId,omitzero"`
	// The call sign for this mission.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Flag indicating this is a close watch mission.
	Cw param.Opt[bool] `json:"cw,omitzero"`
	// Identifier of the Diplomatic Clearance Worksheet used to coordinate aircraft
	// clearance requests.
	DipWorksheetName param.Opt[string] `json:"dipWorksheetName,omitzero"`
	// The International Civil Aviation Organization (ICAO) site code of first cargo
	// pick up.
	FirstPickUp param.Opt[string] `json:"firstPickUp,omitzero"`
	// Global Decision Support System (GDSS) mission unique identifier.
	GdssMissionID param.Opt[string] `json:"gdssMissionId,omitzero"`
	// Highest Joint Chiefs of Staff priority of this mission.
	JcsPriority param.Opt[string] `json:"jcsPriority,omitzero"`
	// The International Civil Aviation Organization (ICAO) site code of last cargo
	// drop off.
	LastDropOff param.Opt[string] `json:"lastDropOff,omitzero"`
	// Load type of this mission (e.g. CARGO, MIXED, PASSENGER).
	LoadCategoryType param.Opt[string] `json:"loadCategoryType,omitzero"`
	// Mode-1 interrogation response (mission code), indicating mission or aircraft
	// type.
	M1 param.Opt[string] `json:"m1,omitzero"`
	// Mode-2 interrogation response (military identification code).
	M2 param.Opt[string] `json:"m2,omitzero"`
	// Mode-3/A interrogation response (aircraft identification), provides a 4-digit
	// octal identification code for the aircraft, assigned by the air traffic
	// controller. Mode-3/A is shared military/civilian use.
	M3a param.Opt[string] `json:"m3a,omitzero"`
	// Numbered Air Force (NAF) organization that owns the mission.
	Naf param.Opt[string] `json:"naf,omitzero"`
	// Air Mobility Command (AMC) mission identifier of the next air transport mission.
	// Provides a method for AMC to link air transport missions together
	// chronologically for tasking and planning purposes.
	NextAmcMissionID param.Opt[string] `json:"nextAMCMissionId,omitzero"`
	// Unique identifier of the next mission provided by the originating source.
	// Provides a method for the data provider to link air transport missions together
	// chronologically for tasking and planning purposes.
	NextMissionID param.Opt[string] `json:"nextMissionId,omitzero"`
	// Designates the location responsible for mission transportation, logistics, or
	// distribution activities for an Area of Responsibility (AOR) within USTRANSCOM.
	Node param.Opt[string] `json:"node,omitzero"`
	// A description of this mission's objective.
	Objective param.Opt[string] `json:"objective,omitzero"`
	// The name of the operation that this mission supports.
	Operation param.Opt[string] `json:"operation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The mission identifier provided by the originating source.
	OrigMissionID param.Opt[string] `json:"origMissionId,omitzero"`
	// Air Mobility Command (AMC) mission identifier of the previous air transport
	// mission. Provides a method for AMC to link air transport missions together
	// chronologically for tasking and planning purposes.
	PrevAmcMissionID param.Opt[string] `json:"prevAMCMissionId,omitzero"`
	// Unique identifier of the previous air transport mission provided by the
	// originating source. Provides a method for the data provider to link air
	// transport missions together chronologically for tasking and planning purposes.
	PrevMissionID param.Opt[string] `json:"prevMissionId,omitzero"`
	// A description of this mission's purpose (e.g. why this mission needs to happen,
	// what is the mission supporting, etc.).
	Purpose param.Opt[string] `json:"purpose,omitzero"`
	// The number of minutes a mission is off schedule based on the source system's
	// business rules. Positive numbers are early, negative numbers are late.
	SourceSysDeviation param.Opt[float64] `json:"sourceSysDeviation,omitzero"`
	// Current state of the mission.
	State param.Opt[string] `json:"state,omitzero"`
	// The type of mission (e.g. SAAM, CHNL, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// Collection of Hazardous Material information planned to be associated with this
	// Air Transport Mission.
	HazMat []AirTransportMissionUpdateParamsHazMat `json:"hazMat,omitzero"`
	// Information related to the planning, load, status, and deployment or dispatch of
	// one aircraft to carry out a mission.
	Remarks []AirTransportMissionUpdateParamsRemark `json:"remarks,omitzero"`
	// Collection of Requirements planned to be associated with this Air Transport
	// Mission.
	Requirements []AirTransportMissionUpdateParamsRequirement `json:"requirements,omitzero"`
	paramObj
}

func (r AirTransportMissionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AirTransportMissionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirTransportMissionUpdateParams) UnmarshalJSON(data []byte) error {
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
type AirTransportMissionUpdateParamsDataMode string

const (
	AirTransportMissionUpdateParamsDataModeReal      AirTransportMissionUpdateParamsDataMode = "REAL"
	AirTransportMissionUpdateParamsDataModeTest      AirTransportMissionUpdateParamsDataMode = "TEST"
	AirTransportMissionUpdateParamsDataModeSimulated AirTransportMissionUpdateParamsDataMode = "SIMULATED"
	AirTransportMissionUpdateParamsDataModeExercise  AirTransportMissionUpdateParamsDataMode = "EXERCISE"
)

// Collection of Hazardous Material information planned to be associated with this
// Air Transport Mission.
type AirTransportMissionUpdateParamsHazMat struct {
	// Comma delimited list of Note IDs for Item Class Segregation groups, specific to
	// GDSS systems.
	ApplicableNotes param.Opt[string] `json:"applicableNotes,omitzero"`
	// Compatibility group code used to specify the controls for the transportation and
	// storage of hazardous materials according to the Hazardous Materials Regulations
	// issued by the U.S. Department of Transportation.
	Cgc param.Opt[string] `json:"cgc,omitzero"`
	// Comma delimited list of Note IDs for compatibility groups, specific to GDSS
	// systems.
	Cgn param.Opt[string] `json:"cgn,omitzero"`
	// Class and division of the hazardous material according to the Hazardous
	// Materials Regulations issued by the U.S. Department of Transportation.
	ClassDiv param.Opt[float64] `json:"classDiv,omitzero"`
	// The hazMat identifier provided by the originating source.
	ExtHazMatID param.Opt[string] `json:"extHazMatId,omitzero"`
	// United Nations proper shipping name of the hazardous material according to the
	// Hazardous Materials Regulations issued by the U.S. Department of Transportation.
	ItemName param.Opt[string] `json:"itemName,omitzero"`
	// Net explosive weight of the hazardous material, in kilograms.
	NetExpWt param.Opt[float64] `json:"netExpWt,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is unloaded.
	OffIcao param.Opt[string] `json:"offICAO,omitzero"`
	// Itinerary number that identifies where the hazardous material is unloaded.
	OffItin param.Opt[int64] `json:"offItin,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is loaded.
	OnIcao param.Opt[string] `json:"onICAO,omitzero"`
	// Itinerary number that identifies where the hazardous material is loaded.
	OnItin param.Opt[int64] `json:"onItin,omitzero"`
	// Number of pieces of hazardous cargo.
	Pieces param.Opt[int64] `json:"pieces,omitzero"`
	// Flag indicating if hazardous material is associated with this air transport
	// mission. Possible values are P (planned to be associated with the mission) or A
	// (actually associated with the mission). Enum: [P, A].
	Planned param.Opt[string] `json:"planned,omitzero"`
	// United Nations number or North America number that identifies hazardous
	// materials according to the Hazardous Materials Regulations issued by the U.S.
	// Department of Transportation.
	UnNum param.Opt[string] `json:"unNum,omitzero"`
	// Total weight of hazardous cargo, including non-explosive parts, in kilograms.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r AirTransportMissionUpdateParamsHazMat) MarshalJSON() (data []byte, err error) {
	type shadow AirTransportMissionUpdateParamsHazMat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirTransportMissionUpdateParamsHazMat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Remarks associated with this Air Transport Mission.
type AirTransportMissionUpdateParamsRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Global Decision Support System (GDSS) remark identifier.
	GdssRemarkID param.Opt[string] `json:"gdssRemarkId,omitzero"`
	// If the remark is sortie specific, this is the number of the sortie it applies
	// to.
	ItineraryNum param.Opt[int64] `json:"itineraryNum,omitzero"`
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// Remark type.
	Type param.Opt[string] `json:"type,omitzero"`
	// User who published the remark.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r AirTransportMissionUpdateParamsRemark) MarshalJSON() (data []byte, err error) {
	type shadow AirTransportMissionUpdateParamsRemark
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirTransportMissionUpdateParamsRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Requirements planned to be associated with this Air Transport
// Mission.
type AirTransportMissionUpdateParamsRequirement struct {
	// Total weight of the bulk cargo, in kilograms.
	BulkWeight param.Opt[float64] `json:"bulkWeight,omitzero"`
	// Earliest available date the cargo can be picked up, in ISO 8601 UTC format with
	// millisecond precision.
	Ead param.Opt[time.Time] `json:"ead,omitzero" format:"date-time"`
	// Global Decision Support System (GDSS) mission requirement identifier.
	GdssReqID param.Opt[string] `json:"gdssReqId,omitzero"`
	// Latest available date the cargo may be delivered, in ISO 8601 UTC format with
	// millisecond precision.
	Lad param.Opt[time.Time] `json:"lad,omitzero" format:"date-time"`
	// Number of ambulatory patients tasked for the mission.
	NumAmbulatory param.Opt[int64] `json:"numAmbulatory,omitzero"`
	// Number of attendants tasked for the mission.
	NumAttendant param.Opt[int64] `json:"numAttendant,omitzero"`
	// Number of litter patients tasked for the mission.
	NumLitter param.Opt[int64] `json:"numLitter,omitzero"`
	// Number of passengers associated with the mission.
	NumPax param.Opt[int64] `json:"numPax,omitzero"`
	// Identifier of the offload itinerary location.
	OffloadID param.Opt[int64] `json:"offloadId,omitzero"`
	// Offload location code.
	OffloadLoCode param.Opt[string] `json:"offloadLOCode,omitzero"`
	// Identifier of the onload itinerary location.
	OnloadID param.Opt[int64] `json:"onloadId,omitzero"`
	// Onload location code.
	OnloadLoCode param.Opt[string] `json:"onloadLOCode,omitzero"`
	// Identification number of the Operation Plan (OPLAN) associated with this
	// mission.
	Oplan param.Opt[string] `json:"oplan,omitzero"`
	// Total weight of the outsize cargo, in kilograms.
	OutsizeWeight param.Opt[float64] `json:"outsizeWeight,omitzero"`
	// Total weight of the oversized cargo, in kilograms.
	OversizeWeight param.Opt[float64] `json:"oversizeWeight,omitzero"`
	// Project name.
	ProjName param.Opt[string] `json:"projName,omitzero"`
	// Transportation requirement number.
	TransReqNum param.Opt[string] `json:"transReqNum,omitzero"`
	// Unit line number.
	Uln param.Opt[string] `json:"uln,omitzero"`
	paramObj
}

func (r AirTransportMissionUpdateParamsRequirement) MarshalJSON() (data []byte, err error) {
	type shadow AirTransportMissionUpdateParamsRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirTransportMissionUpdateParamsRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirTransportMissionListParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirTransportMissionListParams]'s query parameters as
// `url.Values`.
func (r AirTransportMissionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirTransportMissionCountParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirTransportMissionCountParams]'s query parameters as
// `url.Values`.
func (r AirTransportMissionCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirTransportMissionTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirTransportMissionTupleParams]'s query parameters as
// `url.Values`.
func (r AirTransportMissionTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
