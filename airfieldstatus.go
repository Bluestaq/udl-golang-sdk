// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
)

// AirfieldStatusService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirfieldStatusService] method instead.
type AirfieldStatusService struct {
	Options []option.RequestOption
}

// NewAirfieldStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAirfieldStatusService(opts ...option.RequestOption) (r AirfieldStatusService) {
	r = AirfieldStatusService{}
	r.Options = opts
	return
}

// Service operation to get a single airfield status record by its unique ID passed
// as a path parameter.
func (r *AirfieldStatusService) Get(ctx context.Context, id string, query AirfieldStatusGetParams, opts ...option.RequestOption) (res *AirfieldstatusFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfieldstatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single airfield status record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *AirfieldStatusService) Update(ctx context.Context, id string, body AirfieldStatusUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfieldstatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to delete a Status object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *AirfieldStatusService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfieldstatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
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
func (r *AirfieldStatusService) Tuple(ctx context.Context, query AirfieldStatusTupleParams, opts ...option.RequestOption) (res *[]AirfieldstatusFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airfieldstatus/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AirfieldStatusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldStatusGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirfieldStatusGetParams]'s query parameters as
// `url.Values`.
func (r AirfieldStatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldStatusUpdateParams struct {
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
	DataMode AirfieldStatusUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the Airfield for which this status is referencing.
	IDAirfield string `json:"idAirfield,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate airfield identifier provided by the source.
	AltAirfieldID param.Opt[string] `json:"altAirfieldId,omitzero"`
	// The name of the person who approved the airfield survey review.
	ApprovedBy param.Opt[string] `json:"approvedBy,omitzero"`
	// The date that survey review changes were approved for this airfield, in ISO 8601
	// UTC format with millisecond precision.
	ApprovedDate param.Opt[time.Time] `json:"approvedDate,omitzero" format:"date-time"`
	// The category of aircraft rescue and fire fighting (ARFF) services that are
	// currently available at the airfield. Entries should include the code (FAA or
	// ICAO) and the category.
	ArffCat param.Opt[string] `json:"arffCat,omitzero"`
	// Maximum on ground (MOG) number of high-reach/wide-body cargo aircraft that can
	// be serviced simultaneously based on spacing and manpower at the time of status.
	CargoMog param.Opt[int64] `json:"cargoMOG,omitzero"`
	// Maximum on ground (MOG) number of fleet aircraft that can be serviced
	// simultaneously based on spacing and manpower at the time of status.
	FleetServiceMog param.Opt[int64] `json:"fleetServiceMOG,omitzero"`
	// Maximum on ground (MOG) number of aircraft that can be simultaneously refueled
	// based on spacing and manpower at the time of status.
	FuelMog param.Opt[int64] `json:"fuelMOG,omitzero"`
	// The expected time to receive ground support equipment (e.g. power units, air
	// units, cables, hoses, etc.), in minutes.
	GseTime param.Opt[int64] `json:"gseTime,omitzero"`
	// The level of medical support and capabilities available at the airfield.
	MedCap param.Opt[string] `json:"medCap,omitzero"`
	// Description of the current status of the airfield.
	Message param.Opt[string] `json:"message,omitzero"`
	// Maximum on ground (MOG) number of aircraft that can be simultaneously ground
	// handled for standard maintenance based on spacing and manpower at the time of
	// status.
	MxMog param.Opt[int64] `json:"mxMOG,omitzero"`
	// Maximum on ground (MOG) number of parking narrow-body aircraft based on spacing
	// and manpower at the time of status.
	NarrowParkingMog param.Opt[int64] `json:"narrowParkingMOG,omitzero"`
	// Maximum on ground (MOG) number of working narrow-body aircraft based on spacing
	// and manpower at the time of status.
	NarrowWorkingMog param.Opt[int64] `json:"narrowWorkingMOG,omitzero"`
	// The number of aircraft that are currently on ground (COG) at the airfield.
	NumCog param.Opt[int64] `json:"numCOG,omitzero"`
	// Maximum on ground (MOG) number of aircraft due to items not directly related to
	// the airfield infrastructure or aircraft servicing capability based on spacing
	// and manpower at the time of status.
	OperatingMog param.Opt[int64] `json:"operatingMOG,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Maximum on ground (MOG) number of high-reach/wide-body passenger aircraft that
	// can be serviced simultaneously based on spacing and manpower at the time of
	// status.
	PassengerServiceMog param.Opt[int64] `json:"passengerServiceMOG,omitzero"`
	// The primary frequency which the airfield is currently operating, in megahertz.
	PriFreq param.Opt[float64] `json:"priFreq,omitzero"`
	// The number or ID of primary runway at the airfield.
	PriRwyNum param.Opt[string] `json:"priRwyNum,omitzero"`
	// The name of the person who reviewed the airfield survey.
	ReviewedBy param.Opt[string] `json:"reviewedBy,omitzero"`
	// The date the airfield survey was reviewed, in ISO 8601 UTC format with
	// millisecond precision.
	ReviewedDate param.Opt[time.Time] `json:"reviewedDate,omitzero" format:"date-time"`
	// The primary runway condition reading value used for determining runway braking
	// action, from 0 to 26. A value of 0 indicates braking action is poor or
	// non-existent, where a value of 26 indicates braking action is good.
	RwyCondReading param.Opt[int64] `json:"rwyCondReading,omitzero"`
	// The primary runway friction factor which is dependent on the surface friction
	// between the tires of the aircraft and the runway surface, from 0 to 100. A lower
	// number indicates less friction and less braking response.
	RwyFrictionFactor param.Opt[int64] `json:"rwyFrictionFactor,omitzero"`
	// The date the airfield survey was performed, in ISO 8601 UTC format with
	// millisecond precision.
	SurveyDate param.Opt[time.Time] `json:"surveyDate,omitzero" format:"date-time"`
	// Maximum on ground (MOG) number of parking wide-body aircraft based on spacing
	// and manpower at the time of status. Additional information about this field as
	// it pertains to specific aircraft type may be available in an associated
	// SiteOperations record.
	WideParkingMog param.Opt[int64] `json:"wideParkingMOG,omitzero"`
	// Maximum on ground (MOG) number of working wide-body aircraft based on spacing
	// and manpower at the time of status. Additional information about this field as
	// it pertains to specific aircraft type may be available in an associated
	// SiteOperations record.
	WideWorkingMog param.Opt[int64] `json:"wideWorkingMOG,omitzero"`
	// Array of quantities for each fuel type at the airfield, in kilograms. The values
	// in this array must correspond to the position index in fuelTypes. This array
	// must be the same length as fuelTypes.
	FuelQtys []float64 `json:"fuelQtys,omitzero"`
	// Array of fuel types available at the airfield. This array must be the same
	// length as fuelQtys.
	FuelTypes []string `json:"fuelTypes,omitzero"`
	// Array of quantities for each material handling equipment types at the airfield.
	// The values in this array must correspond to the position index in mheTypes. This
	// array must be the same length as mheTypes.
	MheQtys []int64 `json:"mheQtys,omitzero"`
	// Array of material handling equipment types at the airfield. This array must be
	// the same length as mheQtys.
	MheTypes []string `json:"mheTypes,omitzero"`
	// Array of markings currently on the primary runway.
	RwyMarkings []string `json:"rwyMarkings,omitzero"`
	// Array of slot types that an airfield requires a particular aircraft provide in
	// order to consume a slot at this location.
	SlotTypesReq []string `json:"slotTypesReq,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldStatusUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r AirfieldStatusUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AirfieldStatusUpdateParams
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
type AirfieldStatusUpdateParamsDataMode string

const (
	AirfieldStatusUpdateParamsDataModeReal      AirfieldStatusUpdateParamsDataMode = "REAL"
	AirfieldStatusUpdateParamsDataModeTest      AirfieldStatusUpdateParamsDataMode = "TEST"
	AirfieldStatusUpdateParamsDataModeSimulated AirfieldStatusUpdateParamsDataMode = "SIMULATED"
	AirfieldStatusUpdateParamsDataModeExercise  AirfieldStatusUpdateParamsDataMode = "EXERCISE"
)

type AirfieldStatusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldStatusTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirfieldStatusTupleParams]'s query parameters as
// `url.Values`.
func (r AirfieldStatusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
