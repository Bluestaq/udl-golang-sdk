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

// SensorStatingService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSensorStatingService] method instead.
type SensorStatingService struct {
	Options []option.RequestOption
}

// NewSensorStatingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSensorStatingService(opts ...option.RequestOption) (r SensorStatingService) {
	r = SensorStatingService{}
	r.Options = opts
	return
}

// Service operation to take a single SensorStaging record as a POST body and
// ingest into the staging database. This API allows users to create, manage, and
// review SensorStaging records in a staging environment before their incorporation
// into the production UDL. It supports workflows involving validation, review, and
// approval of sensor data to ensure consistency, compliance, and data quality. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *SensorStatingService) New(ctx context.Context, body SensorStatingNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/sensorstaging"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single SensorStaging record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SensorStatingService) Update(ctx context.Context, id string, body SensorStatingUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensorstaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SensorStatingService) List(ctx context.Context, query SensorStatingListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SensorStatingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sensorstaging"
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
func (r *SensorStatingService) ListAutoPaging(ctx context.Context, query SensorStatingListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SensorStatingListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a SensorStaging record specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SensorStatingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensorstaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to take multiple SensorStaging records as a POST body and
// ingest into the staging database. This API allows users to create, manage, and
// review SensorStaging records in a staging environment before their incorporation
// into the production UDL. It supports workflows involving validation, review, and
// approval of sensor data to ensure consistency, compliance, and data quality. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *SensorStatingService) NewBulk(ctx context.Context, body SensorStatingNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/sensorstaging/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single SensorStaging record by its unique ID passed
// as a path parameter.
func (r *SensorStatingService) Get(ctx context.Context, id string, query SensorStatingGetParams, opts ...option.RequestOption) (res *SensorStatingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sensorstaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SensorStatingService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SensorStatingQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/sensorstaging/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Model representation of a nominal sensor. This entity contains minimal
// information used to stage sensor entities.
type SensorStatingListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Sensor altitude relative to WGS-84 ellipsoid, in meters. Positive values
	// indicate a sensor height above ellipsoid, and negative values indicate a sensor
	// height below ellipsoid.
	Altitude float64 `json:"altitude"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// WGS-84 latitude of the sensor, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the sensor equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry string `json:"locationCountry"`
	// WGS-84 longitude of the sensor, in degrees. -180 to 180 degrees (negative values
	// west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the sensor. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	OwnerCountry string `json:"ownerCountry"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber int64 `json:"sensorNumber"`
	// The observation measurement type produced by this sensor.
	SensorObservationType string `json:"sensorObservationType"`
	// The specific sensor type and/or surveillance capability of this sensor.
	SensorType string `json:"sensorType"`
	// Optional short or abbreviated name of this sensor.
	ShortName string `json:"shortName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		SensorName            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Altitude              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Lat                   respjson.Field
		LocationCountry       respjson.Field
		Lon                   respjson.Field
		OwnerCountry          respjson.Field
		SensorNumber          respjson.Field
		SensorObservationType respjson.Field
		SensorType            respjson.Field
		ShortName             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorStatingListResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorStatingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a nominal sensor. This entity contains minimal
// information used to stage sensor entities.
type SensorStatingGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Sensor altitude relative to WGS-84 ellipsoid, in meters. Positive values
	// indicate a sensor height above ellipsoid, and negative values indicate a sensor
	// height below ellipsoid.
	Altitude float64 `json:"altitude"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// WGS-84 latitude of the sensor, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the sensor equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry string `json:"locationCountry"`
	// WGS-84 longitude of the sensor, in degrees. -180 to 180 degrees (negative values
	// west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the sensor. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	OwnerCountry string `json:"ownerCountry"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber int64 `json:"sensorNumber"`
	// The observation measurement type produced by this sensor.
	SensorObservationType string `json:"sensorObservationType"`
	// The specific sensor type and/or surveillance capability of this sensor.
	SensorType string `json:"sensorType"`
	// Optional short or abbreviated name of this sensor.
	ShortName string `json:"shortName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		SensorName            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Altitude              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Lat                   respjson.Field
		LocationCountry       respjson.Field
		Lon                   respjson.Field
		OwnerCountry          respjson.Field
		SensorNumber          respjson.Field
		SensorObservationType respjson.Field
		SensorType            respjson.Field
		ShortName             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SensorStatingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorStatingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorStatingQueryhelpResponse struct {
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
func (r SensorStatingQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SensorStatingQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorStatingNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Sensor altitude relative to WGS-84 ellipsoid, in meters. Positive values
	// indicate a sensor height above ellipsoid, and negative values indicate a sensor
	// height below ellipsoid.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// WGS-84 latitude of the sensor, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the sensor equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry param.Opt[string] `json:"locationCountry,omitzero"`
	// WGS-84 longitude of the sensor, in degrees. -180 to 180 degrees (negative values
	// west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the sensor. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber param.Opt[int64] `json:"sensorNumber,omitzero"`
	// The observation measurement type produced by this sensor.
	SensorObservationType param.Opt[string] `json:"sensorObservationType,omitzero"`
	// The specific sensor type and/or surveillance capability of this sensor.
	SensorType param.Opt[string] `json:"sensorType,omitzero"`
	// Optional short or abbreviated name of this sensor.
	ShortName param.Opt[string] `json:"shortName,omitzero"`
	paramObj
}

func (r SensorStatingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SensorStatingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorStatingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorStatingUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Sensor altitude relative to WGS-84 ellipsoid, in meters. Positive values
	// indicate a sensor height above ellipsoid, and negative values indicate a sensor
	// height below ellipsoid.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// WGS-84 latitude of the sensor, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the sensor equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry param.Opt[string] `json:"locationCountry,omitzero"`
	// WGS-84 longitude of the sensor, in degrees. -180 to 180 degrees (negative values
	// west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the sensor. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber param.Opt[int64] `json:"sensorNumber,omitzero"`
	// The observation measurement type produced by this sensor.
	SensorObservationType param.Opt[string] `json:"sensorObservationType,omitzero"`
	// The specific sensor type and/or surveillance capability of this sensor.
	SensorType param.Opt[string] `json:"sensorType,omitzero"`
	// Optional short or abbreviated name of this sensor.
	ShortName param.Opt[string] `json:"shortName,omitzero"`
	paramObj
}

func (r SensorStatingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SensorStatingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorStatingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorStatingListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorStatingListParams]'s query parameters as
// `url.Values`.
func (r SensorStatingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SensorStatingNewBulkParams struct {
	Body []SensorStatingNewBulkParamsBody
	paramObj
}

func (r SensorStatingNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SensorStatingNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Model representation of a nominal sensor. This entity contains minimal
// information used to stage sensor entities.
//
// The properties ClassificationMarking, SensorName, Source are required.
type SensorStatingNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Unique name of this sensor.
	SensorName string `json:"sensorName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Sensor altitude relative to WGS-84 ellipsoid, in meters. Positive values
	// indicate a sensor height above ellipsoid, and negative values indicate a sensor
	// height below ellipsoid.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// WGS-84 latitude of the sensor, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the sensor equipment geographic coordinates reside. This value is
	// typically the ISO 3166 Alpha-2 two-character country code, however it can also
	// represent various consortiums that do not appear in the ISO document. The code
	// must correspond to an existing country in the UDL’s country API. Call
	// udl/country/{code} to get any associated FIPS code, ISO Alpha-3 code, or
	// alternate code values that exist for the specified country code.
	LocationCountry param.Opt[string] `json:"locationCountry,omitzero"`
	// WGS-84 longitude of the sensor, in degrees. -180 to 180 degrees (negative values
	// west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity
	// owning the sensor. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	OwnerCountry param.Opt[string] `json:"ownerCountry,omitzero"`
	// Number assigned to this sensor. Since there is no authoritative numbering
	// scheme, these numbers sometimes collide across sensors (especially commercial
	// sensors). It is therefore not a unique identifier.
	SensorNumber param.Opt[int64] `json:"sensorNumber,omitzero"`
	// The observation measurement type produced by this sensor.
	SensorObservationType param.Opt[string] `json:"sensorObservationType,omitzero"`
	// The specific sensor type and/or surveillance capability of this sensor.
	SensorType param.Opt[string] `json:"sensorType,omitzero"`
	// Optional short or abbreviated name of this sensor.
	ShortName param.Opt[string] `json:"shortName,omitzero"`
	paramObj
}

func (r SensorStatingNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SensorStatingNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SensorStatingNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SensorStatingGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SensorStatingGetParams]'s query parameters as `url.Values`.
func (r SensorStatingGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
