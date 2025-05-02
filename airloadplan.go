// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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
)

// AirLoadPlanService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirLoadPlanService] method instead.
type AirLoadPlanService struct {
	Options []option.RequestOption
}

// NewAirLoadPlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAirLoadPlanService(opts ...option.RequestOption) (r AirLoadPlanService) {
	r = AirLoadPlanService{}
	r.Options = opts
	return
}

// Service operation to take a single airloadplan record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AirLoadPlanService) New(ctx context.Context, body AirLoadPlanNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airloadplan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single airloadplan record by its unique ID passed as
// a path parameter.
func (r *AirLoadPlanService) Get(ctx context.Context, id string, query AirLoadPlanGetParams, opts ...option.RequestOption) (res *AirloadplanFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airloadplan/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirLoadPlanService) List(ctx context.Context, query AirLoadPlanListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AirloadplanAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/airloadplan"
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
func (r *AirLoadPlanService) ListAutoPaging(ctx context.Context, query AirLoadPlanListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AirloadplanAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirLoadPlanService) Count(ctx context.Context, query AirLoadPlanCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airloadplan/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AirLoadPlanService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airloadplan/queryhelp"
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
func (r *AirLoadPlanService) Tuple(ctx context.Context, query AirLoadPlanTupleParams, opts ...option.RequestOption) (res *[]AirloadplanFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airloadplan/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information related to how an aircraft is loaded with cargo, equipment, and
// passengers.
type AirloadplanAbridged struct {
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
	DataMode AirloadplanAbridgedDataMode `json:"dataMode,required"`
	// The current estimated time that the aircraft is planned to depart, in ISO 8601
	// UTC format with millisecond precision.
	EstDepTime time.Time `json:"estDepTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Allowable Cabin Load (ACL) onboard the aircraft. The maximum weight of
	// passengers, baggage, and cargo that can be safely transported in the aircraft
	// cabin, in kilograms.
	ACLOnboard float64 `json:"aclOnboard"`
	// Allowable Cabin Load (ACL) released this leg. The weight of passengers, baggage,
	// and cargo released from the aircraft cabin, in kilograms.
	ACLReleased float64 `json:"aclReleased"`
	// The Model Design Series designation of the aircraft supporting this load plan.
	AircraftMds string `json:"aircraftMDS"`
	// Collection of hazmat actuals associated with this load plan.
	AirLoadPlanHazmatActuals []AirloadplanAbridgedAirLoadPlanHazmatActual `json:"airLoadPlanHazmatActuals"`
	// Collection of human remains transport information associated with this load
	// plan.
	AirLoadPlanHr []AirloadplanAbridgedAirLoadPlanHr `json:"airLoadPlanHR"`
	// Collection of cargo information located at the pallet positions associated with
	// this load plan.
	AirLoadPlanPalletDetails []AirloadplanAbridgedAirLoadPlanPalletDetail `json:"airLoadPlanPalletDetails"`
	// Collection of passenger and cargo details associated with this load plan for
	// this leg of the mission.
	AirLoadPlanPaxCargo []AirloadplanAbridgedAirLoadPlanPaxCargo `json:"airLoadPlanPaxCargo"`
	// Collection of unit line number actuals associated with this load plan.
	AirLoadPlanUlnActuals []AirloadplanAbridgedAirLoadPlanUlnActual `json:"airLoadPlanULNActuals"`
	// Optional identifier of arrival airfield with no International Civil Organization
	// (ICAO) code.
	ArrAirfield string `json:"arrAirfield"`
	// The arrival International Civil Organization (ICAO) code of the landing
	// airfield.
	ArrIcao string `json:"arrICAO"`
	// Time the loadmaster or boom operator is available for cargo loading/unloading,
	// in ISO 8601 UTC format with millisecond precision.
	AvailableTime time.Time `json:"availableTime" format:"date-time"`
	// The basic weight of the aircraft multiplied by the distance between the
	// reference datum and the aircraft's center of gravity, in Newton-meters.
	BasicMoment float64 `json:"basicMoment"`
	// The weight of the aircraft without passengers, cargo, equipment, or usable fuel,
	// in kilograms.
	BasicWeight float64 `json:"basicWeight"`
	// Time the cargo briefing was given to the loadmaster or boom operator, in ISO
	// 8601 UTC format with millisecond precision.
	BriefTime time.Time `json:"briefTime" format:"date-time"`
	// The call sign of the mission supporting this load plan.
	CallSign string `json:"callSign"`
	// Maximum fuselage station (FS) where cargo can be stored. FS is the distance from
	// the reference datum, in meters.
	CargoBayFsMax float64 `json:"cargoBayFSMax"`
	// Minimum fuselage station (FS) where cargo can be stored. FS is the distance from
	// the reference datum, in meters.
	CargoBayFsMin float64 `json:"cargoBayFSMin"`
	// Width of the cargo bay, in meters.
	CargoBayWidth float64 `json:"cargoBayWidth"`
	// The cargo configuration required for this leg (e.g. C-1, C-2, C-3, DV-1, DV-2,
	// AE-1, etc.). Configuration meanings are determined by the data source.
	CargoConfig string `json:"cargoConfig"`
	// The sum of cargo moments of all cargo on board the aircraft, in Newton-meters.
	// Each individual cargo moment is the weight of the cargo multiplied by the
	// distance between the reference datum and the cargo's center of gravity.
	CargoMoment float64 `json:"cargoMoment"`
	// Volume of cargo space in the aircraft, in cubic meters.
	CargoVolume float64 `json:"cargoVolume"`
	// The weight of the cargo on board the aircraft, in kilograms.
	CargoWeight float64 `json:"cargoWeight"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The number of crew members on the aircraft.
	CrewSize int64 `json:"crewSize"`
	// Optional identifier of departure airfield with no International Civil
	// Organization (ICAO) code.
	DepAirfield string `json:"depAirfield"`
	// The departure International Civil Organization (ICAO) code of the departure
	// airfield.
	DepIcao string `json:"depICAO"`
	// Description of the equipment configuration (e.g. Standard, Ferry, JBLM, CHS,
	// Combat, etc.). Configuration meanings are determined by the data source.
	EquipConfig string `json:"equipConfig"`
	// The current estimated time that the aircraft is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime time.Time `json:"estArrTime" format:"date-time"`
	// The estimated weight of usable fuel upon landing multiplied by the distance
	// between the reference datum and the fuel's center of gravity, in Newton-meters.
	EstLandingFuelMoment float64 `json:"estLandingFuelMoment"`
	// The estimated weight of usable fuel upon landing, in kilograms.
	EstLandingFuelWeight float64 `json:"estLandingFuelWeight"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// The fuel weight on board the aircraft multiplied by the distance between the
	// reference datum and the fuel's center of gravity, in Newton-meters.
	FuelMoment float64 `json:"fuelMoment"`
	// The weight of usable fuel on board the aircraft, in kilograms.
	FuelWeight float64 `json:"fuelWeight"`
	// The center of gravity of the aircraft using the gross weight and gross moment,
	// as a percentage of the mean aerodynamic chord (%MAC).
	GrossCg float64 `json:"grossCG"`
	// The sum of moments of all items making up the gross weight of the aircraft, in
	// Newton-meters.
	GrossMoment float64 `json:"grossMoment"`
	// The total weight of the aircraft at takeoff including passengers, cargo,
	// equipment, and usable fuel, in kilograms.
	GrossWeight float64 `json:"grossWeight"`
	// The UDL ID of the mission this record is associated with.
	IDMission string `json:"idMission"`
	// The UDL ID of the aircraft sortie this record is associated with.
	IDSortie string `json:"idSortie"`
	// The center of gravity of the aircraft using the landing weight and landing
	// moment, as a percentage of the mean aerodynamic chord (%MAC).
	LandingCg float64 `json:"landingCG"`
	// The sum of moments of all items making up the gross weight of the aircraft upon
	// landing, in Newton-meters.
	LandingMoment float64 `json:"landingMoment"`
	// The gross weight of the aircraft upon landing, in kilograms.
	LandingWeight float64 `json:"landingWeight"`
	// The leg number of the mission supporting this load plan.
	LegNum int64 `json:"legNum"`
	// Name of the loadmaster or boom operator who received the cargo briefing.
	LoadmasterName string `json:"loadmasterName"`
	// Rank of the loadmaster or boom operator overseeing cargo loading/unloading.
	LoadmasterRank string `json:"loadmasterRank"`
	// Remarks concerning this load plan.
	LoadRemarks string `json:"loadRemarks"`
	// The mission number of the mission supporting this load plan.
	MissionNumber string `json:"missionNumber"`
	// The operating weight of the aircraft multiplied by the distance between the
	// reference datum and the aircraft's center of gravity, in Newton-meters.
	OperatingMoment float64 `json:"operatingMoment"`
	// The basic weight of the aircraft including passengers and equipment, in
	// kilograms.
	OperatingWeight float64 `json:"operatingWeight"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Number of pallet positions on the aircraft.
	PpOnboard int64 `json:"ppOnboard"`
	// Number of pallet positions released this leg.
	PpReleased int64 `json:"ppReleased"`
	// Time the loadmaster or boom operator is scheduled to begin overseeing cargo
	// loading/unloading, in ISO 8601 UTC format with millisecond precision.
	SchedTime time.Time `json:"schedTime" format:"date-time"`
	// Number of passenger seats on the aircraft.
	SeatsOnboard int64 `json:"seatsOnboard"`
	// Number of passenger seats released this leg.
	SeatsReleased int64 `json:"seatsReleased"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The tail number of the aircraft supporting this load plan.
	TailNumber string `json:"tailNumber"`
	// Description of the fuel tank(s) configuration (e.g. ER, NON-ER, etc.).
	// Configuration meanings are determined by the data source.
	TankConfig string `json:"tankConfig"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Alphanumeric code that describes general cargo-related utilization and
	// characteristics for an itinerary point.
	UtilCode string `json:"utilCode"`
	// The center of gravity of the aircraft using the zero fuel weight and zero fuel
	// total moment, as a percentage of the mean aerodynamic chord (%MAC).
	ZeroFuelCg float64 `json:"zeroFuelCG"`
	// The zero fuel weight of the aircraft multiplied by the distance between the
	// reference datum and the aircraft's center of gravity, in Newton-meters.
	ZeroFuelMoment float64 `json:"zeroFuelMoment"`
	// The operating weight of the aircraft including cargo, mail, baggage, and
	// passengers, but without usable fuel, in kilograms.
	ZeroFuelWeight float64 `json:"zeroFuelWeight"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking    resp.Field
		DataMode                 resp.Field
		EstDepTime               resp.Field
		Source                   resp.Field
		ID                       resp.Field
		ACLOnboard               resp.Field
		ACLReleased              resp.Field
		AircraftMds              resp.Field
		AirLoadPlanHazmatActuals resp.Field
		AirLoadPlanHr            resp.Field
		AirLoadPlanPalletDetails resp.Field
		AirLoadPlanPaxCargo      resp.Field
		AirLoadPlanUlnActuals    resp.Field
		ArrAirfield              resp.Field
		ArrIcao                  resp.Field
		AvailableTime            resp.Field
		BasicMoment              resp.Field
		BasicWeight              resp.Field
		BriefTime                resp.Field
		CallSign                 resp.Field
		CargoBayFsMax            resp.Field
		CargoBayFsMin            resp.Field
		CargoBayWidth            resp.Field
		CargoConfig              resp.Field
		CargoMoment              resp.Field
		CargoVolume              resp.Field
		CargoWeight              resp.Field
		CreatedAt                resp.Field
		CreatedBy                resp.Field
		CrewSize                 resp.Field
		DepAirfield              resp.Field
		DepIcao                  resp.Field
		EquipConfig              resp.Field
		EstArrTime               resp.Field
		EstLandingFuelMoment     resp.Field
		EstLandingFuelWeight     resp.Field
		ExternalID               resp.Field
		FuelMoment               resp.Field
		FuelWeight               resp.Field
		GrossCg                  resp.Field
		GrossMoment              resp.Field
		GrossWeight              resp.Field
		IDMission                resp.Field
		IDSortie                 resp.Field
		LandingCg                resp.Field
		LandingMoment            resp.Field
		LandingWeight            resp.Field
		LegNum                   resp.Field
		LoadmasterName           resp.Field
		LoadmasterRank           resp.Field
		LoadRemarks              resp.Field
		MissionNumber            resp.Field
		OperatingMoment          resp.Field
		OperatingWeight          resp.Field
		Origin                   resp.Field
		OrigNetwork              resp.Field
		PpOnboard                resp.Field
		PpReleased               resp.Field
		SchedTime                resp.Field
		SeatsOnboard             resp.Field
		SeatsReleased            resp.Field
		SourceDl                 resp.Field
		TailNumber               resp.Field
		TankConfig               resp.Field
		UpdatedAt                resp.Field
		UpdatedBy                resp.Field
		UtilCode                 resp.Field
		ZeroFuelCg               resp.Field
		ZeroFuelMoment           resp.Field
		ZeroFuelWeight           resp.Field
		ExtraFields              map[string]resp.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanAbridged) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanAbridged) UnmarshalJSON(data []byte) error {
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
type AirloadplanAbridgedDataMode string

const (
	AirloadplanAbridgedDataModeReal      AirloadplanAbridgedDataMode = "REAL"
	AirloadplanAbridgedDataModeTest      AirloadplanAbridgedDataMode = "TEST"
	AirloadplanAbridgedDataModeSimulated AirloadplanAbridgedDataMode = "SIMULATED"
	AirloadplanAbridgedDataModeExercise  AirloadplanAbridgedDataMode = "EXERCISE"
)

// Collection of hazmat actuals associated with this load plan.
type AirloadplanAbridgedAirLoadPlanHazmatActual struct {
	// The Air Special Handling Code (ASHC) indicates the type of special handling
	// required for hazardous cargo.
	Ashc string `json:"ashc"`
	// Compatibility group code used to specify the controls for the transportation and
	// storage of hazardous materials according to the Hazardous Materials Regulations
	// issued by the U.S. Department of Transportation.
	Cgc string `json:"cgc"`
	// Class and division of the hazardous material according to the Hazardous
	// Materials Regulations issued by the U.S. Department of Transportation.
	ClassDiv string `json:"classDiv"`
	// Description of the hazardous item.
	HazDescription string `json:"hazDescription"`
	// Remarks concerning this hazardous material.
	HazmatRemarks string `json:"hazmatRemarks"`
	// United Nations number or North American number that identifies hazardous
	// materials according to the Hazardous Materials Regulations issued by the U.S.
	// Department of Transportation.
	HazNum string `json:"hazNum"`
	// Designates the type of hazmat number for the item (UN for United Nations or NA
	// for North American).
	HazNumType string `json:"hazNumType"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is unloaded.
	HazOffIcao string `json:"hazOffICAO"`
	// Itinerary number that identifies where the hazardous material is unloaded.
	HazOffItin int64 `json:"hazOffItin"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is loaded.
	HazOnIcao string `json:"hazOnICAO"`
	// Itinerary number that identifies where the hazardous material is loaded.
	HazOnItin int64 `json:"hazOnItin"`
	// Number of pieces of hazardous cargo.
	HazPieces int64 `json:"hazPieces"`
	// Transportation Control Number (TCN) of the hazardous item.
	HazTcn string `json:"hazTcn"`
	// Total weight of hazardous cargo, including non-explosive parts, in kilograms.
	HazWeight float64 `json:"hazWeight"`
	// United Nations proper shipping name of the hazardous material according to the
	// Hazardous Materials Regulations issued by the U.S. Department of Transportation.
	ItemName string `json:"itemName"`
	// Manufacturer's lot number for identification of the hazardous material.
	LotNum string `json:"lotNum"`
	// Net explosive weight of the hazardous material, in kilograms.
	NetExpWt float64 `json:"netExpWt"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Ashc           resp.Field
		Cgc            resp.Field
		ClassDiv       resp.Field
		HazDescription resp.Field
		HazmatRemarks  resp.Field
		HazNum         resp.Field
		HazNumType     resp.Field
		HazOffIcao     resp.Field
		HazOffItin     resp.Field
		HazOnIcao      resp.Field
		HazOnItin      resp.Field
		HazPieces      resp.Field
		HazTcn         resp.Field
		HazWeight      resp.Field
		ItemName       resp.Field
		LotNum         resp.Field
		NetExpWt       resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanAbridgedAirLoadPlanHazmatActual) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanAbridgedAirLoadPlanHazmatActual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of human remains transport information associated with this load
// plan.
type AirloadplanAbridgedAirLoadPlanHr struct {
	// Type of transfer case used.
	Container string `json:"container"`
	// Name of the escort for the remains.
	Escort string `json:"escort"`
	// The current estimated time of arrival for the remains in ISO 8601 UTC format
	// with millisecond precision.
	HrEstArrTime time.Time `json:"hrEstArrTime" format:"date-time"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// remains are unloaded.
	HrOffIcao string `json:"hrOffICAO"`
	// Itinerary number that identifies where the remains are unloaded.
	HrOffItin int64 `json:"hrOffItin"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// remains are loaded.
	HrOnIcao string `json:"hrOnICAO"`
	// Itinerary number that identifies where the remains are loaded.
	HrOnItin int64 `json:"hrOnItin"`
	// Remarks concerning the remains.
	HrRemarks string `json:"hrRemarks"`
	// Name of the deceased.
	Name string `json:"name"`
	// Rank of the deceased.
	Rank string `json:"rank"`
	// Name of the receiving agency or funeral home to which the remains are being
	// delivered.
	RecAgency string `json:"recAgency"`
	// Branch of service of the deceased.
	Service string `json:"service"`
	// Flag indicating if the remains are viewable.
	Viewable bool `json:"viewable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Container    resp.Field
		Escort       resp.Field
		HrEstArrTime resp.Field
		HrOffIcao    resp.Field
		HrOffItin    resp.Field
		HrOnIcao     resp.Field
		HrOnItin     resp.Field
		HrRemarks    resp.Field
		Name         resp.Field
		Rank         resp.Field
		RecAgency    resp.Field
		Service      resp.Field
		Viewable     resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanAbridgedAirLoadPlanHr) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanAbridgedAirLoadPlanHr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of cargo information located at the pallet positions associated with
// this load plan.
type AirloadplanAbridgedAirLoadPlanPalletDetail struct {
	// Category of special interest cargo.
	Category string `json:"category"`
	// Pallet position of the cargo.
	Pp string `json:"pp"`
	// Description of the cargo.
	PpDescription string `json:"ppDescription"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// cargo is unloaded.
	PpOffIcao string `json:"ppOffICAO"`
	// Number of pieces included in the Transportation Control Number (TCN).
	PpPieces int64 `json:"ppPieces"`
	// Remarks concerning the cargo at this pallet position.
	PpRemarks string `json:"ppRemarks"`
	// Transportation Control Number (TCN) of the cargo.
	PpTcn string `json:"ppTcn"`
	// Total weight of the cargo at this pallet position in kilograms.
	PpWeight float64 `json:"ppWeight"`
	// Flag indicating if this cargo is considered special interest.
	SpecialInterest bool `json:"specialInterest"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Category        resp.Field
		Pp              resp.Field
		PpDescription   resp.Field
		PpOffIcao       resp.Field
		PpPieces        resp.Field
		PpRemarks       resp.Field
		PpTcn           resp.Field
		PpWeight        resp.Field
		SpecialInterest resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanAbridgedAirLoadPlanPalletDetail) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanAbridgedAirLoadPlanPalletDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of passenger and cargo details associated with this load plan for
// this leg of the mission.
type AirloadplanAbridgedAirLoadPlanPaxCargo struct {
	// Number of ambulatory medical passengers in this group.
	AmbPax int64 `json:"ambPax"`
	// Number of patient attendant passengers in this group.
	AttPax int64 `json:"attPax"`
	// Number of space available passengers in this group.
	AvailablePax int64 `json:"availablePax"`
	// Weight of baggage in this group in kilograms.
	BagWeight float64 `json:"bagWeight"`
	// Number of civilian passengers in this group.
	CivPax int64 `json:"civPax"`
	// Number of distinguished visitor passengers in this group.
	DvPax int64 `json:"dvPax"`
	// Number of foreign national passengers in this group.
	FnPax int64 `json:"fnPax"`
	// Weight of cargo in this group in kilograms.
	GroupCargoWeight float64 `json:"groupCargoWeight"`
	// Describes the status or action needed for this group of passenger and cargo data
	// (e.g. ARRONBD, OFFTHIS, THROUGH, ONTHIS, DEPONBD, OFFNEXT).
	GroupType string `json:"groupType"`
	// Number of litter-bound passengers in this group.
	LitPax int64 `json:"litPax"`
	// Weight of mail in this group in kilograms.
	MailWeight float64 `json:"mailWeight"`
	// Number of cargo pallets in this group.
	NumPallet int64 `json:"numPallet"`
	// Weight of pallets, chains, and devices in this group in kilograms.
	PalletWeight float64 `json:"palletWeight"`
	// Weight of passengers in this group in kilograms.
	PaxWeight float64 `json:"paxWeight"`
	// Number of space required passengers in this group.
	RequiredPax int64 `json:"requiredPax"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AmbPax           resp.Field
		AttPax           resp.Field
		AvailablePax     resp.Field
		BagWeight        resp.Field
		CivPax           resp.Field
		DvPax            resp.Field
		FnPax            resp.Field
		GroupCargoWeight resp.Field
		GroupType        resp.Field
		LitPax           resp.Field
		MailWeight       resp.Field
		NumPallet        resp.Field
		PalletWeight     resp.Field
		PaxWeight        resp.Field
		RequiredPax      resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanAbridgedAirLoadPlanPaxCargo) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanAbridgedAirLoadPlanPaxCargo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of unit line number actuals associated with this load plan.
type AirloadplanAbridgedAirLoadPlanUlnActual struct {
	// Number of ambulatory patients associated with this load plan.
	NumAmbulatory int64 `json:"numAmbulatory"`
	// Number of attendants associated with this load plan.
	NumAttendant int64 `json:"numAttendant"`
	// Number of litter patients associated with this load plan.
	NumLitter int64 `json:"numLitter"`
	// Number of passengers associated with this load plan.
	NumPax int64 `json:"numPax"`
	// Identifier of the offload itinerary location.
	OffloadID int64 `json:"offloadId"`
	// Offload location code.
	OffloadLoCode string `json:"offloadLOCode"`
	// Identifier of the onload itinerary location.
	OnloadID int64 `json:"onloadId"`
	// Onload location code.
	OnloadLoCode string `json:"onloadLOCode"`
	// Identification number of the Operation Plan (OPLAN) associated with this load
	// plan.
	Oplan string `json:"oplan"`
	// Project name.
	ProjName string `json:"projName"`
	// Unit line number.
	Uln string `json:"uln"`
	// Total weight of all cargo items for this unit line number in kilograms.
	UlnCargoWeight float64 `json:"ulnCargoWeight"`
	// Remarks concerning these unit line number actuals.
	UlnRemarks string `json:"ulnRemarks"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		NumAmbulatory  resp.Field
		NumAttendant   resp.Field
		NumLitter      resp.Field
		NumPax         resp.Field
		OffloadID      resp.Field
		OffloadLoCode  resp.Field
		OnloadID       resp.Field
		OnloadLoCode   resp.Field
		Oplan          resp.Field
		ProjName       resp.Field
		Uln            resp.Field
		UlnCargoWeight resp.Field
		UlnRemarks     resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanAbridgedAirLoadPlanUlnActual) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanAbridgedAirLoadPlanUlnActual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information related to how an aircraft is loaded with cargo, equipment, and
// passengers.
type AirloadplanFull struct {
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
	DataMode AirloadplanFullDataMode `json:"dataMode,required"`
	// The current estimated time that the aircraft is planned to depart, in ISO 8601
	// UTC format with millisecond precision.
	EstDepTime time.Time `json:"estDepTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Allowable Cabin Load (ACL) onboard the aircraft. The maximum weight of
	// passengers, baggage, and cargo that can be safely transported in the aircraft
	// cabin, in kilograms.
	ACLOnboard float64 `json:"aclOnboard"`
	// Allowable Cabin Load (ACL) released this leg. The weight of passengers, baggage,
	// and cargo released from the aircraft cabin, in kilograms.
	ACLReleased float64 `json:"aclReleased"`
	// The Model Design Series designation of the aircraft supporting this load plan.
	AircraftMds string `json:"aircraftMDS"`
	// Collection of hazmat actuals associated with this load plan.
	AirLoadPlanHazmatActuals []AirloadplanFullAirLoadPlanHazmatActual `json:"airLoadPlanHazmatActuals"`
	// Collection of human remains transport information associated with this load
	// plan.
	AirLoadPlanHr []AirloadplanFullAirLoadPlanHr `json:"airLoadPlanHR"`
	// Collection of cargo information located at the pallet positions associated with
	// this load plan.
	AirLoadPlanPalletDetails []AirloadplanFullAirLoadPlanPalletDetail `json:"airLoadPlanPalletDetails"`
	// Collection of passenger and cargo details associated with this load plan for
	// this leg of the mission.
	AirLoadPlanPaxCargo []AirloadplanFullAirLoadPlanPaxCargo `json:"airLoadPlanPaxCargo"`
	// Collection of unit line number actuals associated with this load plan.
	AirLoadPlanUlnActuals []AirloadplanFullAirLoadPlanUlnActual `json:"airLoadPlanULNActuals"`
	// Optional identifier of arrival airfield with no International Civil Organization
	// (ICAO) code.
	ArrAirfield string `json:"arrAirfield"`
	// The arrival International Civil Organization (ICAO) code of the landing
	// airfield.
	ArrIcao string `json:"arrICAO"`
	// Time the loadmaster or boom operator is available for cargo loading/unloading,
	// in ISO 8601 UTC format with millisecond precision.
	AvailableTime time.Time `json:"availableTime" format:"date-time"`
	// The basic weight of the aircraft multiplied by the distance between the
	// reference datum and the aircraft's center of gravity, in Newton-meters.
	BasicMoment float64 `json:"basicMoment"`
	// The weight of the aircraft without passengers, cargo, equipment, or usable fuel,
	// in kilograms.
	BasicWeight float64 `json:"basicWeight"`
	// Time the cargo briefing was given to the loadmaster or boom operator, in ISO
	// 8601 UTC format with millisecond precision.
	BriefTime time.Time `json:"briefTime" format:"date-time"`
	// The call sign of the mission supporting this load plan.
	CallSign string `json:"callSign"`
	// Maximum fuselage station (FS) where cargo can be stored. FS is the distance from
	// the reference datum, in meters.
	CargoBayFsMax float64 `json:"cargoBayFSMax"`
	// Minimum fuselage station (FS) where cargo can be stored. FS is the distance from
	// the reference datum, in meters.
	CargoBayFsMin float64 `json:"cargoBayFSMin"`
	// Width of the cargo bay, in meters.
	CargoBayWidth float64 `json:"cargoBayWidth"`
	// The cargo configuration required for this leg (e.g. C-1, C-2, C-3, DV-1, DV-2,
	// AE-1, etc.). Configuration meanings are determined by the data source.
	CargoConfig string `json:"cargoConfig"`
	// The sum of cargo moments of all cargo on board the aircraft, in Newton-meters.
	// Each individual cargo moment is the weight of the cargo multiplied by the
	// distance between the reference datum and the cargo's center of gravity.
	CargoMoment float64 `json:"cargoMoment"`
	// Volume of cargo space in the aircraft, in cubic meters.
	CargoVolume float64 `json:"cargoVolume"`
	// The weight of the cargo on board the aircraft, in kilograms.
	CargoWeight float64 `json:"cargoWeight"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The number of crew members on the aircraft.
	CrewSize int64 `json:"crewSize"`
	// Optional identifier of departure airfield with no International Civil
	// Organization (ICAO) code.
	DepAirfield string `json:"depAirfield"`
	// The departure International Civil Organization (ICAO) code of the departure
	// airfield.
	DepIcao string `json:"depICAO"`
	// Description of the equipment configuration (e.g. Standard, Ferry, JBLM, CHS,
	// Combat, etc.). Configuration meanings are determined by the data source.
	EquipConfig string `json:"equipConfig"`
	// The current estimated time that the aircraft is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime time.Time `json:"estArrTime" format:"date-time"`
	// The estimated weight of usable fuel upon landing multiplied by the distance
	// between the reference datum and the fuel's center of gravity, in Newton-meters.
	EstLandingFuelMoment float64 `json:"estLandingFuelMoment"`
	// The estimated weight of usable fuel upon landing, in kilograms.
	EstLandingFuelWeight float64 `json:"estLandingFuelWeight"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// The fuel weight on board the aircraft multiplied by the distance between the
	// reference datum and the fuel's center of gravity, in Newton-meters.
	FuelMoment float64 `json:"fuelMoment"`
	// The weight of usable fuel on board the aircraft, in kilograms.
	FuelWeight float64 `json:"fuelWeight"`
	// The center of gravity of the aircraft using the gross weight and gross moment,
	// as a percentage of the mean aerodynamic chord (%MAC).
	GrossCg float64 `json:"grossCG"`
	// The sum of moments of all items making up the gross weight of the aircraft, in
	// Newton-meters.
	GrossMoment float64 `json:"grossMoment"`
	// The total weight of the aircraft at takeoff including passengers, cargo,
	// equipment, and usable fuel, in kilograms.
	GrossWeight float64 `json:"grossWeight"`
	// The UDL ID of the mission this record is associated with.
	IDMission string `json:"idMission"`
	// The UDL ID of the aircraft sortie this record is associated with.
	IDSortie string `json:"idSortie"`
	// The center of gravity of the aircraft using the landing weight and landing
	// moment, as a percentage of the mean aerodynamic chord (%MAC).
	LandingCg float64 `json:"landingCG"`
	// The sum of moments of all items making up the gross weight of the aircraft upon
	// landing, in Newton-meters.
	LandingMoment float64 `json:"landingMoment"`
	// The gross weight of the aircraft upon landing, in kilograms.
	LandingWeight float64 `json:"landingWeight"`
	// The leg number of the mission supporting this load plan.
	LegNum int64 `json:"legNum"`
	// Name of the loadmaster or boom operator who received the cargo briefing.
	LoadmasterName string `json:"loadmasterName"`
	// Rank of the loadmaster or boom operator overseeing cargo loading/unloading.
	LoadmasterRank string `json:"loadmasterRank"`
	// Remarks concerning this load plan.
	LoadRemarks string `json:"loadRemarks"`
	// The mission number of the mission supporting this load plan.
	MissionNumber string `json:"missionNumber"`
	// The operating weight of the aircraft multiplied by the distance between the
	// reference datum and the aircraft's center of gravity, in Newton-meters.
	OperatingMoment float64 `json:"operatingMoment"`
	// The basic weight of the aircraft including passengers and equipment, in
	// kilograms.
	OperatingWeight float64 `json:"operatingWeight"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Number of pallet positions on the aircraft.
	PpOnboard int64 `json:"ppOnboard"`
	// Number of pallet positions released this leg.
	PpReleased int64 `json:"ppReleased"`
	// Time the loadmaster or boom operator is scheduled to begin overseeing cargo
	// loading/unloading, in ISO 8601 UTC format with millisecond precision.
	SchedTime time.Time `json:"schedTime" format:"date-time"`
	// Number of passenger seats on the aircraft.
	SeatsOnboard int64 `json:"seatsOnboard"`
	// Number of passenger seats released this leg.
	SeatsReleased int64 `json:"seatsReleased"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The tail number of the aircraft supporting this load plan.
	TailNumber string `json:"tailNumber"`
	// Description of the fuel tank(s) configuration (e.g. ER, NON-ER, etc.).
	// Configuration meanings are determined by the data source.
	TankConfig string `json:"tankConfig"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Alphanumeric code that describes general cargo-related utilization and
	// characteristics for an itinerary point.
	UtilCode string `json:"utilCode"`
	// The center of gravity of the aircraft using the zero fuel weight and zero fuel
	// total moment, as a percentage of the mean aerodynamic chord (%MAC).
	ZeroFuelCg float64 `json:"zeroFuelCG"`
	// The zero fuel weight of the aircraft multiplied by the distance between the
	// reference datum and the aircraft's center of gravity, in Newton-meters.
	ZeroFuelMoment float64 `json:"zeroFuelMoment"`
	// The operating weight of the aircraft including cargo, mail, baggage, and
	// passengers, but without usable fuel, in kilograms.
	ZeroFuelWeight float64 `json:"zeroFuelWeight"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking    resp.Field
		DataMode                 resp.Field
		EstDepTime               resp.Field
		Source                   resp.Field
		ID                       resp.Field
		ACLOnboard               resp.Field
		ACLReleased              resp.Field
		AircraftMds              resp.Field
		AirLoadPlanHazmatActuals resp.Field
		AirLoadPlanHr            resp.Field
		AirLoadPlanPalletDetails resp.Field
		AirLoadPlanPaxCargo      resp.Field
		AirLoadPlanUlnActuals    resp.Field
		ArrAirfield              resp.Field
		ArrIcao                  resp.Field
		AvailableTime            resp.Field
		BasicMoment              resp.Field
		BasicWeight              resp.Field
		BriefTime                resp.Field
		CallSign                 resp.Field
		CargoBayFsMax            resp.Field
		CargoBayFsMin            resp.Field
		CargoBayWidth            resp.Field
		CargoConfig              resp.Field
		CargoMoment              resp.Field
		CargoVolume              resp.Field
		CargoWeight              resp.Field
		CreatedAt                resp.Field
		CreatedBy                resp.Field
		CrewSize                 resp.Field
		DepAirfield              resp.Field
		DepIcao                  resp.Field
		EquipConfig              resp.Field
		EstArrTime               resp.Field
		EstLandingFuelMoment     resp.Field
		EstLandingFuelWeight     resp.Field
		ExternalID               resp.Field
		FuelMoment               resp.Field
		FuelWeight               resp.Field
		GrossCg                  resp.Field
		GrossMoment              resp.Field
		GrossWeight              resp.Field
		IDMission                resp.Field
		IDSortie                 resp.Field
		LandingCg                resp.Field
		LandingMoment            resp.Field
		LandingWeight            resp.Field
		LegNum                   resp.Field
		LoadmasterName           resp.Field
		LoadmasterRank           resp.Field
		LoadRemarks              resp.Field
		MissionNumber            resp.Field
		OperatingMoment          resp.Field
		OperatingWeight          resp.Field
		Origin                   resp.Field
		OrigNetwork              resp.Field
		PpOnboard                resp.Field
		PpReleased               resp.Field
		SchedTime                resp.Field
		SeatsOnboard             resp.Field
		SeatsReleased            resp.Field
		SourceDl                 resp.Field
		TailNumber               resp.Field
		TankConfig               resp.Field
		UpdatedAt                resp.Field
		UpdatedBy                resp.Field
		UtilCode                 resp.Field
		ZeroFuelCg               resp.Field
		ZeroFuelMoment           resp.Field
		ZeroFuelWeight           resp.Field
		ExtraFields              map[string]resp.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanFull) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanFull) UnmarshalJSON(data []byte) error {
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
type AirloadplanFullDataMode string

const (
	AirloadplanFullDataModeReal      AirloadplanFullDataMode = "REAL"
	AirloadplanFullDataModeTest      AirloadplanFullDataMode = "TEST"
	AirloadplanFullDataModeSimulated AirloadplanFullDataMode = "SIMULATED"
	AirloadplanFullDataModeExercise  AirloadplanFullDataMode = "EXERCISE"
)

// Collection of hazmat actuals associated with this load plan.
type AirloadplanFullAirLoadPlanHazmatActual struct {
	// The Air Special Handling Code (ASHC) indicates the type of special handling
	// required for hazardous cargo.
	Ashc string `json:"ashc"`
	// Compatibility group code used to specify the controls for the transportation and
	// storage of hazardous materials according to the Hazardous Materials Regulations
	// issued by the U.S. Department of Transportation.
	Cgc string `json:"cgc"`
	// Class and division of the hazardous material according to the Hazardous
	// Materials Regulations issued by the U.S. Department of Transportation.
	ClassDiv string `json:"classDiv"`
	// Description of the hazardous item.
	HazDescription string `json:"hazDescription"`
	// Remarks concerning this hazardous material.
	HazmatRemarks string `json:"hazmatRemarks"`
	// United Nations number or North American number that identifies hazardous
	// materials according to the Hazardous Materials Regulations issued by the U.S.
	// Department of Transportation.
	HazNum string `json:"hazNum"`
	// Designates the type of hazmat number for the item (UN for United Nations or NA
	// for North American).
	HazNumType string `json:"hazNumType"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is unloaded.
	HazOffIcao string `json:"hazOffICAO"`
	// Itinerary number that identifies where the hazardous material is unloaded.
	HazOffItin int64 `json:"hazOffItin"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is loaded.
	HazOnIcao string `json:"hazOnICAO"`
	// Itinerary number that identifies where the hazardous material is loaded.
	HazOnItin int64 `json:"hazOnItin"`
	// Number of pieces of hazardous cargo.
	HazPieces int64 `json:"hazPieces"`
	// Transportation Control Number (TCN) of the hazardous item.
	HazTcn string `json:"hazTcn"`
	// Total weight of hazardous cargo, including non-explosive parts, in kilograms.
	HazWeight float64 `json:"hazWeight"`
	// United Nations proper shipping name of the hazardous material according to the
	// Hazardous Materials Regulations issued by the U.S. Department of Transportation.
	ItemName string `json:"itemName"`
	// Manufacturer's lot number for identification of the hazardous material.
	LotNum string `json:"lotNum"`
	// Net explosive weight of the hazardous material, in kilograms.
	NetExpWt float64 `json:"netExpWt"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Ashc           resp.Field
		Cgc            resp.Field
		ClassDiv       resp.Field
		HazDescription resp.Field
		HazmatRemarks  resp.Field
		HazNum         resp.Field
		HazNumType     resp.Field
		HazOffIcao     resp.Field
		HazOffItin     resp.Field
		HazOnIcao      resp.Field
		HazOnItin      resp.Field
		HazPieces      resp.Field
		HazTcn         resp.Field
		HazWeight      resp.Field
		ItemName       resp.Field
		LotNum         resp.Field
		NetExpWt       resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanFullAirLoadPlanHazmatActual) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanFullAirLoadPlanHazmatActual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of human remains transport information associated with this load
// plan.
type AirloadplanFullAirLoadPlanHr struct {
	// Type of transfer case used.
	Container string `json:"container"`
	// Name of the escort for the remains.
	Escort string `json:"escort"`
	// The current estimated time of arrival for the remains in ISO 8601 UTC format
	// with millisecond precision.
	HrEstArrTime time.Time `json:"hrEstArrTime" format:"date-time"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// remains are unloaded.
	HrOffIcao string `json:"hrOffICAO"`
	// Itinerary number that identifies where the remains are unloaded.
	HrOffItin int64 `json:"hrOffItin"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// remains are loaded.
	HrOnIcao string `json:"hrOnICAO"`
	// Itinerary number that identifies where the remains are loaded.
	HrOnItin int64 `json:"hrOnItin"`
	// Remarks concerning the remains.
	HrRemarks string `json:"hrRemarks"`
	// Name of the deceased.
	Name string `json:"name"`
	// Rank of the deceased.
	Rank string `json:"rank"`
	// Name of the receiving agency or funeral home to which the remains are being
	// delivered.
	RecAgency string `json:"recAgency"`
	// Branch of service of the deceased.
	Service string `json:"service"`
	// Flag indicating if the remains are viewable.
	Viewable bool `json:"viewable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Container    resp.Field
		Escort       resp.Field
		HrEstArrTime resp.Field
		HrOffIcao    resp.Field
		HrOffItin    resp.Field
		HrOnIcao     resp.Field
		HrOnItin     resp.Field
		HrRemarks    resp.Field
		Name         resp.Field
		Rank         resp.Field
		RecAgency    resp.Field
		Service      resp.Field
		Viewable     resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanFullAirLoadPlanHr) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanFullAirLoadPlanHr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of cargo information located at the pallet positions associated with
// this load plan.
type AirloadplanFullAirLoadPlanPalletDetail struct {
	// Category of special interest cargo.
	Category string `json:"category"`
	// Pallet position of the cargo.
	Pp string `json:"pp"`
	// Description of the cargo.
	PpDescription string `json:"ppDescription"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// cargo is unloaded.
	PpOffIcao string `json:"ppOffICAO"`
	// Number of pieces included in the Transportation Control Number (TCN).
	PpPieces int64 `json:"ppPieces"`
	// Remarks concerning the cargo at this pallet position.
	PpRemarks string `json:"ppRemarks"`
	// Transportation Control Number (TCN) of the cargo.
	PpTcn string `json:"ppTcn"`
	// Total weight of the cargo at this pallet position in kilograms.
	PpWeight float64 `json:"ppWeight"`
	// Flag indicating if this cargo is considered special interest.
	SpecialInterest bool `json:"specialInterest"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Category        resp.Field
		Pp              resp.Field
		PpDescription   resp.Field
		PpOffIcao       resp.Field
		PpPieces        resp.Field
		PpRemarks       resp.Field
		PpTcn           resp.Field
		PpWeight        resp.Field
		SpecialInterest resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanFullAirLoadPlanPalletDetail) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanFullAirLoadPlanPalletDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of passenger and cargo details associated with this load plan for
// this leg of the mission.
type AirloadplanFullAirLoadPlanPaxCargo struct {
	// Number of ambulatory medical passengers in this group.
	AmbPax int64 `json:"ambPax"`
	// Number of patient attendant passengers in this group.
	AttPax int64 `json:"attPax"`
	// Number of space available passengers in this group.
	AvailablePax int64 `json:"availablePax"`
	// Weight of baggage in this group in kilograms.
	BagWeight float64 `json:"bagWeight"`
	// Number of civilian passengers in this group.
	CivPax int64 `json:"civPax"`
	// Number of distinguished visitor passengers in this group.
	DvPax int64 `json:"dvPax"`
	// Number of foreign national passengers in this group.
	FnPax int64 `json:"fnPax"`
	// Weight of cargo in this group in kilograms.
	GroupCargoWeight float64 `json:"groupCargoWeight"`
	// Describes the status or action needed for this group of passenger and cargo data
	// (e.g. ARRONBD, OFFTHIS, THROUGH, ONTHIS, DEPONBD, OFFNEXT).
	GroupType string `json:"groupType"`
	// Number of litter-bound passengers in this group.
	LitPax int64 `json:"litPax"`
	// Weight of mail in this group in kilograms.
	MailWeight float64 `json:"mailWeight"`
	// Number of cargo pallets in this group.
	NumPallet int64 `json:"numPallet"`
	// Weight of pallets, chains, and devices in this group in kilograms.
	PalletWeight float64 `json:"palletWeight"`
	// Weight of passengers in this group in kilograms.
	PaxWeight float64 `json:"paxWeight"`
	// Number of space required passengers in this group.
	RequiredPax int64 `json:"requiredPax"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AmbPax           resp.Field
		AttPax           resp.Field
		AvailablePax     resp.Field
		BagWeight        resp.Field
		CivPax           resp.Field
		DvPax            resp.Field
		FnPax            resp.Field
		GroupCargoWeight resp.Field
		GroupType        resp.Field
		LitPax           resp.Field
		MailWeight       resp.Field
		NumPallet        resp.Field
		PalletWeight     resp.Field
		PaxWeight        resp.Field
		RequiredPax      resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanFullAirLoadPlanPaxCargo) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanFullAirLoadPlanPaxCargo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of unit line number actuals associated with this load plan.
type AirloadplanFullAirLoadPlanUlnActual struct {
	// Number of ambulatory patients associated with this load plan.
	NumAmbulatory int64 `json:"numAmbulatory"`
	// Number of attendants associated with this load plan.
	NumAttendant int64 `json:"numAttendant"`
	// Number of litter patients associated with this load plan.
	NumLitter int64 `json:"numLitter"`
	// Number of passengers associated with this load plan.
	NumPax int64 `json:"numPax"`
	// Identifier of the offload itinerary location.
	OffloadID int64 `json:"offloadId"`
	// Offload location code.
	OffloadLoCode string `json:"offloadLOCode"`
	// Identifier of the onload itinerary location.
	OnloadID int64 `json:"onloadId"`
	// Onload location code.
	OnloadLoCode string `json:"onloadLOCode"`
	// Identification number of the Operation Plan (OPLAN) associated with this load
	// plan.
	Oplan string `json:"oplan"`
	// Project name.
	ProjName string `json:"projName"`
	// Unit line number.
	Uln string `json:"uln"`
	// Total weight of all cargo items for this unit line number in kilograms.
	UlnCargoWeight float64 `json:"ulnCargoWeight"`
	// Remarks concerning these unit line number actuals.
	UlnRemarks string `json:"ulnRemarks"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		NumAmbulatory  resp.Field
		NumAttendant   resp.Field
		NumLitter      resp.Field
		NumPax         resp.Field
		OffloadID      resp.Field
		OffloadLoCode  resp.Field
		OnloadID       resp.Field
		OnloadLoCode   resp.Field
		Oplan          resp.Field
		ProjName       resp.Field
		Uln            resp.Field
		UlnCargoWeight resp.Field
		UlnRemarks     resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirloadplanFullAirLoadPlanUlnActual) RawJSON() string { return r.JSON.raw }
func (r *AirloadplanFullAirLoadPlanUlnActual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirLoadPlanNewParams struct {
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
	DataMode AirLoadPlanNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The current estimated time that the aircraft is planned to depart, in ISO 8601
	// UTC format with millisecond precision.
	EstDepTime time.Time `json:"estDepTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Allowable Cabin Load (ACL) onboard the aircraft. The maximum weight of
	// passengers, baggage, and cargo that can be safely transported in the aircraft
	// cabin, in kilograms.
	ACLOnboard param.Opt[float64] `json:"aclOnboard,omitzero"`
	// Allowable Cabin Load (ACL) released this leg. The weight of passengers, baggage,
	// and cargo released from the aircraft cabin, in kilograms.
	ACLReleased param.Opt[float64] `json:"aclReleased,omitzero"`
	// The Model Design Series designation of the aircraft supporting this load plan.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Optional identifier of arrival airfield with no International Civil Organization
	// (ICAO) code.
	ArrAirfield param.Opt[string] `json:"arrAirfield,omitzero"`
	// The arrival International Civil Organization (ICAO) code of the landing
	// airfield.
	ArrIcao param.Opt[string] `json:"arrICAO,omitzero"`
	// Time the loadmaster or boom operator is available for cargo loading/unloading,
	// in ISO 8601 UTC format with millisecond precision.
	AvailableTime param.Opt[time.Time] `json:"availableTime,omitzero" format:"date-time"`
	// The basic weight of the aircraft multiplied by the distance between the
	// reference datum and the aircraft's center of gravity, in Newton-meters.
	BasicMoment param.Opt[float64] `json:"basicMoment,omitzero"`
	// The weight of the aircraft without passengers, cargo, equipment, or usable fuel,
	// in kilograms.
	BasicWeight param.Opt[float64] `json:"basicWeight,omitzero"`
	// Time the cargo briefing was given to the loadmaster or boom operator, in ISO
	// 8601 UTC format with millisecond precision.
	BriefTime param.Opt[time.Time] `json:"briefTime,omitzero" format:"date-time"`
	// The call sign of the mission supporting this load plan.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// Maximum fuselage station (FS) where cargo can be stored. FS is the distance from
	// the reference datum, in meters.
	CargoBayFsMax param.Opt[float64] `json:"cargoBayFSMax,omitzero"`
	// Minimum fuselage station (FS) where cargo can be stored. FS is the distance from
	// the reference datum, in meters.
	CargoBayFsMin param.Opt[float64] `json:"cargoBayFSMin,omitzero"`
	// Width of the cargo bay, in meters.
	CargoBayWidth param.Opt[float64] `json:"cargoBayWidth,omitzero"`
	// The cargo configuration required for this leg (e.g. C-1, C-2, C-3, DV-1, DV-2,
	// AE-1, etc.). Configuration meanings are determined by the data source.
	CargoConfig param.Opt[string] `json:"cargoConfig,omitzero"`
	// The sum of cargo moments of all cargo on board the aircraft, in Newton-meters.
	// Each individual cargo moment is the weight of the cargo multiplied by the
	// distance between the reference datum and the cargo's center of gravity.
	CargoMoment param.Opt[float64] `json:"cargoMoment,omitzero"`
	// Volume of cargo space in the aircraft, in cubic meters.
	CargoVolume param.Opt[float64] `json:"cargoVolume,omitzero"`
	// The weight of the cargo on board the aircraft, in kilograms.
	CargoWeight param.Opt[float64] `json:"cargoWeight,omitzero"`
	// The number of crew members on the aircraft.
	CrewSize param.Opt[int64] `json:"crewSize,omitzero"`
	// Optional identifier of departure airfield with no International Civil
	// Organization (ICAO) code.
	DepAirfield param.Opt[string] `json:"depAirfield,omitzero"`
	// The departure International Civil Organization (ICAO) code of the departure
	// airfield.
	DepIcao param.Opt[string] `json:"depICAO,omitzero"`
	// Description of the equipment configuration (e.g. Standard, Ferry, JBLM, CHS,
	// Combat, etc.). Configuration meanings are determined by the data source.
	EquipConfig param.Opt[string] `json:"equipConfig,omitzero"`
	// The current estimated time that the aircraft is planned to arrive, in ISO 8601
	// UTC format with millisecond precision.
	EstArrTime param.Opt[time.Time] `json:"estArrTime,omitzero" format:"date-time"`
	// The estimated weight of usable fuel upon landing multiplied by the distance
	// between the reference datum and the fuel's center of gravity, in Newton-meters.
	EstLandingFuelMoment param.Opt[float64] `json:"estLandingFuelMoment,omitzero"`
	// The estimated weight of usable fuel upon landing, in kilograms.
	EstLandingFuelWeight param.Opt[float64] `json:"estLandingFuelWeight,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The fuel weight on board the aircraft multiplied by the distance between the
	// reference datum and the fuel's center of gravity, in Newton-meters.
	FuelMoment param.Opt[float64] `json:"fuelMoment,omitzero"`
	// The weight of usable fuel on board the aircraft, in kilograms.
	FuelWeight param.Opt[float64] `json:"fuelWeight,omitzero"`
	// The center of gravity of the aircraft using the gross weight and gross moment,
	// as a percentage of the mean aerodynamic chord (%MAC).
	GrossCg param.Opt[float64] `json:"grossCG,omitzero"`
	// The sum of moments of all items making up the gross weight of the aircraft, in
	// Newton-meters.
	GrossMoment param.Opt[float64] `json:"grossMoment,omitzero"`
	// The total weight of the aircraft at takeoff including passengers, cargo,
	// equipment, and usable fuel, in kilograms.
	GrossWeight param.Opt[float64] `json:"grossWeight,omitzero"`
	// The UDL ID of the mission this record is associated with.
	IDMission param.Opt[string] `json:"idMission,omitzero"`
	// The UDL ID of the aircraft sortie this record is associated with.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// The center of gravity of the aircraft using the landing weight and landing
	// moment, as a percentage of the mean aerodynamic chord (%MAC).
	LandingCg param.Opt[float64] `json:"landingCG,omitzero"`
	// The sum of moments of all items making up the gross weight of the aircraft upon
	// landing, in Newton-meters.
	LandingMoment param.Opt[float64] `json:"landingMoment,omitzero"`
	// The gross weight of the aircraft upon landing, in kilograms.
	LandingWeight param.Opt[float64] `json:"landingWeight,omitzero"`
	// The leg number of the mission supporting this load plan.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// Name of the loadmaster or boom operator who received the cargo briefing.
	LoadmasterName param.Opt[string] `json:"loadmasterName,omitzero"`
	// Rank of the loadmaster or boom operator overseeing cargo loading/unloading.
	LoadmasterRank param.Opt[string] `json:"loadmasterRank,omitzero"`
	// Remarks concerning this load plan.
	LoadRemarks param.Opt[string] `json:"loadRemarks,omitzero"`
	// The mission number of the mission supporting this load plan.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// The operating weight of the aircraft multiplied by the distance between the
	// reference datum and the aircraft's center of gravity, in Newton-meters.
	OperatingMoment param.Opt[float64] `json:"operatingMoment,omitzero"`
	// The basic weight of the aircraft including passengers and equipment, in
	// kilograms.
	OperatingWeight param.Opt[float64] `json:"operatingWeight,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Number of pallet positions on the aircraft.
	PpOnboard param.Opt[int64] `json:"ppOnboard,omitzero"`
	// Number of pallet positions released this leg.
	PpReleased param.Opt[int64] `json:"ppReleased,omitzero"`
	// Time the loadmaster or boom operator is scheduled to begin overseeing cargo
	// loading/unloading, in ISO 8601 UTC format with millisecond precision.
	SchedTime param.Opt[time.Time] `json:"schedTime,omitzero" format:"date-time"`
	// Number of passenger seats on the aircraft.
	SeatsOnboard param.Opt[int64] `json:"seatsOnboard,omitzero"`
	// Number of passenger seats released this leg.
	SeatsReleased param.Opt[int64] `json:"seatsReleased,omitzero"`
	// The tail number of the aircraft supporting this load plan.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// Description of the fuel tank(s) configuration (e.g. ER, NON-ER, etc.).
	// Configuration meanings are determined by the data source.
	TankConfig param.Opt[string] `json:"tankConfig,omitzero"`
	// Alphanumeric code that describes general cargo-related utilization and
	// characteristics for an itinerary point.
	UtilCode param.Opt[string] `json:"utilCode,omitzero"`
	// The center of gravity of the aircraft using the zero fuel weight and zero fuel
	// total moment, as a percentage of the mean aerodynamic chord (%MAC).
	ZeroFuelCg param.Opt[float64] `json:"zeroFuelCG,omitzero"`
	// The zero fuel weight of the aircraft multiplied by the distance between the
	// reference datum and the aircraft's center of gravity, in Newton-meters.
	ZeroFuelMoment param.Opt[float64] `json:"zeroFuelMoment,omitzero"`
	// The operating weight of the aircraft including cargo, mail, baggage, and
	// passengers, but without usable fuel, in kilograms.
	ZeroFuelWeight param.Opt[float64] `json:"zeroFuelWeight,omitzero"`
	// Collection of hazmat actuals associated with this load plan.
	AirLoadPlanHazmatActuals []AirLoadPlanNewParamsAirLoadPlanHazmatActual `json:"airLoadPlanHazmatActuals,omitzero"`
	// Collection of human remains transport information associated with this load
	// plan.
	AirLoadPlanHr []AirLoadPlanNewParamsAirLoadPlanHr `json:"airLoadPlanHR,omitzero"`
	// Collection of cargo information located at the pallet positions associated with
	// this load plan.
	AirLoadPlanPalletDetails []AirLoadPlanNewParamsAirLoadPlanPalletDetail `json:"airLoadPlanPalletDetails,omitzero"`
	// Collection of passenger and cargo details associated with this load plan for
	// this leg of the mission.
	AirLoadPlanPaxCargo []AirLoadPlanNewParamsAirLoadPlanPaxCargo `json:"airLoadPlanPaxCargo,omitzero"`
	// Collection of unit line number actuals associated with this load plan.
	AirLoadPlanUlnActuals []AirLoadPlanNewParamsAirLoadPlanUlnActual `json:"airLoadPlanULNActuals,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r AirLoadPlanNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AirLoadPlanNewParams
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
type AirLoadPlanNewParamsDataMode string

const (
	AirLoadPlanNewParamsDataModeReal      AirLoadPlanNewParamsDataMode = "REAL"
	AirLoadPlanNewParamsDataModeTest      AirLoadPlanNewParamsDataMode = "TEST"
	AirLoadPlanNewParamsDataModeSimulated AirLoadPlanNewParamsDataMode = "SIMULATED"
	AirLoadPlanNewParamsDataModeExercise  AirLoadPlanNewParamsDataMode = "EXERCISE"
)

// Collection of hazmat actuals associated with this load plan.
type AirLoadPlanNewParamsAirLoadPlanHazmatActual struct {
	// The Air Special Handling Code (ASHC) indicates the type of special handling
	// required for hazardous cargo.
	Ashc param.Opt[string] `json:"ashc,omitzero"`
	// Compatibility group code used to specify the controls for the transportation and
	// storage of hazardous materials according to the Hazardous Materials Regulations
	// issued by the U.S. Department of Transportation.
	Cgc param.Opt[string] `json:"cgc,omitzero"`
	// Class and division of the hazardous material according to the Hazardous
	// Materials Regulations issued by the U.S. Department of Transportation.
	ClassDiv param.Opt[string] `json:"classDiv,omitzero"`
	// Description of the hazardous item.
	HazDescription param.Opt[string] `json:"hazDescription,omitzero"`
	// Remarks concerning this hazardous material.
	HazmatRemarks param.Opt[string] `json:"hazmatRemarks,omitzero"`
	// United Nations number or North American number that identifies hazardous
	// materials according to the Hazardous Materials Regulations issued by the U.S.
	// Department of Transportation.
	HazNum param.Opt[string] `json:"hazNum,omitzero"`
	// Designates the type of hazmat number for the item (UN for United Nations or NA
	// for North American).
	HazNumType param.Opt[string] `json:"hazNumType,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is unloaded.
	HazOffIcao param.Opt[string] `json:"hazOffICAO,omitzero"`
	// Itinerary number that identifies where the hazardous material is unloaded.
	HazOffItin param.Opt[int64] `json:"hazOffItin,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// hazardous material is loaded.
	HazOnIcao param.Opt[string] `json:"hazOnICAO,omitzero"`
	// Itinerary number that identifies where the hazardous material is loaded.
	HazOnItin param.Opt[int64] `json:"hazOnItin,omitzero"`
	// Number of pieces of hazardous cargo.
	HazPieces param.Opt[int64] `json:"hazPieces,omitzero"`
	// Transportation Control Number (TCN) of the hazardous item.
	HazTcn param.Opt[string] `json:"hazTcn,omitzero"`
	// Total weight of hazardous cargo, including non-explosive parts, in kilograms.
	HazWeight param.Opt[float64] `json:"hazWeight,omitzero"`
	// United Nations proper shipping name of the hazardous material according to the
	// Hazardous Materials Regulations issued by the U.S. Department of Transportation.
	ItemName param.Opt[string] `json:"itemName,omitzero"`
	// Manufacturer's lot number for identification of the hazardous material.
	LotNum param.Opt[string] `json:"lotNum,omitzero"`
	// Net explosive weight of the hazardous material, in kilograms.
	NetExpWt param.Opt[float64] `json:"netExpWt,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanNewParamsAirLoadPlanHazmatActual) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AirLoadPlanNewParamsAirLoadPlanHazmatActual) MarshalJSON() (data []byte, err error) {
	type shadow AirLoadPlanNewParamsAirLoadPlanHazmatActual
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of human remains transport information associated with this load
// plan.
type AirLoadPlanNewParamsAirLoadPlanHr struct {
	// Type of transfer case used.
	Container param.Opt[string] `json:"container,omitzero"`
	// Name of the escort for the remains.
	Escort param.Opt[string] `json:"escort,omitzero"`
	// The current estimated time of arrival for the remains in ISO 8601 UTC format
	// with millisecond precision.
	HrEstArrTime param.Opt[time.Time] `json:"hrEstArrTime,omitzero" format:"date-time"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// remains are unloaded.
	HrOffIcao param.Opt[string] `json:"hrOffICAO,omitzero"`
	// Itinerary number that identifies where the remains are unloaded.
	HrOffItin param.Opt[int64] `json:"hrOffItin,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// remains are loaded.
	HrOnIcao param.Opt[string] `json:"hrOnICAO,omitzero"`
	// Itinerary number that identifies where the remains are loaded.
	HrOnItin param.Opt[int64] `json:"hrOnItin,omitzero"`
	// Remarks concerning the remains.
	HrRemarks param.Opt[string] `json:"hrRemarks,omitzero"`
	// Name of the deceased.
	Name param.Opt[string] `json:"name,omitzero"`
	// Rank of the deceased.
	Rank param.Opt[string] `json:"rank,omitzero"`
	// Name of the receiving agency or funeral home to which the remains are being
	// delivered.
	RecAgency param.Opt[string] `json:"recAgency,omitzero"`
	// Branch of service of the deceased.
	Service param.Opt[string] `json:"service,omitzero"`
	// Flag indicating if the remains are viewable.
	Viewable param.Opt[bool] `json:"viewable,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanNewParamsAirLoadPlanHr) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AirLoadPlanNewParamsAirLoadPlanHr) MarshalJSON() (data []byte, err error) {
	type shadow AirLoadPlanNewParamsAirLoadPlanHr
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of cargo information located at the pallet positions associated with
// this load plan.
type AirLoadPlanNewParamsAirLoadPlanPalletDetail struct {
	// Category of special interest cargo.
	Category param.Opt[string] `json:"category,omitzero"`
	// Pallet position of the cargo.
	Pp param.Opt[string] `json:"pp,omitzero"`
	// Description of the cargo.
	PpDescription param.Opt[string] `json:"ppDescription,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the site where the
	// cargo is unloaded.
	PpOffIcao param.Opt[string] `json:"ppOffICAO,omitzero"`
	// Number of pieces included in the Transportation Control Number (TCN).
	PpPieces param.Opt[int64] `json:"ppPieces,omitzero"`
	// Remarks concerning the cargo at this pallet position.
	PpRemarks param.Opt[string] `json:"ppRemarks,omitzero"`
	// Transportation Control Number (TCN) of the cargo.
	PpTcn param.Opt[string] `json:"ppTcn,omitzero"`
	// Total weight of the cargo at this pallet position in kilograms.
	PpWeight param.Opt[float64] `json:"ppWeight,omitzero"`
	// Flag indicating if this cargo is considered special interest.
	SpecialInterest param.Opt[bool] `json:"specialInterest,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanNewParamsAirLoadPlanPalletDetail) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AirLoadPlanNewParamsAirLoadPlanPalletDetail) MarshalJSON() (data []byte, err error) {
	type shadow AirLoadPlanNewParamsAirLoadPlanPalletDetail
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of passenger and cargo details associated with this load plan for
// this leg of the mission.
type AirLoadPlanNewParamsAirLoadPlanPaxCargo struct {
	// Number of ambulatory medical passengers in this group.
	AmbPax param.Opt[int64] `json:"ambPax,omitzero"`
	// Number of patient attendant passengers in this group.
	AttPax param.Opt[int64] `json:"attPax,omitzero"`
	// Number of space available passengers in this group.
	AvailablePax param.Opt[int64] `json:"availablePax,omitzero"`
	// Weight of baggage in this group in kilograms.
	BagWeight param.Opt[float64] `json:"bagWeight,omitzero"`
	// Number of civilian passengers in this group.
	CivPax param.Opt[int64] `json:"civPax,omitzero"`
	// Number of distinguished visitor passengers in this group.
	DvPax param.Opt[int64] `json:"dvPax,omitzero"`
	// Number of foreign national passengers in this group.
	FnPax param.Opt[int64] `json:"fnPax,omitzero"`
	// Weight of cargo in this group in kilograms.
	GroupCargoWeight param.Opt[float64] `json:"groupCargoWeight,omitzero"`
	// Describes the status or action needed for this group of passenger and cargo data
	// (e.g. ARRONBD, OFFTHIS, THROUGH, ONTHIS, DEPONBD, OFFNEXT).
	GroupType param.Opt[string] `json:"groupType,omitzero"`
	// Number of litter-bound passengers in this group.
	LitPax param.Opt[int64] `json:"litPax,omitzero"`
	// Weight of mail in this group in kilograms.
	MailWeight param.Opt[float64] `json:"mailWeight,omitzero"`
	// Number of cargo pallets in this group.
	NumPallet param.Opt[int64] `json:"numPallet,omitzero"`
	// Weight of pallets, chains, and devices in this group in kilograms.
	PalletWeight param.Opt[float64] `json:"palletWeight,omitzero"`
	// Weight of passengers in this group in kilograms.
	PaxWeight param.Opt[float64] `json:"paxWeight,omitzero"`
	// Number of space required passengers in this group.
	RequiredPax param.Opt[int64] `json:"requiredPax,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanNewParamsAirLoadPlanPaxCargo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AirLoadPlanNewParamsAirLoadPlanPaxCargo) MarshalJSON() (data []byte, err error) {
	type shadow AirLoadPlanNewParamsAirLoadPlanPaxCargo
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of unit line number actuals associated with this load plan.
type AirLoadPlanNewParamsAirLoadPlanUlnActual struct {
	// Number of ambulatory patients associated with this load plan.
	NumAmbulatory param.Opt[int64] `json:"numAmbulatory,omitzero"`
	// Number of attendants associated with this load plan.
	NumAttendant param.Opt[int64] `json:"numAttendant,omitzero"`
	// Number of litter patients associated with this load plan.
	NumLitter param.Opt[int64] `json:"numLitter,omitzero"`
	// Number of passengers associated with this load plan.
	NumPax param.Opt[int64] `json:"numPax,omitzero"`
	// Identifier of the offload itinerary location.
	OffloadID param.Opt[int64] `json:"offloadId,omitzero"`
	// Offload location code.
	OffloadLoCode param.Opt[string] `json:"offloadLOCode,omitzero"`
	// Identifier of the onload itinerary location.
	OnloadID param.Opt[int64] `json:"onloadId,omitzero"`
	// Onload location code.
	OnloadLoCode param.Opt[string] `json:"onloadLOCode,omitzero"`
	// Identification number of the Operation Plan (OPLAN) associated with this load
	// plan.
	Oplan param.Opt[string] `json:"oplan,omitzero"`
	// Project name.
	ProjName param.Opt[string] `json:"projName,omitzero"`
	// Unit line number.
	Uln param.Opt[string] `json:"uln,omitzero"`
	// Total weight of all cargo items for this unit line number in kilograms.
	UlnCargoWeight param.Opt[float64] `json:"ulnCargoWeight,omitzero"`
	// Remarks concerning these unit line number actuals.
	UlnRemarks param.Opt[string] `json:"ulnRemarks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanNewParamsAirLoadPlanUlnActual) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AirLoadPlanNewParamsAirLoadPlanUlnActual) MarshalJSON() (data []byte, err error) {
	type shadow AirLoadPlanNewParamsAirLoadPlanUlnActual
	return param.MarshalObject(r, (*shadow)(&r))
}

type AirLoadPlanGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirLoadPlanGetParams]'s query parameters as `url.Values`.
func (r AirLoadPlanGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirLoadPlanListParams struct {
	// The current estimated time that the aircraft is planned to depart, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EstDepTime  time.Time        `query:"estDepTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirLoadPlanListParams]'s query parameters as `url.Values`.
func (r AirLoadPlanListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirLoadPlanCountParams struct {
	// The current estimated time that the aircraft is planned to depart, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EstDepTime  time.Time        `query:"estDepTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirLoadPlanCountParams]'s query parameters as `url.Values`.
func (r AirLoadPlanCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirLoadPlanTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The current estimated time that the aircraft is planned to depart, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EstDepTime  time.Time        `query:"estDepTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirLoadPlanTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirLoadPlanTupleParams]'s query parameters as `url.Values`.
func (r AirLoadPlanTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
