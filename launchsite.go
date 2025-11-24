// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
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

// LaunchSiteService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaunchSiteService] method instead.
type LaunchSiteService struct {
	Options []option.RequestOption
}

// NewLaunchSiteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaunchSiteService(opts ...option.RequestOption) (r LaunchSiteService) {
	r = LaunchSiteService{}
	r.Options = opts
	return
}

// Service operation to take a single launch site as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LaunchSiteService) New(ctx context.Context, body LaunchSiteNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/launchsite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a launch. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *LaunchSiteService) Update(ctx context.Context, id string, body LaunchSiteUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchsite/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaunchSiteService) List(ctx context.Context, query LaunchSiteListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaunchSiteListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/launchsite"
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
func (r *LaunchSiteService) ListAutoPaging(ctx context.Context, query LaunchSiteListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaunchSiteListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a launch site specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *LaunchSiteService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchsite/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LaunchSiteService) Count(ctx context.Context, query LaunchSiteCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/launchsite/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single launch site by its unique ID passed as a path
// parameter.
func (r *LaunchSiteService) Get(ctx context.Context, id string, query LaunchSiteGetParams, opts ...option.RequestOption) (res *LaunchSiteGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchsite/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *LaunchSiteService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *LaunchSiteQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/launchsite/queryhelp"
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
func (r *LaunchSiteService) Tuple(ctx context.Context, query LaunchSiteTupleParams, opts ...option.RequestOption) (res *[]LaunchSiteTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/launchsite/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of known launch sites.
type LaunchSiteListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Launch site code.
	Code string `json:"code,required"`
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
	DataMode LaunchSiteListResponseDataMode `json:"dataMode,required"`
	// Launch site name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate launch site code.
	AltCode string `json:"altCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent site associated with this launch site.
	IDSite string `json:"idSite"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Abbreviated short name for the launch site.
	ShortCode string `json:"shortCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Code                  respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AltCode               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDSite                respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ShortCode             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchSiteListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteListResponse) UnmarshalJSON(data []byte) error {
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
type LaunchSiteListResponseDataMode string

const (
	LaunchSiteListResponseDataModeReal      LaunchSiteListResponseDataMode = "REAL"
	LaunchSiteListResponseDataModeTest      LaunchSiteListResponseDataMode = "TEST"
	LaunchSiteListResponseDataModeSimulated LaunchSiteListResponseDataMode = "SIMULATED"
	LaunchSiteListResponseDataModeExercise  LaunchSiteListResponseDataMode = "EXERCISE"
)

// Model representation of known launch sites.
type LaunchSiteGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Launch site code.
	Code string `json:"code,required"`
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
	DataMode LaunchSiteGetResponseDataMode `json:"dataMode,required"`
	// Launch site name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate launch site code.
	AltCode string `json:"altCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent site associated with this launch site.
	IDSite string `json:"idSite"`
	// Read-only collection of additional LaunchSiteDetails by various sources for this
	// launch site, ignored on create/update. These details must be created separately
	// via the /udl/launchsitedetails operations.
	LaunchSiteDetails []LaunchSiteGetResponseLaunchSiteDetail `json:"launchSiteDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Abbreviated short name for the launch site.
	ShortCode string `json:"shortCode"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Code                  respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AltCode               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDSite                respjson.Field
		LaunchSiteDetails     respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ShortCode             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchSiteGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteGetResponse) UnmarshalJSON(data []byte) error {
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
type LaunchSiteGetResponseDataMode string

const (
	LaunchSiteGetResponseDataModeReal      LaunchSiteGetResponseDataMode = "REAL"
	LaunchSiteGetResponseDataModeTest      LaunchSiteGetResponseDataMode = "TEST"
	LaunchSiteGetResponseDataModeSimulated LaunchSiteGetResponseDataMode = "SIMULATED"
	LaunchSiteGetResponseDataModeExercise  LaunchSiteGetResponseDataMode = "EXERCISE"
)

// Model representation of details compiled/collected on a launch site by a
// particular source. A launch site may have several details records.
type LaunchSiteGetResponseLaunchSiteDetail struct {
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
func (r LaunchSiteGetResponseLaunchSiteDetail) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteGetResponseLaunchSiteDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaunchSiteQueryhelpResponse struct {
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
func (r LaunchSiteQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of known launch sites.
type LaunchSiteTupleResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Launch site code.
	Code string `json:"code,required"`
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
	DataMode LaunchSiteTupleResponseDataMode `json:"dataMode,required"`
	// Launch site name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate launch site code.
	AltCode string `json:"altCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent site associated with this launch site.
	IDSite string `json:"idSite"`
	// Read-only collection of additional LaunchSiteDetails by various sources for this
	// launch site, ignored on create/update. These details must be created separately
	// via the /udl/launchsitedetails operations.
	LaunchSiteDetails []LaunchSiteTupleResponseLaunchSiteDetail `json:"launchSiteDetails"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Abbreviated short name for the launch site.
	ShortCode string `json:"shortCode"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Code                  respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AltCode               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		IDSite                respjson.Field
		LaunchSiteDetails     respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ShortCode             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchSiteTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteTupleResponse) UnmarshalJSON(data []byte) error {
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
type LaunchSiteTupleResponseDataMode string

const (
	LaunchSiteTupleResponseDataModeReal      LaunchSiteTupleResponseDataMode = "REAL"
	LaunchSiteTupleResponseDataModeTest      LaunchSiteTupleResponseDataMode = "TEST"
	LaunchSiteTupleResponseDataModeSimulated LaunchSiteTupleResponseDataMode = "SIMULATED"
	LaunchSiteTupleResponseDataModeExercise  LaunchSiteTupleResponseDataMode = "EXERCISE"
)

// Model representation of details compiled/collected on a launch site by a
// particular source. A launch site may have several details records.
type LaunchSiteTupleResponseLaunchSiteDetail struct {
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
func (r LaunchSiteTupleResponseLaunchSiteDetail) RawJSON() string { return r.JSON.raw }
func (r *LaunchSiteTupleResponseLaunchSiteDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaunchSiteNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Launch site code.
	Code string `json:"code,required"`
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
	DataMode LaunchSiteNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Launch site name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate launch site code.
	AltCode param.Opt[string] `json:"altCode,omitzero"`
	// Unique identifier of the parent site associated with this launch site.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Abbreviated short name for the launch site.
	ShortCode param.Opt[string] `json:"shortCode,omitzero"`
	paramObj
}

func (r LaunchSiteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchSiteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchSiteNewParams) UnmarshalJSON(data []byte) error {
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
type LaunchSiteNewParamsDataMode string

const (
	LaunchSiteNewParamsDataModeReal      LaunchSiteNewParamsDataMode = "REAL"
	LaunchSiteNewParamsDataModeTest      LaunchSiteNewParamsDataMode = "TEST"
	LaunchSiteNewParamsDataModeSimulated LaunchSiteNewParamsDataMode = "SIMULATED"
	LaunchSiteNewParamsDataModeExercise  LaunchSiteNewParamsDataMode = "EXERCISE"
)

type LaunchSiteUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Launch site code.
	Code string `json:"code,required"`
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
	DataMode LaunchSiteUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Launch site name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate launch site code.
	AltCode param.Opt[string] `json:"altCode,omitzero"`
	// Unique identifier of the parent site associated with this launch site.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Abbreviated short name for the launch site.
	ShortCode param.Opt[string] `json:"shortCode,omitzero"`
	paramObj
}

func (r LaunchSiteUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchSiteUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchSiteUpdateParams) UnmarshalJSON(data []byte) error {
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
type LaunchSiteUpdateParamsDataMode string

const (
	LaunchSiteUpdateParamsDataModeReal      LaunchSiteUpdateParamsDataMode = "REAL"
	LaunchSiteUpdateParamsDataModeTest      LaunchSiteUpdateParamsDataMode = "TEST"
	LaunchSiteUpdateParamsDataModeSimulated LaunchSiteUpdateParamsDataMode = "SIMULATED"
	LaunchSiteUpdateParamsDataModeExercise  LaunchSiteUpdateParamsDataMode = "EXERCISE"
)

type LaunchSiteListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchSiteListParams]'s query parameters as `url.Values`.
func (r LaunchSiteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchSiteCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchSiteCountParams]'s query parameters as `url.Values`.
func (r LaunchSiteCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchSiteGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchSiteGetParams]'s query parameters as `url.Values`.
func (r LaunchSiteGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchSiteTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchSiteTupleParams]'s query parameters as `url.Values`.
func (r LaunchSiteTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
