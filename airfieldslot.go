// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// AirfieldSlotService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirfieldSlotService] method instead.
type AirfieldSlotService struct {
	Options []option.RequestOption
}

// NewAirfieldSlotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAirfieldSlotService(opts ...option.RequestOption) (r AirfieldSlotService) {
	r = AirfieldSlotService{}
	r.Options = opts
	return
}

// Service operation to take a single airfieldslot record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AirfieldSlotService) New(ctx context.Context, body AirfieldSlotNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airfieldslot"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single airfieldslot record by its unique ID passed as
// a path parameter.
func (r *AirfieldSlotService) Get(ctx context.Context, id string, query AirfieldSlotGetParams, opts ...option.RequestOption) (res *AirfieldslotFull, err error) {
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
func (r *AirfieldSlotService) Update(ctx context.Context, id string, body AirfieldSlotUpdateParams, opts ...option.RequestOption) (err error) {
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

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirfieldSlotService) List(ctx context.Context, query AirfieldSlotListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AirfieldslotAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/airfieldslot"
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
func (r *AirfieldSlotService) ListAutoPaging(ctx context.Context, query AirfieldSlotListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AirfieldslotAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an airfieldslot record specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *AirfieldSlotService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
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
func (r *AirfieldSlotService) Count(ctx context.Context, query AirfieldSlotCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airfieldslot/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AirfieldSlotService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
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
func (r *AirfieldSlotService) Tuple(ctx context.Context, query AirfieldSlotTupleParams, opts ...option.RequestOption) (res *[]AirfieldslotFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airfieldslot/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Airfield capacity data. Contains data associated with the airfieldslots
// available for parking, working, takeoff, and landing at the airfield, as well as
// the types of aircraft that can be accommodated.
type AirfieldslotAbridged struct {
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
	DataMode AirfieldslotAbridgedDataMode `json:"dataMode,required"`
	// Name of this slot.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Largest category of aircraft supported in this slot (WIDE, NARROW, HELO, ALL,
	// OTHER).
	//
	// Any of "WIDE", "NARROW", "HELO", "ALL", "OTHER".
	AcSlotCat AirfieldslotAbridgedAcSlotCat `json:"acSlotCat"`
	// Alternate airfield identifier provided by the source.
	AltAirfieldID string `json:"altAirfieldId"`
	// Number of aircraft that can fit in this slot at the same time.
	Capacity int64 `json:"capacity"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Latest zulu time this slot is available based on daily standard hours. Not
	// applicable to slots with type PARKING. Abnormal hours, such as holidays, should
	// be marked via the AirfieldSlotConsumption schema.
	EndTime string `json:"endTime"`
	// The International Civil Aviation Organization (ICAO) code of the airfield.
	Icao string `json:"icao"`
	// Unique identifier of the Airfield for which this slot information applies.
	IDAirfield string `json:"idAirfield"`
	// Minimum time that must elapse between different aircraft leaving and entering
	// this slot, in minutes.
	MinSeparation int64 `json:"minSeparation"`
	// Optional notes/comments for this airfield slot.
	Notes string `json:"notes"`
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
	// Zulu time this slot is first available based on daily standard hours. Not
	// applicable to slots with type PARKING. Abnormal hours, such as holidays, should
	// be marked via the AirfieldSlotConsumption schema.
	StartTime string `json:"startTime"`
	// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
	// OTHER).
	//
	// Any of "WORKING", "PARKING", "TAKEOFF", "LANDING", "OTHER".
	Type AirfieldslotAbridgedType `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AirfieldName          resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Name                  resp.Field
		Source                resp.Field
		ID                    resp.Field
		AcSlotCat             resp.Field
		AltAirfieldID         resp.Field
		Capacity              resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EndTime               resp.Field
		Icao                  resp.Field
		IDAirfield            resp.Field
		MinSeparation         resp.Field
		Notes                 resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		SourceDl              resp.Field
		StartTime             resp.Field
		Type                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirfieldslotAbridged) RawJSON() string { return r.JSON.raw }
func (r *AirfieldslotAbridged) UnmarshalJSON(data []byte) error {
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
type AirfieldslotAbridgedDataMode string

const (
	AirfieldslotAbridgedDataModeReal      AirfieldslotAbridgedDataMode = "REAL"
	AirfieldslotAbridgedDataModeTest      AirfieldslotAbridgedDataMode = "TEST"
	AirfieldslotAbridgedDataModeSimulated AirfieldslotAbridgedDataMode = "SIMULATED"
	AirfieldslotAbridgedDataModeExercise  AirfieldslotAbridgedDataMode = "EXERCISE"
)

// Largest category of aircraft supported in this slot (WIDE, NARROW, HELO, ALL,
// OTHER).
type AirfieldslotAbridgedAcSlotCat string

const (
	AirfieldslotAbridgedAcSlotCatWide   AirfieldslotAbridgedAcSlotCat = "WIDE"
	AirfieldslotAbridgedAcSlotCatNarrow AirfieldslotAbridgedAcSlotCat = "NARROW"
	AirfieldslotAbridgedAcSlotCatHelo   AirfieldslotAbridgedAcSlotCat = "HELO"
	AirfieldslotAbridgedAcSlotCatAll    AirfieldslotAbridgedAcSlotCat = "ALL"
	AirfieldslotAbridgedAcSlotCatOther  AirfieldslotAbridgedAcSlotCat = "OTHER"
)

// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
// OTHER).
type AirfieldslotAbridgedType string

const (
	AirfieldslotAbridgedTypeWorking AirfieldslotAbridgedType = "WORKING"
	AirfieldslotAbridgedTypeParking AirfieldslotAbridgedType = "PARKING"
	AirfieldslotAbridgedTypeTakeoff AirfieldslotAbridgedType = "TAKEOFF"
	AirfieldslotAbridgedTypeLanding AirfieldslotAbridgedType = "LANDING"
	AirfieldslotAbridgedTypeOther   AirfieldslotAbridgedType = "OTHER"
)

// Airfield capacity data. Contains data associated with the airfieldslots
// available for parking, working, takeoff, and landing at the airfield, as well as
// the types of aircraft that can be accommodated.
type AirfieldslotFull struct {
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
	DataMode AirfieldslotFullDataMode `json:"dataMode,required"`
	// Name of this slot.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Largest category of aircraft supported in this slot (WIDE, NARROW, HELO, ALL,
	// OTHER).
	//
	// Any of "WIDE", "NARROW", "HELO", "ALL", "OTHER".
	AcSlotCat AirfieldslotFullAcSlotCat `json:"acSlotCat"`
	// Alternate airfield identifier provided by the source.
	AltAirfieldID string `json:"altAirfieldId"`
	// Number of aircraft that can fit in this slot at the same time.
	Capacity int64 `json:"capacity"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Latest zulu time this slot is available based on daily standard hours. Not
	// applicable to slots with type PARKING. Abnormal hours, such as holidays, should
	// be marked via the AirfieldSlotConsumption schema.
	EndTime string `json:"endTime"`
	// The International Civil Aviation Organization (ICAO) code of the airfield.
	Icao string `json:"icao"`
	// Unique identifier of the Airfield for which this slot information applies.
	IDAirfield string `json:"idAirfield"`
	// Minimum time that must elapse between different aircraft leaving and entering
	// this slot, in minutes.
	MinSeparation int64 `json:"minSeparation"`
	// Optional notes/comments for this airfield slot.
	Notes string `json:"notes"`
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
	// Zulu time this slot is first available based on daily standard hours. Not
	// applicable to slots with type PARKING. Abnormal hours, such as holidays, should
	// be marked via the AirfieldSlotConsumption schema.
	StartTime string `json:"startTime"`
	// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
	// OTHER).
	//
	// Any of "WORKING", "PARKING", "TAKEOFF", "LANDING", "OTHER".
	Type AirfieldslotFullType `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AirfieldName          resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Name                  resp.Field
		Source                resp.Field
		ID                    resp.Field
		AcSlotCat             resp.Field
		AltAirfieldID         resp.Field
		Capacity              resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		EndTime               resp.Field
		Icao                  resp.Field
		IDAirfield            resp.Field
		MinSeparation         resp.Field
		Notes                 resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		SourceDl              resp.Field
		StartTime             resp.Field
		Type                  resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirfieldslotFull) RawJSON() string { return r.JSON.raw }
func (r *AirfieldslotFull) UnmarshalJSON(data []byte) error {
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
type AirfieldslotFullDataMode string

const (
	AirfieldslotFullDataModeReal      AirfieldslotFullDataMode = "REAL"
	AirfieldslotFullDataModeTest      AirfieldslotFullDataMode = "TEST"
	AirfieldslotFullDataModeSimulated AirfieldslotFullDataMode = "SIMULATED"
	AirfieldslotFullDataModeExercise  AirfieldslotFullDataMode = "EXERCISE"
)

// Largest category of aircraft supported in this slot (WIDE, NARROW, HELO, ALL,
// OTHER).
type AirfieldslotFullAcSlotCat string

const (
	AirfieldslotFullAcSlotCatWide   AirfieldslotFullAcSlotCat = "WIDE"
	AirfieldslotFullAcSlotCatNarrow AirfieldslotFullAcSlotCat = "NARROW"
	AirfieldslotFullAcSlotCatHelo   AirfieldslotFullAcSlotCat = "HELO"
	AirfieldslotFullAcSlotCatAll    AirfieldslotFullAcSlotCat = "ALL"
	AirfieldslotFullAcSlotCatOther  AirfieldslotFullAcSlotCat = "OTHER"
)

// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
// OTHER).
type AirfieldslotFullType string

const (
	AirfieldslotFullTypeWorking AirfieldslotFullType = "WORKING"
	AirfieldslotFullTypeParking AirfieldslotFullType = "PARKING"
	AirfieldslotFullTypeTakeoff AirfieldslotFullType = "TAKEOFF"
	AirfieldslotFullTypeLanding AirfieldslotFullType = "LANDING"
	AirfieldslotFullTypeOther   AirfieldslotFullType = "OTHER"
)

type AirfieldSlotNewParams struct {
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
	DataMode AirfieldSlotNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	AcSlotCat AirfieldSlotNewParamsAcSlotCat `json:"acSlotCat,omitzero"`
	// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
	// OTHER).
	//
	// Any of "WORKING", "PARKING", "TAKEOFF", "LANDING", "OTHER".
	Type AirfieldSlotNewParamsType `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldSlotNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r AirfieldSlotNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AirfieldSlotNewParams
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
type AirfieldSlotNewParamsDataMode string

const (
	AirfieldSlotNewParamsDataModeReal      AirfieldSlotNewParamsDataMode = "REAL"
	AirfieldSlotNewParamsDataModeTest      AirfieldSlotNewParamsDataMode = "TEST"
	AirfieldSlotNewParamsDataModeSimulated AirfieldSlotNewParamsDataMode = "SIMULATED"
	AirfieldSlotNewParamsDataModeExercise  AirfieldSlotNewParamsDataMode = "EXERCISE"
)

// Largest category of aircraft supported in this slot (WIDE, NARROW, HELO, ALL,
// OTHER).
type AirfieldSlotNewParamsAcSlotCat string

const (
	AirfieldSlotNewParamsAcSlotCatWide   AirfieldSlotNewParamsAcSlotCat = "WIDE"
	AirfieldSlotNewParamsAcSlotCatNarrow AirfieldSlotNewParamsAcSlotCat = "NARROW"
	AirfieldSlotNewParamsAcSlotCatHelo   AirfieldSlotNewParamsAcSlotCat = "HELO"
	AirfieldSlotNewParamsAcSlotCatAll    AirfieldSlotNewParamsAcSlotCat = "ALL"
	AirfieldSlotNewParamsAcSlotCatOther  AirfieldSlotNewParamsAcSlotCat = "OTHER"
)

// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
// OTHER).
type AirfieldSlotNewParamsType string

const (
	AirfieldSlotNewParamsTypeWorking AirfieldSlotNewParamsType = "WORKING"
	AirfieldSlotNewParamsTypeParking AirfieldSlotNewParamsType = "PARKING"
	AirfieldSlotNewParamsTypeTakeoff AirfieldSlotNewParamsType = "TAKEOFF"
	AirfieldSlotNewParamsTypeLanding AirfieldSlotNewParamsType = "LANDING"
	AirfieldSlotNewParamsTypeOther   AirfieldSlotNewParamsType = "OTHER"
)

type AirfieldSlotGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldSlotGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirfieldSlotGetParams]'s query parameters as `url.Values`.
func (r AirfieldSlotGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldSlotUpdateParams struct {
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
	DataMode AirfieldSlotUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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
	AcSlotCat AirfieldSlotUpdateParamsAcSlotCat `json:"acSlotCat,omitzero"`
	// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
	// OTHER).
	//
	// Any of "WORKING", "PARKING", "TAKEOFF", "LANDING", "OTHER".
	Type AirfieldSlotUpdateParamsType `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldSlotUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r AirfieldSlotUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AirfieldSlotUpdateParams
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
type AirfieldSlotUpdateParamsDataMode string

const (
	AirfieldSlotUpdateParamsDataModeReal      AirfieldSlotUpdateParamsDataMode = "REAL"
	AirfieldSlotUpdateParamsDataModeTest      AirfieldSlotUpdateParamsDataMode = "TEST"
	AirfieldSlotUpdateParamsDataModeSimulated AirfieldSlotUpdateParamsDataMode = "SIMULATED"
	AirfieldSlotUpdateParamsDataModeExercise  AirfieldSlotUpdateParamsDataMode = "EXERCISE"
)

// Largest category of aircraft supported in this slot (WIDE, NARROW, HELO, ALL,
// OTHER).
type AirfieldSlotUpdateParamsAcSlotCat string

const (
	AirfieldSlotUpdateParamsAcSlotCatWide   AirfieldSlotUpdateParamsAcSlotCat = "WIDE"
	AirfieldSlotUpdateParamsAcSlotCatNarrow AirfieldSlotUpdateParamsAcSlotCat = "NARROW"
	AirfieldSlotUpdateParamsAcSlotCatHelo   AirfieldSlotUpdateParamsAcSlotCat = "HELO"
	AirfieldSlotUpdateParamsAcSlotCatAll    AirfieldSlotUpdateParamsAcSlotCat = "ALL"
	AirfieldSlotUpdateParamsAcSlotCatOther  AirfieldSlotUpdateParamsAcSlotCat = "OTHER"
)

// Designates how this slot can be used (WORKING, PARKING, TAKEOFF, LANDING,
// OTHER).
type AirfieldSlotUpdateParamsType string

const (
	AirfieldSlotUpdateParamsTypeWorking AirfieldSlotUpdateParamsType = "WORKING"
	AirfieldSlotUpdateParamsTypeParking AirfieldSlotUpdateParamsType = "PARKING"
	AirfieldSlotUpdateParamsTypeTakeoff AirfieldSlotUpdateParamsType = "TAKEOFF"
	AirfieldSlotUpdateParamsTypeLanding AirfieldSlotUpdateParamsType = "LANDING"
	AirfieldSlotUpdateParamsTypeOther   AirfieldSlotUpdateParamsType = "OTHER"
)

type AirfieldSlotListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldSlotListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirfieldSlotListParams]'s query parameters as `url.Values`.
func (r AirfieldSlotListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldSlotCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AirfieldSlotCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirfieldSlotCountParams]'s query parameters as
// `url.Values`.
func (r AirfieldSlotCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldSlotTupleParams struct {
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
func (f AirfieldSlotTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AirfieldSlotTupleParams]'s query parameters as
// `url.Values`.
func (r AirfieldSlotTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
