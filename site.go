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

// SiteService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteService] method instead.
type SiteService struct {
	Options    []option.RequestOption
	Operations SiteOperationService
}

// NewSiteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteService(opts ...option.RequestOption) (r SiteService) {
	r = SiteService{}
	r.Options = opts
	r.Operations = NewSiteOperationService(opts...)
	return
}

// Service operation to take a single Site as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *SiteService) New(ctx context.Context, body SiteNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/site"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Site. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *SiteService) Update(ctx context.Context, id string, body SiteUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/site/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SiteService) List(ctx context.Context, query SiteListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SiteListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/site"
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
func (r *SiteService) ListAutoPaging(ctx context.Context, query SiteListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SiteListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SiteService) Count(ctx context.Context, query SiteCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/site/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Site record by its unique ID passed as a path
// parameter.
func (r *SiteService) Get(ctx context.Context, id string, query SiteGetParams, opts ...option.RequestOption) (res *SiteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/site/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SiteService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/site/queryhelp"
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
func (r *SiteService) Tuple(ctx context.Context, query SiteTupleParams, opts ...option.RequestOption) (res *[]SiteTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/site/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Properties and characteristics of a site entity, such as an airbase, airfield,
// naval station, etc.
type SiteListResponse struct {
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
	DataMode SiteListResponseDataMode `json:"dataMode,required"`
	// The name of this site.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Indicates the function or mission of an entity, which that entity may or may not
	// be engaged in at any particular time. Typically refers to a unit, organization,
	// or installation/site performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Activity string `json:"activity"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea string `json:"airDefArea"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the site owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance string `json:"allegiance"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance string `json:"altAllegiance"`
	// The Basic Encyclopedia Number associated with the Site. Uniquely identifies the
	// installation of a site. The beNumber is generated based on the value input for
	// the COORD to determine the appropriate World Aeronautical Chart (WAC) location
	// identifier, the system assigned record originator and a one-up-number.
	BeNumber string `json:"beNumber"`
	// The category code that represents the associated site purpose within the target
	// system.
	CatCode string `json:"catCode"`
	// Textual Description of Site catCode.
	CatText string `json:"catText"`
	// Indicates the importance of the entity to the OES or MIR system. This data
	// element is restricted to update by DIA (DB-4). Valid values are:
	//
	// 0 - Does not meet criteria above
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U].
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
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum, in meters.
	ElevMsl float64 `json:"elevMsl"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl int64 `json:"elevMslConfLvl"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc float64 `json:"elevMslDerivAcc"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval int64 `json:"eval"`
	// The Federal Aviation Administration (FAA) Location ID of this site, if
	// applicable.
	Faa string `json:"faa"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa string `json:"fpa"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary string `json:"functPrimary"`
	// Geographical region code used by the Requirements Management System (RMS) as
	// specified by National Geospatial Agency (NGA) in Flight Information Publications
	// (FIPS) 10-4, Appendix 3 - Country Code and Geographic Region Codes. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	GeoArea string `json:"geoArea"`
	// The distance between Mean Sea Level and a referenced ellipsoid, in meters.
	GeoidalMslSep float64 `json:"geoidalMslSep"`
	// Indicates the amount or degree of deviation from the horizontal represented as a
	// percent. Grade is determined by the formula: vertical distance (VD) divided by
	// horizontal distance (HD) times 100. VD is the difference between the highest and
	// lowest elevation within the entity. HD is the linear distance between the
	// highest and lowest elevation.
	Grade int64 `json:"grade"`
	// The International Air Transport Association (IATA) code of this site, if
	// applicable.
	Iata string `json:"iata"`
	// The International Civil Aviation Organization (ICAO) code of this site, if
	// applicable.
	Icao string `json:"icao"`
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
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity string `json:"idEntity"`
	// Unique identifier of the Parent Site record associated with this Site record.
	IDParentSite string `json:"idParentSite"`
	// Indicates the normal usage of the Landing Zone (LZ). Intended as, but not
	// constrained to MIDB Helocopter Landing Area usage value definitions:
	//
	// # AF - Airfield
	//
	// # FD - Field
	//
	// HC - High Crop. 1 meter and over.
	//
	// # HY - Highway
	//
	// # LB - Lake Bed
	//
	// LC - Low Crop. 0-1 meters
	//
	// O - Other. Explain In Remarks.
	//
	// # PD - Paddy
	//
	// # PK - Park
	//
	// # PS - Pasture
	//
	// # RB - Riverbed
	//
	// # SP - Sport Field
	//
	// # U - Unknown
	//
	// Z - Inconclusive Analysis.
	LzUsage string `json:"lzUsage"`
	// The length of the longest runway at this site, if applicable, in meters.
	MaxRunwayLength int64 `json:"maxRunwayLength"`
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
	// Indicates the principal type of mission that an entity is organized and equipped
	// to perform. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	MsnPrimary string `json:"msnPrimary"`
	// Indicates the principal specialty type of mission that an entity is organized
	// and equipped to perform. The specific usage and enumerations contained in this
	// field may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpec string `json:"msnPrimarySpec"`
	// Optional notes/comments for the site.
	Notes string `json:"notes"`
	// A sites ability to conduct nuclear warfare. Valid Values are:
	//
	// # A - Nuclear Ammo Or Warheads Available
	//
	// # N - No Nuclear Offense
	//
	// O - Other. Explain in Remarks
	//
	// # U - Unknown
	//
	// # W - Nuclear Weapons Available
	//
	// # Y - Nuclear Warfare Offensive Capability
	//
	// Z - Inconclusive Analysis.
	NucCap string `json:"nucCap"`
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
	// Unique identifier of the LZ record from the originating system.
	OrigLzID string `json:"origLzId"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Unique identifier of the Site record from the originating system.
	OrigSiteID string `json:"origSiteID"`
	// The O-suffix associated with this site. The O-suffix is a five-character
	// alpha/numeric system used to identify a site, or demographic area, within an
	// installation. The Installation Basic Encyclopedia (beNumber), in conjunction
	// with the O-suffix, uniquely identifies the Site. The Installation beNumber and
	// oSuffix are also used in conjunction with the catCode to classify the function
	// or purpose of the facility.
	Osuffix string `json:"osuffix"`
	// Site number of a specific electronic site or its associated equipment.
	Pin string `json:"pin"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv string `json:"polSubdiv"`
	// Indicates whether the facility is in or outside of a populated area. True, the
	// facility is in or within 5 NM of a populated area. False, the facility is
	// outside a populated area.
	PopArea bool `json:"popArea"`
	// Indicates the distance to nearest populated area (over 1,000 people) in nautical
	// miles.
	PopAreaProx float64 `json:"popAreaProx"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs.
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
	// The reference documentation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc string `json:"referenceDoc"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd string `json:"resProd"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency, in ISO8601 UTC format. This date cannot be
	// greater than the current date.
	ReviewDate time.Time `json:"reviewDate" format:"date"`
	// The number of runways at the site, if applicable.
	Runways int64 `json:"runways"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element ident.
	SymCode string `json:"symCode"`
	// The type of this site (AIRBASE, AIRFIELD, AIRPORT, NAVAL STATION, etc.).
	Type string `json:"type"`
	// The use authorization type of this site (e.g MILITARY, CIVIL, JOINT-USE, etc.).
	Usage string `json:"usage"`
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
	// Maximum expected height of the vegetation in the Landing Zone (LZ), in meters.
	VegHt float64 `json:"vegHt"`
	// The predominant vegetation found in the Landing Zone (LZ). The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	VegType string `json:"vegType"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac string `json:"wac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Activity              respjson.Field
		AirDefArea            respjson.Field
		Allegiance            respjson.Field
		AltAllegiance         respjson.Field
		BeNumber              respjson.Field
		CatCode               respjson.Field
		CatText               respjson.Field
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
		Eval                  respjson.Field
		Faa                   respjson.Field
		Fpa                   respjson.Field
		FunctPrimary          respjson.Field
		GeoArea               respjson.Field
		GeoidalMslSep         respjson.Field
		Grade                 respjson.Field
		Iata                  respjson.Field
		Icao                  respjson.Field
		Ident                 respjson.Field
		IDEntity              respjson.Field
		IDParentSite          respjson.Field
		LzUsage               respjson.Field
		MaxRunwayLength       respjson.Field
		MilGrid               respjson.Field
		MilGridSys            respjson.Field
		MsnPrimary            respjson.Field
		MsnPrimarySpec        respjson.Field
		Notes                 respjson.Field
		NucCap                respjson.Field
		OperStatus            respjson.Field
		Origin                respjson.Field
		OrigLzID              respjson.Field
		OrigNetwork           respjson.Field
		OrigSiteID            respjson.Field
		Osuffix               respjson.Field
		Pin                   respjson.Field
		PolSubdiv             respjson.Field
		PopArea               respjson.Field
		PopAreaProx           respjson.Field
		RecStatus             respjson.Field
		ReferenceDoc          respjson.Field
		ResProd               respjson.Field
		ReviewDate            respjson.Field
		Runways               respjson.Field
		SymCode               respjson.Field
		Type                  respjson.Field
		Usage                 respjson.Field
		Utm                   respjson.Field
		VegHt                 respjson.Field
		VegType               respjson.Field
		Wac                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteListResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteListResponse) UnmarshalJSON(data []byte) error {
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
type SiteListResponseDataMode string

const (
	SiteListResponseDataModeReal      SiteListResponseDataMode = "REAL"
	SiteListResponseDataModeTest      SiteListResponseDataMode = "TEST"
	SiteListResponseDataModeSimulated SiteListResponseDataMode = "SIMULATED"
	SiteListResponseDataModeExercise  SiteListResponseDataMode = "EXERCISE"
)

// Properties and characteristics of a site entity, such as an airbase, airfield,
// naval station, etc.
type SiteGetResponse struct {
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
	DataMode SiteGetResponseDataMode `json:"dataMode,required"`
	// The name of this site.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Indicates the function or mission of an entity, which that entity may or may not
	// be engaged in at any particular time. Typically refers to a unit, organization,
	// or installation/site performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Activity string `json:"activity"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea string `json:"airDefArea"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the site owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance string `json:"allegiance"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance string `json:"altAllegiance"`
	// The Basic Encyclopedia Number associated with the Site. Uniquely identifies the
	// installation of a site. The beNumber is generated based on the value input for
	// the COORD to determine the appropriate World Aeronautical Chart (WAC) location
	// identifier, the system assigned record originator and a one-up-number.
	BeNumber string `json:"beNumber"`
	// The category code that represents the associated site purpose within the target
	// system.
	CatCode string `json:"catCode"`
	// Textual Description of Site catCode.
	CatText string `json:"catText"`
	// Indicates the importance of the entity to the OES or MIR system. This data
	// element is restricted to update by DIA (DB-4). Valid values are:
	//
	// 0 - Does not meet criteria above
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U].
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
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum, in meters.
	ElevMsl float64 `json:"elevMsl"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl int64 `json:"elevMslConfLvl"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc float64 `json:"elevMslDerivAcc"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityFull `json:"entity"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval int64 `json:"eval"`
	// The Federal Aviation Administration (FAA) Location ID of this site, if
	// applicable.
	Faa string `json:"faa"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa string `json:"fpa"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary string `json:"functPrimary"`
	// Geographical region code used by the Requirements Management System (RMS) as
	// specified by National Geospatial Agency (NGA) in Flight Information Publications
	// (FIPS) 10-4, Appendix 3 - Country Code and Geographic Region Codes. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	GeoArea string `json:"geoArea"`
	// The distance between Mean Sea Level and a referenced ellipsoid, in meters.
	GeoidalMslSep float64 `json:"geoidalMslSep"`
	// Indicates the amount or degree of deviation from the horizontal represented as a
	// percent. Grade is determined by the formula: vertical distance (VD) divided by
	// horizontal distance (HD) times 100. VD is the difference between the highest and
	// lowest elevation within the entity. HD is the linear distance between the
	// highest and lowest elevation.
	Grade int64 `json:"grade"`
	// The International Air Transport Association (IATA) code of this site, if
	// applicable.
	Iata string `json:"iata"`
	// The International Civil Aviation Organization (ICAO) code of this site, if
	// applicable.
	Icao string `json:"icao"`
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
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity string `json:"idEntity"`
	// Unique identifier of the Parent Site record associated with this Site record.
	IDParentSite string `json:"idParentSite"`
	// Indicates the normal usage of the Landing Zone (LZ). Intended as, but not
	// constrained to MIDB Helocopter Landing Area usage value definitions:
	//
	// # AF - Airfield
	//
	// # FD - Field
	//
	// HC - High Crop. 1 meter and over.
	//
	// # HY - Highway
	//
	// # LB - Lake Bed
	//
	// LC - Low Crop. 0-1 meters
	//
	// O - Other. Explain In Remarks.
	//
	// # PD - Paddy
	//
	// # PK - Park
	//
	// # PS - Pasture
	//
	// # RB - Riverbed
	//
	// # SP - Sport Field
	//
	// # U - Unknown
	//
	// Z - Inconclusive Analysis.
	LzUsage string `json:"lzUsage"`
	// The length of the longest runway at this site, if applicable, in meters.
	MaxRunwayLength int64 `json:"maxRunwayLength"`
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
	// Indicates the principal type of mission that an entity is organized and equipped
	// to perform. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	MsnPrimary string `json:"msnPrimary"`
	// Indicates the principal specialty type of mission that an entity is organized
	// and equipped to perform. The specific usage and enumerations contained in this
	// field may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpec string `json:"msnPrimarySpec"`
	// Optional notes/comments for the site.
	Notes string `json:"notes"`
	// A sites ability to conduct nuclear warfare. Valid Values are:
	//
	// # A - Nuclear Ammo Or Warheads Available
	//
	// # N - No Nuclear Offense
	//
	// O - Other. Explain in Remarks
	//
	// # U - Unknown
	//
	// # W - Nuclear Weapons Available
	//
	// # Y - Nuclear Warfare Offensive Capability
	//
	// Z - Inconclusive Analysis.
	NucCap string `json:"nucCap"`
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
	// Unique identifier of the LZ record from the originating system.
	OrigLzID string `json:"origLzId"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Unique identifier of the Site record from the originating system.
	OrigSiteID string `json:"origSiteID"`
	// The O-suffix associated with this site. The O-suffix is a five-character
	// alpha/numeric system used to identify a site, or demographic area, within an
	// installation. The Installation Basic Encyclopedia (beNumber), in conjunction
	// with the O-suffix, uniquely identifies the Site. The Installation beNumber and
	// oSuffix are also used in conjunction with the catCode to classify the function
	// or purpose of the facility.
	Osuffix string `json:"osuffix"`
	// Site number of a specific electronic site or its associated equipment.
	Pin string `json:"pin"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv string `json:"polSubdiv"`
	// Indicates whether the facility is in or outside of a populated area. True, the
	// facility is in or within 5 NM of a populated area. False, the facility is
	// outside a populated area.
	PopArea bool `json:"popArea"`
	// Indicates the distance to nearest populated area (over 1,000 people) in nautical
	// miles.
	PopAreaProx float64 `json:"popAreaProx"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs.
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
	// The reference documentation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc string `json:"referenceDoc"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd string `json:"resProd"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency, in ISO8601 UTC format. This date cannot be
	// greater than the current date.
	ReviewDate time.Time `json:"reviewDate" format:"date"`
	// The number of runways at the site, if applicable.
	Runways        int64                          `json:"runways"`
	SiteOperations []SiteGetResponseSiteOperation `json:"siteOperations"`
	// Remarks contain amplifying information for a specific service. The information
	// may contain context and interpretations for consumer use.
	SiteRemarks []SiteGetResponseSiteRemark `json:"siteRemarks"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element ident.
	SymCode string `json:"symCode"`
	// The type of this site (AIRBASE, AIRFIELD, AIRPORT, NAVAL STATION, etc.).
	Type string `json:"type"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The use authorization type of this site (e.g MILITARY, CIVIL, JOINT-USE, etc.).
	Usage string `json:"usage"`
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
	// Maximum expected height of the vegetation in the Landing Zone (LZ), in meters.
	VegHt float64 `json:"vegHt"`
	// The predominant vegetation found in the Landing Zone (LZ). The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	VegType string `json:"vegType"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac string `json:"wac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Activity              respjson.Field
		AirDefArea            respjson.Field
		Allegiance            respjson.Field
		AltAllegiance         respjson.Field
		BeNumber              respjson.Field
		CatCode               respjson.Field
		CatText               respjson.Field
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
		Entity                respjson.Field
		Eval                  respjson.Field
		Faa                   respjson.Field
		Fpa                   respjson.Field
		FunctPrimary          respjson.Field
		GeoArea               respjson.Field
		GeoidalMslSep         respjson.Field
		Grade                 respjson.Field
		Iata                  respjson.Field
		Icao                  respjson.Field
		Ident                 respjson.Field
		IDEntity              respjson.Field
		IDParentSite          respjson.Field
		LzUsage               respjson.Field
		MaxRunwayLength       respjson.Field
		MilGrid               respjson.Field
		MilGridSys            respjson.Field
		MsnPrimary            respjson.Field
		MsnPrimarySpec        respjson.Field
		Notes                 respjson.Field
		NucCap                respjson.Field
		OperStatus            respjson.Field
		Origin                respjson.Field
		OrigLzID              respjson.Field
		OrigNetwork           respjson.Field
		OrigSiteID            respjson.Field
		Osuffix               respjson.Field
		Pin                   respjson.Field
		PolSubdiv             respjson.Field
		PopArea               respjson.Field
		PopAreaProx           respjson.Field
		RecStatus             respjson.Field
		ReferenceDoc          respjson.Field
		ResProd               respjson.Field
		ReviewDate            respjson.Field
		Runways               respjson.Field
		SiteOperations        respjson.Field
		SiteRemarks           respjson.Field
		SymCode               respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Usage                 respjson.Field
		Utm                   respjson.Field
		VegHt                 respjson.Field
		VegType               respjson.Field
		Wac                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponse) UnmarshalJSON(data []byte) error {
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
type SiteGetResponseDataMode string

const (
	SiteGetResponseDataModeReal      SiteGetResponseDataMode = "REAL"
	SiteGetResponseDataModeTest      SiteGetResponseDataMode = "TEST"
	SiteGetResponseDataModeSimulated SiteGetResponseDataMode = "SIMULATED"
	SiteGetResponseDataModeExercise  SiteGetResponseDataMode = "EXERCISE"
)

// Site operating details concerning the hours of operation, operational
// limitations, site navigation, and waivers associated with the Site.
type SiteGetResponseSiteOperation struct {
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
	// The ID of the parent site.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Collection providing hours of operation and other information specific to a day
	// of the week.
	DailyOperations []SiteGetResponseSiteOperationDailyOperation `json:"dailyOperations"`
	// The name of the person who made the most recent change to data in the
	// DailyOperations collection.
	DopsLastChangedBy string `json:"dopsLastChangedBy"`
	// The datetime of the most recent change made to data in the DailyOperations
	// collection, in ISO 8601 UTC format with millisecond precision.
	DopsLastChangedDate time.Time `json:"dopsLastChangedDate" format:"date-time"`
	// The reason for the most recent change to data in the dailyOperations collection.
	DopsLastChangedReason string `json:"dopsLastChangedReason"`
	// Id of the associated launchSite entity.
	IDLaunchSite string `json:"idLaunchSite"`
	// Collection providing maximum on ground (MOG) information for specific aircraft
	// at the site associated with this SiteOperations record.
	MaximumOnGrounds []SiteGetResponseSiteOperationMaximumOnGround `json:"maximumOnGrounds"`
	// The name of the person who made the most recent change to data in the
	// MaximumOnGrounds collection.
	MogsLastChangedBy string `json:"mogsLastChangedBy"`
	// The datetime of the most recent change made to data in the MaximumOnGrounds
	// collection, in ISO 8601 UTC format with millisecond precision.
	MogsLastChangedDate time.Time `json:"mogsLastChangedDate" format:"date-time"`
	// The reason for the most recent change to data in the MaximumOnGrounds
	// collection.
	MogsLastChangedReason string `json:"mogsLastChangedReason"`
	// Collection providing relevant information in the event of deviations/exceptions
	// to normal operations.
	OperationalDeviations []SiteGetResponseSiteOperationOperationalDeviation `json:"operationalDeviations"`
	// Collection of planning information associated with this SiteOperations record.
	OperationalPlannings []SiteGetResponseSiteOperationOperationalPlanning `json:"operationalPlannings"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Collection detailing operational pathways at the Site associated with this
	// SiteOperations record.
	Pathways []SiteGetResponseSiteOperationPathway `json:"pathways"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Collection documenting operational waivers that have been issued for the Site
	// associated with this record.
	Waivers []SiteGetResponseSiteOperationWaiver `json:"waivers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSite                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DailyOperations       respjson.Field
		DopsLastChangedBy     respjson.Field
		DopsLastChangedDate   respjson.Field
		DopsLastChangedReason respjson.Field
		IDLaunchSite          respjson.Field
		MaximumOnGrounds      respjson.Field
		MogsLastChangedBy     respjson.Field
		MogsLastChangedDate   respjson.Field
		MogsLastChangedReason respjson.Field
		OperationalDeviations respjson.Field
		OperationalPlannings  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Pathways              respjson.Field
		SourceDl              respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Waivers               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponseSiteOperation) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponseSiteOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing hours of operation and other information specific to a day
// of the week.
type SiteGetResponseSiteOperationDailyOperation struct {
	// The day of the week to which this operational information pertains.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DayOfWeek string `json:"dayOfWeek"`
	// A collection containing the operational start and stop times scheduled for the
	// day of the week specified.
	OperatingHours []SiteGetResponseSiteOperationDailyOperationOperatingHour `json:"operatingHours"`
	// The name or type of operation to which this information pertains.
	OperationName string `json:"operationName"`
	// The name of the person who made the most recent change to this DailyOperation
	// data.
	OphrsLastChangedBy string `json:"ophrsLastChangedBy"`
	// The datetime of the most recent change made to this DailyOperation data, in ISO
	// 8601 UTC format with millisecond precision.
	OphrsLastChangedDate time.Time `json:"ophrsLastChangedDate" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayOfWeek            respjson.Field
		OperatingHours       respjson.Field
		OperationName        respjson.Field
		OphrsLastChangedBy   respjson.Field
		OphrsLastChangedDate respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponseSiteOperationDailyOperation) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponseSiteOperationDailyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection containing the operational start and stop times scheduled for the
// day of the week specified.
type SiteGetResponseSiteOperationDailyOperationOperatingHour struct {
	// The Zulu (UTC) operational start time, expressed in ISO 8601 format as HH:MM.
	OpStartTime string `json:"opStartTime"`
	// The Zulu (UTC) operational stop time, expressed in ISO 8601 format as HH:MM.
	OpStopTime string `json:"opStopTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpStartTime respjson.Field
		OpStopTime  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponseSiteOperationDailyOperationOperatingHour) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponseSiteOperationDailyOperationOperatingHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing maximum on ground (MOG) information for specific aircraft
// at the site associated with this SiteOperations record.
type SiteGetResponseSiteOperationMaximumOnGround struct {
	// The Model Design Series (MDS) designation of the aircraft to which this maximum
	// on ground (MOG) data pertains.
	AircraftMds string `json:"aircraftMDS"`
	// Maximum on ground (MOG) number of contingent aircraft based on spacing and
	// manpower, for the aircraft type specified.
	ContingencyMog int64 `json:"contingencyMOG"`
	// The name of the person who made the most recent change to this maximum on ground
	// data.
	MogLastChangedBy string `json:"mogLastChangedBy"`
	// The datetime of the most recent change made to this maximum on ground data, in
	// ISO 8601 UTC format with millisecond precision.
	MogLastChangedDate time.Time `json:"mogLastChangedDate" format:"date-time"`
	// Maximum on ground (MOG) number of parking wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideParkingMog int64 `json:"wideParkingMOG"`
	// Maximum on ground (MOG) number of working wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideWorkingMog int64 `json:"wideWorkingMOG"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AircraftMds        respjson.Field
		ContingencyMog     respjson.Field
		MogLastChangedBy   respjson.Field
		MogLastChangedDate respjson.Field
		WideParkingMog     respjson.Field
		WideWorkingMog     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponseSiteOperationMaximumOnGround) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponseSiteOperationMaximumOnGround) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing relevant information in the event of deviations/exceptions
// to normal operations.
type SiteGetResponseSiteOperationOperationalDeviation struct {
	// The Model Design Series (MDS) designation of the aircraft affected by this
	// operational deviation.
	AffectedAircraftMds string `json:"affectedAircraftMDS"`
	// The maximum on ground (MOG) number for aircraft affected by this operational
	// deviation.
	AffectedMog int64 `json:"affectedMOG"`
	// On ground time for aircraft affected by this operational deviation.
	AircraftOnGroundTime string `json:"aircraftOnGroundTime"`
	// Rest time for crew affected by this operational deviation.
	CrewRestTime string `json:"crewRestTime"`
	// The name of the person who made the most recent change to this
	// OperationalDeviation data.
	OdLastChangedBy string `json:"odLastChangedBy"`
	// The datetime of the most recent change made to this OperationalDeviation data,
	// in ISO 8601 UTC format with millisecond precision.
	OdLastChangedDate time.Time `json:"odLastChangedDate" format:"date-time"`
	// Text remark regarding this operational deviation.
	OdRemark string `json:"odRemark"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AffectedAircraftMds  respjson.Field
		AffectedMog          respjson.Field
		AircraftOnGroundTime respjson.Field
		CrewRestTime         respjson.Field
		OdLastChangedBy      respjson.Field
		OdLastChangedDate    respjson.Field
		OdRemark             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponseSiteOperationOperationalDeviation) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponseSiteOperationOperationalDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of planning information associated with this SiteOperations record.
type SiteGetResponseSiteOperationOperationalPlanning struct {
	// The end date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpEndDate time.Time `json:"opEndDate" format:"date-time"`
	// The name of the person who made the most recent change made to this
	// OperationalPlanning data.
	OpLastChangedBy string `json:"opLastChangedBy"`
	// The datetime of the most recent change made to this OperationalPlanning data, in
	// ISO8601 UTC format with millisecond precision.
	OpLastChangedDate time.Time `json:"opLastChangedDate" format:"date-time"`
	// Remark text regarding this operation planning.
	OpRemark string `json:"opRemark"`
	// The person, unit, organization, etc. responsible for this operation planning.
	OpSource string `json:"opSource"`
	// The start date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpStartDate time.Time `json:"opStartDate" format:"date-time"`
	// The status of this operational planning.
	OpStatus string `json:"opStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpEndDate         respjson.Field
		OpLastChangedBy   respjson.Field
		OpLastChangedDate respjson.Field
		OpRemark          respjson.Field
		OpSource          respjson.Field
		OpStartDate       respjson.Field
		OpStatus          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponseSiteOperationOperationalPlanning) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponseSiteOperationOperationalPlanning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection detailing operational pathways at the Site associated with this
// SiteOperations record.
type SiteGetResponseSiteOperationPathway struct {
	// Text defining this pathway from its constituent parts.
	PwDefinition string `json:"pwDefinition"`
	// The name of the person who made the most recent change to this Pathway data.
	PwLastChangedBy string `json:"pwLastChangedBy"`
	// The datetime of the most recent change made to this Pathway data, in ISO 8601
	// UTC format with millisecond precision.
	PwLastChangedDate time.Time `json:"pwLastChangedDate" format:"date-time"`
	// The type of paths that constitute this pathway.
	PwType string `json:"pwType"`
	// The intended use of this pathway.
	PwUsage string `json:"pwUsage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PwDefinition      respjson.Field
		PwLastChangedBy   respjson.Field
		PwLastChangedDate respjson.Field
		PwType            respjson.Field
		PwUsage           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponseSiteOperationPathway) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponseSiteOperationPathway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection documenting operational waivers that have been issued for the Site
// associated with this record.
type SiteGetResponseSiteOperationWaiver struct {
	// The expiration date of this waiver, in ISO8601 UTC format with millisecond
	// precision.
	ExpirationDate time.Time `json:"expirationDate" format:"date-time"`
	// Boolean indicating whether or not this waiver has expired.
	HasExpired bool `json:"hasExpired"`
	// The issue date of this waiver, in ISO8601 UTC format with millisecond precision.
	IssueDate time.Time `json:"issueDate" format:"date-time"`
	// The name of the person who issued this waiver.
	IssuerName string `json:"issuerName"`
	// The name of the person requesting this waiver.
	RequesterName string `json:"requesterName"`
	// The phone number of the person requesting this waiver.
	RequesterPhoneNumber string `json:"requesterPhoneNumber"`
	// The unit requesting this waiver.
	RequestingUnit string `json:"requestingUnit"`
	// Description of the entities to which this waiver applies.
	WaiverAppliesTo string `json:"waiverAppliesTo"`
	// The description of this waiver.
	WaiverDescription string `json:"waiverDescription"`
	// The name of the person who made the most recent change to this Waiver data.
	WaiverLastChangedBy string `json:"waiverLastChangedBy"`
	// The datetime of the most recent change made to this waiver data, in ISO8601 UTC
	// format with millisecond precision.
	WaiverLastChangedDate time.Time `json:"waiverLastChangedDate" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpirationDate        respjson.Field
		HasExpired            respjson.Field
		IssueDate             respjson.Field
		IssuerName            respjson.Field
		RequesterName         respjson.Field
		RequesterPhoneNumber  respjson.Field
		RequestingUnit        respjson.Field
		WaiverAppliesTo       respjson.Field
		WaiverDescription     respjson.Field
		WaiverLastChangedBy   respjson.Field
		WaiverLastChangedDate respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponseSiteOperationWaiver) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponseSiteOperationWaiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
type SiteGetResponseSiteRemark struct {
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
	// The ID of the site to which this remark applies.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Unique identifier of the Site Remark record from the originating system.
	OrigRmkID string `json:"origRmkId"`
	// The remark type (e.g. Caution, Information, Misc, Restriction, etc.).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSite                respjson.Field
		Source                respjson.Field
		Text                  respjson.Field
		ID                    respjson.Field
		Code                  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Name                  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigRmkID             respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteGetResponseSiteRemark) RawJSON() string { return r.JSON.raw }
func (r *SiteGetResponseSiteRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Properties and characteristics of a site entity, such as an airbase, airfield,
// naval station, etc.
type SiteTupleResponse struct {
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
	DataMode SiteTupleResponseDataMode `json:"dataMode,required"`
	// The name of this site.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Indicates the function or mission of an entity, which that entity may or may not
	// be engaged in at any particular time. Typically refers to a unit, organization,
	// or installation/site performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Activity string `json:"activity"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea string `json:"airDefArea"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the site owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance string `json:"allegiance"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance string `json:"altAllegiance"`
	// The Basic Encyclopedia Number associated with the Site. Uniquely identifies the
	// installation of a site. The beNumber is generated based on the value input for
	// the COORD to determine the appropriate World Aeronautical Chart (WAC) location
	// identifier, the system assigned record originator and a one-up-number.
	BeNumber string `json:"beNumber"`
	// The category code that represents the associated site purpose within the target
	// system.
	CatCode string `json:"catCode"`
	// Textual Description of Site catCode.
	CatText string `json:"catText"`
	// Indicates the importance of the entity to the OES or MIR system. This data
	// element is restricted to update by DIA (DB-4). Valid values are:
	//
	// 0 - Does not meet criteria above
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U].
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
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum, in meters.
	ElevMsl float64 `json:"elevMsl"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl int64 `json:"elevMslConfLvl"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc float64 `json:"elevMslDerivAcc"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityFull `json:"entity"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval int64 `json:"eval"`
	// The Federal Aviation Administration (FAA) Location ID of this site, if
	// applicable.
	Faa string `json:"faa"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa string `json:"fpa"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary string `json:"functPrimary"`
	// Geographical region code used by the Requirements Management System (RMS) as
	// specified by National Geospatial Agency (NGA) in Flight Information Publications
	// (FIPS) 10-4, Appendix 3 - Country Code and Geographic Region Codes. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	GeoArea string `json:"geoArea"`
	// The distance between Mean Sea Level and a referenced ellipsoid, in meters.
	GeoidalMslSep float64 `json:"geoidalMslSep"`
	// Indicates the amount or degree of deviation from the horizontal represented as a
	// percent. Grade is determined by the formula: vertical distance (VD) divided by
	// horizontal distance (HD) times 100. VD is the difference between the highest and
	// lowest elevation within the entity. HD is the linear distance between the
	// highest and lowest elevation.
	Grade int64 `json:"grade"`
	// The International Air Transport Association (IATA) code of this site, if
	// applicable.
	Iata string `json:"iata"`
	// The International Civil Aviation Organization (ICAO) code of this site, if
	// applicable.
	Icao string `json:"icao"`
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
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity string `json:"idEntity"`
	// Unique identifier of the Parent Site record associated with this Site record.
	IDParentSite string `json:"idParentSite"`
	// Indicates the normal usage of the Landing Zone (LZ). Intended as, but not
	// constrained to MIDB Helocopter Landing Area usage value definitions:
	//
	// # AF - Airfield
	//
	// # FD - Field
	//
	// HC - High Crop. 1 meter and over.
	//
	// # HY - Highway
	//
	// # LB - Lake Bed
	//
	// LC - Low Crop. 0-1 meters
	//
	// O - Other. Explain In Remarks.
	//
	// # PD - Paddy
	//
	// # PK - Park
	//
	// # PS - Pasture
	//
	// # RB - Riverbed
	//
	// # SP - Sport Field
	//
	// # U - Unknown
	//
	// Z - Inconclusive Analysis.
	LzUsage string `json:"lzUsage"`
	// The length of the longest runway at this site, if applicable, in meters.
	MaxRunwayLength int64 `json:"maxRunwayLength"`
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
	// Indicates the principal type of mission that an entity is organized and equipped
	// to perform. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	MsnPrimary string `json:"msnPrimary"`
	// Indicates the principal specialty type of mission that an entity is organized
	// and equipped to perform. The specific usage and enumerations contained in this
	// field may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpec string `json:"msnPrimarySpec"`
	// Optional notes/comments for the site.
	Notes string `json:"notes"`
	// A sites ability to conduct nuclear warfare. Valid Values are:
	//
	// # A - Nuclear Ammo Or Warheads Available
	//
	// # N - No Nuclear Offense
	//
	// O - Other. Explain in Remarks
	//
	// # U - Unknown
	//
	// # W - Nuclear Weapons Available
	//
	// # Y - Nuclear Warfare Offensive Capability
	//
	// Z - Inconclusive Analysis.
	NucCap string `json:"nucCap"`
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
	// Unique identifier of the LZ record from the originating system.
	OrigLzID string `json:"origLzId"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Unique identifier of the Site record from the originating system.
	OrigSiteID string `json:"origSiteID"`
	// The O-suffix associated with this site. The O-suffix is a five-character
	// alpha/numeric system used to identify a site, or demographic area, within an
	// installation. The Installation Basic Encyclopedia (beNumber), in conjunction
	// with the O-suffix, uniquely identifies the Site. The Installation beNumber and
	// oSuffix are also used in conjunction with the catCode to classify the function
	// or purpose of the facility.
	Osuffix string `json:"osuffix"`
	// Site number of a specific electronic site or its associated equipment.
	Pin string `json:"pin"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv string `json:"polSubdiv"`
	// Indicates whether the facility is in or outside of a populated area. True, the
	// facility is in or within 5 NM of a populated area. False, the facility is
	// outside a populated area.
	PopArea bool `json:"popArea"`
	// Indicates the distance to nearest populated area (over 1,000 people) in nautical
	// miles.
	PopAreaProx float64 `json:"popAreaProx"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs.
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
	// The reference documentation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc string `json:"referenceDoc"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd string `json:"resProd"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency, in ISO8601 UTC format. This date cannot be
	// greater than the current date.
	ReviewDate time.Time `json:"reviewDate" format:"date"`
	// The number of runways at the site, if applicable.
	Runways        int64                            `json:"runways"`
	SiteOperations []SiteTupleResponseSiteOperation `json:"siteOperations"`
	// Remarks contain amplifying information for a specific service. The information
	// may contain context and interpretations for consumer use.
	SiteRemarks []SiteTupleResponseSiteRemark `json:"siteRemarks"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element ident.
	SymCode string `json:"symCode"`
	// The type of this site (AIRBASE, AIRFIELD, AIRPORT, NAVAL STATION, etc.).
	Type string `json:"type"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The use authorization type of this site (e.g MILITARY, CIVIL, JOINT-USE, etc.).
	Usage string `json:"usage"`
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
	// Maximum expected height of the vegetation in the Landing Zone (LZ), in meters.
	VegHt float64 `json:"vegHt"`
	// The predominant vegetation found in the Landing Zone (LZ). The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	VegType string `json:"vegType"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac string `json:"wac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Activity              respjson.Field
		AirDefArea            respjson.Field
		Allegiance            respjson.Field
		AltAllegiance         respjson.Field
		BeNumber              respjson.Field
		CatCode               respjson.Field
		CatText               respjson.Field
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
		Entity                respjson.Field
		Eval                  respjson.Field
		Faa                   respjson.Field
		Fpa                   respjson.Field
		FunctPrimary          respjson.Field
		GeoArea               respjson.Field
		GeoidalMslSep         respjson.Field
		Grade                 respjson.Field
		Iata                  respjson.Field
		Icao                  respjson.Field
		Ident                 respjson.Field
		IDEntity              respjson.Field
		IDParentSite          respjson.Field
		LzUsage               respjson.Field
		MaxRunwayLength       respjson.Field
		MilGrid               respjson.Field
		MilGridSys            respjson.Field
		MsnPrimary            respjson.Field
		MsnPrimarySpec        respjson.Field
		Notes                 respjson.Field
		NucCap                respjson.Field
		OperStatus            respjson.Field
		Origin                respjson.Field
		OrigLzID              respjson.Field
		OrigNetwork           respjson.Field
		OrigSiteID            respjson.Field
		Osuffix               respjson.Field
		Pin                   respjson.Field
		PolSubdiv             respjson.Field
		PopArea               respjson.Field
		PopAreaProx           respjson.Field
		RecStatus             respjson.Field
		ReferenceDoc          respjson.Field
		ResProd               respjson.Field
		ReviewDate            respjson.Field
		Runways               respjson.Field
		SiteOperations        respjson.Field
		SiteRemarks           respjson.Field
		SymCode               respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Usage                 respjson.Field
		Utm                   respjson.Field
		VegHt                 respjson.Field
		VegType               respjson.Field
		Wac                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteTupleResponse) UnmarshalJSON(data []byte) error {
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
type SiteTupleResponseDataMode string

const (
	SiteTupleResponseDataModeReal      SiteTupleResponseDataMode = "REAL"
	SiteTupleResponseDataModeTest      SiteTupleResponseDataMode = "TEST"
	SiteTupleResponseDataModeSimulated SiteTupleResponseDataMode = "SIMULATED"
	SiteTupleResponseDataModeExercise  SiteTupleResponseDataMode = "EXERCISE"
)

// Site operating details concerning the hours of operation, operational
// limitations, site navigation, and waivers associated with the Site.
type SiteTupleResponseSiteOperation struct {
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
	// The ID of the parent site.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Collection providing hours of operation and other information specific to a day
	// of the week.
	DailyOperations []SiteTupleResponseSiteOperationDailyOperation `json:"dailyOperations"`
	// The name of the person who made the most recent change to data in the
	// DailyOperations collection.
	DopsLastChangedBy string `json:"dopsLastChangedBy"`
	// The datetime of the most recent change made to data in the DailyOperations
	// collection, in ISO 8601 UTC format with millisecond precision.
	DopsLastChangedDate time.Time `json:"dopsLastChangedDate" format:"date-time"`
	// The reason for the most recent change to data in the dailyOperations collection.
	DopsLastChangedReason string `json:"dopsLastChangedReason"`
	// Id of the associated launchSite entity.
	IDLaunchSite string `json:"idLaunchSite"`
	// Collection providing maximum on ground (MOG) information for specific aircraft
	// at the site associated with this SiteOperations record.
	MaximumOnGrounds []SiteTupleResponseSiteOperationMaximumOnGround `json:"maximumOnGrounds"`
	// The name of the person who made the most recent change to data in the
	// MaximumOnGrounds collection.
	MogsLastChangedBy string `json:"mogsLastChangedBy"`
	// The datetime of the most recent change made to data in the MaximumOnGrounds
	// collection, in ISO 8601 UTC format with millisecond precision.
	MogsLastChangedDate time.Time `json:"mogsLastChangedDate" format:"date-time"`
	// The reason for the most recent change to data in the MaximumOnGrounds
	// collection.
	MogsLastChangedReason string `json:"mogsLastChangedReason"`
	// Collection providing relevant information in the event of deviations/exceptions
	// to normal operations.
	OperationalDeviations []SiteTupleResponseSiteOperationOperationalDeviation `json:"operationalDeviations"`
	// Collection of planning information associated with this SiteOperations record.
	OperationalPlannings []SiteTupleResponseSiteOperationOperationalPlanning `json:"operationalPlannings"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Collection detailing operational pathways at the Site associated with this
	// SiteOperations record.
	Pathways []SiteTupleResponseSiteOperationPathway `json:"pathways"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Collection documenting operational waivers that have been issued for the Site
	// associated with this record.
	Waivers []SiteTupleResponseSiteOperationWaiver `json:"waivers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSite                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DailyOperations       respjson.Field
		DopsLastChangedBy     respjson.Field
		DopsLastChangedDate   respjson.Field
		DopsLastChangedReason respjson.Field
		IDLaunchSite          respjson.Field
		MaximumOnGrounds      respjson.Field
		MogsLastChangedBy     respjson.Field
		MogsLastChangedDate   respjson.Field
		MogsLastChangedReason respjson.Field
		OperationalDeviations respjson.Field
		OperationalPlannings  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Pathways              respjson.Field
		SourceDl              respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Waivers               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponseSiteOperation) RawJSON() string { return r.JSON.raw }
func (r *SiteTupleResponseSiteOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing hours of operation and other information specific to a day
// of the week.
type SiteTupleResponseSiteOperationDailyOperation struct {
	// The day of the week to which this operational information pertains.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DayOfWeek string `json:"dayOfWeek"`
	// A collection containing the operational start and stop times scheduled for the
	// day of the week specified.
	OperatingHours []SiteTupleResponseSiteOperationDailyOperationOperatingHour `json:"operatingHours"`
	// The name or type of operation to which this information pertains.
	OperationName string `json:"operationName"`
	// The name of the person who made the most recent change to this DailyOperation
	// data.
	OphrsLastChangedBy string `json:"ophrsLastChangedBy"`
	// The datetime of the most recent change made to this DailyOperation data, in ISO
	// 8601 UTC format with millisecond precision.
	OphrsLastChangedDate time.Time `json:"ophrsLastChangedDate" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayOfWeek            respjson.Field
		OperatingHours       respjson.Field
		OperationName        respjson.Field
		OphrsLastChangedBy   respjson.Field
		OphrsLastChangedDate respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponseSiteOperationDailyOperation) RawJSON() string { return r.JSON.raw }
func (r *SiteTupleResponseSiteOperationDailyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection containing the operational start and stop times scheduled for the
// day of the week specified.
type SiteTupleResponseSiteOperationDailyOperationOperatingHour struct {
	// The Zulu (UTC) operational start time, expressed in ISO 8601 format as HH:MM.
	OpStartTime string `json:"opStartTime"`
	// The Zulu (UTC) operational stop time, expressed in ISO 8601 format as HH:MM.
	OpStopTime string `json:"opStopTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpStartTime respjson.Field
		OpStopTime  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponseSiteOperationDailyOperationOperatingHour) RawJSON() string {
	return r.JSON.raw
}
func (r *SiteTupleResponseSiteOperationDailyOperationOperatingHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing maximum on ground (MOG) information for specific aircraft
// at the site associated with this SiteOperations record.
type SiteTupleResponseSiteOperationMaximumOnGround struct {
	// The Model Design Series (MDS) designation of the aircraft to which this maximum
	// on ground (MOG) data pertains.
	AircraftMds string `json:"aircraftMDS"`
	// Maximum on ground (MOG) number of contingent aircraft based on spacing and
	// manpower, for the aircraft type specified.
	ContingencyMog int64 `json:"contingencyMOG"`
	// The name of the person who made the most recent change to this maximum on ground
	// data.
	MogLastChangedBy string `json:"mogLastChangedBy"`
	// The datetime of the most recent change made to this maximum on ground data, in
	// ISO 8601 UTC format with millisecond precision.
	MogLastChangedDate time.Time `json:"mogLastChangedDate" format:"date-time"`
	// Maximum on ground (MOG) number of parking wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideParkingMog int64 `json:"wideParkingMOG"`
	// Maximum on ground (MOG) number of working wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideWorkingMog int64 `json:"wideWorkingMOG"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AircraftMds        respjson.Field
		ContingencyMog     respjson.Field
		MogLastChangedBy   respjson.Field
		MogLastChangedDate respjson.Field
		WideParkingMog     respjson.Field
		WideWorkingMog     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponseSiteOperationMaximumOnGround) RawJSON() string { return r.JSON.raw }
func (r *SiteTupleResponseSiteOperationMaximumOnGround) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing relevant information in the event of deviations/exceptions
// to normal operations.
type SiteTupleResponseSiteOperationOperationalDeviation struct {
	// The Model Design Series (MDS) designation of the aircraft affected by this
	// operational deviation.
	AffectedAircraftMds string `json:"affectedAircraftMDS"`
	// The maximum on ground (MOG) number for aircraft affected by this operational
	// deviation.
	AffectedMog int64 `json:"affectedMOG"`
	// On ground time for aircraft affected by this operational deviation.
	AircraftOnGroundTime string `json:"aircraftOnGroundTime"`
	// Rest time for crew affected by this operational deviation.
	CrewRestTime string `json:"crewRestTime"`
	// The name of the person who made the most recent change to this
	// OperationalDeviation data.
	OdLastChangedBy string `json:"odLastChangedBy"`
	// The datetime of the most recent change made to this OperationalDeviation data,
	// in ISO 8601 UTC format with millisecond precision.
	OdLastChangedDate time.Time `json:"odLastChangedDate" format:"date-time"`
	// Text remark regarding this operational deviation.
	OdRemark string `json:"odRemark"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AffectedAircraftMds  respjson.Field
		AffectedMog          respjson.Field
		AircraftOnGroundTime respjson.Field
		CrewRestTime         respjson.Field
		OdLastChangedBy      respjson.Field
		OdLastChangedDate    respjson.Field
		OdRemark             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponseSiteOperationOperationalDeviation) RawJSON() string { return r.JSON.raw }
func (r *SiteTupleResponseSiteOperationOperationalDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of planning information associated with this SiteOperations record.
type SiteTupleResponseSiteOperationOperationalPlanning struct {
	// The end date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpEndDate time.Time `json:"opEndDate" format:"date-time"`
	// The name of the person who made the most recent change made to this
	// OperationalPlanning data.
	OpLastChangedBy string `json:"opLastChangedBy"`
	// The datetime of the most recent change made to this OperationalPlanning data, in
	// ISO8601 UTC format with millisecond precision.
	OpLastChangedDate time.Time `json:"opLastChangedDate" format:"date-time"`
	// Remark text regarding this operation planning.
	OpRemark string `json:"opRemark"`
	// The person, unit, organization, etc. responsible for this operation planning.
	OpSource string `json:"opSource"`
	// The start date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpStartDate time.Time `json:"opStartDate" format:"date-time"`
	// The status of this operational planning.
	OpStatus string `json:"opStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpEndDate         respjson.Field
		OpLastChangedBy   respjson.Field
		OpLastChangedDate respjson.Field
		OpRemark          respjson.Field
		OpSource          respjson.Field
		OpStartDate       respjson.Field
		OpStatus          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponseSiteOperationOperationalPlanning) RawJSON() string { return r.JSON.raw }
func (r *SiteTupleResponseSiteOperationOperationalPlanning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection detailing operational pathways at the Site associated with this
// SiteOperations record.
type SiteTupleResponseSiteOperationPathway struct {
	// Text defining this pathway from its constituent parts.
	PwDefinition string `json:"pwDefinition"`
	// The name of the person who made the most recent change to this Pathway data.
	PwLastChangedBy string `json:"pwLastChangedBy"`
	// The datetime of the most recent change made to this Pathway data, in ISO 8601
	// UTC format with millisecond precision.
	PwLastChangedDate time.Time `json:"pwLastChangedDate" format:"date-time"`
	// The type of paths that constitute this pathway.
	PwType string `json:"pwType"`
	// The intended use of this pathway.
	PwUsage string `json:"pwUsage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PwDefinition      respjson.Field
		PwLastChangedBy   respjson.Field
		PwLastChangedDate respjson.Field
		PwType            respjson.Field
		PwUsage           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponseSiteOperationPathway) RawJSON() string { return r.JSON.raw }
func (r *SiteTupleResponseSiteOperationPathway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection documenting operational waivers that have been issued for the Site
// associated with this record.
type SiteTupleResponseSiteOperationWaiver struct {
	// The expiration date of this waiver, in ISO8601 UTC format with millisecond
	// precision.
	ExpirationDate time.Time `json:"expirationDate" format:"date-time"`
	// Boolean indicating whether or not this waiver has expired.
	HasExpired bool `json:"hasExpired"`
	// The issue date of this waiver, in ISO8601 UTC format with millisecond precision.
	IssueDate time.Time `json:"issueDate" format:"date-time"`
	// The name of the person who issued this waiver.
	IssuerName string `json:"issuerName"`
	// The name of the person requesting this waiver.
	RequesterName string `json:"requesterName"`
	// The phone number of the person requesting this waiver.
	RequesterPhoneNumber string `json:"requesterPhoneNumber"`
	// The unit requesting this waiver.
	RequestingUnit string `json:"requestingUnit"`
	// Description of the entities to which this waiver applies.
	WaiverAppliesTo string `json:"waiverAppliesTo"`
	// The description of this waiver.
	WaiverDescription string `json:"waiverDescription"`
	// The name of the person who made the most recent change to this Waiver data.
	WaiverLastChangedBy string `json:"waiverLastChangedBy"`
	// The datetime of the most recent change made to this waiver data, in ISO8601 UTC
	// format with millisecond precision.
	WaiverLastChangedDate time.Time `json:"waiverLastChangedDate" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpirationDate        respjson.Field
		HasExpired            respjson.Field
		IssueDate             respjson.Field
		IssuerName            respjson.Field
		RequesterName         respjson.Field
		RequesterPhoneNumber  respjson.Field
		RequestingUnit        respjson.Field
		WaiverAppliesTo       respjson.Field
		WaiverDescription     respjson.Field
		WaiverLastChangedBy   respjson.Field
		WaiverLastChangedDate respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponseSiteOperationWaiver) RawJSON() string { return r.JSON.raw }
func (r *SiteTupleResponseSiteOperationWaiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remarks contain amplifying information for a specific service. The information
// may contain context and interpretations for consumer use.
type SiteTupleResponseSiteRemark struct {
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
	// The ID of the site to which this remark applies.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The text of the remark.
	Text string `json:"text,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Unique identifier of the Site Remark record from the originating system.
	OrigRmkID string `json:"origRmkId"`
	// The remark type (e.g. Caution, Information, Misc, Restriction, etc.).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSite                respjson.Field
		Source                respjson.Field
		Text                  respjson.Field
		ID                    respjson.Field
		Code                  respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Name                  respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigRmkID             respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteTupleResponseSiteRemark) RawJSON() string { return r.JSON.raw }
func (r *SiteTupleResponseSiteRemark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteNewParams struct {
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
	DataMode SiteNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The name of this site.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Indicates the function or mission of an entity, which that entity may or may not
	// be engaged in at any particular time. Typically refers to a unit, organization,
	// or installation/site performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Activity param.Opt[string] `json:"activity,omitzero"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea param.Opt[string] `json:"airDefArea,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the site owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance param.Opt[string] `json:"allegiance,omitzero"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance param.Opt[string] `json:"altAllegiance,omitzero"`
	// The Basic Encyclopedia Number associated with the Site. Uniquely identifies the
	// installation of a site. The beNumber is generated based on the value input for
	// the COORD to determine the appropriate World Aeronautical Chart (WAC) location
	// identifier, the system assigned record originator and a one-up-number.
	BeNumber param.Opt[string] `json:"beNumber,omitzero"`
	// The category code that represents the associated site purpose within the target
	// system.
	CatCode param.Opt[string] `json:"catCode,omitzero"`
	// Textual Description of Site catCode.
	CatText param.Opt[string] `json:"catText,omitzero"`
	// Indicates the importance of the entity to the OES or MIR system. This data
	// element is restricted to update by DIA (DB-4). Valid values are:
	//
	// 0 - Does not meet criteria above
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U].
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
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum, in meters.
	ElevMsl param.Opt[float64] `json:"elevMsl,omitzero"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl param.Opt[int64] `json:"elevMslConfLvl,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc param.Opt[float64] `json:"elevMslDerivAcc,omitzero"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval param.Opt[int64] `json:"eval,omitzero"`
	// The Federal Aviation Administration (FAA) Location ID of this site, if
	// applicable.
	Faa param.Opt[string] `json:"faa,omitzero"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa param.Opt[string] `json:"fpa,omitzero"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary param.Opt[string] `json:"functPrimary,omitzero"`
	// Geographical region code used by the Requirements Management System (RMS) as
	// specified by National Geospatial Agency (NGA) in Flight Information Publications
	// (FIPS) 10-4, Appendix 3 - Country Code and Geographic Region Codes. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	GeoArea param.Opt[string] `json:"geoArea,omitzero"`
	// The distance between Mean Sea Level and a referenced ellipsoid, in meters.
	GeoidalMslSep param.Opt[float64] `json:"geoidalMslSep,omitzero"`
	// Indicates the amount or degree of deviation from the horizontal represented as a
	// percent. Grade is determined by the formula: vertical distance (VD) divided by
	// horizontal distance (HD) times 100. VD is the difference between the highest and
	// lowest elevation within the entity. HD is the linear distance between the
	// highest and lowest elevation.
	Grade param.Opt[int64] `json:"grade,omitzero"`
	// The International Air Transport Association (IATA) code of this site, if
	// applicable.
	Iata param.Opt[string] `json:"iata,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of this site, if
	// applicable.
	Icao param.Opt[string] `json:"icao,omitzero"`
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
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Unique identifier of the Parent Site record associated with this Site record.
	IDParentSite param.Opt[string] `json:"idParentSite,omitzero"`
	// Indicates the normal usage of the Landing Zone (LZ). Intended as, but not
	// constrained to MIDB Helocopter Landing Area usage value definitions:
	//
	// # AF - Airfield
	//
	// # FD - Field
	//
	// HC - High Crop. 1 meter and over.
	//
	// # HY - Highway
	//
	// # LB - Lake Bed
	//
	// LC - Low Crop. 0-1 meters
	//
	// O - Other. Explain In Remarks.
	//
	// # PD - Paddy
	//
	// # PK - Park
	//
	// # PS - Pasture
	//
	// # RB - Riverbed
	//
	// # SP - Sport Field
	//
	// # U - Unknown
	//
	// Z - Inconclusive Analysis.
	LzUsage param.Opt[string] `json:"lzUsage,omitzero"`
	// The length of the longest runway at this site, if applicable, in meters.
	MaxRunwayLength param.Opt[int64] `json:"maxRunwayLength,omitzero"`
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
	// Indicates the principal type of mission that an entity is organized and equipped
	// to perform. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	MsnPrimary param.Opt[string] `json:"msnPrimary,omitzero"`
	// Indicates the principal specialty type of mission that an entity is organized
	// and equipped to perform. The specific usage and enumerations contained in this
	// field may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpec param.Opt[string] `json:"msnPrimarySpec,omitzero"`
	// Optional notes/comments for the site.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// A sites ability to conduct nuclear warfare. Valid Values are:
	//
	// # A - Nuclear Ammo Or Warheads Available
	//
	// # N - No Nuclear Offense
	//
	// O - Other. Explain in Remarks
	//
	// # U - Unknown
	//
	// # W - Nuclear Weapons Available
	//
	// # Y - Nuclear Warfare Offensive Capability
	//
	// Z - Inconclusive Analysis.
	NucCap param.Opt[string] `json:"nucCap,omitzero"`
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
	// Unique identifier of the LZ record from the originating system.
	OrigLzID param.Opt[string] `json:"origLzId,omitzero"`
	// Unique identifier of the Site record from the originating system.
	OrigSiteID param.Opt[string] `json:"origSiteID,omitzero"`
	// The O-suffix associated with this site. The O-suffix is a five-character
	// alpha/numeric system used to identify a site, or demographic area, within an
	// installation. The Installation Basic Encyclopedia (beNumber), in conjunction
	// with the O-suffix, uniquely identifies the Site. The Installation beNumber and
	// oSuffix are also used in conjunction with the catCode to classify the function
	// or purpose of the facility.
	Osuffix param.Opt[string] `json:"osuffix,omitzero"`
	// Site number of a specific electronic site or its associated equipment.
	Pin param.Opt[string] `json:"pin,omitzero"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv param.Opt[string] `json:"polSubdiv,omitzero"`
	// Indicates whether the facility is in or outside of a populated area. True, the
	// facility is in or within 5 NM of a populated area. False, the facility is
	// outside a populated area.
	PopArea param.Opt[bool] `json:"popArea,omitzero"`
	// Indicates the distance to nearest populated area (over 1,000 people) in nautical
	// miles.
	PopAreaProx param.Opt[float64] `json:"popAreaProx,omitzero"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs.
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
	// The reference documentation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc param.Opt[string] `json:"referenceDoc,omitzero"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd param.Opt[string] `json:"resProd,omitzero"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency, in ISO8601 UTC format. This date cannot be
	// greater than the current date.
	ReviewDate param.Opt[time.Time] `json:"reviewDate,omitzero" format:"date"`
	// The number of runways at the site, if applicable.
	Runways param.Opt[int64] `json:"runways,omitzero"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element ident.
	SymCode param.Opt[string] `json:"symCode,omitzero"`
	// The type of this site (AIRBASE, AIRFIELD, AIRPORT, NAVAL STATION, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// The use authorization type of this site (e.g MILITARY, CIVIL, JOINT-USE, etc.).
	Usage param.Opt[string] `json:"usage,omitzero"`
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
	// Maximum expected height of the vegetation in the Landing Zone (LZ), in meters.
	VegHt param.Opt[float64] `json:"vegHt,omitzero"`
	// The predominant vegetation found in the Landing Zone (LZ). The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	VegType param.Opt[string] `json:"vegType,omitzero"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r SiteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SiteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteNewParams) UnmarshalJSON(data []byte) error {
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
type SiteNewParamsDataMode string

const (
	SiteNewParamsDataModeReal      SiteNewParamsDataMode = "REAL"
	SiteNewParamsDataModeTest      SiteNewParamsDataMode = "TEST"
	SiteNewParamsDataModeSimulated SiteNewParamsDataMode = "SIMULATED"
	SiteNewParamsDataModeExercise  SiteNewParamsDataMode = "EXERCISE"
)

type SiteUpdateParams struct {
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
	DataMode SiteUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The name of this site.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Indicates the function or mission of an entity, which that entity may or may not
	// be engaged in at any particular time. Typically refers to a unit, organization,
	// or installation/site performing a specific function or mission such as a
	// redistribution center or naval shipyard. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Activity param.Opt[string] `json:"activity,omitzero"`
	// Air Defense District (ADD) or Air Defense Area (ADA) in which the geographic
	// coordinates reside.
	AirDefArea param.Opt[string] `json:"airDefArea,omitzero"`
	// The DoD Standard Country Code designator for the country or political entity to
	// which the site owes its allegiance. This field will be set to "OTHR" if the
	// source value does not match a UDL Country code value (ISO-3166-ALPHA-2).
	Allegiance param.Opt[string] `json:"allegiance,omitzero"`
	// Specifies an alternate allegiance code if the data provider code is not part of
	// an official Country Code standard such as ISO-3166 or FIPS. This field will be
	// set to the value provided by the source and should be used for all Queries
	// specifying allegiance.
	AltAllegiance param.Opt[string] `json:"altAllegiance,omitzero"`
	// The Basic Encyclopedia Number associated with the Site. Uniquely identifies the
	// installation of a site. The beNumber is generated based on the value input for
	// the COORD to determine the appropriate World Aeronautical Chart (WAC) location
	// identifier, the system assigned record originator and a one-up-number.
	BeNumber param.Opt[string] `json:"beNumber,omitzero"`
	// The category code that represents the associated site purpose within the target
	// system.
	CatCode param.Opt[string] `json:"catCode,omitzero"`
	// Textual Description of Site catCode.
	CatText param.Opt[string] `json:"catText,omitzero"`
	// Indicates the importance of the entity to the OES or MIR system. This data
	// element is restricted to update by DIA (DB-4). Valid values are:
	//
	// 0 - Does not meet criteria above
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
	// Pos. 1-21. Unknown Latitude and Unknown Longitude [000000000U000000000U].
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
	// Ground elevation of the geographic coordinates referenced to (above or below)
	// Mean Sea Level (MSL) vertical datum, in meters.
	ElevMsl param.Opt[float64] `json:"elevMsl,omitzero"`
	// Indicates the confidence level expressed as a percent that a specific geometric
	// spatial element, ELEVATION_MSL linear accuracy, has been vertically positioned
	// to within a specified vertical accuracy.
	ElevMslConfLvl param.Opt[int64] `json:"elevMslConfLvl,omitzero"`
	// Indicates the plus or minus error assessed against the method used to derive the
	// elevation.
	ElevMslDerivAcc param.Opt[float64] `json:"elevMslDerivAcc,omitzero"`
	// Eval represents the Intelligence Confidence Level or the Reliability/degree of
	// confidence that the analyst has assigned to the data within this record. The
	// numerical range is from 1 to 9 with 1 representing the highest confidence level.
	Eval param.Opt[int64] `json:"eval,omitzero"`
	// The Federal Aviation Administration (FAA) Location ID of this site, if
	// applicable.
	Faa param.Opt[string] `json:"faa,omitzero"`
	// Functional Production Area (FPA) under the Shared Production Program (SPP).
	// Producers are defined per country per FPA. The specific usage and enumerations
	// contained in this field may be found in the documentation provided in the
	// referenceDoc field. If referenceDoc not provided, users may consult the data
	// provider.
	Fpa param.Opt[string] `json:"fpa,omitzero"`
	// Principal operational function being performed. The specific usage and
	// enumerations contained in this field may be found in the documentation provided
	// in the referenceDoc field. If referenceDoc not provided, users may consult the
	// data provider.
	FunctPrimary param.Opt[string] `json:"functPrimary,omitzero"`
	// Geographical region code used by the Requirements Management System (RMS) as
	// specified by National Geospatial Agency (NGA) in Flight Information Publications
	// (FIPS) 10-4, Appendix 3 - Country Code and Geographic Region Codes. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	GeoArea param.Opt[string] `json:"geoArea,omitzero"`
	// The distance between Mean Sea Level and a referenced ellipsoid, in meters.
	GeoidalMslSep param.Opt[float64] `json:"geoidalMslSep,omitzero"`
	// Indicates the amount or degree of deviation from the horizontal represented as a
	// percent. Grade is determined by the formula: vertical distance (VD) divided by
	// horizontal distance (HD) times 100. VD is the difference between the highest and
	// lowest elevation within the entity. HD is the linear distance between the
	// highest and lowest elevation.
	Grade param.Opt[int64] `json:"grade,omitzero"`
	// The International Air Transport Association (IATA) code of this site, if
	// applicable.
	Iata param.Opt[string] `json:"iata,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of this site, if
	// applicable.
	Icao param.Opt[string] `json:"icao,omitzero"`
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
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// Unique identifier of the Parent Site record associated with this Site record.
	IDParentSite param.Opt[string] `json:"idParentSite,omitzero"`
	// Indicates the normal usage of the Landing Zone (LZ). Intended as, but not
	// constrained to MIDB Helocopter Landing Area usage value definitions:
	//
	// # AF - Airfield
	//
	// # FD - Field
	//
	// HC - High Crop. 1 meter and over.
	//
	// # HY - Highway
	//
	// # LB - Lake Bed
	//
	// LC - Low Crop. 0-1 meters
	//
	// O - Other. Explain In Remarks.
	//
	// # PD - Paddy
	//
	// # PK - Park
	//
	// # PS - Pasture
	//
	// # RB - Riverbed
	//
	// # SP - Sport Field
	//
	// # U - Unknown
	//
	// Z - Inconclusive Analysis.
	LzUsage param.Opt[string] `json:"lzUsage,omitzero"`
	// The length of the longest runway at this site, if applicable, in meters.
	MaxRunwayLength param.Opt[int64] `json:"maxRunwayLength,omitzero"`
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
	// Indicates the principal type of mission that an entity is organized and equipped
	// to perform. The specific usage and enumerations contained in this field may be
	// found in the documentation provided in the referenceDoc field. If referenceDoc
	// not provided, users may consult the data provider.
	MsnPrimary param.Opt[string] `json:"msnPrimary,omitzero"`
	// Indicates the principal specialty type of mission that an entity is organized
	// and equipped to perform. The specific usage and enumerations contained in this
	// field may be found in the documentation provided in the referenceDoc field. If
	// referenceDoc not provided, users may consult the data provider.
	MsnPrimarySpec param.Opt[string] `json:"msnPrimarySpec,omitzero"`
	// Optional notes/comments for the site.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// A sites ability to conduct nuclear warfare. Valid Values are:
	//
	// # A - Nuclear Ammo Or Warheads Available
	//
	// # N - No Nuclear Offense
	//
	// O - Other. Explain in Remarks
	//
	// # U - Unknown
	//
	// # W - Nuclear Weapons Available
	//
	// # Y - Nuclear Warfare Offensive Capability
	//
	// Z - Inconclusive Analysis.
	NucCap param.Opt[string] `json:"nucCap,omitzero"`
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
	// Unique identifier of the LZ record from the originating system.
	OrigLzID param.Opt[string] `json:"origLzId,omitzero"`
	// Unique identifier of the Site record from the originating system.
	OrigSiteID param.Opt[string] `json:"origSiteID,omitzero"`
	// The O-suffix associated with this site. The O-suffix is a five-character
	// alpha/numeric system used to identify a site, or demographic area, within an
	// installation. The Installation Basic Encyclopedia (beNumber), in conjunction
	// with the O-suffix, uniquely identifies the Site. The Installation beNumber and
	// oSuffix are also used in conjunction with the catCode to classify the function
	// or purpose of the facility.
	Osuffix param.Opt[string] `json:"osuffix,omitzero"`
	// Site number of a specific electronic site or its associated equipment.
	Pin param.Opt[string] `json:"pin,omitzero"`
	// Political subdivision in which the geographic coordinates reside. The specific
	// usage and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	PolSubdiv param.Opt[string] `json:"polSubdiv,omitzero"`
	// Indicates whether the facility is in or outside of a populated area. True, the
	// facility is in or within 5 NM of a populated area. False, the facility is
	// outside a populated area.
	PopArea param.Opt[bool] `json:"popArea,omitzero"`
	// Indicates the distance to nearest populated area (over 1,000 people) in nautical
	// miles.
	PopAreaProx param.Opt[float64] `json:"popAreaProx,omitzero"`
	// Validity and currency of the data in the record to be used in conjunction with
	// the other elements in the record as defined by SOPs.
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
	// The reference documentation that specifies the usage and enumerations contained
	// in this record. If referenceDoc not provided, users may consult the data
	// provider.
	ReferenceDoc param.Opt[string] `json:"referenceDoc,omitzero"`
	// Responsible Producer - Organization that is responsible for the maintenance of
	// the record.
	ResProd param.Opt[string] `json:"resProd,omitzero"`
	// Date on which the data in the record was last reviewed by the responsible
	// analyst for accuracy and currency, in ISO8601 UTC format. This date cannot be
	// greater than the current date.
	ReviewDate param.Opt[time.Time] `json:"reviewDate,omitzero" format:"date"`
	// The number of runways at the site, if applicable.
	Runways param.Opt[int64] `json:"runways,omitzero"`
	// A standard scheme for symbol coding enabling the transfer, display and use of
	// symbols and graphics among information systems, as per MIL-STD 2525B, and
	// supported by the element ident.
	SymCode param.Opt[string] `json:"symCode,omitzero"`
	// The type of this site (AIRBASE, AIRFIELD, AIRPORT, NAVAL STATION, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// The use authorization type of this site (e.g MILITARY, CIVIL, JOINT-USE, etc.).
	Usage param.Opt[string] `json:"usage,omitzero"`
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
	// Maximum expected height of the vegetation in the Landing Zone (LZ), in meters.
	VegHt param.Opt[float64] `json:"vegHt,omitzero"`
	// The predominant vegetation found in the Landing Zone (LZ). The specific usage
	// and enumerations contained in this field may be found in the documentation
	// provided in the referenceDoc field. If referenceDoc not provided, users may
	// consult the data provider.
	VegType param.Opt[string] `json:"vegType,omitzero"`
	// World Aeronautical Chart identifier for the area in which a designated place is
	// located.
	Wac param.Opt[string] `json:"wac,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r SiteUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SiteUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteUpdateParams) UnmarshalJSON(data []byte) error {
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
type SiteUpdateParamsDataMode string

const (
	SiteUpdateParamsDataModeReal      SiteUpdateParamsDataMode = "REAL"
	SiteUpdateParamsDataModeTest      SiteUpdateParamsDataMode = "TEST"
	SiteUpdateParamsDataModeSimulated SiteUpdateParamsDataMode = "SIMULATED"
	SiteUpdateParamsDataModeExercise  SiteUpdateParamsDataMode = "EXERCISE"
)

type SiteListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteListParams]'s query parameters as `url.Values`.
func (r SiteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteCountParams]'s query parameters as `url.Values`.
func (r SiteCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteGetParams]'s query parameters as `url.Values`.
func (r SiteGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteTupleParams]'s query parameters as `url.Values`.
func (r SiteTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
