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

// SurfaceService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSurfaceService] method instead.
type SurfaceService struct {
	Options []option.RequestOption
}

// NewSurfaceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSurfaceService(opts ...option.RequestOption) (r SurfaceService) {
	r = SurfaceService{}
	r.Options = opts
	return
}

// Service operation to take a single Surface as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *SurfaceService) New(ctx context.Context, body SurfaceNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/surface"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Surface. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *SurfaceService) Update(ctx context.Context, id string, body SurfaceUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/surface/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SurfaceService) List(ctx context.Context, query SurfaceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SurfaceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/surface"
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
func (r *SurfaceService) ListAutoPaging(ctx context.Context, query SurfaceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SurfaceListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Surface object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *SurfaceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/surface/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SurfaceService) Count(ctx context.Context, query SurfaceCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/surface/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Surface record by its unique ID passed as a
// path parameter.
func (r *SurfaceService) Get(ctx context.Context, id string, query SurfaceGetParams, opts ...option.RequestOption) (res *SurfaceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/surface/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SurfaceService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SurfaceQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/surface/queryhelp"
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
func (r *SurfaceService) Tuple(ctx context.Context, query SurfaceTupleParams, opts ...option.RequestOption) (res *[]SurfaceTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/surface/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Surface information contains properties related to an airfield's runway,
// taxiway, and parking. The surface types and characteristics can dictate the
// airfield's capability of hosting a specific aircraft.
type SurfaceListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceListResponseDataMode `json:"dataMode,required"`
	// The surface name or identifier.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The surface type of this record (e.g. RUNWAY, TAXIWAY, PARKING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate site identifier provided by the source.
	AltSiteID string `json:"altSiteId"`
	// The surface condition (e.g. GOOD, FAIR, POOR, SERIOUS, FAILED, CLOSED, UNKNOWN).
	Condition string `json:"condition"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKip float64 `json:"ddtWtKip"`
	// The modified max weight allowable on this surface type for a DDT-type (double
	// dual tandem) landing gear configuration, in kilopounds (kip).
	DdtWtKipMod float64 `json:"ddtWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a DDT-type (double dual
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdtWtKipModNote string `json:"ddtWtKipModNote"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKn float64 `json:"ddtWtKN"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKip float64 `json:"ddWtKip"`
	// The modified max weight allowable on this surface type for an FAA 2D-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	DdWtKipMod float64 `json:"ddWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an FAA 2D-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdWtKipModNote string `json:"ddWtKipModNote"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKn float64 `json:"ddWtKN"`
	// WGS-84 latitude of the coordinate representing the end-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	EndLat float64 `json:"endLat"`
	// WGS-84 longitude of the coordinate representing the end-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	EndLon float64 `json:"endLon"`
	// The unique identifier of the Site record that contains location information
	// about this surface.
	IDSite string `json:"idSite"`
	// Load classification number or pavement rating which ranks aircraft on a scale of
	// 1 to 120.
	Lcn int64 `json:"lcn"`
	// The landing distance available for the runway, in feet. Applicable for runway
	// surface types only.
	LdaFt float64 `json:"ldaFt"`
	// The landing distance available for the runway, in meters. Applicable for runway
	// surface types only.
	LdaM float64 `json:"ldaM"`
	// The length of the surface type, in feet. Applicable for runway and parking
	// surface types.
	LengthFt float64 `json:"lengthFt"`
	// The length of the surface type, in meters. Applicable for runway and parking
	// surface types.
	LengthM float64 `json:"lengthM"`
	// Flag indicating the surface has lighting.
	Lighting bool `json:"lighting"`
	// Flag indicating the runway has approach lights. Applicable for runway surface
	// types only.
	LightsAprch bool `json:"lightsAPRCH"`
	// Flag indicating the runway has centerline lights. Applicable for runway surface
	// types only.
	LightsCl bool `json:"lightsCL"`
	// Flag indicating the runway has Optical Landing System (OLS) lights. Applicable
	// for runway surface types only.
	LightsOls bool `json:"lightsOLS"`
	// Flag indicating the runway has Precision Approach Path Indicator (PAPI) lights.
	// Applicable for runway surface types only.
	LightsPapi bool `json:"lightsPAPI"`
	// Flag indicating the runway has Runway End Identifier Lights (REIL). Applicable
	// for runway surface types only.
	LightsReil bool `json:"lightsREIL"`
	// Flag indicating the runway has edge lighting. Applicable for runway surface
	// types only.
	LightsRwy bool `json:"lightsRWY"`
	// Flag indicating the runway has Sequential Flashing (SEQFL) lights. Applicable
	// for runway surface types only.
	LightsSeqfl bool `json:"lightsSEQFL"`
	// Flag indicating the runway has Touchdown Zone lights. Applicable for runway
	// surface types only.
	LightsTdzl bool `json:"lightsTDZL"`
	// Flag indicating the runway lighting is unknown. Applicable for runway surface
	// types only.
	LightsUnkn bool `json:"lightsUNKN"`
	// Flag indicating the runway has Visual Approach Slope Indicator (VASI) lights.
	// Applicable for runway surface types only.
	LightsVasi bool `json:"lightsVASI"`
	// The surface material (e.g. Asphalt, Concrete, Dirt).
	Material string `json:"material"`
	// Flag indicating that one or more obstacles or obstructions exist in the
	// proximity of this surface.
	Obstacle bool `json:"obstacle"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Pavement classification number (PCN) and tire pressure code.
	Pcn string `json:"pcn"`
	// Description of the surface and its dimensions or how it is measured or oriented.
	PointReference string `json:"pointReference"`
	// Flag indicating this is the primary runway. Applicable for runway surface types
	// only.
	Primary bool `json:"primary"`
	// Raw weight bearing capacity value or pavement strength.
	RawWbc string `json:"rawWBC"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilopounds (kip).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKip float64 `json:"sbttWtKip"`
	// The modified max weight allowable on this surface type for an SBTT-type (single
	// belly twin tandem) landing gear configuration, in kilopounds (kip).
	SbttWtKipMod float64 `json:"sbttWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an SBTT-type (single
	// belly twin tandem) landing gear configuration. For more information, contact the
	// provider source.
	SbttWtKipModNote string `json:"sbttWtKipModNote"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilonewtons (kN).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKn float64 `json:"sbttWtKN"`
	// WGS-84 latitude of the coordinate representing the start-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	StartLat float64 `json:"startLat"`
	// WGS-84 longitude of the coordinate representing the start-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	StartLon float64 `json:"startLon"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKip float64 `json:"stWtKip"`
	// The modified max weight allowable on this surface type for an ST-type (single
	// tandem) landing gear configuration, in kilopounds (kip).
	StWtKipMod float64 `json:"stWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an ST-type (single
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	StWtKipModNote string `json:"stWtKipModNote"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKn float64 `json:"stWtKN"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilopounds (kip). Note: The corresponding equivalent
	// field is not converted by the UDL and may or may not be supplied by the
	// provider. The provider/consumer is responsible for all unit conversions.
	SWtKip float64 `json:"sWtKip"`
	// The modified max weight allowable on this surface type for an S-type (single)
	// landing gear configuration, in kilopounds (kip).
	SWtKipMod float64 `json:"sWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an S-type (single)
	// landing gear configuration. For more information, contact the provider source.
	SWtKipModNote string `json:"sWtKipModNote"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	SWtKn float64 `json:"sWtKN"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtkip float64 `json:"tdtWtkip"`
	// The modified max weight allowable on this surface type for a TDT-type (twin
	// delta tandem) landing gear configuration, in kilopounds (kip).
	TdtWtKipMod float64 `json:"tdtWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TDT-type (twin delta
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TdtWtKipModNote string `json:"tdtWtKipModNote"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtKn float64 `json:"tdtWtKN"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKip float64 `json:"trtWtKip"`
	// The modified max weight allowable on this surface type for a TRT-type (triple
	// tandem) landing gear configuration, in kilopounds (kip).
	TrtWtKipMod float64 `json:"trtWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TRT-type (triple
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TrtWtKipModNote string `json:"trtWtKipModNote"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKn float64 `json:"trtWtKN"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKip float64 `json:"ttWtKip"`
	// The modified max weight allowable on this surface type for a GDSS TT-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	TtWtKipMod float64 `json:"ttWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a GDSS TT-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TtWtKipModNote string `json:"ttWtKipModNote"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKn float64 `json:"ttWtKN"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilopounds (kip).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKip float64 `json:"tWtKip"`
	// The modified max weight allowable on this surface type for a T-type (twin
	// (dual)) landing gear configuration, in kilopounds (kip).
	TWtKipMod float64 `json:"tWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a T-type (twin(dual))
	// landing gear configuration. For more information, contact the provider source.
	TWtKipModNote string `json:"tWtKipModNote"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKn float64 `json:"tWtKN"`
	// The width of the surface type, in feet.
	WidthFt float64 `json:"widthFt"`
	// The width of the surface type, in meters.
	WidthM float64 `json:"widthM"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AltSiteID             respjson.Field
		Condition             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DdtWtKip              respjson.Field
		DdtWtKipMod           respjson.Field
		DdtWtKipModNote       respjson.Field
		DdtWtKn               respjson.Field
		DdWtKip               respjson.Field
		DdWtKipMod            respjson.Field
		DdWtKipModNote        respjson.Field
		DdWtKn                respjson.Field
		EndLat                respjson.Field
		EndLon                respjson.Field
		IDSite                respjson.Field
		Lcn                   respjson.Field
		LdaFt                 respjson.Field
		LdaM                  respjson.Field
		LengthFt              respjson.Field
		LengthM               respjson.Field
		Lighting              respjson.Field
		LightsAprch           respjson.Field
		LightsCl              respjson.Field
		LightsOls             respjson.Field
		LightsPapi            respjson.Field
		LightsReil            respjson.Field
		LightsRwy             respjson.Field
		LightsSeqfl           respjson.Field
		LightsTdzl            respjson.Field
		LightsUnkn            respjson.Field
		LightsVasi            respjson.Field
		Material              respjson.Field
		Obstacle              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Pcn                   respjson.Field
		PointReference        respjson.Field
		Primary               respjson.Field
		RawWbc                respjson.Field
		SbttWtKip             respjson.Field
		SbttWtKipMod          respjson.Field
		SbttWtKipModNote      respjson.Field
		SbttWtKn              respjson.Field
		StartLat              respjson.Field
		StartLon              respjson.Field
		StWtKip               respjson.Field
		StWtKipMod            respjson.Field
		StWtKipModNote        respjson.Field
		StWtKn                respjson.Field
		SWtKip                respjson.Field
		SWtKipMod             respjson.Field
		SWtKipModNote         respjson.Field
		SWtKn                 respjson.Field
		TdtWtkip              respjson.Field
		TdtWtKipMod           respjson.Field
		TdtWtKipModNote       respjson.Field
		TdtWtKn               respjson.Field
		TrtWtKip              respjson.Field
		TrtWtKipMod           respjson.Field
		TrtWtKipModNote       respjson.Field
		TrtWtKn               respjson.Field
		TtWtKip               respjson.Field
		TtWtKipMod            respjson.Field
		TtWtKipModNote        respjson.Field
		TtWtKn                respjson.Field
		TWtKip                respjson.Field
		TWtKipMod             respjson.Field
		TWtKipModNote         respjson.Field
		TWtKn                 respjson.Field
		WidthFt               respjson.Field
		WidthM                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SurfaceListResponse) RawJSON() string { return r.JSON.raw }
func (r *SurfaceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceListResponseDataMode string

const (
	SurfaceListResponseDataModeReal      SurfaceListResponseDataMode = "REAL"
	SurfaceListResponseDataModeTest      SurfaceListResponseDataMode = "TEST"
	SurfaceListResponseDataModeSimulated SurfaceListResponseDataMode = "SIMULATED"
	SurfaceListResponseDataModeExercise  SurfaceListResponseDataMode = "EXERCISE"
)

// Surface information contains properties related to an airfield's runway,
// taxiway, and parking. The surface types and characteristics can dictate the
// airfield's capability of hosting a specific aircraft.
type SurfaceGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceGetResponseDataMode `json:"dataMode,required"`
	// The surface name or identifier.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The surface type of this record (e.g. RUNWAY, TAXIWAY, PARKING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate site identifier provided by the source.
	AltSiteID string `json:"altSiteId"`
	// The surface condition (e.g. GOOD, FAIR, POOR, SERIOUS, FAILED, CLOSED, UNKNOWN).
	Condition string `json:"condition"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKip float64 `json:"ddtWtKip"`
	// The modified max weight allowable on this surface type for a DDT-type (double
	// dual tandem) landing gear configuration, in kilopounds (kip).
	DdtWtKipMod float64 `json:"ddtWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a DDT-type (double dual
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdtWtKipModNote string `json:"ddtWtKipModNote"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKn float64 `json:"ddtWtKN"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKip float64 `json:"ddWtKip"`
	// The modified max weight allowable on this surface type for an FAA 2D-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	DdWtKipMod float64 `json:"ddWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an FAA 2D-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdWtKipModNote string `json:"ddWtKipModNote"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKn float64 `json:"ddWtKN"`
	// WGS-84 latitude of the coordinate representing the end-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	EndLat float64 `json:"endLat"`
	// WGS-84 longitude of the coordinate representing the end-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	EndLon float64 `json:"endLon"`
	// The unique identifier of the Site record that contains location information
	// about this surface.
	IDSite string `json:"idSite"`
	// Load classification number or pavement rating which ranks aircraft on a scale of
	// 1 to 120.
	Lcn int64 `json:"lcn"`
	// The landing distance available for the runway, in feet. Applicable for runway
	// surface types only.
	LdaFt float64 `json:"ldaFt"`
	// The landing distance available for the runway, in meters. Applicable for runway
	// surface types only.
	LdaM float64 `json:"ldaM"`
	// The length of the surface type, in feet. Applicable for runway and parking
	// surface types.
	LengthFt float64 `json:"lengthFt"`
	// The length of the surface type, in meters. Applicable for runway and parking
	// surface types.
	LengthM float64 `json:"lengthM"`
	// Flag indicating the surface has lighting.
	Lighting bool `json:"lighting"`
	// Flag indicating the runway has approach lights. Applicable for runway surface
	// types only.
	LightsAprch bool `json:"lightsAPRCH"`
	// Flag indicating the runway has centerline lights. Applicable for runway surface
	// types only.
	LightsCl bool `json:"lightsCL"`
	// Flag indicating the runway has Optical Landing System (OLS) lights. Applicable
	// for runway surface types only.
	LightsOls bool `json:"lightsOLS"`
	// Flag indicating the runway has Precision Approach Path Indicator (PAPI) lights.
	// Applicable for runway surface types only.
	LightsPapi bool `json:"lightsPAPI"`
	// Flag indicating the runway has Runway End Identifier Lights (REIL). Applicable
	// for runway surface types only.
	LightsReil bool `json:"lightsREIL"`
	// Flag indicating the runway has edge lighting. Applicable for runway surface
	// types only.
	LightsRwy bool `json:"lightsRWY"`
	// Flag indicating the runway has Sequential Flashing (SEQFL) lights. Applicable
	// for runway surface types only.
	LightsSeqfl bool `json:"lightsSEQFL"`
	// Flag indicating the runway has Touchdown Zone lights. Applicable for runway
	// surface types only.
	LightsTdzl bool `json:"lightsTDZL"`
	// Flag indicating the runway lighting is unknown. Applicable for runway surface
	// types only.
	LightsUnkn bool `json:"lightsUNKN"`
	// Flag indicating the runway has Visual Approach Slope Indicator (VASI) lights.
	// Applicable for runway surface types only.
	LightsVasi bool `json:"lightsVASI"`
	// The surface material (e.g. Asphalt, Concrete, Dirt).
	Material string `json:"material"`
	// Flag indicating that one or more obstacles or obstructions exist in the
	// proximity of this surface.
	Obstacle bool `json:"obstacle"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Pavement classification number (PCN) and tire pressure code.
	Pcn string `json:"pcn"`
	// Description of the surface and its dimensions or how it is measured or oriented.
	PointReference string `json:"pointReference"`
	// Flag indicating this is the primary runway. Applicable for runway surface types
	// only.
	Primary bool `json:"primary"`
	// Raw weight bearing capacity value or pavement strength.
	RawWbc string `json:"rawWBC"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilopounds (kip).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKip float64 `json:"sbttWtKip"`
	// The modified max weight allowable on this surface type for an SBTT-type (single
	// belly twin tandem) landing gear configuration, in kilopounds (kip).
	SbttWtKipMod float64 `json:"sbttWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an SBTT-type (single
	// belly twin tandem) landing gear configuration. For more information, contact the
	// provider source.
	SbttWtKipModNote string `json:"sbttWtKipModNote"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilonewtons (kN).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKn float64 `json:"sbttWtKN"`
	// WGS-84 latitude of the coordinate representing the start-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	StartLat float64 `json:"startLat"`
	// WGS-84 longitude of the coordinate representing the start-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	StartLon float64 `json:"startLon"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKip float64 `json:"stWtKip"`
	// The modified max weight allowable on this surface type for an ST-type (single
	// tandem) landing gear configuration, in kilopounds (kip).
	StWtKipMod float64 `json:"stWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an ST-type (single
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	StWtKipModNote string `json:"stWtKipModNote"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKn float64 `json:"stWtKN"`
	// An array of SurfaceNavigation objects representing information about
	// obstructions proximal to this surface. This is a read-only field that will be
	// populated with any associated SurfaceObstruction objects for queries that return
	// the full schema.
	SurfaceObstructions []SurfaceGetResponseSurfaceObstruction `json:"surfaceObstructions"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilopounds (kip). Note: The corresponding equivalent
	// field is not converted by the UDL and may or may not be supplied by the
	// provider. The provider/consumer is responsible for all unit conversions.
	SWtKip float64 `json:"sWtKip"`
	// The modified max weight allowable on this surface type for an S-type (single)
	// landing gear configuration, in kilopounds (kip).
	SWtKipMod float64 `json:"sWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an S-type (single)
	// landing gear configuration. For more information, contact the provider source.
	SWtKipModNote string `json:"sWtKipModNote"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	SWtKn float64 `json:"sWtKN"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtkip float64 `json:"tdtWtkip"`
	// The modified max weight allowable on this surface type for a TDT-type (twin
	// delta tandem) landing gear configuration, in kilopounds (kip).
	TdtWtKipMod float64 `json:"tdtWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TDT-type (twin delta
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TdtWtKipModNote string `json:"tdtWtKipModNote"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtKn float64 `json:"tdtWtKN"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKip float64 `json:"trtWtKip"`
	// The modified max weight allowable on this surface type for a TRT-type (triple
	// tandem) landing gear configuration, in kilopounds (kip).
	TrtWtKipMod float64 `json:"trtWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TRT-type (triple
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TrtWtKipModNote string `json:"trtWtKipModNote"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKn float64 `json:"trtWtKN"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKip float64 `json:"ttWtKip"`
	// The modified max weight allowable on this surface type for a GDSS TT-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	TtWtKipMod float64 `json:"ttWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a GDSS TT-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TtWtKipModNote string `json:"ttWtKipModNote"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKn float64 `json:"ttWtKN"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilopounds (kip).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKip float64 `json:"tWtKip"`
	// The modified max weight allowable on this surface type for a T-type (twin
	// (dual)) landing gear configuration, in kilopounds (kip).
	TWtKipMod float64 `json:"tWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a T-type (twin(dual))
	// landing gear configuration. For more information, contact the provider source.
	TWtKipModNote string `json:"tWtKipModNote"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKn float64 `json:"tWtKN"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The width of the surface type, in feet.
	WidthFt float64 `json:"widthFt"`
	// The width of the surface type, in meters.
	WidthM float64 `json:"widthM"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AltSiteID             respjson.Field
		Condition             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DdtWtKip              respjson.Field
		DdtWtKipMod           respjson.Field
		DdtWtKipModNote       respjson.Field
		DdtWtKn               respjson.Field
		DdWtKip               respjson.Field
		DdWtKipMod            respjson.Field
		DdWtKipModNote        respjson.Field
		DdWtKn                respjson.Field
		EndLat                respjson.Field
		EndLon                respjson.Field
		IDSite                respjson.Field
		Lcn                   respjson.Field
		LdaFt                 respjson.Field
		LdaM                  respjson.Field
		LengthFt              respjson.Field
		LengthM               respjson.Field
		Lighting              respjson.Field
		LightsAprch           respjson.Field
		LightsCl              respjson.Field
		LightsOls             respjson.Field
		LightsPapi            respjson.Field
		LightsReil            respjson.Field
		LightsRwy             respjson.Field
		LightsSeqfl           respjson.Field
		LightsTdzl            respjson.Field
		LightsUnkn            respjson.Field
		LightsVasi            respjson.Field
		Material              respjson.Field
		Obstacle              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Pcn                   respjson.Field
		PointReference        respjson.Field
		Primary               respjson.Field
		RawWbc                respjson.Field
		SbttWtKip             respjson.Field
		SbttWtKipMod          respjson.Field
		SbttWtKipModNote      respjson.Field
		SbttWtKn              respjson.Field
		StartLat              respjson.Field
		StartLon              respjson.Field
		StWtKip               respjson.Field
		StWtKipMod            respjson.Field
		StWtKipModNote        respjson.Field
		StWtKn                respjson.Field
		SurfaceObstructions   respjson.Field
		SWtKip                respjson.Field
		SWtKipMod             respjson.Field
		SWtKipModNote         respjson.Field
		SWtKn                 respjson.Field
		TdtWtkip              respjson.Field
		TdtWtKipMod           respjson.Field
		TdtWtKipModNote       respjson.Field
		TdtWtKn               respjson.Field
		TrtWtKip              respjson.Field
		TrtWtKipMod           respjson.Field
		TrtWtKipModNote       respjson.Field
		TrtWtKn               respjson.Field
		TtWtKip               respjson.Field
		TtWtKipMod            respjson.Field
		TtWtKipModNote        respjson.Field
		TtWtKn                respjson.Field
		TWtKip                respjson.Field
		TWtKipMod             respjson.Field
		TWtKipModNote         respjson.Field
		TWtKn                 respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		WidthFt               respjson.Field
		WidthM                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SurfaceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SurfaceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceGetResponseDataMode string

const (
	SurfaceGetResponseDataModeReal      SurfaceGetResponseDataMode = "REAL"
	SurfaceGetResponseDataModeTest      SurfaceGetResponseDataMode = "TEST"
	SurfaceGetResponseDataModeSimulated SurfaceGetResponseDataMode = "SIMULATED"
	SurfaceGetResponseDataModeExercise  SurfaceGetResponseDataMode = "EXERCISE"
)

type SurfaceGetResponseSurfaceObstruction struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,required"`
	// The unique identifier of the associated surface record. This field is required
	// when posting, updating, or deleting a SurfaceObstruction record.
	IDSurface string `json:"idSurface,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an advisory for usage.
	AdvisoryRequired []string `json:"advisoryRequired"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an approval for usage.
	ApprovalRequired []string `json:"approvalRequired"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The distance from the surface center line to this obstruction, in feet.
	DistanceFromCenterLine float64 `json:"distanceFromCenterLine"`
	// The distance from the surface edge to this obstruction, in feet.
	DistanceFromEdge float64 `json:"distanceFromEdge"`
	// The distance from the surface threshold to this obstruction, in feet.
	DistanceFromThreshold float64 `json:"distanceFromThreshold"`
	// The unique identifier of the associated NavigationalObstruction record.
	IDNavigationalObstruction string `json:"idNavigationalObstruction"`
	// Description of this surface obstruction.
	ObstructionDesc string `json:"obstructionDesc"`
	// The height above ground level of the surface obstruction, in feet.
	ObstructionHeight float64 `json:"obstructionHeight"`
	// A code that indicates which side of the surface end is affected by this
	// obstruction.
	ObstructionSideCode string `json:"obstructionSideCode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking     respjson.Field
		DataMode                  respjson.Field
		IDSurface                 respjson.Field
		Source                    respjson.Field
		ID                        respjson.Field
		AdvisoryRequired          respjson.Field
		ApprovalRequired          respjson.Field
		CreatedAt                 respjson.Field
		CreatedBy                 respjson.Field
		DistanceFromCenterLine    respjson.Field
		DistanceFromEdge          respjson.Field
		DistanceFromThreshold     respjson.Field
		IDNavigationalObstruction respjson.Field
		ObstructionDesc           respjson.Field
		ObstructionHeight         respjson.Field
		ObstructionSideCode       respjson.Field
		Origin                    respjson.Field
		OrigNetwork               respjson.Field
		SourceDl                  respjson.Field
		UpdatedAt                 respjson.Field
		UpdatedBy                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SurfaceGetResponseSurfaceObstruction) RawJSON() string { return r.JSON.raw }
func (r *SurfaceGetResponseSurfaceObstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SurfaceQueryhelpResponse struct {
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
func (r SurfaceQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SurfaceQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Surface information contains properties related to an airfield's runway,
// taxiway, and parking. The surface types and characteristics can dictate the
// airfield's capability of hosting a specific aircraft.
type SurfaceTupleResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceTupleResponseDataMode `json:"dataMode,required"`
	// The surface name or identifier.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The surface type of this record (e.g. RUNWAY, TAXIWAY, PARKING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate site identifier provided by the source.
	AltSiteID string `json:"altSiteId"`
	// The surface condition (e.g. GOOD, FAIR, POOR, SERIOUS, FAILED, CLOSED, UNKNOWN).
	Condition string `json:"condition"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKip float64 `json:"ddtWtKip"`
	// The modified max weight allowable on this surface type for a DDT-type (double
	// dual tandem) landing gear configuration, in kilopounds (kip).
	DdtWtKipMod float64 `json:"ddtWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a DDT-type (double dual
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdtWtKipModNote string `json:"ddtWtKipModNote"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKn float64 `json:"ddtWtKN"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKip float64 `json:"ddWtKip"`
	// The modified max weight allowable on this surface type for an FAA 2D-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	DdWtKipMod float64 `json:"ddWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an FAA 2D-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdWtKipModNote string `json:"ddWtKipModNote"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKn float64 `json:"ddWtKN"`
	// WGS-84 latitude of the coordinate representing the end-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	EndLat float64 `json:"endLat"`
	// WGS-84 longitude of the coordinate representing the end-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	EndLon float64 `json:"endLon"`
	// The unique identifier of the Site record that contains location information
	// about this surface.
	IDSite string `json:"idSite"`
	// Load classification number or pavement rating which ranks aircraft on a scale of
	// 1 to 120.
	Lcn int64 `json:"lcn"`
	// The landing distance available for the runway, in feet. Applicable for runway
	// surface types only.
	LdaFt float64 `json:"ldaFt"`
	// The landing distance available for the runway, in meters. Applicable for runway
	// surface types only.
	LdaM float64 `json:"ldaM"`
	// The length of the surface type, in feet. Applicable for runway and parking
	// surface types.
	LengthFt float64 `json:"lengthFt"`
	// The length of the surface type, in meters. Applicable for runway and parking
	// surface types.
	LengthM float64 `json:"lengthM"`
	// Flag indicating the surface has lighting.
	Lighting bool `json:"lighting"`
	// Flag indicating the runway has approach lights. Applicable for runway surface
	// types only.
	LightsAprch bool `json:"lightsAPRCH"`
	// Flag indicating the runway has centerline lights. Applicable for runway surface
	// types only.
	LightsCl bool `json:"lightsCL"`
	// Flag indicating the runway has Optical Landing System (OLS) lights. Applicable
	// for runway surface types only.
	LightsOls bool `json:"lightsOLS"`
	// Flag indicating the runway has Precision Approach Path Indicator (PAPI) lights.
	// Applicable for runway surface types only.
	LightsPapi bool `json:"lightsPAPI"`
	// Flag indicating the runway has Runway End Identifier Lights (REIL). Applicable
	// for runway surface types only.
	LightsReil bool `json:"lightsREIL"`
	// Flag indicating the runway has edge lighting. Applicable for runway surface
	// types only.
	LightsRwy bool `json:"lightsRWY"`
	// Flag indicating the runway has Sequential Flashing (SEQFL) lights. Applicable
	// for runway surface types only.
	LightsSeqfl bool `json:"lightsSEQFL"`
	// Flag indicating the runway has Touchdown Zone lights. Applicable for runway
	// surface types only.
	LightsTdzl bool `json:"lightsTDZL"`
	// Flag indicating the runway lighting is unknown. Applicable for runway surface
	// types only.
	LightsUnkn bool `json:"lightsUNKN"`
	// Flag indicating the runway has Visual Approach Slope Indicator (VASI) lights.
	// Applicable for runway surface types only.
	LightsVasi bool `json:"lightsVASI"`
	// The surface material (e.g. Asphalt, Concrete, Dirt).
	Material string `json:"material"`
	// Flag indicating that one or more obstacles or obstructions exist in the
	// proximity of this surface.
	Obstacle bool `json:"obstacle"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Pavement classification number (PCN) and tire pressure code.
	Pcn string `json:"pcn"`
	// Description of the surface and its dimensions or how it is measured or oriented.
	PointReference string `json:"pointReference"`
	// Flag indicating this is the primary runway. Applicable for runway surface types
	// only.
	Primary bool `json:"primary"`
	// Raw weight bearing capacity value or pavement strength.
	RawWbc string `json:"rawWBC"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilopounds (kip).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKip float64 `json:"sbttWtKip"`
	// The modified max weight allowable on this surface type for an SBTT-type (single
	// belly twin tandem) landing gear configuration, in kilopounds (kip).
	SbttWtKipMod float64 `json:"sbttWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an SBTT-type (single
	// belly twin tandem) landing gear configuration. For more information, contact the
	// provider source.
	SbttWtKipModNote string `json:"sbttWtKipModNote"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilonewtons (kN).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKn float64 `json:"sbttWtKN"`
	// WGS-84 latitude of the coordinate representing the start-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	StartLat float64 `json:"startLat"`
	// WGS-84 longitude of the coordinate representing the start-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	StartLon float64 `json:"startLon"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKip float64 `json:"stWtKip"`
	// The modified max weight allowable on this surface type for an ST-type (single
	// tandem) landing gear configuration, in kilopounds (kip).
	StWtKipMod float64 `json:"stWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an ST-type (single
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	StWtKipModNote string `json:"stWtKipModNote"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKn float64 `json:"stWtKN"`
	// An array of SurfaceNavigation objects representing information about
	// obstructions proximal to this surface. This is a read-only field that will be
	// populated with any associated SurfaceObstruction objects for queries that return
	// the full schema.
	SurfaceObstructions []SurfaceTupleResponseSurfaceObstruction `json:"surfaceObstructions"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilopounds (kip). Note: The corresponding equivalent
	// field is not converted by the UDL and may or may not be supplied by the
	// provider. The provider/consumer is responsible for all unit conversions.
	SWtKip float64 `json:"sWtKip"`
	// The modified max weight allowable on this surface type for an S-type (single)
	// landing gear configuration, in kilopounds (kip).
	SWtKipMod float64 `json:"sWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an S-type (single)
	// landing gear configuration. For more information, contact the provider source.
	SWtKipModNote string `json:"sWtKipModNote"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	SWtKn float64 `json:"sWtKN"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtkip float64 `json:"tdtWtkip"`
	// The modified max weight allowable on this surface type for a TDT-type (twin
	// delta tandem) landing gear configuration, in kilopounds (kip).
	TdtWtKipMod float64 `json:"tdtWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TDT-type (twin delta
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TdtWtKipModNote string `json:"tdtWtKipModNote"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtKn float64 `json:"tdtWtKN"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKip float64 `json:"trtWtKip"`
	// The modified max weight allowable on this surface type for a TRT-type (triple
	// tandem) landing gear configuration, in kilopounds (kip).
	TrtWtKipMod float64 `json:"trtWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TRT-type (triple
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TrtWtKipModNote string `json:"trtWtKipModNote"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKn float64 `json:"trtWtKN"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKip float64 `json:"ttWtKip"`
	// The modified max weight allowable on this surface type for a GDSS TT-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	TtWtKipMod float64 `json:"ttWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a GDSS TT-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TtWtKipModNote string `json:"ttWtKipModNote"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKn float64 `json:"ttWtKN"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilopounds (kip).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKip float64 `json:"tWtKip"`
	// The modified max weight allowable on this surface type for a T-type (twin
	// (dual)) landing gear configuration, in kilopounds (kip).
	TWtKipMod float64 `json:"tWtKipMod"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a T-type (twin(dual))
	// landing gear configuration. For more information, contact the provider source.
	TWtKipModNote string `json:"tWtKipModNote"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKn float64 `json:"tWtKN"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The width of the surface type, in feet.
	WidthFt float64 `json:"widthFt"`
	// The width of the surface type, in meters.
	WidthM float64 `json:"widthM"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AltSiteID             respjson.Field
		Condition             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DdtWtKip              respjson.Field
		DdtWtKipMod           respjson.Field
		DdtWtKipModNote       respjson.Field
		DdtWtKn               respjson.Field
		DdWtKip               respjson.Field
		DdWtKipMod            respjson.Field
		DdWtKipModNote        respjson.Field
		DdWtKn                respjson.Field
		EndLat                respjson.Field
		EndLon                respjson.Field
		IDSite                respjson.Field
		Lcn                   respjson.Field
		LdaFt                 respjson.Field
		LdaM                  respjson.Field
		LengthFt              respjson.Field
		LengthM               respjson.Field
		Lighting              respjson.Field
		LightsAprch           respjson.Field
		LightsCl              respjson.Field
		LightsOls             respjson.Field
		LightsPapi            respjson.Field
		LightsReil            respjson.Field
		LightsRwy             respjson.Field
		LightsSeqfl           respjson.Field
		LightsTdzl            respjson.Field
		LightsUnkn            respjson.Field
		LightsVasi            respjson.Field
		Material              respjson.Field
		Obstacle              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Pcn                   respjson.Field
		PointReference        respjson.Field
		Primary               respjson.Field
		RawWbc                respjson.Field
		SbttWtKip             respjson.Field
		SbttWtKipMod          respjson.Field
		SbttWtKipModNote      respjson.Field
		SbttWtKn              respjson.Field
		StartLat              respjson.Field
		StartLon              respjson.Field
		StWtKip               respjson.Field
		StWtKipMod            respjson.Field
		StWtKipModNote        respjson.Field
		StWtKn                respjson.Field
		SurfaceObstructions   respjson.Field
		SWtKip                respjson.Field
		SWtKipMod             respjson.Field
		SWtKipModNote         respjson.Field
		SWtKn                 respjson.Field
		TdtWtkip              respjson.Field
		TdtWtKipMod           respjson.Field
		TdtWtKipModNote       respjson.Field
		TdtWtKn               respjson.Field
		TrtWtKip              respjson.Field
		TrtWtKipMod           respjson.Field
		TrtWtKipModNote       respjson.Field
		TrtWtKn               respjson.Field
		TtWtKip               respjson.Field
		TtWtKipMod            respjson.Field
		TtWtKipModNote        respjson.Field
		TtWtKn                respjson.Field
		TWtKip                respjson.Field
		TWtKipMod             respjson.Field
		TWtKipModNote         respjson.Field
		TWtKn                 respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		WidthFt               respjson.Field
		WidthM                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SurfaceTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SurfaceTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceTupleResponseDataMode string

const (
	SurfaceTupleResponseDataModeReal      SurfaceTupleResponseDataMode = "REAL"
	SurfaceTupleResponseDataModeTest      SurfaceTupleResponseDataMode = "TEST"
	SurfaceTupleResponseDataModeSimulated SurfaceTupleResponseDataMode = "SIMULATED"
	SurfaceTupleResponseDataModeExercise  SurfaceTupleResponseDataMode = "EXERCISE"
)

type SurfaceTupleResponseSurfaceObstruction struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode string `json:"dataMode,required"`
	// The unique identifier of the associated surface record. This field is required
	// when posting, updating, or deleting a SurfaceObstruction record.
	IDSurface string `json:"idSurface,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an advisory for usage.
	AdvisoryRequired []string `json:"advisoryRequired"`
	// Array of all vehicles that are affected by this obstruction at the surface
	// end-point, and require an approval for usage.
	ApprovalRequired []string `json:"approvalRequired"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The distance from the surface center line to this obstruction, in feet.
	DistanceFromCenterLine float64 `json:"distanceFromCenterLine"`
	// The distance from the surface edge to this obstruction, in feet.
	DistanceFromEdge float64 `json:"distanceFromEdge"`
	// The distance from the surface threshold to this obstruction, in feet.
	DistanceFromThreshold float64 `json:"distanceFromThreshold"`
	// The unique identifier of the associated NavigationalObstruction record.
	IDNavigationalObstruction string `json:"idNavigationalObstruction"`
	// Description of this surface obstruction.
	ObstructionDesc string `json:"obstructionDesc"`
	// The height above ground level of the surface obstruction, in feet.
	ObstructionHeight float64 `json:"obstructionHeight"`
	// A code that indicates which side of the surface end is affected by this
	// obstruction.
	ObstructionSideCode string `json:"obstructionSideCode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking     respjson.Field
		DataMode                  respjson.Field
		IDSurface                 respjson.Field
		Source                    respjson.Field
		ID                        respjson.Field
		AdvisoryRequired          respjson.Field
		ApprovalRequired          respjson.Field
		CreatedAt                 respjson.Field
		CreatedBy                 respjson.Field
		DistanceFromCenterLine    respjson.Field
		DistanceFromEdge          respjson.Field
		DistanceFromThreshold     respjson.Field
		IDNavigationalObstruction respjson.Field
		ObstructionDesc           respjson.Field
		ObstructionHeight         respjson.Field
		ObstructionSideCode       respjson.Field
		Origin                    respjson.Field
		OrigNetwork               respjson.Field
		SourceDl                  respjson.Field
		UpdatedAt                 respjson.Field
		UpdatedBy                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SurfaceTupleResponseSurfaceObstruction) RawJSON() string { return r.JSON.raw }
func (r *SurfaceTupleResponseSurfaceObstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SurfaceNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The surface name or identifier.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The surface type of this record (e.g. RUNWAY, TAXIWAY, PARKING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate site identifier provided by the source.
	AltSiteID param.Opt[string] `json:"altSiteId,omitzero"`
	// The surface condition (e.g. GOOD, FAIR, POOR, SERIOUS, FAILED, CLOSED, UNKNOWN).
	Condition param.Opt[string] `json:"condition,omitzero"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKip param.Opt[float64] `json:"ddtWtKip,omitzero"`
	// The modified max weight allowable on this surface type for a DDT-type (double
	// dual tandem) landing gear configuration, in kilopounds (kip).
	DdtWtKipMod param.Opt[float64] `json:"ddtWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a DDT-type (double dual
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdtWtKipModNote param.Opt[string] `json:"ddtWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKn param.Opt[float64] `json:"ddtWtKN,omitzero"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKip param.Opt[float64] `json:"ddWtKip,omitzero"`
	// The modified max weight allowable on this surface type for an FAA 2D-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	DdWtKipMod param.Opt[float64] `json:"ddWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an FAA 2D-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdWtKipModNote param.Opt[string] `json:"ddWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKn param.Opt[float64] `json:"ddWtKN,omitzero"`
	// WGS-84 latitude of the coordinate representing the end-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	EndLat param.Opt[float64] `json:"endLat,omitzero"`
	// WGS-84 longitude of the coordinate representing the end-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	EndLon param.Opt[float64] `json:"endLon,omitzero"`
	// The unique identifier of the Site record that contains location information
	// about this surface.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Load classification number or pavement rating which ranks aircraft on a scale of
	// 1 to 120.
	Lcn param.Opt[int64] `json:"lcn,omitzero"`
	// The landing distance available for the runway, in feet. Applicable for runway
	// surface types only.
	LdaFt param.Opt[float64] `json:"ldaFt,omitzero"`
	// The landing distance available for the runway, in meters. Applicable for runway
	// surface types only.
	LdaM param.Opt[float64] `json:"ldaM,omitzero"`
	// The length of the surface type, in feet. Applicable for runway and parking
	// surface types.
	LengthFt param.Opt[float64] `json:"lengthFt,omitzero"`
	// The length of the surface type, in meters. Applicable for runway and parking
	// surface types.
	LengthM param.Opt[float64] `json:"lengthM,omitzero"`
	// Flag indicating the surface has lighting.
	Lighting param.Opt[bool] `json:"lighting,omitzero"`
	// Flag indicating the runway has approach lights. Applicable for runway surface
	// types only.
	LightsAprch param.Opt[bool] `json:"lightsAPRCH,omitzero"`
	// Flag indicating the runway has centerline lights. Applicable for runway surface
	// types only.
	LightsCl param.Opt[bool] `json:"lightsCL,omitzero"`
	// Flag indicating the runway has Optical Landing System (OLS) lights. Applicable
	// for runway surface types only.
	LightsOls param.Opt[bool] `json:"lightsOLS,omitzero"`
	// Flag indicating the runway has Precision Approach Path Indicator (PAPI) lights.
	// Applicable for runway surface types only.
	LightsPapi param.Opt[bool] `json:"lightsPAPI,omitzero"`
	// Flag indicating the runway has Runway End Identifier Lights (REIL). Applicable
	// for runway surface types only.
	LightsReil param.Opt[bool] `json:"lightsREIL,omitzero"`
	// Flag indicating the runway has edge lighting. Applicable for runway surface
	// types only.
	LightsRwy param.Opt[bool] `json:"lightsRWY,omitzero"`
	// Flag indicating the runway has Sequential Flashing (SEQFL) lights. Applicable
	// for runway surface types only.
	LightsSeqfl param.Opt[bool] `json:"lightsSEQFL,omitzero"`
	// Flag indicating the runway has Touchdown Zone lights. Applicable for runway
	// surface types only.
	LightsTdzl param.Opt[bool] `json:"lightsTDZL,omitzero"`
	// Flag indicating the runway lighting is unknown. Applicable for runway surface
	// types only.
	LightsUnkn param.Opt[bool] `json:"lightsUNKN,omitzero"`
	// Flag indicating the runway has Visual Approach Slope Indicator (VASI) lights.
	// Applicable for runway surface types only.
	LightsVasi param.Opt[bool] `json:"lightsVASI,omitzero"`
	// The surface material (e.g. Asphalt, Concrete, Dirt).
	Material param.Opt[string] `json:"material,omitzero"`
	// Flag indicating that one or more obstacles or obstructions exist in the
	// proximity of this surface.
	Obstacle param.Opt[bool] `json:"obstacle,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Pavement classification number (PCN) and tire pressure code.
	Pcn param.Opt[string] `json:"pcn,omitzero"`
	// Description of the surface and its dimensions or how it is measured or oriented.
	PointReference param.Opt[string] `json:"pointReference,omitzero"`
	// Flag indicating this is the primary runway. Applicable for runway surface types
	// only.
	Primary param.Opt[bool] `json:"primary,omitzero"`
	// Raw weight bearing capacity value or pavement strength.
	RawWbc param.Opt[string] `json:"rawWBC,omitzero"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilopounds (kip).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKip param.Opt[float64] `json:"sbttWtKip,omitzero"`
	// The modified max weight allowable on this surface type for an SBTT-type (single
	// belly twin tandem) landing gear configuration, in kilopounds (kip).
	SbttWtKipMod param.Opt[float64] `json:"sbttWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an SBTT-type (single
	// belly twin tandem) landing gear configuration. For more information, contact the
	// provider source.
	SbttWtKipModNote param.Opt[string] `json:"sbttWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilonewtons (kN).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKn param.Opt[float64] `json:"sbttWtKN,omitzero"`
	// WGS-84 latitude of the coordinate representing the start-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	StartLat param.Opt[float64] `json:"startLat,omitzero"`
	// WGS-84 longitude of the coordinate representing the start-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	StartLon param.Opt[float64] `json:"startLon,omitzero"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKip param.Opt[float64] `json:"stWtKip,omitzero"`
	// The modified max weight allowable on this surface type for an ST-type (single
	// tandem) landing gear configuration, in kilopounds (kip).
	StWtKipMod param.Opt[float64] `json:"stWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an ST-type (single
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	StWtKipModNote param.Opt[string] `json:"stWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKn param.Opt[float64] `json:"stWtKN,omitzero"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilopounds (kip). Note: The corresponding equivalent
	// field is not converted by the UDL and may or may not be supplied by the
	// provider. The provider/consumer is responsible for all unit conversions.
	SWtKip param.Opt[float64] `json:"sWtKip,omitzero"`
	// The modified max weight allowable on this surface type for an S-type (single)
	// landing gear configuration, in kilopounds (kip).
	SWtKipMod param.Opt[float64] `json:"sWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an S-type (single)
	// landing gear configuration. For more information, contact the provider source.
	SWtKipModNote param.Opt[string] `json:"sWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	SWtKn param.Opt[float64] `json:"sWtKN,omitzero"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtkip param.Opt[float64] `json:"tdtWtkip,omitzero"`
	// The modified max weight allowable on this surface type for a TDT-type (twin
	// delta tandem) landing gear configuration, in kilopounds (kip).
	TdtWtKipMod param.Opt[float64] `json:"tdtWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TDT-type (twin delta
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TdtWtKipModNote param.Opt[string] `json:"tdtWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtKn param.Opt[float64] `json:"tdtWtKN,omitzero"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKip param.Opt[float64] `json:"trtWtKip,omitzero"`
	// The modified max weight allowable on this surface type for a TRT-type (triple
	// tandem) landing gear configuration, in kilopounds (kip).
	TrtWtKipMod param.Opt[float64] `json:"trtWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TRT-type (triple
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TrtWtKipModNote param.Opt[string] `json:"trtWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKn param.Opt[float64] `json:"trtWtKN,omitzero"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKip param.Opt[float64] `json:"ttWtKip,omitzero"`
	// The modified max weight allowable on this surface type for a GDSS TT-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	TtWtKipMod param.Opt[float64] `json:"ttWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a GDSS TT-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TtWtKipModNote param.Opt[string] `json:"ttWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKn param.Opt[float64] `json:"ttWtKN,omitzero"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilopounds (kip).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKip param.Opt[float64] `json:"tWtKip,omitzero"`
	// The modified max weight allowable on this surface type for a T-type (twin
	// (dual)) landing gear configuration, in kilopounds (kip).
	TWtKipMod param.Opt[float64] `json:"tWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a T-type (twin(dual))
	// landing gear configuration. For more information, contact the provider source.
	TWtKipModNote param.Opt[string] `json:"tWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKn param.Opt[float64] `json:"tWtKN,omitzero"`
	// The width of the surface type, in feet.
	WidthFt param.Opt[float64] `json:"widthFt,omitzero"`
	// The width of the surface type, in meters.
	WidthM param.Opt[float64] `json:"widthM,omitzero"`
	paramObj
}

func (r SurfaceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SurfaceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SurfaceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceNewParamsDataMode string

const (
	SurfaceNewParamsDataModeReal      SurfaceNewParamsDataMode = "REAL"
	SurfaceNewParamsDataModeTest      SurfaceNewParamsDataMode = "TEST"
	SurfaceNewParamsDataModeSimulated SurfaceNewParamsDataMode = "SIMULATED"
	SurfaceNewParamsDataModeExercise  SurfaceNewParamsDataMode = "EXERCISE"
)

type SurfaceUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode SurfaceUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The surface name or identifier.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The surface type of this record (e.g. RUNWAY, TAXIWAY, PARKING).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate site identifier provided by the source.
	AltSiteID param.Opt[string] `json:"altSiteId,omitzero"`
	// The surface condition (e.g. GOOD, FAIR, POOR, SERIOUS, FAILED, CLOSED, UNKNOWN).
	Condition param.Opt[string] `json:"condition,omitzero"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKip param.Opt[float64] `json:"ddtWtKip,omitzero"`
	// The modified max weight allowable on this surface type for a DDT-type (double
	// dual tandem) landing gear configuration, in kilopounds (kip).
	DdtWtKipMod param.Opt[float64] `json:"ddtWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a DDT-type (double dual
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdtWtKipModNote param.Opt[string] `json:"ddtWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a DDT-type (double dual
	// tandem) landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdtWtKn param.Opt[float64] `json:"ddtWtKN,omitzero"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKip param.Opt[float64] `json:"ddWtKip,omitzero"`
	// The modified max weight allowable on this surface type for an FAA 2D-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	DdWtKipMod param.Opt[float64] `json:"ddWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an FAA 2D-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	DdWtKipModNote param.Opt[string] `json:"ddWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for an FAA 2D-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	DdWtKn param.Opt[float64] `json:"ddWtKN,omitzero"`
	// WGS-84 latitude of the coordinate representing the end-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	EndLat param.Opt[float64] `json:"endLat,omitzero"`
	// WGS-84 longitude of the coordinate representing the end-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	EndLon param.Opt[float64] `json:"endLon,omitzero"`
	// The unique identifier of the Site record that contains location information
	// about this surface.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Load classification number or pavement rating which ranks aircraft on a scale of
	// 1 to 120.
	Lcn param.Opt[int64] `json:"lcn,omitzero"`
	// The landing distance available for the runway, in feet. Applicable for runway
	// surface types only.
	LdaFt param.Opt[float64] `json:"ldaFt,omitzero"`
	// The landing distance available for the runway, in meters. Applicable for runway
	// surface types only.
	LdaM param.Opt[float64] `json:"ldaM,omitzero"`
	// The length of the surface type, in feet. Applicable for runway and parking
	// surface types.
	LengthFt param.Opt[float64] `json:"lengthFt,omitzero"`
	// The length of the surface type, in meters. Applicable for runway and parking
	// surface types.
	LengthM param.Opt[float64] `json:"lengthM,omitzero"`
	// Flag indicating the surface has lighting.
	Lighting param.Opt[bool] `json:"lighting,omitzero"`
	// Flag indicating the runway has approach lights. Applicable for runway surface
	// types only.
	LightsAprch param.Opt[bool] `json:"lightsAPRCH,omitzero"`
	// Flag indicating the runway has centerline lights. Applicable for runway surface
	// types only.
	LightsCl param.Opt[bool] `json:"lightsCL,omitzero"`
	// Flag indicating the runway has Optical Landing System (OLS) lights. Applicable
	// for runway surface types only.
	LightsOls param.Opt[bool] `json:"lightsOLS,omitzero"`
	// Flag indicating the runway has Precision Approach Path Indicator (PAPI) lights.
	// Applicable for runway surface types only.
	LightsPapi param.Opt[bool] `json:"lightsPAPI,omitzero"`
	// Flag indicating the runway has Runway End Identifier Lights (REIL). Applicable
	// for runway surface types only.
	LightsReil param.Opt[bool] `json:"lightsREIL,omitzero"`
	// Flag indicating the runway has edge lighting. Applicable for runway surface
	// types only.
	LightsRwy param.Opt[bool] `json:"lightsRWY,omitzero"`
	// Flag indicating the runway has Sequential Flashing (SEQFL) lights. Applicable
	// for runway surface types only.
	LightsSeqfl param.Opt[bool] `json:"lightsSEQFL,omitzero"`
	// Flag indicating the runway has Touchdown Zone lights. Applicable for runway
	// surface types only.
	LightsTdzl param.Opt[bool] `json:"lightsTDZL,omitzero"`
	// Flag indicating the runway lighting is unknown. Applicable for runway surface
	// types only.
	LightsUnkn param.Opt[bool] `json:"lightsUNKN,omitzero"`
	// Flag indicating the runway has Visual Approach Slope Indicator (VASI) lights.
	// Applicable for runway surface types only.
	LightsVasi param.Opt[bool] `json:"lightsVASI,omitzero"`
	// The surface material (e.g. Asphalt, Concrete, Dirt).
	Material param.Opt[string] `json:"material,omitzero"`
	// Flag indicating that one or more obstacles or obstructions exist in the
	// proximity of this surface.
	Obstacle param.Opt[bool] `json:"obstacle,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Pavement classification number (PCN) and tire pressure code.
	Pcn param.Opt[string] `json:"pcn,omitzero"`
	// Description of the surface and its dimensions or how it is measured or oriented.
	PointReference param.Opt[string] `json:"pointReference,omitzero"`
	// Flag indicating this is the primary runway. Applicable for runway surface types
	// only.
	Primary param.Opt[bool] `json:"primary,omitzero"`
	// Raw weight bearing capacity value or pavement strength.
	RawWbc param.Opt[string] `json:"rawWBC,omitzero"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilopounds (kip).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKip param.Opt[float64] `json:"sbttWtKip,omitzero"`
	// The modified max weight allowable on this surface type for an SBTT-type (single
	// belly twin tandem) landing gear configuration, in kilopounds (kip).
	SbttWtKipMod param.Opt[float64] `json:"sbttWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an SBTT-type (single
	// belly twin tandem) landing gear configuration. For more information, contact the
	// provider source.
	SbttWtKipModNote param.Opt[string] `json:"sbttWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for an SBTT-type (single belly
	// twin tandem) landing gear configuration, in kilonewtons (kN).Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	SbttWtKn param.Opt[float64] `json:"sbttWtKN,omitzero"`
	// WGS-84 latitude of the coordinate representing the start-point of a surface, in
	// degrees. -90 to 90 degrees (negative values south of equator).
	StartLat param.Opt[float64] `json:"startLat,omitzero"`
	// WGS-84 longitude of the coordinate representing the start-point of a surface, in
	// degrees. -180 to 180 degrees (negative values west of Prime Meridian).
	StartLon param.Opt[float64] `json:"startLon,omitzero"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKip param.Opt[float64] `json:"stWtKip,omitzero"`
	// The modified max weight allowable on this surface type for an ST-type (single
	// tandem) landing gear configuration, in kilopounds (kip).
	StWtKipMod param.Opt[float64] `json:"stWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an ST-type (single
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	StWtKipModNote param.Opt[string] `json:"stWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for an ST-type (single tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	StWtKn param.Opt[float64] `json:"stWtKN,omitzero"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilopounds (kip). Note: The corresponding equivalent
	// field is not converted by the UDL and may or may not be supplied by the
	// provider. The provider/consumer is responsible for all unit conversions.
	SWtKip param.Opt[float64] `json:"sWtKip,omitzero"`
	// The modified max weight allowable on this surface type for an S-type (single)
	// landing gear configuration, in kilopounds (kip).
	SWtKipMod param.Opt[float64] `json:"sWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for an S-type (single)
	// landing gear configuration. For more information, contact the provider source.
	SWtKipModNote param.Opt[string] `json:"sWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for an S-type (single) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	SWtKn param.Opt[float64] `json:"sWtKN,omitzero"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtkip param.Opt[float64] `json:"tdtWtkip,omitzero"`
	// The modified max weight allowable on this surface type for a TDT-type (twin
	// delta tandem) landing gear configuration, in kilopounds (kip).
	TdtWtKipMod param.Opt[float64] `json:"tdtWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TDT-type (twin delta
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TdtWtKipModNote param.Opt[string] `json:"tdtWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a TDT-type (twin delta tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TdtWtKn param.Opt[float64] `json:"tdtWtKN,omitzero"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKip param.Opt[float64] `json:"trtWtKip,omitzero"`
	// The modified max weight allowable on this surface type for a TRT-type (triple
	// tandem) landing gear configuration, in kilopounds (kip).
	TrtWtKipMod param.Opt[float64] `json:"trtWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a TRT-type (triple
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TrtWtKipModNote param.Opt[string] `json:"trtWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a TRT-type (triple tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TrtWtKn param.Opt[float64] `json:"trtWtKN,omitzero"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilopounds (kip).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKip param.Opt[float64] `json:"ttWtKip,omitzero"`
	// The modified max weight allowable on this surface type for a GDSS TT-type (twin
	// tandem) landing gear configuration, in kilopounds (kip).
	TtWtKipMod param.Opt[float64] `json:"ttWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a GDSS TT-type (twin
	// tandem) landing gear configuration. For more information, contact the provider
	// source.
	TtWtKipModNote param.Opt[string] `json:"ttWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a GDSS TT-type (twin tandem)
	// landing gear configuration, in kilonewtons (kN).Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	TtWtKn param.Opt[float64] `json:"ttWtKN,omitzero"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilopounds (kip).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKip param.Opt[float64] `json:"tWtKip,omitzero"`
	// The modified max weight allowable on this surface type for a T-type (twin
	// (dual)) landing gear configuration, in kilopounds (kip).
	TWtKipMod param.Opt[float64] `json:"tWtKipMod,omitzero"`
	// User input rationale for the manual modification of the prescribed value
	// indicating the max weight allowable on this surface for a T-type (twin(dual))
	// landing gear configuration. For more information, contact the provider source.
	TWtKipModNote param.Opt[string] `json:"tWtKipModNote,omitzero"`
	// The max weight allowable on this surface type for a T-type (twin (dual)) landing
	// gear configuration, in kilonewtons (kN).Note: The corresponding equivalent field
	// is not converted by the UDL and may or may not be supplied by the provider. The
	// provider/consumer is responsible for all unit conversions.
	TWtKn param.Opt[float64] `json:"tWtKN,omitzero"`
	// The width of the surface type, in feet.
	WidthFt param.Opt[float64] `json:"widthFt,omitzero"`
	// The width of the surface type, in meters.
	WidthM param.Opt[float64] `json:"widthM,omitzero"`
	paramObj
}

func (r SurfaceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SurfaceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SurfaceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, SIMULATED, or EXERCISE data.
type SurfaceUpdateParamsDataMode string

const (
	SurfaceUpdateParamsDataModeReal      SurfaceUpdateParamsDataMode = "REAL"
	SurfaceUpdateParamsDataModeTest      SurfaceUpdateParamsDataMode = "TEST"
	SurfaceUpdateParamsDataModeSimulated SurfaceUpdateParamsDataMode = "SIMULATED"
	SurfaceUpdateParamsDataModeExercise  SurfaceUpdateParamsDataMode = "EXERCISE"
)

type SurfaceListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SurfaceListParams]'s query parameters as `url.Values`.
func (r SurfaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SurfaceCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SurfaceCountParams]'s query parameters as `url.Values`.
func (r SurfaceCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SurfaceGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SurfaceGetParams]'s query parameters as `url.Values`.
func (r SurfaceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SurfaceTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SurfaceTupleParams]'s query parameters as `url.Values`.
func (r SurfaceTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
