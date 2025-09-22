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

// DeconflictsetService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeconflictsetService] method instead.
type DeconflictsetService struct {
	Options []option.RequestOption
	History DeconflictsetHistoryService
}

// NewDeconflictsetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeconflictsetService(opts ...option.RequestOption) (r DeconflictsetService) {
	r = DeconflictsetService{}
	r.Options = opts
	r.History = NewDeconflictsetHistoryService(opts...)
	return
}

// Service operation to take a single DeconflictSet record as a POST body and
// ingest into the database. This operation does not persist any DeconflictWindow
// datatypes that may be present in the body of the request. This operation is not
// intended to be used for automated feeds into UDL. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *DeconflictsetService) New(ctx context.Context, body DeconflictsetNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/deconflictset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *DeconflictsetService) List(ctx context.Context, query DeconflictsetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[DeconflictsetListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/deconflictset"
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
func (r *DeconflictsetService) ListAutoPaging(ctx context.Context, query DeconflictsetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[DeconflictsetListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *DeconflictsetService) Count(ctx context.Context, query DeconflictsetCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/deconflictset/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single DeconflictSet record by its unique ID passed
// as a path parameter.
func (r *DeconflictsetService) Get(ctx context.Context, id string, query DeconflictsetGetParams, opts ...option.RequestOption) (res *DeconflictsetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/deconflictset/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *DeconflictsetService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *DeconflictsetQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/deconflictset/queryhelp"
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
func (r *DeconflictsetService) Tuple(ctx context.Context, query DeconflictsetTupleParams, opts ...option.RequestOption) (res *[]DeconflictsetTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/deconflictset/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a single DeconflictSet record as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *DeconflictsetService) UnvalidatedPublish(ctx context.Context, body DeconflictsetUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-deconflictset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The DeconflictSet service provides access to a set of DeconflictWindows and
// metadata about those data. A DeconflictWindow describes a time window during
// which an action, such as target engagement, may either occur or is prohibited
// from occurring. The DeconflictWindow model includes information about the
// spatial details for specific target types. A flag is provided to specify whether
// the window should be associated with taking action (OPEN), or if no action
// should occur (CLOSED).
type DeconflictsetListResponse struct {
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
	DataMode DeconflictsetListResponseDataMode `json:"dataMode,required"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// The number of windows provided by this DeconflictSet record.
	NumWindows int64 `json:"numWindows,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The time at which the window calculations completed, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationEndTime time.Time `json:"calculationEndTime" format:"date-time"`
	// The algorithm execution id associated with the generation of this DeconflictSet.
	CalculationID string `json:"calculationId"`
	// The time at which the window calculations started, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationStartTime time.Time `json:"calculationStartTime" format:"date-time"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array of error messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Errors []string `json:"errors"`
	// The end time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventEndTime time.Time `json:"eventEndTime" format:"date-time"`
	// The type of event associated with this DeconflictSet record.
	EventType string `json:"eventType"`
	// The id of the LaserDeconflictRequest record used as input in the generation of
	// this DeconflictSet, if applicable.
	IDLaserDeconflictRequest string `json:"idLaserDeconflictRequest"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	ReferenceFrame string `json:"referenceFrame"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Array of warning messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		EventStartTime           respjson.Field
		NumWindows               respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		CalculationEndTime       respjson.Field
		CalculationID            respjson.Field
		CalculationStartTime     respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		Errors                   respjson.Field
		EventEndTime             respjson.Field
		EventType                respjson.Field
		IDLaserDeconflictRequest respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		ReferenceFrame           respjson.Field
		SourceDl                 respjson.Field
		Tags                     respjson.Field
		TransactionID            respjson.Field
		Warnings                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeconflictsetListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeconflictsetListResponse) UnmarshalJSON(data []byte) error {
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
type DeconflictsetListResponseDataMode string

const (
	DeconflictsetListResponseDataModeReal      DeconflictsetListResponseDataMode = "REAL"
	DeconflictsetListResponseDataModeTest      DeconflictsetListResponseDataMode = "TEST"
	DeconflictsetListResponseDataModeSimulated DeconflictsetListResponseDataMode = "SIMULATED"
	DeconflictsetListResponseDataModeExercise  DeconflictsetListResponseDataMode = "EXERCISE"
)

// The DeconflictSet service provides access to a set of DeconflictWindows and
// metadata about those data. A DeconflictWindow describes a time window during
// which an action, such as target engagement, may either occur or is prohibited
// from occurring. The DeconflictWindow model includes information about the
// spatial details for specific target types. A flag is provided to specify whether
// the window should be associated with taking action (OPEN), or if no action
// should occur (CLOSED).
type DeconflictsetGetResponse struct {
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
	DataMode DeconflictsetGetResponseDataMode `json:"dataMode,required"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// The number of windows provided by this DeconflictSet record.
	NumWindows int64 `json:"numWindows,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The time at which the window calculations completed, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationEndTime time.Time `json:"calculationEndTime" format:"date-time"`
	// The algorithm execution id associated with the generation of this DeconflictSet.
	CalculationID string `json:"calculationId"`
	// The time at which the window calculations started, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationStartTime time.Time `json:"calculationStartTime" format:"date-time"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array of DeconflictWindow records associated with this DeconflictSet.
	DeconflictWindows []DeconflictsetGetResponseDeconflictWindow `json:"deconflictWindows"`
	// Array of error messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Errors []string `json:"errors"`
	// The end time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventEndTime time.Time `json:"eventEndTime" format:"date-time"`
	// The type of event associated with this DeconflictSet record.
	EventType string `json:"eventType"`
	// The id of the LaserDeconflictRequest record used as input in the generation of
	// this DeconflictSet, if applicable.
	IDLaserDeconflictRequest string `json:"idLaserDeconflictRequest"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	ReferenceFrame string `json:"referenceFrame"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Array of warning messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		EventStartTime           respjson.Field
		NumWindows               respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		CalculationEndTime       respjson.Field
		CalculationID            respjson.Field
		CalculationStartTime     respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		DeconflictWindows        respjson.Field
		Errors                   respjson.Field
		EventEndTime             respjson.Field
		EventType                respjson.Field
		IDLaserDeconflictRequest respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		ReferenceFrame           respjson.Field
		SourceDl                 respjson.Field
		Tags                     respjson.Field
		TransactionID            respjson.Field
		Warnings                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeconflictsetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DeconflictsetGetResponse) UnmarshalJSON(data []byte) error {
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
type DeconflictsetGetResponseDataMode string

const (
	DeconflictsetGetResponseDataModeReal      DeconflictsetGetResponseDataMode = "REAL"
	DeconflictsetGetResponseDataModeTest      DeconflictsetGetResponseDataMode = "TEST"
	DeconflictsetGetResponseDataModeSimulated DeconflictsetGetResponseDataMode = "SIMULATED"
	DeconflictsetGetResponseDataModeExercise  DeconflictsetGetResponseDataMode = "EXERCISE"
)

// A DeconflictWindow describes a time window during which an action, such as
// target engagement, may either occur or is prohibited from occurring. The
// DeconflictWindow model includes information about the spatial details for
// specific target types. A flag is provided to specify whether the window should
// be associated with taking action (OPEN), or if no action should occur (CLOSED).
type DeconflictsetGetResponseDeconflictWindow struct {
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
	DataMode string `json:"dataMode,required"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The window start time, in ISO 8601 UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The window stop time, in ISO 8601 UTC format with millisecond precision.
	StopTime time.Time `json:"stopTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The angle at which the victim enters the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfEntry float64 `json:"angleOfEntry"`
	// The angle at which the victim exits the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfExit float64 `json:"angleOfExit"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The X, Y, Z coordinates of entry, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	EntryCoords []float64 `json:"entryCoords"`
	// The type of event associated with the window status.
	EventType string `json:"eventType"`
	// The X, Y, Z coordinates of exit, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	ExitCoords []float64 `json:"exitCoords"`
	// Unique identifier of the parent DeconflictSet, auto-generated by the system. The
	// idDeconflictSet is used to identify all individual windows associated with a
	// parent DeconflictSet record.
	IDDeconflictSet string `json:"idDeconflictSet"`
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
	// The target identifier. If the target is a satellite, the target is the
	// satellite/catalog number of the target on-orbit object.
	Target string `json:"target"`
	// The target type associated with this window (e.g. VICTIM, EARTH, etc.).
	TargetType string `json:"targetType"`
	// The victim identifier associated with this window. If the victim is a satellite,
	// the victim is the satellite/catalog number of the target on-orbit object.
	Victim string `json:"victim"`
	// The window status indicating whether possibility of action may occur. In other
	// words, OPEN is akin to a "green light," during which taking action is warranted
	// or authorized (though not necessarily required) over this timeframe, while
	// CLOSED represents a "red light," meaning that absolutely no action is warranted
	// or authorized to take place during this timeframe.
	WindowType string `json:"windowType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EventStartTime        respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		StopTime              respjson.Field
		ID                    respjson.Field
		AngleOfEntry          respjson.Field
		AngleOfExit           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EntryCoords           respjson.Field
		EventType             respjson.Field
		ExitCoords            respjson.Field
		IDDeconflictSet       respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SourceDl              respjson.Field
		Target                respjson.Field
		TargetType            respjson.Field
		Victim                respjson.Field
		WindowType            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeconflictsetGetResponseDeconflictWindow) RawJSON() string { return r.JSON.raw }
func (r *DeconflictsetGetResponseDeconflictWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeconflictsetQueryhelpResponse struct {
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
func (r DeconflictsetQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *DeconflictsetQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The DeconflictSet service provides access to a set of DeconflictWindows and
// metadata about those data. A DeconflictWindow describes a time window during
// which an action, such as target engagement, may either occur or is prohibited
// from occurring. The DeconflictWindow model includes information about the
// spatial details for specific target types. A flag is provided to specify whether
// the window should be associated with taking action (OPEN), or if no action
// should occur (CLOSED).
type DeconflictsetTupleResponse struct {
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
	DataMode DeconflictsetTupleResponseDataMode `json:"dataMode,required"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// The number of windows provided by this DeconflictSet record.
	NumWindows int64 `json:"numWindows,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The time at which the window calculations completed, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationEndTime time.Time `json:"calculationEndTime" format:"date-time"`
	// The algorithm execution id associated with the generation of this DeconflictSet.
	CalculationID string `json:"calculationId"`
	// The time at which the window calculations started, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationStartTime time.Time `json:"calculationStartTime" format:"date-time"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array of DeconflictWindow records associated with this DeconflictSet.
	DeconflictWindows []DeconflictsetTupleResponseDeconflictWindow `json:"deconflictWindows"`
	// Array of error messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Errors []string `json:"errors"`
	// The end time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventEndTime time.Time `json:"eventEndTime" format:"date-time"`
	// The type of event associated with this DeconflictSet record.
	EventType string `json:"eventType"`
	// The id of the LaserDeconflictRequest record used as input in the generation of
	// this DeconflictSet, if applicable.
	IDLaserDeconflictRequest string `json:"idLaserDeconflictRequest"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	ReferenceFrame string `json:"referenceFrame"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Array of warning messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking    respjson.Field
		DataMode                 respjson.Field
		EventStartTime           respjson.Field
		NumWindows               respjson.Field
		Source                   respjson.Field
		ID                       respjson.Field
		CalculationEndTime       respjson.Field
		CalculationID            respjson.Field
		CalculationStartTime     respjson.Field
		CreatedAt                respjson.Field
		CreatedBy                respjson.Field
		DeconflictWindows        respjson.Field
		Errors                   respjson.Field
		EventEndTime             respjson.Field
		EventType                respjson.Field
		IDLaserDeconflictRequest respjson.Field
		Origin                   respjson.Field
		OrigNetwork              respjson.Field
		ReferenceFrame           respjson.Field
		SourceDl                 respjson.Field
		Tags                     respjson.Field
		TransactionID            respjson.Field
		Warnings                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeconflictsetTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *DeconflictsetTupleResponse) UnmarshalJSON(data []byte) error {
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
type DeconflictsetTupleResponseDataMode string

const (
	DeconflictsetTupleResponseDataModeReal      DeconflictsetTupleResponseDataMode = "REAL"
	DeconflictsetTupleResponseDataModeTest      DeconflictsetTupleResponseDataMode = "TEST"
	DeconflictsetTupleResponseDataModeSimulated DeconflictsetTupleResponseDataMode = "SIMULATED"
	DeconflictsetTupleResponseDataModeExercise  DeconflictsetTupleResponseDataMode = "EXERCISE"
)

// A DeconflictWindow describes a time window during which an action, such as
// target engagement, may either occur or is prohibited from occurring. The
// DeconflictWindow model includes information about the spatial details for
// specific target types. A flag is provided to specify whether the window should
// be associated with taking action (OPEN), or if no action should occur (CLOSED).
type DeconflictsetTupleResponseDeconflictWindow struct {
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
	DataMode string `json:"dataMode,required"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The window start time, in ISO 8601 UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The window stop time, in ISO 8601 UTC format with millisecond precision.
	StopTime time.Time `json:"stopTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// The angle at which the victim enters the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfEntry float64 `json:"angleOfEntry"`
	// The angle at which the victim exits the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfExit float64 `json:"angleOfExit"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The X, Y, Z coordinates of entry, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	EntryCoords []float64 `json:"entryCoords"`
	// The type of event associated with the window status.
	EventType string `json:"eventType"`
	// The X, Y, Z coordinates of exit, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	ExitCoords []float64 `json:"exitCoords"`
	// Unique identifier of the parent DeconflictSet, auto-generated by the system. The
	// idDeconflictSet is used to identify all individual windows associated with a
	// parent DeconflictSet record.
	IDDeconflictSet string `json:"idDeconflictSet"`
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
	// The target identifier. If the target is a satellite, the target is the
	// satellite/catalog number of the target on-orbit object.
	Target string `json:"target"`
	// The target type associated with this window (e.g. VICTIM, EARTH, etc.).
	TargetType string `json:"targetType"`
	// The victim identifier associated with this window. If the victim is a satellite,
	// the victim is the satellite/catalog number of the target on-orbit object.
	Victim string `json:"victim"`
	// The window status indicating whether possibility of action may occur. In other
	// words, OPEN is akin to a "green light," during which taking action is warranted
	// or authorized (though not necessarily required) over this timeframe, while
	// CLOSED represents a "red light," meaning that absolutely no action is warranted
	// or authorized to take place during this timeframe.
	WindowType string `json:"windowType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EventStartTime        respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		StopTime              respjson.Field
		ID                    respjson.Field
		AngleOfEntry          respjson.Field
		AngleOfExit           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EntryCoords           respjson.Field
		EventType             respjson.Field
		ExitCoords            respjson.Field
		IDDeconflictSet       respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SourceDl              respjson.Field
		Target                respjson.Field
		TargetType            respjson.Field
		Victim                respjson.Field
		WindowType            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeconflictsetTupleResponseDeconflictWindow) RawJSON() string { return r.JSON.raw }
func (r *DeconflictsetTupleResponseDeconflictWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeconflictsetNewParams struct {
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
	DataMode DeconflictsetNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// The number of windows provided by this DeconflictSet record.
	NumWindows int64 `json:"numWindows,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The time at which the window calculations completed, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationEndTime param.Opt[time.Time] `json:"calculationEndTime,omitzero" format:"date-time"`
	// The algorithm execution id associated with the generation of this DeconflictSet.
	CalculationID param.Opt[string] `json:"calculationId,omitzero"`
	// The time at which the window calculations started, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationStartTime param.Opt[time.Time] `json:"calculationStartTime,omitzero" format:"date-time"`
	// The end time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventEndTime param.Opt[time.Time] `json:"eventEndTime,omitzero" format:"date-time"`
	// The type of event associated with this DeconflictSet record.
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// The id of the LaserDeconflictRequest record used as input in the generation of
	// this DeconflictSet, if applicable.
	IDLaserDeconflictRequest param.Opt[string] `json:"idLaserDeconflictRequest,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	ReferenceFrame param.Opt[string] `json:"referenceFrame,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Array of DeconflictWindow records associated with this DeconflictSet.
	DeconflictWindows []DeconflictsetNewParamsDeconflictWindow `json:"deconflictWindows,omitzero"`
	// Array of error messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Errors []string `json:"errors,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// Array of warning messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Warnings []string `json:"warnings,omitzero"`
	paramObj
}

func (r DeconflictsetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DeconflictsetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeconflictsetNewParams) UnmarshalJSON(data []byte) error {
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
type DeconflictsetNewParamsDataMode string

const (
	DeconflictsetNewParamsDataModeReal      DeconflictsetNewParamsDataMode = "REAL"
	DeconflictsetNewParamsDataModeTest      DeconflictsetNewParamsDataMode = "TEST"
	DeconflictsetNewParamsDataModeSimulated DeconflictsetNewParamsDataMode = "SIMULATED"
	DeconflictsetNewParamsDataModeExercise  DeconflictsetNewParamsDataMode = "EXERCISE"
)

// A DeconflictWindow describes a time window during which an action, such as
// target engagement, may either occur or is prohibited from occurring. The
// DeconflictWindow model includes information about the spatial details for
// specific target types. A flag is provided to specify whether the window should
// be associated with taking action (OPEN), or if no action should occur (CLOSED).
//
// The properties ClassificationMarking, DataMode, EventStartTime, Source,
// StartTime, StopTime are required.
type DeconflictsetNewParamsDeconflictWindow struct {
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
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The window start time, in ISO 8601 UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The window stop time, in ISO 8601 UTC format with millisecond precision.
	StopTime time.Time `json:"stopTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The angle at which the victim enters the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfEntry param.Opt[float64] `json:"angleOfEntry,omitzero"`
	// The angle at which the victim exits the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfExit param.Opt[float64] `json:"angleOfExit,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The type of event associated with the window status.
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// Unique identifier of the parent DeconflictSet, auto-generated by the system. The
	// idDeconflictSet is used to identify all individual windows associated with a
	// parent DeconflictSet record.
	IDDeconflictSet param.Opt[string] `json:"idDeconflictSet,omitzero"`
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
	// The target identifier. If the target is a satellite, the target is the
	// satellite/catalog number of the target on-orbit object.
	Target param.Opt[string] `json:"target,omitzero"`
	// The target type associated with this window (e.g. VICTIM, EARTH, etc.).
	TargetType param.Opt[string] `json:"targetType,omitzero"`
	// The victim identifier associated with this window. If the victim is a satellite,
	// the victim is the satellite/catalog number of the target on-orbit object.
	Victim param.Opt[string] `json:"victim,omitzero"`
	// The window status indicating whether possibility of action may occur. In other
	// words, OPEN is akin to a "green light," during which taking action is warranted
	// or authorized (though not necessarily required) over this timeframe, while
	// CLOSED represents a "red light," meaning that absolutely no action is warranted
	// or authorized to take place during this timeframe.
	WindowType param.Opt[string] `json:"windowType,omitzero"`
	// The X, Y, Z coordinates of entry, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	EntryCoords []float64 `json:"entryCoords,omitzero"`
	// The X, Y, Z coordinates of exit, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	ExitCoords []float64 `json:"exitCoords,omitzero"`
	paramObj
}

func (r DeconflictsetNewParamsDeconflictWindow) MarshalJSON() (data []byte, err error) {
	type shadow DeconflictsetNewParamsDeconflictWindow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeconflictsetNewParamsDeconflictWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DeconflictsetNewParamsDeconflictWindow](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type DeconflictsetListParams struct {
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EventStartTime time.Time        `query:"eventStartTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeconflictsetListParams]'s query parameters as
// `url.Values`.
func (r DeconflictsetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeconflictsetCountParams struct {
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EventStartTime time.Time        `query:"eventStartTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeconflictsetCountParams]'s query parameters as
// `url.Values`.
func (r DeconflictsetCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeconflictsetGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeconflictsetGetParams]'s query parameters as `url.Values`.
func (r DeconflictsetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeconflictsetTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EventStartTime time.Time        `query:"eventStartTime,required" format:"date-time" json:"-"`
	FirstResult    param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults     param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeconflictsetTupleParams]'s query parameters as
// `url.Values`.
func (r DeconflictsetTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeconflictsetUnvalidatedPublishParams struct {
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
	DataMode DeconflictsetUnvalidatedPublishParamsDataMode `json:"dataMode,omitzero,required"`
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// The number of windows provided by this DeconflictSet record.
	NumWindows int64 `json:"numWindows,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The time at which the window calculations completed, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationEndTime param.Opt[time.Time] `json:"calculationEndTime,omitzero" format:"date-time"`
	// The algorithm execution id associated with the generation of this DeconflictSet.
	CalculationID param.Opt[string] `json:"calculationId,omitzero"`
	// The time at which the window calculations started, in ISO 8601 UTC format with
	// millisecond precision.
	CalculationStartTime param.Opt[time.Time] `json:"calculationStartTime,omitzero" format:"date-time"`
	// The end time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventEndTime param.Opt[time.Time] `json:"eventEndTime,omitzero" format:"date-time"`
	// The type of event associated with this DeconflictSet record.
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// The id of the LaserDeconflictRequest record used as input in the generation of
	// this DeconflictSet, if applicable.
	IDLaserDeconflictRequest param.Opt[string] `json:"idLaserDeconflictRequest,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	ReferenceFrame param.Opt[string] `json:"referenceFrame,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Array of DeconflictWindow records associated with this DeconflictSet.
	DeconflictWindows []DeconflictsetUnvalidatedPublishParamsDeconflictWindow `json:"deconflictWindows,omitzero"`
	// Array of error messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Errors []string `json:"errors,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// Array of warning messages that potentially contain information about the reasons
	// this deconflict response calculation may be inaccurate, or why it failed.
	Warnings []string `json:"warnings,omitzero"`
	paramObj
}

func (r DeconflictsetUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	type shadow DeconflictsetUnvalidatedPublishParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeconflictsetUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
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
type DeconflictsetUnvalidatedPublishParamsDataMode string

const (
	DeconflictsetUnvalidatedPublishParamsDataModeReal      DeconflictsetUnvalidatedPublishParamsDataMode = "REAL"
	DeconflictsetUnvalidatedPublishParamsDataModeTest      DeconflictsetUnvalidatedPublishParamsDataMode = "TEST"
	DeconflictsetUnvalidatedPublishParamsDataModeSimulated DeconflictsetUnvalidatedPublishParamsDataMode = "SIMULATED"
	DeconflictsetUnvalidatedPublishParamsDataModeExercise  DeconflictsetUnvalidatedPublishParamsDataMode = "EXERCISE"
)

// A DeconflictWindow describes a time window during which an action, such as
// target engagement, may either occur or is prohibited from occurring. The
// DeconflictWindow model includes information about the spatial details for
// specific target types. A flag is provided to specify whether the window should
// be associated with taking action (OPEN), or if no action should occur (CLOSED).
//
// The properties ClassificationMarking, DataMode, EventStartTime, Source,
// StartTime, StopTime are required.
type DeconflictsetUnvalidatedPublishParamsDeconflictWindow struct {
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
	// The start time of the event associated with the set of DeconflictWindow records,
	// in ISO 8601 UTC format with millisecond precision.
	EventStartTime time.Time `json:"eventStartTime,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The window start time, in ISO 8601 UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The window stop time, in ISO 8601 UTC format with millisecond precision.
	StopTime time.Time `json:"stopTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// The angle at which the victim enters the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfEntry param.Opt[float64] `json:"angleOfEntry,omitzero"`
	// The angle at which the victim exits the target zone in reference to the emitter
	// source location, in degrees.
	AngleOfExit param.Opt[float64] `json:"angleOfExit,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The type of event associated with the window status.
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// Unique identifier of the parent DeconflictSet, auto-generated by the system. The
	// idDeconflictSet is used to identify all individual windows associated with a
	// parent DeconflictSet record.
	IDDeconflictSet param.Opt[string] `json:"idDeconflictSet,omitzero"`
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
	// The target identifier. If the target is a satellite, the target is the
	// satellite/catalog number of the target on-orbit object.
	Target param.Opt[string] `json:"target,omitzero"`
	// The target type associated with this window (e.g. VICTIM, EARTH, etc.).
	TargetType param.Opt[string] `json:"targetType,omitzero"`
	// The victim identifier associated with this window. If the victim is a satellite,
	// the victim is the satellite/catalog number of the target on-orbit object.
	Victim param.Opt[string] `json:"victim,omitzero"`
	// The window status indicating whether possibility of action may occur. In other
	// words, OPEN is akin to a "green light," during which taking action is warranted
	// or authorized (though not necessarily required) over this timeframe, while
	// CLOSED represents a "red light," meaning that absolutely no action is warranted
	// or authorized to take place during this timeframe.
	WindowType param.Opt[string] `json:"windowType,omitzero"`
	// The X, Y, Z coordinates of entry, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	EntryCoords []float64 `json:"entryCoords,omitzero"`
	// The X, Y, Z coordinates of exit, in the reference frame specified by the parent
	// DeconflictSet record, in meters.
	ExitCoords []float64 `json:"exitCoords,omitzero"`
	paramObj
}

func (r DeconflictsetUnvalidatedPublishParamsDeconflictWindow) MarshalJSON() (data []byte, err error) {
	type shadow DeconflictsetUnvalidatedPublishParamsDeconflictWindow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeconflictsetUnvalidatedPublishParamsDeconflictWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DeconflictsetUnvalidatedPublishParamsDeconflictWindow](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
