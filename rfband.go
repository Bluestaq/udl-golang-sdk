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

// Service operation to take a single RFBand as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RfBandService) New(ctx context.Context, body RfBandNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfband"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an RFBand. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *RfBandService) Update(ctx context.Context, id string, body RfBandUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append(r.Options[:], opts...)
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

// Service operation to delete an RFBand specified by the passed ID path parameter.
// A specific role is required to perform this service operation. Please contact
// the UDL team for assistance.
func (r *RfBandService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/rfband/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single RFBand by its unique ID passed as a path
// parameter.
func (r *RfBandService) Get(ctx context.Context, id string, query RfBandGetParams, opts ...option.RequestOption) (res *RfBandGetResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *RfBandService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfband/queryhelp"
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
func (r *RfBandService) Tuple(ctx context.Context, query RfBandTupleParams, opts ...option.RequestOption) (res *[]RfBandTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	// RF Band frequency range bandwidth in Mhz.
	Bandwidth float64 `json:"bandwidth"`
	// Angle between the half-power (-3 dB) points of the main lobe of the antenna, in
	// degrees.
	Beamwidth float64 `json:"beamwidth"`
	// Center frequency of RF frequency range, if applicable, in Mhz.
	CenterFreq float64 `json:"centerFreq"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// RF Range edge gain, in dBi.
	EdgeGain float64 `json:"edgeGain"`
	// EIRP is defined as the RMS power input in decibel watts required to a lossless
	// half-wave dipole antenna to give the same maximum power density far from the
	// antenna as the actual transmitter. It is equal to the power input to the
	// transmitter's antenna multiplied by the antenna gain relative to a half-wave
	// dipole. Effective radiated power and effective isotropic radiated power both
	// measure the amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Eirp float64 `json:"eirp"`
	// Effective Radiated Power (ERP) is the total power in decibel watts radiated by
	// an actual antenna relative to a half-wave dipole rather than a theoretical
	// isotropic antenna. A half-wave dipole has a gain of 2.15 dB compared to an
	// isotropic antenna. EIRP(dB) = ERP (dB)+2.15 dB or EIRP (W) = 1.64\*ERP(W).
	// Effective radiated power and effective isotropic radiated power both measure the
	// amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Erp float64 `json:"erp"`
	// End/maximum of transmit RF frequency range, if applicable, in Mhz.
	FreqMax float64 `json:"freqMax"`
	// Start/minimum of transmit RF frequency range, if applicable, in Mhz.
	FreqMin float64 `json:"freqMin"`
	// RF Band mode (e.g. TX, RX).
	//
	// Any of "TX", "RX".
	Mode RfBandListResponseMode `json:"mode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// RF Range maximum gain, in dBi.
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
	// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
	//
	// Any of "COMM", "TTC", "OPS", "OTHER".
	Purpose RfBandListResponsePurpose `json:"purpose"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDEntity              resp.Field
		Name                  resp.Field
		Source                resp.Field
		ID                    resp.Field
		Band                  resp.Field
		Bandwidth             resp.Field
		Beamwidth             resp.Field
		CenterFreq            resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EdgeGain              resp.Field
		Eirp                  resp.Field
		Erp                   resp.Field
		FreqMax               resp.Field
		FreqMin               resp.Field
		Mode                  resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		PeakGain              resp.Field
		Polarization          resp.Field
		Purpose               resp.Field
		ExtraFields           map[string]resp.Field
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
// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
type RfBandListResponsePurpose string

const (
	RfBandListResponsePurposeComm  RfBandListResponsePurpose = "COMM"
	RfBandListResponsePurposeTtc   RfBandListResponsePurpose = "TTC"
	RfBandListResponsePurposeOps   RfBandListResponsePurpose = "OPS"
	RfBandListResponsePurposeOther RfBandListResponsePurpose = "OTHER"
)

// Details on a particular Radio Frequency (RF) band, also known as a carrier,
// which may be in use by any type of Entity for communications or operations.
type RfBandGetResponse struct {
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
	DataMode RfBandGetResponseDataMode `json:"dataMode,required"`
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
	// RF Band frequency range bandwidth in Mhz.
	Bandwidth float64 `json:"bandwidth"`
	// Angle between the half-power (-3 dB) points of the main lobe of the antenna, in
	// degrees.
	Beamwidth float64 `json:"beamwidth"`
	// Center frequency of RF frequency range, if applicable, in Mhz.
	CenterFreq float64 `json:"centerFreq"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// RF Range edge gain, in dBi.
	EdgeGain float64 `json:"edgeGain"`
	// EIRP is defined as the RMS power input in decibel watts required to a lossless
	// half-wave dipole antenna to give the same maximum power density far from the
	// antenna as the actual transmitter. It is equal to the power input to the
	// transmitter's antenna multiplied by the antenna gain relative to a half-wave
	// dipole. Effective radiated power and effective isotropic radiated power both
	// measure the amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Eirp float64 `json:"eirp"`
	// Effective Radiated Power (ERP) is the total power in decibel watts radiated by
	// an actual antenna relative to a half-wave dipole rather than a theoretical
	// isotropic antenna. A half-wave dipole has a gain of 2.15 dB compared to an
	// isotropic antenna. EIRP(dB) = ERP (dB)+2.15 dB or EIRP (W) = 1.64\*ERP(W).
	// Effective radiated power and effective isotropic radiated power both measure the
	// amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Erp float64 `json:"erp"`
	// End/maximum of transmit RF frequency range, if applicable, in Mhz.
	FreqMax float64 `json:"freqMax"`
	// Start/minimum of transmit RF frequency range, if applicable, in Mhz.
	FreqMin float64 `json:"freqMin"`
	// RF Band mode (e.g. TX, RX).
	//
	// Any of "TX", "RX".
	Mode RfBandGetResponseMode `json:"mode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// RF Range maximum gain, in dBi.
	PeakGain float64 `json:"peakGain"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	Polarization RfBandGetResponsePolarization `json:"polarization"`
	// Purpose or use of the RF Band -- COMM = communications, TTC =
	// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
	//
	// Any of "COMM", "TTC", "OPS", "OTHER".
	Purpose RfBandGetResponsePurpose `json:"purpose"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDEntity              resp.Field
		Name                  resp.Field
		Source                resp.Field
		ID                    resp.Field
		Band                  resp.Field
		Bandwidth             resp.Field
		Beamwidth             resp.Field
		CenterFreq            resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EdgeGain              resp.Field
		Eirp                  resp.Field
		Erp                   resp.Field
		FreqMax               resp.Field
		FreqMin               resp.Field
		Mode                  resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		PeakGain              resp.Field
		Polarization          resp.Field
		Purpose               resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfBandGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RfBandGetResponse) UnmarshalJSON(data []byte) error {
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
type RfBandGetResponseDataMode string

const (
	RfBandGetResponseDataModeReal      RfBandGetResponseDataMode = "REAL"
	RfBandGetResponseDataModeTest      RfBandGetResponseDataMode = "TEST"
	RfBandGetResponseDataModeSimulated RfBandGetResponseDataMode = "SIMULATED"
	RfBandGetResponseDataModeExercise  RfBandGetResponseDataMode = "EXERCISE"
)

// RF Band mode (e.g. TX, RX).
type RfBandGetResponseMode string

const (
	RfBandGetResponseModeTx RfBandGetResponseMode = "TX"
	RfBandGetResponseModeRx RfBandGetResponseMode = "RX"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
// surface.
type RfBandGetResponsePolarization string

const (
	RfBandGetResponsePolarizationH RfBandGetResponsePolarization = "H"
	RfBandGetResponsePolarizationV RfBandGetResponsePolarization = "V"
	RfBandGetResponsePolarizationR RfBandGetResponsePolarization = "R"
	RfBandGetResponsePolarizationL RfBandGetResponsePolarization = "L"
)

// Purpose or use of the RF Band -- COMM = communications, TTC =
// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
type RfBandGetResponsePurpose string

const (
	RfBandGetResponsePurposeComm  RfBandGetResponsePurpose = "COMM"
	RfBandGetResponsePurposeTtc   RfBandGetResponsePurpose = "TTC"
	RfBandGetResponsePurposeOps   RfBandGetResponsePurpose = "OPS"
	RfBandGetResponsePurposeOther RfBandGetResponsePurpose = "OTHER"
)

// Details on a particular Radio Frequency (RF) band, also known as a carrier,
// which may be in use by any type of Entity for communications or operations.
type RfBandTupleResponse struct {
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
	DataMode RfBandTupleResponseDataMode `json:"dataMode,required"`
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
	// RF Band frequency range bandwidth in Mhz.
	Bandwidth float64 `json:"bandwidth"`
	// Angle between the half-power (-3 dB) points of the main lobe of the antenna, in
	// degrees.
	Beamwidth float64 `json:"beamwidth"`
	// Center frequency of RF frequency range, if applicable, in Mhz.
	CenterFreq float64 `json:"centerFreq"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// RF Range edge gain, in dBi.
	EdgeGain float64 `json:"edgeGain"`
	// EIRP is defined as the RMS power input in decibel watts required to a lossless
	// half-wave dipole antenna to give the same maximum power density far from the
	// antenna as the actual transmitter. It is equal to the power input to the
	// transmitter's antenna multiplied by the antenna gain relative to a half-wave
	// dipole. Effective radiated power and effective isotropic radiated power both
	// measure the amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Eirp float64 `json:"eirp"`
	// Effective Radiated Power (ERP) is the total power in decibel watts radiated by
	// an actual antenna relative to a half-wave dipole rather than a theoretical
	// isotropic antenna. A half-wave dipole has a gain of 2.15 dB compared to an
	// isotropic antenna. EIRP(dB) = ERP (dB)+2.15 dB or EIRP (W) = 1.64\*ERP(W).
	// Effective radiated power and effective isotropic radiated power both measure the
	// amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Erp float64 `json:"erp"`
	// End/maximum of transmit RF frequency range, if applicable, in Mhz.
	FreqMax float64 `json:"freqMax"`
	// Start/minimum of transmit RF frequency range, if applicable, in Mhz.
	FreqMin float64 `json:"freqMin"`
	// RF Band mode (e.g. TX, RX).
	//
	// Any of "TX", "RX".
	Mode RfBandTupleResponseMode `json:"mode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// RF Range maximum gain, in dBi.
	PeakGain float64 `json:"peakGain"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	Polarization RfBandTupleResponsePolarization `json:"polarization"`
	// Purpose or use of the RF Band -- COMM = communications, TTC =
	// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
	//
	// Any of "COMM", "TTC", "OPS", "OTHER".
	Purpose RfBandTupleResponsePurpose `json:"purpose"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDEntity              resp.Field
		Name                  resp.Field
		Source                resp.Field
		ID                    resp.Field
		Band                  resp.Field
		Bandwidth             resp.Field
		Beamwidth             resp.Field
		CenterFreq            resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EdgeGain              resp.Field
		Eirp                  resp.Field
		Erp                   resp.Field
		FreqMax               resp.Field
		FreqMin               resp.Field
		Mode                  resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		PeakGain              resp.Field
		Polarization          resp.Field
		Purpose               resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfBandTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *RfBandTupleResponse) UnmarshalJSON(data []byte) error {
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
type RfBandTupleResponseDataMode string

const (
	RfBandTupleResponseDataModeReal      RfBandTupleResponseDataMode = "REAL"
	RfBandTupleResponseDataModeTest      RfBandTupleResponseDataMode = "TEST"
	RfBandTupleResponseDataModeSimulated RfBandTupleResponseDataMode = "SIMULATED"
	RfBandTupleResponseDataModeExercise  RfBandTupleResponseDataMode = "EXERCISE"
)

// RF Band mode (e.g. TX, RX).
type RfBandTupleResponseMode string

const (
	RfBandTupleResponseModeTx RfBandTupleResponseMode = "TX"
	RfBandTupleResponseModeRx RfBandTupleResponseMode = "RX"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
// surface.
type RfBandTupleResponsePolarization string

const (
	RfBandTupleResponsePolarizationH RfBandTupleResponsePolarization = "H"
	RfBandTupleResponsePolarizationV RfBandTupleResponsePolarization = "V"
	RfBandTupleResponsePolarizationR RfBandTupleResponsePolarization = "R"
	RfBandTupleResponsePolarizationL RfBandTupleResponsePolarization = "L"
)

// Purpose or use of the RF Band -- COMM = communications, TTC =
// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
type RfBandTupleResponsePurpose string

const (
	RfBandTupleResponsePurposeComm  RfBandTupleResponsePurpose = "COMM"
	RfBandTupleResponsePurposeTtc   RfBandTupleResponsePurpose = "TTC"
	RfBandTupleResponsePurposeOps   RfBandTupleResponsePurpose = "OPS"
	RfBandTupleResponsePurposeOther RfBandTupleResponsePurpose = "OTHER"
)

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
	// RF Band frequency range bandwidth in Mhz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Angle between the half-power (-3 dB) points of the main lobe of the antenna, in
	// degrees.
	Beamwidth param.Opt[float64] `json:"beamwidth,omitzero"`
	// Center frequency of RF frequency range, if applicable, in Mhz.
	CenterFreq param.Opt[float64] `json:"centerFreq,omitzero"`
	// RF Range edge gain, in dBi.
	EdgeGain param.Opt[float64] `json:"edgeGain,omitzero"`
	// EIRP is defined as the RMS power input in decibel watts required to a lossless
	// half-wave dipole antenna to give the same maximum power density far from the
	// antenna as the actual transmitter. It is equal to the power input to the
	// transmitter's antenna multiplied by the antenna gain relative to a half-wave
	// dipole. Effective radiated power and effective isotropic radiated power both
	// measure the amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Eirp param.Opt[float64] `json:"eirp,omitzero"`
	// Effective Radiated Power (ERP) is the total power in decibel watts radiated by
	// an actual antenna relative to a half-wave dipole rather than a theoretical
	// isotropic antenna. A half-wave dipole has a gain of 2.15 dB compared to an
	// isotropic antenna. EIRP(dB) = ERP (dB)+2.15 dB or EIRP (W) = 1.64\*ERP(W).
	// Effective radiated power and effective isotropic radiated power both measure the
	// amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Erp param.Opt[float64] `json:"erp,omitzero"`
	// End/maximum of transmit RF frequency range, if applicable, in Mhz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// Start/minimum of transmit RF frequency range, if applicable, in Mhz.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// RF Range maximum gain, in dBi.
	PeakGain param.Opt[float64] `json:"peakGain,omitzero"`
	// RF Band mode (e.g. TX, RX).
	//
	// Any of "TX", "RX".
	Mode RfBandNewParamsMode `json:"mode,omitzero"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	Polarization RfBandNewParamsPolarization `json:"polarization,omitzero"`
	// Purpose or use of the RF Band -- COMM = communications, TTC =
	// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
	//
	// Any of "COMM", "TTC", "OPS", "OTHER".
	Purpose RfBandNewParamsPurpose `json:"purpose,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfBandNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RfBandNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RfBandNewParams
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
// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
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
	// RF Band frequency range bandwidth in Mhz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Angle between the half-power (-3 dB) points of the main lobe of the antenna, in
	// degrees.
	Beamwidth param.Opt[float64] `json:"beamwidth,omitzero"`
	// Center frequency of RF frequency range, if applicable, in Mhz.
	CenterFreq param.Opt[float64] `json:"centerFreq,omitzero"`
	// RF Range edge gain, in dBi.
	EdgeGain param.Opt[float64] `json:"edgeGain,omitzero"`
	// EIRP is defined as the RMS power input in decibel watts required to a lossless
	// half-wave dipole antenna to give the same maximum power density far from the
	// antenna as the actual transmitter. It is equal to the power input to the
	// transmitter's antenna multiplied by the antenna gain relative to a half-wave
	// dipole. Effective radiated power and effective isotropic radiated power both
	// measure the amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Eirp param.Opt[float64] `json:"eirp,omitzero"`
	// Effective Radiated Power (ERP) is the total power in decibel watts radiated by
	// an actual antenna relative to a half-wave dipole rather than a theoretical
	// isotropic antenna. A half-wave dipole has a gain of 2.15 dB compared to an
	// isotropic antenna. EIRP(dB) = ERP (dB)+2.15 dB or EIRP (W) = 1.64\*ERP(W).
	// Effective radiated power and effective isotropic radiated power both measure the
	// amount of power a radio transmitter and antenna (or other source of
	// electromagnetic waves) radiates in a specific direction: in the direction of
	// maximum signal strength (the "main lobe") of its radiation pattern.
	Erp param.Opt[float64] `json:"erp,omitzero"`
	// End/maximum of transmit RF frequency range, if applicable, in Mhz.
	FreqMax param.Opt[float64] `json:"freqMax,omitzero"`
	// Start/minimum of transmit RF frequency range, if applicable, in Mhz.
	FreqMin param.Opt[float64] `json:"freqMin,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// RF Range maximum gain, in dBi.
	PeakGain param.Opt[float64] `json:"peakGain,omitzero"`
	// RF Band mode (e.g. TX, RX).
	//
	// Any of "TX", "RX".
	Mode RfBandUpdateParamsMode `json:"mode,omitzero"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the Earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the Earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	Polarization RfBandUpdateParamsPolarization `json:"polarization,omitzero"`
	// Purpose or use of the RF Band -- COMM = communications, TTC =
	// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
	//
	// Any of "COMM", "TTC", "OPS", "OTHER".
	Purpose RfBandUpdateParamsPurpose `json:"purpose,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfBandUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RfBandUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RfBandUpdateParams
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
// Telemetry/Tracking/Control, OPS = Operations, OTHER = Other).
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfBandListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfBandCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfBandGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfBandTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RfBandTupleParams]'s query parameters as `url.Values`.
func (r RfBandTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
