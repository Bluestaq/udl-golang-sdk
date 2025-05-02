// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
)

// AirfieldslotService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirfieldslotService] method instead.
type AirfieldslotService struct {
	Options []option.RequestOption
}

// NewAirfieldslotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAirfieldslotService(opts ...option.RequestOption) (r AirfieldslotService) {
	r = AirfieldslotService{}
	r.Options = opts
	return
}

// Service operation to get a single airfieldslot record by its unique ID passed as
// a path parameter.
func (r *AirfieldslotService) Get(ctx context.Context, id string, query AirfieldslotGetParams, opts ...option.RequestOption) (res *AirfieldslotFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfieldslot/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single airfieldslot record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *AirfieldslotService) Update(ctx context.Context, id string, body AirfieldslotUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfieldslot/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to delete an airfieldslot record specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *AirfieldslotService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfieldslot/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirfieldslotService) Count(ctx context.Context, query AirfieldslotCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airfieldslot/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AirfieldslotService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airfieldslot/queryhelp"
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
func (r *AirfieldslotService) Tuple(ctx context.Context, query AirfieldslotTupleParams, opts ...option.RequestOption) (res *[]AirfieldslotFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airfieldslot/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AirfieldslotGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldslotGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirfieldslotGetParams]'s query parameters as `url.Values`.
func (r AirfieldslotGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldslotUpdateParams struct {
	// The name of the airfield where this slot is located.
	AirfieldName string `json:"airfieldName,required"`
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
	DataMode AirfieldslotUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Name of this slot.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate airfield identifier provided by the source.
	AltAirfieldID param.Opt[string] `json:"altAirfieldId,omitzero"`
	// Number of aircraft that can fit in this slot at the same time.
	Capacity param.Opt[int64] `json:"capacity,omitzero"`
	// Latest zulu time this slot is available based on daily standard hours. Not
	// applicable to slots with type PARKING. Abnormal hours, such as holidays, should
	// be marked via the AirfieldSlotConsumption schema.
	EndTime param.Opt[string] `json:"endTime,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the airfield.
	Icao param.Opt[string] `json:"icao,omitzero"`
	// Unique identifier of the Airfield for which this slot information applies.
	IDAirfield param.Opt[string] `json:"idAirfield,omitzero"`
	// Minimum time that must elapse between different aircraft leaving and entering
	// this slot, in minutes.
	MinSeparation param.Opt[int64] `json:"minSeparation,omitzero"`
	// Optional notes/comments for this airfield slot.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Zulu time this slot is first available based on daily standard hours. Not
	// applicable to slots with type PARKING. Abnormal hours, such as holidays, should
	// be marked via the AirfieldSlotConsumption schema.
	StartTime param.Opt[string] `json:"startTime,omitzero"`
	// Largest category of aircraft supported in this slot (WIDE, NARROW, HELO, ALL,
	// OTHER).
	//
	// Any of "WIDE", "NARROW", "HELO", "ALL", "OTHER".
	AcSlotCat AirfieldslotUpdateParamsAcSlotCat `json:"acSlotCat,omitzero"`
	// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
	// OTHER).
	//
	// Any of "WORKING", "PARKING", "TAKEOFF", "LANDING", "OTHER".
	Type AirfieldslotUpdateParamsType `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldslotUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r AirfieldslotUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AirfieldslotUpdateParams
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
type AirfieldslotUpdateParamsDataMode string

const (
	AirfieldslotUpdateParamsDataModeReal      AirfieldslotUpdateParamsDataMode = "REAL"
	AirfieldslotUpdateParamsDataModeTest      AirfieldslotUpdateParamsDataMode = "TEST"
	AirfieldslotUpdateParamsDataModeSimulated AirfieldslotUpdateParamsDataMode = "SIMULATED"
	AirfieldslotUpdateParamsDataModeExercise  AirfieldslotUpdateParamsDataMode = "EXERCISE"
)

// Largest category of aircraft supported in this slot (WIDE, NARROW, HELO, ALL,
// OTHER).
type AirfieldslotUpdateParamsAcSlotCat string

const (
	AirfieldslotUpdateParamsAcSlotCatWide   AirfieldslotUpdateParamsAcSlotCat = "WIDE"
	AirfieldslotUpdateParamsAcSlotCatNarrow AirfieldslotUpdateParamsAcSlotCat = "NARROW"
	AirfieldslotUpdateParamsAcSlotCatHelo   AirfieldslotUpdateParamsAcSlotCat = "HELO"
	AirfieldslotUpdateParamsAcSlotCatAll    AirfieldslotUpdateParamsAcSlotCat = "ALL"
	AirfieldslotUpdateParamsAcSlotCatOther  AirfieldslotUpdateParamsAcSlotCat = "OTHER"
)

// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
// OTHER).
type AirfieldslotUpdateParamsType string

const (
	AirfieldslotUpdateParamsTypeWorking AirfieldslotUpdateParamsType = "WORKING"
	AirfieldslotUpdateParamsTypeParking AirfieldslotUpdateParamsType = "PARKING"
	AirfieldslotUpdateParamsTypeTakeoff AirfieldslotUpdateParamsType = "TAKEOFF"
	AirfieldslotUpdateParamsTypeLanding AirfieldslotUpdateParamsType = "LANDING"
	AirfieldslotUpdateParamsTypeOther   AirfieldslotUpdateParamsType = "OTHER"
)

type AirfieldslotCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldslotCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirfieldslotCountParams]'s query parameters as
// `url.Values`.
func (r AirfieldslotCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldslotTupleParams struct {
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
func (f AirfieldslotTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirfieldslotTupleParams]'s query parameters as
// `url.Values`.
func (r AirfieldslotTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
