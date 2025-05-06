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
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// ObservationRadarobservationHistoryService contains methods and other services
// that help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationRadarobservationHistoryService] method instead.
type ObservationRadarobservationHistoryService struct {
	Options []option.RequestOption
}

// NewObservationRadarobservationHistoryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewObservationRadarobservationHistoryService(opts ...option.RequestOption) (r ObservationRadarobservationHistoryService) {
	r = ObservationRadarobservationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationRadarobservationHistoryService) List(ctx context.Context, query ObservationRadarobservationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ObservationRadarobservationHistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/radarobservation/history"
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
func (r *ObservationRadarobservationHistoryService) ListAutoPaging(ctx context.Context, query ObservationRadarobservationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ObservationRadarobservationHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationRadarobservationHistoryService) Aodr(ctx context.Context, query ObservationRadarobservationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/radarobservation/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationRadarobservationHistoryService) Count(ctx context.Context, query ObservationRadarobservationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/radarobservation/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of observation data for radar based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
type ObservationRadarobservationHistoryListResponse struct {
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
	DataMode ObservationRadarobservationHistoryListResponseDataMode `json:"dataMode,required"`
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// azimuth angle in degrees and topocentric frame.
	Azimuth float64 `json:"azimuth"`
	// Sensor azimuth angle bias in degrees.
	AzimuthBias float64 `json:"azimuthBias"`
	// Optional flag indicating whether the azimuth value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	AzimuthMeasured bool `json:"azimuthMeasured"`
	// Rate of change of the line of sight azimuth in degrees per second.
	AzimuthRate float64 `json:"azimuthRate"`
	// One sigma uncertainty in the line of sight azimuth angle measurement, in
	// degrees.
	AzimuthUnc float64 `json:"azimuthUnc"`
	// ID of the beam that produced this observation.
	Beam float64 `json:"beam"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Line of sight declination angle in degrees and J2000 coordinate frame.
	Declination float64 `json:"declination"`
	// Optional flag indicating whether the declination value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	DeclinationMeasured bool `json:"declinationMeasured"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Corrected doppler measurement in meters per second.
	Doppler float64 `json:"doppler"`
	// One sigma uncertainty in the corrected doppler measurement, in meters/second.
	DopplerUnc float64 `json:"dopplerUnc"`
	// Line of sight elevation in degrees and topocentric frame.
	Elevation float64 `json:"elevation"`
	// Sensor elevation bias in degrees.
	ElevationBias float64 `json:"elevationBias"`
	// Optional flag indicating whether the elevation value is measured (true) or
	// computed (false). If null, consumers may consult the data provider for
	// information regarding whether the corresponding value is computed or measured.
	ElevationMeasured bool `json:"elevationMeasured"`
	// Rate of change of the line of sight elevation in degrees per second.
	ElevationRate float64 `json:"elevationRate"`
	// One sigma uncertainty in the line of sight elevation angle measurement, in
	// degrees.
	ElevationUnc float64 `json:"elevationUnc"`
	// Unique identifier of the target on-orbit object, if correlated.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The position of this observation within a track (FENCE, FIRST, IN, LAST,
	// SINGLE). This identifier is optional and, if null, no assumption should be made
	// regarding whether other observations may or may not exist to compose a track.
	ObPosition string `json:"obPosition"`
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
	// Radar cross section in meters squared for orthogonal polarization.
	OrthogonalRcs float64 `json:"orthogonalRcs"`
	// one sigma uncertainty in orthogonal polarization Radar Cross Section, in
	// meters^2.
	OrthogonalRcsUnc float64 `json:"orthogonalRcsUnc"`
	// Line of sight right ascension in degrees and J2000 coordinate frame.
	Ra float64 `json:"ra"`
	// Optional flag indicating whether the ra value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RaMeasured bool `json:"raMeasured"`
	// Target range in km.
	Range float64 `json:"range"`
	// Range accelaration in km/s2.
	RangeAccel float64 `json:"rangeAccel"`
	// One sigma uncertainty in the range acceleration measurement, in
	// kilometers/(second^2).
	RangeAccelUnc float64 `json:"rangeAccelUnc"`
	// Sensor range bias in km.
	RangeBias float64 `json:"rangeBias"`
	// Optional flag indicating whether the range value is measured (true) or computed
	// (false). If null, consumers may consult the data provider for information
	// regarding whether the corresponding value is computed or measured.
	RangeMeasured bool `json:"rangeMeasured"`
	// Rate of change of the line of sight range in km/sec.
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
	// Radar cross section in meters squared for polarization principal.
	Rcs float64 `json:"rcs"`
	// one sigma uncertainty in principal polarization Radar Cross Section, in
	// meters^2.
	RcsUnc float64 `json:"rcsUnc"`
	// Satellite/Catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// The reference frame of the observing sensor state. If the senReferenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	SenReferenceFrame ObservationRadarobservationHistoryListResponseSenReferenceFrame `json:"senReferenceFrame"`
	// Sensor x position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senx float64 `json:"senx"`
	// Sensor y position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Seny float64 `json:"seny"`
	// Sensor z position in km at obTime (if mobile/onorbit) in the specified
	// senReferenceFrame. If senReferenceFrame is null then J2000 should be assumed.
	Senz float64 `json:"senz"`
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to indicate the specific tasking which produced this
	// observation.
	TaskID string `json:"taskId"`
	// Sensor timing bias in seconds.
	TimingBias float64 `json:"timingBias"`
	// Optional identifier of the track to which this observation belongs.
	TrackID string `json:"trackId"`
	// The beam type (or tracking state) in use at the time of collection of this
	// observation. Values include:
	//
	// INIT ACQ WITH INIT VALUES: Initial acquisition based on predefined initial
	// values such as position, velocity, or other specific parameters.
	//
	// INIT ACQ: Initial acquisition when no prior information or initial values such
	// as position or velocity are available.
	//
	// TRACKING SINGLE BEAM: Continuously tracks and monitors a single target using one
	// specific radar beam.
	//
	// TRACKING SEQUENTIAL ROVING: Sequentially tracks different targets or areas by
	// "roving" from one sector to the next in a systematic order.
	//
	// SELF ACQ WITH INIT VALUES: Autonomously acquires targets using predefined
	// starting parameters or initial values.
	//
	// SELF ACQ: Automatically detects and locks onto targets without the need for
	// predefined initial settings.
	//
	// NON-TRACKING: Non-tracking.
	TrackingState string `json:"trackingState"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Read only enumeration specifying the type of observation (e.g. OPTICAL, RADAR,
	// RF, etc).
	Type string `json:"type"`
	// Boolean indicating this observation is part of an uncorrelated track or was
	// unable to be correlated to a known object. This flag should only be set to true
	// by data providers after an attempt to correlate to an on-orbit object was made
	// and failed. If unable to correlate, the 'origObjectId' field may be populated
	// with an internal data provider specific identifier.
	Uct bool `json:"uct"`
	// X position of target in km in J2000 coordinate frame.
	X float64 `json:"x"`
	// X velocity of target in km/sec in J2000 coordinate frame.
	Xvel float64 `json:"xvel"`
	// Y position of target in km in J2000 coordinate frame.
	Y float64 `json:"y"`
	// Y velocity of target in km/sec in J2000 coordinate frame.
	Yvel float64 `json:"yvel"`
	// Z position of target in km in J2000 coordinate frame.
	Z float64 `json:"z"`
	// Z velocity of target in km/sec in J2000 coordinate frame.
	Zvel float64 `json:"zvel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ObTime                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Azimuth               respjson.Field
		AzimuthBias           respjson.Field
		AzimuthMeasured       respjson.Field
		AzimuthRate           respjson.Field
		AzimuthUnc            respjson.Field
		Beam                  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Declination           respjson.Field
		DeclinationMeasured   respjson.Field
		Descriptor            respjson.Field
		Doppler               respjson.Field
		DopplerUnc            respjson.Field
		Elevation             respjson.Field
		ElevationBias         respjson.Field
		ElevationMeasured     respjson.Field
		ElevationRate         respjson.Field
		ElevationUnc          respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		ObPosition            respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		OrthogonalRcs         respjson.Field
		OrthogonalRcsUnc      respjson.Field
		Ra                    respjson.Field
		RaMeasured            respjson.Field
		Range                 respjson.Field
		RangeAccel            respjson.Field
		RangeAccelUnc         respjson.Field
		RangeBias             respjson.Field
		RangeMeasured         respjson.Field
		RangeRate             respjson.Field
		RangeRateMeasured     respjson.Field
		RangeRateUnc          respjson.Field
		RangeUnc              respjson.Field
		RawFileUri            respjson.Field
		Rcs                   respjson.Field
		RcsUnc                respjson.Field
		SatNo                 respjson.Field
		SenReferenceFrame     respjson.Field
		Senx                  respjson.Field
		Seny                  respjson.Field
		Senz                  respjson.Field
		Snr                   respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		TaskID                respjson.Field
		TimingBias            respjson.Field
		TrackID               respjson.Field
		TrackingState         respjson.Field
		TransactionID         respjson.Field
		Type                  respjson.Field
		Uct                   respjson.Field
		X                     respjson.Field
		Xvel                  respjson.Field
		Y                     respjson.Field
		Yvel                  respjson.Field
		Z                     respjson.Field
		Zvel                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObservationRadarobservationHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationRadarobservationHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type ObservationRadarobservationHistoryListResponseDataMode string

const (
	ObservationRadarobservationHistoryListResponseDataModeReal      ObservationRadarobservationHistoryListResponseDataMode = "REAL"
	ObservationRadarobservationHistoryListResponseDataModeTest      ObservationRadarobservationHistoryListResponseDataMode = "TEST"
	ObservationRadarobservationHistoryListResponseDataModeSimulated ObservationRadarobservationHistoryListResponseDataMode = "SIMULATED"
	ObservationRadarobservationHistoryListResponseDataModeExercise  ObservationRadarobservationHistoryListResponseDataMode = "EXERCISE"
)

// The reference frame of the observing sensor state. If the senReferenceFrame is
// null it is assumed to be J2000.
type ObservationRadarobservationHistoryListResponseSenReferenceFrame string

const (
	ObservationRadarobservationHistoryListResponseSenReferenceFrameJ2000   ObservationRadarobservationHistoryListResponseSenReferenceFrame = "J2000"
	ObservationRadarobservationHistoryListResponseSenReferenceFrameEfgTdr  ObservationRadarobservationHistoryListResponseSenReferenceFrame = "EFG/TDR"
	ObservationRadarobservationHistoryListResponseSenReferenceFrameEcrEcef ObservationRadarobservationHistoryListResponseSenReferenceFrame = "ECR/ECEF"
	ObservationRadarobservationHistoryListResponseSenReferenceFrameTeme    ObservationRadarobservationHistoryListResponseSenReferenceFrame = "TEME"
	ObservationRadarobservationHistoryListResponseSenReferenceFrameItrf    ObservationRadarobservationHistoryListResponseSenReferenceFrame = "ITRF"
	ObservationRadarobservationHistoryListResponseSenReferenceFrameGcrf    ObservationRadarobservationHistoryListResponseSenReferenceFrame = "GCRF"
)

type ObservationRadarobservationHistoryListParams struct {
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

// URLQuery serializes [ObservationRadarobservationHistoryListParams]'s query
// parameters as `url.Values`.
func (r ObservationRadarobservationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationRadarobservationHistoryAodrParams struct {
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

// URLQuery serializes [ObservationRadarobservationHistoryAodrParams]'s query
// parameters as `url.Values`.
func (r ObservationRadarobservationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationRadarobservationHistoryCountParams struct {
	// Ob detection time in ISO 8601 UTC with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationRadarobservationHistoryCountParams]'s query
// parameters as `url.Values`.
func (r ObservationRadarobservationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
