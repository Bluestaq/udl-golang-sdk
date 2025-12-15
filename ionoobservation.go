// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// IonoObservationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIonoObservationService] method instead.
type IonoObservationService struct {
	Options []option.RequestOption
	History IonoObservationHistoryService
}

// NewIonoObservationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIonoObservationService(opts ...option.RequestOption) (r IonoObservationService) {
	r = IonoObservationService{}
	r.Options = opts
	r.History = NewIonoObservationHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *IonoObservationService) List(ctx context.Context, query IonoObservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[IonoObservationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/ionoobservation"
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
func (r *IonoObservationService) ListAutoPaging(ctx context.Context, query IonoObservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[IonoObservationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *IonoObservationService) Count(ctx context.Context, query IonoObservationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/ionoobservation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// IonoObservation records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *IonoObservationService) NewBulk(ctx context.Context, body IonoObservationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/ionoobservation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *IonoObservationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *IonoObservationQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/ionoobservation/queryhelp"
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
func (r *IonoObservationService) Tuple(ctx context.Context, query IonoObservationTupleParams, opts ...option.RequestOption) (res *[]IonoObservationTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/ionoobservation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take Ionospheric Observation entries as a POST body and
// ingest into the database with or without dupe detection. Default is no dupe
// checking. This operation is intended to be used for automated feeds into UDL. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *IonoObservationService) UnvalidatedPublish(ctx context.Context, body IonoObservationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-ionoobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// These services provide operations for posting and querying ionospheric
// observation data. Characteristics are defined by the CHARS: URSI IIWG format for
// archiving monthly ionospheric characteristics, INAG Bulletin No. 62
// specification. Qualifying and Descriptive letters are defined by the URSI
// Handbook for Ionogram Interpretation and reduction, Report UAG, No. 23A
// specification.
type IonoObservationListResponse struct {
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
	DataMode IonoObservationListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Sounding Start time in ISO8601 UTC format.
	StartTimeUtc time.Time `json:"startTimeUTC,required" format:"date-time"`
	// URSI code for station or stations producing the ionosonde.
	StationID string `json:"stationId,required"`
	// Ionosonde hardware type or data collection type together with possible
	// additional descriptors.
	System string `json:"system,required"`
	// Names of settings.
	SystemInfo string `json:"systemInfo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID                     string                                            `json:"id"`
	Amplitude              IonoObservationListResponseAmplitude              `json:"amplitude"`
	AntennaElementPosition IonoObservationListResponseAntennaElementPosition `json:"antennaElementPosition"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem IonoObservationListResponseAntennaElementPositionCoordinateSystem `json:"antennaElementPositionCoordinateSystem"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                            `json:"artistFlags"`
	Azimuth     IonoObservationListResponseAzimuth `json:"azimuth"`
	// IRI thickness parameter in km. URSI ID: D0.
	B0 float64 `json:"b0"`
	// IRI profile shape parameter. URSI ID: D1.
	B1 float64 `json:"b1"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonoObservationListResponseCharAtt `json:"charAtts"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Distance for MUF calculation in km.
	D float64 `json:"d"`
	// IRI profile shape parameter, F1 layer. URSI ID: D2.
	D1    float64                          `json:"d1"`
	Datum IonoObservationListResponseDatum `json:"datum"`
	// Adjustment to the scaled foF2 during profile inversion in MHz.
	DeltafoF2 float64 `json:"deltafoF2"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonoObservationListResponseDensityProfile `json:"densityProfile"`
	Doppler        IonoObservationListResponseDoppler        `json:"doppler"`
	// Lowering of E trace to the leading edge in km.
	DownE float64 `json:"downE"`
	// Lowering of Es trace to the leading edge in km.
	DownEs float64 `json:"downEs"`
	// Lowering of F trace to the leading edge in km.
	DownF float64 `json:"downF"`
	// Array of electron densities in cm^-3 (must match the size of the height and
	// plasmaFrequency arrays).
	ElectronDensity []float64 `json:"electronDensity"`
	// Uncertainty in specifying the electron density at each height point of the
	// profile (must match the size of the electronDensity array).
	ElectronDensityUncertainty []float64                            `json:"electronDensityUncertainty"`
	Elevation                  IonoObservationListResponseElevation `json:"elevation"`
	// The blanketing frequency of layer used to derive foEs in MHz. URSI ID: 32.
	FbEs float64 `json:"fbEs"`
	// Frequency spread beyond foE in MHz. URSI ID: 87.
	Fe float64 `json:"fe"`
	// Frequency spread between fxF2 and FxI in MHz. URSI ID: 86.
	Ff float64 `json:"ff"`
	// The frequency at which hprimeF is measured in MHz. URSI ID: 61.
	FhprimeF float64 `json:"fhprimeF"`
	// The frequency at which hprimeF2 is measured in MHz. URSI ID: 60.
	FhprimeF2 float64 `json:"fhprimeF2"`
	// Lowest frequency at which echo traces are observed on the ionogram, in MHz. URSI
	// ID: 42.
	Fmin float64 `json:"fmin"`
	// Minimum frequency of E layer echoes in MHz. URSI ID: 81.
	FminE float64 `json:"fminE"`
	// Minimum frequency of Es layer in MHz.
	FminEs float64 `json:"fminEs"`
	// Minimum frequency of F layer echoes in MHz. URSI ID: 80.
	FminF float64 `json:"fminF"`
	// MUF/OblFactor in MHz.
	Fmuf float64 `json:"fmuf"`
	// The ordinary wave critical frequency of the lowest thick layer which causes a
	// discontinuity, in MHz. URSI ID: 20.
	FoE float64 `json:"foE"`
	// Critical frequency of night time auroral E layer in MHz. URSI ID: 23.
	FoEa float64 `json:"foEa"`
	// Predicted value of foE in MHz.
	FoEp float64 `json:"foEp"`
	// Highest ordinary wave frequency at which a mainly continuous Es trace is
	// observed, in MHz. URSI ID: 30.
	FoEs float64 `json:"foEs"`
	// The ordinary wave F1 critical frequency, in MHz. URSI ID: 10.
	FoF1 float64 `json:"foF1"`
	// Predicted value of foF1 in MHz.
	FoF1p float64 `json:"foF1p"`
	// The ordinary wave critical frequency of the highest stratification in the F
	// region, specified in MHz. URSI ID: 00.
	FoF2 float64 `json:"foF2"`
	// Predicted value of foF2 in MHz.
	FoF2p float64 `json:"foF2p"`
	// Highest ordinary wave critical frequency of F region patch trace in MHz. URSI
	// ID: 55.
	FoP       float64                              `json:"foP"`
	Frequency IonoObservationListResponseFrequency `json:"frequency"`
	// The extraordinary wave E critical frequency, in MHz. URSI ID: 21.
	FxE float64 `json:"fxE"`
	// The extraordinary wave F1 critical frequency, in MHz. URSI ID: 11.
	FxF1 float64 `json:"fxF1"`
	// The extraordinary wave F2 critical frequency, in MHz. URSI ID: 01.
	FxF2 float64 `json:"fxF2"`
	// The highest frequency of F-trace in MHz. Note: fxI is with capital i. URSI
	// ID: 51.
	FxI float64 `json:"fxI"`
	// Array of altitudes above station level for plasma frequency/density arrays in km
	// (must match the size of the plasmaFrequency and electronDensity Arrays).
	Height []float64 `json:"height"`
	// True height of the E peak in km. URSI ID: CE.
	HmE float64 `json:"hmE"`
	// True height of the F1 peak in km. URSI ID: BE.
	HmF1 float64 `json:"hmF1"`
	// True height of the F2 peak in km. URSI ID: AE.
	HmF2 float64 `json:"hmF2"`
	// The minimum virtual height of the normal E layer trace in km. URSI ID: 24.
	HprimeE float64 `json:"hprimeE"`
	// Minimum virtual height of night time auroral E layer trace in km. URSI ID: 27.
	HprimeEa float64 `json:"hprimeEa"`
	// The minimum height of the trace used to give foEs in km. URSI ID: 34.
	HprimeEs float64 `json:"hprimeEs"`
	// The minimum virtual height of the ordinary wave trace taken as a whole, in km.
	// URSI ID: 16.
	HprimeF float64 `json:"hprimeF"`
	// The minimum virtual height of reflection at a point where the trace is
	// horizontal in the F region in km. URSI ID: 14.
	HprimeF1 float64 `json:"hprimeF1"`
	// The minimum virtual height of ordinary wave trace for the highest stable
	// stratification in the F region in km. URSI ID: 4.
	HprimeF2 float64 `json:"hprimeF2"`
	// Virtual height at MUF/OblFactor frequency in MHz.
	HprimefMuf float64 `json:"hprimefMUF"`
	// Minimum virtual height of the trace used to determine foP in km. URSI ID: 56.
	HprimeP float64 `json:"hprimeP"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// Lowest usable frequency.
	Luf float64 `json:"luf"`
	// MUF(D)/foF2.
	Md float64 `json:"md"`
	// Maximum Usable Frequency for ground distance D in MHz.
	Mufd float64 `json:"mufd"`
	// Name of the algorithm used for the electron density profile.
	NeProfileName string `json:"neProfileName"`
	// Version of the algorithm used for the electron density profile.
	NeProfileVersion float64 `json:"neProfileVersion"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID string                           `json:"origSensorId"`
	Phase        IonoObservationListResponsePhase `json:"phase"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64 `json:"plasmaFrequencyUncertainty"`
	// Equipment location.
	PlatformName string                                  `json:"platformName"`
	Polarization IonoObservationListResponsePolarization `json:"polarization"`
	Power        IonoObservationListResponsePower        `json:"power"`
	// Average range spread of E layer in km. URSI ID: 85.
	Qe float64 `json:"qe"`
	// Average range spread of F layer in km. URSI ID: 84.
	Qf    float64                          `json:"qf"`
	Range IonoObservationListResponseRange `json:"range"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	ReceiveCoordinates [][]float64 `json:"receiveCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	ReceiveSensorType IonoObservationListResponseReceiveSensorType `json:"receiveSensorType"`
	// Array of restricted frequencies.
	RestrictedFrequency []float64 `json:"restrictedFrequency"`
	// Notes for the restrictedFrequency data.
	RestrictedFrequencyNotes string `json:"restrictedFrequencyNotes"`
	// Effective scale height at hmF2 Titheridge method in km. URSI ID: 69.
	ScaleHeightF2Peak float64 `json:"scaleHeightF2Peak"`
	// The ScalerInfo record describes the person or system who interpreted the
	// ionogram in IonoObservation.
	ScalerInfo IonoObservationListResponseScalerInfo `json:"scalerInfo"`
	Stokes     IonoObservationListResponseStokes     `json:"stokes"`
	// Details concerning the composition/intention/interpretation/audience/etc. of any
	// data recorded here. This field may contain all of the intended information e.g.
	// info on signal waveforms used, antenna setup, etc. OR may describe the
	// data/settings to be provided in the “data” field.
	SystemNotes string `json:"systemNotes"`
	// Total Ionospheric Electron Content \*10^16e/m^2. 1 TEC Unit (TECU) = 10^16
	// electrons/m^2. URSI ID: 72.
	Tec float64 `json:"tec"`
	// Array of degrees clockwise from true North of the TID.
	TidAzimuth []float64 `json:"tidAzimuth"`
	// Array of 1/frequency of the TID wave.
	TidPeriods []float64 `json:"tidPeriods"`
	// Array of speed in m/s at which the disturbance travels through the ionosphere.
	TidPhaseSpeeds []float64                               `json:"tidPhaseSpeeds"`
	Time           IonoObservationListResponseTime         `json:"time"`
	TraceGeneric   IonoObservationListResponseTraceGeneric `json:"traceGeneric"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	TransmitCoordinates [][]float64 `json:"transmitCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	TransmitSensorType IonoObservationListResponseTransmitSensorType `json:"transmitSensorType"`
	// Characterization of the shape of Es trace. URSI ID: 36.
	TypeEs string `json:"typeEs"`
	// Time the row was updated in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Parabolic E layer semi-thickness in km. URSI ID: 83.
	YE float64 `json:"yE"`
	// Parabolic F1 layer semi-thickness in km. URSI ID: 95.
	YF1 float64 `json:"yF1"`
	// Parabolic F2 layer semi-thickness in km. URSI ID: 94.
	YF2 float64 `json:"yF2"`
	// True height at half peak electron density in the F2 layer in km. URSI ID: 93.
	ZhalfNm float64 `json:"zhalfNm"`
	// Peak height of E-layer in km. URSI ID: 90.
	ZmE float64 `json:"zmE"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                  respjson.Field
		DataMode                               respjson.Field
		Source                                 respjson.Field
		StartTimeUtc                           respjson.Field
		StationID                              respjson.Field
		System                                 respjson.Field
		SystemInfo                             respjson.Field
		ID                                     respjson.Field
		Amplitude                              respjson.Field
		AntennaElementPosition                 respjson.Field
		AntennaElementPositionCoordinateSystem respjson.Field
		ArtistFlags                            respjson.Field
		Azimuth                                respjson.Field
		B0                                     respjson.Field
		B1                                     respjson.Field
		CharAtts                               respjson.Field
		CreatedAt                              respjson.Field
		CreatedBy                              respjson.Field
		D                                      respjson.Field
		D1                                     respjson.Field
		Datum                                  respjson.Field
		DeltafoF2                              respjson.Field
		DensityProfile                         respjson.Field
		Doppler                                respjson.Field
		DownE                                  respjson.Field
		DownEs                                 respjson.Field
		DownF                                  respjson.Field
		ElectronDensity                        respjson.Field
		ElectronDensityUncertainty             respjson.Field
		Elevation                              respjson.Field
		FbEs                                   respjson.Field
		Fe                                     respjson.Field
		Ff                                     respjson.Field
		FhprimeF                               respjson.Field
		FhprimeF2                              respjson.Field
		Fmin                                   respjson.Field
		FminE                                  respjson.Field
		FminEs                                 respjson.Field
		FminF                                  respjson.Field
		Fmuf                                   respjson.Field
		FoE                                    respjson.Field
		FoEa                                   respjson.Field
		FoEp                                   respjson.Field
		FoEs                                   respjson.Field
		FoF1                                   respjson.Field
		FoF1p                                  respjson.Field
		FoF2                                   respjson.Field
		FoF2p                                  respjson.Field
		FoP                                    respjson.Field
		Frequency                              respjson.Field
		FxE                                    respjson.Field
		FxF1                                   respjson.Field
		FxF2                                   respjson.Field
		FxI                                    respjson.Field
		Height                                 respjson.Field
		HmE                                    respjson.Field
		HmF1                                   respjson.Field
		HmF2                                   respjson.Field
		HprimeE                                respjson.Field
		HprimeEa                               respjson.Field
		HprimeEs                               respjson.Field
		HprimeF                                respjson.Field
		HprimeF1                               respjson.Field
		HprimeF2                               respjson.Field
		HprimefMuf                             respjson.Field
		HprimeP                                respjson.Field
		IDSensor                               respjson.Field
		Luf                                    respjson.Field
		Md                                     respjson.Field
		Mufd                                   respjson.Field
		NeProfileName                          respjson.Field
		NeProfileVersion                       respjson.Field
		Origin                                 respjson.Field
		OrigNetwork                            respjson.Field
		OrigSensorID                           respjson.Field
		Phase                                  respjson.Field
		PlasmaFrequency                        respjson.Field
		PlasmaFrequencyUncertainty             respjson.Field
		PlatformName                           respjson.Field
		Polarization                           respjson.Field
		Power                                  respjson.Field
		Qe                                     respjson.Field
		Qf                                     respjson.Field
		Range                                  respjson.Field
		ReceiveCoordinates                     respjson.Field
		ReceiveSensorType                      respjson.Field
		RestrictedFrequency                    respjson.Field
		RestrictedFrequencyNotes               respjson.Field
		ScaleHeightF2Peak                      respjson.Field
		ScalerInfo                             respjson.Field
		Stokes                                 respjson.Field
		SystemNotes                            respjson.Field
		Tec                                    respjson.Field
		TidAzimuth                             respjson.Field
		TidPeriods                             respjson.Field
		TidPhaseSpeeds                         respjson.Field
		Time                                   respjson.Field
		TraceGeneric                           respjson.Field
		TransmitCoordinates                    respjson.Field
		TransmitSensorType                     respjson.Field
		TypeEs                                 respjson.Field
		UpdatedAt                              respjson.Field
		UpdatedBy                              respjson.Field
		YE                                     respjson.Field
		YF1                                    respjson.Field
		YF2                                    respjson.Field
		ZhalfNm                                respjson.Field
		ZmE                                    respjson.Field
		ExtraFields                            map[string]respjson.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponse) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponse) UnmarshalJSON(data []byte) error {
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
type IonoObservationListResponseDataMode string

const (
	IonoObservationListResponseDataModeReal      IonoObservationListResponseDataMode = "REAL"
	IonoObservationListResponseDataModeTest      IonoObservationListResponseDataMode = "TEST"
	IonoObservationListResponseDataModeSimulated IonoObservationListResponseDataMode = "SIMULATED"
	IonoObservationListResponseDataModeExercise  IonoObservationListResponseDataMode = "EXERCISE"
)

type IonoObservationListResponseAmplitude struct {
	// Array of amplitude data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for amplitude dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the amplitude data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseAmplitude) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseAmplitude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponseAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseAntennaElementPosition) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseAntennaElementPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
// (lat, long, radii).
type IonoObservationListResponseAntennaElementPositionCoordinateSystem string

const (
	IonoObservationListResponseAntennaElementPositionCoordinateSystemJ2000                   IonoObservationListResponseAntennaElementPositionCoordinateSystem = "J2000"
	IonoObservationListResponseAntennaElementPositionCoordinateSystemEcrEcef                 IonoObservationListResponseAntennaElementPositionCoordinateSystem = "ECR/ECEF"
	IonoObservationListResponseAntennaElementPositionCoordinateSystemTeme                    IonoObservationListResponseAntennaElementPositionCoordinateSystem = "TEME"
	IonoObservationListResponseAntennaElementPositionCoordinateSystemGcrf                    IonoObservationListResponseAntennaElementPositionCoordinateSystem = "GCRF"
	IonoObservationListResponseAntennaElementPositionCoordinateSystemWgs84GeoDeticLatLongAlt IonoObservationListResponseAntennaElementPositionCoordinateSystem = "WGS84 (GEODetic lat, long, alt)"
	IonoObservationListResponseAntennaElementPositionCoordinateSystemGeoCentricLatLongRadii  IonoObservationListResponseAntennaElementPositionCoordinateSystem = "GEOCentric (lat, long, radii)"
)

type IonoObservationListResponseAzimuth struct {
	// Array of incoming azimuth at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the azimuth array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the azimuth data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseAzimuth) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseAzimuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Characteristic attributes of a IonoObservation.
type IonoObservationListResponseCharAtt struct {
	// Characteristic name. This value should reflect the UDL field name for the
	// corresponding characteristic.
	CharName string `json:"charName"`
	// Input parameters for the climate model.
	ClimateModelInputParams []string `json:"climateModelInputParams"`
	// Name of the climate model.
	ClimateModelName string `json:"climateModelName"`
	// List of options for the climate model.
	ClimateModelOptions []string `json:"climateModelOptions"`
	// Descriptive letter (D) for the characteristic specified by URSI ID. Describes
	// specific ionospheric conditions, beyond numerical values.
	D string `json:"d" format:"char"`
	// Specified characteristic's lower bound. Should be less than or equal to the
	// characteristic's current value as set in this record.
	LowerBound float64 `json:"lowerBound"`
	// Qualifying letter (Q) for the characteristic specified by URSI ID. Describes
	// specific ionospheric conditions, beyond numerical values.
	Q string `json:"q" format:"char"`
	// Uncertainty Bounds (lower and upper) define an interval around reported value
	// that contains true value at the specified probability level. Probability levels
	// are specified in terms of percentile (from 1 to 100) or the standard deviation,
	// sigma (e.g. 1sigma, 2sigma, 3sigma, 5percentile, 10percentile, 25percentile).
	UncertaintyBoundType string `json:"uncertaintyBoundType"`
	// Specified characteristic's upper bound. Should be greater than or equal to the
	// characteristic's current value as set in this record.
	UpperBound float64 `json:"upperBound"`
	// Characteristic's URSI ID. See the characteristic's description for its
	// corresponding URSI ID.
	UrsiID string `json:"ursiID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharName                respjson.Field
		ClimateModelInputParams respjson.Field
		ClimateModelName        respjson.Field
		ClimateModelOptions     respjson.Field
		D                       respjson.Field
		LowerBound              respjson.Field
		Q                       respjson.Field
		UncertaintyBoundType    respjson.Field
		UpperBound              respjson.Field
		UrsiID                  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseCharAtt) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseCharAtt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponseDatum struct {
	// Array to support sparse data collections.
	Data []float64 `json:"data"`
	// Notes for the datum with details of what the data is, units, etc.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Notes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDatum) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseDatum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonoObservationListResponseDensityProfile struct {
	// Full set of the IRI formalism coefficients.
	Iri IonoObservationListResponseDensityProfileIri `json:"iri"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonoObservationListResponseDensityProfileParabolic `json:"parabolic"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonoObservationListResponseDensityProfileQuasiParabolic `json:"quasiParabolic"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonoObservationListResponseDensityProfileShiftedChebyshev `json:"shiftedChebyshev"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonoObservationListResponseDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonoObservationListResponseDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs"`
	// Description of the valley model and parameters.
	ValleyModelDescription string `json:"valleyModelDescription"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iri                          respjson.Field
		Parabolic                    respjson.Field
		QuasiParabolic               respjson.Field
		ShiftedChebyshev             respjson.Field
		TopsideExtensionChapmanConst respjson.Field
		TopsideExtensionVaryChap     respjson.Field
		ValleyModelCoeffs            respjson.Field
		ValleyModelDescription       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfile) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseDensityProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full set of the IRI formalism coefficients.
type IonoObservationListResponseDensityProfileIri struct {
	// B0 parameter of the F2 layer shape.
	B0 float64 `json:"b0"`
	// B1 parameter of the F2 layer shape.
	B1 float64 `json:"b1"`
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi float64 `json:"chi"`
	// D1 parameter of the F1 layer shape.
	D1 float64 `json:"d1"`
	// Description of IRI implementation.
	Description string `json:"description"`
	// TBD.
	Fp1 float64 `json:"fp1"`
	// TBD.
	Fp2 float64 `json:"fp2"`
	// TBD.
	Fp30 float64 `json:"fp30"`
	// TBD.
	Fp3U float64 `json:"fp3U"`
	// Starting height of the D layer, in kilometers.
	Ha float64 `json:"ha"`
	// Height of the intermediate region at the top of D region, in kilometers.
	Hdx float64 `json:"hdx"`
	// Peak height of the D layer, in kilometers.
	HmD float64 `json:"hmD"`
	// Peak height of the F2 layer, in kilometers.
	HmE float64 `json:"hmE"`
	// Peak height of the F1 layer, in kilometers.
	HmF1 float64 `json:"hmF1"`
	// Peak height of F2 layer, in kilometers.
	HmF2 float64 `json:"hmF2"`
	// The valley height, in kilometers.
	HValTop float64 `json:"hValTop"`
	// Height HZ of the interim layer, in kilometers.
	Hz float64 `json:"hz"`
	// Peak density of the D layer, in per cubic centimeter.
	NmD float64 `json:"nmD"`
	// Peak density of the E layer, in per cubic centimeter.
	NmE float64 `json:"nmE"`
	// Peak density of the F1 layer, in grams per cubic centimeter.
	NmF1 float64 `json:"nmF1"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 float64 `json:"nmF2"`
	// The valley depth, in grams per cubic centimeter.
	NValB float64 `json:"nValB"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		B0          respjson.Field
		B1          respjson.Field
		Chi         respjson.Field
		D1          respjson.Field
		Description respjson.Field
		Fp1         respjson.Field
		Fp2         respjson.Field
		Fp30        respjson.Field
		Fp3U        respjson.Field
		Ha          respjson.Field
		Hdx         respjson.Field
		HmD         respjson.Field
		HmE         respjson.Field
		HmF1        respjson.Field
		HmF2        respjson.Field
		HValTop     respjson.Field
		Hz          respjson.Field
		NmD         respjson.Field
		NmE         respjson.Field
		NmF1        respjson.Field
		NmF2        respjson.Field
		NValB       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfileIri) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseDensityProfileIri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationListResponseDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description string `json:"description"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonoObservationListResponseDensityProfileParabolicParabolicItem `json:"parabolicItems"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		ParabolicItems respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfileParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseDensityProfileParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationListResponseDensityProfileParabolicParabolicItem struct {
	// Plasma frequency at the layer peak, in MHz.
	F float64 `json:"f"`
	// Ionospheric plasma layer (E, F1, or F2).
	Layer string `json:"layer"`
	// Half-thickness of the layer, in kilometers.
	Y float64 `json:"y"`
	// Height of the layer peak, in kilometers.
	Z float64 `json:"z"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		F           respjson.Field
		Layer       respjson.Field
		Y           respjson.Field
		Z           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfileParabolicParabolicItem) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationListResponseDensityProfileParabolicParabolicItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonoObservationListResponseDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description string `json:"description"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius float64 `json:"earthRadius"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonoObservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description            respjson.Field
		EarthRadius            respjson.Field
		QuasiParabolicSegments respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfileQuasiParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseDensityProfileQuasiParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonoObservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment struct {
	// Coefficient A.
	A float64 `json:"a"`
	// Coefficient B.
	B float64 `json:"b"`
	// Coefficient C.
	C float64 `json:"c"`
	// Best-fit error.
	Error float64 `json:"error"`
	// The index of this segment in the list, from 1 to NumSegments.
	Index int64 `json:"index"`
	// Starting range of the segment, in kilometers from the Earth's center.
	Rb float64 `json:"rb"`
	// Ending range of the segment, in kilometers from the Earth's center.
	Re float64 `json:"re"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		A           respjson.Field
		B           respjson.Field
		C           respjson.Field
		Error       respjson.Field
		Index       respjson.Field
		Rb          respjson.Field
		Re          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonoObservationListResponseDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description string `json:"description"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonoObservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description       respjson.Field
		ShiftedChebyshevs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfileShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationListResponseDensityProfileShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonoObservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev struct {
	// Array of coefficients.
	Coeffs []float64 `json:"coeffs"`
	// Best fit error.
	Error float64 `json:"error"`
	// Start frequency of the layer, in MHz.
	Fstart float64 `json:"fstart"`
	// Stop frequency of the layer, in MHz.
	Fstop float64 `json:"fstop"`
	// The ionospheric plasma layer.
	Layer string `json:"layer"`
	// Number of coefficients in the expansion.
	N int64 `json:"n"`
	// Peak height of the layer, in kilometers.
	Peakheight float64 `json:"peakheight"`
	// Height at which density is half of the peak Nm, in kilometers.
	ZhalfNm float64 `json:"zhalfNm"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coeffs      respjson.Field
		Error       respjson.Field
		Fstart      respjson.Field
		Fstop       respjson.Field
		Layer       respjson.Field
		N           respjson.Field
		Peakheight  respjson.Field
		ZhalfNm     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters of the constant-scale-height Chapman layer.
type IonoObservationListResponseDensityProfileTopsideExtensionChapmanConst struct {
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi float64 `json:"chi"`
	// Description of the Chapman computation technique.
	Description string `json:"description"`
	// Peak height of F2 layer, in kilometers.
	HmF2 float64 `json:"hmF2"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 float64 `json:"nmF2"`
	// Scale height if F2 layer at the peak, in kilometers.
	ScaleF2 float64 `json:"scaleF2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chi         respjson.Field
		Description respjson.Field
		HmF2        respjson.Field
		NmF2        respjson.Field
		ScaleF2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfileTopsideExtensionChapmanConst) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationListResponseDensityProfileTopsideExtensionChapmanConst) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Varying scale height Chapman topside layer.
type IonoObservationListResponseDensityProfileTopsideExtensionVaryChap struct {
	// Alpha parameter of the profile shape.
	Alpha float64 `json:"alpha"`
	// Beta parameter of the profile shape.
	Beta float64 `json:"beta"`
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi float64 `json:"chi"`
	// Description of the Chapman computation technique.
	Description string `json:"description"`
	// Peak height of F2 layer, in kilometers.
	HmF2 float64 `json:"hmF2"`
	// Transition height, in kilometers.
	Ht float64 `json:"ht"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 float64 `json:"nmF2"`
	// Scale height if F2 layer at the peak, in kilometers.
	ScaleF2 float64 `json:"scaleF2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alpha       respjson.Field
		Beta        respjson.Field
		Chi         respjson.Field
		Description respjson.Field
		HmF2        respjson.Field
		Ht          respjson.Field
		NmF2        respjson.Field
		ScaleF2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDensityProfileTopsideExtensionVaryChap) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationListResponseDensityProfileTopsideExtensionVaryChap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponseDoppler struct {
	// Array of received doppler shifts in Hz.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the doppler array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the doppler data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseDoppler) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseDoppler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponseElevation struct {
	// Array of incoming elevation at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the elevation array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the elevation data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseElevation) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseElevation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponseFrequency struct {
	// Array of frequency data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for frequency dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the frequency array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the frequency data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseFrequency) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseFrequency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponsePhase struct {
	// Array of phase data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for phase dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the phase array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the phase data. Orientation and position for each antenna element
	// across the antenna_element dimension.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponsePhase) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponsePhase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponsePolarization struct {
	// Array of polarization data.
	Data [][][][][][][]string `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for polarization dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the polarization data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponsePolarization) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponsePolarization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponsePower struct {
	// Array of received power in db.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the power array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the power data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponsePower) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponsePower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponseRange struct {
	// Array of range data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for range dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the range array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the range data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseRange) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonoObservationListResponseReceiveSensorType string

const (
	IonoObservationListResponseReceiveSensorTypeMobile IonoObservationListResponseReceiveSensorType = "Mobile"
	IonoObservationListResponseReceiveSensorTypeStatic IonoObservationListResponseReceiveSensorType = "Static"
)

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonoObservationListResponseScalerInfo struct {
	// Scaler confidence level.
	ConfidenceLevel int64 `json:"confidenceLevel"`
	// Scaler confidence score.
	ConfidenceScore int64 `json:"confidenceScore"`
	// Scaler name.
	Name string `json:"name"`
	// Scaler organization.
	Organization string `json:"organization"`
	// Scaler type (MANUAL, AUTOMATIC or UNKNOWN).
	Type string `json:"type"`
	// Scaler version.
	Version float64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfidenceLevel respjson.Field
		ConfidenceScore respjson.Field
		Name            respjson.Field
		Organization    respjson.Field
		Type            respjson.Field
		Version         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseScalerInfo) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseScalerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponseStokes struct {
	// Array of received stokes data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the stoke array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the stokes data.
	Notes string `json:"notes"`
	// S1, S2, and S3 (the normalized Stokes parameters 1, 2, and 3).
	S []float64 `json:"s"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		S             respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseStokes) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseStokes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponseTime struct {
	// Array of times in number of seconds passed since January 1st, 1970 with the same
	// dimensions as power.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the time array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// The notes indicate the scheme and accuracy.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseTime) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationListResponseTraceGeneric struct {
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName"`
	// Notes for the trace generic data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationListResponseTraceGeneric) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationListResponseTraceGeneric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonoObservationListResponseTransmitSensorType string

const (
	IonoObservationListResponseTransmitSensorTypeMobile IonoObservationListResponseTransmitSensorType = "Mobile"
	IonoObservationListResponseTransmitSensorTypeStatic IonoObservationListResponseTransmitSensorType = "Static"
)

type IonoObservationQueryhelpResponse struct {
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
func (r IonoObservationQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These services provide operations for posting and querying ionospheric
// observation data. Characteristics are defined by the CHARS: URSI IIWG format for
// archiving monthly ionospheric characteristics, INAG Bulletin No. 62
// specification. Qualifying and Descriptive letters are defined by the URSI
// Handbook for Ionogram Interpretation and reduction, Report UAG, No. 23A
// specification.
type IonoObservationTupleResponse struct {
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
	DataMode IonoObservationTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Sounding Start time in ISO8601 UTC format.
	StartTimeUtc time.Time `json:"startTimeUTC,required" format:"date-time"`
	// URSI code for station or stations producing the ionosonde.
	StationID string `json:"stationId,required"`
	// Ionosonde hardware type or data collection type together with possible
	// additional descriptors.
	System string `json:"system,required"`
	// Names of settings.
	SystemInfo string `json:"systemInfo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID                     string                                             `json:"id"`
	Amplitude              IonoObservationTupleResponseAmplitude              `json:"amplitude"`
	AntennaElementPosition IonoObservationTupleResponseAntennaElementPosition `json:"antennaElementPosition"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem IonoObservationTupleResponseAntennaElementPositionCoordinateSystem `json:"antennaElementPositionCoordinateSystem"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                             `json:"artistFlags"`
	Azimuth     IonoObservationTupleResponseAzimuth `json:"azimuth"`
	// IRI thickness parameter in km. URSI ID: D0.
	B0 float64 `json:"b0"`
	// IRI profile shape parameter. URSI ID: D1.
	B1 float64 `json:"b1"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonoObservationTupleResponseCharAtt `json:"charAtts"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Distance for MUF calculation in km.
	D float64 `json:"d"`
	// IRI profile shape parameter, F1 layer. URSI ID: D2.
	D1    float64                           `json:"d1"`
	Datum IonoObservationTupleResponseDatum `json:"datum"`
	// Adjustment to the scaled foF2 during profile inversion in MHz.
	DeltafoF2 float64 `json:"deltafoF2"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonoObservationTupleResponseDensityProfile `json:"densityProfile"`
	Doppler        IonoObservationTupleResponseDoppler        `json:"doppler"`
	// Lowering of E trace to the leading edge in km.
	DownE float64 `json:"downE"`
	// Lowering of Es trace to the leading edge in km.
	DownEs float64 `json:"downEs"`
	// Lowering of F trace to the leading edge in km.
	DownF float64 `json:"downF"`
	// Array of electron densities in cm^-3 (must match the size of the height and
	// plasmaFrequency arrays).
	ElectronDensity []float64 `json:"electronDensity"`
	// Uncertainty in specifying the electron density at each height point of the
	// profile (must match the size of the electronDensity array).
	ElectronDensityUncertainty []float64                             `json:"electronDensityUncertainty"`
	Elevation                  IonoObservationTupleResponseElevation `json:"elevation"`
	// The blanketing frequency of layer used to derive foEs in MHz. URSI ID: 32.
	FbEs float64 `json:"fbEs"`
	// Frequency spread beyond foE in MHz. URSI ID: 87.
	Fe float64 `json:"fe"`
	// Frequency spread between fxF2 and FxI in MHz. URSI ID: 86.
	Ff float64 `json:"ff"`
	// The frequency at which hprimeF is measured in MHz. URSI ID: 61.
	FhprimeF float64 `json:"fhprimeF"`
	// The frequency at which hprimeF2 is measured in MHz. URSI ID: 60.
	FhprimeF2 float64 `json:"fhprimeF2"`
	// Lowest frequency at which echo traces are observed on the ionogram, in MHz. URSI
	// ID: 42.
	Fmin float64 `json:"fmin"`
	// Minimum frequency of E layer echoes in MHz. URSI ID: 81.
	FminE float64 `json:"fminE"`
	// Minimum frequency of Es layer in MHz.
	FminEs float64 `json:"fminEs"`
	// Minimum frequency of F layer echoes in MHz. URSI ID: 80.
	FminF float64 `json:"fminF"`
	// MUF/OblFactor in MHz.
	Fmuf float64 `json:"fmuf"`
	// The ordinary wave critical frequency of the lowest thick layer which causes a
	// discontinuity, in MHz. URSI ID: 20.
	FoE float64 `json:"foE"`
	// Critical frequency of night time auroral E layer in MHz. URSI ID: 23.
	FoEa float64 `json:"foEa"`
	// Predicted value of foE in MHz.
	FoEp float64 `json:"foEp"`
	// Highest ordinary wave frequency at which a mainly continuous Es trace is
	// observed, in MHz. URSI ID: 30.
	FoEs float64 `json:"foEs"`
	// The ordinary wave F1 critical frequency, in MHz. URSI ID: 10.
	FoF1 float64 `json:"foF1"`
	// Predicted value of foF1 in MHz.
	FoF1p float64 `json:"foF1p"`
	// The ordinary wave critical frequency of the highest stratification in the F
	// region, specified in MHz. URSI ID: 00.
	FoF2 float64 `json:"foF2"`
	// Predicted value of foF2 in MHz.
	FoF2p float64 `json:"foF2p"`
	// Highest ordinary wave critical frequency of F region patch trace in MHz. URSI
	// ID: 55.
	FoP       float64                               `json:"foP"`
	Frequency IonoObservationTupleResponseFrequency `json:"frequency"`
	// The extraordinary wave E critical frequency, in MHz. URSI ID: 21.
	FxE float64 `json:"fxE"`
	// The extraordinary wave F1 critical frequency, in MHz. URSI ID: 11.
	FxF1 float64 `json:"fxF1"`
	// The extraordinary wave F2 critical frequency, in MHz. URSI ID: 01.
	FxF2 float64 `json:"fxF2"`
	// The highest frequency of F-trace in MHz. Note: fxI is with capital i. URSI
	// ID: 51.
	FxI float64 `json:"fxI"`
	// Array of altitudes above station level for plasma frequency/density arrays in km
	// (must match the size of the plasmaFrequency and electronDensity Arrays).
	Height []float64 `json:"height"`
	// True height of the E peak in km. URSI ID: CE.
	HmE float64 `json:"hmE"`
	// True height of the F1 peak in km. URSI ID: BE.
	HmF1 float64 `json:"hmF1"`
	// True height of the F2 peak in km. URSI ID: AE.
	HmF2 float64 `json:"hmF2"`
	// The minimum virtual height of the normal E layer trace in km. URSI ID: 24.
	HprimeE float64 `json:"hprimeE"`
	// Minimum virtual height of night time auroral E layer trace in km. URSI ID: 27.
	HprimeEa float64 `json:"hprimeEa"`
	// The minimum height of the trace used to give foEs in km. URSI ID: 34.
	HprimeEs float64 `json:"hprimeEs"`
	// The minimum virtual height of the ordinary wave trace taken as a whole, in km.
	// URSI ID: 16.
	HprimeF float64 `json:"hprimeF"`
	// The minimum virtual height of reflection at a point where the trace is
	// horizontal in the F region in km. URSI ID: 14.
	HprimeF1 float64 `json:"hprimeF1"`
	// The minimum virtual height of ordinary wave trace for the highest stable
	// stratification in the F region in km. URSI ID: 4.
	HprimeF2 float64 `json:"hprimeF2"`
	// Virtual height at MUF/OblFactor frequency in MHz.
	HprimefMuf float64 `json:"hprimefMUF"`
	// Minimum virtual height of the trace used to determine foP in km. URSI ID: 56.
	HprimeP float64 `json:"hprimeP"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// Lowest usable frequency.
	Luf float64 `json:"luf"`
	// MUF(D)/foF2.
	Md float64 `json:"md"`
	// Maximum Usable Frequency for ground distance D in MHz.
	Mufd float64 `json:"mufd"`
	// Name of the algorithm used for the electron density profile.
	NeProfileName string `json:"neProfileName"`
	// Version of the algorithm used for the electron density profile.
	NeProfileVersion float64 `json:"neProfileVersion"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID string                            `json:"origSensorId"`
	Phase        IonoObservationTupleResponsePhase `json:"phase"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64 `json:"plasmaFrequencyUncertainty"`
	// Equipment location.
	PlatformName string                                   `json:"platformName"`
	Polarization IonoObservationTupleResponsePolarization `json:"polarization"`
	Power        IonoObservationTupleResponsePower        `json:"power"`
	// Average range spread of E layer in km. URSI ID: 85.
	Qe float64 `json:"qe"`
	// Average range spread of F layer in km. URSI ID: 84.
	Qf    float64                           `json:"qf"`
	Range IonoObservationTupleResponseRange `json:"range"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	ReceiveCoordinates [][]float64 `json:"receiveCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	ReceiveSensorType IonoObservationTupleResponseReceiveSensorType `json:"receiveSensorType"`
	// Array of restricted frequencies.
	RestrictedFrequency []float64 `json:"restrictedFrequency"`
	// Notes for the restrictedFrequency data.
	RestrictedFrequencyNotes string `json:"restrictedFrequencyNotes"`
	// Effective scale height at hmF2 Titheridge method in km. URSI ID: 69.
	ScaleHeightF2Peak float64 `json:"scaleHeightF2Peak"`
	// The ScalerInfo record describes the person or system who interpreted the
	// ionogram in IonoObservation.
	ScalerInfo IonoObservationTupleResponseScalerInfo `json:"scalerInfo"`
	Stokes     IonoObservationTupleResponseStokes     `json:"stokes"`
	// Details concerning the composition/intention/interpretation/audience/etc. of any
	// data recorded here. This field may contain all of the intended information e.g.
	// info on signal waveforms used, antenna setup, etc. OR may describe the
	// data/settings to be provided in the “data” field.
	SystemNotes string `json:"systemNotes"`
	// Total Ionospheric Electron Content \*10^16e/m^2. 1 TEC Unit (TECU) = 10^16
	// electrons/m^2. URSI ID: 72.
	Tec float64 `json:"tec"`
	// Array of degrees clockwise from true North of the TID.
	TidAzimuth []float64 `json:"tidAzimuth"`
	// Array of 1/frequency of the TID wave.
	TidPeriods []float64 `json:"tidPeriods"`
	// Array of speed in m/s at which the disturbance travels through the ionosphere.
	TidPhaseSpeeds []float64                                `json:"tidPhaseSpeeds"`
	Time           IonoObservationTupleResponseTime         `json:"time"`
	TraceGeneric   IonoObservationTupleResponseTraceGeneric `json:"traceGeneric"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	TransmitCoordinates [][]float64 `json:"transmitCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	TransmitSensorType IonoObservationTupleResponseTransmitSensorType `json:"transmitSensorType"`
	// Characterization of the shape of Es trace. URSI ID: 36.
	TypeEs string `json:"typeEs"`
	// Time the row was updated in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Parabolic E layer semi-thickness in km. URSI ID: 83.
	YE float64 `json:"yE"`
	// Parabolic F1 layer semi-thickness in km. URSI ID: 95.
	YF1 float64 `json:"yF1"`
	// Parabolic F2 layer semi-thickness in km. URSI ID: 94.
	YF2 float64 `json:"yF2"`
	// True height at half peak electron density in the F2 layer in km. URSI ID: 93.
	ZhalfNm float64 `json:"zhalfNm"`
	// Peak height of E-layer in km. URSI ID: 90.
	ZmE float64 `json:"zmE"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                  respjson.Field
		DataMode                               respjson.Field
		Source                                 respjson.Field
		StartTimeUtc                           respjson.Field
		StationID                              respjson.Field
		System                                 respjson.Field
		SystemInfo                             respjson.Field
		ID                                     respjson.Field
		Amplitude                              respjson.Field
		AntennaElementPosition                 respjson.Field
		AntennaElementPositionCoordinateSystem respjson.Field
		ArtistFlags                            respjson.Field
		Azimuth                                respjson.Field
		B0                                     respjson.Field
		B1                                     respjson.Field
		CharAtts                               respjson.Field
		CreatedAt                              respjson.Field
		CreatedBy                              respjson.Field
		D                                      respjson.Field
		D1                                     respjson.Field
		Datum                                  respjson.Field
		DeltafoF2                              respjson.Field
		DensityProfile                         respjson.Field
		Doppler                                respjson.Field
		DownE                                  respjson.Field
		DownEs                                 respjson.Field
		DownF                                  respjson.Field
		ElectronDensity                        respjson.Field
		ElectronDensityUncertainty             respjson.Field
		Elevation                              respjson.Field
		FbEs                                   respjson.Field
		Fe                                     respjson.Field
		Ff                                     respjson.Field
		FhprimeF                               respjson.Field
		FhprimeF2                              respjson.Field
		Fmin                                   respjson.Field
		FminE                                  respjson.Field
		FminEs                                 respjson.Field
		FminF                                  respjson.Field
		Fmuf                                   respjson.Field
		FoE                                    respjson.Field
		FoEa                                   respjson.Field
		FoEp                                   respjson.Field
		FoEs                                   respjson.Field
		FoF1                                   respjson.Field
		FoF1p                                  respjson.Field
		FoF2                                   respjson.Field
		FoF2p                                  respjson.Field
		FoP                                    respjson.Field
		Frequency                              respjson.Field
		FxE                                    respjson.Field
		FxF1                                   respjson.Field
		FxF2                                   respjson.Field
		FxI                                    respjson.Field
		Height                                 respjson.Field
		HmE                                    respjson.Field
		HmF1                                   respjson.Field
		HmF2                                   respjson.Field
		HprimeE                                respjson.Field
		HprimeEa                               respjson.Field
		HprimeEs                               respjson.Field
		HprimeF                                respjson.Field
		HprimeF1                               respjson.Field
		HprimeF2                               respjson.Field
		HprimefMuf                             respjson.Field
		HprimeP                                respjson.Field
		IDSensor                               respjson.Field
		Luf                                    respjson.Field
		Md                                     respjson.Field
		Mufd                                   respjson.Field
		NeProfileName                          respjson.Field
		NeProfileVersion                       respjson.Field
		Origin                                 respjson.Field
		OrigNetwork                            respjson.Field
		OrigSensorID                           respjson.Field
		Phase                                  respjson.Field
		PlasmaFrequency                        respjson.Field
		PlasmaFrequencyUncertainty             respjson.Field
		PlatformName                           respjson.Field
		Polarization                           respjson.Field
		Power                                  respjson.Field
		Qe                                     respjson.Field
		Qf                                     respjson.Field
		Range                                  respjson.Field
		ReceiveCoordinates                     respjson.Field
		ReceiveSensorType                      respjson.Field
		RestrictedFrequency                    respjson.Field
		RestrictedFrequencyNotes               respjson.Field
		ScaleHeightF2Peak                      respjson.Field
		ScalerInfo                             respjson.Field
		Stokes                                 respjson.Field
		SystemNotes                            respjson.Field
		Tec                                    respjson.Field
		TidAzimuth                             respjson.Field
		TidPeriods                             respjson.Field
		TidPhaseSpeeds                         respjson.Field
		Time                                   respjson.Field
		TraceGeneric                           respjson.Field
		TransmitCoordinates                    respjson.Field
		TransmitSensorType                     respjson.Field
		TypeEs                                 respjson.Field
		UpdatedAt                              respjson.Field
		UpdatedBy                              respjson.Field
		YE                                     respjson.Field
		YF1                                    respjson.Field
		YF2                                    respjson.Field
		ZhalfNm                                respjson.Field
		ZmE                                    respjson.Field
		ExtraFields                            map[string]respjson.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponse) UnmarshalJSON(data []byte) error {
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
type IonoObservationTupleResponseDataMode string

const (
	IonoObservationTupleResponseDataModeReal      IonoObservationTupleResponseDataMode = "REAL"
	IonoObservationTupleResponseDataModeTest      IonoObservationTupleResponseDataMode = "TEST"
	IonoObservationTupleResponseDataModeSimulated IonoObservationTupleResponseDataMode = "SIMULATED"
	IonoObservationTupleResponseDataModeExercise  IonoObservationTupleResponseDataMode = "EXERCISE"
)

type IonoObservationTupleResponseAmplitude struct {
	// Array of amplitude data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for amplitude dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the amplitude data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseAmplitude) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseAmplitude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponseAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseAntennaElementPosition) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseAntennaElementPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
// (lat, long, radii).
type IonoObservationTupleResponseAntennaElementPositionCoordinateSystem string

const (
	IonoObservationTupleResponseAntennaElementPositionCoordinateSystemJ2000                   IonoObservationTupleResponseAntennaElementPositionCoordinateSystem = "J2000"
	IonoObservationTupleResponseAntennaElementPositionCoordinateSystemEcrEcef                 IonoObservationTupleResponseAntennaElementPositionCoordinateSystem = "ECR/ECEF"
	IonoObservationTupleResponseAntennaElementPositionCoordinateSystemTeme                    IonoObservationTupleResponseAntennaElementPositionCoordinateSystem = "TEME"
	IonoObservationTupleResponseAntennaElementPositionCoordinateSystemGcrf                    IonoObservationTupleResponseAntennaElementPositionCoordinateSystem = "GCRF"
	IonoObservationTupleResponseAntennaElementPositionCoordinateSystemWgs84GeoDeticLatLongAlt IonoObservationTupleResponseAntennaElementPositionCoordinateSystem = "WGS84 (GEODetic lat, long, alt)"
	IonoObservationTupleResponseAntennaElementPositionCoordinateSystemGeoCentricLatLongRadii  IonoObservationTupleResponseAntennaElementPositionCoordinateSystem = "GEOCentric (lat, long, radii)"
)

type IonoObservationTupleResponseAzimuth struct {
	// Array of incoming azimuth at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the azimuth array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the azimuth data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseAzimuth) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseAzimuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Characteristic attributes of a IonoObservation.
type IonoObservationTupleResponseCharAtt struct {
	// Characteristic name. This value should reflect the UDL field name for the
	// corresponding characteristic.
	CharName string `json:"charName"`
	// Input parameters for the climate model.
	ClimateModelInputParams []string `json:"climateModelInputParams"`
	// Name of the climate model.
	ClimateModelName string `json:"climateModelName"`
	// List of options for the climate model.
	ClimateModelOptions []string `json:"climateModelOptions"`
	// Descriptive letter (D) for the characteristic specified by URSI ID. Describes
	// specific ionospheric conditions, beyond numerical values.
	D string `json:"d" format:"char"`
	// Specified characteristic's lower bound. Should be less than or equal to the
	// characteristic's current value as set in this record.
	LowerBound float64 `json:"lowerBound"`
	// Qualifying letter (Q) for the characteristic specified by URSI ID. Describes
	// specific ionospheric conditions, beyond numerical values.
	Q string `json:"q" format:"char"`
	// Uncertainty Bounds (lower and upper) define an interval around reported value
	// that contains true value at the specified probability level. Probability levels
	// are specified in terms of percentile (from 1 to 100) or the standard deviation,
	// sigma (e.g. 1sigma, 2sigma, 3sigma, 5percentile, 10percentile, 25percentile).
	UncertaintyBoundType string `json:"uncertaintyBoundType"`
	// Specified characteristic's upper bound. Should be greater than or equal to the
	// characteristic's current value as set in this record.
	UpperBound float64 `json:"upperBound"`
	// Characteristic's URSI ID. See the characteristic's description for its
	// corresponding URSI ID.
	UrsiID string `json:"ursiID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharName                respjson.Field
		ClimateModelInputParams respjson.Field
		ClimateModelName        respjson.Field
		ClimateModelOptions     respjson.Field
		D                       respjson.Field
		LowerBound              respjson.Field
		Q                       respjson.Field
		UncertaintyBoundType    respjson.Field
		UpperBound              respjson.Field
		UrsiID                  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseCharAtt) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseCharAtt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponseDatum struct {
	// Array to support sparse data collections.
	Data []float64 `json:"data"`
	// Notes for the datum with details of what the data is, units, etc.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Notes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDatum) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseDatum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonoObservationTupleResponseDensityProfile struct {
	// Full set of the IRI formalism coefficients.
	Iri IonoObservationTupleResponseDensityProfileIri `json:"iri"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonoObservationTupleResponseDensityProfileParabolic `json:"parabolic"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonoObservationTupleResponseDensityProfileQuasiParabolic `json:"quasiParabolic"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonoObservationTupleResponseDensityProfileShiftedChebyshev `json:"shiftedChebyshev"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonoObservationTupleResponseDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonoObservationTupleResponseDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs"`
	// Description of the valley model and parameters.
	ValleyModelDescription string `json:"valleyModelDescription"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iri                          respjson.Field
		Parabolic                    respjson.Field
		QuasiParabolic               respjson.Field
		ShiftedChebyshev             respjson.Field
		TopsideExtensionChapmanConst respjson.Field
		TopsideExtensionVaryChap     respjson.Field
		ValleyModelCoeffs            respjson.Field
		ValleyModelDescription       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfile) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseDensityProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full set of the IRI formalism coefficients.
type IonoObservationTupleResponseDensityProfileIri struct {
	// B0 parameter of the F2 layer shape.
	B0 float64 `json:"b0"`
	// B1 parameter of the F2 layer shape.
	B1 float64 `json:"b1"`
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi float64 `json:"chi"`
	// D1 parameter of the F1 layer shape.
	D1 float64 `json:"d1"`
	// Description of IRI implementation.
	Description string `json:"description"`
	// TBD.
	Fp1 float64 `json:"fp1"`
	// TBD.
	Fp2 float64 `json:"fp2"`
	// TBD.
	Fp30 float64 `json:"fp30"`
	// TBD.
	Fp3U float64 `json:"fp3U"`
	// Starting height of the D layer, in kilometers.
	Ha float64 `json:"ha"`
	// Height of the intermediate region at the top of D region, in kilometers.
	Hdx float64 `json:"hdx"`
	// Peak height of the D layer, in kilometers.
	HmD float64 `json:"hmD"`
	// Peak height of the F2 layer, in kilometers.
	HmE float64 `json:"hmE"`
	// Peak height of the F1 layer, in kilometers.
	HmF1 float64 `json:"hmF1"`
	// Peak height of F2 layer, in kilometers.
	HmF2 float64 `json:"hmF2"`
	// The valley height, in kilometers.
	HValTop float64 `json:"hValTop"`
	// Height HZ of the interim layer, in kilometers.
	Hz float64 `json:"hz"`
	// Peak density of the D layer, in per cubic centimeter.
	NmD float64 `json:"nmD"`
	// Peak density of the E layer, in per cubic centimeter.
	NmE float64 `json:"nmE"`
	// Peak density of the F1 layer, in grams per cubic centimeter.
	NmF1 float64 `json:"nmF1"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 float64 `json:"nmF2"`
	// The valley depth, in grams per cubic centimeter.
	NValB float64 `json:"nValB"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		B0          respjson.Field
		B1          respjson.Field
		Chi         respjson.Field
		D1          respjson.Field
		Description respjson.Field
		Fp1         respjson.Field
		Fp2         respjson.Field
		Fp30        respjson.Field
		Fp3U        respjson.Field
		Ha          respjson.Field
		Hdx         respjson.Field
		HmD         respjson.Field
		HmE         respjson.Field
		HmF1        respjson.Field
		HmF2        respjson.Field
		HValTop     respjson.Field
		Hz          respjson.Field
		NmD         respjson.Field
		NmE         respjson.Field
		NmF1        respjson.Field
		NmF2        respjson.Field
		NValB       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfileIri) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseDensityProfileIri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationTupleResponseDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description string `json:"description"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonoObservationTupleResponseDensityProfileParabolicParabolicItem `json:"parabolicItems"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		ParabolicItems respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfileParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseDensityProfileParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationTupleResponseDensityProfileParabolicParabolicItem struct {
	// Plasma frequency at the layer peak, in MHz.
	F float64 `json:"f"`
	// Ionospheric plasma layer (E, F1, or F2).
	Layer string `json:"layer"`
	// Half-thickness of the layer, in kilometers.
	Y float64 `json:"y"`
	// Height of the layer peak, in kilometers.
	Z float64 `json:"z"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		F           respjson.Field
		Layer       respjson.Field
		Y           respjson.Field
		Z           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfileParabolicParabolicItem) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationTupleResponseDensityProfileParabolicParabolicItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonoObservationTupleResponseDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description string `json:"description"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius float64 `json:"earthRadius"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonoObservationTupleResponseDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description            respjson.Field
		EarthRadius            respjson.Field
		QuasiParabolicSegments respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfileQuasiParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseDensityProfileQuasiParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonoObservationTupleResponseDensityProfileQuasiParabolicQuasiParabolicSegment struct {
	// Coefficient A.
	A float64 `json:"a"`
	// Coefficient B.
	B float64 `json:"b"`
	// Coefficient C.
	C float64 `json:"c"`
	// Best-fit error.
	Error float64 `json:"error"`
	// The index of this segment in the list, from 1 to NumSegments.
	Index int64 `json:"index"`
	// Starting range of the segment, in kilometers from the Earth's center.
	Rb float64 `json:"rb"`
	// Ending range of the segment, in kilometers from the Earth's center.
	Re float64 `json:"re"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		A           respjson.Field
		B           respjson.Field
		C           respjson.Field
		Error       respjson.Field
		Index       respjson.Field
		Rb          respjson.Field
		Re          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfileQuasiParabolicQuasiParabolicSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationTupleResponseDensityProfileQuasiParabolicQuasiParabolicSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonoObservationTupleResponseDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description string `json:"description"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonoObservationTupleResponseDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description       respjson.Field
		ShiftedChebyshevs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfileShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationTupleResponseDensityProfileShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonoObservationTupleResponseDensityProfileShiftedChebyshevShiftedChebyshev struct {
	// Array of coefficients.
	Coeffs []float64 `json:"coeffs"`
	// Best fit error.
	Error float64 `json:"error"`
	// Start frequency of the layer, in MHz.
	Fstart float64 `json:"fstart"`
	// Stop frequency of the layer, in MHz.
	Fstop float64 `json:"fstop"`
	// The ionospheric plasma layer.
	Layer string `json:"layer"`
	// Number of coefficients in the expansion.
	N int64 `json:"n"`
	// Peak height of the layer, in kilometers.
	Peakheight float64 `json:"peakheight"`
	// Height at which density is half of the peak Nm, in kilometers.
	ZhalfNm float64 `json:"zhalfNm"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coeffs      respjson.Field
		Error       respjson.Field
		Fstart      respjson.Field
		Fstop       respjson.Field
		Layer       respjson.Field
		N           respjson.Field
		Peakheight  respjson.Field
		ZhalfNm     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfileShiftedChebyshevShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationTupleResponseDensityProfileShiftedChebyshevShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters of the constant-scale-height Chapman layer.
type IonoObservationTupleResponseDensityProfileTopsideExtensionChapmanConst struct {
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi float64 `json:"chi"`
	// Description of the Chapman computation technique.
	Description string `json:"description"`
	// Peak height of F2 layer, in kilometers.
	HmF2 float64 `json:"hmF2"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 float64 `json:"nmF2"`
	// Scale height if F2 layer at the peak, in kilometers.
	ScaleF2 float64 `json:"scaleF2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chi         respjson.Field
		Description respjson.Field
		HmF2        respjson.Field
		NmF2        respjson.Field
		ScaleF2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfileTopsideExtensionChapmanConst) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationTupleResponseDensityProfileTopsideExtensionChapmanConst) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Varying scale height Chapman topside layer.
type IonoObservationTupleResponseDensityProfileTopsideExtensionVaryChap struct {
	// Alpha parameter of the profile shape.
	Alpha float64 `json:"alpha"`
	// Beta parameter of the profile shape.
	Beta float64 `json:"beta"`
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi float64 `json:"chi"`
	// Description of the Chapman computation technique.
	Description string `json:"description"`
	// Peak height of F2 layer, in kilometers.
	HmF2 float64 `json:"hmF2"`
	// Transition height, in kilometers.
	Ht float64 `json:"ht"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 float64 `json:"nmF2"`
	// Scale height if F2 layer at the peak, in kilometers.
	ScaleF2 float64 `json:"scaleF2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alpha       respjson.Field
		Beta        respjson.Field
		Chi         respjson.Field
		Description respjson.Field
		HmF2        respjson.Field
		Ht          respjson.Field
		NmF2        respjson.Field
		ScaleF2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDensityProfileTopsideExtensionVaryChap) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationTupleResponseDensityProfileTopsideExtensionVaryChap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponseDoppler struct {
	// Array of received doppler shifts in Hz.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the doppler array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the doppler data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseDoppler) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseDoppler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponseElevation struct {
	// Array of incoming elevation at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the elevation array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the elevation data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseElevation) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseElevation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponseFrequency struct {
	// Array of frequency data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for frequency dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the frequency array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the frequency data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseFrequency) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseFrequency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponsePhase struct {
	// Array of phase data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for phase dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the phase array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the phase data. Orientation and position for each antenna element
	// across the antenna_element dimension.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponsePhase) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponsePhase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponsePolarization struct {
	// Array of polarization data.
	Data [][][][][][][]string `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for polarization dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the polarization data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponsePolarization) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponsePolarization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponsePower struct {
	// Array of received power in db.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the power array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the power data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponsePower) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponsePower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponseRange struct {
	// Array of range data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for range dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the range array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the range data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseRange) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonoObservationTupleResponseReceiveSensorType string

const (
	IonoObservationTupleResponseReceiveSensorTypeMobile IonoObservationTupleResponseReceiveSensorType = "Mobile"
	IonoObservationTupleResponseReceiveSensorTypeStatic IonoObservationTupleResponseReceiveSensorType = "Static"
)

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonoObservationTupleResponseScalerInfo struct {
	// Scaler confidence level.
	ConfidenceLevel int64 `json:"confidenceLevel"`
	// Scaler confidence score.
	ConfidenceScore int64 `json:"confidenceScore"`
	// Scaler name.
	Name string `json:"name"`
	// Scaler organization.
	Organization string `json:"organization"`
	// Scaler type (MANUAL, AUTOMATIC or UNKNOWN).
	Type string `json:"type"`
	// Scaler version.
	Version float64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfidenceLevel respjson.Field
		ConfidenceScore respjson.Field
		Name            respjson.Field
		Organization    respjson.Field
		Type            respjson.Field
		Version         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseScalerInfo) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseScalerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponseStokes struct {
	// Array of received stokes data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the stoke array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the stokes data.
	Notes string `json:"notes"`
	// S1, S2, and S3 (the normalized Stokes parameters 1, 2, and 3).
	S []float64 `json:"s"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		S             respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseStokes) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseStokes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponseTime struct {
	// Array of times in number of seconds passed since January 1st, 1970 with the same
	// dimensions as power.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the time array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// The notes indicate the scheme and accuracy.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Dimensions    respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseTime) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleResponseTraceGeneric struct {
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName"`
	// Notes for the trace generic data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		DimensionName respjson.Field
		Notes         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationTupleResponseTraceGeneric) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationTupleResponseTraceGeneric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonoObservationTupleResponseTransmitSensorType string

const (
	IonoObservationTupleResponseTransmitSensorTypeMobile IonoObservationTupleResponseTransmitSensorType = "Mobile"
	IonoObservationTupleResponseTransmitSensorTypeStatic IonoObservationTupleResponseTransmitSensorType = "Static"
)

type IonoObservationListParams struct {
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time        `query:"startTimeUTC,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IonoObservationListParams]'s query parameters as
// `url.Values`.
func (r IonoObservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonoObservationCountParams struct {
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time        `query:"startTimeUTC,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IonoObservationCountParams]'s query parameters as
// `url.Values`.
func (r IonoObservationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonoObservationNewBulkParams struct {
	Body []IonoObservationNewBulkParamsBody
	paramObj
}

func (r IonoObservationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *IonoObservationNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide operations for posting and querying ionospheric
// observation data. Characteristics are defined by the CHARS: URSI IIWG format for
// archiving monthly ionospheric characteristics, INAG Bulletin No. 62
// specification. Qualifying and Descriptive letters are defined by the URSI
// Handbook for Ionogram Interpretation and reduction, Report UAG, No. 23A
// specification.
//
// The properties ClassificationMarking, DataMode, Source, StartTimeUtc, StationID,
// System, SystemInfo are required.
type IonoObservationNewBulkParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Sounding Start time in ISO8601 UTC format.
	StartTimeUtc time.Time `json:"startTimeUTC,required" format:"date-time"`
	// URSI code for station or stations producing the ionosonde.
	StationID string `json:"stationId,required"`
	// Ionosonde hardware type or data collection type together with possible
	// additional descriptors.
	System string `json:"system,required"`
	// Names of settings.
	SystemInfo string `json:"systemInfo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// IRI thickness parameter in km. URSI ID: D0.
	B0 param.Opt[float64] `json:"b0,omitzero"`
	// IRI profile shape parameter. URSI ID: D1.
	B1 param.Opt[float64] `json:"b1,omitzero"`
	// Distance for MUF calculation in km.
	D param.Opt[float64] `json:"d,omitzero"`
	// IRI profile shape parameter, F1 layer. URSI ID: D2.
	D1 param.Opt[float64] `json:"d1,omitzero"`
	// Adjustment to the scaled foF2 during profile inversion in MHz.
	DeltafoF2 param.Opt[float64] `json:"deltafoF2,omitzero"`
	// Lowering of E trace to the leading edge in km.
	DownE param.Opt[float64] `json:"downE,omitzero"`
	// Lowering of Es trace to the leading edge in km.
	DownEs param.Opt[float64] `json:"downEs,omitzero"`
	// Lowering of F trace to the leading edge in km.
	DownF param.Opt[float64] `json:"downF,omitzero"`
	// The blanketing frequency of layer used to derive foEs in MHz. URSI ID: 32.
	FbEs param.Opt[float64] `json:"fbEs,omitzero"`
	// Frequency spread beyond foE in MHz. URSI ID: 87.
	Fe param.Opt[float64] `json:"fe,omitzero"`
	// Frequency spread between fxF2 and FxI in MHz. URSI ID: 86.
	Ff param.Opt[float64] `json:"ff,omitzero"`
	// The frequency at which hprimeF is measured in MHz. URSI ID: 61.
	FhprimeF param.Opt[float64] `json:"fhprimeF,omitzero"`
	// The frequency at which hprimeF2 is measured in MHz. URSI ID: 60.
	FhprimeF2 param.Opt[float64] `json:"fhprimeF2,omitzero"`
	// Lowest frequency at which echo traces are observed on the ionogram, in MHz. URSI
	// ID: 42.
	Fmin param.Opt[float64] `json:"fmin,omitzero"`
	// Minimum frequency of E layer echoes in MHz. URSI ID: 81.
	FminE param.Opt[float64] `json:"fminE,omitzero"`
	// Minimum frequency of Es layer in MHz.
	FminEs param.Opt[float64] `json:"fminEs,omitzero"`
	// Minimum frequency of F layer echoes in MHz. URSI ID: 80.
	FminF param.Opt[float64] `json:"fminF,omitzero"`
	// MUF/OblFactor in MHz.
	Fmuf param.Opt[float64] `json:"fmuf,omitzero"`
	// The ordinary wave critical frequency of the lowest thick layer which causes a
	// discontinuity, in MHz. URSI ID: 20.
	FoE param.Opt[float64] `json:"foE,omitzero"`
	// Critical frequency of night time auroral E layer in MHz. URSI ID: 23.
	FoEa param.Opt[float64] `json:"foEa,omitzero"`
	// Predicted value of foE in MHz.
	FoEp param.Opt[float64] `json:"foEp,omitzero"`
	// Highest ordinary wave frequency at which a mainly continuous Es trace is
	// observed, in MHz. URSI ID: 30.
	FoEs param.Opt[float64] `json:"foEs,omitzero"`
	// The ordinary wave F1 critical frequency, in MHz. URSI ID: 10.
	FoF1 param.Opt[float64] `json:"foF1,omitzero"`
	// Predicted value of foF1 in MHz.
	FoF1p param.Opt[float64] `json:"foF1p,omitzero"`
	// The ordinary wave critical frequency of the highest stratification in the F
	// region, specified in MHz. URSI ID: 00.
	FoF2 param.Opt[float64] `json:"foF2,omitzero"`
	// Predicted value of foF2 in MHz.
	FoF2p param.Opt[float64] `json:"foF2p,omitzero"`
	// Highest ordinary wave critical frequency of F region patch trace in MHz. URSI
	// ID: 55.
	FoP param.Opt[float64] `json:"foP,omitzero"`
	// The extraordinary wave E critical frequency, in MHz. URSI ID: 21.
	FxE param.Opt[float64] `json:"fxE,omitzero"`
	// The extraordinary wave F1 critical frequency, in MHz. URSI ID: 11.
	FxF1 param.Opt[float64] `json:"fxF1,omitzero"`
	// The extraordinary wave F2 critical frequency, in MHz. URSI ID: 01.
	FxF2 param.Opt[float64] `json:"fxF2,omitzero"`
	// The highest frequency of F-trace in MHz. Note: fxI is with capital i. URSI
	// ID: 51.
	FxI param.Opt[float64] `json:"fxI,omitzero"`
	// True height of the E peak in km. URSI ID: CE.
	HmE param.Opt[float64] `json:"hmE,omitzero"`
	// True height of the F1 peak in km. URSI ID: BE.
	HmF1 param.Opt[float64] `json:"hmF1,omitzero"`
	// True height of the F2 peak in km. URSI ID: AE.
	HmF2 param.Opt[float64] `json:"hmF2,omitzero"`
	// The minimum virtual height of the normal E layer trace in km. URSI ID: 24.
	HprimeE param.Opt[float64] `json:"hprimeE,omitzero"`
	// Minimum virtual height of night time auroral E layer trace in km. URSI ID: 27.
	HprimeEa param.Opt[float64] `json:"hprimeEa,omitzero"`
	// The minimum height of the trace used to give foEs in km. URSI ID: 34.
	HprimeEs param.Opt[float64] `json:"hprimeEs,omitzero"`
	// The minimum virtual height of the ordinary wave trace taken as a whole, in km.
	// URSI ID: 16.
	HprimeF param.Opt[float64] `json:"hprimeF,omitzero"`
	// The minimum virtual height of reflection at a point where the trace is
	// horizontal in the F region in km. URSI ID: 14.
	HprimeF1 param.Opt[float64] `json:"hprimeF1,omitzero"`
	// The minimum virtual height of ordinary wave trace for the highest stable
	// stratification in the F region in km. URSI ID: 4.
	HprimeF2 param.Opt[float64] `json:"hprimeF2,omitzero"`
	// Virtual height at MUF/OblFactor frequency in MHz.
	HprimefMuf param.Opt[float64] `json:"hprimefMUF,omitzero"`
	// Minimum virtual height of the trace used to determine foP in km. URSI ID: 56.
	HprimeP param.Opt[float64] `json:"hprimeP,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Lowest usable frequency.
	Luf param.Opt[float64] `json:"luf,omitzero"`
	// MUF(D)/foF2.
	Md param.Opt[float64] `json:"md,omitzero"`
	// Maximum Usable Frequency for ground distance D in MHz.
	Mufd param.Opt[float64] `json:"mufd,omitzero"`
	// Name of the algorithm used for the electron density profile.
	NeProfileName param.Opt[string] `json:"neProfileName,omitzero"`
	// Version of the algorithm used for the electron density profile.
	NeProfileVersion param.Opt[float64] `json:"neProfileVersion,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Equipment location.
	PlatformName param.Opt[string] `json:"platformName,omitzero"`
	// Average range spread of E layer in km. URSI ID: 85.
	Qe param.Opt[float64] `json:"qe,omitzero"`
	// Average range spread of F layer in km. URSI ID: 84.
	Qf param.Opt[float64] `json:"qf,omitzero"`
	// Notes for the restrictedFrequency data.
	RestrictedFrequencyNotes param.Opt[string] `json:"restrictedFrequencyNotes,omitzero"`
	// Effective scale height at hmF2 Titheridge method in km. URSI ID: 69.
	ScaleHeightF2Peak param.Opt[float64] `json:"scaleHeightF2Peak,omitzero"`
	// Details concerning the composition/intention/interpretation/audience/etc. of any
	// data recorded here. This field may contain all of the intended information e.g.
	// info on signal waveforms used, antenna setup, etc. OR may describe the
	// data/settings to be provided in the “data” field.
	SystemNotes param.Opt[string] `json:"systemNotes,omitzero"`
	// Total Ionospheric Electron Content \*10^16e/m^2. 1 TEC Unit (TECU) = 10^16
	// electrons/m^2. URSI ID: 72.
	Tec param.Opt[float64] `json:"tec,omitzero"`
	// Characterization of the shape of Es trace. URSI ID: 36.
	TypeEs param.Opt[string] `json:"typeEs,omitzero"`
	// Parabolic E layer semi-thickness in km. URSI ID: 83.
	YE param.Opt[float64] `json:"yE,omitzero"`
	// Parabolic F1 layer semi-thickness in km. URSI ID: 95.
	YF1 param.Opt[float64] `json:"yF1,omitzero"`
	// Parabolic F2 layer semi-thickness in km. URSI ID: 94.
	YF2 param.Opt[float64] `json:"yF2,omitzero"`
	// True height at half peak electron density in the F2 layer in km. URSI ID: 93.
	ZhalfNm param.Opt[float64] `json:"zhalfNm,omitzero"`
	// Peak height of E-layer in km. URSI ID: 90.
	ZmE                    param.Opt[float64]                                     `json:"zmE,omitzero"`
	Amplitude              IonoObservationNewBulkParamsBodyAmplitude              `json:"amplitude,omitzero"`
	AntennaElementPosition IonoObservationNewBulkParamsBodyAntennaElementPosition `json:"antennaElementPosition,omitzero"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem string `json:"antennaElementPositionCoordinateSystem,omitzero"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                                 `json:"artistFlags,omitzero"`
	Azimuth     IonoObservationNewBulkParamsBodyAzimuth `json:"azimuth,omitzero"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonoObservationNewBulkParamsBodyCharAtt `json:"charAtts,omitzero"`
	Datum    IonoObservationNewBulkParamsBodyDatum     `json:"datum,omitzero"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonoObservationNewBulkParamsBodyDensityProfile `json:"densityProfile,omitzero"`
	Doppler        IonoObservationNewBulkParamsBodyDoppler        `json:"doppler,omitzero"`
	// Array of electron densities in cm^-3 (must match the size of the height and
	// plasmaFrequency arrays).
	ElectronDensity []float64 `json:"electronDensity,omitzero"`
	// Uncertainty in specifying the electron density at each height point of the
	// profile (must match the size of the electronDensity array).
	ElectronDensityUncertainty []float64                                 `json:"electronDensityUncertainty,omitzero"`
	Elevation                  IonoObservationNewBulkParamsBodyElevation `json:"elevation,omitzero"`
	Frequency                  IonoObservationNewBulkParamsBodyFrequency `json:"frequency,omitzero"`
	// Array of altitudes above station level for plasma frequency/density arrays in km
	// (must match the size of the plasmaFrequency and electronDensity Arrays).
	Height []float64                             `json:"height,omitzero"`
	Phase  IonoObservationNewBulkParamsBodyPhase `json:"phase,omitzero"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency,omitzero"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64                                    `json:"plasmaFrequencyUncertainty,omitzero"`
	Polarization               IonoObservationNewBulkParamsBodyPolarization `json:"polarization,omitzero"`
	Power                      IonoObservationNewBulkParamsBodyPower        `json:"power,omitzero"`
	Range                      IonoObservationNewBulkParamsBodyRange        `json:"range,omitzero"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	ReceiveCoordinates [][]float64 `json:"receiveCoordinates,omitzero"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	ReceiveSensorType string `json:"receiveSensorType,omitzero"`
	// Array of restricted frequencies.
	RestrictedFrequency []float64 `json:"restrictedFrequency,omitzero"`
	// The ScalerInfo record describes the person or system who interpreted the
	// ionogram in IonoObservation.
	ScalerInfo IonoObservationNewBulkParamsBodyScalerInfo `json:"scalerInfo,omitzero"`
	Stokes     IonoObservationNewBulkParamsBodyStokes     `json:"stokes,omitzero"`
	// Array of degrees clockwise from true North of the TID.
	TidAzimuth []float64 `json:"tidAzimuth,omitzero"`
	// Array of 1/frequency of the TID wave.
	TidPeriods []float64 `json:"tidPeriods,omitzero"`
	// Array of speed in m/s at which the disturbance travels through the ionosphere.
	TidPhaseSpeeds []float64                                    `json:"tidPhaseSpeeds,omitzero"`
	Time           IonoObservationNewBulkParamsBodyTime         `json:"time,omitzero"`
	TraceGeneric   IonoObservationNewBulkParamsBodyTraceGeneric `json:"traceGeneric,omitzero"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	TransmitCoordinates [][]float64 `json:"transmitCoordinates,omitzero"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	TransmitSensorType string `json:"transmitSensorType,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IonoObservationNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[IonoObservationNewBulkParamsBody](
		"antennaElementPositionCoordinateSystem", "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)", "GEOCentric (lat, long, radii)",
	)
	apijson.RegisterFieldValidator[IonoObservationNewBulkParamsBody](
		"receiveSensorType", "Mobile", "Static",
	)
	apijson.RegisterFieldValidator[IonoObservationNewBulkParamsBody](
		"transmitSensorType", "Mobile", "Static",
	)
}

type IonoObservationNewBulkParamsBodyAmplitude struct {
	// Notes for the amplitude data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of amplitude data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers for amplitude dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyAmplitude) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyAmplitude
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyAmplitude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyAntennaElementPosition) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyAntennaElementPosition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyAntennaElementPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyAzimuth struct {
	// Notes for the azimuth data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of incoming azimuth at the receiver.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the azimuth array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyAzimuth) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyAzimuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyAzimuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Characteristic attributes of a IonoObservation.
type IonoObservationNewBulkParamsBodyCharAtt struct {
	// Characteristic name. This value should reflect the UDL field name for the
	// corresponding characteristic.
	CharName param.Opt[string] `json:"charName,omitzero"`
	// Name of the climate model.
	ClimateModelName param.Opt[string] `json:"climateModelName,omitzero"`
	// Descriptive letter (D) for the characteristic specified by URSI ID. Describes
	// specific ionospheric conditions, beyond numerical values.
	D param.Opt[string] `json:"d,omitzero" format:"char"`
	// Specified characteristic's lower bound. Should be less than or equal to the
	// characteristic's current value as set in this record.
	LowerBound param.Opt[float64] `json:"lowerBound,omitzero"`
	// Qualifying letter (Q) for the characteristic specified by URSI ID. Describes
	// specific ionospheric conditions, beyond numerical values.
	Q param.Opt[string] `json:"q,omitzero" format:"char"`
	// Uncertainty Bounds (lower and upper) define an interval around reported value
	// that contains true value at the specified probability level. Probability levels
	// are specified in terms of percentile (from 1 to 100) or the standard deviation,
	// sigma (e.g. 1sigma, 2sigma, 3sigma, 5percentile, 10percentile, 25percentile).
	UncertaintyBoundType param.Opt[string] `json:"uncertaintyBoundType,omitzero"`
	// Specified characteristic's upper bound. Should be greater than or equal to the
	// characteristic's current value as set in this record.
	UpperBound param.Opt[float64] `json:"upperBound,omitzero"`
	// Characteristic's URSI ID. See the characteristic's description for its
	// corresponding URSI ID.
	UrsiID param.Opt[string] `json:"ursiID,omitzero"`
	// Input parameters for the climate model.
	ClimateModelInputParams []string `json:"climateModelInputParams,omitzero"`
	// List of options for the climate model.
	ClimateModelOptions []string `json:"climateModelOptions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyCharAtt) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyCharAtt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyCharAtt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyDatum struct {
	// Notes for the datum with details of what the data is, units, etc.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array to support sparse data collections.
	Data []float64 `json:"data,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDatum) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDatum
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDatum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonoObservationNewBulkParamsBodyDensityProfile struct {
	// Description of the valley model and parameters.
	ValleyModelDescription param.Opt[string] `json:"valleyModelDescription,omitzero"`
	// Full set of the IRI formalism coefficients.
	Iri IonoObservationNewBulkParamsBodyDensityProfileIri `json:"iri,omitzero"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonoObservationNewBulkParamsBodyDensityProfileParabolic `json:"parabolic,omitzero"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolic `json:"quasiParabolic,omitzero"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshev `json:"shiftedChebyshev,omitzero"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst,omitzero"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap,omitzero"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfile) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full set of the IRI formalism coefficients.
type IonoObservationNewBulkParamsBodyDensityProfileIri struct {
	// B0 parameter of the F2 layer shape.
	B0 param.Opt[float64] `json:"b0,omitzero"`
	// B1 parameter of the F2 layer shape.
	B1 param.Opt[float64] `json:"b1,omitzero"`
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi param.Opt[float64] `json:"chi,omitzero"`
	// D1 parameter of the F1 layer shape.
	D1 param.Opt[float64] `json:"d1,omitzero"`
	// Description of IRI implementation.
	Description param.Opt[string] `json:"description,omitzero"`
	// TBD.
	Fp1 param.Opt[float64] `json:"fp1,omitzero"`
	// TBD.
	Fp2 param.Opt[float64] `json:"fp2,omitzero"`
	// TBD.
	Fp30 param.Opt[float64] `json:"fp30,omitzero"`
	// TBD.
	Fp3U param.Opt[float64] `json:"fp3U,omitzero"`
	// Starting height of the D layer, in kilometers.
	Ha param.Opt[float64] `json:"ha,omitzero"`
	// Height of the intermediate region at the top of D region, in kilometers.
	Hdx param.Opt[float64] `json:"hdx,omitzero"`
	// Peak height of the D layer, in kilometers.
	HmD param.Opt[float64] `json:"hmD,omitzero"`
	// Peak height of the F2 layer, in kilometers.
	HmE param.Opt[float64] `json:"hmE,omitzero"`
	// Peak height of the F1 layer, in kilometers.
	HmF1 param.Opt[float64] `json:"hmF1,omitzero"`
	// Peak height of F2 layer, in kilometers.
	HmF2 param.Opt[float64] `json:"hmF2,omitzero"`
	// The valley height, in kilometers.
	HValTop param.Opt[float64] `json:"hValTop,omitzero"`
	// Height HZ of the interim layer, in kilometers.
	Hz param.Opt[float64] `json:"hz,omitzero"`
	// Peak density of the D layer, in per cubic centimeter.
	NmD param.Opt[float64] `json:"nmD,omitzero"`
	// Peak density of the E layer, in per cubic centimeter.
	NmE param.Opt[float64] `json:"nmE,omitzero"`
	// Peak density of the F1 layer, in grams per cubic centimeter.
	NmF1 param.Opt[float64] `json:"nmF1,omitzero"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 param.Opt[float64] `json:"nmF2,omitzero"`
	// The valley depth, in grams per cubic centimeter.
	NValB param.Opt[float64] `json:"nValB,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfileIri) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfileIri
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfileIri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationNewBulkParamsBodyDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonoObservationNewBulkParamsBodyDensityProfileParabolicParabolicItem `json:"parabolicItems,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfileParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfileParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfileParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationNewBulkParamsBodyDensityProfileParabolicParabolicItem struct {
	// Plasma frequency at the layer peak, in MHz.
	F param.Opt[float64] `json:"f,omitzero"`
	// Ionospheric plasma layer (E, F1, or F2).
	Layer param.Opt[string] `json:"layer,omitzero"`
	// Half-thickness of the layer, in kilometers.
	Y param.Opt[float64] `json:"y,omitzero"`
	// Height of the layer peak, in kilometers.
	Z param.Opt[float64] `json:"z,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfileParabolicParabolicItem) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfileParabolicParabolicItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfileParabolicParabolicItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius param.Opt[float64] `json:"earthRadius,omitzero"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment struct {
	// Coefficient A.
	A param.Opt[float64] `json:"a,omitzero"`
	// Coefficient B.
	B param.Opt[float64] `json:"b,omitzero"`
	// Coefficient C.
	C param.Opt[float64] `json:"c,omitzero"`
	// Best-fit error.
	Error param.Opt[float64] `json:"error,omitzero"`
	// The index of this segment in the list, from 1 to NumSegments.
	Index param.Opt[int64] `json:"index,omitzero"`
	// Starting range of the segment, in kilometers from the Earth's center.
	Rb param.Opt[float64] `json:"rb,omitzero"`
	// Ending range of the segment, in kilometers from the Earth's center.
	Re param.Opt[float64] `json:"re,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev struct {
	// Best fit error.
	Error param.Opt[float64] `json:"error,omitzero"`
	// Start frequency of the layer, in MHz.
	Fstart param.Opt[float64] `json:"fstart,omitzero"`
	// Stop frequency of the layer, in MHz.
	Fstop param.Opt[float64] `json:"fstop,omitzero"`
	// The ionospheric plasma layer.
	Layer param.Opt[string] `json:"layer,omitzero"`
	// Number of coefficients in the expansion.
	N param.Opt[int64] `json:"n,omitzero"`
	// Peak height of the layer, in kilometers.
	Peakheight param.Opt[float64] `json:"peakheight,omitzero"`
	// Height at which density is half of the peak Nm, in kilometers.
	ZhalfNm param.Opt[float64] `json:"zhalfNm,omitzero"`
	// Array of coefficients.
	Coeffs []float64 `json:"coeffs,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters of the constant-scale-height Chapman layer.
type IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst struct {
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi param.Opt[float64] `json:"chi,omitzero"`
	// Description of the Chapman computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Peak height of F2 layer, in kilometers.
	HmF2 param.Opt[float64] `json:"hmF2,omitzero"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 param.Opt[float64] `json:"nmF2,omitzero"`
	// Scale height if F2 layer at the peak, in kilometers.
	ScaleF2 param.Opt[float64] `json:"scaleF2,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Varying scale height Chapman topside layer.
type IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap struct {
	// Alpha parameter of the profile shape.
	Alpha param.Opt[float64] `json:"alpha,omitzero"`
	// Beta parameter of the profile shape.
	Beta param.Opt[float64] `json:"beta,omitzero"`
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi param.Opt[float64] `json:"chi,omitzero"`
	// Description of the Chapman computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Peak height of F2 layer, in kilometers.
	HmF2 param.Opt[float64] `json:"hmF2,omitzero"`
	// Transition height, in kilometers.
	Ht param.Opt[float64] `json:"ht,omitzero"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 param.Opt[float64] `json:"nmF2,omitzero"`
	// Scale height if F2 layer at the peak, in kilometers.
	ScaleF2 param.Opt[float64] `json:"scaleF2,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyDoppler struct {
	// Notes for the doppler data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of received doppler shifts in Hz.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the doppler array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyDoppler) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyDoppler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyDoppler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyElevation struct {
	// Notes for the elevation data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of incoming elevation at the receiver.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the elevation array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyElevation) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyElevation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyElevation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyFrequency struct {
	// Notes for the frequency data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of frequency data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for frequency dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the frequency array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyFrequency) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyFrequency
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyFrequency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyPhase struct {
	// Notes for the phase data. Orientation and position for each antenna element
	// across the antenna_element dimension.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of phase data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for phase dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the phase array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyPhase) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyPhase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyPhase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyPolarization struct {
	// Notes for the polarization data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of polarization data.
	Data [][][][][][][]string `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers for polarization dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyPolarization) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyPolarization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyPolarization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyPower struct {
	// Notes for the power data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of received power in db.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the power array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyPower) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyPower
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyPower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyRange struct {
	// Notes for the range data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of range data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for range dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the range array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyRange) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonoObservationNewBulkParamsBodyScalerInfo struct {
	// Scaler confidence level.
	ConfidenceLevel param.Opt[int64] `json:"confidenceLevel,omitzero"`
	// Scaler confidence score.
	ConfidenceScore param.Opt[int64] `json:"confidenceScore,omitzero"`
	// Scaler name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Scaler organization.
	Organization param.Opt[string] `json:"organization,omitzero"`
	// Scaler type (MANUAL, AUTOMATIC or UNKNOWN).
	Type param.Opt[string] `json:"type,omitzero"`
	// Scaler version.
	Version param.Opt[float64] `json:"version,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyScalerInfo) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyScalerInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyScalerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyStokes struct {
	// Notes for the stokes data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of received stokes data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the stoke array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	// S1, S2, and S3 (the normalized Stokes parameters 1, 2, and 3).
	S []float64 `json:"s,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyStokes) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyStokes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyStokes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyTime struct {
	// The notes indicate the scheme and accuracy.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of times in number of seconds passed since January 1st, 1970 with the same
	// dimensions as power.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the time array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyTime) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyTime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationNewBulkParamsBodyTraceGeneric struct {
	// Notes for the trace generic data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data,omitzero"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName,omitzero"`
	paramObj
}

func (r IonoObservationNewBulkParamsBodyTraceGeneric) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationNewBulkParamsBodyTraceGeneric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationNewBulkParamsBodyTraceGeneric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time        `query:"startTimeUTC,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IonoObservationTupleParams]'s query parameters as
// `url.Values`.
func (r IonoObservationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonoObservationUnvalidatedPublishParams struct {
	Body []IonoObservationUnvalidatedPublishParamsBody
	paramObj
}

func (r IonoObservationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *IonoObservationUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// These services provide operations for posting and querying ionospheric
// observation data. Characteristics are defined by the CHARS: URSI IIWG format for
// archiving monthly ionospheric characteristics, INAG Bulletin No. 62
// specification. Qualifying and Descriptive letters are defined by the URSI
// Handbook for Ionogram Interpretation and reduction, Report UAG, No. 23A
// specification.
//
// The properties ClassificationMarking, DataMode, Source, StartTimeUtc, StationID,
// System, SystemInfo are required.
type IonoObservationUnvalidatedPublishParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Sounding Start time in ISO8601 UTC format.
	StartTimeUtc time.Time `json:"startTimeUTC,required" format:"date-time"`
	// URSI code for station or stations producing the ionosonde.
	StationID string `json:"stationId,required"`
	// Ionosonde hardware type or data collection type together with possible
	// additional descriptors.
	System string `json:"system,required"`
	// Names of settings.
	SystemInfo string `json:"systemInfo,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// IRI thickness parameter in km. URSI ID: D0.
	B0 param.Opt[float64] `json:"b0,omitzero"`
	// IRI profile shape parameter. URSI ID: D1.
	B1 param.Opt[float64] `json:"b1,omitzero"`
	// Distance for MUF calculation in km.
	D param.Opt[float64] `json:"d,omitzero"`
	// IRI profile shape parameter, F1 layer. URSI ID: D2.
	D1 param.Opt[float64] `json:"d1,omitzero"`
	// Adjustment to the scaled foF2 during profile inversion in MHz.
	DeltafoF2 param.Opt[float64] `json:"deltafoF2,omitzero"`
	// Lowering of E trace to the leading edge in km.
	DownE param.Opt[float64] `json:"downE,omitzero"`
	// Lowering of Es trace to the leading edge in km.
	DownEs param.Opt[float64] `json:"downEs,omitzero"`
	// Lowering of F trace to the leading edge in km.
	DownF param.Opt[float64] `json:"downF,omitzero"`
	// The blanketing frequency of layer used to derive foEs in MHz. URSI ID: 32.
	FbEs param.Opt[float64] `json:"fbEs,omitzero"`
	// Frequency spread beyond foE in MHz. URSI ID: 87.
	Fe param.Opt[float64] `json:"fe,omitzero"`
	// Frequency spread between fxF2 and FxI in MHz. URSI ID: 86.
	Ff param.Opt[float64] `json:"ff,omitzero"`
	// The frequency at which hprimeF is measured in MHz. URSI ID: 61.
	FhprimeF param.Opt[float64] `json:"fhprimeF,omitzero"`
	// The frequency at which hprimeF2 is measured in MHz. URSI ID: 60.
	FhprimeF2 param.Opt[float64] `json:"fhprimeF2,omitzero"`
	// Lowest frequency at which echo traces are observed on the ionogram, in MHz. URSI
	// ID: 42.
	Fmin param.Opt[float64] `json:"fmin,omitzero"`
	// Minimum frequency of E layer echoes in MHz. URSI ID: 81.
	FminE param.Opt[float64] `json:"fminE,omitzero"`
	// Minimum frequency of Es layer in MHz.
	FminEs param.Opt[float64] `json:"fminEs,omitzero"`
	// Minimum frequency of F layer echoes in MHz. URSI ID: 80.
	FminF param.Opt[float64] `json:"fminF,omitzero"`
	// MUF/OblFactor in MHz.
	Fmuf param.Opt[float64] `json:"fmuf,omitzero"`
	// The ordinary wave critical frequency of the lowest thick layer which causes a
	// discontinuity, in MHz. URSI ID: 20.
	FoE param.Opt[float64] `json:"foE,omitzero"`
	// Critical frequency of night time auroral E layer in MHz. URSI ID: 23.
	FoEa param.Opt[float64] `json:"foEa,omitzero"`
	// Predicted value of foE in MHz.
	FoEp param.Opt[float64] `json:"foEp,omitzero"`
	// Highest ordinary wave frequency at which a mainly continuous Es trace is
	// observed, in MHz. URSI ID: 30.
	FoEs param.Opt[float64] `json:"foEs,omitzero"`
	// The ordinary wave F1 critical frequency, in MHz. URSI ID: 10.
	FoF1 param.Opt[float64] `json:"foF1,omitzero"`
	// Predicted value of foF1 in MHz.
	FoF1p param.Opt[float64] `json:"foF1p,omitzero"`
	// The ordinary wave critical frequency of the highest stratification in the F
	// region, specified in MHz. URSI ID: 00.
	FoF2 param.Opt[float64] `json:"foF2,omitzero"`
	// Predicted value of foF2 in MHz.
	FoF2p param.Opt[float64] `json:"foF2p,omitzero"`
	// Highest ordinary wave critical frequency of F region patch trace in MHz. URSI
	// ID: 55.
	FoP param.Opt[float64] `json:"foP,omitzero"`
	// The extraordinary wave E critical frequency, in MHz. URSI ID: 21.
	FxE param.Opt[float64] `json:"fxE,omitzero"`
	// The extraordinary wave F1 critical frequency, in MHz. URSI ID: 11.
	FxF1 param.Opt[float64] `json:"fxF1,omitzero"`
	// The extraordinary wave F2 critical frequency, in MHz. URSI ID: 01.
	FxF2 param.Opt[float64] `json:"fxF2,omitzero"`
	// The highest frequency of F-trace in MHz. Note: fxI is with capital i. URSI
	// ID: 51.
	FxI param.Opt[float64] `json:"fxI,omitzero"`
	// True height of the E peak in km. URSI ID: CE.
	HmE param.Opt[float64] `json:"hmE,omitzero"`
	// True height of the F1 peak in km. URSI ID: BE.
	HmF1 param.Opt[float64] `json:"hmF1,omitzero"`
	// True height of the F2 peak in km. URSI ID: AE.
	HmF2 param.Opt[float64] `json:"hmF2,omitzero"`
	// The minimum virtual height of the normal E layer trace in km. URSI ID: 24.
	HprimeE param.Opt[float64] `json:"hprimeE,omitzero"`
	// Minimum virtual height of night time auroral E layer trace in km. URSI ID: 27.
	HprimeEa param.Opt[float64] `json:"hprimeEa,omitzero"`
	// The minimum height of the trace used to give foEs in km. URSI ID: 34.
	HprimeEs param.Opt[float64] `json:"hprimeEs,omitzero"`
	// The minimum virtual height of the ordinary wave trace taken as a whole, in km.
	// URSI ID: 16.
	HprimeF param.Opt[float64] `json:"hprimeF,omitzero"`
	// The minimum virtual height of reflection at a point where the trace is
	// horizontal in the F region in km. URSI ID: 14.
	HprimeF1 param.Opt[float64] `json:"hprimeF1,omitzero"`
	// The minimum virtual height of ordinary wave trace for the highest stable
	// stratification in the F region in km. URSI ID: 4.
	HprimeF2 param.Opt[float64] `json:"hprimeF2,omitzero"`
	// Virtual height at MUF/OblFactor frequency in MHz.
	HprimefMuf param.Opt[float64] `json:"hprimefMUF,omitzero"`
	// Minimum virtual height of the trace used to determine foP in km. URSI ID: 56.
	HprimeP param.Opt[float64] `json:"hprimeP,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Lowest usable frequency.
	Luf param.Opt[float64] `json:"luf,omitzero"`
	// MUF(D)/foF2.
	Md param.Opt[float64] `json:"md,omitzero"`
	// Maximum Usable Frequency for ground distance D in MHz.
	Mufd param.Opt[float64] `json:"mufd,omitzero"`
	// Name of the algorithm used for the electron density profile.
	NeProfileName param.Opt[string] `json:"neProfileName,omitzero"`
	// Version of the algorithm used for the electron density profile.
	NeProfileVersion param.Opt[float64] `json:"neProfileVersion,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Equipment location.
	PlatformName param.Opt[string] `json:"platformName,omitzero"`
	// Average range spread of E layer in km. URSI ID: 85.
	Qe param.Opt[float64] `json:"qe,omitzero"`
	// Average range spread of F layer in km. URSI ID: 84.
	Qf param.Opt[float64] `json:"qf,omitzero"`
	// Notes for the restrictedFrequency data.
	RestrictedFrequencyNotes param.Opt[string] `json:"restrictedFrequencyNotes,omitzero"`
	// Effective scale height at hmF2 Titheridge method in km. URSI ID: 69.
	ScaleHeightF2Peak param.Opt[float64] `json:"scaleHeightF2Peak,omitzero"`
	// Details concerning the composition/intention/interpretation/audience/etc. of any
	// data recorded here. This field may contain all of the intended information e.g.
	// info on signal waveforms used, antenna setup, etc. OR may describe the
	// data/settings to be provided in the “data” field.
	SystemNotes param.Opt[string] `json:"systemNotes,omitzero"`
	// Total Ionospheric Electron Content \*10^16e/m^2. 1 TEC Unit (TECU) = 10^16
	// electrons/m^2. URSI ID: 72.
	Tec param.Opt[float64] `json:"tec,omitzero"`
	// Characterization of the shape of Es trace. URSI ID: 36.
	TypeEs param.Opt[string] `json:"typeEs,omitzero"`
	// Parabolic E layer semi-thickness in km. URSI ID: 83.
	YE param.Opt[float64] `json:"yE,omitzero"`
	// Parabolic F1 layer semi-thickness in km. URSI ID: 95.
	YF1 param.Opt[float64] `json:"yF1,omitzero"`
	// Parabolic F2 layer semi-thickness in km. URSI ID: 94.
	YF2 param.Opt[float64] `json:"yF2,omitzero"`
	// True height at half peak electron density in the F2 layer in km. URSI ID: 93.
	ZhalfNm param.Opt[float64] `json:"zhalfNm,omitzero"`
	// Peak height of E-layer in km. URSI ID: 90.
	ZmE                    param.Opt[float64]                                                `json:"zmE,omitzero"`
	Amplitude              IonoObservationUnvalidatedPublishParamsBodyAmplitude              `json:"amplitude,omitzero"`
	AntennaElementPosition IonoObservationUnvalidatedPublishParamsBodyAntennaElementPosition `json:"antennaElementPosition,omitzero"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem string `json:"antennaElementPositionCoordinateSystem,omitzero"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                                            `json:"artistFlags,omitzero"`
	Azimuth     IonoObservationUnvalidatedPublishParamsBodyAzimuth `json:"azimuth,omitzero"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonoObservationUnvalidatedPublishParamsBodyCharAtt `json:"charAtts,omitzero"`
	Datum    IonoObservationUnvalidatedPublishParamsBodyDatum     `json:"datum,omitzero"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonoObservationUnvalidatedPublishParamsBodyDensityProfile `json:"densityProfile,omitzero"`
	Doppler        IonoObservationUnvalidatedPublishParamsBodyDoppler        `json:"doppler,omitzero"`
	// Array of electron densities in cm^-3 (must match the size of the height and
	// plasmaFrequency arrays).
	ElectronDensity []float64 `json:"electronDensity,omitzero"`
	// Uncertainty in specifying the electron density at each height point of the
	// profile (must match the size of the electronDensity array).
	ElectronDensityUncertainty []float64                                            `json:"electronDensityUncertainty,omitzero"`
	Elevation                  IonoObservationUnvalidatedPublishParamsBodyElevation `json:"elevation,omitzero"`
	Frequency                  IonoObservationUnvalidatedPublishParamsBodyFrequency `json:"frequency,omitzero"`
	// Array of altitudes above station level for plasma frequency/density arrays in km
	// (must match the size of the plasmaFrequency and electronDensity Arrays).
	Height []float64                                        `json:"height,omitzero"`
	Phase  IonoObservationUnvalidatedPublishParamsBodyPhase `json:"phase,omitzero"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency,omitzero"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64                                               `json:"plasmaFrequencyUncertainty,omitzero"`
	Polarization               IonoObservationUnvalidatedPublishParamsBodyPolarization `json:"polarization,omitzero"`
	Power                      IonoObservationUnvalidatedPublishParamsBodyPower        `json:"power,omitzero"`
	Range                      IonoObservationUnvalidatedPublishParamsBodyRange        `json:"range,omitzero"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	ReceiveCoordinates [][]float64 `json:"receiveCoordinates,omitzero"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	ReceiveSensorType string `json:"receiveSensorType,omitzero"`
	// Array of restricted frequencies.
	RestrictedFrequency []float64 `json:"restrictedFrequency,omitzero"`
	// The ScalerInfo record describes the person or system who interpreted the
	// ionogram in IonoObservation.
	ScalerInfo IonoObservationUnvalidatedPublishParamsBodyScalerInfo `json:"scalerInfo,omitzero"`
	Stokes     IonoObservationUnvalidatedPublishParamsBodyStokes     `json:"stokes,omitzero"`
	// Array of degrees clockwise from true North of the TID.
	TidAzimuth []float64 `json:"tidAzimuth,omitzero"`
	// Array of 1/frequency of the TID wave.
	TidPeriods []float64 `json:"tidPeriods,omitzero"`
	// Array of speed in m/s at which the disturbance travels through the ionosphere.
	TidPhaseSpeeds []float64                                               `json:"tidPhaseSpeeds,omitzero"`
	Time           IonoObservationUnvalidatedPublishParamsBodyTime         `json:"time,omitzero"`
	TraceGeneric   IonoObservationUnvalidatedPublishParamsBodyTraceGeneric `json:"traceGeneric,omitzero"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	TransmitCoordinates [][]float64 `json:"transmitCoordinates,omitzero"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	TransmitSensorType string `json:"transmitSensorType,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IonoObservationUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[IonoObservationUnvalidatedPublishParamsBody](
		"antennaElementPositionCoordinateSystem", "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)", "GEOCentric (lat, long, radii)",
	)
	apijson.RegisterFieldValidator[IonoObservationUnvalidatedPublishParamsBody](
		"receiveSensorType", "Mobile", "Static",
	)
	apijson.RegisterFieldValidator[IonoObservationUnvalidatedPublishParamsBody](
		"transmitSensorType", "Mobile", "Static",
	)
}

type IonoObservationUnvalidatedPublishParamsBodyAmplitude struct {
	// Notes for the amplitude data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of amplitude data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers for amplitude dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyAmplitude) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyAmplitude
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyAmplitude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyAntennaElementPosition) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyAntennaElementPosition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyAntennaElementPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyAzimuth struct {
	// Notes for the azimuth data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of incoming azimuth at the receiver.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the azimuth array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyAzimuth) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyAzimuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyAzimuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Characteristic attributes of a IonoObservation.
type IonoObservationUnvalidatedPublishParamsBodyCharAtt struct {
	// Characteristic name. This value should reflect the UDL field name for the
	// corresponding characteristic.
	CharName param.Opt[string] `json:"charName,omitzero"`
	// Name of the climate model.
	ClimateModelName param.Opt[string] `json:"climateModelName,omitzero"`
	// Descriptive letter (D) for the characteristic specified by URSI ID. Describes
	// specific ionospheric conditions, beyond numerical values.
	D param.Opt[string] `json:"d,omitzero" format:"char"`
	// Specified characteristic's lower bound. Should be less than or equal to the
	// characteristic's current value as set in this record.
	LowerBound param.Opt[float64] `json:"lowerBound,omitzero"`
	// Qualifying letter (Q) for the characteristic specified by URSI ID. Describes
	// specific ionospheric conditions, beyond numerical values.
	Q param.Opt[string] `json:"q,omitzero" format:"char"`
	// Uncertainty Bounds (lower and upper) define an interval around reported value
	// that contains true value at the specified probability level. Probability levels
	// are specified in terms of percentile (from 1 to 100) or the standard deviation,
	// sigma (e.g. 1sigma, 2sigma, 3sigma, 5percentile, 10percentile, 25percentile).
	UncertaintyBoundType param.Opt[string] `json:"uncertaintyBoundType,omitzero"`
	// Specified characteristic's upper bound. Should be greater than or equal to the
	// characteristic's current value as set in this record.
	UpperBound param.Opt[float64] `json:"upperBound,omitzero"`
	// Characteristic's URSI ID. See the characteristic's description for its
	// corresponding URSI ID.
	UrsiID param.Opt[string] `json:"ursiID,omitzero"`
	// Input parameters for the climate model.
	ClimateModelInputParams []string `json:"climateModelInputParams,omitzero"`
	// List of options for the climate model.
	ClimateModelOptions []string `json:"climateModelOptions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyCharAtt) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyCharAtt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyCharAtt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyDatum struct {
	// Notes for the datum with details of what the data is, units, etc.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array to support sparse data collections.
	Data []float64 `json:"data,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDatum) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDatum
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDatum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonoObservationUnvalidatedPublishParamsBodyDensityProfile struct {
	// Description of the valley model and parameters.
	ValleyModelDescription param.Opt[string] `json:"valleyModelDescription,omitzero"`
	// Full set of the IRI formalism coefficients.
	Iri IonoObservationUnvalidatedPublishParamsBodyDensityProfileIri `json:"iri,omitzero"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolic `json:"parabolic,omitzero"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic `json:"quasiParabolic,omitzero"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev `json:"shiftedChebyshev,omitzero"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst,omitzero"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap,omitzero"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfile) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full set of the IRI formalism coefficients.
type IonoObservationUnvalidatedPublishParamsBodyDensityProfileIri struct {
	// B0 parameter of the F2 layer shape.
	B0 param.Opt[float64] `json:"b0,omitzero"`
	// B1 parameter of the F2 layer shape.
	B1 param.Opt[float64] `json:"b1,omitzero"`
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi param.Opt[float64] `json:"chi,omitzero"`
	// D1 parameter of the F1 layer shape.
	D1 param.Opt[float64] `json:"d1,omitzero"`
	// Description of IRI implementation.
	Description param.Opt[string] `json:"description,omitzero"`
	// TBD.
	Fp1 param.Opt[float64] `json:"fp1,omitzero"`
	// TBD.
	Fp2 param.Opt[float64] `json:"fp2,omitzero"`
	// TBD.
	Fp30 param.Opt[float64] `json:"fp30,omitzero"`
	// TBD.
	Fp3U param.Opt[float64] `json:"fp3U,omitzero"`
	// Starting height of the D layer, in kilometers.
	Ha param.Opt[float64] `json:"ha,omitzero"`
	// Height of the intermediate region at the top of D region, in kilometers.
	Hdx param.Opt[float64] `json:"hdx,omitzero"`
	// Peak height of the D layer, in kilometers.
	HmD param.Opt[float64] `json:"hmD,omitzero"`
	// Peak height of the F2 layer, in kilometers.
	HmE param.Opt[float64] `json:"hmE,omitzero"`
	// Peak height of the F1 layer, in kilometers.
	HmF1 param.Opt[float64] `json:"hmF1,omitzero"`
	// Peak height of F2 layer, in kilometers.
	HmF2 param.Opt[float64] `json:"hmF2,omitzero"`
	// The valley height, in kilometers.
	HValTop param.Opt[float64] `json:"hValTop,omitzero"`
	// Height HZ of the interim layer, in kilometers.
	Hz param.Opt[float64] `json:"hz,omitzero"`
	// Peak density of the D layer, in per cubic centimeter.
	NmD param.Opt[float64] `json:"nmD,omitzero"`
	// Peak density of the E layer, in per cubic centimeter.
	NmE param.Opt[float64] `json:"nmE,omitzero"`
	// Peak density of the F1 layer, in grams per cubic centimeter.
	NmF1 param.Opt[float64] `json:"nmF1,omitzero"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 param.Opt[float64] `json:"nmF2,omitzero"`
	// The valley depth, in grams per cubic centimeter.
	NValB param.Opt[float64] `json:"nValB,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfileIri) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfileIri
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfileIri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem `json:"parabolicItems,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem struct {
	// Plasma frequency at the layer peak, in MHz.
	F param.Opt[float64] `json:"f,omitzero"`
	// Ionospheric plasma layer (E, F1, or F2).
	Layer param.Opt[string] `json:"layer,omitzero"`
	// Half-thickness of the layer, in kilometers.
	Y param.Opt[float64] `json:"y,omitzero"`
	// Height of the layer peak, in kilometers.
	Z param.Opt[float64] `json:"z,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius param.Opt[float64] `json:"earthRadius,omitzero"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment struct {
	// Coefficient A.
	A param.Opt[float64] `json:"a,omitzero"`
	// Coefficient B.
	B param.Opt[float64] `json:"b,omitzero"`
	// Coefficient C.
	C param.Opt[float64] `json:"c,omitzero"`
	// Best-fit error.
	Error param.Opt[float64] `json:"error,omitzero"`
	// The index of this segment in the list, from 1 to NumSegments.
	Index param.Opt[int64] `json:"index,omitzero"`
	// Starting range of the segment, in kilometers from the Earth's center.
	Rb param.Opt[float64] `json:"rb,omitzero"`
	// Ending range of the segment, in kilometers from the Earth's center.
	Re param.Opt[float64] `json:"re,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev struct {
	// Best fit error.
	Error param.Opt[float64] `json:"error,omitzero"`
	// Start frequency of the layer, in MHz.
	Fstart param.Opt[float64] `json:"fstart,omitzero"`
	// Stop frequency of the layer, in MHz.
	Fstop param.Opt[float64] `json:"fstop,omitzero"`
	// The ionospheric plasma layer.
	Layer param.Opt[string] `json:"layer,omitzero"`
	// Number of coefficients in the expansion.
	N param.Opt[int64] `json:"n,omitzero"`
	// Peak height of the layer, in kilometers.
	Peakheight param.Opt[float64] `json:"peakheight,omitzero"`
	// Height at which density is half of the peak Nm, in kilometers.
	ZhalfNm param.Opt[float64] `json:"zhalfNm,omitzero"`
	// Array of coefficients.
	Coeffs []float64 `json:"coeffs,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters of the constant-scale-height Chapman layer.
type IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst struct {
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi param.Opt[float64] `json:"chi,omitzero"`
	// Description of the Chapman computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Peak height of F2 layer, in kilometers.
	HmF2 param.Opt[float64] `json:"hmF2,omitzero"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 param.Opt[float64] `json:"nmF2,omitzero"`
	// Scale height if F2 layer at the peak, in kilometers.
	ScaleF2 param.Opt[float64] `json:"scaleF2,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Varying scale height Chapman topside layer.
type IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap struct {
	// Alpha parameter of the profile shape.
	Alpha param.Opt[float64] `json:"alpha,omitzero"`
	// Beta parameter of the profile shape.
	Beta param.Opt[float64] `json:"beta,omitzero"`
	// Peak Density Thickness (PDT) for description of the flat-nose shape, in
	// kilometers.
	Chi param.Opt[float64] `json:"chi,omitzero"`
	// Description of the Chapman computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Peak height of F2 layer, in kilometers.
	HmF2 param.Opt[float64] `json:"hmF2,omitzero"`
	// Transition height, in kilometers.
	Ht param.Opt[float64] `json:"ht,omitzero"`
	// Peak density of F2 layer, in grams per cubic centimeter.
	NmF2 param.Opt[float64] `json:"nmF2,omitzero"`
	// Scale height if F2 layer at the peak, in kilometers.
	ScaleF2 param.Opt[float64] `json:"scaleF2,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyDoppler struct {
	// Notes for the doppler data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of received doppler shifts in Hz.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the doppler array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyDoppler) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyDoppler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyDoppler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyElevation struct {
	// Notes for the elevation data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of incoming elevation at the receiver.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the elevation array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyElevation) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyElevation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyElevation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyFrequency struct {
	// Notes for the frequency data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of frequency data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for frequency dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the frequency array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyFrequency) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyFrequency
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyFrequency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyPhase struct {
	// Notes for the phase data. Orientation and position for each antenna element
	// across the antenna_element dimension.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of phase data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for phase dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the phase array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyPhase) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyPhase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyPhase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyPolarization struct {
	// Notes for the polarization data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of polarization data.
	Data [][][][][][][]string `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers for polarization dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyPolarization) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyPolarization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyPolarization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyPower struct {
	// Notes for the power data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of received power in db.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the power array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyPower) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyPower
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyPower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyRange struct {
	// Notes for the range data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of range data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for range dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the range array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyRange) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonoObservationUnvalidatedPublishParamsBodyScalerInfo struct {
	// Scaler confidence level.
	ConfidenceLevel param.Opt[int64] `json:"confidenceLevel,omitzero"`
	// Scaler confidence score.
	ConfidenceScore param.Opt[int64] `json:"confidenceScore,omitzero"`
	// Scaler name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Scaler organization.
	Organization param.Opt[string] `json:"organization,omitzero"`
	// Scaler type (MANUAL, AUTOMATIC or UNKNOWN).
	Type param.Opt[string] `json:"type,omitzero"`
	// Scaler version.
	Version param.Opt[float64] `json:"version,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyScalerInfo) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyScalerInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyScalerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyStokes struct {
	// Notes for the stokes data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of received stokes data.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the stoke array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	// S1, S2, and S3 (the normalized Stokes parameters 1, 2, and 3).
	S []float64 `json:"s,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyStokes) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyStokes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyStokes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyTime struct {
	// The notes indicate the scheme and accuracy.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array of times in number of seconds passed since January 1st, 1970 with the same
	// dimensions as power.
	Data [][][][][][][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the time array dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyTime) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyTime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationUnvalidatedPublishParamsBodyTraceGeneric struct {
	// Notes for the trace generic data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data,omitzero"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName,omitzero"`
	paramObj
}

func (r IonoObservationUnvalidatedPublishParamsBodyTraceGeneric) MarshalJSON() (data []byte, err error) {
	type shadow IonoObservationUnvalidatedPublishParamsBodyTraceGeneric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IonoObservationUnvalidatedPublishParamsBodyTraceGeneric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
