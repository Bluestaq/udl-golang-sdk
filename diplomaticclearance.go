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
)

// DiplomaticClearanceService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiplomaticClearanceService] method instead.
type DiplomaticClearanceService struct {
	Options []option.RequestOption
	History DiplomaticClearanceHistoryService
	Country DiplomaticClearanceCountryService
}

// NewDiplomaticClearanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDiplomaticClearanceService(opts ...option.RequestOption) (r DiplomaticClearanceService) {
	r = DiplomaticClearanceService{}
	r.Options = opts
	r.History = NewDiplomaticClearanceHistoryService(opts...)
	r.Country = NewDiplomaticClearanceCountryService(opts...)
	return
}

// Service operation to take a single diplomatic clearance record as a POST body
// and ingest into the database. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *DiplomaticClearanceService) New(ctx context.Context, body DiplomaticClearanceNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/diplomaticclearance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single diplomatic clearance record by its unique ID
// passed as a path parameter.
func (r *DiplomaticClearanceService) Get(ctx context.Context, id string, query DiplomaticClearanceGetParams, opts ...option.RequestOption) (res *DiplomaticclearanceFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/diplomaticclearance/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single diplomatic clearance record. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *DiplomaticClearanceService) Update(ctx context.Context, id string, body DiplomaticClearanceUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/diplomaticclearance/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *DiplomaticClearanceService) List(ctx context.Context, query DiplomaticClearanceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[DiplomaticclearanceAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/diplomaticclearance"
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
func (r *DiplomaticClearanceService) ListAutoPaging(ctx context.Context, query DiplomaticClearanceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[DiplomaticclearanceAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a diplomatic clearance record specified by the
// passed ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *DiplomaticClearanceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/diplomaticclearance/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *DiplomaticClearanceService) Count(ctx context.Context, query DiplomaticClearanceCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/diplomaticclearance/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// diplomaticclearance records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *DiplomaticClearanceService) NewBulk(ctx context.Context, body DiplomaticClearanceNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/diplomaticclearance/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *DiplomaticClearanceService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/diplomaticclearance/queryhelp"
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
func (r *DiplomaticClearanceService) Tuple(ctx context.Context, query DiplomaticClearanceTupleParams, opts ...option.RequestOption) (res *[]DiplomaticclearanceFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/diplomaticclearance/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DiplomaticClearanceNewParams struct {
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
	DataMode DiplomaticClearanceNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The First Departure Date (FDD) the mission is scheduled for departure, in ISO
	// 8601 UTC format with millisecond precision.
	FirstDepDate time.Time `json:"firstDepDate,required" format:"date-time"`
	// Unique identifier of the Mission associated with this diplomatic clearance
	// record.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Aircraft and Personnel Automated Clearance System (APACS) system identifier
	// used to process and approve this clearance request.
	ApacsID param.Opt[string] `json:"apacsId,omitzero"`
	// Identifier of the Diplomatic Clearance Worksheet used to coordinate aircraft
	// clearance requests.
	DipWorksheetName param.Opt[string] `json:"dipWorksheetName,omitzero"`
	// Suspense date for the diplomatic clearance worksheet to be worked, in ISO 8601
	// UTC format with millisecond precision.
	DocDeadline param.Opt[time.Time] `json:"docDeadline,omitzero" format:"date-time"`
	// Optional diplomatic clearance worksheet ID from external systems. This field has
	// no meaning within UDL and is provided as a convenience for systems that require
	// tracking of an internal system generated ID.
	ExternalWorksheetID param.Opt[string] `json:"externalWorksheetId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Collection of diplomatic clearance details.
	DiplomaticClearanceDetails []DiplomaticClearanceNewParamsDiplomaticClearanceDetail `json:"diplomaticClearanceDetails,omitzero"`
	// Collection of diplomatic clearance remarks.
	DiplomaticClearanceRemarks []DiplomaticClearanceNewParamsDiplomaticClearanceRemark `json:"diplomaticClearanceRemarks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r DiplomaticClearanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceNewParams
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
type DiplomaticClearanceNewParamsDataMode string

const (
	DiplomaticClearanceNewParamsDataModeReal      DiplomaticClearanceNewParamsDataMode = "REAL"
	DiplomaticClearanceNewParamsDataModeTest      DiplomaticClearanceNewParamsDataMode = "TEST"
	DiplomaticClearanceNewParamsDataModeSimulated DiplomaticClearanceNewParamsDataMode = "SIMULATED"
	DiplomaticClearanceNewParamsDataModeExercise  DiplomaticClearanceNewParamsDataMode = "EXERCISE"
)

// Collection of diplomatic clearance details.
type DiplomaticClearanceNewParamsDiplomaticClearanceDetail struct {
	// The type of action the aircraft can take with this diplomatic clearance (e.g. O
	// for Overfly, L for Land, etc.).
	Action param.Opt[string] `json:"action,omitzero"`
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS. This field will be set
	// to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Identifier of this diplomatic clearance issued by the host country.
	ClearanceID param.Opt[string] `json:"clearanceId,omitzero"`
	// Remarks concerning this diplomatic clearance.
	ClearanceRemark param.Opt[string] `json:"clearanceRemark,omitzero"`
	// The call sign of the sortie cleared with this diplomatic clearance.
	ClearedCallSign param.Opt[string] `json:"clearedCallSign,omitzero"`
	// The DoD Standard Country Code designator for the country issuing the diplomatic
	// clearance. This field will be set to "OTHR" if the source value does not match a
	// UDL Country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Name of the country issuing this diplomatic clearance.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// Earliest time the aircraft may enter the country, in ISO 8601 UTC format with
	// millisecond precision.
	EntryNet param.Opt[time.Time] `json:"entryNET,omitzero" format:"date-time"`
	// The navigation point name where the aircraft must enter the country.
	EntryPoint param.Opt[string] `json:"entryPoint,omitzero"`
	// Latest time the aircraft may exit the country, in ISO 8601 UTC format with
	// millisecond precision.
	ExitNlt param.Opt[time.Time] `json:"exitNLT,omitzero" format:"date-time"`
	// The navigation point name where the aircraft must exit the country.
	ExitPoint param.Opt[string] `json:"exitPoint,omitzero"`
	// Optional clearance ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalClearanceID param.Opt[string] `json:"externalClearanceId,omitzero"`
	// Unique identifier of the Aircraft Sortie associated with this diplomatic
	// clearance record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile param.Opt[string] `json:"profile,omitzero"`
	// Flag indicating whether the clearance request requires ICAO specific
	// information.
	ReqIcao param.Opt[bool] `json:"reqICAO,omitzero"`
	// Flag indicating whether entry/exit points are required for clearances.
	ReqPoint param.Opt[bool] `json:"reqPoint,omitzero"`
	// The 1801 fileable route of flight string associated with this diplomatic
	// clearance. The route of flight string contains route designators, significant
	// points, change of speed/altitude, change of flight rules, and cruise climbs.
	RouteString param.Opt[string] `json:"routeString,omitzero"`
	// The placement of this diplomatic clearance within a sequence of clearances used
	// on a sortie. For example, a sequence value of 3 means that it is the third
	// diplomatic clearance the aircraft will use.
	SequenceNum param.Opt[int64] `json:"sequenceNum,omitzero"`
	// Indicates the current status of the diplomatic clearance request.
	Status param.Opt[string] `json:"status,omitzero"`
	// Description of when this diplomatic clearance is valid.
	ValidDesc param.Opt[string] `json:"validDesc,omitzero"`
	// The end time of the validity of this diplomatic clearance, in ISO 8601 UTC
	// format with millisecond precision.
	ValidEndTime param.Opt[time.Time] `json:"validEndTime,omitzero" format:"date-time"`
	// The start time of the validity of this diplomatic clearance, in ISO 8601 UTC
	// format with millisecond precision.
	ValidStartTime param.Opt[time.Time] `json:"validStartTime,omitzero" format:"date-time"`
	// Remarks concerning the valid diplomatic clearance window.
	WindowRemark param.Opt[string] `json:"windowRemark,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceNewParamsDiplomaticClearanceDetail) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r DiplomaticClearanceNewParamsDiplomaticClearanceDetail) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceNewParamsDiplomaticClearanceDetail
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of diplomatic clearance remarks.
type DiplomaticClearanceNewParamsDiplomaticClearanceRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Global Decision Support System (GDSS) remark identifier.
	GdssRemarkID param.Opt[string] `json:"gdssRemarkId,omitzero"`
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// User who published the remark.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceNewParamsDiplomaticClearanceRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r DiplomaticClearanceNewParamsDiplomaticClearanceRemark) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceNewParamsDiplomaticClearanceRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

type DiplomaticClearanceGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [DiplomaticClearanceGetParams]'s query parameters as
// `url.Values`.
func (r DiplomaticClearanceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DiplomaticClearanceUpdateParams struct {
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
	DataMode DiplomaticClearanceUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The First Departure Date (FDD) the mission is scheduled for departure, in ISO
	// 8601 UTC format with millisecond precision.
	FirstDepDate time.Time `json:"firstDepDate,required" format:"date-time"`
	// Unique identifier of the Mission associated with this diplomatic clearance
	// record.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Aircraft and Personnel Automated Clearance System (APACS) system identifier
	// used to process and approve this clearance request.
	ApacsID param.Opt[string] `json:"apacsId,omitzero"`
	// Identifier of the Diplomatic Clearance Worksheet used to coordinate aircraft
	// clearance requests.
	DipWorksheetName param.Opt[string] `json:"dipWorksheetName,omitzero"`
	// Suspense date for the diplomatic clearance worksheet to be worked, in ISO 8601
	// UTC format with millisecond precision.
	DocDeadline param.Opt[time.Time] `json:"docDeadline,omitzero" format:"date-time"`
	// Optional diplomatic clearance worksheet ID from external systems. This field has
	// no meaning within UDL and is provided as a convenience for systems that require
	// tracking of an internal system generated ID.
	ExternalWorksheetID param.Opt[string] `json:"externalWorksheetId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Collection of diplomatic clearance details.
	DiplomaticClearanceDetails []DiplomaticClearanceUpdateParamsDiplomaticClearanceDetail `json:"diplomaticClearanceDetails,omitzero"`
	// Collection of diplomatic clearance remarks.
	DiplomaticClearanceRemarks []DiplomaticClearanceUpdateParamsDiplomaticClearanceRemark `json:"diplomaticClearanceRemarks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r DiplomaticClearanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceUpdateParams
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
type DiplomaticClearanceUpdateParamsDataMode string

const (
	DiplomaticClearanceUpdateParamsDataModeReal      DiplomaticClearanceUpdateParamsDataMode = "REAL"
	DiplomaticClearanceUpdateParamsDataModeTest      DiplomaticClearanceUpdateParamsDataMode = "TEST"
	DiplomaticClearanceUpdateParamsDataModeSimulated DiplomaticClearanceUpdateParamsDataMode = "SIMULATED"
	DiplomaticClearanceUpdateParamsDataModeExercise  DiplomaticClearanceUpdateParamsDataMode = "EXERCISE"
)

// Collection of diplomatic clearance details.
type DiplomaticClearanceUpdateParamsDiplomaticClearanceDetail struct {
	// The type of action the aircraft can take with this diplomatic clearance (e.g. O
	// for Overfly, L for Land, etc.).
	Action param.Opt[string] `json:"action,omitzero"`
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS. This field will be set
	// to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Identifier of this diplomatic clearance issued by the host country.
	ClearanceID param.Opt[string] `json:"clearanceId,omitzero"`
	// Remarks concerning this diplomatic clearance.
	ClearanceRemark param.Opt[string] `json:"clearanceRemark,omitzero"`
	// The call sign of the sortie cleared with this diplomatic clearance.
	ClearedCallSign param.Opt[string] `json:"clearedCallSign,omitzero"`
	// The DoD Standard Country Code designator for the country issuing the diplomatic
	// clearance. This field will be set to "OTHR" if the source value does not match a
	// UDL Country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Name of the country issuing this diplomatic clearance.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// Earliest time the aircraft may enter the country, in ISO 8601 UTC format with
	// millisecond precision.
	EntryNet param.Opt[time.Time] `json:"entryNET,omitzero" format:"date-time"`
	// The navigation point name where the aircraft must enter the country.
	EntryPoint param.Opt[string] `json:"entryPoint,omitzero"`
	// Latest time the aircraft may exit the country, in ISO 8601 UTC format with
	// millisecond precision.
	ExitNlt param.Opt[time.Time] `json:"exitNLT,omitzero" format:"date-time"`
	// The navigation point name where the aircraft must exit the country.
	ExitPoint param.Opt[string] `json:"exitPoint,omitzero"`
	// Optional clearance ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalClearanceID param.Opt[string] `json:"externalClearanceId,omitzero"`
	// Unique identifier of the Aircraft Sortie associated with this diplomatic
	// clearance record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile param.Opt[string] `json:"profile,omitzero"`
	// Flag indicating whether the clearance request requires ICAO specific
	// information.
	ReqIcao param.Opt[bool] `json:"reqICAO,omitzero"`
	// Flag indicating whether entry/exit points are required for clearances.
	ReqPoint param.Opt[bool] `json:"reqPoint,omitzero"`
	// The 1801 fileable route of flight string associated with this diplomatic
	// clearance. The route of flight string contains route designators, significant
	// points, change of speed/altitude, change of flight rules, and cruise climbs.
	RouteString param.Opt[string] `json:"routeString,omitzero"`
	// The placement of this diplomatic clearance within a sequence of clearances used
	// on a sortie. For example, a sequence value of 3 means that it is the third
	// diplomatic clearance the aircraft will use.
	SequenceNum param.Opt[int64] `json:"sequenceNum,omitzero"`
	// Indicates the current status of the diplomatic clearance request.
	Status param.Opt[string] `json:"status,omitzero"`
	// Description of when this diplomatic clearance is valid.
	ValidDesc param.Opt[string] `json:"validDesc,omitzero"`
	// The end time of the validity of this diplomatic clearance, in ISO 8601 UTC
	// format with millisecond precision.
	ValidEndTime param.Opt[time.Time] `json:"validEndTime,omitzero" format:"date-time"`
	// The start time of the validity of this diplomatic clearance, in ISO 8601 UTC
	// format with millisecond precision.
	ValidStartTime param.Opt[time.Time] `json:"validStartTime,omitzero" format:"date-time"`
	// Remarks concerning the valid diplomatic clearance window.
	WindowRemark param.Opt[string] `json:"windowRemark,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceUpdateParamsDiplomaticClearanceDetail) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r DiplomaticClearanceUpdateParamsDiplomaticClearanceDetail) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceUpdateParamsDiplomaticClearanceDetail
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of diplomatic clearance remarks.
type DiplomaticClearanceUpdateParamsDiplomaticClearanceRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Global Decision Support System (GDSS) remark identifier.
	GdssRemarkID param.Opt[string] `json:"gdssRemarkId,omitzero"`
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// User who published the remark.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceUpdateParamsDiplomaticClearanceRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r DiplomaticClearanceUpdateParamsDiplomaticClearanceRemark) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceUpdateParamsDiplomaticClearanceRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

type DiplomaticClearanceListParams struct {
	// The First Departure Date (FDD) the mission is scheduled for departure, in ISO
	// 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	FirstDepDate time.Time        `query:"firstDepDate,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [DiplomaticClearanceListParams]'s query parameters as
// `url.Values`.
func (r DiplomaticClearanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DiplomaticClearanceCountParams struct {
	// The First Departure Date (FDD) the mission is scheduled for departure, in ISO
	// 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	FirstDepDate time.Time        `query:"firstDepDate,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [DiplomaticClearanceCountParams]'s query parameters as
// `url.Values`.
func (r DiplomaticClearanceCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DiplomaticClearanceNewBulkParams struct {
	Body []DiplomaticClearanceNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r DiplomaticClearanceNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// A diplomatic clearance is an authorization for an aircraft to traverse or land
// within a specified country.
//
// The properties ClassificationMarking, DataMode, FirstDepDate, IDMission, Source
// are required.
type DiplomaticClearanceNewBulkParamsBody struct {
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
	// The First Departure Date (FDD) the mission is scheduled for departure, in ISO
	// 8601 UTC format with millisecond precision.
	FirstDepDate time.Time `json:"firstDepDate,required" format:"date-time"`
	// Unique identifier of the Mission associated with this diplomatic clearance
	// record.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Aircraft and Personnel Automated Clearance System (APACS) system identifier
	// used to process and approve this clearance request.
	ApacsID param.Opt[string] `json:"apacsId,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Identifier of the Diplomatic Clearance Worksheet used to coordinate aircraft
	// clearance requests.
	DipWorksheetName param.Opt[string] `json:"dipWorksheetName,omitzero"`
	// Suspense date for the diplomatic clearance worksheet to be worked, in ISO 8601
	// UTC format with millisecond precision.
	DocDeadline param.Opt[time.Time] `json:"docDeadline,omitzero" format:"date-time"`
	// Optional diplomatic clearance worksheet ID from external systems. This field has
	// no meaning within UDL and is provided as a convenience for systems that require
	// tracking of an internal system generated ID.
	ExternalWorksheetID param.Opt[string] `json:"externalWorksheetId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Collection of diplomatic clearance details.
	DiplomaticClearanceDetails []DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceDetail `json:"diplomaticClearanceDetails,omitzero"`
	// Collection of diplomatic clearance remarks.
	DiplomaticClearanceRemarks []DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceRemark `json:"diplomaticClearanceRemarks,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceNewBulkParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r DiplomaticClearanceNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[DiplomaticClearanceNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of diplomatic clearance details.
type DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceDetail struct {
	// The type of action the aircraft can take with this diplomatic clearance (e.g. O
	// for Overfly, L for Land, etc.).
	Action param.Opt[string] `json:"action,omitzero"`
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS. This field will be set
	// to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Identifier of this diplomatic clearance issued by the host country.
	ClearanceID param.Opt[string] `json:"clearanceId,omitzero"`
	// Remarks concerning this diplomatic clearance.
	ClearanceRemark param.Opt[string] `json:"clearanceRemark,omitzero"`
	// The call sign of the sortie cleared with this diplomatic clearance.
	ClearedCallSign param.Opt[string] `json:"clearedCallSign,omitzero"`
	// The DoD Standard Country Code designator for the country issuing the diplomatic
	// clearance. This field will be set to "OTHR" if the source value does not match a
	// UDL Country code value (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Name of the country issuing this diplomatic clearance.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// Earliest time the aircraft may enter the country, in ISO 8601 UTC format with
	// millisecond precision.
	EntryNet param.Opt[time.Time] `json:"entryNET,omitzero" format:"date-time"`
	// The navigation point name where the aircraft must enter the country.
	EntryPoint param.Opt[string] `json:"entryPoint,omitzero"`
	// Latest time the aircraft may exit the country, in ISO 8601 UTC format with
	// millisecond precision.
	ExitNlt param.Opt[time.Time] `json:"exitNLT,omitzero" format:"date-time"`
	// The navigation point name where the aircraft must exit the country.
	ExitPoint param.Opt[string] `json:"exitPoint,omitzero"`
	// Optional clearance ID from external systems. This field has no meaning within
	// UDL and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalClearanceID param.Opt[string] `json:"externalClearanceId,omitzero"`
	// Unique identifier of the Aircraft Sortie associated with this diplomatic
	// clearance record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// Identifies the Itinerary point of a sortie where an air event occurs.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile param.Opt[string] `json:"profile,omitzero"`
	// Flag indicating whether the clearance request requires ICAO specific
	// information.
	ReqIcao param.Opt[bool] `json:"reqICAO,omitzero"`
	// Flag indicating whether entry/exit points are required for clearances.
	ReqPoint param.Opt[bool] `json:"reqPoint,omitzero"`
	// The 1801 fileable route of flight string associated with this diplomatic
	// clearance. The route of flight string contains route designators, significant
	// points, change of speed/altitude, change of flight rules, and cruise climbs.
	RouteString param.Opt[string] `json:"routeString,omitzero"`
	// The placement of this diplomatic clearance within a sequence of clearances used
	// on a sortie. For example, a sequence value of 3 means that it is the third
	// diplomatic clearance the aircraft will use.
	SequenceNum param.Opt[int64] `json:"sequenceNum,omitzero"`
	// Indicates the current status of the diplomatic clearance request.
	Status param.Opt[string] `json:"status,omitzero"`
	// Description of when this diplomatic clearance is valid.
	ValidDesc param.Opt[string] `json:"validDesc,omitzero"`
	// The end time of the validity of this diplomatic clearance, in ISO 8601 UTC
	// format with millisecond precision.
	ValidEndTime param.Opt[time.Time] `json:"validEndTime,omitzero" format:"date-time"`
	// The start time of the validity of this diplomatic clearance, in ISO 8601 UTC
	// format with millisecond precision.
	ValidStartTime param.Opt[time.Time] `json:"validStartTime,omitzero" format:"date-time"`
	// Remarks concerning the valid diplomatic clearance window.
	WindowRemark param.Opt[string] `json:"windowRemark,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceDetail) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceDetail) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceDetail
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of diplomatic clearance remarks.
type DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceRemark struct {
	// Date the remark was published, in ISO 8601 UTC format, with millisecond
	// precision.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Global Decision Support System (GDSS) remark identifier.
	GdssRemarkID param.Opt[string] `json:"gdssRemarkId,omitzero"`
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// User who published the remark.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceRemark) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceRemark) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceRemark
	return param.MarshalObject(r, (*shadow)(&r))
}

type DiplomaticClearanceTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The First Departure Date (FDD) the mission is scheduled for departure, in ISO
	// 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	FirstDepDate time.Time        `query:"firstDepDate,required" format:"date-time" json:"-"`
	FirstResult  param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults   param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DiplomaticClearanceTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [DiplomaticClearanceTupleParams]'s query parameters as
// `url.Values`.
func (r DiplomaticClearanceTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
