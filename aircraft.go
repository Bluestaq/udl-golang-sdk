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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// AircraftService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAircraftService] method instead.
type AircraftService struct {
	Options []option.RequestOption
}

// NewAircraftService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAircraftService(opts ...option.RequestOption) (r AircraftService) {
	r = AircraftService{}
	r.Options = opts
	return
}

// Service operation to take a single Aircraft as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *AircraftService) New(ctx context.Context, body AircraftNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aircraft"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Aircraft record by its unique ID passed as a
// path parameter.
func (r *AircraftService) Get(ctx context.Context, id string, query AircraftGetParams, opts ...option.RequestOption) (res *AircraftFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraft/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single Aircraft. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *AircraftService) Update(ctx context.Context, id string, body AircraftUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aircraft/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AircraftService) List(ctx context.Context, query AircraftListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AircraftAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/aircraft"
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
func (r *AircraftService) ListAutoPaging(ctx context.Context, query AircraftListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AircraftAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AircraftService) Count(ctx context.Context, query AircraftCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/aircraft/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AircraftService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aircraft/queryhelp"
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
func (r *AircraftService) TupleQuery(ctx context.Context, query AircraftTupleQueryParams, opts ...option.RequestOption) (res *[]AircraftFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/aircraft/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// General aircraft designation, characteristics, and capabilities. The aircraft
// schema contains static data of specific aircraft, including tail number, cruise
// speed, max speed, and minimum required runway length, etc.
type AircraftAbridged struct {
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds string `json:"aircraftMDS,required"`
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
	DataMode AircraftAbridgedDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The category of aircraft (e.g. M = Military, C = Commercial).
	Category string `json:"category"`
	// The Air Force major command (MAJCOM) overseeing the aircraft.
	Command string `json:"command"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The cruise speed of the aircraft, in kilometers/hour.
	CruiseSpeed float64 `json:"cruiseSpeed"`
	// Military data network data transfer device ID for this aircraft.
	Dtd string `json:"dtd"`
	// ID of the parent entity for this aircraft.
	IDEntity string `json:"idEntity"`
	// The maximum air speed of the aircraft, in kilometers/hour.
	MaxSpeed float64 `json:"maxSpeed"`
	// The minimum length of runway required to land this aircraft, in feet. Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	MinReqRunwayFt int64 `json:"minReqRunwayFt"`
	// The minimum length of runway required to land this aircraft, in meters. Note:
	// The corresponding equivalent field is not converted by the UDL and may or may
	// not be supplied by the provider. The provider/consumer is responsible for all
	// unit conversions.
	MinReqRunwayM int64 `json:"minReqRunwayM"`
	// The nominal turnaround time for this aircraft, in minutes.
	NominalTaTime int64 `json:"nominalTATime"`
	// Optional notes/comments for this aircraft.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The wing or unit that owns the aircraft.
	Owner string `json:"owner"`
	// Full serial number of the aircraft.
	SerialNumber string `json:"serialNumber"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The tail number of this aircraft.
	TailNumber string `json:"tailNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AircraftMds           respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Category              respjson.Field
		Command               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CruiseSpeed           respjson.Field
		Dtd                   respjson.Field
		IDEntity              respjson.Field
		MaxSpeed              respjson.Field
		MinReqRunwayFt        respjson.Field
		MinReqRunwayM         respjson.Field
		NominalTaTime         respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Owner                 respjson.Field
		SerialNumber          respjson.Field
		SourceDl              respjson.Field
		TailNumber            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AircraftAbridged) RawJSON() string { return r.JSON.raw }
func (r *AircraftAbridged) UnmarshalJSON(data []byte) error {
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
type AircraftAbridgedDataMode string

const (
	AircraftAbridgedDataModeReal      AircraftAbridgedDataMode = "REAL"
	AircraftAbridgedDataModeTest      AircraftAbridgedDataMode = "TEST"
	AircraftAbridgedDataModeSimulated AircraftAbridgedDataMode = "SIMULATED"
	AircraftAbridgedDataModeExercise  AircraftAbridgedDataMode = "EXERCISE"
)

// General aircraft designation, characteristics, and capabilities. The aircraft
// schema contains static data of specific aircraft, including tail number, cruise
// speed, max speed, and minimum required runway length, etc.
type AircraftFull struct {
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds string `json:"aircraftMDS,required"`
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
	DataMode AircraftFullDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The category of aircraft (e.g. M = Military, C = Commercial).
	Category string `json:"category"`
	// The Air Force major command (MAJCOM) overseeing the aircraft.
	Command string `json:"command"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The cruise speed of the aircraft, in kilometers/hour.
	CruiseSpeed float64 `json:"cruiseSpeed"`
	// Military data network data transfer device ID for this aircraft.
	Dtd string `json:"dtd"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityFull `json:"entity"`
	// ID of the parent entity for this aircraft.
	IDEntity string `json:"idEntity"`
	// The maximum air speed of the aircraft, in kilometers/hour.
	MaxSpeed float64 `json:"maxSpeed"`
	// The minimum length of runway required to land this aircraft, in feet. Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	MinReqRunwayFt int64 `json:"minReqRunwayFt"`
	// The minimum length of runway required to land this aircraft, in meters. Note:
	// The corresponding equivalent field is not converted by the UDL and may or may
	// not be supplied by the provider. The provider/consumer is responsible for all
	// unit conversions.
	MinReqRunwayM int64 `json:"minReqRunwayM"`
	// The nominal turnaround time for this aircraft, in minutes.
	NominalTaTime int64 `json:"nominalTATime"`
	// Optional notes/comments for this aircraft.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The wing or unit that owns the aircraft.
	Owner string `json:"owner"`
	// Full serial number of the aircraft.
	SerialNumber string `json:"serialNumber"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The tail number of this aircraft.
	TailNumber string `json:"tailNumber"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AircraftMds           respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Category              respjson.Field
		Command               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CruiseSpeed           respjson.Field
		Dtd                   respjson.Field
		Entity                respjson.Field
		IDEntity              respjson.Field
		MaxSpeed              respjson.Field
		MinReqRunwayFt        respjson.Field
		MinReqRunwayM         respjson.Field
		NominalTaTime         respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Owner                 respjson.Field
		SerialNumber          respjson.Field
		SourceDl              respjson.Field
		TailNumber            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AircraftFull) RawJSON() string { return r.JSON.raw }
func (r *AircraftFull) UnmarshalJSON(data []byte) error {
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
type AircraftFullDataMode string

const (
	AircraftFullDataModeReal      AircraftFullDataMode = "REAL"
	AircraftFullDataModeTest      AircraftFullDataMode = "TEST"
	AircraftFullDataModeSimulated AircraftFullDataMode = "SIMULATED"
	AircraftFullDataModeExercise  AircraftFullDataMode = "EXERCISE"
)

type AircraftNewParams struct {
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds string `json:"aircraftMDS,required"`
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
	DataMode AircraftNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The category of aircraft (e.g. M = Military, C = Commercial).
	Category param.Opt[string] `json:"category,omitzero"`
	// The Air Force major command (MAJCOM) overseeing the aircraft.
	Command param.Opt[string] `json:"command,omitzero"`
	// The cruise speed of the aircraft, in kilometers/hour.
	CruiseSpeed param.Opt[float64] `json:"cruiseSpeed,omitzero"`
	// Military data network data transfer device ID for this aircraft.
	Dtd param.Opt[string] `json:"dtd,omitzero"`
	// ID of the parent entity for this aircraft.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// The maximum air speed of the aircraft, in kilometers/hour.
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// The minimum length of runway required to land this aircraft, in feet. Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	MinReqRunwayFt param.Opt[int64] `json:"minReqRunwayFt,omitzero"`
	// The minimum length of runway required to land this aircraft, in meters. Note:
	// The corresponding equivalent field is not converted by the UDL and may or may
	// not be supplied by the provider. The provider/consumer is responsible for all
	// unit conversions.
	MinReqRunwayM param.Opt[int64] `json:"minReqRunwayM,omitzero"`
	// The nominal turnaround time for this aircraft, in minutes.
	NominalTaTime param.Opt[int64] `json:"nominalTATime,omitzero"`
	// Optional notes/comments for this aircraft.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The wing or unit that owns the aircraft.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// Full serial number of the aircraft.
	SerialNumber param.Opt[string] `json:"serialNumber,omitzero"`
	// The tail number of this aircraft.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r AircraftNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AircraftNewParams
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
type AircraftNewParamsDataMode string

const (
	AircraftNewParamsDataModeReal      AircraftNewParamsDataMode = "REAL"
	AircraftNewParamsDataModeTest      AircraftNewParamsDataMode = "TEST"
	AircraftNewParamsDataModeSimulated AircraftNewParamsDataMode = "SIMULATED"
	AircraftNewParamsDataModeExercise  AircraftNewParamsDataMode = "EXERCISE"
)

type AircraftGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftGetParams]'s query parameters as `url.Values`.
func (r AircraftGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftUpdateParams struct {
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of this aircraft. Intended as, but not constrained
	// to, MIL-STD-6016 environment dependent specific type designations.
	AircraftMds string `json:"aircraftMDS,required"`
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
	DataMode AircraftUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The category of aircraft (e.g. M = Military, C = Commercial).
	Category param.Opt[string] `json:"category,omitzero"`
	// The Air Force major command (MAJCOM) overseeing the aircraft.
	Command param.Opt[string] `json:"command,omitzero"`
	// The cruise speed of the aircraft, in kilometers/hour.
	CruiseSpeed param.Opt[float64] `json:"cruiseSpeed,omitzero"`
	// Military data network data transfer device ID for this aircraft.
	Dtd param.Opt[string] `json:"dtd,omitzero"`
	// ID of the parent entity for this aircraft.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// The maximum air speed of the aircraft, in kilometers/hour.
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// The minimum length of runway required to land this aircraft, in feet. Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	MinReqRunwayFt param.Opt[int64] `json:"minReqRunwayFt,omitzero"`
	// The minimum length of runway required to land this aircraft, in meters. Note:
	// The corresponding equivalent field is not converted by the UDL and may or may
	// not be supplied by the provider. The provider/consumer is responsible for all
	// unit conversions.
	MinReqRunwayM param.Opt[int64] `json:"minReqRunwayM,omitzero"`
	// The nominal turnaround time for this aircraft, in minutes.
	NominalTaTime param.Opt[int64] `json:"nominalTATime,omitzero"`
	// Optional notes/comments for this aircraft.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The wing or unit that owns the aircraft.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// Full serial number of the aircraft.
	SerialNumber param.Opt[string] `json:"serialNumber,omitzero"`
	// The tail number of this aircraft.
	TailNumber param.Opt[string] `json:"tailNumber,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r AircraftUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AircraftUpdateParams
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
type AircraftUpdateParamsDataMode string

const (
	AircraftUpdateParamsDataModeReal      AircraftUpdateParamsDataMode = "REAL"
	AircraftUpdateParamsDataModeTest      AircraftUpdateParamsDataMode = "TEST"
	AircraftUpdateParamsDataModeSimulated AircraftUpdateParamsDataMode = "SIMULATED"
	AircraftUpdateParamsDataModeExercise  AircraftUpdateParamsDataMode = "EXERCISE"
)

type AircraftListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftListParams]'s query parameters as `url.Values`.
func (r AircraftListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftCountParams]'s query parameters as `url.Values`.
func (r AircraftCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AircraftTupleQueryParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AircraftTupleQueryParams]'s query parameters as
// `url.Values`.
func (r AircraftTupleQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
