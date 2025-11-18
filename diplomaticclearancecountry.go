// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

// DiplomaticClearanceCountryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiplomaticClearanceCountryService] method instead.
type DiplomaticClearanceCountryService struct {
	Options []option.RequestOption
}

// NewDiplomaticClearanceCountryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDiplomaticClearanceCountryService(opts ...option.RequestOption) (r DiplomaticClearanceCountryService) {
	r = DiplomaticClearanceCountryService{}
	r.Options = opts
	return
}

// Service operation to take a single diplomaticclearancecountry record as a POST
// body and ingest into the database. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *DiplomaticClearanceCountryService) New(ctx context.Context, body DiplomaticClearanceCountryNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/diplomaticclearancecountry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single diplomaticclearancecountry record by its
// unique ID passed as a path parameter.
func (r *DiplomaticClearanceCountryService) Get(ctx context.Context, id string, query DiplomaticClearanceCountryGetParams, opts ...option.RequestOption) (res *DiplomaticClearanceCountryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/diplomaticclearancecountry/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single diplomaticclearancecountry record. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *DiplomaticClearanceCountryService) Update(ctx context.Context, id string, body DiplomaticClearanceCountryUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/diplomaticclearancecountry/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *DiplomaticClearanceCountryService) List(ctx context.Context, query DiplomaticClearanceCountryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[DiplomaticClearanceCountryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/diplomaticclearancecountry"
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
func (r *DiplomaticClearanceCountryService) ListAutoPaging(ctx context.Context, query DiplomaticClearanceCountryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[DiplomaticClearanceCountryListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a diplomaticclearancecountry record specified by the
// passed ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *DiplomaticClearanceCountryService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/diplomaticclearancecountry/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *DiplomaticClearanceCountryService) Count(ctx context.Context, query DiplomaticClearanceCountryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/diplomaticclearancecountry/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// diplomaticclearancecountry records as a POST body and ingest into the database.
// This operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *DiplomaticClearanceCountryService) NewBulk(ctx context.Context, body DiplomaticClearanceCountryNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/diplomaticclearancecountry/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *DiplomaticClearanceCountryService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *DiplomaticClearanceCountryQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/diplomaticclearancecountry/queryhelp"
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
func (r *DiplomaticClearanceCountryService) Tuple(ctx context.Context, query DiplomaticClearanceCountryTupleParams, opts ...option.RequestOption) (res *[]DiplomaticClearanceCountryTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/diplomaticclearancecountry/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple diplomaticclearancecountry records as a POST
// body and ingest into the database. This operation is intended to be used for
// automated feeds into UDL. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *DiplomaticClearanceCountryService) UnvalidatedPublish(ctx context.Context, body DiplomaticClearanceCountryUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-diplomaticclearancecountry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Diplomatic Clearance Country provides information such as entry/exit points,
// requirements, and points of contact for countries diplomatic clearances are
// being created for.
type DiplomaticClearanceCountryGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country for which the
	// diplomatic clearance will be issued. This field should be set to "OTHR" if the
	// source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	DataMode DiplomaticClearanceCountryGetResponseDataMode `json:"dataMode,required"`
	// Last time this country's diplomatic clearance profile information was updated,
	// in ISO 8601 UTC format with millisecond precision.
	LastChangedDate time.Time `json:"lastChangedDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages using the Defense Message System (DMS).
	AcceptsDms bool `json:"acceptsDMS"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via email.
	AcceptsEmail bool `json:"acceptsEmail"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via fax.
	AcceptsFax bool `json:"acceptsFax"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via SIPRNet.
	AcceptsSiprNet bool `json:"acceptsSIPRNet"`
	// The source agency of the diplomatic clearance country data.
	Agency string `json:"agency"`
	// Specifies an alternate country code if the data provider code does not match a
	// UDL Country code value (ISO-3166-ALPHA-2). This field will be set to the value
	// provided by the source and should be used for all Queries specifying a Country
	// Code.
	AltCountryCode string `json:"altCountryCode"`
	// Zulu closing time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	CloseTime string `json:"closeTime"`
	// System generated code used to identify a country.
	CountryID string `json:"countryId"`
	// Name of the country for which the diplomatic clearance will be issued.
	CountryName string `json:"countryName"`
	// Remarks concerning the country for which the diplomatic clearance will be
	// issued.
	CountryRemark string `json:"countryRemark"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryContacts []DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryContact `json:"diplomaticClearanceCountryContacts"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryEntryExitPoints []DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryEntryExitPoint `json:"diplomaticClearanceCountryEntryExitPoints"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryProfiles []DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryProfile `json:"diplomaticClearanceCountryProfiles"`
	// Flag indicating whether a diplomatic clearance profile exists for this country.
	ExistingProfile bool `json:"existingProfile"`
	// Time difference between the location of the country for which the diplomatic
	// clearance will be issued and the Greenwich Mean Time (GMT), expressed as
	// +/-HH:MM. Time zones east of Greenwich have positive offsets and time zones west
	// of Greenwich are negative.
	GmtOffset string `json:"gmtOffset"`
	// Name of this country's diplomatic clearance office.
	OfficeName string `json:"officeName"`
	// Name of the point of contact for this country's diplomatic clearance office.
	OfficePoc string `json:"officePOC"`
	// Remarks concerning this country's diplomatic clearance office.
	OfficeRemark string `json:"officeRemark"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Friday.
	OpenFri bool `json:"openFri"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Monday.
	OpenMon bool `json:"openMon"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Saturday.
	OpenSat bool `json:"openSat"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Sunday.
	OpenSun bool `json:"openSun"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Thursday.
	OpenThu bool `json:"openThu"`
	// Zulu opening time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	OpenTime string `json:"openTime"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Tuesday.
	OpenTue bool `json:"openTue"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Wednesday.
	OpenWed bool `json:"openWed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                     respjson.Field
		CountryCode                               respjson.Field
		DataMode                                  respjson.Field
		LastChangedDate                           respjson.Field
		Source                                    respjson.Field
		ID                                        respjson.Field
		AcceptsDms                                respjson.Field
		AcceptsEmail                              respjson.Field
		AcceptsFax                                respjson.Field
		AcceptsSiprNet                            respjson.Field
		Agency                                    respjson.Field
		AltCountryCode                            respjson.Field
		CloseTime                                 respjson.Field
		CountryID                                 respjson.Field
		CountryName                               respjson.Field
		CountryRemark                             respjson.Field
		CreatedAt                                 respjson.Field
		CreatedBy                                 respjson.Field
		DiplomaticClearanceCountryContacts        respjson.Field
		DiplomaticClearanceCountryEntryExitPoints respjson.Field
		DiplomaticClearanceCountryProfiles        respjson.Field
		ExistingProfile                           respjson.Field
		GmtOffset                                 respjson.Field
		OfficeName                                respjson.Field
		OfficePoc                                 respjson.Field
		OfficeRemark                              respjson.Field
		OpenFri                                   respjson.Field
		OpenMon                                   respjson.Field
		OpenSat                                   respjson.Field
		OpenSun                                   respjson.Field
		OpenThu                                   respjson.Field
		OpenTime                                  respjson.Field
		OpenTue                                   respjson.Field
		OpenWed                                   respjson.Field
		Origin                                    respjson.Field
		OrigNetwork                               respjson.Field
		SourceDl                                  respjson.Field
		UpdatedAt                                 respjson.Field
		UpdatedBy                                 respjson.Field
		ExtraFields                               map[string]respjson.Field
		raw                                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DiplomaticClearanceCountryGetResponse) UnmarshalJSON(data []byte) error {
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
type DiplomaticClearanceCountryGetResponseDataMode string

const (
	DiplomaticClearanceCountryGetResponseDataModeReal      DiplomaticClearanceCountryGetResponseDataMode = "REAL"
	DiplomaticClearanceCountryGetResponseDataModeTest      DiplomaticClearanceCountryGetResponseDataMode = "TEST"
	DiplomaticClearanceCountryGetResponseDataModeSimulated DiplomaticClearanceCountryGetResponseDataMode = "SIMULATED"
	DiplomaticClearanceCountryGetResponseDataModeExercise  DiplomaticClearanceCountryGetResponseDataMode = "EXERCISE"
)

// Collection of contact information for this country.
type DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryContact struct {
	// Phone number for this contact after regular business hours.
	AhNum string `json:"ahNum"`
	// Speed dial code for this contact after regular business hours.
	AhSpdDialCode string `json:"ahSpdDialCode"`
	// Commercial phone number for this contact.
	CommNum string `json:"commNum"`
	// Commercial speed dial code for this contact.
	CommSpdDialCode string `json:"commSpdDialCode"`
	// Identifier of the contact for this country.
	ContactID string `json:"contactId"`
	// Name of the contact for this country.
	ContactName string `json:"contactName"`
	// Remarks about this contact.
	ContactRemark string `json:"contactRemark"`
	// Defense Switched Network (DSN) phone number for this contact.
	DsnNum string `json:"dsnNum"`
	// Defense Switched Network (DSN) speed dial code for this contact.
	DsnSpdDialCode string `json:"dsnSpdDialCode"`
	// Fax number for this contact.
	FaxNum string `json:"faxNum"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure NIPR line.
	NiprNum string `json:"niprNum"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure SIPR line.
	SiprNum string `json:"siprNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AhNum           respjson.Field
		AhSpdDialCode   respjson.Field
		CommNum         respjson.Field
		CommSpdDialCode respjson.Field
		ContactID       respjson.Field
		ContactName     respjson.Field
		ContactRemark   respjson.Field
		DsnNum          respjson.Field
		DsnSpdDialCode  respjson.Field
		FaxNum          respjson.Field
		NiprNum         respjson.Field
		SiprNum         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryContact) RawJSON() string {
	return r.JSON.raw
}
func (r *DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of entry and exit points for this country.
type DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryEntryExitPoint struct {
	// Flag indicating whether this is a point of entry for this country.
	IsEntry bool `json:"isEntry"`
	// Flag indicating whether this is a point of exit for this country.
	IsExit bool `json:"isExit"`
	// Name of this entry/exit point.
	PointName string `json:"pointName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsEntry     respjson.Field
		IsExit      respjson.Field
		PointName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryEntryExitPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryEntryExitPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of diplomatic clearance profile information for this country.
type DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryProfile struct {
	// Remarks concerning aircraft cargo and passenger information for this country
	// profile.
	CargoPaxRemark string `json:"cargoPaxRemark"`
	// Identifier of the associated diplomatic clearance issued by the host country.
	ClearanceID string `json:"clearanceId"`
	// Remarks concerning crew information for this country profile.
	CrewInfoRemark string `json:"crewInfoRemark"`
	// Code denoting the status of the default diplomatic clearance (e.g., A for
	// Approved via APACS, E for Requested via email, etc.).
	DefClearanceStatus string `json:"defClearanceStatus"`
	// Remarks concerning the default entry point for this country.
	DefEntryRemark string `json:"defEntryRemark"`
	// Zulu default entry time for this country expressed in HH:MM format.
	DefEntryTime string `json:"defEntryTime"`
	// Remarks concerning the default exit point for this country.
	DefExitRemark string `json:"defExitRemark"`
	// Zulu default exit time for this country expressed in HH:MM format.
	DefExitTime string `json:"defExitTime"`
	// Remarks concerning flight information for this country profile.
	FltInfoRemark string `json:"fltInfoRemark"`
	// Remarks concerning hazardous material information for this country profile.
	HazInfoRemark string `json:"hazInfoRemark"`
	// Flag indicating whether this is the default landing profile for this country.
	LandDefProf bool `json:"landDefProf"`
	// Amount of lead time required for an aircraft to land in this country. Units need
	// to be specified in the landLeadTimeUnit field.
	LandLeadTime int64 `json:"landLeadTime"`
	// Remarks concerning the landing lead time required for this country.
	LandLeadTimeRemark string `json:"landLeadTimeRemark"`
	// Unit of time specified for the landLeadTime field to indicate the landing lead
	// time required for this country.
	LandLeadTimeUnit string `json:"landLeadTimeUnit"`
	// Amount of time before the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodMinus int64 `json:"landValidPeriodMinus"`
	// Amount of time after the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodPlus int64 `json:"landValidPeriodPlus"`
	// Remarks concerning the valid landing time period for this country.
	LandValidPeriodRemark string `json:"landValidPeriodRemark"`
	// Unit of time specified for the landValidPeriodPlus and landValidPeriodMinus
	// fields to indicate the valid landing period for this country.
	LandValidPeriodUnit string `json:"landValidPeriodUnit"`
	// Flag indicating whether this is the default overfly profile for this country.
	OverflyDefProf bool `json:"overflyDefProf"`
	// Amount of lead time required for an aircraft to enter and fly over the airspace
	// of this country. Units need to be specified in the overflyLeadTimeUnit field.
	OverflyLeadTime int64 `json:"overflyLeadTime"`
	// Remarks concerning the overfly lead time required for this country.
	OverflyLeadTimeRemark string `json:"overflyLeadTimeRemark"`
	// Unit of time specified for the overflyLeadTime field to indicate the overfly
	// lead time required for this country.
	OverflyLeadTimeUnit string `json:"overflyLeadTimeUnit"`
	// Amount of time before the overfly valid period that an aircraft is allowed to
	// fly over this country for this profile. The unit of time should be specified in
	// the overflyValidPeriodUnit field.
	OverflyValidPeriodMinus int64 `json:"overflyValidPeriodMinus"`
	// Amount of time after the overfly valid period that an aircraft is allowed to fly
	// over this country for this profile. The unit of time should be specified in the
	// overflyValidPeriodUnit field.
	OverflyValidPeriodPlus int64 `json:"overflyValidPeriodPlus"`
	// Remarks concerning the valid overfly time period for this country.
	OverflyValidPeriodRemark string `json:"overflyValidPeriodRemark"`
	// Unit of time specified for the overflyValidPeriodPlus and
	// overflyValidPeriodMinus fields to indicate the valid overfly period for this
	// country.
	OverflyValidPeriodUnit string `json:"overflyValidPeriodUnit"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile string `json:"profile"`
	// The agency to which this profile applies.
	ProfileAgency string `json:"profileAgency"`
	// Identifier of the diplomatic clearance country profile.
	ProfileID string `json:"profileId"`
	// Remarks concerning this country profile.
	ProfileRemark string `json:"profileRemark"`
	// Flag indicating whether alternate aircraft names are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAcAltName bool `json:"reqACAltName"`
	// Flag indicating whether all hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqAllHazInfo bool `json:"reqAllHazInfo"`
	// Flag indicating whether standard AMC information is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAmcStdInfo bool `json:"reqAMCStdInfo"`
	// Flag indicating whether a cargo list is required to be reported to the country
	// using this diplomatic clearance profile.
	ReqCargoList bool `json:"reqCargoList"`
	// Flag indicating whether aircraft cargo and passenger information is required to
	// be reported to the country using this diplomatic clearance profile.
	ReqCargoPax bool `json:"reqCargoPax"`
	// Flag indicating whether Class 1 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass1Info bool `json:"reqClass1Info"`
	// Flag indicating whether Class 9 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass9Info bool `json:"reqClass9Info"`
	// Flag indicating whether the number of crew members on a flight is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqCrewComp bool `json:"reqCrewComp"`
	// Flag indicating whether the names, ranks, and positions of crew members are
	// required to be reported to the country using this diplomatic clearance profile.
	ReqCrewDetail bool `json:"reqCrewDetail"`
	// Flag indicating whether crew information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqCrewInfo bool `json:"reqCrewInfo"`
	// Flag indicating whether Division 1.1 hazardous material information is required
	// to be reported to the country using this diplomatic clearance profile.
	ReqDiv1Info bool `json:"reqDiv1Info"`
	// Flag indicating whether distinguished visitors are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqDv bool `json:"reqDV"`
	// Flag indicating whether entry/exit coordinates need to be specified for this
	// diplomatic clearance profile.
	ReqEntryExitCoord bool `json:"reqEntryExitCoord"`
	// Flag indicating whether flight information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltInfo bool `json:"reqFltInfo"`
	// Flag indicating whether a flight plan route is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltPlanRoute bool `json:"reqFltPlanRoute"`
	// Flag indicating whether aviation funding sources are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqFundSource bool `json:"reqFundSource"`
	// Flag indicating whether hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqHazInfo bool `json:"reqHazInfo"`
	// Flag indicating whether this diplomatic clearance profile applies to specific
	// ICAO(s). These specific ICAO(s) should be clarified in the fltInfoRemark field.
	ReqIcao bool `json:"reqICAO"`
	// Flag indicating whether passport information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqPassportInfo bool `json:"reqPassportInfo"`
	// Flag indicating whether ravens are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRaven bool `json:"reqRaven"`
	// Flag indicating whether changes are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRepChange bool `json:"reqRepChange"`
	// Flag indicating whether an aircraft tail number is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqTailNum bool `json:"reqTailNum"`
	// Flag indicating whether weapons information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqWeaponsInfo bool `json:"reqWeaponsInfo"`
	// Flag indicating whether crew reporting is undefined for the country using this
	// diplomatic clearance profile.
	UndefinedCrewReporting bool `json:"undefinedCrewReporting"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CargoPaxRemark           respjson.Field
		ClearanceID              respjson.Field
		CrewInfoRemark           respjson.Field
		DefClearanceStatus       respjson.Field
		DefEntryRemark           respjson.Field
		DefEntryTime             respjson.Field
		DefExitRemark            respjson.Field
		DefExitTime              respjson.Field
		FltInfoRemark            respjson.Field
		HazInfoRemark            respjson.Field
		LandDefProf              respjson.Field
		LandLeadTime             respjson.Field
		LandLeadTimeRemark       respjson.Field
		LandLeadTimeUnit         respjson.Field
		LandValidPeriodMinus     respjson.Field
		LandValidPeriodPlus      respjson.Field
		LandValidPeriodRemark    respjson.Field
		LandValidPeriodUnit      respjson.Field
		OverflyDefProf           respjson.Field
		OverflyLeadTime          respjson.Field
		OverflyLeadTimeRemark    respjson.Field
		OverflyLeadTimeUnit      respjson.Field
		OverflyValidPeriodMinus  respjson.Field
		OverflyValidPeriodPlus   respjson.Field
		OverflyValidPeriodRemark respjson.Field
		OverflyValidPeriodUnit   respjson.Field
		Profile                  respjson.Field
		ProfileAgency            respjson.Field
		ProfileID                respjson.Field
		ProfileRemark            respjson.Field
		ReqAcAltName             respjson.Field
		ReqAllHazInfo            respjson.Field
		ReqAmcStdInfo            respjson.Field
		ReqCargoList             respjson.Field
		ReqCargoPax              respjson.Field
		ReqClass1Info            respjson.Field
		ReqClass9Info            respjson.Field
		ReqCrewComp              respjson.Field
		ReqCrewDetail            respjson.Field
		ReqCrewInfo              respjson.Field
		ReqDiv1Info              respjson.Field
		ReqDv                    respjson.Field
		ReqEntryExitCoord        respjson.Field
		ReqFltInfo               respjson.Field
		ReqFltPlanRoute          respjson.Field
		ReqFundSource            respjson.Field
		ReqHazInfo               respjson.Field
		ReqIcao                  respjson.Field
		ReqPassportInfo          respjson.Field
		ReqRaven                 respjson.Field
		ReqRepChange             respjson.Field
		ReqTailNum               respjson.Field
		ReqWeaponsInfo           respjson.Field
		UndefinedCrewReporting   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryProfile) RawJSON() string {
	return r.JSON.raw
}
func (r *DiplomaticClearanceCountryGetResponseDiplomaticClearanceCountryProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Diplomatic Clearance Country provides information such as entry/exit points,
// requirements, and points of contact for countries diplomatic clearances are
// being created for.
type DiplomaticClearanceCountryListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country for which the
	// diplomatic clearance will be issued. This field should be set to "OTHR" if the
	// source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	DataMode DiplomaticClearanceCountryListResponseDataMode `json:"dataMode,required"`
	// Last time this country's diplomatic clearance profile information was updated,
	// in ISO 8601 UTC format with millisecond precision.
	LastChangedDate time.Time `json:"lastChangedDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages using the Defense Message System (DMS).
	AcceptsDms bool `json:"acceptsDMS"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via email.
	AcceptsEmail bool `json:"acceptsEmail"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via fax.
	AcceptsFax bool `json:"acceptsFax"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via SIPRNet.
	AcceptsSiprNet bool `json:"acceptsSIPRNet"`
	// The source agency of the diplomatic clearance country data.
	Agency string `json:"agency"`
	// Specifies an alternate country code if the data provider code does not match a
	// UDL Country code value (ISO-3166-ALPHA-2). This field will be set to the value
	// provided by the source and should be used for all Queries specifying a Country
	// Code.
	AltCountryCode string `json:"altCountryCode"`
	// Zulu closing time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	CloseTime string `json:"closeTime"`
	// System generated code used to identify a country.
	CountryID string `json:"countryId"`
	// Name of the country for which the diplomatic clearance will be issued.
	CountryName string `json:"countryName"`
	// Remarks concerning the country for which the diplomatic clearance will be
	// issued.
	CountryRemark string `json:"countryRemark"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryContacts []DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryContact `json:"diplomaticClearanceCountryContacts"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryEntryExitPoints []DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryEntryExitPoint `json:"diplomaticClearanceCountryEntryExitPoints"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryProfiles []DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryProfile `json:"diplomaticClearanceCountryProfiles"`
	// Flag indicating whether a diplomatic clearance profile exists for this country.
	ExistingProfile bool `json:"existingProfile"`
	// Time difference between the location of the country for which the diplomatic
	// clearance will be issued and the Greenwich Mean Time (GMT), expressed as
	// +/-HH:MM. Time zones east of Greenwich have positive offsets and time zones west
	// of Greenwich are negative.
	GmtOffset string `json:"gmtOffset"`
	// Name of this country's diplomatic clearance office.
	OfficeName string `json:"officeName"`
	// Name of the point of contact for this country's diplomatic clearance office.
	OfficePoc string `json:"officePOC"`
	// Remarks concerning this country's diplomatic clearance office.
	OfficeRemark string `json:"officeRemark"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Friday.
	OpenFri bool `json:"openFri"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Monday.
	OpenMon bool `json:"openMon"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Saturday.
	OpenSat bool `json:"openSat"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Sunday.
	OpenSun bool `json:"openSun"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Thursday.
	OpenThu bool `json:"openThu"`
	// Zulu opening time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	OpenTime string `json:"openTime"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Tuesday.
	OpenTue bool `json:"openTue"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Wednesday.
	OpenWed bool `json:"openWed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                     respjson.Field
		CountryCode                               respjson.Field
		DataMode                                  respjson.Field
		LastChangedDate                           respjson.Field
		Source                                    respjson.Field
		ID                                        respjson.Field
		AcceptsDms                                respjson.Field
		AcceptsEmail                              respjson.Field
		AcceptsFax                                respjson.Field
		AcceptsSiprNet                            respjson.Field
		Agency                                    respjson.Field
		AltCountryCode                            respjson.Field
		CloseTime                                 respjson.Field
		CountryID                                 respjson.Field
		CountryName                               respjson.Field
		CountryRemark                             respjson.Field
		CreatedAt                                 respjson.Field
		CreatedBy                                 respjson.Field
		DiplomaticClearanceCountryContacts        respjson.Field
		DiplomaticClearanceCountryEntryExitPoints respjson.Field
		DiplomaticClearanceCountryProfiles        respjson.Field
		ExistingProfile                           respjson.Field
		GmtOffset                                 respjson.Field
		OfficeName                                respjson.Field
		OfficePoc                                 respjson.Field
		OfficeRemark                              respjson.Field
		OpenFri                                   respjson.Field
		OpenMon                                   respjson.Field
		OpenSat                                   respjson.Field
		OpenSun                                   respjson.Field
		OpenThu                                   respjson.Field
		OpenTime                                  respjson.Field
		OpenTue                                   respjson.Field
		OpenWed                                   respjson.Field
		Origin                                    respjson.Field
		OrigNetwork                               respjson.Field
		SourceDl                                  respjson.Field
		UpdatedAt                                 respjson.Field
		UpdatedBy                                 respjson.Field
		ExtraFields                               map[string]respjson.Field
		raw                                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryListResponse) RawJSON() string { return r.JSON.raw }
func (r *DiplomaticClearanceCountryListResponse) UnmarshalJSON(data []byte) error {
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
type DiplomaticClearanceCountryListResponseDataMode string

const (
	DiplomaticClearanceCountryListResponseDataModeReal      DiplomaticClearanceCountryListResponseDataMode = "REAL"
	DiplomaticClearanceCountryListResponseDataModeTest      DiplomaticClearanceCountryListResponseDataMode = "TEST"
	DiplomaticClearanceCountryListResponseDataModeSimulated DiplomaticClearanceCountryListResponseDataMode = "SIMULATED"
	DiplomaticClearanceCountryListResponseDataModeExercise  DiplomaticClearanceCountryListResponseDataMode = "EXERCISE"
)

// Collection of contact information for this country.
type DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryContact struct {
	// Phone number for this contact after regular business hours.
	AhNum string `json:"ahNum"`
	// Speed dial code for this contact after regular business hours.
	AhSpdDialCode string `json:"ahSpdDialCode"`
	// Commercial phone number for this contact.
	CommNum string `json:"commNum"`
	// Commercial speed dial code for this contact.
	CommSpdDialCode string `json:"commSpdDialCode"`
	// Identifier of the contact for this country.
	ContactID string `json:"contactId"`
	// Name of the contact for this country.
	ContactName string `json:"contactName"`
	// Remarks about this contact.
	ContactRemark string `json:"contactRemark"`
	// Defense Switched Network (DSN) phone number for this contact.
	DsnNum string `json:"dsnNum"`
	// Defense Switched Network (DSN) speed dial code for this contact.
	DsnSpdDialCode string `json:"dsnSpdDialCode"`
	// Fax number for this contact.
	FaxNum string `json:"faxNum"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure NIPR line.
	NiprNum string `json:"niprNum"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure SIPR line.
	SiprNum string `json:"siprNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AhNum           respjson.Field
		AhSpdDialCode   respjson.Field
		CommNum         respjson.Field
		CommSpdDialCode respjson.Field
		ContactID       respjson.Field
		ContactName     respjson.Field
		ContactRemark   respjson.Field
		DsnNum          respjson.Field
		DsnSpdDialCode  respjson.Field
		FaxNum          respjson.Field
		NiprNum         respjson.Field
		SiprNum         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryContact) RawJSON() string {
	return r.JSON.raw
}
func (r *DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of entry and exit points for this country.
type DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryEntryExitPoint struct {
	// Flag indicating whether this is a point of entry for this country.
	IsEntry bool `json:"isEntry"`
	// Flag indicating whether this is a point of exit for this country.
	IsExit bool `json:"isExit"`
	// Name of this entry/exit point.
	PointName string `json:"pointName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsEntry     respjson.Field
		IsExit      respjson.Field
		PointName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryEntryExitPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryEntryExitPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of diplomatic clearance profile information for this country.
type DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryProfile struct {
	// Remarks concerning aircraft cargo and passenger information for this country
	// profile.
	CargoPaxRemark string `json:"cargoPaxRemark"`
	// Identifier of the associated diplomatic clearance issued by the host country.
	ClearanceID string `json:"clearanceId"`
	// Remarks concerning crew information for this country profile.
	CrewInfoRemark string `json:"crewInfoRemark"`
	// Code denoting the status of the default diplomatic clearance (e.g., A for
	// Approved via APACS, E for Requested via email, etc.).
	DefClearanceStatus string `json:"defClearanceStatus"`
	// Remarks concerning the default entry point for this country.
	DefEntryRemark string `json:"defEntryRemark"`
	// Zulu default entry time for this country expressed in HH:MM format.
	DefEntryTime string `json:"defEntryTime"`
	// Remarks concerning the default exit point for this country.
	DefExitRemark string `json:"defExitRemark"`
	// Zulu default exit time for this country expressed in HH:MM format.
	DefExitTime string `json:"defExitTime"`
	// Remarks concerning flight information for this country profile.
	FltInfoRemark string `json:"fltInfoRemark"`
	// Remarks concerning hazardous material information for this country profile.
	HazInfoRemark string `json:"hazInfoRemark"`
	// Flag indicating whether this is the default landing profile for this country.
	LandDefProf bool `json:"landDefProf"`
	// Amount of lead time required for an aircraft to land in this country. Units need
	// to be specified in the landLeadTimeUnit field.
	LandLeadTime int64 `json:"landLeadTime"`
	// Remarks concerning the landing lead time required for this country.
	LandLeadTimeRemark string `json:"landLeadTimeRemark"`
	// Unit of time specified for the landLeadTime field to indicate the landing lead
	// time required for this country.
	LandLeadTimeUnit string `json:"landLeadTimeUnit"`
	// Amount of time before the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodMinus int64 `json:"landValidPeriodMinus"`
	// Amount of time after the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodPlus int64 `json:"landValidPeriodPlus"`
	// Remarks concerning the valid landing time period for this country.
	LandValidPeriodRemark string `json:"landValidPeriodRemark"`
	// Unit of time specified for the landValidPeriodPlus and landValidPeriodMinus
	// fields to indicate the valid landing period for this country.
	LandValidPeriodUnit string `json:"landValidPeriodUnit"`
	// Flag indicating whether this is the default overfly profile for this country.
	OverflyDefProf bool `json:"overflyDefProf"`
	// Amount of lead time required for an aircraft to enter and fly over the airspace
	// of this country. Units need to be specified in the overflyLeadTimeUnit field.
	OverflyLeadTime int64 `json:"overflyLeadTime"`
	// Remarks concerning the overfly lead time required for this country.
	OverflyLeadTimeRemark string `json:"overflyLeadTimeRemark"`
	// Unit of time specified for the overflyLeadTime field to indicate the overfly
	// lead time required for this country.
	OverflyLeadTimeUnit string `json:"overflyLeadTimeUnit"`
	// Amount of time before the overfly valid period that an aircraft is allowed to
	// fly over this country for this profile. The unit of time should be specified in
	// the overflyValidPeriodUnit field.
	OverflyValidPeriodMinus int64 `json:"overflyValidPeriodMinus"`
	// Amount of time after the overfly valid period that an aircraft is allowed to fly
	// over this country for this profile. The unit of time should be specified in the
	// overflyValidPeriodUnit field.
	OverflyValidPeriodPlus int64 `json:"overflyValidPeriodPlus"`
	// Remarks concerning the valid overfly time period for this country.
	OverflyValidPeriodRemark string `json:"overflyValidPeriodRemark"`
	// Unit of time specified for the overflyValidPeriodPlus and
	// overflyValidPeriodMinus fields to indicate the valid overfly period for this
	// country.
	OverflyValidPeriodUnit string `json:"overflyValidPeriodUnit"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile string `json:"profile"`
	// The agency to which this profile applies.
	ProfileAgency string `json:"profileAgency"`
	// Identifier of the diplomatic clearance country profile.
	ProfileID string `json:"profileId"`
	// Remarks concerning this country profile.
	ProfileRemark string `json:"profileRemark"`
	// Flag indicating whether alternate aircraft names are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAcAltName bool `json:"reqACAltName"`
	// Flag indicating whether all hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqAllHazInfo bool `json:"reqAllHazInfo"`
	// Flag indicating whether standard AMC information is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAmcStdInfo bool `json:"reqAMCStdInfo"`
	// Flag indicating whether a cargo list is required to be reported to the country
	// using this diplomatic clearance profile.
	ReqCargoList bool `json:"reqCargoList"`
	// Flag indicating whether aircraft cargo and passenger information is required to
	// be reported to the country using this diplomatic clearance profile.
	ReqCargoPax bool `json:"reqCargoPax"`
	// Flag indicating whether Class 1 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass1Info bool `json:"reqClass1Info"`
	// Flag indicating whether Class 9 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass9Info bool `json:"reqClass9Info"`
	// Flag indicating whether the number of crew members on a flight is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqCrewComp bool `json:"reqCrewComp"`
	// Flag indicating whether the names, ranks, and positions of crew members are
	// required to be reported to the country using this diplomatic clearance profile.
	ReqCrewDetail bool `json:"reqCrewDetail"`
	// Flag indicating whether crew information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqCrewInfo bool `json:"reqCrewInfo"`
	// Flag indicating whether Division 1.1 hazardous material information is required
	// to be reported to the country using this diplomatic clearance profile.
	ReqDiv1Info bool `json:"reqDiv1Info"`
	// Flag indicating whether distinguished visitors are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqDv bool `json:"reqDV"`
	// Flag indicating whether entry/exit coordinates need to be specified for this
	// diplomatic clearance profile.
	ReqEntryExitCoord bool `json:"reqEntryExitCoord"`
	// Flag indicating whether flight information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltInfo bool `json:"reqFltInfo"`
	// Flag indicating whether a flight plan route is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltPlanRoute bool `json:"reqFltPlanRoute"`
	// Flag indicating whether aviation funding sources are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqFundSource bool `json:"reqFundSource"`
	// Flag indicating whether hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqHazInfo bool `json:"reqHazInfo"`
	// Flag indicating whether this diplomatic clearance profile applies to specific
	// ICAO(s). These specific ICAO(s) should be clarified in the fltInfoRemark field.
	ReqIcao bool `json:"reqICAO"`
	// Flag indicating whether passport information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqPassportInfo bool `json:"reqPassportInfo"`
	// Flag indicating whether ravens are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRaven bool `json:"reqRaven"`
	// Flag indicating whether changes are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRepChange bool `json:"reqRepChange"`
	// Flag indicating whether an aircraft tail number is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqTailNum bool `json:"reqTailNum"`
	// Flag indicating whether weapons information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqWeaponsInfo bool `json:"reqWeaponsInfo"`
	// Flag indicating whether crew reporting is undefined for the country using this
	// diplomatic clearance profile.
	UndefinedCrewReporting bool `json:"undefinedCrewReporting"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CargoPaxRemark           respjson.Field
		ClearanceID              respjson.Field
		CrewInfoRemark           respjson.Field
		DefClearanceStatus       respjson.Field
		DefEntryRemark           respjson.Field
		DefEntryTime             respjson.Field
		DefExitRemark            respjson.Field
		DefExitTime              respjson.Field
		FltInfoRemark            respjson.Field
		HazInfoRemark            respjson.Field
		LandDefProf              respjson.Field
		LandLeadTime             respjson.Field
		LandLeadTimeRemark       respjson.Field
		LandLeadTimeUnit         respjson.Field
		LandValidPeriodMinus     respjson.Field
		LandValidPeriodPlus      respjson.Field
		LandValidPeriodRemark    respjson.Field
		LandValidPeriodUnit      respjson.Field
		OverflyDefProf           respjson.Field
		OverflyLeadTime          respjson.Field
		OverflyLeadTimeRemark    respjson.Field
		OverflyLeadTimeUnit      respjson.Field
		OverflyValidPeriodMinus  respjson.Field
		OverflyValidPeriodPlus   respjson.Field
		OverflyValidPeriodRemark respjson.Field
		OverflyValidPeriodUnit   respjson.Field
		Profile                  respjson.Field
		ProfileAgency            respjson.Field
		ProfileID                respjson.Field
		ProfileRemark            respjson.Field
		ReqAcAltName             respjson.Field
		ReqAllHazInfo            respjson.Field
		ReqAmcStdInfo            respjson.Field
		ReqCargoList             respjson.Field
		ReqCargoPax              respjson.Field
		ReqClass1Info            respjson.Field
		ReqClass9Info            respjson.Field
		ReqCrewComp              respjson.Field
		ReqCrewDetail            respjson.Field
		ReqCrewInfo              respjson.Field
		ReqDiv1Info              respjson.Field
		ReqDv                    respjson.Field
		ReqEntryExitCoord        respjson.Field
		ReqFltInfo               respjson.Field
		ReqFltPlanRoute          respjson.Field
		ReqFundSource            respjson.Field
		ReqHazInfo               respjson.Field
		ReqIcao                  respjson.Field
		ReqPassportInfo          respjson.Field
		ReqRaven                 respjson.Field
		ReqRepChange             respjson.Field
		ReqTailNum               respjson.Field
		ReqWeaponsInfo           respjson.Field
		UndefinedCrewReporting   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryProfile) RawJSON() string {
	return r.JSON.raw
}
func (r *DiplomaticClearanceCountryListResponseDiplomaticClearanceCountryProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DiplomaticClearanceCountryQueryHelpResponse struct {
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
func (r DiplomaticClearanceCountryQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *DiplomaticClearanceCountryQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Diplomatic Clearance Country provides information such as entry/exit points,
// requirements, and points of contact for countries diplomatic clearances are
// being created for.
type DiplomaticClearanceCountryTupleResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country for which the
	// diplomatic clearance will be issued. This field should be set to "OTHR" if the
	// source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	DataMode DiplomaticClearanceCountryTupleResponseDataMode `json:"dataMode,required"`
	// Last time this country's diplomatic clearance profile information was updated,
	// in ISO 8601 UTC format with millisecond precision.
	LastChangedDate time.Time `json:"lastChangedDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages using the Defense Message System (DMS).
	AcceptsDms bool `json:"acceptsDMS"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via email.
	AcceptsEmail bool `json:"acceptsEmail"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via fax.
	AcceptsFax bool `json:"acceptsFax"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via SIPRNet.
	AcceptsSiprNet bool `json:"acceptsSIPRNet"`
	// The source agency of the diplomatic clearance country data.
	Agency string `json:"agency"`
	// Specifies an alternate country code if the data provider code does not match a
	// UDL Country code value (ISO-3166-ALPHA-2). This field will be set to the value
	// provided by the source and should be used for all Queries specifying a Country
	// Code.
	AltCountryCode string `json:"altCountryCode"`
	// Zulu closing time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	CloseTime string `json:"closeTime"`
	// System generated code used to identify a country.
	CountryID string `json:"countryId"`
	// Name of the country for which the diplomatic clearance will be issued.
	CountryName string `json:"countryName"`
	// Remarks concerning the country for which the diplomatic clearance will be
	// issued.
	CountryRemark string `json:"countryRemark"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryContacts []DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryContact `json:"diplomaticClearanceCountryContacts"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryEntryExitPoints []DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryEntryExitPoint `json:"diplomaticClearanceCountryEntryExitPoints"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryProfiles []DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryProfile `json:"diplomaticClearanceCountryProfiles"`
	// Flag indicating whether a diplomatic clearance profile exists for this country.
	ExistingProfile bool `json:"existingProfile"`
	// Time difference between the location of the country for which the diplomatic
	// clearance will be issued and the Greenwich Mean Time (GMT), expressed as
	// +/-HH:MM. Time zones east of Greenwich have positive offsets and time zones west
	// of Greenwich are negative.
	GmtOffset string `json:"gmtOffset"`
	// Name of this country's diplomatic clearance office.
	OfficeName string `json:"officeName"`
	// Name of the point of contact for this country's diplomatic clearance office.
	OfficePoc string `json:"officePOC"`
	// Remarks concerning this country's diplomatic clearance office.
	OfficeRemark string `json:"officeRemark"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Friday.
	OpenFri bool `json:"openFri"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Monday.
	OpenMon bool `json:"openMon"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Saturday.
	OpenSat bool `json:"openSat"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Sunday.
	OpenSun bool `json:"openSun"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Thursday.
	OpenThu bool `json:"openThu"`
	// Zulu opening time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	OpenTime string `json:"openTime"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Tuesday.
	OpenTue bool `json:"openTue"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Wednesday.
	OpenWed bool `json:"openWed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                     respjson.Field
		CountryCode                               respjson.Field
		DataMode                                  respjson.Field
		LastChangedDate                           respjson.Field
		Source                                    respjson.Field
		ID                                        respjson.Field
		AcceptsDms                                respjson.Field
		AcceptsEmail                              respjson.Field
		AcceptsFax                                respjson.Field
		AcceptsSiprNet                            respjson.Field
		Agency                                    respjson.Field
		AltCountryCode                            respjson.Field
		CloseTime                                 respjson.Field
		CountryID                                 respjson.Field
		CountryName                               respjson.Field
		CountryRemark                             respjson.Field
		CreatedAt                                 respjson.Field
		CreatedBy                                 respjson.Field
		DiplomaticClearanceCountryContacts        respjson.Field
		DiplomaticClearanceCountryEntryExitPoints respjson.Field
		DiplomaticClearanceCountryProfiles        respjson.Field
		ExistingProfile                           respjson.Field
		GmtOffset                                 respjson.Field
		OfficeName                                respjson.Field
		OfficePoc                                 respjson.Field
		OfficeRemark                              respjson.Field
		OpenFri                                   respjson.Field
		OpenMon                                   respjson.Field
		OpenSat                                   respjson.Field
		OpenSun                                   respjson.Field
		OpenThu                                   respjson.Field
		OpenTime                                  respjson.Field
		OpenTue                                   respjson.Field
		OpenWed                                   respjson.Field
		Origin                                    respjson.Field
		OrigNetwork                               respjson.Field
		SourceDl                                  respjson.Field
		UpdatedAt                                 respjson.Field
		UpdatedBy                                 respjson.Field
		ExtraFields                               map[string]respjson.Field
		raw                                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *DiplomaticClearanceCountryTupleResponse) UnmarshalJSON(data []byte) error {
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
type DiplomaticClearanceCountryTupleResponseDataMode string

const (
	DiplomaticClearanceCountryTupleResponseDataModeReal      DiplomaticClearanceCountryTupleResponseDataMode = "REAL"
	DiplomaticClearanceCountryTupleResponseDataModeTest      DiplomaticClearanceCountryTupleResponseDataMode = "TEST"
	DiplomaticClearanceCountryTupleResponseDataModeSimulated DiplomaticClearanceCountryTupleResponseDataMode = "SIMULATED"
	DiplomaticClearanceCountryTupleResponseDataModeExercise  DiplomaticClearanceCountryTupleResponseDataMode = "EXERCISE"
)

// Collection of contact information for this country.
type DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryContact struct {
	// Phone number for this contact after regular business hours.
	AhNum string `json:"ahNum"`
	// Speed dial code for this contact after regular business hours.
	AhSpdDialCode string `json:"ahSpdDialCode"`
	// Commercial phone number for this contact.
	CommNum string `json:"commNum"`
	// Commercial speed dial code for this contact.
	CommSpdDialCode string `json:"commSpdDialCode"`
	// Identifier of the contact for this country.
	ContactID string `json:"contactId"`
	// Name of the contact for this country.
	ContactName string `json:"contactName"`
	// Remarks about this contact.
	ContactRemark string `json:"contactRemark"`
	// Defense Switched Network (DSN) phone number for this contact.
	DsnNum string `json:"dsnNum"`
	// Defense Switched Network (DSN) speed dial code for this contact.
	DsnSpdDialCode string `json:"dsnSpdDialCode"`
	// Fax number for this contact.
	FaxNum string `json:"faxNum"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure NIPR line.
	NiprNum string `json:"niprNum"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure SIPR line.
	SiprNum string `json:"siprNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AhNum           respjson.Field
		AhSpdDialCode   respjson.Field
		CommNum         respjson.Field
		CommSpdDialCode respjson.Field
		ContactID       respjson.Field
		ContactName     respjson.Field
		ContactRemark   respjson.Field
		DsnNum          respjson.Field
		DsnSpdDialCode  respjson.Field
		FaxNum          respjson.Field
		NiprNum         respjson.Field
		SiprNum         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryContact) RawJSON() string {
	return r.JSON.raw
}
func (r *DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of entry and exit points for this country.
type DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryEntryExitPoint struct {
	// Flag indicating whether this is a point of entry for this country.
	IsEntry bool `json:"isEntry"`
	// Flag indicating whether this is a point of exit for this country.
	IsExit bool `json:"isExit"`
	// Name of this entry/exit point.
	PointName string `json:"pointName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsEntry     respjson.Field
		IsExit      respjson.Field
		PointName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryEntryExitPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryEntryExitPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of diplomatic clearance profile information for this country.
type DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryProfile struct {
	// Remarks concerning aircraft cargo and passenger information for this country
	// profile.
	CargoPaxRemark string `json:"cargoPaxRemark"`
	// Identifier of the associated diplomatic clearance issued by the host country.
	ClearanceID string `json:"clearanceId"`
	// Remarks concerning crew information for this country profile.
	CrewInfoRemark string `json:"crewInfoRemark"`
	// Code denoting the status of the default diplomatic clearance (e.g., A for
	// Approved via APACS, E for Requested via email, etc.).
	DefClearanceStatus string `json:"defClearanceStatus"`
	// Remarks concerning the default entry point for this country.
	DefEntryRemark string `json:"defEntryRemark"`
	// Zulu default entry time for this country expressed in HH:MM format.
	DefEntryTime string `json:"defEntryTime"`
	// Remarks concerning the default exit point for this country.
	DefExitRemark string `json:"defExitRemark"`
	// Zulu default exit time for this country expressed in HH:MM format.
	DefExitTime string `json:"defExitTime"`
	// Remarks concerning flight information for this country profile.
	FltInfoRemark string `json:"fltInfoRemark"`
	// Remarks concerning hazardous material information for this country profile.
	HazInfoRemark string `json:"hazInfoRemark"`
	// Flag indicating whether this is the default landing profile for this country.
	LandDefProf bool `json:"landDefProf"`
	// Amount of lead time required for an aircraft to land in this country. Units need
	// to be specified in the landLeadTimeUnit field.
	LandLeadTime int64 `json:"landLeadTime"`
	// Remarks concerning the landing lead time required for this country.
	LandLeadTimeRemark string `json:"landLeadTimeRemark"`
	// Unit of time specified for the landLeadTime field to indicate the landing lead
	// time required for this country.
	LandLeadTimeUnit string `json:"landLeadTimeUnit"`
	// Amount of time before the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodMinus int64 `json:"landValidPeriodMinus"`
	// Amount of time after the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodPlus int64 `json:"landValidPeriodPlus"`
	// Remarks concerning the valid landing time period for this country.
	LandValidPeriodRemark string `json:"landValidPeriodRemark"`
	// Unit of time specified for the landValidPeriodPlus and landValidPeriodMinus
	// fields to indicate the valid landing period for this country.
	LandValidPeriodUnit string `json:"landValidPeriodUnit"`
	// Flag indicating whether this is the default overfly profile for this country.
	OverflyDefProf bool `json:"overflyDefProf"`
	// Amount of lead time required for an aircraft to enter and fly over the airspace
	// of this country. Units need to be specified in the overflyLeadTimeUnit field.
	OverflyLeadTime int64 `json:"overflyLeadTime"`
	// Remarks concerning the overfly lead time required for this country.
	OverflyLeadTimeRemark string `json:"overflyLeadTimeRemark"`
	// Unit of time specified for the overflyLeadTime field to indicate the overfly
	// lead time required for this country.
	OverflyLeadTimeUnit string `json:"overflyLeadTimeUnit"`
	// Amount of time before the overfly valid period that an aircraft is allowed to
	// fly over this country for this profile. The unit of time should be specified in
	// the overflyValidPeriodUnit field.
	OverflyValidPeriodMinus int64 `json:"overflyValidPeriodMinus"`
	// Amount of time after the overfly valid period that an aircraft is allowed to fly
	// over this country for this profile. The unit of time should be specified in the
	// overflyValidPeriodUnit field.
	OverflyValidPeriodPlus int64 `json:"overflyValidPeriodPlus"`
	// Remarks concerning the valid overfly time period for this country.
	OverflyValidPeriodRemark string `json:"overflyValidPeriodRemark"`
	// Unit of time specified for the overflyValidPeriodPlus and
	// overflyValidPeriodMinus fields to indicate the valid overfly period for this
	// country.
	OverflyValidPeriodUnit string `json:"overflyValidPeriodUnit"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile string `json:"profile"`
	// The agency to which this profile applies.
	ProfileAgency string `json:"profileAgency"`
	// Identifier of the diplomatic clearance country profile.
	ProfileID string `json:"profileId"`
	// Remarks concerning this country profile.
	ProfileRemark string `json:"profileRemark"`
	// Flag indicating whether alternate aircraft names are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAcAltName bool `json:"reqACAltName"`
	// Flag indicating whether all hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqAllHazInfo bool `json:"reqAllHazInfo"`
	// Flag indicating whether standard AMC information is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAmcStdInfo bool `json:"reqAMCStdInfo"`
	// Flag indicating whether a cargo list is required to be reported to the country
	// using this diplomatic clearance profile.
	ReqCargoList bool `json:"reqCargoList"`
	// Flag indicating whether aircraft cargo and passenger information is required to
	// be reported to the country using this diplomatic clearance profile.
	ReqCargoPax bool `json:"reqCargoPax"`
	// Flag indicating whether Class 1 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass1Info bool `json:"reqClass1Info"`
	// Flag indicating whether Class 9 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass9Info bool `json:"reqClass9Info"`
	// Flag indicating whether the number of crew members on a flight is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqCrewComp bool `json:"reqCrewComp"`
	// Flag indicating whether the names, ranks, and positions of crew members are
	// required to be reported to the country using this diplomatic clearance profile.
	ReqCrewDetail bool `json:"reqCrewDetail"`
	// Flag indicating whether crew information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqCrewInfo bool `json:"reqCrewInfo"`
	// Flag indicating whether Division 1.1 hazardous material information is required
	// to be reported to the country using this diplomatic clearance profile.
	ReqDiv1Info bool `json:"reqDiv1Info"`
	// Flag indicating whether distinguished visitors are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqDv bool `json:"reqDV"`
	// Flag indicating whether entry/exit coordinates need to be specified for this
	// diplomatic clearance profile.
	ReqEntryExitCoord bool `json:"reqEntryExitCoord"`
	// Flag indicating whether flight information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltInfo bool `json:"reqFltInfo"`
	// Flag indicating whether a flight plan route is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltPlanRoute bool `json:"reqFltPlanRoute"`
	// Flag indicating whether aviation funding sources are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqFundSource bool `json:"reqFundSource"`
	// Flag indicating whether hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqHazInfo bool `json:"reqHazInfo"`
	// Flag indicating whether this diplomatic clearance profile applies to specific
	// ICAO(s). These specific ICAO(s) should be clarified in the fltInfoRemark field.
	ReqIcao bool `json:"reqICAO"`
	// Flag indicating whether passport information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqPassportInfo bool `json:"reqPassportInfo"`
	// Flag indicating whether ravens are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRaven bool `json:"reqRaven"`
	// Flag indicating whether changes are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRepChange bool `json:"reqRepChange"`
	// Flag indicating whether an aircraft tail number is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqTailNum bool `json:"reqTailNum"`
	// Flag indicating whether weapons information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqWeaponsInfo bool `json:"reqWeaponsInfo"`
	// Flag indicating whether crew reporting is undefined for the country using this
	// diplomatic clearance profile.
	UndefinedCrewReporting bool `json:"undefinedCrewReporting"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CargoPaxRemark           respjson.Field
		ClearanceID              respjson.Field
		CrewInfoRemark           respjson.Field
		DefClearanceStatus       respjson.Field
		DefEntryRemark           respjson.Field
		DefEntryTime             respjson.Field
		DefExitRemark            respjson.Field
		DefExitTime              respjson.Field
		FltInfoRemark            respjson.Field
		HazInfoRemark            respjson.Field
		LandDefProf              respjson.Field
		LandLeadTime             respjson.Field
		LandLeadTimeRemark       respjson.Field
		LandLeadTimeUnit         respjson.Field
		LandValidPeriodMinus     respjson.Field
		LandValidPeriodPlus      respjson.Field
		LandValidPeriodRemark    respjson.Field
		LandValidPeriodUnit      respjson.Field
		OverflyDefProf           respjson.Field
		OverflyLeadTime          respjson.Field
		OverflyLeadTimeRemark    respjson.Field
		OverflyLeadTimeUnit      respjson.Field
		OverflyValidPeriodMinus  respjson.Field
		OverflyValidPeriodPlus   respjson.Field
		OverflyValidPeriodRemark respjson.Field
		OverflyValidPeriodUnit   respjson.Field
		Profile                  respjson.Field
		ProfileAgency            respjson.Field
		ProfileID                respjson.Field
		ProfileRemark            respjson.Field
		ReqAcAltName             respjson.Field
		ReqAllHazInfo            respjson.Field
		ReqAmcStdInfo            respjson.Field
		ReqCargoList             respjson.Field
		ReqCargoPax              respjson.Field
		ReqClass1Info            respjson.Field
		ReqClass9Info            respjson.Field
		ReqCrewComp              respjson.Field
		ReqCrewDetail            respjson.Field
		ReqCrewInfo              respjson.Field
		ReqDiv1Info              respjson.Field
		ReqDv                    respjson.Field
		ReqEntryExitCoord        respjson.Field
		ReqFltInfo               respjson.Field
		ReqFltPlanRoute          respjson.Field
		ReqFundSource            respjson.Field
		ReqHazInfo               respjson.Field
		ReqIcao                  respjson.Field
		ReqPassportInfo          respjson.Field
		ReqRaven                 respjson.Field
		ReqRepChange             respjson.Field
		ReqTailNum               respjson.Field
		ReqWeaponsInfo           respjson.Field
		UndefinedCrewReporting   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryProfile) RawJSON() string {
	return r.JSON.raw
}
func (r *DiplomaticClearanceCountryTupleResponseDiplomaticClearanceCountryProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DiplomaticClearanceCountryNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country for which the
	// diplomatic clearance will be issued. This field should be set to "OTHR" if the
	// source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	DataMode DiplomaticClearanceCountryNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Last time this country's diplomatic clearance profile information was updated,
	// in ISO 8601 UTC format with millisecond precision.
	LastChangedDate time.Time `json:"lastChangedDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages using the Defense Message System (DMS).
	AcceptsDms param.Opt[bool] `json:"acceptsDMS,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via email.
	AcceptsEmail param.Opt[bool] `json:"acceptsEmail,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via fax.
	AcceptsFax param.Opt[bool] `json:"acceptsFax,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via SIPRNet.
	AcceptsSiprNet param.Opt[bool] `json:"acceptsSIPRNet,omitzero"`
	// The source agency of the diplomatic clearance country data.
	Agency param.Opt[string] `json:"agency,omitzero"`
	// Specifies an alternate country code if the data provider code does not match a
	// UDL Country code value (ISO-3166-ALPHA-2). This field will be set to the value
	// provided by the source and should be used for all Queries specifying a Country
	// Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Zulu closing time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	CloseTime param.Opt[string] `json:"closeTime,omitzero"`
	// System generated code used to identify a country.
	CountryID param.Opt[string] `json:"countryId,omitzero"`
	// Name of the country for which the diplomatic clearance will be issued.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// Remarks concerning the country for which the diplomatic clearance will be
	// issued.
	CountryRemark param.Opt[string] `json:"countryRemark,omitzero"`
	// Flag indicating whether a diplomatic clearance profile exists for this country.
	ExistingProfile param.Opt[bool] `json:"existingProfile,omitzero"`
	// Time difference between the location of the country for which the diplomatic
	// clearance will be issued and the Greenwich Mean Time (GMT), expressed as
	// +/-HH:MM. Time zones east of Greenwich have positive offsets and time zones west
	// of Greenwich are negative.
	GmtOffset param.Opt[string] `json:"gmtOffset,omitzero"`
	// Name of this country's diplomatic clearance office.
	OfficeName param.Opt[string] `json:"officeName,omitzero"`
	// Name of the point of contact for this country's diplomatic clearance office.
	OfficePoc param.Opt[string] `json:"officePOC,omitzero"`
	// Remarks concerning this country's diplomatic clearance office.
	OfficeRemark param.Opt[string] `json:"officeRemark,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Friday.
	OpenFri param.Opt[bool] `json:"openFri,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Monday.
	OpenMon param.Opt[bool] `json:"openMon,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Saturday.
	OpenSat param.Opt[bool] `json:"openSat,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Sunday.
	OpenSun param.Opt[bool] `json:"openSun,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Thursday.
	OpenThu param.Opt[bool] `json:"openThu,omitzero"`
	// Zulu opening time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	OpenTime param.Opt[string] `json:"openTime,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Tuesday.
	OpenTue param.Opt[bool] `json:"openTue,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Wednesday.
	OpenWed param.Opt[bool] `json:"openWed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryContacts []DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryContact `json:"diplomaticClearanceCountryContacts,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryEntryExitPoints []DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryEntryExitPoint `json:"diplomaticClearanceCountryEntryExitPoints,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryProfiles []DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryProfile `json:"diplomaticClearanceCountryProfiles,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryNewParams) UnmarshalJSON(data []byte) error {
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
type DiplomaticClearanceCountryNewParamsDataMode string

const (
	DiplomaticClearanceCountryNewParamsDataModeReal      DiplomaticClearanceCountryNewParamsDataMode = "REAL"
	DiplomaticClearanceCountryNewParamsDataModeTest      DiplomaticClearanceCountryNewParamsDataMode = "TEST"
	DiplomaticClearanceCountryNewParamsDataModeSimulated DiplomaticClearanceCountryNewParamsDataMode = "SIMULATED"
	DiplomaticClearanceCountryNewParamsDataModeExercise  DiplomaticClearanceCountryNewParamsDataMode = "EXERCISE"
)

// Collection of contact information for this country.
type DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryContact struct {
	// Phone number for this contact after regular business hours.
	AhNum param.Opt[string] `json:"ahNum,omitzero"`
	// Speed dial code for this contact after regular business hours.
	AhSpdDialCode param.Opt[string] `json:"ahSpdDialCode,omitzero"`
	// Commercial phone number for this contact.
	CommNum param.Opt[string] `json:"commNum,omitzero"`
	// Commercial speed dial code for this contact.
	CommSpdDialCode param.Opt[string] `json:"commSpdDialCode,omitzero"`
	// Identifier of the contact for this country.
	ContactID param.Opt[string] `json:"contactId,omitzero"`
	// Name of the contact for this country.
	ContactName param.Opt[string] `json:"contactName,omitzero"`
	// Remarks about this contact.
	ContactRemark param.Opt[string] `json:"contactRemark,omitzero"`
	// Defense Switched Network (DSN) phone number for this contact.
	DsnNum param.Opt[string] `json:"dsnNum,omitzero"`
	// Defense Switched Network (DSN) speed dial code for this contact.
	DsnSpdDialCode param.Opt[string] `json:"dsnSpdDialCode,omitzero"`
	// Fax number for this contact.
	FaxNum param.Opt[string] `json:"faxNum,omitzero"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure NIPR line.
	NiprNum param.Opt[string] `json:"niprNum,omitzero"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure SIPR line.
	SiprNum param.Opt[string] `json:"siprNum,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryContact) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of entry and exit points for this country.
type DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryEntryExitPoint struct {
	// Flag indicating whether this is a point of entry for this country.
	IsEntry param.Opt[bool] `json:"isEntry,omitzero"`
	// Flag indicating whether this is a point of exit for this country.
	IsExit param.Opt[bool] `json:"isExit,omitzero"`
	// Name of this entry/exit point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryEntryExitPoint) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryEntryExitPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryEntryExitPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of diplomatic clearance profile information for this country.
type DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryProfile struct {
	// Remarks concerning aircraft cargo and passenger information for this country
	// profile.
	CargoPaxRemark param.Opt[string] `json:"cargoPaxRemark,omitzero"`
	// Identifier of the associated diplomatic clearance issued by the host country.
	ClearanceID param.Opt[string] `json:"clearanceId,omitzero"`
	// Remarks concerning crew information for this country profile.
	CrewInfoRemark param.Opt[string] `json:"crewInfoRemark,omitzero"`
	// Code denoting the status of the default diplomatic clearance (e.g., A for
	// Approved via APACS, E for Requested via email, etc.).
	DefClearanceStatus param.Opt[string] `json:"defClearanceStatus,omitzero"`
	// Remarks concerning the default entry point for this country.
	DefEntryRemark param.Opt[string] `json:"defEntryRemark,omitzero"`
	// Zulu default entry time for this country expressed in HH:MM format.
	DefEntryTime param.Opt[string] `json:"defEntryTime,omitzero"`
	// Remarks concerning the default exit point for this country.
	DefExitRemark param.Opt[string] `json:"defExitRemark,omitzero"`
	// Zulu default exit time for this country expressed in HH:MM format.
	DefExitTime param.Opt[string] `json:"defExitTime,omitzero"`
	// Remarks concerning flight information for this country profile.
	FltInfoRemark param.Opt[string] `json:"fltInfoRemark,omitzero"`
	// Remarks concerning hazardous material information for this country profile.
	HazInfoRemark param.Opt[string] `json:"hazInfoRemark,omitzero"`
	// Flag indicating whether this is the default landing profile for this country.
	LandDefProf param.Opt[bool] `json:"landDefProf,omitzero"`
	// Amount of lead time required for an aircraft to land in this country. Units need
	// to be specified in the landLeadTimeUnit field.
	LandLeadTime param.Opt[int64] `json:"landLeadTime,omitzero"`
	// Remarks concerning the landing lead time required for this country.
	LandLeadTimeRemark param.Opt[string] `json:"landLeadTimeRemark,omitzero"`
	// Unit of time specified for the landLeadTime field to indicate the landing lead
	// time required for this country.
	LandLeadTimeUnit param.Opt[string] `json:"landLeadTimeUnit,omitzero"`
	// Amount of time before the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodMinus param.Opt[int64] `json:"landValidPeriodMinus,omitzero"`
	// Amount of time after the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodPlus param.Opt[int64] `json:"landValidPeriodPlus,omitzero"`
	// Remarks concerning the valid landing time period for this country.
	LandValidPeriodRemark param.Opt[string] `json:"landValidPeriodRemark,omitzero"`
	// Unit of time specified for the landValidPeriodPlus and landValidPeriodMinus
	// fields to indicate the valid landing period for this country.
	LandValidPeriodUnit param.Opt[string] `json:"landValidPeriodUnit,omitzero"`
	// Flag indicating whether this is the default overfly profile for this country.
	OverflyDefProf param.Opt[bool] `json:"overflyDefProf,omitzero"`
	// Amount of lead time required for an aircraft to enter and fly over the airspace
	// of this country. Units need to be specified in the overflyLeadTimeUnit field.
	OverflyLeadTime param.Opt[int64] `json:"overflyLeadTime,omitzero"`
	// Remarks concerning the overfly lead time required for this country.
	OverflyLeadTimeRemark param.Opt[string] `json:"overflyLeadTimeRemark,omitzero"`
	// Unit of time specified for the overflyLeadTime field to indicate the overfly
	// lead time required for this country.
	OverflyLeadTimeUnit param.Opt[string] `json:"overflyLeadTimeUnit,omitzero"`
	// Amount of time before the overfly valid period that an aircraft is allowed to
	// fly over this country for this profile. The unit of time should be specified in
	// the overflyValidPeriodUnit field.
	OverflyValidPeriodMinus param.Opt[int64] `json:"overflyValidPeriodMinus,omitzero"`
	// Amount of time after the overfly valid period that an aircraft is allowed to fly
	// over this country for this profile. The unit of time should be specified in the
	// overflyValidPeriodUnit field.
	OverflyValidPeriodPlus param.Opt[int64] `json:"overflyValidPeriodPlus,omitzero"`
	// Remarks concerning the valid overfly time period for this country.
	OverflyValidPeriodRemark param.Opt[string] `json:"overflyValidPeriodRemark,omitzero"`
	// Unit of time specified for the overflyValidPeriodPlus and
	// overflyValidPeriodMinus fields to indicate the valid overfly period for this
	// country.
	OverflyValidPeriodUnit param.Opt[string] `json:"overflyValidPeriodUnit,omitzero"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile param.Opt[string] `json:"profile,omitzero"`
	// The agency to which this profile applies.
	ProfileAgency param.Opt[string] `json:"profileAgency,omitzero"`
	// Identifier of the diplomatic clearance country profile.
	ProfileID param.Opt[string] `json:"profileId,omitzero"`
	// Remarks concerning this country profile.
	ProfileRemark param.Opt[string] `json:"profileRemark,omitzero"`
	// Flag indicating whether alternate aircraft names are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAcAltName param.Opt[bool] `json:"reqACAltName,omitzero"`
	// Flag indicating whether all hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqAllHazInfo param.Opt[bool] `json:"reqAllHazInfo,omitzero"`
	// Flag indicating whether standard AMC information is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAmcStdInfo param.Opt[bool] `json:"reqAMCStdInfo,omitzero"`
	// Flag indicating whether a cargo list is required to be reported to the country
	// using this diplomatic clearance profile.
	ReqCargoList param.Opt[bool] `json:"reqCargoList,omitzero"`
	// Flag indicating whether aircraft cargo and passenger information is required to
	// be reported to the country using this diplomatic clearance profile.
	ReqCargoPax param.Opt[bool] `json:"reqCargoPax,omitzero"`
	// Flag indicating whether Class 1 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass1Info param.Opt[bool] `json:"reqClass1Info,omitzero"`
	// Flag indicating whether Class 9 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass9Info param.Opt[bool] `json:"reqClass9Info,omitzero"`
	// Flag indicating whether the number of crew members on a flight is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqCrewComp param.Opt[bool] `json:"reqCrewComp,omitzero"`
	// Flag indicating whether the names, ranks, and positions of crew members are
	// required to be reported to the country using this diplomatic clearance profile.
	ReqCrewDetail param.Opt[bool] `json:"reqCrewDetail,omitzero"`
	// Flag indicating whether crew information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqCrewInfo param.Opt[bool] `json:"reqCrewInfo,omitzero"`
	// Flag indicating whether Division 1.1 hazardous material information is required
	// to be reported to the country using this diplomatic clearance profile.
	ReqDiv1Info param.Opt[bool] `json:"reqDiv1Info,omitzero"`
	// Flag indicating whether distinguished visitors are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqDv param.Opt[bool] `json:"reqDV,omitzero"`
	// Flag indicating whether entry/exit coordinates need to be specified for this
	// diplomatic clearance profile.
	ReqEntryExitCoord param.Opt[bool] `json:"reqEntryExitCoord,omitzero"`
	// Flag indicating whether flight information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltInfo param.Opt[bool] `json:"reqFltInfo,omitzero"`
	// Flag indicating whether a flight plan route is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltPlanRoute param.Opt[bool] `json:"reqFltPlanRoute,omitzero"`
	// Flag indicating whether aviation funding sources are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqFundSource param.Opt[bool] `json:"reqFundSource,omitzero"`
	// Flag indicating whether hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqHazInfo param.Opt[bool] `json:"reqHazInfo,omitzero"`
	// Flag indicating whether this diplomatic clearance profile applies to specific
	// ICAO(s). These specific ICAO(s) should be clarified in the fltInfoRemark field.
	ReqIcao param.Opt[bool] `json:"reqICAO,omitzero"`
	// Flag indicating whether passport information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqPassportInfo param.Opt[bool] `json:"reqPassportInfo,omitzero"`
	// Flag indicating whether ravens are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRaven param.Opt[bool] `json:"reqRaven,omitzero"`
	// Flag indicating whether changes are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRepChange param.Opt[bool] `json:"reqRepChange,omitzero"`
	// Flag indicating whether an aircraft tail number is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqTailNum param.Opt[bool] `json:"reqTailNum,omitzero"`
	// Flag indicating whether weapons information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqWeaponsInfo param.Opt[bool] `json:"reqWeaponsInfo,omitzero"`
	// Flag indicating whether crew reporting is undefined for the country using this
	// diplomatic clearance profile.
	UndefinedCrewReporting param.Opt[bool] `json:"undefinedCrewReporting,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryProfile) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DiplomaticClearanceCountryGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DiplomaticClearanceCountryGetParams]'s query parameters as
// `url.Values`.
func (r DiplomaticClearanceCountryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DiplomaticClearanceCountryUpdateParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country for which the
	// diplomatic clearance will be issued. This field should be set to "OTHR" if the
	// source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	DataMode DiplomaticClearanceCountryUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Last time this country's diplomatic clearance profile information was updated,
	// in ISO 8601 UTC format with millisecond precision.
	LastChangedDate time.Time `json:"lastChangedDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages using the Defense Message System (DMS).
	AcceptsDms param.Opt[bool] `json:"acceptsDMS,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via email.
	AcceptsEmail param.Opt[bool] `json:"acceptsEmail,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via fax.
	AcceptsFax param.Opt[bool] `json:"acceptsFax,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via SIPRNet.
	AcceptsSiprNet param.Opt[bool] `json:"acceptsSIPRNet,omitzero"`
	// The source agency of the diplomatic clearance country data.
	Agency param.Opt[string] `json:"agency,omitzero"`
	// Specifies an alternate country code if the data provider code does not match a
	// UDL Country code value (ISO-3166-ALPHA-2). This field will be set to the value
	// provided by the source and should be used for all Queries specifying a Country
	// Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Zulu closing time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	CloseTime param.Opt[string] `json:"closeTime,omitzero"`
	// System generated code used to identify a country.
	CountryID param.Opt[string] `json:"countryId,omitzero"`
	// Name of the country for which the diplomatic clearance will be issued.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// Remarks concerning the country for which the diplomatic clearance will be
	// issued.
	CountryRemark param.Opt[string] `json:"countryRemark,omitzero"`
	// Flag indicating whether a diplomatic clearance profile exists for this country.
	ExistingProfile param.Opt[bool] `json:"existingProfile,omitzero"`
	// Time difference between the location of the country for which the diplomatic
	// clearance will be issued and the Greenwich Mean Time (GMT), expressed as
	// +/-HH:MM. Time zones east of Greenwich have positive offsets and time zones west
	// of Greenwich are negative.
	GmtOffset param.Opt[string] `json:"gmtOffset,omitzero"`
	// Name of this country's diplomatic clearance office.
	OfficeName param.Opt[string] `json:"officeName,omitzero"`
	// Name of the point of contact for this country's diplomatic clearance office.
	OfficePoc param.Opt[string] `json:"officePOC,omitzero"`
	// Remarks concerning this country's diplomatic clearance office.
	OfficeRemark param.Opt[string] `json:"officeRemark,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Friday.
	OpenFri param.Opt[bool] `json:"openFri,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Monday.
	OpenMon param.Opt[bool] `json:"openMon,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Saturday.
	OpenSat param.Opt[bool] `json:"openSat,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Sunday.
	OpenSun param.Opt[bool] `json:"openSun,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Thursday.
	OpenThu param.Opt[bool] `json:"openThu,omitzero"`
	// Zulu opening time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	OpenTime param.Opt[string] `json:"openTime,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Tuesday.
	OpenTue param.Opt[bool] `json:"openTue,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Wednesday.
	OpenWed param.Opt[bool] `json:"openWed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryContacts []DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryContact `json:"diplomaticClearanceCountryContacts,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryEntryExitPoints []DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryEntryExitPoint `json:"diplomaticClearanceCountryEntryExitPoints,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryProfiles []DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryProfile `json:"diplomaticClearanceCountryProfiles,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryUpdateParams) UnmarshalJSON(data []byte) error {
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
type DiplomaticClearanceCountryUpdateParamsDataMode string

const (
	DiplomaticClearanceCountryUpdateParamsDataModeReal      DiplomaticClearanceCountryUpdateParamsDataMode = "REAL"
	DiplomaticClearanceCountryUpdateParamsDataModeTest      DiplomaticClearanceCountryUpdateParamsDataMode = "TEST"
	DiplomaticClearanceCountryUpdateParamsDataModeSimulated DiplomaticClearanceCountryUpdateParamsDataMode = "SIMULATED"
	DiplomaticClearanceCountryUpdateParamsDataModeExercise  DiplomaticClearanceCountryUpdateParamsDataMode = "EXERCISE"
)

// Collection of contact information for this country.
type DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryContact struct {
	// Phone number for this contact after regular business hours.
	AhNum param.Opt[string] `json:"ahNum,omitzero"`
	// Speed dial code for this contact after regular business hours.
	AhSpdDialCode param.Opt[string] `json:"ahSpdDialCode,omitzero"`
	// Commercial phone number for this contact.
	CommNum param.Opt[string] `json:"commNum,omitzero"`
	// Commercial speed dial code for this contact.
	CommSpdDialCode param.Opt[string] `json:"commSpdDialCode,omitzero"`
	// Identifier of the contact for this country.
	ContactID param.Opt[string] `json:"contactId,omitzero"`
	// Name of the contact for this country.
	ContactName param.Opt[string] `json:"contactName,omitzero"`
	// Remarks about this contact.
	ContactRemark param.Opt[string] `json:"contactRemark,omitzero"`
	// Defense Switched Network (DSN) phone number for this contact.
	DsnNum param.Opt[string] `json:"dsnNum,omitzero"`
	// Defense Switched Network (DSN) speed dial code for this contact.
	DsnSpdDialCode param.Opt[string] `json:"dsnSpdDialCode,omitzero"`
	// Fax number for this contact.
	FaxNum param.Opt[string] `json:"faxNum,omitzero"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure NIPR line.
	NiprNum param.Opt[string] `json:"niprNum,omitzero"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure SIPR line.
	SiprNum param.Opt[string] `json:"siprNum,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryContact) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of entry and exit points for this country.
type DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryEntryExitPoint struct {
	// Flag indicating whether this is a point of entry for this country.
	IsEntry param.Opt[bool] `json:"isEntry,omitzero"`
	// Flag indicating whether this is a point of exit for this country.
	IsExit param.Opt[bool] `json:"isExit,omitzero"`
	// Name of this entry/exit point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryEntryExitPoint) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryEntryExitPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryEntryExitPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of diplomatic clearance profile information for this country.
type DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryProfile struct {
	// Remarks concerning aircraft cargo and passenger information for this country
	// profile.
	CargoPaxRemark param.Opt[string] `json:"cargoPaxRemark,omitzero"`
	// Identifier of the associated diplomatic clearance issued by the host country.
	ClearanceID param.Opt[string] `json:"clearanceId,omitzero"`
	// Remarks concerning crew information for this country profile.
	CrewInfoRemark param.Opt[string] `json:"crewInfoRemark,omitzero"`
	// Code denoting the status of the default diplomatic clearance (e.g., A for
	// Approved via APACS, E for Requested via email, etc.).
	DefClearanceStatus param.Opt[string] `json:"defClearanceStatus,omitzero"`
	// Remarks concerning the default entry point for this country.
	DefEntryRemark param.Opt[string] `json:"defEntryRemark,omitzero"`
	// Zulu default entry time for this country expressed in HH:MM format.
	DefEntryTime param.Opt[string] `json:"defEntryTime,omitzero"`
	// Remarks concerning the default exit point for this country.
	DefExitRemark param.Opt[string] `json:"defExitRemark,omitzero"`
	// Zulu default exit time for this country expressed in HH:MM format.
	DefExitTime param.Opt[string] `json:"defExitTime,omitzero"`
	// Remarks concerning flight information for this country profile.
	FltInfoRemark param.Opt[string] `json:"fltInfoRemark,omitzero"`
	// Remarks concerning hazardous material information for this country profile.
	HazInfoRemark param.Opt[string] `json:"hazInfoRemark,omitzero"`
	// Flag indicating whether this is the default landing profile for this country.
	LandDefProf param.Opt[bool] `json:"landDefProf,omitzero"`
	// Amount of lead time required for an aircraft to land in this country. Units need
	// to be specified in the landLeadTimeUnit field.
	LandLeadTime param.Opt[int64] `json:"landLeadTime,omitzero"`
	// Remarks concerning the landing lead time required for this country.
	LandLeadTimeRemark param.Opt[string] `json:"landLeadTimeRemark,omitzero"`
	// Unit of time specified for the landLeadTime field to indicate the landing lead
	// time required for this country.
	LandLeadTimeUnit param.Opt[string] `json:"landLeadTimeUnit,omitzero"`
	// Amount of time before the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodMinus param.Opt[int64] `json:"landValidPeriodMinus,omitzero"`
	// Amount of time after the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodPlus param.Opt[int64] `json:"landValidPeriodPlus,omitzero"`
	// Remarks concerning the valid landing time period for this country.
	LandValidPeriodRemark param.Opt[string] `json:"landValidPeriodRemark,omitzero"`
	// Unit of time specified for the landValidPeriodPlus and landValidPeriodMinus
	// fields to indicate the valid landing period for this country.
	LandValidPeriodUnit param.Opt[string] `json:"landValidPeriodUnit,omitzero"`
	// Flag indicating whether this is the default overfly profile for this country.
	OverflyDefProf param.Opt[bool] `json:"overflyDefProf,omitzero"`
	// Amount of lead time required for an aircraft to enter and fly over the airspace
	// of this country. Units need to be specified in the overflyLeadTimeUnit field.
	OverflyLeadTime param.Opt[int64] `json:"overflyLeadTime,omitzero"`
	// Remarks concerning the overfly lead time required for this country.
	OverflyLeadTimeRemark param.Opt[string] `json:"overflyLeadTimeRemark,omitzero"`
	// Unit of time specified for the overflyLeadTime field to indicate the overfly
	// lead time required for this country.
	OverflyLeadTimeUnit param.Opt[string] `json:"overflyLeadTimeUnit,omitzero"`
	// Amount of time before the overfly valid period that an aircraft is allowed to
	// fly over this country for this profile. The unit of time should be specified in
	// the overflyValidPeriodUnit field.
	OverflyValidPeriodMinus param.Opt[int64] `json:"overflyValidPeriodMinus,omitzero"`
	// Amount of time after the overfly valid period that an aircraft is allowed to fly
	// over this country for this profile. The unit of time should be specified in the
	// overflyValidPeriodUnit field.
	OverflyValidPeriodPlus param.Opt[int64] `json:"overflyValidPeriodPlus,omitzero"`
	// Remarks concerning the valid overfly time period for this country.
	OverflyValidPeriodRemark param.Opt[string] `json:"overflyValidPeriodRemark,omitzero"`
	// Unit of time specified for the overflyValidPeriodPlus and
	// overflyValidPeriodMinus fields to indicate the valid overfly period for this
	// country.
	OverflyValidPeriodUnit param.Opt[string] `json:"overflyValidPeriodUnit,omitzero"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile param.Opt[string] `json:"profile,omitzero"`
	// The agency to which this profile applies.
	ProfileAgency param.Opt[string] `json:"profileAgency,omitzero"`
	// Identifier of the diplomatic clearance country profile.
	ProfileID param.Opt[string] `json:"profileId,omitzero"`
	// Remarks concerning this country profile.
	ProfileRemark param.Opt[string] `json:"profileRemark,omitzero"`
	// Flag indicating whether alternate aircraft names are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAcAltName param.Opt[bool] `json:"reqACAltName,omitzero"`
	// Flag indicating whether all hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqAllHazInfo param.Opt[bool] `json:"reqAllHazInfo,omitzero"`
	// Flag indicating whether standard AMC information is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAmcStdInfo param.Opt[bool] `json:"reqAMCStdInfo,omitzero"`
	// Flag indicating whether a cargo list is required to be reported to the country
	// using this diplomatic clearance profile.
	ReqCargoList param.Opt[bool] `json:"reqCargoList,omitzero"`
	// Flag indicating whether aircraft cargo and passenger information is required to
	// be reported to the country using this diplomatic clearance profile.
	ReqCargoPax param.Opt[bool] `json:"reqCargoPax,omitzero"`
	// Flag indicating whether Class 1 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass1Info param.Opt[bool] `json:"reqClass1Info,omitzero"`
	// Flag indicating whether Class 9 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass9Info param.Opt[bool] `json:"reqClass9Info,omitzero"`
	// Flag indicating whether the number of crew members on a flight is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqCrewComp param.Opt[bool] `json:"reqCrewComp,omitzero"`
	// Flag indicating whether the names, ranks, and positions of crew members are
	// required to be reported to the country using this diplomatic clearance profile.
	ReqCrewDetail param.Opt[bool] `json:"reqCrewDetail,omitzero"`
	// Flag indicating whether crew information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqCrewInfo param.Opt[bool] `json:"reqCrewInfo,omitzero"`
	// Flag indicating whether Division 1.1 hazardous material information is required
	// to be reported to the country using this diplomatic clearance profile.
	ReqDiv1Info param.Opt[bool] `json:"reqDiv1Info,omitzero"`
	// Flag indicating whether distinguished visitors are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqDv param.Opt[bool] `json:"reqDV,omitzero"`
	// Flag indicating whether entry/exit coordinates need to be specified for this
	// diplomatic clearance profile.
	ReqEntryExitCoord param.Opt[bool] `json:"reqEntryExitCoord,omitzero"`
	// Flag indicating whether flight information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltInfo param.Opt[bool] `json:"reqFltInfo,omitzero"`
	// Flag indicating whether a flight plan route is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltPlanRoute param.Opt[bool] `json:"reqFltPlanRoute,omitzero"`
	// Flag indicating whether aviation funding sources are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqFundSource param.Opt[bool] `json:"reqFundSource,omitzero"`
	// Flag indicating whether hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqHazInfo param.Opt[bool] `json:"reqHazInfo,omitzero"`
	// Flag indicating whether this diplomatic clearance profile applies to specific
	// ICAO(s). These specific ICAO(s) should be clarified in the fltInfoRemark field.
	ReqIcao param.Opt[bool] `json:"reqICAO,omitzero"`
	// Flag indicating whether passport information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqPassportInfo param.Opt[bool] `json:"reqPassportInfo,omitzero"`
	// Flag indicating whether ravens are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRaven param.Opt[bool] `json:"reqRaven,omitzero"`
	// Flag indicating whether changes are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRepChange param.Opt[bool] `json:"reqRepChange,omitzero"`
	// Flag indicating whether an aircraft tail number is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqTailNum param.Opt[bool] `json:"reqTailNum,omitzero"`
	// Flag indicating whether weapons information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqWeaponsInfo param.Opt[bool] `json:"reqWeaponsInfo,omitzero"`
	// Flag indicating whether crew reporting is undefined for the country using this
	// diplomatic clearance profile.
	UndefinedCrewReporting param.Opt[bool] `json:"undefinedCrewReporting,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryProfile) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DiplomaticClearanceCountryListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DiplomaticClearanceCountryListParams]'s query parameters as
// `url.Values`.
func (r DiplomaticClearanceCountryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DiplomaticClearanceCountryCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DiplomaticClearanceCountryCountParams]'s query parameters
// as `url.Values`.
func (r DiplomaticClearanceCountryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DiplomaticClearanceCountryNewBulkParams struct {
	Body []DiplomaticClearanceCountryNewBulkParamsBody
	paramObj
}

func (r DiplomaticClearanceCountryNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *DiplomaticClearanceCountryNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Diplomatic Clearance Country provides information such as entry/exit points,
// requirements, and points of contact for countries diplomatic clearances are
// being created for.
//
// The properties ClassificationMarking, CountryCode, DataMode, LastChangedDate,
// Source are required.
type DiplomaticClearanceCountryNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country for which the
	// diplomatic clearance will be issued. This field should be set to "OTHR" if the
	// source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	// Last time this country's diplomatic clearance profile information was updated,
	// in ISO 8601 UTC format with millisecond precision.
	LastChangedDate time.Time `json:"lastChangedDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages using the Defense Message System (DMS).
	AcceptsDms param.Opt[bool] `json:"acceptsDMS,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via email.
	AcceptsEmail param.Opt[bool] `json:"acceptsEmail,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via fax.
	AcceptsFax param.Opt[bool] `json:"acceptsFax,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via SIPRNet.
	AcceptsSiprNet param.Opt[bool] `json:"acceptsSIPRNet,omitzero"`
	// The source agency of the diplomatic clearance country data.
	Agency param.Opt[string] `json:"agency,omitzero"`
	// Specifies an alternate country code if the data provider code does not match a
	// UDL Country code value (ISO-3166-ALPHA-2). This field will be set to the value
	// provided by the source and should be used for all Queries specifying a Country
	// Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Zulu closing time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	CloseTime param.Opt[string] `json:"closeTime,omitzero"`
	// System generated code used to identify a country.
	CountryID param.Opt[string] `json:"countryId,omitzero"`
	// Name of the country for which the diplomatic clearance will be issued.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// Remarks concerning the country for which the diplomatic clearance will be
	// issued.
	CountryRemark param.Opt[string] `json:"countryRemark,omitzero"`
	// Flag indicating whether a diplomatic clearance profile exists for this country.
	ExistingProfile param.Opt[bool] `json:"existingProfile,omitzero"`
	// Time difference between the location of the country for which the diplomatic
	// clearance will be issued and the Greenwich Mean Time (GMT), expressed as
	// +/-HH:MM. Time zones east of Greenwich have positive offsets and time zones west
	// of Greenwich are negative.
	GmtOffset param.Opt[string] `json:"gmtOffset,omitzero"`
	// Name of this country's diplomatic clearance office.
	OfficeName param.Opt[string] `json:"officeName,omitzero"`
	// Name of the point of contact for this country's diplomatic clearance office.
	OfficePoc param.Opt[string] `json:"officePOC,omitzero"`
	// Remarks concerning this country's diplomatic clearance office.
	OfficeRemark param.Opt[string] `json:"officeRemark,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Friday.
	OpenFri param.Opt[bool] `json:"openFri,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Monday.
	OpenMon param.Opt[bool] `json:"openMon,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Saturday.
	OpenSat param.Opt[bool] `json:"openSat,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Sunday.
	OpenSun param.Opt[bool] `json:"openSun,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Thursday.
	OpenThu param.Opt[bool] `json:"openThu,omitzero"`
	// Zulu opening time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	OpenTime param.Opt[string] `json:"openTime,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Tuesday.
	OpenTue param.Opt[bool] `json:"openTue,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Wednesday.
	OpenWed param.Opt[bool] `json:"openWed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryContacts []DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryContact `json:"diplomaticClearanceCountryContacts,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryEntryExitPoints []DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryEntryExitPoint `json:"diplomaticClearanceCountryEntryExitPoints,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryProfiles []DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryProfile `json:"diplomaticClearanceCountryProfiles,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DiplomaticClearanceCountryNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of contact information for this country.
type DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryContact struct {
	// Phone number for this contact after regular business hours.
	AhNum param.Opt[string] `json:"ahNum,omitzero"`
	// Speed dial code for this contact after regular business hours.
	AhSpdDialCode param.Opt[string] `json:"ahSpdDialCode,omitzero"`
	// Commercial phone number for this contact.
	CommNum param.Opt[string] `json:"commNum,omitzero"`
	// Commercial speed dial code for this contact.
	CommSpdDialCode param.Opt[string] `json:"commSpdDialCode,omitzero"`
	// Identifier of the contact for this country.
	ContactID param.Opt[string] `json:"contactId,omitzero"`
	// Name of the contact for this country.
	ContactName param.Opt[string] `json:"contactName,omitzero"`
	// Remarks about this contact.
	ContactRemark param.Opt[string] `json:"contactRemark,omitzero"`
	// Defense Switched Network (DSN) phone number for this contact.
	DsnNum param.Opt[string] `json:"dsnNum,omitzero"`
	// Defense Switched Network (DSN) speed dial code for this contact.
	DsnSpdDialCode param.Opt[string] `json:"dsnSpdDialCode,omitzero"`
	// Fax number for this contact.
	FaxNum param.Opt[string] `json:"faxNum,omitzero"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure NIPR line.
	NiprNum param.Opt[string] `json:"niprNum,omitzero"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure SIPR line.
	SiprNum param.Opt[string] `json:"siprNum,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryContact) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of entry and exit points for this country.
type DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryEntryExitPoint struct {
	// Flag indicating whether this is a point of entry for this country.
	IsEntry param.Opt[bool] `json:"isEntry,omitzero"`
	// Flag indicating whether this is a point of exit for this country.
	IsExit param.Opt[bool] `json:"isExit,omitzero"`
	// Name of this entry/exit point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryEntryExitPoint) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryEntryExitPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryEntryExitPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of diplomatic clearance profile information for this country.
type DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryProfile struct {
	// Remarks concerning aircraft cargo and passenger information for this country
	// profile.
	CargoPaxRemark param.Opt[string] `json:"cargoPaxRemark,omitzero"`
	// Identifier of the associated diplomatic clearance issued by the host country.
	ClearanceID param.Opt[string] `json:"clearanceId,omitzero"`
	// Remarks concerning crew information for this country profile.
	CrewInfoRemark param.Opt[string] `json:"crewInfoRemark,omitzero"`
	// Code denoting the status of the default diplomatic clearance (e.g., A for
	// Approved via APACS, E for Requested via email, etc.).
	DefClearanceStatus param.Opt[string] `json:"defClearanceStatus,omitzero"`
	// Remarks concerning the default entry point for this country.
	DefEntryRemark param.Opt[string] `json:"defEntryRemark,omitzero"`
	// Zulu default entry time for this country expressed in HH:MM format.
	DefEntryTime param.Opt[string] `json:"defEntryTime,omitzero"`
	// Remarks concerning the default exit point for this country.
	DefExitRemark param.Opt[string] `json:"defExitRemark,omitzero"`
	// Zulu default exit time for this country expressed in HH:MM format.
	DefExitTime param.Opt[string] `json:"defExitTime,omitzero"`
	// Remarks concerning flight information for this country profile.
	FltInfoRemark param.Opt[string] `json:"fltInfoRemark,omitzero"`
	// Remarks concerning hazardous material information for this country profile.
	HazInfoRemark param.Opt[string] `json:"hazInfoRemark,omitzero"`
	// Flag indicating whether this is the default landing profile for this country.
	LandDefProf param.Opt[bool] `json:"landDefProf,omitzero"`
	// Amount of lead time required for an aircraft to land in this country. Units need
	// to be specified in the landLeadTimeUnit field.
	LandLeadTime param.Opt[int64] `json:"landLeadTime,omitzero"`
	// Remarks concerning the landing lead time required for this country.
	LandLeadTimeRemark param.Opt[string] `json:"landLeadTimeRemark,omitzero"`
	// Unit of time specified for the landLeadTime field to indicate the landing lead
	// time required for this country.
	LandLeadTimeUnit param.Opt[string] `json:"landLeadTimeUnit,omitzero"`
	// Amount of time before the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodMinus param.Opt[int64] `json:"landValidPeriodMinus,omitzero"`
	// Amount of time after the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodPlus param.Opt[int64] `json:"landValidPeriodPlus,omitzero"`
	// Remarks concerning the valid landing time period for this country.
	LandValidPeriodRemark param.Opt[string] `json:"landValidPeriodRemark,omitzero"`
	// Unit of time specified for the landValidPeriodPlus and landValidPeriodMinus
	// fields to indicate the valid landing period for this country.
	LandValidPeriodUnit param.Opt[string] `json:"landValidPeriodUnit,omitzero"`
	// Flag indicating whether this is the default overfly profile for this country.
	OverflyDefProf param.Opt[bool] `json:"overflyDefProf,omitzero"`
	// Amount of lead time required for an aircraft to enter and fly over the airspace
	// of this country. Units need to be specified in the overflyLeadTimeUnit field.
	OverflyLeadTime param.Opt[int64] `json:"overflyLeadTime,omitzero"`
	// Remarks concerning the overfly lead time required for this country.
	OverflyLeadTimeRemark param.Opt[string] `json:"overflyLeadTimeRemark,omitzero"`
	// Unit of time specified for the overflyLeadTime field to indicate the overfly
	// lead time required for this country.
	OverflyLeadTimeUnit param.Opt[string] `json:"overflyLeadTimeUnit,omitzero"`
	// Amount of time before the overfly valid period that an aircraft is allowed to
	// fly over this country for this profile. The unit of time should be specified in
	// the overflyValidPeriodUnit field.
	OverflyValidPeriodMinus param.Opt[int64] `json:"overflyValidPeriodMinus,omitzero"`
	// Amount of time after the overfly valid period that an aircraft is allowed to fly
	// over this country for this profile. The unit of time should be specified in the
	// overflyValidPeriodUnit field.
	OverflyValidPeriodPlus param.Opt[int64] `json:"overflyValidPeriodPlus,omitzero"`
	// Remarks concerning the valid overfly time period for this country.
	OverflyValidPeriodRemark param.Opt[string] `json:"overflyValidPeriodRemark,omitzero"`
	// Unit of time specified for the overflyValidPeriodPlus and
	// overflyValidPeriodMinus fields to indicate the valid overfly period for this
	// country.
	OverflyValidPeriodUnit param.Opt[string] `json:"overflyValidPeriodUnit,omitzero"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile param.Opt[string] `json:"profile,omitzero"`
	// The agency to which this profile applies.
	ProfileAgency param.Opt[string] `json:"profileAgency,omitzero"`
	// Identifier of the diplomatic clearance country profile.
	ProfileID param.Opt[string] `json:"profileId,omitzero"`
	// Remarks concerning this country profile.
	ProfileRemark param.Opt[string] `json:"profileRemark,omitzero"`
	// Flag indicating whether alternate aircraft names are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAcAltName param.Opt[bool] `json:"reqACAltName,omitzero"`
	// Flag indicating whether all hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqAllHazInfo param.Opt[bool] `json:"reqAllHazInfo,omitzero"`
	// Flag indicating whether standard AMC information is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAmcStdInfo param.Opt[bool] `json:"reqAMCStdInfo,omitzero"`
	// Flag indicating whether a cargo list is required to be reported to the country
	// using this diplomatic clearance profile.
	ReqCargoList param.Opt[bool] `json:"reqCargoList,omitzero"`
	// Flag indicating whether aircraft cargo and passenger information is required to
	// be reported to the country using this diplomatic clearance profile.
	ReqCargoPax param.Opt[bool] `json:"reqCargoPax,omitzero"`
	// Flag indicating whether Class 1 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass1Info param.Opt[bool] `json:"reqClass1Info,omitzero"`
	// Flag indicating whether Class 9 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass9Info param.Opt[bool] `json:"reqClass9Info,omitzero"`
	// Flag indicating whether the number of crew members on a flight is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqCrewComp param.Opt[bool] `json:"reqCrewComp,omitzero"`
	// Flag indicating whether the names, ranks, and positions of crew members are
	// required to be reported to the country using this diplomatic clearance profile.
	ReqCrewDetail param.Opt[bool] `json:"reqCrewDetail,omitzero"`
	// Flag indicating whether crew information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqCrewInfo param.Opt[bool] `json:"reqCrewInfo,omitzero"`
	// Flag indicating whether Division 1.1 hazardous material information is required
	// to be reported to the country using this diplomatic clearance profile.
	ReqDiv1Info param.Opt[bool] `json:"reqDiv1Info,omitzero"`
	// Flag indicating whether distinguished visitors are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqDv param.Opt[bool] `json:"reqDV,omitzero"`
	// Flag indicating whether entry/exit coordinates need to be specified for this
	// diplomatic clearance profile.
	ReqEntryExitCoord param.Opt[bool] `json:"reqEntryExitCoord,omitzero"`
	// Flag indicating whether flight information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltInfo param.Opt[bool] `json:"reqFltInfo,omitzero"`
	// Flag indicating whether a flight plan route is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltPlanRoute param.Opt[bool] `json:"reqFltPlanRoute,omitzero"`
	// Flag indicating whether aviation funding sources are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqFundSource param.Opt[bool] `json:"reqFundSource,omitzero"`
	// Flag indicating whether hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqHazInfo param.Opt[bool] `json:"reqHazInfo,omitzero"`
	// Flag indicating whether this diplomatic clearance profile applies to specific
	// ICAO(s). These specific ICAO(s) should be clarified in the fltInfoRemark field.
	ReqIcao param.Opt[bool] `json:"reqICAO,omitzero"`
	// Flag indicating whether passport information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqPassportInfo param.Opt[bool] `json:"reqPassportInfo,omitzero"`
	// Flag indicating whether ravens are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRaven param.Opt[bool] `json:"reqRaven,omitzero"`
	// Flag indicating whether changes are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRepChange param.Opt[bool] `json:"reqRepChange,omitzero"`
	// Flag indicating whether an aircraft tail number is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqTailNum param.Opt[bool] `json:"reqTailNum,omitzero"`
	// Flag indicating whether weapons information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqWeaponsInfo param.Opt[bool] `json:"reqWeaponsInfo,omitzero"`
	// Flag indicating whether crew reporting is undefined for the country using this
	// diplomatic clearance profile.
	UndefinedCrewReporting param.Opt[bool] `json:"undefinedCrewReporting,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryProfile) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DiplomaticClearanceCountryTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DiplomaticClearanceCountryTupleParams]'s query parameters
// as `url.Values`.
func (r DiplomaticClearanceCountryTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DiplomaticClearanceCountryUnvalidatedPublishParams struct {
	Body []DiplomaticClearanceCountryUnvalidatedPublishParamsBody
	paramObj
}

func (r DiplomaticClearanceCountryUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *DiplomaticClearanceCountryUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Diplomatic Clearance Country provides information such as entry/exit points,
// requirements, and points of contact for countries diplomatic clearances are
// being created for.
//
// The properties ClassificationMarking, CountryCode, DataMode, LastChangedDate,
// Source are required.
type DiplomaticClearanceCountryUnvalidatedPublishParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The DoD Standard Country Code designator for the country for which the
	// diplomatic clearance will be issued. This field should be set to "OTHR" if the
	// source value does not match a UDL country code value (ISO-3166-ALPHA-2).
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
	// Last time this country's diplomatic clearance profile information was updated,
	// in ISO 8601 UTC format with millisecond precision.
	LastChangedDate time.Time `json:"lastChangedDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages using the Defense Message System (DMS).
	AcceptsDms param.Opt[bool] `json:"acceptsDMS,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via email.
	AcceptsEmail param.Opt[bool] `json:"acceptsEmail,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via fax.
	AcceptsFax param.Opt[bool] `json:"acceptsFax,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office can receive
	// messages via SIPRNet.
	AcceptsSiprNet param.Opt[bool] `json:"acceptsSIPRNet,omitzero"`
	// The source agency of the diplomatic clearance country data.
	Agency param.Opt[string] `json:"agency,omitzero"`
	// Specifies an alternate country code if the data provider code does not match a
	// UDL Country code value (ISO-3166-ALPHA-2). This field will be set to the value
	// provided by the source and should be used for all Queries specifying a Country
	// Code.
	AltCountryCode param.Opt[string] `json:"altCountryCode,omitzero"`
	// Zulu closing time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	CloseTime param.Opt[string] `json:"closeTime,omitzero"`
	// System generated code used to identify a country.
	CountryID param.Opt[string] `json:"countryId,omitzero"`
	// Name of the country for which the diplomatic clearance will be issued.
	CountryName param.Opt[string] `json:"countryName,omitzero"`
	// Remarks concerning the country for which the diplomatic clearance will be
	// issued.
	CountryRemark param.Opt[string] `json:"countryRemark,omitzero"`
	// Flag indicating whether a diplomatic clearance profile exists for this country.
	ExistingProfile param.Opt[bool] `json:"existingProfile,omitzero"`
	// Time difference between the location of the country for which the diplomatic
	// clearance will be issued and the Greenwich Mean Time (GMT), expressed as
	// +/-HH:MM. Time zones east of Greenwich have positive offsets and time zones west
	// of Greenwich are negative.
	GmtOffset param.Opt[string] `json:"gmtOffset,omitzero"`
	// Name of this country's diplomatic clearance office.
	OfficeName param.Opt[string] `json:"officeName,omitzero"`
	// Name of the point of contact for this country's diplomatic clearance office.
	OfficePoc param.Opt[string] `json:"officePOC,omitzero"`
	// Remarks concerning this country's diplomatic clearance office.
	OfficeRemark param.Opt[string] `json:"officeRemark,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Friday.
	OpenFri param.Opt[bool] `json:"openFri,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Monday.
	OpenMon param.Opt[bool] `json:"openMon,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Saturday.
	OpenSat param.Opt[bool] `json:"openSat,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Sunday.
	OpenSun param.Opt[bool] `json:"openSun,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Thursday.
	OpenThu param.Opt[bool] `json:"openThu,omitzero"`
	// Zulu opening time of this country's diplomatic clearance office expressed in
	// HH:MM format.
	OpenTime param.Opt[string] `json:"openTime,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Tuesday.
	OpenTue param.Opt[bool] `json:"openTue,omitzero"`
	// Flag indicating whether this country's diplomatic clearance office is open on
	// Wednesday.
	OpenWed param.Opt[bool] `json:"openWed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryContacts []DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryContact `json:"diplomaticClearanceCountryContacts,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryEntryExitPoints []DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryEntryExitPoint `json:"diplomaticClearanceCountryEntryExitPoints,omitzero"`
	// Collection of diplomatic clearance profile information for this country.
	DiplomaticClearanceCountryProfiles []DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryProfile `json:"diplomaticClearanceCountryProfiles,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DiplomaticClearanceCountryUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of contact information for this country.
type DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryContact struct {
	// Phone number for this contact after regular business hours.
	AhNum param.Opt[string] `json:"ahNum,omitzero"`
	// Speed dial code for this contact after regular business hours.
	AhSpdDialCode param.Opt[string] `json:"ahSpdDialCode,omitzero"`
	// Commercial phone number for this contact.
	CommNum param.Opt[string] `json:"commNum,omitzero"`
	// Commercial speed dial code for this contact.
	CommSpdDialCode param.Opt[string] `json:"commSpdDialCode,omitzero"`
	// Identifier of the contact for this country.
	ContactID param.Opt[string] `json:"contactId,omitzero"`
	// Name of the contact for this country.
	ContactName param.Opt[string] `json:"contactName,omitzero"`
	// Remarks about this contact.
	ContactRemark param.Opt[string] `json:"contactRemark,omitzero"`
	// Defense Switched Network (DSN) phone number for this contact.
	DsnNum param.Opt[string] `json:"dsnNum,omitzero"`
	// Defense Switched Network (DSN) speed dial code for this contact.
	DsnSpdDialCode param.Opt[string] `json:"dsnSpdDialCode,omitzero"`
	// Fax number for this contact.
	FaxNum param.Opt[string] `json:"faxNum,omitzero"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure NIPR line.
	NiprNum param.Opt[string] `json:"niprNum,omitzero"`
	// Phone number to contact the Diplomatic Attache Office (DAO) for this country
	// over a secure SIPR line.
	SiprNum param.Opt[string] `json:"siprNum,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryContact) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of entry and exit points for this country.
type DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryEntryExitPoint struct {
	// Flag indicating whether this is a point of entry for this country.
	IsEntry param.Opt[bool] `json:"isEntry,omitzero"`
	// Flag indicating whether this is a point of exit for this country.
	IsExit param.Opt[bool] `json:"isExit,omitzero"`
	// Name of this entry/exit point.
	PointName param.Opt[string] `json:"pointName,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryEntryExitPoint) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryEntryExitPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryEntryExitPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of diplomatic clearance profile information for this country.
type DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryProfile struct {
	// Remarks concerning aircraft cargo and passenger information for this country
	// profile.
	CargoPaxRemark param.Opt[string] `json:"cargoPaxRemark,omitzero"`
	// Identifier of the associated diplomatic clearance issued by the host country.
	ClearanceID param.Opt[string] `json:"clearanceId,omitzero"`
	// Remarks concerning crew information for this country profile.
	CrewInfoRemark param.Opt[string] `json:"crewInfoRemark,omitzero"`
	// Code denoting the status of the default diplomatic clearance (e.g., A for
	// Approved via APACS, E for Requested via email, etc.).
	DefClearanceStatus param.Opt[string] `json:"defClearanceStatus,omitzero"`
	// Remarks concerning the default entry point for this country.
	DefEntryRemark param.Opt[string] `json:"defEntryRemark,omitzero"`
	// Zulu default entry time for this country expressed in HH:MM format.
	DefEntryTime param.Opt[string] `json:"defEntryTime,omitzero"`
	// Remarks concerning the default exit point for this country.
	DefExitRemark param.Opt[string] `json:"defExitRemark,omitzero"`
	// Zulu default exit time for this country expressed in HH:MM format.
	DefExitTime param.Opt[string] `json:"defExitTime,omitzero"`
	// Remarks concerning flight information for this country profile.
	FltInfoRemark param.Opt[string] `json:"fltInfoRemark,omitzero"`
	// Remarks concerning hazardous material information for this country profile.
	HazInfoRemark param.Opt[string] `json:"hazInfoRemark,omitzero"`
	// Flag indicating whether this is the default landing profile for this country.
	LandDefProf param.Opt[bool] `json:"landDefProf,omitzero"`
	// Amount of lead time required for an aircraft to land in this country. Units need
	// to be specified in the landLeadTimeUnit field.
	LandLeadTime param.Opt[int64] `json:"landLeadTime,omitzero"`
	// Remarks concerning the landing lead time required for this country.
	LandLeadTimeRemark param.Opt[string] `json:"landLeadTimeRemark,omitzero"`
	// Unit of time specified for the landLeadTime field to indicate the landing lead
	// time required for this country.
	LandLeadTimeUnit param.Opt[string] `json:"landLeadTimeUnit,omitzero"`
	// Amount of time before the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodMinus param.Opt[int64] `json:"landValidPeriodMinus,omitzero"`
	// Amount of time after the landing valid period that an aircraft is allowed to
	// land in this country for this profile. The unit of time should be specified in
	// the landValidPeriodUnit field.
	LandValidPeriodPlus param.Opt[int64] `json:"landValidPeriodPlus,omitzero"`
	// Remarks concerning the valid landing time period for this country.
	LandValidPeriodRemark param.Opt[string] `json:"landValidPeriodRemark,omitzero"`
	// Unit of time specified for the landValidPeriodPlus and landValidPeriodMinus
	// fields to indicate the valid landing period for this country.
	LandValidPeriodUnit param.Opt[string] `json:"landValidPeriodUnit,omitzero"`
	// Flag indicating whether this is the default overfly profile for this country.
	OverflyDefProf param.Opt[bool] `json:"overflyDefProf,omitzero"`
	// Amount of lead time required for an aircraft to enter and fly over the airspace
	// of this country. Units need to be specified in the overflyLeadTimeUnit field.
	OverflyLeadTime param.Opt[int64] `json:"overflyLeadTime,omitzero"`
	// Remarks concerning the overfly lead time required for this country.
	OverflyLeadTimeRemark param.Opt[string] `json:"overflyLeadTimeRemark,omitzero"`
	// Unit of time specified for the overflyLeadTime field to indicate the overfly
	// lead time required for this country.
	OverflyLeadTimeUnit param.Opt[string] `json:"overflyLeadTimeUnit,omitzero"`
	// Amount of time before the overfly valid period that an aircraft is allowed to
	// fly over this country for this profile. The unit of time should be specified in
	// the overflyValidPeriodUnit field.
	OverflyValidPeriodMinus param.Opt[int64] `json:"overflyValidPeriodMinus,omitzero"`
	// Amount of time after the overfly valid period that an aircraft is allowed to fly
	// over this country for this profile. The unit of time should be specified in the
	// overflyValidPeriodUnit field.
	OverflyValidPeriodPlus param.Opt[int64] `json:"overflyValidPeriodPlus,omitzero"`
	// Remarks concerning the valid overfly time period for this country.
	OverflyValidPeriodRemark param.Opt[string] `json:"overflyValidPeriodRemark,omitzero"`
	// Unit of time specified for the overflyValidPeriodPlus and
	// overflyValidPeriodMinus fields to indicate the valid overfly period for this
	// country.
	OverflyValidPeriodUnit param.Opt[string] `json:"overflyValidPeriodUnit,omitzero"`
	// The diplomatic clearance profile name used within clearance management systems.
	Profile param.Opt[string] `json:"profile,omitzero"`
	// The agency to which this profile applies.
	ProfileAgency param.Opt[string] `json:"profileAgency,omitzero"`
	// Identifier of the diplomatic clearance country profile.
	ProfileID param.Opt[string] `json:"profileId,omitzero"`
	// Remarks concerning this country profile.
	ProfileRemark param.Opt[string] `json:"profileRemark,omitzero"`
	// Flag indicating whether alternate aircraft names are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAcAltName param.Opt[bool] `json:"reqACAltName,omitzero"`
	// Flag indicating whether all hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqAllHazInfo param.Opt[bool] `json:"reqAllHazInfo,omitzero"`
	// Flag indicating whether standard AMC information is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqAmcStdInfo param.Opt[bool] `json:"reqAMCStdInfo,omitzero"`
	// Flag indicating whether a cargo list is required to be reported to the country
	// using this diplomatic clearance profile.
	ReqCargoList param.Opt[bool] `json:"reqCargoList,omitzero"`
	// Flag indicating whether aircraft cargo and passenger information is required to
	// be reported to the country using this diplomatic clearance profile.
	ReqCargoPax param.Opt[bool] `json:"reqCargoPax,omitzero"`
	// Flag indicating whether Class 1 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass1Info param.Opt[bool] `json:"reqClass1Info,omitzero"`
	// Flag indicating whether Class 9 hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqClass9Info param.Opt[bool] `json:"reqClass9Info,omitzero"`
	// Flag indicating whether the number of crew members on a flight is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqCrewComp param.Opt[bool] `json:"reqCrewComp,omitzero"`
	// Flag indicating whether the names, ranks, and positions of crew members are
	// required to be reported to the country using this diplomatic clearance profile.
	ReqCrewDetail param.Opt[bool] `json:"reqCrewDetail,omitzero"`
	// Flag indicating whether crew information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqCrewInfo param.Opt[bool] `json:"reqCrewInfo,omitzero"`
	// Flag indicating whether Division 1.1 hazardous material information is required
	// to be reported to the country using this diplomatic clearance profile.
	ReqDiv1Info param.Opt[bool] `json:"reqDiv1Info,omitzero"`
	// Flag indicating whether distinguished visitors are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqDv param.Opt[bool] `json:"reqDV,omitzero"`
	// Flag indicating whether entry/exit coordinates need to be specified for this
	// diplomatic clearance profile.
	ReqEntryExitCoord param.Opt[bool] `json:"reqEntryExitCoord,omitzero"`
	// Flag indicating whether flight information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltInfo param.Opt[bool] `json:"reqFltInfo,omitzero"`
	// Flag indicating whether a flight plan route is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqFltPlanRoute param.Opt[bool] `json:"reqFltPlanRoute,omitzero"`
	// Flag indicating whether aviation funding sources are required to be reported to
	// the country using this diplomatic clearance profile.
	ReqFundSource param.Opt[bool] `json:"reqFundSource,omitzero"`
	// Flag indicating whether hazardous material information is required to be
	// reported to the country using this diplomatic clearance profile.
	ReqHazInfo param.Opt[bool] `json:"reqHazInfo,omitzero"`
	// Flag indicating whether this diplomatic clearance profile applies to specific
	// ICAO(s). These specific ICAO(s) should be clarified in the fltInfoRemark field.
	ReqIcao param.Opt[bool] `json:"reqICAO,omitzero"`
	// Flag indicating whether passport information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqPassportInfo param.Opt[bool] `json:"reqPassportInfo,omitzero"`
	// Flag indicating whether ravens are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRaven param.Opt[bool] `json:"reqRaven,omitzero"`
	// Flag indicating whether changes are required to be reported to the country using
	// this diplomatic clearance profile.
	ReqRepChange param.Opt[bool] `json:"reqRepChange,omitzero"`
	// Flag indicating whether an aircraft tail number is required to be reported to
	// the country using this diplomatic clearance profile.
	ReqTailNum param.Opt[bool] `json:"reqTailNum,omitzero"`
	// Flag indicating whether weapons information is required to be reported to the
	// country using this diplomatic clearance profile.
	ReqWeaponsInfo param.Opt[bool] `json:"reqWeaponsInfo,omitzero"`
	// Flag indicating whether crew reporting is undefined for the country using this
	// diplomatic clearance profile.
	UndefinedCrewReporting param.Opt[bool] `json:"undefinedCrewReporting,omitzero"`
	paramObj
}

func (r DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryProfile) MarshalJSON() (data []byte, err error) {
	type shadow DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
