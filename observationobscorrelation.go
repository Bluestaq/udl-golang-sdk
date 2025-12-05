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

// ObservationObscorrelationService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationObscorrelationService] method instead.
type ObservationObscorrelationService struct {
	Options []option.RequestOption
	History ObservationObscorrelationHistoryService
}

// NewObservationObscorrelationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservationObscorrelationService(opts ...option.RequestOption) (r ObservationObscorrelationService) {
	r = ObservationObscorrelationService{}
	r.Options = opts
	r.History = NewObservationObscorrelationHistoryService(opts...)
	return
}

// Service operation to take a single ObsCorrelation record as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *ObservationObscorrelationService) New(ctx context.Context, body ObservationObscorrelationNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/obscorrelation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Correlation record by its unique ID passed as
// a path parameter.
func (r *ObservationObscorrelationService) Get(ctx context.Context, id string, query ObservationObscorrelationGetParams, opts ...option.RequestOption) (res *ObservationObscorrelationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/obscorrelation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationObscorrelationService) List(ctx context.Context, query ObservationObscorrelationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ObservationObscorrelationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/obscorrelation"
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
func (r *ObservationObscorrelationService) ListAutoPaging(ctx context.Context, query ObservationObscorrelationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ObservationObscorrelationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationObscorrelationService) Count(ctx context.Context, query ObservationObscorrelationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/obscorrelation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// ObsCorrelation records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *ObservationObscorrelationService) NewBulk(ctx context.Context, body ObservationObscorrelationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/obscorrelation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ObservationObscorrelationService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *ObservationObscorrelationQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/obscorrelation/queryhelp"
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
func (r *ObservationObscorrelationService) Tuple(ctx context.Context, query ObservationObscorrelationTupleParams, opts ...option.RequestOption) (res *[]ObservationObscorrelationTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/obscorrelation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple ObsCorrelation records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *ObservationObscorrelationService) UnvalidatedPublish(ctx context.Context, body ObservationObscorrelationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-obscorrelation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Model representation supporting post-pass correlation of UCTs and re-correlation
// of mis-tagged electro-optical (EO), radar, RF, and DOA track/observations.
type ObservationObscorrelationGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the type of correlation is OBSERVATION or TRACK:
	// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
	// the entire track of observations is being correlated.
	//
	// Any of "OBSERVATION", "TRACK".
	CorrType ObservationObscorrelationGetResponseCorrType `json:"corrType,required"`
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
	DataMode ObservationObscorrelationGetResponseDataMode `json:"dataMode,required"`
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
	ObType ObservationObscorrelationGetResponseObType `json:"obType,required"`
	// Identifier of the orbit state used for correlation.
	ReferenceOrbitID string `json:"referenceOrbitId,required"`
	// Indicator of whether the reference orbit type used for correlation is an ELSET,
	// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
	// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
	// State Vector.
	//
	// Any of "ELSET", "ESID", "SV".
	ReferenceOrbitType ObservationObscorrelationGetResponseReferenceOrbitType `json:"referenceOrbitType,required"`
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
func (r ObservationObscorrelationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationObscorrelationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the type of correlation is OBSERVATION or TRACK:
// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
// the entire track of observations is being correlated.
type ObservationObscorrelationGetResponseCorrType string

const (
	ObservationObscorrelationGetResponseCorrTypeObservation ObservationObscorrelationGetResponseCorrType = "OBSERVATION"
	ObservationObscorrelationGetResponseCorrTypeTrack       ObservationObscorrelationGetResponseCorrType = "TRACK"
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
type ObservationObscorrelationGetResponseDataMode string

const (
	ObservationObscorrelationGetResponseDataModeReal      ObservationObscorrelationGetResponseDataMode = "REAL"
	ObservationObscorrelationGetResponseDataModeTest      ObservationObscorrelationGetResponseDataMode = "TEST"
	ObservationObscorrelationGetResponseDataModeSimulated ObservationObscorrelationGetResponseDataMode = "SIMULATED"
	ObservationObscorrelationGetResponseDataModeExercise  ObservationObscorrelationGetResponseDataMode = "EXERCISE"
)

// Indicator of whether the type of Observation(s) being correlated is DOA, EO,
// PASSIVE_RADAR, RADAR, RF, SAR, or SOISET: DOA: The observation type being
// correlated is Difference of Arrival. EO: The observation type being correlated
// is Electro-Optical. PASSIVE_RADAR: The observation type being correlated is
// Passive Radar. RADAR: The observation type being correlated is Radar. RF: The
// observation type being correlated is Radio Frequency. SAR: The observation type
// being correlated is Synthetic Aperture Radar. SOISET: The observation type being
// correlated is Space Object Identification Observation Set.
type ObservationObscorrelationGetResponseObType string

const (
	ObservationObscorrelationGetResponseObTypeDoa          ObservationObscorrelationGetResponseObType = "DOA"
	ObservationObscorrelationGetResponseObTypeEo           ObservationObscorrelationGetResponseObType = "EO"
	ObservationObscorrelationGetResponseObTypePassiveRadar ObservationObscorrelationGetResponseObType = "PASSIVE_RADAR"
	ObservationObscorrelationGetResponseObTypeRadar        ObservationObscorrelationGetResponseObType = "RADAR"
	ObservationObscorrelationGetResponseObTypeRf           ObservationObscorrelationGetResponseObType = "RF"
	ObservationObscorrelationGetResponseObTypeSar          ObservationObscorrelationGetResponseObType = "SAR"
	ObservationObscorrelationGetResponseObTypeSoiset       ObservationObscorrelationGetResponseObType = "SOISET"
)

// Indicator of whether the reference orbit type used for correlation is an ELSET,
// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
// State Vector.
type ObservationObscorrelationGetResponseReferenceOrbitType string

const (
	ObservationObscorrelationGetResponseReferenceOrbitTypeElset ObservationObscorrelationGetResponseReferenceOrbitType = "ELSET"
	ObservationObscorrelationGetResponseReferenceOrbitTypeEsid  ObservationObscorrelationGetResponseReferenceOrbitType = "ESID"
	ObservationObscorrelationGetResponseReferenceOrbitTypeSv    ObservationObscorrelationGetResponseReferenceOrbitType = "SV"
)

// Model representation supporting post-pass correlation of UCTs and re-correlation
// of mis-tagged electro-optical (EO), radar, RF, and DOA track/observations.
type ObservationObscorrelationListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the type of correlation is OBSERVATION or TRACK:
	// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
	// the entire track of observations is being correlated.
	//
	// Any of "OBSERVATION", "TRACK".
	CorrType ObservationObscorrelationListResponseCorrType `json:"corrType,required"`
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
	DataMode ObservationObscorrelationListResponseDataMode `json:"dataMode,required"`
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
	ObType ObservationObscorrelationListResponseObType `json:"obType,required"`
	// Identifier of the orbit state used for correlation.
	ReferenceOrbitID string `json:"referenceOrbitId,required"`
	// Indicator of whether the reference orbit type used for correlation is an ELSET,
	// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
	// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
	// State Vector.
	//
	// Any of "ELSET", "ESID", "SV".
	ReferenceOrbitType ObservationObscorrelationListResponseReferenceOrbitType `json:"referenceOrbitType,required"`
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
		TrackID               respjson.Field
		TransactionID         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObservationObscorrelationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationObscorrelationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the type of correlation is OBSERVATION or TRACK:
// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
// the entire track of observations is being correlated.
type ObservationObscorrelationListResponseCorrType string

const (
	ObservationObscorrelationListResponseCorrTypeObservation ObservationObscorrelationListResponseCorrType = "OBSERVATION"
	ObservationObscorrelationListResponseCorrTypeTrack       ObservationObscorrelationListResponseCorrType = "TRACK"
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
type ObservationObscorrelationListResponseDataMode string

const (
	ObservationObscorrelationListResponseDataModeReal      ObservationObscorrelationListResponseDataMode = "REAL"
	ObservationObscorrelationListResponseDataModeTest      ObservationObscorrelationListResponseDataMode = "TEST"
	ObservationObscorrelationListResponseDataModeSimulated ObservationObscorrelationListResponseDataMode = "SIMULATED"
	ObservationObscorrelationListResponseDataModeExercise  ObservationObscorrelationListResponseDataMode = "EXERCISE"
)

// Indicator of whether the type of Observation(s) being correlated is DOA, EO,
// PASSIVE_RADAR, RADAR, RF, SAR, or SOISET: DOA: The observation type being
// correlated is Difference of Arrival. EO: The observation type being correlated
// is Electro-Optical. PASSIVE_RADAR: The observation type being correlated is
// Passive Radar. RADAR: The observation type being correlated is Radar. RF: The
// observation type being correlated is Radio Frequency. SAR: The observation type
// being correlated is Synthetic Aperture Radar. SOISET: The observation type being
// correlated is Space Object Identification Observation Set.
type ObservationObscorrelationListResponseObType string

const (
	ObservationObscorrelationListResponseObTypeDoa          ObservationObscorrelationListResponseObType = "DOA"
	ObservationObscorrelationListResponseObTypeEo           ObservationObscorrelationListResponseObType = "EO"
	ObservationObscorrelationListResponseObTypePassiveRadar ObservationObscorrelationListResponseObType = "PASSIVE_RADAR"
	ObservationObscorrelationListResponseObTypeRadar        ObservationObscorrelationListResponseObType = "RADAR"
	ObservationObscorrelationListResponseObTypeRf           ObservationObscorrelationListResponseObType = "RF"
	ObservationObscorrelationListResponseObTypeSar          ObservationObscorrelationListResponseObType = "SAR"
	ObservationObscorrelationListResponseObTypeSoiset       ObservationObscorrelationListResponseObType = "SOISET"
)

// Indicator of whether the reference orbit type used for correlation is an ELSET,
// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
// State Vector.
type ObservationObscorrelationListResponseReferenceOrbitType string

const (
	ObservationObscorrelationListResponseReferenceOrbitTypeElset ObservationObscorrelationListResponseReferenceOrbitType = "ELSET"
	ObservationObscorrelationListResponseReferenceOrbitTypeEsid  ObservationObscorrelationListResponseReferenceOrbitType = "ESID"
	ObservationObscorrelationListResponseReferenceOrbitTypeSv    ObservationObscorrelationListResponseReferenceOrbitType = "SV"
)

type ObservationObscorrelationQueryHelpResponse struct {
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
func (r ObservationObscorrelationQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationObscorrelationQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation supporting post-pass correlation of UCTs and re-correlation
// of mis-tagged electro-optical (EO), radar, RF, and DOA track/observations.
type ObservationObscorrelationTupleResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the type of correlation is OBSERVATION or TRACK:
	// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
	// the entire track of observations is being correlated.
	//
	// Any of "OBSERVATION", "TRACK".
	CorrType ObservationObscorrelationTupleResponseCorrType `json:"corrType,required"`
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
	DataMode ObservationObscorrelationTupleResponseDataMode `json:"dataMode,required"`
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
	ObType ObservationObscorrelationTupleResponseObType `json:"obType,required"`
	// Identifier of the orbit state used for correlation.
	ReferenceOrbitID string `json:"referenceOrbitId,required"`
	// Indicator of whether the reference orbit type used for correlation is an ELSET,
	// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
	// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
	// State Vector.
	//
	// Any of "ELSET", "ESID", "SV".
	ReferenceOrbitType ObservationObscorrelationTupleResponseReferenceOrbitType `json:"referenceOrbitType,required"`
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
func (r ObservationObscorrelationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationObscorrelationTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the type of correlation is OBSERVATION or TRACK:
// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
// the entire track of observations is being correlated.
type ObservationObscorrelationTupleResponseCorrType string

const (
	ObservationObscorrelationTupleResponseCorrTypeObservation ObservationObscorrelationTupleResponseCorrType = "OBSERVATION"
	ObservationObscorrelationTupleResponseCorrTypeTrack       ObservationObscorrelationTupleResponseCorrType = "TRACK"
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
type ObservationObscorrelationTupleResponseDataMode string

const (
	ObservationObscorrelationTupleResponseDataModeReal      ObservationObscorrelationTupleResponseDataMode = "REAL"
	ObservationObscorrelationTupleResponseDataModeTest      ObservationObscorrelationTupleResponseDataMode = "TEST"
	ObservationObscorrelationTupleResponseDataModeSimulated ObservationObscorrelationTupleResponseDataMode = "SIMULATED"
	ObservationObscorrelationTupleResponseDataModeExercise  ObservationObscorrelationTupleResponseDataMode = "EXERCISE"
)

// Indicator of whether the type of Observation(s) being correlated is DOA, EO,
// PASSIVE_RADAR, RADAR, RF, SAR, or SOISET: DOA: The observation type being
// correlated is Difference of Arrival. EO: The observation type being correlated
// is Electro-Optical. PASSIVE_RADAR: The observation type being correlated is
// Passive Radar. RADAR: The observation type being correlated is Radar. RF: The
// observation type being correlated is Radio Frequency. SAR: The observation type
// being correlated is Synthetic Aperture Radar. SOISET: The observation type being
// correlated is Space Object Identification Observation Set.
type ObservationObscorrelationTupleResponseObType string

const (
	ObservationObscorrelationTupleResponseObTypeDoa          ObservationObscorrelationTupleResponseObType = "DOA"
	ObservationObscorrelationTupleResponseObTypeEo           ObservationObscorrelationTupleResponseObType = "EO"
	ObservationObscorrelationTupleResponseObTypePassiveRadar ObservationObscorrelationTupleResponseObType = "PASSIVE_RADAR"
	ObservationObscorrelationTupleResponseObTypeRadar        ObservationObscorrelationTupleResponseObType = "RADAR"
	ObservationObscorrelationTupleResponseObTypeRf           ObservationObscorrelationTupleResponseObType = "RF"
	ObservationObscorrelationTupleResponseObTypeSar          ObservationObscorrelationTupleResponseObType = "SAR"
	ObservationObscorrelationTupleResponseObTypeSoiset       ObservationObscorrelationTupleResponseObType = "SOISET"
)

// Indicator of whether the reference orbit type used for correlation is an ELSET,
// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
// State Vector.
type ObservationObscorrelationTupleResponseReferenceOrbitType string

const (
	ObservationObscorrelationTupleResponseReferenceOrbitTypeElset ObservationObscorrelationTupleResponseReferenceOrbitType = "ELSET"
	ObservationObscorrelationTupleResponseReferenceOrbitTypeEsid  ObservationObscorrelationTupleResponseReferenceOrbitType = "ESID"
	ObservationObscorrelationTupleResponseReferenceOrbitTypeSv    ObservationObscorrelationTupleResponseReferenceOrbitType = "SV"
)

type ObservationObscorrelationNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the type of correlation is OBSERVATION or TRACK:
	// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
	// the entire track of observations is being correlated.
	//
	// Any of "OBSERVATION", "TRACK".
	CorrType ObservationObscorrelationNewParamsCorrType `json:"corrType,omitzero,required"`
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
	DataMode ObservationObscorrelationNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	ObType ObservationObscorrelationNewParamsObType `json:"obType,omitzero,required"`
	// Identifier of the orbit state used for correlation.
	ReferenceOrbitID string `json:"referenceOrbitId,required"`
	// Indicator of whether the reference orbit type used for correlation is an ELSET,
	// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
	// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
	// State Vector.
	//
	// Any of "ELSET", "ESID", "SV".
	ReferenceOrbitType ObservationObscorrelationNewParamsReferenceOrbitType `json:"referenceOrbitType,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Type of algorithm used for this correlation (e.g. ROTAS, GEOMETRIC, STATISTICAL,
	// MAHALANOBIS, AI/ML, OTHER).
	AlgorithmCorrType param.Opt[string] `json:"algorithmCorrType,omitzero"`
	// Name of the alternate catalog.
	AltCatalog param.Opt[string] `json:"altCatalog,omitzero"`
	// Associates one or more alternate catalogs with a source provider or system.
	// Namespaces may be defined by their respective data providers or systems (e.g.
	// JCO, 18SDS, EOSSS, EXO, KBR, KRTL, LeoLabs, NorthStar, SAFRAN, Slingshot).
	AltNamespace param.Opt[string] `json:"altNamespace,omitzero"`
	// Alternate unique object ID within the namespace.
	AltObjectID param.Opt[string] `json:"altObjectId,omitzero"`
	// Boolean indicating whether the observation or track can be correlated to the
	// alternate object specified under altObjectId. This flag should only be set to
	// true by data providers after an attempt to correlate to an on-orbit object was
	// made and failed. If unable to correlate, the 'origObjectId' field may be
	// populated with an internal data provider specific identifier.
	AltUct param.Opt[bool] `json:"altUct,omitzero"`
	// Astrostandard ROTAS correlation result (0 - 4), if applicable. Refer to ROTAS
	// documentation for an explanation of ASTAT values.
	Astat param.Opt[int64] `json:"astat,omitzero"`
	// Correlation score ranging from 0.0 to 1.0. A score of 1.0 represents perfect
	// correlation to the orbit of the corresponding satellite, such as when all
	// observation residuals equal 0.
	CorrQuality param.Opt[float64] `json:"corrQuality,omitzero"`
	// Identifier of the ObsCorrelation record from which this ObsCorrelation record
	// originated. This behavior allows for different source providers/systems to make
	// changes to a given correlation and maintain traceability back to the original
	// correlation.
	IDParentCorrelation param.Opt[string] `json:"idParentCorrelation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier indicates the target on-orbit object being correlated. This
	// may be an internal identifier and not necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Current 18th SDS satellite/catalog number of the target on-orbit object. Useful
	// to know in the case where an observation is correlated to another
	// satellite/catalog number.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Identifier of the Track associated with this ObsCorrelation.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ObservationObscorrelationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObservationObscorrelationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationObscorrelationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the type of correlation is OBSERVATION or TRACK:
// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
// the entire track of observations is being correlated.
type ObservationObscorrelationNewParamsCorrType string

const (
	ObservationObscorrelationNewParamsCorrTypeObservation ObservationObscorrelationNewParamsCorrType = "OBSERVATION"
	ObservationObscorrelationNewParamsCorrTypeTrack       ObservationObscorrelationNewParamsCorrType = "TRACK"
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
type ObservationObscorrelationNewParamsDataMode string

const (
	ObservationObscorrelationNewParamsDataModeReal      ObservationObscorrelationNewParamsDataMode = "REAL"
	ObservationObscorrelationNewParamsDataModeTest      ObservationObscorrelationNewParamsDataMode = "TEST"
	ObservationObscorrelationNewParamsDataModeSimulated ObservationObscorrelationNewParamsDataMode = "SIMULATED"
	ObservationObscorrelationNewParamsDataModeExercise  ObservationObscorrelationNewParamsDataMode = "EXERCISE"
)

// Indicator of whether the type of Observation(s) being correlated is DOA, EO,
// PASSIVE_RADAR, RADAR, RF, SAR, or SOISET: DOA: The observation type being
// correlated is Difference of Arrival. EO: The observation type being correlated
// is Electro-Optical. PASSIVE_RADAR: The observation type being correlated is
// Passive Radar. RADAR: The observation type being correlated is Radar. RF: The
// observation type being correlated is Radio Frequency. SAR: The observation type
// being correlated is Synthetic Aperture Radar. SOISET: The observation type being
// correlated is Space Object Identification Observation Set.
type ObservationObscorrelationNewParamsObType string

const (
	ObservationObscorrelationNewParamsObTypeDoa          ObservationObscorrelationNewParamsObType = "DOA"
	ObservationObscorrelationNewParamsObTypeEo           ObservationObscorrelationNewParamsObType = "EO"
	ObservationObscorrelationNewParamsObTypePassiveRadar ObservationObscorrelationNewParamsObType = "PASSIVE_RADAR"
	ObservationObscorrelationNewParamsObTypeRadar        ObservationObscorrelationNewParamsObType = "RADAR"
	ObservationObscorrelationNewParamsObTypeRf           ObservationObscorrelationNewParamsObType = "RF"
	ObservationObscorrelationNewParamsObTypeSar          ObservationObscorrelationNewParamsObType = "SAR"
	ObservationObscorrelationNewParamsObTypeSoiset       ObservationObscorrelationNewParamsObType = "SOISET"
)

// Indicator of whether the reference orbit type used for correlation is an ELSET,
// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
// State Vector.
type ObservationObscorrelationNewParamsReferenceOrbitType string

const (
	ObservationObscorrelationNewParamsReferenceOrbitTypeElset ObservationObscorrelationNewParamsReferenceOrbitType = "ELSET"
	ObservationObscorrelationNewParamsReferenceOrbitTypeEsid  ObservationObscorrelationNewParamsReferenceOrbitType = "ESID"
	ObservationObscorrelationNewParamsReferenceOrbitTypeSv    ObservationObscorrelationNewParamsReferenceOrbitType = "SV"
)

type ObservationObscorrelationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationObscorrelationGetParams]'s query parameters as
// `url.Values`.
func (r ObservationObscorrelationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationObscorrelationListParams struct {
	// Correlation message generation time, in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTs       time.Time        `query:"msgTs,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationObscorrelationListParams]'s query parameters as
// `url.Values`.
func (r ObservationObscorrelationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationObscorrelationCountParams struct {
	// Correlation message generation time, in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTs       time.Time        `query:"msgTs,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationObscorrelationCountParams]'s query parameters as
// `url.Values`.
func (r ObservationObscorrelationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationObscorrelationNewBulkParams struct {
	Body []ObservationObscorrelationNewBulkParamsBody
	paramObj
}

func (r ObservationObscorrelationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationObscorrelationNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Model representation supporting post-pass correlation of UCTs and re-correlation
// of mis-tagged electro-optical (EO), radar, RF, and DOA track/observations.
//
// The properties ClassificationMarking, CorrType, DataMode, MsgTs, ObID, ObType,
// ReferenceOrbitID, ReferenceOrbitType, Source are required.
type ObservationObscorrelationNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the type of correlation is OBSERVATION or TRACK:
	// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
	// the entire track of observations is being correlated.
	//
	// Any of "OBSERVATION", "TRACK".
	CorrType string `json:"corrType,omitzero,required"`
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
	ObType string `json:"obType,omitzero,required"`
	// Identifier of the orbit state used for correlation.
	ReferenceOrbitID string `json:"referenceOrbitId,required"`
	// Indicator of whether the reference orbit type used for correlation is an ELSET,
	// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
	// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
	// State Vector.
	//
	// Any of "ELSET", "ESID", "SV".
	ReferenceOrbitType string `json:"referenceOrbitType,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Type of algorithm used for this correlation (e.g. ROTAS, GEOMETRIC, STATISTICAL,
	// MAHALANOBIS, AI/ML, OTHER).
	AlgorithmCorrType param.Opt[string] `json:"algorithmCorrType,omitzero"`
	// Name of the alternate catalog.
	AltCatalog param.Opt[string] `json:"altCatalog,omitzero"`
	// Associates one or more alternate catalogs with a source provider or system.
	// Namespaces may be defined by their respective data providers or systems (e.g.
	// JCO, 18SDS, EOSSS, EXO, KBR, KRTL, LeoLabs, NorthStar, SAFRAN, Slingshot).
	AltNamespace param.Opt[string] `json:"altNamespace,omitzero"`
	// Alternate unique object ID within the namespace.
	AltObjectID param.Opt[string] `json:"altObjectId,omitzero"`
	// Boolean indicating whether the observation or track can be correlated to the
	// alternate object specified under altObjectId. This flag should only be set to
	// true by data providers after an attempt to correlate to an on-orbit object was
	// made and failed. If unable to correlate, the 'origObjectId' field may be
	// populated with an internal data provider specific identifier.
	AltUct param.Opt[bool] `json:"altUct,omitzero"`
	// Astrostandard ROTAS correlation result (0 - 4), if applicable. Refer to ROTAS
	// documentation for an explanation of ASTAT values.
	Astat param.Opt[int64] `json:"astat,omitzero"`
	// Correlation score ranging from 0.0 to 1.0. A score of 1.0 represents perfect
	// correlation to the orbit of the corresponding satellite, such as when all
	// observation residuals equal 0.
	CorrQuality param.Opt[float64] `json:"corrQuality,omitzero"`
	// Identifier of the ObsCorrelation record from which this ObsCorrelation record
	// originated. This behavior allows for different source providers/systems to make
	// changes to a given correlation and maintain traceability back to the original
	// correlation.
	IDParentCorrelation param.Opt[string] `json:"idParentCorrelation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier indicates the target on-orbit object being correlated. This
	// may be an internal identifier and not necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Current 18th SDS satellite/catalog number of the target on-orbit object. Useful
	// to know in the case where an observation is correlated to another
	// satellite/catalog number.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Identifier of the Track associated with this ObsCorrelation.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ObservationObscorrelationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationObscorrelationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationObscorrelationNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationObscorrelationNewBulkParamsBody](
		"corrType", "OBSERVATION", "TRACK",
	)
	apijson.RegisterFieldValidator[ObservationObscorrelationNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ObservationObscorrelationNewBulkParamsBody](
		"obType", "DOA", "EO", "PASSIVE_RADAR", "RADAR", "RF", "SAR", "SOISET",
	)
	apijson.RegisterFieldValidator[ObservationObscorrelationNewBulkParamsBody](
		"referenceOrbitType", "ELSET", "ESID", "SV",
	)
}

type ObservationObscorrelationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Correlation message generation time, in ISO 8601 UTC format with millisecond
	// precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	MsgTs       time.Time        `query:"msgTs,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationObscorrelationTupleParams]'s query parameters as
// `url.Values`.
func (r ObservationObscorrelationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationObscorrelationUnvalidatedPublishParams struct {
	Body []ObservationObscorrelationUnvalidatedPublishParamsBody
	paramObj
}

func (r ObservationObscorrelationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationObscorrelationUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Model representation supporting post-pass correlation of UCTs and re-correlation
// of mis-tagged electro-optical (EO), radar, RF, and DOA track/observations.
//
// The properties ClassificationMarking, CorrType, DataMode, MsgTs, ObID, ObType,
// ReferenceOrbitID, ReferenceOrbitType, Source are required.
type ObservationObscorrelationUnvalidatedPublishParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the type of correlation is OBSERVATION or TRACK:
	// OBSERVATION: Identifies an observation is being correlated. TRACK: Identifies
	// the entire track of observations is being correlated.
	//
	// Any of "OBSERVATION", "TRACK".
	CorrType string `json:"corrType,omitzero,required"`
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
	ObType string `json:"obType,omitzero,required"`
	// Identifier of the orbit state used for correlation.
	ReferenceOrbitID string `json:"referenceOrbitId,required"`
	// Indicator of whether the reference orbit type used for correlation is an ELSET,
	// ESID, or SV: ELSET: The reference orbit type is an Element Set. ESID: The
	// reference orbit type is an Ephemeris Set. SV: The reference orbit type is a
	// State Vector.
	//
	// Any of "ELSET", "ESID", "SV".
	ReferenceOrbitType string `json:"referenceOrbitType,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Type of algorithm used for this correlation (e.g. ROTAS, GEOMETRIC, STATISTICAL,
	// MAHALANOBIS, AI/ML, OTHER).
	AlgorithmCorrType param.Opt[string] `json:"algorithmCorrType,omitzero"`
	// Name of the alternate catalog.
	AltCatalog param.Opt[string] `json:"altCatalog,omitzero"`
	// Associates one or more alternate catalogs with a source provider or system.
	// Namespaces may be defined by their respective data providers or systems (e.g.
	// JCO, 18SDS, EOSSS, EXO, KBR, KRTL, LeoLabs, NorthStar, SAFRAN, Slingshot).
	AltNamespace param.Opt[string] `json:"altNamespace,omitzero"`
	// Alternate unique object ID within the namespace.
	AltObjectID param.Opt[string] `json:"altObjectId,omitzero"`
	// Boolean indicating whether the observation or track can be correlated to the
	// alternate object specified under altObjectId. This flag should only be set to
	// true by data providers after an attempt to correlate to an on-orbit object was
	// made and failed. If unable to correlate, the 'origObjectId' field may be
	// populated with an internal data provider specific identifier.
	AltUct param.Opt[bool] `json:"altUct,omitzero"`
	// Astrostandard ROTAS correlation result (0 - 4), if applicable. Refer to ROTAS
	// documentation for an explanation of ASTAT values.
	Astat param.Opt[int64] `json:"astat,omitzero"`
	// Correlation score ranging from 0.0 to 1.0. A score of 1.0 represents perfect
	// correlation to the orbit of the corresponding satellite, such as when all
	// observation residuals equal 0.
	CorrQuality param.Opt[float64] `json:"corrQuality,omitzero"`
	// Identifier of the ObsCorrelation record from which this ObsCorrelation record
	// originated. This behavior allows for different source providers/systems to make
	// changes to a given correlation and maintain traceability back to the original
	// correlation.
	IDParentCorrelation param.Opt[string] `json:"idParentCorrelation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier indicates the target on-orbit object being correlated. This
	// may be an internal identifier and not necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Current 18th SDS satellite/catalog number of the target on-orbit object. Useful
	// to know in the case where an observation is correlated to another
	// satellite/catalog number.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Identifier of the Track associated with this ObsCorrelation.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ObservationObscorrelationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationObscorrelationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationObscorrelationUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationObscorrelationUnvalidatedPublishParamsBody](
		"corrType", "OBSERVATION", "TRACK",
	)
	apijson.RegisterFieldValidator[ObservationObscorrelationUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[ObservationObscorrelationUnvalidatedPublishParamsBody](
		"obType", "DOA", "EO", "PASSIVE_RADAR", "RADAR", "RF", "SAR", "SOISET",
	)
	apijson.RegisterFieldValidator[ObservationObscorrelationUnvalidatedPublishParamsBody](
		"referenceOrbitType", "ELSET", "ESID", "SV",
	)
}
