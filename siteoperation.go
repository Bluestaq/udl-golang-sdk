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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
)

// SiteOperationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteOperationService] method instead.
type SiteOperationService struct {
	Options []option.RequestOption
}

// NewSiteOperationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSiteOperationService(opts ...option.RequestOption) (r SiteOperationService) {
	r = SiteOperationService{}
	r.Options = opts
	return
}

// Service operation to take a single siteoperations object as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SiteOperationService) New(ctx context.Context, body SiteOperationNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/siteoperations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single siteoperations record by its unique ID passed
// as a path parameter.
func (r *SiteOperationService) Get(ctx context.Context, id string, query SiteOperationGetParams, opts ...option.RequestOption) (res *SiteOperationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/siteoperations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single siteoperations record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SiteOperationService) Update(ctx context.Context, id string, body SiteOperationUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/siteoperations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SiteOperationService) List(ctx context.Context, query SiteOperationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SiteOperationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/siteoperations"
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
func (r *SiteOperationService) ListAutoPaging(ctx context.Context, query SiteOperationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SiteOperationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a siteoperations record specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SiteOperationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/siteoperations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SiteOperationService) Count(ctx context.Context, query SiteOperationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/siteoperations/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// siteoperations records as a POST body and ingest into the database. This
// operation is not intended to be used for automated feeds into UDL. Data
// providers should contact the UDL team for specific role assignments and for
// instructions on setting up a permanent feed through an alternate mechanism.
func (r *SiteOperationService) NewBulk(ctx context.Context, body SiteOperationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/siteoperations/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SiteOperationService) QueryHelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/siteoperations/queryhelp"
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
func (r *SiteOperationService) Tuple(ctx context.Context, query SiteOperationTupleParams, opts ...option.RequestOption) (res *[]SiteOperationTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/siteoperations/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple siteoperations records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SiteOperationService) UnvalidatedPublish(ctx context.Context, body SiteOperationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-siteoperations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Site operating details concerning the hours of operation, operational
// limitations, site navigation, and waivers associated with the Site.
type SiteOperationGetResponse struct {
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
	DataMode SiteOperationGetResponseDataMode `json:"dataMode,required"`
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
	DailyOperations []SiteOperationGetResponseDailyOperation `json:"dailyOperations"`
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
	MaximumOnGrounds []SiteOperationGetResponseMaximumOnGround `json:"maximumOnGrounds"`
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
	OperationalDeviations []SiteOperationGetResponseOperationalDeviation `json:"operationalDeviations"`
	// Collection of planning information associated with this SiteOperations record.
	OperationalPlannings []SiteOperationGetResponseOperationalPlanning `json:"operationalPlannings"`
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
	Pathways []SiteOperationGetResponsePathway `json:"pathways"`
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
	Waivers []SiteOperationGetResponseWaiver `json:"waivers"`
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
func (r SiteOperationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationGetResponse) UnmarshalJSON(data []byte) error {
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
type SiteOperationGetResponseDataMode string

const (
	SiteOperationGetResponseDataModeReal      SiteOperationGetResponseDataMode = "REAL"
	SiteOperationGetResponseDataModeTest      SiteOperationGetResponseDataMode = "TEST"
	SiteOperationGetResponseDataModeSimulated SiteOperationGetResponseDataMode = "SIMULATED"
	SiteOperationGetResponseDataModeExercise  SiteOperationGetResponseDataMode = "EXERCISE"
)

// Collection providing hours of operation and other information specific to a day
// of the week.
type SiteOperationGetResponseDailyOperation struct {
	// The day of the week to which this operational information pertains.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DayOfWeek string `json:"dayOfWeek"`
	// A collection containing the operational start and stop times scheduled for the
	// day of the week specified.
	OperatingHours []SiteOperationGetResponseDailyOperationOperatingHour `json:"operatingHours"`
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
func (r SiteOperationGetResponseDailyOperation) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationGetResponseDailyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection containing the operational start and stop times scheduled for the
// day of the week specified.
type SiteOperationGetResponseDailyOperationOperatingHour struct {
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
func (r SiteOperationGetResponseDailyOperationOperatingHour) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationGetResponseDailyOperationOperatingHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing maximum on ground (MOG) information for specific aircraft
// at the site associated with this SiteOperations record.
type SiteOperationGetResponseMaximumOnGround struct {
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
func (r SiteOperationGetResponseMaximumOnGround) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationGetResponseMaximumOnGround) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing relevant information in the event of deviations/exceptions
// to normal operations.
type SiteOperationGetResponseOperationalDeviation struct {
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
func (r SiteOperationGetResponseOperationalDeviation) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationGetResponseOperationalDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of planning information associated with this SiteOperations record.
type SiteOperationGetResponseOperationalPlanning struct {
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
func (r SiteOperationGetResponseOperationalPlanning) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationGetResponseOperationalPlanning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection detailing operational pathways at the Site associated with this
// SiteOperations record.
type SiteOperationGetResponsePathway struct {
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
func (r SiteOperationGetResponsePathway) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationGetResponsePathway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection documenting operational waivers that have been issued for the Site
// associated with this record.
type SiteOperationGetResponseWaiver struct {
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
func (r SiteOperationGetResponseWaiver) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationGetResponseWaiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Site operating details concerning the hours of operation, operational
// limitations, site navigation, and waivers associated with the Site.
type SiteOperationListResponse struct {
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
	DataMode SiteOperationListResponseDataMode `json:"dataMode,required"`
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
	DailyOperations []SiteOperationListResponseDailyOperation `json:"dailyOperations"`
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
	MaximumOnGrounds []SiteOperationListResponseMaximumOnGround `json:"maximumOnGrounds"`
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
	OperationalDeviations []SiteOperationListResponseOperationalDeviation `json:"operationalDeviations"`
	// Collection of planning information associated with this SiteOperations record.
	OperationalPlannings []SiteOperationListResponseOperationalPlanning `json:"operationalPlannings"`
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
	Pathways []SiteOperationListResponsePathway `json:"pathways"`
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
	Waivers []SiteOperationListResponseWaiver `json:"waivers"`
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
func (r SiteOperationListResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationListResponse) UnmarshalJSON(data []byte) error {
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
type SiteOperationListResponseDataMode string

const (
	SiteOperationListResponseDataModeReal      SiteOperationListResponseDataMode = "REAL"
	SiteOperationListResponseDataModeTest      SiteOperationListResponseDataMode = "TEST"
	SiteOperationListResponseDataModeSimulated SiteOperationListResponseDataMode = "SIMULATED"
	SiteOperationListResponseDataModeExercise  SiteOperationListResponseDataMode = "EXERCISE"
)

// Collection providing hours of operation and other information specific to a day
// of the week.
type SiteOperationListResponseDailyOperation struct {
	// The day of the week to which this operational information pertains.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DayOfWeek string `json:"dayOfWeek"`
	// A collection containing the operational start and stop times scheduled for the
	// day of the week specified.
	OperatingHours []SiteOperationListResponseDailyOperationOperatingHour `json:"operatingHours"`
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
func (r SiteOperationListResponseDailyOperation) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationListResponseDailyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection containing the operational start and stop times scheduled for the
// day of the week specified.
type SiteOperationListResponseDailyOperationOperatingHour struct {
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
func (r SiteOperationListResponseDailyOperationOperatingHour) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationListResponseDailyOperationOperatingHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing maximum on ground (MOG) information for specific aircraft
// at the site associated with this SiteOperations record.
type SiteOperationListResponseMaximumOnGround struct {
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
func (r SiteOperationListResponseMaximumOnGround) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationListResponseMaximumOnGround) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing relevant information in the event of deviations/exceptions
// to normal operations.
type SiteOperationListResponseOperationalDeviation struct {
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
func (r SiteOperationListResponseOperationalDeviation) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationListResponseOperationalDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of planning information associated with this SiteOperations record.
type SiteOperationListResponseOperationalPlanning struct {
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
func (r SiteOperationListResponseOperationalPlanning) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationListResponseOperationalPlanning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection detailing operational pathways at the Site associated with this
// SiteOperations record.
type SiteOperationListResponsePathway struct {
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
func (r SiteOperationListResponsePathway) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationListResponsePathway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection documenting operational waivers that have been issued for the Site
// associated with this record.
type SiteOperationListResponseWaiver struct {
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
func (r SiteOperationListResponseWaiver) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationListResponseWaiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Site operating details concerning the hours of operation, operational
// limitations, site navigation, and waivers associated with the Site.
type SiteOperationTupleResponse struct {
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
	DataMode SiteOperationTupleResponseDataMode `json:"dataMode,required"`
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
	DailyOperations []SiteOperationTupleResponseDailyOperation `json:"dailyOperations"`
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
	MaximumOnGrounds []SiteOperationTupleResponseMaximumOnGround `json:"maximumOnGrounds"`
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
	OperationalDeviations []SiteOperationTupleResponseOperationalDeviation `json:"operationalDeviations"`
	// Collection of planning information associated with this SiteOperations record.
	OperationalPlannings []SiteOperationTupleResponseOperationalPlanning `json:"operationalPlannings"`
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
	Pathways []SiteOperationTupleResponsePathway `json:"pathways"`
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
	Waivers []SiteOperationTupleResponseWaiver `json:"waivers"`
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
func (r SiteOperationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationTupleResponse) UnmarshalJSON(data []byte) error {
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
type SiteOperationTupleResponseDataMode string

const (
	SiteOperationTupleResponseDataModeReal      SiteOperationTupleResponseDataMode = "REAL"
	SiteOperationTupleResponseDataModeTest      SiteOperationTupleResponseDataMode = "TEST"
	SiteOperationTupleResponseDataModeSimulated SiteOperationTupleResponseDataMode = "SIMULATED"
	SiteOperationTupleResponseDataModeExercise  SiteOperationTupleResponseDataMode = "EXERCISE"
)

// Collection providing hours of operation and other information specific to a day
// of the week.
type SiteOperationTupleResponseDailyOperation struct {
	// The day of the week to which this operational information pertains.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DayOfWeek string `json:"dayOfWeek"`
	// A collection containing the operational start and stop times scheduled for the
	// day of the week specified.
	OperatingHours []SiteOperationTupleResponseDailyOperationOperatingHour `json:"operatingHours"`
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
func (r SiteOperationTupleResponseDailyOperation) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationTupleResponseDailyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection containing the operational start and stop times scheduled for the
// day of the week specified.
type SiteOperationTupleResponseDailyOperationOperatingHour struct {
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
func (r SiteOperationTupleResponseDailyOperationOperatingHour) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationTupleResponseDailyOperationOperatingHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing maximum on ground (MOG) information for specific aircraft
// at the site associated with this SiteOperations record.
type SiteOperationTupleResponseMaximumOnGround struct {
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
func (r SiteOperationTupleResponseMaximumOnGround) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationTupleResponseMaximumOnGround) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing relevant information in the event of deviations/exceptions
// to normal operations.
type SiteOperationTupleResponseOperationalDeviation struct {
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
func (r SiteOperationTupleResponseOperationalDeviation) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationTupleResponseOperationalDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of planning information associated with this SiteOperations record.
type SiteOperationTupleResponseOperationalPlanning struct {
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
func (r SiteOperationTupleResponseOperationalPlanning) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationTupleResponseOperationalPlanning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection detailing operational pathways at the Site associated with this
// SiteOperations record.
type SiteOperationTupleResponsePathway struct {
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
func (r SiteOperationTupleResponsePathway) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationTupleResponsePathway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection documenting operational waivers that have been issued for the Site
// associated with this record.
type SiteOperationTupleResponseWaiver struct {
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
func (r SiteOperationTupleResponseWaiver) RawJSON() string { return r.JSON.raw }
func (r *SiteOperationTupleResponseWaiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteOperationNewParams struct {
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
	DataMode SiteOperationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The ID of the parent site.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The name of the person who made the most recent change to data in the
	// DailyOperations collection.
	DopsLastChangedBy param.Opt[string] `json:"dopsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to data in the DailyOperations
	// collection, in ISO 8601 UTC format with millisecond precision.
	DopsLastChangedDate param.Opt[time.Time] `json:"dopsLastChangedDate,omitzero" format:"date-time"`
	// The reason for the most recent change to data in the dailyOperations collection.
	DopsLastChangedReason param.Opt[string] `json:"dopsLastChangedReason,omitzero"`
	// Id of the associated launchSite entity.
	IDLaunchSite param.Opt[string] `json:"idLaunchSite,omitzero"`
	// The name of the person who made the most recent change to data in the
	// MaximumOnGrounds collection.
	MogsLastChangedBy param.Opt[string] `json:"mogsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to data in the MaximumOnGrounds
	// collection, in ISO 8601 UTC format with millisecond precision.
	MogsLastChangedDate param.Opt[time.Time] `json:"mogsLastChangedDate,omitzero" format:"date-time"`
	// The reason for the most recent change to data in the MaximumOnGrounds
	// collection.
	MogsLastChangedReason param.Opt[string] `json:"mogsLastChangedReason,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Collection providing hours of operation and other information specific to a day
	// of the week.
	DailyOperations []SiteOperationNewParamsDailyOperation `json:"dailyOperations,omitzero"`
	// Collection providing maximum on ground (MOG) information for specific aircraft
	// at the site associated with this SiteOperations record.
	MaximumOnGrounds []SiteOperationNewParamsMaximumOnGround `json:"maximumOnGrounds,omitzero"`
	// Collection providing relevant information in the event of deviations/exceptions
	// to normal operations.
	OperationalDeviations []SiteOperationNewParamsOperationalDeviation `json:"operationalDeviations,omitzero"`
	// Collection of planning information associated with this SiteOperations record.
	OperationalPlannings []SiteOperationNewParamsOperationalPlanning `json:"operationalPlannings,omitzero"`
	// Collection detailing operational pathways at the Site associated with this
	// SiteOperations record.
	Pathways []SiteOperationNewParamsPathway `json:"pathways,omitzero"`
	// Collection documenting operational waivers that have been issued for the Site
	// associated with this record.
	Waivers []SiteOperationNewParamsWaiver `json:"waivers,omitzero"`
	paramObj
}

func (r SiteOperationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewParams) UnmarshalJSON(data []byte) error {
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
type SiteOperationNewParamsDataMode string

const (
	SiteOperationNewParamsDataModeReal      SiteOperationNewParamsDataMode = "REAL"
	SiteOperationNewParamsDataModeTest      SiteOperationNewParamsDataMode = "TEST"
	SiteOperationNewParamsDataModeSimulated SiteOperationNewParamsDataMode = "SIMULATED"
	SiteOperationNewParamsDataModeExercise  SiteOperationNewParamsDataMode = "EXERCISE"
)

// Collection providing hours of operation and other information specific to a day
// of the week.
type SiteOperationNewParamsDailyOperation struct {
	// The name or type of operation to which this information pertains.
	OperationName param.Opt[string] `json:"operationName,omitzero"`
	// The name of the person who made the most recent change to this DailyOperation
	// data.
	OphrsLastChangedBy param.Opt[string] `json:"ophrsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this DailyOperation data, in ISO
	// 8601 UTC format with millisecond precision.
	OphrsLastChangedDate param.Opt[time.Time] `json:"ophrsLastChangedDate,omitzero" format:"date-time"`
	// The day of the week to which this operational information pertains.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DayOfWeek string `json:"dayOfWeek,omitzero"`
	// A collection containing the operational start and stop times scheduled for the
	// day of the week specified.
	OperatingHours []SiteOperationNewParamsDailyOperationOperatingHour `json:"operatingHours,omitzero"`
	paramObj
}

func (r SiteOperationNewParamsDailyOperation) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewParamsDailyOperation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewParamsDailyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SiteOperationNewParamsDailyOperation](
		"dayOfWeek", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY",
	)
}

// A collection containing the operational start and stop times scheduled for the
// day of the week specified.
type SiteOperationNewParamsDailyOperationOperatingHour struct {
	// The Zulu (UTC) operational start time, expressed in ISO 8601 format as HH:MM.
	OpStartTime param.Opt[string] `json:"opStartTime,omitzero"`
	// The Zulu (UTC) operational stop time, expressed in ISO 8601 format as HH:MM.
	OpStopTime param.Opt[string] `json:"opStopTime,omitzero"`
	paramObj
}

func (r SiteOperationNewParamsDailyOperationOperatingHour) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewParamsDailyOperationOperatingHour
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewParamsDailyOperationOperatingHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing maximum on ground (MOG) information for specific aircraft
// at the site associated with this SiteOperations record.
type SiteOperationNewParamsMaximumOnGround struct {
	// The Model Design Series (MDS) designation of the aircraft to which this maximum
	// on ground (MOG) data pertains.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Maximum on ground (MOG) number of contingent aircraft based on spacing and
	// manpower, for the aircraft type specified.
	ContingencyMog param.Opt[int64] `json:"contingencyMOG,omitzero"`
	// The name of the person who made the most recent change to this maximum on ground
	// data.
	MogLastChangedBy param.Opt[string] `json:"mogLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this maximum on ground data, in
	// ISO 8601 UTC format with millisecond precision.
	MogLastChangedDate param.Opt[time.Time] `json:"mogLastChangedDate,omitzero" format:"date-time"`
	// Maximum on ground (MOG) number of parking wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideParkingMog param.Opt[int64] `json:"wideParkingMOG,omitzero"`
	// Maximum on ground (MOG) number of working wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideWorkingMog param.Opt[int64] `json:"wideWorkingMOG,omitzero"`
	paramObj
}

func (r SiteOperationNewParamsMaximumOnGround) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewParamsMaximumOnGround
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewParamsMaximumOnGround) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing relevant information in the event of deviations/exceptions
// to normal operations.
type SiteOperationNewParamsOperationalDeviation struct {
	// The Model Design Series (MDS) designation of the aircraft affected by this
	// operational deviation.
	AffectedAircraftMds param.Opt[string] `json:"affectedAircraftMDS,omitzero"`
	// The maximum on ground (MOG) number for aircraft affected by this operational
	// deviation.
	AffectedMog param.Opt[int64] `json:"affectedMOG,omitzero"`
	// On ground time for aircraft affected by this operational deviation.
	AircraftOnGroundTime param.Opt[string] `json:"aircraftOnGroundTime,omitzero"`
	// Rest time for crew affected by this operational deviation.
	CrewRestTime param.Opt[string] `json:"crewRestTime,omitzero"`
	// The name of the person who made the most recent change to this
	// OperationalDeviation data.
	OdLastChangedBy param.Opt[string] `json:"odLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this OperationalDeviation data,
	// in ISO 8601 UTC format with millisecond precision.
	OdLastChangedDate param.Opt[time.Time] `json:"odLastChangedDate,omitzero" format:"date-time"`
	// Text remark regarding this operational deviation.
	OdRemark param.Opt[string] `json:"odRemark,omitzero"`
	paramObj
}

func (r SiteOperationNewParamsOperationalDeviation) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewParamsOperationalDeviation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewParamsOperationalDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of planning information associated with this SiteOperations record.
type SiteOperationNewParamsOperationalPlanning struct {
	// The end date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpEndDate param.Opt[time.Time] `json:"opEndDate,omitzero" format:"date-time"`
	// The name of the person who made the most recent change made to this
	// OperationalPlanning data.
	OpLastChangedBy param.Opt[string] `json:"opLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this OperationalPlanning data, in
	// ISO8601 UTC format with millisecond precision.
	OpLastChangedDate param.Opt[time.Time] `json:"opLastChangedDate,omitzero" format:"date-time"`
	// Remark text regarding this operation planning.
	OpRemark param.Opt[string] `json:"opRemark,omitzero"`
	// The person, unit, organization, etc. responsible for this operation planning.
	OpSource param.Opt[string] `json:"opSource,omitzero"`
	// The start date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpStartDate param.Opt[time.Time] `json:"opStartDate,omitzero" format:"date-time"`
	// The status of this operational planning.
	OpStatus param.Opt[string] `json:"opStatus,omitzero"`
	paramObj
}

func (r SiteOperationNewParamsOperationalPlanning) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewParamsOperationalPlanning
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewParamsOperationalPlanning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection detailing operational pathways at the Site associated with this
// SiteOperations record.
type SiteOperationNewParamsPathway struct {
	// Text defining this pathway from its constituent parts.
	PwDefinition param.Opt[string] `json:"pwDefinition,omitzero"`
	// The name of the person who made the most recent change to this Pathway data.
	PwLastChangedBy param.Opt[string] `json:"pwLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this Pathway data, in ISO 8601
	// UTC format with millisecond precision.
	PwLastChangedDate param.Opt[time.Time] `json:"pwLastChangedDate,omitzero" format:"date-time"`
	// The type of paths that constitute this pathway.
	PwType param.Opt[string] `json:"pwType,omitzero"`
	// The intended use of this pathway.
	PwUsage param.Opt[string] `json:"pwUsage,omitzero"`
	paramObj
}

func (r SiteOperationNewParamsPathway) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewParamsPathway
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewParamsPathway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection documenting operational waivers that have been issued for the Site
// associated with this record.
type SiteOperationNewParamsWaiver struct {
	// The expiration date of this waiver, in ISO8601 UTC format with millisecond
	// precision.
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	// Boolean indicating whether or not this waiver has expired.
	HasExpired param.Opt[bool] `json:"hasExpired,omitzero"`
	// The issue date of this waiver, in ISO8601 UTC format with millisecond precision.
	IssueDate param.Opt[time.Time] `json:"issueDate,omitzero" format:"date-time"`
	// The name of the person who issued this waiver.
	IssuerName param.Opt[string] `json:"issuerName,omitzero"`
	// The name of the person requesting this waiver.
	RequesterName param.Opt[string] `json:"requesterName,omitzero"`
	// The phone number of the person requesting this waiver.
	RequesterPhoneNumber param.Opt[string] `json:"requesterPhoneNumber,omitzero"`
	// The unit requesting this waiver.
	RequestingUnit param.Opt[string] `json:"requestingUnit,omitzero"`
	// Description of the entities to which this waiver applies.
	WaiverAppliesTo param.Opt[string] `json:"waiverAppliesTo,omitzero"`
	// The description of this waiver.
	WaiverDescription param.Opt[string] `json:"waiverDescription,omitzero"`
	// The name of the person who made the most recent change to this Waiver data.
	WaiverLastChangedBy param.Opt[string] `json:"waiverLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this waiver data, in ISO8601 UTC
	// format with millisecond precision.
	WaiverLastChangedDate param.Opt[time.Time] `json:"waiverLastChangedDate,omitzero" format:"date-time"`
	paramObj
}

func (r SiteOperationNewParamsWaiver) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewParamsWaiver
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewParamsWaiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteOperationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteOperationGetParams]'s query parameters as `url.Values`.
func (r SiteOperationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteOperationUpdateParams struct {
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
	DataMode SiteOperationUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The ID of the parent site.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The name of the person who made the most recent change to data in the
	// DailyOperations collection.
	DopsLastChangedBy param.Opt[string] `json:"dopsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to data in the DailyOperations
	// collection, in ISO 8601 UTC format with millisecond precision.
	DopsLastChangedDate param.Opt[time.Time] `json:"dopsLastChangedDate,omitzero" format:"date-time"`
	// The reason for the most recent change to data in the dailyOperations collection.
	DopsLastChangedReason param.Opt[string] `json:"dopsLastChangedReason,omitzero"`
	// Id of the associated launchSite entity.
	IDLaunchSite param.Opt[string] `json:"idLaunchSite,omitzero"`
	// The name of the person who made the most recent change to data in the
	// MaximumOnGrounds collection.
	MogsLastChangedBy param.Opt[string] `json:"mogsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to data in the MaximumOnGrounds
	// collection, in ISO 8601 UTC format with millisecond precision.
	MogsLastChangedDate param.Opt[time.Time] `json:"mogsLastChangedDate,omitzero" format:"date-time"`
	// The reason for the most recent change to data in the MaximumOnGrounds
	// collection.
	MogsLastChangedReason param.Opt[string] `json:"mogsLastChangedReason,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Collection providing hours of operation and other information specific to a day
	// of the week.
	DailyOperations []SiteOperationUpdateParamsDailyOperation `json:"dailyOperations,omitzero"`
	// Collection providing maximum on ground (MOG) information for specific aircraft
	// at the site associated with this SiteOperations record.
	MaximumOnGrounds []SiteOperationUpdateParamsMaximumOnGround `json:"maximumOnGrounds,omitzero"`
	// Collection providing relevant information in the event of deviations/exceptions
	// to normal operations.
	OperationalDeviations []SiteOperationUpdateParamsOperationalDeviation `json:"operationalDeviations,omitzero"`
	// Collection of planning information associated with this SiteOperations record.
	OperationalPlannings []SiteOperationUpdateParamsOperationalPlanning `json:"operationalPlannings,omitzero"`
	// Collection detailing operational pathways at the Site associated with this
	// SiteOperations record.
	Pathways []SiteOperationUpdateParamsPathway `json:"pathways,omitzero"`
	// Collection documenting operational waivers that have been issued for the Site
	// associated with this record.
	Waivers []SiteOperationUpdateParamsWaiver `json:"waivers,omitzero"`
	paramObj
}

func (r SiteOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUpdateParams) UnmarshalJSON(data []byte) error {
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
type SiteOperationUpdateParamsDataMode string

const (
	SiteOperationUpdateParamsDataModeReal      SiteOperationUpdateParamsDataMode = "REAL"
	SiteOperationUpdateParamsDataModeTest      SiteOperationUpdateParamsDataMode = "TEST"
	SiteOperationUpdateParamsDataModeSimulated SiteOperationUpdateParamsDataMode = "SIMULATED"
	SiteOperationUpdateParamsDataModeExercise  SiteOperationUpdateParamsDataMode = "EXERCISE"
)

// Collection providing hours of operation and other information specific to a day
// of the week.
type SiteOperationUpdateParamsDailyOperation struct {
	// The name or type of operation to which this information pertains.
	OperationName param.Opt[string] `json:"operationName,omitzero"`
	// The name of the person who made the most recent change to this DailyOperation
	// data.
	OphrsLastChangedBy param.Opt[string] `json:"ophrsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this DailyOperation data, in ISO
	// 8601 UTC format with millisecond precision.
	OphrsLastChangedDate param.Opt[time.Time] `json:"ophrsLastChangedDate,omitzero" format:"date-time"`
	// The day of the week to which this operational information pertains.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DayOfWeek string `json:"dayOfWeek,omitzero"`
	// A collection containing the operational start and stop times scheduled for the
	// day of the week specified.
	OperatingHours []SiteOperationUpdateParamsDailyOperationOperatingHour `json:"operatingHours,omitzero"`
	paramObj
}

func (r SiteOperationUpdateParamsDailyOperation) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUpdateParamsDailyOperation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUpdateParamsDailyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SiteOperationUpdateParamsDailyOperation](
		"dayOfWeek", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY",
	)
}

// A collection containing the operational start and stop times scheduled for the
// day of the week specified.
type SiteOperationUpdateParamsDailyOperationOperatingHour struct {
	// The Zulu (UTC) operational start time, expressed in ISO 8601 format as HH:MM.
	OpStartTime param.Opt[string] `json:"opStartTime,omitzero"`
	// The Zulu (UTC) operational stop time, expressed in ISO 8601 format as HH:MM.
	OpStopTime param.Opt[string] `json:"opStopTime,omitzero"`
	paramObj
}

func (r SiteOperationUpdateParamsDailyOperationOperatingHour) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUpdateParamsDailyOperationOperatingHour
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUpdateParamsDailyOperationOperatingHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing maximum on ground (MOG) information for specific aircraft
// at the site associated with this SiteOperations record.
type SiteOperationUpdateParamsMaximumOnGround struct {
	// The Model Design Series (MDS) designation of the aircraft to which this maximum
	// on ground (MOG) data pertains.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Maximum on ground (MOG) number of contingent aircraft based on spacing and
	// manpower, for the aircraft type specified.
	ContingencyMog param.Opt[int64] `json:"contingencyMOG,omitzero"`
	// The name of the person who made the most recent change to this maximum on ground
	// data.
	MogLastChangedBy param.Opt[string] `json:"mogLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this maximum on ground data, in
	// ISO 8601 UTC format with millisecond precision.
	MogLastChangedDate param.Opt[time.Time] `json:"mogLastChangedDate,omitzero" format:"date-time"`
	// Maximum on ground (MOG) number of parking wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideParkingMog param.Opt[int64] `json:"wideParkingMOG,omitzero"`
	// Maximum on ground (MOG) number of working wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideWorkingMog param.Opt[int64] `json:"wideWorkingMOG,omitzero"`
	paramObj
}

func (r SiteOperationUpdateParamsMaximumOnGround) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUpdateParamsMaximumOnGround
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUpdateParamsMaximumOnGround) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing relevant information in the event of deviations/exceptions
// to normal operations.
type SiteOperationUpdateParamsOperationalDeviation struct {
	// The Model Design Series (MDS) designation of the aircraft affected by this
	// operational deviation.
	AffectedAircraftMds param.Opt[string] `json:"affectedAircraftMDS,omitzero"`
	// The maximum on ground (MOG) number for aircraft affected by this operational
	// deviation.
	AffectedMog param.Opt[int64] `json:"affectedMOG,omitzero"`
	// On ground time for aircraft affected by this operational deviation.
	AircraftOnGroundTime param.Opt[string] `json:"aircraftOnGroundTime,omitzero"`
	// Rest time for crew affected by this operational deviation.
	CrewRestTime param.Opt[string] `json:"crewRestTime,omitzero"`
	// The name of the person who made the most recent change to this
	// OperationalDeviation data.
	OdLastChangedBy param.Opt[string] `json:"odLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this OperationalDeviation data,
	// in ISO 8601 UTC format with millisecond precision.
	OdLastChangedDate param.Opt[time.Time] `json:"odLastChangedDate,omitzero" format:"date-time"`
	// Text remark regarding this operational deviation.
	OdRemark param.Opt[string] `json:"odRemark,omitzero"`
	paramObj
}

func (r SiteOperationUpdateParamsOperationalDeviation) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUpdateParamsOperationalDeviation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUpdateParamsOperationalDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of planning information associated with this SiteOperations record.
type SiteOperationUpdateParamsOperationalPlanning struct {
	// The end date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpEndDate param.Opt[time.Time] `json:"opEndDate,omitzero" format:"date-time"`
	// The name of the person who made the most recent change made to this
	// OperationalPlanning data.
	OpLastChangedBy param.Opt[string] `json:"opLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this OperationalPlanning data, in
	// ISO8601 UTC format with millisecond precision.
	OpLastChangedDate param.Opt[time.Time] `json:"opLastChangedDate,omitzero" format:"date-time"`
	// Remark text regarding this operation planning.
	OpRemark param.Opt[string] `json:"opRemark,omitzero"`
	// The person, unit, organization, etc. responsible for this operation planning.
	OpSource param.Opt[string] `json:"opSource,omitzero"`
	// The start date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpStartDate param.Opt[time.Time] `json:"opStartDate,omitzero" format:"date-time"`
	// The status of this operational planning.
	OpStatus param.Opt[string] `json:"opStatus,omitzero"`
	paramObj
}

func (r SiteOperationUpdateParamsOperationalPlanning) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUpdateParamsOperationalPlanning
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUpdateParamsOperationalPlanning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection detailing operational pathways at the Site associated with this
// SiteOperations record.
type SiteOperationUpdateParamsPathway struct {
	// Text defining this pathway from its constituent parts.
	PwDefinition param.Opt[string] `json:"pwDefinition,omitzero"`
	// The name of the person who made the most recent change to this Pathway data.
	PwLastChangedBy param.Opt[string] `json:"pwLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this Pathway data, in ISO 8601
	// UTC format with millisecond precision.
	PwLastChangedDate param.Opt[time.Time] `json:"pwLastChangedDate,omitzero" format:"date-time"`
	// The type of paths that constitute this pathway.
	PwType param.Opt[string] `json:"pwType,omitzero"`
	// The intended use of this pathway.
	PwUsage param.Opt[string] `json:"pwUsage,omitzero"`
	paramObj
}

func (r SiteOperationUpdateParamsPathway) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUpdateParamsPathway
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUpdateParamsPathway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection documenting operational waivers that have been issued for the Site
// associated with this record.
type SiteOperationUpdateParamsWaiver struct {
	// The expiration date of this waiver, in ISO8601 UTC format with millisecond
	// precision.
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	// Boolean indicating whether or not this waiver has expired.
	HasExpired param.Opt[bool] `json:"hasExpired,omitzero"`
	// The issue date of this waiver, in ISO8601 UTC format with millisecond precision.
	IssueDate param.Opt[time.Time] `json:"issueDate,omitzero" format:"date-time"`
	// The name of the person who issued this waiver.
	IssuerName param.Opt[string] `json:"issuerName,omitzero"`
	// The name of the person requesting this waiver.
	RequesterName param.Opt[string] `json:"requesterName,omitzero"`
	// The phone number of the person requesting this waiver.
	RequesterPhoneNumber param.Opt[string] `json:"requesterPhoneNumber,omitzero"`
	// The unit requesting this waiver.
	RequestingUnit param.Opt[string] `json:"requestingUnit,omitzero"`
	// Description of the entities to which this waiver applies.
	WaiverAppliesTo param.Opt[string] `json:"waiverAppliesTo,omitzero"`
	// The description of this waiver.
	WaiverDescription param.Opt[string] `json:"waiverDescription,omitzero"`
	// The name of the person who made the most recent change to this Waiver data.
	WaiverLastChangedBy param.Opt[string] `json:"waiverLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this waiver data, in ISO8601 UTC
	// format with millisecond precision.
	WaiverLastChangedDate param.Opt[time.Time] `json:"waiverLastChangedDate,omitzero" format:"date-time"`
	paramObj
}

func (r SiteOperationUpdateParamsWaiver) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUpdateParamsWaiver
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUpdateParamsWaiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteOperationListParams struct {
	// The ID of the parent site.
	IDSite      string           `query:"idSite,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteOperationListParams]'s query parameters as
// `url.Values`.
func (r SiteOperationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteOperationCountParams struct {
	// The ID of the parent site.
	IDSite      string           `query:"idSite,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteOperationCountParams]'s query parameters as
// `url.Values`.
func (r SiteOperationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteOperationNewBulkParams struct {
	Body []SiteOperationNewBulkParamsBody
	paramObj
}

func (r SiteOperationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *SiteOperationNewBulkParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// Site operating details concerning the hours of operation, operational
// limitations, site navigation, and waivers associated with the Site.
//
// The properties ClassificationMarking, DataMode, IDSite, Source are required.
type SiteOperationNewBulkParamsBody struct {
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
	// The ID of the parent site.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The name of the person who made the most recent change to data in the
	// DailyOperations collection.
	DopsLastChangedBy param.Opt[string] `json:"dopsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to data in the DailyOperations
	// collection, in ISO 8601 UTC format with millisecond precision.
	DopsLastChangedDate param.Opt[time.Time] `json:"dopsLastChangedDate,omitzero" format:"date-time"`
	// The reason for the most recent change to data in the dailyOperations collection.
	DopsLastChangedReason param.Opt[string] `json:"dopsLastChangedReason,omitzero"`
	// Id of the associated launchSite entity.
	IDLaunchSite param.Opt[string] `json:"idLaunchSite,omitzero"`
	// The name of the person who made the most recent change to data in the
	// MaximumOnGrounds collection.
	MogsLastChangedBy param.Opt[string] `json:"mogsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to data in the MaximumOnGrounds
	// collection, in ISO 8601 UTC format with millisecond precision.
	MogsLastChangedDate param.Opt[time.Time] `json:"mogsLastChangedDate,omitzero" format:"date-time"`
	// The reason for the most recent change to data in the MaximumOnGrounds
	// collection.
	MogsLastChangedReason param.Opt[string] `json:"mogsLastChangedReason,omitzero"`
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
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Collection providing hours of operation and other information specific to a day
	// of the week.
	DailyOperations []SiteOperationNewBulkParamsBodyDailyOperation `json:"dailyOperations,omitzero"`
	// Collection providing maximum on ground (MOG) information for specific aircraft
	// at the site associated with this SiteOperations record.
	MaximumOnGrounds []SiteOperationNewBulkParamsBodyMaximumOnGround `json:"maximumOnGrounds,omitzero"`
	// Collection providing relevant information in the event of deviations/exceptions
	// to normal operations.
	OperationalDeviations []SiteOperationNewBulkParamsBodyOperationalDeviation `json:"operationalDeviations,omitzero"`
	// Collection of planning information associated with this SiteOperations record.
	OperationalPlannings []SiteOperationNewBulkParamsBodyOperationalPlanning `json:"operationalPlannings,omitzero"`
	// Collection detailing operational pathways at the Site associated with this
	// SiteOperations record.
	Pathways []SiteOperationNewBulkParamsBodyPathway `json:"pathways,omitzero"`
	// Collection documenting operational waivers that have been issued for the Site
	// associated with this record.
	Waivers []SiteOperationNewBulkParamsBodyWaiver `json:"waivers,omitzero"`
	paramObj
}

func (r SiteOperationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SiteOperationNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection providing hours of operation and other information specific to a day
// of the week.
type SiteOperationNewBulkParamsBodyDailyOperation struct {
	// The name or type of operation to which this information pertains.
	OperationName param.Opt[string] `json:"operationName,omitzero"`
	// The name of the person who made the most recent change to this DailyOperation
	// data.
	OphrsLastChangedBy param.Opt[string] `json:"ophrsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this DailyOperation data, in ISO
	// 8601 UTC format with millisecond precision.
	OphrsLastChangedDate param.Opt[time.Time] `json:"ophrsLastChangedDate,omitzero" format:"date-time"`
	// The day of the week to which this operational information pertains.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DayOfWeek string `json:"dayOfWeek,omitzero"`
	// A collection containing the operational start and stop times scheduled for the
	// day of the week specified.
	OperatingHours []SiteOperationNewBulkParamsBodyDailyOperationOperatingHour `json:"operatingHours,omitzero"`
	paramObj
}

func (r SiteOperationNewBulkParamsBodyDailyOperation) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewBulkParamsBodyDailyOperation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewBulkParamsBodyDailyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SiteOperationNewBulkParamsBodyDailyOperation](
		"dayOfWeek", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY",
	)
}

// A collection containing the operational start and stop times scheduled for the
// day of the week specified.
type SiteOperationNewBulkParamsBodyDailyOperationOperatingHour struct {
	// The Zulu (UTC) operational start time, expressed in ISO 8601 format as HH:MM.
	OpStartTime param.Opt[string] `json:"opStartTime,omitzero"`
	// The Zulu (UTC) operational stop time, expressed in ISO 8601 format as HH:MM.
	OpStopTime param.Opt[string] `json:"opStopTime,omitzero"`
	paramObj
}

func (r SiteOperationNewBulkParamsBodyDailyOperationOperatingHour) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewBulkParamsBodyDailyOperationOperatingHour
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewBulkParamsBodyDailyOperationOperatingHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing maximum on ground (MOG) information for specific aircraft
// at the site associated with this SiteOperations record.
type SiteOperationNewBulkParamsBodyMaximumOnGround struct {
	// The Model Design Series (MDS) designation of the aircraft to which this maximum
	// on ground (MOG) data pertains.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Maximum on ground (MOG) number of contingent aircraft based on spacing and
	// manpower, for the aircraft type specified.
	ContingencyMog param.Opt[int64] `json:"contingencyMOG,omitzero"`
	// The name of the person who made the most recent change to this maximum on ground
	// data.
	MogLastChangedBy param.Opt[string] `json:"mogLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this maximum on ground data, in
	// ISO 8601 UTC format with millisecond precision.
	MogLastChangedDate param.Opt[time.Time] `json:"mogLastChangedDate,omitzero" format:"date-time"`
	// Maximum on ground (MOG) number of parking wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideParkingMog param.Opt[int64] `json:"wideParkingMOG,omitzero"`
	// Maximum on ground (MOG) number of working wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideWorkingMog param.Opt[int64] `json:"wideWorkingMOG,omitzero"`
	paramObj
}

func (r SiteOperationNewBulkParamsBodyMaximumOnGround) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewBulkParamsBodyMaximumOnGround
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewBulkParamsBodyMaximumOnGround) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing relevant information in the event of deviations/exceptions
// to normal operations.
type SiteOperationNewBulkParamsBodyOperationalDeviation struct {
	// The Model Design Series (MDS) designation of the aircraft affected by this
	// operational deviation.
	AffectedAircraftMds param.Opt[string] `json:"affectedAircraftMDS,omitzero"`
	// The maximum on ground (MOG) number for aircraft affected by this operational
	// deviation.
	AffectedMog param.Opt[int64] `json:"affectedMOG,omitzero"`
	// On ground time for aircraft affected by this operational deviation.
	AircraftOnGroundTime param.Opt[string] `json:"aircraftOnGroundTime,omitzero"`
	// Rest time for crew affected by this operational deviation.
	CrewRestTime param.Opt[string] `json:"crewRestTime,omitzero"`
	// The name of the person who made the most recent change to this
	// OperationalDeviation data.
	OdLastChangedBy param.Opt[string] `json:"odLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this OperationalDeviation data,
	// in ISO 8601 UTC format with millisecond precision.
	OdLastChangedDate param.Opt[time.Time] `json:"odLastChangedDate,omitzero" format:"date-time"`
	// Text remark regarding this operational deviation.
	OdRemark param.Opt[string] `json:"odRemark,omitzero"`
	paramObj
}

func (r SiteOperationNewBulkParamsBodyOperationalDeviation) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewBulkParamsBodyOperationalDeviation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewBulkParamsBodyOperationalDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of planning information associated with this SiteOperations record.
type SiteOperationNewBulkParamsBodyOperationalPlanning struct {
	// The end date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpEndDate param.Opt[time.Time] `json:"opEndDate,omitzero" format:"date-time"`
	// The name of the person who made the most recent change made to this
	// OperationalPlanning data.
	OpLastChangedBy param.Opt[string] `json:"opLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this OperationalPlanning data, in
	// ISO8601 UTC format with millisecond precision.
	OpLastChangedDate param.Opt[time.Time] `json:"opLastChangedDate,omitzero" format:"date-time"`
	// Remark text regarding this operation planning.
	OpRemark param.Opt[string] `json:"opRemark,omitzero"`
	// The person, unit, organization, etc. responsible for this operation planning.
	OpSource param.Opt[string] `json:"opSource,omitzero"`
	// The start date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpStartDate param.Opt[time.Time] `json:"opStartDate,omitzero" format:"date-time"`
	// The status of this operational planning.
	OpStatus param.Opt[string] `json:"opStatus,omitzero"`
	paramObj
}

func (r SiteOperationNewBulkParamsBodyOperationalPlanning) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewBulkParamsBodyOperationalPlanning
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewBulkParamsBodyOperationalPlanning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection detailing operational pathways at the Site associated with this
// SiteOperations record.
type SiteOperationNewBulkParamsBodyPathway struct {
	// Text defining this pathway from its constituent parts.
	PwDefinition param.Opt[string] `json:"pwDefinition,omitzero"`
	// The name of the person who made the most recent change to this Pathway data.
	PwLastChangedBy param.Opt[string] `json:"pwLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this Pathway data, in ISO 8601
	// UTC format with millisecond precision.
	PwLastChangedDate param.Opt[time.Time] `json:"pwLastChangedDate,omitzero" format:"date-time"`
	// The type of paths that constitute this pathway.
	PwType param.Opt[string] `json:"pwType,omitzero"`
	// The intended use of this pathway.
	PwUsage param.Opt[string] `json:"pwUsage,omitzero"`
	paramObj
}

func (r SiteOperationNewBulkParamsBodyPathway) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewBulkParamsBodyPathway
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewBulkParamsBodyPathway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection documenting operational waivers that have been issued for the Site
// associated with this record.
type SiteOperationNewBulkParamsBodyWaiver struct {
	// The expiration date of this waiver, in ISO8601 UTC format with millisecond
	// precision.
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	// Boolean indicating whether or not this waiver has expired.
	HasExpired param.Opt[bool] `json:"hasExpired,omitzero"`
	// The issue date of this waiver, in ISO8601 UTC format with millisecond precision.
	IssueDate param.Opt[time.Time] `json:"issueDate,omitzero" format:"date-time"`
	// The name of the person who issued this waiver.
	IssuerName param.Opt[string] `json:"issuerName,omitzero"`
	// The name of the person requesting this waiver.
	RequesterName param.Opt[string] `json:"requesterName,omitzero"`
	// The phone number of the person requesting this waiver.
	RequesterPhoneNumber param.Opt[string] `json:"requesterPhoneNumber,omitzero"`
	// The unit requesting this waiver.
	RequestingUnit param.Opt[string] `json:"requestingUnit,omitzero"`
	// Description of the entities to which this waiver applies.
	WaiverAppliesTo param.Opt[string] `json:"waiverAppliesTo,omitzero"`
	// The description of this waiver.
	WaiverDescription param.Opt[string] `json:"waiverDescription,omitzero"`
	// The name of the person who made the most recent change to this Waiver data.
	WaiverLastChangedBy param.Opt[string] `json:"waiverLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this waiver data, in ISO8601 UTC
	// format with millisecond precision.
	WaiverLastChangedDate param.Opt[time.Time] `json:"waiverLastChangedDate,omitzero" format:"date-time"`
	paramObj
}

func (r SiteOperationNewBulkParamsBodyWaiver) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationNewBulkParamsBodyWaiver
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationNewBulkParamsBodyWaiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteOperationTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The ID of the parent site.
	IDSite      string           `query:"idSite,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteOperationTupleParams]'s query parameters as
// `url.Values`.
func (r SiteOperationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteOperationUnvalidatedPublishParams struct {
	Body []SiteOperationUnvalidatedPublishParamsBody
	paramObj
}

func (r SiteOperationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *SiteOperationUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// Site operating details concerning the hours of operation, operational
// limitations, site navigation, and waivers associated with the Site.
//
// The properties ClassificationMarking, DataMode, IDSite, Source are required.
type SiteOperationUnvalidatedPublishParamsBody struct {
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
	// The ID of the parent site.
	IDSite string `json:"idSite,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The name of the person who made the most recent change to data in the
	// DailyOperations collection.
	DopsLastChangedBy param.Opt[string] `json:"dopsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to data in the DailyOperations
	// collection, in ISO 8601 UTC format with millisecond precision.
	DopsLastChangedDate param.Opt[time.Time] `json:"dopsLastChangedDate,omitzero" format:"date-time"`
	// The reason for the most recent change to data in the dailyOperations collection.
	DopsLastChangedReason param.Opt[string] `json:"dopsLastChangedReason,omitzero"`
	// Id of the associated launchSite entity.
	IDLaunchSite param.Opt[string] `json:"idLaunchSite,omitzero"`
	// The name of the person who made the most recent change to data in the
	// MaximumOnGrounds collection.
	MogsLastChangedBy param.Opt[string] `json:"mogsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to data in the MaximumOnGrounds
	// collection, in ISO 8601 UTC format with millisecond precision.
	MogsLastChangedDate param.Opt[time.Time] `json:"mogsLastChangedDate,omitzero" format:"date-time"`
	// The reason for the most recent change to data in the MaximumOnGrounds
	// collection.
	MogsLastChangedReason param.Opt[string] `json:"mogsLastChangedReason,omitzero"`
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
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Collection providing hours of operation and other information specific to a day
	// of the week.
	DailyOperations []SiteOperationUnvalidatedPublishParamsBodyDailyOperation `json:"dailyOperations,omitzero"`
	// Collection providing maximum on ground (MOG) information for specific aircraft
	// at the site associated with this SiteOperations record.
	MaximumOnGrounds []SiteOperationUnvalidatedPublishParamsBodyMaximumOnGround `json:"maximumOnGrounds,omitzero"`
	// Collection providing relevant information in the event of deviations/exceptions
	// to normal operations.
	OperationalDeviations []SiteOperationUnvalidatedPublishParamsBodyOperationalDeviation `json:"operationalDeviations,omitzero"`
	// Collection of planning information associated with this SiteOperations record.
	OperationalPlannings []SiteOperationUnvalidatedPublishParamsBodyOperationalPlanning `json:"operationalPlannings,omitzero"`
	// Collection detailing operational pathways at the Site associated with this
	// SiteOperations record.
	Pathways []SiteOperationUnvalidatedPublishParamsBodyPathway `json:"pathways,omitzero"`
	// Collection documenting operational waivers that have been issued for the Site
	// associated with this record.
	Waivers []SiteOperationUnvalidatedPublishParamsBodyWaiver `json:"waivers,omitzero"`
	paramObj
}

func (r SiteOperationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SiteOperationUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

// Collection providing hours of operation and other information specific to a day
// of the week.
type SiteOperationUnvalidatedPublishParamsBodyDailyOperation struct {
	// The name or type of operation to which this information pertains.
	OperationName param.Opt[string] `json:"operationName,omitzero"`
	// The name of the person who made the most recent change to this DailyOperation
	// data.
	OphrsLastChangedBy param.Opt[string] `json:"ophrsLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this DailyOperation data, in ISO
	// 8601 UTC format with millisecond precision.
	OphrsLastChangedDate param.Opt[time.Time] `json:"ophrsLastChangedDate,omitzero" format:"date-time"`
	// The day of the week to which this operational information pertains.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DayOfWeek string `json:"dayOfWeek,omitzero"`
	// A collection containing the operational start and stop times scheduled for the
	// day of the week specified.
	OperatingHours []SiteOperationUnvalidatedPublishParamsBodyDailyOperationOperatingHour `json:"operatingHours,omitzero"`
	paramObj
}

func (r SiteOperationUnvalidatedPublishParamsBodyDailyOperation) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUnvalidatedPublishParamsBodyDailyOperation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUnvalidatedPublishParamsBodyDailyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SiteOperationUnvalidatedPublishParamsBodyDailyOperation](
		"dayOfWeek", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY",
	)
}

// A collection containing the operational start and stop times scheduled for the
// day of the week specified.
type SiteOperationUnvalidatedPublishParamsBodyDailyOperationOperatingHour struct {
	// The Zulu (UTC) operational start time, expressed in ISO 8601 format as HH:MM.
	OpStartTime param.Opt[string] `json:"opStartTime,omitzero"`
	// The Zulu (UTC) operational stop time, expressed in ISO 8601 format as HH:MM.
	OpStopTime param.Opt[string] `json:"opStopTime,omitzero"`
	paramObj
}

func (r SiteOperationUnvalidatedPublishParamsBodyDailyOperationOperatingHour) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUnvalidatedPublishParamsBodyDailyOperationOperatingHour
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUnvalidatedPublishParamsBodyDailyOperationOperatingHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing maximum on ground (MOG) information for specific aircraft
// at the site associated with this SiteOperations record.
type SiteOperationUnvalidatedPublishParamsBodyMaximumOnGround struct {
	// The Model Design Series (MDS) designation of the aircraft to which this maximum
	// on ground (MOG) data pertains.
	AircraftMds param.Opt[string] `json:"aircraftMDS,omitzero"`
	// Maximum on ground (MOG) number of contingent aircraft based on spacing and
	// manpower, for the aircraft type specified.
	ContingencyMog param.Opt[int64] `json:"contingencyMOG,omitzero"`
	// The name of the person who made the most recent change to this maximum on ground
	// data.
	MogLastChangedBy param.Opt[string] `json:"mogLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this maximum on ground data, in
	// ISO 8601 UTC format with millisecond precision.
	MogLastChangedDate param.Opt[time.Time] `json:"mogLastChangedDate,omitzero" format:"date-time"`
	// Maximum on ground (MOG) number of parking wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideParkingMog param.Opt[int64] `json:"wideParkingMOG,omitzero"`
	// Maximum on ground (MOG) number of working wide-body aircraft based on spacing
	// and manpower, for the aircraft type specified.
	WideWorkingMog param.Opt[int64] `json:"wideWorkingMOG,omitzero"`
	paramObj
}

func (r SiteOperationUnvalidatedPublishParamsBodyMaximumOnGround) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUnvalidatedPublishParamsBodyMaximumOnGround
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUnvalidatedPublishParamsBodyMaximumOnGround) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection providing relevant information in the event of deviations/exceptions
// to normal operations.
type SiteOperationUnvalidatedPublishParamsBodyOperationalDeviation struct {
	// The Model Design Series (MDS) designation of the aircraft affected by this
	// operational deviation.
	AffectedAircraftMds param.Opt[string] `json:"affectedAircraftMDS,omitzero"`
	// The maximum on ground (MOG) number for aircraft affected by this operational
	// deviation.
	AffectedMog param.Opt[int64] `json:"affectedMOG,omitzero"`
	// On ground time for aircraft affected by this operational deviation.
	AircraftOnGroundTime param.Opt[string] `json:"aircraftOnGroundTime,omitzero"`
	// Rest time for crew affected by this operational deviation.
	CrewRestTime param.Opt[string] `json:"crewRestTime,omitzero"`
	// The name of the person who made the most recent change to this
	// OperationalDeviation data.
	OdLastChangedBy param.Opt[string] `json:"odLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this OperationalDeviation data,
	// in ISO 8601 UTC format with millisecond precision.
	OdLastChangedDate param.Opt[time.Time] `json:"odLastChangedDate,omitzero" format:"date-time"`
	// Text remark regarding this operational deviation.
	OdRemark param.Opt[string] `json:"odRemark,omitzero"`
	paramObj
}

func (r SiteOperationUnvalidatedPublishParamsBodyOperationalDeviation) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUnvalidatedPublishParamsBodyOperationalDeviation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUnvalidatedPublishParamsBodyOperationalDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection of planning information associated with this SiteOperations record.
type SiteOperationUnvalidatedPublishParamsBodyOperationalPlanning struct {
	// The end date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpEndDate param.Opt[time.Time] `json:"opEndDate,omitzero" format:"date-time"`
	// The name of the person who made the most recent change made to this
	// OperationalPlanning data.
	OpLastChangedBy param.Opt[string] `json:"opLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this OperationalPlanning data, in
	// ISO8601 UTC format with millisecond precision.
	OpLastChangedDate param.Opt[time.Time] `json:"opLastChangedDate,omitzero" format:"date-time"`
	// Remark text regarding this operation planning.
	OpRemark param.Opt[string] `json:"opRemark,omitzero"`
	// The person, unit, organization, etc. responsible for this operation planning.
	OpSource param.Opt[string] `json:"opSource,omitzero"`
	// The start date of this operational planning, in ISO8601 UTC format with
	// millisecond precision.
	OpStartDate param.Opt[time.Time] `json:"opStartDate,omitzero" format:"date-time"`
	// The status of this operational planning.
	OpStatus param.Opt[string] `json:"opStatus,omitzero"`
	paramObj
}

func (r SiteOperationUnvalidatedPublishParamsBodyOperationalPlanning) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUnvalidatedPublishParamsBodyOperationalPlanning
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUnvalidatedPublishParamsBodyOperationalPlanning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection detailing operational pathways at the Site associated with this
// SiteOperations record.
type SiteOperationUnvalidatedPublishParamsBodyPathway struct {
	// Text defining this pathway from its constituent parts.
	PwDefinition param.Opt[string] `json:"pwDefinition,omitzero"`
	// The name of the person who made the most recent change to this Pathway data.
	PwLastChangedBy param.Opt[string] `json:"pwLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this Pathway data, in ISO 8601
	// UTC format with millisecond precision.
	PwLastChangedDate param.Opt[time.Time] `json:"pwLastChangedDate,omitzero" format:"date-time"`
	// The type of paths that constitute this pathway.
	PwType param.Opt[string] `json:"pwType,omitzero"`
	// The intended use of this pathway.
	PwUsage param.Opt[string] `json:"pwUsage,omitzero"`
	paramObj
}

func (r SiteOperationUnvalidatedPublishParamsBodyPathway) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUnvalidatedPublishParamsBodyPathway
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUnvalidatedPublishParamsBodyPathway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection documenting operational waivers that have been issued for the Site
// associated with this record.
type SiteOperationUnvalidatedPublishParamsBodyWaiver struct {
	// The expiration date of this waiver, in ISO8601 UTC format with millisecond
	// precision.
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	// Boolean indicating whether or not this waiver has expired.
	HasExpired param.Opt[bool] `json:"hasExpired,omitzero"`
	// The issue date of this waiver, in ISO8601 UTC format with millisecond precision.
	IssueDate param.Opt[time.Time] `json:"issueDate,omitzero" format:"date-time"`
	// The name of the person who issued this waiver.
	IssuerName param.Opt[string] `json:"issuerName,omitzero"`
	// The name of the person requesting this waiver.
	RequesterName param.Opt[string] `json:"requesterName,omitzero"`
	// The phone number of the person requesting this waiver.
	RequesterPhoneNumber param.Opt[string] `json:"requesterPhoneNumber,omitzero"`
	// The unit requesting this waiver.
	RequestingUnit param.Opt[string] `json:"requestingUnit,omitzero"`
	// Description of the entities to which this waiver applies.
	WaiverAppliesTo param.Opt[string] `json:"waiverAppliesTo,omitzero"`
	// The description of this waiver.
	WaiverDescription param.Opt[string] `json:"waiverDescription,omitzero"`
	// The name of the person who made the most recent change to this Waiver data.
	WaiverLastChangedBy param.Opt[string] `json:"waiverLastChangedBy,omitzero"`
	// The datetime of the most recent change made to this waiver data, in ISO8601 UTC
	// format with millisecond precision.
	WaiverLastChangedDate param.Opt[time.Time] `json:"waiverLastChangedDate,omitzero" format:"date-time"`
	paramObj
}

func (r SiteOperationUnvalidatedPublishParamsBodyWaiver) MarshalJSON() (data []byte, err error) {
	type shadow SiteOperationUnvalidatedPublishParamsBodyWaiver
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiteOperationUnvalidatedPublishParamsBodyWaiver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
