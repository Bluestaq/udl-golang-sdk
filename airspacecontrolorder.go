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

// AirspaceControlOrderService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirspaceControlOrderService] method instead.
type AirspaceControlOrderService struct {
	Options []option.RequestOption
}

// NewAirspaceControlOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAirspaceControlOrderService(opts ...option.RequestOption) (r AirspaceControlOrderService) {
	r = AirspaceControlOrderService{}
	r.Options = opts
	return
}

// Service operation to take a single AirspaceControlOrder record as a POST body
// and ingest into the database. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *AirspaceControlOrderService) New(ctx context.Context, body AirspaceControlOrderNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airspacecontrolorder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single AirspaceControlOrder record by its unique ID
// passed as a path parameter.
func (r *AirspaceControlOrderService) Get(ctx context.Context, id string, query AirspaceControlOrderGetParams, opts ...option.RequestOption) (res *AirspacecontrolorderFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airspacecontrolorder/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirspaceControlOrderService) List(ctx context.Context, query AirspaceControlOrderListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AirspacecontrolorderAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/airspacecontrolorder"
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
func (r *AirspaceControlOrderService) ListAutoPaging(ctx context.Context, query AirspaceControlOrderListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AirspacecontrolorderAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirspaceControlOrderService) Count(ctx context.Context, query AirspaceControlOrderCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airspacecontrolorder/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// AirspaceControlOrder records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *AirspaceControlOrderService) NewBulk(ctx context.Context, body AirspaceControlOrderNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airspacecontrolorder/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AirspaceControlOrderService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airspacecontrolorder/queryhelp"
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
func (r *AirspaceControlOrderService) Tuple(ctx context.Context, query AirspaceControlOrderTupleParams, opts ...option.RequestOption) (res *[]AirspacecontrolorderFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airspacecontrolorder/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Beta Version Airspace Control Order: Contains airspace coordination information
// and instructions that have been issued by an airspace control authority.
type AirspacecontrolorderAbridged struct {
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
	DataMode AirspacecontrolorderAbridgedDataMode `json:"dataMode,required"`
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExName string `json:"opExName,required"`
	// The identifier of the originator of this message.
	Originator string `json:"originator,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the effective time period of this airspace control order, in ISO
	// 8601 UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Free text information expressed in natural language.
	AcoComments string `json:"acoComments"`
	// The serial number of this airspace control order.
	AcoSerialNum string `json:"acoSerialNum"`
	// Mandatory nested segment to report multiple airspace control means statuses
	// within an ACOID.
	AirspaceControlMeansStatus []AirspacecontrolorderAbridgedAirspaceControlMeansStatus `json:"airspaceControlMeansStatus"`
	// The airspaceControlReferences set provides both USMTF and non-USMTF references
	// for this airspace control order.
	AirspaceControlOrderReferences []AirspacecontrolorderAbridgedAirspaceControlOrderReference `json:"airspaceControlOrderReferences"`
	// Name of the area of the command for which the ACO is valid.
	AreaOfValidity string `json:"areaOfValidity"`
	// Mandatory if classSource uses the "IORIG" designator. Must be a REASON FOR
	// CLASSIFICATION code.
	ClassReasons []string `json:"classReasons"`
	// Markings defining the source material or the original classification authority
	// for the ACO message.
	ClassSource string `json:"classSource"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Coded entries that provide justification for exemption from automatic
	// downgrading or declassification of the airspace control order.
	DeclassExemptionCodes []string `json:"declassExemptionCodes"`
	// Markings providing the literal guidance or date for downgrading or declassifying
	// the airspace control order.
	DowngradeInsDates []string `json:"downgradeInsDates"`
	// Specifies the geodetic datum by which the spatial coordinates of the controlled
	// airspace are calculated.
	GeoDatum string `json:"geoDatum"`
	// The month in which the message originated.
	Month string `json:"month"`
	// Supplementary name that can be used to further identify exercise nicknames, or
	// to provide the primary nickname of the option or the alternative of an
	// operational plan.
	OpExInfo string `json:"opExInfo"`
	// The secondary supplementary nickname of the option or the alternative of the
	// operational plan or order.
	OpExInfoAlt string `json:"opExInfoAlt"`
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
	// The qualifier which caveats the message status.
	Qualifier string `json:"qualifier"`
	// The serial number associated with the message qualifier.
	QualSn int64 `json:"qualSN"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The unique message identifier sequentially assigned by the originator.
	SerialNum string `json:"serialNum"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// A qualifier for the end of the effective time period of this airspace control
	// order, such as AFTER, ASOF, NLT, etc. Used with field stopTime to indicate a
	// relative time.
	StopQualifier string `json:"stopQualifier"`
	// The end of the effective time period of this airspace control order, in ISO 8601
	// UTC format with millisecond precision.
	StopTime time.Time `json:"stopTime" format:"date-time"`
	// Array of unique link 16 identifiers that will be assigned to a future airspace
	// control means.
	UndLnkTrks []string `json:"undLnkTrks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking          respjson.Field
		DataMode                       respjson.Field
		OpExName                       respjson.Field
		Originator                     respjson.Field
		Source                         respjson.Field
		StartTime                      respjson.Field
		ID                             respjson.Field
		AcoComments                    respjson.Field
		AcoSerialNum                   respjson.Field
		AirspaceControlMeansStatus     respjson.Field
		AirspaceControlOrderReferences respjson.Field
		AreaOfValidity                 respjson.Field
		ClassReasons                   respjson.Field
		ClassSource                    respjson.Field
		CreatedAt                      respjson.Field
		CreatedBy                      respjson.Field
		DeclassExemptionCodes          respjson.Field
		DowngradeInsDates              respjson.Field
		GeoDatum                       respjson.Field
		Month                          respjson.Field
		OpExInfo                       respjson.Field
		OpExInfoAlt                    respjson.Field
		Origin                         respjson.Field
		OrigNetwork                    respjson.Field
		PlanOrigNum                    respjson.Field
		Qualifier                      respjson.Field
		QualSn                         respjson.Field
		RawFileUri                     respjson.Field
		SerialNum                      respjson.Field
		SourceDl                       respjson.Field
		StopQualifier                  respjson.Field
		StopTime                       respjson.Field
		UndLnkTrks                     respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderAbridged) RawJSON() string { return r.JSON.raw }
func (r *AirspacecontrolorderAbridged) UnmarshalJSON(data []byte) error {
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
type AirspacecontrolorderAbridgedDataMode string

const (
	AirspacecontrolorderAbridgedDataModeReal      AirspacecontrolorderAbridgedDataMode = "REAL"
	AirspacecontrolorderAbridgedDataModeTest      AirspacecontrolorderAbridgedDataMode = "TEST"
	AirspacecontrolorderAbridgedDataModeSimulated AirspacecontrolorderAbridgedDataMode = "SIMULATED"
	AirspacecontrolorderAbridgedDataModeExercise  AirspacecontrolorderAbridgedDataMode = "EXERCISE"
)

// Mandatory nested segment to report multiple airspace control means statuses
// within an ACOID.
type AirspacecontrolorderAbridgedAirspaceControlMeansStatus struct {
	// A conditional nested segment to report multiple airspace control means within a
	// particular airspace control means status.
	AirspaceControlMeans []AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMean `json:"airspaceControlMeans"`
	// Status of Airspace Control Means. Must be ADD, CHANGE, or DELETE.
	CmStat string `json:"cmStat"`
	// Airspace control means name or designator. Mandatory if acmStat equals "DELETE,"
	// otherwise this field is prohibited.
	CmStatID []string `json:"cmStatId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AirspaceControlMeans respjson.Field
		CmStat               respjson.Field
		CmStatID             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderAbridgedAirspaceControlMeansStatus) RawJSON() string { return r.JSON.raw }
func (r *AirspacecontrolorderAbridgedAirspaceControlMeansStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A conditional nested segment to report multiple airspace control means within a
// particular airspace control means status.
type AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMean struct {
	// The controlPoint set describes any reference/controlling/rendezvous point for a
	// given airspace control means.
	AirspaceControlPoint []AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint `json:"airspaceControlPoint"`
	// The timePeriods set describes the effective datetime for a given airspace
	// control means.
	AirspaceTimePeriod []AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod `json:"airspaceTimePeriod"`
	// A bearing measured from true North, in angular degrees. If cmShape is set to
	// "POLYARC" or "RADARC", this field is required and is mapped to the "beginning"
	// radial bearing parameter.
	Bearing0 float64 `json:"bearing0"`
	// A bearing measured from true North, in angular degrees. If cmShape is set to
	// "POLYARC" or "RADARC", this field is required and is mapped to the "ending"
	// radial bearing parameter.
	Bearing1 float64 `json:"bearing1"`
	// Airspace control means name or designator.
	CmID string `json:"cmId"`
	// Designates the geometric type that defines the airspace shape. One of CIRCLE,
	// CORRIDOR, LINE, ORBIT, etc.
	//
	// Any of "POLYARC", "1TRACK", "POLYGON", "CIRCLE", "CORRIDOR", "APOINT", "AORBIT",
	// "GEOLINE".
	CmShape string `json:"cmShape"`
	// The code for the type of airspace control means.
	CmType string `json:"cmType"`
	// The commander responsible within a specified geographical area for the airspace
	// control operation assigned to him.
	CntrlAuth string `json:"cntrlAuth"`
	// The frequency for the airspace control authority. Can specify HZ, KHZ, MHZ, GHZ
	// or a DESIG frequency designator code.
	CntrlAuthFreqs []string `json:"cntrlAuthFreqs"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format. The fields coord0 and coord1 should be used in the specification of any
	// airspace control shape that requires exactly one (1) or two (2) reference points
	// for construction. For shapes requiring one reference point, for instance, when
	// cmShape is set to "APOINT", this field is required and singularly defines the
	// shape. Similarly, this field is required to define the center point of a
	// "CIRCLE" shape, or the "origin of bearing" for arcs.
	Coord0 string `json:"coord0"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format. The fields coord0 and coord1 should be used in the specification of any
	// airspace control shape that requires exactly one (1) or two (2) reference points
	// for construction. For shapes requiring one reference point, for instance, when
	// cmShape is set to "APOINT", this field is required and singularly defines the
	// shape. Similarly, this field is required to define the center point of a
	// "CIRCLE" shape, or the "origin of bearing" for arcs.
	Coord1 string `json:"coord1"`
	// An array of at least two alphanumeric symbols used to serially identify the
	// corridor waypoints. If cmShape is set to "CORRIDOR", one of either corrWayPoints
	// or polyCoord is required to specify the centerline of the corridor path.
	CorrWayPoints []string `json:"corrWayPoints"`
	// Description of the airspace vertical dimension.
	EffVDim string `json:"effVDim"`
	// General informat detailing the transit instruction for the airspace control
	// means.
	FreeText string `json:"freeText"`
	// Used to provide transit instructions for the airspace control means.
	GenTextInd string `json:"genTextInd"`
	// Specifies the geodetic datum by which the spatial coordinates of the controlled
	// airspace are calculated, if different from the top level ACO datum.
	GeoDatumAlt string `json:"geoDatumAlt"`
	// Unique Link 16 identifier assigned to the airspace control means.
	Link16ID string `json:"link16Id"`
	// Orbit alignment look-up code. Can be C=Center, L=Left, R=Right.
	OrbitAlignment string `json:"orbitAlignment"`
	// A set of geospatial coordinates specified in DMS (Degrees, Minutes, Seconds)
	// format which determine the vertices of a one or two dimensional geospatial
	// shape. When cmShape is set to "POLYARC" or "POLYGON", this field is required as
	// applied in the construction of the area boundary. If cmShape is set to
	// "CORRIDOR" or "GEOLINE", this field is required and can be interpreted as an
	// ordered set of points along a path in space.
	PolyCoord []string `json:"polyCoord"`
	// A distance that represents a radial magnitude. If cmShape is set to "CIRCLE" or
	// "POLYARC", one of either fields radMag0 or radMag1 is required. If cmShape is
	// set to "RADARC", this field is required and maps to the "inner" radial magnitude
	// arc limit. If provided, the field radMagUnit is required.
	RadMag0 float64 `json:"radMag0"`
	// A distance that represents a radial magnitude. If cmShape is set to "CIRCLE" or
	// "POLYARC", one of either fields radMag0 or radMag1 is required. If cmShape is
	// set to "RADARC", this field is required and maps to the "outer" radial magnitude
	// arc limit. If provided, the field radMagUnit is required.
	RadMag1 float64 `json:"radMag1"`
	// Specifies the unit of length in which radial magnitudes are given. Use M for
	// meters, KM for kilometers, or NM for nautical miles.
	RadMagUnit string `json:"radMagUnit"`
	// Index of a segment in an airtrack, which is defined by an ordered set of points.
	TrackLeg int64 `json:"trackLeg"`
	// The altitude at or below which the vertical position of an aircraft is
	// controlled by reference to true altitude.
	TransAltitude string `json:"transAltitude"`
	// Designates the means by which a defined airspace control means is to be used.
	Usage string `json:"usage"`
	// Used to describe the "side to side" distance of a target, object or area. If
	// cmShape is set to "CORRIDOR" or "AORBIT", this field is required and is mapped
	// to the width parameter. If provided, the field widthUnit is required.
	Width float64 `json:"width"`
	// Given an ordered pair of spatial coordinates (p0, p1), defines a distance
	// extending into the LEFT half-plane relative to the direction of the vector that
	// maps p0 to p1. If cmShape is set to "1TRACK", this field is required to define
	// the width of the airspace track as measured from the left of the track segment
	// line. If provided, the field widthUnit is required.
	WidthLeft float64 `json:"widthLeft"`
	// Given an ordered pair of spatial coordinates (p0, p1), defines a distance
	// extending into the RIGHT half-plane relative to the direction of the vector that
	// maps p0 to p1. If cmShape is set to "1TRACK", this field is required to define
	// the width of the airspace track as measured from the right of the track segment
	// line. If provided, the field widthUnit is required.
	WidthRight float64 `json:"widthRight"`
	// Specifies the unit of length for which widths are given. Use M for meters, KM
	// for kilometers, or NM for nautical miles.
	WidthUnit string `json:"widthUnit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AirspaceControlPoint respjson.Field
		AirspaceTimePeriod   respjson.Field
		Bearing0             respjson.Field
		Bearing1             respjson.Field
		CmID                 respjson.Field
		CmShape              respjson.Field
		CmType               respjson.Field
		CntrlAuth            respjson.Field
		CntrlAuthFreqs       respjson.Field
		Coord0               respjson.Field
		Coord1               respjson.Field
		CorrWayPoints        respjson.Field
		EffVDim              respjson.Field
		FreeText             respjson.Field
		GenTextInd           respjson.Field
		GeoDatumAlt          respjson.Field
		Link16ID             respjson.Field
		OrbitAlignment       respjson.Field
		PolyCoord            respjson.Field
		RadMag0              respjson.Field
		RadMag1              respjson.Field
		RadMagUnit           respjson.Field
		TrackLeg             respjson.Field
		TransAltitude        respjson.Field
		Usage                respjson.Field
		Width                respjson.Field
		WidthLeft            respjson.Field
		WidthRight           respjson.Field
		WidthUnit            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMean) RawJSON() string {
	return r.JSON.raw
}
func (r *AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The controlPoint set describes any reference/controlling/rendezvous point for a
// given airspace control means.
type AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint struct {
	// The altitude of the control point.
	CtrlPtAltitude string `json:"ctrlPtAltitude"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format that represents the location of the control point.
	CtrlPtLocation string `json:"ctrlPtLocation"`
	// The name applied to the control point, used as a reference.
	CtrlPtName string `json:"ctrlPtName"`
	// One of possible control point type codes, such as CP, ER, OT, etc.
	CtrlPtType string `json:"ctrlPtType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CtrlPtAltitude respjson.Field
		CtrlPtLocation respjson.Field
		CtrlPtName     respjson.Field
		CtrlPtType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The timePeriods set describes the effective datetime for a given airspace
// control means.
type AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod struct {
	// Mandatory if timeMode is INTERVAL. Can be a numerical multiplier on an interval
	// frequency code, a stop time qualifier code such as AFTER, NET, UFN, etc, or a
	// datetime like string.
	IntDur []string `json:"intDur"`
	// Mandatory if timeMode is INTERVAL. Can be one of the interval frequency codes,
	// such as BIWEEKLY, DAILY, YEARLY, etc.
	IntFreq []string `json:"intFreq"`
	// The end time designating that the airspace control order is no longer active.
	// Can contain datetime information or a stop time qualifier code, such as AFTER,
	// NET, UFN, etc.
	TimeEnd string `json:"timeEnd"`
	// The airspace time code associated with the ACO. Can be DISCRETE, a fixed time
	// block, or INTERVAL, a repeating time block.
	TimeMode string `json:"timeMode"`
	// The start time designating that the airspace control order is active.
	TimeStart string `json:"timeStart"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntDur      respjson.Field
		IntFreq     respjson.Field
		TimeEnd     respjson.Field
		TimeMode    respjson.Field
		TimeStart   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod) RawJSON() string {
	return r.JSON.raw
}
func (r *AirspacecontrolorderAbridgedAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The airspaceControlReferences set provides both USMTF and non-USMTF references
// for this airspace control order.
type AirspacecontrolorderAbridgedAirspaceControlOrderReference struct {
	// The originator of this reference.
	RefOriginator string `json:"refOriginator"`
	// The reference serial number.
	RefSerialNum string `json:"refSerialNum"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of the document
	// being referenced.
	RefSiCs []string `json:"refSICs"`
	// Specifies an alphabetic serial number identifying a reference pertaining to this
	// message.
	RefSID string `json:"refSId"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this reference.
	RefSpecialNotation string `json:"refSpecialNotation"`
	// Timestamp of the referenced message, in ISO 8601 UTC format with millisecond
	// precision.
	RefTs time.Time `json:"refTs" format:"date-time"`
	// Specifies the type for this reference.
	RefType string `json:"refType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RefOriginator      respjson.Field
		RefSerialNum       respjson.Field
		RefSiCs            respjson.Field
		RefSID             respjson.Field
		RefSpecialNotation respjson.Field
		RefTs              respjson.Field
		RefType            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderAbridgedAirspaceControlOrderReference) RawJSON() string {
	return r.JSON.raw
}
func (r *AirspacecontrolorderAbridgedAirspaceControlOrderReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Beta Version Airspace Control Order: Contains airspace coordination information
// and instructions that have been issued by an airspace control authority.
type AirspacecontrolorderFull struct {
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
	DataMode AirspacecontrolorderFullDataMode `json:"dataMode,required"`
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExName string `json:"opExName,required"`
	// The identifier of the originator of this message.
	Originator string `json:"originator,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the effective time period of this airspace control order, in ISO
	// 8601 UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Free text information expressed in natural language.
	AcoComments string `json:"acoComments"`
	// The serial number of this airspace control order.
	AcoSerialNum string `json:"acoSerialNum"`
	// Mandatory nested segment to report multiple airspace control means statuses
	// within an ACOID.
	AirspaceControlMeansStatus []AirspacecontrolorderFullAirspaceControlMeansStatus `json:"airspaceControlMeansStatus"`
	// The airspaceControlReferences set provides both USMTF and non-USMTF references
	// for this airspace control order.
	AirspaceControlOrderReferences []AirspacecontrolorderFullAirspaceControlOrderReference `json:"airspaceControlOrderReferences"`
	// Name of the area of the command for which the ACO is valid.
	AreaOfValidity string `json:"areaOfValidity"`
	// Mandatory if classSource uses the "IORIG" designator. Must be a REASON FOR
	// CLASSIFICATION code.
	ClassReasons []string `json:"classReasons"`
	// Markings defining the source material or the original classification authority
	// for the ACO message.
	ClassSource string `json:"classSource"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Coded entries that provide justification for exemption from automatic
	// downgrading or declassification of the airspace control order.
	DeclassExemptionCodes []string `json:"declassExemptionCodes"`
	// Markings providing the literal guidance or date for downgrading or declassifying
	// the airspace control order.
	DowngradeInsDates []string `json:"downgradeInsDates"`
	// Specifies the geodetic datum by which the spatial coordinates of the controlled
	// airspace are calculated.
	GeoDatum string `json:"geoDatum"`
	// The month in which the message originated.
	Month string `json:"month"`
	// Supplementary name that can be used to further identify exercise nicknames, or
	// to provide the primary nickname of the option or the alternative of an
	// operational plan.
	OpExInfo string `json:"opExInfo"`
	// The secondary supplementary nickname of the option or the alternative of the
	// operational plan or order.
	OpExInfoAlt string `json:"opExInfoAlt"`
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
	// The qualifier which caveats the message status.
	Qualifier string `json:"qualifier"`
	// The serial number associated with the message qualifier.
	QualSn int64 `json:"qualSN"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The unique message identifier sequentially assigned by the originator.
	SerialNum string `json:"serialNum"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// A qualifier for the end of the effective time period of this airspace control
	// order, such as AFTER, ASOF, NLT, etc. Used with field stopTime to indicate a
	// relative time.
	StopQualifier string `json:"stopQualifier"`
	// The end of the effective time period of this airspace control order, in ISO 8601
	// UTC format with millisecond precision.
	StopTime time.Time `json:"stopTime" format:"date-time"`
	// Array of unique link 16 identifiers that will be assigned to a future airspace
	// control means.
	UndLnkTrks []string `json:"undLnkTrks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking          respjson.Field
		DataMode                       respjson.Field
		OpExName                       respjson.Field
		Originator                     respjson.Field
		Source                         respjson.Field
		StartTime                      respjson.Field
		ID                             respjson.Field
		AcoComments                    respjson.Field
		AcoSerialNum                   respjson.Field
		AirspaceControlMeansStatus     respjson.Field
		AirspaceControlOrderReferences respjson.Field
		AreaOfValidity                 respjson.Field
		ClassReasons                   respjson.Field
		ClassSource                    respjson.Field
		CreatedAt                      respjson.Field
		CreatedBy                      respjson.Field
		DeclassExemptionCodes          respjson.Field
		DowngradeInsDates              respjson.Field
		GeoDatum                       respjson.Field
		Month                          respjson.Field
		OpExInfo                       respjson.Field
		OpExInfoAlt                    respjson.Field
		Origin                         respjson.Field
		OrigNetwork                    respjson.Field
		PlanOrigNum                    respjson.Field
		Qualifier                      respjson.Field
		QualSn                         respjson.Field
		RawFileUri                     respjson.Field
		SerialNum                      respjson.Field
		SourceDl                       respjson.Field
		StopQualifier                  respjson.Field
		StopTime                       respjson.Field
		UndLnkTrks                     respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderFull) RawJSON() string { return r.JSON.raw }
func (r *AirspacecontrolorderFull) UnmarshalJSON(data []byte) error {
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
type AirspacecontrolorderFullDataMode string

const (
	AirspacecontrolorderFullDataModeReal      AirspacecontrolorderFullDataMode = "REAL"
	AirspacecontrolorderFullDataModeTest      AirspacecontrolorderFullDataMode = "TEST"
	AirspacecontrolorderFullDataModeSimulated AirspacecontrolorderFullDataMode = "SIMULATED"
	AirspacecontrolorderFullDataModeExercise  AirspacecontrolorderFullDataMode = "EXERCISE"
)

// Mandatory nested segment to report multiple airspace control means statuses
// within an ACOID.
type AirspacecontrolorderFullAirspaceControlMeansStatus struct {
	// A conditional nested segment to report multiple airspace control means within a
	// particular airspace control means status.
	AirspaceControlMeans []AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMean `json:"airspaceControlMeans"`
	// Status of Airspace Control Means. Must be ADD, CHANGE, or DELETE.
	CmStat string `json:"cmStat"`
	// Airspace control means name or designator. Mandatory if acmStat equals "DELETE,"
	// otherwise this field is prohibited.
	CmStatID []string `json:"cmStatId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AirspaceControlMeans respjson.Field
		CmStat               respjson.Field
		CmStatID             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderFullAirspaceControlMeansStatus) RawJSON() string { return r.JSON.raw }
func (r *AirspacecontrolorderFullAirspaceControlMeansStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A conditional nested segment to report multiple airspace control means within a
// particular airspace control means status.
type AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMean struct {
	// The controlPoint set describes any reference/controlling/rendezvous point for a
	// given airspace control means.
	AirspaceControlPoint []AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint `json:"airspaceControlPoint"`
	// The timePeriods set describes the effective datetime for a given airspace
	// control means.
	AirspaceTimePeriod []AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod `json:"airspaceTimePeriod"`
	// A bearing measured from true North, in angular degrees. If cmShape is set to
	// "POLYARC" or "RADARC", this field is required and is mapped to the "beginning"
	// radial bearing parameter.
	Bearing0 float64 `json:"bearing0"`
	// A bearing measured from true North, in angular degrees. If cmShape is set to
	// "POLYARC" or "RADARC", this field is required and is mapped to the "ending"
	// radial bearing parameter.
	Bearing1 float64 `json:"bearing1"`
	// Airspace control means name or designator.
	CmID string `json:"cmId"`
	// Designates the geometric type that defines the airspace shape. One of CIRCLE,
	// CORRIDOR, LINE, ORBIT, etc.
	//
	// Any of "POLYARC", "1TRACK", "POLYGON", "CIRCLE", "CORRIDOR", "APOINT", "AORBIT",
	// "GEOLINE".
	CmShape string `json:"cmShape"`
	// The code for the type of airspace control means.
	CmType string `json:"cmType"`
	// The commander responsible within a specified geographical area for the airspace
	// control operation assigned to him.
	CntrlAuth string `json:"cntrlAuth"`
	// The frequency for the airspace control authority. Can specify HZ, KHZ, MHZ, GHZ
	// or a DESIG frequency designator code.
	CntrlAuthFreqs []string `json:"cntrlAuthFreqs"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format. The fields coord0 and coord1 should be used in the specification of any
	// airspace control shape that requires exactly one (1) or two (2) reference points
	// for construction. For shapes requiring one reference point, for instance, when
	// cmShape is set to "APOINT", this field is required and singularly defines the
	// shape. Similarly, this field is required to define the center point of a
	// "CIRCLE" shape, or the "origin of bearing" for arcs.
	Coord0 string `json:"coord0"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format. The fields coord0 and coord1 should be used in the specification of any
	// airspace control shape that requires exactly one (1) or two (2) reference points
	// for construction. For shapes requiring one reference point, for instance, when
	// cmShape is set to "APOINT", this field is required and singularly defines the
	// shape. Similarly, this field is required to define the center point of a
	// "CIRCLE" shape, or the "origin of bearing" for arcs.
	Coord1 string `json:"coord1"`
	// An array of at least two alphanumeric symbols used to serially identify the
	// corridor waypoints. If cmShape is set to "CORRIDOR", one of either corrWayPoints
	// or polyCoord is required to specify the centerline of the corridor path.
	CorrWayPoints []string `json:"corrWayPoints"`
	// Description of the airspace vertical dimension.
	EffVDim string `json:"effVDim"`
	// General informat detailing the transit instruction for the airspace control
	// means.
	FreeText string `json:"freeText"`
	// Used to provide transit instructions for the airspace control means.
	GenTextInd string `json:"genTextInd"`
	// Specifies the geodetic datum by which the spatial coordinates of the controlled
	// airspace are calculated, if different from the top level ACO datum.
	GeoDatumAlt string `json:"geoDatumAlt"`
	// Unique Link 16 identifier assigned to the airspace control means.
	Link16ID string `json:"link16Id"`
	// Orbit alignment look-up code. Can be C=Center, L=Left, R=Right.
	OrbitAlignment string `json:"orbitAlignment"`
	// A set of geospatial coordinates specified in DMS (Degrees, Minutes, Seconds)
	// format which determine the vertices of a one or two dimensional geospatial
	// shape. When cmShape is set to "POLYARC" or "POLYGON", this field is required as
	// applied in the construction of the area boundary. If cmShape is set to
	// "CORRIDOR" or "GEOLINE", this field is required and can be interpreted as an
	// ordered set of points along a path in space.
	PolyCoord []string `json:"polyCoord"`
	// A distance that represents a radial magnitude. If cmShape is set to "CIRCLE" or
	// "POLYARC", one of either fields radMag0 or radMag1 is required. If cmShape is
	// set to "RADARC", this field is required and maps to the "inner" radial magnitude
	// arc limit. If provided, the field radMagUnit is required.
	RadMag0 float64 `json:"radMag0"`
	// A distance that represents a radial magnitude. If cmShape is set to "CIRCLE" or
	// "POLYARC", one of either fields radMag0 or radMag1 is required. If cmShape is
	// set to "RADARC", this field is required and maps to the "outer" radial magnitude
	// arc limit. If provided, the field radMagUnit is required.
	RadMag1 float64 `json:"radMag1"`
	// Specifies the unit of length in which radial magnitudes are given. Use M for
	// meters, KM for kilometers, or NM for nautical miles.
	RadMagUnit string `json:"radMagUnit"`
	// Index of a segment in an airtrack, which is defined by an ordered set of points.
	TrackLeg int64 `json:"trackLeg"`
	// The altitude at or below which the vertical position of an aircraft is
	// controlled by reference to true altitude.
	TransAltitude string `json:"transAltitude"`
	// Designates the means by which a defined airspace control means is to be used.
	Usage string `json:"usage"`
	// Used to describe the "side to side" distance of a target, object or area. If
	// cmShape is set to "CORRIDOR" or "AORBIT", this field is required and is mapped
	// to the width parameter. If provided, the field widthUnit is required.
	Width float64 `json:"width"`
	// Given an ordered pair of spatial coordinates (p0, p1), defines a distance
	// extending into the LEFT half-plane relative to the direction of the vector that
	// maps p0 to p1. If cmShape is set to "1TRACK", this field is required to define
	// the width of the airspace track as measured from the left of the track segment
	// line. If provided, the field widthUnit is required.
	WidthLeft float64 `json:"widthLeft"`
	// Given an ordered pair of spatial coordinates (p0, p1), defines a distance
	// extending into the RIGHT half-plane relative to the direction of the vector that
	// maps p0 to p1. If cmShape is set to "1TRACK", this field is required to define
	// the width of the airspace track as measured from the right of the track segment
	// line. If provided, the field widthUnit is required.
	WidthRight float64 `json:"widthRight"`
	// Specifies the unit of length for which widths are given. Use M for meters, KM
	// for kilometers, or NM for nautical miles.
	WidthUnit string `json:"widthUnit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AirspaceControlPoint respjson.Field
		AirspaceTimePeriod   respjson.Field
		Bearing0             respjson.Field
		Bearing1             respjson.Field
		CmID                 respjson.Field
		CmShape              respjson.Field
		CmType               respjson.Field
		CntrlAuth            respjson.Field
		CntrlAuthFreqs       respjson.Field
		Coord0               respjson.Field
		Coord1               respjson.Field
		CorrWayPoints        respjson.Field
		EffVDim              respjson.Field
		FreeText             respjson.Field
		GenTextInd           respjson.Field
		GeoDatumAlt          respjson.Field
		Link16ID             respjson.Field
		OrbitAlignment       respjson.Field
		PolyCoord            respjson.Field
		RadMag0              respjson.Field
		RadMag1              respjson.Field
		RadMagUnit           respjson.Field
		TrackLeg             respjson.Field
		TransAltitude        respjson.Field
		Usage                respjson.Field
		Width                respjson.Field
		WidthLeft            respjson.Field
		WidthRight           respjson.Field
		WidthUnit            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMean) RawJSON() string {
	return r.JSON.raw
}
func (r *AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The controlPoint set describes any reference/controlling/rendezvous point for a
// given airspace control means.
type AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint struct {
	// The altitude of the control point.
	CtrlPtAltitude string `json:"ctrlPtAltitude"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format that represents the location of the control point.
	CtrlPtLocation string `json:"ctrlPtLocation"`
	// The name applied to the control point, used as a reference.
	CtrlPtName string `json:"ctrlPtName"`
	// One of possible control point type codes, such as CP, ER, OT, etc.
	CtrlPtType string `json:"ctrlPtType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CtrlPtAltitude respjson.Field
		CtrlPtLocation respjson.Field
		CtrlPtName     respjson.Field
		CtrlPtType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The timePeriods set describes the effective datetime for a given airspace
// control means.
type AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod struct {
	// Mandatory if timeMode is INTERVAL. Can be a numerical multiplier on an interval
	// frequency code, a stop time qualifier code such as AFTER, NET, UFN, etc, or a
	// datetime like string.
	IntDur []string `json:"intDur"`
	// Mandatory if timeMode is INTERVAL. Can be one of the interval frequency codes,
	// such as BIWEEKLY, DAILY, YEARLY, etc.
	IntFreq []string `json:"intFreq"`
	// The end time designating that the airspace control order is no longer active.
	// Can contain datetime information or a stop time qualifier code, such as AFTER,
	// NET, UFN, etc.
	TimeEnd string `json:"timeEnd"`
	// The airspace time code associated with the ACO. Can be DISCRETE, a fixed time
	// block, or INTERVAL, a repeating time block.
	TimeMode string `json:"timeMode"`
	// The start time designating that the airspace control order is active.
	TimeStart string `json:"timeStart"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntDur      respjson.Field
		IntFreq     respjson.Field
		TimeEnd     respjson.Field
		TimeMode    respjson.Field
		TimeStart   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod) RawJSON() string {
	return r.JSON.raw
}
func (r *AirspacecontrolorderFullAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The airspaceControlReferences set provides both USMTF and non-USMTF references
// for this airspace control order.
type AirspacecontrolorderFullAirspaceControlOrderReference struct {
	// The originator of this reference.
	RefOriginator string `json:"refOriginator"`
	// The reference serial number.
	RefSerialNum string `json:"refSerialNum"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of the document
	// being referenced.
	RefSiCs []string `json:"refSICs"`
	// Specifies an alphabetic serial number identifying a reference pertaining to this
	// message.
	RefSID string `json:"refSId"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this reference.
	RefSpecialNotation string `json:"refSpecialNotation"`
	// Timestamp of the referenced message, in ISO 8601 UTC format with millisecond
	// precision.
	RefTs time.Time `json:"refTs" format:"date-time"`
	// Specifies the type for this reference.
	RefType string `json:"refType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RefOriginator      respjson.Field
		RefSerialNum       respjson.Field
		RefSiCs            respjson.Field
		RefSID             respjson.Field
		RefSpecialNotation respjson.Field
		RefTs              respjson.Field
		RefType            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirspacecontrolorderFullAirspaceControlOrderReference) RawJSON() string { return r.JSON.raw }
func (r *AirspacecontrolorderFullAirspaceControlOrderReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AirspaceControlOrderNewParams struct {
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
	DataMode AirspaceControlOrderNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Specifies the unique operation or exercise name, nickname, or codeword assigned
	// to a joint exercise or operation plan.
	OpExName string `json:"opExName,required"`
	// The identifier of the originator of this message.
	Originator string `json:"originator,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the effective time period of this airspace control order, in ISO
	// 8601 UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Free text information expressed in natural language.
	AcoComments param.Opt[string] `json:"acoComments,omitzero"`
	// The serial number of this airspace control order.
	AcoSerialNum param.Opt[string] `json:"acoSerialNum,omitzero"`
	// Name of the area of the command for which the ACO is valid.
	AreaOfValidity param.Opt[string] `json:"areaOfValidity,omitzero"`
	// Markings defining the source material or the original classification authority
	// for the ACO message.
	ClassSource param.Opt[string] `json:"classSource,omitzero"`
	// Specifies the geodetic datum by which the spatial coordinates of the controlled
	// airspace are calculated.
	GeoDatum param.Opt[string] `json:"geoDatum,omitzero"`
	// The month in which the message originated.
	Month param.Opt[string] `json:"month,omitzero"`
	// Supplementary name that can be used to further identify exercise nicknames, or
	// to provide the primary nickname of the option or the alternative of an
	// operational plan.
	OpExInfo param.Opt[string] `json:"opExInfo,omitzero"`
	// The secondary supplementary nickname of the option or the alternative of the
	// operational plan or order.
	OpExInfoAlt param.Opt[string] `json:"opExInfoAlt,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The official identifier of the military establishment responsible for the
	// operation plan and the identification number assigned to this plan.
	PlanOrigNum param.Opt[string] `json:"planOrigNum,omitzero"`
	// The qualifier which caveats the message status.
	Qualifier param.Opt[string] `json:"qualifier,omitzero"`
	// The serial number associated with the message qualifier.
	QualSn param.Opt[int64] `json:"qualSN,omitzero"`
	// The unique message identifier sequentially assigned by the originator.
	SerialNum param.Opt[string] `json:"serialNum,omitzero"`
	// A qualifier for the end of the effective time period of this airspace control
	// order, such as AFTER, ASOF, NLT, etc. Used with field stopTime to indicate a
	// relative time.
	StopQualifier param.Opt[string] `json:"stopQualifier,omitzero"`
	// The end of the effective time period of this airspace control order, in ISO 8601
	// UTC format with millisecond precision.
	StopTime param.Opt[time.Time] `json:"stopTime,omitzero" format:"date-time"`
	// Mandatory nested segment to report multiple airspace control means statuses
	// within an ACOID.
	AirspaceControlMeansStatus []AirspaceControlOrderNewParamsAirspaceControlMeansStatus `json:"airspaceControlMeansStatus,omitzero"`
	// The airspaceControlReferences set provides both USMTF and non-USMTF references
	// for this airspace control order.
	AirspaceControlOrderReferences []AirspaceControlOrderNewParamsAirspaceControlOrderReference `json:"airspaceControlOrderReferences,omitzero"`
	// Mandatory if classSource uses the "IORIG" designator. Must be a REASON FOR
	// CLASSIFICATION code.
	ClassReasons []string `json:"classReasons,omitzero"`
	// Coded entries that provide justification for exemption from automatic
	// downgrading or declassification of the airspace control order.
	DeclassExemptionCodes []string `json:"declassExemptionCodes,omitzero"`
	// Markings providing the literal guidance or date for downgrading or declassifying
	// the airspace control order.
	DowngradeInsDates []string `json:"downgradeInsDates,omitzero"`
	// Array of unique link 16 identifiers that will be assigned to a future airspace
	// control means.
	UndLnkTrks []string `json:"undLnkTrks,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewParams
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
type AirspaceControlOrderNewParamsDataMode string

const (
	AirspaceControlOrderNewParamsDataModeReal      AirspaceControlOrderNewParamsDataMode = "REAL"
	AirspaceControlOrderNewParamsDataModeTest      AirspaceControlOrderNewParamsDataMode = "TEST"
	AirspaceControlOrderNewParamsDataModeSimulated AirspaceControlOrderNewParamsDataMode = "SIMULATED"
	AirspaceControlOrderNewParamsDataModeExercise  AirspaceControlOrderNewParamsDataMode = "EXERCISE"
)

// Mandatory nested segment to report multiple airspace control means statuses
// within an ACOID.
type AirspaceControlOrderNewParamsAirspaceControlMeansStatus struct {
	// Status of Airspace Control Means. Must be ADD, CHANGE, or DELETE.
	CmStat param.Opt[string] `json:"cmStat,omitzero"`
	// A conditional nested segment to report multiple airspace control means within a
	// particular airspace control means status.
	AirspaceControlMeans []AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMean `json:"airspaceControlMeans,omitzero"`
	// Airspace control means name or designator. Mandatory if acmStat equals "DELETE,"
	// otherwise this field is prohibited.
	CmStatID []string `json:"cmStatId,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewParamsAirspaceControlMeansStatus) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewParamsAirspaceControlMeansStatus
	return param.MarshalObject(r, (*shadow)(&r))
}

// A conditional nested segment to report multiple airspace control means within a
// particular airspace control means status.
type AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMean struct {
	// A bearing measured from true North, in angular degrees. If cmShape is set to
	// "POLYARC" or "RADARC", this field is required and is mapped to the "beginning"
	// radial bearing parameter.
	Bearing0 param.Opt[float64] `json:"bearing0,omitzero"`
	// A bearing measured from true North, in angular degrees. If cmShape is set to
	// "POLYARC" or "RADARC", this field is required and is mapped to the "ending"
	// radial bearing parameter.
	Bearing1 param.Opt[float64] `json:"bearing1,omitzero"`
	// Airspace control means name or designator.
	CmID param.Opt[string] `json:"cmId,omitzero"`
	// The code for the type of airspace control means.
	CmType param.Opt[string] `json:"cmType,omitzero"`
	// The commander responsible within a specified geographical area for the airspace
	// control operation assigned to him.
	CntrlAuth param.Opt[string] `json:"cntrlAuth,omitzero"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format. The fields coord0 and coord1 should be used in the specification of any
	// airspace control shape that requires exactly one (1) or two (2) reference points
	// for construction. For shapes requiring one reference point, for instance, when
	// cmShape is set to "APOINT", this field is required and singularly defines the
	// shape. Similarly, this field is required to define the center point of a
	// "CIRCLE" shape, or the "origin of bearing" for arcs.
	Coord0 param.Opt[string] `json:"coord0,omitzero"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format. The fields coord0 and coord1 should be used in the specification of any
	// airspace control shape that requires exactly one (1) or two (2) reference points
	// for construction. For shapes requiring one reference point, for instance, when
	// cmShape is set to "APOINT", this field is required and singularly defines the
	// shape. Similarly, this field is required to define the center point of a
	// "CIRCLE" shape, or the "origin of bearing" for arcs.
	Coord1 param.Opt[string] `json:"coord1,omitzero"`
	// Description of the airspace vertical dimension.
	EffVDim param.Opt[string] `json:"effVDim,omitzero"`
	// General informat detailing the transit instruction for the airspace control
	// means.
	FreeText param.Opt[string] `json:"freeText,omitzero"`
	// Used to provide transit instructions for the airspace control means.
	GenTextInd param.Opt[string] `json:"genTextInd,omitzero"`
	// Specifies the geodetic datum by which the spatial coordinates of the controlled
	// airspace are calculated, if different from the top level ACO datum.
	GeoDatumAlt param.Opt[string] `json:"geoDatumAlt,omitzero"`
	// Unique Link 16 identifier assigned to the airspace control means.
	Link16ID param.Opt[string] `json:"link16Id,omitzero"`
	// Orbit alignment look-up code. Can be C=Center, L=Left, R=Right.
	OrbitAlignment param.Opt[string] `json:"orbitAlignment,omitzero"`
	// A distance that represents a radial magnitude. If cmShape is set to "CIRCLE" or
	// "POLYARC", one of either fields radMag0 or radMag1 is required. If cmShape is
	// set to "RADARC", this field is required and maps to the "inner" radial magnitude
	// arc limit. If provided, the field radMagUnit is required.
	RadMag0 param.Opt[float64] `json:"radMag0,omitzero"`
	// A distance that represents a radial magnitude. If cmShape is set to "CIRCLE" or
	// "POLYARC", one of either fields radMag0 or radMag1 is required. If cmShape is
	// set to "RADARC", this field is required and maps to the "outer" radial magnitude
	// arc limit. If provided, the field radMagUnit is required.
	RadMag1 param.Opt[float64] `json:"radMag1,omitzero"`
	// Specifies the unit of length in which radial magnitudes are given. Use M for
	// meters, KM for kilometers, or NM for nautical miles.
	RadMagUnit param.Opt[string] `json:"radMagUnit,omitzero"`
	// Index of a segment in an airtrack, which is defined by an ordered set of points.
	TrackLeg param.Opt[int64] `json:"trackLeg,omitzero"`
	// The altitude at or below which the vertical position of an aircraft is
	// controlled by reference to true altitude.
	TransAltitude param.Opt[string] `json:"transAltitude,omitzero"`
	// Designates the means by which a defined airspace control means is to be used.
	Usage param.Opt[string] `json:"usage,omitzero"`
	// Used to describe the "side to side" distance of a target, object or area. If
	// cmShape is set to "CORRIDOR" or "AORBIT", this field is required and is mapped
	// to the width parameter. If provided, the field widthUnit is required.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Given an ordered pair of spatial coordinates (p0, p1), defines a distance
	// extending into the LEFT half-plane relative to the direction of the vector that
	// maps p0 to p1. If cmShape is set to "1TRACK", this field is required to define
	// the width of the airspace track as measured from the left of the track segment
	// line. If provided, the field widthUnit is required.
	WidthLeft param.Opt[float64] `json:"widthLeft,omitzero"`
	// Given an ordered pair of spatial coordinates (p0, p1), defines a distance
	// extending into the RIGHT half-plane relative to the direction of the vector that
	// maps p0 to p1. If cmShape is set to "1TRACK", this field is required to define
	// the width of the airspace track as measured from the right of the track segment
	// line. If provided, the field widthUnit is required.
	WidthRight param.Opt[float64] `json:"widthRight,omitzero"`
	// Specifies the unit of length for which widths are given. Use M for meters, KM
	// for kilometers, or NM for nautical miles.
	WidthUnit param.Opt[string] `json:"widthUnit,omitzero"`
	// The controlPoint set describes any reference/controlling/rendezvous point for a
	// given airspace control means.
	AirspaceControlPoint []AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint `json:"airspaceControlPoint,omitzero"`
	// The timePeriods set describes the effective datetime for a given airspace
	// control means.
	AirspaceTimePeriod []AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod `json:"airspaceTimePeriod,omitzero"`
	// Designates the geometric type that defines the airspace shape. One of CIRCLE,
	// CORRIDOR, LINE, ORBIT, etc.
	//
	// Any of "POLYARC", "1TRACK", "POLYGON", "CIRCLE", "CORRIDOR", "APOINT", "AORBIT",
	// "GEOLINE".
	CmShape string `json:"cmShape,omitzero"`
	// The frequency for the airspace control authority. Can specify HZ, KHZ, MHZ, GHZ
	// or a DESIG frequency designator code.
	CntrlAuthFreqs []string `json:"cntrlAuthFreqs,omitzero"`
	// An array of at least two alphanumeric symbols used to serially identify the
	// corridor waypoints. If cmShape is set to "CORRIDOR", one of either corrWayPoints
	// or polyCoord is required to specify the centerline of the corridor path.
	CorrWayPoints []string `json:"corrWayPoints,omitzero"`
	// A set of geospatial coordinates specified in DMS (Degrees, Minutes, Seconds)
	// format which determine the vertices of a one or two dimensional geospatial
	// shape. When cmShape is set to "POLYARC" or "POLYGON", this field is required as
	// applied in the construction of the area boundary. If cmShape is set to
	// "CORRIDOR" or "GEOLINE", this field is required and can be interpreted as an
	// ordered set of points along a path in space.
	PolyCoord []string `json:"polyCoord,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMean) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMean
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMean](
		"CmShape", false, "POLYARC", "1TRACK", "POLYGON", "CIRCLE", "CORRIDOR", "APOINT", "AORBIT", "GEOLINE",
	)
}

// The controlPoint set describes any reference/controlling/rendezvous point for a
// given airspace control means.
type AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint struct {
	// The altitude of the control point.
	CtrlPtAltitude param.Opt[string] `json:"ctrlPtAltitude,omitzero"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format that represents the location of the control point.
	CtrlPtLocation param.Opt[string] `json:"ctrlPtLocation,omitzero"`
	// The name applied to the control point, used as a reference.
	CtrlPtName param.Opt[string] `json:"ctrlPtName,omitzero"`
	// One of possible control point type codes, such as CP, ER, OT, etc.
	CtrlPtType param.Opt[string] `json:"ctrlPtType,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint
	return param.MarshalObject(r, (*shadow)(&r))
}

// The timePeriods set describes the effective datetime for a given airspace
// control means.
type AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod struct {
	// The end time designating that the airspace control order is no longer active.
	// Can contain datetime information or a stop time qualifier code, such as AFTER,
	// NET, UFN, etc.
	TimeEnd param.Opt[string] `json:"timeEnd,omitzero"`
	// The airspace time code associated with the ACO. Can be DISCRETE, a fixed time
	// block, or INTERVAL, a repeating time block.
	TimeMode param.Opt[string] `json:"timeMode,omitzero"`
	// The start time designating that the airspace control order is active.
	TimeStart param.Opt[string] `json:"timeStart,omitzero"`
	// Mandatory if timeMode is INTERVAL. Can be a numerical multiplier on an interval
	// frequency code, a stop time qualifier code such as AFTER, NET, UFN, etc, or a
	// datetime like string.
	IntDur []string `json:"intDur,omitzero"`
	// Mandatory if timeMode is INTERVAL. Can be one of the interval frequency codes,
	// such as BIWEEKLY, DAILY, YEARLY, etc.
	IntFreq []string `json:"intFreq,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod
	return param.MarshalObject(r, (*shadow)(&r))
}

// The airspaceControlReferences set provides both USMTF and non-USMTF references
// for this airspace control order.
type AirspaceControlOrderNewParamsAirspaceControlOrderReference struct {
	// The originator of this reference.
	RefOriginator param.Opt[string] `json:"refOriginator,omitzero"`
	// The reference serial number.
	RefSerialNum param.Opt[string] `json:"refSerialNum,omitzero"`
	// Specifies an alphabetic serial number identifying a reference pertaining to this
	// message.
	RefSID param.Opt[string] `json:"refSId,omitzero"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this reference.
	RefSpecialNotation param.Opt[string] `json:"refSpecialNotation,omitzero"`
	// Timestamp of the referenced message, in ISO 8601 UTC format with millisecond
	// precision.
	RefTs param.Opt[time.Time] `json:"refTs,omitzero" format:"date-time"`
	// Specifies the type for this reference.
	RefType param.Opt[string] `json:"refType,omitzero"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of the document
	// being referenced.
	RefSiCs []string `json:"refSICs,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewParamsAirspaceControlOrderReference) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewParamsAirspaceControlOrderReference
	return param.MarshalObject(r, (*shadow)(&r))
}

type AirspaceControlOrderGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirspaceControlOrderGetParams]'s query parameters as
// `url.Values`.
func (r AirspaceControlOrderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirspaceControlOrderListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirspaceControlOrderListParams]'s query parameters as
// `url.Values`.
func (r AirspaceControlOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirspaceControlOrderCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirspaceControlOrderCountParams]'s query parameters as
// `url.Values`.
func (r AirspaceControlOrderCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirspaceControlOrderNewBulkParams struct {
	Body []AirspaceControlOrderNewBulkParamsBody
	paramObj
}

func (r AirspaceControlOrderNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Beta Version Airspace Control Order: Contains airspace coordination information
// and instructions that have been issued by an airspace control authority.
//
// The properties ClassificationMarking, DataMode, OpExName, Originator, Source,
// StartTime are required.
type AirspaceControlOrderNewBulkParamsBody struct {
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
	OpExName string `json:"opExName,required"`
	// The identifier of the originator of this message.
	Originator string `json:"originator,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The start of the effective time period of this airspace control order, in ISO
	// 8601 UTC format with millisecond precision.
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Free text information expressed in natural language.
	AcoComments param.Opt[string] `json:"acoComments,omitzero"`
	// The serial number of this airspace control order.
	AcoSerialNum param.Opt[string] `json:"acoSerialNum,omitzero"`
	// Name of the area of the command for which the ACO is valid.
	AreaOfValidity param.Opt[string] `json:"areaOfValidity,omitzero"`
	// Markings defining the source material or the original classification authority
	// for the ACO message.
	ClassSource param.Opt[string] `json:"classSource,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Specifies the geodetic datum by which the spatial coordinates of the controlled
	// airspace are calculated.
	GeoDatum param.Opt[string] `json:"geoDatum,omitzero"`
	// The month in which the message originated.
	Month param.Opt[string] `json:"month,omitzero"`
	// Supplementary name that can be used to further identify exercise nicknames, or
	// to provide the primary nickname of the option or the alternative of an
	// operational plan.
	OpExInfo param.Opt[string] `json:"opExInfo,omitzero"`
	// The secondary supplementary nickname of the option or the alternative of the
	// operational plan or order.
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
	// The qualifier which caveats the message status.
	Qualifier param.Opt[string] `json:"qualifier,omitzero"`
	// The serial number associated with the message qualifier.
	QualSn param.Opt[int64] `json:"qualSN,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The unique message identifier sequentially assigned by the originator.
	SerialNum param.Opt[string] `json:"serialNum,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// A qualifier for the end of the effective time period of this airspace control
	// order, such as AFTER, ASOF, NLT, etc. Used with field stopTime to indicate a
	// relative time.
	StopQualifier param.Opt[string] `json:"stopQualifier,omitzero"`
	// The end of the effective time period of this airspace control order, in ISO 8601
	// UTC format with millisecond precision.
	StopTime param.Opt[time.Time] `json:"stopTime,omitzero" format:"date-time"`
	// Mandatory nested segment to report multiple airspace control means statuses
	// within an ACOID.
	AirspaceControlMeansStatus []AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatus `json:"airspaceControlMeansStatus,omitzero"`
	// The airspaceControlReferences set provides both USMTF and non-USMTF references
	// for this airspace control order.
	AirspaceControlOrderReferences []AirspaceControlOrderNewBulkParamsBodyAirspaceControlOrderReference `json:"airspaceControlOrderReferences,omitzero"`
	// Mandatory if classSource uses the "IORIG" designator. Must be a REASON FOR
	// CLASSIFICATION code.
	ClassReasons []string `json:"classReasons,omitzero"`
	// Coded entries that provide justification for exemption from automatic
	// downgrading or declassification of the airspace control order.
	DeclassExemptionCodes []string `json:"declassExemptionCodes,omitzero"`
	// Markings providing the literal guidance or date for downgrading or declassifying
	// the airspace control order.
	DowngradeInsDates []string `json:"downgradeInsDates,omitzero"`
	// Array of unique link 16 identifiers that will be assigned to a future airspace
	// control means.
	UndLnkTrks []string `json:"undLnkTrks,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[AirspaceControlOrderNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Mandatory nested segment to report multiple airspace control means statuses
// within an ACOID.
type AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatus struct {
	// Status of Airspace Control Means. Must be ADD, CHANGE, or DELETE.
	CmStat param.Opt[string] `json:"cmStat,omitzero"`
	// A conditional nested segment to report multiple airspace control means within a
	// particular airspace control means status.
	AirspaceControlMeans []AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMean `json:"airspaceControlMeans,omitzero"`
	// Airspace control means name or designator. Mandatory if acmStat equals "DELETE,"
	// otherwise this field is prohibited.
	CmStatID []string `json:"cmStatId,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatus) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatus
	return param.MarshalObject(r, (*shadow)(&r))
}

// A conditional nested segment to report multiple airspace control means within a
// particular airspace control means status.
type AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMean struct {
	// A bearing measured from true North, in angular degrees. If cmShape is set to
	// "POLYARC" or "RADARC", this field is required and is mapped to the "beginning"
	// radial bearing parameter.
	Bearing0 param.Opt[float64] `json:"bearing0,omitzero"`
	// A bearing measured from true North, in angular degrees. If cmShape is set to
	// "POLYARC" or "RADARC", this field is required and is mapped to the "ending"
	// radial bearing parameter.
	Bearing1 param.Opt[float64] `json:"bearing1,omitzero"`
	// Airspace control means name or designator.
	CmID param.Opt[string] `json:"cmId,omitzero"`
	// The code for the type of airspace control means.
	CmType param.Opt[string] `json:"cmType,omitzero"`
	// The commander responsible within a specified geographical area for the airspace
	// control operation assigned to him.
	CntrlAuth param.Opt[string] `json:"cntrlAuth,omitzero"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format. The fields coord0 and coord1 should be used in the specification of any
	// airspace control shape that requires exactly one (1) or two (2) reference points
	// for construction. For shapes requiring one reference point, for instance, when
	// cmShape is set to "APOINT", this field is required and singularly defines the
	// shape. Similarly, this field is required to define the center point of a
	// "CIRCLE" shape, or the "origin of bearing" for arcs.
	Coord0 param.Opt[string] `json:"coord0,omitzero"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format. The fields coord0 and coord1 should be used in the specification of any
	// airspace control shape that requires exactly one (1) or two (2) reference points
	// for construction. For shapes requiring one reference point, for instance, when
	// cmShape is set to "APOINT", this field is required and singularly defines the
	// shape. Similarly, this field is required to define the center point of a
	// "CIRCLE" shape, or the "origin of bearing" for arcs.
	Coord1 param.Opt[string] `json:"coord1,omitzero"`
	// Description of the airspace vertical dimension.
	EffVDim param.Opt[string] `json:"effVDim,omitzero"`
	// General informat detailing the transit instruction for the airspace control
	// means.
	FreeText param.Opt[string] `json:"freeText,omitzero"`
	// Used to provide transit instructions for the airspace control means.
	GenTextInd param.Opt[string] `json:"genTextInd,omitzero"`
	// Specifies the geodetic datum by which the spatial coordinates of the controlled
	// airspace are calculated, if different from the top level ACO datum.
	GeoDatumAlt param.Opt[string] `json:"geoDatumAlt,omitzero"`
	// Unique Link 16 identifier assigned to the airspace control means.
	Link16ID param.Opt[string] `json:"link16Id,omitzero"`
	// Orbit alignment look-up code. Can be C=Center, L=Left, R=Right.
	OrbitAlignment param.Opt[string] `json:"orbitAlignment,omitzero"`
	// A distance that represents a radial magnitude. If cmShape is set to "CIRCLE" or
	// "POLYARC", one of either fields radMag0 or radMag1 is required. If cmShape is
	// set to "RADARC", this field is required and maps to the "inner" radial magnitude
	// arc limit. If provided, the field radMagUnit is required.
	RadMag0 param.Opt[float64] `json:"radMag0,omitzero"`
	// A distance that represents a radial magnitude. If cmShape is set to "CIRCLE" or
	// "POLYARC", one of either fields radMag0 or radMag1 is required. If cmShape is
	// set to "RADARC", this field is required and maps to the "outer" radial magnitude
	// arc limit. If provided, the field radMagUnit is required.
	RadMag1 param.Opt[float64] `json:"radMag1,omitzero"`
	// Specifies the unit of length in which radial magnitudes are given. Use M for
	// meters, KM for kilometers, or NM for nautical miles.
	RadMagUnit param.Opt[string] `json:"radMagUnit,omitzero"`
	// Index of a segment in an airtrack, which is defined by an ordered set of points.
	TrackLeg param.Opt[int64] `json:"trackLeg,omitzero"`
	// The altitude at or below which the vertical position of an aircraft is
	// controlled by reference to true altitude.
	TransAltitude param.Opt[string] `json:"transAltitude,omitzero"`
	// Designates the means by which a defined airspace control means is to be used.
	Usage param.Opt[string] `json:"usage,omitzero"`
	// Used to describe the "side to side" distance of a target, object or area. If
	// cmShape is set to "CORRIDOR" or "AORBIT", this field is required and is mapped
	// to the width parameter. If provided, the field widthUnit is required.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Given an ordered pair of spatial coordinates (p0, p1), defines a distance
	// extending into the LEFT half-plane relative to the direction of the vector that
	// maps p0 to p1. If cmShape is set to "1TRACK", this field is required to define
	// the width of the airspace track as measured from the left of the track segment
	// line. If provided, the field widthUnit is required.
	WidthLeft param.Opt[float64] `json:"widthLeft,omitzero"`
	// Given an ordered pair of spatial coordinates (p0, p1), defines a distance
	// extending into the RIGHT half-plane relative to the direction of the vector that
	// maps p0 to p1. If cmShape is set to "1TRACK", this field is required to define
	// the width of the airspace track as measured from the right of the track segment
	// line. If provided, the field widthUnit is required.
	WidthRight param.Opt[float64] `json:"widthRight,omitzero"`
	// Specifies the unit of length for which widths are given. Use M for meters, KM
	// for kilometers, or NM for nautical miles.
	WidthUnit param.Opt[string] `json:"widthUnit,omitzero"`
	// The controlPoint set describes any reference/controlling/rendezvous point for a
	// given airspace control means.
	AirspaceControlPoint []AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint `json:"airspaceControlPoint,omitzero"`
	// The timePeriods set describes the effective datetime for a given airspace
	// control means.
	AirspaceTimePeriod []AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod `json:"airspaceTimePeriod,omitzero"`
	// Designates the geometric type that defines the airspace shape. One of CIRCLE,
	// CORRIDOR, LINE, ORBIT, etc.
	//
	// Any of "POLYARC", "1TRACK", "POLYGON", "CIRCLE", "CORRIDOR", "APOINT", "AORBIT",
	// "GEOLINE".
	CmShape string `json:"cmShape,omitzero"`
	// The frequency for the airspace control authority. Can specify HZ, KHZ, MHZ, GHZ
	// or a DESIG frequency designator code.
	CntrlAuthFreqs []string `json:"cntrlAuthFreqs,omitzero"`
	// An array of at least two alphanumeric symbols used to serially identify the
	// corridor waypoints. If cmShape is set to "CORRIDOR", one of either corrWayPoints
	// or polyCoord is required to specify the centerline of the corridor path.
	CorrWayPoints []string `json:"corrWayPoints,omitzero"`
	// A set of geospatial coordinates specified in DMS (Degrees, Minutes, Seconds)
	// format which determine the vertices of a one or two dimensional geospatial
	// shape. When cmShape is set to "POLYARC" or "POLYGON", this field is required as
	// applied in the construction of the area boundary. If cmShape is set to
	// "CORRIDOR" or "GEOLINE", this field is required and can be interpreted as an
	// ordered set of points along a path in space.
	PolyCoord []string `json:"polyCoord,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMean) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMean
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMean](
		"CmShape", false, "POLYARC", "1TRACK", "POLYGON", "CIRCLE", "CORRIDOR", "APOINT", "AORBIT", "GEOLINE",
	)
}

// The controlPoint set describes any reference/controlling/rendezvous point for a
// given airspace control means.
type AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint struct {
	// The altitude of the control point.
	CtrlPtAltitude param.Opt[string] `json:"ctrlPtAltitude,omitzero"`
	// A geospatial point coordinate specified in DMS (Degrees, Minutes, Seconds)
	// format that represents the location of the control point.
	CtrlPtLocation param.Opt[string] `json:"ctrlPtLocation,omitzero"`
	// The name applied to the control point, used as a reference.
	CtrlPtName param.Opt[string] `json:"ctrlPtName,omitzero"`
	// One of possible control point type codes, such as CP, ER, OT, etc.
	CtrlPtType param.Opt[string] `json:"ctrlPtType,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint
	return param.MarshalObject(r, (*shadow)(&r))
}

// The timePeriods set describes the effective datetime for a given airspace
// control means.
type AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod struct {
	// The end time designating that the airspace control order is no longer active.
	// Can contain datetime information or a stop time qualifier code, such as AFTER,
	// NET, UFN, etc.
	TimeEnd param.Opt[string] `json:"timeEnd,omitzero"`
	// The airspace time code associated with the ACO. Can be DISCRETE, a fixed time
	// block, or INTERVAL, a repeating time block.
	TimeMode param.Opt[string] `json:"timeMode,omitzero"`
	// The start time designating that the airspace control order is active.
	TimeStart param.Opt[string] `json:"timeStart,omitzero"`
	// Mandatory if timeMode is INTERVAL. Can be a numerical multiplier on an interval
	// frequency code, a stop time qualifier code such as AFTER, NET, UFN, etc, or a
	// datetime like string.
	IntDur []string `json:"intDur,omitzero"`
	// Mandatory if timeMode is INTERVAL. Can be one of the interval frequency codes,
	// such as BIWEEKLY, DAILY, YEARLY, etc.
	IntFreq []string `json:"intFreq,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod
	return param.MarshalObject(r, (*shadow)(&r))
}

// The airspaceControlReferences set provides both USMTF and non-USMTF references
// for this airspace control order.
type AirspaceControlOrderNewBulkParamsBodyAirspaceControlOrderReference struct {
	// The originator of this reference.
	RefOriginator param.Opt[string] `json:"refOriginator,omitzero"`
	// The reference serial number.
	RefSerialNum param.Opt[string] `json:"refSerialNum,omitzero"`
	// Specifies an alphabetic serial number identifying a reference pertaining to this
	// message.
	RefSID param.Opt[string] `json:"refSId,omitzero"`
	// Indicates any special actions, restrictions, guidance, or information relating
	// to this reference.
	RefSpecialNotation param.Opt[string] `json:"refSpecialNotation,omitzero"`
	// Timestamp of the referenced message, in ISO 8601 UTC format with millisecond
	// precision.
	RefTs param.Opt[time.Time] `json:"refTs,omitzero" format:"date-time"`
	// Specifies the type for this reference.
	RefType param.Opt[string] `json:"refType,omitzero"`
	// Array of NATO Subject Indicator Codes (SIC) or filing numbers of the document
	// being referenced.
	RefSiCs []string `json:"refSICs,omitzero"`
	paramObj
}

func (r AirspaceControlOrderNewBulkParamsBodyAirspaceControlOrderReference) MarshalJSON() (data []byte, err error) {
	type shadow AirspaceControlOrderNewBulkParamsBodyAirspaceControlOrderReference
	return param.MarshalObject(r, (*shadow)(&r))
}

type AirspaceControlOrderTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirspaceControlOrderTupleParams]'s query parameters as
// `url.Values`.
func (r AirspaceControlOrderTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
