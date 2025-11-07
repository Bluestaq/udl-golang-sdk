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

// EventEvolutionService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventEvolutionService] method instead.
type EventEvolutionService struct {
	Options []option.RequestOption
	History EventEvolutionHistoryService
}

// NewEventEvolutionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventEvolutionService(opts ...option.RequestOption) (r EventEvolutionService) {
	r = EventEvolutionService{}
	r.Options = opts
	r.History = NewEventEvolutionHistoryService(opts...)
	return
}

// Service operation to take a single EventEvolution object as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *EventEvolutionService) New(ctx context.Context, body EventEvolutionNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/eventevolution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single EventEvolution by its unique ID passed as a
// path parameter.
func (r *EventEvolutionService) Get(ctx context.Context, id string, query EventEvolutionGetParams, opts ...option.RequestOption) (res *shared.EventEvolutionFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/eventevolution/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EventEvolutionService) List(ctx context.Context, query EventEvolutionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EventEvolutionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/eventevolution"
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
func (r *EventEvolutionService) ListAutoPaging(ctx context.Context, query EventEvolutionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EventEvolutionListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EventEvolutionService) Count(ctx context.Context, query EventEvolutionCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/eventevolution/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// EventEvolution records as a POST body and ingest into the database. Requires
// specific roles, please contact the UDL team to gain access. This operation is
// not intended to be used for automated feeds into UDL...data providers should
// contact the UDL team for instructions on setting up a permanent feed through an
// alternate mechanism.
func (r *EventEvolutionService) NewBulk(ctx context.Context, body EventEvolutionNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/eventevolution/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EventEvolutionService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *EventEvolutionQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/eventevolution/queryhelp"
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
func (r *EventEvolutionService) Tuple(ctx context.Context, query EventEvolutionTupleParams, opts ...option.RequestOption) (res *[]shared.EventEvolutionFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/eventevolution/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a list of EventEvolution records as a POST body and
// ingest into the database. Requires a specific role, please contact the UDL team
// to gain access. This operation is intended to be used for automated feeds into
// UDL.
func (r *EventEvolutionService) UnvalidatedPublish(ctx context.Context, body EventEvolutionUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-eventevolution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Event Evolution is a unique service supporting the association of UDL records of
// various data types to a common event or activity. The associations may be a one
// time summary, aggregating sources of a past event, or of an ongoing activity
// that evolves over a period of time.
type EventEvolutionListResponse struct {
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
	DataMode EventEvolutionListResponseDataMode `json:"dataMode,required"`
	// User-provided unique identifier of this activity or event. This ID should remain
	// the same on subsequent updates in order to associate all records pertaining to
	// the activity or event.
	EventID string `json:"eventId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual or estimated start time of the activity or event, in ISO 8601 UTC
	// format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Summary or description of the activity or event.
	Summary string `json:"summary,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground.
	Atype string `json:"atype"`
	// The activity or event type associated with this record (e.g. BREAKUP, DIRECT
	// FIRE, IED, LAUNCH, PROTEST, etc.). For Significant Activities, recommended but
	// not constrained to, CAMEO.Manual.1.1b3 Chapter 6. Note that the evolution of an
	// event may incorporate records of various types, for example, a LAUNCH event may
	// evolve into a BREAKUP event.
	Category string `json:"category"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional description of the relationship between the records provided in the
	// srcTyps/srcIds and the activity or event.
	DataDescription string `json:"dataDescription"`
	// The actual or estimated start time of the activity or event, in ISO 8601 UTC
	// format.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Administrative boundaries of the first sub-national level. Level 1 is simply the
	// largest demarcation under whatever demarcation criteria has been determined by
	// the governing body. For example this may be the state/province in which a
	// terrestrial event takes place, or with which the event is attributed for
	// non-localized or non-terrestrial activity.
	GeoAdminLevel1 string `json:"geoAdminLevel1"`
	// Administrative boundaries of the second sub-national level. Level 2 is simply
	// the second largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example this may be the city/district in
	// which a terrestrial event takes place, or with which the event is attributed for
	// non-localized or non-terrestrial activity.
	GeoAdminLevel2 string `json:"geoAdminLevel2"`
	// Administrative boundaries of the third sub-national level. Level 3 is simply the
	// third largest demarcation under whatever demarcation criteria has been
	// determined by the governing body.
	GeoAdminLevel3 string `json:"geoAdminLevel3"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Flag indicating that this record is for the purpose of redacting one or more
	// previously specified records from association with this activity or event. If
	// this flag is set then all records indicated in srcTyps/srcIds should be removed
	// from event association.
	Redact bool `json:"redact"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this activity or event. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size. See the corresponding srcTyps
	// array element for the data type of the UUID and use the appropriate API
	// operation to retrieve that object.
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK) that
	// are related to this activity or event. See the associated 'srcIds' array for the
	// record UUIDs, positionally corresponding to the record types in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// The status of this activity or event. (ACTIVE, CONCLUDED, UNKNOWN).
	Status string `json:"status"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// List of URLs to before/after images of this point of interest entity.
	URL []string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EventID               respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		Summary               respjson.Field
		ID                    respjson.Field
		Agjson                respjson.Field
		Andims                respjson.Field
		Asrid                 respjson.Field
		Atext                 respjson.Field
		Atype                 respjson.Field
		Category              respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataDescription       respjson.Field
		EndTime               respjson.Field
		GeoAdminLevel1        respjson.Field
		GeoAdminLevel2        respjson.Field
		GeoAdminLevel3        respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Redact                respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		Status                respjson.Field
		Tags                  respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventEvolutionListResponse) RawJSON() string { return r.JSON.raw }
func (r *EventEvolutionListResponse) UnmarshalJSON(data []byte) error {
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
type EventEvolutionListResponseDataMode string

const (
	EventEvolutionListResponseDataModeReal      EventEvolutionListResponseDataMode = "REAL"
	EventEvolutionListResponseDataModeTest      EventEvolutionListResponseDataMode = "TEST"
	EventEvolutionListResponseDataModeSimulated EventEvolutionListResponseDataMode = "SIMULATED"
	EventEvolutionListResponseDataModeExercise  EventEvolutionListResponseDataMode = "EXERCISE"
)

type EventEvolutionQueryhelpResponse struct {
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
func (r EventEvolutionQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *EventEvolutionQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventEvolutionNewParams struct {
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
	DataMode EventEvolutionNewParamsDataMode `json:"dataMode,omitzero,required"`
	// User-provided unique identifier of this activity or event. This ID should remain
	// the same on subsequent updates in order to associate all records pertaining to
	// the activity or event.
	EventID string `json:"eventId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual or estimated start time of the activity or event, in ISO 8601 UTC
	// format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Summary or description of the activity or event.
	Summary string `json:"summary,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the point of interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// The activity or event type associated with this record (e.g. BREAKUP, DIRECT
	// FIRE, IED, LAUNCH, PROTEST, etc.). For Significant Activities, recommended but
	// not constrained to, CAMEO.Manual.1.1b3 Chapter 6. Note that the evolution of an
	// event may incorporate records of various types, for example, a LAUNCH event may
	// evolve into a BREAKUP event.
	Category param.Opt[string] `json:"category,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Optional description of the relationship between the records provided in the
	// srcTyps/srcIds and the activity or event.
	DataDescription param.Opt[string] `json:"dataDescription,omitzero"`
	// The actual or estimated start time of the activity or event, in ISO 8601 UTC
	// format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Administrative boundaries of the first sub-national level. Level 1 is simply the
	// largest demarcation under whatever demarcation criteria has been determined by
	// the governing body. For example this may be the state/province in which a
	// terrestrial event takes place, or with which the event is attributed for
	// non-localized or non-terrestrial activity.
	GeoAdminLevel1 param.Opt[string] `json:"geoAdminLevel1,omitzero"`
	// Administrative boundaries of the second sub-national level. Level 2 is simply
	// the second largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example this may be the city/district in
	// which a terrestrial event takes place, or with which the event is attributed for
	// non-localized or non-terrestrial activity.
	GeoAdminLevel2 param.Opt[string] `json:"geoAdminLevel2,omitzero"`
	// Administrative boundaries of the third sub-national level. Level 3 is simply the
	// third largest demarcation under whatever demarcation criteria has been
	// determined by the governing body.
	GeoAdminLevel3 param.Opt[string] `json:"geoAdminLevel3,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Flag indicating that this record is for the purpose of redacting one or more
	// previously specified records from association with this activity or event. If
	// this flag is set then all records indicated in srcTyps/srcIds should be removed
	// from event association.
	Redact param.Opt[bool] `json:"redact,omitzero"`
	// The status of this activity or event. (ACTIVE, CONCLUDED, UNKNOWN).
	Status param.Opt[string] `json:"status,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this activity or event. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size. See the corresponding srcTyps
	// array element for the data type of the UUID and use the appropriate API
	// operation to retrieve that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK) that
	// are related to this activity or event. See the associated 'srcIds' array for the
	// record UUIDs, positionally corresponding to the record types in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// List of URLs to before/after images of this point of interest entity.
	URL []string `json:"url,omitzero"`
	paramObj
}

func (r EventEvolutionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EventEvolutionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventEvolutionNewParams) UnmarshalJSON(data []byte) error {
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
type EventEvolutionNewParamsDataMode string

const (
	EventEvolutionNewParamsDataModeReal      EventEvolutionNewParamsDataMode = "REAL"
	EventEvolutionNewParamsDataModeTest      EventEvolutionNewParamsDataMode = "TEST"
	EventEvolutionNewParamsDataModeSimulated EventEvolutionNewParamsDataMode = "SIMULATED"
	EventEvolutionNewParamsDataModeExercise  EventEvolutionNewParamsDataMode = "EXERCISE"
)

type EventEvolutionGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventEvolutionGetParams]'s query parameters as
// `url.Values`.
func (r EventEvolutionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventEvolutionListParams struct {
	// (One or more of fields 'eventId, startTime' are required.) User-provided unique
	// identifier of this activity or event. This ID should remain the same on
	// subsequent updates in order to associate all records pertaining to the activity
	// or event.
	EventID     param.Opt[string] `query:"eventId,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'eventId, startTime' are required.) The actual or
	// estimated start time of the activity or event, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EventEvolutionListParams]'s query parameters as
// `url.Values`.
func (r EventEvolutionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventEvolutionCountParams struct {
	// (One or more of fields 'eventId, startTime' are required.) User-provided unique
	// identifier of this activity or event. This ID should remain the same on
	// subsequent updates in order to associate all records pertaining to the activity
	// or event.
	EventID     param.Opt[string] `query:"eventId,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'eventId, startTime' are required.) The actual or
	// estimated start time of the activity or event, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EventEvolutionCountParams]'s query parameters as
// `url.Values`.
func (r EventEvolutionCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventEvolutionNewBulkParams struct {
	Body []EventEvolutionNewBulkParamsBody
	paramObj
}

func (r EventEvolutionNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *EventEvolutionNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Event Evolution is a unique service supporting the association of UDL records of
// various data types to a common event or activity. The associations may be a one
// time summary, aggregating sources of a past event, or of an ongoing activity
// that evolves over a period of time.
//
// The properties ClassificationMarking, DataMode, EventID, Source, StartTime,
// Summary are required.
type EventEvolutionNewBulkParamsBody struct {
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
	// User-provided unique identifier of this activity or event. This ID should remain
	// the same on subsequent updates in order to associate all records pertaining to
	// the activity or event.
	EventID string `json:"eventId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual or estimated start time of the activity or event, in ISO 8601 UTC
	// format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Summary or description of the activity or event.
	Summary string `json:"summary,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the point of interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// The activity or event type associated with this record (e.g. BREAKUP, DIRECT
	// FIRE, IED, LAUNCH, PROTEST, etc.). For Significant Activities, recommended but
	// not constrained to, CAMEO.Manual.1.1b3 Chapter 6. Note that the evolution of an
	// event may incorporate records of various types, for example, a LAUNCH event may
	// evolve into a BREAKUP event.
	Category param.Opt[string] `json:"category,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Optional description of the relationship between the records provided in the
	// srcTyps/srcIds and the activity or event.
	DataDescription param.Opt[string] `json:"dataDescription,omitzero"`
	// The actual or estimated start time of the activity or event, in ISO 8601 UTC
	// format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Administrative boundaries of the first sub-national level. Level 1 is simply the
	// largest demarcation under whatever demarcation criteria has been determined by
	// the governing body. For example this may be the state/province in which a
	// terrestrial event takes place, or with which the event is attributed for
	// non-localized or non-terrestrial activity.
	GeoAdminLevel1 param.Opt[string] `json:"geoAdminLevel1,omitzero"`
	// Administrative boundaries of the second sub-national level. Level 2 is simply
	// the second largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example this may be the city/district in
	// which a terrestrial event takes place, or with which the event is attributed for
	// non-localized or non-terrestrial activity.
	GeoAdminLevel2 param.Opt[string] `json:"geoAdminLevel2,omitzero"`
	// Administrative boundaries of the third sub-national level. Level 3 is simply the
	// third largest demarcation under whatever demarcation criteria has been
	// determined by the governing body.
	GeoAdminLevel3 param.Opt[string] `json:"geoAdminLevel3,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Flag indicating that this record is for the purpose of redacting one or more
	// previously specified records from association with this activity or event. If
	// this flag is set then all records indicated in srcTyps/srcIds should be removed
	// from event association.
	Redact param.Opt[bool] `json:"redact,omitzero"`
	// The status of this activity or event. (ACTIVE, CONCLUDED, UNKNOWN).
	Status param.Opt[string] `json:"status,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this activity or event. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size. See the corresponding srcTyps
	// array element for the data type of the UUID and use the appropriate API
	// operation to retrieve that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK) that
	// are related to this activity or event. See the associated 'srcIds' array for the
	// record UUIDs, positionally corresponding to the record types in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// List of URLs to before/after images of this point of interest entity.
	URL []string `json:"url,omitzero"`
	paramObj
}

func (r EventEvolutionNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EventEvolutionNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventEvolutionNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EventEvolutionNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type EventEvolutionTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// (One or more of fields 'eventId, startTime' are required.) User-provided unique
	// identifier of this activity or event. This ID should remain the same on
	// subsequent updates in order to associate all records pertaining to the activity
	// or event.
	EventID     param.Opt[string] `query:"eventId,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'eventId, startTime' are required.) The actual or
	// estimated start time of the activity or event, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EventEvolutionTupleParams]'s query parameters as
// `url.Values`.
func (r EventEvolutionTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventEvolutionUnvalidatedPublishParams struct {
	Body []EventEvolutionUnvalidatedPublishParamsBody
	paramObj
}

func (r EventEvolutionUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *EventEvolutionUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Event Evolution is a unique service supporting the association of UDL records of
// various data types to a common event or activity. The associations may be a one
// time summary, aggregating sources of a past event, or of an ongoing activity
// that evolves over a period of time.
//
// The properties ClassificationMarking, DataMode, EventID, Source, StartTime,
// Summary are required.
type EventEvolutionUnvalidatedPublishParamsBody struct {
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
	// User-provided unique identifier of this activity or event. This ID should remain
	// the same on subsequent updates in order to associate all records pertaining to
	// the activity or event.
	EventID string `json:"eventId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual or estimated start time of the activity or event, in ISO 8601 UTC
	// format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Summary or description of the activity or event.
	Summary string `json:"summary,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the point of interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// The activity or event type associated with this record (e.g. BREAKUP, DIRECT
	// FIRE, IED, LAUNCH, PROTEST, etc.). For Significant Activities, recommended but
	// not constrained to, CAMEO.Manual.1.1b3 Chapter 6. Note that the evolution of an
	// event may incorporate records of various types, for example, a LAUNCH event may
	// evolve into a BREAKUP event.
	Category param.Opt[string] `json:"category,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Optional description of the relationship between the records provided in the
	// srcTyps/srcIds and the activity or event.
	DataDescription param.Opt[string] `json:"dataDescription,omitzero"`
	// The actual or estimated start time of the activity or event, in ISO 8601 UTC
	// format.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Administrative boundaries of the first sub-national level. Level 1 is simply the
	// largest demarcation under whatever demarcation criteria has been determined by
	// the governing body. For example this may be the state/province in which a
	// terrestrial event takes place, or with which the event is attributed for
	// non-localized or non-terrestrial activity.
	GeoAdminLevel1 param.Opt[string] `json:"geoAdminLevel1,omitzero"`
	// Administrative boundaries of the second sub-national level. Level 2 is simply
	// the second largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example this may be the city/district in
	// which a terrestrial event takes place, or with which the event is attributed for
	// non-localized or non-terrestrial activity.
	GeoAdminLevel2 param.Opt[string] `json:"geoAdminLevel2,omitzero"`
	// Administrative boundaries of the third sub-national level. Level 3 is simply the
	// third largest demarcation under whatever demarcation criteria has been
	// determined by the governing body.
	GeoAdminLevel3 param.Opt[string] `json:"geoAdminLevel3,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Flag indicating that this record is for the purpose of redacting one or more
	// previously specified records from association with this activity or event. If
	// this flag is set then all records indicated in srcTyps/srcIds should be removed
	// from event association.
	Redact param.Opt[bool] `json:"redact,omitzero"`
	// The status of this activity or event. (ACTIVE, CONCLUDED, UNKNOWN).
	Status param.Opt[string] `json:"status,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this activity or event. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size. See the corresponding srcTyps
	// array element for the data type of the UUID and use the appropriate API
	// operation to retrieve that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK) that
	// are related to this activity or event. See the associated 'srcIds' array for the
	// record UUIDs, positionally corresponding to the record types in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// List of URLs to before/after images of this point of interest entity.
	URL []string `json:"url,omitzero"`
	paramObj
}

func (r EventEvolutionUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EventEvolutionUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventEvolutionUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EventEvolutionUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
