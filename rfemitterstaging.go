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

// RfEmitterStagingService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRfEmitterStagingService] method instead.
type RfEmitterStagingService struct {
	Options []option.RequestOption
}

// NewRfEmitterStagingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRfEmitterStagingService(opts ...option.RequestOption) (r RfEmitterStagingService) {
	r = RfEmitterStagingService{}
	r.Options = opts
	return
}

// Service operation to take a single RFEmitterStaging record as a POST body and
// ingest into the staging database. This API allows users to create, manage, and
// review RFEmitter records in a staging environment before their incorporation
// into the production UDL. It supports workflows involving validation, review, and
// approval of emitter data to ensure consistency, compliance, and data quality. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *RfEmitterStagingService) New(ctx context.Context, body RfEmitterStagingNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfemitterstaging"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single RFEmitterStaging record by its unique ID
// passed as a path parameter.
func (r *RfEmitterStagingService) Get(ctx context.Context, id string, query RfEmitterStagingGetParams, opts ...option.RequestOption) (res *RfEmitterStagingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfemitterstaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single RFEmitterStaging record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *RfEmitterStagingService) Update(ctx context.Context, id string, body RfEmitterStagingUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfemitterstaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *RfEmitterStagingService) List(ctx context.Context, query RfEmitterStagingListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[RfEmitterStagingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/rfemitterstaging"
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
func (r *RfEmitterStagingService) ListAutoPaging(ctx context.Context, query RfEmitterStagingListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[RfEmitterStagingListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a RFEmitterStaging record specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *RfEmitterStagingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfemitterstaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to take multiple RFEmitterStaging records as a POST body and
// ingest into the staging database. This API allows users to create, manage, and
// review RFEmitter records in a staging environment before their incorporation
// into the production UDL. It supports workflows involving validation, review, and
// approval of emitter data to ensure consistency, compliance, and data quality. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *RfEmitterStagingService) NewBulk(ctx context.Context, body RfEmitterStagingNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfemitterstaging/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *RfEmitterStagingService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *RfEmitterStagingQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/rfemitterstaging/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Model representation of a nominal RF emitter. This entity contains minimal
// information used to stage RF emitter entities.
type RfEmitterStagingGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this RF Emitter.
	Name string `json:"name,required"`
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
	// The originating system ID for the RF Emitter.
	ExtSysID string `json:"extSysId"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the RF Emitter equipment geographic coordinates reside. This value is
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
	// owning the RF Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry string `json:"ownerCountry"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype string `json:"subtype"`
	// Type of this RF Emitter.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Altitude              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ExtSysID              respjson.Field
		Lat                   respjson.Field
		LocationCountry       respjson.Field
		Lon                   respjson.Field
		OwnerCountry          respjson.Field
		Subtype               respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterStagingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterStagingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a nominal RF emitter. This entity contains minimal
// information used to stage RF emitter entities.
type RfEmitterStagingListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this RF Emitter.
	Name string `json:"name,required"`
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
	// The originating system ID for the RF Emitter.
	ExtSysID string `json:"extSysId"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the RF Emitter equipment geographic coordinates reside. This value is
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
	// owning the RF Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry string `json:"ownerCountry"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype string `json:"subtype"`
	// Type of this RF Emitter.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Altitude              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ExtSysID              respjson.Field
		Lat                   respjson.Field
		LocationCountry       respjson.Field
		Lon                   respjson.Field
		OwnerCountry          respjson.Field
		Subtype               respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfEmitterStagingListResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterStagingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterStagingQueryhelpResponse struct {
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
func (r RfEmitterStagingQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *RfEmitterStagingQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterStagingNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this RF Emitter.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Emitter altitude relative to WGS-84 ellipsoid, in kilometers. Positive values
	// indicate an emitter height above ellipsoid, and negative values indicate an
	// emitter height below ellipsoid.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// The originating system ID for the RF Emitter.
	ExtSysID param.Opt[string] `json:"extSysId,omitzero"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the RF Emitter equipment geographic coordinates reside. This value is
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
	// owning the RF Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype param.Opt[string] `json:"subtype,omitzero"`
	// Type of this RF Emitter.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r RfEmitterStagingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterStagingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterStagingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterStagingGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterStagingGetParams]'s query parameters as
// `url.Values`.
func (r RfEmitterStagingGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfEmitterStagingUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this RF Emitter.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Emitter altitude relative to WGS-84 ellipsoid, in kilometers. Positive values
	// indicate an emitter height above ellipsoid, and negative values indicate an
	// emitter height below ellipsoid.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// The originating system ID for the RF Emitter.
	ExtSysID param.Opt[string] `json:"extSysId,omitzero"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the RF Emitter equipment geographic coordinates reside. This value is
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
	// owning the RF Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype param.Opt[string] `json:"subtype,omitzero"`
	// Type of this RF Emitter.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r RfEmitterStagingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterStagingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterStagingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfEmitterStagingListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfEmitterStagingListParams]'s query parameters as
// `url.Values`.
func (r RfEmitterStagingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfEmitterStagingNewBulkParams struct {
	Body []RfEmitterStagingNewBulkParamsBody
	paramObj
}

func (r RfEmitterStagingNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *RfEmitterStagingNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Model representation of a nominal RF emitter. This entity contains minimal
// information used to stage RF emitter entities.
//
// The properties ClassificationMarking, Name, Source are required.
type RfEmitterStagingNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this RF Emitter.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Emitter altitude relative to WGS-84 ellipsoid, in kilometers. Positive values
	// indicate an emitter height above ellipsoid, and negative values indicate an
	// emitter height below ellipsoid.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The originating system ID for the RF Emitter.
	ExtSysID param.Opt[string] `json:"extSysId,omitzero"`
	// WGS-84 latitude of the emitter, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the RF Emitter equipment geographic coordinates reside. This value is
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
	// owning the RF Emitter. This value is typically the ISO 3166 Alpha-2
	// two-character country code, however it can also represent various consortiums
	// that do not appear in the ISO document. The code must correspond to an existing
	// country in the UDL’s country API. Call udl/country/{code} to get any associated
	// FIPS code, ISO Alpha-3 code, or alternate code values that exist for the
	// specified country code.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	// The RF Emitter subtype, which can distinguish specialized deployments (e.g.
	// BLOCK_0_AVL, BLOCK_0_DS1, BLOCK_0_TEST, BLOCK_1, BLOCK_1_TEST, NONE).
	Subtype param.Opt[string] `json:"subtype,omitzero"`
	// Type of this RF Emitter.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r RfEmitterStagingNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow RfEmitterStagingNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfEmitterStagingNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
