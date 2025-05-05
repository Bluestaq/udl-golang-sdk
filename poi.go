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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// PoiService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPoiService] method instead.
type PoiService struct {
	Options []option.RequestOption
}

// NewPoiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPoiService(opts ...option.RequestOption) (r PoiService) {
	r = PoiService{}
	r.Options = opts
	return
}

// Service operation to take a single POI as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *PoiService) New(ctx context.Context, body PoiNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/poi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *PoiService) List(ctx context.Context, query PoiListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[PoiListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/poi"
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
func (r *PoiService) ListAutoPaging(ctx context.Context, query PoiListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[PoiListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *PoiService) Count(ctx context.Context, query PoiCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/poi/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of POIs
// as a POST body and ingest into the database. This operation is not intended to
// be used for automated feeds into UDL. Data providers should contact the UDL team
// for specific role assignments and for instructions on setting up a permanent
// feed through an alternate mechanism.
func (r *PoiService) NewBulk(ctx context.Context, body PoiNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/poi/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single POI by its unique ID passed as a path
// parameter.
func (r *PoiService) Get(ctx context.Context, id string, query PoiGetParams, opts ...option.RequestOption) (res *PoiGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/poi/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *PoiService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/poi/queryhelp"
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
func (r *PoiService) Tuple(ctx context.Context, query PoiTupleParams, opts ...option.RequestOption) (res *[]PoiTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/poi/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a list of POIs as a POST body and ingest into the
// database. This operation is intended to be used for automated feeds into UDL. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *PoiService) UnvalidatedPublish(ctx context.Context, body PoiUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-poi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// A Point of Interest is loosely based on the MITRE CoT (Cursor on Target) schema
// (https://www.mitre.org/publications/technical-papers/cursorontarget-message-router-users-guide)
// and provides a simple way to specify a point on the earth for a variety of
// purposes (tasking, targeting, etc).
type PoiListResponse struct {
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
	DataMode PoiListResponseDataMode `json:"dataMode,required"`
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
	// Stale timestamp (optional), in ISO8601 UTC format.
	Stale time.Time `json:"stale" format:"date-time"`
	// Start time of event validity (optional), in ISO8601 UTC format.
	Start time.Time `json:"start" format:"date-time"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Name                  resp.Field
		Poiid                 resp.Field
		Source                resp.Field
		Ts                    resp.Field
		ID                    resp.Field
		Activity              resp.Field
		Agjson                resp.Field
		Alt                   resp.Field
		Andims                resp.Field
		Asrid                 resp.Field
		Asset                 resp.Field
		Atext                 resp.Field
		Atype                 resp.Field
		Az                    resp.Field
		BeNumber              resp.Field
		Ce                    resp.Field
		Cntct                 resp.Field
		Conf                  resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Desc                  resp.Field
		El                    resp.Field
		Elle                  resp.Field
		Env                   resp.Field
		Groups                resp.Field
		How                   resp.Field
		Ident                 resp.Field
		IDWeatherReport       resp.Field
		Lat                   resp.Field
		Le                    resp.Field
		Lon                   resp.Field
		Msnid                 resp.Field
		Orientation           resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Plat                  resp.Field
		Pps                   resp.Field
		Pri                   resp.Field
		SourceDl              resp.Field
		Spec                  resp.Field
		Stale                 resp.Field
		Start                 resp.Field
		TransactionID         resp.Field
		Trkid                 resp.Field
		Type                  resp.Field
		URLs                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoiListResponse) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponse) UnmarshalJSON(data []byte) error {
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
type PoiListResponseDataMode string

const (
	PoiListResponseDataModeReal      PoiListResponseDataMode = "REAL"
	PoiListResponseDataModeTest      PoiListResponseDataMode = "TEST"
	PoiListResponseDataModeSimulated PoiListResponseDataMode = "SIMULATED"
	PoiListResponseDataModeExercise  PoiListResponseDataMode = "EXERCISE"
)

// A Point of Interest is loosely based on the MITRE CoT (Cursor on Target) schema
// (https://www.mitre.org/publications/technical-papers/cursorontarget-message-router-users-guide)
// and provides a simple way to specify a point on the earth for a variety of
// purposes (tasking, targeting, etc).
type PoiGetResponse struct {
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
	DataMode PoiGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Name                  resp.Field
		Poiid                 resp.Field
		Source                resp.Field
		Ts                    resp.Field
		ID                    resp.Field
		Activity              resp.Field
		Agjson                resp.Field
		Alt                   resp.Field
		Andims                resp.Field
		Area                  resp.Field
		Asrid                 resp.Field
		Asset                 resp.Field
		Atext                 resp.Field
		Atype                 resp.Field
		Az                    resp.Field
		BeNumber              resp.Field
		Ce                    resp.Field
		Cntct                 resp.Field
		Conf                  resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Desc                  resp.Field
		El                    resp.Field
		Elle                  resp.Field
		Env                   resp.Field
		Groups                resp.Field
		How                   resp.Field
		Ident                 resp.Field
		IDWeatherReport       resp.Field
		Lat                   resp.Field
		Le                    resp.Field
		Lon                   resp.Field
		Msnid                 resp.Field
		Orientation           resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Plat                  resp.Field
		Pps                   resp.Field
		Pri                   resp.Field
		SourceDl              resp.Field
		Spec                  resp.Field
		SrcIDs                resp.Field
		SrcTyps               resp.Field
		Stale                 resp.Field
		Start                 resp.Field
		Tags                  resp.Field
		TransactionID         resp.Field
		Trkid                 resp.Field
		Type                  resp.Field
		URLs                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoiGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PoiGetResponse) UnmarshalJSON(data []byte) error {
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
type PoiGetResponseDataMode string

const (
	PoiGetResponseDataModeReal      PoiGetResponseDataMode = "REAL"
	PoiGetResponseDataModeTest      PoiGetResponseDataMode = "TEST"
	PoiGetResponseDataModeSimulated PoiGetResponseDataMode = "SIMULATED"
	PoiGetResponseDataModeExercise  PoiGetResponseDataMode = "EXERCISE"
)

// A Point of Interest is loosely based on the MITRE CoT (Cursor on Target) schema
// (https://www.mitre.org/publications/technical-papers/cursorontarget-message-router-users-guide)
// and provides a simple way to specify a point on the earth for a variety of
// purposes (tasking, targeting, etc).
type PoiTupleResponse struct {
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
	DataMode PoiTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Name                  resp.Field
		Poiid                 resp.Field
		Source                resp.Field
		Ts                    resp.Field
		ID                    resp.Field
		Activity              resp.Field
		Agjson                resp.Field
		Alt                   resp.Field
		Andims                resp.Field
		Area                  resp.Field
		Asrid                 resp.Field
		Asset                 resp.Field
		Atext                 resp.Field
		Atype                 resp.Field
		Az                    resp.Field
		BeNumber              resp.Field
		Ce                    resp.Field
		Cntct                 resp.Field
		Conf                  resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Desc                  resp.Field
		El                    resp.Field
		Elle                  resp.Field
		Env                   resp.Field
		Groups                resp.Field
		How                   resp.Field
		Ident                 resp.Field
		IDWeatherReport       resp.Field
		Lat                   resp.Field
		Le                    resp.Field
		Lon                   resp.Field
		Msnid                 resp.Field
		Orientation           resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Plat                  resp.Field
		Pps                   resp.Field
		Pri                   resp.Field
		SourceDl              resp.Field
		Spec                  resp.Field
		SrcIDs                resp.Field
		SrcTyps               resp.Field
		Stale                 resp.Field
		Start                 resp.Field
		Tags                  resp.Field
		TransactionID         resp.Field
		Trkid                 resp.Field
		Type                  resp.Field
		URLs                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoiTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *PoiTupleResponse) UnmarshalJSON(data []byte) error {
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
type PoiTupleResponseDataMode string

const (
	PoiTupleResponseDataModeReal      PoiTupleResponseDataMode = "REAL"
	PoiTupleResponseDataModeTest      PoiTupleResponseDataMode = "TEST"
	PoiTupleResponseDataModeSimulated PoiTupleResponseDataMode = "SIMULATED"
	PoiTupleResponseDataModeExercise  PoiTupleResponseDataMode = "EXERCISE"
)

type PoiNewParams struct {
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
	DataMode PoiNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	ID param.Opt[string] `json:"id,omitzero"`
	// The activity in which the POI subject is engaged. Intended as, but not
	// constrained to, MIL-STD-6016 environment dependent activity designations. The
	// activity can be reported as either a combination of the code and environment
	// (e.g. 30/LAND) or as the descriptive enumeration (e.g. TRAINING), which are
	// equivalent.
	Activity param.Opt[string] `json:"activity,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Point height above ellipsoid (WGS-84), in meters.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the Point of Interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// ID/name of the platform or entity providing the POI data.
	Asset param.Opt[string] `json:"asset,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// Target object pointing azimuth angle, in degrees (for target with sensing or
	// emitting capability).
	Az param.Opt[float64] `json:"az,omitzero"`
	// The Basic Encyclopedia Number associated with the POI, if applicable.
	BeNumber param.Opt[string] `json:"beNumber,omitzero"`
	// Radius of circular area about lat/lon point, in meters (1-sigma, if representing
	// error).
	Ce param.Opt[float64] `json:"ce,omitzero"`
	// Contact information for assets reporting PPLI (Precise Participant Location and
	// Identification). PPLI is a Link 16 message that is used by units to transmit
	// complete location, identification, and limited status information.
	Cntct param.Opt[string] `json:"cntct,omitzero"`
	// POI confidence estimate (not standardized, but typically a value between 0 and
	// 1, with 0 indicating lowest confidence.
	Conf param.Opt[float64] `json:"conf,omitzero"`
	// Description of the POI target object.
	Desc param.Opt[string] `json:"desc,omitzero"`
	// Target object pointing elevation angle, in degrees (for target with sensing or
	// emitting capability).
	El param.Opt[float64] `json:"el,omitzero"`
	// POI environment type (e.g., LAND, SURFACE, SUBSURFACE, UNKNOWN, etc.).
	Env param.Opt[string] `json:"env,omitzero"`
	// How the event point was generated, in CoT object heirarchy notation (optional,
	// CoT).
	How param.Opt[string] `json:"how,omitzero"`
	// Estimated identity of the point/object (e.g., FRIEND, HOSTILE, SUSPECT,
	// ASSUMED_FRIEND, UNKNOWN, etc.).
	Ident param.Opt[string] `json:"ident,omitzero"`
	// WGS-84 latitude of the POI, in degrees (+N, -S), -90 to 90.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Height above lat/lon point, in meters (1-sigma, if representing linear error).
	Le param.Opt[float64] `json:"le,omitzero"`
	// WGS-84 longitude of the POI, in degrees (+E, -W), -180 to 180.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Optional mission ID related to the POI.
	Msnid param.Opt[string] `json:"msnid,omitzero"`
	// The orientation of a vehicle, platform or other entity described by the POI. The
	// orientation is defined as the pointing direction of the front/nose of the object
	// in degrees clockwise from true North at the object point.
	Orientation param.Opt[float64] `json:"orientation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// POI/object platform type (e.g., 14/GROUND, COMBAT_VEHICLE, etc.).
	Plat param.Opt[string] `json:"plat,omitzero"`
	// The purpose of this Point of Interest record (e.g., BDA, EQPT, EVENT, GEOL,
	// HZRD, PPLI, SHOTBOX, SURVL, TGT, TSK, WTHR).
	Pps param.Opt[string] `json:"pps,omitzero"`
	// Priority of the POI target object.
	Pri param.Opt[int64] `json:"pri,omitzero"`
	// Specific point/object type (e.g., 82/GROUND, LIGHT_TANK, etc.).
	Spec param.Opt[string] `json:"spec,omitzero"`
	// Stale timestamp (optional), in ISO8601 UTC format.
	Stale param.Opt[time.Time] `json:"stale,omitzero" format:"date-time"`
	// Start time of event validity (optional), in ISO8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Optional ID of an associated track related to the POI object, if applicable.
	// This track ID should correlate the Point of Interest to a track from the Track
	// service.
	Trkid param.Opt[string] `json:"trkid,omitzero"`
	// Event type, in CoT object heirarchy notation (optional, CoT).
	Type param.Opt[string] `json:"type,omitzero"`
	// Elliptical area about the lat/lon point, specified as [semi-major axis (m),
	// semi-minor axis (m), orientation (deg) off true North at POI].
	Elle []float64 `json:"elle,omitzero"`
	// Optional array of groups used when a POI msg originates from a TAK server. Each
	// group must be no longer than 256 characters. Groups identify a set of users
	// targeted by the cot/poi msg.
	Groups []string `json:"groups,omitzero"`
	// Array of one or more unique identifiers of weather records associated with this
	// POI. Each element in array must be 36 characters or less in length.
	IDWeatherReport []string `json:"idWeatherReport,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this Point of Interest. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size. See the corresponding srcTyps
	// array element for the data type of the UUID and use the appropriate API
	// operation to retrieve that object (e.g. /udl/rfobservation/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (GROUNDIMAGE, RFOBS) that are related to the
	// determination of this Point of Interest. See the associated 'srcIds' array for
	// the record UUIDs, positionally corresponding to the record types in this array.
	// The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// List of URLs to before/after images of this Point of Interest entity.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PoiNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r PoiNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PoiNewParams
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
type PoiNewParamsDataMode string

const (
	PoiNewParamsDataModeReal      PoiNewParamsDataMode = "REAL"
	PoiNewParamsDataModeTest      PoiNewParamsDataMode = "TEST"
	PoiNewParamsDataModeSimulated PoiNewParamsDataMode = "SIMULATED"
	PoiNewParamsDataModeExercise  PoiNewParamsDataMode = "EXERCISE"
)

type PoiListParams struct {
	// Activity/POI timestamp in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PoiListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [PoiListParams]'s query parameters as `url.Values`.
func (r PoiListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PoiCountParams struct {
	// Activity/POI timestamp in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PoiCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [PoiCountParams]'s query parameters as `url.Values`.
func (r PoiCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PoiNewBulkParams struct {
	Body []PoiNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PoiNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r PoiNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// A Point of Interest is loosely based on the MITRE CoT (Cursor on Target) schema
// (https://www.mitre.org/publications/technical-papers/cursorontarget-message-router-users-guide)
// and provides a simple way to specify a point on the earth for a variety of
// purposes (tasking, targeting, etc).
//
// The properties ClassificationMarking, DataMode, Name, Poiid, Source, Ts are
// required.
type PoiNewBulkParamsBody struct {
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
	ID param.Opt[string] `json:"id,omitzero"`
	// The activity in which the POI subject is engaged. Intended as, but not
	// constrained to, MIL-STD-6016 environment dependent activity designations. The
	// activity can be reported as either a combination of the code and environment
	// (e.g. 30/LAND) or as the descriptive enumeration (e.g. TRAINING), which are
	// equivalent.
	Activity param.Opt[string] `json:"activity,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Point height above ellipsoid (WGS-84), in meters.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the Point of Interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// ID/name of the platform or entity providing the POI data.
	Asset param.Opt[string] `json:"asset,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// Target object pointing azimuth angle, in degrees (for target with sensing or
	// emitting capability).
	Az param.Opt[float64] `json:"az,omitzero"`
	// The Basic Encyclopedia Number associated with the POI, if applicable.
	BeNumber param.Opt[string] `json:"beNumber,omitzero"`
	// Radius of circular area about lat/lon point, in meters (1-sigma, if representing
	// error).
	Ce param.Opt[float64] `json:"ce,omitzero"`
	// Contact information for assets reporting PPLI (Precise Participant Location and
	// Identification). PPLI is a Link 16 message that is used by units to transmit
	// complete location, identification, and limited status information.
	Cntct param.Opt[string] `json:"cntct,omitzero"`
	// POI confidence estimate (not standardized, but typically a value between 0 and
	// 1, with 0 indicating lowest confidence.
	Conf param.Opt[float64] `json:"conf,omitzero"`
	// Time the row was created in the database.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Description of the POI target object.
	Desc param.Opt[string] `json:"desc,omitzero"`
	// Target object pointing elevation angle, in degrees (for target with sensing or
	// emitting capability).
	El param.Opt[float64] `json:"el,omitzero"`
	// POI environment type (e.g., LAND, SURFACE, SUBSURFACE, UNKNOWN, etc.).
	Env param.Opt[string] `json:"env,omitzero"`
	// How the event point was generated, in CoT object heirarchy notation (optional,
	// CoT).
	How param.Opt[string] `json:"how,omitzero"`
	// Estimated identity of the point/object (e.g., FRIEND, HOSTILE, SUSPECT,
	// ASSUMED_FRIEND, UNKNOWN, etc.).
	Ident param.Opt[string] `json:"ident,omitzero"`
	// WGS-84 latitude of the POI, in degrees (+N, -S), -90 to 90.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Height above lat/lon point, in meters (1-sigma, if representing linear error).
	Le param.Opt[float64] `json:"le,omitzero"`
	// WGS-84 longitude of the POI, in degrees (+E, -W), -180 to 180.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Optional mission ID related to the POI.
	Msnid param.Opt[string] `json:"msnid,omitzero"`
	// The orientation of a vehicle, platform or other entity described by the POI. The
	// orientation is defined as the pointing direction of the front/nose of the object
	// in degrees clockwise from true North at the object point.
	Orientation param.Opt[float64] `json:"orientation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// POI/object platform type (e.g., 14/GROUND, COMBAT_VEHICLE, etc.).
	Plat param.Opt[string] `json:"plat,omitzero"`
	// The purpose of this Point of Interest record (e.g., BDA, EQPT, EVENT, GEOL,
	// HZRD, PPLI, SHOTBOX, SURVL, TGT, TSK, WTHR).
	Pps param.Opt[string] `json:"pps,omitzero"`
	// Priority of the POI target object.
	Pri param.Opt[int64] `json:"pri,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Specific point/object type (e.g., 82/GROUND, LIGHT_TANK, etc.).
	Spec param.Opt[string] `json:"spec,omitzero"`
	// Stale timestamp (optional), in ISO8601 UTC format.
	Stale param.Opt[time.Time] `json:"stale,omitzero" format:"date-time"`
	// Start time of event validity (optional), in ISO8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Optional ID of an associated track related to the POI object, if applicable.
	// This track ID should correlate the Point of Interest to a track from the Track
	// service.
	Trkid param.Opt[string] `json:"trkid,omitzero"`
	// Event type, in CoT object heirarchy notation (optional, CoT).
	Type param.Opt[string] `json:"type,omitzero"`
	// Elliptical area about the lat/lon point, specified as [semi-major axis (m),
	// semi-minor axis (m), orientation (deg) off true North at POI].
	Elle []float64 `json:"elle,omitzero"`
	// Optional array of groups used when a POI msg originates from a TAK server. Each
	// group must be no longer than 256 characters. Groups identify a set of users
	// targeted by the cot/poi msg.
	Groups []string `json:"groups,omitzero"`
	// Array of one or more unique identifiers of weather records associated with this
	// POI. Each element in array must be 36 characters or less in length.
	IDWeatherReport []string `json:"idWeatherReport,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this Point of Interest. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size. See the corresponding srcTyps
	// array element for the data type of the UUID and use the appropriate API
	// operation to retrieve that object (e.g. /udl/rfobservation/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (GROUNDIMAGE, RFOBS) that are related to the
	// determination of this Point of Interest. See the associated 'srcIds' array for
	// the record UUIDs, positionally corresponding to the record types in this array.
	// The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// List of URLs to before/after images of this Point of Interest entity.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PoiNewBulkParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r PoiNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow PoiNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PoiNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type PoiGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PoiGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [PoiGetParams]'s query parameters as `url.Values`.
func (r PoiGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PoiTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Activity/POI timestamp in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PoiTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [PoiTupleParams]'s query parameters as `url.Values`.
func (r PoiTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PoiUnvalidatedPublishParams struct {
	Body []PoiUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PoiUnvalidatedPublishParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r PoiUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// A Point of Interest is loosely based on the MITRE CoT (Cursor on Target) schema
// (https://www.mitre.org/publications/technical-papers/cursorontarget-message-router-users-guide)
// and provides a simple way to specify a point on the earth for a variety of
// purposes (tasking, targeting, etc).
//
// The properties ClassificationMarking, DataMode, Name, Poiid, Source, Ts are
// required.
type PoiUnvalidatedPublishParamsBody struct {
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
	ID param.Opt[string] `json:"id,omitzero"`
	// The activity in which the POI subject is engaged. Intended as, but not
	// constrained to, MIL-STD-6016 environment dependent activity designations. The
	// activity can be reported as either a combination of the code and environment
	// (e.g. 30/LAND) or as the descriptive enumeration (e.g. TRAINING), which are
	// equivalent.
	Activity param.Opt[string] `json:"activity,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Point height above ellipsoid (WGS-84), in meters.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the Point of Interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// ID/name of the platform or entity providing the POI data.
	Asset param.Opt[string] `json:"asset,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// Target object pointing azimuth angle, in degrees (for target with sensing or
	// emitting capability).
	Az param.Opt[float64] `json:"az,omitzero"`
	// The Basic Encyclopedia Number associated with the POI, if applicable.
	BeNumber param.Opt[string] `json:"beNumber,omitzero"`
	// Radius of circular area about lat/lon point, in meters (1-sigma, if representing
	// error).
	Ce param.Opt[float64] `json:"ce,omitzero"`
	// Contact information for assets reporting PPLI (Precise Participant Location and
	// Identification). PPLI is a Link 16 message that is used by units to transmit
	// complete location, identification, and limited status information.
	Cntct param.Opt[string] `json:"cntct,omitzero"`
	// POI confidence estimate (not standardized, but typically a value between 0 and
	// 1, with 0 indicating lowest confidence.
	Conf param.Opt[float64] `json:"conf,omitzero"`
	// Time the row was created in the database.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Description of the POI target object.
	Desc param.Opt[string] `json:"desc,omitzero"`
	// Target object pointing elevation angle, in degrees (for target with sensing or
	// emitting capability).
	El param.Opt[float64] `json:"el,omitzero"`
	// POI environment type (e.g., LAND, SURFACE, SUBSURFACE, UNKNOWN, etc.).
	Env param.Opt[string] `json:"env,omitzero"`
	// How the event point was generated, in CoT object heirarchy notation (optional,
	// CoT).
	How param.Opt[string] `json:"how,omitzero"`
	// Estimated identity of the point/object (e.g., FRIEND, HOSTILE, SUSPECT,
	// ASSUMED_FRIEND, UNKNOWN, etc.).
	Ident param.Opt[string] `json:"ident,omitzero"`
	// WGS-84 latitude of the POI, in degrees (+N, -S), -90 to 90.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Height above lat/lon point, in meters (1-sigma, if representing linear error).
	Le param.Opt[float64] `json:"le,omitzero"`
	// WGS-84 longitude of the POI, in degrees (+E, -W), -180 to 180.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Optional mission ID related to the POI.
	Msnid param.Opt[string] `json:"msnid,omitzero"`
	// The orientation of a vehicle, platform or other entity described by the POI. The
	// orientation is defined as the pointing direction of the front/nose of the object
	// in degrees clockwise from true North at the object point.
	Orientation param.Opt[float64] `json:"orientation,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// POI/object platform type (e.g., 14/GROUND, COMBAT_VEHICLE, etc.).
	Plat param.Opt[string] `json:"plat,omitzero"`
	// The purpose of this Point of Interest record (e.g., BDA, EQPT, EVENT, GEOL,
	// HZRD, PPLI, SHOTBOX, SURVL, TGT, TSK, WTHR).
	Pps param.Opt[string] `json:"pps,omitzero"`
	// Priority of the POI target object.
	Pri param.Opt[int64] `json:"pri,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Specific point/object type (e.g., 82/GROUND, LIGHT_TANK, etc.).
	Spec param.Opt[string] `json:"spec,omitzero"`
	// Stale timestamp (optional), in ISO8601 UTC format.
	Stale param.Opt[time.Time] `json:"stale,omitzero" format:"date-time"`
	// Start time of event validity (optional), in ISO8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Optional ID of an associated track related to the POI object, if applicable.
	// This track ID should correlate the Point of Interest to a track from the Track
	// service.
	Trkid param.Opt[string] `json:"trkid,omitzero"`
	// Event type, in CoT object heirarchy notation (optional, CoT).
	Type param.Opt[string] `json:"type,omitzero"`
	// Elliptical area about the lat/lon point, specified as [semi-major axis (m),
	// semi-minor axis (m), orientation (deg) off true North at POI].
	Elle []float64 `json:"elle,omitzero"`
	// Optional array of groups used when a POI msg originates from a TAK server. Each
	// group must be no longer than 256 characters. Groups identify a set of users
	// targeted by the cot/poi msg.
	Groups []string `json:"groups,omitzero"`
	// Array of one or more unique identifiers of weather records associated with this
	// POI. Each element in array must be 36 characters or less in length.
	IDWeatherReport []string `json:"idWeatherReport,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the determination of
	// this Point of Interest. See the associated 'srcTyps' array for the specific
	// types of data, positionally corresponding to the UUIDs in this array. The
	// 'srcTyps' and 'srcIds' arrays must match in size. See the corresponding srcTyps
	// array element for the data type of the UUID and use the appropriate API
	// operation to retrieve that object (e.g. /udl/rfobservation/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (GROUNDIMAGE, RFOBS) that are related to the
	// determination of this Point of Interest. See the associated 'srcIds' array for
	// the record UUIDs, positionally corresponding to the record types in this array.
	// The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// List of URLs to before/after images of this Point of Interest entity.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PoiUnvalidatedPublishParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r PoiUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow PoiUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PoiUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
