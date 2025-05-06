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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// PortService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortService] method instead.
type PortService struct {
	Options []option.RequestOption
}

// NewPortService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPortService(opts ...option.RequestOption) (r PortService) {
	r = PortService{}
	r.Options = opts
	return
}

// Service operation to take a single port record as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *PortService) New(ctx context.Context, body PortNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/port"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single port record. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *PortService) Update(ctx context.Context, id string, body PortUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/port/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *PortService) List(ctx context.Context, query PortListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[PortListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/port"
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
func (r *PortService) ListAutoPaging(ctx context.Context, query PortListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[PortListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *PortService) Count(ctx context.Context, query PortCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/port/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of port
// records as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *PortService) NewBulk(ctx context.Context, body PortNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/port/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single port record by its unique ID passed as a path
// parameter.
func (r *PortService) Get(ctx context.Context, id string, query PortGetParams, opts ...option.RequestOption) (res *PortGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/port/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *PortService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/port/queryhelp"
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
func (r *PortService) Tuple(ctx context.Context, query PortTupleParams, opts ...option.RequestOption) (res *[]PortTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/port/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Properties and characteristics of a maritime port, which includes location, port
// identifiers, and remarks.
type PortListResponse struct {
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
	DataMode PortListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Average time for a vessel at this port in hours.
	AvgDuration float64 `json:"avgDuration"`
	// The country where this port is located.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// The size of the harbor for this port measured in square kilometers.
	HarborSize float64 `json:"harborSize"`
	// The type of harbor for this port. The harbor type refers to how a port is
	// physically positioned.
	//
	// COASTAL BREAKWATER (CB)
	//
	// COASTAL NATURAL (CN)
	//
	// COASTAL TIDE GATE (CT)
	//
	// LAKE OR CANAL (LC)
	//
	// OPEN ROADSTEAD (OR)
	//
	// RIVER BASIN (RB)
	//
	// RIVER NATURAL (RN)
	//
	// RIVER TIDE GATE (RT)
	//
	// TYPHOON HARBOR (TH).
	HarborType string `json:"harborType"`
	// Unique identifier of the Site Entity associated with the Port record.
	IDSite string `json:"idSite"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// The five-character United Nations Code for Trade and Transport Locations
	// (UN/LOCODE) of this port. The first two letters of the code contains the ISO
	// 3166-1 alpha-2 country designation of the port country. The three remaining
	// characters identify a location within that country. Letters are preferred, but
	// if necessary digits 2 through 9 may be used, excluding "0" and "1" to avoid
	// confusion with the letters "O" and "I" respectively.
	Locode string `json:"locode"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// Maximum allowed vessel draught. Draught is the principal dimensions of any
	// waterborne vessel defined as the distance between the ship’s keel and the
	// waterline of the vessel measured in meters.
	MaxDraught float64 `json:"maxDraught"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Flag indicating whether a pilot is required at this port.
	PilotReqd bool `json:"pilotReqd"`
	// The name of this port.
	PortName string `json:"portName"`
	// The shelter afforded from wind, sea, and swell refers to the area where normal
	// port operations are conducted, usually the wharf area. Shelter afforded by the
	// anchorage area may be given for ports where cargo is handled by lighters. Values
	// given are EXCELLENT, FAIR, GOOD, POOR, or NONE.
	Shelter string `json:"shelter"`
	// The tide range of this port in meters.
	TideRange float64 `json:"tideRange"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AvgDuration           respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ExternalID            respjson.Field
		HarborSize            respjson.Field
		HarborType            respjson.Field
		IDSite                respjson.Field
		Lat                   respjson.Field
		Locode                respjson.Field
		Lon                   respjson.Field
		MaxDraught            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PilotReqd             respjson.Field
		PortName              respjson.Field
		Shelter               respjson.Field
		TideRange             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortListResponse) UnmarshalJSON(data []byte) error {
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
type PortListResponseDataMode string

const (
	PortListResponseDataModeReal      PortListResponseDataMode = "REAL"
	PortListResponseDataModeTest      PortListResponseDataMode = "TEST"
	PortListResponseDataModeSimulated PortListResponseDataMode = "SIMULATED"
	PortListResponseDataModeExercise  PortListResponseDataMode = "EXERCISE"
)

// Properties and characteristics of a maritime port, which includes location, port
// identifiers, and remarks.
type PortGetResponse struct {
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
	DataMode PortGetResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Average time for a vessel at this port in hours.
	AvgDuration float64 `json:"avgDuration"`
	// The country where this port is located.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// The size of the harbor for this port measured in square kilometers.
	HarborSize float64 `json:"harborSize"`
	// The type of harbor for this port. The harbor type refers to how a port is
	// physically positioned.
	//
	// COASTAL BREAKWATER (CB)
	//
	// COASTAL NATURAL (CN)
	//
	// COASTAL TIDE GATE (CT)
	//
	// LAKE OR CANAL (LC)
	//
	// OPEN ROADSTEAD (OR)
	//
	// RIVER BASIN (RB)
	//
	// RIVER NATURAL (RN)
	//
	// RIVER TIDE GATE (RT)
	//
	// TYPHOON HARBOR (TH).
	HarborType string `json:"harborType"`
	// Unique identifier of the Site Entity associated with the Port record.
	IDSite string `json:"idSite"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// The five-character United Nations Code for Trade and Transport Locations
	// (UN/LOCODE) of this port. The first two letters of the code contains the ISO
	// 3166-1 alpha-2 country designation of the port country. The three remaining
	// characters identify a location within that country. Letters are preferred, but
	// if necessary digits 2 through 9 may be used, excluding "0" and "1" to avoid
	// confusion with the letters "O" and "I" respectively.
	Locode string `json:"locode"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// Maximum allowed vessel draught. Draught is the principal dimensions of any
	// waterborne vessel defined as the distance between the ship’s keel and the
	// waterline of the vessel measured in meters.
	MaxDraught float64 `json:"maxDraught"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Flag indicating whether a pilot is required at this port.
	PilotReqd bool `json:"pilotReqd"`
	// The name of this port.
	PortName string `json:"portName"`
	// The shelter afforded from wind, sea, and swell refers to the area where normal
	// port operations are conducted, usually the wharf area. Shelter afforded by the
	// anchorage area may be given for ports where cargo is handled by lighters. Values
	// given are EXCELLENT, FAIR, GOOD, POOR, or NONE.
	Shelter string `json:"shelter"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The tide range of this port in meters.
	TideRange float64 `json:"tideRange"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AvgDuration           respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ExternalID            respjson.Field
		HarborSize            respjson.Field
		HarborType            respjson.Field
		IDSite                respjson.Field
		Lat                   respjson.Field
		Locode                respjson.Field
		Lon                   respjson.Field
		MaxDraught            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PilotReqd             respjson.Field
		PortName              respjson.Field
		Shelter               respjson.Field
		SourceDl              respjson.Field
		TideRange             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PortGetResponse) UnmarshalJSON(data []byte) error {
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
type PortGetResponseDataMode string

const (
	PortGetResponseDataModeReal      PortGetResponseDataMode = "REAL"
	PortGetResponseDataModeTest      PortGetResponseDataMode = "TEST"
	PortGetResponseDataModeSimulated PortGetResponseDataMode = "SIMULATED"
	PortGetResponseDataModeExercise  PortGetResponseDataMode = "EXERCISE"
)

// Properties and characteristics of a maritime port, which includes location, port
// identifiers, and remarks.
type PortTupleResponse struct {
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
	DataMode PortTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Average time for a vessel at this port in hours.
	AvgDuration float64 `json:"avgDuration"`
	// The country where this port is located.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// The size of the harbor for this port measured in square kilometers.
	HarborSize float64 `json:"harborSize"`
	// The type of harbor for this port. The harbor type refers to how a port is
	// physically positioned.
	//
	// COASTAL BREAKWATER (CB)
	//
	// COASTAL NATURAL (CN)
	//
	// COASTAL TIDE GATE (CT)
	//
	// LAKE OR CANAL (LC)
	//
	// OPEN ROADSTEAD (OR)
	//
	// RIVER BASIN (RB)
	//
	// RIVER NATURAL (RN)
	//
	// RIVER TIDE GATE (RT)
	//
	// TYPHOON HARBOR (TH).
	HarborType string `json:"harborType"`
	// Unique identifier of the Site Entity associated with the Port record.
	IDSite string `json:"idSite"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// The five-character United Nations Code for Trade and Transport Locations
	// (UN/LOCODE) of this port. The first two letters of the code contains the ISO
	// 3166-1 alpha-2 country designation of the port country. The three remaining
	// characters identify a location within that country. Letters are preferred, but
	// if necessary digits 2 through 9 may be used, excluding "0" and "1" to avoid
	// confusion with the letters "O" and "I" respectively.
	Locode string `json:"locode"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// Maximum allowed vessel draught. Draught is the principal dimensions of any
	// waterborne vessel defined as the distance between the ship’s keel and the
	// waterline of the vessel measured in meters.
	MaxDraught float64 `json:"maxDraught"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Flag indicating whether a pilot is required at this port.
	PilotReqd bool `json:"pilotReqd"`
	// The name of this port.
	PortName string `json:"portName"`
	// The shelter afforded from wind, sea, and swell refers to the area where normal
	// port operations are conducted, usually the wharf area. Shelter afforded by the
	// anchorage area may be given for ports where cargo is handled by lighters. Values
	// given are EXCELLENT, FAIR, GOOD, POOR, or NONE.
	Shelter string `json:"shelter"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The tide range of this port in meters.
	TideRange float64 `json:"tideRange"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AvgDuration           respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ExternalID            respjson.Field
		HarborSize            respjson.Field
		HarborType            respjson.Field
		IDSite                respjson.Field
		Lat                   respjson.Field
		Locode                respjson.Field
		Lon                   respjson.Field
		MaxDraught            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PilotReqd             respjson.Field
		PortName              respjson.Field
		Shelter               respjson.Field
		SourceDl              respjson.Field
		TideRange             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *PortTupleResponse) UnmarshalJSON(data []byte) error {
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
type PortTupleResponseDataMode string

const (
	PortTupleResponseDataModeReal      PortTupleResponseDataMode = "REAL"
	PortTupleResponseDataModeTest      PortTupleResponseDataMode = "TEST"
	PortTupleResponseDataModeSimulated PortTupleResponseDataMode = "SIMULATED"
	PortTupleResponseDataModeExercise  PortTupleResponseDataMode = "EXERCISE"
)

type PortNewParams struct {
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
	DataMode PortNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Average time for a vessel at this port in hours.
	AvgDuration param.Opt[float64] `json:"avgDuration,omitzero"`
	// The country where this port is located.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The size of the harbor for this port measured in square kilometers.
	HarborSize param.Opt[float64] `json:"harborSize,omitzero"`
	// The type of harbor for this port. The harbor type refers to how a port is
	// physically positioned.
	//
	// COASTAL BREAKWATER (CB)
	//
	// COASTAL NATURAL (CN)
	//
	// COASTAL TIDE GATE (CT)
	//
	// LAKE OR CANAL (LC)
	//
	// OPEN ROADSTEAD (OR)
	//
	// RIVER BASIN (RB)
	//
	// RIVER NATURAL (RN)
	//
	// RIVER TIDE GATE (RT)
	//
	// TYPHOON HARBOR (TH).
	HarborType param.Opt[string] `json:"harborType,omitzero"`
	// Unique identifier of the Site Entity associated with the Port record.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The five-character United Nations Code for Trade and Transport Locations
	// (UN/LOCODE) of this port. The first two letters of the code contains the ISO
	// 3166-1 alpha-2 country designation of the port country. The three remaining
	// characters identify a location within that country. Letters are preferred, but
	// if necessary digits 2 through 9 may be used, excluding "0" and "1" to avoid
	// confusion with the letters "O" and "I" respectively.
	Locode param.Opt[string] `json:"locode,omitzero"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Maximum allowed vessel draught. Draught is the principal dimensions of any
	// waterborne vessel defined as the distance between the ship’s keel and the
	// waterline of the vessel measured in meters.
	MaxDraught param.Opt[float64] `json:"maxDraught,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Flag indicating whether a pilot is required at this port.
	PilotReqd param.Opt[bool] `json:"pilotReqd,omitzero"`
	// The name of this port.
	PortName param.Opt[string] `json:"portName,omitzero"`
	// The shelter afforded from wind, sea, and swell refers to the area where normal
	// port operations are conducted, usually the wharf area. Shelter afforded by the
	// anchorage area may be given for ports where cargo is handled by lighters. Values
	// given are EXCELLENT, FAIR, GOOD, POOR, or NONE.
	Shelter param.Opt[string] `json:"shelter,omitzero"`
	// The tide range of this port in meters.
	TideRange param.Opt[float64] `json:"tideRange,omitzero"`
	paramObj
}

func (r PortNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortNewParams
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
type PortNewParamsDataMode string

const (
	PortNewParamsDataModeReal      PortNewParamsDataMode = "REAL"
	PortNewParamsDataModeTest      PortNewParamsDataMode = "TEST"
	PortNewParamsDataModeSimulated PortNewParamsDataMode = "SIMULATED"
	PortNewParamsDataModeExercise  PortNewParamsDataMode = "EXERCISE"
)

type PortUpdateParams struct {
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
	DataMode PortUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Average time for a vessel at this port in hours.
	AvgDuration param.Opt[float64] `json:"avgDuration,omitzero"`
	// The country where this port is located.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The size of the harbor for this port measured in square kilometers.
	HarborSize param.Opt[float64] `json:"harborSize,omitzero"`
	// The type of harbor for this port. The harbor type refers to how a port is
	// physically positioned.
	//
	// COASTAL BREAKWATER (CB)
	//
	// COASTAL NATURAL (CN)
	//
	// COASTAL TIDE GATE (CT)
	//
	// LAKE OR CANAL (LC)
	//
	// OPEN ROADSTEAD (OR)
	//
	// RIVER BASIN (RB)
	//
	// RIVER NATURAL (RN)
	//
	// RIVER TIDE GATE (RT)
	//
	// TYPHOON HARBOR (TH).
	HarborType param.Opt[string] `json:"harborType,omitzero"`
	// Unique identifier of the Site Entity associated with the Port record.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The five-character United Nations Code for Trade and Transport Locations
	// (UN/LOCODE) of this port. The first two letters of the code contains the ISO
	// 3166-1 alpha-2 country designation of the port country. The three remaining
	// characters identify a location within that country. Letters are preferred, but
	// if necessary digits 2 through 9 may be used, excluding "0" and "1" to avoid
	// confusion with the letters "O" and "I" respectively.
	Locode param.Opt[string] `json:"locode,omitzero"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Maximum allowed vessel draught. Draught is the principal dimensions of any
	// waterborne vessel defined as the distance between the ship’s keel and the
	// waterline of the vessel measured in meters.
	MaxDraught param.Opt[float64] `json:"maxDraught,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Flag indicating whether a pilot is required at this port.
	PilotReqd param.Opt[bool] `json:"pilotReqd,omitzero"`
	// The name of this port.
	PortName param.Opt[string] `json:"portName,omitzero"`
	// The shelter afforded from wind, sea, and swell refers to the area where normal
	// port operations are conducted, usually the wharf area. Shelter afforded by the
	// anchorage area may be given for ports where cargo is handled by lighters. Values
	// given are EXCELLENT, FAIR, GOOD, POOR, or NONE.
	Shelter param.Opt[string] `json:"shelter,omitzero"`
	// The tide range of this port in meters.
	TideRange param.Opt[float64] `json:"tideRange,omitzero"`
	paramObj
}

func (r PortUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PortUpdateParams
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
type PortUpdateParamsDataMode string

const (
	PortUpdateParamsDataModeReal      PortUpdateParamsDataMode = "REAL"
	PortUpdateParamsDataModeTest      PortUpdateParamsDataMode = "TEST"
	PortUpdateParamsDataModeSimulated PortUpdateParamsDataMode = "SIMULATED"
	PortUpdateParamsDataModeExercise  PortUpdateParamsDataMode = "EXERCISE"
)

type PortListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortListParams]'s query parameters as `url.Values`.
func (r PortListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortCountParams]'s query parameters as `url.Values`.
func (r PortCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortNewBulkParams struct {
	Body []PortNewBulkParamsBody
	paramObj
}

func (r PortNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Properties and characteristics of a maritime port, which includes location, port
// identifiers, and remarks.
//
// The properties ClassificationMarking, DataMode, Source are required.
type PortNewBulkParamsBody struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Average time for a vessel at this port in hours.
	AvgDuration param.Opt[float64] `json:"avgDuration,omitzero"`
	// The country where this port is located.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The size of the harbor for this port measured in square kilometers.
	HarborSize param.Opt[float64] `json:"harborSize,omitzero"`
	// The type of harbor for this port. The harbor type refers to how a port is
	// physically positioned.
	//
	// COASTAL BREAKWATER (CB)
	//
	// COASTAL NATURAL (CN)
	//
	// COASTAL TIDE GATE (CT)
	//
	// LAKE OR CANAL (LC)
	//
	// OPEN ROADSTEAD (OR)
	//
	// RIVER BASIN (RB)
	//
	// RIVER NATURAL (RN)
	//
	// RIVER TIDE GATE (RT)
	//
	// TYPHOON HARBOR (TH).
	HarborType param.Opt[string] `json:"harborType,omitzero"`
	// Unique identifier of the Site Entity associated with the Port record.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The five-character United Nations Code for Trade and Transport Locations
	// (UN/LOCODE) of this port. The first two letters of the code contains the ISO
	// 3166-1 alpha-2 country designation of the port country. The three remaining
	// characters identify a location within that country. Letters are preferred, but
	// if necessary digits 2 through 9 may be used, excluding "0" and "1" to avoid
	// confusion with the letters "O" and "I" respectively.
	Locode param.Opt[string] `json:"locode,omitzero"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Maximum allowed vessel draught. Draught is the principal dimensions of any
	// waterborne vessel defined as the distance between the ship’s keel and the
	// waterline of the vessel measured in meters.
	MaxDraught param.Opt[float64] `json:"maxDraught,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Flag indicating whether a pilot is required at this port.
	PilotReqd param.Opt[bool] `json:"pilotReqd,omitzero"`
	// The name of this port.
	PortName param.Opt[string] `json:"portName,omitzero"`
	// The shelter afforded from wind, sea, and swell refers to the area where normal
	// port operations are conducted, usually the wharf area. Shelter afforded by the
	// anchorage area may be given for ports where cargo is handled by lighters. Values
	// given are EXCELLENT, FAIR, GOOD, POOR, or NONE.
	Shelter param.Opt[string] `json:"shelter,omitzero"`
	// The tide range of this port in meters.
	TideRange param.Opt[float64] `json:"tideRange,omitzero"`
	paramObj
}

func (r PortNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow PortNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PortNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type PortGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortGetParams]'s query parameters as `url.Values`.
func (r PortGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortTupleParams]'s query parameters as `url.Values`.
func (r PortTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
