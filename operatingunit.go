// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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

// OperatingunitService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperatingunitService] method instead.
type OperatingunitService struct {
	Options []option.RequestOption
}

// NewOperatingunitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperatingunitService(opts ...option.RequestOption) (r OperatingunitService) {
	r = OperatingunitService{}
	r.Options = opts
	return
}

// Service operation to take a single Operatingunit as a POST body and ingest into
// the database. Operatingunit defines a unit or organization which operates or
// controls a space-related Entity. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *OperatingunitService) New(ctx context.Context, body OperatingunitNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/operatingunit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Operatingunit. Operatingunit defines a unit
// or organization which operates or controls a space-related Entity. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *OperatingunitService) Update(ctx context.Context, id string, body OperatingunitUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/operatingunit/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OperatingunitService) List(ctx context.Context, query OperatingunitListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OperatingunitListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/operatingunit"
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
func (r *OperatingunitService) ListAutoPaging(ctx context.Context, query OperatingunitListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OperatingunitListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an Operatingunit object specified by the passed ID
// path parameter. Operatingunit defines a unit or organization which operates or
// controls a space-related Entity. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *OperatingunitService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/operatingunit/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OperatingunitService) Count(ctx context.Context, query OperatingunitCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/operatingunit/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Operatingunit record by its unique ID passed
// as a path parameter. Operatingunit defines a unit or organization which operates
// or controls a space-related Entity.
func (r *OperatingunitService) Get(ctx context.Context, id string, query OperatingunitGetParams, opts ...option.RequestOption) (res *OperatingunitGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/operatingunit/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *OperatingunitService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/operatingunit/queryhelp"
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
func (r *OperatingunitService) Tuple(ctx context.Context, query OperatingunitTupleParams, opts ...option.RequestOption) (res *[]OperatingunitTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/operatingunit/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of a unit or organization which operates or controls a
// space-related Entity such as an on-orbit payload, a sensor, etc. A contact may
// belong to an organization.
type OperatingunitListResponse struct {
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
	DataMode OperatingunitListResponseDataMode `json:"dataMode,required"`
	// Name of the operating unit.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea string `json:"airDefArea"`
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit owes its allegiance. This field will be set to "OTHR"
	// if the source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	// Unique identifier of the operating unit record from the originating system.
	AltOperatingUnitID string `json:"altOperatingUnitId"`
	// Indicates the importance of the operating unit to the OES or MIR system. This
	// data element is restricted to update by DIA (DB-4). Valid values are: 0 - Does
	// not meet criteria above 1 - Primary importance to system 2 - Secondary
	// importance to system 3 - Tertiary importance to system O - Other. Explain in
	// Remarks.
	ClassRating string `json:"classRating"`
	// The physical manner of being or state of existence of the operating unit. A
	// physical condition that must be considered in the determining of a course of
	// action. The specific usage and enumerations contained in this field may be found
	// in the documentation provided in the referenceDoc field. If referenceDoc not
	// provided, users may consult the data provider.
	Condition string `json:"condition"`
	// Availability of the operating unit relative to its condition. Indicates the
	// reason the operating unit is not fully operational. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	ConditionAvail string `json:"conditionAvail"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc.
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]
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
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit geographic coordinates reside . This field will be set
	// to "OTHR" if the source value does not match a UDL country code value
	// (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A code describing the amount of operating unit participation in a deployment.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	DeployStatus string `json:"deployStatus"`
	// Description of the operating unit.
	Description string `json:"description"`
	// Combat status of a divisional or equivalent operating unit. Currently, this data
	// element applies only to operating units of the Former Soviet Union. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	DivCat string `json:"divCat"`
	// Organizational level of the operating unit. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Echelon string `json:"echelon"`
	// Indicates the major group or level to which an echelon belongs. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	EchelonTier string `json:"echelonTier"`
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum.
	ElevMsl float64 `json:"elevMsl"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl int64 `json:"elevMslConfLvl"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc float64 `json:"elevMslDerivAcc"`
	// The Intelligence Confidence Level or the Reliability/degree of confidence that
	// the analyst has assigned to the data within this record. The numerical range is
	// from 1 to 9 with 1 representing the highest confidence level.
	Eval int64 `json:"eval"`
	// The country code of the observed flag flown.
	FlagFlown string `json:"flagFlown"`
	// Naval fleet to which an operating unit is assigned. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FleetID string `json:"fleetId"`
	// An aggregation of military units within a single service (i.e., ARMY, AIR FORCE,
	// etc.) which operates under a single authority to accomplish a common mission.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Force string `json:"force"`
	// The specific name for a given force. For example, Force = ADF (Air Defense
	// Force) and Force Name = Army Air Defense Force.
	ForceName string `json:"forceName"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa string `json:"fpa"`
	// Principal combat-related role that an operating unit is organized, structured
	// and equipped to perform. Or, the specialized military or paramilitary branch in
	// which an individual serves, their specialization. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctRole string `json:"functRole"`
	// The distance between Mean Sea Level and a referenced ellipsoid.
	GeoidalMslSep float64 `json:"geoidalMslSep"`
	// Unique identifier of the contact for this operating unit.
	IDContact string `json:"idContact"`
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
	// Unique identifier of the location record for this operating unit.
	IDLocation string `json:"idLocation"`
	// Unique identifier of the record, auto-generated by the system.
	IDOperatingUnit string `json:"idOperatingUnit"`
	// Unique identifier of the organization record for this operating unit.
	IDOrganization string `json:"idOrganization"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// Location name for the coordinates.
	LocName string `json:"locName"`
	// Indicates the reason that the operating unit is at that location. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason string `json:"locReason"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// This field contains a value indicating whether the record is a master unit
	// record (True) or a detail record (False). Master records contain basic
	// information that does not change over time for each unit that has been selected
	// to be projected.
	MasterUnit bool `json:"masterUnit"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts: 4Q (grid zone
	// designator, GZD) FJ (the 100,000-meter square identifier) 12345678 (numerical
	// location; easting is 1234 and northing is 5678, in this case specifying a
	// location with 10 m resolution).
	MilGrid string `json:"milGrid"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator
	MilGridSys string `json:"milGridSys"`
	// Indicates the principal type of mission that an operating unit is organized and
	// equipped to perform. The specific usage and enumerations contained in this field
	// may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimary string `json:"msnPrimary"`
	// Indicates the principal specialty type of mission that an operating unit is
	// organized and equipped to perform. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpecialty string `json:"msnPrimarySpecialty"`
	// The Degree to which an operating unit is ready to perform the overall
	// operational mission(s) for which it was organized and equipped. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
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
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Values are: A - Active I -
	// Inactive K - Acknowledged L - Local Q - A nominated (NOM) or Data Change Request
	// (DCR) record R - Production reduced by CMD decision W - Working Record.
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
	// This field contains a value indicating whether the record is a stylized
	// operating unit record (True) or a regular operating unit record (False). A
	// stylized operating unit is a type of operating unit with one set of equipment
	// that can be assigned to one or more superiors. A stylized operating unit is
	// generally useful for lower echelon operating units where the number of operating
	// units and types of equipment are equal for multiple organizations. In lieu of
	// creating unique operating unit records for each operating unit, a template is
	// created for the operating unit and its equipment. This template enables the user
	// to assign the operating unit to multiple organizations.
	StylizedUnit bool `json:"stylizedUnit"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode string `json:"symCode"`
	// An optional identifier for this operating unit that may be composed from items
	// such as the originating organization, allegiance, one-up number, etc.
	UnitIdentifier string `json:"unitIdentifier"`
	// Universal Transverse Mercator (UTM) grid coordinates. Pos. 1-2, UTM Zone Column
	// [01-60 Pos. 3, UTM Zone Row [C-HJ-NP-X] Pos. 4, UTM False Easting [0-9] Pos.
	// 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9] Pos. 10-11, UTM False Northing
	// [0-9][0-9] Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm string `json:"utm"`
	// World Aeronautical Chart identifier for the area in which a designated operating
	// unit is located.
	Wac string `json:"wac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		AirDefArea            respjson.Field
		Allegiance            respjson.Field
		AltAllegiance         respjson.Field
		AltCountryCode        respjson.Field
		AltOperatingUnitID    respjson.Field
		ClassRating           respjson.Field
		Condition             respjson.Field
		ConditionAvail        respjson.Field
		Coord                 respjson.Field
		CoordDatum            respjson.Field
		CoordDerivAcc         respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeployStatus          respjson.Field
		Description           respjson.Field
		DivCat                respjson.Field
		Echelon               respjson.Field
		EchelonTier           respjson.Field
		ElevMsl               respjson.Field
		ElevMslConfLvl        respjson.Field
		ElevMslDerivAcc       respjson.Field
		Eval                  respjson.Field
		FlagFlown             respjson.Field
		FleetID               respjson.Field
		Force                 respjson.Field
		ForceName             respjson.Field
		Fpa                   respjson.Field
		FunctRole             respjson.Field
		GeoidalMslSep         respjson.Field
		IDContact             respjson.Field
		Ident                 respjson.Field
		IDLocation            respjson.Field
		IDOperatingUnit       respjson.Field
		IDOrganization        respjson.Field
		Lat                   respjson.Field
		LocName               respjson.Field
		LocReason             respjson.Field
		Lon                   respjson.Field
		MasterUnit            respjson.Field
		MilGrid               respjson.Field
		MilGridSys            respjson.Field
		MsnPrimary            respjson.Field
		MsnPrimarySpecialty   respjson.Field
		OperStatus            respjson.Field
		Origin                respjson.Field
		PolSubdiv             respjson.Field
		RecStatus             respjson.Field
		ReferenceDoc          respjson.Field
		ResProd               respjson.Field
		ReviewDate            respjson.Field
		StylizedUnit          respjson.Field
		SymCode               respjson.Field
		UnitIdentifier        respjson.Field
		Utm                   respjson.Field
		Wac                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperatingunitListResponse) RawJSON() string { return r.JSON.raw }
func (r *OperatingunitListResponse) UnmarshalJSON(data []byte) error {
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
type OperatingunitListResponseDataMode string

const (
	OperatingunitListResponseDataModeReal      OperatingunitListResponseDataMode = "REAL"
	OperatingunitListResponseDataModeTest      OperatingunitListResponseDataMode = "TEST"
	OperatingunitListResponseDataModeSimulated OperatingunitListResponseDataMode = "SIMULATED"
	OperatingunitListResponseDataModeExercise  OperatingunitListResponseDataMode = "EXERCISE"
)

// Model representation of a unit or organization which operates or controls a
// space-related Entity such as an on-orbit payload, a sensor, etc. A contact may
// belong to an organization.
type OperatingunitGetResponse struct {
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
	DataMode OperatingunitGetResponseDataMode `json:"dataMode,required"`
	// Name of the operating unit.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea string `json:"airDefArea"`
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit owes its allegiance. This field will be set to "OTHR"
	// if the source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	// Unique identifier of the operating unit record from the originating system.
	AltOperatingUnitID string `json:"altOperatingUnitId"`
	// Indicates the importance of the operating unit to the OES or MIR system. This
	// data element is restricted to update by DIA (DB-4). Valid values are: 0 - Does
	// not meet criteria above 1 - Primary importance to system 2 - Secondary
	// importance to system 3 - Tertiary importance to system O - Other. Explain in
	// Remarks.
	ClassRating string `json:"classRating"`
	// The physical manner of being or state of existence of the operating unit. A
	// physical condition that must be considered in the determining of a course of
	// action. The specific usage and enumerations contained in this field may be found
	// in the documentation provided in the referenceDoc field. If referenceDoc not
	// provided, users may consult the data provider.
	Condition string `json:"condition"`
	// Availability of the operating unit relative to its condition. Indicates the
	// reason the operating unit is not fully operational. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	ConditionAvail string `json:"conditionAvail"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc.
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]
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
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit geographic coordinates reside . This field will be set
	// to "OTHR" if the source value does not match a UDL country code value
	// (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A code describing the amount of operating unit participation in a deployment.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	DeployStatus string `json:"deployStatus"`
	// Description of the operating unit.
	Description string `json:"description"`
	// Combat status of a divisional or equivalent operating unit. Currently, this data
	// element applies only to operating units of the Former Soviet Union. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	DivCat string `json:"divCat"`
	// Organizational level of the operating unit. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Echelon string `json:"echelon"`
	// Indicates the major group or level to which an echelon belongs. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	EchelonTier string `json:"echelonTier"`
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum.
	ElevMsl float64 `json:"elevMsl"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl int64 `json:"elevMslConfLvl"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc float64 `json:"elevMslDerivAcc"`
	// The Intelligence Confidence Level or the Reliability/degree of confidence that
	// the analyst has assigned to the data within this record. The numerical range is
	// from 1 to 9 with 1 representing the highest confidence level.
	Eval int64 `json:"eval"`
	// The country code of the observed flag flown.
	FlagFlown string `json:"flagFlown"`
	// Naval fleet to which an operating unit is assigned. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FleetID string `json:"fleetId"`
	// An aggregation of military units within a single service (i.e., ARMY, AIR FORCE,
	// etc.) which operates under a single authority to accomplish a common mission.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Force string `json:"force"`
	// The specific name for a given force. For example, Force = ADF (Air Defense
	// Force) and Force Name = Army Air Defense Force.
	ForceName string `json:"forceName"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa string `json:"fpa"`
	// Principal combat-related role that an operating unit is organized, structured
	// and equipped to perform. Or, the specialized military or paramilitary branch in
	// which an individual serves, their specialization. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctRole string `json:"functRole"`
	// The distance between Mean Sea Level and a referenced ellipsoid.
	GeoidalMslSep float64 `json:"geoidalMslSep"`
	// Unique identifier of the contact for this operating unit.
	IDContact string `json:"idContact"`
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
	// Unique identifier of the location record for this operating unit.
	IDLocation string `json:"idLocation"`
	// Unique identifier of the record, auto-generated by the system.
	IDOperatingUnit string `json:"idOperatingUnit"`
	// Unique identifier of the organization record for this operating unit.
	IDOrganization string `json:"idOrganization"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LocationFull `json:"location"`
	// Location name for the coordinates.
	LocName string `json:"locName"`
	// Indicates the reason that the operating unit is at that location. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason string `json:"locReason"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// This field contains a value indicating whether the record is a master unit
	// record (True) or a detail record (False). Master records contain basic
	// information that does not change over time for each unit that has been selected
	// to be projected.
	MasterUnit bool `json:"masterUnit"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts: 4Q (grid zone
	// designator, GZD) FJ (the 100,000-meter square identifier) 12345678 (numerical
	// location; easting is 1234 and northing is 5678, in this case specifying a
	// location with 10 m resolution).
	MilGrid string `json:"milGrid"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator
	MilGridSys string `json:"milGridSys"`
	// Indicates the principal type of mission that an operating unit is organized and
	// equipped to perform. The specific usage and enumerations contained in this field
	// may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimary string `json:"msnPrimary"`
	// Indicates the principal specialty type of mission that an operating unit is
	// organized and equipped to perform. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpecialty string `json:"msnPrimarySpecialty"`
	// Remarks contain amplifying information for a specific service. The information
	// may contain context and interpretations for consumer use.
	OperatingUnitRemarks []OperatingunitGetResponseOperatingUnitRemark `json:"operatingUnitRemarks"`
	// The Degree to which an operating unit is ready to perform the overall
	// operational mission(s) for which it was organized and equipped. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	OperStatus string `json:"operStatus"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	Organization OrganizationFull `json:"organization"`
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
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Values are: A - Active I -
	// Inactive K - Acknowledged L - Local Q - A nominated (NOM) or Data Change Request
	// (DCR) record R - Production reduced by CMD decision W - Working Record.
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
	// This field contains a value indicating whether the record is a stylized
	// operating unit record (True) or a regular operating unit record (False). A
	// stylized operating unit is a type of operating unit with one set of equipment
	// that can be assigned to one or more superiors. A stylized operating unit is
	// generally useful for lower echelon operating units where the number of operating
	// units and types of equipment are equal for multiple organizations. In lieu of
	// creating unique operating unit records for each operating unit, a template is
	// created for the operating unit and its equipment. This template enables the user
	// to assign the operating unit to multiple organizations.
	StylizedUnit bool `json:"stylizedUnit"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode string `json:"symCode"`
	// An optional identifier for this operating unit that may be composed from items
	// such as the originating organization, allegiance, one-up number, etc.
	UnitIdentifier string `json:"unitIdentifier"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Universal Transverse Mercator (UTM) grid coordinates. Pos. 1-2, UTM Zone Column
	// [01-60 Pos. 3, UTM Zone Row [C-HJ-NP-X] Pos. 4, UTM False Easting [0-9] Pos.
	// 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9] Pos. 10-11, UTM False Northing
	// [0-9][0-9] Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm string `json:"utm"`
	// World Aeronautical Chart identifier for the area in which a designated operating
	// unit is located.
	Wac string `json:"wac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		AirDefArea            respjson.Field
		Allegiance            respjson.Field
		AltAllegiance         respjson.Field
		AltCountryCode        respjson.Field
		AltOperatingUnitID    respjson.Field
		ClassRating           respjson.Field
		Condition             respjson.Field
		ConditionAvail        respjson.Field
		Coord                 respjson.Field
		CoordDatum            respjson.Field
		CoordDerivAcc         respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeployStatus          respjson.Field
		Description           respjson.Field
		DivCat                respjson.Field
		Echelon               respjson.Field
		EchelonTier           respjson.Field
		ElevMsl               respjson.Field
		ElevMslConfLvl        respjson.Field
		ElevMslDerivAcc       respjson.Field
		Eval                  respjson.Field
		FlagFlown             respjson.Field
		FleetID               respjson.Field
		Force                 respjson.Field
		ForceName             respjson.Field
		Fpa                   respjson.Field
		FunctRole             respjson.Field
		GeoidalMslSep         respjson.Field
		IDContact             respjson.Field
		Ident                 respjson.Field
		IDLocation            respjson.Field
		IDOperatingUnit       respjson.Field
		IDOrganization        respjson.Field
		Lat                   respjson.Field
		Location              respjson.Field
		LocName               respjson.Field
		LocReason             respjson.Field
		Lon                   respjson.Field
		MasterUnit            respjson.Field
		MilGrid               respjson.Field
		MilGridSys            respjson.Field
		MsnPrimary            respjson.Field
		MsnPrimarySpecialty   respjson.Field
		OperatingUnitRemarks  respjson.Field
		OperStatus            respjson.Field
		Organization          respjson.Field
		Origin                respjson.Field
		PolSubdiv             respjson.Field
		RecStatus             respjson.Field
		ReferenceDoc          respjson.Field
		ResProd               respjson.Field
		ReviewDate            respjson.Field
		StylizedUnit          respjson.Field
		SymCode               respjson.Field
		UnitIdentifier        respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Utm                   respjson.Field
		Wac                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperatingunitGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OperatingunitGetResponse) UnmarshalJSON(data []byte) error {
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
type OperatingunitGetResponseDataMode string

const (
	OperatingunitGetResponseDataModeReal      OperatingunitGetResponseDataMode = "REAL"
	OperatingunitGetResponseDataModeTest      OperatingunitGetResponseDataMode = "TEST"
	OperatingunitGetResponseDataModeSimulated OperatingunitGetResponseDataMode = "SIMULATED"
	OperatingunitGetResponseDataModeExercise  OperatingunitGetResponseDataMode = "EXERCISE"
)

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
type OperatingunitGetResponseOperatingUnitRemark struct {
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
	// The ID of the operating unit to which this remark applies.
	IDOperatingUnit string `json:"idOperatingUnit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Unique identifier of the unit remark record from the originating system.
	AltRmkID string `json:"altRmkId"`
	// The remark type identifier. For example, the Mobility Air Forces (MAF) remark
	// code, defined in the Airfield Suitability and Restriction Report (ASRR).
	Code string `json:"code"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The name of the remark.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The remark type (e.g. Caution, Information, Misc, Restriction, etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOperatingUnit       respjson.Field
		Source                respjson.Field
		Text                  respjson.Field
		ID                    respjson.Field
		AltRmkID              respjson.Field
		Code                  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Name                  respjson.Field
		Origin                respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperatingunitGetResponseOperatingUnitRemark) RawJSON() string { return r.JSON.raw }
func (r *OperatingunitGetResponseOperatingUnitRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representation of a unit or organization which operates or controls a
// space-related Entity such as an on-orbit payload, a sensor, etc. A contact may
// belong to an organization.
type OperatingunitTupleResponse struct {
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
	DataMode OperatingunitTupleResponseDataMode `json:"dataMode,required"`
	// Name of the operating unit.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea string `json:"airDefArea"`
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit owes its allegiance. This field will be set to "OTHR"
	// if the source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	// Unique identifier of the operating unit record from the originating system.
	AltOperatingUnitID string `json:"altOperatingUnitId"`
	// Indicates the importance of the operating unit to the OES or MIR system. This
	// data element is restricted to update by DIA (DB-4). Valid values are: 0 - Does
	// not meet criteria above 1 - Primary importance to system 2 - Secondary
	// importance to system 3 - Tertiary importance to system O - Other. Explain in
	// Remarks.
	ClassRating string `json:"classRating"`
	// The physical manner of being or state of existence of the operating unit. A
	// physical condition that must be considered in the determining of a course of
	// action. The specific usage and enumerations contained in this field may be found
	// in the documentation provided in the referenceDoc field. If referenceDoc not
	// provided, users may consult the data provider.
	Condition string `json:"condition"`
	// Availability of the operating unit relative to its condition. Indicates the
	// reason the operating unit is not fully operational. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	ConditionAvail string `json:"conditionAvail"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc.
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]
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
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit geographic coordinates reside . This field will be set
	// to "OTHR" if the source value does not match a UDL country code value
	// (ISO-3166-ALPHA-2).
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// A code describing the amount of operating unit participation in a deployment.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	DeployStatus string `json:"deployStatus"`
	// Description of the operating unit.
	Description string `json:"description"`
	// Combat status of a divisional or equivalent operating unit. Currently, this data
	// element applies only to operating units of the Former Soviet Union. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	DivCat string `json:"divCat"`
	// Organizational level of the operating unit. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Echelon string `json:"echelon"`
	// Indicates the major group or level to which an echelon belongs. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	EchelonTier string `json:"echelonTier"`
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum.
	ElevMsl float64 `json:"elevMsl"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl int64 `json:"elevMslConfLvl"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc float64 `json:"elevMslDerivAcc"`
	// The Intelligence Confidence Level or the Reliability/degree of confidence that
	// the analyst has assigned to the data within this record. The numerical range is
	// from 1 to 9 with 1 representing the highest confidence level.
	Eval int64 `json:"eval"`
	// The country code of the observed flag flown.
	FlagFlown string `json:"flagFlown"`
	// Naval fleet to which an operating unit is assigned. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FleetID string `json:"fleetId"`
	// An aggregation of military units within a single service (i.e., ARMY, AIR FORCE,
	// etc.) which operates under a single authority to accomplish a common mission.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Force string `json:"force"`
	// The specific name for a given force. For example, Force = ADF (Air Defense
	// Force) and Force Name = Army Air Defense Force.
	ForceName string `json:"forceName"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa string `json:"fpa"`
	// Principal combat-related role that an operating unit is organized, structured
	// and equipped to perform. Or, the specialized military or paramilitary branch in
	// which an individual serves, their specialization. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctRole string `json:"functRole"`
	// The distance between Mean Sea Level and a referenced ellipsoid.
	GeoidalMslSep float64 `json:"geoidalMslSep"`
	// Unique identifier of the contact for this operating unit.
	IDContact string `json:"idContact"`
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
	// Unique identifier of the location record for this operating unit.
	IDLocation string `json:"idLocation"`
	// Unique identifier of the record, auto-generated by the system.
	IDOperatingUnit string `json:"idOperatingUnit"`
	// Unique identifier of the organization record for this operating unit.
	IDOrganization string `json:"idOrganization"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// Model representation of a location, which is a specific fixed point on the earth
	// and is used to denote the locations of fixed sensors, operating units, etc.
	Location LocationFull `json:"location"`
	// Location name for the coordinates.
	LocName string `json:"locName"`
	// Indicates the reason that the operating unit is at that location. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason string `json:"locReason"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// This field contains a value indicating whether the record is a master unit
	// record (True) or a detail record (False). Master records contain basic
	// information that does not change over time for each unit that has been selected
	// to be projected.
	MasterUnit bool `json:"masterUnit"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts: 4Q (grid zone
	// designator, GZD) FJ (the 100,000-meter square identifier) 12345678 (numerical
	// location; easting is 1234 and northing is 5678, in this case specifying a
	// location with 10 m resolution).
	MilGrid string `json:"milGrid"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator
	MilGridSys string `json:"milGridSys"`
	// Indicates the principal type of mission that an operating unit is organized and
	// equipped to perform. The specific usage and enumerations contained in this field
	// may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimary string `json:"msnPrimary"`
	// Indicates the principal specialty type of mission that an operating unit is
	// organized and equipped to perform. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpecialty string `json:"msnPrimarySpecialty"`
	// Remarks contain amplifying information for a specific service. The information
	// may contain context and interpretations for consumer use.
	OperatingUnitRemarks []OperatingunitTupleResponseOperatingUnitRemark `json:"operatingUnitRemarks"`
	// The Degree to which an operating unit is ready to perform the overall
	// operational mission(s) for which it was organized and equipped. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	OperStatus string `json:"operStatus"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	Organization OrganizationFull `json:"organization"`
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
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Values are: A - Active I -
	// Inactive K - Acknowledged L - Local Q - A nominated (NOM) or Data Change Request
	// (DCR) record R - Production reduced by CMD decision W - Working Record.
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
	// This field contains a value indicating whether the record is a stylized
	// operating unit record (True) or a regular operating unit record (False). A
	// stylized operating unit is a type of operating unit with one set of equipment
	// that can be assigned to one or more superiors. A stylized operating unit is
	// generally useful for lower echelon operating units where the number of operating
	// units and types of equipment are equal for multiple organizations. In lieu of
	// creating unique operating unit records for each operating unit, a template is
	// created for the operating unit and its equipment. This template enables the user
	// to assign the operating unit to multiple organizations.
	StylizedUnit bool `json:"stylizedUnit"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode string `json:"symCode"`
	// An optional identifier for this operating unit that may be composed from items
	// such as the originating organization, allegiance, one-up number, etc.
	UnitIdentifier string `json:"unitIdentifier"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Universal Transverse Mercator (UTM) grid coordinates. Pos. 1-2, UTM Zone Column
	// [01-60 Pos. 3, UTM Zone Row [C-HJ-NP-X] Pos. 4, UTM False Easting [0-9] Pos.
	// 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9] Pos. 10-11, UTM False Northing
	// [0-9][0-9] Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm string `json:"utm"`
	// World Aeronautical Chart identifier for the area in which a designated operating
	// unit is located.
	Wac string `json:"wac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		AirDefArea            respjson.Field
		Allegiance            respjson.Field
		AltAllegiance         respjson.Field
		AltCountryCode        respjson.Field
		AltOperatingUnitID    respjson.Field
		ClassRating           respjson.Field
		Condition             respjson.Field
		ConditionAvail        respjson.Field
		Coord                 respjson.Field
		CoordDatum            respjson.Field
		CoordDerivAcc         respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DeployStatus          respjson.Field
		Description           respjson.Field
		DivCat                respjson.Field
		Echelon               respjson.Field
		EchelonTier           respjson.Field
		ElevMsl               respjson.Field
		ElevMslConfLvl        respjson.Field
		ElevMslDerivAcc       respjson.Field
		Eval                  respjson.Field
		FlagFlown             respjson.Field
		FleetID               respjson.Field
		Force                 respjson.Field
		ForceName             respjson.Field
		Fpa                   respjson.Field
		FunctRole             respjson.Field
		GeoidalMslSep         respjson.Field
		IDContact             respjson.Field
		Ident                 respjson.Field
		IDLocation            respjson.Field
		IDOperatingUnit       respjson.Field
		IDOrganization        respjson.Field
		Lat                   respjson.Field
		Location              respjson.Field
		LocName               respjson.Field
		LocReason             respjson.Field
		Lon                   respjson.Field
		MasterUnit            respjson.Field
		MilGrid               respjson.Field
		MilGridSys            respjson.Field
		MsnPrimary            respjson.Field
		MsnPrimarySpecialty   respjson.Field
		OperatingUnitRemarks  respjson.Field
		OperStatus            respjson.Field
		Organization          respjson.Field
		Origin                respjson.Field
		PolSubdiv             respjson.Field
		RecStatus             respjson.Field
		ReferenceDoc          respjson.Field
		ResProd               respjson.Field
		ReviewDate            respjson.Field
		StylizedUnit          respjson.Field
		SymCode               respjson.Field
		UnitIdentifier        respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Utm                   respjson.Field
		Wac                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperatingunitTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *OperatingunitTupleResponse) UnmarshalJSON(data []byte) error {
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
type OperatingunitTupleResponseDataMode string

const (
	OperatingunitTupleResponseDataModeReal      OperatingunitTupleResponseDataMode = "REAL"
	OperatingunitTupleResponseDataModeTest      OperatingunitTupleResponseDataMode = "TEST"
	OperatingunitTupleResponseDataModeSimulated OperatingunitTupleResponseDataMode = "SIMULATED"
	OperatingunitTupleResponseDataModeExercise  OperatingunitTupleResponseDataMode = "EXERCISE"
)

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
type OperatingunitTupleResponseOperatingUnitRemark struct {
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
	// The ID of the operating unit to which this remark applies.
	IDOperatingUnit string `json:"idOperatingUnit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Unique identifier of the unit remark record from the originating system.
	AltRmkID string `json:"altRmkId"`
	// The remark type identifier. For example, the Mobility Air Forces (MAF) remark
	// code, defined in the Airfield Suitability and Restriction Report (ASRR).
	Code string `json:"code"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The name of the remark.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The remark type (e.g. Caution, Information, Misc, Restriction, etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOperatingUnit       respjson.Field
		Source                respjson.Field
		Text                  respjson.Field
		ID                    respjson.Field
		AltRmkID              respjson.Field
		Code                  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Name                  respjson.Field
		Origin                respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperatingunitTupleResponseOperatingUnitRemark) RawJSON() string { return r.JSON.raw }
func (r *OperatingunitTupleResponseOperatingUnitRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperatingunitNewParams struct {
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
	DataMode OperatingunitNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Name of the operating unit.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea param.Opt[string] `json:"airDefArea,omitzero"`
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit owes its allegiance. This field will be set to "OTHR"
	// if the source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	// Unique identifier of the operating unit record from the originating system.
	AltOperatingUnitID param.Opt[string] `json:"altOperatingUnitId,omitzero"`
	// Indicates the importance of the operating unit to the OES or MIR system. This
	// data element is restricted to update by DIA (DB-4). Valid values are: 0 - Does
	// not meet criteria above 1 - Primary importance to system 2 - Secondary
	// importance to system 3 - Tertiary importance to system O - Other. Explain in
	// Remarks.
	ClassRating param.Opt[string] `json:"classRating,omitzero"`
	// The physical manner of being or state of existence of the operating unit. A
	// physical condition that must be considered in the determining of a course of
	// action. The specific usage and enumerations contained in this field may be found
	// in the documentation provided in the referenceDoc field. If referenceDoc not
	// provided, users may consult the data provider.
	Condition param.Opt[string] `json:"condition,omitzero"`
	// Availability of the operating unit relative to its condition. Indicates the
	// reason the operating unit is not fully operational. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	ConditionAvail param.Opt[string] `json:"conditionAvail,omitzero"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc.
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]
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
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit geographic coordinates reside . This field will be set
	// to "OTHR" if the source value does not match a UDL country code value
	// (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// A code describing the amount of operating unit participation in a deployment.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	DeployStatus param.Opt[string] `json:"deployStatus,omitzero"`
	// Description of the operating unit.
	Description param.Opt[string] `json:"description,omitzero"`
	// Combat status of a divisional or equivalent operating unit. Currently, this data
	// element applies only to operating units of the Former Soviet Union. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	DivCat param.Opt[string] `json:"divCat,omitzero"`
	// Organizational level of the operating unit. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Echelon param.Opt[string] `json:"echelon,omitzero"`
	// Indicates the major group or level to which an echelon belongs. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	EchelonTier param.Opt[string] `json:"echelonTier,omitzero"`
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum.
	ElevMsl param.Opt[float64] `json:"elevMsl,omitzero"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl param.Opt[int64] `json:"elevMslConfLvl,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc param.Opt[float64] `json:"elevMslDerivAcc,omitzero"`
	// The Intelligence Confidence Level or the Reliability/degree of confidence that
	// the analyst has assigned to the data within this record. The numerical range is
	// from 1 to 9 with 1 representing the highest confidence level.
	Eval param.Opt[int64] `json:"eval,omitzero"`
	// The country code of the observed flag flown.
	FlagFlown param.Opt[string] `json:"flagFlown,omitzero"`
	// Naval fleet to which an operating unit is assigned. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FleetID param.Opt[string] `json:"fleetId,omitzero"`
	// An aggregation of military units within a single service (i.e., ARMY, AIR FORCE,
	// etc.) which operates under a single authority to accomplish a common mission.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Force param.Opt[string] `json:"force,omitzero"`
	// The specific name for a given force. For example, Force = ADF (Air Defense
	// Force) and Force Name = Army Air Defense Force.
	ForceName param.Opt[string] `json:"forceName,omitzero"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa param.Opt[string] `json:"fpa,omitzero"`
	// Principal combat-related role that an operating unit is organized, structured
	// and equipped to perform. Or, the specialized military or paramilitary branch in
	// which an individual serves, their specialization. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctRole param.Opt[string] `json:"functRole,omitzero"`
	// The distance between Mean Sea Level and a referenced ellipsoid.
	GeoidalMslSep param.Opt[float64] `json:"geoidalMslSep,omitzero"`
	// Unique identifier of the contact for this operating unit.
	IDContact param.Opt[string] `json:"idContact,omitzero"`
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
	// Unique identifier of the location record for this operating unit.
	IDLocation param.Opt[string] `json:"idLocation,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDOperatingUnit param.Opt[string] `json:"idOperatingUnit,omitzero"`
	// Unique identifier of the organization record for this operating unit.
	IDOrganization param.Opt[string] `json:"idOrganization,omitzero"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Location name for the coordinates.
	LocName param.Opt[string] `json:"locName,omitzero"`
	// Indicates the reason that the operating unit is at that location. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason param.Opt[string] `json:"locReason,omitzero"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// This field contains a value indicating whether the record is a master unit
	// record (True) or a detail record (False). Master records contain basic
	// information that does not change over time for each unit that has been selected
	// to be projected.
	MasterUnit param.Opt[bool] `json:"masterUnit,omitzero"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts: 4Q (grid zone
	// designator, GZD) FJ (the 100,000-meter square identifier) 12345678 (numerical
	// location; easting is 1234 and northing is 5678, in this case specifying a
	// location with 10 m resolution).
	MilGrid param.Opt[string] `json:"milGrid,omitzero"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator
	MilGridSys param.Opt[string] `json:"milGridSys,omitzero"`
	// Indicates the principal type of mission that an operating unit is organized and
	// equipped to perform. The specific usage and enumerations contained in this field
	// may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimary param.Opt[string] `json:"msnPrimary,omitzero"`
	// Indicates the principal specialty type of mission that an operating unit is
	// organized and equipped to perform. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpecialty param.Opt[string] `json:"msnPrimarySpecialty,omitzero"`
	// The Degree to which an operating unit is ready to perform the overall
	// operational mission(s) for which it was organized and equipped. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
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
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Values are: A - Active I -
	// Inactive K - Acknowledged L - Local Q - A nominated (NOM) or Data Change Request
	// (DCR) record R - Production reduced by CMD decision W - Working Record.
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
	// This field contains a value indicating whether the record is a stylized
	// operating unit record (True) or a regular operating unit record (False). A
	// stylized operating unit is a type of operating unit with one set of equipment
	// that can be assigned to one or more superiors. A stylized operating unit is
	// generally useful for lower echelon operating units where the number of operating
	// units and types of equipment are equal for multiple organizations. In lieu of
	// creating unique operating unit records for each operating unit, a template is
	// created for the operating unit and its equipment. This template enables the user
	// to assign the operating unit to multiple organizations.
	StylizedUnit param.Opt[bool] `json:"stylizedUnit,omitzero"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode param.Opt[string] `json:"symCode,omitzero"`
	// An optional identifier for this operating unit that may be composed from items
	// such as the originating organization, allegiance, one-up number, etc.
	UnitIdentifier param.Opt[string] `json:"unitIdentifier,omitzero"`
	// Universal Transverse Mercator (UTM) grid coordinates. Pos. 1-2, UTM Zone Column
	// [01-60 Pos. 3, UTM Zone Row [C-HJ-NP-X] Pos. 4, UTM False Easting [0-9] Pos.
	// 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9] Pos. 10-11, UTM False Northing
	// [0-9][0-9] Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm param.Opt[string] `json:"utm,omitzero"`
	// World Aeronautical Chart identifier for the area in which a designated operating
	// unit is located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	paramObj
}

func (r OperatingunitNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OperatingunitNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OperatingunitNewParams) UnmarshalJSON(data []byte) error {
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
type OperatingunitNewParamsDataMode string

const (
	OperatingunitNewParamsDataModeReal      OperatingunitNewParamsDataMode = "REAL"
	OperatingunitNewParamsDataModeTest      OperatingunitNewParamsDataMode = "TEST"
	OperatingunitNewParamsDataModeSimulated OperatingunitNewParamsDataMode = "SIMULATED"
	OperatingunitNewParamsDataModeExercise  OperatingunitNewParamsDataMode = "EXERCISE"
)

type OperatingunitUpdateParams struct {
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
	DataMode OperatingunitUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Name of the operating unit.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea param.Opt[string] `json:"airDefArea,omitzero"`
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit owes its allegiance. This field will be set to "OTHR"
	// if the source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	// Unique identifier of the operating unit record from the originating system.
	AltOperatingUnitID param.Opt[string] `json:"altOperatingUnitId,omitzero"`
	// Indicates the importance of the operating unit to the OES or MIR system. This
	// data element is restricted to update by DIA (DB-4). Valid values are: 0 - Does
	// not meet criteria above 1 - Primary importance to system 2 - Secondary
	// importance to system 3 - Tertiary importance to system O - Other. Explain in
	// Remarks.
	ClassRating param.Opt[string] `json:"classRating,omitzero"`
	// The physical manner of being or state of existence of the operating unit. A
	// physical condition that must be considered in the determining of a course of
	// action. The specific usage and enumerations contained in this field may be found
	// in the documentation provided in the referenceDoc field. If referenceDoc not
	// provided, users may consult the data provider.
	Condition param.Opt[string] `json:"condition,omitzero"`
	// Availability of the operating unit relative to its condition. Indicates the
	// reason the operating unit is not fully operational. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	ConditionAvail param.Opt[string] `json:"conditionAvail,omitzero"`
	// Indicates any of the magnitudes that serve to define the position of a point by
	// reference to a fixed figure, system of lines, etc.
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U]
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
	// The DoD Standard country code designator for the country or political entity to
	// which the operating unit geographic coordinates reside . This field will be set
	// to "OTHR" if the source value does not match a UDL country code value
	// (ISO-3166-ALPHA-2).
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// A code describing the amount of operating unit participation in a deployment.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	DeployStatus param.Opt[string] `json:"deployStatus,omitzero"`
	// Description of the operating unit.
	Description param.Opt[string] `json:"description,omitzero"`
	// Combat status of a divisional or equivalent operating unit. Currently, this data
	// element applies only to operating units of the Former Soviet Union. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	DivCat param.Opt[string] `json:"divCat,omitzero"`
	// Organizational level of the operating unit. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Echelon param.Opt[string] `json:"echelon,omitzero"`
	// Indicates the major group or level to which an echelon belongs. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	EchelonTier param.Opt[string] `json:"echelonTier,omitzero"`
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum.
	ElevMsl param.Opt[float64] `json:"elevMsl,omitzero"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl param.Opt[int64] `json:"elevMslConfLvl,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc param.Opt[float64] `json:"elevMslDerivAcc,omitzero"`
	// The Intelligence Confidence Level or the Reliability/degree of confidence that
	// the analyst has assigned to the data within this record. The numerical range is
	// from 1 to 9 with 1 representing the highest confidence level.
	Eval param.Opt[int64] `json:"eval,omitzero"`
	// The country code of the observed flag flown.
	FlagFlown param.Opt[string] `json:"flagFlown,omitzero"`
	// Naval fleet to which an operating unit is assigned. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FleetID param.Opt[string] `json:"fleetId,omitzero"`
	// An aggregation of military units within a single service (i.e., ARMY, AIR FORCE,
	// etc.) which operates under a single authority to accomplish a common mission.
	// The specific usage and enumerations contained in this field may be found in the
	// documentation provided in the referenceDoc field. If referenceDoc not provided,
	// users may consult the data provider.
	Force param.Opt[string] `json:"force,omitzero"`
	// The specific name for a given force. For example, Force = ADF (Air Defense
	// Force) and Force Name = Army Air Defense Force.
	ForceName param.Opt[string] `json:"forceName,omitzero"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa param.Opt[string] `json:"fpa,omitzero"`
	// Principal combat-related role that an operating unit is organized, structured
	// and equipped to perform. Or, the specialized military or paramilitary branch in
	// which an individual serves, their specialization. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctRole param.Opt[string] `json:"functRole,omitzero"`
	// The distance between Mean Sea Level and a referenced ellipsoid.
	GeoidalMslSep param.Opt[float64] `json:"geoidalMslSep,omitzero"`
	// Unique identifier of the contact for this operating unit.
	IDContact param.Opt[string] `json:"idContact,omitzero"`
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
	// Unique identifier of the location record for this operating unit.
	IDLocation param.Opt[string] `json:"idLocation,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDOperatingUnit param.Opt[string] `json:"idOperatingUnit,omitzero"`
	// Unique identifier of the organization record for this operating unit.
	IDOrganization param.Opt[string] `json:"idOrganization,omitzero"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Location name for the coordinates.
	LocName param.Opt[string] `json:"locName,omitzero"`
	// Indicates the reason that the operating unit is at that location. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	LocReason param.Opt[string] `json:"locReason,omitzero"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// This field contains a value indicating whether the record is a master unit
	// record (True) or a detail record (False). Master records contain basic
	// information that does not change over time for each unit that has been selected
	// to be projected.
	MasterUnit param.Opt[bool] `json:"masterUnit,omitzero"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts: 4Q (grid zone
	// designator, GZD) FJ (the 100,000-meter square identifier) 12345678 (numerical
	// location; easting is 1234 and northing is 5678, in this case specifying a
	// location with 10 m resolution).
	MilGrid param.Opt[string] `json:"milGrid,omitzero"`
	// Indicates the grid system used in the development of the milGrid coordinates.
	// Values are:
	//
	// # UPS - Universal Polar System
	//
	// UTM - Universal Transverse Mercator
	MilGridSys param.Opt[string] `json:"milGridSys,omitzero"`
	// Indicates the principal type of mission that an operating unit is organized and
	// equipped to perform. The specific usage and enumerations contained in this field
	// may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimary param.Opt[string] `json:"msnPrimary,omitzero"`
	// Indicates the principal specialty type of mission that an operating unit is
	// organized and equipped to perform. The specific usage and enumerations contained
	// in this field may be found in the documentation provided in the referenceDoc
	// field. If referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpecialty param.Opt[string] `json:"msnPrimarySpecialty,omitzero"`
	// The Degree to which an operating unit is ready to perform the overall
	// operational mission(s) for which it was organized and equipped. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
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
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs. Values are: A - Active I -
	// Inactive K - Acknowledged L - Local Q - A nominated (NOM) or Data Change Request
	// (DCR) record R - Production reduced by CMD decision W - Working Record.
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
	// This field contains a value indicating whether the record is a stylized
	// operating unit record (True) or a regular operating unit record (False). A
	// stylized operating unit is a type of operating unit with one set of equipment
	// that can be assigned to one or more superiors. A stylized operating unit is
	// generally useful for lower echelon operating units where the number of operating
	// units and types of equipment are equal for multiple organizations. In lieu of
	// creating unique operating unit records for each operating unit, a template is
	// created for the operating unit and its equipment. This template enables the user
	// to assign the operating unit to multiple organizations.
	StylizedUnit param.Opt[bool] `json:"stylizedUnit,omitzero"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element AFFILIATION.
	SymCode param.Opt[string] `json:"symCode,omitzero"`
	// An optional identifier for this operating unit that may be composed from items
	// such as the originating organization, allegiance, one-up number, etc.
	UnitIdentifier param.Opt[string] `json:"unitIdentifier,omitzero"`
	// Universal Transverse Mercator (UTM) grid coordinates. Pos. 1-2, UTM Zone Column
	// [01-60 Pos. 3, UTM Zone Row [C-HJ-NP-X] Pos. 4, UTM False Easting [0-9] Pos.
	// 5-9, UTM Meter Easting [0-9][0-9][0-9][0-9][0-9] Pos. 10-11, UTM False Northing
	// [0-9][0-9] Pos. 12-16, UTM Meter Northing [0-9][0-9][0-9][0-9][0-9].
	Utm param.Opt[string] `json:"utm,omitzero"`
	// World Aeronautical Chart identifier for the area in which a designated operating
	// unit is located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	paramObj
}

func (r OperatingunitUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OperatingunitUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OperatingunitUpdateParams) UnmarshalJSON(data []byte) error {
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
type OperatingunitUpdateParamsDataMode string

const (
	OperatingunitUpdateParamsDataModeReal      OperatingunitUpdateParamsDataMode = "REAL"
	OperatingunitUpdateParamsDataModeTest      OperatingunitUpdateParamsDataMode = "TEST"
	OperatingunitUpdateParamsDataModeSimulated OperatingunitUpdateParamsDataMode = "SIMULATED"
	OperatingunitUpdateParamsDataModeExercise  OperatingunitUpdateParamsDataMode = "EXERCISE"
)

type OperatingunitListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperatingunitListParams]'s query parameters as
// `url.Values`.
func (r OperatingunitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperatingunitCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperatingunitCountParams]'s query parameters as
// `url.Values`.
func (r OperatingunitCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperatingunitGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperatingunitGetParams]'s query parameters as `url.Values`.
func (r OperatingunitGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperatingunitTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperatingunitTupleParams]'s query parameters as
// `url.Values`.
func (r OperatingunitTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
