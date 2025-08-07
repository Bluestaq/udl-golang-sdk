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

// OnorbitService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitService] method instead.
type OnorbitService struct {
	Options        []option.RequestOption
	AntennaDetails OnorbitAntennaDetailService
}

// NewOnorbitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOnorbitService(opts ...option.RequestOption) (r OnorbitService) {
	r = OnorbitService{}
	r.Options = opts
	r.AntennaDetails = NewOnorbitAntennaDetailService(opts...)
	return
}

// Service operation to take a single onorbit object as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *OnorbitService) New(ctx context.Context, body OnorbitNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single OnOrbit object. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *OnorbitService) Update(ctx context.Context, id string, body OnorbitUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbit/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitService) List(ctx context.Context, query OnorbitListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbit"
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
func (r *OnorbitService) ListAutoPaging(ctx context.Context, query OnorbitListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an OnOrbit object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *OnorbitService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbit/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnorbitService) Count(ctx context.Context, query OnorbitCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onorbit/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single OnOrbit object by its unique ID passed as a
// path parameter.
func (r *OnorbitService) Get(ctx context.Context, id string, query OnorbitGetParams, opts ...option.RequestOption) (res *shared.OnorbitFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbit/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This service queries common information across Radar, EO, and RF observations.
func (r *OnorbitService) GetSignature(ctx context.Context, query OnorbitGetSignatureParams, opts ...option.RequestOption) (res *OnorbitGetSignatureResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/onorbit/getSignature"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *OnorbitService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *OnorbitQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/onorbit/queryhelp"
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
func (r *OnorbitService) Tuple(ctx context.Context, query OnorbitTupleParams, opts ...option.RequestOption) (res *[]shared.OnorbitFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/onorbit/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model object representing on-orbit objects or satellites in the system.
//
// The properties ClassificationMarking, DataMode, SatNo, Source are required.
type OnorbitIngestParam struct {
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
	DataMode OnorbitIngestDataMode `json:"dataMode,omitzero,required"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Alternate name of the on-orbit object.
	AltName param.Opt[string] `json:"altName,omitzero"`
	// Common name of the on-orbit object.
	CommonName param.Opt[string] `json:"commonName,omitzero"`
	// Constellation to which this satellite belongs.
	Constellation param.Opt[string] `json:"constellation,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Date of decay.
	DecayDate param.Opt[time.Time] `json:"decayDate,omitzero" format:"date-time"`
	// For the public catalog, the idOnOrbit is typically the satellite number as a
	// string, but may be a UUID for analyst or other unknown or untracked satellites,
	// auto-generated by the system.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// International Designator, typically of the format YYYYLLLAAA, where YYYY is the
	// launch year, LLL is the sequential launch number of that year, and AAA is an
	// optional launch piece designator for the launch.
	IntlDes param.Opt[string] `json:"intlDes,omitzero"`
	// Date of launch.
	LaunchDate param.Opt[time.Time] `json:"launchDate,omitzero" format:"date"`
	// Id of the associated launchSite entity.
	LaunchSiteID param.Opt[string] `json:"launchSiteId,omitzero"`
	// Estimated lifetime of the on-orbit payload, if known.
	LifetimeYears param.Opt[int64] `json:"lifetimeYears,omitzero"`
	// Mission number of the on-orbit object.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Category of the on-orbit object. (Unknown, On-Orbit, Decayed, Cataloged Without
	// State, Launch Nominal, Analyst Satellite, Cislunar, Lunar, Hyperbolic,
	// Heliocentric, Interplanetary, Lagrangian, Docked).
	//
	// Any of "Unknown", "On-Orbit", "Decayed", "Cataloged Without State", "Launch
	// Nominal", "Analyst Satellite", "Cislunar", "Lunar", "Hyperbolic",
	// "Heliocentric", "Interplanetary", "Lagrangian", "Docked".
	Category OnorbitIngestCategory `json:"category,omitzero"`
	// Type of on-orbit object: ROCKET BODY, DEBRIS, PAYLOAD, PLATFORM, MANNED,
	// UNKNOWN.
	//
	// Any of "ROCKET BODY", "DEBRIS", "PAYLOAD", "PLATFORM", "MANNED", "UNKNOWN".
	ObjectType OnorbitIngestObjectType `json:"objectType,omitzero"`
	paramObj
}

func (r OnorbitIngestParam) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitIngestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitIngestParam) UnmarshalJSON(data []byte) error {
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
type OnorbitIngestDataMode string

const (
	OnorbitIngestDataModeReal      OnorbitIngestDataMode = "REAL"
	OnorbitIngestDataModeTest      OnorbitIngestDataMode = "TEST"
	OnorbitIngestDataModeSimulated OnorbitIngestDataMode = "SIMULATED"
	OnorbitIngestDataModeExercise  OnorbitIngestDataMode = "EXERCISE"
)

// Category of the on-orbit object. (Unknown, On-Orbit, Decayed, Cataloged Without
// State, Launch Nominal, Analyst Satellite, Cislunar, Lunar, Hyperbolic,
// Heliocentric, Interplanetary, Lagrangian, Docked).
type OnorbitIngestCategory string

const (
	OnorbitIngestCategoryUnknown               OnorbitIngestCategory = "Unknown"
	OnorbitIngestCategoryOnOrbit               OnorbitIngestCategory = "On-Orbit"
	OnorbitIngestCategoryDecayed               OnorbitIngestCategory = "Decayed"
	OnorbitIngestCategoryCatalogedWithoutState OnorbitIngestCategory = "Cataloged Without State"
	OnorbitIngestCategoryLaunchNominal         OnorbitIngestCategory = "Launch Nominal"
	OnorbitIngestCategoryAnalystSatellite      OnorbitIngestCategory = "Analyst Satellite"
	OnorbitIngestCategoryCislunar              OnorbitIngestCategory = "Cislunar"
	OnorbitIngestCategoryLunar                 OnorbitIngestCategory = "Lunar"
	OnorbitIngestCategoryHyperbolic            OnorbitIngestCategory = "Hyperbolic"
	OnorbitIngestCategoryHeliocentric          OnorbitIngestCategory = "Heliocentric"
	OnorbitIngestCategoryInterplanetary        OnorbitIngestCategory = "Interplanetary"
	OnorbitIngestCategoryLagrangian            OnorbitIngestCategory = "Lagrangian"
	OnorbitIngestCategoryDocked                OnorbitIngestCategory = "Docked"
)

// Type of on-orbit object: ROCKET BODY, DEBRIS, PAYLOAD, PLATFORM, MANNED,
// UNKNOWN.
type OnorbitIngestObjectType string

const (
	OnorbitIngestObjectTypeRocketBody OnorbitIngestObjectType = "ROCKET BODY"
	OnorbitIngestObjectTypeDebris     OnorbitIngestObjectType = "DEBRIS"
	OnorbitIngestObjectTypePayload    OnorbitIngestObjectType = "PAYLOAD"
	OnorbitIngestObjectTypePlatform   OnorbitIngestObjectType = "PLATFORM"
	OnorbitIngestObjectTypeManned     OnorbitIngestObjectType = "MANNED"
	OnorbitIngestObjectTypeUnknown    OnorbitIngestObjectType = "UNKNOWN"
)

// Model object representing on-orbit objects or satellites in the system.
type OnorbitListResponse struct {
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
	DataMode OnorbitListResponseDataMode `json:"dataMode,required"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Alternate name of the on-orbit object.
	AltName string `json:"altName"`
	// Category of the on-orbit object. (Unknown, On-Orbit, Decayed, Cataloged Without
	// State, Launch Nominal, Analyst Satellite, Cislunar, Lunar, Hyperbolic,
	// Heliocentric, Interplanetary, Lagrangian, Docked).
	//
	// Any of "Unknown", "On-Orbit", "Decayed", "Cataloged Without State", "Launch
	// Nominal", "Analyst Satellite", "Cislunar", "Lunar", "Hyperbolic",
	// "Heliocentric", "Interplanetary", "Lagrangian", "Docked".
	Category OnorbitListResponseCategory `json:"category"`
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
	ObjectType OnorbitListResponseObjectType `json:"objectType"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		SatNo                 respjson.Field
		Source                respjson.Field
		AltName               respjson.Field
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
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitListResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitListResponseDataMode string

const (
	OnorbitListResponseDataModeReal      OnorbitListResponseDataMode = "REAL"
	OnorbitListResponseDataModeTest      OnorbitListResponseDataMode = "TEST"
	OnorbitListResponseDataModeSimulated OnorbitListResponseDataMode = "SIMULATED"
	OnorbitListResponseDataModeExercise  OnorbitListResponseDataMode = "EXERCISE"
)

// Category of the on-orbit object. (Unknown, On-Orbit, Decayed, Cataloged Without
// State, Launch Nominal, Analyst Satellite, Cislunar, Lunar, Hyperbolic,
// Heliocentric, Interplanetary, Lagrangian, Docked).
type OnorbitListResponseCategory string

const (
	OnorbitListResponseCategoryUnknown               OnorbitListResponseCategory = "Unknown"
	OnorbitListResponseCategoryOnOrbit               OnorbitListResponseCategory = "On-Orbit"
	OnorbitListResponseCategoryDecayed               OnorbitListResponseCategory = "Decayed"
	OnorbitListResponseCategoryCatalogedWithoutState OnorbitListResponseCategory = "Cataloged Without State"
	OnorbitListResponseCategoryLaunchNominal         OnorbitListResponseCategory = "Launch Nominal"
	OnorbitListResponseCategoryAnalystSatellite      OnorbitListResponseCategory = "Analyst Satellite"
	OnorbitListResponseCategoryCislunar              OnorbitListResponseCategory = "Cislunar"
	OnorbitListResponseCategoryLunar                 OnorbitListResponseCategory = "Lunar"
	OnorbitListResponseCategoryHyperbolic            OnorbitListResponseCategory = "Hyperbolic"
	OnorbitListResponseCategoryHeliocentric          OnorbitListResponseCategory = "Heliocentric"
	OnorbitListResponseCategoryInterplanetary        OnorbitListResponseCategory = "Interplanetary"
	OnorbitListResponseCategoryLagrangian            OnorbitListResponseCategory = "Lagrangian"
	OnorbitListResponseCategoryDocked                OnorbitListResponseCategory = "Docked"
)

// Type of on-orbit object: ROCKET BODY, DEBRIS, PAYLOAD, PLATFORM, MANNED,
// UNKNOWN.
type OnorbitListResponseObjectType string

const (
	OnorbitListResponseObjectTypeRocketBody OnorbitListResponseObjectType = "ROCKET BODY"
	OnorbitListResponseObjectTypeDebris     OnorbitListResponseObjectType = "DEBRIS"
	OnorbitListResponseObjectTypePayload    OnorbitListResponseObjectType = "PAYLOAD"
	OnorbitListResponseObjectTypePlatform   OnorbitListResponseObjectType = "PLATFORM"
	OnorbitListResponseObjectTypeManned     OnorbitListResponseObjectType = "MANNED"
	OnorbitListResponseObjectTypeUnknown    OnorbitListResponseObjectType = "UNKNOWN"
)

// Contains a list of common information across both Radar and EO observations.
type OnorbitGetSignatureResponse struct {
	// Model representation of observation data for electro-optical based sensor
	// phenomenologies. ECI J2K is the preferred reference frame for EOObservations,
	// however, several user-specified reference frames are accommodated. Users should
	// check the EOObservation record as well as the 'Discover' tab in the storefront
	// to confirm the coordinate frames used by the data provider.
	EoObservation EoObservationAbridged `json:"eoObservation"`
	// Model representation of observation data for radar based sensor phenomenologies.
	// J2000 is the preferred coordinate frame for all observations, but in some cases
	// observations may be in another frame depending on the provider. Please see the
	// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
	RadarObservation OnorbitGetSignatureResponseRadarObservation `json:"radarObservation"`
	// Model representation of observation data for active/passive radio frequency (RF)
	// based sensor phenomenologies. J2000 is the preferred coordinate frame for all
	// observations, but in some cases observations may be in another frame depending
	// on the provider. Please see the 'Discover' tab in the storefront to confirm
	// coordinate frames by data provider. RF observations include several optional
	// ordered arrays which are used to provide detailed information on recorded
	// signals such as power spectral density lists or active signals (code taps/fills,
	// etc). In these cases, the sizes of the arrays must match and can be assumed to
	// have consistent indexing across arrays (e.g. powers[0] is the measured power at
	// frequencies[0]).
	RfObservation OnorbitGetSignatureResponseRfObservation `json:"rfObservation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EoObservation    respjson.Field
		RadarObservation respjson.Field
		RfObservation    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitGetSignatureResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitGetSignatureResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of observation data for radar based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
type OnorbitGetSignatureResponseRadarObservation struct {
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
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// azimuth angle in degrees and topocentric frame.
	Azimuth float64 `json:"azimuth"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias float64 `json:"azimuthBias"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured bool `json:"azimuthMeasured"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// ID of the beam that produced this observation.
	Beam float64 `json:"beam"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination float64 `json:"declination"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured bool `json:"declinationMeasured"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Corrected doppler measurement in meters per second.
	Doppler float64 `json:"doppler"`
	// One sigma uncertainty in the corrected doppler measurement, in meters/second.
	DopplerUnc float64 `json:"dopplerUnc"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation float64 `json:"elevation"`
	// Sensor elevation bias in degrees.
	ElevationBias float64 `json:"elevationBias"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured bool `json:"elevationMeasured"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate float64 `json:"elevationRate"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc float64 `json:"elevationUnc"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The position of this observation within a track (FENCE, FIRST, IN, LAST,
	// SINGLE). This identifier is optional and, if null, no assumption should be made
	// regarding whether other observations may or may not exist to compose a track.
	ObPosition string `json:"obPosition"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Radar cross section in meters squared for orthogonal polarization.
	OrthogonalRcs float64 `json:"orthogonalRcs"`
	// one sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc float64 `json:"orthogonalRcsUnc"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra float64 `json:"ra"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured bool `json:"raMeasured"`
	// Target range in km.
	Range float64 `json:"range"`
	// Range accelaration in km/s2.
	RangeAccel float64 `json:"rangeAccel"`
	// One sigma uncertainty in the range acceleration measurement, in
	// kilometers/(second^2).
	RangeAccelUnc float64 `json:"rangeAccelUnc"`
	// Sensor range bias in km.
	RangeBias float64 `json:"rangeBias"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured bool `json:"rangeMeasured"`
	// Rate of change of the line of sight range in km/sec.
	RangeRate float64 `json:"rangeRate"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured bool `json:"rangeRateMeasured"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc float64 `json:"rangeRateUnc"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc float64 `json:"rangeUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Radar cross section in meters squared for polarization principal.
	Rcs float64 `json:"rcs"`
	// one sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc float64 `json:"rcsUnc"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame string `json:"senReferenceFrame"`
	// Sensor x position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senx float64 `json:"senx"`
	// Sensor y position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Seny float64 `json:"seny"`
	// Sensor z position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senz float64 `json:"senz"`
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID string `json:"taskId"`
	// Sensor timing bias in seconds.
	TimingBias float64 `json:"timingBias"`
	// Optional identifier of the track to which this observation belongs.
	TrackID string `json:"trackId"`
	// The beam type (or tracking state) in use at the time of collection of this
	// observation. Values include:
	//
	// INIT ACQ WITH INIT VALUES: Initial acquisition based on predefined initial
	// values such as position, velocity, or other specific parameters.
	//
	// INIT ACQ: Initial acquisition when no prior information or initial values such
	// as position or velocity are available.
	//
	// TRACKING SINGLE BEAM: Continuously tracks and monitors a single target using one
	// specific radar beam.
	//
	// TRACKING SEQUENTIAL ROVING: Sequentially tracks different targets or areas by
	// "roving" from one sector to the next in a systematic order.
	//
	// SELF ACQ WITH INIT VALUES: Autonomously acquires targets using predefined
	// starting parameters or initial values.
	//
	// SELF ACQ: Automatically detects and locks onto targets without the need for
	// predefined initial settings.
	//
	// NON-TRACKING: Non-tracking.
	TrackingState string `json:"trackingState"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Read only enumeration specifying the type of observation (e.g. OPTICAL, RADAR,
	// RF, etc).
	Type string `json:"type"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct bool `json:"uct"`
	// X position of target in km in J2000 coordinate frame.
	X float64 `json:"x"`
	// X velocity of target in km/sec in J2000 coordinate frame.
	Xvel float64 `json:"xvel"`
	// Y position of target in km in J2000 coordinate frame.
	Y float64 `json:"y"`
	// Y velocity of target in km/sec in J2000 coordinate frame.
	Yvel float64 `json:"yvel"`
	// Z position of target in km in J2000 coordinate frame.
	Z float64 `json:"z"`
	// Z velocity of target in km/sec in J2000 coordinate frame.
	Zvel float64 `json:"zvel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Azimuth               respjson.Field
		AzimuthBias           respjson.Field
		AzimuthMeasured       respjson.Field
		AzimuthRate           respjson.Field
		AzimuthUnc            respjson.Field
		Beam                  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Declination           respjson.Field
		DeclinationMeasured   respjson.Field
		Descriptor            respjson.Field
		Doppler               respjson.Field
		DopplerUnc            respjson.Field
		Elevation             respjson.Field
		ElevationBias         respjson.Field
		ElevationMeasured     respjson.Field
		ElevationRate         respjson.Field
		ElevationUnc          respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		ObPosition            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		OrthogonalRcs         respjson.Field
		OrthogonalRcsUnc      respjson.Field
		Ra                    respjson.Field
		RaMeasured            respjson.Field
		Range                 respjson.Field
		RangeAccel            respjson.Field
		RangeAccelUnc         respjson.Field
		RangeBias             respjson.Field
		RangeMeasured         respjson.Field
		RangeRate             respjson.Field
		RangeRateMeasured     respjson.Field
		RangeRateUnc          respjson.Field
		RangeUnc              respjson.Field
		RawFileUri            respjson.Field
		Rcs                   respjson.Field
		RcsUnc                respjson.Field
		SatNo                 respjson.Field
		SenReferenceFrame     respjson.Field
		Senx                  respjson.Field
		Seny                  respjson.Field
		Senz                  respjson.Field
		Snr                   respjson.Field
		SourceDl              respjson.Field
		TaskID                respjson.Field
		TimingBias            respjson.Field
		TrackID               respjson.Field
		TrackingState         respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		Uct                   respjson.Field
		X                     respjson.Field
		Xvel                  respjson.Field
		Y                     respjson.Field
		Yvel                  respjson.Field
		Z                     respjson.Field
		Zvel                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitGetSignatureResponseRadarObservation) RawJSON() string { return r.JSON.raw }
func (r *OnorbitGetSignatureResponseRadarObservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of observation data for active/passive radio frequency (RF)
// based sensor phenomenologies. J2000 is the preferred coordinate frame for all
// observations, but in some cases observations may be in another frame depending
// on the provider. Please see the 'Discover' tab in the storefront to confirm
// coordinate frames by data provider. RF observations include several optional
// ordered arrays which are used to provide detailed information on recorded
// signals such as power spectral density lists or active signals (code taps/fills,
// etc). In these cases, the sizes of the arrays must match and can be assumed to
// have consistent indexing across arrays (e.g. powers[0] is the measured power at
// frequencies[0]).
type OnorbitGetSignatureResponseRfObservation struct {
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
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of RF ob (e.g. RF, RF-SOSI, PSD, RFI, SPOOF, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Antenna name of the RFObservation record.
	AntennaName string `json:"antennaName"`
	// azimuth angle in degrees and J2000 coordinate frame.
	Azimuth float64 `json:"azimuth"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured bool `json:"azimuthMeasured"`
	// Rate of change of the azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the azimuth angle measurement, in degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// Measured bandwidth in Hz.
	Bandwidth float64 `json:"bandwidth"`
	// Baud rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per second.
	BaudRate float64 `json:"baudRate"`
	// The ratio of bit errors per number of received bits.
	BitErrorRate float64 `json:"bitErrorRate"`
	// Carrier standard (e.g. DVB-S2, 802.11g, etc.).
	CarrierStandard string `json:"carrierStandard"`
	// Channel of the RFObservation record.
	Channel int64 `json:"channel"`
	// Collection mode (e.g. SURVEY, SPOT_SEARCH, NEIGHBORHOOD_WATCH, DIRECTED_SEARCH,
	// MANUAL, etc).
	CollectionMode string `json:"collectionMode"`
	// Confidence in the signal and its measurements and characterization.
	Confidence float64 `json:"confidence"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Detection status (e.g. DETECTED, CARRIER_ACQUIRING, CARRIER_DETECTED,
	// NOT_DETECTED, etc).
	DetectionStatus string `json:"detectionStatus"`
	// Measured Equivalent Isotopically Radiated Power in dBW.
	Eirp float64 `json:"eirp"`
	// elevation in degrees and J2000 coordinate frame.
	Elevation float64 `json:"elevation"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured bool `json:"elevationMeasured"`
	// Rate of change of the elevation in degrees per second.
	ElevationRate float64 `json:"elevationRate"`
	// One sigma uncertainty in the elevation angle measurement, in degrees.
	ElevationUnc float64 `json:"elevationUnc"`
	// ELINT notation.
	Elnot string `json:"elnot"`
	// End carrier frequency in Hz.
	EndFrequency float64 `json:"endFrequency"`
	// Center carrier frequency in Hz.
	Frequency float64 `json:"frequency"`
	// Frequency Shift of the RFObservation record.
	FrequencyShift float64 `json:"frequencyShift"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// True if the signal is incoming, false if outgoing.
	Incoming bool `json:"incoming"`
	// Inner forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	InnerCodingRate int64 `json:"innerCodingRate"`
	// Maximum measured PSD value of the trace in dBW.
	MaxPsd float64 `json:"maxPSD"`
	// Minimum measured PSD value of the trace in dBW.
	MinPsd float64 `json:"minPSD"`
	// Transponder modulation (e.g. Auto, QPSK, 8PSK, etc).
	Modulation string `json:"modulation"`
	// Noise power density, in dBW-Hz.
	NoisePwrDensity float64 `json:"noisePwrDensity"`
	// Expected bandwidth in Hz.
	NominalBandwidth float64 `json:"nominalBandwidth"`
	// Expected Equivalent Isotopically Radiated Power in dBW.
	NominalEirp float64 `json:"nominalEirp"`
	// Nominal or expected center carrier frequency in Hz.
	NominalFrequency float64 `json:"nominalFrequency"`
	// Expected carrier power over noise (dBW/Hz).
	NominalPowerOverNoise float64 `json:"nominalPowerOverNoise"`
	// Nominal or expected signal to noise ratio, in dB.
	NominalSnr float64 `json:"nominalSnr"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Outer forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	OuterCodingRate int64 `json:"outerCodingRate"`
	// Peak of the RFObservation record.
	Peak bool `json:"peak"`
	// A pulse group repetition interval (PGRI) is a pulse train in which there are
	// groups of closely spaced pulses separated by much longer times between these
	// pulse groups.
	Pgri float64 `json:"pgri"`
	// The antenna pointing dependent polarizer angle, in degrees.
	Polarity float64 `json:"polarity"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	PolarityType string `json:"polarityType"`
	// Measured carrier power over noise (dBW/Hz).
	PowerOverNoise float64 `json:"powerOverNoise"`
	// Target range in km.
	Range float64 `json:"range"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured bool `json:"rangeMeasured"`
	// Rate of change of the range in km/sec.
	RangeRate float64 `json:"rangeRate"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured bool `json:"rangeRateMeasured"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc float64 `json:"rangeRateUnc"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc float64 `json:"rangeUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Reference signal level, in dBW.
	ReferenceLevel float64 `json:"referenceLevel"`
	// Measured power of the center carrier frequency in dBW.
	RelativeCarrierPower float64 `json:"relativeCarrierPower"`
	// The measure of the signal created from the sum of all the noise sources and
	// unwanted signals within the measurement system, in dBW.
	RelativeNoiseFloor float64 `json:"relativeNoiseFloor"`
	// Resolution bandwidth in Hz.
	ResolutionBandwidth float64 `json:"resolutionBandwidth"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sensor altitude at obTime (if mobile/onorbit) in km. If null, can be obtained
	// from sensor info.
	Senalt float64 `json:"senalt"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon float64 `json:"senlon"`
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Measured spectrum analyzer power of the center carrier frequency in dBW.
	SpectrumAnalyzerPower float64 `json:"spectrumAnalyzerPower"`
	// Start carrier frequency in Hz.
	StartFrequency float64 `json:"startFrequency"`
	// Switch Point of the RFObservation record.
	SwitchPoint int64 `json:"switchPoint"`
	// Symbol to noise ratio, in dB.
	SymbolToNoiseRatio float64 `json:"symbolToNoiseRatio"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID string `json:"taskId"`
	// Optional identifier of the track to which this observation belongs.
	TrackID string `json:"trackId"`
	// Target track or apparent range in km.
	TrackRange float64 `json:"trackRange"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Transmit pulse shaping filter roll-off value.
	TransmitFilterRollOff float64 `json:"transmitFilterRollOff"`
	// Transmit pulse shaping filter typ (e.g. RRC).
	TransmitFilterType string `json:"transmitFilterType"`
	// Optional identifier provided by observation source to indicate the transponder
	// used for this measurement.
	Transponder string `json:"transponder"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct bool `json:"uct"`
	// Optional URL containing additional information on this observation.
	URL string `json:"url"`
	// Video bandwidth in Hz.
	VideoBandwidth float64 `json:"videoBandwidth"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AntennaName           respjson.Field
		Azimuth               respjson.Field
		AzimuthMeasured       respjson.Field
		AzimuthRate           respjson.Field
		AzimuthUnc            respjson.Field
		Bandwidth             respjson.Field
		BaudRate              respjson.Field
		BitErrorRate          respjson.Field
		CarrierStandard       respjson.Field
		Channel               respjson.Field
		CollectionMode        respjson.Field
		Confidence            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		DetectionStatus       respjson.Field
		Eirp                  respjson.Field
		Elevation             respjson.Field
		ElevationMeasured     respjson.Field
		ElevationRate         respjson.Field
		ElevationUnc          respjson.Field
		Elnot                 respjson.Field
		EndFrequency          respjson.Field
		Frequency             respjson.Field
		FrequencyShift        respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		Incoming              respjson.Field
		InnerCodingRate       respjson.Field
		MaxPsd                respjson.Field
		MinPsd                respjson.Field
		Modulation            respjson.Field
		NoisePwrDensity       respjson.Field
		NominalBandwidth      respjson.Field
		NominalEirp           respjson.Field
		NominalFrequency      respjson.Field
		NominalPowerOverNoise respjson.Field
		NominalSnr            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		OuterCodingRate       respjson.Field
		Peak                  respjson.Field
		Pgri                  respjson.Field
		Polarity              respjson.Field
		PolarityType          respjson.Field
		PowerOverNoise        respjson.Field
		Range                 respjson.Field
		RangeMeasured         respjson.Field
		RangeRate             respjson.Field
		RangeRateMeasured     respjson.Field
		RangeRateUnc          respjson.Field
		RangeUnc              respjson.Field
		RawFileUri            respjson.Field
		ReferenceLevel        respjson.Field
		RelativeCarrierPower  respjson.Field
		RelativeNoiseFloor    respjson.Field
		ResolutionBandwidth   respjson.Field
		SatNo                 respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		Snr                   respjson.Field
		SourceDl              respjson.Field
		SpectrumAnalyzerPower respjson.Field
		StartFrequency        respjson.Field
		SwitchPoint           respjson.Field
		SymbolToNoiseRatio    respjson.Field
		TaskID                respjson.Field
		TrackID               respjson.Field
		TrackRange            respjson.Field
		TransactionID         respjson.Field
		TransmitFilterRollOff respjson.Field
		TransmitFilterType    respjson.Field
		Transponder           respjson.Field
		Uct                   respjson.Field
		URL                   respjson.Field
		VideoBandwidth        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitGetSignatureResponseRfObservation) RawJSON() string { return r.JSON.raw }
func (r *OnorbitGetSignatureResponseRfObservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbitQueryhelpResponse struct {
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
func (r OnorbitQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbitNewParams struct {
	// Model object representing on-orbit objects or satellites in the system.
	OnorbitIngest OnorbitIngestParam
	paramObj
}

func (r OnorbitNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.OnorbitIngest)
}
func (r *OnorbitNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.OnorbitIngest)
}

type OnorbitUpdateParams struct {
	// Model object representing on-orbit objects or satellites in the system.
	OnorbitIngest OnorbitIngestParam
	paramObj
}

func (r OnorbitUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.OnorbitIngest)
}
func (r *OnorbitUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.OnorbitIngest)
}

type OnorbitListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitListParams]'s query parameters as `url.Values`.
func (r OnorbitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitCountParams]'s query parameters as `url.Values`.
func (r OnorbitCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitGetParams]'s query parameters as `url.Values`.
func (r OnorbitGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitGetSignatureParams struct {
	// ID of the Onorbit object.
	IDOnOrbit   string           `query:"idOnOrbit,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitGetSignatureParams]'s query parameters as
// `url.Values`.
func (r OnorbitGetSignatureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitTupleParams]'s query parameters as `url.Values`.
func (r OnorbitTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
