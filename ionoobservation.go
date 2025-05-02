// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// IonoobservationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIonoobservationService] method instead.
type IonoobservationService struct {
	Options []option.RequestOption
}

// NewIonoobservationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIonoobservationService(opts ...option.RequestOption) (r IonoobservationService) {
	r = IonoobservationService{}
	r.Options = opts
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *IonoobservationService) List(ctx context.Context, query IonoobservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[IonoobservationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *IonoobservationService) ListAutoPaging(ctx context.Context, query IonoobservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[IonoobservationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *IonoobservationService) Count(ctx context.Context, query IonoobservationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *IonoobservationService) NewBulk(ctx context.Context, body IonoobservationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/ionoobservation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *IonoobservationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/ionoobservation/queryhelp"
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
func (r *IonoobservationService) Tuple(ctx context.Context, query IonoobservationTupleParams, opts ...option.RequestOption) (res *[]IonoObservationFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/ionoobservation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take Ionospheric Observation entries as a POST body and
// ingest into the database with or without dupe detection. Default is no dupe
// checking. This operation is intended to be used for automated feeds into UDL. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *IonoobservationService) UnvalidatedPublish(ctx context.Context, body IonoobservationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
type IonoobservationListResponse struct {
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
	DataMode IonoobservationListResponseDataMode `json:"dataMode,required"`
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
	Amplitude              IonoobservationListResponseAmplitude              `json:"amplitude"`
	AntennaElementPosition IonoobservationListResponseAntennaElementPosition `json:"antennaElementPosition"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem IonoobservationListResponseAntennaElementPositionCoordinateSystem `json:"antennaElementPositionCoordinateSystem"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                            `json:"artistFlags"`
	Azimuth     IonoobservationListResponseAzimuth `json:"azimuth"`
	// IRI thickness parameter in km. URSI ID: D0.
	B0 float64 `json:"b0"`
	// IRI profile shape parameter. URSI ID: D1.
	B1 float64 `json:"b1"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonoobservationListResponseCharAtt `json:"charAtts"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Distance for MUF calculation in km.
	D float64 `json:"d"`
	// IRI profile shape parameter, F1 layer. URSI ID: D2.
	D1    float64                          `json:"d1"`
	Datum IonoobservationListResponseDatum `json:"datum"`
	// Adjustment to the scaled foF2 during profile inversion in MHz.
	DeltafoF2 float64 `json:"deltafoF2"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonoobservationListResponseDensityProfile `json:"densityProfile"`
	Doppler        IonoobservationListResponseDoppler        `json:"doppler"`
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
	Elevation                  IonoobservationListResponseElevation `json:"elevation"`
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
	Frequency IonoobservationListResponseFrequency `json:"frequency"`
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
	Phase        IonoobservationListResponsePhase `json:"phase"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64 `json:"plasmaFrequencyUncertainty"`
	// Equipment location.
	PlatformName string                                  `json:"platformName"`
	Polarization IonoobservationListResponsePolarization `json:"polarization"`
	Power        IonoobservationListResponsePower        `json:"power"`
	// Average range spread of E layer in km. URSI ID: 85.
	Qe float64 `json:"qe"`
	// Average range spread of F layer in km. URSI ID: 84.
	Qf    float64                          `json:"qf"`
	Range IonoobservationListResponseRange `json:"range"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	ReceiveCoordinates [][]float64 `json:"receiveCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	ReceiveSensorType IonoobservationListResponseReceiveSensorType `json:"receiveSensorType"`
	// Array of restricted frequencies.
	RestrictedFrequency []float64 `json:"restrictedFrequency"`
	// Notes for the restrictedFrequency data.
	RestrictedFrequencyNotes string `json:"restrictedFrequencyNotes"`
	// Effective scale height at hmF2 Titheridge method in km. URSI ID: 69.
	ScaleHeightF2Peak float64 `json:"scaleHeightF2Peak"`
	// The ScalerInfo record describes the person or system who interpreted the
	// ionogram in IonoObservation.
	ScalerInfo IonoobservationListResponseScalerInfo `json:"scalerInfo"`
	Stokes     IonoobservationListResponseStokes     `json:"stokes"`
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
	Time           IonoobservationListResponseTime         `json:"time"`
	TraceGeneric   IonoobservationListResponseTraceGeneric `json:"traceGeneric"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	TransmitCoordinates [][]float64 `json:"transmitCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	TransmitSensorType IonoobservationListResponseTransmitSensorType `json:"transmitSensorType"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                  resp.Field
		DataMode                               resp.Field
		Source                                 resp.Field
		StartTimeUtc                           resp.Field
		StationID                              resp.Field
		System                                 resp.Field
		SystemInfo                             resp.Field
		ID                                     resp.Field
		Amplitude                              resp.Field
		AntennaElementPosition                 resp.Field
		AntennaElementPositionCoordinateSystem resp.Field
		ArtistFlags                            resp.Field
		Azimuth                                resp.Field
		B0                                     resp.Field
		B1                                     resp.Field
		CharAtts                               resp.Field
		CreatedAt                              resp.Field
		CreatedBy                              resp.Field
		D                                      resp.Field
		D1                                     resp.Field
		Datum                                  resp.Field
		DeltafoF2                              resp.Field
		DensityProfile                         resp.Field
		Doppler                                resp.Field
		DownE                                  resp.Field
		DownEs                                 resp.Field
		DownF                                  resp.Field
		ElectronDensity                        resp.Field
		ElectronDensityUncertainty             resp.Field
		Elevation                              resp.Field
		FbEs                                   resp.Field
		Fe                                     resp.Field
		Ff                                     resp.Field
		FhprimeF                               resp.Field
		FhprimeF2                              resp.Field
		Fmin                                   resp.Field
		FminE                                  resp.Field
		FminEs                                 resp.Field
		FminF                                  resp.Field
		Fmuf                                   resp.Field
		FoE                                    resp.Field
		FoEa                                   resp.Field
		FoEp                                   resp.Field
		FoEs                                   resp.Field
		FoF1                                   resp.Field
		FoF1p                                  resp.Field
		FoF2                                   resp.Field
		FoF2p                                  resp.Field
		FoP                                    resp.Field
		Frequency                              resp.Field
		FxE                                    resp.Field
		FxF1                                   resp.Field
		FxF2                                   resp.Field
		FxI                                    resp.Field
		Height                                 resp.Field
		HmE                                    resp.Field
		HmF1                                   resp.Field
		HmF2                                   resp.Field
		HprimeE                                resp.Field
		HprimeEa                               resp.Field
		HprimeEs                               resp.Field
		HprimeF                                resp.Field
		HprimeF1                               resp.Field
		HprimeF2                               resp.Field
		HprimefMuf                             resp.Field
		HprimeP                                resp.Field
		IDSensor                               resp.Field
		Luf                                    resp.Field
		Md                                     resp.Field
		Mufd                                   resp.Field
		NeProfileName                          resp.Field
		NeProfileVersion                       resp.Field
		Origin                                 resp.Field
		OrigNetwork                            resp.Field
		OrigSensorID                           resp.Field
		Phase                                  resp.Field
		PlasmaFrequency                        resp.Field
		PlasmaFrequencyUncertainty             resp.Field
		PlatformName                           resp.Field
		Polarization                           resp.Field
		Power                                  resp.Field
		Qe                                     resp.Field
		Qf                                     resp.Field
		Range                                  resp.Field
		ReceiveCoordinates                     resp.Field
		ReceiveSensorType                      resp.Field
		RestrictedFrequency                    resp.Field
		RestrictedFrequencyNotes               resp.Field
		ScaleHeightF2Peak                      resp.Field
		ScalerInfo                             resp.Field
		Stokes                                 resp.Field
		SystemNotes                            resp.Field
		Tec                                    resp.Field
		TidAzimuth                             resp.Field
		TidPeriods                             resp.Field
		TidPhaseSpeeds                         resp.Field
		Time                                   resp.Field
		TraceGeneric                           resp.Field
		TransmitCoordinates                    resp.Field
		TransmitSensorType                     resp.Field
		TypeEs                                 resp.Field
		UpdatedAt                              resp.Field
		UpdatedBy                              resp.Field
		YE                                     resp.Field
		YF1                                    resp.Field
		YF2                                    resp.Field
		ZhalfNm                                resp.Field
		ZmE                                    resp.Field
		ExtraFields                            map[string]resp.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponse) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponse) UnmarshalJSON(data []byte) error {
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
type IonoobservationListResponseDataMode string

const (
	IonoobservationListResponseDataModeReal      IonoobservationListResponseDataMode = "REAL"
	IonoobservationListResponseDataModeTest      IonoobservationListResponseDataMode = "TEST"
	IonoobservationListResponseDataModeSimulated IonoobservationListResponseDataMode = "SIMULATED"
	IonoobservationListResponseDataModeExercise  IonoobservationListResponseDataMode = "EXERCISE"
)

type IonoobservationListResponseAmplitude struct {
	// Array of amplitude data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for amplitude dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the amplitude data.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseAmplitude) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseAmplitude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponseAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseAntennaElementPosition) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseAntennaElementPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
// (lat, long, radii).
type IonoobservationListResponseAntennaElementPositionCoordinateSystem string

const (
	IonoobservationListResponseAntennaElementPositionCoordinateSystemJ2000                   IonoobservationListResponseAntennaElementPositionCoordinateSystem = "J2000"
	IonoobservationListResponseAntennaElementPositionCoordinateSystemEcrEcef                 IonoobservationListResponseAntennaElementPositionCoordinateSystem = "ECR/ECEF"
	IonoobservationListResponseAntennaElementPositionCoordinateSystemTeme                    IonoobservationListResponseAntennaElementPositionCoordinateSystem = "TEME"
	IonoobservationListResponseAntennaElementPositionCoordinateSystemGcrf                    IonoobservationListResponseAntennaElementPositionCoordinateSystem = "GCRF"
	IonoobservationListResponseAntennaElementPositionCoordinateSystemWgs84GeoDeticLatLongAlt IonoobservationListResponseAntennaElementPositionCoordinateSystem = "WGS84 (GEODetic lat, long, alt)"
	IonoobservationListResponseAntennaElementPositionCoordinateSystemGeoCentricLatLongRadii  IonoobservationListResponseAntennaElementPositionCoordinateSystem = "GEOCentric (lat, long, radii)"
)

type IonoobservationListResponseAzimuth struct {
	// Array of incoming azimuth at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the azimuth array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the azimuth data.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseAzimuth) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseAzimuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Characteristic attributes of a IonoObservation.
type IonoobservationListResponseCharAtt struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CharName                resp.Field
		ClimateModelInputParams resp.Field
		ClimateModelName        resp.Field
		ClimateModelOptions     resp.Field
		D                       resp.Field
		LowerBound              resp.Field
		Q                       resp.Field
		UncertaintyBoundType    resp.Field
		UpperBound              resp.Field
		UrsiID                  resp.Field
		ExtraFields             map[string]resp.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseCharAtt) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseCharAtt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponseDatum struct {
	// Array to support sparse data collections.
	Data []float64 `json:"data"`
	// Notes for the datum with details of what the data is, units, etc.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data        resp.Field
		Notes       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDatum) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseDatum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonoobservationListResponseDensityProfile struct {
	// Full set of the IRI formalism coefficients.
	Iri IonoobservationListResponseDensityProfileIri `json:"iri"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonoobservationListResponseDensityProfileParabolic `json:"parabolic"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonoobservationListResponseDensityProfileQuasiParabolic `json:"quasiParabolic"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonoobservationListResponseDensityProfileShiftedChebyshev `json:"shiftedChebyshev"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonoobservationListResponseDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonoobservationListResponseDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs"`
	// Description of the valley model and parameters.
	ValleyModelDescription string `json:"valleyModelDescription"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Iri                          resp.Field
		Parabolic                    resp.Field
		QuasiParabolic               resp.Field
		ShiftedChebyshev             resp.Field
		TopsideExtensionChapmanConst resp.Field
		TopsideExtensionVaryChap     resp.Field
		ValleyModelCoeffs            resp.Field
		ValleyModelDescription       resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfile) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseDensityProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full set of the IRI formalism coefficients.
type IonoobservationListResponseDensityProfileIri struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		B0          resp.Field
		B1          resp.Field
		Chi         resp.Field
		D1          resp.Field
		Description resp.Field
		Fp1         resp.Field
		Fp2         resp.Field
		Fp30        resp.Field
		Fp3U        resp.Field
		Ha          resp.Field
		Hdx         resp.Field
		HmD         resp.Field
		HmE         resp.Field
		HmF1        resp.Field
		HmF2        resp.Field
		HValTop     resp.Field
		Hz          resp.Field
		NmD         resp.Field
		NmE         resp.Field
		NmF1        resp.Field
		NmF2        resp.Field
		NValB       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfileIri) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseDensityProfileIri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonoobservationListResponseDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description string `json:"description"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonoobservationListResponseDensityProfileParabolicParabolicItem `json:"parabolicItems"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Description    resp.Field
		ParabolicItems resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfileParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseDensityProfileParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonoobservationListResponseDensityProfileParabolicParabolicItem struct {
	// Plasma frequency at the layer peak, in MHz.
	F float64 `json:"f"`
	// Ionospheric plasma layer (E, F1, or F2).
	Layer string `json:"layer"`
	// Half-thickness of the layer, in kilometers.
	Y float64 `json:"y"`
	// Height of the layer peak, in kilometers.
	Z float64 `json:"z"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		F           resp.Field
		Layer       resp.Field
		Y           resp.Field
		Z           resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfileParabolicParabolicItem) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoobservationListResponseDensityProfileParabolicParabolicItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonoobservationListResponseDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description string `json:"description"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius float64 `json:"earthRadius"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonoobservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Description            resp.Field
		EarthRadius            resp.Field
		QuasiParabolicSegments resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfileQuasiParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseDensityProfileQuasiParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonoobservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		A           resp.Field
		B           resp.Field
		C           resp.Field
		Error       resp.Field
		Index       resp.Field
		Rb          resp.Field
		Re          resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoobservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonoobservationListResponseDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description string `json:"description"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonoobservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Description       resp.Field
		ShiftedChebyshevs resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfileShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoobservationListResponseDensityProfileShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonoobservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Coeffs      resp.Field
		Error       resp.Field
		Fstart      resp.Field
		Fstop       resp.Field
		Layer       resp.Field
		N           resp.Field
		Peakheight  resp.Field
		ZhalfNm     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoobservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters of the constant-scale-height Chapman layer.
type IonoobservationListResponseDensityProfileTopsideExtensionChapmanConst struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Chi         resp.Field
		Description resp.Field
		HmF2        resp.Field
		NmF2        resp.Field
		ScaleF2     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfileTopsideExtensionChapmanConst) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoobservationListResponseDensityProfileTopsideExtensionChapmanConst) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Varying scale height Chapman topside layer.
type IonoobservationListResponseDensityProfileTopsideExtensionVaryChap struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Alpha       resp.Field
		Beta        resp.Field
		Chi         resp.Field
		Description resp.Field
		HmF2        resp.Field
		Ht          resp.Field
		NmF2        resp.Field
		ScaleF2     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDensityProfileTopsideExtensionVaryChap) RawJSON() string {
	return r.JSON.raw
}
func (r *IonoobservationListResponseDensityProfileTopsideExtensionVaryChap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponseDoppler struct {
	// Array of received doppler shifts in Hz.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the doppler array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the doppler data.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseDoppler) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseDoppler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponseElevation struct {
	// Array of incoming elevation at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the elevation array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the elevation data.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseElevation) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseElevation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponseFrequency struct {
	// Array of frequency data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for frequency dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the frequency array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the frequency data.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseFrequency) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseFrequency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponsePhase struct {
	// Array of phase data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for phase dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the phase array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the phase data. Orientation and position for each antenna element
	// across the antenna_element dimension.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponsePhase) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponsePhase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponsePolarization struct {
	// Array of polarization data.
	Data [][][][][][][]string `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for polarization dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the polarization data.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponsePolarization) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponsePolarization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponsePower struct {
	// Array of received power in db.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the power array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the power data.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponsePower) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponsePower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponseRange struct {
	// Array of range data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for range dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the range array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the range data.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseRange) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonoobservationListResponseReceiveSensorType string

const (
	IonoobservationListResponseReceiveSensorTypeMobile IonoobservationListResponseReceiveSensorType = "Mobile"
	IonoobservationListResponseReceiveSensorTypeStatic IonoobservationListResponseReceiveSensorType = "Static"
)

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonoobservationListResponseScalerInfo struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ConfidenceLevel resp.Field
		ConfidenceScore resp.Field
		Name            resp.Field
		Organization    resp.Field
		Type            resp.Field
		Version         resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseScalerInfo) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseScalerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponseStokes struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		S             resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseStokes) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseStokes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponseTime struct {
	// Array of times in number of seconds passed since January 1st, 1970 with the same
	// dimensions as power.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the time array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// The notes indicate the scheme and accuracy.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseTime) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonoobservationListResponseTraceGeneric struct {
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName"`
	// Notes for the trace generic data.
	Notes string `json:"notes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonoobservationListResponseTraceGeneric) RawJSON() string { return r.JSON.raw }
func (r *IonoobservationListResponseTraceGeneric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonoobservationListResponseTransmitSensorType string

const (
	IonoobservationListResponseTransmitSensorTypeMobile IonoobservationListResponseTransmitSensorType = "Mobile"
	IonoobservationListResponseTransmitSensorTypeStatic IonoobservationListResponseTransmitSensorType = "Static"
)

type IonoobservationListParams struct {
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time        `query:"startTimeUTC,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [IonoobservationListParams]'s query parameters as
// `url.Values`.
func (r IonoobservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonoobservationCountParams struct {
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time        `query:"startTimeUTC,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [IonoobservationCountParams]'s query parameters as
// `url.Values`.
func (r IonoobservationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonoobservationNewBulkParams struct {
	Body []IonoobservationNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r IonoobservationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
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
type IonoobservationNewBulkParamsBody struct {
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
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
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
	// Time the row was updated in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
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
	Amplitude              IonoobservationNewBulkParamsBodyAmplitude              `json:"amplitude,omitzero"`
	AntennaElementPosition IonoobservationNewBulkParamsBodyAntennaElementPosition `json:"antennaElementPosition,omitzero"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem string `json:"antennaElementPositionCoordinateSystem,omitzero"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                                 `json:"artistFlags,omitzero"`
	Azimuth     IonoobservationNewBulkParamsBodyAzimuth `json:"azimuth,omitzero"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonoobservationNewBulkParamsBodyCharAtt `json:"charAtts,omitzero"`
	Datum    IonoobservationNewBulkParamsBodyDatum     `json:"datum,omitzero"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonoobservationNewBulkParamsBodyDensityProfile `json:"densityProfile,omitzero"`
	Doppler        IonoobservationNewBulkParamsBodyDoppler        `json:"doppler,omitzero"`
	// Array of electron densities in cm^-3 (must match the size of the height and
	// plasmaFrequency arrays).
	ElectronDensity []float64 `json:"electronDensity,omitzero"`
	// Uncertainty in specifying the electron density at each height point of the
	// profile (must match the size of the electronDensity array).
	ElectronDensityUncertainty []float64                                 `json:"electronDensityUncertainty,omitzero"`
	Elevation                  IonoobservationNewBulkParamsBodyElevation `json:"elevation,omitzero"`
	Frequency                  IonoobservationNewBulkParamsBodyFrequency `json:"frequency,omitzero"`
	// Array of altitudes above station level for plasma frequency/density arrays in km
	// (must match the size of the plasmaFrequency and electronDensity Arrays).
	Height []float64                             `json:"height,omitzero"`
	Phase  IonoobservationNewBulkParamsBodyPhase `json:"phase,omitzero"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency,omitzero"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64                                    `json:"plasmaFrequencyUncertainty,omitzero"`
	Polarization               IonoobservationNewBulkParamsBodyPolarization `json:"polarization,omitzero"`
	Power                      IonoobservationNewBulkParamsBodyPower        `json:"power,omitzero"`
	Range                      IonoobservationNewBulkParamsBodyRange        `json:"range,omitzero"`
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
	ScalerInfo IonoobservationNewBulkParamsBodyScalerInfo `json:"scalerInfo,omitzero"`
	Stokes     IonoobservationNewBulkParamsBodyStokes     `json:"stokes,omitzero"`
	// Array of degrees clockwise from true North of the TID.
	TidAzimuth []float64 `json:"tidAzimuth,omitzero"`
	// Array of 1/frequency of the TID wave.
	TidPeriods []float64 `json:"tidPeriods,omitzero"`
	// Array of speed in m/s at which the disturbance travels through the ionosphere.
	TidPhaseSpeeds []float64                                    `json:"tidPhaseSpeeds,omitzero"`
	Time           IonoobservationNewBulkParamsBodyTime         `json:"time,omitzero"`
	TraceGeneric   IonoobservationNewBulkParamsBodyTraceGeneric `json:"traceGeneric,omitzero"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r IonoobservationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[IonoobservationNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[IonoobservationNewBulkParamsBody](
		"AntennaElementPositionCoordinateSystem", false, "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)", "GEOCentric (lat, long, radii)",
	)
	apijson.RegisterFieldValidator[IonoobservationNewBulkParamsBody](
		"ReceiveSensorType", false, "Mobile", "Static",
	)
	apijson.RegisterFieldValidator[IonoobservationNewBulkParamsBody](
		"TransmitSensorType", false, "Mobile", "Static",
	)
}

type IonoobservationNewBulkParamsBodyAmplitude struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyAmplitude) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyAmplitude) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyAmplitude
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyAntennaElementPosition) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyAntennaElementPosition) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyAntennaElementPosition
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyAzimuth struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyAzimuth) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyAzimuth) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyAzimuth
	return param.MarshalObject(r, (*shadow)(&r))
}

// Characteristic attributes of a IonoObservation.
type IonoobservationNewBulkParamsBodyCharAtt struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyCharAtt) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyCharAtt) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyCharAtt
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyDatum struct {
	// Notes for the datum with details of what the data is, units, etc.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array to support sparse data collections.
	Data []float64 `json:"data,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDatum) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDatum) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDatum
	return param.MarshalObject(r, (*shadow)(&r))
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonoobservationNewBulkParamsBodyDensityProfile struct {
	// Description of the valley model and parameters.
	ValleyModelDescription param.Opt[string] `json:"valleyModelDescription,omitzero"`
	// Full set of the IRI formalism coefficients.
	Iri IonoobservationNewBulkParamsBodyDensityProfileIri `json:"iri,omitzero"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonoobservationNewBulkParamsBodyDensityProfileParabolic `json:"parabolic,omitzero"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolic `json:"quasiParabolic,omitzero"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshev `json:"shiftedChebyshev,omitzero"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst,omitzero"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap,omitzero"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfile) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

// Full set of the IRI formalism coefficients.
type IonoobservationNewBulkParamsBodyDensityProfileIri struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfileIri) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfileIri) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfileIri
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonoobservationNewBulkParamsBodyDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonoobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem `json:"parabolicItems,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfileParabolic) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfileParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfileParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonoobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius param.Opt[float64] `json:"earthRadius,omitzero"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolic) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshev) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameters of the constant-scale-height Chapman layer.
type IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst
	return param.MarshalObject(r, (*shadow)(&r))
}

// Varying scale height Chapman topside layer.
type IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyDoppler struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyDoppler) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyDoppler) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyDoppler
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyElevation struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyElevation) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyElevation) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyElevation
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyFrequency struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyFrequency) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyFrequency) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyFrequency
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyPhase struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyPhase) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyPhase) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyPhase
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyPolarization struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyPolarization) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyPolarization) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyPolarization
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyPower struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyPower) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyPower) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyPower
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyRange struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyRange) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyRange) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyRange
	return param.MarshalObject(r, (*shadow)(&r))
}

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonoobservationNewBulkParamsBodyScalerInfo struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyScalerInfo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyScalerInfo) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyScalerInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyStokes struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyStokes) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyStokes) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyStokes
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyTime struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyTime) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyTime) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyTime
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationNewBulkParamsBodyTraceGeneric struct {
	// Notes for the trace generic data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data,omitzero"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationNewBulkParamsBodyTraceGeneric) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationNewBulkParamsBodyTraceGeneric) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationNewBulkParamsBodyTraceGeneric
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [IonoobservationTupleParams]'s query parameters as
// `url.Values`.
func (r IonoobservationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonoobservationUnvalidatedPublishParams struct {
	Body []IonoobservationUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r IonoobservationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
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
type IonoobservationUnvalidatedPublishParamsBody struct {
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
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
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
	// Time the row was updated in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
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
	Amplitude              IonoobservationUnvalidatedPublishParamsBodyAmplitude              `json:"amplitude,omitzero"`
	AntennaElementPosition IonoobservationUnvalidatedPublishParamsBodyAntennaElementPosition `json:"antennaElementPosition,omitzero"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem string `json:"antennaElementPositionCoordinateSystem,omitzero"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                                            `json:"artistFlags,omitzero"`
	Azimuth     IonoobservationUnvalidatedPublishParamsBodyAzimuth `json:"azimuth,omitzero"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonoobservationUnvalidatedPublishParamsBodyCharAtt `json:"charAtts,omitzero"`
	Datum    IonoobservationUnvalidatedPublishParamsBodyDatum     `json:"datum,omitzero"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonoobservationUnvalidatedPublishParamsBodyDensityProfile `json:"densityProfile,omitzero"`
	Doppler        IonoobservationUnvalidatedPublishParamsBodyDoppler        `json:"doppler,omitzero"`
	// Array of electron densities in cm^-3 (must match the size of the height and
	// plasmaFrequency arrays).
	ElectronDensity []float64 `json:"electronDensity,omitzero"`
	// Uncertainty in specifying the electron density at each height point of the
	// profile (must match the size of the electronDensity array).
	ElectronDensityUncertainty []float64                                            `json:"electronDensityUncertainty,omitzero"`
	Elevation                  IonoobservationUnvalidatedPublishParamsBodyElevation `json:"elevation,omitzero"`
	Frequency                  IonoobservationUnvalidatedPublishParamsBodyFrequency `json:"frequency,omitzero"`
	// Array of altitudes above station level for plasma frequency/density arrays in km
	// (must match the size of the plasmaFrequency and electronDensity Arrays).
	Height []float64                                        `json:"height,omitzero"`
	Phase  IonoobservationUnvalidatedPublishParamsBodyPhase `json:"phase,omitzero"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency,omitzero"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64                                               `json:"plasmaFrequencyUncertainty,omitzero"`
	Polarization               IonoobservationUnvalidatedPublishParamsBodyPolarization `json:"polarization,omitzero"`
	Power                      IonoobservationUnvalidatedPublishParamsBodyPower        `json:"power,omitzero"`
	Range                      IonoobservationUnvalidatedPublishParamsBodyRange        `json:"range,omitzero"`
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
	ScalerInfo IonoobservationUnvalidatedPublishParamsBodyScalerInfo `json:"scalerInfo,omitzero"`
	Stokes     IonoobservationUnvalidatedPublishParamsBodyStokes     `json:"stokes,omitzero"`
	// Array of degrees clockwise from true North of the TID.
	TidAzimuth []float64 `json:"tidAzimuth,omitzero"`
	// Array of 1/frequency of the TID wave.
	TidPeriods []float64 `json:"tidPeriods,omitzero"`
	// Array of speed in m/s at which the disturbance travels through the ionosphere.
	TidPhaseSpeeds []float64                                               `json:"tidPhaseSpeeds,omitzero"`
	Time           IonoobservationUnvalidatedPublishParamsBodyTime         `json:"time,omitzero"`
	TraceGeneric   IonoobservationUnvalidatedPublishParamsBodyTraceGeneric `json:"traceGeneric,omitzero"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[IonoobservationUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[IonoobservationUnvalidatedPublishParamsBody](
		"AntennaElementPositionCoordinateSystem", false, "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)", "GEOCentric (lat, long, radii)",
	)
	apijson.RegisterFieldValidator[IonoobservationUnvalidatedPublishParamsBody](
		"ReceiveSensorType", false, "Mobile", "Static",
	)
	apijson.RegisterFieldValidator[IonoobservationUnvalidatedPublishParamsBody](
		"TransmitSensorType", false, "Mobile", "Static",
	)
}

type IonoobservationUnvalidatedPublishParamsBodyAmplitude struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyAmplitude) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyAmplitude) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyAmplitude
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyAntennaElementPosition) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyAntennaElementPosition) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyAntennaElementPosition
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyAzimuth struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyAzimuth) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyAzimuth) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyAzimuth
	return param.MarshalObject(r, (*shadow)(&r))
}

// Characteristic attributes of a IonoObservation.
type IonoobservationUnvalidatedPublishParamsBodyCharAtt struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyCharAtt) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyCharAtt) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyCharAtt
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyDatum struct {
	// Notes for the datum with details of what the data is, units, etc.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array to support sparse data collections.
	Data []float64 `json:"data,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDatum) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDatum) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDatum
	return param.MarshalObject(r, (*shadow)(&r))
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonoobservationUnvalidatedPublishParamsBodyDensityProfile struct {
	// Description of the valley model and parameters.
	ValleyModelDescription param.Opt[string] `json:"valleyModelDescription,omitzero"`
	// Full set of the IRI formalism coefficients.
	Iri IonoobservationUnvalidatedPublishParamsBodyDensityProfileIri `json:"iri,omitzero"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolic `json:"parabolic,omitzero"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic `json:"quasiParabolic,omitzero"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev `json:"shiftedChebyshev,omitzero"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst,omitzero"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap,omitzero"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfile) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

// Full set of the IRI formalism coefficients.
type IonoobservationUnvalidatedPublishParamsBodyDensityProfileIri struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfileIri) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfileIri) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfileIri
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem `json:"parabolicItems,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolic) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius param.Opt[float64] `json:"earthRadius,omitzero"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameters of the constant-scale-height Chapman layer.
type IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst
	return param.MarshalObject(r, (*shadow)(&r))
}

// Varying scale height Chapman topside layer.
type IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyDoppler struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyDoppler) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyDoppler) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyDoppler
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyElevation struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyElevation) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyElevation) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyElevation
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyFrequency struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyFrequency) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyFrequency) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyFrequency
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyPhase struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyPhase) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyPhase) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyPhase
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyPolarization struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyPolarization) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyPolarization) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyPolarization
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyPower struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyPower) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyPower) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyPower
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyRange struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyRange) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyRange) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyRange
	return param.MarshalObject(r, (*shadow)(&r))
}

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonoobservationUnvalidatedPublishParamsBodyScalerInfo struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyScalerInfo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyScalerInfo) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyScalerInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyStokes struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyStokes) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyStokes) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyStokes
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyTime struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyTime) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyTime) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyTime
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonoobservationUnvalidatedPublishParamsBodyTraceGeneric struct {
	// Notes for the trace generic data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data,omitzero"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IonoobservationUnvalidatedPublishParamsBodyTraceGeneric) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r IonoobservationUnvalidatedPublishParamsBodyTraceGeneric) MarshalJSON() (data []byte, err error) {
	type shadow IonoobservationUnvalidatedPublishParamsBodyTraceGeneric
	return param.MarshalObject(r, (*shadow)(&r))
}
