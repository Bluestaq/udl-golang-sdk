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

// SeradataOpticalPayloadService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeradataOpticalPayloadService] method instead.
type SeradataOpticalPayloadService struct {
	Options []option.RequestOption
}

// NewSeradataOpticalPayloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSeradataOpticalPayloadService(opts ...option.RequestOption) (r SeradataOpticalPayloadService) {
	r = SeradataOpticalPayloadService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataOpticalPayload as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeradataOpticalPayloadService) New(ctx context.Context, body SeradataOpticalPayloadNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/seradataopticalpayload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataOpticalPayload. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SeradataOpticalPayloadService) Update(ctx context.Context, id string, body SeradataOpticalPayloadUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataopticalpayload/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SeradataOpticalPayloadService) List(ctx context.Context, query SeradataOpticalPayloadListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeradataOpticalPayloadListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/seradataopticalpayload"
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
func (r *SeradataOpticalPayloadService) ListAutoPaging(ctx context.Context, query SeradataOpticalPayloadListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeradataOpticalPayloadListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataOpticalPayload specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SeradataOpticalPayloadService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataopticalpayload/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SeradataOpticalPayloadService) Count(ctx context.Context, query SeradataOpticalPayloadCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradataopticalpayload/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataOpticalPayload by its unique ID passed
// as a path parameter.
func (r *SeradataOpticalPayloadService) Get(ctx context.Context, id string, query SeradataOpticalPayloadGetParams, opts ...option.RequestOption) (res *SeradataOpticalPayloadGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataopticalpayload/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SeradataOpticalPayloadService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SeradataOpticalPayloadQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/seradataopticalpayload/queryhelp"
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
func (r *SeradataOpticalPayloadService) Tuple(ctx context.Context, query SeradataOpticalPayloadTupleParams, opts ...option.RequestOption) (res *[]SeradataOpticalPayloadTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/seradataopticalpayload/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details for an optical payload from Seradata.
type SeradataOpticalPayloadListResponse struct {
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
	DataMode SeradataOpticalPayloadListResponseDataMode `json:"dataMode,required"`
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
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataOpticalPayloadListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataOpticalPayloadListResponse) UnmarshalJSON(data []byte) error {
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
type SeradataOpticalPayloadListResponseDataMode string

const (
	SeradataOpticalPayloadListResponseDataModeReal      SeradataOpticalPayloadListResponseDataMode = "REAL"
	SeradataOpticalPayloadListResponseDataModeTest      SeradataOpticalPayloadListResponseDataMode = "TEST"
	SeradataOpticalPayloadListResponseDataModeSimulated SeradataOpticalPayloadListResponseDataMode = "SIMULATED"
	SeradataOpticalPayloadListResponseDataModeExercise  SeradataOpticalPayloadListResponseDataMode = "EXERCISE"
)

// Details for an optical payload from Seradata.
type SeradataOpticalPayloadGetResponse struct {
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
	DataMode SeradataOpticalPayloadGetResponseDataMode `json:"dataMode,required"`
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
func (r SeradataOpticalPayloadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataOpticalPayloadGetResponse) UnmarshalJSON(data []byte) error {
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
type SeradataOpticalPayloadGetResponseDataMode string

const (
	SeradataOpticalPayloadGetResponseDataModeReal      SeradataOpticalPayloadGetResponseDataMode = "REAL"
	SeradataOpticalPayloadGetResponseDataModeTest      SeradataOpticalPayloadGetResponseDataMode = "TEST"
	SeradataOpticalPayloadGetResponseDataModeSimulated SeradataOpticalPayloadGetResponseDataMode = "SIMULATED"
	SeradataOpticalPayloadGetResponseDataModeExercise  SeradataOpticalPayloadGetResponseDataMode = "EXERCISE"
)

type SeradataOpticalPayloadQueryhelpResponse struct {
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
func (r SeradataOpticalPayloadQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataOpticalPayloadQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for an optical payload from Seradata.
type SeradataOpticalPayloadTupleResponse struct {
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
	DataMode SeradataOpticalPayloadTupleResponseDataMode `json:"dataMode,required"`
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
func (r SeradataOpticalPayloadTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataOpticalPayloadTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeradataOpticalPayloadTupleResponseDataMode string

const (
	SeradataOpticalPayloadTupleResponseDataModeReal      SeradataOpticalPayloadTupleResponseDataMode = "REAL"
	SeradataOpticalPayloadTupleResponseDataModeTest      SeradataOpticalPayloadTupleResponseDataMode = "TEST"
	SeradataOpticalPayloadTupleResponseDataModeSimulated SeradataOpticalPayloadTupleResponseDataMode = "SIMULATED"
	SeradataOpticalPayloadTupleResponseDataModeExercise  SeradataOpticalPayloadTupleResponseDataMode = "EXERCISE"
)

type SeradataOpticalPayloadNewParams struct {
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
	DataMode SeradataOpticalPayloadNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Best resolution.
	BestResolution param.Opt[float64] `json:"bestResolution,omitzero"`
	// Field of regard of this radar in degrees.
	FieldOfRegard param.Opt[float64] `json:"fieldOfRegard,omitzero"`
	// Field of view of this radar in kilometers.
	FieldOfView param.Opt[float64] `json:"fieldOfView,omitzero"`
	// Ground Station Locations for this payload.
	GroundStationLocations param.Opt[string] `json:"groundStationLocations,omitzero"`
	// Ground Station info for this payload.
	GroundStations param.Opt[string] `json:"groundStations,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the Sensor record.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Imaging category for this payload, e.g. Multispectral, Infrared, Panchromatic.
	ImagingPayloadCategory param.Opt[string] `json:"imagingPayloadCategory,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Sensor name from Seradata, e.g. TOURNESOL, MESSR (Multispectral Self-Scanning
	// Radiometer), AWFI, etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// Payload notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of film return canisters.
	NumberOfFilmReturnCanisters param.Opt[int64] `json:"numberOfFilmReturnCanisters,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Point method for this radar, e.g. Spacecraft.
	PointingMethod param.Opt[string] `json:"pointingMethod,omitzero"`
	// Recorder size.
	RecorderSize param.Opt[string] `json:"recorderSize,omitzero"`
	// Spectral Band supported by this payload, e.g. Green, Red, Mid-wave infrared,
	// etc.
	SpectralBand param.Opt[string] `json:"spectralBand,omitzero"`
	// Frequency limit for this payload, e.g. 0.51 - 0.59.
	SpectralFrequencyLimits param.Opt[string] `json:"spectralFrequencyLimits,omitzero"`
	// Swath width in kilometers.
	SwathWidth param.Opt[float64] `json:"swathWidth,omitzero"`
	paramObj
}

func (r SeradataOpticalPayloadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataOpticalPayloadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeradataOpticalPayloadNewParams) UnmarshalJSON(data []byte) error {
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
type SeradataOpticalPayloadNewParamsDataMode string

const (
	SeradataOpticalPayloadNewParamsDataModeReal      SeradataOpticalPayloadNewParamsDataMode = "REAL"
	SeradataOpticalPayloadNewParamsDataModeTest      SeradataOpticalPayloadNewParamsDataMode = "TEST"
	SeradataOpticalPayloadNewParamsDataModeSimulated SeradataOpticalPayloadNewParamsDataMode = "SIMULATED"
	SeradataOpticalPayloadNewParamsDataModeExercise  SeradataOpticalPayloadNewParamsDataMode = "EXERCISE"
)

type SeradataOpticalPayloadUpdateParams struct {
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
	DataMode SeradataOpticalPayloadUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Best resolution.
	BestResolution param.Opt[float64] `json:"bestResolution,omitzero"`
	// Field of regard of this radar in degrees.
	FieldOfRegard param.Opt[float64] `json:"fieldOfRegard,omitzero"`
	// Field of view of this radar in kilometers.
	FieldOfView param.Opt[float64] `json:"fieldOfView,omitzero"`
	// Ground Station Locations for this payload.
	GroundStationLocations param.Opt[string] `json:"groundStationLocations,omitzero"`
	// Ground Station info for this payload.
	GroundStations param.Opt[string] `json:"groundStations,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the Sensor record.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Imaging category for this payload, e.g. Multispectral, Infrared, Panchromatic.
	ImagingPayloadCategory param.Opt[string] `json:"imagingPayloadCategory,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Sensor name from Seradata, e.g. TOURNESOL, MESSR (Multispectral Self-Scanning
	// Radiometer), AWFI, etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// Payload notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of film return canisters.
	NumberOfFilmReturnCanisters param.Opt[int64] `json:"numberOfFilmReturnCanisters,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Point method for this radar, e.g. Spacecraft.
	PointingMethod param.Opt[string] `json:"pointingMethod,omitzero"`
	// Recorder size.
	RecorderSize param.Opt[string] `json:"recorderSize,omitzero"`
	// Spectral Band supported by this payload, e.g. Green, Red, Mid-wave infrared,
	// etc.
	SpectralBand param.Opt[string] `json:"spectralBand,omitzero"`
	// Frequency limit for this payload, e.g. 0.51 - 0.59.
	SpectralFrequencyLimits param.Opt[string] `json:"spectralFrequencyLimits,omitzero"`
	// Swath width in kilometers.
	SwathWidth param.Opt[float64] `json:"swathWidth,omitzero"`
	paramObj
}

func (r SeradataOpticalPayloadUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataOpticalPayloadUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeradataOpticalPayloadUpdateParams) UnmarshalJSON(data []byte) error {
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
type SeradataOpticalPayloadUpdateParamsDataMode string

const (
	SeradataOpticalPayloadUpdateParamsDataModeReal      SeradataOpticalPayloadUpdateParamsDataMode = "REAL"
	SeradataOpticalPayloadUpdateParamsDataModeTest      SeradataOpticalPayloadUpdateParamsDataMode = "TEST"
	SeradataOpticalPayloadUpdateParamsDataModeSimulated SeradataOpticalPayloadUpdateParamsDataMode = "SIMULATED"
	SeradataOpticalPayloadUpdateParamsDataModeExercise  SeradataOpticalPayloadUpdateParamsDataMode = "EXERCISE"
)

type SeradataOpticalPayloadListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataOpticalPayloadListParams]'s query parameters as
// `url.Values`.
func (r SeradataOpticalPayloadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataOpticalPayloadCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataOpticalPayloadCountParams]'s query parameters as
// `url.Values`.
func (r SeradataOpticalPayloadCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataOpticalPayloadGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataOpticalPayloadGetParams]'s query parameters as
// `url.Values`.
func (r SeradataOpticalPayloadGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataOpticalPayloadTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataOpticalPayloadTupleParams]'s query parameters as
// `url.Values`.
func (r SeradataOpticalPayloadTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
