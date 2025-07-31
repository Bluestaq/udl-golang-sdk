// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
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
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// LaunchSiteDetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaunchSiteDetailService] method instead.
type LaunchSiteDetailService struct {
	Options []option.RequestOption
}

// NewLaunchSiteDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLaunchSiteDetailService(opts ...option.RequestOption) (r LaunchSiteDetailService) {
	r = LaunchSiteDetailService{}
	r.Options = opts
	return
}

// Service operation to take a single LaunchSiteDetails as a POST body and ingest
// into the database. A LaunchSiteDetails represents details compiled/collected on
// a launch site by a particular source. A launch site may have several details
// records. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *LaunchSiteDetailService) New(ctx context.Context, body LaunchSiteDetailNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchsitedetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a LaunchSiteDetails. A LaunchSiteDetails represents
// details compiled/collected on a launch site by a particular source. A launch
// site may have several details records. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *LaunchSiteDetailService) Update(ctx context.Context, id string, body LaunchSiteDetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchsitedetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaunchSiteDetailService) List(ctx context.Context, query LaunchSiteDetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaunchSiteDetailListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/launchsitedetails"
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
func (r *LaunchSiteDetailService) ListAutoPaging(ctx context.Context, query LaunchSiteDetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaunchSiteDetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a LaunchSiteDetails specified by the passed ID path
// parameter. A LaunchSiteDetails represents details compiled/collected on a launch
// site by a particular source. A launch site may have several details records. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *LaunchSiteDetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchsitedetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to get a single LaunchSiteDetails by a source passed as a
// query parameter. A LaunchSiteDetails represents details compiled/collected on a
// launch site by a particular source. A launch site may have several details
// records.
func (r *LaunchSiteDetailService) FindBySource(ctx context.Context, query LaunchSiteDetailFindBySourceParams, opts ...option.RequestOption) (res *[]LaunchSiteDetailFindBySourceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/launchsitedetails/findBySource"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single LaunchSiteDetails by its unique ID passed as a
// path parameter. A LaunchSiteDetails represents details compiled/collected on a
// launch site by a particular source. A launch site may have several details
// records.
func (r *LaunchSiteDetailService) Get(ctx context.Context, id string, query LaunchSiteDetailGetParams, opts ...option.RequestOption) (res *LaunchSiteDetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchsitedetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of details compiled/collected on a launch site by a
// particular source. A launch site may have several details records.
type LaunchSiteDetailListResponse struct {
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
	DataMode LaunchSiteDetailListResponseDataMode `json:"dataMode,required"`
	// Identifier of the parent launch site record.
	IDLaunchSite string `json:"idLaunchSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of supported launch inclinations known for this site. The array is always
	// sized as a multiple of two and includes start/end values to support ranges. For
	// example, if a site support inclinations of 10 and 12-14, the array would have
	// the following values: [10,10, 12,14].
	AvailableInclinations []float64 `json:"availableInclinations"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Notes/description of the launch site.
	Description string `json:"description"`
	// ID of the location data for the launch site, or null if mobile (e.g. sea launch
	// platform in international waters).
	IDLocation string `json:"idLocation"`
	// Launch site group name. Multiple launch sites may be colocated within a launch
	// ”group”.
	LaunchGroup string `json:"launchGroup"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LaunchSiteDetailListResponseLocation `json:"location"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDLaunchSite          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AvailableInclinations respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Description           respjson.Field
		IDLocation            respjson.Field
		LaunchGroup           respjson.Field
		Location              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchSiteDetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteDetailListResponse) UnmarshalJSON(data []byte) error {
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
type LaunchSiteDetailListResponseDataMode string

const (
	LaunchSiteDetailListResponseDataModeReal      LaunchSiteDetailListResponseDataMode = "REAL"
	LaunchSiteDetailListResponseDataModeTest      LaunchSiteDetailListResponseDataMode = "TEST"
	LaunchSiteDetailListResponseDataModeSimulated LaunchSiteDetailListResponseDataMode = "SIMULATED"
	LaunchSiteDetailListResponseDataModeExercise  LaunchSiteDetailListResponseDataMode = "EXERCISE"
)

// Model representation of a location, which is a specific fixed point on the earth
// and is used to denote the locations of fixed sensors, operating units, etc.
type LaunchSiteDetailListResponseLocation struct {
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
	DataMode string `json:"dataMode,required"`
	// Location name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Altitude of the location, in kilometers.
	Altitude float64 `json:"altitude"`
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
	// Unique identifier of the location, auto-generated by the system.
	IDLocation string `json:"idLocation"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		Altitude              respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDLocation            respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchSiteDetailListResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteDetailListResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of details compiled/collected on a launch site by a
// particular source. A launch site may have several details records.
type LaunchSiteDetailFindBySourceResponse struct {
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
	DataMode LaunchSiteDetailFindBySourceResponseDataMode `json:"dataMode,required"`
	// Identifier of the parent launch site record.
	IDLaunchSite string `json:"idLaunchSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of supported launch inclinations known for this site. The array is always
	// sized as a multiple of two and includes start/end values to support ranges. For
	// example, if a site support inclinations of 10 and 12-14, the array would have
	// the following values: [10,10, 12,14].
	AvailableInclinations []float64 `json:"availableInclinations"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Notes/description of the launch site.
	Description string `json:"description"`
	// ID of the location data for the launch site, or null if mobile (e.g. sea launch
	// platform in international waters).
	IDLocation string `json:"idLocation"`
	// Launch site group name. Multiple launch sites may be colocated within a launch
	// ”group”.
	LaunchGroup string `json:"launchGroup"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LaunchSiteDetailFindBySourceResponseLocation `json:"location"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDLaunchSite          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AvailableInclinations respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Description           respjson.Field
		IDLocation            respjson.Field
		LaunchGroup           respjson.Field
		Location              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchSiteDetailFindBySourceResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteDetailFindBySourceResponse) UnmarshalJSON(data []byte) error {
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
type LaunchSiteDetailFindBySourceResponseDataMode string

const (
	LaunchSiteDetailFindBySourceResponseDataModeReal      LaunchSiteDetailFindBySourceResponseDataMode = "REAL"
	LaunchSiteDetailFindBySourceResponseDataModeTest      LaunchSiteDetailFindBySourceResponseDataMode = "TEST"
	LaunchSiteDetailFindBySourceResponseDataModeSimulated LaunchSiteDetailFindBySourceResponseDataMode = "SIMULATED"
	LaunchSiteDetailFindBySourceResponseDataModeExercise  LaunchSiteDetailFindBySourceResponseDataMode = "EXERCISE"
)

// Model representation of a location, which is a specific fixed point on the earth
// and is used to denote the locations of fixed sensors, operating units, etc.
type LaunchSiteDetailFindBySourceResponseLocation struct {
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
	DataMode string `json:"dataMode,required"`
	// Location name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Altitude of the location, in kilometers.
	Altitude float64 `json:"altitude"`
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
	// Unique identifier of the location, auto-generated by the system.
	IDLocation string `json:"idLocation"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		Altitude              respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDLocation            respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchSiteDetailFindBySourceResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteDetailFindBySourceResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of details compiled/collected on a launch site by a
// particular source. A launch site may have several details records.
type LaunchSiteDetailGetResponse struct {
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
	DataMode LaunchSiteDetailGetResponseDataMode `json:"dataMode,required"`
	// Identifier of the parent launch site record.
	IDLaunchSite string `json:"idLaunchSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of supported launch inclinations known for this site. The array is always
	// sized as a multiple of two and includes start/end values to support ranges. For
	// example, if a site support inclinations of 10 and 12-14, the array would have
	// the following values: [10,10, 12,14].
	AvailableInclinations []float64 `json:"availableInclinations"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Notes/description of the launch site.
	Description string `json:"description"`
	// ID of the location data for the launch site, or null if mobile (e.g. sea launch
	// platform in international waters).
	IDLocation string `json:"idLocation"`
	// Launch site group name. Multiple launch sites may be colocated within a launch
	// ”group”.
	LaunchGroup string `json:"launchGroup"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location shared.LocationFull `json:"location"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDLaunchSite          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AvailableInclinations respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Description           respjson.Field
		IDLocation            respjson.Field
		LaunchGroup           respjson.Field
		Location              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Tags                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchSiteDetailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteDetailGetResponse) UnmarshalJSON(data []byte) error {
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
type LaunchSiteDetailGetResponseDataMode string

const (
	LaunchSiteDetailGetResponseDataModeReal      LaunchSiteDetailGetResponseDataMode = "REAL"
	LaunchSiteDetailGetResponseDataModeTest      LaunchSiteDetailGetResponseDataMode = "TEST"
	LaunchSiteDetailGetResponseDataModeSimulated LaunchSiteDetailGetResponseDataMode = "SIMULATED"
	LaunchSiteDetailGetResponseDataModeExercise  LaunchSiteDetailGetResponseDataMode = "EXERCISE"
)

type LaunchSiteDetailNewParams struct {
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
	DataMode LaunchSiteDetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the parent launch site record.
	IDLaunchSite string `json:"idLaunchSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Notes/description of the launch site.
	Description param.Opt[string] `json:"description,omitzero"`
	// ID of the location data for the launch site, or null if mobile (e.g. sea launch
	// platform in international waters).
	IDLocation param.Opt[string] `json:"idLocation,omitzero"`
	// Launch site group name. Multiple launch sites may be colocated within a launch
	// ”group”.
	LaunchGroup param.Opt[string] `json:"launchGroup,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Array of supported launch inclinations known for this site. The array is always
	// sized as a multiple of two and includes start/end values to support ranges. For
	// example, if a site support inclinations of 10 and 12-14, the array would have
	// the following values: [10,10, 12,14].
	AvailableInclinations []float64 `json:"availableInclinations,omitzero"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LocationIngestParam `json:"location,omitzero"`
	paramObj
}

func (r LaunchSiteDetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchSiteDetailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchSiteDetailNewParams) UnmarshalJSON(data []byte) error {
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
type LaunchSiteDetailNewParamsDataMode string

const (
	LaunchSiteDetailNewParamsDataModeReal      LaunchSiteDetailNewParamsDataMode = "REAL"
	LaunchSiteDetailNewParamsDataModeTest      LaunchSiteDetailNewParamsDataMode = "TEST"
	LaunchSiteDetailNewParamsDataModeSimulated LaunchSiteDetailNewParamsDataMode = "SIMULATED"
	LaunchSiteDetailNewParamsDataModeExercise  LaunchSiteDetailNewParamsDataMode = "EXERCISE"
)

type LaunchSiteDetailUpdateParams struct {
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
	DataMode LaunchSiteDetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the parent launch site record.
	IDLaunchSite string `json:"idLaunchSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Notes/description of the launch site.
	Description param.Opt[string] `json:"description,omitzero"`
	// ID of the location data for the launch site, or null if mobile (e.g. sea launch
	// platform in international waters).
	IDLocation param.Opt[string] `json:"idLocation,omitzero"`
	// Launch site group name. Multiple launch sites may be colocated within a launch
	// ”group”.
	LaunchGroup param.Opt[string] `json:"launchGroup,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Array of supported launch inclinations known for this site. The array is always
	// sized as a multiple of two and includes start/end values to support ranges. For
	// example, if a site support inclinations of 10 and 12-14, the array would have
	// the following values: [10,10, 12,14].
	AvailableInclinations []float64 `json:"availableInclinations,omitzero"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LocationIngestParam `json:"location,omitzero"`
	paramObj
}

func (r LaunchSiteDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchSiteDetailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchSiteDetailUpdateParams) UnmarshalJSON(data []byte) error {
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
type LaunchSiteDetailUpdateParamsDataMode string

const (
	LaunchSiteDetailUpdateParamsDataModeReal      LaunchSiteDetailUpdateParamsDataMode = "REAL"
	LaunchSiteDetailUpdateParamsDataModeTest      LaunchSiteDetailUpdateParamsDataMode = "TEST"
	LaunchSiteDetailUpdateParamsDataModeSimulated LaunchSiteDetailUpdateParamsDataMode = "SIMULATED"
	LaunchSiteDetailUpdateParamsDataModeExercise  LaunchSiteDetailUpdateParamsDataMode = "EXERCISE"
)

type LaunchSiteDetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchSiteDetailListParams]'s query parameters as
// `url.Values`.
func (r LaunchSiteDetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchSiteDetailFindBySourceParams struct {
	// The source of the LaunchSiteDetails records to find.
	Source      string           `query:"source,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchSiteDetailFindBySourceParams]'s query parameters as
// `url.Values`.
func (r LaunchSiteDetailFindBySourceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchSiteDetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchSiteDetailGetParams]'s query parameters as
// `url.Values`.
func (r LaunchSiteDetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
