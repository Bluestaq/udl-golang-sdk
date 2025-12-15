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

// LaseremitterService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaseremitterService] method instead.
type LaseremitterService struct {
	Options []option.RequestOption
	Staging LaseremitterStagingService
}

// NewLaseremitterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaseremitterService(opts ...option.RequestOption) (r LaseremitterService) {
	r = LaseremitterService{}
	r.Options = opts
	r.Staging = NewLaseremitterStagingService(opts...)
	return
}

// Service operation to take a single LaserEmitter record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *LaseremitterService) New(ctx context.Context, body LaseremitterNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/laseremitter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single LaserEmitter record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *LaseremitterService) Update(ctx context.Context, id string, body LaseremitterUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/laseremitter/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaseremitterService) List(ctx context.Context, query LaseremitterListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaseremitterListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/laseremitter"
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
func (r *LaseremitterService) ListAutoPaging(ctx context.Context, query LaseremitterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaseremitterListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a LaserEmitter record specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LaseremitterService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/laseremitter/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LaseremitterService) Count(ctx context.Context, query LaseremitterCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/laseremitter/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single LaserEmitter record by its unique ID passed as
// a path parameter.
func (r *LaseremitterService) Get(ctx context.Context, id string, query LaseremitterGetParams, opts ...option.RequestOption) (res *LaseremitterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/laseremitter/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *LaseremitterService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *LaseremitterQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/laseremitter/queryhelp"
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
func (r *LaseremitterService) Tuple(ctx context.Context, query LaseremitterTupleParams, opts ...option.RequestOption) (res *[]LaseremitterTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/laseremitter/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of a laser beam emitter/director which represents a source
// of directed energy.
type LaseremitterListResponse struct {
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
	DataMode LaseremitterListResponseDataMode `json:"dataMode,required"`
	// The name of this laser within the laser system or facility.
	LaserName string `json:"laserName,required"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType string `json:"laserType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Maximum possible laser-to-space atmospheric transmission (a value between zero
	// and one, usually assumed to be unity for all lasers except for lasers with
	// wavelengths heavily absorbed by the atmosphere).
	AtmosphericTransmission float64 `json:"atmosphericTransmission"`
	// The beam director aperture average (equivalent CW) output power is the total
	// laser average power that is transmitted away from the final exit aperture of the
	// optical system. For a CW laser, this is equivalent to the laser device CW power
	// multiplied by the throughput of the optical system, including the beam director
	// telescope. For a pulsed laser, this is equivalent to the laser device energy per
	// pulse multiplied by the pulse repetition frequency (PRF) multiplied by the
	// throughput of the optical system including the beam director telescope. Measured
	// in Watts.
	BeamOutputPower float64 `json:"beamOutputPower"`
	// The beam waist (radius) of this laser at the exit aperture, in centimeters.
	BeamWaist float64 `json:"beamWaist"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Minimum possible divergence half-angle of this laser, referenced from the
	// optical axis to the 1/e point of the beam's far field irradiance profile,
	// measured in microradians.
	Divergence float64 `json:"divergence"`
	// Unique identifier of the parent entity. idEntity is required for PUT.
	IDEntity string `json:"idEntity"`
	// A Boolean indicating whether or not this laser emitter is operational. False
	// indicates that the laser specified in this record is no longer available and can
	// be considered defunct. This field is true by default.
	IsOperational bool `json:"isOperational"`
	// The maximum time that the laser would be "on" for a given test(s) for laser
	// activity, in seconds.
	MaxDuration float64 `json:"maxDuration"`
	// The maximum possible focus range of this laser, measured in kilometers.
	MaxFocusRange float64 `json:"maxFocusRange"`
	// The minimum possible focus range of this laser, measured in kilometers.
	MinFocusRange float64 `json:"minFocusRange"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The amount of energy in each laser pulse exiting from the beam
	// director/telescope, measured in Joules.
	PulseFluence float64 `json:"pulseFluence"`
	// The instantaneous single pulse peak power exiting from the laser device,
	// measured in Watts.
	PulsePeakPower float64 `json:"pulsePeakPower"`
	// The pulse repetition frequency of this laser, measured in kilohertz.
	PulseRepFreq float64 `json:"pulseRepFreq"`
	// The pulse shape (waveform) of this laser, e.g. "RECTANGULAR," "SAWTOOTH,"
	// "GAUSSIAN," etc.
	PulseShape string `json:"pulseShape"`
	// The laser device pulse duration, measured in seconds.
	PulseWidth float64 `json:"pulseWidth"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The center wavelength of this laser, in micrometers.
	Wavelength float64 `json:"wavelength"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking   respjson.Field
		DataMode                respjson.Field
		LaserName               respjson.Field
		LaserType               respjson.Field
		Source                  respjson.Field
		ID                      respjson.Field
		AtmosphericTransmission respjson.Field
		BeamOutputPower         respjson.Field
		BeamWaist               respjson.Field
		CreatedAt               respjson.Field
		CreatedBy               respjson.Field
		Divergence              respjson.Field
		IDEntity                respjson.Field
		IsOperational           respjson.Field
		MaxDuration             respjson.Field
		MaxFocusRange           respjson.Field
		MinFocusRange           respjson.Field
		Origin                  respjson.Field
		OrigNetwork             respjson.Field
		PulseFluence            respjson.Field
		PulsePeakPower          respjson.Field
		PulseRepFreq            respjson.Field
		PulseShape              respjson.Field
		PulseWidth              respjson.Field
		SourceDl                respjson.Field
		Wavelength              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaseremitterListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaseremitterListResponse) UnmarshalJSON(data []byte) error {
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
type LaseremitterListResponseDataMode string

const (
	LaseremitterListResponseDataModeReal      LaseremitterListResponseDataMode = "REAL"
	LaseremitterListResponseDataModeTest      LaseremitterListResponseDataMode = "TEST"
	LaseremitterListResponseDataModeSimulated LaseremitterListResponseDataMode = "SIMULATED"
	LaseremitterListResponseDataModeExercise  LaseremitterListResponseDataMode = "EXERCISE"
)

// Model representation of a laser beam emitter/director which represents a source
// of directed energy.
type LaseremitterGetResponse struct {
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
	DataMode LaseremitterGetResponseDataMode `json:"dataMode,required"`
	// The name of this laser within the laser system or facility.
	LaserName string `json:"laserName,required"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType string `json:"laserType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Maximum possible laser-to-space atmospheric transmission (a value between zero
	// and one, usually assumed to be unity for all lasers except for lasers with
	// wavelengths heavily absorbed by the atmosphere).
	AtmosphericTransmission float64 `json:"atmosphericTransmission"`
	// The beam director aperture average (equivalent CW) output power is the total
	// laser average power that is transmitted away from the final exit aperture of the
	// optical system. For a CW laser, this is equivalent to the laser device CW power
	// multiplied by the throughput of the optical system, including the beam director
	// telescope. For a pulsed laser, this is equivalent to the laser device energy per
	// pulse multiplied by the pulse repetition frequency (PRF) multiplied by the
	// throughput of the optical system including the beam director telescope. Measured
	// in Watts.
	BeamOutputPower float64 `json:"beamOutputPower"`
	// The beam waist (radius) of this laser at the exit aperture, in centimeters.
	BeamWaist float64 `json:"beamWaist"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Minimum possible divergence half-angle of this laser, referenced from the
	// optical axis to the 1/e point of the beam's far field irradiance profile,
	// measured in microradians.
	Divergence float64 `json:"divergence"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity shared.EntityFull `json:"entity"`
	// Unique identifier of the parent entity. idEntity is required for PUT.
	IDEntity string `json:"idEntity"`
	// A Boolean indicating whether or not this laser emitter is operational. False
	// indicates that the laser specified in this record is no longer available and can
	// be considered defunct. This field is true by default.
	IsOperational bool `json:"isOperational"`
	// The maximum time that the laser would be "on" for a given test(s) for laser
	// activity, in seconds.
	MaxDuration float64 `json:"maxDuration"`
	// The maximum possible focus range of this laser, measured in kilometers.
	MaxFocusRange float64 `json:"maxFocusRange"`
	// The minimum possible focus range of this laser, measured in kilometers.
	MinFocusRange float64 `json:"minFocusRange"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The amount of energy in each laser pulse exiting from the beam
	// director/telescope, measured in Joules.
	PulseFluence float64 `json:"pulseFluence"`
	// The instantaneous single pulse peak power exiting from the laser device,
	// measured in Watts.
	PulsePeakPower float64 `json:"pulsePeakPower"`
	// The pulse repetition frequency of this laser, measured in kilohertz.
	PulseRepFreq float64 `json:"pulseRepFreq"`
	// The pulse shape (waveform) of this laser, e.g. "RECTANGULAR," "SAWTOOTH,"
	// "GAUSSIAN," etc.
	PulseShape string `json:"pulseShape"`
	// The laser device pulse duration, measured in seconds.
	PulseWidth float64 `json:"pulseWidth"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The center wavelength of this laser, in micrometers.
	Wavelength float64 `json:"wavelength"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking   respjson.Field
		DataMode                respjson.Field
		LaserName               respjson.Field
		LaserType               respjson.Field
		Source                  respjson.Field
		ID                      respjson.Field
		AtmosphericTransmission respjson.Field
		BeamOutputPower         respjson.Field
		BeamWaist               respjson.Field
		CreatedAt               respjson.Field
		CreatedBy               respjson.Field
		Divergence              respjson.Field
		Entity                  respjson.Field
		IDEntity                respjson.Field
		IsOperational           respjson.Field
		MaxDuration             respjson.Field
		MaxFocusRange           respjson.Field
		MinFocusRange           respjson.Field
		Origin                  respjson.Field
		OrigNetwork             respjson.Field
		PulseFluence            respjson.Field
		PulsePeakPower          respjson.Field
		PulseRepFreq            respjson.Field
		PulseShape              respjson.Field
		PulseWidth              respjson.Field
		SourceDl                respjson.Field
		UpdatedAt               respjson.Field
		UpdatedBy               respjson.Field
		Wavelength              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaseremitterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaseremitterGetResponse) UnmarshalJSON(data []byte) error {
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
type LaseremitterGetResponseDataMode string

const (
	LaseremitterGetResponseDataModeReal      LaseremitterGetResponseDataMode = "REAL"
	LaseremitterGetResponseDataModeTest      LaseremitterGetResponseDataMode = "TEST"
	LaseremitterGetResponseDataModeSimulated LaseremitterGetResponseDataMode = "SIMULATED"
	LaseremitterGetResponseDataModeExercise  LaseremitterGetResponseDataMode = "EXERCISE"
)

type LaseremitterQueryhelpResponse struct {
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
func (r LaseremitterQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *LaseremitterQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a laser beam emitter/director which represents a source
// of directed energy.
type LaseremitterTupleResponse struct {
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
	DataMode LaseremitterTupleResponseDataMode `json:"dataMode,required"`
	// The name of this laser within the laser system or facility.
	LaserName string `json:"laserName,required"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType string `json:"laserType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Maximum possible laser-to-space atmospheric transmission (a value between zero
	// and one, usually assumed to be unity for all lasers except for lasers with
	// wavelengths heavily absorbed by the atmosphere).
	AtmosphericTransmission float64 `json:"atmosphericTransmission"`
	// The beam director aperture average (equivalent CW) output power is the total
	// laser average power that is transmitted away from the final exit aperture of the
	// optical system. For a CW laser, this is equivalent to the laser device CW power
	// multiplied by the throughput of the optical system, including the beam director
	// telescope. For a pulsed laser, this is equivalent to the laser device energy per
	// pulse multiplied by the pulse repetition frequency (PRF) multiplied by the
	// throughput of the optical system including the beam director telescope. Measured
	// in Watts.
	BeamOutputPower float64 `json:"beamOutputPower"`
	// The beam waist (radius) of this laser at the exit aperture, in centimeters.
	BeamWaist float64 `json:"beamWaist"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Minimum possible divergence half-angle of this laser, referenced from the
	// optical axis to the 1/e point of the beam's far field irradiance profile,
	// measured in microradians.
	Divergence float64 `json:"divergence"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity shared.EntityFull `json:"entity"`
	// Unique identifier of the parent entity. idEntity is required for PUT.
	IDEntity string `json:"idEntity"`
	// A Boolean indicating whether or not this laser emitter is operational. False
	// indicates that the laser specified in this record is no longer available and can
	// be considered defunct. This field is true by default.
	IsOperational bool `json:"isOperational"`
	// The maximum time that the laser would be "on" for a given test(s) for laser
	// activity, in seconds.
	MaxDuration float64 `json:"maxDuration"`
	// The maximum possible focus range of this laser, measured in kilometers.
	MaxFocusRange float64 `json:"maxFocusRange"`
	// The minimum possible focus range of this laser, measured in kilometers.
	MinFocusRange float64 `json:"minFocusRange"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The amount of energy in each laser pulse exiting from the beam
	// director/telescope, measured in Joules.
	PulseFluence float64 `json:"pulseFluence"`
	// The instantaneous single pulse peak power exiting from the laser device,
	// measured in Watts.
	PulsePeakPower float64 `json:"pulsePeakPower"`
	// The pulse repetition frequency of this laser, measured in kilohertz.
	PulseRepFreq float64 `json:"pulseRepFreq"`
	// The pulse shape (waveform) of this laser, e.g. "RECTANGULAR," "SAWTOOTH,"
	// "GAUSSIAN," etc.
	PulseShape string `json:"pulseShape"`
	// The laser device pulse duration, measured in seconds.
	PulseWidth float64 `json:"pulseWidth"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The center wavelength of this laser, in micrometers.
	Wavelength float64 `json:"wavelength"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking   respjson.Field
		DataMode                respjson.Field
		LaserName               respjson.Field
		LaserType               respjson.Field
		Source                  respjson.Field
		ID                      respjson.Field
		AtmosphericTransmission respjson.Field
		BeamOutputPower         respjson.Field
		BeamWaist               respjson.Field
		CreatedAt               respjson.Field
		CreatedBy               respjson.Field
		Divergence              respjson.Field
		Entity                  respjson.Field
		IDEntity                respjson.Field
		IsOperational           respjson.Field
		MaxDuration             respjson.Field
		MaxFocusRange           respjson.Field
		MinFocusRange           respjson.Field
		Origin                  respjson.Field
		OrigNetwork             respjson.Field
		PulseFluence            respjson.Field
		PulsePeakPower          respjson.Field
		PulseRepFreq            respjson.Field
		PulseShape              respjson.Field
		PulseWidth              respjson.Field
		SourceDl                respjson.Field
		UpdatedAt               respjson.Field
		UpdatedBy               respjson.Field
		Wavelength              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaseremitterTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LaseremitterTupleResponse) UnmarshalJSON(data []byte) error {
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
type LaseremitterTupleResponseDataMode string

const (
	LaseremitterTupleResponseDataModeReal      LaseremitterTupleResponseDataMode = "REAL"
	LaseremitterTupleResponseDataModeTest      LaseremitterTupleResponseDataMode = "TEST"
	LaseremitterTupleResponseDataModeSimulated LaseremitterTupleResponseDataMode = "SIMULATED"
	LaseremitterTupleResponseDataModeExercise  LaseremitterTupleResponseDataMode = "EXERCISE"
)

type LaseremitterNewParams struct {
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
	DataMode LaseremitterNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The name of this laser within the laser system or facility.
	LaserName string `json:"laserName,required"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType string `json:"laserType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Maximum possible laser-to-space atmospheric transmission (a value between zero
	// and one, usually assumed to be unity for all lasers except for lasers with
	// wavelengths heavily absorbed by the atmosphere).
	AtmosphericTransmission param.Opt[float64] `json:"atmosphericTransmission,omitzero"`
	// The beam director aperture average (equivalent CW) output power is the total
	// laser average power that is transmitted away from the final exit aperture of the
	// optical system. For a CW laser, this is equivalent to the laser device CW power
	// multiplied by the throughput of the optical system, including the beam director
	// telescope. For a pulsed laser, this is equivalent to the laser device energy per
	// pulse multiplied by the pulse repetition frequency (PRF) multiplied by the
	// throughput of the optical system including the beam director telescope. Measured
	// in Watts.
	BeamOutputPower param.Opt[float64] `json:"beamOutputPower,omitzero"`
	// The beam waist (radius) of this laser at the exit aperture, in centimeters.
	BeamWaist param.Opt[float64] `json:"beamWaist,omitzero"`
	// Minimum possible divergence half-angle of this laser, referenced from the
	// optical axis to the 1/e point of the beam's far field irradiance profile,
	// measured in microradians.
	Divergence param.Opt[float64] `json:"divergence,omitzero"`
	// Unique identifier of the parent entity. idEntity is required for PUT.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// A Boolean indicating whether or not this laser emitter is operational. False
	// indicates that the laser specified in this record is no longer available and can
	// be considered defunct. This field is true by default.
	IsOperational param.Opt[bool] `json:"isOperational,omitzero"`
	// The maximum time that the laser would be "on" for a given test(s) for laser
	// activity, in seconds.
	MaxDuration param.Opt[float64] `json:"maxDuration,omitzero"`
	// The maximum possible focus range of this laser, measured in kilometers.
	MaxFocusRange param.Opt[float64] `json:"maxFocusRange,omitzero"`
	// The minimum possible focus range of this laser, measured in kilometers.
	MinFocusRange param.Opt[float64] `json:"minFocusRange,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The amount of energy in each laser pulse exiting from the beam
	// director/telescope, measured in Joules.
	PulseFluence param.Opt[float64] `json:"pulseFluence,omitzero"`
	// The instantaneous single pulse peak power exiting from the laser device,
	// measured in Watts.
	PulsePeakPower param.Opt[float64] `json:"pulsePeakPower,omitzero"`
	// The pulse repetition frequency of this laser, measured in kilohertz.
	PulseRepFreq param.Opt[float64] `json:"pulseRepFreq,omitzero"`
	// The pulse shape (waveform) of this laser, e.g. "RECTANGULAR," "SAWTOOTH,"
	// "GAUSSIAN," etc.
	PulseShape param.Opt[string] `json:"pulseShape,omitzero"`
	// The laser device pulse duration, measured in seconds.
	PulseWidth param.Opt[float64] `json:"pulseWidth,omitzero"`
	// The center wavelength of this laser, in micrometers.
	Wavelength param.Opt[float64] `json:"wavelength,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r LaseremitterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LaseremitterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaseremitterNewParams) UnmarshalJSON(data []byte) error {
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
type LaseremitterNewParamsDataMode string

const (
	LaseremitterNewParamsDataModeReal      LaseremitterNewParamsDataMode = "REAL"
	LaseremitterNewParamsDataModeTest      LaseremitterNewParamsDataMode = "TEST"
	LaseremitterNewParamsDataModeSimulated LaseremitterNewParamsDataMode = "SIMULATED"
	LaseremitterNewParamsDataModeExercise  LaseremitterNewParamsDataMode = "EXERCISE"
)

type LaseremitterUpdateParams struct {
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
	DataMode LaseremitterUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The name of this laser within the laser system or facility.
	LaserName string `json:"laserName,required"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType string `json:"laserType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Maximum possible laser-to-space atmospheric transmission (a value between zero
	// and one, usually assumed to be unity for all lasers except for lasers with
	// wavelengths heavily absorbed by the atmosphere).
	AtmosphericTransmission param.Opt[float64] `json:"atmosphericTransmission,omitzero"`
	// The beam director aperture average (equivalent CW) output power is the total
	// laser average power that is transmitted away from the final exit aperture of the
	// optical system. For a CW laser, this is equivalent to the laser device CW power
	// multiplied by the throughput of the optical system, including the beam director
	// telescope. For a pulsed laser, this is equivalent to the laser device energy per
	// pulse multiplied by the pulse repetition frequency (PRF) multiplied by the
	// throughput of the optical system including the beam director telescope. Measured
	// in Watts.
	BeamOutputPower param.Opt[float64] `json:"beamOutputPower,omitzero"`
	// The beam waist (radius) of this laser at the exit aperture, in centimeters.
	BeamWaist param.Opt[float64] `json:"beamWaist,omitzero"`
	// Minimum possible divergence half-angle of this laser, referenced from the
	// optical axis to the 1/e point of the beam's far field irradiance profile,
	// measured in microradians.
	Divergence param.Opt[float64] `json:"divergence,omitzero"`
	// Unique identifier of the parent entity. idEntity is required for PUT.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// A Boolean indicating whether or not this laser emitter is operational. False
	// indicates that the laser specified in this record is no longer available and can
	// be considered defunct. This field is true by default.
	IsOperational param.Opt[bool] `json:"isOperational,omitzero"`
	// The maximum time that the laser would be "on" for a given test(s) for laser
	// activity, in seconds.
	MaxDuration param.Opt[float64] `json:"maxDuration,omitzero"`
	// The maximum possible focus range of this laser, measured in kilometers.
	MaxFocusRange param.Opt[float64] `json:"maxFocusRange,omitzero"`
	// The minimum possible focus range of this laser, measured in kilometers.
	MinFocusRange param.Opt[float64] `json:"minFocusRange,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The amount of energy in each laser pulse exiting from the beam
	// director/telescope, measured in Joules.
	PulseFluence param.Opt[float64] `json:"pulseFluence,omitzero"`
	// The instantaneous single pulse peak power exiting from the laser device,
	// measured in Watts.
	PulsePeakPower param.Opt[float64] `json:"pulsePeakPower,omitzero"`
	// The pulse repetition frequency of this laser, measured in kilohertz.
	PulseRepFreq param.Opt[float64] `json:"pulseRepFreq,omitzero"`
	// The pulse shape (waveform) of this laser, e.g. "RECTANGULAR," "SAWTOOTH,"
	// "GAUSSIAN," etc.
	PulseShape param.Opt[string] `json:"pulseShape,omitzero"`
	// The laser device pulse duration, measured in seconds.
	PulseWidth param.Opt[float64] `json:"pulseWidth,omitzero"`
	// The center wavelength of this laser, in micrometers.
	Wavelength param.Opt[float64] `json:"wavelength,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r LaseremitterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LaseremitterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaseremitterUpdateParams) UnmarshalJSON(data []byte) error {
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
type LaseremitterUpdateParamsDataMode string

const (
	LaseremitterUpdateParamsDataModeReal      LaseremitterUpdateParamsDataMode = "REAL"
	LaseremitterUpdateParamsDataModeTest      LaseremitterUpdateParamsDataMode = "TEST"
	LaseremitterUpdateParamsDataModeSimulated LaseremitterUpdateParamsDataMode = "SIMULATED"
	LaseremitterUpdateParamsDataModeExercise  LaseremitterUpdateParamsDataMode = "EXERCISE"
)

type LaseremitterListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaseremitterListParams]'s query parameters as `url.Values`.
func (r LaseremitterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaseremitterCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaseremitterCountParams]'s query parameters as
// `url.Values`.
func (r LaseremitterCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaseremitterGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaseremitterGetParams]'s query parameters as `url.Values`.
func (r LaseremitterGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaseremitterTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaseremitterTupleParams]'s query parameters as
// `url.Values`.
func (r LaseremitterTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
