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

// SeradataradarpayloadService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeradataradarpayloadService] method instead.
type SeradataradarpayloadService struct {
	Options []option.RequestOption
}

// NewSeradataradarpayloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSeradataradarpayloadService(opts ...option.RequestOption) (r SeradataradarpayloadService) {
	r = SeradataradarpayloadService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataRadarPayload as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeradataradarpayloadService) New(ctx context.Context, body SeradataradarpayloadNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradataradarpayload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataRadarPayload. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *SeradataradarpayloadService) Update(ctx context.Context, id string, body SeradataradarpayloadUpdateParams, opts ...option.RequestOption) (err error) {
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
func (r *SeradataradarpayloadService) List(ctx context.Context, query SeradataradarpayloadListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeradataradarpayloadListResponse], err error) {
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
func (r *SeradataradarpayloadService) ListAutoPaging(ctx context.Context, query SeradataradarpayloadListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeradataradarpayloadListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataRadarPayload specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SeradataradarpayloadService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
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
func (r *SeradataradarpayloadService) Count(ctx context.Context, query SeradataradarpayloadCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradataradarpayload/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataRadarPayload by its unique ID passed
// as a path parameter.
func (r *SeradataradarpayloadService) Get(ctx context.Context, id string, query SeradataradarpayloadGetParams, opts ...option.RequestOption) (res *SeradataradarpayloadGetResponse, err error) {
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
func (r *SeradataradarpayloadService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
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
func (r *SeradataradarpayloadService) Tuple(ctx context.Context, query SeradataradarpayloadTupleParams, opts ...option.RequestOption) (res *[]SeradataradarpayloadTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/seradataradarpayload/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details for an radar payload from Seradata.
type SeradataradarpayloadListResponse struct {
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
	DataMode SeradataradarpayloadListResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                  resp.Field
		DataMode                               resp.Field
		Source                                 resp.Field
		SpacecraftID                           resp.Field
		ID                                     resp.Field
		Bandwidth                              resp.Field
		BestResolution                         resp.Field
		Category                               resp.Field
		ConstellationInterferometricCapability resp.Field
		CreatedAt                              resp.Field
		CreatedBy                              resp.Field
		DutyCycle                              resp.Field
		FieldOfRegard                          resp.Field
		FieldOfView                            resp.Field
		Frequency                              resp.Field
		FrequencyBand                          resp.Field
		GroundStationLocations                 resp.Field
		GroundStations                         resp.Field
		HostedForCompanyOrgID                  resp.Field
		IDSensor                               resp.Field
		ManufacturerOrgID                      resp.Field
		Name                                   resp.Field
		Notes                                  resp.Field
		Origin                                 resp.Field
		OrigNetwork                            resp.Field
		PartnerSpacecraft                      resp.Field
		PointingMethod                         resp.Field
		ReceivePolarization                    resp.Field
		RecorderSize                           resp.Field
		SwathWidth                             resp.Field
		TransmitPolarization                   resp.Field
		WaveLength                             resp.Field
		ExtraFields                            map[string]resp.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataradarpayloadListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataradarpayloadListResponse) UnmarshalJSON(data []byte) error {
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
type SeradataradarpayloadListResponseDataMode string

const (
	SeradataradarpayloadListResponseDataModeReal      SeradataradarpayloadListResponseDataMode = "REAL"
	SeradataradarpayloadListResponseDataModeTest      SeradataradarpayloadListResponseDataMode = "TEST"
	SeradataradarpayloadListResponseDataModeSimulated SeradataradarpayloadListResponseDataMode = "SIMULATED"
	SeradataradarpayloadListResponseDataModeExercise  SeradataradarpayloadListResponseDataMode = "EXERCISE"
)

// Details for an radar payload from Seradata.
type SeradataradarpayloadGetResponse struct {
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
	DataMode SeradataradarpayloadGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                  resp.Field
		DataMode                               resp.Field
		Source                                 resp.Field
		SpacecraftID                           resp.Field
		ID                                     resp.Field
		Bandwidth                              resp.Field
		BestResolution                         resp.Field
		Category                               resp.Field
		ConstellationInterferometricCapability resp.Field
		CreatedAt                              resp.Field
		CreatedBy                              resp.Field
		DutyCycle                              resp.Field
		FieldOfRegard                          resp.Field
		FieldOfView                            resp.Field
		Frequency                              resp.Field
		FrequencyBand                          resp.Field
		GroundStationLocations                 resp.Field
		GroundStations                         resp.Field
		HostedForCompanyOrgID                  resp.Field
		IDSensor                               resp.Field
		ManufacturerOrgID                      resp.Field
		Name                                   resp.Field
		Notes                                  resp.Field
		Origin                                 resp.Field
		OrigNetwork                            resp.Field
		PartnerSpacecraft                      resp.Field
		PointingMethod                         resp.Field
		ReceivePolarization                    resp.Field
		RecorderSize                           resp.Field
		SwathWidth                             resp.Field
		TransmitPolarization                   resp.Field
		UpdatedAt                              resp.Field
		UpdatedBy                              resp.Field
		WaveLength                             resp.Field
		ExtraFields                            map[string]resp.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataradarpayloadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataradarpayloadGetResponse) UnmarshalJSON(data []byte) error {
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
type SeradataradarpayloadGetResponseDataMode string

const (
	SeradataradarpayloadGetResponseDataModeReal      SeradataradarpayloadGetResponseDataMode = "REAL"
	SeradataradarpayloadGetResponseDataModeTest      SeradataradarpayloadGetResponseDataMode = "TEST"
	SeradataradarpayloadGetResponseDataModeSimulated SeradataradarpayloadGetResponseDataMode = "SIMULATED"
	SeradataradarpayloadGetResponseDataModeExercise  SeradataradarpayloadGetResponseDataMode = "EXERCISE"
)

// Details for an radar payload from Seradata.
type SeradataradarpayloadTupleResponse struct {
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
	DataMode SeradataradarpayloadTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                  resp.Field
		DataMode                               resp.Field
		Source                                 resp.Field
		SpacecraftID                           resp.Field
		ID                                     resp.Field
		Bandwidth                              resp.Field
		BestResolution                         resp.Field
		Category                               resp.Field
		ConstellationInterferometricCapability resp.Field
		CreatedAt                              resp.Field
		CreatedBy                              resp.Field
		DutyCycle                              resp.Field
		FieldOfRegard                          resp.Field
		FieldOfView                            resp.Field
		Frequency                              resp.Field
		FrequencyBand                          resp.Field
		GroundStationLocations                 resp.Field
		GroundStations                         resp.Field
		HostedForCompanyOrgID                  resp.Field
		IDSensor                               resp.Field
		ManufacturerOrgID                      resp.Field
		Name                                   resp.Field
		Notes                                  resp.Field
		Origin                                 resp.Field
		OrigNetwork                            resp.Field
		PartnerSpacecraft                      resp.Field
		PointingMethod                         resp.Field
		ReceivePolarization                    resp.Field
		RecorderSize                           resp.Field
		SwathWidth                             resp.Field
		TransmitPolarization                   resp.Field
		UpdatedAt                              resp.Field
		UpdatedBy                              resp.Field
		WaveLength                             resp.Field
		ExtraFields                            map[string]resp.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataradarpayloadTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataradarpayloadTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeradataradarpayloadTupleResponseDataMode string

const (
	SeradataradarpayloadTupleResponseDataModeReal      SeradataradarpayloadTupleResponseDataMode = "REAL"
	SeradataradarpayloadTupleResponseDataModeTest      SeradataradarpayloadTupleResponseDataMode = "TEST"
	SeradataradarpayloadTupleResponseDataModeSimulated SeradataradarpayloadTupleResponseDataMode = "SIMULATED"
	SeradataradarpayloadTupleResponseDataModeExercise  SeradataradarpayloadTupleResponseDataMode = "EXERCISE"
)

type SeradataradarpayloadNewParams struct {
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
	DataMode SeradataradarpayloadNewParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataradarpayloadNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SeradataradarpayloadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataradarpayloadNewParams
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
type SeradataradarpayloadNewParamsDataMode string

const (
	SeradataradarpayloadNewParamsDataModeReal      SeradataradarpayloadNewParamsDataMode = "REAL"
	SeradataradarpayloadNewParamsDataModeTest      SeradataradarpayloadNewParamsDataMode = "TEST"
	SeradataradarpayloadNewParamsDataModeSimulated SeradataradarpayloadNewParamsDataMode = "SIMULATED"
	SeradataradarpayloadNewParamsDataModeExercise  SeradataradarpayloadNewParamsDataMode = "EXERCISE"
)

type SeradataradarpayloadUpdateParams struct {
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
	DataMode SeradataradarpayloadUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataradarpayloadUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SeradataradarpayloadUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataradarpayloadUpdateParams
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
type SeradataradarpayloadUpdateParamsDataMode string

const (
	SeradataradarpayloadUpdateParamsDataModeReal      SeradataradarpayloadUpdateParamsDataMode = "REAL"
	SeradataradarpayloadUpdateParamsDataModeTest      SeradataradarpayloadUpdateParamsDataMode = "TEST"
	SeradataradarpayloadUpdateParamsDataModeSimulated SeradataradarpayloadUpdateParamsDataMode = "SIMULATED"
	SeradataradarpayloadUpdateParamsDataModeExercise  SeradataradarpayloadUpdateParamsDataMode = "EXERCISE"
)

type SeradataradarpayloadListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataradarpayloadListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradataradarpayloadListParams]'s query parameters as
// `url.Values`.
func (r SeradataradarpayloadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataradarpayloadCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataradarpayloadCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradataradarpayloadCountParams]'s query parameters as
// `url.Values`.
func (r SeradataradarpayloadCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataradarpayloadGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradataradarpayloadGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradataradarpayloadGetParams]'s query parameters as
// `url.Values`.
func (r SeradataradarpayloadGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataradarpayloadTupleParams struct {
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
func (f SeradataradarpayloadTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradataradarpayloadTupleParams]'s query parameters as
// `url.Values`.
func (r SeradataradarpayloadTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
