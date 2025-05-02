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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// AviationriskmanagementService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAviationriskmanagementService] method instead.
type AviationriskmanagementService struct {
	Options []option.RequestOption
}

// NewAviationriskmanagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAviationriskmanagementService(opts ...option.RequestOption) (r AviationriskmanagementService) {
	r = AviationriskmanagementService{}
	r.Options = opts
	return
}

// Service operation to take a single Aviation Risk Management record as a POST
// body and ingest into the database. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *AviationriskmanagementService) New(ctx context.Context, body AviationriskmanagementNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aviationriskmanagement"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Aviation Risk Management record by its unique
// ID passed as a path parameter.
func (r *AviationriskmanagementService) Get(ctx context.Context, id string, query AviationriskmanagementGetParams, opts ...option.RequestOption) (res *AviationriskmanagementGetResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *AviationriskmanagementService) Update(ctx context.Context, id string, body AviationriskmanagementUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/aviationriskmanagement/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to delete an Aviation Risk Management record specified by the
// passed ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AviationriskmanagementService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
func (r *AviationriskmanagementService) Count(ctx context.Context, query AviationriskmanagementCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *AviationriskmanagementService) NewBulk(ctx context.Context, body AviationriskmanagementNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aviationriskmanagement/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *AviationriskmanagementService) Query(ctx context.Context, query AviationriskmanagementQueryParams, opts ...option.RequestOption) (res *[]AviationriskmanagementQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/aviationriskmanagement"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *AviationriskmanagementService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/aviationriskmanagement/queryhelp"
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
func (r *AviationriskmanagementService) Tuple(ctx context.Context, query AviationriskmanagementTupleParams, opts ...option.RequestOption) (res *[]AviationriskmanagementTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/aviationriskmanagement/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple Aviation Risk Management records as a POST
// body and ingest into the database. This operation is intended to be used for
// automated feeds into UDL. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *AviationriskmanagementService) UnvalidatedPublish(ctx context.Context, body AviationriskmanagementUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-aviationriskmanagement"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
type AviationriskmanagementGetResponse struct {
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
	DataMode AviationriskmanagementGetResponseDataMode `json:"dataMode,required"`
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                 resp.Field
		DataMode                              resp.Field
		IDMission                             resp.Field
		Source                                resp.Field
		ID                                    resp.Field
		AviationRiskManagementWorksheetRecord resp.Field
		CreatedAt                             resp.Field
		CreatedBy                             resp.Field
		ExtMissionID                          resp.Field
		MissionNumber                         resp.Field
		OrgID                                 resp.Field
		Origin                                resp.Field
		OrigNetwork                           resp.Field
		SourceDl                              resp.Field
		UnitID                                resp.Field
		UpdatedAt                             resp.Field
		UpdatedBy                             resp.Field
		ExtraFields                           map[string]resp.Field
		raw                                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AviationriskmanagementGetResponse) UnmarshalJSON(data []byte) error {
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
type AviationriskmanagementGetResponseDataMode string

const (
	AviationriskmanagementGetResponseDataModeReal      AviationriskmanagementGetResponseDataMode = "REAL"
	AviationriskmanagementGetResponseDataModeTest      AviationriskmanagementGetResponseDataMode = "TEST"
	AviationriskmanagementGetResponseDataModeSimulated AviationriskmanagementGetResponseDataMode = "SIMULATED"
	AviationriskmanagementGetResponseDataModeExercise  AviationriskmanagementGetResponseDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
type AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecord struct {
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
	AviationRiskManagementWorksheetScore []AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		MissionDate                          resp.Field
		AircraftMds                          resp.Field
		ApprovalPending                      resp.Field
		Approved                             resp.Field
		AviationRiskManagementWorksheetScore resp.Field
		DispositionComments                  resp.Field
		ExtRecordID                          resp.Field
		Itinerary                            resp.Field
		LastUpdatedAt                        resp.Field
		Remarks                              resp.Field
		SeverityLevel                        resp.Field
		SubmissionDate                       resp.Field
		TierNumber                           resp.Field
		TotalScore                           resp.Field
		UserID                               resp.Field
		ExtraFields                          map[string]resp.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecord) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
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
	AviationRiskManagementSortie []AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ApprovalDate                 resp.Field
		ApprovedBy                   resp.Field
		ApprovedCode                 resp.Field
		AviationRiskManagementSortie resp.Field
		ExtScoreID                   resp.Field
		RiskCategory                 resp.Field
		RiskDescription              resp.Field
		RiskKey                      resp.Field
		RiskName                     resp.Field
		Score                        resp.Field
		ScoreRemark                  resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ExtSortieID resp.Field
		IDSortie    resp.Field
		LegNum      resp.Field
		SortieScore resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationriskmanagementGetResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
type AviationriskmanagementQueryResponse struct {
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
	DataMode AviationriskmanagementQueryResponseDataMode `json:"dataMode,required"`
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                 resp.Field
		DataMode                              resp.Field
		IDMission                             resp.Field
		Source                                resp.Field
		ID                                    resp.Field
		AviationRiskManagementWorksheetRecord resp.Field
		CreatedAt                             resp.Field
		CreatedBy                             resp.Field
		ExtMissionID                          resp.Field
		MissionNumber                         resp.Field
		OrgID                                 resp.Field
		Origin                                resp.Field
		OrigNetwork                           resp.Field
		SourceDl                              resp.Field
		UnitID                                resp.Field
		UpdatedAt                             resp.Field
		UpdatedBy                             resp.Field
		ExtraFields                           map[string]resp.Field
		raw                                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *AviationriskmanagementQueryResponse) UnmarshalJSON(data []byte) error {
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
type AviationriskmanagementQueryResponseDataMode string

const (
	AviationriskmanagementQueryResponseDataModeReal      AviationriskmanagementQueryResponseDataMode = "REAL"
	AviationriskmanagementQueryResponseDataModeTest      AviationriskmanagementQueryResponseDataMode = "TEST"
	AviationriskmanagementQueryResponseDataModeSimulated AviationriskmanagementQueryResponseDataMode = "SIMULATED"
	AviationriskmanagementQueryResponseDataModeExercise  AviationriskmanagementQueryResponseDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
type AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecord struct {
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
	AviationRiskManagementWorksheetScore []AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		MissionDate                          resp.Field
		AircraftMds                          resp.Field
		ApprovalPending                      resp.Field
		Approved                             resp.Field
		AviationRiskManagementWorksheetScore resp.Field
		DispositionComments                  resp.Field
		ExtRecordID                          resp.Field
		Itinerary                            resp.Field
		LastUpdatedAt                        resp.Field
		Remarks                              resp.Field
		SeverityLevel                        resp.Field
		SubmissionDate                       resp.Field
		TierNumber                           resp.Field
		TotalScore                           resp.Field
		UserID                               resp.Field
		ExtraFields                          map[string]resp.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecord) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
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
	AviationRiskManagementSortie []AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ApprovalDate                 resp.Field
		ApprovedBy                   resp.Field
		ApprovedCode                 resp.Field
		AviationRiskManagementSortie resp.Field
		ExtScoreID                   resp.Field
		RiskCategory                 resp.Field
		RiskDescription              resp.Field
		RiskKey                      resp.Field
		RiskName                     resp.Field
		Score                        resp.Field
		ScoreRemark                  resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ExtSortieID resp.Field
		IDSortie    resp.Field
		LegNum      resp.Field
		SortieScore resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationriskmanagementQueryResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
type AviationriskmanagementTupleResponse struct {
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
	DataMode AviationriskmanagementTupleResponseDataMode `json:"dataMode,required"`
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission string `json:"idMission,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                 resp.Field
		DataMode                              resp.Field
		IDMission                             resp.Field
		Source                                resp.Field
		ID                                    resp.Field
		AviationRiskManagementWorksheetRecord resp.Field
		CreatedAt                             resp.Field
		CreatedBy                             resp.Field
		ExtMissionID                          resp.Field
		MissionNumber                         resp.Field
		OrgID                                 resp.Field
		Origin                                resp.Field
		OrigNetwork                           resp.Field
		SourceDl                              resp.Field
		UnitID                                resp.Field
		UpdatedAt                             resp.Field
		UpdatedBy                             resp.Field
		ExtraFields                           map[string]resp.Field
		raw                                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *AviationriskmanagementTupleResponse) UnmarshalJSON(data []byte) error {
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
type AviationriskmanagementTupleResponseDataMode string

const (
	AviationriskmanagementTupleResponseDataModeReal      AviationriskmanagementTupleResponseDataMode = "REAL"
	AviationriskmanagementTupleResponseDataModeTest      AviationriskmanagementTupleResponseDataMode = "TEST"
	AviationriskmanagementTupleResponseDataModeSimulated AviationriskmanagementTupleResponseDataMode = "SIMULATED"
	AviationriskmanagementTupleResponseDataModeExercise  AviationriskmanagementTupleResponseDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
type AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecord struct {
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
	AviationRiskManagementWorksheetScore []AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		MissionDate                          resp.Field
		AircraftMds                          resp.Field
		ApprovalPending                      resp.Field
		Approved                             resp.Field
		AviationRiskManagementWorksheetScore resp.Field
		DispositionComments                  resp.Field
		ExtRecordID                          resp.Field
		Itinerary                            resp.Field
		LastUpdatedAt                        resp.Field
		Remarks                              resp.Field
		SeverityLevel                        resp.Field
		SubmissionDate                       resp.Field
		TierNumber                           resp.Field
		TotalScore                           resp.Field
		UserID                               resp.Field
		ExtraFields                          map[string]resp.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecord) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
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
	AviationRiskManagementSortie []AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ApprovalDate                 resp.Field
		ApprovedBy                   resp.Field
		ApprovedCode                 resp.Field
		AviationRiskManagementSortie resp.Field
		ExtScoreID                   resp.Field
		RiskCategory                 resp.Field
		RiskDescription              resp.Field
		RiskKey                      resp.Field
		RiskName                     resp.Field
		Score                        resp.Field
		ScoreRemark                  resp.Field
		ExtraFields                  map[string]resp.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ExtSortieID resp.Field
		IDSortie    resp.Field
		LegNum      resp.Field
		SortieScore resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) RawJSON() string {
	return r.JSON.raw
}
func (r *AviationriskmanagementTupleResponseAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AviationriskmanagementNewParams struct {
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
	DataMode AviationriskmanagementNewParamsDataMode `json:"dataMode,omitzero,required"`
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
	AviationRiskManagementWorksheetRecord []AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r AviationriskmanagementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementNewParams
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
type AviationriskmanagementNewParamsDataMode string

const (
	AviationriskmanagementNewParamsDataModeReal      AviationriskmanagementNewParamsDataMode = "REAL"
	AviationriskmanagementNewParamsDataModeTest      AviationriskmanagementNewParamsDataMode = "TEST"
	AviationriskmanagementNewParamsDataModeSimulated AviationriskmanagementNewParamsDataMode = "SIMULATED"
	AviationriskmanagementNewParamsDataModeExercise  AviationriskmanagementNewParamsDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
//
// The property MissionDate is required.
type AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecord struct {
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
	AviationRiskManagementWorksheetScore []AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecord) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecord) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecord
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
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
	AviationRiskManagementSortie []AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie
	return param.MarshalObject(r, (*shadow)(&r))
}

type AviationriskmanagementGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [AviationriskmanagementGetParams]'s query parameters as
// `url.Values`.
func (r AviationriskmanagementGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AviationriskmanagementUpdateParams struct {
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
	DataMode AviationriskmanagementUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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
	AviationRiskManagementWorksheetRecord []AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementUpdateParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r AviationriskmanagementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementUpdateParams
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
type AviationriskmanagementUpdateParamsDataMode string

const (
	AviationriskmanagementUpdateParamsDataModeReal      AviationriskmanagementUpdateParamsDataMode = "REAL"
	AviationriskmanagementUpdateParamsDataModeTest      AviationriskmanagementUpdateParamsDataMode = "TEST"
	AviationriskmanagementUpdateParamsDataModeSimulated AviationriskmanagementUpdateParamsDataMode = "SIMULATED"
	AviationriskmanagementUpdateParamsDataModeExercise  AviationriskmanagementUpdateParamsDataMode = "EXERCISE"
)

// Collection of Aviation Risk Management Worksheet Records.
//
// The property MissionDate is required.
type AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecord struct {
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
	AviationRiskManagementWorksheetScore []AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecord) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecord) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecord
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
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
	AviationRiskManagementSortie []AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie
	return param.MarshalObject(r, (*shadow)(&r))
}

type AviationriskmanagementCountParams struct {
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission   string           `query:"idMission,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementCountParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [AviationriskmanagementCountParams]'s query parameters as
// `url.Values`.
func (r AviationriskmanagementCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AviationriskmanagementNewBulkParams struct {
	Body []AviationriskmanagementNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementNewBulkParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r AviationriskmanagementNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
//
// The properties ClassificationMarking, DataMode, IDMission, Source are required.
type AviationriskmanagementNewBulkParamsBody struct {
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
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Identifier for the unit which this risk management record is evaluated.
	UnitID param.Opt[string] `json:"unitId,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementNewBulkParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[AviationriskmanagementNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of Aviation Risk Management Worksheet Records.
//
// The property MissionDate is required.
type AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord struct {
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
	AviationRiskManagementWorksheetScore []AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
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
	AviationRiskManagementSortie []AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie
	return param.MarshalObject(r, (*shadow)(&r))
}

type AviationriskmanagementQueryParams struct {
	// The unique identifier of the mission to which this risk management record is
	// assigned.
	IDMission   string           `query:"idMission,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementQueryParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [AviationriskmanagementQueryParams]'s query parameters as
// `url.Values`.
func (r AviationriskmanagementQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AviationriskmanagementTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementTupleParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [AviationriskmanagementTupleParams]'s query parameters as
// `url.Values`.
func (r AviationriskmanagementTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AviationriskmanagementUnvalidatedPublishParams struct {
	Body []AviationriskmanagementUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r AviationriskmanagementUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Aviation Risk Management is used to identify, evaluate, and track risks when
// mission planning by accounting for factors such as crew fatigue and mission
// complexity.
//
// The properties ClassificationMarking, DataMode, IDMission, Source are required.
type AviationriskmanagementUnvalidatedPublishParamsBody struct {
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
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
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
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Identifier for the unit which this risk management record is evaluated.
	UnitID param.Opt[string] `json:"unitId,omitzero"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Collection of Aviation Risk Management Worksheet Records.
	AviationRiskManagementWorksheetRecord []AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord `json:"aviationRiskManagementWorksheetRecord,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[AviationriskmanagementUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection of Aviation Risk Management Worksheet Records.
//
// The property MissionDate is required.
type AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord struct {
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
	AviationRiskManagementWorksheetScore []AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore `json:"aviationRiskManagementWorksheetScore,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of Aviation Risk Management worksheet record scores.
type AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore struct {
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
	AviationRiskManagementSortie []AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie `json:"aviationRiskManagementSortie,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore
	return param.MarshalObject(r, (*shadow)(&r))
}

// Collection of aviation risk management worksheet record score aircraft sorties.
type AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie) MarshalJSON() (data []byte, err error) {
	type shadow AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie
	return param.MarshalObject(r, (*shadow)(&r))
}
