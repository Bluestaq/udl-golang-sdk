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

// BeamContourService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBeamContourService] method instead.
type BeamContourService struct {
	Options []option.RequestOption
}

// NewBeamContourService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBeamContourService(opts ...option.RequestOption) (r BeamContourService) {
	r = BeamContourService{}
	r.Options = opts
	return
}

// Service operation to take a single BeamContour as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *BeamContourService) New(ctx context.Context, body BeamContourNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/beamcontour"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single BeamContour by its unique ID passed as a path
// parameter.
func (r *BeamContourService) Get(ctx context.Context, id string, query BeamContourGetParams, opts ...option.RequestOption) (res *shared.BeamcontourFull, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/beamcontour/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single BeamContour. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *BeamContourService) Update(ctx context.Context, id string, body BeamContourUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/beamcontour/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *BeamContourService) List(ctx context.Context, query BeamContourListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[BeamcontourAbridged], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/beamcontour"
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
func (r *BeamContourService) ListAutoPaging(ctx context.Context, query BeamContourListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[BeamcontourAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a BeamContour object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *BeamContourService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/beamcontour/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *BeamContourService) Count(ctx context.Context, query BeamContourCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/beamcontour/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a list of BeamContours as a POST body and ingest into
// the database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *BeamContourService) NewBulk(ctx context.Context, body BeamContourNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/beamcontour/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *BeamContourService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *BeamContourQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/beamcontour/queryhelp"
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
func (r *BeamContourService) Tuple(ctx context.Context, query BeamContourTupleParams, opts ...option.RequestOption) (res *[]shared.BeamcontourFull, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/beamcontour/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Describes the beam contour associated with a beam entity. Beam contours are the
// geographic representation of the relative gain levels of beam power off of the
// maximum gain boresight points.
type BeamcontourAbridged struct {
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
	DataMode BeamcontourAbridgedDataMode `json:"dataMode,required"`
	// ID of the beam.
	IDBeam string `json:"idBeam,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of object represented in this record (BORESIGHT, CONTOUR, SVC AREA).
	// Boresight refers to the point of maximum/peak gain, and should not be confused
	// with the 'aim point' of the related beam. Gain contours are regions of coverage
	// referenced to the relative gain of the related beam. Service Areas are composed
	// of one or more service regions, with each region being either discrete point(s)
	// or a continuous contour.
	//
	// Any of "BORESIGHT", "CONTOUR", "SVC AREA".
	Type BeamcontourAbridgedType `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The index number of this contour. The value is required if type = CONTOUR.
	ContourIdx int64 `json:"contourIdx"`
	// Time the row was created in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The relative gain level in dB associated with this boresight or contour. Gain
	// does not apply to service area records. The value is required if type =
	// BORESIGHT or CONTOUR.
	Gain float64 `json:"gain"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	GeographyJson string `json:"geographyJson"`
	// Number of dimensions of the geometry depicted by region.
	GeographyNdims int64 `json:"geographyNdims"`
	// Geographical spatial_ref_sys for region.
	GeographySrid int64 `json:"geographySrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	GeographyText string `json:"geographyText"`
	// Type of region as projected.
	GeographyType string `json:"geographyType"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The region name within the service area.
	RegionName string `json:"regionName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDBeam                respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		ContourIdx            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Gain                  respjson.Field
		GeographyJson         respjson.Field
		GeographyNdims        respjson.Field
		GeographySrid         respjson.Field
		GeographyText         respjson.Field
		GeographyType         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RegionName            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BeamcontourAbridged) RawJSON() string { return r.JSON.raw }
func (r *BeamcontourAbridged) UnmarshalJSON(data []byte) error {
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
type BeamcontourAbridgedDataMode string

const (
	BeamcontourAbridgedDataModeReal      BeamcontourAbridgedDataMode = "REAL"
	BeamcontourAbridgedDataModeTest      BeamcontourAbridgedDataMode = "TEST"
	BeamcontourAbridgedDataModeSimulated BeamcontourAbridgedDataMode = "SIMULATED"
	BeamcontourAbridgedDataModeExercise  BeamcontourAbridgedDataMode = "EXERCISE"
)

// The type of object represented in this record (BORESIGHT, CONTOUR, SVC AREA).
// Boresight refers to the point of maximum/peak gain, and should not be confused
// with the 'aim point' of the related beam. Gain contours are regions of coverage
// referenced to the relative gain of the related beam. Service Areas are composed
// of one or more service regions, with each region being either discrete point(s)
// or a continuous contour.
type BeamcontourAbridgedType string

const (
	BeamcontourAbridgedTypeBoresight BeamcontourAbridgedType = "BORESIGHT"
	BeamcontourAbridgedTypeContour   BeamcontourAbridgedType = "CONTOUR"
	BeamcontourAbridgedTypeSvcArea   BeamcontourAbridgedType = "SVC AREA"
)

type BeamContourQueryHelpResponse struct {
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
func (r BeamContourQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *BeamContourQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BeamContourNewParams struct {
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
	DataMode BeamContourNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the beam.
	IDBeam string `json:"idBeam,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of object represented in this record (BORESIGHT, CONTOUR, SVC AREA).
	// Boresight refers to the point of maximum/peak gain, and should not be confused
	// with the 'aim point' of the related beam. Gain contours are regions of coverage
	// referenced to the relative gain of the related beam. Service Areas are composed
	// of one or more service regions, with each region being either discrete point(s)
	// or a continuous contour.
	//
	// Any of "BORESIGHT", "CONTOUR", "SVC AREA".
	Type BeamContourNewParamsType `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The index number of this contour. The value is required if type = CONTOUR.
	ContourIdx param.Opt[int64] `json:"contourIdx,omitzero"`
	// The relative gain level in dB associated with this boresight or contour. Gain
	// does not apply to service area records. The value is required if type =
	// BORESIGHT or CONTOUR.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// GeoJSON or Well Known Text expression of the boresight point, service area point
	// or region, or the gain contour region in geographic longitude, latitude pairs.
	// Boresight and service area point(s) are represented as a 'Point' or
	// 'MultiPoint', service areas and closed gain contours as 'Polygon', and open
	// contours as 'LineString'. This is an optional convenience field only used for
	// create operations. The system will auto-detect the format (Well Known Text or
	// GeoJSON) and populate both geographyText and geographyJson fields appropriately.
	// A create request must contain one of the geography, geographyText, or
	// geographyJson.
	Geography param.Opt[string] `json:"geography,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	GeographyJson param.Opt[string] `json:"geographyJson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	GeographyNdims param.Opt[int64] `json:"geographyNdims,omitzero"`
	// Geographical spatial_ref_sys for region.
	GeographySrid param.Opt[int64] `json:"geographySrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	GeographyText param.Opt[string] `json:"geographyText,omitzero"`
	// Type of region as projected.
	GeographyType param.Opt[string] `json:"geographyType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The region name within the service area.
	RegionName param.Opt[string] `json:"regionName,omitzero"`
	paramObj
}

func (r BeamContourNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BeamContourNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BeamContourNewParams) UnmarshalJSON(data []byte) error {
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
type BeamContourNewParamsDataMode string

const (
	BeamContourNewParamsDataModeReal      BeamContourNewParamsDataMode = "REAL"
	BeamContourNewParamsDataModeTest      BeamContourNewParamsDataMode = "TEST"
	BeamContourNewParamsDataModeSimulated BeamContourNewParamsDataMode = "SIMULATED"
	BeamContourNewParamsDataModeExercise  BeamContourNewParamsDataMode = "EXERCISE"
)

// The type of object represented in this record (BORESIGHT, CONTOUR, SVC AREA).
// Boresight refers to the point of maximum/peak gain, and should not be confused
// with the 'aim point' of the related beam. Gain contours are regions of coverage
// referenced to the relative gain of the related beam. Service Areas are composed
// of one or more service regions, with each region being either discrete point(s)
// or a continuous contour.
type BeamContourNewParamsType string

const (
	BeamContourNewParamsTypeBoresight BeamContourNewParamsType = "BORESIGHT"
	BeamContourNewParamsTypeContour   BeamContourNewParamsType = "CONTOUR"
	BeamContourNewParamsTypeSvcArea   BeamContourNewParamsType = "SVC AREA"
)

type BeamContourGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BeamContourGetParams]'s query parameters as `url.Values`.
func (r BeamContourGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BeamContourUpdateParams struct {
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
	DataMode BeamContourUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the beam.
	IDBeam string `json:"idBeam,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of object represented in this record (BORESIGHT, CONTOUR, SVC AREA).
	// Boresight refers to the point of maximum/peak gain, and should not be confused
	// with the 'aim point' of the related beam. Gain contours are regions of coverage
	// referenced to the relative gain of the related beam. Service Areas are composed
	// of one or more service regions, with each region being either discrete point(s)
	// or a continuous contour.
	//
	// Any of "BORESIGHT", "CONTOUR", "SVC AREA".
	Type BeamContourUpdateParamsType `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The index number of this contour. The value is required if type = CONTOUR.
	ContourIdx param.Opt[int64] `json:"contourIdx,omitzero"`
	// The relative gain level in dB associated with this boresight or contour. Gain
	// does not apply to service area records. The value is required if type =
	// BORESIGHT or CONTOUR.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// GeoJSON or Well Known Text expression of the boresight point, service area point
	// or region, or the gain contour region in geographic longitude, latitude pairs.
	// Boresight and service area point(s) are represented as a 'Point' or
	// 'MultiPoint', service areas and closed gain contours as 'Polygon', and open
	// contours as 'LineString'. This is an optional convenience field only used for
	// create operations. The system will auto-detect the format (Well Known Text or
	// GeoJSON) and populate both geographyText and geographyJson fields appropriately.
	// A create request must contain one of the geography, geographyText, or
	// geographyJson.
	Geography param.Opt[string] `json:"geography,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	GeographyJson param.Opt[string] `json:"geographyJson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	GeographyNdims param.Opt[int64] `json:"geographyNdims,omitzero"`
	// Geographical spatial_ref_sys for region.
	GeographySrid param.Opt[int64] `json:"geographySrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	GeographyText param.Opt[string] `json:"geographyText,omitzero"`
	// Type of region as projected.
	GeographyType param.Opt[string] `json:"geographyType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The region name within the service area.
	RegionName param.Opt[string] `json:"regionName,omitzero"`
	paramObj
}

func (r BeamContourUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BeamContourUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BeamContourUpdateParams) UnmarshalJSON(data []byte) error {
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
type BeamContourUpdateParamsDataMode string

const (
	BeamContourUpdateParamsDataModeReal      BeamContourUpdateParamsDataMode = "REAL"
	BeamContourUpdateParamsDataModeTest      BeamContourUpdateParamsDataMode = "TEST"
	BeamContourUpdateParamsDataModeSimulated BeamContourUpdateParamsDataMode = "SIMULATED"
	BeamContourUpdateParamsDataModeExercise  BeamContourUpdateParamsDataMode = "EXERCISE"
)

// The type of object represented in this record (BORESIGHT, CONTOUR, SVC AREA).
// Boresight refers to the point of maximum/peak gain, and should not be confused
// with the 'aim point' of the related beam. Gain contours are regions of coverage
// referenced to the relative gain of the related beam. Service Areas are composed
// of one or more service regions, with each region being either discrete point(s)
// or a continuous contour.
type BeamContourUpdateParamsType string

const (
	BeamContourUpdateParamsTypeBoresight BeamContourUpdateParamsType = "BORESIGHT"
	BeamContourUpdateParamsTypeContour   BeamContourUpdateParamsType = "CONTOUR"
	BeamContourUpdateParamsTypeSvcArea   BeamContourUpdateParamsType = "SVC AREA"
)

type BeamContourListParams struct {
	// ID of the beam.
	IDBeam      string           `query:"idBeam,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BeamContourListParams]'s query parameters as `url.Values`.
func (r BeamContourListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BeamContourCountParams struct {
	// ID of the beam.
	IDBeam      string           `query:"idBeam,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BeamContourCountParams]'s query parameters as `url.Values`.
func (r BeamContourCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BeamContourNewBulkParams struct {
	Body []BeamContourNewBulkParamsBody
	paramObj
}

func (r BeamContourNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *BeamContourNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Describes the beam contour associated with a beam entity. Beam contours are the
// geographic representation of the relative gain levels of beam power off of the
// maximum gain boresight points.
//
// The properties ClassificationMarking, DataMode, IDBeam, Source, Type are
// required.
type BeamContourNewBulkParamsBody struct {
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
	// ID of the beam.
	IDBeam string `json:"idBeam,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The type of object represented in this record (BORESIGHT, CONTOUR, SVC AREA).
	// Boresight refers to the point of maximum/peak gain, and should not be confused
	// with the 'aim point' of the related beam. Gain contours are regions of coverage
	// referenced to the relative gain of the related beam. Service Areas are composed
	// of one or more service regions, with each region being either discrete point(s)
	// or a continuous contour.
	//
	// Any of "BORESIGHT", "CONTOUR", "SVC AREA".
	Type string `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The index number of this contour. The value is required if type = CONTOUR.
	ContourIdx param.Opt[int64] `json:"contourIdx,omitzero"`
	// The relative gain level in dB associated with this boresight or contour. Gain
	// does not apply to service area records. The value is required if type =
	// BORESIGHT or CONTOUR.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// GeoJSON or Well Known Text expression of the boresight point, service area point
	// or region, or the gain contour region in geographic longitude, latitude pairs.
	// Boresight and service area point(s) are represented as a 'Point' or
	// 'MultiPoint', service areas and closed gain contours as 'Polygon', and open
	// contours as 'LineString'. This is an optional convenience field only used for
	// create operations. The system will auto-detect the format (Well Known Text or
	// GeoJSON) and populate both geographyText and geographyJson fields appropriately.
	// A create request must contain one of the geography, geographyText, or
	// geographyJson.
	Geography param.Opt[string] `json:"geography,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	GeographyJson param.Opt[string] `json:"geographyJson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	GeographyNdims param.Opt[int64] `json:"geographyNdims,omitzero"`
	// Geographical spatial_ref_sys for region.
	GeographySrid param.Opt[int64] `json:"geographySrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	GeographyText param.Opt[string] `json:"geographyText,omitzero"`
	// Type of region as projected.
	GeographyType param.Opt[string] `json:"geographyType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The region name within the service area.
	RegionName param.Opt[string] `json:"regionName,omitzero"`
	paramObj
}

func (r BeamContourNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow BeamContourNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BeamContourNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BeamContourNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[BeamContourNewBulkParamsBody](
		"type", "BORESIGHT", "CONTOUR", "SVC AREA",
	)
}

type BeamContourTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// ID of the beam.
	IDBeam      string           `query:"idBeam,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BeamContourTupleParams]'s query parameters as `url.Values`.
func (r BeamContourTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
