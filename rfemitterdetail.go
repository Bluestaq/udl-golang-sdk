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

// RfEmitterDetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRfEmitterDetailService] method instead.
type RfEmitterDetailService struct {
	Options []option.RequestOption
}

// NewRfEmitterDetailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRfEmitterDetailService(opts ...option.RequestOption) (r RfEmitterDetailService) {
	r = RfEmitterDetailService{}
	r.Options = opts
	return
}

// Service operation to take a single RFEmitterDetails as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *RfEmitterDetailService) New(ctx context.Context, body RfEmitterDetailNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/rfemitterdetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single RFEmitterDetails record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *RfEmitterDetailService) Update(ctx context.Context, id string, body RfEmitterDetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfemitterdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *RfEmitterDetailService) List(ctx context.Context, query RfEmitterDetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[RfEmitterDetailListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/rfemitterdetails"
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
func (r *RfEmitterDetailService) ListAutoPaging(ctx context.Context, query RfEmitterDetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[RfEmitterDetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a single RFEmitterDetails record specified by the
// passed ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *RfEmitterDetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfemitterdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *RfEmitterDetailService) Count(ctx context.Context, query RfEmitterDetailCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/rfemitterdetails/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single RFEmitterDetails record by its unique ID
// passed as a path parameter.
func (r *RfEmitterDetailService) Get(ctx context.Context, id string, query RfEmitterDetailGetParams, opts ...option.RequestOption) (res *RfEmitterDetailGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfemitterdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *RfEmitterDetailService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *RfEmitterDetailQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/rfemitterdetails/queryhelp"
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
func (r *RfEmitterDetailService) Tuple(ctx context.Context, query RfEmitterDetailTupleParams, opts ...option.RequestOption) (res *[]RfEmitterDetailTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/rfemitterdetails/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details for a particular RF Emitter, collected by a particular source. An RF
// Emitter may have multiple details records from various sources.
type RfEmitterDetailListResponse struct {
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
	DataMode RfEmitterDetailListResponseDataMode `json:"dataMode,required"`
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
	Amplifier RfEmitterDetailListResponseAmplifier `json:"amplifier"`
	// The set of antennas hosted on this EW Emitter system.
	Antennas []RfEmitterDetailListResponseAntenna `json:"antennas"`
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
	PowerOffsets []RfEmitterDetailListResponsePowerOffset `json:"powerOffsets"`
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
	Services []RfEmitterDetailListResponseService `json:"services"`
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
	Ttps []RfEmitterDetailListResponseTtp `json:"ttps"`
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
func (r RfEmitterDetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponse) UnmarshalJSON(data []byte) error {
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
type RfEmitterDetailListResponseDataMode string

const (
	RfEmitterDetailListResponseDataModeReal      RfEmitterDetailListResponseDataMode = "REAL"
	RfEmitterDetailListResponseDataModeTest      RfEmitterDetailListResponseDataMode = "TEST"
	RfEmitterDetailListResponseDataModeSimulated RfEmitterDetailListResponseDataMode = "SIMULATED"
	RfEmitterDetailListResponseDataModeExercise  RfEmitterDetailListResponseDataMode = "EXERCISE"
)

// An RF Amplifier associated with an RF Emitter Details.
type RfEmitterDetailListResponseAmplifier struct {
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
func (r RfEmitterDetailListResponseAmplifier) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponseAmplifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna associated with an RF Emitter Details.
type RfEmitterDetailListResponseAntenna struct {
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter float64 `json:"antennaDiameter"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize"`
	// A flag to indicate whether the antenna points at a fixed azimuth/elevation.
	AzElFixed bool `json:"azElFixed"`
	// The set of antenna feeds for this antenna.
	Feeds []RfEmitterDetailListResponseAntennaFeed `json:"feeds"`
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
	ReceiverChannels []RfEmitterDetailListResponseAntennaReceiverChannel `json:"receiverChannels"`
	// The set of transmit channels for this antenna.
	TransmitChannels []RfEmitterDetailListResponseAntennaTransmitChannel `json:"transmitChannels"`
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
func (r RfEmitterDetailListResponseAntenna) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponseAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Feed associated with an RF Antenna.
type RfEmitterDetailListResponseAntennaFeed struct {
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
func (r RfEmitterDetailListResponseAntennaFeed) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponseAntennaFeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Receiver Channel associated with an RF Antenna.
type RfEmitterDetailListResponseAntennaReceiverChannel struct {
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
func (r RfEmitterDetailListResponseAntennaReceiverChannel) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponseAntennaReceiverChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Transmit Channel associated with an RF Antenna.
type RfEmitterDetailListResponseAntennaTransmitChannel struct {
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
func (r RfEmitterDetailListResponseAntennaTransmitChannel) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponseAntennaTransmitChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Power Offset associated with an RF Emitter Details.
type RfEmitterDetailListResponsePowerOffset struct {
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
func (r RfEmitterDetailListResponsePowerOffset) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponsePowerOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter SW Service associated with an RF Emitter Details.
type RfEmitterDetailListResponseService struct {
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
func (r RfEmitterDetailListResponseService) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter TTP associated with an RF Emitter Details.
type RfEmitterDetailListResponseTtp struct {
	// The name of the output signal.
	OutputSignalName string `json:"outputSignalName"`
	// The set of TTPs affected by this signal.
	TechniqueDefinitions []RfEmitterDetailListResponseTtpTechniqueDefinition `json:"techniqueDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutputSignalName     respjson.Field
		TechniqueDefinitions respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterDetailListResponseTtp) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponseTtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Definition associated with an RF Emitter TTP.
type RfEmitterDetailListResponseTtpTechniqueDefinition struct {
	// The EW Emitter system technique name.
	Name string `json:"name"`
	// The set of required/optional parameters for this technique.
	ParamDefinitions []RfEmitterDetailListResponseTtpTechniqueDefinitionParamDefinition `json:"paramDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name             respjson.Field
		ParamDefinitions respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterDetailListResponseTtpTechniqueDefinition) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailListResponseTtpTechniqueDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Parameter Definition associated with an RF Emitter
// Technique Definition.
type RfEmitterDetailListResponseTtpTechniqueDefinitionParamDefinition struct {
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
func (r RfEmitterDetailListResponseTtpTechniqueDefinitionParamDefinition) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterDetailListResponseTtpTechniqueDefinitionParamDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for a particular RF Emitter, collected by a particular source. An RF
// Emitter may have multiple details records from various sources.
type RfEmitterDetailGetResponse struct {
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
	DataMode RfEmitterDetailGetResponseDataMode `json:"dataMode,required"`
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
	Amplifier RfEmitterDetailGetResponseAmplifier `json:"amplifier"`
	// The set of antennas hosted on this EW Emitter system.
	Antennas []RfEmitterDetailGetResponseAntenna `json:"antennas"`
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
	PowerOffsets []RfEmitterDetailGetResponsePowerOffset `json:"powerOffsets"`
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
	Services []RfEmitterDetailGetResponseService `json:"services"`
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
	Ttps []RfEmitterDetailGetResponseTtp `json:"ttps"`
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
func (r RfEmitterDetailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponse) UnmarshalJSON(data []byte) error {
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
type RfEmitterDetailGetResponseDataMode string

const (
	RfEmitterDetailGetResponseDataModeReal      RfEmitterDetailGetResponseDataMode = "REAL"
	RfEmitterDetailGetResponseDataModeTest      RfEmitterDetailGetResponseDataMode = "TEST"
	RfEmitterDetailGetResponseDataModeSimulated RfEmitterDetailGetResponseDataMode = "SIMULATED"
	RfEmitterDetailGetResponseDataModeExercise  RfEmitterDetailGetResponseDataMode = "EXERCISE"
)

// An RF Amplifier associated with an RF Emitter Details.
type RfEmitterDetailGetResponseAmplifier struct {
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
func (r RfEmitterDetailGetResponseAmplifier) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponseAmplifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna associated with an RF Emitter Details.
type RfEmitterDetailGetResponseAntenna struct {
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter float64 `json:"antennaDiameter"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize"`
	// A flag to indicate whether the antenna points at a fixed azimuth/elevation.
	AzElFixed bool `json:"azElFixed"`
	// The set of antenna feeds for this antenna.
	Feeds []RfEmitterDetailGetResponseAntennaFeed `json:"feeds"`
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
	ReceiverChannels []RfEmitterDetailGetResponseAntennaReceiverChannel `json:"receiverChannels"`
	// The set of transmit channels for this antenna.
	TransmitChannels []RfEmitterDetailGetResponseAntennaTransmitChannel `json:"transmitChannels"`
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
func (r RfEmitterDetailGetResponseAntenna) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponseAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Feed associated with an RF Antenna.
type RfEmitterDetailGetResponseAntennaFeed struct {
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
func (r RfEmitterDetailGetResponseAntennaFeed) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponseAntennaFeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Receiver Channel associated with an RF Antenna.
type RfEmitterDetailGetResponseAntennaReceiverChannel struct {
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
func (r RfEmitterDetailGetResponseAntennaReceiverChannel) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponseAntennaReceiverChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Transmit Channel associated with an RF Antenna.
type RfEmitterDetailGetResponseAntennaTransmitChannel struct {
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
func (r RfEmitterDetailGetResponseAntennaTransmitChannel) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponseAntennaTransmitChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Power Offset associated with an RF Emitter Details.
type RfEmitterDetailGetResponsePowerOffset struct {
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
func (r RfEmitterDetailGetResponsePowerOffset) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponsePowerOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter SW Service associated with an RF Emitter Details.
type RfEmitterDetailGetResponseService struct {
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
func (r RfEmitterDetailGetResponseService) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter TTP associated with an RF Emitter Details.
type RfEmitterDetailGetResponseTtp struct {
	// The name of the output signal.
	OutputSignalName string `json:"outputSignalName"`
	// The set of TTPs affected by this signal.
	TechniqueDefinitions []RfEmitterDetailGetResponseTtpTechniqueDefinition `json:"techniqueDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutputSignalName     respjson.Field
		TechniqueDefinitions respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterDetailGetResponseTtp) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponseTtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Definition associated with an RF Emitter TTP.
type RfEmitterDetailGetResponseTtpTechniqueDefinition struct {
	// The EW Emitter system technique name.
	Name string `json:"name"`
	// The set of required/optional parameters for this technique.
	ParamDefinitions []RfEmitterDetailGetResponseTtpTechniqueDefinitionParamDefinition `json:"paramDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name             respjson.Field
		ParamDefinitions respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterDetailGetResponseTtpTechniqueDefinition) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailGetResponseTtpTechniqueDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Parameter Definition associated with an RF Emitter
// Technique Definition.
type RfEmitterDetailGetResponseTtpTechniqueDefinitionParamDefinition struct {
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
func (r RfEmitterDetailGetResponseTtpTechniqueDefinitionParamDefinition) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterDetailGetResponseTtpTechniqueDefinitionParamDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterDetailQueryhelpResponse struct {
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
func (r RfEmitterDetailQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for a particular RF Emitter, collected by a particular source. An RF
// Emitter may have multiple details records from various sources.
type RfEmitterDetailTupleResponse struct {
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
	DataMode RfEmitterDetailTupleResponseDataMode `json:"dataMode,required"`
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
	Amplifier RfEmitterDetailTupleResponseAmplifier `json:"amplifier"`
	// The set of antennas hosted on this EW Emitter system.
	Antennas []RfEmitterDetailTupleResponseAntenna `json:"antennas"`
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
	PowerOffsets []RfEmitterDetailTupleResponsePowerOffset `json:"powerOffsets"`
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
	Services []RfEmitterDetailTupleResponseService `json:"services"`
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
	Ttps []RfEmitterDetailTupleResponseTtp `json:"ttps"`
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
func (r RfEmitterDetailTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponse) UnmarshalJSON(data []byte) error {
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
type RfEmitterDetailTupleResponseDataMode string

const (
	RfEmitterDetailTupleResponseDataModeReal      RfEmitterDetailTupleResponseDataMode = "REAL"
	RfEmitterDetailTupleResponseDataModeTest      RfEmitterDetailTupleResponseDataMode = "TEST"
	RfEmitterDetailTupleResponseDataModeSimulated RfEmitterDetailTupleResponseDataMode = "SIMULATED"
	RfEmitterDetailTupleResponseDataModeExercise  RfEmitterDetailTupleResponseDataMode = "EXERCISE"
)

// An RF Amplifier associated with an RF Emitter Details.
type RfEmitterDetailTupleResponseAmplifier struct {
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
func (r RfEmitterDetailTupleResponseAmplifier) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponseAmplifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna associated with an RF Emitter Details.
type RfEmitterDetailTupleResponseAntenna struct {
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter float64 `json:"antennaDiameter"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize"`
	// A flag to indicate whether the antenna points at a fixed azimuth/elevation.
	AzElFixed bool `json:"azElFixed"`
	// The set of antenna feeds for this antenna.
	Feeds []RfEmitterDetailTupleResponseAntennaFeed `json:"feeds"`
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
	ReceiverChannels []RfEmitterDetailTupleResponseAntennaReceiverChannel `json:"receiverChannels"`
	// The set of transmit channels for this antenna.
	TransmitChannels []RfEmitterDetailTupleResponseAntennaTransmitChannel `json:"transmitChannels"`
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
func (r RfEmitterDetailTupleResponseAntenna) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponseAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Feed associated with an RF Antenna.
type RfEmitterDetailTupleResponseAntennaFeed struct {
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
func (r RfEmitterDetailTupleResponseAntennaFeed) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponseAntennaFeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Receiver Channel associated with an RF Antenna.
type RfEmitterDetailTupleResponseAntennaReceiverChannel struct {
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
func (r RfEmitterDetailTupleResponseAntennaReceiverChannel) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponseAntennaReceiverChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Transmit Channel associated with an RF Antenna.
type RfEmitterDetailTupleResponseAntennaTransmitChannel struct {
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
func (r RfEmitterDetailTupleResponseAntennaTransmitChannel) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponseAntennaTransmitChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Power Offset associated with an RF Emitter Details.
type RfEmitterDetailTupleResponsePowerOffset struct {
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
func (r RfEmitterDetailTupleResponsePowerOffset) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponsePowerOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter SW Service associated with an RF Emitter Details.
type RfEmitterDetailTupleResponseService struct {
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
func (r RfEmitterDetailTupleResponseService) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter TTP associated with an RF Emitter Details.
type RfEmitterDetailTupleResponseTtp struct {
	// The name of the output signal.
	OutputSignalName string `json:"outputSignalName"`
	// The set of TTPs affected by this signal.
	TechniqueDefinitions []RfEmitterDetailTupleResponseTtpTechniqueDefinition `json:"techniqueDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutputSignalName     respjson.Field
		TechniqueDefinitions respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterDetailTupleResponseTtp) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponseTtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Definition associated with an RF Emitter TTP.
type RfEmitterDetailTupleResponseTtpTechniqueDefinition struct {
	// The EW Emitter system technique name.
	Name string `json:"name"`
	// The set of required/optional parameters for this technique.
	ParamDefinitions []RfEmitterDetailTupleResponseTtpTechniqueDefinitionParamDefinition `json:"paramDefinitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name             respjson.Field
		ParamDefinitions respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterDetailTupleResponseTtpTechniqueDefinition) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterDetailTupleResponseTtpTechniqueDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Parameter Definition associated with an RF Emitter
// Technique Definition.
type RfEmitterDetailTupleResponseTtpTechniqueDefinitionParamDefinition struct {
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
func (r RfEmitterDetailTupleResponseTtpTechniqueDefinitionParamDefinition) RawJSON() string {
	return r.JSON.raw
}
func (r *RfEmitterDetailTupleResponseTtpTechniqueDefinitionParamDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterDetailNewParams struct {
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
	DataMode RfEmitterDetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent RF Emitter.
	IDRfEmitter string `json:"idRFEmitter,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate facility name for this RF Emitter.
	AlternateFacilityName param.Opt[string] `json:"alternateFacilityName,omitzero"`
	// Optional alternate name or alias for this RF Emitter.
	AltName param.Opt[string] `json:"altName,omitzero"`
	// Barrage noise bandwidth, in megahertz.
	BarrageNoiseBandwidth param.Opt[float64] `json:"barrageNoiseBandwidth,omitzero"`
	// The length of time, in seconds, for the RF Emitter built-in test to run to
	// completion.
	BitRunTime param.Opt[float64] `json:"bitRunTime,omitzero"`
	// Detailed description of the RF Emitter.
	Description param.Opt[string] `json:"description,omitzero"`
	// Designator of this RF Emitter.
	Designator param.Opt[string] `json:"designator,omitzero"`
	// Doppler noise value, in megahertz.
	DopplerNoise param.Opt[float64] `json:"dopplerNoise,omitzero"`
	// Digital Form Radio Memory instantaneous bandwidth in megahertz.
	DrfmInstantaneousBandwidth param.Opt[float64] `json:"drfmInstantaneousBandwidth,omitzero"`
	// Family of this RF Emitter type.
	Family param.Opt[string] `json:"family,omitzero"`
	// A fixed attenuation value to be set on the SRF Emitter HPA when commanding an
	// Electronic Attack/Techniques Tactics and Procedures task, in decibels.
	FixedAttenuation param.Opt[float64] `json:"fixedAttenuation,omitzero"`
	// Unique identifier of the organization which manufactured this RF Emitter.
	IDManufacturerOrg param.Opt[string] `json:"idManufacturerOrg,omitzero"`
	// Unique identifier of the location of the production facility for this RF
	// Emitter.
	IDProductionFacilityLocation param.Opt[string] `json:"idProductionFacilityLocation,omitzero"`
	// COCOM that has temporary responsibility for scheduling and management of the RF
	// Emitter (e.g. SPACEFOR-CENT, SPACEFOR-EURAF, SPACEFOR-INDOPAC, SPACEFOR-KOR,
	// SPACEFOR-STRATNORTH, SPACESOC, NONE).
	LoanedToCocom param.Opt[string] `json:"loanedToCocom,omitzero"`
	// Notes on the RF Emitter.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of bits.
	NumBits param.Opt[int64] `json:"numBits,omitzero"`
	// Number of channels.
	NumChannels param.Opt[int64] `json:"numChannels,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The length of time, in seconds, for the RF Emitter to prepare for a task,
	// including sufficient time to slew the antenna and configure the equipment.
	PrepTime param.Opt[float64] `json:"prepTime,omitzero"`
	// Primary COCOM that is responsible for scheduling and management of the RF
	// Emitter (e.g. SPACEFOR-CENT, SPACEFOR-EURAF, SPACEFOR-INDOPAC, SPACEFOR-KOR,
	// SPACEFOR-STRATNORTH, SPACESOC, NONE).
	PrimaryCocom param.Opt[string] `json:"primaryCocom,omitzero"`
	// Name of the production facility for this RF Emitter.
	ProductionFacilityName param.Opt[string] `json:"productionFacilityName,omitzero"`
	// Type or name of receiver.
	ReceiverType param.Opt[string] `json:"receiverType,omitzero"`
	// Secondary notes on the RF Emitter.
	SecondaryNotes param.Opt[string] `json:"secondaryNotes,omitzero"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. End sensitivity range, in
	// decibel-milliwatts.
	SystemSensitivityEnd param.Opt[float64] `json:"systemSensitivityEnd,omitzero"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. Start sensitivity range, in
	// decibel-milliwatts.
	SystemSensitivityStart param.Opt[float64] `json:"systemSensitivityStart,omitzero"`
	// An RF Amplifier associated with an RF Emitter Details.
	Amplifier RfEmitterDetailNewParamsAmplifier `json:"amplifier,omitzero"`
	// The set of antennas hosted on this EW Emitter system.
	Antennas []RfEmitterDetailNewParamsAntenna `json:"antennas,omitzero"`
	// A set of system/frequency band adjustments to the power offset commanded in an
	// EA/TTP task.
	PowerOffsets []RfEmitterDetailNewParamsPowerOffset `json:"powerOffsets,omitzero"`
	// The set of software services running on this EW Emitter system.
	Services []RfEmitterDetailNewParamsService `json:"services,omitzero"`
	// The set of EA/TTP techniques that are supported by this EW Emitter system.
	Ttps []RfEmitterDetailNewParamsTtp `json:"ttps,omitzero"`
	// Array of URLs containing additional information on this RF Emitter.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParams) UnmarshalJSON(data []byte) error {
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
type RfEmitterDetailNewParamsDataMode string

const (
	RfEmitterDetailNewParamsDataModeReal      RfEmitterDetailNewParamsDataMode = "REAL"
	RfEmitterDetailNewParamsDataModeTest      RfEmitterDetailNewParamsDataMode = "TEST"
	RfEmitterDetailNewParamsDataModeSimulated RfEmitterDetailNewParamsDataMode = "SIMULATED"
	RfEmitterDetailNewParamsDataModeExercise  RfEmitterDetailNewParamsDataMode = "EXERCISE"
)

// An RF Amplifier associated with an RF Emitter Details.
type RfEmitterDetailNewParamsAmplifier struct {
	// The device identifier of the amplifier.
	DeviceIdentifier param.Opt[string] `json:"deviceIdentifier,omitzero"`
	// The manufacturer of the amplifier.
	Manufacturer param.Opt[string] `json:"manufacturer,omitzero"`
	// The model name of the amplifier.
	ModelName param.Opt[string] `json:"modelName,omitzero"`
	// The amplifier power level, in watts.
	Power param.Opt[float64] `json:"power,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsAmplifier) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsAmplifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsAmplifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna associated with an RF Emitter Details.
type RfEmitterDetailNewParamsAntenna struct {
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter param.Opt[float64] `json:"antennaDiameter,omitzero"`
	// A flag to indicate whether the antenna points at a fixed azimuth/elevation.
	AzElFixed param.Opt[bool] `json:"azElFixed,omitzero"`
	// Antenna azimuth, in degrees clockwise from true North, for a fixed antenna.
	FixedAzimuth param.Opt[float64] `json:"fixedAzimuth,omitzero"`
	// Antenna elevation, in degrees, for a fixed antenna.
	FixedElevation param.Opt[float64] `json:"fixedElevation,omitzero"`
	// Maximum elevation, in degrees.
	MaxElevation param.Opt[float64] `json:"maxElevation,omitzero"`
	// Minimum elevation, in degrees.
	MinElevation param.Opt[float64] `json:"minElevation,omitzero"`
	// The name of the antenna.
	Name param.Opt[string] `json:"name,omitzero"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize,omitzero"`
	// The set of antenna feeds for this antenna.
	Feeds []RfEmitterDetailNewParamsAntennaFeed `json:"feeds,omitzero"`
	// Array of maximum azimuths, in degrees.
	MaxAzimuths []float64 `json:"maxAzimuths,omitzero"`
	// Array of minimum azimuths, in degrees.
	MinAzimuths []float64 `json:"minAzimuths,omitzero"`
	// The set of receiver channels for this antenna.
	ReceiverChannels []RfEmitterDetailNewParamsAntennaReceiverChannel `json:"receiverChannels,omitzero"`
	// The set of transmit channels for this antenna.
	TransmitChannels []RfEmitterDetailNewParamsAntennaTransmitChannel `json:"transmitChannels,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsAntenna) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsAntenna
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Feed associated with an RF Antenna.
type RfEmitterDetailNewParamsAntennaFeed struct {
	// Maximum frequency, in megahertz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// Minimum frequency, in megahertz.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// The feed name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The antenna feed linear/circular polarization (e.g. HORIZONTAL, VERTICAL,
	// LEFT_HAND_CIRCULAR, RIGHT_HAND_CIRCULAR).
	Polarization param.Opt[string] `json:"polarization,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsAntennaFeed) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsAntennaFeed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsAntennaFeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Receiver Channel associated with an RF Antenna.
type RfEmitterDetailNewParamsAntennaReceiverChannel struct {
	// The receiver bandwidth, in megahertz, must satisfy the constraint: minBandwidth
	// ≤ bandwidth ≤ maxBandwidth.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// The receive channel number.
	ChannelNum param.Opt[string] `json:"channelNum,omitzero"`
	// The receive channel device identifier.
	DeviceIdentifier param.Opt[string] `json:"deviceIdentifier,omitzero"`
	// Maximum frequency, in megahertz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// Minimum frequency, in megahertz.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// The maximum receiver bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MaxBandwidth param.Opt[float64] `json:"maxBandwidth,omitzero"`
	// The receiver bandwidth, in megahertz, must satisfy the constraint: minBandwidth
	// ≤ bandwidth ≤ maxBandwidth.
	MinBandwidth param.Opt[float64] `json:"minBandwidth,omitzero"`
	// Receiver sensitivity, in decibel-milliwatts.
	Sensitivity param.Opt[float64] `json:"sensitivity,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsAntennaReceiverChannel) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsAntennaReceiverChannel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsAntennaReceiverChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Transmit Channel associated with an RF Antenna.
//
// The property Power is required.
type RfEmitterDetailNewParamsAntennaTransmitChannel struct {
	// Transmit power, in watts.
	Power float64 `json:"power,required"`
	// The transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// The transmit channel number.
	ChannelNum param.Opt[string] `json:"channelNum,omitzero"`
	// The transmit channel device identifier.
	DeviceIdentifier param.Opt[string] `json:"deviceIdentifier,omitzero"`
	// The transmitter frequency, in megahertz, must satisfy the constraint: freqMin <=
	// freq <= freqMax.
	Freq param.Opt[float64] `json:"freq,omitzero"`
	// The maximum transmitter frequency, in megahertz, must satisfy the constraint:
	// freqMin ≤ freq ≤ freqMax.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// The minimum transmitter frequency, in megahertz, must satisfy the constraint:
	// freqMin ≤ freq ≤ freqMax.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// The hardware sample rate, in bits per second for this transmit channel.
	HardwareSampleRate param.Opt[int64] `json:"hardwareSampleRate,omitzero"`
	// The maximum transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MaxBandwidth param.Opt[float64] `json:"maxBandwidth,omitzero"`
	// Maximum gain, in decibels.
	MaxGain param.Opt[float64] `json:"maxGain,omitzero"`
	// The minimum transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MinBandwidth param.Opt[float64] `json:"minBandwidth,omitzero"`
	// Minimum gain, in decibels.
	MinGain param.Opt[float64] `json:"minGain,omitzero"`
	// The set of sample rates supported by this transmit channel, in bits per second.
	SampleRates []float64 `json:"sampleRates,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsAntennaTransmitChannel) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsAntennaTransmitChannel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsAntennaTransmitChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Power Offset associated with an RF Emitter Details.
type RfEmitterDetailNewParamsPowerOffset struct {
	// The RF frequency band (e.g. HF, VHF, P, UHF, L, S, C, X, KU, K, KA, V, W, MM).
	FrequencyBand param.Opt[string] `json:"frequencyBand,omitzero"`
	// Power offset, in decibels.
	PowerOffset param.Opt[float64] `json:"powerOffset,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsPowerOffset) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsPowerOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsPowerOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter SW Service associated with an RF Emitter Details.
type RfEmitterDetailNewParamsService struct {
	// The name for this software service.
	Name param.Opt[string] `json:"name,omitzero"`
	// The version for this software service.
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsService) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter TTP associated with an RF Emitter Details.
type RfEmitterDetailNewParamsTtp struct {
	// The name of the output signal.
	OutputSignalName param.Opt[string] `json:"outputSignalName,omitzero"`
	// The set of TTPs affected by this signal.
	TechniqueDefinitions []RfEmitterDetailNewParamsTtpTechniqueDefinition `json:"techniqueDefinitions,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsTtp) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsTtp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsTtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Definition associated with an RF Emitter TTP.
type RfEmitterDetailNewParamsTtpTechniqueDefinition struct {
	// The EW Emitter system technique name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The set of required/optional parameters for this technique.
	ParamDefinitions []RfEmitterDetailNewParamsTtpTechniqueDefinitionParamDefinition `json:"paramDefinitions,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsTtpTechniqueDefinition) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsTtpTechniqueDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsTtpTechniqueDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Parameter Definition associated with an RF Emitter
// Technique Definition.
type RfEmitterDetailNewParamsTtpTechniqueDefinitionParamDefinition struct {
	// Default parameter value used if not overridden in a SEW task definition.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Maximum allowable value for a numeric parameter.
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowable value for a numeric parameter.
	Min param.Opt[float64] `json:"min,omitzero"`
	// The name of the parameter.
	Name param.Opt[string] `json:"name,omitzero"`
	// A flag to specify that a parameter is optional.
	Optional param.Opt[bool] `json:"optional,omitzero"`
	// The type of parameter (e.g. STRING, DOUBLE, INT, LIST).
	Type param.Opt[string] `json:"type,omitzero"`
	// Units (degrees, seconds, decibels, etc.) for a numeric parameter.
	Units param.Opt[string] `json:"units,omitzero"`
	// Valid values for strictly defined parameters.
	ValidValues []string `json:"validValues,omitzero"`
	paramObj
}

func (r RfEmitterDetailNewParamsTtpTechniqueDefinitionParamDefinition) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailNewParamsTtpTechniqueDefinitionParamDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailNewParamsTtpTechniqueDefinitionParamDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterDetailUpdateParams struct {
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
	DataMode RfEmitterDetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent RF Emitter.
	IDRfEmitter string `json:"idRFEmitter,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate facility name for this RF Emitter.
	AlternateFacilityName param.Opt[string] `json:"alternateFacilityName,omitzero"`
	// Optional alternate name or alias for this RF Emitter.
	AltName param.Opt[string] `json:"altName,omitzero"`
	// Barrage noise bandwidth, in megahertz.
	BarrageNoiseBandwidth param.Opt[float64] `json:"barrageNoiseBandwidth,omitzero"`
	// The length of time, in seconds, for the RF Emitter built-in test to run to
	// completion.
	BitRunTime param.Opt[float64] `json:"bitRunTime,omitzero"`
	// Detailed description of the RF Emitter.
	Description param.Opt[string] `json:"description,omitzero"`
	// Designator of this RF Emitter.
	Designator param.Opt[string] `json:"designator,omitzero"`
	// Doppler noise value, in megahertz.
	DopplerNoise param.Opt[float64] `json:"dopplerNoise,omitzero"`
	// Digital Form Radio Memory instantaneous bandwidth in megahertz.
	DrfmInstantaneousBandwidth param.Opt[float64] `json:"drfmInstantaneousBandwidth,omitzero"`
	// Family of this RF Emitter type.
	Family param.Opt[string] `json:"family,omitzero"`
	// A fixed attenuation value to be set on the SRF Emitter HPA when commanding an
	// Electronic Attack/Techniques Tactics and Procedures task, in decibels.
	FixedAttenuation param.Opt[float64] `json:"fixedAttenuation,omitzero"`
	// Unique identifier of the organization which manufactured this RF Emitter.
	IDManufacturerOrg param.Opt[string] `json:"idManufacturerOrg,omitzero"`
	// Unique identifier of the location of the production facility for this RF
	// Emitter.
	IDProductionFacilityLocation param.Opt[string] `json:"idProductionFacilityLocation,omitzero"`
	// COCOM that has temporary responsibility for scheduling and management of the RF
	// Emitter (e.g. SPACEFOR-CENT, SPACEFOR-EURAF, SPACEFOR-INDOPAC, SPACEFOR-KOR,
	// SPACEFOR-STRATNORTH, SPACESOC, NONE).
	LoanedToCocom param.Opt[string] `json:"loanedToCocom,omitzero"`
	// Notes on the RF Emitter.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of bits.
	NumBits param.Opt[int64] `json:"numBits,omitzero"`
	// Number of channels.
	NumChannels param.Opt[int64] `json:"numChannels,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The length of time, in seconds, for the RF Emitter to prepare for a task,
	// including sufficient time to slew the antenna and configure the equipment.
	PrepTime param.Opt[float64] `json:"prepTime,omitzero"`
	// Primary COCOM that is responsible for scheduling and management of the RF
	// Emitter (e.g. SPACEFOR-CENT, SPACEFOR-EURAF, SPACEFOR-INDOPAC, SPACEFOR-KOR,
	// SPACEFOR-STRATNORTH, SPACESOC, NONE).
	PrimaryCocom param.Opt[string] `json:"primaryCocom,omitzero"`
	// Name of the production facility for this RF Emitter.
	ProductionFacilityName param.Opt[string] `json:"productionFacilityName,omitzero"`
	// Type or name of receiver.
	ReceiverType param.Opt[string] `json:"receiverType,omitzero"`
	// Secondary notes on the RF Emitter.
	SecondaryNotes param.Opt[string] `json:"secondaryNotes,omitzero"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. End sensitivity range, in
	// decibel-milliwatts.
	SystemSensitivityEnd param.Opt[float64] `json:"systemSensitivityEnd,omitzero"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. Start sensitivity range, in
	// decibel-milliwatts.
	SystemSensitivityStart param.Opt[float64] `json:"systemSensitivityStart,omitzero"`
	// An RF Amplifier associated with an RF Emitter Details.
	Amplifier RfEmitterDetailUpdateParamsAmplifier `json:"amplifier,omitzero"`
	// The set of antennas hosted on this EW Emitter system.
	Antennas []RfEmitterDetailUpdateParamsAntenna `json:"antennas,omitzero"`
	// A set of system/frequency band adjustments to the power offset commanded in an
	// EA/TTP task.
	PowerOffsets []RfEmitterDetailUpdateParamsPowerOffset `json:"powerOffsets,omitzero"`
	// The set of software services running on this EW Emitter system.
	Services []RfEmitterDetailUpdateParamsService `json:"services,omitzero"`
	// The set of EA/TTP techniques that are supported by this EW Emitter system.
	Ttps []RfEmitterDetailUpdateParamsTtp `json:"ttps,omitzero"`
	// Array of URLs containing additional information on this RF Emitter.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParams) UnmarshalJSON(data []byte) error {
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
type RfEmitterDetailUpdateParamsDataMode string

const (
	RfEmitterDetailUpdateParamsDataModeReal      RfEmitterDetailUpdateParamsDataMode = "REAL"
	RfEmitterDetailUpdateParamsDataModeTest      RfEmitterDetailUpdateParamsDataMode = "TEST"
	RfEmitterDetailUpdateParamsDataModeSimulated RfEmitterDetailUpdateParamsDataMode = "SIMULATED"
	RfEmitterDetailUpdateParamsDataModeExercise  RfEmitterDetailUpdateParamsDataMode = "EXERCISE"
)

// An RF Amplifier associated with an RF Emitter Details.
type RfEmitterDetailUpdateParamsAmplifier struct {
	// The device identifier of the amplifier.
	DeviceIdentifier param.Opt[string] `json:"deviceIdentifier,omitzero"`
	// The manufacturer of the amplifier.
	Manufacturer param.Opt[string] `json:"manufacturer,omitzero"`
	// The model name of the amplifier.
	ModelName param.Opt[string] `json:"modelName,omitzero"`
	// The amplifier power level, in watts.
	Power param.Opt[float64] `json:"power,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsAmplifier) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsAmplifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsAmplifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna associated with an RF Emitter Details.
type RfEmitterDetailUpdateParamsAntenna struct {
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter param.Opt[float64] `json:"antennaDiameter,omitzero"`
	// A flag to indicate whether the antenna points at a fixed azimuth/elevation.
	AzElFixed param.Opt[bool] `json:"azElFixed,omitzero"`
	// Antenna azimuth, in degrees clockwise from true North, for a fixed antenna.
	FixedAzimuth param.Opt[float64] `json:"fixedAzimuth,omitzero"`
	// Antenna elevation, in degrees, for a fixed antenna.
	FixedElevation param.Opt[float64] `json:"fixedElevation,omitzero"`
	// Maximum elevation, in degrees.
	MaxElevation param.Opt[float64] `json:"maxElevation,omitzero"`
	// Minimum elevation, in degrees.
	MinElevation param.Opt[float64] `json:"minElevation,omitzero"`
	// The name of the antenna.
	Name param.Opt[string] `json:"name,omitzero"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize,omitzero"`
	// The set of antenna feeds for this antenna.
	Feeds []RfEmitterDetailUpdateParamsAntennaFeed `json:"feeds,omitzero"`
	// Array of maximum azimuths, in degrees.
	MaxAzimuths []float64 `json:"maxAzimuths,omitzero"`
	// Array of minimum azimuths, in degrees.
	MinAzimuths []float64 `json:"minAzimuths,omitzero"`
	// The set of receiver channels for this antenna.
	ReceiverChannels []RfEmitterDetailUpdateParamsAntennaReceiverChannel `json:"receiverChannels,omitzero"`
	// The set of transmit channels for this antenna.
	TransmitChannels []RfEmitterDetailUpdateParamsAntennaTransmitChannel `json:"transmitChannels,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsAntenna) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsAntenna
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsAntenna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Feed associated with an RF Antenna.
type RfEmitterDetailUpdateParamsAntennaFeed struct {
	// Maximum frequency, in megahertz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// Minimum frequency, in megahertz.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// The feed name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The antenna feed linear/circular polarization (e.g. HORIZONTAL, VERTICAL,
	// LEFT_HAND_CIRCULAR, RIGHT_HAND_CIRCULAR).
	Polarization param.Opt[string] `json:"polarization,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsAntennaFeed) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsAntennaFeed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsAntennaFeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Receiver Channel associated with an RF Antenna.
type RfEmitterDetailUpdateParamsAntennaReceiverChannel struct {
	// The receiver bandwidth, in megahertz, must satisfy the constraint: minBandwidth
	// ≤ bandwidth ≤ maxBandwidth.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// The receive channel number.
	ChannelNum param.Opt[string] `json:"channelNum,omitzero"`
	// The receive channel device identifier.
	DeviceIdentifier param.Opt[string] `json:"deviceIdentifier,omitzero"`
	// Maximum frequency, in megahertz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// Minimum frequency, in megahertz.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// The maximum receiver bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MaxBandwidth param.Opt[float64] `json:"maxBandwidth,omitzero"`
	// The receiver bandwidth, in megahertz, must satisfy the constraint: minBandwidth
	// ≤ bandwidth ≤ maxBandwidth.
	MinBandwidth param.Opt[float64] `json:"minBandwidth,omitzero"`
	// Receiver sensitivity, in decibel-milliwatts.
	Sensitivity param.Opt[float64] `json:"sensitivity,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsAntennaReceiverChannel) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsAntennaReceiverChannel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsAntennaReceiverChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Antenna Transmit Channel associated with an RF Antenna.
//
// The property Power is required.
type RfEmitterDetailUpdateParamsAntennaTransmitChannel struct {
	// Transmit power, in watts.
	Power float64 `json:"power,required"`
	// The transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// The transmit channel number.
	ChannelNum param.Opt[string] `json:"channelNum,omitzero"`
	// The transmit channel device identifier.
	DeviceIdentifier param.Opt[string] `json:"deviceIdentifier,omitzero"`
	// The transmitter frequency, in megahertz, must satisfy the constraint: freqMin <=
	// freq <= freqMax.
	Freq param.Opt[float64] `json:"freq,omitzero"`
	// The maximum transmitter frequency, in megahertz, must satisfy the constraint:
	// freqMin ≤ freq ≤ freqMax.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// The minimum transmitter frequency, in megahertz, must satisfy the constraint:
	// freqMin ≤ freq ≤ freqMax.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// The hardware sample rate, in bits per second for this transmit channel.
	HardwareSampleRate param.Opt[int64] `json:"hardwareSampleRate,omitzero"`
	// The maximum transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MaxBandwidth param.Opt[float64] `json:"maxBandwidth,omitzero"`
	// Maximum gain, in decibels.
	MaxGain param.Opt[float64] `json:"maxGain,omitzero"`
	// The minimum transmitter bandwidth, in megahertz, must satisfy the constraint:
	// minBandwidth ≤ bandwidth ≤ maxBandwidth.
	MinBandwidth param.Opt[float64] `json:"minBandwidth,omitzero"`
	// Minimum gain, in decibels.
	MinGain param.Opt[float64] `json:"minGain,omitzero"`
	// The set of sample rates supported by this transmit channel, in bits per second.
	SampleRates []float64 `json:"sampleRates,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsAntennaTransmitChannel) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsAntennaTransmitChannel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsAntennaTransmitChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Power Offset associated with an RF Emitter Details.
type RfEmitterDetailUpdateParamsPowerOffset struct {
	// The RF frequency band (e.g. HF, VHF, P, UHF, L, S, C, X, KU, K, KA, V, W, MM).
	FrequencyBand param.Opt[string] `json:"frequencyBand,omitzero"`
	// Power offset, in decibels.
	PowerOffset param.Opt[float64] `json:"powerOffset,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsPowerOffset) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsPowerOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsPowerOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter SW Service associated with an RF Emitter Details.
type RfEmitterDetailUpdateParamsService struct {
	// The name for this software service.
	Name param.Opt[string] `json:"name,omitzero"`
	// The version for this software service.
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsService) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter TTP associated with an RF Emitter Details.
type RfEmitterDetailUpdateParamsTtp struct {
	// The name of the output signal.
	OutputSignalName param.Opt[string] `json:"outputSignalName,omitzero"`
	// The set of TTPs affected by this signal.
	TechniqueDefinitions []RfEmitterDetailUpdateParamsTtpTechniqueDefinition `json:"techniqueDefinitions,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsTtp) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsTtp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsTtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Definition associated with an RF Emitter TTP.
type RfEmitterDetailUpdateParamsTtpTechniqueDefinition struct {
	// The EW Emitter system technique name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The set of required/optional parameters for this technique.
	ParamDefinitions []RfEmitterDetailUpdateParamsTtpTechniqueDefinitionParamDefinition `json:"paramDefinitions,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsTtpTechniqueDefinition) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsTtpTechniqueDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsTtpTechniqueDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An RF Emitter Technique Parameter Definition associated with an RF Emitter
// Technique Definition.
type RfEmitterDetailUpdateParamsTtpTechniqueDefinitionParamDefinition struct {
	// Default parameter value used if not overridden in a SEW task definition.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Maximum allowable value for a numeric parameter.
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowable value for a numeric parameter.
	Min param.Opt[float64] `json:"min,omitzero"`
	// The name of the parameter.
	Name param.Opt[string] `json:"name,omitzero"`
	// A flag to specify that a parameter is optional.
	Optional param.Opt[bool] `json:"optional,omitzero"`
	// The type of parameter (e.g. STRING, DOUBLE, INT, LIST).
	Type param.Opt[string] `json:"type,omitzero"`
	// Units (degrees, seconds, decibels, etc.) for a numeric parameter.
	Units param.Opt[string] `json:"units,omitzero"`
	// Valid values for strictly defined parameters.
	ValidValues []string `json:"validValues,omitzero"`
	paramObj
}

func (r RfEmitterDetailUpdateParamsTtpTechniqueDefinitionParamDefinition) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterDetailUpdateParamsTtpTechniqueDefinitionParamDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterDetailUpdateParamsTtpTechniqueDefinitionParamDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterDetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterDetailListParams]'s query parameters as
// `url.Values`.
func (r RfEmitterDetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfEmitterDetailCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterDetailCountParams]'s query parameters as
// `url.Values`.
func (r RfEmitterDetailCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfEmitterDetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterDetailGetParams]'s query parameters as
// `url.Values`.
func (r RfEmitterDetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfEmitterDetailTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterDetailTupleParams]'s query parameters as
// `url.Values`.
func (r RfEmitterDetailTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
