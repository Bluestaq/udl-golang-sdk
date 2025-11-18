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

// LaunchVehicleDetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaunchVehicleDetailService] method instead.
type LaunchVehicleDetailService struct {
	Options []option.RequestOption
}

// NewLaunchVehicleDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLaunchVehicleDetailService(opts ...option.RequestOption) (r LaunchVehicleDetailService) {
	r = LaunchVehicleDetailService{}
	r.Options = opts
	return
}

// Service operation to take a single LaunchVehicleDetails as a POST body and
// ingest into the database. LaunchVehicleDetails represents launch vehicle details
// and characteristics, compiled by a particular source. A vehicle may have
// multiple details records from various sources. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *LaunchVehicleDetailService) New(ctx context.Context, body LaunchVehicleDetailNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/launchvehicledetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single LaunchVehicleDetails. LaunchVehicleDetails
// represents launch vehicle details and characteristics, compiled by a particular
// source. A vehicle may have multiple details records from various sources. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *LaunchVehicleDetailService) Update(ctx context.Context, id string, body LaunchVehicleDetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchvehicledetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaunchVehicleDetailService) List(ctx context.Context, query LaunchVehicleDetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaunchVehicleDetailListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/launchvehicledetails"
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
func (r *LaunchVehicleDetailService) ListAutoPaging(ctx context.Context, query LaunchVehicleDetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaunchVehicleDetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a LaunchVehicleDetails object specified by the
// passed ID path parameter. LaunchVehicleDetails represents launch vehicle details
// and characteristics, compiled by a particular source. A vehicle may have
// multiple details records from various sources. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *LaunchVehicleDetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchvehicledetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to get a single LaunchVehicleDetails record by its unique ID
// passed as a path parameter. LaunchVehicleDetails represents launch vehicle
// details and characteristics, compiled by a particular source. A vehicle may have
// multiple details records from various sources.
func (r *LaunchVehicleDetailService) Get(ctx context.Context, id string, query LaunchVehicleDetailGetParams, opts ...option.RequestOption) (res *LaunchVehicleDetailGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchvehicledetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of launch vehicle details and characteristics, compiled by
// a particular source. A vehicle may have multiple details records from various
// sources.
type LaunchVehicleDetailListResponse struct {
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
	DataMode LaunchVehicleDetailListResponseDataMode `json:"dataMode,required"`
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
		Variant                             respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchVehicleDetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleDetailListResponse) UnmarshalJSON(data []byte) error {
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
type LaunchVehicleDetailListResponseDataMode string

const (
	LaunchVehicleDetailListResponseDataModeReal      LaunchVehicleDetailListResponseDataMode = "REAL"
	LaunchVehicleDetailListResponseDataModeTest      LaunchVehicleDetailListResponseDataMode = "TEST"
	LaunchVehicleDetailListResponseDataModeSimulated LaunchVehicleDetailListResponseDataMode = "SIMULATED"
	LaunchVehicleDetailListResponseDataModeExercise  LaunchVehicleDetailListResponseDataMode = "EXERCISE"
)

// Model representation of launch vehicle details and characteristics, compiled by
// a particular source. A vehicle may have multiple details records from various
// sources.
type LaunchVehicleDetailGetResponse struct {
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
	DataMode LaunchVehicleDetailGetResponseDataMode `json:"dataMode,required"`
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
func (r LaunchVehicleDetailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchVehicleDetailGetResponse) UnmarshalJSON(data []byte) error {
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
type LaunchVehicleDetailGetResponseDataMode string

const (
	LaunchVehicleDetailGetResponseDataModeReal      LaunchVehicleDetailGetResponseDataMode = "REAL"
	LaunchVehicleDetailGetResponseDataModeTest      LaunchVehicleDetailGetResponseDataMode = "TEST"
	LaunchVehicleDetailGetResponseDataModeSimulated LaunchVehicleDetailGetResponseDataMode = "SIMULATED"
	LaunchVehicleDetailGetResponseDataModeExercise  LaunchVehicleDetailGetResponseDataMode = "EXERCISE"
)

type LaunchVehicleDetailNewParams struct {
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
	DataMode LaunchVehicleDetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the parent launch vehicle record.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Launch vehicle attitude accuracy (degrees).
	AttitudeAccuracy param.Opt[float64] `json:"attitudeAccuracy,omitzero"`
	// Vehicle category.
	Category param.Opt[string] `json:"category,omitzero"`
	// Launch vehicle deployment rotation rate in RPM.
	DeploymentRotationRate param.Opt[float64] `json:"deploymentRotationRate,omitzero"`
	// Vehicle diameter in meters.
	Diameter param.Opt[float64] `json:"diameter,omitzero"`
	// Launch vehicle estimated launch price in US dollars.
	EstLaunchPrice param.Opt[float64] `json:"estLaunchPrice,omitzero"`
	// Launch vehicle typical estimated launch price in US dollars.
	EstLaunchPriceTypical param.Opt[float64] `json:"estLaunchPriceTypical,omitzero"`
	// Vehicle fairing maximum external diameter in meters.
	FairingExternalDiameter param.Opt[float64] `json:"fairingExternalDiameter,omitzero"`
	// Vehicle fairing maximum internal diameter in meters.
	FairingInternalDiameter param.Opt[float64] `json:"fairingInternalDiameter,omitzero"`
	// Vehicle fairing length in meters.
	FairingLength param.Opt[float64] `json:"fairingLength,omitzero"`
	// Vehicle fairing mass in kg.
	FairingMass param.Opt[float64] `json:"fairingMass,omitzero"`
	// Fairing material.
	FairingMaterial param.Opt[string] `json:"fairingMaterial,omitzero"`
	// Name of the fairing.
	FairingName param.Opt[string] `json:"fairingName,omitzero"`
	// Notes/Description of the launch vehicle fairing.
	FairingNotes param.Opt[string] `json:"fairingNotes,omitzero"`
	// Vehicle family.
	Family param.Opt[string] `json:"family,omitzero"`
	// Maximum vehicle payload mass to GEO orbit in kg.
	GeoPayloadMass param.Opt[float64] `json:"geoPayloadMass,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Apogee Margin (degrees).
	GtoInj3SigAccuracyApogeeMargin param.Opt[float64] `json:"gtoInj3SigAccuracyApogeeMargin,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Apogee Target (degrees).
	GtoInj3SigAccuracyApogeeTarget param.Opt[float64] `json:"gtoInj3SigAccuracyApogeeTarget,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Inclination Margin (degrees).
	GtoInj3SigAccuracyInclinationMargin param.Opt[float64] `json:"gtoInj3SigAccuracyInclinationMargin,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Inclination Target (degrees).
	GtoInj3SigAccuracyInclinationTarget param.Opt[float64] `json:"gtoInj3SigAccuracyInclinationTarget,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Perigee Margin (degrees).
	GtoInj3SigAccuracyPerigeeMargin param.Opt[float64] `json:"gtoInj3SigAccuracyPerigeeMargin,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Perigee Target (degrees).
	GtoInj3SigAccuracyPerigeeTarget param.Opt[float64] `json:"gtoInj3SigAccuracyPerigeeTarget,omitzero"`
	// Max vehicle payload mass to Geo-Transfer Orbit in kg.
	GtoPayloadMass param.Opt[float64] `json:"gtoPayloadMass,omitzero"`
	// Vehicle total mass at launch time in kg (including all boosters).
	LaunchMass param.Opt[float64] `json:"launchMass,omitzero"`
	// Vehicle launch prefix.
	LaunchPrefix param.Opt[string] `json:"launchPrefix,omitzero"`
	// Vehicle length in meters.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Max vehicle payload mass to LEO orbit in kg.
	LeoPayloadMass param.Opt[float64] `json:"leoPayloadMass,omitzero"`
	// ID of the organization that manufactures the launch vehicle.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Vehicle maximum acceleration load in g.
	MaxAccelLoad param.Opt[float64] `json:"maxAccelLoad,omitzero"`
	// Vehicle maximum acoustic level in dB.
	MaxAcousticLevel param.Opt[float64] `json:"maxAcousticLevel,omitzero"`
	// Vehicle maximum acoustic level range in Hz.
	MaxAcousticLevelRange param.Opt[float64] `json:"maxAcousticLevelRange,omitzero"`
	// Vehicle fairing maximum pressure change in kPa/sec.
	MaxFairingPressureChange param.Opt[float64] `json:"maxFairingPressureChange,omitzero"`
	// Vehicle maximum flight shock force in g.
	MaxFlightShockForce param.Opt[float64] `json:"maxFlightShockForce,omitzero"`
	// Vehicle maximum flight shock frequency in Hz.
	MaxFlightShockFreq param.Opt[float64] `json:"maxFlightShockFreq,omitzero"`
	// Vehicle maximum payload lateral frequency in Hz.
	MaxPayloadFreqLat param.Opt[float64] `json:"maxPayloadFreqLat,omitzero"`
	// Vehicle maximum payload longitudinal frequency in Hz.
	MaxPayloadFreqLon param.Opt[float64] `json:"maxPayloadFreqLon,omitzero"`
	// Vehicle minor variant.
	MinorVariant param.Opt[string] `json:"minorVariant,omitzero"`
	// Notes/Description of the launch vehicle.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Oxidizer type.
	Oxidizer param.Opt[string] `json:"oxidizer,omitzero"`
	// Notes/Description of the launch vehicle payload.
	PayloadNotes param.Opt[string] `json:"payloadNotes,omitzero"`
	// Launch vehicle payload separation rate in m/s.
	PayloadSeparationRate param.Opt[float64] `json:"payloadSeparationRate,omitzero"`
	// Propellant type.
	Propellant param.Opt[string] `json:"propellant,omitzero"`
	// Vehicle overall sound pressure level in dB.
	SoundPressureLevel param.Opt[float64] `json:"soundPressureLevel,omitzero"`
	// Optional URL for additional information on the vehicle.
	SourceURL param.Opt[string] `json:"sourceURL,omitzero"`
	// Max vehicle payload mass to Sun-Synchronous Orbit in kg.
	SSOPayloadMass param.Opt[float64] `json:"ssoPayloadMass,omitzero"`
	// Vehicle variant.
	Variant param.Opt[string] `json:"variant,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r LaunchVehicleDetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchVehicleDetailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchVehicleDetailNewParams) UnmarshalJSON(data []byte) error {
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
type LaunchVehicleDetailNewParamsDataMode string

const (
	LaunchVehicleDetailNewParamsDataModeReal      LaunchVehicleDetailNewParamsDataMode = "REAL"
	LaunchVehicleDetailNewParamsDataModeTest      LaunchVehicleDetailNewParamsDataMode = "TEST"
	LaunchVehicleDetailNewParamsDataModeSimulated LaunchVehicleDetailNewParamsDataMode = "SIMULATED"
	LaunchVehicleDetailNewParamsDataModeExercise  LaunchVehicleDetailNewParamsDataMode = "EXERCISE"
)

type LaunchVehicleDetailUpdateParams struct {
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
	DataMode LaunchVehicleDetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the parent launch vehicle record.
	IDLaunchVehicle string `json:"idLaunchVehicle,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Launch vehicle attitude accuracy (degrees).
	AttitudeAccuracy param.Opt[float64] `json:"attitudeAccuracy,omitzero"`
	// Vehicle category.
	Category param.Opt[string] `json:"category,omitzero"`
	// Launch vehicle deployment rotation rate in RPM.
	DeploymentRotationRate param.Opt[float64] `json:"deploymentRotationRate,omitzero"`
	// Vehicle diameter in meters.
	Diameter param.Opt[float64] `json:"diameter,omitzero"`
	// Launch vehicle estimated launch price in US dollars.
	EstLaunchPrice param.Opt[float64] `json:"estLaunchPrice,omitzero"`
	// Launch vehicle typical estimated launch price in US dollars.
	EstLaunchPriceTypical param.Opt[float64] `json:"estLaunchPriceTypical,omitzero"`
	// Vehicle fairing maximum external diameter in meters.
	FairingExternalDiameter param.Opt[float64] `json:"fairingExternalDiameter,omitzero"`
	// Vehicle fairing maximum internal diameter in meters.
	FairingInternalDiameter param.Opt[float64] `json:"fairingInternalDiameter,omitzero"`
	// Vehicle fairing length in meters.
	FairingLength param.Opt[float64] `json:"fairingLength,omitzero"`
	// Vehicle fairing mass in kg.
	FairingMass param.Opt[float64] `json:"fairingMass,omitzero"`
	// Fairing material.
	FairingMaterial param.Opt[string] `json:"fairingMaterial,omitzero"`
	// Name of the fairing.
	FairingName param.Opt[string] `json:"fairingName,omitzero"`
	// Notes/Description of the launch vehicle fairing.
	FairingNotes param.Opt[string] `json:"fairingNotes,omitzero"`
	// Vehicle family.
	Family param.Opt[string] `json:"family,omitzero"`
	// Maximum vehicle payload mass to GEO orbit in kg.
	GeoPayloadMass param.Opt[float64] `json:"geoPayloadMass,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Apogee Margin (degrees).
	GtoInj3SigAccuracyApogeeMargin param.Opt[float64] `json:"gtoInj3SigAccuracyApogeeMargin,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Apogee Target (degrees).
	GtoInj3SigAccuracyApogeeTarget param.Opt[float64] `json:"gtoInj3SigAccuracyApogeeTarget,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Inclination Margin (degrees).
	GtoInj3SigAccuracyInclinationMargin param.Opt[float64] `json:"gtoInj3SigAccuracyInclinationMargin,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Inclination Target (degrees).
	GtoInj3SigAccuracyInclinationTarget param.Opt[float64] `json:"gtoInj3SigAccuracyInclinationTarget,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Perigee Margin (degrees).
	GtoInj3SigAccuracyPerigeeMargin param.Opt[float64] `json:"gtoInj3SigAccuracyPerigeeMargin,omitzero"`
	// Launch vehicle GTO Injection 3 Sigma Accuracy Perigee Target (degrees).
	GtoInj3SigAccuracyPerigeeTarget param.Opt[float64] `json:"gtoInj3SigAccuracyPerigeeTarget,omitzero"`
	// Max vehicle payload mass to Geo-Transfer Orbit in kg.
	GtoPayloadMass param.Opt[float64] `json:"gtoPayloadMass,omitzero"`
	// Vehicle total mass at launch time in kg (including all boosters).
	LaunchMass param.Opt[float64] `json:"launchMass,omitzero"`
	// Vehicle launch prefix.
	LaunchPrefix param.Opt[string] `json:"launchPrefix,omitzero"`
	// Vehicle length in meters.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Max vehicle payload mass to LEO orbit in kg.
	LeoPayloadMass param.Opt[float64] `json:"leoPayloadMass,omitzero"`
	// ID of the organization that manufactures the launch vehicle.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Vehicle maximum acceleration load in g.
	MaxAccelLoad param.Opt[float64] `json:"maxAccelLoad,omitzero"`
	// Vehicle maximum acoustic level in dB.
	MaxAcousticLevel param.Opt[float64] `json:"maxAcousticLevel,omitzero"`
	// Vehicle maximum acoustic level range in Hz.
	MaxAcousticLevelRange param.Opt[float64] `json:"maxAcousticLevelRange,omitzero"`
	// Vehicle fairing maximum pressure change in kPa/sec.
	MaxFairingPressureChange param.Opt[float64] `json:"maxFairingPressureChange,omitzero"`
	// Vehicle maximum flight shock force in g.
	MaxFlightShockForce param.Opt[float64] `json:"maxFlightShockForce,omitzero"`
	// Vehicle maximum flight shock frequency in Hz.
	MaxFlightShockFreq param.Opt[float64] `json:"maxFlightShockFreq,omitzero"`
	// Vehicle maximum payload lateral frequency in Hz.
	MaxPayloadFreqLat param.Opt[float64] `json:"maxPayloadFreqLat,omitzero"`
	// Vehicle maximum payload longitudinal frequency in Hz.
	MaxPayloadFreqLon param.Opt[float64] `json:"maxPayloadFreqLon,omitzero"`
	// Vehicle minor variant.
	MinorVariant param.Opt[string] `json:"minorVariant,omitzero"`
	// Notes/Description of the launch vehicle.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Oxidizer type.
	Oxidizer param.Opt[string] `json:"oxidizer,omitzero"`
	// Notes/Description of the launch vehicle payload.
	PayloadNotes param.Opt[string] `json:"payloadNotes,omitzero"`
	// Launch vehicle payload separation rate in m/s.
	PayloadSeparationRate param.Opt[float64] `json:"payloadSeparationRate,omitzero"`
	// Propellant type.
	Propellant param.Opt[string] `json:"propellant,omitzero"`
	// Vehicle overall sound pressure level in dB.
	SoundPressureLevel param.Opt[float64] `json:"soundPressureLevel,omitzero"`
	// Optional URL for additional information on the vehicle.
	SourceURL param.Opt[string] `json:"sourceURL,omitzero"`
	// Max vehicle payload mass to Sun-Synchronous Orbit in kg.
	SSOPayloadMass param.Opt[float64] `json:"ssoPayloadMass,omitzero"`
	// Vehicle variant.
	Variant param.Opt[string] `json:"variant,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r LaunchVehicleDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchVehicleDetailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchVehicleDetailUpdateParams) UnmarshalJSON(data []byte) error {
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
type LaunchVehicleDetailUpdateParamsDataMode string

const (
	LaunchVehicleDetailUpdateParamsDataModeReal      LaunchVehicleDetailUpdateParamsDataMode = "REAL"
	LaunchVehicleDetailUpdateParamsDataModeTest      LaunchVehicleDetailUpdateParamsDataMode = "TEST"
	LaunchVehicleDetailUpdateParamsDataModeSimulated LaunchVehicleDetailUpdateParamsDataMode = "SIMULATED"
	LaunchVehicleDetailUpdateParamsDataModeExercise  LaunchVehicleDetailUpdateParamsDataMode = "EXERCISE"
)

type LaunchVehicleDetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchVehicleDetailListParams]'s query parameters as
// `url.Values`.
func (r LaunchVehicleDetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchVehicleDetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchVehicleDetailGetParams]'s query parameters as
// `url.Values`.
func (r LaunchVehicleDetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
