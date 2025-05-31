// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// AirOperationAirTaskingOrderService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirOperationAirTaskingOrderService] method instead.
type AirOperationAirTaskingOrderService struct {
	Options []option.RequestOption
}

// NewAirOperationAirTaskingOrderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAirOperationAirTaskingOrderService(opts ...option.RequestOption) (r AirOperationAirTaskingOrderService) {
	r = AirOperationAirTaskingOrderService{}
	r.Options = opts
	return
}

// Service operation to take a single airtaskingorder record as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AirOperationAirTaskingOrderService) New(ctx context.Context, body AirOperationAirTaskingOrderNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airtaskingorder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single airtaskingorder record by its unique ID passed
// as a path parameter.
func (r *AirOperationAirTaskingOrderService) Get(ctx context.Context, id string, query AirOperationAirTaskingOrderGetParams, opts ...option.RequestOption) (res *AirTaskingOrderFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airtaskingorder/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirOperationAirTaskingOrderService) List(ctx context.Context, query AirOperationAirTaskingOrderListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AirtaskingorderAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/airtaskingorder"
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
func (r *AirOperationAirTaskingOrderService) ListAutoPaging(ctx context.Context, query AirOperationAirTaskingOrderListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AirtaskingorderAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirOperationAirTaskingOrderService) Count(ctx context.Context, query AirOperationAirTaskingOrderCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airtaskingorder/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AirOperationAirTaskingOrderService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *AirOperationAirTaskingOrderQueryHelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airtaskingorder/queryhelp"
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
func (r *AirOperationAirTaskingOrderService) Tuple(ctx context.Context, query AirOperationAirTaskingOrderTupleParams, opts ...option.RequestOption) (res *[]AirTaskingOrderFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airtaskingorder/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple airtaskingorder records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *AirOperationAirTaskingOrderService) UnvalidatedPublish(ctx context.Context, body AirOperationAirTaskingOrderUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-airtaskingorder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Beta Version Air Tasking Order: The ATO is used to task air missions, assign
// cross force tasking as well as intraservice tasking.
type AirTaskingOrderFull struct {
	// The effective begin time for this ATO in ISO 8601 UTC format with millisecond
	// precision.
	BeginTs time.Time `json:"beginTs,required" format:"date-time"`
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
	DataMode AirTaskingOrderFullDataMode `json:"dataMode,required"`
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExerName string `json:"opExerName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The indicator specifying an affirmative or a negatice condition for this
	// message.
	AckReqInd string `json:"ackReqInd"`
	// Specifies textual data amplifying the data contained in the acknowledgement
	// requirement indicator (ackRedInd) field or the unit required to acknowledge.
	AckUnitInstructions string `json:"ackUnitInstructions"`
	// A collection that specifies the tasked country, tasked service, unit and mission
	// level tasking for this ATO.
	AcMsnTasking []AirTaskingOrderFullAcMsnTasking `json:"acMsnTasking"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// The effective end time for this ATO in ISO 8601 UTC format with millisecond
	// precision.
	EndTs time.Time `json:"endTs" format:"date-time"`
	// A collection that details special instructions, important information, guidance,
	// and amplifying information regarding this ATO.
	GenText []AirTaskingOrderFullGenText `json:"genText"`
	// The month in which the message originated.
	MsgMonth string `json:"msgMonth"`
	// The identifier of the originator of the message.
	MsgOriginator string `json:"msgOriginator"`
	// The qualifier which caveats the message status.
	MsgQualifier string `json:"msgQualifier"`
	// The unique message identifier sequentially assigned by the originator.
	MsgSn string `json:"msgSN"`
	// A collection that specifies the naval flight operations for this ATO.
	NavalFltOps []AirTaskingOrderFullNavalFltOp `json:"navalFltOps"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BeginTs               respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		OpExerName            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AckReqInd             respjson.Field
		AckUnitInstructions   respjson.Field
		AcMsnTasking          respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndTs                 respjson.Field
		GenText               respjson.Field
		MsgMonth              respjson.Field
		MsgOriginator         respjson.Field
		MsgQualifier          respjson.Field
		MsgSn                 respjson.Field
		NavalFltOps           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RawFileUri            respjson.Field
		SourceDl              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTaskingOrderFull) RawJSON() string { return r.JSON.raw }
func (r *AirTaskingOrderFull) UnmarshalJSON(data []byte) error {
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
type AirTaskingOrderFullDataMode string

const (
	AirTaskingOrderFullDataModeReal      AirTaskingOrderFullDataMode = "REAL"
	AirTaskingOrderFullDataModeTest      AirTaskingOrderFullDataMode = "TEST"
	AirTaskingOrderFullDataModeSimulated AirTaskingOrderFullDataMode = "SIMULATED"
	AirTaskingOrderFullDataModeExercise  AirTaskingOrderFullDataMode = "EXERCISE"
)

// Collection that specifies the tasked country, tasked service, unit and mission
// level tasking for this ATO.
type AirTaskingOrderFullAcMsnTasking struct {
	// The country code responsible for conducting this aircraft mission tasking for
	// the exercise or operation.
	CountryCode string `json:"countryCode,required"`
	// The service tasked with conducting this aircraft mission tasking for the
	// exercise or operation.
	TaskedService string `json:"taskedService,required"`
	// The designator of the unit that is tasked to perform this aircraft mission
	// tasking.
	UnitDesignator string `json:"unitDesignator,required"`
	// A collection of aircraft mission location information for this aircraft mission
	// tasking.
	AcMsnLocSeg []AirTaskingOrderFullAcMsnTaskingAcMsnLocSeg `json:"acMsnLocSeg"`
	// The readiness status expressed in time (minutes) for an aircraft to be airborne
	// after the launch order is received or the time required for a missile unit to
	// assume battle stations.
	AlertStatus int64 `json:"alertStatus"`
	// The AMC number assigned to identify one aircraft from another.
	AmcMsnNum string `json:"amcMsnNum"`
	// WGS-84 latitude of the departure location, in degrees. -90 to 90 degrees
	// (negative values south of equator) for this tasked air mission.
	DepLocLat float64 `json:"depLocLat"`
	// WGS-84 longitude of the departure location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian) for this tasked air mission.
	DepLocLon float64 `json:"depLocLon"`
	// The location or name specified for the departure of the tasked air mission.
	DepLocName string `json:"depLocName"`
	// The departure location specified in UTM (100 meter) coordinates for the tasked
	// air mission.
	DepLocUtm string `json:"depLocUTM"`
	// The time of departure for the tasked air mission in ISO8601 UTC format with
	// millisecond precision.
	DepTime time.Time `json:"depTime" format:"date-time"`
	// A collection of the individual aircraft assigned to this aircraft mission
	// tasking.
	IndAcTasking []AirTaskingOrderFullAcMsnTaskingIndAcTasking `json:"indACTasking"`
	// The commander responsible for the planning and execution of the forces necessary
	// to achieve desired objectives.
	MsnCommander string `json:"msnCommander"`
	// The mission number assigned to this mission.
	MsnNum string `json:"msnNum"`
	// The identifier for the composite set of missions for this operation/exercise.
	PkgID string `json:"pkgId"`
	// The code for the preferred type or designator for a tasked air mission.
	PriMsnType string `json:"priMsnType"`
	// An array of WGS-84 latitude of the recovery locations, in degrees. -90 to 90
	// degrees (negative values south of equator) for this tasked air mission.
	RcvyLocLat []float64 `json:"rcvyLocLat"`
	// An array of WGS-84 longitude of the recovery locations, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian) for this tasked air mission.
	RcvyLocLon []float64 `json:"rcvyLocLon"`
	// An array of locations specified for the recovery of the tasked air mission
	// represented by varying formats.
	RcvyLocName []string `json:"rcvyLocName"`
	// An array of recovery locations specified in UTM (100 meter) coordinates for the
	// tasked air mission.
	RcvyLocUtm []string `json:"rcvyLocUTM"`
	// An array of recovery times for the tasked air mission in ISO8601 UTC format with
	// millisecond precision.
	RcvyTime []time.Time `json:"rcvyTime" format:"date-time"`
	// An indicator of whether a mission is or will be a residual mission.
	ResMsnInd string `json:"resMsnInd" format:"char"`
	// The code for the alternative type of a tasked air mission.
	SecMsnType string `json:"secMsnType"`
	// The tasked units location expressed as an ICAO or a place name.
	UnitLocName string `json:"unitLocName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode    respjson.Field
		TaskedService  respjson.Field
		UnitDesignator respjson.Field
		AcMsnLocSeg    respjson.Field
		AlertStatus    respjson.Field
		AmcMsnNum      respjson.Field
		DepLocLat      respjson.Field
		DepLocLon      respjson.Field
		DepLocName     respjson.Field
		DepLocUtm      respjson.Field
		DepTime        respjson.Field
		IndAcTasking   respjson.Field
		MsnCommander   respjson.Field
		MsnNum         respjson.Field
		PkgID          respjson.Field
		PriMsnType     respjson.Field
		RcvyLocLat     respjson.Field
		RcvyLocLon     respjson.Field
		RcvyLocName    respjson.Field
		RcvyLocUtm     respjson.Field
		RcvyTime       respjson.Field
		ResMsnInd      respjson.Field
		SecMsnType     respjson.Field
		UnitLocName    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTaskingOrderFullAcMsnTasking) RawJSON() string { return r.JSON.raw }
func (r *AirTaskingOrderFullAcMsnTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aircraft mission location information for this aircraft mission
// tasking.
type AirTaskingOrderFullAcMsnTaskingAcMsnLocSeg struct {
	// The start time of this mission in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The code for the priority assigned to this mission.
	AirMsnPri string `json:"airMsnPri"`
	// The altitude for this mission represented as hundreds of feet above MSL.
	Alt int64 `json:"alt"`
	// The radius of the circle around the location being reported in feet.
	AreaGeoRad int64 `json:"areaGeoRad"`
	// The end time of this mission in ISO 8601 UTC format with millisecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The name that identifies the location at which this mission is to be performed.
	// This can be the name of a general target area, orbit, cap point, station, etc.
	MsnLocName string `json:"msnLocName"`
	// The alpha-numeric specified location for this mission specified as a bearing
	// angle in degrees relative to true north and a range in nautical miles (NM).
	MsnLocPtBarT string `json:"msnLocPtBarT"`
	// WGS-84 latitude of the mission location, in degrees. -90 to 90 degrees (negative
	// values south of equator) for this tasked air mission.
	MsnLocPtLat float64 `json:"msnLocPtLat"`
	// WGS-84 longitude of the mission location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian) for this tasked air mission.
	MsnLocPtLon float64 `json:"msnLocPtLon"`
	// The location name for this mission.
	MsnLocPtName string `json:"msnLocPtName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartTime    respjson.Field
		AirMsnPri    respjson.Field
		Alt          respjson.Field
		AreaGeoRad   respjson.Field
		EndTime      respjson.Field
		MsnLocName   respjson.Field
		MsnLocPtBarT respjson.Field
		MsnLocPtLat  respjson.Field
		MsnLocPtLon  respjson.Field
		MsnLocPtName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTaskingOrderFullAcMsnTaskingAcMsnLocSeg) RawJSON() string { return r.JSON.raw }
func (r *AirTaskingOrderFullAcMsnTaskingAcMsnLocSeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that specifies the naval flight operations for this ATO.
type AirTaskingOrderFullAcMsnTaskingIndAcTasking struct {
	// The type and model number for the aircraft. The field may specify a value of an
	// aircraft not yet assigned an aircraft code contained in the aircraft codes list.
	AcftType string `json:"acftType,required"`
	// The call sign assigned to this mission aircraft.
	CallSign string `json:"callSign"`
	// The mode 1 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode1Code string `json:"iffSifMode1Code"`
	// The mode 2 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode2Code string `json:"iffSifMode2Code"`
	// The mode 3 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode3Code string `json:"iffSifMode3Code"`
	// An optional array of link 16 octal track numbers assigned as the primary JTIDS
	// Unit (JU) address for the mission aircraft.
	JuAddress []int64 `json:"juAddress"`
	// The Link 16 abbreviated call sign assigned to the ACA. This is normally the
	// first and last letter and the last two numbers of the call sign.
	Link16CallSign string `json:"link16CallSign"`
	// The number of aircraft participating in this mission.
	NumAcft int64 `json:"numAcft"`
	// The code that indicates the ordinance mix carried on this mission aircraft.
	PriConfigCode string `json:"priConfigCode"`
	// The code for the secondary ordinance mix carried on this mission aircraft.
	SecConfigCode string `json:"secConfigCode"`
	// The TACAN channel assigned to this mission aircraft.
	TacanChan int64 `json:"tacanChan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcftType        respjson.Field
		CallSign        respjson.Field
		IffSifMode1Code respjson.Field
		IffSifMode2Code respjson.Field
		IffSifMode3Code respjson.Field
		JuAddress       respjson.Field
		Link16CallSign  respjson.Field
		NumAcft         respjson.Field
		PriConfigCode   respjson.Field
		SecConfigCode   respjson.Field
		TacanChan       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTaskingOrderFullAcMsnTaskingIndAcTasking) RawJSON() string { return r.JSON.raw }
func (r *AirTaskingOrderFullAcMsnTaskingIndAcTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that details special instructions, important information, guidance,
// and amplifying information regarding this ATO.
type AirTaskingOrderFullGenText struct {
	// The free text that describes the information specific to the text indicator.
	Text string `json:"text"`
	// The indicator for the general text block. Examples include "OPENING REMARKS" and
	// "GENERAL SPINS INFORMATION".
	TextInd string `json:"textInd"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		TextInd     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTaskingOrderFullGenText) RawJSON() string { return r.JSON.raw }
func (r *AirTaskingOrderFullGenText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that specifies the naval flight operations for this ATO.
type AirTaskingOrderFullNavalFltOp struct {
	// The name of a ship or maritime vessel. Specify UNKNOWN if name is not known.
	ShipName string `json:"shipName,required"`
	// The time when flight operations begin in ISO8601 UTC format with millisecond
	// precision.
	FltOpStart time.Time `json:"fltOpStart" format:"date-time"`
	// The time when flight operations end in ISO8601 UTC format with millisecond
	// precision.
	FltOpStop time.Time `json:"fltOpStop" format:"date-time"`
	// An array of times at which an aircraft will be launched and/or recovered in
	// ISO8601 UTC format with millisecond precision.
	SchdLaunchRcvyTime []time.Time `json:"schdLaunchRcvyTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ShipName           respjson.Field
		FltOpStart         respjson.Field
		FltOpStop          respjson.Field
		SchdLaunchRcvyTime respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirTaskingOrderFullNavalFltOp) RawJSON() string { return r.JSON.raw }
func (r *AirTaskingOrderFullNavalFltOp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Beta Version Air Tasking Order: The ATO is used to task air missions, assign
// cross force tasking as well as intraservice tasking.
type AirtaskingorderAbridged struct {
	// The effective begin time for this ATO in ISO 8601 UTC format with millisecond
	// precision.
	BeginTs time.Time `json:"beginTs,required" format:"date-time"`
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
	DataMode AirtaskingorderAbridgedDataMode `json:"dataMode,required"`
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExerName string `json:"opExerName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The indicator specifying an affirmative or a negatice condition for this
	// message.
	AckReqInd string `json:"ackReqInd"`
	// Specifies textual data amplifying the data contained in the acknowledgement
	// requirement indicator (ackRedInd) field or the unit required to acknowledge.
	AckUnitInstructions string `json:"ackUnitInstructions"`
	// A collection that specifies the tasked country, tasked service, unit and mission
	// level tasking for this ATO.
	AcMsnTasking []AirtaskingorderAbridgedAcMsnTasking `json:"acMsnTasking"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// The effective end time for this ATO in ISO 8601 UTC format with millisecond
	// precision.
	EndTs time.Time `json:"endTs" format:"date-time"`
	// A collection that details special instructions, important information, guidance,
	// and amplifying information regarding this ATO.
	GenText []AirtaskingorderAbridgedGenText `json:"genText"`
	// The month in which the message originated.
	MsgMonth string `json:"msgMonth"`
	// The identifier of the originator of the message.
	MsgOriginator string `json:"msgOriginator"`
	// The qualifier which caveats the message status.
	MsgQualifier string `json:"msgQualifier"`
	// The unique message identifier sequentially assigned by the originator.
	MsgSn string `json:"msgSN"`
	// A collection that specifies the naval flight operations for this ATO.
	NavalFltOps []AirtaskingorderAbridgedNavalFltOp `json:"navalFltOps"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BeginTs               respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		OpExerName            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AckReqInd             respjson.Field
		AckUnitInstructions   respjson.Field
		AcMsnTasking          respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndTs                 respjson.Field
		GenText               respjson.Field
		MsgMonth              respjson.Field
		MsgOriginator         respjson.Field
		MsgQualifier          respjson.Field
		MsgSn                 respjson.Field
		NavalFltOps           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		RawFileUri            respjson.Field
		SourceDl              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirtaskingorderAbridged) RawJSON() string { return r.JSON.raw }
func (r *AirtaskingorderAbridged) UnmarshalJSON(data []byte) error {
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
type AirtaskingorderAbridgedDataMode string

const (
	AirtaskingorderAbridgedDataModeReal      AirtaskingorderAbridgedDataMode = "REAL"
	AirtaskingorderAbridgedDataModeTest      AirtaskingorderAbridgedDataMode = "TEST"
	AirtaskingorderAbridgedDataModeSimulated AirtaskingorderAbridgedDataMode = "SIMULATED"
	AirtaskingorderAbridgedDataModeExercise  AirtaskingorderAbridgedDataMode = "EXERCISE"
)

// Collection that specifies the tasked country, tasked service, unit and mission
// level tasking for this ATO.
type AirtaskingorderAbridgedAcMsnTasking struct {
	// The country code responsible for conducting this aircraft mission tasking for
	// the exercise or operation.
	CountryCode string `json:"countryCode,required"`
	// The service tasked with conducting this aircraft mission tasking for the
	// exercise or operation.
	TaskedService string `json:"taskedService,required"`
	// The designator of the unit that is tasked to perform this aircraft mission
	// tasking.
	UnitDesignator string `json:"unitDesignator,required"`
	// A collection of aircraft mission location information for this aircraft mission
	// tasking.
	AcMsnLocSeg []AirtaskingorderAbridgedAcMsnTaskingAcMsnLocSeg `json:"acMsnLocSeg"`
	// The readiness status expressed in time (minutes) for an aircraft to be airborne
	// after the launch order is received or the time required for a missile unit to
	// assume battle stations.
	AlertStatus int64 `json:"alertStatus"`
	// The AMC number assigned to identify one aircraft from another.
	AmcMsnNum string `json:"amcMsnNum"`
	// WGS-84 latitude of the departure location, in degrees. -90 to 90 degrees
	// (negative values south of equator) for this tasked air mission.
	DepLocLat float64 `json:"depLocLat"`
	// WGS-84 longitude of the departure location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian) for this tasked air mission.
	DepLocLon float64 `json:"depLocLon"`
	// The location or name specified for the departure of the tasked air mission.
	DepLocName string `json:"depLocName"`
	// The departure location specified in UTM (100 meter) coordinates for the tasked
	// air mission.
	DepLocUtm string `json:"depLocUTM"`
	// The time of departure for the tasked air mission in ISO8601 UTC format with
	// millisecond precision.
	DepTime time.Time `json:"depTime" format:"date-time"`
	// A collection of the individual aircraft assigned to this aircraft mission
	// tasking.
	IndAcTasking []AirtaskingorderAbridgedAcMsnTaskingIndAcTasking `json:"indACTasking"`
	// The commander responsible for the planning and execution of the forces necessary
	// to achieve desired objectives.
	MsnCommander string `json:"msnCommander"`
	// The mission number assigned to this mission.
	MsnNum string `json:"msnNum"`
	// The identifier for the composite set of missions for this operation/exercise.
	PkgID string `json:"pkgId"`
	// The code for the preferred type or designator for a tasked air mission.
	PriMsnType string `json:"priMsnType"`
	// An array of WGS-84 latitude of the recovery locations, in degrees. -90 to 90
	// degrees (negative values south of equator) for this tasked air mission.
	RcvyLocLat []float64 `json:"rcvyLocLat"`
	// An array of WGS-84 longitude of the recovery locations, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian) for this tasked air mission.
	RcvyLocLon []float64 `json:"rcvyLocLon"`
	// An array of locations specified for the recovery of the tasked air mission
	// represented by varying formats.
	RcvyLocName []string `json:"rcvyLocName"`
	// An array of recovery locations specified in UTM (100 meter) coordinates for the
	// tasked air mission.
	RcvyLocUtm []string `json:"rcvyLocUTM"`
	// An array of recovery times for the tasked air mission in ISO8601 UTC format with
	// millisecond precision.
	RcvyTime []time.Time `json:"rcvyTime" format:"date-time"`
	// An indicator of whether a mission is or will be a residual mission.
	ResMsnInd string `json:"resMsnInd" format:"char"`
	// The code for the alternative type of a tasked air mission.
	SecMsnType string `json:"secMsnType"`
	// The tasked units location expressed as an ICAO or a place name.
	UnitLocName string `json:"unitLocName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode    respjson.Field
		TaskedService  respjson.Field
		UnitDesignator respjson.Field
		AcMsnLocSeg    respjson.Field
		AlertStatus    respjson.Field
		AmcMsnNum      respjson.Field
		DepLocLat      respjson.Field
		DepLocLon      respjson.Field
		DepLocName     respjson.Field
		DepLocUtm      respjson.Field
		DepTime        respjson.Field
		IndAcTasking   respjson.Field
		MsnCommander   respjson.Field
		MsnNum         respjson.Field
		PkgID          respjson.Field
		PriMsnType     respjson.Field
		RcvyLocLat     respjson.Field
		RcvyLocLon     respjson.Field
		RcvyLocName    respjson.Field
		RcvyLocUtm     respjson.Field
		RcvyTime       respjson.Field
		ResMsnInd      respjson.Field
		SecMsnType     respjson.Field
		UnitLocName    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirtaskingorderAbridgedAcMsnTasking) RawJSON() string { return r.JSON.raw }
func (r *AirtaskingorderAbridgedAcMsnTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aircraft mission location information for this aircraft mission
// tasking.
type AirtaskingorderAbridgedAcMsnTaskingAcMsnLocSeg struct {
	// The start time of this mission in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The code for the priority assigned to this mission.
	AirMsnPri string `json:"airMsnPri"`
	// The altitude for this mission represented as hundreds of feet above MSL.
	Alt int64 `json:"alt"`
	// The radius of the circle around the location being reported in feet.
	AreaGeoRad int64 `json:"areaGeoRad"`
	// The end time of this mission in ISO 8601 UTC format with millisecond precision.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// The name that identifies the location at which this mission is to be performed.
	// This can be the name of a general target area, orbit, cap point, station, etc.
	MsnLocName string `json:"msnLocName"`
	// The alpha-numeric specified location for this mission specified as a bearing
	// angle in degrees relative to true north and a range in nautical miles (NM).
	MsnLocPtBarT string `json:"msnLocPtBarT"`
	// WGS-84 latitude of the mission location, in degrees. -90 to 90 degrees (negative
	// values south of equator) for this tasked air mission.
	MsnLocPtLat float64 `json:"msnLocPtLat"`
	// WGS-84 longitude of the mission location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian) for this tasked air mission.
	MsnLocPtLon float64 `json:"msnLocPtLon"`
	// The location name for this mission.
	MsnLocPtName string `json:"msnLocPtName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartTime    respjson.Field
		AirMsnPri    respjson.Field
		Alt          respjson.Field
		AreaGeoRad   respjson.Field
		EndTime      respjson.Field
		MsnLocName   respjson.Field
		MsnLocPtBarT respjson.Field
		MsnLocPtLat  respjson.Field
		MsnLocPtLon  respjson.Field
		MsnLocPtName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirtaskingorderAbridgedAcMsnTaskingAcMsnLocSeg) RawJSON() string { return r.JSON.raw }
func (r *AirtaskingorderAbridgedAcMsnTaskingAcMsnLocSeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that specifies the naval flight operations for this ATO.
type AirtaskingorderAbridgedAcMsnTaskingIndAcTasking struct {
	// The type and model number for the aircraft. The field may specify a value of an
	// aircraft not yet assigned an aircraft code contained in the aircraft codes list.
	AcftType string `json:"acftType,required"`
	// The call sign assigned to this mission aircraft.
	CallSign string `json:"callSign"`
	// The mode 1 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode1Code string `json:"iffSifMode1Code"`
	// The mode 2 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode2Code string `json:"iffSifMode2Code"`
	// The mode 3 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode3Code string `json:"iffSifMode3Code"`
	// An optional array of link 16 octal track numbers assigned as the primary JTIDS
	// Unit (JU) address for the mission aircraft.
	JuAddress []int64 `json:"juAddress"`
	// The Link 16 abbreviated call sign assigned to the ACA. This is normally the
	// first and last letter and the last two numbers of the call sign.
	Link16CallSign string `json:"link16CallSign"`
	// The number of aircraft participating in this mission.
	NumAcft int64 `json:"numAcft"`
	// The code that indicates the ordinance mix carried on this mission aircraft.
	PriConfigCode string `json:"priConfigCode"`
	// The code for the secondary ordinance mix carried on this mission aircraft.
	SecConfigCode string `json:"secConfigCode"`
	// The TACAN channel assigned to this mission aircraft.
	TacanChan int64 `json:"tacanChan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcftType        respjson.Field
		CallSign        respjson.Field
		IffSifMode1Code respjson.Field
		IffSifMode2Code respjson.Field
		IffSifMode3Code respjson.Field
		JuAddress       respjson.Field
		Link16CallSign  respjson.Field
		NumAcft         respjson.Field
		PriConfigCode   respjson.Field
		SecConfigCode   respjson.Field
		TacanChan       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirtaskingorderAbridgedAcMsnTaskingIndAcTasking) RawJSON() string { return r.JSON.raw }
func (r *AirtaskingorderAbridgedAcMsnTaskingIndAcTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that details special instructions, important information, guidance,
// and amplifying information regarding this ATO.
type AirtaskingorderAbridgedGenText struct {
	// The free text that describes the information specific to the text indicator.
	Text string `json:"text"`
	// The indicator for the general text block. Examples include "OPENING REMARKS" and
	// "GENERAL SPINS INFORMATION".
	TextInd string `json:"textInd"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		TextInd     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirtaskingorderAbridgedGenText) RawJSON() string { return r.JSON.raw }
func (r *AirtaskingorderAbridgedGenText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that specifies the naval flight operations for this ATO.
type AirtaskingorderAbridgedNavalFltOp struct {
	// The name of a ship or maritime vessel. Specify UNKNOWN if name is not known.
	ShipName string `json:"shipName,required"`
	// The time when flight operations begin in ISO8601 UTC format with millisecond
	// precision.
	FltOpStart time.Time `json:"fltOpStart" format:"date-time"`
	// The time when flight operations end in ISO8601 UTC format with millisecond
	// precision.
	FltOpStop time.Time `json:"fltOpStop" format:"date-time"`
	// An array of times at which an aircraft will be launched and/or recovered in
	// ISO8601 UTC format with millisecond precision.
	SchdLaunchRcvyTime []time.Time `json:"schdLaunchRcvyTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ShipName           respjson.Field
		FltOpStart         respjson.Field
		FltOpStop          respjson.Field
		SchdLaunchRcvyTime respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirtaskingorderAbridgedNavalFltOp) RawJSON() string { return r.JSON.raw }
func (r *AirtaskingorderAbridgedNavalFltOp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirOperationAirTaskingOrderQueryHelpResponse struct {
	AodrSupported         bool                                                    `json:"aodrSupported"`
	ClassificationMarking string                                                  `json:"classificationMarking"`
	Description           string                                                  `json:"description"`
	HistorySupported      bool                                                    `json:"historySupported"`
	Name                  string                                                  `json:"name"`
	Parameters            []AirOperationAirTaskingOrderQueryHelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                                `json:"requiredRoles"`
	RestSupported         bool                                                    `json:"restSupported"`
	SortSupported         bool                                                    `json:"sortSupported"`
	TypeName              string                                                  `json:"typeName"`
	Uri                   string                                                  `json:"uri"`
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
func (r AirOperationAirTaskingOrderQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *AirOperationAirTaskingOrderQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirOperationAirTaskingOrderQueryHelpResponseParameter struct {
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
func (r AirOperationAirTaskingOrderQueryHelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *AirOperationAirTaskingOrderQueryHelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirOperationAirTaskingOrderNewParams struct {
	// The effective begin time for this ATO in ISO 8601 UTC format with millisecond
	// precision.
	BeginTs time.Time `json:"beginTs,required" format:"date-time"`
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
	DataMode AirOperationAirTaskingOrderNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExerName string `json:"opExerName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The indicator specifying an affirmative or a negatice condition for this
	// message.
	AckReqInd param.Opt[string] `json:"ackReqInd,omitzero"`
	// Specifies textual data amplifying the data contained in the acknowledgement
	// requirement indicator (ackRedInd) field or the unit required to acknowledge.
	AckUnitInstructions param.Opt[string] `json:"ackUnitInstructions,omitzero"`
	// The effective end time for this ATO in ISO 8601 UTC format with millisecond
	// precision.
	EndTs param.Opt[time.Time] `json:"endTs,omitzero" format:"date-time"`
	// The month in which the message originated.
	MsgMonth param.Opt[string] `json:"msgMonth,omitzero"`
	// The identifier of the originator of the message.
	MsgOriginator param.Opt[string] `json:"msgOriginator,omitzero"`
	// The qualifier which caveats the message status.
	MsgQualifier param.Opt[string] `json:"msgQualifier,omitzero"`
	// The unique message identifier sequentially assigned by the originator.
	MsgSn param.Opt[string] `json:"msgSN,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// A collection that specifies the tasked country, tasked service, unit and mission
	// level tasking for this ATO.
	AcMsnTasking []AirOperationAirTaskingOrderNewParamsAcMsnTasking `json:"acMsnTasking,omitzero"`
	// A collection that details special instructions, important information, guidance,
	// and amplifying information regarding this ATO.
	GenText []AirOperationAirTaskingOrderNewParamsGenText `json:"genText,omitzero"`
	// A collection that specifies the naval flight operations for this ATO.
	NavalFltOps []AirOperationAirTaskingOrderNewParamsNavalFltOp `json:"navalFltOps,omitzero"`
	paramObj
}

func (r AirOperationAirTaskingOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderNewParams) UnmarshalJSON(data []byte) error {
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
type AirOperationAirTaskingOrderNewParamsDataMode string

const (
	AirOperationAirTaskingOrderNewParamsDataModeReal      AirOperationAirTaskingOrderNewParamsDataMode = "REAL"
	AirOperationAirTaskingOrderNewParamsDataModeTest      AirOperationAirTaskingOrderNewParamsDataMode = "TEST"
	AirOperationAirTaskingOrderNewParamsDataModeSimulated AirOperationAirTaskingOrderNewParamsDataMode = "SIMULATED"
	AirOperationAirTaskingOrderNewParamsDataModeExercise  AirOperationAirTaskingOrderNewParamsDataMode = "EXERCISE"
)

// Collection that specifies the tasked country, tasked service, unit and mission
// level tasking for this ATO.
//
// The properties CountryCode, TaskedService, UnitDesignator are required.
type AirOperationAirTaskingOrderNewParamsAcMsnTasking struct {
	// The country code responsible for conducting this aircraft mission tasking for
	// the exercise or operation.
	CountryCode string `json:"countryCode,required"`
	// The service tasked with conducting this aircraft mission tasking for the
	// exercise or operation.
	TaskedService string `json:"taskedService,required"`
	// The designator of the unit that is tasked to perform this aircraft mission
	// tasking.
	UnitDesignator string `json:"unitDesignator,required"`
	// The readiness status expressed in time (minutes) for an aircraft to be airborne
	// after the launch order is received or the time required for a missile unit to
	// assume battle stations.
	AlertStatus param.Opt[int64] `json:"alertStatus,omitzero"`
	// The AMC number assigned to identify one aircraft from another.
	AmcMsnNum param.Opt[string] `json:"amcMsnNum,omitzero"`
	// WGS-84 latitude of the departure location, in degrees. -90 to 90 degrees
	// (negative values south of equator) for this tasked air mission.
	DepLocLat param.Opt[float64] `json:"depLocLat,omitzero"`
	// WGS-84 longitude of the departure location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian) for this tasked air mission.
	DepLocLon param.Opt[float64] `json:"depLocLon,omitzero"`
	// The location or name specified for the departure of the tasked air mission.
	DepLocName param.Opt[string] `json:"depLocName,omitzero"`
	// The departure location specified in UTM (100 meter) coordinates for the tasked
	// air mission.
	DepLocUtm param.Opt[string] `json:"depLocUTM,omitzero"`
	// The time of departure for the tasked air mission in ISO8601 UTC format with
	// millisecond precision.
	DepTime param.Opt[time.Time] `json:"depTime,omitzero" format:"date-time"`
	// The commander responsible for the planning and execution of the forces necessary
	// to achieve desired objectives.
	MsnCommander param.Opt[string] `json:"msnCommander,omitzero"`
	// The mission number assigned to this mission.
	MsnNum param.Opt[string] `json:"msnNum,omitzero"`
	// The identifier for the composite set of missions for this operation/exercise.
	PkgID param.Opt[string] `json:"pkgId,omitzero"`
	// The code for the preferred type or designator for a tasked air mission.
	PriMsnType param.Opt[string] `json:"priMsnType,omitzero"`
	// An indicator of whether a mission is or will be a residual mission.
	ResMsnInd param.Opt[string] `json:"resMsnInd,omitzero" format:"char"`
	// The code for the alternative type of a tasked air mission.
	SecMsnType param.Opt[string] `json:"secMsnType,omitzero"`
	// The tasked units location expressed as an ICAO or a place name.
	UnitLocName param.Opt[string] `json:"unitLocName,omitzero"`
	// A collection of aircraft mission location information for this aircraft mission
	// tasking.
	AcMsnLocSeg []AirOperationAirTaskingOrderNewParamsAcMsnTaskingAcMsnLocSeg `json:"acMsnLocSeg,omitzero"`
	// A collection of the individual aircraft assigned to this aircraft mission
	// tasking.
	IndAcTasking []AirOperationAirTaskingOrderNewParamsAcMsnTaskingIndAcTasking `json:"indACTasking,omitzero"`
	// An array of WGS-84 latitude of the recovery locations, in degrees. -90 to 90
	// degrees (negative values south of equator) for this tasked air mission.
	RcvyLocLat []float64 `json:"rcvyLocLat,omitzero"`
	// An array of WGS-84 longitude of the recovery locations, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian) for this tasked air mission.
	RcvyLocLon []float64 `json:"rcvyLocLon,omitzero"`
	// An array of locations specified for the recovery of the tasked air mission
	// represented by varying formats.
	RcvyLocName []string `json:"rcvyLocName,omitzero"`
	// An array of recovery locations specified in UTM (100 meter) coordinates for the
	// tasked air mission.
	RcvyLocUtm []string `json:"rcvyLocUTM,omitzero"`
	// An array of recovery times for the tasked air mission in ISO8601 UTC format with
	// millisecond precision.
	RcvyTime []time.Time `json:"rcvyTime,omitzero" format:"date-time"`
	paramObj
}

func (r AirOperationAirTaskingOrderNewParamsAcMsnTasking) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderNewParamsAcMsnTasking
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderNewParamsAcMsnTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aircraft mission location information for this aircraft mission
// tasking.
//
// The property StartTime is required.
type AirOperationAirTaskingOrderNewParamsAcMsnTaskingAcMsnLocSeg struct {
	// The start time of this mission in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The code for the priority assigned to this mission.
	AirMsnPri param.Opt[string] `json:"airMsnPri,omitzero"`
	// The altitude for this mission represented as hundreds of feet above MSL.
	Alt param.Opt[int64] `json:"alt,omitzero"`
	// The radius of the circle around the location being reported in feet.
	AreaGeoRad param.Opt[int64] `json:"areaGeoRad,omitzero"`
	// The end time of this mission in ISO 8601 UTC format with millisecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The name that identifies the location at which this mission is to be performed.
	// This can be the name of a general target area, orbit, cap point, station, etc.
	MsnLocName param.Opt[string] `json:"msnLocName,omitzero"`
	// The alpha-numeric specified location for this mission specified as a bearing
	// angle in degrees relative to true north and a range in nautical miles (NM).
	MsnLocPtBarT param.Opt[string] `json:"msnLocPtBarT,omitzero"`
	// WGS-84 latitude of the mission location, in degrees. -90 to 90 degrees (negative
	// values south of equator) for this tasked air mission.
	MsnLocPtLat param.Opt[float64] `json:"msnLocPtLat,omitzero"`
	// WGS-84 longitude of the mission location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian) for this tasked air mission.
	MsnLocPtLon param.Opt[float64] `json:"msnLocPtLon,omitzero"`
	// The location name for this mission.
	MsnLocPtName param.Opt[string] `json:"msnLocPtName,omitzero"`
	paramObj
}

func (r AirOperationAirTaskingOrderNewParamsAcMsnTaskingAcMsnLocSeg) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderNewParamsAcMsnTaskingAcMsnLocSeg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderNewParamsAcMsnTaskingAcMsnLocSeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that specifies the naval flight operations for this ATO.
//
// The property AcftType is required.
type AirOperationAirTaskingOrderNewParamsAcMsnTaskingIndAcTasking struct {
	// The type and model number for the aircraft. The field may specify a value of an
	// aircraft not yet assigned an aircraft code contained in the aircraft codes list.
	AcftType string `json:"acftType,required"`
	// The call sign assigned to this mission aircraft.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// The mode 1 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode1Code param.Opt[string] `json:"iffSifMode1Code,omitzero"`
	// The mode 2 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode2Code param.Opt[string] `json:"iffSifMode2Code,omitzero"`
	// The mode 3 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode3Code param.Opt[string] `json:"iffSifMode3Code,omitzero"`
	// The Link 16 abbreviated call sign assigned to the ACA. This is normally the
	// first and last letter and the last two numbers of the call sign.
	Link16CallSign param.Opt[string] `json:"link16CallSign,omitzero"`
	// The number of aircraft participating in this mission.
	NumAcft param.Opt[int64] `json:"numAcft,omitzero"`
	// The code that indicates the ordinance mix carried on this mission aircraft.
	PriConfigCode param.Opt[string] `json:"priConfigCode,omitzero"`
	// The code for the secondary ordinance mix carried on this mission aircraft.
	SecConfigCode param.Opt[string] `json:"secConfigCode,omitzero"`
	// The TACAN channel assigned to this mission aircraft.
	TacanChan param.Opt[int64] `json:"tacanChan,omitzero"`
	// An optional array of link 16 octal track numbers assigned as the primary JTIDS
	// Unit (JU) address for the mission aircraft.
	JuAddress []int64 `json:"juAddress,omitzero"`
	paramObj
}

func (r AirOperationAirTaskingOrderNewParamsAcMsnTaskingIndAcTasking) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderNewParamsAcMsnTaskingIndAcTasking
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderNewParamsAcMsnTaskingIndAcTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that details special instructions, important information, guidance,
// and amplifying information regarding this ATO.
type AirOperationAirTaskingOrderNewParamsGenText struct {
	// The free text that describes the information specific to the text indicator.
	Text param.Opt[string] `json:"text,omitzero"`
	// The indicator for the general text block. Examples include "OPENING REMARKS" and
	// "GENERAL SPINS INFORMATION".
	TextInd param.Opt[string] `json:"textInd,omitzero"`
	paramObj
}

func (r AirOperationAirTaskingOrderNewParamsGenText) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderNewParamsGenText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderNewParamsGenText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that specifies the naval flight operations for this ATO.
//
// The property ShipName is required.
type AirOperationAirTaskingOrderNewParamsNavalFltOp struct {
	// The name of a ship or maritime vessel. Specify UNKNOWN if name is not known.
	ShipName string `json:"shipName,required"`
	// The time when flight operations begin in ISO8601 UTC format with millisecond
	// precision.
	FltOpStart param.Opt[time.Time] `json:"fltOpStart,omitzero" format:"date-time"`
	// The time when flight operations end in ISO8601 UTC format with millisecond
	// precision.
	FltOpStop param.Opt[time.Time] `json:"fltOpStop,omitzero" format:"date-time"`
	// An array of times at which an aircraft will be launched and/or recovered in
	// ISO8601 UTC format with millisecond precision.
	SchdLaunchRcvyTime []time.Time `json:"schdLaunchRcvyTime,omitzero" format:"date-time"`
	paramObj
}

func (r AirOperationAirTaskingOrderNewParamsNavalFltOp) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderNewParamsNavalFltOp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderNewParamsNavalFltOp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirOperationAirTaskingOrderGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationAirTaskingOrderGetParams]'s query parameters as
// `url.Values`.
func (r AirOperationAirTaskingOrderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirOperationAirTaskingOrderListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationAirTaskingOrderListParams]'s query parameters
// as `url.Values`.
func (r AirOperationAirTaskingOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirOperationAirTaskingOrderCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationAirTaskingOrderCountParams]'s query parameters
// as `url.Values`.
func (r AirOperationAirTaskingOrderCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirOperationAirTaskingOrderTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationAirTaskingOrderTupleParams]'s query parameters
// as `url.Values`.
func (r AirOperationAirTaskingOrderTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirOperationAirTaskingOrderUnvalidatedPublishParams struct {
	Body []AirOperationAirTaskingOrderUnvalidatedPublishParamsBody
	paramObj
}

func (r AirOperationAirTaskingOrderUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *AirOperationAirTaskingOrderUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Beta Version Air Tasking Order: The ATO is used to task air missions, assign
// cross force tasking as well as intraservice tasking.
//
// The properties BeginTs, ClassificationMarking, DataMode, OpExerName, Source are
// required.
type AirOperationAirTaskingOrderUnvalidatedPublishParamsBody struct {
	// The effective begin time for this ATO in ISO 8601 UTC format with millisecond
	// precision.
	BeginTs time.Time `json:"beginTs,required" format:"date-time"`
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
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExerName string `json:"opExerName,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The indicator specifying an affirmative or a negatice condition for this
	// message.
	AckReqInd param.Opt[string] `json:"ackReqInd,omitzero"`
	// Specifies textual data amplifying the data contained in the acknowledgement
	// requirement indicator (ackRedInd) field or the unit required to acknowledge.
	AckUnitInstructions param.Opt[string] `json:"ackUnitInstructions,omitzero"`
	// Time the row was created in the database.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The effective end time for this ATO in ISO 8601 UTC format with millisecond
	// precision.
	EndTs param.Opt[time.Time] `json:"endTs,omitzero" format:"date-time"`
	// The month in which the message originated.
	MsgMonth param.Opt[string] `json:"msgMonth,omitzero"`
	// The identifier of the originator of the message.
	MsgOriginator param.Opt[string] `json:"msgOriginator,omitzero"`
	// The qualifier which caveats the message status.
	MsgQualifier param.Opt[string] `json:"msgQualifier,omitzero"`
	// The unique message identifier sequentially assigned by the originator.
	MsgSn param.Opt[string] `json:"msgSN,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// A collection that specifies the tasked country, tasked service, unit and mission
	// level tasking for this ATO.
	AcMsnTasking []AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTasking `json:"acMsnTasking,omitzero"`
	// A collection that details special instructions, important information, guidance,
	// and amplifying information regarding this ATO.
	GenText []AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyGenText `json:"genText,omitzero"`
	// A collection that specifies the naval flight operations for this ATO.
	NavalFltOps []AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyNavalFltOp `json:"navalFltOps,omitzero"`
	paramObj
}

func (r AirOperationAirTaskingOrderUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AirOperationAirTaskingOrderUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection that specifies the tasked country, tasked service, unit and mission
// level tasking for this ATO.
//
// The properties CountryCode, TaskedService, UnitDesignator are required.
type AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTasking struct {
	// The country code responsible for conducting this aircraft mission tasking for
	// the exercise or operation.
	CountryCode string `json:"countryCode,required"`
	// The service tasked with conducting this aircraft mission tasking for the
	// exercise or operation.
	TaskedService string `json:"taskedService,required"`
	// The designator of the unit that is tasked to perform this aircraft mission
	// tasking.
	UnitDesignator string `json:"unitDesignator,required"`
	// The readiness status expressed in time (minutes) for an aircraft to be airborne
	// after the launch order is received or the time required for a missile unit to
	// assume battle stations.
	AlertStatus param.Opt[int64] `json:"alertStatus,omitzero"`
	// The AMC number assigned to identify one aircraft from another.
	AmcMsnNum param.Opt[string] `json:"amcMsnNum,omitzero"`
	// WGS-84 latitude of the departure location, in degrees. -90 to 90 degrees
	// (negative values south of equator) for this tasked air mission.
	DepLocLat param.Opt[float64] `json:"depLocLat,omitzero"`
	// WGS-84 longitude of the departure location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian) for this tasked air mission.
	DepLocLon param.Opt[float64] `json:"depLocLon,omitzero"`
	// The location or name specified for the departure of the tasked air mission.
	DepLocName param.Opt[string] `json:"depLocName,omitzero"`
	// The departure location specified in UTM (100 meter) coordinates for the tasked
	// air mission.
	DepLocUtm param.Opt[string] `json:"depLocUTM,omitzero"`
	// The time of departure for the tasked air mission in ISO8601 UTC format with
	// millisecond precision.
	DepTime param.Opt[time.Time] `json:"depTime,omitzero" format:"date-time"`
	// The commander responsible for the planning and execution of the forces necessary
	// to achieve desired objectives.
	MsnCommander param.Opt[string] `json:"msnCommander,omitzero"`
	// The mission number assigned to this mission.
	MsnNum param.Opt[string] `json:"msnNum,omitzero"`
	// The identifier for the composite set of missions for this operation/exercise.
	PkgID param.Opt[string] `json:"pkgId,omitzero"`
	// The code for the preferred type or designator for a tasked air mission.
	PriMsnType param.Opt[string] `json:"priMsnType,omitzero"`
	// An indicator of whether a mission is or will be a residual mission.
	ResMsnInd param.Opt[string] `json:"resMsnInd,omitzero" format:"char"`
	// The code for the alternative type of a tasked air mission.
	SecMsnType param.Opt[string] `json:"secMsnType,omitzero"`
	// The tasked units location expressed as an ICAO or a place name.
	UnitLocName param.Opt[string] `json:"unitLocName,omitzero"`
	// A collection of aircraft mission location information for this aircraft mission
	// tasking.
	AcMsnLocSeg []AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingAcMsnLocSeg `json:"acMsnLocSeg,omitzero"`
	// A collection of the individual aircraft assigned to this aircraft mission
	// tasking.
	IndAcTasking []AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingIndAcTasking `json:"indACTasking,omitzero"`
	// An array of WGS-84 latitude of the recovery locations, in degrees. -90 to 90
	// degrees (negative values south of equator) for this tasked air mission.
	RcvyLocLat []float64 `json:"rcvyLocLat,omitzero"`
	// An array of WGS-84 longitude of the recovery locations, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian) for this tasked air mission.
	RcvyLocLon []float64 `json:"rcvyLocLon,omitzero"`
	// An array of locations specified for the recovery of the tasked air mission
	// represented by varying formats.
	RcvyLocName []string `json:"rcvyLocName,omitzero"`
	// An array of recovery locations specified in UTM (100 meter) coordinates for the
	// tasked air mission.
	RcvyLocUtm []string `json:"rcvyLocUTM,omitzero"`
	// An array of recovery times for the tasked air mission in ISO8601 UTC format with
	// millisecond precision.
	RcvyTime []time.Time `json:"rcvyTime,omitzero" format:"date-time"`
	paramObj
}

func (r AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTasking) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTasking
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aircraft mission location information for this aircraft mission
// tasking.
//
// The property StartTime is required.
type AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingAcMsnLocSeg struct {
	// The start time of this mission in ISO 8601 UTC format with millisecond
	// precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// The code for the priority assigned to this mission.
	AirMsnPri param.Opt[string] `json:"airMsnPri,omitzero"`
	// The altitude for this mission represented as hundreds of feet above MSL.
	Alt param.Opt[int64] `json:"alt,omitzero"`
	// The radius of the circle around the location being reported in feet.
	AreaGeoRad param.Opt[int64] `json:"areaGeoRad,omitzero"`
	// The end time of this mission in ISO 8601 UTC format with millisecond precision.
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// The name that identifies the location at which this mission is to be performed.
	// This can be the name of a general target area, orbit, cap point, station, etc.
	MsnLocName param.Opt[string] `json:"msnLocName,omitzero"`
	// The alpha-numeric specified location for this mission specified as a bearing
	// angle in degrees relative to true north and a range in nautical miles (NM).
	MsnLocPtBarT param.Opt[string] `json:"msnLocPtBarT,omitzero"`
	// WGS-84 latitude of the mission location, in degrees. -90 to 90 degrees (negative
	// values south of equator) for this tasked air mission.
	MsnLocPtLat param.Opt[float64] `json:"msnLocPtLat,omitzero"`
	// WGS-84 longitude of the mission location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian) for this tasked air mission.
	MsnLocPtLon param.Opt[float64] `json:"msnLocPtLon,omitzero"`
	// The location name for this mission.
	MsnLocPtName param.Opt[string] `json:"msnLocPtName,omitzero"`
	paramObj
}

func (r AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingAcMsnLocSeg) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingAcMsnLocSeg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingAcMsnLocSeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that specifies the naval flight operations for this ATO.
//
// The property AcftType is required.
type AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingIndAcTasking struct {
	// The type and model number for the aircraft. The field may specify a value of an
	// aircraft not yet assigned an aircraft code contained in the aircraft codes list.
	AcftType string `json:"acftType,required"`
	// The call sign assigned to this mission aircraft.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// The mode 1 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode1Code param.Opt[string] `json:"iffSifMode1Code,omitzero"`
	// The mode 2 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode2Code param.Opt[string] `json:"iffSifMode2Code,omitzero"`
	// The mode 3 and code of the Identification Friend or FOE (IFF) or Selective
	// Identification Feature (SIF).
	IffSifMode3Code param.Opt[string] `json:"iffSifMode3Code,omitzero"`
	// The Link 16 abbreviated call sign assigned to the ACA. This is normally the
	// first and last letter and the last two numbers of the call sign.
	Link16CallSign param.Opt[string] `json:"link16CallSign,omitzero"`
	// The number of aircraft participating in this mission.
	NumAcft param.Opt[int64] `json:"numAcft,omitzero"`
	// The code that indicates the ordinance mix carried on this mission aircraft.
	PriConfigCode param.Opt[string] `json:"priConfigCode,omitzero"`
	// The code for the secondary ordinance mix carried on this mission aircraft.
	SecConfigCode param.Opt[string] `json:"secConfigCode,omitzero"`
	// The TACAN channel assigned to this mission aircraft.
	TacanChan param.Opt[int64] `json:"tacanChan,omitzero"`
	// An optional array of link 16 octal track numbers assigned as the primary JTIDS
	// Unit (JU) address for the mission aircraft.
	JuAddress []int64 `json:"juAddress,omitzero"`
	paramObj
}

func (r AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingIndAcTasking) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingIndAcTasking
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingIndAcTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that details special instructions, important information, guidance,
// and amplifying information regarding this ATO.
type AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyGenText struct {
	// The free text that describes the information specific to the text indicator.
	Text param.Opt[string] `json:"text,omitzero"`
	// The indicator for the general text block. Examples include "OPENING REMARKS" and
	// "GENERAL SPINS INFORMATION".
	TextInd param.Opt[string] `json:"textInd,omitzero"`
	paramObj
}

func (r AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyGenText) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyGenText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyGenText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection that specifies the naval flight operations for this ATO.
//
// The property ShipName is required.
type AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyNavalFltOp struct {
	// The name of a ship or maritime vessel. Specify UNKNOWN if name is not known.
	ShipName string `json:"shipName,required"`
	// The time when flight operations begin in ISO8601 UTC format with millisecond
	// precision.
	FltOpStart param.Opt[time.Time] `json:"fltOpStart,omitzero" format:"date-time"`
	// The time when flight operations end in ISO8601 UTC format with millisecond
	// precision.
	FltOpStop param.Opt[time.Time] `json:"fltOpStop,omitzero" format:"date-time"`
	// An array of times at which an aircraft will be launched and/or recovered in
	// ISO8601 UTC format with millisecond precision.
	SchdLaunchRcvyTime []time.Time `json:"schdLaunchRcvyTime,omitzero" format:"date-time"`
	paramObj
}

func (r AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyNavalFltOp) MarshalJSON() (data []byte, err error) {
	type shadow AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyNavalFltOp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyNavalFltOp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
