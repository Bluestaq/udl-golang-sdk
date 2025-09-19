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

// EmireportHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmireportHistoryService] method instead.
type EmireportHistoryService struct {
	Options []option.RequestOption
}

// NewEmireportHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmireportHistoryService(opts ...option.RequestOption) (r EmireportHistoryService) {
	r = EmireportHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EmireportHistoryService) List(ctx context.Context, query EmireportHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EmireportHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/emireport/history"
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
func (r *EmireportHistoryService) ListAutoPaging(ctx context.Context, query EmireportHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EmireportHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EmireportHistoryService) Aodr(ctx context.Context, query EmireportHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/emireport/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EmireportHistoryService) Count(ctx context.Context, query EmireportHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/emireport/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The EMI Report service supports the reporting, response, and
// resolution/mitigation for spectrum interference, and provides details regarding
// electromagnetic interference (EMI) detection, characterization, reporting,
// identification, geo-location, and resolution data for space-based and
// terrestrial systems.
type EmireportHistoryListResponse struct {
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
	DataMode EmireportHistoryListResponseDataMode `json:"dataMode,required"`
	// Flag indicating whether the affected mission is an ISR mission.
	Isr bool `json:"isr,required"`
	// User generated report identifier. This ID should remain the same on subsequent
	// updates to this report.
	ReportID string `json:"reportId,required"`
	// The reporting time of this EMI Report record, in ISO 8601 UTC format, with
	// millisecond precision.
	ReportTime time.Time `json:"reportTime,required" format:"date-time"`
	// The type of Electromagnetic Interference (EMI) being reported (GPS, SATCOM,
	// TERRESTRIAL).
	ReportType string `json:"reportType,required"`
	// Flag indicating whether assistance is being requested to address this EMI.
	RequestAssist bool `json:"requestAssist,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The EMI start time in ISO 8601 UTC format, with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Actions(s) taken to troubleshoot, mitigate, work-around, and/or resolve the EMI
	// impacts.
	ActionsTaken string `json:"actionsTaken"`
	// The specific type of activity affected by the reported EMI (e.g. UPLINK,
	// DOWNLINK, HF COMM, etc.).
	AffActivity string `json:"affActivity"`
	// Altitude of the affected receiver, expressed in meters above WGS-84 ellipsoid.
	Alt float64 `json:"alt"`
	// The Area Of Responsibility (AOR), Organization, or Combatant Command under which
	// the reported EMI pertains (AFRICOM, CENTCOM, EUCOM, INDOPACOM, NORTHCOM, SOCOM,
	// SOUTHCOM, SPACECOM, STRATCOM, TRANSCOM, UNKNOWN).
	Aor string `json:"aor"`
	// The band (EHF, SHF, UHF, etc.) affected by the EMI.
	Band string `json:"band"`
	// The beam pattern in use.
	BeamPattern string `json:"beamPattern"`
	// The channel affected by the EMI.
	Channel string `json:"channel"`
	// Flag indicating whether this interference appears to be illegally passing
	// traffic over a known channel.
	ChanPirate bool `json:"chanPirate"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Text description of the EMI particulars and other supporting information which
	// may be relevant to the cause and/or possible resolution of the issue.
	Description string `json:"description"`
	// Duration, Nature, Extent of impact.
	DneImpact string `json:"dneImpact"`
	// The type of EMI (i.e. BARRAGE, CARRIER WAVE, etc.), if known.
	EmiType string `json:"emiType"`
	// The EMI end time in ISO 8601 UTC format, with millisecond precision. The endTime
	// may be excluded if EMI is ongoing.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The affected frequency, in MHz.
	Frequency float64 `json:"frequency"`
	// Confidence ellipse centered about the detection location [semi-major axis (m),
	// semi-minor axis (m), orientation (deg) measured clockwise (0 - 360) from true
	// North].
	GeoLocErrEllp []float64 `json:"geoLocErrEllp"`
	// Flag indicating whether encryption is in use on the affected GPS frequency.
	GpsEncrypted bool `json:"gpsEncrypted"`
	// The affected GPS Frequency (L1, L2, etc.).
	GpsFreq string `json:"gpsFreq"`
	// The highest affected frequency, in MHz.
	HighAffectedFrequency float64 `json:"highAffectedFrequency"`
	// Unique identifier of the affected on-orbit object. For the public catalog, the
	// idOnOrbit is typically the satellite number as a string, but may be a UUID for
	// analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Flag indicating whether the EMI is a decipherable intercept over the affected
	// receiver. Additional information may be included in the description field
	// content of this record.
	Intercept bool `json:"intercept"`
	// The language heard over the intercepted source. Applicable when interceptType =
	// VOICE.
	InterceptLang string `json:"interceptLang"`
	// The type of transmission being intercepted (e.g. VOICE, etc.). Applicable when
	// intercept = TRUE.
	InterceptType string `json:"interceptType"`
	// The relative amplitude, in decibels (dB), of the interfering source, if known.
	IntSrcAmplitude float64 `json:"intSrcAmplitude"`
	// The bandwidth, in MHz, of the interfering source, if known.
	IntSrcBandwidth float64 `json:"intSrcBandwidth"`
	// The center frequency, in MHz, of the interfering source, if known.
	IntSrcCentFreq float64 `json:"intSrcCentFreq"`
	// Flag indicating whether the interfering source is encrypted.
	IntSrcEncrypted bool `json:"intSrcEncrypted"`
	// The modulation method (e.g. AM, FM, FSK, PSK, etc.) of the interfering source,
	// if known.
	IntSrcModulation string `json:"intSrcModulation"`
	// Flag indicating whether this EMI is impacting ISR collection.
	IsrCollectionImpact bool `json:"isrCollectionImpact"`
	// The location of the affected receiver, reported as a kill box.
	KillBox string `json:"killBox"`
	// WGS-84 latitude of the affected receiver, represented as -90 to 90 degrees
	// (negative values south of equator).
	Lat float64 `json:"lat"`
	// The name or identifier of the affected link.
	Link string `json:"link"`
	// WGS-84 longitude of the affected receiver, represented as -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The Military Grid Reference System (MGRS) location of the affected receiver. The
	// Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of a milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts:
	//
	// 4Q (grid zone designator, GZD) FJ (the 100,000-meter square identifier) 12345678
	// (numerical location; easting is 1234 and northing is 5678, in this case
	// specifying a location with 10 m resolution).
	MilGrid string `json:"milGrid"`
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
	// Optional identifier provided by the reporting source to indicate the affected
	// object of this report. This may be an internal identifier and not necessarily
	// map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// The persistence status (e.g. CONTINUOUS, INTERMITTENT, RANDOM, etc.) of the EMI.
	Persistence string `json:"persistence"`
	// The name or identifier of the affected platform.
	Platform string `json:"platform"`
	// The demodulation method (e.g. AM, FM, FSK, PSK, etc.) setting of the affected
	// receiver.
	RcvrDemod string `json:"rcvrDemod"`
	// The gain setting of the affected receiver, in decibels (dB).
	RcvrGain float64 `json:"rcvrGain"`
	// Description of the affected receiver location.
	RcvrLocation string `json:"rcvrLocation"`
	// The affected antenna type (e.g. DISH, OMNI, PHASED ARRAY, etc.) experiencing the
	// EMI.
	RcvrType string `json:"rcvrType"`
	// The responsible service branch under which this EMI falls (AIR FORCE, ARMY,
	// COAST GUARD, MARINES, NAVY).
	RespService string `json:"respService"`
	// The priority (LOW, MEDIUM, HIGH) of the affected SATCOM.
	SatcomPriority string `json:"satcomPriority"`
	// The downlink frequency, in MHz, of the impacted link.
	SatDownlinkFrequency float64 `json:"satDownlinkFrequency"`
	// The downlink polarization e.g. H - (Horizontally Polarized), V - (Vertically
	// Polarized), L - (Left Hand Circularly Polarized), R - (Right Hand Circularly
	// Polarized).
	SatDownlinkPolarization string `json:"satDownlinkPolarization"`
	// The name of the spacecraft whose link is being affected by the EMI.
	SatName string `json:"satName"`
	// Satellite/Catalog number of the affected OnOrbit object.
	SatNo int64 `json:"satNo"`
	// The name or identifier of the affected sat transponder.
	SatTransponderID string `json:"satTransponderId"`
	// The uplink frequency, in MHz, of the impacted link.
	SatUplinkFrequency float64 `json:"satUplinkFrequency"`
	// The uplink polarization e.g. H - (Horizontally Polarized), V - (Vertically
	// Polarized), L - (Left Hand Circularly Polarized), R - (Right Hand Circularly
	// Polarized).
	SatUplinkPolarization string `json:"satUplinkPolarization"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The reporting status (INITIAL, UPDATE, RESOLVED) of this EMI issue.
	Status string `json:"status"`
	// The ISR role of the impacted asset.
	SupportedIsrRole string `json:"supportedISRRole"`
	// The name or identifier of the affected system.
	System string `json:"system"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// The alternate country identifier in which the EMI occurred or is occurring.
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS.
	VictimAltCountry string `json:"victimAltCountry"`
	// The country code in which the EMI occurred or is occurring. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	VictimCountryCode string `json:"victimCountryCode"`
	// The victim functional impacts (e.g. C2, COMM DATA LINK, ISR SENSOR, PNT, etc.).
	VictimFuncImpacts string `json:"victimFuncImpacts"`
	// The e-mail contact of the reporting POC.
	VictimPocMail string `json:"victimPOCMail"`
	// The Point of Contact (POC) for this EMI Report.
	VictimPocName string `json:"victimPOCName"`
	// The phone number of the reporting POC, represented as digits only, no spaces or
	// special characters.
	VictimPocPhone string `json:"victimPOCPhone"`
	// The Unit or Organization of the reporting POC.
	VictimPocUnit string `json:"victimPOCUnit"`
	// The victim reaction (e.g. LOITER ORBIT, RETASK ASSET, RETURN TO BASE,
	// TROUBLESHOOT, etc.).
	VictimReaction string `json:"victimReaction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking   respjson.Field
		DataMode                respjson.Field
		Isr                     respjson.Field
		ReportID                respjson.Field
		ReportTime              respjson.Field
		ReportType              respjson.Field
		RequestAssist           respjson.Field
		Source                  respjson.Field
		StartTime               respjson.Field
		ID                      respjson.Field
		ActionsTaken            respjson.Field
		AffActivity             respjson.Field
		Alt                     respjson.Field
		Aor                     respjson.Field
		Band                    respjson.Field
		BeamPattern             respjson.Field
		Channel                 respjson.Field
		ChanPirate              respjson.Field
		CreatedAt               respjson.Field
		CreatedBy               respjson.Field
		Description             respjson.Field
		DneImpact               respjson.Field
		EmiType                 respjson.Field
		EndTime                 respjson.Field
		Frequency               respjson.Field
		GeoLocErrEllp           respjson.Field
		GpsEncrypted            respjson.Field
		GpsFreq                 respjson.Field
		HighAffectedFrequency   respjson.Field
		IDOnOrbit               respjson.Field
		Intercept               respjson.Field
		InterceptLang           respjson.Field
		InterceptType           respjson.Field
		IntSrcAmplitude         respjson.Field
		IntSrcBandwidth         respjson.Field
		IntSrcCentFreq          respjson.Field
		IntSrcEncrypted         respjson.Field
		IntSrcModulation        respjson.Field
		IsrCollectionImpact     respjson.Field
		KillBox                 respjson.Field
		Lat                     respjson.Field
		Link                    respjson.Field
		Lon                     respjson.Field
		MilGrid                 respjson.Field
		OnOrbit                 respjson.Field
		Origin                  respjson.Field
		OrigNetwork             respjson.Field
		OrigObjectID            respjson.Field
		Persistence             respjson.Field
		Platform                respjson.Field
		RcvrDemod               respjson.Field
		RcvrGain                respjson.Field
		RcvrLocation            respjson.Field
		RcvrType                respjson.Field
		RespService             respjson.Field
		SatcomPriority          respjson.Field
		SatDownlinkFrequency    respjson.Field
		SatDownlinkPolarization respjson.Field
		SatName                 respjson.Field
		SatNo                   respjson.Field
		SatTransponderID        respjson.Field
		SatUplinkFrequency      respjson.Field
		SatUplinkPolarization   respjson.Field
		SourceDl                respjson.Field
		Status                  respjson.Field
		SupportedIsrRole        respjson.Field
		System                  respjson.Field
		Tags                    respjson.Field
		TransactionID           respjson.Field
		VictimAltCountry        respjson.Field
		VictimCountryCode       respjson.Field
		VictimFuncImpacts       respjson.Field
		VictimPocMail           respjson.Field
		VictimPocName           respjson.Field
		VictimPocPhone          respjson.Field
		VictimPocUnit           respjson.Field
		VictimReaction          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmireportHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmireportHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type EmireportHistoryListResponseDataMode string

const (
	EmireportHistoryListResponseDataModeReal      EmireportHistoryListResponseDataMode = "REAL"
	EmireportHistoryListResponseDataModeTest      EmireportHistoryListResponseDataMode = "TEST"
	EmireportHistoryListResponseDataModeSimulated EmireportHistoryListResponseDataMode = "SIMULATED"
	EmireportHistoryListResponseDataModeExercise  EmireportHistoryListResponseDataMode = "EXERCISE"
)

type EmireportHistoryListParams struct {
	// The reporting time of this EMI Report record, in ISO 8601 UTC format, with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	ReportTime time.Time `query:"reportTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmireportHistoryListParams]'s query parameters as
// `url.Values`.
func (r EmireportHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmireportHistoryAodrParams struct {
	// The reporting time of this EMI Report record, in ISO 8601 UTC format, with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	ReportTime time.Time `query:"reportTime,required" format:"date-time" json:"-"`
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

// URLQuery serializes [EmireportHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r EmireportHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmireportHistoryCountParams struct {
	// The reporting time of this EMI Report record, in ISO 8601 UTC format, with
	// millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	ReportTime  time.Time        `query:"reportTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmireportHistoryCountParams]'s query parameters as
// `url.Values`.
func (r EmireportHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
