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

// AviationRiskManagementService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAviationRiskManagementService] method instead.
type AviationRiskManagementService struct {
	Options []option.RequestOption
}

// NewAviationRiskManagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAviationRiskManagementService(opts ...option.RequestOption) (r AviationRiskManagementService) {
	r = AviationRiskManagementService{}
	r.Options = opts
	return
}

// Service operation to take a single Aviation Risk Management record as a POST
// body and ingest into the database. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *AviationRiskManagementService) New(ctx context.Context, body AviationRiskManagementNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/aviationriskmanagement"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Aviation Risk Management record by its unique
// ID passed as a path parameter.
func (r *AviationRiskManagementService) Get(ctx context.Context, id string, query AviationRiskManagementGetParams, opts ...option.RequestOption) (res *AviationRiskManagementGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aviationriskmanagement/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single Aviation Risk Management record. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *AviationRiskManagementService) Update(ctx context.Context, id string, body AviationRiskManagementUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aviationriskmanagement/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AviationRiskManagementService) List(ctx context.Context, query AviationRiskManagementListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AviationRiskManagementListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/aviationriskmanagement"
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
func (r *AviationRiskManagementService) ListAutoPaging(ctx context.Context, query AviationRiskManagementListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AviationRiskManagementListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an Aviation Risk Management record specified by the
// passed ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AviationRiskManagementService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aviationriskmanagement/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *AviationRiskManagementService) Count(ctx context.Context, query AviationRiskManagementCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/aviationriskmanagement/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// Aviation Risk Management records as a POST body and ingest into the database.
// This operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *AviationRiskManagementService) NewBulk(ctx context.Context, body AviationRiskManagementNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/aviationriskmanagement/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AviationRiskManagementService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (res *AviationRiskManagementQueryHelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/aviationriskmanagement/queryhelp"
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
func (r *AviationRiskManagementService) Tuple(ctx context.Context, query AviationRiskManagementTupleParams, opts ...option.RequestOption) (res *[]AviationRiskManagementTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/aviationriskmanagement/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple Aviation Risk Management records as a POST
// body and ingest into the database. This operation is intended to be used for
// automated feeds into UDL. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AviationRiskManagementService) UnvalidatedPublish(ctx context.Context, body AviationRiskManagementUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-aviationriskmanagement"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
type AviationRiskManagementGetResponse struct {
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
	DataMode AviationRiskManagementGetResponseDataMode `json:"dataMode,required"`
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional mission ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtMissionID string `json:"extMissionId"`
	// The mission number of the mission associated with this record.
	MissionNumber string `json:"missionNumber"`
	// Identifier for the organization which this risk management record is evaluated.
	OrgID string `json:"orgId"`
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
	// Identifier for the unit which this risk management record is evaluated.
	UnitID string `json:"unitId"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                 respjson.Field
		DataMode                              respjson.Field
		IDMission                             respjson.Field
		Source                                respjson.Field
		ID                                    respjson.Field
		AviationRiskManagementWorksheetRecord respjson.Field
		CreatedAt                             respjson.Field
		CreatedBy                             respjson.Field
		ExtMissionID                          respjson.Field
		MissionNumber                         respjson.Field
		OrgID                                 respjson.Field
		Origin                                respjson.Field
		OrigNetwork                           respjson.Field
		SourceDl                              respjson.Field
		UnitID                                respjson.Field
		UpdatedAt                             respjson.Field
		UpdatedBy                             respjson.Field
		ExtraFields                           map[string]respjson.Field
		raw                                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AviationRiskManagementGetResponse) UnmarshalJSON(data []byte) error {
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
type AviationRiskManagementGetResponseDataMode string

const (
	AviationRiskManagementGetResponseDataModeReal      AviationRiskManagementGetResponseDataMode = "REAL"
	AviationRiskManagementGetResponseDataModeTest      AviationRiskManagementGetResponseDataMode = "TEST"
	AviationRiskManagementGetResponseDataModeSimulated AviationRiskManagementGetResponseDataMode = "SIMULATED"
	AviationRiskManagementGetResponseDataModeExercise  AviationRiskManagementGetResponseDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
type AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecord struct {
	// Date of the mission in ISO 8601 date-only format (YYYY-MM-DD).
	MissionDate time.Time `json:"missionDate,required" format:"date"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this risk
	// management worksheet record. Intended as, but not constrained to, MIL-STD-6016
	// environment dependent specific type designations.
	AircraftMds string `json:"aircraftMDS"`
	// Flag indicating the worksheet record is pending approval.
	ApprovalPending bool `json:"approvalPending"`
	// Flag indicating the worksheet record is approved.
	Approved bool `json:"approved"`
	// Collection of Aviation Risk Management worksheet record scores.
	AviationRiskManagementWorksheetScore []AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore"`
	// Comment(s) explaining why the worksheet record has been approved or disapproved.
	DispositionComments string `json:"dispositionComments"`
	// Optional identifier of the worksheet record provided by the data source. This
	// field has no meaning within UDL and is provided as a convenience for systems
	// that require tracking of an internal system generated ID.
	ExtRecordID string `json:"extRecordId"`
	// Sequential order of itinerary locations associated for the mission.
	Itinerary string `json:"itinerary"`
	// Timestamp the worksheet record was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdatedAt time.Time `json:"lastUpdatedAt" format:"date-time"`
	// Remarks and/or comments regarding the worksheet record.
	Remarks string `json:"remarks"`
	// Numeric assignment for the worksheet record severity. 0 - LOW; 1 - MODERATE; 2 -
	// HIGH; 3 - SEVERE.
	SeverityLevel int64 `json:"severityLevel"`
	// Timestamp the worksheet record was submitted, in ISO 8601 UTC format with
	// millisecond precision.
	SubmissionDate time.Time `json:"submissionDate" format:"date-time"`
	// Tier number which the mission is being scored as determined by the data source.
	// For example, Tier 1 may indicate mission planners, Tier 2 may indicate
	// operations personnel, Tier 3 may indicate squadron leadership, and Tier 4 may
	// indicate the aircrew.
	TierNumber int64 `json:"tierNumber"`
	// Total score for the worksheet record as defined by the data source. Larger
	// values indicate a higher risk level. For example, values between 0-10 may
	// indicate a low risk level, where values greater then 40 indicate a severe risk
	// level.
	TotalScore int64 `json:"totalScore"`
	// User identifier associated to the worksheet record.
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionDate                          respjson.Field
		AircraftMds                          respjson.Field
		ApprovalPending                      respjson.Field
		Approved                             respjson.Field
		AviationRiskManagementWorksheetScore respjson.Field
		DispositionComments                  respjson.Field
		ExtRecordID                          respjson.Field
		Itinerary                            respjson.Field
		LastUpdatedAt                        respjson.Field
		Remarks                              respjson.Field
		SeverityLevel                        respjson.Field
		SubmissionDate                       respjson.Field
		TierNumber                           respjson.Field
		TotalScore                           respjson.Field
		UserID                               respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecord) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
	// Timestamp the worksheet record score was approval or disapproval, in ISO 8601
	// UTC format with millisecond precision.
	ApprovalDate time.Time `json:"approvalDate" format:"date-time"`
	// Name of the individual who approved or disapproved of the score.
	ApprovedBy string `json:"approvedBy"`
	// Numeric assignment used to determine score approval. 0 - APPROVAL PENDING (used
	// when score value is 2 or 3); 1 - APPROVED (used when score value is 2 or 3); 2 -
	// DISAPPROVED.
	ApprovedCode int64 `json:"approvedCode"`
	// Collection of aviation risk management worksheet record score aircraft sorties.
	AviationRiskManagementSortie []AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie"`
	// Optional identifier of the worksheet record score provided by the data source.
	// This field has no meaning within UDL and is provided as a convenience for
	// systems that require tracking of an internal system generated ID.
	ExtScoreID string `json:"extScoreId"`
	// The category of the risk factor.
	RiskCategory string `json:"riskCategory"`
	// Description of the risk factor.
	RiskDescription string `json:"riskDescription"`
	// Code or identifier of the risk factor category as defined by the data source.
	RiskKey string `json:"riskKey"`
	// Name of the risk factor.
	RiskName string `json:"riskName"`
	// Score of the worksheet record risk factor. Value ranges from 0 to 3, where a
	// value of 0 indicates a low and a value of 3 indicates severe. A value of -1
	// indicates no score.
	Score int64 `json:"score"`
	// Remarks and/or comments regarding the worksheet score.
	ScoreRemark string `json:"scoreRemark"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalDate                 respjson.Field
		ApprovedBy                   respjson.Field
		ApprovedCode                 respjson.Field
		AviationRiskManagementSortie respjson.Field
		ExtScoreID                   respjson.Field
		RiskCategory                 respjson.Field
		RiskDescription              respjson.Field
		RiskKey                      respjson.Field
		RiskName                     respjson.Field
		Score                        respjson.Field
		ScoreRemark                  respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
	// Optional aircraft sortie ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExtSortieID string `json:"extSortieId"`
	// Unique identifier of an associated Aircraft Sortie that is assigned to this risk
	// management record.
	IDSortie string `json:"idSortie"`
	// The leg number of the sortie.
	LegNum int64 `json:"legNum"`
	// The score of the associated aircraft sortie as defined by the data source. Value
	// ranges from 0 to 3, where a value of 0 indicates a low and a value of 3
	// indicates severe. A value of -1 indicates no score.
	SortieScore int64 `json:"sortieScore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtSortieID respjson.Field
		IDSortie    respjson.Field
		LegNum      respjson.Field
		SortieScore respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationRiskManagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
type AviationRiskManagementListResponse struct {
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
	DataMode AviationRiskManagementListResponseDataMode `json:"dataMode,required"`
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationRiskManagementListResponseAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional mission ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtMissionID string `json:"extMissionId"`
	// The mission number of the mission associated with this record.
	MissionNumber string `json:"missionNumber"`
	// Identifier for the organization which this risk management record is evaluated.
	OrgID string `json:"orgId"`
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
	// Identifier for the unit which this risk management record is evaluated.
	UnitID string `json:"unitId"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                 respjson.Field
		DataMode                              respjson.Field
		IDMission                             respjson.Field
		Source                                respjson.Field
		ID                                    respjson.Field
		AviationRiskManagementWorksheetRecord respjson.Field
		CreatedAt                             respjson.Field
		CreatedBy                             respjson.Field
		ExtMissionID                          respjson.Field
		MissionNumber                         respjson.Field
		OrgID                                 respjson.Field
		Origin                                respjson.Field
		OrigNetwork                           respjson.Field
		SourceDl                              respjson.Field
		UnitID                                respjson.Field
		UpdatedAt                             respjson.Field
		UpdatedBy                             respjson.Field
		ExtraFields                           map[string]respjson.Field
		raw                                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementListResponse) RawJSON() string { return r.JSON.raw }
func (r *AviationRiskManagementListResponse) UnmarshalJSON(data []byte) error {
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
type AviationRiskManagementListResponseDataMode string

const (
	AviationRiskManagementListResponseDataModeReal      AviationRiskManagementListResponseDataMode = "REAL"
	AviationRiskManagementListResponseDataModeTest      AviationRiskManagementListResponseDataMode = "TEST"
	AviationRiskManagementListResponseDataModeSimulated AviationRiskManagementListResponseDataMode = "SIMULATED"
	AviationRiskManagementListResponseDataModeExercise  AviationRiskManagementListResponseDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
type AviationRiskManagementListResponseAviationRiskManagementWorksheetRecord struct {
	// Date of the mission in ISO 8601 date-only format (YYYY-MM-DD).
	MissionDate time.Time `json:"missionDate,required" format:"date"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this risk
	// management worksheet record. Intended as, but not constrained to, MIL-STD-6016
	// environment dependent specific type designations.
	AircraftMds string `json:"aircraftMDS"`
	// Flag indicating the worksheet record is pending approval.
	ApprovalPending bool `json:"approvalPending"`
	// Flag indicating the worksheet record is approved.
	Approved bool `json:"approved"`
	// Collection of Aviation Risk Management worksheet record scores.
	AviationRiskManagementWorksheetScore []AviationRiskManagementListResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore"`
	// Comment(s) explaining why the worksheet record has been approved or disapproved.
	DispositionComments string `json:"dispositionComments"`
	// Optional identifier of the worksheet record provided by the data source. This
	// field has no meaning within UDL and is provided as a convenience for systems
	// that require tracking of an internal system generated ID.
	ExtRecordID string `json:"extRecordId"`
	// Sequential order of itinerary locations associated for the mission.
	Itinerary string `json:"itinerary"`
	// Timestamp the worksheet record was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdatedAt time.Time `json:"lastUpdatedAt" format:"date-time"`
	// Remarks and/or comments regarding the worksheet record.
	Remarks string `json:"remarks"`
	// Numeric assignment for the worksheet record severity. 0 - LOW; 1 - MODERATE; 2 -
	// HIGH; 3 - SEVERE.
	SeverityLevel int64 `json:"severityLevel"`
	// Timestamp the worksheet record was submitted, in ISO 8601 UTC format with
	// millisecond precision.
	SubmissionDate time.Time `json:"submissionDate" format:"date-time"`
	// Tier number which the mission is being scored as determined by the data source.
	// For example, Tier 1 may indicate mission planners, Tier 2 may indicate
	// operations personnel, Tier 3 may indicate squadron leadership, and Tier 4 may
	// indicate the aircrew.
	TierNumber int64 `json:"tierNumber"`
	// Total score for the worksheet record as defined by the data source. Larger
	// values indicate a higher risk level. For example, values between 0-10 may
	// indicate a low risk level, where values greater then 40 indicate a severe risk
	// level.
	TotalScore int64 `json:"totalScore"`
	// User identifier associated to the worksheet record.
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionDate                          respjson.Field
		AircraftMds                          respjson.Field
		ApprovalPending                      respjson.Field
		Approved                             respjson.Field
		AviationRiskManagementWorksheetScore respjson.Field
		DispositionComments                  respjson.Field
		ExtRecordID                          respjson.Field
		Itinerary                            respjson.Field
		LastUpdatedAt                        respjson.Field
		Remarks                              respjson.Field
		SeverityLevel                        respjson.Field
		SubmissionDate                       respjson.Field
		TierNumber                           respjson.Field
		TotalScore                           respjson.Field
		UserID                               respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementListResponseAviationRiskManagementWorksheetRecord) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationRiskManagementListResponseAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationRiskManagementListResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
	// Timestamp the worksheet record score was approval or disapproval, in ISO 8601
	// UTC format with millisecond precision.
	ApprovalDate time.Time `json:"approvalDate" format:"date-time"`
	// Name of the individual who approved or disapproved of the score.
	ApprovedBy string `json:"approvedBy"`
	// Numeric assignment used to determine score approval. 0 - APPROVAL PENDING (used
	// when score value is 2 or 3); 1 - APPROVED (used when score value is 2 or 3); 2 -
	// DISAPPROVED.
	ApprovedCode int64 `json:"approvedCode"`
	// Collection of aviation risk management worksheet record score aircraft sorties.
	AviationRiskManagementSortie []AviationRiskManagementListResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie"`
	// Optional identifier of the worksheet record score provided by the data source.
	// This field has no meaning within UDL and is provided as a convenience for
	// systems that require tracking of an internal system generated ID.
	ExtScoreID string `json:"extScoreId"`
	// The category of the risk factor.
	RiskCategory string `json:"riskCategory"`
	// Description of the risk factor.
	RiskDescription string `json:"riskDescription"`
	// Code or identifier of the risk factor category as defined by the data source.
	RiskKey string `json:"riskKey"`
	// Name of the risk factor.
	RiskName string `json:"riskName"`
	// Score of the worksheet record risk factor. Value ranges from 0 to 3, where a
	// value of 0 indicates a low and a value of 3 indicates severe. A value of -1
	// indicates no score.
	Score int64 `json:"score"`
	// Remarks and/or comments regarding the worksheet score.
	ScoreRemark string `json:"scoreRemark"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalDate                 respjson.Field
		ApprovedBy                   respjson.Field
		ApprovedCode                 respjson.Field
		AviationRiskManagementSortie respjson.Field
		ExtScoreID                   respjson.Field
		RiskCategory                 respjson.Field
		RiskDescription              respjson.Field
		RiskKey                      respjson.Field
		RiskName                     respjson.Field
		Score                        respjson.Field
		ScoreRemark                  respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementListResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationRiskManagementListResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationRiskManagementListResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
	// Optional aircraft sortie ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExtSortieID string `json:"extSortieId"`
	// Unique identifier of an associated Aircraft Sortie that is assigned to this risk
	// management record.
	IDSortie string `json:"idSortie"`
	// The leg number of the sortie.
	LegNum int64 `json:"legNum"`
	// The score of the associated aircraft sortie as defined by the data source. Value
	// ranges from 0 to 3, where a value of 0 indicates a low and a value of 3
	// indicates severe. A value of -1 indicates no score.
	SortieScore int64 `json:"sortieScore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtSortieID respjson.Field
		IDSortie    respjson.Field
		LegNum      respjson.Field
		SortieScore respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementListResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationRiskManagementListResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AviationRiskManagementQueryHelpResponse struct {
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
func (r AviationRiskManagementQueryHelpResponse) RawJSON() string { return r.JSON.raw }
func (r *AviationRiskManagementQueryHelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
type AviationRiskManagementTupleResponse struct {
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
	DataMode AviationRiskManagementTupleResponseDataMode `json:"dataMode,required"`
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional mission ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtMissionID string `json:"extMissionId"`
	// The mission number of the mission associated with this record.
	MissionNumber string `json:"missionNumber"`
	// Identifier for the organization which this risk management record is evaluated.
	OrgID string `json:"orgId"`
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
	// Identifier for the unit which this risk management record is evaluated.
	UnitID string `json:"unitId"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                 respjson.Field
		DataMode                              respjson.Field
		IDMission                             respjson.Field
		Source                                respjson.Field
		ID                                    respjson.Field
		AviationRiskManagementWorksheetRecord respjson.Field
		CreatedAt                             respjson.Field
		CreatedBy                             respjson.Field
		ExtMissionID                          respjson.Field
		MissionNumber                         respjson.Field
		OrgID                                 respjson.Field
		Origin                                respjson.Field
		OrigNetwork                           respjson.Field
		SourceDl                              respjson.Field
		UnitID                                respjson.Field
		UpdatedAt                             respjson.Field
		UpdatedBy                             respjson.Field
		ExtraFields                           map[string]respjson.Field
		raw                                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *AviationRiskManagementTupleResponse) UnmarshalJSON(data []byte) error {
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
type AviationRiskManagementTupleResponseDataMode string

const (
	AviationRiskManagementTupleResponseDataModeReal      AviationRiskManagementTupleResponseDataMode = "REAL"
	AviationRiskManagementTupleResponseDataModeTest      AviationRiskManagementTupleResponseDataMode = "TEST"
	AviationRiskManagementTupleResponseDataModeSimulated AviationRiskManagementTupleResponseDataMode = "SIMULATED"
	AviationRiskManagementTupleResponseDataModeExercise  AviationRiskManagementTupleResponseDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
type AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecord struct {
	// Date of the mission in ISO 8601 date-only format (YYYY-MM-DD).
	MissionDate time.Time `json:"missionDate,required" format:"date"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this risk
	// management worksheet record. Intended as, but not constrained to, MIL-STD-6016
	// environment dependent specific type designations.
	AircraftMds string `json:"aircraftMDS"`
	// Flag indicating the worksheet record is pending approval.
	ApprovalPending bool `json:"approvalPending"`
	// Flag indicating the worksheet record is approved.
	Approved bool `json:"approved"`
	// Collection of Aviation Risk Management worksheet record scores.
	AviationRiskManagementWorksheetScore []AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore"`
	// Comment(s) explaining why the worksheet record has been approved or disapproved.
	DispositionComments string `json:"dispositionComments"`
	// Optional identifier of the worksheet record provided by the data source. This
	// field has no meaning within UDL and is provided as a convenience for systems
	// that require tracking of an internal system generated ID.
	ExtRecordID string `json:"extRecordId"`
	// Sequential order of itinerary locations associated for the mission.
	Itinerary string `json:"itinerary"`
	// Timestamp the worksheet record was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdatedAt time.Time `json:"lastUpdatedAt" format:"date-time"`
	// Remarks and/or comments regarding the worksheet record.
	Remarks string `json:"remarks"`
	// Numeric assignment for the worksheet record severity. 0 - LOW; 1 - MODERATE; 2 -
	// HIGH; 3 - SEVERE.
	SeverityLevel int64 `json:"severityLevel"`
	// Timestamp the worksheet record was submitted, in ISO 8601 UTC format with
	// millisecond precision.
	SubmissionDate time.Time `json:"submissionDate" format:"date-time"`
	// Tier number which the mission is being scored as determined by the data source.
	// For example, Tier 1 may indicate mission planners, Tier 2 may indicate
	// operations personnel, Tier 3 may indicate squadron leadership, and Tier 4 may
	// indicate the aircrew.
	TierNumber int64 `json:"tierNumber"`
	// Total score for the worksheet record as defined by the data source. Larger
	// values indicate a higher risk level. For example, values between 0-10 may
	// indicate a low risk level, where values greater then 40 indicate a severe risk
	// level.
	TotalScore int64 `json:"totalScore"`
	// User identifier associated to the worksheet record.
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionDate                          respjson.Field
		AircraftMds                          respjson.Field
		ApprovalPending                      respjson.Field
		Approved                             respjson.Field
		AviationRiskManagementWorksheetScore respjson.Field
		DispositionComments                  respjson.Field
		ExtRecordID                          respjson.Field
		Itinerary                            respjson.Field
		LastUpdatedAt                        respjson.Field
		Remarks                              respjson.Field
		SeverityLevel                        respjson.Field
		SubmissionDate                       respjson.Field
		TierNumber                           respjson.Field
		TotalScore                           respjson.Field
		UserID                               respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecord) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
	// Timestamp the worksheet record score was approval or disapproval, in ISO 8601
	// UTC format with millisecond precision.
	ApprovalDate time.Time `json:"approvalDate" format:"date-time"`
	// Name of the individual who approved or disapproved of the score.
	ApprovedBy string `json:"approvedBy"`
	// Numeric assignment used to determine score approval. 0 - APPROVAL PENDING (used
	// when score value is 2 or 3); 1 - APPROVED (used when score value is 2 or 3); 2 -
	// DISAPPROVED.
	ApprovedCode int64 `json:"approvedCode"`
	// Collection of aviation risk management worksheet record score aircraft sorties.
	AviationRiskManagementSortie []AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie"`
	// Optional identifier of the worksheet record score provided by the data source.
	// This field has no meaning within UDL and is provided as a convenience for
	// systems that require tracking of an internal system generated ID.
	ExtScoreID string `json:"extScoreId"`
	// The category of the risk factor.
	RiskCategory string `json:"riskCategory"`
	// Description of the risk factor.
	RiskDescription string `json:"riskDescription"`
	// Code or identifier of the risk factor category as defined by the data source.
	RiskKey string `json:"riskKey"`
	// Name of the risk factor.
	RiskName string `json:"riskName"`
	// Score of the worksheet record risk factor. Value ranges from 0 to 3, where a
	// value of 0 indicates a low and a value of 3 indicates severe. A value of -1
	// indicates no score.
	Score int64 `json:"score"`
	// Remarks and/or comments regarding the worksheet score.
	ScoreRemark string `json:"scoreRemark"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalDate                 respjson.Field
		ApprovedBy                   respjson.Field
		ApprovedCode                 respjson.Field
		AviationRiskManagementSortie respjson.Field
		ExtScoreID                   respjson.Field
		RiskCategory                 respjson.Field
		RiskDescription              respjson.Field
		RiskKey                      respjson.Field
		RiskName                     respjson.Field
		Score                        respjson.Field
		ScoreRemark                  respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
	// Optional aircraft sortie ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExtSortieID string `json:"extSortieId"`
	// Unique identifier of an associated Aircraft Sortie that is assigned to this risk
	// management record.
	IDSortie string `json:"idSortie"`
	// The leg number of the sortie.
	LegNum int64 `json:"legNum"`
	// The score of the associated aircraft sortie as defined by the data source. Value
	// ranges from 0 to 3, where a value of 0 indicates a low and a value of 3
	// indicates severe. A value of -1 indicates no score.
	SortieScore int64 `json:"sortieScore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtSortieID respjson.Field
		IDSortie    respjson.Field
		LegNum      respjson.Field
		SortieScore respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationRiskManagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AviationRiskManagementNewParams struct {
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
	DataMode AviationRiskManagementNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional mission ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtMissionID param.Opt[string] `json:"extMissionId,omitzero"`
	// The mission number of the mission associated with this record.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// Identifier for the organization which this risk management record is evaluated.
	OrgID param.Opt[string] `json:"orgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Identifier for the unit which this risk management record is evaluated.
	UnitID param.Opt[string] `json:"unitId,omitzero"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord,omitzero"`
	paramObj
}

func (r AviationRiskManagementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementNewParams) UnmarshalJSON(data []byte) error {
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
type AviationRiskManagementNewParamsDataMode string

const (
	AviationRiskManagementNewParamsDataModeReal      AviationRiskManagementNewParamsDataMode = "REAL"
	AviationRiskManagementNewParamsDataModeTest      AviationRiskManagementNewParamsDataMode = "TEST"
	AviationRiskManagementNewParamsDataModeSimulated AviationRiskManagementNewParamsDataMode = "SIMULATED"
	AviationRiskManagementNewParamsDataModeExercise  AviationRiskManagementNewParamsDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
//
// The property MissionDate is required.
type AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecord struct {
	// Date of the mission in ISO 8601 date-only format (YYYY-MM-DD).
	MissionDate time.Time `json:"missionDate,required" format:"date"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this risk
	// management worksheet record. Intended as, but not constrained to, MIL-STD-6016
	// environment dependent specific type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Flag indicating the worksheet record is pending approval.
	ApprovalPending param.Opt[bool] `json:"approvalPending,omitzero"`
	// Flag indicating the worksheet record is approved.
	Approved param.Opt[bool] `json:"approved,omitzero"`
	// Comment(s) explaining why the worksheet record has been approved or disapproved.
	DispositionComments param.Opt[string] `json:"dispositionComments,omitzero"`
	// Optional identifier of the worksheet record provided by the data source. This
	// field has no meaning within UDL and is provided as a convenience for systems
	// that require tracking of an internal system generated ID.
	ExtRecordID param.Opt[string] `json:"extRecordId,omitzero"`
	// Sequential order of itinerary locations associated for the mission.
	Itinerary param.Opt[string] `json:"itinerary,omitzero"`
	// Timestamp the worksheet record was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdatedAt param.Opt[time.Time] `json:"lastUpdatedAt,omitzero" format:"date-time"`
	// Remarks and/or comments regarding the worksheet record.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// Numeric assignment for the worksheet record severity. 0 - LOW; 1 - MODERATE; 2 -
	// HIGH; 3 - SEVERE.
	SeverityLevel param.Opt[int64] `json:"severityLevel,omitzero"`
	// Timestamp the worksheet record was submitted, in ISO 8601 UTC format with
	// millisecond precision.
	SubmissionDate param.Opt[time.Time] `json:"submissionDate,omitzero" format:"date-time"`
	// Tier number which the mission is being scored as determined by the data source.
	// For example, Tier 1 may indicate mission planners, Tier 2 may indicate
	// operations personnel, Tier 3 may indicate squadron leadership, and Tier 4 may
	// indicate the aircrew.
	TierNumber param.Opt[int64] `json:"tierNumber,omitzero"`
	// Total score for the worksheet record as defined by the data source. Larger
	// values indicate a higher risk level. For example, values between 0-10 may
	// indicate a low risk level, where values greater then 40 indicate a severe risk
	// level.
	TotalScore param.Opt[int64] `json:"totalScore,omitzero"`
	// User identifier associated to the worksheet record.
	UserID param.Opt[string] `json:"userId,omitzero"`
	// Collection of Aviation Risk Management worksheet record scores.
	AviationRiskManagementWorksheetScore []AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore,omitzero"`
	paramObj
}

func (r AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecord) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecord
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
	// Timestamp the worksheet record score was approval or disapproval, in ISO 8601
	// UTC format with millisecond precision.
	ApprovalDate param.Opt[time.Time] `json:"approvalDate,omitzero" format:"date-time"`
	// Name of the individual who approved or disapproved of the score.
	ApprovedBy param.Opt[string] `json:"approvedBy,omitzero"`
	// Numeric assignment used to determine score approval. 0 - APPROVAL PENDING (used
	// when score value is 2 or 3); 1 - APPROVED (used when score value is 2 or 3); 2 -
	// DISAPPROVED.
	ApprovedCode param.Opt[int64] `json:"approvedCode,omitzero"`
	// Optional identifier of the worksheet record score provided by the data source.
	// This field has no meaning within UDL and is provided as a convenience for
	// systems that require tracking of an internal system generated ID.
	ExtScoreID param.Opt[string] `json:"extScoreId,omitzero"`
	// The category of the risk factor.
	RiskCategory param.Opt[string] `json:"riskCategory,omitzero"`
	// Description of the risk factor.
	RiskDescription param.Opt[string] `json:"riskDescription,omitzero"`
	// Code or identifier of the risk factor category as defined by the data source.
	RiskKey param.Opt[string] `json:"riskKey,omitzero"`
	// Name of the risk factor.
	RiskName param.Opt[string] `json:"riskName,omitzero"`
	// Score of the worksheet record risk factor. Value ranges from 0 to 3, where a
	// value of 0 indicates a low and a value of 3 indicates severe. A value of -1
	// indicates no score.
	Score param.Opt[int64] `json:"score,omitzero"`
	// Remarks and/or comments regarding the worksheet score.
	ScoreRemark param.Opt[string] `json:"scoreRemark,omitzero"`
	// Collection of aviation risk management worksheet record score aircraft sorties.
	AviationRiskManagementSortie []AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie,omitzero"`
	paramObj
}

func (r AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
	// Optional aircraft sortie ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExtSortieID param.Opt[string] `json:"extSortieId,omitzero"`
	// Unique identifier of an associated Aircraft Sortie that is assigned to this risk
	// management record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// The leg number of the sortie.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The score of the associated aircraft sortie as defined by the data source. Value
	// ranges from 0 to 3, where a value of 0 indicates a low and a value of 3
	// indicates severe. A value of -1 indicates no score.
	SortieScore param.Opt[int64] `json:"sortieScore,omitzero"`
	paramObj
}

func (r AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AviationRiskManagementGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AviationRiskManagementGetParams]'s query parameters as
// `url.Values`.
func (r AviationRiskManagementGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AviationRiskManagementUpdateParams struct {
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
	DataMode AviationRiskManagementUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional mission ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtMissionID param.Opt[string] `json:"extMissionId,omitzero"`
	// The mission number of the mission associated with this record.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// Identifier for the organization which this risk management record is evaluated.
	OrgID param.Opt[string] `json:"orgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Identifier for the unit which this risk management record is evaluated.
	UnitID param.Opt[string] `json:"unitId,omitzero"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord,omitzero"`
	paramObj
}

func (r AviationRiskManagementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementUpdateParams) UnmarshalJSON(data []byte) error {
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
type AviationRiskManagementUpdateParamsDataMode string

const (
	AviationRiskManagementUpdateParamsDataModeReal      AviationRiskManagementUpdateParamsDataMode = "REAL"
	AviationRiskManagementUpdateParamsDataModeTest      AviationRiskManagementUpdateParamsDataMode = "TEST"
	AviationRiskManagementUpdateParamsDataModeSimulated AviationRiskManagementUpdateParamsDataMode = "SIMULATED"
	AviationRiskManagementUpdateParamsDataModeExercise  AviationRiskManagementUpdateParamsDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
//
// The property MissionDate is required.
type AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecord struct {
	// Date of the mission in ISO 8601 date-only format (YYYY-MM-DD).
	MissionDate time.Time `json:"missionDate,required" format:"date"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this risk
	// management worksheet record. Intended as, but not constrained to, MIL-STD-6016
	// environment dependent specific type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Flag indicating the worksheet record is pending approval.
	ApprovalPending param.Opt[bool] `json:"approvalPending,omitzero"`
	// Flag indicating the worksheet record is approved.
	Approved param.Opt[bool] `json:"approved,omitzero"`
	// Comment(s) explaining why the worksheet record has been approved or disapproved.
	DispositionComments param.Opt[string] `json:"dispositionComments,omitzero"`
	// Optional identifier of the worksheet record provided by the data source. This
	// field has no meaning within UDL and is provided as a convenience for systems
	// that require tracking of an internal system generated ID.
	ExtRecordID param.Opt[string] `json:"extRecordId,omitzero"`
	// Sequential order of itinerary locations associated for the mission.
	Itinerary param.Opt[string] `json:"itinerary,omitzero"`
	// Timestamp the worksheet record was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdatedAt param.Opt[time.Time] `json:"lastUpdatedAt,omitzero" format:"date-time"`
	// Remarks and/or comments regarding the worksheet record.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// Numeric assignment for the worksheet record severity. 0 - LOW; 1 - MODERATE; 2 -
	// HIGH; 3 - SEVERE.
	SeverityLevel param.Opt[int64] `json:"severityLevel,omitzero"`
	// Timestamp the worksheet record was submitted, in ISO 8601 UTC format with
	// millisecond precision.
	SubmissionDate param.Opt[time.Time] `json:"submissionDate,omitzero" format:"date-time"`
	// Tier number which the mission is being scored as determined by the data source.
	// For example, Tier 1 may indicate mission planners, Tier 2 may indicate
	// operations personnel, Tier 3 may indicate squadron leadership, and Tier 4 may
	// indicate the aircrew.
	TierNumber param.Opt[int64] `json:"tierNumber,omitzero"`
	// Total score for the worksheet record as defined by the data source. Larger
	// values indicate a higher risk level. For example, values between 0-10 may
	// indicate a low risk level, where values greater then 40 indicate a severe risk
	// level.
	TotalScore param.Opt[int64] `json:"totalScore,omitzero"`
	// User identifier associated to the worksheet record.
	UserID param.Opt[string] `json:"userId,omitzero"`
	// Collection of Aviation Risk Management worksheet record scores.
	AviationRiskManagementWorksheetScore []AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore,omitzero"`
	paramObj
}

func (r AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecord) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecord
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
	// Timestamp the worksheet record score was approval or disapproval, in ISO 8601
	// UTC format with millisecond precision.
	ApprovalDate param.Opt[time.Time] `json:"approvalDate,omitzero" format:"date-time"`
	// Name of the individual who approved or disapproved of the score.
	ApprovedBy param.Opt[string] `json:"approvedBy,omitzero"`
	// Numeric assignment used to determine score approval. 0 - APPROVAL PENDING (used
	// when score value is 2 or 3); 1 - APPROVED (used when score value is 2 or 3); 2 -
	// DISAPPROVED.
	ApprovedCode param.Opt[int64] `json:"approvedCode,omitzero"`
	// Optional identifier of the worksheet record score provided by the data source.
	// This field has no meaning within UDL and is provided as a convenience for
	// systems that require tracking of an internal system generated ID.
	ExtScoreID param.Opt[string] `json:"extScoreId,omitzero"`
	// The category of the risk factor.
	RiskCategory param.Opt[string] `json:"riskCategory,omitzero"`
	// Description of the risk factor.
	RiskDescription param.Opt[string] `json:"riskDescription,omitzero"`
	// Code or identifier of the risk factor category as defined by the data source.
	RiskKey param.Opt[string] `json:"riskKey,omitzero"`
	// Name of the risk factor.
	RiskName param.Opt[string] `json:"riskName,omitzero"`
	// Score of the worksheet record risk factor. Value ranges from 0 to 3, where a
	// value of 0 indicates a low and a value of 3 indicates severe. A value of -1
	// indicates no score.
	Score param.Opt[int64] `json:"score,omitzero"`
	// Remarks and/or comments regarding the worksheet score.
	ScoreRemark param.Opt[string] `json:"scoreRemark,omitzero"`
	// Collection of aviation risk management worksheet record score aircraft sorties.
	AviationRiskManagementSortie []AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie,omitzero"`
	paramObj
}

func (r AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
	// Optional aircraft sortie ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExtSortieID param.Opt[string] `json:"extSortieId,omitzero"`
	// Unique identifier of an associated Aircraft Sortie that is assigned to this risk
	// management record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// The leg number of the sortie.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The score of the associated aircraft sortie as defined by the data source. Value
	// ranges from 0 to 3, where a value of 0 indicates a low and a value of 3
	// indicates severe. A value of -1 indicates no score.
	SortieScore param.Opt[int64] `json:"sortieScore,omitzero"`
	paramObj
}

func (r AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AviationRiskManagementListParams struct {
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission   string           `query:"idMission,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AviationRiskManagementListParams]'s query parameters as
// `url.Values`.
func (r AviationRiskManagementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AviationRiskManagementCountParams struct {
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission   string           `query:"idMission,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AviationRiskManagementCountParams]'s query parameters as
// `url.Values`.
func (r AviationRiskManagementCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AviationRiskManagementNewBulkParams struct {
	Body []AviationRiskManagementNewBulkParamsBody
	paramObj
}

func (r AviationRiskManagementNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AviationRiskManagementNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
//
// The properties ClassificationMarking, DataMode, IDMission, Source are required.
type AviationRiskManagementNewBulkParamsBody struct {
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
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional mission ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtMissionID param.Opt[string] `json:"extMissionId,omitzero"`
	// The mission number of the mission associated with this record.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// Identifier for the organization which this risk management record is evaluated.
	OrgID param.Opt[string] `json:"orgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Identifier for the unit which this risk management record is evaluated.
	UnitID param.Opt[string] `json:"unitId,omitzero"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord,omitzero"`
	paramObj
}

func (r AviationRiskManagementNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AviationRiskManagementNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of Aviation Risk Management Worksheet Records.
//
// The property MissionDate is required.
type AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord struct {
	// Date of the mission in ISO 8601 date-only format (YYYY-MM-DD).
	MissionDate time.Time `json:"missionDate,required" format:"date"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this risk
	// management worksheet record. Intended as, but not constrained to, MIL-STD-6016
	// environment dependent specific type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Flag indicating the worksheet record is pending approval.
	ApprovalPending param.Opt[bool] `json:"approvalPending,omitzero"`
	// Flag indicating the worksheet record is approved.
	Approved param.Opt[bool] `json:"approved,omitzero"`
	// Comment(s) explaining why the worksheet record has been approved or disapproved.
	DispositionComments param.Opt[string] `json:"dispositionComments,omitzero"`
	// Optional identifier of the worksheet record provided by the data source. This
	// field has no meaning within UDL and is provided as a convenience for systems
	// that require tracking of an internal system generated ID.
	ExtRecordID param.Opt[string] `json:"extRecordId,omitzero"`
	// Sequential order of itinerary locations associated for the mission.
	Itinerary param.Opt[string] `json:"itinerary,omitzero"`
	// Timestamp the worksheet record was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdatedAt param.Opt[time.Time] `json:"lastUpdatedAt,omitzero" format:"date-time"`
	// Remarks and/or comments regarding the worksheet record.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// Numeric assignment for the worksheet record severity. 0 - LOW; 1 - MODERATE; 2 -
	// HIGH; 3 - SEVERE.
	SeverityLevel param.Opt[int64] `json:"severityLevel,omitzero"`
	// Timestamp the worksheet record was submitted, in ISO 8601 UTC format with
	// millisecond precision.
	SubmissionDate param.Opt[time.Time] `json:"submissionDate,omitzero" format:"date-time"`
	// Tier number which the mission is being scored as determined by the data source.
	// For example, Tier 1 may indicate mission planners, Tier 2 may indicate
	// operations personnel, Tier 3 may indicate squadron leadership, and Tier 4 may
	// indicate the aircrew.
	TierNumber param.Opt[int64] `json:"tierNumber,omitzero"`
	// Total score for the worksheet record as defined by the data source. Larger
	// values indicate a higher risk level. For example, values between 0-10 may
	// indicate a low risk level, where values greater then 40 indicate a severe risk
	// level.
	TotalScore param.Opt[int64] `json:"totalScore,omitzero"`
	// User identifier associated to the worksheet record.
	UserID param.Opt[string] `json:"userId,omitzero"`
	// Collection of Aviation Risk Management worksheet record scores.
	AviationRiskManagementWorksheetScore []AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore,omitzero"`
	paramObj
}

func (r AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
	// Timestamp the worksheet record score was approval or disapproval, in ISO 8601
	// UTC format with millisecond precision.
	ApprovalDate param.Opt[time.Time] `json:"approvalDate,omitzero" format:"date-time"`
	// Name of the individual who approved or disapproved of the score.
	ApprovedBy param.Opt[string] `json:"approvedBy,omitzero"`
	// Numeric assignment used to determine score approval. 0 - APPROVAL PENDING (used
	// when score value is 2 or 3); 1 - APPROVED (used when score value is 2 or 3); 2 -
	// DISAPPROVED.
	ApprovedCode param.Opt[int64] `json:"approvedCode,omitzero"`
	// Optional identifier of the worksheet record score provided by the data source.
	// This field has no meaning within UDL and is provided as a convenience for
	// systems that require tracking of an internal system generated ID.
	ExtScoreID param.Opt[string] `json:"extScoreId,omitzero"`
	// The category of the risk factor.
	RiskCategory param.Opt[string] `json:"riskCategory,omitzero"`
	// Description of the risk factor.
	RiskDescription param.Opt[string] `json:"riskDescription,omitzero"`
	// Code or identifier of the risk factor category as defined by the data source.
	RiskKey param.Opt[string] `json:"riskKey,omitzero"`
	// Name of the risk factor.
	RiskName param.Opt[string] `json:"riskName,omitzero"`
	// Score of the worksheet record risk factor. Value ranges from 0 to 3, where a
	// value of 0 indicates a low and a value of 3 indicates severe. A value of -1
	// indicates no score.
	Score param.Opt[int64] `json:"score,omitzero"`
	// Remarks and/or comments regarding the worksheet score.
	ScoreRemark param.Opt[string] `json:"scoreRemark,omitzero"`
	// Collection of aviation risk management worksheet record score aircraft sorties.
	AviationRiskManagementSortie []AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie,omitzero"`
	paramObj
}

func (r AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
	// Optional aircraft sortie ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExtSortieID param.Opt[string] `json:"extSortieId,omitzero"`
	// Unique identifier of an associated Aircraft Sortie that is assigned to this risk
	// management record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// The leg number of the sortie.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The score of the associated aircraft sortie as defined by the data source. Value
	// ranges from 0 to 3, where a value of 0 indicates a low and a value of 3
	// indicates severe. A value of -1 indicates no score.
	SortieScore param.Opt[int64] `json:"sortieScore,omitzero"`
	paramObj
}

func (r AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AviationRiskManagementTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission   string           `query:"idMission,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AviationRiskManagementTupleParams]'s query parameters as
// `url.Values`.
func (r AviationRiskManagementTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AviationRiskManagementUnvalidatedPublishParams struct {
	Body []AviationRiskManagementUnvalidatedPublishParamsBody
	paramObj
}

func (r AviationRiskManagementUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AviationRiskManagementUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
//
// The properties ClassificationMarking, DataMode, IDMission, Source are required.
type AviationRiskManagementUnvalidatedPublishParamsBody struct {
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
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional mission ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExtMissionID param.Opt[string] `json:"extMissionId,omitzero"`
	// The mission number of the mission associated with this record.
	MissionNumber param.Opt[string] `json:"missionNumber,omitzero"`
	// Identifier for the organization which this risk management record is evaluated.
	OrgID param.Opt[string] `json:"orgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Identifier for the unit which this risk management record is evaluated.
	UnitID param.Opt[string] `json:"unitId,omitzero"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord,omitzero"`
	paramObj
}

func (r AviationRiskManagementUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AviationRiskManagementUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of Aviation Risk Management Worksheet Records.
//
// The property MissionDate is required.
type AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord struct {
	// Date of the mission in ISO 8601 date-only format (YYYY-MM-DD).
	MissionDate time.Time `json:"missionDate,required" format:"date"`
	// The aircraft Model Design Series (MDS) designation (e.g. E-2C HAWKEYE, F-15
	// EAGLE, KC-130 HERCULES, etc.) of the aircraft associated with this risk
	// management worksheet record. Intended as, but not constrained to, MIL-STD-6016
	// environment dependent specific type designations.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Flag indicating the worksheet record is pending approval.
	ApprovalPending param.Opt[bool] `json:"approvalPending,omitzero"`
	// Flag indicating the worksheet record is approved.
	Approved param.Opt[bool] `json:"approved,omitzero"`
	// Comment(s) explaining why the worksheet record has been approved or disapproved.
	DispositionComments param.Opt[string] `json:"dispositionComments,omitzero"`
	// Optional identifier of the worksheet record provided by the data source. This
	// field has no meaning within UDL and is provided as a convenience for systems
	// that require tracking of an internal system generated ID.
	ExtRecordID param.Opt[string] `json:"extRecordId,omitzero"`
	// Sequential order of itinerary locations associated for the mission.
	Itinerary param.Opt[string] `json:"itinerary,omitzero"`
	// Timestamp the worksheet record was updated, in ISO 8601 UTC format with
	// millisecond precision.
	LastUpdatedAt param.Opt[time.Time] `json:"lastUpdatedAt,omitzero" format:"date-time"`
	// Remarks and/or comments regarding the worksheet record.
	Remarks param.Opt[string] `json:"remarks,omitzero"`
	// Numeric assignment for the worksheet record severity. 0 - LOW; 1 - MODERATE; 2 -
	// HIGH; 3 - SEVERE.
	SeverityLevel param.Opt[int64] `json:"severityLevel,omitzero"`
	// Timestamp the worksheet record was submitted, in ISO 8601 UTC format with
	// millisecond precision.
	SubmissionDate param.Opt[time.Time] `json:"submissionDate,omitzero" format:"date-time"`
	// Tier number which the mission is being scored as determined by the data source.
	// For example, Tier 1 may indicate mission planners, Tier 2 may indicate
	// operations personnel, Tier 3 may indicate squadron leadership, and Tier 4 may
	// indicate the aircrew.
	TierNumber param.Opt[int64] `json:"tierNumber,omitzero"`
	// Total score for the worksheet record as defined by the data source. Larger
	// values indicate a higher risk level. For example, values between 0-10 may
	// indicate a low risk level, where values greater then 40 indicate a severe risk
	// level.
	TotalScore param.Opt[int64] `json:"totalScore,omitzero"`
	// User identifier associated to the worksheet record.
	UserID param.Opt[string] `json:"userId,omitzero"`
	// Collection of Aviation Risk Management worksheet record scores.
	AviationRiskManagementWorksheetScore []AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore,omitzero"`
	paramObj
}

func (r AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
	// Timestamp the worksheet record score was approval or disapproval, in ISO 8601
	// UTC format with millisecond precision.
	ApprovalDate param.Opt[time.Time] `json:"approvalDate,omitzero" format:"date-time"`
	// Name of the individual who approved or disapproved of the score.
	ApprovedBy param.Opt[string] `json:"approvedBy,omitzero"`
	// Numeric assignment used to determine score approval. 0 - APPROVAL PENDING (used
	// when score value is 2 or 3); 1 - APPROVED (used when score value is 2 or 3); 2 -
	// DISAPPROVED.
	ApprovedCode param.Opt[int64] `json:"approvedCode,omitzero"`
	// Optional identifier of the worksheet record score provided by the data source.
	// This field has no meaning within UDL and is provided as a convenience for
	// systems that require tracking of an internal system generated ID.
	ExtScoreID param.Opt[string] `json:"extScoreId,omitzero"`
	// The category of the risk factor.
	RiskCategory param.Opt[string] `json:"riskCategory,omitzero"`
	// Description of the risk factor.
	RiskDescription param.Opt[string] `json:"riskDescription,omitzero"`
	// Code or identifier of the risk factor category as defined by the data source.
	RiskKey param.Opt[string] `json:"riskKey,omitzero"`
	// Name of the risk factor.
	RiskName param.Opt[string] `json:"riskName,omitzero"`
	// Score of the worksheet record risk factor. Value ranges from 0 to 3, where a
	// value of 0 indicates a low and a value of 3 indicates severe. A value of -1
	// indicates no score.
	Score param.Opt[int64] `json:"score,omitzero"`
	// Remarks and/or comments regarding the worksheet score.
	ScoreRemark param.Opt[string] `json:"scoreRemark,omitzero"`
	// Collection of aviation risk management worksheet record score aircraft sorties.
	AviationRiskManagementSortie []AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie,omitzero"`
	paramObj
}

func (r AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
	// Optional aircraft sortie ID from external systems. This field has no meaning
	// within UDL and is provided as a convenience for systems that require tracking of
	// an internal system generated ID.
	ExtSortieID param.Opt[string] `json:"extSortieId,omitzero"`
	// Unique identifier of an associated Aircraft Sortie that is assigned to this risk
	// management record.
	IDSortie param.Opt[string] `json:"idSortie,omitzero"`
	// The leg number of the sortie.
	LegNum param.Opt[int64] `json:"legNum,omitzero"`
	// The score of the associated aircraft sortie as defined by the data source. Value
	// ranges from 0 to 3, where a value of 0 indicates a low and a value of 3
	// indicates severe. A value of -1 indicates no score.
	SortieScore param.Opt[int64] `json:"sortieScore,omitzero"`
	paramObj
}

func (r AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) MarshalJSON() (data []byte, err error) {
	type shadow AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
