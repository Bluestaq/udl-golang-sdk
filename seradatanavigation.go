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
)

// SeraDataNavigationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeraDataNavigationService] method instead.
type SeraDataNavigationService struct {
	Options []option.RequestOption
}

// NewSeraDataNavigationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSeraDataNavigationService(opts ...option.RequestOption) (r SeraDataNavigationService) {
	r = SeraDataNavigationService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataNavigation as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeraDataNavigationService) New(ctx context.Context, body SeraDataNavigationNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradatanavigation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataNavigation. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *SeraDataNavigationService) Update(ctx context.Context, id string, body SeraDataNavigationUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradatanavigation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SeraDataNavigationService) List(ctx context.Context, query SeraDataNavigationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeraDataNavigationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/seradatanavigation"
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
func (r *SeraDataNavigationService) ListAutoPaging(ctx context.Context, query SeraDataNavigationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeraDataNavigationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataNavigation specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SeraDataNavigationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradatanavigation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SeraDataNavigationService) Count(ctx context.Context, query SeraDataNavigationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradatanavigation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataNavigation by its unique ID passed as
// a path parameter.
func (r *SeraDataNavigationService) Get(ctx context.Context, id string, query SeraDataNavigationGetParams, opts ...option.RequestOption) (res *SeraDataNavigationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradatanavigation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SeraDataNavigationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SeraDataNavigationQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/seradatanavigation/queryhelp"
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
func (r *SeraDataNavigationService) Tuple(ctx context.Context, query SeraDataNavigationTupleParams, opts ...option.RequestOption) (res *[]SeraDataNavigationTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/seradatanavigation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details for a navigation payload from Seradata.
type SeraDataNavigationListResponse struct {
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
	DataMode SeraDataNavigationListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Area of coverage, e.g. Worldwide, India, etc.
	AreaCoverage string `json:"areaCoverage"`
	// Type of clock, e.g. Rubidium, Hydrogen Maser, etc.
	ClockType string `json:"clockType"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the parent Navigation record.
	IDNavigation string `json:"idNavigation"`
	// Location accuracy in meters.
	LocationAccuracy float64 `json:"locationAccuracy"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Frequency for this payload.
	ModeFrequency string `json:"modeFrequency"`
	// Modes of operation.
	Modes string `json:"modes"`
	// Sensor name from Seradata, e.g. WAAS GEO-5, etc.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID string `json:"partnerSpacecraftId"`
	// Navigation payload type, e.g. WAAS, GAGAN, etc.
	PayloadType string `json:"payloadType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		SpacecraftID          respjson.Field
		ID                    respjson.Field
		AreaCoverage          respjson.Field
		ClockType             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		HostedForCompanyOrgID respjson.Field
		IDNavigation          respjson.Field
		LocationAccuracy      respjson.Field
		ManufacturerOrgID     respjson.Field
		ModeFrequency         respjson.Field
		Modes                 respjson.Field
		Name                  respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PartnerSpacecraftID   respjson.Field
		PayloadType           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataNavigationListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataNavigationListResponse) UnmarshalJSON(data []byte) error {
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
type SeraDataNavigationListResponseDataMode string

const (
	SeraDataNavigationListResponseDataModeReal      SeraDataNavigationListResponseDataMode = "REAL"
	SeraDataNavigationListResponseDataModeTest      SeraDataNavigationListResponseDataMode = "TEST"
	SeraDataNavigationListResponseDataModeSimulated SeraDataNavigationListResponseDataMode = "SIMULATED"
	SeraDataNavigationListResponseDataModeExercise  SeraDataNavigationListResponseDataMode = "EXERCISE"
)

// Details for a navigation payload from Seradata.
type SeraDataNavigationGetResponse struct {
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
	DataMode SeraDataNavigationGetResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Area of coverage, e.g. Worldwide, India, etc.
	AreaCoverage string `json:"areaCoverage"`
	// Type of clock, e.g. Rubidium, Hydrogen Maser, etc.
	ClockType string `json:"clockType"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the parent Navigation record.
	IDNavigation string `json:"idNavigation"`
	// Location accuracy in meters.
	LocationAccuracy float64 `json:"locationAccuracy"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Frequency for this payload.
	ModeFrequency string `json:"modeFrequency"`
	// Modes of operation.
	Modes string `json:"modes"`
	// Sensor name from Seradata, e.g. WAAS GEO-5, etc.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID string `json:"partnerSpacecraftId"`
	// Navigation payload type, e.g. WAAS, GAGAN, etc.
	PayloadType string `json:"payloadType"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		SpacecraftID          respjson.Field
		ID                    respjson.Field
		AreaCoverage          respjson.Field
		ClockType             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		HostedForCompanyOrgID respjson.Field
		IDNavigation          respjson.Field
		LocationAccuracy      respjson.Field
		ManufacturerOrgID     respjson.Field
		ModeFrequency         respjson.Field
		Modes                 respjson.Field
		Name                  respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PartnerSpacecraftID   respjson.Field
		PayloadType           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataNavigationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataNavigationGetResponse) UnmarshalJSON(data []byte) error {
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
type SeraDataNavigationGetResponseDataMode string

const (
	SeraDataNavigationGetResponseDataModeReal      SeraDataNavigationGetResponseDataMode = "REAL"
	SeraDataNavigationGetResponseDataModeTest      SeraDataNavigationGetResponseDataMode = "TEST"
	SeraDataNavigationGetResponseDataModeSimulated SeraDataNavigationGetResponseDataMode = "SIMULATED"
	SeraDataNavigationGetResponseDataModeExercise  SeraDataNavigationGetResponseDataMode = "EXERCISE"
)

type SeraDataNavigationQueryhelpResponse struct {
	AodrSupported         bool                                           `json:"aodrSupported"`
	ClassificationMarking string                                         `json:"classificationMarking"`
	Description           string                                         `json:"description"`
	HistorySupported      bool                                           `json:"historySupported"`
	Name                  string                                         `json:"name"`
	Parameters            []SeraDataNavigationQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                       `json:"requiredRoles"`
	RestSupported         bool                                           `json:"restSupported"`
	SortSupported         bool                                           `json:"sortSupported"`
	TypeName              string                                         `json:"typeName"`
	Uri                   string                                         `json:"uri"`
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
func (r SeraDataNavigationQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataNavigationQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SeraDataNavigationQueryhelpResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataNavigationQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *SeraDataNavigationQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details for a navigation payload from Seradata.
type SeraDataNavigationTupleResponse struct {
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
	DataMode SeraDataNavigationTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Area of coverage, e.g. Worldwide, India, etc.
	AreaCoverage string `json:"areaCoverage"`
	// Type of clock, e.g. Rubidium, Hydrogen Maser, etc.
	ClockType string `json:"clockType"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the parent Navigation record.
	IDNavigation string `json:"idNavigation"`
	// Location accuracy in meters.
	LocationAccuracy float64 `json:"locationAccuracy"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Frequency for this payload.
	ModeFrequency string `json:"modeFrequency"`
	// Modes of operation.
	Modes string `json:"modes"`
	// Sensor name from Seradata, e.g. WAAS GEO-5, etc.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID string `json:"partnerSpacecraftId"`
	// Navigation payload type, e.g. WAAS, GAGAN, etc.
	PayloadType string `json:"payloadType"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		SpacecraftID          respjson.Field
		ID                    respjson.Field
		AreaCoverage          respjson.Field
		ClockType             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		HostedForCompanyOrgID respjson.Field
		IDNavigation          respjson.Field
		LocationAccuracy      respjson.Field
		ManufacturerOrgID     respjson.Field
		ModeFrequency         respjson.Field
		Modes                 respjson.Field
		Name                  respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PartnerSpacecraftID   respjson.Field
		PayloadType           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataNavigationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataNavigationTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeraDataNavigationTupleResponseDataMode string

const (
	SeraDataNavigationTupleResponseDataModeReal      SeraDataNavigationTupleResponseDataMode = "REAL"
	SeraDataNavigationTupleResponseDataModeTest      SeraDataNavigationTupleResponseDataMode = "TEST"
	SeraDataNavigationTupleResponseDataModeSimulated SeraDataNavigationTupleResponseDataMode = "SIMULATED"
	SeraDataNavigationTupleResponseDataModeExercise  SeraDataNavigationTupleResponseDataMode = "EXERCISE"
)

type SeraDataNavigationNewParams struct {
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
	DataMode SeraDataNavigationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Area of coverage, e.g. Worldwide, India, etc.
	AreaCoverage param.Opt[string] `json:"areaCoverage,omitzero"`
	// Type of clock, e.g. Rubidium, Hydrogen Maser, etc.
	ClockType param.Opt[string] `json:"clockType,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the parent Navigation record.
	IDNavigation param.Opt[string] `json:"idNavigation,omitzero"`
	// Location accuracy in meters.
	LocationAccuracy param.Opt[float64] `json:"locationAccuracy,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Frequency for this payload.
	ModeFrequency param.Opt[string] `json:"modeFrequency,omitzero"`
	// Modes of operation.
	Modes param.Opt[string] `json:"modes,omitzero"`
	// Sensor name from Seradata, e.g. WAAS GEO-5, etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// Payload notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID param.Opt[string] `json:"partnerSpacecraftId,omitzero"`
	// Navigation payload type, e.g. WAAS, GAGAN, etc.
	PayloadType param.Opt[string] `json:"payloadType,omitzero"`
	paramObj
}

func (r SeraDataNavigationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeraDataNavigationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeraDataNavigationNewParams) UnmarshalJSON(data []byte) error {
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
type SeraDataNavigationNewParamsDataMode string

const (
	SeraDataNavigationNewParamsDataModeReal      SeraDataNavigationNewParamsDataMode = "REAL"
	SeraDataNavigationNewParamsDataModeTest      SeraDataNavigationNewParamsDataMode = "TEST"
	SeraDataNavigationNewParamsDataModeSimulated SeraDataNavigationNewParamsDataMode = "SIMULATED"
	SeraDataNavigationNewParamsDataModeExercise  SeraDataNavigationNewParamsDataMode = "EXERCISE"
)

type SeraDataNavigationUpdateParams struct {
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
	DataMode SeraDataNavigationUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Area of coverage, e.g. Worldwide, India, etc.
	AreaCoverage param.Opt[string] `json:"areaCoverage,omitzero"`
	// Type of clock, e.g. Rubidium, Hydrogen Maser, etc.
	ClockType param.Opt[string] `json:"clockType,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the parent Navigation record.
	IDNavigation param.Opt[string] `json:"idNavigation,omitzero"`
	// Location accuracy in meters.
	LocationAccuracy param.Opt[float64] `json:"locationAccuracy,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Frequency for this payload.
	ModeFrequency param.Opt[string] `json:"modeFrequency,omitzero"`
	// Modes of operation.
	Modes param.Opt[string] `json:"modes,omitzero"`
	// Sensor name from Seradata, e.g. WAAS GEO-5, etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// Payload notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	PartnerSpacecraftID param.Opt[string] `json:"partnerSpacecraftId,omitzero"`
	// Navigation payload type, e.g. WAAS, GAGAN, etc.
	PayloadType param.Opt[string] `json:"payloadType,omitzero"`
	paramObj
}

func (r SeraDataNavigationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeraDataNavigationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeraDataNavigationUpdateParams) UnmarshalJSON(data []byte) error {
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
type SeraDataNavigationUpdateParamsDataMode string

const (
	SeraDataNavigationUpdateParamsDataModeReal      SeraDataNavigationUpdateParamsDataMode = "REAL"
	SeraDataNavigationUpdateParamsDataModeTest      SeraDataNavigationUpdateParamsDataMode = "TEST"
	SeraDataNavigationUpdateParamsDataModeSimulated SeraDataNavigationUpdateParamsDataMode = "SIMULATED"
	SeraDataNavigationUpdateParamsDataModeExercise  SeraDataNavigationUpdateParamsDataMode = "EXERCISE"
)

type SeraDataNavigationListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeraDataNavigationListParams]'s query parameters as
// `url.Values`.
func (r SeraDataNavigationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeraDataNavigationCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeraDataNavigationCountParams]'s query parameters as
// `url.Values`.
func (r SeraDataNavigationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeraDataNavigationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeraDataNavigationGetParams]'s query parameters as
// `url.Values`.
func (r SeraDataNavigationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeraDataNavigationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeraDataNavigationTupleParams]'s query parameters as
// `url.Values`.
func (r SeraDataNavigationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
