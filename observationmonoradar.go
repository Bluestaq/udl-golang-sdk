// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// ObservationMonoradarService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationMonoradarService] method instead.
type ObservationMonoradarService struct {
	Options []option.RequestOption
	History ObservationMonoradarHistoryService
}

// NewObservationMonoradarService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObservationMonoradarService(opts ...option.RequestOption) (r ObservationMonoradarService) {
	r = ObservationMonoradarService{}
	r.Options = opts
	r.History = NewObservationMonoradarHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ObservationMonoradarService) List(ctx context.Context, query ObservationMonoradarListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ObservationMonoradarListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/monoradar"
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
func (r *ObservationMonoradarService) ListAutoPaging(ctx context.Context, query ObservationMonoradarListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ObservationMonoradarListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ObservationMonoradarService) Count(ctx context.Context, query ObservationMonoradarCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/monoradar/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// MonoRadar records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *ObservationMonoradarService) NewBulk(ctx context.Context, body ObservationMonoradarNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/monoradar/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ObservationMonoradarService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *ObservationMonoradarQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/monoradar/queryhelp"
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
func (r *ObservationMonoradarService) Tuple(ctx context.Context, query ObservationMonoradarTupleParams, opts ...option.RequestOption) (res *[]ObservationMonoradarTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/monoradar/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a list of MonoRadar records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *ObservationMonoradarService) UnvalidatedPublish(ctx context.Context, body ObservationMonoradarUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/monoradar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// A monoradar record contains the raw, and in some cases, processed target reports
// from primary and secondary air surveillance radars. All target positions for
// monoradar reports are recorded as range and azimuth from geographical North
// relative to the detecting radar site. In the case of secondary surveillance
// radars, interrogation response codes are provided as well as quality and
// validation characteristics, when available in the particular record type used to
// generate the record.
type ObservationMonoradarListResponse struct {
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
	DataMode ObservationMonoradarListResponseDataMode `json:"dataMode,required"`
	// Message format received (i.e. 'ASR9', 'CAT48', 'TPS70', etc..).
	Msgfmt string `json:"msgfmt,required"`
	// Message time, in ISO 8601 UTC format with microsecond precision. This is the
	// time that the data message was released from the site.
	Msgts time.Time `json:"msgts,required" format:"date-time"`
	// Message report type received (i.e. 'SRCH', 'BCN', 'REINF', 'BRTQC', 'PSR',
	// etc..).
	Msgtyp string `json:"msgtyp,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Target detection time, in ISO 8601 UTC format with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Azimuth Change Pulse count at target detection.
	Acp int64 `json:"acp"`
	// Aircraft address (24-bits Mode S address) assigned uniquely to each aircraft.
	Addr string `json:"addr"`
	// Flag indicating military message.
	Af bool `json:"af"`
	// Flag indicating whether AIMS present.
	Aims bool `json:"aims"`
	// Measured height of the target, in km. (for 3D radars).
	Alt3d float64 `json:"alt3d"`
	// ARTS quality.
	Artsqual string `json:"artsqual"`
	// Target azimuth, measured from the observing site, in degrees from true North. If
	// Azimuth Change Pulse (acp) count is provided, az represents the computed angle.
	Az float64 `json:"az"`
	// Target azimuth delta between PSR and SSR (reference PSR-SSR), in degrees.
	Azdelt float64 `json:"azdelt"`
	// Number of beacon hits received on the target.
	Bcnhits int64 `json:"bcnhits"`
	// Array of local 2d-cartesian [x, y] coordinates of target, in km.
	Cartpos []float64 `json:"cartpos"`
	// Climbing/Descending mode indicator.
	Cdm string `json:"cdm"`
	// 7500 squawk present (hijack).
	Code7500 bool `json:"code7500"`
	// 7600 squawk present (loss of comm).
	Code7600 bool `json:"code7600"`
	// 7700 squawk present (general emergency).
	Code7700 bool `json:"code7700"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Flag indicating FAA/Civ message.
	Faa bool `json:"faa"`
	// Target ground speed, in meters/second.
	Grndspd float64 `json:"grndspd"`
	// Target heading, in degrees from true North at the target position.
	Hdng float64 `json:"hdng"`
	// ID of the radar site or station providing the data.
	IDSensor string `json:"idSensor"`
	// Mode-1 interrogation response (mission code).
	M1 string `json:"m1"`
	// Indicator that the mode-1 response was garbled.
	M1g bool `json:"m1g"`
	// Status of the mode-1 validity bit.
	M1v string `json:"m1v"`
	// Mode-2 interrogation response (military identification code).
	M2 string `json:"m2"`
	// Indicator that the mode-2 response was garbled.
	M2g bool `json:"m2g"`
	// Status of the mode-2 validity bit.
	M2v string `json:"m2v"`
	// Status of the mode-2 X-Pulse response validation.
	M2xv string `json:"m2xv"`
	// Mode-3/A interrogation response (aircraft identification).
	M3a string `json:"m3a"`
	// Indicator that the mode-3/A response was garbled.
	M3ag bool `json:"m3ag"`
	// Status of the mode-3/A validity bit.
	M3av string `json:"m3av"`
	// Status of the mode-3 X-Pulse response validation.
	M3axv string `json:"m3axv"`
	// Mode-4 interrogation response (Identification Friend/Foe).
	M4 string `json:"m4"`
	// Mode-4 D1 & D2 response status.
	M4d1d2 string `json:"m4d1d2"`
	// Status of the mode-4 validity bit.
	M4v string `json:"m4v"`
	// Indication of Horizontal Maneuver detection.
	Mah string `json:"mah"`
	// Mode-C altitude (uncorrected pressure altitude), in km.
	Mc float64 `json:"mc"`
	// Indicator that the mode-C response was garbled.
	Mcg bool `json:"mcg"`
	// Status of the mode-C validity bit.
	Mcv string `json:"mcv"`
	// Flag indicating military emergency.
	Milemrgcy bool `json:"milemrgcy"`
	// Flag indicating report separated from different responses at same range. Azimuth
	// may have larger than normal error when present.
	Mrgrpt bool `json:"mrgrpt"`
	// Mode-S Comm B message data.
	Mscommb string `json:"mscommb"`
	// Flag indicating that target was detected using data from an MTI receiver.
	Mti bool `json:"mti"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation.This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Runlength of the primary surveillance radar track, in degrees.
	Psrrl float64 `json:"psrrl"`
	// Type of sensor(s) maintaining track.
	Rad string `json:"rad"`
	// Measured slant range to target from the observing site, in km.
	Rng float64 `json:"rng"`
	// Target range delta between PSR and SSR (reference PSR-SSR), in km.
	Rngdelt float64 `json:"rngdelt"`
	// System Area Code.
	Sac int64 `json:"sac"`
	// Sensor altitude, in kilometers, at time of observation (ts).
	Senalt float64 `json:"senalt"`
	// Sensor WGS84 latitude, in degrees, at time of observation (ts). -90 to 90
	// degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor WGS84 longitude, in degrees, at time of observation (ts). -180 to 180
	// degrees (negative values west of Prime Meridian).
	Senlon float64 `json:"senlon"`
	// System Identification Code.
	Sic int64 `json:"sic"`
	// Flag indicating whether Special Position Indicator (SPI) present in
	// interrogation response.
	Spi bool `json:"spi"`
	// Runlength of the secondary surveillance radar track, in degrees.
	Ssrl float64 `json:"ssrl"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Target confidence.
	Tgtconf string `json:"tgtconf"`
	// Target correlation flag.
	Tgtcorr string `json:"tgtcorr"`
	// Aircraft identification from an aircraft equipped with a Mode S transponder.
	Tgtid string `json:"tgtid"`
	// Data time-in-storage, in seconds. This is the amount of time elapsed between
	// target detection and message transmission.
	Tis float64 `json:"tis"`
	// Track eligibility flag.
	Trkelig string `json:"trkelig"`
	// Value representing a unique reference to a track record within a particular
	// track file. Included when the radar station outputs tracks.
	Trknum int64 `json:"trknum"`
	// Test target indicator.
	Tti string `json:"tti"`
	// Warning/Error Conditions and Target Classification.
	Wectc []string `json:"wectc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Msgfmt                respjson.Field
		Msgts                 respjson.Field
		Msgtyp                respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		Acp                   respjson.Field
		Addr                  respjson.Field
		Af                    respjson.Field
		Aims                  respjson.Field
		Alt3d                 respjson.Field
		Artsqual              respjson.Field
		Az                    respjson.Field
		Azdelt                respjson.Field
		Bcnhits               respjson.Field
		Cartpos               respjson.Field
		Cdm                   respjson.Field
		Code7500              respjson.Field
		Code7600              respjson.Field
		Code7700              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Faa                   respjson.Field
		Grndspd               respjson.Field
		Hdng                  respjson.Field
		IDSensor              respjson.Field
		M1                    respjson.Field
		M1g                   respjson.Field
		M1v                   respjson.Field
		M2                    respjson.Field
		M2g                   respjson.Field
		M2v                   respjson.Field
		M2xv                  respjson.Field
		M3a                   respjson.Field
		M3ag                  respjson.Field
		M3av                  respjson.Field
		M3axv                 respjson.Field
		M4                    respjson.Field
		M4d1d2                respjson.Field
		M4v                   respjson.Field
		Mah                   respjson.Field
		Mc                    respjson.Field
		Mcg                   respjson.Field
		Mcv                   respjson.Field
		Milemrgcy             respjson.Field
		Mrgrpt                respjson.Field
		Mscommb               respjson.Field
		Mti                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		Psrrl                 respjson.Field
		Rad                   respjson.Field
		Rng                   respjson.Field
		Rngdelt               respjson.Field
		Sac                   respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		Sic                   respjson.Field
		Spi                   respjson.Field
		Ssrl                  respjson.Field
		Tags                  respjson.Field
		Tgtconf               respjson.Field
		Tgtcorr               respjson.Field
		Tgtid                 respjson.Field
		Tis                   respjson.Field
		Trkelig               respjson.Field
		Trknum                respjson.Field
		Tti                   respjson.Field
		Wectc                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObservationMonoradarListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationMonoradarListResponse) UnmarshalJSON(data []byte) error {
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
type ObservationMonoradarListResponseDataMode string

const (
	ObservationMonoradarListResponseDataModeReal      ObservationMonoradarListResponseDataMode = "REAL"
	ObservationMonoradarListResponseDataModeTest      ObservationMonoradarListResponseDataMode = "TEST"
	ObservationMonoradarListResponseDataModeSimulated ObservationMonoradarListResponseDataMode = "SIMULATED"
	ObservationMonoradarListResponseDataModeExercise  ObservationMonoradarListResponseDataMode = "EXERCISE"
)

type ObservationMonoradarQueryhelpResponse struct {
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
func (r ObservationMonoradarQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationMonoradarQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A monoradar record contains the raw, and in some cases, processed target reports
// from primary and secondary air surveillance radars. All target positions for
// monoradar reports are recorded as range and azimuth from geographical North
// relative to the detecting radar site. In the case of secondary surveillance
// radars, interrogation response codes are provided as well as quality and
// validation characteristics, when available in the particular record type used to
// generate the record.
type ObservationMonoradarTupleResponse struct {
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
	DataMode ObservationMonoradarTupleResponseDataMode `json:"dataMode,required"`
	// Message format received (i.e. 'ASR9', 'CAT48', 'TPS70', etc..).
	Msgfmt string `json:"msgfmt,required"`
	// Message time, in ISO 8601 UTC format with microsecond precision. This is the
	// time that the data message was released from the site.
	Msgts time.Time `json:"msgts,required" format:"date-time"`
	// Message report type received (i.e. 'SRCH', 'BCN', 'REINF', 'BRTQC', 'PSR',
	// etc..).
	Msgtyp string `json:"msgtyp,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Target detection time, in ISO 8601 UTC format with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Azimuth Change Pulse count at target detection.
	Acp int64 `json:"acp"`
	// Aircraft address (24-bits Mode S address) assigned uniquely to each aircraft.
	Addr string `json:"addr"`
	// Flag indicating military message.
	Af bool `json:"af"`
	// Flag indicating whether AIMS present.
	Aims bool `json:"aims"`
	// Measured height of the target, in km. (for 3D radars).
	Alt3d float64 `json:"alt3d"`
	// ARTS quality.
	Artsqual string `json:"artsqual"`
	// Target azimuth, measured from the observing site, in degrees from true North. If
	// Azimuth Change Pulse (acp) count is provided, az represents the computed angle.
	Az float64 `json:"az"`
	// Target azimuth delta between PSR and SSR (reference PSR-SSR), in degrees.
	Azdelt float64 `json:"azdelt"`
	// Number of beacon hits received on the target.
	Bcnhits int64 `json:"bcnhits"`
	// Array of local 2d-cartesian [x, y] coordinates of target, in km.
	Cartpos []float64 `json:"cartpos"`
	// Climbing/Descending mode indicator.
	Cdm string `json:"cdm"`
	// 7500 squawk present (hijack).
	Code7500 bool `json:"code7500"`
	// 7600 squawk present (loss of comm).
	Code7600 bool `json:"code7600"`
	// 7700 squawk present (general emergency).
	Code7700 bool `json:"code7700"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Flag indicating FAA/Civ message.
	Faa bool `json:"faa"`
	// Target ground speed, in meters/second.
	Grndspd float64 `json:"grndspd"`
	// Target heading, in degrees from true North at the target position.
	Hdng float64 `json:"hdng"`
	// ID of the radar site or station providing the data.
	IDSensor string `json:"idSensor"`
	// Mode-1 interrogation response (mission code).
	M1 string `json:"m1"`
	// Indicator that the mode-1 response was garbled.
	M1g bool `json:"m1g"`
	// Status of the mode-1 validity bit.
	M1v string `json:"m1v"`
	// Mode-2 interrogation response (military identification code).
	M2 string `json:"m2"`
	// Indicator that the mode-2 response was garbled.
	M2g bool `json:"m2g"`
	// Status of the mode-2 validity bit.
	M2v string `json:"m2v"`
	// Status of the mode-2 X-Pulse response validation.
	M2xv string `json:"m2xv"`
	// Mode-3/A interrogation response (aircraft identification).
	M3a string `json:"m3a"`
	// Indicator that the mode-3/A response was garbled.
	M3ag bool `json:"m3ag"`
	// Status of the mode-3/A validity bit.
	M3av string `json:"m3av"`
	// Status of the mode-3 X-Pulse response validation.
	M3axv string `json:"m3axv"`
	// Mode-4 interrogation response (Identification Friend/Foe).
	M4 string `json:"m4"`
	// Mode-4 D1 & D2 response status.
	M4d1d2 string `json:"m4d1d2"`
	// Status of the mode-4 validity bit.
	M4v string `json:"m4v"`
	// Indication of Horizontal Maneuver detection.
	Mah string `json:"mah"`
	// Mode-C altitude (uncorrected pressure altitude), in km.
	Mc float64 `json:"mc"`
	// Indicator that the mode-C response was garbled.
	Mcg bool `json:"mcg"`
	// Status of the mode-C validity bit.
	Mcv string `json:"mcv"`
	// Flag indicating military emergency.
	Milemrgcy bool `json:"milemrgcy"`
	// Flag indicating report separated from different responses at same range. Azimuth
	// may have larger than normal error when present.
	Mrgrpt bool `json:"mrgrpt"`
	// Mode-S Comm B message data.
	Mscommb string `json:"mscommb"`
	// Flag indicating that target was detected using data from an MTI receiver.
	Mti bool `json:"mti"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation.This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// Runlength of the primary surveillance radar track, in degrees.
	Psrrl float64 `json:"psrrl"`
	// Type of sensor(s) maintaining track.
	Rad string `json:"rad"`
	// Measured slant range to target from the observing site, in km.
	Rng float64 `json:"rng"`
	// Target range delta between PSR and SSR (reference PSR-SSR), in km.
	Rngdelt float64 `json:"rngdelt"`
	// System Area Code.
	Sac int64 `json:"sac"`
	// Sensor altitude, in kilometers, at time of observation (ts).
	Senalt float64 `json:"senalt"`
	// Sensor WGS84 latitude, in degrees, at time of observation (ts). -90 to 90
	// degrees (negative values south of equator).
	Senlat float64 `json:"senlat"`
	// Sensor WGS84 longitude, in degrees, at time of observation (ts). -180 to 180
	// degrees (negative values west of Prime Meridian).
	Senlon float64 `json:"senlon"`
	// System Identification Code.
	Sic int64 `json:"sic"`
	// Flag indicating whether Special Position Indicator (SPI) present in
	// interrogation response.
	Spi bool `json:"spi"`
	// Runlength of the secondary surveillance radar track, in degrees.
	Ssrl float64 `json:"ssrl"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Target confidence.
	Tgtconf string `json:"tgtconf"`
	// Target correlation flag.
	Tgtcorr string `json:"tgtcorr"`
	// Aircraft identification from an aircraft equipped with a Mode S transponder.
	Tgtid string `json:"tgtid"`
	// Data time-in-storage, in seconds. This is the amount of time elapsed between
	// target detection and message transmission.
	Tis float64 `json:"tis"`
	// Track eligibility flag.
	Trkelig string `json:"trkelig"`
	// Value representing a unique reference to a track record within a particular
	// track file. Included when the radar station outputs tracks.
	Trknum int64 `json:"trknum"`
	// Test target indicator.
	Tti string `json:"tti"`
	// Warning/Error Conditions and Target Classification.
	Wectc []string `json:"wectc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Msgfmt                respjson.Field
		Msgts                 respjson.Field
		Msgtyp                respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		Acp                   respjson.Field
		Addr                  respjson.Field
		Af                    respjson.Field
		Aims                  respjson.Field
		Alt3d                 respjson.Field
		Artsqual              respjson.Field
		Az                    respjson.Field
		Azdelt                respjson.Field
		Bcnhits               respjson.Field
		Cartpos               respjson.Field
		Cdm                   respjson.Field
		Code7500              respjson.Field
		Code7600              respjson.Field
		Code7700              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Faa                   respjson.Field
		Grndspd               respjson.Field
		Hdng                  respjson.Field
		IDSensor              respjson.Field
		M1                    respjson.Field
		M1g                   respjson.Field
		M1v                   respjson.Field
		M2                    respjson.Field
		M2g                   respjson.Field
		M2v                   respjson.Field
		M2xv                  respjson.Field
		M3a                   respjson.Field
		M3ag                  respjson.Field
		M3av                  respjson.Field
		M3axv                 respjson.Field
		M4                    respjson.Field
		M4d1d2                respjson.Field
		M4v                   respjson.Field
		Mah                   respjson.Field
		Mc                    respjson.Field
		Mcg                   respjson.Field
		Mcv                   respjson.Field
		Milemrgcy             respjson.Field
		Mrgrpt                respjson.Field
		Mscommb               respjson.Field
		Mti                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		Psrrl                 respjson.Field
		Rad                   respjson.Field
		Rng                   respjson.Field
		Rngdelt               respjson.Field
		Sac                   respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		Sic                   respjson.Field
		Spi                   respjson.Field
		Ssrl                  respjson.Field
		Tags                  respjson.Field
		Tgtconf               respjson.Field
		Tgtcorr               respjson.Field
		Tgtid                 respjson.Field
		Tis                   respjson.Field
		Trkelig               respjson.Field
		Trknum                respjson.Field
		Tti                   respjson.Field
		Wectc                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObservationMonoradarTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ObservationMonoradarTupleResponse) UnmarshalJSON(data []byte) error {
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
type ObservationMonoradarTupleResponseDataMode string

const (
	ObservationMonoradarTupleResponseDataModeReal      ObservationMonoradarTupleResponseDataMode = "REAL"
	ObservationMonoradarTupleResponseDataModeTest      ObservationMonoradarTupleResponseDataMode = "TEST"
	ObservationMonoradarTupleResponseDataModeSimulated ObservationMonoradarTupleResponseDataMode = "SIMULATED"
	ObservationMonoradarTupleResponseDataModeExercise  ObservationMonoradarTupleResponseDataMode = "EXERCISE"
)

type ObservationMonoradarListParams struct {
	// Target detection time, in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationMonoradarListParams]'s query parameters as
// `url.Values`.
func (r ObservationMonoradarListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationMonoradarCountParams struct {
	// Target detection time, in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationMonoradarCountParams]'s query parameters as
// `url.Values`.
func (r ObservationMonoradarCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationMonoradarNewBulkParams struct {
	Body []ObservationMonoradarNewBulkParamsBody
	paramObj
}

func (r ObservationMonoradarNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationMonoradarNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// A monoradar record contains the raw, and in some cases, processed target reports
// from primary and secondary air surveillance radars. All target positions for
// monoradar reports are recorded as range and azimuth from geographical North
// relative to the detecting radar site. In the case of secondary surveillance
// radars, interrogation response codes are provided as well as quality and
// validation characteristics, when available in the particular record type used to
// generate the record.
//
// The properties ClassificationMarking, DataMode, Msgfmt, Msgts, Msgtyp, Source,
// Ts are required.
type ObservationMonoradarNewBulkParamsBody struct {
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
	// Message format received (i.e. 'ASR9', 'CAT48', 'TPS70', etc..).
	Msgfmt string `json:"msgfmt,required"`
	// Message time, in ISO 8601 UTC format with microsecond precision. This is the
	// time that the data message was released from the site.
	Msgts time.Time `json:"msgts,required" format:"date-time"`
	// Message report type received (i.e. 'SRCH', 'BCN', 'REINF', 'BRTQC', 'PSR',
	// etc..).
	Msgtyp string `json:"msgtyp,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Target detection time, in ISO 8601 UTC format with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Azimuth Change Pulse count at target detection.
	Acp param.Opt[int64] `json:"acp,omitzero"`
	// Aircraft address (24-bits Mode S address) assigned uniquely to each aircraft.
	Addr param.Opt[string] `json:"addr,omitzero"`
	// Flag indicating military message.
	Af param.Opt[bool] `json:"af,omitzero"`
	// Flag indicating whether AIMS present.
	Aims param.Opt[bool] `json:"aims,omitzero"`
	// Measured height of the target, in km. (for 3D radars).
	Alt3d param.Opt[float64] `json:"alt3d,omitzero"`
	// ARTS quality.
	Artsqual param.Opt[string] `json:"artsqual,omitzero"`
	// Target azimuth, measured from the observing site, in degrees from true North. If
	// Azimuth Change Pulse (acp) count is provided, az represents the computed angle.
	Az param.Opt[float64] `json:"az,omitzero"`
	// Target azimuth delta between PSR and SSR (reference PSR-SSR), in degrees.
	Azdelt param.Opt[float64] `json:"azdelt,omitzero"`
	// Number of beacon hits received on the target.
	Bcnhits param.Opt[int64] `json:"bcnhits,omitzero"`
	// Climbing/Descending mode indicator.
	Cdm param.Opt[string] `json:"cdm,omitzero"`
	// 7500 squawk present (hijack).
	Code7500 param.Opt[bool] `json:"code7500,omitzero"`
	// 7600 squawk present (loss of comm).
	Code7600 param.Opt[bool] `json:"code7600,omitzero"`
	// 7700 squawk present (general emergency).
	Code7700 param.Opt[bool] `json:"code7700,omitzero"`
	// Time the row was created in the database.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Flag indicating FAA/Civ message.
	Faa param.Opt[bool] `json:"faa,omitzero"`
	// Target ground speed, in meters/second.
	Grndspd param.Opt[float64] `json:"grndspd,omitzero"`
	// Target heading, in degrees from true North at the target position.
	Hdng param.Opt[float64] `json:"hdng,omitzero"`
	// ID of the radar site or station providing the data.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Mode-1 interrogation response (mission code).
	M1 param.Opt[string] `json:"m1,omitzero"`
	// Indicator that the mode-1 response was garbled.
	M1g param.Opt[bool] `json:"m1g,omitzero"`
	// Status of the mode-1 validity bit.
	M1v param.Opt[string] `json:"m1v,omitzero"`
	// Mode-2 interrogation response (military identification code).
	M2 param.Opt[string] `json:"m2,omitzero"`
	// Indicator that the mode-2 response was garbled.
	M2g param.Opt[bool] `json:"m2g,omitzero"`
	// Status of the mode-2 validity bit.
	M2v param.Opt[string] `json:"m2v,omitzero"`
	// Status of the mode-2 X-Pulse response validation.
	M2xv param.Opt[string] `json:"m2xv,omitzero"`
	// Mode-3/A interrogation response (aircraft identification).
	M3a param.Opt[string] `json:"m3a,omitzero"`
	// Indicator that the mode-3/A response was garbled.
	M3ag param.Opt[bool] `json:"m3ag,omitzero"`
	// Status of the mode-3/A validity bit.
	M3av param.Opt[string] `json:"m3av,omitzero"`
	// Status of the mode-3 X-Pulse response validation.
	M3axv param.Opt[string] `json:"m3axv,omitzero"`
	// Mode-4 interrogation response (Identification Friend/Foe).
	M4 param.Opt[string] `json:"m4,omitzero"`
	// Mode-4 D1 & D2 response status.
	M4d1d2 param.Opt[string] `json:"m4d1d2,omitzero"`
	// Status of the mode-4 validity bit.
	M4v param.Opt[string] `json:"m4v,omitzero"`
	// Indication of Horizontal Maneuver detection.
	Mah param.Opt[string] `json:"mah,omitzero"`
	// Mode-C altitude (uncorrected pressure altitude), in km.
	Mc param.Opt[float64] `json:"mc,omitzero"`
	// Indicator that the mode-C response was garbled.
	Mcg param.Opt[bool] `json:"mcg,omitzero"`
	// Status of the mode-C validity bit.
	Mcv param.Opt[string] `json:"mcv,omitzero"`
	// Flag indicating military emergency.
	Milemrgcy param.Opt[bool] `json:"milemrgcy,omitzero"`
	// Flag indicating report separated from different responses at same range. Azimuth
	// may have larger than normal error when present.
	Mrgrpt param.Opt[bool] `json:"mrgrpt,omitzero"`
	// Mode-S Comm B message data.
	Mscommb param.Opt[string] `json:"mscommb,omitzero"`
	// Flag indicating that target was detected using data from an MTI receiver.
	Mti param.Opt[bool] `json:"mti,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation.This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Runlength of the primary surveillance radar track, in degrees.
	Psrrl param.Opt[float64] `json:"psrrl,omitzero"`
	// Type of sensor(s) maintaining track.
	Rad param.Opt[string] `json:"rad,omitzero"`
	// Measured slant range to target from the observing site, in km.
	Rng param.Opt[float64] `json:"rng,omitzero"`
	// Target range delta between PSR and SSR (reference PSR-SSR), in km.
	Rngdelt param.Opt[float64] `json:"rngdelt,omitzero"`
	// System Area Code.
	Sac param.Opt[int64] `json:"sac,omitzero"`
	// Sensor altitude, in kilometers, at time of observation (ts).
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude, in degrees, at time of observation (ts). -90 to 90
	// degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude, in degrees, at time of observation (ts). -180 to 180
	// degrees (negative values west of Prime Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// System Identification Code.
	Sic param.Opt[int64] `json:"sic,omitzero"`
	// Flag indicating whether Special Position Indicator (SPI) present in
	// interrogation response.
	Spi param.Opt[bool] `json:"spi,omitzero"`
	// Runlength of the secondary surveillance radar track, in degrees.
	Ssrl param.Opt[float64] `json:"ssrl,omitzero"`
	// Target confidence.
	Tgtconf param.Opt[string] `json:"tgtconf,omitzero"`
	// Target correlation flag.
	Tgtcorr param.Opt[string] `json:"tgtcorr,omitzero"`
	// Aircraft identification from an aircraft equipped with a Mode S transponder.
	Tgtid param.Opt[string] `json:"tgtid,omitzero"`
	// Data time-in-storage, in seconds. This is the amount of time elapsed between
	// target detection and message transmission.
	Tis param.Opt[float64] `json:"tis,omitzero"`
	// Track eligibility flag.
	Trkelig param.Opt[string] `json:"trkelig,omitzero"`
	// Value representing a unique reference to a track record within a particular
	// track file. Included when the radar station outputs tracks.
	Trknum param.Opt[int64] `json:"trknum,omitzero"`
	// Test target indicator.
	Tti param.Opt[string] `json:"tti,omitzero"`
	// Array of local 2d-cartesian [x, y] coordinates of target, in km.
	Cartpos []float64 `json:"cartpos,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// Warning/Error Conditions and Target Classification.
	Wectc []string `json:"wectc,omitzero"`
	paramObj
}

func (r ObservationMonoradarNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationMonoradarNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationMonoradarNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationMonoradarNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type ObservationMonoradarTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Target detection time, in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObservationMonoradarTupleParams]'s query parameters as
// `url.Values`.
func (r ObservationMonoradarTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObservationMonoradarUnvalidatedPublishParams struct {
	Body []ObservationMonoradarUnvalidatedPublishParamsBody
	paramObj
}

func (r ObservationMonoradarUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObservationMonoradarUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// A monoradar record contains the raw, and in some cases, processed target reports
// from primary and secondary air surveillance radars. All target positions for
// monoradar reports are recorded as range and azimuth from geographical North
// relative to the detecting radar site. In the case of secondary surveillance
// radars, interrogation response codes are provided as well as quality and
// validation characteristics, when available in the particular record type used to
// generate the record.
//
// The properties ClassificationMarking, DataMode, Msgfmt, Msgts, Msgtyp, Source,
// Ts are required.
type ObservationMonoradarUnvalidatedPublishParamsBody struct {
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
	// Message format received (i.e. 'ASR9', 'CAT48', 'TPS70', etc..).
	Msgfmt string `json:"msgfmt,required"`
	// Message time, in ISO 8601 UTC format with microsecond precision. This is the
	// time that the data message was released from the site.
	Msgts time.Time `json:"msgts,required" format:"date-time"`
	// Message report type received (i.e. 'SRCH', 'BCN', 'REINF', 'BRTQC', 'PSR',
	// etc..).
	Msgtyp string `json:"msgtyp,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Target detection time, in ISO 8601 UTC format with microsecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Azimuth Change Pulse count at target detection.
	Acp param.Opt[int64] `json:"acp,omitzero"`
	// Aircraft address (24-bits Mode S address) assigned uniquely to each aircraft.
	Addr param.Opt[string] `json:"addr,omitzero"`
	// Flag indicating military message.
	Af param.Opt[bool] `json:"af,omitzero"`
	// Flag indicating whether AIMS present.
	Aims param.Opt[bool] `json:"aims,omitzero"`
	// Measured height of the target, in km. (for 3D radars).
	Alt3d param.Opt[float64] `json:"alt3d,omitzero"`
	// ARTS quality.
	Artsqual param.Opt[string] `json:"artsqual,omitzero"`
	// Target azimuth, measured from the observing site, in degrees from true North. If
	// Azimuth Change Pulse (acp) count is provided, az represents the computed angle.
	Az param.Opt[float64] `json:"az,omitzero"`
	// Target azimuth delta between PSR and SSR (reference PSR-SSR), in degrees.
	Azdelt param.Opt[float64] `json:"azdelt,omitzero"`
	// Number of beacon hits received on the target.
	Bcnhits param.Opt[int64] `json:"bcnhits,omitzero"`
	// Climbing/Descending mode indicator.
	Cdm param.Opt[string] `json:"cdm,omitzero"`
	// 7500 squawk present (hijack).
	Code7500 param.Opt[bool] `json:"code7500,omitzero"`
	// 7600 squawk present (loss of comm).
	Code7600 param.Opt[bool] `json:"code7600,omitzero"`
	// 7700 squawk present (general emergency).
	Code7700 param.Opt[bool] `json:"code7700,omitzero"`
	// Time the row was created in the database.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Flag indicating FAA/Civ message.
	Faa param.Opt[bool] `json:"faa,omitzero"`
	// Target ground speed, in meters/second.
	Grndspd param.Opt[float64] `json:"grndspd,omitzero"`
	// Target heading, in degrees from true North at the target position.
	Hdng param.Opt[float64] `json:"hdng,omitzero"`
	// ID of the radar site or station providing the data.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Mode-1 interrogation response (mission code).
	M1 param.Opt[string] `json:"m1,omitzero"`
	// Indicator that the mode-1 response was garbled.
	M1g param.Opt[bool] `json:"m1g,omitzero"`
	// Status of the mode-1 validity bit.
	M1v param.Opt[string] `json:"m1v,omitzero"`
	// Mode-2 interrogation response (military identification code).
	M2 param.Opt[string] `json:"m2,omitzero"`
	// Indicator that the mode-2 response was garbled.
	M2g param.Opt[bool] `json:"m2g,omitzero"`
	// Status of the mode-2 validity bit.
	M2v param.Opt[string] `json:"m2v,omitzero"`
	// Status of the mode-2 X-Pulse response validation.
	M2xv param.Opt[string] `json:"m2xv,omitzero"`
	// Mode-3/A interrogation response (aircraft identification).
	M3a param.Opt[string] `json:"m3a,omitzero"`
	// Indicator that the mode-3/A response was garbled.
	M3ag param.Opt[bool] `json:"m3ag,omitzero"`
	// Status of the mode-3/A validity bit.
	M3av param.Opt[string] `json:"m3av,omitzero"`
	// Status of the mode-3 X-Pulse response validation.
	M3axv param.Opt[string] `json:"m3axv,omitzero"`
	// Mode-4 interrogation response (Identification Friend/Foe).
	M4 param.Opt[string] `json:"m4,omitzero"`
	// Mode-4 D1 & D2 response status.
	M4d1d2 param.Opt[string] `json:"m4d1d2,omitzero"`
	// Status of the mode-4 validity bit.
	M4v param.Opt[string] `json:"m4v,omitzero"`
	// Indication of Horizontal Maneuver detection.
	Mah param.Opt[string] `json:"mah,omitzero"`
	// Mode-C altitude (uncorrected pressure altitude), in km.
	Mc param.Opt[float64] `json:"mc,omitzero"`
	// Indicator that the mode-C response was garbled.
	Mcg param.Opt[bool] `json:"mcg,omitzero"`
	// Status of the mode-C validity bit.
	Mcv param.Opt[string] `json:"mcv,omitzero"`
	// Flag indicating military emergency.
	Milemrgcy param.Opt[bool] `json:"milemrgcy,omitzero"`
	// Flag indicating report separated from different responses at same range. Azimuth
	// may have larger than normal error when present.
	Mrgrpt param.Opt[bool] `json:"mrgrpt,omitzero"`
	// Mode-S Comm B message data.
	Mscommb param.Opt[string] `json:"mscommb,omitzero"`
	// Flag indicating that target was detected using data from an MTI receiver.
	Mti param.Opt[bool] `json:"mti,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by observation source to indicate the sensor
	// identifier which produced this observation.This may be an internal identifier
	// and not necessarily a valid sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// Runlength of the primary surveillance radar track, in degrees.
	Psrrl param.Opt[float64] `json:"psrrl,omitzero"`
	// Type of sensor(s) maintaining track.
	Rad param.Opt[string] `json:"rad,omitzero"`
	// Measured slant range to target from the observing site, in km.
	Rng param.Opt[float64] `json:"rng,omitzero"`
	// Target range delta between PSR and SSR (reference PSR-SSR), in km.
	Rngdelt param.Opt[float64] `json:"rngdelt,omitzero"`
	// System Area Code.
	Sac param.Opt[int64] `json:"sac,omitzero"`
	// Sensor altitude, in kilometers, at time of observation (ts).
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// Sensor WGS84 latitude, in degrees, at time of observation (ts). -90 to 90
	// degrees (negative values south of equator).
	Senlat param.Opt[float64] `json:"senlat,omitzero"`
	// Sensor WGS84 longitude, in degrees, at time of observation (ts). -180 to 180
	// degrees (negative values west of Prime Meridian).
	Senlon param.Opt[float64] `json:"senlon,omitzero"`
	// System Identification Code.
	Sic param.Opt[int64] `json:"sic,omitzero"`
	// Flag indicating whether Special Position Indicator (SPI) present in
	// interrogation response.
	Spi param.Opt[bool] `json:"spi,omitzero"`
	// Runlength of the secondary surveillance radar track, in degrees.
	Ssrl param.Opt[float64] `json:"ssrl,omitzero"`
	// Target confidence.
	Tgtconf param.Opt[string] `json:"tgtconf,omitzero"`
	// Target correlation flag.
	Tgtcorr param.Opt[string] `json:"tgtcorr,omitzero"`
	// Aircraft identification from an aircraft equipped with a Mode S transponder.
	Tgtid param.Opt[string] `json:"tgtid,omitzero"`
	// Data time-in-storage, in seconds. This is the amount of time elapsed between
	// target detection and message transmission.
	Tis param.Opt[float64] `json:"tis,omitzero"`
	// Track eligibility flag.
	Trkelig param.Opt[string] `json:"trkelig,omitzero"`
	// Value representing a unique reference to a track record within a particular
	// track file. Included when the radar station outputs tracks.
	Trknum param.Opt[int64] `json:"trknum,omitzero"`
	// Test target indicator.
	Tti param.Opt[string] `json:"tti,omitzero"`
	// Array of local 2d-cartesian [x, y] coordinates of target, in km.
	Cartpos []float64 `json:"cartpos,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	// Warning/Error Conditions and Target Classification.
	Wectc []string `json:"wectc,omitzero"`
	paramObj
}

func (r ObservationMonoradarUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObservationMonoradarUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservationMonoradarUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObservationMonoradarUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
