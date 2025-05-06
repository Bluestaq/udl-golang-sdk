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

// EquipmentService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEquipmentService] method instead.
type EquipmentService struct {
	Options []option.RequestOption
}

// NewEquipmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEquipmentService(opts ...option.RequestOption) (r EquipmentService) {
	r = EquipmentService{}
	r.Options = opts
	return
}

// Service operation to take a single equipment record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *EquipmentService) New(ctx context.Context, body EquipmentNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/equipment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single equipment record by its unique ID passed as a
// path parameter.
func (r *EquipmentService) Get(ctx context.Context, id string, query EquipmentGetParams, opts ...option.RequestOption) (res *EquipmentFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/equipment/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single equipment record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *EquipmentService) Update(ctx context.Context, id string, body EquipmentUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/equipment/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EquipmentService) List(ctx context.Context, query EquipmentListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EquipmentAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/equipment"
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
func (r *EquipmentService) ListAutoPaging(ctx context.Context, query EquipmentListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EquipmentAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a equipment record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *EquipmentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/equipment/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EquipmentService) Count(ctx context.Context, query EquipmentCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/equipment/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// Equipment records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *EquipmentService) NewBulk(ctx context.Context, body EquipmentNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/equipment/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EquipmentService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/equipment/queryhelp"
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
func (r *EquipmentService) Tuple(ctx context.Context, query EquipmentTupleParams, opts ...option.RequestOption) (res *[]EquipmentFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/equipment/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Properties and characteristics of equipment that can be associated with a site
// or other entity.
type EquipmentAbridged struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment geographic coordinates reside. This field will be set to
	// "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode,required"`
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
	DataMode EquipmentAbridgedDataMode `json:"dataMode,required"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea string `json:"airDefArea"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance string `json:"allegiance"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance string `json:"altAllegiance"`
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS. This field will be set
	// to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode string `json:"altCountryCode"`
	// Unique identifier of the Equipment record from the originating system.
	AltEqpID string `json:"altEqpId"`
	// Indicates the importance of the equipment. Referenced, but not constrained to,
	// the following class ratings type classifications.
	//
	// 0 - Not of significant importance of the system
	//
	// 1 - Primary importance to system
	//
	// 2 - Secondary importance to system
	//
	// 3 - Tertiary importance to system
	//
	// O - Other. Explain in Remarks.
	ClassRating string `json:"classRating"`
	// The physical manner of being or state of existence of the entity. A physical
	// condition that must be considered in the determining of a course of action. The
	// specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Condition string `json:"condition"`
	// Availability of the entity relative to its condition. Indicates the reason the
	// entity is not fully operational. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	ConditionAvail string `json:"conditionAvail"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc. specified in degrees, minute,
	// and seconds.
	//
	// Pos. 1-2. Latitude Degrees [00-90]
	//
	// Pos. 3-4. Latitude Minutes [00-59]
	//
	// Pos. 5-6. Latitude Seconds [00-59]
	//
	// Pos. 7-9. Latitude Thousandths Of Seconds [000-999]
	//
	// Pos. 10. Latitude Hemisphere [NS]
	//
	// Pos. 11-13. Longitude Degrees [00-180]
	//
	// Pos. 14-15. Longitude Minutes [00-59]
	//
	// Pos. 16-17. Longitude Seconds [00-59]
	//
	// Pos. 18-20. Longitude Thousandths Of Seconds [000-999]
	//
	// Pos. 21. Longitude Hemisphere [EW]
	//
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]].
	Coord string `json:"coord"`
	// A mathematical model of the earth used to calculate coordinates on a map. US
	// Forces use the World Geodetic System 1984 (WGS 84), but also use maps by allied
	// countries with local datums. The datum must be specified to ensure accuracy of
	// coordinates. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	CoordDatum string `json:"coordDatum"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// coordinate.
	CoordDerivAcc float64 `json:"coordDerivAcc"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Ground elevation, in meters, of the geographic coordinates referenced to (above
	// or below) Mean Sea Level (MSL) vertical datum.
	ElevMsl float64 `json:"elevMsl"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy. Confidence level has a range of values
	// from 0 to 100, with 100 being highest level of confidence.
	ElevMslConfLvl int64 `json:"elevMslConfLvl"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation, measured in meters.
	ElevMslDerivAcc float64 `json:"elevMslDerivAcc"`
	// Designated equipment code assigned to the item of equipment or an abbreviation
	// record type unique identifier. Users should consult the data provider for
	// information on the equipment code structure.
	EqpCode string `json:"eqpCode"`
	// Uniquely identifies each item or group of equipment associated with a unit,
	// facility or site.
	EqpIDNum string `json:"eqpIdNum"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval int64 `json:"eval"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa string `json:"fpa"`
	// Indicates the function or mission of this equipment, which may or may not be
	// engaged in at any particular time. Typically refers to a unit, organization, or
	// installation/facility performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Function string `json:"function"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary string `json:"functPrimary"`
	// The distance between Mean Sea Level and a referenced ellipsoid, measured in
	// meters.
	GeoidalMslSep float64 `json:"geoidalMslSep"`
	// Estimated identity of the Site (ASSUMED FRIEND, FRIEND, HOSTILE, FAKER, JOKER,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	Ident string `json:"ident"`
	// Unique identifier of the operating unit associated with the equipment record.
	IDOperatingUnit string `json:"idOperatingUnit"`
	// Unique identifier of the Parent equipment record associated with this equipment
	// record.
	IDParentEquipment string `json:"idParentEquipment"`
	// Unique identifier of the Site Entity associated with the equipment record.
	IDSite string `json:"idSite"`
	// Indicates the reason that the equipment is at that location. The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason string `json:"locReason"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts:
	//
	// 4Q (grid zone designator, GZD)
	//
	// FJ (the 100,000-meter square identifier)
	//
	// 12345678 (numerical location; easting is 1234 and northing is 5678, in this case
	// specifying a location with 10 m resolution).
	MilGrid string `json:"milGrid"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator.
	MilGridSys string `json:"milGridSys"`
	// Generic type this specific piece of equipment belongs to, and the identifying
	// nomenclature which describes the equipment.
	Nomen string `json:"nomen"`
	// Internationally recognized water area in which the vessel is most likely to be
	// deployed or in which it normally operates most frequently.
	OperAreaPrimary string `json:"operAreaPrimary"`
	// The Degree to which an entity is ready to perform the overall operational
	// mission(s) for which it was organized and equipped. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	OperStatus string `json:"operStatus"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv string `json:"polSubdiv"`
	// Relative to the parent entity, the total number of military personnel or
	// equipment assessed to be on-hand (OH).
	QtyOh int64 `json:"qtyOH"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Referenced, but not
	// constrained to, the following record status type classifications.
	//
	// # A - Active
	//
	// # I - Inactive
	//
	// # K - Acknowledged
	//
	// # L - Local
	//
	// # Q - A nominated (NOM) or Data Change Request (DCR) record
	//
	// # R - Production reduced by CMD decision
	//
	// W - Working Record.
	RecStatus string `json:"recStatus"`
	// The reference documentiation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc string `json:"referenceDoc"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd string `json:"resProd"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency. This date cannot be greater than the current
	// date.
	ReviewDate time.Time `json:"reviewDate" format:"date"`
	// Provider specific sequential number assigned to the equipment.
	SeqNum int64 `json:"seqNum"`
	// Array of UUID(s) of the UDL data record(s) that are related to this equipment
	// record. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types such as AIRCRAFT, VESSEL, EO, MTI that are related to
	// this equipment record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode string `json:"symCode"`
	// Universal Transverse Mercator (UTM) grid coordinates.
	//
	// Pos. 1-2, UTM Zone Column [01-60
	//
	// Pos. 3, UTM Zone Row [C-HJ-NP-X]
	//
	// Pos. 4, UTM False Easting [0-9]
	//
	// Pos. 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9]
	//
	// Pos. 10-11, UTM False Northing [0-9][0-9]
	//
	// Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm string `json:"utm"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac string `json:"wac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		CountryCode           respjson.Field
		DataMode              respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AirDefArea            respjson.Field
		Allegiance            respjson.Field
		AltAllegiance         respjson.Field
		AltCountryCode        respjson.Field
		AltEqpID              respjson.Field
		ClassRating           respjson.Field
		Condition             respjson.Field
		ConditionAvail        respjson.Field
		Coord                 respjson.Field
		CoordDatum            respjson.Field
		CoordDerivAcc         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ElevMsl               respjson.Field
		ElevMslConfLvl        respjson.Field
		ElevMslDerivAcc       respjson.Field
		EqpCode               respjson.Field
		EqpIDNum              respjson.Field
		Eval                  respjson.Field
		Fpa                   respjson.Field
		Function              respjson.Field
		FunctPrimary          respjson.Field
		GeoidalMslSep         respjson.Field
		Ident                 respjson.Field
		IDOperatingUnit       respjson.Field
		IDParentEquipment     respjson.Field
		IDSite                respjson.Field
		LocReason             respjson.Field
		MilGrid               respjson.Field
		MilGridSys            respjson.Field
		Nomen                 respjson.Field
		OperAreaPrimary       respjson.Field
		OperStatus            respjson.Field
		Origin                respjson.Field
		PolSubdiv             respjson.Field
		QtyOh                 respjson.Field
		RecStatus             respjson.Field
		ReferenceDoc          respjson.Field
		ResProd               respjson.Field
		ReviewDate            respjson.Field
		SeqNum                respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		SymCode               respjson.Field
		Utm                   respjson.Field
		Wac                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EquipmentAbridged) RawJSON() string { return r.JSON.raw }
func (r *EquipmentAbridged) UnmarshalJSON(data []byte) error {
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
type EquipmentAbridgedDataMode string

const (
	EquipmentAbridgedDataModeReal      EquipmentAbridgedDataMode = "REAL"
	EquipmentAbridgedDataModeTest      EquipmentAbridgedDataMode = "TEST"
	EquipmentAbridgedDataModeSimulated EquipmentAbridgedDataMode = "SIMULATED"
	EquipmentAbridgedDataModeExercise  EquipmentAbridgedDataMode = "EXERCISE"
)

// Properties and characteristics of equipment that can be associated with a site
// or other entity.
type EquipmentFull struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment geographic coordinates reside. This field will be set to
	// "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode,required"`
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
	DataMode EquipmentFullDataMode `json:"dataMode,required"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea string `json:"airDefArea"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance string `json:"allegiance"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance string `json:"altAllegiance"`
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS. This field will be set
	// to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode string `json:"altCountryCode"`
	// Unique identifier of the Equipment record from the originating system.
	AltEqpID string `json:"altEqpId"`
	// Indicates the importance of the equipment. Referenced, but not constrained to,
	// the following class ratings type classifications.
	//
	// 0 - Not of significant importance of the system
	//
	// 1 - Primary importance to system
	//
	// 2 - Secondary importance to system
	//
	// 3 - Tertiary importance to system
	//
	// O - Other. Explain in Remarks.
	ClassRating string `json:"classRating"`
	// The physical manner of being or state of existence of the entity. A physical
	// condition that must be considered in the determining of a course of action. The
	// specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Condition string `json:"condition"`
	// Availability of the entity relative to its condition. Indicates the reason the
	// entity is not fully operational. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	ConditionAvail string `json:"conditionAvail"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc. specified in degrees, minute,
	// and seconds.
	//
	// Pos. 1-2. Latitude Degrees [00-90]
	//
	// Pos. 3-4. Latitude Minutes [00-59]
	//
	// Pos. 5-6. Latitude Seconds [00-59]
	//
	// Pos. 7-9. Latitude Thousandths Of Seconds [000-999]
	//
	// Pos. 10. Latitude Hemisphere [NS]
	//
	// Pos. 11-13. Longitude Degrees [00-180]
	//
	// Pos. 14-15. Longitude Minutes [00-59]
	//
	// Pos. 16-17. Longitude Seconds [00-59]
	//
	// Pos. 18-20. Longitude Thousandths Of Seconds [000-999]
	//
	// Pos. 21. Longitude Hemisphere [EW]
	//
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]].
	Coord string `json:"coord"`
	// A mathematical model of the earth used to calculate coordinates on a map. US
	// Forces use the World Geodetic System 1984 (WGS 84), but also use maps by allied
	// countries with local datums. The datum must be specified to ensure accuracy of
	// coordinates. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	CoordDatum string `json:"coordDatum"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// coordinate.
	CoordDerivAcc float64 `json:"coordDerivAcc"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Ground elevation, in meters, of the geographic coordinates referenced to (above
	// or below) Mean Sea Level (MSL) vertical datum.
	ElevMsl float64 `json:"elevMsl"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy. Confidence level has a range of values
	// from 0 to 100, with 100 being highest level of confidence.
	ElevMslConfLvl int64 `json:"elevMslConfLvl"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation, measured in meters.
	ElevMslDerivAcc float64 `json:"elevMslDerivAcc"`
	// Designated equipment code assigned to the item of equipment or an abbreviation
	// record type unique identifier. Users should consult the data provider for
	// information on the equipment code structure.
	EqpCode string `json:"eqpCode"`
	// Uniquely identifies each item or group of equipment associated with a unit,
	// facility or site.
	EqpIDNum string `json:"eqpIdNum"`
	// Remarks contain amplifying information for a specific service. The information
	// may contain context and interpretations for consumer use.
	EquipmentRemarks []EquipmentRemarkFull `json:"equipmentRemarks"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval int64 `json:"eval"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa string `json:"fpa"`
	// Indicates the function or mission of this equipment, which may or may not be
	// engaged in at any particular time. Typically refers to a unit, organization, or
	// installation/facility performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Function string `json:"function"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary string `json:"functPrimary"`
	// The distance between Mean Sea Level and a referenced ellipsoid, measured in
	// meters.
	GeoidalMslSep float64 `json:"geoidalMslSep"`
	// Estimated identity of the Site (ASSUMED FRIEND, FRIEND, HOSTILE, FAKER, JOKER,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	Ident string `json:"ident"`
	// Unique identifier of the operating unit associated with the equipment record.
	IDOperatingUnit string `json:"idOperatingUnit"`
	// Unique identifier of the Parent equipment record associated with this equipment
	// record.
	IDParentEquipment string `json:"idParentEquipment"`
	// Unique identifier of the Site Entity associated with the equipment record.
	IDSite string `json:"idSite"`
	// Indicates the reason that the equipment is at that location. The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason string `json:"locReason"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts:
	//
	// 4Q (grid zone designator, GZD)
	//
	// FJ (the 100,000-meter square identifier)
	//
	// 12345678 (numerical location; easting is 1234 and northing is 5678, in this case
	// specifying a location with 10 m resolution).
	MilGrid string `json:"milGrid"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator.
	MilGridSys string `json:"milGridSys"`
	// Generic type this specific piece of equipment belongs to, and the identifying
	// nomenclature which describes the equipment.
	Nomen string `json:"nomen"`
	// Internationally recognized water area in which the vessel is most likely to be
	// deployed or in which it normally operates most frequently.
	OperAreaPrimary string `json:"operAreaPrimary"`
	// The Degree to which an entity is ready to perform the overall operational
	// mission(s) for which it was organized and equipped. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	OperStatus string `json:"operStatus"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv string `json:"polSubdiv"`
	// Relative to the parent entity, the total number of military personnel or
	// equipment assessed to be on-hand (OH).
	QtyOh int64 `json:"qtyOH"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Referenced, but not
	// constrained to, the following record status type classifications.
	//
	// # A - Active
	//
	// # I - Inactive
	//
	// # K - Acknowledged
	//
	// # L - Local
	//
	// # Q - A nominated (NOM) or Data Change Request (DCR) record
	//
	// # R - Production reduced by CMD decision
	//
	// W - Working Record.
	RecStatus string `json:"recStatus"`
	// The reference documentiation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc string `json:"referenceDoc"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd string `json:"resProd"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency. This date cannot be greater than the current
	// date.
	ReviewDate time.Time `json:"reviewDate" format:"date"`
	// Provider specific sequential number assigned to the equipment.
	SeqNum int64 `json:"seqNum"`
	// Array of UUID(s) of the UDL data record(s) that are related to this equipment
	// record. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types such as AIRCRAFT, VESSEL, EO, MTI that are related to
	// this equipment record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode string `json:"symCode"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Universal Transverse Mercator (UTM) grid coordinates.
	//
	// Pos. 1-2, UTM Zone Column [01-60
	//
	// Pos. 3, UTM Zone Row [C-HJ-NP-X]
	//
	// Pos. 4, UTM False Easting [0-9]
	//
	// Pos. 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9]
	//
	// Pos. 10-11, UTM False Northing [0-9][0-9]
	//
	// Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm string `json:"utm"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac string `json:"wac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		CountryCode           respjson.Field
		DataMode              respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AirDefArea            respjson.Field
		Allegiance            respjson.Field
		AltAllegiance         respjson.Field
		AltCountryCode        respjson.Field
		AltEqpID              respjson.Field
		ClassRating           respjson.Field
		Condition             respjson.Field
		ConditionAvail        respjson.Field
		Coord                 respjson.Field
		CoordDatum            respjson.Field
		CoordDerivAcc         respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		ElevMsl               respjson.Field
		ElevMslConfLvl        respjson.Field
		ElevMslDerivAcc       respjson.Field
		EqpCode               respjson.Field
		EqpIDNum              respjson.Field
		EquipmentRemarks      respjson.Field
		Eval                  respjson.Field
		Fpa                   respjson.Field
		Function              respjson.Field
		FunctPrimary          respjson.Field
		GeoidalMslSep         respjson.Field
		Ident                 respjson.Field
		IDOperatingUnit       respjson.Field
		IDParentEquipment     respjson.Field
		IDSite                respjson.Field
		LocReason             respjson.Field
		MilGrid               respjson.Field
		MilGridSys            respjson.Field
		Nomen                 respjson.Field
		OperAreaPrimary       respjson.Field
		OperStatus            respjson.Field
		Origin                respjson.Field
		PolSubdiv             respjson.Field
		QtyOh                 respjson.Field
		RecStatus             respjson.Field
		ReferenceDoc          respjson.Field
		ResProd               respjson.Field
		ReviewDate            respjson.Field
		SeqNum                respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		SymCode               respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Utm                   respjson.Field
		Wac                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EquipmentFull) RawJSON() string { return r.JSON.raw }
func (r *EquipmentFull) UnmarshalJSON(data []byte) error {
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
type EquipmentFullDataMode string

const (
	EquipmentFullDataModeReal      EquipmentFullDataMode = "REAL"
	EquipmentFullDataModeTest      EquipmentFullDataMode = "TEST"
	EquipmentFullDataModeSimulated EquipmentFullDataMode = "SIMULATED"
	EquipmentFullDataModeExercise  EquipmentFullDataMode = "EXERCISE"
)

type EquipmentNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment geographic coordinates reside. This field will be set to
	// "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode,required"`
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
	DataMode EquipmentNewParamsDataMode `json:"dataMode,omitzero,required"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea param.Opt[string] `json:"airDefArea,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance param.Opt[string] `json:"allegiance,omitzero"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance param.Opt[string] `json:"altAllegiance,omitzero"`
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS. This field will be set
	// to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Unique identifier of the Equipment record from the originating system.
	AltEqpID param.Opt[string] `json:"altEqpId,omitzero"`
	// Indicates the importance of the equipment. Referenced, but not constrained to,
	// the following class ratings type classifications.
	//
	// 0 - Not of significant importance of the system
	//
	// 1 - Primary importance to system
	//
	// 2 - Secondary importance to system
	//
	// 3 - Tertiary importance to system
	//
	// O - Other. Explain in Remarks.
	ClassRating param.Opt[string] `json:"classRating,omitzero"`
	// The physical manner of being or state of existence of the entity. A physical
	// condition that must be considered in the determining of a course of action. The
	// specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Condition param.Opt[string] `json:"condition,omitzero"`
	// Availability of the entity relative to its condition. Indicates the reason the
	// entity is not fully operational. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	ConditionAvail param.Opt[string] `json:"conditionAvail,omitzero"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc. specified in degrees, minute,
	// and seconds.
	//
	// Pos. 1-2. Latitude Degrees [00-90]
	//
	// Pos. 3-4. Latitude Minutes [00-59]
	//
	// Pos. 5-6. Latitude Seconds [00-59]
	//
	// Pos. 7-9. Latitude Thousandths Of Seconds [000-999]
	//
	// Pos. 10. Latitude Hemisphere [NS]
	//
	// Pos. 11-13. Longitude Degrees [00-180]
	//
	// Pos. 14-15. Longitude Minutes [00-59]
	//
	// Pos. 16-17. Longitude Seconds [00-59]
	//
	// Pos. 18-20. Longitude Thousandths Of Seconds [000-999]
	//
	// Pos. 21. Longitude Hemisphere [EW]
	//
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]].
	Coord param.Opt[string] `json:"coord,omitzero"`
	// A mathematical model of the earth used to calculate coordinates on a map. US
	// Forces use the World Geodetic System 1984 (WGS 84), but also use maps by allied
	// countries with local datums. The datum must be specified to ensure accuracy of
	// coordinates. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	CoordDatum param.Opt[string] `json:"coordDatum,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// coordinate.
	CoordDerivAcc param.Opt[float64] `json:"coordDerivAcc,omitzero"`
	// Ground elevation, in meters, of the geographic coordinates referenced to (above
	// or below) Mean Sea Level (MSL) vertical datum.
	ElevMsl param.Opt[float64] `json:"elevMsl,omitzero"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy. Confidence level has a range of values
	// from 0 to 100, with 100 being highest level of confidence.
	ElevMslConfLvl param.Opt[int64] `json:"elevMslConfLvl,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation, measured in meters.
	ElevMslDerivAcc param.Opt[float64] `json:"elevMslDerivAcc,omitzero"`
	// Designated equipment code assigned to the item of equipment or an abbreviation
	// record type unique identifier. Users should consult the data provider for
	// information on the equipment code structure.
	EqpCode param.Opt[string] `json:"eqpCode,omitzero"`
	// Uniquely identifies each item or group of equipment associated with a unit,
	// facility or site.
	EqpIDNum param.Opt[string] `json:"eqpIdNum,omitzero"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval param.Opt[int64] `json:"eval,omitzero"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa param.Opt[string] `json:"fpa,omitzero"`
	// Indicates the function or mission of this equipment, which may or may not be
	// engaged in at any particular time. Typically refers to a unit, organization, or
	// installation/facility performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Function param.Opt[string] `json:"function,omitzero"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary param.Opt[string] `json:"functPrimary,omitzero"`
	// The distance between Mean Sea Level and a referenced ellipsoid, measured in
	// meters.
	GeoidalMslSep param.Opt[float64] `json:"geoidalMslSep,omitzero"`
	// Estimated identity of the Site (ASSUMED FRIEND, FRIEND, HOSTILE, FAKER, JOKER,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	Ident param.Opt[string] `json:"ident,omitzero"`
	// Unique identifier of the operating unit associated with the equipment record.
	IDOperatingUnit param.Opt[string] `json:"idOperatingUnit,omitzero"`
	// Unique identifier of the Parent equipment record associated with this equipment
	// record.
	IDParentEquipment param.Opt[string] `json:"idParentEquipment,omitzero"`
	// Unique identifier of the Site Entity associated with the equipment record.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Indicates the reason that the equipment is at that location. The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason param.Opt[string] `json:"locReason,omitzero"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts:
	//
	// 4Q (grid zone designator, GZD)
	//
	// FJ (the 100,000-meter square identifier)
	//
	// 12345678 (numerical location; easting is 1234 and northing is 5678, in this case
	// specifying a location with 10 m resolution).
	MilGrid param.Opt[string] `json:"milGrid,omitzero"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator.
	MilGridSys param.Opt[string] `json:"milGridSys,omitzero"`
	// Generic type this specific piece of equipment belongs to, and the identifying
	// nomenclature which describes the equipment.
	Nomen param.Opt[string] `json:"nomen,omitzero"`
	// Internationally recognized water area in which the vessel is most likely to be
	// deployed or in which it normally operates most frequently.
	OperAreaPrimary param.Opt[string] `json:"operAreaPrimary,omitzero"`
	// The Degree to which an entity is ready to perform the overall operational
	// mission(s) for which it was organized and equipped. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	OperStatus param.Opt[string] `json:"operStatus,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv param.Opt[string] `json:"polSubdiv,omitzero"`
	// Relative to the parent entity, the total number of military personnel or
	// equipment assessed to be on-hand (OH).
	QtyOh param.Opt[int64] `json:"qtyOH,omitzero"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Referenced, but not
	// constrained to, the following record status type classifications.
	//
	// # A - Active
	//
	// # I - Inactive
	//
	// # K - Acknowledged
	//
	// # L - Local
	//
	// # Q - A nominated (NOM) or Data Change Request (DCR) record
	//
	// # R - Production reduced by CMD decision
	//
	// W - Working Record.
	RecStatus param.Opt[string] `json:"recStatus,omitzero"`
	// The reference documentiation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc param.Opt[string] `json:"referenceDoc,omitzero"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd param.Opt[string] `json:"resProd,omitzero"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency. This date cannot be greater than the current
	// date.
	ReviewDate param.Opt[time.Time] `json:"reviewDate,omitzero" format:"date"`
	// Provider specific sequential number assigned to the equipment.
	SeqNum param.Opt[int64] `json:"seqNum,omitzero"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode param.Opt[string] `json:"symCode,omitzero"`
	// Universal Transverse Mercator (UTM) grid coordinates.
	//
	// Pos. 1-2, UTM Zone Column [01-60
	//
	// Pos. 3, UTM Zone Row [C-HJ-NP-X]
	//
	// Pos. 4, UTM False Easting [0-9]
	//
	// Pos. 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9]
	//
	// Pos. 10-11, UTM False Northing [0-9][0-9]
	//
	// Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm param.Opt[string] `json:"utm,omitzero"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	// Array of UUID(s) of the UDL data record(s) that are related to this equipment
	// record. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types such as AIRCRAFT, VESSEL, EO, MTI that are related to
	// this equipment record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	paramObj
}

func (r EquipmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EquipmentNewParams
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
type EquipmentNewParamsDataMode string

const (
	EquipmentNewParamsDataModeReal      EquipmentNewParamsDataMode = "REAL"
	EquipmentNewParamsDataModeTest      EquipmentNewParamsDataMode = "TEST"
	EquipmentNewParamsDataModeSimulated EquipmentNewParamsDataMode = "SIMULATED"
	EquipmentNewParamsDataModeExercise  EquipmentNewParamsDataMode = "EXERCISE"
)

type EquipmentGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EquipmentGetParams]'s query parameters as `url.Values`.
func (r EquipmentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EquipmentUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment geographic coordinates reside. This field will be set to
	// "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode,required"`
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
	DataMode EquipmentUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea param.Opt[string] `json:"airDefArea,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance param.Opt[string] `json:"allegiance,omitzero"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance param.Opt[string] `json:"altAllegiance,omitzero"`
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS. This field will be set
	// to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Unique identifier of the Equipment record from the originating system.
	AltEqpID param.Opt[string] `json:"altEqpId,omitzero"`
	// Indicates the importance of the equipment. Referenced, but not constrained to,
	// the following class ratings type classifications.
	//
	// 0 - Not of significant importance of the system
	//
	// 1 - Primary importance to system
	//
	// 2 - Secondary importance to system
	//
	// 3 - Tertiary importance to system
	//
	// O - Other. Explain in Remarks.
	ClassRating param.Opt[string] `json:"classRating,omitzero"`
	// The physical manner of being or state of existence of the entity. A physical
	// condition that must be considered in the determining of a course of action. The
	// specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Condition param.Opt[string] `json:"condition,omitzero"`
	// Availability of the entity relative to its condition. Indicates the reason the
	// entity is not fully operational. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	ConditionAvail param.Opt[string] `json:"conditionAvail,omitzero"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc. specified in degrees, minute,
	// and seconds.
	//
	// Pos. 1-2. Latitude Degrees [00-90]
	//
	// Pos. 3-4. Latitude Minutes [00-59]
	//
	// Pos. 5-6. Latitude Seconds [00-59]
	//
	// Pos. 7-9. Latitude Thousandths Of Seconds [000-999]
	//
	// Pos. 10. Latitude Hemisphere [NS]
	//
	// Pos. 11-13. Longitude Degrees [00-180]
	//
	// Pos. 14-15. Longitude Minutes [00-59]
	//
	// Pos. 16-17. Longitude Seconds [00-59]
	//
	// Pos. 18-20. Longitude Thousandths Of Seconds [000-999]
	//
	// Pos. 21. Longitude Hemisphere [EW]
	//
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]].
	Coord param.Opt[string] `json:"coord,omitzero"`
	// A mathematical model of the earth used to calculate coordinates on a map. US
	// Forces use the World Geodetic System 1984 (WGS 84), but also use maps by allied
	// countries with local datums. The datum must be specified to ensure accuracy of
	// coordinates. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	CoordDatum param.Opt[string] `json:"coordDatum,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// coordinate.
	CoordDerivAcc param.Opt[float64] `json:"coordDerivAcc,omitzero"`
	// Ground elevation, in meters, of the geographic coordinates referenced to (above
	// or below) Mean Sea Level (MSL) vertical datum.
	ElevMsl param.Opt[float64] `json:"elevMsl,omitzero"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy. Confidence level has a range of values
	// from 0 to 100, with 100 being highest level of confidence.
	ElevMslConfLvl param.Opt[int64] `json:"elevMslConfLvl,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation, measured in meters.
	ElevMslDerivAcc param.Opt[float64] `json:"elevMslDerivAcc,omitzero"`
	// Designated equipment code assigned to the item of equipment or an abbreviation
	// record type unique identifier. Users should consult the data provider for
	// information on the equipment code structure.
	EqpCode param.Opt[string] `json:"eqpCode,omitzero"`
	// Uniquely identifies each item or group of equipment associated with a unit,
	// facility or site.
	EqpIDNum param.Opt[string] `json:"eqpIdNum,omitzero"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval param.Opt[int64] `json:"eval,omitzero"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa param.Opt[string] `json:"fpa,omitzero"`
	// Indicates the function or mission of this equipment, which may or may not be
	// engaged in at any particular time. Typically refers to a unit, organization, or
	// installation/facility performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Function param.Opt[string] `json:"function,omitzero"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary param.Opt[string] `json:"functPrimary,omitzero"`
	// The distance between Mean Sea Level and a referenced ellipsoid, measured in
	// meters.
	GeoidalMslSep param.Opt[float64] `json:"geoidalMslSep,omitzero"`
	// Estimated identity of the Site (ASSUMED FRIEND, FRIEND, HOSTILE, FAKER, JOKER,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	Ident param.Opt[string] `json:"ident,omitzero"`
	// Unique identifier of the operating unit associated with the equipment record.
	IDOperatingUnit param.Opt[string] `json:"idOperatingUnit,omitzero"`
	// Unique identifier of the Parent equipment record associated with this equipment
	// record.
	IDParentEquipment param.Opt[string] `json:"idParentEquipment,omitzero"`
	// Unique identifier of the Site Entity associated with the equipment record.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Indicates the reason that the equipment is at that location. The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason param.Opt[string] `json:"locReason,omitzero"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts:
	//
	// 4Q (grid zone designator, GZD)
	//
	// FJ (the 100,000-meter square identifier)
	//
	// 12345678 (numerical location; easting is 1234 and northing is 5678, in this case
	// specifying a location with 10 m resolution).
	MilGrid param.Opt[string] `json:"milGrid,omitzero"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator.
	MilGridSys param.Opt[string] `json:"milGridSys,omitzero"`
	// Generic type this specific piece of equipment belongs to, and the identifying
	// nomenclature which describes the equipment.
	Nomen param.Opt[string] `json:"nomen,omitzero"`
	// Internationally recognized water area in which the vessel is most likely to be
	// deployed or in which it normally operates most frequently.
	OperAreaPrimary param.Opt[string] `json:"operAreaPrimary,omitzero"`
	// The Degree to which an entity is ready to perform the overall operational
	// mission(s) for which it was organized and equipped. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	OperStatus param.Opt[string] `json:"operStatus,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv param.Opt[string] `json:"polSubdiv,omitzero"`
	// Relative to the parent entity, the total number of military personnel or
	// equipment assessed to be on-hand (OH).
	QtyOh param.Opt[int64] `json:"qtyOH,omitzero"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Referenced, but not
	// constrained to, the following record status type classifications.
	//
	// # A - Active
	//
	// # I - Inactive
	//
	// # K - Acknowledged
	//
	// # L - Local
	//
	// # Q - A nominated (NOM) or Data Change Request (DCR) record
	//
	// # R - Production reduced by CMD decision
	//
	// W - Working Record.
	RecStatus param.Opt[string] `json:"recStatus,omitzero"`
	// The reference documentiation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc param.Opt[string] `json:"referenceDoc,omitzero"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd param.Opt[string] `json:"resProd,omitzero"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency. This date cannot be greater than the current
	// date.
	ReviewDate param.Opt[time.Time] `json:"reviewDate,omitzero" format:"date"`
	// Provider specific sequential number assigned to the equipment.
	SeqNum param.Opt[int64] `json:"seqNum,omitzero"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode param.Opt[string] `json:"symCode,omitzero"`
	// Universal Transverse Mercator (UTM) grid coordinates.
	//
	// Pos. 1-2, UTM Zone Column [01-60
	//
	// Pos. 3, UTM Zone Row [C-HJ-NP-X]
	//
	// Pos. 4, UTM False Easting [0-9]
	//
	// Pos. 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9]
	//
	// Pos. 10-11, UTM False Northing [0-9][0-9]
	//
	// Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm param.Opt[string] `json:"utm,omitzero"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	// Array of UUID(s) of the UDL data record(s) that are related to this equipment
	// record. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types such as AIRCRAFT, VESSEL, EO, MTI that are related to
	// this equipment record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	paramObj
}

func (r EquipmentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EquipmentUpdateParams
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
type EquipmentUpdateParamsDataMode string

const (
	EquipmentUpdateParamsDataModeReal      EquipmentUpdateParamsDataMode = "REAL"
	EquipmentUpdateParamsDataModeTest      EquipmentUpdateParamsDataMode = "TEST"
	EquipmentUpdateParamsDataModeSimulated EquipmentUpdateParamsDataMode = "SIMULATED"
	EquipmentUpdateParamsDataModeExercise  EquipmentUpdateParamsDataMode = "EXERCISE"
)

type EquipmentListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EquipmentListParams]'s query parameters as `url.Values`.
func (r EquipmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EquipmentCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EquipmentCountParams]'s query parameters as `url.Values`.
func (r EquipmentCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EquipmentNewBulkParams struct {
	Body []EquipmentNewBulkParamsBody
	paramObj
}

func (r EquipmentNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Properties and characteristics of equipment that can be associated with a site
// or other entity.
//
// The properties ClassificationMarking, CountryCode, DataMode, Lat, Lon, Source
// are required.
type EquipmentNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment geographic coordinates reside. This field will be set to
	// "OTHR" if the source value does not match a UDL Country code value
	// (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode,required"`
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
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat,required"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea param.Opt[string] `json:"airDefArea,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the equipment owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance param.Opt[string] `json:"allegiance,omitzero"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance param.Opt[string] `json:"altAllegiance,omitzero"`
	// Specifies an alternate country code if the data provider code is not part of an
	// official Country Code standard such as ISO-3166 or FIPS. This field will be set
	// to the value provided by the source and should be used for all Queries
	// specifying a Country Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Unique identifier of the Equipment record from the originating system.
	AltEqpID param.Opt[string] `json:"altEqpId,omitzero"`
	// Indicates the importance of the equipment. Referenced, but not constrained to,
	// the following class ratings type classifications.
	//
	// 0 - Not of significant importance of the system
	//
	// 1 - Primary importance to system
	//
	// 2 - Secondary importance to system
	//
	// 3 - Tertiary importance to system
	//
	// O - Other. Explain in Remarks.
	ClassRating param.Opt[string] `json:"classRating,omitzero"`
	// The physical manner of being or state of existence of the entity. A physical
	// condition that must be considered in the determining of a course of action. The
	// specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Condition param.Opt[string] `json:"condition,omitzero"`
	// Availability of the entity relative to its condition. Indicates the reason the
	// entity is not fully operational. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	ConditionAvail param.Opt[string] `json:"conditionAvail,omitzero"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc. specified in degrees, minute,
	// and seconds.
	//
	// Pos. 1-2. Latitude Degrees [00-90]
	//
	// Pos. 3-4. Latitude Minutes [00-59]
	//
	// Pos. 5-6. Latitude Seconds [00-59]
	//
	// Pos. 7-9. Latitude Thousandths Of Seconds [000-999]
	//
	// Pos. 10. Latitude Hemisphere [NS]
	//
	// Pos. 11-13. Longitude Degrees [00-180]
	//
	// Pos. 14-15. Longitude Minutes [00-59]
	//
	// Pos. 16-17. Longitude Seconds [00-59]
	//
	// Pos. 18-20. Longitude Thousandths Of Seconds [000-999]
	//
	// Pos. 21. Longitude Hemisphere [EW]
	//
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]].
	Coord param.Opt[string] `json:"coord,omitzero"`
	// A mathematical model of the earth used to calculate coordinates on a map. US
	// Forces use the World Geodetic System 1984 (WGS 84), but also use maps by allied
	// countries with local datums. The datum must be specified to ensure accuracy of
	// coordinates. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	CoordDatum param.Opt[string] `json:"coordDatum,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// coordinate.
	CoordDerivAcc param.Opt[float64] `json:"coordDerivAcc,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Ground elevation, in meters, of the geographic coordinates referenced to (above
	// or below) Mean Sea Level (MSL) vertical datum.
	ElevMsl param.Opt[float64] `json:"elevMsl,omitzero"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy. Confidence level has a range of values
	// from 0 to 100, with 100 being highest level of confidence.
	ElevMslConfLvl param.Opt[int64] `json:"elevMslConfLvl,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation, measured in meters.
	ElevMslDerivAcc param.Opt[float64] `json:"elevMslDerivAcc,omitzero"`
	// Designated equipment code assigned to the item of equipment or an abbreviation
	// record type unique identifier. Users should consult the data provider for
	// information on the equipment code structure.
	EqpCode param.Opt[string] `json:"eqpCode,omitzero"`
	// Uniquely identifies each item or group of equipment associated with a unit,
	// facility or site.
	EqpIDNum param.Opt[string] `json:"eqpIdNum,omitzero"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval param.Opt[int64] `json:"eval,omitzero"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa param.Opt[string] `json:"fpa,omitzero"`
	// Indicates the function or mission of this equipment, which may or may not be
	// engaged in at any particular time. Typically refers to a unit, organization, or
	// installation/facility performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Function param.Opt[string] `json:"function,omitzero"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary param.Opt[string] `json:"functPrimary,omitzero"`
	// The distance between Mean Sea Level and a referenced ellipsoid, measured in
	// meters.
	GeoidalMslSep param.Opt[float64] `json:"geoidalMslSep,omitzero"`
	// Estimated identity of the Site (ASSUMED FRIEND, FRIEND, HOSTILE, FAKER, JOKER,
	// NEUTRAL, PENDING, SUSPECT, UNKNOWN):
	//
	// ASSUMED FRIEND: Track assumed to be a friend due to the object characteristics,
	// behavior, and/or origin.
	//
	// FRIEND: Track object supporting friendly forces and belonging to a declared
	// friendly nation or entity.
	//
	// HOSTILE: Track object belonging to an opposing nation, party, group, or entity
	// deemed to contribute to a threat to friendly forces or their mission due to its
	// behavior, characteristics, nationality, or origin.
	//
	// FAKER: Friendly track, object, or entity acting as an exercise hostile.
	//
	// JOKER: Friendly track, object, or entity acting as an exercise suspect.
	//
	// NEUTRAL: Track object whose characteristics, behavior, nationality, and/or
	// origin indicate that it is neither supporting nor opposing friendly forces or
	// their mission.
	//
	// PENDING: Track object which has not been evaluated.
	//
	// SUSPECT: Track object deemed potentially hostile due to the object
	// characteristics, behavior, nationality, and/or origin.
	//
	// UNKNOWN: Track object which has been evaluated and does not meet criteria for
	// any standard identity.
	Ident param.Opt[string] `json:"ident,omitzero"`
	// Unique identifier of the operating unit associated with the equipment record.
	IDOperatingUnit param.Opt[string] `json:"idOperatingUnit,omitzero"`
	// Unique identifier of the Parent equipment record associated with this equipment
	// record.
	IDParentEquipment param.Opt[string] `json:"idParentEquipment,omitzero"`
	// Unique identifier of the Site Entity associated with the equipment record.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// Indicates the reason that the equipment is at that location. The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason param.Opt[string] `json:"locReason,omitzero"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts:
	//
	// 4Q (grid zone designator, GZD)
	//
	// FJ (the 100,000-meter square identifier)
	//
	// 12345678 (numerical location; easting is 1234 and northing is 5678, in this case
	// specifying a location with 10 m resolution).
	MilGrid param.Opt[string] `json:"milGrid,omitzero"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator.
	MilGridSys param.Opt[string] `json:"milGridSys,omitzero"`
	// Generic type this specific piece of equipment belongs to, and the identifying
	// nomenclature which describes the equipment.
	Nomen param.Opt[string] `json:"nomen,omitzero"`
	// Internationally recognized water area in which the vessel is most likely to be
	// deployed or in which it normally operates most frequently.
	OperAreaPrimary param.Opt[string] `json:"operAreaPrimary,omitzero"`
	// The Degree to which an entity is ready to perform the overall operational
	// mission(s) for which it was organized and equipped. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	OperStatus param.Opt[string] `json:"operStatus,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv param.Opt[string] `json:"polSubdiv,omitzero"`
	// Relative to the parent entity, the total number of military personnel or
	// equipment assessed to be on-hand (OH).
	QtyOh param.Opt[int64] `json:"qtyOH,omitzero"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Referenced, but not
	// constrained to, the following record status type classifications.
	//
	// # A - Active
	//
	// # I - Inactive
	//
	// # K - Acknowledged
	//
	// # L - Local
	//
	// # Q - A nominated (NOM) or Data Change Request (DCR) record
	//
	// # R - Production reduced by CMD decision
	//
	// W - Working Record.
	RecStatus param.Opt[string] `json:"recStatus,omitzero"`
	// The reference documentiation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc param.Opt[string] `json:"referenceDoc,omitzero"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd param.Opt[string] `json:"resProd,omitzero"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency. This date cannot be greater than the current
	// date.
	ReviewDate param.Opt[time.Time] `json:"reviewDate,omitzero" format:"date"`
	// Provider specific sequential number assigned to the equipment.
	SeqNum param.Opt[int64] `json:"seqNum,omitzero"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode param.Opt[string] `json:"symCode,omitzero"`
	// Universal Transverse Mercator (UTM) grid coordinates.
	//
	// Pos. 1-2, UTM Zone Column [01-60
	//
	// Pos. 3, UTM Zone Row [C-HJ-NP-X]
	//
	// Pos. 4, UTM False Easting [0-9]
	//
	// Pos. 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9]
	//
	// Pos. 10-11, UTM False Northing [0-9][0-9]
	//
	// Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm param.Opt[string] `json:"utm,omitzero"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	// Array of UUID(s) of the UDL data record(s) that are related to this equipment
	// record. See the associated 'srcTyps' array for the specific types of data,
	// positionally corresponding to the UUIDs in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size. See the corresponding srcTyps array element
	// for the data type of the UUID and use the appropriate API operation to retrieve
	// that object.
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types such as AIRCRAFT, VESSEL, EO, MTI that are related to
	// this equipment record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	paramObj
}

func (r EquipmentNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EquipmentNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EquipmentNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type EquipmentTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EquipmentTupleParams]'s query parameters as `url.Values`.
func (r EquipmentTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
