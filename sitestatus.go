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

// SitestatusService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSitestatusService] method instead.
type SitestatusService struct {
	Options []option.RequestOption
	History SitestatusHistoryService
}

// NewSitestatusService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSitestatusService(opts ...option.RequestOption) (r SitestatusService) {
	r = SitestatusService{}
	r.Options = opts
	r.History = NewSitestatusHistoryService(opts...)
	return
}

// Service operation to take a single SiteStatus object as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SitestatusService) New(ctx context.Context, body SitestatusNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sitestatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single SiteStatus object. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SitestatusService) Update(ctx context.Context, id string, body SitestatusUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
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
func (r *SitestatusService) List(ctx context.Context, query SitestatusListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SitestatusListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *SitestatusService) ListAutoPaging(ctx context.Context, query SitestatusListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SitestatusListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SiteStatus object specified by the passed ID path
// parameter. Note, delete operations do not remove data from historical or
// publish/subscribe stores. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SitestatusService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
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
func (r *SitestatusService) Count(ctx context.Context, query SitestatusCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sitestatus/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SiteStatus record by its unique ID passed as a
// path parameter.
func (r *SitestatusService) Get(ctx context.Context, id string, query SitestatusGetParams, opts ...option.RequestOption) (res *SitestatusFull, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *SitestatusService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sitestatus/queryhelp"
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
func (r *SitestatusService) Tuple(ctx context.Context, query SitestatusTupleParams, opts ...option.RequestOption) (res *[]SitestatusFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sitestatus/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SitestatusListResponse struct {
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
	DataMode SitestatusListResponseDataMode `json:"dataMode,required"`
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
	Cat SitestatusListResponseCat `json:"cat"`
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
	Cpcon SitestatusListResponseCpcon `json:"cpcon"`
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
	Eoc SitestatusListResponseEoc `json:"eoc"`
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
	Fpcon SitestatusListResponseFpcon `json:"fpcon"`
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
	Hpcon SitestatusListResponseHpcon `json:"hpcon"`
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
	InstStatus SitestatusListResponseInstStatus `json:"instStatus"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDSite                resp.Field
		Source                resp.Field
		ID                    resp.Field
		Cat                   resp.Field
		ColdInventory         resp.Field
		CommImpairment        resp.Field
		Cpcon                 resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Eoc                   resp.Field
		Fpcon                 resp.Field
		HotInventory          resp.Field
		Hpcon                 resp.Field
		InstStatus            resp.Field
		Link                  resp.Field
		LinkStatus            resp.Field
		Missile               resp.Field
		MissileInventory      resp.Field
		MobileAltID           resp.Field
		OpsCapability         resp.Field
		OpsImpairment         resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Pes                   resp.Field
		Poiid                 resp.Field
		RadarStatus           resp.Field
		RadarSystem           resp.Field
		RadiateMode           resp.Field
		ReportTime            resp.Field
		SamMode               resp.Field
		SiteType              resp.Field
		TimeFunction          resp.Field
		TrackID               resp.Field
		TrackRefL16           resp.Field
		WeatherMessage        resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitestatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *SitestatusListResponse) UnmarshalJSON(data []byte) error {
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
type SitestatusListResponseDataMode string

const (
	SitestatusListResponseDataModeReal      SitestatusListResponseDataMode = "REAL"
	SitestatusListResponseDataModeTest      SitestatusListResponseDataMode = "TEST"
	SitestatusListResponseDataModeSimulated SitestatusListResponseDataMode = "SIMULATED"
	SitestatusListResponseDataModeExercise  SitestatusListResponseDataMode = "EXERCISE"
)

// Crisis Action Team (CAT).
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SitestatusListResponseCat string

const (
	SitestatusListResponseCatCold SitestatusListResponseCat = "COLD"
	SitestatusListResponseCatWarm SitestatusListResponseCat = "WARM"
	SitestatusListResponseCatHot  SitestatusListResponseCat = "HOT"
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
type SitestatusListResponseCpcon string

const (
	SitestatusListResponseCpcon1 SitestatusListResponseCpcon = "1"
	SitestatusListResponseCpcon2 SitestatusListResponseCpcon = "2"
	SitestatusListResponseCpcon3 SitestatusListResponseCpcon = "3"
	SitestatusListResponseCpcon4 SitestatusListResponseCpcon = "4"
	SitestatusListResponseCpcon5 SitestatusListResponseCpcon = "5"
)

// Emergency Operations Center (EOC) status.
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SitestatusListResponseEoc string

const (
	SitestatusListResponseEocCold SitestatusListResponseEoc = "COLD"
	SitestatusListResponseEocWarm SitestatusListResponseEoc = "WARM"
	SitestatusListResponseEocHot  SitestatusListResponseEoc = "HOT"
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
type SitestatusListResponseFpcon string

const (
	SitestatusListResponseFpconNormal  SitestatusListResponseFpcon = "NORMAL"
	SitestatusListResponseFpconAlpha   SitestatusListResponseFpcon = "ALPHA"
	SitestatusListResponseFpconBravo   SitestatusListResponseFpcon = "BRAVO"
	SitestatusListResponseFpconCharlie SitestatusListResponseFpcon = "CHARLIE"
	SitestatusListResponseFpconDelta   SitestatusListResponseFpcon = "DELTA"
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
type SitestatusListResponseHpcon string

const (
	SitestatusListResponseHpcon0       SitestatusListResponseHpcon = "0"
	SitestatusListResponseHpconAlpha   SitestatusListResponseHpcon = "ALPHA"
	SitestatusListResponseHpconBravo   SitestatusListResponseHpcon = "BRAVO"
	SitestatusListResponseHpconCharlie SitestatusListResponseHpcon = "CHARLIE"
	SitestatusListResponseHpconDelta   SitestatusListResponseHpcon = "DELTA"
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
type SitestatusListResponseInstStatus string

const (
	SitestatusListResponseInstStatusFmc SitestatusListResponseInstStatus = "FMC"
	SitestatusListResponseInstStatusPmc SitestatusListResponseInstStatus = "PMC"
	SitestatusListResponseInstStatusNmc SitestatusListResponseInstStatus = "NMC"
	SitestatusListResponseInstStatusUnk SitestatusListResponseInstStatus = "UNK"
)

type SitestatusNewParams struct {
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
	DataMode SitestatusNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	Cat SitestatusNewParamsCat `json:"cat,omitzero"`
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
	Cpcon SitestatusNewParamsCpcon `json:"cpcon,omitzero"`
	// Emergency Operations Center (EOC) status.
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Eoc SitestatusNewParamsEoc `json:"eoc,omitzero"`
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
	Fpcon SitestatusNewParamsFpcon `json:"fpcon,omitzero"`
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
	Hpcon SitestatusNewParamsHpcon `json:"hpcon,omitzero"`
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
	InstStatus SitestatusNewParamsInstStatus `json:"instStatus,omitzero"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SitestatusNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SitestatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SitestatusNewParams
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
type SitestatusNewParamsDataMode string

const (
	SitestatusNewParamsDataModeReal      SitestatusNewParamsDataMode = "REAL"
	SitestatusNewParamsDataModeTest      SitestatusNewParamsDataMode = "TEST"
	SitestatusNewParamsDataModeSimulated SitestatusNewParamsDataMode = "SIMULATED"
	SitestatusNewParamsDataModeExercise  SitestatusNewParamsDataMode = "EXERCISE"
)

// Crisis Action Team (CAT).
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SitestatusNewParamsCat string

const (
	SitestatusNewParamsCatCold SitestatusNewParamsCat = "COLD"
	SitestatusNewParamsCatWarm SitestatusNewParamsCat = "WARM"
	SitestatusNewParamsCatHot  SitestatusNewParamsCat = "HOT"
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
type SitestatusNewParamsCpcon string

const (
	SitestatusNewParamsCpcon1 SitestatusNewParamsCpcon = "1"
	SitestatusNewParamsCpcon2 SitestatusNewParamsCpcon = "2"
	SitestatusNewParamsCpcon3 SitestatusNewParamsCpcon = "3"
	SitestatusNewParamsCpcon4 SitestatusNewParamsCpcon = "4"
	SitestatusNewParamsCpcon5 SitestatusNewParamsCpcon = "5"
)

// Emergency Operations Center (EOC) status.
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SitestatusNewParamsEoc string

const (
	SitestatusNewParamsEocCold SitestatusNewParamsEoc = "COLD"
	SitestatusNewParamsEocWarm SitestatusNewParamsEoc = "WARM"
	SitestatusNewParamsEocHot  SitestatusNewParamsEoc = "HOT"
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
type SitestatusNewParamsFpcon string

const (
	SitestatusNewParamsFpconNormal  SitestatusNewParamsFpcon = "NORMAL"
	SitestatusNewParamsFpconAlpha   SitestatusNewParamsFpcon = "ALPHA"
	SitestatusNewParamsFpconBravo   SitestatusNewParamsFpcon = "BRAVO"
	SitestatusNewParamsFpconCharlie SitestatusNewParamsFpcon = "CHARLIE"
	SitestatusNewParamsFpconDelta   SitestatusNewParamsFpcon = "DELTA"
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
type SitestatusNewParamsHpcon string

const (
	SitestatusNewParamsHpcon0       SitestatusNewParamsHpcon = "0"
	SitestatusNewParamsHpconAlpha   SitestatusNewParamsHpcon = "ALPHA"
	SitestatusNewParamsHpconBravo   SitestatusNewParamsHpcon = "BRAVO"
	SitestatusNewParamsHpconCharlie SitestatusNewParamsHpcon = "CHARLIE"
	SitestatusNewParamsHpconDelta   SitestatusNewParamsHpcon = "DELTA"
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
type SitestatusNewParamsInstStatus string

const (
	SitestatusNewParamsInstStatusFmc SitestatusNewParamsInstStatus = "FMC"
	SitestatusNewParamsInstStatusPmc SitestatusNewParamsInstStatus = "PMC"
	SitestatusNewParamsInstStatusNmc SitestatusNewParamsInstStatus = "NMC"
	SitestatusNewParamsInstStatusUnk SitestatusNewParamsInstStatus = "UNK"
)

type SitestatusUpdateParams struct {
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
	DataMode SitestatusUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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
	Cat SitestatusUpdateParamsCat `json:"cat,omitzero"`
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
	Cpcon SitestatusUpdateParamsCpcon `json:"cpcon,omitzero"`
	// Emergency Operations Center (EOC) status.
	//
	// COLD - Not in use.
	//
	// WARM - Facility prepped/possible skeleton crew.
	//
	// HOT - Fully active.
	//
	// Any of "COLD", "WARM", "HOT".
	Eoc SitestatusUpdateParamsEoc `json:"eoc,omitzero"`
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
	Fpcon SitestatusUpdateParamsFpcon `json:"fpcon,omitzero"`
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
	Hpcon SitestatusUpdateParamsHpcon `json:"hpcon,omitzero"`
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
	InstStatus SitestatusUpdateParamsInstStatus `json:"instStatus,omitzero"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SitestatusUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SitestatusUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SitestatusUpdateParams
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
type SitestatusUpdateParamsDataMode string

const (
	SitestatusUpdateParamsDataModeReal      SitestatusUpdateParamsDataMode = "REAL"
	SitestatusUpdateParamsDataModeTest      SitestatusUpdateParamsDataMode = "TEST"
	SitestatusUpdateParamsDataModeSimulated SitestatusUpdateParamsDataMode = "SIMULATED"
	SitestatusUpdateParamsDataModeExercise  SitestatusUpdateParamsDataMode = "EXERCISE"
)

// Crisis Action Team (CAT).
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SitestatusUpdateParamsCat string

const (
	SitestatusUpdateParamsCatCold SitestatusUpdateParamsCat = "COLD"
	SitestatusUpdateParamsCatWarm SitestatusUpdateParamsCat = "WARM"
	SitestatusUpdateParamsCatHot  SitestatusUpdateParamsCat = "HOT"
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
type SitestatusUpdateParamsCpcon string

const (
	SitestatusUpdateParamsCpcon1 SitestatusUpdateParamsCpcon = "1"
	SitestatusUpdateParamsCpcon2 SitestatusUpdateParamsCpcon = "2"
	SitestatusUpdateParamsCpcon3 SitestatusUpdateParamsCpcon = "3"
	SitestatusUpdateParamsCpcon4 SitestatusUpdateParamsCpcon = "4"
	SitestatusUpdateParamsCpcon5 SitestatusUpdateParamsCpcon = "5"
)

// Emergency Operations Center (EOC) status.
//
// COLD - Not in use.
//
// WARM - Facility prepped/possible skeleton crew.
//
// HOT - Fully active.
type SitestatusUpdateParamsEoc string

const (
	SitestatusUpdateParamsEocCold SitestatusUpdateParamsEoc = "COLD"
	SitestatusUpdateParamsEocWarm SitestatusUpdateParamsEoc = "WARM"
	SitestatusUpdateParamsEocHot  SitestatusUpdateParamsEoc = "HOT"
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
type SitestatusUpdateParamsFpcon string

const (
	SitestatusUpdateParamsFpconNormal  SitestatusUpdateParamsFpcon = "NORMAL"
	SitestatusUpdateParamsFpconAlpha   SitestatusUpdateParamsFpcon = "ALPHA"
	SitestatusUpdateParamsFpconBravo   SitestatusUpdateParamsFpcon = "BRAVO"
	SitestatusUpdateParamsFpconCharlie SitestatusUpdateParamsFpcon = "CHARLIE"
	SitestatusUpdateParamsFpconDelta   SitestatusUpdateParamsFpcon = "DELTA"
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
type SitestatusUpdateParamsHpcon string

const (
	SitestatusUpdateParamsHpcon0       SitestatusUpdateParamsHpcon = "0"
	SitestatusUpdateParamsHpconAlpha   SitestatusUpdateParamsHpcon = "ALPHA"
	SitestatusUpdateParamsHpconBravo   SitestatusUpdateParamsHpcon = "BRAVO"
	SitestatusUpdateParamsHpconCharlie SitestatusUpdateParamsHpcon = "CHARLIE"
	SitestatusUpdateParamsHpconDelta   SitestatusUpdateParamsHpcon = "DELTA"
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
type SitestatusUpdateParamsInstStatus string

const (
	SitestatusUpdateParamsInstStatusFmc SitestatusUpdateParamsInstStatus = "FMC"
	SitestatusUpdateParamsInstStatusPmc SitestatusUpdateParamsInstStatus = "PMC"
	SitestatusUpdateParamsInstStatusNmc SitestatusUpdateParamsInstStatus = "NMC"
	SitestatusUpdateParamsInstStatusUnk SitestatusUpdateParamsInstStatus = "UNK"
)

type SitestatusListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SitestatusListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SitestatusListParams]'s query parameters as `url.Values`.
func (r SitestatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SitestatusCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SitestatusCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SitestatusCountParams]'s query parameters as `url.Values`.
func (r SitestatusCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SitestatusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SitestatusGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SitestatusGetParams]'s query parameters as `url.Values`.
func (r SitestatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SitestatusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SitestatusTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SitestatusTupleParams]'s query parameters as `url.Values`.
func (r SitestatusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
