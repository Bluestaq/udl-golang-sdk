// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
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

// GnssObservationsetService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGnssObservationsetService] method instead.
type GnssObservationsetService struct {
	Options []option.RequestOption
	History GnssObservationsetHistoryService
}

// NewGnssObservationsetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGnssObservationsetService(opts ...option.RequestOption) (r GnssObservationsetService) {
	r = GnssObservationsetService{}
	r.Options = opts
	r.History = NewGnssObservationsetHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *GnssObservationsetService) List(ctx context.Context, query GnssObservationsetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GnssObservationsetListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/gnssobservationset"
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
func (r *GnssObservationsetService) ListAutoPaging(ctx context.Context, query GnssObservationsetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GnssObservationsetListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *GnssObservationsetService) Count(ctx context.Context, query GnssObservationsetCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/gnssobservationset/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of Track
// Details records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *GnssObservationsetService) NewBulk(ctx context.Context, body GnssObservationsetNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/gnssobservationset/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *GnssObservationsetService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *GnssObservationsetQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/gnssobservationset/queryhelp"
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
func (r *GnssObservationsetService) Tuple(ctx context.Context, query GnssObservationsetTupleParams, opts ...option.RequestOption) (res *[]GnssObservationSetFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/gnssobservationset/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to accept one or more GNSSObservationSet(s) and associated
// GNSS Observation(s) as a POST body and ingest into the database. This operation
// is intended to be used for automated feeds into UDL. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *GnssObservationsetService) UnvalidatedPublish(ctx context.Context, body GnssObservationsetUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-gnssobset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Set of GNSSObservation data.
type GnssObservationsetListResponse struct {
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
	DataMode GnssObservationsetListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Observation Time, in ISO8601 UTC format with microsecond precision. This
	// timestamp applies to all observations within the set.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// GNSS Automatic Gain Control State.
	AgcState int64 `json:"agcState"`
	// Spacecraft altitude at observation time (ts), expressed in kilometers above
	// WGS-84 ellipsoid.
	Alt float64 `json:"alt"`
	// unit vector of the outward facing direction of the receiver boresight in a
	// body-fixed coordinate system.
	Boresight []float64 `json:"boresight"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent Ephemeris Set, if this data is correlated with
	// an Ephemeris. If reporting for a spacecraft with multiple onboard GNSS
	// receivers, this ID may be associated with multiple GNSS Observation records if
	// each receiver is synced to the ephemeris points.
	EsID string `json:"esId"`
	// Optional source-provided identifier for this collection event. This field can be
	// used to associate records related to the same event.
	EventID string `json:"eventId"`
	// Geometric Dilution of Precision.
	GDop float64 `json:"gDop"`
	// GNSSObservations associated with this GNSSObservationSet.
	GnssObservationList []GnssObservationsetListResponseGnssObservationList `json:"gnssObservationList"`
	// Horizontal Dilution of Precision.
	HDop float64 `json:"hDop"`
	// Unique identifier of the primary satellite on-orbit object.
	IDOnOrbit string `json:"idOnOrbit"`
	// WGS-84 spacecraft latitude sub-point at observation time (ts), represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat float64 `json:"lat"`
	// WGS-84 spacecraft longitude sub-point at observation time (ts), represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The marker type of the observing receiver (AIRBORNE, ANIMAL, BALLISTIC, FIXED
	// BUOY, FLOATING BUOY, FLOATING ICE, GEODETIC, GLACIER, GROUNDCRAFT, HUMAN, NON
	// GEODETIC, NON PHYSICAL, SPACEBORNE, WATERCRAFT). Reference RINEX 3+ for further
	// information concerning marker types.
	MarkerType string `json:"markerType"`
	// The current navigation status as defined by the data source. In general the
	// navigation status specifies whether the signal is normal, degraded, or
	// unavailable. For status value definitions please reach out to data source
	// provider.
	NavigationStatus string `json:"navigationStatus"`
	// Array of the strings containing the individual observation code sets that are
	// contained within this GNSS Observation set. Each string is a three-character
	// representation of the measurement type, the channel, and the coding, in
	// compliance with the RINEX 3+ standard (Pseudorange (C), Carrier Phase (L),
	// Doppler (D), Signal Strength C/No (S), or Channel Number (X)). See the GNSS
	// Observation ob field for the units of measure associated with each observation
	// type.
	ObsCodes []string `json:"obsCodes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by ephemeris source to indicate the target object
	// of this ephemeris. This may be an internal identifier and not necessarily map to
	// a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Time, in seconds, that the receiver has been unable to compute a GNSS fix.
	Outage int64 `json:"outage"`
	// Position Dilution of Precision.
	PDop float64 `json:"pDop"`
	// The quaternion describing the rotation of the body-fixed frame used for this
	// system into the local geodetic frame, at observation time (ts). The array
	// element order convention is scalar component first, followed by the three vector
	// components. For a vector u in the body-fixed frame, the corresponding vector u'
	// in the geodetic frame should satisfy u' = quq\*, where q is this quaternion.
	Quat []float64 `json:"quat"`
	// The number or ID of the GNSS receiver associated with this data. Each GNSS
	// Observation Set is associated with only one receiver. If reporting for multiple
	// receivers a separate set should be generated for each. A null value is assumed
	// to indicate that only one receiver is present, or reported.
	Receiver string `json:"receiver"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo int64 `json:"satNo"`
	// Three element array, expressing the ECI J2K spacecraft position components, in
	// kilometers, at observation time (ts). The array element order is [x, y, z].
	SatPosition []float64 `json:"satPosition"`
	// Three element array, expressing the ECI J2K spacecraft velocity components, in
	// km/second, at observation time (ts). The array element order is [xvel, yvel,
	// zvel].
	SatVelocity []float64 `json:"satVelocity"`
	// Array of UUIDs of the UDL data records that are related to this GNSS Observation
	// Set. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object (e.g. /udl/statevector/{uuid}).
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK) that
	// are related to this GNSS Observation Set. See the associated 'srcIds' array for
	// the record UUIDs, positionally corresponding to the record types in this array.
	// The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Time Dilution of Precision.
	TDop float64 `json:"tDop"`
	// Status of the GNSS receiver signal. Status options are 0, 1 or 2 (0 being the
	// best).
	TrackingStatus int64 `json:"trackingStatus"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Vertical Dilution of Precision.
	VDop float64 `json:"vDop"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		AgcState              respjson.Field
		Alt                   respjson.Field
		Boresight             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EsID                  respjson.Field
		EventID               respjson.Field
		GDop                  respjson.Field
		GnssObservationList   respjson.Field
		HDop                  respjson.Field
		IDOnOrbit             respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		MarkerType            respjson.Field
		NavigationStatus      respjson.Field
		ObsCodes              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		Outage                respjson.Field
		PDop                  respjson.Field
		Quat                  respjson.Field
		Receiver              respjson.Field
		SatNo                 respjson.Field
		SatPosition           respjson.Field
		SatVelocity           respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		Tags                  respjson.Field
		TDop                  respjson.Field
		TrackingStatus        respjson.Field
		TransactionID         respjson.Field
		VDop                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GnssObservationsetListResponse) RawJSON() string { return r.JSON.raw }
func (r *GnssObservationsetListResponse) UnmarshalJSON(data []byte) error {
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
type GnssObservationsetListResponseDataMode string

const (
	GnssObservationsetListResponseDataModeReal      GnssObservationsetListResponseDataMode = "REAL"
	GnssObservationsetListResponseDataModeTest      GnssObservationsetListResponseDataMode = "TEST"
	GnssObservationsetListResponseDataModeSimulated GnssObservationsetListResponseDataMode = "SIMULATED"
	GnssObservationsetListResponseDataModeExercise  GnssObservationsetListResponseDataMode = "EXERCISE"
)

// Information for Global Navigation Satellite Systems (GNSS) Observations
// collected from GNSS receivers, including the specific GNSS sat from which each
// signal was received, and the observation codes of each observation in the
// record. Each GNSS Observation is associated with a GNSS Observation Set record
// containing data which applies to all observations in the set, including
// observation time, receiver location, and Dilution of Precision (DOP) values.
// Users can Reference RINEX 3+ documentation for further information concerning
// many of the standards and conventions for GNSS observations.
type GnssObservationsetListResponseGnssObservationList struct {
	// GNSS Automatic Gain Control State.
	AgcState int64 `json:"agcState"`
	// RINEX 3+ compliant GNSS System and Satellite Identifier (represented as SNN,
	// where S is the system code, and NN is the satellite identifier) associated with
	// this observation:
	//
	// G - GPS (NN = PRN)
	//
	// R - GLONASS (NN = Slot Number)
	//
	// S - SBAS Payload (NN = PRN-100)
	//
	// E - Galileo (NN = PRN)
	//
	// C - BeiDou (NN = PRN)
	//
	// J - QZSS (NN = PRN - 1923)
	//
	// I - IRNSS (NN = PRN)
	GnssSatID string `json:"gnssSatId"`
	// Array of observation(s). The ob array must be the same length as the obsCodeSet.
	// Pseudorange (C) is expressed meters, carrier phase (L) in cycles, doppler (D) in
	// Hz where + values indicate approaching sats, and signal strength C/No (S) in
	// dB-Hz.
	Ob []float64 `json:"ob"`
	// The observation code set that applies to this observation record. Reference
	// RINEX 3+ for further information concerning observation code set conventions.
	ObsCodeSet []string `json:"obsCodeSet"`
	// Status of the GNSS receiver signal. Status options are 0, 1 or 2 (0 being the
	// best).
	TrackingStatus int64 `json:"trackingStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgcState       respjson.Field
		GnssSatID      respjson.Field
		Ob             respjson.Field
		ObsCodeSet     respjson.Field
		TrackingStatus respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GnssObservationsetListResponseGnssObservationList) RawJSON() string { return r.JSON.raw }
func (r *GnssObservationsetListResponseGnssObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GnssObservationsetQueryhelpResponse struct {
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
func (r GnssObservationsetQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *GnssObservationsetQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GnssObservationsetListParams struct {
	// Observation Time, in ISO8601 UTC format with microsecond precision. This
	// timestamp applies to all observations within the set.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GnssObservationsetListParams]'s query parameters as
// `url.Values`.
func (r GnssObservationsetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GnssObservationsetCountParams struct {
	// Observation Time, in ISO8601 UTC format with microsecond precision. This
	// timestamp applies to all observations within the set.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GnssObservationsetCountParams]'s query parameters as
// `url.Values`.
func (r GnssObservationsetCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GnssObservationsetNewBulkParams struct {
	Body []GnssObservationsetNewBulkParamsBody
	paramObj
}

func (r GnssObservationsetNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *GnssObservationsetNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Set of GNSSObservation data.
//
// The properties ClassificationMarking, DataMode, Source, Ts are required.
type GnssObservationsetNewBulkParamsBody struct {
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
	// Observation Time, in ISO8601 UTC format with microsecond precision. This
	// timestamp applies to all observations within the set.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// GNSS Automatic Gain Control State.
	AgcState param.Opt[int64] `json:"agcState,omitzero"`
	// Spacecraft altitude at observation time (ts), expressed in kilometers above
	// WGS-84 ellipsoid.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of the parent Ephemeris Set, if this data is correlated with
	// an Ephemeris. If reporting for a spacecraft with multiple onboard GNSS
	// receivers, this ID may be associated with multiple GNSS Observation records if
	// each receiver is synced to the ephemeris points.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// Optional source-provided identifier for this collection event. This field can be
	// used to associate records related to the same event.
	EventID param.Opt[string] `json:"eventId,omitzero"`
	// Geometric Dilution of Precision.
	GDop param.Opt[float64] `json:"gDop,omitzero"`
	// Horizontal Dilution of Precision.
	HDop param.Opt[float64] `json:"hDop,omitzero"`
	// Unique identifier of the primary satellite on-orbit object.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// WGS-84 spacecraft latitude sub-point at observation time (ts), represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 spacecraft longitude sub-point at observation time (ts), represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The marker type of the observing receiver (AIRBORNE, ANIMAL, BALLISTIC, FIXED
	// BUOY, FLOATING BUOY, FLOATING ICE, GEODETIC, GLACIER, GROUNDCRAFT, HUMAN, NON
	// GEODETIC, NON PHYSICAL, SPACEBORNE, WATERCRAFT). Reference RINEX 3+ for further
	// information concerning marker types.
	MarkerType param.Opt[string] `json:"markerType,omitzero"`
	// The current navigation status as defined by the data source. In general the
	// navigation status specifies whether the signal is normal, degraded, or
	// unavailable. For status value definitions please reach out to data source
	// provider.
	NavigationStatus param.Opt[string] `json:"navigationStatus,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by ephemeris source to indicate the target object
	// of this ephemeris. This may be an internal identifier and not necessarily map to
	// a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Time, in seconds, that the receiver has been unable to compute a GNSS fix.
	Outage param.Opt[int64] `json:"outage,omitzero"`
	// Position Dilution of Precision.
	PDop param.Opt[float64] `json:"pDop,omitzero"`
	// The number or ID of the GNSS receiver associated with this data. Each GNSS
	// Observation Set is associated with only one receiver. If reporting for multiple
	// receivers a separate set should be generated for each. A null value is assumed
	// to indicate that only one receiver is present, or reported.
	Receiver param.Opt[string] `json:"receiver,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Time Dilution of Precision.
	TDop param.Opt[float64] `json:"tDop,omitzero"`
	// Status of the GNSS receiver signal. Status options are 0, 1 or 2 (0 being the
	// best).
	TrackingStatus param.Opt[int64] `json:"trackingStatus,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Vertical Dilution of Precision.
	VDop param.Opt[float64] `json:"vDop,omitzero"`
	// unit vector of the outward facing direction of the receiver boresight in a
	// body-fixed coordinate system.
	Boresight []float64 `json:"boresight,omitzero"`
	// GNSSObservations associated with this GNSSObservationSet.
	GnssObservationList []GnssObservationsetNewBulkParamsBodyGnssObservationList `json:"gnssObservationList,omitzero"`
	// Array of the strings containing the individual observation code sets that are
	// contained within this GNSS Observation set. Each string is a three-character
	// representation of the measurement type, the channel, and the coding, in
	// compliance with the RINEX 3+ standard (Pseudorange (C), Carrier Phase (L),
	// Doppler (D), Signal Strength C/No (S), or Channel Number (X)). See the GNSS
	// Observation ob field for the units of measure associated with each observation
	// type.
	ObsCodes []string `json:"obsCodes,omitzero"`
	// The quaternion describing the rotation of the body-fixed frame used for this
	// system into the local geodetic frame, at observation time (ts). The array
	// element order convention is scalar component first, followed by the three vector
	// components. For a vector u in the body-fixed frame, the corresponding vector u'
	// in the geodetic frame should satisfy u' = quq\*, where q is this quaternion.
	Quat []float64 `json:"quat,omitzero"`
	// Three element array, expressing the ECI J2K spacecraft position components, in
	// kilometers, at observation time (ts). The array element order is [x, y, z].
	SatPosition []float64 `json:"satPosition,omitzero"`
	// Three element array, expressing the ECI J2K spacecraft velocity components, in
	// km/second, at observation time (ts). The array element order is [xvel, yvel,
	// zvel].
	SatVelocity []float64 `json:"satVelocity,omitzero"`
	// Array of UUIDs of the UDL data records that are related to this GNSS Observation
	// Set. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object (e.g. /udl/statevector/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK) that
	// are related to this GNSS Observation Set. See the associated 'srcIds' array for
	// the record UUIDs, positionally corresponding to the record types in this array.
	// The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r GnssObservationsetNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow GnssObservationsetNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GnssObservationsetNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GnssObservationsetNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Information for Global Navigation Satellite Systems (GNSS) Observations
// collected from GNSS receivers, including the specific GNSS sat from which each
// signal was received, and the observation codes of each observation in the
// record. Each GNSS Observation is associated with a GNSS Observation Set record
// containing data which applies to all observations in the set, including
// observation time, receiver location, and Dilution of Precision (DOP) values.
// Users can Reference RINEX 3+ documentation for further information concerning
// many of the standards and conventions for GNSS observations.
type GnssObservationsetNewBulkParamsBodyGnssObservationList struct {
	// GNSS Automatic Gain Control State.
	AgcState param.Opt[int64] `json:"agcState,omitzero"`
	// RINEX 3+ compliant GNSS System and Satellite Identifier (represented as SNN,
	// where S is the system code, and NN is the satellite identifier) associated with
	// this observation:
	//
	// G - GPS (NN = PRN)
	//
	// R - GLONASS (NN = Slot Number)
	//
	// S - SBAS Payload (NN = PRN-100)
	//
	// E - Galileo (NN = PRN)
	//
	// C - BeiDou (NN = PRN)
	//
	// J - QZSS (NN = PRN - 1923)
	//
	// I - IRNSS (NN = PRN)
	GnssSatID param.Opt[string] `json:"gnssSatId,omitzero"`
	// Status of the GNSS receiver signal. Status options are 0, 1 or 2 (0 being the
	// best).
	TrackingStatus param.Opt[int64] `json:"trackingStatus,omitzero"`
	// Array of observation(s). The ob array must be the same length as the obsCodeSet.
	// Pseudorange (C) is expressed meters, carrier phase (L) in cycles, doppler (D) in
	// Hz where + values indicate approaching sats, and signal strength C/No (S) in
	// dB-Hz.
	Ob []float64 `json:"ob,omitzero"`
	// The observation code set that applies to this observation record. Reference
	// RINEX 3+ for further information concerning observation code set conventions.
	ObsCodeSet []string `json:"obsCodeSet,omitzero"`
	paramObj
}

func (r GnssObservationsetNewBulkParamsBodyGnssObservationList) MarshalJSON() (data []byte, err error) {
	type shadow GnssObservationsetNewBulkParamsBodyGnssObservationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GnssObservationsetNewBulkParamsBodyGnssObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GnssObservationsetTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Observation Time, in ISO8601 UTC format with microsecond precision. This
	// timestamp applies to all observations within the set.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GnssObservationsetTupleParams]'s query parameters as
// `url.Values`.
func (r GnssObservationsetTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GnssObservationsetUnvalidatedPublishParams struct {
	Body []GnssObservationsetUnvalidatedPublishParamsBody
	paramObj
}

func (r GnssObservationsetUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *GnssObservationsetUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Set of GNSSObservation data.
//
// The properties ClassificationMarking, DataMode, Source, Ts are required.
type GnssObservationsetUnvalidatedPublishParamsBody struct {
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
	// Observation Time, in ISO8601 UTC format with microsecond precision. This
	// timestamp applies to all observations within the set.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// GNSS Automatic Gain Control State.
	AgcState param.Opt[int64] `json:"agcState,omitzero"`
	// Spacecraft altitude at observation time (ts), expressed in kilometers above
	// WGS-84 ellipsoid.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Unique identifier of the parent Ephemeris Set, if this data is correlated with
	// an Ephemeris. If reporting for a spacecraft with multiple onboard GNSS
	// receivers, this ID may be associated with multiple GNSS Observation records if
	// each receiver is synced to the ephemeris points.
	EsID param.Opt[string] `json:"esId,omitzero"`
	// Optional source-provided identifier for this collection event. This field can be
	// used to associate records related to the same event.
	EventID param.Opt[string] `json:"eventId,omitzero"`
	// Geometric Dilution of Precision.
	GDop param.Opt[float64] `json:"gDop,omitzero"`
	// Horizontal Dilution of Precision.
	HDop param.Opt[float64] `json:"hDop,omitzero"`
	// Unique identifier of the primary satellite on-orbit object.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// WGS-84 spacecraft latitude sub-point at observation time (ts), represented as
	// -90 to 90 degrees (negative values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 spacecraft longitude sub-point at observation time (ts), represented as
	// -180 to 180 degrees (negative values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The marker type of the observing receiver (AIRBORNE, ANIMAL, BALLISTIC, FIXED
	// BUOY, FLOATING BUOY, FLOATING ICE, GEODETIC, GLACIER, GROUNDCRAFT, HUMAN, NON
	// GEODETIC, NON PHYSICAL, SPACEBORNE, WATERCRAFT). Reference RINEX 3+ for further
	// information concerning marker types.
	MarkerType param.Opt[string] `json:"markerType,omitzero"`
	// The current navigation status as defined by the data source. In general the
	// navigation status specifies whether the signal is normal, degraded, or
	// unavailable. For status value definitions please reach out to data source
	// provider.
	NavigationStatus param.Opt[string] `json:"navigationStatus,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by ephemeris source to indicate the target object
	// of this ephemeris. This may be an internal identifier and not necessarily map to
	// a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Time, in seconds, that the receiver has been unable to compute a GNSS fix.
	Outage param.Opt[int64] `json:"outage,omitzero"`
	// Position Dilution of Precision.
	PDop param.Opt[float64] `json:"pDop,omitzero"`
	// The number or ID of the GNSS receiver associated with this data. Each GNSS
	// Observation Set is associated with only one receiver. If reporting for multiple
	// receivers a separate set should be generated for each. A null value is assumed
	// to indicate that only one receiver is present, or reported.
	Receiver param.Opt[string] `json:"receiver,omitzero"`
	// Satellite/catalog number of the target on-orbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Time Dilution of Precision.
	TDop param.Opt[float64] `json:"tDop,omitzero"`
	// Status of the GNSS receiver signal. Status options are 0, 1 or 2 (0 being the
	// best).
	TrackingStatus param.Opt[int64] `json:"trackingStatus,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Vertical Dilution of Precision.
	VDop param.Opt[float64] `json:"vDop,omitzero"`
	// unit vector of the outward facing direction of the receiver boresight in a
	// body-fixed coordinate system.
	Boresight []float64 `json:"boresight,omitzero"`
	// GNSSObservations associated with this GNSSObservationSet.
	GnssObservationList []GnssObservationsetUnvalidatedPublishParamsBodyGnssObservationList `json:"gnssObservationList,omitzero"`
	// Array of the strings containing the individual observation code sets that are
	// contained within this GNSS Observation set. Each string is a three-character
	// representation of the measurement type, the channel, and the coding, in
	// compliance with the RINEX 3+ standard (Pseudorange (C), Carrier Phase (L),
	// Doppler (D), Signal Strength C/No (S), or Channel Number (X)). See the GNSS
	// Observation ob field for the units of measure associated with each observation
	// type.
	ObsCodes []string `json:"obsCodes,omitzero"`
	// The quaternion describing the rotation of the body-fixed frame used for this
	// system into the local geodetic frame, at observation time (ts). The array
	// element order convention is scalar component first, followed by the three vector
	// components. For a vector u in the body-fixed frame, the corresponding vector u'
	// in the geodetic frame should satisfy u' = quq\*, where q is this quaternion.
	Quat []float64 `json:"quat,omitzero"`
	// Three element array, expressing the ECI J2K spacecraft position components, in
	// kilometers, at observation time (ts). The array element order is [x, y, z].
	SatPosition []float64 `json:"satPosition,omitzero"`
	// Three element array, expressing the ECI J2K spacecraft velocity components, in
	// km/second, at observation time (ts). The array element order is [xvel, yvel,
	// zvel].
	SatVelocity []float64 `json:"satVelocity,omitzero"`
	// Array of UUIDs of the UDL data records that are related to this GNSS Observation
	// Set. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object (e.g. /udl/statevector/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (AIS, CONJUNCTION, DOA, ELSET, EO, ESID, GROUNDIMAGE,
	// POI, MANEUVER, MTI, NOTIFICATION, RADAR, RF, SIGACT, SKYIMAGE, SV, TRACK) that
	// are related to this GNSS Observation Set. See the associated 'srcIds' array for
	// the record UUIDs, positionally corresponding to the record types in this array.
	// The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r GnssObservationsetUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow GnssObservationsetUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GnssObservationsetUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GnssObservationsetUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Information for Global Navigation Satellite Systems (GNSS) Observations
// collected from GNSS receivers, including the specific GNSS sat from which each
// signal was received, and the observation codes of each observation in the
// record. Each GNSS Observation is associated with a GNSS Observation Set record
// containing data which applies to all observations in the set, including
// observation time, receiver location, and Dilution of Precision (DOP) values.
// Users can Reference RINEX 3+ documentation for further information concerning
// many of the standards and conventions for GNSS observations.
type GnssObservationsetUnvalidatedPublishParamsBodyGnssObservationList struct {
	// GNSS Automatic Gain Control State.
	AgcState param.Opt[int64] `json:"agcState,omitzero"`
	// RINEX 3+ compliant GNSS System and Satellite Identifier (represented as SNN,
	// where S is the system code, and NN is the satellite identifier) associated with
	// this observation:
	//
	// G - GPS (NN = PRN)
	//
	// R - GLONASS (NN = Slot Number)
	//
	// S - SBAS Payload (NN = PRN-100)
	//
	// E - Galileo (NN = PRN)
	//
	// C - BeiDou (NN = PRN)
	//
	// J - QZSS (NN = PRN - 1923)
	//
	// I - IRNSS (NN = PRN)
	GnssSatID param.Opt[string] `json:"gnssSatId,omitzero"`
	// Status of the GNSS receiver signal. Status options are 0, 1 or 2 (0 being the
	// best).
	TrackingStatus param.Opt[int64] `json:"trackingStatus,omitzero"`
	// Array of observation(s). The ob array must be the same length as the obsCodeSet.
	// Pseudorange (C) is expressed meters, carrier phase (L) in cycles, doppler (D) in
	// Hz where + values indicate approaching sats, and signal strength C/No (S) in
	// dB-Hz.
	Ob []float64 `json:"ob,omitzero"`
	// The observation code set that applies to this observation record. Reference
	// RINEX 3+ for further information concerning observation code set conventions.
	ObsCodeSet []string `json:"obsCodeSet,omitzero"`
	paramObj
}

func (r GnssObservationsetUnvalidatedPublishParamsBodyGnssObservationList) MarshalJSON() (data []byte, err error) {
	type shadow GnssObservationsetUnvalidatedPublishParamsBodyGnssObservationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GnssObservationsetUnvalidatedPublishParamsBodyGnssObservationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
