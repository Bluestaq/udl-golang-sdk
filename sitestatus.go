// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
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

// SiteStatusService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteStatusService] method instead.
type SiteStatusService struct {
	Options []option.RequestOption
	History SiteStatusHistoryService
}

// NewSiteStatusService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSiteStatusService(opts ...option.RequestOption) (r SiteStatusService) {
	r = SiteStatusService{}
	r.Options = opts
	r.History = NewSiteStatusHistoryService(opts...)
	return
}

// Service operation to take a single SiteStatus object as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SiteStatusService) New(ctx context.Context, body SiteStatusNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sitestatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single SiteStatus object. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SiteStatusService) Update(ctx context.Context, id string, body SiteStatusUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sitestatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SiteStatusService) List(ctx context.Context, query SiteStatusListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SiteStatusListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sitestatus"
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
func (r *SiteStatusService) ListAutoPaging(ctx context.Context, query SiteStatusListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SiteStatusListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SiteStatus object specified by the passed ID path
// parameter. Note, delete operations do not remove data from historical or
// publish/subscribe stores. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SiteStatusService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sitestatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SiteStatusService) Count(ctx context.Context, query SiteStatusCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sitestatus/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SiteStatus record by its unique ID passed as a
// path parameter.
func (r *SiteStatusService) Get(ctx context.Context, id string, query SiteStatusGetParams, opts ...option.RequestOption) (res *SiteStatusGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sitestatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SiteStatusService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SiteStatusQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/sitestatus/queryhelp"
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
func (r *SiteStatusService) Tuple(ctx context.Context, query SiteStatusTupleParams, opts ...option.RequestOption) (res *[]SiteStatusTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/sitestatus/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SiteStatusListResponse struct {
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
	DataMode SiteStatusListResponseDataMode `json:"dataMode,required"`
	// The ID of the site, if this status is associated with a fixed site or platform.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Crisis Action Team (CAT).
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Cat SiteStatusListResponseCat `json:"cat"`
	// Estimated number of cold missiles of all types remaining in weapons system
	// inventory.
	ColdInventory int64 `json:"coldInventory"`
	// The communications component causing the platform or system to be less than
	// fully operational.
	CommImpairment string `json:"commImpairment"`
	// Cyberspace Protection Condition (CPCON).
	//
	// 1 - VERY HIGH - Critical functions.
	//
	// 2 - HIGH - Critical and essential functions.
	//
	// 3 - MEDIUM - Critical, essential, and support functions.
	//
	// 4 - LOW - All functions.
	//
	// 5 - VERY LOW - All functions.
	//
	// Any of "1", "2", "3", "4", "5".
	Cpcon SiteStatusListResponseCpcon `json:"cpcon"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Emergency Operations Center (EOC) status.
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Eoc SiteStatusListResponseEoc `json:"eoc"`
	// Force Protection Condition (FPCON).
	//
	// NORMAL - Applies when a general global threat of possible terrorist activity
	// exists and warrants a routine security posture.
	//
	// ALPHA - Applies when an increased general threat of possible terrorist activity
	// against personnel or facilities. Nature and extent of threat are unpredictable.
	//
	// BRAVO - Applies when an increased or predictable threat of terrorist activity
	// exists.
	//
	// CHARLIE - Applies when an incident occurs or intelligence is received indicating
	// some form of terrorist action against personnel and facilities is imminent.
	//
	// DELTA - Applies in the immediate area where an attack has occurred or when
	// intelligence is received indicating terrorist action against a location is
	// imminent.
	//
	// Any of "NORMAL", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Fpcon SiteStatusListResponseFpcon `json:"fpcon"`
	// Estimated number of hot missiles of all types remaining in weapons system
	// inventory.
	HotInventory int64 `json:"hotInventory"`
	// Health Protection Condition (HPCON).
	//
	// 0 - Routine, no community transmission.
	//
	// ALPHA - Limited, community transmission beginning.
	//
	// BRAVO - Moderate, increased community transmission.
	//
	// CHARLIE - Substantial, sustained community transmission.
	//
	// DELTA - Severe, widespread community transmission.
	//
	// Any of "0", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Hpcon SiteStatusListResponseHpcon `json:"hpcon"`
	// The status of the installation.
	//
	// # FMC - Fully Mission Capable
	//
	// # PMC - Partially Mission Capable
	//
	// # NMC - Non Mission Capable
	//
	// UNK - Unknown.
	//
	// Any of "FMC", "PMC", "NMC", "UNK".
	InstStatus SiteStatusListResponseInstStatus `json:"instStatus"`
	// Array of Link item(s) for which status is available and reported (ATDL, IJMS,
	// LINK-1, LINK-11, LINK-11B, LINK-16). This array must be the same length as the
	// linkStatus array.
	Link []string `json:"link"`
	// Array of the status (AVAILABLE, DEGRADED, NOT AVAILABLE, etc.) for each links in
	// the link array. This array must be the same length as the link array, and the
	// status must correspond to the appropriate position index in the link array.
	LinkStatus []string `json:"linkStatus"`
	// Array of specific missile types for which an estimated inventory count is
	// available (e.g. GMD TYPE A, HARPOON, TOMAHAWK, etc.). This array must be the
	// same length as the missileInventory array.
	Missile []string `json:"missile"`
	// Array of the quantity of each of the missile items. This array must be the same
	// length as the missile array, and the values must correspond to appropriate
	// position index in the missile array.
	MissileInventory []int64 `json:"missileInventory"`
	// Alternate Identifier for a mobile or transportable platform provided by source.
	MobileAltID string `json:"mobileAltId"`
	// The operational status of the platform (e.g. Fully Operational, Partially
	// Operational, Not Operational, etc.).
	OpsCapability string `json:"opsCapability"`
	// The primary component degrading the operational capability of the platform or
	// system.
	OpsImpairment string `json:"opsImpairment"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Position Engagement Status flag, Indicating whether this platform is initiating
	// multiple simultaneous engagements. A value of 1/True indicates the platform is
	// initiating multiple simultaneous engagements.
	Pes bool `json:"pes"`
	// The POI (point of interest) ID related to this platform, if available.
	Poiid string `json:"poiid"`
	// Array of the status (NON-OPERATIONAL, OPERATIONAL, OFF) for each radar system in
	// the radarSystem array. This array must be the same length as the radarSystem
	// array, and the status must correspond to the appropriate position index in the
	// radarSystem array.
	RadarStatus []string `json:"radarStatus"`
	// Array of radar system(s) for which status is available and reported
	// (ACQUISITION, IFFSIF, ILLUMINATING, MODE-4, PRIMARY SURVEILLANCE, SECONDARY
	// SURVEILLANCE, TERTIARY SURVEILLANCE). This array must be the same length as the
	// radarStatus array.
	RadarSystem []string `json:"radarSystem"`
	// SAM sensor radar surveillance mode (Active, Passive, Off).
	RadiateMode string `json:"radiateMode"`
	// Time of report, in ISO8601 UTC format.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// The state of a SAM unit (e.g. Initialization, Standby, Reorientation, etc.).
	SamMode string `json:"samMode"`
	// Optional site type or further detail of type. Intended for, but not limited to,
	// Link-16 site type specifications (e.g. ADOC, GACC, SOC, TACC, etc.).
	SiteType string `json:"siteType"`
	// Description of the time function associated with the reportTime (e.g.
	// Activation, Deactivation, Arrival, Departure, etc.), if applicable.
	TimeFunction string `json:"timeFunction"`
	// The track ID related to this platform (if mobile or transportable), if
	// available.
	TrackID string `json:"trackId"`
	// Link-16 specific reference track number.
	TrackRefL16 string `json:"trackRefL16"`
	// Description of the current weather conditions over a site.
	WeatherMessage string `json:"weatherMessage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSite                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Cat                   respjson.Field
		ColdInventory         respjson.Field
		CommImpairment        respjson.Field
		Cpcon                 respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Eoc                   respjson.Field
		Fpcon                 respjson.Field
		HotInventory          respjson.Field
		Hpcon                 respjson.Field
		InstStatus            respjson.Field
		Link                  respjson.Field
		LinkStatus            respjson.Field
		Missile               respjson.Field
		MissileInventory      respjson.Field
		MobileAltID           respjson.Field
		OpsCapability         respjson.Field
		OpsImpairment         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Pes                   respjson.Field
		Poiid                 respjson.Field
		RadarStatus           respjson.Field
		RadarSystem           respjson.Field
		RadiateMode           respjson.Field
		ReportTime            respjson.Field
		SamMode               respjson.Field
		SiteType              respjson.Field
		TimeFunction          respjson.Field
		TrackID               respjson.Field
		TrackRefL16           respjson.Field
		WeatherMessage        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteStatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteStatusListResponse) UnmarshalJSON(data []byte) error {
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
type SiteStatusListResponseDataMode string

const (
	SiteStatusListResponseDataModeReal      SiteStatusListResponseDataMode = "REAL"
	SiteStatusListResponseDataModeTest      SiteStatusListResponseDataMode = "TEST"
	SiteStatusListResponseDataModeSimulated SiteStatusListResponseDataMode = "SIMULATED"
	SiteStatusListResponseDataModeExercise  SiteStatusListResponseDataMode = "EXERCISE"
)

// Crisis Action Team (CAT).
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusListResponseCat string

const (
	SiteStatusListResponseCatCold SiteStatusListResponseCat = "COLD"
	SiteStatusListResponseCatWarm SiteStatusListResponseCat = "WARM"
	SiteStatusListResponseCatHot  SiteStatusListResponseCat = "HOT"
)

// Cyberspace Protection Condition (CPCON).
//
// 1 - VERY HIGH - Critical functions.
//
// 2 - HIGH - Critical and essential functions.
//
// 3 - MEDIUM - Critical, essential, and support functions.
//
// 4 - LOW - All functions.
//
// 5 - VERY LOW - All functions.
type SiteStatusListResponseCpcon string

const (
	SiteStatusListResponseCpcon1 SiteStatusListResponseCpcon = "1"
	SiteStatusListResponseCpcon2 SiteStatusListResponseCpcon = "2"
	SiteStatusListResponseCpcon3 SiteStatusListResponseCpcon = "3"
	SiteStatusListResponseCpcon4 SiteStatusListResponseCpcon = "4"
	SiteStatusListResponseCpcon5 SiteStatusListResponseCpcon = "5"
)

// Emergency Operations Center (EOC) status.
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusListResponseEoc string

const (
	SiteStatusListResponseEocCold SiteStatusListResponseEoc = "COLD"
	SiteStatusListResponseEocWarm SiteStatusListResponseEoc = "WARM"
	SiteStatusListResponseEocHot  SiteStatusListResponseEoc = "HOT"
)

// Force Protection Condition (FPCON).
//
// NORMAL - Applies when a general global threat of possible terrorist activity
// exists and warrants a routine security posture.
//
// ALPHA - Applies when an increased general threat of possible terrorist activity
// against personnel or facilities. Nature and extent of threat are unpredictable.
//
// BRAVO - Applies when an increased or predictable threat of terrorist activity
// exists.
//
// CHARLIE - Applies when an incident occurs or intelligence is received indicating
// some form of terrorist action against personnel and facilities is imminent.
//
// DELTA - Applies in the immediate area where an attack has occurred or when
// intelligence is received indicating terrorist action against a location is
// imminent.
type SiteStatusListResponseFpcon string

const (
	SiteStatusListResponseFpconNormal  SiteStatusListResponseFpcon = "NORMAL"
	SiteStatusListResponseFpconAlpha   SiteStatusListResponseFpcon = "ALPHA"
	SiteStatusListResponseFpconBravo   SiteStatusListResponseFpcon = "BRAVO"
	SiteStatusListResponseFpconCharlie SiteStatusListResponseFpcon = "CHARLIE"
	SiteStatusListResponseFpconDelta   SiteStatusListResponseFpcon = "DELTA"
)

// Health Protection Condition (HPCON).
//
// 0 - Routine, no community transmission.
//
// ALPHA - Limited, community transmission beginning.
//
// BRAVO - Moderate, increased community transmission.
//
// CHARLIE - Substantial, sustained community transmission.
//
// DELTA - Severe, widespread community transmission.
type SiteStatusListResponseHpcon string

const (
	SiteStatusListResponseHpcon0       SiteStatusListResponseHpcon = "0"
	SiteStatusListResponseHpconAlpha   SiteStatusListResponseHpcon = "ALPHA"
	SiteStatusListResponseHpconBravo   SiteStatusListResponseHpcon = "BRAVO"
	SiteStatusListResponseHpconCharlie SiteStatusListResponseHpcon = "CHARLIE"
	SiteStatusListResponseHpconDelta   SiteStatusListResponseHpcon = "DELTA"
)

// The status of the installation.
//
// # FMC - Fully Mission Capable
//
// # PMC - Partially Mission Capable
//
// # NMC - Non Mission Capable
//
// UNK - Unknown.
type SiteStatusListResponseInstStatus string

const (
	SiteStatusListResponseInstStatusFmc SiteStatusListResponseInstStatus = "FMC"
	SiteStatusListResponseInstStatusPmc SiteStatusListResponseInstStatus = "PMC"
	SiteStatusListResponseInstStatusNmc SiteStatusListResponseInstStatus = "NMC"
	SiteStatusListResponseInstStatusUnk SiteStatusListResponseInstStatus = "UNK"
)

type SiteStatusGetResponse struct {
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
	DataMode SiteStatusGetResponseDataMode `json:"dataMode,required"`
	// The ID of the site, if this status is associated with a fixed site or platform.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Crisis Action Team (CAT).
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Cat SiteStatusGetResponseCat `json:"cat"`
	// Estimated number of cold missiles of all types remaining in weapons system
	// inventory.
	ColdInventory int64 `json:"coldInventory"`
	// The communications component causing the platform or system to be less than
	// fully operational.
	CommImpairment string `json:"commImpairment"`
	// Cyberspace Protection Condition (CPCON).
	//
	// 1 - VERY HIGH - Critical functions.
	//
	// 2 - HIGH - Critical and essential functions.
	//
	// 3 - MEDIUM - Critical, essential, and support functions.
	//
	// 4 - LOW - All functions.
	//
	// 5 - VERY LOW - All functions.
	//
	// Any of "1", "2", "3", "4", "5".
	Cpcon SiteStatusGetResponseCpcon `json:"cpcon"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Emergency Operations Center (EOC) status.
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Eoc SiteStatusGetResponseEoc `json:"eoc"`
	// Force Protection Condition (FPCON).
	//
	// NORMAL - Applies when a general global threat of possible terrorist activity
	// exists and warrants a routine security posture.
	//
	// ALPHA - Applies when an increased general threat of possible terrorist activity
	// against personnel or facilities. Nature and extent of threat are unpredictable.
	//
	// BRAVO - Applies when an increased or predictable threat of terrorist activity
	// exists.
	//
	// CHARLIE - Applies when an incident occurs or intelligence is received indicating
	// some form of terrorist action against personnel and facilities is imminent.
	//
	// DELTA - Applies in the immediate area where an attack has occurred or when
	// intelligence is received indicating terrorist action against a location is
	// imminent.
	//
	// Any of "NORMAL", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Fpcon SiteStatusGetResponseFpcon `json:"fpcon"`
	// Estimated number of hot missiles of all types remaining in weapons system
	// inventory.
	HotInventory int64 `json:"hotInventory"`
	// Health Protection Condition (HPCON).
	//
	// 0 - Routine, no community transmission.
	//
	// ALPHA - Limited, community transmission beginning.
	//
	// BRAVO - Moderate, increased community transmission.
	//
	// CHARLIE - Substantial, sustained community transmission.
	//
	// DELTA - Severe, widespread community transmission.
	//
	// Any of "0", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Hpcon SiteStatusGetResponseHpcon `json:"hpcon"`
	// The status of the installation.
	//
	// # FMC - Fully Mission Capable
	//
	// # PMC - Partially Mission Capable
	//
	// # NMC - Non Mission Capable
	//
	// UNK - Unknown.
	//
	// Any of "FMC", "PMC", "NMC", "UNK".
	InstStatus SiteStatusGetResponseInstStatus `json:"instStatus"`
	// Array of Link item(s) for which status is available and reported (ATDL, IJMS,
	// LINK-1, LINK-11, LINK-11B, LINK-16). This array must be the same length as the
	// linkStatus array.
	Link []string `json:"link"`
	// Array of the status (AVAILABLE, DEGRADED, NOT AVAILABLE, etc.) for each links in
	// the link array. This array must be the same length as the link array, and the
	// status must correspond to the appropriate position index in the link array.
	LinkStatus []string `json:"linkStatus"`
	// Array of specific missile types for which an estimated inventory count is
	// available (e.g. GMD TYPE A, HARPOON, TOMAHAWK, etc.). This array must be the
	// same length as the missileInventory array.
	Missile []string `json:"missile"`
	// Array of the quantity of each of the missile items. This array must be the same
	// length as the missile array, and the values must correspond to appropriate
	// position index in the missile array.
	MissileInventory []int64 `json:"missileInventory"`
	// Alternate Identifier for a mobile or transportable platform provided by source.
	MobileAltID string `json:"mobileAltId"`
	// The operational status of the platform (e.g. Fully Operational, Partially
	// Operational, Not Operational, etc.).
	OpsCapability string `json:"opsCapability"`
	// The primary component degrading the operational capability of the platform or
	// system.
	OpsImpairment string `json:"opsImpairment"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Position Engagement Status flag, Indicating whether this platform is initiating
	// multiple simultaneous engagements. A value of 1/True indicates the platform is
	// initiating multiple simultaneous engagements.
	Pes bool `json:"pes"`
	// The POI (point of interest) ID related to this platform, if available.
	Poiid string `json:"poiid"`
	// Array of the status (NON-OPERATIONAL, OPERATIONAL, OFF) for each radar system in
	// the radarSystem array. This array must be the same length as the radarSystem
	// array, and the status must correspond to the appropriate position index in the
	// radarSystem array.
	RadarStatus []string `json:"radarStatus"`
	// Array of radar system(s) for which status is available and reported
	// (ACQUISITION, IFFSIF, ILLUMINATING, MODE-4, PRIMARY SURVEILLANCE, SECONDARY
	// SURVEILLANCE, TERTIARY SURVEILLANCE). This array must be the same length as the
	// radarStatus array.
	RadarSystem []string `json:"radarSystem"`
	// SAM sensor radar surveillance mode (Active, Passive, Off).
	RadiateMode string `json:"radiateMode"`
	// Time of report, in ISO8601 UTC format.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// The state of a SAM unit (e.g. Initialization, Standby, Reorientation, etc.).
	SamMode string `json:"samMode"`
	// Optional site type or further detail of type. Intended for, but not limited to,
	// Link-16 site type specifications (e.g. ADOC, GACC, SOC, TACC, etc.).
	SiteType string `json:"siteType"`
	// Description of the time function associated with the reportTime (e.g.
	// Activation, Deactivation, Arrival, Departure, etc.), if applicable.
	TimeFunction string `json:"timeFunction"`
	// The track ID related to this platform (if mobile or transportable), if
	// available.
	TrackID string `json:"trackId"`
	// Link-16 specific reference track number.
	TrackRefL16 string `json:"trackRefL16"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Description of the current weather conditions over a site.
	WeatherMessage string `json:"weatherMessage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSite                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Cat                   respjson.Field
		ColdInventory         respjson.Field
		CommImpairment        respjson.Field
		Cpcon                 respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Eoc                   respjson.Field
		Fpcon                 respjson.Field
		HotInventory          respjson.Field
		Hpcon                 respjson.Field
		InstStatus            respjson.Field
		Link                  respjson.Field
		LinkStatus            respjson.Field
		Missile               respjson.Field
		MissileInventory      respjson.Field
		MobileAltID           respjson.Field
		OpsCapability         respjson.Field
		OpsImpairment         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Pes                   respjson.Field
		Poiid                 respjson.Field
		RadarStatus           respjson.Field
		RadarSystem           respjson.Field
		RadiateMode           respjson.Field
		ReportTime            respjson.Field
		SamMode               respjson.Field
		SiteType              respjson.Field
		TimeFunction          respjson.Field
		TrackID               respjson.Field
		TrackRefL16           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		WeatherMessage        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteStatusGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteStatusGetResponse) UnmarshalJSON(data []byte) error {
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
type SiteStatusGetResponseDataMode string

const (
	SiteStatusGetResponseDataModeReal      SiteStatusGetResponseDataMode = "REAL"
	SiteStatusGetResponseDataModeTest      SiteStatusGetResponseDataMode = "TEST"
	SiteStatusGetResponseDataModeSimulated SiteStatusGetResponseDataMode = "SIMULATED"
	SiteStatusGetResponseDataModeExercise  SiteStatusGetResponseDataMode = "EXERCISE"
)

// Crisis Action Team (CAT).
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusGetResponseCat string

const (
	SiteStatusGetResponseCatCold SiteStatusGetResponseCat = "COLD"
	SiteStatusGetResponseCatWarm SiteStatusGetResponseCat = "WARM"
	SiteStatusGetResponseCatHot  SiteStatusGetResponseCat = "HOT"
)

// Cyberspace Protection Condition (CPCON).
//
// 1 - VERY HIGH - Critical functions.
//
// 2 - HIGH - Critical and essential functions.
//
// 3 - MEDIUM - Critical, essential, and support functions.
//
// 4 - LOW - All functions.
//
// 5 - VERY LOW - All functions.
type SiteStatusGetResponseCpcon string

const (
	SiteStatusGetResponseCpcon1 SiteStatusGetResponseCpcon = "1"
	SiteStatusGetResponseCpcon2 SiteStatusGetResponseCpcon = "2"
	SiteStatusGetResponseCpcon3 SiteStatusGetResponseCpcon = "3"
	SiteStatusGetResponseCpcon4 SiteStatusGetResponseCpcon = "4"
	SiteStatusGetResponseCpcon5 SiteStatusGetResponseCpcon = "5"
)

// Emergency Operations Center (EOC) status.
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusGetResponseEoc string

const (
	SiteStatusGetResponseEocCold SiteStatusGetResponseEoc = "COLD"
	SiteStatusGetResponseEocWarm SiteStatusGetResponseEoc = "WARM"
	SiteStatusGetResponseEocHot  SiteStatusGetResponseEoc = "HOT"
)

// Force Protection Condition (FPCON).
//
// NORMAL - Applies when a general global threat of possible terrorist activity
// exists and warrants a routine security posture.
//
// ALPHA - Applies when an increased general threat of possible terrorist activity
// against personnel or facilities. Nature and extent of threat are unpredictable.
//
// BRAVO - Applies when an increased or predictable threat of terrorist activity
// exists.
//
// CHARLIE - Applies when an incident occurs or intelligence is received indicating
// some form of terrorist action against personnel and facilities is imminent.
//
// DELTA - Applies in the immediate area where an attack has occurred or when
// intelligence is received indicating terrorist action against a location is
// imminent.
type SiteStatusGetResponseFpcon string

const (
	SiteStatusGetResponseFpconNormal  SiteStatusGetResponseFpcon = "NORMAL"
	SiteStatusGetResponseFpconAlpha   SiteStatusGetResponseFpcon = "ALPHA"
	SiteStatusGetResponseFpconBravo   SiteStatusGetResponseFpcon = "BRAVO"
	SiteStatusGetResponseFpconCharlie SiteStatusGetResponseFpcon = "CHARLIE"
	SiteStatusGetResponseFpconDelta   SiteStatusGetResponseFpcon = "DELTA"
)

// Health Protection Condition (HPCON).
//
// 0 - Routine, no community transmission.
//
// ALPHA - Limited, community transmission beginning.
//
// BRAVO - Moderate, increased community transmission.
//
// CHARLIE - Substantial, sustained community transmission.
//
// DELTA - Severe, widespread community transmission.
type SiteStatusGetResponseHpcon string

const (
	SiteStatusGetResponseHpcon0       SiteStatusGetResponseHpcon = "0"
	SiteStatusGetResponseHpconAlpha   SiteStatusGetResponseHpcon = "ALPHA"
	SiteStatusGetResponseHpconBravo   SiteStatusGetResponseHpcon = "BRAVO"
	SiteStatusGetResponseHpconCharlie SiteStatusGetResponseHpcon = "CHARLIE"
	SiteStatusGetResponseHpconDelta   SiteStatusGetResponseHpcon = "DELTA"
)

// The status of the installation.
//
// # FMC - Fully Mission Capable
//
// # PMC - Partially Mission Capable
//
// # NMC - Non Mission Capable
//
// UNK - Unknown.
type SiteStatusGetResponseInstStatus string

const (
	SiteStatusGetResponseInstStatusFmc SiteStatusGetResponseInstStatus = "FMC"
	SiteStatusGetResponseInstStatusPmc SiteStatusGetResponseInstStatus = "PMC"
	SiteStatusGetResponseInstStatusNmc SiteStatusGetResponseInstStatus = "NMC"
	SiteStatusGetResponseInstStatusUnk SiteStatusGetResponseInstStatus = "UNK"
)

type SiteStatusQueryhelpResponse struct {
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
func (r SiteStatusQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteStatusQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteStatusTupleResponse struct {
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
	DataMode SiteStatusTupleResponseDataMode `json:"dataMode,required"`
	// The ID of the site, if this status is associated with a fixed site or platform.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Crisis Action Team (CAT).
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Cat SiteStatusTupleResponseCat `json:"cat"`
	// Estimated number of cold missiles of all types remaining in weapons system
	// inventory.
	ColdInventory int64 `json:"coldInventory"`
	// The communications component causing the platform or system to be less than
	// fully operational.
	CommImpairment string `json:"commImpairment"`
	// Cyberspace Protection Condition (CPCON).
	//
	// 1 - VERY HIGH - Critical functions.
	//
	// 2 - HIGH - Critical and essential functions.
	//
	// 3 - MEDIUM - Critical, essential, and support functions.
	//
	// 4 - LOW - All functions.
	//
	// 5 - VERY LOW - All functions.
	//
	// Any of "1", "2", "3", "4", "5".
	Cpcon SiteStatusTupleResponseCpcon `json:"cpcon"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Emergency Operations Center (EOC) status.
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Eoc SiteStatusTupleResponseEoc `json:"eoc"`
	// Force Protection Condition (FPCON).
	//
	// NORMAL - Applies when a general global threat of possible terrorist activity
	// exists and warrants a routine security posture.
	//
	// ALPHA - Applies when an increased general threat of possible terrorist activity
	// against personnel or facilities. Nature and extent of threat are unpredictable.
	//
	// BRAVO - Applies when an increased or predictable threat of terrorist activity
	// exists.
	//
	// CHARLIE - Applies when an incident occurs or intelligence is received indicating
	// some form of terrorist action against personnel and facilities is imminent.
	//
	// DELTA - Applies in the immediate area where an attack has occurred or when
	// intelligence is received indicating terrorist action against a location is
	// imminent.
	//
	// Any of "NORMAL", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Fpcon SiteStatusTupleResponseFpcon `json:"fpcon"`
	// Estimated number of hot missiles of all types remaining in weapons system
	// inventory.
	HotInventory int64 `json:"hotInventory"`
	// Health Protection Condition (HPCON).
	//
	// 0 - Routine, no community transmission.
	//
	// ALPHA - Limited, community transmission beginning.
	//
	// BRAVO - Moderate, increased community transmission.
	//
	// CHARLIE - Substantial, sustained community transmission.
	//
	// DELTA - Severe, widespread community transmission.
	//
	// Any of "0", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Hpcon SiteStatusTupleResponseHpcon `json:"hpcon"`
	// The status of the installation.
	//
	// # FMC - Fully Mission Capable
	//
	// # PMC - Partially Mission Capable
	//
	// # NMC - Non Mission Capable
	//
	// UNK - Unknown.
	//
	// Any of "FMC", "PMC", "NMC", "UNK".
	InstStatus SiteStatusTupleResponseInstStatus `json:"instStatus"`
	// Array of Link item(s) for which status is available and reported (ATDL, IJMS,
	// LINK-1, LINK-11, LINK-11B, LINK-16). This array must be the same length as the
	// linkStatus array.
	Link []string `json:"link"`
	// Array of the status (AVAILABLE, DEGRADED, NOT AVAILABLE, etc.) for each links in
	// the link array. This array must be the same length as the link array, and the
	// status must correspond to the appropriate position index in the link array.
	LinkStatus []string `json:"linkStatus"`
	// Array of specific missile types for which an estimated inventory count is
	// available (e.g. GMD TYPE A, HARPOON, TOMAHAWK, etc.). This array must be the
	// same length as the missileInventory array.
	Missile []string `json:"missile"`
	// Array of the quantity of each of the missile items. This array must be the same
	// length as the missile array, and the values must correspond to appropriate
	// position index in the missile array.
	MissileInventory []int64 `json:"missileInventory"`
	// Alternate Identifier for a mobile or transportable platform provided by source.
	MobileAltID string `json:"mobileAltId"`
	// The operational status of the platform (e.g. Fully Operational, Partially
	// Operational, Not Operational, etc.).
	OpsCapability string `json:"opsCapability"`
	// The primary component degrading the operational capability of the platform or
	// system.
	OpsImpairment string `json:"opsImpairment"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Position Engagement Status flag, Indicating whether this platform is initiating
	// multiple simultaneous engagements. A value of 1/True indicates the platform is
	// initiating multiple simultaneous engagements.
	Pes bool `json:"pes"`
	// The POI (point of interest) ID related to this platform, if available.
	Poiid string `json:"poiid"`
	// Array of the status (NON-OPERATIONAL, OPERATIONAL, OFF) for each radar system in
	// the radarSystem array. This array must be the same length as the radarSystem
	// array, and the status must correspond to the appropriate position index in the
	// radarSystem array.
	RadarStatus []string `json:"radarStatus"`
	// Array of radar system(s) for which status is available and reported
	// (ACQUISITION, IFFSIF, ILLUMINATING, MODE-4, PRIMARY SURVEILLANCE, SECONDARY
	// SURVEILLANCE, TERTIARY SURVEILLANCE). This array must be the same length as the
	// radarStatus array.
	RadarSystem []string `json:"radarSystem"`
	// SAM sensor radar surveillance mode (Active, Passive, Off).
	RadiateMode string `json:"radiateMode"`
	// Time of report, in ISO8601 UTC format.
	ReportTime time.Time `json:"reportTime" format:"date-time"`
	// The state of a SAM unit (e.g. Initialization, Standby, Reorientation, etc.).
	SamMode string `json:"samMode"`
	// Optional site type or further detail of type. Intended for, but not limited to,
	// Link-16 site type specifications (e.g. ADOC, GACC, SOC, TACC, etc.).
	SiteType string `json:"siteType"`
	// Description of the time function associated with the reportTime (e.g.
	// Activation, Deactivation, Arrival, Departure, etc.), if applicable.
	TimeFunction string `json:"timeFunction"`
	// The track ID related to this platform (if mobile or transportable), if
	// available.
	TrackID string `json:"trackId"`
	// Link-16 specific reference track number.
	TrackRefL16 string `json:"trackRefL16"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Description of the current weather conditions over a site.
	WeatherMessage string `json:"weatherMessage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDSite                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Cat                   respjson.Field
		ColdInventory         respjson.Field
		CommImpairment        respjson.Field
		Cpcon                 respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Eoc                   respjson.Field
		Fpcon                 respjson.Field
		HotInventory          respjson.Field
		Hpcon                 respjson.Field
		InstStatus            respjson.Field
		Link                  respjson.Field
		LinkStatus            respjson.Field
		Missile               respjson.Field
		MissileInventory      respjson.Field
		MobileAltID           respjson.Field
		OpsCapability         respjson.Field
		OpsImpairment         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Pes                   respjson.Field
		Poiid                 respjson.Field
		RadarStatus           respjson.Field
		RadarSystem           respjson.Field
		RadiateMode           respjson.Field
		ReportTime            respjson.Field
		SamMode               respjson.Field
		SiteType              respjson.Field
		TimeFunction          respjson.Field
		TrackID               respjson.Field
		TrackRefL16           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		WeatherMessage        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiteStatusTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteStatusTupleResponse) UnmarshalJSON(data []byte) error {
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
type SiteStatusTupleResponseDataMode string

const (
	SiteStatusTupleResponseDataModeReal      SiteStatusTupleResponseDataMode = "REAL"
	SiteStatusTupleResponseDataModeTest      SiteStatusTupleResponseDataMode = "TEST"
	SiteStatusTupleResponseDataModeSimulated SiteStatusTupleResponseDataMode = "SIMULATED"
	SiteStatusTupleResponseDataModeExercise  SiteStatusTupleResponseDataMode = "EXERCISE"
)

// Crisis Action Team (CAT).
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusTupleResponseCat string

const (
	SiteStatusTupleResponseCatCold SiteStatusTupleResponseCat = "COLD"
	SiteStatusTupleResponseCatWarm SiteStatusTupleResponseCat = "WARM"
	SiteStatusTupleResponseCatHot  SiteStatusTupleResponseCat = "HOT"
)

// Cyberspace Protection Condition (CPCON).
//
// 1 - VERY HIGH - Critical functions.
//
// 2 - HIGH - Critical and essential functions.
//
// 3 - MEDIUM - Critical, essential, and support functions.
//
// 4 - LOW - All functions.
//
// 5 - VERY LOW - All functions.
type SiteStatusTupleResponseCpcon string

const (
	SiteStatusTupleResponseCpcon1 SiteStatusTupleResponseCpcon = "1"
	SiteStatusTupleResponseCpcon2 SiteStatusTupleResponseCpcon = "2"
	SiteStatusTupleResponseCpcon3 SiteStatusTupleResponseCpcon = "3"
	SiteStatusTupleResponseCpcon4 SiteStatusTupleResponseCpcon = "4"
	SiteStatusTupleResponseCpcon5 SiteStatusTupleResponseCpcon = "5"
)

// Emergency Operations Center (EOC) status.
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusTupleResponseEoc string

const (
	SiteStatusTupleResponseEocCold SiteStatusTupleResponseEoc = "COLD"
	SiteStatusTupleResponseEocWarm SiteStatusTupleResponseEoc = "WARM"
	SiteStatusTupleResponseEocHot  SiteStatusTupleResponseEoc = "HOT"
)

// Force Protection Condition (FPCON).
//
// NORMAL - Applies when a general global threat of possible terrorist activity
// exists and warrants a routine security posture.
//
// ALPHA - Applies when an increased general threat of possible terrorist activity
// against personnel or facilities. Nature and extent of threat are unpredictable.
//
// BRAVO - Applies when an increased or predictable threat of terrorist activity
// exists.
//
// CHARLIE - Applies when an incident occurs or intelligence is received indicating
// some form of terrorist action against personnel and facilities is imminent.
//
// DELTA - Applies in the immediate area where an attack has occurred or when
// intelligence is received indicating terrorist action against a location is
// imminent.
type SiteStatusTupleResponseFpcon string

const (
	SiteStatusTupleResponseFpconNormal  SiteStatusTupleResponseFpcon = "NORMAL"
	SiteStatusTupleResponseFpconAlpha   SiteStatusTupleResponseFpcon = "ALPHA"
	SiteStatusTupleResponseFpconBravo   SiteStatusTupleResponseFpcon = "BRAVO"
	SiteStatusTupleResponseFpconCharlie SiteStatusTupleResponseFpcon = "CHARLIE"
	SiteStatusTupleResponseFpconDelta   SiteStatusTupleResponseFpcon = "DELTA"
)

// Health Protection Condition (HPCON).
//
// 0 - Routine, no community transmission.
//
// ALPHA - Limited, community transmission beginning.
//
// BRAVO - Moderate, increased community transmission.
//
// CHARLIE - Substantial, sustained community transmission.
//
// DELTA - Severe, widespread community transmission.
type SiteStatusTupleResponseHpcon string

const (
	SiteStatusTupleResponseHpcon0       SiteStatusTupleResponseHpcon = "0"
	SiteStatusTupleResponseHpconAlpha   SiteStatusTupleResponseHpcon = "ALPHA"
	SiteStatusTupleResponseHpconBravo   SiteStatusTupleResponseHpcon = "BRAVO"
	SiteStatusTupleResponseHpconCharlie SiteStatusTupleResponseHpcon = "CHARLIE"
	SiteStatusTupleResponseHpconDelta   SiteStatusTupleResponseHpcon = "DELTA"
)

// The status of the installation.
//
// # FMC - Fully Mission Capable
//
// # PMC - Partially Mission Capable
//
// # NMC - Non Mission Capable
//
// UNK - Unknown.
type SiteStatusTupleResponseInstStatus string

const (
	SiteStatusTupleResponseInstStatusFmc SiteStatusTupleResponseInstStatus = "FMC"
	SiteStatusTupleResponseInstStatusPmc SiteStatusTupleResponseInstStatus = "PMC"
	SiteStatusTupleResponseInstStatusNmc SiteStatusTupleResponseInstStatus = "NMC"
	SiteStatusTupleResponseInstStatusUnk SiteStatusTupleResponseInstStatus = "UNK"
)

type SiteStatusNewParams struct {
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
	DataMode SiteStatusNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The ID of the site, if this status is associated with a fixed site or platform.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Estimated number of cold missiles of all types remaining in weapons system
	// inventory.
	ColdInventory param.Opt[int64] `json:"coldInventory,omitzero"`
	// The communications component causing the platform or system to be less than
	// fully operational.
	CommImpairment param.Opt[string] `json:"commImpairment,omitzero"`
	// Estimated number of hot missiles of all types remaining in weapons system
	// inventory.
	HotInventory param.Opt[int64] `json:"hotInventory,omitzero"`
	// Alternate Identifier for a mobile or transportable platform provided by source.
	MobileAltID param.Opt[string] `json:"mobileAltId,omitzero"`
	// The operational status of the platform (e.g. Fully Operational, Partially
	// Operational, Not Operational, etc.).
	OpsCapability param.Opt[string] `json:"opsCapability,omitzero"`
	// The primary component degrading the operational capability of the platform or
	// system.
	OpsImpairment param.Opt[string] `json:"opsImpairment,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Position Engagement Status flag, Indicating whether this platform is initiating
	// multiple simultaneous engagements. A value of 1/True indicates the platform is
	// initiating multiple simultaneous engagements.
	Pes param.Opt[bool] `json:"pes,omitzero"`
	// The POI (point of interest) ID related to this platform, if available.
	Poiid param.Opt[string] `json:"poiid,omitzero"`
	// SAM sensor radar surveillance mode (Active, Passive, Off).
	RadiateMode param.Opt[string] `json:"radiateMode,omitzero"`
	// Time of report, in ISO8601 UTC format.
	ReportTime param.Opt[time.Time] `json:"reportTime,omitzero" format:"date-time"`
	// The state of a SAM unit (e.g. Initialization, Standby, Reorientation, etc.).
	SamMode param.Opt[string] `json:"samMode,omitzero"`
	// Optional site type or further detail of type. Intended for, but not limited to,
	// Link-16 site type specifications (e.g. ADOC, GACC, SOC, TACC, etc.).
	SiteType param.Opt[string] `json:"siteType,omitzero"`
	// Description of the time function associated with the reportTime (e.g.
	// Activation, Deactivation, Arrival, Departure, etc.), if applicable.
	TimeFunction param.Opt[string] `json:"timeFunction,omitzero"`
	// The track ID related to this platform (if mobile or transportable), if
	// available.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Link-16 specific reference track number.
	TrackRefL16 param.Opt[string] `json:"trackRefL16,omitzero"`
	// Description of the current weather conditions over a site.
	WeatherMessage param.Opt[string] `json:"weatherMessage,omitzero"`
	// Crisis Action Team (CAT).
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Cat SiteStatusNewParamsCat `json:"cat,omitzero"`
	// Cyberspace Protection Condition (CPCON).
	//
	// 1 - VERY HIGH - Critical functions.
	//
	// 2 - HIGH - Critical and essential functions.
	//
	// 3 - MEDIUM - Critical, essential, and support functions.
	//
	// 4 - LOW - All functions.
	//
	// 5 - VERY LOW - All functions.
	//
	// Any of "1", "2", "3", "4", "5".
	Cpcon SiteStatusNewParamsCpcon `json:"cpcon,omitzero"`
	// Emergency Operations Center (EOC) status.
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Eoc SiteStatusNewParamsEoc `json:"eoc,omitzero"`
	// Force Protection Condition (FPCON).
	//
	// NORMAL - Applies when a general global threat of possible terrorist activity
	// exists and warrants a routine security posture.
	//
	// ALPHA - Applies when an increased general threat of possible terrorist activity
	// against personnel or facilities. Nature and extent of threat are unpredictable.
	//
	// BRAVO - Applies when an increased or predictable threat of terrorist activity
	// exists.
	//
	// CHARLIE - Applies when an incident occurs or intelligence is received indicating
	// some form of terrorist action against personnel and facilities is imminent.
	//
	// DELTA - Applies in the immediate area where an attack has occurred or when
	// intelligence is received indicating terrorist action against a location is
	// imminent.
	//
	// Any of "NORMAL", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Fpcon SiteStatusNewParamsFpcon `json:"fpcon,omitzero"`
	// Health Protection Condition (HPCON).
	//
	// 0 - Routine, no community transmission.
	//
	// ALPHA - Limited, community transmission beginning.
	//
	// BRAVO - Moderate, increased community transmission.
	//
	// CHARLIE - Substantial, sustained community transmission.
	//
	// DELTA - Severe, widespread community transmission.
	//
	// Any of "0", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Hpcon SiteStatusNewParamsHpcon `json:"hpcon,omitzero"`
	// The status of the installation.
	//
	// # FMC - Fully Mission Capable
	//
	// # PMC - Partially Mission Capable
	//
	// # NMC - Non Mission Capable
	//
	// UNK - Unknown.
	//
	// Any of "FMC", "PMC", "NMC", "UNK".
	InstStatus SiteStatusNewParamsInstStatus `json:"instStatus,omitzero"`
	// Array of Link item(s) for which status is available and reported (ATDL, IJMS,
	// LINK-1, LINK-11, LINK-11B, LINK-16). This array must be the same length as the
	// linkStatus array.
	Link []string `json:"link,omitzero"`
	// Array of the status (AVAILABLE, DEGRADED, NOT AVAILABLE, etc.) for each links in
	// the link array. This array must be the same length as the link array, and the
	// status must correspond to the appropriate position index in the link array.
	LinkStatus []string `json:"linkStatus,omitzero"`
	// Array of specific missile types for which an estimated inventory count is
	// available (e.g. GMD TYPE A, HARPOON, TOMAHAWK, etc.). This array must be the
	// same length as the missileInventory array.
	Missile []string `json:"missile,omitzero"`
	// Array of the quantity of each of the missile items. This array must be the same
	// length as the missile array, and the values must correspond to appropriate
	// position index in the missile array.
	MissileInventory []int64 `json:"missileInventory,omitzero"`
	// Array of the status (NON-OPERATIONAL, OPERATIONAL, OFF) for each radar system in
	// the radarSystem array. This array must be the same length as the radarSystem
	// array, and the status must correspond to the appropriate position index in the
	// radarSystem array.
	RadarStatus []string `json:"radarStatus,omitzero"`
	// Array of radar system(s) for which status is available and reported
	// (ACQUISITION, IFFSIF, ILLUMINATING, MODE-4, PRIMARY SURVEILLANCE, SECONDARY
	// SURVEILLANCE, TERTIARY SURVEILLANCE). This array must be the same length as the
	// radarStatus array.
	RadarSystem []string `json:"radarSystem,omitzero"`
	paramObj
}

func (r SiteStatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SiteStatusNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteStatusNewParams) UnmarshalJSON(data []byte) error {
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
type SiteStatusNewParamsDataMode string

const (
	SiteStatusNewParamsDataModeReal      SiteStatusNewParamsDataMode = "REAL"
	SiteStatusNewParamsDataModeTest      SiteStatusNewParamsDataMode = "TEST"
	SiteStatusNewParamsDataModeSimulated SiteStatusNewParamsDataMode = "SIMULATED"
	SiteStatusNewParamsDataModeExercise  SiteStatusNewParamsDataMode = "EXERCISE"
)

// Crisis Action Team (CAT).
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusNewParamsCat string

const (
	SiteStatusNewParamsCatCold SiteStatusNewParamsCat = "COLD"
	SiteStatusNewParamsCatWarm SiteStatusNewParamsCat = "WARM"
	SiteStatusNewParamsCatHot  SiteStatusNewParamsCat = "HOT"
)

// Cyberspace Protection Condition (CPCON).
//
// 1 - VERY HIGH - Critical functions.
//
// 2 - HIGH - Critical and essential functions.
//
// 3 - MEDIUM - Critical, essential, and support functions.
//
// 4 - LOW - All functions.
//
// 5 - VERY LOW - All functions.
type SiteStatusNewParamsCpcon string

const (
	SiteStatusNewParamsCpcon1 SiteStatusNewParamsCpcon = "1"
	SiteStatusNewParamsCpcon2 SiteStatusNewParamsCpcon = "2"
	SiteStatusNewParamsCpcon3 SiteStatusNewParamsCpcon = "3"
	SiteStatusNewParamsCpcon4 SiteStatusNewParamsCpcon = "4"
	SiteStatusNewParamsCpcon5 SiteStatusNewParamsCpcon = "5"
)

// Emergency Operations Center (EOC) status.
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusNewParamsEoc string

const (
	SiteStatusNewParamsEocCold SiteStatusNewParamsEoc = "COLD"
	SiteStatusNewParamsEocWarm SiteStatusNewParamsEoc = "WARM"
	SiteStatusNewParamsEocHot  SiteStatusNewParamsEoc = "HOT"
)

// Force Protection Condition (FPCON).
//
// NORMAL - Applies when a general global threat of possible terrorist activity
// exists and warrants a routine security posture.
//
// ALPHA - Applies when an increased general threat of possible terrorist activity
// against personnel or facilities. Nature and extent of threat are unpredictable.
//
// BRAVO - Applies when an increased or predictable threat of terrorist activity
// exists.
//
// CHARLIE - Applies when an incident occurs or intelligence is received indicating
// some form of terrorist action against personnel and facilities is imminent.
//
// DELTA - Applies in the immediate area where an attack has occurred or when
// intelligence is received indicating terrorist action against a location is
// imminent.
type SiteStatusNewParamsFpcon string

const (
	SiteStatusNewParamsFpconNormal  SiteStatusNewParamsFpcon = "NORMAL"
	SiteStatusNewParamsFpconAlpha   SiteStatusNewParamsFpcon = "ALPHA"
	SiteStatusNewParamsFpconBravo   SiteStatusNewParamsFpcon = "BRAVO"
	SiteStatusNewParamsFpconCharlie SiteStatusNewParamsFpcon = "CHARLIE"
	SiteStatusNewParamsFpconDelta   SiteStatusNewParamsFpcon = "DELTA"
)

// Health Protection Condition (HPCON).
//
// 0 - Routine, no community transmission.
//
// ALPHA - Limited, community transmission beginning.
//
// BRAVO - Moderate, increased community transmission.
//
// CHARLIE - Substantial, sustained community transmission.
//
// DELTA - Severe, widespread community transmission.
type SiteStatusNewParamsHpcon string

const (
	SiteStatusNewParamsHpcon0       SiteStatusNewParamsHpcon = "0"
	SiteStatusNewParamsHpconAlpha   SiteStatusNewParamsHpcon = "ALPHA"
	SiteStatusNewParamsHpconBravo   SiteStatusNewParamsHpcon = "BRAVO"
	SiteStatusNewParamsHpconCharlie SiteStatusNewParamsHpcon = "CHARLIE"
	SiteStatusNewParamsHpconDelta   SiteStatusNewParamsHpcon = "DELTA"
)

// The status of the installation.
//
// # FMC - Fully Mission Capable
//
// # PMC - Partially Mission Capable
//
// # NMC - Non Mission Capable
//
// UNK - Unknown.
type SiteStatusNewParamsInstStatus string

const (
	SiteStatusNewParamsInstStatusFmc SiteStatusNewParamsInstStatus = "FMC"
	SiteStatusNewParamsInstStatusPmc SiteStatusNewParamsInstStatus = "PMC"
	SiteStatusNewParamsInstStatusNmc SiteStatusNewParamsInstStatus = "NMC"
	SiteStatusNewParamsInstStatusUnk SiteStatusNewParamsInstStatus = "UNK"
)

type SiteStatusUpdateParams struct {
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
	DataMode SiteStatusUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The ID of the site, if this status is associated with a fixed site or platform.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Estimated number of cold missiles of all types remaining in weapons system
	// inventory.
	ColdInventory param.Opt[int64] `json:"coldInventory,omitzero"`
	// The communications component causing the platform or system to be less than
	// fully operational.
	CommImpairment param.Opt[string] `json:"commImpairment,omitzero"`
	// Estimated number of hot missiles of all types remaining in weapons system
	// inventory.
	HotInventory param.Opt[int64] `json:"hotInventory,omitzero"`
	// Alternate Identifier for a mobile or transportable platform provided by source.
	MobileAltID param.Opt[string] `json:"mobileAltId,omitzero"`
	// The operational status of the platform (e.g. Fully Operational, Partially
	// Operational, Not Operational, etc.).
	OpsCapability param.Opt[string] `json:"opsCapability,omitzero"`
	// The primary component degrading the operational capability of the platform or
	// system.
	OpsImpairment param.Opt[string] `json:"opsImpairment,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Position Engagement Status flag, Indicating whether this platform is initiating
	// multiple simultaneous engagements. A value of 1/True indicates the platform is
	// initiating multiple simultaneous engagements.
	Pes param.Opt[bool] `json:"pes,omitzero"`
	// The POI (point of interest) ID related to this platform, if available.
	Poiid param.Opt[string] `json:"poiid,omitzero"`
	// SAM sensor radar surveillance mode (Active, Passive, Off).
	RadiateMode param.Opt[string] `json:"radiateMode,omitzero"`
	// Time of report, in ISO8601 UTC format.
	ReportTime param.Opt[time.Time] `json:"reportTime,omitzero" format:"date-time"`
	// The state of a SAM unit (e.g. Initialization, Standby, Reorientation, etc.).
	SamMode param.Opt[string] `json:"samMode,omitzero"`
	// Optional site type or further detail of type. Intended for, but not limited to,
	// Link-16 site type specifications (e.g. ADOC, GACC, SOC, TACC, etc.).
	SiteType param.Opt[string] `json:"siteType,omitzero"`
	// Description of the time function associated with the reportTime (e.g.
	// Activation, Deactivation, Arrival, Departure, etc.), if applicable.
	TimeFunction param.Opt[string] `json:"timeFunction,omitzero"`
	// The track ID related to this platform (if mobile or transportable), if
	// available.
	TrackID param.Opt[string] `json:"trackId,omitzero"`
	// Link-16 specific reference track number.
	TrackRefL16 param.Opt[string] `json:"trackRefL16,omitzero"`
	// Description of the current weather conditions over a site.
	WeatherMessage param.Opt[string] `json:"weatherMessage,omitzero"`
	// Crisis Action Team (CAT).
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Cat SiteStatusUpdateParamsCat `json:"cat,omitzero"`
	// Cyberspace Protection Condition (CPCON).
	//
	// 1 - VERY HIGH - Critical functions.
	//
	// 2 - HIGH - Critical and essential functions.
	//
	// 3 - MEDIUM - Critical, essential, and support functions.
	//
	// 4 - LOW - All functions.
	//
	// 5 - VERY LOW - All functions.
	//
	// Any of "1", "2", "3", "4", "5".
	Cpcon SiteStatusUpdateParamsCpcon `json:"cpcon,omitzero"`
	// Emergency Operations Center (EOC) status.
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Eoc SiteStatusUpdateParamsEoc `json:"eoc,omitzero"`
	// Force Protection Condition (FPCON).
	//
	// NORMAL - Applies when a general global threat of possible terrorist activity
	// exists and warrants a routine security posture.
	//
	// ALPHA - Applies when an increased general threat of possible terrorist activity
	// against personnel or facilities. Nature and extent of threat are unpredictable.
	//
	// BRAVO - Applies when an increased or predictable threat of terrorist activity
	// exists.
	//
	// CHARLIE - Applies when an incident occurs or intelligence is received indicating
	// some form of terrorist action against personnel and facilities is imminent.
	//
	// DELTA - Applies in the immediate area where an attack has occurred or when
	// intelligence is received indicating terrorist action against a location is
	// imminent.
	//
	// Any of "NORMAL", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Fpcon SiteStatusUpdateParamsFpcon `json:"fpcon,omitzero"`
	// Health Protection Condition (HPCON).
	//
	// 0 - Routine, no community transmission.
	//
	// ALPHA - Limited, community transmission beginning.
	//
	// BRAVO - Moderate, increased community transmission.
	//
	// CHARLIE - Substantial, sustained community transmission.
	//
	// DELTA - Severe, widespread community transmission.
	//
	// Any of "0", "ALPHA", "BRAVO", "CHARLIE", "DELTA".
	Hpcon SiteStatusUpdateParamsHpcon `json:"hpcon,omitzero"`
	// The status of the installation.
	//
	// # FMC - Fully Mission Capable
	//
	// # PMC - Partially Mission Capable
	//
	// # NMC - Non Mission Capable
	//
	// UNK - Unknown.
	//
	// Any of "FMC", "PMC", "NMC", "UNK".
	InstStatus SiteStatusUpdateParamsInstStatus `json:"instStatus,omitzero"`
	// Array of Link item(s) for which status is available and reported (ATDL, IJMS,
	// LINK-1, LINK-11, LINK-11B, LINK-16). This array must be the same length as the
	// linkStatus array.
	Link []string `json:"link,omitzero"`
	// Array of the status (AVAILABLE, DEGRADED, NOT AVAILABLE, etc.) for each links in
	// the link array. This array must be the same length as the link array, and the
	// status must correspond to the appropriate position index in the link array.
	LinkStatus []string `json:"linkStatus,omitzero"`
	// Array of specific missile types for which an estimated inventory count is
	// available (e.g. GMD TYPE A, HARPOON, TOMAHAWK, etc.). This array must be the
	// same length as the missileInventory array.
	Missile []string `json:"missile,omitzero"`
	// Array of the quantity of each of the missile items. This array must be the same
	// length as the missile array, and the values must correspond to appropriate
	// position index in the missile array.
	MissileInventory []int64 `json:"missileInventory,omitzero"`
	// Array of the status (NON-OPERATIONAL, OPERATIONAL, OFF) for each radar system in
	// the radarSystem array. This array must be the same length as the radarSystem
	// array, and the status must correspond to the appropriate position index in the
	// radarSystem array.
	RadarStatus []string `json:"radarStatus,omitzero"`
	// Array of radar system(s) for which status is available and reported
	// (ACQUISITION, IFFSIF, ILLUMINATING, MODE-4, PRIMARY SURVEILLANCE, SECONDARY
	// SURVEILLANCE, TERTIARY SURVEILLANCE). This array must be the same length as the
	// radarStatus array.
	RadarSystem []string `json:"radarSystem,omitzero"`
	paramObj
}

func (r SiteStatusUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SiteStatusUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteStatusUpdateParams) UnmarshalJSON(data []byte) error {
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
type SiteStatusUpdateParamsDataMode string

const (
	SiteStatusUpdateParamsDataModeReal      SiteStatusUpdateParamsDataMode = "REAL"
	SiteStatusUpdateParamsDataModeTest      SiteStatusUpdateParamsDataMode = "TEST"
	SiteStatusUpdateParamsDataModeSimulated SiteStatusUpdateParamsDataMode = "SIMULATED"
	SiteStatusUpdateParamsDataModeExercise  SiteStatusUpdateParamsDataMode = "EXERCISE"
)

// Crisis Action Team (CAT).
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusUpdateParamsCat string

const (
	SiteStatusUpdateParamsCatCold SiteStatusUpdateParamsCat = "COLD"
	SiteStatusUpdateParamsCatWarm SiteStatusUpdateParamsCat = "WARM"
	SiteStatusUpdateParamsCatHot  SiteStatusUpdateParamsCat = "HOT"
)

// Cyberspace Protection Condition (CPCON).
//
// 1 - VERY HIGH - Critical functions.
//
// 2 - HIGH - Critical and essential functions.
//
// 3 - MEDIUM - Critical, essential, and support functions.
//
// 4 - LOW - All functions.
//
// 5 - VERY LOW - All functions.
type SiteStatusUpdateParamsCpcon string

const (
	SiteStatusUpdateParamsCpcon1 SiteStatusUpdateParamsCpcon = "1"
	SiteStatusUpdateParamsCpcon2 SiteStatusUpdateParamsCpcon = "2"
	SiteStatusUpdateParamsCpcon3 SiteStatusUpdateParamsCpcon = "3"
	SiteStatusUpdateParamsCpcon4 SiteStatusUpdateParamsCpcon = "4"
	SiteStatusUpdateParamsCpcon5 SiteStatusUpdateParamsCpcon = "5"
)

// Emergency Operations Center (EOC) status.
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SiteStatusUpdateParamsEoc string

const (
	SiteStatusUpdateParamsEocCold SiteStatusUpdateParamsEoc = "COLD"
	SiteStatusUpdateParamsEocWarm SiteStatusUpdateParamsEoc = "WARM"
	SiteStatusUpdateParamsEocHot  SiteStatusUpdateParamsEoc = "HOT"
)

// Force Protection Condition (FPCON).
//
// NORMAL - Applies when a general global threat of possible terrorist activity
// exists and warrants a routine security posture.
//
// ALPHA - Applies when an increased general threat of possible terrorist activity
// against personnel or facilities. Nature and extent of threat are unpredictable.
//
// BRAVO - Applies when an increased or predictable threat of terrorist activity
// exists.
//
// CHARLIE - Applies when an incident occurs or intelligence is received indicating
// some form of terrorist action against personnel and facilities is imminent.
//
// DELTA - Applies in the immediate area where an attack has occurred or when
// intelligence is received indicating terrorist action against a location is
// imminent.
type SiteStatusUpdateParamsFpcon string

const (
	SiteStatusUpdateParamsFpconNormal  SiteStatusUpdateParamsFpcon = "NORMAL"
	SiteStatusUpdateParamsFpconAlpha   SiteStatusUpdateParamsFpcon = "ALPHA"
	SiteStatusUpdateParamsFpconBravo   SiteStatusUpdateParamsFpcon = "BRAVO"
	SiteStatusUpdateParamsFpconCharlie SiteStatusUpdateParamsFpcon = "CHARLIE"
	SiteStatusUpdateParamsFpconDelta   SiteStatusUpdateParamsFpcon = "DELTA"
)

// Health Protection Condition (HPCON).
//
// 0 - Routine, no community transmission.
//
// ALPHA - Limited, community transmission beginning.
//
// BRAVO - Moderate, increased community transmission.
//
// CHARLIE - Substantial, sustained community transmission.
//
// DELTA - Severe, widespread community transmission.
type SiteStatusUpdateParamsHpcon string

const (
	SiteStatusUpdateParamsHpcon0       SiteStatusUpdateParamsHpcon = "0"
	SiteStatusUpdateParamsHpconAlpha   SiteStatusUpdateParamsHpcon = "ALPHA"
	SiteStatusUpdateParamsHpconBravo   SiteStatusUpdateParamsHpcon = "BRAVO"
	SiteStatusUpdateParamsHpconCharlie SiteStatusUpdateParamsHpcon = "CHARLIE"
	SiteStatusUpdateParamsHpconDelta   SiteStatusUpdateParamsHpcon = "DELTA"
)

// The status of the installation.
//
// # FMC - Fully Mission Capable
//
// # PMC - Partially Mission Capable
//
// # NMC - Non Mission Capable
//
// UNK - Unknown.
type SiteStatusUpdateParamsInstStatus string

const (
	SiteStatusUpdateParamsInstStatusFmc SiteStatusUpdateParamsInstStatus = "FMC"
	SiteStatusUpdateParamsInstStatusPmc SiteStatusUpdateParamsInstStatus = "PMC"
	SiteStatusUpdateParamsInstStatusNmc SiteStatusUpdateParamsInstStatus = "NMC"
	SiteStatusUpdateParamsInstStatusUnk SiteStatusUpdateParamsInstStatus = "UNK"
)

type SiteStatusListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteStatusListParams]'s query parameters as `url.Values`.
func (r SiteStatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteStatusCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteStatusCountParams]'s query parameters as `url.Values`.
func (r SiteStatusCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteStatusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteStatusGetParams]'s query parameters as `url.Values`.
func (r SiteStatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteStatusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteStatusTupleParams]'s query parameters as `url.Values`.
func (r SiteStatusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
