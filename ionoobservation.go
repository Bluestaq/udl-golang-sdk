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

// IonOobservationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIonOobservationService] method instead.
type IonOobservationService struct {
	Options []option.RequestOption
}

// NewIonOobservationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIonOobservationService(opts ...option.RequestOption) (r IonOobservationService) {
	r = IonOobservationService{}
	r.Options = opts
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *IonOobservationService) List(ctx context.Context, query IonOobservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[IonOobservationListResponse], err error) {
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
func (r *IonOobservationService) ListAutoPaging(ctx context.Context, query IonOobservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[IonOobservationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *IonOobservationService) Count(ctx context.Context, query IonOobservationCountParams, opts ...option.RequestOption) (res *string, err error) {
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
func (r *IonOobservationService) NewBulk(ctx context.Context, body IonOobservationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/ionoobservation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *IonOobservationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
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
func (r *IonOobservationService) Tuple(ctx context.Context, query IonOobservationTupleParams, opts ...option.RequestOption) (res *[]IonOobservationTupleResponse, err error) {
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
func (r *IonOobservationService) UnvalidatedPublish(ctx context.Context, body IonOobservationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
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
type IonOobservationListResponse struct {
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
	DataMode IonOobservationListResponseDataMode `json:"dataMode,required"`
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
	Amplitude              IonOobservationListResponseAmplitude              `json:"amplitude"`
	AntennaElementPosition IonOobservationListResponseAntennaElementPosition `json:"antennaElementPosition"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem IonOobservationListResponseAntennaElementPositionCoordinateSystem `json:"antennaElementPositionCoordinateSystem"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                            `json:"artistFlags"`
	Azimuth     IonOobservationListResponseAzimuth `json:"azimuth"`
	// IRI thickness parameter in km. URSI ID: D0.
	B0 float64 `json:"b0"`
	// IRI profile shape parameter. URSI ID: D1.
	B1 float64 `json:"b1"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonOobservationListResponseCharAtt `json:"charAtts"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Distance for MUF calculation in km.
	D float64 `json:"d"`
	// IRI profile shape parameter, F1 layer. URSI ID: D2.
	D1    float64                          `json:"d1"`
	Datum IonOobservationListResponseDatum `json:"datum"`
	// Adjustment to the scaled foF2 during profile inversion in MHz.
	DeltafoF2 float64 `json:"deltafoF2"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonOobservationListResponseDensityProfile `json:"densityProfile"`
	Doppler        IonOobservationListResponseDoppler        `json:"doppler"`
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
	Elevation                  IonOobservationListResponseElevation `json:"elevation"`
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
	Frequency IonOobservationListResponseFrequency `json:"frequency"`
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
	Phase        IonOobservationListResponsePhase `json:"phase"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64 `json:"plasmaFrequencyUncertainty"`
	// Equipment location.
	PlatformName string                                  `json:"platformName"`
	Polarization IonOobservationListResponsePolarization `json:"polarization"`
	Power        IonOobservationListResponsePower        `json:"power"`
	// Average range spread of E layer in km. URSI ID: 85.
	Qe float64 `json:"qe"`
	// Average range spread of F layer in km. URSI ID: 84.
	Qf    float64                          `json:"qf"`
	Range IonOobservationListResponseRange `json:"range"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	ReceiveCoordinates [][]float64 `json:"receiveCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	ReceiveSensorType IonOobservationListResponseReceiveSensorType `json:"receiveSensorType"`
	// Array of restricted frequencies.
	RestrictedFrequency []float64 `json:"restrictedFrequency"`
	// Notes for the restrictedFrequency data.
	RestrictedFrequencyNotes string `json:"restrictedFrequencyNotes"`
	// Effective scale height at hmF2 Titheridge method in km. URSI ID: 69.
	ScaleHeightF2Peak float64 `json:"scaleHeightF2Peak"`
	// The ScalerInfo record describes the person or system who interpreted the
	// ionogram in IonoObservation.
	ScalerInfo IonOobservationListResponseScalerInfo `json:"scalerInfo"`
	Stokes     IonOobservationListResponseStokes     `json:"stokes"`
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
	Time           IonOobservationListResponseTime         `json:"time"`
	TraceGeneric   IonOobservationListResponseTraceGeneric `json:"traceGeneric"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	TransmitCoordinates [][]float64 `json:"transmitCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	TransmitSensorType IonOobservationListResponseTransmitSensorType `json:"transmitSensorType"`
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponse) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponse) UnmarshalJSON(data []byte) error {
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
type IonOobservationListResponseDataMode string

const (
	IonOobservationListResponseDataModeReal      IonOobservationListResponseDataMode = "REAL"
	IonOobservationListResponseDataModeTest      IonOobservationListResponseDataMode = "TEST"
	IonOobservationListResponseDataModeSimulated IonOobservationListResponseDataMode = "SIMULATED"
	IonOobservationListResponseDataModeExercise  IonOobservationListResponseDataMode = "EXERCISE"
)

type IonOobservationListResponseAmplitude struct {
	// Array of amplitude data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for amplitude dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the amplitude data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseAmplitude) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseAmplitude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponseAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationListResponseAntennaElementPosition) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseAntennaElementPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
// (lat, long, radii).
type IonOobservationListResponseAntennaElementPositionCoordinateSystem string

const (
	IonOobservationListResponseAntennaElementPositionCoordinateSystemJ2000                   IonOobservationListResponseAntennaElementPositionCoordinateSystem = "J2000"
	IonOobservationListResponseAntennaElementPositionCoordinateSystemEcrEcef                 IonOobservationListResponseAntennaElementPositionCoordinateSystem = "ECR/ECEF"
	IonOobservationListResponseAntennaElementPositionCoordinateSystemTeme                    IonOobservationListResponseAntennaElementPositionCoordinateSystem = "TEME"
	IonOobservationListResponseAntennaElementPositionCoordinateSystemGcrf                    IonOobservationListResponseAntennaElementPositionCoordinateSystem = "GCRF"
	IonOobservationListResponseAntennaElementPositionCoordinateSystemWgs84GeoDeticLatLongAlt IonOobservationListResponseAntennaElementPositionCoordinateSystem = "WGS84 (GEODetic lat, long, alt)"
	IonOobservationListResponseAntennaElementPositionCoordinateSystemGeoCentricLatLongRadii  IonOobservationListResponseAntennaElementPositionCoordinateSystem = "GEOCentric (lat, long, radii)"
)

type IonOobservationListResponseAzimuth struct {
	// Array of incoming azimuth at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the azimuth array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the azimuth data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseAzimuth) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseAzimuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Characteristic attributes of a IonoObservation.
type IonOobservationListResponseCharAtt struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseCharAtt) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseCharAtt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponseDatum struct {
	// Array to support sparse data collections.
	Data []float64 `json:"data"`
	// Notes for the datum with details of what the data is, units, etc.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Data        resp.Field
		Notes       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationListResponseDatum) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseDatum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonOobservationListResponseDensityProfile struct {
	// Full set of the IRI formalism coefficients.
	Iri IonOobservationListResponseDensityProfileIri `json:"iri"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonOobservationListResponseDensityProfileParabolic `json:"parabolic"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonOobservationListResponseDensityProfileQuasiParabolic `json:"quasiParabolic"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonOobservationListResponseDensityProfileShiftedChebyshev `json:"shiftedChebyshev"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonOobservationListResponseDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonOobservationListResponseDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs"`
	// Description of the valley model and parameters.
	ValleyModelDescription string `json:"valleyModelDescription"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseDensityProfile) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseDensityProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full set of the IRI formalism coefficients.
type IonOobservationListResponseDensityProfileIri struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseDensityProfileIri) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseDensityProfileIri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonOobservationListResponseDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description string `json:"description"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonOobservationListResponseDensityProfileParabolicParabolicItem `json:"parabolicItems"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Description    resp.Field
		ParabolicItems resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationListResponseDensityProfileParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseDensityProfileParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonOobservationListResponseDensityProfileParabolicParabolicItem struct {
	// Plasma frequency at the layer peak, in MHz.
	F float64 `json:"f"`
	// Ionospheric plasma layer (E, F1, or F2).
	Layer string `json:"layer"`
	// Half-thickness of the layer, in kilometers.
	Y float64 `json:"y"`
	// Height of the layer peak, in kilometers.
	Z float64 `json:"z"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseDensityProfileParabolicParabolicItem) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationListResponseDensityProfileParabolicParabolicItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonOobservationListResponseDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description string `json:"description"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius float64 `json:"earthRadius"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonOobservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Description            resp.Field
		EarthRadius            resp.Field
		QuasiParabolicSegments resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationListResponseDensityProfileQuasiParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseDensityProfileQuasiParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonOobservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationListResponseDensityProfileQuasiParabolicQuasiParabolicSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonOobservationListResponseDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description string `json:"description"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonOobservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Description       resp.Field
		ShiftedChebyshevs resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationListResponseDensityProfileShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationListResponseDensityProfileShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonOobservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationListResponseDensityProfileShiftedChebyshevShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters of the constant-scale-height Chapman layer.
type IonOobservationListResponseDensityProfileTopsideExtensionChapmanConst struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseDensityProfileTopsideExtensionChapmanConst) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationListResponseDensityProfileTopsideExtensionChapmanConst) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Varying scale height Chapman topside layer.
type IonOobservationListResponseDensityProfileTopsideExtensionVaryChap struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseDensityProfileTopsideExtensionVaryChap) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationListResponseDensityProfileTopsideExtensionVaryChap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponseDoppler struct {
	// Array of received doppler shifts in Hz.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the doppler array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the doppler data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseDoppler) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseDoppler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponseElevation struct {
	// Array of incoming elevation at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the elevation array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the elevation data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseElevation) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseElevation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponseFrequency struct {
	// Array of frequency data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for frequency dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the frequency array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the frequency data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseFrequency) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseFrequency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponsePhase struct {
	// Array of phase data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for phase dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the phase array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the phase data. Orientation and position for each antenna element
	// across the antenna_element dimension.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponsePhase) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponsePhase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponsePolarization struct {
	// Array of polarization data.
	Data [][][][][][][]string `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for polarization dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the polarization data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponsePolarization) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponsePolarization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponsePower struct {
	// Array of received power in db.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the power array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the power data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponsePower) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponsePower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponseRange struct {
	// Array of range data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for range dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the range array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the range data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseRange) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonOobservationListResponseReceiveSensorType string

const (
	IonOobservationListResponseReceiveSensorTypeMobile IonOobservationListResponseReceiveSensorType = "Mobile"
	IonOobservationListResponseReceiveSensorTypeStatic IonOobservationListResponseReceiveSensorType = "Static"
)

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonOobservationListResponseScalerInfo struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseScalerInfo) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseScalerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponseStokes struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseStokes) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseStokes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponseTime struct {
	// Array of times in number of seconds passed since January 1st, 1970 with the same
	// dimensions as power.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the time array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// The notes indicate the scheme and accuracy.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationListResponseTime) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationListResponseTraceGeneric struct {
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName"`
	// Notes for the trace generic data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationListResponseTraceGeneric) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationListResponseTraceGeneric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonOobservationListResponseTransmitSensorType string

const (
	IonOobservationListResponseTransmitSensorTypeMobile IonOobservationListResponseTransmitSensorType = "Mobile"
	IonOobservationListResponseTransmitSensorTypeStatic IonOobservationListResponseTransmitSensorType = "Static"
)

// These services provide operations for posting and querying ionospheric
// observation data. Characteristics are defined by the CHARS: URSI IIWG format for
// archiving monthly ionospheric characteristics, INAG Bulletin No. 62
// specification. Qualifying and Descriptive letters are defined by the URSI
// Handbook for Ionogram Interpretation and reduction, Report UAG, No. 23A
// specification.
type IonOobservationTupleResponse struct {
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
	DataMode IonOobservationTupleResponseDataMode `json:"dataMode,required"`
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
	Amplitude              IonOobservationTupleResponseAmplitude              `json:"amplitude"`
	AntennaElementPosition IonOobservationTupleResponseAntennaElementPosition `json:"antennaElementPosition"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem IonOobservationTupleResponseAntennaElementPositionCoordinateSystem `json:"antennaElementPositionCoordinateSystem"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                             `json:"artistFlags"`
	Azimuth     IonOobservationTupleResponseAzimuth `json:"azimuth"`
	// IRI thickness parameter in km. URSI ID: D0.
	B0 float64 `json:"b0"`
	// IRI profile shape parameter. URSI ID: D1.
	B1 float64 `json:"b1"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonOobservationTupleResponseCharAtt `json:"charAtts"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Distance for MUF calculation in km.
	D float64 `json:"d"`
	// IRI profile shape parameter, F1 layer. URSI ID: D2.
	D1    float64                           `json:"d1"`
	Datum IonOobservationTupleResponseDatum `json:"datum"`
	// Adjustment to the scaled foF2 during profile inversion in MHz.
	DeltafoF2 float64 `json:"deltafoF2"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonOobservationTupleResponseDensityProfile `json:"densityProfile"`
	Doppler        IonOobservationTupleResponseDoppler        `json:"doppler"`
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
	Elevation                  IonOobservationTupleResponseElevation `json:"elevation"`
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
	Frequency IonOobservationTupleResponseFrequency `json:"frequency"`
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
	Phase        IonOobservationTupleResponsePhase `json:"phase"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64 `json:"plasmaFrequencyUncertainty"`
	// Equipment location.
	PlatformName string                                   `json:"platformName"`
	Polarization IonOobservationTupleResponsePolarization `json:"polarization"`
	Power        IonOobservationTupleResponsePower        `json:"power"`
	// Average range spread of E layer in km. URSI ID: 85.
	Qe float64 `json:"qe"`
	// Average range spread of F layer in km. URSI ID: 84.
	Qf    float64                           `json:"qf"`
	Range IonOobservationTupleResponseRange `json:"range"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	ReceiveCoordinates [][]float64 `json:"receiveCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	ReceiveSensorType IonOobservationTupleResponseReceiveSensorType `json:"receiveSensorType"`
	// Array of restricted frequencies.
	RestrictedFrequency []float64 `json:"restrictedFrequency"`
	// Notes for the restrictedFrequency data.
	RestrictedFrequencyNotes string `json:"restrictedFrequencyNotes"`
	// Effective scale height at hmF2 Titheridge method in km. URSI ID: 69.
	ScaleHeightF2Peak float64 `json:"scaleHeightF2Peak"`
	// The ScalerInfo record describes the person or system who interpreted the
	// ionogram in IonoObservation.
	ScalerInfo IonOobservationTupleResponseScalerInfo `json:"scalerInfo"`
	Stokes     IonOobservationTupleResponseStokes     `json:"stokes"`
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
	Time           IonOobservationTupleResponseTime         `json:"time"`
	TraceGeneric   IonOobservationTupleResponseTraceGeneric `json:"traceGeneric"`
	// List of Geodetic Latitude, Longitude, and Altitude coordinates in km of the
	// receiver. Coordinates List must always have 3 elements. Valid ranges are -90.0
	// to 90.0 for Latitude and -180.0 to 180.0 for Longitude. Altitude is in km and
	// its value must be 0 or greater.
	TransmitCoordinates [][]float64 `json:"transmitCoordinates"`
	// Enums: Mobile, Static.
	//
	// Any of "Mobile", "Static".
	TransmitSensorType IonOobservationTupleResponseTransmitSensorType `json:"transmitSensorType"`
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponse) UnmarshalJSON(data []byte) error {
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
type IonOobservationTupleResponseDataMode string

const (
	IonOobservationTupleResponseDataModeReal      IonOobservationTupleResponseDataMode = "REAL"
	IonOobservationTupleResponseDataModeTest      IonOobservationTupleResponseDataMode = "TEST"
	IonOobservationTupleResponseDataModeSimulated IonOobservationTupleResponseDataMode = "SIMULATED"
	IonOobservationTupleResponseDataModeExercise  IonOobservationTupleResponseDataMode = "EXERCISE"
)

type IonOobservationTupleResponseAmplitude struct {
	// Array of amplitude data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for amplitude dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the amplitude data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseAmplitude) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseAmplitude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponseAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Dimensions    resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationTupleResponseAntennaElementPosition) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseAntennaElementPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
// (lat, long, radii).
type IonOobservationTupleResponseAntennaElementPositionCoordinateSystem string

const (
	IonOobservationTupleResponseAntennaElementPositionCoordinateSystemJ2000                   IonOobservationTupleResponseAntennaElementPositionCoordinateSystem = "J2000"
	IonOobservationTupleResponseAntennaElementPositionCoordinateSystemEcrEcef                 IonOobservationTupleResponseAntennaElementPositionCoordinateSystem = "ECR/ECEF"
	IonOobservationTupleResponseAntennaElementPositionCoordinateSystemTeme                    IonOobservationTupleResponseAntennaElementPositionCoordinateSystem = "TEME"
	IonOobservationTupleResponseAntennaElementPositionCoordinateSystemGcrf                    IonOobservationTupleResponseAntennaElementPositionCoordinateSystem = "GCRF"
	IonOobservationTupleResponseAntennaElementPositionCoordinateSystemWgs84GeoDeticLatLongAlt IonOobservationTupleResponseAntennaElementPositionCoordinateSystem = "WGS84 (GEODetic lat, long, alt)"
	IonOobservationTupleResponseAntennaElementPositionCoordinateSystemGeoCentricLatLongRadii  IonOobservationTupleResponseAntennaElementPositionCoordinateSystem = "GEOCentric (lat, long, radii)"
)

type IonOobservationTupleResponseAzimuth struct {
	// Array of incoming azimuth at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the azimuth array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the azimuth data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseAzimuth) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseAzimuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Characteristic attributes of a IonoObservation.
type IonOobservationTupleResponseCharAtt struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseCharAtt) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseCharAtt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponseDatum struct {
	// Array to support sparse data collections.
	Data []float64 `json:"data"`
	// Notes for the datum with details of what the data is, units, etc.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Data        resp.Field
		Notes       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationTupleResponseDatum) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseDatum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonOobservationTupleResponseDensityProfile struct {
	// Full set of the IRI formalism coefficients.
	Iri IonOobservationTupleResponseDensityProfileIri `json:"iri"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonOobservationTupleResponseDensityProfileParabolic `json:"parabolic"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonOobservationTupleResponseDensityProfileQuasiParabolic `json:"quasiParabolic"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonOobservationTupleResponseDensityProfileShiftedChebyshev `json:"shiftedChebyshev"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonOobservationTupleResponseDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonOobservationTupleResponseDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs"`
	// Description of the valley model and parameters.
	ValleyModelDescription string `json:"valleyModelDescription"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseDensityProfile) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseDensityProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full set of the IRI formalism coefficients.
type IonOobservationTupleResponseDensityProfileIri struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseDensityProfileIri) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseDensityProfileIri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonOobservationTupleResponseDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description string `json:"description"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonOobservationTupleResponseDensityProfileParabolicParabolicItem `json:"parabolicItems"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Description    resp.Field
		ParabolicItems resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationTupleResponseDensityProfileParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseDensityProfileParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonOobservationTupleResponseDensityProfileParabolicParabolicItem struct {
	// Plasma frequency at the layer peak, in MHz.
	F float64 `json:"f"`
	// Ionospheric plasma layer (E, F1, or F2).
	Layer string `json:"layer"`
	// Half-thickness of the layer, in kilometers.
	Y float64 `json:"y"`
	// Height of the layer peak, in kilometers.
	Z float64 `json:"z"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseDensityProfileParabolicParabolicItem) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationTupleResponseDensityProfileParabolicParabolicItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonOobservationTupleResponseDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description string `json:"description"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius float64 `json:"earthRadius"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonOobservationTupleResponseDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Description            resp.Field
		EarthRadius            resp.Field
		QuasiParabolicSegments resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationTupleResponseDensityProfileQuasiParabolic) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseDensityProfileQuasiParabolic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonOobservationTupleResponseDensityProfileQuasiParabolicQuasiParabolicSegment struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseDensityProfileQuasiParabolicQuasiParabolicSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationTupleResponseDensityProfileQuasiParabolicQuasiParabolicSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonOobservationTupleResponseDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description string `json:"description"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonOobservationTupleResponseDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Description       resp.Field
		ShiftedChebyshevs resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationTupleResponseDensityProfileShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationTupleResponseDensityProfileShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonOobservationTupleResponseDensityProfileShiftedChebyshevShiftedChebyshev struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseDensityProfileShiftedChebyshevShiftedChebyshev) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationTupleResponseDensityProfileShiftedChebyshevShiftedChebyshev) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters of the constant-scale-height Chapman layer.
type IonOobservationTupleResponseDensityProfileTopsideExtensionChapmanConst struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseDensityProfileTopsideExtensionChapmanConst) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationTupleResponseDensityProfileTopsideExtensionChapmanConst) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Varying scale height Chapman topside layer.
type IonOobservationTupleResponseDensityProfileTopsideExtensionVaryChap struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseDensityProfileTopsideExtensionVaryChap) RawJSON() string {
	return r.JSON.raw
}
func (r *IonOobservationTupleResponseDensityProfileTopsideExtensionVaryChap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponseDoppler struct {
	// Array of received doppler shifts in Hz.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the doppler array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the doppler data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseDoppler) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseDoppler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponseElevation struct {
	// Array of incoming elevation at the receiver.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the elevation array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the elevation data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseElevation) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseElevation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponseFrequency struct {
	// Array of frequency data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for frequency dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the frequency array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the frequency data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseFrequency) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseFrequency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponsePhase struct {
	// Array of phase data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for phase dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the phase array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the phase data. Orientation and position for each antenna element
	// across the antenna_element dimension.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponsePhase) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponsePhase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponsePolarization struct {
	// Array of polarization data.
	Data [][][][][][][]string `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers for polarization dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the polarization data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponsePolarization) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponsePolarization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponsePower struct {
	// Array of received power in db.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the power array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the power data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponsePower) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponsePower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponseRange struct {
	// Array of range data.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for range dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the range array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// Notes for the range data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseRange) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonOobservationTupleResponseReceiveSensorType string

const (
	IonOobservationTupleResponseReceiveSensorTypeMobile IonOobservationTupleResponseReceiveSensorType = "Mobile"
	IonOobservationTupleResponseReceiveSensorTypeStatic IonOobservationTupleResponseReceiveSensorType = "Static"
)

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonOobservationTupleResponseScalerInfo struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseScalerInfo) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseScalerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponseStokes struct {
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
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseStokes) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseStokes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponseTime struct {
	// Array of times in number of seconds passed since January 1st, 1970 with the same
	// dimensions as power.
	Data [][][][][][][]float64 `json:"data"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName"`
	// Array of integers of the time array dimensions.
	Dimensions []int64 `json:"dimensions"`
	// The notes indicate the scheme and accuracy.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
func (r IonOobservationTupleResponseTime) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IonOobservationTupleResponseTraceGeneric struct {
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName"`
	// Notes for the trace generic data.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Data          resp.Field
		DimensionName resp.Field
		Notes         resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IonOobservationTupleResponseTraceGeneric) RawJSON() string { return r.JSON.raw }
func (r *IonOobservationTupleResponseTraceGeneric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enums: Mobile, Static.
type IonOobservationTupleResponseTransmitSensorType string

const (
	IonOobservationTupleResponseTransmitSensorTypeMobile IonOobservationTupleResponseTransmitSensorType = "Mobile"
	IonOobservationTupleResponseTransmitSensorTypeStatic IonOobservationTupleResponseTransmitSensorType = "Static"
)

type IonOobservationListParams struct {
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time        `query:"startTimeUTC,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IonOobservationListParams]'s query parameters as
// `url.Values`.
func (r IonOobservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonOobservationCountParams struct {
	// Sounding Start time in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTimeUtc time.Time        `query:"startTimeUTC,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IonOobservationCountParams]'s query parameters as
// `url.Values`.
func (r IonOobservationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonOobservationNewBulkParams struct {
	Body []IonOobservationNewBulkParamsBody
	paramObj
}

func (r IonOobservationNewBulkParams) MarshalJSON() (data []byte, err error) {
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
type IonOobservationNewBulkParamsBody struct {
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
	Amplitude              IonOobservationNewBulkParamsBodyAmplitude              `json:"amplitude,omitzero"`
	AntennaElementPosition IonOobservationNewBulkParamsBodyAntennaElementPosition `json:"antennaElementPosition,omitzero"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem string `json:"antennaElementPositionCoordinateSystem,omitzero"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                                 `json:"artistFlags,omitzero"`
	Azimuth     IonOobservationNewBulkParamsBodyAzimuth `json:"azimuth,omitzero"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonOobservationNewBulkParamsBodyCharAtt `json:"charAtts,omitzero"`
	Datum    IonOobservationNewBulkParamsBodyDatum     `json:"datum,omitzero"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonOobservationNewBulkParamsBodyDensityProfile `json:"densityProfile,omitzero"`
	Doppler        IonOobservationNewBulkParamsBodyDoppler        `json:"doppler,omitzero"`
	// Array of electron densities in cm^-3 (must match the size of the height and
	// plasmaFrequency arrays).
	ElectronDensity []float64 `json:"electronDensity,omitzero"`
	// Uncertainty in specifying the electron density at each height point of the
	// profile (must match the size of the electronDensity array).
	ElectronDensityUncertainty []float64                                 `json:"electronDensityUncertainty,omitzero"`
	Elevation                  IonOobservationNewBulkParamsBodyElevation `json:"elevation,omitzero"`
	Frequency                  IonOobservationNewBulkParamsBodyFrequency `json:"frequency,omitzero"`
	// Array of altitudes above station level for plasma frequency/density arrays in km
	// (must match the size of the plasmaFrequency and electronDensity Arrays).
	Height []float64                             `json:"height,omitzero"`
	Phase  IonOobservationNewBulkParamsBodyPhase `json:"phase,omitzero"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency,omitzero"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64                                    `json:"plasmaFrequencyUncertainty,omitzero"`
	Polarization               IonOobservationNewBulkParamsBodyPolarization `json:"polarization,omitzero"`
	Power                      IonOobservationNewBulkParamsBodyPower        `json:"power,omitzero"`
	Range                      IonOobservationNewBulkParamsBodyRange        `json:"range,omitzero"`
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
	ScalerInfo IonOobservationNewBulkParamsBodyScalerInfo `json:"scalerInfo,omitzero"`
	Stokes     IonOobservationNewBulkParamsBodyStokes     `json:"stokes,omitzero"`
	// Array of degrees clockwise from true North of the TID.
	TidAzimuth []float64 `json:"tidAzimuth,omitzero"`
	// Array of 1/frequency of the TID wave.
	TidPeriods []float64 `json:"tidPeriods,omitzero"`
	// Array of speed in m/s at which the disturbance travels through the ionosphere.
	TidPhaseSpeeds []float64                                    `json:"tidPhaseSpeeds,omitzero"`
	Time           IonOobservationNewBulkParamsBodyTime         `json:"time,omitzero"`
	TraceGeneric   IonOobservationNewBulkParamsBodyTraceGeneric `json:"traceGeneric,omitzero"`
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

func (r IonOobservationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[IonOobservationNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[IonOobservationNewBulkParamsBody](
		"AntennaElementPositionCoordinateSystem", false, "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)", "GEOCentric (lat, long, radii)",
	)
	apijson.RegisterFieldValidator[IonOobservationNewBulkParamsBody](
		"ReceiveSensorType", false, "Mobile", "Static",
	)
	apijson.RegisterFieldValidator[IonOobservationNewBulkParamsBody](
		"TransmitSensorType", false, "Mobile", "Static",
	)
}

type IonOobservationNewBulkParamsBodyAmplitude struct {
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

func (r IonOobservationNewBulkParamsBodyAmplitude) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyAmplitude
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonOobservationNewBulkParamsBodyAntennaElementPosition) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyAntennaElementPosition
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyAzimuth struct {
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

func (r IonOobservationNewBulkParamsBodyAzimuth) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyAzimuth
	return param.MarshalObject(r, (*shadow)(&r))
}

// Characteristic attributes of a IonoObservation.
type IonOobservationNewBulkParamsBodyCharAtt struct {
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

func (r IonOobservationNewBulkParamsBodyCharAtt) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyCharAtt
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyDatum struct {
	// Notes for the datum with details of what the data is, units, etc.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array to support sparse data collections.
	Data []float64 `json:"data,omitzero"`
	paramObj
}

func (r IonOobservationNewBulkParamsBodyDatum) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDatum
	return param.MarshalObject(r, (*shadow)(&r))
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonOobservationNewBulkParamsBodyDensityProfile struct {
	// Description of the valley model and parameters.
	ValleyModelDescription param.Opt[string] `json:"valleyModelDescription,omitzero"`
	// Full set of the IRI formalism coefficients.
	Iri IonOobservationNewBulkParamsBodyDensityProfileIri `json:"iri,omitzero"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonOobservationNewBulkParamsBodyDensityProfileParabolic `json:"parabolic,omitzero"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolic `json:"quasiParabolic,omitzero"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshev `json:"shiftedChebyshev,omitzero"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst,omitzero"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap,omitzero"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs,omitzero"`
	paramObj
}

func (r IonOobservationNewBulkParamsBodyDensityProfile) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

// Full set of the IRI formalism coefficients.
type IonOobservationNewBulkParamsBodyDensityProfileIri struct {
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

func (r IonOobservationNewBulkParamsBodyDensityProfileIri) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfileIri
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonOobservationNewBulkParamsBodyDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonOobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem `json:"parabolicItems,omitzero"`
	paramObj
}

func (r IonOobservationNewBulkParamsBodyDensityProfileParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfileParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonOobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem struct {
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

func (r IonOobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius param.Opt[float64] `json:"earthRadius,omitzero"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments,omitzero"`
	paramObj
}

func (r IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment struct {
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

func (r IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs,omitzero"`
	paramObj
}

func (r IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev struct {
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

func (r IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameters of the constant-scale-height Chapman layer.
type IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst struct {
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

func (r IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst
	return param.MarshalObject(r, (*shadow)(&r))
}

// Varying scale height Chapman topside layer.
type IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap struct {
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

func (r IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyDoppler struct {
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

func (r IonOobservationNewBulkParamsBodyDoppler) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyDoppler
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyElevation struct {
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

func (r IonOobservationNewBulkParamsBodyElevation) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyElevation
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyFrequency struct {
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

func (r IonOobservationNewBulkParamsBodyFrequency) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyFrequency
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyPhase struct {
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

func (r IonOobservationNewBulkParamsBodyPhase) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyPhase
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyPolarization struct {
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

func (r IonOobservationNewBulkParamsBodyPolarization) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyPolarization
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyPower struct {
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

func (r IonOobservationNewBulkParamsBodyPower) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyPower
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyRange struct {
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

func (r IonOobservationNewBulkParamsBodyRange) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyRange
	return param.MarshalObject(r, (*shadow)(&r))
}

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonOobservationNewBulkParamsBodyScalerInfo struct {
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

func (r IonOobservationNewBulkParamsBodyScalerInfo) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyScalerInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyStokes struct {
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

func (r IonOobservationNewBulkParamsBodyStokes) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyStokes
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyTime struct {
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

func (r IonOobservationNewBulkParamsBodyTime) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyTime
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationNewBulkParamsBodyTraceGeneric struct {
	// Notes for the trace generic data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data,omitzero"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName,omitzero"`
	paramObj
}

func (r IonOobservationNewBulkParamsBodyTraceGeneric) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationNewBulkParamsBodyTraceGeneric
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationTupleParams struct {
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

// URLQuery serializes [IonOobservationTupleParams]'s query parameters as
// `url.Values`.
func (r IonOobservationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IonOobservationUnvalidatedPublishParams struct {
	Body []IonOobservationUnvalidatedPublishParamsBody
	paramObj
}

func (r IonOobservationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
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
type IonOobservationUnvalidatedPublishParamsBody struct {
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
	Amplitude              IonOobservationUnvalidatedPublishParamsBodyAmplitude              `json:"amplitude,omitzero"`
	AntennaElementPosition IonOobservationUnvalidatedPublishParamsBodyAntennaElementPosition `json:"antennaElementPosition,omitzero"`
	// Enums: J2000, ECR/ECEF, TEME, GCRF, WGS84 (GEODetic lat, long, alt), GEOCentric
	// (lat, long, radii).
	//
	// Any of "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)",
	// "GEOCentric (lat, long, radii)".
	AntennaElementPositionCoordinateSystem string `json:"antennaElementPositionCoordinateSystem,omitzero"`
	// Array of Legacy Artist Flags.
	ArtistFlags []int64                                            `json:"artistFlags,omitzero"`
	Azimuth     IonOobservationUnvalidatedPublishParamsBodyAzimuth `json:"azimuth,omitzero"`
	// List of attributes that are associated with the specified characteristics.
	// Characteristics are defined by the CHARS: URSI IIWG format for archiving monthly
	// ionospheric characteristics, INAG Bulletin No. 62 specification. Qualifying and
	// Descriptive letters are defined by the URSI Handbook for Ionogram Interpretation
	// and reduction, Report UAG, No. 23A specification.
	CharAtts []IonOobservationUnvalidatedPublishParamsBodyCharAtt `json:"charAtts,omitzero"`
	Datum    IonOobservationUnvalidatedPublishParamsBodyDatum     `json:"datum,omitzero"`
	// Profile of electron densities in the ionosphere associated with an
	// IonoObservation.
	DensityProfile IonOobservationUnvalidatedPublishParamsBodyDensityProfile `json:"densityProfile,omitzero"`
	Doppler        IonOobservationUnvalidatedPublishParamsBodyDoppler        `json:"doppler,omitzero"`
	// Array of electron densities in cm^-3 (must match the size of the height and
	// plasmaFrequency arrays).
	ElectronDensity []float64 `json:"electronDensity,omitzero"`
	// Uncertainty in specifying the electron density at each height point of the
	// profile (must match the size of the electronDensity array).
	ElectronDensityUncertainty []float64                                            `json:"electronDensityUncertainty,omitzero"`
	Elevation                  IonOobservationUnvalidatedPublishParamsBodyElevation `json:"elevation,omitzero"`
	Frequency                  IonOobservationUnvalidatedPublishParamsBodyFrequency `json:"frequency,omitzero"`
	// Array of altitudes above station level for plasma frequency/density arrays in km
	// (must match the size of the plasmaFrequency and electronDensity Arrays).
	Height []float64                                        `json:"height,omitzero"`
	Phase  IonOobservationUnvalidatedPublishParamsBodyPhase `json:"phase,omitzero"`
	// Array of plasma frequencies in MHz (must match the size of the height and
	// electronDensity arrays).
	PlasmaFrequency []float64 `json:"plasmaFrequency,omitzero"`
	// Uncertainty in specifying the electron plasma frequency at each height point of
	// the profile (must match the size of the plasmaFrequency array).
	PlasmaFrequencyUncertainty []float64                                               `json:"plasmaFrequencyUncertainty,omitzero"`
	Polarization               IonOobservationUnvalidatedPublishParamsBodyPolarization `json:"polarization,omitzero"`
	Power                      IonOobservationUnvalidatedPublishParamsBodyPower        `json:"power,omitzero"`
	Range                      IonOobservationUnvalidatedPublishParamsBodyRange        `json:"range,omitzero"`
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
	ScalerInfo IonOobservationUnvalidatedPublishParamsBodyScalerInfo `json:"scalerInfo,omitzero"`
	Stokes     IonOobservationUnvalidatedPublishParamsBodyStokes     `json:"stokes,omitzero"`
	// Array of degrees clockwise from true North of the TID.
	TidAzimuth []float64 `json:"tidAzimuth,omitzero"`
	// Array of 1/frequency of the TID wave.
	TidPeriods []float64 `json:"tidPeriods,omitzero"`
	// Array of speed in m/s at which the disturbance travels through the ionosphere.
	TidPhaseSpeeds []float64                                               `json:"tidPhaseSpeeds,omitzero"`
	Time           IonOobservationUnvalidatedPublishParamsBodyTime         `json:"time,omitzero"`
	TraceGeneric   IonOobservationUnvalidatedPublishParamsBodyTraceGeneric `json:"traceGeneric,omitzero"`
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

func (r IonOobservationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[IonOobservationUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[IonOobservationUnvalidatedPublishParamsBody](
		"AntennaElementPositionCoordinateSystem", false, "J2000", "ECR/ECEF", "TEME", "GCRF", "WGS84 (GEODetic lat, long, alt)", "GEOCentric (lat, long, radii)",
	)
	apijson.RegisterFieldValidator[IonOobservationUnvalidatedPublishParamsBody](
		"ReceiveSensorType", false, "Mobile", "Static",
	)
	apijson.RegisterFieldValidator[IonOobservationUnvalidatedPublishParamsBody](
		"TransmitSensorType", false, "Mobile", "Static",
	)
}

type IonOobservationUnvalidatedPublishParamsBodyAmplitude struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyAmplitude) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyAmplitude
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyAntennaElementPosition struct {
	// Array of 3-element tuples (x,y,z) in km.
	Data [][]float64 `json:"data,omitzero"`
	// Array of names for dimensions.
	DimensionName []string `json:"dimensionName,omitzero"`
	// Array of integers of the antenna_element dimensions.
	Dimensions []int64 `json:"dimensions,omitzero"`
	paramObj
}

func (r IonOobservationUnvalidatedPublishParamsBodyAntennaElementPosition) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyAntennaElementPosition
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyAzimuth struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyAzimuth) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyAzimuth
	return param.MarshalObject(r, (*shadow)(&r))
}

// Characteristic attributes of a IonoObservation.
type IonOobservationUnvalidatedPublishParamsBodyCharAtt struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyCharAtt) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyCharAtt
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyDatum struct {
	// Notes for the datum with details of what the data is, units, etc.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Array to support sparse data collections.
	Data []float64 `json:"data,omitzero"`
	paramObj
}

func (r IonOobservationUnvalidatedPublishParamsBodyDatum) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDatum
	return param.MarshalObject(r, (*shadow)(&r))
}

// Profile of electron densities in the ionosphere associated with an
// IonoObservation.
type IonOobservationUnvalidatedPublishParamsBodyDensityProfile struct {
	// Description of the valley model and parameters.
	ValleyModelDescription param.Opt[string] `json:"valleyModelDescription,omitzero"`
	// Full set of the IRI formalism coefficients.
	Iri IonOobservationUnvalidatedPublishParamsBodyDensityProfileIri `json:"iri,omitzero"`
	// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
	Parabolic IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolic `json:"parabolic,omitzero"`
	// Coefficients to describe profile shape as quasi-parabolic segments.
	QuasiParabolic IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic `json:"quasiParabolic,omitzero"`
	// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
	// the height uncertainty boundaries.
	ShiftedChebyshev IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev `json:"shiftedChebyshev,omitzero"`
	// Parameters of the constant-scale-height Chapman layer.
	TopsideExtensionChapmanConst IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst `json:"topsideExtensionChapmanConst,omitzero"`
	// Varying scale height Chapman topside layer.
	TopsideExtensionVaryChap IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap `json:"topsideExtensionVaryChap,omitzero"`
	// Array of valley model coefficients.
	ValleyModelCoeffs []float64 `json:"valleyModelCoeffs,omitzero"`
	paramObj
}

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfile) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

// Full set of the IRI formalism coefficients.
type IonOobservationUnvalidatedPublishParamsBodyDensityProfileIri struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfileIri) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfileIri
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe the E, F1, and F2 layers as parabolic-shape segments.
type IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolic struct {
	// General description of the QP computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Describes the E, F1, and F2 layers as parabolic-shape segments.
	ParabolicItems []IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem `json:"parabolicItems,omitzero"`
	paramObj
}

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}

// Describes the E, F1, and F2 layers as parabolic-shape segments.
type IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe profile shape as quasi-parabolic segments.
type IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic struct {
	// General description of the quasi-parabolic computation algorithm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Value of the Earth's radius, in kilometers, used for computations.
	EarthRadius param.Opt[float64] `json:"earthRadius,omitzero"`
	// Array of quasi-parabolic segments. Each segment is the best-fit 3-parameter
	// quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3
	// groups for E, F1, and F2 layers, but additional segments may be used to improve
	// accuracy.
	QuasiParabolicSegments []IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment `json:"quasiParabolicSegments,omitzero"`
	paramObj
}

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic
	return param.MarshalObject(r, (*shadow)(&r))
}

// A quasi-parabolic segment is the best-fit 3-parameter quasi-parabolas defined
// via A, B, C coefficients, f^2=A/r^2+B/r+C”. Usually 3 groups for E, F1, and F2
// layers, but additional segments may be used to improve accuracy.
type IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients to describe either the E, F1, and bottomside F2 profile shapes or
// the height uncertainty boundaries.
type IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev struct {
	// Description of the computation technique.
	Description param.Opt[string] `json:"description,omitzero"`
	// Up to 3 groups of coefficients, using “shiftedChebyshev” sub-field, to describe
	// E, F1, and bottomside F2 profile shapes, or up to 6 groups of coefficients to
	// describe height uncertainty boundaries (upper and lower).
	ShiftedChebyshevs []IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev `json:"shiftedChebyshevs,omitzero"`
	paramObj
}

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}

// Coefficients, using ‘shiftedChebyshev’ sub-field, to describe E, F1, and
// bottomside F2 profile shapes, or height uncertainty boundaries (upper and
// lower).
type IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameters of the constant-scale-height Chapman layer.
type IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst
	return param.MarshalObject(r, (*shadow)(&r))
}

// Varying scale height Chapman topside layer.
type IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyDoppler struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyDoppler) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyDoppler
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyElevation struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyElevation) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyElevation
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyFrequency struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyFrequency) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyFrequency
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyPhase struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyPhase) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyPhase
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyPolarization struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyPolarization) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyPolarization
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyPower struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyPower) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyPower
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyRange struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyRange) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyRange
	return param.MarshalObject(r, (*shadow)(&r))
}

// The ScalerInfo record describes the person or system who interpreted the
// ionogram in IonoObservation.
type IonOobservationUnvalidatedPublishParamsBodyScalerInfo struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyScalerInfo) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyScalerInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyStokes struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyStokes) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyStokes
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyTime struct {
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

func (r IonOobservationUnvalidatedPublishParamsBodyTime) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyTime
	return param.MarshalObject(r, (*shadow)(&r))
}

type IonOobservationUnvalidatedPublishParamsBodyTraceGeneric struct {
	// Notes for the trace generic data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Multi-dimensional Array. The 1st dimension spans points along the trace while
	// the 2nd dimension spans frequency-range pairs.
	Data [][][]float64 `json:"data,omitzero"`
	// Array of dimension names for trace generic data.
	DimensionName []string `json:"dimensionName,omitzero"`
	paramObj
}

func (r IonOobservationUnvalidatedPublishParamsBodyTraceGeneric) MarshalJSON() (data []byte, err error) {
	type shadow IonOobservationUnvalidatedPublishParamsBodyTraceGeneric
	return param.MarshalObject(r, (*shadow)(&r))
}
