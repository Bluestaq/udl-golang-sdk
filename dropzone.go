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

// DropzoneService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDropzoneService] method instead.
type DropzoneService struct {
	Options []option.RequestOption
}

// NewDropzoneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDropzoneService(opts ...option.RequestOption) (r DropzoneService) {
	r = DropzoneService{}
	r.Options = opts
	return
}

// Service operation to take a single dropzone record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *DropzoneService) New(ctx context.Context, body DropzoneNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/dropzone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single dropzone record by its unique ID passed as a
// path parameter.
func (r *DropzoneService) Get(ctx context.Context, id string, query DropzoneGetParams, opts ...option.RequestOption) (res *DropzoneGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/dropzone/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single dropzone record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *DropzoneService) Update(ctx context.Context, id string, body DropzoneUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/dropzone/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *DropzoneService) List(ctx context.Context, query DropzoneListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[DropzoneListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/dropzone"
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
func (r *DropzoneService) ListAutoPaging(ctx context.Context, query DropzoneListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[DropzoneListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a dropzone record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *DropzoneService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/dropzone/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *DropzoneService) Count(ctx context.Context, query DropzoneCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/dropzone/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// dropzone records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *DropzoneService) NewBulk(ctx context.Context, body DropzoneNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/dropzone/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *DropzoneService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *DropzoneQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/dropzone/queryhelp"
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
func (r *DropzoneService) Tuple(ctx context.Context, query DropzoneTupleParams, opts ...option.RequestOption) (res *[]DropzoneTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/dropzone/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple dropzone records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *DropzoneService) UnvalidatedPublish(ctx context.Context, body DropzoneUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-dropzone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Properties and characteristics of a Drop Zone, including name, location, shape,
// type code, survey date, and remarks.
type DropzoneGetResponse struct {
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
	DataMode DropzoneGetResponseDataMode `json:"dataMode,required"`
	// WGS84 latitude of the drop zone, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the drop zone, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// The name of the drop zone.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Specifies an alternate country code for the drop zone if the data provider code
	// is not part of an official Country Code standard such as ISO-3166 or FIPS. This
	// field will be set to the value provided by the source and should be used for all
	// Queries specifying a Country Code.
	AltCountryCode string `json:"altCountryCode"`
	// Specifies the country name associated with the source provided alternate country
	// code.
	AltCountryName string `json:"altCountryName"`
	// The date the drop zone survey was approved, in ISO 8601 UTC format with
	// millisecond precision.
	ApprovalDate time.Time `json:"approvalDate" format:"date-time"`
	// The type code for the drop zone.
	Code string `json:"code"`
	// The Country Code where the drop zone is located. This value is typically the ISO
	// 3166 Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode string `json:"countryCode"`
	// The country name of the location for the drop zone.
	CountryName string `json:"countryName"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The survey expiration date of the drop zone, in ISO 8601 UTC format with
	// millisecond precision.
	ExpirationDate time.Time `json:"expirationDate" format:"date-time"`
	// The external identifier assigned to the drop zone.
	ExtIdentifier string `json:"extIdentifier"`
	// The ID of the site associated with the drop zone.
	IDSite string `json:"idSite"`
	// Last time the drop zone information was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdate time.Time `json:"lastUpdate" format:"date-time"`
	// The length dimension of the drop zone in meters for non-circular drop zones.
	Length float64 `json:"length"`
	// The Major Command (MAJCOM) responsible for management of the drop zone.
	Majcom string `json:"majcom"`
	// The nearest reference location to the drop zone.
	NearestLoc string `json:"nearestLoc"`
	// The approval date for the drop zone by an air drop authority certifying
	// operational usage, in ISO 8601 UTC format with millisecond precision.
	OperationalApprovalDate time.Time `json:"operationalApprovalDate" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The name assigned to the drop zone point.
	PointName string `json:"pointName"`
	// The radius dimension of the drop zone in meters for circular drop zones.
	Radius float64 `json:"radius"`
	// The date the drop zone was recertified, in ISO 8601 UTC format with millisecond
	// precision.
	RecertDate time.Time `json:"recertDate" format:"date-time"`
	// Remarks concerning the drop zone.
	Remark string `json:"remark"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The US alphabetical code for the state where the drop zone is located.
	StateAbbr string `json:"stateAbbr"`
	// The name of the state where the drop zone is located.
	StateName string `json:"stateName"`
	// The date the drop zone survey was performed, in ISO 8601 UTC format with
	// millisecond precision.
	SurveyDate time.Time `json:"surveyDate" format:"date-time"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The width dimension of the drop zone in meters for non-circular drop zones.
	Width float64 `json:"width"`
	// The identifier of the Zone Availability Report (ZAR) for the drop zone.
	ZarID string `json:"zarId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking   respjson.Field
		DataMode                respjson.Field
		Lat                     respjson.Field
		Lon                     respjson.Field
		Name                    respjson.Field
		Source                  respjson.Field
		ID                      respjson.Field
		AltCountryCode          respjson.Field
		AltCountryName          respjson.Field
		ApprovalDate            respjson.Field
		Code                    respjson.Field
		CountryCode             respjson.Field
		CountryName             respjson.Field
		CreatedAt               respjson.Field
		CreatedBy               respjson.Field
		ExpirationDate          respjson.Field
		ExtIdentifier           respjson.Field
		IDSite                  respjson.Field
		LastUpdate              respjson.Field
		Length                  respjson.Field
		Majcom                  respjson.Field
		NearestLoc              respjson.Field
		OperationalApprovalDate respjson.Field
		Origin                  respjson.Field
		OrigNetwork             respjson.Field
		PointName               respjson.Field
		Radius                  respjson.Field
		RecertDate              respjson.Field
		Remark                  respjson.Field
		SourceDl                respjson.Field
		StateAbbr               respjson.Field
		StateName               respjson.Field
		SurveyDate              respjson.Field
		UpdatedAt               respjson.Field
		UpdatedBy               respjson.Field
		Width                   respjson.Field
		ZarID                   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DropzoneGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DropzoneGetResponse) UnmarshalJSON(data []byte) error {
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
type DropzoneGetResponseDataMode string

const (
	DropzoneGetResponseDataModeReal      DropzoneGetResponseDataMode = "REAL"
	DropzoneGetResponseDataModeTest      DropzoneGetResponseDataMode = "TEST"
	DropzoneGetResponseDataModeSimulated DropzoneGetResponseDataMode = "SIMULATED"
	DropzoneGetResponseDataModeExercise  DropzoneGetResponseDataMode = "EXERCISE"
)

// Properties and characteristics of a Drop Zone, including name, location, shape,
// type code, survey date, and remarks.
type DropzoneListResponse struct {
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
	DataMode DropzoneListResponseDataMode `json:"dataMode,required"`
	// WGS84 latitude of the drop zone, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the drop zone, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// The name of the drop zone.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Specifies an alternate country code for the drop zone if the data provider code
	// is not part of an official Country Code standard such as ISO-3166 or FIPS. This
	// field will be set to the value provided by the source and should be used for all
	// Queries specifying a Country Code.
	AltCountryCode string `json:"altCountryCode"`
	// Specifies the country name associated with the source provided alternate country
	// code.
	AltCountryName string `json:"altCountryName"`
	// The date the drop zone survey was approved, in ISO 8601 UTC format with
	// millisecond precision.
	ApprovalDate time.Time `json:"approvalDate" format:"date-time"`
	// The type code for the drop zone.
	Code string `json:"code"`
	// The Country Code where the drop zone is located. This value is typically the ISO
	// 3166 Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode string `json:"countryCode"`
	// The country name of the location for the drop zone.
	CountryName string `json:"countryName"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The survey expiration date of the drop zone, in ISO 8601 UTC format with
	// millisecond precision.
	ExpirationDate time.Time `json:"expirationDate" format:"date-time"`
	// The external identifier assigned to the drop zone.
	ExtIdentifier string `json:"extIdentifier"`
	// The ID of the site associated with the drop zone.
	IDSite string `json:"idSite"`
	// Last time the drop zone information was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdate time.Time `json:"lastUpdate" format:"date-time"`
	// The length dimension of the drop zone in meters for non-circular drop zones.
	Length float64 `json:"length"`
	// The Major Command (MAJCOM) responsible for management of the drop zone.
	Majcom string `json:"majcom"`
	// The nearest reference location to the drop zone.
	NearestLoc string `json:"nearestLoc"`
	// The approval date for the drop zone by an air drop authority certifying
	// operational usage, in ISO 8601 UTC format with millisecond precision.
	OperationalApprovalDate time.Time `json:"operationalApprovalDate" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The name assigned to the drop zone point.
	PointName string `json:"pointName"`
	// The radius dimension of the drop zone in meters for circular drop zones.
	Radius float64 `json:"radius"`
	// The date the drop zone was recertified, in ISO 8601 UTC format with millisecond
	// precision.
	RecertDate time.Time `json:"recertDate" format:"date-time"`
	// Remarks concerning the drop zone.
	Remark string `json:"remark"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The US alphabetical code for the state where the drop zone is located.
	StateAbbr string `json:"stateAbbr"`
	// The name of the state where the drop zone is located.
	StateName string `json:"stateName"`
	// The date the drop zone survey was performed, in ISO 8601 UTC format with
	// millisecond precision.
	SurveyDate time.Time `json:"surveyDate" format:"date-time"`
	// The width dimension of the drop zone in meters for non-circular drop zones.
	Width float64 `json:"width"`
	// The identifier of the Zone Availability Report (ZAR) for the drop zone.
	ZarID string `json:"zarId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking   respjson.Field
		DataMode                respjson.Field
		Lat                     respjson.Field
		Lon                     respjson.Field
		Name                    respjson.Field
		Source                  respjson.Field
		ID                      respjson.Field
		AltCountryCode          respjson.Field
		AltCountryName          respjson.Field
		ApprovalDate            respjson.Field
		Code                    respjson.Field
		CountryCode             respjson.Field
		CountryName             respjson.Field
		CreatedAt               respjson.Field
		CreatedBy               respjson.Field
		ExpirationDate          respjson.Field
		ExtIdentifier           respjson.Field
		IDSite                  respjson.Field
		LastUpdate              respjson.Field
		Length                  respjson.Field
		Majcom                  respjson.Field
		NearestLoc              respjson.Field
		OperationalApprovalDate respjson.Field
		Origin                  respjson.Field
		OrigNetwork             respjson.Field
		PointName               respjson.Field
		Radius                  respjson.Field
		RecertDate              respjson.Field
		Remark                  respjson.Field
		SourceDl                respjson.Field
		StateAbbr               respjson.Field
		StateName               respjson.Field
		SurveyDate              respjson.Field
		Width                   respjson.Field
		ZarID                   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DropzoneListResponse) RawJSON() string { return r.JSON.raw }
func (r *DropzoneListResponse) UnmarshalJSON(data []byte) error {
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
type DropzoneListResponseDataMode string

const (
	DropzoneListResponseDataModeReal      DropzoneListResponseDataMode = "REAL"
	DropzoneListResponseDataModeTest      DropzoneListResponseDataMode = "TEST"
	DropzoneListResponseDataModeSimulated DropzoneListResponseDataMode = "SIMULATED"
	DropzoneListResponseDataModeExercise  DropzoneListResponseDataMode = "EXERCISE"
)

type DropzoneQueryHelpResponse struct {
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
func (r DropzoneQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *DropzoneQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Properties and characteristics of a Drop Zone, including name, location, shape,
// type code, survey date, and remarks.
type DropzoneTupleResponse struct {
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
	DataMode DropzoneTupleResponseDataMode `json:"dataMode,required"`
	// WGS84 latitude of the drop zone, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the drop zone, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// The name of the drop zone.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Specifies an alternate country code for the drop zone if the data provider code
	// is not part of an official Country Code standard such as ISO-3166 or FIPS. This
	// field will be set to the value provided by the source and should be used for all
	// Queries specifying a Country Code.
	AltCountryCode string `json:"altCountryCode"`
	// Specifies the country name associated with the source provided alternate country
	// code.
	AltCountryName string `json:"altCountryName"`
	// The date the drop zone survey was approved, in ISO 8601 UTC format with
	// millisecond precision.
	ApprovalDate time.Time `json:"approvalDate" format:"date-time"`
	// The type code for the drop zone.
	Code string `json:"code"`
	// The Country Code where the drop zone is located. This value is typically the ISO
	// 3166 Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode string `json:"countryCode"`
	// The country name of the location for the drop zone.
	CountryName string `json:"countryName"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The survey expiration date of the drop zone, in ISO 8601 UTC format with
	// millisecond precision.
	ExpirationDate time.Time `json:"expirationDate" format:"date-time"`
	// The external identifier assigned to the drop zone.
	ExtIdentifier string `json:"extIdentifier"`
	// The ID of the site associated with the drop zone.
	IDSite string `json:"idSite"`
	// Last time the drop zone information was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdate time.Time `json:"lastUpdate" format:"date-time"`
	// The length dimension of the drop zone in meters for non-circular drop zones.
	Length float64 `json:"length"`
	// The Major Command (MAJCOM) responsible for management of the drop zone.
	Majcom string `json:"majcom"`
	// The nearest reference location to the drop zone.
	NearestLoc string `json:"nearestLoc"`
	// The approval date for the drop zone by an air drop authority certifying
	// operational usage, in ISO 8601 UTC format with millisecond precision.
	OperationalApprovalDate time.Time `json:"operationalApprovalDate" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The name assigned to the drop zone point.
	PointName string `json:"pointName"`
	// The radius dimension of the drop zone in meters for circular drop zones.
	Radius float64 `json:"radius"`
	// The date the drop zone was recertified, in ISO 8601 UTC format with millisecond
	// precision.
	RecertDate time.Time `json:"recertDate" format:"date-time"`
	// Remarks concerning the drop zone.
	Remark string `json:"remark"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The US alphabetical code for the state where the drop zone is located.
	StateAbbr string `json:"stateAbbr"`
	// The name of the state where the drop zone is located.
	StateName string `json:"stateName"`
	// The date the drop zone survey was performed, in ISO 8601 UTC format with
	// millisecond precision.
	SurveyDate time.Time `json:"surveyDate" format:"date-time"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The width dimension of the drop zone in meters for non-circular drop zones.
	Width float64 `json:"width"`
	// The identifier of the Zone Availability Report (ZAR) for the drop zone.
	ZarID string `json:"zarId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking   respjson.Field
		DataMode                respjson.Field
		Lat                     respjson.Field
		Lon                     respjson.Field
		Name                    respjson.Field
		Source                  respjson.Field
		ID                      respjson.Field
		AltCountryCode          respjson.Field
		AltCountryName          respjson.Field
		ApprovalDate            respjson.Field
		Code                    respjson.Field
		CountryCode             respjson.Field
		CountryName             respjson.Field
		CreatedAt               respjson.Field
		CreatedBy               respjson.Field
		ExpirationDate          respjson.Field
		ExtIdentifier           respjson.Field
		IDSite                  respjson.Field
		LastUpdate              respjson.Field
		Length                  respjson.Field
		Majcom                  respjson.Field
		NearestLoc              respjson.Field
		OperationalApprovalDate respjson.Field
		Origin                  respjson.Field
		OrigNetwork             respjson.Field
		PointName               respjson.Field
		Radius                  respjson.Field
		RecertDate              respjson.Field
		Remark                  respjson.Field
		SourceDl                respjson.Field
		StateAbbr               respjson.Field
		StateName               respjson.Field
		SurveyDate              respjson.Field
		UpdatedAt               respjson.Field
		UpdatedBy               respjson.Field
		Width                   respjson.Field
		ZarID                   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DropzoneTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *DropzoneTupleResponse) UnmarshalJSON(data []byte) error {
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
type DropzoneTupleResponseDataMode string

const (
	DropzoneTupleResponseDataModeReal      DropzoneTupleResponseDataMode = "REAL"
	DropzoneTupleResponseDataModeTest      DropzoneTupleResponseDataMode = "TEST"
	DropzoneTupleResponseDataModeSimulated DropzoneTupleResponseDataMode = "SIMULATED"
	DropzoneTupleResponseDataModeExercise  DropzoneTupleResponseDataMode = "EXERCISE"
)

type DropzoneNewParams struct {
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
	DataMode DropzoneNewParamsDataMode `json:"dataMode,omitzero,required"`
	// WGS84 latitude of the drop zone, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the drop zone, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// The name of the drop zone.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Specifies an alternate country code for the drop zone if the data provider code
	// is not part of an official Country Code standard such as ISO-3166 or FIPS. This
	// field will be set to the value provided by the source and should be used for all
	// Queries specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Specifies the country name associated with the source provided alternate country
	// code.
	AltCountryName param.Opt[string] `json:"altCountryName,omitzero"`
	// The date the drop zone survey was approved, in ISO 8601 UTC format with
	// millisecond precision.
	ApprovalDate param.Opt[time.Time] `json:"approvalDate,omitzero" format:"date-time"`
	// The type code for the drop zone.
	Code param.Opt[string] `json:"code,omitzero"`
	// The Country Code where the drop zone is located. This value is typically the ISO
	// 3166 Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// The country name of the location for the drop zone.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// The survey expiration date of the drop zone, in ISO 8601 UTC format with
	// millisecond precision.
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	// The external identifier assigned to the drop zone.
	ExtIdentifier param.Opt[string] `json:"extIdentifier,omitzero"`
	// The ID of the site associated with the drop zone.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Last time the drop zone information was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdate param.Opt[time.Time] `json:"lastUpdate,omitzero" format:"date-time"`
	// The length dimension of the drop zone in meters for non-circular drop zones.
	Length param.Opt[float64] `json:"length,omitzero"`
	// The Major Command (MAJCOM) responsible for management of the drop zone.
	Majcom param.Opt[string] `json:"majcom,omitzero"`
	// The nearest reference location to the drop zone.
	NearestLoc param.Opt[string] `json:"nearestLoc,omitzero"`
	// The approval date for the drop zone by an air drop authority certifying
	// operational usage, in ISO 8601 UTC format with millisecond precision.
	OperationalApprovalDate param.Opt[time.Time] `json:"operationalApprovalDate,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The name assigned to the drop zone point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	// The radius dimension of the drop zone in meters for circular drop zones.
	Radius param.Opt[float64] `json:"radius,omitzero"`
	// The date the drop zone was recertified, in ISO 8601 UTC format with millisecond
	// precision.
	RecertDate param.Opt[time.Time] `json:"recertDate,omitzero" format:"date-time"`
	// Remarks concerning the drop zone.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The US alphabetical code for the state where the drop zone is located.
	StateAbbr param.Opt[string] `json:"stateAbbr,omitzero"`
	// The name of the state where the drop zone is located.
	StateName param.Opt[string] `json:"stateName,omitzero"`
	// The date the drop zone survey was performed, in ISO 8601 UTC format with
	// millisecond precision.
	SurveyDate param.Opt[time.Time] `json:"surveyDate,omitzero" format:"date-time"`
	// The width dimension of the drop zone in meters for non-circular drop zones.
	Width param.Opt[float64] `json:"width,omitzero"`
	// The identifier of the Zone Availability Report (ZAR) for the drop zone.
	ZarID param.Opt[string] `json:"zarId,omitzero"`
	paramObj
}

func (r DropzoneNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DropzoneNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DropzoneNewParams) UnmarshalJSON(data []byte) error {
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
type DropzoneNewParamsDataMode string

const (
	DropzoneNewParamsDataModeReal      DropzoneNewParamsDataMode = "REAL"
	DropzoneNewParamsDataModeTest      DropzoneNewParamsDataMode = "TEST"
	DropzoneNewParamsDataModeSimulated DropzoneNewParamsDataMode = "SIMULATED"
	DropzoneNewParamsDataModeExercise  DropzoneNewParamsDataMode = "EXERCISE"
)

type DropzoneGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DropzoneGetParams]'s query parameters as `url.Values`.
func (r DropzoneGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DropzoneUpdateParams struct {
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
	DataMode DropzoneUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// WGS84 latitude of the drop zone, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the drop zone, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// The name of the drop zone.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Specifies an alternate country code for the drop zone if the data provider code
	// is not part of an official Country Code standard such as ISO-3166 or FIPS. This
	// field will be set to the value provided by the source and should be used for all
	// Queries specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Specifies the country name associated with the source provided alternate country
	// code.
	AltCountryName param.Opt[string] `json:"altCountryName,omitzero"`
	// The date the drop zone survey was approved, in ISO 8601 UTC format with
	// millisecond precision.
	ApprovalDate param.Opt[time.Time] `json:"approvalDate,omitzero" format:"date-time"`
	// The type code for the drop zone.
	Code param.Opt[string] `json:"code,omitzero"`
	// The Country Code where the drop zone is located. This value is typically the ISO
	// 3166 Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// The country name of the location for the drop zone.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// The survey expiration date of the drop zone, in ISO 8601 UTC format with
	// millisecond precision.
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	// The external identifier assigned to the drop zone.
	ExtIdentifier param.Opt[string] `json:"extIdentifier,omitzero"`
	// The ID of the site associated with the drop zone.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Last time the drop zone information was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdate param.Opt[time.Time] `json:"lastUpdate,omitzero" format:"date-time"`
	// The length dimension of the drop zone in meters for non-circular drop zones.
	Length param.Opt[float64] `json:"length,omitzero"`
	// The Major Command (MAJCOM) responsible for management of the drop zone.
	Majcom param.Opt[string] `json:"majcom,omitzero"`
	// The nearest reference location to the drop zone.
	NearestLoc param.Opt[string] `json:"nearestLoc,omitzero"`
	// The approval date for the drop zone by an air drop authority certifying
	// operational usage, in ISO 8601 UTC format with millisecond precision.
	OperationalApprovalDate param.Opt[time.Time] `json:"operationalApprovalDate,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The name assigned to the drop zone point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	// The radius dimension of the drop zone in meters for circular drop zones.
	Radius param.Opt[float64] `json:"radius,omitzero"`
	// The date the drop zone was recertified, in ISO 8601 UTC format with millisecond
	// precision.
	RecertDate param.Opt[time.Time] `json:"recertDate,omitzero" format:"date-time"`
	// Remarks concerning the drop zone.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The US alphabetical code for the state where the drop zone is located.
	StateAbbr param.Opt[string] `json:"stateAbbr,omitzero"`
	// The name of the state where the drop zone is located.
	StateName param.Opt[string] `json:"stateName,omitzero"`
	// The date the drop zone survey was performed, in ISO 8601 UTC format with
	// millisecond precision.
	SurveyDate param.Opt[time.Time] `json:"surveyDate,omitzero" format:"date-time"`
	// The width dimension of the drop zone in meters for non-circular drop zones.
	Width param.Opt[float64] `json:"width,omitzero"`
	// The identifier of the Zone Availability Report (ZAR) for the drop zone.
	ZarID param.Opt[string] `json:"zarId,omitzero"`
	paramObj
}

func (r DropzoneUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DropzoneUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DropzoneUpdateParams) UnmarshalJSON(data []byte) error {
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
type DropzoneUpdateParamsDataMode string

const (
	DropzoneUpdateParamsDataModeReal      DropzoneUpdateParamsDataMode = "REAL"
	DropzoneUpdateParamsDataModeTest      DropzoneUpdateParamsDataMode = "TEST"
	DropzoneUpdateParamsDataModeSimulated DropzoneUpdateParamsDataMode = "SIMULATED"
	DropzoneUpdateParamsDataModeExercise  DropzoneUpdateParamsDataMode = "EXERCISE"
)

type DropzoneListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DropzoneListParams]'s query parameters as `url.Values`.
func (r DropzoneListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DropzoneCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DropzoneCountParams]'s query parameters as `url.Values`.
func (r DropzoneCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DropzoneNewBulkParams struct {
	Body []DropzoneNewBulkParamsBody
	paramObj
}

func (r DropzoneNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *DropzoneNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Properties and characteristics of a Drop Zone, including name, location, shape,
// type code, survey date, and remarks.
//
// The properties ClassificationMarking, DataMode, Lat, Lon, Name, Source are
// required.
type DropzoneNewBulkParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// WGS84 latitude of the drop zone, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the drop zone, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// The name of the drop zone.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Specifies an alternate country code for the drop zone if the data provider code
	// is not part of an official Country Code standard such as ISO-3166 or FIPS. This
	// field will be set to the value provided by the source and should be used for all
	// Queries specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Specifies the country name associated with the source provided alternate country
	// code.
	AltCountryName param.Opt[string] `json:"altCountryName,omitzero"`
	// The date the drop zone survey was approved, in ISO 8601 UTC format with
	// millisecond precision.
	ApprovalDate param.Opt[time.Time] `json:"approvalDate,omitzero" format:"date-time"`
	// The type code for the drop zone.
	Code param.Opt[string] `json:"code,omitzero"`
	// The Country Code where the drop zone is located. This value is typically the ISO
	// 3166 Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// The country name of the location for the drop zone.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// The survey expiration date of the drop zone, in ISO 8601 UTC format with
	// millisecond precision.
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	// The external identifier assigned to the drop zone.
	ExtIdentifier param.Opt[string] `json:"extIdentifier,omitzero"`
	// The ID of the site associated with the drop zone.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Last time the drop zone information was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdate param.Opt[time.Time] `json:"lastUpdate,omitzero" format:"date-time"`
	// The length dimension of the drop zone in meters for non-circular drop zones.
	Length param.Opt[float64] `json:"length,omitzero"`
	// The Major Command (MAJCOM) responsible for management of the drop zone.
	Majcom param.Opt[string] `json:"majcom,omitzero"`
	// The nearest reference location to the drop zone.
	NearestLoc param.Opt[string] `json:"nearestLoc,omitzero"`
	// The approval date for the drop zone by an air drop authority certifying
	// operational usage, in ISO 8601 UTC format with millisecond precision.
	OperationalApprovalDate param.Opt[time.Time] `json:"operationalApprovalDate,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The name assigned to the drop zone point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	// The radius dimension of the drop zone in meters for circular drop zones.
	Radius param.Opt[float64] `json:"radius,omitzero"`
	// The date the drop zone was recertified, in ISO 8601 UTC format with millisecond
	// precision.
	RecertDate param.Opt[time.Time] `json:"recertDate,omitzero" format:"date-time"`
	// Remarks concerning the drop zone.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The US alphabetical code for the state where the drop zone is located.
	StateAbbr param.Opt[string] `json:"stateAbbr,omitzero"`
	// The name of the state where the drop zone is located.
	StateName param.Opt[string] `json:"stateName,omitzero"`
	// The date the drop zone survey was performed, in ISO 8601 UTC format with
	// millisecond precision.
	SurveyDate param.Opt[time.Time] `json:"surveyDate,omitzero" format:"date-time"`
	// The width dimension of the drop zone in meters for non-circular drop zones.
	Width param.Opt[float64] `json:"width,omitzero"`
	// The identifier of the Zone Availability Report (ZAR) for the drop zone.
	ZarID param.Opt[string] `json:"zarId,omitzero"`
	paramObj
}

func (r DropzoneNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow DropzoneNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DropzoneNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DropzoneNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type DropzoneTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DropzoneTupleParams]'s query parameters as `url.Values`.
func (r DropzoneTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DropzoneUnvalidatedPublishParams struct {
	Body []DropzoneUnvalidatedPublishParamsBody
	paramObj
}

func (r DropzoneUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *DropzoneUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Properties and characteristics of a Drop Zone, including name, location, shape,
// type code, survey date, and remarks.
//
// The properties ClassificationMarking, DataMode, Lat, Lon, Name, Source are
// required.
type DropzoneUnvalidatedPublishParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// WGS84 latitude of the drop zone, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the drop zone, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// The name of the drop zone.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Specifies an alternate country code for the drop zone if the data provider code
	// is not part of an official Country Code standard such as ISO-3166 or FIPS. This
	// field will be set to the value provided by the source and should be used for all
	// Queries specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Specifies the country name associated with the source provided alternate country
	// code.
	AltCountryName param.Opt[string] `json:"altCountryName,omitzero"`
	// The date the drop zone survey was approved, in ISO 8601 UTC format with
	// millisecond precision.
	ApprovalDate param.Opt[time.Time] `json:"approvalDate,omitzero" format:"date-time"`
	// The type code for the drop zone.
	Code param.Opt[string] `json:"code,omitzero"`
	// The Country Code where the drop zone is located. This value is typically the ISO
	// 3166 Alpha-2 two-character country code, however it can also represent various
	// consortiums that do not appear in the ISO document. The code must correspond to
	// an existing country in the UDL’s country API. Call udl/country/{code} to get any
	// associated FIPS code, ISO Alpha-3 code, or alternate code values that exist for
	// the specified country code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// The country name of the location for the drop zone.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// The survey expiration date of the drop zone, in ISO 8601 UTC format with
	// millisecond precision.
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	// The external identifier assigned to the drop zone.
	ExtIdentifier param.Opt[string] `json:"extIdentifier,omitzero"`
	// The ID of the site associated with the drop zone.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Last time the drop zone information was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdate param.Opt[time.Time] `json:"lastUpdate,omitzero" format:"date-time"`
	// The length dimension of the drop zone in meters for non-circular drop zones.
	Length param.Opt[float64] `json:"length,omitzero"`
	// The Major Command (MAJCOM) responsible for management of the drop zone.
	Majcom param.Opt[string] `json:"majcom,omitzero"`
	// The nearest reference location to the drop zone.
	NearestLoc param.Opt[string] `json:"nearestLoc,omitzero"`
	// The approval date for the drop zone by an air drop authority certifying
	// operational usage, in ISO 8601 UTC format with millisecond precision.
	OperationalApprovalDate param.Opt[time.Time] `json:"operationalApprovalDate,omitzero" format:"date-time"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The name assigned to the drop zone point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	// The radius dimension of the drop zone in meters for circular drop zones.
	Radius param.Opt[float64] `json:"radius,omitzero"`
	// The date the drop zone was recertified, in ISO 8601 UTC format with millisecond
	// precision.
	RecertDate param.Opt[time.Time] `json:"recertDate,omitzero" format:"date-time"`
	// Remarks concerning the drop zone.
	Remark param.Opt[string] `json:"remark,omitzero"`
	// The US alphabetical code for the state where the drop zone is located.
	StateAbbr param.Opt[string] `json:"stateAbbr,omitzero"`
	// The name of the state where the drop zone is located.
	StateName param.Opt[string] `json:"stateName,omitzero"`
	// The date the drop zone survey was performed, in ISO 8601 UTC format with
	// millisecond precision.
	SurveyDate param.Opt[time.Time] `json:"surveyDate,omitzero" format:"date-time"`
	// The width dimension of the drop zone in meters for non-circular drop zones.
	Width param.Opt[float64] `json:"width,omitzero"`
	// The identifier of the Zone Availability Report (ZAR) for the drop zone.
	ZarID param.Opt[string] `json:"zarId,omitzero"`
	paramObj
}

func (r DropzoneUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow DropzoneUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DropzoneUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DropzoneUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
