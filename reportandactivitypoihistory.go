// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// ReportAndActivityPoiHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportAndActivityPoiHistoryService] method instead.
type ReportAndActivityPoiHistoryService struct {
	Options []option.RequestOption
}

// NewReportAndActivityPoiHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewReportAndActivityPoiHistoryService(opts ...option.RequestOption) (r ReportAndActivityPoiHistoryService) {
	r = ReportAndActivityPoiHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ReportAndActivityPoiHistoryService) List(ctx context.Context, query ReportAndActivityPoiHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ReportAndActivityPoiHistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/poi/history"
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
func (r *ReportAndActivityPoiHistoryService) ListAutoPaging(ctx context.Context, query ReportAndActivityPoiHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ReportAndActivityPoiHistoryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ReportAndActivityPoiHistoryService) Aodr(ctx context.Context, query ReportAndActivityPoiHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/poi/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ReportAndActivityPoiHistoryService) Count(ctx context.Context, query ReportAndActivityPoiHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/poi/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A Point of Interest is loosely based on the MITRE CoT (Cursor on Target) schema
// (https://www.mitre.org/publications/technical-papers/cursorontarget-message-router-users-guide)
// and provides a simple way to specify a point on the earth for a variety of
// purposes (tasking, targeting, etc).
type ReportAndActivityPoiHistoryListResponse struct {
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
	DataMode ReportAndActivityPoiHistoryListResponseDataMode `json:"dataMode,required"`
	// Name of the POI target object.
	Name string `json:"name,required"`
	// Identifier of the actual Point of Interest or target object, which should remain
	// the same on subsequent POI records of the same Point of Interest.
	Poiid string `json:"poiid,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Activity/POI timestamp in ISO8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The activity in which the POI subject is engaged. Intended as, but not
	// constrained to, MIL-STD-6016 environment dependent activity designations. The
	// activity can be reported as either a combination of the code and environment
	// (e.g. 30/LAND) or as the descriptive enumeration (e.g. TRAINING), which are
	// equivalent.
	Activity string `json:"activity"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// Point height above ellipsoid (WGS-84), in meters.
	Alt float64 `json:"alt"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the Point of Interest as projected on the ground.
	Area string `json:"area"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// ID/name of the platform or entity providing the POI data.
	Asset string `json:"asset"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground.
	Atype string `json:"atype"`
	// Target object pointing azimuth angle, in degrees (for target with sensing or
	// emitting capability).
	Az float64 `json:"az"`
	// The Basic Encyclopedia Number associated with the POI, if applicable.
	BeNumber string `json:"beNumber"`
	// Radius of circular area about lat/lon point, in meters (1-sigma, if representing
	// error).
	Ce float64 `json:"ce"`
	// Contact information for assets reporting PPLI (Precise Participant Location and
	// Identification). PPLI is a Link 16 message that is used by units to transmit
	// complete location, identification, and limited status information.
	Cntct string `json:"cntct"`
	// POI confidence estimate (not standardized, but typically a value between 0 and
	// 1, with 0 indicating lowest confidence.
	Conf float64 `json:"conf"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Description of the POI target object.
	Desc string `json:"desc"`
	// Target object pointing elevation angle, in degrees (for target with sensing or
	// emitting capability).
	El float64 `json:"el"`
	// Elliptical area about the lat/lon point, specified as [semi-major axis (m),
	// semi-minor axis (m), orientation (deg) off true North at POI].
	Elle []float64 `json:"elle"`
	// POI environment type (e.g., LAND, SURFACE, SUBSURFACE, UNKNOWN, etc.).
	Env string `json:"env"`
	// Optional array of groups used when a POI msg originates from a TAK server. Each
	// group must be no longer than 256 characters. Groups identify a set of users
	// targeted by the cot/poi msg.
	Groups []string `json:"groups"`
	// How the event point was generated, in CoT object heirarchy notation (optional,
	// CoT).
	How string `json:"how"`
	// Estimated identity of the point/object (e.g., FRIEND, HOSTILE, SUSPECT,
	// ASSUMED_FRIEND, UNKNOWN, etc.).
	Ident string `json:"ident"`
	// Array of one or more unique identifiers of weather records associated with this
	// POI. Each element in array must be 36 characters or less in length.
	IDWeatherReport []string `json:"idWeatherReport"`
	// WGS-84 latitude of the POI, in degrees (+N, -S), -90 to 90.
	Lat float64 `json:"lat"`
	// Height above lat/lon point, in meters (1-sigma, if representing linear error).
	Le float64 `json:"le"`
	// WGS-84 longitude of the POI, in degrees (+E, -W), -180 to 180.
	Lon float64 `json:"lon"`
	// Optional mission ID related to the POI.
	Msnid string `json:"msnid"`
	// The orientation of a vehicle, platform or other entity described by the POI. The
	// orientation is defined as the pointing direction of the front/nose of the object
	// in degrees clockwise from true North at the object point.
	Orientation float64 `json:"orientation"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// POI/object platform type (e.g., 14/GROUND, COMBAT_VEHICLE, etc.).
	Plat string `json:"plat"`
	// The purpose of this Point of Interest record (e.g., BDA, EQPT, EVENT, GEOL,
	// HZRD, PPLI, SHOTBOX, SURVL, TGT, TSK, WTHR).
	Pps string `json:"pps"`
	// Priority of the POI target object.
	Pri int64 `json:"pri"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Specific point/object type (e.g., 82/GROUND, LIGHT_TANK, etc.).
	Spec string `json:"spec"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this Point of Interest. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size. See the corresponding srcTyps
	// array element for the data type of the UUID and use the appropriate API
	// operation to retrieve that object (e.g. /udl/rfobservation/{uuid}).
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (GROUNDIMAGE, RFOBS) that are related to the
	// determination of this Point of Interest. See the associated 'srcIds' array for
	// the record UUIDs, positionally corresponding to the record types in this array.
	// The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// Stale timestamp (optional), in ISO8601 UTC format.
	Stale time.Time `json:"stale" format:"date-time"`
	// Start time of event validity (optional), in ISO8601 UTC format.
	Start time.Time `json:"start" format:"date-time"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Optional ID of an associated track related to the POI object, if applicable.
	// This track ID should correlate the Point of Interest to a track from the Track
	// service.
	Trkid string `json:"trkid"`
	// Event type, in CoT object heirarchy notation (optional, CoT).
	Type string `json:"type"`
	// List of URLs to before/after images of this Point of Interest entity.
	URLs []string `json:"urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Poiid                 respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		Activity              respjson.Field
		Agjson                respjson.Field
		Alt                   respjson.Field
		Andims                respjson.Field
		Area                  respjson.Field
		Asrid                 respjson.Field
		Asset                 respjson.Field
		Atext                 respjson.Field
		Atype                 respjson.Field
		Az                    respjson.Field
		BeNumber              respjson.Field
		Ce                    respjson.Field
		Cntct                 respjson.Field
		Conf                  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Desc                  respjson.Field
		El                    respjson.Field
		Elle                  respjson.Field
		Env                   respjson.Field
		Groups                respjson.Field
		How                   respjson.Field
		Ident                 respjson.Field
		IDWeatherReport       respjson.Field
		Lat                   respjson.Field
		Le                    respjson.Field
		Lon                   respjson.Field
		Msnid                 respjson.Field
		Orientation           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Plat                  respjson.Field
		Pps                   respjson.Field
		Pri                   respjson.Field
		SourceDl              respjson.Field
		Spec                  respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		Stale                 respjson.Field
		Start                 respjson.Field
		Tags                  respjson.Field
		TransactionID         respjson.Field
		Trkid                 respjson.Field
		Type                  respjson.Field
		URLs                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportAndActivityPoiHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportAndActivityPoiHistoryListResponse) UnmarshalJSON(data []byte) error {
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
type ReportAndActivityPoiHistoryListResponseDataMode string

const (
	ReportAndActivityPoiHistoryListResponseDataModeReal      ReportAndActivityPoiHistoryListResponseDataMode = "REAL"
	ReportAndActivityPoiHistoryListResponseDataModeTest      ReportAndActivityPoiHistoryListResponseDataMode = "TEST"
	ReportAndActivityPoiHistoryListResponseDataModeSimulated ReportAndActivityPoiHistoryListResponseDataMode = "SIMULATED"
	ReportAndActivityPoiHistoryListResponseDataModeExercise  ReportAndActivityPoiHistoryListResponseDataMode = "EXERCISE"
)

type ReportAndActivityPoiHistoryListParams struct {
	// Activity/POI timestamp in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts time.Time `query:"ts,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportAndActivityPoiHistoryListParams]'s query parameters
// as `url.Values`.
func (r ReportAndActivityPoiHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReportAndActivityPoiHistoryAodrParams struct {
	// Activity/POI timestamp in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts time.Time `query:"ts,required" format:"date-time" json:"-"`
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

// URLQuery serializes [ReportAndActivityPoiHistoryAodrParams]'s query parameters
// as `url.Values`.
func (r ReportAndActivityPoiHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReportAndActivityPoiHistoryCountParams struct {
	// Activity/POI timestamp in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportAndActivityPoiHistoryCountParams]'s query parameters
// as `url.Values`.
func (r ReportAndActivityPoiHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
