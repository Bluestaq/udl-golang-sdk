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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// AirfieldService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirfieldService] method instead.
type AirfieldService struct {
	Options []option.RequestOption
}

// NewAirfieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAirfieldService(opts ...option.RequestOption) (r AirfieldService) {
	r = AirfieldService{}
	r.Options = opts
	return
}

// Service operation to take a single Airfield as a POST body and ingest into the
// database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *AirfieldService) New(ctx context.Context, body AirfieldNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airfield"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Airfield by its unique ID passed as a path
// parameter.
func (r *AirfieldService) Get(ctx context.Context, id string, query AirfieldGetParams, opts ...option.RequestOption) (res *AirfieldFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfield/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single Airfield. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *AirfieldService) Update(ctx context.Context, id string, body AirfieldUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/airfield/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AirfieldService) List(ctx context.Context, query AirfieldListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AirfieldAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/airfield"
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
func (r *AirfieldService) ListAutoPaging(ctx context.Context, query AirfieldListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AirfieldAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AirfieldService) Count(ctx context.Context, query AirfieldCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/airfield/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AirfieldService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/airfield/queryhelp"
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
func (r *AirfieldService) Tuple(ctx context.Context, query AirfieldTupleParams, opts ...option.RequestOption) (res *[]AirfieldFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/airfield/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Properties and characteristics of an airfield, which includes location, airfield
// codes, suitability codes, and remarks.
type AirfieldAbridged struct {
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
	DataMode AirfieldAbridgedDataMode `json:"dataMode,required"`
	// The name of the airfield.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The airfield activity use type (e.g. Commercial, Airport, Heliport, Gliderport,
	// etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate Airfield identifier provided by source.
	AltAirfieldID string `json:"altAirfieldId"`
	// Alternative names for this airfield.
	AlternativeNames []string `json:"alternativeNames"`
	// The closest city to the location of this airfield.
	City string `json:"city"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// The country name where this airfield is located.
	CountryName string `json:"countryName"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Information regarding daylight saving time as is relevant to the location and
	// operation of this airfield.
	DstInfo string `json:"dstInfo"`
	// Elevation of the airfield above mean sea level, in feet. Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	ElevFt float64 `json:"elevFt"`
	// Elevation of the airfield above mean sea level, in meters. Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	ElevM float64 `json:"elevM"`
	// The Federal Aviation Administration (FAA) location identifier of this airfield.
	Faa string `json:"faa"`
	// Air Force geographic location code of the airfield.
	Geoloc string `json:"geoloc"`
	// Time difference between the location of the airfield and the Greenwich Mean Time
	// (GMT), expressed as +/-HH:MM. Time zones east of Greenwich have positive offsets
	// and time zones west of Greenwich are negative.
	GmtOffset string `json:"gmtOffset"`
	// The host nation code of this airfield, used for non-DoD/FAA locations.
	HostNatCode string `json:"hostNatCode"`
	// The International Aviation Transport Association (IATA) code of the airfield.
	Iata string `json:"iata"`
	// The International Civil Aviation Organization (ICAO) code of the airfield.
	Icao string `json:"icao"`
	// The ID of the parent site.
	IDSite string `json:"idSite"`
	// The URL link to information about airfield.
	InfoURL string `json:"infoURL"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The magnetic declination/variation of the airfield's location from true north,
	// in degrees. Positive values east of true north and negative values west of true
	// north.
	MagDec float64 `json:"magDec"`
	// The length of the longest runway at this airfield in feet.
	MaxRunwayLength int64 `json:"maxRunwayLength"`
	// Applicable miscellaneous codes according to the Airfield Suitability and
	// Restrictions Report (ASRR) for this airfield.
	MiscCodes string `json:"miscCodes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The regional authority of the airfield.
	RegionalAuthority string `json:"regionalAuthority"`
	// Region where the airfield resides.
	RegionName string `json:"regionName"`
	// The number of runways at the site.
	Runways int64 `json:"runways"`
	// The secondary ICAO code for this airfield. Some airfields have two associated
	// ICAO codes, this can occur in cases when a single airfield supports both
	// military and civilian operations.
	SecondaryIcao string `json:"secondaryICAO"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// State or province of the airfield's location.
	State string `json:"state"`
	// The code for the state or province in which this airfield is located. Intended
	// as, but not constrained to, FIPS 10-4 region code designations.
	StateProvinceCode string `json:"stateProvinceCode"`
	// Array of descriptions for given suitability codes. The index of the description
	// corresponds to the position of the letter code in the string provided in the
	// suitabilityCodes field.
	SuitabilityCodeDescs []string `json:"suitabilityCodeDescs"`
	// Associated suitability codes according to the Airfield Suitability and
	// Restrictions Report (ASRR) for this airfield.
	SuitabilityCodes string `json:"suitabilityCodes"`
	// The airfield's World Area Code installation number (WAC-INNR).
	WacInnr string `json:"wacINNR"`
	// Air Mobility Command (AMC) Zone availability Report identifier.
	ZarID string `json:"zarId"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Name                  resp.Field
		Source                resp.Field
		Type                  resp.Field
		ID                    resp.Field
		AltAirfieldID         resp.Field
		AlternativeNames      resp.Field
		City                  resp.Field
		CountryCode           resp.Field
		CountryName           resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DstInfo               resp.Field
		ElevFt                resp.Field
		ElevM                 resp.Field
		Faa                   resp.Field
		Geoloc                resp.Field
		GmtOffset             resp.Field
		HostNatCode           resp.Field
		Iata                  resp.Field
		Icao                  resp.Field
		IDSite                resp.Field
		InfoURL               resp.Field
		Lat                   resp.Field
		Lon                   resp.Field
		MagDec                resp.Field
		MaxRunwayLength       resp.Field
		MiscCodes             resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		RegionalAuthority     resp.Field
		RegionName            resp.Field
		Runways               resp.Field
		SecondaryIcao         resp.Field
		SourceDl              resp.Field
		State                 resp.Field
		StateProvinceCode     resp.Field
		SuitabilityCodeDescs  resp.Field
		SuitabilityCodes      resp.Field
		WacInnr               resp.Field
		ZarID                 resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirfieldAbridged) RawJSON() string { return r.JSON.raw }
func (r *AirfieldAbridged) UnmarshalJSON(data []byte) error {
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
type AirfieldAbridgedDataMode string

const (
	AirfieldAbridgedDataModeReal      AirfieldAbridgedDataMode = "REAL"
	AirfieldAbridgedDataModeTest      AirfieldAbridgedDataMode = "TEST"
	AirfieldAbridgedDataModeSimulated AirfieldAbridgedDataMode = "SIMULATED"
	AirfieldAbridgedDataModeExercise  AirfieldAbridgedDataMode = "EXERCISE"
)

// Properties and characteristics of an airfield, which includes location, airfield
// codes, suitability codes, and remarks.
type AirfieldFull struct {
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
	DataMode AirfieldFullDataMode `json:"dataMode,required"`
	// The name of the airfield.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The airfield activity use type (e.g. Commercial, Airport, Heliport, Gliderport,
	// etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Alternate Airfield identifier provided by source.
	AltAirfieldID string `json:"altAirfieldId"`
	// Alternative names for this airfield.
	AlternativeNames []string `json:"alternativeNames"`
	// The closest city to the location of this airfield.
	City string `json:"city"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// The country name where this airfield is located.
	CountryName string `json:"countryName"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Information regarding daylight saving time as is relevant to the location and
	// operation of this airfield.
	DstInfo string `json:"dstInfo"`
	// Elevation of the airfield above mean sea level, in feet. Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	ElevFt float64 `json:"elevFt"`
	// Elevation of the airfield above mean sea level, in meters. Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	ElevM float64 `json:"elevM"`
	// The Federal Aviation Administration (FAA) location identifier of this airfield.
	Faa string `json:"faa"`
	// Air Force geographic location code of the airfield.
	Geoloc string `json:"geoloc"`
	// Time difference between the location of the airfield and the Greenwich Mean Time
	// (GMT), expressed as +/-HH:MM. Time zones east of Greenwich have positive offsets
	// and time zones west of Greenwich are negative.
	GmtOffset string `json:"gmtOffset"`
	// The host nation code of this airfield, used for non-DoD/FAA locations.
	HostNatCode string `json:"hostNatCode"`
	// The International Aviation Transport Association (IATA) code of the airfield.
	Iata string `json:"iata"`
	// The International Civil Aviation Organization (ICAO) code of the airfield.
	Icao string `json:"icao"`
	// The ID of the parent site.
	IDSite string `json:"idSite"`
	// The URL link to information about airfield.
	InfoURL string `json:"infoURL"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat float64 `json:"lat"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The magnetic declination/variation of the airfield's location from true north,
	// in degrees. Positive values east of true north and negative values west of true
	// north.
	MagDec float64 `json:"magDec"`
	// The length of the longest runway at this airfield in feet.
	MaxRunwayLength int64 `json:"maxRunwayLength"`
	// Applicable miscellaneous codes according to the Airfield Suitability and
	// Restrictions Report (ASRR) for this airfield.
	MiscCodes string `json:"miscCodes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The regional authority of the airfield.
	RegionalAuthority string `json:"regionalAuthority"`
	// Region where the airfield resides.
	RegionName string `json:"regionName"`
	// The number of runways at the site.
	Runways int64 `json:"runways"`
	// The secondary ICAO code for this airfield. Some airfields have two associated
	// ICAO codes, this can occur in cases when a single airfield supports both
	// military and civilian operations.
	SecondaryIcao string `json:"secondaryICAO"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// State or province of the airfield's location.
	State string `json:"state"`
	// The code for the state or province in which this airfield is located. Intended
	// as, but not constrained to, FIPS 10-4 region code designations.
	StateProvinceCode string `json:"stateProvinceCode"`
	// Array of descriptions for given suitability codes. The index of the description
	// corresponds to the position of the letter code in the string provided in the
	// suitabilityCodes field.
	SuitabilityCodeDescs []string `json:"suitabilityCodeDescs"`
	// Associated suitability codes according to the Airfield Suitability and
	// Restrictions Report (ASRR) for this airfield.
	SuitabilityCodes string `json:"suitabilityCodes"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The airfield's World Area Code installation number (WAC-INNR).
	WacInnr string `json:"wacINNR"`
	// Air Mobility Command (AMC) Zone availability Report identifier.
	ZarID string `json:"zarId"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Name                  resp.Field
		Source                resp.Field
		Type                  resp.Field
		ID                    resp.Field
		AltAirfieldID         resp.Field
		AlternativeNames      resp.Field
		City                  resp.Field
		CountryCode           resp.Field
		CountryName           resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DstInfo               resp.Field
		ElevFt                resp.Field
		ElevM                 resp.Field
		Faa                   resp.Field
		Geoloc                resp.Field
		GmtOffset             resp.Field
		HostNatCode           resp.Field
		Iata                  resp.Field
		Icao                  resp.Field
		IDSite                resp.Field
		InfoURL               resp.Field
		Lat                   resp.Field
		Lon                   resp.Field
		MagDec                resp.Field
		MaxRunwayLength       resp.Field
		MiscCodes             resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		RegionalAuthority     resp.Field
		RegionName            resp.Field
		Runways               resp.Field
		SecondaryIcao         resp.Field
		SourceDl              resp.Field
		State                 resp.Field
		StateProvinceCode     resp.Field
		SuitabilityCodeDescs  resp.Field
		SuitabilityCodes      resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		WacInnr               resp.Field
		ZarID                 resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AirfieldFull) RawJSON() string { return r.JSON.raw }
func (r *AirfieldFull) UnmarshalJSON(data []byte) error {
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
type AirfieldFullDataMode string

const (
	AirfieldFullDataModeReal      AirfieldFullDataMode = "REAL"
	AirfieldFullDataModeTest      AirfieldFullDataMode = "TEST"
	AirfieldFullDataModeSimulated AirfieldFullDataMode = "SIMULATED"
	AirfieldFullDataModeExercise  AirfieldFullDataMode = "EXERCISE"
)

type AirfieldNewParams struct {
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
	DataMode AirfieldNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The name of the airfield.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The airfield activity use type (e.g. Commercial, Airport, Heliport, Gliderport,
	// etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate Airfield identifier provided by source.
	AltAirfieldID param.Opt[string] `json:"altAirfieldId,omitzero"`
	// The closest city to the location of this airfield.
	City param.Opt[string] `json:"city,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// The country name where this airfield is located.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// Information regarding daylight saving time as is relevant to the location and
	// operation of this airfield.
	DstInfo param.Opt[string] `json:"dstInfo,omitzero"`
	// Elevation of the airfield above mean sea level, in feet. Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	ElevFt param.Opt[float64] `json:"elevFt,omitzero"`
	// Elevation of the airfield above mean sea level, in meters. Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	ElevM param.Opt[float64] `json:"elevM,omitzero"`
	// The Federal Aviation Administration (FAA) location identifier of this airfield.
	Faa param.Opt[string] `json:"faa,omitzero"`
	// Air Force geographic location code of the airfield.
	Geoloc param.Opt[string] `json:"geoloc,omitzero"`
	// Time difference between the location of the airfield and the Greenwich Mean Time
	// (GMT), expressed as +/-HH:MM. Time zones east of Greenwich have positive offsets
	// and time zones west of Greenwich are negative.
	GmtOffset param.Opt[string] `json:"gmtOffset,omitzero"`
	// The host nation code of this airfield, used for non-DoD/FAA locations.
	HostNatCode param.Opt[string] `json:"hostNatCode,omitzero"`
	// The International Aviation Transport Association (IATA) code of the airfield.
	Iata param.Opt[string] `json:"iata,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the airfield.
	Icao param.Opt[string] `json:"icao,omitzero"`
	// The ID of the parent site.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// The URL link to information about airfield.
	InfoURL param.Opt[string] `json:"infoURL,omitzero"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The magnetic declination/variation of the airfield's location from true north,
	// in degrees. Positive values east of true north and negative values west of true
	// north.
	MagDec param.Opt[float64] `json:"magDec,omitzero"`
	// The length of the longest runway at this airfield in feet.
	MaxRunwayLength param.Opt[int64] `json:"maxRunwayLength,omitzero"`
	// Applicable miscellaneous codes according to the Airfield Suitability and
	// Restrictions Report (ASRR) for this airfield.
	MiscCodes param.Opt[string] `json:"miscCodes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The regional authority of the airfield.
	RegionalAuthority param.Opt[string] `json:"regionalAuthority,omitzero"`
	// Region where the airfield resides.
	RegionName param.Opt[string] `json:"regionName,omitzero"`
	// The number of runways at the site.
	Runways param.Opt[int64] `json:"runways,omitzero"`
	// The secondary ICAO code for this airfield. Some airfields have two associated
	// ICAO codes, this can occur in cases when a single airfield supports both
	// military and civilian operations.
	SecondaryIcao param.Opt[string] `json:"secondaryICAO,omitzero"`
	// State or province of the airfield's location.
	State param.Opt[string] `json:"state,omitzero"`
	// The code for the state or province in which this airfield is located. Intended
	// as, but not constrained to, FIPS 10-4 region code designations.
	StateProvinceCode param.Opt[string] `json:"stateProvinceCode,omitzero"`
	// Associated suitability codes according to the Airfield Suitability and
	// Restrictions Report (ASRR) for this airfield.
	SuitabilityCodes param.Opt[string] `json:"suitabilityCodes,omitzero"`
	// The airfield's World Area Code installation number (WAC-INNR).
	WacInnr param.Opt[string] `json:"wacINNR,omitzero"`
	// Air Mobility Command (AMC) Zone availability Report identifier.
	ZarID param.Opt[string] `json:"zarId,omitzero"`
	// Alternative names for this airfield.
	AlternativeNames []string `json:"alternativeNames,omitzero"`
	// Array of descriptions for given suitability codes. The index of the description
	// corresponds to the position of the letter code in the string provided in the
	// suitabilityCodes field.
	SuitabilityCodeDescs []string `json:"suitabilityCodeDescs,omitzero"`
	paramObj
}

func (r AirfieldNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AirfieldNewParams
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
type AirfieldNewParamsDataMode string

const (
	AirfieldNewParamsDataModeReal      AirfieldNewParamsDataMode = "REAL"
	AirfieldNewParamsDataModeTest      AirfieldNewParamsDataMode = "TEST"
	AirfieldNewParamsDataModeSimulated AirfieldNewParamsDataMode = "SIMULATED"
	AirfieldNewParamsDataModeExercise  AirfieldNewParamsDataMode = "EXERCISE"
)

type AirfieldGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirfieldGetParams]'s query parameters as `url.Values`.
func (r AirfieldGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldUpdateParams struct {
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
	DataMode AirfieldUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The name of the airfield.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The airfield activity use type (e.g. Commercial, Airport, Heliport, Gliderport,
	// etc.).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Alternate Airfield identifier provided by source.
	AltAirfieldID param.Opt[string] `json:"altAirfieldId,omitzero"`
	// The closest city to the location of this airfield.
	City param.Opt[string] `json:"city,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// The country name where this airfield is located.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// Information regarding daylight saving time as is relevant to the location and
	// operation of this airfield.
	DstInfo param.Opt[string] `json:"dstInfo,omitzero"`
	// Elevation of the airfield above mean sea level, in feet. Note: The corresponding
	// equivalent field is not converted by the UDL and may or may not be supplied by
	// the provider. The provider/consumer is responsible for all unit conversions.
	ElevFt param.Opt[float64] `json:"elevFt,omitzero"`
	// Elevation of the airfield above mean sea level, in meters. Note: The
	// corresponding equivalent field is not converted by the UDL and may or may not be
	// supplied by the provider. The provider/consumer is responsible for all unit
	// conversions.
	ElevM param.Opt[float64] `json:"elevM,omitzero"`
	// The Federal Aviation Administration (FAA) location identifier of this airfield.
	Faa param.Opt[string] `json:"faa,omitzero"`
	// Air Force geographic location code of the airfield.
	Geoloc param.Opt[string] `json:"geoloc,omitzero"`
	// Time difference between the location of the airfield and the Greenwich Mean Time
	// (GMT), expressed as +/-HH:MM. Time zones east of Greenwich have positive offsets
	// and time zones west of Greenwich are negative.
	GmtOffset param.Opt[string] `json:"gmtOffset,omitzero"`
	// The host nation code of this airfield, used for non-DoD/FAA locations.
	HostNatCode param.Opt[string] `json:"hostNatCode,omitzero"`
	// The International Aviation Transport Association (IATA) code of the airfield.
	Iata param.Opt[string] `json:"iata,omitzero"`
	// The International Civil Aviation Organization (ICAO) code of the airfield.
	Icao param.Opt[string] `json:"icao,omitzero"`
	// The ID of the parent site.
	IDSite param.Opt[string] `json:"idSite,omitzero"`
	// The URL link to information about airfield.
	InfoURL param.Opt[string] `json:"infoURL,omitzero"`
	// WGS84 latitude of the location, in degrees. -90 to 90 degrees (negative values
	// south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS84 longitude of the location, in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The magnetic declination/variation of the airfield's location from true north,
	// in degrees. Positive values east of true north and negative values west of true
	// north.
	MagDec param.Opt[float64] `json:"magDec,omitzero"`
	// The length of the longest runway at this airfield in feet.
	MaxRunwayLength param.Opt[int64] `json:"maxRunwayLength,omitzero"`
	// Applicable miscellaneous codes according to the Airfield Suitability and
	// Restrictions Report (ASRR) for this airfield.
	MiscCodes param.Opt[string] `json:"miscCodes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The regional authority of the airfield.
	RegionalAuthority param.Opt[string] `json:"regionalAuthority,omitzero"`
	// Region where the airfield resides.
	RegionName param.Opt[string] `json:"regionName,omitzero"`
	// The number of runways at the site.
	Runways param.Opt[int64] `json:"runways,omitzero"`
	// The secondary ICAO code for this airfield. Some airfields have two associated
	// ICAO codes, this can occur in cases when a single airfield supports both
	// military and civilian operations.
	SecondaryIcao param.Opt[string] `json:"secondaryICAO,omitzero"`
	// State or province of the airfield's location.
	State param.Opt[string] `json:"state,omitzero"`
	// The code for the state or province in which this airfield is located. Intended
	// as, but not constrained to, FIPS 10-4 region code designations.
	StateProvinceCode param.Opt[string] `json:"stateProvinceCode,omitzero"`
	// Associated suitability codes according to the Airfield Suitability and
	// Restrictions Report (ASRR) for this airfield.
	SuitabilityCodes param.Opt[string] `json:"suitabilityCodes,omitzero"`
	// The airfield's World Area Code installation number (WAC-INNR).
	WacInnr param.Opt[string] `json:"wacINNR,omitzero"`
	// Air Mobility Command (AMC) Zone availability Report identifier.
	ZarID param.Opt[string] `json:"zarId,omitzero"`
	// Alternative names for this airfield.
	AlternativeNames []string `json:"alternativeNames,omitzero"`
	// Array of descriptions for given suitability codes. The index of the description
	// corresponds to the position of the letter code in the string provided in the
	// suitabilityCodes field.
	SuitabilityCodeDescs []string `json:"suitabilityCodeDescs,omitzero"`
	paramObj
}

func (r AirfieldUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AirfieldUpdateParams
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
type AirfieldUpdateParamsDataMode string

const (
	AirfieldUpdateParamsDataModeReal      AirfieldUpdateParamsDataMode = "REAL"
	AirfieldUpdateParamsDataModeTest      AirfieldUpdateParamsDataMode = "TEST"
	AirfieldUpdateParamsDataModeSimulated AirfieldUpdateParamsDataMode = "SIMULATED"
	AirfieldUpdateParamsDataModeExercise  AirfieldUpdateParamsDataMode = "EXERCISE"
)

type AirfieldListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirfieldListParams]'s query parameters as `url.Values`.
func (r AirfieldListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirfieldCountParams]'s query parameters as `url.Values`.
func (r AirfieldCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirfieldTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AirfieldTupleParams]'s query parameters as `url.Values`.
func (r AirfieldTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
