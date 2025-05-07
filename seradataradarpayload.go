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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// SeradataRadarPayloadService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeradataRadarPayloadService] method instead.
type SeradataRadarPayloadService struct {
	Options []option.RequestOption
}

// NewSeradataRadarPayloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSeradataRadarPayloadService(opts ...option.RequestOption) (r SeradataRadarPayloadService) {
	r = SeradataRadarPayloadService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataRadarPayload as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeradataRadarPayloadService) New(ctx context.Context, body SeradataRadarPayloadNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradataradarpayload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataRadarPayload. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *SeradataRadarPayloadService) Update(ctx context.Context, id string, body SeradataRadarPayloadUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataradarpayload/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SeradataRadarPayloadService) List(ctx context.Context, query SeradataRadarPayloadListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeradataRadarPayloadListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/seradataradarpayload"
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
func (r *SeradataRadarPayloadService) ListAutoPaging(ctx context.Context, query SeradataRadarPayloadListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeradataRadarPayloadListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataRadarPayload specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SeradataRadarPayloadService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataradarpayload/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SeradataRadarPayloadService) Count(ctx context.Context, query SeradataRadarPayloadCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradataradarpayload/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataRadarPayload by its unique ID passed
// as a path parameter.
func (r *SeradataRadarPayloadService) Get(ctx context.Context, id string, query SeradataRadarPayloadGetParams, opts ...option.RequestOption) (res *SeradataRadarPayloadGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradataradarpayload/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SeradataRadarPayloadService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradataradarpayload/queryhelp"
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
func (r *SeradataRadarPayloadService) Tuple(ctx context.Context, query SeradataRadarPayloadTupleParams, opts ...option.RequestOption) (res *[]SeradataRadarPayloadTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/seradataradarpayload/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details for an radar payload from Seradata.
type SeradataRadarPayloadListResponse struct {
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
	DataMode SeradataRadarPayloadListResponseDataMode `json:"dataMode,required"`
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
		WaveLength                             respjson.Field
		ExtraFields                            map[string]respjson.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataRadarPayloadListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataRadarPayloadListResponse) UnmarshalJSON(data []byte) error {
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
type SeradataRadarPayloadListResponseDataMode string

const (
	SeradataRadarPayloadListResponseDataModeReal      SeradataRadarPayloadListResponseDataMode = "REAL"
	SeradataRadarPayloadListResponseDataModeTest      SeradataRadarPayloadListResponseDataMode = "TEST"
	SeradataRadarPayloadListResponseDataModeSimulated SeradataRadarPayloadListResponseDataMode = "SIMULATED"
	SeradataRadarPayloadListResponseDataModeExercise  SeradataRadarPayloadListResponseDataMode = "EXERCISE"
)

// Details for an radar payload from Seradata.
type SeradataRadarPayloadGetResponse struct {
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
	DataMode SeradataRadarPayloadGetResponseDataMode `json:"dataMode,required"`
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
func (r SeradataRadarPayloadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataRadarPayloadGetResponse) UnmarshalJSON(data []byte) error {
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
type SeradataRadarPayloadGetResponseDataMode string

const (
	SeradataRadarPayloadGetResponseDataModeReal      SeradataRadarPayloadGetResponseDataMode = "REAL"
	SeradataRadarPayloadGetResponseDataModeTest      SeradataRadarPayloadGetResponseDataMode = "TEST"
	SeradataRadarPayloadGetResponseDataModeSimulated SeradataRadarPayloadGetResponseDataMode = "SIMULATED"
	SeradataRadarPayloadGetResponseDataModeExercise  SeradataRadarPayloadGetResponseDataMode = "EXERCISE"
)

// Details for an radar payload from Seradata.
type SeradataRadarPayloadTupleResponse struct {
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
	DataMode SeradataRadarPayloadTupleResponseDataMode `json:"dataMode,required"`
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
func (r SeradataRadarPayloadTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataRadarPayloadTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeradataRadarPayloadTupleResponseDataMode string

const (
	SeradataRadarPayloadTupleResponseDataModeReal      SeradataRadarPayloadTupleResponseDataMode = "REAL"
	SeradataRadarPayloadTupleResponseDataModeTest      SeradataRadarPayloadTupleResponseDataMode = "TEST"
	SeradataRadarPayloadTupleResponseDataModeSimulated SeradataRadarPayloadTupleResponseDataMode = "SIMULATED"
	SeradataRadarPayloadTupleResponseDataModeExercise  SeradataRadarPayloadTupleResponseDataMode = "EXERCISE"
)

type SeradataRadarPayloadNewParams struct {
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
	DataMode SeradataRadarPayloadNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Radar bandwidth in mega hertz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Best resolution in meters.
	BestResolution param.Opt[float64] `json:"bestResolution,omitzero"`
	// Radar category, e.g. SAR, Surface Search, etc.
	Category param.Opt[string] `json:"category,omitzero"`
	// Constellation interferometric capability.
	ConstellationInterferometricCapability param.Opt[string] `json:"constellationInterferometricCapability,omitzero"`
	// Duty cycle.
	DutyCycle param.Opt[string] `json:"dutyCycle,omitzero"`
	// Field of regard of this radar in degrees.
	FieldOfRegard param.Opt[float64] `json:"fieldOfRegard,omitzero"`
	// Field of view of this radar in kilometers.
	FieldOfView param.Opt[float64] `json:"fieldOfView,omitzero"`
	// Frequency in giga hertz.
	Frequency param.Opt[float64] `json:"frequency,omitzero"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	FrequencyBand param.Opt[string] `json:"frequencyBand,omitzero"`
	// Ground Station Locations for this payload.
	GroundStationLocations param.Opt[string] `json:"groundStationLocations,omitzero"`
	// Ground Station info for this payload.
	GroundStations param.Opt[string] `json:"groundStations,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the Sensor record.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Sensor name from Seradata, e.g. ALT (Radar Altimeter), COSI (Corea SAR
	// Instrument), etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// Payload notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Partner seradata-spacecraft.
	PartnerSpacecraft param.Opt[string] `json:"partnerSpacecraft,omitzero"`
	// Point method for this radar, e.g. Spacecraft.
	PointingMethod param.Opt[string] `json:"pointingMethod,omitzero"`
	// Receive polarization, e.g. Lin Dual, Lin vert, etc.
	ReceivePolarization param.Opt[string] `json:"receivePolarization,omitzero"`
	// Recorder size, e.g. 256.
	RecorderSize param.Opt[string] `json:"recorderSize,omitzero"`
	// Swath width in kilometers.
	SwathWidth param.Opt[float64] `json:"swathWidth,omitzero"`
	// Transmit polarization, e.g. Lin Dual, Lin vert, etc.
	TransmitPolarization param.Opt[string] `json:"transmitPolarization,omitzero"`
	// Wave length in meters.
	WaveLength param.Opt[float64] `json:"waveLength,omitzero"`
	paramObj
}

func (r SeradataRadarPayloadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataRadarPayloadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeradataRadarPayloadNewParams) UnmarshalJSON(data []byte) error {
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
type SeradataRadarPayloadNewParamsDataMode string

const (
	SeradataRadarPayloadNewParamsDataModeReal      SeradataRadarPayloadNewParamsDataMode = "REAL"
	SeradataRadarPayloadNewParamsDataModeTest      SeradataRadarPayloadNewParamsDataMode = "TEST"
	SeradataRadarPayloadNewParamsDataModeSimulated SeradataRadarPayloadNewParamsDataMode = "SIMULATED"
	SeradataRadarPayloadNewParamsDataModeExercise  SeradataRadarPayloadNewParamsDataMode = "EXERCISE"
)

type SeradataRadarPayloadUpdateParams struct {
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
	DataMode SeradataRadarPayloadUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Radar bandwidth in mega hertz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Best resolution in meters.
	BestResolution param.Opt[float64] `json:"bestResolution,omitzero"`
	// Radar category, e.g. SAR, Surface Search, etc.
	Category param.Opt[string] `json:"category,omitzero"`
	// Constellation interferometric capability.
	ConstellationInterferometricCapability param.Opt[string] `json:"constellationInterferometricCapability,omitzero"`
	// Duty cycle.
	DutyCycle param.Opt[string] `json:"dutyCycle,omitzero"`
	// Field of regard of this radar in degrees.
	FieldOfRegard param.Opt[float64] `json:"fieldOfRegard,omitzero"`
	// Field of view of this radar in kilometers.
	FieldOfView param.Opt[float64] `json:"fieldOfView,omitzero"`
	// Frequency in giga hertz.
	Frequency param.Opt[float64] `json:"frequency,omitzero"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	FrequencyBand param.Opt[string] `json:"frequencyBand,omitzero"`
	// Ground Station Locations for this payload.
	GroundStationLocations param.Opt[string] `json:"groundStationLocations,omitzero"`
	// Ground Station info for this payload.
	GroundStations param.Opt[string] `json:"groundStations,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the Sensor record.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Sensor name from Seradata, e.g. ALT (Radar Altimeter), COSI (Corea SAR
	// Instrument), etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// Payload notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Partner seradata-spacecraft.
	PartnerSpacecraft param.Opt[string] `json:"partnerSpacecraft,omitzero"`
	// Point method for this radar, e.g. Spacecraft.
	PointingMethod param.Opt[string] `json:"pointingMethod,omitzero"`
	// Receive polarization, e.g. Lin Dual, Lin vert, etc.
	ReceivePolarization param.Opt[string] `json:"receivePolarization,omitzero"`
	// Recorder size, e.g. 256.
	RecorderSize param.Opt[string] `json:"recorderSize,omitzero"`
	// Swath width in kilometers.
	SwathWidth param.Opt[float64] `json:"swathWidth,omitzero"`
	// Transmit polarization, e.g. Lin Dual, Lin vert, etc.
	TransmitPolarization param.Opt[string] `json:"transmitPolarization,omitzero"`
	// Wave length in meters.
	WaveLength param.Opt[float64] `json:"waveLength,omitzero"`
	paramObj
}

func (r SeradataRadarPayloadUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataRadarPayloadUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeradataRadarPayloadUpdateParams) UnmarshalJSON(data []byte) error {
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
type SeradataRadarPayloadUpdateParamsDataMode string

const (
	SeradataRadarPayloadUpdateParamsDataModeReal      SeradataRadarPayloadUpdateParamsDataMode = "REAL"
	SeradataRadarPayloadUpdateParamsDataModeTest      SeradataRadarPayloadUpdateParamsDataMode = "TEST"
	SeradataRadarPayloadUpdateParamsDataModeSimulated SeradataRadarPayloadUpdateParamsDataMode = "SIMULATED"
	SeradataRadarPayloadUpdateParamsDataModeExercise  SeradataRadarPayloadUpdateParamsDataMode = "EXERCISE"
)

type SeradataRadarPayloadListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataRadarPayloadListParams]'s query parameters as
// `url.Values`.
func (r SeradataRadarPayloadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataRadarPayloadCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataRadarPayloadCountParams]'s query parameters as
// `url.Values`.
func (r SeradataRadarPayloadCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataRadarPayloadGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataRadarPayloadGetParams]'s query parameters as
// `url.Values`.
func (r SeradataRadarPayloadGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataRadarPayloadTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataRadarPayloadTupleParams]'s query parameters as
// `url.Values`.
func (r SeradataRadarPayloadTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
