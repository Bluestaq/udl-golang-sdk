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

// AIService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	Options []option.RequestOption
	History AIHistoryService
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r AIService) {
	r = AIService{}
	r.Options = opts
	r.History = NewAIHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AIService) List(ctx context.Context, query AIListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AIsAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/ais"
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
func (r *AIService) ListAutoPaging(ctx context.Context, query AIListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AIsAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AIService) Count(ctx context.Context, query AICountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/ais/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of AIS
// records as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *AIService) NewBulk(ctx context.Context, body AINewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/ais/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AIService) HistoryCount(ctx context.Context, query AIHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/ais/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AIService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *AIQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/ais/queryhelp"
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
func (r *AIService) Tuple(ctx context.Context, query AITupleParams, opts ...option.RequestOption) (res *[]shared.AIsFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/ais/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Self-reported information obtained from Automatic Identification System (AIS)
// equipment. This contains information such as unique identification, status,
// position, course, and speed. The AIS is an automatic tracking system that uses
// transceivers on ships and is used by vessel traffic services. Although
// technically and operationally distinct, the AIS system is analogous to ADS-B
// that performs a similar function for aircraft. AIS is intended to assist a
// vessel's watchstanding officers and allow maritime authorities to track and
// monitor vessel movements. AIS integrates a standardized VHF transceiver with a
// positioning system such as Global Positioning System receiver, with other
// electronic navigation sensors, such as gyrocompass or rate of turn indicator.
// Vessels fitted with AIS transceivers can be tracked by AIS base stations located
// along coast lines or, when out of range of terrestrial networks, through a
// growing number of satellites that are fitted with special AIS receivers which
// are capable of deconflicting a large number of signatures.
type AIsAbridged struct {
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
	DataMode AIsAbridgedDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp that the vessel position was recorded, in ISO 8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The reference dimensions of the vessel, reported as [A, B, C, D], in meters.
	// Where the array values represent the distance fore (A), aft (B), to port (C),
	// and to starboard (D) of the navigation antenna. Array with values A = C = 0 and
	// B, D > 0 indicate the length (B) and width (D) of the vessel without antenna
	// position reference.
	AntennaRefDimensions []float64 `json:"antennaRefDimensions"`
	// The average speed, in kilometers/hour, calculated for the subject vessel during
	// the latest voyage (port to port).
	AvgSpeed float64 `json:"avgSpeed"`
	// A uniquely designated identifier for the vessel's transmitter station.
	CallSign string `json:"callSign"`
	// The reported cargo type. Intended as, but not constrained to, the USCG NAVCEN
	// AIS cargo definitions. Users should refer to USCG Navigation Center
	// documentation for specific definitions associated with ship and cargo types.
	// USCG NAVCEN documentation may be found at https://www.navcen.uscg.gov.
	CargoType string `json:"cargoType"`
	// The course-over-ground reported by the vessel, in degrees.
	Course float64 `json:"course"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The US Geographic Unique Identifier of the current port hosting the vessel.
	CurrentPortGuid string `json:"currentPortGUID"`
	// The UN Location Code of the current port hosting the vessel.
	CurrentPortLocode string `json:"currentPortLOCODE"`
	// The destination of the vessel according to the AIS transmission.
	Destination string `json:"destination"`
	// The Estimated Time of Arrival of the vessel at the destination, in ISO 8601 UTC
	// format.
	DestinationEta time.Time `json:"destinationETA" format:"date-time"`
	// The remaining distance, in kilometers, for the vessel to reach the reported
	// destination.
	DistanceToGo float64 `json:"distanceToGo"`
	// The distance, in kilometers, that the vessel has travelled since departing the
	// last port.
	DistanceTravelled float64 `json:"distanceTravelled"`
	// The maximum static draught, in meters, of the vessel according to the AIS
	// transmission.
	Draught float64 `json:"draught"`
	// The activity that the vessel is engaged in. This entry applies only when the
	// shipType = Other.
	EngagedIn string `json:"engagedIn"`
	// The Estimated Time of Arrival of the vessel at the destination port, according
	// to MarineTraffic calculations, in ISO 8601 UTC format.
	EtaCalculated time.Time `json:"etaCalculated" format:"date-time"`
	// The date and time that the ETA was calculated by MarineTraffic, in ISO 8601 UTC
	// format.
	EtaUpdated time.Time `json:"etaUpdated" format:"date-time"`
	// Unique identifier of the Track.
	IDTrack string `json:"idTrack"`
	// Unique identifier of the vessel.
	IDVessel string `json:"idVessel"`
	// The International Maritime Organization Number of the vessel. IMON is a
	// seven-digit number that uniquely identifies the vessel.
	Imon int64 `json:"imon"`
	// The US Geographic Unique Identifier of the last port visited by the vessel.
	LastPortGuid string `json:"lastPortGUID"`
	// The UN Location Code of the last port visited by the vessel.
	LastPortLocode string `json:"lastPortLOCODE"`
	// WGS-84 latitude of the vessel position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	Lat float64 `json:"lat"`
	// The overall length of the vessel, in meters. A value of 511 indicates a vessel
	// length of 511 meters or greater.
	Length float64 `json:"length"`
	// WGS-84 longitude of the vessel position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The maximum speed, in kilometers/hour, reported by the subject vessel during the
	// latest voyage (port to port).
	MaxSpeed float64 `json:"maxSpeed"`
	// The Maritime Mobile Service Identity of the vessel. MMSI is a nine-digit number
	// that identifies the transmitter station of the vessel.
	Mmsi int64 `json:"mmsi"`
	// The AIS Navigational Status of the vessel (e.g. Underway Using Engine, Moored,
	// Aground, etc.). Intended as, but not constrained to, the USCG NAVCEN navigation
	// status definitions. Users should refer to USCG Navigation Center documentation
	// for specific definitions associated with navigation status. USCG NAVCEN
	// documentation may be found at https://www.navcen.uscg.gov.
	NavStatus string `json:"navStatus"`
	// The US Geographic Unique Identifier of the next destination port of the vessel.
	NextPortGuid string `json:"nextPortGUID"`
	// The UN Location Code of the next destination port of the vessel.
	NextPortLocode string `json:"nextPortLOCODE"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The type of electronic position fixing device (e.g. GPS, GLONASS, etc.).
	// Intended as, but not constrained to, the USCG NAVCEN electronic position fixing
	// device definitions. Users should refer to USCG Navigation Center documentation
	// for specific device type information. USCG NAVCEN documentation may be found at
	// https://www.navcen.uscg.gov.
	PosDeviceType string `json:"posDeviceType"`
	// Flag indicating high reported position accuracy (less than or equal to 10
	// meters). A value of 0/false indicates low accuracy (greater than 10 meters).
	PosHiAccuracy bool `json:"posHiAccuracy"`
	// Flag indicating high reported position latency (greater than 5 seconds). A value
	// of 0/false indicates low latency (less than 5 seconds).
	PosHiLatency bool `json:"posHiLatency"`
	// The Rate-of-Turn for the vessel, in degrees/minute. Positive value indicates
	// that the vessel is turning right.
	RateOfTurn float64 `json:"rateOfTurn"`
	// Further description or explanation of the vessel or type.
	ShipDescription string `json:"shipDescription"`
	// The name of the vessel. Vessel names that exceed the AIS 20 character are
	// shortened (not truncated) to 15 character-spaces, followed by an underscore and
	// the last 4 characters-spaces of the vessel full name.
	ShipName string `json:"shipName"`
	// The reported ship type (e.g. Passenger, Tanker, Cargo, Other, etc.). See the
	// engagedIn and specialCraft entries for additional information on certain types
	// of vessels.
	ShipType string `json:"shipType"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The type of special craft designation of the vessel. This entry applies only
	// when the shipType = Special Craft.
	SpecialCraft string `json:"specialCraft"`
	// Flag indicating that the vessel is engaged in a special maneuver (e.g. Waterway
	// Navigation).
	SpecialManeuver bool `json:"specialManeuver"`
	// The speed-over-ground reported by the vessel, in kilometers/hour.
	Speed float64 `json:"speed"`
	// The true heading reported by the vessel, in degrees.
	TrueHeading float64 `json:"trueHeading"`
	// The flag of the subject vessel according to AIS transmission.
	VesselFlag string `json:"vesselFlag"`
	// The breadth of the vessel, in meters. A value of 63 indicates a vessel breadth
	// of 63 meters or greater.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		AntennaRefDimensions  respjson.Field
		AvgSpeed              respjson.Field
		CallSign              respjson.Field
		CargoType             respjson.Field
		Course                respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		CurrentPortGuid       respjson.Field
		CurrentPortLocode     respjson.Field
		Destination           respjson.Field
		DestinationEta        respjson.Field
		DistanceToGo          respjson.Field
		DistanceTravelled     respjson.Field
		Draught               respjson.Field
		EngagedIn             respjson.Field
		EtaCalculated         respjson.Field
		EtaUpdated            respjson.Field
		IDTrack               respjson.Field
		IDVessel              respjson.Field
		Imon                  respjson.Field
		LastPortGuid          respjson.Field
		LastPortLocode        respjson.Field
		Lat                   respjson.Field
		Length                respjson.Field
		Lon                   respjson.Field
		MaxSpeed              respjson.Field
		Mmsi                  respjson.Field
		NavStatus             respjson.Field
		NextPortGuid          respjson.Field
		NextPortLocode        respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PosDeviceType         respjson.Field
		PosHiAccuracy         respjson.Field
		PosHiLatency          respjson.Field
		RateOfTurn            respjson.Field
		ShipDescription       respjson.Field
		ShipName              respjson.Field
		ShipType              respjson.Field
		SourceDl              respjson.Field
		SpecialCraft          respjson.Field
		SpecialManeuver       respjson.Field
		Speed                 respjson.Field
		TrueHeading           respjson.Field
		VesselFlag            respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIsAbridged) RawJSON() string { return r.JSON.raw }
func (r *AIsAbridged) UnmarshalJSON(data []byte) error {
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
type AIsAbridgedDataMode string

const (
	AIsAbridgedDataModeReal      AIsAbridgedDataMode = "REAL"
	AIsAbridgedDataModeTest      AIsAbridgedDataMode = "TEST"
	AIsAbridgedDataModeSimulated AIsAbridgedDataMode = "SIMULATED"
	AIsAbridgedDataModeExercise  AIsAbridgedDataMode = "EXERCISE"
)

type AIQueryhelpResponse struct {
	AodrSupported         bool                           `json:"aodrSupported"`
	ClassificationMarking string                         `json:"classificationMarking"`
	Description           string                         `json:"description"`
	HistorySupported      bool                           `json:"historySupported"`
	Name                  string                         `json:"name"`
	Parameters            []AIQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                       `json:"requiredRoles"`
	RestSupported         bool                           `json:"restSupported"`
	SortSupported         bool                           `json:"sortSupported"`
	TypeName              string                         `json:"typeName"`
	Uri                   string                         `json:"uri"`
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
func (r AIQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *AIQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIQueryhelpResponseParameter struct {
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
func (r AIQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *AIQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIListParams struct {
	// The timestamp that the vessel position was recorded, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIListParams]'s query parameters as `url.Values`.
func (r AIListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AICountParams struct {
	// The timestamp that the vessel position was recorded, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AICountParams]'s query parameters as `url.Values`.
func (r AICountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AINewBulkParams struct {
	Body []AINewBulkParamsBody
	paramObj
}

func (r AINewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *AINewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Self-reported information obtained from Automatic Identification System (AIS)
// equipment. This contains information such as unique identification, status,
// position, course, and speed. The AIS is an automatic tracking system that uses
// transceivers on ships and is used by vessel traffic services. Although
// technically and operationally distinct, the AIS system is analogous to ADS-B
// that performs a similar function for aircraft. AIS is intended to assist a
// vessel's watchstanding officers and allow maritime authorities to track and
// monitor vessel movements. AIS integrates a standardized VHF transceiver with a
// positioning system such as Global Positioning System receiver, with other
// electronic navigation sensors, such as gyrocompass or rate of turn indicator.
// Vessels fitted with AIS transceivers can be tracked by AIS base stations located
// along coast lines or, when out of range of terrestrial networks, through a
// growing number of satellites that are fitted with special AIS receivers which
// are capable of deconflicting a large number of signatures.
//
// The properties ClassificationMarking, DataMode, Source, Ts are required.
type AINewBulkParamsBody struct {
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
	// The timestamp that the vessel position was recorded, in ISO 8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The average speed, in kilometers/hour, calculated for the subject vessel during
	// the latest voyage (port to port).
	AvgSpeed param.Opt[float64] `json:"avgSpeed,omitzero"`
	// A uniquely designated identifier for the vessel's transmitter station.
	CallSign param.Opt[string] `json:"callSign,omitzero"`
	// The reported cargo type. Intended as, but not constrained to, the USCG NAVCEN
	// AIS cargo definitions. Users should refer to USCG Navigation Center
	// documentation for specific definitions associated with ship and cargo types.
	// USCG NAVCEN documentation may be found at https://www.navcen.uscg.gov.
	CargoType param.Opt[string] `json:"cargoType,omitzero"`
	// The course-over-ground reported by the vessel, in degrees.
	Course param.Opt[float64] `json:"course,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The US Geographic Unique Identifier of the current port hosting the vessel.
	CurrentPortGuid param.Opt[string] `json:"currentPortGUID,omitzero"`
	// The UN Location Code of the current port hosting the vessel.
	CurrentPortLocode param.Opt[string] `json:"currentPortLOCODE,omitzero"`
	// The destination of the vessel according to the AIS transmission.
	Destination param.Opt[string] `json:"destination,omitzero"`
	// The Estimated Time of Arrival of the vessel at the destination, in ISO 8601 UTC
	// format.
	DestinationEta param.Opt[time.Time] `json:"destinationETA,omitzero" format:"date-time"`
	// The remaining distance, in kilometers, for the vessel to reach the reported
	// destination.
	DistanceToGo param.Opt[float64] `json:"distanceToGo,omitzero"`
	// The distance, in kilometers, that the vessel has travelled since departing the
	// last port.
	DistanceTravelled param.Opt[float64] `json:"distanceTravelled,omitzero"`
	// The maximum static draught, in meters, of the vessel according to the AIS
	// transmission.
	Draught param.Opt[float64] `json:"draught,omitzero"`
	// The activity that the vessel is engaged in. This entry applies only when the
	// shipType = Other.
	EngagedIn param.Opt[string] `json:"engagedIn,omitzero"`
	// The Estimated Time of Arrival of the vessel at the destination port, according
	// to MarineTraffic calculations, in ISO 8601 UTC format.
	EtaCalculated param.Opt[time.Time] `json:"etaCalculated,omitzero" format:"date-time"`
	// The date and time that the ETA was calculated by MarineTraffic, in ISO 8601 UTC
	// format.
	EtaUpdated param.Opt[time.Time] `json:"etaUpdated,omitzero" format:"date-time"`
	// Unique identifier of the Track.
	IDTrack param.Opt[string] `json:"idTrack,omitzero"`
	// Unique identifier of the vessel.
	IDVessel param.Opt[string] `json:"idVessel,omitzero"`
	// The International Maritime Organization Number of the vessel. IMON is a
	// seven-digit number that uniquely identifies the vessel.
	Imon param.Opt[int64] `json:"imon,omitzero"`
	// The US Geographic Unique Identifier of the last port visited by the vessel.
	LastPortGuid param.Opt[string] `json:"lastPortGUID,omitzero"`
	// The UN Location Code of the last port visited by the vessel.
	LastPortLocode param.Opt[string] `json:"lastPortLOCODE,omitzero"`
	// WGS-84 latitude of the vessel position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// The overall length of the vessel, in meters. A value of 511 indicates a vessel
	// length of 511 meters or greater.
	Length param.Opt[float64] `json:"length,omitzero"`
	// WGS-84 longitude of the vessel position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The maximum speed, in kilometers/hour, reported by the subject vessel during the
	// latest voyage (port to port).
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// The Maritime Mobile Service Identity of the vessel. MMSI is a nine-digit number
	// that identifies the transmitter station of the vessel.
	Mmsi param.Opt[int64] `json:"mmsi,omitzero"`
	// The AIS Navigational Status of the vessel (e.g. Underway Using Engine, Moored,
	// Aground, etc.). Intended as, but not constrained to, the USCG NAVCEN navigation
	// status definitions. Users should refer to USCG Navigation Center documentation
	// for specific definitions associated with navigation status. USCG NAVCEN
	// documentation may be found at https://www.navcen.uscg.gov.
	NavStatus param.Opt[string] `json:"navStatus,omitzero"`
	// The US Geographic Unique Identifier of the next destination port of the vessel.
	NextPortGuid param.Opt[string] `json:"nextPortGUID,omitzero"`
	// The UN Location Code of the next destination port of the vessel.
	NextPortLocode param.Opt[string] `json:"nextPortLOCODE,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The type of electronic position fixing device (e.g. GPS, GLONASS, etc.).
	// Intended as, but not constrained to, the USCG NAVCEN electronic position fixing
	// device definitions. Users should refer to USCG Navigation Center documentation
	// for specific device type information. USCG NAVCEN documentation may be found at
	// https://www.navcen.uscg.gov.
	PosDeviceType param.Opt[string] `json:"posDeviceType,omitzero"`
	// Flag indicating high reported position accuracy (less than or equal to 10
	// meters). A value of 0/false indicates low accuracy (greater than 10 meters).
	PosHiAccuracy param.Opt[bool] `json:"posHiAccuracy,omitzero"`
	// Flag indicating high reported position latency (greater than 5 seconds). A value
	// of 0/false indicates low latency (less than 5 seconds).
	PosHiLatency param.Opt[bool] `json:"posHiLatency,omitzero"`
	// The Rate-of-Turn for the vessel, in degrees/minute. Positive value indicates
	// that the vessel is turning right.
	RateOfTurn param.Opt[float64] `json:"rateOfTurn,omitzero"`
	// Further description or explanation of the vessel or type.
	ShipDescription param.Opt[string] `json:"shipDescription,omitzero"`
	// The name of the vessel. Vessel names that exceed the AIS 20 character are
	// shortened (not truncated) to 15 character-spaces, followed by an underscore and
	// the last 4 characters-spaces of the vessel full name.
	ShipName param.Opt[string] `json:"shipName,omitzero"`
	// The reported ship type (e.g. Passenger, Tanker, Cargo, Other, etc.). See the
	// engagedIn and specialCraft entries for additional information on certain types
	// of vessels.
	ShipType param.Opt[string] `json:"shipType,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The type of special craft designation of the vessel. This entry applies only
	// when the shipType = Special Craft.
	SpecialCraft param.Opt[string] `json:"specialCraft,omitzero"`
	// Flag indicating that the vessel is engaged in a special maneuver (e.g. Waterway
	// Navigation).
	SpecialManeuver param.Opt[bool] `json:"specialManeuver,omitzero"`
	// The speed-over-ground reported by the vessel, in kilometers/hour.
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// The true heading reported by the vessel, in degrees.
	TrueHeading param.Opt[float64] `json:"trueHeading,omitzero"`
	// The flag of the subject vessel according to AIS transmission.
	VesselFlag param.Opt[string] `json:"vesselFlag,omitzero"`
	// The breadth of the vessel, in meters. A value of 63 indicates a vessel breadth
	// of 63 meters or greater.
	Width param.Opt[float64] `json:"width,omitzero"`
	// The reference dimensions of the vessel, reported as [A, B, C, D], in meters.
	// Where the array values represent the distance fore (A), aft (B), to port (C),
	// and to starboard (D) of the navigation antenna. Array with values A = C = 0 and
	// B, D > 0 indicate the length (B) and width (D) of the vessel without antenna
	// position reference.
	AntennaRefDimensions []float64 `json:"antennaRefDimensions,omitzero"`
	paramObj
}

func (r AINewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AINewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AINewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AINewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type AIHistoryCountParams struct {
	// The timestamp that the vessel position was recorded, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIHistoryCountParams]'s query parameters as `url.Values`.
func (r AIHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AITupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The timestamp that the vessel position was recorded, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AITupleParams]'s query parameters as `url.Values`.
func (r AITupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
