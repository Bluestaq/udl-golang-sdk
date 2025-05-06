// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// ObservationRfObservationService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationRfObservationService] method instead.
type ObservationRfObservationService struct {
	Options []option.RequestOption
	History ObservationRfObservationHistoryService
}

// NewObservationRfObservationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservationRfObservationService(opts ...option.RequestOption) (r ObservationRfObservationService) {
	r = ObservationRfObservationService{}
	r.Options = opts
	r.History = NewObservationRfObservationHistoryService(opts...)
	return
}

// Service operation to take a single RF observation as a POST body and ingest into
// the database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *ObservationRfObservationService) New(ctx context.Context, body ObservationRfObservationNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfobservation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationRfObservationService) List(ctx context.Context, query ObservationRfObservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ObservationRfObservationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/rfobservation"
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
func (r *ObservationRfObservationService) ListAutoPaging(ctx context.Context, query ObservationRfObservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ObservationRfObservationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationRfObservationService) Count(ctx context.Context, query ObservationRfObservationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/rfobservation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of RF
// observations as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *ObservationRfObservationService) NewBulk(ctx context.Context, body ObservationRfObservationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfobservation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single RF observation by its unique ID passed as a
// path parameter.
func (r *ObservationRfObservationService) Get(ctx context.Context, id string, query ObservationRfObservationGetParams, opts ...option.RequestOption) (res *ObservationRfObservationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfobservation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ObservationRfObservationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfobservation/queryhelp"
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
func (r *ObservationRfObservationService) Tuple(ctx context.Context, query ObservationRfObservationTupleParams, opts ...option.RequestOption) (res *[]ObservationRfObservationTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/rfobservation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple RF observations as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *ObservationRfObservationService) UnvalidatedPublish(ctx context.Context, body ObservationRfObservationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-rf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
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
type ObservationRfObservationListResponse struct {
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
	DataMode ObservationRfObservationListResponseDataMode `json:"dataMode,required"`
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
	// azimuth angle in degrees and J2000 coordinate frame.
	Azimuth float64 `json:"azimuth"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured bool `json:"azimuthMeasured"`
	// Rate of change of the azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the azimuth angle measurement, in degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// Measured bandwidth in Hz.
	Bandwidth float64 `json:"bandwidth"`
	// Baud rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per second.
	BaudRate float64 `json:"baudRate"`
	// The ratio of bit errors per number of received bits.
	BitErrorRate float64 `json:"bitErrorRate"`
	// Carrier standard (e.g. DVB-S2, 802.11g, etc.).
	CarrierStandard string `json:"carrierStandard"`
	// Channel of the RFObservation record.
	Channel int64 `json:"channel"`
	// Collection mode (e.g. SURVEY, SPOT_SEARCH, NEIGHBORHOOD_WATCH, DIRECTED_SEARCH,
	// MANUAL, etc).
	CollectionMode string `json:"collectionMode"`
	// Confidence in the signal and its measurements and characterization.
	Confidence float64 `json:"confidence"`
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
	// Measured Equivalent Isotopically Radiated Power in dBW.
	Eirp float64 `json:"eirp"`
	// elevation in degrees and J2000 coordinate frame.
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
	// End carrier frequency in Hz.
	EndFrequency float64 `json:"endFrequency"`
	// Center carrier frequency in Hz.
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
	// Maximum measured PSD value of the trace in dBW.
	MaxPsd float64 `json:"maxPSD"`
	// Minimum measured PSD value of the trace in dBW.
	MinPsd float64 `json:"minPSD"`
	// Transponder modulation (e.g. Auto, QPSK, 8PSK, etc).
	Modulation string `json:"modulation"`
	// Noise power density, in dBW-Hz.
	NoisePwrDensity float64 `json:"noisePwrDensity"`
	// Expected bandwidth in Hz.
	NominalBandwidth float64 `json:"nominalBandwidth"`
	// Expected Equivalent Isotopically Radiated Power in dBW.
	NominalEirp float64 `json:"nominalEirp"`
	// Nominal or expected center carrier frequency in Hz.
	NominalFrequency float64 `json:"nominalFrequency"`
	// Expected carrier power over noise (dBW/Hz).
	NominalPowerOverNoise float64 `json:"nominalPowerOverNoise"`
	// Nominal or expected signal to noise ratio, in dB.
	NominalSnr float64 `json:"nominalSnr"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
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
	// pulse groups.
	Pgri float64 `json:"pgri"`
	// The antenna pointing dependent polarizer angle, in degrees.
	Polarity float64 `json:"polarity"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	PolarityType ObservationRfObservationListResponsePolarityType `json:"polarityType"`
	// Measured carrier power over noise (dBW/Hz).
	PowerOverNoise float64 `json:"powerOverNoise"`
	// Target range in km.
	Range float64 `json:"range"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured bool `json:"rangeMeasured"`
	// Rate of change of the range in km/sec.
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
	// Reference signal level, in dBW.
	ReferenceLevel float64 `json:"referenceLevel"`
	// Measured power of the center carrier frequency in dBW.
	RelativeCarrierPower float64 `json:"relativeCarrierPower"`
	// The measure of the signal created from the sum of all the noise sources and
	// unwanted signals within the measurement system, in dBW.
	RelativeNoiseFloor float64 `json:"relativeNoiseFloor"`
	// Resolution bandwidth in Hz.
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
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Measured spectrum analyzer power of the center carrier frequency in dBW.
	SpectrumAnalyzerPower float64 `json:"spectrumAnalyzerPower"`
	// Start carrier frequency in Hz.
	StartFrequency float64 `json:"startFrequency"`
	// Switch Point of the RFObservation record.
	SwitchPoint int64 `json:"switchPoint"`
	// Symbol to noise ratio, in dB.
	SymbolToNoiseRatio float64 `json:"symbolToNoiseRatio"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID string `json:"taskId"`
	// Optional identifier of the track to which this observation belongs.
	TrackID string `json:"trackId"`
	// Target track or apparent range in km.
	TrackRange float64 `json:"trackRange"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Transmit pulse shaping filter roll-off value.
	TransmitFilterRollOff float64 `json:"transmitFilterRollOff"`
	// Transmit pulse shaping filter typ (e.g. RRC).
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
	// Video bandwidth in Hz.
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
		BitErrorRate          respjson.Field
		CarrierStandard       respjson.Field
		Channel               respjson.Field
		CollectionMode        respjson.Field
		Confidence            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		DetectionStatus       respjson.Field
		Eirp                  respjson.Field
		Elevation             respjson.Field
		ElevationMeasured     respjson.Field
		ElevationRate         respjson.Field
		ElevationUnc          respjson.Field
		Elnot                 respjson.Field
		EndFrequency          respjson.Field
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
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		OuterCodingRate       respjson.Field
		Peak                  respjson.Field
		Pgri                  respjson.Field
		Polarity              respjson.Field
		PolarityType          respjson.Field
		PowerOverNoise        respjson.Field
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
		Snr                   respjson.Field
		SourceDl              respjson.Field
		SpectrumAnalyzerPower respjson.Field
		StartFrequency        respjson.Field
		SwitchPoint           respjson.Field
		SymbolToNoiseRatio    respjson.Field
		TaskID                respjson.Field
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
func (r ObservationRfObservationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationRfObservationListResponse) UnmarshalJSON(data []byte) error {
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
type ObservationRfObservationListResponseDataMode string

const (
	ObservationRfObservationListResponseDataModeReal      ObservationRfObservationListResponseDataMode = "REAL"
	ObservationRfObservationListResponseDataModeTest      ObservationRfObservationListResponseDataMode = "TEST"
	ObservationRfObservationListResponseDataModeSimulated ObservationRfObservationListResponseDataMode = "SIMULATED"
	ObservationRfObservationListResponseDataModeExercise  ObservationRfObservationListResponseDataMode = "EXERCISE"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
// surface.
type ObservationRfObservationListResponsePolarityType string

const (
	ObservationRfObservationListResponsePolarityTypeH ObservationRfObservationListResponsePolarityType = "H"
	ObservationRfObservationListResponsePolarityTypeV ObservationRfObservationListResponsePolarityType = "V"
	ObservationRfObservationListResponsePolarityTypeR ObservationRfObservationListResponsePolarityType = "R"
	ObservationRfObservationListResponsePolarityTypeL ObservationRfObservationListResponsePolarityType = "L"
)

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
type ObservationRfObservationGetResponse struct {
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
	DataMode ObservationRfObservationGetResponseDataMode `json:"dataMode,required"`
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
	// azimuth angle in degrees and J2000 coordinate frame.
	Azimuth float64 `json:"azimuth"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured bool `json:"azimuthMeasured"`
	// Rate of change of the azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the azimuth angle measurement, in degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// Measured bandwidth in Hz.
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
	// Collection mode (e.g. SURVEY, SPOT_SEARCH, NEIGHBORHOOD_WATCH, DIRECTED_SEARCH,
	// MANUAL, etc).
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
	// Array of detection statuses (e.g. DETECTED, CARRIER_DETECTED, NOT_DETECTED) for
	// each measured signal.
	DetectionStatuses []string `json:"detectionStatuses"`
	// Measured Equivalent Isotopically Radiated Power in dBW.
	Eirp float64 `json:"eirp"`
	// elevation in degrees and J2000 coordinate frame.
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
	// End carrier frequency in Hz.
	EndFrequency float64 `json:"endFrequency"`
	// Array of individual PSD frequencies of the signal in Hz. This array should
	// correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies"`
	// Center carrier frequency in Hz.
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
	// Maximum measured PSD value of the trace in dBW.
	MaxPsd float64 `json:"maxPSD"`
	// Minimum measured PSD value of the trace in dBW.
	MinPsd float64 `json:"minPSD"`
	// Transponder modulation (e.g. Auto, QPSK, 8PSK, etc).
	Modulation string `json:"modulation"`
	// Noise power density, in dBW-Hz.
	NoisePwrDensity float64 `json:"noisePwrDensity"`
	// Expected bandwidth in Hz.
	NominalBandwidth float64 `json:"nominalBandwidth"`
	// Expected Equivalent Isotopically Radiated Power in dBW.
	NominalEirp float64 `json:"nominalEirp"`
	// Nominal or expected center carrier frequency in Hz.
	NominalFrequency float64 `json:"nominalFrequency"`
	// Expected carrier power over noise (dBW/Hz).
	NominalPowerOverNoise float64 `json:"nominalPowerOverNoise"`
	// Nominal or expected signal to noise ratio, in dB.
	NominalSnr float64 `json:"nominalSnr"`
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
	// pulse groups.
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
	PolarityType ObservationRfObservationGetResponsePolarityType `json:"polarityType"`
	// Measured carrier power over noise (dBW/Hz).
	PowerOverNoise float64 `json:"powerOverNoise"`
	// Array of individual measured PSD powers of the signal in dBW. This array should
	// correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers"`
	// Target range in km.
	Range float64 `json:"range"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured bool `json:"rangeMeasured"`
	// Rate of change of the range in km/sec.
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
	// Reference signal level, in dBW.
	ReferenceLevel float64 `json:"referenceLevel"`
	// Measured power of the center carrier frequency in dBW.
	RelativeCarrierPower float64 `json:"relativeCarrierPower"`
	// The measure of the signal created from the sum of all the noise sources and
	// unwanted signals within the measurement system, in dBW.
	RelativeNoiseFloor float64 `json:"relativeNoiseFloor"`
	// Resolution bandwidth in Hz.
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
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
	// Array of signal to noise ratios of the signals, in dB.
	Snrs []float64 `json:"snrs"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Measured spectrum analyzer power of the center carrier frequency in dBW.
	SpectrumAnalyzerPower float64 `json:"spectrumAnalyzerPower"`
	// Start carrier frequency in Hz.
	StartFrequency float64 `json:"startFrequency"`
	// Switch Point of the RFObservation record.
	SwitchPoint int64 `json:"switchPoint"`
	// Symbol to noise ratio, in dB.
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
	// Target track or apparent range in km.
	TrackRange float64 `json:"trackRange"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Transmit pulse shaping filter roll-off value.
	TransmitFilterRollOff float64 `json:"transmitFilterRollOff"`
	// Transmit pulse shaping filter typ (e.g. RRC).
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
	// Video bandwidth in Hz.
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
		Origin                respjson.Field
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
func (r ObservationRfObservationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationRfObservationGetResponse) UnmarshalJSON(data []byte) error {
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
type ObservationRfObservationGetResponseDataMode string

const (
	ObservationRfObservationGetResponseDataModeReal      ObservationRfObservationGetResponseDataMode = "REAL"
	ObservationRfObservationGetResponseDataModeTest      ObservationRfObservationGetResponseDataMode = "TEST"
	ObservationRfObservationGetResponseDataModeSimulated ObservationRfObservationGetResponseDataMode = "SIMULATED"
	ObservationRfObservationGetResponseDataModeExercise  ObservationRfObservationGetResponseDataMode = "EXERCISE"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
// surface.
type ObservationRfObservationGetResponsePolarityType string

const (
	ObservationRfObservationGetResponsePolarityTypeH ObservationRfObservationGetResponsePolarityType = "H"
	ObservationRfObservationGetResponsePolarityTypeV ObservationRfObservationGetResponsePolarityType = "V"
	ObservationRfObservationGetResponsePolarityTypeR ObservationRfObservationGetResponsePolarityType = "R"
	ObservationRfObservationGetResponsePolarityTypeL ObservationRfObservationGetResponsePolarityType = "L"
)

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
type ObservationRfObservationTupleResponse struct {
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
	DataMode ObservationRfObservationTupleResponseDataMode `json:"dataMode,required"`
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
	// azimuth angle in degrees and J2000 coordinate frame.
	Azimuth float64 `json:"azimuth"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured bool `json:"azimuthMeasured"`
	// Rate of change of the azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the azimuth angle measurement, in degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// Measured bandwidth in Hz.
	Bandwidth float64 `json:"bandwidth"`
	// Baud rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per second.
	BaudRate float64 `json:"baudRate"`
	// The ratio of bit errors per number of received bits.
	BitErrorRate float64 `json:"bitErrorRate"`
	// Carrier standard (e.g. DVB-S2, 802.11g, etc.).
	CarrierStandard string `json:"carrierStandard"`
	// Channel of the RFObservation record.
	Channel int64 `json:"channel"`
	// Collection mode (e.g. SURVEY, SPOT_SEARCH, NEIGHBORHOOD_WATCH, DIRECTED_SEARCH,
	// MANUAL, etc).
	CollectionMode string `json:"collectionMode"`
	// Confidence in the signal and its measurements and characterization.
	Confidence float64 `json:"confidence"`
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
	// Measured Equivalent Isotopically Radiated Power in dBW.
	Eirp float64 `json:"eirp"`
	// elevation in degrees and J2000 coordinate frame.
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
	// End carrier frequency in Hz.
	EndFrequency float64 `json:"endFrequency"`
	// Center carrier frequency in Hz.
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
	// Maximum measured PSD value of the trace in dBW.
	MaxPsd float64 `json:"maxPSD"`
	// Minimum measured PSD value of the trace in dBW.
	MinPsd float64 `json:"minPSD"`
	// Transponder modulation (e.g. Auto, QPSK, 8PSK, etc).
	Modulation string `json:"modulation"`
	// Noise power density, in dBW-Hz.
	NoisePwrDensity float64 `json:"noisePwrDensity"`
	// Expected bandwidth in Hz.
	NominalBandwidth float64 `json:"nominalBandwidth"`
	// Expected Equivalent Isotopically Radiated Power in dBW.
	NominalEirp float64 `json:"nominalEirp"`
	// Nominal or expected center carrier frequency in Hz.
	NominalFrequency float64 `json:"nominalFrequency"`
	// Expected carrier power over noise (dBW/Hz).
	NominalPowerOverNoise float64 `json:"nominalPowerOverNoise"`
	// Nominal or expected signal to noise ratio, in dB.
	NominalSnr float64 `json:"nominalSnr"`
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
	// pulse groups.
	Pgri float64 `json:"pgri"`
	// The antenna pointing dependent polarizer angle, in degrees.
	Polarity float64 `json:"polarity"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	PolarityType ObservationRfObservationTupleResponsePolarityType `json:"polarityType"`
	// Measured carrier power over noise (dBW/Hz).
	PowerOverNoise float64 `json:"powerOverNoise"`
	// Target range in km.
	Range float64 `json:"range"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured bool `json:"rangeMeasured"`
	// Rate of change of the range in km/sec.
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
	// Reference signal level, in dBW.
	ReferenceLevel float64 `json:"referenceLevel"`
	// Measured power of the center carrier frequency in dBW.
	RelativeCarrierPower float64 `json:"relativeCarrierPower"`
	// The measure of the signal created from the sum of all the noise sources and
	// unwanted signals within the measurement system, in dBW.
	RelativeNoiseFloor float64 `json:"relativeNoiseFloor"`
	// Resolution bandwidth in Hz.
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
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Measured spectrum analyzer power of the center carrier frequency in dBW.
	SpectrumAnalyzerPower float64 `json:"spectrumAnalyzerPower"`
	// Start carrier frequency in Hz.
	StartFrequency float64 `json:"startFrequency"`
	// Switch Point of the RFObservation record.
	SwitchPoint int64 `json:"switchPoint"`
	// Symbol to noise ratio, in dB.
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
	// Optional identifier of the track to which this observation belongs.
	TrackID string `json:"trackId"`
	// Target track or apparent range in km.
	TrackRange float64 `json:"trackRange"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Transmit pulse shaping filter roll-off value.
	TransmitFilterRollOff float64 `json:"transmitFilterRollOff"`
	// Transmit pulse shaping filter typ (e.g. RRC).
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
	// Video bandwidth in Hz.
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
		BitErrorRate          respjson.Field
		CarrierStandard       respjson.Field
		Channel               respjson.Field
		CollectionMode        respjson.Field
		Confidence            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		DetectionStatus       respjson.Field
		Eirp                  respjson.Field
		Elevation             respjson.Field
		ElevationMeasured     respjson.Field
		ElevationRate         respjson.Field
		ElevationUnc          respjson.Field
		Elnot                 respjson.Field
		EndFrequency          respjson.Field
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
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		OuterCodingRate       respjson.Field
		Peak                  respjson.Field
		Pgri                  respjson.Field
		Polarity              respjson.Field
		PolarityType          respjson.Field
		PowerOverNoise        respjson.Field
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
		Snr                   respjson.Field
		SourceDl              respjson.Field
		SpectrumAnalyzerPower respjson.Field
		StartFrequency        respjson.Field
		SwitchPoint           respjson.Field
		SymbolToNoiseRatio    respjson.Field
		Tags                  respjson.Field
		TaskID                respjson.Field
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
func (r ObservationRfObservationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationRfObservationTupleResponse) UnmarshalJSON(data []byte) error {
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
type ObservationRfObservationTupleResponseDataMode string

const (
	ObservationRfObservationTupleResponseDataModeReal      ObservationRfObservationTupleResponseDataMode = "REAL"
	ObservationRfObservationTupleResponseDataModeTest      ObservationRfObservationTupleResponseDataMode = "TEST"
	ObservationRfObservationTupleResponseDataModeSimulated ObservationRfObservationTupleResponseDataMode = "SIMULATED"
	ObservationRfObservationTupleResponseDataModeExercise  ObservationRfObservationTupleResponseDataMode = "EXERCISE"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
// surface.
type ObservationRfObservationTupleResponsePolarityType string

const (
	ObservationRfObservationTupleResponsePolarityTypeH ObservationRfObservationTupleResponsePolarityType = "H"
	ObservationRfObservationTupleResponsePolarityTypeV ObservationRfObservationTupleResponsePolarityType = "V"
	ObservationRfObservationTupleResponsePolarityTypeR ObservationRfObservationTupleResponsePolarityType = "R"
	ObservationRfObservationTupleResponsePolarityTypeL ObservationRfObservationTupleResponsePolarityType = "L"
)

type ObservationRfObservationNewParams struct {
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
	DataMode ObservationRfObservationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of RF ob (e.g. RF, RF-SOSI, PSD, RFI, SPOOF, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Antenna name of the RFObservation record.
	AntennaName param.Opt[string] `json:"antennaName,omitzero"`
	// azimuth angle in degrees and J2000 coordinate frame.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured param.Opt[bool] `json:"azimuthMeasured,omitzero"`
	// Rate of change of the azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the azimuth angle measurement, in degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// Measured bandwidth in Hz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Baud rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per second.
	BaudRate param.Opt[float64] `json:"baudRate,omitzero"`
	// The ratio of bit errors per number of received bits.
	BitErrorRate param.Opt[float64] `json:"bitErrorRate,omitzero"`
	// Carrier standard (e.g. DVB-S2, 802.11g, etc.).
	CarrierStandard param.Opt[string] `json:"carrierStandard,omitzero"`
	// Channel of the RFObservation record.
	Channel param.Opt[int64] `json:"channel,omitzero"`
	// Collection mode (e.g. SURVEY, SPOT_SEARCH, NEIGHBORHOOD_WATCH, DIRECTED_SEARCH,
	// MANUAL, etc).
	CollectionMode param.Opt[string] `json:"collectionMode,omitzero"`
	// Confidence in the signal and its measurements and characterization.
	Confidence param.Opt[float64] `json:"confidence,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Detection status (e.g. DETECTED, CARRIER_ACQUIRING, CARRIER_DETECTED,
	// NOT_DETECTED, etc).
	DetectionStatus param.Opt[string] `json:"detectionStatus,omitzero"`
	// Measured Equivalent Isotopically Radiated Power in dBW.
	Eirp param.Opt[float64] `json:"eirp,omitzero"`
	// elevation in degrees and J2000 coordinate frame.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured param.Opt[bool] `json:"elevationMeasured,omitzero"`
	// Rate of change of the elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the elevation angle measurement, in degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// ELINT notation.
	Elnot param.Opt[string] `json:"elnot,omitzero"`
	// End carrier frequency in Hz.
	EndFrequency param.Opt[float64] `json:"endFrequency,omitzero"`
	// Center carrier frequency in Hz.
	Frequency param.Opt[float64] `json:"frequency,omitzero"`
	// Frequency Shift of the RFObservation record.
	FrequencyShift param.Opt[float64] `json:"frequencyShift,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// True if the signal is incoming, false if outgoing.
	Incoming param.Opt[bool] `json:"incoming,omitzero"`
	// Inner forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	InnerCodingRate param.Opt[int64] `json:"innerCodingRate,omitzero"`
	// Maximum measured PSD value of the trace in dBW.
	MaxPsd param.Opt[float64] `json:"maxPSD,omitzero"`
	// Minimum measured PSD value of the trace in dBW.
	MinPsd param.Opt[float64] `json:"minPSD,omitzero"`
	// Transponder modulation (e.g. Auto, QPSK, 8PSK, etc).
	Modulation param.Opt[string] `json:"modulation,omitzero"`
	// Noise power density, in dBW-Hz.
	NoisePwrDensity param.Opt[float64] `json:"noisePwrDensity,omitzero"`
	// Expected bandwidth in Hz.
	NominalBandwidth param.Opt[float64] `json:"nominalBandwidth,omitzero"`
	// Expected Equivalent Isotopically Radiated Power in dBW.
	NominalEirp param.Opt[float64] `json:"nominalEirp,omitzero"`
	// Nominal or expected center carrier frequency in Hz.
	NominalFrequency param.Opt[float64] `json:"nominalFrequency,omitzero"`
	// Expected carrier power over noise (dBW/Hz).
	NominalPowerOverNoise param.Opt[float64] `json:"nominalPowerOverNoise,omitzero"`
	// Nominal or expected signal to noise ratio, in dB.
	NominalSnr param.Opt[float64] `json:"nominalSnr,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Outer forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	OuterCodingRate param.Opt[int64] `json:"outerCodingRate,omitzero"`
	// Peak of the RFObservation record.
	Peak param.Opt[bool] `json:"peak,omitzero"`
	// A pulse group repetition interval (PGRI) is a pulse train in which there are
	// groups of closely spaced pulses separated by much longer times between these
	// pulse groups.
	Pgri param.Opt[float64] `json:"pgri,omitzero"`
	// The antenna pointing dependent polarizer angle, in degrees.
	Polarity param.Opt[float64] `json:"polarity,omitzero"`
	// Measured carrier power over noise (dBW/Hz).
	PowerOverNoise param.Opt[float64] `json:"powerOverNoise,omitzero"`
	// Target range in km.
	Range param.Opt[float64] `json:"range,omitzero"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured param.Opt[bool] `json:"rangeMeasured,omitzero"`
	// Rate of change of the range in km/sec.
	RangeRate param.Opt[float64] `json:"rangeRate,omitzero"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured param.Opt[bool] `json:"rangeRateMeasured,omitzero"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc param.Opt[float64] `json:"rangeRateUnc,omitzero"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc param.Opt[float64] `json:"rangeUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Reference signal level, in dBW.
	ReferenceLevel param.Opt[float64] `json:"referenceLevel,omitzero"`
	// Measured power of the center carrier frequency in dBW.
	RelativeCarrierPower param.Opt[float64] `json:"relativeCarrierPower,omitzero"`
	// The measure of the signal created from the sum of all the noise sources and
	// unwanted signals within the measurement system, in dBW.
	RelativeNoiseFloor param.Opt[float64] `json:"relativeNoiseFloor,omitzero"`
	// Resolution bandwidth in Hz.
	ResolutionBandwidth param.Opt[float64] `json:"resolutionBandwidth,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude at obTime (if mobile/onorbit) in km. If null, can be obtained
	// from sensor info.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// Measured spectrum analyzer power of the center carrier frequency in dBW.
	SpectrumAnalyzerPower param.Opt[float64] `json:"spectrumAnalyzerPower,omitzero"`
	// Start carrier frequency in Hz.
	StartFrequency param.Opt[float64] `json:"startFrequency,omitzero"`
	// Switch Point of the RFObservation record.
	SwitchPoint param.Opt[int64] `json:"switchPoint,omitzero"`
	// Symbol to noise ratio, in dB.
	SymbolToNoiseRatio param.Opt[float64] `json:"symbolToNoiseRatio,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Optional identifier of the track to which this observation belongs.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Target track or apparent range in km.
	TrackRange param.Opt[float64] `json:"trackRange,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Transmit pulse shaping filter roll-off value.
	TransmitFilterRollOff param.Opt[float64] `json:"transmitFilterRollOff,omitzero"`
	// Transmit pulse shaping filter typ (e.g. RRC).
	TransmitFilterType param.Opt[string] `json:"transmitFilterType,omitzero"`
	// Optional identifier provided by observation source to indicate the transponder
	// used for this measurement.
	Transponder param.Opt[string] `json:"transponder,omitzero"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Optional URL containing additional information on this observation.
	URL param.Opt[string] `json:"url,omitzero"`
	// Video bandwidth in Hz.
	VideoBandwidth param.Opt[float64] `json:"videoBandwidth,omitzero"`
	// Array of measured signal baud rates.
	BaudRates []float64 `json:"baudRates,omitzero"`
	// Array of chipRates.
	ChipRates []float64 `json:"chipRates,omitzero"`
	// Array of code fills.
	CodeFills []string `json:"codeFills,omitzero"`
	// Array of code lengths.
	CodeLengths []float64 `json:"codeLengths,omitzero"`
	// Array of code taps.
	CodeTaps []string `json:"codeTaps,omitzero"`
	// Array of measurement confidences.
	Confidences []float64 `json:"confidences,omitzero"`
	// Array of individual x-coordinates for demodulated signal constellation. This
	// array should correspond with the same-sized array of constellationYPoints.
	ConstellationXPoints []float64 `json:"constellationXPoints,omitzero"`
	// Array of individual y-coordinates for demodulated signal constellation. This
	// array should correspond with the same-sized array of constellationXPoints.
	ConstellationYPoints []float64 `json:"constellationYPoints,omitzero"`
	// Array of detection statuses (e.g. DETECTED, CARRIER_DETECTED, NOT_DETECTED) for
	// each measured signal.
	DetectionStatuses []string `json:"detectionStatuses,omitzero"`
	// Array of individual PSD frequencies of the signal in Hz. This array should
	// correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies,omitzero"`
	// Array of pnOrder.
	PnOrders []int64 `json:"pnOrders,omitzero"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	PolarityType ObservationRfObservationNewParamsPolarityType `json:"polarityType,omitzero"`
	// Array of individual measured PSD powers of the signal in dBW. This array should
	// correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers,omitzero"`
	// Array of optional source provided identifiers of the measurements/signals.
	SignalIDs []string `json:"signalIds,omitzero"`
	// Array of signal to noise ratios of the signals, in dB.
	Snrs []float64 `json:"snrs,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// Array of optional source provided telemetry identifiers of the
	// measurements/signals.
	TelemetryIDs []string `json:"telemetryIds,omitzero"`
	paramObj
}

func (r ObservationRfObservationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObservationRfObservationNewParams
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
type ObservationRfObservationNewParamsDataMode string

const (
	ObservationRfObservationNewParamsDataModeReal      ObservationRfObservationNewParamsDataMode = "REAL"
	ObservationRfObservationNewParamsDataModeTest      ObservationRfObservationNewParamsDataMode = "TEST"
	ObservationRfObservationNewParamsDataModeSimulated ObservationRfObservationNewParamsDataMode = "SIMULATED"
	ObservationRfObservationNewParamsDataModeExercise  ObservationRfObservationNewParamsDataMode = "EXERCISE"
)

// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
// surface.
type ObservationRfObservationNewParamsPolarityType string

const (
	ObservationRfObservationNewParamsPolarityTypeH ObservationRfObservationNewParamsPolarityType = "H"
	ObservationRfObservationNewParamsPolarityTypeV ObservationRfObservationNewParamsPolarityType = "V"
	ObservationRfObservationNewParamsPolarityTypeR ObservationRfObservationNewParamsPolarityType = "R"
	ObservationRfObservationNewParamsPolarityTypeL ObservationRfObservationNewParamsPolarityType = "L"
)

type ObservationRfObservationListParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationRfObservationListParams]'s query parameters as
// `url.Values`.
func (r ObservationRfObservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationRfObservationCountParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationRfObservationCountParams]'s query parameters as
// `url.Values`.
func (r ObservationRfObservationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationRfObservationNewBulkParams struct {
	Body []ObservationRfObservationNewBulkParamsBody
	paramObj
}

func (r ObservationRfObservationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
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
//
// The properties ClassificationMarking, DataMode, ObTime, Source, Type are
// required.
type ObservationRfObservationNewBulkParamsBody struct {
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
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of RF ob (e.g. RF, RF-SOSI, PSD, RFI, SPOOF, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Antenna name of the RFObservation record.
	AntennaName param.Opt[string] `json:"antennaName,omitzero"`
	// azimuth angle in degrees and J2000 coordinate frame.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured param.Opt[bool] `json:"azimuthMeasured,omitzero"`
	// Rate of change of the azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the azimuth angle measurement, in degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// Measured bandwidth in Hz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Baud rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per second.
	BaudRate param.Opt[float64] `json:"baudRate,omitzero"`
	// The ratio of bit errors per number of received bits.
	BitErrorRate param.Opt[float64] `json:"bitErrorRate,omitzero"`
	// Carrier standard (e.g. DVB-S2, 802.11g, etc.).
	CarrierStandard param.Opt[string] `json:"carrierStandard,omitzero"`
	// Channel of the RFObservation record.
	Channel param.Opt[int64] `json:"channel,omitzero"`
	// Collection mode (e.g. SURVEY, SPOT_SEARCH, NEIGHBORHOOD_WATCH, DIRECTED_SEARCH,
	// MANUAL, etc).
	CollectionMode param.Opt[string] `json:"collectionMode,omitzero"`
	// Confidence in the signal and its measurements and characterization.
	Confidence param.Opt[float64] `json:"confidence,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Detection status (e.g. DETECTED, CARRIER_ACQUIRING, CARRIER_DETECTED,
	// NOT_DETECTED, etc).
	DetectionStatus param.Opt[string] `json:"detectionStatus,omitzero"`
	// Measured Equivalent Isotopically Radiated Power in dBW.
	Eirp param.Opt[float64] `json:"eirp,omitzero"`
	// elevation in degrees and J2000 coordinate frame.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured param.Opt[bool] `json:"elevationMeasured,omitzero"`
	// Rate of change of the elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the elevation angle measurement, in degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// ELINT notation.
	Elnot param.Opt[string] `json:"elnot,omitzero"`
	// End carrier frequency in Hz.
	EndFrequency param.Opt[float64] `json:"endFrequency,omitzero"`
	// Center carrier frequency in Hz.
	Frequency param.Opt[float64] `json:"frequency,omitzero"`
	// Frequency Shift of the RFObservation record.
	FrequencyShift param.Opt[float64] `json:"frequencyShift,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// True if the signal is incoming, false if outgoing.
	Incoming param.Opt[bool] `json:"incoming,omitzero"`
	// Inner forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	InnerCodingRate param.Opt[int64] `json:"innerCodingRate,omitzero"`
	// Maximum measured PSD value of the trace in dBW.
	MaxPsd param.Opt[float64] `json:"maxPSD,omitzero"`
	// Minimum measured PSD value of the trace in dBW.
	MinPsd param.Opt[float64] `json:"minPSD,omitzero"`
	// Transponder modulation (e.g. Auto, QPSK, 8PSK, etc).
	Modulation param.Opt[string] `json:"modulation,omitzero"`
	// Noise power density, in dBW-Hz.
	NoisePwrDensity param.Opt[float64] `json:"noisePwrDensity,omitzero"`
	// Expected bandwidth in Hz.
	NominalBandwidth param.Opt[float64] `json:"nominalBandwidth,omitzero"`
	// Expected Equivalent Isotopically Radiated Power in dBW.
	NominalEirp param.Opt[float64] `json:"nominalEirp,omitzero"`
	// Nominal or expected center carrier frequency in Hz.
	NominalFrequency param.Opt[float64] `json:"nominalFrequency,omitzero"`
	// Expected carrier power over noise (dBW/Hz).
	NominalPowerOverNoise param.Opt[float64] `json:"nominalPowerOverNoise,omitzero"`
	// Nominal or expected signal to noise ratio, in dB.
	NominalSnr param.Opt[float64] `json:"nominalSnr,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Outer forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	OuterCodingRate param.Opt[int64] `json:"outerCodingRate,omitzero"`
	// Peak of the RFObservation record.
	Peak param.Opt[bool] `json:"peak,omitzero"`
	// A pulse group repetition interval (PGRI) is a pulse train in which there are
	// groups of closely spaced pulses separated by much longer times between these
	// pulse groups.
	Pgri param.Opt[float64] `json:"pgri,omitzero"`
	// The antenna pointing dependent polarizer angle, in degrees.
	Polarity param.Opt[float64] `json:"polarity,omitzero"`
	// Measured carrier power over noise (dBW/Hz).
	PowerOverNoise param.Opt[float64] `json:"powerOverNoise,omitzero"`
	// Target range in km.
	Range param.Opt[float64] `json:"range,omitzero"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured param.Opt[bool] `json:"rangeMeasured,omitzero"`
	// Rate of change of the range in km/sec.
	RangeRate param.Opt[float64] `json:"rangeRate,omitzero"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured param.Opt[bool] `json:"rangeRateMeasured,omitzero"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc param.Opt[float64] `json:"rangeRateUnc,omitzero"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc param.Opt[float64] `json:"rangeUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Reference signal level, in dBW.
	ReferenceLevel param.Opt[float64] `json:"referenceLevel,omitzero"`
	// Measured power of the center carrier frequency in dBW.
	RelativeCarrierPower param.Opt[float64] `json:"relativeCarrierPower,omitzero"`
	// The measure of the signal created from the sum of all the noise sources and
	// unwanted signals within the measurement system, in dBW.
	RelativeNoiseFloor param.Opt[float64] `json:"relativeNoiseFloor,omitzero"`
	// Resolution bandwidth in Hz.
	ResolutionBandwidth param.Opt[float64] `json:"resolutionBandwidth,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude at obTime (if mobile/onorbit) in km. If null, can be obtained
	// from sensor info.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Measured spectrum analyzer power of the center carrier frequency in dBW.
	SpectrumAnalyzerPower param.Opt[float64] `json:"spectrumAnalyzerPower,omitzero"`
	// Start carrier frequency in Hz.
	StartFrequency param.Opt[float64] `json:"startFrequency,omitzero"`
	// Switch Point of the RFObservation record.
	SwitchPoint param.Opt[int64] `json:"switchPoint,omitzero"`
	// Symbol to noise ratio, in dB.
	SymbolToNoiseRatio param.Opt[float64] `json:"symbolToNoiseRatio,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Optional identifier of the track to which this observation belongs.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Target track or apparent range in km.
	TrackRange param.Opt[float64] `json:"trackRange,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Transmit pulse shaping filter roll-off value.
	TransmitFilterRollOff param.Opt[float64] `json:"transmitFilterRollOff,omitzero"`
	// Transmit pulse shaping filter typ (e.g. RRC).
	TransmitFilterType param.Opt[string] `json:"transmitFilterType,omitzero"`
	// Optional identifier provided by observation source to indicate the transponder
	// used for this measurement.
	Transponder param.Opt[string] `json:"transponder,omitzero"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Optional URL containing additional information on this observation.
	URL param.Opt[string] `json:"url,omitzero"`
	// Video bandwidth in Hz.
	VideoBandwidth param.Opt[float64] `json:"videoBandwidth,omitzero"`
	// Array of measured signal baud rates.
	BaudRates []float64 `json:"baudRates,omitzero"`
	// Array of chipRates.
	ChipRates []float64 `json:"chipRates,omitzero"`
	// Array of code fills.
	CodeFills []string `json:"codeFills,omitzero"`
	// Array of code lengths.
	CodeLengths []float64 `json:"codeLengths,omitzero"`
	// Array of code taps.
	CodeTaps []string `json:"codeTaps,omitzero"`
	// Array of measurement confidences.
	Confidences []float64 `json:"confidences,omitzero"`
	// Array of individual x-coordinates for demodulated signal constellation. This
	// array should correspond with the same-sized array of constellationYPoints.
	ConstellationXPoints []float64 `json:"constellationXPoints,omitzero"`
	// Array of individual y-coordinates for demodulated signal constellation. This
	// array should correspond with the same-sized array of constellationXPoints.
	ConstellationYPoints []float64 `json:"constellationYPoints,omitzero"`
	// Array of detection statuses (e.g. DETECTED, CARRIER_DETECTED, NOT_DETECTED) for
	// each measured signal.
	DetectionStatuses []string `json:"detectionStatuses,omitzero"`
	// Array of individual PSD frequencies of the signal in Hz. This array should
	// correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies,omitzero"`
	// Array of pnOrder.
	PnOrders []int64 `json:"pnOrders,omitzero"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	PolarityType string `json:"polarityType,omitzero"`
	// Array of individual measured PSD powers of the signal in dBW. This array should
	// correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers,omitzero"`
	// Array of optional source provided identifiers of the measurements/signals.
	SignalIDs []string `json:"signalIds,omitzero"`
	// Array of signal to noise ratios of the signals, in dB.
	Snrs []float64 `json:"snrs,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// Array of optional source provided telemetry identifiers of the
	// measurements/signals.
	TelemetryIDs []string `json:"telemetryIds,omitzero"`
	paramObj
}

func (r ObservationRfObservationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationRfObservationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ObservationRfObservationNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ObservationRfObservationNewBulkParamsBody](
		"PolarityType", false, "H", "V", "R", "L",
	)
}

type ObservationRfObservationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationRfObservationGetParams]'s query parameters as
// `url.Values`.
func (r ObservationRfObservationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationRfObservationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationRfObservationTupleParams]'s query parameters as
// `url.Values`.
func (r ObservationRfObservationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationRfObservationUnvalidatedPublishParams struct {
	Body []ObservationRfObservationUnvalidatedPublishParamsBody
	paramObj
}

func (r ObservationRfObservationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
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
//
// The properties ClassificationMarking, DataMode, ObTime, Source, Type are
// required.
type ObservationRfObservationUnvalidatedPublishParamsBody struct {
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
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Type of RF ob (e.g. RF, RF-SOSI, PSD, RFI, SPOOF, etc).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Antenna name of the RFObservation record.
	AntennaName param.Opt[string] `json:"antennaName,omitzero"`
	// azimuth angle in degrees and J2000 coordinate frame.
	Azimuth param.Opt[float64] `json:"azimuth,omitzero"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured param.Opt[bool] `json:"azimuthMeasured,omitzero"`
	// Rate of change of the azimuth in degrees per second.
	AzimuthRate param.Opt[float64] `json:"azimuthRate,omitzero"`
	// One sigma uncertainty in the azimuth angle measurement, in degrees.
	AzimuthUnc param.Opt[float64] `json:"azimuthUnc,omitzero"`
	// Measured bandwidth in Hz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Baud rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per second.
	BaudRate param.Opt[float64] `json:"baudRate,omitzero"`
	// The ratio of bit errors per number of received bits.
	BitErrorRate param.Opt[float64] `json:"bitErrorRate,omitzero"`
	// Carrier standard (e.g. DVB-S2, 802.11g, etc.).
	CarrierStandard param.Opt[string] `json:"carrierStandard,omitzero"`
	// Channel of the RFObservation record.
	Channel param.Opt[int64] `json:"channel,omitzero"`
	// Collection mode (e.g. SURVEY, SPOT_SEARCH, NEIGHBORHOOD_WATCH, DIRECTED_SEARCH,
	// MANUAL, etc).
	CollectionMode param.Opt[string] `json:"collectionMode,omitzero"`
	// Confidence in the signal and its measurements and characterization.
	Confidence param.Opt[float64] `json:"confidence,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Detection status (e.g. DETECTED, CARRIER_ACQUIRING, CARRIER_DETECTED,
	// NOT_DETECTED, etc).
	DetectionStatus param.Opt[string] `json:"detectionStatus,omitzero"`
	// Measured Equivalent Isotopically Radiated Power in dBW.
	Eirp param.Opt[float64] `json:"eirp,omitzero"`
	// elevation in degrees and J2000 coordinate frame.
	Elevation param.Opt[float64] `json:"elevation,omitzero"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured param.Opt[bool] `json:"elevationMeasured,omitzero"`
	// Rate of change of the elevation in degrees per second.
	ElevationRate param.Opt[float64] `json:"elevationRate,omitzero"`
	// One sigma uncertainty in the elevation angle measurement, in degrees.
	ElevationUnc param.Opt[float64] `json:"elevationUnc,omitzero"`
	// ELINT notation.
	Elnot param.Opt[string] `json:"elnot,omitzero"`
	// End carrier frequency in Hz.
	EndFrequency param.Opt[float64] `json:"endFrequency,omitzero"`
	// Center carrier frequency in Hz.
	Frequency param.Opt[float64] `json:"frequency,omitzero"`
	// Frequency Shift of the RFObservation record.
	FrequencyShift param.Opt[float64] `json:"frequencyShift,omitzero"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// True if the signal is incoming, false if outgoing.
	Incoming param.Opt[bool] `json:"incoming,omitzero"`
	// Inner forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	InnerCodingRate param.Opt[int64] `json:"innerCodingRate,omitzero"`
	// Maximum measured PSD value of the trace in dBW.
	MaxPsd param.Opt[float64] `json:"maxPSD,omitzero"`
	// Minimum measured PSD value of the trace in dBW.
	MinPsd param.Opt[float64] `json:"minPSD,omitzero"`
	// Transponder modulation (e.g. Auto, QPSK, 8PSK, etc).
	Modulation param.Opt[string] `json:"modulation,omitzero"`
	// Noise power density, in dBW-Hz.
	NoisePwrDensity param.Opt[float64] `json:"noisePwrDensity,omitzero"`
	// Expected bandwidth in Hz.
	NominalBandwidth param.Opt[float64] `json:"nominalBandwidth,omitzero"`
	// Expected Equivalent Isotopically Radiated Power in dBW.
	NominalEirp param.Opt[float64] `json:"nominalEirp,omitzero"`
	// Nominal or expected center carrier frequency in Hz.
	NominalFrequency param.Opt[float64] `json:"nominalFrequency,omitzero"`
	// Expected carrier power over noise (dBW/Hz).
	NominalPowerOverNoise param.Opt[float64] `json:"nominalPowerOverNoise,omitzero"`
	// Nominal or expected signal to noise ratio, in dB.
	NominalSnr param.Opt[float64] `json:"nominalSnr,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by observation source to indicate the target
	// onorbit object of this observation. This may be an internal identifier and not
	// necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation. This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Outer forward error correction rate: 0 = Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 =
	// 5/6, 5 = 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	OuterCodingRate param.Opt[int64] `json:"outerCodingRate,omitzero"`
	// Peak of the RFObservation record.
	Peak param.Opt[bool] `json:"peak,omitzero"`
	// A pulse group repetition interval (PGRI) is a pulse train in which there are
	// groups of closely spaced pulses separated by much longer times between these
	// pulse groups.
	Pgri param.Opt[float64] `json:"pgri,omitzero"`
	// The antenna pointing dependent polarizer angle, in degrees.
	Polarity param.Opt[float64] `json:"polarity,omitzero"`
	// Measured carrier power over noise (dBW/Hz).
	PowerOverNoise param.Opt[float64] `json:"powerOverNoise,omitzero"`
	// Target range in km.
	Range param.Opt[float64] `json:"range,omitzero"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured param.Opt[bool] `json:"rangeMeasured,omitzero"`
	// Rate of change of the range in km/sec.
	RangeRate param.Opt[float64] `json:"rangeRate,omitzero"`
	// Optional flag indicating whether the rangeRate value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	RangeRateMeasured param.Opt[bool] `json:"rangeRateMeasured,omitzero"`
	// One sigma uncertainty in the range rate measurement, in kilometers/second.
	RangeRateUnc param.Opt[float64] `json:"rangeRateUnc,omitzero"`
	// One sigma uncertainty in the range measurement, in kilometers.
	RangeUnc param.Opt[float64] `json:"rangeUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Reference signal level, in dBW.
	ReferenceLevel param.Opt[float64] `json:"referenceLevel,omitzero"`
	// Measured power of the center carrier frequency in dBW.
	RelativeCarrierPower param.Opt[float64] `json:"relativeCarrierPower,omitzero"`
	// The measure of the signal created from the sum of all the noise sources and
	// unwanted signals within the measurement system, in dBW.
	RelativeNoiseFloor param.Opt[float64] `json:"relativeNoiseFloor,omitzero"`
	// Resolution bandwidth in Hz.
	ResolutionBandwidth param.Opt[float64] `json:"resolutionBandwidth,omitzero"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude at obTime (if mobile/onorbit) in km. If null, can be obtained
	// from sensor info.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -90 to 90 degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude at obTime (if mobile/onorbit) in degrees. If null, can be
	// obtained from sensor info. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Measured spectrum analyzer power of the center carrier frequency in dBW.
	SpectrumAnalyzerPower param.Opt[float64] `json:"spectrumAnalyzerPower,omitzero"`
	// Start carrier frequency in Hz.
	StartFrequency param.Opt[float64] `json:"startFrequency,omitzero"`
	// Switch Point of the RFObservation record.
	SwitchPoint param.Opt[int64] `json:"switchPoint,omitzero"`
	// Symbol to noise ratio, in dB.
	SymbolToNoiseRatio param.Opt[float64] `json:"symbolToNoiseRatio,omitzero"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Optional identifier of the track to which this observation belongs.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Target track or apparent range in km.
	TrackRange param.Opt[float64] `json:"trackRange,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Transmit pulse shaping filter roll-off value.
	TransmitFilterRollOff param.Opt[float64] `json:"transmitFilterRollOff,omitzero"`
	// Transmit pulse shaping filter typ (e.g. RRC).
	TransmitFilterType param.Opt[string] `json:"transmitFilterType,omitzero"`
	// Optional identifier provided by observation source to indicate the transponder
	// used for this measurement.
	Transponder param.Opt[string] `json:"transponder,omitzero"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Optional URL containing additional information on this observation.
	URL param.Opt[string] `json:"url,omitzero"`
	// Video bandwidth in Hz.
	VideoBandwidth param.Opt[float64] `json:"videoBandwidth,omitzero"`
	// Array of measured signal baud rates.
	BaudRates []float64 `json:"baudRates,omitzero"`
	// Array of chipRates.
	ChipRates []float64 `json:"chipRates,omitzero"`
	// Array of code fills.
	CodeFills []string `json:"codeFills,omitzero"`
	// Array of code lengths.
	CodeLengths []float64 `json:"codeLengths,omitzero"`
	// Array of code taps.
	CodeTaps []string `json:"codeTaps,omitzero"`
	// Array of measurement confidences.
	Confidences []float64 `json:"confidences,omitzero"`
	// Array of individual x-coordinates for demodulated signal constellation. This
	// array should correspond with the same-sized array of constellationYPoints.
	ConstellationXPoints []float64 `json:"constellationXPoints,omitzero"`
	// Array of individual y-coordinates for demodulated signal constellation. This
	// array should correspond with the same-sized array of constellationXPoints.
	ConstellationYPoints []float64 `json:"constellationYPoints,omitzero"`
	// Array of detection statuses (e.g. DETECTED, CARRIER_DETECTED, NOT_DETECTED) for
	// each measured signal.
	DetectionStatuses []string `json:"detectionStatuses,omitzero"`
	// Array of individual PSD frequencies of the signal in Hz. This array should
	// correspond with the same-sized array of powers.
	Frequencies []float64 `json:"frequencies,omitzero"`
	// Array of pnOrder.
	PnOrders []int64 `json:"pnOrders,omitzero"`
	// Transponder polarization e.g. H - (Horizontally Polarized) Perpendicular to
	// Earth's surface, V - (Vertically Polarized) Parallel to Earth's surface, L -
	// (Left Hand Circularly Polarized) Rotating left relative to the earth's surface,
	// R - (Right Hand Circularly Polarized) Rotating right relative to the earth's
	// surface.
	//
	// Any of "H", "V", "R", "L".
	PolarityType string `json:"polarityType,omitzero"`
	// Array of individual measured PSD powers of the signal in dBW. This array should
	// correspond with the same-sized array of frequencies.
	Powers []float64 `json:"powers,omitzero"`
	// Array of optional source provided identifiers of the measurements/signals.
	SignalIDs []string `json:"signalIds,omitzero"`
	// Array of signal to noise ratios of the signals, in dB.
	Snrs []float64 `json:"snrs,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// Array of optional source provided telemetry identifiers of the
	// measurements/signals.
	TelemetryIDs []string `json:"telemetryIds,omitzero"`
	paramObj
}

func (r ObservationRfObservationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationRfObservationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ObservationRfObservationUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ObservationRfObservationUnvalidatedPublishParamsBody](
		"PolarityType", false, "H", "V", "R", "L",
	)
}
