// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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

// ObservationRfObservationHistoryService contains methods and other services that
// help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationRfObservationHistoryService] method instead.
type ObservationRfObservationHistoryService struct {
	Options []option.RequestOption
}

// NewObservationRfObservationHistoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservationRfObservationHistoryService(opts ...option.RequestOption) (r ObservationRfObservationHistoryService) {
	r = ObservationRfObservationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationRfObservationHistoryService) List(ctx context.Context, query ObservationRfObservationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ObservationRfObservationHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/rfobservation/history"
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
func (r *ObservationRfObservationHistoryService) ListAutoPaging(ctx context.Context, query ObservationRfObservationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ObservationRfObservationHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationRfObservationHistoryService) Aodr(ctx context.Context, query ObservationRfObservationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/rfobservation/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationRfObservationHistoryService) Count(ctx context.Context, query ObservationRfObservationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/rfobservation/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of observation data for active/passive radio frequency (RF)
// based sensor phenomenologies. J2000 is the preferred coordinate frame for all
// observations, but in some cases observations may be in another frame depending
// on the provider. Please see the 'Discover' tab in the storefront to confirm
// coordinate frames by data provider. RF observations include several optional
// ordered arrays which are used to provide detailed information on recorded
// signals such as power spectral density lists or active signals (code taps/fills,
// etc). In these cases, the sizes of the arrays must match and can be assumed to
// have consistent indexing across arrays (e.g. powers[0] is the measured power at
// frequencies[0]).
type ObservationRfObservationHistoryListResponse struct {
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
	DataMode ObservationRfObservationHistoryListResponseDataMode `json:"dataMode,required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of RF ob (e.g. RF, RF-SOSI, PSD, RFI, SPOOF, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Antenna name of the RFObservation record.
	AntennaName string `json:"antennaName"`
	// Azimuth angle in degrees and topocentric coordinate frame.
	Azimuth float64 `json:"azimuth"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured bool `json:"azimuthMeasured"`
	// Rate of change of the azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the azimuth angle measurement, in degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// Measured bandwidth in hertz.
	Bandwidth float64 `json:"bandwidth"`
	// Baud rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per second.
	BaudRate float64 `json:"baudRate"`
	// Array of measured signal baud rates.
	BaudRates []float64 `json:"baudRates"`
	// The ratio of bit errors per number of received bits.
	BitErrorRate float64 `json:"bitErrorRate"`
	// Carrier standard (e.g. DVB-S2, 802.11g, etc.).
	CarrierStandard string `json:"carrierStandard"`
	// Channel of the RFObservation record.
	Channel int64 `json:"channel"`
	// Array of chipRates.
	ChipRates []float64 `json:"chipRates"`
	// Array of code fills.
	CodeFills []string `json:"codeFills"`
	// Array of code lengths.
	CodeLengths []float64 `json:"codeLengths"`
	// Array of code taps.
	CodeTaps []string `json:"codeTaps"`
	// Collection mode (e.g. CONTINUOUS, MANUAL, NEIGHBORHOOD_WATCH, DIRECTED_SEARCH,
	// SPOT_SEARCH, SURVEY, etc).
	CollectionMode string `json:"collectionMode"`
	// Confidence in the signal and its measurements and characterization.
	Confidence float64 `json:"confidence"`
	// Array of measurement confidences.
	Confidences []float64 `json:"confidences"`
	// Array of individual x-coordinates for demodulated signal constellation. This
	// array should correspond with the same-sized array of constellationYPoints.
	ConstellationXPoints []float64 `json:"constellationXPoints"`
	// Array of individual y-coordinates for demodulated signal constellation. This
	// array should correspond with the same-sized array of constellationXPoints.
	ConstellationYPoints []float64 `json:"constellationYPoints"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Detection status (e.g. DETECTED, CARRIER_ACQUIRING, CARRIER_DETECTED,
	// NOT_DETECTED, etc).
	DetectionStatus string `json:"detectionStatus"`
	// Array of detection statuses (e.g. CARRIER_DETECTED, DETECTED, NOT_DETECTED) for
	// each measured signal.
	DetectionStatuses []string `json:"detectionStatuses"`
	// Measured Equivalent Isotopically Radiated Power in decibel watts.
	Eirp float64 `json:"eirp"`
	// Elevation in degrees and topocentric coordinate frame.
	Elevation float64 `json:"elevation"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured bool `json:"elevationMeasured"`
	// Rate of change of the elevation in degrees per second.
	ElevationRate float64 `json:"elevationRate"`
	// One sigma uncertainty in the elevation angle measurement, in degrees.
	ElevationUnc float64 `json:"elevationUnc"`
	// ELINT notation.
	Elnot string `json:"elnot"`
	// End carrier frequency in hertz.
	EndFrequency float64 `json:"endFrequency"`
	// Array of imaginary components of the complex Fast Fourier Transform (FFT)
	// coefficients from the signal. Used together with the same-sized fftRealCoeffs
	// array to preserve both amplitude and phase information. This array should
	// correspond with the same-sized array of frequencies.
	FftImagCoeffs []float64 `json:"fftImagCoeffs"`
	// Array of real components of the complex Fast Fourier Transform (FFT)
	// coefficients from the signal. Used together with the same-sized fftImagCoeffs
	// array to preserve both amplitude and phase information. This array should
	// correspond with the same-sized array of frequencies.
	FftRealCoeffs []float64 `json:"fftRealCoeffs"`
	// Array of individual PSD frequencies of the signal in hertz. This array should
	// correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies"`
	// Center carrier frequency in hertz.
	Frequency float64 `json:"frequency"`
	// Frequency Shift of the RFObservation record.
	FrequencyShift float64 `json:"frequencyShift"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// True if the signal is incoming, false if outgoing.
	Incoming bool `json:"incoming"`
	// Inner forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	InnerCodingRate int64 `json:"innerCodingRate"`
	// Maximum measured PSD value of the trace in decibel watts.
	MaxPsd float64 `json:"maxPSD"`
	// Minimum measured PSD value of the trace in decibel watts.
	MinPsd float64 `json:"minPSD"`
	// Transponder modulation (e.g. Auto, QPSK, 8PSK, etc).
	Modulation string `json:"modulation"`
	// Noise power density, in decibel watts per hertz.
	NoisePwrDensity float64 `json:"noisePwrDensity"`
	// Expected bandwidth in hertz.
	NominalBandwidth float64 `json:"nominalBandwidth"`
	// Expected Equivalent Isotopically Radiated Power in decibel watts.
	NominalEirp float64 `json:"nominalEirp"`
	// Nominal or expected center carrier frequency in hertz.
	NominalFrequency float64 `json:"nominalFrequency"`
	// Expected carrier power over noise (decibel watts per hertz).
	NominalPowerOverNoise float64 `json:"nominalPowerOverNoise"`
	// Nominal or expected signal to noise ratio, in decibels.
	NominalSnr float64 `json:"nominalSnr"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Country of origin in which the data was originally posted.
	OrigCountry string `json:"origCountry"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// Original security marking that the data was marked with.
	OrigMarking string `json:"origMarking"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Outer forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	OuterCodingRate int64 `json:"outerCodingRate"`
	// Peak of the RFObservation record.
	Peak bool `json:"peak"`
	// A pulse group repetition interval (PGRI) is a pulse train in which there are
	// groups of closely spaced pulses separated by much longer times between these
	// pulse groups. The PGRI is measured in seconds.
	Pgri float64 `json:"pgri"`
	// Array of pnOrder.
	PnOrders []int64 `json:"pnOrders"`
	// The antenna pointing dependent polarizer angle, in degrees.
	Polarity float64 `json:"polarity"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	PolarityType ObservationRfObservationHistoryListResponsePolarityType `json:"polarityType"`
	// Measured carrier power over noise (decibel watts per hertz).
	PowerOverNoise float64 `json:"powerOverNoise"`
	// Array of individual measured PSD powers of the signal in decibel watts. This
	// array should correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers"`
	// Target range in kilometers.
	Range float64 `json:"range"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured bool `json:"rangeMeasured"`
	// Rate of change of the range in kilometers per second.
	RangeRate float64 `json:"rangeRate"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured bool `json:"rangeRateMeasured"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc float64 `json:"rangeRateUnc"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc float64 `json:"rangeUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Reference signal level, in decibel watts.
	ReferenceLevel float64 `json:"referenceLevel"`
	// Measured power of the center carrier frequency in decibel watts.
	RelativeCarrierPower float64 `json:"relativeCarrierPower"`
	// The measure of the signal created from the sum of all the noise sources and
	// unwanted signals within the measurement system, in decibel watts.
	RelativeNoiseFloor float64 `json:"relativeNoiseFloor"`
	// Resolution bandwidth in hertz.
	ResolutionBandwidth float64 `json:"resolutionBandwidth"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Sensor altitude at obTime (if mobile/onorbit) in km. If null, can be obtained
	// from sensor info.
	Senalt float64 `json:"senalt"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon float64 `json:"senlon"`
	// Array of optional source provided identifiers of the measurements/signals.
	SignalIDs []string `json:"signalIds"`
	// Signal to noise ratio, in decibels.
	Snr float64 `json:"snr"`
	// Array of signal to noise ratios of the signals, in decibels.
	Snrs []float64 `json:"snrs"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Measured spectrum analyzer power of the center carrier frequency in decibel
	// watts.
	SpectrumAnalyzerPower float64 `json:"spectrumAnalyzerPower"`
	// Start carrier frequency in hertz.
	StartFrequency float64 `json:"startFrequency"`
	// Switch Point of the RFObservation record.
	SwitchPoint int64 `json:"switchPoint"`
	// Symbol to noise ratio, in decibels.
	SymbolToNoiseRatio float64 `json:"symbolToNoiseRatio"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID string `json:"taskId"`
	// Array of optional source provided telemetry identifiers of the
	// measurements/signals.
	TelemetryIDs []string `json:"telemetryIds"`
	// Optional identifier of the track to which this observation belongs.
	TrackID string `json:"trackId"`
	// Target track or apparent range in kilometers.
	TrackRange float64 `json:"trackRange"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Transmit pulse shaping filter roll-off value.
	TransmitFilterRollOff float64 `json:"transmitFilterRollOff"`
	// Transmit pulse shaping filter type (e.g. RRC).
	TransmitFilterType string `json:"transmitFilterType"`
	// Optional identifier provided by observation source to indicate the transponder
	// used for this measurement.
	Transponder string `json:"transponder"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct bool `json:"uct"`
	// Optional URL containing additional information on this observation.
	URL string `json:"url"`
	// Video bandwidth in hertz.
	VideoBandwidth float64 `json:"videoBandwidth"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AntennaName           respjson.Field
		Azimuth               respjson.Field
		AzimuthMeasured       respjson.Field
		AzimuthRate           respjson.Field
		AzimuthUnc            respjson.Field
		Bandwidth             respjson.Field
		BaudRate              respjson.Field
		BaudRates             respjson.Field
		BitErrorRate          respjson.Field
		CarrierStandard       respjson.Field
		Channel               respjson.Field
		ChipRates             respjson.Field
		CodeFills             respjson.Field
		CodeLengths           respjson.Field
		CodeTaps              respjson.Field
		CollectionMode        respjson.Field
		Confidence            respjson.Field
		Confidences           respjson.Field
		ConstellationXPoints  respjson.Field
		ConstellationYPoints  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		DetectionStatus       respjson.Field
		DetectionStatuses     respjson.Field
		Eirp                  respjson.Field
		Elevation             respjson.Field
		ElevationMeasured     respjson.Field
		ElevationRate         respjson.Field
		ElevationUnc          respjson.Field
		Elnot                 respjson.Field
		EndFrequency          respjson.Field
		FftImagCoeffs         respjson.Field
		FftRealCoeffs         respjson.Field
		Frequencies           respjson.Field
		Frequency             respjson.Field
		FrequencyShift        respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		Incoming              respjson.Field
		InnerCodingRate       respjson.Field
		MaxPsd                respjson.Field
		MinPsd                respjson.Field
		Modulation            respjson.Field
		NoisePwrDensity       respjson.Field
		NominalBandwidth      respjson.Field
		NominalEirp           respjson.Field
		NominalFrequency      respjson.Field
		NominalPowerOverNoise respjson.Field
		NominalSnr            respjson.Field
		OnOrbit               respjson.Field
		OrigCountry           respjson.Field
		Origin                respjson.Field
		OrigMarking           respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		OuterCodingRate       respjson.Field
		Peak                  respjson.Field
		Pgri                  respjson.Field
		PnOrders              respjson.Field
		Polarity              respjson.Field
		PolarityType          respjson.Field
		PowerOverNoise        respjson.Field
		Powers                respjson.Field
		Range                 respjson.Field
		RangeMeasured         respjson.Field
		RangeRate             respjson.Field
		RangeRateMeasured     respjson.Field
		RangeRateUnc          respjson.Field
		RangeUnc              respjson.Field
		RawFileUri            respjson.Field
		ReferenceLevel        respjson.Field
		RelativeCarrierPower  respjson.Field
		RelativeNoiseFloor    respjson.Field
		ResolutionBandwidth   respjson.Field
		SatNo                 respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		SignalIDs             respjson.Field
		Snr                   respjson.Field
		Snrs                  respjson.Field
		SourceDl              respjson.Field
		SpectrumAnalyzerPower respjson.Field
		StartFrequency        respjson.Field
		SwitchPoint           respjson.Field
		SymbolToNoiseRatio    respjson.Field
		Tags                  respjson.Field
		TaskID                respjson.Field
		TelemetryIDs          respjson.Field
		TrackID               respjson.Field
		TrackRange            respjson.Field
		TransactionID         respjson.Field
		TransmitFilterRollOff respjson.Field
		TransmitFilterType    respjson.Field
		Transponder           respjson.Field
		Uct                   respjson.Field
		URL                   respjson.Field
		VideoBandwidth        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObservationRfObservationHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationRfObservationHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type ObservationRfObservationHistoryListResponseDataMode string

const (
	ObservationRfObservationHistoryListResponseDataModeReal      ObservationRfObservationHistoryListResponseDataMode = "REAL"
	ObservationRfObservationHistoryListResponseDataModeTest      ObservationRfObservationHistoryListResponseDataMode = "TEST"
	ObservationRfObservationHistoryListResponseDataModeSimulated ObservationRfObservationHistoryListResponseDataMode = "SIMULATED"
	ObservationRfObservationHistoryListResponseDataModeExercise  ObservationRfObservationHistoryListResponseDataMode = "EXERCISE"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
// surface.
type ObservationRfObservationHistoryListResponsePolarityType string

const (
	ObservationRfObservationHistoryListResponsePolarityTypeH ObservationRfObservationHistoryListResponsePolarityType = "H"
	ObservationRfObservationHistoryListResponsePolarityTypeV ObservationRfObservationHistoryListResponsePolarityType = "V"
	ObservationRfObservationHistoryListResponsePolarityTypeR ObservationRfObservationHistoryListResponsePolarityType = "R"
	ObservationRfObservationHistoryListResponsePolarityTypeL ObservationRfObservationHistoryListResponsePolarityType = "L"
)

type ObservationRfObservationHistoryListParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime time.Time `query:"obTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationRfObservationHistoryListParams]'s query
// parameters as `url.Values`.
func (r ObservationRfObservationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationRfObservationHistoryAodrParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime time.Time `query:"obTime,required" format:"date-time" json:"-"`
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

// URLQuery serializes [ObservationRfObservationHistoryAodrParams]'s query
// parameters as `url.Values`.
func (r ObservationRfObservationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationRfObservationHistoryCountParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationRfObservationHistoryCountParams]'s query
// parameters as `url.Values`.
func (r ObservationRfObservationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
