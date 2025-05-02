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

// RfemitterdetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRfemitterdetailService] method instead.
type RfemitterdetailService struct {
	Options []option.RequestOption
}

// NewRfemitterdetailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRfemitterdetailService(opts ...option.RequestOption) (r RfemitterdetailService) {
	r = RfemitterdetailService{}
	r.Options = opts
	return
}

// Service operation to take a single RFEmitterDetails as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *RfemitterdetailService) New(ctx context.Context, body RfemitterdetailNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfemitterdetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an RFEmitterDetails. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *RfemitterdetailService) Update(ctx context.Context, id string, body RfemitterdetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
func (r *RfemitterdetailService) List(ctx context.Context, query RfemitterdetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[RfemitterdetailListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *RfemitterdetailService) ListAutoPaging(ctx context.Context, query RfemitterdetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[RfemitterdetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an RFEmitterDetails specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RfemitterdetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
func (r *RfemitterdetailService) Count(ctx context.Context, query RfemitterdetailCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/rfemitterdetails/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single RFEmitterDetails by its unique ID passed as a
// path parameter.
func (r *RfemitterdetailService) Get(ctx context.Context, id string, query RfemitterdetailGetParams, opts ...option.RequestOption) (res *RfemitterdetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *RfemitterdetailService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfemitterdetails/queryhelp"
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
func (r *RfemitterdetailService) Tuple(ctx context.Context, query RfemitterdetailTupleParams, opts ...option.RequestOption) (res *[]RfemitterdetailTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/rfemitterdetails/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details for a particular RF Emitter, collected by a particular source. An RF
// Emitter may have multiple details records from various sources.
type RfemitterdetailListResponse struct {
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
	DataMode RfemitterdetailListResponseDataMode `json:"dataMode,required"`
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
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter float64 `json:"antennaDiameter"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize"`
	// Barrage noise bandwidth in Mhz.
	BarrageNoiseBandwidth float64 `json:"barrageNoiseBandwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Detailed description of the RF Emitter.
	Description string `json:"description"`
	// Designator of this RF Emitter.
	Designator string `json:"designator"`
	// Doppler noise value in Mhz.
	DopplerNoise float64 `json:"dopplerNoise"`
	// Digital Form Radio Memory instantaneous bandwidth in Mhz.
	DrfmInstantaneousBandwidth float64 `json:"drfmInstantaneousBandwidth"`
	// Family of this RF Emitter type.
	Family string `json:"family"`
	// Unique identifier of the organization which manufactures this RF Emitter.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
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
	// Unique identifier of the location of the production facility for this RF
	// Emitter.
	ProductionFacilityLocationID string `json:"productionFacilityLocationId"`
	// Name of the production facility for this RF Emitter.
	ProductionFacilityName string `json:"productionFacilityName"`
	// Receiver bandwidth in Mhz.
	ReceiverBandwidth float64 `json:"receiverBandwidth"`
	// Receiver sensitivity in dBm.
	ReceiverSensitivity float64 `json:"receiverSensitivity"`
	// Type or name of receiver.
	ReceiverType string `json:"receiverType"`
	// Secondary notes on the RF Emitter.
	SecondaryNotes string `json:"secondaryNotes"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. End sensitivity range, in dBm.
	SystemSensitivityEnd float64 `json:"systemSensitivityEnd"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. Start sensitivity range, in dBm.
	SystemSensitivityStart float64 `json:"systemSensitivityStart"`
	// Transmit power in Watts.
	TransmitPower float64 `json:"transmitPower"`
	// Transmitter bandwidth in Mhz.
	TransmitterBandwidth float64 `json:"transmitterBandwidth"`
	// Transmitter frequency in Mhz.
	TransmitterFrequency float64 `json:"transmitterFrequency"`
	// Array of URLs containing additional information on this RF Emitter.
	URLs []string `json:"urls"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking        resp.Field
		DataMode                     resp.Field
		IDRfEmitter                  resp.Field
		Source                       resp.Field
		ID                           resp.Field
		AlternateFacilityName        resp.Field
		AltName                      resp.Field
		AntennaDiameter              resp.Field
		AntennaSize                  resp.Field
		BarrageNoiseBandwidth        resp.Field
		CreatedAt                    resp.Field
		CreatedBy                    resp.Field
		Description                  resp.Field
		Designator                   resp.Field
		DopplerNoise                 resp.Field
		DrfmInstantaneousBandwidth   resp.Field
		Family                       resp.Field
		ManufacturerOrgID            resp.Field
		Notes                        resp.Field
		NumBits                      resp.Field
		NumChannels                  resp.Field
		Origin                       resp.Field
		OrigNetwork                  resp.Field
		ProductionFacilityLocationID resp.Field
		ProductionFacilityName       resp.Field
		ReceiverBandwidth            resp.Field
		ReceiverSensitivity          resp.Field
		ReceiverType                 resp.Field
		SecondaryNotes               resp.Field
		SystemSensitivityEnd         resp.Field
		SystemSensitivityStart       resp.Field
		TransmitPower                resp.Field
		TransmitterBandwidth         resp.Field
		TransmitterFrequency         resp.Field
		URLs                         resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfemitterdetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *RfemitterdetailListResponse) UnmarshalJSON(data []byte) error {
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
type RfemitterdetailListResponseDataMode string

const (
	RfemitterdetailListResponseDataModeReal      RfemitterdetailListResponseDataMode = "REAL"
	RfemitterdetailListResponseDataModeTest      RfemitterdetailListResponseDataMode = "TEST"
	RfemitterdetailListResponseDataModeSimulated RfemitterdetailListResponseDataMode = "SIMULATED"
	RfemitterdetailListResponseDataModeExercise  RfemitterdetailListResponseDataMode = "EXERCISE"
)

// Details for a particular RF Emitter, collected by a particular source. An RF
// Emitter may have multiple details records from various sources.
type RfemitterdetailGetResponse struct {
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
	DataMode RfemitterdetailGetResponseDataMode `json:"dataMode,required"`
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
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter float64 `json:"antennaDiameter"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize"`
	// Barrage noise bandwidth in Mhz.
	BarrageNoiseBandwidth float64 `json:"barrageNoiseBandwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Detailed description of the RF Emitter.
	Description string `json:"description"`
	// Designator of this RF Emitter.
	Designator string `json:"designator"`
	// Doppler noise value in Mhz.
	DopplerNoise float64 `json:"dopplerNoise"`
	// Digital Form Radio Memory instantaneous bandwidth in Mhz.
	DrfmInstantaneousBandwidth float64 `json:"drfmInstantaneousBandwidth"`
	// Family of this RF Emitter type.
	Family string `json:"family"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	ManufacturerOrg OrganizationFull `json:"manufacturerOrg"`
	// Unique identifier of the organization which manufactures this RF Emitter.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
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
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	ProductionFacilityLocation LocationFull `json:"productionFacilityLocation"`
	// Unique identifier of the location of the production facility for this RF
	// Emitter.
	ProductionFacilityLocationID string `json:"productionFacilityLocationId"`
	// Name of the production facility for this RF Emitter.
	ProductionFacilityName string `json:"productionFacilityName"`
	// Receiver bandwidth in Mhz.
	ReceiverBandwidth float64 `json:"receiverBandwidth"`
	// Receiver sensitivity in dBm.
	ReceiverSensitivity float64 `json:"receiverSensitivity"`
	// Type or name of receiver.
	ReceiverType string `json:"receiverType"`
	// Secondary notes on the RF Emitter.
	SecondaryNotes string `json:"secondaryNotes"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. End sensitivity range, in dBm.
	SystemSensitivityEnd float64 `json:"systemSensitivityEnd"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. Start sensitivity range, in dBm.
	SystemSensitivityStart float64 `json:"systemSensitivityStart"`
	// Transmit power in Watts.
	TransmitPower float64 `json:"transmitPower"`
	// Transmitter bandwidth in Mhz.
	TransmitterBandwidth float64 `json:"transmitterBandwidth"`
	// Transmitter frequency in Mhz.
	TransmitterFrequency float64 `json:"transmitterFrequency"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Array of URLs containing additional information on this RF Emitter.
	URLs []string `json:"urls"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking        resp.Field
		DataMode                     resp.Field
		IDRfEmitter                  resp.Field
		Source                       resp.Field
		ID                           resp.Field
		AlternateFacilityName        resp.Field
		AltName                      resp.Field
		AntennaDiameter              resp.Field
		AntennaSize                  resp.Field
		BarrageNoiseBandwidth        resp.Field
		CreatedAt                    resp.Field
		CreatedBy                    resp.Field
		Description                  resp.Field
		Designator                   resp.Field
		DopplerNoise                 resp.Field
		DrfmInstantaneousBandwidth   resp.Field
		Family                       resp.Field
		ManufacturerOrg              resp.Field
		ManufacturerOrgID            resp.Field
		Notes                        resp.Field
		NumBits                      resp.Field
		NumChannels                  resp.Field
		Origin                       resp.Field
		OrigNetwork                  resp.Field
		ProductionFacilityLocation   resp.Field
		ProductionFacilityLocationID resp.Field
		ProductionFacilityName       resp.Field
		ReceiverBandwidth            resp.Field
		ReceiverSensitivity          resp.Field
		ReceiverType                 resp.Field
		SecondaryNotes               resp.Field
		SystemSensitivityEnd         resp.Field
		SystemSensitivityStart       resp.Field
		TransmitPower                resp.Field
		TransmitterBandwidth         resp.Field
		TransmitterFrequency         resp.Field
		UpdatedAt                    resp.Field
		UpdatedBy                    resp.Field
		URLs                         resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfemitterdetailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RfemitterdetailGetResponse) UnmarshalJSON(data []byte) error {
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
type RfemitterdetailGetResponseDataMode string

const (
	RfemitterdetailGetResponseDataModeReal      RfemitterdetailGetResponseDataMode = "REAL"
	RfemitterdetailGetResponseDataModeTest      RfemitterdetailGetResponseDataMode = "TEST"
	RfemitterdetailGetResponseDataModeSimulated RfemitterdetailGetResponseDataMode = "SIMULATED"
	RfemitterdetailGetResponseDataModeExercise  RfemitterdetailGetResponseDataMode = "EXERCISE"
)

// Details for a particular RF Emitter, collected by a particular source. An RF
// Emitter may have multiple details records from various sources.
type RfemitterdetailTupleResponse struct {
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
	DataMode RfemitterdetailTupleResponseDataMode `json:"dataMode,required"`
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
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter float64 `json:"antennaDiameter"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize"`
	// Barrage noise bandwidth in Mhz.
	BarrageNoiseBandwidth float64 `json:"barrageNoiseBandwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Detailed description of the RF Emitter.
	Description string `json:"description"`
	// Designator of this RF Emitter.
	Designator string `json:"designator"`
	// Doppler noise value in Mhz.
	DopplerNoise float64 `json:"dopplerNoise"`
	// Digital Form Radio Memory instantaneous bandwidth in Mhz.
	DrfmInstantaneousBandwidth float64 `json:"drfmInstantaneousBandwidth"`
	// Family of this RF Emitter type.
	Family string `json:"family"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	ManufacturerOrg OrganizationFull `json:"manufacturerOrg"`
	// Unique identifier of the organization which manufactures this RF Emitter.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
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
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	ProductionFacilityLocation LocationFull `json:"productionFacilityLocation"`
	// Unique identifier of the location of the production facility for this RF
	// Emitter.
	ProductionFacilityLocationID string `json:"productionFacilityLocationId"`
	// Name of the production facility for this RF Emitter.
	ProductionFacilityName string `json:"productionFacilityName"`
	// Receiver bandwidth in Mhz.
	ReceiverBandwidth float64 `json:"receiverBandwidth"`
	// Receiver sensitivity in dBm.
	ReceiverSensitivity float64 `json:"receiverSensitivity"`
	// Type or name of receiver.
	ReceiverType string `json:"receiverType"`
	// Secondary notes on the RF Emitter.
	SecondaryNotes string `json:"secondaryNotes"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. End sensitivity range, in dBm.
	SystemSensitivityEnd float64 `json:"systemSensitivityEnd"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. Start sensitivity range, in dBm.
	SystemSensitivityStart float64 `json:"systemSensitivityStart"`
	// Transmit power in Watts.
	TransmitPower float64 `json:"transmitPower"`
	// Transmitter bandwidth in Mhz.
	TransmitterBandwidth float64 `json:"transmitterBandwidth"`
	// Transmitter frequency in Mhz.
	TransmitterFrequency float64 `json:"transmitterFrequency"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Array of URLs containing additional information on this RF Emitter.
	URLs []string `json:"urls"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking        resp.Field
		DataMode                     resp.Field
		IDRfEmitter                  resp.Field
		Source                       resp.Field
		ID                           resp.Field
		AlternateFacilityName        resp.Field
		AltName                      resp.Field
		AntennaDiameter              resp.Field
		AntennaSize                  resp.Field
		BarrageNoiseBandwidth        resp.Field
		CreatedAt                    resp.Field
		CreatedBy                    resp.Field
		Description                  resp.Field
		Designator                   resp.Field
		DopplerNoise                 resp.Field
		DrfmInstantaneousBandwidth   resp.Field
		Family                       resp.Field
		ManufacturerOrg              resp.Field
		ManufacturerOrgID            resp.Field
		Notes                        resp.Field
		NumBits                      resp.Field
		NumChannels                  resp.Field
		Origin                       resp.Field
		OrigNetwork                  resp.Field
		ProductionFacilityLocation   resp.Field
		ProductionFacilityLocationID resp.Field
		ProductionFacilityName       resp.Field
		ReceiverBandwidth            resp.Field
		ReceiverSensitivity          resp.Field
		ReceiverType                 resp.Field
		SecondaryNotes               resp.Field
		SystemSensitivityEnd         resp.Field
		SystemSensitivityStart       resp.Field
		TransmitPower                resp.Field
		TransmitterBandwidth         resp.Field
		TransmitterFrequency         resp.Field
		UpdatedAt                    resp.Field
		UpdatedBy                    resp.Field
		URLs                         resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfemitterdetailTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *RfemitterdetailTupleResponse) UnmarshalJSON(data []byte) error {
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
type RfemitterdetailTupleResponseDataMode string

const (
	RfemitterdetailTupleResponseDataModeReal      RfemitterdetailTupleResponseDataMode = "REAL"
	RfemitterdetailTupleResponseDataModeTest      RfemitterdetailTupleResponseDataMode = "TEST"
	RfemitterdetailTupleResponseDataModeSimulated RfemitterdetailTupleResponseDataMode = "SIMULATED"
	RfemitterdetailTupleResponseDataModeExercise  RfemitterdetailTupleResponseDataMode = "EXERCISE"
)

type RfemitterdetailNewParams struct {
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
	DataMode RfemitterdetailNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter param.Opt[float64] `json:"antennaDiameter,omitzero"`
	// Barrage noise bandwidth in Mhz.
	BarrageNoiseBandwidth param.Opt[float64] `json:"barrageNoiseBandwidth,omitzero"`
	// Detailed description of the RF Emitter.
	Description param.Opt[string] `json:"description,omitzero"`
	// Designator of this RF Emitter.
	Designator param.Opt[string] `json:"designator,omitzero"`
	// Doppler noise value in Mhz.
	DopplerNoise param.Opt[float64] `json:"dopplerNoise,omitzero"`
	// Digital Form Radio Memory instantaneous bandwidth in Mhz.
	DrfmInstantaneousBandwidth param.Opt[float64] `json:"drfmInstantaneousBandwidth,omitzero"`
	// Family of this RF Emitter type.
	Family param.Opt[string] `json:"family,omitzero"`
	// Unique identifier of the organization which manufactures this RF Emitter.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
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
	// Unique identifier of the location of the production facility for this RF
	// Emitter.
	ProductionFacilityLocationID param.Opt[string] `json:"productionFacilityLocationId,omitzero"`
	// Name of the production facility for this RF Emitter.
	ProductionFacilityName param.Opt[string] `json:"productionFacilityName,omitzero"`
	// Receiver bandwidth in Mhz.
	ReceiverBandwidth param.Opt[float64] `json:"receiverBandwidth,omitzero"`
	// Receiver sensitivity in dBm.
	ReceiverSensitivity param.Opt[float64] `json:"receiverSensitivity,omitzero"`
	// Type or name of receiver.
	ReceiverType param.Opt[string] `json:"receiverType,omitzero"`
	// Secondary notes on the RF Emitter.
	SecondaryNotes param.Opt[string] `json:"secondaryNotes,omitzero"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. End sensitivity range, in dBm.
	SystemSensitivityEnd param.Opt[float64] `json:"systemSensitivityEnd,omitzero"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. Start sensitivity range, in dBm.
	SystemSensitivityStart param.Opt[float64] `json:"systemSensitivityStart,omitzero"`
	// Transmit power in Watts.
	TransmitPower param.Opt[float64] `json:"transmitPower,omitzero"`
	// Transmitter bandwidth in Mhz.
	TransmitterBandwidth param.Opt[float64] `json:"transmitterBandwidth,omitzero"`
	// Transmitter frequency in Mhz.
	TransmitterFrequency param.Opt[float64] `json:"transmitterFrequency,omitzero"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize,omitzero"`
	// Array of URLs containing additional information on this RF Emitter.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfemitterdetailNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RfemitterdetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RfemitterdetailNewParams
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
type RfemitterdetailNewParamsDataMode string

const (
	RfemitterdetailNewParamsDataModeReal      RfemitterdetailNewParamsDataMode = "REAL"
	RfemitterdetailNewParamsDataModeTest      RfemitterdetailNewParamsDataMode = "TEST"
	RfemitterdetailNewParamsDataModeSimulated RfemitterdetailNewParamsDataMode = "SIMULATED"
	RfemitterdetailNewParamsDataModeExercise  RfemitterdetailNewParamsDataMode = "EXERCISE"
)

type RfemitterdetailUpdateParams struct {
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
	DataMode RfemitterdetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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
	// For parabolic/dish antennas, the diameter of the antenna in meters.
	AntennaDiameter param.Opt[float64] `json:"antennaDiameter,omitzero"`
	// Barrage noise bandwidth in Mhz.
	BarrageNoiseBandwidth param.Opt[float64] `json:"barrageNoiseBandwidth,omitzero"`
	// Detailed description of the RF Emitter.
	Description param.Opt[string] `json:"description,omitzero"`
	// Designator of this RF Emitter.
	Designator param.Opt[string] `json:"designator,omitzero"`
	// Doppler noise value in Mhz.
	DopplerNoise param.Opt[float64] `json:"dopplerNoise,omitzero"`
	// Digital Form Radio Memory instantaneous bandwidth in Mhz.
	DrfmInstantaneousBandwidth param.Opt[float64] `json:"drfmInstantaneousBandwidth,omitzero"`
	// Family of this RF Emitter type.
	Family param.Opt[string] `json:"family,omitzero"`
	// Unique identifier of the organization which manufactures this RF Emitter.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
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
	// Unique identifier of the location of the production facility for this RF
	// Emitter.
	ProductionFacilityLocationID param.Opt[string] `json:"productionFacilityLocationId,omitzero"`
	// Name of the production facility for this RF Emitter.
	ProductionFacilityName param.Opt[string] `json:"productionFacilityName,omitzero"`
	// Receiver bandwidth in Mhz.
	ReceiverBandwidth param.Opt[float64] `json:"receiverBandwidth,omitzero"`
	// Receiver sensitivity in dBm.
	ReceiverSensitivity param.Opt[float64] `json:"receiverSensitivity,omitzero"`
	// Type or name of receiver.
	ReceiverType param.Opt[string] `json:"receiverType,omitzero"`
	// Secondary notes on the RF Emitter.
	SecondaryNotes param.Opt[string] `json:"secondaryNotes,omitzero"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. End sensitivity range, in dBm.
	SystemSensitivityEnd param.Opt[float64] `json:"systemSensitivityEnd,omitzero"`
	// Receiver sensitivity is the lowest power level at which the receiver can detect
	// an RF signal and demodulate data. Sensitivity is purely a receiver specification
	// and is independent of the transmitter. Start sensitivity range, in dBm.
	SystemSensitivityStart param.Opt[float64] `json:"systemSensitivityStart,omitzero"`
	// Transmit power in Watts.
	TransmitPower param.Opt[float64] `json:"transmitPower,omitzero"`
	// Transmitter bandwidth in Mhz.
	TransmitterBandwidth param.Opt[float64] `json:"transmitterBandwidth,omitzero"`
	// Transmitter frequency in Mhz.
	TransmitterFrequency param.Opt[float64] `json:"transmitterFrequency,omitzero"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	AntennaSize []float64 `json:"antennaSize,omitzero"`
	// Array of URLs containing additional information on this RF Emitter.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfemitterdetailUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RfemitterdetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RfemitterdetailUpdateParams
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
type RfemitterdetailUpdateParamsDataMode string

const (
	RfemitterdetailUpdateParamsDataModeReal      RfemitterdetailUpdateParamsDataMode = "REAL"
	RfemitterdetailUpdateParamsDataModeTest      RfemitterdetailUpdateParamsDataMode = "TEST"
	RfemitterdetailUpdateParamsDataModeSimulated RfemitterdetailUpdateParamsDataMode = "SIMULATED"
	RfemitterdetailUpdateParamsDataModeExercise  RfemitterdetailUpdateParamsDataMode = "EXERCISE"
)

type RfemitterdetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfemitterdetailListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RfemitterdetailListParams]'s query parameters as
// `url.Values`.
func (r RfemitterdetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfemitterdetailCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfemitterdetailCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RfemitterdetailCountParams]'s query parameters as
// `url.Values`.
func (r RfemitterdetailCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfemitterdetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RfemitterdetailGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RfemitterdetailGetParams]'s query parameters as
// `url.Values`.
func (r RfemitterdetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfemitterdetailTupleParams struct {
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
func (f RfemitterdetailTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RfemitterdetailTupleParams]'s query parameters as
// `url.Values`.
func (r RfemitterdetailTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
