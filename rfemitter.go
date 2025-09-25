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

// RfEmitterService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRfEmitterService] method instead.
type RfEmitterService struct {
	Options []option.RequestOption
	Staging RfEmitterStagingService
	Details RfEmitterDetailService
}

// NewRfEmitterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRfEmitterService(opts ...option.RequestOption) (r RfEmitterService) {
	r = RfEmitterService{}
	r.Options = opts
	r.Staging = NewRfEmitterStagingService(opts...)
	r.Details = NewRfEmitterDetailService(opts...)
	return
}

// Service operation to take a single RFEmitter as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RfEmitterService) New(ctx context.Context, body RfEmitterNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfemitter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single RFEmitter record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *RfEmitterService) Update(ctx context.Context, id string, body RfEmitterUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfemitter/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *RfEmitterService) List(ctx context.Context, query RfEmitterListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[RfEmitterListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/rfemitter"
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
func (r *RfEmitterService) ListAutoPaging(ctx context.Context, query RfEmitterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[RfEmitterListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a RFEmitter record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RfEmitterService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfemitter/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *RfEmitterService) Count(ctx context.Context, query RfEmitterCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/rfemitter/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single RFEmitter record by its unique ID passed as a
// path parameter.
func (r *RfEmitterService) Get(ctx context.Context, id string, query RfEmitterGetParams, opts ...option.RequestOption) (res *RfEmitterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfemitter/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *RfEmitterService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *RfEmitterQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/rfemitter/queryhelp"
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
func (r *RfEmitterService) Tuple(ctx context.Context, query RfEmitterTupleParams, opts ...option.RequestOption) (res *[]RfEmitterTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/rfemitter/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An RF Emitter is a source of active Radio Frequency (RF) signals which could
// potentially interfere with other RF receivers.
type RfEmitterListResponse struct {
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
	DataMode RfEmitterListResponseDataMode `json:"dataMode,required"`
	// Unique name of this RF Emitter.
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
	// The originating system ID for the RF Emitter.
	ExtSysID string `json:"extSysId"`
	// ID by reference of the parent entity for this RFEmitter.
	IDEntity string `json:"idEntity"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype string `json:"subtype"`
	// Type of this RF Emitter.
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
		ExtSysID              respjson.Field
		IDEntity              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Subtype               respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterListResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterListResponse) UnmarshalJSON(data []byte) error {
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
type RfEmitterListResponseDataMode string

const (
	RfEmitterListResponseDataModeReal      RfEmitterListResponseDataMode = "REAL"
	RfEmitterListResponseDataModeTest      RfEmitterListResponseDataMode = "TEST"
	RfEmitterListResponseDataModeSimulated RfEmitterListResponseDataMode = "SIMULATED"
	RfEmitterListResponseDataModeExercise  RfEmitterListResponseDataMode = "EXERCISE"
)

// An RF Emitter is a source of active Radio Frequency (RF) signals which could
// potentially interfere with other RF receivers.
type RfEmitterGetResponse struct {
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
	DataMode RfEmitterGetResponseDataMode `json:"dataMode,required"`
	// Unique name of this RF Emitter.
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
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity shared.EntityFull `json:"entity"`
	// The originating system ID for the RF Emitter.
	ExtSysID string `json:"extSysId"`
	// ID by reference of the parent entity for this RFEmitter.
	IDEntity string `json:"idEntity"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Details about this RF Emitter.
	RfEmitterDetails []RfEmitterGetResponseRfEmitterDetail `json:"rfEmitterDetails"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype string `json:"subtype"`
	// Type of this RF Emitter.
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
		Entity                respjson.Field
		ExtSysID              respjson.Field
		IDEntity              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RfEmitterDetails      respjson.Field
		Subtype               respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterGetResponse) UnmarshalJSON(data []byte) error {
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
type RfEmitterGetResponseDataMode string

const (
	RfEmitterGetResponseDataModeReal      RfEmitterGetResponseDataMode = "REAL"
	RfEmitterGetResponseDataModeTest      RfEmitterGetResponseDataMode = "TEST"
	RfEmitterGetResponseDataModeSimulated RfEmitterGetResponseDataMode = "SIMULATED"
	RfEmitterGetResponseDataModeExercise  RfEmitterGetResponseDataMode = "EXERCISE"
)

// Details for a particular RF Emitter, collected by a particular source. An RF
// Emitter may have multiple details records from various sources.
type RfEmitterGetResponseRfEmitterDetail struct {
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
	// Unique identifier of the parent RF Emitter.
	IDRfEmitter string `json:"idRFEmitter,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate facility name for this RF Emitter.
	AlternateFacilityName string `json:"alternateFacilityName"`
	// Optional alternate name or alias for this RF Emitter.
	AltName string `json:"altName"`
	// An RF Amplifier associated with an RF Emitter Details.
	Amplifier RfEmitterGetResponseRfEmitterDetailAmplifier `json:"amplifier"`
	// The set of antennas hosted on this EW Emitter system.
	Antennas []RfEmitterGetResponseRfEmitterDetailAntenna `json:"antennas"`
	// Barrage noise bandwidth, in megahertz.
	BarrageNoiseBandwidth float64 `json:"barrageNoiseBandwidth"`
	// The length of time, in seconds, for the RF Emitter built-in test to run to
	// completion.
	BitRunTime float64 `json:"bitRunTime"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Detailed description of the RF Emitter.
	Description string `json:"description"`
	// Designator of this RF Emitter.
	Designator string `json:"designator"`
	// Doppler noise value, in megahertz.
	DopplerNoise float64 `json:"dopplerNoise"`
	// Digital Form Radio Memory instantaneous bandwidth in megahertz.
	DrfmInstantaneousBandwidth float64 `json:"drfmInstantaneousBandwidth"`
	// Family of this RF Emitter type.
	Family string `json:"family"`
	// A fixed attenuation value to be set on the SRF Emitter HPA when commanding an
	// Electronic Attack/Techniques Tactics and Procedures task, in decibels.
	FixedAttenuation float64 `json:"fixedAttenuation"`
	// Unique identifier of the organization which manufactured this RF Emitter.
	IDManufacturerOrg string `json:"idManufacturerOrg"`
	// Unique identifier of the location of the production facility for this RF
	// Emitter.
	IDProductionFacilityLocation string `json:"idProductionFacilityLocation"`
	// COCOM that has temporary responsibility for scheduling and management of the RF
	// Emitter (e.g. SPACEFOR-CENT, SPACEFOR-EURAF, SPACEFOR-INDOPAC, SPACEFOR-KOR,
	// SPACEFOR-STRATNORTH, SPACESOC, NONE).
	LoanedToCocom string `json:"loanedToCocom"`
	// Notes on the RF Emitter.
	Notes string `json:"notes"`
	// Number of bits.
	NumBits int64 `json:"numBits"`
	// Number of channels.
	NumChannels int64 `json:"numChannels"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// A set of system/frequency band adjustments to the power offset commanded in an
	// EA/TTP task.
	PowerOffsets []RfEmitterGetResponseRfEmitterDetailPowerOffset `json:"powerOffsets"`
	// The length of time, in seconds, for the RF Emitter to prepare for a task,
	// including sufficient time to slew the antenna and configure the equipment.
	PrepTime float64 `json:"prepTime"`
	// Primary COCOM that is responsible for scheduling and management of the RF
	// Emitter (e.g. SPACEFOR-CENT, SPACEFOR-EURAF, SPACEFOR-INDOPAC, SPACEFOR-KOR,
	// SPACEFOR-STRATNORTH, SPACESOC, NONE).
	PrimaryCocom string `json:"primaryCocom"`
	// Name of the production facility for this RF Emitter.
	ProductionFacilityName string `json:"productionFacilityName"`
	// Type or name of receiver.
	ReceiverType string `json:"receiverType"`
	// Secondary notes on the RF Emitter.
	SecondaryNotes string `json:"secondaryNotes"`
	// The set of software services running on this EW Emitter system.
	Services []RfEmitterGetResponseRfEmitterDetailService `json:"services"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. End sensitivity range, in
	// decibel-milliwatts.
	SystemSensitivityEnd float64 `json:"systemSensitivityEnd"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. Start sensitivity range, in
	// decibel-milliwatts.
	SystemSensitivityStart float64 `json:"systemSensitivityStart"`
	// The set of EA/TTP techniques that are supported by this EW Emitter system.
	Ttps []RfEmitterGetResponseRfEmitterDetailTtp `json:"ttps"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Array of URLs containing additional information on this RF Emitter.
	URLs []string `json:"urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking        respjson.Field
		DataMode                     respjson.Field
		IDRfEmitter                  respjson.Field
		Source                       respjson.Field
		ID                           respjson.Field
		AlternateFacilityName        respjson.Field
		AltName                      respjson.Field
		Amplifier                    respjson.Field
		Antennas                     respjson.Field
		BarrageNoiseBandwidth        respjson.Field
		BitRunTime                   respjson.Field
		CreatedAt                    respjson.Field
		CreatedBy                    respjson.Field
		Description                  respjson.Field
		Designator                   respjson.Field
		DopplerNoise                 respjson.Field
		DrfmInstantaneousBandwidth   respjson.Field
		Family                       respjson.Field
		FixedAttenuation             respjson.Field
		IDManufacturerOrg            respjson.Field
		IDProductionFacilityLocation respjson.Field
		LoanedToCocom                respjson.Field
		Notes                        respjson.Field
		NumBits                      respjson.Field
		NumChannels                  respjson.Field
		Origin                       respjson.Field
		OrigNetwork                  respjson.Field
		PowerOffsets                 respjson.Field
		PrepTime                     respjson.Field
		PrimaryCocom                 respjson.Field
		ProductionFacilityName       respjson.Field
		ReceiverType                 respjson.Field
		SecondaryNotes               respjson.Field
		Services                     respjson.Field
		SystemSensitivityEnd         respjson.Field
		SystemSensitivityStart       respjson.Field
		Ttps                         respjson.Field
		UpdatedAt                    respjson.Field
		UpdatedBy                    respjson.Field
		URLs                         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetail) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterGetResponseRfEmitterDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Amplifier associated with an RF Emitter Details.
type RfEmitterGetResponseRfEmitterDetailAmplifier struct {
	// The device identifier of the amplifier.
	DeviceIdentifier string `json:"deviceIdentifier"`
	// The manufacturer of the amplifier.
	Manufacturer string `json:"manufacturer"`
	// The model name of the amplifier.
	ModelName string `json:"modelName"`
	// The amplifier power level, in watts.
	Power float64 `json:"power"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceIdentifier respjson.Field
		Manufacturer     respjson.Field
		ModelName        respjson.Field
		Power            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailAmplifier) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterGetResponseRfEmitterDetailAmplifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna associated with an RF Emitter Details.
type RfEmitterGetResponseRfEmitterDetailAntenna struct {
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter float64 `json:"antennaDiameter"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize"`
	// A flag to indicate whether the antenna points at a fixed azimuth/elevation.
	AzElFixed bool `json:"azElFixed"`
	// The set of antenna feeds for this antenna.
	Feeds []RfEmitterGetResponseRfEmitterDetailAntennaFeed `json:"feeds"`
	// Antenna azimuth, in degrees clockwise from true North, for a fixed antenna.
	FixedAzimuth float64 `json:"fixedAzimuth"`
	// Antenna elevation, in degrees, for a fixed antenna.
	FixedElevation float64 `json:"fixedElevation"`
	// Array of maximum azimuths, in degrees.
	MaxAzimuths []float64 `json:"maxAzimuths"`
	// Maximum elevation, in degrees.
	MaxElevation float64 `json:"maxElevation"`
	// Array of minimum azimuths, in degrees.
	MinAzimuths []float64 `json:"minAzimuths"`
	// Minimum elevation, in degrees.
	MinElevation float64 `json:"minElevation"`
	// The name of the antenna.
	Name string `json:"name"`
	// The set of receiver channels for this antenna.
	ReceiverChannels []RfEmitterGetResponseRfEmitterDetailAntennaReceiverChannel `json:"receiverChannels"`
	// The set of transmit channels for this antenna.
	TransmitChannels []RfEmitterGetResponseRfEmitterDetailAntennaTransmitChannel `json:"transmitChannels"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AntennaDiameter  respjson.Field
		AntennaSize      respjson.Field
		AzElFixed        respjson.Field
		Feeds            respjson.Field
		FixedAzimuth     respjson.Field
		FixedElevation   respjson.Field
		MaxAzimuths      respjson.Field
		MaxElevation     respjson.Field
		MinAzimuths      respjson.Field
		MinElevation     respjson.Field
		Name             respjson.Field
		ReceiverChannels respjson.Field
		TransmitChannels respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailAntenna) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterGetResponseRfEmitterDetailAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Feed associated with an RF Antenna.
type RfEmitterGetResponseRfEmitterDetailAntennaFeed struct {
	// Maximum frequency, in megahertz.
	FreqMax float64 `json:"freqMax"`
	// Minimum frequency, in megahertz.
	FreqMin float64 `json:"freqMin"`
	// The feed name.
	Name string `json:"name"`
	// The antenna feed linear/circular polarization (e.g. HORIZONTAL, VERTICAL,
	// LEFT_HAND_CIRCULAR, RIGHT_HAND_CIRCULAR).
	Polarization string `json:"polarization"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FreqMax      respjson.Field
		FreqMin      respjson.Field
		Name         respjson.Field
		Polarization respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailAntennaFeed) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterGetResponseRfEmitterDetailAntennaFeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Receiver Channel associated with an RF Antenna.
type RfEmitterGetResponseRfEmitterDetailAntennaReceiverChannel struct {
	// The receiver bandwidth, in megahertz, must satisfy the constraint: minBandwidth
	// ≤ bandwidth ≤ maxBandwidth.
	Bandwidth float64 `json:"bandwidth"`
	// The receive channel number.
	ChannelNum string `json:"channelNum"`
	// The receive channel device identifier.
	DeviceIdentifier string `json:"deviceIdentifier"`
	// Maximum frequency, in megahertz.
	FreqMax float64 `json:"freqMax"`
	// Minimum frequency, in megahertz.
	FreqMin float64 `json:"freqMin"`
	// The maximum receiver bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MaxBandwidth float64 `json:"maxBandwidth"`
	// The receiver bandwidth, in megahertz, must satisfy the constraint: minBandwidth
	// ≤ bandwidth ≤ maxBandwidth.
	MinBandwidth float64 `json:"minBandwidth"`
	// Receiver sensitivity, in decibel-milliwatts.
	Sensitivity float64 `json:"sensitivity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bandwidth        respjson.Field
		ChannelNum       respjson.Field
		DeviceIdentifier respjson.Field
		FreqMax          respjson.Field
		FreqMin          respjson.Field
		MaxBandwidth     respjson.Field
		MinBandwidth     respjson.Field
		Sensitivity      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailAntennaReceiverChannel) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterGetResponseRfEmitterDetailAntennaReceiverChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Transmit Channel associated with an RF Antenna.
type RfEmitterGetResponseRfEmitterDetailAntennaTransmitChannel struct {
	// Transmit power, in watts.
	Power float64 `json:"power,required"`
	// The transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	Bandwidth float64 `json:"bandwidth"`
	// The transmit channel number.
	ChannelNum string `json:"channelNum"`
	// The transmit channel device identifier.
	DeviceIdentifier string `json:"deviceIdentifier"`
	// The transmitter frequency, in megahertz, must satisfy the constraint: freqMin <=
	// freq <= freqMax.
	Freq float64 `json:"freq"`
	// The maximum transmitter frequency, in megahertz, must satisfy the constraint:
	// freqMin ≤ freq ≤ freqMax.
	FreqMax float64 `json:"freqMax"`
	// The minimum transmitter frequency, in megahertz, must satisfy the constraint:
	// freqMin ≤ freq ≤ freqMax.
	FreqMin float64 `json:"freqMin"`
	// The hardware sample rate, in bits per second for this transmit channel.
	HardwareSampleRate int64 `json:"hardwareSampleRate"`
	// The maximum transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MaxBandwidth float64 `json:"maxBandwidth"`
	// Maximum gain, in decibels.
	MaxGain float64 `json:"maxGain"`
	// The minimum transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MinBandwidth float64 `json:"minBandwidth"`
	// Minimum gain, in decibels.
	MinGain float64 `json:"minGain"`
	// The set of sample rates supported by this transmit channel, in bits per second.
	SampleRates []float64 `json:"sampleRates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Power              respjson.Field
		Bandwidth          respjson.Field
		ChannelNum         respjson.Field
		DeviceIdentifier   respjson.Field
		Freq               respjson.Field
		FreqMax            respjson.Field
		FreqMin            respjson.Field
		HardwareSampleRate respjson.Field
		MaxBandwidth       respjson.Field
		MaxGain            respjson.Field
		MinBandwidth       respjson.Field
		MinGain            respjson.Field
		SampleRates        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailAntennaTransmitChannel) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterGetResponseRfEmitterDetailAntennaTransmitChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Power Offset associated with an RF Emitter Details.
type RfEmitterGetResponseRfEmitterDetailPowerOffset struct {
	// The RF frequency band (e.g. HF, VHF, P, UHF, L, S, C, X, KU, K, KA, V, W, MM).
	FrequencyBand string `json:"frequencyBand"`
	// Power offset, in decibels.
	PowerOffset float64 `json:"powerOffset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FrequencyBand respjson.Field
		PowerOffset   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailPowerOffset) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterGetResponseRfEmitterDetailPowerOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter SW Service associated with an RF Emitter Details.
type RfEmitterGetResponseRfEmitterDetailService struct {
	// The name for this software service.
	Name string `json:"name"`
	// The version for this software service.
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailService) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterGetResponseRfEmitterDetailService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter TTP associated with an RF Emitter Details.
type RfEmitterGetResponseRfEmitterDetailTtp struct {
	// The name of the output signal.
	OutputSignalName string `json:"outputSignalName"`
	// The set of TTPs affected by this signal.
	TechniqueDefinitions []RfEmitterGetResponseRfEmitterDetailTtpTechniqueDefinition `json:"techniqueDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutputSignalName     respjson.Field
		TechniqueDefinitions respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailTtp) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterGetResponseRfEmitterDetailTtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Definition associated with an RF Emitter TTP.
type RfEmitterGetResponseRfEmitterDetailTtpTechniqueDefinition struct {
	// The EW Emitter system technique name.
	Name string `json:"name"`
	// The set of required/optional parameters for this technique.
	ParamDefinitions []RfEmitterGetResponseRfEmitterDetailTtpTechniqueDefinitionParamDefinition `json:"paramDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name             respjson.Field
		ParamDefinitions respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailTtpTechniqueDefinition) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterGetResponseRfEmitterDetailTtpTechniqueDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Parameter Definition associated with an RF Emitter
// Technique Definition.
type RfEmitterGetResponseRfEmitterDetailTtpTechniqueDefinitionParamDefinition struct {
	// Default parameter value used if not overridden in a SEW task definition.
	DefaultValue string `json:"defaultValue"`
	// Maximum allowable value for a numeric parameter.
	Max float64 `json:"max"`
	// Minimum allowable value for a numeric parameter.
	Min float64 `json:"min"`
	// The name of the parameter.
	Name string `json:"name"`
	// A flag to specify that a parameter is optional.
	Optional bool `json:"optional"`
	// The type of parameter (e.g. STRING, DOUBLE, INT, LIST).
	Type string `json:"type"`
	// Units (degrees, seconds, decibels, etc.) for a numeric parameter.
	Units string `json:"units"`
	// Valid values for strictly defined parameters.
	ValidValues []string `json:"validValues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultValue respjson.Field
		Max          respjson.Field
		Min          respjson.Field
		Name         respjson.Field
		Optional     respjson.Field
		Type         respjson.Field
		Units        respjson.Field
		ValidValues  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterGetResponseRfEmitterDetailTtpTechniqueDefinitionParamDefinition) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterGetResponseRfEmitterDetailTtpTechniqueDefinitionParamDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterQueryhelpResponse struct {
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
func (r RfEmitterQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter is a source of active Radio Frequency (RF) signals which could
// potentially interfere with other RF receivers.
type RfEmitterTupleResponse struct {
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
	DataMode RfEmitterTupleResponseDataMode `json:"dataMode,required"`
	// Unique name of this RF Emitter.
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
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity shared.EntityFull `json:"entity"`
	// The originating system ID for the RF Emitter.
	ExtSysID string `json:"extSysId"`
	// ID by reference of the parent entity for this RFEmitter.
	IDEntity string `json:"idEntity"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Details about this RF Emitter.
	RfEmitterDetails []RfEmitterTupleResponseRfEmitterDetail `json:"rfEmitterDetails"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype string `json:"subtype"`
	// Type of this RF Emitter.
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
		Entity                respjson.Field
		ExtSysID              respjson.Field
		IDEntity              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RfEmitterDetails      respjson.Field
		Subtype               respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterTupleResponse) UnmarshalJSON(data []byte) error {
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
type RfEmitterTupleResponseDataMode string

const (
	RfEmitterTupleResponseDataModeReal      RfEmitterTupleResponseDataMode = "REAL"
	RfEmitterTupleResponseDataModeTest      RfEmitterTupleResponseDataMode = "TEST"
	RfEmitterTupleResponseDataModeSimulated RfEmitterTupleResponseDataMode = "SIMULATED"
	RfEmitterTupleResponseDataModeExercise  RfEmitterTupleResponseDataMode = "EXERCISE"
)

// Details for a particular RF Emitter, collected by a particular source. An RF
// Emitter may have multiple details records from various sources.
type RfEmitterTupleResponseRfEmitterDetail struct {
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
	// Unique identifier of the parent RF Emitter.
	IDRfEmitter string `json:"idRFEmitter,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate facility name for this RF Emitter.
	AlternateFacilityName string `json:"alternateFacilityName"`
	// Optional alternate name or alias for this RF Emitter.
	AltName string `json:"altName"`
	// An RF Amplifier associated with an RF Emitter Details.
	Amplifier RfEmitterTupleResponseRfEmitterDetailAmplifier `json:"amplifier"`
	// The set of antennas hosted on this EW Emitter system.
	Antennas []RfEmitterTupleResponseRfEmitterDetailAntenna `json:"antennas"`
	// Barrage noise bandwidth, in megahertz.
	BarrageNoiseBandwidth float64 `json:"barrageNoiseBandwidth"`
	// The length of time, in seconds, for the RF Emitter built-in test to run to
	// completion.
	BitRunTime float64 `json:"bitRunTime"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Detailed description of the RF Emitter.
	Description string `json:"description"`
	// Designator of this RF Emitter.
	Designator string `json:"designator"`
	// Doppler noise value, in megahertz.
	DopplerNoise float64 `json:"dopplerNoise"`
	// Digital Form Radio Memory instantaneous bandwidth in megahertz.
	DrfmInstantaneousBandwidth float64 `json:"drfmInstantaneousBandwidth"`
	// Family of this RF Emitter type.
	Family string `json:"family"`
	// A fixed attenuation value to be set on the SRF Emitter HPA when commanding an
	// Electronic Attack/Techniques Tactics and Procedures task, in decibels.
	FixedAttenuation float64 `json:"fixedAttenuation"`
	// Unique identifier of the organization which manufactured this RF Emitter.
	IDManufacturerOrg string `json:"idManufacturerOrg"`
	// Unique identifier of the location of the production facility for this RF
	// Emitter.
	IDProductionFacilityLocation string `json:"idProductionFacilityLocation"`
	// COCOM that has temporary responsibility for scheduling and management of the RF
	// Emitter (e.g. SPACEFOR-CENT, SPACEFOR-EURAF, SPACEFOR-INDOPAC, SPACEFOR-KOR,
	// SPACEFOR-STRATNORTH, SPACESOC, NONE).
	LoanedToCocom string `json:"loanedToCocom"`
	// Notes on the RF Emitter.
	Notes string `json:"notes"`
	// Number of bits.
	NumBits int64 `json:"numBits"`
	// Number of channels.
	NumChannels int64 `json:"numChannels"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// A set of system/frequency band adjustments to the power offset commanded in an
	// EA/TTP task.
	PowerOffsets []RfEmitterTupleResponseRfEmitterDetailPowerOffset `json:"powerOffsets"`
	// The length of time, in seconds, for the RF Emitter to prepare for a task,
	// including sufficient time to slew the antenna and configure the equipment.
	PrepTime float64 `json:"prepTime"`
	// Primary COCOM that is responsible for scheduling and management of the RF
	// Emitter (e.g. SPACEFOR-CENT, SPACEFOR-EURAF, SPACEFOR-INDOPAC, SPACEFOR-KOR,
	// SPACEFOR-STRATNORTH, SPACESOC, NONE).
	PrimaryCocom string `json:"primaryCocom"`
	// Name of the production facility for this RF Emitter.
	ProductionFacilityName string `json:"productionFacilityName"`
	// Type or name of receiver.
	ReceiverType string `json:"receiverType"`
	// Secondary notes on the RF Emitter.
	SecondaryNotes string `json:"secondaryNotes"`
	// The set of software services running on this EW Emitter system.
	Services []RfEmitterTupleResponseRfEmitterDetailService `json:"services"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. End sensitivity range, in
	// decibel-milliwatts.
	SystemSensitivityEnd float64 `json:"systemSensitivityEnd"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. Start sensitivity range, in
	// decibel-milliwatts.
	SystemSensitivityStart float64 `json:"systemSensitivityStart"`
	// The set of EA/TTP techniques that are supported by this EW Emitter system.
	Ttps []RfEmitterTupleResponseRfEmitterDetailTtp `json:"ttps"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Array of URLs containing additional information on this RF Emitter.
	URLs []string `json:"urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking        respjson.Field
		DataMode                     respjson.Field
		IDRfEmitter                  respjson.Field
		Source                       respjson.Field
		ID                           respjson.Field
		AlternateFacilityName        respjson.Field
		AltName                      respjson.Field
		Amplifier                    respjson.Field
		Antennas                     respjson.Field
		BarrageNoiseBandwidth        respjson.Field
		BitRunTime                   respjson.Field
		CreatedAt                    respjson.Field
		CreatedBy                    respjson.Field
		Description                  respjson.Field
		Designator                   respjson.Field
		DopplerNoise                 respjson.Field
		DrfmInstantaneousBandwidth   respjson.Field
		Family                       respjson.Field
		FixedAttenuation             respjson.Field
		IDManufacturerOrg            respjson.Field
		IDProductionFacilityLocation respjson.Field
		LoanedToCocom                respjson.Field
		Notes                        respjson.Field
		NumBits                      respjson.Field
		NumChannels                  respjson.Field
		Origin                       respjson.Field
		OrigNetwork                  respjson.Field
		PowerOffsets                 respjson.Field
		PrepTime                     respjson.Field
		PrimaryCocom                 respjson.Field
		ProductionFacilityName       respjson.Field
		ReceiverType                 respjson.Field
		SecondaryNotes               respjson.Field
		Services                     respjson.Field
		SystemSensitivityEnd         respjson.Field
		SystemSensitivityStart       respjson.Field
		Ttps                         respjson.Field
		UpdatedAt                    respjson.Field
		UpdatedBy                    respjson.Field
		URLs                         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetail) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterTupleResponseRfEmitterDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Amplifier associated with an RF Emitter Details.
type RfEmitterTupleResponseRfEmitterDetailAmplifier struct {
	// The device identifier of the amplifier.
	DeviceIdentifier string `json:"deviceIdentifier"`
	// The manufacturer of the amplifier.
	Manufacturer string `json:"manufacturer"`
	// The model name of the amplifier.
	ModelName string `json:"modelName"`
	// The amplifier power level, in watts.
	Power float64 `json:"power"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceIdentifier respjson.Field
		Manufacturer     respjson.Field
		ModelName        respjson.Field
		Power            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailAmplifier) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterTupleResponseRfEmitterDetailAmplifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna associated with an RF Emitter Details.
type RfEmitterTupleResponseRfEmitterDetailAntenna struct {
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter float64 `json:"antennaDiameter"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize"`
	// A flag to indicate whether the antenna points at a fixed azimuth/elevation.
	AzElFixed bool `json:"azElFixed"`
	// The set of antenna feeds for this antenna.
	Feeds []RfEmitterTupleResponseRfEmitterDetailAntennaFeed `json:"feeds"`
	// Antenna azimuth, in degrees clockwise from true North, for a fixed antenna.
	FixedAzimuth float64 `json:"fixedAzimuth"`
	// Antenna elevation, in degrees, for a fixed antenna.
	FixedElevation float64 `json:"fixedElevation"`
	// Array of maximum azimuths, in degrees.
	MaxAzimuths []float64 `json:"maxAzimuths"`
	// Maximum elevation, in degrees.
	MaxElevation float64 `json:"maxElevation"`
	// Array of minimum azimuths, in degrees.
	MinAzimuths []float64 `json:"minAzimuths"`
	// Minimum elevation, in degrees.
	MinElevation float64 `json:"minElevation"`
	// The name of the antenna.
	Name string `json:"name"`
	// The set of receiver channels for this antenna.
	ReceiverChannels []RfEmitterTupleResponseRfEmitterDetailAntennaReceiverChannel `json:"receiverChannels"`
	// The set of transmit channels for this antenna.
	TransmitChannels []RfEmitterTupleResponseRfEmitterDetailAntennaTransmitChannel `json:"transmitChannels"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AntennaDiameter  respjson.Field
		AntennaSize      respjson.Field
		AzElFixed        respjson.Field
		Feeds            respjson.Field
		FixedAzimuth     respjson.Field
		FixedElevation   respjson.Field
		MaxAzimuths      respjson.Field
		MaxElevation     respjson.Field
		MinAzimuths      respjson.Field
		MinElevation     respjson.Field
		Name             respjson.Field
		ReceiverChannels respjson.Field
		TransmitChannels respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailAntenna) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterTupleResponseRfEmitterDetailAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Feed associated with an RF Antenna.
type RfEmitterTupleResponseRfEmitterDetailAntennaFeed struct {
	// Maximum frequency, in megahertz.
	FreqMax float64 `json:"freqMax"`
	// Minimum frequency, in megahertz.
	FreqMin float64 `json:"freqMin"`
	// The feed name.
	Name string `json:"name"`
	// The antenna feed linear/circular polarization (e.g. HORIZONTAL, VERTICAL,
	// LEFT_HAND_CIRCULAR, RIGHT_HAND_CIRCULAR).
	Polarization string `json:"polarization"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FreqMax      respjson.Field
		FreqMin      respjson.Field
		Name         respjson.Field
		Polarization respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailAntennaFeed) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterTupleResponseRfEmitterDetailAntennaFeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Receiver Channel associated with an RF Antenna.
type RfEmitterTupleResponseRfEmitterDetailAntennaReceiverChannel struct {
	// The receiver bandwidth, in megahertz, must satisfy the constraint: minBandwidth
	// ≤ bandwidth ≤ maxBandwidth.
	Bandwidth float64 `json:"bandwidth"`
	// The receive channel number.
	ChannelNum string `json:"channelNum"`
	// The receive channel device identifier.
	DeviceIdentifier string `json:"deviceIdentifier"`
	// Maximum frequency, in megahertz.
	FreqMax float64 `json:"freqMax"`
	// Minimum frequency, in megahertz.
	FreqMin float64 `json:"freqMin"`
	// The maximum receiver bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MaxBandwidth float64 `json:"maxBandwidth"`
	// The receiver bandwidth, in megahertz, must satisfy the constraint: minBandwidth
	// ≤ bandwidth ≤ maxBandwidth.
	MinBandwidth float64 `json:"minBandwidth"`
	// Receiver sensitivity, in decibel-milliwatts.
	Sensitivity float64 `json:"sensitivity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bandwidth        respjson.Field
		ChannelNum       respjson.Field
		DeviceIdentifier respjson.Field
		FreqMax          respjson.Field
		FreqMin          respjson.Field
		MaxBandwidth     respjson.Field
		MinBandwidth     respjson.Field
		Sensitivity      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailAntennaReceiverChannel) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterTupleResponseRfEmitterDetailAntennaReceiverChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Transmit Channel associated with an RF Antenna.
type RfEmitterTupleResponseRfEmitterDetailAntennaTransmitChannel struct {
	// Transmit power, in watts.
	Power float64 `json:"power,required"`
	// The transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	Bandwidth float64 `json:"bandwidth"`
	// The transmit channel number.
	ChannelNum string `json:"channelNum"`
	// The transmit channel device identifier.
	DeviceIdentifier string `json:"deviceIdentifier"`
	// The transmitter frequency, in megahertz, must satisfy the constraint: freqMin <=
	// freq <= freqMax.
	Freq float64 `json:"freq"`
	// The maximum transmitter frequency, in megahertz, must satisfy the constraint:
	// freqMin ≤ freq ≤ freqMax.
	FreqMax float64 `json:"freqMax"`
	// The minimum transmitter frequency, in megahertz, must satisfy the constraint:
	// freqMin ≤ freq ≤ freqMax.
	FreqMin float64 `json:"freqMin"`
	// The hardware sample rate, in bits per second for this transmit channel.
	HardwareSampleRate int64 `json:"hardwareSampleRate"`
	// The maximum transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MaxBandwidth float64 `json:"maxBandwidth"`
	// Maximum gain, in decibels.
	MaxGain float64 `json:"maxGain"`
	// The minimum transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MinBandwidth float64 `json:"minBandwidth"`
	// Minimum gain, in decibels.
	MinGain float64 `json:"minGain"`
	// The set of sample rates supported by this transmit channel, in bits per second.
	SampleRates []float64 `json:"sampleRates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Power              respjson.Field
		Bandwidth          respjson.Field
		ChannelNum         respjson.Field
		DeviceIdentifier   respjson.Field
		Freq               respjson.Field
		FreqMax            respjson.Field
		FreqMin            respjson.Field
		HardwareSampleRate respjson.Field
		MaxBandwidth       respjson.Field
		MaxGain            respjson.Field
		MinBandwidth       respjson.Field
		MinGain            respjson.Field
		SampleRates        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailAntennaTransmitChannel) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterTupleResponseRfEmitterDetailAntennaTransmitChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Power Offset associated with an RF Emitter Details.
type RfEmitterTupleResponseRfEmitterDetailPowerOffset struct {
	// The RF frequency band (e.g. HF, VHF, P, UHF, L, S, C, X, KU, K, KA, V, W, MM).
	FrequencyBand string `json:"frequencyBand"`
	// Power offset, in decibels.
	PowerOffset float64 `json:"powerOffset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FrequencyBand respjson.Field
		PowerOffset   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailPowerOffset) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterTupleResponseRfEmitterDetailPowerOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter SW Service associated with an RF Emitter Details.
type RfEmitterTupleResponseRfEmitterDetailService struct {
	// The name for this software service.
	Name string `json:"name"`
	// The version for this software service.
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailService) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterTupleResponseRfEmitterDetailService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter TTP associated with an RF Emitter Details.
type RfEmitterTupleResponseRfEmitterDetailTtp struct {
	// The name of the output signal.
	OutputSignalName string `json:"outputSignalName"`
	// The set of TTPs affected by this signal.
	TechniqueDefinitions []RfEmitterTupleResponseRfEmitterDetailTtpTechniqueDefinition `json:"techniqueDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutputSignalName     respjson.Field
		TechniqueDefinitions respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailTtp) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterTupleResponseRfEmitterDetailTtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Definition associated with an RF Emitter TTP.
type RfEmitterTupleResponseRfEmitterDetailTtpTechniqueDefinition struct {
	// The EW Emitter system technique name.
	Name string `json:"name"`
	// The set of required/optional parameters for this technique.
	ParamDefinitions []RfEmitterTupleResponseRfEmitterDetailTtpTechniqueDefinitionParamDefinition `json:"paramDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name             respjson.Field
		ParamDefinitions respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailTtpTechniqueDefinition) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterTupleResponseRfEmitterDetailTtpTechniqueDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Parameter Definition associated with an RF Emitter
// Technique Definition.
type RfEmitterTupleResponseRfEmitterDetailTtpTechniqueDefinitionParamDefinition struct {
	// Default parameter value used if not overridden in a SEW task definition.
	DefaultValue string `json:"defaultValue"`
	// Maximum allowable value for a numeric parameter.
	Max float64 `json:"max"`
	// Minimum allowable value for a numeric parameter.
	Min float64 `json:"min"`
	// The name of the parameter.
	Name string `json:"name"`
	// A flag to specify that a parameter is optional.
	Optional bool `json:"optional"`
	// The type of parameter (e.g. STRING, DOUBLE, INT, LIST).
	Type string `json:"type"`
	// Units (degrees, seconds, decibels, etc.) for a numeric parameter.
	Units string `json:"units"`
	// Valid values for strictly defined parameters.
	ValidValues []string `json:"validValues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultValue respjson.Field
		Max          respjson.Field
		Min          respjson.Field
		Name         respjson.Field
		Optional     respjson.Field
		Type         respjson.Field
		Units        respjson.Field
		ValidValues  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterTupleResponseRfEmitterDetailTtpTechniqueDefinitionParamDefinition) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterTupleResponseRfEmitterDetailTtpTechniqueDefinitionParamDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterNewParams struct {
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
	DataMode RfEmitterNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique name of this RF Emitter.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The originating system ID for the RF Emitter.
	ExtSysID param.Opt[string] `json:"extSysId,omitzero"`
	// ID by reference of the parent entity for this RFEmitter.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype param.Opt[string] `json:"subtype,omitzero"`
	// Type of this RF Emitter.
	Type param.Opt[string] `json:"type,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r RfEmitterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterNewParams) UnmarshalJSON(data []byte) error {
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
type RfEmitterNewParamsDataMode string

const (
	RfEmitterNewParamsDataModeReal      RfEmitterNewParamsDataMode = "REAL"
	RfEmitterNewParamsDataModeTest      RfEmitterNewParamsDataMode = "TEST"
	RfEmitterNewParamsDataModeSimulated RfEmitterNewParamsDataMode = "SIMULATED"
	RfEmitterNewParamsDataModeExercise  RfEmitterNewParamsDataMode = "EXERCISE"
)

type RfEmitterUpdateParams struct {
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
	DataMode RfEmitterUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique name of this RF Emitter.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The originating system ID for the RF Emitter.
	ExtSysID param.Opt[string] `json:"extSysId,omitzero"`
	// ID by reference of the parent entity for this RFEmitter.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype param.Opt[string] `json:"subtype,omitzero"`
	// Type of this RF Emitter.
	Type param.Opt[string] `json:"type,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r RfEmitterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterUpdateParams) UnmarshalJSON(data []byte) error {
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
type RfEmitterUpdateParamsDataMode string

const (
	RfEmitterUpdateParamsDataModeReal      RfEmitterUpdateParamsDataMode = "REAL"
	RfEmitterUpdateParamsDataModeTest      RfEmitterUpdateParamsDataMode = "TEST"
	RfEmitterUpdateParamsDataModeSimulated RfEmitterUpdateParamsDataMode = "SIMULATED"
	RfEmitterUpdateParamsDataModeExercise  RfEmitterUpdateParamsDataMode = "EXERCISE"
)

type RfEmitterListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterListParams]'s query parameters as `url.Values`.
func (r RfEmitterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfEmitterCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterCountParams]'s query parameters as `url.Values`.
func (r RfEmitterCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfEmitterGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterGetParams]'s query parameters as `url.Values`.
func (r RfEmitterGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfEmitterTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterTupleParams]'s query parameters as `url.Values`.
func (r RfEmitterTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
