// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// MtiService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMtiService] method instead.
type MtiService struct {
	Options []option.RequestOption
	History MtiHistoryService
}

// NewMtiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMtiService(opts ...option.RequestOption) (r MtiService) {
	r = MtiService{}
	r.Options = opts
	r.History = NewMtiHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *MtiService) List(ctx context.Context, query MtiListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[MtiListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/mti"
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
func (r *MtiService) ListAutoPaging(ctx context.Context, query MtiListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[MtiListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *MtiService) Count(ctx context.Context, query MtiCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/mti/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// Moving Target Indicator (MTI) records as a POST body and ingest into the
// database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *MtiService) NewBulk(ctx context.Context, body MtiNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/mti/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *MtiService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *MtiQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/mti/queryhelp"
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
func (r *MtiService) Tuple(ctx context.Context, query MtiTupleParams, opts ...option.RequestOption) (res *[]MtiFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/mti/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take a list of Moving Target Indicator (MTI) formatted data
// as a POST body and ingest into the database. This operation is intended to be
// used for automated feeds into UDL. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *MtiService) UnvalidatedPublish(ctx context.Context, body MtiUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-mti"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Information on the mission and flight plans, the type and configuration of the
// platform, and the reference time.
type MtiListResponse struct {
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
	DataMode MtiListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	Dwells []MtiListResponseDwell `json:"dwells"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	FreeTexts []MtiListResponseFreeText `json:"freeTexts"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	Hrrs []MtiListResponseHrr `json:"hrrs"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	JobDefs []MtiListResponseJobDef `json:"jobDefs"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	JobRequests []MtiListResponseJobRequest `json:"jobRequests"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	Missions []MtiListResponseMission `json:"missions"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	P10 int64 `json:"p10"`
	// Nationality of the platform providing the data.
	P3 string `json:"p3"`
	// Control / handling marking.
	P6 string `json:"p6"`
	// Data record exercise indicator.
	P7 string `json:"p7"`
	// ID of the platform providing the data (tail number for air platform, name and
	// numerical designator for space-based platforms).
	P8 string `json:"p8"`
	// Integer field, assigned by the platform, that uniquely identifies the mission
	// for the platform.
	P9 int64 `json:"p9"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	PlatformLocs []MtiListResponsePlatformLoc `json:"platformLocs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Dwells                respjson.Field
		FreeTexts             respjson.Field
		Hrrs                  respjson.Field
		JobDefs               respjson.Field
		JobRequests           respjson.Field
		Missions              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		P10                   respjson.Field
		P3                    respjson.Field
		P6                    respjson.Field
		P7                    respjson.Field
		P8                    respjson.Field
		P9                    respjson.Field
		PlatformLocs          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponse) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponse) UnmarshalJSON(data []byte) error {
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
type MtiListResponseDataMode string

const (
	MtiListResponseDataModeReal      MtiListResponseDataMode = "REAL"
	MtiListResponseDataModeTest      MtiListResponseDataMode = "TEST"
	MtiListResponseDataModeSimulated MtiListResponseDataMode = "SIMULATED"
	MtiListResponseDataModeExercise  MtiListResponseDataMode = "EXERCISE"
)

type MtiListResponseDwell struct {
	// Factor which modifies the value of the reported target latitude (Delta Latitude,
	// field D32.4).
	D10 float64 `json:"d10"`
	// Factor which modifies the value of the reported target longitude (Delta
	// Longitude, field D32.5).
	D11 float64 `json:"d11"`
	// Standard deviation in the estimated horizontal sensor location at the time of
	// the dwell, measured along the sensor track direction (field D15), in
	// centimeters.
	D12 int64 `json:"d12"`
	// Standard deviation in the estimated horizontal sensor location at the time of
	// the dwell, measured orthogonal to the sensor track direction (field D15), in
	// centimeters.
	D13 int64 `json:"d13"`
	// Standard deviation of the sensor altitude estimate (field D9), in centimeters.
	D14 int64 `json:"d14"`
	// Ground track of the sensor at the time of the dwell, as the angle in degrees
	// (clockwise) from True North.
	D15 float64 `json:"d15"`
	// Ground speed of the sensor at the time of the dwell, in millimeters per second.
	D16 int64 `json:"d16"`
	// Velocity of the sensor in the vertical direction, in decimeters per second.
	D17 int64 `json:"d17"`
	// Standard deviation of the estimate of the sensor track, in degrees.
	D18 int64 `json:"d18"`
	// Standard deviation of estimate of the sensor speed, in millimeters per second.
	D19 int64 `json:"d19"`
	// Sequential count of a revisit of the bounding area in the last sent Job
	// Definition Segment, where a Revisit Index of '0' indicates the first revisit.
	D2 int64 `json:"d2"`
	// Standard deviation of estimate of the sensor vertical velocity, expressed in
	// centimeters per second.
	D20 int64 `json:"d20"`
	// Heading of the platform at the time of the dwell, as the angle in degrees
	// (clockwise) from True North to the roll axis of the platform.
	D21 float64 `json:"d21"`
	// Pitch angle of the platform at the time of the dwell, in degrees.
	D22 float64 `json:"d22"`
	// Roll angle of the platform at the time of the dwell, in degrees.
	D23 float64 `json:"d23"`
	// The North-South position of the center of the dwell area, expressed as degrees
	// North (positive) or South (negative) of the Equator.
	D24 float64 `json:"d24"`
	// The East-West position of the center of the dwell area, expressed as degrees
	// East (positive, 0 to 180) or West (negative, 0 to -180) of the Prime Meridian.
	D25 float64 `json:"d25"`
	// Distance on the earth surface, expressed in kilometers, from the near edge to
	// the center of the dwell area.
	D26 float64 `json:"d26"`
	// For dwell based radars, one-half of the 3-dB beamwidth. For non-dwell based
	// radars, the angle between the beginning of the dwell to the center of the dwell.
	// Measured in degrees.
	D27 float64 `json:"d27"`
	// Rotation of the sensor broadside face about the local vertical axis of the
	// platform, in degrees.
	D28 float64 `json:"d28"`
	// Rotation angle of the sensor about the transverse axis of the sensor broadside,
	// in degrees.
	D29 float64 `json:"d29"`
	// Temporally sequential count of a dwell within the revisit of a particular
	// bounding area for a given job ID.
	D3 int64 `json:"d3"`
	// Rotation angle of the sensor about the transverse axis of the sensor broadside,
	// in degrees.
	D30 float64 `json:"d30"`
	// Minimum velocity component, along the line of sight, which can be detected by
	// the sensor, in decimeters per second.
	D31 int64 `json:"d31"`
	// Minimum velocity component, along the line of sight, which can be detected by
	// the sensor, in decimeters per second.
	D32 []MtiListResponseDwellD32 `json:"d32"`
	// Flag indicating the last dwell of the revisit.
	D4 bool `json:"d4"`
	// Count of the total number of targets reported during this dwell and sent in this
	// Dwell Segment.
	D5 int64 `json:"d5"`
	// Elapsed time, expressed in milliseconds, from midnight at the beginning of the
	// day specified in the Reference Time fields (missionRefTime) of the Mission
	// Segment.
	D6 int64 `json:"d6"`
	// North-South position of the sensor at the temporal center of the dwell, in
	// degrees.
	D7 float64 `json:"d7"`
	// The East-West position of the sensor at the temporal center of the dwell, in
	// degrees East (positive, 0 to 180) or West (negative, 0 to -180) of the Prime
	// Meridian.
	D8 float64 `json:"d8"`
	// The altitude of the sensor at temporal center of the dwell, above the WGS 84
	// ellipsoid, expressed in centimeters.
	D9 int64 `json:"d9"`
	// Dwell timestamp in ISO8601 UTC format.
	Dwellts time.Time `json:"dwellts" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		D10         respjson.Field
		D11         respjson.Field
		D12         respjson.Field
		D13         respjson.Field
		D14         respjson.Field
		D15         respjson.Field
		D16         respjson.Field
		D17         respjson.Field
		D18         respjson.Field
		D19         respjson.Field
		D2          respjson.Field
		D20         respjson.Field
		D21         respjson.Field
		D22         respjson.Field
		D23         respjson.Field
		D24         respjson.Field
		D25         respjson.Field
		D26         respjson.Field
		D27         respjson.Field
		D28         respjson.Field
		D29         respjson.Field
		D3          respjson.Field
		D30         respjson.Field
		D31         respjson.Field
		D32         respjson.Field
		D4          respjson.Field
		D5          respjson.Field
		D6          respjson.Field
		D7          respjson.Field
		D8          respjson.Field
		D9          respjson.Field
		Dwellts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponseDwell) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponseDwell) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A grouping of zero or more target reports for which the sensor provides a single
// time, sensor position, reference position on the ground with simple estimates
// for the observed area at the reported time, and other pertinent data.
type MtiListResponseDwellD32 struct {
	// Sequential count of this MTI report within the dwell.
	D32_1 int64 `json:"d32_1"`
	// The classification of the target (i.e. vehicle, aircraft, …).
	D32_10 string `json:"d32_10"`
	// Estimated probability that the target classification field is correctly
	// classified.
	D32_11 int64 `json:"d32_11"`
	// Standard deviation of the estimated slant range of the reported detection, in
	// centimeters.
	D32_12 int64 `json:"d32_12"`
	// Standard deviation of the position estimate, in the cross-range direction, of
	// the reported detection, in decimeters.
	D32_13 int64 `json:"d32_13"`
	// Standard deviation of the estimated geodetic height, in meters.
	D32_14 int64 `json:"d32_14"`
	// Standard deviation of the measured line-of-sight velocity component, in
	// centimeters per second.
	D32_15 int64 `json:"d32_15"`
	// The Truth Tag- Application is the Application Field truncated to 8 bits, from
	// the Entity State Protocol Data Unit (PDU) used to generate the MTI Target.
	D32_16 int64 `json:"d32_16"`
	// The Truth Tag - Entity is the Entity Field from the Entity State PDU used to
	// generate the MTI Target.
	D32_17 int64 `json:"d32_17"`
	// Estimated radar cross section of the target return, in half-decibels.
	D32_18 int64 `json:"d32_18"`
	// The North-South position of the reported detection, expressed as degrees North
	// (positive) or South (negative) of the Equator.
	D32_2 float64 `json:"d32_2"`
	// The East-West position of the reported detection, expressed as degrees East
	// (positive) from the Prime Meridian.
	D32_3 float64 `json:"d32_3"`
	// The North-South position of the reported detection, expressed as degrees North
	// (positive) or South (negative) from the Dwell Area Center Latitude.
	D32_4 int64 `json:"d32_4"`
	// The East-West position of the reported detection, expressed as degrees East
	// (positive, 0 to 180) or West (negative, 0 to -180) of the Prime Meridian from
	// the Dwell Area Center Longitude.
	D32_5 int64 `json:"d32_5"`
	// Height of the reported detection, referenced to its position above the WGS 84
	// ellipsoid, in meters.
	D32_6 int64 `json:"d32_6"`
	// The component of velocity for the reported detection, expressed in centimeters
	// per second, corrected for platform motion, along the line of sight between the
	// sensor and the reported detection, where the positive direction is away from the
	// sensor.
	D32_7 int64 `json:"d32_7"`
	// The target wrap velocity permits trackers to un-wrap velocities for targets with
	// line-of-sight components large enough to exceed the first velocity period.
	// Expressed in centimeters/sec.
	D32_8 int64 `json:"d32_8"`
	// Estimated signal-to-noise ratio (SNR) of the target return, in decibels.
	D32_9 int64 `json:"d32_9"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		D32_1       respjson.Field
		D32_10      respjson.Field
		D32_11      respjson.Field
		D32_12      respjson.Field
		D32_13      respjson.Field
		D32_14      respjson.Field
		D32_15      respjson.Field
		D32_16      respjson.Field
		D32_17      respjson.Field
		D32_18      respjson.Field
		D32_2       respjson.Field
		D32_3       respjson.Field
		D32_4       respjson.Field
		D32_5       respjson.Field
		D32_6       respjson.Field
		D32_7       respjson.Field
		D32_8       respjson.Field
		D32_9       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponseDwellD32) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponseDwellD32) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiListResponseFreeText struct {
	// The originator of the Free Text message.
	F1 string `json:"f1"`
	// The recipient for which the Free Text message is intended.
	F2 string `json:"f2"`
	// Free text data message.
	F3 string `json:"f3"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		F1          respjson.Field
		F2          respjson.Field
		F3          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponseFreeText) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponseFreeText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiListResponseHrr struct {
	// Detection threshold used to isolate significant target scatterer pixels,
	// expressed as power relative to clutter mean in negative quarter-decibels.
	H10 int64 `json:"h10"`
	// 3dB range impulse response of the radar, expressed in centimeters.
	H11 float64 `json:"h11"`
	// Slant Range pixel spacing after over sampling, expressed in centimeters.
	H12 float64 `json:"h12"`
	// 3dB Doppler resolution of the radar, expressed in Hertz.
	H13 float64 `json:"h13"`
	// Doppler pixel spacing after over sampling, expressed in Hertz.
	H14 float64 `json:"h14"`
	// Center Frequency of the radar in GHz.
	H15 float64 `json:"h15"`
	// Enumeration table denoting the compression technique used.
	H16 string `json:"h16"`
	// Enumeration table indicating the spectral weighting used in the range
	// compression process.
	H17 string `json:"h17"`
	// Enumeration table indicating the spectral weighting used in the cross-range or
	// Doppler compression process.
	H18 string `json:"h18"`
	// Initial power of the peak scatterer, expressed in dB.
	H19 float64 `json:"h19"`
	// Sequential count of a revisit of the bounding area for a given job ID.
	H2 int64 `json:"h2"`
	// RCS of the peak scatterer, expressed in half-decibels (dB/2).
	H20 int64 `json:"h20"`
	// When the RDM does not correlate to a single MTI report index or when the center
	// range bin does not correlate to the center of the dwell; provide the range
	// sample offset in meters from Dwell Center (positive is away from the sensor) of
	// the first scatterer record.
	H21 int64 `json:"h21"`
	// When the RDM does not correlate to a single MTI report index or the center
	// doppler bin does not correlate to the doppler centroid of the dwell; Doppler
	// sample value in Hz of the first scatterer record.
	H22 int64 `json:"h22"`
	// Enumeration field which designates the type of data being delivered.
	H23 string `json:"h23"`
	// Flag field to indicate the additional signal processing techniques applied to
	// the data.
	H24 string `json:"h24"`
	// Number of pixels in the range dimension of the chip.
	H27 int64 `json:"h27"`
	// Distance from Range Bin to closest edge in the entire chip, expressed in
	// centimeters.
	H28 int64 `json:"h28"`
	// Relative velocity to skin line.
	H29 int64 `json:"h29"`
	// Sequential count of a dwell within the revisit of a particular bounding area for
	// a given job ID.
	H3 int64 `json:"h3"`
	// Computed object length based upon HRR profile, in meters.
	H30 int64 `json:"h30"`
	// Standard deviation of estimate of the object length, expressed in meters.
	H31 int64 `json:"h31"`
	// Standard deviation of estimate of the object length, expressed in meters.
	H32 []MtiListResponseHrrH32 `json:"h32"`
	// Flag to indicate the last dwell of the revisit.
	H4 bool `json:"h4"`
	// Sequential index of the associated MTI Report.
	H5 int64 `json:"h5"`
	// Number of Range Doppler pixels that exceed target scatterer threshold and are
	// reported in this segment.
	H6 int64 `json:"h6"`
	// Number of Range Bins/Samples in a Range Doppler Chip.
	H7 int64 `json:"h7"`
	// Number of Doppler bins in a Range-Doppler chip.
	H8 int64 `json:"h8"`
	// The Peak Scatter returns the maximum power level (e.g. in milliwatts, or dBm)
	// registered by the sensor.
	H9 int64 `json:"h9"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H10         respjson.Field
		H11         respjson.Field
		H12         respjson.Field
		H13         respjson.Field
		H14         respjson.Field
		H15         respjson.Field
		H16         respjson.Field
		H17         respjson.Field
		H18         respjson.Field
		H19         respjson.Field
		H2          respjson.Field
		H20         respjson.Field
		H21         respjson.Field
		H22         respjson.Field
		H23         respjson.Field
		H24         respjson.Field
		H27         respjson.Field
		H28         respjson.Field
		H29         respjson.Field
		H3          respjson.Field
		H30         respjson.Field
		H31         respjson.Field
		H32         respjson.Field
		H4          respjson.Field
		H5          respjson.Field
		H6          respjson.Field
		H7          respjson.Field
		H8          respjson.Field
		H9          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponseHrr) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponseHrr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HRR Scatterer record for a target pixel that exceeds the target detection
// threshold.
type MtiListResponseHrrH32 struct {
	// Scatterer’s power magnitude.
	H32_1 int64 `json:"h32_1"`
	// Scatterer’s complex phase, in degrees.
	H32_2 int64 `json:"h32_2"`
	// Scatterer’s Range index relative to Range-Doppler chip, where increasing index
	// equates to increasing range.
	H32_3 int64 `json:"h32_3"`
	// Scatterer’s Doppler index relative to Range-Doppler chip, where increasing index
	// equates to increasing Doppler.
	H32_4 int64 `json:"h32_4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H32_1       respjson.Field
		H32_2       respjson.Field
		H32_3       respjson.Field
		H32_4       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponseHrrH32) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponseHrrH32) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The means for the platform to pass information pertaining to the sensor job that
// will be performed and details of the location parameters (terrain elevation
// model and geoid model) used in the measurement.
type MtiListResponseJobDef struct {
	// North-South position of the third corner (Point C) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J10 float64 `json:"j10"`
	// East-West position of the third corner (Point C) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J11 float64 `json:"j11"`
	// North-South position of the fourth corner (Point D) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J12 float64 `json:"j12"`
	// East-West position of the fourth corner (Point D) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J13 float64 `json:"j13"`
	// Mode in which the radar will operate for the given job ID.
	J14 string `json:"j14"`
	// The nominal revisit interval for the job ID, expressed in deciseconds. Value of
	// zero, indicates that the sensor is not revisiting the previous area.
	J15 int64 `json:"j15"`
	// Nominal estimate of the standard deviation in the estimated horizontal (along
	// track) sensor location, expressed in decimeters. measured along the sensor track
	// direction defined in the Dwell segment.
	J16 int64 `json:"j16"`
	// Nominal estimate of the standard deviation in the estimated horizontal sensor
	// location, measured orthogonal to the track direction, expressed in decimeters.
	J17 int64 `json:"j17"`
	// Nominal estimate of the standard deviation of the measured sensor altitude,
	// expressed in decimeters.
	J18 int64 `json:"j18"`
	// Standard deviation of the estimate of sensor track heading, expressed in
	// degrees.
	J19 int64 `json:"j19"`
	// The type of sensor or the platform.
	J2 string `json:"j2"`
	// Nominal standard deviation of the estimate of sensor speed, expressed in
	// millimeters per second.
	J20 int64 `json:"j20"`
	// Nominal standard deviation of the slant range of the reported detection,
	// expressed in centimeters.
	J21 int64 `json:"j21"`
	// Nominal standard deviation of the measured cross angle to the reported
	// detection, expressed in degrees.
	J22 float64 `json:"j22"`
	// Nominal standard deviation of the velocity line-of-sight component, expressed in
	// centimeters per second.
	J23 int64 `json:"j23"`
	// Nominal minimum velocity component along the line of sight, which can be
	// detected by the sensor, expressed in decimeters per second.
	J24 int64 `json:"j24"`
	// Nominal probability that an unobscured ten square-meter target will be detected
	// within the given area of surveillance.
	J25 int64 `json:"j25"`
	// The expected density of False Alarms (FA), expressed as the negative of the
	// decibel value.
	J26 int64 `json:"j26"`
	// The terrain elevation model used for developing the target reports.
	J27 string `json:"j27"`
	// The geoid model used for developing the target reports.
	J28 string `json:"j28"`
	// Identifier of the particular variant of the sensor type.
	J3 string `json:"j3"`
	// Flag field indicating whether filtering has been applied to the targets detected
	// within the dwell area.
	J4 int64 `json:"j4"`
	// Priority of this tasking request relative to all other active tasking requests
	// scheduled for execution on the specified platform.
	J5 int64 `json:"j5"`
	// North-South position of the first corner (Point A) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J6 float64 `json:"j6"`
	// East-West position of the first corner (Point A) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J7 float64 `json:"j7"`
	// North-South position of the second corner (Point B) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J8 float64 `json:"j8"`
	// East-West position of the second corner (Point B) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J9 float64 `json:"j9"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		J10         respjson.Field
		J11         respjson.Field
		J12         respjson.Field
		J13         respjson.Field
		J14         respjson.Field
		J15         respjson.Field
		J16         respjson.Field
		J17         respjson.Field
		J18         respjson.Field
		J19         respjson.Field
		J2          respjson.Field
		J20         respjson.Field
		J21         respjson.Field
		J22         respjson.Field
		J23         respjson.Field
		J24         respjson.Field
		J25         respjson.Field
		J26         respjson.Field
		J27         respjson.Field
		J28         respjson.Field
		J3          respjson.Field
		J4          respjson.Field
		J5          respjson.Field
		J6          respjson.Field
		J7          respjson.Field
		J8          respjson.Field
		J9          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponseJobDef) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponseJobDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiListResponseJobRequest struct {
	// Specifies the Earliest Start Time for which the service is requested. Composite
	// of fields R15-R20.
	JobReqEst time.Time `json:"jobReqEst" format:"date-time"`
	// The requestor of the sensor service.
	R1 string `json:"r1"`
	// North-South position of the fourth corner (Point D) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R10 float64 `json:"r10"`
	// East-West position of the fourth corner (Point D) defining the requested area
	// for service, expressed as degrees East (positive, 0 to 180) or West (negative, 0
	// to -180) of the Prime Meridian.
	R11 float64 `json:"r11"`
	// Identifies the radar mode requested by the requestor.
	R12 string `json:"r12"`
	// Specifies the radar range resolution requested by the requestor, expressed in
	// centimeters.
	R13 int64 `json:"r13"`
	// Specifies the radar cross-range resolution requested by the requestor, expressed
	// in decimeters.
	R14 int64 `json:"r14"`
	// Identifier for the tasking message sent by the requesting station.
	R2 string `json:"r2"`
	// Specifies the maximum time from the requested start time after which the request
	// is to be abandoned, expressed in seconds.
	R21 int64 `json:"r21"`
	// Specifies the time duration for the radar job, measured from the actual start of
	// the job, expressed in seconds.
	R22 int64 `json:"r22"`
	// Specifies the revisit interval for the radar job, expressed in deciseconds.
	R23 int64 `json:"r23"`
	// the type of sensor or the platform.
	R24 string `json:"r24"`
	// The particular variant of the sensor type.
	R25 string `json:"r25"`
	// Flag field indicating that it is an initial request (flag = 0) or the desire of
	// the requestor to cancel (flag = 1) the requested job.
	R26 bool `json:"r26"`
	// The priority of the request relative to other requests originated by the
	// requesting station.
	R3 int64 `json:"r3"`
	// North-South position of the first corner (Point A) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R4 float64 `json:"r4"`
	// East-West position of the first corner (Point A) defining the requested area for
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	R5 float64 `json:"r5"`
	// North-South position of the second corner (Point B) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R6 float64 `json:"r6"`
	// East-West position of the second corner (Point B) defining the requested area
	// for service, expressed as degrees East (positive, 0 to 180) or West (negative, 0
	// to -180) of the Prime Meridian.
	R7 float64 `json:"r7"`
	// North-South position of the third corner (Point C) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R8 float64 `json:"r8"`
	// East-West position of the third corner (Point C) defining the requested area for
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	R9 float64 `json:"r9"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobReqEst   respjson.Field
		R1          respjson.Field
		R10         respjson.Field
		R11         respjson.Field
		R12         respjson.Field
		R13         respjson.Field
		R14         respjson.Field
		R2          respjson.Field
		R21         respjson.Field
		R22         respjson.Field
		R23         respjson.Field
		R24         respjson.Field
		R25         respjson.Field
		R26         respjson.Field
		R3          respjson.Field
		R4          respjson.Field
		R5          respjson.Field
		R6          respjson.Field
		R7          respjson.Field
		R8          respjson.Field
		R9          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponseJobRequest) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponseJobRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiListResponseMission struct {
	// The mission plan id.
	M1 string `json:"m1"`
	// Unique identification of the flight plan.
	M2 string `json:"m2"`
	// Platform type that originated the data.
	M3 string `json:"m3"`
	// Identification of the platform variant, modifications, etc.
	M4 string `json:"m4"`
	// Mission origination date.
	MsnRefTs time.Time `json:"msnRefTs" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		M1          respjson.Field
		M2          respjson.Field
		M3          respjson.Field
		M4          respjson.Field
		MsnRefTs    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponseMission) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponseMission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiListResponsePlatformLoc struct {
	// Elapsed time, expressed in milliseconds, from midnight at the beginning of the
	// day specified in the Reference Time fields of the Mission Segment to the time
	// the report is prepared.
	L1 int64 `json:"l1"`
	// North-South position of the platform at the time the report is prepared,
	// expressed as degrees North (positive) or South (negative) of the Equator.
	L2 float64 `json:"l2"`
	// East-West position of the platform at the time the report is prepared, expressed
	// as degrees East (positive) from the Prime Meridian.
	L3 float64 `json:"l3"`
	// Altitude of the platform at the time the report is prepared, referenced to its
	// position above the WGS-84 ellipsoid, in centimeters.
	L4 int64 `json:"l4"`
	// Ground track of the platform at the time the report is prepared, expressed as
	// the angle in degrees (clockwise) from True North.
	L5 float64 `json:"l5"`
	// Ground speed of the platform at the time the report is prepared, expressed as
	// millimeters per second.
	L6 int64 `json:"l6"`
	// Velocity of the platform in the vertical direction, expressed as decimeters per
	// second.
	L7 int64 `json:"l7"`
	// Platform location timestamp in ISO8601 UTC format.
	Platlocts time.Time `json:"platlocts" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		L1          respjson.Field
		L2          respjson.Field
		L3          respjson.Field
		L4          respjson.Field
		L5          respjson.Field
		L6          respjson.Field
		L7          respjson.Field
		Platlocts   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiListResponsePlatformLoc) RawJSON() string { return r.JSON.raw }
func (r *MtiListResponsePlatformLoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiQueryhelpResponse struct {
	AodrSupported         bool                            `json:"aodrSupported"`
	ClassificationMarking string                          `json:"classificationMarking"`
	Description           string                          `json:"description"`
	HistorySupported      bool                            `json:"historySupported"`
	Name                  string                          `json:"name"`
	Parameters            []MtiQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                        `json:"requiredRoles"`
	RestSupported         bool                            `json:"restSupported"`
	SortSupported         bool                            `json:"sortSupported"`
	TypeName              string                          `json:"typeName"`
	Uri                   string                          `json:"uri"`
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
func (r MtiQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *MtiQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiQueryhelpResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MtiQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *MtiQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiListParams struct {
	// Time the row was created in the database. (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MtiListParams]'s query parameters as `url.Values`.
func (r MtiListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MtiCountParams struct {
	// Time the row was created in the database. (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MtiCountParams]'s query parameters as `url.Values`.
func (r MtiCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MtiNewBulkParams struct {
	Body []MtiNewBulkParamsBody
	paramObj
}

func (r MtiNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *MtiNewBulkParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// Information on the mission and flight plans, the type and configuration of the
// platform, and the reference time.
//
// The properties ClassificationMarking, DataMode, Source are required.
type MtiNewBulkParamsBody struct {
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
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	P10 param.Opt[int64] `json:"p10,omitzero"`
	// Nationality of the platform providing the data.
	P3 param.Opt[string] `json:"p3,omitzero"`
	// Control / handling marking.
	P6 param.Opt[string] `json:"p6,omitzero"`
	// Data record exercise indicator.
	P7 param.Opt[string] `json:"p7,omitzero"`
	// ID of the platform providing the data (tail number for air platform, name and
	// numerical designator for space-based platforms).
	P8 param.Opt[string] `json:"p8,omitzero"`
	// Integer field, assigned by the platform, that uniquely identifies the mission
	// for the platform.
	P9 param.Opt[int64] `json:"p9,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	Dwells []MtiNewBulkParamsBodyDwell `json:"dwells,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	FreeTexts []MtiNewBulkParamsBodyFreeText `json:"freeTexts,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	Hrrs []MtiNewBulkParamsBodyHrr `json:"hrrs,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	JobDefs []MtiNewBulkParamsBodyJobDef `json:"jobDefs,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	JobRequests []MtiNewBulkParamsBodyJobRequest `json:"jobRequests,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	Missions []MtiNewBulkParamsBodyMission `json:"missions,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	PlatformLocs []MtiNewBulkParamsBodyPlatformLoc `json:"platformLocs,omitzero"`
	paramObj
}

func (r MtiNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MtiNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type MtiNewBulkParamsBodyDwell struct {
	// Factor which modifies the value of the reported target latitude (Delta Latitude,
	// field D32.4).
	D10 param.Opt[float64] `json:"d10,omitzero"`
	// Factor which modifies the value of the reported target longitude (Delta
	// Longitude, field D32.5).
	D11 param.Opt[float64] `json:"d11,omitzero"`
	// Standard deviation in the estimated horizontal sensor location at the time of
	// the dwell, measured along the sensor track direction (field D15), in
	// centimeters.
	D12 param.Opt[int64] `json:"d12,omitzero"`
	// Standard deviation in the estimated horizontal sensor location at the time of
	// the dwell, measured orthogonal to the sensor track direction (field D15), in
	// centimeters.
	D13 param.Opt[int64] `json:"d13,omitzero"`
	// Standard deviation of the sensor altitude estimate (field D9), in centimeters.
	D14 param.Opt[int64] `json:"d14,omitzero"`
	// Ground track of the sensor at the time of the dwell, as the angle in degrees
	// (clockwise) from True North.
	D15 param.Opt[float64] `json:"d15,omitzero"`
	// Ground speed of the sensor at the time of the dwell, in millimeters per second.
	D16 param.Opt[int64] `json:"d16,omitzero"`
	// Velocity of the sensor in the vertical direction, in decimeters per second.
	D17 param.Opt[int64] `json:"d17,omitzero"`
	// Standard deviation of the estimate of the sensor track, in degrees.
	D18 param.Opt[int64] `json:"d18,omitzero"`
	// Standard deviation of estimate of the sensor speed, in millimeters per second.
	D19 param.Opt[int64] `json:"d19,omitzero"`
	// Sequential count of a revisit of the bounding area in the last sent Job
	// Definition Segment, where a Revisit Index of '0' indicates the first revisit.
	D2 param.Opt[int64] `json:"d2,omitzero"`
	// Standard deviation of estimate of the sensor vertical velocity, expressed in
	// centimeters per second.
	D20 param.Opt[int64] `json:"d20,omitzero"`
	// Heading of the platform at the time of the dwell, as the angle in degrees
	// (clockwise) from True North to the roll axis of the platform.
	D21 param.Opt[float64] `json:"d21,omitzero"`
	// Pitch angle of the platform at the time of the dwell, in degrees.
	D22 param.Opt[float64] `json:"d22,omitzero"`
	// Roll angle of the platform at the time of the dwell, in degrees.
	D23 param.Opt[float64] `json:"d23,omitzero"`
	// The North-South position of the center of the dwell area, expressed as degrees
	// North (positive) or South (negative) of the Equator.
	D24 param.Opt[float64] `json:"d24,omitzero"`
	// The East-West position of the center of the dwell area, expressed as degrees
	// East (positive, 0 to 180) or West (negative, 0 to -180) of the Prime Meridian.
	D25 param.Opt[float64] `json:"d25,omitzero"`
	// Distance on the earth surface, expressed in kilometers, from the near edge to
	// the center of the dwell area.
	D26 param.Opt[float64] `json:"d26,omitzero"`
	// For dwell based radars, one-half of the 3-dB beamwidth. For non-dwell based
	// radars, the angle between the beginning of the dwell to the center of the dwell.
	// Measured in degrees.
	D27 param.Opt[float64] `json:"d27,omitzero"`
	// Rotation of the sensor broadside face about the local vertical axis of the
	// platform, in degrees.
	D28 param.Opt[float64] `json:"d28,omitzero"`
	// Rotation angle of the sensor about the transverse axis of the sensor broadside,
	// in degrees.
	D29 param.Opt[float64] `json:"d29,omitzero"`
	// Temporally sequential count of a dwell within the revisit of a particular
	// bounding area for a given job ID.
	D3 param.Opt[int64] `json:"d3,omitzero"`
	// Rotation angle of the sensor about the transverse axis of the sensor broadside,
	// in degrees.
	D30 param.Opt[float64] `json:"d30,omitzero"`
	// Minimum velocity component, along the line of sight, which can be detected by
	// the sensor, in decimeters per second.
	D31 param.Opt[int64] `json:"d31,omitzero"`
	// Flag indicating the last dwell of the revisit.
	D4 param.Opt[bool] `json:"d4,omitzero"`
	// Count of the total number of targets reported during this dwell and sent in this
	// Dwell Segment.
	D5 param.Opt[int64] `json:"d5,omitzero"`
	// Elapsed time, expressed in milliseconds, from midnight at the beginning of the
	// day specified in the Reference Time fields (missionRefTime) of the Mission
	// Segment.
	D6 param.Opt[int64] `json:"d6,omitzero"`
	// North-South position of the sensor at the temporal center of the dwell, in
	// degrees.
	D7 param.Opt[float64] `json:"d7,omitzero"`
	// The East-West position of the sensor at the temporal center of the dwell, in
	// degrees East (positive, 0 to 180) or West (negative, 0 to -180) of the Prime
	// Meridian.
	D8 param.Opt[float64] `json:"d8,omitzero"`
	// The altitude of the sensor at temporal center of the dwell, above the WGS 84
	// ellipsoid, expressed in centimeters.
	D9 param.Opt[int64] `json:"d9,omitzero"`
	// Dwell timestamp in ISO8601 UTC format.
	Dwellts param.Opt[time.Time] `json:"dwellts,omitzero" format:"date-time"`
	// Minimum velocity component, along the line of sight, which can be detected by
	// the sensor, in decimeters per second.
	D32 []MtiNewBulkParamsBodyDwellD32 `json:"d32,omitzero"`
	paramObj
}

func (r MtiNewBulkParamsBodyDwell) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBodyDwell
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBodyDwell) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A grouping of zero or more target reports for which the sensor provides a single
// time, sensor position, reference position on the ground with simple estimates
// for the observed area at the reported time, and other pertinent data.
type MtiNewBulkParamsBodyDwellD32 struct {
	// Sequential count of this MTI report within the dwell.
	D32_1 param.Opt[int64] `json:"d32_1,omitzero"`
	// The classification of the target (i.e. vehicle, aircraft, …).
	D32_10 param.Opt[string] `json:"d32_10,omitzero"`
	// Estimated probability that the target classification field is correctly
	// classified.
	D32_11 param.Opt[int64] `json:"d32_11,omitzero"`
	// Standard deviation of the estimated slant range of the reported detection, in
	// centimeters.
	D32_12 param.Opt[int64] `json:"d32_12,omitzero"`
	// Standard deviation of the position estimate, in the cross-range direction, of
	// the reported detection, in decimeters.
	D32_13 param.Opt[int64] `json:"d32_13,omitzero"`
	// Standard deviation of the estimated geodetic height, in meters.
	D32_14 param.Opt[int64] `json:"d32_14,omitzero"`
	// Standard deviation of the measured line-of-sight velocity component, in
	// centimeters per second.
	D32_15 param.Opt[int64] `json:"d32_15,omitzero"`
	// The Truth Tag- Application is the Application Field truncated to 8 bits, from
	// the Entity State Protocol Data Unit (PDU) used to generate the MTI Target.
	D32_16 param.Opt[int64] `json:"d32_16,omitzero"`
	// The Truth Tag - Entity is the Entity Field from the Entity State PDU used to
	// generate the MTI Target.
	D32_17 param.Opt[int64] `json:"d32_17,omitzero"`
	// Estimated radar cross section of the target return, in half-decibels.
	D32_18 param.Opt[int64] `json:"d32_18,omitzero"`
	// The North-South position of the reported detection, expressed as degrees North
	// (positive) or South (negative) of the Equator.
	D32_2 param.Opt[float64] `json:"d32_2,omitzero"`
	// The East-West position of the reported detection, expressed as degrees East
	// (positive) from the Prime Meridian.
	D32_3 param.Opt[float64] `json:"d32_3,omitzero"`
	// The North-South position of the reported detection, expressed as degrees North
	// (positive) or South (negative) from the Dwell Area Center Latitude.
	D32_4 param.Opt[int64] `json:"d32_4,omitzero"`
	// The East-West position of the reported detection, expressed as degrees East
	// (positive, 0 to 180) or West (negative, 0 to -180) of the Prime Meridian from
	// the Dwell Area Center Longitude.
	D32_5 param.Opt[int64] `json:"d32_5,omitzero"`
	// Height of the reported detection, referenced to its position above the WGS 84
	// ellipsoid, in meters.
	D32_6 param.Opt[int64] `json:"d32_6,omitzero"`
	// The component of velocity for the reported detection, expressed in centimeters
	// per second, corrected for platform motion, along the line of sight between the
	// sensor and the reported detection, where the positive direction is away from the
	// sensor.
	D32_7 param.Opt[int64] `json:"d32_7,omitzero"`
	// The target wrap velocity permits trackers to un-wrap velocities for targets with
	// line-of-sight components large enough to exceed the first velocity period.
	// Expressed in centimeters/sec.
	D32_8 param.Opt[int64] `json:"d32_8,omitzero"`
	// Estimated signal-to-noise ratio (SNR) of the target return, in decibels.
	D32_9 param.Opt[int64] `json:"d32_9,omitzero"`
	paramObj
}

func (r MtiNewBulkParamsBodyDwellD32) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBodyDwellD32
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBodyDwellD32) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiNewBulkParamsBodyFreeText struct {
	// The originator of the Free Text message.
	F1 param.Opt[string] `json:"f1,omitzero"`
	// The recipient for which the Free Text message is intended.
	F2 param.Opt[string] `json:"f2,omitzero"`
	// Free text data message.
	F3 param.Opt[string] `json:"f3,omitzero"`
	paramObj
}

func (r MtiNewBulkParamsBodyFreeText) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBodyFreeText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBodyFreeText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiNewBulkParamsBodyHrr struct {
	// Detection threshold used to isolate significant target scatterer pixels,
	// expressed as power relative to clutter mean in negative quarter-decibels.
	H10 param.Opt[int64] `json:"h10,omitzero"`
	// 3dB range impulse response of the radar, expressed in centimeters.
	H11 param.Opt[float64] `json:"h11,omitzero"`
	// Slant Range pixel spacing after over sampling, expressed in centimeters.
	H12 param.Opt[float64] `json:"h12,omitzero"`
	// 3dB Doppler resolution of the radar, expressed in Hertz.
	H13 param.Opt[float64] `json:"h13,omitzero"`
	// Doppler pixel spacing after over sampling, expressed in Hertz.
	H14 param.Opt[float64] `json:"h14,omitzero"`
	// Center Frequency of the radar in GHz.
	H15 param.Opt[float64] `json:"h15,omitzero"`
	// Enumeration table denoting the compression technique used.
	H16 param.Opt[string] `json:"h16,omitzero"`
	// Enumeration table indicating the spectral weighting used in the range
	// compression process.
	H17 param.Opt[string] `json:"h17,omitzero"`
	// Enumeration table indicating the spectral weighting used in the cross-range or
	// Doppler compression process.
	H18 param.Opt[string] `json:"h18,omitzero"`
	// Initial power of the peak scatterer, expressed in dB.
	H19 param.Opt[float64] `json:"h19,omitzero"`
	// Sequential count of a revisit of the bounding area for a given job ID.
	H2 param.Opt[int64] `json:"h2,omitzero"`
	// RCS of the peak scatterer, expressed in half-decibels (dB/2).
	H20 param.Opt[int64] `json:"h20,omitzero"`
	// When the RDM does not correlate to a single MTI report index or when the center
	// range bin does not correlate to the center of the dwell; provide the range
	// sample offset in meters from Dwell Center (positive is away from the sensor) of
	// the first scatterer record.
	H21 param.Opt[int64] `json:"h21,omitzero"`
	// When the RDM does not correlate to a single MTI report index or the center
	// doppler bin does not correlate to the doppler centroid of the dwell; Doppler
	// sample value in Hz of the first scatterer record.
	H22 param.Opt[int64] `json:"h22,omitzero"`
	// Enumeration field which designates the type of data being delivered.
	H23 param.Opt[string] `json:"h23,omitzero"`
	// Flag field to indicate the additional signal processing techniques applied to
	// the data.
	H24 param.Opt[string] `json:"h24,omitzero"`
	// Number of pixels in the range dimension of the chip.
	H27 param.Opt[int64] `json:"h27,omitzero"`
	// Distance from Range Bin to closest edge in the entire chip, expressed in
	// centimeters.
	H28 param.Opt[int64] `json:"h28,omitzero"`
	// Relative velocity to skin line.
	H29 param.Opt[int64] `json:"h29,omitzero"`
	// Sequential count of a dwell within the revisit of a particular bounding area for
	// a given job ID.
	H3 param.Opt[int64] `json:"h3,omitzero"`
	// Computed object length based upon HRR profile, in meters.
	H30 param.Opt[int64] `json:"h30,omitzero"`
	// Standard deviation of estimate of the object length, expressed in meters.
	H31 param.Opt[int64] `json:"h31,omitzero"`
	// Flag to indicate the last dwell of the revisit.
	H4 param.Opt[bool] `json:"h4,omitzero"`
	// Sequential index of the associated MTI Report.
	H5 param.Opt[int64] `json:"h5,omitzero"`
	// Number of Range Doppler pixels that exceed target scatterer threshold and are
	// reported in this segment.
	H6 param.Opt[int64] `json:"h6,omitzero"`
	// Number of Range Bins/Samples in a Range Doppler Chip.
	H7 param.Opt[int64] `json:"h7,omitzero"`
	// Number of Doppler bins in a Range-Doppler chip.
	H8 param.Opt[int64] `json:"h8,omitzero"`
	// The Peak Scatter returns the maximum power level (e.g. in milliwatts, or dBm)
	// registered by the sensor.
	H9 param.Opt[int64] `json:"h9,omitzero"`
	// Standard deviation of estimate of the object length, expressed in meters.
	H32 []MtiNewBulkParamsBodyHrrH32 `json:"h32,omitzero"`
	paramObj
}

func (r MtiNewBulkParamsBodyHrr) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBodyHrr
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBodyHrr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HRR Scatterer record for a target pixel that exceeds the target detection
// threshold.
type MtiNewBulkParamsBodyHrrH32 struct {
	// Scatterer’s power magnitude.
	H32_1 param.Opt[int64] `json:"h32_1,omitzero"`
	// Scatterer’s complex phase, in degrees.
	H32_2 param.Opt[int64] `json:"h32_2,omitzero"`
	// Scatterer’s Range index relative to Range-Doppler chip, where increasing index
	// equates to increasing range.
	H32_3 param.Opt[int64] `json:"h32_3,omitzero"`
	// Scatterer’s Doppler index relative to Range-Doppler chip, where increasing index
	// equates to increasing Doppler.
	H32_4 param.Opt[int64] `json:"h32_4,omitzero"`
	paramObj
}

func (r MtiNewBulkParamsBodyHrrH32) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBodyHrrH32
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBodyHrrH32) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The means for the platform to pass information pertaining to the sensor job that
// will be performed and details of the location parameters (terrain elevation
// model and geoid model) used in the measurement.
type MtiNewBulkParamsBodyJobDef struct {
	// A platform assigned number identifying the specific request or task to which the
	// specific dwell pertains.
	J1 param.Opt[int64] `json:"j1,omitzero"`
	// North-South position of the third corner (Point C) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J10 param.Opt[float64] `json:"j10,omitzero"`
	// East-West position of the third corner (Point C) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J11 param.Opt[float64] `json:"j11,omitzero"`
	// North-South position of the fourth corner (Point D) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J12 param.Opt[float64] `json:"j12,omitzero"`
	// East-West position of the fourth corner (Point D) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J13 param.Opt[float64] `json:"j13,omitzero"`
	// Mode in which the radar will operate for the given job ID.
	J14 param.Opt[string] `json:"j14,omitzero"`
	// The nominal revisit interval for the job ID, expressed in deciseconds. Value of
	// zero, indicates that the sensor is not revisiting the previous area.
	J15 param.Opt[int64] `json:"j15,omitzero"`
	// Nominal estimate of the standard deviation in the estimated horizontal (along
	// track) sensor location, expressed in decimeters. measured along the sensor track
	// direction defined in the Dwell segment.
	J16 param.Opt[int64] `json:"j16,omitzero"`
	// Nominal estimate of the standard deviation in the estimated horizontal sensor
	// location, measured orthogonal to the track direction, expressed in decimeters.
	J17 param.Opt[int64] `json:"j17,omitzero"`
	// Nominal estimate of the standard deviation of the measured sensor altitude,
	// expressed in decimeters.
	J18 param.Opt[int64] `json:"j18,omitzero"`
	// Standard deviation of the estimate of sensor track heading, expressed in
	// degrees.
	J19 param.Opt[int64] `json:"j19,omitzero"`
	// The type of sensor or the platform.
	J2 param.Opt[string] `json:"j2,omitzero"`
	// Nominal standard deviation of the estimate of sensor speed, expressed in
	// millimeters per second.
	J20 param.Opt[int64] `json:"j20,omitzero"`
	// Nominal standard deviation of the slant range of the reported detection,
	// expressed in centimeters.
	J21 param.Opt[int64] `json:"j21,omitzero"`
	// Nominal standard deviation of the measured cross angle to the reported
	// detection, expressed in degrees.
	J22 param.Opt[float64] `json:"j22,omitzero"`
	// Nominal standard deviation of the velocity line-of-sight component, expressed in
	// centimeters per second.
	J23 param.Opt[int64] `json:"j23,omitzero"`
	// Nominal minimum velocity component along the line of sight, which can be
	// detected by the sensor, expressed in decimeters per second.
	J24 param.Opt[int64] `json:"j24,omitzero"`
	// Nominal probability that an unobscured ten square-meter target will be detected
	// within the given area of surveillance.
	J25 param.Opt[int64] `json:"j25,omitzero"`
	// The expected density of False Alarms (FA), expressed as the negative of the
	// decibel value.
	J26 param.Opt[int64] `json:"j26,omitzero"`
	// The terrain elevation model used for developing the target reports.
	J27 param.Opt[string] `json:"j27,omitzero"`
	// The geoid model used for developing the target reports.
	J28 param.Opt[string] `json:"j28,omitzero"`
	// Identifier of the particular variant of the sensor type.
	J3 param.Opt[string] `json:"j3,omitzero"`
	// Flag field indicating whether filtering has been applied to the targets detected
	// within the dwell area.
	J4 param.Opt[int64] `json:"j4,omitzero"`
	// Priority of this tasking request relative to all other active tasking requests
	// scheduled for execution on the specified platform.
	J5 param.Opt[int64] `json:"j5,omitzero"`
	// North-South position of the first corner (Point A) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J6 param.Opt[float64] `json:"j6,omitzero"`
	// East-West position of the first corner (Point A) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J7 param.Opt[float64] `json:"j7,omitzero"`
	// North-South position of the second corner (Point B) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J8 param.Opt[float64] `json:"j8,omitzero"`
	// East-West position of the second corner (Point B) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J9 param.Opt[float64] `json:"j9,omitzero"`
	paramObj
}

func (r MtiNewBulkParamsBodyJobDef) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBodyJobDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBodyJobDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiNewBulkParamsBodyJobRequest struct {
	// Specifies the Earliest Start Time for which the service is requested. Composite
	// of fields R15-R20.
	JobReqEst param.Opt[time.Time] `json:"jobReqEst,omitzero" format:"date-time"`
	// The requestor of the sensor service.
	R1 param.Opt[string] `json:"r1,omitzero"`
	// North-South position of the fourth corner (Point D) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R10 param.Opt[float64] `json:"r10,omitzero"`
	// East-West position of the fourth corner (Point D) defining the requested area
	// for service, expressed as degrees East (positive, 0 to 180) or West (negative, 0
	// to -180) of the Prime Meridian.
	R11 param.Opt[float64] `json:"r11,omitzero"`
	// Identifies the radar mode requested by the requestor.
	R12 param.Opt[string] `json:"r12,omitzero"`
	// Specifies the radar range resolution requested by the requestor, expressed in
	// centimeters.
	R13 param.Opt[int64] `json:"r13,omitzero"`
	// Specifies the radar cross-range resolution requested by the requestor, expressed
	// in decimeters.
	R14 param.Opt[int64] `json:"r14,omitzero"`
	// Identifier for the tasking message sent by the requesting station.
	R2 param.Opt[string] `json:"r2,omitzero"`
	// Specifies the maximum time from the requested start time after which the request
	// is to be abandoned, expressed in seconds.
	R21 param.Opt[int64] `json:"r21,omitzero"`
	// Specifies the time duration for the radar job, measured from the actual start of
	// the job, expressed in seconds.
	R22 param.Opt[int64] `json:"r22,omitzero"`
	// Specifies the revisit interval for the radar job, expressed in deciseconds.
	R23 param.Opt[int64] `json:"r23,omitzero"`
	// the type of sensor or the platform.
	R24 param.Opt[string] `json:"r24,omitzero"`
	// The particular variant of the sensor type.
	R25 param.Opt[string] `json:"r25,omitzero"`
	// Flag field indicating that it is an initial request (flag = 0) or the desire of
	// the requestor to cancel (flag = 1) the requested job.
	R26 param.Opt[bool] `json:"r26,omitzero"`
	// The priority of the request relative to other requests originated by the
	// requesting station.
	R3 param.Opt[int64] `json:"r3,omitzero"`
	// North-South position of the first corner (Point A) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R4 param.Opt[float64] `json:"r4,omitzero"`
	// East-West position of the first corner (Point A) defining the requested area for
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	R5 param.Opt[float64] `json:"r5,omitzero"`
	// North-South position of the second corner (Point B) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R6 param.Opt[float64] `json:"r6,omitzero"`
	// East-West position of the second corner (Point B) defining the requested area
	// for service, expressed as degrees East (positive, 0 to 180) or West (negative, 0
	// to -180) of the Prime Meridian.
	R7 param.Opt[float64] `json:"r7,omitzero"`
	// North-South position of the third corner (Point C) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R8 param.Opt[float64] `json:"r8,omitzero"`
	// East-West position of the third corner (Point C) defining the requested area for
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	R9 param.Opt[float64] `json:"r9,omitzero"`
	paramObj
}

func (r MtiNewBulkParamsBodyJobRequest) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBodyJobRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBodyJobRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiNewBulkParamsBodyMission struct {
	// The mission plan id.
	M1 param.Opt[string] `json:"m1,omitzero"`
	// Unique identification of the flight plan.
	M2 param.Opt[string] `json:"m2,omitzero"`
	// Platform type that originated the data.
	M3 param.Opt[string] `json:"m3,omitzero"`
	// Identification of the platform variant, modifications, etc.
	M4 param.Opt[string] `json:"m4,omitzero"`
	// Mission origination date.
	MsnRefTs param.Opt[time.Time] `json:"msnRefTs,omitzero" format:"date"`
	paramObj
}

func (r MtiNewBulkParamsBodyMission) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBodyMission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBodyMission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiNewBulkParamsBodyPlatformLoc struct {
	// Elapsed time, expressed in milliseconds, from midnight at the beginning of the
	// day specified in the Reference Time fields of the Mission Segment to the time
	// the report is prepared.
	L1 param.Opt[int64] `json:"l1,omitzero"`
	// North-South position of the platform at the time the report is prepared,
	// expressed as degrees North (positive) or South (negative) of the Equator.
	L2 param.Opt[float64] `json:"l2,omitzero"`
	// East-West position of the platform at the time the report is prepared, expressed
	// as degrees East (positive) from the Prime Meridian.
	L3 param.Opt[float64] `json:"l3,omitzero"`
	// Altitude of the platform at the time the report is prepared, referenced to its
	// position above the WGS-84 ellipsoid, in centimeters.
	L4 param.Opt[int64] `json:"l4,omitzero"`
	// Ground track of the platform at the time the report is prepared, expressed as
	// the angle in degrees (clockwise) from True North.
	L5 param.Opt[float64] `json:"l5,omitzero"`
	// Ground speed of the platform at the time the report is prepared, expressed as
	// millimeters per second.
	L6 param.Opt[int64] `json:"l6,omitzero"`
	// Velocity of the platform in the vertical direction, expressed as decimeters per
	// second.
	L7 param.Opt[int64] `json:"l7,omitzero"`
	// Platform location timestamp in ISO8601 UTC format.
	Platlocts param.Opt[time.Time] `json:"platlocts,omitzero" format:"date-time"`
	paramObj
}

func (r MtiNewBulkParamsBodyPlatformLoc) MarshalJSON() (data []byte, err error) {
	type shadow MtiNewBulkParamsBodyPlatformLoc
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiNewBulkParamsBodyPlatformLoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time the row was created in the database. (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MtiTupleParams]'s query parameters as `url.Values`.
func (r MtiTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MtiUnvalidatedPublishParams struct {
	Body []MtiUnvalidatedPublishParamsBody
	paramObj
}

func (r MtiUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *MtiUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// Information on the mission and flight plans, the type and configuration of the
// platform, and the reference time.
//
// The properties ClassificationMarking, DataMode, Source are required.
type MtiUnvalidatedPublishParamsBody struct {
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
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	P10 param.Opt[int64] `json:"p10,omitzero"`
	// Nationality of the platform providing the data.
	P3 param.Opt[string] `json:"p3,omitzero"`
	// Control / handling marking.
	P6 param.Opt[string] `json:"p6,omitzero"`
	// Data record exercise indicator.
	P7 param.Opt[string] `json:"p7,omitzero"`
	// ID of the platform providing the data (tail number for air platform, name and
	// numerical designator for space-based platforms).
	P8 param.Opt[string] `json:"p8,omitzero"`
	// Integer field, assigned by the platform, that uniquely identifies the mission
	// for the platform.
	P9 param.Opt[int64] `json:"p9,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	Dwells []MtiUnvalidatedPublishParamsBodyDwell `json:"dwells,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	FreeTexts []MtiUnvalidatedPublishParamsBodyFreeText `json:"freeTexts,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	Hrrs []MtiUnvalidatedPublishParamsBodyHrr `json:"hrrs,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	JobDefs []MtiUnvalidatedPublishParamsBodyJobDef `json:"jobDefs,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	JobRequests []MtiUnvalidatedPublishParamsBodyJobRequest `json:"jobRequests,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	Missions []MtiUnvalidatedPublishParamsBodyMission `json:"missions,omitzero"`
	// A platform-assigned number identifying the specific request or task that
	// pertains to all Dwell, HRR, and Range-Doppler segments in the packet. Job ID is
	// unique within a mission.
	PlatformLocs []MtiUnvalidatedPublishParamsBodyPlatformLoc `json:"platformLocs,omitzero"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MtiUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type MtiUnvalidatedPublishParamsBodyDwell struct {
	// Factor which modifies the value of the reported target latitude (Delta Latitude,
	// field D32.4).
	D10 param.Opt[float64] `json:"d10,omitzero"`
	// Factor which modifies the value of the reported target longitude (Delta
	// Longitude, field D32.5).
	D11 param.Opt[float64] `json:"d11,omitzero"`
	// Standard deviation in the estimated horizontal sensor location at the time of
	// the dwell, measured along the sensor track direction (field D15), in
	// centimeters.
	D12 param.Opt[int64] `json:"d12,omitzero"`
	// Standard deviation in the estimated horizontal sensor location at the time of
	// the dwell, measured orthogonal to the sensor track direction (field D15), in
	// centimeters.
	D13 param.Opt[int64] `json:"d13,omitzero"`
	// Standard deviation of the sensor altitude estimate (field D9), in centimeters.
	D14 param.Opt[int64] `json:"d14,omitzero"`
	// Ground track of the sensor at the time of the dwell, as the angle in degrees
	// (clockwise) from True North.
	D15 param.Opt[float64] `json:"d15,omitzero"`
	// Ground speed of the sensor at the time of the dwell, in millimeters per second.
	D16 param.Opt[int64] `json:"d16,omitzero"`
	// Velocity of the sensor in the vertical direction, in decimeters per second.
	D17 param.Opt[int64] `json:"d17,omitzero"`
	// Standard deviation of the estimate of the sensor track, in degrees.
	D18 param.Opt[int64] `json:"d18,omitzero"`
	// Standard deviation of estimate of the sensor speed, in millimeters per second.
	D19 param.Opt[int64] `json:"d19,omitzero"`
	// Sequential count of a revisit of the bounding area in the last sent Job
	// Definition Segment, where a Revisit Index of '0' indicates the first revisit.
	D2 param.Opt[int64] `json:"d2,omitzero"`
	// Standard deviation of estimate of the sensor vertical velocity, expressed in
	// centimeters per second.
	D20 param.Opt[int64] `json:"d20,omitzero"`
	// Heading of the platform at the time of the dwell, as the angle in degrees
	// (clockwise) from True North to the roll axis of the platform.
	D21 param.Opt[float64] `json:"d21,omitzero"`
	// Pitch angle of the platform at the time of the dwell, in degrees.
	D22 param.Opt[float64] `json:"d22,omitzero"`
	// Roll angle of the platform at the time of the dwell, in degrees.
	D23 param.Opt[float64] `json:"d23,omitzero"`
	// The North-South position of the center of the dwell area, expressed as degrees
	// North (positive) or South (negative) of the Equator.
	D24 param.Opt[float64] `json:"d24,omitzero"`
	// The East-West position of the center of the dwell area, expressed as degrees
	// East (positive, 0 to 180) or West (negative, 0 to -180) of the Prime Meridian.
	D25 param.Opt[float64] `json:"d25,omitzero"`
	// Distance on the earth surface, expressed in kilometers, from the near edge to
	// the center of the dwell area.
	D26 param.Opt[float64] `json:"d26,omitzero"`
	// For dwell based radars, one-half of the 3-dB beamwidth. For non-dwell based
	// radars, the angle between the beginning of the dwell to the center of the dwell.
	// Measured in degrees.
	D27 param.Opt[float64] `json:"d27,omitzero"`
	// Rotation of the sensor broadside face about the local vertical axis of the
	// platform, in degrees.
	D28 param.Opt[float64] `json:"d28,omitzero"`
	// Rotation angle of the sensor about the transverse axis of the sensor broadside,
	// in degrees.
	D29 param.Opt[float64] `json:"d29,omitzero"`
	// Temporally sequential count of a dwell within the revisit of a particular
	// bounding area for a given job ID.
	D3 param.Opt[int64] `json:"d3,omitzero"`
	// Rotation angle of the sensor about the transverse axis of the sensor broadside,
	// in degrees.
	D30 param.Opt[float64] `json:"d30,omitzero"`
	// Minimum velocity component, along the line of sight, which can be detected by
	// the sensor, in decimeters per second.
	D31 param.Opt[int64] `json:"d31,omitzero"`
	// Flag indicating the last dwell of the revisit.
	D4 param.Opt[bool] `json:"d4,omitzero"`
	// Count of the total number of targets reported during this dwell and sent in this
	// Dwell Segment.
	D5 param.Opt[int64] `json:"d5,omitzero"`
	// Elapsed time, expressed in milliseconds, from midnight at the beginning of the
	// day specified in the Reference Time fields (missionRefTime) of the Mission
	// Segment.
	D6 param.Opt[int64] `json:"d6,omitzero"`
	// North-South position of the sensor at the temporal center of the dwell, in
	// degrees.
	D7 param.Opt[float64] `json:"d7,omitzero"`
	// The East-West position of the sensor at the temporal center of the dwell, in
	// degrees East (positive, 0 to 180) or West (negative, 0 to -180) of the Prime
	// Meridian.
	D8 param.Opt[float64] `json:"d8,omitzero"`
	// The altitude of the sensor at temporal center of the dwell, above the WGS 84
	// ellipsoid, expressed in centimeters.
	D9 param.Opt[int64] `json:"d9,omitzero"`
	// Dwell timestamp in ISO8601 UTC format.
	Dwellts param.Opt[time.Time] `json:"dwellts,omitzero" format:"date-time"`
	// Minimum velocity component, along the line of sight, which can be detected by
	// the sensor, in decimeters per second.
	D32 []MtiUnvalidatedPublishParamsBodyDwellD32 `json:"d32,omitzero"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBodyDwell) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBodyDwell
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBodyDwell) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A grouping of zero or more target reports for which the sensor provides a single
// time, sensor position, reference position on the ground with simple estimates
// for the observed area at the reported time, and other pertinent data.
type MtiUnvalidatedPublishParamsBodyDwellD32 struct {
	// Sequential count of this MTI report within the dwell.
	D32_1 param.Opt[int64] `json:"d32_1,omitzero"`
	// The classification of the target (i.e. vehicle, aircraft, …).
	D32_10 param.Opt[string] `json:"d32_10,omitzero"`
	// Estimated probability that the target classification field is correctly
	// classified.
	D32_11 param.Opt[int64] `json:"d32_11,omitzero"`
	// Standard deviation of the estimated slant range of the reported detection, in
	// centimeters.
	D32_12 param.Opt[int64] `json:"d32_12,omitzero"`
	// Standard deviation of the position estimate, in the cross-range direction, of
	// the reported detection, in decimeters.
	D32_13 param.Opt[int64] `json:"d32_13,omitzero"`
	// Standard deviation of the estimated geodetic height, in meters.
	D32_14 param.Opt[int64] `json:"d32_14,omitzero"`
	// Standard deviation of the measured line-of-sight velocity component, in
	// centimeters per second.
	D32_15 param.Opt[int64] `json:"d32_15,omitzero"`
	// The Truth Tag- Application is the Application Field truncated to 8 bits, from
	// the Entity State Protocol Data Unit (PDU) used to generate the MTI Target.
	D32_16 param.Opt[int64] `json:"d32_16,omitzero"`
	// The Truth Tag - Entity is the Entity Field from the Entity State PDU used to
	// generate the MTI Target.
	D32_17 param.Opt[int64] `json:"d32_17,omitzero"`
	// Estimated radar cross section of the target return, in half-decibels.
	D32_18 param.Opt[int64] `json:"d32_18,omitzero"`
	// The North-South position of the reported detection, expressed as degrees North
	// (positive) or South (negative) of the Equator.
	D32_2 param.Opt[float64] `json:"d32_2,omitzero"`
	// The East-West position of the reported detection, expressed as degrees East
	// (positive) from the Prime Meridian.
	D32_3 param.Opt[float64] `json:"d32_3,omitzero"`
	// The North-South position of the reported detection, expressed as degrees North
	// (positive) or South (negative) from the Dwell Area Center Latitude.
	D32_4 param.Opt[int64] `json:"d32_4,omitzero"`
	// The East-West position of the reported detection, expressed as degrees East
	// (positive, 0 to 180) or West (negative, 0 to -180) of the Prime Meridian from
	// the Dwell Area Center Longitude.
	D32_5 param.Opt[int64] `json:"d32_5,omitzero"`
	// Height of the reported detection, referenced to its position above the WGS 84
	// ellipsoid, in meters.
	D32_6 param.Opt[int64] `json:"d32_6,omitzero"`
	// The component of velocity for the reported detection, expressed in centimeters
	// per second, corrected for platform motion, along the line of sight between the
	// sensor and the reported detection, where the positive direction is away from the
	// sensor.
	D32_7 param.Opt[int64] `json:"d32_7,omitzero"`
	// The target wrap velocity permits trackers to un-wrap velocities for targets with
	// line-of-sight components large enough to exceed the first velocity period.
	// Expressed in centimeters/sec.
	D32_8 param.Opt[int64] `json:"d32_8,omitzero"`
	// Estimated signal-to-noise ratio (SNR) of the target return, in decibels.
	D32_9 param.Opt[int64] `json:"d32_9,omitzero"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBodyDwellD32) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBodyDwellD32
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBodyDwellD32) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiUnvalidatedPublishParamsBodyFreeText struct {
	// The originator of the Free Text message.
	F1 param.Opt[string] `json:"f1,omitzero"`
	// The recipient for which the Free Text message is intended.
	F2 param.Opt[string] `json:"f2,omitzero"`
	// Free text data message.
	F3 param.Opt[string] `json:"f3,omitzero"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBodyFreeText) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBodyFreeText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBodyFreeText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiUnvalidatedPublishParamsBodyHrr struct {
	// Detection threshold used to isolate significant target scatterer pixels,
	// expressed as power relative to clutter mean in negative quarter-decibels.
	H10 param.Opt[int64] `json:"h10,omitzero"`
	// 3dB range impulse response of the radar, expressed in centimeters.
	H11 param.Opt[float64] `json:"h11,omitzero"`
	// Slant Range pixel spacing after over sampling, expressed in centimeters.
	H12 param.Opt[float64] `json:"h12,omitzero"`
	// 3dB Doppler resolution of the radar, expressed in Hertz.
	H13 param.Opt[float64] `json:"h13,omitzero"`
	// Doppler pixel spacing after over sampling, expressed in Hertz.
	H14 param.Opt[float64] `json:"h14,omitzero"`
	// Center Frequency of the radar in GHz.
	H15 param.Opt[float64] `json:"h15,omitzero"`
	// Enumeration table denoting the compression technique used.
	H16 param.Opt[string] `json:"h16,omitzero"`
	// Enumeration table indicating the spectral weighting used in the range
	// compression process.
	H17 param.Opt[string] `json:"h17,omitzero"`
	// Enumeration table indicating the spectral weighting used in the cross-range or
	// Doppler compression process.
	H18 param.Opt[string] `json:"h18,omitzero"`
	// Initial power of the peak scatterer, expressed in dB.
	H19 param.Opt[float64] `json:"h19,omitzero"`
	// Sequential count of a revisit of the bounding area for a given job ID.
	H2 param.Opt[int64] `json:"h2,omitzero"`
	// RCS of the peak scatterer, expressed in half-decibels (dB/2).
	H20 param.Opt[int64] `json:"h20,omitzero"`
	// When the RDM does not correlate to a single MTI report index or when the center
	// range bin does not correlate to the center of the dwell; provide the range
	// sample offset in meters from Dwell Center (positive is away from the sensor) of
	// the first scatterer record.
	H21 param.Opt[int64] `json:"h21,omitzero"`
	// When the RDM does not correlate to a single MTI report index or the center
	// doppler bin does not correlate to the doppler centroid of the dwell; Doppler
	// sample value in Hz of the first scatterer record.
	H22 param.Opt[int64] `json:"h22,omitzero"`
	// Enumeration field which designates the type of data being delivered.
	H23 param.Opt[string] `json:"h23,omitzero"`
	// Flag field to indicate the additional signal processing techniques applied to
	// the data.
	H24 param.Opt[string] `json:"h24,omitzero"`
	// Number of pixels in the range dimension of the chip.
	H27 param.Opt[int64] `json:"h27,omitzero"`
	// Distance from Range Bin to closest edge in the entire chip, expressed in
	// centimeters.
	H28 param.Opt[int64] `json:"h28,omitzero"`
	// Relative velocity to skin line.
	H29 param.Opt[int64] `json:"h29,omitzero"`
	// Sequential count of a dwell within the revisit of a particular bounding area for
	// a given job ID.
	H3 param.Opt[int64] `json:"h3,omitzero"`
	// Computed object length based upon HRR profile, in meters.
	H30 param.Opt[int64] `json:"h30,omitzero"`
	// Standard deviation of estimate of the object length, expressed in meters.
	H31 param.Opt[int64] `json:"h31,omitzero"`
	// Flag to indicate the last dwell of the revisit.
	H4 param.Opt[bool] `json:"h4,omitzero"`
	// Sequential index of the associated MTI Report.
	H5 param.Opt[int64] `json:"h5,omitzero"`
	// Number of Range Doppler pixels that exceed target scatterer threshold and are
	// reported in this segment.
	H6 param.Opt[int64] `json:"h6,omitzero"`
	// Number of Range Bins/Samples in a Range Doppler Chip.
	H7 param.Opt[int64] `json:"h7,omitzero"`
	// Number of Doppler bins in a Range-Doppler chip.
	H8 param.Opt[int64] `json:"h8,omitzero"`
	// The Peak Scatter returns the maximum power level (e.g. in milliwatts, or dBm)
	// registered by the sensor.
	H9 param.Opt[int64] `json:"h9,omitzero"`
	// Standard deviation of estimate of the object length, expressed in meters.
	H32 []MtiUnvalidatedPublishParamsBodyHrrH32 `json:"h32,omitzero"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBodyHrr) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBodyHrr
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBodyHrr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HRR Scatterer record for a target pixel that exceeds the target detection
// threshold.
type MtiUnvalidatedPublishParamsBodyHrrH32 struct {
	// Scatterer’s power magnitude.
	H32_1 param.Opt[int64] `json:"h32_1,omitzero"`
	// Scatterer’s complex phase, in degrees.
	H32_2 param.Opt[int64] `json:"h32_2,omitzero"`
	// Scatterer’s Range index relative to Range-Doppler chip, where increasing index
	// equates to increasing range.
	H32_3 param.Opt[int64] `json:"h32_3,omitzero"`
	// Scatterer’s Doppler index relative to Range-Doppler chip, where increasing index
	// equates to increasing Doppler.
	H32_4 param.Opt[int64] `json:"h32_4,omitzero"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBodyHrrH32) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBodyHrrH32
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBodyHrrH32) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The means for the platform to pass information pertaining to the sensor job that
// will be performed and details of the location parameters (terrain elevation
// model and geoid model) used in the measurement.
type MtiUnvalidatedPublishParamsBodyJobDef struct {
	// A platform assigned number identifying the specific request or task to which the
	// specific dwell pertains.
	J1 param.Opt[int64] `json:"j1,omitzero"`
	// North-South position of the third corner (Point C) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J10 param.Opt[float64] `json:"j10,omitzero"`
	// East-West position of the third corner (Point C) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J11 param.Opt[float64] `json:"j11,omitzero"`
	// North-South position of the fourth corner (Point D) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J12 param.Opt[float64] `json:"j12,omitzero"`
	// East-West position of the fourth corner (Point D) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J13 param.Opt[float64] `json:"j13,omitzero"`
	// Mode in which the radar will operate for the given job ID.
	J14 param.Opt[string] `json:"j14,omitzero"`
	// The nominal revisit interval for the job ID, expressed in deciseconds. Value of
	// zero, indicates that the sensor is not revisiting the previous area.
	J15 param.Opt[int64] `json:"j15,omitzero"`
	// Nominal estimate of the standard deviation in the estimated horizontal (along
	// track) sensor location, expressed in decimeters. measured along the sensor track
	// direction defined in the Dwell segment.
	J16 param.Opt[int64] `json:"j16,omitzero"`
	// Nominal estimate of the standard deviation in the estimated horizontal sensor
	// location, measured orthogonal to the track direction, expressed in decimeters.
	J17 param.Opt[int64] `json:"j17,omitzero"`
	// Nominal estimate of the standard deviation of the measured sensor altitude,
	// expressed in decimeters.
	J18 param.Opt[int64] `json:"j18,omitzero"`
	// Standard deviation of the estimate of sensor track heading, expressed in
	// degrees.
	J19 param.Opt[int64] `json:"j19,omitzero"`
	// The type of sensor or the platform.
	J2 param.Opt[string] `json:"j2,omitzero"`
	// Nominal standard deviation of the estimate of sensor speed, expressed in
	// millimeters per second.
	J20 param.Opt[int64] `json:"j20,omitzero"`
	// Nominal standard deviation of the slant range of the reported detection,
	// expressed in centimeters.
	J21 param.Opt[int64] `json:"j21,omitzero"`
	// Nominal standard deviation of the measured cross angle to the reported
	// detection, expressed in degrees.
	J22 param.Opt[float64] `json:"j22,omitzero"`
	// Nominal standard deviation of the velocity line-of-sight component, expressed in
	// centimeters per second.
	J23 param.Opt[int64] `json:"j23,omitzero"`
	// Nominal minimum velocity component along the line of sight, which can be
	// detected by the sensor, expressed in decimeters per second.
	J24 param.Opt[int64] `json:"j24,omitzero"`
	// Nominal probability that an unobscured ten square-meter target will be detected
	// within the given area of surveillance.
	J25 param.Opt[int64] `json:"j25,omitzero"`
	// The expected density of False Alarms (FA), expressed as the negative of the
	// decibel value.
	J26 param.Opt[int64] `json:"j26,omitzero"`
	// The terrain elevation model used for developing the target reports.
	J27 param.Opt[string] `json:"j27,omitzero"`
	// The geoid model used for developing the target reports.
	J28 param.Opt[string] `json:"j28,omitzero"`
	// Identifier of the particular variant of the sensor type.
	J3 param.Opt[string] `json:"j3,omitzero"`
	// Flag field indicating whether filtering has been applied to the targets detected
	// within the dwell area.
	J4 param.Opt[int64] `json:"j4,omitzero"`
	// Priority of this tasking request relative to all other active tasking requests
	// scheduled for execution on the specified platform.
	J5 param.Opt[int64] `json:"j5,omitzero"`
	// North-South position of the first corner (Point A) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J6 param.Opt[float64] `json:"j6,omitzero"`
	// East-West position of the first corner (Point A) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J7 param.Opt[float64] `json:"j7,omitzero"`
	// North-South position of the second corner (Point B) defining the area for sensor
	// service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	J8 param.Opt[float64] `json:"j8,omitzero"`
	// East-West position of the second corner (Point B) defining the area for sensor
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	J9 param.Opt[float64] `json:"j9,omitzero"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBodyJobDef) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBodyJobDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBodyJobDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiUnvalidatedPublishParamsBodyJobRequest struct {
	// Specifies the Earliest Start Time for which the service is requested. Composite
	// of fields R15-R20.
	JobReqEst param.Opt[time.Time] `json:"jobReqEst,omitzero" format:"date-time"`
	// The requestor of the sensor service.
	R1 param.Opt[string] `json:"r1,omitzero"`
	// North-South position of the fourth corner (Point D) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R10 param.Opt[float64] `json:"r10,omitzero"`
	// East-West position of the fourth corner (Point D) defining the requested area
	// for service, expressed as degrees East (positive, 0 to 180) or West (negative, 0
	// to -180) of the Prime Meridian.
	R11 param.Opt[float64] `json:"r11,omitzero"`
	// Identifies the radar mode requested by the requestor.
	R12 param.Opt[string] `json:"r12,omitzero"`
	// Specifies the radar range resolution requested by the requestor, expressed in
	// centimeters.
	R13 param.Opt[int64] `json:"r13,omitzero"`
	// Specifies the radar cross-range resolution requested by the requestor, expressed
	// in decimeters.
	R14 param.Opt[int64] `json:"r14,omitzero"`
	// Identifier for the tasking message sent by the requesting station.
	R2 param.Opt[string] `json:"r2,omitzero"`
	// Specifies the maximum time from the requested start time after which the request
	// is to be abandoned, expressed in seconds.
	R21 param.Opt[int64] `json:"r21,omitzero"`
	// Specifies the time duration for the radar job, measured from the actual start of
	// the job, expressed in seconds.
	R22 param.Opt[int64] `json:"r22,omitzero"`
	// Specifies the revisit interval for the radar job, expressed in deciseconds.
	R23 param.Opt[int64] `json:"r23,omitzero"`
	// the type of sensor or the platform.
	R24 param.Opt[string] `json:"r24,omitzero"`
	// The particular variant of the sensor type.
	R25 param.Opt[string] `json:"r25,omitzero"`
	// Flag field indicating that it is an initial request (flag = 0) or the desire of
	// the requestor to cancel (flag = 1) the requested job.
	R26 param.Opt[bool] `json:"r26,omitzero"`
	// The priority of the request relative to other requests originated by the
	// requesting station.
	R3 param.Opt[int64] `json:"r3,omitzero"`
	// North-South position of the first corner (Point A) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R4 param.Opt[float64] `json:"r4,omitzero"`
	// East-West position of the first corner (Point A) defining the requested area for
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	R5 param.Opt[float64] `json:"r5,omitzero"`
	// North-South position of the second corner (Point B) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R6 param.Opt[float64] `json:"r6,omitzero"`
	// East-West position of the second corner (Point B) defining the requested area
	// for service, expressed as degrees East (positive, 0 to 180) or West (negative, 0
	// to -180) of the Prime Meridian.
	R7 param.Opt[float64] `json:"r7,omitzero"`
	// North-South position of the third corner (Point C) defining the requested area
	// for service, expressed as degrees North (positive) or South (negative) of the
	// Equator.
	R8 param.Opt[float64] `json:"r8,omitzero"`
	// East-West position of the third corner (Point C) defining the requested area for
	// service, expressed as degrees East (positive, 0 to 180) or West (negative, 0 to
	// -180) of the Prime Meridian.
	R9 param.Opt[float64] `json:"r9,omitzero"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBodyJobRequest) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBodyJobRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBodyJobRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiUnvalidatedPublishParamsBodyMission struct {
	// The mission plan id.
	M1 param.Opt[string] `json:"m1,omitzero"`
	// Unique identification of the flight plan.
	M2 param.Opt[string] `json:"m2,omitzero"`
	// Platform type that originated the data.
	M3 param.Opt[string] `json:"m3,omitzero"`
	// Identification of the platform variant, modifications, etc.
	M4 param.Opt[string] `json:"m4,omitzero"`
	// Mission origination date.
	MsnRefTs param.Opt[time.Time] `json:"msnRefTs,omitzero" format:"date"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBodyMission) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBodyMission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBodyMission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MtiUnvalidatedPublishParamsBodyPlatformLoc struct {
	// Elapsed time, expressed in milliseconds, from midnight at the beginning of the
	// day specified in the Reference Time fields of the Mission Segment to the time
	// the report is prepared.
	L1 param.Opt[int64] `json:"l1,omitzero"`
	// North-South position of the platform at the time the report is prepared,
	// expressed as degrees North (positive) or South (negative) of the Equator.
	L2 param.Opt[float64] `json:"l2,omitzero"`
	// East-West position of the platform at the time the report is prepared, expressed
	// as degrees East (positive) from the Prime Meridian.
	L3 param.Opt[float64] `json:"l3,omitzero"`
	// Altitude of the platform at the time the report is prepared, referenced to its
	// position above the WGS-84 ellipsoid, in centimeters.
	L4 param.Opt[int64] `json:"l4,omitzero"`
	// Ground track of the platform at the time the report is prepared, expressed as
	// the angle in degrees (clockwise) from True North.
	L5 param.Opt[float64] `json:"l5,omitzero"`
	// Ground speed of the platform at the time the report is prepared, expressed as
	// millimeters per second.
	L6 param.Opt[int64] `json:"l6,omitzero"`
	// Velocity of the platform in the vertical direction, expressed as decimeters per
	// second.
	L7 param.Opt[int64] `json:"l7,omitzero"`
	// Platform location timestamp in ISO8601 UTC format.
	Platlocts param.Opt[time.Time] `json:"platlocts,omitzero" format:"date-time"`
	paramObj
}

func (r MtiUnvalidatedPublishParamsBodyPlatformLoc) MarshalJSON() (data []byte, err error) {
	type shadow MtiUnvalidatedPublishParamsBodyPlatformLoc
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MtiUnvalidatedPublishParamsBodyPlatformLoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
