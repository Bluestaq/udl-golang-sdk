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
)

// ObservationObscorrelationHistoryService contains methods and other services that
// help with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationObscorrelationHistoryService] method instead.
type ObservationObscorrelationHistoryService struct {
	Options []option.RequestOption
}

// NewObservationObscorrelationHistoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservationObscorrelationHistoryService(opts ...option.RequestOption) (r ObservationObscorrelationHistoryService) {
	r = ObservationObscorrelationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationObscorrelationHistoryService) List(ctx context.Context, query ObservationObscorrelationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ObservationObscorrelationHistoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/obscorrelation/history"
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
func (r *ObservationObscorrelationHistoryService) ListAutoPaging(ctx context.Context, query ObservationObscorrelationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ObservationObscorrelationHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationObscorrelationHistoryService) Aodr(ctx context.Context, query ObservationObscorrelationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/obscorrelation/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationObscorrelationHistoryService) Count(ctx context.Context, query ObservationObscorrelationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/obscorrelation/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation supporting post-pass correlation of UCTs and re-correlation
// of mis-tagged electro-optical (EO), radar, RF, and DOA track/observations.
type ObservationObscorrelationHistoryListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the type of correlation is OBSERVATION or TRACK:
	// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
	// the entire track of observations is being correlated.
	//
	// Any of "OBSERVATION", "TRACK".
	CorrType ObservationObscorrelationHistoryListResponseCorrType `json:"corrType,required"`
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
	DataMode ObservationObscorrelationHistoryListResponseDataMode `json:"dataMode,required"`
	// Correlation message generation time, in ISO 8601 UTC format with millisecond
	// precision.
	MsgTs time.Time `json:"msgTs,required" format:"date-time"`
	// Identifier of the Observation associated with this Correlation. If
	// corrType=TRACK then this field should reference the first Observation in the
	// track. Note: To retrieve all remaining Observations in the track, the GET query
	// should include this Observation's source and origin field values, along with the
	// trackId.
	ObID string `json:"obId,required"`
	// Indicator of whether the type of Observation(s) being correlated is DOA, EO,
	// PASSIVE_RADAR, RADAR, RF, SAR, or SOISET: DOA: The observation type being
	// correlated is Difference of Arrival. EO: The observation type being correlated
	// is Electro-Optical. PASSIVE_RADAR: The observation type being correlated is
	// Passive Radar. RADAR: The observation type being correlated is Radar. RF: The
	// observation type being correlated is Radio Frequency. SAR: The observation type
	// being correlated is Synthetic Aperture Radar. SOISET: The observation type being
	// correlated is Space Object Identification Observation Set.
	//
	// Any of "DOA", "EO", "PASSIVE_RADAR", "RADAR", "RF", "SAR", "SOISET".
	ObType ObservationObscorrelationHistoryListResponseObType `json:"obType,required"`
	// Identifier of the orbit state used for correlation.
	ReferenceOrbitID string `json:"referenceOrbitId,required"`
	// Indicator of whether the reference orbit type used for correlation is an ELSET,
	// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
	// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
	// State Vector.
	//
	// Any of "ELSET", "ESID", "SV".
	ReferenceOrbitType ObservationObscorrelationHistoryListResponseReferenceOrbitType `json:"referenceOrbitType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Type of algorithm used for this correlation (e.g. ROTAS, GEOMETRIC, STATISTICAL,
	// MAHALANOBIS, AI/ML, OTHER).
	AlgorithmCorrType string `json:"algorithmCorrType"`
	// Name of the alternate catalog.
	AltCatalog string `json:"altCatalog"`
	// Associates one or more alternate catalogs with a source provider or system.
	// Namespaces may be defined by their respective data providers or systems (e.g.
	// JCO, 18SDS, EOSSS, EXO, KBR, KRTL, LeoLabs, NorthStar, SAFRAN, Slingshot).
	AltNamespace string `json:"altNamespace"`
	// Alternate unique object ID within the namespace.
	AltObjectID string `json:"altObjectId"`
	// Boolean indicating whether the observation or track can be correlated to the
	// alternate object specified under altObjectId. This flag should only be set to
	// true by data providers after an attempt to correlate to an on-orbit object was
	// made and failed. If unable to correlate, the 'origObjectId' field may be
	// populated with an internal data provider specific identifier.
	AltUct bool `json:"altUct"`
	// Astrostandard ROTAS correlation result (0 - 4), if applicable. Refer to ROTAS
	// documentation for an explanation of ASTAT values.
	Astat int64 `json:"astat"`
	// Correlation score ranging from 0.0 to 1.0. A score of 1.0 represents perfect
	// correlation to the orbit of the corresponding satellite, such as when all
	// observation residuals equal 0.
	CorrQuality float64 `json:"corrQuality"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Identifier of the correlated target on-orbit object, if associated with a valid
	// satNo.
	IDOnOrbit string `json:"idOnOrbit"`
	// Identifier of the ObsCorrelation record from which this ObsCorrelation record
	// originated. This behavior allows for different source providers/systems to make
	// changes to a given correlation and maintain traceability back to the original
	// correlation.
	IDParentCorrelation string `json:"idParentCorrelation"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier indicates the target on-orbit object being correlated. This
	// may be an internal identifier and not necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Current 18th SDS satellite/catalog number of the target on-orbit object. Useful
	// to know in the case where an observation is correlated to another
	// satellite/catalog number.
	SatNo int64 `json:"satNo"`
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
	// Identifier of the Track associated with this ObsCorrelation.
	TrackID string `json:"trackId"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		CorrType              respjson.Field
		DataMode              respjson.Field
		MsgTs                 respjson.Field
		ObID                  respjson.Field
		ObType                respjson.Field
		ReferenceOrbitID      respjson.Field
		ReferenceOrbitType    respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AlgorithmCorrType     respjson.Field
		AltCatalog            respjson.Field
		AltNamespace          respjson.Field
		AltObjectID           respjson.Field
		AltUct                respjson.Field
		Astat                 respjson.Field
		CorrQuality           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDOnOrbit             respjson.Field
		IDParentCorrelation   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		SatNo                 respjson.Field
		SourceDl              respjson.Field
		Tags                  respjson.Field
		TrackID               respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObservationObscorrelationHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationObscorrelationHistoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the type of correlation is OBSERVATION or TRACK:
// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
// the entire track of observations is being correlated.
type ObservationObscorrelationHistoryListResponseCorrType string

const (
	ObservationObscorrelationHistoryListResponseCorrTypeObservation ObservationObscorrelationHistoryListResponseCorrType = "OBSERVATION"
	ObservationObscorrelationHistoryListResponseCorrTypeTrack       ObservationObscorrelationHistoryListResponseCorrType = "TRACK"
)

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
type ObservationObscorrelationHistoryListResponseDataMode string

const (
	ObservationObscorrelationHistoryListResponseDataModeReal      ObservationObscorrelationHistoryListResponseDataMode = "REAL"
	ObservationObscorrelationHistoryListResponseDataModeTest      ObservationObscorrelationHistoryListResponseDataMode = "TEST"
	ObservationObscorrelationHistoryListResponseDataModeSimulated ObservationObscorrelationHistoryListResponseDataMode = "SIMULATED"
	ObservationObscorrelationHistoryListResponseDataModeExercise  ObservationObscorrelationHistoryListResponseDataMode = "EXERCISE"
)

// Indicator of whether the type of Observation(s) being correlated is DOA, EO,
// PASSIVE_RADAR, RADAR, RF, SAR, or SOISET: DOA: The observation type being
// correlated is Difference of Arrival. EO: The observation type being correlated
// is Electro-Optical. PASSIVE_RADAR: The observation type being correlated is
// Passive Radar. RADAR: The observation type being correlated is Radar. RF: The
// observation type being correlated is Radio Frequency. SAR: The observation type
// being correlated is Synthetic Aperture Radar. SOISET: The observation type being
// correlated is Space Object Identification Observation Set.
type ObservationObscorrelationHistoryListResponseObType string

const (
	ObservationObscorrelationHistoryListResponseObTypeDoa          ObservationObscorrelationHistoryListResponseObType = "DOA"
	ObservationObscorrelationHistoryListResponseObTypeEo           ObservationObscorrelationHistoryListResponseObType = "EO"
	ObservationObscorrelationHistoryListResponseObTypePassiveRadar ObservationObscorrelationHistoryListResponseObType = "PASSIVE_RADAR"
	ObservationObscorrelationHistoryListResponseObTypeRadar        ObservationObscorrelationHistoryListResponseObType = "RADAR"
	ObservationObscorrelationHistoryListResponseObTypeRf           ObservationObscorrelationHistoryListResponseObType = "RF"
	ObservationObscorrelationHistoryListResponseObTypeSar          ObservationObscorrelationHistoryListResponseObType = "SAR"
	ObservationObscorrelationHistoryListResponseObTypeSoiset       ObservationObscorrelationHistoryListResponseObType = "SOISET"
)

// Indicator of whether the reference orbit type used for correlation is an ELSET,
// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
// State Vector.
type ObservationObscorrelationHistoryListResponseReferenceOrbitType string

const (
	ObservationObscorrelationHistoryListResponseReferenceOrbitTypeElset ObservationObscorrelationHistoryListResponseReferenceOrbitType = "ELSET"
	ObservationObscorrelationHistoryListResponseReferenceOrbitTypeEsid  ObservationObscorrelationHistoryListResponseReferenceOrbitType = "ESID"
	ObservationObscorrelationHistoryListResponseReferenceOrbitTypeSv    ObservationObscorrelationHistoryListResponseReferenceOrbitType = "SV"
)

type ObservationObscorrelationHistoryListParams struct {
	// Correlation message generation time, in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTs time.Time `query:"msgTs,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationObscorrelationHistoryListParams]'s query
// parameters as `url.Values`.
func (r ObservationObscorrelationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationObscorrelationHistoryAodrParams struct {
	// Correlation message generation time, in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTs time.Time `query:"msgTs,required" format:"date-time" json:"-"`
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

// URLQuery serializes [ObservationObscorrelationHistoryAodrParams]'s query
// parameters as `url.Values`.
func (r ObservationObscorrelationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationObscorrelationHistoryCountParams struct {
	// Correlation message generation time, in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTs       time.Time        `query:"msgTs,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationObscorrelationHistoryCountParams]'s query
// parameters as `url.Values`.
func (r ObservationObscorrelationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
