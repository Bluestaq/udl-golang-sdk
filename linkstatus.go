// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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
)

// LinkStatusService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLinkStatusService] method instead.
type LinkStatusService struct {
	Options  []option.RequestOption
	Datalink LinkStatusDatalinkService
	History  LinkStatusHistoryService
}

// NewLinkStatusService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLinkStatusService(opts ...option.RequestOption) (r LinkStatusService) {
	r = LinkStatusService{}
	r.Options = opts
	r.Datalink = NewLinkStatusDatalinkService(opts...)
	r.History = NewLinkStatusHistoryService(opts...)
	return
}

// Service operation to take a single LinkStatus as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *LinkStatusService) New(ctx context.Context, body LinkStatusNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/linkstatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LinkStatusService) List(ctx context.Context, query LinkStatusListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LinkStatusListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/linkstatus"
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
func (r *LinkStatusService) ListAutoPaging(ctx context.Context, query LinkStatusListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LinkStatusListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LinkStatusService) Count(ctx context.Context, query LinkStatusCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/linkstatus/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single LinkStatus record by its unique ID passed as a
// path parameter.
func (r *LinkStatusService) Get(ctx context.Context, id string, query LinkStatusGetParams, opts ...option.RequestOption) (res *LinkStatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/linkstatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *LinkStatusService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/linkstatus/queryhelp"
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
func (r *LinkStatusService) Tuple(ctx context.Context, query LinkStatusTupleParams, opts ...option.RequestOption) (res *[]LinkStatusTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/linkstatus/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Captures link status.
type LinkStatusListResponse struct {
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
	DataMode LinkStatusListResponseDataMode `json:"dataMode,required"`
	// Latitude of link endpoint-1, WGS-84 in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndPoint1Lat float64 `json:"endPoint1Lat,required"`
	// Longitude of link endpoint-1, WGS-84 longitude in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndPoint1Lon float64 `json:"endPoint1Lon,required"`
	// The name or description of link endpoint-1, corresponding to beam-1.
	EndPoint1Name string `json:"endPoint1Name,required"`
	// Latitude of link endpoint-2, WGS-84 in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndPoint2Lat float64 `json:"endPoint2Lat,required"`
	// Longitude of link endpoint-2, WGS-84 longitude in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndPoint2Lon float64 `json:"endPoint2Lon,required"`
	// The name or description of link endpoint-2, corresponding to beam-2.
	EndPoint2Name string `json:"endPoint2Name,required"`
	// The name or description of the link.
	LinkName string `json:"linkName,required"`
	// The link establishment time, or the time that the link becomes available for
	// use, in ISO8601 UTC format.
	LinkStartTime time.Time `json:"linkStartTime,required" format:"date-time"`
	// The link termination time, or the time that the link becomes unavailable for
	// use, in ISO8601 UTC format.
	LinkStopTime time.Time `json:"linkStopTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The RF band employed by the link (e.g. MIL-KA, COM-KA, X-BAND, C-BAND, etc.).
	Band string `json:"band"`
	// The constellation name if the link is established over a LEO/MEO constellation.
	// In this case, idOnOrbit1 and idOnOrbit2 will be null.
	Constellation string `json:"constellation"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The endpoint-1 to endpoint-2 data rate, in kbps.
	DataRate1To2 float64 `json:"dataRate1To2"`
	// The endpoint-2 to endpoint-1 data rate, in kbps.
	DataRate2To1 float64 `json:"dataRate2To1"`
	// The ID of beam-1 forming the link. In the case of two sat link, beam-1
	// corresponds to Sat-1.
	IDBeam1 string `json:"idBeam1"`
	// The ID of beam-2 forming the link. In the case of two sat link, beam-2
	// corresponds to Sat-2.
	IDBeam2 string `json:"idBeam2"`
	// Unique ID of the on-orbit satellite (Sat-1) forming the link. A null value for
	// idOnOrbit1 indicates that the link is formed over a LEO/MEO constellation.
	IDOnOrbit1 string `json:"idOnOrbit1"`
	// Unique ID of the on-orbit satellite (Sat-2) forming the link. A null value for
	// idOnOrbit2 indicates either a link employing only Sat-1 or a link formed over a
	// LEO/MEO constellation.
	IDOnOrbit2 string `json:"idOnOrbit2"`
	// The state of the link (e.g. OK, DEGRADED-WEATHER, DEGRADED-EMI, etc.).
	LinkState string `json:"linkState"`
	// The type of the link.
	LinkType string `json:"linkType"`
	// The OPSCAP mission status of the system(s) forming the link.
	OpsCap string `json:"opsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Satellite/catalog number of the target on-orbit primary object.
	SatNo1 int64 `json:"satNo1"`
	// Satellite/catalog number of the target on-orbit secondary object.
	SatNo2 int64 `json:"satNo2"`
	// The SYSCAP mission status of the system(s) forming the link.
	SysCap string `json:"sysCap"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EndPoint1Lat          respjson.Field
		EndPoint1Lon          respjson.Field
		EndPoint1Name         respjson.Field
		EndPoint2Lat          respjson.Field
		EndPoint2Lon          respjson.Field
		EndPoint2Name         respjson.Field
		LinkName              respjson.Field
		LinkStartTime         respjson.Field
		LinkStopTime          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Band                  respjson.Field
		Constellation         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataRate1To2          respjson.Field
		DataRate2To1          respjson.Field
		IDBeam1               respjson.Field
		IDBeam2               respjson.Field
		IDOnOrbit1            respjson.Field
		IDOnOrbit2            respjson.Field
		LinkState             respjson.Field
		LinkType              respjson.Field
		OpsCap                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SatNo1                respjson.Field
		SatNo2                respjson.Field
		SysCap                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusListResponse) UnmarshalJSON(data []byte) error {
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
type LinkStatusListResponseDataMode string

const (
	LinkStatusListResponseDataModeReal      LinkStatusListResponseDataMode = "REAL"
	LinkStatusListResponseDataModeTest      LinkStatusListResponseDataMode = "TEST"
	LinkStatusListResponseDataModeSimulated LinkStatusListResponseDataMode = "SIMULATED"
	LinkStatusListResponseDataModeExercise  LinkStatusListResponseDataMode = "EXERCISE"
)

// Captures link status.
type LinkStatusGetResponse struct {
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
	DataMode LinkStatusGetResponseDataMode `json:"dataMode,required"`
	// Latitude of link endpoint-1, WGS-84 in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndPoint1Lat float64 `json:"endPoint1Lat,required"`
	// Longitude of link endpoint-1, WGS-84 longitude in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndPoint1Lon float64 `json:"endPoint1Lon,required"`
	// The name or description of link endpoint-1, corresponding to beam-1.
	EndPoint1Name string `json:"endPoint1Name,required"`
	// Latitude of link endpoint-2, WGS-84 in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndPoint2Lat float64 `json:"endPoint2Lat,required"`
	// Longitude of link endpoint-2, WGS-84 longitude in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndPoint2Lon float64 `json:"endPoint2Lon,required"`
	// The name or description of link endpoint-2, corresponding to beam-2.
	EndPoint2Name string `json:"endPoint2Name,required"`
	// The name or description of the link.
	LinkName string `json:"linkName,required"`
	// The link establishment time, or the time that the link becomes available for
	// use, in ISO8601 UTC format.
	LinkStartTime time.Time `json:"linkStartTime,required" format:"date-time"`
	// The link termination time, or the time that the link becomes unavailable for
	// use, in ISO8601 UTC format.
	LinkStopTime time.Time `json:"linkStopTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The RF band employed by the link (e.g. MIL-KA, COM-KA, X-BAND, C-BAND, etc.).
	Band string `json:"band"`
	// The constellation name if the link is established over a LEO/MEO constellation.
	// In this case, idOnOrbit1 and idOnOrbit2 will be null.
	Constellation string `json:"constellation"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The endpoint-1 to endpoint-2 data rate, in kbps.
	DataRate1To2 float64 `json:"dataRate1To2"`
	// The endpoint-2 to endpoint-1 data rate, in kbps.
	DataRate2To1 float64 `json:"dataRate2To1"`
	// The ID of beam-1 forming the link. In the case of two sat link, beam-1
	// corresponds to Sat-1.
	IDBeam1 string `json:"idBeam1"`
	// The ID of beam-2 forming the link. In the case of two sat link, beam-2
	// corresponds to Sat-2.
	IDBeam2 string `json:"idBeam2"`
	// Unique ID of the on-orbit satellite (Sat-1) forming the link. A null value for
	// idOnOrbit1 indicates that the link is formed over a LEO/MEO constellation.
	IDOnOrbit1 string `json:"idOnOrbit1"`
	// Unique ID of the on-orbit satellite (Sat-2) forming the link. A null value for
	// idOnOrbit2 indicates either a link employing only Sat-1 or a link formed over a
	// LEO/MEO constellation.
	IDOnOrbit2 string `json:"idOnOrbit2"`
	// The state of the link (e.g. OK, DEGRADED-WEATHER, DEGRADED-EMI, etc.).
	LinkState string `json:"linkState"`
	// The type of the link.
	LinkType string `json:"linkType"`
	// The OPSCAP mission status of the system(s) forming the link.
	OpsCap string `json:"opsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Satellite/catalog number of the target on-orbit primary object.
	SatNo1 int64 `json:"satNo1"`
	// Satellite/catalog number of the target on-orbit secondary object.
	SatNo2 int64 `json:"satNo2"`
	// The SYSCAP mission status of the system(s) forming the link.
	SysCap string `json:"sysCap"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EndPoint1Lat          respjson.Field
		EndPoint1Lon          respjson.Field
		EndPoint1Name         respjson.Field
		EndPoint2Lat          respjson.Field
		EndPoint2Lon          respjson.Field
		EndPoint2Name         respjson.Field
		LinkName              respjson.Field
		LinkStartTime         respjson.Field
		LinkStopTime          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Band                  respjson.Field
		Constellation         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataRate1To2          respjson.Field
		DataRate2To1          respjson.Field
		IDBeam1               respjson.Field
		IDBeam2               respjson.Field
		IDOnOrbit1            respjson.Field
		IDOnOrbit2            respjson.Field
		LinkState             respjson.Field
		LinkType              respjson.Field
		OpsCap                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SatNo1                respjson.Field
		SatNo2                respjson.Field
		SysCap                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusGetResponse) UnmarshalJSON(data []byte) error {
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
type LinkStatusGetResponseDataMode string

const (
	LinkStatusGetResponseDataModeReal      LinkStatusGetResponseDataMode = "REAL"
	LinkStatusGetResponseDataModeTest      LinkStatusGetResponseDataMode = "TEST"
	LinkStatusGetResponseDataModeSimulated LinkStatusGetResponseDataMode = "SIMULATED"
	LinkStatusGetResponseDataModeExercise  LinkStatusGetResponseDataMode = "EXERCISE"
)

// Captures link status.
type LinkStatusTupleResponse struct {
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
	DataMode LinkStatusTupleResponseDataMode `json:"dataMode,required"`
	// Latitude of link endpoint-1, WGS-84 in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndPoint1Lat float64 `json:"endPoint1Lat,required"`
	// Longitude of link endpoint-1, WGS-84 longitude in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndPoint1Lon float64 `json:"endPoint1Lon,required"`
	// The name or description of link endpoint-1, corresponding to beam-1.
	EndPoint1Name string `json:"endPoint1Name,required"`
	// Latitude of link endpoint-2, WGS-84 in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndPoint2Lat float64 `json:"endPoint2Lat,required"`
	// Longitude of link endpoint-2, WGS-84 longitude in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndPoint2Lon float64 `json:"endPoint2Lon,required"`
	// The name or description of link endpoint-2, corresponding to beam-2.
	EndPoint2Name string `json:"endPoint2Name,required"`
	// The name or description of the link.
	LinkName string `json:"linkName,required"`
	// The link establishment time, or the time that the link becomes available for
	// use, in ISO8601 UTC format.
	LinkStartTime time.Time `json:"linkStartTime,required" format:"date-time"`
	// The link termination time, or the time that the link becomes unavailable for
	// use, in ISO8601 UTC format.
	LinkStopTime time.Time `json:"linkStopTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The RF band employed by the link (e.g. MIL-KA, COM-KA, X-BAND, C-BAND, etc.).
	Band string `json:"band"`
	// The constellation name if the link is established over a LEO/MEO constellation.
	// In this case, idOnOrbit1 and idOnOrbit2 will be null.
	Constellation string `json:"constellation"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The endpoint-1 to endpoint-2 data rate, in kbps.
	DataRate1To2 float64 `json:"dataRate1To2"`
	// The endpoint-2 to endpoint-1 data rate, in kbps.
	DataRate2To1 float64 `json:"dataRate2To1"`
	// The ID of beam-1 forming the link. In the case of two sat link, beam-1
	// corresponds to Sat-1.
	IDBeam1 string `json:"idBeam1"`
	// The ID of beam-2 forming the link. In the case of two sat link, beam-2
	// corresponds to Sat-2.
	IDBeam2 string `json:"idBeam2"`
	// Unique ID of the on-orbit satellite (Sat-1) forming the link. A null value for
	// idOnOrbit1 indicates that the link is formed over a LEO/MEO constellation.
	IDOnOrbit1 string `json:"idOnOrbit1"`
	// Unique ID of the on-orbit satellite (Sat-2) forming the link. A null value for
	// idOnOrbit2 indicates either a link employing only Sat-1 or a link formed over a
	// LEO/MEO constellation.
	IDOnOrbit2 string `json:"idOnOrbit2"`
	// The state of the link (e.g. OK, DEGRADED-WEATHER, DEGRADED-EMI, etc.).
	LinkState string `json:"linkState"`
	// The type of the link.
	LinkType string `json:"linkType"`
	// The OPSCAP mission status of the system(s) forming the link.
	OpsCap string `json:"opsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Satellite/catalog number of the target on-orbit primary object.
	SatNo1 int64 `json:"satNo1"`
	// Satellite/catalog number of the target on-orbit secondary object.
	SatNo2 int64 `json:"satNo2"`
	// The SYSCAP mission status of the system(s) forming the link.
	SysCap string `json:"sysCap"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EndPoint1Lat          respjson.Field
		EndPoint1Lon          respjson.Field
		EndPoint1Name         respjson.Field
		EndPoint2Lat          respjson.Field
		EndPoint2Lon          respjson.Field
		EndPoint2Name         respjson.Field
		LinkName              respjson.Field
		LinkStartTime         respjson.Field
		LinkStopTime          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Band                  respjson.Field
		Constellation         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataRate1To2          respjson.Field
		DataRate2To1          respjson.Field
		IDBeam1               respjson.Field
		IDBeam2               respjson.Field
		IDOnOrbit1            respjson.Field
		IDOnOrbit2            respjson.Field
		LinkState             respjson.Field
		LinkType              respjson.Field
		OpsCap                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SatNo1                respjson.Field
		SatNo2                respjson.Field
		SysCap                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusTupleResponse) UnmarshalJSON(data []byte) error {
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
type LinkStatusTupleResponseDataMode string

const (
	LinkStatusTupleResponseDataModeReal      LinkStatusTupleResponseDataMode = "REAL"
	LinkStatusTupleResponseDataModeTest      LinkStatusTupleResponseDataMode = "TEST"
	LinkStatusTupleResponseDataModeSimulated LinkStatusTupleResponseDataMode = "SIMULATED"
	LinkStatusTupleResponseDataModeExercise  LinkStatusTupleResponseDataMode = "EXERCISE"
)

type LinkStatusNewParams struct {
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
	DataMode LinkStatusNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Latitude of link endpoint-1, WGS-84 in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndPoint1Lat float64 `json:"endPoint1Lat,required"`
	// Longitude of link endpoint-1, WGS-84 longitude in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndPoint1Lon float64 `json:"endPoint1Lon,required"`
	// The name or description of link endpoint-1, corresponding to beam-1.
	EndPoint1Name string `json:"endPoint1Name,required"`
	// Latitude of link endpoint-2, WGS-84 in degrees. -90 to 90 degrees (negative
	// values south of equator).
	EndPoint2Lat float64 `json:"endPoint2Lat,required"`
	// Longitude of link endpoint-2, WGS-84 longitude in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	EndPoint2Lon float64 `json:"endPoint2Lon,required"`
	// The name or description of link endpoint-2, corresponding to beam-2.
	EndPoint2Name string `json:"endPoint2Name,required"`
	// The name or description of the link.
	LinkName string `json:"linkName,required"`
	// The link establishment time, or the time that the link becomes available for
	// use, in ISO8601 UTC format.
	LinkStartTime time.Time `json:"linkStartTime,required" format:"date-time"`
	// The link termination time, or the time that the link becomes unavailable for
	// use, in ISO8601 UTC format.
	LinkStopTime time.Time `json:"linkStopTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The RF band employed by the link (e.g. MIL-KA, COM-KA, X-BAND, C-BAND, etc.).
	Band param.Opt[string] `json:"band,omitzero"`
	// The constellation name if the link is established over a LEO/MEO constellation.
	// In this case, idOnOrbit1 and idOnOrbit2 will be null.
	Constellation param.Opt[string] `json:"constellation,omitzero"`
	// The endpoint-1 to endpoint-2 data rate, in kbps.
	DataRate1To2 param.Opt[float64] `json:"dataRate1To2,omitzero"`
	// The endpoint-2 to endpoint-1 data rate, in kbps.
	DataRate2To1 param.Opt[float64] `json:"dataRate2To1,omitzero"`
	// The ID of beam-1 forming the link. In the case of two sat link, beam-1
	// corresponds to Sat-1.
	IDBeam1 param.Opt[string] `json:"idBeam1,omitzero"`
	// The ID of beam-2 forming the link. In the case of two sat link, beam-2
	// corresponds to Sat-2.
	IDBeam2 param.Opt[string] `json:"idBeam2,omitzero"`
	// The state of the link (e.g. OK, DEGRADED-WEATHER, DEGRADED-EMI, etc.).
	LinkState param.Opt[string] `json:"linkState,omitzero"`
	// The type of the link.
	LinkType param.Opt[string] `json:"linkType,omitzero"`
	// The OPSCAP mission status of the system(s) forming the link.
	OpsCap param.Opt[string] `json:"opsCap,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Satellite/catalog number of the target on-orbit primary object.
	SatNo1 param.Opt[int64] `json:"satNo1,omitzero"`
	// Satellite/catalog number of the target on-orbit secondary object.
	SatNo2 param.Opt[int64] `json:"satNo2,omitzero"`
	// The SYSCAP mission status of the system(s) forming the link.
	SysCap param.Opt[string] `json:"sysCap,omitzero"`
	paramObj
}

func (r LinkStatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LinkStatusNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkStatusNewParams) UnmarshalJSON(data []byte) error {
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
type LinkStatusNewParamsDataMode string

const (
	LinkStatusNewParamsDataModeReal      LinkStatusNewParamsDataMode = "REAL"
	LinkStatusNewParamsDataModeTest      LinkStatusNewParamsDataMode = "TEST"
	LinkStatusNewParamsDataModeSimulated LinkStatusNewParamsDataMode = "SIMULATED"
	LinkStatusNewParamsDataModeExercise  LinkStatusNewParamsDataMode = "EXERCISE"
)

type LinkStatusListParams struct {
	// (One or more of fields 'createdAt, linkStartTime, linkStopTime' are required.)
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   param.Opt[time.Time] `query:"createdAt,omitzero" format:"date" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'createdAt, linkStartTime, linkStopTime' are required.)
	// The link establishment time, or the time that the link becomes available for
	// use, in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	LinkStartTime param.Opt[time.Time] `query:"linkStartTime,omitzero" format:"date-time" json:"-"`
	// (One or more of fields 'createdAt, linkStartTime, linkStopTime' are required.)
	// The link termination time, or the time that the link becomes unavailable for
	// use, in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	LinkStopTime param.Opt[time.Time] `query:"linkStopTime,omitzero" format:"date-time" json:"-"`
	MaxResults   param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LinkStatusListParams]'s query parameters as `url.Values`.
func (r LinkStatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LinkStatusCountParams struct {
	// (One or more of fields 'createdAt, linkStartTime, linkStopTime' are required.)
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   param.Opt[time.Time] `query:"createdAt,omitzero" format:"date" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'createdAt, linkStartTime, linkStopTime' are required.)
	// The link establishment time, or the time that the link becomes available for
	// use, in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	LinkStartTime param.Opt[time.Time] `query:"linkStartTime,omitzero" format:"date-time" json:"-"`
	// (One or more of fields 'createdAt, linkStartTime, linkStopTime' are required.)
	// The link termination time, or the time that the link becomes unavailable for
	// use, in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	LinkStopTime param.Opt[time.Time] `query:"linkStopTime,omitzero" format:"date-time" json:"-"`
	MaxResults   param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LinkStatusCountParams]'s query parameters as `url.Values`.
func (r LinkStatusCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LinkStatusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LinkStatusGetParams]'s query parameters as `url.Values`.
func (r LinkStatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LinkStatusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// (One or more of fields 'createdAt, linkStartTime, linkStopTime' are required.)
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   param.Opt[time.Time] `query:"createdAt,omitzero" format:"date" json:"-"`
	FirstResult param.Opt[int64]     `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'createdAt, linkStartTime, linkStopTime' are required.)
	// The link establishment time, or the time that the link becomes available for
	// use, in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	LinkStartTime param.Opt[time.Time] `query:"linkStartTime,omitzero" format:"date-time" json:"-"`
	// (One or more of fields 'createdAt, linkStartTime, linkStopTime' are required.)
	// The link termination time, or the time that the link becomes unavailable for
	// use, in ISO8601 UTC format. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	LinkStopTime param.Opt[time.Time] `query:"linkStopTime,omitzero" format:"date-time" json:"-"`
	MaxResults   param.Opt[int64]     `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LinkStatusTupleParams]'s query parameters as `url.Values`.
func (r LinkStatusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
