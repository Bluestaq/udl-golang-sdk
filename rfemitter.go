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

// RfEmitterService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRfEmitterService] method instead.
type RfEmitterService struct {
	Options []option.RequestOption
}

// NewRfEmitterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRfEmitterService(opts ...option.RequestOption) (r RfEmitterService) {
	r = RfEmitterService{}
	r.Options = opts
	return
}

// Service operation to take a single RFEmitter as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RfEmitterService) New(ctx context.Context, body RfEmitterNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfemitter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an RFEmitter. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *RfEmitterService) Update(ctx context.Context, id string, body RfEmitterUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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

// Service operation to delete an RFEmitter specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RfEmitterService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/rfemitter/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single RFEmitter by its unique ID passed as a path
// parameter.
func (r *RfEmitterService) Get(ctx context.Context, id string, query RfEmitterGetParams, opts ...option.RequestOption) (res *RfEmitterGetResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	// ID of the parent entity for this rfemitter.
	IDEntity string `json:"idEntity"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
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
		IDEntity              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
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
	Entity EntityFull `json:"entity"`
	// ID of the parent entity for this rfemitter.
	IDEntity string `json:"idEntity"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only details for this RFEmitter.
	RfEmitterDetails []RfEmitterGetResponseRfEmitterDetail `json:"rfEmitterDetails"`
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
		IDEntity              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RfEmitterDetails      respjson.Field
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking        respjson.Field
		DataMode                     respjson.Field
		IDRfEmitter                  respjson.Field
		Source                       respjson.Field
		ID                           respjson.Field
		AlternateFacilityName        respjson.Field
		AltName                      respjson.Field
		AntennaDiameter              respjson.Field
		AntennaSize                  respjson.Field
		BarrageNoiseBandwidth        respjson.Field
		CreatedAt                    respjson.Field
		CreatedBy                    respjson.Field
		Description                  respjson.Field
		Designator                   respjson.Field
		DopplerNoise                 respjson.Field
		DrfmInstantaneousBandwidth   respjson.Field
		Family                       respjson.Field
		ManufacturerOrg              respjson.Field
		ManufacturerOrgID            respjson.Field
		Notes                        respjson.Field
		NumBits                      respjson.Field
		NumChannels                  respjson.Field
		Origin                       respjson.Field
		OrigNetwork                  respjson.Field
		ProductionFacilityLocation   respjson.Field
		ProductionFacilityLocationID respjson.Field
		ProductionFacilityName       respjson.Field
		ReceiverBandwidth            respjson.Field
		ReceiverSensitivity          respjson.Field
		ReceiverType                 respjson.Field
		SecondaryNotes               respjson.Field
		SystemSensitivityEnd         respjson.Field
		SystemSensitivityStart       respjson.Field
		TransmitPower                respjson.Field
		TransmitterBandwidth         respjson.Field
		TransmitterFrequency         respjson.Field
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

type RfEmitterQueryhelpResponse struct {
	AodrSupported         bool                                  `json:"aodrSupported"`
	ClassificationMarking string                                `json:"classificationMarking"`
	Description           string                                `json:"description"`
	HistorySupported      bool                                  `json:"historySupported"`
	Name                  string                                `json:"name"`
	Parameters            []RfEmitterQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                              `json:"requiredRoles"`
	RestSupported         bool                                  `json:"restSupported"`
	SortSupported         bool                                  `json:"sortSupported"`
	TypeName              string                                `json:"typeName"`
	Uri                   string                                `json:"uri"`
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

type RfEmitterQueryhelpResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
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
	Entity EntityFull `json:"entity"`
	// ID of the parent entity for this rfemitter.
	IDEntity string `json:"idEntity"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Read-only details for this RFEmitter.
	RfEmitterDetails []RfEmitterTupleResponseRfEmitterDetail `json:"rfEmitterDetails"`
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
		IDEntity              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RfEmitterDetails      respjson.Field
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking        respjson.Field
		DataMode                     respjson.Field
		IDRfEmitter                  respjson.Field
		Source                       respjson.Field
		ID                           respjson.Field
		AlternateFacilityName        respjson.Field
		AltName                      respjson.Field
		AntennaDiameter              respjson.Field
		AntennaSize                  respjson.Field
		BarrageNoiseBandwidth        respjson.Field
		CreatedAt                    respjson.Field
		CreatedBy                    respjson.Field
		Description                  respjson.Field
		Designator                   respjson.Field
		DopplerNoise                 respjson.Field
		DrfmInstantaneousBandwidth   respjson.Field
		Family                       respjson.Field
		ManufacturerOrg              respjson.Field
		ManufacturerOrgID            respjson.Field
		Notes                        respjson.Field
		NumBits                      respjson.Field
		NumChannels                  respjson.Field
		Origin                       respjson.Field
		OrigNetwork                  respjson.Field
		ProductionFacilityLocation   respjson.Field
		ProductionFacilityLocationID respjson.Field
		ProductionFacilityName       respjson.Field
		ReceiverBandwidth            respjson.Field
		ReceiverSensitivity          respjson.Field
		ReceiverType                 respjson.Field
		SecondaryNotes               respjson.Field
		SystemSensitivityEnd         respjson.Field
		SystemSensitivityStart       respjson.Field
		TransmitPower                respjson.Field
		TransmitterBandwidth         respjson.Field
		TransmitterFrequency         respjson.Field
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
	// ID of the parent entity for this rfemitter.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
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
	// ID of the parent entity for this rfemitter.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
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
