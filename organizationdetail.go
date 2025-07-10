// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// OrganizationdetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationdetailService] method instead.
type OrganizationdetailService struct {
	Options []option.RequestOption
}

// NewOrganizationdetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationdetailService(opts ...option.RequestOption) (r OrganizationdetailService) {
	r = OrganizationdetailService{}
	r.Options = opts
	return
}

// Service operation to take a single OrganizationDetails as a POST body and ingest
// into the database. OrganizationDetails represent details of organizations such
// as a corporation, manufacturer, consortium, government, etc. An organization can
// have detail records from several sources. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *OrganizationdetailService) New(ctx context.Context, body OrganizationdetailNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/organizationdetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an OrganizationDetails object. OrganizationDetails
// represent details of organizations such as a corporation, manufacturer,
// consortium, government, etc. An organization can have detail records from
// several sources. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *OrganizationdetailService) Update(ctx context.Context, id string, body OrganizationdetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/organizationdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OrganizationdetailService) List(ctx context.Context, query OrganizationdetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OrganizationdetailListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/organizationdetails"
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
func (r *OrganizationdetailService) ListAutoPaging(ctx context.Context, query OrganizationdetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OrganizationdetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an OrganizationDetails specified by the passed ID
// path parameter. OrganizationDetails represent details of organizations such as a
// corporation, manufacturer, consortium, government, etc. An organization can have
// detail records from several sources. A specific role is required to perform this
// service operation. Please contact the UDL team for assistance.
func (r *OrganizationdetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/organizationdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to get a single OrganizationDetails by a source passed as a
// query parameter. OrganizationDetails represent details of organizations such as
// a corporation, manufacturer, consortium, government, etc. An organization can
// have detail records from several sources.
func (r *OrganizationdetailService) FindBySource(ctx context.Context, query OrganizationdetailFindBySourceParams, opts ...option.RequestOption) (res *[]OrganizationdetailFindBySourceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/organizationdetails/findBySource"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single OrganizationDetails by its unique ID passed as
// a path parameter. OrganizationDetails represent details of organizations such as
// a corporation, manufacturer, consortium, government, etc. An organization can
// have detail records from several sources.
func (r *OrganizationdetailService) Get(ctx context.Context, id string, query OrganizationdetailGetParams, opts ...option.RequestOption) (res *OrganizationDetailsFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/organizationdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of additional detailed organization data as collected by a
// particular source.
type OrganizationDetailsFull struct {
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
	DataMode OrganizationDetailsFullDataMode `json:"dataMode,required"`
	// Unique identifier of the parent organization.
	IDOrganization string `json:"idOrganization,required"`
	// Organization details name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Street number of the organization.
	Address1 string `json:"address1"`
	// Field for additional organization address information such as PO Box and unit
	// number.
	Address2 string `json:"address2"`
	// Contains the third line of address information for an organization.
	Address3 string `json:"address3"`
	// Designated broker for this organization.
	Broker string `json:"broker"`
	// For organizations of type CORPORATION, the name of the Chief Executive Officer.
	Ceo string `json:"ceo"`
	// For organizations of type CORPORATION, the name of the Chief Financial Officer.
	Cfo string `json:"cfo"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// For organizations of type CORPORATION, the name of the Chief Technology Officer.
	Cto string `json:"cto"`
	// Organization description.
	Description string `json:"description"`
	// For organizations of type CORPORATION, the company EBITDA value as of
	// financialYearEndDate in US Dollars.
	Ebitda float64 `json:"ebitda"`
	// Listed contact email address for the organization.
	Email string `json:"email"`
	// For organizations of type CORPORATION, notes on company financials.
	FinancialNotes string `json:"financialNotes"`
	// For organizations of type CORPORATION, the effective financial year end date for
	// revenue, EBITDA, and profit values.
	FinancialYearEndDate time.Time `json:"financialYearEndDate" format:"date-time"`
	// Satellite fleet planning notes for this organization.
	FleetPlanNotes string `json:"fleetPlanNotes"`
	// Former organization ID (if this organization previously existed as another
	// organization).
	FormerOrgID string `json:"formerOrgId"`
	// Total number of FTEs in this organization.
	Ftes int64 `json:"ftes"`
	// Administrative boundaries of the first sub-national level. Level 1 is simply the
	// largest demarcation under whatever demarcation criteria has been determined by
	// the governing body. For example, this may be a state or province.
	GeoAdminLevel1 string `json:"geoAdminLevel1"`
	// Administrative boundaries of the second sub-national level. Level 2 is simply
	// the second largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a county or district.
	GeoAdminLevel2 string `json:"geoAdminLevel2"`
	// Administrative boundaries of the third sub-national level. Level 3 is simply the
	// third largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a city or township.
	GeoAdminLevel3 string `json:"geoAdminLevel3"`
	// Mass ranking for this organization.
	MassRanking int64 `json:"massRanking"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Parent organization ID of this organization if it is a child organization.
	ParentOrgID string `json:"parentOrgId"`
	// A postal code, such as PIN or ZIP Code, is a series of letters or digits or both
	// included in the postal address of the organization.
	PostalCode string `json:"postalCode"`
	// For organizations of type CORPORATION, total annual profit as of
	// financialYearEndDate in US Dollars.
	Profit float64 `json:"profit"`
	// For organizations of type CORPORATION, total annual revenue as of
	// financialYearEndDate in US Dollars.
	Revenue float64 `json:"revenue"`
	// Revenue ranking for this organization.
	RevenueRanking int64 `json:"revenueRanking"`
	// The name of the risk manager for the organization.
	RiskManager string `json:"riskManager"`
	// Notes on the services provided by the organization.
	ServicesNotes string `json:"servicesNotes"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOrganization        respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Address1              respjson.Field
		Address2              respjson.Field
		Address3              respjson.Field
		Broker                respjson.Field
		Ceo                   respjson.Field
		Cfo                   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Cto                   respjson.Field
		Description           respjson.Field
		Ebitda                respjson.Field
		Email                 respjson.Field
		FinancialNotes        respjson.Field
		FinancialYearEndDate  respjson.Field
		FleetPlanNotes        respjson.Field
		FormerOrgID           respjson.Field
		Ftes                  respjson.Field
		GeoAdminLevel1        respjson.Field
		GeoAdminLevel2        respjson.Field
		GeoAdminLevel3        respjson.Field
		MassRanking           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ParentOrgID           respjson.Field
		PostalCode            respjson.Field
		Profit                respjson.Field
		Revenue               respjson.Field
		RevenueRanking        respjson.Field
		RiskManager           respjson.Field
		ServicesNotes         respjson.Field
		Tags                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationDetailsFull) RawJSON() string { return r.JSON.raw }
func (r *OrganizationDetailsFull) UnmarshalJSON(data []byte) error {
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
type OrganizationDetailsFullDataMode string

const (
	OrganizationDetailsFullDataModeReal      OrganizationDetailsFullDataMode = "REAL"
	OrganizationDetailsFullDataModeTest      OrganizationDetailsFullDataMode = "TEST"
	OrganizationDetailsFullDataModeSimulated OrganizationDetailsFullDataMode = "SIMULATED"
	OrganizationDetailsFullDataModeExercise  OrganizationDetailsFullDataMode = "EXERCISE"
)

// Model representation of additional detailed organization data as collected by a
// particular source.
type OrganizationdetailListResponse struct {
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
	DataMode OrganizationdetailListResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the parent organization.
	IDOrganization string `json:"idOrganization,required"`
	// Organization details name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Street number of the organization.
	Address1 string `json:"address1"`
	// Field for additional organization address information such as PO Box and unit
	// number.
	Address2 string `json:"address2"`
	// Contains the third line of address information for an organization.
	Address3 string `json:"address3"`
	// Designated broker for this organization.
	Broker string `json:"broker"`
	// For organizations of type CORPORATION, the name of the Chief Executive Officer.
	Ceo string `json:"ceo"`
	// For organizations of type CORPORATION, the name of the Chief Financial Officer.
	Cfo string `json:"cfo"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// For organizations of type CORPORATION, the name of the Chief Technology Officer.
	Cto string `json:"cto"`
	// Organization description.
	Description string `json:"description"`
	// For organizations of type CORPORATION, the company EBITDA value as of
	// financialYearEndDate in US Dollars.
	Ebitda float64 `json:"ebitda"`
	// Listed contact email address for the organization.
	Email string `json:"email"`
	// For organizations of type CORPORATION, notes on company financials.
	FinancialNotes string `json:"financialNotes"`
	// For organizations of type CORPORATION, the effective financial year end date for
	// revenue, EBITDA, and profit values.
	FinancialYearEndDate time.Time `json:"financialYearEndDate" format:"date-time"`
	// Satellite fleet planning notes for this organization.
	FleetPlanNotes string `json:"fleetPlanNotes"`
	// Former organization ID (if this organization previously existed as another
	// organization).
	FormerOrgID string `json:"formerOrgId"`
	// Total number of FTEs in this organization.
	Ftes int64 `json:"ftes"`
	// Administrative boundaries of the first sub-national level. Level 1 is simply the
	// largest demarcation under whatever demarcation criteria has been determined by
	// the governing body. For example, this may be a state or province.
	GeoAdminLevel1 string `json:"geoAdminLevel1"`
	// Administrative boundaries of the second sub-national level. Level 2 is simply
	// the second largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a county or district.
	GeoAdminLevel2 string `json:"geoAdminLevel2"`
	// Administrative boundaries of the third sub-national level. Level 3 is simply the
	// third largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a city or township.
	GeoAdminLevel3 string `json:"geoAdminLevel3"`
	// Mass ranking for this organization.
	MassRanking int64 `json:"massRanking"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Parent organization ID of this organization if it is a child organization.
	ParentOrgID string `json:"parentOrgId"`
	// A postal code, such as PIN or ZIP Code, is a series of letters or digits or both
	// included in the postal address of the organization.
	PostalCode string `json:"postalCode"`
	// For organizations of type CORPORATION, total annual profit as of
	// financialYearEndDate in US Dollars.
	Profit float64 `json:"profit"`
	// For organizations of type CORPORATION, total annual revenue as of
	// financialYearEndDate in US Dollars.
	Revenue float64 `json:"revenue"`
	// Revenue ranking for this organization.
	RevenueRanking int64 `json:"revenueRanking"`
	// The name of the risk manager for the organization.
	RiskManager string `json:"riskManager"`
	// Notes on the services provided by the organization.
	ServicesNotes string `json:"servicesNotes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOrganization        respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Address1              respjson.Field
		Address2              respjson.Field
		Address3              respjson.Field
		Broker                respjson.Field
		Ceo                   respjson.Field
		Cfo                   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Cto                   respjson.Field
		Description           respjson.Field
		Ebitda                respjson.Field
		Email                 respjson.Field
		FinancialNotes        respjson.Field
		FinancialYearEndDate  respjson.Field
		FleetPlanNotes        respjson.Field
		FormerOrgID           respjson.Field
		Ftes                  respjson.Field
		GeoAdminLevel1        respjson.Field
		GeoAdminLevel2        respjson.Field
		GeoAdminLevel3        respjson.Field
		MassRanking           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ParentOrgID           respjson.Field
		PostalCode            respjson.Field
		Profit                respjson.Field
		Revenue               respjson.Field
		RevenueRanking        respjson.Field
		RiskManager           respjson.Field
		ServicesNotes         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationdetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationdetailListResponse) UnmarshalJSON(data []byte) error {
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
type OrganizationdetailListResponseDataMode string

const (
	OrganizationdetailListResponseDataModeReal      OrganizationdetailListResponseDataMode = "REAL"
	OrganizationdetailListResponseDataModeTest      OrganizationdetailListResponseDataMode = "TEST"
	OrganizationdetailListResponseDataModeSimulated OrganizationdetailListResponseDataMode = "SIMULATED"
	OrganizationdetailListResponseDataModeExercise  OrganizationdetailListResponseDataMode = "EXERCISE"
)

// Model representation of additional detailed organization data as collected by a
// particular source.
type OrganizationdetailFindBySourceResponse struct {
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
	DataMode OrganizationdetailFindBySourceResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the parent organization.
	IDOrganization string `json:"idOrganization,required"`
	// Organization details name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Street number of the organization.
	Address1 string `json:"address1"`
	// Field for additional organization address information such as PO Box and unit
	// number.
	Address2 string `json:"address2"`
	// Contains the third line of address information for an organization.
	Address3 string `json:"address3"`
	// Designated broker for this organization.
	Broker string `json:"broker"`
	// For organizations of type CORPORATION, the name of the Chief Executive Officer.
	Ceo string `json:"ceo"`
	// For organizations of type CORPORATION, the name of the Chief Financial Officer.
	Cfo string `json:"cfo"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// For organizations of type CORPORATION, the name of the Chief Technology Officer.
	Cto string `json:"cto"`
	// Organization description.
	Description string `json:"description"`
	// For organizations of type CORPORATION, the company EBITDA value as of
	// financialYearEndDate in US Dollars.
	Ebitda float64 `json:"ebitda"`
	// Listed contact email address for the organization.
	Email string `json:"email"`
	// For organizations of type CORPORATION, notes on company financials.
	FinancialNotes string `json:"financialNotes"`
	// For organizations of type CORPORATION, the effective financial year end date for
	// revenue, EBITDA, and profit values.
	FinancialYearEndDate time.Time `json:"financialYearEndDate" format:"date-time"`
	// Satellite fleet planning notes for this organization.
	FleetPlanNotes string `json:"fleetPlanNotes"`
	// Former organization ID (if this organization previously existed as another
	// organization).
	FormerOrgID string `json:"formerOrgId"`
	// Total number of FTEs in this organization.
	Ftes int64 `json:"ftes"`
	// Administrative boundaries of the first sub-national level. Level 1 is simply the
	// largest demarcation under whatever demarcation criteria has been determined by
	// the governing body. For example, this may be a state or province.
	GeoAdminLevel1 string `json:"geoAdminLevel1"`
	// Administrative boundaries of the second sub-national level. Level 2 is simply
	// the second largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a county or district.
	GeoAdminLevel2 string `json:"geoAdminLevel2"`
	// Administrative boundaries of the third sub-national level. Level 3 is simply the
	// third largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a city or township.
	GeoAdminLevel3 string `json:"geoAdminLevel3"`
	// Mass ranking for this organization.
	MassRanking int64 `json:"massRanking"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Parent organization ID of this organization if it is a child organization.
	ParentOrgID string `json:"parentOrgId"`
	// A postal code, such as PIN or ZIP Code, is a series of letters or digits or both
	// included in the postal address of the organization.
	PostalCode string `json:"postalCode"`
	// For organizations of type CORPORATION, total annual profit as of
	// financialYearEndDate in US Dollars.
	Profit float64 `json:"profit"`
	// For organizations of type CORPORATION, total annual revenue as of
	// financialYearEndDate in US Dollars.
	Revenue float64 `json:"revenue"`
	// Revenue ranking for this organization.
	RevenueRanking int64 `json:"revenueRanking"`
	// The name of the risk manager for the organization.
	RiskManager string `json:"riskManager"`
	// Notes on the services provided by the organization.
	ServicesNotes string `json:"servicesNotes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOrganization        respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Address1              respjson.Field
		Address2              respjson.Field
		Address3              respjson.Field
		Broker                respjson.Field
		Ceo                   respjson.Field
		Cfo                   respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Cto                   respjson.Field
		Description           respjson.Field
		Ebitda                respjson.Field
		Email                 respjson.Field
		FinancialNotes        respjson.Field
		FinancialYearEndDate  respjson.Field
		FleetPlanNotes        respjson.Field
		FormerOrgID           respjson.Field
		Ftes                  respjson.Field
		GeoAdminLevel1        respjson.Field
		GeoAdminLevel2        respjson.Field
		GeoAdminLevel3        respjson.Field
		MassRanking           respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ParentOrgID           respjson.Field
		PostalCode            respjson.Field
		Profit                respjson.Field
		Revenue               respjson.Field
		RevenueRanking        respjson.Field
		RiskManager           respjson.Field
		ServicesNotes         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationdetailFindBySourceResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationdetailFindBySourceResponse) UnmarshalJSON(data []byte) error {
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
type OrganizationdetailFindBySourceResponseDataMode string

const (
	OrganizationdetailFindBySourceResponseDataModeReal      OrganizationdetailFindBySourceResponseDataMode = "REAL"
	OrganizationdetailFindBySourceResponseDataModeTest      OrganizationdetailFindBySourceResponseDataMode = "TEST"
	OrganizationdetailFindBySourceResponseDataModeSimulated OrganizationdetailFindBySourceResponseDataMode = "SIMULATED"
	OrganizationdetailFindBySourceResponseDataModeExercise  OrganizationdetailFindBySourceResponseDataMode = "EXERCISE"
)

type OrganizationdetailNewParams struct {
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
	DataMode OrganizationdetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent organization.
	IDOrganization string `json:"idOrganization,required"`
	// Organization details name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Street number of the organization.
	Address1 param.Opt[string] `json:"address1,omitzero"`
	// Field for additional organization address information such as PO Box and unit
	// number.
	Address2 param.Opt[string] `json:"address2,omitzero"`
	// Contains the third line of address information for an organization.
	Address3 param.Opt[string] `json:"address3,omitzero"`
	// Designated broker for this organization.
	Broker param.Opt[string] `json:"broker,omitzero"`
	// For organizations of type CORPORATION, the name of the Chief Executive Officer.
	Ceo param.Opt[string] `json:"ceo,omitzero"`
	// For organizations of type CORPORATION, the name of the Chief Financial Officer.
	Cfo param.Opt[string] `json:"cfo,omitzero"`
	// For organizations of type CORPORATION, the name of the Chief Technology Officer.
	Cto param.Opt[string] `json:"cto,omitzero"`
	// Organization description.
	Description param.Opt[string] `json:"description,omitzero"`
	// For organizations of type CORPORATION, the company EBITDA value as of
	// financialYearEndDate in US Dollars.
	Ebitda param.Opt[float64] `json:"ebitda,omitzero"`
	// Listed contact email address for the organization.
	Email param.Opt[string] `json:"email,omitzero"`
	// For organizations of type CORPORATION, notes on company financials.
	FinancialNotes param.Opt[string] `json:"financialNotes,omitzero"`
	// For organizations of type CORPORATION, the effective financial year end date for
	// revenue, EBITDA, and profit values.
	FinancialYearEndDate param.Opt[time.Time] `json:"financialYearEndDate,omitzero" format:"date-time"`
	// Satellite fleet planning notes for this organization.
	FleetPlanNotes param.Opt[string] `json:"fleetPlanNotes,omitzero"`
	// Former organization ID (if this organization previously existed as another
	// organization).
	FormerOrgID param.Opt[string] `json:"formerOrgId,omitzero"`
	// Total number of FTEs in this organization.
	Ftes param.Opt[int64] `json:"ftes,omitzero"`
	// Administrative boundaries of the first sub-national level. Level 1 is simply the
	// largest demarcation under whatever demarcation criteria has been determined by
	// the governing body. For example, this may be a state or province.
	GeoAdminLevel1 param.Opt[string] `json:"geoAdminLevel1,omitzero"`
	// Administrative boundaries of the second sub-national level. Level 2 is simply
	// the second largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a county or district.
	GeoAdminLevel2 param.Opt[string] `json:"geoAdminLevel2,omitzero"`
	// Administrative boundaries of the third sub-national level. Level 3 is simply the
	// third largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a city or township.
	GeoAdminLevel3 param.Opt[string] `json:"geoAdminLevel3,omitzero"`
	// Mass ranking for this organization.
	MassRanking param.Opt[int64] `json:"massRanking,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Parent organization ID of this organization if it is a child organization.
	ParentOrgID param.Opt[string] `json:"parentOrgId,omitzero"`
	// A postal code, such as PIN or ZIP Code, is a series of letters or digits or both
	// included in the postal address of the organization.
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// For organizations of type CORPORATION, total annual profit as of
	// financialYearEndDate in US Dollars.
	Profit param.Opt[float64] `json:"profit,omitzero"`
	// For organizations of type CORPORATION, total annual revenue as of
	// financialYearEndDate in US Dollars.
	Revenue param.Opt[float64] `json:"revenue,omitzero"`
	// Revenue ranking for this organization.
	RevenueRanking param.Opt[int64] `json:"revenueRanking,omitzero"`
	// The name of the risk manager for the organization.
	RiskManager param.Opt[string] `json:"riskManager,omitzero"`
	// Notes on the services provided by the organization.
	ServicesNotes param.Opt[string] `json:"servicesNotes,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r OrganizationdetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationdetailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationdetailNewParams) UnmarshalJSON(data []byte) error {
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
type OrganizationdetailNewParamsDataMode string

const (
	OrganizationdetailNewParamsDataModeReal      OrganizationdetailNewParamsDataMode = "REAL"
	OrganizationdetailNewParamsDataModeTest      OrganizationdetailNewParamsDataMode = "TEST"
	OrganizationdetailNewParamsDataModeSimulated OrganizationdetailNewParamsDataMode = "SIMULATED"
	OrganizationdetailNewParamsDataModeExercise  OrganizationdetailNewParamsDataMode = "EXERCISE"
)

type OrganizationdetailUpdateParams struct {
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
	DataMode OrganizationdetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent organization.
	IDOrganization string `json:"idOrganization,required"`
	// Organization details name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Street number of the organization.
	Address1 param.Opt[string] `json:"address1,omitzero"`
	// Field for additional organization address information such as PO Box and unit
	// number.
	Address2 param.Opt[string] `json:"address2,omitzero"`
	// Contains the third line of address information for an organization.
	Address3 param.Opt[string] `json:"address3,omitzero"`
	// Designated broker for this organization.
	Broker param.Opt[string] `json:"broker,omitzero"`
	// For organizations of type CORPORATION, the name of the Chief Executive Officer.
	Ceo param.Opt[string] `json:"ceo,omitzero"`
	// For organizations of type CORPORATION, the name of the Chief Financial Officer.
	Cfo param.Opt[string] `json:"cfo,omitzero"`
	// For organizations of type CORPORATION, the name of the Chief Technology Officer.
	Cto param.Opt[string] `json:"cto,omitzero"`
	// Organization description.
	Description param.Opt[string] `json:"description,omitzero"`
	// For organizations of type CORPORATION, the company EBITDA value as of
	// financialYearEndDate in US Dollars.
	Ebitda param.Opt[float64] `json:"ebitda,omitzero"`
	// Listed contact email address for the organization.
	Email param.Opt[string] `json:"email,omitzero"`
	// For organizations of type CORPORATION, notes on company financials.
	FinancialNotes param.Opt[string] `json:"financialNotes,omitzero"`
	// For organizations of type CORPORATION, the effective financial year end date for
	// revenue, EBITDA, and profit values.
	FinancialYearEndDate param.Opt[time.Time] `json:"financialYearEndDate,omitzero" format:"date-time"`
	// Satellite fleet planning notes for this organization.
	FleetPlanNotes param.Opt[string] `json:"fleetPlanNotes,omitzero"`
	// Former organization ID (if this organization previously existed as another
	// organization).
	FormerOrgID param.Opt[string] `json:"formerOrgId,omitzero"`
	// Total number of FTEs in this organization.
	Ftes param.Opt[int64] `json:"ftes,omitzero"`
	// Administrative boundaries of the first sub-national level. Level 1 is simply the
	// largest demarcation under whatever demarcation criteria has been determined by
	// the governing body. For example, this may be a state or province.
	GeoAdminLevel1 param.Opt[string] `json:"geoAdminLevel1,omitzero"`
	// Administrative boundaries of the second sub-national level. Level 2 is simply
	// the second largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a county or district.
	GeoAdminLevel2 param.Opt[string] `json:"geoAdminLevel2,omitzero"`
	// Administrative boundaries of the third sub-national level. Level 3 is simply the
	// third largest demarcation under whatever demarcation criteria has been
	// determined by the governing body. For example, this may be a city or township.
	GeoAdminLevel3 param.Opt[string] `json:"geoAdminLevel3,omitzero"`
	// Mass ranking for this organization.
	MassRanking param.Opt[int64] `json:"massRanking,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Parent organization ID of this organization if it is a child organization.
	ParentOrgID param.Opt[string] `json:"parentOrgId,omitzero"`
	// A postal code, such as PIN or ZIP Code, is a series of letters or digits or both
	// included in the postal address of the organization.
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// For organizations of type CORPORATION, total annual profit as of
	// financialYearEndDate in US Dollars.
	Profit param.Opt[float64] `json:"profit,omitzero"`
	// For organizations of type CORPORATION, total annual revenue as of
	// financialYearEndDate in US Dollars.
	Revenue param.Opt[float64] `json:"revenue,omitzero"`
	// Revenue ranking for this organization.
	RevenueRanking param.Opt[int64] `json:"revenueRanking,omitzero"`
	// The name of the risk manager for the organization.
	RiskManager param.Opt[string] `json:"riskManager,omitzero"`
	// Notes on the services provided by the organization.
	ServicesNotes param.Opt[string] `json:"servicesNotes,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r OrganizationdetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationdetailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationdetailUpdateParams) UnmarshalJSON(data []byte) error {
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
type OrganizationdetailUpdateParamsDataMode string

const (
	OrganizationdetailUpdateParamsDataModeReal      OrganizationdetailUpdateParamsDataMode = "REAL"
	OrganizationdetailUpdateParamsDataModeTest      OrganizationdetailUpdateParamsDataMode = "TEST"
	OrganizationdetailUpdateParamsDataModeSimulated OrganizationdetailUpdateParamsDataMode = "SIMULATED"
	OrganizationdetailUpdateParamsDataModeExercise  OrganizationdetailUpdateParamsDataMode = "EXERCISE"
)

type OrganizationdetailListParams struct {
	// Organization details name.
	Name        string           `query:"name,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationdetailListParams]'s query parameters as
// `url.Values`.
func (r OrganizationdetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationdetailFindBySourceParams struct {
	// Organization details name.
	Name string `query:"name,required" json:"-"`
	// The source of the OrganizationDetails records to find.
	Source      string           `query:"source,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationdetailFindBySourceParams]'s query parameters as
// `url.Values`.
func (r OrganizationdetailFindBySourceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationdetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationdetailGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationdetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
