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
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// LinkStatusDatalinkService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLinkStatusDatalinkService] method instead.
type LinkStatusDatalinkService struct {
	Options []option.RequestOption
}

// NewLinkStatusDatalinkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLinkStatusDatalinkService(opts ...option.RequestOption) (r LinkStatusDatalinkService) {
	r = LinkStatusDatalinkService{}
	r.Options = opts
	return
}

// Service operation to take a single DataLink record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *LinkStatusDatalinkService) New(ctx context.Context, body LinkStatusDatalinkNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/datalink"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LinkStatusDatalinkService) List(ctx context.Context, query LinkStatusDatalinkListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LinkStatusDatalinkListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/datalink"
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
func (r *LinkStatusDatalinkService) ListAutoPaging(ctx context.Context, query LinkStatusDatalinkListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LinkStatusDatalinkListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LinkStatusDatalinkService) Count(ctx context.Context, query LinkStatusDatalinkCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/datalink/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *LinkStatusDatalinkService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *LinkStatusDatalinkQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/datalink/queryhelp"
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
func (r *LinkStatusDatalinkService) Tuple(ctx context.Context, query LinkStatusDatalinkTupleParams, opts ...option.RequestOption) (res *[]LinkStatusDatalinkTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/datalink/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple datalink records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *LinkStatusDatalinkService) UnvalidatedPublish(ctx context.Context, body LinkStatusDatalinkUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-datalink"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Beta Version DataLink: Detailed instructions regarding the operations of data
// links.
//
// The properties ClassificationMarking, DataMode, OpExName, Originator, Source,
// StartTime are required.
type DatalinkIngestParam struct {
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
	DataMode DatalinkIngestDataMode `json:"dataMode,omitzero,required"`
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExName string `json:"opExName,required"`
	// The identifier of the originator of this message.
	Originator string `json:"originator,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the effective time period of this data link message, in ISO 8601
	// UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Flag Indicating if formal acknowledgement is required for the particular data
	// link message being sent.
	AckReq param.Opt[bool] `json:"ackReq,omitzero"`
	// Maximum altitude difference between two air tracks, in thousands of feet.
	// Required if sysDefaultCode field is "MAN". Allowable entires are 5 to 50 in
	// increments of 5000 feet.
	AltDiff param.Opt[int64] `json:"altDiff,omitzero"`
	// The identifier for this data link message cancellation.
	CanxID param.Opt[string] `json:"canxId,omitzero"`
	// The originator of this data link message cancellation.
	CanxOriginator param.Opt[string] `json:"canxOriginator,omitzero"`
	// Serial number assigned to this data link message cancellation.
	CanxSerialNum param.Opt[string] `json:"canxSerialNum,omitzero"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this data link message cancellation.
	CanxSpecialNotation param.Opt[string] `json:"canxSpecialNotation,omitzero"`
	// Timestamp of the data link message cancellation, in ISO 8601 UTC format with
	// millisecond precision.
	CanxTs param.Opt[time.Time] `json:"canxTs,omitzero" format:"date-time"`
	// Markings that define the source material or the original classification
	// authority for this data link message.
	ClassSource param.Opt[string] `json:"classSource,omitzero"`
	// Number of consecutive remote track reports that must meet the decorrelation
	// criteria before the decorrelation is executed. Required if sysDefaultCode field
	// is "MAN". Allowable entries are integers from 1 to 5.
	ConsecDecorr param.Opt[int64] `json:"consecDecorr,omitzero"`
	// Maximum difference between the reported course of the remote track and the
	// calculated course of the local track. Required if sysDefaultCode field is "MAN".
	// Allowable entries are 15 to 90 in increments of 15 degrees.
	CourseDiff param.Opt[int64] `json:"courseDiff,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Distance between the common and remote track is to exceed the applicable
	// correlation window for the two tracks in order to be decorrelated. Required if
	// sysDefaultCode field is "MAN". Allowable entries are 1.0 to 2.0 in increments of
	// 0.1.
	DecorrWinMult param.Opt[float64] `json:"decorrWinMult,omitzero"`
	// The code for the point of reference from which the coordinates and networks are
	// computed.
	GeoDatum param.Opt[string] `json:"geoDatum,omitzero"`
	// Call sign which identifies one or more communications facilities, commands,
	// authorities, or activities for Joint Range Extension (JRE) units.
	JreCallSign param.Opt[string] `json:"jreCallSign,omitzero"`
	// Joint Range Extension (JRE) unit details.
	JreDetails param.Opt[string] `json:"jreDetails,omitzero"`
	// Link-16 octal track number assigned as the primary JTIDS unit address.
	JrePriAdd param.Opt[int64] `json:"jrePriAdd,omitzero"`
	// Link-16 octal track number assigned as the secondary JTIDS unit address.
	JreSecAdd param.Opt[int64] `json:"jreSecAdd,omitzero"`
	// Designator of the unit for Joint Range Extension (JRE).
	JreUnitDes param.Opt[string] `json:"jreUnitDes,omitzero"`
	// Number used for maximum geodetic position quality. Required if sysDefaultCode
	// field is "MAN". Allowable entires are integers from 1 to 15.
	MaxGeoPosQual param.Opt[int64] `json:"maxGeoPosQual,omitzero"`
	// Track quality to prevent correlation windows from being unrealistically small.
	// Required if sysDefaultCode field is "MAN". Allowable entries are integers from 8
	// to 15.
	MaxTrackQual param.Opt[int64] `json:"maxTrackQual,omitzero"`
	// Data link management code word.
	MgmtCode param.Opt[string] `json:"mgmtCode,omitzero"`
	// Data link management code word meaning.
	MgmtCodeMeaning param.Opt[string] `json:"mgmtCodeMeaning,omitzero"`
	// Number used for minimum geodetic position quality. Required if sysDefaultCode
	// field is "MAN". Allowable entries are integers from 1 to 5.
	MinGeoPosQual param.Opt[int64] `json:"minGeoPosQual,omitzero"`
	// Track quality to prevent correlation windows from being unrealistically large.
	// Required if sysDefaultCode field is "MAN". Allowable entries are integers from 3
	// to 7.
	MinTrackQual param.Opt[int64] `json:"minTrackQual,omitzero"`
	// The month in which this message originated.
	Month param.Opt[string] `json:"month,omitzero"`
	// Provides an additional caveat further identifying the exercise or modifies the
	// exercise nickname.
	OpExInfo param.Opt[string] `json:"opExInfo,omitzero"`
	// The secondary nickname of the option or the alternative of the operational plan
	// or order.
	OpExInfoAlt param.Opt[string] `json:"opExInfoAlt,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The official identifier of the military establishment responsible for the
	// operation plan and the identification number assigned to this plan.
	PlanOrigNum param.Opt[string] `json:"planOrigNum,omitzero"`
	// The unit identifier or call sign of the point of contact for this data link
	// message.
	PocCallSign param.Opt[string] `json:"pocCallSign,omitzero"`
	// WGS84 latitude of the point of contact for this data link message, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	PocLat param.Opt[float64] `json:"pocLat,omitzero"`
	// The location name of the point of contact for this data link message.
	PocLocName param.Opt[string] `json:"pocLocName,omitzero"`
	// WGS84 longitude of the point of contact for this data link message, in degrees.
	// -180 to 180 degrees (negative values west of Prime Meridian).
	PocLon param.Opt[float64] `json:"pocLon,omitzero"`
	// The name of the point of contact for this data link message.
	PocName param.Opt[string] `json:"pocName,omitzero"`
	// The rank or position of the point of contact for this data link message in a
	// military or civilian organization.
	PocRank param.Opt[string] `json:"pocRank,omitzero"`
	// The qualifier which caveats the message status such as AMP (Amplification), CHG
	// (Change), etc.
	Qualifier param.Opt[string] `json:"qualifier,omitzero"`
	// The serial number associated with the message qualifier.
	QualSn param.Opt[int64] `json:"qualSN,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// Track quality to enter if too many duals involving low track quality tracks are
	// occurring. Required if sysDefaultCode field is "MAN". Allowable entries are
	// integers from 2 to 6.
	ResTrackQual param.Opt[int64] `json:"resTrackQual,omitzero"`
	// The unique message identifier assigned by the originator.
	SerialNum param.Opt[string] `json:"serialNum,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Maximum percentage the faster track speed may differ from the slower track
	// speed. Required if sysDefaultCode field is "MAN". Allowable entries are 10 to
	// 100 in increments of 10.
	SpeedDiff param.Opt[int64] `json:"speedDiff,omitzero"`
	// The end of the effective time period of this data link message, in ISO 8601 UTC
	// format with millisecond precision. This may be a relative stop time if used with
	// stopTimeMod.
	StopTime param.Opt[time.Time] `json:"stopTime,omitzero" format:"date-time"`
	// A qualifier for the end of the effective time period of this data link message,
	// such as AFTER, ASOF, NLT, etc. Used with field stopTime to indicate a relative
	// time.
	StopTimeMod param.Opt[string] `json:"stopTimeMod,omitzero"`
	// Indicates the data terminal settings the system defaults to, either automatic
	// correlation/decorrelation (AUTO) or manual (MAN).
	SysDefaultCode param.Opt[string] `json:"sysDefaultCode,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Number added to the basic window calculated from track qualities to ensure that
	// windows still allow valid correlations. Required if sysDefaultCode field is
	// "MAN". Allowable entries are 0.0 to 2.0 in increments of 0.25.
	WinSizeMin param.Opt[float64] `json:"winSizeMin,omitzero"`
	// The correlation window size multiplier to stretch or reduce the window size.
	// Required if sysDefaultCode field is "MAN". Allowable entries are 0.5 to 3.0 in
	// increments of 0.1.
	WinSizeMult param.Opt[float64] `json:"winSizeMult,omitzero"`
	// Array of instructions for acknowledging and the force or units required to
	// acknowledge the data link message being sent.
	AckInstUnits []string `json:"ackInstUnits,omitzero"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of this data link
	// message or document being cancelled.
	CanxSiCs []string `json:"canxSICs,omitzero"`
	// Array of codes that indicate the reasons material is classified.
	ClassReasons []string `json:"classReasons,omitzero"`
	// Array of codes that provide justification for exemption from automatic
	// downgrading or declassification.
	DecExemptCodes []string `json:"decExemptCodes,omitzero"`
	// Array of markings that provide the literal guidance or dates for the downgrading
	// or declassification of this data link message.
	DecInstDates []string `json:"decInstDates,omitzero"`
	// Collection of contact and identification information for designated multilink
	// coordinator duty assignments. There can be 0 to many DataLinkMultiDuty
	// collections within the datalink service.
	MultiDuty []DatalinkIngestMultiDutyParam `json:"multiDuty,omitzero"`
	// Array of non-link specific data unit designators.
	NonLinkUnitDes []string `json:"nonLinkUnitDes,omitzero"`
	// Collection of information describing the establishment and detailed operation of
	// tactical data links. There can be 0 to many DataLinkOps collections within the
	// datalink service.
	Ops []DatalinkIngestOpParam `json:"ops,omitzero"`
	// Array of telephone numbers, radio frequency values, or email addresses of the
	// point of contact for this data link message.
	PocNums []string `json:"pocNums,omitzero"`
	// Collection of reference information. There can be 0 to many DataLinkReferences
	// collections within the datalink service.
	References []DatalinkIngestReferenceParam `json:"references,omitzero"`
	// Collection that identifies points of reference used in the establishment of the
	// data links. There can be 1 to many DataLinkRefPoints collections within the
	// datalink service.
	RefPoints []DatalinkIngestRefPointParam `json:"refPoints,omitzero"`
	// Collection of remarks associated with this data link message.
	Remarks []DatalinkIngestRemarkParam `json:"remarks,omitzero"`
	// Collection of special track numbers used on the data links. There can be 0 to
	// many DataLinkSpecTracks collections within the datalink service.
	SpecTracks []DatalinkIngestSpecTrackParam `json:"specTracks,omitzero"`
	// Array of Link-16 octal track numbers used as the lower limit of a track block.
	TrackNumBlockLLs []int64 `json:"trackNumBlockLLs,omitzero"`
	// Array of defined ranges of Link-11/11B track numbers assigned to a participating
	// unit or reporting unit.
	TrackNumBlocks []string `json:"trackNumBlocks,omitzero"`
	// Collection of information regarding the function, frequency, and priority of
	// interface control and coordination nets for this data link message. There can be
	// 1 to many DataLinkVoiceCoord collections within the datalink service.
	VoiceCoord []DatalinkIngestVoiceCoordParam `json:"voiceCoord,omitzero"`
	paramObj
}

func (r DatalinkIngestParam) MarshalJSON() (data []byte, err error) {
	type shadow DatalinkIngestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatalinkIngestParam) UnmarshalJSON(data []byte) error {
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
type DatalinkIngestDataMode string

const (
	DatalinkIngestDataModeReal      DatalinkIngestDataMode = "REAL"
	DatalinkIngestDataModeTest      DatalinkIngestDataMode = "TEST"
	DatalinkIngestDataModeSimulated DatalinkIngestDataMode = "SIMULATED"
	DatalinkIngestDataModeExercise  DatalinkIngestDataMode = "EXERCISE"
)

// Collection of contact and identification information for designated multilink
// coordinator duty assignments. There can be 0 to many DataLinkMultiDuty
// collections within the datalink service.
type DatalinkIngestMultiDutyParam struct {
	// Specific duties assigned for multilink coordination (e.g. ICO, RICO, SICO).
	Duty param.Opt[string] `json:"duty,omitzero"`
	// The name of the person to be contacted for multilink coordination.
	Name param.Opt[string] `json:"name,omitzero"`
	// The rank or position of the person to be contacted for multilink coordination.
	Rank param.Opt[string] `json:"rank,omitzero"`
	// Designated force of unit specified by ship name, unit call sign, or unit
	// designator.
	UnitDes param.Opt[string] `json:"unitDes,omitzero"`
	// Array of telephone numbers or the frequency values for radio transmission of the
	// person to be contacted for multilink coordination.
	DutyTeleFreqNums []string `json:"dutyTeleFreqNums,omitzero"`
	// Collection of information regarding the function, frequency, and priority of
	// interface control and coordination nets for multilink coordination. There can be
	// 0 to many DataLinkMultiVoiceCoord collections within a DataLinkMultiDuty
	// collection.
	MultiDutyVoiceCoord []DatalinkIngestMultiDutyMultiDutyVoiceCoordParam `json:"multiDutyVoiceCoord,omitzero"`
	paramObj
}

func (r DatalinkIngestMultiDutyParam) MarshalJSON() (data []byte, err error) {
	type shadow DatalinkIngestMultiDutyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatalinkIngestMultiDutyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of information regarding the function, frequency, and priority of
// interface control and coordination nets for multilink coordination. There can be
// 0 to many DataLinkMultiVoiceCoord collections within a DataLinkMultiDuty
// collection.
type DatalinkIngestMultiDutyMultiDutyVoiceCoordParam struct {
	// Priority of a communication circuit, channel or frequency for multilink
	// coordination (e.g. P - Primary, M - Monitor).
	MultiCommPri param.Opt[string] `json:"multiCommPri,omitzero"`
	// Designator used in nonsecure communications to refer to a radio frequency for
	// multilink coordination.
	MultiFreqDes param.Opt[string] `json:"multiFreqDes,omitzero"`
	// Designator assigned to a voice interface control and coordination net for
	// multilink coordination (e.g. ADCCN, DCN, VPN, etc.).
	MultiVoiceNetDes param.Opt[string] `json:"multiVoiceNetDes,omitzero"`
	// Array of telephone numbers or contact frequencies used for interface control for
	// multilink coordination.
	MultiTeleFreqNums []string `json:"multiTeleFreqNums,omitzero"`
	paramObj
}

func (r DatalinkIngestMultiDutyMultiDutyVoiceCoordParam) MarshalJSON() (data []byte, err error) {
	type shadow DatalinkIngestMultiDutyMultiDutyVoiceCoordParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatalinkIngestMultiDutyMultiDutyVoiceCoordParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of information describing the establishment and detailed operation of
// tactical data links. There can be 0 to many DataLinkOps collections within the
// datalink service.
type DatalinkIngestOpParam struct {
	// Detailed characteristics of the data link.
	LinkDetails param.Opt[string] `json:"linkDetails,omitzero"`
	// Name of the data link.
	LinkName param.Opt[string] `json:"linkName,omitzero"`
	// The start of the effective time period of the data link, in ISO 8601 UTC format
	// with millisecond precision.
	LinkStartTime param.Opt[time.Time] `json:"linkStartTime,omitzero" format:"date-time"`
	// The end of the effective time period of the data link, in ISO 8601 UTC format
	// with millisecond precision.
	LinkStopTime param.Opt[time.Time] `json:"linkStopTime,omitzero" format:"date-time"`
	// A qualifier for the end of the effective time period of this data link, such as
	// AFTER, ASOF, NLT, etc. Used with field linkStopTimeMod to indicate a relative
	// time.
	LinkStopTimeMod param.Opt[string] `json:"linkStopTimeMod,omitzero"`
	paramObj
}

func (r DatalinkIngestOpParam) MarshalJSON() (data []byte, err error) {
	type shadow DatalinkIngestOpParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatalinkIngestOpParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of reference information. There can be 0 to many DataLinkReferences
// collections within the datalink service.
type DatalinkIngestReferenceParam struct {
	// The originator of this reference.
	RefOriginator param.Opt[string] `json:"refOriginator,omitzero"`
	// Specifies an alphabetic serial identifier a reference pertaining to the data
	// link message.
	RefSerialID param.Opt[string] `json:"refSerialId,omitzero"`
	// Serial number assigned to this reference.
	RefSerialNum param.Opt[string] `json:"refSerialNum,omitzero"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this reference.
	RefSpecialNotation param.Opt[string] `json:"refSpecialNotation,omitzero"`
	// Timestamp of the referenced message, in ISO 8601 UTC format with millisecond
	// precision.
	RefTs param.Opt[time.Time] `json:"refTs,omitzero" format:"date-time"`
	// Specifies the type of document referenced.
	RefType param.Opt[string] `json:"refType,omitzero"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of the document
	// being referenced.
	RefSiCs []string `json:"refSICs,omitzero"`
	paramObj
}

func (r DatalinkIngestReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow DatalinkIngestReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatalinkIngestReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that identifies points of reference used in the establishment of the
// data links. There can be 1 to many DataLinkRefPoints collections within the
// datalink service.
type DatalinkIngestRefPointParam struct {
	// Indicates when a particular event or nickname becomes effective or the old event
	// or nickname is deleted, in ISO 8601 UTC format with millisecond precision.
	EffEventTime param.Opt[time.Time] `json:"effEventTime,omitzero" format:"date-time"`
	// Identifier to designate a reference point.
	RefDes param.Opt[string] `json:"refDes,omitzero"`
	// WGS84 latitude of the reference point for this data link message, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	RefLat param.Opt[float64] `json:"refLat,omitzero"`
	// The location name of the point of reference for this data link message.
	RefLocName param.Opt[string] `json:"refLocName,omitzero"`
	// WGS84 longitude of the reference point for this data link message, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	RefLon param.Opt[float64] `json:"refLon,omitzero"`
	// Type of data link reference point or grid origin.
	RefPointType param.Opt[string] `json:"refPointType,omitzero"`
	paramObj
}

func (r DatalinkIngestRefPointParam) MarshalJSON() (data []byte, err error) {
	type shadow DatalinkIngestRefPointParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatalinkIngestRefPointParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this data link message.
type DatalinkIngestRemarkParam struct {
	// Text of the remark.
	Text param.Opt[string] `json:"text,omitzero"`
	// Indicates the subject matter of the remark.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r DatalinkIngestRemarkParam) MarshalJSON() (data []byte, err error) {
	type shadow DatalinkIngestRemarkParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatalinkIngestRemarkParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of special track numbers used on the data links. There can be 0 to
// many DataLinkSpecTracks collections within the datalink service.
type DatalinkIngestSpecTrackParam struct {
	// The special track number used on the data link entered as an octal reference
	// number. Used to identify a particular type of platform (e.g. MPA, KRESTA) or
	// platform name (e.g. TROMP, MOUNT WHITNEY) which is not included in assigned
	// track blocks.
	SpecTrackNum param.Opt[string] `json:"specTrackNum,omitzero"`
	// Description of the special track number.
	SpecTrackNumDesc param.Opt[string] `json:"specTrackNumDesc,omitzero"`
	paramObj
}

func (r DatalinkIngestSpecTrackParam) MarshalJSON() (data []byte, err error) {
	type shadow DatalinkIngestSpecTrackParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatalinkIngestSpecTrackParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of information regarding the function, frequency, and priority of
// interface control and coordination nets for this data link message. There can be
// 1 to many DataLinkVoiceCoord collections within the datalink service.
type DatalinkIngestVoiceCoordParam struct {
	// Priority of a communication circuit, channel or frequency for this data link
	// message such as P (Primary), M (Monitor), etc.
	CommPri param.Opt[string] `json:"commPri,omitzero"`
	// Designator used in nonsecure communications to refer to a radio frequency for
	// this data link message.
	FreqDes param.Opt[string] `json:"freqDes,omitzero"`
	// Designator assigned to a voice interface control and coordination net for this
	// data link message (e.g. ADCCN, DCN, VPN, etc.).
	VoiceNetDes param.Opt[string] `json:"voiceNetDes,omitzero"`
	// Array of telephone numbers or contact frequencies used for interface control for
	// this data link message.
	TeleFreqNums []string `json:"teleFreqNums,omitzero"`
	paramObj
}

func (r DatalinkIngestVoiceCoordParam) MarshalJSON() (data []byte, err error) {
	type shadow DatalinkIngestVoiceCoordParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatalinkIngestVoiceCoordParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Beta Version DataLink: Detailed instructions regarding the operations of data
// links.
type LinkStatusDatalinkListResponse struct {
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
	DataMode LinkStatusDatalinkListResponseDataMode `json:"dataMode,required"`
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExName string `json:"opExName,required"`
	// The identifier of the originator of this message.
	Originator string `json:"originator,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the effective time period of this data link message, in ISO 8601
	// UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Array of instructions for acknowledging and the force or units required to
	// acknowledge the data link message being sent.
	AckInstUnits []string `json:"ackInstUnits"`
	// Flag Indicating if formal acknowledgement is required for the particular data
	// link message being sent.
	AckReq bool `json:"ackReq"`
	// Maximum altitude difference between two air tracks, in thousands of feet.
	// Required if sysDefaultCode field is "MAN". Allowable entires are 5 to 50 in
	// increments of 5000 feet.
	AltDiff int64 `json:"altDiff"`
	// The identifier for this data link message cancellation.
	CanxID string `json:"canxId"`
	// The originator of this data link message cancellation.
	CanxOriginator string `json:"canxOriginator"`
	// Serial number assigned to this data link message cancellation.
	CanxSerialNum string `json:"canxSerialNum"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of this data link
	// message or document being cancelled.
	CanxSiCs []string `json:"canxSICs"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this data link message cancellation.
	CanxSpecialNotation string `json:"canxSpecialNotation"`
	// Timestamp of the data link message cancellation, in ISO 8601 UTC format with
	// millisecond precision.
	CanxTs time.Time `json:"canxTs" format:"date-time"`
	// Array of codes that indicate the reasons material is classified.
	ClassReasons []string `json:"classReasons"`
	// Markings that define the source material or the original classification
	// authority for this data link message.
	ClassSource string `json:"classSource"`
	// Number of consecutive remote track reports that must meet the decorrelation
	// criteria before the decorrelation is executed. Required if sysDefaultCode field
	// is "MAN". Allowable entries are integers from 1 to 5.
	ConsecDecorr int64 `json:"consecDecorr"`
	// Maximum difference between the reported course of the remote track and the
	// calculated course of the local track. Required if sysDefaultCode field is "MAN".
	// Allowable entries are 15 to 90 in increments of 15 degrees.
	CourseDiff int64 `json:"courseDiff"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array of codes that provide justification for exemption from automatic
	// downgrading or declassification.
	DecExemptCodes []string `json:"decExemptCodes"`
	// Array of markings that provide the literal guidance or dates for the downgrading
	// or declassification of this data link message.
	DecInstDates []string `json:"decInstDates"`
	// Distance between the common and remote track is to exceed the applicable
	// correlation window for the two tracks in order to be decorrelated. Required if
	// sysDefaultCode field is "MAN". Allowable entries are 1.0 to 2.0 in increments of
	// 0.1.
	DecorrWinMult float64 `json:"decorrWinMult"`
	// The code for the point of reference from which the coordinates and networks are
	// computed.
	GeoDatum string `json:"geoDatum"`
	// Call sign which identifies one or more communications facilities, commands,
	// authorities, or activities for Joint Range Extension (JRE) units.
	JreCallSign string `json:"jreCallSign"`
	// Joint Range Extension (JRE) unit details.
	JreDetails string `json:"jreDetails"`
	// Link-16 octal track number assigned as the primary JTIDS unit address.
	JrePriAdd int64 `json:"jrePriAdd"`
	// Link-16 octal track number assigned as the secondary JTIDS unit address.
	JreSecAdd int64 `json:"jreSecAdd"`
	// Designator of the unit for Joint Range Extension (JRE).
	JreUnitDes string `json:"jreUnitDes"`
	// Number used for maximum geodetic position quality. Required if sysDefaultCode
	// field is "MAN". Allowable entires are integers from 1 to 15.
	MaxGeoPosQual int64 `json:"maxGeoPosQual"`
	// Track quality to prevent correlation windows from being unrealistically small.
	// Required if sysDefaultCode field is "MAN". Allowable entries are integers from 8
	// to 15.
	MaxTrackQual int64 `json:"maxTrackQual"`
	// Data link management code word.
	MgmtCode string `json:"mgmtCode"`
	// Data link management code word meaning.
	MgmtCodeMeaning string `json:"mgmtCodeMeaning"`
	// Number used for minimum geodetic position quality. Required if sysDefaultCode
	// field is "MAN". Allowable entries are integers from 1 to 5.
	MinGeoPosQual int64 `json:"minGeoPosQual"`
	// Track quality to prevent correlation windows from being unrealistically large.
	// Required if sysDefaultCode field is "MAN". Allowable entries are integers from 3
	// to 7.
	MinTrackQual int64 `json:"minTrackQual"`
	// The month in which this message originated.
	Month string `json:"month"`
	// Collection of contact and identification information for designated multilink
	// coordinator duty assignments. There can be 0 to many DataLinkMultiDuty
	// collections within the datalink service.
	MultiDuty []LinkStatusDatalinkListResponseMultiDuty `json:"multiDuty"`
	// Array of non-link specific data unit designators.
	NonLinkUnitDes []string `json:"nonLinkUnitDes"`
	// Provides an additional caveat further identifying the exercise or modifies the
	// exercise nickname.
	OpExInfo string `json:"opExInfo"`
	// The secondary nickname of the option or the alternative of the operational plan
	// or order.
	OpExInfoAlt string `json:"opExInfoAlt"`
	// Collection of information describing the establishment and detailed operation of
	// tactical data links. There can be 0 to many DataLinkOps collections within the
	// datalink service.
	Ops []LinkStatusDatalinkListResponseOp `json:"ops"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The official identifier of the military establishment responsible for the
	// operation plan and the identification number assigned to this plan.
	PlanOrigNum string `json:"planOrigNum"`
	// The unit identifier or call sign of the point of contact for this data link
	// message.
	PocCallSign string `json:"pocCallSign"`
	// WGS84 latitude of the point of contact for this data link message, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	PocLat float64 `json:"pocLat"`
	// The location name of the point of contact for this data link message.
	PocLocName string `json:"pocLocName"`
	// WGS84 longitude of the point of contact for this data link message, in degrees.
	// -180 to 180 degrees (negative values west of Prime Meridian).
	PocLon float64 `json:"pocLon"`
	// The name of the point of contact for this data link message.
	PocName string `json:"pocName"`
	// Array of telephone numbers, radio frequency values, or email addresses of the
	// point of contact for this data link message.
	PocNums []string `json:"pocNums"`
	// The rank or position of the point of contact for this data link message in a
	// military or civilian organization.
	PocRank string `json:"pocRank"`
	// The qualifier which caveats the message status such as AMP (Amplification), CHG
	// (Change), etc.
	Qualifier string `json:"qualifier"`
	// The serial number associated with the message qualifier.
	QualSn int64 `json:"qualSN"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Collection of reference information. There can be 0 to many DataLinkReferences
	// collections within the datalink service.
	References []LinkStatusDatalinkListResponseReference `json:"references"`
	// Collection that identifies points of reference used in the establishment of the
	// data links. There can be 1 to many DataLinkRefPoints collections within the
	// datalink service.
	RefPoints []LinkStatusDatalinkListResponseRefPoint `json:"refPoints"`
	// Collection of remarks associated with this data link message.
	Remarks []LinkStatusDatalinkListResponseRemark `json:"remarks"`
	// Track quality to enter if too many duals involving low track quality tracks are
	// occurring. Required if sysDefaultCode field is "MAN". Allowable entries are
	// integers from 2 to 6.
	ResTrackQual int64 `json:"resTrackQual"`
	// The unique message identifier assigned by the originator.
	SerialNum string `json:"serialNum"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Collection of special track numbers used on the data links. There can be 0 to
	// many DataLinkSpecTracks collections within the datalink service.
	SpecTracks []LinkStatusDatalinkListResponseSpecTrack `json:"specTracks"`
	// Maximum percentage the faster track speed may differ from the slower track
	// speed. Required if sysDefaultCode field is "MAN". Allowable entries are 10 to
	// 100 in increments of 10.
	SpeedDiff int64 `json:"speedDiff"`
	// The end of the effective time period of this data link message, in ISO 8601 UTC
	// format with millisecond precision. This may be a relative stop time if used with
	// stopTimeMod.
	StopTime time.Time `json:"stopTime" format:"date-time"`
	// A qualifier for the end of the effective time period of this data link message,
	// such as AFTER, ASOF, NLT, etc. Used with field stopTime to indicate a relative
	// time.
	StopTimeMod string `json:"stopTimeMod"`
	// Indicates the data terminal settings the system defaults to, either automatic
	// correlation/decorrelation (AUTO) or manual (MAN).
	SysDefaultCode string `json:"sysDefaultCode"`
	// Array of Link-16 octal track numbers used as the lower limit of a track block.
	TrackNumBlockLLs []int64 `json:"trackNumBlockLLs"`
	// Array of defined ranges of Link-11/11B track numbers assigned to a participating
	// unit or reporting unit.
	TrackNumBlocks []string `json:"trackNumBlocks"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Collection of information regarding the function, frequency, and priority of
	// interface control and coordination nets for this data link message. There can be
	// 1 to many DataLinkVoiceCoord collections within the datalink service.
	VoiceCoord []LinkStatusDatalinkListResponseVoiceCoord `json:"voiceCoord"`
	// Number added to the basic window calculated from track qualities to ensure that
	// windows still allow valid correlations. Required if sysDefaultCode field is
	// "MAN". Allowable entries are 0.0 to 2.0 in increments of 0.25.
	WinSizeMin float64 `json:"winSizeMin"`
	// The correlation window size multiplier to stretch or reduce the window size.
	// Required if sysDefaultCode field is "MAN". Allowable entries are 0.5 to 3.0 in
	// increments of 0.1.
	WinSizeMult float64 `json:"winSizeMult"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		OpExName              respjson.Field
		Originator            respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		AckInstUnits          respjson.Field
		AckReq                respjson.Field
		AltDiff               respjson.Field
		CanxID                respjson.Field
		CanxOriginator        respjson.Field
		CanxSerialNum         respjson.Field
		CanxSiCs              respjson.Field
		CanxSpecialNotation   respjson.Field
		CanxTs                respjson.Field
		ClassReasons          respjson.Field
		ClassSource           respjson.Field
		ConsecDecorr          respjson.Field
		CourseDiff            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecExemptCodes        respjson.Field
		DecInstDates          respjson.Field
		DecorrWinMult         respjson.Field
		GeoDatum              respjson.Field
		JreCallSign           respjson.Field
		JreDetails            respjson.Field
		JrePriAdd             respjson.Field
		JreSecAdd             respjson.Field
		JreUnitDes            respjson.Field
		MaxGeoPosQual         respjson.Field
		MaxTrackQual          respjson.Field
		MgmtCode              respjson.Field
		MgmtCodeMeaning       respjson.Field
		MinGeoPosQual         respjson.Field
		MinTrackQual          respjson.Field
		Month                 respjson.Field
		MultiDuty             respjson.Field
		NonLinkUnitDes        respjson.Field
		OpExInfo              respjson.Field
		OpExInfoAlt           respjson.Field
		Ops                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PlanOrigNum           respjson.Field
		PocCallSign           respjson.Field
		PocLat                respjson.Field
		PocLocName            respjson.Field
		PocLon                respjson.Field
		PocName               respjson.Field
		PocNums               respjson.Field
		PocRank               respjson.Field
		Qualifier             respjson.Field
		QualSn                respjson.Field
		RawFileUri            respjson.Field
		References            respjson.Field
		RefPoints             respjson.Field
		Remarks               respjson.Field
		ResTrackQual          respjson.Field
		SerialNum             respjson.Field
		SourceDl              respjson.Field
		SpecTracks            respjson.Field
		SpeedDiff             respjson.Field
		StopTime              respjson.Field
		StopTimeMod           respjson.Field
		SysDefaultCode        respjson.Field
		TrackNumBlockLLs      respjson.Field
		TrackNumBlocks        respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VoiceCoord            respjson.Field
		WinSizeMin            respjson.Field
		WinSizeMult           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkListResponse) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkListResponse) UnmarshalJSON(data []byte) error {
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
type LinkStatusDatalinkListResponseDataMode string

const (
	LinkStatusDatalinkListResponseDataModeReal      LinkStatusDatalinkListResponseDataMode = "REAL"
	LinkStatusDatalinkListResponseDataModeTest      LinkStatusDatalinkListResponseDataMode = "TEST"
	LinkStatusDatalinkListResponseDataModeSimulated LinkStatusDatalinkListResponseDataMode = "SIMULATED"
	LinkStatusDatalinkListResponseDataModeExercise  LinkStatusDatalinkListResponseDataMode = "EXERCISE"
)

// Collection of contact and identification information for designated multilink
// coordinator duty assignments. There can be 0 to many DataLinkMultiDuty
// collections within the datalink service.
type LinkStatusDatalinkListResponseMultiDuty struct {
	// Specific duties assigned for multilink coordination (e.g. ICO, RICO, SICO).
	Duty string `json:"duty"`
	// Array of telephone numbers or the frequency values for radio transmission of the
	// person to be contacted for multilink coordination.
	DutyTeleFreqNums []string `json:"dutyTeleFreqNums"`
	// Collection of information regarding the function, frequency, and priority of
	// interface control and coordination nets for multilink coordination. There can be
	// 0 to many DataLinkMultiVoiceCoord collections within a DataLinkMultiDuty
	// collection.
	MultiDutyVoiceCoord []LinkStatusDatalinkListResponseMultiDutyMultiDutyVoiceCoord `json:"multiDutyVoiceCoord"`
	// The name of the person to be contacted for multilink coordination.
	Name string `json:"name"`
	// The rank or position of the person to be contacted for multilink coordination.
	Rank string `json:"rank"`
	// Designated force of unit specified by ship name, unit call sign, or unit
	// designator.
	UnitDes string `json:"unitDes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duty                respjson.Field
		DutyTeleFreqNums    respjson.Field
		MultiDutyVoiceCoord respjson.Field
		Name                respjson.Field
		Rank                respjson.Field
		UnitDes             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkListResponseMultiDuty) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkListResponseMultiDuty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of information regarding the function, frequency, and priority of
// interface control and coordination nets for multilink coordination. There can be
// 0 to many DataLinkMultiVoiceCoord collections within a DataLinkMultiDuty
// collection.
type LinkStatusDatalinkListResponseMultiDutyMultiDutyVoiceCoord struct {
	// Priority of a communication circuit, channel or frequency for multilink
	// coordination (e.g. P - Primary, M - Monitor).
	MultiCommPri string `json:"multiCommPri"`
	// Designator used in nonsecure communications to refer to a radio frequency for
	// multilink coordination.
	MultiFreqDes string `json:"multiFreqDes"`
	// Array of telephone numbers or contact frequencies used for interface control for
	// multilink coordination.
	MultiTeleFreqNums []string `json:"multiTeleFreqNums"`
	// Designator assigned to a voice interface control and coordination net for
	// multilink coordination (e.g. ADCCN, DCN, VPN, etc.).
	MultiVoiceNetDes string `json:"multiVoiceNetDes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MultiCommPri      respjson.Field
		MultiFreqDes      respjson.Field
		MultiTeleFreqNums respjson.Field
		MultiVoiceNetDes  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkListResponseMultiDutyMultiDutyVoiceCoord) RawJSON() string {
	return r.JSON.raw
}
func (r *LinkStatusDatalinkListResponseMultiDutyMultiDutyVoiceCoord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of information describing the establishment and detailed operation of
// tactical data links. There can be 0 to many DataLinkOps collections within the
// datalink service.
type LinkStatusDatalinkListResponseOp struct {
	// Detailed characteristics of the data link.
	LinkDetails string `json:"linkDetails"`
	// Name of the data link.
	LinkName string `json:"linkName"`
	// The start of the effective time period of the data link, in ISO 8601 UTC format
	// with millisecond precision.
	LinkStartTime time.Time `json:"linkStartTime" format:"date-time"`
	// The end of the effective time period of the data link, in ISO 8601 UTC format
	// with millisecond precision.
	LinkStopTime time.Time `json:"linkStopTime" format:"date-time"`
	// A qualifier for the end of the effective time period of this data link, such as
	// AFTER, ASOF, NLT, etc. Used with field linkStopTimeMod to indicate a relative
	// time.
	LinkStopTimeMod string `json:"linkStopTimeMod"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkDetails     respjson.Field
		LinkName        respjson.Field
		LinkStartTime   respjson.Field
		LinkStopTime    respjson.Field
		LinkStopTimeMod respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkListResponseOp) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkListResponseOp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of reference information. There can be 0 to many DataLinkReferences
// collections within the datalink service.
type LinkStatusDatalinkListResponseReference struct {
	// The originator of this reference.
	RefOriginator string `json:"refOriginator"`
	// Specifies an alphabetic serial identifier a reference pertaining to the data
	// link message.
	RefSerialID string `json:"refSerialId"`
	// Serial number assigned to this reference.
	RefSerialNum string `json:"refSerialNum"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of the document
	// being referenced.
	RefSiCs []string `json:"refSICs"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this reference.
	RefSpecialNotation string `json:"refSpecialNotation"`
	// Timestamp of the referenced message, in ISO 8601 UTC format with millisecond
	// precision.
	RefTs time.Time `json:"refTs" format:"date-time"`
	// Specifies the type of document referenced.
	RefType string `json:"refType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RefOriginator      respjson.Field
		RefSerialID        respjson.Field
		RefSerialNum       respjson.Field
		RefSiCs            respjson.Field
		RefSpecialNotation respjson.Field
		RefTs              respjson.Field
		RefType            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkListResponseReference) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkListResponseReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that identifies points of reference used in the establishment of the
// data links. There can be 1 to many DataLinkRefPoints collections within the
// datalink service.
type LinkStatusDatalinkListResponseRefPoint struct {
	// Indicates when a particular event or nickname becomes effective or the old event
	// or nickname is deleted, in ISO 8601 UTC format with millisecond precision.
	EffEventTime time.Time `json:"effEventTime" format:"date-time"`
	// Identifier to designate a reference point.
	RefDes string `json:"refDes"`
	// WGS84 latitude of the reference point for this data link message, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	RefLat float64 `json:"refLat"`
	// The location name of the point of reference for this data link message.
	RefLocName string `json:"refLocName"`
	// WGS84 longitude of the reference point for this data link message, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	RefLon float64 `json:"refLon"`
	// Type of data link reference point or grid origin.
	RefPointType string `json:"refPointType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EffEventTime respjson.Field
		RefDes       respjson.Field
		RefLat       respjson.Field
		RefLocName   respjson.Field
		RefLon       respjson.Field
		RefPointType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkListResponseRefPoint) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkListResponseRefPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this data link message.
type LinkStatusDatalinkListResponseRemark struct {
	// Text of the remark.
	Text string `json:"text"`
	// Indicates the subject matter of the remark.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkListResponseRemark) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkListResponseRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of special track numbers used on the data links. There can be 0 to
// many DataLinkSpecTracks collections within the datalink service.
type LinkStatusDatalinkListResponseSpecTrack struct {
	// The special track number used on the data link entered as an octal reference
	// number. Used to identify a particular type of platform (e.g. MPA, KRESTA) or
	// platform name (e.g. TROMP, MOUNT WHITNEY) which is not included in assigned
	// track blocks.
	SpecTrackNum string `json:"specTrackNum"`
	// Description of the special track number.
	SpecTrackNumDesc string `json:"specTrackNumDesc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SpecTrackNum     respjson.Field
		SpecTrackNumDesc respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkListResponseSpecTrack) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkListResponseSpecTrack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of information regarding the function, frequency, and priority of
// interface control and coordination nets for this data link message. There can be
// 1 to many DataLinkVoiceCoord collections within the datalink service.
type LinkStatusDatalinkListResponseVoiceCoord struct {
	// Priority of a communication circuit, channel or frequency for this data link
	// message such as P (Primary), M (Monitor), etc.
	CommPri string `json:"commPri"`
	// Designator used in nonsecure communications to refer to a radio frequency for
	// this data link message.
	FreqDes string `json:"freqDes"`
	// Array of telephone numbers or contact frequencies used for interface control for
	// this data link message.
	TeleFreqNums []string `json:"teleFreqNums"`
	// Designator assigned to a voice interface control and coordination net for this
	// data link message (e.g. ADCCN, DCN, VPN, etc.).
	VoiceNetDes string `json:"voiceNetDes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommPri      respjson.Field
		FreqDes      respjson.Field
		TeleFreqNums respjson.Field
		VoiceNetDes  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkListResponseVoiceCoord) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkListResponseVoiceCoord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkStatusDatalinkQueryhelpResponse struct {
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
func (r LinkStatusDatalinkQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Beta Version DataLink: Detailed instructions regarding the operations of data
// links.
type LinkStatusDatalinkTupleResponse struct {
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
	DataMode LinkStatusDatalinkTupleResponseDataMode `json:"dataMode,required"`
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExName string `json:"opExName,required"`
	// The identifier of the originator of this message.
	Originator string `json:"originator,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the effective time period of this data link message, in ISO 8601
	// UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Array of instructions for acknowledging and the force or units required to
	// acknowledge the data link message being sent.
	AckInstUnits []string `json:"ackInstUnits"`
	// Flag Indicating if formal acknowledgement is required for the particular data
	// link message being sent.
	AckReq bool `json:"ackReq"`
	// Maximum altitude difference between two air tracks, in thousands of feet.
	// Required if sysDefaultCode field is "MAN". Allowable entires are 5 to 50 in
	// increments of 5000 feet.
	AltDiff int64 `json:"altDiff"`
	// The identifier for this data link message cancellation.
	CanxID string `json:"canxId"`
	// The originator of this data link message cancellation.
	CanxOriginator string `json:"canxOriginator"`
	// Serial number assigned to this data link message cancellation.
	CanxSerialNum string `json:"canxSerialNum"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of this data link
	// message or document being cancelled.
	CanxSiCs []string `json:"canxSICs"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this data link message cancellation.
	CanxSpecialNotation string `json:"canxSpecialNotation"`
	// Timestamp of the data link message cancellation, in ISO 8601 UTC format with
	// millisecond precision.
	CanxTs time.Time `json:"canxTs" format:"date-time"`
	// Array of codes that indicate the reasons material is classified.
	ClassReasons []string `json:"classReasons"`
	// Markings that define the source material or the original classification
	// authority for this data link message.
	ClassSource string `json:"classSource"`
	// Number of consecutive remote track reports that must meet the decorrelation
	// criteria before the decorrelation is executed. Required if sysDefaultCode field
	// is "MAN". Allowable entries are integers from 1 to 5.
	ConsecDecorr int64 `json:"consecDecorr"`
	// Maximum difference between the reported course of the remote track and the
	// calculated course of the local track. Required if sysDefaultCode field is "MAN".
	// Allowable entries are 15 to 90 in increments of 15 degrees.
	CourseDiff int64 `json:"courseDiff"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Array of codes that provide justification for exemption from automatic
	// downgrading or declassification.
	DecExemptCodes []string `json:"decExemptCodes"`
	// Array of markings that provide the literal guidance or dates for the downgrading
	// or declassification of this data link message.
	DecInstDates []string `json:"decInstDates"`
	// Distance between the common and remote track is to exceed the applicable
	// correlation window for the two tracks in order to be decorrelated. Required if
	// sysDefaultCode field is "MAN". Allowable entries are 1.0 to 2.0 in increments of
	// 0.1.
	DecorrWinMult float64 `json:"decorrWinMult"`
	// The code for the point of reference from which the coordinates and networks are
	// computed.
	GeoDatum string `json:"geoDatum"`
	// Call sign which identifies one or more communications facilities, commands,
	// authorities, or activities for Joint Range Extension (JRE) units.
	JreCallSign string `json:"jreCallSign"`
	// Joint Range Extension (JRE) unit details.
	JreDetails string `json:"jreDetails"`
	// Link-16 octal track number assigned as the primary JTIDS unit address.
	JrePriAdd int64 `json:"jrePriAdd"`
	// Link-16 octal track number assigned as the secondary JTIDS unit address.
	JreSecAdd int64 `json:"jreSecAdd"`
	// Designator of the unit for Joint Range Extension (JRE).
	JreUnitDes string `json:"jreUnitDes"`
	// Number used for maximum geodetic position quality. Required if sysDefaultCode
	// field is "MAN". Allowable entires are integers from 1 to 15.
	MaxGeoPosQual int64 `json:"maxGeoPosQual"`
	// Track quality to prevent correlation windows from being unrealistically small.
	// Required if sysDefaultCode field is "MAN". Allowable entries are integers from 8
	// to 15.
	MaxTrackQual int64 `json:"maxTrackQual"`
	// Data link management code word.
	MgmtCode string `json:"mgmtCode"`
	// Data link management code word meaning.
	MgmtCodeMeaning string `json:"mgmtCodeMeaning"`
	// Number used for minimum geodetic position quality. Required if sysDefaultCode
	// field is "MAN". Allowable entries are integers from 1 to 5.
	MinGeoPosQual int64 `json:"minGeoPosQual"`
	// Track quality to prevent correlation windows from being unrealistically large.
	// Required if sysDefaultCode field is "MAN". Allowable entries are integers from 3
	// to 7.
	MinTrackQual int64 `json:"minTrackQual"`
	// The month in which this message originated.
	Month string `json:"month"`
	// Collection of contact and identification information for designated multilink
	// coordinator duty assignments. There can be 0 to many DataLinkMultiDuty
	// collections within the datalink service.
	MultiDuty []LinkStatusDatalinkTupleResponseMultiDuty `json:"multiDuty"`
	// Array of non-link specific data unit designators.
	NonLinkUnitDes []string `json:"nonLinkUnitDes"`
	// Provides an additional caveat further identifying the exercise or modifies the
	// exercise nickname.
	OpExInfo string `json:"opExInfo"`
	// The secondary nickname of the option or the alternative of the operational plan
	// or order.
	OpExInfoAlt string `json:"opExInfoAlt"`
	// Collection of information describing the establishment and detailed operation of
	// tactical data links. There can be 0 to many DataLinkOps collections within the
	// datalink service.
	Ops []LinkStatusDatalinkTupleResponseOp `json:"ops"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The official identifier of the military establishment responsible for the
	// operation plan and the identification number assigned to this plan.
	PlanOrigNum string `json:"planOrigNum"`
	// The unit identifier or call sign of the point of contact for this data link
	// message.
	PocCallSign string `json:"pocCallSign"`
	// WGS84 latitude of the point of contact for this data link message, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	PocLat float64 `json:"pocLat"`
	// The location name of the point of contact for this data link message.
	PocLocName string `json:"pocLocName"`
	// WGS84 longitude of the point of contact for this data link message, in degrees.
	// -180 to 180 degrees (negative values west of Prime Meridian).
	PocLon float64 `json:"pocLon"`
	// The name of the point of contact for this data link message.
	PocName string `json:"pocName"`
	// Array of telephone numbers, radio frequency values, or email addresses of the
	// point of contact for this data link message.
	PocNums []string `json:"pocNums"`
	// The rank or position of the point of contact for this data link message in a
	// military or civilian organization.
	PocRank string `json:"pocRank"`
	// The qualifier which caveats the message status such as AMP (Amplification), CHG
	// (Change), etc.
	Qualifier string `json:"qualifier"`
	// The serial number associated with the message qualifier.
	QualSn int64 `json:"qualSN"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// Collection of reference information. There can be 0 to many DataLinkReferences
	// collections within the datalink service.
	References []LinkStatusDatalinkTupleResponseReference `json:"references"`
	// Collection that identifies points of reference used in the establishment of the
	// data links. There can be 1 to many DataLinkRefPoints collections within the
	// datalink service.
	RefPoints []LinkStatusDatalinkTupleResponseRefPoint `json:"refPoints"`
	// Collection of remarks associated with this data link message.
	Remarks []LinkStatusDatalinkTupleResponseRemark `json:"remarks"`
	// Track quality to enter if too many duals involving low track quality tracks are
	// occurring. Required if sysDefaultCode field is "MAN". Allowable entries are
	// integers from 2 to 6.
	ResTrackQual int64 `json:"resTrackQual"`
	// The unique message identifier assigned by the originator.
	SerialNum string `json:"serialNum"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Collection of special track numbers used on the data links. There can be 0 to
	// many DataLinkSpecTracks collections within the datalink service.
	SpecTracks []LinkStatusDatalinkTupleResponseSpecTrack `json:"specTracks"`
	// Maximum percentage the faster track speed may differ from the slower track
	// speed. Required if sysDefaultCode field is "MAN". Allowable entries are 10 to
	// 100 in increments of 10.
	SpeedDiff int64 `json:"speedDiff"`
	// The end of the effective time period of this data link message, in ISO 8601 UTC
	// format with millisecond precision. This may be a relative stop time if used with
	// stopTimeMod.
	StopTime time.Time `json:"stopTime" format:"date-time"`
	// A qualifier for the end of the effective time period of this data link message,
	// such as AFTER, ASOF, NLT, etc. Used with field stopTime to indicate a relative
	// time.
	StopTimeMod string `json:"stopTimeMod"`
	// Indicates the data terminal settings the system defaults to, either automatic
	// correlation/decorrelation (AUTO) or manual (MAN).
	SysDefaultCode string `json:"sysDefaultCode"`
	// Array of Link-16 octal track numbers used as the lower limit of a track block.
	TrackNumBlockLLs []int64 `json:"trackNumBlockLLs"`
	// Array of defined ranges of Link-11/11B track numbers assigned to a participating
	// unit or reporting unit.
	TrackNumBlocks []string `json:"trackNumBlocks"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Collection of information regarding the function, frequency, and priority of
	// interface control and coordination nets for this data link message. There can be
	// 1 to many DataLinkVoiceCoord collections within the datalink service.
	VoiceCoord []LinkStatusDatalinkTupleResponseVoiceCoord `json:"voiceCoord"`
	// Number added to the basic window calculated from track qualities to ensure that
	// windows still allow valid correlations. Required if sysDefaultCode field is
	// "MAN". Allowable entries are 0.0 to 2.0 in increments of 0.25.
	WinSizeMin float64 `json:"winSizeMin"`
	// The correlation window size multiplier to stretch or reduce the window size.
	// Required if sysDefaultCode field is "MAN". Allowable entries are 0.5 to 3.0 in
	// increments of 0.1.
	WinSizeMult float64 `json:"winSizeMult"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		OpExName              respjson.Field
		Originator            respjson.Field
		Source                respjson.Field
		StartTime             respjson.Field
		ID                    respjson.Field
		AckInstUnits          respjson.Field
		AckReq                respjson.Field
		AltDiff               respjson.Field
		CanxID                respjson.Field
		CanxOriginator        respjson.Field
		CanxSerialNum         respjson.Field
		CanxSiCs              respjson.Field
		CanxSpecialNotation   respjson.Field
		CanxTs                respjson.Field
		ClassReasons          respjson.Field
		ClassSource           respjson.Field
		ConsecDecorr          respjson.Field
		CourseDiff            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecExemptCodes        respjson.Field
		DecInstDates          respjson.Field
		DecorrWinMult         respjson.Field
		GeoDatum              respjson.Field
		JreCallSign           respjson.Field
		JreDetails            respjson.Field
		JrePriAdd             respjson.Field
		JreSecAdd             respjson.Field
		JreUnitDes            respjson.Field
		MaxGeoPosQual         respjson.Field
		MaxTrackQual          respjson.Field
		MgmtCode              respjson.Field
		MgmtCodeMeaning       respjson.Field
		MinGeoPosQual         respjson.Field
		MinTrackQual          respjson.Field
		Month                 respjson.Field
		MultiDuty             respjson.Field
		NonLinkUnitDes        respjson.Field
		OpExInfo              respjson.Field
		OpExInfoAlt           respjson.Field
		Ops                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PlanOrigNum           respjson.Field
		PocCallSign           respjson.Field
		PocLat                respjson.Field
		PocLocName            respjson.Field
		PocLon                respjson.Field
		PocName               respjson.Field
		PocNums               respjson.Field
		PocRank               respjson.Field
		Qualifier             respjson.Field
		QualSn                respjson.Field
		RawFileUri            respjson.Field
		References            respjson.Field
		RefPoints             respjson.Field
		Remarks               respjson.Field
		ResTrackQual          respjson.Field
		SerialNum             respjson.Field
		SourceDl              respjson.Field
		SpecTracks            respjson.Field
		SpeedDiff             respjson.Field
		StopTime              respjson.Field
		StopTimeMod           respjson.Field
		SysDefaultCode        respjson.Field
		TrackNumBlockLLs      respjson.Field
		TrackNumBlocks        respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VoiceCoord            respjson.Field
		WinSizeMin            respjson.Field
		WinSizeMult           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkTupleResponse) UnmarshalJSON(data []byte) error {
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
type LinkStatusDatalinkTupleResponseDataMode string

const (
	LinkStatusDatalinkTupleResponseDataModeReal      LinkStatusDatalinkTupleResponseDataMode = "REAL"
	LinkStatusDatalinkTupleResponseDataModeTest      LinkStatusDatalinkTupleResponseDataMode = "TEST"
	LinkStatusDatalinkTupleResponseDataModeSimulated LinkStatusDatalinkTupleResponseDataMode = "SIMULATED"
	LinkStatusDatalinkTupleResponseDataModeExercise  LinkStatusDatalinkTupleResponseDataMode = "EXERCISE"
)

// Collection of contact and identification information for designated multilink
// coordinator duty assignments. There can be 0 to many DataLinkMultiDuty
// collections within the datalink service.
type LinkStatusDatalinkTupleResponseMultiDuty struct {
	// Specific duties assigned for multilink coordination (e.g. ICO, RICO, SICO).
	Duty string `json:"duty"`
	// Array of telephone numbers or the frequency values for radio transmission of the
	// person to be contacted for multilink coordination.
	DutyTeleFreqNums []string `json:"dutyTeleFreqNums"`
	// Collection of information regarding the function, frequency, and priority of
	// interface control and coordination nets for multilink coordination. There can be
	// 0 to many DataLinkMultiVoiceCoord collections within a DataLinkMultiDuty
	// collection.
	MultiDutyVoiceCoord []LinkStatusDatalinkTupleResponseMultiDutyMultiDutyVoiceCoord `json:"multiDutyVoiceCoord"`
	// The name of the person to be contacted for multilink coordination.
	Name string `json:"name"`
	// The rank or position of the person to be contacted for multilink coordination.
	Rank string `json:"rank"`
	// Designated force of unit specified by ship name, unit call sign, or unit
	// designator.
	UnitDes string `json:"unitDes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duty                respjson.Field
		DutyTeleFreqNums    respjson.Field
		MultiDutyVoiceCoord respjson.Field
		Name                respjson.Field
		Rank                respjson.Field
		UnitDes             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkTupleResponseMultiDuty) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkTupleResponseMultiDuty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of information regarding the function, frequency, and priority of
// interface control and coordination nets for multilink coordination. There can be
// 0 to many DataLinkMultiVoiceCoord collections within a DataLinkMultiDuty
// collection.
type LinkStatusDatalinkTupleResponseMultiDutyMultiDutyVoiceCoord struct {
	// Priority of a communication circuit, channel or frequency for multilink
	// coordination (e.g. P - Primary, M - Monitor).
	MultiCommPri string `json:"multiCommPri"`
	// Designator used in nonsecure communications to refer to a radio frequency for
	// multilink coordination.
	MultiFreqDes string `json:"multiFreqDes"`
	// Array of telephone numbers or contact frequencies used for interface control for
	// multilink coordination.
	MultiTeleFreqNums []string `json:"multiTeleFreqNums"`
	// Designator assigned to a voice interface control and coordination net for
	// multilink coordination (e.g. ADCCN, DCN, VPN, etc.).
	MultiVoiceNetDes string `json:"multiVoiceNetDes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MultiCommPri      respjson.Field
		MultiFreqDes      respjson.Field
		MultiTeleFreqNums respjson.Field
		MultiVoiceNetDes  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkTupleResponseMultiDutyMultiDutyVoiceCoord) RawJSON() string {
	return r.JSON.raw
}
func (r *LinkStatusDatalinkTupleResponseMultiDutyMultiDutyVoiceCoord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of information describing the establishment and detailed operation of
// tactical data links. There can be 0 to many DataLinkOps collections within the
// datalink service.
type LinkStatusDatalinkTupleResponseOp struct {
	// Detailed characteristics of the data link.
	LinkDetails string `json:"linkDetails"`
	// Name of the data link.
	LinkName string `json:"linkName"`
	// The start of the effective time period of the data link, in ISO 8601 UTC format
	// with millisecond precision.
	LinkStartTime time.Time `json:"linkStartTime" format:"date-time"`
	// The end of the effective time period of the data link, in ISO 8601 UTC format
	// with millisecond precision.
	LinkStopTime time.Time `json:"linkStopTime" format:"date-time"`
	// A qualifier for the end of the effective time period of this data link, such as
	// AFTER, ASOF, NLT, etc. Used with field linkStopTimeMod to indicate a relative
	// time.
	LinkStopTimeMod string `json:"linkStopTimeMod"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkDetails     respjson.Field
		LinkName        respjson.Field
		LinkStartTime   respjson.Field
		LinkStopTime    respjson.Field
		LinkStopTimeMod respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkTupleResponseOp) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkTupleResponseOp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of reference information. There can be 0 to many DataLinkReferences
// collections within the datalink service.
type LinkStatusDatalinkTupleResponseReference struct {
	// The originator of this reference.
	RefOriginator string `json:"refOriginator"`
	// Specifies an alphabetic serial identifier a reference pertaining to the data
	// link message.
	RefSerialID string `json:"refSerialId"`
	// Serial number assigned to this reference.
	RefSerialNum string `json:"refSerialNum"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of the document
	// being referenced.
	RefSiCs []string `json:"refSICs"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this reference.
	RefSpecialNotation string `json:"refSpecialNotation"`
	// Timestamp of the referenced message, in ISO 8601 UTC format with millisecond
	// precision.
	RefTs time.Time `json:"refTs" format:"date-time"`
	// Specifies the type of document referenced.
	RefType string `json:"refType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RefOriginator      respjson.Field
		RefSerialID        respjson.Field
		RefSerialNum       respjson.Field
		RefSiCs            respjson.Field
		RefSpecialNotation respjson.Field
		RefTs              respjson.Field
		RefType            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkTupleResponseReference) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkTupleResponseReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that identifies points of reference used in the establishment of the
// data links. There can be 1 to many DataLinkRefPoints collections within the
// datalink service.
type LinkStatusDatalinkTupleResponseRefPoint struct {
	// Indicates when a particular event or nickname becomes effective or the old event
	// or nickname is deleted, in ISO 8601 UTC format with millisecond precision.
	EffEventTime time.Time `json:"effEventTime" format:"date-time"`
	// Identifier to designate a reference point.
	RefDes string `json:"refDes"`
	// WGS84 latitude of the reference point for this data link message, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	RefLat float64 `json:"refLat"`
	// The location name of the point of reference for this data link message.
	RefLocName string `json:"refLocName"`
	// WGS84 longitude of the reference point for this data link message, in degrees.
	// -90 to 90 degrees (negative values south of equator).
	RefLon float64 `json:"refLon"`
	// Type of data link reference point or grid origin.
	RefPointType string `json:"refPointType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EffEventTime respjson.Field
		RefDes       respjson.Field
		RefLat       respjson.Field
		RefLocName   respjson.Field
		RefLon       respjson.Field
		RefPointType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkTupleResponseRefPoint) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkTupleResponseRefPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of remarks associated with this data link message.
type LinkStatusDatalinkTupleResponseRemark struct {
	// Text of the remark.
	Text string `json:"text"`
	// Indicates the subject matter of the remark.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkTupleResponseRemark) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkTupleResponseRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of special track numbers used on the data links. There can be 0 to
// many DataLinkSpecTracks collections within the datalink service.
type LinkStatusDatalinkTupleResponseSpecTrack struct {
	// The special track number used on the data link entered as an octal reference
	// number. Used to identify a particular type of platform (e.g. MPA, KRESTA) or
	// platform name (e.g. TROMP, MOUNT WHITNEY) which is not included in assigned
	// track blocks.
	SpecTrackNum string `json:"specTrackNum"`
	// Description of the special track number.
	SpecTrackNumDesc string `json:"specTrackNumDesc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SpecTrackNum     respjson.Field
		SpecTrackNumDesc respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkTupleResponseSpecTrack) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkTupleResponseSpecTrack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of information regarding the function, frequency, and priority of
// interface control and coordination nets for this data link message. There can be
// 1 to many DataLinkVoiceCoord collections within the datalink service.
type LinkStatusDatalinkTupleResponseVoiceCoord struct {
	// Priority of a communication circuit, channel or frequency for this data link
	// message such as P (Primary), M (Monitor), etc.
	CommPri string `json:"commPri"`
	// Designator used in nonsecure communications to refer to a radio frequency for
	// this data link message.
	FreqDes string `json:"freqDes"`
	// Array of telephone numbers or contact frequencies used for interface control for
	// this data link message.
	TeleFreqNums []string `json:"teleFreqNums"`
	// Designator assigned to a voice interface control and coordination net for this
	// data link message (e.g. ADCCN, DCN, VPN, etc.).
	VoiceNetDes string `json:"voiceNetDes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommPri      respjson.Field
		FreqDes      respjson.Field
		TeleFreqNums respjson.Field
		VoiceNetDes  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkStatusDatalinkTupleResponseVoiceCoord) RawJSON() string { return r.JSON.raw }
func (r *LinkStatusDatalinkTupleResponseVoiceCoord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkStatusDatalinkNewParams struct {
	// Beta Version DataLink: Detailed instructions regarding the operations of data
	// links.
	DatalinkIngest DatalinkIngestParam
	paramObj
}

func (r LinkStatusDatalinkNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DatalinkIngest)
}
func (r *LinkStatusDatalinkNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DatalinkIngest)
}

type LinkStatusDatalinkListParams struct {
	// The start of the effective time period of this data link message, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LinkStatusDatalinkListParams]'s query parameters as
// `url.Values`.
func (r LinkStatusDatalinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LinkStatusDatalinkCountParams struct {
	// The start of the effective time period of this data link message, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LinkStatusDatalinkCountParams]'s query parameters as
// `url.Values`.
func (r LinkStatusDatalinkCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LinkStatusDatalinkTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The start of the effective time period of this data link message, in ISO 8601
	// UTC format with millisecond precision. (YYYY-MM-DDTHH:MM:SS.sssZ)
	StartTime   time.Time        `query:"startTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LinkStatusDatalinkTupleParams]'s query parameters as
// `url.Values`.
func (r LinkStatusDatalinkTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LinkStatusDatalinkUnvalidatedPublishParams struct {
	Body []DatalinkIngestParam
	paramObj
}

func (r LinkStatusDatalinkUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LinkStatusDatalinkUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
