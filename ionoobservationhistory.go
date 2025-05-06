// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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

// IonoObservationHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIonoObservationHistoryService] method instead.
type IonoObservationHistoryService struct {
	Options []option.RequestOption
}

// NewIonoObservationHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIonoObservationHistoryService(opts ...option.RequestOption) (r IonoObservationHistoryService) {
	r = IonoObservationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *IonoObservationHistoryService) List(ctx context.Context, query IonoObservationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[IonoObservationHistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/ionoobservation/history"
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

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *IonoObservationHistoryService) ListAutoPaging(ctx context.Context, query IonoObservationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[IonoObservationHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *IonoObservationHistoryService) Aodr(ctx context.Context, query IonoObservationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/ionoobservation/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *IonoObservationHistoryService) Count(ctx context.Context, query IonoObservationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/ionoobservation/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// These services provide operations for posting and querying ionospheric
// observation data. Characteristics are defined by the CHARS: URSI IIWG format for
// archiving monthly ionospheric characteristics, INAG Bulletin No. 62
// specification. Qualifying and Descriptive letters are defined by the URSI
// Handbook for Ionogram Interpretation and reduction, Report UAG, No. 23A
// specification.
type IonoObservationHistoryListResponse struct {
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
	DataMode IonoObservationHistoryListResponseDataMode `json:"dataMode,required"`
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
	ID                     string                                                   `json:"id"`
	Amplitude              IonoObservationHistoryListResponseAmplitude              `json:"amplitude"`
	AntennaElementPosition IonoObservationHistoryListResponseAntennaElementPosition `json:"antennaElementPosition"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystem `json:"antennaElementPositionCoordinateSystem"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                                   `json:"artistFlags"`
	Azimuth     IonoObservationHistoryListResponseAzimuth `json:"azimuth"`
	// IRI thickness parameter in km. URSI ID: D0.
	B0 float64 `json:"b0"`
	// IRI profile shape parameter. URSI ID: D1.
	B1 float64 `json:"b1"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonoObservationHistoryListResponseCharAtt `json:"charAtts"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Distance for MUF calculation in km.
	D float64 `json:"d"`
	// IRI profile shape parameter, F1 layer. URSI ID: D2.
	D1    float64                                 `json:"d1"`
	Datum IonoObservationHistoryListResponseDatum `json:"datum"`
	// Adjustment to the scaled foF2 during profile inversion in MHz.
	DeltafoF2 float64 `json:"deltafoF2"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonoObservationHistoryListResponseDensityProfile `json:"densityProfile"`
	Doppler        IonoObservationHistoryListResponseDoppler        `json:"doppler"`
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
	ElectronDensityUncertainty []float64                                   `json:"electronDensityUncertainty"`
	Elevation                  IonoObservationHistoryListResponseElevation `json:"elevation"`
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
	FoP       float64                                     `json:"foP"`
	Frequency IonoObservationHistoryListResponseFrequency `json:"frequency"`
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
	OrigSensorID string                                  `json:"origSensorId"`
	Phase        IonoObservationHistoryListResponsePhase `json:"phase"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64 `json:"plasmaFrequencyUncertainty"`
	// Equipment location.
	PlatformName string                                         `json:"platformName"`
	Polarization IonoObservationHistoryListResponsePolarization `json:"polarization"`
	Power        IonoObservationHistoryListResponsePower        `json:"power"`
	// Average range spread of E layer in km. URSI ID: 85.
	Qe float64 `json:"qe"`
	// Average range spread of F layer in km. URSI ID: 84.
	Qf    float64                                 `json:"qf"`
	Range IonoObservationHistoryListResponseRange `json:"range"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	ReceiveCoordinates [][]float64 `json:"receiveCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	ReceiveSensorType IonoObservationHistoryListResponseReceiveSensorType `json:"receiveSensorType"`
	// Array of restricted frequencies.
	RestrictedFrequency []float64 `json:"restrictedFrequency"`
	// Notes for the restrictedFrequency data.
	RestrictedFrequencyNotes string `json:"restrictedFrequencyNotes"`
	// Effective scale height at hmF2 Titheridge method in km. URSI ID: 69.
	ScaleHeightF2Peak float64 `json:"scaleHeightF2Peak"`
	// The ScalerInfo record describes the person or system who interpreted the
	// ionogram in IonoObservation.
	ScalerInfo IonoObservationHistoryListResponseScalerInfo `json:"scalerInfo"`
	Stokes     IonoObservationHistoryListResponseStokes     `json:"stokes"`
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
	TidPhaseSpeeds []float64                                      `json:"tidPhaseSpeeds"`
	Time           IonoObservationHistoryListResponseTime         `json:"time"`
	TraceGeneric   IonoObservationHistoryListResponseTraceGeneric `json:"traceGeneric"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	TransmitCoordinates [][]float64 `json:"transmitCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	TransmitSensorType IonoObservationHistoryListResponseTransmitSensorType `json:"transmitSensorType"`
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
func (r IonoObservationHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type IonoObservationHistoryListResponseDataMode string

const (
	IonoObservationHistoryListResponseDataModeReal      IonoObservationHistoryListResponseDataMode = "REAL"
	IonoObservationHistoryListResponseDataModeTest      IonoObservationHistoryListResponseDataMode = "TEST"
	IonoObservationHistoryListResponseDataModeSimulated IonoObservationHistoryListResponseDataMode = "SIMULATED"
	IonoObservationHistoryListResponseDataModeExercise  IonoObservationHistoryListResponseDataMode = "EXERCISE"
)

type IonoObservationHistoryListResponseAmplitude struct {
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
func (r IonoObservationHistoryListResponseAmplitude) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseAmplitude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponseAntennaElementPosition struct {
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
func (r IonoObservationHistoryListResponseAntennaElementPosition) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseAntennaElementPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
// (lat, long, radii).
type IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystem string

const (
	IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystemJ2000                   IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystem = "J2000"
	IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystemEcrEcef                 IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystem = "ECR/ECEF"
	IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystemTeme                    IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystem = "TEME"
	IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystemGcrf                    IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystem = "GCRF"
	IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystemWgs84GeoDeticLatLongAlt IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystem = "WGS84 (GEODetic lat, long, alt)"
	IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystemGeoCentricLatLongRadii  IonoObservationHistoryListResponseAntennaElementPositionCoordinateSystem = "GEOCentric (lat, long, radii)"
)

type IonoObservationHistoryListResponseAzimuth struct {
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
func (r IonoObservationHistoryListResponseAzimuth) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseAzimuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Characteristic attributes of a IonoObservation.
type IonoObservationHistoryListResponseCharAtt struct {
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
func (r IonoObservationHistoryListResponseCharAtt) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseCharAtt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponseDatum struct {
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
func (r IonoObservationHistoryListResponseDatum) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseDatum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonoObservationHistoryListResponseDensityProfile struct {
	// Full set of the IRI formalism coefficients.
	Iri IonoObservationHistoryListResponseDensityProfileIri `json:"iri"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonoObservationHistoryListResponseDensityProfileParabolic `json:"parabolic"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonoObservationHistoryListResponseDensityProfileQuasiParabolic `json:"quasiParabolic"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonoObservationHistoryListResponseDensityProfileShiftedChebyshev `json:"shiftedChebyshev"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonoObservationHistoryListResponseDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonoObservationHistoryListResponseDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap"`
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
func (r IonoObservationHistoryListResponseDensityProfile) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseDensityProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full set of the IRI formalism coefficients.
type IonoObservationHistoryListResponseDensityProfileIri struct {
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
func (r IonoObservationHistoryListResponseDensityProfileIri) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseDensityProfileIri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationHistoryListResponseDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description string `json:"description"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonoObservationHistoryListResponseDensityProfileParabolicParabolicItem `json:"parabolicItems"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		ParabolicItems respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationHistoryListResponseDensityProfileParabolic) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationHistoryListResponseDensityProfileParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonoObservationHistoryListResponseDensityProfileParabolicParabolicItem struct {
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
func (r IonoObservationHistoryListResponseDensityProfileParabolicParabolicItem) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationHistoryListResponseDensityProfileParabolicParabolicItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonoObservationHistoryListResponseDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description string `json:"description"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius float64 `json:"earthRadius"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonoObservationHistoryListResponseDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments"`
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
func (r IonoObservationHistoryListResponseDensityProfileQuasiParabolic) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationHistoryListResponseDensityProfileQuasiParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonoObservationHistoryListResponseDensityProfileQuasiParabolicQuasiParabolicSegment struct {
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
func (r IonoObservationHistoryListResponseDensityProfileQuasiParabolicQuasiParabolicSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationHistoryListResponseDensityProfileQuasiParabolicQuasiParabolicSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonoObservationHistoryListResponseDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description string `json:"description"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonoObservationHistoryListResponseDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description       respjson.Field
		ShiftedChebyshevs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoObservationHistoryListResponseDensityProfileShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationHistoryListResponseDensityProfileShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonoObservationHistoryListResponseDensityProfileShiftedChebyshevShiftedChebyshev struct {
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
func (r IonoObservationHistoryListResponseDensityProfileShiftedChebyshevShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationHistoryListResponseDensityProfileShiftedChebyshevShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters of the constant-scale-height Chapman layer.
type IonoObservationHistoryListResponseDensityProfileTopsideExtensionChapmanConst struct {
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
func (r IonoObservationHistoryListResponseDensityProfileTopsideExtensionChapmanConst) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationHistoryListResponseDensityProfileTopsideExtensionChapmanConst) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Varying scale height Chapman topside layer.
type IonoObservationHistoryListResponseDensityProfileTopsideExtensionVaryChap struct {
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
func (r IonoObservationHistoryListResponseDensityProfileTopsideExtensionVaryChap) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoObservationHistoryListResponseDensityProfileTopsideExtensionVaryChap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponseDoppler struct {
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
func (r IonoObservationHistoryListResponseDoppler) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseDoppler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponseElevation struct {
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
func (r IonoObservationHistoryListResponseElevation) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseElevation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponseFrequency struct {
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
func (r IonoObservationHistoryListResponseFrequency) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseFrequency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponsePhase struct {
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
func (r IonoObservationHistoryListResponsePhase) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponsePhase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponsePolarization struct {
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
func (r IonoObservationHistoryListResponsePolarization) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponsePolarization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponsePower struct {
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
func (r IonoObservationHistoryListResponsePower) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponsePower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponseRange struct {
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
func (r IonoObservationHistoryListResponseRange) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonoObservationHistoryListResponseReceiveSensorType string

const (
	IonoObservationHistoryListResponseReceiveSensorTypeMobile IonoObservationHistoryListResponseReceiveSensorType = "Mobile"
	IonoObservationHistoryListResponseReceiveSensorTypeStatic IonoObservationHistoryListResponseReceiveSensorType = "Static"
)

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonoObservationHistoryListResponseScalerInfo struct {
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
func (r IonoObservationHistoryListResponseScalerInfo) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseScalerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponseStokes struct {
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
func (r IonoObservationHistoryListResponseStokes) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseStokes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponseTime struct {
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
func (r IonoObservationHistoryListResponseTime) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoObservationHistoryListResponseTraceGeneric struct {
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
func (r IonoObservationHistoryListResponseTraceGeneric) RawJSON() string { return r.JSON.raw }
func (r *IonoObservationHistoryListResponseTraceGeneric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonoObservationHistoryListResponseTransmitSensorType string

const (
	IonoObservationHistoryListResponseTransmitSensorTypeMobile IonoObservationHistoryListResponseTransmitSensorType = "Mobile"
	IonoObservationHistoryListResponseTransmitSensorTypeStatic IonoObservationHistoryListResponseTransmitSensorType = "Static"
)

type IonoObservationHistoryListParams struct {
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time `query:"startTimeUTC,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IonoObservationHistoryListParams]'s query parameters as
// `url.Values`.
func (r IonoObservationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonoObservationHistoryAodrParams struct {
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time `query:"startTimeUTC,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// optional, notification method for the created file link. When omitted, EMAIL is
	// assumed. Current valid values are: EMAIL, SMS.
	Notification param.Opt[string] `query:"notification,omitzero" json:"-"`
	// optional, field delimiter when the created file is not JSON. Must be a single
	// character chosen from this set: (',', ';', ':', '|'). When omitted, "," is used.
	// It is strongly encouraged that your field delimiter be a character unlikely to
	// occur within the data.
	OutputDelimiter param.Opt[string] `query:"outputDelimiter,omitzero" json:"-"`
	// optional, output format for the file. When omitted, JSON is assumed. Current
	// valid values are: JSON and CSV.
	OutputFormat param.Opt[string] `query:"outputFormat,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IonoObservationHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r IonoObservationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonoObservationHistoryCountParams struct {
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time        `query:"startTimeUTC,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IonoObservationHistoryCountParams]'s query parameters as
// `url.Values`.
func (r IonoObservationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
