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

// RfBandService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRfBandService] method instead.
type RfBandService struct {
	Options []option.RequestOption
}

// NewRfBandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRfBandService(opts ...option.RequestOption) (r RfBandService) {
	r = RfBandService{}
	r.Options = opts
	return
}

// Service operation to take a single RFBand record as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *RfBandService) New(ctx context.Context, body RfBandNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/rfband"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single RFBand record. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *RfBandService) Update(ctx context.Context, id string, body RfBandUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfband/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *RfBandService) List(ctx context.Context, query RfBandListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[RfBandListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/rfband"
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
func (r *RfBandService) ListAutoPaging(ctx context.Context, query RfBandListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[RfBandListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a RFBand record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RfBandService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfband/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *RfBandService) Count(ctx context.Context, query RfBandCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/rfband/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single RFBand record by its unique ID passed as a
// path parameter.
func (r *RfBandService) Get(ctx context.Context, id string, query RfBandGetParams, opts ...option.RequestOption) (res *shared.RfBandFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfband/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *RfBandService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *RfBandQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/rfband/queryhelp"
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
func (r *RfBandService) Tuple(ctx context.Context, query RfBandTupleParams, opts ...option.RequestOption) (res *[]shared.RfBandFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/rfband/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details on a particular Radio Frequency (RF) band, also known as a carrier,
// which may be in use by any type of Entity for communications or operations.
type RfBandListResponse struct {
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
	DataMode RfBandListResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the parent Entity which uses this band.
	IDEntity string `json:"idEntity,required"`
	// RF Band name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band string `json:"band"`
	// RF Band frequency range bandwidth in megahertz.
	Bandwidth float64 `json:"bandwidth"`
	// Array of frequency range bandwidth settings, in megahertz for this RFBand. If
	// this array is specified then it must be the same size as the frequencySettings
	// array. A null value may be used for one or more of the frequencies in the
	// frequencySettings array if there is no corresponding value for a given
	// frequency.
	BandwidthSettings []float64 `json:"bandwidthSettings"`
	// Angle between the half-power (-3 dB) points of the main lobe of the antenna, in
	// degrees.
	Beamwidth float64 `json:"beamwidth"`
	// Array of beamwidth settings, in degrees for this RFBand. If this array is
	// specified then it must be the same size as the frequencySettings array. A null
	// value may be used for one or more of the frequencies in the frequencySettings
	// array if there is no corresponding value for a given frequency.
	BeamwidthSettings []float64 `json:"beamwidthSettings"`
	// Center frequency of RF frequency range, if applicable, in megahertz.
	CenterFreq float64 `json:"centerFreq"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array of delay settings, in seconds for this RFBand. If this array is specified
	// then it must be the same size as the frequencySettings array. A null value may
	// be used for one or more of the frequencies in the frequencySettings array if
	// there is no corresponding value for a given frequency.
	DelaySettings []float64 `json:"delaySettings"`
	// RF Range edge gain, in decibel relative to isotrope.
	EdgeGain float64 `json:"edgeGain"`
	// EIRP is defined as the RMS power input in decibel watts required to a lossless
	// half-wave dipole antenna to give the same maximum power density far from the
	// antenna as the actual transmitter. It is equal to the power input to the
	// transmitter's antenna multiplied by the antenna gain relative to a half-wave
	// dipole. Effective radiated power and effective isotropic radiated power both
	// measure the amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the main lobe) of its radiation pattern.
	Eirp float64 `json:"eirp"`
	// Effective Radiated Power (ERP) is the total power in decibel watts radiated by
	// an actual antenna relative to a half-wave dipole rather than a theoretical
	// isotropic antenna. A half-wave dipole has a gain of 2.15 dB compared to an
	// isotropic antenna. EIRP(dB) = ERP (dB)+2.15 dB or EIRP (W) = 1.64\*ERP(W).
	// Effective radiated power and effective isotropic radiated power both measure the
	// amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the main lobe) of its radiation pattern.
	Erp float64 `json:"erp"`
	// End/maximum of transmit RF frequency range, if applicable, in megahertz.
	FreqMax float64 `json:"freqMax"`
	// Start/minimum of transmit RF frequency range, if applicable, in megahertz.
	FreqMin float64 `json:"freqMin"`
	// Array of frequency settings, in megahertz for this RFBand. This array and the
	// settings arrays must match in size.
	FrequencySettings []float64 `json:"frequencySettings"`
	// Array of gain settings, in decibels for this RFBand. If this array is specified
	// then it must be the same size as the frequencySettings array. A null value may
	// be used for one or more of the frequencies in the frequencySettings array if
	// there is no corresponding value for a given frequency.
	GainSettings []float64 `json:"gainSettings"`
	// RF Band mode (e.g. TX, RX).
	//
	// Any of "TX", "RX".
	Mode RfBandListResponseMode `json:"mode"`
	// Array of signal noise settings, in decibels for this RFBand. If this array is
	// specified then it must be the same size as the frequencySettings array. A null
	// value may be used for one or more of the frequencies in the frequencySettings
	// array if there is no corresponding value for a given frequency.
	NoiseSettings []float64 `json:"noiseSettings"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// RF Range maximum gain, in decibel relative to isotrope.
	PeakGain float64 `json:"peakGain"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	Polarization RfBandListResponsePolarization `json:"polarization"`
	// Purpose or use of the RF Band -- COMM = communications, TTC =
	// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other.
	//
	// Any of "COMM", "TTC", "OPS", "OTHER".
	Purpose RfBandListResponsePurpose `json:"purpose"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDEntity              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Band                  respjson.Field
		Bandwidth             respjson.Field
		BandwidthSettings     respjson.Field
		Beamwidth             respjson.Field
		BeamwidthSettings     respjson.Field
		CenterFreq            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DelaySettings         respjson.Field
		EdgeGain              respjson.Field
		Eirp                  respjson.Field
		Erp                   respjson.Field
		FreqMax               respjson.Field
		FreqMin               respjson.Field
		FrequencySettings     respjson.Field
		GainSettings          respjson.Field
		Mode                  respjson.Field
		NoiseSettings         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PeakGain              respjson.Field
		Polarization          respjson.Field
		Purpose               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfBandListResponse) RawJSON() string { return r.JSON.raw }
func (r *RfBandListResponse) UnmarshalJSON(data []byte) error {
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
type RfBandListResponseDataMode string

const (
	RfBandListResponseDataModeReal      RfBandListResponseDataMode = "REAL"
	RfBandListResponseDataModeTest      RfBandListResponseDataMode = "TEST"
	RfBandListResponseDataModeSimulated RfBandListResponseDataMode = "SIMULATED"
	RfBandListResponseDataModeExercise  RfBandListResponseDataMode = "EXERCISE"
)

// RF Band mode (e.g. TX, RX).
type RfBandListResponseMode string

const (
	RfBandListResponseModeTx RfBandListResponseMode = "TX"
	RfBandListResponseModeRx RfBandListResponseMode = "RX"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
// surface.
type RfBandListResponsePolarization string

const (
	RfBandListResponsePolarizationH RfBandListResponsePolarization = "H"
	RfBandListResponsePolarizationV RfBandListResponsePolarization = "V"
	RfBandListResponsePolarizationR RfBandListResponsePolarization = "R"
	RfBandListResponsePolarizationL RfBandListResponsePolarization = "L"
)

// Purpose or use of the RF Band -- COMM = communications, TTC =
// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other.
type RfBandListResponsePurpose string

const (
	RfBandListResponsePurposeComm  RfBandListResponsePurpose = "COMM"
	RfBandListResponsePurposeTtc   RfBandListResponsePurpose = "TTC"
	RfBandListResponsePurposeOps   RfBandListResponsePurpose = "OPS"
	RfBandListResponsePurposeOther RfBandListResponsePurpose = "OTHER"
)

type RfBandQueryhelpResponse struct {
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
func (r RfBandQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *RfBandQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfBandNewParams struct {
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
	DataMode RfBandNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent Entity which uses this band.
	IDEntity string `json:"idEntity,required"`
	// RF Band name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band param.Opt[string] `json:"band,omitzero"`
	// RF Band frequency range bandwidth in megahertz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Angle between the half-power (-3 dB) points of the main lobe of the antenna, in
	// degrees.
	Beamwidth param.Opt[float64] `json:"beamwidth,omitzero"`
	// Center frequency of RF frequency range, if applicable, in megahertz.
	CenterFreq param.Opt[float64] `json:"centerFreq,omitzero"`
	// RF Range edge gain, in decibel relative to isotrope.
	EdgeGain param.Opt[float64] `json:"edgeGain,omitzero"`
	// EIRP is defined as the RMS power input in decibel watts required to a lossless
	// half-wave dipole antenna to give the same maximum power density far from the
	// antenna as the actual transmitter. It is equal to the power input to the
	// transmitter's antenna multiplied by the antenna gain relative to a half-wave
	// dipole. Effective radiated power and effective isotropic radiated power both
	// measure the amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the main lobe) of its radiation pattern.
	Eirp param.Opt[float64] `json:"eirp,omitzero"`
	// Effective Radiated Power (ERP) is the total power in decibel watts radiated by
	// an actual antenna relative to a half-wave dipole rather than a theoretical
	// isotropic antenna. A half-wave dipole has a gain of 2.15 dB compared to an
	// isotropic antenna. EIRP(dB) = ERP (dB)+2.15 dB or EIRP (W) = 1.64\*ERP(W).
	// Effective radiated power and effective isotropic radiated power both measure the
	// amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the main lobe) of its radiation pattern.
	Erp param.Opt[float64] `json:"erp,omitzero"`
	// End/maximum of transmit RF frequency range, if applicable, in megahertz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// Start/minimum of transmit RF frequency range, if applicable, in megahertz.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// RF Range maximum gain, in decibel relative to isotrope.
	PeakGain param.Opt[float64] `json:"peakGain,omitzero"`
	// Array of frequency range bandwidth settings, in megahertz for this RFBand. If
	// this array is specified then it must be the same size as the frequencySettings
	// array. A null value may be used for one or more of the frequencies in the
	// frequencySettings array if there is no corresponding value for a given
	// frequency.
	BandwidthSettings []float64 `json:"bandwidthSettings,omitzero"`
	// Array of beamwidth settings, in degrees for this RFBand. If this array is
	// specified then it must be the same size as the frequencySettings array. A null
	// value may be used for one or more of the frequencies in the frequencySettings
	// array if there is no corresponding value for a given frequency.
	BeamwidthSettings []float64 `json:"beamwidthSettings,omitzero"`
	// Array of delay settings, in seconds for this RFBand. If this array is specified
	// then it must be the same size as the frequencySettings array. A null value may
	// be used for one or more of the frequencies in the frequencySettings array if
	// there is no corresponding value for a given frequency.
	DelaySettings []float64 `json:"delaySettings,omitzero"`
	// Array of frequency settings, in megahertz for this RFBand. This array and the
	// settings arrays must match in size.
	FrequencySettings []float64 `json:"frequencySettings,omitzero"`
	// Array of gain settings, in decibels for this RFBand. If this array is specified
	// then it must be the same size as the frequencySettings array. A null value may
	// be used for one or more of the frequencies in the frequencySettings array if
	// there is no corresponding value for a given frequency.
	GainSettings []float64 `json:"gainSettings,omitzero"`
	// RF Band mode (e.g. TX, RX).
	//
	// Any of "TX", "RX".
	Mode RfBandNewParamsMode `json:"mode,omitzero"`
	// Array of signal noise settings, in decibels for this RFBand. If this array is
	// specified then it must be the same size as the frequencySettings array. A null
	// value may be used for one or more of the frequencies in the frequencySettings
	// array if there is no corresponding value for a given frequency.
	NoiseSettings []float64 `json:"noiseSettings,omitzero"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	Polarization RfBandNewParamsPolarization `json:"polarization,omitzero"`
	// Purpose or use of the RF Band -- COMM = communications, TTC =
	// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other.
	//
	// Any of "COMM", "TTC", "OPS", "OTHER".
	Purpose RfBandNewParamsPurpose `json:"purpose,omitzero"`
	paramObj
}

func (r RfBandNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RfBandNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfBandNewParams) UnmarshalJSON(data []byte) error {
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
type RfBandNewParamsDataMode string

const (
	RfBandNewParamsDataModeReal      RfBandNewParamsDataMode = "REAL"
	RfBandNewParamsDataModeTest      RfBandNewParamsDataMode = "TEST"
	RfBandNewParamsDataModeSimulated RfBandNewParamsDataMode = "SIMULATED"
	RfBandNewParamsDataModeExercise  RfBandNewParamsDataMode = "EXERCISE"
)

// RF Band mode (e.g. TX, RX).
type RfBandNewParamsMode string

const (
	RfBandNewParamsModeTx RfBandNewParamsMode = "TX"
	RfBandNewParamsModeRx RfBandNewParamsMode = "RX"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
// surface.
type RfBandNewParamsPolarization string

const (
	RfBandNewParamsPolarizationH RfBandNewParamsPolarization = "H"
	RfBandNewParamsPolarizationV RfBandNewParamsPolarization = "V"
	RfBandNewParamsPolarizationR RfBandNewParamsPolarization = "R"
	RfBandNewParamsPolarizationL RfBandNewParamsPolarization = "L"
)

// Purpose or use of the RF Band -- COMM = communications, TTC =
// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other.
type RfBandNewParamsPurpose string

const (
	RfBandNewParamsPurposeComm  RfBandNewParamsPurpose = "COMM"
	RfBandNewParamsPurposeTtc   RfBandNewParamsPurpose = "TTC"
	RfBandNewParamsPurposeOps   RfBandNewParamsPurpose = "OPS"
	RfBandNewParamsPurposeOther RfBandNewParamsPurpose = "OTHER"
)

type RfBandUpdateParams struct {
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
	DataMode RfBandUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent Entity which uses this band.
	IDEntity string `json:"idEntity,required"`
	// RF Band name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band param.Opt[string] `json:"band,omitzero"`
	// RF Band frequency range bandwidth in megahertz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Angle between the half-power (-3 dB) points of the main lobe of the antenna, in
	// degrees.
	Beamwidth param.Opt[float64] `json:"beamwidth,omitzero"`
	// Center frequency of RF frequency range, if applicable, in megahertz.
	CenterFreq param.Opt[float64] `json:"centerFreq,omitzero"`
	// RF Range edge gain, in decibel relative to isotrope.
	EdgeGain param.Opt[float64] `json:"edgeGain,omitzero"`
	// EIRP is defined as the RMS power input in decibel watts required to a lossless
	// half-wave dipole antenna to give the same maximum power density far from the
	// antenna as the actual transmitter. It is equal to the power input to the
	// transmitter's antenna multiplied by the antenna gain relative to a half-wave
	// dipole. Effective radiated power and effective isotropic radiated power both
	// measure the amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the main lobe) of its radiation pattern.
	Eirp param.Opt[float64] `json:"eirp,omitzero"`
	// Effective Radiated Power (ERP) is the total power in decibel watts radiated by
	// an actual antenna relative to a half-wave dipole rather than a theoretical
	// isotropic antenna. A half-wave dipole has a gain of 2.15 dB compared to an
	// isotropic antenna. EIRP(dB) = ERP (dB)+2.15 dB or EIRP (W) = 1.64\*ERP(W).
	// Effective radiated power and effective isotropic radiated power both measure the
	// amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the main lobe) of its radiation pattern.
	Erp param.Opt[float64] `json:"erp,omitzero"`
	// End/maximum of transmit RF frequency range, if applicable, in megahertz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// Start/minimum of transmit RF frequency range, if applicable, in megahertz.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// RF Range maximum gain, in decibel relative to isotrope.
	PeakGain param.Opt[float64] `json:"peakGain,omitzero"`
	// Array of frequency range bandwidth settings, in megahertz for this RFBand. If
	// this array is specified then it must be the same size as the frequencySettings
	// array. A null value may be used for one or more of the frequencies in the
	// frequencySettings array if there is no corresponding value for a given
	// frequency.
	BandwidthSettings []float64 `json:"bandwidthSettings,omitzero"`
	// Array of beamwidth settings, in degrees for this RFBand. If this array is
	// specified then it must be the same size as the frequencySettings array. A null
	// value may be used for one or more of the frequencies in the frequencySettings
	// array if there is no corresponding value for a given frequency.
	BeamwidthSettings []float64 `json:"beamwidthSettings,omitzero"`
	// Array of delay settings, in seconds for this RFBand. If this array is specified
	// then it must be the same size as the frequencySettings array. A null value may
	// be used for one or more of the frequencies in the frequencySettings array if
	// there is no corresponding value for a given frequency.
	DelaySettings []float64 `json:"delaySettings,omitzero"`
	// Array of frequency settings, in megahertz for this RFBand. This array and the
	// settings arrays must match in size.
	FrequencySettings []float64 `json:"frequencySettings,omitzero"`
	// Array of gain settings, in decibels for this RFBand. If this array is specified
	// then it must be the same size as the frequencySettings array. A null value may
	// be used for one or more of the frequencies in the frequencySettings array if
	// there is no corresponding value for a given frequency.
	GainSettings []float64 `json:"gainSettings,omitzero"`
	// RF Band mode (e.g. TX, RX).
	//
	// Any of "TX", "RX".
	Mode RfBandUpdateParamsMode `json:"mode,omitzero"`
	// Array of signal noise settings, in decibels for this RFBand. If this array is
	// specified then it must be the same size as the frequencySettings array. A null
	// value may be used for one or more of the frequencies in the frequencySettings
	// array if there is no corresponding value for a given frequency.
	NoiseSettings []float64 `json:"noiseSettings,omitzero"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	Polarization RfBandUpdateParamsPolarization `json:"polarization,omitzero"`
	// Purpose or use of the RF Band -- COMM = communications, TTC =
	// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other.
	//
	// Any of "COMM", "TTC", "OPS", "OTHER".
	Purpose RfBandUpdateParamsPurpose `json:"purpose,omitzero"`
	paramObj
}

func (r RfBandUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RfBandUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfBandUpdateParams) UnmarshalJSON(data []byte) error {
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
type RfBandUpdateParamsDataMode string

const (
	RfBandUpdateParamsDataModeReal      RfBandUpdateParamsDataMode = "REAL"
	RfBandUpdateParamsDataModeTest      RfBandUpdateParamsDataMode = "TEST"
	RfBandUpdateParamsDataModeSimulated RfBandUpdateParamsDataMode = "SIMULATED"
	RfBandUpdateParamsDataModeExercise  RfBandUpdateParamsDataMode = "EXERCISE"
)

// RF Band mode (e.g. TX, RX).
type RfBandUpdateParamsMode string

const (
	RfBandUpdateParamsModeTx RfBandUpdateParamsMode = "TX"
	RfBandUpdateParamsModeRx RfBandUpdateParamsMode = "RX"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
// surface.
type RfBandUpdateParamsPolarization string

const (
	RfBandUpdateParamsPolarizationH RfBandUpdateParamsPolarization = "H"
	RfBandUpdateParamsPolarizationV RfBandUpdateParamsPolarization = "V"
	RfBandUpdateParamsPolarizationR RfBandUpdateParamsPolarization = "R"
	RfBandUpdateParamsPolarizationL RfBandUpdateParamsPolarization = "L"
)

// Purpose or use of the RF Band -- COMM = communications, TTC =
// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other.
type RfBandUpdateParamsPurpose string

const (
	RfBandUpdateParamsPurposeComm  RfBandUpdateParamsPurpose = "COMM"
	RfBandUpdateParamsPurposeTtc   RfBandUpdateParamsPurpose = "TTC"
	RfBandUpdateParamsPurposeOps   RfBandUpdateParamsPurpose = "OPS"
	RfBandUpdateParamsPurposeOther RfBandUpdateParamsPurpose = "OTHER"
)

type RfBandListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfBandListParams]'s query parameters as `url.Values`.
func (r RfBandListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfBandCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfBandCountParams]'s query parameters as `url.Values`.
func (r RfBandCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfBandGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfBandGetParams]'s query parameters as `url.Values`.
func (r RfBandGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfBandTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfBandTupleParams]'s query parameters as `url.Values`.
func (r RfBandTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
