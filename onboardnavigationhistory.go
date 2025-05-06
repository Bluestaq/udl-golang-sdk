// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// OnboardnavigationHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnboardnavigationHistoryService] method instead.
type OnboardnavigationHistoryService struct {
	Options []option.RequestOption
}

// NewOnboardnavigationHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOnboardnavigationHistoryService(opts ...option.RequestOption) (r OnboardnavigationHistoryService) {
	r = OnboardnavigationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnboardnavigationHistoryService) List(ctx context.Context, query OnboardnavigationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnboardnavigationFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onboardnavigation/history"
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

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnboardnavigationHistoryService) ListAutoPaging(ctx context.Context, query OnboardnavigationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnboardnavigationFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnboardnavigationHistoryService) Aodr(ctx context.Context, query OnboardnavigationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onboardnavigation/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OnboardnavigationHistoryService) Count(ctx context.Context, query OnboardnavigationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/onboardnavigation/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// These services provide spacecraft positional data derived from on-board
// navigational sensors. Navigational points are provided in kilometers and in a
// user specified reference frame, with ECI J2K being preferred.
type OnboardnavigationFull struct {
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
	DataMode OnboardnavigationFullDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Start time of the sensor data, in ISO 8601 UTC format.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Arrays of spacecraft delta position (X, Y, Z), in km, in the specified
	// referenceFrame, between the onboard state and the last accepted sensor position.
	DeltaPos [][]float64 `json:"deltaPos"`
	// End time of the sensor data, in ISO 8601 UTC format.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Unique identifier of the parent EphemerisSet, if this data is correlated with an
	// Ephemeris. If multiple nav sensor records are required, this ID may be
	// associated with each of those records if each is synced to the ephemeris points.
	EsID string `json:"esId"`
	// Unique identifier of the primary satellite on-orbit object.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the last onboard state vector.
	IDStateVector string `json:"idStateVector"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard magnetometer, at each epoch.
	Mag [][]float64 `json:"mag"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the data source to indicate the target object of
	// this record. This may be an internal identifier and not necessarily map to a
	// valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// The reference frame in which the sensor derived positions are provided. If the
	// referenceFrame is null it is assumed to be J2000. Note that this frame is
	// assumed to apply to all of the sensor data in this record. If onboard sensors
	// process positional estimates in different frames then separate records should be
	// generated.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame OnboardnavigationFullReferenceFrame `json:"referenceFrame"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// The last load time of the current star catalog onboard this spacecraft.
	StarCatLoadTime time.Time `json:"starCatLoadTime" format:"date-time"`
	// The name or identifier the star catalog in use.
	StarCatName string `json:"starCatName"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard star tracker, at each epoch.
	StarTracker [][]float64 `json:"starTracker"`
	// Arrays of spacecraft position (X, Y, Z), in km, in the specified referenceFrame,
	// based on the onboard sun sensor, at each epoch.
	SunSensor [][]float64 `json:"sunSensor"`
	// Array of epochs of the observations, in ISO 8601 UTC format. The epochs are
	// assumed to correspond to all sensor data in this record. If sensors do not share
	// a common epoch then separate records should be generated.
	Ts []time.Time `json:"ts" format:"date-time"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Source                resp.Field
		StartTime             resp.Field
		ID                    resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DeltaPos              resp.Field
		EndTime               resp.Field
		EsID                  resp.Field
		IDOnOrbit             resp.Field
		IDStateVector         resp.Field
		Mag                   resp.Field
		OnOrbit               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		OrigObjectID          resp.Field
		ReferenceFrame        resp.Field
		SatNo                 resp.Field
		StarCatLoadTime       resp.Field
		StarCatName           resp.Field
		StarTracker           resp.Field
		SunSensor             resp.Field
		Ts                    resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnboardnavigationFull) RawJSON() string { return r.JSON.raw }
func (r *OnboardnavigationFull) UnmarshalJSON(data []byte) error {
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
type OnboardnavigationFullDataMode string

const (
	OnboardnavigationFullDataModeReal      OnboardnavigationFullDataMode = "REAL"
	OnboardnavigationFullDataModeTest      OnboardnavigationFullDataMode = "TEST"
	OnboardnavigationFullDataModeSimulated OnboardnavigationFullDataMode = "SIMULATED"
	OnboardnavigationFullDataModeExercise  OnboardnavigationFullDataMode = "EXERCISE"
)

// The reference frame in which the sensor derived positions are provided. If the
// referenceFrame is null it is assumed to be J2000. Note that this frame is
// assumed to apply to all of the sensor data in this record. If onboard sensors
// process positional estimates in different frames then separate records should be
// generated.
type OnboardnavigationFullReferenceFrame string

const (
	OnboardnavigationFullReferenceFrameJ2000   OnboardnavigationFullReferenceFrame = "J2000"
	OnboardnavigationFullReferenceFrameEfgTdr  OnboardnavigationFullReferenceFrame = "EFG/TDR"
	OnboardnavigationFullReferenceFrameEcrEcef OnboardnavigationFullReferenceFrame = "ECR/ECEF"
	OnboardnavigationFullReferenceFrameTeme    OnboardnavigationFullReferenceFrame = "TEME"
	OnboardnavigationFullReferenceFrameItrf    OnboardnavigationFullReferenceFrame = "ITRF"
	OnboardnavigationFullReferenceFrameGcrf    OnboardnavigationFullReferenceFrame = "GCRF"
)

type OnboardnavigationHistoryListParams struct {
	// Start time of the sensor data, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime time.Time `query:"startTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnboardnavigationHistoryListParams]'s query parameters as
// `url.Values`.
func (r OnboardnavigationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnboardnavigationHistoryAodrParams struct {
	// Start time of the sensor data, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime time.Time `query:"startTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// optional, notification method for the created file link. When omitted, EMAIL is
	// assumed. Current valid values are: EMAIL, SMS.
	Notification param.Opt[string] `query:"notification,omitzero" json:"-"`
	// optional, field delimiter when the created file is not JSON. Must be a single
	// character chosen from this set: (',', ';', ':', '|'). When omitted, "," is used.
	// It is strongly encouraged that your field delimiter be a character unlikely to
	// occur within the data.
	OutputDelimiter param.Opt[string] `query:"outputDelimiter,omitzero" json:"-"`
	// optional, output format for the file. When omitted, JSON is assumed. Current
	// valid values are: JSON and CSV.
	OutputFormat param.Opt[string] `query:"outputFormat,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnboardnavigationHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r OnboardnavigationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnboardnavigationHistoryCountParams struct {
	// Start time of the sensor data, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnboardnavigationHistoryCountParams]'s query parameters as
// `url.Values`.
func (r OnboardnavigationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
