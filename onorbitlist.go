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

// OnorbitlistService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitlistService] method instead.
type OnorbitlistService struct {
	Options []option.RequestOption
}

// NewOnorbitlistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOnorbitlistService(opts ...option.RequestOption) (r OnorbitlistService) {
	r = OnorbitlistService{}
	r.Options = opts
	return
}

// Service operation to take a single OnOrbitList as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *OnorbitlistService) New(ctx context.Context, body OnorbitlistNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/onorbitlist"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single OnOrbitList record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *OnorbitlistService) Update(ctx context.Context, id string, body OnorbitlistUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitlist/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitlistService) List(ctx context.Context, query OnorbitlistListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitlistListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitlist"
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
func (r *OnorbitlistService) ListAutoPaging(ctx context.Context, query OnorbitlistListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitlistListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a single OnOrbitList record specified by the passed
// ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *OnorbitlistService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitlist/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnorbitlistService) Count(ctx context.Context, query OnorbitlistCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onorbitlist/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single OnOrbitList record by its unique ID passed as
// a path parameter.
func (r *OnorbitlistService) Get(ctx context.Context, id string, query OnorbitlistGetParams, opts ...option.RequestOption) (res *OnorbitlistGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitlist/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *OnorbitlistService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *OnorbitlistQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/onorbitlist/queryhelp"
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
func (r *OnorbitlistService) Tuple(ctx context.Context, query OnorbitlistTupleParams, opts ...option.RequestOption) (res *[]OnorbitlistTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/onorbitlist/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Table for maintaining generic lists of OnOrbit objects (e.g. Favorites, HIO,
// SHIO, HVA, etc).
type OnorbitlistListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
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
	DataMode OnorbitlistListResponseDataMode `json:"dataMode,required"`
	// Unique name of the list.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Default revisit rate in minutes for all objects in this list.
	DefaultRevisitRateMins float64 `json:"defaultRevisitRateMins"`
	// Description of the list.
	Description string `json:"description"`
	// Numerical priority of this orbit list relative to other orbit lists; lower
	// values indicate higher priority. Decimal values allowed for fine granularity.
	// Consumers should contact the provider for details on the priority.
	ListPriority float64 `json:"listPriority"`
	// Defined naming system that ensures each satellite or space object has a unique
	// and unambiguous identifier within the name space (e.g. JCO, 18SDS). If null, it
	// is assumed to be 18th Space Defense Squadron (18SDS).
	Namespace string `json:"namespace"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
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
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		Name                   respjson.Field
		Source                 respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DefaultRevisitRateMins respjson.Field
		Description            respjson.Field
		ListPriority           respjson.Field
		Namespace              respjson.Field
		Origin                 respjson.Field
		SourceDl               respjson.Field
		Tags                   respjson.Field
		TransactionID          respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitlistListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitlistListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
type OnorbitlistListResponseDataMode string

const (
	OnorbitlistListResponseDataModeReal      OnorbitlistListResponseDataMode = "REAL"
	OnorbitlistListResponseDataModeTest      OnorbitlistListResponseDataMode = "TEST"
	OnorbitlistListResponseDataModeSimulated OnorbitlistListResponseDataMode = "SIMULATED"
	OnorbitlistListResponseDataModeExercise  OnorbitlistListResponseDataMode = "EXERCISE"
)

// Table for maintaining generic lists of OnOrbit objects (e.g. Favorites, HIO,
// SHIO, HVA, etc).
type OnorbitlistGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
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
	DataMode OnorbitlistGetResponseDataMode `json:"dataMode,required"`
	// Unique name of the list.
	Name string `json:"name,required"`
	// This is a list of onOrbitListItems that will be related one-to-one with an
	// onOrbit entry.
	OnOrbitListItems []OnorbitlistGetResponseOnOrbitListItem `json:"onOrbitListItems,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Default revisit rate in minutes for all objects in this list.
	DefaultRevisitRateMins float64 `json:"defaultRevisitRateMins"`
	// Description of the list.
	Description string `json:"description"`
	// Numerical priority of this orbit list relative to other orbit lists; lower
	// values indicate higher priority. Decimal values allowed for fine granularity.
	// Consumers should contact the provider for details on the priority.
	ListPriority float64 `json:"listPriority"`
	// Defined naming system that ensures each satellite or space object has a unique
	// and unambiguous identifier within the name space (e.g. JCO, 18SDS). If null, it
	// is assumed to be 18th Space Defense Squadron (18SDS).
	Namespace string `json:"namespace"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
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
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		Name                   respjson.Field
		OnOrbitListItems       respjson.Field
		Source                 respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DefaultRevisitRateMins respjson.Field
		Description            respjson.Field
		ListPriority           respjson.Field
		Namespace              respjson.Field
		Origin                 respjson.Field
		SourceDl               respjson.Field
		Tags                   respjson.Field
		TransactionID          respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitlistGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitlistGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
type OnorbitlistGetResponseDataMode string

const (
	OnorbitlistGetResponseDataModeReal      OnorbitlistGetResponseDataMode = "REAL"
	OnorbitlistGetResponseDataModeTest      OnorbitlistGetResponseDataMode = "TEST"
	OnorbitlistGetResponseDataModeSimulated OnorbitlistGetResponseDataMode = "SIMULATED"
	OnorbitlistGetResponseDataModeExercise  OnorbitlistGetResponseDataMode = "EXERCISE"
)

// Items associated with this onOrbitList record.
type OnorbitlistGetResponseOnOrbitListItem struct {
	// Height of a box, in degrees, volume expected to be cleared by sensor providers,
	// if CLEARING is selected.
	ClearingBoxCrossTrack float64 `json:"clearingBoxCrossTrack"`
	// Width of a box volume, in degrees, expected to be cleared by sensor providers,
	// if CLEARING is selected.
	ClearingBoxInTrack float64 `json:"clearingBoxInTrack"`
	// Radius, in degrees, of a spherical volume expected to be cleared by sensor
	// providers, if CLEARING is selected.
	ClearingRadius float64 `json:"clearingRadius"`
	// Common name of the onorbit object.
	CommonName string `json:"commonName"`
	// This value is typically the ISO 3166 Alpha-3 three-character country code,
	// however it can also represent various consortiums that do not appear in the ISO
	// document.
	CountryCode string `json:"countryCode"`
	// Datetime expiration of a satellite on this list, allowing for the maintenance of
	// a history of when satellites entered and when they exited the list in ISO 8601
	// UTC datetime format with millisecond precision.
	ExpiredOn time.Time `json:"expiredOn" format:"date-time"`
	// Frequency of additional routine, in minutes, tasking identified in and
	// corresponding to the monitoringType array.
	FreqMins float64 `json:"freqMins"`
	// Routine tasking that should be applied to this object. REVISIT_RATE allows users
	// to define custom revisit rates for individual satellites, HVA_CLEARING allows
	// users to define custom volumes that are expected to be clear of unknown objects,
	// and POL would be collects on a specified increment in support of collecting data
	// that feeds into Pattern of Life (PoL) assessments.
	MonitoringType string `json:"monitoringType"`
	// Unique identifier of the on-orbit object. This is typically the USSF 18th SDS
	// satellite number (also sometimes referred to as NORAD ID/number) but could be an
	// identifier from another satellite catalog namespace. See the ‘namespace’ field
	// for the appropriate identifier namespace. If namespace is null, 18SDS satellite
	// number is assumed.
	ObjectID string `json:"objectId"`
	// Orbit Regime refers to a classification of a satellite's orbit based on its
	// altitude, inclination, and other orbital characteristics. Common orbit regimes
	// include Low Earth Orbit (LEO), Medium Earth Orbit (MEO), Geostationary Orbit
	// (GEO), and Highly Elliptical Orbit (HEO).
	OrbitRegime string `json:"orbitRegime"`
	// Optional identifier indicates the on-orbit object being referenced. This may be
	// an internal system identifier and not necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Payload priority based on the type of payload that has been identified or that
	// is suspected. Priority of the payload as a number. (1=highest priority).
	PayloadPriority float64 `json:"payloadPriority"`
	// Rank refers to the assigned position or level of priority given to a satellite
	// based on its relative importance, urgency, or operational relevance as
	// determined by the applicable operations unit.
	Rank int64 `json:"rank"`
	// Tasking urgency, typically will be on a 1-10 scale. Urgency as a number.
	// (1=highest priority).
	Urgency float64 `json:"urgency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClearingBoxCrossTrack respjson.Field
		ClearingBoxInTrack    respjson.Field
		ClearingRadius        respjson.Field
		CommonName            respjson.Field
		CountryCode           respjson.Field
		ExpiredOn             respjson.Field
		FreqMins              respjson.Field
		MonitoringType        respjson.Field
		ObjectID              respjson.Field
		OrbitRegime           respjson.Field
		OrigObjectID          respjson.Field
		PayloadPriority       respjson.Field
		Rank                  respjson.Field
		Urgency               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitlistGetResponseOnOrbitListItem) RawJSON() string { return r.JSON.raw }
func (r *OnorbitlistGetResponseOnOrbitListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbitlistQueryhelpResponse struct {
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
func (r OnorbitlistQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitlistQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Table for maintaining generic lists of OnOrbit objects (e.g. Favorites, HIO,
// SHIO, HVA, etc).
type OnorbitlistTupleResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
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
	DataMode OnorbitlistTupleResponseDataMode `json:"dataMode,required"`
	// Unique name of the list.
	Name string `json:"name,required"`
	// This is a list of onOrbitListItems that will be related one-to-one with an
	// onOrbit entry.
	OnOrbitListItems []OnorbitlistTupleResponseOnOrbitListItem `json:"onOrbitListItems,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Default revisit rate in minutes for all objects in this list.
	DefaultRevisitRateMins float64 `json:"defaultRevisitRateMins"`
	// Description of the list.
	Description string `json:"description"`
	// Numerical priority of this orbit list relative to other orbit lists; lower
	// values indicate higher priority. Decimal values allowed for fine granularity.
	// Consumers should contact the provider for details on the priority.
	ListPriority float64 `json:"listPriority"`
	// Defined naming system that ensures each satellite or space object has a unique
	// and unambiguous identifier within the name space (e.g. JCO, 18SDS). If null, it
	// is assumed to be 18th Space Defense Squadron (18SDS).
	Namespace string `json:"namespace"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
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
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		Name                   respjson.Field
		OnOrbitListItems       respjson.Field
		Source                 respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DefaultRevisitRateMins respjson.Field
		Description            respjson.Field
		ListPriority           respjson.Field
		Namespace              respjson.Field
		Origin                 respjson.Field
		SourceDl               respjson.Field
		Tags                   respjson.Field
		TransactionID          respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitlistTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitlistTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
type OnorbitlistTupleResponseDataMode string

const (
	OnorbitlistTupleResponseDataModeReal      OnorbitlistTupleResponseDataMode = "REAL"
	OnorbitlistTupleResponseDataModeTest      OnorbitlistTupleResponseDataMode = "TEST"
	OnorbitlistTupleResponseDataModeSimulated OnorbitlistTupleResponseDataMode = "SIMULATED"
	OnorbitlistTupleResponseDataModeExercise  OnorbitlistTupleResponseDataMode = "EXERCISE"
)

// Items associated with this onOrbitList record.
type OnorbitlistTupleResponseOnOrbitListItem struct {
	// Height of a box, in degrees, volume expected to be cleared by sensor providers,
	// if CLEARING is selected.
	ClearingBoxCrossTrack float64 `json:"clearingBoxCrossTrack"`
	// Width of a box volume, in degrees, expected to be cleared by sensor providers,
	// if CLEARING is selected.
	ClearingBoxInTrack float64 `json:"clearingBoxInTrack"`
	// Radius, in degrees, of a spherical volume expected to be cleared by sensor
	// providers, if CLEARING is selected.
	ClearingRadius float64 `json:"clearingRadius"`
	// Common name of the onorbit object.
	CommonName string `json:"commonName"`
	// This value is typically the ISO 3166 Alpha-3 three-character country code,
	// however it can also represent various consortiums that do not appear in the ISO
	// document.
	CountryCode string `json:"countryCode"`
	// Datetime expiration of a satellite on this list, allowing for the maintenance of
	// a history of when satellites entered and when they exited the list in ISO 8601
	// UTC datetime format with millisecond precision.
	ExpiredOn time.Time `json:"expiredOn" format:"date-time"`
	// Frequency of additional routine, in minutes, tasking identified in and
	// corresponding to the monitoringType array.
	FreqMins float64 `json:"freqMins"`
	// Routine tasking that should be applied to this object. REVISIT_RATE allows users
	// to define custom revisit rates for individual satellites, HVA_CLEARING allows
	// users to define custom volumes that are expected to be clear of unknown objects,
	// and POL would be collects on a specified increment in support of collecting data
	// that feeds into Pattern of Life (PoL) assessments.
	MonitoringType string `json:"monitoringType"`
	// Unique identifier of the on-orbit object. This is typically the USSF 18th SDS
	// satellite number (also sometimes referred to as NORAD ID/number) but could be an
	// identifier from another satellite catalog namespace. See the ‘namespace’ field
	// for the appropriate identifier namespace. If namespace is null, 18SDS satellite
	// number is assumed.
	ObjectID string `json:"objectId"`
	// Orbit Regime refers to a classification of a satellite's orbit based on its
	// altitude, inclination, and other orbital characteristics. Common orbit regimes
	// include Low Earth Orbit (LEO), Medium Earth Orbit (MEO), Geostationary Orbit
	// (GEO), and Highly Elliptical Orbit (HEO).
	OrbitRegime string `json:"orbitRegime"`
	// Optional identifier indicates the on-orbit object being referenced. This may be
	// an internal system identifier and not necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Payload priority based on the type of payload that has been identified or that
	// is suspected. Priority of the payload as a number. (1=highest priority).
	PayloadPriority float64 `json:"payloadPriority"`
	// Rank refers to the assigned position or level of priority given to a satellite
	// based on its relative importance, urgency, or operational relevance as
	// determined by the applicable operations unit.
	Rank int64 `json:"rank"`
	// Tasking urgency, typically will be on a 1-10 scale. Urgency as a number.
	// (1=highest priority).
	Urgency float64 `json:"urgency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClearingBoxCrossTrack respjson.Field
		ClearingBoxInTrack    respjson.Field
		ClearingRadius        respjson.Field
		CommonName            respjson.Field
		CountryCode           respjson.Field
		ExpiredOn             respjson.Field
		FreqMins              respjson.Field
		MonitoringType        respjson.Field
		ObjectID              respjson.Field
		OrbitRegime           respjson.Field
		OrigObjectID          respjson.Field
		PayloadPriority       respjson.Field
		Rank                  respjson.Field
		Urgency               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitlistTupleResponseOnOrbitListItem) RawJSON() string { return r.JSON.raw }
func (r *OnorbitlistTupleResponseOnOrbitListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbitlistNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
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
	DataMode OnorbitlistNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique name of the list.
	Name string `json:"name,required"`
	// This is a list of onOrbitListItems that will be related one-to-one with an
	// onOrbit entry.
	OnOrbitListItems []OnorbitlistNewParamsOnOrbitListItem `json:"onOrbitListItems,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Default revisit rate in minutes for all objects in this list.
	DefaultRevisitRateMins param.Opt[float64] `json:"defaultRevisitRateMins,omitzero"`
	// Description of the list.
	Description param.Opt[string] `json:"description,omitzero"`
	// Numerical priority of this orbit list relative to other orbit lists; lower
	// values indicate higher priority. Decimal values allowed for fine granularity.
	// Consumers should contact the provider for details on the priority.
	ListPriority param.Opt[float64] `json:"listPriority,omitzero"`
	// Defined naming system that ensures each satellite or space object has a unique
	// and unambiguous identifier within the name space (e.g. JCO, 18SDS). If null, it
	// is assumed to be 18th Space Defense Squadron (18SDS).
	Namespace param.Opt[string] `json:"namespace,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
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

func (r OnorbitlistNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitlistNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitlistNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
type OnorbitlistNewParamsDataMode string

const (
	OnorbitlistNewParamsDataModeReal      OnorbitlistNewParamsDataMode = "REAL"
	OnorbitlistNewParamsDataModeTest      OnorbitlistNewParamsDataMode = "TEST"
	OnorbitlistNewParamsDataModeSimulated OnorbitlistNewParamsDataMode = "SIMULATED"
	OnorbitlistNewParamsDataModeExercise  OnorbitlistNewParamsDataMode = "EXERCISE"
)

// Items associated with this onOrbitList record.
type OnorbitlistNewParamsOnOrbitListItem struct {
	// Height of a box, in degrees, volume expected to be cleared by sensor providers,
	// if CLEARING is selected.
	ClearingBoxCrossTrack param.Opt[float64] `json:"clearingBoxCrossTrack,omitzero"`
	// Width of a box volume, in degrees, expected to be cleared by sensor providers,
	// if CLEARING is selected.
	ClearingBoxInTrack param.Opt[float64] `json:"clearingBoxInTrack,omitzero"`
	// Radius, in degrees, of a spherical volume expected to be cleared by sensor
	// providers, if CLEARING is selected.
	ClearingRadius param.Opt[float64] `json:"clearingRadius,omitzero"`
	// Common name of the onorbit object.
	CommonName param.Opt[string] `json:"commonName,omitzero"`
	// This value is typically the ISO 3166 Alpha-3 three-character country code,
	// however it can also represent various consortiums that do not appear in the ISO
	// document.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Datetime expiration of a satellite on this list, allowing for the maintenance of
	// a history of when satellites entered and when they exited the list in ISO 8601
	// UTC datetime format with millisecond precision.
	ExpiredOn param.Opt[time.Time] `json:"expiredOn,omitzero" format:"date-time"`
	// Frequency of additional routine, in minutes, tasking identified in and
	// corresponding to the monitoringType array.
	FreqMins param.Opt[float64] `json:"freqMins,omitzero"`
	// Routine tasking that should be applied to this object. REVISIT_RATE allows users
	// to define custom revisit rates for individual satellites, HVA_CLEARING allows
	// users to define custom volumes that are expected to be clear of unknown objects,
	// and POL would be collects on a specified increment in support of collecting data
	// that feeds into Pattern of Life (PoL) assessments.
	MonitoringType param.Opt[string] `json:"monitoringType,omitzero"`
	// Unique identifier of the on-orbit object. This is typically the USSF 18th SDS
	// satellite number (also sometimes referred to as NORAD ID/number) but could be an
	// identifier from another satellite catalog namespace. See the ‘namespace’ field
	// for the appropriate identifier namespace. If namespace is null, 18SDS satellite
	// number is assumed.
	ObjectID param.Opt[string] `json:"objectId,omitzero"`
	// Orbit Regime refers to a classification of a satellite's orbit based on its
	// altitude, inclination, and other orbital characteristics. Common orbit regimes
	// include Low Earth Orbit (LEO), Medium Earth Orbit (MEO), Geostationary Orbit
	// (GEO), and Highly Elliptical Orbit (HEO).
	OrbitRegime param.Opt[string] `json:"orbitRegime,omitzero"`
	// Optional identifier indicates the on-orbit object being referenced. This may be
	// an internal system identifier and not necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Payload priority based on the type of payload that has been identified or that
	// is suspected. Priority of the payload as a number. (1=highest priority).
	PayloadPriority param.Opt[float64] `json:"payloadPriority,omitzero"`
	// Rank refers to the assigned position or level of priority given to a satellite
	// based on its relative importance, urgency, or operational relevance as
	// determined by the applicable operations unit.
	Rank param.Opt[int64] `json:"rank,omitzero"`
	// Tasking urgency, typically will be on a 1-10 scale. Urgency as a number.
	// (1=highest priority).
	Urgency param.Opt[float64] `json:"urgency,omitzero"`
	paramObj
}

func (r OnorbitlistNewParamsOnOrbitListItem) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitlistNewParamsOnOrbitListItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitlistNewParamsOnOrbitListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbitlistUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
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
	DataMode OnorbitlistUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique name of the list.
	Name string `json:"name,required"`
	// This is a list of onOrbitListItems that will be related one-to-one with an
	// onOrbit entry.
	OnOrbitListItems []OnorbitlistUpdateParamsOnOrbitListItem `json:"onOrbitListItems,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Default revisit rate in minutes for all objects in this list.
	DefaultRevisitRateMins param.Opt[float64] `json:"defaultRevisitRateMins,omitzero"`
	// Description of the list.
	Description param.Opt[string] `json:"description,omitzero"`
	// Numerical priority of this orbit list relative to other orbit lists; lower
	// values indicate higher priority. Decimal values allowed for fine granularity.
	// Consumers should contact the provider for details on the priority.
	ListPriority param.Opt[float64] `json:"listPriority,omitzero"`
	// Defined naming system that ensures each satellite or space object has a unique
	// and unambiguous identifier within the name space (e.g. JCO, 18SDS). If null, it
	// is assumed to be 18th Space Defense Squadron (18SDS).
	Namespace param.Opt[string] `json:"namespace,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
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

func (r OnorbitlistUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitlistUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitlistUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
type OnorbitlistUpdateParamsDataMode string

const (
	OnorbitlistUpdateParamsDataModeReal      OnorbitlistUpdateParamsDataMode = "REAL"
	OnorbitlistUpdateParamsDataModeTest      OnorbitlistUpdateParamsDataMode = "TEST"
	OnorbitlistUpdateParamsDataModeSimulated OnorbitlistUpdateParamsDataMode = "SIMULATED"
	OnorbitlistUpdateParamsDataModeExercise  OnorbitlistUpdateParamsDataMode = "EXERCISE"
)

// Items associated with this onOrbitList record.
type OnorbitlistUpdateParamsOnOrbitListItem struct {
	// Height of a box, in degrees, volume expected to be cleared by sensor providers,
	// if CLEARING is selected.
	ClearingBoxCrossTrack param.Opt[float64] `json:"clearingBoxCrossTrack,omitzero"`
	// Width of a box volume, in degrees, expected to be cleared by sensor providers,
	// if CLEARING is selected.
	ClearingBoxInTrack param.Opt[float64] `json:"clearingBoxInTrack,omitzero"`
	// Radius, in degrees, of a spherical volume expected to be cleared by sensor
	// providers, if CLEARING is selected.
	ClearingRadius param.Opt[float64] `json:"clearingRadius,omitzero"`
	// Common name of the onorbit object.
	CommonName param.Opt[string] `json:"commonName,omitzero"`
	// This value is typically the ISO 3166 Alpha-3 three-character country code,
	// however it can also represent various consortiums that do not appear in the ISO
	// document.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Datetime expiration of a satellite on this list, allowing for the maintenance of
	// a history of when satellites entered and when they exited the list in ISO 8601
	// UTC datetime format with millisecond precision.
	ExpiredOn param.Opt[time.Time] `json:"expiredOn,omitzero" format:"date-time"`
	// Frequency of additional routine, in minutes, tasking identified in and
	// corresponding to the monitoringType array.
	FreqMins param.Opt[float64] `json:"freqMins,omitzero"`
	// Routine tasking that should be applied to this object. REVISIT_RATE allows users
	// to define custom revisit rates for individual satellites, HVA_CLEARING allows
	// users to define custom volumes that are expected to be clear of unknown objects,
	// and POL would be collects on a specified increment in support of collecting data
	// that feeds into Pattern of Life (PoL) assessments.
	MonitoringType param.Opt[string] `json:"monitoringType,omitzero"`
	// Unique identifier of the on-orbit object. This is typically the USSF 18th SDS
	// satellite number (also sometimes referred to as NORAD ID/number) but could be an
	// identifier from another satellite catalog namespace. See the ‘namespace’ field
	// for the appropriate identifier namespace. If namespace is null, 18SDS satellite
	// number is assumed.
	ObjectID param.Opt[string] `json:"objectId,omitzero"`
	// Orbit Regime refers to a classification of a satellite's orbit based on its
	// altitude, inclination, and other orbital characteristics. Common orbit regimes
	// include Low Earth Orbit (LEO), Medium Earth Orbit (MEO), Geostationary Orbit
	// (GEO), and Highly Elliptical Orbit (HEO).
	OrbitRegime param.Opt[string] `json:"orbitRegime,omitzero"`
	// Optional identifier indicates the on-orbit object being referenced. This may be
	// an internal system identifier and not necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Payload priority based on the type of payload that has been identified or that
	// is suspected. Priority of the payload as a number. (1=highest priority).
	PayloadPriority param.Opt[float64] `json:"payloadPriority,omitzero"`
	// Rank refers to the assigned position or level of priority given to a satellite
	// based on its relative importance, urgency, or operational relevance as
	// determined by the applicable operations unit.
	Rank param.Opt[int64] `json:"rank,omitzero"`
	// Tasking urgency, typically will be on a 1-10 scale. Urgency as a number.
	// (1=highest priority).
	Urgency param.Opt[float64] `json:"urgency,omitzero"`
	paramObj
}

func (r OnorbitlistUpdateParamsOnOrbitListItem) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitlistUpdateParamsOnOrbitListItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnorbitlistUpdateParamsOnOrbitListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbitlistListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitlistListParams]'s query parameters as `url.Values`.
func (r OnorbitlistListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitlistCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitlistCountParams]'s query parameters as `url.Values`.
func (r OnorbitlistCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitlistGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitlistGetParams]'s query parameters as `url.Values`.
func (r OnorbitlistGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitlistTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitlistTupleParams]'s query parameters as `url.Values`.
func (r OnorbitlistTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
