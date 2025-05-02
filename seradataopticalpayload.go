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

// SeradataopticalpayloadService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeradataopticalpayloadService] method instead.
type SeradataopticalpayloadService struct {
	Options []option.RequestOption
}

// NewSeradataopticalpayloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSeradataopticalpayloadService(opts ...option.RequestOption) (r SeradataopticalpayloadService) {
	r = SeradataopticalpayloadService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataOpticalPayload as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeradataopticalpayloadService) New(ctx context.Context, body SeradataopticalpayloadNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradataopticalpayload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataOpticalPayload. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SeradataopticalpayloadService) Update(ctx context.Context, id string, body SeradataopticalpayloadUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
func (r *SeradataopticalpayloadService) List(ctx context.Context, query SeradataopticalpayloadListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeradataopticalpayloadListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *SeradataopticalpayloadService) ListAutoPaging(ctx context.Context, query SeradataopticalpayloadListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeradataopticalpayloadListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataOpticalPayload specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SeradataopticalpayloadService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
func (r *SeradataopticalpayloadService) Count(ctx context.Context, query SeradataopticalpayloadCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradataopticalpayload/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataOpticalPayload by its unique ID passed
// as a path parameter.
func (r *SeradataopticalpayloadService) Get(ctx context.Context, id string, query SeradataopticalpayloadGetParams, opts ...option.RequestOption) (res *SeradataopticalpayloadGetResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *SeradataopticalpayloadService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradataopticalpayload/queryhelp"
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
func (r *SeradataopticalpayloadService) Tuple(ctx context.Context, query SeradataopticalpayloadTupleParams, opts ...option.RequestOption) (res *[]SeradataopticalpayloadTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/seradataopticalpayload/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details for an optical payload from Seradata.
type SeradataopticalpayloadListResponse struct {
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
	DataMode SeradataopticalpayloadListResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking       resp.Field
		DataMode                    resp.Field
		Source                      resp.Field
		SpacecraftID                resp.Field
		ID                          resp.Field
		BestResolution              resp.Field
		CreatedAt                   resp.Field
		CreatedBy                   resp.Field
		FieldOfRegard               resp.Field
		FieldOfView                 resp.Field
		GroundStationLocations      resp.Field
		GroundStations              resp.Field
		HostedForCompanyOrgID       resp.Field
		IDSensor                    resp.Field
		ImagingPayloadCategory      resp.Field
		ManufacturerOrgID           resp.Field
		Name                        resp.Field
		Notes                       resp.Field
		NumberOfFilmReturnCanisters resp.Field
		Origin                      resp.Field
		OrigNetwork                 resp.Field
		PointingMethod              resp.Field
		RecorderSize                resp.Field
		SpectralBand                resp.Field
		SpectralFrequencyLimits     resp.Field
		SwathWidth                  resp.Field
		ExtraFields                 map[string]resp.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataopticalpayloadListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataopticalpayloadListResponse) UnmarshalJSON(data []byte) error {
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
type SeradataopticalpayloadListResponseDataMode string

const (
	SeradataopticalpayloadListResponseDataModeReal      SeradataopticalpayloadListResponseDataMode = "REAL"
	SeradataopticalpayloadListResponseDataModeTest      SeradataopticalpayloadListResponseDataMode = "TEST"
	SeradataopticalpayloadListResponseDataModeSimulated SeradataopticalpayloadListResponseDataMode = "SIMULATED"
	SeradataopticalpayloadListResponseDataModeExercise  SeradataopticalpayloadListResponseDataMode = "EXERCISE"
)

// Details for an optical payload from Seradata.
type SeradataopticalpayloadGetResponse struct {
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
	DataMode SeradataopticalpayloadGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking       resp.Field
		DataMode                    resp.Field
		Source                      resp.Field
		SpacecraftID                resp.Field
		ID                          resp.Field
		BestResolution              resp.Field
		CreatedAt                   resp.Field
		CreatedBy                   resp.Field
		FieldOfRegard               resp.Field
		FieldOfView                 resp.Field
		GroundStationLocations      resp.Field
		GroundStations              resp.Field
		HostedForCompanyOrgID       resp.Field
		IDSensor                    resp.Field
		ImagingPayloadCategory      resp.Field
		ManufacturerOrgID           resp.Field
		Name                        resp.Field
		Notes                       resp.Field
		NumberOfFilmReturnCanisters resp.Field
		Origin                      resp.Field
		OrigNetwork                 resp.Field
		PointingMethod              resp.Field
		RecorderSize                resp.Field
		SpectralBand                resp.Field
		SpectralFrequencyLimits     resp.Field
		SwathWidth                  resp.Field
		UpdatedAt                   resp.Field
		UpdatedBy                   resp.Field
		ExtraFields                 map[string]resp.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataopticalpayloadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataopticalpayloadGetResponse) UnmarshalJSON(data []byte) error {
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
type SeradataopticalpayloadGetResponseDataMode string

const (
	SeradataopticalpayloadGetResponseDataModeReal      SeradataopticalpayloadGetResponseDataMode = "REAL"
	SeradataopticalpayloadGetResponseDataModeTest      SeradataopticalpayloadGetResponseDataMode = "TEST"
	SeradataopticalpayloadGetResponseDataModeSimulated SeradataopticalpayloadGetResponseDataMode = "SIMULATED"
	SeradataopticalpayloadGetResponseDataModeExercise  SeradataopticalpayloadGetResponseDataMode = "EXERCISE"
)

// Details for an optical payload from Seradata.
type SeradataopticalpayloadTupleResponse struct {
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
	DataMode SeradataopticalpayloadTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking       resp.Field
		DataMode                    resp.Field
		Source                      resp.Field
		SpacecraftID                resp.Field
		ID                          resp.Field
		BestResolution              resp.Field
		CreatedAt                   resp.Field
		CreatedBy                   resp.Field
		FieldOfRegard               resp.Field
		FieldOfView                 resp.Field
		GroundStationLocations      resp.Field
		GroundStations              resp.Field
		HostedForCompanyOrgID       resp.Field
		IDSensor                    resp.Field
		ImagingPayloadCategory      resp.Field
		ManufacturerOrgID           resp.Field
		Name                        resp.Field
		Notes                       resp.Field
		NumberOfFilmReturnCanisters resp.Field
		Origin                      resp.Field
		OrigNetwork                 resp.Field
		PointingMethod              resp.Field
		RecorderSize                resp.Field
		SpectralBand                resp.Field
		SpectralFrequencyLimits     resp.Field
		SwathWidth                  resp.Field
		UpdatedAt                   resp.Field
		UpdatedBy                   resp.Field
		ExtraFields                 map[string]resp.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataopticalpayloadTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataopticalpayloadTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeradataopticalpayloadTupleResponseDataMode string

const (
	SeradataopticalpayloadTupleResponseDataModeReal      SeradataopticalpayloadTupleResponseDataMode = "REAL"
	SeradataopticalpayloadTupleResponseDataModeTest      SeradataopticalpayloadTupleResponseDataMode = "TEST"
	SeradataopticalpayloadTupleResponseDataModeSimulated SeradataopticalpayloadTupleResponseDataMode = "SIMULATED"
	SeradataopticalpayloadTupleResponseDataModeExercise  SeradataopticalpayloadTupleResponseDataMode = "EXERCISE"
)

type SeradataopticalpayloadNewParams struct {
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
	DataMode SeradataopticalpayloadNewParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataopticalpayloadNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SeradataopticalpayloadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataopticalpayloadNewParams
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
type SeradataopticalpayloadNewParamsDataMode string

const (
	SeradataopticalpayloadNewParamsDataModeReal      SeradataopticalpayloadNewParamsDataMode = "REAL"
	SeradataopticalpayloadNewParamsDataModeTest      SeradataopticalpayloadNewParamsDataMode = "TEST"
	SeradataopticalpayloadNewParamsDataModeSimulated SeradataopticalpayloadNewParamsDataMode = "SIMULATED"
	SeradataopticalpayloadNewParamsDataModeExercise  SeradataopticalpayloadNewParamsDataMode = "EXERCISE"
)

type SeradataopticalpayloadUpdateParams struct {
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
	DataMode SeradataopticalpayloadUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataopticalpayloadUpdateParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r SeradataopticalpayloadUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataopticalpayloadUpdateParams
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
type SeradataopticalpayloadUpdateParamsDataMode string

const (
	SeradataopticalpayloadUpdateParamsDataModeReal      SeradataopticalpayloadUpdateParamsDataMode = "REAL"
	SeradataopticalpayloadUpdateParamsDataModeTest      SeradataopticalpayloadUpdateParamsDataMode = "TEST"
	SeradataopticalpayloadUpdateParamsDataModeSimulated SeradataopticalpayloadUpdateParamsDataMode = "SIMULATED"
	SeradataopticalpayloadUpdateParamsDataModeExercise  SeradataopticalpayloadUpdateParamsDataMode = "EXERCISE"
)

type SeradataopticalpayloadListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataopticalpayloadListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradataopticalpayloadListParams]'s query parameters as
// `url.Values`.
func (r SeradataopticalpayloadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataopticalpayloadCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataopticalpayloadCountParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [SeradataopticalpayloadCountParams]'s query parameters as
// `url.Values`.
func (r SeradataopticalpayloadCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataopticalpayloadGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataopticalpayloadGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradataopticalpayloadGetParams]'s query parameters as
// `url.Values`.
func (r SeradataopticalpayloadGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataopticalpayloadTupleParams struct {
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
func (f SeradataopticalpayloadTupleParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [SeradataopticalpayloadTupleParams]'s query parameters as
// `url.Values`.
func (r SeradataopticalpayloadTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
