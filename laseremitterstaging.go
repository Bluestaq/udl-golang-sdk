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

// LaseremitterStagingService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaseremitterStagingService] method instead.
type LaseremitterStagingService struct {
	Options []option.RequestOption
}

// NewLaseremitterStagingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLaseremitterStagingService(opts ...option.RequestOption) (r LaseremitterStagingService) {
	r = LaseremitterStagingService{}
	r.Options = opts
	return
}

// Service operation to take a single LaserEmitterStaging record as a POST body and
// ingest into the staging database. This API allows users to create, manage, and
// review LaserEmitterStaging records in a staging environment before their
// incorporation into the production UDL. It supports workflows involving
// validation, review, and approval of emitter data to ensure consistency,
// compliance, and data quality. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *LaseremitterStagingService) New(ctx context.Context, body LaseremitterStagingNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/laseremitterstaging"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single LaserEmitterStaging record by its unique ID
// passed as a path parameter.
func (r *LaseremitterStagingService) Get(ctx context.Context, id string, query LaseremitterStagingGetParams, opts ...option.RequestOption) (res *LaseremitterStagingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/laseremitterstaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single LaserEmitterStaging record. A specific role
// is required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *LaseremitterStagingService) Update(ctx context.Context, id string, body LaseremitterStagingUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/laseremitterstaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaseremitterStagingService) List(ctx context.Context, query LaseremitterStagingListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaseremitterStagingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/laseremitterstaging"
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
func (r *LaseremitterStagingService) ListAutoPaging(ctx context.Context, query LaseremitterStagingListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaseremitterStagingListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a LaserEmitterStaging record specified by the passed
// ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *LaseremitterStagingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/laseremitterstaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to take multiple LaserEmitterStaging records as a POST body
// and ingest into the staging database. This API allows users to create, manage,
// and review LaserEmitterStaging records in a staging environment before their
// incorporation into the production UDL. It supports workflows involving
// validation, review, and approval of emitter data to ensure consistency,
// compliance, and data quality. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *LaseremitterStagingService) NewBulk(ctx context.Context, body LaseremitterStagingNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/laseremitterstaging/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *LaseremitterStagingService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *LaseremitterStagingQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/laseremitterstaging/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Model representation of a nominal laser emitter. This entity contains minimal
// information used to stage laser emitter entities.
type LaseremitterStagingGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this Laser Emitter.
	LaserName string `json:"laserName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Emitter altitude relative to WGS-84 ellipsoid, in kilometers. Positive values
	// indicate an emitter height above ellipsoid, and negative values indicate an
	// emitter height below ellipsoid.
	Altitude float64 `json:"altitude"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType string `json:"laserType"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the Laser Emitter equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry string `json:"locationCountry"`
	// WGS-84 longitude of the emitter, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the Laser Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry string `json:"ownerCountry"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		LaserName             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Altitude              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		LaserType             respjson.Field
		Lat                   respjson.Field
		LocationCountry       respjson.Field
		Lon                   respjson.Field
		OwnerCountry          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaseremitterStagingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaseremitterStagingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a nominal laser emitter. This entity contains minimal
// information used to stage laser emitter entities.
type LaseremitterStagingListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this Laser Emitter.
	LaserName string `json:"laserName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Emitter altitude relative to WGS-84 ellipsoid, in kilometers. Positive values
	// indicate an emitter height above ellipsoid, and negative values indicate an
	// emitter height below ellipsoid.
	Altitude float64 `json:"altitude"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType string `json:"laserType"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the Laser Emitter equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry string `json:"locationCountry"`
	// WGS-84 longitude of the emitter, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the Laser Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry string `json:"ownerCountry"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		LaserName             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Altitude              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		LaserType             respjson.Field
		Lat                   respjson.Field
		LocationCountry       respjson.Field
		Lon                   respjson.Field
		OwnerCountry          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaseremitterStagingListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaseremitterStagingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaseremitterStagingQueryhelpResponse struct {
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
func (r LaseremitterStagingQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *LaseremitterStagingQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaseremitterStagingNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this Laser Emitter.
	LaserName string `json:"laserName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Emitter altitude relative to WGS-84 ellipsoid, in kilometers. Positive values
	// indicate an emitter height above ellipsoid, and negative values indicate an
	// emitter height below ellipsoid.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType param.Opt[string] `json:"laserType,omitzero"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the Laser Emitter equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry param.Opt[string] `json:"locationCountry,omitzero"`
	// WGS-84 longitude of the emitter, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the Laser Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	paramObj
}

func (r LaseremitterStagingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LaseremitterStagingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaseremitterStagingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaseremitterStagingGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaseremitterStagingGetParams]'s query parameters as
// `url.Values`.
func (r LaseremitterStagingGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaseremitterStagingUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this Laser Emitter.
	LaserName string `json:"laserName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Emitter altitude relative to WGS-84 ellipsoid, in kilometers. Positive values
	// indicate an emitter height above ellipsoid, and negative values indicate an
	// emitter height below ellipsoid.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType param.Opt[string] `json:"laserType,omitzero"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the Laser Emitter equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry param.Opt[string] `json:"locationCountry,omitzero"`
	// WGS-84 longitude of the emitter, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the Laser Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	paramObj
}

func (r LaseremitterStagingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LaseremitterStagingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaseremitterStagingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaseremitterStagingListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaseremitterStagingListParams]'s query parameters as
// `url.Values`.
func (r LaseremitterStagingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaseremitterStagingNewBulkParams struct {
	Body []LaseremitterStagingNewBulkParamsBody
	paramObj
}

func (r LaseremitterStagingNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LaseremitterStagingNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Model representation of a nominal laser emitter. This entity contains minimal
// information used to stage laser emitter entities.
//
// The properties ClassificationMarking, LaserName, Source are required.
type LaseremitterStagingNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this Laser Emitter.
	LaserName string `json:"laserName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Emitter altitude relative to WGS-84 ellipsoid, in kilometers. Positive values
	// indicate an emitter height above ellipsoid, and negative values indicate an
	// emitter height below ellipsoid.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// The type of laser (e.g. CONTINUOUS WAVE, PULSED, etc.).
	LaserType param.Opt[string] `json:"laserType,omitzero"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the Laser Emitter equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry param.Opt[string] `json:"locationCountry,omitzero"`
	// WGS-84 longitude of the emitter, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the Laser Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	paramObj
}

func (r LaseremitterStagingNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaseremitterStagingNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaseremitterStagingNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
