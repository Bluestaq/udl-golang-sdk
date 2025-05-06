// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// IsrCollectionService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIsrCollectionService] method instead.
type IsrCollectionService struct {
	Options []option.RequestOption
	History IsrCollectionHistoryService
}

// NewIsrCollectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIsrCollectionService(opts ...option.RequestOption) (r IsrCollectionService) {
	r = IsrCollectionService{}
	r.Options = opts
	r.History = NewIsrCollectionHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *IsrCollectionService) List(ctx context.Context, query IsrCollectionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[IsrCollectionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/isrcollection"
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
func (r *IsrCollectionService) ListAutoPaging(ctx context.Context, query IsrCollectionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[IsrCollectionListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *IsrCollectionService) Count(ctx context.Context, query IsrCollectionCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/isrcollection/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// ISRCollection records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *IsrCollectionService) NewBulk(ctx context.Context, body IsrCollectionNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/isrcollection/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *IsrCollectionService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/isrcollection/queryhelp"
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
func (r *IsrCollectionService) Tuple(ctx context.Context, query IsrCollectionTupleParams, opts ...option.RequestOption) (res *[]IsrCollectionTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/isrcollection/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple ISR Collections as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *IsrCollectionService) UnvalidatedPublish(ctx context.Context, body IsrCollectionUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-isrcollection"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// ISR Collection data.
type IsrCollectionListResponse struct {
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
	DataMode IsrCollectionListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Mission desired collection requirements.
	CollectionRequirements []IsrCollectionListResponseCollectionRequirement `json:"collectionRequirements"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Version of the IDEX software the request came from for compatibility.
	IdexVersion int64 `json:"idexVersion"`
	// Designation of mission Area Of Responsibility.
	MissionAor string `json:"missionAOR"`
	// Mission geographical collection area.
	MissionCollectionArea string `json:"missionCollectionArea"`
	// Country code of the mission. A Country may represent countries, multi-national
	// consortiums, and international organizations.
	MissionCountry string `json:"missionCountry"`
	// Text version of what we are emphasizing in this mission.
	MissionEmphasis string `json:"missionEmphasis"`
	// Mission Identifier.
	MissionID string `json:"missionId"`
	// Joint Operations Area.
	MissionJoa string `json:"missionJoa"`
	// Mission operation name.
	MissionOperation string `json:"missionOperation"`
	// Primary type of intelligence to be collected during the mission.
	MissionPrimaryIntelDiscipline string `json:"missionPrimaryIntelDiscipline"`
	// Sub category of primary intelligence to be collected.
	MissionPrimarySubCategory string `json:"missionPrimarySubCategory"`
	// Mission Priority (1-n).
	MissionPriority int64 `json:"missionPriority"`
	// Region of the mission.
	MissionRegion string `json:"missionRegion"`
	// What is the primary objective(Role) of this mission.
	MissionRole string `json:"missionRole"`
	// Type of intelligence to be collected second.
	MissionSecondaryIntelDiscipline string `json:"missionSecondaryIntelDiscipline"`
	// Mission sub category for secondary intelligence discipline to be collected.
	MissionSecondarySubCategory string `json:"missionSecondarySubCategory"`
	// WGS-84 latitude of the start position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	MissionStartPointLat float64 `json:"missionStartPointLat"`
	// WGS-84 longitude of the start position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	MissionStartPointLong float64 `json:"missionStartPointLong"`
	// Subregion of the mission.
	MissionSubRegion string `json:"missionSubRegion"`
	// Name of the Supporting unit/Location that is performing this mission.
	MissionSupportedUnit string `json:"missionSupportedUnit"`
	// A synchronization matrix is used to organize the logistics synchronization
	// process during a mission.
	MissionSyncMatrixBin string `json:"missionSyncMatrixBin"`
	// Human readable Mission Name.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Individual taskings to complete the mission.
	Taskings []IsrCollectionListResponseTasking `json:"taskings"`
	// Object for data dissemination.
	Transit []IsrCollectionListResponseTransit `json:"transit"`
	// Time the row was updated in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking           respjson.Field
		DataMode                        respjson.Field
		Source                          respjson.Field
		ID                              respjson.Field
		CollectionRequirements          respjson.Field
		CreatedAt                       respjson.Field
		CreatedBy                       respjson.Field
		IdexVersion                     respjson.Field
		MissionAor                      respjson.Field
		MissionCollectionArea           respjson.Field
		MissionCountry                  respjson.Field
		MissionEmphasis                 respjson.Field
		MissionID                       respjson.Field
		MissionJoa                      respjson.Field
		MissionOperation                respjson.Field
		MissionPrimaryIntelDiscipline   respjson.Field
		MissionPrimarySubCategory       respjson.Field
		MissionPriority                 respjson.Field
		MissionRegion                   respjson.Field
		MissionRole                     respjson.Field
		MissionSecondaryIntelDiscipline respjson.Field
		MissionSecondarySubCategory     respjson.Field
		MissionStartPointLat            respjson.Field
		MissionStartPointLong           respjson.Field
		MissionSubRegion                respjson.Field
		MissionSupportedUnit            respjson.Field
		MissionSyncMatrixBin            respjson.Field
		Name                            respjson.Field
		Origin                          respjson.Field
		OrigNetwork                     respjson.Field
		Taskings                        respjson.Field
		Transit                         respjson.Field
		UpdatedAt                       respjson.Field
		UpdatedBy                       respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionListResponse) UnmarshalJSON(data []byte) error {
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
type IsrCollectionListResponseDataMode string

const (
	IsrCollectionListResponseDataModeReal      IsrCollectionListResponseDataMode = "REAL"
	IsrCollectionListResponseDataModeTest      IsrCollectionListResponseDataMode = "TEST"
	IsrCollectionListResponseDataModeSimulated IsrCollectionListResponseDataMode = "SIMULATED"
	IsrCollectionListResponseDataModeExercise  IsrCollectionListResponseDataMode = "EXERCISE"
)

type IsrCollectionListResponseCollectionRequirement struct {
	// Collection Requirement Unique Identifier.
	ID string `json:"id"`
	// Country code of the collection requirement. A Country may represent countries,
	// multi-national consortiums, and international organizations.
	Country string `json:"country"`
	// Collection Requirement Unique Identifier.
	CridNumbers   string                                                      `json:"cridNumbers"`
	CriticalTimes IsrCollectionListResponseCollectionRequirementCriticalTimes `json:"criticalTimes"`
	// Is this collection requirement an emphasized/critical requirement.
	Emphasized              bool                                                                  `json:"emphasized"`
	ExploitationRequirement IsrCollectionListResponseCollectionRequirementExploitationRequirement `json:"exploitationRequirement"`
	// Encryption hashing algorithm.
	Hash string `json:"hash"`
	// Primary type of intelligence to be collected for this requirement.
	IntelDiscipline string `json:"intelDiscipline"`
	// Is this collection request for the Prism system?.
	IsPrismCr bool `json:"isPrismCr"`
	// Human readable name for this operation.
	Operation string `json:"operation"`
	// 1-n priority for this collection requirement.
	Priority float64 `json:"priority"`
	// Reconnaissance Survey information the operator needs.
	ReconSurvey string `json:"reconSurvey"`
	// Record id.
	RecordID string `json:"recordId"`
	// Region of the collection requirement.
	Region string `json:"region"`
	// Sub category of primary intelligence to be collected for this requirement.
	Secondary bool `json:"secondary"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance string `json:"specialComGuidance"`
	// Start time for this requirement, should be within the mission time window.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time for this requirement, should be within the mission time window.
	Stop time.Time `json:"stop" format:"date-time"`
	// Subregion of the collection requirement.
	Subregion string `json:"subregion"`
	// The name of the military unit that this assigned collection requirement will
	// support.
	SupportedUnit string `json:"supportedUnit"`
	// Array of POI Id's for the targets being tasked.
	TargetList []string `json:"targetList"`
	// Type collection this requirement applies to.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Country                 respjson.Field
		CridNumbers             respjson.Field
		CriticalTimes           respjson.Field
		Emphasized              respjson.Field
		ExploitationRequirement respjson.Field
		Hash                    respjson.Field
		IntelDiscipline         respjson.Field
		IsPrismCr               respjson.Field
		Operation               respjson.Field
		Priority                respjson.Field
		ReconSurvey             respjson.Field
		RecordID                respjson.Field
		Region                  respjson.Field
		Secondary               respjson.Field
		SpecialComGuidance      respjson.Field
		Start                   respjson.Field
		Stop                    respjson.Field
		Subregion               respjson.Field
		SupportedUnit           respjson.Field
		TargetList              respjson.Field
		Type                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseCollectionRequirement) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionListResponseCollectionRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseCollectionRequirementCriticalTimes struct {
	// Critical start time to collect an image for this requirement.
	EarliestImagingTime time.Time `json:"earliestImagingTime,required" format:"date-time"`
	// Critical stop time to collect an image for this requirement.
	LatestImagingTime time.Time `json:"latestImagingTime,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EarliestImagingTime respjson.Field
		LatestImagingTime   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseCollectionRequirementCriticalTimes) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionListResponseCollectionRequirementCriticalTimes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseCollectionRequirementExploitationRequirement struct {
	// Exploitation requirement id.
	ID string `json:"id"`
	// Amplifying data for the exploitation requirement.
	Amplification string `json:"amplification"`
	// List of e-mails to disseminate collection verification information.
	Dissemination string `json:"dissemination"`
	// Essential Elements of Information.
	Eei string                                                                   `json:"eei"`
	Poc IsrCollectionListResponseCollectionRequirementExploitationRequirementPoc `json:"poc"`
	// The reporting criteria of the collection requirement.
	ReportingCriteria string `json:"reportingCriteria"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Amplification     respjson.Field
		Dissemination     respjson.Field
		Eei               respjson.Field
		Poc               respjson.Field
		ReportingCriteria respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseCollectionRequirementExploitationRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionListResponseCollectionRequirementExploitationRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseCollectionRequirementExploitationRequirementPoc struct {
	// Unique identifier of the collection requirement POC.
	ID string `json:"id"`
	// Callsign of the POC.
	Callsign string `json:"callsign"`
	// Chat name of the POC.
	ChatName string `json:"chatName"`
	// Chat system the POC is accessing.
	ChatSystem string `json:"chatSystem"`
	// Email address of the POC.
	Email string `json:"email"`
	// Name of the POC.
	Name string `json:"name"`
	// Amplifying notes about the POC.
	Notes string `json:"notes"`
	// Phone number of the POC.
	Phone string `json:"phone"`
	// Radio Frequency the POC is on.
	RadioFrequency float64 `json:"radioFrequency"`
	// Unit the POC belongs to.
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Callsign       respjson.Field
		ChatName       respjson.Field
		ChatSystem     respjson.Field
		Email          respjson.Field
		Name           respjson.Field
		Notes          respjson.Field
		Phone          respjson.Field
		RadioFrequency respjson.Field
		Unit           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseCollectionRequirementExploitationRequirementPoc) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionListResponseCollectionRequirementExploitationRequirementPoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTasking struct {
	// Tasking Unique Identifier.
	ID                string                                            `json:"id"`
	CollectionPeriods IsrCollectionListResponseTaskingCollectionPeriods `json:"collectionPeriods"`
	// Type of collection tasked.
	//
	// Any of "Simultaneous", "Sequential", "Operationally", "Driven", "Priority",
	// "Order".
	CollectionType string `json:"collectionType"`
	// Eight line.
	EightLine string `json:"eightLine"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance string `json:"specialComGuidance"`
	// Value of the Sensitive Reconnaissance Operations Track.
	SroTrack string `json:"sroTrack"`
	// Human readable definition of this taskings Area Of Responsibility.
	TaskingAor string `json:"taskingAOR"`
	// Tasking geographical collection area.
	TaskingCollectionArea string `json:"taskingCollectionArea"`
	// Tasking desired collection requirements.
	TaskingCollectionRequirements []IsrCollectionListResponseTaskingTaskingCollectionRequirement `json:"taskingCollectionRequirements"`
	// Country code of the tasking. A Country may represent countries, multi-national
	// consortiums, and international organizations.
	TaskingCountry string `json:"taskingCountry"`
	// Tasking emphasis.
	TaskingEmphasis string `json:"taskingEmphasis"`
	// Joint Operations Area.
	TaskingJoa string `json:"taskingJoa"`
	// Tasking operation name.
	TaskingOperation string `json:"taskingOperation"`
	// Primary type of intelligence to be collected during the mission.
	TaskingPrimaryIntelDiscipline string `json:"taskingPrimaryIntelDiscipline"`
	// Sub category of primary intelligence to be collected.
	TaskingPrimarySubCategory string `json:"taskingPrimarySubCategory"`
	// Tasking Priority (1-n).
	TaskingPriority float64 `json:"taskingPriority"`
	// Region of the tasking.
	TaskingRegion string `json:"taskingRegion"`
	// Time of retasking, in ISO 8601 UTC format.
	TaskingRetaskTime time.Time `json:"taskingRetaskTime" format:"date-time"`
	// What is the primary objective (role) of this task.
	TaskingRole string `json:"taskingRole"`
	// Type of tasking intelligence to be collected second.
	TaskingSecondaryIntelDiscipline string `json:"taskingSecondaryIntelDiscipline"`
	// Mission sub category for secondary intelligence discipline to be collected.
	TaskingSecondarySubCategory string `json:"taskingSecondarySubCategory"`
	// WGS-84 latitude of the start position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	TaskingStartPointLat float64 `json:"taskingStartPointLat"`
	// WGS-84 longitude of the start position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	TaskingStartPointLong float64 `json:"taskingStartPointLong"`
	// Subregion of the tasking.
	TaskingSubRegion string `json:"taskingSubRegion"`
	// Military Base to transmit the dissemination of this data.
	TaskingSupportedUnit string `json:"taskingSupportedUnit"`
	// A synchronization matrix is used to organize the logistics synchronization
	// process during a mission.
	TaskingSyncMatrixBin string `json:"taskingSyncMatrixBin"`
	// Type of tasking.
	//
	// Any of "Deliberate", "Dynamic", "Training", "Transit".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		CollectionPeriods               respjson.Field
		CollectionType                  respjson.Field
		EightLine                       respjson.Field
		SpecialComGuidance              respjson.Field
		SroTrack                        respjson.Field
		TaskingAor                      respjson.Field
		TaskingCollectionArea           respjson.Field
		TaskingCollectionRequirements   respjson.Field
		TaskingCountry                  respjson.Field
		TaskingEmphasis                 respjson.Field
		TaskingJoa                      respjson.Field
		TaskingOperation                respjson.Field
		TaskingPrimaryIntelDiscipline   respjson.Field
		TaskingPrimarySubCategory       respjson.Field
		TaskingPriority                 respjson.Field
		TaskingRegion                   respjson.Field
		TaskingRetaskTime               respjson.Field
		TaskingRole                     respjson.Field
		TaskingSecondaryIntelDiscipline respjson.Field
		TaskingSecondarySubCategory     respjson.Field
		TaskingStartPointLat            respjson.Field
		TaskingStartPointLong           respjson.Field
		TaskingSubRegion                respjson.Field
		TaskingSupportedUnit            respjson.Field
		TaskingSyncMatrixBin            respjson.Field
		Type                            respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTasking) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionListResponseTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTaskingCollectionPeriods struct {
	// Actual start and stop for the collection.
	Actual  []IsrCollectionListResponseTaskingCollectionPeriodsActual `json:"actual"`
	Planned IsrCollectionListResponseTaskingCollectionPeriodsPlanned  `json:"planned"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Planned     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTaskingCollectionPeriods) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionListResponseTaskingCollectionPeriods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTaskingCollectionPeriodsActual struct {
	// Unique Identifier of actual collection period for historical archive.
	ID string `json:"id"`
	// Start time the collection actually occurred, in ISO 8601 UTC format.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time the collection actually occurred, in ISO 8601 UTC format.
	Stop time.Time `json:"stop" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Start       respjson.Field
		Stop        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTaskingCollectionPeriodsActual) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionListResponseTaskingCollectionPeriodsActual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTaskingCollectionPeriodsPlanned struct {
	// Additional start and stop for the collection.
	Additional []IsrCollectionListResponseTaskingCollectionPeriodsPlannedAdditional `json:"additional"`
	// Start time of collection, in ISO 8601 UTC format.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time of collection, in ISO 8601 UTC format.
	Stop time.Time `json:"stop" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Additional  respjson.Field
		Start       respjson.Field
		Stop        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTaskingCollectionPeriodsPlanned) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionListResponseTaskingCollectionPeriodsPlanned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTaskingCollectionPeriodsPlannedAdditional struct {
	// Unique Identifier of additional collection period.
	ID string `json:"id"`
	// Start time of collection, in ISO 8601 UTC format.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time of collection, in ISO 8601 UTC format.
	Stop time.Time `json:"stop" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Start       respjson.Field
		Stop        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTaskingCollectionPeriodsPlannedAdditional) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionListResponseTaskingCollectionPeriodsPlannedAdditional) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTaskingTaskingCollectionRequirement struct {
	// Collection Requirement Unique Identifier.
	ID string `json:"id"`
	// Country code of the collection requirement. A Country may represent countries,
	// multi-national consortiums, and international organizations.
	Country string `json:"country"`
	// Collection Requirement Unique Identifier.
	CridNumbers   string                                                                    `json:"cridNumbers"`
	CriticalTimes IsrCollectionListResponseTaskingTaskingCollectionRequirementCriticalTimes `json:"criticalTimes"`
	// Is this collection requirement an emphasized/critical requirement.
	Emphasized              bool                                                                                `json:"emphasized"`
	ExploitationRequirement IsrCollectionListResponseTaskingTaskingCollectionRequirementExploitationRequirement `json:"exploitationRequirement"`
	// Encryption hashing algorithm.
	Hash string `json:"hash"`
	// Primary type of intelligence to be collected for this requirement.
	IntelDiscipline string `json:"intelDiscipline"`
	// Is this collection request for the Prism system?.
	IsPrismCr bool `json:"isPrismCr"`
	// Human readable name for this operation.
	Operation string `json:"operation"`
	// 1-n priority for this collection requirement.
	Priority float64 `json:"priority"`
	// Reconnaissance Survey information the operator needs.
	ReconSurvey string `json:"reconSurvey"`
	// Record id.
	RecordID string `json:"recordId"`
	// Region of the collection requirement.
	Region string `json:"region"`
	// Sub category of primary intelligence to be collected for this requirement.
	Secondary bool `json:"secondary"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance string `json:"specialComGuidance"`
	// Start time for this requirement, should be within the mission time window.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time for this requirement, should be within the mission time window.
	Stop time.Time `json:"stop" format:"date-time"`
	// Subregion of the collection requirement.
	Subregion string `json:"subregion"`
	// The name of the military unit that this assigned collection requirement will
	// support.
	SupportedUnit string `json:"supportedUnit"`
	// Array of POI Id's for the targets being tasked.
	TargetList []string `json:"targetList"`
	// Type collection this requirement applies to.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Country                 respjson.Field
		CridNumbers             respjson.Field
		CriticalTimes           respjson.Field
		Emphasized              respjson.Field
		ExploitationRequirement respjson.Field
		Hash                    respjson.Field
		IntelDiscipline         respjson.Field
		IsPrismCr               respjson.Field
		Operation               respjson.Field
		Priority                respjson.Field
		ReconSurvey             respjson.Field
		RecordID                respjson.Field
		Region                  respjson.Field
		Secondary               respjson.Field
		SpecialComGuidance      respjson.Field
		Start                   respjson.Field
		Stop                    respjson.Field
		Subregion               respjson.Field
		SupportedUnit           respjson.Field
		TargetList              respjson.Field
		Type                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTaskingTaskingCollectionRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionListResponseTaskingTaskingCollectionRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTaskingTaskingCollectionRequirementCriticalTimes struct {
	// Critical start time to collect an image for this requirement.
	EarliestImagingTime time.Time `json:"earliestImagingTime,required" format:"date-time"`
	// Critical stop time to collect an image for this requirement.
	LatestImagingTime time.Time `json:"latestImagingTime,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EarliestImagingTime respjson.Field
		LatestImagingTime   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTaskingTaskingCollectionRequirementCriticalTimes) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionListResponseTaskingTaskingCollectionRequirementCriticalTimes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTaskingTaskingCollectionRequirementExploitationRequirement struct {
	// Exploitation requirement id.
	ID string `json:"id"`
	// Amplifying data for the exploitation requirement.
	Amplification string `json:"amplification"`
	// List of e-mails to disseminate collection verification information.
	Dissemination string `json:"dissemination"`
	// Essential Elements of Information.
	Eei string                                                                                 `json:"eei"`
	Poc IsrCollectionListResponseTaskingTaskingCollectionRequirementExploitationRequirementPoc `json:"poc"`
	// The reporting criteria of the collection requirement.
	ReportingCriteria string `json:"reportingCriteria"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Amplification     respjson.Field
		Dissemination     respjson.Field
		Eei               respjson.Field
		Poc               respjson.Field
		ReportingCriteria respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTaskingTaskingCollectionRequirementExploitationRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionListResponseTaskingTaskingCollectionRequirementExploitationRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTaskingTaskingCollectionRequirementExploitationRequirementPoc struct {
	// Unique identifier of the collection requirement POC.
	ID string `json:"id"`
	// Callsign of the POC.
	Callsign string `json:"callsign"`
	// Chat name of the POC.
	ChatName string `json:"chatName"`
	// Chat system the POC is accessing.
	ChatSystem string `json:"chatSystem"`
	// Email address of the POC.
	Email string `json:"email"`
	// Name of the POC.
	Name string `json:"name"`
	// Amplifying notes about the POC.
	Notes string `json:"notes"`
	// Phone number of the POC.
	Phone string `json:"phone"`
	// Radio Frequency the POC is on.
	RadioFrequency float64 `json:"radioFrequency"`
	// Unit the POC belongs to.
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Callsign       respjson.Field
		ChatName       respjson.Field
		ChatSystem     respjson.Field
		Email          respjson.Field
		Name           respjson.Field
		Notes          respjson.Field
		Phone          respjson.Field
		RadioFrequency respjson.Field
		Unit           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTaskingTaskingCollectionRequirementExploitationRequirementPoc) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionListResponseTaskingTaskingCollectionRequirementExploitationRequirementPoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListResponseTransit struct {
	// Transit Unique Identifier.
	ID string `json:"id"`
	// Military Base to transmit the dissemination of this data.
	Base string `json:"base"`
	// Length of mission in milliseconds.
	Duration float64 `json:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Base        respjson.Field
		Duration    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionListResponseTransit) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionListResponseTransit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ISR Collection data.
type IsrCollectionTupleResponse struct {
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
	DataMode IsrCollectionTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Mission desired collection requirements.
	CollectionRequirements []IsrCollectionTupleResponseCollectionRequirement `json:"collectionRequirements"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Version of the IDEX software the request came from for compatibility.
	IdexVersion int64 `json:"idexVersion"`
	// Designation of mission Area Of Responsibility.
	MissionAor string `json:"missionAOR"`
	// Mission geographical collection area.
	MissionCollectionArea string `json:"missionCollectionArea"`
	// Country code of the mission. A Country may represent countries, multi-national
	// consortiums, and international organizations.
	MissionCountry string `json:"missionCountry"`
	// Text version of what we are emphasizing in this mission.
	MissionEmphasis string `json:"missionEmphasis"`
	// Mission Identifier.
	MissionID string `json:"missionId"`
	// Joint Operations Area.
	MissionJoa string `json:"missionJoa"`
	// Mission operation name.
	MissionOperation string `json:"missionOperation"`
	// Primary type of intelligence to be collected during the mission.
	MissionPrimaryIntelDiscipline string `json:"missionPrimaryIntelDiscipline"`
	// Sub category of primary intelligence to be collected.
	MissionPrimarySubCategory string `json:"missionPrimarySubCategory"`
	// Mission Priority (1-n).
	MissionPriority int64 `json:"missionPriority"`
	// Region of the mission.
	MissionRegion string `json:"missionRegion"`
	// What is the primary objective(Role) of this mission.
	MissionRole string `json:"missionRole"`
	// Type of intelligence to be collected second.
	MissionSecondaryIntelDiscipline string `json:"missionSecondaryIntelDiscipline"`
	// Mission sub category for secondary intelligence discipline to be collected.
	MissionSecondarySubCategory string `json:"missionSecondarySubCategory"`
	// WGS-84 latitude of the start position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	MissionStartPointLat float64 `json:"missionStartPointLat"`
	// WGS-84 longitude of the start position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	MissionStartPointLong float64 `json:"missionStartPointLong"`
	// Subregion of the mission.
	MissionSubRegion string `json:"missionSubRegion"`
	// Name of the Supporting unit/Location that is performing this mission.
	MissionSupportedUnit string `json:"missionSupportedUnit"`
	// A synchronization matrix is used to organize the logistics synchronization
	// process during a mission.
	MissionSyncMatrixBin string `json:"missionSyncMatrixBin"`
	// Human readable Mission Name.
	Name string `json:"name"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Individual taskings to complete the mission.
	Taskings []IsrCollectionTupleResponseTasking `json:"taskings"`
	// Object for data dissemination.
	Transit []IsrCollectionTupleResponseTransit `json:"transit"`
	// Time the row was updated in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking           respjson.Field
		DataMode                        respjson.Field
		Source                          respjson.Field
		ID                              respjson.Field
		CollectionRequirements          respjson.Field
		CreatedAt                       respjson.Field
		CreatedBy                       respjson.Field
		IdexVersion                     respjson.Field
		MissionAor                      respjson.Field
		MissionCollectionArea           respjson.Field
		MissionCountry                  respjson.Field
		MissionEmphasis                 respjson.Field
		MissionID                       respjson.Field
		MissionJoa                      respjson.Field
		MissionOperation                respjson.Field
		MissionPrimaryIntelDiscipline   respjson.Field
		MissionPrimarySubCategory       respjson.Field
		MissionPriority                 respjson.Field
		MissionRegion                   respjson.Field
		MissionRole                     respjson.Field
		MissionSecondaryIntelDiscipline respjson.Field
		MissionSecondarySubCategory     respjson.Field
		MissionStartPointLat            respjson.Field
		MissionStartPointLong           respjson.Field
		MissionSubRegion                respjson.Field
		MissionSupportedUnit            respjson.Field
		MissionSyncMatrixBin            respjson.Field
		Name                            respjson.Field
		Origin                          respjson.Field
		OrigNetwork                     respjson.Field
		Taskings                        respjson.Field
		Transit                         respjson.Field
		UpdatedAt                       respjson.Field
		UpdatedBy                       respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionTupleResponse) UnmarshalJSON(data []byte) error {
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
type IsrCollectionTupleResponseDataMode string

const (
	IsrCollectionTupleResponseDataModeReal      IsrCollectionTupleResponseDataMode = "REAL"
	IsrCollectionTupleResponseDataModeTest      IsrCollectionTupleResponseDataMode = "TEST"
	IsrCollectionTupleResponseDataModeSimulated IsrCollectionTupleResponseDataMode = "SIMULATED"
	IsrCollectionTupleResponseDataModeExercise  IsrCollectionTupleResponseDataMode = "EXERCISE"
)

type IsrCollectionTupleResponseCollectionRequirement struct {
	// Collection Requirement Unique Identifier.
	ID string `json:"id"`
	// Country code of the collection requirement. A Country may represent countries,
	// multi-national consortiums, and international organizations.
	Country string `json:"country"`
	// Collection Requirement Unique Identifier.
	CridNumbers   string                                                       `json:"cridNumbers"`
	CriticalTimes IsrCollectionTupleResponseCollectionRequirementCriticalTimes `json:"criticalTimes"`
	// Is this collection requirement an emphasized/critical requirement.
	Emphasized              bool                                                                   `json:"emphasized"`
	ExploitationRequirement IsrCollectionTupleResponseCollectionRequirementExploitationRequirement `json:"exploitationRequirement"`
	// Encryption hashing algorithm.
	Hash string `json:"hash"`
	// Primary type of intelligence to be collected for this requirement.
	IntelDiscipline string `json:"intelDiscipline"`
	// Is this collection request for the Prism system?.
	IsPrismCr bool `json:"isPrismCr"`
	// Human readable name for this operation.
	Operation string `json:"operation"`
	// 1-n priority for this collection requirement.
	Priority float64 `json:"priority"`
	// Reconnaissance Survey information the operator needs.
	ReconSurvey string `json:"reconSurvey"`
	// Record id.
	RecordID string `json:"recordId"`
	// Region of the collection requirement.
	Region string `json:"region"`
	// Sub category of primary intelligence to be collected for this requirement.
	Secondary bool `json:"secondary"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance string `json:"specialComGuidance"`
	// Start time for this requirement, should be within the mission time window.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time for this requirement, should be within the mission time window.
	Stop time.Time `json:"stop" format:"date-time"`
	// Subregion of the collection requirement.
	Subregion string `json:"subregion"`
	// The name of the military unit that this assigned collection requirement will
	// support.
	SupportedUnit string `json:"supportedUnit"`
	// Array of POI Id's for the targets being tasked.
	TargetList []string `json:"targetList"`
	// Type collection this requirement applies to.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Country                 respjson.Field
		CridNumbers             respjson.Field
		CriticalTimes           respjson.Field
		Emphasized              respjson.Field
		ExploitationRequirement respjson.Field
		Hash                    respjson.Field
		IntelDiscipline         respjson.Field
		IsPrismCr               respjson.Field
		Operation               respjson.Field
		Priority                respjson.Field
		ReconSurvey             respjson.Field
		RecordID                respjson.Field
		Region                  respjson.Field
		Secondary               respjson.Field
		SpecialComGuidance      respjson.Field
		Start                   respjson.Field
		Stop                    respjson.Field
		Subregion               respjson.Field
		SupportedUnit           respjson.Field
		TargetList              respjson.Field
		Type                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseCollectionRequirement) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionTupleResponseCollectionRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseCollectionRequirementCriticalTimes struct {
	// Critical start time to collect an image for this requirement.
	EarliestImagingTime time.Time `json:"earliestImagingTime,required" format:"date-time"`
	// Critical stop time to collect an image for this requirement.
	LatestImagingTime time.Time `json:"latestImagingTime,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EarliestImagingTime respjson.Field
		LatestImagingTime   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseCollectionRequirementCriticalTimes) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionTupleResponseCollectionRequirementCriticalTimes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseCollectionRequirementExploitationRequirement struct {
	// Exploitation requirement id.
	ID string `json:"id"`
	// Amplifying data for the exploitation requirement.
	Amplification string `json:"amplification"`
	// List of e-mails to disseminate collection verification information.
	Dissemination string `json:"dissemination"`
	// Essential Elements of Information.
	Eei string                                                                    `json:"eei"`
	Poc IsrCollectionTupleResponseCollectionRequirementExploitationRequirementPoc `json:"poc"`
	// The reporting criteria of the collection requirement.
	ReportingCriteria string `json:"reportingCriteria"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Amplification     respjson.Field
		Dissemination     respjson.Field
		Eei               respjson.Field
		Poc               respjson.Field
		ReportingCriteria respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseCollectionRequirementExploitationRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionTupleResponseCollectionRequirementExploitationRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseCollectionRequirementExploitationRequirementPoc struct {
	// Unique identifier of the collection requirement POC.
	ID string `json:"id"`
	// Callsign of the POC.
	Callsign string `json:"callsign"`
	// Chat name of the POC.
	ChatName string `json:"chatName"`
	// Chat system the POC is accessing.
	ChatSystem string `json:"chatSystem"`
	// Email address of the POC.
	Email string `json:"email"`
	// Name of the POC.
	Name string `json:"name"`
	// Amplifying notes about the POC.
	Notes string `json:"notes"`
	// Phone number of the POC.
	Phone string `json:"phone"`
	// Radio Frequency the POC is on.
	RadioFrequency float64 `json:"radioFrequency"`
	// Unit the POC belongs to.
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Callsign       respjson.Field
		ChatName       respjson.Field
		ChatSystem     respjson.Field
		Email          respjson.Field
		Name           respjson.Field
		Notes          respjson.Field
		Phone          respjson.Field
		RadioFrequency respjson.Field
		Unit           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseCollectionRequirementExploitationRequirementPoc) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionTupleResponseCollectionRequirementExploitationRequirementPoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTasking struct {
	// Tasking Unique Identifier.
	ID                string                                             `json:"id"`
	CollectionPeriods IsrCollectionTupleResponseTaskingCollectionPeriods `json:"collectionPeriods"`
	// Type of collection tasked.
	//
	// Any of "Simultaneous", "Sequential", "Operationally", "Driven", "Priority",
	// "Order".
	CollectionType string `json:"collectionType"`
	// Eight line.
	EightLine string `json:"eightLine"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance string `json:"specialComGuidance"`
	// Value of the Sensitive Reconnaissance Operations Track.
	SroTrack string `json:"sroTrack"`
	// Human readable definition of this taskings Area Of Responsibility.
	TaskingAor string `json:"taskingAOR"`
	// Tasking geographical collection area.
	TaskingCollectionArea string `json:"taskingCollectionArea"`
	// Tasking desired collection requirements.
	TaskingCollectionRequirements []IsrCollectionTupleResponseTaskingTaskingCollectionRequirement `json:"taskingCollectionRequirements"`
	// Country code of the tasking. A Country may represent countries, multi-national
	// consortiums, and international organizations.
	TaskingCountry string `json:"taskingCountry"`
	// Tasking emphasis.
	TaskingEmphasis string `json:"taskingEmphasis"`
	// Joint Operations Area.
	TaskingJoa string `json:"taskingJoa"`
	// Tasking operation name.
	TaskingOperation string `json:"taskingOperation"`
	// Primary type of intelligence to be collected during the mission.
	TaskingPrimaryIntelDiscipline string `json:"taskingPrimaryIntelDiscipline"`
	// Sub category of primary intelligence to be collected.
	TaskingPrimarySubCategory string `json:"taskingPrimarySubCategory"`
	// Tasking Priority (1-n).
	TaskingPriority float64 `json:"taskingPriority"`
	// Region of the tasking.
	TaskingRegion string `json:"taskingRegion"`
	// Time of retasking, in ISO 8601 UTC format.
	TaskingRetaskTime time.Time `json:"taskingRetaskTime" format:"date-time"`
	// What is the primary objective (role) of this task.
	TaskingRole string `json:"taskingRole"`
	// Type of tasking intelligence to be collected second.
	TaskingSecondaryIntelDiscipline string `json:"taskingSecondaryIntelDiscipline"`
	// Mission sub category for secondary intelligence discipline to be collected.
	TaskingSecondarySubCategory string `json:"taskingSecondarySubCategory"`
	// WGS-84 latitude of the start position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	TaskingStartPointLat float64 `json:"taskingStartPointLat"`
	// WGS-84 longitude of the start position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	TaskingStartPointLong float64 `json:"taskingStartPointLong"`
	// Subregion of the tasking.
	TaskingSubRegion string `json:"taskingSubRegion"`
	// Military Base to transmit the dissemination of this data.
	TaskingSupportedUnit string `json:"taskingSupportedUnit"`
	// A synchronization matrix is used to organize the logistics synchronization
	// process during a mission.
	TaskingSyncMatrixBin string `json:"taskingSyncMatrixBin"`
	// Type of tasking.
	//
	// Any of "Deliberate", "Dynamic", "Training", "Transit".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		CollectionPeriods               respjson.Field
		CollectionType                  respjson.Field
		EightLine                       respjson.Field
		SpecialComGuidance              respjson.Field
		SroTrack                        respjson.Field
		TaskingAor                      respjson.Field
		TaskingCollectionArea           respjson.Field
		TaskingCollectionRequirements   respjson.Field
		TaskingCountry                  respjson.Field
		TaskingEmphasis                 respjson.Field
		TaskingJoa                      respjson.Field
		TaskingOperation                respjson.Field
		TaskingPrimaryIntelDiscipline   respjson.Field
		TaskingPrimarySubCategory       respjson.Field
		TaskingPriority                 respjson.Field
		TaskingRegion                   respjson.Field
		TaskingRetaskTime               respjson.Field
		TaskingRole                     respjson.Field
		TaskingSecondaryIntelDiscipline respjson.Field
		TaskingSecondarySubCategory     respjson.Field
		TaskingStartPointLat            respjson.Field
		TaskingStartPointLong           respjson.Field
		TaskingSubRegion                respjson.Field
		TaskingSupportedUnit            respjson.Field
		TaskingSyncMatrixBin            respjson.Field
		Type                            respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTasking) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionTupleResponseTasking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTaskingCollectionPeriods struct {
	// Actual start and stop for the collection.
	Actual  []IsrCollectionTupleResponseTaskingCollectionPeriodsActual `json:"actual"`
	Planned IsrCollectionTupleResponseTaskingCollectionPeriodsPlanned  `json:"planned"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Planned     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTaskingCollectionPeriods) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionTupleResponseTaskingCollectionPeriods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTaskingCollectionPeriodsActual struct {
	// Unique Identifier of actual collection period for historical archive.
	ID string `json:"id"`
	// Start time the collection actually occurred, in ISO 8601 UTC format.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time the collection actually occurred, in ISO 8601 UTC format.
	Stop time.Time `json:"stop" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Start       respjson.Field
		Stop        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTaskingCollectionPeriodsActual) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionTupleResponseTaskingCollectionPeriodsActual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTaskingCollectionPeriodsPlanned struct {
	// Additional start and stop for the collection.
	Additional []IsrCollectionTupleResponseTaskingCollectionPeriodsPlannedAdditional `json:"additional"`
	// Start time of collection, in ISO 8601 UTC format.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time of collection, in ISO 8601 UTC format.
	Stop time.Time `json:"stop" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Additional  respjson.Field
		Start       respjson.Field
		Stop        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTaskingCollectionPeriodsPlanned) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionTupleResponseTaskingCollectionPeriodsPlanned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTaskingCollectionPeriodsPlannedAdditional struct {
	// Unique Identifier of additional collection period.
	ID string `json:"id"`
	// Start time of collection, in ISO 8601 UTC format.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time of collection, in ISO 8601 UTC format.
	Stop time.Time `json:"stop" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Start       respjson.Field
		Stop        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTaskingCollectionPeriodsPlannedAdditional) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionTupleResponseTaskingCollectionPeriodsPlannedAdditional) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTaskingTaskingCollectionRequirement struct {
	// Collection Requirement Unique Identifier.
	ID string `json:"id"`
	// Country code of the collection requirement. A Country may represent countries,
	// multi-national consortiums, and international organizations.
	Country string `json:"country"`
	// Collection Requirement Unique Identifier.
	CridNumbers   string                                                                     `json:"cridNumbers"`
	CriticalTimes IsrCollectionTupleResponseTaskingTaskingCollectionRequirementCriticalTimes `json:"criticalTimes"`
	// Is this collection requirement an emphasized/critical requirement.
	Emphasized              bool                                                                                 `json:"emphasized"`
	ExploitationRequirement IsrCollectionTupleResponseTaskingTaskingCollectionRequirementExploitationRequirement `json:"exploitationRequirement"`
	// Encryption hashing algorithm.
	Hash string `json:"hash"`
	// Primary type of intelligence to be collected for this requirement.
	IntelDiscipline string `json:"intelDiscipline"`
	// Is this collection request for the Prism system?.
	IsPrismCr bool `json:"isPrismCr"`
	// Human readable name for this operation.
	Operation string `json:"operation"`
	// 1-n priority for this collection requirement.
	Priority float64 `json:"priority"`
	// Reconnaissance Survey information the operator needs.
	ReconSurvey string `json:"reconSurvey"`
	// Record id.
	RecordID string `json:"recordId"`
	// Region of the collection requirement.
	Region string `json:"region"`
	// Sub category of primary intelligence to be collected for this requirement.
	Secondary bool `json:"secondary"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance string `json:"specialComGuidance"`
	// Start time for this requirement, should be within the mission time window.
	Start time.Time `json:"start" format:"date-time"`
	// Stop time for this requirement, should be within the mission time window.
	Stop time.Time `json:"stop" format:"date-time"`
	// Subregion of the collection requirement.
	Subregion string `json:"subregion"`
	// The name of the military unit that this assigned collection requirement will
	// support.
	SupportedUnit string `json:"supportedUnit"`
	// Array of POI Id's for the targets being tasked.
	TargetList []string `json:"targetList"`
	// Type collection this requirement applies to.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Country                 respjson.Field
		CridNumbers             respjson.Field
		CriticalTimes           respjson.Field
		Emphasized              respjson.Field
		ExploitationRequirement respjson.Field
		Hash                    respjson.Field
		IntelDiscipline         respjson.Field
		IsPrismCr               respjson.Field
		Operation               respjson.Field
		Priority                respjson.Field
		ReconSurvey             respjson.Field
		RecordID                respjson.Field
		Region                  respjson.Field
		Secondary               respjson.Field
		SpecialComGuidance      respjson.Field
		Start                   respjson.Field
		Stop                    respjson.Field
		Subregion               respjson.Field
		SupportedUnit           respjson.Field
		TargetList              respjson.Field
		Type                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTaskingTaskingCollectionRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionTupleResponseTaskingTaskingCollectionRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTaskingTaskingCollectionRequirementCriticalTimes struct {
	// Critical start time to collect an image for this requirement.
	EarliestImagingTime time.Time `json:"earliestImagingTime,required" format:"date-time"`
	// Critical stop time to collect an image for this requirement.
	LatestImagingTime time.Time `json:"latestImagingTime,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EarliestImagingTime respjson.Field
		LatestImagingTime   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTaskingTaskingCollectionRequirementCriticalTimes) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionTupleResponseTaskingTaskingCollectionRequirementCriticalTimes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTaskingTaskingCollectionRequirementExploitationRequirement struct {
	// Exploitation requirement id.
	ID string `json:"id"`
	// Amplifying data for the exploitation requirement.
	Amplification string `json:"amplification"`
	// List of e-mails to disseminate collection verification information.
	Dissemination string `json:"dissemination"`
	// Essential Elements of Information.
	Eei string                                                                                  `json:"eei"`
	Poc IsrCollectionTupleResponseTaskingTaskingCollectionRequirementExploitationRequirementPoc `json:"poc"`
	// The reporting criteria of the collection requirement.
	ReportingCriteria string `json:"reportingCriteria"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Amplification     respjson.Field
		Dissemination     respjson.Field
		Eei               respjson.Field
		Poc               respjson.Field
		ReportingCriteria respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTaskingTaskingCollectionRequirementExploitationRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionTupleResponseTaskingTaskingCollectionRequirementExploitationRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTaskingTaskingCollectionRequirementExploitationRequirementPoc struct {
	// Unique identifier of the collection requirement POC.
	ID string `json:"id"`
	// Callsign of the POC.
	Callsign string `json:"callsign"`
	// Chat name of the POC.
	ChatName string `json:"chatName"`
	// Chat system the POC is accessing.
	ChatSystem string `json:"chatSystem"`
	// Email address of the POC.
	Email string `json:"email"`
	// Name of the POC.
	Name string `json:"name"`
	// Amplifying notes about the POC.
	Notes string `json:"notes"`
	// Phone number of the POC.
	Phone string `json:"phone"`
	// Radio Frequency the POC is on.
	RadioFrequency float64 `json:"radioFrequency"`
	// Unit the POC belongs to.
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Callsign       respjson.Field
		ChatName       respjson.Field
		ChatSystem     respjson.Field
		Email          respjson.Field
		Name           respjson.Field
		Notes          respjson.Field
		Phone          respjson.Field
		RadioFrequency respjson.Field
		Unit           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTaskingTaskingCollectionRequirementExploitationRequirementPoc) RawJSON() string {
	return r.JSON.raw
}
func (r *IsrCollectionTupleResponseTaskingTaskingCollectionRequirementExploitationRequirementPoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionTupleResponseTransit struct {
	// Transit Unique Identifier.
	ID string `json:"id"`
	// Military Base to transmit the dissemination of this data.
	Base string `json:"base"`
	// Length of mission in milliseconds.
	Duration float64 `json:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Base        respjson.Field
		Duration    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsrCollectionTupleResponseTransit) RawJSON() string { return r.JSON.raw }
func (r *IsrCollectionTupleResponseTransit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsrCollectionListParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IsrCollectionListParams]'s query parameters as
// `url.Values`.
func (r IsrCollectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IsrCollectionCountParams struct {
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IsrCollectionCountParams]'s query parameters as
// `url.Values`.
func (r IsrCollectionCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IsrCollectionNewBulkParams struct {
	Body []IsrCollectionNewBulkParamsBody
	paramObj
}

func (r IsrCollectionNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// ISR Collection data.
//
// The properties ClassificationMarking, DataMode, Source are required.
type IsrCollectionNewBulkParamsBody struct {
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
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Version of the IDEX software the request came from for compatibility.
	IdexVersion param.Opt[int64] `json:"idexVersion,omitzero"`
	// Designation of mission Area Of Responsibility.
	MissionAor param.Opt[string] `json:"missionAOR,omitzero"`
	// Mission geographical collection area.
	MissionCollectionArea param.Opt[string] `json:"missionCollectionArea,omitzero"`
	// Country code of the mission. A Country may represent countries, multi-national
	// consortiums, and international organizations.
	MissionCountry param.Opt[string] `json:"missionCountry,omitzero"`
	// Text version of what we are emphasizing in this mission.
	MissionEmphasis param.Opt[string] `json:"missionEmphasis,omitzero"`
	// Mission Identifier.
	MissionID param.Opt[string] `json:"missionId,omitzero"`
	// Joint Operations Area.
	MissionJoa param.Opt[string] `json:"missionJoa,omitzero"`
	// Mission operation name.
	MissionOperation param.Opt[string] `json:"missionOperation,omitzero"`
	// Primary type of intelligence to be collected during the mission.
	MissionPrimaryIntelDiscipline param.Opt[string] `json:"missionPrimaryIntelDiscipline,omitzero"`
	// Sub category of primary intelligence to be collected.
	MissionPrimarySubCategory param.Opt[string] `json:"missionPrimarySubCategory,omitzero"`
	// Mission Priority (1-n).
	MissionPriority param.Opt[int64] `json:"missionPriority,omitzero"`
	// Region of the mission.
	MissionRegion param.Opt[string] `json:"missionRegion,omitzero"`
	// What is the primary objective(Role) of this mission.
	MissionRole param.Opt[string] `json:"missionRole,omitzero"`
	// Type of intelligence to be collected second.
	MissionSecondaryIntelDiscipline param.Opt[string] `json:"missionSecondaryIntelDiscipline,omitzero"`
	// Mission sub category for secondary intelligence discipline to be collected.
	MissionSecondarySubCategory param.Opt[string] `json:"missionSecondarySubCategory,omitzero"`
	// WGS-84 latitude of the start position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	MissionStartPointLat param.Opt[float64] `json:"missionStartPointLat,omitzero"`
	// WGS-84 longitude of the start position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	MissionStartPointLong param.Opt[float64] `json:"missionStartPointLong,omitzero"`
	// Subregion of the mission.
	MissionSubRegion param.Opt[string] `json:"missionSubRegion,omitzero"`
	// Name of the Supporting unit/Location that is performing this mission.
	MissionSupportedUnit param.Opt[string] `json:"missionSupportedUnit,omitzero"`
	// A synchronization matrix is used to organize the logistics synchronization
	// process during a mission.
	MissionSyncMatrixBin param.Opt[string] `json:"missionSyncMatrixBin,omitzero"`
	// Human readable Mission Name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Time the row was updated in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Mission desired collection requirements.
	CollectionRequirements []IsrCollectionNewBulkParamsBodyCollectionRequirement `json:"collectionRequirements,omitzero"`
	// Individual taskings to complete the mission.
	Taskings []IsrCollectionNewBulkParamsBodyTasking `json:"taskings,omitzero"`
	// Object for data dissemination.
	Transit []IsrCollectionNewBulkParamsBodyTransit `json:"transit,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[IsrCollectionNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type IsrCollectionNewBulkParamsBodyCollectionRequirement struct {
	// Collection Requirement Unique Identifier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Country code of the collection requirement. A Country may represent countries,
	// multi-national consortiums, and international organizations.
	Country param.Opt[string] `json:"country,omitzero"`
	// Collection Requirement Unique Identifier.
	CridNumbers param.Opt[string] `json:"cridNumbers,omitzero"`
	// Is this collection requirement an emphasized/critical requirement.
	Emphasized param.Opt[bool] `json:"emphasized,omitzero"`
	// Encryption hashing algorithm.
	Hash param.Opt[string] `json:"hash,omitzero"`
	// Primary type of intelligence to be collected for this requirement.
	IntelDiscipline param.Opt[string] `json:"intelDiscipline,omitzero"`
	// Is this collection request for the Prism system?.
	IsPrismCr param.Opt[bool] `json:"isPrismCr,omitzero"`
	// Human readable name for this operation.
	Operation param.Opt[string] `json:"operation,omitzero"`
	// 1-n priority for this collection requirement.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Reconnaissance Survey information the operator needs.
	ReconSurvey param.Opt[string] `json:"reconSurvey,omitzero"`
	// Record id.
	RecordID param.Opt[string] `json:"recordId,omitzero"`
	// Region of the collection requirement.
	Region param.Opt[string] `json:"region,omitzero"`
	// Sub category of primary intelligence to be collected for this requirement.
	Secondary param.Opt[bool] `json:"secondary,omitzero"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance param.Opt[string] `json:"specialComGuidance,omitzero"`
	// Start time for this requirement, should be within the mission time window.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time for this requirement, should be within the mission time window.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	// Subregion of the collection requirement.
	Subregion param.Opt[string] `json:"subregion,omitzero"`
	// The name of the military unit that this assigned collection requirement will
	// support.
	SupportedUnit param.Opt[string] `json:"supportedUnit,omitzero"`
	// Type collection this requirement applies to.
	Type                    param.Opt[string]                                                          `json:"type,omitzero"`
	CriticalTimes           IsrCollectionNewBulkParamsBodyCollectionRequirementCriticalTimes           `json:"criticalTimes,omitzero"`
	ExploitationRequirement IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirement `json:"exploitationRequirement,omitzero"`
	// Array of POI Id's for the targets being tasked.
	TargetList []string `json:"targetList,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyCollectionRequirement) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyCollectionRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties EarliestImagingTime, LatestImagingTime are required.
type IsrCollectionNewBulkParamsBodyCollectionRequirementCriticalTimes struct {
	// Critical start time to collect an image for this requirement.
	EarliestImagingTime time.Time `json:"earliestImagingTime,required" format:"date-time"`
	// Critical stop time to collect an image for this requirement.
	LatestImagingTime time.Time `json:"latestImagingTime,required" format:"date-time"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyCollectionRequirementCriticalTimes) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyCollectionRequirementCriticalTimes
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirement struct {
	// Exploitation requirement id.
	ID param.Opt[string] `json:"id,omitzero"`
	// Amplifying data for the exploitation requirement.
	Amplification param.Opt[string] `json:"amplification,omitzero"`
	// List of e-mails to disseminate collection verification information.
	Dissemination param.Opt[string] `json:"dissemination,omitzero"`
	// Essential Elements of Information.
	Eei param.Opt[string] `json:"eei,omitzero"`
	// The reporting criteria of the collection requirement.
	ReportingCriteria param.Opt[string]                                                             `json:"reportingCriteria,omitzero"`
	Poc               IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirementPoc `json:"poc,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirement) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirementPoc struct {
	// Unique identifier of the collection requirement POC.
	ID param.Opt[string] `json:"id,omitzero"`
	// Callsign of the POC.
	Callsign param.Opt[string] `json:"callsign,omitzero"`
	// Chat name of the POC.
	ChatName param.Opt[string] `json:"chatName,omitzero"`
	// Chat system the POC is accessing.
	ChatSystem param.Opt[string] `json:"chatSystem,omitzero"`
	// Email address of the POC.
	Email param.Opt[string] `json:"email,omitzero"`
	// Name of the POC.
	Name param.Opt[string] `json:"name,omitzero"`
	// Amplifying notes about the POC.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Phone number of the POC.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Radio Frequency the POC is on.
	RadioFrequency param.Opt[float64] `json:"radioFrequency,omitzero"`
	// Unit the POC belongs to.
	Unit param.Opt[string] `json:"unit,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirementPoc) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirementPoc
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyTasking struct {
	// Tasking Unique Identifier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Eight line.
	EightLine param.Opt[string] `json:"eightLine,omitzero"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance param.Opt[string] `json:"specialComGuidance,omitzero"`
	// Value of the Sensitive Reconnaissance Operations Track.
	SroTrack param.Opt[string] `json:"sroTrack,omitzero"`
	// Human readable definition of this taskings Area Of Responsibility.
	TaskingAor param.Opt[string] `json:"taskingAOR,omitzero"`
	// Tasking geographical collection area.
	TaskingCollectionArea param.Opt[string] `json:"taskingCollectionArea,omitzero"`
	// Country code of the tasking. A Country may represent countries, multi-national
	// consortiums, and international organizations.
	TaskingCountry param.Opt[string] `json:"taskingCountry,omitzero"`
	// Tasking emphasis.
	TaskingEmphasis param.Opt[string] `json:"taskingEmphasis,omitzero"`
	// Joint Operations Area.
	TaskingJoa param.Opt[string] `json:"taskingJoa,omitzero"`
	// Tasking operation name.
	TaskingOperation param.Opt[string] `json:"taskingOperation,omitzero"`
	// Primary type of intelligence to be collected during the mission.
	TaskingPrimaryIntelDiscipline param.Opt[string] `json:"taskingPrimaryIntelDiscipline,omitzero"`
	// Sub category of primary intelligence to be collected.
	TaskingPrimarySubCategory param.Opt[string] `json:"taskingPrimarySubCategory,omitzero"`
	// Tasking Priority (1-n).
	TaskingPriority param.Opt[float64] `json:"taskingPriority,omitzero"`
	// Region of the tasking.
	TaskingRegion param.Opt[string] `json:"taskingRegion,omitzero"`
	// Time of retasking, in ISO 8601 UTC format.
	TaskingRetaskTime param.Opt[time.Time] `json:"taskingRetaskTime,omitzero" format:"date-time"`
	// What is the primary objective (role) of this task.
	TaskingRole param.Opt[string] `json:"taskingRole,omitzero"`
	// Type of tasking intelligence to be collected second.
	TaskingSecondaryIntelDiscipline param.Opt[string] `json:"taskingSecondaryIntelDiscipline,omitzero"`
	// Mission sub category for secondary intelligence discipline to be collected.
	TaskingSecondarySubCategory param.Opt[string] `json:"taskingSecondarySubCategory,omitzero"`
	// WGS-84 latitude of the start position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	TaskingStartPointLat param.Opt[float64] `json:"taskingStartPointLat,omitzero"`
	// WGS-84 longitude of the start position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	TaskingStartPointLong param.Opt[float64] `json:"taskingStartPointLong,omitzero"`
	// Subregion of the tasking.
	TaskingSubRegion param.Opt[string] `json:"taskingSubRegion,omitzero"`
	// Military Base to transmit the dissemination of this data.
	TaskingSupportedUnit param.Opt[string] `json:"taskingSupportedUnit,omitzero"`
	// A synchronization matrix is used to organize the logistics synchronization
	// process during a mission.
	TaskingSyncMatrixBin param.Opt[string]                                      `json:"taskingSyncMatrixBin,omitzero"`
	CollectionPeriods    IsrCollectionNewBulkParamsBodyTaskingCollectionPeriods `json:"collectionPeriods,omitzero"`
	// Type of collection tasked.
	//
	// Any of "Simultaneous", "Sequential", "Operationally", "Driven", "Priority",
	// "Order".
	CollectionType string `json:"collectionType,omitzero"`
	// Tasking desired collection requirements.
	TaskingCollectionRequirements []IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirement `json:"taskingCollectionRequirements,omitzero"`
	// Type of tasking.
	//
	// Any of "Deliberate", "Dynamic", "Training", "Transit".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTasking) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTasking
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[IsrCollectionNewBulkParamsBodyTasking](
		"CollectionType", false, "Simultaneous", "Sequential", "Operationally", "Driven", "Priority", "Order",
	)
	apijson.RegisterFieldValidator[IsrCollectionNewBulkParamsBodyTasking](
		"Type", false, "Deliberate", "Dynamic", "Training", "Transit",
	)
}

type IsrCollectionNewBulkParamsBodyTaskingCollectionPeriods struct {
	// Actual start and stop for the collection.
	Actual  []IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsActual `json:"actual,omitzero"`
	Planned IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlanned  `json:"planned,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTaskingCollectionPeriods) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTaskingCollectionPeriods
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsActual struct {
	// Unique Identifier of actual collection period for historical archive.
	ID param.Opt[string] `json:"id,omitzero"`
	// Start time the collection actually occurred, in ISO 8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time the collection actually occurred, in ISO 8601 UTC format.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsActual) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsActual
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlanned struct {
	// Start time of collection, in ISO 8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time of collection, in ISO 8601 UTC format.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	// Additional start and stop for the collection.
	Additional []IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlannedAdditional `json:"additional,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlanned) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlanned
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlannedAdditional struct {
	// Unique Identifier of additional collection period.
	ID param.Opt[string] `json:"id,omitzero"`
	// Start time of collection, in ISO 8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time of collection, in ISO 8601 UTC format.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlannedAdditional) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlannedAdditional
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirement struct {
	// Collection Requirement Unique Identifier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Country code of the collection requirement. A Country may represent countries,
	// multi-national consortiums, and international organizations.
	Country param.Opt[string] `json:"country,omitzero"`
	// Collection Requirement Unique Identifier.
	CridNumbers param.Opt[string] `json:"cridNumbers,omitzero"`
	// Is this collection requirement an emphasized/critical requirement.
	Emphasized param.Opt[bool] `json:"emphasized,omitzero"`
	// Encryption hashing algorithm.
	Hash param.Opt[string] `json:"hash,omitzero"`
	// Primary type of intelligence to be collected for this requirement.
	IntelDiscipline param.Opt[string] `json:"intelDiscipline,omitzero"`
	// Is this collection request for the Prism system?.
	IsPrismCr param.Opt[bool] `json:"isPrismCr,omitzero"`
	// Human readable name for this operation.
	Operation param.Opt[string] `json:"operation,omitzero"`
	// 1-n priority for this collection requirement.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Reconnaissance Survey information the operator needs.
	ReconSurvey param.Opt[string] `json:"reconSurvey,omitzero"`
	// Record id.
	RecordID param.Opt[string] `json:"recordId,omitzero"`
	// Region of the collection requirement.
	Region param.Opt[string] `json:"region,omitzero"`
	// Sub category of primary intelligence to be collected for this requirement.
	Secondary param.Opt[bool] `json:"secondary,omitzero"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance param.Opt[string] `json:"specialComGuidance,omitzero"`
	// Start time for this requirement, should be within the mission time window.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time for this requirement, should be within the mission time window.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	// Subregion of the collection requirement.
	Subregion param.Opt[string] `json:"subregion,omitzero"`
	// The name of the military unit that this assigned collection requirement will
	// support.
	SupportedUnit param.Opt[string] `json:"supportedUnit,omitzero"`
	// Type collection this requirement applies to.
	Type                    param.Opt[string]                                                                        `json:"type,omitzero"`
	CriticalTimes           IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementCriticalTimes           `json:"criticalTimes,omitzero"`
	ExploitationRequirement IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement `json:"exploitationRequirement,omitzero"`
	// Array of POI Id's for the targets being tasked.
	TargetList []string `json:"targetList,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirement) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties EarliestImagingTime, LatestImagingTime are required.
type IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementCriticalTimes struct {
	// Critical start time to collect an image for this requirement.
	EarliestImagingTime time.Time `json:"earliestImagingTime,required" format:"date-time"`
	// Critical stop time to collect an image for this requirement.
	LatestImagingTime time.Time `json:"latestImagingTime,required" format:"date-time"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementCriticalTimes) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementCriticalTimes
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement struct {
	// Exploitation requirement id.
	ID param.Opt[string] `json:"id,omitzero"`
	// Amplifying data for the exploitation requirement.
	Amplification param.Opt[string] `json:"amplification,omitzero"`
	// List of e-mails to disseminate collection verification information.
	Dissemination param.Opt[string] `json:"dissemination,omitzero"`
	// Essential Elements of Information.
	Eei param.Opt[string] `json:"eei,omitzero"`
	// The reporting criteria of the collection requirement.
	ReportingCriteria param.Opt[string]                                                                           `json:"reportingCriteria,omitzero"`
	Poc               IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc `json:"poc,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc struct {
	// Unique identifier of the collection requirement POC.
	ID param.Opt[string] `json:"id,omitzero"`
	// Callsign of the POC.
	Callsign param.Opt[string] `json:"callsign,omitzero"`
	// Chat name of the POC.
	ChatName param.Opt[string] `json:"chatName,omitzero"`
	// Chat system the POC is accessing.
	ChatSystem param.Opt[string] `json:"chatSystem,omitzero"`
	// Email address of the POC.
	Email param.Opt[string] `json:"email,omitzero"`
	// Name of the POC.
	Name param.Opt[string] `json:"name,omitzero"`
	// Amplifying notes about the POC.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Phone number of the POC.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Radio Frequency the POC is on.
	RadioFrequency param.Opt[float64] `json:"radioFrequency,omitzero"`
	// Unit the POC belongs to.
	Unit param.Opt[string] `json:"unit,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionNewBulkParamsBodyTransit struct {
	// Transit Unique Identifier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Military Base to transmit the dissemination of this data.
	Base param.Opt[string] `json:"base,omitzero"`
	// Length of mission in milliseconds.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	paramObj
}

func (r IsrCollectionNewBulkParamsBodyTransit) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionNewBulkParamsBodyTransit
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time the row was created in the database, auto-populated by the system.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	CreatedAt   time.Time        `query:"createdAt,required" format:"date" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IsrCollectionTupleParams]'s query parameters as
// `url.Values`.
func (r IsrCollectionTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IsrCollectionUnvalidatedPublishParams struct {
	Body []IsrCollectionUnvalidatedPublishParamsBody
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// ISR Collection data.
//
// The properties ClassificationMarking, DataMode, Source are required.
type IsrCollectionUnvalidatedPublishParamsBody struct {
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
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Version of the IDEX software the request came from for compatibility.
	IdexVersion param.Opt[int64] `json:"idexVersion,omitzero"`
	// Designation of mission Area Of Responsibility.
	MissionAor param.Opt[string] `json:"missionAOR,omitzero"`
	// Mission geographical collection area.
	MissionCollectionArea param.Opt[string] `json:"missionCollectionArea,omitzero"`
	// Country code of the mission. A Country may represent countries, multi-national
	// consortiums, and international organizations.
	MissionCountry param.Opt[string] `json:"missionCountry,omitzero"`
	// Text version of what we are emphasizing in this mission.
	MissionEmphasis param.Opt[string] `json:"missionEmphasis,omitzero"`
	// Mission Identifier.
	MissionID param.Opt[string] `json:"missionId,omitzero"`
	// Joint Operations Area.
	MissionJoa param.Opt[string] `json:"missionJoa,omitzero"`
	// Mission operation name.
	MissionOperation param.Opt[string] `json:"missionOperation,omitzero"`
	// Primary type of intelligence to be collected during the mission.
	MissionPrimaryIntelDiscipline param.Opt[string] `json:"missionPrimaryIntelDiscipline,omitzero"`
	// Sub category of primary intelligence to be collected.
	MissionPrimarySubCategory param.Opt[string] `json:"missionPrimarySubCategory,omitzero"`
	// Mission Priority (1-n).
	MissionPriority param.Opt[int64] `json:"missionPriority,omitzero"`
	// Region of the mission.
	MissionRegion param.Opt[string] `json:"missionRegion,omitzero"`
	// What is the primary objective(Role) of this mission.
	MissionRole param.Opt[string] `json:"missionRole,omitzero"`
	// Type of intelligence to be collected second.
	MissionSecondaryIntelDiscipline param.Opt[string] `json:"missionSecondaryIntelDiscipline,omitzero"`
	// Mission sub category for secondary intelligence discipline to be collected.
	MissionSecondarySubCategory param.Opt[string] `json:"missionSecondarySubCategory,omitzero"`
	// WGS-84 latitude of the start position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	MissionStartPointLat param.Opt[float64] `json:"missionStartPointLat,omitzero"`
	// WGS-84 longitude of the start position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	MissionStartPointLong param.Opt[float64] `json:"missionStartPointLong,omitzero"`
	// Subregion of the mission.
	MissionSubRegion param.Opt[string] `json:"missionSubRegion,omitzero"`
	// Name of the Supporting unit/Location that is performing this mission.
	MissionSupportedUnit param.Opt[string] `json:"missionSupportedUnit,omitzero"`
	// A synchronization matrix is used to organize the logistics synchronization
	// process during a mission.
	MissionSyncMatrixBin param.Opt[string] `json:"missionSyncMatrixBin,omitzero"`
	// Human readable Mission Name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Time the row was updated in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Mission desired collection requirements.
	CollectionRequirements []IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirement `json:"collectionRequirements,omitzero"`
	// Individual taskings to complete the mission.
	Taskings []IsrCollectionUnvalidatedPublishParamsBodyTasking `json:"taskings,omitzero"`
	// Object for data dissemination.
	Transit []IsrCollectionUnvalidatedPublishParamsBodyTransit `json:"transit,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[IsrCollectionUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirement struct {
	// Collection Requirement Unique Identifier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Country code of the collection requirement. A Country may represent countries,
	// multi-national consortiums, and international organizations.
	Country param.Opt[string] `json:"country,omitzero"`
	// Collection Requirement Unique Identifier.
	CridNumbers param.Opt[string] `json:"cridNumbers,omitzero"`
	// Is this collection requirement an emphasized/critical requirement.
	Emphasized param.Opt[bool] `json:"emphasized,omitzero"`
	// Encryption hashing algorithm.
	Hash param.Opt[string] `json:"hash,omitzero"`
	// Primary type of intelligence to be collected for this requirement.
	IntelDiscipline param.Opt[string] `json:"intelDiscipline,omitzero"`
	// Is this collection request for the Prism system?.
	IsPrismCr param.Opt[bool] `json:"isPrismCr,omitzero"`
	// Human readable name for this operation.
	Operation param.Opt[string] `json:"operation,omitzero"`
	// 1-n priority for this collection requirement.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Reconnaissance Survey information the operator needs.
	ReconSurvey param.Opt[string] `json:"reconSurvey,omitzero"`
	// Record id.
	RecordID param.Opt[string] `json:"recordId,omitzero"`
	// Region of the collection requirement.
	Region param.Opt[string] `json:"region,omitzero"`
	// Sub category of primary intelligence to be collected for this requirement.
	Secondary param.Opt[bool] `json:"secondary,omitzero"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance param.Opt[string] `json:"specialComGuidance,omitzero"`
	// Start time for this requirement, should be within the mission time window.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time for this requirement, should be within the mission time window.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	// Subregion of the collection requirement.
	Subregion param.Opt[string] `json:"subregion,omitzero"`
	// The name of the military unit that this assigned collection requirement will
	// support.
	SupportedUnit param.Opt[string] `json:"supportedUnit,omitzero"`
	// Type collection this requirement applies to.
	Type                    param.Opt[string]                                                                     `json:"type,omitzero"`
	CriticalTimes           IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementCriticalTimes           `json:"criticalTimes,omitzero"`
	ExploitationRequirement IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirement `json:"exploitationRequirement,omitzero"`
	// Array of POI Id's for the targets being tasked.
	TargetList []string `json:"targetList,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirement) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties EarliestImagingTime, LatestImagingTime are required.
type IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementCriticalTimes struct {
	// Critical start time to collect an image for this requirement.
	EarliestImagingTime time.Time `json:"earliestImagingTime,required" format:"date-time"`
	// Critical stop time to collect an image for this requirement.
	LatestImagingTime time.Time `json:"latestImagingTime,required" format:"date-time"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementCriticalTimes) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementCriticalTimes
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirement struct {
	// Exploitation requirement id.
	ID param.Opt[string] `json:"id,omitzero"`
	// Amplifying data for the exploitation requirement.
	Amplification param.Opt[string] `json:"amplification,omitzero"`
	// List of e-mails to disseminate collection verification information.
	Dissemination param.Opt[string] `json:"dissemination,omitzero"`
	// Essential Elements of Information.
	Eei param.Opt[string] `json:"eei,omitzero"`
	// The reporting criteria of the collection requirement.
	ReportingCriteria param.Opt[string]                                                                        `json:"reportingCriteria,omitzero"`
	Poc               IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirementPoc `json:"poc,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirement) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirementPoc struct {
	// Unique identifier of the collection requirement POC.
	ID param.Opt[string] `json:"id,omitzero"`
	// Callsign of the POC.
	Callsign param.Opt[string] `json:"callsign,omitzero"`
	// Chat name of the POC.
	ChatName param.Opt[string] `json:"chatName,omitzero"`
	// Chat system the POC is accessing.
	ChatSystem param.Opt[string] `json:"chatSystem,omitzero"`
	// Email address of the POC.
	Email param.Opt[string] `json:"email,omitzero"`
	// Name of the POC.
	Name param.Opt[string] `json:"name,omitzero"`
	// Amplifying notes about the POC.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Phone number of the POC.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Radio Frequency the POC is on.
	RadioFrequency param.Opt[float64] `json:"radioFrequency,omitzero"`
	// Unit the POC belongs to.
	Unit param.Opt[string] `json:"unit,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirementPoc) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirementPoc
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyTasking struct {
	// Tasking Unique Identifier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Eight line.
	EightLine param.Opt[string] `json:"eightLine,omitzero"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance param.Opt[string] `json:"specialComGuidance,omitzero"`
	// Value of the Sensitive Reconnaissance Operations Track.
	SroTrack param.Opt[string] `json:"sroTrack,omitzero"`
	// Human readable definition of this taskings Area Of Responsibility.
	TaskingAor param.Opt[string] `json:"taskingAOR,omitzero"`
	// Tasking geographical collection area.
	TaskingCollectionArea param.Opt[string] `json:"taskingCollectionArea,omitzero"`
	// Country code of the tasking. A Country may represent countries, multi-national
	// consortiums, and international organizations.
	TaskingCountry param.Opt[string] `json:"taskingCountry,omitzero"`
	// Tasking emphasis.
	TaskingEmphasis param.Opt[string] `json:"taskingEmphasis,omitzero"`
	// Joint Operations Area.
	TaskingJoa param.Opt[string] `json:"taskingJoa,omitzero"`
	// Tasking operation name.
	TaskingOperation param.Opt[string] `json:"taskingOperation,omitzero"`
	// Primary type of intelligence to be collected during the mission.
	TaskingPrimaryIntelDiscipline param.Opt[string] `json:"taskingPrimaryIntelDiscipline,omitzero"`
	// Sub category of primary intelligence to be collected.
	TaskingPrimarySubCategory param.Opt[string] `json:"taskingPrimarySubCategory,omitzero"`
	// Tasking Priority (1-n).
	TaskingPriority param.Opt[float64] `json:"taskingPriority,omitzero"`
	// Region of the tasking.
	TaskingRegion param.Opt[string] `json:"taskingRegion,omitzero"`
	// Time of retasking, in ISO 8601 UTC format.
	TaskingRetaskTime param.Opt[time.Time] `json:"taskingRetaskTime,omitzero" format:"date-time"`
	// What is the primary objective (role) of this task.
	TaskingRole param.Opt[string] `json:"taskingRole,omitzero"`
	// Type of tasking intelligence to be collected second.
	TaskingSecondaryIntelDiscipline param.Opt[string] `json:"taskingSecondaryIntelDiscipline,omitzero"`
	// Mission sub category for secondary intelligence discipline to be collected.
	TaskingSecondarySubCategory param.Opt[string] `json:"taskingSecondarySubCategory,omitzero"`
	// WGS-84 latitude of the start position, in degrees. -90 to 90 degrees (negative
	// values south of equator).
	TaskingStartPointLat param.Opt[float64] `json:"taskingStartPointLat,omitzero"`
	// WGS-84 longitude of the start position, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	TaskingStartPointLong param.Opt[float64] `json:"taskingStartPointLong,omitzero"`
	// Subregion of the tasking.
	TaskingSubRegion param.Opt[string] `json:"taskingSubRegion,omitzero"`
	// Military Base to transmit the dissemination of this data.
	TaskingSupportedUnit param.Opt[string] `json:"taskingSupportedUnit,omitzero"`
	// A synchronization matrix is used to organize the logistics synchronization
	// process during a mission.
	TaskingSyncMatrixBin param.Opt[string]                                                 `json:"taskingSyncMatrixBin,omitzero"`
	CollectionPeriods    IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriods `json:"collectionPeriods,omitzero"`
	// Type of collection tasked.
	//
	// Any of "Simultaneous", "Sequential", "Operationally", "Driven", "Priority",
	// "Order".
	CollectionType string `json:"collectionType,omitzero"`
	// Tasking desired collection requirements.
	TaskingCollectionRequirements []IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirement `json:"taskingCollectionRequirements,omitzero"`
	// Type of tasking.
	//
	// Any of "Deliberate", "Dynamic", "Training", "Transit".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTasking) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTasking
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[IsrCollectionUnvalidatedPublishParamsBodyTasking](
		"CollectionType", false, "Simultaneous", "Sequential", "Operationally", "Driven", "Priority", "Order",
	)
	apijson.RegisterFieldValidator[IsrCollectionUnvalidatedPublishParamsBodyTasking](
		"Type", false, "Deliberate", "Dynamic", "Training", "Transit",
	)
}

type IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriods struct {
	// Actual start and stop for the collection.
	Actual  []IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsActual `json:"actual,omitzero"`
	Planned IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlanned  `json:"planned,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriods) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriods
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsActual struct {
	// Unique Identifier of actual collection period for historical archive.
	ID param.Opt[string] `json:"id,omitzero"`
	// Start time the collection actually occurred, in ISO 8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time the collection actually occurred, in ISO 8601 UTC format.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsActual) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsActual
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlanned struct {
	// Start time of collection, in ISO 8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time of collection, in ISO 8601 UTC format.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	// Additional start and stop for the collection.
	Additional []IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlannedAdditional `json:"additional,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlanned) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlanned
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlannedAdditional struct {
	// Unique Identifier of additional collection period.
	ID param.Opt[string] `json:"id,omitzero"`
	// Start time of collection, in ISO 8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time of collection, in ISO 8601 UTC format.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlannedAdditional) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlannedAdditional
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirement struct {
	// Collection Requirement Unique Identifier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Country code of the collection requirement. A Country may represent countries,
	// multi-national consortiums, and international organizations.
	Country param.Opt[string] `json:"country,omitzero"`
	// Collection Requirement Unique Identifier.
	CridNumbers param.Opt[string] `json:"cridNumbers,omitzero"`
	// Is this collection requirement an emphasized/critical requirement.
	Emphasized param.Opt[bool] `json:"emphasized,omitzero"`
	// Encryption hashing algorithm.
	Hash param.Opt[string] `json:"hash,omitzero"`
	// Primary type of intelligence to be collected for this requirement.
	IntelDiscipline param.Opt[string] `json:"intelDiscipline,omitzero"`
	// Is this collection request for the Prism system?.
	IsPrismCr param.Opt[bool] `json:"isPrismCr,omitzero"`
	// Human readable name for this operation.
	Operation param.Opt[string] `json:"operation,omitzero"`
	// 1-n priority for this collection requirement.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Reconnaissance Survey information the operator needs.
	ReconSurvey param.Opt[string] `json:"reconSurvey,omitzero"`
	// Record id.
	RecordID param.Opt[string] `json:"recordId,omitzero"`
	// Region of the collection requirement.
	Region param.Opt[string] `json:"region,omitzero"`
	// Sub category of primary intelligence to be collected for this requirement.
	Secondary param.Opt[bool] `json:"secondary,omitzero"`
	// Free text field for the user to specify special instructions needed for this
	// collection.
	SpecialComGuidance param.Opt[string] `json:"specialComGuidance,omitzero"`
	// Start time for this requirement, should be within the mission time window.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Stop time for this requirement, should be within the mission time window.
	Stop param.Opt[time.Time] `json:"stop,omitzero" format:"date-time"`
	// Subregion of the collection requirement.
	Subregion param.Opt[string] `json:"subregion,omitzero"`
	// The name of the military unit that this assigned collection requirement will
	// support.
	SupportedUnit param.Opt[string] `json:"supportedUnit,omitzero"`
	// Type collection this requirement applies to.
	Type                    param.Opt[string]                                                                                   `json:"type,omitzero"`
	CriticalTimes           IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementCriticalTimes           `json:"criticalTimes,omitzero"`
	ExploitationRequirement IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement `json:"exploitationRequirement,omitzero"`
	// Array of POI Id's for the targets being tasked.
	TargetList []string `json:"targetList,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirement) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties EarliestImagingTime, LatestImagingTime are required.
type IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementCriticalTimes struct {
	// Critical start time to collect an image for this requirement.
	EarliestImagingTime time.Time `json:"earliestImagingTime,required" format:"date-time"`
	// Critical stop time to collect an image for this requirement.
	LatestImagingTime time.Time `json:"latestImagingTime,required" format:"date-time"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementCriticalTimes) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementCriticalTimes
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement struct {
	// Exploitation requirement id.
	ID param.Opt[string] `json:"id,omitzero"`
	// Amplifying data for the exploitation requirement.
	Amplification param.Opt[string] `json:"amplification,omitzero"`
	// List of e-mails to disseminate collection verification information.
	Dissemination param.Opt[string] `json:"dissemination,omitzero"`
	// Essential Elements of Information.
	Eei param.Opt[string] `json:"eei,omitzero"`
	// The reporting criteria of the collection requirement.
	ReportingCriteria param.Opt[string]                                                                                      `json:"reportingCriteria,omitzero"`
	Poc               IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc `json:"poc,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc struct {
	// Unique identifier of the collection requirement POC.
	ID param.Opt[string] `json:"id,omitzero"`
	// Callsign of the POC.
	Callsign param.Opt[string] `json:"callsign,omitzero"`
	// Chat name of the POC.
	ChatName param.Opt[string] `json:"chatName,omitzero"`
	// Chat system the POC is accessing.
	ChatSystem param.Opt[string] `json:"chatSystem,omitzero"`
	// Email address of the POC.
	Email param.Opt[string] `json:"email,omitzero"`
	// Name of the POC.
	Name param.Opt[string] `json:"name,omitzero"`
	// Amplifying notes about the POC.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Phone number of the POC.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Radio Frequency the POC is on.
	RadioFrequency param.Opt[float64] `json:"radioFrequency,omitzero"`
	// Unit the POC belongs to.
	Unit param.Opt[string] `json:"unit,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc
	return param.MarshalObject(r, (*shadow)(&r))
}

type IsrCollectionUnvalidatedPublishParamsBodyTransit struct {
	// Transit Unique Identifier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Military Base to transmit the dissemination of this data.
	Base param.Opt[string] `json:"base,omitzero"`
	// Length of mission in milliseconds.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	paramObj
}

func (r IsrCollectionUnvalidatedPublishParamsBodyTransit) MarshalJSON() (data []byte, err error) {
	type shadow IsrCollectionUnvalidatedPublishParamsBodyTransit
	return param.MarshalObject(r, (*shadow)(&r))
}
