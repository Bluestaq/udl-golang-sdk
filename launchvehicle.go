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

// LaunchVehicleService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaunchVehicleService] method instead.
type LaunchVehicleService struct {
	Options []option.RequestOption
}

// NewLaunchVehicleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaunchVehicleService(opts ...option.RequestOption) (r LaunchVehicleService) {
	r = LaunchVehicleService{}
	r.Options = opts
	return
}

// Service operation to take a single LaunchVehicle as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LaunchVehicleService) New(ctx context.Context, body LaunchVehicleNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchvehicle"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single LaunchVehicle. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *LaunchVehicleService) Update(ctx context.Context, id string, body LaunchVehicleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchvehicle/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaunchVehicleService) List(ctx context.Context, query LaunchVehicleListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaunchVehicleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/launchvehicle"
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
func (r *LaunchVehicleService) ListAutoPaging(ctx context.Context, query LaunchVehicleListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaunchVehicleListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a LaunchVehicle object specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LaunchVehicleService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchvehicle/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LaunchVehicleService) Count(ctx context.Context, query LaunchVehicleCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/launchvehicle/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single LaunchVehicle record by its unique ID passed
// as a path parameter.
func (r *LaunchVehicleService) Get(ctx context.Context, id string, query LaunchVehicleGetParams, opts ...option.RequestOption) (res *LaunchVehicleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchvehicle/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *LaunchVehicleService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *LaunchVehicleQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/launchvehicle/queryhelp"
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
func (r *LaunchVehicleService) Tuple(ctx context.Context, query LaunchVehicleTupleParams, opts ...option.RequestOption) (res *[]LaunchVehicleTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/launchvehicle/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of basic information about known launch vehicles. A launch
// vehicle may have several details records including characteristics and
// information compiled by a particular source.
type LaunchVehicleListResponse struct {
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
	DataMode LaunchVehicleListResponseDataMode `json:"dataMode,required"`
	// Launch vehicle name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Vehicle type.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchVehicleListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleListResponse) UnmarshalJSON(data []byte) error {
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
type LaunchVehicleListResponseDataMode string

const (
	LaunchVehicleListResponseDataModeReal      LaunchVehicleListResponseDataMode = "REAL"
	LaunchVehicleListResponseDataModeTest      LaunchVehicleListResponseDataMode = "TEST"
	LaunchVehicleListResponseDataModeSimulated LaunchVehicleListResponseDataMode = "SIMULATED"
	LaunchVehicleListResponseDataModeExercise  LaunchVehicleListResponseDataMode = "EXERCISE"
)

// Model representation of basic information about known launch vehicles. A launch
// vehicle may have several details records including characteristics and
// information compiled by a particular source.
type LaunchVehicleGetResponse struct {
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
	DataMode LaunchVehicleGetResponseDataMode `json:"dataMode,required"`
	// Launch vehicle name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Read-only collection of additional LaunchVehicleDetails by various sources for
	// this launch vehicle, ignored on create/update. These details must be created
	// separately via the /udl/launchvehicledetails operations.
	LaunchVehicleDetails []LaunchVehicleGetResponseLaunchVehicleDetail `json:"launchVehicleDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only collection of stages for this launch vehicle, ignored on
	// create/update.
	Stages []LaunchVehicleGetResponseStage `json:"stages"`
	// Vehicle type.
	Type string `json:"type"`
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
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		LaunchVehicleDetails  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Stages                respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchVehicleGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleGetResponse) UnmarshalJSON(data []byte) error {
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
type LaunchVehicleGetResponseDataMode string

const (
	LaunchVehicleGetResponseDataModeReal      LaunchVehicleGetResponseDataMode = "REAL"
	LaunchVehicleGetResponseDataModeTest      LaunchVehicleGetResponseDataMode = "TEST"
	LaunchVehicleGetResponseDataModeSimulated LaunchVehicleGetResponseDataMode = "SIMULATED"
	LaunchVehicleGetResponseDataModeExercise  LaunchVehicleGetResponseDataMode = "EXERCISE"
)

// Model representation of launch vehicle details and characteristics, compiled by
// a particular source. A vehicle may have multiple details records from various
// sources.
type LaunchVehicleGetResponseLaunchVehicleDetail struct {
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
	// Identifier of the parent launch vehicle record.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Launch vehicle attitude accuracy (degrees).
	AttitudeAccuracy float64 `json:"attitudeAccuracy"`
	// Vehicle category.
	Category string `json:"category"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Launch vehicle deployment rotation rate in RPM.
	DeploymentRotationRate float64 `json:"deploymentRotationRate"`
	// Vehicle diameter in meters.
	Diameter float64 `json:"diameter"`
	// Launch vehicle estimated launch price in US dollars.
	EstLaunchPrice float64 `json:"estLaunchPrice"`
	// Launch vehicle typical estimated launch price in US dollars.
	EstLaunchPriceTypical float64 `json:"estLaunchPriceTypical"`
	// Vehicle fairing maximum external diameter in meters.
	FairingExternalDiameter float64 `json:"fairingExternalDiameter"`
	// Vehicle fairing maximum internal diameter in meters.
	FairingInternalDiameter float64 `json:"fairingInternalDiameter"`
	// Vehicle fairing length in meters.
	FairingLength float64 `json:"fairingLength"`
	// Vehicle fairing mass in kg.
	FairingMass float64 `json:"fairingMass"`
	// Fairing material.
	FairingMaterial string `json:"fairingMaterial"`
	// Name of the fairing.
	FairingName string `json:"fairingName"`
	// Notes/Description of the launch vehicle fairing.
	FairingNotes string `json:"fairingNotes"`
	// Vehicle family.
	Family string `json:"family"`
	// Maximum vehicle payload mass to GEO orbit in kg.
	GeoPayloadMass float64 `json:"geoPayloadMass"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Apogee Margin (degrees).
	GtoInj3SigAccuracyApogeeMargin float64 `json:"gtoInj3SigAccuracyApogeeMargin"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Apogee Target (degrees).
	GtoInj3SigAccuracyApogeeTarget float64 `json:"gtoInj3SigAccuracyApogeeTarget"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Inclination Margin (degrees).
	GtoInj3SigAccuracyInclinationMargin float64 `json:"gtoInj3SigAccuracyInclinationMargin"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Inclination Target (degrees).
	GtoInj3SigAccuracyInclinationTarget float64 `json:"gtoInj3SigAccuracyInclinationTarget"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Perigee Margin (degrees).
	GtoInj3SigAccuracyPerigeeMargin float64 `json:"gtoInj3SigAccuracyPerigeeMargin"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Perigee Target (degrees).
	GtoInj3SigAccuracyPerigeeTarget float64 `json:"gtoInj3SigAccuracyPerigeeTarget"`
	// Max vehicle payload mass to Geo-Transfer Orbit in kg.
	GtoPayloadMass float64 `json:"gtoPayloadMass"`
	// Vehicle total mass at launch time in kg (including all boosters).
	LaunchMass float64 `json:"launchMass"`
	// Vehicle launch prefix.
	LaunchPrefix string `json:"launchPrefix"`
	// Vehicle length in meters.
	Length float64 `json:"length"`
	// Max vehicle payload mass to LEO orbit in kg.
	LeoPayloadMass float64 `json:"leoPayloadMass"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	ManufacturerOrg shared.OrganizationFull `json:"manufacturerOrg"`
	// ID of the organization that manufactures the launch vehicle.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Vehicle maximum acceleration load in g.
	MaxAccelLoad float64 `json:"maxAccelLoad"`
	// Vehicle maximum acoustic level in dB.
	MaxAcousticLevel float64 `json:"maxAcousticLevel"`
	// Vehicle maximum acoustic level range in Hz.
	MaxAcousticLevelRange float64 `json:"maxAcousticLevelRange"`
	// Vehicle fairing maximum pressure change in kPa/sec.
	MaxFairingPressureChange float64 `json:"maxFairingPressureChange"`
	// Vehicle maximum flight shock force in g.
	MaxFlightShockForce float64 `json:"maxFlightShockForce"`
	// Vehicle maximum flight shock frequency in Hz.
	MaxFlightShockFreq float64 `json:"maxFlightShockFreq"`
	// Vehicle maximum payload lateral frequency in Hz.
	MaxPayloadFreqLat float64 `json:"maxPayloadFreqLat"`
	// Vehicle maximum payload longitudinal frequency in Hz.
	MaxPayloadFreqLon float64 `json:"maxPayloadFreqLon"`
	// Vehicle minor variant.
	MinorVariant string `json:"minorVariant"`
	// Notes/Description of the launch vehicle.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Oxidizer type.
	Oxidizer string `json:"oxidizer"`
	// Notes/Description of the launch vehicle payload.
	PayloadNotes string `json:"payloadNotes"`
	// Launch vehicle payload separation rate in m/s.
	PayloadSeparationRate float64 `json:"payloadSeparationRate"`
	// Propellant type.
	Propellant string `json:"propellant"`
	// Vehicle overall sound pressure level in dB.
	SoundPressureLevel float64 `json:"soundPressureLevel"`
	// Optional URL for additional information on the vehicle.
	SourceURL string `json:"sourceURL"`
	// Max vehicle payload mass to Sun-Synchronous Orbit in kg.
	SSOPayloadMass float64 `json:"ssoPayloadMass"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Vehicle variant.
	Variant string `json:"variant"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking               respjson.Field
		DataMode                            respjson.Field
		IDLaunchVehicle                     respjson.Field
		Source                              respjson.Field
		ID                                  respjson.Field
		AttitudeAccuracy                    respjson.Field
		Category                            respjson.Field
		CreatedAt                           respjson.Field
		CreatedBy                           respjson.Field
		DeploymentRotationRate              respjson.Field
		Diameter                            respjson.Field
		EstLaunchPrice                      respjson.Field
		EstLaunchPriceTypical               respjson.Field
		FairingExternalDiameter             respjson.Field
		FairingInternalDiameter             respjson.Field
		FairingLength                       respjson.Field
		FairingMass                         respjson.Field
		FairingMaterial                     respjson.Field
		FairingName                         respjson.Field
		FairingNotes                        respjson.Field
		Family                              respjson.Field
		GeoPayloadMass                      respjson.Field
		GtoInj3SigAccuracyApogeeMargin      respjson.Field
		GtoInj3SigAccuracyApogeeTarget      respjson.Field
		GtoInj3SigAccuracyInclinationMargin respjson.Field
		GtoInj3SigAccuracyInclinationTarget respjson.Field
		GtoInj3SigAccuracyPerigeeMargin     respjson.Field
		GtoInj3SigAccuracyPerigeeTarget     respjson.Field
		GtoPayloadMass                      respjson.Field
		LaunchMass                          respjson.Field
		LaunchPrefix                        respjson.Field
		Length                              respjson.Field
		LeoPayloadMass                      respjson.Field
		ManufacturerOrg                     respjson.Field
		ManufacturerOrgID                   respjson.Field
		MaxAccelLoad                        respjson.Field
		MaxAcousticLevel                    respjson.Field
		MaxAcousticLevelRange               respjson.Field
		MaxFairingPressureChange            respjson.Field
		MaxFlightShockForce                 respjson.Field
		MaxFlightShockFreq                  respjson.Field
		MaxPayloadFreqLat                   respjson.Field
		MaxPayloadFreqLon                   respjson.Field
		MinorVariant                        respjson.Field
		Notes                               respjson.Field
		Origin                              respjson.Field
		OrigNetwork                         respjson.Field
		Oxidizer                            respjson.Field
		PayloadNotes                        respjson.Field
		PayloadSeparationRate               respjson.Field
		Propellant                          respjson.Field
		SoundPressureLevel                  respjson.Field
		SourceURL                           respjson.Field
		SSOPayloadMass                      respjson.Field
		Tags                                respjson.Field
		UpdatedAt                           respjson.Field
		UpdatedBy                           respjson.Field
		Variant                             respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchVehicleGetResponseLaunchVehicleDetail) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleGetResponseLaunchVehicleDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Launch stage information for a particular launch vehicle. A launch vehicle can
// have several stages, each with 1 to many engines.
type LaunchVehicleGetResponseStage struct {
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
	// Identifier of the Engine record for this stage.
	IDEngine string `json:"idEngine,required"`
	// Identifier of the launch vehicle record for this stage.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Description/notes of the stage avionics.
	AvionicsNotes string `json:"avionicsNotes"`
	// Total burn time of the stage engines in seconds.
	BurnTime float64 `json:"burnTime"`
	// Control thruster 1 type.
	ControlThruster1 string `json:"controlThruster1"`
	// Control thruster 2 type.
	ControlThruster2 string `json:"controlThruster2"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Stage maximum external diameter in meters.
	Diameter float64 `json:"diameter"`
	// Known launch vehicle engines and their performance characteristics and limits. A
	// launch vehicle has 1 to many engines per stage.
	Engine shared.Engine `json:"engine"`
	// Stage length in meters.
	Length float64 `json:"length"`
	// Thrust of the stage main engine at sea level in kN.
	MainEngineThrustSeaLevel float64 `json:"mainEngineThrustSeaLevel"`
	// Thrust of the stage main engine in a vacuum in kN.
	MainEngineThrustVacuum float64 `json:"mainEngineThrustVacuum"`
	// ID of the organization that manufactures this launch stage.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Stage gross mass in kg.
	Mass float64 `json:"mass"`
	// Description/notes of the stage.
	Notes string `json:"notes"`
	// Number of burns for the stage engines.
	NumBurns int64 `json:"numBurns"`
	// Number of type control thruster 1.
	NumControlThruster1 int64 `json:"numControlThruster1"`
	// Number of type control thruster 2.
	NumControlThruster2 int64 `json:"numControlThruster2"`
	// The number of the specified engines on this launch stage.
	NumEngines int64 `json:"numEngines"`
	// Number of launch stage elements used in this stage.
	NumStageElements int64 `json:"numStageElements"`
	// Number of vernier or additional engines.
	NumVernier int64 `json:"numVernier"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Array of URLs of photos of the stage.
	PhotoURLs []string `json:"photoURLs"`
	// Boolean indicating if this launch stage can be restarted.
	Restartable bool `json:"restartable"`
	// Boolean indicating if this launch stage is reusable.
	Reusable bool `json:"reusable"`
	// The stage number of this launch stage.
	StageNumber int64 `json:"stageNumber"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Total thrust of the stage at sea level in kN.
	ThrustSeaLevel float64 `json:"thrustSeaLevel"`
	// Total thrust of the stage in a vacuum in kN.
	ThrustVacuum float64 `json:"thrustVacuum"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Engine vernier or additional engine type.
	Vernier string `json:"vernier"`
	// Total burn time of the vernier or additional stage engines in seconds.
	VernierBurnTime float64 `json:"vernierBurnTime"`
	// Total number of burns of the vernier or additional stage engines.
	VernierNumBurns int64 `json:"vernierNumBurns"`
	// Total thrust of one of the vernier or additional engines at sea level in kN.
	VernierThrustSeaLevel float64 `json:"vernierThrustSeaLevel"`
	// Total thrust of one of the vernier or additional engines in a vacuum in kN.
	VernierThrustVacuum float64 `json:"vernierThrustVacuum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		IDEngine                 respjson.Field
		IDLaunchVehicle          respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		AvionicsNotes            respjson.Field
		BurnTime                 respjson.Field
		ControlThruster1         respjson.Field
		ControlThruster2         respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		Diameter                 respjson.Field
		Engine                   respjson.Field
		Length                   respjson.Field
		MainEngineThrustSeaLevel respjson.Field
		MainEngineThrustVacuum   respjson.Field
		ManufacturerOrgID        respjson.Field
		Mass                     respjson.Field
		Notes                    respjson.Field
		NumBurns                 respjson.Field
		NumControlThruster1      respjson.Field
		NumControlThruster2      respjson.Field
		NumEngines               respjson.Field
		NumStageElements         respjson.Field
		NumVernier               respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		PhotoURLs                respjson.Field
		Restartable              respjson.Field
		Reusable                 respjson.Field
		StageNumber              respjson.Field
		Tags                     respjson.Field
		ThrustSeaLevel           respjson.Field
		ThrustVacuum             respjson.Field
		Type                     respjson.Field
		UpdatedAt                respjson.Field
		UpdatedBy                respjson.Field
		Vernier                  respjson.Field
		VernierBurnTime          respjson.Field
		VernierNumBurns          respjson.Field
		VernierThrustSeaLevel    respjson.Field
		VernierThrustVacuum      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchVehicleGetResponseStage) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleGetResponseStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaunchVehicleQueryhelpResponse struct {
	AodrSupported         bool                                      `json:"aodrSupported"`
	ClassificationMarking string                                    `json:"classificationMarking"`
	Description           string                                    `json:"description"`
	HistorySupported      bool                                      `json:"historySupported"`
	Name                  string                                    `json:"name"`
	Parameters            []LaunchVehicleQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                  `json:"requiredRoles"`
	RestSupported         bool                                      `json:"restSupported"`
	SortSupported         bool                                      `json:"sortSupported"`
	TypeName              string                                    `json:"typeName"`
	Uri                   string                                    `json:"uri"`
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
func (r LaunchVehicleQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaunchVehicleQueryhelpResponseParameter struct {
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
func (r LaunchVehicleQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of basic information about known launch vehicles. A launch
// vehicle may have several details records including characteristics and
// information compiled by a particular source.
type LaunchVehicleTupleResponse struct {
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
	DataMode LaunchVehicleTupleResponseDataMode `json:"dataMode,required"`
	// Launch vehicle name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Read-only collection of additional LaunchVehicleDetails by various sources for
	// this launch vehicle, ignored on create/update. These details must be created
	// separately via the /udl/launchvehicledetails operations.
	LaunchVehicleDetails []LaunchVehicleTupleResponseLaunchVehicleDetail `json:"launchVehicleDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only collection of stages for this launch vehicle, ignored on
	// create/update.
	Stages []LaunchVehicleTupleResponseStage `json:"stages"`
	// Vehicle type.
	Type string `json:"type"`
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
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		LaunchVehicleDetails  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Stages                respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchVehicleTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleTupleResponse) UnmarshalJSON(data []byte) error {
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
type LaunchVehicleTupleResponseDataMode string

const (
	LaunchVehicleTupleResponseDataModeReal      LaunchVehicleTupleResponseDataMode = "REAL"
	LaunchVehicleTupleResponseDataModeTest      LaunchVehicleTupleResponseDataMode = "TEST"
	LaunchVehicleTupleResponseDataModeSimulated LaunchVehicleTupleResponseDataMode = "SIMULATED"
	LaunchVehicleTupleResponseDataModeExercise  LaunchVehicleTupleResponseDataMode = "EXERCISE"
)

// Model representation of launch vehicle details and characteristics, compiled by
// a particular source. A vehicle may have multiple details records from various
// sources.
type LaunchVehicleTupleResponseLaunchVehicleDetail struct {
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
	// Identifier of the parent launch vehicle record.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Launch vehicle attitude accuracy (degrees).
	AttitudeAccuracy float64 `json:"attitudeAccuracy"`
	// Vehicle category.
	Category string `json:"category"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Launch vehicle deployment rotation rate in RPM.
	DeploymentRotationRate float64 `json:"deploymentRotationRate"`
	// Vehicle diameter in meters.
	Diameter float64 `json:"diameter"`
	// Launch vehicle estimated launch price in US dollars.
	EstLaunchPrice float64 `json:"estLaunchPrice"`
	// Launch vehicle typical estimated launch price in US dollars.
	EstLaunchPriceTypical float64 `json:"estLaunchPriceTypical"`
	// Vehicle fairing maximum external diameter in meters.
	FairingExternalDiameter float64 `json:"fairingExternalDiameter"`
	// Vehicle fairing maximum internal diameter in meters.
	FairingInternalDiameter float64 `json:"fairingInternalDiameter"`
	// Vehicle fairing length in meters.
	FairingLength float64 `json:"fairingLength"`
	// Vehicle fairing mass in kg.
	FairingMass float64 `json:"fairingMass"`
	// Fairing material.
	FairingMaterial string `json:"fairingMaterial"`
	// Name of the fairing.
	FairingName string `json:"fairingName"`
	// Notes/Description of the launch vehicle fairing.
	FairingNotes string `json:"fairingNotes"`
	// Vehicle family.
	Family string `json:"family"`
	// Maximum vehicle payload mass to GEO orbit in kg.
	GeoPayloadMass float64 `json:"geoPayloadMass"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Apogee Margin (degrees).
	GtoInj3SigAccuracyApogeeMargin float64 `json:"gtoInj3SigAccuracyApogeeMargin"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Apogee Target (degrees).
	GtoInj3SigAccuracyApogeeTarget float64 `json:"gtoInj3SigAccuracyApogeeTarget"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Inclination Margin (degrees).
	GtoInj3SigAccuracyInclinationMargin float64 `json:"gtoInj3SigAccuracyInclinationMargin"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Inclination Target (degrees).
	GtoInj3SigAccuracyInclinationTarget float64 `json:"gtoInj3SigAccuracyInclinationTarget"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Perigee Margin (degrees).
	GtoInj3SigAccuracyPerigeeMargin float64 `json:"gtoInj3SigAccuracyPerigeeMargin"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Perigee Target (degrees).
	GtoInj3SigAccuracyPerigeeTarget float64 `json:"gtoInj3SigAccuracyPerigeeTarget"`
	// Max vehicle payload mass to Geo-Transfer Orbit in kg.
	GtoPayloadMass float64 `json:"gtoPayloadMass"`
	// Vehicle total mass at launch time in kg (including all boosters).
	LaunchMass float64 `json:"launchMass"`
	// Vehicle launch prefix.
	LaunchPrefix string `json:"launchPrefix"`
	// Vehicle length in meters.
	Length float64 `json:"length"`
	// Max vehicle payload mass to LEO orbit in kg.
	LeoPayloadMass float64 `json:"leoPayloadMass"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	ManufacturerOrg shared.OrganizationFull `json:"manufacturerOrg"`
	// ID of the organization that manufactures the launch vehicle.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Vehicle maximum acceleration load in g.
	MaxAccelLoad float64 `json:"maxAccelLoad"`
	// Vehicle maximum acoustic level in dB.
	MaxAcousticLevel float64 `json:"maxAcousticLevel"`
	// Vehicle maximum acoustic level range in Hz.
	MaxAcousticLevelRange float64 `json:"maxAcousticLevelRange"`
	// Vehicle fairing maximum pressure change in kPa/sec.
	MaxFairingPressureChange float64 `json:"maxFairingPressureChange"`
	// Vehicle maximum flight shock force in g.
	MaxFlightShockForce float64 `json:"maxFlightShockForce"`
	// Vehicle maximum flight shock frequency in Hz.
	MaxFlightShockFreq float64 `json:"maxFlightShockFreq"`
	// Vehicle maximum payload lateral frequency in Hz.
	MaxPayloadFreqLat float64 `json:"maxPayloadFreqLat"`
	// Vehicle maximum payload longitudinal frequency in Hz.
	MaxPayloadFreqLon float64 `json:"maxPayloadFreqLon"`
	// Vehicle minor variant.
	MinorVariant string `json:"minorVariant"`
	// Notes/Description of the launch vehicle.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Oxidizer type.
	Oxidizer string `json:"oxidizer"`
	// Notes/Description of the launch vehicle payload.
	PayloadNotes string `json:"payloadNotes"`
	// Launch vehicle payload separation rate in m/s.
	PayloadSeparationRate float64 `json:"payloadSeparationRate"`
	// Propellant type.
	Propellant string `json:"propellant"`
	// Vehicle overall sound pressure level in dB.
	SoundPressureLevel float64 `json:"soundPressureLevel"`
	// Optional URL for additional information on the vehicle.
	SourceURL string `json:"sourceURL"`
	// Max vehicle payload mass to Sun-Synchronous Orbit in kg.
	SSOPayloadMass float64 `json:"ssoPayloadMass"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Vehicle variant.
	Variant string `json:"variant"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking               respjson.Field
		DataMode                            respjson.Field
		IDLaunchVehicle                     respjson.Field
		Source                              respjson.Field
		ID                                  respjson.Field
		AttitudeAccuracy                    respjson.Field
		Category                            respjson.Field
		CreatedAt                           respjson.Field
		CreatedBy                           respjson.Field
		DeploymentRotationRate              respjson.Field
		Diameter                            respjson.Field
		EstLaunchPrice                      respjson.Field
		EstLaunchPriceTypical               respjson.Field
		FairingExternalDiameter             respjson.Field
		FairingInternalDiameter             respjson.Field
		FairingLength                       respjson.Field
		FairingMass                         respjson.Field
		FairingMaterial                     respjson.Field
		FairingName                         respjson.Field
		FairingNotes                        respjson.Field
		Family                              respjson.Field
		GeoPayloadMass                      respjson.Field
		GtoInj3SigAccuracyApogeeMargin      respjson.Field
		GtoInj3SigAccuracyApogeeTarget      respjson.Field
		GtoInj3SigAccuracyInclinationMargin respjson.Field
		GtoInj3SigAccuracyInclinationTarget respjson.Field
		GtoInj3SigAccuracyPerigeeMargin     respjson.Field
		GtoInj3SigAccuracyPerigeeTarget     respjson.Field
		GtoPayloadMass                      respjson.Field
		LaunchMass                          respjson.Field
		LaunchPrefix                        respjson.Field
		Length                              respjson.Field
		LeoPayloadMass                      respjson.Field
		ManufacturerOrg                     respjson.Field
		ManufacturerOrgID                   respjson.Field
		MaxAccelLoad                        respjson.Field
		MaxAcousticLevel                    respjson.Field
		MaxAcousticLevelRange               respjson.Field
		MaxFairingPressureChange            respjson.Field
		MaxFlightShockForce                 respjson.Field
		MaxFlightShockFreq                  respjson.Field
		MaxPayloadFreqLat                   respjson.Field
		MaxPayloadFreqLon                   respjson.Field
		MinorVariant                        respjson.Field
		Notes                               respjson.Field
		Origin                              respjson.Field
		OrigNetwork                         respjson.Field
		Oxidizer                            respjson.Field
		PayloadNotes                        respjson.Field
		PayloadSeparationRate               respjson.Field
		Propellant                          respjson.Field
		SoundPressureLevel                  respjson.Field
		SourceURL                           respjson.Field
		SSOPayloadMass                      respjson.Field
		Tags                                respjson.Field
		UpdatedAt                           respjson.Field
		UpdatedBy                           respjson.Field
		Variant                             respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchVehicleTupleResponseLaunchVehicleDetail) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleTupleResponseLaunchVehicleDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Launch stage information for a particular launch vehicle. A launch vehicle can
// have several stages, each with 1 to many engines.
type LaunchVehicleTupleResponseStage struct {
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
	// Identifier of the Engine record for this stage.
	IDEngine string `json:"idEngine,required"`
	// Identifier of the launch vehicle record for this stage.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Description/notes of the stage avionics.
	AvionicsNotes string `json:"avionicsNotes"`
	// Total burn time of the stage engines in seconds.
	BurnTime float64 `json:"burnTime"`
	// Control thruster 1 type.
	ControlThruster1 string `json:"controlThruster1"`
	// Control thruster 2 type.
	ControlThruster2 string `json:"controlThruster2"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Stage maximum external diameter in meters.
	Diameter float64 `json:"diameter"`
	// Known launch vehicle engines and their performance characteristics and limits. A
	// launch vehicle has 1 to many engines per stage.
	Engine shared.Engine `json:"engine"`
	// Stage length in meters.
	Length float64 `json:"length"`
	// Thrust of the stage main engine at sea level in kN.
	MainEngineThrustSeaLevel float64 `json:"mainEngineThrustSeaLevel"`
	// Thrust of the stage main engine in a vacuum in kN.
	MainEngineThrustVacuum float64 `json:"mainEngineThrustVacuum"`
	// ID of the organization that manufactures this launch stage.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Stage gross mass in kg.
	Mass float64 `json:"mass"`
	// Description/notes of the stage.
	Notes string `json:"notes"`
	// Number of burns for the stage engines.
	NumBurns int64 `json:"numBurns"`
	// Number of type control thruster 1.
	NumControlThruster1 int64 `json:"numControlThruster1"`
	// Number of type control thruster 2.
	NumControlThruster2 int64 `json:"numControlThruster2"`
	// The number of the specified engines on this launch stage.
	NumEngines int64 `json:"numEngines"`
	// Number of launch stage elements used in this stage.
	NumStageElements int64 `json:"numStageElements"`
	// Number of vernier or additional engines.
	NumVernier int64 `json:"numVernier"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Array of URLs of photos of the stage.
	PhotoURLs []string `json:"photoURLs"`
	// Boolean indicating if this launch stage can be restarted.
	Restartable bool `json:"restartable"`
	// Boolean indicating if this launch stage is reusable.
	Reusable bool `json:"reusable"`
	// The stage number of this launch stage.
	StageNumber int64 `json:"stageNumber"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Total thrust of the stage at sea level in kN.
	ThrustSeaLevel float64 `json:"thrustSeaLevel"`
	// Total thrust of the stage in a vacuum in kN.
	ThrustVacuum float64 `json:"thrustVacuum"`
	// Engine cycle type (e.g. Electrostatic Ion, Pressure Fed, Hall, Catalytic
	// Decomposition, etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Engine vernier or additional engine type.
	Vernier string `json:"vernier"`
	// Total burn time of the vernier or additional stage engines in seconds.
	VernierBurnTime float64 `json:"vernierBurnTime"`
	// Total number of burns of the vernier or additional stage engines.
	VernierNumBurns int64 `json:"vernierNumBurns"`
	// Total thrust of one of the vernier or additional engines at sea level in kN.
	VernierThrustSeaLevel float64 `json:"vernierThrustSeaLevel"`
	// Total thrust of one of the vernier or additional engines in a vacuum in kN.
	VernierThrustVacuum float64 `json:"vernierThrustVacuum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		IDEngine                 respjson.Field
		IDLaunchVehicle          respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		AvionicsNotes            respjson.Field
		BurnTime                 respjson.Field
		ControlThruster1         respjson.Field
		ControlThruster2         respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		Diameter                 respjson.Field
		Engine                   respjson.Field
		Length                   respjson.Field
		MainEngineThrustSeaLevel respjson.Field
		MainEngineThrustVacuum   respjson.Field
		ManufacturerOrgID        respjson.Field
		Mass                     respjson.Field
		Notes                    respjson.Field
		NumBurns                 respjson.Field
		NumControlThruster1      respjson.Field
		NumControlThruster2      respjson.Field
		NumEngines               respjson.Field
		NumStageElements         respjson.Field
		NumVernier               respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		PhotoURLs                respjson.Field
		Restartable              respjson.Field
		Reusable                 respjson.Field
		StageNumber              respjson.Field
		Tags                     respjson.Field
		ThrustSeaLevel           respjson.Field
		ThrustVacuum             respjson.Field
		Type                     respjson.Field
		UpdatedAt                respjson.Field
		UpdatedBy                respjson.Field
		Vernier                  respjson.Field
		VernierBurnTime          respjson.Field
		VernierNumBurns          respjson.Field
		VernierThrustSeaLevel    respjson.Field
		VernierThrustVacuum      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchVehicleTupleResponseStage) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleTupleResponseStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaunchVehicleNewParams struct {
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
	DataMode LaunchVehicleNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Launch vehicle name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Vehicle type.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r LaunchVehicleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchVehicleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchVehicleNewParams) UnmarshalJSON(data []byte) error {
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
type LaunchVehicleNewParamsDataMode string

const (
	LaunchVehicleNewParamsDataModeReal      LaunchVehicleNewParamsDataMode = "REAL"
	LaunchVehicleNewParamsDataModeTest      LaunchVehicleNewParamsDataMode = "TEST"
	LaunchVehicleNewParamsDataModeSimulated LaunchVehicleNewParamsDataMode = "SIMULATED"
	LaunchVehicleNewParamsDataModeExercise  LaunchVehicleNewParamsDataMode = "EXERCISE"
)

type LaunchVehicleUpdateParams struct {
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
	DataMode LaunchVehicleUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Launch vehicle name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Vehicle type.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r LaunchVehicleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchVehicleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchVehicleUpdateParams) UnmarshalJSON(data []byte) error {
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
type LaunchVehicleUpdateParamsDataMode string

const (
	LaunchVehicleUpdateParamsDataModeReal      LaunchVehicleUpdateParamsDataMode = "REAL"
	LaunchVehicleUpdateParamsDataModeTest      LaunchVehicleUpdateParamsDataMode = "TEST"
	LaunchVehicleUpdateParamsDataModeSimulated LaunchVehicleUpdateParamsDataMode = "SIMULATED"
	LaunchVehicleUpdateParamsDataModeExercise  LaunchVehicleUpdateParamsDataMode = "EXERCISE"
)

type LaunchVehicleListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchVehicleListParams]'s query parameters as
// `url.Values`.
func (r LaunchVehicleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchVehicleCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchVehicleCountParams]'s query parameters as
// `url.Values`.
func (r LaunchVehicleCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchVehicleGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchVehicleGetParams]'s query parameters as `url.Values`.
func (r LaunchVehicleGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchVehicleTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchVehicleTupleParams]'s query parameters as
// `url.Values`.
func (r LaunchVehicleTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
