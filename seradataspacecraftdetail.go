// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
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
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// SeradataSpacecraftDetailService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeradataSpacecraftDetailService] method instead.
type SeradataSpacecraftDetailService struct {
	Options []option.RequestOption
}

// NewSeradataSpacecraftDetailService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSeradataSpacecraftDetailService(opts ...option.RequestOption) (r SeradataSpacecraftDetailService) {
	r = SeradataSpacecraftDetailService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataSpacecraftDetails as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeradataSpacecraftDetailService) New(ctx context.Context, body SeradataSpacecraftDetailNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/seradataspacecraftdetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataSpacecraftDetails. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SeradataSpacecraftDetailService) Update(ctx context.Context, id string, body SeradataSpacecraftDetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataspacecraftdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SeradataSpacecraftDetailService) List(ctx context.Context, query SeradataSpacecraftDetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeradataSpacecraftDetailListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/seradataspacecraftdetails"
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
func (r *SeradataSpacecraftDetailService) ListAutoPaging(ctx context.Context, query SeradataSpacecraftDetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeradataSpacecraftDetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataSpacecraftDetails specified by the passed
// ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeradataSpacecraftDetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataspacecraftdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SeradataSpacecraftDetailService) Count(ctx context.Context, query SeradataSpacecraftDetailCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradataspacecraftdetails/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataSpacecraftDetails by its unique ID
// passed as a path parameter.
func (r *SeradataSpacecraftDetailService) Get(ctx context.Context, id string, query SeradataSpacecraftDetailGetParams, opts ...option.RequestOption) (res *SeradataSpacecraftDetailGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataspacecraftdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SeradataSpacecraftDetailService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SeradataSpacecraftDetailQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/seradataspacecraftdetails/queryhelp"
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
func (r *SeradataSpacecraftDetailService) Tuple(ctx context.Context, query SeradataSpacecraftDetailTupleParams, opts ...option.RequestOption) (res *[]SeradataSpacecraftDetailTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/seradataspacecraftdetails/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// On-orbit spacecraft details compiled by Seradata for a particular satellite.
type SeradataSpacecraftDetailListResponse struct {
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
	DataMode SeradataSpacecraftDetailListResponseDataMode `json:"dataMode,required"`
	// Spacecraft name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Spacecraft additional missions and groups.
	AdditionalMissionsGroups string `json:"additionalMissionsGroups"`
	// Spacecraft latest altitude in km.
	Altitude float64 `json:"altitude"`
	// Annual insured depreciaion factor as a percent fraction.
	AnnualInsuredDepreciationFactor float64 `json:"annualInsuredDepreciationFactor"`
	// Boolean indicating if the spacecraft annualInsuredDepreciationFactor is
	// estimated.
	AnnualInsuredDepreciationFactorEstimated bool `json:"annualInsuredDepreciationFactorEstimated"`
	// Apogee in km.
	Apogee float64 `json:"apogee"`
	// Spacecraft Bus ID.
	BusID string `json:"busId"`
	// Total capability lost as a percent fraction.
	CapabilityLost float64 `json:"capabilityLost"`
	// Total capacity lost as a percent fraction.
	CapacityLost float64 `json:"capacityLost"`
	// NORAD satellite number if available.
	CatalogNumber int64 `json:"catalogNumber"`
	// Spacecraft collision risk 1cm sqm latest.
	CollisionRiskCm float64 `json:"collisionRiskCM"`
	// Spacecraft collision risk 1mm sqm latest.
	CollisionRiskMm float64 `json:"collisionRiskMM"`
	// Boolean indicating if the spacecraft combined new cost is estimated.
	CombinedCostEstimated bool `json:"combinedCostEstimated"`
	// Combined cost of spacecraft at new in M USD.
	CombinedNewCost float64 `json:"combinedNewCost"`
	// Boolean indicating if the launch was commercial.
	CommercialLaunch bool `json:"commercialLaunch"`
	// Spacecraft constellation.
	Constellation string `json:"constellation"`
	// Boolean indicating if the spacecraft cost is estimated.
	CostEstimated bool `json:"costEstimated"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Cubesat dispenser type.
	CubesatDispenserType string `json:"cubesatDispenserType"`
	// Current age in years.
	CurrentAge float64 `json:"currentAge"`
	// Spacecraft date of observation.
	DateOfObservation time.Time `json:"dateOfObservation" format:"date-time"`
	// Description associated with the spacecraft.
	Description string `json:"description"`
	// Spacecraft design life in days.
	DesignLife int64 `json:"designLife"`
	// Mass dry in kg.
	DryMass float64 `json:"dryMass"`
	// Spacecraft expected life in days.
	ExpectedLife int64 `json:"expectedLife"`
	// WGS84 longitude of the spacecraft’s latest GEO position, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	GeoPosition float64 `json:"geoPosition"`
	// UUID of the parent Onorbit record, if available.
	IDOnOrbit string `json:"idOnOrbit"`
	// Seradata provided inclination in degrees.
	Inclination float64 `json:"inclination"`
	// Spacecraft total insurance losses as a fraction.
	InsuranceLossesTotal float64 `json:"insuranceLossesTotal"`
	// Insurance notes for the spacecraft.
	InsuranceNotes string `json:"insuranceNotes"`
	// Insurance premium at launch in M USD.
	InsurancePremiumAtLaunch float64 `json:"insurancePremiumAtLaunch"`
	// Boolean indicating if the spacecraft insurancePremiumAtLaunch is estimated.
	InsurancePremiumAtLaunchEstimated bool `json:"insurancePremiumAtLaunchEstimated"`
	// Boolean indicating if the spacecraft was insured at launch.
	InsuredAtLaunch bool `json:"insuredAtLaunch"`
	// Insured value of spacecraft at launch in M USD.
	InsuredValueAtLaunch float64 `json:"insuredValueAtLaunch"`
	// Boolean indicating if the spacecraft insured value at launch is estimated.
	InsuredValueLaunchEstimated bool `json:"insuredValueLaunchEstimated"`
	// Seradata international number.
	IntlNumber string `json:"intlNumber"`
	// Spacecraft latest latitude in degrees.
	Lat float64 `json:"lat"`
	// Spacecraft launch arranger.
	LaunchArranger string `json:"launchArranger"`
	// Spacecraft launch arranger country.
	LaunchArrangerCountry string `json:"launchArrangerCountry"`
	// Seradata launch characteristic (e.g. Expendable, Reusable (New), etc).
	LaunchCharacteristic string `json:"launchCharacteristic"`
	// Cost of launch in M USD.
	LaunchCost float64 `json:"launchCost"`
	// Boolean indicating if the spacecraft launch cost is estimated.
	LaunchCostEstimated bool `json:"launchCostEstimated"`
	// Seradata launch country.
	LaunchCountry string `json:"launchCountry"`
	// Launch date.
	LaunchDate time.Time `json:"launchDate" format:"date-time"`
	// Seradata remarks on launch date.
	LaunchDateRemarks string `json:"launchDateRemarks"`
	// Seradata launch ID.
	LaunchID string `json:"launchId"`
	// Mass at launch in kg.
	LaunchMass float64 `json:"launchMass"`
	// Insurance notes for the spacecraft.
	LaunchNotes string `json:"launchNotes"`
	// Seradata launch number.
	LaunchNumber string `json:"launchNumber"`
	// Seradata launch provider.
	LaunchProvider string `json:"launchProvider"`
	// Seradata launch provider country.
	LaunchProviderCountry string `json:"launchProviderCountry"`
	// Seradata launch vehicle family.
	LaunchProviderFlightNumber string `json:"launchProviderFlightNumber"`
	// Seradata Launch Site ID.
	LaunchSiteID string `json:"launchSiteId"`
	// Launch Site Name.
	LaunchSiteName string `json:"launchSiteName"`
	// Seradata launch type (e.g. Launched, Future, etc).
	LaunchType string `json:"launchType"`
	// Seradata launch ID.
	LaunchVehicleID string `json:"launchVehicleId"`
	// Boolean indicating if the spacecraft was leased.
	Leased bool `json:"leased"`
	// Spacecraft life lost as a percent fraction.
	LifeLost float64 `json:"lifeLost"`
	// Spacecraft latest longitude in degrees.
	Lon float64 `json:"lon"`
	// Mass category (e.g. 2500 - 3500kg - Large Satellite, 10 - 100 kg -
	// Microsatellite, etc).
	MassCategory string `json:"massCategory"`
	// Spacecraft name at launch.
	NameAtLaunch string `json:"nameAtLaunch"`
	// Cost of spacecraft at new in M USD.
	NewCost float64 `json:"newCost"`
	// Notes on the spacecraft.
	Notes string `json:"notes"`
	// Number of humans carried on spacecraft.
	NumHumans int64 `json:"numHumans"`
	// Spacecraft operator name.
	Operator string `json:"operator"`
	// Spacecraft operator country.
	OperatorCountry string `json:"operatorCountry"`
	// Spacecraft orbit category (e.g GEO, LEO, etc).
	OrbitCategory string `json:"orbitCategory"`
	// Spacecraft sub orbit category (e.g LEO - Sun-synchronous, Geostationary, etc).
	OrbitSubCategory string `json:"orbitSubCategory"`
	// Spacecraft order date.
	OrderDate time.Time `json:"orderDate" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Spacecraft owner name.
	Owner string `json:"owner"`
	// Spacecraft owner country.
	OwnerCountry string `json:"ownerCountry"`
	// Perigee in km.
	Perigee float64 `json:"perigee"`
	// Spacecraft period in minutes.
	Period float64 `json:"period"`
	// Spacecraft primary mission and group.
	PrimaryMissionGroup string `json:"primaryMissionGroup"`
	// UUID of the prime manufacturer organization, if available.
	PrimeManufacturerOrgID string `json:"primeManufacturerOrgId"`
	// Spacecraft program name.
	ProgramName string `json:"programName"`
	// Spacecraft quantity.
	Quantity int64 `json:"quantity"`
	// Spacecraft reusable flights.
	ReusableFlights string `json:"reusableFlights"`
	// Spacecraft reused hull name.
	ReusedHullName string `json:"reusedHullName"`
	// Seradata sector (e.g. Commercial, Military, Civil/Other).
	Sector string `json:"sector"`
	// Spacecraft serial number.
	SerialNumber string `json:"serialNumber"`
	// Spacecraft stabilizer (e.g. 3-Axis, Gravity Gradiant, etc).
	Stabilizer string `json:"stabilizer"`
	// Spacecraft status (e.g. Inactive - Retired, Inactive - Re-entered, Active, etc).
	Status string `json:"status"`
	// Number of insurance claims for this spacecraft.
	TotalClaims int64 `json:"totalClaims"`
	// Number of fatalities related to this spacecraft.
	TotalFatalities int64 `json:"totalFatalities"`
	// Number of injuries related to this spacecraft.
	TotalInjuries int64 `json:"totalInjuries"`
	// Mass dry in kg.
	TotalPayloadPower float64 `json:"totalPayloadPower"`
	// Youtube link of launch.
	YoutubeLaunchLink string `json:"youtubeLaunchLink"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                    respjson.Field
		DataMode                                 respjson.Field
		Name                                     respjson.Field
		Source                                   respjson.Field
		ID                                       respjson.Field
		AdditionalMissionsGroups                 respjson.Field
		Altitude                                 respjson.Field
		AnnualInsuredDepreciationFactor          respjson.Field
		AnnualInsuredDepreciationFactorEstimated respjson.Field
		Apogee                                   respjson.Field
		BusID                                    respjson.Field
		CapabilityLost                           respjson.Field
		CapacityLost                             respjson.Field
		CatalogNumber                            respjson.Field
		CollisionRiskCm                          respjson.Field
		CollisionRiskMm                          respjson.Field
		CombinedCostEstimated                    respjson.Field
		CombinedNewCost                          respjson.Field
		CommercialLaunch                         respjson.Field
		Constellation                            respjson.Field
		CostEstimated                            respjson.Field
		CreatedAt                                respjson.Field
		CreatedBy                                respjson.Field
		CubesatDispenserType                     respjson.Field
		CurrentAge                               respjson.Field
		DateOfObservation                        respjson.Field
		Description                              respjson.Field
		DesignLife                               respjson.Field
		DryMass                                  respjson.Field
		ExpectedLife                             respjson.Field
		GeoPosition                              respjson.Field
		IDOnOrbit                                respjson.Field
		Inclination                              respjson.Field
		InsuranceLossesTotal                     respjson.Field
		InsuranceNotes                           respjson.Field
		InsurancePremiumAtLaunch                 respjson.Field
		InsurancePremiumAtLaunchEstimated        respjson.Field
		InsuredAtLaunch                          respjson.Field
		InsuredValueAtLaunch                     respjson.Field
		InsuredValueLaunchEstimated              respjson.Field
		IntlNumber                               respjson.Field
		Lat                                      respjson.Field
		LaunchArranger                           respjson.Field
		LaunchArrangerCountry                    respjson.Field
		LaunchCharacteristic                     respjson.Field
		LaunchCost                               respjson.Field
		LaunchCostEstimated                      respjson.Field
		LaunchCountry                            respjson.Field
		LaunchDate                               respjson.Field
		LaunchDateRemarks                        respjson.Field
		LaunchID                                 respjson.Field
		LaunchMass                               respjson.Field
		LaunchNotes                              respjson.Field
		LaunchNumber                             respjson.Field
		LaunchProvider                           respjson.Field
		LaunchProviderCountry                    respjson.Field
		LaunchProviderFlightNumber               respjson.Field
		LaunchSiteID                             respjson.Field
		LaunchSiteName                           respjson.Field
		LaunchType                               respjson.Field
		LaunchVehicleID                          respjson.Field
		Leased                                   respjson.Field
		LifeLost                                 respjson.Field
		Lon                                      respjson.Field
		MassCategory                             respjson.Field
		NameAtLaunch                             respjson.Field
		NewCost                                  respjson.Field
		Notes                                    respjson.Field
		NumHumans                                respjson.Field
		Operator                                 respjson.Field
		OperatorCountry                          respjson.Field
		OrbitCategory                            respjson.Field
		OrbitSubCategory                         respjson.Field
		OrderDate                                respjson.Field
		Origin                                   respjson.Field
		OrigNetwork                              respjson.Field
		Owner                                    respjson.Field
		OwnerCountry                             respjson.Field
		Perigee                                  respjson.Field
		Period                                   respjson.Field
		PrimaryMissionGroup                      respjson.Field
		PrimeManufacturerOrgID                   respjson.Field
		ProgramName                              respjson.Field
		Quantity                                 respjson.Field
		ReusableFlights                          respjson.Field
		ReusedHullName                           respjson.Field
		Sector                                   respjson.Field
		SerialNumber                             respjson.Field
		Stabilizer                               respjson.Field
		Status                                   respjson.Field
		TotalClaims                              respjson.Field
		TotalFatalities                          respjson.Field
		TotalInjuries                            respjson.Field
		TotalPayloadPower                        respjson.Field
		YoutubeLaunchLink                        respjson.Field
		ExtraFields                              map[string]respjson.Field
		raw                                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailListResponse) UnmarshalJSON(data []byte) error {
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
type SeradataSpacecraftDetailListResponseDataMode string

const (
	SeradataSpacecraftDetailListResponseDataModeReal      SeradataSpacecraftDetailListResponseDataMode = "REAL"
	SeradataSpacecraftDetailListResponseDataModeTest      SeradataSpacecraftDetailListResponseDataMode = "TEST"
	SeradataSpacecraftDetailListResponseDataModeSimulated SeradataSpacecraftDetailListResponseDataMode = "SIMULATED"
	SeradataSpacecraftDetailListResponseDataModeExercise  SeradataSpacecraftDetailListResponseDataMode = "EXERCISE"
)

// On-orbit spacecraft details compiled by Seradata for a particular satellite.
type SeradataSpacecraftDetailGetResponse struct {
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
	DataMode SeradataSpacecraftDetailGetResponseDataMode `json:"dataMode,required"`
	// Spacecraft name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Spacecraft additional missions and groups.
	AdditionalMissionsGroups string `json:"additionalMissionsGroups"`
	// Spacecraft latest altitude in km.
	Altitude float64 `json:"altitude"`
	// Annual insured depreciaion factor as a percent fraction.
	AnnualInsuredDepreciationFactor float64 `json:"annualInsuredDepreciationFactor"`
	// Boolean indicating if the spacecraft annualInsuredDepreciationFactor is
	// estimated.
	AnnualInsuredDepreciationFactorEstimated bool `json:"annualInsuredDepreciationFactorEstimated"`
	// Apogee in km.
	Apogee float64 `json:"apogee"`
	// Spacecraft Bus ID.
	BusID string `json:"busId"`
	// Total capability lost as a percent fraction.
	CapabilityLost float64 `json:"capabilityLost"`
	// Total capacity lost as a percent fraction.
	CapacityLost float64 `json:"capacityLost"`
	// NORAD satellite number if available.
	CatalogNumber int64 `json:"catalogNumber"`
	// Spacecraft collision risk 1cm sqm latest.
	CollisionRiskCm float64 `json:"collisionRiskCM"`
	// Spacecraft collision risk 1mm sqm latest.
	CollisionRiskMm float64 `json:"collisionRiskMM"`
	// Boolean indicating if the spacecraft combined new cost is estimated.
	CombinedCostEstimated bool `json:"combinedCostEstimated"`
	// Combined cost of spacecraft at new in M USD.
	CombinedNewCost float64 `json:"combinedNewCost"`
	// Boolean indicating if the launch was commercial.
	CommercialLaunch bool `json:"commercialLaunch"`
	// Spacecraft constellation.
	Constellation string `json:"constellation"`
	// Boolean indicating if the spacecraft cost is estimated.
	CostEstimated bool `json:"costEstimated"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Cubesat dispenser type.
	CubesatDispenserType string `json:"cubesatDispenserType"`
	// Current age in years.
	CurrentAge float64 `json:"currentAge"`
	// Spacecraft date of observation.
	DateOfObservation time.Time `json:"dateOfObservation" format:"date-time"`
	// Description associated with the spacecraft.
	Description string `json:"description"`
	// Spacecraft design life in days.
	DesignLife int64 `json:"designLife"`
	// Mass dry in kg.
	DryMass float64 `json:"dryMass"`
	// Spacecraft expected life in days.
	ExpectedLife int64 `json:"expectedLife"`
	// WGS84 longitude of the spacecraft’s latest GEO position, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	GeoPosition float64 `json:"geoPosition"`
	// UUID of the parent Onorbit record, if available.
	IDOnOrbit string `json:"idOnOrbit"`
	// Seradata provided inclination in degrees.
	Inclination float64 `json:"inclination"`
	// Spacecraft total insurance losses as a fraction.
	InsuranceLossesTotal float64 `json:"insuranceLossesTotal"`
	// Insurance notes for the spacecraft.
	InsuranceNotes string `json:"insuranceNotes"`
	// Insurance premium at launch in M USD.
	InsurancePremiumAtLaunch float64 `json:"insurancePremiumAtLaunch"`
	// Boolean indicating if the spacecraft insurancePremiumAtLaunch is estimated.
	InsurancePremiumAtLaunchEstimated bool `json:"insurancePremiumAtLaunchEstimated"`
	// Boolean indicating if the spacecraft was insured at launch.
	InsuredAtLaunch bool `json:"insuredAtLaunch"`
	// Insured value of spacecraft at launch in M USD.
	InsuredValueAtLaunch float64 `json:"insuredValueAtLaunch"`
	// Boolean indicating if the spacecraft insured value at launch is estimated.
	InsuredValueLaunchEstimated bool `json:"insuredValueLaunchEstimated"`
	// Seradata international number.
	IntlNumber string `json:"intlNumber"`
	// Spacecraft latest latitude in degrees.
	Lat float64 `json:"lat"`
	// Spacecraft launch arranger.
	LaunchArranger string `json:"launchArranger"`
	// Spacecraft launch arranger country.
	LaunchArrangerCountry string `json:"launchArrangerCountry"`
	// Seradata launch characteristic (e.g. Expendable, Reusable (New), etc).
	LaunchCharacteristic string `json:"launchCharacteristic"`
	// Cost of launch in M USD.
	LaunchCost float64 `json:"launchCost"`
	// Boolean indicating if the spacecraft launch cost is estimated.
	LaunchCostEstimated bool `json:"launchCostEstimated"`
	// Seradata launch country.
	LaunchCountry string `json:"launchCountry"`
	// Launch date.
	LaunchDate time.Time `json:"launchDate" format:"date-time"`
	// Seradata remarks on launch date.
	LaunchDateRemarks string `json:"launchDateRemarks"`
	// Seradata launch ID.
	LaunchID string `json:"launchId"`
	// Mass at launch in kg.
	LaunchMass float64 `json:"launchMass"`
	// Insurance notes for the spacecraft.
	LaunchNotes string `json:"launchNotes"`
	// Seradata launch number.
	LaunchNumber string `json:"launchNumber"`
	// Seradata launch provider.
	LaunchProvider string `json:"launchProvider"`
	// Seradata launch provider country.
	LaunchProviderCountry string `json:"launchProviderCountry"`
	// Seradata launch vehicle family.
	LaunchProviderFlightNumber string `json:"launchProviderFlightNumber"`
	// Seradata Launch Site ID.
	LaunchSiteID string `json:"launchSiteId"`
	// Launch Site Name.
	LaunchSiteName string `json:"launchSiteName"`
	// Seradata launch type (e.g. Launched, Future, etc).
	LaunchType string `json:"launchType"`
	// Seradata launch ID.
	LaunchVehicleID string `json:"launchVehicleId"`
	// Boolean indicating if the spacecraft was leased.
	Leased bool `json:"leased"`
	// Spacecraft life lost as a percent fraction.
	LifeLost float64 `json:"lifeLost"`
	// Spacecraft latest longitude in degrees.
	Lon float64 `json:"lon"`
	// Mass category (e.g. 2500 - 3500kg - Large Satellite, 10 - 100 kg -
	// Microsatellite, etc).
	MassCategory string `json:"massCategory"`
	// Spacecraft name at launch.
	NameAtLaunch string `json:"nameAtLaunch"`
	// Cost of spacecraft at new in M USD.
	NewCost float64 `json:"newCost"`
	// Notes on the spacecraft.
	Notes string `json:"notes"`
	// Number of humans carried on spacecraft.
	NumHumans int64 `json:"numHumans"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit SeradataSpacecraftDetailGetResponseOnOrbit `json:"onOrbit"`
	// Spacecraft operator name.
	Operator string `json:"operator"`
	// Spacecraft operator country.
	OperatorCountry string `json:"operatorCountry"`
	// Spacecraft orbit category (e.g GEO, LEO, etc).
	OrbitCategory string `json:"orbitCategory"`
	// Spacecraft sub orbit category (e.g LEO - Sun-synchronous, Geostationary, etc).
	OrbitSubCategory string `json:"orbitSubCategory"`
	// Spacecraft order date.
	OrderDate time.Time `json:"orderDate" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Spacecraft owner name.
	Owner string `json:"owner"`
	// Spacecraft owner country.
	OwnerCountry string `json:"ownerCountry"`
	// Perigee in km.
	Perigee float64 `json:"perigee"`
	// Spacecraft period in minutes.
	Period float64 `json:"period"`
	// Spacecraft primary mission and group.
	PrimaryMissionGroup string `json:"primaryMissionGroup"`
	// UUID of the prime manufacturer organization, if available.
	PrimeManufacturerOrgID string `json:"primeManufacturerOrgId"`
	// Spacecraft program name.
	ProgramName string `json:"programName"`
	// Spacecraft quantity.
	Quantity int64 `json:"quantity"`
	// Spacecraft reusable flights.
	ReusableFlights string `json:"reusableFlights"`
	// Spacecraft reused hull name.
	ReusedHullName string `json:"reusedHullName"`
	// Read-only details of the Scientific object, only used during detail queries (not
	// to be provided on create/update operations).
	Scientific []SeradataSpacecraftDetailGetResponseScientific `json:"scientific"`
	// Seradata sector (e.g. Commercial, Military, Civil/Other).
	Sector string `json:"sector"`
	// Read-only details of the SeradataCommDetails object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataCommDetails []SeradataSpacecraftDetailGetResponseSeradataCommDetail `json:"seradataCommDetails"`
	// Read-only details of the SeradataEarlyWarning object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataEarlyWarning []SeradataSpacecraftDetailGetResponseSeradataEarlyWarning `json:"seradataEarlyWarning"`
	// Read-only details of the SeradataNavigation object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataNavigation []SeradataSpacecraftDetailGetResponseSeradataNavigation `json:"seradataNavigation"`
	// Read-only details of the SeradataOpticalPayload object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataOpticalPayload []SeradataSpacecraftDetailGetResponseSeradataOpticalPayload `json:"seradataOpticalPayload"`
	// Read-only details of the SeradataRadarPayload object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataRadarPayload []SeradataSpacecraftDetailGetResponseSeradataRadarPayload `json:"seradataRadarPayload"`
	// Read-only details of the SeradataSigIntPayload object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataSigIntPayload []SeradataSpacecraftDetailGetResponseSeradataSigIntPayload `json:"seradataSigIntPayload"`
	// Spacecraft serial number.
	SerialNumber string `json:"serialNumber"`
	// Spacecraft stabilizer (e.g. 3-Axis, Gravity Gradiant, etc).
	Stabilizer string `json:"stabilizer"`
	// Spacecraft status (e.g. Inactive - Retired, Inactive - Re-entered, Active, etc).
	Status string `json:"status"`
	// Number of insurance claims for this spacecraft.
	TotalClaims int64 `json:"totalClaims"`
	// Number of fatalities related to this spacecraft.
	TotalFatalities int64 `json:"totalFatalities"`
	// Number of injuries related to this spacecraft.
	TotalInjuries int64 `json:"totalInjuries"`
	// Mass dry in kg.
	TotalPayloadPower float64 `json:"totalPayloadPower"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Youtube link of launch.
	YoutubeLaunchLink string `json:"youtubeLaunchLink"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                    respjson.Field
		DataMode                                 respjson.Field
		Name                                     respjson.Field
		Source                                   respjson.Field
		ID                                       respjson.Field
		AdditionalMissionsGroups                 respjson.Field
		Altitude                                 respjson.Field
		AnnualInsuredDepreciationFactor          respjson.Field
		AnnualInsuredDepreciationFactorEstimated respjson.Field
		Apogee                                   respjson.Field
		BusID                                    respjson.Field
		CapabilityLost                           respjson.Field
		CapacityLost                             respjson.Field
		CatalogNumber                            respjson.Field
		CollisionRiskCm                          respjson.Field
		CollisionRiskMm                          respjson.Field
		CombinedCostEstimated                    respjson.Field
		CombinedNewCost                          respjson.Field
		CommercialLaunch                         respjson.Field
		Constellation                            respjson.Field
		CostEstimated                            respjson.Field
		CreatedAt                                respjson.Field
		CreatedBy                                respjson.Field
		CubesatDispenserType                     respjson.Field
		CurrentAge                               respjson.Field
		DateOfObservation                        respjson.Field
		Description                              respjson.Field
		DesignLife                               respjson.Field
		DryMass                                  respjson.Field
		ExpectedLife                             respjson.Field
		GeoPosition                              respjson.Field
		IDOnOrbit                                respjson.Field
		Inclination                              respjson.Field
		InsuranceLossesTotal                     respjson.Field
		InsuranceNotes                           respjson.Field
		InsurancePremiumAtLaunch                 respjson.Field
		InsurancePremiumAtLaunchEstimated        respjson.Field
		InsuredAtLaunch                          respjson.Field
		InsuredValueAtLaunch                     respjson.Field
		InsuredValueLaunchEstimated              respjson.Field
		IntlNumber                               respjson.Field
		Lat                                      respjson.Field
		LaunchArranger                           respjson.Field
		LaunchArrangerCountry                    respjson.Field
		LaunchCharacteristic                     respjson.Field
		LaunchCost                               respjson.Field
		LaunchCostEstimated                      respjson.Field
		LaunchCountry                            respjson.Field
		LaunchDate                               respjson.Field
		LaunchDateRemarks                        respjson.Field
		LaunchID                                 respjson.Field
		LaunchMass                               respjson.Field
		LaunchNotes                              respjson.Field
		LaunchNumber                             respjson.Field
		LaunchProvider                           respjson.Field
		LaunchProviderCountry                    respjson.Field
		LaunchProviderFlightNumber               respjson.Field
		LaunchSiteID                             respjson.Field
		LaunchSiteName                           respjson.Field
		LaunchType                               respjson.Field
		LaunchVehicleID                          respjson.Field
		Leased                                   respjson.Field
		LifeLost                                 respjson.Field
		Lon                                      respjson.Field
		MassCategory                             respjson.Field
		NameAtLaunch                             respjson.Field
		NewCost                                  respjson.Field
		Notes                                    respjson.Field
		NumHumans                                respjson.Field
		OnOrbit                                  respjson.Field
		Operator                                 respjson.Field
		OperatorCountry                          respjson.Field
		OrbitCategory                            respjson.Field
		OrbitSubCategory                         respjson.Field
		OrderDate                                respjson.Field
		Origin                                   respjson.Field
		OrigNetwork                              respjson.Field
		Owner                                    respjson.Field
		OwnerCountry                             respjson.Field
		Perigee                                  respjson.Field
		Period                                   respjson.Field
		PrimaryMissionGroup                      respjson.Field
		PrimeManufacturerOrgID                   respjson.Field
		ProgramName                              respjson.Field
		Quantity                                 respjson.Field
		ReusableFlights                          respjson.Field
		ReusedHullName                           respjson.Field
		Scientific                               respjson.Field
		Sector                                   respjson.Field
		SeradataCommDetails                      respjson.Field
		SeradataEarlyWarning                     respjson.Field
		SeradataNavigation                       respjson.Field
		SeradataOpticalPayload                   respjson.Field
		SeradataRadarPayload                     respjson.Field
		SeradataSigIntPayload                    respjson.Field
		SerialNumber                             respjson.Field
		Stabilizer                               respjson.Field
		Status                                   respjson.Field
		TotalClaims                              respjson.Field
		TotalFatalities                          respjson.Field
		TotalInjuries                            respjson.Field
		TotalPayloadPower                        respjson.Field
		UpdatedAt                                respjson.Field
		UpdatedBy                                respjson.Field
		YoutubeLaunchLink                        respjson.Field
		ExtraFields                              map[string]respjson.Field
		raw                                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailGetResponse) UnmarshalJSON(data []byte) error {
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
type SeradataSpacecraftDetailGetResponseDataMode string

const (
	SeradataSpacecraftDetailGetResponseDataModeReal      SeradataSpacecraftDetailGetResponseDataMode = "REAL"
	SeradataSpacecraftDetailGetResponseDataModeTest      SeradataSpacecraftDetailGetResponseDataMode = "TEST"
	SeradataSpacecraftDetailGetResponseDataModeSimulated SeradataSpacecraftDetailGetResponseDataMode = "SIMULATED"
	SeradataSpacecraftDetailGetResponseDataModeExercise  SeradataSpacecraftDetailGetResponseDataMode = "EXERCISE"
)

// Model object representing on-orbit objects or satellites in the system.
type SeradataSpacecraftDetailGetResponseOnOrbit struct {
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
	DataMode string `json:"dataMode,required"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Alternate name of the on-orbit object.
	AltName string `json:"altName"`
	// Read-only collection of antennas on this on-orbit object.
	Antennas []shared.OnorbitAntennaFull `json:"antennas"`
	// Read-only collection of batteries on this on-orbit object.
	Batteries []shared.OnorbitBatteryFull `json:"batteries"`
	// Category of the on-orbit object. (Unknown, On-Orbit, Decayed, Cataloged Without
	// State, Launch Nominal, Analyst Satellite, Cislunar, Lunar, Hyperbolic,
	// Heliocentric, Interplanetary, Lagrangian, Docked).
	//
	// Any of "Unknown", "On-Orbit", "Decayed", "Cataloged Without State", "Launch
	// Nominal", "Analyst Satellite", "Cislunar", "Lunar", "Hyperbolic",
	// "Heliocentric", "Interplanetary", "Lagrangian", "Docked".
	Category string `json:"category"`
	// Common name of the on-orbit object.
	CommonName string `json:"commonName"`
	// Constellation to which this satellite belongs.
	Constellation string `json:"constellation"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Date of decay.
	DecayDate time.Time `json:"decayDate" format:"date-time"`
	// For the public catalog, the idOnOrbit is typically the satellite number as a
	// string, but may be a UUID for analyst or other unknown or untracked satellites,
	// auto-generated by the system.
	IDOnOrbit string `json:"idOnOrbit"`
	// International Designator, typically of the format YYYYLLLAAA, where YYYY is the
	// launch year, LLL is the sequential launch number of that year, and AAA is an
	// optional launch piece designator for the launch.
	IntlDes string `json:"intlDes"`
	// Date of launch.
	LaunchDate time.Time `json:"launchDate" format:"date"`
	// Id of the associated launchSite entity.
	LaunchSiteID string `json:"launchSiteId"`
	// Estimated lifetime of the on-orbit payload, if known.
	LifetimeYears int64 `json:"lifetimeYears"`
	// Mission number of the on-orbit object.
	MissionNumber string `json:"missionNumber"`
	// Type of on-orbit object: ROCKET BODY, DEBRIS, PAYLOAD, PLATFORM, MANNED,
	// UNKNOWN.
	//
	// Any of "ROCKET BODY", "DEBRIS", "PAYLOAD", "PLATFORM", "MANNED", "UNKNOWN".
	ObjectType string `json:"objectType"`
	// Read-only collection of details for this on-orbit object.
	OnorbitDetails []shared.OnorbitDetailsFull `json:"onorbitDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only collection of solar arrays on this on-orbit object.
	SolarArrays []shared.OnorbitSolarArrayFull `json:"solarArrays"`
	// Read-only collection of thrusters (engines) on this on-orbit object.
	Thrusters []shared.OnorbitThrusterFull `json:"thrusters"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		SatNo                 respjson.Field
		Source                respjson.Field
		AltName               respjson.Field
		Antennas              respjson.Field
		Batteries             respjson.Field
		Category              respjson.Field
		CommonName            respjson.Field
		Constellation         respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecayDate             respjson.Field
		IDOnOrbit             respjson.Field
		IntlDes               respjson.Field
		LaunchDate            respjson.Field
		LaunchSiteID          respjson.Field
		LifetimeYears         respjson.Field
		MissionNumber         respjson.Field
		ObjectType            respjson.Field
		OnorbitDetails        respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SolarArrays           respjson.Field
		Thrusters             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailGetResponseOnOrbit) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailGetResponseOnOrbit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scientific or other data from Seradata.
type SeradataSpacecraftDetailGetResponseScientific struct {
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
	DataMode string `json:"dataMode,required"`
	// Sensor name from sera data, e.g. SEM/MAG (SEM / Magnetometer).
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity shared.EntityFull `json:"entity"`
	// Frequency band, e.g. Gamma.
	FrequencyBand string `json:"frequencyBand"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// ID of the parent entity for this Scientific.
	IDEntity string `json:"idEntity"`
	// Unique identifier of the organization which manufactures this bus.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Notes associated with the payload.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Payload category, e.g. Magnetometer, Radiometer, Sensor, etc.
	PayloadCategory string `json:"payloadCategory"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		SpacecraftID          respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Entity                respjson.Field
		FrequencyBand         respjson.Field
		HostedForCompanyOrgID respjson.Field
		IDEntity              respjson.Field
		ManufacturerOrgID     respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PayloadCategory       respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailGetResponseScientific) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailGetResponseScientific) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Seradata-compiled information on communications payloads.
type SeradataSpacecraftDetailGetResponseSeradataCommDetail struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band string `json:"band"`
	// Comm bandwidth in Mhz.
	Bandwidth float64 `json:"bandwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Effective isotropic radiated power in dB.
	Eirp float64 `json:"eirp"`
	// Comm estimated HtsTotalCapacity in Gbps.
	EstHtsTotalCapacity float64 `json:"estHtsTotalCapacity"`
	// Comm estimated HtsTotalUserDownlinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserDownlinkBandwidthPerBeam float64 `json:"estHtsTotalUserDownlinkBandwidthPerBeam"`
	// Comm estimated HtsTotalUserUplinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserUplinkBandwidthPerBeam float64 `json:"estHtsTotalUserUplinkBandwidthPerBeam"`
	// Comm gatewayDownlinkFrom in Ghz.
	GatewayDownlinkFrom float64 `json:"gatewayDownlinkFrom"`
	// Comm gatewayDownlinkTo in Ghz.
	GatewayDownlinkTo float64 `json:"gatewayDownlinkTo"`
	// Comm gatewayUplinkFrom in Ghz.
	GatewayUplinkFrom float64 `json:"gatewayUplinkFrom"`
	// Comm gatewayUplinkTo in Ghz.
	GatewayUplinkTo float64 `json:"gatewayUplinkTo"`
	// Comm hostedForCompanyOrgId.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// Comm htsNumUserSpotBeams.
	HtsNumUserSpotBeams int64 `json:"htsNumUserSpotBeams"`
	// Comm htsUserDownlinkBandwidthPerBeam in Mhz.
	HtsUserDownlinkBandwidthPerBeam float64 `json:"htsUserDownlinkBandwidthPerBeam"`
	// Comm htsUserUplinkBandwidthPerBeam in Mhz.
	HtsUserUplinkBandwidthPerBeam float64 `json:"htsUserUplinkBandwidthPerBeam"`
	// UUID of the parent Comm record.
	IDComm string `json:"idComm"`
	// Comm manufacturerOrgId.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Comm num36MhzEquivalentTransponders.
	Num36MhzEquivalentTransponders int64 `json:"num36MhzEquivalentTransponders"`
	// Comm numOperationalTransponders.
	NumOperationalTransponders int64 `json:"numOperationalTransponders"`
	// Comm numSpareTransponders.
	NumSpareTransponders int64 `json:"numSpareTransponders"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Payload notes.
	PayloadNotes string `json:"payloadNotes"`
	// Comm polarization.
	Polarization string `json:"polarization"`
	// Solid state power amplifier, in Watts.
	SolidStatePowerAmp float64 `json:"solidStatePowerAmp"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId"`
	// Comm tradeLeaseOrgId.
	TradeLeaseOrgID string `json:"tradeLeaseOrgId"`
	// Comm travelingWaveTubeAmplifier in Watts.
	TravelingWaveTubeAmplifier float64 `json:"travelingWaveTubeAmplifier"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Comm userDownlinkFrom in Ghz.
	UserDownlinkFrom float64 `json:"userDownlinkFrom"`
	// Comm userDownlinkTo in Ghz.
	UserDownlinkTo float64 `json:"userDownlinkTo"`
	// Comm userUplinkFrom in Ghz.
	UserUplinkFrom float64 `json:"userUplinkFrom"`
	// Comm userUplinkTo in Ghz.
	UserUplinkTo float64 `json:"userUplinkTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                   respjson.Field
		DataMode                                respjson.Field
		Source                                  respjson.Field
		ID                                      respjson.Field
		Band                                    respjson.Field
		Bandwidth                               respjson.Field
		CreatedAt                               respjson.Field
		CreatedBy                               respjson.Field
		Eirp                                    respjson.Field
		EstHtsTotalCapacity                     respjson.Field
		EstHtsTotalUserDownlinkBandwidthPerBeam respjson.Field
		EstHtsTotalUserUplinkBandwidthPerBeam   respjson.Field
		GatewayDownlinkFrom                     respjson.Field
		GatewayDownlinkTo                       respjson.Field
		GatewayUplinkFrom                       respjson.Field
		GatewayUplinkTo                         respjson.Field
		HostedForCompanyOrgID                   respjson.Field
		HtsNumUserSpotBeams                     respjson.Field
		HtsUserDownlinkBandwidthPerBeam         respjson.Field
		HtsUserUplinkBandwidthPerBeam           respjson.Field
		IDComm                                  respjson.Field
		ManufacturerOrgID                       respjson.Field
		Num36MhzEquivalentTransponders          respjson.Field
		NumOperationalTransponders              respjson.Field
		NumSpareTransponders                    respjson.Field
		Origin                                  respjson.Field
		OrigNetwork                             respjson.Field
		PayloadNotes                            respjson.Field
		Polarization                            respjson.Field
		SolidStatePowerAmp                      respjson.Field
		SpacecraftID                            respjson.Field
		TradeLeaseOrgID                         respjson.Field
		TravelingWaveTubeAmplifier              respjson.Field
		UpdatedAt                               respjson.Field
		UpdatedBy                               respjson.Field
		UserDownlinkFrom                        respjson.Field
		UserDownlinkTo                          respjson.Field
		UserUplinkFrom                          respjson.Field
		UserUplinkTo                            respjson.Field
		ExtraFields                             map[string]respjson.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailGetResponseSeradataCommDetail) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailGetResponseSeradataCommDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for an early warning payload from Seradata.
type SeradataSpacecraftDetailGetResponseSeradataEarlyWarning struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Best resolution for this IR in meters.
	BestResolution float64 `json:"bestResolution"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Is the sensor Earth Pointing.
	EarthPointing bool `json:"earthPointing"`
	// Frequency Limits for this IR.
	FrequencyLimits string `json:"frequencyLimits"`
	// Ground Station Locations for this IR.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this IR.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the parent IR record.
	IDIr string `json:"idIR"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Missile Launch Phase Detection Ability.
	MissileLaunchPhaseDetectionAbility string `json:"missileLaunchPhaseDetectionAbility"`
	// Sensor name from Seradata, e.g. Infra red telescope, Schmidt Telescope, etc.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID string `json:"partnerSpacecraftId"`
	// Payload notes.
	PayloadNotes string `json:"payloadNotes"`
	// Spectral Bands, e.g. Infra-Red.
	SpectralBands string `json:"spectralBands"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking              respjson.Field
		DataMode                           respjson.Field
		Source                             respjson.Field
		SpacecraftID                       respjson.Field
		ID                                 respjson.Field
		BestResolution                     respjson.Field
		CreatedAt                          respjson.Field
		CreatedBy                          respjson.Field
		EarthPointing                      respjson.Field
		FrequencyLimits                    respjson.Field
		GroundStationLocations             respjson.Field
		GroundStations                     respjson.Field
		HostedForCompanyOrgID              respjson.Field
		IDIr                               respjson.Field
		ManufacturerOrgID                  respjson.Field
		MissileLaunchPhaseDetectionAbility respjson.Field
		Name                               respjson.Field
		Origin                             respjson.Field
		OrigNetwork                        respjson.Field
		PartnerSpacecraftID                respjson.Field
		PayloadNotes                       respjson.Field
		SpectralBands                      respjson.Field
		UpdatedAt                          respjson.Field
		UpdatedBy                          respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailGetResponseSeradataEarlyWarning) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailGetResponseSeradataEarlyWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for a navigation payload from Seradata.
type SeradataSpacecraftDetailGetResponseSeradataNavigation struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Area of coverage, e.g. Worldwide, India, etc.
	AreaCoverage string `json:"areaCoverage"`
	// Type of clock, e.g. Rubidium, Hydrogen Maser, etc.
	ClockType string `json:"clockType"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the parent Navigation record.
	IDNavigation string `json:"idNavigation"`
	// Location accuracy in meters.
	LocationAccuracy float64 `json:"locationAccuracy"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Frequency for this payload.
	ModeFrequency string `json:"modeFrequency"`
	// Modes of operation.
	Modes string `json:"modes"`
	// Sensor name from Seradata, e.g. WAAS GEO-5, etc.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID string `json:"partnerSpacecraftId"`
	// Navigation payload type, e.g. WAAS, GAGAN, etc.
	PayloadType string `json:"payloadType"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		SpacecraftID          respjson.Field
		ID                    respjson.Field
		AreaCoverage          respjson.Field
		ClockType             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		HostedForCompanyOrgID respjson.Field
		IDNavigation          respjson.Field
		LocationAccuracy      respjson.Field
		ManufacturerOrgID     respjson.Field
		ModeFrequency         respjson.Field
		Modes                 respjson.Field
		Name                  respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PartnerSpacecraftID   respjson.Field
		PayloadType           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailGetResponseSeradataNavigation) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailGetResponseSeradataNavigation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for an optical payload from Seradata.
type SeradataSpacecraftDetailGetResponseSeradataOpticalPayload struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Best resolution.
	BestResolution float64 `json:"bestResolution"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Field of regard of this radar in degrees.
	FieldOfRegard float64 `json:"fieldOfRegard"`
	// Field of view of this radar in kilometers.
	FieldOfView float64 `json:"fieldOfView"`
	// Ground Station Locations for this payload.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this payload.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the Sensor record.
	IDSensor string `json:"idSensor"`
	// Imaging category for this payload, e.g. Multispectral, Infrared, Panchromatic.
	ImagingPayloadCategory string `json:"imagingPayloadCategory"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Sensor name from Seradata, e.g. TOURNESOL, MESSR (Multispectral Self-Scanning
	// Radiometer), AWFI, etc.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Number of film return canisters.
	NumberOfFilmReturnCanisters int64 `json:"numberOfFilmReturnCanisters"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Point method for this radar, e.g. Spacecraft.
	PointingMethod string `json:"pointingMethod"`
	// Recorder size.
	RecorderSize string `json:"recorderSize"`
	// Spectral Band supported by this payload, e.g. Green, Red, Mid-wave infrared,
	// etc.
	SpectralBand string `json:"spectralBand"`
	// Frequency limit for this payload, e.g. 0.51 - 0.59.
	SpectralFrequencyLimits string `json:"spectralFrequencyLimits"`
	// Swath width in kilometers.
	SwathWidth float64 `json:"swathWidth"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking       respjson.Field
		DataMode                    respjson.Field
		Source                      respjson.Field
		SpacecraftID                respjson.Field
		ID                          respjson.Field
		BestResolution              respjson.Field
		CreatedAt                   respjson.Field
		CreatedBy                   respjson.Field
		FieldOfRegard               respjson.Field
		FieldOfView                 respjson.Field
		GroundStationLocations      respjson.Field
		GroundStations              respjson.Field
		HostedForCompanyOrgID       respjson.Field
		IDSensor                    respjson.Field
		ImagingPayloadCategory      respjson.Field
		ManufacturerOrgID           respjson.Field
		Name                        respjson.Field
		Notes                       respjson.Field
		NumberOfFilmReturnCanisters respjson.Field
		Origin                      respjson.Field
		OrigNetwork                 respjson.Field
		PointingMethod              respjson.Field
		RecorderSize                respjson.Field
		SpectralBand                respjson.Field
		SpectralFrequencyLimits     respjson.Field
		SwathWidth                  respjson.Field
		UpdatedAt                   respjson.Field
		UpdatedBy                   respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailGetResponseSeradataOpticalPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *SeradataSpacecraftDetailGetResponseSeradataOpticalPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for an radar payload from Seradata.
type SeradataSpacecraftDetailGetResponseSeradataRadarPayload struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Radar bandwidth in mega hertz.
	Bandwidth float64 `json:"bandwidth"`
	// Best resolution in meters.
	BestResolution float64 `json:"bestResolution"`
	// Radar category, e.g. SAR, Surface Search, etc.
	Category string `json:"category"`
	// Constellation interferometric capability.
	ConstellationInterferometricCapability string `json:"constellationInterferometricCapability"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Duty cycle.
	DutyCycle string `json:"dutyCycle"`
	// Field of regard of this radar in degrees.
	FieldOfRegard float64 `json:"fieldOfRegard"`
	// Field of view of this radar in kilometers.
	FieldOfView float64 `json:"fieldOfView"`
	// Frequency in giga hertz.
	Frequency float64 `json:"frequency"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	FrequencyBand string `json:"frequencyBand"`
	// Ground Station Locations for this payload.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this payload.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the Sensor record.
	IDSensor string `json:"idSensor"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Sensor name from Seradata, e.g. ALT (Radar Altimeter), COSI (Corea SAR
	// Instrument), etc.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Partner seradata-spacecraft.
	PartnerSpacecraft string `json:"partnerSpacecraft"`
	// Point method for this radar, e.g. Spacecraft.
	PointingMethod string `json:"pointingMethod"`
	// Receive polarization, e.g. Lin Dual, Lin vert, etc.
	ReceivePolarization string `json:"receivePolarization"`
	// Recorder size, e.g. 256.
	RecorderSize string `json:"recorderSize"`
	// Swath width in kilometers.
	SwathWidth float64 `json:"swathWidth"`
	// Transmit polarization, e.g. Lin Dual, Lin vert, etc.
	TransmitPolarization string `json:"transmitPolarization"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Wave length in meters.
	WaveLength float64 `json:"waveLength"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                  respjson.Field
		DataMode                               respjson.Field
		Source                                 respjson.Field
		SpacecraftID                           respjson.Field
		ID                                     respjson.Field
		Bandwidth                              respjson.Field
		BestResolution                         respjson.Field
		Category                               respjson.Field
		ConstellationInterferometricCapability respjson.Field
		CreatedAt                              respjson.Field
		CreatedBy                              respjson.Field
		DutyCycle                              respjson.Field
		FieldOfRegard                          respjson.Field
		FieldOfView                            respjson.Field
		Frequency                              respjson.Field
		FrequencyBand                          respjson.Field
		GroundStationLocations                 respjson.Field
		GroundStations                         respjson.Field
		HostedForCompanyOrgID                  respjson.Field
		IDSensor                               respjson.Field
		ManufacturerOrgID                      respjson.Field
		Name                                   respjson.Field
		Notes                                  respjson.Field
		Origin                                 respjson.Field
		OrigNetwork                            respjson.Field
		PartnerSpacecraft                      respjson.Field
		PointingMethod                         respjson.Field
		ReceivePolarization                    respjson.Field
		RecorderSize                           respjson.Field
		SwathWidth                             respjson.Field
		TransmitPolarization                   respjson.Field
		UpdatedAt                              respjson.Field
		UpdatedBy                              respjson.Field
		WaveLength                             respjson.Field
		ExtraFields                            map[string]respjson.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailGetResponseSeradataRadarPayload) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailGetResponseSeradataRadarPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for an sigint payload from Seradata.
type SeradataSpacecraftDetailGetResponseSeradataSigIntPayload struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Frequency coverage for this payload.
	FrequencyCoverage string `json:"frequencyCoverage"`
	// Ground Station Locations for this payload.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this payload.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the Sensor record.
	IDSensor string `json:"idSensor"`
	// Intercept parameters.
	InterceptParameters string `json:"interceptParameters"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Sensor name from Seradata.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Positional Accuracy for this payload.
	PositionalAccuracy string `json:"positionalAccuracy"`
	// Swath Width in kilometers.
	SwathWidth float64 `json:"swathWidth"`
	// SIGINT Payload type, e.g. Comint, Elint, etc.
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		Source                 respjson.Field
		SpacecraftID           respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		FrequencyCoverage      respjson.Field
		GroundStationLocations respjson.Field
		GroundStations         respjson.Field
		HostedForCompanyOrgID  respjson.Field
		IDSensor               respjson.Field
		InterceptParameters    respjson.Field
		ManufacturerOrgID      respjson.Field
		Name                   respjson.Field
		Notes                  respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		PositionalAccuracy     respjson.Field
		SwathWidth             respjson.Field
		Type                   respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailGetResponseSeradataSigIntPayload) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailGetResponseSeradataSigIntPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SeradataSpacecraftDetailQueryhelpResponse struct {
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
func (r SeradataSpacecraftDetailQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// On-orbit spacecraft details compiled by Seradata for a particular satellite.
type SeradataSpacecraftDetailTupleResponse struct {
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
	DataMode SeradataSpacecraftDetailTupleResponseDataMode `json:"dataMode,required"`
	// Spacecraft name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Spacecraft additional missions and groups.
	AdditionalMissionsGroups string `json:"additionalMissionsGroups"`
	// Spacecraft latest altitude in km.
	Altitude float64 `json:"altitude"`
	// Annual insured depreciaion factor as a percent fraction.
	AnnualInsuredDepreciationFactor float64 `json:"annualInsuredDepreciationFactor"`
	// Boolean indicating if the spacecraft annualInsuredDepreciationFactor is
	// estimated.
	AnnualInsuredDepreciationFactorEstimated bool `json:"annualInsuredDepreciationFactorEstimated"`
	// Apogee in km.
	Apogee float64 `json:"apogee"`
	// Spacecraft Bus ID.
	BusID string `json:"busId"`
	// Total capability lost as a percent fraction.
	CapabilityLost float64 `json:"capabilityLost"`
	// Total capacity lost as a percent fraction.
	CapacityLost float64 `json:"capacityLost"`
	// NORAD satellite number if available.
	CatalogNumber int64 `json:"catalogNumber"`
	// Spacecraft collision risk 1cm sqm latest.
	CollisionRiskCm float64 `json:"collisionRiskCM"`
	// Spacecraft collision risk 1mm sqm latest.
	CollisionRiskMm float64 `json:"collisionRiskMM"`
	// Boolean indicating if the spacecraft combined new cost is estimated.
	CombinedCostEstimated bool `json:"combinedCostEstimated"`
	// Combined cost of spacecraft at new in M USD.
	CombinedNewCost float64 `json:"combinedNewCost"`
	// Boolean indicating if the launch was commercial.
	CommercialLaunch bool `json:"commercialLaunch"`
	// Spacecraft constellation.
	Constellation string `json:"constellation"`
	// Boolean indicating if the spacecraft cost is estimated.
	CostEstimated bool `json:"costEstimated"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Cubesat dispenser type.
	CubesatDispenserType string `json:"cubesatDispenserType"`
	// Current age in years.
	CurrentAge float64 `json:"currentAge"`
	// Spacecraft date of observation.
	DateOfObservation time.Time `json:"dateOfObservation" format:"date-time"`
	// Description associated with the spacecraft.
	Description string `json:"description"`
	// Spacecraft design life in days.
	DesignLife int64 `json:"designLife"`
	// Mass dry in kg.
	DryMass float64 `json:"dryMass"`
	// Spacecraft expected life in days.
	ExpectedLife int64 `json:"expectedLife"`
	// WGS84 longitude of the spacecraft’s latest GEO position, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	GeoPosition float64 `json:"geoPosition"`
	// UUID of the parent Onorbit record, if available.
	IDOnOrbit string `json:"idOnOrbit"`
	// Seradata provided inclination in degrees.
	Inclination float64 `json:"inclination"`
	// Spacecraft total insurance losses as a fraction.
	InsuranceLossesTotal float64 `json:"insuranceLossesTotal"`
	// Insurance notes for the spacecraft.
	InsuranceNotes string `json:"insuranceNotes"`
	// Insurance premium at launch in M USD.
	InsurancePremiumAtLaunch float64 `json:"insurancePremiumAtLaunch"`
	// Boolean indicating if the spacecraft insurancePremiumAtLaunch is estimated.
	InsurancePremiumAtLaunchEstimated bool `json:"insurancePremiumAtLaunchEstimated"`
	// Boolean indicating if the spacecraft was insured at launch.
	InsuredAtLaunch bool `json:"insuredAtLaunch"`
	// Insured value of spacecraft at launch in M USD.
	InsuredValueAtLaunch float64 `json:"insuredValueAtLaunch"`
	// Boolean indicating if the spacecraft insured value at launch is estimated.
	InsuredValueLaunchEstimated bool `json:"insuredValueLaunchEstimated"`
	// Seradata international number.
	IntlNumber string `json:"intlNumber"`
	// Spacecraft latest latitude in degrees.
	Lat float64 `json:"lat"`
	// Spacecraft launch arranger.
	LaunchArranger string `json:"launchArranger"`
	// Spacecraft launch arranger country.
	LaunchArrangerCountry string `json:"launchArrangerCountry"`
	// Seradata launch characteristic (e.g. Expendable, Reusable (New), etc).
	LaunchCharacteristic string `json:"launchCharacteristic"`
	// Cost of launch in M USD.
	LaunchCost float64 `json:"launchCost"`
	// Boolean indicating if the spacecraft launch cost is estimated.
	LaunchCostEstimated bool `json:"launchCostEstimated"`
	// Seradata launch country.
	LaunchCountry string `json:"launchCountry"`
	// Launch date.
	LaunchDate time.Time `json:"launchDate" format:"date-time"`
	// Seradata remarks on launch date.
	LaunchDateRemarks string `json:"launchDateRemarks"`
	// Seradata launch ID.
	LaunchID string `json:"launchId"`
	// Mass at launch in kg.
	LaunchMass float64 `json:"launchMass"`
	// Insurance notes for the spacecraft.
	LaunchNotes string `json:"launchNotes"`
	// Seradata launch number.
	LaunchNumber string `json:"launchNumber"`
	// Seradata launch provider.
	LaunchProvider string `json:"launchProvider"`
	// Seradata launch provider country.
	LaunchProviderCountry string `json:"launchProviderCountry"`
	// Seradata launch vehicle family.
	LaunchProviderFlightNumber string `json:"launchProviderFlightNumber"`
	// Seradata Launch Site ID.
	LaunchSiteID string `json:"launchSiteId"`
	// Launch Site Name.
	LaunchSiteName string `json:"launchSiteName"`
	// Seradata launch type (e.g. Launched, Future, etc).
	LaunchType string `json:"launchType"`
	// Seradata launch ID.
	LaunchVehicleID string `json:"launchVehicleId"`
	// Boolean indicating if the spacecraft was leased.
	Leased bool `json:"leased"`
	// Spacecraft life lost as a percent fraction.
	LifeLost float64 `json:"lifeLost"`
	// Spacecraft latest longitude in degrees.
	Lon float64 `json:"lon"`
	// Mass category (e.g. 2500 - 3500kg - Large Satellite, 10 - 100 kg -
	// Microsatellite, etc).
	MassCategory string `json:"massCategory"`
	// Spacecraft name at launch.
	NameAtLaunch string `json:"nameAtLaunch"`
	// Cost of spacecraft at new in M USD.
	NewCost float64 `json:"newCost"`
	// Notes on the spacecraft.
	Notes string `json:"notes"`
	// Number of humans carried on spacecraft.
	NumHumans int64 `json:"numHumans"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit SeradataSpacecraftDetailTupleResponseOnOrbit `json:"onOrbit"`
	// Spacecraft operator name.
	Operator string `json:"operator"`
	// Spacecraft operator country.
	OperatorCountry string `json:"operatorCountry"`
	// Spacecraft orbit category (e.g GEO, LEO, etc).
	OrbitCategory string `json:"orbitCategory"`
	// Spacecraft sub orbit category (e.g LEO - Sun-synchronous, Geostationary, etc).
	OrbitSubCategory string `json:"orbitSubCategory"`
	// Spacecraft order date.
	OrderDate time.Time `json:"orderDate" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Spacecraft owner name.
	Owner string `json:"owner"`
	// Spacecraft owner country.
	OwnerCountry string `json:"ownerCountry"`
	// Perigee in km.
	Perigee float64 `json:"perigee"`
	// Spacecraft period in minutes.
	Period float64 `json:"period"`
	// Spacecraft primary mission and group.
	PrimaryMissionGroup string `json:"primaryMissionGroup"`
	// UUID of the prime manufacturer organization, if available.
	PrimeManufacturerOrgID string `json:"primeManufacturerOrgId"`
	// Spacecraft program name.
	ProgramName string `json:"programName"`
	// Spacecraft quantity.
	Quantity int64 `json:"quantity"`
	// Spacecraft reusable flights.
	ReusableFlights string `json:"reusableFlights"`
	// Spacecraft reused hull name.
	ReusedHullName string `json:"reusedHullName"`
	// Read-only details of the Scientific object, only used during detail queries (not
	// to be provided on create/update operations).
	Scientific []SeradataSpacecraftDetailTupleResponseScientific `json:"scientific"`
	// Seradata sector (e.g. Commercial, Military, Civil/Other).
	Sector string `json:"sector"`
	// Read-only details of the SeradataCommDetails object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataCommDetails []SeradataSpacecraftDetailTupleResponseSeradataCommDetail `json:"seradataCommDetails"`
	// Read-only details of the SeradataEarlyWarning object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataEarlyWarning []SeradataSpacecraftDetailTupleResponseSeradataEarlyWarning `json:"seradataEarlyWarning"`
	// Read-only details of the SeradataNavigation object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataNavigation []SeradataSpacecraftDetailTupleResponseSeradataNavigation `json:"seradataNavigation"`
	// Read-only details of the SeradataOpticalPayload object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataOpticalPayload []SeradataSpacecraftDetailTupleResponseSeradataOpticalPayload `json:"seradataOpticalPayload"`
	// Read-only details of the SeradataRadarPayload object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataRadarPayload []SeradataSpacecraftDetailTupleResponseSeradataRadarPayload `json:"seradataRadarPayload"`
	// Read-only details of the SeradataSigIntPayload object, only used during detail
	// queries (not to be provided on create/update operations).
	SeradataSigIntPayload []SeradataSpacecraftDetailTupleResponseSeradataSigIntPayload `json:"seradataSigIntPayload"`
	// Spacecraft serial number.
	SerialNumber string `json:"serialNumber"`
	// Spacecraft stabilizer (e.g. 3-Axis, Gravity Gradiant, etc).
	Stabilizer string `json:"stabilizer"`
	// Spacecraft status (e.g. Inactive - Retired, Inactive - Re-entered, Active, etc).
	Status string `json:"status"`
	// Number of insurance claims for this spacecraft.
	TotalClaims int64 `json:"totalClaims"`
	// Number of fatalities related to this spacecraft.
	TotalFatalities int64 `json:"totalFatalities"`
	// Number of injuries related to this spacecraft.
	TotalInjuries int64 `json:"totalInjuries"`
	// Mass dry in kg.
	TotalPayloadPower float64 `json:"totalPayloadPower"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Youtube link of launch.
	YoutubeLaunchLink string `json:"youtubeLaunchLink"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                    respjson.Field
		DataMode                                 respjson.Field
		Name                                     respjson.Field
		Source                                   respjson.Field
		ID                                       respjson.Field
		AdditionalMissionsGroups                 respjson.Field
		Altitude                                 respjson.Field
		AnnualInsuredDepreciationFactor          respjson.Field
		AnnualInsuredDepreciationFactorEstimated respjson.Field
		Apogee                                   respjson.Field
		BusID                                    respjson.Field
		CapabilityLost                           respjson.Field
		CapacityLost                             respjson.Field
		CatalogNumber                            respjson.Field
		CollisionRiskCm                          respjson.Field
		CollisionRiskMm                          respjson.Field
		CombinedCostEstimated                    respjson.Field
		CombinedNewCost                          respjson.Field
		CommercialLaunch                         respjson.Field
		Constellation                            respjson.Field
		CostEstimated                            respjson.Field
		CreatedAt                                respjson.Field
		CreatedBy                                respjson.Field
		CubesatDispenserType                     respjson.Field
		CurrentAge                               respjson.Field
		DateOfObservation                        respjson.Field
		Description                              respjson.Field
		DesignLife                               respjson.Field
		DryMass                                  respjson.Field
		ExpectedLife                             respjson.Field
		GeoPosition                              respjson.Field
		IDOnOrbit                                respjson.Field
		Inclination                              respjson.Field
		InsuranceLossesTotal                     respjson.Field
		InsuranceNotes                           respjson.Field
		InsurancePremiumAtLaunch                 respjson.Field
		InsurancePremiumAtLaunchEstimated        respjson.Field
		InsuredAtLaunch                          respjson.Field
		InsuredValueAtLaunch                     respjson.Field
		InsuredValueLaunchEstimated              respjson.Field
		IntlNumber                               respjson.Field
		Lat                                      respjson.Field
		LaunchArranger                           respjson.Field
		LaunchArrangerCountry                    respjson.Field
		LaunchCharacteristic                     respjson.Field
		LaunchCost                               respjson.Field
		LaunchCostEstimated                      respjson.Field
		LaunchCountry                            respjson.Field
		LaunchDate                               respjson.Field
		LaunchDateRemarks                        respjson.Field
		LaunchID                                 respjson.Field
		LaunchMass                               respjson.Field
		LaunchNotes                              respjson.Field
		LaunchNumber                             respjson.Field
		LaunchProvider                           respjson.Field
		LaunchProviderCountry                    respjson.Field
		LaunchProviderFlightNumber               respjson.Field
		LaunchSiteID                             respjson.Field
		LaunchSiteName                           respjson.Field
		LaunchType                               respjson.Field
		LaunchVehicleID                          respjson.Field
		Leased                                   respjson.Field
		LifeLost                                 respjson.Field
		Lon                                      respjson.Field
		MassCategory                             respjson.Field
		NameAtLaunch                             respjson.Field
		NewCost                                  respjson.Field
		Notes                                    respjson.Field
		NumHumans                                respjson.Field
		OnOrbit                                  respjson.Field
		Operator                                 respjson.Field
		OperatorCountry                          respjson.Field
		OrbitCategory                            respjson.Field
		OrbitSubCategory                         respjson.Field
		OrderDate                                respjson.Field
		Origin                                   respjson.Field
		OrigNetwork                              respjson.Field
		Owner                                    respjson.Field
		OwnerCountry                             respjson.Field
		Perigee                                  respjson.Field
		Period                                   respjson.Field
		PrimaryMissionGroup                      respjson.Field
		PrimeManufacturerOrgID                   respjson.Field
		ProgramName                              respjson.Field
		Quantity                                 respjson.Field
		ReusableFlights                          respjson.Field
		ReusedHullName                           respjson.Field
		Scientific                               respjson.Field
		Sector                                   respjson.Field
		SeradataCommDetails                      respjson.Field
		SeradataEarlyWarning                     respjson.Field
		SeradataNavigation                       respjson.Field
		SeradataOpticalPayload                   respjson.Field
		SeradataRadarPayload                     respjson.Field
		SeradataSigIntPayload                    respjson.Field
		SerialNumber                             respjson.Field
		Stabilizer                               respjson.Field
		Status                                   respjson.Field
		TotalClaims                              respjson.Field
		TotalFatalities                          respjson.Field
		TotalInjuries                            respjson.Field
		TotalPayloadPower                        respjson.Field
		UpdatedAt                                respjson.Field
		UpdatedBy                                respjson.Field
		YoutubeLaunchLink                        respjson.Field
		ExtraFields                              map[string]respjson.Field
		raw                                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeradataSpacecraftDetailTupleResponseDataMode string

const (
	SeradataSpacecraftDetailTupleResponseDataModeReal      SeradataSpacecraftDetailTupleResponseDataMode = "REAL"
	SeradataSpacecraftDetailTupleResponseDataModeTest      SeradataSpacecraftDetailTupleResponseDataMode = "TEST"
	SeradataSpacecraftDetailTupleResponseDataModeSimulated SeradataSpacecraftDetailTupleResponseDataMode = "SIMULATED"
	SeradataSpacecraftDetailTupleResponseDataModeExercise  SeradataSpacecraftDetailTupleResponseDataMode = "EXERCISE"
)

// Model object representing on-orbit objects or satellites in the system.
type SeradataSpacecraftDetailTupleResponseOnOrbit struct {
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
	DataMode string `json:"dataMode,required"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Alternate name of the on-orbit object.
	AltName string `json:"altName"`
	// Read-only collection of antennas on this on-orbit object.
	Antennas []shared.OnorbitAntennaFull `json:"antennas"`
	// Read-only collection of batteries on this on-orbit object.
	Batteries []shared.OnorbitBatteryFull `json:"batteries"`
	// Category of the on-orbit object. (Unknown, On-Orbit, Decayed, Cataloged Without
	// State, Launch Nominal, Analyst Satellite, Cislunar, Lunar, Hyperbolic,
	// Heliocentric, Interplanetary, Lagrangian, Docked).
	//
	// Any of "Unknown", "On-Orbit", "Decayed", "Cataloged Without State", "Launch
	// Nominal", "Analyst Satellite", "Cislunar", "Lunar", "Hyperbolic",
	// "Heliocentric", "Interplanetary", "Lagrangian", "Docked".
	Category string `json:"category"`
	// Common name of the on-orbit object.
	CommonName string `json:"commonName"`
	// Constellation to which this satellite belongs.
	Constellation string `json:"constellation"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Date of decay.
	DecayDate time.Time `json:"decayDate" format:"date-time"`
	// For the public catalog, the idOnOrbit is typically the satellite number as a
	// string, but may be a UUID for analyst or other unknown or untracked satellites,
	// auto-generated by the system.
	IDOnOrbit string `json:"idOnOrbit"`
	// International Designator, typically of the format YYYYLLLAAA, where YYYY is the
	// launch year, LLL is the sequential launch number of that year, and AAA is an
	// optional launch piece designator for the launch.
	IntlDes string `json:"intlDes"`
	// Date of launch.
	LaunchDate time.Time `json:"launchDate" format:"date"`
	// Id of the associated launchSite entity.
	LaunchSiteID string `json:"launchSiteId"`
	// Estimated lifetime of the on-orbit payload, if known.
	LifetimeYears int64 `json:"lifetimeYears"`
	// Mission number of the on-orbit object.
	MissionNumber string `json:"missionNumber"`
	// Type of on-orbit object: ROCKET BODY, DEBRIS, PAYLOAD, PLATFORM, MANNED,
	// UNKNOWN.
	//
	// Any of "ROCKET BODY", "DEBRIS", "PAYLOAD", "PLATFORM", "MANNED", "UNKNOWN".
	ObjectType string `json:"objectType"`
	// Read-only collection of details for this on-orbit object.
	OnorbitDetails []shared.OnorbitDetailsFull `json:"onorbitDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only collection of solar arrays on this on-orbit object.
	SolarArrays []shared.OnorbitSolarArrayFull `json:"solarArrays"`
	// Read-only collection of thrusters (engines) on this on-orbit object.
	Thrusters []shared.OnorbitThrusterFull `json:"thrusters"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		SatNo                 respjson.Field
		Source                respjson.Field
		AltName               respjson.Field
		Antennas              respjson.Field
		Batteries             respjson.Field
		Category              respjson.Field
		CommonName            respjson.Field
		Constellation         respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecayDate             respjson.Field
		IDOnOrbit             respjson.Field
		IntlDes               respjson.Field
		LaunchDate            respjson.Field
		LaunchSiteID          respjson.Field
		LifetimeYears         respjson.Field
		MissionNumber         respjson.Field
		ObjectType            respjson.Field
		OnorbitDetails        respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SolarArrays           respjson.Field
		Thrusters             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailTupleResponseOnOrbit) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailTupleResponseOnOrbit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scientific or other data from Seradata.
type SeradataSpacecraftDetailTupleResponseScientific struct {
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
	DataMode string `json:"dataMode,required"`
	// Sensor name from sera data, e.g. SEM/MAG (SEM / Magnetometer).
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity shared.EntityFull `json:"entity"`
	// Frequency band, e.g. Gamma.
	FrequencyBand string `json:"frequencyBand"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// ID of the parent entity for this Scientific.
	IDEntity string `json:"idEntity"`
	// Unique identifier of the organization which manufactures this bus.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Notes associated with the payload.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Payload category, e.g. Magnetometer, Radiometer, Sensor, etc.
	PayloadCategory string `json:"payloadCategory"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		SpacecraftID          respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Entity                respjson.Field
		FrequencyBand         respjson.Field
		HostedForCompanyOrgID respjson.Field
		IDEntity              respjson.Field
		ManufacturerOrgID     respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PayloadCategory       respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailTupleResponseScientific) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailTupleResponseScientific) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Seradata-compiled information on communications payloads.
type SeradataSpacecraftDetailTupleResponseSeradataCommDetail struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band string `json:"band"`
	// Comm bandwidth in Mhz.
	Bandwidth float64 `json:"bandwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Effective isotropic radiated power in dB.
	Eirp float64 `json:"eirp"`
	// Comm estimated HtsTotalCapacity in Gbps.
	EstHtsTotalCapacity float64 `json:"estHtsTotalCapacity"`
	// Comm estimated HtsTotalUserDownlinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserDownlinkBandwidthPerBeam float64 `json:"estHtsTotalUserDownlinkBandwidthPerBeam"`
	// Comm estimated HtsTotalUserUplinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserUplinkBandwidthPerBeam float64 `json:"estHtsTotalUserUplinkBandwidthPerBeam"`
	// Comm gatewayDownlinkFrom in Ghz.
	GatewayDownlinkFrom float64 `json:"gatewayDownlinkFrom"`
	// Comm gatewayDownlinkTo in Ghz.
	GatewayDownlinkTo float64 `json:"gatewayDownlinkTo"`
	// Comm gatewayUplinkFrom in Ghz.
	GatewayUplinkFrom float64 `json:"gatewayUplinkFrom"`
	// Comm gatewayUplinkTo in Ghz.
	GatewayUplinkTo float64 `json:"gatewayUplinkTo"`
	// Comm hostedForCompanyOrgId.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// Comm htsNumUserSpotBeams.
	HtsNumUserSpotBeams int64 `json:"htsNumUserSpotBeams"`
	// Comm htsUserDownlinkBandwidthPerBeam in Mhz.
	HtsUserDownlinkBandwidthPerBeam float64 `json:"htsUserDownlinkBandwidthPerBeam"`
	// Comm htsUserUplinkBandwidthPerBeam in Mhz.
	HtsUserUplinkBandwidthPerBeam float64 `json:"htsUserUplinkBandwidthPerBeam"`
	// UUID of the parent Comm record.
	IDComm string `json:"idComm"`
	// Comm manufacturerOrgId.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Comm num36MhzEquivalentTransponders.
	Num36MhzEquivalentTransponders int64 `json:"num36MhzEquivalentTransponders"`
	// Comm numOperationalTransponders.
	NumOperationalTransponders int64 `json:"numOperationalTransponders"`
	// Comm numSpareTransponders.
	NumSpareTransponders int64 `json:"numSpareTransponders"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Payload notes.
	PayloadNotes string `json:"payloadNotes"`
	// Comm polarization.
	Polarization string `json:"polarization"`
	// Solid state power amplifier, in Watts.
	SolidStatePowerAmp float64 `json:"solidStatePowerAmp"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId"`
	// Comm tradeLeaseOrgId.
	TradeLeaseOrgID string `json:"tradeLeaseOrgId"`
	// Comm travelingWaveTubeAmplifier in Watts.
	TravelingWaveTubeAmplifier float64 `json:"travelingWaveTubeAmplifier"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Comm userDownlinkFrom in Ghz.
	UserDownlinkFrom float64 `json:"userDownlinkFrom"`
	// Comm userDownlinkTo in Ghz.
	UserDownlinkTo float64 `json:"userDownlinkTo"`
	// Comm userUplinkFrom in Ghz.
	UserUplinkFrom float64 `json:"userUplinkFrom"`
	// Comm userUplinkTo in Ghz.
	UserUplinkTo float64 `json:"userUplinkTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                   respjson.Field
		DataMode                                respjson.Field
		Source                                  respjson.Field
		ID                                      respjson.Field
		Band                                    respjson.Field
		Bandwidth                               respjson.Field
		CreatedAt                               respjson.Field
		CreatedBy                               respjson.Field
		Eirp                                    respjson.Field
		EstHtsTotalCapacity                     respjson.Field
		EstHtsTotalUserDownlinkBandwidthPerBeam respjson.Field
		EstHtsTotalUserUplinkBandwidthPerBeam   respjson.Field
		GatewayDownlinkFrom                     respjson.Field
		GatewayDownlinkTo                       respjson.Field
		GatewayUplinkFrom                       respjson.Field
		GatewayUplinkTo                         respjson.Field
		HostedForCompanyOrgID                   respjson.Field
		HtsNumUserSpotBeams                     respjson.Field
		HtsUserDownlinkBandwidthPerBeam         respjson.Field
		HtsUserUplinkBandwidthPerBeam           respjson.Field
		IDComm                                  respjson.Field
		ManufacturerOrgID                       respjson.Field
		Num36MhzEquivalentTransponders          respjson.Field
		NumOperationalTransponders              respjson.Field
		NumSpareTransponders                    respjson.Field
		Origin                                  respjson.Field
		OrigNetwork                             respjson.Field
		PayloadNotes                            respjson.Field
		Polarization                            respjson.Field
		SolidStatePowerAmp                      respjson.Field
		SpacecraftID                            respjson.Field
		TradeLeaseOrgID                         respjson.Field
		TravelingWaveTubeAmplifier              respjson.Field
		UpdatedAt                               respjson.Field
		UpdatedBy                               respjson.Field
		UserDownlinkFrom                        respjson.Field
		UserDownlinkTo                          respjson.Field
		UserUplinkFrom                          respjson.Field
		UserUplinkTo                            respjson.Field
		ExtraFields                             map[string]respjson.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailTupleResponseSeradataCommDetail) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailTupleResponseSeradataCommDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for an early warning payload from Seradata.
type SeradataSpacecraftDetailTupleResponseSeradataEarlyWarning struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Best resolution for this IR in meters.
	BestResolution float64 `json:"bestResolution"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Is the sensor Earth Pointing.
	EarthPointing bool `json:"earthPointing"`
	// Frequency Limits for this IR.
	FrequencyLimits string `json:"frequencyLimits"`
	// Ground Station Locations for this IR.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this IR.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the parent IR record.
	IDIr string `json:"idIR"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Missile Launch Phase Detection Ability.
	MissileLaunchPhaseDetectionAbility string `json:"missileLaunchPhaseDetectionAbility"`
	// Sensor name from Seradata, e.g. Infra red telescope, Schmidt Telescope, etc.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID string `json:"partnerSpacecraftId"`
	// Payload notes.
	PayloadNotes string `json:"payloadNotes"`
	// Spectral Bands, e.g. Infra-Red.
	SpectralBands string `json:"spectralBands"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking              respjson.Field
		DataMode                           respjson.Field
		Source                             respjson.Field
		SpacecraftID                       respjson.Field
		ID                                 respjson.Field
		BestResolution                     respjson.Field
		CreatedAt                          respjson.Field
		CreatedBy                          respjson.Field
		EarthPointing                      respjson.Field
		FrequencyLimits                    respjson.Field
		GroundStationLocations             respjson.Field
		GroundStations                     respjson.Field
		HostedForCompanyOrgID              respjson.Field
		IDIr                               respjson.Field
		ManufacturerOrgID                  respjson.Field
		MissileLaunchPhaseDetectionAbility respjson.Field
		Name                               respjson.Field
		Origin                             respjson.Field
		OrigNetwork                        respjson.Field
		PartnerSpacecraftID                respjson.Field
		PayloadNotes                       respjson.Field
		SpectralBands                      respjson.Field
		UpdatedAt                          respjson.Field
		UpdatedBy                          respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailTupleResponseSeradataEarlyWarning) RawJSON() string {
	return r.JSON.raw
}
func (r *SeradataSpacecraftDetailTupleResponseSeradataEarlyWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for a navigation payload from Seradata.
type SeradataSpacecraftDetailTupleResponseSeradataNavigation struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Area of coverage, e.g. Worldwide, India, etc.
	AreaCoverage string `json:"areaCoverage"`
	// Type of clock, e.g. Rubidium, Hydrogen Maser, etc.
	ClockType string `json:"clockType"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the parent Navigation record.
	IDNavigation string `json:"idNavigation"`
	// Location accuracy in meters.
	LocationAccuracy float64 `json:"locationAccuracy"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Frequency for this payload.
	ModeFrequency string `json:"modeFrequency"`
	// Modes of operation.
	Modes string `json:"modes"`
	// Sensor name from Seradata, e.g. WAAS GEO-5, etc.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID string `json:"partnerSpacecraftId"`
	// Navigation payload type, e.g. WAAS, GAGAN, etc.
	PayloadType string `json:"payloadType"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		SpacecraftID          respjson.Field
		ID                    respjson.Field
		AreaCoverage          respjson.Field
		ClockType             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		HostedForCompanyOrgID respjson.Field
		IDNavigation          respjson.Field
		LocationAccuracy      respjson.Field
		ManufacturerOrgID     respjson.Field
		ModeFrequency         respjson.Field
		Modes                 respjson.Field
		Name                  respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PartnerSpacecraftID   respjson.Field
		PayloadType           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailTupleResponseSeradataNavigation) RawJSON() string { return r.JSON.raw }
func (r *SeradataSpacecraftDetailTupleResponseSeradataNavigation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for an optical payload from Seradata.
type SeradataSpacecraftDetailTupleResponseSeradataOpticalPayload struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Best resolution.
	BestResolution float64 `json:"bestResolution"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Field of regard of this radar in degrees.
	FieldOfRegard float64 `json:"fieldOfRegard"`
	// Field of view of this radar in kilometers.
	FieldOfView float64 `json:"fieldOfView"`
	// Ground Station Locations for this payload.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this payload.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the Sensor record.
	IDSensor string `json:"idSensor"`
	// Imaging category for this payload, e.g. Multispectral, Infrared, Panchromatic.
	ImagingPayloadCategory string `json:"imagingPayloadCategory"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Sensor name from Seradata, e.g. TOURNESOL, MESSR (Multispectral Self-Scanning
	// Radiometer), AWFI, etc.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Number of film return canisters.
	NumberOfFilmReturnCanisters int64 `json:"numberOfFilmReturnCanisters"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Point method for this radar, e.g. Spacecraft.
	PointingMethod string `json:"pointingMethod"`
	// Recorder size.
	RecorderSize string `json:"recorderSize"`
	// Spectral Band supported by this payload, e.g. Green, Red, Mid-wave infrared,
	// etc.
	SpectralBand string `json:"spectralBand"`
	// Frequency limit for this payload, e.g. 0.51 - 0.59.
	SpectralFrequencyLimits string `json:"spectralFrequencyLimits"`
	// Swath width in kilometers.
	SwathWidth float64 `json:"swathWidth"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking       respjson.Field
		DataMode                    respjson.Field
		Source                      respjson.Field
		SpacecraftID                respjson.Field
		ID                          respjson.Field
		BestResolution              respjson.Field
		CreatedAt                   respjson.Field
		CreatedBy                   respjson.Field
		FieldOfRegard               respjson.Field
		FieldOfView                 respjson.Field
		GroundStationLocations      respjson.Field
		GroundStations              respjson.Field
		HostedForCompanyOrgID       respjson.Field
		IDSensor                    respjson.Field
		ImagingPayloadCategory      respjson.Field
		ManufacturerOrgID           respjson.Field
		Name                        respjson.Field
		Notes                       respjson.Field
		NumberOfFilmReturnCanisters respjson.Field
		Origin                      respjson.Field
		OrigNetwork                 respjson.Field
		PointingMethod              respjson.Field
		RecorderSize                respjson.Field
		SpectralBand                respjson.Field
		SpectralFrequencyLimits     respjson.Field
		SwathWidth                  respjson.Field
		UpdatedAt                   respjson.Field
		UpdatedBy                   respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailTupleResponseSeradataOpticalPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *SeradataSpacecraftDetailTupleResponseSeradataOpticalPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for an radar payload from Seradata.
type SeradataSpacecraftDetailTupleResponseSeradataRadarPayload struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Radar bandwidth in mega hertz.
	Bandwidth float64 `json:"bandwidth"`
	// Best resolution in meters.
	BestResolution float64 `json:"bestResolution"`
	// Radar category, e.g. SAR, Surface Search, etc.
	Category string `json:"category"`
	// Constellation interferometric capability.
	ConstellationInterferometricCapability string `json:"constellationInterferometricCapability"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Duty cycle.
	DutyCycle string `json:"dutyCycle"`
	// Field of regard of this radar in degrees.
	FieldOfRegard float64 `json:"fieldOfRegard"`
	// Field of view of this radar in kilometers.
	FieldOfView float64 `json:"fieldOfView"`
	// Frequency in giga hertz.
	Frequency float64 `json:"frequency"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	FrequencyBand string `json:"frequencyBand"`
	// Ground Station Locations for this payload.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this payload.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the Sensor record.
	IDSensor string `json:"idSensor"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Sensor name from Seradata, e.g. ALT (Radar Altimeter), COSI (Corea SAR
	// Instrument), etc.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Partner seradata-spacecraft.
	PartnerSpacecraft string `json:"partnerSpacecraft"`
	// Point method for this radar, e.g. Spacecraft.
	PointingMethod string `json:"pointingMethod"`
	// Receive polarization, e.g. Lin Dual, Lin vert, etc.
	ReceivePolarization string `json:"receivePolarization"`
	// Recorder size, e.g. 256.
	RecorderSize string `json:"recorderSize"`
	// Swath width in kilometers.
	SwathWidth float64 `json:"swathWidth"`
	// Transmit polarization, e.g. Lin Dual, Lin vert, etc.
	TransmitPolarization string `json:"transmitPolarization"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Wave length in meters.
	WaveLength float64 `json:"waveLength"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                  respjson.Field
		DataMode                               respjson.Field
		Source                                 respjson.Field
		SpacecraftID                           respjson.Field
		ID                                     respjson.Field
		Bandwidth                              respjson.Field
		BestResolution                         respjson.Field
		Category                               respjson.Field
		ConstellationInterferometricCapability respjson.Field
		CreatedAt                              respjson.Field
		CreatedBy                              respjson.Field
		DutyCycle                              respjson.Field
		FieldOfRegard                          respjson.Field
		FieldOfView                            respjson.Field
		Frequency                              respjson.Field
		FrequencyBand                          respjson.Field
		GroundStationLocations                 respjson.Field
		GroundStations                         respjson.Field
		HostedForCompanyOrgID                  respjson.Field
		IDSensor                               respjson.Field
		ManufacturerOrgID                      respjson.Field
		Name                                   respjson.Field
		Notes                                  respjson.Field
		Origin                                 respjson.Field
		OrigNetwork                            respjson.Field
		PartnerSpacecraft                      respjson.Field
		PointingMethod                         respjson.Field
		ReceivePolarization                    respjson.Field
		RecorderSize                           respjson.Field
		SwathWidth                             respjson.Field
		TransmitPolarization                   respjson.Field
		UpdatedAt                              respjson.Field
		UpdatedBy                              respjson.Field
		WaveLength                             respjson.Field
		ExtraFields                            map[string]respjson.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailTupleResponseSeradataRadarPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *SeradataSpacecraftDetailTupleResponseSeradataRadarPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for an sigint payload from Seradata.
type SeradataSpacecraftDetailTupleResponseSeradataSigIntPayload struct {
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
	DataMode string `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Frequency coverage for this payload.
	FrequencyCoverage string `json:"frequencyCoverage"`
	// Ground Station Locations for this payload.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this payload.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the Sensor record.
	IDSensor string `json:"idSensor"`
	// Intercept parameters.
	InterceptParameters string `json:"interceptParameters"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Sensor name from Seradata.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Positional Accuracy for this payload.
	PositionalAccuracy string `json:"positionalAccuracy"`
	// Swath Width in kilometers.
	SwathWidth float64 `json:"swathWidth"`
	// SIGINT Payload type, e.g. Comint, Elint, etc.
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		Source                 respjson.Field
		SpacecraftID           respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		FrequencyCoverage      respjson.Field
		GroundStationLocations respjson.Field
		GroundStations         respjson.Field
		HostedForCompanyOrgID  respjson.Field
		IDSensor               respjson.Field
		InterceptParameters    respjson.Field
		ManufacturerOrgID      respjson.Field
		Name                   respjson.Field
		Notes                  respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		PositionalAccuracy     respjson.Field
		SwathWidth             respjson.Field
		Type                   respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSpacecraftDetailTupleResponseSeradataSigIntPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *SeradataSpacecraftDetailTupleResponseSeradataSigIntPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SeradataSpacecraftDetailNewParams struct {
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
	DataMode SeradataSpacecraftDetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Spacecraft name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Spacecraft additional missions and groups.
	AdditionalMissionsGroups param.Opt[string] `json:"additionalMissionsGroups,omitzero"`
	// Spacecraft latest altitude in km.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// Annual insured depreciaion factor as a percent fraction.
	AnnualInsuredDepreciationFactor param.Opt[float64] `json:"annualInsuredDepreciationFactor,omitzero"`
	// Boolean indicating if the spacecraft annualInsuredDepreciationFactor is
	// estimated.
	AnnualInsuredDepreciationFactorEstimated param.Opt[bool] `json:"annualInsuredDepreciationFactorEstimated,omitzero"`
	// Apogee in km.
	Apogee param.Opt[float64] `json:"apogee,omitzero"`
	// Spacecraft Bus ID.
	BusID param.Opt[string] `json:"busId,omitzero"`
	// Total capability lost as a percent fraction.
	CapabilityLost param.Opt[float64] `json:"capabilityLost,omitzero"`
	// Total capacity lost as a percent fraction.
	CapacityLost param.Opt[float64] `json:"capacityLost,omitzero"`
	// NORAD satellite number if available.
	CatalogNumber param.Opt[int64] `json:"catalogNumber,omitzero"`
	// Spacecraft collision risk 1cm sqm latest.
	CollisionRiskCm param.Opt[float64] `json:"collisionRiskCM,omitzero"`
	// Spacecraft collision risk 1mm sqm latest.
	CollisionRiskMm param.Opt[float64] `json:"collisionRiskMM,omitzero"`
	// Boolean indicating if the spacecraft combined new cost is estimated.
	CombinedCostEstimated param.Opt[bool] `json:"combinedCostEstimated,omitzero"`
	// Combined cost of spacecraft at new in M USD.
	CombinedNewCost param.Opt[float64] `json:"combinedNewCost,omitzero"`
	// Boolean indicating if the launch was commercial.
	CommercialLaunch param.Opt[bool] `json:"commercialLaunch,omitzero"`
	// Spacecraft constellation.
	Constellation param.Opt[string] `json:"constellation,omitzero"`
	// Boolean indicating if the spacecraft cost is estimated.
	CostEstimated param.Opt[bool] `json:"costEstimated,omitzero"`
	// Cubesat dispenser type.
	CubesatDispenserType param.Opt[string] `json:"cubesatDispenserType,omitzero"`
	// Current age in years.
	CurrentAge param.Opt[float64] `json:"currentAge,omitzero"`
	// Spacecraft date of observation.
	DateOfObservation param.Opt[time.Time] `json:"dateOfObservation,omitzero" format:"date-time"`
	// Description associated with the spacecraft.
	Description param.Opt[string] `json:"description,omitzero"`
	// Spacecraft design life in days.
	DesignLife param.Opt[int64] `json:"designLife,omitzero"`
	// Mass dry in kg.
	DryMass param.Opt[float64] `json:"dryMass,omitzero"`
	// Spacecraft expected life in days.
	ExpectedLife param.Opt[int64] `json:"expectedLife,omitzero"`
	// WGS84 longitude of the spacecraft’s latest GEO position, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	GeoPosition param.Opt[float64] `json:"geoPosition,omitzero"`
	// UUID of the parent Onorbit record, if available.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Seradata provided inclination in degrees.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Spacecraft total insurance losses as a fraction.
	InsuranceLossesTotal param.Opt[float64] `json:"insuranceLossesTotal,omitzero"`
	// Insurance notes for the spacecraft.
	InsuranceNotes param.Opt[string] `json:"insuranceNotes,omitzero"`
	// Insurance premium at launch in M USD.
	InsurancePremiumAtLaunch param.Opt[float64] `json:"insurancePremiumAtLaunch,omitzero"`
	// Boolean indicating if the spacecraft insurancePremiumAtLaunch is estimated.
	InsurancePremiumAtLaunchEstimated param.Opt[bool] `json:"insurancePremiumAtLaunchEstimated,omitzero"`
	// Boolean indicating if the spacecraft was insured at launch.
	InsuredAtLaunch param.Opt[bool] `json:"insuredAtLaunch,omitzero"`
	// Insured value of spacecraft at launch in M USD.
	InsuredValueAtLaunch param.Opt[float64] `json:"insuredValueAtLaunch,omitzero"`
	// Boolean indicating if the spacecraft insured value at launch is estimated.
	InsuredValueLaunchEstimated param.Opt[bool] `json:"insuredValueLaunchEstimated,omitzero"`
	// Seradata international number.
	IntlNumber param.Opt[string] `json:"intlNumber,omitzero"`
	// Spacecraft latest latitude in degrees.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Spacecraft launch arranger.
	LaunchArranger param.Opt[string] `json:"launchArranger,omitzero"`
	// Spacecraft launch arranger country.
	LaunchArrangerCountry param.Opt[string] `json:"launchArrangerCountry,omitzero"`
	// Seradata launch characteristic (e.g. Expendable, Reusable (New), etc).
	LaunchCharacteristic param.Opt[string] `json:"launchCharacteristic,omitzero"`
	// Cost of launch in M USD.
	LaunchCost param.Opt[float64] `json:"launchCost,omitzero"`
	// Boolean indicating if the spacecraft launch cost is estimated.
	LaunchCostEstimated param.Opt[bool] `json:"launchCostEstimated,omitzero"`
	// Seradata launch country.
	LaunchCountry param.Opt[string] `json:"launchCountry,omitzero"`
	// Launch date.
	LaunchDate param.Opt[time.Time] `json:"launchDate,omitzero" format:"date-time"`
	// Seradata remarks on launch date.
	LaunchDateRemarks param.Opt[string] `json:"launchDateRemarks,omitzero"`
	// Seradata launch ID.
	LaunchID param.Opt[string] `json:"launchId,omitzero"`
	// Mass at launch in kg.
	LaunchMass param.Opt[float64] `json:"launchMass,omitzero"`
	// Insurance notes for the spacecraft.
	LaunchNotes param.Opt[string] `json:"launchNotes,omitzero"`
	// Seradata launch number.
	LaunchNumber param.Opt[string] `json:"launchNumber,omitzero"`
	// Seradata launch provider.
	LaunchProvider param.Opt[string] `json:"launchProvider,omitzero"`
	// Seradata launch provider country.
	LaunchProviderCountry param.Opt[string] `json:"launchProviderCountry,omitzero"`
	// Seradata launch vehicle family.
	LaunchProviderFlightNumber param.Opt[string] `json:"launchProviderFlightNumber,omitzero"`
	// Seradata Launch Site ID.
	LaunchSiteID param.Opt[string] `json:"launchSiteId,omitzero"`
	// Launch Site Name.
	LaunchSiteName param.Opt[string] `json:"launchSiteName,omitzero"`
	// Seradata launch type (e.g. Launched, Future, etc).
	LaunchType param.Opt[string] `json:"launchType,omitzero"`
	// Seradata launch ID.
	LaunchVehicleID param.Opt[string] `json:"launchVehicleId,omitzero"`
	// Boolean indicating if the spacecraft was leased.
	Leased param.Opt[bool] `json:"leased,omitzero"`
	// Spacecraft life lost as a percent fraction.
	LifeLost param.Opt[float64] `json:"lifeLost,omitzero"`
	// Spacecraft latest longitude in degrees.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Mass category (e.g. 2500 - 3500kg - Large Satellite, 10 - 100 kg -
	// Microsatellite, etc).
	MassCategory param.Opt[string] `json:"massCategory,omitzero"`
	// Spacecraft name at launch.
	NameAtLaunch param.Opt[string] `json:"nameAtLaunch,omitzero"`
	// Cost of spacecraft at new in M USD.
	NewCost param.Opt[float64] `json:"newCost,omitzero"`
	// Notes on the spacecraft.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of humans carried on spacecraft.
	NumHumans param.Opt[int64] `json:"numHumans,omitzero"`
	// Spacecraft operator name.
	Operator param.Opt[string] `json:"operator,omitzero"`
	// Spacecraft operator country.
	OperatorCountry param.Opt[string] `json:"operatorCountry,omitzero"`
	// Spacecraft orbit category (e.g GEO, LEO, etc).
	OrbitCategory param.Opt[string] `json:"orbitCategory,omitzero"`
	// Spacecraft sub orbit category (e.g LEO - Sun-synchronous, Geostationary, etc).
	OrbitSubCategory param.Opt[string] `json:"orbitSubCategory,omitzero"`
	// Spacecraft order date.
	OrderDate param.Opt[time.Time] `json:"orderDate,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Spacecraft owner name.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// Spacecraft owner country.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	// Perigee in km.
	Perigee param.Opt[float64] `json:"perigee,omitzero"`
	// Spacecraft period in minutes.
	Period param.Opt[float64] `json:"period,omitzero"`
	// Spacecraft primary mission and group.
	PrimaryMissionGroup param.Opt[string] `json:"primaryMissionGroup,omitzero"`
	// UUID of the prime manufacturer organization, if available.
	PrimeManufacturerOrgID param.Opt[string] `json:"primeManufacturerOrgId,omitzero"`
	// Spacecraft program name.
	ProgramName param.Opt[string] `json:"programName,omitzero"`
	// Spacecraft quantity.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Spacecraft reusable flights.
	ReusableFlights param.Opt[string] `json:"reusableFlights,omitzero"`
	// Spacecraft reused hull name.
	ReusedHullName param.Opt[string] `json:"reusedHullName,omitzero"`
	// Seradata sector (e.g. Commercial, Military, Civil/Other).
	Sector param.Opt[string] `json:"sector,omitzero"`
	// Spacecraft serial number.
	SerialNumber param.Opt[string] `json:"serialNumber,omitzero"`
	// Spacecraft stabilizer (e.g. 3-Axis, Gravity Gradiant, etc).
	Stabilizer param.Opt[string] `json:"stabilizer,omitzero"`
	// Spacecraft status (e.g. Inactive - Retired, Inactive - Re-entered, Active, etc).
	Status param.Opt[string] `json:"status,omitzero"`
	// Number of insurance claims for this spacecraft.
	TotalClaims param.Opt[int64] `json:"totalClaims,omitzero"`
	// Number of fatalities related to this spacecraft.
	TotalFatalities param.Opt[int64] `json:"totalFatalities,omitzero"`
	// Number of injuries related to this spacecraft.
	TotalInjuries param.Opt[int64] `json:"totalInjuries,omitzero"`
	// Mass dry in kg.
	TotalPayloadPower param.Opt[float64] `json:"totalPayloadPower,omitzero"`
	// Youtube link of launch.
	YoutubeLaunchLink param.Opt[string] `json:"youtubeLaunchLink,omitzero"`
	paramObj
}

func (r SeradataSpacecraftDetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataSpacecraftDetailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeradataSpacecraftDetailNewParams) UnmarshalJSON(data []byte) error {
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
type SeradataSpacecraftDetailNewParamsDataMode string

const (
	SeradataSpacecraftDetailNewParamsDataModeReal      SeradataSpacecraftDetailNewParamsDataMode = "REAL"
	SeradataSpacecraftDetailNewParamsDataModeTest      SeradataSpacecraftDetailNewParamsDataMode = "TEST"
	SeradataSpacecraftDetailNewParamsDataModeSimulated SeradataSpacecraftDetailNewParamsDataMode = "SIMULATED"
	SeradataSpacecraftDetailNewParamsDataModeExercise  SeradataSpacecraftDetailNewParamsDataMode = "EXERCISE"
)

type SeradataSpacecraftDetailUpdateParams struct {
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
	DataMode SeradataSpacecraftDetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Spacecraft name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Spacecraft additional missions and groups.
	AdditionalMissionsGroups param.Opt[string] `json:"additionalMissionsGroups,omitzero"`
	// Spacecraft latest altitude in km.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// Annual insured depreciaion factor as a percent fraction.
	AnnualInsuredDepreciationFactor param.Opt[float64] `json:"annualInsuredDepreciationFactor,omitzero"`
	// Boolean indicating if the spacecraft annualInsuredDepreciationFactor is
	// estimated.
	AnnualInsuredDepreciationFactorEstimated param.Opt[bool] `json:"annualInsuredDepreciationFactorEstimated,omitzero"`
	// Apogee in km.
	Apogee param.Opt[float64] `json:"apogee,omitzero"`
	// Spacecraft Bus ID.
	BusID param.Opt[string] `json:"busId,omitzero"`
	// Total capability lost as a percent fraction.
	CapabilityLost param.Opt[float64] `json:"capabilityLost,omitzero"`
	// Total capacity lost as a percent fraction.
	CapacityLost param.Opt[float64] `json:"capacityLost,omitzero"`
	// NORAD satellite number if available.
	CatalogNumber param.Opt[int64] `json:"catalogNumber,omitzero"`
	// Spacecraft collision risk 1cm sqm latest.
	CollisionRiskCm param.Opt[float64] `json:"collisionRiskCM,omitzero"`
	// Spacecraft collision risk 1mm sqm latest.
	CollisionRiskMm param.Opt[float64] `json:"collisionRiskMM,omitzero"`
	// Boolean indicating if the spacecraft combined new cost is estimated.
	CombinedCostEstimated param.Opt[bool] `json:"combinedCostEstimated,omitzero"`
	// Combined cost of spacecraft at new in M USD.
	CombinedNewCost param.Opt[float64] `json:"combinedNewCost,omitzero"`
	// Boolean indicating if the launch was commercial.
	CommercialLaunch param.Opt[bool] `json:"commercialLaunch,omitzero"`
	// Spacecraft constellation.
	Constellation param.Opt[string] `json:"constellation,omitzero"`
	// Boolean indicating if the spacecraft cost is estimated.
	CostEstimated param.Opt[bool] `json:"costEstimated,omitzero"`
	// Cubesat dispenser type.
	CubesatDispenserType param.Opt[string] `json:"cubesatDispenserType,omitzero"`
	// Current age in years.
	CurrentAge param.Opt[float64] `json:"currentAge,omitzero"`
	// Spacecraft date of observation.
	DateOfObservation param.Opt[time.Time] `json:"dateOfObservation,omitzero" format:"date-time"`
	// Description associated with the spacecraft.
	Description param.Opt[string] `json:"description,omitzero"`
	// Spacecraft design life in days.
	DesignLife param.Opt[int64] `json:"designLife,omitzero"`
	// Mass dry in kg.
	DryMass param.Opt[float64] `json:"dryMass,omitzero"`
	// Spacecraft expected life in days.
	ExpectedLife param.Opt[int64] `json:"expectedLife,omitzero"`
	// WGS84 longitude of the spacecraft’s latest GEO position, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	GeoPosition param.Opt[float64] `json:"geoPosition,omitzero"`
	// UUID of the parent Onorbit record, if available.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Seradata provided inclination in degrees.
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Spacecraft total insurance losses as a fraction.
	InsuranceLossesTotal param.Opt[float64] `json:"insuranceLossesTotal,omitzero"`
	// Insurance notes for the spacecraft.
	InsuranceNotes param.Opt[string] `json:"insuranceNotes,omitzero"`
	// Insurance premium at launch in M USD.
	InsurancePremiumAtLaunch param.Opt[float64] `json:"insurancePremiumAtLaunch,omitzero"`
	// Boolean indicating if the spacecraft insurancePremiumAtLaunch is estimated.
	InsurancePremiumAtLaunchEstimated param.Opt[bool] `json:"insurancePremiumAtLaunchEstimated,omitzero"`
	// Boolean indicating if the spacecraft was insured at launch.
	InsuredAtLaunch param.Opt[bool] `json:"insuredAtLaunch,omitzero"`
	// Insured value of spacecraft at launch in M USD.
	InsuredValueAtLaunch param.Opt[float64] `json:"insuredValueAtLaunch,omitzero"`
	// Boolean indicating if the spacecraft insured value at launch is estimated.
	InsuredValueLaunchEstimated param.Opt[bool] `json:"insuredValueLaunchEstimated,omitzero"`
	// Seradata international number.
	IntlNumber param.Opt[string] `json:"intlNumber,omitzero"`
	// Spacecraft latest latitude in degrees.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Spacecraft launch arranger.
	LaunchArranger param.Opt[string] `json:"launchArranger,omitzero"`
	// Spacecraft launch arranger country.
	LaunchArrangerCountry param.Opt[string] `json:"launchArrangerCountry,omitzero"`
	// Seradata launch characteristic (e.g. Expendable, Reusable (New), etc).
	LaunchCharacteristic param.Opt[string] `json:"launchCharacteristic,omitzero"`
	// Cost of launch in M USD.
	LaunchCost param.Opt[float64] `json:"launchCost,omitzero"`
	// Boolean indicating if the spacecraft launch cost is estimated.
	LaunchCostEstimated param.Opt[bool] `json:"launchCostEstimated,omitzero"`
	// Seradata launch country.
	LaunchCountry param.Opt[string] `json:"launchCountry,omitzero"`
	// Launch date.
	LaunchDate param.Opt[time.Time] `json:"launchDate,omitzero" format:"date-time"`
	// Seradata remarks on launch date.
	LaunchDateRemarks param.Opt[string] `json:"launchDateRemarks,omitzero"`
	// Seradata launch ID.
	LaunchID param.Opt[string] `json:"launchId,omitzero"`
	// Mass at launch in kg.
	LaunchMass param.Opt[float64] `json:"launchMass,omitzero"`
	// Insurance notes for the spacecraft.
	LaunchNotes param.Opt[string] `json:"launchNotes,omitzero"`
	// Seradata launch number.
	LaunchNumber param.Opt[string] `json:"launchNumber,omitzero"`
	// Seradata launch provider.
	LaunchProvider param.Opt[string] `json:"launchProvider,omitzero"`
	// Seradata launch provider country.
	LaunchProviderCountry param.Opt[string] `json:"launchProviderCountry,omitzero"`
	// Seradata launch vehicle family.
	LaunchProviderFlightNumber param.Opt[string] `json:"launchProviderFlightNumber,omitzero"`
	// Seradata Launch Site ID.
	LaunchSiteID param.Opt[string] `json:"launchSiteId,omitzero"`
	// Launch Site Name.
	LaunchSiteName param.Opt[string] `json:"launchSiteName,omitzero"`
	// Seradata launch type (e.g. Launched, Future, etc).
	LaunchType param.Opt[string] `json:"launchType,omitzero"`
	// Seradata launch ID.
	LaunchVehicleID param.Opt[string] `json:"launchVehicleId,omitzero"`
	// Boolean indicating if the spacecraft was leased.
	Leased param.Opt[bool] `json:"leased,omitzero"`
	// Spacecraft life lost as a percent fraction.
	LifeLost param.Opt[float64] `json:"lifeLost,omitzero"`
	// Spacecraft latest longitude in degrees.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Mass category (e.g. 2500 - 3500kg - Large Satellite, 10 - 100 kg -
	// Microsatellite, etc).
	MassCategory param.Opt[string] `json:"massCategory,omitzero"`
	// Spacecraft name at launch.
	NameAtLaunch param.Opt[string] `json:"nameAtLaunch,omitzero"`
	// Cost of spacecraft at new in M USD.
	NewCost param.Opt[float64] `json:"newCost,omitzero"`
	// Notes on the spacecraft.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of humans carried on spacecraft.
	NumHumans param.Opt[int64] `json:"numHumans,omitzero"`
	// Spacecraft operator name.
	Operator param.Opt[string] `json:"operator,omitzero"`
	// Spacecraft operator country.
	OperatorCountry param.Opt[string] `json:"operatorCountry,omitzero"`
	// Spacecraft orbit category (e.g GEO, LEO, etc).
	OrbitCategory param.Opt[string] `json:"orbitCategory,omitzero"`
	// Spacecraft sub orbit category (e.g LEO - Sun-synchronous, Geostationary, etc).
	OrbitSubCategory param.Opt[string] `json:"orbitSubCategory,omitzero"`
	// Spacecraft order date.
	OrderDate param.Opt[time.Time] `json:"orderDate,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Spacecraft owner name.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// Spacecraft owner country.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	// Perigee in km.
	Perigee param.Opt[float64] `json:"perigee,omitzero"`
	// Spacecraft period in minutes.
	Period param.Opt[float64] `json:"period,omitzero"`
	// Spacecraft primary mission and group.
	PrimaryMissionGroup param.Opt[string] `json:"primaryMissionGroup,omitzero"`
	// UUID of the prime manufacturer organization, if available.
	PrimeManufacturerOrgID param.Opt[string] `json:"primeManufacturerOrgId,omitzero"`
	// Spacecraft program name.
	ProgramName param.Opt[string] `json:"programName,omitzero"`
	// Spacecraft quantity.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Spacecraft reusable flights.
	ReusableFlights param.Opt[string] `json:"reusableFlights,omitzero"`
	// Spacecraft reused hull name.
	ReusedHullName param.Opt[string] `json:"reusedHullName,omitzero"`
	// Seradata sector (e.g. Commercial, Military, Civil/Other).
	Sector param.Opt[string] `json:"sector,omitzero"`
	// Spacecraft serial number.
	SerialNumber param.Opt[string] `json:"serialNumber,omitzero"`
	// Spacecraft stabilizer (e.g. 3-Axis, Gravity Gradiant, etc).
	Stabilizer param.Opt[string] `json:"stabilizer,omitzero"`
	// Spacecraft status (e.g. Inactive - Retired, Inactive - Re-entered, Active, etc).
	Status param.Opt[string] `json:"status,omitzero"`
	// Number of insurance claims for this spacecraft.
	TotalClaims param.Opt[int64] `json:"totalClaims,omitzero"`
	// Number of fatalities related to this spacecraft.
	TotalFatalities param.Opt[int64] `json:"totalFatalities,omitzero"`
	// Number of injuries related to this spacecraft.
	TotalInjuries param.Opt[int64] `json:"totalInjuries,omitzero"`
	// Mass dry in kg.
	TotalPayloadPower param.Opt[float64] `json:"totalPayloadPower,omitzero"`
	// Youtube link of launch.
	YoutubeLaunchLink param.Opt[string] `json:"youtubeLaunchLink,omitzero"`
	paramObj
}

func (r SeradataSpacecraftDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataSpacecraftDetailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeradataSpacecraftDetailUpdateParams) UnmarshalJSON(data []byte) error {
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
type SeradataSpacecraftDetailUpdateParamsDataMode string

const (
	SeradataSpacecraftDetailUpdateParamsDataModeReal      SeradataSpacecraftDetailUpdateParamsDataMode = "REAL"
	SeradataSpacecraftDetailUpdateParamsDataModeTest      SeradataSpacecraftDetailUpdateParamsDataMode = "TEST"
	SeradataSpacecraftDetailUpdateParamsDataModeSimulated SeradataSpacecraftDetailUpdateParamsDataMode = "SIMULATED"
	SeradataSpacecraftDetailUpdateParamsDataModeExercise  SeradataSpacecraftDetailUpdateParamsDataMode = "EXERCISE"
)

type SeradataSpacecraftDetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataSpacecraftDetailListParams]'s query parameters as
// `url.Values`.
func (r SeradataSpacecraftDetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataSpacecraftDetailCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataSpacecraftDetailCountParams]'s query parameters as
// `url.Values`.
func (r SeradataSpacecraftDetailCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataSpacecraftDetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataSpacecraftDetailGetParams]'s query parameters as
// `url.Values`.
func (r SeradataSpacecraftDetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataSpacecraftDetailTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataSpacecraftDetailTupleParams]'s query parameters as
// `url.Values`.
func (r SeradataSpacecraftDetailTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
