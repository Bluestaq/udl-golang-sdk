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

// SeraDataEarlyWarningService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeraDataEarlyWarningService] method instead.
type SeraDataEarlyWarningService struct {
	Options []option.RequestOption
}

// NewSeraDataEarlyWarningService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSeraDataEarlyWarningService(opts ...option.RequestOption) (r SeraDataEarlyWarningService) {
	r = SeraDataEarlyWarningService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataEarlyWarning as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeraDataEarlyWarningService) New(ctx context.Context, body SeraDataEarlyWarningNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradataearlywarning"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataEarlyWarning. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *SeraDataEarlyWarningService) Update(ctx context.Context, id string, body SeraDataEarlyWarningUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataearlywarning/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SeraDataEarlyWarningService) List(ctx context.Context, query SeraDataEarlyWarningListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeraDataEarlyWarningListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/seradataearlywarning"
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
func (r *SeraDataEarlyWarningService) ListAutoPaging(ctx context.Context, query SeraDataEarlyWarningListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeraDataEarlyWarningListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataEarlyWarning specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SeraDataEarlyWarningService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataearlywarning/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SeraDataEarlyWarningService) Count(ctx context.Context, query SeraDataEarlyWarningCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradataearlywarning/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataEarlyWarning by its unique ID passed
// as a path parameter.
func (r *SeraDataEarlyWarningService) Get(ctx context.Context, id string, query SeraDataEarlyWarningGetParams, opts ...option.RequestOption) (res *SeraDataEarlyWarningGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataearlywarning/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SeraDataEarlyWarningService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradataearlywarning/queryhelp"
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
func (r *SeraDataEarlyWarningService) Tuple(ctx context.Context, query SeraDataEarlyWarningTupleParams, opts ...option.RequestOption) (res *[]SeraDataEarlyWarningTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/seradataearlywarning/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details for an early warning payload from Seradata.
type SeraDataEarlyWarningListResponse struct {
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
	DataMode SeraDataEarlyWarningListResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking              resp.Field
		DataMode                           resp.Field
		Source                             resp.Field
		SpacecraftID                       resp.Field
		ID                                 resp.Field
		BestResolution                     resp.Field
		CreatedAt                          resp.Field
		CreatedBy                          resp.Field
		EarthPointing                      resp.Field
		FrequencyLimits                    resp.Field
		GroundStationLocations             resp.Field
		GroundStations                     resp.Field
		HostedForCompanyOrgID              resp.Field
		IDIr                               resp.Field
		ManufacturerOrgID                  resp.Field
		MissileLaunchPhaseDetectionAbility resp.Field
		Name                               resp.Field
		Origin                             resp.Field
		OrigNetwork                        resp.Field
		PartnerSpacecraftID                resp.Field
		PayloadNotes                       resp.Field
		SpectralBands                      resp.Field
		ExtraFields                        map[string]resp.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataEarlyWarningListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataEarlyWarningListResponse) UnmarshalJSON(data []byte) error {
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
type SeraDataEarlyWarningListResponseDataMode string

const (
	SeraDataEarlyWarningListResponseDataModeReal      SeraDataEarlyWarningListResponseDataMode = "REAL"
	SeraDataEarlyWarningListResponseDataModeTest      SeraDataEarlyWarningListResponseDataMode = "TEST"
	SeraDataEarlyWarningListResponseDataModeSimulated SeraDataEarlyWarningListResponseDataMode = "SIMULATED"
	SeraDataEarlyWarningListResponseDataModeExercise  SeraDataEarlyWarningListResponseDataMode = "EXERCISE"
)

// Details for an early warning payload from Seradata.
type SeraDataEarlyWarningGetResponse struct {
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
	DataMode SeraDataEarlyWarningGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking              resp.Field
		DataMode                           resp.Field
		Source                             resp.Field
		SpacecraftID                       resp.Field
		ID                                 resp.Field
		BestResolution                     resp.Field
		CreatedAt                          resp.Field
		CreatedBy                          resp.Field
		EarthPointing                      resp.Field
		FrequencyLimits                    resp.Field
		GroundStationLocations             resp.Field
		GroundStations                     resp.Field
		HostedForCompanyOrgID              resp.Field
		IDIr                               resp.Field
		ManufacturerOrgID                  resp.Field
		MissileLaunchPhaseDetectionAbility resp.Field
		Name                               resp.Field
		Origin                             resp.Field
		OrigNetwork                        resp.Field
		PartnerSpacecraftID                resp.Field
		PayloadNotes                       resp.Field
		SpectralBands                      resp.Field
		UpdatedAt                          resp.Field
		UpdatedBy                          resp.Field
		ExtraFields                        map[string]resp.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataEarlyWarningGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataEarlyWarningGetResponse) UnmarshalJSON(data []byte) error {
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
type SeraDataEarlyWarningGetResponseDataMode string

const (
	SeraDataEarlyWarningGetResponseDataModeReal      SeraDataEarlyWarningGetResponseDataMode = "REAL"
	SeraDataEarlyWarningGetResponseDataModeTest      SeraDataEarlyWarningGetResponseDataMode = "TEST"
	SeraDataEarlyWarningGetResponseDataModeSimulated SeraDataEarlyWarningGetResponseDataMode = "SIMULATED"
	SeraDataEarlyWarningGetResponseDataModeExercise  SeraDataEarlyWarningGetResponseDataMode = "EXERCISE"
)

// Details for an early warning payload from Seradata.
type SeraDataEarlyWarningTupleResponse struct {
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
	DataMode SeraDataEarlyWarningTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking              resp.Field
		DataMode                           resp.Field
		Source                             resp.Field
		SpacecraftID                       resp.Field
		ID                                 resp.Field
		BestResolution                     resp.Field
		CreatedAt                          resp.Field
		CreatedBy                          resp.Field
		EarthPointing                      resp.Field
		FrequencyLimits                    resp.Field
		GroundStationLocations             resp.Field
		GroundStations                     resp.Field
		HostedForCompanyOrgID              resp.Field
		IDIr                               resp.Field
		ManufacturerOrgID                  resp.Field
		MissileLaunchPhaseDetectionAbility resp.Field
		Name                               resp.Field
		Origin                             resp.Field
		OrigNetwork                        resp.Field
		PartnerSpacecraftID                resp.Field
		PayloadNotes                       resp.Field
		SpectralBands                      resp.Field
		UpdatedAt                          resp.Field
		UpdatedBy                          resp.Field
		ExtraFields                        map[string]resp.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataEarlyWarningTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataEarlyWarningTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeraDataEarlyWarningTupleResponseDataMode string

const (
	SeraDataEarlyWarningTupleResponseDataModeReal      SeraDataEarlyWarningTupleResponseDataMode = "REAL"
	SeraDataEarlyWarningTupleResponseDataModeTest      SeraDataEarlyWarningTupleResponseDataMode = "TEST"
	SeraDataEarlyWarningTupleResponseDataModeSimulated SeraDataEarlyWarningTupleResponseDataMode = "SIMULATED"
	SeraDataEarlyWarningTupleResponseDataModeExercise  SeraDataEarlyWarningTupleResponseDataMode = "EXERCISE"
)

type SeraDataEarlyWarningNewParams struct {
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
	DataMode SeraDataEarlyWarningNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Best resolution for this IR in meters.
	BestResolution param.Opt[float64] `json:"bestResolution,omitzero"`
	// Is the sensor Earth Pointing.
	EarthPointing param.Opt[bool] `json:"earthPointing,omitzero"`
	// Frequency Limits for this IR.
	FrequencyLimits param.Opt[string] `json:"frequencyLimits,omitzero"`
	// Ground Station Locations for this IR.
	GroundStationLocations param.Opt[string] `json:"groundStationLocations,omitzero"`
	// Ground Station info for this IR.
	GroundStations param.Opt[string] `json:"groundStations,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the parent IR record.
	IDIr param.Opt[string] `json:"idIR,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Missile Launch Phase Detection Ability.
	MissileLaunchPhaseDetectionAbility param.Opt[string] `json:"missileLaunchPhaseDetectionAbility,omitzero"`
	// Sensor name from Seradata, e.g. Infra red telescope, Schmidt Telescope, etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID param.Opt[string] `json:"partnerSpacecraftId,omitzero"`
	// Payload notes.
	PayloadNotes param.Opt[string] `json:"payloadNotes,omitzero"`
	// Spectral Bands, e.g. Infra-Red.
	SpectralBands param.Opt[string] `json:"spectralBands,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeraDataEarlyWarningNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SeraDataEarlyWarningNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeraDataEarlyWarningNewParams
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
type SeraDataEarlyWarningNewParamsDataMode string

const (
	SeraDataEarlyWarningNewParamsDataModeReal      SeraDataEarlyWarningNewParamsDataMode = "REAL"
	SeraDataEarlyWarningNewParamsDataModeTest      SeraDataEarlyWarningNewParamsDataMode = "TEST"
	SeraDataEarlyWarningNewParamsDataModeSimulated SeraDataEarlyWarningNewParamsDataMode = "SIMULATED"
	SeraDataEarlyWarningNewParamsDataModeExercise  SeraDataEarlyWarningNewParamsDataMode = "EXERCISE"
)

type SeraDataEarlyWarningUpdateParams struct {
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
	DataMode SeraDataEarlyWarningUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Best resolution for this IR in meters.
	BestResolution param.Opt[float64] `json:"bestResolution,omitzero"`
	// Is the sensor Earth Pointing.
	EarthPointing param.Opt[bool] `json:"earthPointing,omitzero"`
	// Frequency Limits for this IR.
	FrequencyLimits param.Opt[string] `json:"frequencyLimits,omitzero"`
	// Ground Station Locations for this IR.
	GroundStationLocations param.Opt[string] `json:"groundStationLocations,omitzero"`
	// Ground Station info for this IR.
	GroundStations param.Opt[string] `json:"groundStations,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the parent IR record.
	IDIr param.Opt[string] `json:"idIR,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Missile Launch Phase Detection Ability.
	MissileLaunchPhaseDetectionAbility param.Opt[string] `json:"missileLaunchPhaseDetectionAbility,omitzero"`
	// Sensor name from Seradata, e.g. Infra red telescope, Schmidt Telescope, etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID param.Opt[string] `json:"partnerSpacecraftId,omitzero"`
	// Payload notes.
	PayloadNotes param.Opt[string] `json:"payloadNotes,omitzero"`
	// Spectral Bands, e.g. Infra-Red.
	SpectralBands param.Opt[string] `json:"spectralBands,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeraDataEarlyWarningUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SeraDataEarlyWarningUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeraDataEarlyWarningUpdateParams
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
type SeraDataEarlyWarningUpdateParamsDataMode string

const (
	SeraDataEarlyWarningUpdateParamsDataModeReal      SeraDataEarlyWarningUpdateParamsDataMode = "REAL"
	SeraDataEarlyWarningUpdateParamsDataModeTest      SeraDataEarlyWarningUpdateParamsDataMode = "TEST"
	SeraDataEarlyWarningUpdateParamsDataModeSimulated SeraDataEarlyWarningUpdateParamsDataMode = "SIMULATED"
	SeraDataEarlyWarningUpdateParamsDataModeExercise  SeraDataEarlyWarningUpdateParamsDataMode = "EXERCISE"
)

type SeraDataEarlyWarningListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeraDataEarlyWarningListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeraDataEarlyWarningListParams]'s query parameters as
// `url.Values`.
func (r SeraDataEarlyWarningListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeraDataEarlyWarningCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeraDataEarlyWarningCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeraDataEarlyWarningCountParams]'s query parameters as
// `url.Values`.
func (r SeraDataEarlyWarningCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeraDataEarlyWarningGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeraDataEarlyWarningGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeraDataEarlyWarningGetParams]'s query parameters as
// `url.Values`.
func (r SeraDataEarlyWarningGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeraDataEarlyWarningTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeraDataEarlyWarningTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeraDataEarlyWarningTupleParams]'s query parameters as
// `url.Values`.
func (r SeraDataEarlyWarningTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
