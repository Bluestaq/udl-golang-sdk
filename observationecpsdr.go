// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// ObservationEcpsdrService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationEcpsdrService] method instead.
type ObservationEcpsdrService struct {
	Options []option.RequestOption
}

// NewObservationEcpsdrService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObservationEcpsdrService(opts ...option.RequestOption) (r ObservationEcpsdrService) {
	r = ObservationEcpsdrService{}
	r.Options = opts
	return
}

// Service operation to take a single ECPSDR as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *ObservationEcpsdrService) New(ctx context.Context, body ObservationEcpsdrNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/ecpsdr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single ECPSDR by its unique ID passed as a path
// parameter.
func (r *ObservationEcpsdrService) Get(ctx context.Context, id string, query ObservationEcpsdrGetParams, opts ...option.RequestOption) (res *Ecpsdr, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/ecpsdr/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationEcpsdrService) List(ctx context.Context, query ObservationEcpsdrListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EcpsdrAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/ecpsdr"
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
func (r *ObservationEcpsdrService) ListAutoPaging(ctx context.Context, query ObservationEcpsdrListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EcpsdrAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationEcpsdrService) Count(ctx context.Context, query ObservationEcpsdrCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/ecpsdr/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// ECPSDR as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *ObservationEcpsdrService) NewBulk(ctx context.Context, body ObservationEcpsdrNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/ecpsdr/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ObservationEcpsdrService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *ObservationEcpsdrQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/ecpsdr/queryhelp"
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
func (r *ObservationEcpsdrService) Tuple(ctx context.Context, query ObservationEcpsdrTupleParams, opts ...option.RequestOption) (res *[]Ecpsdr, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/ecpsdr/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple ECPSDR as a POST body and ingest into the
// database. This operation is intended to be used for automated feeds into UDL. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *ObservationEcpsdrService) UnvalidatedPublish(ctx context.Context, body ObservationEcpsdrUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-ecpsdr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Supports Sensor Data Records (SDR) from space-borne Energetic Charged Particle
// (ECP) Sensors. SDR contains sensor status telemetry and raw dosimeter
// measurements of the space environment.
type Ecpsdr struct {
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
	DataMode EcpsdrDataMode `json:"dataMode,required"`
	// Time stamp of time packet receipt on ground, in ISO 8601 UTC format with
	// millisecond precision.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of data associated with this record (STANDARD, TRANSIENT).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Internal 5V current monitor for analog supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Asl5VCurrMon int64 `json:"asl5VCurrMon"`
	// CDS Charge Plate voltage monitor. See vRef for conversion factor to Volts.
	CdsPlateVMon int64 `json:"cdsPlateVMon"`
	// CDS reference voltage monitor. See vRef for conversion factor to Volts.
	CdsRefVMon int64 `json:"cdsRefVMon"`
	// CDS Threshold setting for ESD detection threshold. The CDS Threshold is the
	// adjustable sensitivity of recording/digitizing an ESD as a transient packet.
	CdsThreshold int64 `json:"cdsThreshold"`
	// CDS throttle number of seconds between CDS transient capture readouts.
	CdsThrottle int64 `json:"cdsThrottle"`
	// Two byte CRC-16-CCITT checksum (ordered as first byte, second byte).
	Checksum int64 `json:"checksum"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unitless dosimeter detector bias for MedLET and HiLET. MedLET (Linear Energy
	// Transfer) and HiLET subsensors detect particles above LET thresholds, 300keV and
	// 1MeV, respectively.
	DosBias int64 `json:"dosBias"`
	// Internal 5V current monitor for digital supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Dsl5VCurrMon int64 `json:"dsl5VCurrMon"`
	// Number of ESD triggers, high byte of 2-byte counter.
	EsdTrigCountH int64 `json:"esdTrigCountH"`
	// Number of ESD triggers, low byte of 2-byte counter.
	EsdTrigCountL int64 `json:"esdTrigCountL"`
	// HiLET dosimeter low range output. Low byte of scaler (HiLET) dosimeter output.
	HiLetL int64 `json:"hiLetL"`
	// Unitless HiLET dosimeter medium range output. Medium byte of (HiLET) dosimeter
	// output.
	HiLetM int64 `json:"hiLetM"`
	// Unique identifier of the on-orbit satellite hosting the sensor.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// LowLET dosimeter low range output. Low byte of (LowLET) dosimeter output.
	LowLetL int64 `json:"lowLetL"`
	// LowLET dosimeter medium range output. Medium byte of (LowLET) dosimeter output.
	LowLetM int64 `json:"lowLetM"`
	// MedLET1 dosimeter low range output. Low byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1L int64 `json:"medLet1L"`
	// MedLET1 dosimeter medium range output. Medium byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1M int64 `json:"medLet1M"`
	// MedLET2 dosimeter low range output. Low byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2L int64 `json:"medLet2L"`
	// MedLET2 dosimeter medium range output. Medium byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2M int64 `json:"medLet2M"`
	// MedLET3 dosimeter low range output. Low byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3L int64 `json:"medLet3L"`
	// MedLET3 dosimeter medium range output. Medium byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3M int64 `json:"medLet3M"`
	// MedLET4 dosimeter low range output. Low byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4L int64 `json:"medLet4L"`
	// MedLET4 dosimeter medium range output. Medium byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4M int64 `json:"medLet4M"`
	// Unitless sensor mounting plate temperature.
	MpTemp int64 `json:"mpTemp"`
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	ObTime time.Time `json:"obTime" format:"date-time"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the record source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Photodiode 1 signal level.
	Pd1SigLev int64 `json:"pd1SigLev"`
	// Photodiode 2 signal level.
	Pd2SigLev int64 `json:"pd2SigLev"`
	// Power supply temperature monitor. This is sensor status telemetry.
	PsTempMon int64 `json:"psTempMon"`
	// Flag indicating whether this record is an original or re-transmitted dataset
	// (TRUE indicates a retransmit from the host).
	Retransmit bool `json:"retransmit"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo int64 `json:"satNo"`
	// The sensor mode associated with this measurements (NORMAL, TEST).
	SenMode string `json:"senMode"`
	// Surface dosimeter charge rate high output (converts to pico-amps/bit). High byte
	// of 2 bytes.
	SurfDosChargeH int64 `json:"surfDosChargeH"`
	// Surface dosimeter charge rate low output (converts to pico-amps/bit). Low byte
	// of 2 bytes.
	SurfDosChargeL int64 `json:"surfDosChargeL"`
	// Surface dosimeter high range output (converts to pico-coulombs/bit). High byte
	// of 3 bytes.
	SurfDosH int64 `json:"surfDosH"`
	// Surface dosimeter low range output (converts to pico-coulombs/bit). Low byte of
	// 3 bytes.
	SurfDosL int64 `json:"surfDosL"`
	// Surface dosimeter medium range output (converts to pico-coulombs/bit). Middle
	// byte of 3 bytes.
	SurfDosM int64 `json:"surfDosM"`
	// Surface dosimeter status byte.
	SurfDosStat int64 `json:"surfDosStat"`
	// Array of 144 digitized samples of ESD waveform for transient packets.
	TransientData []int64 `json:"transientData"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Reference voltage (volts/bit). Conversion factor used to convert analog V
	// monitor data from bytes to volts.
	VRef int64 `json:"vRef"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		MsgTime               respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		Asl5VCurrMon          respjson.Field
		CdsPlateVMon          respjson.Field
		CdsRefVMon            respjson.Field
		CdsThreshold          respjson.Field
		CdsThrottle           respjson.Field
		Checksum              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DosBias               respjson.Field
		Dsl5VCurrMon          respjson.Field
		EsdTrigCountH         respjson.Field
		EsdTrigCountL         respjson.Field
		HiLetL                respjson.Field
		HiLetM                respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		LowLetL               respjson.Field
		LowLetM               respjson.Field
		MedLet1L              respjson.Field
		MedLet1M              respjson.Field
		MedLet2L              respjson.Field
		MedLet2M              respjson.Field
		MedLet3L              respjson.Field
		MedLet3M              respjson.Field
		MedLet4L              respjson.Field
		MedLet4M              respjson.Field
		MpTemp                respjson.Field
		ObTime                respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		Pd1SigLev             respjson.Field
		Pd2SigLev             respjson.Field
		PsTempMon             respjson.Field
		Retransmit            respjson.Field
		SatNo                 respjson.Field
		SenMode               respjson.Field
		SurfDosChargeH        respjson.Field
		SurfDosChargeL        respjson.Field
		SurfDosH              respjson.Field
		SurfDosL              respjson.Field
		SurfDosM              respjson.Field
		SurfDosStat           respjson.Field
		TransientData         respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VRef                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Ecpsdr) RawJSON() string { return r.JSON.raw }
func (r *Ecpsdr) UnmarshalJSON(data []byte) error {
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
type EcpsdrDataMode string

const (
	EcpsdrDataModeReal      EcpsdrDataMode = "REAL"
	EcpsdrDataModeTest      EcpsdrDataMode = "TEST"
	EcpsdrDataModeSimulated EcpsdrDataMode = "SIMULATED"
	EcpsdrDataModeExercise  EcpsdrDataMode = "EXERCISE"
)

// Supports Sensor Data Records (SDR) from space-borne Energetic Charged Particle
// (ECP) Sensors. SDR contains sensor status telemetry and raw dosimeter
// measurements of the space environment.
type EcpsdrAbridged struct {
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
	DataMode EcpsdrAbridgedDataMode `json:"dataMode,required"`
	// Time stamp of time packet receipt on ground, in ISO 8601 UTC format with
	// millisecond precision.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of data associated with this record (STANDARD, TRANSIENT).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Internal 5V current monitor for analog supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Asl5VCurrMon int64 `json:"asl5VCurrMon"`
	// CDS Charge Plate voltage monitor. See vRef for conversion factor to Volts.
	CdsPlateVMon int64 `json:"cdsPlateVMon"`
	// CDS reference voltage monitor. See vRef for conversion factor to Volts.
	CdsRefVMon int64 `json:"cdsRefVMon"`
	// CDS Threshold setting for ESD detection threshold. The CDS Threshold is the
	// adjustable sensitivity of recording/digitizing an ESD as a transient packet.
	CdsThreshold int64 `json:"cdsThreshold"`
	// CDS throttle number of seconds between CDS transient capture readouts.
	CdsThrottle int64 `json:"cdsThrottle"`
	// Two byte CRC-16-CCITT checksum (ordered as first byte, second byte).
	Checksum int64 `json:"checksum"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unitless dosimeter detector bias for MedLET and HiLET. MedLET (Linear Energy
	// Transfer) and HiLET subsensors detect particles above LET thresholds, 300keV and
	// 1MeV, respectively.
	DosBias int64 `json:"dosBias"`
	// Internal 5V current monitor for digital supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Dsl5VCurrMon int64 `json:"dsl5VCurrMon"`
	// Number of ESD triggers, high byte of 2-byte counter.
	EsdTrigCountH int64 `json:"esdTrigCountH"`
	// Number of ESD triggers, low byte of 2-byte counter.
	EsdTrigCountL int64 `json:"esdTrigCountL"`
	// HiLET dosimeter low range output. Low byte of scaler (HiLET) dosimeter output.
	HiLetL int64 `json:"hiLetL"`
	// Unitless HiLET dosimeter medium range output. Medium byte of (HiLET) dosimeter
	// output.
	HiLetM int64 `json:"hiLetM"`
	// Unique identifier of the on-orbit satellite hosting the sensor.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// LowLET dosimeter low range output. Low byte of (LowLET) dosimeter output.
	LowLetL int64 `json:"lowLetL"`
	// LowLET dosimeter medium range output. Medium byte of (LowLET) dosimeter output.
	LowLetM int64 `json:"lowLetM"`
	// MedLET1 dosimeter low range output. Low byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1L int64 `json:"medLet1L"`
	// MedLET1 dosimeter medium range output. Medium byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1M int64 `json:"medLet1M"`
	// MedLET2 dosimeter low range output. Low byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2L int64 `json:"medLet2L"`
	// MedLET2 dosimeter medium range output. Medium byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2M int64 `json:"medLet2M"`
	// MedLET3 dosimeter low range output. Low byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3L int64 `json:"medLet3L"`
	// MedLET3 dosimeter medium range output. Medium byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3M int64 `json:"medLet3M"`
	// MedLET4 dosimeter low range output. Low byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4L int64 `json:"medLet4L"`
	// MedLET4 dosimeter medium range output. Medium byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4M int64 `json:"medLet4M"`
	// Unitless sensor mounting plate temperature.
	MpTemp int64 `json:"mpTemp"`
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	ObTime time.Time `json:"obTime" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the record source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Photodiode 1 signal level.
	Pd1SigLev int64 `json:"pd1SigLev"`
	// Photodiode 2 signal level.
	Pd2SigLev int64 `json:"pd2SigLev"`
	// Power supply temperature monitor. This is sensor status telemetry.
	PsTempMon int64 `json:"psTempMon"`
	// Flag indicating whether this record is an original or re-transmitted dataset
	// (TRUE indicates a retransmit from the host).
	Retransmit bool `json:"retransmit"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo int64 `json:"satNo"`
	// The sensor mode associated with this measurements (NORMAL, TEST).
	SenMode string `json:"senMode"`
	// Surface dosimeter charge rate high output (converts to pico-amps/bit). High byte
	// of 2 bytes.
	SurfDosChargeH int64 `json:"surfDosChargeH"`
	// Surface dosimeter charge rate low output (converts to pico-amps/bit). Low byte
	// of 2 bytes.
	SurfDosChargeL int64 `json:"surfDosChargeL"`
	// Surface dosimeter high range output (converts to pico-coulombs/bit). High byte
	// of 3 bytes.
	SurfDosH int64 `json:"surfDosH"`
	// Surface dosimeter low range output (converts to pico-coulombs/bit). Low byte of
	// 3 bytes.
	SurfDosL int64 `json:"surfDosL"`
	// Surface dosimeter medium range output (converts to pico-coulombs/bit). Middle
	// byte of 3 bytes.
	SurfDosM int64 `json:"surfDosM"`
	// Surface dosimeter status byte.
	SurfDosStat int64 `json:"surfDosStat"`
	// Array of 144 digitized samples of ESD waveform for transient packets.
	TransientData []int64 `json:"transientData"`
	// Reference voltage (volts/bit). Conversion factor used to convert analog V
	// monitor data from bytes to volts.
	VRef int64 `json:"vRef"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		MsgTime               respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		Asl5VCurrMon          respjson.Field
		CdsPlateVMon          respjson.Field
		CdsRefVMon            respjson.Field
		CdsThreshold          respjson.Field
		CdsThrottle           respjson.Field
		Checksum              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DosBias               respjson.Field
		Dsl5VCurrMon          respjson.Field
		EsdTrigCountH         respjson.Field
		EsdTrigCountL         respjson.Field
		HiLetL                respjson.Field
		HiLetM                respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		LowLetL               respjson.Field
		LowLetM               respjson.Field
		MedLet1L              respjson.Field
		MedLet1M              respjson.Field
		MedLet2L              respjson.Field
		MedLet2M              respjson.Field
		MedLet3L              respjson.Field
		MedLet3M              respjson.Field
		MedLet4L              respjson.Field
		MedLet4M              respjson.Field
		MpTemp                respjson.Field
		ObTime                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		Pd1SigLev             respjson.Field
		Pd2SigLev             respjson.Field
		PsTempMon             respjson.Field
		Retransmit            respjson.Field
		SatNo                 respjson.Field
		SenMode               respjson.Field
		SurfDosChargeH        respjson.Field
		SurfDosChargeL        respjson.Field
		SurfDosH              respjson.Field
		SurfDosL              respjson.Field
		SurfDosM              respjson.Field
		SurfDosStat           respjson.Field
		TransientData         respjson.Field
		VRef                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EcpsdrAbridged) RawJSON() string { return r.JSON.raw }
func (r *EcpsdrAbridged) UnmarshalJSON(data []byte) error {
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
type EcpsdrAbridgedDataMode string

const (
	EcpsdrAbridgedDataModeReal      EcpsdrAbridgedDataMode = "REAL"
	EcpsdrAbridgedDataModeTest      EcpsdrAbridgedDataMode = "TEST"
	EcpsdrAbridgedDataModeSimulated EcpsdrAbridgedDataMode = "SIMULATED"
	EcpsdrAbridgedDataModeExercise  EcpsdrAbridgedDataMode = "EXERCISE"
)

type ObservationEcpsdrQueryHelpResponse struct {
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
func (r ObservationEcpsdrQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationEcpsdrQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObservationEcpsdrNewParams struct {
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
	DataMode ObservationEcpsdrNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Time stamp of time packet receipt on ground, in ISO 8601 UTC format with
	// millisecond precision.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of data associated with this record (STANDARD, TRANSIENT).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Internal 5V current monitor for analog supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Asl5VCurrMon param.Opt[int64] `json:"asl5VCurrMon,omitzero"`
	// CDS Charge Plate voltage monitor. See vRef for conversion factor to Volts.
	CdsPlateVMon param.Opt[int64] `json:"cdsPlateVMon,omitzero"`
	// CDS reference voltage monitor. See vRef for conversion factor to Volts.
	CdsRefVMon param.Opt[int64] `json:"cdsRefVMon,omitzero"`
	// CDS Threshold setting for ESD detection threshold. The CDS Threshold is the
	// adjustable sensitivity of recording/digitizing an ESD as a transient packet.
	CdsThreshold param.Opt[int64] `json:"cdsThreshold,omitzero"`
	// CDS throttle number of seconds between CDS transient capture readouts.
	CdsThrottle param.Opt[int64] `json:"cdsThrottle,omitzero"`
	// Two byte CRC-16-CCITT checksum (ordered as first byte, second byte).
	Checksum param.Opt[int64] `json:"checksum,omitzero"`
	// Unitless dosimeter detector bias for MedLET and HiLET. MedLET (Linear Energy
	// Transfer) and HiLET subsensors detect particles above LET thresholds, 300keV and
	// 1MeV, respectively.
	DosBias param.Opt[int64] `json:"dosBias,omitzero"`
	// Internal 5V current monitor for digital supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Dsl5VCurrMon param.Opt[int64] `json:"dsl5VCurrMon,omitzero"`
	// Number of ESD triggers, high byte of 2-byte counter.
	EsdTrigCountH param.Opt[int64] `json:"esdTrigCountH,omitzero"`
	// Number of ESD triggers, low byte of 2-byte counter.
	EsdTrigCountL param.Opt[int64] `json:"esdTrigCountL,omitzero"`
	// HiLET dosimeter low range output. Low byte of scaler (HiLET) dosimeter output.
	HiLetL param.Opt[int64] `json:"hiLetL,omitzero"`
	// Unitless HiLET dosimeter medium range output. Medium byte of (HiLET) dosimeter
	// output.
	HiLetM param.Opt[int64] `json:"hiLetM,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// LowLET dosimeter low range output. Low byte of (LowLET) dosimeter output.
	LowLetL param.Opt[int64] `json:"lowLetL,omitzero"`
	// LowLET dosimeter medium range output. Medium byte of (LowLET) dosimeter output.
	LowLetM param.Opt[int64] `json:"lowLetM,omitzero"`
	// MedLET1 dosimeter low range output. Low byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1L param.Opt[int64] `json:"medLet1L,omitzero"`
	// MedLET1 dosimeter medium range output. Medium byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1M param.Opt[int64] `json:"medLet1M,omitzero"`
	// MedLET2 dosimeter low range output. Low byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2L param.Opt[int64] `json:"medLet2L,omitzero"`
	// MedLET2 dosimeter medium range output. Medium byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2M param.Opt[int64] `json:"medLet2M,omitzero"`
	// MedLET3 dosimeter low range output. Low byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3L param.Opt[int64] `json:"medLet3L,omitzero"`
	// MedLET3 dosimeter medium range output. Medium byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3M param.Opt[int64] `json:"medLet3M,omitzero"`
	// MedLET4 dosimeter low range output. Low byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4L param.Opt[int64] `json:"medLet4L,omitzero"`
	// MedLET4 dosimeter medium range output. Medium byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4M param.Opt[int64] `json:"medLet4M,omitzero"`
	// Unitless sensor mounting plate temperature.
	MpTemp param.Opt[int64] `json:"mpTemp,omitzero"`
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	ObTime param.Opt[time.Time] `json:"obTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the record source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Photodiode 1 signal level.
	Pd1SigLev param.Opt[int64] `json:"pd1SigLev,omitzero"`
	// Photodiode 2 signal level.
	Pd2SigLev param.Opt[int64] `json:"pd2SigLev,omitzero"`
	// Power supply temperature monitor. This is sensor status telemetry.
	PsTempMon param.Opt[int64] `json:"psTempMon,omitzero"`
	// Flag indicating whether this record is an original or re-transmitted dataset
	// (TRUE indicates a retransmit from the host).
	Retransmit param.Opt[bool] `json:"retransmit,omitzero"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The sensor mode associated with this measurements (NORMAL, TEST).
	SenMode param.Opt[string] `json:"senMode,omitzero"`
	// Surface dosimeter charge rate high output (converts to pico-amps/bit). High byte
	// of 2 bytes.
	SurfDosChargeH param.Opt[int64] `json:"surfDosChargeH,omitzero"`
	// Surface dosimeter charge rate low output (converts to pico-amps/bit). Low byte
	// of 2 bytes.
	SurfDosChargeL param.Opt[int64] `json:"surfDosChargeL,omitzero"`
	// Surface dosimeter high range output (converts to pico-coulombs/bit). High byte
	// of 3 bytes.
	SurfDosH param.Opt[int64] `json:"surfDosH,omitzero"`
	// Surface dosimeter low range output (converts to pico-coulombs/bit). Low byte of
	// 3 bytes.
	SurfDosL param.Opt[int64] `json:"surfDosL,omitzero"`
	// Surface dosimeter medium range output (converts to pico-coulombs/bit). Middle
	// byte of 3 bytes.
	SurfDosM param.Opt[int64] `json:"surfDosM,omitzero"`
	// Surface dosimeter status byte.
	SurfDosStat param.Opt[int64] `json:"surfDosStat,omitzero"`
	// Reference voltage (volts/bit). Conversion factor used to convert analog V
	// monitor data from bytes to volts.
	VRef param.Opt[int64] `json:"vRef,omitzero"`
	// Array of 144 digitized samples of ESD waveform for transient packets.
	TransientData []int64 `json:"transientData,omitzero"`
	paramObj
}

func (r ObservationEcpsdrNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObservationEcpsdrNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationEcpsdrNewParams) UnmarshalJSON(data []byte) error {
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
type ObservationEcpsdrNewParamsDataMode string

const (
	ObservationEcpsdrNewParamsDataModeReal      ObservationEcpsdrNewParamsDataMode = "REAL"
	ObservationEcpsdrNewParamsDataModeTest      ObservationEcpsdrNewParamsDataMode = "TEST"
	ObservationEcpsdrNewParamsDataModeSimulated ObservationEcpsdrNewParamsDataMode = "SIMULATED"
	ObservationEcpsdrNewParamsDataModeExercise  ObservationEcpsdrNewParamsDataMode = "EXERCISE"
)

type ObservationEcpsdrGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationEcpsdrGetParams]'s query parameters as
// `url.Values`.
func (r ObservationEcpsdrGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEcpsdrListParams struct {
	// Time stamp of time packet receipt on ground, in ISO 8601 UTC format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationEcpsdrListParams]'s query parameters as
// `url.Values`.
func (r ObservationEcpsdrListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEcpsdrCountParams struct {
	// Time stamp of time packet receipt on ground, in ISO 8601 UTC format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationEcpsdrCountParams]'s query parameters as
// `url.Values`.
func (r ObservationEcpsdrCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEcpsdrNewBulkParams struct {
	Body []ObservationEcpsdrNewBulkParamsBody
	paramObj
}

func (r ObservationEcpsdrNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationEcpsdrNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Supports Sensor Data Records (SDR) from space-borne Energetic Charged Particle
// (ECP) Sensors. SDR contains sensor status telemetry and raw dosimeter
// measurements of the space environment.
//
// The properties ClassificationMarking, DataMode, MsgTime, Source, Type are
// required.
type ObservationEcpsdrNewBulkParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Time stamp of time packet receipt on ground, in ISO 8601 UTC format with
	// millisecond precision.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of data associated with this record (STANDARD, TRANSIENT).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Internal 5V current monitor for analog supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Asl5VCurrMon param.Opt[int64] `json:"asl5VCurrMon,omitzero"`
	// CDS Charge Plate voltage monitor. See vRef for conversion factor to Volts.
	CdsPlateVMon param.Opt[int64] `json:"cdsPlateVMon,omitzero"`
	// CDS reference voltage monitor. See vRef for conversion factor to Volts.
	CdsRefVMon param.Opt[int64] `json:"cdsRefVMon,omitzero"`
	// CDS Threshold setting for ESD detection threshold. The CDS Threshold is the
	// adjustable sensitivity of recording/digitizing an ESD as a transient packet.
	CdsThreshold param.Opt[int64] `json:"cdsThreshold,omitzero"`
	// CDS throttle number of seconds between CDS transient capture readouts.
	CdsThrottle param.Opt[int64] `json:"cdsThrottle,omitzero"`
	// Two byte CRC-16-CCITT checksum (ordered as first byte, second byte).
	Checksum param.Opt[int64] `json:"checksum,omitzero"`
	// Unitless dosimeter detector bias for MedLET and HiLET. MedLET (Linear Energy
	// Transfer) and HiLET subsensors detect particles above LET thresholds, 300keV and
	// 1MeV, respectively.
	DosBias param.Opt[int64] `json:"dosBias,omitzero"`
	// Internal 5V current monitor for digital supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Dsl5VCurrMon param.Opt[int64] `json:"dsl5VCurrMon,omitzero"`
	// Number of ESD triggers, high byte of 2-byte counter.
	EsdTrigCountH param.Opt[int64] `json:"esdTrigCountH,omitzero"`
	// Number of ESD triggers, low byte of 2-byte counter.
	EsdTrigCountL param.Opt[int64] `json:"esdTrigCountL,omitzero"`
	// HiLET dosimeter low range output. Low byte of scaler (HiLET) dosimeter output.
	HiLetL param.Opt[int64] `json:"hiLetL,omitzero"`
	// Unitless HiLET dosimeter medium range output. Medium byte of (HiLET) dosimeter
	// output.
	HiLetM param.Opt[int64] `json:"hiLetM,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// LowLET dosimeter low range output. Low byte of (LowLET) dosimeter output.
	LowLetL param.Opt[int64] `json:"lowLetL,omitzero"`
	// LowLET dosimeter medium range output. Medium byte of (LowLET) dosimeter output.
	LowLetM param.Opt[int64] `json:"lowLetM,omitzero"`
	// MedLET1 dosimeter low range output. Low byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1L param.Opt[int64] `json:"medLet1L,omitzero"`
	// MedLET1 dosimeter medium range output. Medium byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1M param.Opt[int64] `json:"medLet1M,omitzero"`
	// MedLET2 dosimeter low range output. Low byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2L param.Opt[int64] `json:"medLet2L,omitzero"`
	// MedLET2 dosimeter medium range output. Medium byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2M param.Opt[int64] `json:"medLet2M,omitzero"`
	// MedLET3 dosimeter low range output. Low byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3L param.Opt[int64] `json:"medLet3L,omitzero"`
	// MedLET3 dosimeter medium range output. Medium byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3M param.Opt[int64] `json:"medLet3M,omitzero"`
	// MedLET4 dosimeter low range output. Low byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4L param.Opt[int64] `json:"medLet4L,omitzero"`
	// MedLET4 dosimeter medium range output. Medium byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4M param.Opt[int64] `json:"medLet4M,omitzero"`
	// Unitless sensor mounting plate temperature.
	MpTemp param.Opt[int64] `json:"mpTemp,omitzero"`
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	ObTime param.Opt[time.Time] `json:"obTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the record source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Photodiode 1 signal level.
	Pd1SigLev param.Opt[int64] `json:"pd1SigLev,omitzero"`
	// Photodiode 2 signal level.
	Pd2SigLev param.Opt[int64] `json:"pd2SigLev,omitzero"`
	// Power supply temperature monitor. This is sensor status telemetry.
	PsTempMon param.Opt[int64] `json:"psTempMon,omitzero"`
	// Flag indicating whether this record is an original or re-transmitted dataset
	// (TRUE indicates a retransmit from the host).
	Retransmit param.Opt[bool] `json:"retransmit,omitzero"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The sensor mode associated with this measurements (NORMAL, TEST).
	SenMode param.Opt[string] `json:"senMode,omitzero"`
	// Surface dosimeter charge rate high output (converts to pico-amps/bit). High byte
	// of 2 bytes.
	SurfDosChargeH param.Opt[int64] `json:"surfDosChargeH,omitzero"`
	// Surface dosimeter charge rate low output (converts to pico-amps/bit). Low byte
	// of 2 bytes.
	SurfDosChargeL param.Opt[int64] `json:"surfDosChargeL,omitzero"`
	// Surface dosimeter high range output (converts to pico-coulombs/bit). High byte
	// of 3 bytes.
	SurfDosH param.Opt[int64] `json:"surfDosH,omitzero"`
	// Surface dosimeter low range output (converts to pico-coulombs/bit). Low byte of
	// 3 bytes.
	SurfDosL param.Opt[int64] `json:"surfDosL,omitzero"`
	// Surface dosimeter medium range output (converts to pico-coulombs/bit). Middle
	// byte of 3 bytes.
	SurfDosM param.Opt[int64] `json:"surfDosM,omitzero"`
	// Surface dosimeter status byte.
	SurfDosStat param.Opt[int64] `json:"surfDosStat,omitzero"`
	// Reference voltage (volts/bit). Conversion factor used to convert analog V
	// monitor data from bytes to volts.
	VRef param.Opt[int64] `json:"vRef,omitzero"`
	// Array of 144 digitized samples of ESD waveform for transient packets.
	TransientData []int64 `json:"transientData,omitzero"`
	paramObj
}

func (r ObservationEcpsdrNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationEcpsdrNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationEcpsdrNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationEcpsdrNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type ObservationEcpsdrTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time stamp of time packet receipt on ground, in ISO 8601 UTC format with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTime     time.Time        `query:"msgTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationEcpsdrTupleParams]'s query parameters as
// `url.Values`.
func (r ObservationEcpsdrTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationEcpsdrUnvalidatedPublishParams struct {
	Body []ObservationEcpsdrUnvalidatedPublishParamsBody
	paramObj
}

func (r ObservationEcpsdrUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationEcpsdrUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Supports Sensor Data Records (SDR) from space-borne Energetic Charged Particle
// (ECP) Sensors. SDR contains sensor status telemetry and raw dosimeter
// measurements of the space environment.
//
// The properties ClassificationMarking, DataMode, MsgTime, Source, Type are
// required.
type ObservationEcpsdrUnvalidatedPublishParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Time stamp of time packet receipt on ground, in ISO 8601 UTC format with
	// millisecond precision.
	MsgTime time.Time `json:"msgTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of data associated with this record (STANDARD, TRANSIENT).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Internal 5V current monitor for analog supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Asl5VCurrMon param.Opt[int64] `json:"asl5VCurrMon,omitzero"`
	// CDS Charge Plate voltage monitor. See vRef for conversion factor to Volts.
	CdsPlateVMon param.Opt[int64] `json:"cdsPlateVMon,omitzero"`
	// CDS reference voltage monitor. See vRef for conversion factor to Volts.
	CdsRefVMon param.Opt[int64] `json:"cdsRefVMon,omitzero"`
	// CDS Threshold setting for ESD detection threshold. The CDS Threshold is the
	// adjustable sensitivity of recording/digitizing an ESD as a transient packet.
	CdsThreshold param.Opt[int64] `json:"cdsThreshold,omitzero"`
	// CDS throttle number of seconds between CDS transient capture readouts.
	CdsThrottle param.Opt[int64] `json:"cdsThrottle,omitzero"`
	// Two byte CRC-16-CCITT checksum (ordered as first byte, second byte).
	Checksum param.Opt[int64] `json:"checksum,omitzero"`
	// Unitless dosimeter detector bias for MedLET and HiLET. MedLET (Linear Energy
	// Transfer) and HiLET subsensors detect particles above LET thresholds, 300keV and
	// 1MeV, respectively.
	DosBias param.Opt[int64] `json:"dosBias,omitzero"`
	// Internal 5V current monitor for digital supply line. This is sensor status
	// telemetry. See vRef for conversion factor to Volts.
	Dsl5VCurrMon param.Opt[int64] `json:"dsl5VCurrMon,omitzero"`
	// Number of ESD triggers, high byte of 2-byte counter.
	EsdTrigCountH param.Opt[int64] `json:"esdTrigCountH,omitzero"`
	// Number of ESD triggers, low byte of 2-byte counter.
	EsdTrigCountL param.Opt[int64] `json:"esdTrigCountL,omitzero"`
	// HiLET dosimeter low range output. Low byte of scaler (HiLET) dosimeter output.
	HiLetL param.Opt[int64] `json:"hiLetL,omitzero"`
	// Unitless HiLET dosimeter medium range output. Medium byte of (HiLET) dosimeter
	// output.
	HiLetM param.Opt[int64] `json:"hiLetM,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// LowLET dosimeter low range output. Low byte of (LowLET) dosimeter output.
	LowLetL param.Opt[int64] `json:"lowLetL,omitzero"`
	// LowLET dosimeter medium range output. Medium byte of (LowLET) dosimeter output.
	LowLetM param.Opt[int64] `json:"lowLetM,omitzero"`
	// MedLET1 dosimeter low range output. Low byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1L param.Opt[int64] `json:"medLet1L,omitzero"`
	// MedLET1 dosimeter medium range output. Medium byte of the 1st (MedLET) dosimeter
	// output.
	MedLet1M param.Opt[int64] `json:"medLet1M,omitzero"`
	// MedLET2 dosimeter low range output. Low byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2L param.Opt[int64] `json:"medLet2L,omitzero"`
	// MedLET2 dosimeter medium range output. Medium byte of the 2nd (MedLET) dosimeter
	// output.
	MedLet2M param.Opt[int64] `json:"medLet2M,omitzero"`
	// MedLET3 dosimeter low range output. Low byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3L param.Opt[int64] `json:"medLet3L,omitzero"`
	// MedLET3 dosimeter medium range output. Medium byte of the 3rd (MedLET) dosimeter
	// output.
	MedLet3M param.Opt[int64] `json:"medLet3M,omitzero"`
	// MedLET4 dosimeter low range output. Low byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4L param.Opt[int64] `json:"medLet4L,omitzero"`
	// MedLET4 dosimeter medium range output. Medium byte of the 4th (MedLET) dosimeter
	// output.
	MedLet4M param.Opt[int64] `json:"medLet4M,omitzero"`
	// Unitless sensor mounting plate temperature.
	MpTemp param.Opt[int64] `json:"mpTemp,omitzero"`
	// Time of the observation, in ISO 8601 UTC format with millisecond precision.
	ObTime param.Opt[time.Time] `json:"obTime,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source to indicate the satellite
	// hosting the sensor. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the record source to indicate the sensor
	// identifier which produced this data. This may be an internal identifier and not
	// necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Photodiode 1 signal level.
	Pd1SigLev param.Opt[int64] `json:"pd1SigLev,omitzero"`
	// Photodiode 2 signal level.
	Pd2SigLev param.Opt[int64] `json:"pd2SigLev,omitzero"`
	// Power supply temperature monitor. This is sensor status telemetry.
	PsTempMon param.Opt[int64] `json:"psTempMon,omitzero"`
	// Flag indicating whether this record is an original or re-transmitted dataset
	// (TRUE indicates a retransmit from the host).
	Retransmit param.Opt[bool] `json:"retransmit,omitzero"`
	// Satellite/catalog number of the on-orbit satellite hosting the sensor.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// The sensor mode associated with this measurements (NORMAL, TEST).
	SenMode param.Opt[string] `json:"senMode,omitzero"`
	// Surface dosimeter charge rate high output (converts to pico-amps/bit). High byte
	// of 2 bytes.
	SurfDosChargeH param.Opt[int64] `json:"surfDosChargeH,omitzero"`
	// Surface dosimeter charge rate low output (converts to pico-amps/bit). Low byte
	// of 2 bytes.
	SurfDosChargeL param.Opt[int64] `json:"surfDosChargeL,omitzero"`
	// Surface dosimeter high range output (converts to pico-coulombs/bit). High byte
	// of 3 bytes.
	SurfDosH param.Opt[int64] `json:"surfDosH,omitzero"`
	// Surface dosimeter low range output (converts to pico-coulombs/bit). Low byte of
	// 3 bytes.
	SurfDosL param.Opt[int64] `json:"surfDosL,omitzero"`
	// Surface dosimeter medium range output (converts to pico-coulombs/bit). Middle
	// byte of 3 bytes.
	SurfDosM param.Opt[int64] `json:"surfDosM,omitzero"`
	// Surface dosimeter status byte.
	SurfDosStat param.Opt[int64] `json:"surfDosStat,omitzero"`
	// Reference voltage (volts/bit). Conversion factor used to convert analog V
	// monitor data from bytes to volts.
	VRef param.Opt[int64] `json:"vRef,omitzero"`
	// Array of 144 digitized samples of ESD waveform for transient packets.
	TransientData []int64 `json:"transientData,omitzero"`
	paramObj
}

func (r ObservationEcpsdrUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationEcpsdrUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationEcpsdrUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationEcpsdrUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
