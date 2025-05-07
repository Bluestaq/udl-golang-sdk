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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// AttitudeSetService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttitudeSetService] method instead.
type AttitudeSetService struct {
	Options []option.RequestOption
	History AttitudeSetHistoryService
}

// NewAttitudeSetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttitudeSetService(opts ...option.RequestOption) (r AttitudeSetService) {
	r = AttitudeSetService{}
	r.Options = opts
	r.History = NewAttitudeSetHistoryService(opts...)
	return
}

// Service operation intended for initial integration only. Takes a single
// AttitudeSet as a POST body and ingest into the database. This operation does not
// persist any Onorbit Attitude points that may be present in the body of the
// request. This operation is not intended to be used for automated feeds into UDL.
// A specific role is required to perform this service operation. Please contact
// the UDL team for assistance.
//
// The following rules apply to this operation:
//
// <h3>
//   - Attitude Set numPoints value must correspond exactly to the number of Onorbit Attitude states associated with that Attitude Set.  The numPoints value is checked against the actual posted number of states and mismatch will result in the post being rejected.
//   - Attitude Set startTime and endTime must correspond to the earliest and latest state times, respectively, of the associated Onorbit Attitude states.
//   - Either satNo, idOnOrbit, or origObjectId must be provided.  The preferred option is to post with satNo for a cataloged object with (only) origObjectId for an unknown, uncatalogued, or internal/test object.  There are several cases for the logic associated with these fields:
//   - If satNo is provided and correlates to a known UDL sat number then the idOnOrbit will be populated appropriately in addition to the satNo.
//   - If satNo is provided and does not correlate to a known UDL sat number then the provided satNo value will be moved to the origObjectId field and satNo left null.
//   - If origObjectId and a valid satNo or idOnOrbit are provided then both the satNo/idOnOrbit and origObjectId will maintain the provided values.
//   - If only origObjectId is provided then origObjectId will be populated with the posted value.  In this case, no checks are made against existing UDL sat numbers.
//
// </h3>
func (r *AttitudeSetService) New(ctx context.Context, body AttitudeSetNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/attitudeset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AttitudeSetService) List(ctx context.Context, query AttitudeSetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AttitudesetAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/attitudeset"
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
func (r *AttitudeSetService) ListAutoPaging(ctx context.Context, query AttitudeSetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AttitudesetAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AttitudeSetService) Count(ctx context.Context, query AttitudeSetCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/attitudeset/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AttitudeSetService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/attitudeset/queryhelp"
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
func (r *AttitudeSetService) Tuple(ctx context.Context, query AttitudeSetTupleParams, opts ...option.RequestOption) (res *[]shared.AttitudesetFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/attitudeset/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a single Attitude Set and many associated Onorbit
// Attitude records as a POST body for ingest into the database. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
//
// The following rules apply to this operation:
//
// <h3>
//   - Attitude Set numPoints value must correspond exactly to the number of Onorbit Attitude states associated with that Attitude Set. The numPoints value is checked against the actual posted number of states and mismatch will result in the post being rejected.
//   - Attitude Set startTime and endTime must correspond to the earliest and latest state times, respectively, of the associated Onorbit Attitude states.
//   - Either satNo, idOnOrbit, or origObjectId must be provided. The preferred option is to post with satNo for a cataloged object, and with (only) origObjectId for an unknown, uncatalogued, or internal/test object.  There are several cases for the logic associated with these fields:
//   - If satNo is provided and correlates to a known UDL sat number then the idOnOrbit will be populated appropriately in addition to the satNo.
//   - If satNo is provided and does not correlate to a known UDL sat number then the provided satNo value will be moved to the origObjectId field and satNo left null.
//   - If origObjectId and a valid satNo or idOnOrbit are provided then both the satNo/idOnOrbit and origObjectId will maintain the provided values.
//   - If only origObjectId is provided then origObjectId will be populated with the posted value.  In this case, no checks are made against existing UDL sat numbers.
//
// </h3>
func (r *AttitudeSetService) UnvalidatedPublish(ctx context.Context, body AttitudeSetUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-attitudeset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// AttitudeSet represents a wrapper or collection of Onorbit Attitude 'points' and
// meta data indicating the specifics of the orientation of an on-orbit object.
// Attitude is typically distributed in a flat file containing details of the
// attitude generation as well as a large collection of individual points at
// varying time steps. AttitudeSet is analogous to this flat file.
type AttitudesetAbridged struct {
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
	DataMode AttitudesetAbridgedDataMode `json:"dataMode,required"`
	// The end time of the attitude ephemeris, in ISO 8601 UTC format, with microsecond
	// precision. If this set is constituted by a single epoch attitude message then
	// endTime should match the startTime.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Reference frame 1 of the quaternion or Euler angle transformation utilized in
	// this attitude parameter or attitude ephemeris. The UDL convention is that
	// transformations occur FROM frame1 TO frame2. A specific spacecraft frame or
	// instrument name may be provided with the assumption the consumer understands the
	// location of these frames (ex. SC BODY, J2000, LVLH, ICRF, INSTRUMENTx,
	// THRUSTERx, etc.).
	Frame1 string `json:"frame1,required"`
	// Reference frame 2 of the quaternion or Euler angle transformation utilized in
	// this attitude parameter or attitude ephemeris. The UDL convention is that
	// transformations occur FROM frame1 TO frame2. A specific spacecraft frame or
	// instrument name may be provided with the assumption the consumer understands the
	// location of these frames (ex. SC BODY, J2000, LVLH, ICRF, INSTRUMENTx,
	// THRUSTERx, etc.).
	Frame2 string `json:"frame2,required"`
	// Number of attitude records contained in this set.
	NumPoints int64 `json:"numPoints,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The epoch or start time of the attitude parameter or attitude ephemeris, in ISO
	// 8601 UTC format, with microsecond precision. If this set is constituted by a
	// single attitude parameter message then startTime is the epoch.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The type of attitude message or messages associated with this set.
	//
	// AEM = Attitude Ephemeris Message, specifying the attitude state of a single
	// object at multiple epochs.
	//
	// APM = Attitude Parameters Message, specifying the attitude state of a single
	// object at a single epoch.
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent (positional) Ephemeris Set, if this data is
	// correlated with an Ephemeris.
	EsID string `json:"esId"`
	// The rotation sequence of the Euler angles in which attitude reference frame
	// transformation occurs (from left to right). One, two, or three axis rotations
	// are supported and are represented by one, two, or three characters respectively.
	// Repeated axis rotations are also supported, however, these rotations should not
	// be sequential. The numeric sequence values correspond to the body angles/rates
	// as follows: 1 - xAngle/xRate, 2 - yAngle/yRate, and 3 - zAngle/zRate. Valid
	// sequences are: 123, 132, 213, 231, 312, 321, 121, 131, 212, 232, 313, 323, 12,
	// 13, 21, 23, 31, 32, 1, 2, and 3.
	//
	// The following represent examples of possible rotation sequences: A single
	// rotation about the Y-axis can be expressed as '2', a double rotation with X-Z
	// sequence can be expressed as '13', and a triple rotation with Z-X-Y sequence can
	// be expressed as '312'.
	EulerRotSeq string `json:"eulerRotSeq"`
	// Unique identifier of the on-orbit satellite to which this attitude set applies.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the sensor to which this attitude set applies IF this set
	// is reporting a single sensor orientation.
	IDSensor string `json:"idSensor"`
	// Recommended interpolation method for estimating attitude ephemeris data.
	Interpolator string `json:"interpolator"`
	// Recommended polynomial interpolation degree.
	InterpolatorDegree int64 `json:"interpolatorDegree"`
	// Optional notes/comments for this attitude set.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the record source to indicate the target object
	// of this attitude set. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the record source to indicate the sensor
	// identifier to which this attitude set applies IF this set is reporting a single
	// sensor orientation. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Initial precession angle (ECI J2000 frame) in degrees.
	PrecAngleInit float64 `json:"precAngleInit"`
	// Satellite/catalog number of the on-orbit object to which this attitude set
	// applies.
	SatNo int64 `json:"satNo"`
	// Initial spin angle (ECI J2000 frame) in degrees.
	SpinAngleInit float64 `json:"spinAngleInit"`
	// Attitude ephemeris step size, in seconds. This applies to Attitude Ephemeris
	// Messages (AEM) that employ a fixed step size.
	StepSize int64 `json:"stepSize"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EndTime               respjson.Field
		Frame1                respjson.Field
		Frame2                respjson.Field
		NumPoints             respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EsID                  respjson.Field
		EulerRotSeq           respjson.Field
		IDOnOrbit             respjson.Field
		IDSensor              respjson.Field
		Interpolator          respjson.Field
		InterpolatorDegree    respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		OrigSensorID          respjson.Field
		PrecAngleInit         respjson.Field
		SatNo                 respjson.Field
		SpinAngleInit         respjson.Field
		StepSize              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttitudesetAbridged) RawJSON() string { return r.JSON.raw }
func (r *AttitudesetAbridged) UnmarshalJSON(data []byte) error {
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
type AttitudesetAbridgedDataMode string

const (
	AttitudesetAbridgedDataModeReal      AttitudesetAbridgedDataMode = "REAL"
	AttitudesetAbridgedDataModeTest      AttitudesetAbridgedDataMode = "TEST"
	AttitudesetAbridgedDataModeSimulated AttitudesetAbridgedDataMode = "SIMULATED"
	AttitudesetAbridgedDataModeExercise  AttitudesetAbridgedDataMode = "EXERCISE"
)

type AttitudeSetNewParams struct {
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
	DataMode AttitudeSetNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The end time of the attitude ephemeris, in ISO 8601 UTC format, with microsecond
	// precision. If this set is constituted by a single epoch attitude message then
	// endTime should match the startTime.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Reference frame 1 of the quaternion or Euler angle transformation utilized in
	// this attitude parameter or attitude ephemeris. The UDL convention is that
	// transformations occur FROM frame1 TO frame2. A specific spacecraft frame or
	// instrument name may be provided with the assumption the consumer understands the
	// location of these frames (ex. SC BODY, J2000, LVLH, ICRF, INSTRUMENTx,
	// THRUSTERx, etc.).
	Frame1 string `json:"frame1,required"`
	// Reference frame 2 of the quaternion or Euler angle transformation utilized in
	// this attitude parameter or attitude ephemeris. The UDL convention is that
	// transformations occur FROM frame1 TO frame2. A specific spacecraft frame or
	// instrument name may be provided with the assumption the consumer understands the
	// location of these frames (ex. SC BODY, J2000, LVLH, ICRF, INSTRUMENTx,
	// THRUSTERx, etc.).
	Frame2 string `json:"frame2,required"`
	// Number of attitude records contained in this set.
	NumPoints int64 `json:"numPoints,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The epoch or start time of the attitude parameter or attitude ephemeris, in ISO
	// 8601 UTC format, with microsecond precision. If this set is constituted by a
	// single attitude parameter message then startTime is the epoch.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The type of attitude message or messages associated with this set.
	//
	// AEM = Attitude Ephemeris Message, specifying the attitude state of a single
	// object at multiple epochs.
	//
	// APM = Attitude Parameters Message, specifying the attitude state of a single
	// object at a single epoch.
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Unique identifier of the parent (positional) Ephemeris Set, if this data is
	// correlated with an Ephemeris.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// The rotation sequence of the Euler angles in which attitude reference frame
	// transformation occurs (from left to right). One, two, or three axis rotations
	// are supported and are represented by one, two, or three characters respectively.
	// Repeated axis rotations are also supported, however, these rotations should not
	// be sequential. The numeric sequence values correspond to the body angles/rates
	// as follows: 1 - xAngle/xRate, 2 - yAngle/yRate, and 3 - zAngle/zRate. Valid
	// sequences are: 123, 132, 213, 231, 312, 321, 121, 131, 212, 232, 313, 323, 12,
	// 13, 21, 23, 31, 32, 1, 2, and 3.
	//
	// The following represent examples of possible rotation sequences: A single
	// rotation about the Y-axis can be expressed as '2', a double rotation with X-Z
	// sequence can be expressed as '13', and a triple rotation with Z-X-Y sequence can
	// be expressed as '312'.
	EulerRotSeq param.Opt[string] `json:"eulerRotSeq,omitzero"`
	// Unique identifier of the sensor to which this attitude set applies IF this set
	// is reporting a single sensor orientation.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Recommended interpolation method for estimating attitude ephemeris data.
	Interpolator param.Opt[string] `json:"interpolator,omitzero"`
	// Recommended polynomial interpolation degree.
	InterpolatorDegree param.Opt[int64] `json:"interpolatorDegree,omitzero"`
	// Optional notes/comments for this attitude set.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source to indicate the target object
	// of this attitude set. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the record source to indicate the sensor
	// identifier to which this attitude set applies IF this set is reporting a single
	// sensor orientation. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Initial precession angle (ECI J2000 frame) in degrees.
	PrecAngleInit param.Opt[float64] `json:"precAngleInit,omitzero"`
	// Satellite/catalog number of the on-orbit object to which this attitude set
	// applies.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Initial spin angle (ECI J2000 frame) in degrees.
	SpinAngleInit param.Opt[float64] `json:"spinAngleInit,omitzero"`
	// Attitude ephemeris step size, in seconds. This applies to Attitude Ephemeris
	// Messages (AEM) that employ a fixed step size.
	StepSize param.Opt[int64] `json:"stepSize,omitzero"`
	// Array of UDL UUIDs of one or more AttitudeSet records associated with this set.
	// For example, a spacecraft Attitude Ephemeris Set might include a reference to an
	// Attitude Parameter Message defining the sensor to body frame transformation for
	// a sensor onboard the spacecraft, which allows for calculation of the sensor
	// orientation in frame2 of the attitude ephemeris.
	AsRef []string `json:"asRef,omitzero"`
	// Collection of attitude data associated with this Attitude Set.
	AttitudeList []AttitudeSetNewParamsAttitudeList `json:"attitudeList,omitzero"`
	paramObj
}

func (r AttitudeSetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AttitudeSetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttitudeSetNewParams) UnmarshalJSON(data []byte) error {
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
type AttitudeSetNewParamsDataMode string

const (
	AttitudeSetNewParamsDataModeReal      AttitudeSetNewParamsDataMode = "REAL"
	AttitudeSetNewParamsDataModeTest      AttitudeSetNewParamsDataMode = "TEST"
	AttitudeSetNewParamsDataModeSimulated AttitudeSetNewParamsDataMode = "SIMULATED"
	AttitudeSetNewParamsDataModeExercise  AttitudeSetNewParamsDataMode = "EXERCISE"
)

// These services provide operations for posting and querying attitude of on-orbit
// objects. Attitude describes the orientation of an object, which can be
// represented by quaternions or euler angles. The AttitudeSet ID (asId) identifies
// the 'AttitudeSet' record which contains details of the underlying data as well
// as a collection of attitude points. Points must be retrieved by first
// identifying a desired AttitudeSet and pulling its points by that AttitudeSet ID
// 'asId'.
//
// The properties ClassificationMarking, DataMode, Source, Ts are required.
type AttitudeSetNewParamsAttitudeList struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Time associated with this attitude record, in ISO 8601 UTC format, with
	// microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Unique identifier of the parent AttitudeSet associated with this record.
	AsID param.Opt[string] `json:"asId,omitzero"`
	// Coning angle in degrees.
	ConingAngle param.Opt[float64] `json:"coningAngle,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Precession axis declination (ECI J2000 frame) in degrees.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// Unique identifier of the on-orbit satellite to which this attitude record
	// applies.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Label specifying type of rotational motion of target.
	MotionType param.Opt[string] `json:"motionType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the record source to indicate the target object
	// of this attitude record. This may be an internal identifier and not necessarily
	// map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Precession period in seconds.
	PrecPeriod param.Opt[float64] `json:"precPeriod,omitzero"`
	// Quaternion vector component 1.
	Q1 param.Opt[float64] `json:"q1,omitzero"`
	// Derivative of quaternion vector component 1.
	Q1Dot param.Opt[float64] `json:"q1Dot,omitzero"`
	// Quaternion vector component 2.
	Q2 param.Opt[float64] `json:"q2,omitzero"`
	// Derivative of quaternion vector component 2.
	Q2Dot param.Opt[float64] `json:"q2Dot,omitzero"`
	// Quaternion vector component 3.
	Q3 param.Opt[float64] `json:"q3,omitzero"`
	// Derivative of quaternion vector component 3.
	Q3Dot param.Opt[float64] `json:"q3Dot,omitzero"`
	// Quaternion scalar component.
	Qc param.Opt[float64] `json:"qc,omitzero"`
	// Derivative of quaternion scalar component.
	QcDot param.Opt[float64] `json:"qcDot,omitzero"`
	// Precession axis right ascension (ECI J2000 frame) in degrees.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Satellite/catalog number of the on-orbit object to which this attitude record
	// applies.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Spin period in seconds.
	SpinPeriod param.Opt[float64] `json:"spinPeriod,omitzero"`
	// Array of X body rotation Euler angle(s), in degrees (-180 to 180). For repeated
	// axis rotations, the array elements should be placed in the order that the angles
	// apply in the sequence.
	XAngle []float64 `json:"xAngle,omitzero"`
	// Array of X body rotation rate(s), in degrees per second. For repeated axis
	// rotations, the array elements should be placed in the order that the rates apply
	// in the sequence. Attitude rates are expressed in frame1 with respect to frame2.
	XRate []float64 `json:"xRate,omitzero"`
	// Array of Y body rotation Euler angle(s), in degrees (-180 to 180). For repeated
	// axis rotations, the array elements should be placed in the order that the angles
	// apply in the sequence.
	YAngle []float64 `json:"yAngle,omitzero"`
	// Array of Y body rotation rate(s), in degrees per second. For repeated axis
	// rotations, the array elements should be placed in the order that the rates apply
	// in the sequence. Attitude rates are expressed in frame1 with respect to frame2.
	YRate []float64 `json:"yRate,omitzero"`
	// Array of Z body rotation Euler angle(s), in degrees (-180 to 180). For repeated
	// axis rotations, the array elements should be placed in the order that the angles
	// apply in the sequence.
	ZAngle []float64 `json:"zAngle,omitzero"`
	// Array of Z body rotation rate(s), in degrees per second. For repeated axis
	// rotations, the array elements should be placed in the order that the rates apply
	// in the sequence Attitude rates are expressed in frame1 with respect to frame2.
	ZRate []float64 `json:"zRate,omitzero"`
	paramObj
}

func (r AttitudeSetNewParamsAttitudeList) MarshalJSON() (data []byte, err error) {
	type shadow AttitudeSetNewParamsAttitudeList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttitudeSetNewParamsAttitudeList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttitudeSetNewParamsAttitudeList](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type AttitudeSetListParams struct {
	// The epoch or start time of the attitude parameter or attitude ephemeris, in ISO
	// 8601 UTC format, with microsecond precision. If this set is constituted by a
	// single attitude parameter message then startTime is the epoch.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AttitudeSetListParams]'s query parameters as `url.Values`.
func (r AttitudeSetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttitudeSetCountParams struct {
	// The epoch or start time of the attitude parameter or attitude ephemeris, in ISO
	// 8601 UTC format, with microsecond precision. If this set is constituted by a
	// single attitude parameter message then startTime is the epoch.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AttitudeSetCountParams]'s query parameters as `url.Values`.
func (r AttitudeSetCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttitudeSetTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The epoch or start time of the attitude parameter or attitude ephemeris, in ISO
	// 8601 UTC format, with microsecond precision. If this set is constituted by a
	// single attitude parameter message then startTime is the epoch.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AttitudeSetTupleParams]'s query parameters as `url.Values`.
func (r AttitudeSetTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttitudeSetUnvalidatedPublishParams struct {
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
	DataMode AttitudeSetUnvalidatedPublishParamsDataMode `json:"dataMode,omitzero,required"`
	// The end time of the attitude ephemeris, in ISO 8601 UTC format, with microsecond
	// precision. If this set is constituted by a single epoch attitude message then
	// endTime should match the startTime.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Reference frame 1 of the quaternion or Euler angle transformation utilized in
	// this attitude parameter or attitude ephemeris. The UDL convention is that
	// transformations occur FROM frame1 TO frame2. A specific spacecraft frame or
	// instrument name may be provided with the assumption the consumer understands the
	// location of these frames (ex. SC BODY, J2000, LVLH, ICRF, INSTRUMENTx,
	// THRUSTERx, etc.).
	Frame1 string `json:"frame1,required"`
	// Reference frame 2 of the quaternion or Euler angle transformation utilized in
	// this attitude parameter or attitude ephemeris. The UDL convention is that
	// transformations occur FROM frame1 TO frame2. A specific spacecraft frame or
	// instrument name may be provided with the assumption the consumer understands the
	// location of these frames (ex. SC BODY, J2000, LVLH, ICRF, INSTRUMENTx,
	// THRUSTERx, etc.).
	Frame2 string `json:"frame2,required"`
	// Number of attitude records contained in this set.
	NumPoints int64 `json:"numPoints,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The epoch or start time of the attitude parameter or attitude ephemeris, in ISO
	// 8601 UTC format, with microsecond precision. If this set is constituted by a
	// single attitude parameter message then startTime is the epoch.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The type of attitude message or messages associated with this set.
	//
	// AEM = Attitude Ephemeris Message, specifying the attitude state of a single
	// object at multiple epochs.
	//
	// APM = Attitude Parameters Message, specifying the attitude state of a single
	// object at a single epoch.
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Unique identifier of the parent (positional) Ephemeris Set, if this data is
	// correlated with an Ephemeris.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// The rotation sequence of the Euler angles in which attitude reference frame
	// transformation occurs (from left to right). One, two, or three axis rotations
	// are supported and are represented by one, two, or three characters respectively.
	// Repeated axis rotations are also supported, however, these rotations should not
	// be sequential. The numeric sequence values correspond to the body angles/rates
	// as follows: 1 - xAngle/xRate, 2 - yAngle/yRate, and 3 - zAngle/zRate. Valid
	// sequences are: 123, 132, 213, 231, 312, 321, 121, 131, 212, 232, 313, 323, 12,
	// 13, 21, 23, 31, 32, 1, 2, and 3.
	//
	// The following represent examples of possible rotation sequences: A single
	// rotation about the Y-axis can be expressed as '2', a double rotation with X-Z
	// sequence can be expressed as '13', and a triple rotation with Z-X-Y sequence can
	// be expressed as '312'.
	EulerRotSeq param.Opt[string] `json:"eulerRotSeq,omitzero"`
	// Unique identifier of the sensor to which this attitude set applies IF this set
	// is reporting a single sensor orientation.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Recommended interpolation method for estimating attitude ephemeris data.
	Interpolator param.Opt[string] `json:"interpolator,omitzero"`
	// Recommended polynomial interpolation degree.
	InterpolatorDegree param.Opt[int64] `json:"interpolatorDegree,omitzero"`
	// Optional notes/comments for this attitude set.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the record source to indicate the target object
	// of this attitude set. This may be an internal identifier and not necessarily map
	// to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the record source to indicate the sensor
	// identifier to which this attitude set applies IF this set is reporting a single
	// sensor orientation. This may be an internal identifier and not necessarily a
	// valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Initial precession angle (ECI J2000 frame) in degrees.
	PrecAngleInit param.Opt[float64] `json:"precAngleInit,omitzero"`
	// Satellite/catalog number of the on-orbit object to which this attitude set
	// applies.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Initial spin angle (ECI J2000 frame) in degrees.
	SpinAngleInit param.Opt[float64] `json:"spinAngleInit,omitzero"`
	// Attitude ephemeris step size, in seconds. This applies to Attitude Ephemeris
	// Messages (AEM) that employ a fixed step size.
	StepSize param.Opt[int64] `json:"stepSize,omitzero"`
	// Array of UDL UUIDs of one or more AttitudeSet records associated with this set.
	// For example, a spacecraft Attitude Ephemeris Set might include a reference to an
	// Attitude Parameter Message defining the sensor to body frame transformation for
	// a sensor onboard the spacecraft, which allows for calculation of the sensor
	// orientation in frame2 of the attitude ephemeris.
	AsRef []string `json:"asRef,omitzero"`
	// Collection of attitude data associated with this Attitude Set.
	AttitudeList []AttitudeSetUnvalidatedPublishParamsAttitudeList `json:"attitudeList,omitzero"`
	paramObj
}

func (r AttitudeSetUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	type shadow AttitudeSetUnvalidatedPublishParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttitudeSetUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
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
type AttitudeSetUnvalidatedPublishParamsDataMode string

const (
	AttitudeSetUnvalidatedPublishParamsDataModeReal      AttitudeSetUnvalidatedPublishParamsDataMode = "REAL"
	AttitudeSetUnvalidatedPublishParamsDataModeTest      AttitudeSetUnvalidatedPublishParamsDataMode = "TEST"
	AttitudeSetUnvalidatedPublishParamsDataModeSimulated AttitudeSetUnvalidatedPublishParamsDataMode = "SIMULATED"
	AttitudeSetUnvalidatedPublishParamsDataModeExercise  AttitudeSetUnvalidatedPublishParamsDataMode = "EXERCISE"
)

// These services provide operations for posting and querying attitude of on-orbit
// objects. Attitude describes the orientation of an object, which can be
// represented by quaternions or euler angles. The AttitudeSet ID (asId) identifies
// the 'AttitudeSet' record which contains details of the underlying data as well
// as a collection of attitude points. Points must be retrieved by first
// identifying a desired AttitudeSet and pulling its points by that AttitudeSet ID
// 'asId'.
//
// The properties ClassificationMarking, DataMode, Source, Ts are required.
type AttitudeSetUnvalidatedPublishParamsAttitudeList struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Time associated with this attitude record, in ISO 8601 UTC format, with
	// microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Unique identifier of the parent AttitudeSet associated with this record.
	AsID param.Opt[string] `json:"asId,omitzero"`
	// Coning angle in degrees.
	ConingAngle param.Opt[float64] `json:"coningAngle,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Precession axis declination (ECI J2000 frame) in degrees.
	Declination param.Opt[float64] `json:"declination,omitzero"`
	// Unique identifier of the on-orbit satellite to which this attitude record
	// applies.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Label specifying type of rotational motion of target.
	MotionType param.Opt[string] `json:"motionType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the record source to indicate the target object
	// of this attitude record. This may be an internal identifier and not necessarily
	// map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Precession period in seconds.
	PrecPeriod param.Opt[float64] `json:"precPeriod,omitzero"`
	// Quaternion vector component 1.
	Q1 param.Opt[float64] `json:"q1,omitzero"`
	// Derivative of quaternion vector component 1.
	Q1Dot param.Opt[float64] `json:"q1Dot,omitzero"`
	// Quaternion vector component 2.
	Q2 param.Opt[float64] `json:"q2,omitzero"`
	// Derivative of quaternion vector component 2.
	Q2Dot param.Opt[float64] `json:"q2Dot,omitzero"`
	// Quaternion vector component 3.
	Q3 param.Opt[float64] `json:"q3,omitzero"`
	// Derivative of quaternion vector component 3.
	Q3Dot param.Opt[float64] `json:"q3Dot,omitzero"`
	// Quaternion scalar component.
	Qc param.Opt[float64] `json:"qc,omitzero"`
	// Derivative of quaternion scalar component.
	QcDot param.Opt[float64] `json:"qcDot,omitzero"`
	// Precession axis right ascension (ECI J2000 frame) in degrees.
	Ra param.Opt[float64] `json:"ra,omitzero"`
	// Satellite/catalog number of the on-orbit object to which this attitude record
	// applies.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Spin period in seconds.
	SpinPeriod param.Opt[float64] `json:"spinPeriod,omitzero"`
	// Array of X body rotation Euler angle(s), in degrees (-180 to 180). For repeated
	// axis rotations, the array elements should be placed in the order that the angles
	// apply in the sequence.
	XAngle []float64 `json:"xAngle,omitzero"`
	// Array of X body rotation rate(s), in degrees per second. For repeated axis
	// rotations, the array elements should be placed in the order that the rates apply
	// in the sequence. Attitude rates are expressed in frame1 with respect to frame2.
	XRate []float64 `json:"xRate,omitzero"`
	// Array of Y body rotation Euler angle(s), in degrees (-180 to 180). For repeated
	// axis rotations, the array elements should be placed in the order that the angles
	// apply in the sequence.
	YAngle []float64 `json:"yAngle,omitzero"`
	// Array of Y body rotation rate(s), in degrees per second. For repeated axis
	// rotations, the array elements should be placed in the order that the rates apply
	// in the sequence. Attitude rates are expressed in frame1 with respect to frame2.
	YRate []float64 `json:"yRate,omitzero"`
	// Array of Z body rotation Euler angle(s), in degrees (-180 to 180). For repeated
	// axis rotations, the array elements should be placed in the order that the angles
	// apply in the sequence.
	ZAngle []float64 `json:"zAngle,omitzero"`
	// Array of Z body rotation rate(s), in degrees per second. For repeated axis
	// rotations, the array elements should be placed in the order that the rates apply
	// in the sequence Attitude rates are expressed in frame1 with respect to frame2.
	ZRate []float64 `json:"zRate,omitzero"`
	paramObj
}

func (r AttitudeSetUnvalidatedPublishParamsAttitudeList) MarshalJSON() (data []byte, err error) {
	type shadow AttitudeSetUnvalidatedPublishParamsAttitudeList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttitudeSetUnvalidatedPublishParamsAttitudeList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttitudeSetUnvalidatedPublishParamsAttitudeList](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
